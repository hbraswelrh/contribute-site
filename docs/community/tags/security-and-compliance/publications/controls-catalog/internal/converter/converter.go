package converter

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/gemaraproj/go-gemara"
	"github.com/goccy/go-yaml"
)

//go:embed markdown.tmpl
var markdownTemplate string

// DefaultGroupFileOrder is the merge order for per-group YAML files (basename without .yaml).
var DefaultGroupFileOrder = []string{
	"access",
	"compute",
	"deploy",
	"develop",
	"distribute",
	"securing-artefacts",
	"securing-build-pipelines",
	"securing-materials",
	"securing-the-source-code",
	"security-assurance",
	"storage",
}

// catalogYAMLRoot matches the on-disk metadata.yaml shape (Gemara v1.0.0-rc.1–style).
type catalogYAMLRoot struct {
	Title    string `yaml:"title"`
	Type     string `yaml:"type"`
	Metadata struct {
		ID                  string                    `yaml:"id"`
		Type                string                    `yaml:"type"`
		GemaraVersion       string                    `yaml:"gemara-version"`
		Version             string                    `yaml:"version"`
		Description         string                    `yaml:"description"`
		Author              gemara.Actor              `yaml:"author"`
		MappingReferences   []gemara.MappingReference `yaml:"mapping-references"`
		ApplicabilityGroups []struct {
			ID          string `yaml:"id"`
			Title       string `yaml:"title"`
			Description string `yaml:"description"`
		} `yaml:"applicability-groups"`
	} `yaml:"metadata"`
}

type groupsYAML struct {
	Groups []gemara.Family `yaml:"groups"`
}

// guidelineYAML matches per-file guidelines (group, not family).
type guidelineYAML struct {
	ID              string   `yaml:"id"`
	Title           string   `yaml:"title"`
	Objective       string   `yaml:"objective"`
	State           string   `yaml:"state"`
	Group           string   `yaml:"group"`
	Applicability   []string `yaml:"applicability"`
	Recommendations []string `yaml:"recommendations"`
}

func (g guidelineYAML) toGemara() gemara.Guideline {
	return gemara.Guideline{
		Id:              g.ID,
		Title:           g.Title,
		Objective:       g.Objective,
		Family:          g.Group,
		Applicability:   g.Applicability,
		Recommendations: g.Recommendations,
	}
}

// ToGemara loads metadata.yaml, groups.yaml, and per-group guideline files into a GuidanceCatalog.
// groupFileOrder lists per-group file basenames (without .yaml); use DefaultGroupFileOrder for normal builds.
func ToGemara(catalogDir string, groupFileOrder []string) (*gemara.GuidanceCatalog, error) {
	metadataPath := filepath.Join(catalogDir, "metadata.yaml")
	rawMeta, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("read metadata.yaml: %w", err)
	}

	var root catalogYAMLRoot
	if err := yaml.Unmarshal(rawMeta, &root); err != nil {
		return nil, fmt.Errorf("parse metadata.yaml: %w", err)
	}

	md := gemara.Metadata{
		Id:                root.Metadata.ID,
		Version:           root.Metadata.Version,
		Description:       root.Metadata.Description,
		Author:            root.Metadata.Author,
		MappingReferences: root.Metadata.MappingReferences,
	}
	for _, ag := range root.Metadata.ApplicabilityGroups {
		md.ApplicabilityCategories = append(md.ApplicabilityCategories, gemara.Category{
			Id:          ag.ID,
			Title:       ag.Title,
			Description: ag.Description,
		})
	}

	groupsPath := filepath.Join(catalogDir, "groups.yaml")
	groupsData, err := os.ReadFile(groupsPath)
	if err != nil {
		return nil, fmt.Errorf("read groups.yaml: %w", err)
	}
	var groupsDoc groupsYAML
	if err := yaml.Unmarshal(groupsData, &groupsDoc); err != nil {
		return nil, fmt.Errorf("parse groups.yaml: %w", err)
	}

	var allGuidelines []gemara.Guideline
	for _, slug := range groupFileOrder {
		p := filepath.Join(catalogDir, fmt.Sprintf("%s.yaml", slug))
		if _, err := os.Stat(p); os.IsNotExist(err) {
			continue
		}
		data, err := os.ReadFile(p)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", p, err)
		}
		var file struct {
			Guidelines []guidelineYAML `yaml:"guidelines"`
		}
		if err := yaml.Unmarshal(data, &file); err != nil {
			return nil, fmt.Errorf("parse %s: %w", p, err)
		}
		for _, g := range file.Guidelines {
			allGuidelines = append(allGuidelines, g.toGemara())
		}
	}

	sort.Slice(allGuidelines, func(i, j int) bool {
		return allGuidelines[i].Id < allGuidelines[j].Id
	})

	return &gemara.GuidanceCatalog{
		Title:        root.Title,
		GuidanceType: gemara.GuidanceType(root.Type),
		Metadata:     md,
		Families:     groupsDoc.Groups,
		Guidelines:   allGuidelines,
	}, nil
}

// familyWithGuidelines represents a group (family in go-gemara) with its guidelines for template rendering.
type familyWithGuidelines struct {
	gemara.Family
	Guidelines []gemara.Guideline
}

// templateData holds the data structure for the markdown template.
type templateData struct {
	GroupsWithGuidelines []familyWithGuidelines
	ApplicabilityTitles  map[string]string
}

// ToMarkdown converts a GuidanceCatalog to Markdown format for website rendering.
func ToMarkdown(doc *gemara.GuidanceCatalog, catalogDir string) (string, error) {
	byGroup := make(map[string][]gemara.Guideline)
	for _, guideline := range doc.Guidelines {
		if guideline.Family != "" {
			byGroup[guideline.Family] = append(byGroup[guideline.Family], guideline)
		}
	}
	for gid := range byGroup {
		sort.Slice(byGroup[gid], func(i, j int) bool {
			return byGroup[gid][i].Id < byGroup[gid][j].Id
		})
	}

	var groupsWithGuidelines []familyWithGuidelines
	for _, grp := range doc.Families {
		gl, ok := byGroup[grp.Id]
		if ok && len(gl) > 0 {
			groupsWithGuidelines = append(groupsWithGuidelines, familyWithGuidelines{
				Family:     grp,
				Guidelines: gl,
			})
		}
	}

	applicabilityTitles := make(map[string]string)
	for _, cat := range doc.Metadata.ApplicabilityCategories {
		applicabilityTitles[cat.Id] = cat.Title
	}
	if catalogDir != "" {
		if err := mergeApplicabilityGroupTitles(filepath.Join(catalogDir, "metadata.yaml"), applicabilityTitles); err != nil {
			return "", err
		}
	}

	data := templateData{
		GroupsWithGuidelines: groupsWithGuidelines,
		ApplicabilityTitles:  applicabilityTitles,
	}

	tmpl, err := template.New("markdown").Funcs(template.FuncMap{
		"anchor": func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
		"applicabilityTitle": func(id string, titles map[string]string) string {
			if title, ok := titles[id]; ok {
				return title
			}
			return id
		},
	}).Parse(markdownTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var b strings.Builder
	if err := tmpl.Execute(&b, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return b.String(), nil
}

// metadataApplicabilityGroups mirrors optional metadata.applicability-groups in metadata.yaml.
type metadataApplicabilityGroups struct {
	Metadata struct {
		ApplicabilityGroups []struct {
			Id    string `yaml:"id"`
			Title string `yaml:"title"`
		} `yaml:"applicability-groups"`
	} `yaml:"metadata"`
}

func mergeApplicabilityGroupTitles(metadataPath string, titles map[string]string) error {
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("read metadata for applicability-groups: %w", err)
	}
	var ext metadataApplicabilityGroups
	if err := yaml.Unmarshal(data, &ext); err != nil {
		return fmt.Errorf("parse applicability-groups: %w", err)
	}
	for _, g := range ext.Metadata.ApplicabilityGroups {
		if g.Id == "" {
			continue
		}
		if _, exists := titles[g.Id]; !exists && g.Title != "" {
			titles[g.Id] = g.Title
		}
	}
	return nil
}

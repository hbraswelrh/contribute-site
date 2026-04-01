package converter

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

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

// Actor is metadata.author (Gemara-style Layer 1 catalog).
type Actor struct {
	Id   string `yaml:"id"`
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// MappingReference is an external framework reference (e.g. NIST).
type MappingReference struct {
	Id          string `yaml:"id"`
	Title       string `yaml:"title"`
	Version     string `yaml:"version"`
	URL         string `yaml:"url,omitempty"`
	Description string `yaml:"description,omitempty"`
}

// Category is an applicability group (originating document).
type Category struct {
	Id          string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description,omitempty"`
}

// Metadata holds catalog-level metadata for templates and merged YAML export.
type Metadata struct {
	Id                      string             `yaml:"id"`
	Type                    string             `yaml:"type,omitempty"`
	Version                 string             `yaml:"version"`
	GemaraVersion           string             `yaml:"gemara-version,omitempty"`
	Description             string             `yaml:"description"`
	Author                  Actor              `yaml:"author"`
	MappingReferences       []MappingReference `yaml:"mapping-references,omitempty"`
	ApplicabilityCategories []Category         `yaml:"applicability-categories,omitempty"`
}

// Family is one security domain group from groups.yaml (shown as a “family” on the site).
type Family struct {
	Id          string `yaml:"id"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

// Guideline is one catalog entry.
type Guideline struct {
	Id              string   `yaml:"id"`
	Title           string   `yaml:"title"`
	Objective       string   `yaml:"objective"`
	Family          string   `yaml:"family"`
	Applicability   []string `yaml:"applicability,omitempty"`
	Recommendations []string `yaml:"recommendations,omitempty"`
}

// Catalog is the merged in-memory catalog used for markdown and optional Layer 1 YAML export.
type Catalog struct {
	Title      string      `yaml:"title"`
	Metadata   Metadata    `yaml:"metadata"`
	Type       string      `yaml:"type"`
	Families   []Family    `yaml:"families"`
	Guidelines []Guideline `yaml:"guidelines"`
}

// catalogYAMLRoot matches metadata.yaml on disk (applicability uses applicability-groups).
type catalogYAMLRoot struct {
	Title    string `yaml:"title"`
	Type     string `yaml:"type"`
	Metadata struct {
		Id                  string             `yaml:"id"`
		Type                string             `yaml:"type"`
		GemaraVersion       string             `yaml:"gemara-version"`
		Version             string             `yaml:"version"`
		Description         string             `yaml:"description"`
		Author              Actor              `yaml:"author"`
		MappingReferences   []MappingReference `yaml:"mapping-references"`
		ApplicabilityGroups []Category         `yaml:"applicability-groups"`
	} `yaml:"metadata"`
}

type groupsYAML struct {
	Groups []Family `yaml:"groups"`
}

// guidelineYAML matches per-file guidelines (field group → Guideline.Family).
type guidelineYAML struct {
	ID              string   `yaml:"id"`
	Title           string   `yaml:"title"`
	Objective       string   `yaml:"objective"`
	State           string   `yaml:"state"`
	Group           string   `yaml:"group"`
	Applicability   []string `yaml:"applicability"`
	Recommendations []string `yaml:"recommendations"`
}

func (g guidelineYAML) toGuideline() Guideline {
	return Guideline{
		Id:              g.ID,
		Title:           g.Title,
		Objective:       g.Objective,
		Family:          g.Group,
		Applicability:   g.Applicability,
		Recommendations: g.Recommendations,
	}
}

// LoadCatalog reads metadata.yaml, groups.yaml, and per-group guideline files into a Catalog.
// groupFileOrder lists per-group file basenames (without .yaml); use DefaultGroupFileOrder for normal builds.
func LoadCatalog(catalogDir string, groupFileOrder []string) (*Catalog, error) {
	metadataPath := filepath.Join(catalogDir, "metadata.yaml")
	rawMeta, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("read metadata.yaml: %w", err)
	}

	var root catalogYAMLRoot
	if err := yaml.Unmarshal(rawMeta, &root); err != nil {
		return nil, fmt.Errorf("parse metadata.yaml: %w", err)
	}

	md := Metadata{
		Id:                root.Metadata.Id,
		Type:              root.Metadata.Type,
		Version:           root.Metadata.Version,
		GemaraVersion:     root.Metadata.GemaraVersion,
		Description:       root.Metadata.Description,
		Author:            root.Metadata.Author,
		MappingReferences: root.Metadata.MappingReferences,
	}
	for _, ag := range root.Metadata.ApplicabilityGroups {
		md.ApplicabilityCategories = append(md.ApplicabilityCategories, Category{
			Id:          ag.Id,
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

	var allGuidelines []Guideline
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
			allGuidelines = append(allGuidelines, g.toGuideline())
		}
	}

	sort.Slice(allGuidelines, func(i, j int) bool {
		return allGuidelines[i].Id < allGuidelines[j].Id
	})

	return &Catalog{
		Title:      root.Title,
		Type:       root.Type,
		Metadata:   md,
		Families:   groupsDoc.Groups,
		Guidelines: allGuidelines,
	}, nil
}

// familyWithGuidelines is one group plus its guidelines for template rendering.
type familyWithGuidelines struct {
	Family
	Guidelines []Guideline
}

// templateData holds the data structure for the markdown template.
type templateData struct {
	GroupsWithGuidelines []familyWithGuidelines
	ApplicabilityTitles  map[string]string
}

// ToMarkdown converts a Catalog to Markdown for the site.
func ToMarkdown(doc *Catalog, catalogDir string) (string, error) {
	byGroup := make(map[string][]Guideline)
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

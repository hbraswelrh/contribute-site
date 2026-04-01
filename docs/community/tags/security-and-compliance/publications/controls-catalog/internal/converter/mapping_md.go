package converter

import (
	_ "embed"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/goccy/go-yaml"
)

//go:embed mapping_markdown.tmpl
var mappingMarkdownTemplates string

type mappingTargetEntry struct {
	EntryID string `yaml:"entry-id"`
}

type mappingAuthorYAML struct {
	Name string `yaml:"name"`
}

// mappingDocumentYAML is the subset of Gemara MappingDocument needed for site rendering.
type mappingDocumentYAML struct {
	Title    string `yaml:"title"`
	Metadata struct {
		ID                string            `yaml:"id"`
		Type              string            `yaml:"type"`
		Version           string            `yaml:"version"`
		GemaraVersion     string            `yaml:"gemara-version"`
		Description       string            `yaml:"description"`
		Author            mappingAuthorYAML `yaml:"author"`
		MappingReferences []struct {
			ID          string `yaml:"id"`
			Title       string `yaml:"title"`
			Version     string `yaml:"version"`
			Description string `yaml:"description"`
			URL         string `yaml:"url"`
		} `yaml:"mapping-references"`
	} `yaml:"metadata"`
	SourceReference struct {
		ReferenceID string `yaml:"reference-id"`
		EntryType   string `yaml:"entry-type"`
	} `yaml:"source-reference"`
	TargetReference struct {
		ReferenceID string `yaml:"reference-id"`
		EntryType   string `yaml:"entry-type"`
	} `yaml:"target-reference"`
	Mappings []struct {
		ID           string               `yaml:"id"`
		Source       string               `yaml:"source"`
		Targets      []mappingTargetEntry `yaml:"targets"`
		Relationship string               `yaml:"relationship"`
		Remarks      string               `yaml:"remarks"`
	} `yaml:"mappings"`
	Remarks string `yaml:"remarks"`
}

type mappingMetaView struct {
	ID                string
	Type              string
	Version           string
	GemaraVersion     string
	Description       string
	Author            mappingAuthorYAML
	MappingReferences []mappingRefTableRow
}

type mappingRefTableRow struct {
	ID, Title, Version string
}

type mappingRefView struct {
	ReferenceID, EntryType string
}

type mappingRowView struct {
	ID, Source, Relationship, Remarks string
	Targets                           []mappingTargetEntry
}

// MergedFlatPart is one atomic YAML mapping row grouped under a single CNSC guideline.
type MergedFlatPart struct {
	Targets   []mappingTargetEntry
	Remarks   string
	MappingID string
}

// MergedFlatRow is one Cross-walked Controls table row (one CNSC guideline, possibly several NIST lines).
type MergedFlatRow struct {
	Source         string
	Relationship   string
	Parts          []MergedFlatPart
	MappingIDs     []string
}

// mappingPageFamily is one catalog family with guidelines and optional NIST rows.
type mappingPageFamily struct {
	Id, Title, Description string
	Guidelines             []mappingPageGuideline
}

type mappingPageGuideline struct {
	Id, Title, Objective string
	Applicability        []string
	Rows                 []mappingRowView
	MergedParts          []MergedFlatPart // same rows as NIST lines for merged table rendering
}

type mappingPageTemplateData struct {
	Title               string
	Remarks             string
	Metadata            mappingMetaView
	SourceReference     mappingRefView
	TargetReference     mappingRefView
	Families            []mappingPageFamily
	FlatMappingsMerged  []MergedFlatRow
	ApplicabilityTitles map[string]string
	// Relative page links (no .md suffix), e.g. ./cnsc-nist-800-53-mapping
	MappingDocumentRel string
	ByFamilyRel        string
}

func targetEntryIDs(targets []mappingTargetEntry) string {
	if len(targets) == 0 {
		return "—"
	}
	var parts []string
	for _, t := range targets {
		parts = append(parts, "`"+t.EntryID+"`")
	}
	return strings.Join(parts, ", ")
}

func firstTargetEntryID(r mappingRowView) string {
	if len(r.Targets) == 0 {
		return "—"
	}
	return r.Targets[0].EntryID
}

// tableCell escapes content for use inside a Markdown pipe table cell.
func tableCell(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "|", "\\|")
	return strings.TrimSpace(s)
}

func anchorID(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}

// mappingRowAnchor returns a stable fragment id for an atomic mapping row (flat index + deep links).
func mappingRowAnchor(mappingID string) string {
	s := strings.ToLower(strings.TrimSpace(mappingID))
	s = strings.ReplaceAll(s, " ", "-")
	return "mapping-row-" + s
}

// buildMappingFamilies groups catalog guidelines by family and attaches NIST rows by CNSC id.
func buildMappingFamilies(cat *Catalog, rowsBySource map[string][]mappingRowView) []mappingPageFamily {
	byGroup := make(map[string][]Guideline)
	for _, g := range cat.Guidelines {
		if g.Family != "" {
			byGroup[g.Family] = append(byGroup[g.Family], g)
		}
	}
	for gid := range byGroup {
		sort.Slice(byGroup[gid], func(i, j int) bool {
			return byGroup[gid][i].Id < byGroup[gid][j].Id
		})
	}

	var out []mappingPageFamily
	for _, grp := range cat.Families {
		gl, ok := byGroup[grp.Id]
		if !ok || len(gl) == 0 {
			continue
		}
		var mgl []mappingPageGuideline
		for _, g := range gl {
			rows := rowsBySource[g.Id]
			if rows == nil {
				rows = []mappingRowView{}
			}
			mgl = append(mgl, mappingPageGuideline{
				Id:            g.Id,
				Title:         g.Title,
				Objective:     g.Objective,
				Applicability: g.Applicability,
				Rows:          rows,
				MergedParts:   rowViewsToMergedParts(rows),
			})
		}
		out = append(out, mappingPageFamily{
			Id:          grp.Id,
			Title:       grp.Title,
			Description: grp.Description,
			Guidelines:  mgl,
		})
	}
	return out
}

func sortFlatMappings(rows []mappingRowView) {
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Source != rows[j].Source {
			return rows[i].Source < rows[j].Source
		}
		a, b := firstTargetEntryID(rows[i]), firstTargetEntryID(rows[j])
		if a != b {
			return a < b
		}
		return rows[i].ID < rows[j].ID
	})
}

func rowViewsToMergedParts(rows []mappingRowView) []MergedFlatPart {
	var p []MergedFlatPart
	for _, r := range rows {
		p = append(p, MergedFlatPart{
			Targets:   r.Targets,
			Remarks:   r.Remarks,
			MappingID: r.ID,
		})
	}
	return p
}

func mergeFlatMappingsBySource(rows []mappingRowView) []MergedFlatRow {
	bySource := make(map[string][]mappingRowView)
	for _, r := range rows {
		bySource[r.Source] = append(bySource[r.Source], r)
	}
	sources := make([]string, 0, len(bySource))
	for s := range bySource {
		sources = append(sources, s)
	}
	sort.Strings(sources)

	var out []MergedFlatRow
	for _, src := range sources {
		group := bySource[src]
		sort.Slice(group, func(i, j int) bool {
			a, b := firstTargetEntryID(group[i]), firstTargetEntryID(group[j])
			if a != b {
				return a < b
			}
			return group[i].ID < group[j].ID
		})

		relSet := make(map[string]struct{})
		for _, r := range group {
			if strings.TrimSpace(r.Relationship) != "" {
				relSet[r.Relationship] = struct{}{}
			}
		}
		rels := make([]string, 0, len(relSet))
		for rs := range relSet {
			rels = append(rels, rs)
		}
		sort.Strings(rels)
		relDisplay := strings.Join(rels, "; ")
		if relDisplay == "" {
			relDisplay = "—"
		}

		var parts []MergedFlatPart
		var ids []string
		for _, r := range group {
			parts = append(parts, MergedFlatPart{
				Targets:   r.Targets,
				Remarks:   r.Remarks,
				MappingID: r.ID,
			})
			ids = append(ids, r.ID)
		}
		out = append(out, MergedFlatRow{
			Source:       src,
			Relationship: relDisplay,
			Parts:        parts,
			MappingIDs:   ids,
		})
	}
	return out
}

// mergedNISTCell renders the NIST control(s) column: one line per YAML mapping when a CNSC has several rows (controls only).
func mergedNISTCell(parts []MergedFlatPart) string {
	if len(parts) == 0 {
		return "—"
	}
	if len(parts) == 1 {
		return targetEntryIDs(parts[0].Targets)
	}
	var b strings.Builder
	for i, p := range parts {
		if i > 0 {
			b.WriteString("<br />")
		}
		b.WriteString(targetEntryIDs(p.Targets))
	}
	// Pipe characters in cell text break markdown pipe tables even inside HTML fragments.
	return strings.ReplaceAll(b.String(), "|", "\\|")
}

// mergedRemarksCell renders the Remarks column; multiple mappings become a bulleted list.
func mergedRemarksCell(parts []MergedFlatPart) string {
	if len(parts) == 0 {
		return "—"
	}
	if len(parts) == 1 {
		return tableCell(parts[0].Remarks)
	}
	var b strings.Builder
	b.WriteString("<ul>")
	for _, p := range parts {
		rem := strings.TrimSpace(strings.ReplaceAll(p.Remarks, "\n", " "))
		if rem == "" {
			rem = "—"
		} else {
			rem = html.EscapeString(rem)
		}
		rem = strings.ReplaceAll(rem, "|", "\\|")
		b.WriteString("<li>")
		b.WriteString(rem)
		b.WriteString("</li>")
	}
	b.WriteString("</ul>")
	return b.String()
}

func mappingTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"targetList":         targetEntryIDs,
		"tableCell":          tableCell,
		"firstTargetEntryID": firstTargetEntryID,
		"anchor":             anchorID,
		"mappingRowAnchor":   mappingRowAnchor,
		"mergedNISTCell":     mergedNISTCell,
		"mergedRemarksCell":  mergedRemarksCell,
		"lower":              strings.ToLower,
		"applicabilityTitle": func(id string, titles map[string]string) string {
			if title, ok := titles[id]; ok {
				return title
			}
			return id
		},
	}
}

func relPageLink(outPath, defaultBasename string) string {
	base := defaultBasename
	if outPath != "" {
		base = filepath.Base(outPath)
	}
	return "./" + strings.TrimSuffix(base, ".md")
}

// buildMappingPageData loads the mapping YAML and catalog into template data (both pages share it).
func buildMappingPageData(catalogDir, mappingYAMLPath string) (mappingPageTemplateData, error) {
	data, err := os.ReadFile(mappingYAMLPath)
	if err != nil {
		return mappingPageTemplateData{}, fmt.Errorf("read mapping yaml: %w", err)
	}
	var mapDoc mappingDocumentYAML
	if err := yaml.Unmarshal(data, &mapDoc); err != nil {
		return mappingPageTemplateData{}, fmt.Errorf("parse mapping yaml: %w", err)
	}

	cat, err := LoadCatalog(catalogDir, DefaultGroupFileOrder)
	if err != nil {
		return mappingPageTemplateData{}, fmt.Errorf("load catalog for mapping page: %w", err)
	}

	rowsBySource := make(map[string][]mappingRowView)
	var flat []mappingRowView
	for _, m := range mapDoc.Mappings {
		r := mappingRowView{
			ID:           m.ID,
			Source:       m.Source,
			Relationship: m.Relationship,
			Remarks:      m.Remarks,
			Targets:      m.Targets,
		}
		rowsBySource[m.Source] = append(rowsBySource[m.Source], r)
		flat = append(flat, r)
	}
	sortFlatMappings(flat)
	flatMerged := mergeFlatMappingsBySource(flat)

	var refs []mappingRefTableRow
	for _, r := range mapDoc.Metadata.MappingReferences {
		refs = append(refs, mappingRefTableRow{ID: r.ID, Title: r.Title, Version: r.Version})
	}

	meta := mappingMetaView{
		ID:                mapDoc.Metadata.ID,
		Type:              mapDoc.Metadata.Type,
		Version:           mapDoc.Metadata.Version,
		GemaraVersion:     mapDoc.Metadata.GemaraVersion,
		Description:       mapDoc.Metadata.Description,
		Author:            mapDoc.Metadata.Author,
		MappingReferences: refs,
	}

	applicabilityTitles := make(map[string]string)
	for _, ac := range cat.Metadata.ApplicabilityCategories {
		applicabilityTitles[ac.Id] = ac.Title
	}
	if catalogDir != "" {
		if err := mergeApplicabilityGroupTitles(filepath.Join(catalogDir, "metadata.yaml"), applicabilityTitles); err != nil {
			return mappingPageTemplateData{}, err
		}
	}

	return mappingPageTemplateData{
		Title:               mapDoc.Title,
		Remarks:             mapDoc.Remarks,
		Metadata:            meta,
		SourceReference:     mappingRefView{ReferenceID: mapDoc.SourceReference.ReferenceID, EntryType: mapDoc.SourceReference.EntryType},
		TargetReference:     mappingRefView{ReferenceID: mapDoc.TargetReference.ReferenceID, EntryType: mapDoc.TargetReference.EntryType},
		Families:            buildMappingFamilies(cat, rowsBySource),
		FlatMappingsMerged:  flatMerged,
		ApplicabilityTitles: applicabilityTitles,
	}, nil
}

func parseMappingMarkdownRoot() (*template.Template, error) {
	return template.New("mappingMarkdown").Funcs(mappingTemplateFuncMap()).Parse(mappingMarkdownTemplates)
}

func renderMappingTemplate(tmpl *template.Template, name string, tv mappingPageTemplateData) (string, error) {
	var b strings.Builder
	if err := tmpl.ExecuteTemplate(&b, name, tv); err != nil {
		return "", fmt.Errorf("execute mapping template %q: %w", name, err)
	}
	return b.String(), nil
}

// DefaultMappingYAMLName is the mapping artifact filename inside the catalog directory.
const DefaultMappingYAMLName = "cnsc-nist-800-53-mapping.yaml"

// DefaultMappingMDName is the Gemara MappingDocument markdown page next to the YAML source.
const DefaultMappingMDName = "cnsc-nist-800-53-mapping.md"

// DefaultMappingByFamilyMDName is the per–guideline-family NIST mapping tables page.
const DefaultMappingByFamilyMDName = "cnsc-nist-800-53-by-family.md"

// WriteMappingMarkdown writes the mapping document and/or the by-family page. Pass "" for a path to skip that file.
// Layout for both pages lives in mapping_markdown.tmpl (templates "mappingDocument" and "mappingByFamily").
func WriteMappingMarkdown(catalogDir, mappingYAMLFile, documentOutPath, byFamilyOutPath string) error {
	yamlPath := mappingYAMLFile
	if !filepath.IsAbs(yamlPath) {
		yamlPath = filepath.Join(catalogDir, mappingYAMLFile)
	}
	tv, err := buildMappingPageData(catalogDir, yamlPath)
	if err != nil {
		return err
	}
	tv.MappingDocumentRel = relPageLink(documentOutPath, DefaultMappingMDName)
	tv.ByFamilyRel = relPageLink(byFamilyOutPath, DefaultMappingByFamilyMDName)

	tmpl, err := parseMappingMarkdownRoot()
	if err != nil {
		return fmt.Errorf("parse mapping_markdown.tmpl: %w", err)
	}

	if documentOutPath != "" {
		md, err := renderMappingTemplate(tmpl, "mappingDocument", tv)
		if err != nil {
			return err
		}
		out := documentOutPath
		if !filepath.IsAbs(out) {
			out = filepath.Join(catalogDir, documentOutPath)
		}
		if err := os.WriteFile(out, []byte(md), 0644); err != nil {
			return err
		}
	}
	if byFamilyOutPath != "" {
		md, err := renderMappingTemplate(tmpl, "mappingByFamily", tv)
		if err != nil {
			return err
		}
		out := byFamilyOutPath
		if !filepath.IsAbs(out) {
			out = filepath.Join(catalogDir, byFamilyOutPath)
		}
		if err := os.WriteFile(out, []byte(md), 0644); err != nil {
			return err
		}
	}
	return nil
}

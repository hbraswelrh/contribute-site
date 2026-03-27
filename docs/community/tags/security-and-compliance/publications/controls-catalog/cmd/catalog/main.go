package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cncf/contribute-site/controls-catalog/internal/converter"
	"github.com/gemaraproj/go-gemara"
	"github.com/goccy/go-yaml"
)

func main() {
	var (
		catalogDir         = flag.String("dir", ".", "Directory containing YAML files (metadata.yaml, groups.yaml, and per-group guideline files)")
		outputYAML         = flag.String("yaml", "", "Output YAML file (Gemara Layer 1 document). If empty, YAML output is skipped.")
		outputMD           = flag.String("md", "index.md", "Output markdown file. Defaults to index.md. If empty, markdown output is skipped.")
		mappingYAML        = flag.String("mapping-yaml", converter.DefaultMappingYAMLName, "Input Gemara MappingDocument YAML (basename or path).")
		mappingMDOut       = flag.String("mapping-md", converter.DefaultMappingMDName, "Output markdown for the Gemara mapping document. If empty, that file is skipped.")
		mappingByFamilyOut = flag.String("mapping-by-family-md", converter.DefaultMappingByFamilyMDName, "Output markdown for NIST mappings grouped by catalog family. If empty, that file is skipped.")
	)
	flag.Parse()

	if *outputYAML == "" && *outputMD == "" && *mappingMDOut == "" && *mappingByFamilyOut == "" {
		fmt.Fprintf(os.Stderr, "Error: at least one of -yaml, -md, -mapping-md, or -mapping-by-family-md must be specified\n")
		flag.Usage()
		os.Exit(1)
	}

	var doc *gemara.GuidanceCatalog
	if *outputYAML != "" || *outputMD != "" {
		var err error
		doc, err = converter.ToGemara(*catalogDir, converter.DefaultGroupFileOrder)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}

	if *outputYAML != "" {
		outputData, err := yaml.Marshal(doc)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to marshal YAML: %v\n", err)
			os.Exit(1)
		}

		if err := os.WriteFile(*outputYAML, outputData, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing YAML output: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Generated %s with %d groups and %d total guidelines\n", *outputYAML, len(doc.Families), len(doc.Guidelines))
		fmt.Printf("   Metadata ID: %s\n", doc.Metadata.Id)
	}

	if *outputMD != "" {
		markdown, err := converter.ToMarkdown(doc, *catalogDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating markdown: %v\n", err)
			os.Exit(1)
		}

		if err := os.WriteFile(*outputMD, []byte(markdown), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing markdown output: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Generated %s with %d groups and %d total guidelines\n", *outputMD, len(doc.Families), len(doc.Guidelines))
	}

	if *mappingMDOut != "" || *mappingByFamilyOut != "" {
		if err := converter.WriteMappingMarkdown(*catalogDir, *mappingYAML, *mappingMDOut, *mappingByFamilyOut); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating mapping markdown: %v\n", err)
			os.Exit(1)
		}
		if *mappingMDOut != "" {
			fmt.Printf("✅ Generated %s from %s\n", *mappingMDOut, *mappingYAML)
		}
		if *mappingByFamilyOut != "" {
			fmt.Printf("✅ Generated %s from %s\n", *mappingByFamilyOut, *mappingYAML)
		}
	}
}

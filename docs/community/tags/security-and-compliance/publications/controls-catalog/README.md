# Cloud Native Security Controls Catalog

This guide covers Gemara format, file organization, and CLI usage for the controls catalog.

This catalog is expressed in **Gemara Layer 1** (Guidance Catalog).
See [Gemara Documentation](https://gemara.openssf.org/) for specification and schema details.

## File breakdown by group

Guidelines are organized into 11 groups (`groups.yaml`). Each group has a dedicated YAML file:

| File                            | Group ID                   | Group title              |
|---------------------------------|----------------------------|--------------------------|
| `access.yaml`                   | `Access`                   | Access Control           |
| `compute.yaml`                  | `Compute`                  | Compute                  |
| `deploy.yaml`                   | `Deploy`                   | Deploy                   |
| `develop.yaml`                  | `Develop`                  | Develop                  |
| `distribute.yaml`               | `Distribute`               | Distribute               |
| `securing-artefacts.yaml`       | `Securing Artefacts`       | Securing Artefacts       |
| `securing-build-pipelines.yaml` | `Securing Build Pipelines` | Securing Build Pipelines |
| `securing-materials.yaml`       | `Securing Materials`       | Securing Materials       |
| `securing-the-source-code.yaml` | `Securing the Source Code` | Securing the Source Code |
| `security-assurance.yaml`       | `Security Assurance`       | Security Assurance       |
| `storage.yaml`                  | `Storage`                  | Storage                  |

Each guideline must have a `group` field matching the group `id` in `groups.yaml` exactly (case-sensitive).

NIST SP 800-53 mappings live in `cnsc-nist-800-53-mapping.yaml` (Gemara **MappingDocument**). Generation produces two markdown pages:

- `cnsc-nist-800-53-mapping.md` — mapping **document** (metadata, references, source/target, **Complete Mapping Index**).
- `cnsc-nist-800-53-by-family.md` — **Cloud Native Security Controls Catalog** view of NIST alignments **grouped by family** (matches `index.md` layout). The mapping document **Complete Mapping Index** is the flat table of all rows (with stable anchors, no separate mapping-id column).

## CLI Usage

The converter tool merges the YAML files into a single Gemara schema-compliant artefact.

### Prerequisites

- Go 1.24 or later

### Validate

Validates YAML files and generates Gemara Layer 1 document:

```bash
go run cmd/catalog/main.go -yaml catalog.yaml
```

**Output**: Creates `catalog.yaml` with a Gemara schema-compliant document. Exits with error if validation fails.

### Generate Markdown

Generates the catalog page and both NIST mapping pages:

```bash
go run cmd/catalog/main.go -md index.md
```

**Output**:

- `index.md` — Cloud Native Security Controls Catalog (GuidanceCatalog).
- `cnsc-nist-800-53-mapping.md` — Gemara mapping document (from `cnsc-nist-800-53-mapping.yaml`).
- `cnsc-nist-800-53-by-family.md` — NIST tables grouped by catalog family (YAML + catalog for layout).

To skip either mapping output, pass `-mapping-md=""` and/or `-mapping-by-family-md=""`.

Regenerate only the mapping pages (leave `index.md` unchanged):

```bash
go run cmd/catalog/main.go -yaml "" -md "" -mapping-md cnsc-nist-800-53-mapping.md -mapping-by-family-md cnsc-nist-800-53-by-family.md
```

### Generate Both YAML and Markdown

```bash
go run cmd/catalog/main.go -yaml catalog.yaml -md index.md
```

### Specify Custom Directory

```bash
go run cmd/catalog/main.go -dir /path/to/catalog -yaml catalog.yaml -md index.md
```

## Control ID Format

- Control ID: `CNSC-{number}` (e.g., `CNSC-1`, `CNSC-200`)
- Statement ID: `CNSC-{number}.{sub}` (e.g., `CNSC-1.1`, `CNSC-200.2`)

IDs must be unique across all group YAML files. Use sequential numbering within each group.

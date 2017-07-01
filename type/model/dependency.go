package model

// Dependency has project dependency and source distination.
type Dependency struct {
	ProjectUUID          string `json:"project_uuid"`
	DependentProjectUUID string `json:"dependent_project_uuid"`
	SourceCD             string `json:"source_cd"`
}

package model

// SourceCD
// 0 = unknown
// 1 = carthage
// 2 = cocoapods
// 3 = submodule

// Dependency has project dependency and source distination.
type Dependency struct {
	ProjectUUID          string `json:"project_uuid" gorm:"ForeignKey:UUID"`
	DependentProjectUUID string `json:"dependent_project_uuid" gorm:"ForeignKey:UUID"`

	SourceCD int `json:"source_cd"`
}

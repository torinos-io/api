package model

import (
	"github.com/guregu/null"
)

// StateCD
// 0 = pending
// 1 = fetching
// 2 = analyzing
// 3 = done
// 4 = error

// Project has uploaded repository information.
type Project struct {
	UUID   string   `json:"uuid" gorm:"primary_key"`
	UserID null.Int `json:"user_id"`

	CartfileContent       string `json:"cartfile_content"`
	PodfileLockContent    string `json:"podfile_content"`
	PbxprojContent        string `json:"pbxproj_content"`
	SupportedSwiftVersion string `json:"supported_swift_version"`
	Repository            string `json:"repository"`
	LastFetchedAt         Time   `json:"last_fetched_at"`
	StateCD               int    `json:"status_cd"`
}

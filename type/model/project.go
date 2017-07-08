package model

import (
	"gopkg.in/guregu/null.v3"
)

// StateCD
// 0 = pending
// 1 = fetching
// 2 = analyzing
// 3 = done
// 4 = error

// Project has uploaded repository information.
type Project struct {
	UUID   string   `json:"uuid"`
	UserID null.Int `json:"user_id"`

	CartfileContent       string `json:"cartfile_content"`
	PodfileLockContent    string `json:"podfile_content"`
	PBXprojectContent     string `json:"pbxporj_content"`
	SupportedSwiftVersion string `json:"supported_swift_version"`
	Repository            string `json:"repository"`
	LastFetchedAt         Time   `json:"last_fetched_at"`
	StateCD               int    `json:"status_cd"`
}

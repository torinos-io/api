package model

// StateCD
// 0 = pending
// 1 = fetching
// 2 = analyzing
// 3 = done
// 4 = error

// Project has uploaded repository information.
type Project struct {
	UUID   string `json:"uuid"`
	UserID int    `json:"user_id"`

	CartFileContent       string `json:"cartfile_content"`
	PodFileLockContent    string `json:"podfile_content"`
	PBXproject            string `json:"pbxporj_content"`
	SupportedSwiftVersion string `json:"supported_swift_version"`
	Repository            string `json:"repository"`
	LastFetchedAt         Time   `json:"last_fetched_at"`
	StateCD               int    `json:"status_cd"`
}

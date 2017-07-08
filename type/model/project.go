package model

// StateCD
// 0 = pending
// 1 = fetching
// 2 = analyzing
// 3 = done
// 4 = error

// Project has uploaded repository information.
type Project struct {
	UUID string `json:"uuid"`

	CartFileContent       string `json:"cart_file_content"`
	PodFileLockContent    string `json:"pods_file_content"`
	PBXproject            string `json:"xcode_xml_content"`
	SupportedSwiftVersion string `json:"supported_swift_version"`
	Repository            string `json:"repository"`
	LastFetchedAt         Time   `json:"last_fetched_at"`
	StateCD               int    `json:"status_cd"`

	UserID int `json:"user_id"`
}

package model

// Project has uploaded repository information.
type Project struct {
	UUID            string `json:"uuid"`
	UserID          int    `json:"user_id"`
	CartFileContent string `json:"cart_file_content"`
	PodsFileContent string `json:"pods_file_content"`
	XcodeXMLContent string `json:"xcode_xml_content"`
}

package service

// Upload uploads files to Analyze service
func (s *service) Upload() error {
	return s.ProjectStore.Upload()
}

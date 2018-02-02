package builds

type Service interface {
	// UploadBuild()
	// UploadBuildChunk()
	// MigrateBuild()
}

type service struct {
}

func NewService() Service {
	return &service{}
}

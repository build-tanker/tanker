package builds

type BuildsService interface {
	// UploadBuild()
	// UploadBuildChunk()
	// MigrateBuild()
}

type buildsService struct {
}

func NewBuildsService() BuildsService {
	return &buildsService{}
}

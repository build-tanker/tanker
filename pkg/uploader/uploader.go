package uploader

import (
	"source.golabs.io/core/tanker/pkg/appcontext"
)

type Uploader interface {
	Upload(accessKey string, bundle string, file string) error
}

type uploader struct {
	ctx *appcontext.AppContext
}

func NewUploader(ctx *appcontext.AppContext) Uploader {
	return &uploader{
		ctx: ctx,
	}
}

func (u *uploader) Upload(accessKey string, bundle string, file string) error {
	log := u.ctx.GetLogger()
	log.Infoln("key:", accessKey, "bundle:", bundle, "file:", file)

	return nil
}

// 	toUpload, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}
// 	defer toUpload.Close()

// 	serverURL := fmt.Sprintf("%s?key=%s&bundle=%s&file=%s", config.UploadServer(), key, bundle, file)
// 	logger.Infof(serverURL)

// 	response, err := http.Post(serverURL, "binary/octet-stream", toUpload)
// 	if err != nil {
// 		return err
// 	}
// 	defer response.Body.Close()

// 	message, _ := ioutil.ReadAll(response.Body)
// 	logger.Infoln(string(message))

// 	return nil
// }

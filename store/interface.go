package store

type Uploader interface {
	Upload(bucketName, objectKey, fileName string) error
}

type Downloader interface {
	DownLoad(file_name string, BucketName string) error
}

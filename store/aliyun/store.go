package aliyun

import (
	"SDK/store"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliOssStore struct {
	client *oss.Client
	//依赖listener 的实习
	listener oss.ProgressListener
}

type Opts struct {
	ossEndpoint     string
	accessKeyId     string
	accessKeySecret string
}

// 接口是否实现
var _ store.Uploader = &AliOssStore{}

func NewDefaultOss() (*AliOssStore, error) {
	opts := &Opts{ossEndpoint: os.Getenv("ALI_OSS_ENDPOINT"), accessKeyId: os.Getenv("ALI_AK"), accessKeySecret: os.Getenv("ALI_SK")}

	client, err := oss.New(opts.ossEndpoint, opts.accessKeyId, opts.accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: client, listener: NewDefaultProgressListener(os.Args[1] == "upload")}, nil

}
func NewAliOssStore(opts *Opts) (*AliOssStore, error) {
	client, err := oss.New(opts.ossEndpoint, opts.accessKeyId, opts.accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: client}, nil
}

func (alioss *AliOssStore) Upload(bucketName, objectKey, fileName string) error {
	buket, err := alioss.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	filename1 := filepath.Base(fileName)

	err = buket.PutObjectFromFile(filename1, fileName, oss.Progress(alioss.listener))

	return err

}

func (alioss *AliOssStore) DownLoad(file_name string, BucketName string) error {

	bucket, err := alioss.client.Bucket(BucketName)
	if err != nil {
		return err
	}
	userProfile := os.Getenv("USERPROFILE")
	desktopPath := filepath.Join(userProfile, "Desktop", file_name)
	err = bucket.GetObjectToFile(file_name, desktopPath, oss.Progress(alioss.listener))
	if err != nil {
		return err
	}
	return nil
}

func (alioss *AliOssStore) List(BucketName string) error {
	bucket, err := alioss.client.Bucket(BucketName)
	if err != nil {
		return err
	}

	lsRes, err := bucket.ListObjects()
	if err != nil {
		return err
	}

	for _, object := range lsRes.Objects {
		fmt.Println("文件名:", object.Key)
	}
	return nil
}

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
	return &AliOssStore{client: client}, nil

}
func NewAliOssStore(opts *Opts) (*AliOssStore, error) {
	client, err := oss.New(opts.ossEndpoint, opts.accessKeyId, opts.accessKeySecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{client: client}, nil
}

func (S *AliOssStore) Upload(bucketName, objectKey, fileName string) error {
	buket, err := S.client.Bucket(bucketName)
	if err != nil {
		return err
	}
	filename := filepath.Base(fileName)
	fmt.Printf("文件%s上传成功", filename)
	return buket.PutObjectFromFile(objectKey, filename)
}

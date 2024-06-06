package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 全局client 在init 初始化
var client *oss.Client
var (
	AccessKeyId     = os.Getenv("ALI_AK")
	AccessKeySecret = os.Getenv("ALI_SK")
	OssEndpoint     = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName      = os.Getenv("ALI_BUCKET_NAME")
)

// 测试aliyum OssSDK
func TestBuketlist(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		HandleError(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("ss")
		fmt.Println("Buckets:", bucket.Name)
	}

}

func HandleError(err error) {
	fmt.Println(err)
}
func TestUploadFile(t *testing.T) {

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		HandleError(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		HandleError(err)
	}
}

// 初始化一个Oss client 给所有测试用例
func init() {
	c, err := oss.New(OssEndpoint, AccessKeyId, AccessKeySecret)

	if err != nil {
		HandleError(err)
	}
	client = c

}

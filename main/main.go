package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 修改变量控制程序运行
var (
	//程序内置
	AccessKeyId     = "LTAI5tKoddUi4B2SNyB3zbz3"
	AccessKeySecret = "KWjCmw9p8XSk6Ey1V70zoINp8LkpCO"
	OssEndpoint     = "oss-cn-beijing.aliyuncs.com"
	//默认配置
	BucketName = "devcloud-suyi"
	//cli
	uploadFile = ""

	help = false
)

// 文件上传
func Upload(File_path string) error {
	client, err := oss.New(OssEndpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return err
	}
	buket, err := client.Bucket(BucketName)
	if err != nil {
		return err
	}
	filename := filepath.Base(File_path)
	fmt.Println(filename)

	return buket.PutObjectFromFile(filename, File_path)
}

func HandleError(err error) {
	fmt.Println(err)
}

// 参数合法检查
func Validate() error {
	if OssEndpoint == "" || AccessKeyId == "" || AccessKeySecret == "" {
		return fmt.Errorf("OssEndpoint, AccessKeyId ,AccessKeySecret exits empty")
	}
	if uploadFile == "" {
		return fmt.Errorf("uploadFile is empty")
	}
	return nil
}
func LoadParams() {
	flag.BoolVar(&help, "h", false, "打印帮助信息")
	flag.StringVar(&uploadFile, "f", "", "上传本地文件路径")
	flag.Parse()
	if help {
		Usage()
	}
}

func Usage() {
	//打印描述信息
	fmt.Fprint(os.Stderr, `shaoyong version:v0.1
Usage: shaoyong -f <upload_file_path>
Options:
-list 输出仓库所有文件名
-upload 上传文件
-download 下载文件到桌面
-h 查看帮助信息
`)
	//打印参数
	flag.PrintDefaults()
}
func Listload() error {
	client, err := oss.New(OssEndpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return err
	}
	bucket, err := client.Bucket(BucketName)
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

func Switch(arg []string) {
	if len(arg) == 1 {
		Usage()
		return
	}
	switch arg[1] {
	case "-h":
		Usage()
	case "-upload":
		err := Upload(arg[2])
		if err == nil {
			fmt.Println("文件上传成功" + arg[2])
		} else {
			fmt.Println("文件上传失败,检查文件是否存在或者路径错误" + arg[2])
		}
	case "-download":
		err := DownLoad(arg[2])
		if err == nil {
			fmt.Println("文件下载成功" + arg[2])
		} else {
			fmt.Println("文件下载失败" + arg[2])
		}
	case "-list":
		Listload()
	default:
		Usage()
	}

}

func DownLoad(name string) error {
	client, err := oss.New(OssEndpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return err
	}
	userProfile := os.Getenv("USERPROFILE")
	desktopPath := filepath.Join(userProfile, "Desktop", name)
	err = bucket.GetObjectToFile(name, desktopPath)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	str := os.Args
	// fmt.Println(str[1])
	// //参数加载
	// LoadParams()
	// if err := Validate(); err != nil {
	// 	fmt.Printf("参数校验异常 %s\n", err)

	// 	os.Exit(1)
	// }

	// if err := Upload(uploadFile); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Printf("文件上传完成 :%s", uploadFile)
	Switch(str)

}

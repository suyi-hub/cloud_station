package cli

import (
	"SDK/store"
	"SDK/store/aliyun"
	"fmt"

	"github.com/spf13/cobra"
)

var uploader store.Uploader
var ossProvider string
var file_name string
var bucket_name string
var UploadCmd = &cobra.Command{
	Use: "upload",
	//详细命令看见的标签
	Long: "upload 文件上传",
	//根命令看见的标签
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch ossProvider {
		case "aliyun":
			aliyun, err := aliyun.NewDefaultOss()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("ssss")
			uploader = aliyun
		case "tx":
		case "aws":
		default:
			return fmt.Errorf("not support oss storage provider")
		}

		uploader.Upload(bucket_name, file_name, file_name)
		fmt.Println(file_name)
		return nil
	},
}

func init() {
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&bucket_name, "bucket_name", "b", "devcloud-suyi", "oss bucket_name ")
	f.StringVarP(&file_name, "file_name", "f", "", "up file_name")
	RootCmd.AddCommand(UploadCmd)

}

package cli

import (
	"SDK/store"
	"SDK/store/aliyun"
	"fmt"

	"github.com/spf13/cobra"
)

var download store.Downloader
var DownloaderCmd = &cobra.Command{
	Use:     "download",
	Long:    "download 文件下载",
	Short:   "download 文件下载",
	Example: "shaoyong cmds",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch ossProvider {
		case "aliyun":
			aliyun, err := aliyun.NewDefaultOss()
			if err != nil {
				fmt.Println(err)
			}
			download = aliyun
		default:
		}
		err := download.DownLoad(file_name, bucket_name)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(DownloaderCmd)
	f := DownloaderCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&bucket_name, "bucket_name", "b", "devcloud-suyi", "oss bucket_name ")
	f.StringVarP(&file_name, "filepath", "f", "", "文件下载路径")
}

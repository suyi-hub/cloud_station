package cli

import (
	"SDK/store"
	"SDK/store/aliyun"
	"fmt"

	"github.com/spf13/cobra"
)

var Lister store.Lister
var List bool
var listCmd = &cobra.Command{
	Use:   "list",
	Long:  "list all filename",
	Short: "list all filename",
	RunE: func(cmd *cobra.Command, args []string) error {
		switch ossProvider {
		case "aliyun":
			aliyun, err := aliyun.NewDefaultOss()
			if err != nil {
				fmt.Println(err)
			}
			Lister = aliyun
		case "tx":
		case "aws":
		default:
		}

		Lister.List(bucket_name)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	f := listCmd.PersistentFlags()
	f.BoolVarP(&List, "list", "l", false, "查看所有文件")
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider [aliyun/tx/aws]")
	f.StringVarP(&bucket_name, "bucket_name", "b", "devcloud-suyi", "oss bucket_name ")
}

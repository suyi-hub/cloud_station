package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version bool

var RootCmd = &cobra.Command{
	Use:     "cloud-station",
	Long:    "cloud-station 云中转站",
	Short:   "cloud-station 阿毅",
	Example: "shaoyong cmds",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("cloud-station v0.0.1")
		}
		return nil
	},
}

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "version", "v", false, "cloud-station 版本信息")
}

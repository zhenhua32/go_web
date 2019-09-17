package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"tzh.com/web/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info of server",
	Long:  "Print the version info of server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("命令参数: ", args)
		printVersion()
	},
}

func printVersion() {
	info := version.Get()
	infoj, err := json.MarshalIndent(&info, "", " ") // 加一点缩进
	if err != nil {
		fmt.Printf("遇到了错误: %v\n", err)
	}
	fmt.Println(string(infoj))
}

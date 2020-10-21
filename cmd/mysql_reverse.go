package cmd

import (
	"fmt"
	"github.com/selftool/config"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

const (
	GoPath      string = "/Users/fuhui/Code/"
	ProjectPath string = "github.com/selftool/"
)

/*
	/bin/reverse
	代码路径需要指定
*/
var mysqlReverseCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {

		result := exec.Command(config.GetGoPath()+"/bin/reverse", "-f",
			config.GetProjectPath()+"/config/mysql.template.yml")
		if _, err := result.Output(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wsxkzs/go-tour-book/internal/sshcmd"
)

var sshCmd = &cobra.Command{
	Use:   "command",
	Short: "远程执行命令",
	Long:  "远程执行命令",
	Run: func(cmd *cobra.Command, args []string) {
		sshcmd.SshExcuteCmd(host, user, password, port, command)
	},
}

var (
	host     string
	user     string
	password string
	port     int
	command  string
)

func init() {
	sshCmd.Flags().StringVarP(&host, "host", "H", "", "请输入主机IP")
	sshCmd.Flags().StringVarP(&user, "user", "u", "", "请输入用户")
	sshCmd.Flags().StringVarP(&password, "password", "p", "", "请输入密码")
	sshCmd.Flags().IntVarP(&port, "port", "P", 22, "请输入端口")
	sshCmd.Flags().StringVarP(&command, "command", "c", "pwd", "请输入单词内容要执行的命令")

}

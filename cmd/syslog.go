package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"log/syslog"
	"os"
)

var (
	addr    = "localhost:514"
	network = "tcp"
	tag     = "info"
	message = "message"
)

func init() {
	rootCmd.AddCommand(SyslogCmd)
	SyslogCmd.Flags().StringVarP(&addr, "addr", "a", "localhost:514", "syslog address")
	SyslogCmd.Flags().StringVarP(&network, "network", "n", "tcp", "network (tcp/udp)")
	SyslogCmd.Flags().StringVarP(&tag, "tag", "t", "info", "syslog tag name")
	SyslogCmd.Flags().StringVarP(&message, "message", "m", "message", "syslog message content")
}

var SyslogCmd = &cobra.Command{
	Use:   "syslog [address] [message]",
	Short: "Send Syslog to Splunk",
	Long: `The syslog command is send syslog message to Splunk syslog service.

e.g. splunk-go syslog -addr localhost:514 -network tcp -tag info -message "this is message"

`,
	Run: func(cmd *cobra.Command, args []string) {
		sysLog, err := syslog.Dial(network, addr,
			syslog.LOG_WARNING|syslog.LOG_DAEMON, tag)
		if err != nil {
			log.Fatal(err)
		}
		err = sysLog.Info(readMessageFile(message))
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Send Success.")
	},
}

func readMessageFile(message string) string {
	_, err := os.Stat(message)
	if err == nil {
		file, err := os.ReadFile(message)
		if err != nil {
			return message
		}
		return string(file)
	}
	return message
}

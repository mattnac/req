package cmd

import (
	"fmt"

	"github.com/mattnac/req/request"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Fire of requests",
	Long: `The send command is used to fire off a batch of requests to the
  specified target/port/uri combination.`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		port, _ := cmd.Flags().GetInt("port")
		count, _ := cmd.Flags().GetInt("count")
		uri, _ := cmd.Flags().GetString("uri")
		req := request.Fire(target, uri, port, count)
		resultString := `================================
    Final test results:
    ===================
    Number of requests sent: %s
    Number of 200 OK responses: %s
    Number of 300 responses: %s
    Number of 400 errors: %s`

		fmt.Println(resultString, req.twoHundreds, req.threeHundreds, req.fourHundreds)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("target", "t", "", "Target URL to hit with requests.")
	sendCmd.Flags().StringP("uri", "u", "/", "The URI to hit, defaults to /")
	sendCmd.Flags().Int("port", 80, "Target port number.")
	sendCmd.Flags().Int("count", 1, "Number of requests to fire off.")
}

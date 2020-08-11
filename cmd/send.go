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
		fmt.Println("Firing of requests, please hold...")
		req := request.Fire(target, uri, port, count)

		fmt.Printf(`
    ================================
    Final test results:
    ================================
    Number of requests sent: %d
    Number of 200 OK responses: %d
    Number of 300 responses: %d
    Number of 400 errors: %d
    ================================`, count, req.TwoHundreds, req.ThreeHundreds, req.FourHundreds)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("target", "t", "", "Target URL to hit with requests.")
	sendCmd.Flags().StringP("uri", "u", "/", "The URI to hit, defaults to /")
	sendCmd.Flags().Int("port", 80, "Target port number.")
	sendCmd.Flags().Int("count", 1, "Number of requests to fire off.")
}

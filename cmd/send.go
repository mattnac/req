package cmd

import (
	"fmt"
	"os"
	"time"

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
		sendReq(cmd)
	},
}

func sendReq(cmd *cobra.Command) {
	start := time.Now()

	target, _ := cmd.Flags().GetString("target")
	port, _ := cmd.Flags().GetInt("port")
	count, _ := cmd.Flags().GetInt("count")
	uri, _ := cmd.Flags().GetString("uri")
	write, _ := cmd.Flags().GetBool("write")

	fmt.Println("Firing off requests, please hold...")
	req := request.Fire(target, uri, port, count)
	elapsed := time.Since(start)
	resultString := fmt.Sprintf(`
================================
Final test results:
================================
Number of requests sent: %d
Number of 200 OK responses: %d
Number of 300 responses: %d
Number of 400 errors: %d
Execution time: %dms
================================`, count, req.TwoHundreds, req.ThreeHundreds, req.FourHundreds, elapsed.Milliseconds())

	if write {
		f, err := os.Create("/tmp/test-report.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString(resultString)
		f.Sync()
		fmt.Println("Results written to", f.Name())
	} else {
		fmt.Printf(resultString)
	}
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("target", "t", "", "Target URL to hit with requests.")
	sendCmd.Flags().StringP("uri", "u", "/", "The URI to hit, defaults to /")
	sendCmd.Flags().Int("port", 80, "Target port number.")
	sendCmd.Flags().Int("count", 1, "Number of requests to fire off.")
	sendCmd.Flags().BoolP("write", "w", false, "Add this flag to write report to a file.")
}

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/mattnac/loadout/request"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "The command that fires of requests.",
	Long: `The send command is used to fire off a batch of requests to the
  specified target/port/uri combination.`,
	Run: func(cmd *cobra.Command, args []string) {
		sendReq(cmd)
	},
}

func sendReq(cmd *cobra.Command) {
	target, _ := cmd.Flags().GetString("target")
	port, _ := cmd.Flags().GetInt("port")
	count, _ := cmd.Flags().GetInt("count")
	uri, _ := cmd.Flags().GetString("uri")
	write, _ := cmd.Flags().GetBool("write")
	insecure, _ := cmd.Flags().GetBool("insecure")
	delay, _ := cmd.Flags().GetInt("delay")

	start := time.Now()
	// perc95 := 0.95

	fmt.Println("Firing off requests, please hold...")
	req, responseTimes := request.Fire(target, uri, port, count, insecure, delay)
	lastIndex := (len(responseTimes) - 1)
	elapsed := time.Since(start)
	meanTime := int(elapsed.Milliseconds()) / count
	fmt.Println(lastIndex)
	fmt.Println(responseTimes[lastIndex])

	sort.Ints(responseTimes)
	fastestResp := responseTimes[0]
	slowestResp := responseTimes[lastIndex]
	// multiplier := math.Round(float64(len(responseTimes)) * perc95)
	// percentile95th := responseTimes[int(multiplier)]
	resultString := fmt.Sprintf(`
================================
Final test results:
================================
Number of requests sent: %d
Number of 200 OK responses: %d
Number of 300 responses: %d
Number of 400 errors: %d
Execution time: %dms
Mean request time: %dms
Fastest Response: %dms
Slowest Response: %dms
95th: ms
================================`, count, req.TwoHundreds, req.ThreeHundreds, req.FourHundreds, elapsed.Milliseconds(), meanTime, fastestResp, slowestResp)

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
	sendCmd.Flags().BoolP("insecure", "i", false, "This flag disregards certificate checks.")
	sendCmd.Flags().Int("delay", 0, "Set a delay in milliseconds between requests")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

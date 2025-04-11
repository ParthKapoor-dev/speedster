/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("No Url Provided")
			return
		}
		url := args[0]

		fmt.Println("Starting Speed Test")

		start := time.Now()

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Unable to Get the file", err)
			return
		}
		defer resp.Body.Close()

		nBytes, err := io.Copy(io.Discard, resp.Body)
		if err != nil {
			fmt.Println("Unable to count Bytes", err)
		}

		duration := time.Since(start).Seconds()
		mbps := float64(nBytes*8) / duration / 1_000_000

		fmt.Printf("Downloaded %.2f MB in %.2f seconds\n", float64(nBytes)/1_000_000, duration)
		fmt.Printf("Download speed: %.2f Mbps\n", mbps)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "speedster",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		url := "https://nbg1-speed.hetzner.com/100MB.bin"

		fmt.Println("Starting Speed Test")

		start := time.Now()

		resp, err := http.Get(url)

		if err != nil {
			fmt.Println("Unable to Get the file", err)
			return
		}
		defer resp.Body.Close()

		fmt.Println(resp)
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.speedster.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

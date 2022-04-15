package cmd

import (
	"fmt"
	"os"
)

var (
	inputFileFlag  string
	outputFileFlag string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

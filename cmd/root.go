package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/th3shadowbroker/kpm2csv/internal/kpm"
)

func init() {
	rootCmd.Flags().StringVarP(&inputFileFlag, "input", "i", "", "Input file")
	rootCmd.MarkFlagFilename("input")
	rootCmd.MarkFlagRequired("input")

	rootCmd.Flags().StringVarP(&outputFileFlag, "output", "o", "", "Output file")
	rootCmd.MarkFlagFilename("output")
	rootCmd.MarkFlagRequired("output")
}

var rootCmd = &cobra.Command{
	Use:   "kpm2csv",
	Short: "A small command line tool to conver exported databases from Kaspersky Password Manager to CSV for importing it into KeepassXC.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Input file: %s\n", inputFileFlag)
		fmt.Printf("Output file: %s\n", outputFileFlag)

		var websites []kpm.Website
		websites, err := kpm.ParseFile(inputFileFlag)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Found %d websites\n", len(websites))

		var outfile, _ = os.OpenFile(outputFileFlag, os.O_CREATE, 0644)
		defer outfile.Close()

		var csv = csv.NewWriter(outfile)
		defer csv.Flush()

		csv.Comma = rune("\t"[0])
		csv.UseCRLF = true

		csv.Write([]string{"Group", "Title", "Username", "Password", "URL", "Notes"})
		for _, website := range websites {
			err := csv.Write(website.AsSlice())
			if err != nil {
				panic(err)
			}
		}

		absPath, _ := filepath.Abs(outputFileFlag)
		fmt.Printf("CSV has been written to %s\n", absPath)
	},
}

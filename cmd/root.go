package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "ccwc is a CLI tool to count characters, words and lines in a given text",
	Long: `The wc utility displays the number of lines, words, and bytes contained in each input file, or standard input (if
		no file is specified) to the standard output.  A line is defined as a string of characters delimited by a ⟨newline⟩
		character.  Characters beyond the final ⟨newline⟩ character will not be included in the line count.
   
		A word is defined as a string of characters delimited by white space characters.  White space characters are the
		set of characters for which the iswspace(3) function returns true.  If more than one input file is specified, a
		line of cumulative counts for all the files is displayed on a separate line after the output for the last file.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

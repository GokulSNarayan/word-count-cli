package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of ccwc",
	Long: `All software has versions. This is ccwc's`,
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("ccwc v0.1")
	},
}


func init(){
	rootCmd.AddCommand(versionCmd)
}
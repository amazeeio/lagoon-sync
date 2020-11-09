package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/bomoko/lagoon-sync/synchers"
	"os"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync a resource type",
	Long:  `Use Lagoon-Sync to sync an external environments resources with the local environment`,
	Run: func(cmd *cobra.Command, args []string) {
		//For now, let's just try write up a command that generates the strings ...
		lagoonConfigBytestream, err := LoadLagoonConfig("./.lagoon.yml")
		if(err != nil) {
			fmt.Println("Couldn't load lagoon config file")
			os.Exit(1)
		}

		fmt.Println(synchers.UnmarshallLagoonYamlToLagoonSyncStructure(lagoonConfigBytestream))

	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

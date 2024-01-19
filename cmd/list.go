package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:    "list",
	Short:  "List resources",
	Long:   `List resources when given a resource type. For example: list clusters`,
	DisableAutoGenTag: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list command")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}


var clustersCmd = &cobra.Command{
	Use:     "clusters",
	Aliases: []string{"cluster"},
	Short:   "List available clusters",
	Long: `List available clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, clusters := range clustersList {
			if !clusters.hidden || viper.Get("DefaultShowHidden") == "true" {
				fmt.Println(clusters.clusterName)
			}
		}
	},
}

func init() {
	listCmd.AddCommand(clustersCmd)
}
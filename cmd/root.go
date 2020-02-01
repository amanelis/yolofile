package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/amanelis/yolofile/api"
)

var (
	// DryRun does not commit or change data
	DryRun bool

	rootCmd = &cobra.Command{
		Use:   "yolo",
		Run: func(cmd *cobra.Command, args []string) {
			api.Execute(os.Args)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(preConfig)

	// // root
	// rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(dsaCmd)
	// rootCmd.AddCommand(infoCmd)
	//
	// // flags
	// rootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "d", false,
	// 	"dry, no commits to real data")
	//
	// // dsa
	// dsaCmd.PersistentFlags().StringVarP(&dsaType, "type", "t", "",
	// 	"type of key: [ecdsa, eddsa, rsa.....]")
	// dsaCmd.MarkFlagRequired("type")
	//
	// dsaCmd.AddCommand(dsaCreateCmd)
	// dsaCmd.AddCommand(dsaGetCmd)
	// dsaCmd.AddCommand(dsaListCmd)
	// dsaCmd.AddCommand(dsaSignCmd)
	// dsaCmd.AddCommand(dsaVerifyCmd)
	// dsaCmd.AddCommand(dsaImportPubCmd)

	// Fire post configuration
	postConfig()
}

func preConfig() {
	if DryRun {
		fmt.Printf("*** --dry-run enabled, no data will be saved ***\n")
	}
}

func postConfig() {
	// ...
}

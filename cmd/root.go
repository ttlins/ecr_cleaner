package cmd

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ecr_cleaner",
		Short: "Cleans up ECR images based on pattern",
		Long: `AWS ECR stands for elastic cloud registry, and it's amazon's solution for storing
		docker images. This is simply a helper script for cleaning tagged images based on a given
		tag pattern`,
	}

	sess *session.Session
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initSession)

	rootCmd.PersistentFlags().StringP("repository", "r", "", "repository name to search for images (required)")
	rootCmd.MarkPersistentFlagRequired("repository")
	viper.BindPFlag("repository", rootCmd.PersistentFlags().Lookup("repository"))

	rootCmd.PersistentFlags().StringP("pattern", "p", "", "tag pattern to use for filtering images (required)")
	rootCmd.MarkPersistentFlagRequired("pattern")
	viper.BindPFlag("pattern", rootCmd.PersistentFlags().Lookup("pattern"))

	rootCmd.PersistentFlags().BoolP("escape", "e", false, "escapes regex metacharacters")
	viper.BindPFlag("escape", rootCmd.PersistentFlags().Lookup("escape"))

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
}

func initSession() {
	sess = session.Must(session.NewSession())
}

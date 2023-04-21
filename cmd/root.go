package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-self-hosted-usage-check",
	Short: "A CLI tool to check the usage of GitHub Actions self-hosted runners",
}

func init() {
	reportCmd.Flags().String("org-name", "", "The name of the GitHub organization")
	reportCmd.Flags().StringArray("runner-labels", []string{}, "The labels that you use for your jobs (can be both user defined and GitHub defined)")
	reportCmd.Flags().Int("num-workflow-runs-to-evaluate", 1, "The number of workflow runs to evaluate for a workflow")
	err := reportCmd.MarkFlagRequired("org-name")
	if err != nil {
		log.Fatal(err)
	}
	err = reportCmd.MarkFlagRequired("runner-labels")
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(reportCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

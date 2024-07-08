package cmd

import (
	"log"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/ssulei7/gh-runner-usage/internal/actions"
	"github.com/ssulei7/gh-runner-usage/internal/report"
	"github.com/ssulei7/gh-runner-usage/internal/repository"
	"github.com/ssulei7/gh-runner-usage/internal/util"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate a report of the usage of GitHub Actions self-hosted runners across an organization",
	Run: func(cmd *cobra.Command, args []string) {
		orgName, err := cmd.Flags().GetString("org-name")
		if err != nil {
			log.Fatalf("Error getting org name: %v", err)
		}
		runnerLabels, err := cmd.Flags().GetStringSlice("runner-labels")
		pterm.DefaultBox.WithTitle(pterm.LightGreen("Runner labels we are searching for")).WithTitleTopCenter().Println(runnerLabels)
		if err != nil {
			log.Fatalf("Error getting self-hosted labels: %v", err)
		}
		numberOfWorkflowRunsToEvaluate, err := cmd.Flags().GetInt("num-workflow-runs-to-evaluate")
		if err != nil {
			log.Fatalf("Error getting number of workflow runs to evaluate: %v", err)
		}

		outputType, err := cmd.Flags().GetString("output-type")
		if err != nil {
			log.Fatalf("Error getting output type: %v", err)
		}

		report.CreateInitialOutputFile(outputType, orgName)

		// Get repositories in the organization
		spinner, _ := pterm.DefaultSpinner.Start("Getting repositories in organization: " + orgName)
		repositories := repository.GetOrgRepositories(orgName)
		spinner.Success("Done!")

		// For each repository, get the workflows and get average of self-hosted minutes
		progressBar, _ := pterm.DefaultProgressbar.WithTotal(len(repositories)).WithTitle("Calculating workflow run averages in: " + orgName).Start()
		defer progressBar.Stop()
		for _, repo := range repositories {
			workflows := repo.GetRepoWorkflows()
			if len(workflows) == 0 {
				progressBar.Increment()
				continue
			}
			for _, workflow := range workflows {
				workflow_runs, err := repo.GetWorkflowRunsForWorkflow(workflow)
				if err != nil {
					pterm.Warning.Println("Error getting workflow runs for workflow: ", workflow)
					continue
				}
				total := 0.0
				if workflow_runs.TotalCount == 0 {
					pterm.Info.Println("No workflow runs found for workflow: ", workflow)
					continue
				}

				// Filter workflow_runs by status of success
				successfulWorkflowRuns := actions.WorkflowRuns{}
				for _, workflow_run := range workflow_runs.WorkflowRuns {
					if workflow_run.Conclusion == "success" {
						successfulWorkflowRuns.TotalCount += 1
						successfulWorkflowRuns.WorkflowRuns = append(successfulWorkflowRuns.WorkflowRuns, workflow_run)
					}
				}

				labels := []string{}

				// Loop based on number of
				for i := 0; i < numberOfWorkflowRunsToEvaluate; i++ {
					if i >= successfulWorkflowRuns.TotalCount {
						break
					}
					jobs := repo.GetWorkflowJobRuns(successfulWorkflowRuns.WorkflowRuns[i].ID)
					if len(jobs.Jobs) == 0 {
						pterm.Info.Println("No jobs found for workflow run: ", successfulWorkflowRuns.WorkflowRuns[i].ID)
						continue
					}
					// Filter jobs by self-hosted labels
					selfHostedJobs := actions.Jobs{}
					for _, job := range jobs.Jobs {
						if job.CheckIfSelfHosted(runnerLabels) {
							selfHostedJobs = append(selfHostedJobs, job)
							labels = append(labels, job.Labels...)
						}
					}
					if len(selfHostedJobs) == 0 {
						pterm.Info.Println("No jobs with specified labels found for workflow run: ", successfulWorkflowRuns.WorkflowRuns[i].ID)
						continue
					}
					// Get the total self-hosted minutes for the workflow run
					for _, job := range selfHostedJobs {
						total += job.CalculateTotalMinutesForJob()
					}
				}

				// Write a new record to the output file
				//remove any duplicate values from labels
				// Remove duplicate values from labels
				uniqueLabels := util.RemoveDuplicates(labels)
				if len(uniqueLabels) > 0 {
					report.AddRecord(outputType, orgName, uniqueLabels, repo.FullName, workflow, total/float64(numberOfWorkflowRunsToEvaluate))
				}
			}

			progressBar.Increment()
		}

		// Print where output file is located
		pterm.Success.Println("Report generated successfully! Report is located at: " + report.GetOutputFilePath(outputType, orgName) + "!")
	},
}

func Cmd() *cobra.Command {
	return reportCmd
}

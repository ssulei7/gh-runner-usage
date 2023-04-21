package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ssulei7/gh-runner-usage-check/internal/actions"
	"github.com/ssulei7/gh-runner-usage-check/internal/report"
	"github.com/ssulei7/gh-runner-usage-check/internal/repository"
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
		fmt.Println("Labels we are searching for in workflows: ", runnerLabels)
		if err != nil {
			log.Fatalf("Error getting self-hosted labels: %v", err)
		}
		numberOfWorkflowRunsToEvaluate, err := cmd.Flags().GetInt("num-workflow-runs-to-evaluate")
		if err != nil {
			log.Fatalf("Error getting number of workflow runs to evaluate: %v", err)
		}

		// Create the report CSV file
		report.CreateInitialReportCSVFile(orgName)

		// Get repositories in the organization
		repositories := repository.GetOrgRepositories(orgName)

		// For each repository, get the workflows and get average of self-hosted minutes
		for _, repo := range repositories {
			fmt.Println("Current repository is: " + repo.FullName)
			workflows := repo.GetRepoWorkflows()
			if len(workflows) == 0 {
				continue
			}
			for _, workflow := range workflows {
				workflow_runs, err := repo.GetWorkflowRunsForWorkflow(workflow)
				if err != nil {
					fmt.Println("Error getting workflow runs for workflow: ", workflow)
					fmt.Println("Continuing..")
					continue
				}
				total := 0.0
				if workflow_runs.TotalCount == 0 {
					fmt.Println("No workflow runs found for workflow: ", workflow)
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

				// Loop based on number of
				for i := 0; i < numberOfWorkflowRunsToEvaluate; i++ {
					if i >= successfulWorkflowRuns.TotalCount {
						break
					}
					jobs := repo.GetWorkflowJobRuns(successfulWorkflowRuns.WorkflowRuns[i].ID)
					if len(jobs.Jobs) == 0 {
						fmt.Println("No jobs found for workflow run: ", successfulWorkflowRuns.WorkflowRuns[i].ID)
						continue
					}
					// Filter jobs by self-hosted labels
					selfHostedJobs := actions.Jobs{}
					for _, job := range jobs.Jobs {
						if job.CheckIfSelfHosted(runnerLabels) {
							selfHostedJobs = append(selfHostedJobs, job)
						}
					}
					if len(selfHostedJobs) == 0 {
						fmt.Println("No self-hosted jobs found for workflow run: ", successfulWorkflowRuns.WorkflowRuns[i].ID)
						continue
					}
					// Get the total self-hosted minutes for the workflow run
					for _, job := range selfHostedJobs {
						total += job.CalculateTotalMinutesForJob()
					}
				}

				// Write as a row in a CSV file
				report.AddRecordToCSVFile(orgName, repo.FullName, workflow, total/float64(numberOfWorkflowRunsToEvaluate))
			}
		}
	},
}

func Cmd() *cobra.Command {
	return reportCmd
}

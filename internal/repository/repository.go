package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/cli/go-gh"
	"github.com/ssulei7/gh-runner-usage-check/internal/actions"
)

type Repository struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
}

func (r *Repository) GetRepoWorkflows() []string {
	// Create a new client
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get all workflows from the repository
	response, err := client.Request(http.MethodGet, fmt.Sprintf("repos/%s/actions/workflows", r.FullName), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make a new json decoder
	decoder := json.NewDecoder(response.Body)
	var workflows_response map[string]interface{}
	decoder.Decode(&workflows_response)
	workflows := workflows_response["workflows"].([]interface{})
	// Check if workflows is empty
	if len(workflows) == 0 {
		fmt.Println("No workflows found for repository: " + r.FullName)
		return []string{}
	}

	workflow_paths := []string{}
	// Get workflow paths
	for _, workflow := range workflows {
		workflow := workflow.(map[string]interface{})
		// Remove the .github/workflows/ from the path
		workflow_name := strings.Replace(workflow["path"].(string), ".github/workflows/", "", 1)
		workflow_paths = append(workflow_paths, workflow_name)
	}

	// Log workflow_paths
	fmt.Println("Workflows for repository: ", r.FullName, " are: ", workflow_paths)

	return workflow_paths

}

func (r *Repository) GetWorkflowRunsForWorkflow(workflowID string) (actions.WorkflowRuns, error) {
	// Create a new client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return actions.WorkflowRuns{}, err
	}

	// Get all workflows from the repository
	response, err := client.Request(http.MethodGet, fmt.Sprintf("repos/%s/actions/workflows/%s/runs", r.FullName, workflowID), nil)
	if err != nil {
		return actions.WorkflowRuns{}, err
	}

	// Make a new json decoder
	decoder := json.NewDecoder(response.Body)
	var workflows_response actions.WorkflowRuns
	decoder.Decode(&workflows_response)
	return workflows_response, nil
}

func (r *Repository) GetWorkflowJobRuns(runID int) actions.WorkflowJobs {
	// Create a new client
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get all workflows from the repository
	response, err := client.Request(http.MethodGet, fmt.Sprintf("repos/%s/actions/runs/%d/jobs", r.FullName, runID), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make a new json decoder
	decoder := json.NewDecoder(response.Body)
	workflow_run_jobs_response := actions.WorkflowJobs{}
	decoder.Decode(&workflow_run_jobs_response)
	return workflow_run_jobs_response
}

type Repositories []Repository

// Get all repositories from a given organization
func GetOrgRepositories(orgName string) Repositories {
	// Create a new client
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get all repositories from the organization
	repositories := Repositories{}
	currentPage := 1
	noNextPage := false
	fmt.Println("Getting repositories for organization: ", orgName)
	for !noNextPage {
		response, err := client.Request(http.MethodGet, fmt.Sprintf("orgs/%s/repos?per_page=1000&page=%d", orgName, currentPage), nil)
		if err != nil {
			log.Fatal(err)
		}

		// Make a new json decoder
		decoder := json.NewDecoder(response.Body)
		var currentSetRepos Repositories
		decoder.Decode(&currentSetRepos)

		// Append to repositories
		repositories = append(repositories, currentSetRepos...)

		links, pageLinksExist := response.Header["Link"]
		if pageLinksExist {
			// Check and see if we have a next page
			for _, links := range links {
				if strings.Contains(links, "rel=\"next\"") {
					currentPage++
					break
				} else {
					// No next page... BREAK
					noNextPage = true
					break
				}
			}
		} else {
			// Doesn't exist... break loop
			break
		}
	}

	return repositories
}

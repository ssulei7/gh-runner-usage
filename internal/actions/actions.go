package actions

import (
	"math"
	"time"
)

type WorkflowRuns struct {
	TotalCount   int `json:"total_count"`
	WorkflowRuns []struct {
		ID               int       `json:"id"`
		Name             string    `json:"name"`
		NodeID           string    `json:"node_id"`
		CheckSuiteID     int       `json:"check_suite_id"`
		CheckSuiteNodeID string    `json:"check_suite_node_id"`
		HeadBranch       string    `json:"head_branch"`
		HeadSha          string    `json:"head_sha"`
		RunNumber        int       `json:"run_number"`
		Event            string    `json:"event"`
		Status           string    `json:"status"`
		Conclusion       string    `json:"conclusion"`
		WorkflowID       int       `json:"workflow_id"`
		URL              string    `json:"url"`
		HTMLURL          string    `json:"html_url"`
		PullRequests     []any     `json:"pull_requests"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		Actor            struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"actor"`
		RunAttempt      int       `json:"run_attempt"`
		RunStartedAt    time.Time `json:"run_started_at"`
		TriggeringActor struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"triggering_actor"`
		JobsURL       string `json:"jobs_url"`
		LogsURL       string `json:"logs_url"`
		CheckSuiteURL string `json:"check_suite_url"`
		ArtifactsURL  string `json:"artifacts_url"`
		CancelURL     string `json:"cancel_url"`
		RerunURL      string `json:"rerun_url"`
		WorkflowURL   string `json:"workflow_url"`
		HeadCommit    struct {
			ID        string    `json:"id"`
			TreeID    string    `json:"tree_id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Committer struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"head_commit"`
		Repository struct {
			ID       int    `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			Private          bool   `json:"private"`
			HTMLURL          string `json:"html_url"`
			Description      string `json:"description"`
			Fork             bool   `json:"fork"`
			URL              string `json:"url"`
			ArchiveURL       string `json:"archive_url"`
			AssigneesURL     string `json:"assignees_url"`
			BlobsURL         string `json:"blobs_url"`
			BranchesURL      string `json:"branches_url"`
			CollaboratorsURL string `json:"collaborators_url"`
			CommentsURL      string `json:"comments_url"`
			CommitsURL       string `json:"commits_url"`
			CompareURL       string `json:"compare_url"`
			ContentsURL      string `json:"contents_url"`
			ContributorsURL  string `json:"contributors_url"`
			DeploymentsURL   string `json:"deployments_url"`
			DownloadsURL     string `json:"downloads_url"`
			EventsURL        string `json:"events_url"`
			ForksURL         string `json:"forks_url"`
			GitCommitsURL    string `json:"git_commits_url"`
			GitRefsURL       string `json:"git_refs_url"`
			GitTagsURL       string `json:"git_tags_url"`
			GitURL           string `json:"git_url"`
			IssueCommentURL  string `json:"issue_comment_url"`
			IssueEventsURL   string `json:"issue_events_url"`
			IssuesURL        string `json:"issues_url"`
			KeysURL          string `json:"keys_url"`
			LabelsURL        string `json:"labels_url"`
			LanguagesURL     string `json:"languages_url"`
			MergesURL        string `json:"merges_url"`
			MilestonesURL    string `json:"milestones_url"`
			NotificationsURL string `json:"notifications_url"`
			PullsURL         string `json:"pulls_url"`
			ReleasesURL      string `json:"releases_url"`
			SSHURL           string `json:"ssh_url"`
			StargazersURL    string `json:"stargazers_url"`
			StatusesURL      string `json:"statuses_url"`
			SubscribersURL   string `json:"subscribers_url"`
			SubscriptionURL  string `json:"subscription_url"`
			TagsURL          string `json:"tags_url"`
			TeamsURL         string `json:"teams_url"`
			TreesURL         string `json:"trees_url"`
			HooksURL         string `json:"hooks_url"`
		} `json:"repository"`
		HeadRepository struct {
			ID       int    `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int    `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			HTMLURL          string `json:"html_url"`
			Description      any    `json:"description"`
			Fork             bool   `json:"fork"`
			URL              string `json:"url"`
			ForksURL         string `json:"forks_url"`
			KeysURL          string `json:"keys_url"`
			CollaboratorsURL string `json:"collaborators_url"`
			TeamsURL         string `json:"teams_url"`
			HooksURL         string `json:"hooks_url"`
			IssueEventsURL   string `json:"issue_events_url"`
			EventsURL        string `json:"events_url"`
			AssigneesURL     string `json:"assignees_url"`
			BranchesURL      string `json:"branches_url"`
			TagsURL          string `json:"tags_url"`
			BlobsURL         string `json:"blobs_url"`
			GitTagsURL       string `json:"git_tags_url"`
			GitRefsURL       string `json:"git_refs_url"`
			TreesURL         string `json:"trees_url"`
			StatusesURL      string `json:"statuses_url"`
			LanguagesURL     string `json:"languages_url"`
			StargazersURL    string `json:"stargazers_url"`
			ContributorsURL  string `json:"contributors_url"`
			SubscribersURL   string `json:"subscribers_url"`
			SubscriptionURL  string `json:"subscription_url"`
			CommitsURL       string `json:"commits_url"`
			GitCommitsURL    string `json:"git_commits_url"`
			CommentsURL      string `json:"comments_url"`
			IssueCommentURL  string `json:"issue_comment_url"`
			ContentsURL      string `json:"contents_url"`
			CompareURL       string `json:"compare_url"`
			MergesURL        string `json:"merges_url"`
			ArchiveURL       string `json:"archive_url"`
			DownloadsURL     string `json:"downloads_url"`
			IssuesURL        string `json:"issues_url"`
			PullsURL         string `json:"pulls_url"`
			MilestonesURL    string `json:"milestones_url"`
			NotificationsURL string `json:"notifications_url"`
			LabelsURL        string `json:"labels_url"`
			ReleasesURL      string `json:"releases_url"`
			DeploymentsURL   string `json:"deployments_url"`
		} `json:"head_repository"`
	} `json:"workflow_runs"`
}

type WorkflowJobs struct {
	TotalCount int  `json:"total_count"`
	Jobs       Jobs `json:"jobs"`
}

type Job struct {
	ID           int64     `json:"id"`
	RunID        int       `json:"run_id"`
	WorkflowName string    `json:"workflow_name"`
	HeadBranch   string    `json:"head_branch"`
	RunURL       string    `json:"run_url"`
	RunAttempt   int       `json:"run_attempt"`
	NodeID       string    `json:"node_id"`
	HeadSha      string    `json:"head_sha"`
	URL          string    `json:"url"`
	HTMLURL      string    `json:"html_url"`
	Status       string    `json:"status"`
	Conclusion   string    `json:"conclusion"`
	CreatedAt    time.Time `json:"created_at"`
	StartedAt    time.Time `json:"started_at"`
	CompletedAt  time.Time `json:"completed_at"`
	Name         string    `json:"name"`
	Steps        []struct {
		Name        string `json:"name"`
		Status      string `json:"status"`
		Conclusion  string `json:"conclusion"`
		Number      int    `json:"number"`
		StartedAt   string `json:"started_at"`
		CompletedAt string `json:"completed_at"`
	} `json:"steps"`
	CheckRunURL       string   `json:"check_run_url"`
	Labels            []string `json:"labels"`
	LabelDescriptions map[string]string
	RunnerID          int    `json:"runner_id"`
	RunnerName        string `json:"runner_name"`
	RunnerGroupID     int    `json:"runner_group_id"`
	RunnerGroupName   string `json:"runner_group_name"`
}

type Jobs []Job

// CalculateTotalMinutesForJob calculates the total number of minutes that a job took to complete.
// It returns the total number of minutes as a float64 value, rounded to three decimal places.
func (j *Job) CalculateTotalMinutesForJob() float64 {
	// Check if job actually completed... if not, return 0
	if j.CompletedAt.Before(j.StartedAt) {
		return 0.0
	}
	return math.Round(float64(j.CompletedAt.Sub(j.StartedAt).Minutes())*1000) / 1000
}

// CheckIfSelfHosted checks if a job is running on a self-hosted runner by comparing the labels of the job with the provided labels.
// It takes a slice of strings representing the labels to check against and returns a boolean value indicating whether the job is running on a self-hosted runner or not.
func (wj *Job) CheckIfSelfHosted(labels []string) bool {
	for _, label := range labels {
		for _, labelInJob := range wj.Labels {
			if label == labelInJob {
				return true
			}
		}
	}
	return false
}

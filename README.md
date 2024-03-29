# gh-runner-usage

This is a GitHub CLI extension that calculates the average minutes consumed by a workflow run, given a set of labels. It is intended to be used to check the usage of self-hosted runners, but you can also use it to check the usage of GitHub-hosted runners. 

## Installation
    
```sh-session
$ gh extension install ssulei7/gh-runner-usage
```

## Usage of the extension

To use the extension, you can run the following command:

```zsh
Generate a report of the usage of GitHub Actions self-hosted runners across an organization

Usage:
  gh runner-usage report [flags]

Flags:
  -h, --help                                help for report
      --num-workflow-runs-to-evaluate int   The number of workflow runs to evaluate for a workflow (default 1)
      --org-name string                     The name of the GitHub organization
      --output-type string                  The type of output to generate (csv or json) (default "csv") 
      --runner-labels strings               The labels that you use for your jobs (can be both user defined and GitHub defined) comma separated.
```

## Example outputs

### CSV

```csv
Repository,Workflow,Average Runner Minutes
SullyDevSquad/sample-repo,main.yml,0.03
SullyDevSquad/sample-repo,override-example.yml,0.08
SullyDevSquad/sample-repo,sample-dynamic-input.yml,0.00
SullyDevSquad/github-flow-demo,go.yml,0.18
SullyDevSquad/github-flow-demo,linter.yml,2.57
```
### JSON
```json
[
  {
    "Average Runner Minutes": "1.00",
    "Repository": "SullyDevSquad/sample-repo",
    "Workflow": "dynamic-env-input.yml"
  },
  {
    "Average Runner Minutes": "3.00",
    "Repository": "SullyDevSquad/sample-repo",
    "Workflow": "main.yml"
  },
  {
    "Average Runner Minutes": "8.00",
    "Repository": "SullyDevSquad/sample-repo",
    "Workflow": "override-example.yml"
  },
  {
    "Average Runner Minutes": "10.00",
    "Repository": "SullyDevSquad/sample-repo",
    "Workflow": "sample-dynamic-input.yml"
  }
]
```

## Contributing
To contribute to gh-runner-usage, follow these steps:

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Push your changes to your fork
5. Create a pull request

## License
gh-runner-usage-check is released under the MIT License.





# gh-runner-usage-check

This is a GitHub CLI extension that calculates the average minutes consumed by a workflow run, given a set of labels. It is intended to be used to check the usage of self-hosted runners, but you can also use it to check the usage of GitHub-hosted runners. 

## Installation
    
```sh-session
$ gh extension install ssulei7/gh-runner-usage-check
```

## Usage of the extension

To use the extension, you can run the following command:

```sh
Generate a report of the usage of GitHub Actions self-hosted runners across an organization

Usage:
  gh-self-hosted-usage-check report [flags]

Flags:
  -h, --help                                help for report
      --num-workflow-runs-to-evaluate int   The number of workflow runs to evaluate for a workflow (default 1)
      --org-name string                     The name of the GitHub organization
      --runner-labels strings               The labels that you use for your jobs (can be both user defined and GitHub defined)
```

## Example output

```csv
Repository,Workflow,Average Runner Minutes
SullyDevSquad/sample-repo,main.yml,0.03
SullyDevSquad/sample-repo,override-example.yml,0.08
SullyDevSquad/sample-repo,sample-dynamic-input.yml,0.00
SullyDevSquad/github-flow-demo,go.yml,0.18
SullyDevSquad/github-flow-demo,linter.yml,2.57
```

## Contributing
To contribute to gh-runner-usage-check, follow these steps:

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Push your changes to your fork
5. Create a pull request

## License
gh-runner-usage-check is released under the MIT License.





package report

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type JSONRecordArray []map[string]string

func GetOutputFilePath(outputType string, orgName string) string {
	if outputType == "csv" {
		return orgName + "-runner-minute-average-report.csv"
	} else if outputType == "json" {
		return orgName + "-runner-minute-average-report.json"
	}
	return ""
}

func CreateInitialOutputFile(outputType string, orgName string) {
	if outputType == "csv" {
		CreateInitialReportCSVFile(orgName)
	} else if outputType == "json" {
		CreateInitialJSONFile(orgName)
	}
}

func CreateInitialReportCSVFile(orgName string) {
	// Create a new csv file
	file, err := os.Create(orgName + "-runner-minute-average-report.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a csv writer
	writer := csv.NewWriter(file)

	// Write headers to the csv file
	headers := []string{"Repository", "Workflow", "Labels", "Average Runner Minutes"}
	err = writer.Write(headers)
	if err != nil {
		panic(err)
	}

	// Flush the csv writer to write any buffered data to the file
	writer.Flush()
}

func CreateInitialJSONFile(orgName string) {
	// Create a new json file
	file, err := os.Create(orgName + "-runner-minute-average-report.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Write opening bracket and closing bracket to the json file
	file.WriteString("[\n]")
}

func AddRecord(outputType string, orgName string, labels []string, repoName string, workflowName string, averageSelfHostedMinutes float64) {
	if outputType == "csv" {
		AddRecordToCSVFile(orgName, repoName, labels, workflowName, averageSelfHostedMinutes)
	} else if outputType == "json" {
		AddRecordToJSONFile(orgName, repoName, labels, workflowName, averageSelfHostedMinutes)
	}
}

func AddRecordToCSVFile(orgName string, repoName string, labels []string, workflowName string, averageSelfHostedMinutes float64) {
	// Open the csv file
	file, err := os.OpenFile(orgName+"-runner-minute-average-report.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a csv writer
	writer := csv.NewWriter(file)

	// Write the record to the csv file
	record := []string{repoName, workflowName, strings.Join(labels, " "), strconv.FormatFloat(averageSelfHostedMinutes, 'f', 2, 64)}
	err = writer.Write(record)
	if err != nil {
		panic(err)
	}

	// Flush the csv writer to write any buffered data to the file
	writer.Flush()
}

func AddRecordToJSONFile(orgName string, repoName string, labels []string, workflowName string, averageSelfHostedMinutes float64) {
	// Load JSON file as an array of struct
	var records JSONRecordArray

	// Open the json file
	file, err := os.ReadFile(orgName + "-runner-minute-average-report.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal the json file into the records array using json.Unmarshal
	err = json.Unmarshal(file, &records)
	if err != nil {
		panic(err)
	}

	// Create a new record
	record := map[string]string{
		"Repository":             repoName,
		"Workflow":               workflowName,
		"Labels":                 strings.Join(labels, " "),
		"Average Runner Minutes": strconv.FormatFloat(averageSelfHostedMinutes, 'f', 2, 64),
	}

	// Append the new record to the records array
	records = append(records, record)

	// Marshal the records array into a json string
	jsonString, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the json string to the json file
	err = os.WriteFile(orgName+"-runner-minute-average-report.json", jsonString, 0644)
	if err != nil {
		panic(err)
	}
}

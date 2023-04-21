package report

import (
	"encoding/csv"
	"os"
	"strconv"
)

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
	headers := []string{"Repository", "Workflow", "Average Runner Minutes"}
	err = writer.Write(headers)
	if err != nil {
		panic(err)
	}

	// Flush the csv writer to write any buffered data to the file
	writer.Flush()
}

func AddRecordToCSVFile(orgName string, repoName string, workflowName string, averageSelfHostedMinutes float64) {
	// Open the csv file
	file, err := os.OpenFile(orgName+"-self-hosted-minute-average-report.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a csv writer
	writer := csv.NewWriter(file)

	// Write the record to the csv file
	record := []string{repoName, workflowName, strconv.FormatFloat(averageSelfHostedMinutes, 'f', 2, 64)}
	err = writer.Write(record)
	if err != nil {
		panic(err)
	}

	// Flush the csv writer to write any buffered data to the file
	writer.Flush()
}

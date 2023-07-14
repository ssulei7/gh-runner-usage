package report

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestCreateInitialOutputFile(t *testing.T) {
	// Test case 1: csv output type
	orgName1 := "myorg"
	outputType1 := "csv"
	CreateInitialOutputFile(outputType1, orgName1)
	_, err1 := os.Stat(orgName1 + "-runner-minute-average-report.csv")
	if os.IsNotExist(err1) {
		t.Errorf("Test case 1 failed: csv file was not created")
	}

	// Test case 2: json output type
	orgName2 := "myorg"
	outputType2 := "json"
	CreateInitialOutputFile(outputType2, orgName2)
	_, err2 := os.Stat(orgName2 + "-runner-minute-average-report.json")
	if os.IsNotExist(err2) {
		t.Errorf("Test case 2 failed: json file was not created")
	}
}

func TestAddRecord(t *testing.T) {
	// Test case 1: csv output type
	orgName1 := "myorg"
	outputType1 := "csv"
	repoName1 := "myrepo"
	workflowName1 := "myworkflow"
	averageSelfHostedMinutes1 := 10.0
	AddRecord(outputType1, orgName1, repoName1, workflowName1, averageSelfHostedMinutes1)
	file1, err1 := os.Open(orgName1 + "-runner-minute-average-report.csv")
	if err1 != nil {
		t.Errorf("Test case 1 failed: csv file could not be opened")
	}
	defer file1.Close()
	reader1 := csv.NewReader(file1)
	records1, err1 := reader1.ReadAll()
	if err1 != nil {
		t.Errorf("Test case 1 failed: csv file could not be read")
	}
	expected1 := []string{repoName1, workflowName1, strconv.FormatFloat(averageSelfHostedMinutes1, 'f', 2, 64)}
	if !reflect.DeepEqual(records1[1], expected1) {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, records1[1])
	}

	// Test case 2: json output type
	orgName2 := "myorg"
	outputType2 := "json"
	repoName2 := "myrepo"
	workflowName2 := "myworkflow"
	averageSelfHostedMinutes2 := 20.0
	AddRecord(outputType2, orgName2, repoName2, workflowName2, averageSelfHostedMinutes2)
	file2, err2 := os.ReadFile(orgName2 + "-runner-minute-average-report.json")
	if err2 != nil {
		t.Errorf("Test case 2 failed: json file could not be opened")
	}
	var records2 JSONRecordArray
	err2 = json.Unmarshal(file2, &records2)
	if err2 != nil {
		t.Errorf("Test case 2 failed: json file could not be unmarshalled")
	}
	expected2 := map[string]string{
		"Repository":             repoName2,
		"Workflow":               workflowName2,
		"Average Runner Minutes": strconv.FormatFloat(averageSelfHostedMinutes2, 'f', 2, 64),
	}
	if !reflect.DeepEqual(records2[0], expected2) {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, records2[0])
	}

	// Remove files created by tests
	os.Remove(orgName1 + "-runner-minute-average-report.csv")
	os.Remove(orgName2 + "-runner-minute-average-report.json")

}

package actions

import (
	"testing"
	"time"
)

func TestCalculateTotalMinutesForJob(t *testing.T) {
	// Test case 1: CompletedAt is after StartedAt
	j1 := &Job{
		StartedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		CompletedAt: time.Date(2021, 1, 1, 0, 10, 0, 0, time.UTC),
	}
	expected1 := 10.0
	if actual := j1.CalculateTotalMinutesForJob(); actual != expected1 {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, actual)
	}

	// Test case 2: CompletedAt is before StartedAt
	j2 := &Job{
		StartedAt:   time.Date(2021, 1, 1, 0, 10, 0, 0, time.UTC),
		CompletedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	expected2 := -10.0
	if actual := j2.CalculateTotalMinutesForJob(); actual != expected2 {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, actual)
	}

	// Test case 3: CompletedAt is equal to StartedAt
	j3 := &Job{
		StartedAt:   time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		CompletedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	expected3 := 0.0
	if actual := j3.CalculateTotalMinutesForJob(); actual != expected3 {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, actual)
	}
}

func TestCheckIfSelfHosted(t *testing.T) {
	// Test case 1: labels match
	j1 := &Job{
		Labels: []string{"self-hosted", "linux"},
	}
	labels1 := []string{"self-hosted"}
	expected1 := true
	if actual := j1.CheckIfSelfHosted(labels1); actual != expected1 {
		t.Errorf("Test case 1 failed: expected %v but got %v", expected1, actual)
	}

	// Test case 2: labels don't match
	j2 := &Job{
		Labels: []string{"windows", "macos"},
	}
	labels2 := []string{"self-hosted"}
	expected2 := false
	if actual := j2.CheckIfSelfHosted(labels2); actual != expected2 {
		t.Errorf("Test case 2 failed: expected %v but got %v", expected2, actual)
	}

	// Test case 3: empty labels
	j3 := &Job{
		Labels: []string{},
	}
	labels3 := []string{"self-hosted"}
	expected3 := false
	if actual := j3.CheckIfSelfHosted(labels3); actual != expected3 {
		t.Errorf("Test case 3 failed: expected %v but got %v", expected3, actual)
	}

	// Test case 4: empty job labels
	j4 := &Job{
		Labels: []string{},
	}
	labels4 := []string{}
	expected4 := false
	if actual := j4.CheckIfSelfHosted(labels4); actual != expected4 {
		t.Errorf("Test case 4 failed: expected %v but got %v", expected4, actual)
	}
}

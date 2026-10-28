// api/cmd/api/workflow_jobs_test.go
package main

import (
	"os"
	"regexp"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

const (
	apiWorkflowPath   = "../../../.github/workflows/api.yml"
	pointJobsStepName = "Point the jobs at the new image"
)

// jobLoopPattern reads the job names out of the step's `for job in ...; do` loop.
var jobLoopPattern = regexp.MustCompile(`for job in ([^;]+);\s*do`)

type workflowFile struct {
	Jobs map[string]struct {
		Steps []struct {
			Name string `yaml:"name"`
			Run  string `yaml:"run"`
		} `yaml:"steps"`
	} `yaml:"jobs"`
}

// deployLoopJobs returns the job names the deploy's image-update loop moves.
func deployLoopJobs(t *testing.T) map[string]bool {
	t.Helper()
	raw, err := os.ReadFile(apiWorkflowPath)
	if err != nil {
		t.Fatalf("read %s: %v", apiWorkflowPath, err)
	}
	var wf workflowFile
	if err := yaml.Unmarshal(raw, &wf); err != nil {
		t.Fatalf("parse %s: %v", apiWorkflowPath, err)
	}
	for _, job := range wf.Jobs {
		for _, step := range job.Steps {
			if step.Name != pointJobsStepName {
				continue
			}
			m := jobLoopPattern.FindStringSubmatch(step.Run)
			if m == nil {
				t.Fatalf("step %q has no `for job in ...; do` loop", pointJobsStepName)
			}
			names := map[string]bool{}
			for _, name := range strings.Fields(m[1]) {
				names[name] = true
			}
			return names
		}
	}
	t.Fatalf("%s has no step named %q", apiWorkflowPath, pointJobsStepName)
	return nil
}

// Every row of the jobs table is a Cloud Run job built from the service image: the deploy
// must move each one, or it keeps running the old image (and the old pack/model versions).
func TestDeployMovesEveryJobInTheJobsTable(t *testing.T) {
	loop := deployLoopJobs(t)
	for _, job := range extraJobs {
		if !loop[job.name] {
			t.Errorf("the %q step in api.yml does not update the %s job", pointJobsStepName, job.name)
		}
	}
}

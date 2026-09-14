// Package jobs runs the Cloud Run job that parses a whole-file upload.
// The API and the job are the same image: the service calls the Cloud
// Run Admin API to execute it with `parse-report <id>` as its arguments.
package jobs

import (
	"context"
	"fmt"
	"sync"

	"google.golang.org/api/option"
	run "google.golang.org/api/run/v2"
)

// Runner executes the parse job. The service holds it as an interface
// so a test can watch what would have been run without a Google Cloud
// project behind it.
type Runner interface {
	Run(ctx context.Context, args ...string) error
}

// CloudRun executes a Cloud Run job through the Admin API, as the
// service account the API runs under.
type CloudRun struct {
	svc  *run.Service
	name string
}

// NewCloudRun builds a runner for one job. Credentials come from the
// runtime service account; nothing is dialled here.
func NewCloudRun(ctx context.Context, project, region, job string) (*CloudRun, error) {
	// The regional endpoint is required for job execution: the global
	// one cannot run a job in another region.
	return NewCloudRunWith(ctx, project, region, job,
		option.WithEndpoint("https://"+region+"-run.googleapis.com/"))
}

// NewCloudRunWith is NewCloudRun with the client options spelled out,
// so a test can point the runner at an in-process Cloud Run.
func NewCloudRunWith(ctx context.Context, project, region, job string, opts ...option.ClientOption) (*CloudRun, error) {
	svc, err := run.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("jobs: cloud run client: %w", err)
	}
	return &CloudRun{
		svc:  svc,
		name: fmt.Sprintf("projects/%s/locations/%s/jobs/%s", project, region, job),
	}, nil
}

// Run starts one execution with the given container arguments. It
// returns as soon as the execution is accepted: a parse takes up to a
// minute and a half, far longer than the request that asked for it.
func (c *CloudRun) Run(ctx context.Context, args ...string) error {
	_, err := c.svc.Projects.Locations.Jobs.Run(c.name, &run.GoogleCloudRunV2RunJobRequest{
		Overrides: &run.GoogleCloudRunV2Overrides{
			ContainerOverrides: []*run.GoogleCloudRunV2ContainerOverride{{Args: args}},
		},
	}).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("jobs: run %s: %w", c.name, err)
	}
	return nil
}

// Fake records what it was asked to run.
type Fake struct {
	mu    sync.Mutex
	Calls [][]string
	Err   error
}

// Run records one call, or fails with the configured error.
func (f *Fake) Run(_ context.Context, args ...string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.Calls = append(f.Calls, args)
	return nil
}

// Ran reports the arguments of the calls so far.
func (f *Fake) Ran() [][]string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([][]string{}, f.Calls...)
}

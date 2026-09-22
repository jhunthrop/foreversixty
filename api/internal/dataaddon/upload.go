// api/internal/dataaddon/upload.go
package dataaddon

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"google.golang.org/api/option"
	storage "google.golang.org/api/storage/v1"
)

// Uploader publishes one named object's bytes. *GCS satisfies it; a nil
// Uploader (or an empty bucket name) means this deployment has no bucket
// configured, and job.go's Run logs why and skips the upload rather than
// failing the run.
type Uploader interface {
	Upload(ctx context.Context, bucket, object string, data []byte) error
}

// GCS uploads through the Cloud Storage JSON API, using the runtime
// service account's Application Default Credentials -- the same
// credential model api/internal/jobs.CloudRun already uses for the Cloud
// Run Admin API, which is why this follows that file's own constructor
// shape (NewX / NewXWith) closely.
type GCS struct {
	svc *storage.Service
}

// NewGCS builds an uploader using Application Default Credentials. Nothing
// is dialled here.
func NewGCS(ctx context.Context) (*GCS, error) {
	return NewGCSWith(ctx)
}

// NewGCSWith is NewGCS with the client options spelled out, so a test can
// point the uploader at an in-process fake server.
func NewGCSWith(ctx context.Context, opts ...option.ClientOption) (*GCS, error) {
	svc, err := storage.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dataaddon: storage client: %w", err)
	}
	return &GCS{svc: svc}, nil
}

// Upload writes data to gs://bucket/object, overwriting any existing
// object at that name -- every publish is a full overwrite, matching
// "Do not edit by hand: the addon-data-release workflow overwrites this
// file every night."
func (g *GCS) Upload(ctx context.Context, bucket, object string, data []byte) error {
	_, err := g.svc.Objects.Insert(bucket, &storage.Object{Name: object}).
		Media(bytes.NewReader(data)).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("dataaddon: upload gs://%s/%s: %w", bucket, object, err)
	}
	return nil
}

// FakeUploader records what it was asked to upload, for job.go's tests.
type FakeUploader struct {
	mu    sync.Mutex
	Calls []UploadCall
	Err   error
}

// UploadCall is one recorded FakeUploader.Upload invocation.
type UploadCall struct {
	Bucket, Object string
	Data           []byte
}

func (f *FakeUploader) Upload(_ context.Context, bucket, object string, data []byte) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Err != nil {
		return f.Err
	}
	f.Calls = append(f.Calls, UploadCall{Bucket: bucket, Object: object, Data: append([]byte{}, data...)})
	return nil
}

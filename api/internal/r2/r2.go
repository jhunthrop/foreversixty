// Package r2 is the API's Cloudflare R2 client: a store.Putter for every
// report file the engine publishes, a reader for the parse job, and the
// signed multipart URLs a browser uploads a whole log through.
//
// R2 speaks the S3 API, so this is the AWS SDK pointed at the account's
// S3 endpoint with static credentials. Tests point it at an in-process
// fake; the engine's store.Dir covers everything that only needs a
// Putter.
package r2

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// PartSize is the multipart part size the contract fixes: 64 MiB.
const PartSize = 64 << 20

// URLTTL is how long a signed upload or download URL is good for.
const URLTTL = time.Hour

// AccessTTL is the ten minutes a signed URL handed to a private
// report's reader lasts.
const AccessTTL = 10 * time.Minute

// Config addresses one bucket.
type Config struct {
	AccountID       string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	// Endpoint overrides the S3 endpoint derived from AccountID. The
	// tests set it; production leaves it empty.
	Endpoint string
}

// Endpoint is the S3 endpoint for the configured account.
func (c Config) endpoint() string {
	if c.Endpoint != "" {
		return c.Endpoint
	}
	return "https://" + c.AccountID + ".r2.cloudflarestorage.com"
}

// Client is one bucket, ready to read, write, and sign.
type Client struct {
	api     *s3.Client
	presign *s3.PresignClient
	bucket  string
}

// New builds a client. It fails only on an incomplete configuration:
// nothing is dialled until the first call.
func New(cfg Config) (*Client, error) {
	if cfg.Bucket == "" || cfg.AccessKeyID == "" || cfg.SecretAccessKey == "" {
		return nil, fmt.Errorf("r2: bucket, access key id, and secret are all required")
	}
	if cfg.AccountID == "" && cfg.Endpoint == "" {
		return nil, fmt.Errorf("r2: an account id or an endpoint is required")
	}
	api := s3.NewFromConfig(aws.Config{
		// R2 has one region and calls it auto.
		Region:      "auto",
		Credentials: credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
	}, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(cfg.endpoint())
		// Path style keeps the bucket in the path rather than in a
		// hostname, which is what R2's endpoint expects and what makes
		// an in-process fake addressable in tests.
		o.UsePathStyle = true
		// R2 rejects the SDK's default trailing checksums on some
		// paths, and a presigned PUT carrying one cannot be replayed by
		// a browser at all. Ask for a checksum only where the API
		// requires one.
		o.RequestChecksumCalculation = aws.RequestChecksumCalculationWhenRequired
		o.ResponseChecksumValidation = aws.ResponseChecksumValidationWhenRequired
	})
	return &Client{api: api, presign: s3.NewPresignClient(api), bucket: cfg.Bucket}, nil
}

// Put writes one object with the cache headers the caller asks for. It
// is the store.Putter the engine's publisher writes through.
func (c *Client) Put(ctx context.Context, key string, body []byte, o store.PutOptions) error {
	in := &s3.PutObjectInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key), Body: bytes.NewReader(body),
		ContentLength: aws.Int64(int64(len(body))),
	}
	if o.ContentType != "" {
		in.ContentType = aws.String(o.ContentType)
	}
	if o.CacheControl != "" {
		in.CacheControl = aws.String(o.CacheControl)
	}
	if o.ContentEncoding != "" {
		in.ContentEncoding = aws.String(o.ContentEncoding)
	}
	if _, err := c.api.PutObject(ctx, in); err != nil {
		return fmt.Errorf("r2: put %s: %w", key, err)
	}
	return nil
}

// Get opens an object for reading. The caller closes it.
func (c *Client) Get(ctx context.Context, key string) (io.ReadCloser, error) {
	out, err := c.api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("r2: get %s: %w", key, err)
	}
	return out.Body, nil
}

// Delete removes an object. Used when an upload is abandoned.
func (c *Client) Delete(ctx context.Context, key string) error {
	if _, err := c.api.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key),
	}); err != nil {
		return fmt.Errorf("r2: delete %s: %w", key, err)
	}
	return nil
}

// Part is one signed part of a multipart upload.
type Part struct {
	Number int    `json:"number"`
	URL    string `json:"url"`
}

// PartCount is how many PartSize parts an upload of size bytes needs. An
// empty upload still has one part, because S3 has no zero-part upload.
func PartCount(size int64) int {
	if size <= 0 {
		return 1
	}
	return int((size + PartSize - 1) / PartSize)
}

// StartMultipart opens a multipart upload and returns its id.
func (c *Client) StartMultipart(ctx context.Context, key, contentType string) (string, error) {
	out, err := c.api.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("r2: start multipart %s: %w", key, err)
	}
	return aws.ToString(out.UploadId), nil
}

// PresignParts signs one PUT URL per part, each good for ttl.
func (c *Client) PresignParts(ctx context.Context, key, uploadID string, parts int, ttl time.Duration) ([]Part, error) {
	out := make([]Part, 0, parts)
	for n := 1; n <= parts; n++ {
		req, err := c.presign.PresignUploadPart(ctx, &s3.UploadPartInput{
			Bucket: aws.String(c.bucket), Key: aws.String(key),
			UploadId: aws.String(uploadID), PartNumber: aws.Int32(int32(n)),
		}, s3.WithPresignExpires(ttl))
		if err != nil {
			return nil, fmt.Errorf("r2: presign part %d of %s: %w", n, key, err)
		}
		out = append(out, Part{Number: n, URL: req.URL})
	}
	return out, nil
}

// CompletedPart is one uploaded part's number and ETag, as the browser
// reports them back.
type CompletedPart struct {
	Number int    `json:"number"`
	ETag   string `json:"etag"`
}

// CompleteMultipart assembles the parts into the final object.
func (c *Client) CompleteMultipart(ctx context.Context, key, uploadID string, parts []CompletedPart) error {
	done := make([]types.CompletedPart, 0, len(parts))
	for _, p := range parts {
		done = append(done, types.CompletedPart{
			PartNumber: aws.Int32(int32(p.Number)),
			ETag:       aws.String(normalizeETag(p.ETag)),
		})
	}
	if _, err := c.api.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key), UploadId: aws.String(uploadID),
		MultipartUpload: &types.CompletedMultipartUpload{Parts: done},
	}); err != nil {
		return fmt.Errorf("r2: complete multipart %s: %w", key, err)
	}
	return nil
}

// AbortMultipart throws away an upload whose parts will never arrive.
func (c *Client) AbortMultipart(ctx context.Context, key, uploadID string) error {
	if _, err := c.api.AbortMultipartUpload(ctx, &s3.AbortMultipartUploadInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key), UploadId: aws.String(uploadID),
	}); err != nil {
		return fmt.Errorf("r2: abort multipart %s: %w", key, err)
	}
	return nil
}

// PresignGet signs a download URL, for the files of a report that the
// site Worker may not serve from the bucket itself.
func (c *Client) PresignGet(ctx context.Context, key string, ttl time.Duration) (string, error) {
	req, err := c.presign.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(c.bucket), Key: aws.String(key),
	}, s3.WithPresignExpires(ttl))
	if err != nil {
		return "", fmt.Errorf("r2: presign get %s: %w", key, err)
	}
	return req.URL, nil
}

// normalizeETag quotes an ETag if the client reported it bare. S3
// requires the quotes; browsers hand back whatever the response header
// carried, which is usually quoted but not always once it has been
// through JSON.
func normalizeETag(etag string) string {
	etag = strings.TrimSpace(etag)
	if strings.HasPrefix(etag, `"`) && strings.HasSuffix(etag, `"`) {
		return etag
	}
	return `"` + strings.Trim(etag, `"`) + `"`
}

// UploadKey is where a whole-file upload's bytes live.
func UploadKey(uploadID string) string { return "uploads/" + uploadID + "/raw.txt.zst" }

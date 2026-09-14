package reports

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
)

const (
	// MaxUploadBytes is the ceiling on a whole-file upload. A 500 MB
	// raid night is the design's worked example; four gigabytes is room
	// for a week of them in one file and still a bound.
	MaxUploadBytes = 4 << 30
	// uploadIDChars is the length of an upload id.
	uploadIDChars = 16
	// ParseJobCommand is the argument the image dispatches on.
	ParseJobCommand = "parse-report"
)

// Multipart is the part of the R2 client the upload routes use.
type Multipart interface {
	StartMultipart(ctx context.Context, key, contentType string) (string, error)
	PresignParts(ctx context.Context, key, uploadID string, parts int, ttl time.Duration) ([]r2.Part, error)
	CompleteMultipart(ctx context.Context, key, uploadID string, parts []r2.CompletedPart) error
	AbortMultipart(ctx context.Context, key, uploadID string) error
}

// Uploads serves the whole-file upload routes: signed multipart URLs
// out, a completion call back, and a parse job started for it.
type Uploads struct {
	Store      *Store
	R2         Multipart
	Jobs       jobs.Runner
	APIBaseURL string
	Log        *slog.Logger
}

// MountUploads registers the upload routes.
func MountUploads(mux *http.ServeMux, u *Uploads) {
	mux.HandleFunc("POST /v1/uploads", auth.RequireSession(u.start))
	mux.HandleFunc("POST /v1/uploads/{upload_id}/complete", auth.RequireSession(u.complete))
}

func (u *Uploads) logger() *slog.Logger {
	if u.Log != nil {
		return u.Log
	}
	return slog.Default()
}

func (u *Uploads) fail(w http.ResponseWriter, r *http.Request, op string, err error, message string) {
	u.logger().Error("uploads", "id", httpx.RequestIDFrom(r.Context()), "op", op, "err", err)
	httpx.WriteError(w, r, http.StatusInternalServerError, "internal", message, nil)
}

// abort throws away a multipart upload that a later step has no way to
// finish, so it does not sit open in R2 forever. An abort that itself
// fails is logged rather than swallowed: the upload is orphaned in R2
// either way, and this is the only record of it.
func (u *Uploads) abort(r *http.Request, key, r2UploadID string) {
	if err := u.R2.AbortMultipart(r.Context(), key, r2UploadID); err != nil {
		u.logger().Error("uploads", "id", httpx.RequestIDFrom(r.Context()), "op", "abort",
			"key", key, "err", err)
	}
}

// StartInput is the body of POST /v1/uploads.
type StartInput struct {
	SizeBytes int64  `json:"size_bytes"`
	Filename  string `json:"filename"`
}

func (u *Uploads) start(w http.ResponseWriter, r *http.Request) {
	var in StartInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with size_bytes and filename", nil)
		return
	}
	if in.SizeBytes <= 0 || in.SizeBytes > MaxUploadBytes {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that file is too large to upload",
			map[string]string{"size_bytes": "between 1 byte and 4 GiB"})
		return
	}
	id := auth.Base32ID(uploadIDChars)
	key := r2.UploadKey(id)
	r2UploadID, err := u.R2.StartMultipart(r.Context(), key, "application/zstd")
	if err != nil {
		u.fail(w, r, "start", err, "could not start that upload just now")
		return
	}
	parts, err := u.R2.PresignParts(r.Context(), key, r2UploadID, r2.PartCount(in.SizeBytes), r2.URLTTL)
	if err != nil {
		// The multipart upload is open in R2 but nothing here will ever
		// finish it: abort it rather than leave it orphaned.
		u.abort(r, key, r2UploadID)
		u.fail(w, r, "start", err, "could not start that upload just now")
		return
	}
	owner := auth.ActorFrom(r.Context()).UserID
	if err := u.Store.CreateUpload(r.Context(), Upload{
		ID: id, UserID: &owner, ObjectKey: key, R2UploadID: r2UploadID,
		SizeBytes: in.SizeBytes, Filename: trimFilename(in.Filename),
	}); err != nil {
		u.abort(r, key, r2UploadID)
		u.fail(w, r, "start", err, "could not start that upload just now")
		return
	}
	httpx.WriteOK(w, r, http.StatusCreated, map[string]any{
		"upload_id":    id,
		"parts":        parts,
		"complete_url": u.APIBaseURL + "/v1/uploads/" + id + "/complete",
	})
}

// CompleteUploadInput is the body of the completion call.
type CompleteUploadInput struct {
	ETags      []r2.CompletedPart `json:"etags"`
	Title      string             `json:"title"`
	Visibility string             `json:"visibility"`
}

func (u *Uploads) complete(w http.ResponseWriter, r *http.Request) {
	var in CompleteUploadInput
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxJSONBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"body must be JSON with etags, and optionally title and visibility", nil)
		return
	}
	if len(in.ETags) == 0 {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "no parts were uploaded",
			map[string]string{"etags": "one entry per uploaded part"})
		return
	}
	if in.Visibility == "" {
		in.Visibility = Public
	}
	if !ValidVisibility(in.Visibility) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "that is not a visibility",
			map[string]string{"visibility": "one of public, unlisted, private, guild"})
		return
	}
	up, err := u.Store.Upload(r.Context(), r.PathValue("upload_id"))
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such upload", nil)
		return
	}
	if err != nil {
		u.fail(w, r, "complete", err, "could not finish that upload just now")
		return
	}
	actor := auth.ActorFrom(r.Context())
	if up.UserID == nil || *up.UserID != actor.UserID {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden", "that upload is not yours", nil)
		return
	}
	// The report id is reserved on the upload row the moment R2 confirms
	// the object is whole, before the report row itself exists. That
	// makes every step from here on replayable: a retry that finds
	// up.ReportID already set never calls CompleteMultipart again on an
	// upload R2 has already finalized (which would error forever), and
	// never creates a second report row for the same upload. Only the
	// request that actually creates the report row starts the parse job:
	// a retry that finds the report already there hands it back and
	// starts nothing, so a client's at-least-once retry of a completion
	// that already succeeded does not launch a second parse.
	reportID := ""
	if up.ReportID != nil {
		reportID = *up.ReportID
	} else {
		if err := u.R2.CompleteMultipart(r.Context(), up.ObjectKey, up.R2UploadID, in.ETags); err != nil {
			u.fail(w, r, "complete", err, "could not finish that upload just now")
			return
		}
		reportID = auth.NewReportID()
		if err := u.Store.FinishUpload(r.Context(), up.ID, reportID); err != nil {
			u.fail(w, r, "complete", err, "could not finish that upload just now")
			return
		}
	}
	rep, err := u.Store.Get(r.Context(), reportID)
	created := false
	if errors.Is(err, ErrNotFound) {
		uploadID := up.ID
		rep, err = u.Store.Create(r.Context(), Report{
			ID: reportID, OwnerID: &actor.UserID, Title: strings.TrimSpace(in.Title),
			Visibility: in.Visibility, Status: StatusProcessing, UploadID: &uploadID,
		})
		created = true
	}
	if err != nil {
		u.fail(w, r, "complete", err, "could not finish that upload just now")
		return
	}
	if !created {
		// The report already exists: either a fully-successful earlier
		// completion, or an earlier attempt that created it and is still
		// running (or already failed) its own parse. Either way this
		// retry starts no second job.
		httpx.WriteOK(w, r, http.StatusAccepted, map[string]string{"report_id": rep.ID})
		return
	}
	if err := u.Jobs.Run(r.Context(), ParseJobCommand, rep.ID); err != nil {
		// The bytes are in the bucket and the report exists, so the
		// parse can be retried; the report says it failed rather than
		// hanging on "processing" forever.
		u.logger().Error("uploads", "id", httpx.RequestIDFrom(r.Context()), "op", "job",
			"report", rep.ID, "err", err)
		if err := u.Store.SetStatus(r.Context(), rep.ID, StatusFailed, "", nil); err != nil {
			u.fail(w, r, "complete", err, "could not start parsing that upload")
			return
		}
		httpx.WriteError(w, r, http.StatusBadGateway, "upstream",
			"the upload was stored but parsing could not be started", nil)
		return
	}
	httpx.WriteOK(w, r, http.StatusAccepted, map[string]string{"report_id": rep.ID})
}

// trimFilename bounds a filename and keeps its base name only.
func trimFilename(name string) string {
	name = strings.TrimSpace(name)
	if i := strings.LastIndexAny(name, `/\`); i >= 0 {
		name = name[i+1:]
	}
	if len(name) > 120 {
		name = name[:120]
	}
	return name
}

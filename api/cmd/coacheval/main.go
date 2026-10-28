// Command coacheval is the Log Coach pre-launch evaluation (spec §21): it builds packs for
// stored fights, runs each configured model over the same packs through the generator's
// prompt, parser, validator and one retry, and writes a markdown report and a JSONL of raw
// results. It reads the database and bucket only; it writes no coach_reviews row, records
// no spend and counts nothing against any cap.
//
//	coacheval -corpus ./summaries -n 20 -model deepseek/deepseek-v4.1-flash -out ./eval
//	DATABASE_URL=... coacheval -db "$DATABASE_URL" -model a -model b -out ./eval
package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/r2"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := run(ctx, os.Args[1:], os.Getenv, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, "coacheval:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, getenv func(string) string, out, errOut io.Writer) error {
	o, err := ParseOptions(args, getenv, errOut)
	if err != nil {
		return err
	}
	ev, closeFn, err := wire(ctx, o, getenv, slog.New(slog.NewTextHandler(errOut, nil)))
	if err != nil {
		return err
	}
	defer closeFn()
	targets, results, err := ev.Evaluate(ctx)
	if err != nil {
		return err
	}
	meta := RunMeta{Source: sourceName(o), Fake: o.Fake, Planned: len(targets)}
	if err := WriteOutputs(o.Out, targets, results, meta); err != nil {
		return err
	}
	fmt.Fprintf(out, "wrote %s: %d fights, %d models\n", o.Out, len(targets), len(o.Models))
	return nil
}

func sourceName(o Options) string {
	if o.DB != "" {
		return "database (read only) and bucket"
	}
	return "corpus directory"
}

// wire builds the Evaluation from the options. closeFn releases the database pool.
func wire(ctx context.Context, o Options, getenv func(string) string, log *slog.Logger) (Evaluation, func(), error) {
	ev := Evaluation{N: o.N, Now: time.Now, Log: log, CoachDir: o.CoachDir,
		Tables: &coach.FileTables{BuildsDir: o.BuildsDir, CoachDir: o.CoachDir}}
	for _, slug := range o.Models {
		ev.Models = append(ev.Models, ModelSpec{Slug: slug, Temperature: o.Temperature})
	}
	if err := wireClients(&ev, o, getenv); err != nil {
		return Evaluation{}, nil, err
	}
	closeFn := func() {}
	if o.Corpus != "" {
		ev.Source = dirSource{dir: o.Corpus, now: time.Now}
		return ev, closeFn, nil
	}
	pool, err := db.Connect(ctx, o.DB)
	if err != nil {
		return Evaluation{}, nil, err
	}
	closeFn = pool.Close
	bucket, err := bucketClient(getenv)
	if err != nil {
		closeFn()
		return Evaluation{}, nil, err
	}
	ev.Source = dbSource{pool: pool, get: bucket}
	ev.Pool = &coach.Pool{DB: pool}
	return ev, closeFn, nil
}

func wireClients(ev *Evaluation, o Options, getenv func(string) string) error {
	if o.Fake {
		ev.ClientFor = fakeFactory
		return nil
	}
	clients, err := gatewayClients(o, getenv)
	if err != nil {
		return err
	}
	ev.ClientFor = sharedFactory(clients)
	return nil
}

func bucketClient(getenv func(string) string) (*r2.Client, error) {
	bucket := getenv("R2_BUCKET")
	if bucket == "" {
		bucket = "foreversixty-logs"
	}
	return r2.New(r2.Config{AccountID: getenv("R2_ACCOUNT_ID"), AccessKeyID: getenv("R2_ACCESS_KEY_ID"),
		SecretAccessKey: getenv("R2_SECRET_ACCESS_KEY"), Bucket: bucket, Endpoint: getenv("R2_ENDPOINT")})
}

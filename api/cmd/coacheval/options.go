package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	defaultFights      = 20
	defaultGatewayURL  = "https://ai-gateway.vercel.sh/v1"
	defaultTemperature = 0.3
	pairParts          = 2
)

// price is USD per million tokens.
type price struct{ In, Out float64 }

// Options is the parsed command line.
type Options struct {
	DB, Corpus, Out     string
	N                   int
	Models              []string
	Prices              map[string]price
	Providers           map[string][]string
	Fake, DrySpend      bool
	CoachDir, BuildsDir string
	Temperature         float64
}

type listFlag []string

func (l *listFlag) String() string     { return strings.Join(*l, ",") }
func (l *listFlag) Set(v string) error { *l = append(*l, v); return nil }

// ParseOptions reads args; getenv supplies the directory defaults.
func ParseOptions(args []string, getenv func(string) string, errOut io.Writer) (Options, error) {
	fs := flag.NewFlagSet("coacheval", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var models, prices, providers listFlag
	o := Options{}
	fs.StringVar(&o.DB, "db", "", "database URL; read-only queries pick stored kills (summaries come from the R2_* bucket)")
	fs.StringVar(&o.Corpus, "corpus", "", "directory of fight summary JSON files, for offline runs")
	fs.StringVar(&o.Out, "out", "", "directory for report.md, results.jsonl and fights.private.json")
	fs.IntVar(&o.N, "n", defaultFights, "fights to evaluate")
	fs.Var(&models, "model", "gateway model slug; repeat to run each model over the same packs")
	fs.Var(&prices, "price", "slug=inUSDperMTok,outUSDperMTok; repeatable; falls back to LLM_PRICE_IN_PER_MTOK/LLM_PRICE_OUT_PER_MTOK")
	fs.Var(&providers, "provider", "slug=provider1,provider2 pins gateway routing; repeatable; falls back to LLM_PROVIDERS")
	fs.BoolVar(&o.Fake, "fake", false, "use llm.Fake instead of the gateway (smoke test)")
	fs.BoolVar(&o.DrySpend, "dry-spend", true, "record no spend and no database row; always on")
	fs.StringVar(&o.CoachDir, "coach-dir", getenv("COACH_DATA_DIR"), "spec rule files and personas (default $COACH_DATA_DIR)")
	fs.StringVar(&o.BuildsDir, "builds-dir", getenv("TREE_DATA_DIR"), "client build data (default $TREE_DATA_DIR)")
	fs.Float64Var(&o.Temperature, "temperature", defaultTemperature, "sampling temperature")
	if err := fs.Parse(args); err != nil {
		return Options{}, err
	}
	o.Models = models
	var err error
	if o.Prices, err = parsePrices(prices); err != nil {
		return Options{}, err
	}
	if o.Providers, err = parseProviders(providers); err != nil {
		return Options{}, err
	}
	return o, o.validate()
}

func (o Options) validate() error {
	switch {
	case !o.DrySpend:
		return errors.New("-dry-spend=false: the harness never records spend or writes rows")
	case o.Out == "":
		return errors.New("-out is required")
	case o.DB == "" && o.Corpus == "":
		return errors.New("one of -db or -corpus is required")
	case o.DB != "" && o.Corpus != "":
		return errors.New("-db and -corpus are exclusive")
	case o.N < 1:
		return errors.New("-n must be at least 1")
	case len(o.Models) == 0:
		return errors.New("at least one -model is required")
	case o.CoachDir == "" || o.BuildsDir == "":
		return errors.New("-coach-dir and -builds-dir (or COACH_DATA_DIR and TREE_DATA_DIR) are required")
	}
	return nil
}

func splitPair(s string) (string, string, error) {
	slug, rest, ok := strings.Cut(s, "=")
	if !ok || slug == "" || rest == "" {
		return "", "", fmt.Errorf("%q is not slug=value", s)
	}
	return slug, rest, nil
}

func parsePrices(in []string) (map[string]price, error) {
	out := map[string]price{}
	for _, s := range in {
		slug, rest, err := splitPair(s)
		if err != nil {
			return nil, fmt.Errorf("-price: %w", err)
		}
		parts := strings.Split(rest, ",")
		if len(parts) != pairParts {
			return nil, fmt.Errorf("-price %q needs in,out", s)
		}
		pin, err1 := strconv.ParseFloat(parts[0], 64)
		pout, err2 := strconv.ParseFloat(parts[1], 64)
		if err1 != nil || err2 != nil || pin <= 0 || pout <= 0 {
			return nil, fmt.Errorf("-price %q needs two positive numbers", s)
		}
		out[slug] = price{In: pin, Out: pout}
	}
	return out, nil
}

func parseProviders(in []string) (map[string][]string, error) {
	out := map[string][]string{}
	for _, s := range in {
		slug, rest, err := splitPair(s)
		if err != nil {
			return nil, fmt.Errorf("-provider: %w", err)
		}
		out[slug] = strings.Split(rest, ",")
	}
	return out, nil
}

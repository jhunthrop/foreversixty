package api

import (
	"strings"
	"testing"
)

func targetErrorReq(target float64, ceiling int) SimRequest {
	req := runReq()
	req.TargetError = target
	req.Iterations = ceiling
	return req
}

func TestTargetErrorValidation(t *testing.T) {
	cases := []struct {
		name string
		req  SimRequest
		want string
	}{
		{"half a percent up to thirty thousand", targetErrorReq(0.005, 30000), ""},
		{"a ceiling that is not a whole number of steps", targetErrorReq(0.005, 30500), "multiple of 1000"},
		{"a ceiling of zero", targetErrorReq(0.005, 0), "multiple of 1000"},
		{"a ceiling past the largest lane's", targetErrorReq(0.005, 200000), "at most 100000"},
		{"a negative target", targetErrorReq(-0.1, 30000), "target_error"},
		{"a target of one whole", targetErrorReq(1.5, 30000), "target_error"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.req.Validate()
			switch {
			case c.want == "" && err != nil:
				t.Fatalf("a legal request was refused: %v", err)
			case c.want != "" && err == nil:
				t.Fatalf("an illegal request was accepted; the error should mention %q", c.want)
			case c.want != "" && !strings.Contains(err.Error(), c.want):
				t.Errorf("error %q does not mention %q", err, c.want)
			}
		})
	}
}

// Validate uses the largest lane's numbers because the envelope carries
// no lane. ValidateLane is what the api handler and the page call, and
// it is where a browser request asking for a server-sized run is
// refused.
func TestValidateLaneNarrowsToOneLane(t *testing.T) {
	req := targetErrorReq(0.005, 100000)
	if err := req.Validate(); err != nil {
		t.Fatalf("the largest lane's ceiling was refused by Validate: %v", err)
	}
	if err := req.ValidateLane(LaneServer); err != nil {
		t.Fatalf("the server lane refused its own ceiling: %v", err)
	}
	err := req.ValidateLane(LaneBrowser)
	if err == nil {
		t.Fatal("the browser lane accepted 100000 iterations")
	}
	if !strings.Contains(err.Error(), "30000") {
		t.Errorf("error %q does not name the browser's ceiling", err)
	}

	big := gear()
	big.Bulk.Cap = Caps[LaneServer]
	if err := big.ValidateLane(LaneServer); err != nil {
		t.Fatalf("the server lane refused its own cap: %v", err)
	}
	if err := big.ValidateLane(LaneBrowser); err == nil || !strings.Contains(err.Error(), "400") {
		t.Errorf("the browser lane did not refuse a server-sized cap: %v", err)
	}
	if err := req.ValidateLane("moon"); err == nil || !strings.Contains(err.Error(), "lane") {
		t.Errorf("an unknown lane was accepted: %v", err)
	}
}

func TestNeedsMoreIterations(t *testing.T) {
	cases := []struct {
		name  string
		req   SimRequest
		res   SimResult
		want  bool
		steps int // NextStepIterations; -1 means do not check
	}{
		{
			name:  "a fixed-count run never steps",
			req:   runReq(),
			res:   SimResult{IterationsRun: 3000, DPS: Estimate{Mean: 1000, Error: 50}},
			want:  false,
			steps: 0,
		},
		{
			name:  "still outside the target",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 2000, DPS: Estimate{Mean: 1000, Error: 20}},
			want:  true,
			steps: StepIterations,
		},
		{
			name:  "inside the target",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 2000, DPS: Estimate{Mean: 1000, Error: 4}},
			want:  false,
			steps: 0,
		},
		{
			name:  "exactly on the target is inside it",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 2000, DPS: Estimate{Mean: 1000, Error: 5}},
			want:  false,
			steps: 0,
		},
		{
			name:  "the ceiling wins",
			req:   targetErrorReq(0.005, 3000),
			res:   SimResult{IterationsRun: 3000, DPS: Estimate{Mean: 1000, Error: 90}},
			want:  false,
			steps: 0,
		},
		{
			name:  "a short last step rather than an overshoot",
			req:   targetErrorReq(0.005, 3000),
			res:   SimResult{IterationsRun: 2500, DPS: Estimate{Mean: 1000, Error: 90}},
			want:  true,
			steps: 500,
		},
		{
			name:  "a failed run does not step",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 1000, Error: "the engine died", DPS: Estimate{Mean: 0}},
			want:  false,
			steps: 0,
		},
		{
			name:  "a stopped run does not step",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 1000, Aborted: true, DPS: Estimate{Mean: 1000, Error: 90}},
			want:  false,
			steps: 0,
		},
		{
			name:  "no dps yet is not precision",
			req:   targetErrorReq(0.005, 30000),
			res:   SimResult{IterationsRun: 1000, DPS: Estimate{Mean: 0, Error: 0}},
			want:  true,
			steps: StepIterations,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := NeedsMoreIterations(c.res, c.req); got != c.want {
				t.Errorf("NeedsMoreIterations = %v, want %v", got, c.want)
			}
			if c.steps >= 0 {
				if got := NextStepIterations(c.res, c.req); got != c.steps {
					t.Errorf("NextStepIterations = %d, want %d", got, c.steps)
				}
			}
		})
	}
}

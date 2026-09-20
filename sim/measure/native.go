package measure

// NativeIterationsPerCPUSecond is how many sim iterations one CPU
// second of the native engine completes, measured by this package's
// own benchmark on the reference fixture.
//
// Two things are derived from it and must not be derived from
// anything else. api.Caps[api.LaneServer] is the largest bulk
// expansion the premium job can finish inside its 15-minute timeout:
// at this rate, four CPUs and 840 usable seconds, a fast run's three
// stages over 5,000 combinations fits and 20,000 does not. The api
// lane's "too_large" answer at submit time quotes the same
// arithmetic, so a request the planner would refuse and a request the
// job could not finish are the same set.
//
// It is a measurement, not a target: re-measure with
// `go test ./sim/measure/... -bench BenchmarkNativeIteration` after
// an engine bump and update BOTH this number and the comment's date.
// Measured 2026-09-19 on the warrior-fury fixture.
//
// That measurement predates Task 8, which was the first code to put
// non-zero armor on the simulated target (contract A8's
// TargetArmorByLevel); every run before it faced a target with zero
// armor, so this figure may be off by roughly the 40% of physical
// damage armor mitigates - almost certainly on the high side, since a
// mitigated hit resolves in less engine work than an unmitigated one.
// Task 25 (the Makefile/CI task) looked for `BenchmarkNativeIteration`
// to re-measure this and found no such benchmark in this package -
// only the correctness tests above - so re-measuring would mean
// building a benchmarking harness from nothing, which is out of scope
// for a Makefile/CI change. Left as-is; building that harness and
// re-measuring against an armored target is follow-up work.
const NativeIterationsPerCPUSecond = 1218

// NativeJobCPUs and NativeJobSeconds are the premium job's shape, as
// the api lane provisions it: four CPUs and the usable share of a
// 15-minute timeout, leaving a minute for startup and the result
// write.
const (
	NativeJobCPUs    = 4
	NativeJobSeconds = 840
)

// NativeIterationBudget is how many iterations the premium job can
// run in total. It is the one number the cap and the submit-time
// estimate are both read off.
func NativeIterationBudget() int {
	return NativeIterationsPerCPUSecond * NativeJobCPUs * NativeJobSeconds
}

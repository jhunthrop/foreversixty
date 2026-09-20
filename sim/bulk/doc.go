// Package bulk decides WHICH sims run.
//
// Top Gear, Droptimizer and talent compare are one question - a base
// character, a set of substitutions, a ranked answer - so they are one
// planner. Expand turns a request's candidates into every valid
// combination; Plan turns those into the first stage's requests; Rank
// scores a finished stage and returns either the next one or the
// finished SimResult.
//
// Nothing here runs a sim. The browser's worker pool runs each stage's
// requests as whole, unsplit runs and forever-sim runs them in a loop;
// both call these three functions and neither does any statistics of
// its own. That is the point: the ladder, the cuts, the deltas and the
// within-error grouping are written once, in Go, and compiled into
// both artifacts, so the two lanes cannot rank the same candidates
// differently.
//
// This package touches no protobuf. Item and enchant facts both come
// from sim/internal/simdb's plain-Go view of the embedded build.
package bulk

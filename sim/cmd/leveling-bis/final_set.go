package main

// The published set DPS (coordinator scope, 2026-11-05): verify.go's
// 300-iteration baseline carried about 2 to 3 DPS of noise, and the
// site's "View ranking sim" link reruns the set at 3,000, landing up to
// 2.3 errors away. Once the band's set is final (swaps, the guard), it is
// measured once more at finalSetIterations on verifySeed, the link's own
// count and the ranker's own seed, and that run is what set_dps and
// set_dps_error publish. Per-candidate checks stay at verifyIterations.

// finalSetIterations is the iteration count of the published set figure.
const finalSetIterations = 3000

// finalSetRequestIterations is what to ask the runner for so the run
// lands at finalSetIterations: a healer's runner multiplies every request
// by healIterationScale (score_heal.go), so a healer asks for the
// quotient, rounded up (3,003 effective at scale 7).
func finalSetRequestIterations(spec specInfo) int {
	if isHealer(spec) {
		return (finalSetIterations + healIterationScale - 1) / healIterationScale
	}
	return finalSetIterations
}

// finalSetEffectiveIterations is the count the published figure really
// ran at, for the run block.
func finalSetEffectiveIterations(spec specInfo) int {
	if isHealer(spec) {
		return finalSetRequestIterations(spec) * healIterationScale
	}
	return finalSetIterations
}

// measureFinalSet is the published set figure: DPS and its standard error.
func measureFinalSet(runner engineRunner, spec specInfo, race string, level int, talents string, picks map[string]slotPick) (dps, stdErr float64, err error) {
	ch := bandCharacter("final-set", race, spec.ClassSlug, spec.Spec, level, talents, buildGear(picks))
	return runner.RunPlainDPSWithError(plainRequest(spec, ch, finalSetRequestIterations(spec), verifySeed))
}

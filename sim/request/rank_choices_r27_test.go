package request

import "testing"

// Round 27 (orders): the Skyborne druid rows below 20. Round 26 judged the
// rows on night elf and tauren and put each Skyborne race with one of them;
// a joint grid on the Skyborne races themselves (band 20, 10,000 iterations,
// seeds 1 and 50001) wants its own ranks at 10 to 18.
func TestRound27SkyborneBalanceRankRows(t *testing.T) {
	const wr1, wr2, wr3, mf1, mf2 = 5176, 5177, 5178, 8921, 8924
	checkRankRows(t, "druid-balance", "druid", []rankRowCase{
		{"14 Windshaper Wrath 1, Moonfire 1", "windshaper-skyborne", 14, []int32{wr1, mf1}, []int32{wr2, mf2}},
		{"14 High Order keeps Wrath 2, Moonfire 2", "high-order-skyborne", 14, []int32{wr2, mf2}, []int32{wr1, mf1}},
		{"18 High Order Wrath 3, Moonfire 1", "high-order-skyborne", 18, []int32{wr3, mf1}, []int32{wr2, mf2}},
		{"18 Windshaper Wrath 3, Moonfire 1", "windshaper-skyborne", 18, []int32{wr3, mf1}, []int32{wr2, mf2}},
		{"18 tauren keeps Wrath 2, Moonfire 2", "tauren", 18, []int32{wr2, mf2}, []int32{wr3, mf1}},
	})
}

func TestRound27SkyborneRestorationRankRows(t *testing.T) {
	const rj1, rj2, rj3, ht1, ht2 = 774, 1058, 1430, 5185, 5186
	checkRankRows(t, "druid-restoration", "druid", []rankRowCase{
		{"10 High Order rank 2 of everything", "high-order-skyborne", 10, []int32{rj2, ht2}, []int32{rj1, ht1}},
		{"10 Windshaper rank 2 of everything", "windshaper-skyborne", 10, []int32{rj2, ht2}, []int32{rj1, ht1}},
		{"12 High Order tank rank 1, party rank 2", "high-order-skyborne", 12, []int32{rj1, rj2}, nil},
		// At 13 the Nature's Swiftness line still casts the top Healing Touch (rank 2).
		{"13 Windshaper rank 1", "windshaper-skyborne", 13, []int32{rj1, ht1}, []int32{rj2}},
		{"14 Windshaper rank 1", "windshaper-skyborne", 14, []int32{rj1, ht1}, []int32{rj2, ht2}},
		{"15 Windshaper Healing Touch 1", "windshaper-skyborne", 15, []int32{rj1, ht1}, []int32{rj2, ht2}},
		{"16 Windshaper rank 1", "windshaper-skyborne", 16, []int32{rj1, ht1}, []int32{rj2, rj3, ht2}},
	})
}

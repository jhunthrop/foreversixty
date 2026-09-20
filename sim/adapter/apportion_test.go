package adapter

import "testing"

func TestApportionDividesEvenly(t *testing.T) {
	got := apportion(300, []float64{1, 1, 1})
	want := []int64{100, 100, 100}
	if !int64SliceEqual(got, want) {
		t.Errorf("apportion(300, [1,1,1]) = %v, want %v", got, want)
	}
}

func TestApportionDistributesTheRemainder(t *testing.T) {
	// 100/3 = 33.33... each; naive independent rounding gives 33+33+33 =
	// 99, one short of 100. The leftover unit goes to a largest
	// remainder - here all three tie, so it goes to the first part.
	got := apportion(100, []float64{1, 1, 1})
	var sum int64
	for _, v := range got {
		sum += v
		if v != 33 && v != 34 {
			t.Errorf("apportion(100, [1,1,1]) part = %d, want 33 or 34", v)
		}
	}
	if sum != 100 {
		t.Errorf("apportion(100, [1,1,1]) sums to %d, want 100", sum)
	}
	if got[0] != 34 {
		t.Errorf("apportion(100, [1,1,1])[0] = %d, want 34 (ties keep part order)", got[0])
	}
}

func TestApportionSinglePart(t *testing.T) {
	got := apportion(12345, []float64{7})
	want := []int64{12345}
	if !int64SliceEqual(got, want) {
		t.Errorf("apportion(12345, [7]) = %v, want %v", got, want)
	}
}

func TestApportionEmptyParts(t *testing.T) {
	got := apportion(500, nil)
	if len(got) != 0 {
		t.Errorf("apportion(500, nil) = %v, want an empty slice - there is no part to carry the total", got)
	}
}

// This is the case D4/the tank review actually hit: five equal shares
// whose independent rounding would land two short of the real total,
// not the one-off a three-part tie produces.
func TestApportionCorrectsMultiUnitDisagreement(t *testing.T) {
	shares := []float64{10.4, 10.4, 10.4, 10.4, 10.4}
	const total = 52 // sum(shares) = 52.0 exactly; naive round(10.4)=10 each sums to 50.

	var naive int64
	for _, s := range shares {
		naive += int64(s + 0.5) // round-half-up, same idea a caller might reach for
	}
	if naive == total {
		t.Fatalf("test setup: naive per-part rounding already sums to %d; pick shares that actually disagree", total)
	}

	got := apportion(total, shares)
	var sum int64
	for _, v := range got {
		sum += v
		if v != 10 && v != 11 {
			t.Errorf("apportion(52, [10.4]*5) part = %d, want 10 or 11", v)
		}
	}
	if sum != total {
		t.Errorf("apportion(52, [10.4]*5) sums to %d, want %d", sum, total)
	}
}

// A share list that sums to zero (or carries no positive share at all)
// has no proportion to follow; the total is still owed to someone, and
// spreading it evenly is what the function falls back to.
func TestApportionOfZeroShares(t *testing.T) {
	got := apportion(10, []float64{0, 0, 0, 0})
	var sum int64
	for _, v := range got {
		sum += v
	}
	if sum != 10 {
		t.Errorf("apportion(10, [0,0,0,0]) sums to %d, want 10", sum)
	}
}

func TestApportionOfZeroTotal(t *testing.T) {
	got := apportion(0, []float64{5, 3, 1})
	want := []int64{0, 0, 0}
	if !int64SliceEqual(got, want) {
		t.Errorf("apportion(0, [5,3,1]) = %v, want %v", got, want)
	}
}

func int64SliceEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

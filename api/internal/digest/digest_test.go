package digest

import (
	"math"
	"math/rand/v2"
	"sort"
	"testing"
)

// exact is the brute-force answer the digest is measured against.
func exact(values []float64, q float64) float64 {
	sorted := append([]float64{}, values...)
	sort.Float64s(sorted)
	if q <= 0 {
		return sorted[0]
	}
	if q >= 1 {
		return sorted[len(sorted)-1]
	}
	i := int(q * float64(len(sorted)))
	if i >= len(sorted) {
		i = len(sorted) - 1
	}
	return sorted[i]
}

func TestQuantilesTrackBruteForceOnATypicalSpread(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	d := New()
	values := make([]float64, 0, 10000)
	for range 10000 {
		// A DPS distribution: a broad middle with a long upper tail.
		v := 400 + r.NormFloat64()*120 + math.Abs(r.NormFloat64())*80
		values = append(values, v)
		d.Add(v)
	}
	if d.Count() != 10000 {
		t.Fatalf("count = %d", d.Count())
	}
	spread := exact(values, 0.99) - exact(values, 0.01)
	for _, q := range []float64{0.01, 0.1, 0.25, 0.5, 0.75, 0.9, 0.95, 0.99} {
		got, want := d.Quantile(q), exact(values, q)
		if math.Abs(got-want) > spread*0.01 {
			t.Errorf("quantile %.2f = %.1f, brute force says %.1f (spread %.1f)", q, got, want, spread)
		}
	}
}

func TestASmallSampleIsExact(t *testing.T) {
	d := New()
	for _, v := range []float64{10, 20, 30, 40, 50} {
		d.Add(v)
	}
	if got := d.Quantile(0); got != 10 {
		t.Errorf("min = %v", got)
	}
	if got := d.Quantile(1); got != 50 {
		t.Errorf("max = %v", got)
	}
	if got := d.Quantile(0.5); math.Abs(got-30) > 0.001 {
		t.Errorf("median = %v, want 30", got)
	}
}

func TestCDFIsThePercentileAParseLandsIn(t *testing.T) {
	d := New()
	for i := range 1000 {
		d.Add(float64(i))
	}
	for _, tc := range []struct{ value, want float64 }{
		{-1, 0}, {500, 0.5}, {1000, 1},
	} {
		if got := d.CDF(tc.value); math.Abs(got-tc.want) > 0.02 {
			t.Errorf("CDF(%v) = %v, want about %v", tc.value, got, tc.want)
		}
	}
}

func TestMergeIsTheSameAsAddingEverything(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	whole, left, right := New(), New(), New()
	for i := range 5000 {
		v := r.Float64() * 1000
		whole.Add(v)
		if i%2 == 0 {
			left.Add(v)
		} else {
			right.Add(v)
		}
	}
	left.Merge(right)
	if left.Count() != whole.Count() {
		t.Fatalf("count = %d, want %d", left.Count(), whole.Count())
	}
	for _, q := range []float64{0.05, 0.5, 0.95} {
		a, b := left.Quantile(q), whole.Quantile(q)
		if math.Abs(a-b) > 10 {
			t.Errorf("merged quantile %.2f = %.1f, whole = %.1f", q, a, b)
		}
	}
}

func TestMergingNothingIsHarmless(t *testing.T) {
	d := New()
	d.Add(1)
	d.Merge(nil)
	d.Merge(New())
	if d.Count() != 1 {
		t.Fatalf("count = %d, want 1", d.Count())
	}
}

func TestEncodingRoundTripsAndIsStable(t *testing.T) {
	d := New()
	for i := range 500 {
		d.Add(float64(i) * 1.5)
	}
	b, err := d.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	again, err := d.MarshalBinary()
	if err != nil || string(again) != string(b) {
		t.Fatal("encoding the same digest twice must give the same bytes")
	}
	back, err := Unmarshal(b)
	if err != nil {
		t.Fatal(err)
	}
	if back.Count() != d.Count() {
		t.Fatalf("count = %d, want %d", back.Count(), d.Count())
	}
	for _, q := range []float64{0.1, 0.5, 0.9} {
		if math.Abs(back.Quantile(q)-d.Quantile(q)) > 0.0001 {
			t.Errorf("quantile %.1f did not survive the round trip", q)
		}
	}
}

func TestUnmarshalOfNothingIsAnEmptyDigest(t *testing.T) {
	d, err := Unmarshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	if d.Count() != 0 || d.Quantile(0.5) != 0 || d.CDF(1) != 0 {
		t.Fatal("an empty digest should answer zero to everything")
	}
}

func TestUnmarshalRejectsRubbish(t *testing.T) {
	for name, b := range map[string][]byte{
		"short":       []byte("FSTD"),
		"wrong magic": append([]byte("XXXX"), make([]byte, 17)...),
		"wrong format": append(append([]byte("FSTD"), 9),
			make([]byte, 16)...),
		"bad length": append(append([]byte("FSTD"), 1),
			[]byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 9}...),
	} {
		if _, err := Unmarshal(b); err == nil {
			t.Errorf("%s should have been rejected", name)
		}
	}
}

func TestNaNAndInfinityAreIgnored(t *testing.T) {
	d := New()
	d.Add(math.NaN())
	d.Add(math.Inf(1))
	if d.Count() != 0 {
		t.Fatalf("count = %d, want 0", d.Count())
	}
}

func TestPlacementCountsWhatAParseBeats(t *testing.T) {
	d := New()
	for _, v := range []float64{1000, 1380.9, 2000} {
		d.Add(v)
	}
	cases := []struct{ value, want float64 }{
		{2000, 1},     // the best beats both others
		{1380.4, 0.5}, // recomputed to the hundredth, still the middle one
		{1000, 0},     // the bottom beats nobody
		{999, 0},      // below everyone
		{3000, 1},     // above everyone
	}
	for _, tc := range cases {
		if got := d.Placement(tc.value); math.Abs(got-tc.want) > 1e-9 {
			t.Errorf("Placement(%v) = %v, want %v", tc.value, got, tc.want)
		}
	}
	one := New()
	one.Add(500)
	if got := one.Placement(500); got != 1 {
		t.Errorf("a bracket of one places its only kill at %v, want 1", got)
	}
}

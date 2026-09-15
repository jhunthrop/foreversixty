// Package digest is the percentile sketch behind "your parse was in the
// 87th percentile". One digest per encounter, difficulty, spec, phase,
// and metric is kept in Postgres as bytes, merged with each new fight,
// and asked for quantiles by the rankings routes.
//
// It is a merging t-digest, written here rather than taken from a
// library: the store needs a stable binary encoding it can keep in a
// bytea column and merge incrementally, which is a hundred lines, and
// the alternative on offer has had no release since 2019.
package digest

import (
	"encoding/binary"
	"fmt"
	"math"
	"sort"
)

// DefaultCompression trades size for accuracy. At 100 a digest is a few
// hundred centroids at most - under 5 KB - and the median is accurate
// to well under a percent, with the tails better still, which is what
// the rankings show.
const DefaultCompression = 100

// bufferSize is how many raw values are held before a compression pass.
const bufferSize = 256

// format is the first byte of an encoded digest, so a later encoding
// can be told from this one.
const format byte = 1

// magic marks an encoded digest.
var magic = [4]byte{'F', 'S', 'T', 'D'}

type centroid struct {
	Mean   float64
	Weight float64
}

// Digest is a t-digest. The zero value is not usable; call New.
type Digest struct {
	compression float64
	merged      []centroid
	buffer      []centroid
	count       int64
}

// New returns an empty digest at the default compression.
func New() *Digest { return &Digest{compression: DefaultCompression} }

// Count is how many values the digest has seen.
func (d *Digest) Count() int64 { return d.count }

// Add folds one value in.
func (d *Digest) Add(x float64) {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return
	}
	d.buffer = append(d.buffer, centroid{Mean: x, Weight: 1})
	d.count++
	if len(d.buffer) >= bufferSize {
		d.compress()
	}
}

// Merge folds another digest in whole.
func (d *Digest) Merge(o *Digest) {
	if o == nil {
		return
	}
	d.buffer = append(d.buffer, o.merged...)
	d.buffer = append(d.buffer, o.buffer...)
	d.count += o.count
	d.compress()
}

// compress folds the buffer into the merged centroids, bounding each
// centroid's weight by its position: centroids near the middle may be
// heavy, centroids at the tails stay light, which is what makes the
// extreme quantiles accurate.
func (d *Digest) compress() {
	if d.compression == 0 {
		d.compression = DefaultCompression
	}
	all := make([]centroid, 0, len(d.merged)+len(d.buffer))
	all = append(all, d.merged...)
	all = append(all, d.buffer...)
	d.buffer = d.buffer[:0]
	if len(all) == 0 {
		d.merged = nil
		return
	}
	sort.Slice(all, func(i, j int) bool { return all[i].Mean < all[j].Mean })
	var total float64
	for _, c := range all {
		total += c.Weight
	}
	out := make([]centroid, 0, len(all))
	out = append(out, all[0])
	var before float64
	for _, c := range all[1:] {
		cur := &out[len(out)-1]
		proposed := cur.Weight + c.Weight
		q := (before + proposed/2) / total
		if proposed <= 4*total*q*(1-q)/d.compression {
			cur.Mean = (cur.Mean*cur.Weight + c.Mean*c.Weight) / proposed
			cur.Weight = proposed
			continue
		}
		before += cur.Weight
		out = append(out, c)
	}
	d.merged = out
}

// centroids returns the compressed view, folding in anything buffered.
func (d *Digest) centroids() []centroid {
	if len(d.buffer) > 0 {
		d.compress()
	}
	return d.merged
}

// total is the summed weight of the centroids.
func total(cs []centroid) float64 {
	var t float64
	for _, c := range cs {
		t += c.Weight
	}
	return t
}

// Quantile is the value at q, with q in 0..1.
func (d *Digest) Quantile(q float64) float64 {
	cs := d.centroids()
	if len(cs) == 0 {
		return 0
	}
	if q <= 0 {
		return cs[0].Mean
	}
	if q >= 1 {
		return cs[len(cs)-1].Mean
	}
	t := total(cs)
	target := q * t
	var before float64
	for i, c := range cs {
		center := before + c.Weight/2
		if target <= center {
			if i == 0 {
				return c.Mean
			}
			prev := cs[i-1]
			prevCenter := before - prev.Weight/2
			span := center - prevCenter
			if span <= 0 {
				return c.Mean
			}
			frac := (target - prevCenter) / span
			return prev.Mean + frac*(c.Mean-prev.Mean)
		}
		before += c.Weight
	}
	return cs[len(cs)-1].Mean
}

// CDF is the fraction of values at or below x, in 0..1. It is the
// percentile a parse lands in, before multiplying by a hundred.
func (d *Digest) CDF(x float64) float64 {
	cs := d.centroids()
	if len(cs) == 0 {
		return 0
	}
	t := total(cs)
	if x < cs[0].Mean {
		return 0
	}
	if x >= cs[len(cs)-1].Mean {
		return 1
	}
	var before float64
	for i, c := range cs {
		center := before + c.Weight/2
		if x < c.Mean {
			if i == 0 {
				return center / t
			}
			prev := cs[i-1]
			prevCenter := before - prev.Weight/2
			span := c.Mean - prev.Mean
			if span <= 0 {
				return center / t
			}
			frac := (x - prev.Mean) / span
			return (prevCenter + frac*(center-prevCenter)) / t
		}
		before += c.Weight
	}
	return 1
}

// Placement is the share of the other values a parse beats, in 0..1: the
// top of a bracket is 1, the bottom 0, and a bracket of one has nothing to
// beat and reads 1. It counts centroids strictly below x, with a small
// tolerance so a value the page recomputed to the hundredth still matches
// the one the ingest stored, and never interpolates: CDF's interpolation
// between two centroids placed the best of two kills at the 30th
// percentile. Above a few hundred values the centroids are merged and the
// count is the t-digest's estimate, which is what it is for.
func (d *Digest) Placement(x float64) float64 {
	cs := d.centroids()
	t := total(cs)
	if t <= 1 {
		return 1
	}
	floor := x * (1 - 0.0025)
	var below float64
	for _, c := range cs {
		if c.Mean < floor {
			below += c.Weight
		}
	}
	return math.Min(below/(t-1), 1)
}

// MarshalBinary encodes the digest for the percentile_digests column.
// The encoding is deterministic: the same values in the same order
// always produce the same bytes.
func (d *Digest) MarshalBinary() ([]byte, error) {
	cs := d.centroids()
	out := make([]byte, 0, 4+1+8+8+len(cs)*16)
	out = append(out, magic[:]...)
	out = append(out, format)
	out = binary.BigEndian.AppendUint64(out, uint64(d.count))
	out = binary.BigEndian.AppendUint64(out, uint64(len(cs)))
	for _, c := range cs {
		out = binary.BigEndian.AppendUint64(out, math.Float64bits(c.Mean))
		out = binary.BigEndian.AppendUint64(out, math.Float64bits(c.Weight))
	}
	return out, nil
}

// Unmarshal reads a digest back. An empty input is an empty digest, so
// a missing row and a new one behave alike.
func Unmarshal(b []byte) (*Digest, error) {
	d := New()
	if len(b) == 0 {
		return d, nil
	}
	if len(b) < 21 || [4]byte{b[0], b[1], b[2], b[3]} != magic {
		return nil, fmt.Errorf("digest: not a digest")
	}
	if b[4] != format {
		return nil, fmt.Errorf("digest: unknown format %d", b[4])
	}
	d.count = int64(binary.BigEndian.Uint64(b[5:13]))
	declared := binary.BigEndian.Uint64(b[13:21])
	rem := len(b) - 21
	// Checked by division, not by computing 21+n*16: n comes straight off
	// the wire, and multiplying an attacker-controlled 64-bit count by 16
	// can overflow and coincidentally match len(b), which would let a
	// corrupt blob reach the decode loop below instead of being rejected.
	if rem%16 != 0 || declared != uint64(rem/16) {
		return nil, fmt.Errorf("digest: %d centroids do not fit %d bytes", declared, len(b))
	}
	n := int(declared)
	d.merged = make([]centroid, 0, n)
	for i := range n {
		off := 21 + i*16
		d.merged = append(d.merged, centroid{
			Mean:   math.Float64frombits(binary.BigEndian.Uint64(b[off : off+8])),
			Weight: math.Float64frombits(binary.BigEndian.Uint64(b[off+8 : off+16])),
		})
	}
	return d, nil
}

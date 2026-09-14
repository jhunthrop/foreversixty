// logs/engine/lexer/lexer_test.go
package lexer

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func collect(t *testing.T, l *Lexer, chunks []string, offsets []int64) []Line {
	t.Helper()
	var got []Line
	for i, c := range chunks {
		if err := l.Feed([]byte(c), offsets[i], func(ln Line) error {
			got = append(got, ln)
			return nil
		}); err != nil {
			t.Fatalf("Feed(%d): %v", i, err)
		}
	}
	if err := l.Flush(func(ln Line) error {
		got = append(got, ln)
		return nil
	}); err != nil {
		t.Fatalf("Flush: %v", err)
	}
	return got
}

func TestSplitParamsKeepsQuotedCommasAndNesting(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want []string
	}{
		{
			name: "quoted comma",
			in:   `SPELL_CAST_SUCCESS,Player-4184-000000A3,"Morrowlyn, the Cold",0x512,0x0`,
			want: []string{"SPELL_CAST_SUCCESS", "Player-4184-000000A3", "Morrowlyn, the Cold", "0x512", "0x0"},
		},
		{
			name: "escaped quote inside a quoted string",
			in:   `EMOTE,0000000000000000,"The \"Hollow\" Sentinel",0x0`,
			want: []string{"EMOTE", "0000000000000000", `The "Hollow" Sentinel`, "0x0"},
		},
		{
			name: "nested brackets and parens stay one field",
			in:   `COMBATANT_INFO,Player-4184-000000A1,0,(1,2,3),[(175850,183,(),(6788,1487),()),(0,0,(),(),())],184`,
			want: []string{
				"COMBATANT_INFO", "Player-4184-000000A1", "0", "(1,2,3)",
				"[(175850,183,(),(6788,1487),()),(0,0,(),(),())]", "184",
			},
		},
		{
			name: "empty trailing field",
			in:   `A,B,`,
			want: []string{"A", "B", ""},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := SplitParams(tc.in)
			if len(got) != len(tc.want) {
				t.Fatalf("got %d fields %q, want %d %q", len(got), got, len(tc.want), tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("field %d = %q, want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}

func TestFeedSplitsTimestampFromTheRecord(t *testing.T) {
	l := New()
	got := collect(t, l, []string{"9/26 20:10:04.600  SPELL_CAST_START,Player-4184-000000A3,\"Morrowlyn-Nightslayer\",0x512\n"}, []int64{0})
	if len(got) != 1 {
		t.Fatalf("got %d lines", len(got))
	}
	if got[0].Stamp != "9/26 20:10:04.600" {
		t.Errorf("stamp = %q", got[0].Stamp)
	}
	if got[0].Params[0] != "SPELL_CAST_START" || got[0].Params[2] != "Morrowlyn-Nightslayer" {
		t.Errorf("params = %q", got[0].Params)
	}
	if got[0].Number != 1 || got[0].Offset != 0 {
		t.Errorf("number = %d, offset = %d", got[0].Number, got[0].Offset)
	}
}

func TestFeedHandlesTabSeparatorAndMissingTimestamp(t *testing.T) {
	l := New()
	got := collect(t, l, []string{"9/26 20:10:04.600\tUNIT_DIED,A\nCOMBAT_LOG_VERSION,16\n"}, []int64{0})
	if len(got) != 2 {
		t.Fatalf("got %d lines", len(got))
	}
	if got[0].Stamp != "9/26 20:10:04.600" || got[0].Params[0] != "UNIT_DIED" {
		t.Errorf("tab line = %+v", got[0])
	}
	if got[1].Stamp != "" || got[1].Params[0] != "COMBAT_LOG_VERSION" {
		t.Errorf("headerless line = %+v", got[1])
	}
}

func TestFeedCarriesAPartialLineAcrossEveryChunkSize(t *testing.T) {
	src := "9/26 20:10:00.000  A,one\n9/26 20:10:01.000  B,two\n9/26 20:10:02.000  C,three\n"
	want := []string{"A", "B", "C"}
	for _, size := range []int{1, 7, 64, 4096} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			l := New()
			var chunks []string
			var offsets []int64
			for i := 0; i < len(src); i += size {
				end := min(i+size, len(src))
				chunks = append(chunks, src[i:end])
				offsets = append(offsets, int64(i))
			}
			got := collect(t, l, chunks, offsets)
			if len(got) != len(want) {
				t.Fatalf("got %d lines, want %d", len(got), len(want))
			}
			for i := range want {
				if got[i].Params[0] != want[i] {
					t.Errorf("line %d = %q, want %q", i, got[i].Params[0], want[i])
				}
				if got[i].Number != int64(i+1) {
					t.Errorf("line %d number = %d", i, got[i].Number)
				}
			}
			if got[1].Offset != 25 {
				t.Errorf("second line offset = %d, want 25", got[1].Offset)
			}
		})
	}
}

func TestFeedIgnoresBytesItHasAlreadySeen(t *testing.T) {
	src := "9/26 20:10:00.000  A,one\n9/26 20:10:01.000  B,two\n"
	l := New()
	got := collect(t, l, []string{src, src[10:], src}, []int64{0, 10, 0})
	if len(got) != 2 {
		t.Fatalf("got %d lines %v, want 2", len(got), names(got))
	}
	if l.NextOffset() != int64(len(src)) {
		t.Errorf("next offset = %d, want %d", l.NextOffset(), len(src))
	}
}

func TestFeedRejectsAGap(t *testing.T) {
	l := New()
	err := l.Feed([]byte("9/26 20:10:00.000  A,one\n"), 500, func(Line) error { return nil })
	if !errors.Is(err, ErrGap) {
		t.Fatalf("err = %v, want ErrGap", err)
	}
}

func TestFeedDropsAnOverlongLineAndResynchronises(t *testing.T) {
	l := New()
	long := strings.Repeat("x", MaxLineBytes+10)
	var got []Line
	emit := func(ln Line) error { got = append(got, ln); return nil }
	if err := l.Feed([]byte(long), 0, emit); err != nil {
		t.Fatalf("Feed long: %v", err)
	}
	if err := l.Feed([]byte("tail\n9/26 20:10:00.000  A,one\n"), int64(len(long)), emit); err != nil {
		t.Fatalf("Feed tail: %v", err)
	}
	if len(got) != 1 || got[0].Params[0] != "A" {
		t.Fatalf("got %v, want just the line after the resync", names(got))
	}
	if l.Dropped() != 1 {
		t.Errorf("dropped = %d, want 1", l.Dropped())
	}
}

func TestStateAndRestoreResumeMidLine(t *testing.T) {
	src := "9/26 20:10:00.000  A,one\n9/26 20:10:01.000  B,"
	l := New()
	var got []Line
	emit := func(ln Line) error { got = append(got, ln); return nil }
	if err := l.Feed([]byte(src), 0, emit); err != nil {
		t.Fatal(err)
	}
	revived := Restore(l.State())
	if err := revived.Feed([]byte("two\n"), int64(len(src)), emit); err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[1].Params[1] != "two" || got[1].Number != 2 {
		t.Fatalf("got %+v", got)
	}
}

func TestCarriageReturnsAreStripped(t *testing.T) {
	l := New()
	got := collect(t, l, []string{"9/26 20:10:00.000  A,one\r\n"}, []int64{0})
	if len(got) != 1 || got[0].Params[1] != "one" {
		t.Fatalf("got %+v", got)
	}
}

func names(ls []Line) []string {
	out := make([]string, len(ls))
	for i, l := range ls {
		out[i] = l.Params[0]
	}
	return out
}

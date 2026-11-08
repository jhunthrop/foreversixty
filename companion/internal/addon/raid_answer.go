// companion/internal/addon/raid_answer.go
// The addon's raid answers (docs/contracts/addon-v2.md, contract 4: outbox.raid_answer), the
// one part of the game's own file the companion edits. The answers ride up in the outbox like
// every other part; once the site has accepted or refused an answer, ClearRaidAnswers takes it
// out of ForeverSixtyDB.outbox.raid_answer, splicing that one list and leaving every other byte
// of the file as the game wrote it. A deferred answer stays and goes up again.
package addon

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

// OutboxRaidAnswer is the outbox key of the answers, and the name the inbox's accepts lists.
const OutboxRaidAnswer = "raid_answer"

// What the site answers for one raid answer (api/internal/addon raid_answers.go).
const (
	RaidAnswerAccepted = "accepted"
	RaidAnswerRefused  = "refused"
	RaidAnswerDeferred = "deferred"
)

// RaidAnswerResult is the site's answer for one entry, echoing the entry it answers.
type RaidAnswerResult struct {
	Character  string `json:"character"`
	RaidID     int64  `json:"raid_id"`
	AnsweredAt int64  `json:"answered_at"`
	Status     string `json:"status"`
	Reason     string `json:"reason,omitempty"`
}

// RaidAnswerKey names one entry in the game's file: its character, raid and click time.
type RaidAnswerKey struct {
	Character  string
	RaidID     int64
	AnsweredAt int64
}

// answeredKeys are the entries the site is done with: accepted or refused.
func answeredKeys(results []RaidAnswerResult) []RaidAnswerKey {
	var out []RaidAnswerKey
	for _, r := range results {
		if r.Status == RaidAnswerAccepted || r.Status == RaidAnswerRefused {
			out = append(out, RaidAnswerKey{r.Character, r.RaidID, r.AnsweredAt})
		}
	}
	return out
}

// raidAnswerHash is a short hash of the outbox's answers, "" with none, so the outbox's key
// moves when an answer is added or removed.
func raidAnswerHash(outbox *Table) (string, error) {
	answers, ok := outbox.Fields[OutboxRaidAnswer].(*Table)
	if !ok || (len(answers.Items) == 0 && len(answers.Fields) == 0) {
		return "", nil
	}
	b, err := json.Marshal(luaJSON(answers))
	if err != nil {
		return "", fmt.Errorf("the addon's raid answers cannot be sent: %w", err)
	}
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:8]), nil
}

// keyOfEntry is an entry's key as the game wrote it; a field it lacks reads as zero, as the
// site echoes an entry it could not read.
func keyOfEntry(v any) RaidAnswerKey {
	t, _ := v.(*Table)
	k := RaidAnswerKey{Character: t.Get("character")}
	if t != nil {
		if f, ok := t.Fields["raid_id"].(float64); ok {
			k.RaidID = int64(f)
		}
		if f, ok := t.Fields["answered_at"].(float64); ok {
			k.AnsweredAt = int64(f)
		}
	}
	return k
}

// span is a value's place in the source: src[start:end].
type span struct{ start, end int }

// ClearRaidAnswers removes the done entries from ForeverSixtyDB.outbox.raid_answer in the
// addon's SavedVariables file at path and reports how many it removed. It writes only when it
// removes something, only the addon's own file, and only after checking that everything in the
// file but that list reads the same afterwards.
func ClearRaidAnswers(path string, done []RaidAnswerKey) (int, error) {
	if filepath.Base(path) != SavedVariablesName+".lua" {
		return 0, fmt.Errorf("addon: refusing to edit %s: the companion edits only %s.lua", path, SavedVariablesName)
	}
	src, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	list, items, err := raidAnswerSpans(string(src))
	if err != nil || len(items) == 0 {
		return 0, err
	}
	drop := map[RaidAnswerKey]bool{}
	for _, k := range done {
		drop[k] = true
	}
	var kept []string
	for _, it := range items {
		v, err := (&parser{src: string(src), i: it.start}).value()
		if err != nil {
			return 0, err
		}
		if !drop[keyOfEntry(v)] {
			kept = append(kept, string(src[it.start:it.end]))
		}
	}
	removed := len(items) - len(kept)
	if removed == 0 {
		return 0, nil
	}
	out := spliceList(src, list, kept)
	if err := sameButAnswers(string(src), string(out)); err != nil {
		return 0, err
	}
	return removed, replaceFile(path, out)
}

// spliceList is src with the list at l replaced by a list of the kept entries' own text.
func spliceList(src []byte, l span, kept []string) []byte {
	indent := lineIndent(src, l.start)
	var b bytes.Buffer
	b.Write(src[:l.start])
	b.WriteString("{\n")
	for _, k := range kept {
		b.WriteString(indent + "\t" + k + ",\n")
	}
	b.WriteString(indent + "}")
	b.Write(src[l.end:])
	return b.Bytes()
}

// lineIndent is the whitespace that starts the line holding offset i.
func lineIndent(src []byte, i int) string {
	start := bytes.LastIndexByte(src[:i], '\n') + 1
	end := start
	for end < len(src) && (src[end] == ' ' || src[end] == '\t') {
		end++
	}
	return string(src[start:end])
}

// sameButAnswers checks that before and after read the same once the answers are set aside.
func sameButAnswers(before, after string) error {
	strip := func(src string) (map[string]any, error) {
		top, err := ParseLua(src)
		if err != nil {
			return nil, err
		}
		if db, ok := top["ForeverSixtyDB"].(*Table); ok {
			if ob, ok := db.Fields["outbox"].(*Table); ok {
				delete(ob.Fields, OutboxRaidAnswer)
			}
		}
		return top, nil
	}
	a, err := strip(before)
	if err != nil {
		return err
	}
	b, err := strip(after)
	if err != nil {
		return fmt.Errorf("addon: the edited saved variables would not read back: %w", err)
	}
	if !reflect.DeepEqual(a, b) {
		return fmt.Errorf("addon: the edit would change more than the raid answers; left the file alone")
	}
	return nil
}

// replaceFile writes body over path through a temporary file in the same folder, so the game
// never reads half a file.
func replaceFile(path string, body []byte) error {
	mode := os.FileMode(0o644)
	if fi, err := os.Stat(path); err == nil {
		mode = fi.Mode().Perm()
	}
	name := path + ".tmp"
	if err := os.WriteFile(name, body, mode); err != nil {
		return err
	}
	if err := os.Rename(name, path); err != nil {
		os.Remove(name)
		return err
	}
	return nil
}

// raidAnswerSpans finds ForeverSixtyDB.outbox.raid_answer in src: the list's own span and each
// positional entry's. A file without the list answers no spans.
func raidAnswerSpans(src string) (span, []span, error) {
	p := &parser{src: src}
	at, found, err := p.seekTop("ForeverSixtyDB", "outbox", OutboxRaidAnswer)
	if err != nil || !found {
		return span{}, nil, err
	}
	p.i = at
	if p.i >= len(p.src) || p.src[p.i] != '{' {
		return span{}, nil, nil // not a list: nothing the companion would edit
	}
	items, err := p.itemSpans()
	return span{at, p.i}, items, err
}

// seekTop finds the value at path below the top-level assignment path[0] and answers its
// offset.
func (p *parser) seekTop(path ...string) (int, bool, error) {
	for {
		p.space()
		if p.i >= len(p.src) {
			return 0, false, nil
		}
		name, err := p.name()
		if err != nil {
			return 0, false, err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return 0, false, err
		}
		if name == path[0] {
			return p.seek(path[1:])
		}
		if _, err := p.value(); err != nil {
			return 0, false, err
		}
	}
}

// seek finds the value at path inside the value at p.i and answers its offset.
func (p *parser) seek(path []string) (int, bool, error) {
	p.space()
	if len(path) == 0 {
		return p.i, true, nil
	}
	if p.i >= len(p.src) || p.src[p.i] != '{' {
		return 0, false, nil
	}
	p.i++
	for {
		p.space()
		if p.i >= len(p.src) {
			return 0, false, p.errf("the file ends inside a table")
		}
		if p.src[p.i] == '}' {
			return 0, false, nil
		}
		key, keyed, err := p.fieldKey()
		if err != nil {
			return 0, false, err
		}
		if keyed && key == path[0] {
			return p.seek(path[1:])
		}
		if _, err := p.value(); err != nil {
			return 0, false, err
		}
		p.separator()
	}
}

// fieldKey reads a field's `[key] =` or `name =`, if it has one, leaving p.i at its value.
func (p *parser) fieldKey() (string, bool, error) {
	switch {
	case p.src[p.i] == '[':
		p.i++
		k, err := p.value()
		if err != nil {
			return "", false, err
		}
		p.space()
		if err := p.expect(']'); err != nil {
			return "", false, err
		}
		p.space()
		return keyOf(k), true, p.expect('=')
	case isNameStart(p.src[p.i]) && p.assignmentAhead():
		name, err := p.name()
		if err != nil {
			return "", false, err
		}
		p.space()
		return name, true, p.expect('=')
	}
	return "", false, nil
}

// separator skips the space and one , or ; after a field.
func (p *parser) separator() {
	p.space()
	if p.i < len(p.src) && (p.src[p.i] == ',' || p.src[p.i] == ';') {
		p.i++
	}
}

// itemSpans reads the table at p.i and answers the span of each positional entry; a keyed
// field (which the addon never writes here) is left out, and so dropped by a splice.
func (p *parser) itemSpans() ([]span, error) {
	if err := p.expect('{'); err != nil {
		return nil, err
	}
	var out []span
	for {
		p.space()
		if p.i >= len(p.src) {
			return nil, p.errf("the file ends inside a table")
		}
		if p.src[p.i] == '}' {
			p.i++
			return out, nil
		}
		_, keyed, err := p.fieldKey()
		if err != nil {
			return nil, err
		}
		p.space()
		start := p.i
		if _, err := p.value(); err != nil {
			return nil, err
		}
		if !keyed {
			out = append(out, span{start, p.i})
		}
		p.separator()
	}
}

// acceptsLine is the inbox's accepts list (contract 3): the outbox entries beyond export,
// sim_request and roster_snapshot this companion uploads.
var acceptsLine = "\t[\"accepts\"] = {\n\t\t" + quote(OutboxRaidAnswer) + ",\n\t},\n"

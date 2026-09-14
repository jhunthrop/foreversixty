// companion/internal/addon/lua.go
// A reader for the subset of Lua that WoW writes into SavedVariables:
// top-level assignments of strings, numbers, booleans, nil and nested
// tables. It exists because the companion must read what the addon
// saved without embedding a Lua interpreter, and because the export
// strings themselves are opaque to it — FS1 is the addon's format,
// parsed on the site, never here.
package addon

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Table is a decoded Lua table: keyed entries in Fields, positional
// entries in Items, in the order they were written.
type Table struct {
	Fields map[string]any
	Items  []any
}

// Get returns a field as a string, empty when it is absent or is not
// a string.
func (t *Table) Get(key string) string {
	if t == nil {
		return ""
	}
	if v, ok := t.Fields[key].(string); ok {
		return v
	}
	return ""
}

// parser is a recursive-descent reader over the source.
type parser struct {
	src string
	i   int
}

// ParseLua reads a SavedVariables file into its top-level assignments.
func ParseLua(src string) (map[string]any, error) {
	p := &parser{src: src}
	out := map[string]any{}
	for {
		p.space()
		if p.i >= len(p.src) {
			return out, nil
		}
		name, err := p.name()
		if err != nil {
			return nil, err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return nil, err
		}
		v, err := p.value()
		if err != nil {
			return nil, err
		}
		out[name] = v
	}
}

func (p *parser) errf(format string, a ...any) error {
	line := 1 + strings.Count(p.src[:min(p.i, len(p.src))], "\n")
	return fmt.Errorf("lua line %d: %s", line, fmt.Sprintf(format, a...))
}

// space skips whitespace and -- comments, including --[[ blocks.
func (p *parser) space() {
	for p.i < len(p.src) {
		c := p.src[p.i]
		switch {
		case c == ' ' || c == '\t' || c == '\r' || c == '\n':
			p.i++
		case strings.HasPrefix(p.src[p.i:], "--[["):
			end := strings.Index(p.src[p.i:], "]]")
			if end < 0 {
				p.i = len(p.src)
				return
			}
			p.i += end + 2
		case strings.HasPrefix(p.src[p.i:], "--"):
			end := strings.IndexByte(p.src[p.i:], '\n')
			if end < 0 {
				p.i = len(p.src)
				return
			}
			p.i += end + 1
		default:
			return
		}
	}
}

func (p *parser) expect(c byte) error {
	if p.i >= len(p.src) || p.src[p.i] != c {
		return p.errf("expected %q", string(c))
	}
	p.i++
	return nil
}

func (p *parser) name() (string, error) {
	start := p.i
	for p.i < len(p.src) {
		c := rune(p.src[p.i])
		if unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_' || c == '.' {
			p.i++
			continue
		}
		break
	}
	if p.i == start {
		return "", p.errf("expected a name")
	}
	return p.src[start:p.i], nil
}

func (p *parser) value() (any, error) {
	p.space()
	if p.i >= len(p.src) {
		return nil, p.errf("expected a value")
	}
	switch c := p.src[p.i]; {
	case c == '"' || c == '\'':
		return p.str()
	case c == '{':
		return p.table()
	case c == '-' || (c >= '0' && c <= '9'):
		return p.number()
	default:
		word, err := p.name()
		if err != nil {
			return nil, err
		}
		switch word {
		case "true":
			return true, nil
		case "false":
			return false, nil
		case "nil":
			return nil, nil
		}
		return nil, p.errf("unexpected word %q", word)
	}
}

func (p *parser) str() (string, error) {
	quote := p.src[p.i]
	p.i++
	var b strings.Builder
	for p.i < len(p.src) {
		c := p.src[p.i]
		switch {
		case c == quote:
			p.i++
			return b.String(), nil
		case c == '\\':
			p.i++
			if p.i >= len(p.src) {
				return "", p.errf("the file ends inside a string")
			}
			e := p.src[p.i]
			switch {
			case e == 'n':
				b.WriteByte('\n')
				p.i++
			case e == 't':
				b.WriteByte('\t')
				p.i++
			case e == 'r':
				b.WriteByte('\r')
				p.i++
			case e >= '0' && e <= '9':
				j := p.i
				for j < len(p.src) && j-p.i < 3 && p.src[j] >= '0' && p.src[j] <= '9' {
					j++
				}
				n, err := strconv.Atoi(p.src[p.i:j])
				if err != nil || n > 255 {
					return "", p.errf("bad escape")
				}
				b.WriteByte(byte(n))
				p.i = j
			default:
				b.WriteByte(e)
				p.i++
			}
		default:
			b.WriteByte(c)
			p.i++
		}
	}
	return "", p.errf("the file ends inside a string")
}

func (p *parser) number() (float64, error) {
	start := p.i
	if p.src[p.i] == '-' {
		p.i++
	}
	for p.i < len(p.src) {
		c := p.src[p.i]
		if (c >= '0' && c <= '9') || c == '.' || c == 'e' || c == 'E' ||
			((c == '+' || c == '-') && (p.src[p.i-1] == 'e' || p.src[p.i-1] == 'E')) {
			p.i++
			continue
		}
		break
	}
	n, err := strconv.ParseFloat(p.src[start:p.i], 64)
	if err != nil {
		return 0, p.errf("bad number %q", p.src[start:p.i])
	}
	return n, nil
}

func (p *parser) table() (*Table, error) {
	if err := p.expect('{'); err != nil {
		return nil, err
	}
	t := &Table{Fields: map[string]any{}}
	for {
		p.space()
		if p.i >= len(p.src) {
			return nil, p.errf("the file ends inside a table")
		}
		if p.src[p.i] == '}' {
			p.i++
			return t, nil
		}
		if err := p.field(t); err != nil {
			return nil, err
		}
		p.space()
		if p.i < len(p.src) && (p.src[p.i] == ',' || p.src[p.i] == ';') {
			p.i++
		}
	}
}

func (p *parser) field(t *Table) error {
	p.space()
	switch {
	case p.src[p.i] == '[':
		p.i++
		k, err := p.value()
		if err != nil {
			return err
		}
		p.space()
		if err := p.expect(']'); err != nil {
			return err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return err
		}
		v, err := p.value()
		if err != nil {
			return err
		}
		t.Fields[keyOf(k)] = v
		return nil
	case isNameStart(p.src[p.i]) && p.assignmentAhead():
		name, err := p.name()
		if err != nil {
			return err
		}
		p.space()
		if err := p.expect('='); err != nil {
			return err
		}
		v, err := p.value()
		if err != nil {
			return err
		}
		t.Fields[name] = v
		return nil
	default:
		v, err := p.value()
		if err != nil {
			return err
		}
		t.Items = append(t.Items, v)
		return nil
	}
}

// assignmentAhead distinguishes `name = value` from a bare word value
// such as true, false or nil.
func (p *parser) assignmentAhead() bool {
	j := p.i
	for j < len(p.src) && (isNameStart(p.src[j]) || (p.src[j] >= '0' && p.src[j] <= '9')) {
		j++
	}
	for j < len(p.src) && (p.src[j] == ' ' || p.src[j] == '\t') {
		j++
	}
	return j < len(p.src) && p.src[j] == '='
}

func isNameStart(c byte) bool {
	return c == '_' || unicode.IsLetter(rune(c))
}

func keyOf(v any) string {
	switch k := v.(type) {
	case string:
		return k
	case float64:
		return strconv.FormatFloat(k, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(k)
	}
	return ""
}

// quote writes a Lua string literal.
func quote(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for i := range len(s) {
		switch c := s[i]; c {
		case '"':
			b.WriteString(`\"`)
		case '\\':
			b.WriteString(`\\`)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			if c < 0x20 {
				fmt.Fprintf(&b, `\%03d`, c)
				continue
			}
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
	return b.String()
}

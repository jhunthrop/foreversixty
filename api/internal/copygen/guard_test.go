package copygen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// The copy guard: the packages that speak to a reader outside the site (Discord, push, email,
// the app sign-in page) take their words from this package, generated from apps/shared/copy.
// A new sentence written inline in one of them fails here; add it to server.json (or reuse the
// en.json or shared.json key that already says it) and run `node web/scripts/gen-mobile-copy.mjs`.

// guarded is every package (or, for a package that mostly serves the API, the files) the guard
// reads, relative to this directory.
var guarded = []struct {
	dir   string
	files []string // empty reads every non-test Go file
}{
	{dir: "../discordbot"},
	{dir: "../notify"},
	{dir: "../billingmail"},
	{dir: "../mail"},
	{dir: "../auth", files: []string{"app_confirm.go", "signin_mail.go"}},
	{dir: "../subscribe", files: []string{"confirm_mail.go"}},
}

// allowedDecls are the package-level declarations whose literals are not copy, with why.
var allowedDecls = map[string]string{
	"pvpTitles":          "the game's PvP rank titles, data the item source line names (items/source-labels.ts)",
	"appConfirmFaces":    "font family names, which the page's @font-face rules declare",
	"invalidBodyMessage": "the webhook API's request shape, for a developer, not a player",
}

// notCopyCalls are the calls whose string arguments are never a reader's words: errors, logs,
// SQL and patterns.
var notCopyCalls = map[string]bool{
	"Errorf": true, "New": true, "Info": true, "Warn": true, "Error": true, "Debug": true,
	"InfoContext": true, "WarnContext": true, "ErrorContext": true, "DebugContext": true,
	"Printf": true, "Println": true, "Print": true, "panic": true, "WriteError": true, "fail": true,
	"Query": true, "QueryRow": true, "Exec": true, "MustCompile": true,
}

var (
	twoWords        = regexp.MustCompile(`\p{L}{2,}[ ,’']+\p{L}{2,}`)
	sqlText         = regexp.MustCompile(`\$\d|(?i)^\s*(select|insert|update|delete|with)\s`)
	placeholderOnly = regexp.MustCompile(`^\{[^}]*\}$`)
	errorPrefix     = regexp.MustCompile(`^[a-z][a-z0-9]*: `)
	// markup is a multi-line HTML or CSS template; its words come in through its fields.
	markup   = regexp.MustCompile(`(?s)\n.*(<[a-z!/]|\{[a-z-]+:)`)
	sentence = regexp.MustCompile(`\p{L}{2,}[ ,]+\p{L}+[ ,]+\p{L}+.*[.?:]$`)
)

// isProse reports whether a literal reads as words for a reader: two or more words, and either a
// capital first letter, a {placeholder}, or a sentence's end. SQL, a lone placeholder, a Go error's
// "pkg: " prefix and a multi-line HTML or CSS template are not prose. Lower-case fragments (" · your main") pass; the guard
// catches the sentences a lane would write inline, not every word.
func isProse(s string) bool {
	t := strings.TrimSpace(s)
	if markup.MatchString(t) {
		return false
	}
	if sqlText.MatchString(t) || placeholderOnly.MatchString(t) || errorPrefix.MatchString(t) || !twoWords.MatchString(t) {
		return false
	}
	if strings.IndexFunc(t, unicode.IsLower) < 0 {
		return false
	}
	first := []rune(t[strings.IndexFunc(t, unicode.IsLetter):])[0]
	return unicode.IsUpper(first) || strings.Contains(t, "{") || sentence.MatchString(t)
}

// inlineCopy is every prose literal in one file's source, as "line: text".
func inlineCopy(t *testing.T, name string, src []byte) []string {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, name, src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse %s: %v", name, err)
	}
	var found []string
	for _, decl := range f.Decls {
		if skipDecl(decl) {
			continue
		}
		var stack []ast.Node
		ast.Inspect(decl, func(n ast.Node) bool {
			if n == nil {
				stack = stack[:len(stack)-1]
				return true
			}
			stack = append(stack, n)
			if field, ok := n.(*ast.Field); ok && field.Tag != nil {
				return false
			}
			lit, ok := n.(*ast.BasicLit)
			if !ok || lit.Kind != token.STRING {
				return true
			}
			text, err := strconv.Unquote(lit.Value)
			if err == nil && isProse(text) && !insideNotCopyCall(stack) {
				found = append(found, strconv.Itoa(fset.Position(lit.Pos()).Line)+": "+text)
			}
			return true
		})
	}
	return found
}

// skipDecl reports whether a declaration holds no copy: imports, an allowed declaration, and an
// error type's Error method.
func skipDecl(decl ast.Decl) bool {
	switch d := decl.(type) {
	case *ast.GenDecl:
		if d.Tok == token.IMPORT {
			return true
		}
		for _, spec := range d.Specs {
			if vs, ok := spec.(*ast.ValueSpec); ok {
				for _, n := range vs.Names {
					if _, allowed := allowedDecls[n.Name]; allowed {
						return true
					}
				}
			}
		}
	case *ast.FuncDecl:
		return d.Recv != nil && d.Name.Name == "Error"
	}
	return false
}

func insideNotCopyCall(stack []ast.Node) bool {
	for i := len(stack) - 2; i >= 0; i-- {
		call, ok := stack[i].(*ast.CallExpr)
		if !ok {
			continue
		}
		switch fn := call.Fun.(type) {
		case *ast.Ident:
			if notCopyCalls[fn.Name] {
				return true
			}
		case *ast.SelectorExpr:
			if notCopyCalls[fn.Sel.Name] {
				return true
			}
		}
	}
	return false
}

func guardedFiles(t *testing.T, dir string, only []string) []string {
	t.Helper()
	if len(only) > 0 {
		out := make([]string, len(only))
		for i, f := range only {
			out[i] = filepath.Join(dir, f)
		}
		return out
	}
	all, err := filepath.Glob(filepath.Join(dir, "*.go"))
	if err != nil {
		t.Fatal(err)
	}
	return slices.DeleteFunc(all, func(f string) bool {
		return strings.HasSuffix(f, "_test.go") || strings.HasSuffix(f, "_gen.go")
	})
}

func TestServerWordsComeFromSharedCopy(t *testing.T) {
	for _, g := range guarded {
		files := guardedFiles(t, g.dir, g.files)
		if len(files) == 0 {
			t.Fatalf("%s: no Go files; the guard has lost a package", g.dir)
		}
		for _, file := range files {
			src, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("%s: %v", file, err)
			}
			for _, hit := range inlineCopy(t, file, src) {
				t.Errorf("%s:%s\n\twords for a reader belong in apps/shared/copy (server.json), read through copygen", file, hit)
			}
		}
	}
}

// The guard's self-check: it flags a sentence, a capitalised label and a placeholder line, and
// passes errors, logs, SQL, struct tags, an Error method, an allowed declaration, a lower-case
// fragment, a format and a multi-line template.
func TestTheGuardFindsInlineCopy(t *testing.T) {
	src := []byte(`package p

import "fmt"

type row struct {
	Name string ` + "`json:\"name\" doc:\"The Name of it\"`" + `
}

type stopErr struct{}

func (stopErr) Error() string { return "The night could not be saved." }

var pvpTitles = []string{"Knight Lieutenant"}

const sentence = "This raid has ended."

var labels = map[string]string{"in": "Signed up"}

func f(name string) string {
	_ = fmt.Errorf("Could not load %s: %w", name, nil)
	log.Info("Discord did not answer", "err", nil)
	_ = "select id from raids where id = $1"
	_ = " · your main"
	_ = "%s · %s"
	_ = "Template Words\n<b>More Words</b>"
	return "{Name} is In."
}
`)
	got := inlineCopy(t, "self.go", src)
	want := []string{"15: This raid has ended.", "17: Signed up", "26: {Name} is In."}
	if !slices.Equal(got, want) {
		t.Fatalf("guard found %q, want %q", got, want)
	}
}

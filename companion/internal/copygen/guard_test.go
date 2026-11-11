package copygen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// The copy guard (the companion's surface guard): every package of the companion takes the words
// a player reads from this package, generated from apps/shared/copy/companion.json. A new sentence
// written inline anywhere under internal/ or cmd/ fails here; add it to companion.json and run
// `node web/scripts/gen-mobile-copy.mjs`. The window's script has its own guard (ui/copy.test.js).

// roots are the trees the guard walks, relative to this directory; skipped are the directories
// in them that hold no player's words.
var roots = []string{"..", "../../cmd"}

var skipped = map[string]string{
	"copygen":  "the generated words themselves",
	"fakeapi":  "a test double of the site's API, run only by tests",
	"testdata": "test inputs",
	"fixture":  "the synthetic combat log the tests parse",
}

// allowedDecls are the package-level declarations whose literals are not copy, with why.
var allowedDecls = map[string]string{
	"allowedSample": "the self-check's own allowed declaration",
	// Names the system reads, not sentences: the bundle, the Apps & Features entry, the keychain.
	"AppName":     "the Mac app bundle's file name (install)",
	"ProductName": "the Windows Apps & Features entry and Start menu name (install)",
	"Publisher":   "the company Windows names as the publisher (install)",
	"Service":     "the keychain item's service name (secret)",
	// companion.log's words for the site's answer to an outbox: a log, read by whoever opens it.
	"CopyOutboxIgnored": "the companion.log lines block (addon/outbox.go)",
}

// allowedFuncs are the functions whose literals are not copy, with why.
var allowedFuncs = map[string]string{
	"RenderInbox":   "the header comment of the Lua file the addon reads (addon/addon.go), never on screen",
	"Lines":         "companion.log lines for the site's outbox answer (addon/outbox.go)",
	"candidatesFor": "the folders the game installs into, paths not words (wow/wow.go)",
}

// notCopyCalls are the calls whose string arguments are never a reader's words: errors, logs,
// SQL and patterns.
var notCopyCalls = map[string]bool{
	"Errorf": true, "New": true, "Info": true, "Warn": true, "Error": true, "Debug": true,
	"InfoContext": true, "WarnContext": true, "ErrorContext": true, "DebugContext": true,
	"Printf": true, "Println": true, "Print": true, "panic": true, "WriteError": true, "fail": true,
	"MustCompile": true, "Fprintln": true, "Fprintf": true,
	// exec.Command's arguments are a program's, not a player's.
	"Command": true,
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
		if _, allowed := allowedFuncs[d.Name.Name]; allowed {
			return true
		}
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

// guardedFiles is every non-test, non-generated Go file under the roots.
func guardedFiles(t *testing.T) []string {
	t.Helper()
	var files []string
	for _, root := range roots {
		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				if _, skip := skipped[d.Name()]; skip {
					return filepath.SkipDir
				}
				return nil
			}
			name := d.Name()
			if strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go") && !strings.HasSuffix(name, "_gen.go") {
				files = append(files, path)
			}
			return nil
		})
		if err != nil {
			t.Fatal(err)
		}
	}
	return files
}

func TestCompanionWordsComeFromSharedCopy(t *testing.T) {
	files := guardedFiles(t)
	if len(files) < 20 {
		t.Fatalf("the guard read %d files; it has lost the companion's packages", len(files))
	}
	for _, file := range files {
		src, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("%s: %v", file, err)
		}
		for _, hit := range inlineCopy(t, file, src) {
			t.Errorf("%s:%s\n\twords for a player belong in apps/shared/copy/companion.json, read through copygen", file, hit)
		}
	}
}

// The guard's self-check: it flags a sentence, a capitalised label and a placeholder line, and
// passes errors, logs, struct tags, an Error method, an allowed declaration, a lower-case
// fragment, a format and a multi-line template.
func TestTheGuardFindsInlineCopy(t *testing.T) {
	src := []byte(`package p

import "fmt"

type row struct {
	Name string ` + "`json:\"name\" doc:\"The Name of it\"`" + `
}

type stopErr struct{}

func (stopErr) Error() string { return "The upload could not be saved." }

var allowedSample = []string{"Allowed Words"}

const sentence = "Fights upload as they end."

var labels = map[string]string{"open": "Open your reports"}

func f(name string) string {
	_ = fmt.Errorf("Could not read %s: %w", name, nil)
	log.Info("The site did not answer", "err", nil)
	_ = " · up to date"
	_ = "%s · %s"
	_ = "Template Words\n<b>More Words</b>"
	return "{folder} is not a game folder."
}
`)
	got := inlineCopy(t, "self.go", src)
	want := []string{"15: Fights upload as they end.", "17: Open your reports", "25: {folder} is not a game folder."}
	if !slices.Equal(got, want) {
		t.Fatalf("guard found %q, want %q", got, want)
	}
}

package activity

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
)

func isNoRows(err error) bool { return errors.Is(err, pgx.ErrNoRows) }

func cutPipe(s string) (string, string, bool) { return strings.Cut(s, "|") }

func plural(n int, one, many string) string {
	if n == 1 {
		return one
	}
	return many
}

func lowerASCII(s string) string { return strings.ToLower(strings.TrimSpace(s)) }

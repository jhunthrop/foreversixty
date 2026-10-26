package notify

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// failureRule is the five-failure rule over one webhook table: a counted failure adds one
// to the consecutive-failure counter and, at MaxFailures, switches the webhook off by
// setting offColumn to offValue (kept when already set). An uncounted failure (a 429, a
// manual test) only stores the last error. The account and guild stores share it.
type failureRule struct {
	table string
	// where is the row key with its placeholders numbered from $4 (after the rule's own three).
	where     string
	offColumn string
	offValue  string
}

// record applies the rule to the row keys select and reports whether the webhook is off.
// ErrNotFound when no row matches.
func (f failureRule) record(ctx context.Context, pool *pgxpool.Pool, lastErr string, counts bool, keys ...any) (bool, error) {
	step := 0
	if counts {
		step = 1
	}
	sql := fmt.Sprintf(`
		update %[1]s set
		  failures = failures + $1,
		  last_error = $2,
		  %[3]s = case when failures + $1 >= $3 then coalesce(%[3]s, %[4]s) else %[3]s end
		where %[2]s
		returning %[3]s is not null`, f.table, f.where, f.offColumn, f.offValue)
	var disabled bool
	err := pool.QueryRow(ctx, sql, append([]any{step, lastErr, MaxFailures}, keys...)...).Scan(&disabled)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, ErrNotFound
	}
	if err != nil {
		return false, fmt.Errorf("notify: record failure %s %v: %w", f.table, keys, err)
	}
	return disabled, nil
}

// accountFailures is the rule over the members' own webhooks: off is a timestamp.
var accountFailures = failureRule{
	table: "notification_webhooks", where: "user_id = $4 and kind = $5",
	offColumn: "disabled_at", offValue: "now()",
}

// guildFailures is the rule over the guilds' raid channels: off is the reason the
// settings page reads, worded as the account's.
var guildFailures = failureRule{
	table: "guild_notification_webhooks", where: "guild_id = $4",
	offColumn: "disabled_reason", offValue: "'" + reasonFailedFiveTimes + "'",
}

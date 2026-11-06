package guilds

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

const (
	csvContentType   = "text/csv; charset=utf-8"
	exportDateLayout = "2006-01-02"
	exportFailText   = "could not build that export just now"
)

var (
	attendanceExportHeader = []string{"date", "raid", "character", "class", "signed", "outcome", "excused", "pulls_present", "pulls_total", "fixed"}
	lootExportHeader       = []string{"date", "raid", "boss", "item_id", "item", "character"}
	nonAlphanumeric        = regexp.MustCompile(`[^a-z0-9]+`)
)

// fallbackExportSlug names the file of a guild whose name has no ASCII letters or digits.
const fallbackExportSlug = "guild"

// exportFlushEvery is how many rows are written between flushes to the client.
const exportFlushEvery = 500

// formulaPrefixes are the first characters a spreadsheet reads as a formula.
const formulaPrefixes = "=+-@\t\r"

// csvSafe blocks formula injection: a cell that starts with a formula character gets an
// apostrophe in front, so a spreadsheet shows it as text.
func csvSafe(cell string) string {
	if cell != "" && strings.ContainsRune(formulaPrefixes, rune(cell[0])) {
		return "'" + cell
	}
	return cell
}

// exportSlug is a guild name as a filename stem: lower-cased, each run of non-alphanumerics
// one dash, trimmed of dashes.
func exportSlug(name string) string {
	slug := strings.Trim(nonAlphanumeric.ReplaceAllString(strings.ToLower(name), "-"), "-")
	if slug == "" {
		return fallbackExportSlug
	}
	return slug
}

// itemNames resolves an award's item the way Store.Loot does, from the fixed table.
var itemNames = func() map[int]string {
	names := make(map[int]string, len(onyxiaLoot))
	for _, item := range onyxiaLoot {
		names[item.ItemID] = item.Name
	}
	return names
}()

// exportAttendance is GET /v1/guilds/{id}/export/attendance.csv.
func (s *Service) exportAttendance(w http.ResponseWriter, r *http.Request) {
	s.serveExport(w, r, "attendance", attendanceExportHeader, s.Store.streamAttendance)
}

// exportLoot is GET /v1/guilds/{id}/export/loot.csv.
func (s *Service) exportLoot(w http.ResponseWriter, r *http.Request) {
	s.serveExport(w, r, "loot", lootExportHeader, s.Store.streamLoot)
}

// serveExport answers an officer-rank CSV read. Like the rest of the officer reads it never
// consults the plan, so a lapsed guild keeps its history; anyone below officer rank gets 404.
func (s *Service) serveExport(w http.ResponseWriter, r *http.Request, kind string, header []string,
	stream rowStream) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "export_"+kind, err, exportFailText)
		return
	}
	if !allowed && !auth.ActorFrom(r.Context()).IsModerator() {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	guild, err := s.Store.getGuild(r.Context(), guildID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
			return
		}
		s.fail(w, r, "export_"+kind, err, exportFailText)
		return
	}
	err = writeExport(r.Context(), w, guildID, header, stream, func() {
		httpx.CachePrivate(w)
		w.Header().Set("Content-Type", csvContentType)
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s-%s.csv"`, exportSlug(guild.Name), kind))
	})
	var early heldError
	switch {
	case errors.As(err, &early):
		s.fail(w, r, "export_"+kind, early.err, exportFailText)
	case err != nil:
		// The status line is already sent; all that is left is to say so in the log.
		s.logger().Warn("guilds", "id", httpx.RequestIDFrom(r.Context()), "op", "export_"+kind, "guild_id", guildID, "err", err)
	}
}

// heldError is a failure before any byte of the file reached the client, so
// the request can still answer an error instead of a short file.
type heldError struct{ err error }

func (e heldError) Error() string { return e.err.Error() }
func (e heldError) Unwrap() error { return e.err }

// writeExport streams the CSV. The first exportFlushEvery rows are held back,
// so a query that fails early answers an error, not a header-only file;
// headers runs just before the first byte goes out.
func writeExport(ctx context.Context, w http.ResponseWriter, guildID int64, header []string, stream rowStream,
	headers func()) error {
	held := &heldWriter{w: w, onOpen: headers}
	out := csv.NewWriter(held)
	_ = out.Write(header)
	written := 0
	err := stream(ctx, guildID, func(row []string) error {
		if err := out.Write(row); err != nil {
			return err
		}
		if written++; written%exportFlushEvery == 0 {
			return held.flush(out)
		}
		return nil
	})
	switch {
	case err != nil && !held.open:
		return heldError{err: err}
	case err != nil:
		return err
	}
	return held.flush(out)
}

// heldWriter keeps the response's bytes until flush opens it.
type heldWriter struct {
	w      http.ResponseWriter
	onOpen func()
	buf    bytes.Buffer
	open   bool
}

func (h *heldWriter) Write(p []byte) (int, error) {
	if h.open {
		return h.w.Write(p)
	}
	return h.buf.Write(p)
}

// flush sends what out has written, opening the response on the first call.
func (h *heldWriter) flush(out *csv.Writer) error {
	out.Flush()
	if err := out.Error(); err != nil {
		return err
	}
	if h.open {
		return nil
	}
	h.open = true
	h.onOpen()
	_, err := h.w.Write(h.buf.Bytes())
	h.buf.Reset()
	return err
}

// rowStream feeds each CSV row of a guild's export to emit as it is read.
type rowStream func(ctx context.Context, guildID int64, emit func([]string) error) error

// lootNightUnknown is the raid cell of an award that names no raid and was made from no
// report: its date cell is then the day it was awarded.
const lootNightUnknown = "No raid night recorded (date is the award date)"

// streamLoot names each award's character by its export. An award to an invited member names
// the account's synthetic row (syntheticKey), which has no export: it reads as the account's
// main (else first) character, else the stand-in, the rule inviteeName applies on the roster.
// Each award is dated by the raid night it belongs to: its own raid's, else the raid holding the
// report it was made from, else that report's night; with neither, the award date, said so.
func (s *Store) streamLoot(ctx context.Context, guildID int64, emit func([]string) error) error {
	rows, err := s.Pool.Query(ctx, `
		select coalesce(r.starts_at, rr.starts_at, `+logAnchorSQL("rp")+`, la.awarded_at) as night,
		       coalesce(r.title, rr.title, nullif(rp.zone, ''), nullif(rp.log_zone, ''), nullif(rp.title, ''), '`+lootNightUnknown+`'),
		       la.encounter_id, la.item_id,
		       coalesce(nullif(ae.name, ''), inv.name,
		                case when u.id is not null then '`+invitedStandIn+` ' || u.id end, la.character_key)
		from loot_awards la
		left join guild_raids r on r.id = la.raid_id
		left join reports rp on rp.id = la.report_id
		left join lateral (
		  select g.starts_at, g.title from guild_raids g
		  where g.report_id = la.report_id and g.guild_id = la.guild_id
		  order by g.starts_at desc limit 1
		) rr on true
		left join addon_exports ae on ae.character_key = la.character_key
		left join users u on la.character_key like 'account:%'
		  and u.id = nullif(substring(la.character_key from '^account:([0-9]+)$'), '')::bigint
		left join lateral (
		  select c.name from characters c where c.user_id = u.id
		  order by (c.key = u.main_character_key) desc, c.name limit 1
		) inv on true
		where la.guild_id = $1
		order by night, la.awarded_at, la.id`, guildID)
	if err != nil {
		return fmt.Errorf("guilds: export loot: %w", err)
	}
	defer rows.Close()
	loc, err := s.exportLocation(ctx, guildID)
	if err != nil {
		return err
	}
	for rows.Next() {
		var (
			night             time.Time
			raid, character   string
			encounterID, item int64
		)
		if err := rows.Scan(&night, &raid, &encounterID, &item, &character); err != nil {
			return fmt.Errorf("guilds: export loot: %w", err)
		}
		boss := ""
		if encounterID == onyxiaEncounterID {
			boss = encounterName
		}
		if err := emit(safeRow(night.In(loc).Format(exportDateLayout), raid, boss,
			strconv.FormatInt(item, 10), itemNames[int(item)], character)); err != nil {
			return err
		}
	}
	return rows.Err()
}

// exportLocation is the zone both CSVs write night dates in: the guild's raid zone, the one
// every planned raid is posted and shown in, so a Wednesday 7 pm raid in Chicago is dated
// Wednesday. A logged night with no planned raid uses the same zone. A guild that never set a
// raid zone gets UTC: a report stores no time zone of its own (reports.zone is the instance),
// so there is nothing closer to the players' clock to read.
func (s *Store) exportLocation(ctx context.Context, guildID int64) (*time.Location, error) {
	tz, err := s.GuildRaidTimezone(ctx, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: export zone: %w", err)
	}
	return zoneOrUTC(tz), nil
}

// safeRow escapes every cell of a row against formula injection.
func safeRow(cells ...string) []string {
	for i, c := range cells {
		cells[i] = csvSafe(c)
	}
	return cells
}

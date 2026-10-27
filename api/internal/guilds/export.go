package guilds

import (
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
	attendanceExportHeader = []string{"date", "raid", "character", "class", "signed", "outcome", "excused", "pulls_present", "pulls_total"}
	lootExportHeader       = []string{"date", "raid", "boss", "item_id", "item", "character"}
	nonAlphanumeric        = regexp.MustCompile(`[^a-z0-9]+`)
)

// formulaPrefixes are the first characters a spreadsheet reads as a formula.
const formulaPrefixes = "=+-@"

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
	return strings.Trim(nonAlphanumeric.ReplaceAllString(strings.ToLower(name), "-"), "-")
}

// itemNameByID resolves an award's item the way Store.Loot does, from the fixed table.
func itemNameByID(itemID int) string {
	for _, item := range onyxiaLoot {
		if item.ItemID == itemID {
			return item.Name
		}
	}
	return ""
}

// exportAttendance is GET /v1/guilds/{id}/export/attendance.csv.
func (s *Service) exportAttendance(w http.ResponseWriter, r *http.Request) {
	s.serveExport(w, r, "attendance", attendanceExportHeader, s.Store.attendanceExportRows)
}

// exportLoot is GET /v1/guilds/{id}/export/loot.csv.
func (s *Service) exportLoot(w http.ResponseWriter, r *http.Request) {
	s.serveExport(w, r, "loot", lootExportHeader, s.Store.lootExportRows)
}

// serveExport answers an officer-rank CSV read. Like the rest of the officer reads it never
// consults the plan, so a lapsed guild keeps its history; anyone below officer rank gets 404.
func (s *Service) serveExport(w http.ResponseWriter, r *http.Request, kind string, header []string,
	load func(ctx context.Context, guildID int64) ([][]string, error)) {
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
	rows, err := load(r.Context(), guildID)
	if err != nil {
		s.fail(w, r, "export_"+kind, err, exportFailText)
		return
	}
	httpx.CachePrivate(w)
	w.Header().Set("Content-Type", csvContentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s-%s.csv"`, exportSlug(guild.Name), kind))
	out := csv.NewWriter(w)
	_ = out.Write(header)
	_ = out.WriteAll(rows) // WriteAll flushes; a broken client connection has nobody to tell
}

func (s *Store) attendanceExportRows(ctx context.Context, guildID int64) ([][]string, error) {
	rows, err := s.Pool.Query(ctx, `
		select r.starts_at, r.title, a.name, a.class, coalesce(a.signed, ''), a.outcome, a.excused,
		       a.pulls_present, a.pulls_total
		from guild_raid_attendance a join guild_raids r on r.id = a.raid_id
		where r.guild_id = $1
		order by r.starts_at, a.name, a.character_key`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: export attendance: %w", err)
	}
	defer rows.Close()
	out := [][]string{}
	for rows.Next() {
		var (
			startsAt                         time.Time
			title, name, class, signed, kind string
			excused                          bool
			present, total                   int
		)
		if err := rows.Scan(&startsAt, &title, &name, &class, &signed, &kind, &excused, &present, &total); err != nil {
			return nil, fmt.Errorf("guilds: export attendance: %w", err)
		}
		out = append(out, safeRow(startsAt.UTC().Format(exportDateLayout), title, name, class, signed, kind,
			strconv.FormatBool(excused), strconv.Itoa(present), strconv.Itoa(total)))
	}
	return out, rows.Err()
}

func (s *Store) lootExportRows(ctx context.Context, guildID int64) ([][]string, error) {
	rows, err := s.Pool.Query(ctx, `
		select la.awarded_at, coalesce(r.title, ''), la.encounter_id, la.item_id, la.character_key
		from loot_awards la left join guild_raids r on r.id = la.raid_id
		where la.guild_id = $1
		order by la.awarded_at, la.id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: export loot: %w", err)
	}
	defer rows.Close()
	out := [][]string{}
	for rows.Next() {
		var (
			awardedAt         time.Time
			raid, key         string
			encounterID, item int64
		)
		if err := rows.Scan(&awardedAt, &raid, &encounterID, &item, &key); err != nil {
			return nil, fmt.Errorf("guilds: export loot: %w", err)
		}
		boss := ""
		if encounterID == onyxiaEncounterID {
			boss = encounterName
		}
		out = append(out, safeRow(awardedAt.UTC().Format(exportDateLayout), raid, boss,
			strconv.FormatInt(item, 10), itemNameByID(int(item)), key))
	}
	return out, rows.Err()
}

// safeRow escapes every cell of a row against formula injection.
func safeRow(cells ...string) []string {
	for i, c := range cells {
		cells[i] = csvSafe(c)
	}
	return cells
}

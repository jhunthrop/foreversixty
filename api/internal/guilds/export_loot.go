// api/internal/guilds/export_loot.go
//
// Loot history, as CSV (GET .../export/loot.csv) and JSON (GET .../export/loot.json), for an
// optional range of night dates (?from=YYYY-MM-DD&to=YYYY-MM-DD, inclusive, in the guild's raid
// zone as the attendance CSV dates them). Each award carries its decision record: main or off
// spec, the votes at award time, the reason and the officer who awarded it. A field the award
// never recorded (an award made before the council) is empty in the CSV and null in the JSON.
package guilds

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// lootNightUnknown is the raid cell of an award that names no raid and was made from no
// report: its date cell is then the day it was awarded.
const lootNightUnknown = "No raid night recorded (date is the award date)"

// LootHistoryRow is one award of the loot history.
type LootHistoryRow struct {
	Date      string  `json:"date"`
	Raid      string  `json:"raid"`
	Boss      string  `json:"boss"`
	ItemID    int     `json:"item_id"`
	Item      string  `json:"item"`
	Character string  `json:"character"`
	Class     string  `json:"class"`
	MsOs      *string `json:"ms_os"`
	VotesYes  *int    `json:"votes_yes"`
	VotesNo   *int    `json:"votes_no"`
	Reason    *string `json:"reason"`
	AwardedBy string  `json:"awarded_by"`
}

// LootHistoryView is the body of GET .../export/loot.json.
type LootHistoryView struct {
	From *string          `json:"from"`
	To   *string          `json:"to"`
	Rows []LootHistoryRow `json:"rows"`
}

// dateRange is an inclusive range of local night dates; an empty end is open.
type dateRange struct{ from, to string }

func (d dateRange) holds(date string) bool {
	return (d.from == "" || date >= d.from) && (d.to == "" || date <= d.to)
}

// dateRangeFrom reads ?from and ?to, writing the refusal and returning false when either is not a
// date or from is after to.
func dateRangeFrom(w http.ResponseWriter, r *http.Request) (dateRange, bool) {
	q := r.URL.Query()
	d := dateRange{from: q.Get("from"), to: q.Get("to")}
	for _, v := range []string{d.from, d.to} {
		if _, err := time.Parse(exportDateLayout, v); v != "" && err != nil {
			httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "from and to are dates, YYYY-MM-DD", nil)
			return d, false
		}
	}
	if d.from != "" && d.to != "" && d.from > d.to {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "from is after to", nil)
		return d, false
	}
	return d, true
}

// exportLoot is GET /v1/guilds/{id}/export/loot.csv.
func (s *Service) exportLoot(w http.ResponseWriter, r *http.Request) {
	rng, ok := dateRangeFrom(w, r)
	if !ok {
		return
	}
	s.serveExport(w, r, "loot", lootExportHeader, func(ctx context.Context, guildID int64, emit func([]string) error) error {
		rows, err := s.Store.lootHistory(ctx, guildID, rng)
		if err != nil {
			return err
		}
		for _, row := range rows {
			if err := emit(row.cells()); err != nil {
				return err
			}
		}
		return nil
	})
}

// exportLootJSON is GET /v1/guilds/{id}/export/loot.json.
func (s *Service) exportLootJSON(w http.ResponseWriter, r *http.Request) {
	rng, ok := dateRangeFrom(w, r)
	if !ok {
		return
	}
	guildID, _, ok := s.exportGuild(w, r, "loot")
	if !ok {
		return
	}
	rows, err := s.Store.lootHistory(r.Context(), guildID, rng)
	if err != nil {
		s.fail(w, r, "export_loot", err, exportFailText)
		return
	}
	view := LootHistoryView{Rows: rows}
	if rng.from != "" {
		view.From = &rng.from
	}
	if rng.to != "" {
		view.To = &rng.to
	}
	httpx.CachePrivate(w)
	httpx.WriteOK(w, r, http.StatusOK, view)
}

// cells is the row as formula-safe CSV cells.
func (row LootHistoryRow) cells() []string {
	return safeRow(row.Date, row.Raid, row.Boss, strconv.Itoa(row.ItemID), row.Item, row.Character, row.Class,
		optString(row.MsOs), optInt(row.VotesYes), optInt(row.VotesNo), optString(row.Reason), row.AwardedBy)
}

func optString(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func optInt(v *int) string {
	if v == nil {
		return ""
	}
	return strconv.Itoa(*v)
}

// lootHistory is every award of the guild in rng, by night then award time. It names each award's
// character by its export; an award to an invited member names the account's synthetic row
// (syntheticKey), which has no export: it reads as the account's main (else first) character,
// else the stand-in, the rule inviteeName applies on the roster. Each award is dated by the raid
// night it belongs to: its own raid's, else the raid holding the report it was made from, else
// that report's night; with neither, the award date, said so.
func (s *Store) lootHistory(ctx context.Context, guildID int64, rng dateRange) ([]LootHistoryRow, error) {
	loc, err := s.exportLocation(ctx, guildID)
	if err != nil {
		return nil, err
	}
	rows, err := s.Pool.Query(ctx, `
		select coalesce(r.starts_at, rr.starts_at, `+logAnchorSQL("rp")+`, la.awarded_at) as night,
		       coalesce(r.title, rr.title, nullif(rp.zone, ''), nullif(rp.log_zone, ''), nullif(rp.title, ''), '`+lootNightUnknown+`'),
		       la.encounter_id, la.item_id,
		       coalesce(nullif(ae.name, ''), inv.name,
		                case when u.id is not null then '`+invitedStandIn+` ' || u.id end, la.character_key),
		       coalesce(c.class, inv.class, ''), la.ms_os, la.votes_yes, la.votes_no, la.reason, coalesce(la.awarded_by, 0)
		from loot_awards la
		left join guild_raids r on r.id = la.raid_id
		left join reports rp on rp.id = la.report_id
		left join lateral (
		  select g.starts_at, g.title from guild_raids g
		  where g.report_id = la.report_id and g.guild_id = la.guild_id
		  order by g.starts_at desc limit 1
		) rr on true
		left join addon_exports ae on ae.character_key = la.character_key
		left join characters c on c.key = la.character_key
		left join users u on la.character_key like 'account:%'
		  and u.id = nullif(substring(la.character_key from '^account:([0-9]+)$'), '')::bigint
		left join lateral (
		  select c2.name, c2.class from characters c2 where c2.user_id = u.id
		  order by (c2.key = u.main_character_key) desc, c2.name limit 1
		) inv on true
		where la.guild_id = $1
		order by night, la.awarded_at, la.id`, guildID)
	if err != nil {
		return nil, fmt.Errorf("guilds: export loot: %w", err)
	}
	defer rows.Close()
	out := []LootHistoryRow{}
	var awarders []int64
	for rows.Next() {
		var (
			night       time.Time
			encounterID int64
			awardedBy   int64
			row         LootHistoryRow
		)
		if err := rows.Scan(&night, &row.Raid, &encounterID, &row.ItemID, &row.Character, &row.Class,
			&row.MsOs, &row.VotesYes, &row.VotesNo, &row.Reason, &awardedBy); err != nil {
			return nil, fmt.Errorf("guilds: export loot: %w", err)
		}
		row.Date = night.In(loc).Format(exportDateLayout)
		if !rng.holds(row.Date) {
			continue
		}
		if _, ok := lootBossByEncounter(encounterID); ok {
			row.Boss = bossName(encounterID)
		}
		row.Item = itemNames[row.ItemID]
		out = append(out, row)
		awarders = append(awarders, awardedBy)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("guilds: export loot: %w", err)
	}
	return out, s.nameHistoryAwarders(ctx, guildID, out, awarders)
}

// nameHistoryAwarders fills awarded_by with the officer's character (awarderNames), "" for an
// award with no recorded officer.
func (s *Store) nameHistoryAwarders(ctx context.Context, guildID int64, rows []LootHistoryRow, awarders []int64) error {
	names, err := s.awarderNamesOf(ctx, guildID, awarders)
	if err != nil {
		return err
	}
	for i, id := range awarders {
		if id != 0 {
			rows[i].AwardedBy = names.of(id)
		}
	}
	return nil
}

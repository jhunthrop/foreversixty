package activity

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/rating/presence"
	"github.com/jhunthrop/foreversixty/api/internal/textfmt"
)

// ReportWriter writes the feed rows a completed report earns (rulings P1, P2): its upload,
// and the first site kill of each raid boss it holds. It satisfies reports.Completed. Only a
// public, unflagged, complete report writes anything.
type ReportWriter struct {
	Store *Store
	// RaidZones are the log zone names of the raids (data/curated/raids/first-tier.json); a
	// kill in one of them may be the site's first.
	RaidZones map[string]bool
	Log       *slog.Logger
}

// reportFacts is what the sentences need from one report.
type reportFacts struct {
	id, zone, key, name, class, spec string
	level, bosses                    int
	anonymous                        bool
	durationMS                       int64
	at                               time.Time
}

var reportFactsSQL = `select r.id, coalesce(nullif(r.log_zone, ''), r.zone), coalesce(r.logging_character, ''),
	coalesce(c.name, ''), coalesce(c.class, ''), coalesce(c.level, 0), coalesce(u.anonymize, false),
	(select count(distinct f.encounter_id) from fights f where f.report_id = r.id and f.kill and f.encounter_id is not null),
	(select coalesce(max(f.start_ms + f.duration_ms) - min(f.start_ms), 0) from fights f where f.report_id = r.id),
	coalesce((select m.class || '|' || m.spec from fight_metrics m
	          where m.report_id = r.id and m.player_key = r.logging_character and m.spec is not null limit 1), ''),
	coalesce(r.completed_at, now())
	from reports r
	left join characters c on c.key = r.logging_character
	left join users u on u.id = r.owner_id
	where r.id = $1 and r.visibility = 'public' and ` + presence.EvidenceReport("r") + ` and r.status = 'complete'`

// ReportCompleted writes the report's rows. It never fails the completion: a failure is logged.
func (w *ReportWriter) ReportCompleted(ctx context.Context, reportID string) {
	if w == nil || w.Store == nil {
		return
	}
	f, ok, err := w.facts(ctx, reportID)
	if err != nil {
		w.logger().Error("activity", "op", "report", "report", reportID, "err", err)
		return
	}
	if !ok || f.bosses == 0 {
		return
	}
	if _, err := w.Store.Append(ctx, uploadEvent(f)); err != nil {
		w.logger().Error("activity", "op", "report upload", "report", reportID, "err", err)
	}
	if w.RaidZones[f.zone] {
		w.firstKills(ctx, f)
	}
}

func (w *ReportWriter) facts(ctx context.Context, reportID string) (reportFacts, bool, error) {
	var f reportFacts
	var classSpec string
	err := w.Store.Pool.QueryRow(ctx, reportFactsSQL, reportID).Scan(&f.id, &f.zone, &f.key, &f.name, &f.class,
		&f.level, &f.anonymous, &f.bosses, &f.durationMS, &classSpec, &f.at)
	if isNoRows(err) {
		return f, false, nil
	}
	if err != nil {
		return f, false, fmt.Errorf("activity: report facts %s: %w", reportID, err)
	}
	if class, spec, found := cutPipe(classSpec); found {
		if key, ok := SpecKey(class, spec); ok {
			f.spec = key
		}
	}
	return f, true, nil
}

func uploadEvent(f reportFacts) Event {
	length := textfmt.Duration(time.Duration(f.durationMS) * time.Millisecond)
	counts := fmt.Sprintf("%d %s in %s", f.bosses, plural(f.bosses, "boss", "bosses"), length)
	segments := []Segment{Text(f.zone + " uploaded, " + counts)}
	crest := ""
	if !f.anonymous && f.name != "" {
		crest = lowerASCII(f.class)
		segments = []Segment{Name(f.name, crest), Text(" uploaded " + f.zone + ", " + counts)}
	}
	e := Event{
		Kind: KindLogUpload, OccurredAt: f.at, Crest: crest,
		Href: "/reports/" + f.id, Segments: segments, DedupeKey: "upload:" + f.id, ReportID: f.id,
	}
	if !f.anonymous {
		e.CharacterKey = f.key
	}
	if f.spec != "" && f.level > 0 {
		e.Spec, e.Band = f.spec, BandOf(f.level)
	}
	return e
}

// firstKills writes one row for each raid boss this report killed that no earlier report
// killed first: the dedupe key per encounter makes the first writer the only one.
func (w *ReportWriter) firstKills(ctx context.Context, f reportFacts) {
	rows, err := w.Store.Pool.Query(ctx,
		`select distinct on (encounter_id) encounter_id, name from fights
		 where report_id = $1 and kill and encounter_id is not null order by encounter_id, start_ms`, f.id)
	if err != nil {
		w.logger().Error("activity", "op", "first kills", "report", f.id, "err", err)
		return
	}
	type kill struct {
		id   int
		boss string
	}
	var kills []kill
	for rows.Next() {
		var k kill
		if err := rows.Scan(&k.id, &k.boss); err != nil {
			rows.Close()
			w.logger().Error("activity", "op", "first kills scan", "report", f.id, "err", err)
			return
		}
		kills = append(kills, k)
	}
	rows.Close()
	for _, k := range kills {
		if _, err := w.Store.Append(ctx, killEvent(f, k.id, k.boss)); err != nil {
			w.logger().Error("activity", "op", "first kill", "report", f.id, "encounter", k.id, "err", err)
		}
	}
}

func killEvent(f reportFacts, encounterID int, boss string) Event {
	segments := []Segment{Strong(boss), Text(" fell in " + f.zone + ", the first kill on Forever")}
	crest := ""
	if !f.anonymous && f.name != "" {
		crest = lowerASCII(f.class)
		segments = []Segment{Name(f.name, crest), Text(" downed "), Strong(boss), Text(" in " + f.zone + ", the first kill on Forever")}
	}
	e := Event{
		Kind: KindBossKill, OccurredAt: f.at, Crest: crest, Href: "/reports/" + f.id,
		Segments: segments, DedupeKey: "first-kill:" + strconv.Itoa(encounterID), ReportID: f.id,
	}
	if f.spec != "" && f.level > 0 {
		e.Spec, e.Band = f.spec, BandOf(f.level)
	}
	return e
}

func (w *ReportWriter) logger() *slog.Logger {
	if w.Log != nil {
		return w.Log
	}
	return slog.Default()
}

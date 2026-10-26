# Guild Attendance and the Raid-Night Report (Lane G2) Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** A completed log attaches to the raid it was for, attendance is recorded per raider from the sign-ups and the log, and the raid page carries the night's report; the roster, the loot council and the readiness board read from it.

**Architecture:** Three prerequisites the survey found come first: reports learn their guild (attribution at completion from the logging character's verified membership) and their zone (persisted from the parse), and one `ReportCompleted` hook fires from both completion paths. Migration 0052 adds the match columns, the rejected-log table, `guild_raid_attendance`, `loot_awards.raid_id` and a per-raid unique index. The match store auto-matches on completion and takes hand matches; attendance rows are computed from `fight_metrics` through the `presence` rules and recomputed on every match change, with `excused` kept. The raid view gains `match`, `night`, `report`, `attendance`; four officer routes and the `report/post` route; the readiness read gains the `raid=` filter. The web adds the match line and picker, `Your night`, `Report`, `Attendance`, the member record and the readiness toggle. Two PRs: G2a (API, Tasks 1–5) and G2b (web, Tasks 6–10).

**Tech Stack:** Go (pgx, migrations), Astro + Svelte 5, vitest, Playwright.

**Spec:** `design/specs/2026-10-27-guild-attendance.md` (experience; its "Contract this spec assumes" is the API's contract) and `docs/superpowers/specs/2026-10-26-guild-suite-design.md` Lane G2. Survey of the code as it stands: the controller's `g2-anchors.md` (its findings are restated where a task needs them).

## Global Constraints

- Attendance evidence follows `api/internal/rating/presence` (`EvidenceReport`, `In`): private or flagged reports count nobody; `fight_metrics` rows exist only for encounter fights, so `pulls_total` counts encounter fights.
- Times: a log's anchor is `reports.created_at` (upload path) or `completed_at` (live path); "first seen" is log-relative (`start_ms` of the first fight the key appears in, minus the first fight's `start_ms`), never a wall-clock.
- One log matches one raid (partial unique index on `guild_raids.report_id`); a cleared match is never auto-picked again (`guild_raid_rejected_reports`).
- `Attendance{present, nights}` keeps its name and shape everywhere; in raid mode `nights` is per character and every consumer reads its own row's nights.
- The word "no-show" never appears in any player-facing string; `outcome: "no_show"` is a JSON value only.
- Officer writes (match, clear, excuse, post) sit behind rank → `CanForGuild(officer_views)` → claim freeze, rate-limited with the raid routes' limiter; copy and reading are free; a member sees only their own rows.
- Consumables: the parser stores no buff events today; `consumables.log_has_buffs` is `false` and `before` comes from the readiness board; no fake data.
- Copy verbatim from the spec through the spec-table tests; commit trailer `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`; API tests `-p 1`.

## Review Focus

1. A report completed by a member who belongs to no guild, or whose logging character is unverified: no attribution, no match, no error (Task 1).
2. Two raids of the same zone eight hours apart and one log: the window picks one and only one; the other stays `Ended · no log yet` (Task 2).
3. A member signed `in` on Zulmara who came on Duskstrider (same account): `present` with `came_as`; a member signed `in` who is in no fight: `no_show`; an unsigned raider in the log: `present` with `signed: null` (Task 2).
4. Re-matching a raid to another log recomputes every row but keeps `excused` where the character still appears (Task 2).
5. The loot council tiebreak after the switch: two candidates with the same `present` but different per-character `nights` sort by ratio, not by raw present (Task 3).

---

## PR G2a — the API

### Task 1: Report attribution, the persisted zone, and the completion hook

**Files:**
- Create: `api/internal/db/migrations/0052_guild_attendance.up.sql` / `.down.sql` (the whole lane's schema, written here, used by Tasks 2–5), `api/internal/reports/attribution.go` (+test), `api/internal/reports/completed.go` (the hook interface)
- Modify: `api/internal/parse/parse.go` (`Deps.Completed reports.Completed`, called after `announce`, best effort; `writeFight`/`writeReport` persist `log_zone` from the first encounter fight's `Fight.Zone` or the engine header), `api/internal/reports/ingest.go` (`complete` ~621–650: attribution, zone, the hook), `api/internal/reports/handler.go` (the PATCH that sets `guild_id` also fires the hook), `api/internal/reports/store.go` (`SetLogZone`, `SetGuildIfUnset`), `api/cmd/api/main.go` + `cmd/api/notifyservice.go`-style wiring for the parse job (a guilds match store on the job's pool)
- Test: `api/internal/parse/parse_test.go` (the fixture upload completes with `log_zone` set and the hook called), `api/internal/reports/ingest_test.go`, `attribution_test.go`

**Interfaces:**
- Produces: migration 0052 — `reports.log_zone text` (nullable); `guild_raids` + `pre_done_status text`, `matched_by text check in ('auto','hand')`, `matched_at timestamptz`, `report_posted_at timestamptz`, `create unique index guild_raids_report_idx on guild_raids(report_id) where report_id is not null`; `guild_raid_rejected_reports(raid_id fk cascade, report_id text fk cascade, rejected_by, rejected_at, pk(raid_id, report_id))`; `guild_raid_attendance(raid_id fk cascade, character_key text, name, class, spec, signed text null check in (in,late,bench,out), role text null, outcome text check in (present,late,absent,no_show), came_as text null, first_seen_offset_s int null, pulls_present int, pulls_total int, excused boolean default false, computed_at, pk(raid_id, character_key))`; `loot_awards.raid_id bigint null fk guild_raids set null`; drop `loot_awards_drop_idx`, create `loot_awards_drop_idx unique (guild_id, raid_id, encounter_id, item_id)` (NULL raid_id rows stay distinct per Postgres — the pre-raid helper keeps its upsert on `raid_id is null` through a second partial unique index `loot_awards_drop_unmatched_idx unique (guild_id, encounter_id, item_id) where raid_id is null`).
- Produces: `type Completed interface{ ReportCompleted(ctx, reportID string) }` (nil-safe), `func (s *Store) AttributeGuild(ctx, reportID string) (guildID *int64, err error)` — the logging character's verified `guild_characters` row's guild, else the owner's only verified guild, else nil; writes `reports.guild_id` only when unset.

- [ ] Steps: failing tests (attribution three cases; the fixture parse sets `log_zone` = the fixture's zone; the live `complete` sets guild + zone and calls the hook; PATCH fires the hook) → red → implement → `go test -race -p 1 ./internal/reports/ ./internal/parse/ ./internal/db/` → commit `feat(api): reports learn their guild and zone at completion; the ReportCompleted hook; migration 0052`.

### Task 2: The match store and attendance rows

**Files:**
- Create: `api/internal/guilds/raids_match.go` (store: `AutoMatch(ctx, reportID)`, `SetMatch(ctx, guildID, raidID, reportID, by, how)`, `ClearMatch(ctx, guildID, raidID, by)`, `MatchCandidates(ctx, guildID, raidID, now) ([]Candidate, error)`), `api/internal/guilds/raids_attendance.go` (`RecomputeAttendance(ctx, raidID)`, `SetExcused(ctx, raidID, key, bool)`, `AttendanceRows(ctx, raidID)`), `api/internal/guilds/report_completed.go` (`func (s *Service) ReportCompleted(ctx, reportID)` → `AutoMatch` → `RecomputeAttendance`; best effort, logged)
- Test: `raids_match_test.go`, `raids_attendance_test.go` with a new multi-fight, multi-player night seeder (`seedNight(t, pool, guildID, reportID, zone, createdAt, fights []seedFight)` in `raids_fixture_test.go`; call `db.EnsureMetricsPartition` first)

**Interfaces:**
- Produces: `Candidate{ReportID, ZoneName string; UploadedAt time.Time; UploadedByName string; Pulls, Kills int; DurationMS int64; Closest bool; MatchedRaid *RaidRef}`; `AttendanceRow{CharacterKey, Name, Class, Spec string; Signed, Role *string; Outcome string; CameAs *string; FirstSeenOffsetS *int; PullsPresent, PullsTotal int; Excused bool}`; `ErrReportMatched`, `ErrRaidNotEnded`.
- Rules: `AutoMatch` picks among the guild's raids with status `planned|locked`, `report_id null`, zone's `log_zone` = `reports.log_zone`, whose window (`starts_at − 1 h`, `end + 6 h`) contains the log's anchor, not in the rejected table, the nearest `starts_at` on a tie; a match sets `report_id`, `pre_done_status`, `status = done`, `matched_by`, `matched_at` under the status guard; `ClearMatch` restores `pre_done_status`, inserts the rejected row, clears the attendance rows; `MatchCandidates` lists the guild's evidence reports of that zone from the last 14 days (`closest` = nearest anchor; `matched_raid` filled from the unique index). Attendance per Review Focus 3–4: participants from `fight_metrics` (state ok) via `presence`, keys mapped to the guild's roster by character key, `came_as` by the signup character's owner account, `late` = first seen > first pull + 15 min for `in`/`late` signups, `absent` for bench/out, `no_show` for signed `in`/`late` and absent from the log.

- [ ] Steps: failing tests for Review Focus 2–4 plus the candidate list and the clear/rejected rule → red → implement (small functions, SQL in the store files, ≤ 400 lines each) → PASS → commit `feat(api): raids match their logs; attendance rows computed from the sign-ups and the log`.

### Task 3: The attendance switch, the member record, the loot tiebreak

**Files:**
- Modify: `api/internal/guilds/roster_extended.go` (`enrichRoster`: when the guild has ≥ 1 `done` raid with attendance rows, `Attendance` per character = present over raids the character signed in/late or attended, last 8; `attendance_30` the same over 30; `attendance_source: "raids"`; else today's report count with `"reports"`), `home.go` (`HomeView.AttendanceSource`, `RosterRow.Attendance30`), `loot.go` (~182–189, ~267, ~428: each candidate's own nights; the tiebreak by ratio then present), `members.go`, `api/internal/auth/store.go` + `handler.go` (`CharacterGuild.Attendance *Attendance` through an injected `GuildAttendanceReader` interface, nil-safe), `cmd/api/main.go`
- Test: `roster_extended_test.go`, `loot_test.go` (Review Focus 5), `auth/handler_test.go`

- [ ] Steps: red → green → commit `feat(api): attendance from raids once a guild has one; per-character nights on the roster, /me and the loot council`.

### Task 4: Raid view, match routes, attendance route

**Files:**
- Modify: `api/internal/guilds/raids_view.go` (`RaidView` gains `PreDoneStatus *string`, `Match *MatchView`, `Night *NightView`, `Attendance []AttendanceView` (officers all, members own; `null` when none)), `raids_handler.go` (routes: `GET …/raids/{raid}/match-candidates` officer; `PUT …/raids/{raid}/match` `{report_id}` → 409 `report_matched` / `conflict` (raid not ended) / 422; `DELETE …/match`; `PUT …/raids/{raid}/attendance/{region}/{ruleset}/{name}` `{excused}`; all behind the officer gate + the raid limiter), `raids.go` (`RaidRow.RaidID *int64` via a left join on `guild_raids.report_id` so a night links to its raid), `handler.go` routes
- Test: `raids_match_handler_test.go`, `raids_handler_test.go` (view keys per the spec's contract; a member's `attendance` holds only their rows)

- [ ] Steps: red → green → commit `feat(api): the raid view's match, night and attendance; match, clear and excuse routes; nights link to their raids`.

### Task 5: The report view, posting, readiness filter

**Files:**
- Create: `api/internal/guilds/raids_report.go` (`ReportView` per the contract: `kill_order` from `fights` (kill, `kill_ms`, `wipes_before` = earlier non-kill fights of the same encounter, `deaths`), `unnamed` (encounter_id null), `deaths` per raider split on `fight_metrics.kill`, `consumables{log_has_buffs:false, rows}` with `before` from the readiness board's bags rule, `loot` = `loot_awards` where `raid_id` = this raid grouped by encounter with item name/quality from the loot table, `moves` = top five |Δ| of `metric_dps`/`metric_hps` on kills vs the previous `done` raid of the same zone, `posted_at`), `api/internal/notify/report_embed.go` (`ReportReadyEmbed`, colour 8377482, the §6.2 fields) + `GuildDeliverer.ReportReady(...) (posted bool, err error)`
- Modify: `raids_handler.go` (`POST …/raids/{raid}/report/post` → conditional `update … set report_posted_at = now() where report_posted_at is null` first, 409 `already_posted`; the embed on the 8 s budget), `raids_discord.go` (`RaidEvents.ReportReady`), `loot.go` (`CreateLootAward` body gains optional `raid_id`; the helper's default = the guild's most recent `done` raid of that zone within 7 days, else null), `readiness.go` + `handler.go` (`GET …/readiness?raid={id}` → rows limited to the raid's `in`/`late` sign-ups incl. `no_export` rows built from the sign-up snapshot, `raid{id,title,starts_at,cutoff_at}`, per-row `synced_before_cutoff` (export `logged_at` within `READINESS_FRESH_DAYS = 7` before `starts_at`, else a failing check); the unfiltered read adds `next_raid`)
- Test: `raids_report_test.go` (goldens over the seeded night), `report_embed_test.go`, `readiness_test.go` (the filter and the stale rule), `loot_test.go` (`raid_id` on an award; the per-raid unique index keeps last week's award)

- [ ] Steps: red → green → commit `feat(api): the raid-night report, report_ready posting, loot awards per raid, the readiness next-raid filter`.

**PR G2a gate:** `go test -race -p 1 ./...` → CI → fast-forward → deploy api (0052) → verify `schema_migrations` 52 and that an anonymous `GET /v1/guilds/1/raids/1/match-candidates` is 404 → memory.

---

## PR G2b — the web

### Task 6: Contracts, copy, API calls

**Files:** `web/src/lib/guild/api.ts` (the contract's types), `web/src/lib/guild/attendance-copy.ts` (+ spec-table test over `### 6.1`), `web/src/lib/guild/attendance-api.ts` (`fetchMatchCandidates`, `setMatch`, `clearMatch`, `setExcused`, `postReport`, `fetchGuildReadiness(guildId, raidId?)`), `web/src/lib/guild/night.ts` (pure: group the attendance rows into the §3.5 five groups in order; the `Your night` sentence; the copied Discord text ≤ 12 lines), tests.

- [ ] Steps: red → green → commit `feat(web): attendance contracts, copy, grouping and API calls`.

### Task 7: The done band, match line, picker, clear confirm

**Files:** `web/src/components/GuildRaid.svelte` (band lines per §3.1 on `done`/ended; the sr-only live region when the poll flips to done), `guild/RaidMatchLine.svelte`, `guild/RaidMatchPicker.svelte` (officers; `Closest` preselected; a matched-elsewhere row disabled with its sentence; empty and unknown-zone states), `guild/RaidClearConfirm.svelte` (focus `Keep it`, Escape keeps), the controller (`raid-page-controller.svelte.ts`: match/clear actions through the officer-action pattern; `guild_plan_required` flip), `GuildRaids.svelte` (a night row links to its raid via `raid_id`; `#raids-{report_id}` deep link opens that night), tests.

- [ ] Steps: red → green → commit `feat(web): the matched-log band, the match picker and clear confirm, nights linked to raids`.

### Task 8: Your night, Report, Attendance

**Files:** `guild/RaidYourNight.svelte` (in the strip's slot on `done`), `guild/RaidReport.svelte` (+ `RaidKillOrder.svelte`, `RaidDeaths.svelte`, `RaidLoot.svelte` with `ItemHover`, `RaidMoves.svelte`, `RaidConsumables.svelte`; two columns from 1024; `Copy for Discord`; `Post to Discord` once, 409 handled), `guild/RaidAttendance.svelte` (five groups, `Excused` on the first two, compact rows for Present and Bench/out; officer-read-only), mounted in `GuildRaid.svelte` before Notes; every file ≤ 400 lines; tests for each component's states.

- [ ] Steps: red → green → commit `feat(web): Your night, the raid-night report and the attendance groups on the raid page`.

### Task 9: The member record and the readiness toggle

**Files:** `guild/GuildRosterTable.svelte` + `lib/guild/roster-facts.ts`/`nights.ts` (raid mode chip `Came 7 of 8 · 26 of 30` + the one footnote; report mode unchanged), `lib/guild/loot-view.ts` (`7 of 8 raids`), `components/character/CharacterGuildLine.svelte` (` · came 7 of 8 raids`), `guild/GuildReadiness.svelte` + `Guild.svelte` (the `Next raid · Thu 10 Dec` toggle on by default; stale and no-export rows), tests.

- [ ] Steps: red → green → commit `feat(web): attendance on the roster, loot line and character guild line; the readiness next-raid toggle`.

### Task 10: e2e, captures, reviews, changelog

**Files:** `web/tests/e2e/guild-attendance.spec.ts` + `tests/e2e/support/guild-attendance.ts` (stubs for the done raid, candidates, match/clear/excuse/post, readiness with `raid=`), `web/tests/captures/guild-attendance/` (port 4451), `design/reviews/2026-10-27-guild-attendance/README.md`, fixtures (`mock-guild.ts`: a done raid, a report, attendance rows, candidates), `web/src/content/changelog/2026-10-2x-guild-attendance.md`.

- [ ] Steps: the spec §9 e2e list → run (port 4452) → captures rows 1–11, 13 → ux-designer + wow-player reviews → one fix round → commit `test(web): guild attendance e2e, captures, reviews and changelog`.

**PR G2b gate:** full web vitest + full Playwright + lhci → CI → fast-forward → deploy web → verify → memory.

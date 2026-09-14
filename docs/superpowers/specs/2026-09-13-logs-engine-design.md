# Forever Sixty — combat logs design (engine and report experience)

Date: 2026-09-13. Status: approved in conversation, section by section. Replaces the earlier
engine-only draft of the same date.

This is the end-state design for the logs product: one engine, one storage story, real-time
from day one, sized for Classic 2019 launch scale, and cost-efficient enough to run on storage
alone. The engine starts now in its own worktree; the report experience is Phase 3 (Nov 4 to
Dec 9, 2026), targeting the first raids on Dec 9. The bar the user set: match or beat Warcraft
Logs' report experience on day one, since nothing less gets used.

## Decisions

| Question | Decision |
|---|---|
| Scope | Full fidelity: every event type the format defines, every analysis a raid leader opens, not a damage meter, not an MVP. |
| Scale target | Classic 2019 launch: order of 100,000 reports a week at peak, thousands of live raids at once, years of retention. |
| Parse location | In the desktop companion, verified by the server; whole-file web uploads parse server-side with the same engine. |
| Read path | Files on the edge: precomputed summaries as JSON and per-fight events as Parquet on R2 behind Cloudflare's cache; deeper queries run in the browser with an analytical engine compiled to WebAssembly. No server compute per page view. |
| Object storage | Cloudflare R2: no egress fees, S3-compatible, behind the same cache as the site. |
| Rankings | Updated live at fight close, from per-player metrics rows in Postgres with incrementally maintained percentile digests. |
| Retention | Keep everything; raw uploads tier to cold storage after 90 days. Account deletion never deletes reports or rankings. |
| Companion | Native, signed desktop app (Go), macOS and Windows; the browser-based alternative was considered and rejected for seamlessness. |
| Engine location | New top-level `logs/` Go module: library, CLI, tests, WebAssembly not required. |
| Format strategy | Version-keyed decoder layouts so Forever's format, retail-style or Classic-style, is one table row and one roster decoder when the first beta log arrives. |
| Fights | Encounters name themselves from `ENCOUNTER_START`; everything else is "Trash" with a zone label. No encounter tables, phase markers, or per-encounter expectation lists. |
| Advanced logging | Assumed on. Reports without it are accepted, marked, and lose the views that need it. |
| Cross-report analytics | Batch over the lake; ClickHouse self-hosted on one VM if and when an interactive cross-report query is needed, ingesting the same Parquet. |

## 1. Architecture and performance budgets

```
companion (Go, on the player's PC) ── live fights, verified ──▶ ingest (Go, Cloud Run) ──▶ R2
web upload (whole file) ── resumable to R2 via signed URL ──▶ parse job (Cloud Run job) ──▶ R2
                                                              │
                                                              ├─▶ fight metrics ─▶ rankings store (Postgres) ─▶ percentile digests at fight close
                                                              └─▶ report index (Postgres): reports, fights, players, guilds, ownership
report page (static shell + report island) ◀── Cloudflare cache ◀── R2: report.json, fight summaries, fight Parquet
                                            deep queries: DuckDB-WASM in the browser over the fight's Parquet
rankings, character, guild pages ◀── API ◀── Postgres
```

Budgets, measured in CI on the fixture set:

| Path | Budget |
|---|---|
| Live fight visible after it ends | under 3 s from the `ENCOUNTER_END` line being written to the summary being served |
| Whole-file upload, 500 MB raid night | first fight visible within 10 s of upload completion; whole file parsed in under 90 s on two CPUs; fights appear progressively |
| Report page first paint | summaries under 200 KB per fight from the edge cache; LCP within the site's 1.6 s mobile budget |
| Any deep query on a fight | under 500 ms on a laptop after the fight's Parquet (2 to 10 MB compressed) is loaded once |
| Rankings page for an encounter | under 200 ms from Postgres |
| Parse throughput | at least 10 MB per second per CPU; memory bounded by the open fight, never a whole-file load |

The expensive scan is per fight, never per report or across reports. Rankings never touch
events. The edge serves every repeat view of a closed fight without an origin call. Cross-report
ad-hoc queries over raw events, which Warcraft Logs answers from a memory tier, run here as batch
jobs; no report or rankings page needs them.

## 2. Ingest and the engine session

**One engine, two hosts.** `logs/engine` exposes a session: `NewSession(reportID, opts)`,
`Feed(chunk, offset)` returning decoded events and closed fights, `Snapshot()` for the open
fight's running summary, `State()` and `Restore()` for resumption, `Close()`. Batch parsing is one
session fed one file. The companion, the ingest job, and the CLI call the same code.

**Live path.** The companion tails `WoWCombatLog.txt`, runs the session, and on each fight close
sends one bundle: the fight's events as Parquet, its summary, its metrics rows, and a hash of the
raw lines the fight came from. Every few seconds during a fight it sends the snapshot. The raw
file streams up in the background as compressed chunks addressed by offset, resumable.

**Whole-file path.** The browser uploads in resumable chunks straight to R2 through a signed URL;
a Cloud Run job runs the session over the file and writes fights as they close, so the report
fills in progressively.

**Verification.** For every companion bundle the ingest recomputes the metrics from the uploaded
events and rejects the bundle on mismatch. After the raw file arrives, a background job re-parses
a sample of fights from raw bytes and compares; a mismatch flags the report and excludes it from
rankings until reprocessed server-side. The engine version is recorded per fight.

**Idempotency and resumption.** Chunks and bundles are addressed by report id, fight index, and
byte offset; re-sent data is recognized and ignored, so at-least-once delivery is safe. A
companion that disconnects resumes from its own state and offset; the server reconciles by
offset.

**Storage layout on R2**
```
reports/<id>/report.json                    metadata, health, fight list, units, engine version   (5 s cache)
reports/<id>/fights/<n>/summary.json        precomputed tables                                     (immutable, 1 y)
reports/<id>/fights/<n>/events.parquet      typed events, queried in the browser                   (immutable, 1 y)
reports/<id>/fights/<n>/live.json           snapshot while the fight is open                       (5 s cache)
reports/<id>/raw/<offset>.zst               raw chunks; tiered to cold storage after 90 days
```

**Report index in Postgres.** Reports (owner, guild, zone, visibility, engine version), fights
(encounter, difficulty, kill, duration, players), players seen, per-fight metrics rows, percentile
digests, moderation states. Events never enter Postgres.

**Cost.** Per whole-file upload about half a cent of compute; live bundles under a tenth of a
cent. Storage about 120 MB per report (100 MB compressed raw plus 20 MB Parquet and summaries),
roughly $0.015 per GB-month on R2, halved by cold tiering. At 10,000 reports a week: about
$65 a month compute, 5 TB a month of storage growth, about $900 a month of storage by month 12.

## 3. Data model and analyses

**Units.** GUIDs parsed into kind, id, spawn. Players keyed by GUID, name as display data. Class
from `COMBATANT_INFO` when present, otherwise inferred from spells cast and labeled inferred. Pet
owners from summon events and the advanced owner field.

**Fights.** `ENCOUNTER_START` to `ENCOUNTER_END` gives id, name, difficulty, size, kill or wipe.
Everything else is segmented by hostile-combat gaps and labeled "Trash" with the zone; dungeons
the same, kills counted. Each fight records players present, duration, and raid markers.

**Analyses, precomputed per fight at close into `summary.json`**
- Damage done, taken, and taken by ability: per source, target, ability; hits, crits, misses by
  type, blocked, absorbed, resisted, overkill; one-second series.
- Healing done and taken: amount, overheal, absorbs granted, crits; one-second series.
- Deaths: killing blow, last ten damage events with health from advanced fields, auras applied
  and removed in the final seconds, defensives available, time to release.
- Buffs and debuffs: application segments, uptime, stacks over time, applier; raid buffs at pull
  against the roster.
- Casts: count, success versus failed, cast time from start-success pairs, sequence.
- Interrupts and dispels: who, what, on whom; casts that got through.
- Resources: mana, rage, energy series; time at zero.
- Threat: computed from the vanilla threat model per class and ability.
- Combatant info: gear with item ids, enchants, item level; talents; consumables from auras at pull.
- Roster: active time, activity percentage, deaths, and the ranking metrics.

**Ranking metrics per player per boss kill.** Role metric (DPS, HPS, damage taken for tanks),
active time, item level, spec from the talent split, fight duration, encounter id, difficulty,
size, date.

**Events file.** One Parquet per fight, sorted by time then line, dictionary-encoded strings,
typed nullable columns per suffix, the 17 advanced fields as nullable columns, and `raw` for
unknown or errored events.

## 4. The report experience

Same shape as Warcraft Logs so nobody relearns: fight selector (encounter, kill or wipe,
duration, time, phase presets), then modes (Analyze, Compare, Rankings; Mechanics and Replay
later), then views (Tables, Timelines, Events, Queries), then source scope (all friendlies, all
enemies, one player), then the twelve table tabs: Summary, Damage Done, Damage Taken, Healing,
Threat, Buffs, Debuffs, Deaths, Interrupts, Dispels, Resources, Casts. Filters beside the chart:
target, ability, pets, NPCs, weighted damage, overkill, ignore events after player deaths, query
expression. Rows: parse percentile, class-colored name, amount bar segmented by ability, active
time, per-second figure, expander. URL state carries fight, mode, view, tab, source, and window.

**The time window is the primary interaction.** A brushable per-second chart above every table
and timeline; brushing rescopes everything instantly in the browser from the fight's events;
phases are preset windows; clicking a death sets the window to the seconds before it.

**Instant.** The summary opens every table without a query; the events file loads once per fight
and serves every brush, filter, and drill-down from then on. DuckDB-WASM loads lazily on the first
deep interaction, never on page load.

**Where we go past them, by design**
- No ads; the data gets the whole viewport.
- Live within three seconds; the selector marks a fight in progress.
- Everything they gate is free: archive, cross-report comparison, priority processing, queries.
- Deaths and gear link to the player's build: talents and gear at pull, one click into the planner.
- Trash is trash, with a zone label; the report's timeline shows the night with pull counts and
  wipe progression per boss.
- Mobile first: tables collapse to cards at phone width with the chart above.

**Deferred, not dropped:** Replay and positional views once Forever logs prove the positional
fields; Mechanics content per encounter over time.

## 5. The companion

Native Go app, one static binary per platform, macOS and Windows signed and notarized, Linux
unsigned. Menu-bar or tray app with a small native-webview window in the site's design system:
status, current report link, recent reports, sign-in state, upload log.

- Watches the game's `Logs` folder, detects a growing `WoWCombatLog.txt`, runs the engine on it,
  sends fight bundles and snapshots, streams raw chunks in the background, resumable.
- Checks that advanced combat logging is enabled and shows the exact menu path if not.
- Syncs the addon: reads `ForeverSixty.lua` under SavedVariables after logout and uploads the
  character exports; writes the addon's inbox file so builds chosen on the site appear in-game.
- Reports which character is logging for attribution.
- Signs in once via the site: a pairing code issues a per-device upload token, revocable from the
  account page. Never sees passwords or Battle.net credentials.
- Auto-updates from GitHub Releases with signature verification; engine version travels with
  every bundle.
- Offline: fights queue locally and upload in order on reconnect. Engine failure on a fight: raw
  chunks still upload and the server parses that fight itself, reporting the failure with the
  engine version. Never modifies the game's log file; can archive and rotate old logs by setting.
- Prerequisites the user owns: Apple Developer account for notarization; Windows signing via
  Microsoft Trusted Signing (about $10 a month) or SignPath's free program if open source.

## 6. Rankings

Modeled on Warcraft Logs' zone rankings, read directly:

- **Types:** character Damage, Healing, Damage to Bosses; guild Progress, Speed, Execution;
  All Stars per raid as a points total across bosses.
- **Brackets:** content phase, with item level shown per row and as a filter. Phase boundaries
  come from the site's dates data.
- **Specs:** derived from the talent split at pull with a per-class table naming the game's specs
  and the hybrid builds the community recognizes; the same table names builds in the planner.
- **Columns:** rank, name with guild, realm, region; metric; size; date; duration; talent split;
  trinkets; buff count; report flag; optional video link. The talent split links into the planner
  with that tree and gear loaded.
- **Filters:** boss, region, realm, faction, class and spec, phase, item level, date; all in the URL.
- **Moderation:** at-risk and removed states inline with reasons; per rank, per report, or per
  hotfix rule.
- **Live:** rows written at fight close after verification; percentiles from per-group digests
  updated incrementally; the report shows the percentile seconds after the kill; during a fight a
  projection from the snapshot, labeled as such. Today versus all-time views from the same rows.
- **Character and guild pages:** every ranked fight over time, best per encounter, builds seen at
  pull; guild progression per boss, pull counts, kill dates, roster bests.
- **Integrity:** ingest recomputation, raw-sample checks, owner hide, moderator flag.
- **Store:** Postgres partitioned by month, indexed on encounter, spec, bracket, metric.

## 7. Accounts, ownership, visibility, privacy

- **Sign-in:** Battle.net first (client registered), email magic link as fallback; roles user,
  moderator, admin; opaque session cookies on `.foreversixty.gg`, HttpOnly, Secure, SameSite
  Lax, stored in Postgres; CSRF by double-submit; no passwords anywhere.
- **Ownership:** the uploader owns a report, with optional guild and logging character.
  Anonymous whole-file uploads are public, unowned, claimable for 24 hours.
- **Visibility:** public, unlisted, private (no rankings), guild (guild page only, ranks).
- **Guilds:** claimed by an officer per the Battle.net roster; officers and members refreshed on
  sign-in; the game's roster is the truth.
- **Characters:** linked from the Battle.net profile; unclaimed characters still get pages from
  public reports, claimable by signing in.
- **Account deletion:** removes the account, sign-in links, email, tokens, settings, and
  character links. Reports stay: ownership passes to the guild or the report becomes unowned;
  private no-guild reports become unlisted. Rankings rows and character pages stay. Stored events
  are never rewritten.
- **Names:** hide-on-claim removes a character from search and rankings; a read-time pseudonym is
  available on request to anyone, account or not, applied across public reports and rankings.
- **Legal footing:** legitimate interest, plus the position Warcraft Logs' policy states, that
  characters, names, and character data are owned by the game company under Blizzard's User
  Agreement and are not personal information. Launch gates: a privacy policy naming the
  controller, basis, retention, processors, and request channel; processor terms accepted with
  Cloudflare, Google Cloud, Neon, and Resend; a records-of-processing note; a lawyer's review.
- **Minimization:** chat and emote text dropped at parse; no analytics cookies; tokens hashed;
  raw files private and served only to their owner as a download.
- **Abuse:** rate limits per token and account, a size ceiling, moderator hide, flag, and ban.

## 8. Infrastructure and operations

| Component | Runs on | Scaling |
|---|---|---|
| Ingest API | Cloud Run service, Go, existing project | 0 to N, request-driven |
| Whole-file parse | Cloud Run jobs, two CPUs | one per upload |
| Object storage | Cloudflare R2, one bucket, lifecycle rule on `raw/` | unbounded |
| Index and rankings | Neon Postgres, new schema, monthly partitions | vertical |
| Report delivery | Cloudflare CDN in front of R2 | edge |
| Report page | static site plus report island; DuckDB-WASM lazy | edge |
| Rankings and pages | existing Go API | as today |

Live path end to end: bundle to ingest, verify, write R2 and Postgres, purge `report.json` at the
edge, five-second page poll. Caching: only `report.json` and `live.json` are mutable; rankings
pages cache 30 s. Observability: structured logs with report id and engine version, per-report
health in `report.json`, metrics for ingest latency, verification failures, job durations, bucket
growth, alerts on budget breaches; Cloud Logging and Monitoring. Backups: Neon point-in-time
recovery; R2 object versioning on `report.json` and `fights/`; raw files make every derived file
reproducible by a reprocess job keyed on report range and engine version. Security: per-device
hashed tokens scoped to upload; hour-long signed upload URLs bound to a report id; bundle sizes
validated against raw byte ranges; private bucket with signed paths for non-public reports.

## 9. Testing and gates

- Engine: fixture line per event shape and layout; golden files byte for byte; streaming
  equivalence at chunk sizes 1, 7, 64, 4096 and random, with duplicate and overlapping ranges
  ignored and mid-fight serialize-restore; property tests (parse, write, read back; recompute
  summaries and metrics); health fixtures for malformed lines, truncation, missing header, year
  rollover, clock jumps.
- Performance: benchmark asserting throughput and memory in CI; island timing test for a brush
  over a 10 MB fight; Lighthouse on a fixture report page.
- Real logs: the retail sample fetched by script (never committed); the first beta log pins the
  Forever layout and a redacted excerpt becomes the canonical golden file; a conformance suite
  reports unknown events, parse errors, and inferred layouts across every collected log.
- Companion: tail, offsets, upload queue with a fake server; integration with a fake game writer,
  local ingest, drops and restarts; signed build smoke test per platform.
- Ingest and rankings: verification, idempotent re-sends, size limits, token scopes; digest
  versus brute-force percentiles; rankings queries on a seeded month.
- Report experience: Playwright flows for brush, tabs, drill-down, deaths, talent split into the
  planner, live updates from a stubbed ingest, phone layout.
- Launch gates: all green; a real raid night logged live and compared table by table with Warcraft
  Logs' report of the same night; privacy policy reviewed; processor terms accepted.

## 10. Sequencing and open questions

- **Now, `logs-engine` worktree:** the engine, CLI, fixtures, golden files, benchmark.
- **Sept 17 to 18:** first beta log pins Forever's layout; rework bounded to the layout row and
  the roster decoder.
- **Phase 3, Nov 4 to Dec 9:** ingest and whole-file upload first, report page second, companion
  and live third, rankings fourth; accounts, policy, processor terms; live for Dec 9.
- **After Dec 9:** Replay and positions; Mechanics content; ClickHouse when warranted; the
  simulator's log import.

Open on purpose: which logging format Forever writes (answered by the first beta log); whether
positions are present (decides Replay's timing); whether a raid-size live test is possible before
Nov 4 (else a volunteer guild on Dec 9); signing accounts (needed at Phase 3 start).

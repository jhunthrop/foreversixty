# seedguild

A reversible seed tool that populates **one real guild** on the database with a
realistic, clearly-tagged mock roster, raid nights and fight metrics — so the owner can
design and build the guild page UX against a full page instead of a single-member one.

Built for guild id 2, "OLYMPUS XXVII" (region `us`, ruleset `pvp`), owner battletag
`hunthrop#1894`, but takes both as flags and works against any guild the owner's account
already holds a (verified or unverified) character in.

## What it writes

- **The guild claim.** The owner's account claims guild `--guild`, with
  `officer_max_rank_index` set to 1. The guild's prior claim state is snapshotted first
  (`seed_rows`, see below) so `--remove` restores it exactly, including "still unclaimed."
- **The owner's own character(s).** Every `guild_characters` row the owner already holds
  in this guild is promoted to `rank = 'leader'`, `verified_by = 'claim'`,
  `verified_at = now()` — exactly what a real claim leaves behind. Snapshotted first, so
  `--remove` restores the owner's exact prior rank/verification, not a guess.
- **24 mock characters across 22 accounts**, all nine classes, a tank, six healers,
  seventeen DPS, two Skyborne, two officers, three left unverified (so the approve button
  on the roster has rows), and two accounts that each carry a main and an alt character
  (`roster.go`'s `Account` field) — every `users`/`characters`/`addon_exports`/
  `guild_characters`/`guild_members` row this implies.
- **Battle-tag tag.** Every mock account's battletag is `Seed#0001`..`Seed#0022`
  (22 accounts for 24 characters) — never a value a real Battle.net account can hold.
- **Attendance variety.** Each of the 24 mock characters attends 18–22 of the 4 raid
  nights per night (not all 24 on every night): two characters attend exactly one night
  each, and one character has zero attendance in the last 7 days — `fight_metrics` (and
  `fights.players`) carries a row only for raiders actually present at that fight.
- **Readiness variety in the FS1 exports** (`addon_exports.export`, built per
  `addon/ForeverSixty/Codec.lua`'s real section grammar — see "FS1 fidelity" below): most
  raiders are geared exactly to their spec's band-60 BiS pick
  (`data/builds/1.60.1.70009/bis/*.json`); four raiders have 2–4 slots swapped for that
  slot's own real BiS *alternative* item (never an invented id); two raiders are missing
  the enchant on one slot that every other raider's copy of that slot carries; two raiders
  have unspent talent points; every raider carries a `professions=` section (its class's
  own gathering/crafting pair plus first aid); the one `gear_bags`-consent raider carries
  a `bags=` section of real level-60 raid consumables.
- **4 raid nights across Forever's real first raid tier** — 2 Onyxia's Lair, 1 Barrow
  Deeps, 1 Hyjal Summit (`data/curated/loot/forever-raid-phases.json`'s own notes field:
  "the first tier opens on 9 December and is Barrow Deeps, Hyjal Summit and Onyxia";
  Molten Core is a later, unscheduled "Era" raid and is not used here) — 6–10 fights each
  with 1–2 wipes, owned by the owner's account, `visibility = 'guild'`,
  `status = 'complete'`. Onyxia is the only one of the three with a published encounter
  (`logs/engine/mechanics/encounters.json`: id 1084 "Onyxia", read from the Forever client
  itself) — both Onyxia nights end in a kill. Barrow Deeps and Hyjal Summit have no
  published encounter or boss name at all (that same file's own
  `no_client_rows_yet.raids` list); their pulls carry a null `encounter_id` and a generic
  "Pull N" name rather than an invented boss, matching the design spec's own ruling
  (`design/specs/2026-10-04-guild-page.md` §12.1: "no named encounter published").
- **`seed_rows` (migration `0029_seed_rows`).** One row per created/mutated row this tool
  is responsible for, keyed by a tag derived from the guild id
  (`seedguild-<guild id>`). A created row (`users`, `reports`) has `prior = null`;
  `--remove` deletes it outright. A mutated row (`guilds`, the owner's own
  `guild_characters` rows) has `prior` holding a JSON snapshot of its previous column
  values; `--remove` restores those columns instead of deleting anything. This is the one
  table added — no existing table or column was repurposed, because nothing already on
  `guilds`/`guild_characters`/`addon_exports` can tell a mock row from a real one, and a
  battletag-prefix heuristic alone cannot tell "mutated, restore me" from "created, delete
  me."

## Why `visibility = 'guild'`, never `'public'`

A `'public'` report is rankable and shows up in cross-guild leaderboards and the public
character rating trend — a mock raid night with made-up numbers has no business there.
`'guild'` keeps the four raid nights visible only to this guild's own page (home reports,
roster bests, progression), which is everything the owner needs to design against, and
keeps them out of every public-facing surface. (`reports.Ranked` — guild counts as ranked
for the guild's *own* progression/roster-best queries, which is also why those panels
populate even though the reports themselves are guild-only.)

## Dates: a launch-raid stand-in, never a future date

The real design spec this seed is built for
(`design/specs/2026-10-04-guild-page.md`, "Data notes") fixes its own mock boards' raid
nights after 9 December 2026 (when raids actually open) — but that date is in the future
relative to when this tool is actually being built and tested (Forever's beta runs now,
capped at level 30; raids are closed). **A `created_at` in the future is not an honest
database row.** This tool schedules its four raid nights relative to the moment it
*runs* — the last four raid-size nights before `now` (`raid.go`'s `DaysAgo`), two within
the trailing 7 days and two older — not the spec's own fixed 12-10/12-15/12-17/12-22
dates. Once this tool is actually run after 9 December 2026, "the last four raid-size
nights before now" lands after that date on its own, and the Barrow Deeps/Hyjal Summit/
Onyxia level-60 picture stops being a stand-in and starts being the literal production
launch-raid; until then, it is clearly labelled one (see the header comment in `gear.go`
and `raid.go`). Report titles name the *actual* weekday `created_at` falls on (`raid.go`'s
`titleFor`), never a hardcoded day name that could disagree with the real row.

## FS1 fidelity

`addon_exports.export` is built to the real FS1 v2 grammar
(`addon/ForeverSixty/Codec.lua`'s `encodeFS1`): head (`FS1:<build>:<class>:<race>:
<talents>:<gear>`), then `level=`, `bags=`, `professions=` sections in that fixed order.
`bank=`/`sets=`/`loadouts=`/`guild=`/`who=` are not written — this tool has nothing
honest to put in them (no bank contents, no alternate talent loadouts, and guild
membership is already the real `guild_characters` row, not something to restate here).
`api/internal/fs1` (the only Go reader) only ever surfaces `level`, the gear item ids and
the talent point sums — it does not decode `bags=`/`professions=`/gear enchants at all,
so none of that readiness variety is visible through today's API. It exists for the
**raid readiness board** the design spec names as out of scope but references by name
(§9: *"we already have the data (`bags=` in FS1); nobody renders it per-raider"*) — when
that board ships and reads the raw export client-side
(`web/src/lib/planner/fs1.ts`), the data it needs is already real and already there.

## Ratings: not written, and why

`fight_metrics.execution_score` and every `rating_scores` row are only ever produced from
a **stored fight summary object in R2** — `api/internal/rating.Backfill` (job command
`rating-backfill`, dispatched by `api/cmd/api`'s own `main.go`) reads it via
`readSummary`/`Store.Summaries.Get`, the exact object a real ingest
(`api/internal/reports.Ingest`) writes when a fight is uploaded and verified. This tool
writes `fight_metrics` rows directly from a `fightPlan`/`mockCharacter` — it never ran a
real ingest, so no such object exists for any of its fights. Pointed at this seed's
reports, `rating-backfill` would find them (`staleFights` selects any complete,
non-private report with no `rating_scores` row), fail to read a summary that was never
written, log the error, and skip every one — it would write nothing, honest or otherwise.
**Fabricating a summary object to make the job produce something would mean fabricating a
fake combat log** (GUIDs, per-event combatant state, a parseable Parquet stream) — a
different and much larger kind of invention than a wrong item id, and this tool does not
do it. Nothing the guild home itself reads needs ratings anyway: `HomeRoster` never joins
`rating_scores`, and the public character rating trend (`ReadCharacterRating`) only ever
reads `visibility = 'public'` reports, which these guild-only ones are not and should
never become.

If a future roster "readiness" panel wants per-raider execution scores from *real*
fights, the fix is a real raid night logged through the real companion against this same
guild — not a change to this tool.

## Commands

```sh
# apply (default) - writes the roster, raid nights and claim
DATABASE_URL=... go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894'

# preview exactly what --apply would write, without writing anything
DATABASE_URL=... go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894' --dry-run

# undo: delete everything --apply wrote, restore the owner's rank/verification and the
# guild's claim state to what they were before
DATABASE_URL=... go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894' --remove

# preview exactly what --remove would delete/restore, without writing anything
DATABASE_URL=... go run ./cmd/seedguild --guild 2 --owner-battletag 'hunthrop#1894' --remove --dry-run
```

`DATABASE_URL` is read from the environment and never printed or logged by this tool.
Running `--apply` twice is a no-op the second time (it checks `seed_rows` for this guild's
tag first and prints a notice rather than partially re-applying); running `--remove` twice
is a no-op the second time for the same reason.

## Tests

```sh
cd api
go test ./cmd/seedguild/...                 # pure generators (no database)
TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/forever_test?sslmode=disable \
  go test ./cmd/seedguild/... -run TestApplyThenRemoveRoundTrips -v   # + the integration test
```

The integration test applies against a freshly-seeded "one unverified owner" guild (the
exact production shape), asserts `guilds.Store.Home` and `rankings.Store.Guild` return the
expected roster/report/progression/roster-best shapes and counts, asserts a second apply
does not duplicate anything, then removes and asserts the guild is back to a single
unverified member with no `seed_rows` left anywhere.

## Remove before public launch

**2026-10-04: this seed must be removed from any guild it was applied to before Forever
Sixty's public launch.** It writes real-looking accounts, a real-looking claim and four
raid nights into the production `guilds`/`reports`/`fight_metrics` tables; `'guild'`
visibility keeps it off every *public* page today, but a launched product should not carry
staged design-review data in its own database at all. Run `--remove` (see above) once the
guild page build this seed exists to support has shipped and been reviewed.

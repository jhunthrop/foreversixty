# Guild suite — design

Status: design, 2026-10-26. Owner's direction (2026-10-25): "we should do attendance (raid
sign-up) etc.. think of all the features that are needed to manage a guild." This document is
the whole suite seen at once, then cut into lanes that each ship something a guild can use the
week it lands. Each lane gets its own experience spec from the ux-designer before it is built
(tenet 6) and its own plan.

Two owner decisions were asked and not yet answered; the defaults below stand until they are:

- **Discord sign-up buttons need a bot.** A webhook can post a message but cannot receive a
  click. Default: sign-ups happen on the site (one tap from the Discord post's link); the bot,
  and with it in-Discord buttons and the channel name, is its own later lane.
- **Which loot system first.** Default: soft reserves, because every Classic guild already
  runs them and a reserve is one row; loot council scoring and DKP come after.

## 1. Intent

A guild officer runs a raid week from three or four places today: a Discord thread for the
sign-up, a spreadsheet for attendance and reserves, Warcraft Logs for who died, and a notes
channel for assignments. The guild page already knows the roster, the logs and the gear. The
suite moves the week onto that page, one screen per job, so the officer opens Forever Sixty on
Tuesday, posts the raid, and on Thursday after the pull knows who came, who got what, and what
to fix before the next one. Members get the one line they need: when, whether they are in, and
what to bring.

The addon and the Discord webhook are the two channels out of the page: the addon so the
raid sees assignments and reserves in game, the webhook so the guild sees the post where it
already talks.

## 2. Goals, non-goals, anti-goals

Goals

- A raid is a first-class record: when, where, who signed up as what, who came, what dropped,
  who got it. Everything else in the suite hangs off it.
- Attendance comes from the raid record and the uploaded log, never from a hand count.
- Every officer surface is reachable in two taps from the guild page and renders at 360 px.
- Members never see an officer control; they see their own status and what is asked of them.
- Each lane ships a usable thing on its own and is gated by the guild plan the way the officer
  tools are (free until the first raid night for grandfathered guilds, by the scoped grant).

Non-goals

- A general-purpose calendar. One guild, its raids, nothing recurring beyond "every Thursday
  at 8".
- In-Discord interaction (buttons, slash commands) before the bot lane.
- DKP ledgers and loot council scoring in the first loot lane.
- A guild bank that reads the in-game bank. The bank lane is a list the quartermaster keeps,
  fed by the addon's bag export where it can be.

Anti-goals

- A second place the roster lives. The raid record references roster characters; it never
  copies them.
- A wall of settings. Every default comes from what the guild already did last time.
- Notifications that fire for things nobody asked about. Every post to Discord is an event the
  officer chose (raid posted, raid changed, reserves locked, report ready).

## 3. The pieces, and what exists today

| Piece | Exists | This suite adds |
|---|---|---|
| Roster, ranks, officer rank, claim, invite link | yes (`api/internal/guilds`, tables `guild_members`, `guild_characters`) | nothing new; every piece joins to it |
| Readiness board (gear, enchants, consumables from the export) | yes (`readiness.go`, read-only rows, Nudge copies text) | the board for one raid: only the raid's signed-up roster, with the raid's date as the "fresh export" cutoff |
| Loot council helper (Onyxia's 22 drops, Award/Undo) | yes (`loot.go`, `loot_awards`) | the loot table for the raid's zone, soft reserves on it, BiS wishlists beside the candidates |
| Raids tab (past raid nights from logs) | yes (`raids.go`) | the planned raids before them, and the link from a night to its report |
| Progression | yes | unchanged |
| Attendance | computed over the last 8 reports (`roster_extended.go`, `attendanceWindow = 8`) | stored per raid: signed up / came / late / no-show, from sign-ups + the log |
| The member's "before Thursday" line | no: `home.go`'s `NeedsBeforeNextRaid` is a gear-needs list with no date | the line, built from the next planned raid |
| Loot tables | Onyxia's Lair only, in `data/builds/<build>/loot.json`; `data/curated/loot/forever-raid-phases.json` names the three first-tier raids | the zone list and its log/loot names in the curated file; the other two tables when the client data lands (a data lane, not this suite) |
| Discord posts (webhook, per account) | yes (`notify`, user-scoped; its `guild_id` column is Discord's server id) | a guild-scoped webhook the guild's officer connects once, and four guild events |
| Addon | exports gear and bags; receives builds and inbox messages | receives the raid's assignments and reserves (a later addon lane) |

## 4. Lanes, in order

Each lane: product design (this document's section), the ux-designer's experience spec and
mock, plan, build, two persona reviews on captures, ship with a changelog entry.

### Lane G1: raids and sign-ups

The record everything else needs.

- **Zones**: one raid has one zone, picked from `forever-raid-phases.json`'s first tier
  (Barrow Deeps, Hyjal Summit, Onyxia's Lair). That file gains, per raid, `key` (the loot
  key), `log_zone` (the name the combat log reports, filled in as the client data lands) and
  `size`. An officer may plan any of the three now; reserves (G3) exist only where a loot
  table exists. A night of two raids is two rows back to back.
- **Table `guild_raids`**: id, guild_id, zone (the curated `key`), starts_at (timestamptz; the guild's timezone stored on the guild as
  `raid_timezone`, default from the claimant's browser at creation), duration_minutes, title
  (default the zone name), notes (markdown, 2000 chars), status (`planned` | `locked` |
  `done` | `cancelled`), report_id (nullable, set when a log is matched), created_by,
  created_at, updated_at. A guild has at most twelve `planned` raids ahead.
- **Table `guild_raid_signups`**: raid_id, character_key (a verified `guild_characters` row of
  this guild; unverified rows cannot sign up, as the loot helper already rules), plus
  `name`, `class` and `spec` snapshotted at sign-up so the row outlives the roster row (no
  cascade when a character leaves the guild), role (`tank` | `healer` | `dps`), state (`in` | `bench` | `out` | `late`), note (140
  chars), signed_by (user), updated_at; unique (raid_id, character_key). A member signs up
  any of their own characters in the guild; an officer may set anyone's state.
- **Repeat**: "every week at this time" creates the next four as `planned` rows, each
  editable on its own; no recurrence engine.
- **Pages**: the Raids tab gains the planned raids above the past nights (same list, two
  sections). A raid page `/guild/{region}/{ruleset}/{name}/raids/{id}`: the header (zone,
  when in the viewer's local time with the guild's time beside it, countdown), the sign-up
  strip for the viewer's own characters (one tap: In / Bench / Out, role pre-filled from the
  character's spec), the roster by role with counts against the zone's size (20/40), the
  notes, and for officers: edit, lock (no more changes from members), cancel, set anyone's
  state, "copy the sign-up list" as text.
- **Member home line** (the guild page's "before Thursday" line already exists): "Thursday 8
  pm, you are In as Healer, 23 of 40 signed up."
- **Discord**: a **guild-scoped webhook** (`guild_notification_webhooks`, one per guild,
  `guild_id` = our guild, `discord_guild_id` = Discord's server; the same sealed shape and
  poster as the account one, connected by an officer through a guild connect route whose
  state is stored server-side (`notify_connect_states`: state, user, guild, expiry) because
  the cookie carries only a random value today; the callback re-checks officer rank before
  the exchange) and the event `raid_posted` (one embed: zone, when,
  link to the raid page, the sign-up counts) and `raid_changed` (time or cancel). Posts are
  the officer's choice: a checkbox on the raid form, remembered.
- **Gate**: planned raids are visible to every member; creating/editing/locking is an
  officer control behind `CanForGuild(officer_views)` like Award; members' own sign-ups are
  free (a member of a guild without the plan still signs up to a raid an officer posted
  during the grant or by hand).
- **Acceptance**: a member signs up from the raid page and from the home line; an officer
  posts a raid, it appears in the test Discord channel with a working link; the roster counts
  update without a reload (a 10 s poll like server runs); 360 px captures of the raid page in
  five states.

### Lane G2: attendance and the raid-night report

- **Matching a log to a raid**: at the one shared point where a guild report becomes
  complete (both the upload parse in `parse.Report` and the live-logging path in
  `reports/ingest.go` end there; add that point if it does not exist), pick the
  `planned`/`locked` raid whose `log_zone` equals the report's zone and whose window
  (starts_at − 1 h, starts_at + duration + 6 h) contains the report's upload time. Combat
  logs carry no absolute clock, so upload time is the anchor and the window is wide; an
  officer can set or clear the match by hand on the raid page. Set `report_id`, status
  `done`.
- **Attendance rows** (`guild_raid_attendance`: raid_id, character_key, signed state,
  outcome `present` | `late` | `absent` | `no_show` (signed In, not in the log), first_seen_at,
  pulls_present / pulls_total): built from the log's participants (the fight rows the parse
  already stores) against the sign-ups; `late` when first seen after the first pull plus 15
  minutes. Recomputed on re-match; never edited by hand except an officer's "excused" flag.
- **Member record**: on the roster row and the character's guild line: attendance over the
  last 8 and the last 30 raids as `present / signed`. The existing `Attendance{Present,
  Nights}` (the loot council's tiebreak) switches per guild: once the guild has one `done`
  raid with attendance rows it reads those (present over raids the character signed in or
  attended), until then the report count it reads today. The contract field keeps its name
  and shape.
- **Raid-night report** (`/raids/{id}` after the match, the "Report" section): kill order
  with kill time per boss and wipes before it; deaths by raider; consumables used (from the
  log's buff events where the parser has them) against the readiness board's expectations
  (the board gains a "next raid" filter in G1: only the raid's signed-up roster, the raid's
  date as the fresh-export cutoff); the raid's loot awards by boss (`loot_awards` gains a
  nullable `raid_id`; combat logs carry no loot events, so the awards are the record); the five biggest
  parse moves against the previous raid of the same zone. One "copy for Discord" text
  summary and the `report_ready` webhook event.
- **Gate**: the report and attendance rows are officer views; a member sees their own
  attendance and their own row of the report.
- **Acceptance**: an uploaded fixture log matches the fixture raid; attendance rows match a
  hand-counted table; the report's kill times equal the log's; the Discord summary posts.

### Lane G3: soft reserves and wishlists

- **Reserves** (`guild_raid_reserves`: raid_id, character_key, item_id, created_at; one or two
  per character, the raid's `reserve_limit` set by the officer; unique per character+item):
  a member reserves from the zone's loot table on the raid page before `locked`; the list
  shows every reserve by item so the raid sees contention before the pull. Locking the raid
  locks reserves. Officer: "copy reserves" in the common SR text format (`Item - Name`) for
  the in-game addons that read it; our addon reads them directly (later addon lane).
- **Wishlists** (`guild_wishlist_items`: guild_id, character_key, item_id, rank 1–5, source
  `bis` | `manual`): seeded from the character's BiS list for its spec (the BiS page already
  computes it) and editable by the member. The loot council helper shows, beside each
  candidate, "reserved" and the wishlist rank, and sorts reserved candidates first.
- **Loot table**: the hard-coded Onyxia list becomes the zone's table from
  `data/curated/loot` (the BiS pipeline already has every raid drop by boss); the helper and
  the reserve picker read it by zone.
- **Gate**: reserving and wishlists are member actions (free); the council view with ranks is
  the officer view.
- **Acceptance**: two members reserve the same item and both see the contention; the SR text
  matches the format the common addons import; the helper sorts reserved first.

### Lane G4: assignments

- A per-raid, per-boss assignment sheet (`guild_raid_assignments`: raid_id, encounter_id,
  slot (tank1, tank2, healer group 1–8, interrupt 1–n, curses, buffs, kite), character_key,
  note): seeded from the raid's roster by role, edited by officers, read by everyone on the
  raid page and in game (the addon lane). "Copy as text" per boss. The `assignments_posted`
  webhook event on demand.

### Lane G5: Discord bot, guild channel

- A bot the officer installs to the guild's server: sign-up buttons on the raid post, the
  channel name on the account and guild settings pages (the deferred item from the connect
  lane), `/raid` to post the next raid. The bot replaces the webhook for guilds that install
  it; the webhook stays for the rest.

### Lane G6: bank and recruitment

- **Bank**: a list the quartermaster keeps (item, count, holder character, note), with the
  addon's bag export of designated bank characters filling counts; a "needs" list members can
  request from; nothing in game.
- **Recruitment**: a guild's open spots (class/spec/role, count) on its public page, with an
  "apply" that creates a pending `guild_characters` row the officers already approve today;
  the rankings page's guild row shows "recruiting: 2 healers".

## 5. Cross-cutting

- **Time**: every stored time is UTC; every rendered time is the viewer's local with the
  guild's zone beside it when they differ. The guild's `raid_timezone` is an IANA name.
- **Gate**: one check, `CanForGuild(officer_views)`, for every officer write; member writes
  (sign-up, reserve, wishlist) are free so a guild on a lapsed plan still fills its raid; the
  read-only sentence pattern from the officer tools gate is reused verbatim for officer
  controls.
- **Discord**: the guild webhook is the account webhook's twin (same `notify` package, same
  seal, same disable-after-five rule, `guild_id` instead of `user_id`); events are chosen per
  raid, never global.
- **Addon**: G1, G3 and G4 each name what the addon will receive (sign-up state, reserves,
  assignments) but the addon work is one later lane so the in-game half ships together.
- **Polish**: every list is sorted the way the game sorts (role, then class, then name); class
  colours and the ringed crest everywhere a character appears (tenet 7, one crest language).

## 6. Order and dates

G1 → G2 → G3 → G4, each a week or less at the current cadence; G5 and G6 after launch day (4
Nov) and before the first raid night (9 Dec), G5 only if the owner wants the bot. G1 is the
next lane after this document is approved.

## 7. Left to the builder

Table and index names beyond those given; the parser-side hook for G2's match (where
`parse.Report` already posts the Discord event is the natural place); the SR text format's
exact line shape (copy the most common addon's import format); fixture raids and logs; the
raid page's URL slug; how the twelve-planned-raids cap is enforced.

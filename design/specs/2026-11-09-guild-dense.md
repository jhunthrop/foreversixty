# Guild pages, dense: experience spec (batch F, lane guild)

2026-11-09. Owner, verbatim: "the guild pages look incredibly boring and not dense ... this is the premium differentiator for the site. We have to make it feel uber valuable to guild members and officers."

Inputs: the audit (`design/reviews/2026-11-09-f-guild/audit-*.md`: guild-officer NOT YET, info-density THIN), the data inventory (`design/reviews/2026-11-09-f-guild/inventory.md`), the persona `.claude/agents/guild-officer-reviewer.md`. Mocks: `design/mocks/ws_guild_dense.py` (data `design/mocks/guild_dense_data.py`, fixtures `design/mocks/guild_dense_fixtures.py` into `web/tests/e2e/fixtures/guild-parity/dense-*.json`).

Rules carried over: the batch F brief (shared primitives only, compact PageHeader answer slot, game words, no explainer lines, no em dashes) and the web layout spec 2026-11-05 (S0/S1/S2 surfaces, one Plate per region, numbers right-aligned and tabular).

## 1. The bar
An officer opens this instead of the spreadsheet; a member opens it before raid without being told. Every number carries its context: a target, a median, a rank, last week, or a streak. Every officer action is one or two clicks from the row it is about.

## 2. The fixture world (the mocks and the parity states draw it)
- Olympus XXVII, US PvP Alliance, **40 raiders** on the roster plus 3 waiting for approval.
- Clock **Sat 9 Jan 2027, 4:00 pm CST**. Five raid nights logged, Wednesdays 8 pm: 9, 16, 23, 30 Dec and 6 Jan (raids 40 to 44). Onyxia down every night (one wipe on the first). Next raid **Wed 13 Jan** (raid 45), 33 answered.
- Every roster row carries what the API sends: `rating` (overall + six parts), `best_parse` with percentile, `attendance` and `attendance_30`, `professions`, `addon` (build, gear gap, seen), `logged_at`, `level`, `item_level`.
- Loot history: every Onyxia drop of the five nights with its council decision.
- Attendance: per raider per night (present, late, bench, absent, excused) from `attendance/recent`.
- Report parses: every raider's Onyxia figure and rating for raids 43 and 44 (last week and the week before).
The old 24-character fixtures stay for the states this lane does not redraw.

## 3. The officer's week, mapped to the page that answers it

| Job | Page and module | Data |
|---|---|---|
| Can we field Wednesday? | Overview, **Wednesday** strip: In / Late / Bench / Not answered against 40; tanks, healers, melee, ranged against the comp target with the gap in red; ready count of those signed | raid view `counts`, `comp_target`, readiness `failing` |
| Who fills the gaps? | Overview, **Gaps** rows: each short role names its fillers: not answered raiders of that role, then bench | raid view `not_signed_raiders[].role`, signups `state=bench` |
| Who needs a word? | Overview, **Needs a word**: one row per raider with the reason and the action (Ping, Nudge, Profile): not answered, short for Wednesday, missed 2 of the last 4, rating down vs last week, no log in 7 days | `not_signed_raiders`, readiness `failing` + `nudge_text`, `attendance/recent` `slipping`, report parses x2, roster `logged_at` |
| How did last week go? | Overview, **Last night**: kill time vs best, wipes, deaths, raid DPS, the three biggest risers and fallers by rating vs the week before; Report link | raids 44 and 43 `report.parses`, `kill_order`, `night` |
| Who is raid-ready? | **Roster** (one dense table) and **Readiness** (the check matrix) | roster rows + readiness rows |
| Loot fairness | **Loot**, Council table (candidates as columns of figures) and **Fairness** (every raider: items this tier, nights, items per night, last item) | loot candidates, `loot/nights` decisions, `attendance_30` |
| Attendance | Roster column **Last 5** (one cell per night, coloured) and **Streak** | `attendance/recent` cells |
| Recruiting | Overview rail **Recruit for the gaps**: roles short over the last 4 nights' comp, the open applications against them | `comp_target`, `counts`, `applications` |

## 4. The member's week

| Job | Where | Data |
|---|---|---|
| When is it, am I in? | Overview, **Your week** strip: raid, countdown, your answer (In / Late / Bench / Out inline) | raid view `viewer`, signups |
| What do I fix before? | **Before Wednesday** (the checklist, one line per item, no explainer sub-lines) | `me/checklist` |
| How did I do? | **Your Onyxia**: last week's HPS or DPS, rating, vs the week before, rank among the guild's raiders of your role | report parses x2, roster |
| What did I win, am I due? | **Your loot**: items this tier, nights attended, the guild's median items per night | `loot/nights` decisions, roster attendance |
| My standing | Figures: item level and rank of `same_spec_count`, attendance and streak, best parse % | `standing`, roster row |

## 5. Pages

### 5.1 Overview, officer
Header: the shared PageHeader as batch E ships it; answer slot = "Wed 13 Jan · in 4 d 4 h". Tabs as today.
- **Wednesday** (Plate, full width): a Figures row: In 27, Late 2, Bench 2, Not answered 7 (of 40 raiders); then four role rows (Tanks, Healers, Melee, Ranged): Need, In, Late, Gap, and the fillers as class-coloured names with their state word ("Voidcall not answered", "Cinderglass bench"). Officer actions on the Plate: Ping not answered (one click, state shown after: "Pinged 4:02 pm"), Open the planner.
- Main column, below: **Needs a word** (Rows: crest, name, reason with its figure, action button). Then **Last night** (StatTable: raider, rating, change vs the week before, figure; the three up and three down; head line with kill time, wipes, deaths).
- Rail: **Cover** (the buff and debuff table as today, with the provider named on a gap), **Attendance** (the last four nights, the slipping rows), **Loot last night** (drops and who won), **Recruit for the gaps**.
- Action needed stays at the top as one line of buttons when it has items (approvals, applications, drops to award).

### 5.2 Overview, member
- **Your week** (Plate): raid and countdown, your answer segmented control, then Figures: Item level 64 (2nd of 3 Holy Priests), Attendance 5 of 5 (streak 5), Best parse 82 %, Items this tier 1.
- **Before Wednesday**: checklist rows, the reason as data ("Hands, Finger 1 · +11 HPS"), link to the BiS page.
- **Your Onyxia**: last two nights side by side: figure, rating, deaths; change.
- Rail: the raid's comp (as today) and **Last night** (guild line).

### 5.3 Roster (members and officers)
One Plate, one StatTable-grammar table, 40 rows, grouped by role with a group head that carries the comp need ("Healers 9 · need 8"). Columns (desktop): Raider (crest, name in class colour, spec, rank word for officers), Item level (with ± vs median under it), Gap (DPS or HPS from `addon.gear_gap` or readiness), Rating, Best (% with the figure as tooltip), Last 5 (five cells: present, late, bench, absent, before joining), 30 nights (%), Loot (items this tier), Ready (a tick or the count short, links to Readiness), Seen (relative). Sort on every numeric head. Default sort: role group, then rating.
Waiting for approval collapses to one Notice line with Review (opens the existing list) and Approve all.
Phone (390): the row is crest, name and spec; then one line of three figures (ilvl, rating, 30 nights) and the Last 5 cells; the rest behind the row's link to the profile.

### 5.4 Readiness
The matrix of every raider signed for Wednesday, the ready included: columns Attuned, Item level ±, Gap, Enchants (missing slots), Flask, Food, Talents, Fire res (tanks), Synced. Failing cells red, passing muted ticks. Head: "26 of 31 ready" and Nudge all short (one click; each row keeps its Nudge).

### 5.5 Loot
- Left: **Drops** table for the night picker's night: item (icon, quality name), slot, top candidate and gain, state (Open, Awarded to X).
- Right: **Council** for the picked drop as a comparison table: candidate, gain, ilvl ±, nights (30), items this tier, last item, votes, Award. One click to Award; the decision lands in the record below.
- Below, full width: **Fairness**: every raider who attended a night this tier: nights, items, items per night (bar against the guild median), last item and when. Sort on each head.
- **Decision record** keeps its rows, with the voters' names.

## 6. Primitives
Only the shared ones: PageHeader, Plate, Section, Row/Rows, Figures/Figure, StatTable/StatRow, Tag, Button, Segmented, Notice, Meter, ItemIcon, ClassCrestRing. One additive change to the shared library, told to the integrator: `StatTable` gains `dense` (36 px rows, 6 px padding) for 40-row tables, and an attendance cell strip lives in `components/guild/` as a guild component (it is not used elsewhere).

## 7. Needs data (for the owner; not built)
1. Per-parse percentile vs the spec (`percentile_digests` exists; `report.parses[]` carries only `rating`). "vs spec median" uses the rating until this lands.
2. Real sim DPS per raider (`character_sims`) on roster and readiness rows (today the gap is the BiS-file gain).
3. Member tenure (`guild_members.joined_at`) and trial state on the roster.
4. Answer time and who moved a sign-up (`guild_raid_signups.updated_at`, `acted_by`).
5. Avoidable damage and per-fight deaths (`fight_metrics.damage_taken`, `rating_scores.deaths`) in the report.
6. Drop method and time (`raid_drops.loot_method`, `dropped_at`).
7. Flask, food and potions split (one consumables state today).
8. A guild activity feed (`activity_events.guild_id` has no guild read).
9. A fire resistance target per raid (the raid only says it asks for fire resistance; the matrix shows the tanks' totals with no pass mark).
10. Per-voter split and vote times on past decisions (`loot_votes.voter_id`, `voted_at`).
11. Per-raider deaths and their causes on the officer's last-raid movers (`rating_scores.deaths`; the report's `deaths[]` is per raider but not per cause outside the pull view).
12. The in-game rank number (`guild_characters.rank_index`).

Landed in this lane (API): roster rows carry `group` (raidspec.Group); report parses and roster `best_parse` carry `percentile` from `percentile_digests` (the roster's was always null before).

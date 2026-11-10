# Component audit, 10 Nov 2026: copied UI and the consolidation plan

Owner, 2026-11-10: "find every instance of components we copy across the codebase and centralize and unify the experience."

**Scope:** a read-only audit of:
- `web/src` (Astro and Svelte);
- the server-rendered HTML in `api/internal/site` and `api/internal/auth`;
- the Discord embed builders;
- `companion/ui`;
- the iOS and Android apps, only where one platform builds the same thing twice.

**Excluded:** talent trees (`lane/talent-tree`) and page width and frame (`lane/header-width`). Their leftovers are in section 4.

**Method:** grep for markup, class recipes and helper definitions, then read the matching excerpts. Every location was checked with grep at `0637711a0`. `src/` means `web/src/`. "Users" means non-test files that import or call a thing.

**What decides the best basis:** `docs/tenets.md` (2 and 3: an item or ability is never just a name; 9: every state; 10: one design system), `design/specs/2026-11-05-web-layout.md` §3A (the shared primitives), and the review rulings cited in each cluster.

## Summary

The ranking is: what a player sees first, then how often it appears, then risk.
- **Vis** is player visibility. H means the same thing looks or reads differently on two pages; M means a state or control behaves differently; L means code only.
- **Size** is the number of files touched.

| # | Cluster | Vis | Copies | Existing primitive | Size | Risk | Lane |
|---|---|---|---|---|---|---|---|
| 1 | Spell and ability icons | H | 4 components + 17 hand-drawn `<img>` | `ui/SpellIcon` | ~35 | M | A |
| 2 | Class crests and portraits | H | 9 ring recipes, 2 bare | `ClassCrestRing` | ~19 | M | D |
| 3 | Dates and "Updated" stamps | H | ~12 day-month, 8 weekday, 5 relative | `lib/dates.ts` | ~25 | M | B |
| 4 | Parse percentile figures | H | 10 sites | `percentileToken` (colour only) | ~10 | L | B |
| 5 | Raid sign-up control and save | H | 5 `putSignup` callers, 2 label sets | `SignupSegments` | 5 | M-H | D |
| 6 | Zone and raid icons | H | 4 sites | none | 4 | L | A |
| 7 | Numbers and durations | H | ~45 inline, 8 helpers, 9 duration functions | `lib/report/format.ts` | ~30 | L | B |
| 8 | Countdowns | H | 4 formats, 3 tick timers | `ui/Countdown` (1 user) | 4 | L | B |
| 9 | Item icon and inline item name | H | 1 hand-drawn icon, ~13 inline names | `ItemIcon`; no name primitive | ~16 | L | A |
| 10 | Empty states | M | ~40 hand-written vs 6 users | `EmptyState`, `report/EmptyLine` | ~35 | L | E |
| 11 | Error and retry | M | 5 retry styles, 19 "Try again" strings | `LoadError` | ~8 | L | E |
| 12 | Filter chips | M | 4 recipes | `ui/Chips` (3 users) | ~6 | M | E |
| 13 | Copy buttons | M | 12 inline sites, 3 parallel helpers | `createCopyFeedback`, `CopyLinkField` | ~14 | L-M | E |
| 14 | Filter popover, focus trap, confirms | M | 3 + 3 + 6 | `Sheet`, `account/ConfirmButtons` | ~12 | M | E |
| 15 | Section heads | M | 4 hand-written h2, 38 raw h3 | `Section` | ~8 | L | E |
| 16 | Show more and pagination | M | 8 patterns | none | ~7 | L | E |
| 17 | Status pills used instead of Tag | M | 3 files | `Tag` | 3 | L | D |
| 18 | Class-coloured character name | M | 22 inline vs 3 users | `guild/GuildCharacterName` | ~25 | L | D |
| 19 | Tabs: BossPicker | M | 1 | `TabStrip` | 2 | L | E |
| 20 | Server-rendered pages | M | 2 | n/a | 3-4 | L | C |
| 21 | Discord and Go formatting | M | 2 separators, 4 + 4 date layouts | n/a | ~12 | L-M | C |
| 22 | Companion formatters | M | `format.js` | n/a | 2 | L | C |
| 23 | iOS: one component built twice | M | 5 icon tiles, 4 error cards, ~12 formatters | DesignSystem (partly) | ~25 | M | F |
| 24 | Android: one component built twice | M | 4 error regions, 2 + 12 skeletons, ~14 formatters | core/designsystem (partly) | ~30 | M | F |
| 25 | Icon URL builder | L | 36 sites in 33 files | none | ~30 | L | A |
| 26 | Class and spec names, slugs, colours | L | 7 slug lists, 4 spec-name functions, 15 capitalise helpers | `spec-label.ts` (partly) | ~25 | L | B |
| 27 | Gear-slot constants, average item level | L (one real divergence) | 5 slot orders, 2 item-level rules | `FRAME_SLOTS` | ~8 | L | A |
| 28 | Tooltip positioning, disclosures, dead code | L | 3 tooltip systems, 5 `<details>`, 3 dead files | yes | ~10 | L-M | E |

## Clusters

### 1. Spell and ability icons

**Locations: four components**
- `src/components/ui/SpellIcon.svelte:34-54` (6 users, all on report pages).
- `src/components/report/SpellIcon.svelte:10-30` (8 users).
- `src/components/coach/SpellIcon.svelte:14-42` (7 users, including `guild/report/PullDeaths.svelte:12` and `guild/overview/BuffCover.svelte:10`).
- `src/components/sim/ActionIcon.svelte:23-45` (4 users), plus `sim/ActionMark.svelte:54`.

**Locations: hand-drawn `<img>` icons (17)**
- Report and sim: `report/AbilityTable.svelte:120`, `sim/ActionStatRow.svelte:54-64`, `sim/BuffPanel.svelte:109`, `sim/ImbueRows.svelte:44`, `sim/RunReport.svelte:284`.
- Charts: `chart/ChartTooltip.svelte:106`, `chart/TimeSeriesPlot.svelte:793`.
- Guides: `guides/GuideSpellRow.astro:27`, `GuideRacials.astro:22`, `GuideExtrasLane.astro:56,92`, `GuideSpecRow.astro:42,61`.
- Other: `planner/GuideBuilds.svelte:60`, `ui/IconList.svelte:36`, `ui/ControlBar.svelte:163`.

**How the copies differ**

| Copy | Frame | Radius | Empty icon | Hover tooltip | Alt text |
|---|---|---|---|---|---|
| `ui/SpellIcon` | 1 px warm line | 5 px | dark socket | yes (`data-game-tip`) | wrapper hidden |
| `report/SpellIcon` | none | 4 px | faint white ring | no | none |
| `coach/SpellIcon` | black edge and white ring, art scaled 1.08 | 4 px | nothing; socket only on error | no | name as title/aria-label |
| `sim/ActionIcon` | none | 4 px | skeleton while loading | no | the label (the only real alt) |

**What a player sees**
- On one report page, CastTable and AuraTable icons are framed and 28 px. SummaryPanels (:187, :196), GroupBuffs (:94) and DeathsTab (:397, :413) icons are unframed and 24 px.
- Only the framed ones show the spell on hover, which breaks tenet 3.
- The same consumables are bare 20 px icons in BuffPanel and ImbueRows, but framed 28 px with a white quality border in `sim/tools/ConsumableCandidates.svelte:48`.

**Basis:** `ui/SpellIcon`.

**Proposed API:** `SpellIcon {stem?, build?, url?, size: 16|20|24|28|32|36, tip?: spellId, label?, loading?}`.
- Add a static `SpellIcon.astro` twin for the guide pages.
- Delete the report, coach and sim copies.
- Buffs, enchants and consumables use `SpellIcon`, never `ItemIcon`.

**Size and risk:** about 35 files, medium. The e2e testids `coach-spell-icon` and `action-icon` move.

### 2. Class crests and portraits

**Locations.** All copies share `src/lib/class-crest.ts:14`. The ring recipe is copied nine times:
- `ClassCrest.astro:69-113`: plain ground, colour taken from the `classes.json` hex at `:65`.
- `character/ClassCrestRing.svelte:50-104` (56 users): tinted ground, a silhouette for an unknown class, colour from tokens.
- `character/CharacterPortrait.svelte:49-57`: copies the silhouette (`:46`) and has no `hasClassCrest` guard, so an unknown class gives a broken image.
- `nav/SelectorCrest.svelte:21-33` with `styles/nav.css:653-664`.
- `styles/nav.css:1137-1147`.
- `styles/nav.css:260-271`, used by `HeaderPanel.astro:91-100`.
- `styles/typeahead.css:40-54`, built in `lib/character-typeahead-dom.ts:53,80`.
- `ui/CharacterAnswer.svelte:68,186-191`.
- `ui/segmented.css:80-88`, used by `ControlBar.svelte:156`.

**Locations: bare crests with no ring**
- `Rankings.svelte:721-727`: a 28 px `<img class="rounded-full">` passed in as the RankRow lead.
- `rankings/RankingsLanding.svelte:133-139`: the same.
- The same RankRow gets a ringed 36 px crest on /tiers (`tiers/TierRankRow.astro:20`).

**What a player sees:** plain discs on /rankings and ringed crests on /tiers in the same row. Ring weight also differs in the selector, typeahead, answer box and control bar. This breaks the owner's "one crest language" ruling.

**Basis:** `ClassCrestRing`.

**Proposed API:** `Crest {cls, size, as?: 'span'|'a'|'button'|'summary', href?, eager?}`.
- Its CSS moves to one `styles/crest.css`, so Astro, Svelte and DOM-built crests all use one class.
- RankRow draws the crest itself from `classSlug`.

**Size and risk:** about 19 files, medium. It runs after `lane/header-width`, because it touches nav.css and HeaderPanel.

### 3. Dates and "Updated" stamps

**Locations: day-month formatters**

| Location | Output | Time zone |
|---|---|---|
| `lib/dates.ts:5` (months at `:1`) | "Sept", "June", "July" | UTC |
| `lib/dates.ts:84` (a second table in the same file) | "Sep" | UTC |
| `lib/guides/guide-header.ts:5-11` | "Sep" | UTC |
| `lib/sim/head-meta.ts:31-34` | "Sep" (strips "Sept" by hand) | UTC |
| `lib/character/page-view.ts:222` | "Sept" | UTC |
| `lib/home/home-view.ts:61` | "Sept" | viewer |
| `lib/score/score.ts:167` | "Sep" | date only |
| `lib/coach/dates.ts:8`, `:15` | "Sept" | viewer / UTC |
| `lib/sim/result-band.ts:74` | "Sept" | UTC |
| `lib/report/format.ts:101` | "Sept" | UTC |
| `lib/guild/invite-link.ts:15-18` | "Sep" | UTC |

**Locations: weekday dates ("Wed 9 Dec")**
- Guild components: `guild/dense/MemberWeek.svelte:86`, `OfficerWeek.svelte:131`, `FairnessTable.svelte:57`.
- Guild lib: `lib/guild/loot-council-view.ts:290`, `raid-time.ts:154`, `zone-catalogue.ts:43`.
- Coach and home: `lib/coach/last-raid.ts:38`, `lib/home/feed.ts:146`.

**Locations: long dates**
- `lib/home/plaque.ts:104`
- `lib/guild/raid-page.ts:66`
- `lib/sim/phase.ts:118`

**Locations: relative time**
- `lib/dates.ts:52` `relativeTime`.
- `lib/sim/sources.ts:113`: also named `relativeTime`, never switches to a date, so it can print "400 days ago".
- `lib/dates.ts:72` `compactAgo`.
- `lib/home/feed.ts:136` `agoText`: a copy of `compactAgo` with no guard against a bad date.
- `lib/dates.ts:91` `relativeDay`.

**Locations: viewer time zone**
- Three helpers: `lib/dates.ts:198`, `lib/guild/raid-copy.ts:418`, `lib/guild/week-one.ts:203`.
- About 10 inline lookups, e.g. `HomeRaidPlaque.svelte:56`, `GuildOverview.svelte:99`, `RaidBandLine.svelte:93`.

**Locations: clock style**
- 12-hour: `lib/home/plaque.ts:110` and `guild/dense/WednesdayPanel.svelte:66`.
- 24-hour: `lib/dates.ts:111`.

**Locations: "Updated" stamps.** `UpdatedStamp.astro` has one user (`tiers/TierPage.astro:115`). Hand-written ones:
- `Rankings.svelte:461-462`
- `pages/changelog.astro:73`
- `pages/guides/index.astro:63`
- `pages/guides/[class]/index.astro:115`
- `pages/guides/[class]/[spec].astro:250`
- `pages/bis/index.astro:56`

**What a player sees:** the same day reads "Updated 25 Sept" on /tiers, /rankings and /changelog and "25 Sep" on /guides and /bis. Some rows use UTC and others local time, so late-evening events land on different days.

**Basis:** `lib/dates.ts`. Its own comment (`:3`) calls it "the house date".

**Proposed API:** in `lib/dates.ts`, one month table plus:
- `dayMonth(d, {zone?, year?})`;
- `weekdayDayMonth`;
- `longDay`;
- `clock(d, zone)`;
- `compactAgo` and `relativeTime`, renamed so they don't collide;
- `viewerZone()` as the only time-zone lookup;
- `updatedLabel()`, used by `UpdatedStamp.astro` and a new Svelte twin.

**Size and risk:** about 25 files, medium; tests currently pin both spellings. Blocked on the "Sept" or "Sep" decision.

### 4. Parse percentile figures

**Locations**
- The colour ladder is `lib/report/format.ts:335-343` `percentileToken`. It floors on purpose: "Floored, never rounded (Warcraft Logs' own rule)". 9 users.
- Numbers rounded instead:
  - `character/CharacterFightTable.svelte:89,172`
  - `lib/character/page-view.ts:283`
  - `lib/character/history-chart.ts:381`
- Numbers floored:
  - Report: `report/SummaryTab.svelte:121`, `report/ActorTable.svelte:117`.
  - Recruit: `guild/recruit/ApplicantRow.svelte:86`.
  - Dense guild views: `guild/dense/MemberWeek.svelte:132,171,182`, `OfficerRail.svelte:62`, `GuildRaidRoster.svelte:163`.
- Report parse cells are uncoloured. `ApplicantRow.svelte:83-86` colours a dot but leaves the number white.
- Only `CharacterFightTable.svelte:88-90` gives screen readers a label.
- `score/RatingsPage.svelte:34-40` copies `lib/score/score.ts:90`.
- `score.ts:105` hard-codes `#3d94f0` and `#b866f5`, which already exist as tokens (`tokens.css:89-90`).

**What a player sees:** a 98.6 parse reads "99" on the character page and "98" in the guild roster. The character page "99" is orange, because the colour floors to 98. The ruling "one number format" (`design/reviews/2026-11-09-h-audit/round1/slop-rankings.md:29`) applies.

**Proposed API:** `parseFigure(p)` and `parseText(p)` in `lib/report/format.ts`, which floor, colour and label the value.

**Size and risk:** about 10 files, low.

### 5. Raid sign-up control and save

**Locations**
- `SignupSegments` is shared at `guild/SignupRow.svelte:290`, `guild/RaidBandLine.svelte:282,311` and `guild/planner/PlannerRow.svelte:89`.
- `home/HomeRaidPlaque.svelte:108-121` re-implements it with three buttons: "I'm in", "Late", "Can't come" (`lib/home/plaque.ts:20,24`). The guild has four: In, Late, Bench, Out (`lib/guild/raid-copy.ts:69`).
- `putSignup` has five callers with four different save behaviours:
  - `RaidBandLine.svelte:224`: handles refusals and focus at `:212-237`.
  - `SignupRow.svelte:193`: serial, optimistic save.
  - `HomeRaidPlaque.svelte:75`: a generic failure line.
  - `guild/dense/WednesdayPanel.svelte:89`.
  - `lib/guild/raid-page-controller.svelte.ts:270,300`.

**Bug:** `WednesdayPanel.svelte:84-93` `moveIn` has no `catch`, and its caller at `:163` discards the promise (`void`). A failed move is an unhandled rejection and the officer sees nothing.

**What a player sees:** different words and button shapes for the same sign-up on home and the guild page, and different behaviour when it fails.

**Proposed API**
- A shared `useSignupSave({guildId, raidId, onChanged})` in `lib/guild/`, based on SignupRow's serial save plus RaidBandLine's refusal handling.
- The plaque renders `SignupSegments`, with a `choices` subset if three choices stay.

**Size and risk:** about 6 files, medium-high.

### 6. Zone and raid icons

**Locations.** The icon comes from `lib/first-raids.ts:5-14` `raidZoneIcon`. It is drawn with `ItemIcon`, whose default quality is 1 (`ui/ItemIcon.svelte:17`):
- Quality 1, a white frame: `home/HomeFirstRaids.svelte:63` and `RaidsCountdown.svelte:33`.
- Quality 5, an orange frame: `guild/overview/NextRaidRow.svelte:33` and `guild/GuildRaidNights.svelte:110-113`.
- These zone rows have no icon at all: GuildProgression, LastRaid, RankingsEncounters, ReportRow, HomeRaidPlaque, RaidDoneBand, MyReports, EncounterProgress.

**What a player sees:** the same raid in a white frame on home and an orange "legendary" frame on the guild page.

**Proposed API:** `ZoneIcon {zone, size, build}`, using the spell-icon frame and showing nothing for an unknown zone.

**Size and risk:** about 5 files, low.

### 7. Numbers and durations

**Locations: number helpers**
- `lib/report/format.ts:114` (grouping) and `:120` (k/M/B from 100k).
- `lib/chart/value-axis.ts:35-36` (k from 1,000, no B).
- BiS, tiers, planner: `lib/bis/tank-view.ts:19-20,75,80`, `lib/tiers/tier-copy.ts:362-363`, `lib/planner/score-unit.ts:54`.
- Sim: `lib/sim/result-copy.ts:10,15`, `lib/sim/compare.ts:158`, `lib/sim/server-runs-copy.ts:7`.
- Other: `lib/golinks/funnel.ts:21`.

**Locations: inline formatting**
- `toLocaleString(` appears 61 times in about 38 files.
- These have no locale, so a German-locale viewer sees "1.234":
  - `report/EventsView.svelte:163,211,229`
  - `report/QueriesView.svelte:232`
  - `lib/report/events.ts:130,146,149,150`

**Locations: durations.** Nine functions in four styles:
- Clock style:
  - `lib/report/format.ts:71` "1:42" or "4.2s" (about 69 calls)
  - `lib/report/format.ts:29`
  - `lib/sim/encounter.ts:9`
  - `killTimeLabel`
- Words:
  - `lib/guild/attendance-copy.ts:38` "2 h 41 min"
  - `lib/guild/raid-form.ts:50`
  - `lib/logs/report-summary.ts:14` "2 h 10 m"
  - `lib/guides/rotation-lists.ts:114` "1 min 30 s"
  - `lib/guild/raid-copy.ts:455`
- Three different exported functions are all called `durationLabel`.
- Seconds are "4.2s" in reports and "4.2 s" in `sim/DetailsCard.svelte:43`.

**Locations: other small copies**
- Ordinal twice: `lib/report/format.ts:301` and `lib/tiers/tier-copy.ts:230`.
- The empty-cell dash constant is defined in 10 files, plus 18 inline literals. `lib/gear/doll-cells.ts:31` uses an em dash while every other site uses an en dash.

**What a player sees:** a raid is "2 h 10 m" on /logs and "2 h 41 min" on attendance.

**Proposed API**
- `lib/format/number.ts`: `grouped`, `oneDecimal`, `compact`, `ordinal`, `EMPTY_FIGURE`.
- `lib/format/duration.ts`: `clock(ms)`, `spoken(ms)`.

**Size and risk:** about 30 files, low.

### 8. Countdowns

**Locations: four formats**
- `lib/ui/countdown.ts:19-33` `clockParts`, zero-padded. Only one chain uses it: Countdown, then RaidsCountdown, then `guild/GuildWeekOneCard.svelte:110`.
- `lib/home/plaque.ts:38-46` `plaqueClock`: days and hours, unpadded. Used by `home/HomeFirstRaids.svelte:38` and `HomeRaidPlaque.svelte:66`.
- `lib/guild/raid-copy.ts:453-468`: "in 1 day 4 h".
- `lib/character/standing.ts:100-102`.

**Locations: three tick timers**
- Countdown's own timer.
- `lib/guild/minute-clock.svelte.ts:12-27` (11 users).
- A hand-written interval at `Character.svelte:158-162`.

**What a player sees:** "First raids open" shows "32 days 5 hours" on home and "32 days 05 hours 12 min" on the guild week-one card.

**Proposed API:** `clockParts(target, now, {units, pad})` replaces `plaqueClock`. The minute clock moves to `lib/ui/`.

**Size and risk:** 4 files, low.

### 9. Item icon and inline item name

**Locations: hand-drawn item icon.** `guild/RaidLootItem.svelte:72-90`:
- 2 px border (ItemIcon uses 1 px).
- A tile with the item's initials where ItemIcon shows a dark socket (`lib/ui/item-icon.ts:16`).
- The loot tab (`guild/loot/DecisionRow.svelte:32`) uses ItemIcon, so the two loot views differ.

**Locations: non-items drawn as white-bordered items**
- `sim/tools/ConsumableCandidates.svelte:48`
- `sim/tools/EnchantList.svelte:59`
- the zone icons in cluster 6

**Locations: inline item names (about 13), no shared component**
- Home: `home/FeedSentence.svelte:17`, `home/HomeUpgrades.svelte:65`.
- BiS and character: `ItemHover.svelte:217`, `bis/BisSlotLines.svelte:30,39`, `character/GearDoll.svelte:110`.
- Report: `report/AbilityTable.svelte:147`.
- Sim tools: `sim/tools/CandidateRows.svelte:94`, `ItemSearch.svelte:155`, `SourcePicker.svelte:168`, `SubstitutionChips.svelte:66,72`.
- Guild: `guild/ZoneKeyNote.svelte:23`, `guild/RaidLootItem.svelte:95`.
- Guides: `guides/GuideExtrasLane.astro:105`.

**Bug:** `home/HomeUpgrades.svelte:65` puts the tooltip hook on a span with no `tabindex`, so keyboard users can't open that item's tooltip (tenets 2 and 9).

**Proposed API**
- RaidLootItem uses ItemIcon, with a new 16 px size.
- New `ui/ItemName {id, name, quality, unknown?, focusable = true}`.

**Size and risk:** about 16 files, low.

### 10. Empty states

**Locations: report copies of EmptyLine.** `EmptyState` has 6 users and `report/EmptyLine.svelte` has 6. These hand-write the same line:
- `report/EventsView.svelte:183`
- `report/CastTable.svelte:219`
- `report/ExchangeTable.svelte:248`
- `report/ResourceGraphs.svelte:232`
- `report/TimelinesView.svelte:338`
- `report/ThreatTable.svelte:450`
- `report/AuraTable.svelte:336`
- `report/QueriesView.svelte:250`

**Locations: empty states with actions, five layouts**
- `ui/EmptyState.svelte:19-36`
- `MyReports.svelte:141-152`
- `account/CharacterList.svelte:182-192`
- `rankings/RankingsEmpty.svelte:30-37`
- `home/HomeFeed.svelte:99-101`

**Locations: one-line empties that should be EmptyState.** About 25. Outside the guild:
- `HomeTopGuilds.svelte:60`
- `character/ParseChart.svelte:87`
- `planner/BandCompare.svelte:244`
- `admin/GoLinksAdmin.svelte:202`, `admin/AppTestersAdmin.svelte:97`
- `report/ThreatRail.svelte:70`
- `sim/SimGearSection.svelte:108`, `sim/SimYourSims.svelte:97`, `sim/tools/SlotGrid.svelte:52`
- `tiers/TierPage.astro:146,150`

Guild:
- `guild/GuildPlannedRaids.svelte:37`
- `GuildWeekOneCard.svelte:132`
- `RaidMatchPicker.svelte:126`
- `GuildDiscordChannelPick.svelte:162`
- `recruit/OfficerNotes.svelte:48`
- `GuildRaids.svelte:260`
- `planner/PlannerBench.svelte:41`, `planner/PlannerRoster.svelte:63`, `planner/RaidPlanner.svelte:377`
- `loot/LootDrops.svelte:66`, `loot/DecisionRecord.svelte:47`
- `report/ReportLoot.svelte:42`, `report/ReportPulls.svelte:84`, `report/ReportOutput.svelte:34`
- `billing/GuildSeatPicker.svelte:64`

**Locations: loading placeholders**
- Hand-written skeletons with no status role:
  - `planner/Planner.svelte:1139,1270-1273,1306-1308`
  - `Character.svelte:275`
  - `tiers/TierSpecCard.astro:31-34`
- `sim/SimYourSims.svelte:91-95` nests several Skeletons inside another busy region, so a screen reader hears "Loading" again and again.

**What a player sees:** padding, colour and action layout of "nothing here" differ between Reports, Account, Rankings, Home and every guild tab.

**Basis:** `RankingsEmpty`, the most complete version.

**Proposed API:** `EmptyState {line, fill?, action?, link?, inset?}`. `inset` replaces EmptyLine. Add Astro twins of EmptyState and Skeleton.

**Size and risk:** about 35 files, low.

### 11. Error and retry states

**Locations: hand-written errors.** `LoadError` has 26 users. Each copy below has its own retry style:
- `Rankings.svelte:553-562`: a link.
- `account/AccountPage.svelte:392-398`: a primary button.
- `sim/SavedSimLoadFailed.svelte:15-21`: a secondary button.
- `Guild.svelte:1166-1176`: underlined text.
- `report/CompareMode.svelte:603`: no retry at all.

**Locations: wording**
- "Retry" instead of "Try again": `lib/character-selector/copy.ts:51` and `lib/planner/copy.ts:137`.
- 19 separate "Try again" strings in `lib/*/copy.ts`.
- Loading text varies: "Loading", "Loading…", "Loading the chart.", "Loading ratings.", "Loading engine…" (`lib/ui/copy.ts:4,14`, `lib/chart/copy.ts:140`, `lib/report/rating-card.ts:11`, `lib/sim/page-copy.ts:121`).

**Locations: busy buttons, two conventions**
- `BUSY_CLASS` (`lib/ui/busy.ts:5`, no spinner): 18 files.
- A spinner plus a busy word: UiSheet:71, AccountDeleteDialog:178-183, ConfirmButtons and four others.
- `Button` has no `busy` prop (`ui/Button.svelte:7-38`).

**What a player sees:** the retry control is a link, a gold button, a grey button or underlined text depending on the page.

**Proposed API**
- `LoadError {message, onRetry?, title?, variant?: 'line'|'panel'}`.
- `Button {busy?}`.
- One shared retry string and one shared loading string.

**Size and risk:** about 8 files, low.

### 12. Filter chips

**Locations.** `ui/Chips` (`chips.css`) has 3 users. Separate chip styles:
- `bis/BisSources.svelte:250-276`: 32 px only, and shown only below 768 px (`:285-289`).
- `lib/guild/filter-styles.ts:5-7`: used by GuildRosterTable and GuildReadiness.
- `ui/select.css:523-558`.
- `chart/ChartLegend.svelte:174-220`: a different control; keep it.

**Accessibility bugs**
- The roster toggles at `guild/GuildRosterTable.svelte:511-530` don't report pressed/unpressed state to screen readers. The sort chips (`:545`) and `GuildReadiness.svelte:136` do.
- "Role: All" and "Class: All" (`:505-510`) cycle on click instead of being dropdowns.
- `guild/dense/GuildRaidRoster.svelte:261-268` is a raw 32 px search box with no focus style and a hard-coded "Find a raider".

**What a player sees:** four chip looks; 32 px tap targets on the BiS phone layout where every other control is 44 px (spec §3.4).

**Proposed API:** `Chips {chips: {value, label, count?, pressed}[], mode}`.
- Delete `filter-styles.ts`.
- The cycles become `Select`.
- The search box becomes a standard `Field`.

**Size and risk:** about 6 files, medium (guild end-to-end test ids).

### 13. Copy buttons

**Locations: primitives and parallel helpers**
- Primitives: `CopyLinkField` (5 users), `CopyForGame` (4), `createCopyFeedback` (8).
- Parallel helpers: `lib/guild/report-copy.svelte.ts:9-34`, `lib/report/copy-link.ts:8-22`, `lib/addon/copy-for-game.ts:40`.

**Locations: inline clipboard calls**

| Site | Revert | On failure |
|---|---|---|
| `planner/SharePanel.svelte:88-95` | stands until the build changes (deliberate, `:58-67`) | silent |
| `setup/SetupPage.svelte:127-135` | 2 s | silent |
| `report/CopyCsv.svelte:19-26` | 1.5 s | "Copy failed" |
| `sim/SimSaveActions.svelte:68-76` | never | not handled |
| `sim/SavedKind.svelte:228-235` | never | not handled |
| `sim/tools/SaveSimForm.svelte:63-70` | never | not handled |
| `sim/RequestDrawer.svelte:113` | no feedback | not handled |
| `guild/GuildReadiness.svelte:108-119` | 2 s | silent |
| `guild/setup/SetupPaste.svelte:18-25` | a named delay | not handled |
| `guild/RaidOfficerRow.svelte:23` | 4 s (re-declares the delay) | not handled |
| `guild/dense/OfficerWeek.svelte:198` | no feedback | unhandled error |
| `guild/dense/GuildRaidRoster.svelte:201-203` | no feedback | unhandled error |

**Also**
- The roster builds CSV by joining with commas and no escaping (`:195-200`).
- The two sim save forms rebuild `CopyLinkField`.

**What a player sees:** some "Copied" labels never clear, and the officer week and roster copy buttons give no confirmation.

**Proposed API:** `createCopyFeedback({revert: 'timed'|'until-change', failedText?})` absorbs every site. The sim forms use `CopyLinkField`; the roster uses the shared CSV writer.

**Size and risk:** about 14 files, low-medium.

### 14. Filter popover, focus trap, inline confirms

**Locations: filter popover shell, three copies.** About 35 lines of behaviour plus 30 lines of CSS each:
- `bis/BisSources.svelte:47-87` (CSS `:196-215`)
- `rankings/RankingsMoreFilters.svelte:60-100` (CSS `:198-220`)
- `report/ReportFilterPopover.svelte:47-83` (CSS `:157-177`)

The report one has no chevron; focus targets and widths differ.

**Locations: focus trap, three copies**
- `ui/Sheet.svelte:47-65`.
- `account/AccountDeleteDialog.svelte:70-93`, copied verbatim.
- `nav/SelectorPanel.svelte:173-190`, with a narrower selector, so Tab can escape the character selector.

**Locations: confirms with their own Escape handling**
- `guild/RaidClearConfirm.svelte:23`
- `guild/RaidCancelConfirm.svelte:29`
- `guild/billing/GuildUpgradeConfirm.svelte:46`
- `checkout/ClaimConfirmBlock.svelte:65`
- `guild/billing/GuildWhoPays.svelte:98`

`account/ConfirmButtons.svelte` already handles this and is used only in account/.

**Proposed API**
- `ui/FilterPopover {buttonLabel, label, testid, width, focusSelector?, panel, stack}`.
- `lib/ui/focus-trap.ts`.
- `ConfirmButtons` moves into `ui/`.

**Size and risk:** about 12 files, medium.

### 15. Section heads

**Locations**
- Hand-written 22 px bold h2s instead of the shared section heading (`Section.svelte:134-139`; spec §3.3):
  - `character/CharacterFightTable.svelte:122`
  - `guild/RaidMatchPicker.svelte:104-113`
  - `home/HomeHowToStart.svelte:17`
  - `bis/BisItinerary.svelte:28`
- 17 px panel titles:
  - `sim/SavedSimLoadFailed.svelte:16`
  - `checkout/PremiumOwnedPanel.svelte:38`
  - `checkout/GuildStepCard.svelte:66`
- 38 raw h3s in about 12 styles.

**What a player sees:** a few section headings per page in a different face and weight.

**Fix:** use `Section` (`level`, `meta`, `focusId`).

**Size and risk:** about 8 files, low.

### 16. Show more and pagination

**Locations: eight patterns**
- `Rankings.svelte:651-684`: a real pager; keep it.
- `MyReports.svelte:178-181`
- `coach/CoachHistoryList.svelte:110-125`
- `character/HistoryFights.svelte:77-85`
- `guild/GuildPlannedRaids.svelte:46-52`: doesn't tell screen readers it expanded.
- `report/SummaryTab.svelte:194-198`: same, and the text is hard-coded.
- `nav/SelectorPanel.svelte:416-424`
- `account/CharacterList.svelte:243-251`

**Proposed API:** `ui/ShowAll {shown, total, expanded, label, onToggle}` and `ui/LoadOlder {label, loading, failed, onLoad}`.

**Size and risk:** about 7 files, low.

### 17. Status pills used instead of Tag

**Locations**
- `guild/RaidChip.svelte:26-28`
- `guild/billing/GuildSeatRow.svelte:174,187,208`
- `guild/RaidMatchPicker.svelte:163`

These use the provenance `.pill` style for states. `Tag` has 52 users, and its own comment (`ui/Tag.svelte:2-3`) says states use Tag.

**Fix:** a Tag tone.

**Size and risk:** 3 files, low.

### 18. Class-coloured character name

**Locations**
- `guild/GuildCharacterName.svelte` has 3 users.
- 22 files pair a crest with an inline class-coloured name, including GuildRosterTable, GuildRaidRoster, ReadinessBoard, FairnessTable, LastNightMovers, BestByRole, SignupRow, ThreatTable, RatingTab and CompareMode.
- The class-colour call appears 59 times in 47 components.
- Feed-row leads are copied too: `home/HomeFeed.svelte:112-116`, `character/CharacterActivity.svelte:41-46`, `guild/overview/GuildActivity.svelte:38-48`.

**Proposed API:** `ui/CharacterName {name, cls, crestSize?, href?}`, and a `lead` prop on `ui/FeedRow`.

**Size and risk:** about 25 files, low.

### 19. Tabs: BossPicker

**Location:** `character/BossPicker.svelte:65-91` is a hand-written tab list with its own resize and fade logic (`:22-47`).
- Arrow keys stop at the ends instead of wrapping.
- No Home/End.
- Arrows only move focus; activation is manual.

The shared behaviour is in `lib/ui/tabs.ts:65`.

**What a player sees:** the boss tabs answer the keyboard differently from every other tab strip.

**Fix:** use `TabStrip` and add a `count` field to its items (`lib/ui/tabs.ts:7-18`).

**Size and risk:** 2 files, low.

### 20. Server-rendered pages

**Locations**
- `api/internal/auth/app_confirm.go` hand-writes its font stack (`:55-56`) and sizes (`:65-76`), and uses a text wordmark with no seal.
- `api/internal/site/page.go:129-137` restates the page header CSS with no test tying it to the web.
- `api/internal/site/chrome.html` is a snapshot of the web header and footer, refreshed by hand with `go test -run TestChromeSnapshot -update`.
- `api/internal/builds/describe.go:16` has a default colour that isn't a design token.

**What a player sees:** the app's sign-in confirmation shows its title in Optima or Georgia, not Marcellus, and with no seal. Shared build pages can show an old header.

**Fix**
- `design/tokens/build.mjs` also writes font and size tokens for Go.
- Regenerate the chrome snapshot in CI.
- Use a brand token for the default colour.

**Size and risk:** 3-4 files, low.

### 21. Discord and Go formatting

**Locations: number grouping**
- `notify/text.go:51` and `builds/describe.go:144` both group thousands.
- `activity/guildsims.go:112` and `activity/guildsim_passed.go:78` don't, so the home activity feed says "simmed 1234 DPS".

**Locations: durations.** `notify/text.go:65` gives "3 min 12 s", a fourth duration style.

**Locations: date layouts**
- "Mon 2 Jan" four times: `activity/guildraids.go:71`, `discordbot/cmd_attendance.go:17`, `discordbot/autocomplete.go:18`, `guilds/readiness.go:285`.
- "3:04 pm" four times: `notify/raid_embed.go:24`, `discordbot/signup_card.go:21`, `guilds/raids_confirmation.go:21`, `guilds/calendar_ics.go:56`.
- `guilds/raids_events.go:55` is the only 24-hour raid time on any surface.

**Locations: other small copies**
- `discordbot/item_embed.go:43,79` repeat `brand.RGB`.
- `discordbot/copy_lookup.go:99` repeats `sheet.QualityWord`.
- Embed colours otherwise all come from the shared brand palette.

**Proposed API:** `api/internal/textfmt` with `Grouped`, `Duration`, `RaidDay`, `RaidClock` and `Location`.

**Size and risk:** about 12 files, low-medium.

### 22. Companion formatters

**Locations: what differs from the web**
- `companion/ui/format.js:47,59` use a 24-hour clock and "Sep".
- `format.js:70` with `copy.js:68-73` says "less than a minute ago" or "today at 20:14", where the web says "just now" or "5 min ago".
- `companion/ui/styles.css:84-154` copies the web's band, tab and panel styles by hand.

**Locations: what is already shared and tested**
- Design tokens are generated (`design/tokens/build.mjs:25`).
- Icons are generated.
- Art and fonts are checked byte-for-byte against the web (`assets_test.go:46-59,90`).

**Fix:** match wording, month spelling and the 12-hour clock to `lib/dates.ts`, held by a shared fixture file, `apps/shared/formats.json`.

**Size and risk:** 2 files plus the fixture, low.

### 23. iOS: one component built twice

**Locations: square game-icon tile, five copies**
- `Raids/Views/RaidComponents.swift:17`
- `Items/ItemIcon.swift:120`
- `Search/SearchRows.swift:131`: falls back to two letters, which breaks the "never a letter square" rule.
- `Guild/Views/GuildIcons.swift:6`: its own image loader.
- Wrappers: `Guides/TiersView.swift:110`, `Sims/SimRowViews.swift:34`.

**Locations: icon frames**
- `Raids/Flow/RaidRootView.swift:174` draws a double border.
- `Widgets/Shared/Views/WidgetChrome.swift:104` uses a different corner radius.

**Locations: error cards, four shapes**
- `Raids/Views/RaidComponents.swift:176`
- `Reports/ReportsListView.swift:233`
- `Bis/BisView.swift:287`
- Two close variants: `Reports/ReportView.swift:99` and `Guild/Views/GuildCommonViews.swift:53`.

**Locations: offline note.** `Search/SearchViews.swift:356` copies `Sources/DesignSystem/Surfaces.swift:53`.

**Locations: date and time formatters**
- "d MMM" five times, with "Sept" in Planner, Reports and Account and "Sep" in Sims and BiS.
- `Common/ClockText.swift:6-11` follows the phone's locale, so a UK phone shows "18:02". Four other places force "6:02 pm".
- Ordinal twice: `Reports/ReportsCopy.swift:53` and `Raids/Views/RaidNightView.swift:265`.

**What a player sees:** a spec with no art shows a crest, a blank tile or letters depending on the tab; error cards change between tabs; a UK phone mixes 12- and 24-hour times.

**Proposed API (DesignSystem):** `GameIcon`, `ErrorCard`, `MessageCard`, `OfflineStrip(style: .card)`, `FSDate`, `FSNumber`.

**Size and risk:** about 25 files, medium.

### 24. Android: one component built twice

**Locations: error regions, four copies (two share the name `ErrorRegion`)**
- `raids/ui/RaidComponents.kt:221`
- `reports/ui/ReportsChrome.kt:101`
- `sims/ui/SimsList.kt:180`
- `search/ui/ItemPage.kt:277`

**Locations: skeletons**
- `SkeletonBar` twice, in different colours: `reports/ui/ReportsChrome.kt:124` and `guild/ui/GuildChrome.kt:98`.
- About 12 private skeleton functions:
  - Raids, BiS, character: RaidsTab:337, BisScreen:319, CharacterPage:421.
  - Rankings, guides, tiers: RankingsBoard:387, GuidesHome:130, TierLists:275, GuideReader:206.
  - Sims and reports: SimsList:189, ReportsList:245, ReportScreen:141.
  - Other: ItemPage:416, FlowLayout:222.

**Locations: icon tiles, four copies**
- `raids/ui/RaidComponents.kt:101`
- `items/ItemIcon.kt:153`
- `search/ui/SearchRows.kt:118`: letter fallback.
- `guides/ui/TierLists.kt:256`

**Locations: formatters**
- About 14 "d MMM" patterns: "Sep" (US locale) in most, "Sept" (UK) in `account/AccountCopy.kt:147` and `activity/ActivityScreen.kt:205`.
- `OfflineStamp.kt:15-17` and `OfflineDataLogic.kt:109-111` are identical.
- The AM-to-am replace is repeated five times.
- Two different ordinal functions: `raids/ui/RaidNight.kt:213` and `reports/logic/ReportsCopy.kt:234`.

**What a player sees:** placeholder bars are different shades on Reports and Guild; "Sept" on Account and "Sep" on Reports.

**Proposed API (core/designsystem):** `ErrorRegion`, `SkeletonBar`, `GameIcon`, `FsDates`, `FsNumbers`.

**Size and risk:** about 30 files, medium.

### 25. Icon URL builder

**Locations: hand-built icon URLs.** The same `dataUrl(build, "icons/<stem>.webp")` pattern is written out 36 times in 33 files:
- Coach: `lib/coach/names.ts:54`.
- Sim: `lib/sim/action-visuals.ts:72`, `action-mark.ts:18`, `spell-info.ts:43`, `buff-names.ts:52`, `spec-icon.ts:22`.
- Guild: `lib/guild/raid-spec.ts:15`, `cover-view.ts:47`, `report-spell-icons.ts:61,63`.
- Other lib: `lib/billing/coach-sample-icons.ts:15`, `lib/report/role-icon.ts:21`, `lib/guides/build-races.ts:75`, `lib/bis/landing-build.ts:52`.
- Components: `ui/SelectSheet.svelte:87`, `report/ReportView.svelte:676`, `sim/SimThisRunSection.svelte:86`, `sim/RunReport.svelte:174`, `guides/GuideSpecBar.astro:33`, and the cluster 1 image sites.

**Locations: inconsistencies**
- `bis/BisSpecRoute.astro:120` hard-codes the path.
- The check that rejects the placeholder "?" icon (`lib/icons/placeholder.ts:7`) runs in only 3 places, so the red "?" can still show elsewhere. That happened on 2026-11-04; see `lib/bis/rotation-icons.test.ts`.

**Proposed API:** `iconUrl(build, stem): string | null` in `lib/ui/assets.ts`, with the placeholder check built in.

**Size and risk:** about 30 files, low. It lands first in lane A.

### 26. Class and spec names, slugs, colours

**Locations: class colour, two sources**
- The tokens (`tokens.css:70-78`).
- The `color` field in `data/classes.json`, read in five places: `ClassCrest.astro:65`, `Planner.svelte:1031`, `character-typeahead-dom.ts:80`, `landing-build.ts:45,73`, `HeaderPanel.astro:89`.

The values match today, but nothing tests that.

**Locations: the nine-class list, 7 copies**
- `lib/ui/primitives.ts:19`
- `lib/report/format.ts:212`
- `lib/class-crest.ts:19`
- `lib/report/planner-link.ts:105`
- `lib/guild/applications-view.ts:151`
- `lib/guild/raid-comp.ts:20`
- `layouts/Base.astro:157-167`

**Locations: name and slug converters**
- Name to slug, three: `lib/report/tree-sizes.ts:11`, `lib/report/format.ts:226`, `lib/report/planner-link.ts:117`.
- Slug to name, three: `lib/sim/spec-label.ts:43`, `lib/tiers/tier-copy.ts:209`, `lib/guides/guide-answer.ts:66`.

**Locations: spec-name functions, four**
- `spec-label.ts:34`
- `home-b-view.ts:128`
- `leveling-plan.ts:80`
- `guild/spec-name.ts:7`

Two different functions share the name `specName`.

**Locations: capitalise helpers, 15 copies**
- Exported: `lib/character-profile.ts:23`.
- Character selector and home: `lib/character-selector/rows.ts:65`, `paste.ts:18`, `lib/home/home-view.ts:118,122`.
- BiS and items: `lib/bis/source-line.ts:26`, `page-model.ts:315`, `copy.ts:500`, `lib/items/source-labels.ts:16`.
- Guides: `lib/guides/key-talents.ts:44`, `race-table.ts:305`, `rotation-lists.ts:71`.
- Guild: `lib/guild/standing.ts:60`, `readiness-view.ts:135`, `loot-council-view.ts:242`.

**Locations: class pickers, four**
- `home/HomeClassPicker.svelte:20-40`
- `not-found/NotFoundClassPicker.svelte:15-40`
- `guides/ClassSpecPicker.astro`
- `HeaderPanel.astro:83-104`

**Fix**
- `spec-label.ts` becomes the one name module.
- Keep one class list in `lib/ui/primitives.ts`.
- Add `lib/text.ts`.
- Drop the colour field from classes.json.
- Merge the pickers in lane D.

**Size and risk:** about 25 files, low.

### 27. Gear-slot constants and average item level

**Locations: slot orders**
- `lib/ui/paperdoll-frame.ts:20-32` is the main one.
- Copies:
  - `lib/planner/types.ts:55-73`
  - `lib/gear/types.ts:10-27`
  - `lib/report/planner-link.ts:28-70`
  - `lib/ui/paperdoll.ts:60-80`

**Locations: other copies**
- Slot-pair maps: `lib/bis/best-in-slot.ts:8-13`, `lib/bis/source-filter.ts:41-45`.
- Weapon-slot sets: `guides/GuideBisPanel.astro:19`, `lib/home/upgrades.ts:28`.

**Locations: average item level, two rules**
- `lib/gear/paperdoll.ts:76` excludes shirt and tabard and refuses a partial average.
- `lib/sim/gear-card.ts:30` averages whatever it knows.

**What a player sees:** with one unknown item, the character page shows no item level and the sim card shows a partial one.

**Fix:** one `averageItemLevel`, and every slot constant exported from `paperdoll-frame.ts`.

**Size and risk:** about 8 files, low.

### 28. Tooltip positioning, disclosures, dead code

**Locations: tooltip positioning, three systems**
- `lib/items/tooltip-position.ts:103,137`
- `lib/ui/popover.ts:90`
- `chart/ChartTooltip.svelte:67-72`, which clamps inline with assumed screen sizes of 1440×900 during server rendering.

**Locations: hand-written `<details>` instead of `ui/Disclosure`**
- `setup/CompanionDownload.svelte:65`
- `setup/SetupSurfaces.svelte:18`
- `report/RaidCooldowns.svelte:267`
- `report/Glossary.svelte:37`
- `sim/RequestDrawer.svelte:123`

**Locations: name clash.** `sim/Disclosure.svelte` has the same name as `ui/Disclosure.svelte`.

**Locations: dead code**
- `SourcePill.astro`: nothing imports it.
- `guild/RaidDisclosure.svelte`: nothing imports it.
- `FactionMark.astro`: only mentioned in comments; confirm before deleting.

**Fix:** the chart tooltip uses `tooltip-position.ts`, the `<details>` sites use `Disclosure`, rename the sim one, and delete the dead files.

**Size and risk:** about 10 files, low-medium.

## 4. Leftovers for the excluded lanes

**lane/talent-tree**
- `planner/TalentCell.svelte:44,68-77,160,280`
- `planner/TreeGrid.svelte:101`
- `planner/GuideBuilds.svelte:60`
- The tooltip placement in `ui/TalentButton.svelte:62-70` and `ui/PointOrder.svelte:46-52,95`
- `guides/GuideBuildsLane.astro:86`
- The tree tabs in `ui/TalentEditor.svelte:113-128`
- The icon URLs in `ui/TalentTrees.svelte:70`, `planner/PlannerOrder.svelte:60` and `lib/planner/editor-trees.ts:41-43,79` (switch to `iconUrl` from lane A)

**lane/header-width**
- `HeaderPanel.astro:83-104`
- The crest styles in `styles/nav.css:260-271,653-664,779-784,1137-1147`
- The header components CheckoutHead, ReportHead, SimCharacterHead, SimToolHead and GuildViewHeader
- `api/internal/site/page.go:129-137` and the chrome snapshot
- The crest and icon in `ui/ControlBar.svelte:156,163`

## 5. Lanes

At most three lanes run at once. Lanes don't push: one integrator merges, runs the checks and pushes. Each lane adds one guard test written like `surface-guard.test.ts`: it scans the source, fails if the pattern appears outside the shared component, and includes a check that it catches the pattern it was written for. Lanes A, D and E change pixels on most pages, so they need the usual ux-designer, wow-player and ai-slop SHIPs on screenshots.

### Wave 1: start now

**Lane A: game icons** (clusters 25, then 1, 6, 9, 27)
- **Touches:**
  - lib: `lib/ui/assets.ts`, the icon-URL builders, `lib/gear/paperdoll.ts`, `lib/sim/gear-card.ts`.
  - New components: `ui/SpellIcon.svelte`, a static `SpellIcon.astro`, `ZoneIcon`, `ItemName`.
  - Callers: the report, coach, sim, guides and chart icon sites; RaidLootItem, ZoneKeyNote, GuildRaidNights, NextRaidRow; HomeFirstRaids, HomeUpgrades, FeedSentence, RaidsCountdown; BisSpecRoute, BisSlotLines.
- **Guard (`icon-guard.test.ts`)** fails on:
  - a hand-built `icons/` URL outside `assets.ts`;
  - an icon `<img>` outside `ui/`;
  - an import of the deleted SpellIcon copies;
  - the item-tooltip hook outside ItemName, ItemRow, ItemHover and PaperdollFrame;
  - `raidZoneIcon(` outside ZoneIcon.
- **Depends on:** nothing.

**Lane B: formatting** (clusters 3, 4, 7, 8, 26)
- **Touches:**
  - New shared modules: `lib/dates.ts`, `lib/format/{number,duration}.ts`, `lib/text.ts`, `lib/report/format.ts`, `spec-label.ts`.
  - Callers: the date, number and duration call sites; the parse sites; RatingsPage; the countdown sites; `UpdatedStamp` and its Svelte twin; the six "Updated" sites.
- **Guard (`format-guard.test.ts`)** fails on:
  - a month-name list outside `dates.ts`;
  - locale or Intl formatting calls outside the formatting modules;
  - rounding or flooring a percentile outside `format.ts`;
  - any exported `durationLabel`;
  - a second class list.
- **Depends on:** the "Sept" or "Sep" decision. It lands before E, which shares GuildRaidRoster.

**Lane C: Go, server pages, companion** (clusters 20, 21, 22, and the `apps/shared/formats.json` fixture)
- **Touches:**
  - Go: a new `api/internal/textfmt`; the Go sites in cluster 21; `app_confirm.go`; `page.go`.
  - Tokens and companion: `design/tokens/build.mjs`; `companion/ui/format.js` and `copy.js`.
  - Shared: `apps/shared/formats.json`.
- **Guards:**
  - A Go test that fails on the date-layout strings and hand-rolled colour packing outside `textfmt` and `brand`.
  - A companion `format.test.js` checked against the shared fixture.
- **Depends on:** the "Sept" or "Sep" decision. It shares no web files with A or B.

### Wave 2

**Lane D: crests, names, sign-up** (clusters 2, 18, 5, 17, and the class pickers from 26)
- **Touches:**
  - Crests: the crest components and styles (nav, typeahead, segmented, a new crest.css); the answer and typeahead code.
  - Rows and names: RankRow, Rankings, RankingsLanding; CharacterName, FeedRow and the 22 name sites.
  - Sign-up and pills: HomeRaidPlaque, RaidBandLine, SignupRow, RaidChip, RaidMatchPicker, GuildSeatRow, WednesdayPanel, `raid-page-controller`; the class pickers.
- **Guard (`crest-guard.test.ts`)** fails on:
  - building a crest path outside `ui/Crest` and `class-crest.ts`;
  - a round crest rule outside `crest.css`;
  - `putSignup(` outside the shared save helper;
  - `.pill` used outside feed rows.
- **Depends on:** `lane/header-width`, lane A, and the sign-up choices decision.

**Lane E: controls and states** (clusters 10, 11, 12, 13, 14, 15, 16, 19, 28)
- **Touches:**
  - Primitives: the ui primitives listed in those clusters plus the new FilterPopover, ShowAll, LoadOlder and ConfirmButtons; `lib/ui` copy, copy-feedback, focus-trap and tabs.
  - Callers: every caller site in those clusters.
- **Guard (`state-guard.test.ts`)** fails on:
  - clipboard calls outside the copy helper;
  - tab lists, dialogs or popovers outside `ui/`;
  - `<details>` outside `Disclosure`;
  - "Try again", "Retry" or "Copied" written outside `lib/ui/copy.ts`;
  - a "No …" muted paragraph in a file that doesn't use EmptyState (a heuristic, with a self-check);
  - a focus selector constant outside the focus-trap module.
- **Depends on:** lanes A and B, and the copy-revert and busy-button decisions. SelectorPanel also overlaps `lane/header-width`.

**Lane F: mobile** (clusters 23, 24)
- **Touches:** the iOS DesignSystem plus about 25 app files, and Android `core/designsystem` plus about 30 app files.
- **Guards:** an iOS test and an Android test that each fail on a date formatter outside the formats module, or a second definition of ErrorRegion, SkeletonBar or GameIcon. Both check against the shared fixture.
- **Depends on:** lane C's fixture and the date and clock decisions. It shares no web files, so it can take lane C's slot as soon as C lands.

## 6. Decisions you need to make

1. **Month spelling:** "Sept/June/July" (the current house date, `lib/dates.ts:1`) or "Sep/Jun/Jul" (guides, BiS, both apps, the companion). Blocks lanes B, C and F.
2. **Clock:** a fixed lower-case 12-hour "8 pm" everywhere (including `dates.ts:111` and the iOS clock), or follow the device's locale.
3. **Home sign-up:** keep three choices ("I'm in", "Late", "Can't come") or match the guild's four (In, Late, Bench, Out).
4. **"Copied":** clear after a short time, or stay until the copied thing changes (the planner's deliberate rule, `SharePanel.svelte:58-67`).
5. **Busy buttons:** dim with no label change (18 files today), or a spinner with a busy word (7 files).
6. **Parse figures:** this plan assumes floored numbers, matching the colour rule (Warcraft Logs' rule, `format.ts:335-336`).

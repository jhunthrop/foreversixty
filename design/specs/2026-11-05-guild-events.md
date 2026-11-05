# Guild events: dungeon runs and guild events before raids open (experience spec)

Author: ux-designer, 2026-11-05. Round-1 guild audit fix WP-1 (the wow-player's NOT SHIPPABLE reason one: a leveling guild can post nothing until 9 Dec).
Amends `design/specs/2026-10-26-guild-raids.md` §3.1, §3.5, §6.1, §6.2 and ruling R6. Everything not named here is that spec, unchanged: sign-ups, the raid page, the band line, the poll, Discord, the 12-ahead cap.
Code read: `api/internal/guilds/{raids_zones,raids_validate,raids_view,raids_match_candidates}.go`, `api/internal/db/migrations/0050_guild_raids.up.sql`, `data/curated/raids/first-tier.json`, `data/builds/1.60.1.70790/loot.json` (`sources`, kind `dungeon`, `opens`), `web/src/components/guild/RaidForm.svelte`, `web/src/components/ui/Select.svelte`.
No new control, no new token, **no database migration**.

## 0. What the schema allows (verified)

- `guild_raids.zone` is `text not null` with no check constraint (`0050_guild_raids.up.sql`; no later migration adds one). Any key the API accepts can be stored.
- There is no size column. Size comes from the zone (`ZoneView.size`, `Zone.Size`), read at render time.
- `duration_minutes` is checked `between 60 and 240` in the database. Events keep 1 h to 4 h.
- `status`, sign-ups, roles and notes are zone-agnostic.
- Log matching reads `Zone.LogZone`; a zone with none is never auto-matched (`raids_match_candidates.go` line 53), and an officer can still match a log by hand.

So the whole feature is a longer curated zone list plus per-zone facts, all in Go and curated JSON:
- Dungeons get one key each with a sourced size.
- A guild event's size is chosen by the officer by picking one of six keys (`event:guild`, `event:guild-5`, `-10`, `-15`, `-20`, `-40`). That is how an officer-chosen size fits without a size column.
- Raids gain an open date.

## 1. The one question

| Viewer | Question | Answer, no scroll at 360 and 1280 |
|---|---|---|
| Officer, week one | "Can I get a Deadmines run or a guild meet-up on the calendar?" | `Plan an event` on the Raids tab (and as the Before raids open card's button, spec A §3.1); the form's first field lists dungeons and `Guild event` above raids. |
| Member | "What's on, and am I in?" | The `Planned` rows and the band's raid line, unchanged, now holding dungeon runs and events too. |

## 2. Reference, and what we borrow

- **The in-game Calendar's "Create Event"** (Blizzard_Calendar): the event type list is Raid, Dungeon, PvP, Meeting, Other, and Raid or Dungeon opens an instance list. Borrowed: one type-then-instance choice, and a free "Meeting/Other" type. Not borrowed: PvP (no data model) and the calendar grid.
- **Raid-Helper's templates** for 5-man and 10-man runs: the role count against a small cap (`1 tank · 1 healer · 3 DPS`). Borrowed: the same count line, now at 5.

## 3. Layout and behaviour

### 3.1 The zone catalogue (`raids_zones.go`, `data/curated/raids/events.json`)

`Zone` gains `Kind` (`raid`, `dungeon`, `event`) and `Opens *time.Time`. `curatedZones` holds, in this order:

1. **Dungeons**, kind `dungeon`, key `dungeon:{slug}`. The starting list is every `loot.json` source of kind `dungeon` with no `opens` field (open at launch), minus the event variants (`:scourge-invasion`, `:tier-0-5`). Then the lane checks each against the client's own tables for the active build (wago.tools) and records the build number and rows in its report:
   - **Is it a dungeon?** `Map.InstanceType` must be 1 (party) or 2 (raid). Anything else (an arena is 4, a world zone 0) is dropped. Ruins of Lordaeron (`zone_id` 16611 in `loot.json`, Wowhead-sourced, its "bosses" are skeletons and ghouls) and Excavation Site: Wetlands are checked this way; neither ships until it passes.
   - **Wings are separate runs.** A group signs up for one wing, so each `LFGDungeons` row is its own key: Scarlet Monastery `Graveyard`, `Library`, `Armory`, `Cathedral`; Dire Maul `East`, `West`, `North`; Stratholme `Live` and `Undead`; Blackrock Spire `Lower` and `Upper`. Names read `Scarlet Monastery: Library`, `Dire Maul: East`, `Stratholme: Undead`, `Lower Blackrock Spire`, `Upper Blackrock Spire`. Keys: `dungeon:scarlet-monastery-library`, `dungeon:dire-maul-east`, `dungeon:stratholme-undead`, and so on.
   - **Size**: `MapDifficulty.MaxPlayers` for the map. Expected 5 for the five-mans and 10 for Upper Blackrock Spire; the table, not this spec, is the source.
   - **Level range**: `LFGDungeons.MinLevel` and `MaxLevel` for the row. Shown as `17 to 26`.
   - **Key needed**: a row carries `key_item` only when the lane verifies the lock on classic-db (expected: Upper Blackrock Spire, Seal of Ascension; Scholomance, Skeleton Key; Stratholme: Undead by the service entrance, Key to the City). Unverified rows carry none.
   - A row a table does not answer keeps `size: null` (`{n} signed up`, no bar, as R6 rules for Barrow Deeps) or no level range.
   - `LogZone: nil` this round. Dungeon nights are matched to logs by hand.
   - Dungeons sort by `MinLevel`, then `MaxLevel`, then name, so the list reads in the order a guild levels through them.
2. **Guild event**, kind `event`: six keys, one name.

| Key | Name shown | Size |
|---|---|---|
| `event:guild` | `Guild event` | null (no cap) |
| `event:guild-5` | `Guild event` | 5 |
| `event:guild-10` | `Guild event` | 10 |
| `event:guild-15` | `Guild event` | 15 |
| `event:guild-20` | `Guild event` | 20 |
| `event:guild-40` | `Guild event` | 40 |

3. **Raids**, kind `raid`: today's three, each with `Opens` = `raids-1` from `data/curated/phases.json` (`2026-12-09T00:00:00Z`).

`data/curated/raids/events.json` carries the dungeon and event rows (key, name, kind, size, min_level, max_level, key_item, and the client build each fact was read from); `raids_zones_test.go` pins it to the Go slice, as it pins `first-tier.json` today. The API's `zones` array gains `kind`, `opens` (RFC 3339 or null), `min_level`, `max_level` and `key_item` (name or null).

### 3.2 Validation (`raids_validate.go`)

- `zone` must be a curated key (today's rule; the list is longer).
- A raid zone with `Opens` after `starts_at`: 422, field `starts_at`, `{Raid} opens {Wed 9 Dec}. Pick that day or later.` Planning the opening night ahead of time is allowed and expected.
- An event's empty title saves as `Guild event`; the form asks for a title for events (§3.3), so this is a fallback only.
- The 12-ahead cap counts raids only (kind `raid`; the count query filters `zone like 'raid:%'`). A guild can post as many dungeon runs and events as it likes, each still within a year.
- Nothing else changes: duration 1 h to 4 h, `Also plan the next 3 weeks` works for every kind.

### 3.3 The form (`RaidForm.svelte`)

Field 1 is renamed and regrouped. Every other field and its order is the raid spec's §3.5.

1. **`What` select** (`Select.svelte`, 52 px, its native `optgroup` support), placeholder `Choose a dungeon, event or raid`. Groups in this order:
   - `Guild event`: one option, `Guild event`.
   - `Dungeons`: in level order, name, size, then level range, each part only when known: `The Deadmines · 5 · 17 to 26`, `Upper Blackrock Spire · 10 · 58 to 60`.
   - `Raids`: `Barrow Deeps · opens 9 Dec`, `Hyjal Summit · opens 9 Dec`, `Onyxia's Lair · 40 · opens 9 Dec`. From 9 Dec the `· opens` suffix is gone. Raids are selectable; the date rule above applies.
   - Default: the guild's last planned key (its event size included); first time, the placeholder.
   - A dungeon with `key_item`: under the select, 13 px muted `Needs a key: {Seal of Ascension}.` (the key's name as the shared item link, quality colour, tooltip). The same line goes into the event page's band under the when line. A key that opens only one of two ways in carries `key_item.door` and reads `{Key to the City} opens the service entrance.` instead (Stratholme: Undead; groups also enter through the main gate, which has no lock) (amended 2026-11-05, guild round 1).
2. **`Group size` select**, only when `Guild event` is chosen, directly under `What` (same column, half width from 1024): `No limit`, `5`, `10`, `15`, `20`, `40`; default `No limit`, or the last event's size. The pair (`Guild event`, size) maps to the six keys of §3.1. Hidden and not sent for any other choice.
3. **`Title`**: for a guild event the placeholder is `What is it? (a PvP premade, a quest group)` and an empty title shows the field error `Give the event a title.` on submit. For dungeons and raids the placeholder stays the zone name and empty still saves as the zone name.
4. **`Date`**: when a raid is chosen before its open day, the input's `min` is the open date and, under the field, 13 px muted: `{Onyxia's Lair} opens {Wed 9 Dec}.` A date before it gets the §3.2 error.

Labels: the form's h2 and button change from raid words to event words (§5). `Post to Discord`, notes, repeat and time zone are unchanged.

### 3.4 The Raids tab and the rows

- The tab keeps its name, `Raids`; it already holds raid nights and planned nights. The header button becomes `Plan an event` (`button('primary','md')`; gated rules in spec C §3.2).
- A planned row is the raid spec's row. The title cell's second line (zone, shown when the title differs) reads the zone name with size for events: `Guild event · 10`. Count `3/5` with the bar when sized; `3 signed up` without.
- The empty `Planned` lines (raid spec §6.1) become: member `Nothing planned yet. Officers post dungeon runs, events and raids here.`; officer `Nothing planned yet. Post a dungeon run or a guild event and the guild can sign up.`
- The band raid line and its `{Zone} →` link work unchanged for every kind.
- Readiness's `Next raid` chip and the needs line's `before {Weekday}` read the next planned **raid** only (kind `raid`), never a dungeon run; before 9 Dec they fall back to `before the next raid`.

### 3.5 The event page

`GuildRaid.svelte` unchanged in layout. Words that say `raid` follow §5 by kind: eyebrow `RAID` / `DUNGEON` / `EVENT`; `Cancel raid` / `Cancel run` / `Cancel event`; `This raid has ended.` / `This run has ended.` / `This event has ended.`; cancel confirm `Cancel this {raid|run|event}? Members see it as cancelled.`; `Logged as a raid night.` stays raid-only. One copy function per string, keyed by `zone.kind`, in `raid-copy.ts`.

**A guild event is a plain list** (amended 2026-11-05, guild round 3: a guild meeting showed the Tanks / Healers / Melee / Ranged grid and a class row). One rule, `kindCountsRoles(kind)` in `raid-copy.ts`: kind `event` counts no roles; `raid` and `dungeon` do (a run needs a tank and a healer). For kind `event`:
- **Band**: no class row. The count line has no role words: `12 of 20 signed up`, then ` · 2 late` only when someone is late.
- **Your sign-up**: each row is who (crest, name, the class alone: `Priest`), the segments `In` · `Late` · `Bench` · `Out`, and the note. No spec picker, no `Pick a spec first.`; a press saves at once and sends no spec (the character's own role rides along unseen). Desktop grid `minmax(180px,1fr) auto minmax(200px,1.4fr)`; phone: who, segments, note.
- **Sign-ups**: no role grid. One list grouped by state, `In {n}`, `Late {n}`, `Bench {n}`, `Out {n}` (the 11 px `label` heads), side by side from 1280 (four columns), two from 768, stacked below; an empty group is left out. Each row is crest, name in class colour, and the class (`Priest`), never a spec or a role word. `Not signed up` and the officer selects are unchanged.
- **Planned row**: the breakdown is empty, or `2 late`; the viewer chip reads `In`, `Late` with no role.
- **Band raid line** on the guild page: `Sunday 8 pm CST, you are In, 12 of 20 signed up.`, never `In as DPS`.
- **Copy the sign-up list**: no `Classes:` line and no role groups; `In: …`, then `Late`, `Bench`, `Out`.
- Not this round: the Discord post's role fields and the bot's `Which spec tonight?` still follow raid rules on an event (Go, `notify` and `raids_discord*`); a follow-up.

### 3.6 Discord

Same `raid_posted` and `raid_changed` embeds (raid spec §6.2), same channel, same box. The title is `{Title} · Thursday 12 November`. `Signed up` reads `3 of 5`, or `None yet`. No new event type, no thumbnail change.

## 4. States

| Element | Default | Hover | Focus | Disabled | Loading | Empty | Error |
|---|---|---|---|---|---|---|---|
| `What` select | placeholder | border `--border-warm` | 2 px gold | gated officer (spec C) | the form opens only after `zones` loads | n/a | `Choose a dungeon, event or raid.` |
| `Group size` | `No limit` | as Select | as Select | as `What` | n/a | n/a | n/a |
| Title (event) | placeholder | as field | as field | as `What` | n/a | n/a | `Give the event a title.` |
| Date (raid before open) | `min` set, helper line | as field | as field | as `What` | n/a | n/a | `{Raid} opens {date}. Pick that day or later.` |
| `Plan an event` | primary | primary hover | 2 px gold | gated: secondary, disabled (spec C) | n/a | n/a | n/a |

Touch and keyboard as the raid spec: every control 44 px on phone, native selects keep the phone pickers.

## 5. Copy, verbatim

| Where | String |
|---|---|
| Tab button | `Plan an event` |
| Form label | `What`; `Group size`; options `No limit`, `5`, `10`, `15`, `20`, `40` |
| Placeholders | `Choose a dungeon, event or raid`; `What is it? (a PvP premade, a quest group)` |
| Groups | `Guild event`, `Dungeons`, `Raids` |
| Option suffixes | ` · {size}`, ` · {min} to {max}`, ` · opens {9 Dec}` |
| Key note | `Needs a key: {Item}.`; with a keyless way in, `{Item} opens {door}.` (`Key to the City opens the service entrance.`) |
| Wing names | `Scarlet Monastery: Graveyard`, `Scarlet Monastery: Library`, `Scarlet Monastery: Armory`, `Scarlet Monastery: Cathedral`, `Dire Maul: East`, `Dire Maul: West`, `Dire Maul: North`, `Stratholme: Live`, `Stratholme: Undead`, `Lower Blackrock Spire`, `Upper Blackrock Spire` |
| Helper | `{Raid} opens {Wed 9 Dec}.` |
| Errors | `Choose a dungeon, event or raid.`; `Give the event a title.`; `{Raid} opens {Wed 9 Dec}. Pick that day or later.` |
| Submit | `Post the event` (`Posting…`); edit `Save changes` (`Saving…`) |
| Empty Planned | member `Nothing planned yet. Officers post dungeon runs, events and raids here.`; officer `Nothing planned yet. Post a dungeon run or a guild event and the guild can sign up.` |
| Eyebrow | `RAID`, `DUNGEON`, `EVENT` |
| Officer row | `Cancel raid`, `Cancel run`, `Cancel event` |
| Ended | `This raid has ended.`, `This run has ended.`, `This event has ended.` |
| Cancelled | `This raid was cancelled.`, `This run was cancelled.`, `This event was cancelled.` |
| Cancel confirm | `Cancel this {raid|run|event}? Members see it as cancelled.`; `Cancel the {raid|run|event}` |
| Row zone line | `Guild event · {size}`; the event page's h1 follows a title that is not the zone's name with the same words, 15 px muted in their own case, never the h1's uppercase (amended 2026-11-05, guild round 2) |
| Changelog | `Officers can now post dungeon runs and guild events as well as raids. Members sign up the same way.` |

Every other raid-spec string is unchanged. `Choose a raid.` and `Choose a raid` are replaced by the strings above.

## 6. Hierarchy

The Raids tab's one fill stays the header button. In the form the `What` select is first and focused on open. Dungeons and events are not visually ranked above raids beyond their group order.

## 7. Phone and desktop

As the raid spec §7. `Group size` sits under `What` at full width below 1024. Option labels never truncate in the native picker.

## 8. Performance

The `zones` array grows from 3 to about 40 rows inside the existing `GET …/raids` answer. No new request, image or font. Lighthouse budgets unchanged.

## 9. Left out, on purpose

Soft reserves, a composition tool, PvP events, auto log matching for dungeons, and per-event role caps. Each needs either a migration or its own design review.

## 10. Acceptance screenshots (390, 1440; 2000 for row 1)

| # | State |
|---|---|
| 1 | Raids tab, officer, a Deadmines run (`2/5`) and a guild event with no limit planned |
| 2 | Form: `What` open (native picker on phone) showing level order and wings, `Guild event` chosen with `Group size` shown, title error, a keyed dungeon with its note |
| 3 | Form: Onyxia's Lair chosen before 9 Dec (clock stubbed), helper line and the date error |
| 4 | Event page for a dungeon run: eyebrow `DUNGEON`, `Cancel run`, sign-ups at 5 |
| 5 | Band raid line pointing at a guild event |
| 6 | Discord embed for a dungeon run in the desktop client |
| 7 | Member Raids tab, empty Planned |

Tests: `raids_zones_test.go` pins `events.json`; no zone in it has `Map.InstanceType` other than 1 or 2; thirteen dungeon runs and events post past the raid cap; a POST with `event:guild-10` answers `size: 10`; a POST of `raid:onyxias-lair` on 20 Nov answers 422 on `starts_at`; a dungeon run never moves the Readiness `Next raid` chip.

# Guild tab: information density verdict (2026-11-11)

Reviewed: spec 2026-11-10-guild-tab.md, GuildTabPresenter.swift, GuildTabViews.swift, the full-length renders
(ios-full-guild-tab, -officer, -leveler), the guild home fixture and home.go. Android was not walked screen by screen; it
renders the same sections from the same inputs (GuildTabSections.kt, GuildTabParts.kt), so findings 1 to 8 apply to both.
Not reviewed: the Android captures.

Job of the tab: run my guild from my phone, the raid first. The raid block, Your job, Who is coming and Before you pull are
dense and good. The tab's own sections (head, Roster, Loot, Logs, Progression) are not.

## Verdict: THIN

The guild sections read `GET /v1/guilds/{id}/home` and use about a tenth of it. Two panels (Loot, Progression) are one row
holding one phrase, and the roster is a list of names with a clock. The spec's own mapping table limited the tab to what the
home serves, and the presenter drops most of that. Fix the MUSTs below and re-submit.

## MUST

1. **Loot is a count with nothing behind it.** The panel is a title, a row `Last raid: 4 pieces`, and a link. A raider's
   question is who got what. Fix: list the pieces (item icon in quality colour, item name, winner in class colour), the
   first five, `Loot history ›` for the rest. Source: `/v1/guilds/{id}/raids/{raid}` loot rows for `badge.last_raid_loot.raid_id`
   (the guild loot-history call the web Loot page reads); only the count is on `/raid-badge`. If the per-piece call is not
   already generated for the apps, it goes under needs data (#1); until then the panel is the count folded into the Loot
   history row of Pages (as the leveler render already does), not its own panel.
2. **Roster rows carry no figures.** `Recently on` rows show level, spec, `Seen 12 min ago` and stop, though `roster[]` has
   them. Fix: replace the level in the sub line with item level and add the trail's second figure. Concretely
   `60 · Protection Warrior` becomes `Protection Warrior · 58 ilvl`, trail `Seen 12 min ago`. Source:
   `roster[].item_level`, `roster[].addon.gear_gap` (show `gap 3` in the warning tone only when > 0), join `members[]` and
   `roster[]` on `character_key`. Withheld by consent: the field is absent, so the row shows what it has.
3. **Attendance and parse are served and never shown.** Fix: on each Roster row (at 60) a trail `6 of 8 nights`
   (`roster[].attendance.present/nights`) and, as the sub line's second item, the best parse
   (`roster[].best_parse.percentile` with `encounter`), tap opening the report (`best_parse.report_id`, `fight_index`;
   `openReport` already exists). Gives a raider the percentile and the trend context the roster lacks.
4. **Close to 60 duplicates Recently on.** In the leveler render Thornwick (57) is in both lists, one under the other.
   Fix: drop members already shown in Recently on from Close to 60 and take the next five (presenter `close(_:excluding:)`
   gets the `recent` ids), or sort Close to 60 by `seen_at` second. Source: `members[]`.
5. **Progression is one row and a phrase.** `First tier / 1 of 1 down · 6 pulls this tier` is the tier's two numbers with
   no boss, no date, no bar. Fix: a segmented bar of `down` over `named_encounters`, the tier's raids as chips
   (`summary.tier.raids[]`), and `updated_at` as the one Updated stamp. Source: `summary.tier.{name,raids,named_encounters,down,pulls_total}`,
   `summary.updated_at`. Per-boss kills with dates: needs data (#2, already in the spec's own list).
6. **The guild head says nothing at 60.** At 60 it is a crest, a name and `38 members · Alliance · US PvP`; the spread
   figures (31 at 60, 4 at 55 to 59, you) only appear below 60, and the head is then separated from Roster by the whole raid
   block. Fix: at 60 show the same three figures from the same data, with `raiders` and, for officers,
   `waiting_for_approval` / `below_rating_floor` (`summary.raiders`, `summary.waiting_for_approval`, `summary.below_rating_floor`),
   and `counts.not_seen` when > 0. Source: `members[].level`, `summary.*`, `counts.*`. This also removes the `31 of 38 at 60`
   duplicate on the Roster title.

## SHOULD

7. **Recent logs lose their clean-night signal.** `GuildLogLine.clean` is computed and never drawn. Fix: render `7/8 kills`
   in the kill tone when clean, plain otherwise, and add `fight_count` pulls as the trail's second line. Also three logs
   only; `reports[]` is 20 per page, so a `Logs ›` action (the site's guild Logs tab) is missing. Source: `reports[].kill_count`,
   `fight_count`, `duration_ms`.
8. **Officer block has no roster health.** Applications says `1 waiting` and nothing about who is not showing up. Fix: a
   row `Not synced` / `Not seen` with `counts.not_seen`, `counts.in_game_no_account`, `counts.waiting`, tapping the roster
   page; for the raid, the addon's not-synced names already appear under Who is coming (good, keep it there, do not repeat).
9. **Next raid on the Guild tab below 60 lacks the raid's tail.** The next raid's `size` and `status` are served
   (`next_raid.size`, `status`); 40-man shows, status does not (only matters for `cancelled` or `live`, show then).
10. **Assignments has no your-slot jump.** Slots list names but a raider cannot tell which encounter has them; the Your job
    section above covers it, so only a gold left rule on rows holding `names[].own` is needed. Source: `assignments.encounters[].slots[].names[].own`.

## Phones at 390
Findings 2 and 3 add a second figure to the trail: keep the sub line `spec · 58 ilvl` and let the trail hold one value
(seen, or nights) with the parse moving to the row's tap-through, never a third column. Finding 5's chips wrap to two lines.
Finding 6 reuses the existing three-figure row (fits at 390 and at ax5, where it already stacks).

## Needs data
1. Per-piece loot for the last raid on the app side (item, winner) if no generated client call exists; `/raid-badge` serves
   only `last_raid_loot.count` and `raid_id`.
2. Guild kills by boss with dates (already spec needs-data #2).
3. An account-wide applications list (already spec needs-data #1); nothing else new is requested.

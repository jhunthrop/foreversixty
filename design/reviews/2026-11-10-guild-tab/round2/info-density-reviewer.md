# Guild tab: information density, round 2 (2026-11-11)

Reviewed: GuildTabPresenter.swift, the full-length iOS renders (ios-full-guild-tab-officer, -leveler), the loot nights
fixture, home.go, the round-1 response. Android read through GuildTabSections.kt's shared inputs only; its captures were
not walked.

## Verdict: DENSE

Round-1 MUSTs 1, 2, 4 and 6 are closed and 5 is mostly closed:
- Loot lists the night's drops with icon, quality colour, winner in class colour, `Not awarded yet`, the day. (1)
- Recently on rows carry item level; officers get Missing raids with `2 of 8 nights`. (2, part of 3)
- Thornwick no longer appears twice. (4)
- The spread shows at every level and the head carries `US PvP`. (6)
- Progression names the raid and the `down` and `pulls` figures. (5)

Nothing on the tab is a dead panel any more. What follows is still served and unused; none of it blocks.

## Still missing that the API serves

1. **Best parse on Recently on rows.** `roster[].best_parse.{percentile,encounter,report_id,fight_index}` is served and
   never read by the presenter. Fix: the second item of the sub line (`Frost Mage · 61 ilvl · 94`) at 60 only, row tap
   opens the report (`openReport` exists). Trail keeps the one value, so 390 holds.
2. **Logs lose the clean-night signal and the full list.** `GuildLogLine.clean` is computed and not drawn; the render shows
   `7/8 kills` plain. Fix: kill tone on `kills` when clean, and a `Logs ›` action on the section title for the rest
   (`reports[]` is 20 per page, three shown).
3. **Progression has no bar.** `summary.tier.down` over `named_encounters` and `summary.updated_at` are served. Fix: a thin
   segmented bar under the row, the one Updated stamp at the section's end. Per-boss kills with dates stay needs data.
4. **Roster health for officers.** `counts.not_seen`, `counts.in_game_no_account`, `summary.waiting_for_approval`,
   `summary.below_rating_floor` are served; Officer shows only `1 waiting`. Fix: one row `Not seen` / `No account` with the
   counts, tapping the roster page. Show only non-zero counts.
5. **Assignments: your slot.** `assignments.encounters[].slots[].names[].own` is served and the gold name already marks it;
   add the gold left rule on the row so it reads without scanning (the render has no row marker).

## Phones at 390
All five fit: 1 and 2 add to an existing line or title, 3 is one bar, 4 and 5 reuse rows. Nothing needs a drawer.

## Needs data
1. Guild kills by boss with dates (spec needs-data #2).
2. An account-wide list of your applications (spec needs-data #1).
3. Per-raider deaths and avoidable damage.

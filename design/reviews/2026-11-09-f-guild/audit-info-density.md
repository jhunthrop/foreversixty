# Audit (round 0): info-density-reviewer on the current guild pages

Captures: audit/build-<state>-<width>.png.

## Per page (condensed)
- Overview, officer: THIN. Unshown: summary.below_rating_floor, tier/named_encounters_down/pulls_this_tier, reports[], roster rating and attendance_30, addon_snapshot.received_at, readiness median and failing. Next raid is four bare numbers. Add a KPI strip, Last night, Who's short, class mix.
- Overview, member: THIN. standing.rank_by_item_level / same_spec_count unshown; item_level_delta, rating, best_parse, gear_gap gain unshown. Checklist sub-lines are explainer copy.
- Roster: THIN ("the premium page and the one the owner will hit first"). Every GuildRosterRow field but item level and raids is unshown: rating (7), best_parse, attendance_30, professions, addon.build, addon.gear_gap, level, logged_at, last_report_at, is_main/account_key, sim_export. No header sort, no role summary.
- Raids list: THIN. top_dps/top_hps/raiders/deaths/duration unshown; trash rows; one night.
- Raid planner: DENSE WITH ADDITIONS. Sign-up rows lack roster joins; note and also[] unrendered; cover missing; assignment prose.
- Report: DENSE WITH ADDITIONS. Full meter, deaths grouped by ability, trash collapsed, apology panel removed.
- Progression: THIN. pulls_by_night, best_kill_ms, best_by_role, deaths_per_pull, unnamed unshown.
- Readiness: THIN. Needs the full matrix with every GuildReadinessRow field, the ready people included.
- Loot: DENSE WITH ADDITIONS. replaces_item icon, ilvl_delta, attendance, items_won buried; make drops a table and candidates a comparison table.
- Applications: THIN. Profile: THIN. Public: THIN. Recruit: DENSE WITH ADDITIONS. Settings: DENSE.

## Cross-cutting
Dead rails and empty bottom halves; no sortable headers or CSV; no tooltips; no context on numbers (median, percentile, window, delta); explainer copy to remove ("Private to you", "Same units as the BiS page", "Best in slot or no upgrades, enchanted, stocked, every point spent", the "no character sheet" apology); 390 roster opens on filter chips.

## Ranked adds (reviewer's order)
1 Roster as a real table. 2 Readiness matrix. 3 Overview KPI strip + Last night + Who's short + class mix. 4 Report full meter. 5 Progression per-boss rows. 6 Loot table with comparison columns. 7 Member profile. 8 Public/recruit. 9 Raids night table. 10 Applications.

**Overall: THIN.**

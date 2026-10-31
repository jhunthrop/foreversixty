# Player verdict: fight list kill counts (2026-10-31)

Reviewed: sample-fights-1440/390, folded-zone-row, no-mobs-died, old-format-no-count, plus the lane report.

VERDICT: SHIP

## Blockers
None.

## What reads true
- "177 mobs died · your group killed 26" is believable. 675 / 78 on the fold sums the pulls, and old reports show no count instead of a made-up one. That is the right call.
- Hiding the dagger when only outsiders died is right. A raider reads the dagger as "did we wipe people", and a stranger's death in range is not ours. Old reports still show the all-deaths dagger (6† on the fold against 2† now), which is an honest fallback.
- "Died" for the mob figure and "killed" for the group figure is the right pairing.

## Shoulds
1. "your group killed 0" under "115 mobs died" (3:26 pull) is noise on a dozen rows. Say "none by your group", or drop the second half when it is 0 and let the tooltip explain.
2. "1 mob died · your group killed 1" says the same thing twice. When the two are equal, say "1 mob killed" or "all killed by your group".
3. A 115-mob "Trash" pull is an open-world or city crowd, not trash. Players will mock the label. Name the zone or call it "Open world" when most deaths are outsiders.
4. The dagger's "among everyone in range" tooltip contradicts a group-scoped number. Say "among your group", and give the all-in-range figure only on hover if it is worth keeping.
5. The panel header reads "0 pulls" on the Zephras Isle sample with 79 trash pulls, while the boss sample reads "2 pulls". Define what it counts or show the real number. It was there before this change but it sits next to the new counts.
6. "No mobs died" on a pull is fine, but "Trash · 6.4s" with no kills looks like a bug. Consider "Pulled, no kills" or hiding the row.
7. The 390 summary line wraps well, but the long "Trash · 12:43 · 177 mobs died · your group killed 26" header is heavy. Abbreviate on phone, e.g. "177 died · 26 yours".

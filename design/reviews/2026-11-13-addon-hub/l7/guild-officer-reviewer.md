# L7 Officer tab, round 2: guild-officer-reviewer

Checked: round-1-fixes.md, the redrawn l7-raid-day.png, `Officer.coverage`, `Officer.readiness` and the Guild tab's prep code.

## Verdict: SHIP

1. Round 1 MUST 1 resolved: the `Not covered` row from ReadyCheck's coverage over the sign-ups shows a gap before anyone is grouped.
2. Round 1 MUST 2 resolved: a raider's own pre-pull look stays `/fs ready`, the chip line and Guild's Your prep (spec 4.6 and 6).
3. Round 1 items 3 to 6 fixed: the ready-check row lists who lacked what, the muted `-` separates unknown from `Best` and `0`, the roster pages, `Raiders online (24)`, whole-name cuts.
4. SHOULD: a click on the ready-check row could post the missing names to raid chat.
5. SHOULD: approvals are a count with no link (as the old flyout).
6. NICE: Won covers only your own `RaidDrops` record; the hover says so.

Nothing an officer could do in the old flyout or group panel is missing: sorting, the online filter, the push confirm, claim state, approvals and the loot list are rebuilt, and the coverage row replaces the old coverage region.

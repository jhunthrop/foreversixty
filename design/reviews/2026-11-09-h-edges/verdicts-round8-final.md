# Round 8 verdicts: final (captures: round8/)

- ux-designer: VERDICT: SHIP
- wow-player: VERDICT: SHIP
- ai-slop-reviewer: VERDICT: PASS
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS
- guild-officer-reviewer (member profile, night report): VERDICT: WORTH PAYING FOR

Scope ruling (coordinator): the painted ground and the blue band are the owner's round-5 art, out of scope for ai-slop.

## Routed to the coordinator
- The shared attendance label "Joined from the bench or without a sign-up": web/src/lib/guild/attendance-copy.ts line 135 (attendanceCopy.attendance.groups.joined). Lane f-guild changes that file, so this lane left it; the report's officer names read lane copy ("Not in the log", "Late", "Joined").
- RaidLootAward in web/src/lib/guild/attendance-contract.ts (also an f-guild file) should take AwardRecord once both lanes land; ReportLootRow reads it as RaidLootAward & Partial<AwardRecord> until then.

## Needs data (no table or column holds it)
- A raider's trial status and rank history; a guild attendance or at-risk threshold; officer actions from the profile (message, bench, demote); reputation standing and vendor price on a BiS source; week over week per raider (moves.previous needs a second logged night).

## Follow-ups named by reviewers
- "Also an upgrade for N raiders" on each report award (guild-officer); DPS share bars, sort, CSV and fight-log links (density); label reputation and vendor sources (guild-officer, wow-player); the ~60 px guild footer parity drift (site-wide, coordinator).

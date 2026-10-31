# Reports and trash: review verdicts (2026-10-31)

Both reviewers read the captures in this folder.

## ux-designer: SHIP (round 2)

Round 1 was BLOCK on one finding: the new themed scrollbar exposed a stray vertical thumb at the
end of the report's mode tab row. The fix was a `scrollbar-none` utility plus `overflow-y-hidden`
on the mode row and on the planner's TreeTabs. Round 2 confirmed the thumb is gone at 1440 and
390.

Shoulds fixed in this lane:
- "PER PULL ›" sits on its own line under the spec.
- The trash entry's title is in the boss-row colour.
- The capture's session carries a character.

Shoulds left for the shared components:
- the empty state's Upload button (shared EmptyState)
- the 390 row wrap (shared ReportRow)
- the crest in the night Damage table (pre-existing in "All boss pulls")

## wow-player: SHIP (round 2, for this lane's scope)

Round 1 was BLOCK. Each blocker was answered as out of scope, and round 2 accepted the answers.
The owner follow-ups below are logged, and none of them comes from this lane:

1. **Report-level "N fights" totals count trash.** These appear in the report subtitle, the
   /logs and /account rows, the sample row and the unfurls. The API's fight_count has no
   boss/trash split. A raider reads "79 fights" on a bossless report as boss attempts.
   Suggestion: split the count, or call them pulls.
2. **"N mobs died · your group killed M"** (logs-engine-counts wording). "1 mob died · your group
   killed 1" reads as a contradiction; a "kills" label would help.
3. **The night views' Active % has no stated denominator.**
4. **The reports list rows** have no chevron affordance, repeat "Not credited" on every row, and
   carry a "READY" pill that says nothing.

Taken here: "1 deaths" now reads "1 death" in the Summary roster meta.

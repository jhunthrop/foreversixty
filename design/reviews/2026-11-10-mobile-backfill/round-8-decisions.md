# Round 7 findings: what we change in round 8

Round 7: guild-officer WORTH PAYING FOR; brand CONSISTENT; ux-designer SHIP; info-density DENSE WITH
ADDITIONS; ai-slop FAIL (one must-fix). Codes: ID, GO, BR, UX, AS (round-7 sections).

## Both apps
| # | Change | From |
|---|---|---|
| 1 | The Gear sheet's source row has one anatomy on both apps: the source alone, in the shared source words (`bis.source.*`, the `gear.where.*` shapes: `Onyxia, Onyxia's Lair, 8%`, `Crafted: Blacksmithing`, `World drop`), with its figure when the data has one (drop chance; a profession's skill when the loot data carries it). No caption under it (`gear.sheet.where` stays only where it is a section heading, if anywhere), and no item icon beside it, so it cannot read as an item row. AS's tell 22 and UX's "label it" are both answered by making the row self-describing | AS must-fix 1, UX S1, ID S1, BR N1, N2 |
| 2 | The Mention failure title fits one iOS banner line: `Could not mention the healers` (`push.short.failed` is now `Could not mention the {who}`; the body keeps the raid and the retry) | UX S2, GO nit |
| 3 | A zero gain reads `No gain` (`gear.cell.noGain`), not `+0.0 DPS`, on Gear's cells | AS polish (twice), UX nit, BR N3 |
| 4 | Android's sheet: the runner-up row carries its gain when the data has one, as iOS does | ID S3 |

## Captures (round 8)
- The Gear sheet on both apps for a boss drop (with its chance), a crafted piece, and an upgrade piece (its gain line); Android ID S2.
- The Decline failure on both apps (GO S2).
- The iOS drawer Notifications frames from the top of the drawer (UX nit).
- The Mention failure banner on iOS (one line).

## KEEP
- A failed officer push carries no action buttons: the body names the retry and the tap opens the raid or the applications list, which UX clears (S3); re-posting the actions is a platform-specific path for a rare failure (GO S1).
- `Synced` / `Updated … from Battle.net` (round 7 KEEP), the duplicate day on the raid row (round 6 KEEP), Gear cells without a source sub-line (round 7 KEEP).

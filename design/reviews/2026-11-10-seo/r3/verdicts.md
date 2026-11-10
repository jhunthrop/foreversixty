# Round 3 verdicts (shots-r3): final

| Reviewer | Verdict |
|---|---|
| ux-designer | SHIP |
| wow-player | SHIP |
| ai-slop-reviewer | PASS |
| info-density-reviewer | DENSE |

All round 2 musts and shoulds are resolved:
- The weights lead no longer claims the player's own gear.
- The default note names the default.
- The bare preset reads "no buffs".
- The no-JavaScript line agrees with the table.
- The source fields, Load buttons, Run and the stat checkboxes are disabled in the server render, so a paste can't be dropped when the island mounts.

CLS is 0 on /sim/weights, /sim and home at every width, and at most 0.0044 on /planner.

## Applied after round 3

Three reviewers flagged the spec name appearing twice on the default card, so the note now reads "Level 60 until you load your own." The spec name stays in the card head. This is a copy-only change, so no reshoot.

## Open follow-ups (not blocking, raised by the reviewers)

**/sim/weights with JavaScript off:**
- The lone disabled Run and the "240 engine iterations" line.
- The inert source cards above the table.
- The "Change" toggle at 390.

**Site-wide, PageHeader:** put the lead before the meta line.

**Outside the lane (pre-existing):**
- **/sim/weights:**
  - "From the addon" appears twice;
  - the stat picker offers spell schools to a warrior;
  - the table's empty icon gutter;
  - the haste caption's "this level range".
- **/planner:**
  - the "Level 9" meta;
  - the empty counter block and race select with JavaScript off.
- **/rankings:**
  - the 390 CLS from the empty-state rail inserted above the board;
  - the board's skeleton never resolves with JavaScript off.
- **Needs data:** a rankings top-N snapshot in the nightly, so the board's rows can be server-rendered.

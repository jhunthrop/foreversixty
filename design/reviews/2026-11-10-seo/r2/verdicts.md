# Round 2 verdicts (shots-r2)

| Reviewer | Verdict |
|---|---|
| ux-designer | SHIP |
| wow-player | NO-SHIP (one must) |
| ai-slop-reviewer | PASS |
| info-density-reviewer | DENSE |

## ux-designer: SHIP

Every round 1 must and should is resolved:
- the rankings hole;
- the default table is named;
- the weights lead matches the table;
- the /sim, /rankings and /planner leads;
- the planner heads line up.

CLS is 0 on /sim, /sim/weights and home at every width, and at most 0.0044 on /planner.

Findings:
1. **should:** with JavaScript off, the weights shell's bar is a lone disabled Run. A follow-up is fine.
2. **should:** put the lead before the meta line site-wide, as a follow-up issue.
3. **nit:** the weights lead says "your gear" above a default table.
4. **nit:** the noscript line wraps at 390. Shorten it.
5. **nit:** the caution line sits above a default table. This copy is outside the lane.

**Outside the lane:**
- /rankings 390 CLS comes from the empty-state rail that is inserted above the board.
- The rankings skeleton without JavaScript.
- The planner's "Level 9", counter block and race select.
- "From the addon" appears twice.
- The icon gutter.
- The stat picker vocabulary.
- The caption.

## wow-player: NO-SHIP

1. **must:** the weights lead still reads as a claim about the player's own gear above a Fury Warrior table. Say the table shows a level 60 Fury Warrior until a character is loaded.
2. **should:** "Level 60 BiS, bare" does not match the toggle's "None". Use "no buffs". The "raid buffs" half is resolved.
3. **should:** the noscript line "Stat weights runs in your browser and needs JavaScript." sits above a working table. Use "Loading your own character needs JavaScript. The published figures below work without it."
4. **nit:** at 390 the weights lead wraps to four lines.
5. **nit:** the /sim lead is resolved and good.
6. **nit:** the /rankings lead is clear.
7. **nit:** the /planner lead is accurate and the best of the four. The tree heads read like the in-game frame.

**Outside the lane:**
- "Level 9" on a default Warrior.
- The planner's race select and counter with JavaScript off.
- "From the addon" appears twice.
- The stat picker offers spell schools to a warrior.
- The caption.

## ai-slop-reviewer: PASS

1. **should:** with JavaScript off, the weights shell's controls are inert and the cost line is sim talk. Make them hidden or inert.
2. **should:** the noscript line contradicts the table.
3. **should:** the weights lead says "your gear" above the default table.
4. **should (accepted as a decline):** lead before meta, as a site-wide follow-up.
5. **nit:** planner "Level 9" is outside the lane. The heads are right.
6. **nit:** rankings is fixed.
7. **nit:** the copy audit is clean, with no em dashes.

## info-density-reviewer: DENSE

1. **nit:** with JavaScript off, give the noscript line a gap. Optionally hide the inert source cards when JavaScript is off.
2. **nit:** the spec name appears twice, in the card head and in the note.
3. **nit:** the link wraps at 390. Acceptable.
4. **nit:** rankings is fixed.
5. **nit:** the planner heads are right.
6. **nit:** comment on the default faction order.

**Later:**
- Links to every spec's weights.
- Filter the stat picker by spec.
- The caption.
- A rankings top-N snapshot file in the nightly, so the board's rows can be server-rendered.

## What round 3 changed

- **Weights lead:** "What one more point of each stat adds to your DPS. The top stat in the table reads 1.00 and the rest scale from it. Load your character to weigh your own gear, talents and buffs." (wow 1, ai 3, ux 3)
- **Default note:** "Fury Warrior at level 60 until you load your own." The spec label comes from the spec list. (wow 1)
- **Link label:** "Level 60 BiS, no buffs" for the bare preset. (wow 2)
- **Weights noscript line:** with the default table it now reads "Loading your own character needs JavaScript. The published figures below work without it." It has bottom padding. (wow 3, ai 2, ux 4, info 1)
- **Default faction order:** now has a code comment. (info 6)

## Kept as declined

- **The shell's bar and cost line.** They match the island's first frame, so nothing jumps on mount.
- **Lead before meta.** This is a site-wide PageHeader follow-up.

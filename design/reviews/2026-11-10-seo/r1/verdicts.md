# Round 1 verdicts (shots-r1)

| Reviewer | Verdict |
|---|---|
| ux-designer | NO-SHIP |
| wow-player | NO-SHIP |
| ai-slop-reviewer | PASS |
| info-density-reviewer | DENSE |

## ux-designer: NO-SHIP

1. **must, /rankings:** the meta line is now a lone no-break space with no date. It leaves a blank band of about 40 px between "Rankings" and the lead. Without JavaScript the hole is permanent. Move the date's slot under the lead.
2. **should, /sim and /sim/weights:** the caveat meta line sits above the lead. Put the lead first.
3. **should, /sim/weights:** the lead's "best stat reads 1.00" reads against the table note's "Haste: 5.63 per 1%". Say "the top stat in the table".
4. **should, /planner:** with JavaScript off, the loading tree heads sit 7 to 11 px left of the editor's heads. Give them the tree's padding.
5. **nit, /planner at 390:** the lead's left edge does not line up with the skeleton's.

**Outside the lane:**
- The weights table has an empty icon gutter.
- "From the addon" appears twice.

## wow-player: NO-SHIP

1. **must, /sim/weights:** the Fury Warrior table with no character loaded does not say that it is a default. Name it and say "Load your character for your own."
2. **should:** the "Level 60 raid set" link label reads as raid gear before raids open. Use neutral BiS wording.
3. **should:** the table caption says "Hit = 1.00" against the row "Hit rating", and a haste figure sits in the caption. This copy is outside the lane.
4. **should, weights lead:** use "The top stat is 1.00, the rest scale from it."
5. **should, /sim lead:** "Fight a raid boss" overclaims before raids open. Say "a boss".
6. **nit, /sim lead:** use "a breakdown by ability" in place of "where every point of it comes from".
7. **should, /rankings lead:** "parses" is Warcraft Logs vocabulary. Use "your kills".
8. **nit:** the planner lead is accurate and the best of the four.
9. **should, /planner:** "Level 9" sits against the lead's "from level 10". Pre-existing meta.
10. **nit:** the planner's loading tree heads now read right.
11. **nit:** the rankings stamp line is fine.

**Outside the lane:**
- "From the addon" appears twice.
- The no-JavaScript line contradicts the table.
- The stat picker offers every stat for a warrior.

## ai-slop-reviewer: PASS

1. **should:** the weights lead describes "your gear" over a Fury Warrior table. Say whose figures they are.
2. **should, /planner:** "Level 9" sits against "level 10" in the lead.
3. **nit:** the /sim lead's clause is clumsy. Suggested wording: "a damage breakdown by ability and a cast-by-cast timeline".
4. **nit:** check that the companion uploads "as you play". `scoreCopy.explainerLogsBody` says it does.
5. **nit:** the rankings stamp leaves a blank line with JavaScript off.
6. **nit:** with JavaScript off, the weights shell's controls are inert. Move the noscript line under the lead.
7. **nit:** the planner's rail is still skeleton-only with JavaScript off. Acceptable for this lane.

**Outside the lane:**
- "From the addon" appears twice.
- The table caption is qualifier-heavy.
- The rankings encounter skeleton never resolves with JavaScript off.
- The planner's race select renders empty with JavaScript off.

## info-density-reviewer: DENSE

1. **should, /rankings:** blank gap under the h1 (the same issue as ux finding 1).
2. **should, /sim/weights with JavaScript off:** dead controls. Drop the bar and the cost line from the shell, or keep only the content.
3. **should:** the weights lead says "your gear" over the default table. Put the lead before the meta line.
4. **should, /planner with JavaScript off:**
   - an empty counter block;
   - "Level 9";
   - an empty race select.

   These are pre-existing.
5. **nit:** the planner's guide-build list could be server-rendered.
6. **nit:** the rankings rows have no build-time source, which is correct.
7. **nit:** comment that the default faction order makes no claim about which faction is first.

**Outside the lane / later:**
- Links to every spec's published weights.
- Filter the stat picker by spec.
- The haste caption.
- "From the addon" appears twice.

## What round 2 changed

- **Rankings:** the stamp moved from the meta line to the head's notice line, under the lead. That line holds the no-break space, so there is no hole between the title and the lead. (ux 1, info 1, ai 5)
- **Weights:**
  - The default table now says whose figures it shows: "Fury Warrior at level 60. Load your character for your own." The spec label comes from the spec list. (wow 1, ai 1, info 3)
  - The lead now reads "The top stat in the table reads 1.00 and the rest scale from it." (ux 3, wow 4)
  - The link label is now "Level 60 BiS, raid buffs" (or "bare"). (wow 2)
  - The noscript line moved to the top of the mount. (ai 6)
- **/sim lead:** now "Run your gear, talents, buffs and consumables against a boss. You get your DPS, a breakdown by ability and a cast-by-cast timeline of the fight." (wow 5 and 6, ai 3)
- **/rankings lead:** now says "so your kills join the board". (wow 7)
- **/planner:**
  - The lead now says "up to level 60", which no longer reads against "Level 9". (wow 9, ai 2)
  - The loading heads take the editor tree's 12 px padding with no grid gap. (ux 4)

## Declined, with reasons

- **ux 2 and info 3, lead before meta.** PageHeader orders the meta line under the title on every page; reordering it is a cross-site change outside this lane.
- **info 2, drop the bar and cost line from the shell.** The shell matches the island's first frame, so the table does not jump on mount. The noscript line now leads the mount.
- **Pre-existing items** stay out of scope and are listed for the owner:
  - planner "Level 9", the counter block and the race select;
  - "From the addon" appearing twice;
  - the stat picker vocabulary;
  - the table caption;
  - the icon gutter.

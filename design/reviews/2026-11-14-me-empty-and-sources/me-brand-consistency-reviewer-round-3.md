# brand-consistency-reviewer, round 3 (Me with no site activity)

Verdict: PASS

The Round 3 layout and copy changes hold. I found no must-fix items. I read the re-rendered `me-empty-l25.png` and `me-empty-l60.png` and spec section 13. I did not open the `-l5` or `-point` captures this round, or the code.

**Fixed since round 2**
1. **Enchant rows:** `Chest: Minor Stats` and `Wrist: Minor Agility` now show the slot and enchant without truncation. Warning colour stays on the enchant rows only, and the Gear tab keeps its long wording.
2. **Gear in the shorter column:** at level 60 Gear moves under Where you stand. The left column is no longer sparse and the right is no longer overfull. Heading, note and row anatomy are identical in either column.
3. **Upgrade sources:** each source drops the slot because the item name already says it. `The Deadmines · Captain Greenskin · 30% dr...` still clips the drop chance on the one longest row at level 25.
4. **Unlock row:** `New gear list` / `Level 30` is a label and value pair like the other rows, with no colon or sentence.
5. **Site line:** it shows only without the companion, so it is not repeated beside "Companion synced".
6. **Unchanged and correct:**
   - The brand tokens are intact: quality colours, success-green gains, a single gold link, muted labels.
   - Copy is sentence case with no em dashes.
   - "best in slot" and "on plan" match the web.

**Shoulds**
1. **Clipped drop chance:** the Level 25 Impaling Harpoon row cuts at `30% dr...`. The drop chance is the number a player acts on. Cut the boss name first, or drop the middle dot segment, so `30%` always survives. Fix it in `GearTab.sourceWhere` if the Gear tab has the same clip.
2. **Warning red on the full enchant text:** the whole enchant row is red. This is the shared Gear tab style, so it is consistent. If the web's missing-enchant row shows only the label in warning colour, change the shared row, not Me.
3. **Dropped tooltips:** `Next point` still has no spell tooltip. Spec 13 notes the Build hover and the talent tooltip are not done. A player who hovers the talent will notice.
4. **Level 5 `Pick your spec`:** it stays in gold in Right now with the level-5 build row reading `No build followed`. That is by decision (section 13), and it is the same gold act-now style as the web's primary prompts. Fine.

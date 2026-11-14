# brand-consistency-reviewer, final (round 4): Me with no site activity

Captures: me-empty-l25.png, me-empty-l25-point.png, me-empty-l5.png, me-empty-l60.png (design/mocks/addon_hub_me_empty.py). Earlier rounds: me-brand-consistency-reviewer-round-1.md to -round-3.md.

Verdict: PASS

The round 4 change fixes the clipped drop chance, and I found no must-fix items. I read the re-rendered `me-empty-l25.png` and `me-empty-l60.png`. I did not open the `-l5` or `-point` captures, the spec section 14 text or the code.

**Fixed since round 3**
1. **Source lines:** the Impaling Harpoon row now reads `The Deadmines · Captain Greenskin` with no ellipsis. Dropping the drop chance first is the right order. The boss name is the actionable part, and the chance is still available in the item tooltip and on the Gear tab.
2. **Gear placement:** at level 25 Gear is back on the right, which matches the hub's usual column order (act-now content and Gear on the right). At level 60 the left column is clearly shorter (two rows), so Gear sits under Where you stand. The columns stay balanced and the section anatomy is identical in both places.
3. **Level 25 left column:** the site line sits comfortably below the five-row Where you stand block, and the gold link is the only gold on that side.
4. **Unchanged and correct:**
   - **Colour:** quality-coloured item names, green gains, warning-red enchant rows, muted labels.
   - **Wording and enchant rows:** `best in slot` and `on plan` match the web, and the short enchant form fits.
   - **Sim:** "Not simmed yet" with the "Run a sim" button.

**Shoulds**
1. **Placement rule:** the rule "right unless the left is clearly shorter" is invisible to players, so Gear can jump columns across levels. That is acceptable. Make sure the threshold lives in one named constant in `MeFrame.draw`, not a magic number, and has a test.
2. **Whole-row warning red:** the enchant rows are red in full. This is shared with the Gear tab, so it is consistent. If the web shows only the label in warning colour, change the shared row.
3. **Dropped tooltips:** `Next point` still has no spell tooltip and the Build name has no hover. Both are spec decisions, but a player who hovers the talent will notice.
4. **Right column at level 60:** it is mostly empty below the sim row. This is honest, and it matches the "no padding" rule.

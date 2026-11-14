# ai-slop-reviewer, final (round 4): Me with no site activity

Captures: me-empty-l25.png, me-empty-l25-point.png, me-empty-l5.png, me-empty-l60.png (design/mocks/addon_hub_me_empty.py). Earlier rounds: me-ai-slop-reviewer-round-1.md to -round-3.md.

Verdict: PASS

Round 4 changes nothing that was a must-fix. The l25 and l26 layouts are balanced again. Gear sits on the right above the sim, and there is a sensible reason for the l60 exception. The source line now fits without cutting words ("The Deadmines · Captain Greenskin"). There are no visual tells and no truncation.

**Must-fix**
None.

**Shoulds**
1. **l60 is asymmetric.**
   - The left panel holds Gear, and the right panel is about 350 px of blank under "Your sim".
   - The "left is clearly shorter" rule only applies at 60. Even so, putting the sim under Gear on the left, or adding a second Right now row (next raid reset, attunement state), would balance it.
2. **"Your sim" is still three statements: the "Your sim" heading, "Not simmed yet", and "Run a sim".** Make it one row, "Sim: not run" with the button.
3. **l5: the "Builds" link on the right of "Pick your spec" repeats the tab dot.** Drop the link text.
4. **l26: "Next point ... Now" and "Spend 1 point" repeat each other.** Keep the call to action and drop "Now".
5. **l25 and l5: the "No companion" footer and the site sentence ("Your score, raid logs and sims come in through the companion app.") say the same thing.** The footer alone is enough, so keep the URL and cut the sentence.
6. **The drop chance is gone from Impaling Harpoon.** This is correct for width. Check it still appears in a tooltip, because "30% drop" is exactly the number a player wants.

**Authentic, keep**
- Source lines ("Crafted · Leatherworking", "Scholomance · Kirtonos the Herald").
- Enchant rows in the red "Chest: Minor Stats" form.
- Professions with skill numbers.
- "New gear list / Level 30" as label and value.
- The Onyxia's Lair raid line.
- Quality-coloured item names with DPS deltas.

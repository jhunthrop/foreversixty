# ux-designer, final (round 4): Me with no site activity

Captures: me-empty-l25.png, me-empty-l25-point.png, me-empty-l5.png, me-empty-l60.png (design/mocks/addon_hub_me_empty.py). Earlier rounds: me-ux-designer-round-1.md to -round-3.md.

Verdict: SHIP

Round 4 resolves my round 3 shoulds 1 to 3. Nothing blocks.

**Verified in the round 4 captures**

- **l25-point.**
  - The window is about 928 px tall, down from 1072.
  - Gear sits in the right column between Right now and Your sim.
  - The Right now row is the one line `Spend 1 point` with the `Builds` link, so Right now is not duplicated on the left.
  - The columns end within about 50 px of each other, and the right column is the taller of the two.
  - The `Now` note in Where you stand is body colour, not gold.
- **l25.**
  - Gear stays on the right, as it does in l25-point, so Gear no longer jumps columns between 25 and 26.
  - The Harpoon source reads `The Deadmines · Captain Greenskin`, with the chance dropped and the boss whole, as spec section 14 says.
  - The left column has all five rows, the single site line and the gold link.
  - The window is 788 px, which is shorter than the normal Me layouts.
- **l60.**
  - Gear is on the left under Where you stand, which has two rows, and Right now keeps the Guild raid row and Your sim.
  - The columns end within about 60 px of each other, so this is a balanced layout.
  - There is no Experience row and no site line, matching section 12.

**Shoulds**

1. **Dropped drop chance.** The Harpoon's `30% drop` is now gone on Me, but the same item shows it on the Gear tab. This is acceptable. The player can still see the chance when they open Gear, and the boss name is the part that tells them where to go. Mention it in the lane notes.
2. **Empty space on l60.** The lower right below Your sim is empty. It is a normal consequence of this state and needs no change.
3. **Still to log, outside this lane.**
   - The live `PLAYER_XP_UPDATE` redraw is not built.
   - The Set up step 2 question-mark icon is a missing-icon flash.
   - The tooltip hover captures were not producible from the renders. The redraw rule rests on the named tests, which I have not run.

# ux-designer, round 3 (Me with no site activity)

Verdict: SHIP

Round 3 resolves my round 2 shoulds 1 to 3. Nothing blocks. One layout choice makes the window taller than it needs to be.

**Verified in the round 3 captures**

- **Enchant rows.** They now read `Chest: Minor Stats`, `Wrist: Minor Agility` and `Back: Lesser Agility`, with the planned enchant's icon, warning colour, and nothing cut off. This fixes the lost slot name.
- **Point-waiting row.** `Spend 1 point` in client gold with the `Builds` link, on one line, with no truncation. The talent name stays in the left `Next point` row, whose `Now` note is body colour. Right now does not repeat the left column.
- **Sources.** The gear sources no longer repeat the slot. The l25-point capture shows `The Deadmines · Captain Greenskin · 30% drop` in full.
- **Site line.** It shows once, with the gold link, only without the companion: l25 and l5 have it, while l25-point and l60 have none. This matches section 13 as you described it.
- **Labels.** `First talent` / `Level 10` and `New gear list` / `Level 30` replace the colon sentences. They are shorter and read as a table.
- **Your sim.** `Not simmed yet` with `Run a sim` is consistent across captures. The new test `keeps Sim again shown across a redraw` covers the button no longer hiding on redraw, per your test run.
- **Level 60.** There is no Experience row. Professions and Gear at band 60 show. The Guild raid row is not duplicated.

**Shoulds**

1. **Window height in l25-point is 1072 px, against 788 in l25 and 928 for the same state last round.**
   - Gear now sits in the left column under five rows. The right column holds one row and Your sim, then about 480 px of nothing.
   - The "shorter column" rule counted the right column as taller. In the render, the left column is shorter by about 50 px before Gear is added. Check the height arithmetic and which column `buildGear` picks. For a one-row Right now, Gear on the right gives roughly 930 px and balances the columns.
   - Section 8 says the tab must not make the window taller than the tallest of today's Me layouts. Confirm 1072 px is within that, since `Hub.fitHeight` sets the in-game height. If it is not, this is a must-fix.
2. **Gear moves between columns between states.** It is on the right in l25 and l60, and on the left in l25-point. A player who opens the tab, spends the point, and reopens sees it jump. This is a smaller issue than should 1. It could go away if rule 1 puts Gear on the right consistently.
3. **l25 right column.** The Harpoon source line is still cut at `30% dr...` at 576 px. Drop `· 30% drop`, or shrink the item name. Once Gear moves left or the column is wider, it fits.
4. **The XP bar still updates only on open and hub refresh.** This is accepted as a follow-up. Record it in the lane notes.
5. **Not in this lane, still to log.** The question-mark icon on Set up step 2, and `Run a sim` alone on l60 (visible here as the empty right half).

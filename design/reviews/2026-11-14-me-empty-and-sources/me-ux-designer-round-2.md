# ux-designer, round 2 (Me with no site activity)

Verdict: SHIP

Round 1's must-fix 1 is resolved. Must-fix 2 is accepted on the terms below. Nothing blocks, but the truncated rows should be fixed soon.

**Round 1 must-fixes**

1. Resolved. In `me-empty-l25-point.png` the Right now row is now the gold `Spend 1 point: Improved Aspect ...` with a `Builds` link. The build name and points on plan no longer repeat in the right column. The left `Now` note is body colour and matches that row. The "points on plan" contradiction is gone.
2. Accepted.
   - The tooltip hover captures cannot come from the HTML renders. The talent tooltip was cut by amendment.
   - I accept the named test `redraws the quiet state without hiding a control that stays` as the evidence for the redraw rule.
   - I did not run that test or see the code. If the lane owner wants this on record as a pass, they should run `addon/tests/me_tab_spec.lua` and paste the result.
   - The l5-setup board is dropped, and its change (no site line) is tested.

**Verified in the new captures (l25, l25-point, l60)**

- **Layout.** Where you stand is on the left, Gear is on the right above Your sim, and the site line appears once. There are no boxes.
- **Gear rows.** Each row has a 20 px icon, the name in its quality colour, the green gain, and a muted where-line. They follow the Right now row anatomy.
- **Experience and Professions.** The XP track and `Rested 40%` show. Professions shows real skill lines. The Experience row is left out at 60, replaced by the header, which is a reasonable change.
- **Right now.** At l60, `Onyxia's Lair ... Guild` is not duplicated on the left.
- **Your sim.** `Not simmed yet` with `Run a sim` reads correctly. This is a global change and is better than the old "Sim again" with no sim.
- **Site line.** The l25 site copy is the section 12 version. At l25 with the companion running, the line is omitted per section 12.

**Shoulds**

1. **Enchant rows cut the useful part.** `No enchant: Enchant Chest - Minor Stats · Che...` and `Enchant Bracer - Minor Agility · ...` lose the slot, which is what the player needs. The 60 row loses the rest of the enchant name.
   - Drop the `No enchant:` prefix, or put the slot first and cut the enchant name.
   - Check what the Gear tab itself shows at 576 px; fix it there if the shared component is the cause.
2. **The point-waiting row is cut off.** `Spend 1 point: Improved Aspect ...` ends well before the `Builds` link. The talent name is already in the left `Next point` row, so shorten the copy to `Spend 1 talent point` with the name kept in the left row.
3. **Spacing in l25-point.** Right now has a roughly 40 px empty band under its single row before Gear. Confirm this is `sectionGap` plus the row reserve, not leftover space.
4. **The tab no longer redraws on `PLAYER_XP_UPDATE`.** Section 9 required an in-place update at most once a second. Section 12 now says "not done, by decision". Fine for this pass, but it needs a follow-up item, since the XP bar will be stale while the hub stays open.
5. **Section 8 height limit.** l25-point is 928 px tall in the render. Confirm that `Hub.fitHeight` gives the same height in game and that it is no taller than the tallest normal Me layout.
6. **Pre-existing items stay out of scope.** The question-mark Set up icon and the lone `Run a sim` on l60 are not part of this lane. Log both as follow-ups.

# ux-designer, round 1 (Me with no site activity)

Verdict: FIX

The state renders the right facts in the right places. One captured state breaks the spec's repeat rule, and acceptance captures are missing.

**Must fix**

1. **me-empty-l25-point.png: Right now repeats the left column.**
   - Right now shows `Beast Mastery leveling` / `15 of 15 points on plan` with a gold `Builds` link.
   - Where you stand row 2 shows the same build and `15 of 15 on plan`. This breaks the verdict gate "no Right now row repeated in the left column".
   - The row says nothing about the point that is waiting. Spec section 11 says a Builds item that waits still shows there in client gold. Spec capture 4 expects "Right now carries the gold line".
   - `MeTab.rightNow` appears to fall back to the calm lead when the tab is in `QUIET_SAID` and the item is not urgent. Here the waiting point is not reaching Right now as its own line.
   - Ship: the Right now row for a waiting point reads the point itself, for example `Spend your talent point` with the talent name, in client gold. The calm lead stays out.
   - Check that the row's text matches the left `Now` note, and that `15 of 15 points on plan` cannot show alongside a waiting point.
2. **Acceptance captures are incomplete (spec section 10).**
   - Missing:
     - Capture 7: item cell and talent icon hover.
     - Capture 8: a `PLAYER_XP_UPDATE` redraw with the cursor on an item cell, no tooltip flicker.
     - The `Nothing to chase at 60.` alternate for capture 6.
     - The left-column crops.
   - Capture 7 can drop the talent half under the section 11 amendment.
   - I cannot verify the "no Hide or Show under the cursor" rule from stills. Attach a note or a two-frame capture, plus the test that covers it in `addon/tests/me_tab_spec.lua`.

**Verified OK**

- **Section 11 amendments.** Gear sits in the right column. The calm Builds lead is gone from Right now in l25 and l60. The Where you stand heading has no note. The build and talent rows have no tooltip.
- **l25.**
  - All four rows show.
  - The XP track fills at 62% with a muted percent.
  - `Rested 40%` is muted.
  - Three upgrades have quality-coloured names and green `+x DPS`.
  - `2 enchants missing` is in warning colour with the slot names muted.
  - The site line has exactly one gold link.
  - `Next` reads `Level 30, a new gear list`.
- **l25-point.** The `Now` note is body colour, not gold. The companion-running variant of the site line has no link.
- **l5.** There are three rows, no Next point row, and no Gear section. The site line and link show, and the link is gone in the Set up capture, where step 3 carries the address.
- **l60.** `Level` / `60` / `Max level` shows with no track. Build reads `51 of 51 on plan`. Rows 3 and 4 are omitted. Gear shows band 60. The no-companion line is the shorter variant ("after your first logged raid") because the companion is synced in the capture, which matches section 3.3.
- **Layout.** There are no boxes or rules. Window height follows the content, and l25 is shorter than the normal Me layout.

**Shoulds**

3. **l25 and l5: the gap before the site line.** In l25 it is about 40 px from the last row to the site text, and in l5 about 40 px as well. The spec says `sectionGap` 16 plus `y + 2`. Confirm this is the 18 px row pitch plus the gap, not a stray margin.
4. **l25 and l60: the right column starts with Gear at the top, so Gear is the first thing there.** The spec's hierarchy makes Gear gains primary, so this is acceptable. Add a line to the section 11 amendment saying Gear moves up when Right now is empty.
5. **Pre-existing items you flagged.** The `Your sim` heading with only Sim again on l60, and the question-mark icon on Set up step 2, do not block this lane. Log both as follow-ups. The icon is a missing-icon flash and breaks the spec's no-missing-icon rule on that surface.
6. **l5 Right now.** `Pick your spec` with `Builds` is a single row on an otherwise empty column. That is right for the state and needs no change.

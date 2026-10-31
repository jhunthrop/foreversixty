# F6 Talents mock: wow-player verdict (2026-10-31)

Surface: boards f6-talents, f6-rank, f6-states, f6-xl (iOS and Android), against the spec. Question: "Is my build the guide build?" The status line answers it first on screen, in plain words, before any tree. The count sentence is true for the vector (Reverberation +2 over the guide gives 2; Call of Flame and Elemental Warding are 4 short; 28+3+18 = 49; 2 unspent). The rank sheet follows the game tooltip order (Rank n/max, current text, Next rank). Band rule, under-20 state and the empty states read right. No gear, DPS or sim figure anywhere; the footer link goes to /planner.

## Blockers

1. **Differences in a closed tree are invisible.** The tree that opens is the one with the most points, so the player sees dashes only there. If the 2 differing points are in Restoration, the status says "Differs" while the open tree shows no dashed cell, and the tabs ("Elemental 28  Enhancement 3  Restoration 18") give no hint where to look. "The tree is the proof" fails. Put a gold dot or "n differ" on each tab whose tree has a dashed cell (on Your build), with its own spoken value.
2. **Gold sends the player to the wrong talents.** In F6.2 the two unspent points light up Trial by Fire, Eye of the Storm, Earthbound, Elemental Reach and Elemental Focus in gold, while the guide wants those points in Call of Flame and Elemental Warding, which are shown green and dashed. The count sentence says "differ" and the unspent line says "2 points unspent", so a player will read gold as "spend here". Either do not draw gold on the compare screen, or make the cells the guide is short on (your rank below the guide, prerequisites met) the only gold ones. Also, a partly filled talent that can still take a point (1/3) is not "available" in the spec, which is not how the game reads it.

## Shoulds

1. **Dash has no direction.** Reverberation 2/5 (you are over the guide, respec it) and Call of Flame 1/3 (you are under, take points) wear the same dashed ring. Use a different mark: a small minus/plus, or two ring colours. The sentence only counts "over" points and says nothing of the 4 missing, so add "and the guide takes 4 more" or similar when missingCount > 0.
2. **Colours are inverted against the game frame.** In the Classic talent frame a maxed talent's rank is gold and "can still take a point" is green. Here maxed is green and available is gold (the spec says it borrows the game's colours; it does not). It matches the site planner, so fix both or add a colour key to the legend. Players will read 5/5 green as "still open".
3. **Locked cells give no reason.** F6.4 Elemental Mastery shows only "Rank 0/1" and "Next rank". The game tooltip's red lines ("Requires 30 points in Elemental Talents", "Requires 1 point in Lava Burst") are missing, and a locked-looking 0/x cell with no unspent points (Matches board) cannot be told apart from a truly locked tier. Add the requirement lines to the sheet from tier x 5 and prerequisite data.
4. **Prerequisite lines have no arrowheads.** The game draws an arrow into the dependent. Without it Elemental Alacrity to Lightning Overload to Elemental Fury reads as one long bar.
5. **"Set a spec in the planner" (talents.pickSpec) is wrong for Classic players.** A spec comes from spending points, not from a setting. Reword to what the data actually lacks.
6. **iOS large title bleeds.** F6.5 shows "TS" of the title behind the Home pill, F6.6 shows ghost status text above the segments, and XL.2 has huge "toration 18" behind the collapsed title. Give the bar a solid or blurred ground.
7. **Locked/empty rows hide under the tab bar** on the first screen of F6.1-F6.3 (bottom tier half covered); spec says tiers 1 to 6 are enough, but the capstone Elemental Mastery is the talent players look for, so scroll hint or shrink the pitch.

## Nits

- Rank sheet and tab names are all caps condensed ("CALL OF FLAME"); the game writes talent names in mixed case.
- "the 60 build" reads oddly; "the 60 guide build" matches the Matches sentence.
- "Not available yet: no talent export..." on F6.8 wraps to three lines beside a tiny tag; fine but the tag text is hard to read.
- Matches board hides the unspent line correctly; state it in the spec for the case Matches with unspent.

## Data I do not believe (for the accuracy reviewer)

- The guide build (`5510331300001350-...`) spends exactly 30 in Elemental and skips Elemental Mastery (0/1), the capstone that needs 30 points first. A 51-point Elemental build with no Mastery will look wrong to players; confirm it is the guide's intent.
- Band guide strings: the compare takes the band's whole string with no level cut. A level 29 character against the "20 to 29" band, or 59 against "50 to 59", will show "Differs" for the extra points unless the band string is sized to the top of the band. Confirm what level each band string was built for, or the sentence will say "Differs" for most players below 60.
- Elemental Mastery text ("reducing the cast time of Lightning Bolt and Lava Burst by 0.5 sec ... Lasts 15 sec") and the Call of Flame text are Forever client text; confirm verbatim against ranks[].description.

VERDICT: HOLD

## Re-check (after the ux-designer's revision)

Viewed f6-talents-ios, f6-talents-and, f6-rank-ios and the revised spec and ux verdict.

- **Blocker 1 (closed trees): cleared.** F6.2 shows a red dot on Elemental and Restoration, and the vector now has a Restoration difference (Tidal Focus), so the dot is demonstrated. The tab dot has a spoken value.
- **Blocker 2 (gold misdirects): cleared.** Gold "available" is gone on Your build. The cells the guide is short on wear a gold "-2" chip and the over-guide cell wears a red "+2". The line "The guide build takes 5 points you have not." says what is missing. Elemental Mastery now shows the red "Requires 30 points in Elemental Talents" line, and arrowheads are drawn.
- **Rejection: maxed green vs the game's gold. Acceptable.** Parity with the planner and a printed rank on every cell outweigh it, because the cell no longer carries a gold "available" meaning that could clash. Please keep the legend from calling green "maxed". It should say "points".
- **Rejection: band-string length (Differs forever below 60).** Acceptable for the mock, with the `talents.beyond` explanation line. It is a real wart: a level 58 will always read "Differs". The owner should rule on it, and no board shows the line yet.
- Remaining nit: on Differs, the capstone row sits half under the tab bar (accepted as the first scroll).

VERDICT: SHIP

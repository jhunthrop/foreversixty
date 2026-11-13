# info-density-reviewer: addon Builds pane and panel Talents row (2026-11-13)

Verdict: FAIL (3 must-fixes; the structure is right, the one line per row that carries the most space carries the least).

Job: which build do I follow, what do I take next. Reached and unused: `addon-builds.json` / `Data.lua defaultBuilds[spec][kind].order` (51-point sequences, 28 specs), `.why`, `specIcons`, `Progress.pointsAllowed`, off-plan count (panel already shows "5 points off plan").

Must-fixes
1. The why line is cut mid-sentence on 9 of 10 rows (hunter-42, druid-24: "all five points of Efficie...", "Marksmanship take t..."). It is a full line per row that reads as nothing; the full text lives only in the hover tooltip, which the game has no touch for. Replace it with what differs between that spec's two builds, computed from `order` and the talent names in Data.lua: raid row `+Improved Stings, Mortal Shots / -Efficiency, Improved Tracking`, level 60 row `Takes Efficiency 5/5`. Fits 388 px as one muted line, answers "what changes if I follow this one", and drops the ellipsis. Keep the why only in the tooltip. One-build specs show their top-3 first talents or no line (no sentence fragments like "It is built for a tank-healing fight...").
2. Panel Talents row line 2 says `Beast Mastery build`, but the pane now has two kinds per spec. Name the kind (`Beast Mastery level 60 build` / `raid build`, `Follow.buildName` already holds the words) or the player cannot tell which of the two rows he is on. Cut with an ellipsis only on site names.
3. Redundant text: site row sub-line `For Leafcall · from the site, 12 Nov` sits under the `From the site` heading (druid board), drop `from the site`; the hunter board spends two lines on "Builds you pick on foreversixty.gg arrive here through the companion app" under 5 own rows (explain copy), keep it only when no own rows and no site rows exist, otherwise drop. `Tab to switch` footer hint is chrome on every board, drop it.

Should (not blocking)
- Heading `Level 24 · 15 of 15 points` duplicates the panel; when off plan, add the figure the panel already computes (`5 points off plan`, warning colour), since the pane is where the fix is chosen.
- Followed row: `Next talent at 25: Feral Instinct` gets the talent icon from Data (16 px) so it matches the panel's row.
- Phones/app: not applicable (420 frame); the 482-690 px heights hold, the druid's 4 specs x rows stay unscrolled only because of fix 1 saving nothing in height, so if fix 1 makes the line optional it also trims 15 px per row.

Needs data: none.

## Round 2

Verdict: PASS.
1. Whys: every row on hunter-42 and druid-24 now ends whole on one line, and the lines say what the build changes ("Trades Efficiency for Improved Stings and Mortal Shots", "Trades Brutal Impact and Feral Charge for full Rend and Tear"). Fixed.
2. Panel row: the kind shows only when it differs from the level's own (middle "Beast Mastery build", right "Beast Mastery raid build"). Accepted as the deliberate 340 px call: the pane and chat name the kind and the ellipsis is gone. Fixed.
3. Redundant text: "from the site" is gone from the site sub-line, the hunter explainer is gone, and the Tab hint is gone. Fixed. The heading now carries "5 points off plan" in the warning colour, and the followed row carries the talent icon. Both of my Should items are done.
Optional: the "and 3 more" row should open the full site list if it does not already; the "Enter to follow" footer hint is the last chrome line, droppable.
Needs data: none.

# AI-slop review: addon Builds pane and panel Talents row (2026-11-13)

First impression: reads as a WoW addon made by a player. Native frame, real spell icons, real talent names, no glow, gradient, glass or pill. Nothing from the visual list (1, 2, 5, 8, 12, 13, 15) is present.

Verdict: PASS WITH FIXES (none are visual hard tells).

Must-fixes
1. (22) "From the site" heading, then a row captioned "For Leafcall · from the site, 12 Nov" (druid-24-site). The heading already says it. Caption becomes "For Leafcall · 12 Nov".
2. (21/19) "Builds you pick on foreversixty.gg arrive here through the companion app. Run it on your PC." explains the pipeline to a veteran; the pitch line under the paste box. Cut it, or reduce to "Companion delivers site builds."
3. (7) A grey blurb under every one of 9 build rows, all cut mid-word with "…" ("Efficie…", "Marksmansh…", "Veng…"). Reads as a template slot, and the cut text is the point of the sentence. Show a one-line summary that fits (the talent that differs), full text stays in the tooltip.
4. (22) Retribution's spec tagline truncates ("Seal of Command and Veng…") while the others fit; trim the copy so no tagline ends in an ellipsis.
5. Layout: "and 3 more" sits flush against the paste box with no gap (druid-24-site); the "Intimidation · Spend 1 now" row and its Change/Talents buttons overlap tightly in panel-talents-42 (right). Give each a line of air.
6. Copy consistency: the follow line says "Spend 1 now: Intimidation" in one place and "Next talent at 25: Feral Instinct" in another, and "Intimidation · Spend 1 now" in the panel. Pick one shape.

Authentic, keep
- Real spec taglines with named spells, real split strings "Marksmanship 31 / Beast Mastery 15 / Survival 5".
- "Are you Holy or Retribution?" heading at level 13 with the split; it is the one thing that matters then.
- Uneven hierarchy: Following in gold, the Spend line in gold, the failure in red only when it happens.
- Row states board (hover, cursor, held, did not load) is real developer thoroughness; follow buttons look clickable, rows do not pretend.
- Level 60 / Raid / "Level 60 and raid build" labels carry real meaning.
- No em dashes, emoji, eyebrows or numbering in the UI.

## Round 2

Verdict: PASS. Round 1 fixes landed: site caption is "For Leafcall · 12 Nov", the companion pitch line is gone, no blurb is cut mid-word, no tagline ends in "...", "and 3 more" has air, the Spend line has Change/Talents beneath it, and the follow line shape is one pattern. The left panel capture is the replayed bug, so its blank button and "buil..." are expected.
Whys (48 lines): no marketing words, no em dashes, every line names real talents. Fix when convenient, not blocking:
1. "Trades X for Y" opens 16 lines, almost every raid line that differs from its leveling build. It is a true diff, but reads as a formula in the list. Vary a few: "Raid swaps Efficiency for Improved Stings", or just name the new talents.
2. Role formula ("Healer on...", "Shield healer on...", "Block tank on...", "Shield tank on...", "Bear tank on...", "Melee hunter:") opens 8 lines; fine, the role is the point.
3. Small padding words: "full" (full Barrage, full Rend and Tear), "for trash and the way up", "to outlast the fight". Cut "full" where the point count is not the news.

# Brand consistency review: addon hub, Me tab (L2)

## Round 1

Verdict: PASS (no must-fix items; three nice-to-haves).

Checked: me-raid-day.png and me-no-pack.png against the web score card (CharacterScore.svelte, score-parts.ts, score/copy.ts, score-labels.json) and the shipped Builds tab (l1/window-hunter-42-hub.png).

What matches the web:
- Score line: number in its tier colour (Rare blue) with the tier word beside it, and "316 to Epic (1500)" is the web's toNextTier shape with the fixture's floors.
- Part labels and order from the fixture: Gear, Experience, Parses, Attendance, Preparation (the web's label, not "Discipline"). "410 of 500" is the web's ofMax string. "as of 9 Dec" matches asOf.
- Part bars: track length proportional to each part's max, fill coloured on the item quality ladder. Checked the fills against QUALITY_FLOORS: Gear 82% rare blue, Experience 54% uncommon green, Attendance 60% uncommon green, Parses 38% and Preparation 46% common white. All correct. Tier colours are read through Score.tierHex and Theme.QUALITY_HEX, with no raw colours in MeTab.lua or MeFrame.lua.
- Section headings use the Builds tab's title face through BuildsWindow.buildSection, so Me and Builds read as one window. Tabs, header crest, footer hint and the "Sim again" primary button are the same chrome as Builds.
- Sentence case throughout, no em dashes, game words ("No flask", "All up", "Downtime 18 s", "Hunter's Mark late 6 s"). Parse numbers use the quality ladder colours. Gear and item icons keep their quality edges, with the client's tooltip on hover.
- The no-pack state keeps the same anatomy and says what to do in one line.

Nice-to-have (not gating):
1. The owner-approved board hub-me.png shows "Raid-ready list 984.0" next to the sim figure, "51 of 51 points on plan" and a purple (epic) item name on the Gear line. The built render drops the raid-ready figure, says "51 of 51 on plan" (which does match the Builds header wording) and shows the gear name in white. Confirm with the owner that those three board deltas are intended; if the raid-ready figure is still wanted, add it.
2. Preparation's reason ("2 flasks missed") sits in the bar column where the other parts leave the space empty. Fine, but keep the text on one line at 740 px with long consumable names.
3. The sim's "Biggest gains" icons render with a small corner artefact on the icon edge (top-left tick) in the built render; the board does not show it. Check the edge texture crop.

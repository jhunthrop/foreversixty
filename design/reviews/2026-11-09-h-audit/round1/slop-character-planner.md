# AI-slop review: character and planner pages (first captures, 1440 and 390)

Reviewer: ai-slop. Pages: character-visitor-60, character-owner-60-premium, planner, planner-you, planner-guide.

## Verdicts
- character-visitor-60: PASS WITH FIXES
- character-owner-60-premium: PASS WITH FIXES
- planner: PASS WITH FIXES
- planner-you: PASS WITH FIXES
- planner-guide: PASS WITH FIXES

No page has visual tells 1, 2, 5, 12, 13 (one borderline, noted) or 15. Reads as made by a player: paperdoll, real item names, upgrade list with HPS gain, real talent trees, point order as icon strip.

## Findings
Tell 8 (borderline): every block on all five pages is the same dark rounded bordered panel with a serif heading (Score, Builds seen, Best per encounter, Activity, History, Coach history; planner: Level 60 build, In the game, Guide builds, Share, Import, Point order). Planner right rail is six identical stacked cards. Source: components/planner/Planner.svelte, components/character/CharacterOwnerSections.svelte. Fix: merge Share + Import + In the game into one plate with hairline dividers; drop card chrome on Builds seen and Activity.
Tell 22: "Paste in game with /fs." under Copy for the game, all planner pages. lib/planner/copy.ts:88. Fix: delete.
Tell 22: "Include a simmed DPS on the card" toggle. lib/sim/planner-copy.ts:18. Fix: "Show DPS".
Tell 22/18: "Talents loaded from a character." (planner-you) and "Loaded in the Survival guide's levelling order." (planner-guide), grey line above the tree. Planner.svelte ~213-220. Fix: drop; the character chip and Loaded pill say it.
Tell 22: "Covers 21 of your 25 points." lib/planner/copy.ts:261. Fix: "21 of 25".
Tell 22: "Client build 1.60.1.70790" in the talent plate footer, all planner pages. Fix: move to page footer or tooltip.
Tell 21: yellow dash marker before "Felsnap is level 34. Then the Beast Mastery guide's order." in Point order, planner-you 1440/390. lib/planner/copy.ts:95. Fix: plain hairline divider, "Felsnap is 34. Beast Mastery guide order from here."
Tell 19/13: History "Score over time, then parse per boss" and Coach history "Every review the coach wrote for this character", owner 1440/390. CharacterOwnerSections.svelte, components/coach/CoachHistoryList.svelte. Fix: collapsed rows name real content (latest score, review count) or are removed; they read as placeholder empty states.
Tell 22: character band meta "Supporter" and "Addon · 1 day ago", owner 1440/390, plain text with no visible state. Fix: "Synced 1 d ago", Supporter as a crest or dropped.
Tell 22: Builds seen rows "Holy · first seen 3 Nov 2026", character pages. lib/character/page-view.ts. Fix: "Holy, 3 Nov".
Tell 16 (mild): Forever Score card shows Score 1,830, Tier Epic, Rating 65 of 100 plus five bars; three score concepts at equal weight. Fix: slim rail block, drop Tier/Rating rows, keep sparkline and bars.
Tell 11/dead space: planner 1440 left column ends ~1295 while rail runs to ~1510; planner-you and planner-guide leave ~400px empty under Point order; character right rail ends at 915 with a void beside Best per encounter and Activity. Fix: move Guide builds or Share under Point order, or let the left blocks span full width.
Tell 22: character owner 390, four yellow headline links stacked, two lines each, no separation ("Upgrade: Devout Skirt, Stratholme: Baron Rivendare ›"). Fix: one line each on a plate, truncated.
Tell 21: Activity rows repeat the character's own name on their own page (all character pages). lib/character-profile.ts. Fix: start with the verb, "Uploaded Molten Core, 10 bosses in 92 min".
Tell 13 (borderline): planner blank 1440, 60 dashed empty squares with "No points spent yet." in Point order. Fix: collapse until a point is spent, or show the level 60 build order dimmed.
Tell 17: "Loaded" pill on Survival in Guide builds (planner-guide). Real state, keep as the only pill.

## Authentic, keep
Paperdoll with real icons and quality colours; upgrade list with item, slot and HPS gain; +128.9 HPS as the lead number; Best per encounter with parse colour and HPS; planner trees with real icons and ranks; point-order icon strip with per-level rank; guide builds named with split (31/20/0); class-coloured title and crest; Send this build to Felsnap; no gradient or glow buttons, no eyebrows, no icon-per-bullet.

## 390
Character: paperdoll becomes a named list (item, slot), dense and good; Score and Builds seen sink to the bottom correctly. Planner: tree becomes a spec tab strip with point counts; fine. Same copy findings apply.

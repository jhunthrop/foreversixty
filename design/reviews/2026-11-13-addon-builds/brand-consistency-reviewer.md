# Brand consistency review: addon Builds pane and panel Talents row

Reference: web/src/components/planner/GuideBuilds.svelte, lib/planner/copy.ts, lib/planner/guide-builds.ts (source read; live page not fetched).
Boards: builds-hunter-42, builds-druid-24-site, builds-row-states, panel-talents-42 viewed.

VERDICT: PASS. Must-fixes: none.

Matches the web: (1) labels word for word, "Level 60 build" and "Raid build" (copy.ts guideLevelingBuild/guideRaidBuild), with "Level 60 and raid build" a sound extension. (2) Row anatomy: spec's client icon, spec name, label, split, one action. (3) Gold reserved for the one loaded/followed mark, as the web's gold Tag. (4) Class art is the ringed HD crest (Theme.buildCrest/paintCrest) in the heading and the site rows, no square class icons; spec icons are game icons, as on the web. (5) Colours and the Follow button come from Theme/Widgets tokens (Theme.inlineColor, Theme.HEX, Widgets.clientButton); no hardcoded colours in BuildsPane.lua. (6) Sentence case, no em dashes, strings from Locale.lua keys. (7) Panel Talents row: title, talent, gold "Spend 1 now", muted build name; the blank button is fixed and the cut text is gone.

Allowed (platform idiom): stacked label/split lines instead of the web's one-line "label · 31/16/4"; client-style blue Follow button; cursor edge and Enter-to-follow; "Following" where the web says "Loaded" (the addon follows, the planner loads).

NICE (no block):
1. Split shows tree names ("Beast Mastery 31 / Marksmanship 16"), the web shows "31/16/4". Fine in game; keep the biggest-first order the web uses.
2. Druid: row title "Feral" but split says "Feral Combat" (client tree name); consider the spec's display name.
3. Druid board: "and 3 more" sits flush on the paste box; add the 6 px gap.
4. Panel still says "5 points off plan" while the web says "5 points off"; align if that copy is touched.
5. Why lines cut mid-word ("take t...") on the 12 px line; cut at a word boundary.

## Round 2

VERDICT: PASS. Must-fixes: none.

- Labels: "Leveling build" / "Raid build" now differ from web copy.ts guideLevelingBuild ("Level 60 build"). ALLOWED as a deliberate in-game ruling (wow-player: wrong for a level 42 player); the planner keeps its words. Keep it in addon Locale only (L.dialogKindLeveling), never fork the shared copy, and note the divergence in the contract doc. "Raid build" is identical.
- Followed row: gold edge and warm ground on the followed row only; the cursor edge is the neutral body colour (BuildDialog.lua 6.4 comment), so gold still means one thing, as the web's gold Tag. Colours via Theme.HEX, no hardcoded hex found.
- Site-row crest 20 px (BuildDialog.SITE_CREST) is the ringed HD crest via Theme.buildCrest; no square art. Tooltip is title, muted split, why, muted "Click to follow": same order as the web row, native idiom.
- Round 1 NICE: the "5 points off plan" wording (web: "5 points off"), the tree-name split order and the word-boundary cut are unchanged or fine; still NICE, not blocking.

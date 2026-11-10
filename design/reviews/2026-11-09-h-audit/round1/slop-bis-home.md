# AI-slop review: bis landing (out, 60), bis spec, home (out, 60)

Sources: web/src/components/bis/{BisLandingAnswer.astro,BisSide.svelte,BisItinerary.svelte,BisSlotLines.svelte,BisSpecPage.svelte}, lib/bis/{landing-build.ts,frame-view.ts}, components/home/{HomeClassPicker,HomeHowToStart,HomeFeed,HomeGrid,HomeCharacter}.svelte, components/RaidsCountdown.svelte, lib/home/raid-facts.ts.

No visual tell 1, 2, 5, 12 or 15 on any page. No gradients, glows, glass, em dashes. Remaining slop is label redundancy (22), oversize secondary numerals (16), repeated placeholder rows.

## bis-landing-out: PASS WITH FIXES
- Reads as a player tool: icon-led spec rows, class colours, no pitch.
- T8: nine identical plates in a symmetric 3x3. Acceptable because it is a real list; add no more chrome.
- T13: blank dark icon squares (Beast Mastery, Marksmanship, Protection, Fire, Affliction...). Fixture artefact per brief; confirm real build never ships a blank square.
- T6/22: "Level range" label left of the tabs (1440). Tabs are self-evident; drop it.
- T7: "Which spec is strongest? Tier list" aside (1440). Make it "Tier list >" only.

## bis-landing-60: PASS WITH FIXES
- T22: "60: full raid-ready, Phase 1 gear." under tabs (1440, 390), lib/bis/landing-build.ts. Restates the selected tab, odd grammar. Delete.
- T16: oversize "4" in "4 of 17 slots" with bar in the title band (BisLandingAnswer.astro). Normal weight numeral, keep the bar.
- T17: "Yours" pill on Fury; on 390 it wraps to its own line and breaks row height. Use the accent bar plus inline tag, or colour only.
- 390: "Fury gear for 60" button takes its own row under the bar; put it inline.
- T13: same blank icon squares.

## bis-spec: PASS WITH FIXES
- Strongest page: real paperdoll, tooltip, weights, itinerary with real item names.
- T22: "Same DPS" on 12 of 16 itinerary rows (BisItinerary.svelte). Show a dash for ties; keep "+4.0 DPS" for real gains.
- T22: "1.00 = 0.891 DPS per point of Agility" (lib/bis/frame-view.ts). Use "1.00 = 0.891 DPS".
- T22: "Crafted · Leatherworking" stop header, then each row repeats "Leatherworking (skill 300)". Trim row to "Head · skill 300".
- T22: tooltip "Next: -1.8 vs ..." (BisSlotLines.svelte) cryptic; name the item and delta only.
- Layout: "Pre-raid 60 · Alliance 1083.6 DPS" overflows the plaque's right edge (1440).
- Layout: filter strip wraps with "All sources" orphaned on a second row (1440); group into intended rows.
- 390: filter summary truncated mid-word ("Raid-rea..."); fix.
- "Where to go" right-aligned bare counts (2, 2, 1) have no noun; use "2 pieces" or drop.

## home-out: PASS WITH FIXES
- Concrete feed lines, real zones, real countdown.
- T8: How to start plate has an inner inset frame (HomeHowToStart.svelte). Remove inner frame.
- Duplicate CTA: "Sign in with Battle.net" is both the top button and step 1 (visible together on 390). Remove one.
- T13: "Loot not out yet" twice in the raid rows (lib/home/raid-facts.ts). Show boss count or drop those rows.
- Dead band under the feed on 1440 (about 190px before footer); balance column heights.
- Title band "Pick your class" with lone button reads thin; fine.

## home-60: PASS WITH FIXES
- Densest, most authentic page: real sims, upgrades, guild rank, <Olympus XXVII>.
- T16: four equal big numerals "33 8 18 3" in the Onyxia plate (RaidsCountdown/HomeGrid). Countdown is the point; shrink signed up and late.
- T22: "+48.1 DPS in 13 upgrades" footer next to "See all 13" (HomeGrid.svelte). Drop "in 13 upgrades".
- T22: "(Fury Warrior, 60 raid)" parenthetical in feed lines (HomeFeed.svelte). Use "#2 of 14 Fury Warriors".
- T22: upgrade rows run three lines, "over Lieutenant Commander's Plate..." truncated at 1440. Drop slot word where the item name implies it, let "over" wrap.
- T17: solid orange "S" tile fights the tan background; shrink.

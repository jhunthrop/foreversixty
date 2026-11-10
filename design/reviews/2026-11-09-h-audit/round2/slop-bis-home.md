# AI-slop review round 2: bis landing (out, 60), bis spec, home (out, 60)

Captures: round2-captures/, 1440 and 390. Handoff items (segmented.css band wrap, GuideSpecRow "Yours" pill at 390, server feed line "(Fury Warrior, 60 raid)") are noted, not failed on.

First impression: reads as a player tool on every page. No gradients, glows, glass, em dashes, eyebrows or hero slogan. Round-1 copy fixes mostly landed.

## Fixed since round 1 (verified)
- bis-landing: "Level range" label gone; aside is now "Tier list >" only; "60: full raid-ready" line gone.
- bis-spec: "Same DPS" rows gone (ties show nothing), "1.00 = 0.891 DPS", "Head . skill 300" rows, tooltip "-1.8 DPS vs Lieutenant Commander's Chain Helm", plaque header fits ("Pre-raid 60"), Where to go counts now "2 pieces".
- home: "Loot not out yet" replaced by "1 boss . 22 drops"; "in 13 upgrades" footer gone; feed dead band shrunk to ~100px.

## bis-landing-out: PASS
- Only fixture blank icon squares left (T13, artefact per brief; confirm the real build never ships one).

## bis-landing-60: PASS WITH FIXES
- T16: "4" in "4 of 17 slots" still a large bold numeral beside the bar (1440, 390), BisLandingAnswer.astro. Normal weight, same size as "of 17 slots".
- T17 (handoff f-selected): "Yours" pill on Fury; at 390 it wraps to its own line and grows the row.
- 390: "Fury gear for 60" still on its own row under the bar; BisLandingAnswer.astro, put inline with the count.

## bis-spec: PASS WITH FIXES
- 390: filter summary "Marksmanship . 60 . Alliance . Raid-rea..." still truncates mid-word, BisSpecPage.svelte; drop default values or wrap instead of clip.
- 390: "Theramore Expeditionary Force reputation" row wraps its count to "2 / pieces" in Where to go, BisSpecPage.svelte; nowrap the count, let the name wrap.
- 1440: "All sources" select orphaned on a second row (handoff, segmented.css).
- Minor: Where to go repeats every group heading the itinerary prints again below; keep one.

## home-out: FAIL (one blocking tell, T8)
- T8: How to start plate still has an inner 1px inset frame inside the rounded plate (both widths), HomeHowToStart.svelte. Remove the inner border; it is a card inside a card.
- Duplicate CTA: step 1 "Sign in with Battle.net" repeats the top button; at 390 both are on screen together, HomeHowToStart.svelte / home title band. Remove one.
- Onyxia's Lair has a sub-line, Barrow Deeps and Hyjal Summit have none; lib/home/raid-facts.ts, give all three a boss count or none.
- Title band "Pick your class" with lone button reads thin at 1440; acceptable.

## home-60: PASS WITH FIXES
- T16: four equal large numerals "33 8 18 3" in the Onyxia plate (1440, 390), RaidsCountdown/HomeGrid. Keep days and hours large; set signed up and late at body size.
- T22 (handoff, API): "(Fury Warrior, 60 raid)" in feed lines; server-side to "#2 of 14 Fury Warriors".
- T17: solid orange "S" tile in Your spec this week (both widths) fights the tan ground, HomeGrid.svelte; shrink to a coloured letter.
- 1440: "Before the pull" half of the Onyxia plate is two lines over a large empty art panel; fill with real prep state or narrow the panel.

## Authentic, keep
Icon-led class rows on plates in class colours, real item names and paperdoll on bis-spec, guild sims with class-coloured bars, real raid names and countdown, short player-voice feed lines.

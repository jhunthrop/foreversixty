# AI-slop review round 3: bis landing (out, 60), bis spec, home (out, 60)

Captures: round3-captures/, 1440 and 390. Handoffs (segmented.css band wrap, GuideSpecRow "Yours" pill at 390, server feed line "(Fury Warrior, 60 raid)") not failed on. Fixture blank icons ignored.

First impression: reads as a player tool. No gradients, glows, glass, em dashes, eyebrows or slogan.

## Verified fixed since round 2
- bis-landing-60: "4 of 17 slots" is body size now; no big numeral.
- bis-spec 390: filter summary fits ("60 . Alliance . Raid-ready"); Theramore row wraps the name, count stays "2 pieces".
- home-out: inner inset frame gone from How to start; lone title-band button gone (no duplicate CTA); Barrow Deeps, Hyjal Summit, Onyxia's Lair all consistent (no sub-lines).
- home-60: only days and hours are large; signed up / late / tanks etc. at body size. "S" is a coloured letter, not a tile.

## Verdicts
- bis-landing-out: PASS
- bis-landing-60: PASS (nit: 390 "Fury gear for 60" sits on its own row under the count and bar; BisLandingAnswer.astro, optional inline)
- bis-spec: PASS (nit: Where to go list and the itinerary below repeat every group heading; BisSpecPage.svelte, keep one)
- home-out: PASS
- home-60: PASS WITH FIXES

## home-60 remaining
- 1440: "Before the pull" half of the Onyxia plate is two checklist lines over a ~150px empty art panel (dead panel). HomeGrid.svelte / RaidsCountdown: narrow the art column or add real prep rows; do not stretch the plate to the countdown height.

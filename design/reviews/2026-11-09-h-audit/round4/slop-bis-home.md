# AI-slop review round 4: bis landing (out, 60), bis spec, home (out, 60)

Captures: round4-captures/, 1440 and 390. Handed-off items (segmented.css "All sources" wrap, GuideSpecRow "Yours" pill at 390, server feed line "(Fury Warrior, 60 raid)") and fixture blank icons not failed on.

First impression: reads as a player tool. No gradients, glows, glass, em dashes, eyebrows, slogan or invented stats. Real item names, quality colours, class colours by role.

## Verified fixed since round 3
- home-60 1440: the dead art panel beside "Before the pull" is gone; the checklist sits in the plate beside the countdown, no stretched empty column.
- home-60 390: Before the pull stacks under the buttons, nothing dead.
- bis-landing-60: "4 of 17 slots" stays body size at both widths.
- home-out: no sub-lines on raids, no duplicate CTA.

## Verdicts
- bis-landing-out: PASS
- bis-landing-60: PASS (nit: 390 title band is a left-aligned button over ~50px of empty band; BisLandingAnswer.astro, tighten bottom padding)
- bis-spec: PASS (nit: "Where to go" list and the itinerary below repeat every group heading, BisSpecPage.svelte, keep one; nit: 1440 left column ends under the doll with ~100px dead space beside the long list)
- home-out: PASS (nit: 1440 dead band under Latest on Forever; "How to start" plate padding deeper than its two rows need; HomeGrid.svelte)
- home-60: PASS (nit: 1440 "Your spec this week" column ends ~60px short of Guild sims; the "S" rank letter is acceptable, the rank is the point)

No tells from the list found on any of the five pages at 1440 or 390.

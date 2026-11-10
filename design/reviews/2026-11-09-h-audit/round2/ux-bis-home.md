# ux-designer review, round 2: BiS and home (2026-11-09)

Captures: round2-captures/{bis-landing-out,bis-landing-60,bis-spec,home-out,home-60}-{1440,390}.png. Reference: round1/ux-bis-home.md, handoffs.md, directions spec.

## Verdicts
- bis-landing-out: SHIP
- bis-landing-60: SHIP
- bis-spec: SHIP WITH FIXES
- home-out: SHIP WITH FIXES
- home-60: SHIP

## Round-1 items
1. bis-spec paperdoll title bar: FIXED. Title, "Pre-raid 60" and "1083.6 DPS" sit inside the bar at 1440; 390 wraps the title to two lines inside the bar.
2. bis-spec control band: NOT FIXED, owned by f-selected (segmented.css). "All sources" still wraps to row two at 1440; no group gaps; Survival alone has an icon; Marksmanship underline starts left of its label. Noted, not failed.
3. "Same DPS": FIXED. Gain shows only on the four rows that have one, with unit.
4. Tooltip "Next": FIXED, now "-1.8 DPS vs Lieutenant Commander's Chain Helm" with unit. Polish: a loss still uses the yellow gain slot; use muted.
5. home-out hero band: NOT FIXED. At 1440 "Pick your class" sits 17 px under the nav (y 97 vs nav end 80); band about 100 px; art barely shown. Same at 390 (title at y 92, nav ends 64).
6. home-60 title band: IMPROVED to an acceptable compact band (name at y 111, 31 px under nav; DPS and sim stamp right). Passes the compact-band rule.

## Must fix
- bis-spec 390, header strip: "Marksmanship · 60 · Alliance · Raid-rea..." truncates the raid-ready value. Show it on two lines or drop the redundant "Marksmanship" (the doll title names it). Source: web/src/components/bis/BisSpecPage.svelte (mobile summary strip).
- home-out 1440 and 390, title band: 17 px under nav, art unseen. Min-height 160 px (1440) / 120 px (390), title bottom-aligned, 24 px under nav. Source: web/src/pages/index.astro, web/src/components/home/HomeGrid.svelte.

## Polish
- home-out 390: raid plaque still sits under "How to start"; put it first after the class list. 1440: left column ends 913, rail 1018, 105 px dead gap; step 1 repeats the Sign in button in the band.
- bis-landing-60 1440 and 390: "4 of 17 slots" has no caption; at 390 the primary "Fury gear for 60" sits under the progress bar, put it first. Source: web/src/pages/bis/index.astro.
- bis-landing both: empty dark icon squares (Protection, BM, MM, Retribution, Holy, Fire, Affliction, Demonology). Fixture artefact per brief; still reserve a class-crest fallback.
- bis-spec 1440: "Where to go" (10 rows) is repeated verbatim as section headings below; the itinerary rows have a 1000 px empty middle with gain far right. Make the summary jump links or drop it.
- bis-spec 1440: no tinted itinerary row for the selected slot; tooltip text about 11 px (mock 13 px).

## Handoff, not failed
- BiS control band wrap and underline (f-selected, segmented.css); "Yours" pill wraps under Fury at 390 on bis-landing-60 (f-selected, GuideSpecRow); feed line "(Fury Warrior, 60 raid)" (API).

## Passes
- No em dashes on any page. Nav, ground, plates, Marcellus titles match round 5.
- home-60: plaque leads (33 days, 8 hours, 18 signed up, 3 late, roles, I'm in / Late / Can't come); upgrades carry quality colour, slot, source, gain with unit, no truncation; guild bars in class colour, own row selected; spec card with tier S.
- bis-landing: compact title band, class names in class colour, crests, Yours tag, tier list link; the stray "60: full raid-ready" line is gone.
- bis-spec 390: no horizontal scroll; two-column slot list, tooltip under it.

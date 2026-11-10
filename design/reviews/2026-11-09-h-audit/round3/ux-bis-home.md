# ux-designer review, round 3: BiS and home (2026-11-09)

Captures: round3-captures/{bis-landing-out,bis-landing-60,bis-spec,home-out,home-60}-{1440,390}.png. Compared against round2/ux-bis-home.md and handoffs.md.

## Verdicts
- bis-landing-out: SHIP
- bis-landing-60: SHIP
- bis-spec: SHIP WITH FIXES
- home-out: SHIP
- home-60: SHIP

## Round-2 must-fixes
1. bis-spec 390 header strip: FIXED. Reads "60 · Alliance · Raid-ready" with Change; no truncation. Doll title wraps to two lines inside the bar (Marksmanship / Hunter).
2. home-out title band: FIXED. "Pick your class" now sits about 130 px under the nav at 1440 (y 211) and 93 px at 390 (y 157); background art shows above it. Step 1 no longer repeats a Sign in button.

## Remaining (polish, none fail a page)
- bis-spec 1440: "Where to go" (10 rows) is still repeated verbatim as the section headings below. Make the rows jump links or drop the summary. Source: web/src/components/bis/BisSpecPage.svelte.
- bis-spec 1440: item rows below the doll have a 900 px empty middle with the gain at the far right edge; cap the row width or bring the gain next to the name.
- bis-spec 1440 tooltip: still about 11 px text; the loss "-1.8 DPS" still uses the yellow gain slot (use muted).
- home-out 390: raid plaque still sits under "How to start"; put it directly after the class list. 1440: left column ends y 1000, rail y 1105, about 100 px dead gap under Latest on Forever.
- bis-landing-60 390: "Fury gear for 60" still sits under the progress bar; put it first. About 50 px dead band under the button.
- bis-landing both: empty dark icon squares (Protection, BM, MM, Retribution, Holy, Fire, Affliction, Demonology). Fixture artefact; still reserve a class-crest fallback.

## Handed off, not failed
- BiS control band wrap on "All sources", group gaps, Survival-only icon, underline offset (f-selected, segmented.css).
- "Yours" pill wrapping under Fury at 390 on bis-landing-60 (f-selected, GuideSpecRow).
- Feed line "(Fury Warrior, 60 raid)" on home-60 (API).

## Passes
- No em dashes. Nav, ground, plates and Marcellus titles consistent across all five pages at both widths.
- home-60: compact name band with DPS and sim stamp; plaque leads; upgrades with quality colour, slot, source, gain with unit; guild bars in class colour, own row selected; spec card with tier S. No horizontal scroll at 390.
- bis-landing: compact band, class-coloured names on plates, crests, Yours tag, tier list link, progress with unit.
- bis-spec: doll title bar holds title, "Pre-raid 60" and DPS; gains only on rows that have one.

# ux-designer review, round 4: BiS and home (2026-11-09)

Captures: round4-captures/{bis-landing-out,bis-landing-60,bis-spec,home-out,home-60}-{1440,390}.png. Compared against round3/ux-bis-home.md and handoffs.md.

## Verdicts
- bis-landing-out: SHIP
- bis-landing-60: SHIP
- bis-spec: SHIP
- home-out: SHIP
- home-60: SHIP

## Round-3 items
- bis-spec "Where to go" repeated as headings: ACCEPTED, rows are jump links by design.
- bis-spec 390: doll list, tooltip block, "Open in the sim" and gains with unit read in order; no horizontal scroll.
- bis-landing-60 390: "Fury gear for 60" now sits directly under the progress bar and the dead band is gone. Still not first, polish.
- home-60 390: upgrades now lead the page, then Latest, Guild sims, spec. Good.

## Polish (none fail a page)
- bis-spec 1440: item rows keep a ~900 px empty middle with the gain at the far right; cap the row width or move the gain beside the name. Source: web/src/components/bis/BisSpecPage.svelte.
- bis-spec 1440 tooltip: ~11 px text; "-1.8 DPS" loss should use the muted slot, not the gain slot.
- home-out 390: raid plaque still follows "How to start"; put it right after the class list. 1440: left column ends y 1000, rail y 1105, ~100 px dead gap under Latest on Forever.
- bis-landing both: empty dark icon squares (Protection, BM, MM, Retribution, Holy, Fire, Affliction, Demonology); fixture artefact, keep a crest fallback.

## Handed off, not failed
- BiS control band wrap on "All sources", Survival-only icon, Bare/Night Elf control (f-selected).
- "Yours" pill wrap at 390 (GuideSpecRow). Feed line "(Fury Warrior, 60 raid)" (API).

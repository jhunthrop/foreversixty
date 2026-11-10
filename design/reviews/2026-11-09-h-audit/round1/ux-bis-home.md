# ux-designer review, round 1: BiS and home (2026-11-09)

Captures: design/reviews/2026-11-09-h-audit/first/{bis-landing-out,bis-landing-60,bis-spec,home-out,home-60}-{1440,390}.png
Reference: design/specs/2026-11-09-directions.md (round 5), design/mocks/directions/png/g-home-*, g-bis-frame-*.

## Verdicts
- bis-landing-out: SHIP WITH FIXES
- bis-landing-60: SHIP WITH FIXES
- bis-spec (/bis/hunter/marksmanship/60): SHIP WITH FIXES
- home-out: SHIP WITH FIXES
- home-60: SHIP WITH FIXES

## Must fix
1. bis-spec 1440, paperdoll title bar: "1083.6 DPS" runs past the right edge of the frame's title bar (the "DPS" unit sits outside the plate, over the page ground). Spec: title and DPS on one line inside the bar. Fix: shrink/ellipsis the title ("Marksmanship Hunter BiS" + "Pre-raid 60 · Alliance" is longer than the mock's) or drop the "Pre-raid 60 · Alliance" meta to the band. Source: web/src/components/ui/Paperdoll.svelte, web/src/lib/bis/frame-view.ts.
2. bis-spec 1440, control band: one run of 12 undivided controls (spec, range, faction, preset, Night Elf select) plus "All sources" wrapped alone to a second row at the left. Survival alone carries a spec icon; Beast Mastery and Marksmanship none; the Marksmanship underline sits off its label. Mock: one row, three groups (spec, range, faction) with gaps, no icons. Fix: group with 24 px gaps, move preset/race/source into one right-aligned cluster, drop or equalise the icon. Source: web/src/components/bis/BisSpecPage.svelte, BisSources.svelte.
3. bis-spec 1440 and 390, itinerary: every tied pick shows "Same DPS" (15 rows at 1440). Spec: gain only, blank when none. Word plus unit repeated is noise and breaks "a number takes its unit and nothing else". Fix: render nothing for ties; keep "+x.x DPS". Source: web/src/components/bis/BisItinerary.svelte (header comment names "Same DPS").
4. bis-spec 1440 and 390, selected-item tooltip: "Next: −1.8 vs Lieutenant Commander's Chain Helm" has no unit, in the tooltip's yellow gain slot. Spec: tooltip ends "+x.x DPS". Fix: "−1.8 DPS vs ..." muted, or omit. Source: tooltip text builder, not located; start at web/src/lib/bis/doll-view.ts.
5. home-out 1440 and 390, hero band: art band about 100 px tall; "Pick your class" and "Sign in with Battle.net" sit 17 px under the nav. Mock band is 230 px, title at its foot, art visible. Fix: band min-height 200 px (1440) / 150 px (390), title bottom-aligned, 24 px under the nav. Source: web/src/pages/index.astro, web/src/components/home/HomeGrid.svelte.
6. home-60 1440 and 390, title band: same crowding (name at 97 px, nav ends 80), art barely visible. Same fix. Source: web/src/components/home/HomeCharacter.svelte.

## Should fix (polish)
- home-60 1440, Your upgrades: "over Lieutenant Commander's Plate…" truncates; allow two lines or widen rail (HomeUpgrades.svelte ~94-101).
- home-out 1440: left column ends 913 px, rail 1018 px; 105 px dead gap. "Loot not out yet" twice in the plaque is filler.
- home-out 390: raid plaque (the lead) sits 960 px down, under the class list and How to start. Move it above How to start.
- bis-landing 1440: title band 166 px before the control bar; spec says compact. "60: full raid-ready, Phase 1 gear." floats alone under the bar.
- bis-landing-60 390: "4 of 17 slots" has no caption ("Best-in-slot pieces you wear", spec 2026-11-07 s3) and the primary button sits under the figure; put button first.
- bis-landing both widths: empty dark icon squares on several specs (BM, MM, Fury Prot, Holy, Fire, Affliction...). Brief calls blank icons fixture artefacts, so unscored; still reserve a fallback (class crest), never an empty square. Source: web/src/pages/bis/index.astro.
- bis-spec 1440: no itinerary row tinted for the selected slot (mock does). Tooltip text about 11 px, mock 13 px.

## Passes
- Nav, seal, Premium teal, Marcellus titles, ground and 10 px panels match round 5 on all pages.
- Home signed in: plaque leads (33 days, 8 hours, I'm in / Late / Can't come, roles); upgrades carry quality colour, slot, source, gain with unit; guild sims bars in class colour, own row selected.
- BiS landing 60: "Fury gear for 60" primary, Yours tag, class names in class colour, round crests.
- BiS spec 390: header strip, two-column slot list, tooltip under it, no horizontal scroll.
- No em dashes on any of the five pages.

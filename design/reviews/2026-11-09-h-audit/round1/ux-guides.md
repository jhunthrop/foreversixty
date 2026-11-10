# UX review, guides family, round 1 (2026-11-09, ux-designer)

Captures: design/reviews/2026-11-09-h-audit/first/guides{,-you,-class,-spec}-{1440,390}.png. Reference: design/specs/2026-11-09-directions.md (round 5), design/specs/2026-11-08-guides.md, owner rules in BRIEF.md. Blank spell/item icons are fixture artefacts and are not findings except where the markup shows a hole.

## /guides (signed out): SHIP
- Compact title band (Marcellus "Guides", Updated stamp), nine class plates with crest, class-coloured name and spec rows; every spec reachable in 875 px at 1440. Phone: title then Warrior at y 165.
- Polish: Mage and Warlock plates pad to Druid's height, leaving ~50 px of dead plate under Frost and Destruction (web/src/components/guides/ClassSpecPicker.astro). Align to content, or fill the row.

## /guides (signed in, level 34 hunter): SHIP
- Answer sits right of the title at 1440 (name in class colour, "Level 34 · Dwarf Beast Mastery Hunter", yellow "Your guide: Beast Mastery Hunter ›"), spec row marked "Your spec" with the yellow edge. At 390 it stacks under the title and "Your spec" wraps to its own line under Beast Mastery, making that row 70 px against 44 px for the others.
- Polish: at 390 keep the "Your spec" pill on the name line (right-aligned) so the row height does not change (GuideSpecRow.astro).
- Polish: the header character and the answer repeat the same name and level at 1440 (nav chip plus answer). Acceptable; the answer carries the next step.

## /guides/hunter: SHIP WITH FIXES
- Must fix: Races panel shows browser broken-image glyphs for all eight race icons at 1440 (right rail). Source: web/src/components/guides/GuideRaceList.astro -> ui/RaceIcon.svelte. If it is a fixture 404, RaceIcon still needs a no-flash fallback (the faction-coloured disc, as the spec page's racials show a plain square) so a missing file never draws the browser glyph.
- Polish: each spec row reads "DPS" under the name. A veteran knows a hunter spec is damage; owner rule "don't explain the obvious". Show a tag only where it informs (Survival "Melee", ruling Q8 in the 11-08 spec; the page comment says the owner dropped role tags, so decide and apply it consistently).
- Page is dense enough (specs, what changed, races, sources, 1,399 px). The three spec panels are 95 px tall with one link each; fine as rows.

## /guides/hunter/survival: SHIP WITH FIXES
What matches: At a glance plate is the first thing under the band at 1440 (y 323-737) and 390 (tree, Load and Sim, rotation, stats in the first two screens); teal Load/Sim buttons; weights in the Agility = 1.00 convention with hit cap and unit lines; Marcellus section heads; TOC rail left; quality-coloured BiS on a plate; Sources only in the footer block; no em dashes seen.
- Must fix (1440, Talents, Best in slot, Enchants and consumables, Races, Leveling): the prose column is short and the lane column long, so a 320-1,100 px empty strip sits under the prose on the plate. Worst: Best in slot (prose ends y 636 of a 1,300 px lane), Enchants (prose ends 410, lane ends 1,085), Leveling (lane tree ends mid-section, right column empty for the last 700 px). This is the "thin page / dead plate" the owner rules fail. Fix: make the lane `position: sticky` under the nav so it rides beside the prose, and/or give the prose column the lane's height by putting Best in slot and Enchants lists in two columns at 1440. Source: web/src/components/guides/GuideBisLane.astro, GuideExtrasLane.astro, GuideRacesLane.astro, GuideBuildsLane.astro and the shared reading-section layout in pages/guides/[class]/[spec].astro.
- Must fix (1440 and 390, Enchants and consumables): rows carry no icon at all where `row.icon` is undefined (GuideExtrasLane.astro lines 53 and 95 render an invisible span), so the rows read as indented text, against 2026-11-08 MUST 6 and "never an empty icon square" (races show a visible bordered square, these show nothing). Either always resolve an icon or render the same bordered placeholder as GuideSpellRow. Confirm real data has icons; if so this is a fixture artefact and drops to polish.
- Should fix (1440): Best in slot repeats the full 17-slot list that the BiS page owns (one job per page; ruling Q4 asked for one band plus "Full list and alternatives ›"). It is 1,300 px tall. Cut to the row of icons (a paperdoll-style strip, or two columns) with the link.
- Should fix (both): the Builds lane repeats Load this build / Sim this build as text links directly after the At a glance buttons. One primary per region; keep the link in the lane only if the plate is out of view.
- Polish (390): the Leveling slider's 3 trees are 390 px wide with 14 px nodes; touch targets are small, hover tooltips need tap parity (verify).
- Polish (390): "On this page" bar reads "Survival Hunter · On this page" with "Change" at the right; "Change" does not say what changes. Label it "Jump to".
- Polish: stat table's "Attack speed / Not on any item / –" row stays in the full table, which the 11-08 spec allows.

## Acceptance still to attach
390 and 1440 with the Leveling tab/slider moved, the signed-in level 34 Survival variant, and a hover state on one BiS row (tooltip).

# UX review, guides family, round 2 (2026-11-10, ux-designer)

Captures: round2-captures/guides{,-you,-class,-spec}-{1440,390}.png. Baseline: round1/ux-guides.md. Items in handoffs.md (f-selected, f-content) are noted, not failed.

## /guides: SHIP
- Unchanged and correct: compact Marcellus band, nine class plates, crest, class-coloured names, every spec in view at 1440 (to y 875) and 390 (Warrior at y 165).
- Polish (open from round 1): Mage and Warlock plates still pad to Druid's height, ~50 px dead plate under Frost and Destruction (1440). web/src/components/guides/ClassSpecPicker.astro: align plates to content (align-items: start) or fill with a fourth row.
- Blank icons beside Beast Mastery, Marksmanship, Protection and others are fixture artefacts.

## /guides signed in (level 34 hunter): SHIP
- Answer block right of the title at 1440 and stacked at 390, spec row marked with yellow edge.
- Open, owned by f-selected (GuideSpecRow.astro): at 390 "Your spec" still wraps under Beast Mastery, row 70 px against 44. Noted.
- Open, owned by f-content (lib/guides/copy.ts): "Your guide:" prefix. Noted.

## /guides/hunter: SHIP
- Fixed since round 1: race icons are now bordered faction-tinted squares (no browser glyph) at 1440 and 390; the "DPS" tag under each spec name is gone.
- Dense enough: specs, what changed, races, sources.
- Polish (1440): the Races rail ends at y 740 while Sources runs to y 1,100 beside an empty right column. Let Sources span full width or move Races below the spec row beside What changed so the rail is not a stub (web/src/pages/guides/[class]/index.astro).
- Polish: the three spec panels are 95 px with one link; fine.

## /guides/hunter/survival: SHIP WITH FIXES
Held from round 1: At a glance plate first under the band (1440 y 323-737; 390 first two screens); teal Load/Sim; Agility = 1.00 weights; Marcellus heads; TOC rail; quality-coloured BiS; Sources in the footer block; no em dashes seen.

Remaining blockers and must-fix (own lane):
- Must fix (1440 and 390): broken-image glyphs, not blank squares, on spell icons: Key talents (Lightning Reflexes, Resourcefulness), Rotation rows (Aspect of the Beast, Strider Kick), Cooldowns (Rapid Fire), and the talent trees' hit-nodes. The glyph is the browser's, so the fallback is missing. Source: ui/SpellIcon.svelte and guides/GuideSpellRow.astro. Fix: render the same dark bordered placeholder every other missing icon gets, with onerror swapping to it, so no glyph ever shows.
- Must fix (390): the sticky bar reads "Survival Hunter · On this page" with "Change". Source: web/src/components/guides/GuideToc*.astro. Fix: label it "Jump to".
- Should fix (390, Talents): the Builds lane shows "Load this build  Sim this build" as text links 700 px below the same two buttons in At a glance. Drop the lane copy at 390 (also f-content).
- Polish (1440, Enchants): "Enchant Weapon - Agility" appears twice, Weapon and Off hand. Merge to one row "Enchant Weapon - Agility, both weapons". Source: content/guides/hunter/survival.md enchant list.

Open, owned by f-content (noted, not failed):
- Dead plate beside the prose, 1440: Talents (prose ends y 1,690, lane to y 2,240), Stat priority (prose ends ~y 250 of that section, table runs to 600), Best in slot (prose ends ~y 840, lane 1,450), Enchants and Races (prose stops ~700 and ~1,100 px above the lane end), Leveling (lane tree ends mid-section; right column empty for the last ~1,000 px). Sticky lane or two-column lists still not done.
- Best in slot still prints all 17 slots (about 1,500 px) and repeats the BiS page; at 390 it also sits before its own prose.
- Enchants and consumables rows carry no icon at 1440 and 390 (invisible span, GuideExtrasLane.astro).
- Racials: tooltip text pasted verbatim (race-table.ts); racial icons are empty squares (fixture).
- Lede "Survival is a melee spec in Forever: ..."; "DPS · Updated 7 Nov" in the band; tier-gate and Client-build strings are in other lanes.
- Footer "Numbers from client build 1.60.1.70790, checked against launch logs as they come in." (sources block, 1440): meta talk; f-content to cut or move to the Updated stamp.

## Checklist
- [x] Compact title bands, no big header, all four pages
- [x] Race icon fallback on class page
- [x] No role tag where it does not inform
- [x] No em dashes seen
- [x] Plates for class and quality-coloured lists
- [ ] No browser broken-image glyph (spec page)
- [ ] "Jump to" label at 390
- [ ] Dead plate beside prose (f-content)
- [ ] Extras rows with icons (f-content)
- [ ] Best in slot cut to a strip plus link (f-content)
- [ ] Spec-row "Your spec" pill on the name line at 390 (f-selected)

## Acceptance still to attach
390 and 1440 with the Leveling slider moved, the signed-in level 34 Survival variant, a hover state on one BiS row.

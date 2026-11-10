# AI-slop review: guides, round 1 (2026-11-09)

Captures: first/guides, guides-you, guides-class, guides-spec at 1440 and 390.
Tell 23 (owner list, not in the base 22): a definitional blurb, a sentence whose job is to say what a thing is ("X is a Y") to someone who already knows.

## First impression
The index and class page are clean and read as a player tool (class-coloured plates, real spell and spec icons, Updated stamp, no hero, no eyebrow). The spec page is dense and authored. What gives it away is a handful of "X is a ..." sentences and blank icon squares.

## /guides (signed out): PASS WITH FIXES
- 13 Blank spec icons (Beast Mastery, Marksmanship, Protection x2, Retribution, Holy, Fire, Affliction, Demonology): empty squares read as unfinished. Fixture artefact if the build lacks the icon; if production lacks them too, fall back to the spec crest. ClassSpecPicker.astro, 1440 and 390.
- 8 Nine plates in a 3x3 grid. Defensible: it is a picker and the grid is uneven (druid has 4 rows). Do not add per-plate blurbs.
- No tell 23 here. Index has no blurbs. Keep it that way.

## /guides signed in (level 34 hunter): PASS WITH FIXES
- 17 "Your spec" pill on the Beast Mastery row (1440): real state, but the yellow left rule already marks it. Keep the rule, drop the pill. ClassSpecPicker.astro.
- 22 "Level 34 · Dwarf Beast Mastery Hunter" under the name: Beast Mastery implies Hunter and the name is class-coloured. Use "34 Dwarf Beast Mastery". GuideAnswer.astro.
- 22 "Your guide: Beast Mastery Hunter >": "Your guide:" is the qualifier; "Beast Mastery Hunter >" is enough under the character. GuideAnswer.astro.
- 13 Same blank icons as signed out, including the marked row.

## /guides/hunter (class): FAIL (tell 23), then PASS WITH FIXES
- 23 "What changed from 1.12" bullets are definitions (1440 and 390, main column, web/src/content/guides/hunter/index.md ~lines 29-35):
  - "Aimed Shot is a trained spell with a 2 second cast..." lead with the change: "Aimed Shot: 2 second cast, shares a 6 second cooldown with Multi-Shot, auto shots keep firing while you cast."
  - "Heart of the Lion is a hunter spell from level 1: ..." drop "is a hunter spell".
  - "Aspect of the Beast gives melee attack power, which makes it the Survival aspect." state the change only.
  - "Mail from 40. Hunters wear leather until 40, and good leather stays in the mix after it." cut to "Mail from 40."
  - "Trueshot Aura is ranged attack power only, 100 at 60" carries a number; keep.
- 23 Races rail prose "Alliance hunters are Dwarf, Night Elf and, new in Forever, Human. Horde hunters are Orc, Troll and Tauren. Skyborne can be hunters..." restates the icon list above it. Delete; mark Human "new" on its row. hunter/index.md line 39, rendered through [class]/index.astro and GuideRaceList.
- 21 Pets and Aimed Shot bullets are long semicolon chains; split or trim.
- 13 Blank Beast Mastery and Marksmanship icons on the spec plates; plates are mostly empty (name, "DPS", link). Put the point split (5/13/33) or top stat on the plate. [class]/index.astro.
- 22 "DPS" under every spec name when all are DPS: show role only for tank/healer.
- 8 Boxed bullets, boxed Sources, boxed Races stack as same-style plates; make Sources a plain list.

## /guides/hunter/survival (spec): PASS WITH FIXES
- 23 Lede "Survival is a melee spec in Forever: two weapons, Raptor Strike queued..." (1440 under band, 390 first paragraph). Start at the play: "Two weapons, Raptor Strike on cooldown, Mongoose Bite when it lights up, no pet in a raid." content/guides/hunter/survival.md line 48.
- 22 "DPS · Updated 7 Nov" in the band: DPS is implied for most specs; drop for DPS.
- 22 Racials list (Races section, 1440): tooltip text pasted verbatim ("Grants 10% critical strike chance for 15 seconds on a 3 minute cooldown", "Increases hit chance with spells and attacks by 1%") beside prose saying the same. Trim to "10% crit, 15 s / 3 min". lib/guides/race-table.ts.
- 13 Race and a few trait icons render as empty boxes (Races list, 1440 ~y 6650, and broken-image glyphs in the talent column). Fixture artefact only if data is missing.
- 17 "Recommended" pills on 3 of 8 races beside coloured names: keep one marker.
- Keep: hit cap breakdown, weight normalisation, "Full priority", build with ranks.

## Authentic, keep
Real talent trees with ranks, rotation with conditions ("When Lacerate is not ticking"), stat bars with numbers, hit cap breakdown, class-coloured names, compact title band, uneven plate heights, spell tooltips in prose, Sources with provenance. The rotation prose ("Walk in behind the boss with the rogues once the tank has it") is the best copy on the site.

## Verdicts
- /guides: PASS WITH FIXES
- /guides signed in: PASS WITH FIXES
- /guides/hunter: FAIL (tell 23)
- /guides/hunter/survival: PASS WITH FIXES

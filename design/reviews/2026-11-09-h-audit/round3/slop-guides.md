# AI-slop review: guides, round 3 (2026-11-10)

Captures: round3-captures/guides, guides-you, guides-class, guides-spec at 1440 and 390. Fixture blank icons and handed-off items (handoffs.md) do not fail a page.

## /guides: PASS
Plain class plates with spec rows, title band, no copy. Blank spec icons are fixture artefacts.

## /guides-you: PASS WITH FIXES
- 17 "Your spec" pill on Beast Mastery beside the yellow left rule (1440, 390). Drop the pill, keep the rule. GuideSpecRow.astro (pill wrap at 390 is f-selected's).
- 22 "Your guide: Beast Mastery Hunter >" in the band (1440). Prefix is the qualifier. lib/guides/copy.ts signedInCallout (f-content, handed off, not failed on).
- Fixed since round 2: "34 Dwarf Beast Mastery".

## /guides/hunter: PASS WITH FIXES
- Round 2 Races sentence is gone; Human and Skyborne marked "new" on their rows. Tell 23 cleared.
- Spec plates now show point splits (31/20/0). Sources is a plain ruled list. Good.
- 21 "Pets copy 10% of your attack power and all of your crit. No buffs from other players reach them; your Heart of the Lion does." Second sentence explains; trim. content/guides/hunter/index.md.

## /guides/hunter/survival: PASS WITH FIXES (tell 23 lede is f-content's)
- 23 Lede "Survival is a melee spec in Forever: two weapons, Raptor Strike queued..." still present (1440 under tab row, 390 first paragraph). content/guides/hunter/survival.md:48 (f-content). Fix: "Two weapons, Raptor Strike on cooldown, Mongoose Bite when it lights up, no pet in a raid."
- 23 "Survival levels in melee." opens Levels 20 to 60. Cut; start at "Keep a pet out while you level". survival.md.
- 23/21 Races prose restates the rail: "Humans can be Hunters for the first time in Forever" (survival.md:117) and "Skyborne can be Hunters on either faction:" (survival.md:121). Cut both clauses.
- 22 Racials column pastes tooltip text (Elune's Light "Grants 10% critical strike chance for 15 seconds on a 3 minute cooldown", Blood Fury "Increases Attack Power, Ranged Attack Power and Spell Power by 10%...", Command "Increases the damage dealt by Hunter and Warlock pets by 5%"). Trim to "10% crit, 15 s / 3 min". race-table.ts (f-content).
- 17 "Recommended" pills on Night Elf, Orc, Troll beside coloured names; keep one marker. GuideRacesLane.
- Handed off, not failed: "DPS · Updated" band, Races dead plate, blank race icons, broken spell icons.

## Verdicts
- /guides: PASS
- /guides-you: PASS WITH FIXES
- /guides/hunter: PASS WITH FIXES
- /guides/hunter/survival: PASS WITH FIXES

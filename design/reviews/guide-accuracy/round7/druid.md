# Guide accuracy round 7: Druid

Auditor: druid persona. Tree: gr7/audit at 394562f20 (land/guides-r6, all six rounds applied; READ-ONLY), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`). Before starting I read the round 1 to 6 druid reports with their fix logs, round4/druid-levels.md, the round 2 to 5 race-sims files and ledger.md. Since round 6's audit head (367cbadb7), the only druid changes are the round 6 fixes in 394562f20: the four Stat priority paragraphs, the two Wolfshead lines and the `updated:` stamps. No commit touched addon-builds.json, talents/druid.json, trainables/druid.json, spelltips/druid.json, the druid BiS files (last: de044125c) or the druid APLs.

## Decoded builds (unchanged since round 1, all legal, 51 points)

Same as round4/druid.md. Balance `5232220115101151/0/555301`, raid `5232220115501351/0/505003`. Feral `01/50212032123012123001/535033`, raid `05002/50230032023011125001/05503`. Bear `0/5523033202310212551/051`. Restoration `05312001/0/55530051531132`.

I re-decoded the shipped leveling orders (addon-builds.json, index 0 = level 10). They are identical to round5/druid.md. Every Levels 1 to 19 paragraph and every 20 to 60 talent claim matches them: Balance Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Moonkin Form 40. Feral 31 Feral points done with Shifting Power at 40, then Genesis 41, Nature's Focus 42-46, Heart of the Wild 47-49, Naturalist, Natural Shapeshifter and Subtlety 58-60. Bear Ferocity 10-14, Furor 15-19, Lacerate 25, Primal Bite 40. Restoration Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R7-1 | CONFIRMED | talent | feral.md:59 | "**Rend and Tear**: 3% more melee damage a point against a bleeding target, 15% at 5." | talents/druid.json Rend and Tear rank 5 (70755): "Increases damage done by your **melee abilities** on Bleeding targets by 15%." Blizzard's deep dive says the same ("increases damage your melee abilities deal to Bleeding targets"). The engine applies it to abilities only, never to auto attacks or bleed ticks (sim/druid/feral_talents.go:79-85 at d1d9f2775). "Melee damage" also covers your white swings, which are a real share of a cat's damage and get nothing. The bear page already says it right (feral-bear.md:60, "15% more from your melee abilities"). This wording dates from round 1's DRUID-1 fix, which fixed the number but kept the scope. | "**Rend and Tear**: 3% more damage a point from your melee abilities against a bleeding target, 15% at 5. Your auto attacks do not get it. Rip keeps the boss bleeding, so it is on for every Shred and Bite after your first finisher." |
| DRUID-R7-2 | CONFIRMED | spell | feral.md:81 | "Shifting Power costs 55% of your base mana, and mana decides how many you get." | Shifting Power 1322605 costs 55% of base mana (spelltips/druid.json `cost_pct: 55.0`), and its client text says "Shifting Power's cost is reduced by effects that reduce the cost of Shapeshifting". Natural Shapeshifter 3/3 is "Reduces the mana cost of all shapeshifting by 30%" (talents/druid.json), and both cat builds take 3 points (decoded above). The page says so itself at :67 ("Natural Shapeshifter ... makes Shifting Power cheaper"). So with the page's own build each press costs 38.5% of base mana, not 55%, and the Mana paragraph contradicts the talent paragraph. | "**Mana.** Shifting Power costs 55% of your base mana, 30% less with Natural Shapeshifter's three points, and mana decides how many you get." |

Checked and right (no finding):

**Round 6 edits re-checked**
- Balance stat prose: no ranking words. Spell power is every cast. A resisted Starfire wastes 3 seconds of casting (3.5 s base, minus Improved Starfire's 0.5 s). A resisted Wrath banks no Eclipse (round 1 Q3). The spell hit cap rule is right. Starfire and Moonfire are Arcane, and Wrath and Insect Swarm are Nature (spelltips). Crit is doubled by Vengeance ("by 100%") and triggers Nature's Grace ("All non-periodic spell criticals").
- Feral stat prose: Cat Form is "Attack Power by 12 plus Agility", so Agility is attack power as well as crit. Heart of the Wild is 10% Strength in Cat Form at 5 (2% a point, so 6% at 3). Blood Frenzy turns builder crits into combo points. The paw swing is 1.0 s, and Shred is 200% weapon damage plus 80 (deep dive and spelltips 9830).
- Bear stat prose: Bear Form's attack power is flat ("by 30"/"by 120"), so Agility is Dodge, armor and crit, as the page says. The form multiplies the armor items carry (180% and 360%; engine bear_form.go applies it to bonus armor too). The paw swings every 2.5 s. Thick Hide is armor per Defense point over five times your level.
- Restoration stat prose: Bloom is 100% of Spirit. Reflection is 50% at 3. Innervate is 400%. Crit only lands on direct heals.
- Wolfshead Helm (8345): ilvl 45, required level 40, blue, druid-only effect. The stable reasons only.
- No ranking of stats or gear by a nightly value anywhere on the five pages. The one cross-class comparison, "crit is worth more to you than to most casters" (balance.md:50), rests on Vengeance, a stable mechanic (round 6 left it on purpose).

**Spells and talents, every named one re-read (70755 text)**
- Moonkin Form and Leader of the Pack: "all party or raid members within 45 yards", so "everyone within 45 yards" and "everyone in range" hold. Nature's Reach is 4% hit with no spell-only wording. Eclipse is 0.5 s on the next 2 Starfires, with up to 4 charges stored. Improved Wrath is 0.5 s and 50%. Moonglow is 25% on damaging spells. Moonfury is 10% Arcane and Nature. Hurricane has no cooldown.
- Ferocity covers Maul, Primal Bite, Swipe, Claw and Rake. Feral Instinct is 30% Swipe at 3. Savage Fury covers Claw, Rake, Shred, Maul and Swipe. Predatory Discipline: "no longer converts extra Energy into damage, but still deals extra damage based on your excess Energy". Shredding Attacks is 18 energy off Shred and 3 rage off Lacerate. Lacerate is "plus 10% weapon damage per existing application", stacking 5. Primal Bite is "100% normal damage plus 77" on a 6 s cooldown for 20 rage (15 with Ferocity). Natural Reaction is 5% dodge plus 5 rage a dodge, and Feral Swiftness is 4% dodge, so 9%. Furor's text matches the index and both feral pages.
- Enrage gives 10 up front plus 20 over 10 s, and costs 27% (Bear) or 16% (Dire Bear) base armor. Barkskin cuts physical damage taken by 20% for 15 s on a 60 s cooldown. Frenzied Regeneration converts 1% health a rage. Challenging Roar lasts 6 s. Swipe hits 3. Demoralizing Roar rank 5 is 193 at 52, 204 at 60.
- Poised is learned at 32 and works at full energy out of combat. Ravage and Pounce both need "prowling and behind". Rip lasts 12 s. Ravage, Pounce and Cower train at 32, 36 and 28. Rake and Ferocious Bite at 24 are not contradicted by the 20 to 60 text.
- Swiftmend is 20% base mana on a 15 s cooldown. Wild Growth is a 3.0 s cast on a 6 s cooldown and heals the target's party. Tranquility reaches party members within 20 yards. Gift of the Wild ranks come at 50 and 60. Rejuvenation ranks 8, 9 and 10 are at 46, 52 and 58. Innervate trains at 40.
- Trainer levels in all four Levels 1 to 19 blocks still match trainables/druid.json.

**Items and consumables**: Thistle Tea 7676, Mighty Rage Potion 13442, Greater Stoneshield Potion 13455, Major Mana Potion 13444, Demonic Rune 12662, Elemental Sharpening Stone 18262, Dense Weightstone 12643, Brilliant Mana Oil 20748, EZ-Thro Thorium Grenade 260816, Aquarius Belt 16608 and Idol of Shifting Tides 263411 all exist in 70755 items.json. The deep dive (fetched 2026-11-04) confirms: "Rage potions and Thistletea are also now usable by Druids"; weightstones, Crusader and Windfury work in form; most bombs still do not; racials, Barkskin and Faerie Fire work in form; swings are 1.0 and 2.5 s.

**Card notes**: I re-read every rendered druid APL note. None states anything false. The Lacerate note ("each stack makes the next cast hit harder") matches the client text.

**Sources**: every frontmatter link returns 200 (Blizzard deep dive, Blizzard panel recap, Wowhead Skyborne First Look, Talents Forever data.json, /sim/weights, /sim/drops, /bis/druid/feral-bear, /bis/druid/restoration). Each one still follows the ledger's source-link rule.

No em or en dashes. No site, sim, APL, preset or band words in prose. No quoted nightly numbers.

## Decisions

1. **Rend and Tear's scope follows the client text (R7-1):** it covers melee abilities, not auto attacks. Three sources agree: the client description, Blizzard's deep dive and the engine's implementation. Nothing points the other way.
2. **Shifting Power's cost is stated for the page's build (R7-2):** the client text links it to shapeshift cost reductions, and both cat builds take Natural Shapeshifter 3.
3. **Bloom "the best single point you take" (restoration.md:55) stays:** it is a talent judgment, not a nightly value, and no measured per-talent data on 70755 contradicts it. Not filed.

## Sims requested

No new ones. Two are carried and still open:
1. From round 4 (S10): druid-feral leveling, the switch from Claw to Shred. Run ladder levels 20 to 25 with the Shred line gated on Shredding Attacks 3/3 (level 22), against the current Ferocious Bite rank 5 gate.
2. From round 3: druid-restoration with Improved Moonfire 1 moved into Improved Regrowth 1 (`05302001/0/555300515311321`). Heal profile, band 60 raid, seeds 1 and 50001, with a ladder-first row that keeps Improved Regrowth.

## Sim findings

- S1 (carried, still open at d1d9f2775): `rendAndTearDamagePerRank = 0.02` (sim/druid/feral_talents.go:92) against the client's 3% a rank. The ability-only scope there is correct, and R7-1 brings the prose into line with it.
- S7, S8, S9, S11, S12 and SIM-D1 (carried): Nature's Reach melee hit for the feral raid re-run, Feral Charge (Cat) unmodeled, the Pounce behind-target check, expertise in the bear weights, bear Defense noise, and the feral leveling Nature's Focus and Subtlety points.
- No new engine finding.

## Sources

- Client 1.60.1.70755: talents/druid.json (every rank text, including Rend and Tear, Natural Shapeshifter and Predatory Discipline), spelltips/druid.json (Shifting Power 1322605 `cost_pct` 55 and text, Cat Form 768, Bear and Dire Bear Form, Lacerate, Primal Bite, Enrage, Barkskin, Frenzied Regeneration, Poised, Ravage, Pounce, Rip, Swiftmend, Wild Growth, Tranquility, Moonfire, Starfire), trainables/druid.json, items.json, addon-builds.json, bis/druid-*.json; APLs data/curated/apl/druid-*.json.
- Engine d1d9f2775: sim/druid/feral_talents.go:79-117 (Rend and Tear on abilities only), sim/druid/bear_form.go (form armor).
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (fetched 2026-11-04: Rend and Tear "your melee abilities", consumables, Thistle Tea, racials in form, swing speeds).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (200).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops, https://foreversixty.gg/bis/druid/feral-bear, https://foreversixty.gg/bis/druid/restoration (all 200).

## Fix log (round 7)

- DRUID-R7-1 (Decision 1): fixed, feral.md:59. Checked against talents/druid.json Rend and Tear ranks 1 to 5 (70755): "Increases damage done by your melee abilities on Bleeding targets by 3%" ... "15%". Line now reads 3% a point from your melee abilities, 15% at 5, auto attacks do not get it, on for every Shred and Bite after the first finisher. Bite and Shred are abilities; the line makes no claim about Rip ticks.
- DRUID-R7-2 (Decision 2): fixed, feral.md:81. Checked against spelltips/druid.json Shifting Power (`cost_pct` 55.0, "Shifting Power's cost is reduced by effects that reduce the cost of Shapeshifting") and talents/druid.json Natural Shapeshifter 3/3 ("Reduces the mana cost of all shapeshifting by 30%"). Both cat builds take 3 points, and feral.md:67 already says so, so the two paragraphs now agree.
- Decision 3 (Bloom judgment, restoration.md:55): left, as decided.
- Stat priority rule (all four spec pages): balance.md, feral.md, feral-bear.md and restoration.md Stat priority prose already gives lasting reasons with no ranking words and points to the table; no change needed. The prose test passes for druid.
- Other stat or gear ranking: feral-bear.md:116 leveling line "Stamina and Agility first" made neutral ("quest rewards with Stamina and Agility over vendor gear"). Enchant and BiS sections on all pages rank nothing by a nightly value. balance.md:50/115 "crit is worth more to you" sit outside Stat priority and rest on Vengeance, a stable mechanic; kept.
- `updated:` already 2026-11-04 on feral.md and feral-bear.md; index.md, balance.md, restoration.md untouched.
- SIM findings: none new. S1 (engine `rendAndTearDamagePerRank = 0.02` vs client 3%) still open; the two carried sim requests stay open.

# Guide accuracy round 5: Druid

Auditor: druid persona. Tree: gr5/audit at 42f10dda9 (land/guides-r4: 752c02e43 plus the round 4 fix commits 77bd7c12b and a901bae80 and the ledger; READ-ONLY), client 1.60.1.70755, engine pin d1d9f2775 (sim/enginever/version.go).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`). Before starting I read the round 1 to 4 druid reports (with their fix logs), round4/druid-levels.md, round2/3/4 race-sims.md and ledger.md. There is no round5/druid-levels.md; the Levels 1 to 19 notes are in round4/druid-levels.md. I re-read every line of all five pages and every druid APL note. I re-dumped every druid talent rank text, every trainable rank and level and every named spell's 70755 description, and I re-decoded the four shipped leveling orders. Since round 4's audit head, the druid changes are the round 4 fixes only: source entries, the bear race paragraph, the Restoration Improved Moonfire sentence, paragraph breaks in the Levels 1 to 19 blocks and the feral Claw and Shred notes. I checked each of those edits again.

## Decoded builds (unchanged since round 1, all legal, 51 points)

Same as round4/druid.md. Balance `5232220115101151/0/555301` and raid `5232220115501351/0/505003`. Feral `01/50212032123012123001/535033` and raid `05002/50230032023011125001/05503`. Bear `0/5523033202310212551/051`. Restoration `05312001/0/55530051531132`.

Shipped leveling orders (addon-builds.json, index 0 = level 10), re-decoded:
- Balance: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Improved Moonfire 20-21, Nature's Majesty 22-23, Nature's Reach 24-25, Nature's Splendor 26, Insect Swarm 27, Vengeance 28-32, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonfury 36-39 and 41, Moonkin Form 40, Nature's Focus 42-46, Heart of the Wild 47-51, Naturalist 52-56, Reflection 57-59, Subtlety 60.
- Feral: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, Savage Fury 23-24, Feral Charge 25, Sharpened Claws 26-27, Predatory Strikes 28-30, Predatory Discipline 31, Blood Frenzy 32-33, Leader of the Pack 34, Natural Instinct 35-36, Rend and Tear 37-39, Shifting Power 40, Genesis 41, Nature's Focus 42-46, Heart of the Wild 47-49, Naturalist 50-54, Natural Shapeshifter 55-57, Subtlety 58-60.
- Bear: Ferocity 10-14, Furor 15-19, Savage Fury 20-21, Feral Swiftness 22-23, Feral Instinct 24, 26, 27, Lacerate 25, Thick Hide 28, 29, 31, Leader of the Pack 30, Shredding Attacks 32-34, Sharpened Claws 35-36, Predatory Strikes 37-39, Primal Bite 40, Blood Frenzy 41-42, Natural Instinct 43-44, Rend and Tear 45-49, Natural Reaction 50-54, Heart of the Wild 55-59, Naturalist 60.
- Restoration: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 and 21-23, Bloom 20, Naturalist 24-28, Gift of Nature 29 and 31-34, Nature's Swiftness 30, Tranquil Spirit 35-39, Improved Rejuvenation 40-42, Swiftmend 43, Living Spirit 44-46, Improved Tranquility 47-48, Genesis 49-53, Moonglow 54-56, Improved Moonfire 57, Nature's Majesty 58-59, Nature's Splendor 60.

Every Levels 1 to 19 paragraph and every 20 to 60 talent claim matches these orders.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R5-1 | CONFIRMED | talent (round 4 edit) | restoration.md:62 | "The one Balance point left after that sits in Improved Moonfire, which does nothing for a healer; it is there because nothing else in reach does more." | Round 4 added this sentence (fix log, DRUID-R4-3). Its second half is false. A point that does nothing cannot beat every other legal point, and a legal point that heals exists. Improved Regrowth rank 1 is "Increases the critical effect chance of your Regrowth spell by 10%" (talents/druid.json 104913), its prereq Improved Rejuvenation 3 is in the build, and `05302001/0/555300515311321` is legal (round3/druid.md R3-3). The same page spends a point there in its Wild Growth paragraph (:64). Round 3 measured that exact swap at 1014.5 to 1019.2 effective HPS (round3/race-sims.md:124). It was not adopted because it fell under the 1% bar and broke the ladder-first row, not because it heals less. | "The one Balance point left after that sits in Improved Moonfire, which does nothing for a healer. A point of Improved Regrowth, 10% more Regrowth crit, heals a little more if you would rather spend it there." |
| DRUID-R5-2 | CONFIRMED | spell | restoration.md:29 | "Swiftmend as a free extra heal" | Swiftmend costs 20% of base mana (spelltips/druid.json 18562 `cost_pct: 20.0`; Swiftmend APL note). The page says so itself at :72: "all it costs is a global and 20% of your base mana". It does not use up the HoT, but it is not free. | "Restoration heals with time: Regrowth and Rejuvenation on the tank all fight, Swiftmend as an extra heal that leaves the HoT ticking, Bloom when the HoTs run out, and Healing Touch for the holes." |
| DRUID-R5-3 | LIKELY | leveling (round 4 edit) | feral-bear.md:111 | "Furor from 15 to 19 climbs to 10 Rage on every shift into Bear Form by 19" | Furor rank 1: "Gives you a 20% chance to gain 10 Rage when shapeshifting into Bear Form"; rank 5: "100% chance" (talents/druid.json 104939). The amount is always 10. What climbs is the chance. The round 4 wording reads as if the rage grows from a small amount to 10. | "Furor from 15 to 19 adds a 20% chance a point of 10 Rage when you shift into Bear Form, so by 19 every shift gives it, and a pull opens with Maul ready." |
| DRUID-R5-4 | LIKELY | talent | index.md:30 vs feral-bear.md:54 | index:30 "Shifting Power, the Feral capstone"; bear:54 "**Primal Bite** (1 point): the capstone." | Both are tier 6 Feral talents (talents/druid.json: Primal Bite t6c1, prereq Leader of the Pack; Shifting Power t6c2, prereq Predatory Discipline). The class page calls one of them the only Feral capstone, and the bear page gives that title to the other. feral.md:63 ("the capstone") is fine because it means the cat's. | index:30 "Shifting Power, the cat's capstone in Feral, gives a cat 40 energy for mana instead." |
| DRUID-R5-5 | LIKELY | source link | feral-bear.md:12-21 | The bear page lists no best-in-slot source. | Ledger round 4 source-link rule: "healer and tank pages link their own best in slot page". restoration.md:22-24 and warrior/protection.md:36-37 do this. https://foreversixty.gg/bis/druid/feral-bear resolves and is the Feral Bear page ("Druid Feral Bear leveling BiS", bands 20 to 60, tank weights; WebFetch 2026-11-03). Paladin Protection and the priest and shaman healers also lack one, so this is uneven rule coverage rather than a false entry. | Add `- label: 'Forever Sixty, Feral bear best in slot'`, `url: https://foreversixty.gg/bis/druid/feral-bear`, `kind: site`. |

Checked and right (no finding):

**Round 4 edits re-checked**
- The source entries open the pages their labels name. /sim/weights is "Stat weights" for damage specs (Balance and Feral only; Restoration correctly dropped it). /sim/drops is the Droptimizer ("which boss drops are upgrades"). /bis/druid/restoration is the Restoration BiS page (bands 20 to 60). No druid page links /sources. Talents Forever is labelled as beta client data.
- Bear race paragraph: Night Elf takes slightly less damage (465.8 against 473.5 a second) and holds a little more threat (964 against 956), and Tauren's TMI is lower (35.1 against 37.3), per round4/race-sims.md:30-31. Druids are faction-locked (races.json: Night Elf alliance, Tauren horde; combos.json class 11 = races 4, 6, 95, 96). Skyborne is named only as the paid option.
- Feral Claw and Shred notes: Shred 60 minus 6, 12 and 18 (Shredding Attacks at 20, 21 and 22) is 54, 48 and 42 against Claw's 40 with Ferocity, so "8 to 14 more" holds. Shred is 200% plus 22 to 80, and Claw is 140% plus 27 to 115, so "Shred hits harder than Claw at every rank" holds.

**Levels 1 to 19 (all four pages, client data)**
- Trainer levels: Wrath 1 (ranks at 6 and 14), Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18, Healing Touch 1, Rejuvenation 4, Regrowth and Revive 12, Cure Poison 14, Bear Form, Maul, Growl and Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul rank 2 18, Aquatic Form 16, Cat Form 20.
- Talent texts at the level the page means: Improved Wrath 5 is 0.5 s and 50%. Genesis 2 is 4% periodic. Moonglow 3 is 25% on damaging spells. Ferocity 5 is 5 off Maul, Primal Bite, Swipe, Claw and Rake. Feral Swiftness 2 is 30% Cat speed and 4% dodge. Feral Instinct 1 is 10% Swipe. Brutal Impact 2 is +1 s on Bash and Pounce and -30 s on Bash. Heart of the Wild 5 is 10% Intellect and 20% bear Stamina. Reflection 3 is 50%. Nature's Focus 2 is 28%. Tier gates are met at every step.
- Entangling Roots: "only 1 target Rooted at a time". Revive: "Cannot be cast when in combat". Wolfshead Helm (items/druid.json 8345): ilvl 45, required level 40, blue, "5 Rage from activating Enrage and an additional 5 Energy from activating Shifting Power".

**Rest of the pages**
- Moonkin Form and Leader of the Pack are 3% crit within 45 yards. Nature's Grace is 10% casting speed and global cooldown for 3 s. Eclipse lasts 15 s and stores 4. Nature's Splendor is Moonfire and Rejuvenation +3 s, Regrowth +6 s and Insect Swarm +2 s ("longer DoTs" holds). Nature's Reach is 20% range and 4% hit. Nature's Majesty is spells and melee (so it counts for the cat and the healer). Genesis covers "spells and abilities" (Rip included). Predatory Strikes is 150% of level in Cat and Bear. Savage Fury covers Claw, Rake, Shred, Maul and Swipe (both pages' partial lists are true). Blood Frenzy is 5 rage and a combo point at 2/2 and 50% at 1. Natural Reaction plus Feral Swiftness is 9%. Living Spirit is 15%, Tranquil Spirit 10%, Improved Rejuvenation 15% and Gift of Nature 10%. Improved Tranquility (60%) takes Tranquility from 5 min to 2.
- Demoralizing Roar rank 5 is 193 at its spell level 52 with -1.4 a level and a max level of 62 (spellconst 9898), which is 204 at 60, as the bear card says. Enrage costs 16% base armor in Dire Bear Form. Barkskin is physical only, 15 s, 1 min. Frenzied Regeneration is 1% health a rage. Challenging Roar has a 10 min cooldown. Swipe hits 3 targets. Primal Bite costs 20 rage (15 with Ferocity) on a 6 s cooldown. Lacerate costs 15 (12 with Shredding Attacks), stacks 5 and lasts 15 s.
- Rejuvenation lasts 12 s and Regrowth 21 s (15 and 27 with Nature's Splendor, as the APL says). Wild Growth heals over 7 s with a 3.0 s cast (2.5 with Naturalist) on a 6 s cooldown. Swiftmend has a 15 s cooldown. Innervate is 400% for 20 s on 6 min. Gift of the Wild is rank 1 at 50 and rank 2 at 60. Rejuvenation ranks 8, 9 and 10 are at 46, 52 and 58.
- Judgement of Wisdom procs from spells as well as attacks ("attacks and spells used against the judged enemy", spelltips/paladin.json Seal of Wisdom). That supports Balance's "Judgement of Wisdom carries the rest".
- Blizzard deep dive (fetched): Moonkin casts all non-healing spells; the Cat Feral Charge has a 30 s cooldown and does not interrupt or root, which matches feral.md:55, where only the bear charge interrupts; Hurricane has no cooldown; racials and Barkskin work in form; Faerie Fire is free in form on a 6 s cooldown; Revive is at 12. The panel recap names six new race and class pairs, none of them a druid, which matches index.md:36 and combos.json.
- Card notes in all four druid APLs: no note states anything false. I re-read every line that renders.
- No em or en dashes, no site, sim, APL, preset, band or search words in druid prose, and no quoted nightly numbers. statPriority is unchanged and still matches the BiS files (no druid BiS commit since round 3).

## Decisions

1. **Restoration's spare point (R5-1):** the prose admits the point is dead and names Improved Regrowth as the small gain. The build string stays, because the integrator's round 3 call (under the bar, and it breaks the ladder-first row) still stands.
2. **Swiftmend is not "free" (R5-2):** "does not consume the HoT" is the Forever change. The mana cost stays real (client cost_pct 20).
3. **"Capstone" (R5-4):** both tier 6 Feral talents are capstones. Each page names the one for its role.
4. **Natural Reaction wording (bear :39 against :51):** this was softened in round 2 and the two lines no longer contradict each other. Not re-filed.
5. **Gift of Nature and Tranquility:** the beta-log ruling from round 1 stays. Not re-filed.

## Sims requested

1. Carried from round 4 (S10): druid-feral leveling, Claw to Shred switch. Ladder levels 20 to 25 with the Shred line gated on Shredding Attacks 3/3 (level 22), against the current Ferocious Bite rank-5 gate.
2. Carried from round 3 (integrator): druid-restoration Improved Moonfire 1 into Improved Regrowth 1 (`05302001/0/555300515311321`), heal profile at band 60 raid, seeds 1 and 50001, together with a ladder-first row that keeps Improved Regrowth. If it clears the bar, the build moves and R5-1's sentence becomes the build line. R5-1 is correct either way.

## Sim findings

- S1 (carried): `rendAndTearDamagePerRank = 0.02` against the client's 3% a rank.
- S7, S8, S9, S10 and SIM-D1 (carried): Nature's Reach melee hit for the feral raid re-run, Feral Charge (Cat) unmodeled, the Pounce behind-target check, the Claw gate and the feral leveling Nature's Focus/Subtlety points.
- No new engine finding.

## Sources

- Client 1.60.1.70755: talents/druid.json (all ranks), trainables/druid.json, spelltips/druid.json (Swiftmend 18562, Shred/Claw all ranks, Demoralizing Roar 9898, Lacerate, Primal Bite, Rejuvenation, Regrowth, Wild Growth, Revive, Entangling Roots), spellconst/druid.json (9898 points per level, max level), spelltips/paladin.json (Seal of Wisdom), items.json and items/druid.json (8345, 16608, 263411, consumables), races.json, combos.json, addon-builds.json; APLs data/curated/apl/druid-*.json; round3/race-sims.md:124 and round4/race-sims.md:28-33.
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (fetched 2026-11-03).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (fetched; six new pairs, none a druid).
- Wowhead, "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (fetched; title matches).
- Forever Sixty: https://foreversixty.gg/sim/weights ("Stat weights", damage specs only), https://foreversixty.gg/sim/drops ("Droptimizer"), https://foreversixty.gg/bis/druid/restoration and https://foreversixty.gg/bis/druid/feral-bear (BiS pages, bands 20 to 60), all fetched 2026-11-03.

## Fix log (round 5)

- DRUID-R5-1 (CONFIRMED): fixed, restoration.md:62. Improved Moonfire is now called a dead point and Improved Regrowth (10% Regrowth crit) is named as the small gain. Build string unchanged (Decision 1).
- DRUID-R5-2 (CONFIRMED): fixed, restoration.md:29. "free extra heal" is now "an extra heal that leaves the HoT ticking". The 20% base mana cost stays stated at :72 (Decision 2). No other druid line or APL note calls Swiftmend free.
- DRUID-R5-3 (LIKELY, verified against Furor rank texts 20% to 100% chance of 10 Rage): fixed, feral-bear.md:114. The line now says each point adds a 20% chance of 10 Rage, so by 19 every shift gives it. I did not use the suggested text because it had a typo.
- DRUID-R5-4 (LIKELY, verified: Primal Bite and Shifting Power are both tier 6 Feral): fixed, index.md:30, "the cat's capstone in Feral" (Decision 3). bear:54 and feral:63 are unchanged.
- DRUID-R5-5 (LIKELY, verified: the round 4 rule plus the page resolves): fixed, feral-bear.md:21-23. Added "Forever Sixty, Feral bear best in slot", https://foreversixty.gg/bis/druid/feral-bear, kind: site.
- Decisions 4 and 5: no change, as decided.
- Build strings, point orders, card actions and APL notes: unchanged. No new SIM findings. Sims 1 and 2 in "Sims requested" are carried.

## Fix log (round 5, review pass)

- MUST 1 (Furor line): fixed, feral-bear.md:114, the reviewer's line as written. Checked against the Furor rank texts (talents/druid.json 104939): 20%, 40%, 60%, 80%, 100% chance of 10 Rage on a shift into Bear Form or Dire Bear Form.
- MUST 2 (stat prose against the table): fixed, feral-bear.md:90. The prose now follows the rendered band 60 order (Stamina, Agility, Dodge, Hit, Armor, Strength, Attack power, then Defense and crit "Not significant"), with a stable reason per stat, no numbers, and the hit-cap gear dependence. "Counts twice" is dropped. Thick Hide (talents/druid.json 104942) does add 0.67/1.33/2.00 base armor per Defense point over five times your level, so the line keeps "a little armor" and says Defense sits last. Also tightened feral-bear.md:35 to "every point of Defense over five times your level" to match the rank text.
- SHOULD (Elune's Light on Restoration): fixed, restoration.md:96, now "Regrowths and Healing Touches alike".
- SHOULD (Feral Swiftness before Cat Form): fixed, feral.md:113. trainables/druid.json lists Cat Form (768) at level 20, so the line gives the 4% dodge now and the run speed once you learn Cat Form at 20.
- SHOULD (Restoration Improved Moonfire): left as is, per the review; the carried sim decides it.
- Frontmatter statPriority in feral-bear.md is not what renders (the chips come from the data), so it is untouched. Build strings, card actions, APL notes and BiS files: unchanged.

# Guide accuracy round 11: Druid

Auditor: druid persona. Tree: gr11/audit at d54de48aa (main after round 10, READ-ONLY), client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go). I read the round 1 to 10 druid reports, round4/druid-levels.md, the race-sims files and ledger.md first.

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`), and the four druid APL card notes.

## What moved since round 10 (84460cec8 to d54de48aa)

- eabba8f84 (round 10 fixes): feral.md (16 lines), feral-bear.md (8), restoration.md (9), and the druid-feral and druid-restoration APL notes.
- ec60e6abf (the leveling strip, planner and home now use the ladder's order) touches components only. addon-builds.json, talents, trainables, spelltips and the druid bis/*.json are unchanged.
- No druid build string changed.

## Decoded builds (all legal, 51 points, prerequisites and tier gates met)

These are unchanged from round 10:

- Balance leveling `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- Feral leveling `050022/50212032123012123001/05501`, 9/31/11.
- Feral raid `050022/50230032023012125001/055`, 9/32/10. Its Feral tier 1 is Feral Swiftness 2 and Feral Instinct 3, exactly the 5 points that open tier 2, because Brutal Impact and Feral Charge are gone.
- Bear `0/5123033202310212551/055`, 0/41/10.
- Restoration `05302001/0/555300515311321`, 11/0/40. Prerequisites: Nature's Splendor needs Nature's Majesty 2, Nature's Swiftness needs Naturalist 5, and Improved Regrowth needs Improved Rejuvenation 3. All are met.

I re-derived the shipped point orders from addon-builds.json (index 0 = level 10). They are identical to the round 10 listing. Balance and Restoration have the same leveling and raid orders. Bear has one order for both.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| DRUID-R11-1 | CONFIRMED | talent (round 10 edit) | feral.md:67 | "Natural Shapeshifter only trims the mana cost of Shifting Power, so in a raid that point does more in Feral Instinct and Rend and Tear." | (a) talents/druid.json Natural Shapeshifter: "Reduces the mana cost of all shapeshifting by 30%." It cuts every form shift, and Shifting Power is only one of them (Shifting Power 1322605: "cost is reduced by effects that reduce the cost of Shapeshifting"). The same paragraph tells the cat to shift out to Innervate, rez and decurse, and each of those shifts back into Cat Form costs mana. (b) Feral Instinct: "Increases damage done by your Swipe ability by 30% and reduces the chance enemies have to detect you while Prowling". Swipe is a bear ability: the trainable Swipe ranks 779/780/769/9754/9908 are Bear Form, the 1264494 "Swipe" ids are a hunter pet ability (Wowhead Forever tooltip: "Requires Hunter"), and 411128 is the SoD rune. The page itself says "You have no cleave". Feral Instinct therefore adds no raid damage for a cat. The raid build takes it because Feral Swiftness 2 plus Feral Instinct 3 are the 5 tier-1 points that open tier 2 once Brutal Impact and Feral Charge are gone. The round 10 fix log calls this line "the stable reason", but both halves are wrong. | "Natural Shapeshifter only makes your shifts and Shifting Power cheaper, so the raid build spends that point on Rend and Tear. The extra Feral Instinct points do nothing for a cat; they are there because the raid build drops Brutal Impact and Feral Charge, and something has to fill that row to open the next one." |

The other 527 lines were checked and are right (no finding):

**feral.md**
- :55 Feral 31 is done at 40, and the order from 41 is Genesis, Nature's Majesty, Nature's Reach, Heart of the Wild, Naturalist, Natural Shapeshifter 60. The respec diff (Brutal Impact 2, Feral Charge 1, Natural Shapeshifter 1 out; Feral Instinct +2, Rend and Tear +2 in) matches both strings.
- :57 to :65 talent effects:
  - Shred 60 - 18 = 42.
  - Savage Fury 10%.
  - Rend and Tear 3%/15% on melee abilities.
  - Blood Frenzy rank 2 is 100%.
  - Predatory Strikes 150% of level.
  - Predatory Discipline opens Shifting Power (prereq 104951).
  - Heart of the Wild 10% Strength.
  - Naturalist "all damage you deal by 10%".
- :67 Furor's text matches. The Balance points are Genesis 10% periodic, Nature's Majesty 4% "spells and melee attacks", and Nature's Reach "improves your chance to hit by 4%".
- :71 Poised 1328829 "While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10". Ravage and Pounce both need Prowl and the target's back.
- :75 Rake 40 - 5 (Ferocity) = 35.
- :77 Omen of Clarity 16864: "next damage or healing spell or offensive ability". :81 Shifting Power costs 55% of base mana.
- :89 Shred 200% and Ravage 435% of weapon damage.
- :101 races match races.json: Elune's Light 10% / 15 s / 3 min, Endurance 5%, War Stomp 5 targets 2 s, Wind Blessed 2% haste. Skyborne is named only as the paid option.
- :111 to :115 trainer levels: Moonfire 4, Bear Form/Maul/Demoralizing Roar 10, Bash 14, Swipe 16, Aquatic Form 16. Talents: Ferocity 10-14, Feral Swiftness 15-16 (4% dodge, 30% Cat speed), Feral Instinct 17 (10% Swipe), Brutal Impact 18-19 (+1 s stun, -30 s Bash).
- :119 Cat Form 20, Shredding Attacks 20-22, Ravage 32, Pounce 36, Feral Charge 25.
- Card notes: Claw 40 against Shred 54 at 20 and 48 at 21. Poised. Shifting Power. The round 10 Rip note is true whatever the Rip/Bite sim finds.

**feral-bear.md**
- :52 41 Feral and 10 Restoration.
- :54 to :64:
  - Thick Hide text.
  - 9% Dodge.
  - Furor rank 1 is 20% for 10 rage.
  - Lacerate stacks 5 at +10% weapon damage a stack.
  - Primal Bite 20 - 5 = 15 rage, on a 6 s cooldown.
  - Ferocity covers Maul, Swipe and Primal Bite.
  - Shredding Attacks is -3 on Lacerate.
  - Blood Frenzy is 5 rage a crit.
  - Heart of the Wild is 20% Stamina, and Naturalist is 10%.
- :70 Enrage 5229: 10 up front plus 20 over 10 s, and armor down for the duration.
- :81 Swipe hits 3 targets.
- :85 Barkskin 22812: physical damage only, 20%, 15 s, 1 min. Frenzied Regeneration 22842: 10 rage a second, 1% health a rage.
- :113 to :121 trainer levels: Maul, Growl and Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul rank 2 18, Frenzied Regeneration 36, Dire Bear Form 40. Order: Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19, Lacerate 25, Primal Bite 40.
- Card notes: Demoralizing Roar 204 at 60 (193 at 52, scaling, per rounds 2, 5 and 7), Maul 10 with Ferocity, and Swipe 3 targets.

**restoration.md**
- :51 40/11.
- :53 to :63 talent texts:
  - Gift of Nature 10%.
  - Improved Rejuvenation 15%.
  - Bloom is 100% of Spirit.
  - Naturalist is 0.5 s.
  - Reflection 50%.
  - Tranquil Spirit covers Healing Touch and Tranquility.
  - Living Spirit 15%.
  - Improved Tranquility: 100% less threat, 60% off 5 min = 2 min.
  - Heart of the Wild 10% Intellect.
  - Nature's Splendor: +3 s Rejuvenation, +6 s Regrowth.
  - Moonglow is the 3 points that open tier 2.
  - Improved Regrowth 10% a point.
- :65 the Wild Growth swap is legal: prereq Living Spirit 3, and 39 Restoration points before row 7. Wild Growth is a 3.0 s cast (2.5 s with Naturalist) on a 6 s cooldown, 1050 mana.
- :81 rank levels r1 4, r7 40, r8 to r10 46/52/58. With the 17-level grace: r7 85%, r1 0. The page matches the card (rank 7 at raid 60) and rank-choices.json (rank 1 below 60 and at bare 60, rank 4 on the tank from 28 to 49).
- :83 Innervate 29166: +400% regeneration and 100% while casting, 20 s.
- :107 to :113 trainer levels: Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14. Talents: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 (28%, "more than a quarter").

**balance.md**
- :67 the leveling-to-raid diff (Heart of the Wild 5 and Reflection 3 into Improved Starfire +4, Eclipse +2, Subtlety +2) checks out. Starfire is 3.5 - 0.5 = 3.0 s. Eclipse stores up to 4 charges at 0.5 s. Nature's Grace is "non-periodic spell criticals", 3 s.
- :69 to :77 talent texts: Vengeance +100% crit bonus, Moonfury 10%, Improved Wrath 0.5 s/50%, Nature's Reach 4%, Moonglow 25% on damaging spells, Moonkin 45 yd 3%, Nature's Focus 70%, Subtlety threat.
- :91 Omen of Clarity: "not consumed by Wrath".
- :95 Moonkin Form "Omen of Clarity's chance to trigger by 100%".
- :99 Hurricane has no cooldown in trainables.
- :125 to :127 Wrath ranks 1/6/14, Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18. Talents: Improved Wrath 10-14, Genesis 15-16 (4%), Moonglow 17-19.
- :135 Moonkin Form at 40 (order).

**index.md**
- Omen of Clarity baseline at 20 and Nature's Grasp at 10 (trainables).
- Poised 32 and Revive 12.
- The Furor text matches.
- No druid among the new race and class pairs.

**Sources**: every frontmatter link returned 200 to a browser user agent on 2026-11-04: both Blizzard pages, Wowhead Skyborne First Look, Talents Forever data.json, /sim/weights, /sim/drops, /bis/druid/feral-bear and /bis/druid/restoration. Each follows the round 4 source-link rule (no /sources, weights and drops on damage specs only, own best in slot page on tank and healer).

No em or en dashes. No site, sim, APL, preset or band words in prose. No quoted nightly numbers. Stat paragraphs explain and do not rank.

## Shoulds (wording, not findings)

- feral.md:119 "Cat Form arrives at 20 ... Claw builds your points and Rake adds a cheap bleed ... Ferocious Bite on the rest": Rake and Ferocious Bite both train at 24 (trainables). Suggest "Rake from 24" and "Ferocious Bite, from 24, on the rest".
- balance.md:91 "it always lands on a Starfire or a DoT": carried, "almost always" is exact.
- feral.md:73 "so do not pool energy for it": carried, see Sims requested 2.

## Decisions

1. **A round 10 fix sentence that states a talent's effect is held to the tooltip** (R11-1). Natural Shapeshifter is "all shapeshifting". Feral Instinct is Swipe and Prowl detection only, and neither does anything for a cat in a raid. A tier-gate filler is described as one, not as a damage point.
2. **Cat Swipe does not exist in Forever.** The trainer teaches only the bear ranks. 1264494 is a hunter pet Swipe (Wowhead Forever tooltip), and 411128 is a SoD rune (AcquireMethod 3). "You have no cleave" stands.
3. **Rounds 1 to 10 decisions stand**:
   - Swiftmend keeps the HoT.
   - Gift of Nature misses Tranquility.
   - Moonkin casts Rebirth, Innervate and Remove Curse.
   - Pounce and Ravage need you behind the target.
   - Bear Furor is 1.
   - Restoration ranks follow the cards.

## Sims requested

1. Carried, R10-12: druid-feral Rip against Bite at 5 points, at pin 9f7f1f6e4, band 60 raid and bare, seeds 1 and 50001. The prose no longer depends on the result.
2. Carried: druid-feral Bite energy pooling under Predatory Discipline. Compare the shipped list with holding the 5-point Bite until 80 energy or more.
3. Carried: bear race re-check on `0/5123033202310212551/055`, Night Elf against Tauren (threat and chance of death, raid 60, seeds 1 and 50001). feral-bear.md:103 "slightly less damage, a little more threat" is PROVISIONAL.
4. PROVISIONAL BiS and stat order: every druid bis/*.json in the tree is still engine d1d9f2775 (generated 2026-11-04 02:22 to 02:56 UTC), before the final re-rank. After the re-rank, check the statPriority frontmatter on all four specs and that the stat prose still reads with the table.
5. Carried S10: feral leveling, Claw to Shred gate (Shredding Attacks 3/3 at 22 against the Ferocious Bite rank 5 gate), ladder 20 to 25.

## Sim findings

- None new. Carried: S11 bear expertise weight, S12 bear Defense noise.

## Sources

- Client 1.60.1.70755:
  - talents/druid.json: Natural Shapeshifter, Feral Instinct, prerequisites, every talent named on the pages.
  - trainables/druid.json: levels and ranks.
  - spelltips/druid.json: 5229, 22812, 22842, 29166, 16864, 1328829, 1322605, 9908, 9898, 8910 and the cat abilities.
  - spells.json: Swipe ids.
  - races.json.
  - addon-builds.json (orders decoded at d54de48aa).
  - bis/druid-*.json (engine_version d1d9f2775).
- data/curated/apl/druid-*.json card notes. data/curated/rank-choices.json.
- Wowhead Forever tooltip, Swipe 1264494: https://nether.wowhead.com/forever/tooltip/spell/1264494 ("Requires Hunter", "Swipe 3 nearby enemies, inflicting 4 damage").
- Blizzard, "World of Warcraft: Forever Class Deep Dives: Hunter and Druid", https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid (200).
- Blizzard, "World of Warcraft: Forever Deep Dive Panel Recap", https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap (200).
- Wowhead, "Skyborne First Look", https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 (200).
- Talents Forever, https://talentsforever.com/data.json (200).
- Forever Sixty: /sim/weights, /sim/drops, /bis/druid/feral-bear, /bis/druid/restoration (all 200).

## Fix log (round 11)

- DRUID-R11-1: fixed, feral.md:67. Checked against talents/druid.json (Natural Shapeshifter "all shapeshifting by 30%"; Feral Instinct is Swipe and Prowl only). Wording "add nothing to a cat's damage" in place of "do nothing for a cat", since the Prowl half is not damage.
- SHOULD feral.md:119: applied, Rake and Ferocious Bite "from 24" (trainables/druid.json: Rake 24, Ferocious Bite 24).
- SHOULD balance.md:91 ("almost always"): left, carried and not checkable in one line.
- feral.md updated: already 2026-11-04.

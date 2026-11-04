# Guide accuracy round 13: Druid

Auditor: druid persona. Tree: gr13/audit at 604b78a80 (main after round 12, READ-ONLY). Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go; source read from gr10/engine). Before starting I read the round 1 to 12 druid reports and fix logs, round4/druid-levels.md, the race-sims files and ledger.md.

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`), and the card notes in all four druid APLs.

## What moved since round 12 (a8b147a6f to 604b78a80)

Only the round 12 fixes:
- feral.md:67, the Feral Instinct sentence (R12-1)
- druid-restoration.json priorityList[7] note and the top-level notes (R12-2)

addon-builds.json, talents, trainables, spelltips, rank-choices.json and bis/druid-*.json did not change. No druid build string changed.

## Decoded builds (all legal, 51 points, tier gates and prerequisites met)

- Balance leveling `5232220115101151/0/555301`, raid `5232220115501351/0/505003`
- Feral leveling `050022/50212032123012123001/05501`, raid `050022/50230032023012125001/055`
- Bear `0/5123033202310212551/055`
- Restoration `05302001/0/555300515311321`

Prerequisites checked from talents/druid.json:
- Lacerate needs Savage Fury 2. The bear takes Savage Fury at 20 and 21 and Lacerate at 25.
- Blood Frenzy needs Sharpened Claws 2.
- Rend and Tear needs Predatory Strikes 3.
- Primal Bite needs Leader of the Pack.
- Shifting Power needs Predatory Discipline.
- Nature's Splendor needs Nature's Majesty 2.
- Vengeance needs Improved Moonfire 2.
- Nature's Swiftness needs Naturalist 5.
- Improved Regrowth needs Improved Rejuvenation 3.
- Wild Growth needs Living Spirit 3.

Every build meets each prerequisite that applies to it.

## Findings

None. I found no CONFIRMED, LIKELY or NEEDS-SIM finding this round.

## Checks

**Round 12 fixes re-checked**
- feral.md:67 R12-1 holds on the tier math:
  - Feral Instinct and Brutal Impact are tier 1, and Feral Charge is tier 2.
  - With Brutal Impact gone, tier 1 is Feral Swiftness 2 and Feral Instinct 3, which is 10.
- The restoration [7] note R12-2 matches rank-choices.json: authored 9841, one band, 28 to 49, rank 4. It also matches restoration.md:81. Rank 11 (25299) needs the Book of Rejuvenation XI (items.json 21296, an Ahn'Qiraj book; the APL notes say core.IncludeAQ is false). That makes rank 10 the top rank you can train. So "keep the top rank on the tank", "ranks 8, 9 and 10 keep all of it" and the card's "rank 10 in a raid" agree.

**Leveling against addon-builds.json** (order decoded, index 0 = level 10)
- Balance: Improved Wrath 10 to 14, Genesis 15 and 16, Moonglow 17 to 19. Moonkin Form at 40 in both builds. Matches balance.md:127 and :135.
- Feral leveling:
  - Ferocity 10 to 14, Feral Swiftness 15 and 16, Feral Instinct 17, Brutal Impact 18 and 19. Shredding Attacks 20 to 22, Feral Charge 25.
  - Feral points are done at 40 (Shifting Power).
  - From 41: Genesis, then Nature's Majesty and Nature's Reach, then Heart of the Wild and Naturalist. Natural Shapeshifter at 60.
  - Matches feral.md:55, :113 and :119.
- Bear:
  - Ferocity 10 to 14, Furor 15, Feral Swiftness 16 and 17, Feral Instinct 18 and 19.
  - Lacerate 25, Primal Bite 40.
  - Matches feral-bear.md:115 and :121.
- Restoration: Heart of the Wild 10 to 14, Reflection 15 to 17, Nature's Focus 18 and 19. Two ranks of Nature's Focus are 28%, so "more than a quarter" holds. Matches restoration.md:109.
- Trainer levels (trainables/druid.json):
  - Caster spells: Wrath 1, 6 and 14. Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18.
  - Healing spells: Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14.
  - Bear Form spells: Bear Form, Maul, Growl and Demoralizing Roar 10. Enrage 12, Bash 14, Swipe 16, Maul rank 2 18. Frenzied Regeneration 36, Dire Bear Form 40.
  - Cat Form spells: Cat Form, Shred and Claw 20. Rake and Ferocious Bite 24. Cower 28. Ravage and Poised 32. Pounce 36. Omen of Clarity 20.
  - Innervate 40.
  - Every level the pages state is right.
- Quest items: Idol of Shifting Tides 263411 is a reward of Aquatic Form (quests 31 and 5061, level 16). The Aquarius Belt 16608 exists. Both were decided in round 4.

**Client values re-derived on 70755** (spelltips/druid.json)
- Wrath and Starfire casts:
  - Wrath rank 8 is a 2.0 s cast, 1.5 s with Improved Wrath. Starfire is 3.5 s, 3.0 s with Improved Starfire 5.
- Cat costs:
  - Shred is 60 energy at ranks 1 and 5, so 54, 48 and 42 with Shredding Attacks.
  - Claw is 45 (40 with Ferocity).
  - Rake is 40 (35 with Ferocity).
  - Ravage and Pounce both read "Must be prowling and behind the target".
- Rip 6 is a fixed 12 s at every combo point count.
- Bear costs:
  - Maul is 15 rage (10 with Ferocity).
  - Swipe is 20 rage and hits 3 targets.
  - Primal Bite 3 is 20 rage (15 with Ferocity) on a 6 s cooldown.
- Bear forms and cooldowns:
  - Dire Bear Form gives 360% armor from items and +600 health ("a flat block of health").
  - Enrage gives 10 rage plus 20 over 10 s.
  - Frenzied Regeneration converts 10 rage a second, 1% health per rage.
  - Barkskin is 20% for 15 s, every 60 s.
- Balance spells:
  - Hurricane has no cooldown.
  - Moonkin Form gives a 45 yd 3% crit aura, doubles Omen of Clarity's chance and blocks healing spells.
  - Omen of Clarity is "not consumed by Wrath".
- Healing spells:
  - Tranquility 4 heals the party every 2 s for 10 s, on a 5 min cooldown.
  - Wild Growth 3 is a 3.0 s cast on a 6 s cooldown, costs 1050 mana and heals over 7 s.
  - Swiftmend costs 20% of base mana on a 15 s cooldown.
  - Innervate is +400% for 20 s.
- Revive: "Cannot be cast when in combat".
- Wolfshead Helm 8345: level 40 blue, 10 Spirit, "5 Rage from activating Enrage and 5 Energy from activating Shifting Power".

**Talent texts** (talents/druid.json, max rank and the ranks the pages quote)
- Feral Instinct is 10% Swipe a rank.
- Brutal Impact 2: 1 s on Bash and Pounce, 30 s off Bash.
- Furor 1: 20% for 10 rage. Rank 5 restores 100% of your stored energy plus 10 a second outside bear form.
- Naturalist: 2% damage a rank and 0.5 s off Healing Touch and Wild Growth.
- Nature's Focus: 70%.
- Genesis: 2% a rank.
- Vengeance: crit damage bonus +100%.
- Eclipse: 2 charges, 0.5 s each, stores 4.
- Thick Hide, Rend and Tear 15%, Natural Reaction 5% and 5 rage, Heart of the Wild, Bloom, Improved Tranquility 60% and Improved Regrowth 50%/5 all read as the pages say.

**Card notes** (all four APLs)
- Every note matches its action, its condition and the client values above:
  - Feral: Shred 54/48 against Claw 40, and Shifting Power "at 60 or less".
  - Bear: Maul 10 with Ferocity.
  - Balance: Eclipse two charges.
  - Restoration: rank notes against rank-choices.json, including party rank 1 from 10 to 59 and at bare 60.

**Races** (races.json, combos.json)
- Elune's Light is 10% crit for 15 s on a 3 min cooldown. Quickness is 1% dodge. Endurance is 5% health. War Stomp stuns up to 5 targets for 2 s. Wind Blessed is 2% haste. Read Leyline is Alliance Skyborne only.
- combos.json has six new non-Skyborne race and class pairs, and none of them is a druid. Skyborne druids are covered by the sentence before ("can be druids on either faction"). So index.md:36 holds.
- Skyborne is named only as the paid option on every page.

**Header card against prose**
- No prose contradicts the raid-build header on Feral or Balance.
- Every bolded talent with a point count is the same in both builds, or the prose says which build it means.

**Sources**
- Frontmatter entries are unchanged since round 12, which fetched them on 2026-11-04. All follow the round 4 source-link rule.

**Voice**
- No em or en dashes. No site, sim, APL or band words in prose. No quoted nightly numbers.
- Stat paragraphs explain the stats and do not rank them.

## Shoulds (wording, not findings)

- balance.md:105 "a resisted Wrath banks no Eclipse charges": the engine banks charges on cast complete (sim/druid/talents.go applyEclipse, OnCastComplete). The client proc mask is not in the tree, and the Wowhead Forever tooltip does not say whether a resist banks charges. Keep the line until a log settles it, or soften it to "a resisted Wrath is a global for nothing".
- Carried from round 12: feral.md:77 should add "(55 with Wolfshead Helm)"; feral.md:85 "Leader of the Pack crits" should read "your crits"; balance.md:91 "always" should read "almost always".

## Decisions

1. Round 12's two fixes are right and stay.
2. Rank 11 Rejuvenation and Healing Touch are Ahn'Qiraj book ranks. The Restoration rank prose and cards correctly treat rank 10 as the top rank.
3. Rounds 1 to 12 decisions stand. These include:
   - Swiftmend keeps the HoT.
   - Gift of Nature misses Tranquility.
   - Pounce and Ravage need you behind the target.
   - Cat has no Swipe.
   - Bear Furor is 1 point.
   - Feral Charge (Cat) is the 30 s leap.
   - Nature's Reach covers melee hit.
   - Expertise is on a few late sets (bb4d616e0).

## Sims requested

These are all carried. The sim lane already holds the Feral Instinct points against Furor 2 or Thick Hide 2, the race re-runs and the BiS re-rank, so they are not re-filed here.
1. Carried, R10-12: druid-feral Rip against Bite at 5 points, band 60 raid and bare, seeds 1 and 50001. The prose matches the shipped card either way.
2. Carried: druid-feral Bite energy pooling under Predatory Discipline.
3. Carried S10: the feral leveling Claw to Shred gate, ladder 20 to 25.
4. PROVISIONAL: every bis/druid-*.json is still engine d1d9f2775, and its talent strings are not the shipped builds. After the re-rank, check statPriority on all four specs and that the stat prose still reads with the table.

## Sim findings

- New, minor: Eclipse banks charges on a resisted Wrath in the engine. The guide says it does not. Settle this from the client proc flags (SpellAuraOptions 408248) or a combat log.
- Carried: S11 bear expertise weight, S12 bear Defense noise, S13 cat abilities and Starfire on 70753 values.

## Sources

- Client 1.60.1.70755:
  - talents/druid.json
  - trainables/druid.json
  - spelltips/druid.json (every spell id named above)
  - items.json (21296, 21294, 8345, 16608, 263411, 260816)
  - items/druid.json
  - raw/items/item-sources.json (263411)
  - races.json
  - combos.json
  - addon-builds.json
  - bis/druid-*.json
- data/curated/apl/druid-*.json; data/curated/rank-choices.json.
- Engine 9f7f1f6e4: sim/druid/talents.go applyEclipse.
- Wowhead Forever tooltip, Eclipse, https://nether.wowhead.com/forever/tooltip/spell/408248 (fetched 2026-11-04; the text does not cover resists).

# Guide accuracy round 16: Druid

Auditor: druid persona. Tree: gr16/audit at 2c3135a2e (land/guides-r15: fifteen guide rounds on launch-state main; READ-ONLY). Client 1.60.1.70755, engine pin b2c4a2d05 (sim/enginever/version.go). Before starting I read the round 15 druid report and fix log, the earlier druid reports, and ledger.md through round 15.

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`) and Restoration (`restoration.md`). I read every line again. The best in slot rows were checked at band 60 (bare and raid, both factions) and, for the bear, at every band.

## What moved since round 15 (4600b7d82 to 2c3135a2e)

- feral-bear.md:107, Professions. This is the R15-1 fix plus the rendered-review MUST that scoped the line to the raid set. It is the only druid line that changed.
- The tree diff also touches addon-data.json, but only rogue Backstab conditions changed. The rogue and shaman APLs changed too.
- These did not change: druid APLs, rank-choices.json druid rows, addon-builds.json, talents/druid.json, trainables, spelltips, races.json and bis/druid-*.json.

## Decoded builds (unchanged; all legal, 51 points)

Decoded through talents/druid.json `tier`/`column`. Prerequisites were checked against `prereq_talent_id`.

- **Balance leveling** `5232220115101151/0/555301`
  - Balance: Improved Wrath 5, Genesis 2, Moonglow 3, Improved Moonfire 2, Nature's Majesty 2, Nature's Reach 2, Nature's Splendor 1, Insect Swarm 1, Vengeance 5, Improved Starfire 1, Nature's Grace 1, Eclipse 1, Moonfury 5, Moonkin Form 1.
  - Restoration: Nature's Focus 5, Heart of the Wild 5, Naturalist 5, Reflection 3, Subtlety 1.
- **Balance raid** `5232220115501351/0/505003`
  - Same Balance tree as leveling, except Improved Starfire 5 and Eclipse 3.
  - Restoration: Nature's Focus 5, Naturalist 5, Subtlety 3.
- **Feral leveling** `050022/50212032123012123001/05501`
  - Balance: Genesis 5, Nature's Majesty 2, Nature's Reach 2.
  - Feral: Ferocity 5, Feral Swiftness 2, Feral Instinct 1, Brutal Impact 2, Shredding Attacks 3, Savage Fury 2, Feral Charge 1, Sharpened Claws 2, Predatory Strikes 3, Predatory Discipline 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 3, Shifting Power 1.
  - Restoration: Heart of the Wild 5, Naturalist 5, Natural Shapeshifter 1.
- **Feral raid** `050022/50230032023012125001/055`
  - The same as leveling, except: Feral Instinct 3, Rend and Tear 5, and no Brutal Impact, Feral Charge or Natural Shapeshifter.
- **Bear** `0/5123033202310212551/055`
  - Feral: Ferocity 5, Furor 1, Feral Swiftness 2, Feral Instinct 3, Thick Hide 3, Shredding Attacks 3, Savage Fury 2, Sharpened Claws 2, Predatory Strikes 3, Lacerate 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 5, Natural Reaction 5, Primal Bite 1. That is 41.
  - Restoration: Heart of the Wild 5, Naturalist 5.
- **Restoration** `05302001/0/555300515311321`
  - Balance (11): Genesis 5, Moonglow 3, Nature's Majesty 2, Nature's Splendor 1.
  - Restoration (40): Nature's Focus 5, Heart of the Wild 5, Naturalist 5, Reflection 3, Gift of Nature 5, Bloom 1, Tranquil Spirit 5, Improved Rejuvenation 3, Swiftmend 1, Nature's Swiftness 1, Living Spirit 3, Improved Tranquility 2, Improved Regrowth 1.

The Wild Growth alternative at restoration.md:65 is legal:
- Wild Growth needs Living Spirit 3, which the build has.
- Improved Regrowth 2 needs Improved Rejuvenation 3, which the build has.
- Without Improved Tranquility, tiers 0 to 5 still hold 39 points.

## Findings

No CONFIRMED, LIKELY or NEEDS-SIM findings.

## Checks

**Round 15 fix (feral-bear.md:107)**, against every band of bis/druid-feral-bear.json and the bonding in items/druid.json:

- "Nothing in your raid set binds on pickup". This holds. The crafts in both band 60 raid rows all bind when equipped:
  - Outlaw's Collar 279253
  - Fel Cape 279269 (Tailoring)
  - Warbear Harness 15064
  - Hivethrasher's Girdle 275614
  - The band 60 bare rows have the same crafts without the cape.
- "the Leatherworking and Tailoring pieces bind when equipped". This holds.
- "Much of the 20 to 50 leather in your table, the Defender's, Brawler's and Warden's pieces, binds when picked up". This holds. Every one of these named pieces binds on pickup:
  - Band 20: Defender's Leather Pants 252445
  - Band 30: Brawler's Leather Helm 252512, Warden's Leather Belt 252460, Defender's Leather Kilt 252457
  - Band 40: Warden's Leather Gloves 252527, Warden's Leather Shoes 252466
  - Band 50: Warden's Leather Gauntlets 252549, Warden's Leather Waistguard 252475, Warden's Leather Boots 252470
  - The other leather crafts in those bands bind when equipped: Deviate Scale Belt, Barbaric Bracers, Raptor Hide Harness and Warbear Harness. So "much of" is accurate, and "a leveling bear needs Leatherworking for it" is true.
- The line no longer points at the cat's line, and it does not contradict feral.md:105.

**Leveling against addon-builds.json and trainables** (re-derived this round; both files unchanged):

- The trainer levels quoted on the five pages match trainables/druid.json:
  - Wrath 1/6/14; Moonfire 4; Entangling Roots 8; Nature's Grasp 10; Faerie Fire 18.
  - Healing Touch 1; Rejuvenation 4; Regrowth 12; Revive 12; Cure Poison 14.
  - Bear Form, Maul, Growl and Demoralizing Roar 10; Enrage 12; Bash 14; Swipe 16; Maul rank 2 at 18.
  - Cat Form, Claw, Shred, Rip and Omen of Clarity 20; Rake and Ferocious Bite 24; Lacerate rank 1 at 25.
  - Poised and Ravage 32; Pounce and Frenzied Regeneration 36.
  - Dire Bear Form, Innervate, Hurricane and Primal Bite 40.
- The talent orders in the Levels 1 to 19 paragraphs match the orders round 15 decoded. The per-rank values are the talent texts divided by rank:
  - Feral Instinct is 10% a point.
  - Nature's Focus is 14% a point, so 28% is "more than a quarter".
  - Genesis is 2% a point.
- Balance ranks (balance.md:135) match the three druid-balance rows in rank-choices.json:
  - Moonfire rank 2 at 20 to 29, rank 3 at 30 to 35, rank 4 at 36 to 39.
  - Wrath rank 3 at 22 to 26, rank 4 at 30 to 39.
  - Insect Swarm rank 1 at 30 to 38.
- Restoration ranks (restoration.md:81 and :117) match the 9841 and 8910 rows. The level-gap arithmetic at 60 also holds:
  - Ranks 8 to 10 keep the full bonus.
  - Rank 7 keeps 85%.
  - Rank 1 keeps none.

**Client values re-read** (talents/druid.json, spelltips/druid.json):

- Shifting Power and Furor.
- Eclipse: 2 charges, stores 4. Improved Starfire: -0.5 s.
- Nature's Grace: non-periodic spell crits. Vengeance: +100% to the bonus, which "doubles" the 50% spell bonus.
- Moonkin Form: Omen of Clarity +100%, and 3% crit within 45 yards. Leader of the Pack: 3% within 45 yards.
- Omen of Clarity is not consumed by Wrath.
- Thick Hide, Natural Reaction, Blood Frenzy, Natural Instinct and Rend and Tear.
- Heart of the Wild, Naturalist (cast time and 10% damage), Reflection, Gift of Nature and Bloom.
- Improved Tranquility (60%) and Nature's Splendor (+3/+6/+2).
- Pounce and Ravage: "Must be prowling and behind the target".
- Poised: 10 off, at max energy, out of combat.
- Wolfshead Helm 8345: item level 45, needs level 40, 10 Spirit, the Enrage and Shifting Power text, binds when equipped.
- Idol of Shifting Tides 263411 and Aquarius Belt 16608 exist.
- Every quoted number on the pages matches these values.

**Best in slot and professions** (band 60, both presets, both factions):

- Feral (feral.md:105): Molten Heart is in all four rows, and Frozen Heart sits beside it in the Alliance raid row. The Horde raid row has Rune of the Guard Captain instead. Every Leatherworking piece binds when equipped.
- Restoration (restoration.md:101): the only bind-on-pickup crafts are Dormant Heart (Enchanting) and Mender's Leather Gauntlets (Leatherworking). The other crafts, including Idol of Swiftness, bind when equipped.
- Balance (balance.md:119): Frozen Heart (Enchanting, bind on pickup) is in all four rows. Every Tailoring and Leatherworking piece binds when equipped. The line names Enchanting, so nothing it says is false (see Shoulds).
- Cat and bear wear The Unstoppable Force everywhere, which fits the two-hander lines.

**Races**: these all read as the pages say in races.json:
- Night Elf: Elune's Light (10% crit, 15 s, 3 min) and Quickness (1% dodge).
- Tauren: Endurance (5%) and War Stomp (5 enemies, 2 s).
- Skyborne: Wind Blessed (2% haste) and Read Leyline (50% for 15 s on 2 min).

Skyborne is named only as the paid option on every page.

**Sources**: every frontmatter URL was fetched on 2026-11-04 with a browser user agent.
- 200: both Blizzard articles, talentsforever data.json, /sim/weights, /sim/drops, /bis/druid/restoration and /bis/druid/feral-bear.
- 403: Wowhead Skyborne First Look. This is the bot block, so it counts as unknown, not dead. It returned 200 in round 15.
- Labels unchanged.

**Voice**: no em or en dashes. No site, sim, APL, band or preset words in the prose; "site" appears only as a frontmatter source kind. No quoted nightly numbers.

## Shoulds (wording, not findings)

- feral.md:67, "keep that row at ten". The ten points are the first two rows together; Feral Instinct's row holds five. Suggested: "with Brutal Impact gone, they are the two points that bring the tree to the ten that open the next row."
- Carried from round 15:
  - feral.md:89: explain melee haste.
  - restoration.md:101: "Herbalism and Alchemy keep you in mana potions and elixirs".
  - feral.md:67: Feral Instinct hides your Prowl on the walk in.
  - feral.md:73: "let it run into its last two seconds".
  - restoration.md:81: drop "(not 3)".
  - balance.md:119: name Frozen Heart of the Mountain as the bind-on-pickup reason for Enchanting.
  - balance.md:105: explain spell haste.
  - balance.md:105: the resisted-Wrath Eclipse line (sim round 17 L1).
  - feral.md:77: "(55 with Wolfshead Helm)".
  - feral.md:85: "your crits".
  - balance.md:91: "almost always".

## Decisions

1. The round 15 bear Professions line, including the rendered-review leveling sentence, is right and stays.
2. Rounds 1 to 15 decisions stand. These include:
   - Swiftmend keeps the HoT.
   - Gift of Nature misses Tranquility.
   - Pounce and Ravage need you behind the target.
   - Bear Furor is 1 point.
   - Feral Charge (Cat) is the 30 s leap.
   - Nature's Reach covers melee hit.
   - Rejuvenation and Healing Touch rank 11 are Ahn'Qiraj book ranks.

## Sims requested

None new. Carried: druid-feral Bite energy pooling under Predatory Discipline. Sim round 18 owns everything from round 15.

## Sim findings

Carried, still open:
- For a launch log:
  - Sim round 17 L1, Eclipse on a resisted Wrath.
  - S1/S2, the Cat Form spread and the unarmed paw.
  - X2, Read Leyline with Innervate.
- For the BiS lane:
  - Balance raid trinket 2, Uther's Strength on the Alliance.
  - The BiS talent strings are not the shipped builds (post-launch re-rank).

## Sources

- Client 1.60.1.70755:
  - talents/druid.json (texts, tier, column, prerequisites)
  - trainables/druid.json
  - spelltips/druid.json
  - races.json
  - items.json and items/druid.json (8345, 16608, 263411, every crafted best in slot piece named above)
  - addon-builds.json
  - bis/druid-balance.json, druid-feral.json, druid-feral-bear.json (every band), druid-restoration.json
- data/curated/rank-choices.json.
- design/reviews/guide-accuracy/round15/druid.md and ledger.md.
- Frontmatter URLs fetched 2026-11-04 (status above):
  - https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid
  - https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
  - https://talentsforever.com/data.json
  - https://foreversixty.gg/sim/weights
  - https://foreversixty.gg/sim/drops
  - https://foreversixty.gg/bis/druid/restoration
  - https://foreversixty.gg/bis/druid/feral-bear

## Fix log (round 16)

- feral.md:67 SHOULD: fixed, "they are the two points that bring the tree to the ten that open the next row" (Feral Instinct is tier 1, max 3).

# Guide accuracy round 14: Druid

Auditor: druid persona. Tree: gr14/audit at d7fc51fe7 (main after thirteen guide rounds and sim round 16; READ-ONLY). Client 1.60.1.70755, engine pin f637ec32c (sim/enginever/version.go; source read from gr14/engine). Before starting I read the round 1 to 13 druid reports and fix logs, round4/druid-levels.md, every race-sims file, ledger.md and the sim round 16 druid report (design/reviews/2026-11-04-sim-audit-round16/druid.md).

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`), Restoration (`restoration.md`), and the card notes in all four druid APLs. New this round: the live best in slot rows (de044125c, both factions, band 60 bare and raid) and race data checked against every gear, professions, race and stat line.

## What moved since round 13 (604b78a80 to d7fc51fe7)

- balance.md:135, the Levels 20 to 60 rank sentence (sim round 16 W2-3).
- feral.md:73, the Finishers paragraph (Rip at under 2 s, Bite at 5 s or more; sim round 16 C1).
- restoration.md:65, Wild Growth swap now "a second point of Improved Regrowth".
- restoration.md:81, the Ranks paragraph (tank Rejuvenation by level) and :83, the Read Leyline sentence.
- data/curated/apl/druid-feral.json [6] and [7]: Rip and Bite thresholds and notes.
- data/curated/apl/druid-restoration.json: new Read Leyline line [2].
- data/curated/rank-choices.json: druid-balance rows (Moonfire, Wrath, Insect Swarm) and druid-restoration tank Rejuvenation rows rewritten level by level.

talents/druid.json, trainables, races.json, combos.json and bis/druid-*.json did not change. No druid build string or addon-builds.json druid order changed.

## Decoded builds (all legal, 51 points, tier gates and prerequisites met)

- Balance leveling `5232220115101151/0/555301`: Improved Wrath 5, Genesis 2, Moonglow 3, Improved Moonfire 2, Nature's Majesty 2, Nature's Reach 2, Nature's Splendor 1, Insect Swarm 1, Vengeance 5, Improved Starfire 1, Nature's Grace 1, Eclipse 1, Moonfury 5, Moonkin Form 1 / Nature's Focus 5, Heart of the Wild 5, Naturalist 5, Reflection 3, Subtlety 1.
- Balance raid `5232220115501351/0/505003`: as leveling but Improved Starfire 5, Eclipse 3 / Nature's Focus 5, Naturalist 5, Subtlety 3.
- Feral leveling `050022/50212032123012123001/05501`: Genesis 5, Nature's Majesty 2, Nature's Reach 2 / Ferocity 5, Feral Swiftness 2, Feral Instinct 1, Brutal Impact 2, Shredding Attacks 3, Savage Fury 2, Feral Charge 1, Sharpened Claws 2, Predatory Strikes 3, Predatory Discipline 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 3, Shifting Power 1 / Heart of the Wild 5, Naturalist 5, Natural Shapeshifter 1.
- Feral raid `050022/50230032023012125001/055`: Feral Instinct 3, Rend and Tear 5, no Brutal Impact, Feral Charge or Natural Shapeshifter.
- Bear `0/5123033202310212551/055`: Ferocity 5, Furor 1, Feral Swiftness 2, Feral Instinct 3, Thick Hide 3, Shredding Attacks 3, Savage Fury 2, Sharpened Claws 2, Predatory Strikes 3, Lacerate 1, Blood Frenzy 2, Leader of the Pack 1, Natural Instinct 2, Rend and Tear 5, Natural Reaction 5, Primal Bite 1 / Heart of the Wild 5, Naturalist 5.
- Restoration `05302001/0/555300515311321`: Genesis 5, Moonglow 3, Nature's Majesty 2, Nature's Splendor 1 / Nature's Focus 5, Heart of the Wild 5, Naturalist 5, Reflection 3, Gift of Nature 5, Bloom 1, Tranquil Spirit 5, Improved Rejuvenation 3, Swiftmend 1, Nature's Swiftness 1, Living Spirit 3, Improved Tranquility 2, Improved Regrowth 1.

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| DRUID-R14-1 | CONFIRMED | rank (card note) | data/curated/apl/druid-restoration.json priorityList[8] note (tank Rejuvenation, 9841) | "A raid fight's mana pays for the rank that heals the tank most; from 28 to 49, rank 4 keeps your mana going." | rank-choices.json druid-restoration, authored 9841, now has four rows: 20-37 rank 2, 38-49 rank 4, 50-52 rank 6, 58-59 rank 8 (sim round 16 W2-2, 4724faf89). From 28 to 37 the card casts rank 2, not rank 4, and the 50-52 and 58-59 rows are missing. The same page's Ranks paragraph (restoration.md:81) gives the new rows, so the card note now contradicts the prose beside it. The rows were rewritten in 4724faf89 and only the prose followed. | Note: "On the tank whenever it is not ticking, at rank 10 in a raid. Let it run out before you recast it, so Bloom heals him when it does. While leveling your mana runs out first, so a lower rank heals more for it: rank 2 from 20 to 37, rank 4 from 38 to 49, rank 6 from 50 to 52 and rank 8 at 58 and 59." Also correct the top-level notes' "rank 4 from 28 to 49" to the same rows (not rendered, but it is the file's own record). |
| DRUID-R14-2 | CONFIRMED | professions vs BiS | restoration.md:101 | "Herbalism and Alchemy is the healer pair: your own mana potions and elixirs." | The page's own best in slot, band 60 raid, both factions (bis/druid-restoration.json): trinket Dormant Heart of the Mountain 249473 (Enchanting, items/druid.json bonding "pickup") and hands Mender's Leather Gauntlets 252551 (Leatherworking, bonding "pickup"). Two bind-on-pickup crafts in the set, and the line steers you to two professions that can make neither. Every other craft in the set (Living Crown, Elderwild Pants, Mooncloth Vest, Mooncloth Boots, Mooncloth Shoulders, Wisdom of the Timbermaw, Idol of Swiftness) binds when equipped. The Restoration shaman page already says this for the same trinket (shaman/restoration.md:110). | "Your raid set asks for two crafts that bind on pickup: Dormant Heart of the Mountain from Enchanting and Mender's Leather Gauntlets from Leatherworking, so take one of them and you make that piece yourself. Every other craft in your set binds when equipped and can be bought. Herbalism or Alchemy keeps you in mana potions and elixirs, which is most of what a healer spends gold on." |
| DRUID-R14-3 | CONFIRMED | professions vs BiS | feral.md:105 | "Skinning and Leatherworking is the classic cat pair: leather is your armor, and Leatherworking turns it into gear." | bis/druid-feral.json band 60 raid: Molten Heart of the Mountain 249470 in both factions, and Frozen Heart of the Mountain 249469 too on the Alliance; both are Enchanting crafts with bonding "pickup". Every Leatherworking piece in the set (Outlaw's Collar 279253, Timbermaw Tunic 252484, Might of the Timbermaw 19044, Warbear Woolies 15065) binds when equipped. So Leatherworking is not needed for anything in the set, and the one profession the set needs is never named. The rogue pages say this for the same trinkets (rogue/combat.md:116). | "Enchanting is the one profession your raid set asks for: Molten Heart of the Mountain and Frozen Heart of the Mountain are bind-on-pickup trinkets only an enchanter can make. The leather pieces in your set bind when equipped, so buy them; Skinning and Leatherworking still pay for themselves. Engineering works for a cat in Forever: the EZ-Thro and SAF-T bombs and grenades throw from Cat Form, though the old bombs still do not." |

No LIKELY or NEEDS-SIM findings.

## Checks

**Round 16 text re-checked**
- balance.md:135 against rank-choices.json druid-balance and trainables (Moonfire r2 at 10, r3 16, r4 22; Wrath r3 14, r4 22; Insect Swarm r1 the talent, r2 30):
  - Moonfire rank 2 for 20-29, rank 3 for 30-35 and rank 4 for 36-39 all match the rows.
  - Wrath rank 3 for 22-26 and rank 4 for 30-39 match the rows.
  - Insect Swarm rank 1 for 30-38 matches its row. The leveling build takes Insect Swarm at 27.
  - "Everywhere else, and from 40 on, cast the newest rank": no row reaches 40, and 27-29 fall to the default.
- restoration.md:81 against the authored 9841 rows: 2, 4, 6 and 8 at the levels stated. The newest rank at 53 to 57 is rank 9 (52), within 1% of the top per the row note.
- Restoration rank shares at 60 re-derived from the engine (`LevelPenalty`: full until learn level + 17, then 5% a level): rank 8 (46) full, rank 7 (40) 85%, rank 1 (4) 0. So "ranks 8, 9 and 10 keep all of it, rank 7 most of it, rank 1 none" holds. The party rows (8910: rank 1 at 10-59 and bare 60) match "rank 1 is your party Rejuvenation too".
- restoration.md:65: Improved Regrowth is 5 ranks at 10% each and the build has 1, so "a second point" is right. Dropping Improved Tranquility 2 leaves 42 points in tiers 0-4 (gate 25 for Improved Regrowth, 30 for Wild Growth), and Wild Growth's prerequisite (Living Spirit 3) is met.
- feral.md:73 and the [6]/[7] notes against the APL: Rip at 5 points when it is down or under 2 s, Bite at 5 points with 5 s or more, hold the points between. The engine's `Dot.Apply` cancels the old Rip (sim round 16 C1), so "throws away the ticks it had left" holds. Rip 6 is 12 s in 2 s ticks.
- Read Leyline (restoration.md:83, :97, APL [2]): races.json gives 50% health and mana regeneration for 15 s on a 2 min cooldown, spell 1259705. Only High Order (Alliance) Skyborne register it (sim round 16 S3, `racials.go`). "An Alliance Skyborne" is correct.

**Leveling against addon-builds.json** (index 0 = level 10): unchanged since round 13 and still matched.
- balance.md:127: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19.
- feral.md:55 and :113: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Feral Charge 25, Shifting Power 40, then Genesis, Nature's Majesty, Nature's Reach, Heart of the Wild, Naturalist and Natural Shapeshifter at 60. "Brutal Impact, Feral Charge and that Natural Shapeshifter point finish Rend and Tear and fill Feral Instinct" is 4 points out and 4 in.
- feral-bear.md:115 and :121: Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19, Lacerate 25, Primal Bite 40.
- restoration.md:109: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19.

**Best in slot rows against prose** (band 60, bare and raid, both factions)
- feral.md:89 and feral-bear.md:91 ("a high-DPS two-hander", "a low-DPS caster staff costs you threat"): every feral and bear row wears The Unstoppable Force, a two-hand mace.
- Wolfshead Helm (feral.md:93, feral-bear.md:95) is in no row, and the prose does not claim it is.
- balance.md:119 Tailoring and Enchanting: the set's one bind-on-pickup craft is Frozen Heart of the Mountain (Enchanting). The Tailoring and Leatherworking pieces all bind when equipped. Not false (see Shoulds).
- feral-bear.md:107: every craft in the bear set binds when equipped (Outlaw's Collar, Warbear Harness, Hivethrasher's Girdle, Fel Cape).
- Stat prose on all four pages explains and does not rank. No line contradicts the band 60 raid scale factors:
  - Balance: spell haste, spell power, hit, arcane, crit, nature, Intellect.
  - Feral: melee haste, hit, Agility, Strength, crit, attack power.
  - Bear (Alliance): expertise, Stamina, Agility, Dodge, hit, armor, Strength. Defense and crit are insignificant on the Alliance raid row; the prose still only explains them.
  - Restoration: Spirit, MP5, Intellect, healing power, crit.

**Races** (races.json, combos.json; unchanged)
- Elune's Light, Quickness, Endurance, War Stomp, Wind Blessed and Read Leyline all read as the pages say.
- The bear line matches round4/race-sims.md: Night Elf takes about 1.6% less damage with more threat; Tauren has the steadier TMI.
- Skyborne is named only as the paid option on every page.

**Client values** re-checked where round 16 touched them:
- Feral Instinct text: Swipe +30% at 3, plus Prowl detection as if 3 levels higher.
- Shredding Attacks: Shred -18 and Lacerate -3 rage at 3.
- Furor rank 1 is 20%.
- Natural Shapeshifter is 30% at 3.
- Improved Tranquility: -60% cooldown, so 5 min becomes 2.
- Wild Growth heals the party over 7 s.

**Card notes**: every other note in the four APLs matches its action and the client values. Examples:
- Feral [8]/[9]: Claw 40 against Shred 54/48, "8 to 14 more energy".
- Bear: Maul 10 with Ferocity, Demoralizing Roar 204.
- Balance: Eclipse two charges.

**Sources**: every frontmatter URL returned 200 on 2026-11-04. That covers both Blizzard articles, the Wowhead Skyborne article, talentsforever data.json, /sim/weights, /sim/drops, /bis/druid/feral-bear and /bis/druid/restoration. All follow the round 4 source-link rule.

**Voice**: no em or en dashes, and no site, sim, APL, band or preset words in prose. No quoted nightly numbers.

## Shoulds (wording, not findings)

- feral.md:67: say what the raid build's Feral Instinct points do. Its text reduces your chance to be seen while Prowling, which is your Ravage approach, and that is why sim round 16 W3-2 kept it. Suggested: "The extra Feral Instinct points add nothing to a cat's damage, but they keep you hidden on the walk in to your Ravage, and with Brutal Impact gone they keep that row at ten."
- feral.md:73: "so let it run to its last tick" sits beside a recast at under 2 s, which drops that last tick. Better: "so let it run into its last two seconds".
- restoration.md:81: "rank 2 (not 3)" only means something at 20 and 21. Drop "(not 3)".
- balance.md:119: name the bind-on-pickup trinket. Suggested: "Enchanting makes Frozen Heart of the Mountain, a bind-on-pickup trinket in your set, and puts spell power on your gear; the Tailoring and Leatherworking pieces bind when equipped, so you can buy them."
- balance.md:105: the stat paragraph never explains spell haste, though the frontmatter selects it. Suggested addition: "Haste is faster casts, which is more Starfires before the boss dies."
- Carried from round 13:
  - balance.md:105 "a resisted Wrath banks no Eclipse charges" (engine banks on cast complete).
  - feral.md:77 "(55 with Wolfshead Helm)".
  - feral.md:85 "your crits".
  - balance.md:91 "almost always".

## Decisions

1. Sim round 16's druid text is right:
   - the Balance leveling ranks
   - the Restoration tank ranks
   - the Rip and Bite timing
   - Read Leyline for Alliance Skyborne
   - the Wild Growth swap

   The one miss is the tank Rejuvenation card note (R14-1).
2. Profession lines follow the bind-on-pickup crafts in the live best in slot set, as the rogue and shaman pages already do. A craft that binds when equipped is bought, not trained for.
3. Rounds 1 to 13 decisions stand. These include:
   - Swiftmend keeps the HoT.
   - Gift of Nature misses Tranquility.
   - Pounce and Ravage need you behind the target.
   - Bear Furor is 1 point.
   - Feral Charge (Cat) is the 30 s leap.
   - Nature's Reach covers melee hit.
   - Rank 11 Rejuvenation and Healing Touch are Ahn'Qiraj book ranks.

## Sims requested

None new. Carried:
1. R10-12: druid-feral Rip against Bite at 5 points. Sim round 16 C1 has now measured the thresholds (2/5 best), so this one is closed.
2. Bite energy pooling under Predatory Discipline.
3. S10: the Claw to Shred gate. Sim round 16 C2 measured it: under the bar, aligned. Closed.

## Sim findings

- Carried, minor: Eclipse banks charges on a resisted Wrath in the engine (OnCastComplete). Settle from the client proc flags of 408248 or a combat log.
- Carried: S1 and S2 from sim round 16 (form damage spread, unarmed paw), S11 bear expertise weight, S12 bear Defense noise.
- For the BiS lane, carried from sim round 16 L3: Balance raid trinket 2 Uther's Strength.

## Sources

- Client 1.60.1.70755:
  - talents/druid.json
  - trainables/druid.json
  - races.json
  - combos.json
  - items/druid.json (bonding of 249469, 249470, 249473, 252551, 252561, 279250, 279252, 279253, 279254, 252484, 19044, 15064, 15065, 275614, 14139, 279269, 8345, 19323)
  - addon-builds.json
  - bis/druid-balance.json, druid-feral.json, druid-feral-bear.json, druid-restoration.json (band 60, bare and raid, both factions)
- data/curated/rank-choices.json and data/curated/apl/druid-*.json.
- Engine f637ec32c:
  - sim/core/spell.go `LevelPenalty` (grace 12, MaxLevel = learn + 5, 5% a level)
  - sim/druid
- design/reviews/2026-11-04-sim-audit-round16/druid.md (S3, C1, W2-2, W3-1, W3-2); guide-accuracy round4/race-sims.md.

## Fix log (round 14)

- DRUID-R14-1: fixed, druid-restoration.json tank Rejuvenation (9841) note and the top-level notes. Ranks now match rank-choices.json: 2 for 20-37, 4 for 38-49, 6 for 50-52 and 8 for 58-59.
- DRUID-R14-2: fixed, restoration.md:101. Dormant Heart (Enchanting, BoP) and Mender's Leather Gauntlets (Leatherworking, BoP) are in both raid rows.
- DRUID-R14-3: fixed, feral.md:105. Worded "Molten Heart, and on the Alliance Frozen Heart beside it", since the Horde raid row wears Molten only.
- SHOULDs: not applied.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).
- Rendered-review MUST fixed: feral.md Professions now says the Leatherworking pieces bind when equipped (the reputation and drop leather binds on pickup).

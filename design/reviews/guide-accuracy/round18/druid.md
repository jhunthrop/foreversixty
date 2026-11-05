# Guide accuracy round 18: Druid

Auditor: druid persona. Tree: gr18/audit at 588a3ce2b (main after seventeen guide rounds plus sim round 19; READ-ONLY). Client 1.60.1.70755 (the live 70790 carries the same talents, items, racials and numbers), engine pin da54b7b06 (sim/enginever/version.go; source read from gr18/engine). Before starting I read the round 17 druid report with its fix log and merge note, the round 16 report, ledger.md through round 17, the sim round 19 druid report (design/reviews/2026-11-04-sim-audit-round19/druid.md), and skimmed rounds 1 to 15.

Specs audited: class page (`index.md`), Balance (`balance.md`), Feral cat (`feral.md`), Feral bear (`feral-bear.md`) and Restoration (`restoration.md`), every line, plus the card notes in the four druid APLs and their rendered rows in addon-data.json at every level band (10 to 60). The sim round 19 edits and round 17's merge of them were audited as new work.

## What moved since round 17 (919f7e8a7 to 588a3ce2b)

`git diff 919f7e8a7 HEAD` on the druid files:

- balance.md:135 and :137: "Insect Swarm stays at rank 1 until 39" and a new paragraph for the two round 19 rows (Wrath rank 6 at 46 to 48, Insect Swarm rank 3 at 54 to 58).
- restoration.md:81 to :85: Ranks is now three paragraphs (the rule at 60; Ranks while leveling; the Healing Touch paragraph with round 17's Nature's Swiftness sentence); :121 Levels 20 to 60 names each tank rank switch.
- data/curated/apl/druid-restoration.json: the rotation notes, the tank Rejuvenation note [8] and the Healing Touch note [9].
- data/curated/rank-choices.json druid rows: tank Rejuvenation 38-43/44-48/49-56/57-59, party Rejuvenation 40-42, Healing Touch 30-33/34-39/40-45/46-49, Wrath 46-48 rank 6, Insect Swarm 54-58 rank 3.
- data/builds/1.60.1.70755/addon-data.json: regenerated from those rows.
- Unchanged: the four druid build strings, addon-builds.json orders, talents/druid.json, trainables/druid.json, spelltips, races.json, items/druid.json, bis/druid-*.json, the Balance, Feral and bear APLs, index.md, feral.md, feral-bear.md.

## Decoded builds (unchanged; all legal, 51 points)

Decoded in rounds 16 and 17 through talents/druid.json `tier`/`column`; no string changed.

- Balance leveling `5232220115101151/0/555301`, raid `5232220115501351/0/505003`.
- Feral leveling `050022/50212032123012123001/05501`, raid `050022/50230032023012125001/055`.
- Bear `0/5123033202310212551/055`.
- Restoration `05302001/0/555300515311321`.

## Findings

No CONFIRMED, LIKELY or NEEDS-SIM finding on main (588a3ce2b).

## Shoulds (wording, not findings)

- balance.md:135: "Insect Swarm stays at rank 1 until 39" (carried from round 17). The row ends at 38 and rank 2 wins at 39, so "until 39" is right if read as the switch level, but "through 38" cannot be misread. Suggested: "Insect Swarm stays at rank 1 through 38."
- balance.md:115 and restoration.md:101: "entering a raid removes it for good" (carried). Blizzard's notes say Attuned "is removed upon entering a raid" and nothing about reading a leyline again afterwards. Suggested: "but entering a raid removes it, so it never comes with you to a boss."
- restoration.md:57: "your next Healing Touch is instant" (carried). The talent covers your next Nature spell. Suggested: "your next Nature spell is instant, and for you that is a Healing Touch on a dying tank."
- restoration.md:83: "(rank 3 is not worth it)" (carried in another form since round 15). The bracket breaks the rank ladder's rhythm; the row's reason is that rank 2 beats every newer rank at 20 to 37, not only rank 3. Suggested: "rank 2 from 20 through 37".
- Carried from rounds 15 and 16 (none applied yet):
  - feral.md:89: explain melee haste ("faster paw swings, and every swing can proc Omen of Clarity").
  - restoration.md:105: "Herbalism and Alchemy keep you in mana potions and elixirs".
  - feral.md:67: Feral Instinct hides your Prowl on the walk in.
  - feral.md:73: "let it run into its last two seconds".
  - balance.md:119: name Frozen Heart of the Mountain as the bind-on-pickup reason for Enchanting.
  - balance.md:105: explain spell haste.
  - balance.md:105: the resisted-Wrath Eclipse line (sim round 18 S1 is still open on a log).
  - feral.md:77: "(55 with Wolfshead Helm)".
  - feral.md:85: "your crits".
  - balance.md:91: "almost always".

## Checked and right

**Sim round 19's Restoration rank lines (this round's new work)**, restoration.md:83, :85 and :121, druid-restoration.json notes [8] and [9], against the rank-choices.json druid rows and the rendered rows in addon-data.json at every band:

- Tank Rejuvenation (authored 9841): rows rank 2 at 20-37, rank 4 at 38-43, rank 5 at 44-48, rank 6 at 49-56, rank 8 at 57-59. The rendered tank line casts 1058 (rank 2) at 20, 24, 26, 30, 32, 34 and 36; 2090 (rank 4) at 38, 40, 42 and 43; 2091 (rank 5) at 44, 46 and 48; 3627 (rank 6) at 49, 50, 54 and 56; 9839 (rank 8) at 57; 9841 (rank 10) at 60. The page's "rank 2 from 20 through 37, rank 4 from 38 to 43, rank 5 from 44 to 48, rank 6 from 49 to 56 and rank 8 from 57 to 59" and note [8] say exactly that. "Below the top rank all the way to 59" holds: Rejuvenation trains at 4/10/16/22/28/34/40/46/52/58 (trainables/druid.json), so the newest rank at 38-43 is 6 or 7, at 44-48 7 or 8, at 49-56 8 or 9, at 57-59 9 or 10, and every row rank sits under it. Levels 20 to 60 (:121) names the same switch levels, 38, 44, 49 and 57.
- Party Rejuvenation (authored 8910): rows rank 1 at 10-39, rank 2 at 40-42, rank 1 at 43-59, rank 1 at bare 60. Rendered: 774 (rank 1) at every band from 10 to 38 and from 43 to 57; 1058 (rank 2) at 40 and 42; 8910 (rank 7) at 60. The page's "Rank 1 is your party Rejuvenation, except at 40 to 42" (:83 and :121) matches; the raid rule at :81 (rank 7 on a hurt party member, rank 1 when the fight will run you dry) matches the 60 rows.
- Healing Touch filler (authored 9889, line 9 only): rows rank 3 at 30-33, rank 4 at 34-39, rank 5 at 40-45, rank 6 at 46-49. Rendered filler line: 5187 (rank 3) at 30 and 32; 5188 (rank 4) at 34, 36 and 38; 5189 (rank 5) at 40, 42, 43 and 44; 6778 (rank 6) at 46, 48 and 49; the top rank below 30 (5186 at 10 and 12, 5187 at 14 to 18, 5188 at 20 and 24, 5189 at 26) and from 50 (9888 at 50 and 54, 9889 from 56). The page's "rank 3 at 30 to 33, rank 4 at 34 to 39, rank 5 at 40 to 45 and rank 6 at 46 to 49. Below 30 and from 50, cast the top rank" and note [9] match. "A lower rank" holds at every stretch: Healing Touch trains at 1/8/14/20/26/32/38/44/50/56, so the newest rank is 5 or 6 at 30-33, 6 or 7 at 34-39, 7 or 8 at 40-45 and 8 at 46-49.
- The Nature's Swiftness Healing Touch line (line 4) casts the top rank at every band (5186 at 10, 5187 at 14, 5188 at 20, 5189 at 26, 6778 at 32, 8903 at 38, 9758 at 44, 9888 at 50, 9889 at 56 and 60), which is what :85 says. Round 17's sentence is right on the client: Nature's Swiftness (17116) "When activated, your next Nature spell becomes an instant cast spell"; rank 10 Healing Touch still costs 755 mana under it (spelltips/druid.json), so "instant", not free.
- Regrowth and Tranquility cast the newest rank at every band (Regrowth 8936/8938/8939/8940/8941/9750/9856/9857/9858 at 12/18/24/30/36/42/48/54/60; Tranquility 740/8918/9862/9863 at 30/40/50/60), as "Use the top rank of Regrowth" and note [6] say.
- The rule at 60 (:81) still holds at the da54b7b06 pin (sim/core/spell.go: levelPenaltyGrace 12, levelPenaltyPerLevel 0.05, a rank's max level five past its learn level): ranks 8 to 10 (learned 46, 52, 58) keep all of it, rank 7 (learned 40) keeps 85%, rank 1 (learned 4) keeps none.

**Sim round 19's Balance rank lines**, balance.md:135 and :137, against the rows and the rendered Balance card:

- Wrath (authored 9912): rows rank 3 at 22-26, rank 4 at 30-39, rank 6 at 46-48. Rendered: 5178 (rank 3) at 20 and 26; 5179 (rank 4) at 27, 30, 34, 36, 38 and 39; 6780 (rank 6) at 40, 42 and 46; 8905 (rank 7) at 49, 50 and 52; 9912 (rank 8) from 54. Wrath trains at 1/6/14/22/30/38/46/54, so rank 7 is the new rank at 46 and "rank 6 from 46 to 48" is an old rank edging out a new one, as the page says.
- Insect Swarm (authored 24977): rows rank 1 at 30-38, rank 3 at 54-58. Rendered: 5570 (rank 1) at 20 to 38; 24974 (rank 2) at 39; 24975 (rank 3) at 40, 42, 46, 49, 54 and 58; 24976 (rank 4) at 50, 52 and 59; 24977 at 60. Insect Swarm trains at 20/30/40/50/60, so rank 3 at 54-58 is the old rank against rank 4.
- Moonfire rank 2 at 20-29, rank 3 at 30-35, rank 4 at 36-39 and the newest rank from 40 (8929 at 40 and 42, 9833 at 46 and 49, 9834 at 52, 9835 from 58), and Starfire's newest rank everywhere (2912/8949/8950/8951/9875/9876 at 20/26/34/42/50/58): "Everywhere else, cast the newest rank" holds.

**Every Innervate, Read Leyline and Attuned line** (balance.md:58, :63, :95 and :115; restoration.md:36, :47, :87 and :101; feral.md:51 and :67; druid-restoration.json [1] and [2]; druid-balance.json [6]):

- Innervate (29166, spelltips/druid.json): "Increases the target's Mana regeneration by 400% and allows 100% of the target's Mana regeneration to continue while casting. Lasts 20 sec.", 6 min cooldown, 5% of base mana, trained at 40 (trainables/druid.json). "Innervate is yours from level 40", "it multiplies your Spirit regeneration", "give it to a healer with high Spirit, because Innervate multiplies Spirit regeneration" all hold; the engine (sim/core/buffs.go InnervateAura) adds 4 to SpiritRegenMultiplier and forces full regen while casting. Note [1] casts it below 30% mana on yourself; the Balance Wrath note [6] fires only while Innervate is on cooldown, which is what :95 says.
- Read Leyline (1259705, spelltips/shared.json): "increasing your Health and Mana regeneration by 50% for 15 sec. If a leyline is nearby, you will become Attuned for the next 60 min and the effect of this ability will be increased by 100%.", 120000 ms cooldown. The pages' 50%, 15 seconds, 2 minutes, "Attuned for an hour" and "Attuned doubles Read Leyline's regeneration" match. It is the Alliance Skyborne racial (races.json 95, High Order Skyborne).
- Blizzard's Forever notes (26 October 2026, repeated 3 November; fetched 2026-11-04): "Read Leyline now increases Mana and Health regeneration by 50% (was 100%), with a 15-second duration." "This ability cannot be used while Innervate is active, and will be replaced by Innervate as a more powerful effect." "Using Read Leyline at a Ley Line now also grants "Attuned" for 1 hour." "While Attuned, the benefit of Read Leyline is increased to 100% regeneration increase." "Attuned persists through death, and is removed upon entering a raid." The pages and note [2] say exactly that (the "for good" wording is a Should above). The engine at da54b7b06 models both halves (sim/core/racial_read_leyline.go: not castable under any Innervate aura; buffs.go:1182: Innervate ends Energized).
- Wind Blessed: the 16 October note raised it to 3%, the 26 October note "reduced to 2% (was 3%)", and the client says "Increases your spellcasting, melee, and ranged Haste by 2%." The pages say 2%.

**Every talent number the pages quote**, re-read against talents/druid.json (unchanged since round 17): Improved Wrath 0.5 s and 50%; Eclipse 2 Starfires, 0.50 s at 3 points, stores 4; Nature's Grace, non-periodic spell crits, 10% for 3 s; Moonkin Form, Omen of Clarity +100%, 3% crit for party or raid within 45 yards, no healing spells (the 3 November note says the bonus "is again an increase to all types of critical strikes", so "on spells and attacks alike" holds); Leader of the Pack 3% within 45 yards; Furor rank 1 a 20% chance of 10 rage into Bear Form, rank 5 all the Cat Form energy plus 10 a second outside bear forms; Feral Instinct 10% Swipe a point; Nature's Swiftness needs Naturalist; and every other number in the round 17 list (Genesis, Moonglow, Nature's Majesty, Nature's Reach, Nature's Splendor, Vengeance, Improved Starfire, Moonfury, Ferocity, Feral Swiftness, Brutal Impact, Thick Hide, Shredding Attacks, Savage Fury, Sharpened Claws, Predatory Strikes, Lacerate, Predatory Discipline, Blood Frenzy, Natural Instinct, Rend and Tear, Natural Reaction, Primal Bite, Shifting Power, Nature's Focus, Heart of the Wild, Naturalist, Reflection, Natural Shapeshifter, Subtlety, Gift of Nature, Bloom, Tranquil Spirit, Improved Rejuvenation, Swiftmend, Living Spirit, Improved Tranquility, Improved Regrowth, Wild Growth).

**Leveling** (addon-builds.json unchanged; decoded again this round by tree, tier and column):
- Balance leveling: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Improved Moonfire 20-21, Nature's Majesty 22-23, Nature's Reach 24-25, Nature's Splendor 26, Insect Swarm 27, Vengeance 28-32, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonfury 36-39, Moonkin Form 40, Moonfury 41, Nature's Focus 42-46, Heart of the Wild 47-51, Naturalist 52-56, Reflection 57-59, Subtlety 60.
- Feral leveling: Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, Savage Fury 23-24, Feral Charge 25, Sharpened Claws 26-27, Predatory Strikes 28-30, Predatory Discipline 31, Blood Frenzy 32-33, Leader of the Pack 34, Natural Instinct 35-36, Rend and Tear 37-39, Shifting Power 40, Genesis 41-45, Nature's Majesty 46-47, Nature's Reach 48-49, Heart of the Wild 50-54, Naturalist 55-59, Natural Shapeshifter 60. feral.md:55 ("31 Feral points are done at 40. From 41 take Genesis, then Nature's Majesty and Nature's Reach in Balance, then Heart of the Wild, Naturalist and a point of Natural Shapeshifter") and :119 ("once Shredding Attacks has it at 42 energy") and the Feral card notes [8] and [9] ("until Shredding Attacks is full at 22") follow it.
- Bear: Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19, Savage Fury 20-21, Feral Instinct 22, Thick Hide 23-24, Lacerate 25, Thick Hide 26, Shredding Attacks 27-29, Leader of the Pack 30, Sharpened Claws 31-32, Predatory Strikes 33-35, Blood Frenzy 36-37, Natural Instinct 38-39, Primal Bite 40, Rend and Tear 41-45, Natural Reaction 46-50, Heart of the Wild 51-55, Naturalist 56-60. feral-bear.md:115 and :121 follow it (Lacerate at 25, Primal Bite at 40).
- Restoration: Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19, Bloom 20, Nature's Focus 21-23, Naturalist 24-28, Gift of Nature 29, Nature's Swiftness 30, Gift of Nature 31-34, Tranquil Spirit 35-39, Improved Rejuvenation 40-42, Swiftmend 43, Living Spirit 44-46, Improved Tranquility 47-48, Improved Regrowth 49, Genesis 50-54, Moonglow 55-57, Nature's Majesty 58-59, Nature's Splendor 60. restoration.md:113 follows it.
- Every trainer level the five pages quote matches trainables/druid.json (the round 17 list, re-read: Wrath 1/6/14, Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18; Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14; Bear Form, Maul, Growl and Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul rank 2 at 18; Cat Form, Claw, Shred, Rip, Prowl and Omen of Clarity 20; Rake and Ferocious Bite 24; Lacerate 25; Poised and Ravage 32; Pounce and Frenzied Regeneration 36; Dire Bear Form, Innervate, Hurricane and Primal Bite 40; Barkskin 44; Gift of the Wild 50).

**Card notes** (the four APLs and their rendered rows): every note is true at every level it renders on. Restoration [1] Innervate under 30%, [2] Read Leyline, [3]/[4] Nature's Swiftness under 35% with the top-rank Healing Touch, [5] Swiftmend under 75% for 20% of base mana (18562 cost_pct 20, 15 s), [6] Tranquility with two of the party under 70%, [7] Regrowth, [8]/[9] the rank rows above, [10] party rank 7 at 60 and the rank 1 and 2 rows below it. Balance [1]/[2] end of fight, [3]/[4] the DoTs, [5] Eclipse, [6] Wrath behind on mana with Innervate on cooldown, [7] Starfire. Feral [1] Ravage from Prowl, [2] free Faerie Fire, [3]/[4] below 20, [5] Shifting Power at 60 energy or less on 8 s, [6] Rip under 2 s, [7] Bite at 5 s or more, [8]/[9] Shred against Claw at 20 and 21. Bear [1] Barkskin, [2] Frenzied Regeneration under half with 20 rage, [3] Primal Bite 6 s, [4] Lacerate five stacks and 3 s, [5] free Faerie Fire with threat, [6] Demoralizing Roar 204 at 60 and 3 s refresh, [7] Maul at 15 (10 with Ferocity), [8] Swipe above 40.

**Best in slot and professions**: bis/druid-*.json and items/druid.json did not change since round 17, so the round 17 checks stand (Molten and Frozen Heart of the Mountain for the cat, Frozen Heart for Balance, Dormant Heart and Mender's Leather Gauntlets for Restoration, nothing bind-on-pickup in the bear's raid set, the Defender's, Brawler's and Warden's leather at 20 to 50). Stat prose explains and never ranks.

**Races** (races.json, spelltips/shared.json): Elune's Light 10% crit with all spells and attacks for 15 s on 3 min; Quickness 1% dodge; Endurance 5% health; War Stomp 5 enemies within 8 yards for 2 s; Wind Blessed 2% haste; Read Leyline as above. Skyborne is named only as the paid option on every page.

**Sources**: every frontmatter URL fetched 2026-11-04 with a browser user agent. 200: both Blizzard articles, talentsforever data.json, /sim/weights, /sim/drops, /bis/druid/restoration, /bis/druid/feral-bear. 403: Wowhead Skyborne First Look (the bot block; unknown, not dead, as in rounds 16 and 17). Labels follow the round 4 rule.

**Voice**: no em or en dashes on any druid page or in the round 19 notes. No site, sim, APL, band or preset words in prose. No quoted nightly numbers (the rows' percentages stay in rank-choices.json).

## Decisions

1. Sim round 19's Restoration and Balance rank lines, as merged by round 17, stand: every level range on restoration.md:83, :85 and :121, balance.md:135 and :137, and notes [8] and [9] matches its rank row and the rendered card at every band, and every "lower rank" claim holds against the trainer levels.
2. Round 17's Nature's Swiftness sentence stands: the talent makes the next Nature spell instant and leaves its mana cost alone, and the card's save line casts the top rank at every level.
3. The Read Leyline, Attuned and Innervate wording stands on the client tooltips and Blizzard's 26 October and 3 November notes; "for good" stays a Should.
4. Rounds 1 to 17 decisions stand: Swiftmend keeps the HoT; Gift of Nature misses Tranquility; Pounce and Ravage need you behind the target; bear Furor is 1 point; Feral Charge (Cat) is the 30 s leap; Nature's Reach covers melee hit; Rejuvenation and Healing Touch rank 11 are Ahn'Qiraj book ranks.

## Sims requested

None. Nothing on the pages needs a measurement to settle, and the round 19 "measure at refresh" item (Restoration and Balance bands 40 and 50 on re-ranked gear, which can move a switch level by one) belongs to the sim lane after tonight's re-rank, not to this audit.

## Sim findings

- Carried, still open for a launch log: sim round 18 S1 (Eclipse on a resisted Wrath; balance.md:105 says a resisted Wrath banks none) and S2 (the Cat Form spread and the unarmed paw).
- Carried for the BiS lane: Balance raid trinket 2 is Uther's Strength on the Alliance; the BiS talent strings are not the shipped builds; the late druid hotfix in Feral and Balance BiS; the Feral raid build's Feral Instinct points against Furor or Thick Hide (already with the sim lane, not re-filed).
- Blizzard's 3 November note raises the bear form threat bonus from 50% to 115%; no page quotes a bear threat number, so nothing to change, but the engine's bear threat multiplier is worth a read at the next pin.

## Sources

- Client 1.60.1.70755 (data/builds/1.60.1.70755): talents/druid.json; trainables/druid.json; spelltips/druid.json (774, 1058, 2090, 2091, 3627, 8910, 9839, 9841, 5187, 5188, 5189, 6778, 9889, 6780, 8905, 9912, 24975, 24977, 17116, 18562, 29166) and spelltips/shared.json (Read Leyline, Wind Blessed, Elune's Light, Quickness, Endurance, War Stomp); spellranks.json; races.json; addon-builds.json; addon-data.json (rotations druid-restoration and druid-balance at every band, druid-feral and druid-feral-bear notes); bis/druid-*.json and items/druid.json (unchanged since round 17).
- data/curated/rank-choices.json (druid-balance and druid-restoration rows) and data/curated/apl/druid-*.json.
- Engine da54b7b06: sim/core/spell.go (levelPenaltyGrace, levelPenaltyPerLevel, LevelPenalty), sim/core/racial_read_leyline.go, sim/core/buffs.go (InnervateAura).
- design/reviews/2026-11-04-sim-audit-round19/druid.md; guide-accuracy round16/druid.md, round17/druid.md, ledger.md.
- Blizzard, "WoW Classic Forever Beta Patch Notes" (Kaivax; 16 October, 26 October, 27 October and 3 November 2026 sections): https://us.forums.blizzard.com/en/wow/t/wow-classic-forever-beta-patch-notes/2360696 (fetched 2026-11-04).
- Frontmatter URLs fetched 2026-11-04 (status above):
  - https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid
  - https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
  - https://talentsforever.com/data.json
  - https://foreversixty.gg/sim/weights
  - https://foreversixty.gg/sim/drops
  - https://foreversixty.gg/bis/druid/restoration
  - https://foreversixty.gg/bis/druid/feral-bear

## Counts

- CONFIRMED 0, LIKELY 0, NEEDS-SIM 0.
- Shoulds: 1 new (restoration.md:83 bracket), 3 carried from round 17, 10 carried from rounds 15 and 16.

## Fix log (round 18)

- No finding. The three wording SHOULDs (Insect Swarm "through 38", "removed upon entering a raid", the rank 3 bracket) are not taken this round.

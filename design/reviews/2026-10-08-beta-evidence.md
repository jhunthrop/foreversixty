# Beta evidence for open simulator inputs (researched 2026-10-08)

Grades: BLIZZARD (official statement/post), DATA (client/DB mirror of tables), TESTER, INFERENCE, SECONDARY (site summary of an unseen primary; verify).
Several pages were fetched through a summarising fetcher, so quotes marked SECONDARY should be re-checked against the primary.

## 1 Rogue energy

- BLIZZARD (via secondary): Blizzard Sep 30 class-highlights video; Icy Veins/guided.news: "Energy will no longer return in clearly separated intervals and will instead regenerate continuously." https://guided.news/en/news/wow-forever-shaman-totem-systems-rogue-overhaul/ , https://www.icy-veins.com/wow-forever/news/class-change-highlights-for-warrior-shaman-and-rogue-video/ (video gives no per-second number, per a search summary).
- TESTER: GitHub issue "they changed the rogue energy regeneration to a fluid regen, instead of ticks" https://github.com/wowaddonmaker/classicuiforever/issues/17
- SECONDARY: classicwow.gg "Energy now regenerates smoothly rather than in Classic's larger ticks"; Vigor "increases maximum energy, not regeneration" https://classicwow.gg/forever/guides/rogue ; max 100 repeated by guide sites.
- Blizzard Rogue deep dive (https://worldofwarcraft.blizzard.com/en-us/news/24310968) lists no regen-rate or baseline-max change. Flawless Execution: -10 Eviscerate cost; Improved Sinister Strike up to -5.
- GAP: no published rate. Smooth regen with 20/2s = 10/s is an inference only.

## 2 Eviscerate AP scaling

- DATA: Forever tooltip "Finishing move that causes damage per combo point, increased by Attack Power" (Rank 2: 14-22 /25-33 /36-44 /47-55 /58-66 for 1-5 CP) https://forever.rexas.tools/spell/6760/eviscerate
- Blizzard: "Improved Eviscerate: Moved to row 1 from Assassination. Damage bonus increased." and Flawless Execution (above).
- No Forever coefficient found. 0.03 x AP per CP is the Classic formula (vanilla-wow-archive / MMO-Champion), carried over by inference.

## 3 Dual-wield miss, glancing, weapon skill

- BLIZZARD: Deep Dive panel (recap 2026-09-26): weapon skill "Works as before, but less per item", "Stops weapon skill gear being the obvious best pick". https://wofwforever.com/en/guides/wow-forever-weapon-skill-scaling/ ; https://wowforeverguides.com/systems/itemization (also: Blizzard published no hit/crit target for level 60; hit merged across spell/melee/ranged)
- BLIZZARD (Kaivax known issues, Sep 24): glancing "chance ... correct at all levels for melee"; "The damage penalty of a glancing blow is correct against enemies of your own level. It is not correct against higher-level enemies." Character-sheet glancing numbers inaccurate; caster values wrong. https://www.indiekings.com/2026/09/wow-forever-beta-known-issues-pets.html , https://seramate.com/news/forever/2026/09/24/forever-beta-known-issues-update-1f1bced4
  -> implication: glancing chance formula is Classic; the penalty against +3 targets was a known bug on Sep 24, check whether fixed in Oct 1 notes (not confirmed).
- SECONDARY: Method.gg "Dual Wielding requires 24% Hit Chance to never miss with both weapons"; 29% vs raid bosses. https://www.method.gg/wow-forever/how-does-hit-chance-work-in-wow-forever (= 19% + 5% on a level-equal target, +3 level boss adds ~4%... consistent with Classic 19% penalty).
- BLIZZARD: Rogue deep dive: "Dual Wield Specialization: Increase to off-hand damage reduced." (conflicts with classicwow.gg "retains its Classic off-hand damage bonus ... up to 25%"; prefer Blizzard.) Beta notes: warrior Dual Wield Specialization hit bonus bug fixed so it applies to both weapons. https://www.bluetracker.gg/wow/topic/us-en/2360696-wow-forever-beta-development-notes-updated-october-1/
- No statement found that the white-hit miss penalty or glancing formula changed.

## 4 Windfury Totem

- DATA: Rank 3 tooltip "Summons a Windfury Totem with 5 health ... enhances the melee attacks of all party members within 30 yards. Each main hand hit has a 20% chance of granting the attacker 1 extra attack with 246 extra melee attack power. Lasts 5 min." (Rank1 95 AP, Rank2 179 AP; Classic 122/229/315, 20 yd, 2 min). https://theforeverera.com/en/spells/shaman/windfury-totem/ , https://foreverdiff.com/spells/windfury-totem/
- BLIZZARD: weapon procs including Windfury Totem now work while shapeshifted (hunter/druid deep dive) https://news.blizzard.com/en-us/article/24301515/world-of-warcraft-forever-class-deep-dives-hunter-and-druid ; "Pets can no longer receive player buffs that increase their stats" (same page; not specific to the totem).
- BLIZZARD (beta notes): Windfury no longer stacks with Tranquil Air/Grace of Air, nor Flametongue totem; buff now named "Windfury Totem". Known issue: "item enhancements (such as windfury ...) are not tracked" for pets (display).
- "Each main hand hit" does not say auto vs special; no ICD stated; no tester report on pets. Pets: UNKNOWN. Weak INFERENCE: "party members" aura, and Blizzard says pets don't receive player stat buffs.
- Note: forum snippet quotes Windfury Weapon "2 extra attacks, 433 AP" (an imbue, not the totem).

## 5 Shadowfiend and Dark Sacrifice

- DATA: Wowhead Forever: 15 s duration, 5 min cooldown, "Caster receives 5% mana when the Shadowfiend attacks"; Forever 1.60.1 change from 6% (Classic Era) to 5%. https://wow-forever.gg/db/spells/401977-shadowfiend/ (swing speed 1.5 and "up to 11 hits, usually 10" from a search summary of Icy Veins/other, SECONDARY: https://www.icy-veins.com/wow-forever/shadow-priest-ranged-dps-pve-guide, 403 on direct fetch).
- DATA: Dark Sacrifice (Undead racial, learned level 20, 5 ranks): R1 "Cannibalize 400 of your own Health over 15 sec to gain 400 (+100% of Spirit) Mana"; R5 (lvl 60) 1600/1600. https://foreverdb.net/spell/1277324 , https://www.wowhead.com/forever/spell=1277325/dark-sacrifice (Wowhead dump shows 137 mana/3 s dummy, 10 min cd, effect rows only). Blizzard Priest deep dive text: "Cannibalize your own health to gain Mana over time. Mana gain increased by Spirit." https://news.blizzard.com/en-us/article/24301514/world-of-warcraft-forever-class-deep-dives-priest-and-warrior
- Note the user's phrase "Dark Sacrifice" is a priest-racial; check engine handles it as such.

## 6 Spirit mana regen / five-second rule

- BLIZZARD: Priest deep dive: Meditation "Mana benefit from spirit increased substantially" (no number). Hunter deep dive references "Mana regeneration from Spirit while casting".
- SECONDARY: Icy Veins Forever Holy Priest guide describes the 5-second rule as in Classic (page 403 on direct fetch; seen via search summary) https://www.icy-veins.com/wow-forever/holy-priest-healer-pve-guide
- No Blizzard statement of a changed regen formula or FSR found. Classic reference (INFERENCE): warcraft.wiki.gg "Mana regeneration" level-60 formula via Whitetooth is the TBC-era one; the vanilla one differs (sqrt-free table-based); do not mix.

## 7 Rogue highest single-target DPS

- BLIZZARD: Principal Game Designer Kris Zierhut, "WoW: Forever Podcast" Episode 2 "Speed Running Classes ft. Sodapoppin" (published 2 Oct 2026; https://www.youtube.com/watch?v=uDRcv_w85j8). Quotes as reported: "Rogues will be the highest single target damage dealers in the game, given mechanics, given the encounter."; "If Rogues don't do the highest single-target damage, then why are there Rogues?"; "they will do the highest single target damage in the game and that's going to be one of the biggest reasons why you bring [a Rogue] with you."; "Rogues do so much damage from white attacks, plus Slice and Dice, plus poisons, that it is really easy to take them too far."; "Rogues pay nothing and should deal the most single target damage in a standing fight." Hybrids deal "about 5% less damage than pure specs"; rogue is a pure class with no raid utility ("pay nothing").
  URLs: https://us.forums.blizzard.com/en/wow/t/rogue-top-dps-class/2368299 ; https://seramate.com/news/forever/2026/10/01/rogues-highest-single-target-damage-1f1bdc19 ; https://wowforever.be/news/hybrids-deal-about-5-percent-less-damage/ ; https://x.com/nohitjerome/status/2106062499560800445 ; https://www.wowhead.com/forever/news/rogues-will-have-the-highest-single-target-damage-in-wow-forever-383209
- Reasoning rests on: no raid buffs/utility, white damage + Slice and Dice + poisons, standing (Patchwerk) fight. Not on any named Forever rogue change. Timestamp not found; quotes are via written summaries (SECONDARY wording), transcript not available.

## 8 Beta logs / parses

- None found. Search found only simulated numbers (foreverdb.net sim: level-30 Combat rogue 75.1 DPS on a 60 s fight, 90% between 67.0 and 82.9 - a SIM, not a log) https://foreverdb.net/sim/rogue/combat ; tier-list sites mention "the few dungeon logs uploaded so far, mixed gear" without links (https://leprestore.com/guides/world-of-warcraft-forever/forever-best-dps-tier-list/). No Warcraft Logs reports for the beta located. Treat as GAP.

## 9 Unified crit: do Agility crit and Intellect crit sum? (added 2026-10-08)

- Blizzard: "Spell, melee, and ranged critical chance are combined" (Deep Dive recap). The client's `PlayerExpectedStat` table (design/reviews/2026-10-08-expected-stat.md) gives every class both a `CritPerAgility` and a `SpellCritPerIntellect` rate, per level. What no table says is whether the server adds both into the one Crit number (the engine's model: Grace of Air's 77 Agility lifts a balance druid's Starfire crit 3.4 points, a mage's 195 Agility in the engine's own P1 test set lifts its crit 10 points) or keeps two pools and shares only item, talent and racial crit between them (foreverdb.net/stats and Output Lag read it that way, as an inference). Worth 2 to 3% of every caster's raid DPS and the whole value of Agility on caster gear and of Intellect on hunter, retribution and enhancement gear. GAP, testable.
- **The test (any beta character, under a minute):** open the character sheet of a druid, shaman, paladin, warlock, hunter or mage and note the Crit figure (and whether the sheet shows one Crit line or separate melee and spell lines). Equip or unequip any item with Agility and nothing else that matters (a green "of the Monkey" piece, or Grace of Air from a shaman) and note the Crit figure again; do the same with an Intellect item. If one Crit number moves with both stats, the engine's sum is right. If the sheet shows two lines, or one number that moves only with the class's traditional stat (Agility for the melee, Intellect for the casters), the engine should keep two pools and apply item, talent and racial crit to both. Note the level, class, the two readings and the item's stats.

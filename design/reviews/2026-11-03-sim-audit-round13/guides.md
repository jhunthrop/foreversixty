# Round 13: guides lane

Lane `guides`, site branch wt/r13-guides from int/r13 (93c5d2510). No engine branch. Scope: web/src/content/guides/** and the class landing data (data/curated/classes.json `forever_changes`, data/curated/guide-cooldowns.json). Client truth is 1.60.1.70753 (wago CSVs in scratchpad r11/c70753). The committed hotfix push (data/builds/1.60.1.70753/hotfix-push.json) is empty. A hunter hotfix went live on 2026-11-03 (Aimed Shot, Arcane Shot, Serpent Sting, Trueshot Aura). The hotfix lane owns the hunter rotation, priority and ability-value lines, so this lane did not touch them. Wowhead tooltips (read 2026-11-03) are a second source only where they agree with a client row.

Every fix is prose or landing copy. No DPS or HPS moves. Tests: `npx vitest run src/content/guides/` (`_prose`, `_sections`, `builds`) passes, 518 of 518, after each commit. data/tests/test_curated.py::test_the_committed_build_matches_the_curated_facts already failed at 93c5d2510, because round 12's classes.json edit is waiting on the regen. It still fails until the copy is regenerated (see "Files to regenerate").

Commits: S aae1e50a5, S 0b2004a0f, S 3a155b380, S 13497f823.

## 1. Regression over round 12's guide fixes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Unstable Affliction / Haunt line (warlock index:31): "your DoTs are Corruption, Immolate, your bane and, for Affliction, Siphon Life" | No SkillLineAbility row for either spell. Siphon Life is an Affliction tier-4 talent and only the Affliction builds take it. Wrack is a 6 s channel, which the index calls a channel. No other page names UA or Haunt. | Aligned | - | - |
| X2 | Rogue Races panel | Race list matches combos.json class 4, and no new_in_forever rogue pair exists. Ruthless 20597 is A290 +1. Artisan's Precision 1259719 is A54/55 +1. Wind Blessed 1259710 is 2% haste. Will of the Forsaken 7744 removes Charm, Fear and Sleep. Eureka! 1259812 is 8 s / 2 min, energy -10%, damage +10%. No racial on a race skill line grants weapon skill. | Aligned | - | - |
| X3 | Kidney Shot "25 energy, 20 s, up to 6 s at five" (three specs) | 408 and 8643: 25 energy, category 20000 ms. Rank 2 (50) runs 2 to 6 s. | Aligned | - | - |
| X4 | Eureka! on every page (mage, warlock, priest, rogue) | 1259812/13/17/21/23: 8 s, 2 min, no charges. Mage, warlock and priest rows add a periodic +10% (aura 108 misc 22). The warlock "DoTs get nothing" line holds, since the warlock periodic mask is the drains and Wrack only. race-table.ts wording matches every row. | Aligned | - | - |
| X5 | Scatter Shot → Freezing Trap (BM, MM) | Freezing Trap 1499 is a trainer spell at 20 (14310 at 40), 30 s category, freeze up to 20 s at 60 (14309). The coach copies still carry the old Scatter Shot line (regen). | Aligned | - | - |
| X6 | Three build point counts (Blood Frenzy, Lethality/Vile Poisons, Accelerant) | Decoded against talents/*.json: Feral leveling 2/2 and raid 1/2; Assassination raid Lethality 3 and Vile Poisons 2; Frost Accelerant 1 and Flame Throwing 2. | Aligned | - | - |
| X7 | Ruin "raise the crit bonus by 80%" | 4 of 5 points × 20%. | Aligned | - | - |
| X8 | Shaman crit pools (index, Restoration) | Matches the owner's ruling. Enhancement still said "one hit stat" twice (S2 below), which round 12 fixed only in the index. | Aligned (gap logged as S2) | - | - |
| X9 | Revive's cost | 75% of base mana, 10 s cast, from 12. | Aligned | - | - |
| X10 | Feral Genesis / Natural Shapeshifter | Leveling `01/50212032123012123001/535033` decodes to Genesis 1, Heart of the Wild 3, Natural Shapeshifter 3. Raid decodes to Genesis 5, Nature's Majesty 2, Natural Shapeshifter 3. Prose matches. | Aligned | - | - |
| X11 | Feral Claw/Shred line: "Claw's 45 energy buys more damage than Shred's 60" | Ferocity 5/5, a tier-0 point every build takes, makes Claw cost 40. Shredding Attacks takes 6 off Shred per point. A talented cat never pays 45 or 60, and the APL note itself says 40 against 54. | CONFIRMED (minor). New: "before that, Claw buys more damage for its energy." | R12 bafdb0c22 | S 13497f823 |
| X12 | Paladin class landing after S 9efecab6f | The fix removed Voice of Truth and Instrument of Law but left three names. Improved Holy Strike 1310902 has no tree slot and no trainer row. Improved Seal of Fury 1314103 is a trainer spell (SLA 267, class 2). Twist of Light has no row in 70753 at all. Vindication was already a 1.12 Retribution talent, so it is not a "gain" (1.12 tree; single source). | CONFIRMED. New: "Holy gains Reverence, Infusion of Light and Light's Vigil; Protection gains Shield Specialization, Swift Judgement, Templar's Bulwark and Iron Creed; Retribution gains Sacred Arbiter and Champion of the Light." | old 1e0a8ca39 (round 12 looked and missed) | S aae1e50a5, S 0b2004a0f |
| X13 | guide-cooldowns.json: Reincarnation and Inferno | 20608: level 30, category 3,600,000. 1122: level 50, category 3,600,000. Both are trainer rows. I also listed every trainer spell with a cooldown of 3 minutes or more against the file. The gaps (Retaliation, Intimidating Shout, Sprint, Dash, Divine Intervention, Ritual of Doom) are deliberate under the file's own rule (252455d15). Lightwell is carried by the Holy and Discipline rotation lists (rotation-lists.test.ts:464). | Aligned | - | - |
| X14 | Warrior index: "every warrior plays around rage: Bloodrage before the fight" | Contradicts Arms' round-12 opener (Charge first, because Bloodrage puts you in combat). The Arms APL casts no Bloodrage before the pull. | CONFIRMED (contradiction the R12 fix created). New: "Bloodrage at the pull". | R12 70aff3eaa | S 3a155b380 |
| X15 | Fury build change (Anger Management, no Improved Charge or Deflection) | Fury decodes to Arms 17 (Improved Heroic Strike 3, Improved Rend 3, Improved Tactical Mastery 5, Anger Management 1, Deep Wounds 3, Impale 2). The "8 points in Arms" bleed line still counts 3+3+2. | Aligned | - | - |

### Fact checks in forever-refresh-hold.md

| # | claim | evidence | verdict |
|---|---|---|---|
| F1 | Heart of the Lion: 10% stats plus attack power, 280 at 60 (hunter index) | 409583: A99/A124 at 40 + 4 a level, max level 60, BaseLevel 0, SpellLevel 1. The engine (buff_ranks.go:141) and the Wowhead tooltip both give 280. A server that scales from SpellLevel would give 276. | Aligned on two sources. OPEN only on the 4 AP: the in-game tooltip at 60 settles it. |
| F2 | Blessings last 1 h, Greater included | DurationIndex 3,600,000 ms on every trained rank of Might, Wisdom and Kings, and on both Greater ranks. | Aligned |
| F3 | Portal of Summoning: caster plus two others, two at the portal summon, five shards, 1 h | 437169: level 60, cooldown 3,600,000, 2 min channel. The tooltip gives five Soul Shards (no SpellReagents table in the cache, so the shard count is single-source). | Aligned |
| F4 | Chain Lightning 2.0 s cast | 421, 930, 2860, 10605: 2000 ms, 6 s category. | Aligned |

## 2. Fresh re-audit

Method: every page read in full. Builds decoded against talents/<class>.json. Each number checked with a lookup over SpellMisc, SpellLevels, SpellPower, SpellCooldowns, SpellEffect, SpellAuraOptions and SkillLineAbility (and ItemEffect for consumables). Rotation lines checked against the actions and conditions in data/curated/apl/<spec>.json. Stat prose checked against the published raid-60 order in bis/<spec>.json and stat-weights.json.

### Priest

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | index:29 "a heal learned below level 20 keeps its full healing power bonus" | No Forever source says so, and it is no change from 1.12. The engine applies the sub-20 penalty, 1-(20-lvl)×0.0375 (sim/core/spell.go:777-803, round 8 X2), so the claim contradicts every published number. | CONFIRMED. Bullet removed. | landing adcaad714 | S 0b2004a0f |
| P2 | discipline:76, the same claim for the filler | The APL's filler is Lesser Heal rank 3 (2053), learned at 10, 75 mana. | CONFIRMED. New: "Lesser Heal rank 3 is learned below level 20, so it loses part of its healing bonus, but at 75 mana it is still the right top-up." | landing adcaad714 | S 0b2004a0f |
| P3 | discipline:74 and holy:70 "every rank (of Renew) carries the same healing power bonus" | Renew rank 1 is learned at 8 and rank 2 at 14, so both take the penalty. Rank 3 (6075, level 20) is the rank the APLs cast. | CONFIRMED. New: "every rank from level 20 up". | landing adcaad714 | S 0b2004a0f |
| P4 | holy:88 "Spirit and mp5 ... Spirit the best of them" | The published Holy raid-60 order is mp5, then healing power, then Spirit, then Intellect, then crit. | CONFIRMED. The paragraph now follows that order, with no numbers. | landing adcaad714 | S 0b2004a0f |
| P5 | discipline:62 Twilight Focus "no longer pushes back" | Talent 14913 is a 70% chance. | CONFIRMED (minor). New: "rarely pushes back". | landing adcaad714 | S 0b2004a0f |

Aligned: Penance (3 pulses, 12 s, level 30), Divine Aegis, Soul Warding, the Power Infusion swap (tier-6 talent), Prayer of Mending (5 jumps), Heal rank 4 (2.5 s with Divine Fury), Lightwell (3 min or 5 clicks, 10 min cooldown), Shadowform (40% base mana), SW:Death (18 s, 30% backlash), Devouring Plague 24 s, Mind Blast 8 s, the race spells and race masks, Divine Grace, Dark Sacrifice, Contingency Plan at 35%, Elune's Grace -50%, and the Shadow opener and priority against the APL.

### Paladin

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Landing paladin[6]: "Holy Strike is on a 12 second cooldown and had reached rank 5 by level 38 ... taunts for 4 seconds ... where every other taunt is 5" | Holy Strike 679-10333: 10 s category. Rank 5 (2495) is learned at 36. Judgement of Fury 20232 lasts 3000 ms. The guides already say 3 s. | CONFIRMED. New: "Holy Strike is on a 10 second cooldown and reaches rank 5 at level 36; ... Judging it taunts for 3 seconds at 10 yards." | old 1e0a8ca39 | S 0b2004a0f |

Aligned (Holy, Retribution, Protection, index): every talent text and count. Flash of Light rank costs. Light's Vigil 1,340. Holy Shock rank 1 at half cost. Divine Favor 2 min. Oath of Light 5 min. Lay on Hands 20 min. Blessing of Protection 3 min with Guardian's Favor 2/2. Judgement 8 s. Hammer of Wrath made instant, 6 s, from 44. Exorcism and Holy Wrath on Undead and Demons. Repentance 6 s / 1 min. Holy Shield 10 s, 4 charges, ranks at 30/40/50/60. Swift Judgement 60 s. Templar's Bulwark 8 s. Righteous Fury 30 min for 30% base mana. Openers and priorities against the three APLs.

### Rogue

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Landing rogue[4]: "Restless Blades in Combat" | No such talent in the 70753 rogue tree. Mutilate, Venom, Cutthroat and Weapon Expertise (2% dodge/parry reduction) are present. | CONFIRMED. Removed. | old 1e0a8ca39 | S 0b2004a0f |

Aligned: all four pages' talent texts and counts. Backstab and Mutilate at 60 energy. Mutilate needs a dagger and has no facing line. Blade Flurry. Ghostly Strike 20 s. Premeditation 20 s / 2 min. Preparation 10 min. Redirect 15 s. Garrote openers, Cold Blood and finisher thresholds against the APLs. The hit, then Agility, then crit order.

### Shaman

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | elemental:91 "plus the Elements pieces from Stratholme and Blackrock Spire" | The published list has one Elements piece (Blackrock Spire) and none from Stratholme. The line also restated BiS. | CONFIRMED. Clause dropped; the page links BiS. | landing fdf6c87f4 | S 3a155b380 |
| S2 | enhancement:52 "because Forever has one hit stat"; :36 "One hit stat covers..." | Tidal Focus carries separate melee (A54) and spell (A55) hit effects. Owner ruling: the two pools are separate. | CONFIRMED. Reason dropped. :36 now "Hit on gear covers your swings and your spells". | landing b2e4a3bc5 | S 3a155b380 |
| S3 | enhancement:52, the leveling Elemental points leave out Concussion | Leveling Elemental decodes to Convection 3, Concussion 2, Elemental Alacrity 3, Call of Flame 3. | CONFIRMED. Concussion added. | landing b2e4a3bc5 | S 3a155b380 |
| S4 | enhancement:91, stat prose leads with Strength | Published raid-60 order: hit, then Intellect, crit, Agility, Strength. | CONFIRMED. Reordered (hit, Intellect, Agility, Strength), no numbers. | landing fdf6c87f4 | S 3a155b380 |
| S5 | Landing shaman[5] "Ghost Wolf's base cast is 2 seconds" | 2645: 3000 ms cast. Improved Ghost Wolf (2/2) takes off 3.0 s. | CONFIRMED. "3 seconds". | old 1e0a8ca39 | S 3a155b380 |
| S6 | Landing shaman[7] "Stormstrike now increases the Nature damage you deal" | Talent text: your next Lightning Bolt, Chain Lightning or Earth Shock on the target deals 20% more for 12 s. | CONFIRMED. Reworded to that. | old 1e0a8ca39 | S 3a155b380 |

Aligned: totems are raid auras within 30 yards; Mana Spring, Healing Stream and Mana Tide are party only. Lava Burst 2.5 s base, 10 s. Chain Lightning 2.0 s / 6 s. Searing Totem 20 yd. Magma Totem 20 s. Reincarnation 40 min with the talent. Imbue levels (Rockbiter 1, Flametongue 10, Windfury 30). Chain Heal 3 targets at half per jump. Totem of Urgency. The three APL loops.

### Warlock

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | destruction:100 puts crit second | Published raid-60 order: spell power, then hit, Fire power, crit, Shadow power. | CONFIRMED. Hit and Fire power now come before crit. | landing d432b3833 | S 3a155b380 |
| K2 | Landing warlock[4]: "Drain Hope" and "Malevolence, which lets DoTs critically strike" | No spell named Drain Hope in 70753. Malevolence is 5% Shadow crit, and DoTs crit without it (owner ruling). | CONFIRMED. New list: Malevolence (5% Shadow crit), Pandemic, Wrack, Decimation, Demonic Pact, Shadow and Flame, Incinerate, Bane of Havoc. | old 1e0a8ca39 | S 3a155b380 |
| K3 | Landing warlock[3]: Incubus "named by four Demonology talents" | Eight talents name it at their top rank. | CONFIRMED. "eight". | old 1e0a8ca39 | S 3a155b380 |

Aligned: Curse of the Elements covers six schools. Life Tap scales with Spirit. Shadow Trance 10 s. Amplify Curse is off the global cooldown. Decimation 2/2. Shadow and Flame. Bane cast times. Aftermath 10% a point. Summon levels. The three APLs against the prose.

### Warrior

Aligned beyond X14 and X15: every count and talent text in both Arms builds. The Arms opener (Charge, then Battle Shout, then Bloodrage) against the APL. Heroic Strike thresholds. Hamstring only with Windfury. The Berserker swap. Fury's Death Wish and Recklessness timing. Whirlwind 10 s. The Protection APL thresholds (Last Stand, Shield Wall, Thunder Clap refresh). The landing rows: Thunder Clap rank 4 is 20% for 22 s (8205), Shield Block, Shield Wall 60% / 12 s / 15 min, Victory Rush, Spearing Strike (an Arms talent), Death Wish 5%.

### Druid

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| D1 | Landing druid[3]: "Cure Poison is gone" | Cure Poison 8946 has a trainer row (class 1024, level 14) and is in trainables. | CONFIRMED. Clause dropped. | old 1e0a8ca39 | S 13497f823 |
| D2 | Landing druid[4]: "Nature's Majesty and Berserk in Feral; Feral Charge gains a separate Cat version and Primal Fury also grants combo points" | Nature's Majesty is a Balance talent. Berserk 417141 and the Cat Feral Charge are acquire 3 (runes, not learnable). No Primal Fury exists; Blood Frenzy grants the combo points. | CONFIRMED. New: "New talents include Genesis and Nature's Majesty in Balance and Shifting Power, Predatory Discipline, Rend and Tear and Natural Instinct in Feral; Blood Frenzy also adds combo points on Cat builder crits." | old 1e0a8ca39 | S 13497f823 |

Aligned: Omen of Clarity and Nature's Grasp are baseline; Tiger's Fury is acquire 3. Every Feral, Balance, bear and Restoration talent count. Poised at 32. Rip 30 energy. Rake 35 with Ferocity. Shifting Power. Leader of the Pack A290 3%. Starfire 3.5 s. Hurricane with no cooldown. Enrage and Frenzied Regeneration. Wild Growth. Rejuvenation rank 4 at 105 of 360 mana. The four APLs against the prose.

### Hunter (outside the hotfix lane's lines)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | survival:80 Aspect of the Beast gives "more melee attack power than Aspect of the Falcon" | Beast rank 4 (1299447) is 110 melee attack power. Falcon (469145) carries Hawk's 120 (25296). Falcon's class mask is 0, so Deadly Aspects never touches it. Falcon's own proc is 5%; Beast's is Deadly Aspects' 10%. | CONFIRMED. New: "Aspect of the Falcon gives 10 more melee attack power at 60, but Deadly Aspects does nothing under Falcon, and under Beast its attack speed proc ..." | landing 1feda0686 | S 13497f823 |

Aligned: every Beast Mastery, Marksmanship and Survival count and text. Rapid Fire hastes melee too (aura 319). Multi-Shot hits 3 targets. Potions share one 2 min category. Command 5%. Ammo and quiver levels. Human hunter is a new pair. No Aimed Shot, Arcane Shot, Serpent Sting or Trueshot Aura line was edited.

### Mage

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| M1 | fire:81, arcane:79, frost:78 "Gems, runes and the Robe of the Archmage share a 1 minute cooldown". The order "Ruby, potion, Rune, then Citrine two minutes after the Ruby" cannot be played. | ItemEffect puts Mana Ruby, Citrine, Jade, Agate and Demonic Rune in category 1153 at 120 s. The Robe has its own 300 s. The tooltips agree. A Rune after the Ruby blocks the Citrine. | CONFIRMED. New: "...a Mana Ruby when you are 1,200 mana down and a Major Mana Potion, then the Mana Citrine or a Demonic Rune two minutes after the Ruby. Mana gems and Demonic Runes share one 2 minute cooldown, and the Robe of the Archmage has its own 5 minute one." | landing e2b3a9a53 | S 13497f823 |
| M2 | frost:110 "From level 22, AoE grinding with ... Cone of Cold" | Cone of Cold 120 is learned at 26. | CONFIRMED (minor). "From level 26". | landing e2b3a9a53 | S 13497f823 |
| M3 | Landing mage[5] "Hot Streak" | No Hot Streak in 70753. The talent is Heating Up (400624). | CONFIRMED | old 1e0a8ca39 | S 13497f823 |
| M4 | Landing mage[4] Frostfire Bolt "named by four talents" | Three talents name it: Missile Barrage, Accelerant, Heating Up. | CONFIRMED | old 1e0a8ca39 | S 13497f823 |

Aligned: every count in the three builds. Scorch 150 mana. Fire Blast 8 s. Pyroblast 6 s. Frostfire Bolt 3 s. Ice Lance 8 s. Ice Barrier 30 s. Arcane Explosion 390. Blizzard 1,400. Evocation 8 min. Human 10% Spirit. Skyborne mages are Alliance only. The openers against the APLs.

## 3. The casters and warlock lanes' build lines

- **wt/r13-casters 11f258619 (Elemental Tidal Focus):** leveling Restoration `053352` decodes to Tidal Focus 5, Mindfulness 3, Natural Grace 3, Totemic Focus 5 and Improved Reincarnation 2, which matches the new prose. The stat paragraph change ("Tidal Focus gives you 5%") holds for both builds. My S1 edit is on a different line of elemental.md, so the two merge cleanly. Restoration's own Improved Healing Wave line is its own build. No contradiction.
- **wt/r13-warlock 49849fa9a (Improved Bane of Agony):** the line says the raid build drops Improved Drains "since a raid rotation casts no drain". The Affliction APL casts Shadow Bolt, Immolate, Corruption, Life Tap, Bane of Agony, Curse of the Elements, Amplify Curse, Siphon Life and Shoot, and no drain. Improved Bane of Agony is 5% at 1 point. No contradiction. The warlock index's "Pandemic ... Drain Soul" line is talent text, not a build claim.

## Before/after

No sim number moves. Every change is prose or landing copy.

## Files to regenerate

- data/builds/1.60.1.70753/classes.json (the copy of data/curated/classes.json: round 12's paladin edit plus this lane's ten landing rows). test_curated.py::test_the_committed_build_matches_the_curated_facts fails until then, as it did at 93c5d2510.
- Coach starters (data/curated/coach/*.md, generated from guide prose): the priest Renew line (priest-discipline.md:25), Scatter Shot in hunter-beast-mastery.md:40 and hunter-marksmanship.md:37, and the changed shaman, mage, Survival and warrior lines.

## Measure at refresh

None.

## For other lanes

- **healers:** three APL notes repeat P1 and P3. `data/curated/apl/priest-discipline.json:1394` has "keeps its full healing power bonus". `priest-discipline.json:351` and `priest-holy.json:150` have "every rank carries the same healing power bonus". Proposed: "every rank from level 20 up carries the same healing power bonus".
- **casters:** the arcane APL note says "rank 8 [Arcane Missiles] is a book from Ahn'Qiraj", but 25345 is a trainer row at 56. The fire APL note's "Fireball's 395" is rank 11, while rank 12 (25306) costs 410. Shadow keeps Vampiric Embrace up per the guide, but the APL never casts it, so the published number leaves out a global every 30 s.
- **ranks:** the warlock APLs cast Immolate 11668 (rank 7), but 25309 (rank 8) is also level 60. The warrior APLs cast Heroic Strike 11567, but 25286 is level 60.
- **auras:** the client shares one 2 minute category (1153) between mana gems and Demonic Rune. The engine gives each its own cooldown, and mage gems have none.
- **physical:** the Protection warrior guide asks for Demoralizing Shout, and the APL has none.
- **trainables:** Lava Burst is both a level-30 trainer row and a talent. Lacerate has a talent node and a trainer row (1328617 at 25), but the druid landing says "learned after level 30".
- **platform:** the client and the tooltip spell the Skyborne racial "Read Leyline" (1259705). race-table.ts:26/65/66, races.json:576/616/746, balance.md:107 and restoration.md:99 say "Read Ley Line". These pages stay unchanged until one rename lands everywhere.

## Still open

- F1: Heart of the Lion at 60 is 280 or 276 attack power (BaseLevel 0 against SpellLevel 1). The in-game tooltip on a level 60 hunter settles it.
- Gift of Nature and Tranquility (restoration.md:56, "does not reach Tranquility"). The op-0 mask includes Tranquility's bit and the op-22 mask does not, and Tranquility heals through a periodic aura. A server heal log settles it. It is already on the round-12 list.
- Protection warrior and bear stat prose (stamina and survival first) against published tank weights that rank expertise and threat first. Is the tank score meant to drive the prose? hybrids and platform to answer.
- Elemental Mastery has ProcCharges 1 against the guide's "for its 15 seconds" (round-12 casters E2).

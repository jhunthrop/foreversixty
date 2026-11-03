# Round 14: guides lane

Lane `guides`. Site branch wt/r14-guides, started from int/r14 (72471fcb3). No engine branch.

Scope:
- web/src/content/guides/** (40 pages)
- the rendered `rotation` notes in data/curated/apl/*.json (28 files)

Client truth is 1.60.1.70753 with the 2026-11-03 push layered. Lookups use the round-13 hotfix lane's layered raw/ (scratchpad r13/hotfix/root/1.60.1.70753/raw), data/builds/1.60.1.70753 (talents, trainables, races, combos, bis) and data/curated/book-ranks.json. The lookup tools are in the scratchpad: r14/guides/spell.py, build.py, order.py and scan.py.

Method:
1. Hunter and the meta-talk sweep were checked by this lane directly.
2. Three read-only helpers fact-checked the other eight classes against the same tools.
3. Every helper claim listed here was re-checked against the client row before any edit. Rejected helper claims are listed in section 4.

Every fix is prose, a rotation note or the guide's statPriority set. No DPS or HPS moves.

Tests (run after each commit, all passing):
- `npx vitest run src/content/guides/ src/lib/sim/rotation-apl.test.ts src/lib/guides/rotation-lists.test.ts src/lib/bis/rotation-view.test.ts` (645 tests)
- `data/tests/test_apl.py`

Commits: S fcb37cbc8, S 93909a2bd, S 52b1bd3ae, S 9733f7eb9, S 311ecf6b0.

## 1. Sim and site meta-talk

New guard: `web/src/content/guides/_prose.test.ts`, rule `simTalk`.
- **Banned:** sim, sims, simmed, simulator, simulation, SIM, Forever Sixty, the site, our site, the ladder, the ranker, tier list, published number(s), raid numbers, we measured, custom rotation, damage figure, rank-choice, the guide, this guide, this line, the line.
- **Where it applies:** guide bodies, titles and descriptions, and every note under `rotation` in the curated APLs.
- **The one listed exception (`SOURCE_LABEL_ALLOWED`):** a frontmatter source label may cite "Forever Sixty" data, for example "Forever Sixty build data, races.json". That is a citation. No page under guides is about the sim tool itself.
- **Verified red, then green:** before the text fixes, the test failed on 5 classes with 13 violations. After them it passes.
- **Not scanned:** the top-level `notes` and `engineering_notes` of an APL file. They are developer prose and nothing renders them (web/scripts/rotation-apl.mjs:7).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | balance.md:107 said of Read Leyline: "The sim does not cast it on its own; a custom rotation can." | Meta-talk. Client: 1259705 is a 0.5 s cast on a 2 min cooldown and triggers Energized 1270842 (aura 110 and aura 88 at 50%, 15 s). Attuned 1259691 adds +100% near a leyline. | CONFIRMED. New: "Read Leyline: 50% more health and mana regeneration for 15 seconds on a 2 minute cooldown, twice that next to a leyline. Press it when your mana starts to run low." | R13 dc86a6e2e | S fcb37cbc8 |
| T2 | restoration.md:99 (druid) had the same Read Leyline line ("add it to a custom rotation to see it") | as T1 | CONFIRMED. The same fact, now "press it every time it is ready on a long fight". | R13 dc86a6e2e | S fcb37cbc8 |
| T3 | shadow.md:81: "The sim keeps Vampiric Embrace up and counts only your damage" | Meta-talk | CONFIRMED. New: "That healing goes to your group, so it never shows on your own damage meter." | R13 dc86a6e2e | S fcb37cbc8 |
| T4 | affliction:96, demonology:109 and destruction:108 said "The table above is what the raid numbers assume." | Meta-talk | CONFIRMED. New: "The table above is the standard raid kit." | old d432b3833 (round 13 read these pages and missed it) | S fcb37cbc8 |
| T5 | Shadow APL note (Dark Sacrifice): "Other races ... skip the line" | Talks about the priority list | CONFIRMED. New: "Only Undead priests have it." | old 8ed6b286e | S fcb37cbc8 |
| T6 | Shadow APL note (Vampiric Embrace): "the guide keeps it up ... which the damage figure does not count ... below the level that takes the talent" | Meta-talk. 15286 is a 30 s cooldown and a 30 s duration, 20%, and a talent. | CONFIRMED. New: "heals your whole group for 20% of every Shadow spell you land; your own damage meter never shows that healing. Skip it in the last 10 seconds, and until you have the talent." | R13 175d4ff96 | S fcb37cbc8 |
| T7 | Fire APL note (Improved Scorch): "so this line matters after a move" | Meta-talk | CONFIRMED. New: "a refresh of its own matters only after a move or a target swap". | old e2b3a9a53 | S fcb37cbc8 |
| T8 | Restoration shaman APL note: "Nature's Swiftness is saved for the emergency line below" | Meta-talk | CONFIRMED. New: "saved for the emergency Healing Wave on the tank". | old fdf6c87f4 | S fcb37cbc8 |
| T9 | Protection warrior APL note (Battle Shout): "so this line idles" | Meta-talk | CONFIRMED. New: "In a raid at 60 another warrior's Battle Shout is the same aura, so you only shout when nobody else has." | old 2bcc92053 | S fcb37cbc8 |
| T10 | Restoration druid APL note: "it stays rank 4, the rank-choice rows" | Developer talk | CONFIRMED. New: "below 60 and without raid buffs, rank 4 keeps your mana going". The guard now also bans "rank-choice". | R13 a7070f05d | S 9733f7eb9 |
| T11 | feral.md:96: "at 60 the BiS list weighs it against helms with real stats" | Meta-talk | CONFIRMED. New: "at 60 most helms with real stats beat it". | old c0f5432b2 | S 9733f7eb9 |
| T12 | prot paladin:113: "The table weighs every stat by how much it lowers the damage you take ... then by threat." The same paragraph named two different stats as the top one. | Meta-talk about the score, plus a contradiction. The published raid-60 order starts expertise, then defense. | CONFIRMED. New: "Expertise and Defense come first ... Block, Dodge and Parry come next ...", with no talk about the table. | old 6e3e1bc53 | S 52b1bd3ae |

Kept on purpose: "the cap shown above", "the table above" and "the list below". They point at a block on the same page, the way a printed guide would, and they say nothing about the tool behind it.

## 2. Hunter against the 11-03 push

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | Class page, Beast Mastery and Marksmanship said Aimed Shot "hits about twice as hard" since the hotfix | The push doubled the **bonus**: 20904 eff0 effect 121, bp 166 → 350 (SpellEffect 697764). The shot is still normalized weapon damage plus that bonus, so the whole shot is nowhere near doubled. | CONFIRMED. New: "The November hotfix doubled its bonus damage, so it now outhits Multi-Shot when the cast fits." The class page and Beast Mastery get the same fact. | R13 daaf8314d | S 93909a2bd |
| H2 | Both rotation notes said "Since the 3 November hotfix it hits about twice as hard as before" | as H1 | CONFIRMED. Same fix. | R13 35534af90 | S 93909a2bd |
| H3 | survival:80: "Aspect of the Falcon gives 10 more melee attack power at 60" | 469145: "the same amount as their highest rank of Aspect of the Hawk". At 60 the trainer rank is 14322 (rank 6, 110 after the push). 25296 (rank 7, 120) is an Ahn'Qiraj book (data/curated/book-ranks.json:54; the round-13 hotfix lane says the same). Aspect of the Beast 1299447 gives 110. So Falcon equals Beast until you have the book. | CONFIRMED. New: "Aspect of the Falcon matches your best rank of Aspect of the Hawk, so at 60 it gives the same melee attack power as Beast (10 more once you have Hawk's Ahn'Qiraj rank), but Deadly Aspects does nothing under Falcon ..." | R13 13497f823 | S 93909a2bd |

Aligned:
- Shots and shared cooldown: Aimed Shot 2 s cast, sharing category 6000 with Multi-Shot (2643). Arcane Shot has its own 6 s.
- Serpent Sting: rank 8 (13555, 100 a tick) is the trainer rank at 60, and the APLs cast it. Rank 9 is a book.
- Trueshot Aura: 20906 bp 100. Effect 65 is the raid area aura, so "for your party and raid" holds.
- Sniper Shot: 1310687/785/786 at 320/450/590, 4 s cast, 15 s cooldown. The "4 second cast" line holds. It sits at tier 6 and is out of reach with 30 points.
- Rapid Fire 5 min (aura 319 hastes melee too). Bestial Wrath 2 min / 18 s / 50%. Tranquilizing Shot at 60. Feign Death at 30. Aspect of the Cheetah at 20. Freezing Trap at 20/40/60. Disengage on a 5 s category.
- Survival: Strider Kick 8 s. Mongoose Bite opens on Defensive State 5302 (5 s), and Expose Prey 2/2 opens it at 10% for 10 s. Counterattack needs a parry. Wing Clip rank 3 is 80 mana for 50. Resourcefulness takes 60%.
- All talent texts and counts in the five hunter builds (decoded), and the Marksmanship raid trade.
- Stat prose against the published raid-60 order: Beast Mastery Intellect ≈ Agility within error, "close" holds. Marksmanship and Survival orders match.
- Race list: combos class 3, Human new. Both Skyborne are hunters.
- Heart of the Lion 280/276 stays pending (owner ruling 5).

## 3. Fact checks, other classes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Retribution:62: Conviction "does not reach Consecration or Exorcism", and on all three paladin pages it applies "while you hold a melee weapon" | 20117: aura 290, rank text "all spells and attacks" (2/4/6%), with no weapon clause. Owner law: Conviction is aura 290 on every table (ledger:152; 70753 shipped it). | CONFIRMED. Now "6% crit on everything you do ... your spells too, Consecration and Exorcism included". Holy and Protection drop the weapon clause. | new-data (70753) | S 52b1bd3ae |
| P2 | Retribution:103 stat prose: "Strength is your main stat after [hit] ... Intellect also counts". The statPriority set also left out Intellect and Spell power, so the pill row hid the #2 stat. | bis/paladin-retribution.json raid 60: hit, then Intellect (1.07 ± 0.08), then Strength (1.00), crit, Agility, spell power, attack power. The pills show only stats in the set (spec.astro:194, rail-stats.ts). | CONFIRMED. The prose follows the order, with no numbers and without "the weights above". Intellect and Spell power are added to the set. | old 6e3e1bc53 / 682e01aa8 | S 52b1bd3ae |
| P3 | The Protection paladin and Protection warrior statPriority sets leave out Block value | Block value is a significant published raid-60 stat for both tanks, so the pill row hid it | CONFIRMED (render gap). Added. | old ef69a9288 | S 52b1bd3ae |
| P4 | Enhancement statPriority set leaves out Spell power | significant in the raid-60 weights | CONFIRMED (minor). Added. | old 682e01aa8 | S 52b1bd3ae |
| P5 | holy:90: "Holy Shock is your only instant heal" | Lay on Hands 633 has no cast time | CONFIRMED (minor). Now "outside Lay on Hands". | old 6e3e1bc53 | S 52b1bd3ae |
| S1 | elemental:52: "The leveling build spends every point in Elemental until 50" | The build decodes to Elemental 30, Thundering Strikes 3, Restoration 18. The bare bands show Elemental full by 40, Thundering Strikes 1 at 40 and 3 at 50. | CONFIRMED. New: "fills Elemental to 30 points by 40 and puts three points into Thundering Strikes for 3% crit". The raid line now reads "the last two points of Thundering Strikes". | R13 dc86a6e2e | S 52b1bd3ae |
| S2 | elemental:91/105: the raid crafts "bind when equipped, so you can buy them" | Skycaller's Leather Waistguard 252476, ItemSparse Bonding 1 (binds on pickup). It is on the published raid-60 list. | CONFIRMED. "Most bind when equipped ... a piece that binds on pickup needs the profession yourself." No item is named, because the BiS list is nightly-owned. | old fdf6c87f4 | S 52b1bd3ae |
| S3 | Enhancement APL note: "Strength of Earth Totem ... Strength is your best stat" | Strength is fifth in the raid-60 order (hit, Intellect, crit, Agility, Strength) | CONFIRMED. New: "it is Strength for every melee player near you". | old fdf6c87f4 | S 52b1bd3ae |
| G1 | mage index:38: Eureka! "lines up with Arcane Power or Combustion every 2 minutes" | Eureka 1259817 has a 120 s cooldown. Arcane Power 12042 is 180 s. Combustion 11129 is category 180 s. | CONFIRMED. "Arcane Power and Combustion come back every 3, so pair Eureka! with them whenever both are ready." | old e2b3a9a53 | S 9733f7eb9 |
| G2 | Fire APL note: "Fireball's 395 for 3.4 seconds" | 10151 (rank 11, the top trainer rank, 395 mana) is a 3500 ms cast | CONFIRMED (minor). Now "Fireball's 395 and its long cast". | old e2b3a9a53 | S 9733f7eb9 |
| G3 | shadow:73 opener: "Shadow Word: Pain, Devouring Plague, then Mind Blast" | The APL now casts Vampiric Embrace (15286) between SW:P and Devouring Plague | CONFIRMED. Vampiric Embrace is added to the opener. | R13 175d4ff96 (the APL change left the prose behind) | S 9733f7eb9 |
| G4 | affliction:62: the raid build "moves the point from Improved Drains" | Leveling has Improved Drains 3/3, Improved Health Funnel 2, Improved Voidwalker 2. The raid build has none of these and takes Improved Bane of Agony 1/2. | CONFIRMED. Now names all three points and the two Demonology talents dropped. | R13 dc86a6e2e | S 9733f7eb9 |
| G5 | affliction:68 and demonology:81: Curse of the Elements "goes on first" | Both APLs cast Corruption (11672) before Curse of the Elements (1311680) | CONFIRMED. Both pages now put it after Corruption. | old d432b3833 | S 9733f7eb9 |
| G6 | Affliction APL note: Agony beats Doom "because its ticks crit with Malevolence and Pandemic" | Doom (603) is Shadow too, and Pandemic lists Bane of Doom. The leveling build has no Malevolence. Amplify Curse: "your next ... Bane of Agony by 50%". | CONFIRMED. New: "because it starts ticking at once and Amplify Curse adds 50% to it". | old d432b3833 | S 9733f7eb9 |
| G7 | demonology:67: Lash of Pain crits "more with Malevolence" | Neither Demonology build takes Malevolence | CONFIRMED (minor). Clause dropped. | old d432b3833 | S 9733f7eb9 |
| G8 | demonology:101: "Hit is close behind" spell power, with Shadow power after it | Raid-60: spell power, then Shadow power, then hit | CONFIRMED (minor). Reordered, no numbers. | old d432b3833 | S 9733f7eb9 |
| W1 | fury:81: Charging in Battle Stance and swapping back "loses everything above 10 rage" | The Fury build has Improved Tactical Mastery 5/5. TM 1310185 keeps 10 and ITM adds 15, so a swap keeps 25, above Charge's 15. The real reason is that Charge (11578) "cannot be used in combat", and the pre-pull Bloodrage has put you in combat. | CONFIRMED. New: "the Bloodrage you pressed before the pull has already put you in combat, so Charge is gone, and a stance dance only costs you globals." | R12 70aff3eaa (the build gained ITM; the sentence is from 4b1864450) | S 9733f7eb9 |
| R1 | subtlety:51/85: the build line leaves out Cold Blood, and "You have no burst cooldown of your own on a boss" | The build decodes to Lethality 1, Cold Blood 1, Improved Poisons 4 in Assassination. 14177 has a 3 min cooldown. | CONFIRMED. The build line names them, and the burst line now says save Cold Blood for a four or five point Eviscerate. | old 53f1006c3 | S 9733f7eb9 |
| R2 | combat:52: the leveling build differs from raid "in place of Lethality and Relentless Strikes" only | Leveling is 14/31/6 with Camouflage 5 and Master of Deception 1. Raid has Aggression 3 and Opportunity 2. | CONFIRMED (minor). Swap added. | old 53f1006c3 | S 9733f7eb9 |
| R3 | assassination:79: "then 4 to 6" combo points | Combo points stop at 5 (every finisher row lists 1-5) | CONFIRMED (minor). "then to 4 or 5". | old 53f1006c3 | S 9733f7eb9 |
| R4 | subtlety:58/69 treated Premeditation as a pre-pull opener only | 14183's description and the Wowhead Forever tooltip (nether.wowhead.com/forever/tooltip/spell/14183, read 2026-11-03) carry no Stealth requirement: 2 combo points, 20 s, 2 min cooldown. The physical lane's engine now casts it in combat (E e01d5596a). | CONFIRMED. The talent line and the opener paragraph now say to press it again every 2 minutes in the fight at three combo points or fewer. | old 53f1006c3 | S (this commit) |
| D1 | feral-bear:87 and the bear APL note: Barkskin "20% less damage" | 22812 eff1: aura 87 bp -20, misc 1 (Physical). Text: "Physical damage taken is reduced by 20%". | CONFIRMED. "20% off physical damage". | old c0f5432b2 | S 9733f7eb9 |
| D2 | druid index:30 and feral:70: Furor "only keeps your bar when you leave Cat Form" | Furor 5/5: "restores 100% of the Energy you had when last in Cat Form, plus 10 Energy per second spent outside Bear Form ... up to 100" | CONFIRMED (minor). Both now say it gives back the bar plus 10 a second out of form. | old c0f5432b2 | S 9733f7eb9 |
| D3 | Feral APL note: Shred "costs 14 to 20 more energy" than Claw until Shredding Attacks is full | Shred 60 − 6 a point, Claw 40 with Ferocity: at 2/3 the gap is 8 | CONFIRMED (minor). "8 to 20". | R12 7ea0ba759 | S 9733f7eb9 |
| D4 | feral-bear:75 pull: "Then Lacerate, Primal Bite" | The APL ranks Primal Bite before Lacerate, as does the page's own priority at :79 | CONFIRMED (minor). Swapped. | old c0f5432b2 | S 9733f7eb9 |

Aligned, checked spell by spell. The helpers' full lists are condensed here.
- **Paladin:** every talent text and count in the five builds; Holy Shield; Judgement; Hammer of Wrath; Holy Strike 10 s; Repentance; Righteous Fury; Templar's Bulwark; Swift Judgement; Light's Vigil; Flash of Light rank costs; Holy Light rank 9 is a book (book-ranks.json, so holy:84 holds); the Retribution twist gates; the Protection APL order. Round-14 ruling 1 (Templar's Bulwark at 20) touches no guide line.
- **Shaman:** every count; Lava Burst; Call of Flame; Lightning Bolt ranks 9/10; Chain Heal ranks; Water Shield; Mana Tide; Riptide; imbue and totem levels; Stormstrike; Maelstrom Weapon; Rage of the Farseer; races. Restoration's stat paragraph claims no strict order, so it was left.
- **Mage:** every count and talent text in the nine builds; Arcane Blast; Missile Barrage; Arcane Power; Heating Up; Ignite; Scorch; Flamestrike; Frostbolt; Ice Lance (a talent, ranks from 20); Fingers of Frost; Winter's Chill; Blizzard; Ice Block; Evocation; Mage Armor; Skyborne Alliance-only.
- **Priest:** Penance; Divine Aegis; Binding Heal; Lightwell; Shadowform; Shadow Weaving; Devouring Plague; SW:Death; Eureka masks; Dark Sacrifice; Divine Grace; Greater Heal rank 5 is a book.
- **Warlock:** Curse of the Elements; Pandemic; Wrack; Amplify Curse; Improved Corruption; Ruin 4/5; Decimation; Soul Link; Master Summoner; Demonic Pact (tier 6; the raid build takes it, so "at 60" is about the respec); Destruction's raid swap; Incinerate outside Eureka.
- **Druid:** Omen of Clarity; Nature's Grasp; Poised 32; Shifting Power; Eclipse; Moonkin aura; Innervate; Elune's Light; Wolfshead Helm; Lacerate and Primal Bite; Enrage; Frenzied Regeneration; Swiftmend 15 s; Wild Growth; Rejuvenation and Healing Touch top ranks at 60 are books (25299, 25297), so the rank-10 casts are right.
- **Rogue:** every decode; Mutilate; Seal Fate; Cold Blood; Adrenaline Rush; Blade Flurry; Hemorrhage; Ghostly Strike; Premeditation; Preparation; Thistle Tea; poison levels.
- **Warrior:** both Arms decodes and the raid move; Improved Slam; stance masks; Bloodrage; Death Wish; Bloodthirst; Whirlwind at 36; Protection talent texts; Shield Block; Last Stand and Shield Wall thresholds.

## 4. Helper claims rejected

| claim | why rejected |
|---|---|
| The frontmatter `statPriority` order disagrees with the raid-60 order on most spec pages (helpers listed 18 rows) | The page never renders that order. The pills and rail rows sort by the published weights (`web/src/pages/guides/[class]/[spec].astro:194`, `web/src/lib/guides/rail-stats.ts:7-10`). The frontmatter is only the set of stats shown, so only set gaps are bugs (P2-P4). Reordering would chase a nightly-owned order. |
| Starfire, Ferocious Bite, Eviscerate, Backstab and Battle Shout APLs cast below the "top trainer rank" | The higher ranks (25298, 31018, 31016, 25300, 25289) are Ahn'Qiraj books in data/curated/book-ranks.json. Aligned. |
| Round 13's "For other lanes": the Arcane Missiles rank 8 book note and Fireball's 395 are wrong | Both notes are right. ItemSparse 21280 teaches Arcane Missiles VIII. Fireball rank 12 is a book (21279), so rank 11's 395 is the top trainer rank. Round 13's two lines there were wrong. |
| Holy Power "175% instead of 150%" is open | A 50% larger crit bonus on a 50% bonus is 175%. Aligned. |
| Holy Light and Healing Wave rank 9 might be trainer rows | They are books in book-ranks.json. Aligned. |

## 5. Follow-up after the healers and tanks lanes merged (int/r14 c2cd3dc5b)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | Priest guides and notes against the healers lane's new ranks. holy.md said Renew "at a low rank ... rank 3", Heal "rank 4" as the raid filler, and the tank's Flash Heal "the low ranks". discipline.md and both Renew notes said "a low rank". | The curated APLs now cast Renew 139 (rank 1), Heal 2054 (rank 1, learned at 16, 3.0 s, 2.5 with Divine Fury) in a raid with rank 4 on the bare 60 and leveling rows (rank-choices.json priest-holy), and Flash Heal 10917 (rank 7, top trainer rank) on the tank. Wording from healers.md "For other lanes". | CONFIRMED (prose behind the rotation) | the healers lane's round-14 rank change | S (this commit) |
| F2 | The new Demoralizing Shout and Demoralizing Roar rotation notes said "204 attack power off it at 60 takes about a fifth of its melee damage away" | 204 is the client value: 11556 bp -196 and 9898 bp -193, both -1.4 a level, so 204 at 60 for each. "About a fifth" rests on the modeled boss attack power (268, ruling 4), so the note no longer states it. | CONFIRMED (minor). New: "204 less melee attack power at 60 takes a real slice off every swing. One is enough, since it does not stack with a druid's Demoralizing Roar" (or a warrior's Shout). | the tanks lane's round-14 lines | S (this commit) |
| F3 | The bear guide's priority list had no Demoralizing Roar, and the Protection warrior priority paragraph had no Demoralizing Shout. The warrior's own paragraph said "refresh it on a free global". | The tanks lane's APLs keep the roar after Faerie Fire and the shout after Thunder Clap, each refreshed in its last 3 seconds (Shout lasts 45 s, Roar 30 s) | CONFIRMED. The roar is item 4 of the bear priority; the shout follows Thunder Clap; both refresh in their last 3 seconds. | the tanks lane's round-14 lines | S (this commit) |

The tanks lane's rewritten top-level APL `notes` are developer prose that no page renders, so they are left as written. Their measurement figures are fine there.

## Before/after

No sim number moves. Every change is prose, a rotation note, a reason string or a guide stat set.

## Files to regenerate

- `make apl-sync`: sim/request/apl copies of the 12 curated files touched: mage-fire, mage-frost (reason only), priest-shadow, shaman-restoration, shaman-enhancement, warrior-protection, hunter-beast-mastery, hunter-marksmanship, warlock-affliction, druid-feral-bear, druid-feral, druid-restoration. Only text moves.
- Fork ui/*/apls/forever_*.apl.json carry the same notes (for example forever_fire and forever_shadow). They do not reach a page. Sync them with the next fork APL commit of each class lane.
- addon-data.json: the guide priority card reads its notes from here.
- Coach starters (data/curated/coach/*.md, built from guide prose and notes): at least priest-shadow (the old Vampiric Embrace note), hunter-beast-mastery and hunter-marksmanship (Aimed Shot), and the pages changed in S 52b1bd3ae and S 9733f7eb9.

## Measure at refresh

None.

## For other lanes

- **physical (rogue):** Subtlety's build takes Cold Blood, but its APL has no Cold Blood line. So either `autocastOtherCooldowns` fires it whenever it is ready, or nothing casts it at all. The guide now says to save it for a four or five point Eviscerate. Please check which one it is.
- **hybrids or tanks:** retribution:95 says rank 4 Consecration goes further without raid buffs. The APL always casts rank 5 (20924), which is right for the raid fight. No change asked.
- **tanks:** the Protection paladin and Protection warrior pill sets now include Block value. Round 13's open question (tank stat prose against tank weights) still stands. Ruling 4 (boss attack power) will move both orders, so the prose names only the leading stats.
- **platform:** the coach catalogue (data/curated/coach/hunter-*.json) points at Aspect of the Hawk 25296, the Ahn'Qiraj rank. Check that the coach does not tell a fresh 60 to train it.

## Still open

- Swiftmend: the guides say it keeps the HoT. The 18562 tooltip says "the consumed spell", and the client tables cannot settle it. An in-game check settles it.
- Faerie Fire free in Cat and Bear Form on a 6 s cooldown: 9907 shows 115 mana and no cooldown. The form version rests on one source.
- Cat and bear swing timers (1.0 s and 2.5 s): there is no SpellShapeshiftForm table in the raw set.
- Elemental Mastery has ProcCharges 1 against "for its 15 seconds" (carried from rounds 12 and 13).
- Portal of Summoning's five shards: there is no SpellReagents table.
- Hunter spell power weight is significant at raid 60 but not in the guides' stat sets. It stays out until ruling 5 settles whether Arcane Shot takes spell power.
- WCL-dependent checks are blocked (gcloud login expired).

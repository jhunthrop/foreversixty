# Round 22, lane warrior (Arms, Fury, Protection)

Site `wt/r22-warrior` on int/r22 (21bd3af0c). Engine `lane/r22-warrior` on fork int/r22 (170a59c39). Client 1.60.1.70790, which is still the newest Forever build on wago (`/api/builds`, product wowf, read 5 Nov 07:15 UTC).

**Result: 0 confirmed warrior bugs. No engine, APL, rank-row, build or guide change.** The seed (Revenge rank 6, plus the coordinator's Battle Shout rank 7 and Heroic Strike rank 9) is the owner's pending "level-60 trainer book ranks" check. The client cannot tell a trainer rank from a book rank for these three spells, and the engine, the APLs and the cards all agree with the client's book reading. Round 21's open items are still open: nothing new has been published, and no Forever log exists yet.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | **Guides W22-1: Revenge rank 6 (25288) at 60.** `trainables/warrior.json` lists 25288 at 60 with `source: skill_line_ability`, because the client has a SkillLineAbility row for it (13218, skill line 257, AcquireMethod 0, supersedes 11601). The same row with the same values is in **Classic Era 1.15.9.70576** (13218, `257,25288,1,1,11601,0,...`). In Era, rank 6 is learned from the Ahn'Qiraj book, not from a trainer. So the row does not make it a trainer rank. In Forever the book still exists and still teaches it. Wowhead's Forever tooltip for item 21299, "Manual of Revenge VI", reads "Requires Level 60. Use: Teaches Revenge (Rank 6)", and the client's ItemSparse 70790 has it. `spellranks.json` marks 25288 `book: true`. The engine (`revenge.go:34`) caps Revenge at rank 5 while `core.IncludeAQ` is off. The Protection APL and card cast 11601 rank 5. Engine, APL and card agree with the client's book reading. Whether a Forever trainer actually sells rank 6 at 60 can be settled only in game (open the trainer at 60) or by a log that shows 25288. That check is the owner's pending item "level-60 trainer book ranks" (ledger, rounds 13 and 14). It is recorded, not reopened. If rank 6 does train, Revenge gains its 137.7 to 168.3 base against 109.07 to 132.93 (about +32 a hit before Improved Revenge). | wago SkillLineAbility 1.60.1.70790 and 1.15.9.70576; ItemSparse 70790 (21297, 21298, 21299); https://nether.wowhead.com/forever/tooltip/item/21299 and /spell/25288 and /spell/11601 (read 5 Nov 07:05 UTC); `spellranks.json`; `revenge.go:34` | Aligned with the client; trainer question **pending owner** | - | - |
| S2 | **Coordinator seed: Battle Shout rank 7 (25289) and Heroic Strike rank 9 (25286).** Same shape as S1. Their SkillLineAbility rows (13219, 13217) are byte-identical in Era and Forever. The Forever item tooltips for 21298 "Manual of Battle Shout VII" and 21297 "Manual of Heroic Strike IX" ("Teaches Heroic Strike (Rank 9)") are live, and `spellranks.json` marks both `book: true`. The engine caps them under `IncludeAQ` (`shouts.go:107` rank 6; `heroic_strike_cleave.go:22` rank 8). All three APLs cast Battle Shout 11551 rank 6 and Heroic Strike 11567 rank 8. 27578 (a second rank-6 Battle Shout id in `spellranks.json`) has no learn row and is not a player rank. A web search (5 Nov) found no Forever source that says a trainer teaches any of the three. | as S1; https://nether.wowhead.com/forever/tooltip/item/21297, /spell/25286 | Aligned; **pending owner** (same check) | - | - |
| S3 | **Defiance 7/13/20 against 5/10/15 (round 21 S2).** Re-read 5 Nov 07:10 UTC: Wowhead's Forever tooltip for 12792 still reads "an additional 15%", and the one for 7376 still reads "[Defiance: 107)% / 80%]" (1.8 x 1.15). Wowhead's live data has the stance's 80 but not Defiance's 20. The 70790 trait curve is unchanged, and there is no newer client build. Blizzard's thread has no post after 4 Nov 19:35 UTC. | Wowhead tooltips 12792, 7376, 71; wago `/api/builds`; https://us.forums.blizzard.com/en/wow/t/2360696 (10 posts, read 5 Nov 07:10 UTC) | OPEN (carried; next client build or a threat capture) | - | - |
| S4 | **Threat constants with no client column: Revenge (2.25x damage + 2.25 x 2 x level), Shield Slam, Thunder Clap (2.5x).** The client rows carry no threat effect: 11601 has only effect 2, 23925 effects 38 and 2, 11581 effects 2 and 319. Sunder (11597 e1, effect 63, 206) and Demoralizing Shout (11556 e1, effect 63, 43) are the only Protection spells with one, and the engine uses both values. I re-read forever-bugs at 07:00 UTC 5 Nov: nothing has been filed since #448 (4 Nov), and the threat issues (#278 and others) have not changed. | wago SpellEffect 70790; `gh search issues` | OPEN (carried) | - | - |
| S5 | **Deep Wounds' attacker multiplier, applied twice (round 21 S4).** It is still in `deep_wounds.go` OnTick ("Double dips on attackers mods", the fork's long-standing vanilla model). Blizzard's 8 Oct note and #234 change only the rollover and the attack power, and say nothing on modifiers. #434 (procs) still has no answer. No log. | code; forever-bugs #234, #434 | OPEN (carried; a 40+ log with Deep Wounds ticks beside a known weapon and Death Wish/Enrage windows) | - | - |
| S6 | **A missed Heroic Strike and Reckless Swings.** Reckless Swings (1329555) has ProcTypeMask 0x4 (melee auto attack) and the default hit mask, so a white swing that misses spends no charge and grants no rage (`reckless_swings.go`, `result.Landed()`). The log-backed Flurry rule turns a missed Heroic Strike into a missed white swing (Tawl's 70450 log, `talents.go:512`). A missed white swing does not touch Reckless Swings either way, so both readings give the same answer: no charge. A landed Heroic Strike is not a white swing for procs (round 19 ruling, `countsAsWhiteSwing`). | client SpellAuraOptions 1329555 (as round 18 A2); code | Aligned (closes the round 21 open item) | - | - |

## Blizzard's Forever notes, every warrior line (24 Sep to 4 Nov)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | I re-read the thread on 5 Nov at 07:10 UTC. It has 10 posts and the last is 4 Nov 19:35 UTC (no warrior line). The warrior lines are the ones round 21 checked (its N1 to N7): 24 Sep Sunder threat and Slam 18 s / Improved Slam 1.5/3; 1 Oct the Fury pass (Bloodthirst 45% AP, Raging Blows, Gore Drinker, Dual Wield Specialization off-hand only); 8 Oct Deep Wounds without AP and rolling, rage from damage; 16 Oct Slam pushback, Overpower against low-level targets, Weaponmaster not from Shield Bash/Slam; 3 Nov Defensive Stance 80%, Slam 15 s, Improved Tactical Mastery 5/10/15, Two-Handed Specialization 2/4/6, Bloodthrill 10/20, Improved Slam 2.5/5, Booming Voice on Piercing Howl, Reckless Swings, Flurry 10/20/30 for 3 swings, Improved Whirlwind, Bloodthirst behind Death Wish, Defiance 7/13/20 (S3). Nothing new; the engine matches every line except Defiance (S3). | thread JSON | Aligned (S3 open) | - | - |

## Arms

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Raid 60 breakdown read spell by spell** (horde, frozen BiS, guide raid build, 3,000 iterations). Mortal Strike 23.8 casts (6 s cooldown); Slam 10.5 (15 s less Improved Slam's 5); Spearing Strike 6.4 (20 s, Battle Stance only, 75.6% of the fight in Battle Stance); Reckless Swings 11.4 (15 s); Overpower crit 77.4% (Improved Overpower 50); one Charge for 15 rage (the raid build has no Improved Charge); Berserker Stance once, at the execute swap; Recklessness once. Yellow miss 3.4% and white glance 40% against a level 63 boss. No anomaly. | `$R/warrior/out/arms.txt` | Aligned | - | - |
| A2 | **Build points.** The raid build spends nothing dead. Every point is cast, raises a cast's damage or rage, or is a tier or prerequisite gate (Sweeping Strikes for Mortal Strike, Enrage 5 for the Fury tier 3 gate, as round 17 A4). | `talents/warrior.json` decode | Aligned | - | - |

## Fury

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Raid 60 breakdown.** Bloodthirst 22.1 casts (6 s, behind Execute below 20%); Whirlwind 16.3 (8 s with Improved Whirlwind 2) and its off-hand strike 462891 on every cast; Berserker Rage 5.8 (30 s); Death Wish 1.1 and Recklessness 1.0 (execute phase, longer fights sometimes fit a pull Death Wish); Flurry uptime 86.5%; yellow miss 0% and dual-wield white miss 18.5% (9.5% hit against 28%). No anomaly. | `$R/warrior/out/fury.txt` | Aligned | - | - |
| F2 | **Arms points in the Fury build** (Improved Heroic Strike 3, Improved Rend 3, Improved Charge 2, Improved Tactical Mastery 3, Anger Management 1, Deep Wounds 3, Impale 2). Improved Rend and Improved Charge do nothing for Fury. They are needed as gates: tier 1 needs 5 points (Improved Heroic Strike and Improved Rend give 6), tier 2 needs 10, Deep Wounds needs Improved Rend 3, and Anger Management needs Improved Tactical Mastery 3. Every point that could stand in for them (Deflection, Improved Overpower, Iron Will) also does nothing for Fury. So there is no legal swap to a point that does something, and the rule's CONFIRMED-BUG test is not met. Enrage 2 is the Fury tier 3 gate (Cruelty 5, Unbridled Wrath 5, Furious Precision 3, Enrage 2 = 15). | `talents/warrior.json` tiers and prerequisites | Aligned | - | - |

## Protection

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **Raid 60 tank breakdown.** Revenge 30.1 casts (5 s; it lights on 77% of the fight); Shield Slam 30.2 (6 s); Shield Block 29.4 (80% uptime); Thunder Clap 7.3 and Demoralizing Shout 5.2 (30 s debuffs kept up); Sunder 47.5. Threat per cast against the formulas: Sunder about 426 a cast = (206 + 5% AP) x 2.07 x about 72% landed from the front (miss 8%, dodge 6.5%, parry); Demoralizing Shout about 83 = 43 x 2.07 less misses. The round-21 stance change (80% x Defiance 15%) shows in every row. | `$R/warrior/out/prot.txt` | Aligned | - | - |
| P2 | **Focused Rage's reach, against the client mask.** 29787 e0 (aura 107, -3 rage) has the class mask [1852722926, 1057316, 1, 0]. Each spell's SpellClassOptions mask overlaps it for Revenge, Shield Slam, Thunder Clap, Sunder Armor, Heroic Strike, Cleave, Demoralizing Shout, Hamstring, Rend, Overpower, Mortal Strike, Execute, Slam, Spearing Strike, Bloodthirst, Whirlwind, Shield Bash, Pummel and Death Wish. It does not reach Battle Shout, Shield Block, Bloodrage, Berserker Rage, Recklessness, Charge, Reckless Swings, Last Stand, Shield Wall or the stances. The engine's `SpellFlagOffensive` set covers every spell the Protection card casts in the first list, and none of the second. Death Wish is in the mask but not flagged; Protection never casts it. | wago SpellEffect / SpellClassOptions 70790; `talents.go:319` | Aligned | - | - |
| P3 | Bastion (16538, aura 79, misc 1 = physical school, 10 at five points) matches `talents.go:310` (physical only, 2% a point). Improved Revenge (12797, aura 108, mask 1024, 60) and Improved Thunder Clap (12287, -60 tenths = 6 rage) match. Improved Shield Wall stays with the owner (WARRIOR-13). | wago SpellEffect 70790 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | forever-bugs issues no round had cited: #126 (the off-hand swing timer after a weapon swap or a Slam swing reset), #420 (Sweeping Strikes taking armour twice), #65 (block value tooltip). None reaches a published single-target rotation: no rotation swaps weapons, Fury never casts Slam, Arms' Slam runs with Improved Slam, and there is no second target for Sweeping Strikes. Also: #420 matters to an Arms player leveling on packs; #126 to a Fury warrior swapping weapons. | gh issues 126, 420, 65 (read 5 Nov 07:00 UTC) | Aligned (out of the sim's scope) | - | - |

## Before/after

No change in this lane, so the raid-60 numbers are `$R/before.tsv`. Horde seed 1, 3,000 iterations on this branch: Arms 854.7, Fury 910.3, Protection 1490.3 threat/s, all inside the before figures' error.

## Files to regenerate

None.

## Measure at refresh

- Carried from round 21: the Protection BiS on the 80% stance (picks should not move), and the warrior BiS after the CI re-rank.

## For other lanes

- **Platform / data:** `trainables/warrior.json` lists the three book ranks (25286, 25288, 25289) beside trainer ranks with `source: skill_line_ability`. The Era client has the same rows, so a guide or card reader cannot tell them apart. `spellranks.json` marks them `book: true`, but trainables does not. If the owner's in-game check finds them at the trainer, flip `core.IncludeAQ` for these spells (or per spell). If it does not, mark them `book` in trainables too.

## For the guide lane

- None. W22-1 is answered above (book rank, owner check pending); no guide line names Revenge rank 6, Battle Shout rank 7 or Heroic Strike rank 9. R22-1 (the Shield Slam note) is already "a hard hit with a very high amount of threat" on int/r22.

## Guide text not done

- None (no change).

## Still open

- S1/S2 trainer book ranks (owner, in game: a level-60 warrior's trainer list, or a log with 25288, 25286 or 25289).
- S3 Defiance 7/13/20; S4 Revenge / Shield Slam / Thunder Clap threat; S5 Deep Wounds' double multiplier; WARRIOR-13 (owner).
- **Logs, last looked 5 Nov 07:19 UTC:** api.foreversixty.gg `/v1/reports/recent` lists only the 20 Oct report. foreverlogs.gg `/api/reports/public` now answers 403 "API access not allowed". Warcraft Logs v1 `/v1/zones` on www/classic/fresh/vanilla/sod (66/57/21/29/12 zones) has no Forever zone. A log with Revenge, Shield Slam and Thunder Clap threat events, or Deep Wounds ticks, would settle S4 or S5.

## Sources

- wago.tools: SkillLineAbility 1.60.1.70790 (cached `scratchpad/b70790/70790`) and 1.15.9.70576 (`$R/warrior/era/`); SpellEffect, SpellClassOptions, SpellName, ItemSparse 70790; `/api/builds` (read 5 Nov 07:15 UTC).
- Wowhead Forever tooltips (5 Nov 07:05 to 07:10 UTC): item 21297, 21299; spell 25286, 25288, 11601, 12792, 7376, 71, 23925, 11581, 12834.
- Blizzard, WoW Forever patch notes: https://us.forums.blizzard.com/en/wow/t/2360696 (all 10 posts, read 5 Nov 07:10 UTC).
- forever-bugs: newest issues #409 to #448, warrior search, #126, #420, #65 (5 Nov 07:00 UTC).
- Web search, 5 Nov: "WoW Classic Forever trainer rank 6 Revenge / Heroic Strike rank 9 / Manual of Ahn'Qiraj books trainer"; only item-database pages listing the Manual of Heroic Strike IX for Forever and Era, no trainer source.

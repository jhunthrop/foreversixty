# foreverlogs: warriors and hunters, reports 2815-2853 (2026-11-12)

Data: 2815-2849 from scratchpad/fl/data/<rid>.json (physical lane's dump: warriors, rogues, hunters, druids + pets; streams damage, casts, avoidance, resources, buffs).
2850 and 2853 fetched fresh into logs/data/ (warriors, hunters, all pets). No report newer than 2853 on /api/reports/public at fetch time.
Scripts: logs/py (profile, hunter, hunter2, hunter3, pets, warrior, abs). Raw outputs: profiles.json, hunter_raw.json, hunter2_raw.json, pets.json, warrior_raw.json, abs_gaps.json, *.txt.

Log quirk that matters: the casts stream holds both cast START (target null) and SUCCESS (target set). Cast times = start->success.

## Specs found
- No Mortal Strike anywhere. No Lacerate, Strider Kick, Bestial Wrath in any report.
- 2850 / 2853: Dread (2850, lvl>=28, stance-dancing hybrid, 215 swings), Sativa/Adrian/Lenara (2853 tanks), Bazooka (2853, >=20, 25 swings). Hunters all ranged and low (8-16); Mark (2853, >=16, cat Barnabus). No hawk.
- Fury fragments, 2834 (ESW, 8 Nov): Thrump >=42 (BT 2, WW 3, Death Wish, Reckless Swings, Flurry, 17 swings), Deemsters >=40 (BT 3, WW 4, 24 swings). Arms-ish: Big >=36 (WW 6, Sweeping 3, OP 5, Exec 5, 44 swings), Limp >=32 (Sweeping, Deep Wound), Shmeegle >=30, Docked >=32, Pimpin >=24.
- Melee-leaning hunters (none Survival): Spykids 2849 lvl 8-9 (94 Raptor, 255 melee vs 221 Auto Shot), Moe 2848 lvl 20-23 (Monkey 100%, 64 Raptor, 20 Mongoose, 65 Aimed), Ivy 2827 lvl 16-17.

## Hunter measurements: see final report; numbers in hunter2.txt, hunter3.txt, pets.json, abs_gaps.json
## Warrior measurements: warrior.txt

## Findings (value, n, players; Classic expectation)
- Serpent Sting: tick 3.000 s (n=5701, 40 hunters); 5 ticks per full application (411 of 449, 23 hunters; 0 = resist/immune 28, 6 ticks 3); first tick 3.5 s after cast success (3 s after impact); early refresh resets the tick clock. Classic: 15 s / 5 ticks.
- Arcane Shot cd: min 5.943 s, p10 6.003 (n=1900, 68). Classic 6 s.
- Multi-Shot: hasted 0.5 s windup (start->success 0.459 med, 22 hunters, n=1233; 0.357 under Quick Shots; 0.321 under Rapid Fire). Cooldown: success gaps min 6.32 s, p10 6.46 (n=818, 23). Client 70888 cd 6 s. Classic 10 s.
- Aimed Shot: cast 1.82 s (n=412, 21) out of buffs = 2.0 / quiver haste (Dusk, no quiver: 2.003); 1.38-1.40 under QS; 1.29-1.30 under RF; 1.005 under RF+QS. Cd: min success gap 7.30 s = 6 s + cast.
- Auto Shot windup (start->success): 0.458 s out, 0.352 QS, 0.324 RF (83 hunters) = 0.5 s / ranged haste.
- Auto Shot interval ratio: QS 0.762-0.777 (12 hunters; Pun 0.852 outlier); RF 0.708-0.717 (5); RF+QS 0.549-0.553 (2) = multiplicative.
- Rapid Fire duration 15.0 s (n=24, 7). Cd: gaps >=305 s for 6 hunters; one 184 s gap (Zae 2818, BFD 13:02:45 -> 13:05:49).
- Quick Shots buff 12.0 s (n=278, 21).
- Raptor Strike cd min 5.94 s (n=259, 43). Replaces a white swing: 3 of 562 Raptors have a white within 100 ms; prev/next white sit one swing away.
- Mongoose Bite cd min 4.994 s (n=17). Trigger (hunter's own dodge) not in the data (mob-source events not fetched).
- Wing Clip: GCD only (min 1.48 s, n=36).
- Traps: Immolation Trap Effect 3.0 s ticks, 5 ticks over 12 s span (n=112, 12); Explosive Trap Effect initial + 2.0 s ticks, up to 11 events over 20 s (1 hunter). Trap cast gaps min 31.2 s (12 hunters) = client 70888 30 s category cd (Classic 15 s).
- Hunter's Mark: 0 to 99 casts per row; mark on 0-82% of encounters.
- Aspects: Hawk 55-100% for most; Moe Monkey 100%; Grok Cheetah 58%.
- Pets: Bite cd min 10.02-10.6 s (13 pets); Claw GCD 1.5 s, median 1.63; Growl 5 s; Charge ~25 s. Swing interval per pet: 2.00 (boar, wolf, crab, 7 pets), 1.40 (5), 1.20 (3 cats), 1.30, 1.37, 2.45.
- Hawk (pet named Hawk/Faucon), all before 11 Nov: Owen 2823 (lvl>=28, 951 hits, med 10, swing 2.2 s), Draaxil 2833 (>=34, 113, med 12, 2.235 s), Send 2834 (>=28, 68, med 10, 1.76 s), Dusk 2847 (>=25, 207, med 9, 2.50 s). None after 11 Nov.

## Warrior
- Unbridled Wrath: 1 rage each, logged 12-25 ms before a white hit; white only (Dread: 106 of 108 on whites, 0 on ~80 yellow melee hits). Rate per landed white: Dread 0.50 (n=215), Doublewide 0.118 (558), Docked 0.60 (35), Shmeegle 0.49 (55), Deemsters 0.67 (24), Thrump 0.71 (17).
- Overpower: 64 of 65 within 5 s of the warrior's attack being dodged (max 4.96 s); 1 with no dodge and no Rend; no Bloodthrill evidence. Cd min 4.996.
- Rend: tick 3.0 s (n=1788); rank 2 12 s/4 ticks (30), rank 3 15 s/5 (114), rank 4 18 s/6 (3).
- Deep Wound (412613): tick 3.0 s, 4 ticks over 9 s span (12 s), first tick 3.05 s after crit (Garry 44 runs, Limp 2).
- Slam (1240193): cast 1.50 s (Garry n=13); cast gaps min 23 s (client cd 15 s).
- Bloodthirst gaps 6.6, 8.1, 10.7 s (n=3). WW: no 2 casts in one fight. Flurry swing medians (dual wield): Thrump 0.61 in / 0.82 out (7/17), Deemsters 0.80/0.95 (13/19).
- HS casts per landed white 0.02-0.29.
- Victory Rush 30 s, Spearing Strike 20 s, Shield Bash 12 s, Bloodrage 60 s, Charge 15 s.

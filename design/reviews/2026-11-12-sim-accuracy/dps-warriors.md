# Sim accuracy 2026-11-12: Arms and Fury warriors at 50 and 60

Lane `sim-dpswar`, brief `design/reviews/sim-accuracy/brief.md`. Scope: Fury and Arms at 50 and 60, bare and the raid 60 presets, the DPS warrior's first dedicated engine-against-client pass at 60: every spell, talent, proc and set bonus the two raid rotations and builds touch, against 1.60.1.70790 plus the hotfix push layer, the live 1.60.1.70888 rows and Blizzard's Forever posts to 12 November; then the engine against every readable log at 45 and above.

- Site branch `lane/sim-dpswar` from origin/main 62e835b00.
- Fork branch `lane/sim-dpswar` from master 4c595ee6b, the pin at the start.
- Client on the site: 1.60.1.70790 with its push layer. Live client: 1.60.1.70888, hotfix-layered rows (wago.tools `useHotfixes=1`, the 12 November fetch).

**Totals: 1 confirmed bug, old, fixed: Deep Wounds' ticks crit in the engine (the owner's bleed rule) where the client flags Forever's Deep Wound (412613) Can't Crit and 146 logged ticks hold no crit. Arms raid 60 -2.6 to -2.8%, Arms bare 60 -3.3 to -3.7%, Fury raid 60 -2.1 to -2.2%, Fury bare 60 -2.8 to -3.0%; Protection threat moves a little. 0 new data (the 70888 rows and the push move no number these cards use). Open, with strong evidence: every logged Rend tick runs 30 to 130% above the client's base and grows with the warrior, a server-side term the client row does not show; the formula needs a level 60 advanced log. Still no Arms, Fury or level 60 warrior log. Engine pin 4c595ee6b to [PIN].**

## Method

1. **Client.** Every spell, rank, talent, stance, proc and set bonus the Arms and Fury cards cast or the four builds take, read in the fork (worktree at 4c595ee6b) against the hotfixed 70888 tables (SpellEffect, SpellMisc, SpellAuraOptions, SpellCooldowns, SpellPower, SpellCategories, SpellClassOptions, SpellLevels, SpellDuration, SpellEquippedItems, SpellShapeshift) and the site's 70790 `spellconst`, `spelltips` and `talents`; the push layer (`hotfix-push.json`, 1.60.1.70790) touches no warrior row these cards use. The melee-hunters lane's 64-row table (`melee-hunters/client-warrior.md`) stands; this pass re-read the rows that lane listed as questions and the ones its table did not reach: the attributes and proc masks of Deep Wounds (12834, 12162, 412613), Flurry (12319, 12966), Enrage (12317, 12880), Death Wish, Dual Wield Specialization, Impale's class mask, Two-Handed Weapon Specialization, Reckless Swings, Sweeping Strikes, Anger Management, Furious Precision, Recklessness, Berserker Rage, the Berserker Stance passive, Hand of Justice and Felstriker, and the attack table against a level 63 target.
2. **Blizzard.** The cached "WoW: Forever Hotfixes" (t/2403040, posts 1 to 8), "Class Tuning Incoming" (t/2404092) and the patch notes (t/2360696): every warrior line read against the engine; none since 6 November moves a number these cards use (the melee-hunters lane's reading stands).
3. **Real fights.** foreverlogs.gg: all 26 public reports, 2815 to 2854 (2854 is new since the melee lane: a world log with a Protection tank at about 25 and two leveling warriors with under 20 swings). Our API: two new reports today, both with 0 fights. Still no level 60 Arms or Fury log, no level 45+ warrior at all: the highest-level warriors with a usable fight are the 2834 Fury fragments at 40 to 42 (17 to 24 swings) and the tanks at 30 to 43. So the log pass used every warrior in the 26 reports (17 with Rend or Deep Wounds data) plus our 8 November advanced report's tank (attack power and level on every event), and says so where a check is below 60.
4. **Sims** on the Windows box only (`tools/remote-test.sh`, one job at a time, `nice -n 15`, GOMAXPROCS=5): `spec-breakdown` on the published entries at 10,000 iterations, seed 1, at the pin and at the fix (`dps-warriors/job_b_bands.txt`); the talent and rotation searches at 60 on the fixed engine (`job_c_searches.sh`); the fork suites in a throwaway worktree of the box's fork clone (`fork_tests.sh`).

## Engine against the client

The rows this pass added to the melee-hunters table. Every number is the hotfixed 70888 row unless marked; 70790 agrees on all of them.

| spell or talent | client | engine | verdict |
|---|---|---|---|
| Deep Wounds 12834 | aura 42 trigger 12162 (a dummy with Mechanic 15, bleed) on ProcTypeMask 0x11154: done melee auto, melee ability, ranged auto, ranged ability, harmful spell and periodic; the server's own tick spell is **Deep Wound 412613**: Effect 2 school damage, Mechanic 15, SpellMisc Attributes_2 0x20000004, which carries **SPELL_ATTR2_CANT_CRIT** (0x20000000), Attributes_3 0x2000000 (treat as periodic) | the bleed's four 3 s ticks rolled crit at the melee crit chance under the owner's bleed rule | CONFIRMED-BUG, fixed (B1): the ticks never crit |
| Deep Wounds from a periodic crit | the mask's periodic bit (0x10000) would let a Rend crit tick proc it | only a direct melee or ranged crit procs it | aligned with the logs (L2): 46 Deep Wound runs, every one from a melee crit, two Rend crit ticks inside a run that neither started nor extended it |
| Flurry 12319 and its buff 12966 | trigger ProcTypeMask 0x15554 (every done hit kind), the buff ProcCharges 3 on ProcTypeMask 4 (white swings), DurationIndex 8 = 15 s, aura 319 +30 at the max rank | procs on any melee crit, three charges spent by white swings (and a missed Heroic Strike, Tawl's log), no duration | aligned in effect: a dual wielder spends three charges in well under 15 s; the client's 15 s cap is noted |
| Enrage 12317 and 12880 | ProcChance 30 on ProcTypeMask 0x222A8 (every taken damage kind, periodic included); the buff aura 79 physical, 12 s, 2% a rank | 30% on any damaging hit taken, 2% a rank for 12 s | aligned; a raid DPS seat takes no damage in the sim, so the five Arms raid points and the two Fury points buy only the tiers (the guide says so) |
| Death Wish 12328 | aura 79 +20 physical, aura 77 misc 5 (fear immunity), aura 87 +5% taken all schools; 10 rage, 3 min, 30 s | the same | aligned |
| Dual Wield Specialization 23584 | e0 aura 122 (off-hand damage) 5 a point, e1 dummy 5 (the rage clause), ProcTypeMask 0x800004 (melee auto, off-hand swing) | +5% off-hand damage and +10% off-hand rage a point | aligned (the rage clause is the text's) |
| Impale 16493 | aura 108 misc 15 (crit damage bonus) 10 a point, class masks 0xEE604CEE / 0x8140 / 0x1 / 0x2 | +10% a point of the crit bonus on every warrior ability, Rend's ticks and Whirlwind's off-hand strike included | aligned (bits 5, 29, 30 and 31 of word 0 are Rend, Execute, Whirlwind and Heroic Strike; word 1 bit 2 is 462891; word 2 bit 0 is Spearing Strike) |
| Two-Handed Weapon Specialization 12163 | aura 79 misc 1 (physical), EquippedItemSubclass 0x21562 (the two-hand classes), 2/4/6 by curve | x1.02/1.04/1.06 on physical when the main hand is two-handed | aligned |
| Reckless Swings 1329555 | ProcCharges 4 on mask 4, 8 s, 15 s cooldown, 1 s global, PowerType -2 PowerCostPct 25, trigger 1329554 energize 50 | the same; the health cost is paid when tanking | aligned |
| Sweeping Strikes 12292 | ProcCharges 5 on mask 0x14 (melee auto and ability), 20 s, 30 rage, 30 s, Battle Stance | the same, never cast on one target | aligned |
| Anger Management 12296 | e0 aura 85 (power regen, rage) 15; e1 dummy 1, e2 dummy 3, e3 dummy 30: the text's "1 Rage every 3 sec" and the 30% decay cut are the client's own dummies | 1 rage every 3 s | aligned to the client's dummies; the server's tick size is still unlogged |
| Furious Precision 1323963 | aura 54 (hit) +10 at the max rank, EquippedItemInvTypes 1<<22 (off-hand weapon) | +4/7/10% hit on off-hand attacks | aligned |
| Recklessness 1719 | aura 290 +100 crit, aura 87 +20% taken, aura 77 misc 5; 15 s, 30 min, Berserker Stance, 1.5 s global | the same | aligned |
| Berserker Rage 18499 | aura 77 misc 5 and 14 (fear and sap immunity), ProcTypeMask 0x222A8 (rage from damage taken), 10 s, 30 s, Berserker Stance; Improved Berserker Rage 5/10 instant | the same | aligned |
| Berserker Stance passive 7381 | aura 290 +3 crit, aura 87 +10% taken, aura 10 -20 threat, aura 166 at 0 | +3% crit, +10% taken, x0.8 threat | aligned |
| Whirlwind's off-hand strike 462891 | Attributes_3 0x1000200: requires an off-hand weapon, not a proc | cast by Whirlwind when dual wielding, off-hand special mask | aligned |
| Spearing Strike 1310222 | e0 121 normalized weapon damage, e1 31 weapon percent 75, e2 dummy 1 (the giant and dragonkin clause) | 75% of normalized weapon damage, twice against Giants and Dragonkin | aligned |
| Execute 20662 | e0 dummy 600 (EffectChainAmplitude 1.5, unused), e1 trigger 26651 (a 1 s self dummy) | the rank's flat plus 15 a rage over the cost, the cost refunded at 80% on a failed cast, every point spent | aligned on the client; what a failed Execute keeps is server side (open) |
| Hand of Justice 15600 | ProcChance 3, ProcTypeMask 0x14, ProcCategoryRecovery 2000 | 1% a melee hit, 2 s cooldown (round 21's reading of the Dwarf multiplier) | aligned with round 21 |
| Felstriker 16551 | aura 290 +100, aura 54 +100, 3 s | 1 PPM, 3 s of 100% hit and crit | aligned (the PPM is vanilla's, unlogged) |
| the attack table at 60 against 63 | 300 weapon skill against 315 defense (no race carries weapon skill in Forever) | 8% miss plus 1% hit suppression, 6.5% dodge, 40% glancing at 0.55 to 0.75, 1.8% crit suppression on top of 3%; dual wield +19% white miss | Era's table, unchanged and unlogged at 60 (physical-logs L13; open) |
| normalized rage | | 3.46 a point of speed on a one-hander (log-confirmed), 4.50 on a two-hander, half on the off hand, double on a crit | the two-hand and off-hand constants stay unlogged (open) |

The Arms raid build cannot reach Death Wish: Mortal Strike needs 31 Arms points, which leaves 20 for Fury, and Death Wish is the 21st (tier 4 needs 20 beneath it). The talent search's silence on it is right; the probe's +27 DPS for Death Wish is a build that gives up Mortal Strike.

## Logs against the engine

| # | finding | evidence | verdict |
|---|---|---|---|
| L1 | **Deep Wound ticks never crit.** 146 Deep Wound ticks from 17 warriors across 2815 to 2853 hold 0 critical ticks; the same warriors' Rend ticks crit 13 times in 193 (6.7%, their melee crit at 20 to 43). At 6% a tick, 146 plain ticks in a row is a one in ten thousand draw. | foreverlogs.gg damage streams | CONFIRMED-BUG, fixed (B1); the client's Can't Crit flag on 412613 is the rule |
| L2 | **Deep Wounds procs from melee crits only.** 46 Deep Wound runs (a first tick with no tick in the 3.3 s before it on that target): 45 have a melee or ability crit 2.6 to 3.4 s before the first tick, 1 has no crit event in the window (a crit the parser dropped), 0 have only a Rend crit tick there. Two runs with a Rend crit tick inside ran four ticks and stopped. | the same | aligned (the engine procs it off direct crits only) |
| L3 | **Rend ticks run above the client's base at every logged level, and the gap grows with the warrior.** Rank 3 (6547, 45 over 15 s, 9 a tick): 12 to 14 a tick for eight tanks in Defensive Stance at 20 to 28 (13.3 to 15.6 before the stance's 10%), 14 to 15 for Doublewide, 19 for Pronkers (Fury, Battle Stance), 21 to 23 for Garry (Arms, dual wield). Rank 2 (7 a tick): 9 to 11. Rank 4 (11 a tick): 18 for Miniwaz (Defensive), 20 to 21 for Shmeegle and Docked (Battle). The first tick of a fresh run equals every later one (no rollover), no druid Mangle or other bleed debuff is in any of these groups, and Defensive Stance's 10% shows in the Battle-to-Defensive pairs (Garry 23 to 21, Adriän 14 to 12.6). Our 8 November advanced tank (Bùckk, scaled level 21, rank 4) ticks 16 to 19 and the tick rises with his attack power snapshot (16 at 617, 17 at 652 to 674, 18 at 680 to 733, 19 at 763). | foreverlogs.gg 2815 to 2853; our report gxrgvz4yzsi2 | OPEN, a real gap: Forever's Rend carries a server-side term the client row does not show (an attack power or weapon share, 30 to 130% of the base at 20 to 35; the Thunder Clap precedent). The formula does not pin from these logs: the tanks' Improved Rend is unknown, the dual wielder breaks a plain attack power fit, and no log holds a warrior's own attack power at a known level and talent string. Needs a level 60 advanced log with COMBATANT_INFO. Not changed |
| L4 | **Enrage procs off damage taken in the open world.** A leveling Arms warrior in 2854 (Immortal) gains Enrage 4 times while taking hits; the raid DPS seat never does. | 2854 buffs | aligned |
| L5 | **Unbridled Wrath, Overpower's window, Rend's and Deep Wounds' tick timing, Slam's cast.** The melee lane's figures stand; 2854 adds nothing above 30. | melee-hunters/logs-foreverlogs.md | aligned |
| L6 | **No level 60 Arms or Fury log.** Cast mix, rage a second, Flurry uptime, Heroic Strike count, white and glancing rates and damage shares at 60 stay unmeasured. The highest-level warriors with a fight are the 2834 Fury fragments at 40 to 42 (Bloodthirst 2 and 3, Whirlwind 3 and 4, Death Wish, Reckless Swings and Flurry seen, 17 and 24 swings) and the tanks at 30 to 43. | foreverlogs.gg public list, our API feed | open |

## B1: Deep Wounds' ticks never crit (CONFIRMED-BUG, fixed)

**Finding.** The engine's periodic crit rule (owner, 2026-10-31: every bleed crits) reached Deep Wounds, so its four ticks rolled crit at the melee crit chance with the 2.0 physical multiplier. Forever's Deep Wound is not a bleed aura but the server's 412613, a school-damage spell the client marks Can't Crit (SpellMisc Attributes_2 0x20000000) and treat-as-periodic; it is the one bleed in the warrior's kit with the flag (Rend 11574 has neither bit). The launch logs read the same: 0 critical ticks in 146 (L1) against Rend's 6.7%.

**Fix.** `core.DotConfig.NeverCrits`, which keeps a dot off the crit roll whatever the bleed rule says, set on Deep Wounds with the client row and the log count beside it (E [FORKSHA]). Red first: `TestDeepWoundsTicksNeverCrit` (four ticks at 200% crit, want none critical) fails at the pin with 4 critical ticks and passes at the fix; it replaces `TestDeepWoundsTicksCritWithoutRefreshing`, whose comment also misread 12834's mask as having no periodic bit (it has 0x10000; the logs, not the mask, are what rule a periodic proc out, L2).

**Before and after** (spec-breakdown, 10,000 iterations, seed 1, the published entries; `dps-warriors/job_b_bands.txt` at the pin and at the fix):

| entry | before | after | change |
|---|---|---|---|
| Arms raid 60 Horde / Alliance | 945.2 / 965.0 | 920.3 / 938.5 | -2.6% / -2.8% |
| Arms bare 60 Horde / Alliance | 322.9 / 325.0 | 312.1 / 313.0 | -3.3% / -3.7% |
| Arms bare 50 / 40 / 30 / 20 (Horde) | 225.1 / 148.8 / 65.7 / 32.5 | 220.3 / 147.1 / 65.5 / 32.5 | -2.1% / -1.2% / -0.3% / 0.0% |
| Fury raid 60 Horde / Alliance | 1028.3 / 1030.8 | 1006.9 / 1007.6 | -2.1% / -2.2% |
| Fury bare 60 Horde / Alliance | 371.3 / 374.1 | 361.0 / 362.8 | -2.8% / -3.0% |
| Fury bare 50 / 40 / 30 / 20 (Horde) | 232.8 / 148.1 / 63.6 / 36.5 | unchanged | 0.0% |

Errors are 0.6 DPS at raid 60 and 0.2 at bare 60. The cut is Deep Wounds' crit share: a quarter to a third of the bleed at 60, where Arms crits 35 to 40% of its hits and Fury a little less; the Fury bare bands below 60 and Arms band 20 take no Deep Wounds, so they do not move, and Arms band 30 moves 0.3% at a 6% crit chance. The relative gains are what the prose keeps; the absolute figures are the job files' and the nightly's.

Protection warrior's threat moves with it (Deep Wounds is on the tank build): the suite golden `TestP1TankWarrior` regenerated with `TestP1DPSWarrior`; the published tank figures wait for the refresh (MR row below).

## Builds, rotations and rank rows

[SEARCHES]

Arms and Fury have no rank rows. Leveling orders: the round 29 walks and the melee lane's Arms walk stand; Deep Wounds keeps about three quarters of its damage at the levels the orders take it (crit is 5 to 10% there), it is the prerequisite of Impale either way, and nothing else competes for those points, so the orders were not re-walked.

## Held and open

| # | item | why |
|---|---|---|
| O1 | Rend's hidden scaling (L3) | 30 to 130% over the client's base in every logged warrior, growing with the warrior's power; no log holds a warrior's own attack power at a known level and talent string. A level 60 advanced log settles the formula; if it is an attack power share, Arms raid Rend roughly doubles and the card's Rend-first line gains |
| O2 | a failed Execute's extra rage | the engine spends every point before the roll and refunds 80% of the cost; the server's rule is unlogged (no Execute dodge in any log with power snapshots) |
| O3 | Anger Management's tick size | the client's dummies say 1 every 3 s, the engine does that; the server may read aura 85's 15 |
| O4 | two-hand and off-hand normalized rage, the attack table at 60 (glancing 40% at 0.55 to 0.75, dual wield miss 19%), Weaponmaster's sword cooldown, Felstriker's PPM, Hand of Justice's 1% | unlogged at 60 |
| O5 | Flurry's 15 s cap | the client's buff lasts 15 s; the engine's never expires; a dual wielder spends the three charges in under 6 s, so no number moves; a two-hand Fury would |

## Gates

[GATES]

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| Arms and Fury on Deep Wounds without crits | warrior-arms and warrior-fury, every bare band 20 to 60 and raid 60: picks, weights (crit and Agility worth a little less, Strength and attack power a little more), the race sweep, the published set_dps | B1 |
| Protection warrior on Deep Wounds without crits | warrior-protection, every band and raid 60: threat and the published metrics | B1 (shared talent) |

## Shas

[SHAS]

## Sources

- wago.tools DB2 CSVs for 1.60.1.70888 with `useHotfixes=1` (the 12 November fetch), the site's 1.60.1.70790 `spellconst`, `spelltips`, `talents` and `hotfix-push.json`.
- Blizzard: t/2403040 (posts 1 to 8), t/2404092, t/2360696 (cached copies).
- foreverlogs.gg reports 2815 to 2854; foreversixty API `/v1/reports/recent` and the 8 November advanced report gxrgvz4yzsi2.

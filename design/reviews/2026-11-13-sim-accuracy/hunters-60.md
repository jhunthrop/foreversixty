# Sim accuracy 2026-11-13: the three hunters at 50 and 60, bare and raid

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main d730bbcac, fork from master 5515ae0ef (the pin). Client on the site: 1.60.1.70790; live client 1.60.1.70938 (wago.tools, 13 November 01:36 UTC, one build past 70888) with every hotfix push layered (`useHotfixes=1`). The first dedicated engine-against-client pass at 60 for Marksmanship, Beast Mastery and Survival: what the raid builds, raid preset and raid gear touch, after the 12 November melee-hunters lane's spell-by-spell pass (about 115 rows aligned, `2026-11-12-sim-accuracy/melee-hunters/client-hunter.md`).

**Totals: 3 confirmed engine bugs, all fixed. Devilsaur Eye, worn by every hunter band at 50 and 60 except Beast Mastery bare 60 Alliance, carried Zul'Gurub's use (150 attack power and 2% hit for 20 s every 2 minutes); Forever's level 50 Devilsaur Eye uses 24352 to root a Beast on the next hit, on a 5 minute cooldown, with no stats (old). Gloves of Holy Might's Smite ran at one a minute on any melee hit; Blizzard's 9 November fix, native in 70938, makes it 15% on a white swing or a harmful spell with 20 s between procs, never on a melee special (new data, shared: rogues, warriors and four more specs wear the gloves). The Summon Hawk's swing was the beta's after-armor reading times 1.8; our own advanced logs after the hotfix read it before armor at 1.10 a level, 0.96 unbuffed (old, nothing published casts it). Before and after: Marksmanship raid 60 -1.4%, Beast Mastery raid 60 -1.3%, Survival raid 60 -2.1 to -2.2%; bare 60 -1.1 to -2.3%, bare 50 -1.9 to -2.6%. Devilsaur Tooth carried Zul'Gurub's use too (no band wears it). Logs: no hunter at 50 or above in any readable log (raids open 9 December). Held: Summon Hawk still loses to the published builds with the bigger swing (raid -1.3 to -2.8%, bare -3.6 to -6.1%). Engine pin f46237344 to 79358b3c2 (the lane's 6b5b5c079 merged with the fork's master f46237344, the Discipline lane's pin).**

## Method

1. **Client at 60.** Every row the three raid builds, the raid preset and the 50 and 60 band gear touch, against the 70888 and 70938 hotfixed tables (`hunters-60/` keeps the helper; the 70888 cache is the melee-hunters lane's): Aimed Shot, Multi-Shot, Arcane Shot, Sniper Shot, Serpent Sting, Hunter's Mark, Rapid Fire, Trueshot Aura (20906 at 100, native since 70888), Aspect of the Hawk (14322, 110 and Quick Shots), Heart of the Lion (409580 self, 409583 area: 10% stats, 40 + 4 a level attack power to 60), Bestial Wrath (19574: pet only, 50% damage, 18 s, 120 s, 12% base mana), Lone Wolf (1310684, a server dummy; Blizzard's 20%), the pet's 10% attack power and full crit and hit, the raid preset's buffs and their ranged attack power, the racials (Forever has no bow or gun skill racial: 26290 and 20595 are gone; Berserking and Blood Fury match), ammo and quiver by level, the auto shot timer, and the attack table at 63 (Era's: 8% miss and 1% hit suppression at 300 skill, 4.8% crit suppression). Then every item effect on a hunter band at 50 and 60: ItemXItemEffect, ItemEffect and the use spell's own rows.
2. **The 70938 client.** A table diff of 70888 against 70938 (SpellEffect, SpellMisc, SpellCooldowns, SpellPower, SpellAuraOptions, SpellCategories, SpellClassOptions, SpellLevels, SpellDuration, SpellCastTimes, SpellProcsPerMinute, SkillLineAbility, SpellEquippedItems, SpellName, ItemEffect): no hunter spell row moved; ItemEffect 238916 (Gloves of Holy Might) is now native.
3. **Blizzard.** "WoW: Forever Hotfixes" (t/2403040) to post 9 (12 November 19:10 UTC), "Class Tuning Incoming" (t/2404092; Kaivax's only post is 6 November) and Kaivax's posts since 6 November. Hunter lines: 9 November pet families and far aggro; 10 November pets at 115% health, Arcane Shot and Summon Hawk at 11%, the hawk's swing +80%; 11 November Volley procs and channel. Items: 9 November "Gloves of Holy Might now work as intended". Nothing new for hunters since the melee-hunters lane.
4. **Logs** (`hunters-60/logs/logs-60.md`): our own seven reports (the backfill: 1,178 fights, 412 hunter rows), foreverlogs.gg to report 2854 (the newest; 2855 to 2900 are 404), Warcraft Logs (HTTP 403 to automated fetches, no Forever zone found, the API needs a client token).
5. **Sims** on the Windows box only (`tools/remote-test.sh --engine-sha`, one job at a time, `nice -n 15`, GOMAXPROCS=5): `spec-breakdown` at 10,000 iterations, seed 1, on every published band that wears Devilsaur Eye or Gloves of Holy Might, before and after (`jobs_items.txt`, `results/items-*.txt`, `results/delta.txt`); the full breakdowns at bare 50, bare 60 and raid 60 for both factions (`base.sh`, `results/base-*.txt`); Summon Hawk against the published Beast Mastery builds (`jobs_hawk.txt`, `results/hawk-fix.txt`); the talent and rotation searches at raid 60 on the raid builds and bare 60 (`searches.sh`, `results/searches-fix.txt`).

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| D1 | **Devilsaur Eye carried Zul'Gurub's use.** `hunter/items.go` gave 150 attack power, 150 ranged attack power and 2% hit for 20 s on a 2 minute cooldown. Forever's 19991 is a level 50 rare (item level 52, 16 attack power): ItemEffect 101844 uses 24352 on a 300,000 ms cooldown, and 24352 "Devilsaur Glare" is a 10 s aura with one dummy effect (aura 4) and ProcTypeMask 4436, "Your next melee or ranged attack within 10 sec will Root the target for 2.5 sec ... Only works against Beasts." Nothing for a boss. The ranker had put it on 18 of the 20 hunter bands at 50 and 60 for the phantom use. | client 70790, 70888 and 70938 ItemEffect and SpellEffect; Wowhead Forever's tooltip | old | E a46598e62 (no effect registered; the item's stats only) |
| D2 | **Devilsaur Tooth carried Zul'Gurub's use.** The engine gave the pet's next Bite or Claw a guaranteed crit every 2 minutes. Forever's 19992 uses 24353 (ItemEffect 101573, 10 min), aura 108 -100% on Revive Pet's mask: "Reduces the Mana cost of your next Revive Pet cast within 30 sec by 100%." No band wears it. | same | old | E a46598e62 |
| G1 | **Gloves of Holy Might proc on the old trigger.** The engine rolled the Smite (1325229, 112 +/-10% Holy, the ranged table) at one a minute on any melee hit, specials included. Blizzard, 9 November: "Gloves of Holy Might now work as intended"; client 70938 (and the push on 70790) moves ItemEffect 238916 from TriggerType 2 (chance on hit) to TriggerType 1 (equip) on 1330855: aura 42, ProcChance 15, ProcCategoryRecovery 20000, ProcTypeMask 69636 (a melee auto attack, a harmful spell of no class, a harmful magic spell). So 15% on a white swing or a harmful spell, 20 s apart, and never on a melee special or a ranged attack. Shared item code: the gloves are on Survival 40 to 60 and Beast Mastery raid 60, and on Assassination, Combat, Subtlety, Arms, Fury, Retribution, Enhancement and Feral bands. | Blizzard t/2403040 (9 Nov); client 70938 ItemEffect 238916, SpellEffect and SpellAuraOptions 1330855 | new data | E 4bce59129, `common/item_effects/client_procs.go` |
| H1 | **The hawk swung for half what our logs show.** `summon_hawk_assault.go` took the beta logs' non-crit mean (0.32 a level, read off talented hits on armored mobs as if unarmored) and multiplied it by 1.8 for the 10 November hotfix: 0.576 a level, 17.3 at 30. Our own advanced logs give the swing before armor (`base_amount`): 447 landed non-crit swings from 21 hunters at 17 to 36, median 1.10 a level (p10 1.04, p90 1.25), 33 at 30, the same on 9 to 10 November (1.107, n=26) as after (1.100, n=377). The bulk is Beast Mastery at Unleashed Fury 5, which the hawk takes on top, so the unbuffed swing is 1.10 / 1.15 = 0.96 a level. The engine's 18 s life and 2.26 s swing hold (no hawk swings past 18.3 s, n=136; median gap 2.24 s, n=418), and the dive (32 + 5% of ranged attack power) is not contradicted (the logs' attack power field is melee). No published build casts Summon Hawk. | own logs (`logs/logs-60.md`, "Summon Hawk in detail") | old | E b21d938ae |

Red first on the box: `TestForeverDevilsaurTrinketsGiveNoCombatBonus` (attack power 597.8 to 747.8, ranged attack power 675.2 to 825.2, hit 0 to 2; the pet's Bite and Claw crit rating 0 to 100), `TestGlovesOfHolyMightProcsOnWhiteSwingsAtFifteenPercentEveryTwentySeconds` (8 Smites from 328 white swings inside 0.44 s, and 8 from 364 Wing Clips; want 1 and 0) and the hawk's swing test (0.576 a level, want 0.96). Green on the fixes; the fork's whole `./sim/...` with `-tags with_db` passes at b21d938ae (6b5b5c079 changes only the trigger's action id) with no suite golden moving, except `sim/conformance`'s stale `warlock.golden.md`, which fails the same way at the pin 5515ae0ef (for the warlock lanes), and `sim/web`'s setup on the box.

**Before and after** (spec-breakdown, 10,000 iterations, seed 1, the published entries; `results/base-pin.txt`, `base-fix.txt`):

| spec | raid 60 Horde / Alliance | bare 60 Horde / Alliance | bare 50 Horde / Alliance |
|---|---|---|---|
| Marksmanship | -1.39% / -1.41% | -1.65% / -1.10% | -1.94% / -1.87% |
| Beast Mastery | -1.33% / -1.31% | -1.48% / 0.00% (no Eye) | -1.98% / -1.92% |
| Survival | -2.17% / -2.11% | -2.25% / -2.21% | -2.55% / -2.50% |

The other bands that wear the gloves (`results/delta.txt`, both factions): Arms bare 40 to 60 -0.6 to -0.8%, raid 60 -0.4 to -0.5%; Fury bare 40 to 60 -0.1 to -0.4%; Assassination bare 40 and 50 +0.1 to +0.4%, bare 60 -0.2 to -0.3%, raid 60 -0.3%; Combat bare 40 and 50 +0.2 to +0.5%, bare 60 0.0 to +0.1%; Subtlety bare 40 and 50 +0.2 to +0.4%; Retribution bare 40 and 50 +0.2 to +0.3%; Enhancement bare 40 -0.2%; Feral bare 40 +1.2%. A fast white swing gains (15% with a 20 s gap fires more often than one a minute); a slow swing that leaned on its specials loses.

## Checked, nothing to change

- **Raid preset at 60.** The hunter takes Heart of the Lion's area buff (280 attack power and ranged attack power, 10% stats), Trueshot Aura's 100, Hunter's Mark's 143, Battle Shout, Blessing of Might and Strength of Earth only on melee attack power (BM and MM gain nothing from them; Survival does), Leader of the Pack or Moonkin Aura once (3%), Juju Might on both attack powers, and Judgement of Wisdom on every landed shot (about 6,000 mana a fight). The pet refuses every player buff the client flags "players only" and takes Heart of the Lion; Trueshot Aura carries no ranged attack power to a pet that has none.
- **Bestial Wrath, Lone Wolf, the pet's share.** As the client and Blizzard state (method 1). The pet's 10% of the hunter's higher attack power and his full crit and hit stay sourced to the guides and round 20; no 50 or 60 log can test them.
- **Auto Shot at 60.** Eaglehorn Long Bow is 1.8 s in Forever (40 to 76, ItemSparse delay 1800). With the 15% quiver, Quick Shots up 58% of the fight and one Rapid Fire, Marksmanship fires 47 a minute at raid 60 (141.7 in 180 s); the arithmetic closes.
- **Blackhand's Breadth, Burst of Knowledge, Felstriker, Teebu's Blazing Longsword.** Use spell, cooldown and amount match ItemEffect and the spell rows.
- **70938.** No hunter spell row differs from 70888.

## Logs against the engine

No hunter at level 50 or above in any readable log, boss or trash; the highest in combat is 39 (our Scarlet Monastery report). Every measurement asked for at 50 and 60 (cast mix, auto shots a minute, pet share, Aimed Shot cadence, crit, mana, the miss rate at +3) has n = 0. The earliest 60 source is the level 60 dungeons; raids open 9 December. `logs/py/hunter_own.py data/own 50` prints the table once someone qualifies.

Below 50, for context: Summon Hawk is an extra unit beside the tamed pet, two or three out at once (50 of 64 kept swinging after a recast), as the engine models; H1 above. Auto Shot gaps 2.24 to 2.26 s at 30 to 31.

## Builds, rotations and rank rows

On the fixed engine (`results/searches-fix.txt`, `results/rotation-search-fix.txt`; 800 iterations to confirm, seed 7; the rotation search wears the raid build under the raid preset, the Arcane lane's fix):

| spec | talent search, raid 60 (raid build, `raid-keeps.json`) | rotation search, raid 60 | rotation search, bare 60 |
|---|---|---|---|
| Marksmanship | guide best, both factions | Arcane Shot above Serpent Sting, +0.8% (+8.4 +/- 2.1) | autocast Rapid Fire instead of the card's timed line, +0.6% (+2.8 +/- 0.7) |
| Beast Mastery | Improved Stings 1 for Efficiency 1, +0.6% both factions | curated best | curated best |
| Survival | guide best, both factions | Lacerate refreshed at 3 s left, +0.5% (+4.7 +/- 2.4) | curated best |

Every winner is under the adoption rule's 1%, so the builds and cards stand. The Beast Mastery raid search repeats the melee-hunters lane's +0.6% (held there too).

- **Summon Hawk (MR44, held again).** With H1's bigger swing, the published Beast Mastery builds against Endurance Training 3 and Summon Hawk in Arcane Shot's slot (`bm-hawk-a.json` always, `bm-hawk-b.json` with two hawks): raid 60 -2.7% and -1.3% Horde, -2.8% and -1.5% Alliance; bare 60 -6.1% and -3.6% Horde, -6.1% Alliance (always) (`results/hawk-fix.txt`). Still no build takes it.

No guide, card, coach, build, rotation or rank row changed, so no wow-player pass was owed: no hunter guide or card names Devilsaur Eye, Devilsaur Tooth or Gloves of Holy Might, and the trinket and glove picks on the BiS pages move only at the re-rank (MR96, MR97).

## Open

- **Immolation Trap and hit, and what spends the Mongoose Bite window** (melee-hunters): still no Survival in any log.
- **The pet inherits Blackhand's Breadth.** The engine adds the trinket's 5% to the hunter's crit stat and the pet inherits it; in the client it is a mark on the target for the hunter's own attacks (aura 306). Beast Mastery raid 60 at most about 0.1%; no log can tell.
- **The +3 attack table** (Era's 8% miss plus 1% suppression and 4.8% crit suppression for ranged) waits for a level 60 boss log.
- **Gloves of Holy Might's ProcTypeMask_1 32** (a second-word proc flag) is not decoded; nothing a hunter casts depends on it.

## Measure at refresh

- **MR96** hunters on Forever's Devilsaur Eye: hunter-beast-mastery, hunter-marksmanship and hunter-survival bands 50 and 60, bare and raid: the trinket slot (the Eye's use is gone), weights, the race sweep, the published set_dps.
- **MR97** Gloves of Holy Might on the 9 November trigger: every band that wears 867 (Survival 40 to 60 and raid; Beast Mastery raid 60; Assassination, Combat and Subtlety, Arms and Fury, Retribution 40 and 50, Enhancement 40, Feral 40): the hands slot and set_dps.
- **MR98** the client move past 70790 (issue #21): Gloves of Holy Might's trigger is a hand row from 70938; read the generated tables then.

## Sim changes in this landing

- **Fork `lane/sim-hunters60` on master 5515ae0ef:** 61406ad11 and 2654a001c (tests, red), a46598e62 (Devilsaur Eye and Tooth), 4bce59129 (Gloves of Holy Might), d46528ffd (hawk test, red), b21d938ae (hawk swing), 6b5b5c079 (the trigger aura reports only the Smite, so no new action needs a name). Merged with master f46237344 as 79358b3c2; master fast-forwarded to 79358b3c2.
- **Site:** the pin (`sim/enginever/version.go`, `web/src/lib/sim/version.ts`, the three sim fixtures), no ladder golden moved (the ladder wears no Devilsaur Eye or Gloves of Holy Might), this report, `hunters-60/` (scripts, job lists, results, the logs write-up and its scripts; the raw logs stay out of git), the ledger row and MR96 to MR98.

## Sources

- Blizzard: t/2403040 (posts 1 to 9), t/2404092, t/2360696.
- wago.tools DB2 CSVs for 1.60.1.70790, 70888 and 70938, `useHotfixes=1`, and 70888 and 70938 without.
- Our API's reports (read-only), foreverlogs.gg to 2854.

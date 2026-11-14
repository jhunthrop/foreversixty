# Sim accuracy 2026-11-13: warriors on 70938 and the 13 November pushes

Lane `sa-warrior-70938`, brief `design/reviews/sim-accuracy/brief.md`. Scope: Arms, Fury and Protection warrior on the active client 1.60.1.70938 and its hotfix push layer at the engine pin 037798db1: what 70938 and the pushes since 12 November changed for warrior abilities, talents, items and raid buffs; levels 50 to 60 and the raid fight; rage and Rend against our logs. Not redone: everything the 2026-11-12 lanes (`dps-warriors.md`, `protection-warrior.md`, `melee-hunters.md`) and `tanks-60.md` closed.

- Site branch `lane/sa-warrior-70938` from origin/main 31ce477eb.
- Fork branch `lane/sa-warrior-70938` from master 037798db1, the pin at the start.
- Client: 1.60.1.70938; wago.tools hotfixed tables fetched twice (13 Nov 04:49 UTC, the morning cache, and 14 Nov 03:00 UTC), plus the committed push store (`raw/wago-push/entries.json`, newest push 114082).

**Totals: 2 confirmed bugs, both fixed.**

1. **Windfury Totem's attack (new data, shared core).** Push 113887 rebuilt the totem: the rank's spell (8516, 10608, 10610) is no longer a 1 s attack power aura plus one extra white swing but a main-hand weapon strike, and the attack power moved to a dummy on the totem's aura. Our first report after the push logs it exactly so (SPELL_DAMAGE at the instant of the triggering hit, where every earlier report logs SPELL_EXTRA_ATTACKS). The engine still gave an extra white swing. Arms raid 60 +0.7 to +0.8%, Fury +1.9 to +2.0%, Protection warrior threat +0.8%; every melee raid band moves (rogues +3.0 to +3.8%, Survival +6.8%, Enhancement +2.2%, Feral +1.8%, the bear's threat -3.5 to -3.9%). Bare bands carry no totem and do not move.
2. **Rend's attack power term (old, the 2026-11-12 lane's O1, closed).** Every logged Rend tick runs above the client's flat amount. Our 8 November advanced report's raw upload carries the one warrior with a COMBATANT_INFO (talents with no Improved Rend) and a Defensive Stance cast before every Rend: 45 applications at a cast-time attack power of 599 to 769 tick 0.9 x (11 + 0.0129 x attack power), residual 0.47, bootstrap 0.0126 to 0.0132. Arms raid 60 a further +1.4% (+2.0 to +2.2% with the totem), bare 60 +1.5 to +1.6%, bare 50 to 20 +1.8 to +5.4%. Fury and Protection never cast Rend.

Also confirmed from the logs, no change: normalized rage on a two-hander is 4.50 a point of weapon speed (two warriors at 3.60 speed take 16.2 a swing), 3.45 on a one-hander, half on the off hand, double on a crit (the 2026-11-12 lanes' O4, closed). The Fury card takes Hamstring under a Windfury Totem (+1.2 to +1.4% at raid 60, inert without the totem); Arms' raid-only Slam above Execute (+1.1%) is held. Engine pin 037798db1 to f0f90b5f9.

## Method

1. **Client.** The warrior class set (SpellClassSet 4, 518 spells with their triggers) diffed between the 70888 and 70938 hotfixed tables (`c60/diff.py`): one change, float formatting of Execute's unused chain amplitude. The site's `spellconst/warrior.json` 70790 against 70938: two new class-mask rows (1283073 Charge, 1283407 Shield Charge, neither trained), Berserker Rage gaining sap immunity, Mocking Blow's rows (not cast); talents and trainables carry the build stamp only. Then every spell row diffed between the morning and evening hotfixed fetches (`alldiff.py`, 202 rows): Windfury Totem (below), Berserker Rage's new incapacitate immunity (push 114052, Blizzard's "Berserker Rage was not removing Sap"), the flask and elixir attribute flood (Attributes_2 0x100, no number), Wild Growth's category and the Judgement defense types (other lanes'), a hunter set mask.
2. **Blizzard.** "WoW: Forever Hotfixes" (t/2403040) to post 10 (14 Nov 01:54 UTC). Warrior and melee lines: Windfury Totem "now correctly activated by all melee abilities even if it doesn't involve your main hand weapon ... Strikes from abilities that involve your offhand weapon like Whirlwind or Mutilate do not have a chance to activate Windfury"; Berserker Rage and Sap; Taran Icebreaker "now correctly behaves as a 2-handed Mace" (the data already says two-handed, subclass 5, inventory type 17, which is how the engine has always simmed it; it is an Arms band 50 runner-up); Gloves of Holy Might "only activated by spells and abilities that are harmful and deal damage" (push 114008's 0x11154, which the engine carries since the paladins lane); Blessed Band of Light's hit removed (MR110).
3. **Real fights.** Our reports, read only (Cloud SQL proxy, R2 event parquet for all 8 reports, and the one raw upload R2 still holds, the 8 November advanced log). No report since 13 Nov 18:38 UTC; the newest (ju2t73py6pib, Feralas, two Enhancement shamans at 44 to 46 with Windfury Totem rank 2) is the first after push 113887. Warriors with a fight: 61 across the reports, levels 1 to 51 by the advanced log's level field, none at 60. foreverlogs.gg's API refuses automated reads (403 since 13 Nov, per the casters lane) and Warcraft Logs has no Forever partition, so no public report was read.
4. **Sims** on the Windows box only (`tools/remote-test.sh`, one job at a time, `nice -n 15`, GOMAXPROCS=5): `bands.sh` (spec-breakdown on the published entries, 10,000 iterations, seed 1) at the pin, at the Windfury fix and at the Rend fix; the fork suites in a throwaway worktree (`fork_tests.sh`); the rotation and talent searches on the fixed engine (`searches.sh`).

## B1: Windfury Totem procs a main-hand strike (CONFIRMED-BUG, new data, fixed)

**Client.** Morning against evening (push 113887, already in the committed push store):

| row | before | after |
|---|---|---|
| 8516 / 10608 / 10610 SpellEffect | e0 aura 99 (attack power) 95 / 179 / 246 on self; e1 effect 19 (extra attacks) 1 | e0 effect 58 (weapon damage) with no bonus points on the target; e1 removed |
| SpellMisc | DurationIndex 36 (1 s), Attributes_1 0x10000000, Attributes_6 0x400 | no duration; Attributes_3 0x40000600 (0x400 main hand, 0x200 triggered can proc), Attributes_10 0x2000000 |
| SpellAuraOptions | ProcCharges 2 on mask 4 (the buff's two white swings) | removed |
| SpellCategories / SpellCooldowns | none | Category 2619, CategoryRecoveryTime 100 |
| 8515 / 10609 / 10612 (the totem's aura) | e0 aura 42 trigger 8516 at 20%, 100 ms ProcCategoryRecovery, mask 0x14 | the same, plus e1 aura 4 (dummy) 95 / 179 / 246: the "$8515m2" the text's "1 extra attack with ... extra melee attack power" reads |

The text is unchanged. So a proc is one yellow main-hand strike at the weapon's own speed with the rank's attack power: no glancing blow, no dual wield miss penalty, no rage from a white swing, no Flurry charge spent, no queued Heroic Strike taken, and the procs a melee ability rolls (Flurry's trigger, Deep Wounds, Hand of Justice). The engine gave vanilla's extra white swing under a two-charge attack power buff.

**Logs.** ju2t73py6pib (13 Nov 18:38, after the push): Windfury Totem rank 2 as SPELL_DAMAGE 10608 six times and SPELL_MISSED (dodge) three times, each at the same timestamp as the white swing or Stormstrike that triggered it, two crits at double damage (490 and 574 against 242 to 252). Every report before the push (gxrgvz4yzsi2, 3sdrcwdgo2yi, edcgw6rodqs2, eygg26pg4gr2) logs the totem as SPELL_EXTRA_ATTACKS 8516, 728 in all.

**Fix.** `core.ApplyWindfury` registers the strike (MH weapon damage plus the rank's attack power, `OutcomeMeleeWeaponSpecialHitAndCrit`, the shaman's own Windfury Weapon strike's flags) and casts it from the permanent proc aura on a landed main-hand melee hit, 20% with the 100 ms cooldown. The proc aura carries the strike's id, so the Arms and Survival filler gates (`auraIsKnown` 10610) still read whether a totem reaches the character. `CreateExtraAttackAuraCommon`, `WindfuryBuffDuration`, `GetWildStrikesAP` (unused) are gone; `WindfuryBonusAP` and `WindfuryStrikeSpellId` replace the buff arrays. Blizzard's trigger line needs nothing: the engine already took any landed main-hand melee hit, auto or special, and no off-hand strike. Red first: `TestWindfuryTotemProcsAMainHandStrike` ("granted 109 extra white swings", "never struck") at 4d5712668, green at 8e87fe155. The feral, Unsealed Wrath and core client tests now read the strike; Unsealed Wrath's "no echoes from Windfury Totem attacks" holds by construction (a yellow strike is not a Seal swing). Site: the adapter's Windfury extra attack key is gone (no new result carries one; the web still names it in results saved before).

**Before and after** (spec-breakdown, 10,000 iterations, seed 1, the published entries; `bands_before.txt`, `bands_wf.txt`, `bands_after.txt`):

| entry | pin | Windfury fix | with Rend |
|---|---|---|---|
| Arms raid 60 Alliance / Horde | 935.4 / 930.9 | +0.7% / +0.8% | +2.0% / +2.2% |
| Fury raid 60 Alliance / Horde | 1021.0 / 1009.0 | +1.9% / +2.0% | +1.9% / +2.0% |
| Protection warrior raid 60, threat /s | 1673.8 / 1648.3 | +0.8% / +0.8% | +0.8% / +0.8% |
| Assassination / Combat / Subtlety raid 60 | | +3.8 / +3.0 to +3.1 / +3.5 to +3.6% | |
| Retribution / Enhancement / Feral raid 60 | | +0.1 to +0.2 / +2.2 to +2.3 / +1.8% | |
| Survival raid 60 | | +6.8% | |
| Feral bear raid 60, threat /s | 1941.1 / 2147.2 | -3.9% / -3.5% | |
| Protection paladin raid 60, threat /s | | -0.0 to -0.1% | |
| every bare 60 entry above | | 0.0% | |

Errors are 0.4 to 0.7 DPS at raid 60. Who gains is who glanced and dual wielded: the extra white swing glanced 40% of the time at about two thirds damage and missed at the dual wield rate, the strike does neither. The warriors gain least, because their extra swing also paid rage (and took a queued Heroic Strike) and the strike does not; the bear loses, as its threat is rage. Survival's two-hander and Wing Clip filler gain most. Damage taken and chance of death do not move for any tank.

## B2: Rend's tick adds 0.0129 of attack power (CONFIRMED-BUG, old, fixed)

**Client.** Rend 772 to 11574: one effect each, aura 3 periodic damage of 5 / 7 / 9 / 11 / 14 / 18 / 21 every 3 s, no coefficient and no BonusCoefficientFromAP; the text is "$o1 damage over $d". The engine ticked the flat amount.

**Logs.** `rend_pairs.py` pairs each Rend cast's own attack power (the warrior's advanced snapshot on SPELL_CAST_SUCCESS) with that application's non-crit ticks: 229 applications from 35 warriors at levels 1 to 44, ranks 1 to 7. Every tick runs above the flat amount. The warrior that pins it is the 8 November tank, whose raw upload R2 still holds (`rend_stance.py`): COMBATANT_INFO with 23 Protection points and no Improved Rend, Defensive Stance (0.9 damage) cast before every one of 45 Rend applications, 120 ticks of 16 to 20 at an attack power of 599 to 769. Least squares on tick = 0.9 x (11 + k x attack power): k = 0.0129, residual 0.47 (the ticks are whole numbers), bootstrap 95% 0.0126 to 0.0132, no drift with attack power. The other warriors' talents and stance are unknown; at k = 0.0129 every one of them sits within about a point of 0.9, 1.0 or an Improved Rend multiplier (1.12, 1.23, 1.35), across ranks 1 to 7, so the share is a tick's and not a fraction of the whole bleed.

**Fix.** `rendAttackPowerPerTick` = 0.0129 of `MeleeAttackPower` added to the rank's roll at the snapshot (a hand row from the logs, the Thunder Clap precedent; no client column holds it). Red first: `TestRendTickAddsAShareOfAttackPower` at acd929dff, green at 00e13541b. No fork suite golden moved (their rotations do not keep Rend up).

**Before and after,** Arms (the only card that casts it):

| entry | before (Windfury fix) | after | change |
|---|---|---|---|
| raid 60 Alliance / Horde | 941.5 / 938.1 | 954.4 / 951.1 | +1.4% / +1.4% |
| bare 60 Alliance / Horde | 312.1 / 311.6 | 316.7 / 316.5 | +1.5% / +1.6% |
| bare 50 | 219.4 / 221.3 | 223.4 / 225.4 | +1.8% / +1.9% |
| bare 40 | 148.1 / 147.1 | 151.8 / 150.7 | +2.5% / +2.4% |
| bare 30 | 65.4 / 65.5 | 68.1 / 68.3 | +4.1% / +4.3% |
| bare 20 | 31.3 / 32.5 | 33.0 / 34.2 | +5.4% / +5.2% |

At 60 the term is linear past the logged range: 0.0129 x a raid warrior's attack power roughly doubles rank 7's 21 a tick. A level 60 advanced log with a COMBATANT_INFO reads it at that range; until then it is the logged share carried forward.

## Rage against the logs (aligned, closes O4's rage half)

`rage_swings.py`: the warrior's own advanced power on each landed white swing minus its previous snapshot, with no cast, energize or damage taken between (19,636 swings). Rage over the weapon's swing interval (tenths of rage): 34.5 for every one-handed warrior with a clean interval (11 warriors), 45.0 at 3.60 and 3.62 speeds (two two-handers: 16.2 rage a swing), 17.4 to 18.1 for the dual wielders' second cluster (the off hand at half), crits exactly double. The engine's 3.46, 4.50, 0.5 and 2.0 stand. The parser leaves the off-hand flag unset, so the off-hand half is read from the two clusters of each dual wielder, not from a flag.

## Builds, rotations and rank rows

Rotation searches at 60 on the fixed engine (`rotsearch_probe.txt`, `rotsearch.txt`, 800 iterations, seed 7), the winners confirmed with `spec-breakdown -rotation` at 10,000 iterations on two seeds and both factions (`variants.sh`, `variants.txt`):

| spec | search | confirmed | verdict |
|---|---|---|---|
| Fury raid 60 | insert Hamstring after Reckless Swings, +14.4 (+1.4%) | the line gated as Arms' is (a Windfury Totem on you, never from 25% down): +1.3% Alliance and Horde on seed 1, +1.4% and +1.2% on seed 50001 (errors 0.6); bare 50 and 60 identical (the gate is false without a totem) | **adopted**: nothing the card casts is dropped, the gain is beyond 1% and the error everywhere the line can fire |
| Arms raid 60 | Slam above Execute, +9.6 (+1.0%) | +1.1 to +1.2% on both seeds and factions | **held**: the bare 60 search does not find it (its winner is a Rend refresh at 3 s, +0.5%), so it is a raid-only order on a card every band shares, and it turns Arms' execute phase from one rage-dump Execute into Slam with Execute behind it. A raid-only gate is a card design question for the owner, not a sim bug (as `protection-warrior.md` O6) |
| Arms bare 60 | Rend refreshed under 3 s left, +1.6 (+0.5%) | not run | held, under 1% |
| Fury bare 60 | the card is the best found | | |

The probes: every Arms raid line still helps when removed (Overpower +159, Rend +110 with its attack power share, Mortal Strike +80, Slam +48, Execute +43, Hamstring under the totem +33, Reckless Swings +20, Spearing Strike +19); every Fury raid line helps but Berserker Rage and the pre-36 Overpower (Execute +123, Bloodthirst +109, Whirlwind +80, Heroic Strike at 60 rage +70, Reckless Swings +20, Recklessness +19). Rend inserted into Fury measures 0.0 at raid and bare 60: the share does not make it worth a Fury global.

Talent searches at 60 on the fixed engine (`searches.txt`, seed 7, both factions, the raid keep lists): Arms raid, the guide is best on both factions; Fury raid Alliance, the guide is best; Fury raid Horde, 2 Enrage to Blood Craze +3.6 ± 3.0 (+0.35%), under the bar; the bare 60 winners are the 2026-11-12 lane's (Arms drops Iron Will and Deflection for Lingering Rage or Booming Voice, +2.1 to +2.8%, held for the unmodelled Iron Will exactly as then; Fury Horde 2 Enrage to Booming Voice +0.4%). No build moves. No warrior card has rank rows; the leveling orders are unchanged (Rend's share helps Improved Rend, which every order already takes at 10 to 12).

The card change ships with its text: the Fury card's new line and its panel text, the Fury guide's Rage paragraph (it said Fury never Hamstrings) and the coach's priority paragraph; Rend's attack power in the Arms card's Rend note, the Arms coach and the Arms and Fury leveling paragraphs; "the totem's strike" in the Arms Hamstring note; the raid preset's Windfury Totem reason.

## Held and open

| # | item | why |
|---|---|---|
| O1 | Windfury Totem strike's size | in the post-push report the strike lands 1.5 to 1.6 times a white swing of the same shaman (234 to 252 against a median 153), where 179 attack power should add far less; the same shamans' Windfury Weapon strike lands 1.27 times. With no gear in that report (no COMBATANT_INFO) the weapon and talents are unknown, so the client's 95 / 179 / 246 stand. A post-push log with a warrior's COMBATANT_INFO settles it |
| O2 | Rend's share at 60 | the fit is 599 to 769 attack power; a level 60 advanced log reads it at raid attack power |
| O3 | a failed Execute's extra rage, Anger Management's tick size, the attack table at 60, Weaponmaster's sword cooldown, Felstriker's PPM, Hand of Justice's 1%, Flurry's 15 s cap | unchanged from the 2026-11-12 lanes; still no level 60 warrior log |
| O4 | threat factors (Shield Slam, Revenge, Heroic Strike, Sunder's attack power share, Thunder Clap) | no log carries threat (`tanks-60.md` O4) |
| O5 | Retribution's Seals on the strike | the engine's Seals and Unsealed Wrath act on white swings only, so the totem's strike now neither stores nor releases; whether a Seal of Command rolls on the strike in the game is the paladin lane's question |

## Gates

On the box, one run at a time:

- **Fork** (`fork_tests.sh`; outputs `fork_tests.out`, `fork_tests_rend.out`, `fork_tests_apl.out`): `TestWindfuryTotemProcsAMainHandStrike` red at 4d5712668, green at 8e87fe155; `TestRendTickAddsAShareOfAttackPower` red at acd929dff, green at 00e13541b; `go test -tags=with_db -p 1 ./sim/...` at 8e87fe155: every package ok but the two warrior suite goldens (the P1 suite's main-hand "Windfury" imbue is the totem), regenerated at 81e632ff5 (`sim/web` fails setup on the box with or without the lane, a node dependency); at 00e13541b the warrior, core, conformance, druid, paladin, hunter, rogue and shaman packages ok with no golden moved; the Fury card's line moved `TestP1DPSWarrior` once more, regenerated at f0f90b5f9.
- **Site, pinned to f0f90b5f9** (`gates.sh`; `gates.out`, `gates2.out`, `gates3.out`): `npm run lint`; after `npm run sync`, `astro check` 0 errors; `vitest` over `src/lib/sim`, `src/content/guides`, `src/lib/guides` and `src/fixtures`; `apl-check`, every copy matches; `go vet ./...`; `go test ./request/` with the Arms and Fury ladder goldens rewritten (`ladder.patch`), then `./leveling/` (its Windfury rank test now reads the strike from 70938's spellconst), `./rankchoices/`, `./score/`, `./adapter/` (the Windfury extra attack key is gone), `./api/`, `./combine/`, `./cmd/spec-breakdown/`, `./cmd/leveling-bis/`, `./cmd/addon-builds/`, `./cmd/talent-search/`. Two reds found and fixed on the way: the Fury card's `level_notes` are keyed by step and had to follow the new line, and the new note's "never from 25% down" read as a leveling clause at 60 (`input-cards.test.ts`); it now says "while the target is above 25%".
- **Python, on the laptop with uv** (the box has no pytest): `data/tests/test_apl.py`, `test_card_note_levels.py`, `test_addonrotation.py`: 113 passed; the Fury sidecar names Hamstring and the totem, the Arms Rend note's allowlist line takes its new hash. One failure is on origin/main already and not this lane's: the Enhancement "Below a third, keep the mana for Stormstrike" note at band 10.
- **wow-player:** SHIP with SHOULDs, no MUSTs; the 25% reason and "on you"/"target" wording taken; "a strike that cannot glance" left out as explaining the obvious.

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| MR111 Windfury Totem's strike | every raid 60 entry a totem reaches, both factions: warrior-arms, warrior-fury, warrior-protection, the three rogues, paladin-retribution, shaman-enhancement, druid-feral, druid-feral-bear, hunter-survival (and paladin-protection, unchanged within error): picks, weights (hit and crit against weapon damage shift as the glancing extra swing goes), races, set_dps and the tank metrics; the tier list's melee order | B1 |
| MR112 Rend's attack power | warrior-arms every bare band 20 to 60 and raid 60: picks, weights (attack power and Strength a little more), races, set_dps | B2 |

## Shas

- **Fork `lane/sa-warrior-70938` on master 037798db1:** 4d5712668 (Windfury test, red), 8e87fe155 (Windfury strike), 81e632ff5 (warrior suite goldens), acd929dff (Rend test, red), 00e13541b (Rend's share), e64864309 (the Fury card's line, apl-sync), f0f90b5f9 (DPS suite golden). Master fast-forwarded to f0f90b5f9.
- **Site pin:** f0f90b5f9 (`make engine-pin`: `sim/enginever/version.go`, `web/src/lib/sim/version.ts`, the three sim fixtures).
- **Site:** the pin; the Fury card line, its level notes and panel text; the Arms Rend and Hamstring notes; the Fury guide's Rage paragraph and coach priority; the Rend line in the Arms and Fury guides and the Arms coach; the raid preset's Windfury reason; the adapter's dead Windfury key; the leveling Windfury test; the Arms and Fury ladder goldens; the test sidecar and allowlist line; this report, its job files (`warriors-70938/`), the ledger row, MR111 and MR112.
- **Shared code:** `core.ApplyWindfury` (every class a totem reaches); `warrior/rend.go`.

## Sources

- wago.tools DB2 CSVs for 1.60.1.70938 with `useHotfixes=1` (13 Nov 04:49 and 14 Nov 03:00 UTC); the site's 70790 and 70938 `spellconst`, `talents`, `items.json` and `raw/wago-push/entries.json`.
- Blizzard: t/2403040 posts 7 to 10.
- Our reports (8, private ones in aggregate; no player named), the R2 raw upload of gxrgvz4yzsi2.
- Job files: `warriors-70938/` (`bands.sh`, `fork_tests.sh`, `searches.sh`, `gates.sh` and their outputs, `rend_pairs.py`, `rend_stance.py`, `rage_swings.py`, `alldiff.py`).

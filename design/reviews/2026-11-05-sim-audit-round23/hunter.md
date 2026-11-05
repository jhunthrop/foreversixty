# Round 23, lane hunter: Beast Mastery, Marksmanship, Survival (melee)

Client 1.60.1.70790. Site wt/r23-hunter at int/r23 08447fb84 (no site code change); fork lane/r23-hunter from int/r23 46c3ae5b9.
Method: spec-breakdown through `$R/sim.sh`, guide raid builds (`$R/rt`) and `-bis-dir $R/bis-frozen`, 10000 it at seeds 1 and 50001, both factions. Fork tests `-tags with_db`, only `./sim/hunter`.

**Result: 1 CONFIRMED-BUG, fixed (`R22 d4a1e2ca0`). The seed (the druid crit aura counted twice on the pet) is Aligned, now guarded by a test.** No published number moves.

## Seed 1: does the pet take the raid's 3% crit aura twice?

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **No double count.** The pet inherits the hunter's crit (every-table `stats.Crit` plus melee crit, `sim/hunter/pet.go` `makeStatInheritance`), and the hunter's sheet carries the raid's 3% once (`core.RaidHasDruidCritAura`, buffs.go:247). The pet itself takes no player buff: `hp.RefusePlayerBuffs()` (pet.go:108) lets only Heart of the Lion's area buff through (`applyPetBuffEffects`, buffs.go:445). Client: Leader of the Pack 24932 and Moonkin Aura 24907 both carry SpellMisc Attributes_3 = 256 (0x100, "only target players") in 1.60.1.70790 (`scratchpad/b70790/70790/SpellMisc.csv`); 409583 (Heart of the Lion area) has 0. Blizzard 16 Oct (t/2360696 post 6): pets inherit the owner's crit ("whichever is higher, melee or ranged"). (pet.go:107 also quotes "Pets can no longer receive player buffs that increase their stats" as Blizzard's; that line is not in t/2360696 as fetched today, so the client flag is the source for "no buff on the pet".) Test `TestHunterPetTakesTheRaidCritAuraOnceThroughTheHunter`: both druid auras add 3 to the hunter and 3 to the pet; with `RefusePlayerBuffs` removed it fails at +6 (mutation check run, then restored). | client row + Blizzard 16 Oct note + engine test | Aligned | - | E 92be0fd71 (test only) |
| P2 | **Warlock demons take the same path:** `sim/warlock/pet.go:128-129` calls `RefusePlayerBuffs()` and `TakeHeartOfTheLion()`, so a demon never takes the druid aura directly. Demons do not inherit the warlock's crit at all today (fixed own crit, `demonMeleeCritPercent` 0.7, `demonSpellCritPercent` 6.4, pet.go:254-256), so there is no double count; whether they should inherit is the warlock lane's call. | code read | For warlock lane | - | - |

## Round 22's own changes (regressions first)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **The scope's crit reached the pet again on every change to the hunter's stats.** Round 22 put the hunter's ranged-only crit (a scope, `RangedOnlyCritRating`) into the pet's inheritance function as a constant. The engine uses that same function for every in-fight change to the hunter's stats (`core.Pet.Enable` sets `dynamicStatInheritance = statInheritance`; `addOwnerStats` maps each change through it), and the function's own comment says it is linear. So each buff, proc or debuff on the hunter added the scope's 2% to the pet once more: three +1% changes moved the pet +9% instead of +3%. In the game a pet inherits the hunter's higher crit (Blizzard 16 Oct), i.e. his melee crit plus the scope, once. Failing test first `TestHunterPetTakesTheScopeOnceThroughTheFight` (red: 9.00, want 3). Fix: the summon inherits the full sheet with the scope; the hunter's later changes go through the linear shares only (`makeOwnerChangeInheritance`, set in the pet's `OnPetEnable`). Only a hunter with a +crit scope moves (8720; no preset, BiS list or site enchant row carries one). | `sim/hunter/pet.go`, `sim/core/pet.go:110,171-177`; test | CONFIRMED-BUG | R22 d4a1e2ca0 | E bc0291071 |
| R2 | Round 22 rank rows (Arcane Shot, Serpent Sting, Aimed Shot 50-51; S aae25e43b, 0a6bbcacc) and the Survival Raptor Strike note: untouched since round 22, raid 60 identical to the digit (below). Not re-swept: the rank rows are bare-band and nothing in the engine under them changed this round. | measure | Aligned | - | - |

## Seed 2: round 22's open items and guide "Sims requested"

| # | item | verdict |
|---|---|---|
| G1 | Guide round 24 "Sims requested": none new; carried the Marksmanship raid 60 Auto Shot share. Round 22 M3 answered it (46%); the BM breakdown here gives Auto Shot 38.5% with the pet's 20.4% beside it; no engine change since. | Answered |
| G2 | Guide round 24 E1 (Auto Shot through Sniper Shot's 4 s cast): settled in round 22 F1 (client SpellMisc/SpellInterrupts identical to Aimed Shot, Blizzard 16 Oct names Sniper Shot). | Aligned |
| G3 | Round 20's 18-19 low-mana stall (guide rounds 21-23): fixed in round 21 G1; guide round 24 D4 agrees. | Closed |
| G4 | HUNTER-R24-3 / R24-4 (Rapid Fire minutes on the leveling order, the sting note's Aimed Shot reason at 18-19): note text, not sim. | For the guide lane |
| G5 | Round 22 "Still open" (hawk swing, Heart of the Lion on the pet, Focused Fire snapshot, Ranged Weapon Specialization on the sting, Lacerating Strikes, Immolation Trap's 2%, multipliers on the attack power share, the low-rank share): each still needs a live Forever log; none exists (Logs). | OPEN |

## Seed 3: Blizzard's notes since round 22

Thread us.forums.blizzard.com/en/wow/t/2360696 fetched 11:50 UTC 5 Nov: 10 posts, last 4 Nov 19:35 (post 10 is UI and accessibility only; post 9 edited the same minute). Every hunter line in post 9 (3 Nov) was in round 22's N1-N11 table; the two not listed there are "Pets now gain 2 Health per point of Master's Stamina" (health, no damage) and "Roots that occur from Procs now correctly have Diminishing Returns" (Improved Wing Clip, Counterattack roots; no damage). No Forever hotfix thread found (web search 5 Nov). **Aligned, nothing new.**

## Seed 4: fresh re-audit

| spec | checked | verdict |
|---|---|---|
| Beast Mastery | raid build decoded (5420001505001251-00523500014): every point acts or gates (Endurance Training 4 and Bestial Swiftness gate Intimidation/Bestial Wrath; rounds 15, 16, 18); pet Claw 3009 rank 8 (56) and Bite 17261 rank 8 (56) against spellranks.json; pet crit channel includes Lethal Attacks (a `stats.Crit` add, talents.go:490), which the pet should inherit as sheet crit | Aligned (plus R1) |
| Marksmanship | raid build decoded (Lone Wolf, no pet); Improved Tracking 5 counts on the type-neutral target since round 3 X3; Sniper Shot path (G2) | Aligned |
| Survival | raid build decoded; APL order and Wing Clip under Windfury Totem (rounds 13 regmelee); rank sweep round 22 S1 | Aligned |

Not re-done this round: a spell-by-spell coefficient pass (rounds 13, 15, 20, 21 did it on 70755/70790 and the hunter spell rows did not change since), BiS item effects (items lane).

## Before/after

Raid 60, guide builds, frozen BiS, 10000 it, against this lane's int/r23 baseline (`$R/hunter/int-raid.tsv`) and after (`$R/hunter/after-raid.tsv`): identical to the digit for all twelve runs (BM 868.3/868.2 Alliance, 865.5/865.3 Horde; MM 959.3/960.6, 960.0/959.2; SV 851.1/850.4, 850.7/851.2). R1 moves only a hunter with a crit scope (user gear), by about 2% pet crit per in-fight change to the hunter's stats removed.

## Files to regenerate

None. No fork golden changed (`./sim/hunter` passes).

## Measure at refresh

Round 22's list stands.

## For other lanes

- **Warlock:** P2 (demons refuse player buffs; no crit inheritance today).
- **Platform / any lane adding a constant to a pet inheritance function:** `core.Pet` reuses the static inheritance for owner stat changes unless the pet sets its own (`EnableDynamicStats` in `OnPetEnable`, as the hunter pet now does); a constant term in it is added once per owner stat change.

## For the guide lane

G4 (HUNTER-R24-3, HUNTER-R24-4).

## Guide text

No build, APL or card-rank change, so no guide text and no wow-player pass.

## Logs

Last looked 11:54 UTC 5 Nov: api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct); foreverlogs.gg `/api/reports/public` 403; Warcraft Logs v1 `/v1/zones` with the Secret Manager key 401.

## Still open

G5, unchanged from round 22; each needs a live Forever log (a hunter with known sheet ranged attack power, a BM hawk, a Survival trap, a rank 1 Arcane Shot at 40+).

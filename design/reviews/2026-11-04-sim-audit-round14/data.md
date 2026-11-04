# Sim audit round 14: data lane

Scope: racials, raid buffs and the raid preset, consumables, item effects and set bonuses on the published BiS gear, the rank-choice mechanism and its non-healer rows, trainables, and the leveling band walk. Fresh re-audit on client 1.60.1.70753 with the 2026-11-03 14:14 UTC hotfix push layered.

Sources:
- Client tables: the push-layered raw of round 13 (`scratchpad/r13/trainables/root/1.60.1.70753/raw`, the same push the landing committed as `hotfix-push.json`), read with `$R/data/tools/spell.py` and `item.py`.
- Measurements: spec-breakdown, 10,000 iterations, seeds 1 and 50001, `nice -n 15`, GOMAXPROCS 9, one run at a time. Before is `$R/sb-before` from `$R/site-before/sim`. After is engine `lane/r14-data` with site `wt/r14-data`.
- Rank rows: `$R/data/tools/rc.py`. It builds one spec-breakdown per row with that row removed from the embedded table, then runs both binaries on the row's band, preset and races. Results are in `$R/data/rc.tsv`.

Branches: engine `lane/r14-data` (from int/r14 5195a8889), site `wt/r14-data` (from int/r14 72471fcb3).

**Found: 5 CONFIRMED-BUGs, all fixed.** By origin: R13 2, old 1, new-data 2.
- **R13 (2):** the Shadow Devouring Plague race row and the Destruction Corruption row.
- **old (1):** Expansive Mind's rage and energy.
- **new-data (2):** the Beast Mastery Arcane Shot rows, and the Trueshot Aura preset reason. The same hotfix push caused both.
- Raid 60 does not move for any published entry. The bare bands move where a rank row changed.

Round 13's in-game checks are still **pending the owner**, and this lane did not reopen them: the Arcane Shot tooltip with spell power, the level-60 trainer book ranks, the six Wowhead-baseline talents, Heart of the Lion 280 or 276, and Rejuvenation rank 1's tooltip. E 05b72a3b5 stays.

## Racials

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | **A Gnome's Expansive Mind did nothing for rage or energy.**<br>- **What the client says:** three rows, each aura 178 (maximum power percent) at 5. 20591 is mana (misc 0, class mask 400: priest, mage, warlock). 1259802 is rage (misc 1, mask 1: warrior). 1259803 is energy (misc 3, mask 8: rogue). So a Gnome warrior holds 105 rage and a Gnome rogue 105 energy.<br>- **What was wrong:** the racial added a Rage stat and an Energy stat. No code reads either one: the rage cap is the constant 100 in `rage.go`, and the rogue's energy bar is enabled at 100 plus Vigor. Only the mana part worked.<br>- **What changed:** the rage bar now carries its own cap (`MaxRage()`, `MultiplyMaxRage`). The racial raises whichever bar the class has enabled; the class builds the bar before race effects apply. Vigor's +5/+10 is multiplied too, as a percent of maximum is.<br>- **Effect:** at raid 60, a Gnome rogue gains +0.2 to +0.5% and a Gnome warrior +0.04 to +0.18%. No published entry is a Gnome warrior or rogue. | SpellEffect 20591, 1259802, 1259803 (SkillLineAbility skill 753). `racials.go:437-449` (old), `rage.go:243`. Test first: `TestExpansiveMindRaisesMaximumRageAndEnergy` failed ("maximum rage is 100, want 105"; "maximum energy is 100, want 105"), then passed. `go test ./core` is green. | CONFIRMED-BUG | old (5f2da5165, 2026-09-18; round 13 auras R3 checked the class split, not the plumbing) | E 6411690c9 |
| R2 | Re-read against the client on fresh eyes, each published race's racials:<br>- Blood Fury: aura 166, 167 and 317 at 10, 15 s, 2 min, no global cooldown. The multiplier covers attack power, spell power and healing.<br>- Berserking: aura 319, 140 and 65 at 10, 15 s, 3 min.<br>- Elune's Light: aura 290 at 10, 15 s, 3 min.<br>- Ruthless: 1% crit. Artisan's Precision: 1% hit. The Human Spirit: 10%.<br>- Wind Blessed: 2% to all three hastes.<br>- Endurance: 5% health. Quickness: 1% dodge.<br>- The three mob-type 5% racials.<br>- Read Leyline: APL-cast only; no published entry is Skyborne.<br>All match. | SpellEffect, SpellMisc, SpellCooldowns rows; `racials.go`, `racial_blood_fury.go` | Aligned | - | - |
| R3 | `GetFaction()` returns Unknown for both Skyborne rows. Its only reader is Warchief's Blessing, a world buff the preset never uses. | `racials.go:577`, `buffs.go:178,369` | Aligned (latent, no output) | - | - |

## Raid buffs, debuffs and the preset

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **The BiS page said Trueshot Aura gives "50 at 60".** The hotfix push doubled rank 5 (20906) to 100 ranged attack power, and the engine has used 100 since E a3b2af6d4. The reason in `data/curated/presets.json` now says 100 and names the push. | SpellEffect 697683 (20906: aura 124, base 100). `TestRaidPresetReasonsMatchTheClient` gained the case and failed first ("still says 50 at 60"), then passed. | CONFIRMED-BUG (published text) | new-data (the push; round 13's hotfix lane changed the engine, not the reason) | S dafabde80 |
| P2 | Re-read on 70753 with the push:<br>- Gift of the Wild 21850: 385 armor, 16 stats, 27 resistances.<br>- Arcane Brilliance 23028: 31.<br>- Windfury Totem 10614 → 10612 (20% proc) → 10610: 246 attack power, 2 charges, 1 s.<br>- Judgement of the Crusader 20303: 161.<br>- Heart of the Lion 409583: 10% stats; 40 + 4 a level attack power and ranged attack power.<br>- Curse of the Elements 1311680: -75 resistance and +10% on misc 126 (all six magic schools, Holy included). The engine lists all six.<br>The buff ranks are pinned by `sim/leveling/buff_ranks_client_test.go` and `raid_ranks_client_test.go`. | client rows; `core/buff_ranks.go`, `core/buffs.go:1667-1745`, `core/debuffs.go:335-375` | Aligned | - | - |
| P3 | The preset's notes against its lists:<br>- one Blasted Lands buff per spec (category 103);<br>- one potion per spec;<br>- R.O.I.D.S. for strength users and tanks, Ground Scorpok Assay for rogues, hunters and the cat, Cerebral Cortex Compound for casters, Gizzard Gum for druid and priest healers;<br>- poisons and stones together under the totem aura.<br>All are consistent. | `sim/request/raid-consumes.json`, `presets.json` notes | Aligned | - | - |

## Consumables

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Every preset consumable read from ItemXItemEffect → ItemEffect → SpellEffect, against the engine:<br>- Flask of Supreme Power 150; Greater Arcane Elixir 35; Shadow Power 40; Fire Power 10; Frost Power 15; Holy Power 40 (Holy).<br>- Mongoose 25 Agility and 2% aura-290 crit; Juju Power 30; Juju Might 40/40.<br>- R.O.I.D.S., Ground Scorpok Assay, Cerebral Cortex Compound and Gizzard Gum 25, category 103.<br>- Elixir of Sages 25 Spirit and 2% crit; Greater Cleric's 40 healing; Greater Mageblood 20 mana per five; Distilled Wisdom 2000; Titans 1200; Rumsey Rum Black Label 15.<br>- Major Mana Potion 1800 ±25% (1350-2250); Mighty Rage Potion 600 ±25% (45-75 rage) plus 60 Strength for 20 s.<br>- Major Frenzy Potion 80/80 and Major Spellblasting 47 for 30 s. Their 120 s category sits on the spell (SpellCooldowns), not the item.<br>- Thistle Tea: 100 Energy, a 5 min item cooldown and the 1153 category.<br>All match. | client rows (`item.py`); `core/consumes.go:500-660, 1172-1240, 1397-1436` | Aligned | - | - |
| C2 | The "Endless" elixirs and scrolls the push added (ItemEffect 239511-239689, spells 1327736-1328127) are still on no item. No ItemXItemEffect row names them, and the item hotfix layer has no such item. | push-layered ItemEffect against ItemXItemEffect; `raw/hotfixes/ItemSparse.csv` | OPEN (settles when an item carries them; re-check the next push) | - | - |

## Item effects and set bonuses

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| I1 | The BiS files have not changed since 11-01, so round 12's census of item effects on the published gear still applies. I re-read every row the census marks "no engine effect":<br>- defensive or utility only: Ring of Saviors, Guardian Talisman, Arena Master, Arena Grand Master, Uther's Strength (a shield when struck), Petrolspill, Truefaith Vestments (Fade), Robe of the Void (pet heal), Lifeforce Dirk, Forsaken weapons, Evonice's, Parachute Cloak, Kam's, Savage Gladiator Chain (fear resist), Archmagister's Pendant (aura 447);<br>- thrown attacks only: the Satchels of Bombs (1318034 and 1318069 proc on thrown attacks, which a tank's rotation never makes).<br>Nothing missing moves output. Hand of Edward the Odd (Focus 18803) is a melee-hit proc. The casters who wear it at 60 never melee, so it rarely fires; that is correct. | `scratchpad/r12/itemscrit/census.txt`, re-read on the 70753 rows; `common/item_effects.go:1410` | Aligned | - | - |
| I2 | Set bonuses active in any published band: Devilsaur 2 (Survival, rogues, Arms), Ironfeather 2 (Balance, Elemental), Stormshroud 2 and 3 (rogues). Each is in `client_sets_gen.go` and pinned by the item_sets client tests (proc masks, amounts). No other set reaches two pieces in a published band. | ItemSparse ItemSet against ItemSetSpell for every band's slots | Aligned | - | - |

## Rank choices (non-healer rows)

Every non-healer row (61 band × race cells) was re-judged on its own against the default rule. Round 13's hotfix lane had re-measured the hunter rows only all together (its H8). Full table: `$R/data/rc.tsv`.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Beast Mastery's lower Arcane Shot ranks now lose or tie.** Rank 3 at 40 measured -0.27 to -0.33% (error 0.05). Rank 1 at 50 measured -0.03 to -0.10% (error 0.04). Rank 1 at bare 60 measured -0.70 to -0.78%. Every result held at both seeds and both factions. The push raised every Arcane Shot rank (rank 8 217 → 280), so the top rank now pays for its mana. The three rows are removed, and a Beast Mastery hunter casts the highest rank learned. | `rc.tsv`; test first: `TestRound14RankChoicesCastTheWinningRank` failed on all three, then passed | CONFIRMED-BUG | new-data (the push; round 13 H8 measured the rows only all together) | S e7a45f12e |
| K2 | **Shadow's Devouring Plague rank 2 at bare 60 loses for a Human and a Gnome.** Human -0.31/-0.32%, Gnome -0.15/-0.21% (error 0.07). Dwarf, Night Elf and Troll still gain +0.94 to +1.15%. The row now names only those three races. The Gnome is the published Alliance race, so that band's number moves. | `rc.tsv`; same test (Human and Gnome cases) | CONFIRMED-BUG | R13 175d4ff96 (Vampiric Embrace upkeep merged after ranks lane 2fbc44364 measured the row) | S e7a45f12e |
| K3 | **Destruction's Corruption rank 3 at 40 ties for every race:** Troll +0.03, Gnome -0.07, Human 0.00/+0.13, Orc -0.06, Undead +0.07/+0.02 (error 0.07). Round 13 measured it within twice the error and asked for a recheck at integration, which did not happen. The row is removed. | `rc.tsv` and the race runs; same test | CONFIRMED-BUG | R13 e88a33c97 (the random regen phase, merged after the row's measurement) | S e7a45f12e |
| K4 | Every other row still wins beyond error at both seeds, on each race it covers:<br>- Balance: Moonfire, Wrath and Insect Swarm at 30 (+2.5 to +7.0%).<br>- Beast Mastery Aimed Shot rank 1 at 40 and bare 60 (+2.8 to +3.8%); Marksmanship's at bare 60 (+4.3 to +4.5%).<br>- Arcane: Frostbolt at 30 (+0.55 to +1.0%, error 0.12).<br>- Fire: Fire Blast at 30 and 40 (+0.42 to +1.14%, error 0.14); Scorch at 30 to 50 (+1.2 to +6.5%).<br>- Protection paladin: Consecration at 40 to 60 (+1.9 to +4.0% threat); Holy Strike at 40 (+0.42 to +0.68%).<br>- Retribution: Consecration at 30 to 60 (+1.1 to +7.6%).<br>- Shadow: Shadow Word: Pain at 30; every Mind Blast race row (lowest: Night Elf at 50, +0.27/+0.37, error 0.11); Devouring Plague at 30 to 50; Shadow Word: Death at 50 and bare 60.<br>- Elemental: every row (lowest: Lightning Bolt rank 7 pacing at 50 and bare 60, +0.39 to +0.49%, error 0.09).<br>- Destruction: Shadow Bolt at 30 (+0.51 to +0.61%, error 0.06). | `rc.tsv` | Aligned | - | - |
| K5 | The mechanism: the embedded copy matches the curated file (`TestEmbeddedRankChoicesMatchCurated`), and the race key follows `-race`. The coach reads the same bytes (`api/internal/coach` passes), and so do the addon rotation tests (`test_addonrotation.py` passes). | tests named | Aligned | - | - |

## Trainables

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | A fresh `build_trainables` over the push-layered raw reproduces the committed `trainables/*.json` byte for byte for all nine classes. The landing's regen carries round 13's auto-learned rank 1s and its effect-row fix, and the push changes nothing in them. I re-read the module with fresh eyes:<br>- learnable methods 0 and 2;<br>- the unmasked method-2 rule;<br>- the effect-row rule;<br>- the no-learn-row copies;<br>- live-tree tagging.<br>No defect found. | scratch run into `$R/data/trn/`, compared file by file | Aligned | - | - |

## Leveling band walk

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| W1 | Round 13 hybrids X1 listed 43 walk-gap rows. Where each went:<br>- 39 were measured by the walk lane (10 adopted, the rest held) or by hybrids (the bear), or are owner calls (Protection paladin 20, ruling 1, lane tanks).<br>- Resto druid 38 Swiftmend and Resto shaman 38 Nature's Swiftness: rungs with no BiS entry. They follow their 30/40 decisions, which lost.<br>- Shadow Inner Focus at 20 to 50 was argued, not measured. This lane checked it: Inner Focus needs an 11th Discipline point, which a Shadow band can only take from Shadow. Every one of the 31 Shadow points is held by a tier gate: rows 0 to 4 hold exactly the 25 Darkness needs, and Darkness 5 makes Shadowform's 30. Spirit Tap 3→2 is refused ("Darkness needs 25, the build has 24"), and Darkness 5→4 is refused ("Shadowform needs 30, the build has 29"). At 20 to 40 the cost is Shadow points before Shadowform. No legal band build takes it before 60. | probe of `LadderTalentString` (priest-shadow, bands 20 to 60); spec-breakdown refusals on the two candidate strings | Aligned (no legal reservation) | - | - |
| W2 | **Demonology's Improved Imp at bands 20 and 30 (warlock lane D4, coordinator follow-up).** Every band summons the Succubus, so Improved Imp 3 does nothing. Moving those 3 points to Unholy Power on the walk's own band strings gains:<br>- band 20: +1.16/+1.16% horde, +1.13/+1.13% alliance (48.37 to 48.93 at seed 1, horde);<br>- band 30: +1.05/+1.05% horde, +1.05/+1.04% alliance (84.02 to 84.90).<br>Error is under 0.05%, so the swap passes the adoption test at both bands, both seeds and both factions (Shadow-build precedent: a dead point may be replaced). **It cannot be adopted through ladder-first.**<br>- A reservation is capped at the guide's target, and the leveling build's Unholy Power target is 2. Bands 20 and 30 already hold those 2.<br>- The walk only ever adds points toward the level-60 guide build, which keeps Improved Imp 3. So a band can never hold Unholy Power 5 with no Improved Imp and then gain Improved Imp at 40.<br>Adopting it takes one of two things. The guide's leveling `build:` could drop Improved Imp for Unholy Power 5, but that also moves bands 40 to 60 by +0.72 to +0.95% (warlock D4), which is under the 1% bar: an owner call on the precedent. Or the walk could gain a band cap that lets the point order take points back, which the addon's point order does not allow today. | `$R/data/sb` with `-talents` on the probed `LadderTalentString` strings (20: `-2332100...` vs `-2035100...`; 30: `-2332113101210001000-` vs `-2035113101210001000-`); `sim/leveling/ladder_first.go` reservedRanks (capped at `targets`) | OPEN (passes at 20/30; needs the guide build change, owner or warlock/guides lane) | old | - |

## Before/after

Engine `lane/r14-data` 6411690c9 with site `wt/r14-data` against `$R/sb-before`, 10,000 iterations. Seed 1 absolute values are shown first.

| spec, preset, band, race | seed 1 | seed 50001 |
|---|---|---|
| rogue-combat raid 60 gnome | 626.5 → 627.8 (+0.20%) | +0.33% |
| rogue-assassination raid 60 gnome | 625.9 → 629.3 (+0.54%) | +0.37% |
| rogue-subtlety raid 60 gnome | 539.3 → 540.4 (+0.21%) | +0.24% |
| warrior-fury raid 60 gnome | 754.0 → 755.3 (+0.18%) | +0.12% |
| warrior-arms raid 60 gnome | 717.7 → 718.0 (+0.04%) | +0.09% |
| warrior-protection raid 60 gnome (damage) | 390.4 → 391.0 (+0.17%) | +0.17% |
| hunter-beast-mastery bare 40 horde / alliance | 167.4 → 167.9 (+0.32%) / +0.27% | +0.33% / +0.33% |
| hunter-beast-mastery bare 50 horde / alliance | 229.8 → 229.9 (+0.03%) / +0.11% | +0.06% / +0.09% |
| hunter-beast-mastery bare 60 horde / alliance | 327.2 → 329.5 (+0.70%) / +0.77% | +0.74% / +0.78% |
| priest-shadow bare 60 alliance (gnome, published) | 349.4 → 349.9 (+0.15%) | +0.21% |
| priest-shadow bare 60 human | 351.3 → 352.4 (+0.31%) | +0.32% |
| warlock-destruction bare 40 horde / alliance | 147.2 → 147.1 (-0.03%) / +0.07% | -0.03% / +0.06% |

Unchanged to the decimal: raid 60 for every published entry I checked (rogue-combat, mage-arcane and priest-shadow alliance, hunter-beast-mastery horde), and priest-shadow bare 60 horde. No published entry is a Gnome warrior or rogue. The rank rows change only bare bands.

## Files to regenerate

- addon-data (`pipeline addon-data`, then Data.lua). The rank rows ship as `raceSpellId` lines: Beast Mastery Arcane Shot (rows gone), Devouring Plague at bare 60 (race list), Corruption at 40 (row gone).
- sim/request `TestRotationLadder` goldens: hunter-beast-mastery (40 to 59, bare 60), warlock-destruction (40 to 49), priest-shadow (bare 60 rung, if the ladder runs it for a Human or Gnome).
- Nightly at the refresh: BiS bare bands for hunter-beast-mastery 40/50/60, warlock-destruction 40 and priest-shadow 60 (alliance), plus their dps-per-point.
- No fork goldens moved (`go test ./core` green; no fork test sims a Gnome warrior or rogue).

## Measure at refresh

- Rogue and warrior race picks with Expansive Mind's 105 energy and rage. Gnome Combat now reads 627.8 against the published Night Elf's 627.1. That is within 1%, so the race pick stays under the rule; the re-rank will show it.
- Beast Mastery bare 40 to 60 and Shadow bare 60 (Alliance) picks on the new ranks.

## For other lanes

- **core:** `sim/core/rage.go` now carries a per-bar cap (`rageBar.maxRage`, `Unit.MaxRage()`, `MultiplyMaxRage`). `AddRage` caps at it; the `MaxRage` constant stays the default. Code that tests `core.MaxRage` directly, such as Mighty Rage Potion's use gate in `common/item_effects/client_procs.go:221`, still reads 100. That is harmless today. Keep it in mind if a merge touches rage.
- **tanks:** Protection paladin's Consecration rank 1 (40 to 60) and Holy Strike rank 1 (40) rows still win on today's tank boss (+1.9 to +4.0% and +0.4 to +0.7% threat). Re-check them once the boss gets attack power (ruling 4).
- **guides:** no guide or APL note names the removed ranks (checked hunter, priest and warlock prose).

## Still open

0. W2: Demonology Unholy Power for Improved Imp at bands 20 and 30 passes the adoption test, but only a change to the guide's leveling build can carry it (owner call: bands 40 to 60 gain under 1%).

1. C2: the "Endless" elixirs are on no item yet.
2. Kings with Heart of the Lion, and the Judgement of Wisdom/Light 50% roll (carried from round 13).
3. Round 13's in-game checks, pending the owner (listed at the top).

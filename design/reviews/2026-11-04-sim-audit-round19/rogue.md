# Sim-accuracy round 19: rogue

Lane `rogue`. Site wt/r19-rogue (from int/r19), fork lane/r19-rogue (from fork int/r19 01ecf70dc). Client 1.60.1.70755 with the late hotfix re-layer. Specs: Assassination, Combat, Subtlety.

Method: spec-breakdown on the guide raid builds (`$R/rt`) and frozen BiS (`$R/bis-frozen`), raid 60 and bare 20 to 50, and `-level` runs at 43, 47, 49, 53, 55 and 59 for all three specs. Client rows from the 70755 CSVs (SpellEffect, SpellPower, SpellCooldowns, SkillLineAbility, ItemSparse), `trainables/rogue.json`, `talents/rogue.json`, and Wowhead Forever tooltips. The lane's starting engine (int/r19) reproduces `$R/before.tsv` exactly for all twelve rogue rows, so no other lane's merge moved a rogue number.

Totals: **0 CONFIRMED-BUGs.** One card-text fix (guide round 15 ROGUE-R15-S4), wow-player SHIP on the rendered page. No DPS moves.

Commits: site d3146b361 (Sinister Strike and Backstab card notes, combat.md:77), plus this report. Fork fa10d43ab and bdb4a2a15 (notes sync only; `apl_check` ignores notes, so neither is needed for the pin).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K4 | Subtlety band 20: the Sinister Strike fallback holds to 60 energy on the published mace (Diamond Hammer) | Unchanged from round 18: -0.3/-0.5% on that one set. The hold exists so a dagger rogue never spends Backstab's energy on Sinister Strike between 45 and 59. Lifting it needs an APL value that reads the main-hand weapon type; `proto/apl.proto` has none (the nearest, `spellCurrentCost`, reads cost, not the weapon). A new APL value is shared core plus proto plus web card gates, too wide for 0.3% on one band that the BiS refresh fixes by giving band 20 a dagger | OPEN (modelling limit; fixes itself at refresh) | - | - |
| X1 | Occult Poison I/II learnability | Not learnable in Forever, on two independent community sources plus the client: warcraft.wiki.gg "Rogue abilities (Forever)" lists every trainer poison (Instant I-VI, Deadly I-IV, Wound I-IV, Crippling I-II, Mind-numbing I-III) and no Occult Poison; warcraft.wiki.gg "Occult Poison II" marks it Season of Discovery exclusive (taught by a handbook); Warcraft Tavern's Forever poisons guide lists the trainer poisons without it (all read 2026-11-04). In 70755 the craft spells (458822, 1214168) carry the same SkillLineAbility shape as SoD's Numbing, Atrophic and Sebacious poisons, and `trainables/rogue.json` files the poison ids under `class_spell`, the not-learnable bucket. The engine does not model it and no build, APL or guide names it | Aligned (closed) | - | - |
| G1 | ROGUE-R15-S4: the Combat Sinister Strike card described the APL's 60-energy gate ("held to 60 energy so a dagger rogue never spends Backstab's energy here"), telling a sword rogue to pool before every strike | Card note rewritten: "Sinister Strike only when Backstab cannot be cast, because your main hand is not a dagger or the target is facing you." The card cuts a note at its first colon (`clauseOf`, web/src/lib/guides/rotation-lists.ts:82), so the Assassination and Subtlety notes' colon form rendered only "Only when Backstab cannot be cast"; both now take the same comma sentence (their dagger hold stays as the second sentence). wow-player NO-SHIP on the first wording (it parsed as "neither a dagger nor facing"), then SHIP. Its SHOULD applied: the Backstab row and combat.md:77 now say "from behind". Notes and prose only | copy (not a sim bug) | - | S d3146b361, E fa10d43ab, bdb4a2a15 |
| R1 | Re-audit of round 18's Deadly Poison fix (E e10695cf2) | Every application re-takes `AttackerDamageMultiplier`; base damage is added only on a new stack (`makeDeadlyPoison` adds a stack only below 5, so a refresh at five never grows the base). Venom 59.8% uptime, Deadly Poison 46.8 DPS at raid 60. `TestDeadlyPoisonReapplicationTakesVenomOfTheMoment` passes on int/r19 | Aligned (no regression) | - | - |
| R2 | Re-audit of round 18's Hemorrhage fix (E 05ccb0d1d) | `dealBleedTickWith` scales `SnapshotAttackerMultiplier` for the tick and restores it, so Death Mark stores the scaled tick and the next cast's snapshot is untouched. Only this rogue's own Hemorrhage debuff counts (`hemorrhageAuras`). Bleed ticks still roll crit (owner ruling; Rupture/Garrote 33% at raid 60). `TestHemorrhageRaisesRuptureAndGarroteTicks` passes | Aligned (no regression) | - | - |
| R3 | Re-audit of round 18's Thistle Tea fix (E 1aa12229b, shared core) | `thistleTeaFits`: drunk when 100 energy fits under the bar's own cap with at most 10 lost. Raid 60: one tea a fight on all three specs (90 to 92.9 of 100 energy kept). `go test -tags with_db ./sim/core -run ThistleTea` passes | Aligned (no regression) | - | - |
| R4 | Rank rows at 40-59, level by level | The rogue has no rank-choice rows, and none could gain: every rogue ability costs the same energy at every rank, so the top learned rank always wins. The default rewrite was checked at 43, 47, 49, 53, 55 and 59 for all three specs (`-level`, per-spell ids in the breakdown): Backstab 11279@44, 11280@52; Eviscerate 8624@40, 11299@48, 11300@56; Garrote 11289@46, 11290@54; Rupture 11273@44, 11274@52; Shank 1329549@44; Death Mark 1325756@48, 1325757@56; Mutilate rank 2 @50; Deadly Poison III@46, IV@54; Instant Poison IV@44, V@52. Each matches `trainables/rogue.json` and was actually cast. The integrator's "any gain beyond error" correction changes nothing here: no rogue row was ever kept for being under 1% | Aligned | - | - |
| C0 | Cross-class coefficient lead (client EffectBonusCoefficient 0 vs a generator default) | No rogue cast scales with spell power. Poisons: 70755 SpellEffect gives EffectBonusCoefficient 0 and BonusCoefficientFromAP 0 on every Instant Poison (8680...11337), Deadly Poison (2818...11354) and Wound Poison row; the engine's poison spells set no `BonusCoefficient` (poisons.go:158-168, 247-257), so 0 on both sides. Weapon strikes (Backstab, Sinister Strike, Hemorrhage, Ghostly Strike, Mutilate, Shank, Ambush, Eviscerate, Kick, Riposte) use `BonusCoefficient: 1` on effects 31/121/3, which the client states as 1 (16511, 399960, 1241585/6). The only generator coefficient in a rogue file is `GougeSpellCoeff` (gouge.go:55), and Gouge is never cast. Nothing to settle on logs | Aligned | - | - |
| P1 | Shared watch: low ranks proc less | No rogue line casts a lower rank (R4) | Aligned (nothing to flag) | - | - |

## Assassination

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | Raid 60 breakdown, fresh | Mutilate 40.4 casts at 60 energy (2423.8), both hands 44% crit (Malice, Puncturing Wounds 15%), 0.2% miss; Eviscerate 35 energy at four and five; Rupture 25 at four and five with Blood Splatter (8.99 Garrote ticks = 27 s, 9.6 to 9.9 Rupture ticks a cast); Slice and Dice 25, 93% uptime; Garrote opener 40 (50 - 10 Poised, no Composure); Shank 18.24 casts (10 s cooldown), +20 net energy; Cold Blood 1.43; Relentless Strikes 570 energy; Seal Fate 27.6 points. Stormshroud 3 (client 23863: 4% on a melee hit, 30 energy; 361 energy a fight = 12 procs of about 300 hits) and Stormshroud 2 (18979: 10%, 20 Nature) match the client | Aligned | - | - |
| A2 | Bare 20: every point goes into Slice and Dice at four (6.83 casts, no Eviscerate) | Energy-starved Backstab at 20 builds one point every 5.8 s, slower than a four-point Slice and Dice (21 s) runs out, so the Eviscerate line never sees four points. A rotation choice, not an engine fault; not changed | Aligned (noted) | - | - |

## Combat

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | Raid 60 breakdown, fresh | Backstab 60 energy, 40.6 casts, 63.6% crit; Eviscerate at five, 20 energy (35 - 15 Flawless Execution); Slice and Dice at one or two, 73% uptime; Blade Flurry 25 energy, 2 min (client SpellPower 180126, SpellCooldowns 54910: 25, 120000), 2 casts; Adrenaline Rush no cost, 5 min (179497, 54829), 1 cast; Shank 8 s (Flash of Steel 5), 22.5 casts, +35 net energy; Puncturing Wounds 17.5 points (0.45 x 38.8 hits); Hack and Slash 4 = 4% dagger crit (talent text 4/5), sword/axe 1%/rank extra attack with a 200 ms ICD, mace 3%/rank armor | Aligned | - | - |
| C2 | Raid build dead points: Lightning Reflexes 4 | Still the gate for Flawless Execution (tier 3 needs 15 above it: Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3, Composure 2), as round 17 C2 found | Aligned | - | - |

## Subtlety

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| S1 | Raid 60 breakdown, fresh | Hemorrhage 30 energy (35 - 5 Efficient Strikes), 36.5 casts; Ghostly Strike 40 energy, 20 s cooldown (client 179423, 54818), 9.6 casts with Preparation; Garrote 8.9 casts at 30 (opener 20 with Poised), Initiative +1 point each; Death Mark 25 energy, 10 s (315056, 102201), 11.2 casts and 10.85 detonations; Premeditation 2 min (54909), 3 casts with Preparation; Preparation 10 min (54915), 1 cast; Shank 17.9 casts (10 s), +20 net, Boot Knife 2 points each; Rupture ticks crit 33% | Aligned | - | - |
| S2 | Bare 40 and the 43-59 levels | Death Mark rank by level (1316709 to 48, 1325756 to 56, 1325757 from 56); Garrote 8.15 casts and Rupture ahead of Slice and Dice, as round 18 S4 measured | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | Fork `sim/rogue` tests panic with "No item with id: 16707" | The item database is embedded only under the `with_db` build tag (sim/core/database_load.go). `GOWORK=off go test -tags with_db ./sim/rogue/` passes on int/r19. Without the tag every gear-set test panics, at 6ea675594 too. Not a code fault; the tag is the way to run them | Aligned (method note) | - | - |
| X2 | Live logs | `api.foreversixty.gg/v1/reports/recent` still lists only the 2026-10-20 report; foreverlogs.gg's public API answers "Forbidden". Last looked 19:13 UTC, before the 23:00 launch, so no level 40+ log exists yet | OPEN (K1 below) | - | - |

## Before/after

No change moves DPS (notes and prose only). The lane's engine on int/r19 reproduces `$R/before.tsv` exactly (10000 iterations, seeds 1 / 50001):

| spec | alliance | horde |
|---|---|---|
| rogue-assassination | 726.6 / 726.0 | 729.1 / 728.8 |
| rogue-combat | 774.0 / 774.2 | 776.6 / 776.6 |
| rogue-subtlety | 788.4 / 788.6 | 782.7 / 782.8 |

## Files to regenerate

- `addon-data.json` (Sinister Strike notes on all three rogues, Backstab note on Combat), then Data.lua. The rendered review used a temporary hand-patched copy, reverted before commit.

## Measure at refresh

- Subtlety band 20: a dagger main hand replaces Diamond Hammer (K4; carried from rounds 17 and 18).

## For other lanes

- Platform: an APL value for the main-hand weapon type would let the Sinister Strike fallback drop its 60-energy hold on a non-dagger (K4). Small gain, one band; only if the platform wants it.
- Anyone running fork package tests: use `-tags with_db` (X1).

## For the guide lane

- Nothing left over. This lane wrote the card notes and combat.md:77. `_prose.test.ts` and `rotation-lists.test.ts` pass (108 tests).

## Rendered review

wow-player on the rendered Combat Priority list (astro dev on this worktree after `npm run sync`, 1440 px, hand-patched addon-data): first wording NO-SHIP (MUST: "with no dagger in the main hand or the target facing you" reads as "neither"); after the rewrite, **SHIP**. Optional SHOULD left: row 6 wraps under the name while rows 1 to 5 sit inline.

## Sources

- warcraft.wiki.gg, "Rogue abilities (Forever)": https://warcraft.wiki.gg/wiki/Rogue_abilities_(Forever) (read 2026-11-04)
- warcraft.wiki.gg, "Occult Poison II": https://warcraft.wiki.gg/wiki/Occult_Poison_II (search result, 2026-11-04)
- Warcraft Tavern, "Rogue Poisons Guide" (Forever): https://www.warcrafttavern.com/forever/guides/poisons/ (read 2026-11-04)
- Wowhead Forever tooltips 458822, 1214168, items 226374, 234444 (nether.wowhead.com/forever/tooltip, 2026-11-04): the recipe and item exist in the client; they say nothing about a trainer

## Still open

- K1: Death Mark storage and the Mutilate/Shank off-hand penalty, until a level 40+ Forever log exists (Mutilate off-hand stays pending with the owner).
- K4: Subtlety band 20's 60-energy Sinister Strike hold on a mace, until the refresh moves the set to a dagger.
- Carried: Combat Eviscerate at five (owner: keep five).

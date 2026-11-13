# Engine follow, client 1.60.1.70938

Fork pull request #9 (`engine: follow client 1.60.1.70938`, from engine.yml run 43400492893) reviewed, finished and merged. Fork `master` is 10cbe9f5e. The site pins 10cbe9f5e.

**Result: 0 engine behaviour changes, 1 conformance report fix, 7 hand rows settled (4 now read the client, 3 kept), no published number moves.**

## Talent texts

The follow flagged six changed texts. None changes anything the engine models.

| Talent | What changed | Engine |
|---|---|---|
| Swiftmend (druid 104912) | "of the consumed spell" dropped | The engine never consumed the effect: 55 of 57 logged effects ticked on (swiftmend.go). The text now agrees. Comment updated. |
| Mind Flay (priest 105826) | a double space | Wording only. |
| Death Mark (rogue 110866) | `$mcrit` rendered as "Critical Strike chance" | Wording only. Same formula. |
| Wrack (warlock 105909) | 86 a tick and +10% to the warlock's DoTs became 60 a tick, and Siphon Life, Corruption and Bane of Agony tick 100% faster while channelling | Already in the engine from push 113766 (wrack.go: 25.2 + 1.75 a level = 60.2 at 60, 0.2002, clock rate 2). The text is catching up with the push. |
| Soul Siphon (warlock 105911) | "other Affliction effects" became "Corruption, Siphon Life, and Bane of Agony" | Already in the engine (`soulSiphonCountedSpellCodes`, 4/8/12% a rank, at most 3). |
| Gore Drinker (warrior 113569) | Reckless Swings added to its triggers | Not modelled. It heals the warrior, and no simmed number reads warrior self-healing. No published build takes it. Quoted text updated. |

The trees' rank counts, spell ids and rates did not change (`talents_auto_gen.go` moved only its build stamp).

## Hunter conformance: the 7 mismatches

All seven were one row: Explosive Trap, at 38, 40, 50 (ranks 1 and 2) and 60 (ranks 1 to 3). The client's `duration_ms` 60000 was set against the engine's 20000.

- **Decision: report-side fix (the engine's 20 s is the client's).** A trap's own `duration_ms` is how long it stays armed on the ground. The comparable number is its payload's. Since 1.60.1.70790 the client tables name the payload (13812, 14314, 14315, a 20 s burn) "Explosive Trap" instead of "Explosive Trap Effect". 70291 had renamed Freezing and Immolation Trap's payloads the same way. `durationPayloadSpells` still looked the payload up by its old name, so it found nothing and compared the armed lifetime. It now looks up the same-rank "Explosive Trap" with another id, as it does for the other two traps. All 7 rows now match (hunter golden regenerated). These rows were mismatches on `master` before the follow as well. None was caused by 70938.
- Hunter mismatches: 7 to 0. No client row needed an engine change, no hand row is involved, and no allowance was added.

## Hand rows the client move listed as native

Each was read against the 70938 raw tables (`SpellEffect`, `SpellAuraOptions`, `ItemEffect`) and the regenerated fork tables.

| Ledger | Row | 70938 native value | Decision |
|---|---|---|---|
| MR58 | Wild Growth tick coefficient 0.066 (408120, 1238214, 1238215) | 0.066 in the generated and vendored rows | **Removed.** `WildGrowthTickCoefficient`, `wild_growth_client_test.go` and the allowances in `healing_client_test.go`, `restoration/spells_test.go` and `conformance/damage.go` (that map is now empty, so it is gone) are deleted. The table carries the literal 0.066, checked against the client as every druid heal is. |
| MR41 | Imp Firebolt ranks 1 to 4 (681913, 685570 to 685572) | 6/11/15/18 centres, the push's widths and growth | **Removed.** `FireboltDamage` is `clientdamage.FromTable` over the generated columns. Ranks 5 to 7 are unchanged. |
| MR41 | Wrack tick (1349556) | EffectBasePointsF 25.2, 1.75 a level, 0.2002 | **Kept, amount only.** The generated amount reads 25 because the site's spellconst rounds EffectBasePointsF to a whole number (`pipeline/spelltext.py` `_base_points`). Growth and coefficient match. A new test (`TestWrackTickAgreesWithTheGeneratedRow`) checks them against the generated row. |
| MR92 | Bane of Havoc global (SpellCooldowns 102518) | gcd_ms 1500 on 1225228 | **Native.** The conformance report compares the global. Comment only. |
| MR83 | Defensive Stance threat (686475) | 80 | **Native.** The engine's 0.8 is already checked against the row (`TestShieldAbilitiesMatchClient`). No conformance server-hotfix entry existed for it, so nothing drops. **Defiance 12792 still reads 15 in 70938, so its 7/13/20 stays on push 113326.** Bear Form's 700189 also reads 115 natively. Its stale test comment is fixed. |
| MR98 | Gloves of Holy Might (ItemEffect 238916 to 1330855) | TriggerType 1; ProcChance 15, ProcCategoryRecovery 20000, ProcTypeMask 69636 | **Native and equal.** It stays a coded item effect, as every client proc in `client_procs.go` is. No generated proc table exists to read. |
| MR46 | Arcane Shot and Summon Hawk cost | ManaCostPct 11 on every rank | **Removed.** Both read `ArcaneShotManaCostPct[rank]` and `SummonHawkManaCostPct[rank]`. The hand constant is gone. |
| MR46 | Nature's Swiftness mask (695239) | mask_1 0x6000000 (Wild Growth) | **Native.** The engine's class mask is its own spell-mask enum, with no generated reader. Comment only. |
| MR46 | Lightning and Chain Lightning overload rows (408439 to 408484) | half the bolt: 36.5, 54.5, 147.5, 176.5; chain 107.5 | **Kept.** 70938's vendored spellconst carries them but rounds the half points (36.5 reads 36, 54.5 reads 54, 147.5 reads 148, 176.5 reads 176, 107.5 reads 108), past the 0.05 roll tolerance. `TestLightningOverloadRowsAreThe70888Client` keeps pinning the exact rows. |

Unchanged: the MR46 server-side values (Stormstrike 8 s and its conformance server-hotfix entry, Mental Dexterity, Maelstrom Weapon, the imbue and totem rule, Arcane Concentration's split, the hawk's swing). No 70938 client or push row states them.

### Follow-up (not fixed here)

`pipeline/spelltext.py` `_base_points` rounds EffectBasePointsF to an integer. The rounding was designed for tooltips, but spellconst inherits it. In 70938, 107 class-file spells carry a fractional amount. Most are NPC copies. The player rows it bends are Wrack (25.2), the Lightning and Chain Lightning overloads, Health Funnel's effect 3, Rupture 1271514 (8.70) and Judgement of Fury (194.4, 316.8). The engine states Wrack and the overloads exactly. A spellconst `amount` that keeps the float would let those two hand rows go. That change touches every generated table, so it is a data-lane change of its own.

## Fork commits on the follow branch

- 9df0365d4 fix(conformance): Explosive Trap's duration reads its renamed payload
- 363c1c30c refactor(engine): the hand rows client 1.60.1.70938 now carries read the client
- c4b3e448b test(conformance): hunter golden, Explosive Trap's seven rows match

The fork CI passed on c4b3e448b, which was merged as 10cbe9f5e. The touched packages (conformance, druid, hunter, warlock, warrior, shaman, common) and the golden regeneration ran on the Windows box.

## The pin's deltas, raid 60

Measured with `engine-70938/deltas.sh`: each spec's published raid 60 alliance entry, 10,000 iterations, seed 1, one sim at a time on the box. The baseline is main's pin at measurement time (a0dbf6d8c, the Feral lane's Rake fix), not b0a43f004. DPS for damage specs and the heal fight's effective HPS for healers.

| Spec | a0dbf6d8c | 10cbe9f5e | Change |
|---|---|---|---|
| druid-balance | 737.5 | 737.5 | 0 |
| druid-feral-bear | 498.6 | 498.6 | 0 |
| druid-feral | 1082.1 | 1082.1 | 0 |
| druid-restoration | 1025.53 | 1025.53 | 0 |
| hunter-beast-mastery | 971.2 | 971.2 | 0 |
| hunter-marksmanship | 1086.0 | 1086.0 | 0 |
| hunter-survival | 976.1 | 976.1 | 0 |
| mage-arcane | 794.5 | 794.5 | 0 |
| mage-fire | 812.8 | 812.8 | 0 |
| mage-frost | 696.1 | 696.1 | 0 |
| paladin-holy | 719.84 | 719.84 | 0 |
| paladin-protection | 496.8 | 496.8 | 0 |
| paladin-retribution | 833.4 | 833.4 | 0 |
| priest-discipline | 623.20 | 623.20 | 0 |
| priest-holy | 880.28 | 880.28 | 0 |
| priest-shadow | 797.0 | 797.0 | 0 |
| rogue-assassination | 914.8 | 914.8 | 0 |
| rogue-combat | 902.3 | 902.3 | 0 |
| rogue-subtlety | 890.5 | 890.5 | 0 |
| shaman-elemental | 748.0 | 748.0 | 0 |
| shaman-enhancement | 873.4 | 873.4 | 0 |
| shaman-restoration | 890.00 | 890.00 | 0 |
| warlock-affliction | 768.7 | 768.7 | 0 |
| warlock-demonology | 814.1 | 814.1 | 0 |
| warlock-destruction | 888.0 | 888.0 | 0 |
| warrior-arms | 933.4 | 933.4 | 0 |
| warrior-fury | 1007.6 | 1007.6 | 0 |
| warrior-protection | 423.5 | 423.5 | 0 |

Every spec is identical to the decimal. Every behaviour 70938 brings was already in the engine as hand rows. The regenerated base stats and spell tables change nothing a raid 60 entry reads. Replacing hand rows with client reads is value-for-value. The BiS ranking this pin starts should move nothing from the engine. Any change it shows comes from elsewhere (the data lane, or the Feral lane's earlier Rake pin).

## Gates

- Fork: CI `test` green on the PR head. Conformance goldens regenerated on the box, where only `hunter.golden.md` moved.
- Site: `make apl-check` against 10cbe9f5e passed. `data/tests/test_genproto.py` passed. The sim module's vet and tests ran on the box with the pin (see the landing commit).

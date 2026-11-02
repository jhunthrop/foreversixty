# Sim audit round 9: crit lane

The owner's item: "The engine's single crit stat: a sharpening stone must not raise spell crit. Separate melee and spell crit where they differ." Round 8 left it open as "one crit stat for melee and spell".

**What the engine did.** At fork 40021374a one `stats.Crit` served every roll: `PhysicalCritChance`, `SpellCritChance` and `HealingCritChance` (`sim/core/spell_result.go:123, 224, 253`). Agility and Intellect both fed it (`base_stats.go:231-232`). Each class's base crit was the larger of Era's melee and spell values. The results:
- a sharpening stone raised spell crit;
- a wizard oil raised melee crit;
- Agility raised a caster's spell crit;
- Intellect raised a hybrid's melee crit and a hunter's shots.

**What the client says.** The 1.60.1.70707 tables (hotfix push layered in) keep three crit auras apart, and Forever's own new rows name them:
- aura 52 is melee and ranged (weapon) crit: "Increased Critical 1 - Melee and Ranged" 1318954;
- aura 57 is spell crit: "Increased Critical 1 - Spell" 1318956;
- aura 290 is every table: "Increased Critical 1 - All" 436239 (also 448324 and 1223262).

`SpellCategories.DefenseType` decides which table a spell rolls. 1 (magic) rolls spell crit; 2 (melee) and 3 (ranged) roll melee and ranged crit.
- DefenseType 3: Serpent Sting 13552, Arcane Shot 14287, Hammer of Wrath 24239.
- DefenseType 1: wand Shoot 5019, the rogue poisons 8680 and 2818, the weapon procs 18980 and 21992, Consecration 20924.
- DefenseType 2: the seals, e.g. 25713.

An item's crit rating (ItemSparse stat 32) costs 14 points per 1% in the melee, ranged and spell columns of `combatratings.txt`, so it reaches every table. The items lane counted 978 such rows, and none uses stat 19, 20 or 21.

**What "unified crit" left open.** Blizzard says "spell, melee and ranged critical chance are combined". It does not say whether Agility and Intellect still feed separate pools (GAP 9 in `design/reviews/2026-10-08-beta-evidence.md`). The beta logs settle it. Source: public foreverlogs.gg reports at levels 13 to 30, fetched 2026-11-02; scripts `scratchpad/r9/crit/py/pois.py` and `pools.py`.
- **Rogue poisons.** Instant Poison, a spell-table hit, crit 9 of 409 times (2.2%). The same rogues' melee crit was 19% (961 of 5003 swings). One pool would have given about 78 poison crits.
- **Paladin Consecration.** One paladin's Consecration ticks crit 6.1% of 1762 (report 2705, character 48035). Its own melee crit was 12.4%.
- **Every class with a gear snapshot.** For 342 characters, the logs carry a gear snapshot with Agility, Intellect and talent points (level = points + 9). For each one I compared melee-minus-spell crit with each model's prediction:
  - Split model: χ² 3,696. Pooled model: χ² 10,751.
  - Rogues (50): observed gap +18.0 points. Split predicts +12.8, pooled 0.
  - Warriors (33), paladins (96), shamans (68) and druids (26) also favour the split.
  - Priests and mages lean the other way, which their spell-only crit talents explain.
- **Hunters, absolute.** Seven hunters' ranged crit averaged 6.3% in the logs. Agility alone on base 0 predicts 5.7%. The old engine (Era's spell base 3.6 plus Intellect) predicted 11.0%.

So Forever unified item, talent and buff crit, not the tables. The fix keeps one path where the client has one (aura 290 and item crit) and splits only where it differs.

## The change (engine, lane/r9-crit)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | **One crit stat for every table.** Agility raised spell crit, Intellect melee crit, and Era's spell base sat on melee and ranged (hunter +3.6, paladin +3.5 instead of 0.7). The engine now has three crit stats:<br>- `stats.Crit` reaches every table: item crit rating and the aura 290 sources.<br>- New `stats.MeleeCrit` (melee and ranged) and `stats.SpellCrit`, appended to the Stat enum and proto as `StatMeleeCrit` 41 and `StatSpellCrit` 42.<br>Agility feeds MeleeCrit, Intellect SpellCrit, and `ClassBaseCrit` keeps Era's two columns. A physical roll reads Crit + MeleeCrit. A spell or heal reads Crit + SpellCrit. A magic-school spell with a melee or ranged DefenseType (Serpent Sting, the seals, Hammer of Wrath) reads the melee table. | The client's aura rows and DefenseType, above. The logs, above. Tests `TestIntellectIsSpellCritOnly`, `TestMongooseCritReachesEveryTableItsAgilityOnlyMelee` and `TestRogueAgilityDoesNotReachPoisons` (`sim/core/crit_tables_test.go`) failed first. | CONFIRMED-BUG (client tables plus two kinds of log evidence) | E c14bf02a5 |
| C2 | **Elemental Sharpening Stone raised spell crit.** It is now melee crit only. The existing −2 on ranged spells stays. | Enchant 2506 → spell 22755, aura 52, 2. Test `TestSharpeningStoneIsMeleeCritOnly` failed first: +2 on spell and heal. | CONFIRMED-BUG | E c14bf02a5 |
| C3 | **Brilliant Wizard Oil raised melee and ranged crit.** It is now spell crit only. | Enchant 2628 → 25113: aura 13 and aura 135 at 36, aura 57 at 1. Test `TestBrilliantWizardOilIsSpellCritOnly` failed first. | CONFIRMED-BUG | E c14bf02a5 |
| C4 | **Cruelty reached weapon procs.** It is now melee crit. | 12320, aura 52, 5. Its SpellEquippedItems row lists melee weapons only (physical lane note). | CONFIRMED-BUG (warriors move by 0.07% at most) | E c14bf02a5 |
| C5 | **Elemental Devastation.** The proc added unified crit and then took it back off every spell school. It is now melee crit only. Same result, simpler code. | 30165 / 29177 / 29178, aura 52. `TestElementalDevastationIsMeleeOnly` was updated to read MeleeCrit. | Aligned (same numbers) | E c14bf02a5 |
| C6 | **Spell-only crit on every table.** These now give spell crit only:<br>- Atiesh's Power of the Guardian (28142, aura 57, 2)<br>- Cenarion Raiment 5-piece (18382, aura 57, 2)<br>- Bloodvine Garb 3-piece (18382, aura 57, 2)<br>None is in a published band. | ItemSetSpell and SpellEffect rows. | CONFIRMED-BUG (no published number) | E c14bf02a5 |
| C7 | **Firestone lacked its spell crit.** It now has it: 1, 1, 2 and 2% by rank. No preset uses Firestone. | Enchants 1803, 1823, 1824 and 1825 → 23480 to 23483: aura 13 fire damage 10/14/17/21 and aura 57 at 1/1/2/2. | CONFIRMED-BUG (no published number) | E c14bf02a5 |
| C8 | **Client set-bonus decoder.** It folded 52, 57 and 290 into one max(). It now writes each aura to its own table, and a spell stating 52 and 57 alike gives every-table crit. Moves: Black Dragon Mail 3-piece (7598, aura 52) is now melee only; The Ten Storms 5-piece (23570, aura 57) spell only. | `client_sets.go`. Test `TestClientFlatBonusKeepsCritTablesApart`. The conformance golden notes Ten Storms as "client SpellCrit 3". | CONFIRMED-BUG (neither set is in a published band) | E c14bf02a5 |
| C9 | **Pets.** Hunter pets inherit the every-table crit and the hunter's melee and ranged crit (415429's aura 52 channel), not its spell-only crit. Demons' Agility and Intellect feed their own tables. | `hunter/pet.go`, `warlock/pet.go`. | Aligned with C1 | E c14bf02a5 |

Goldens regenerated for the packages the split moves: mage, paladin protection, both shaman specs, both warrior specs, feral tank, and the set conformance golden.

## The change (site, wt/r9-crit)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| S1 | **simdb dropped equip-spell spell crit and all-table crit, and gave aura 52 to every table.** `equip.IGNORED_AURAS` held 57 and 290. Examples:<br>- Power of the Scourge (29468), Magnificent Trollshine and Enchanted Repellent lost their 1% spell crit.<br>- "+2% Crit" enchant 8720 (aura 290) lost its 2%.<br>- Might of the Scourge (29482, aura 52) and the "Critical Hit +N%" enchants (195 to 198, 2506) raised spell crit.<br>Now 52 is `melee_crit`, 57 (all schools) is `spell_crit`, and 290 is `crit`. A spell stating both 52 and 57 gives their shared part as `crit` (Forever's Might of the Scourge 1220743). | SpellItemEnchantment → SpellEffect rows in the layered 70707 tables. Tests `test_the_three_crit_auras_keep_their_tables` and `test_each_crit_table_has_its_own_stat` failed first. | CONFIRMED-BUG (no published preset carries a permanent enchant; the committed simdb.bin is untouched) | S 67f5e9d9b |
| S2 | Item crit rating (stat 32) and Wowhead's `critstrkpct` stay `crit`, every table. That matches the client's equal rating columns. | `combatratings.txt` level 60: Crit - Melee = Crit - Ranged = Crit - Spell = 14. | Aligned | - |
| S3 | spec-breakdown's heal and tank tables now print the table-only crit beside the every-table crit. | `cmd/spec-breakdown/heal.go`, `tank.go`. | tooling | S 67f5e9d9b |

The vendored `data/proto/common.proto` and `common_pb2.py` follow the lane engine. `data/proto/ENGINE_SHA` stays 40021374a for the integrator to move.

## Sources checked, by table (client 70707)

| source | client row | table | engine after |
|---|---|---|---|
| Elixir of the Mongoose | 17538: aura 29 Agility 25, aura 290 2 | all (its Agility melee only) | Crit 2 + Agility |
| Grilled Squid (Well Fed) | 1249523: aura 290 1 | all | Crit |
| Elixir of Sages | 1250941: aura 290 2 | all | Crit |
| Juju Power, Juju Might, R.O.I.D.S., Dense stone, other foods | no crit effect | - | - |
| Leader of the Pack | 24932: aura 290 3. Era's was melee and ranged only. | all (a Forever change) | Crit |
| Moonkin Aura | 24907: aura 290 3. Era's was spell only. | all (a Forever change) | Crit |
| Songflower, Rallying Cry, Slip'kik's | Forever rows are aura 4 (dummy); tooltips say all spells and attacks | all | Crit (gated off: no world buffs) |
| Ruthless (Human), Elune's Light (Night Elf) | 20597 and 1259799: aura 290 | all | Crit |
| Conviction, Malice, Thundering Strikes, Lethal Attacks, Killer Instinct, Nature's Majesty, Sharpened Claws, Arcane Instability, Weaponmaster (axe/polearm), Berserker Stance, Recklessness, Felstriker | 20117, 14138, 16255, 19426, 19370, 1223082, 16942, 15058, 12700, 7381, 1719, 16551: all aura 290 | all | Crit |
| Cruelty | 12320: aura 52 | melee | MeleeCrit (C4) |
| Elemental Devastation | 30165: aura 52 | melee | MeleeCrit (C5) |
| Elemental Sharpening Stone | 22755: aura 52 | melee | MeleeCrit (C2) |
| Brilliant Wizard Oil | 25113: aura 57 | spell | SpellCrit (C3) |
| Item crit rating | ItemSparse stat 32 | all | Crit (S2) |
| Agility / Intellect | PlayerExpectedStat CritPerAgility / SpellCritPerIntellect | melee / spell | MeleeCrit / SpellCrit (C1) |

## Before/after

Band 60 raid. Each cell is the change in the headline metric: DPS, effective HPS for healers, threat per second for tanks. 10000 iterations. Before: fork 40021374a with site origin/main 55a1212f7. After: E c14bf02a5 with S 67f5e9d9b.

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | horde s1 absolute (before -> after) |
|---|---|---|---|---|---|
| druid-balance | -3.96% | -4.02% | -3.55% | -3.60% | 579.5 -> 558.9 |
| druid-feral-bear | -4.18% | -4.22% | -4.12% | -4.10% | 856.3 -> 821.0 |
| druid-feral | -3.03% | -3.03% | -2.99% | -3.02% | 608.9 -> 590.7 |
| druid-restoration | -2.09% | -1.99% | -1.83% | -1.80% | 777.3 -> 763.1 |
| hunter-beast-mastery | -7.20% | -7.14% | -7.13% | -7.14% | 661.8 -> 614.6 |
| hunter-marksmanship | -7.36% | -7.37% | -7.30% | -7.30% | 653.1 -> 605.4 |
| hunter-survival | -7.10% | -7.13% | -7.04% | -7.09% | 759.5 -> 706.0 |
| mage-arcane | -2.36% | -2.32% | -2.08% | -2.06% | 601.4 -> 588.9 |
| mage-fire | -3.20% | -3.22% | -2.83% | -2.96% | 670.8 -> 651.8 |
| mage-frost | -2.17% | -2.15% | -2.11% | -2.11% | 601.5 -> 588.8 |
| paladin-holy | -4.09% | -4.16% | -4.29% | -4.33% | 648.2 -> 620.4 |
| paladin-protection | -4.35% | -4.41% | -4.29% | -4.32% | 1018.2 -> 974.5 |
| paladin-retribution | -5.70% | -5.71% | -5.69% | -5.71% | 658.8 -> 621.3 |
| priest-discipline | -1.86% | -1.89% | -1.58% | -1.68% | 620.0 -> 610.2 |
| priest-holy | -1.30% | -1.37% | -1.32% | -1.29% | 674.1 -> 665.2 |
| priest-shadow | -2.69% | -2.69% | -2.43% | -2.42% | 724.2 -> 706.6 |
| rogue-assassination | -0.59% | -0.59% | -0.61% | -0.61% | 609.3 -> 605.6 |
| rogue-combat | -0.38% | -0.38% | -0.38% | -0.38% | 606.0 -> 603.7 |
| rogue-subtlety | -0.60% | -0.60% | -0.62% | -0.62% | 532.6 -> 529.3 |
| shaman-elemental | -2.95% | -2.97% | -3.02% | -3.02% | 476.8 -> 462.4 |
| shaman-enhancement | -6.52% | -6.67% | -6.36% | -6.50% | 594.0 -> 556.2 |
| shaman-restoration | -2.82% | -2.82% | -2.85% | -2.81% | 838.9 -> 815.0 |
| warlock-affliction | -3.26% | -3.27% | -3.21% | -3.21% | 585.2 -> 566.4 |
| warlock-demonology | -3.08% | -3.09% | -3.01% | -3.06% | 654.9 -> 635.2 |
| warlock-destruction | -2.77% | -2.78% | -2.76% | -2.75% | 649.7 -> 631.8 |
| warrior-arms | +0.00% | +0.00% | +0.00% | +0.00% | 720.8 -> 720.8 |
| warrior-fury | -0.07% | -0.05% | +0.00% | +0.00% | 757.8 -> 757.8 |
| warrior-protection | -0.20% | -0.04% | -0.08% | -0.07% | 1036.9 -> 1036.1 |

Where it comes from, read off the per-spell crit (band 60 raid, horde, seed 1):
- **Hunters, −7.1 to −7.4%.** Ranged crit falls about 8 points, e.g. BM Auto Shot 27.8 → 19.9%. The 3.6 spell base and the Intellect crit come off every shot. The pet follows, because it copies the hunter's crit.
- **Enhancement, −6.4 to −6.7%.** Melee 35.8 → 28.9% (no Intellect crit, base 1.7 not 2.3). Earth Shock 32.9 → 24.1% (no Agility crit).
- **Retribution, −5.7%.** Melee 29.9 → 23.1%, and the seals and judgements (melee table) follow. Consecration (spell table) falls 34.8 → 25.6%.
- **Protection paladin and Feral bear, −4.1 to −4.4% threat.** Feral cat, −3.0%. The Intellect crit and the spell base come off their swings.
- **Casters, −2.1 to −4.0%; healers, −1.3 to −4.3%.** Spells and heals lose Agility crit, mostly the raid's Agility buffs, and the old max() base. Holy paladin loses the most (−4.3%).
- **Rogues, −0.4 to −0.6%.** Melee is unchanged. Poisons and Lightning procs fall to the spell table (Combat's Deadly Poison 40.0 → 27.2%).
- **Warriors, ≤0.2%.** Cruelty and the stone leave their weapon procs.

Bands 20 to 50 (the published bare preset, horde, both seeds):

| spec | 20 | 30 | 40 | 50 | band 40 seed 1 absolute |
|---|---|---|---|---|---|
| shaman-enhancement | -4.14% / -4.42% | -6.67% / -6.67% | -7.26% / -7.41% | -6.99% / -6.98% | 123.9 -> 114.9 |
| paladin-retribution | -5.11% / -5.35% | -4.50% / -4.64% | -6.84% / -6.93% | -8.98% / -8.97% | 108.2 -> 100.8 |
| hunter-survival | -6.72% / -6.87% | -9.72% / -9.72% | -6.92% / -6.87% | -8.08% / -8.08% | 173.3 -> 161.3 |
| rogue-combat | -0.25% / -0.25% | -0.61% / -0.61% | -0.55% / -0.55% | -0.71% / -0.64% | 109.8 -> 109.2 |
| druid-balance | -1.55% / -1.55% | -2.29% / -2.29% | -2.95% / -2.94% | -3.02% / -2.97% | 156.1 -> 151.5 |
| warlock-demonology | -1.45% / -1.45% | -1.24% / -1.12% | -1.20% / -1.13% | -1.15% / -1.15% | 142.0 -> 140.3 |
| mage-fire | -1.99% / -1.98% | -1.69% / -1.69% | -1.70% / -1.70% | -1.76% / -1.92% | 147.0 -> 144.5 |

Survival and Retribution move most, as at 60: Survival −9.7% at 30 and Retribution −9.0% at 50. The bare preset carries no stones or oils, so these moves are all from the pools and the base columns.

## For other lanes

- **Warlock:**
  - **Lash of Pain crit drops with this change.** At band 30 bare it goes from 6.3% to 2.7%; the 29 low-level logs show 7.1%. The Succubus's Agility no longer reaches its spell. Warlock Pet Scaling 416189 carries an aura 57 channel next to its aura 52, so the demon probably inherits the warlock's spell crit for its spells. Whether to wire that is the warlock lane's call. It is under 0.1% of DPS.
  - **Spellstone is wrong.** In Forever it is shadow damage 14/17/21 and cast speed 1/2/2% (enchants 8059 to 8061 → 1237162/4/5: aura 13 and aura 65), not the engine's +1% crit (`warlock/talents.go` applyWeaponImbue). No preset uses it.
- **Items:** Blackhand's Breadth 13965 and Eye of the Beast 13968 carry Forever equip spells: 1318954 is 1% melee and ranged, 1318956 is 1% spell. simdb gives each 2% every-table crit from the Wowhead/classic-db supplement, because neither has a client ItemSparse row. Neither is picked. Sanctified Orb 20512 is a mana item in Forever (24865: restores 319 to 361 mana), not the engine's 3% crit (`paladin/items.go:63`). It is not in a phase.
- **Integrator:**
  - Every spec moves (table above). Ladders, BiS picks and stat weights will shift at the next data run. The "crit" stat weight still means every-table crit, so a crit item's weight stays correct.
  - At the pin, `python -m pipeline simproto --engine <fork>` regenerates the same `common.proto` and `common_pb2.py` this branch carries. Moving `ENGINE_SHA` is your step.
  - Then rebuild simdb.bin; only the enchant rows move.
  - The fork's `ui/core/proto` TypeScript is generated and gitignored; the site does not use it.

## Still open

- **Era's base crit columns.** They are kept as Era states them: melee 0.7 paladin, 1.7 shaman, 0.9 druid, 0 hunter; spell 3.5, 2.3, 1.8, 3.6. Forever's chancetomeleecrit and chancetospellcrit game tables 404. The hunter logs fit base 0 (above). The melee column for priest, mage and warlock (3.0, 3.2, 2.0, equal to Dodge) only touches their melee swings, which no published rotation uses. A level-60 character sheet with gear off would settle all nine.
- **Which hands an aura-52 weapon enchant reaches.** The stone's spell has no SpellEquippedItems row. The engine gives melee crit to both hands and takes 2 back off ranged. Single source either way; unchanged.
- **Hunter pet and the stone.** The pet inherits MeleeCrit, so a Survival hunter's stone reaches the pet's melee. The hunter's own ranged shots take the 2 back off; the pet does not. No published hunter preset carries a stone.
- **Equip-spell crit on the Wowhead and classic-db supplement items.** classic-db's own aura 52 and 71 still map to every-table `crit` on the planner path. Two items are affected (above), neither picked.

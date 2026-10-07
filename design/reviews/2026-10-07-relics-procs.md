# Relics and weapon procs in the ranker, 2026-10-07

Lane `relics-procs`. Fork branch `relics-procs` (wowsims-forever, from `forever` at 42b2acbbc), site branch `relics-procs` (from `main` at b8cbd556). Client build 1.60.1.70009.

## Summary

1. The engine now models the DPS librams, totems and idols as equip effects, so the ranker sims them. Retribution gets Libram of Invocation, Elemental gets Totem of Thunder, and Enhancement's relic slot is empty with `no_dps_value` because its ranked rotation casts nothing a totem touches. None of the three is empty with `effect_not_modelled` any more.
2. Of the six weapons the brief named, only Blackblade of Shahram and Teebu's Blazing Longsword have a proc, and both were already modelled. Shadowsong's Sorrow, Ravencrest's Legacy, Riphook and Crackling Staff carry no proc in the client's item-effect rows or in Wowhead's tooltip; there was nothing to model or flag. The Unstoppable Force's stun stays unmodelled and was already flagged.
3. The flagging gap was real but elsewhere: `effect_unmodelled` and the effect-ranking pool were keyed on the item's `effect_text`, which is empty whenever the client spell has no readable description. Both now read the client's own effect tables.

## 1. The relic model

### Mechanism (fork, commits 6d486dcf6, c88324d55)

A relic's value is an equip spell whose effect is a spell-family modifier. `core.EquipSpellMod` carries the four client numbers (aura 107 flat or 108 percent, property from `EffectMiscValue_0`, amount from `EffectBasePointsF`, family from `EffectSpellClassMask_0..3`). Each class supplies a `core.ClassMaskTable` saying which engine `ClassSpellMask` bits each client family is, and `core.NewEquipModItemEffect` turns the pair into `AddStaticMod` configs. It panics at registration when a modifier reaches no modelled spell or has no engine expression, so an item cannot count as modelled while part of it does nothing.

Supported (aura, property) pairs: damage (flat and percent), duration of a damage-over-time effect (new `SpellMod_DotDuration_Flat`, whole ticks only), crit chance, cast time, cooldown, cost (flat and percent). Aura 13 (MOD_DAMAGE_DONE) is not needed: no relic at item level 40+ in this client uses it, and Crackling Staff's spell damage is a plain stat the item data already holds.

`tools/relicmods/gen.py` writes `sim/<class>/relic_mods_auto_gen.go` from the raw client CSVs: every relic at item level 40+ whose equip spell is a 107/108 modifier, with the client spell id, spell text and raw class mask in comments, and the relics it left out named at the foot of the file. `make relicmods` runs it. The table records every modifier faithfully; which items the engine claims is the explicit registration list in each class's `items.go`.

`core.SpellFingerprints` and `core.ChangedSpells` snapshot a unit's modifier-visible spell state so a test can assert the modifier moves the right spell and no other.

### What is modelled

| Item | Client spell | Effect in the engine | Test |
|---|---|---|---|
| Libram of Invocation 249442 | 1249005, aura 108, cost -5 | Seal of Righteousness, Command and the Crusader cost 5% less mana | `paladin/retribution/relics_test.go` |
| Libram of Law 272435 | 1291086, aura 108, damage +4 | Judgement of Righteousness +4%. The client family does not include Judgement of Command, so a Seal of Command paladin gets nothing | same |
| Libram of Infusion 279248 | 1306429, aura 107, crit +6 | Holy Shock crit +6% | same |
| Libram of Fervor 23203 | 28852, aura 107, effects 3 and 8 | Seal of the Crusader +48 attack power, Judgement of the Crusader +33 holy damage taken. Already hand-coded in `sotc.go` from literals; now registered and read from the client table. Also fixes the aura's expiry, which added the libram's attack power instead of removing it | same (attack power returns to idle when the seal fades) |
| Totem of Thunder 228176 | 461295, aura 107, crit +1 | Lightning Bolt crit +1% | `shaman/elemental/relics_test.go` |
| Burning Totem 272433 | 1291077, aura 107, duration +3000 ms | Flame Shock, every rank, one more 3 s tick | same |
| Totem of the Storm 272432 | 1291078, aura 4 | Lightning Bolt can grant Maelstrom Weapon at half a melee hit's chance (0.5 x 0.5). Hand-written; the talent's trigger now asks `maelstromWeaponChance(spell)` | `shaman/relics_test.go`, `shaman/elemental/relics_test.go` |
| Idol of the Dream 220606 | 446212, aura 107, duration +2000 ms | Rip one more 2 s tick. Rip sets its tick count per cast from combo points, so it now adds `Dot.ModNumberOfTicks` back | `druid/feral/relics_test.go` (a 5-point Rip lasts 18 s, not 16 s) |
| Howling Idol 272427 | 1291059, aura 107, cooldown -3000 ms | Tiger's Fury 27 s | same |
| Talons of Wrath 249441 | 1248996 (50% proc), mana from 1302521 (35) | Wrath that lands: 50% to restore 35 mana. Hand-written | `druid/balance/relics_test.go` |

Class spells a relic reaches gained `ClassSpellMask` bits (the shaman had none). The nine Forever-only relics were added to the item database (`tools/database/overrides.go`; `db.bin` and `db.json` patched in place with nine added items, nothing else regenerated; a full `gen_db` run also reshuffles unrelated enchant and spell-icon rows). Idol of the Moon 23197 (aura 112, the Vanilla class-script form) was already modelled.

### What stays unmodelled, and why

| Item | Why |
|---|---|
| Libram of Grace 22402, Libram of Economy 272436, Libram of Light 23006, Libram of Holy Alacrity 228175 | The engine has no Cleanse, Holy Light or Flash of Light |
| Sentinel's Libram 272434 | Swift Judgement is a talent with no spell in the engine (`paladin/talents.go` reads and discards it) |
| Steadfast Libram 279247 | Holy Shield block value; no DPS spec |
| Libram of Deliverance 213513, Idol of the Heckler 213594, Idol of the Raging Shambler 220915, Idol of the Huntress 227444 | Engraving runes, not equip effects |
| Totem of Ancestral Protection 249443 | Grounding Totem is not registered |
| Tidal Totem 272431, Totem of Urgency 279249, Totem of Life 22396, Totem of Flowing Water 23005 | Healing Wave and Lesser Healing Wave are not registered |
| Totem of the Storm 23199 | Superseded by 272432 in the item data |
| Enraged Idol 272428 | Enrage is commented out in the druid package |
| Idol of Swiftness 279250, Idol of Synthesis 272429, Idol of Health 22399, Idol of Longevity 23004 | Swiftmend and Healing Touch are not registered |
| Idol of the Ursine Twins 279251 | Bear Lacerate and Mangle |
| Swarming Idol 272430 | Its text says Insect Swarm but its class mask is Rip's (the same bit as Idol of the Dream). Registering either reading would be a guess, so it is left to the client to settle |
| Relics below item level 40 (Tenets of the Silver Hand, Mystic Mushroom, Polished Driftwood Icon, Firestorm Totem and the like) | Outside the brief; each is a percentage or regeneration aura, not a spell modifier |

## 2. Weapon procs

The client's `ItemXItemEffect` to `ItemEffect` to `SpellEffect` rows, checked against the tooltip text in `assets/db_inputs/wowhead_item_tooltips.csv`:

| Weapon | Client effect | Engine |
|---|---|---|
| Blackblade of Shahram 12592 | chance on hit, spell 16602 "Shahram" (summon NPC 10718) | Modelled already as the six spells the spirit casts, without the NPC (Vanilla numbers; the client has only the summon row). Added a test for Will (+50 stats) and Flames |
| Teebu's Blazing Longsword 1728 | chance on hit, spell 1300753 "Firebolt", 140 Fire, no coefficient | Modelled already, but with Vanilla's 150 under spell 18086 (this client rows that id at 182). Now 140 under 1300753, with a test. The proc rate stays the assumed 1 PPM; the client has no rate in the rows we read |
| The Unstoppable Force 19323 | +2% crit (stat) and a chance-on-hit 1 s stun | Stun unmodelled, flagged `effect_unmodelled`; it has no DPS effect |
| Shadowsong's Sorrow 21522, Ravencrest's Legacy 21520 | none: no item-effect row, no spell in classic-db, no line in the Wowhead tooltip | Nothing to model |
| Riphook 12653 | equip +22 attack power, folded into the item's stats | Nothing to model |
| Crackling Staff 19102 | equip +15 spell damage, folded into the item's stats (`spell_power` 25) | Nothing to model |

## 3. The flagging fix (site, ranker)

Why a proc weapon went unflagged: `effect_unmodelled` was `EffectText != "" && !effectVerifiedInSim`, and `hasImplementedEffect` (the gate for the effect-ranking pool) also began with `EffectText != ""`. The item pipeline leaves `effect_text` empty when the client spell has no readable description, so a chance-on-hit weapon with a blank spell text was neither flagged when the engine lacked it nor sim-ranked when the engine had it. 64 equippable items with a client proc or use effect have a blank text, five of them engine-implemented (Ravager, The Hand of Antu'sul, Coldrage Dagger, Lord General's Sword, Mark of the Chosen) and 59 not (Kindling Stave, Iceblade Hacker, Warblade of Caer Darrow, Mind Carver, Sul'thraze the Lasher and the rest).

The fix:

* `sim/scripts/clienteffects.py` writes `clienteffects_generated.go` from the client tables: 727 equippable ids that have a chance-on-hit spell, a use spell that is not just an aura, or an equip spell with a behaviour aura (periodic damage or trigger, proc trigger, the 107/108/112 modifiers). A plain stat aura is not listed. The dummy aura is deliberately not a trigger: the client carries Blackhand's Breadth and Eye of the Beast, passive stat sticks, as dummies, and listing them flagged a stat stick as an unsimulated proc (the first draft did exactly that).
* `carriesEffect(c)` is `EffectText != "" || clientEffectItemIDs[c.ID]`. `hasImplementedEffect`, the report's `effect_unmodelled`, `trinketEffectUnmodelled` and the trinket exemption all read it, for every slot.
* `sim/scripts/effectids.py` regenerated from fork 912a458f7 (adds the ten relic ids; it also stops counting `_test.go` files, one of which registered a made-up id). It treats `core.NewEquipModItemEffect(Const, ...)` as a registration.

Effect on the published sets: against the committed nightly output, one pick would newly flag (Abyss Shard, Warlock trinket1, band 50). The 59 unflagged weapon procs are not top picks anywhere today; they are the next ones to be caught.

## 4. Relic slot ranking

A relic scores nothing, so `pick()` chose an arbitrary one, usually unmodelled, and `rankSlotWithEffects`'s 1% margin would have kept that incumbent against any modelled relic gaining less than 1%. A relic slot is now ranked like a trinket slot: `rankRelicSlot` (`relics.go`) sims every relic whose effect a sim exercises (`effectVerifiedInSim`), highest DPS wins, and the winner's gain over an empty relic slot is measured at escalating precision. `rankTrinketSlot`'s tournament body is now `rankByMeasuredGain`, shared by both. `report.go` gates a relic on that gain exactly as it gates a trinket (`trinketGainSignificant`, twice its error): a relic whose simulated gain is noise publishes empty with `no_dps_value`. A slot with no modelled relic still publishes `effect_not_modelled`.

## 5. Re-rank, band 60 (`-weights-iterations 200`)

Gains are the winner's measured DPS over an empty relic slot in the same set at the same seed (the tournament's own numbers, 100 iterations then escalated for the winner). Alliance first.

| Spec | Relic picked | Gain over an empty slot | Others measured |
|---|---|---|---|
| paladin-retribution | Libram of Invocation (both factions) | +3.37 +/- 0.66 (+1.5%), +2.71 +/- 0.67 (+1.2%) | Law +1.3 and +2.2, Fervor and Infusion +0.2 and +1.0; the last two are the same sim as each other (no effect of their own), so that spread is the noise floor at 100 iterations |
| shaman-elemental | Totem of Thunder (both) | +0.82 +/- 0.37, +0.78 +/- 0.37 (+0.7%) | Totem of the Storm and Burning Totem exactly 0: no effect in the ranked Elemental sim (it wears no Maelstrom Weapon melee and casts no Flame Shock that a longer duration reaches) |
| shaman-enhancement | none, `no_dps_value` (both) | 0.00 for all three totems | the ranked Enhancement rotation casts neither Lightning Bolt nor Flame Shock |
| warrior-arms | no relic slot | | |
| rogue-combat | no relic slot | | |

Weapon picks did not change in any of the five: Blackblade of Shahram (Arms, Retribution), Shadowsong's Sorrow and Felstriker (Combat), Crackling Staff and The Unstoppable Force (flagged) are as before. Set DPS against the pre-change ranker on the same fork build: Retribution 242.55 to 245.71 (Alliance) and 221.79 to 227.18 (Horde, where the head slot also moved from Lionheart Helm to Soulforge Greathelm, a re-roll from the changed baseline); Elemental, Enhancement, Arms and Combat unchanged to two places (the old run presumably already wore the same arbitrary relic in the sim and only published the slot empty). Hunter-survival (checked for Teebu) is unchanged: 140 against 150 on a 1 PPM proc is about 0.17 DPS.

## 6. Tests run

* Fork, `go test --tags=with_db` over `./sim/core/... ./sim/common/... ./sim/paladin/... ./sim/shaman/... ./sim/druid/... ./sim/warrior/... ./sim/rogue/... ./sim/hunter/...`: all pass (before the last test-only commit, which was run on its own package).
* Site, `go test ./cmd/leveling-bis/...`: pass. One existing test used Fire Ruby's id for an "effect-free trinket"; the real Fire Ruby has an effect (client spell 24389), so the fixture now uses a stand-in id.

## 7. For the owner

* `sim/go.mod`'s replace still points at the fork worktree locally and is not committed. The site needs the fork pin moved to the fork branch head before the nightly.
* Regenerate with each build: `python3 -I sim/scripts/clienteffects.py data/builds/<build>` (needs the raw CSVs, so run it where they exist) and `python3 -I sim/scripts/effectids.py <fork>`. In the fork, `make relicmods`.
* Open question for the client: Swarming Idol's mask (above).
* The fork's `Dot.ModNumberOfTicks` is only honoured by Rip; Rupture-style dots that assign their tick count per cast would need the same line to take a duration modifier.

# Hunter: Hydra Shot, Lacerate and the new pet abilities (2026-10-07)

Build 1.60.1.70009. Fork branch `hunter-shots`, site branch `hunter-shots`.

## What the client says

| Spell | Id | Learnable? | Client facts |
|---|---|---|---|
| Hydra Shot | 1293020 | **No.** SkillLineAbility (wago, this build) has no row for it, and no talent in `talents/hunter.json` grants or names it. | Level 60, 250 mana, 6 s category cooldown shared with Arcane Shot, effect 121 (weapon damage) plus 260, chain 5 targets, amplitude 0.65 |
| Lacerate | 24118, 24119, 24120, 1299332 | Yes, Survival skill line, levels 30/40/50/60 | 40/65/80/95 mana, GCD only, bleed of 17/28/40/58 per tick (plus 0.5/0.6/0.7/0.8 per level above the spell's own), 7 ticks over 21 s, no attack power coefficient |
| Savage Rend | 1265065-69 | Yes, Pet - Raptor | 50 focus, 60 s, 6 ticks of 26 at rank 5 over 18 s |
| Tendon Rip | 1265038-42 | Yes, Pet - Hyena | 25 focus, 30 s, 3 ticks of 20 at rank 5 over 9 s, 50% slow |
| Web | 1265843, 1265878, 1265880, 1265881, 1265883 | Yes, Pet - Spider | 20 focus, 40 s, 4 ticks of 13 nature at rank 5, root |
| Dismember | 1264758 .. 1264933 | Yes, Pet - Crocilisk | 35 focus, 6 s, 54 direct plus 50% healing reduction |
| Pinch | 1264735 .. 1264742 | Yes, Pet - Crab | 50 focus, 30 s, 95 direct plus slow |
| Sonic Blast | 1264478 .. 1264482 | Rank 5 has no row; ranks 1-4 not checked | 80 focus, 30 s, 84 nature plus cast slow |
| Dust Cloud | 1265899 .. 1265904 | Yes, Pet - Tallstrider | 10 focus, armor reduction 505 at rank 5, no damage |

Findings:

- `LaceratingStrikes` in `sim/hunter` is the Survival talent (a bleed on Mongoose Bite). The trainable Lacerate was missing and is now registered.
- Wowhead's Lacerate tooltip reads 451 over 21 s, but the client table gives 58 per tick (406). 451 equals 58 + 0.8 x 8 levels per tick, so the tooltip is computed at the rank's level cap (68), which no character reaches. The engine and the client-damage test use the level-60 value.
- Pet families come from the Wowhead spell pages ("In the <family> Pet Abilities") and agree with the learn rows. The sim's ranker and ladder use the Cat. The Cat gains none of these abilities.

## What was modelled (fork)

- **Lacerate** (`sim/hunter/lacerate.go`): rank by level, melee special mask, physical bleed (ignores armor), 7 ticks of 3 s, `ClientBaseDamage` declared, melee range gate. Resourcefulness reduces its cost through the melee mask. Savage Strikes does not matter (a non-crit bleed).
- **Pet periodic abilities** (`pet_family_abilities.go`): Savage Rend (Raptor), Tendon Rip (Hyena), Web (Spider), read from the generated rank tables. The pet casts the family ability on cooldown ahead of its focus dump, pooling focus for it.
- **Not modelled on the pet**: Dismember (1.5 damage per focus against 2.0-2.5 for Claw/Bite), Pinch (1.9), Sonic Blast (1.05) return less per focus than the dump they would replace, so the pet would not cast them. Dust Cloud (armor reduction) is not damage, but 505 armor is worth a look for the Tallstrider.
- **Hydra Shot was first built, then removed** on the coordinator's correction that the client has no learn row for it. The conformance golden has no Hydra row.
- Conformance: the new Lacerate rows read `declared, matches` (levels 30, 38, 40, 50, 60). `SUMMARY.md` moved by one hunter row.
- Tests: `lacerate_test.go`, `pet_family_abilities_test.go`, and `TestLacerateDamageMatchesClient` in `spellconst_damage_test.go`.

Talent note: Forever has no "Lethal Shots"; Lethal Attacks is a flat crit stat. Efficiency's text covers "melee abilities" but the engine applies it to Shots, Stings and Volley only (see open questions).

## Rotation search

`rotation-search` per spec, 800 confirm iterations, seed 7, level 60 BiS band and guide build. Lacerate is a candidate for all three specs; Hydra Shot appears as a candidate (the tool reads the client's spellconst) but is not registered, so every Hydra row reads "no effect".

| Spec | Baseline | Lacerate inserted at top (probe) | Search winner | Related to Lacerate? |
|---|---|---|---|---|
| Beast Mastery | 253.1 +- 0.4 | +0.0 (melee gate, never casts) | +2.3 +- 0.5: drop Multi-Shot, Arcane Shot mana gate to 20% | no |
| Marksmanship | 248.0 +- 0.4 | +0.0 | +0.9 +- 0.5: swap Serpent Sting and Multi-Shot | no |
| Survival | 252.5 +- 0.5 | -15.5 +- 0.7 | none beats the incumbent | yes, no gain |

For Survival I also ran Lacerate as the last line behind Immolation Trap, gated on "dot not active": 251.5 +- 0.4 against 252.5 +- 0.5, a change of -1.0, inside the combined error and not a gain. Survival is mana-bound in this build and the bleed (about 21 per second for 95 mana) costs more than it returns once Mongoose Bite, Strider Kick and the trap are already on cooldown.

**Adopted: nothing.** No Lacerate line beats the curated rotation by more than the combined error, so the three curated APLs, the guides' Rotation prose, the ladder goldens and the smoke expectations are unchanged (`go test ./...` under `sim/` passes; `make apl-check` passes). The BM and MM winners come from lines unrelated to this task, so I left them for the rotation lane. They are in `design/reviews/rotation-search/hunter-*.md`.

## Open questions

1. Efficiency's tooltip says it reduces the mana cost of "melee abilities" by 3% per rank, but `applyEfficiency` skips melee. This changes Survival's mana budget, and so Lacerate's value, if corrected.
2. Does Lacerate crit? The Lacerating Strikes patch note says that bleed can crit; nothing says Lacerate does. It is modelled as not critting.
3. Should the sim carry Dust Cloud (Tallstrider, -505 armor) or the Hydra Shot constants the generator still emits? Hydra stays unregistered until a learn row appears.
4. Sonic Blast ranks 1-4 learn rows were not checked, as the family is not wired.
5. The BM and MM search winners (+2.3 and +0.9) are worth a separate adoption pass.

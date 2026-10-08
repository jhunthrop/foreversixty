# Tidal Focus and the Elemental golden, 2026-10-08

## Verdict

Tidal Focus is already implemented and correct; the premise of the lane was wrong on three counts.

1. **It is applied.** `sim/shaman/talents_restoration.go` (fork commit 8c089470e, the healer lane) applies it: a -1% a rank power-cost mod on `ShamanSpellMaskHealing` and `AddStat(stats.Hit, 1% a rank)`. A grep for "tidal" that stops at the generated constants misses it. Existing test: `TestTidalFocusCutsHealingManaCostAndAddsHit` (restoration).
2. **The guide build does not take it.** Restoration order in the live tree is Improved Healing Wave, Totemic Focus, Mindfulness, Natural Grace, Tidal Focus, Improved Reincarnation. The guide's `553302` is 5/5/3/3/0/2: Tidal Focus 0. The raid build has the same Restoration string. A hit baseline of 0 in the published profile is therefore right.
3. **The golden's +6% is Tidal Focus, plus a Mindfulness share.** See below.

## Client rows, spell 16179

| SpellEffect | Index | Aura | Points | Mask | Meaning |
|---|---|---|---|---|---|
| 693720 | 0 | 108 (add percent modifier) | -5 | misc 14 (power cost), class mask 448 / 0 / 16 / 0 | -5% mana cost of the Healing Wave, Lesser Healing Wave, Chain Heal family (and Riptide's word) |
| 1339216 | 1 | 54 (melee and ranged hit chance) | +5 | none | +5% hit |
| 1339217 | 2 | 55 (spell hit chance) | +5 | none | +5% spell hit on every spell |

Points are the five-rank total (the talent has one spell id for all ranks; the description steps 1 to 5). The hit is an unmasked unit stat, so the engine's `stats.Hit` (which serves both auras) is exactly the client's, and the ranker's hit profile sees it as baseline. The cost cut is masked to healing, so a damage spell's cost is untouched.

## Where the golden's movement comes from

`TestElemental` first-listed DPS (Phase 1 Troll, `DefaultTalents` = 550331010002155--50105301005, which carries Improved Healing Wave 5, Mindfulness 1, Tidal Focus 5), by fork commit:

| Commit | DPS | Note |
|---|---|---|
| 808c11774 and every commit up to 37347ac0f | 360.5 | before the healer lane |
| 8c089470e feat(shaman): healing spells and Restoration talents | 399.7 | +10.9%, only shaman files change |
| b9a0ba60e healers drink mana consumables, mana oil and Wisdom values | 406.5 | +1.7% |

Inside 8c089470e (Restoration points zeroed one at a time on that commit):

| Talent removed | DPS | Share |
|---|---|---|
| none | 399.7 | |
| Tidal Focus 5 | 376.5 | -5.8%, so Tidal Focus is +6.2% |
| Mindfulness 1 | 386.4 | -3.3%, so Mindfulness is +3.4% |
| Improved Healing Wave 5 | 399.7 | no effect on a damage build |
| all Restoration | 360.5 | the old golden exactly |

So the +6% is Tidal Focus (5 spell hit, a caster below the 16 point cap). It was attributed correctly. The other third of the jump is Mindfulness (17% of regeneration while casting, the client's text), which the golden adoption message did not name. Neither is a bug: no sim/core change is in 8c089470e, the cooldown gate and healing-power commits do not move this golden, and a damage spell is untouched by the healing-power change.

## Follow-up for the ranker

Tidal Focus 5 is worth about +6% on a bare Phase 1 Elemental, more than any Restoration point the guide spends (Totemic Focus 5 is mana for totems; Improved Healing Wave is a self-heal). The talent search should price the Restoration pool with the engine's real Tidal Focus and Mindfulness. This lane makes no build change.

## Change in this lane

Fork branch `tidal`: `sim/shaman/elemental/tidal_focus_test.go` (Tidal Focus 5 is +5 spell hit on Lightning Bolt, Lightning Bolt's cost is unchanged). No engine code or golden moved; the shaman conformance golden is unchanged.

# Mongoose Bite: the dodge / Expose Prey window (2026-10-07)

Lane `mongoose`. Fork branch `mongoose` (sim/hunter), site branch `mongoose`.

## Client rows (build 1.60.1.70009)

- Mongoose Bite, every rank (1495, 14269, 14270, 14271; levels 16/30/44/58): "Counterattack the enemy for melee weapon damage plus $s1. Can only be performed after you dodge." Cooldown 5 s, GCD 1.5 s.
- Defensive State (5302): "Follows a successful block, dodge or parry." SpellMisc duration index 28 = 5000 ms; SpellAuraOptions: 100% chance, 1 charge, proc type mask 16 (a melee spell consumes the charge).
- Expose Prey (talent node 104985, spell 1310532, max rank 2): "Your attacks against targets with Hunter's Mark have a $s1% chance to activate your Mongoose Bite for $5302d." $s1 = 5 / 10% by rank (the talents/hunter.json rank text); effect 1 is a proc-trigger aura (aura 42) casting 1310726; SpellAuraOptions proc type mask 340 = melee auto (0x4) + melee spell (0x10) + ranged auto (0x40) + ranged spell (0x100): both melee and ranged, autos and specials, not spell damage.
- Expose Prey's trigger, 1310726 "Mongoose Bite activated": 5 s (duration index 28), 1 charge, proc mask 16. Identical rows to 5302, so one aura models both.

## The model (fork `sim/hunter/mongoose_bite.go`, `talents.go`)

- `MongooseBiteWindowAura` (ActionID 5302, 5 s, label "Mongoose Bite Ready"). Mongoose Bite's `ExtraCastCondition` is melee range plus this aura active; the cast spends it before the hit lands, so a proc off that very hit opens a fresh window.
- A dodge of the hunter's opens it (`OnSpellHitTaken`, `DidDodge`). Counterattack keeps its own parry window, untouched. The sim's target never attacks, so without Expose Prey no window opens.
- Expose Prey rolls 5/10% on a landed hit whose proc mask is melee or ranged (autos and specials), only while the target has Hunter's Mark (`HuntersMarkAuraTag`). It no longer resets the cooldown (the old stand-in, now removed).
- Tests first: `sim/hunter/mongoose_bite_window_test.go` (gated without the window, 5 s and spent on cast, dodge opens it and a parry does not, Expose Prey on all four attack kinds, not on spell damage, not without a mark, no cooldown reset).
- Curated rotation: the Mongoose Bite line carries `auraIsActive 5302`; `make apl-sync`/`apl-check` clean with `ENGINE_DIR` at the fork worktree.

## Hunter's Mark and the survival kit (site `sim/leveling/kit.go`)

The engine models Hunter's Mark as a request debuff, not a castable spell, and the curated rotation never cast it. The bare ladder, ranker, talent-search and rotation-search characters carried no mark, so Expose Prey could never roll and Mongoose Bite sat idle (guide 252.5 -> 197.6 DPS in the first talent-search run). `KitBuffs("hunter-survival", level >= 6)` now returns `hunters_mark` (level 6 is the spell's learn level, client spell 1130), the same way a mage carries Arcane Intellect: a Survival hunter marks his target before anything else. The raid preset already carried the mark. Test: `TestKitBuffsSurvivalCarriesHuntersMarkFromSix`. BM and MM kits are unchanged.

## Talent search (guide build, 800 iterations)

Re-run after the engine, kit and rotation changes. Expose Prey is now credited with damage (before: "modeled, no damage", 0.0).

| Talent | DPS/point before | DPS/point now |
|---|---|---|
| Expose Prey (2/2) | +0.00 | +8.2 |
| Lacerating Strikes | +8.7 | +4.0 |
| Strider Kick | - | +12.0 |

Adoption rule (keeps every unmodeled guide talent, Survival majority, at least 1% beyond error, drops nothing the rotation casts):

- "deep Survival" (+16.5%) drops Hawk Eye and Deterrence, unmodeled guide talents: refused.
- "guide, modeled non-damage points re-spent + 1 Counterattack -> Resourcefulness" (+9.4%, 0/20/31, keeps every unmodeled talent) drops Counterattack, which the rotation names: refused, a hold. Counterattack's sim value is 0 because the target never parries, so this "gain" is the model gap, not a better build. Same hold as before this change.
- The guide already takes Expose Prey 2/2 and nothing candidate-wise moves points into or out of it beyond error. **The guide `build:` is unchanged.** The Talents prose now says Expose Prey is a damage talent (about +7% for the two points) that needs Hunter's Mark up.

## Rotation search (after the gate)

With Mongoose Bite gated, the search found one missing line beyond error: Lacerate (rank 4, spell 1299332), `not dotIsActive`, inserted above Strider Kick: +11.3 +/- 0.8 (+5.3%). Adopted. Before the gate Lacerate measured -15.5 (hurts), because Mongoose Bite spam took the globals. A confirming re-run: the curated rotation is the best found (best variant +0.0 within error), Mongoose Bite's removal costs +17.2, Lacerate's +11.8. The search notes "Mongoose Bite requires spell 5302, which nothing in this rotation casts": expected, a dodge or Expose Prey opens it.

Mongoose Bite (14271) is listed `expected_idle` in the curated file: at ladder levels 20 and 30 the guide build cannot reach Expose Prey (tier 4) and the dummy never dodges.

## Published-number movement

Ranker, hunter-survival band 60, scratch `-out` (not committed), guide build, `set_dps`:

| Preset / faction | Before | After | Change |
|---|---|---|---|
| bare alliance | 240.4 | 243.2 | +1.2% |
| bare horde | 243.8 | 244.6 | +0.3% |
| raid alliance | 739.1 | 628.6 | -15.0% |
| raid horde | 750.7 | 637.0 | -15.1% |

Bare: the mark kit and Lacerate offset the lost free Mongoose Bite. Raid: Mongoose Bite was spammed every 5 s with no gate; it now fires only on Expose Prey procs (10% of marked attacks, one cast per window), which is the client's behavior.

Ladder (survival golden regenerated, level 60 DPS 245.2 -> 236.9, level 20 53.8 -> 48.2, level 30 70.7 -> 65.7; Mongoose Bite drops out of levels 20 and 30 and Lacerate now fires from level 30). Guide front matter `updated` moved to 2026-10-07; prose states relative gains only.

## Open

- Dodge never occurs in the sim (the target never attacks); if the target model ever swings, dodges feed the same aura with no further change.
- Counterattack is still parry-only, per the existing engine (the client text for Counterattack itself was not part of this lane).
- The talent-search hold (Counterattack -> Resourcefulness, +9.4%) stays a rotation/model follow-up for the owner.

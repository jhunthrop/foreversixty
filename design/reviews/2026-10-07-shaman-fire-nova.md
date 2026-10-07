# Shaman: Fire Nova, Mana Spring Totem and the two rotations (2026-10-07)

Lane `shaman-nova`. Fork branch `shaman-nova` (two commits), site branch `shaman-nova`. Level-60 numbers use the committed alliance BiS band, the guide's own talents, the default 180 s single-target fight and `sim/cmd/rotation-search`'s paired seeds; "paired" means both rotations were simmed with the same seed. Fire Nova rows are hand-built variants (the search never offers Fire Nova, see Open questions).

## What the client says

- Fire Nova is a trainable: SkillLineAbility rows 8734-8738 (AcquireMethod 0, class mask 64) teach spells 408341-408345 at levels 12/22/32/42/52. Cast data: 95/170/280/395/520 mana, 1.5 s GCD, instant, 10 s category cooldown, 30 yd range. Tooltip: "Instantly inflicts 413 to 459 fire damage to enemies within 10 yd of your active Fire totem" (rank 5, level 60).
- The cast spells carry only a dummy effect (amount 5). The damage rides on the linked spells 8349/8502/8503/11306/11307 (school damage, coefficients 0.100 then 0.143, rank 5 amount 419 + 3.4 per level to 57, variance 0.1098). Those reproduce the tooltip exactly.
- The vanilla Fire Nova Totem ids (1535, 8498, 8499, 11314, 11315) have no learn row in SkillLineAbility, so that registration is removed. `Fire Nova Totem (27623)`, a TBC-era id with source `class_spell`, still appears in the conformance unregistered list; it has no learn row either and I left it.
- Talent text (spell tooltips, nether.wowhead.com/forever): Call of Flame names Fire Nova (+15% at 3/3); Improved Fire Nova (16086, 2 points) is "+20% damage and -2 sec cooldown"; Elemental Fury names Fire spells; Concussion names Lightning Bolt, Chain Lightning and Earth Shock only; Totemic Focus names "totems and spells that summon or move them".
- Mana Spring Totem rank 4 (10497): 100 mana, "restores 10 mana every 2 seconds", duration 300 s (the engine had 60 s and no restore at all, a TODO in the source).

## What is modelled (fork `sim/shaman/fire_nova.go`)

- Five ranks by level. A cast spell per rank (ids 408341-5; mana, 1.5 s GCD, 10 s cooldown, castable only while `ActiveTotems[Fire]` is live) and a damage spell per rank (8349...; `ClientBaseDamage` declared, hits every enemy in the encounter, since the sim has no positions).
- Call of Flame (+5% a point) and Improved Fire Nova (taken as +10% and -1 s a point; the client states one effect for two ranks) stack additively. Elemental Fury applies through the Shaman flag; Elemental Focus and the Clearcasting cost reduction through the cast spell. Elemental Mastery is commented out in the engine today, so it does not apply. Concussion and Totemic Focus do not apply (client text).
- Mana Spring Totem now applies a real MP5 aura (restore x 2.5, 25 at rank 4) for 300 s.
- Tests: `TestFireNova*` and `TestManaSpringTotemRestoresManaForFiveMinutes` in `sim/shaman/elemental/fire_nova_test.go`, `TestFireTotemDamageMatchesClient` (Fire Nova rows) and `TestFireNovaCastMatchesClient` in `sim/shaman/spellconst_damage_test.go`. Conformance: all Fire Nova damage rows read `declared, matches`, cast rows `match` on cost, GCD, cooldown and level; Fire Nova left the unregistered list (Shaman 42 to 41) and the five Mana Spring duration mismatches (300000 to 0) are gone.

## Search results (paired, 2000-3000 iterations, deltas against the previous curated rotation)

Elemental, baseline 119.2 +/- 0.3 DPS (2000 iterations):

| Variant | Delta | +/- |
|---|---|---|
| Fire Nova after the totems, no gate | -12.0 | 0.4 |
| Fire Nova, mana >= 20 / 40 / 60 / 80% | -12.0 / -8.1 / -8.1 / -4.1 | 0.4 |
| Mana Spring Totem in the prepull (-3.5 s) | +7.4 | 0.4 |
| Mana Spring Totem first priority line | +5.0 | 0.4 |
| Mana Spring + Fire Nova, no gate | -6.9 | 0.4 |
| Mana Spring + Fire Nova, mana >= 40 or 60% | +0.3 | 0.4 |
| Mana Spring first priority line, then the search's swap with autocastOtherCooldowns | +6.9 total (swap alone +1.9) | 0.3 (3000 it., seeds 7 and 41 agree) |

Enhancement, baseline 230.5 +/- 0.4 DPS (3000 iterations):

| Variant | Delta | +/- |
|---|---|---|
| Fire Nova after Searing Totem, no gate / 20 / 40 / 60% | -18.6 / -18.6 / -10.0 / -1.9 | 0.7 |
| Fire Nova, mana >= 80% | 0.0 (ties; the gate never opens) | 0.7 |
| Mana Spring Totem prepull at -3.5 s / first priority line | +2.5 / +2.6 | 0.7 |
| Mana Spring prepull at -4.5 s plus a refresh line (adopted) | +3.6 | 0.6 |

Re-running the stock search on the adopted rotations: Elemental baseline 126.2 +/- 0.4, "already the best found"; Enhancement baseline 234.1 +/- 0.8, "already the best found". The stock search's Windfury Totem removal (+2.0 on the old Enhancement rotation) is the solo-sim group-buff artifact its own caveat describes and was not adopted.

## Adopted and why

- Elemental: Mana Spring Totem (10497) as the first priority line, refreshed at <= 1.5 s remaining, ahead of autocastOtherCooldowns (the swap is +1.9 +/- 0.3 on two seeds, so cooldowns start one global later; mechanism not understood). Combined +6.9 DPS (+5.8%) against the old rotation, far beyond error. Mana is the binding constraint of a 180 s Elemental fight and the totem returns about 900 mana for 100.
- Enhancement: Mana Spring Totem first in the prepull at -4.5 s plus a totemRemainingTime refresh line, matching its other totems. +3.6 DPS (+1.6%).
- Not adopted: Fire Nova for either spec. It returns roughly a third of Lightning Bolt's damage per mana (about 1.1 against 3+). No mana gate beat the line's absence.
- Ladder goldens regenerated (Elemental and Enhancement; one more cast per level from 30 up). Guides: the Rotation prose of both specs. `go test ./...` under `sim/` passes; `vitest run src/content src/lib/guides` passes; `make apl-check` passes.

## Open questions

- Improved Fire Nova per-point split (10% and 1 s) is an assumption; the client text is one effect for two points.
- Spell 408423-408428 (the Season of Discovery Fire Nova rune family, coefficient 0.214, rank 5 amount 403) share Fire Nova's family bits but their numbers do not match the tooltip; the 8349 family does. If the live damage uses the 0.214 spells Fire Nova roughly doubles its spell-power term and the verdict above should be re-run (it would still need to beat about 3x worse damage per mana).
- The rotation search's insertion pool is built from effect-2 damage spells, so a dummy-effect cast like Fire Nova is never offered. A site change (outside this lane) is needed for the stock tool to probe it.
- Why swapping Mana Spring ahead of autocastOtherCooldowns gains 1.9 DPS is unexplained; it is stable across seeds. Worth a look at what the autocast fires on the first global.
- Elemental Fury does not reach Searing and Magma Totem damage today (their damage spells carry neither the Shaman nor the Totem flag), although its client text names both. Elemental Mastery is dead code in the engine. Neither was changed here.
- The Mana Spring aura grants MP5 to the shaman only; the raid-buff path in `core/buffs.go` stays the way other players receive it. The mage conformance golden was stale on `forever` before this lane (Frostfire Bolt) and was left to its owner; the regenerated `SUMMARY.md` therefore also carries the mage trainables count.

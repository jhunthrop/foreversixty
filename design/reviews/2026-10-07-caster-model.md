# Caster model audit: Elemental Shaman and Balance Druid

Branch `caster-model` (site) on engine branch `caster-model` (fork), client 1.60.1.70009. The
owner suspected the two lowest casters were low because of model gaps. They were: the engine
carried Classic tooltip numbers and Classic talent rules in a client that states different
ones. Fixing them lowers both specs, because the client's numbers are smaller.

## Method

Every engine registration was compared with `spellconst/{shaman,druid}.json` (flat per-rank
`amount`, `sp_coefficient`), `talents/*.json` text, and the raw SpellEffect / trait-curve rows
behind both. Each defect has a failing test first, then the fix. Damage tables are now
checked rank by rank against the vendored client files
(`sim/shaman/spellconst_damage_test.go`, `sim/druid/spellconst_damage_test.go`).

## Defects found (engine fork paths, before the fix)

| # | Where | Client | Engine was |
|---|---|---|---|
| 1 | `shaman/lightning_bolt.go` tables | rank 10: 196, coefficient 0.714; ranks 1-3 0.429/0.571/0.714 | 428-477, 0.857; rank 8 {145,163} (below rank 7); downrank coefficients |
| 2 | `shaman/chain_lightning.go` | rank 4: 123 at 0.571 | 505-564 at 0.714 |
| 3 | `shaman/flame_shock.go` | rank 6: 166 at 0.214, then 44 at 0.1 on each of 4 ticks | 292; 80 a tick (320 total); low-rank coefficients below 0.214 |
| 4 | `shaman/earth_shock.go`, `frost_shock.go` | Earth Shock 7: 301 at 0.386 (the 4080xx/12207xx ids are the taunting tank variants); Frost Shock 4: 283 | 517-545 (the tank variant's size); 492-520 |
| 5 | `shaman/fire_totems.go` | Searing bolt ("Attack" 10436): 47 at **0.017**; Magma 73; Fire Nova 419 | coefficient 0.083 (almost 5x), rolls |
| 6 | `shaman/talents.go` Concussion | Lightning Bolt, Chain Lightning, Earth Shock (class mask 1048579) | also Flame Shock and Frost Shock |
| 7 | `shaman/talents.go` Elemental Fury | +20% crit damage bonus per rank | +100% at any rank |
| 8 | `shaman/talents.go` Elemental Devastation | melee crit only (aura 52) | raised the unified Crit stat, so the shaman's own spell crit rose; the talent search credited it to an Elemental |
| 9 | `shaman/lightning_bolt.go`, `chain_lightning.go` Overload | overload spell = half amount and half coefficient (408477: 98 at 0.357) | half the base damage, full spell-power share |
| 10 | `druid/wrath.go` | rank 8: 91 at 0.571; rank 1-3 coefficients 0.429/0.486/0.571 | 248-277; 0.123/0.231/0.443 |
| 11 | `druid/starfire.go` | rank 7: 381 at 1.0 | 496-584 |
| 12 | `druid/moonfire.go` | rank 10: 135 at 0.15, then 60 at 0.13 on 4 ticks | 195-228; 96 a tick; low-rank coefficients |
| 13 | `druid/insect_swarm.go` | rank 5: 31 a tick (186) at 0.158 | 54 a tick (324) |
| 14 | `druid/forms.go` Moonkin Form | text: armor, Omen of Clarity, and party aura 24907 = +3% crit, which includes the druid | form registered no effect (comment described a Season of Discovery form); the capstone was worth zero |
| 15 | `druid/talents.go` Nature's Grace | +10% casting speed, GCD -10%, 3 s, any non-periodic crit | flat -0.5 s cast time for 15 s, consumed by the next cast; extra -0.5 s on Wrath's GCD |
| 16 | `druid/talents.go` + `starfire.go` Moonglow | 8/17/25% off Wrath, Starfire, Moonfire, Insect Swarm | 3 points a rank off three spells, Starfire's taken twice |
| 17 | `druid/talents.go` Improved Wrath | -0.1 s and -10% mana per rank | cast time only |
| 18 | `druid/talents.go` Improved Moonfire | 5% a rank on damage and crit (trait curve 5, 10) | 2% a rank |
| 19 | `druid/talents.go` Moonfury | school multiplier on Arcane and Nature (aura 79) | additive on three spells; Insect Swarm missed |

Checked and correct: Lava Burst (220 at 0.714, 2.5 s, 10 s cooldown, +20% on a Flame-Shocked
target, no guaranteed crit in this client's text), Convection, Reverberation, Call of Flame,
Call of Thunder, Elemental Alacrity, Elemental Focus, Chain Lightning's 0.7 bounce, Vengeance,
Genesis, Nature's Splendor, Eclipse, Improved Starfire. Elemental Mastery is not in this
client's tree (nothing to model).

Answers to the open questions: (a) Lava Burst's own numbers were right; it lost to Lightning
Bolt because Lightning Bolt carried 452 base at 0.857 against Lava Burst's client 220 at 0.714.
With client numbers an unconditional Lava Burst is level with leaving it out (below). (b) The
engine put the druid in Moonkin Form at every reset (`balance.go` Reset), so no cast is needed;
the form simply did nothing, now +3% crit. (c) All numbers above.

## Before and after (probe, committed level-60 alliance band, seed 7, 800 iterations)

| Spec | Before | After |
|---|---|---|
| Elemental baseline | 195.2 ± 0.7 | 115.0 ± 0.4 |
| Balance baseline | 183.6 ± 0.8 | 172.8 ± 0.7 |

Elemental probe (delta = DPS the line contributes): Lightning Bolt +103.8 ± 0.8 to +54.2 ±
0.5; Searing Totem +20.8 ± 1.1 to +16.1 ± 0.6; insert Lava Burst -12.3 ± 1.0 to -0.7 ± 0.6
(now level with leaving it out); insert Flame Shock -10.7 ± 1.0 to -5.2 ± 0.6.
Balance probe: Starfire +123.1 ± 0.8 to +120.4 ± 0.7; Insect Swarm +4.5 ± 1.1 to +6.1 ± 1.0.
Balance falls least because it is mostly spell-power bound; Elemental falls hardest because
Lightning Bolt lost over half its base damage and the Searing Totem lost most of its spell-power
share.

## Rotation search on the fixed engine

Both specs: the curated rotation is already the best found (best variant +0.0, within
error). No rotation change is adopted, and Lava Burst stays out. Moonkin Form needs no
prepull line: the engine puts the druid in the form at the pull (a prepull cast would charge
435 mana for a form already up). The guides' Rotation sections now say so, and the Elemental
section drops the stale "+13%" figure and the claim that Lava Burst measured worse.

## Tests

Fork: `go test --tags=with_db ./sim/shaman/... ./sim/druid/...` green after adopting
`TestElemental.results` (P1 average 544 to 379 DPS) and `TestEnhancement.results` (949 to 937),
both explained by the damage numbers. Conformance regenerated with `FOREVER_UPDATE_GOLDEN=1`:
no row moved (that report compares costs, cooldowns, cast times, levels and durations with an
empty talent string, none of which changed). Site: ladder goldens regenerated for druid-balance,
druid-feral (level-10 Wrath/Moonfire), shaman-elemental and shaman-enhancement. `go test
./request/` still fails on `hunter-survival` (engine error, spell 1310533 has no DefenseType)
and `warrior-fury` (golden drift); both are untouched by this branch and the warrior golden
is left as it was.

## Still open

- Mage, priest and others still carry Classic rolls (Fireball 12 is 598-760 in the engine,
  483 in the client), so the cross-class DPS list compares conformed classes with unconformed
  ones; Elemental and Balance look worse than they are against them until those conform.
- The client scales some ranks per level (`EffectRealPointsPerLevel`, e.g. Lightning Bolt 10
  +1.2 a level above 56 to 61, Wrath 8 +1 a level above 54 to 60: +2% to +7% at level 60). The
  spellconst files do not carry it, the warlock lane ignores it too, and it is not modelled.
- Improved Moonfire and Moonglow-style percentage mods apply to base damage only in this
  engine (`BaseDamageMultiplierAdditive`), not to the spell-power share.
- Hurricane is not in Moonglow's mask and its numbers keep the older per-level formula.
- Chain Lightning rank 3's client coefficient (0.517) looks like a transposed 0.571; the
  engine follows the client.
- Variance: the client states `Variance` (about 11%) but the tooltip a flat number; rolls are
  flat, as in the warlock lane. Means are unaffected.
- Moonkin Form's separate value in the Balance build was not re-run; the talent search should
  be repeated on this engine pin.

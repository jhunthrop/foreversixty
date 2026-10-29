# Demonology: the Demonic Sacrifice and Demonic Pact line (2026-10-28)

Engine c89e838c6 (the pin), client 1.60.1.70658, band-60 BiS sets from
`data/builds/1.60.1.70658/bis/warlock-demonology.json` (gnome), 10,000
iterations, seed 1, `sim/cmd/spec-breakdown -rotation <variant> -talents <code>`.
The pet-out rotation is the curated one with Soul Link cast before the pull.

## What the engine models at the pin

- Demonic Sacrifice (18788) is an APL cast that needs a demon out. It grants
  Burning Wish (Imp, +15% Shadow), Fel Stamina (Voidwalker, 2% mana every 4 s),
  Touch of Shadow (Succubus, +15% Fire) or Fel Energy (Felhunter, 3% health every
  4 s), and dismisses the demon (`sim/warlock/talents.go` applyDemonicSacrifice).
- Summoning any demon cancels the buff. With Demonic Pact it survives every
  summon except the sacrificed demon's own (the same function's OnPetEnable hook).
- Summon Imp (688) and Summon Succubus (712) are APL casts with a 10 s cast, 2 s
  less per Master Summoner rank (`sim/warlock/summon_demon.go`, applyMasterSummoner).
- Soul Link has to be cast again after a new summon: a demon leaving takes the
  link with it.

## Variants

The rotation variants are in `design/reviews/rotation-search/variants/`. Each
one is the curated rotation with only the prepull replaced.

| Variant | Prepull |
|---|---|
| pet out (curated) | Soul Link at -2 s |
| `sacrifice-succubus-imp` | Demonic Sacrifice at -15 s (Succubus), Summon Imp at -14 s, Soul Link at -3 s |
| `sacrifice-imp-succubus` | Summon Imp at -30 s, Demonic Sacrifice at -20 s, Summon Succubus at -19 s, Soul Link at -3 s |
| `sacrifice-succubus-alone` | Demonic Sacrifice at -3 s, no demon |

Demonic Pact is a tier-6 talent behind Soul Link, so a Pact build needs a 31st
Demonology point. The point comes out of a tree that has nothing the engine
credits:
- raid: the raid build with Ruin 1 moved to Demonic Pact;
- bare: the leveling build with Soul Harvest 1 moved to Demonic Pact.
  Soul Harvest is a no-op in this engine, because nothing dies mid-fight.

## Raid (raid preset, the guide's `raidBuild`)

| Build | Variant | DPS | ± | vs pet out |
|---|---|---|---|---|
| raidBuild | pet out (Succubus) | 780.1 | 0.33 | - |
| raidBuild | sacrifice Succubus, summon Imp (no Pact) | 694.2 | 0.31 | -11.0% |
| raidBuild | sacrifice Succubus, no demon | 632.0 | 0.29 | -19.0% |
| Pact | pet out (Succubus) | 772.9 | 0.32 | -0.9% |
| Pact | sacrifice Succubus, summon Imp | 759.3 | 0.33 | -2.7% |
| Pact | **sacrifice Imp, summon Succubus** | **791.3** | 0.34 | **+1.4%** |
| Pact, Improved Sayaad 1 to Improved Imp 1 | sacrifice Succubus, summon Imp | 761.3 | 0.33 | -2.4% |

## Bare (bare preset, the guide's leveling `build`)

| Build | Variant | Demon | DPS | ± | vs pet out (Succubus) |
|---|---|---|---|---|---|
| build | pet out | Succubus | 414.6 | 0.20 | - |
| build | pet out | Imp | 420.9 | 0.20 | +1.5% |
| Pact | pet out | Imp | 420.9 | 0.20 | +1.5% |
| Pact | **sacrifice Succubus, summon Imp** | Imp | **457.9** | 0.21 | **+10.4%** |
| Pact | sacrifice Imp, summon Succubus | Succubus | 429.8 | 0.22 | +3.7% |

## Reading

- **The Pact line wins in both presets**, beyond error.
  - Raid: +1.4% against the pet-out raid build.
  - Bare: +8.8% against the best pet-out bare run (Imp out), and +10.4% against the Succubus the sim uses today.
- **The best sacrifice differs by preset:**
  - Bare: sacrifice the Succubus for +15% Fire, then summon the Imp. Soul Fire, Searing Pain and Immolate are about 70% of the warlock's damage. The Imp also adds Master Demonologist's +10% Fire, and the leveling build's Improved Imp feeds Firebolt.
  - Raid: the raid preset hands the Succubus every melee buff (Battle Shout, Windfury, Sunder Armor and the rest). That makes it worth about 147 DPS against the Imp's 35-40, so keeping the Succubus out and sacrificing the Imp for +15% Shadow comes out ahead.
- **Without Pact the sacrifice is a loss.** Sacrificing and summoning another demon cancels the buff. Sacrificing and fighting alone loses Soul Link, Master Demonologist and Demonic Knowledge, and costs 19%.
- The prepull spends mana that the pet-out line keeps: a summon (80% or 100% of base mana, less Master Summoner) plus Soul Link. That cost is already in the totals above. A player would drink it back before the pull, so if anything the sims understate the Pact line slightly.

The curated rotation is one priority list per spec for both presets, so it cannot hold both sacrifices. The guide build is not changed on this evidence alone. See `.superpowers/sdd/2026-10-28-warlock-sims/fix-report.md` for the decision and the talent search that follows.

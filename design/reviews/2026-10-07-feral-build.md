# Feral: the talent-search build adopted (2026-10-07)

Branches: site `feral-build`, fork `feral-build` (off `main` / `forever` at 6cda9f720, not merged).
Client 1.60.1.70009, level 60, the committed alliance band gear, seed 7, 800 iterations.

## Decision

The guide build moves from `0/54232212120032010001/55532` to
`FS1:1.60.1.70009:druid:night-elf:01/55232232121032012001/5053:` (1/37/13): **267.6 +- 0.3 against
221.2 +- 0.3, +46.3 +- 0.5, about +21%**, beyond error and well past the 1% bar.

Against the old guide: +1 Genesis, +1 Heart of the Wild (5/5), +2 Shredding Attacks (3/3),
+1 Shifting Power, +2 Predatory Instincts (2/2); -5 Furor, -2 Natural Shapeshifter. Everything else
is kept, including every talent the engine still does not model.

## Model findings

Question: are Furor and Natural Shapeshifter modelled, and do they reach the cat powershift?

- **Natural Shapeshifter: modelled and wired.** `sim/druid` reads it for the mana cost of Cat Form,
  Moonkin Form and Shifting Power (`100 - 10 * rank` percent). New test
  `TestNaturalShapeshifterCutsTheShiftManaCost` pins 3/3 at 70% of the base cost on Cat Form and on
  Shifting Power.
- **Furor: modelled wrongly, now fixed from the live text.** The engine still had the Era proc (20% a
  rank to set energy to 40 on shifting into Cat Form). The live text (node 104958) is: regain 20% a
  rank of the energy you had when you last left Cat Form, plus 2 energy a rank for each second spent
  out of Bear/Cat/Dire Bear Form, capped at 20 a rank. `sim/druid/furor.go` implements that
  (`furorCatFormEnergy`); Cat Form records its energy and the time when it expires and refills from
  that on the next shift. The cap is read as applying to the whole refill. The Bear Form rage half is
  not modelled (Bear Form is not registered in this fork). Tests: `TestFurorCatFormEnergyFollowsTheLiveText`
  (the formula, ranks 0-5, the cap), `TestFurorReachesThePowershift` (the rotation's own
  `tryPowershift` from 30 energy lands on 0/6/18/30 at Furor 0/1/3/5), `TestPowershiftingNeedsFuror`
  (the hard-coded cat rotation powershifts only with the talent).
- **DPS value, measured.** With the hard-coded `catOptimalRotation` on the gear-free fork test bed
  (300 iterations), Furor 5 + Natural Shapeshifter 2 against none: 465.3 vs 464.8 without Shifting
  Power (+0.1%, the powershift only recovers the energy it spends) and identical (514.1) with it. The
  curated APL never powershifts, so the search probe also reads 0.0 for both. Under the live text a
  powershift hands back a share of the energy you had, so it cannot beat Shifting Power, which grants
  a clean 40. Natural Shapeshifter only trims the mana cost of that spell, and mana does not bind
  (variants that drop it lose nothing).

### Why the classifier called them unmodeled

Not a gap in the engine: a bug in `sim/cmd/talent-search/modeled.go`. Its one regexp matched
`recv.Field` and consumed the receiver, so `druid.Talents.Furor` matched `druid.Talents` and never
the field: every chained read was uncounted. Fixed with a separate chained pattern. The wider scan
needed two exclusions to stay honest, both tested: `_ = druid.Talents.X` lines (the fork's recorded
"changes nothing here" acknowledgement) and files or directories the go tool does not compile
(`_maul.go`, `_tank/`, which still mention Rend and Tear, King of the Jungle, ...). Result for the
druid: Furor and Natural Shapeshifter are "modeled, no damage" (Thick Hide too, as armor);
Subtlety, Nature's Focus, Feral Swiftness, Feral Instinct, Brutal Impact, Feral Charge, Primal Bite
and Natural Reaction stay unmodeled by design (threat, pushback, PvP, bear). Other classes' reports
will also move, because the old scan under-counted them the same way.

## Search result

Re-run on this fork (`talent-search -spec druid-feral`), report `design/reviews/talent-search/druid-feral.md`.
Unconstrained best: deep Feral Combat 277.9 +- 0.3 (+56.7), which drops Subtlety (3), Brutal Impact
(1), Feral Swiftness (2) and Feral Charge (1), all still unmodeled. Same-run ladder: 271.5 (keeps
Feral Instinct 2-3 but drops Subtlety and Nature's Focus), 270.5, and **267.6 the best variant that
drops no unmodeled guide talent**, which is the adopted build.

## Rule check (owner's adoption rule)

| Condition | Adopted build |
|---|---|
| Keeps every unmodeled guide talent | Brutal Impact 2, Feral Charge 1, Feral Instinct 3, Feral Swiftness 2, Nature's Focus 5, Subtlety 3: all kept |
| Drops nothing the rotation casts | Drops Furor and Natural Shapeshifter only; the curated rotation casts neither, and Shred, Rip, Rake, Ferocious Bite and Shifting Power all keep casting (ladder golden: Ferocious Bite now fires, Shifting Power 7-8 casts) |
| Beats the guide by at least 1% beyond error | +46.3 +- 0.5 (+20.9%) |
| Legal (51 points, tier gates) | the tool's legality check passes; 1/37/13 |

Higher-scoring variants (270-278) are not adopted because they drop Subtlety or Nature's Focus,
which the engine does not model. Those two have no DPS effect here (Nature's Focus: pushback
avoidance, nothing interrupts a Patchwerk fight; Subtlety: threat); if the owner wants the extra
+1% to +4% (the 270.5 and 271.5 variants) they cost Subtlety or Nature's Focus points that exist for
a threat or pushback value this sim cannot see. That call is the owner's.

## Changes

Site: `web/src/content/guides/druid/feral.md` (`build:` and the Talents section, plus the Shifting
Power paragraph that said the build did not take it); `sim/request/testdata/ladder/druid-feral.golden.md`
adopted (the re-encoded build; Shifting Power and Ferocious Bite now cast, unresolved ids gone);
`sim/cmd/talent-search/modeled.go` and tests; the regenerated talent-search report.
`data/curated/specs.json` carries no build. `data/builds/*/bis/*` are nightly-owned and were not touched;
they still name the old guide build until the nightly regenerates them.

Fork: `sim/druid/furor.go`, `forms.go`, `druid.go`, and tests. The `TestP1Feral` golden did not move.
`go test ./...` under `sim/` passes, `go test --tags=with_db ./sim/druid/...` passes, web
`vitest run src/content src/lib/guides` passes (14 files).

## Open

1. The Furor energy cap is read as covering the whole refill; the text is ambiguous.
2. Bear Form rage from Furor is not modelled. The Bear tank guide section is unchanged.
3. The engine's rotation helpers still use Era-shaped Furor numbers (`furorCap`, `timeToCast`); they
   only matter to the hard-coded rotation, which the curated rotation does not use.
4. Improved Shifting Power and mana sustain at an 8 s cooldown remain unmeasured.

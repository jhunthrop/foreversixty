# Elemental rotation and parity (2026-10-07)

Branches `ele-rotation` in both repos. Question: Elemental publishes the lowest level-60 raid number of the twenty
specs, and its Lava Burst line (gated on a Flame Shock nothing cast) never fired. What does a Forever Elemental
actually play, and does the engine match the client?

Short answer: the client data and the engine agree (one projectile-speed fix); Flame Shock is a loss in this mana
economy under every variant we could write, so it stays out; Lava Burst goes on cooldown with no Flame Shock
condition and is worth about a third of a percent to anyone who spends the point; the movement in the published
numbers comes from re-spending two dead talent points, which the owner's build rule now allows.

## 1. Client parity

Read from `data/builds/1.60.1.70009/raw` (Spell, SpellEffect, SpellCooldowns, SpellPower, SpellMisc, SpellAuraOptions,
SpellLevels) and `talents/shaman.json`; compared with `sim/core/testdata/conformance/shaman.golden.md`
(`declared, matches` on every elemental row, ranks 1 to top, at every level the golden covers).

| Item | Client | Engine |
|---|---|---|
| Lava Burst 408490 / 1238299 / 1238300 | learned 40 / 50 / 60; base 164 / 196 / 220, +0.9 / +1.1 / +1.3 per level, width 0.2533; coefficient 0.714; cost 165 / 230 / 265; 2.5 s cast, 1.5 s GCD, 10 s category cooldown; **projectile speed 20** | matches, golden `declared, matches`. Speed was missing: **fixed** (fork f5a5b245f), test first |
| Lava Burst Flame Shock bonus | effect 1 is Dummy, amount 20 on every rank, tooltip "20% increased damage". No crit guarantee anywhere in the rows (ranks 2 and 3 carry stray coefficient and per-level values on that dummy, unused at level 60) | flat x1.20 while the caster's own Flame Shock is on the target; the existing test pins it. Not a guaranteed crit |
| Flame Shock rank 6 (29228) | direct 166 +2.1 per level, 0.214; periodic 44 per 3 s tick, 0.1 per tick; 12 s; cost 410; 6 s shock cooldown | matches (direct, per tick, duration, cost, shared shock timer) |
| Lightning Bolt rank 10 (15208) | 196 +1.2 per level, width 0.1084, 0.714; cost 220; 2.5 s cast, no cooldown; speed 20 | matches |
| Chain Lightning rank 4 (10605) | 123 +0.8 per level, width 0.1111, 0.571; cost 485; 2 s cast; 6 s cooldown; 3 targets, 0.7 amplitude | matches. Rank 3's client coefficient reads 0.517 and is registered as stated |
| Earth Shock rank 7 (10414) | 301 +1.9 per level, 0.386; cost 450; 6 s shared cooldown | matches (the duration column mismatch is the interrupt lockout, not damage) |
| Elemental Mastery | not in Forever's Elemental tree (16 nodes, none is it) | nothing to register; `registerElementalMasteryCD` is a commented-out body from the Classic tree |
| Lightning Mastery | not in the tree; Elemental Alacrity replaces it (flat 0.17 / 0.33 / 0.5 s off Lightning Bolt, Chain Lightning, Lava Burst) | `applyElementalAlacrity`, matches |
| Elemental Focus / Clearcasting | 10% on any Fire, Frost or Nature damage spell, aura 16246: one charge, -100% cost on the next damage spell | matches (one stack, 15 s) |
| Convection | -2% per rank on Shock, Lightning Bolt, Lava Burst, Chain Lightning | all four read it |
| Concussion | +1% per rank on Lightning Bolt, Chain Lightning, Earth Shock only | matches |
| Call of Thunder | +3% crit on Lightning Bolt and Chain Lightning | matches |
| Call of Flame | +5% per rank on Fire Totems, Flame Shock, Fire Nova, Lava Burst | matches |
| Lightning Overload | 3 / 7 / 10%, a second spell at half damage and half coefficient | matches |
| Elemental Fury | +20% per rank to the crit damage bonus of Fire, Frost, Nature spells and the Searing and Magma Totems | matches |
| Elemental Devastation | +3% per rank melee crit for 10 s after a spell crit | matches (melee only, as the effect says) |
| Eye of the Storm, Reverberation, Improved Fire Nova | pushback 23 / 47 / 70%; shock cooldown -0.2 s per rank; Fire Nova +10% and -2 s per rank | modeled; none moves the Elemental loop except as noted below |
| Elemental Warding, Elemental Reach, Earthbound | defensive, range, utility | unmodeled by design (`applyUnmodeledTalents`), no outgoing damage |

Fork tests: `go test --tags=with_db ./sim/shaman/...` passes for the package, `elemental` and `warden`.
`TestEnhancement` (the hit weight) and the hunter conformance golden fail on `forever` before this branch and are
left to their lanes; `TestElemental.results` had the same hit-weight drift and is refreshed here.

## 2. The rotation

Why the earlier curated list could not use Lava Burst: the rotation search of 2026-10-06 dropped Flame Shock and left
Lava Burst gated on the Flame Shock dot (so it never cast), and the guide build did not spend the capstone point
either. Both searches were also run on the bare character; the headline is the raid context.

New tooling (site commit c2365830): `rotation-search` and `talent-search` take `-preset raid` (the raid band's gear with
the raid preset's buffs and consumables on the kit, the same character the ranker publishes; bare stays the default
and keeps its report names, raid writes `<spec>-raid.md`) and `rotation-search` takes `-build-code` to wear a build
the guide does not carry (a rotation line for a talent can only be searched on a build that has the talent).

Sweep, raid preset, the adopted build plus the capstone point (`5530311300103051/03/453302`), 10 000 iterations,
paired seed. Delta against the previous curated list (which never cast Lava Burst):

| Variant | Delta |
|---|---|
| Lava Burst first, no condition (adopted) | +0.34% (+1.4, error 0.5) |
| Lava Burst after the cooldowns, or as the prepull cast | +0.3% (same, within error of the adopted line) |
| Flame Shock when the dot is down, Lava Burst whenever Flame Shock is up | -0.9% |
| Flame Shock when the dot is down, no Lava Burst | -4.1% |
| the first variant, Flame Shock only above 90 / 70 / 50% mana | 0.0% / -0.4% / -0.7% |
| Chain Lightning on cooldown at one target (alone, or with Flame Shock) | -18.9% / -14.0% |

Reason, from the cast tables: under the raid kit the character is out of mana for 16 to 20 seconds of the fight with
Lightning Bolt alone. Flame Shock costs 410 mana for about 940 damage (direct plus ticks at raid spell power) against
Lightning Bolt's 220 for about 835; Lava Burst's 20% on a 10 s cooldown does not pay the shock back. Chain Lightning
costs 485 for about 650. Lava Burst costs 265 for about 980, the same damage per mana as Lightning Bolt for more per
second of casting, which is its whole gain. Playing the shock to unlock the bonus moves out-of-mana time from 20 s to
35 s.

Adopted: one new first line (`data/curated/apl/shaman-elemental.json`), Lava Burst rank 3, no condition, with the
reasoning in the line's notes. The Flame Shock gate that made it dead is gone; Flame Shock, Chain Lightning
(multi-target only), Earth Shock (30% mana gate, never fires) and the rest are as before. Synced with
`make apl-sync ENGINE_DIR=...`, `make apl-check` clean (fork 808c11774).

Rotation search results (reports in `design/reviews/rotation-search/`, seed 7, 800 iterations, probe table):

- `shaman-elemental-raid-lava-burst-build.md`: raid, build with the capstone point. Lava Burst removal costs 2.1 DPS
  (error 1.6): the line is right; inserting Flame Shock costs 7.8; nothing else helps. Verdict: already the best found.
- `shaman-elemental-raid.md` and `shaman-elemental.md` (raid and bare, the adopted guide build): Lava Burst reads "no
  effect" because the guide build does not carry the talent; inserting Flame Shock costs 20.1 raid and 6.4 bare.
  Verdict: already the best found.

## 3. Talents

Talent search on the new rotation (`design/reviews/talent-search/`), before the guide moved: the best variant that keeps
every unmodeled guide talent, stays Elemental-majority, is at least 1% ahead beyond error and drops nothing the loop
casts was **Reverberation's two points to Thundering Strikes** (`553031130010305/03/553302`, 30/3/18):
+1.5% raid (error 0.4%) and +1.6% bare (error 0.6%). Reverberation shortens the shock cooldown and the loop casts no
shock. Legality bounds it: moving Elemental Devastation's point too would leave 24 points above Elemental Fury's row,
and the row needs 25.

The capstone point is worth +0.26% to the loop (a spare Improved Healing Wave point moved to Lava Burst), less than a
Thundering Strikes point, so the guide stays without it. The default build therefore never casts Lava Burst; the line
is for builds that take the point (the planner, the ladder's lower bands, any user build).

Guide: frontmatter build, Talents and Rotation prose updated, relative gains only. Re-run after the change, both
presets: "the guide is the best" among the variants that keep every unmodeled talent. (The unconstrained winners, deep
Enhancement and the 30/10/11 respend, drop unmodeled Restoration talents and stay holds under the rule.)

## 4. Movement (ranker, band 60, scratch output, before and after)

| Band 60 | Before | After |
|---|---|---|
| bare, alliance / horde | 137.8 / 137.1 | 140.2 / 140.1 (+1.7%) |
| raid, alliance / horde | 405.1 / 405.1 | 411.0 / 411.0 (+1.5%) |

All of the movement is the Thundering Strikes re-spend; the Lava Burst line adds nothing to the guide build because
the guide build cannot cast it. The ladder golden moved with it (the guide talents at 30 to 60 changed) and gained the
Lava Burst row at 40 to 60 under "learned but unused"; the smoke test declares the capstone as a talent this bare build
does not take. Elemental stays the lowest of the twenty: the rotation is not what holds it down, the mana economy is
(Convection, Elemental Alacrity and Elemental Focus carry about half of the talent value; a bolt that costs 220 mana
against a pool that lasts about 160 s).

## 5. Checks

- Fork: `go test --tags=with_db ./sim/shaman/...` as above; `gofmt` clean.
- Site: `go test ./...` under `sim/` passes; `cd web && npx vitest run src/content src/lib/guides` passes (14 files).
- `make apl-check ENGINE_DIR=/Users/jh/code/wowsims-forever/.worktrees/ele-rotation` clean. `sim/go.mod`'s replace was
  pointed at the fork worktree locally and is not committed.

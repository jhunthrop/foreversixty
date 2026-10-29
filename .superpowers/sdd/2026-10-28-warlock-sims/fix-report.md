# Warlock sims: fix report (2026-10-28)

Branch `worktree-lane-discord-bot`. Engine pin `c89e838c6`
(`sim/enginever/version.go`).

**How everything here was measured.** I fetched the pin into the fork's
remote refs (`git fetch origin`), exported it with `git archive c89e838c6`
to `/private/tmp/claude-501/engine-c89`, and compiled against it with a
scratch `go.work` that replaces `github.com/wowsims/classic` with that
export. I ran `make simdb` first. The fork's checkout was never checked
out or modified. `make apl-sync` was pointed at the export
(`ENGINE_DIR=/private/tmp/claude-501/engine-c89`), so the fork copy it
writes landed in the export, not the fork.

## 1. What the engine has at the pin

All four facts hold at `c89e838c6`:

- **Soul Link (19028)** is an APL cast, not a passive:
  `sim/warlock/talents.go:541-595` (`applySoulLink`).
  - The spell is `SpellFlagAPL`, costs 20% base mana and a GCD, and needs `ActivePet != nil` (:576-594).
  - Its aura gives +3% damage dealt and -30% damage taken to the warlock and the pet (:546-560).
  - A pet leaving deactivates it (:565-570).
- **Demonic Sacrifice (18788)** is an APL cast: `sim/warlock/talents.go:597-722` (`applyDemonicSacrifice`).
  - Buffs: Burning Wish (Imp, +15% Shadow), Fel Stamina (Voidwalker), Touch of Shadow (Succubus, +15% Fire), Fel Energy (Felhunter) (:602-668).
  - Demonic Pact keeps the buff through any summon except the sacrificed demon's own (:670-689).
  - Spell registration: :692-721.
- **The four pets:**
  - `proto/warlock.proto:126-132`: the Summon enum, NoSummon/Imp/Voidwalker/Succubus/Felhunter.
  - `sim/warlock/pet.go:37-48` (`setDefaultActivePet`) and :73-80 (`registerPets`, BasePets = all four).
  - APL-castable summons in `sim/warlock/summon_demon.go`:
    - Felhunter 691 (:26)
    - Imp 688 (:42)
    - Succubus 712 (:60)
    - Voidwalker 697 (:76)
  - Master Summoner takes 2 s off the cast per rank: `talents.go:389-417`.

## 2. Soul Link before the pull (commit b2a47a76f)

**What changed**
- `data/curated/apl/warlock-demonology.json` gains `prepullActions` that cast 19028 at -2 s.
- The `notes` paragraph that blamed spellconst is corrected:
  - 18788 and 19028 are talent-granted ids. The rank check already accepts them through `talents.json` `spell_ids`.
  - The engine registers both itself.
  - The guide never casts the sacrifice; it is only Soul Link's gate.
- `make apl-sync` updated the site copy, `sim/request/apl/warlock-demonology.apl.json`.
- `rotations_smoke_test.go` explains 19028 for the smoke build, which takes no talents.

**Demonology ladder golden** (`sim/request/testdata/ladder/warlock-demonology.golden.md`):

| Level | Before | After | Change |
|---|---|---|---|
| 10-30 | unchanged | unchanged | 19028 joins Unresolved as the expected talent exception; the ladder build has no Soul Link yet |
| 38 | 57.2 | 58.8 | +2.8% |
| 40 | 83.4 | 85.3 | +2.3% |
| 50 | 130.6 | 134.9 | +3.3% |
| 60 | 321.9 | 332.1 | **+3.2%** |

- From level 38 there is one more distinct cast (Soul Link).
- The rest of the cast mix shifts only slightly.
- No violations.

The +3% matches the aura. The smaller gain at 38/40 is plausible: there the
pet's share is larger relative to the warlock's, and the mana that Soul Link
costs before the pull matters more.

**Not done here.** The fork's own copy
(`ui/warlock/apls/forever_demonology.apl.json`) needs the same sync and a
new pin. Until then, `make apl-check` against the pinned fork, which CI's
`rotations` job runs, reports the fork copy drifted. The synced file is in
the export at
`/private/tmp/claude-501/engine-c89/ui/warlock/apls/forever_demonology.apl.json`.

## 3. Each spec picks its own pet (commit d8bb14cbb)

**Code**
- `warlockOptions(summon)` in `sim/request/request.go` returns the options function for one demon.
- Each warlock entry in `specOptions` names its own demon.
- `TestEachWarlockSpecSimsItsOwnDemon` (`mapping_test.go`) pins each spec's demon.

**Measurements**
- Band-60 BiS sets with the guide's builds: `build` bare, `raidBuild` raid.
- `spec-breakdown`, 10,000 iterations, seed 1, total DPS (warlock + pet).
- The demon was switched through a temporary, uncommitted override.

| Spec | Preset | Imp | Succubus | Felhunter | Voidwalker | Best |
|---|---|---|---|---|---|---|
| Affliction | bare | 395.7 | **430.5** | 415.1 | 404.1 | Succubus +3.7% over Felhunter |
| Affliction | raid | 718.2 | **842.7** | 817.9 | 795.5 | Succubus +3.0% |
| Demonology | bare | **420.9** | 414.6 | 386.6 | 375.9 | Imp +1.5% |
| Demonology | raid | 701.6 | **780.1** | 736.6 | 713.6 | Succubus +11.2% over Imp |
| Destruction | bare | 436.3 | **466.2** | 452.8 | 442.8 | Succubus +3.0% |
| Destruction | raid | 760.3 | **883.2** | 859.9 | 837.6 | Succubus +2.7% |

All errors are ±0.18 to ±0.38.

**Choice: the Succubus for all three specs**, so no golden moved.
- The only place it loses is Demonology bare. That spec's leveling build takes Improved Imp 3, and the Imp wins there by 1.5%.
- The raid build takes Improved Sayaad and the Succubus wins by 11%.
- One demon serves both presets, so the Succubus stays.

Pet DPS behind these totals (3,000 iterations, total minus the warlock's own lines):

| Spec | Preset | Imp | Succubus |
|---|---|---|---|
| Affliction | bare | 9 | 44 |
| Affliction | raid | 16 | 140 |
| Demonology | bare | 38 | 45 |
| Demonology | raid | 40 | 148 |

## 4. Demonic Pact line (commit eb7d251e1)

The engine models it at the pin (section 1).

- Variants live in `design/reviews/rotation-search/variants/`.
- The full tables are in `design/reviews/2026-10-28-warlock-demonic-pact.md`.
- Every run: 10,000 iterations, seed 1.
- A Pact build needs a 31st Demonology point:
  - raid: Ruin 1 moved to Demonic Pact;
  - bare: Soul Harvest 1 moved to Demonic Pact. Soul Harvest is a no-op in this engine.

| Preset | Pet-out best | Pact line best | Change |
|---|---|---|---|
| raid | 780.1 (Succubus out, Soul Link) | **791.3**: sacrifice the Imp (+15% Shadow), Succubus back, Soul Link | **+1.4%** |
| bare | 420.9 (Imp out) / 414.6 (Succubus, the sim's demon) | **457.9**: sacrifice the Succubus (+15% Fire), Imp out, Soul Link | **+8.8%** / +10.4% |

- Without Pact, sacrificing and summoning another demon loses 11% (694.2 raid), and sacrificing to fight alone loses 19% (632.0 raid).
- **The Pact line wins in both presets, beyond error.** But the best sacrifice differs by preset:
  - Bare favours Fire: Soul Fire, Searing Pain and Immolate are about 70% of the warlock's damage.
  - Raid keeps the Succubus, because the preset gives it every melee buff.
- The curated rotation is one priority list for both presets, so it cannot express both. I did not adopt it. This is the owner's call; see Concerns.

## 5. The adoption rule and the search (commit c7184c698)

**The rule** (`sim/cmd/talent-search/survival.go`). A talent is survival-only when the client's data says it can add no damage. Both of these must hold:
1. Every rank's own spellconst row is nothing but spell modifiers (aura 107/108).
2. Its top-rank tooltip names at least one class spell, and none of the named spells contributes to damage. A spell contributes to damage if it:
   - deals damage, or puts a negative armor aura on the target;
   - grants a resource;
   - raises damage, crit, haste or attack power;
   - hands off to a dummy, script, trigger or summon.

   A tooltip that names another talent, or names no spell, fails.

**How it applies**
- Survival-only talents are exempt from the "keeps every unmodeled guide talent" clause (`removedUnmodeled`).
- They join the guide re-spend candidate, now labelled "guide, modeled and survival-only non-damage points re-spent".
- The report lists them, and marks them in the credits table.

**Results on this build**
- Warlock: Improved Health Funnel, Demonic Aegis and Improved Voidwalker are survival-only (asserted in `survival_test.go`).
- Demonic Brand, Molten Skin, Grim Reach and Demonic Pact are not.
- Across all classes the rule frees only healing, utility and defensive talents: Improved Renew, Improved Inner Fire, Improved Shield Wall and the like.

**Re-search** at the pin with Soul Link cast:
- bare: no `-keep`;
- raid: `-keep "Soul Link,Suppression"`. Demonic Aegis is no longer kept, because it is survival-only.

| Search | Best overall | Best clean | Adopt? |
|---|---|---|---|
| bare (`build`) | deep Demonology + Aftermath, +12.0 (+2.9%), but it drops Soul Harvest and Grim Reach (protected) | 3 Demonic Embrace -> Unholy Power, **+2.6 ± 1.0 (+0.6%)** | no: under 1% |
| raid (`raidBuild`) | deep Destruction (5/22/24), +22.4 (+2.9%) | the same build, but Destruction holds the most points | no: fails the own-tree rule |
| raid, Demonology-shaped | filler (Health Funnel or Demonic Embrace) -> Improved Sayaad, **+2.0 ± 1.7 (+0.3%)** | | no: under 1% |

**Guide builds unchanged.** I updated the Talents prose in
`web/src/content/guides/warlock/demonology.md`:
- Demonic Sacrifice is taken only as Soul Link's gate and never cast. It also explains why, and what Pact changes.
- Soul Link is new, with a "cast it before the pull, and again after every summon" note.
- The three filler talents add no damage and can go to Improved Sayaad or Unholy Power for well under 1%.

The Rotation prose now opens with Soul Link before the pull, and `updated:`
is now 2026-10-28.

## Verification

- `go vet ./...` and `go test ./...` in `sim/` against the pinned engine: green. The ladder goldens are green.
- `cd data && uv run pytest -q`: 1589 passed, 9 skipped, coverage 92.65%.
- Web: `vitest run src/content/guides src/lib/seo.test.ts`: green.

## Concerns

1. **CI `rotations` job.** `make apl-check ENGINE_DIR=.engine` compares against the pinned fork, which still has the old Demonology copy. The fork needs `make apl-sync` committed and a new pin before this lane merges green. The file to carry is in the export (see section 2).
2. **The Pact line wins and is not adopted.**
   - It is worth +1.4% raid and +8.8% bare, but needs a different pre-pull sacrifice per preset, plus the 31st Demonology point.
   - The talent search cannot value Pact while the curated rotation never sacrifices, so it still reads Pact as zero.
   - Adopting it means:
     - choosing one line, or adding per-preset rotations to the request;
     - a guide build change;
     - a design review of the prose.
3. **The raid preset's Succubus is very strong** (about 145 pet DPS against the Imp's 35-40), because the preset gives the pet every melee party buff (Battle Shout, Windfury, Sunder and the rest). That decides the raid pet and the raid Pact line. If a warlock's party in practice lacks those buffs, the raid answers would flip toward the Imp.
4. **The Imp looks under-modeled.** Affliction's Imp does 9-16 DPS, which suggests the Imp runs out of mana or barely casts. This is worth an engine look before trusting any Imp comparison.
5. The raid search's deep-Destruction winner (+2.9%) says Demonology as a spec trails a Destruction-heavy build on the raid set. Outside this lane.

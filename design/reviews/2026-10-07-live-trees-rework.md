# Live trees rework: the engine, the guides and the goldens follow the overlaid talent trees

2026-10-07, lane `live-trees`. Follows `2026-10-07-live-talents-overlay.md`.

Branches: site `live-trees` (worktree `.worktrees/live-trees`), engine fork `live-trees`
(worktree `wowsims-forever/.worktrees/live-trees`). Neither is merged or pushed.

## What changed in the engine fork

`make talents TALENT_BUILD=1.60.1.70009` regenerated the nine talent protos, the per-class
`talents_auto_gen.go` tables and the UI tree JSONs, then `make proto` regenerated the `.pb.go`
files (protoc was not installed; `brew install protobuf` supplied it). The fork's trees had been
generated from 1.60.1.69893, so the regeneration also cleared drift older than the overlay:

| Class | Layout change in the engine proto |
|---|---|
| Warrior | Fury 18 to 17 talents: Improved Cleave, Precision and Boundless Rage out; Furious Precision, Lingering Rage, Gore Drinker in; Flurry behind Death Wish. Protection: Toughness out, Iron Will in, eight talents moved. |
| Druid | Feral 19 to 20: Mangle is Primal Bite, Primal Fury is Blood Frenzy, King of the Jungle out, Shifting Power and Improved Shifting Power in, Shredding Attacks and Predatory Instincts moved. |
| Hunter | Marksmanship 17 to 16: Improved Serpent Sting out. |
| Paladin | Improved Holy Strike and Crusade (stale proto fields) out. |
| Shaman | Elemental Alacrity and Elemental Fury swapped positions. |
| Mage, Warlock | Hot Streak is Heating Up, Soul Harvesting is Soul Harvest (rename only). |
| Priest, Rogue | Nothing but the build comment. `sim/rogue/talents_auto_gen.go`, `proto/rogue.proto` and `rogue.pb.go` changed by the build string only; no rogue behaviour was touched. |

### Behaviour changes, by source

"Note" is Blizzard's 1 October 2026 patch notes; where the note and the client table disagree the
note is the live state and is cited in a comment at the site. "Text" is Wowhead's hotfix-aware
tooltip in the overlaid tree.

| Change | File | Source |
|---|---|---|
| Furious Precision: off-hand hit 4/7/10% | `sim/warrior/talents.go` (`applyFuriousPrecision`) | Text (hotfix_only talent) |
| Dual Wield Specialization: off-hand rage 10% a point, no hit bonus | `sim/warrior/talents.go` | Note, text |
| Unbridled Wrath: a flat 1 rage, the two-handed doubling is gone | `sim/warrior/talents.go` | Note, text |
| Raging Blows: 3 off Cleave and Whirlwind (was 2, Cleave only) | `sim/warrior/talents.go` | Note |
| Whirlwind strikes with both weapons baseline (was Raging Blows' clause) | `sim/warrior/whirlwind.go` | Note |
| Bloodthirst: 45% of attack power (was 35%) | `sim/warrior/bloodthirst.go` | Note |
| Improved Slam: 3 s of cooldown a rank (was 1.5 s) | `sim/warrior/slam.go` | Note |
| Berserker Rage learned at level 30 (client says 32) | `sim/warrior/berserker_rage.go` | Note |
| Lingering Rage, Gore Drinker: explicit no-DPS-effect notes | `sim/warrior/talents.go` | Text |
| Removed Improved Cleave, Precision, Toughness, Boundless Rage code | `sim/warrior/talents.go` | Tree |
| Shifting Power (instant, 55% of base mana into 40 energy, 16 s cooldown, Natural Shapeshifter discounts the cost) and Improved Shifting Power (cooldown minus 4/8 s) | `sim/druid/shifting_power.go`, `druid.go` | Text plus the lane brief's 55% and 16 s |
| Devouring Plague ticks can crit (`CanCrit`, `OutcomeMagicCritPerTick`) | `sim/priest/devouring_plague.go` | Note |
| Shadow Weaving always applies on a landed Shadow spell (no proc roll, no resist roll) | `sim/priest/talents.go` | Note |
| Inner Focus's crit bonus skips pure DoTs and channels | `sim/priest/talents.go` (`innerFocusAddsCrit`) | Note |
| Champion of the Light 20/40/60% of Intellect (was 33/66/100) | `sim/paladin/talents.go` | Note, text |
| Redoubt 4% block a rank (was 6) | `sim/paladin/talents.go` | Note, text |
| Resourcefulness crit chance 30/60% (was 50/100) | `sim/hunter/talents.go` | Text (the overlay review lists it as a real hotfix) |
| Combustion ends after 3 crits, now `combustionCriticalStrikes` | `sim/mage/talents.go` | Note (the engine already read 3) |
| Improved Serpent Sting damage multiplier removed | `sim/hunter/serpent_sting.go` | Tree |

Already correct and left alone: Holy Shield 30% block, Deflection 1% a rank on the paladin and
warrior, Heating Up's renamed field (still inert, see below).

New shared helper: `core.TalentsStringFromRanks` (`sim/core/talent_string.go`) builds a talent
string from generated proto field names. The druid-feral, hunter and shaman-elemental tests
addressed talents by typed field number or index and broke when the layout moved; they now go
through the helper or look the field up by name.

Fixtures re-fitted to the live trees: `ForeverFuryTalents` is `...-15353100051010501-...`
(Improved Cleave 3 and Precision 3 became Furious Precision 3 and Lingering Rage 3, the latter
with no DPS effect) and `ForeverProtectionTalents` spends Iron Will 5 where Toughness 5 was;
druid-feral `P1Talents` and shaman-elemental `DefaultTalents` are recoded by name (the elemental
golden did not move, which is the check that the recode kept the same talents);
`tools/talentgen/testdata/*.json` are the site's 70009 trees, with edge counts 70 (druid 11,
warrior 7) and the warrior tree sizes 17/17/18.

## What changed on the site

Guide builds, decoded against the pre-overlay trees (`9359537f^`) and re-encoded by tree index
and name:

| Guide | Old | New | What moved |
|---|---|---|---|
| druid/feral | `0/5423222121032010001/55532` | `0/54232212120032010001/55532` | King of the Jungle held 0 points; positions shifted only |
| hunter/marksmanship | `55200004/0050550011500305/5` | `55200004/005055001150305/5` | Improved Serpent Sting held 0 points; positions shifted only |
| warrior/arms | `05325213032310001/0505/5005` | `05325213032310001/0505/055` | Toughness 5 re-spent as Protection Iron Will 5 |
| warrior/fury | `35311103002/350511005050010051/0` | `35311103002/35051105050010501/0` | Improved Cleave 1 re-spent on Furious Precision 1, same tree slot |
| warrior/protection | `0/5555/552531213110010001` | `0/5525003/255513121310001001` | Toughness 5 became Iron Will 5 so Protection stays at 31; the Fury tree's Iron Will 5 became Blood Craze 3 and Lingering Rage 2 (a no-op filler) |

Every other non-rogue guide (mage, warlock, paladin, priest, shaman, hunter BM and survival,
druid balance and restoration) re-encodes to the identical string: the renames kept their
positions. Rogue guides and specs were not touched (`rogue-curate`).

Guide prose: the three warrior guides name the removed and added talents and say where each
removed point went, and state the notes that changed their claims (Unbridled Wrath flat 1 rage,
Dual Wield Specialization without the hit clause, Bloodthirst 45% with no Death Wish
prerequisite, Raging Blows on Cleave and Whirlwind). No renamed or removed mage, warlock, hunter
or druid talent is named in any guide. `data/curated/classes.json` still carries the Heating Up
name in its old spelling because `test_curated` compares it with the committed build files.

Other re-encoded strings: `sim/adapter/testdata/warrior-fury.request.json`; guards in
`sim/internal/enginetalents`, `sim/cmd/leveling-bis` now assert the engine and site layouts agree
(node 104949 resolves to `primal_bite`). No curated list or committed file is keyed to the
reused node ids 105953, 110857 and 104951 outside generated or nightly-owned files.

Ladder goldens (`sim/request/testdata/ladder/*`) are adopted: the strings in each row follow the
new tree widths and the DPS movement is the engine's. At level 60, warrior-fury moves from
257.6 to 255.6 and paladin-retribution from 173.9 to 169.1 (Champion of the Light at 60% of
Intellect); druid-balance and hunter-survival rows differ only by the re-encoded strings.

## Test results

Fork (`--tags=with_db`), run package by package over `./sim/... ./tools/...` minus `sim/web`:
everything passes except `sim/conformance` (below). Warrior, druid, mage, priest, paladin, hunter,
shaman and warlock suites pass. `TestP1DPSWarrior.results` was adopted: DPS rose about 7 to 8%
(Bloodthirst +10 points of attack-power ratio and Whirlwind's off-hand hit), which the changes
explain. Every new behaviour has a test that failed before the change: Furious Precision, Dual
Wield Specialization hit and rage, Raging Blows costs, baseline Whirlwind off-hand, Bloodthirst,
Improved Slam, Berserker Rage level, Shifting Power, Devouring Plague crit, Shadow Weaving,
Inner Focus, Champion of the Light, Redoubt, Resourcefulness, Combustion.

Site: `go test ./...` under `sim/` passes, after `FOREVER_UPDATE_GOLDEN=1 go test ./request/ -run
TestRotationLadder`. Web `vitest run src/content src/lib/guides src/lib/planner`: 618 passed, 1
failed, 1 skipped; the failure is `build-tree.test.ts` "decodes a BiS band's own dash-separated
talent string" (50 points against 51), because the committed `bis/warrior-fury.json` band
strings are nightly-owned and stale; `bis.yml` clears it. `data/tests/test_curated.py`: 22 passed.

## What remains open

1. `sim/conformance` (damage-conformance lane) fails. `talent_gated.go` addresses talents by
   typed (tree, position) and `presets.go` carries old-layout strings. Positions that no longer
   hold the named talent, on the 1-based positions in the new proto: Sniper Shot 16 (was 17, now
   out of range), Divine Favor 12, Holy Shock 14, Death Wish 13, Bloodthirst 17 (was 18, out of
   range), Last Stand 7 (was 6), Berserk 20 (was 19). The harness skips an out-of-range talent
   instead of failing, so the committed goldens lose those rows silently. The goldens also move
   by `required_level 32->30` on Berserker Rage (my hotfix against the client table). Neither
   was touched here.
2. Shadow Weaving: the engine's debuff (`sim/core/debuffs.go`) is 3% a stack, the live text says
   2% a stack and gives three ranks. Not changed; the note only said it cannot fail.
3. Not modelled, by design or lack of a model: Heating Up (Pyroblast cast-time stacks), Primal
   Bite (no bear kit), Sniper Shot's 45 yd range (no range model), Lingering Rage and Gore Drinker
   (no DPS effect), Improved Intercept and the other tank talents.
4. The feral cat rotation is hard-coded Go and does not cast Shifting Power; the spell is
   registered and APL-castable. Whether it is on the GCD is unconfirmed (treated as on it).
5. Shifting Power's 55% base mana and 16 s cooldown come from the brief, not from the client's
   tables (hotfix_only spell 1322605 has no row there).
6. Ladder `zero_casts` violations that this change did not touch (informational unless
   `FOREVER_LADDER_STRICT` is set): hunter-beast-mastery and hunter-marksmanship at level 50 never
   cast Arcane Shot (id 14285 against authored 14287), and mage-arcane at levels 20 and 30 has
   the unresolved id 400573. Not checked against `main`.
7. Fork strings left alone: legacy vanilla-layout talent strings in tests that are skipped by
   `core.SkipAwaitingForeverTalentRewrite` (paladin, priest, rogue, warlock, hunter P1), so the
   talents they read are meaningless before and after.
8. A merge note: the fork's `proto/rogue.proto`, `rogue.pb.go` and `sim/rogue/talents_auto_gen.go`
   differ from `forever` by the build string only, so `rogue-curate` may see a trivial conflict.

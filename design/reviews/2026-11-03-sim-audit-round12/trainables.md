# Round 12: trainables (baseline vs talent, the trainables pipeline, the guide lists)

Lane `trainables`. Site branch `wt/r12-trainables` (from c6bdc855c). No engine change, so no fork branch commits.

Sources used, each cited per row:
- **Client:** wago CSVs for 1.60.1.70753, shipped and push-layered (round 11's copies in `scratchpad/r11/c70753/builds/`). Also 70707 shipped and its `wago-push/`, and the committed raw tables for 70009, 70291, 70374 and 70532 in `data/builds/*/raw/`. Tables: SkillLineAbility (SLA), SpellLevels, TraitDefinition, the legacy Talent, SpellEffect.
- **Wowhead live talent payload:** `https://nether.wowhead.com/forever/data/talents-classic?dv=100`, fetched 2026-11-03 14:20 UTC through `pipeline.wowhead_talents` into `$R/trainables/whroot/` (nothing committed). Its `abilities` key is Wowhead's per-class spellbook. Each entry carries a learn level and a `change` tag (`nowTrainable`, `introduced`, `changedLevels`). Compared with the committed snapshot `data/raw-forever/wowhead-talents-2026-10-07.json`.
- **Wowhead Forever tooltips:** `https://nether.wowhead.com/forever/tooltip/spell/<id>` and the Era ones (`/classic/`), 2026-11-03.
- **Beta logs:** cached foreverlogs healer timelines (`scratchpad/fl`, `r10/healers/fl`) and live `api/reports/<id>/character_abilities` for public reports 2790 to 2811 (2026-10-17 to 10-23). The beta ended about 23 October at levels 13 to 30, so no log reaches level 40, and none postdates the 30 October push.
- **Blizzard and community notes** already in `research/01-official-facts.md` and `research/06-since-announcement.md`.

## 1. Baseline or talent, the five spells (and Lightwell, Voice of Truth, Binding Heal)

How to read the client: an SLA row does not decide this. A talent's spell has the same learn row as a trainer spell (Penance 402174, SLA plus TraitDefinition 135364). In Forever, an ability is a talent exactly when a live trait node teaches it. The legacy `Talent` table is the Era UI leftover (`pipeline/wago.py:55-56`) and lists none of Forever's own talents (Penance, Sniper Shot or Wild Growth: no row).

| # | Spell | Verdict | Evidence | confirmed? |
|---|---|---|---|---|
| B1 | Aimed Shot (19434, 20900-20904) | **Baseline**, trained at 20/28/36/44/52/60 | Client: SLA 10956 and 11870-11875 (Marksmanship, AcquireMethod 0); SpellLevels 20 to 60; no TraitDefinition row in any Forever build; only the legacy Era Talent row 1345. Wowhead live spellbook: hunter 19434 at 20, 20904 at 60 (`changedLevels`). Tooltip: Era "Talent", Forever "Requires level 20". Beta logs: four hunters cast it (report 2790: 19434 x31 and 20900 x17; 2808: 20900 x18; 2809: 20900 x16). Community: "Aimed Shot is trainable rather than a Marksmanship talent" (research/06:77). | Aligned (three sources) |
| B2 | Blessing of Kings (20217; Greater 25898) | **Baseline** at 20 (Greater at 60) | Client: SLA 13455 and 13456 (Protection, class mask 2); levels 20 and 60; no trait row; legacy Talent 1442 only. Wowhead live spellbook: 20217 at 20, tagged **`nowTrainable`**; 25898 at 60. Blizzard recap: "key buff talents (Divine Spirit, Blessing of Kings, ...) are baseline" (research/01:170). Logs: 5 paladins apply 20217 (2701, 2753, 2764, 2811). SpellEffect 696790: EffectBasePointsF 10, so 10%. Wowhead's "Value 11%" is the old die-roll display. | Aligned |
| B3 | Consecration (26573, 20116, 20922-20924) | **Baseline**, 20/30/40/50/60 | Client: SLA 13453, plus 13438-13441 **and** 13444-13447 (every rank 2-5 on two rows; see T4). No trait row; legacy Talent 1435 only. Wowhead live spellbook: 26573 at 20. Blizzard: "Consecration baseline at 20" (research/06:66). Logs: 26573 and 20116 cast by 3 paladins each. | Aligned |
| B4 | Divine Spirit (14752, 14818, 14819, 27841; Prayer of Spirit 27681) | **Baseline**, 30/40/50/60 (Prayer 60) | Client: SLA 7956, 8003, 8004 and 13848 (Discipline, class mask 16); no trait row; legacy Talent 351 only. Wowhead live spellbook: 14752 at 30. Blizzard recap, as for Kings. Logs: one priest casts 14752 (report 2747). | Aligned |
| B5 | Omen of Clarity (16864) | **Baseline** passive at 20 | Client: SLA 9402 (Balance); level 20; no trait row; legacy Talent 788 only. Wowhead live spellbook: 16864 at 20, **`nowTrainable`**. Community: "Omen of Clarity and Nature's Grasp are trainable" (research/06:101). Logs: Clearcasting (16870) procs on 7 druids. | Aligned |
| B6 | Lightwell (724, 27870, 27871) | **Baseline**, 40/50/60. It was never in Forever's tree, so the "removed from the live tree" claim is wrong. | Client: SLA 13838, 13850 and 13851 (Holy, class mask 16); levels 40/50/60. No TraitDefinition row in 70009, 70291, 70374, 70532, 70707 (shipped and pushed) or 70753; only the Era Talent row 1637. Wowhead live spellbook: 724 at 40 and 27871 at 60 (`changedLevels`). The Era tooltip says "Talent"; the Forever one drops it and adds "Requires level 40". The guides already say so (`web/src/content/guides/priest/index.md:26`: "Fear Ward, ..., Binding Heal and Lightwell from the trainer"). No log reaches 40. | Aligned, but on client plus Wowhead only (no cast seen) |
| B7 | Voice of Truth (1310897) | **Baseline** at 20 since the 70707 push | 70707 shipped: TraitDefinition 134825 taught 1310897. The 70707 push repoints 134825 to 1328548 Oath of Light (wago-push TraitDefinition, push 113139), and 70753 ships it that way. SLA 59228 and level 20 stay. Wowhead live spellbook: 1310897 at 20, tagged **`introduced`**; it is absent from the 10-07 snapshot, when it was a talent. | Aligned (client plus Wowhead; post-beta, so no log) |
| B8 | Binding Heal (401937, 1240770-1240774) | **Baseline**, 25/32/38/44/50/56 | 70707 shipped: TraitDefinition 135386 taught 401937; the push repoints it to 1328569 Holy Binding. SLA 50041 and 53923-53927 and levels 25-56 stay. Wowhead live spellbook: ranks 2-6 at 32-56 (`introduced`). Rank 1 is missing there; Wowhead's book also leaves out other 40xxxx ids (Feedback 447549, Touch of Weakness), so read that as its filter, not a removal. Round 6 D2 reached the same verdict. Rank 1 rests on the client alone. | Aligned; rank 1 single-source |

The same push marks Arcane Focus (11222), Piercing Ice (11151), Improved Cleave (12329), Improved Overpower, Toughness and One-Handed Weapon Specialization `nowTrainable` at level 1, and Coup de Grace (1310709) is in the live spellbook. These are the removed talents' effects turned into baseline passives.

## 2. The trainables pipeline against the live tree (all nine classes)

Method: `$R/trainables/tools/disagree.py`, `missing.py` and `whcmp.py` compare `data/builds/1.60.1.70753/trainables/*.json` with `talents/*.json`, the client trait and legacy tables, and Wowhead's live spellbook. The committed files reproduce byte for byte from the 70753 raw (`tools/run.py`; `diff -rq` is clean), so every finding below belongs to the pipeline.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **The "talent the live tree removed is not a trainable" drop was wrong, and on 70753 half dead.** `_removed_by_hotfix` dropped any spell a TraitDefinition row teaches when the live tree has no talent of that name. Its premise ("no player can learn it") is false: the 70707 push *repointed* those nodes, and Wowhead's live spellbook lists the old spells at trainer levels (B7, B8, and the `nowTrainable` list above). On 70753 the shipped trait rows for Voice of Truth and Binding Heal are gone, so the drop no longer fired for them. It still dropped Arcane Focus, Piercing Ice, Improved Cleave and Coup de Grace, all four live baseline passives. | `data/pipeline/trainables.py:323-347` at c6bdc855c; the 70707 push's TraitDefinition rows 134825 and 135386; Wowhead payload of 2026-11-03. The workaround was there to keep Voice of Truth off the paladin cooldown lists (1434a03cf message), which `data/curated/guide-cooldowns.json` (an allowlist) now does on its own. | CONFIRMED (pipeline) | old 1434a03cf (70707 activation, 2026-10-31; round 6 P10 saw the spells, round 11 did not look) | S 1b3646428 |
| T2 | **"Both": 124 talent abilities were listed as plain learn-row spells.** Examples: Penance, Holy Shock, Mortal Strike, Wild Growth, Sniper Shot. Each reader had to re-derive "talent" by matching names against `talents/` (rotation-lists.ts specCooldowns, coachrules `_load_class`). Now an ability a live talent teaches, by name or by any rank's spell id, is `source: "talent"`. Its trainer ranks stay with it (Penance 2-4, druid Lacerate 2-5). Baseline spells keep `skill_line_ability`. | Per class (druid, hunter, mage, paladin, priest, rogue, shaman, warlock, warrior): 9, 10, 43, 8, 6, 7, 9, 7, 25. Scratch regen: 0 trainables named like a live talent and still untagged. | CONFIRMED (data contract) | old daf0fa765 | S 1b3646428 |
| T3 | **Missing: 16 active live talents were in no trainables file.** They have no SLA row, no rank subtext and level 0: Shifting Power; Divine Favor, Oath of Light, Templar's Bulwark; Inner Focus, Silence, Shadowform; Adrenaline Rush, Cold Blood, Ghostly Strike, Preparation, Riposte; Amplify Curse, Curse of Exhaustion, Fel Domination, Soul Link. Each is now listed as `talent`, under its tree. Passive talents stay in the talent file only. | `missing.py`: 16 before, 0 after. The guide code already worked around it with its own talent-file pass (`rotation-lists.ts:536-555`). | CONFIRMED | old daf0fa765 | S 1b3646428 |
| T4 | **Duplicate ranks.** Consecration ranks 2-5 and an Immolate rank were each listed twice, because the client has two SLA rows on the same line (13438-13441 and 13444-13447). One spell is now one rank. | before: `[26573, 20116, 20116, 20922, 20922, ...]`; after: `[26573, 20116, 20922, 20923, 20924]` | CONFIRMED | old daf0fa765 | S 1b3646428 |
| T5 | `write_trainables` now fails fast without the build's `talents/`. data.yml runs `normalize` before `trainables` (`.github/workflows/data.yml:251,264`), so CI is unaffected. | - | n/a | - | S 1b3646428 |
| T6 | Era talent passives that Forever's tree dropped and Wowhead's live spellbook does not list: Improved Taunt 12302, Improved Shield Block 12945, Improved Battle Shout 12318, Improved Demoralizing Shout 12324, Weapon Mastery 29082, Totemic Mastery 16189. They stay listed as baseline passives at level 1 (`active: false`). | Client SLA at level 1; Blizzard calls Improved Battle Shout baseline, and the stream calls Shield Block's two charges baseline (research/06:73). Passives reach no guide list, coach rule or engine report. | OPEN (passives, no effect) | - | - |

Tests first, in `data/tests/test_trainables.py`. These five failed before the fix and pass after it: a talent moved off the tree keeps its learn row as a baseline spell; a talent-taught ability is `talent` (by name); a talent is matched by spell when the names differ; an active talent with no learn row is listed and a passive one is not; a spell on two learn rows is one rank. Two more pass before and after: without live trees nothing is tagged a talent, and a committed-data contract that B1-B8 are `skill_line_ability`. The two old tests that pinned the removal premise are replaced. `pytest tests/test_trainables.py` gives 43 passed; ruff check and format are clean.

Consumers, checked against the scratch regen copied into the worktree and then restored: `web/src/lib/guides/rotation-lists.test.ts` 81 passed, `tests/test_addonrotation.py` and `test_trainables.py` 80 passed, `api/internal/coach` ok. The fork's loader (`sim/conformance/trainables.go`, DisallowUnknownFields) gets no new field; it only sees a new `source` value. It counts every non-`class_spell` source as learnable, so its counts move only by the 16 T3 additions.

### Regenerated diff (scratch only: `$R/trainables/after/` against `before/`)

| class | added | rank lists fixed | retagged `talent` |
|---|---|---|---|
| druid | Shifting Power (talent) | - | 9 |
| hunter | - | - | 10 |
| mage | Arcane Focus, Piercing Ice (baseline passives, T1) | - | 43 |
| paladin | Divine Favor, Oath of Light, Templar's Bulwark (talent) | Consecration | 8 |
| priest | Inner Focus, Shadowform, Silence (talent) | - | 6 |
| rogue | Adrenaline Rush, Cold Blood, Ghostly Strike, Preparation, Riposte (talent); Coup de Grace (baseline passive, T1) | - | 7 |
| shaman | - | - | 9 |
| warlock | Amplify Curse, Curse of Exhaustion, Fel Domination, Soul Link (talent) | Immolate | 7 |
| warrior | Improved Cleave (baseline passive, T1) | - | 25 |

Nothing is removed. Voice of Truth, Binding Heal and Lightwell stay listed as trainer spells, as they are today.

## 3. The list that keeps Voice of Truth and Lightwell off the guide pages

No denylist names them. The pieces:

| # | entry | decision |
|---|---|---|
| L1 | `_removed_by_hotfix` in the pipeline (1434a03cf), the original "keep Voice of Truth off the paladin pages" workaround | **Deleted** (T1). Wrong premise, and redundant: the cooldown lists are an allowlist now. |
| L2 | `rotation-lists.test.ts` listed Lightwell among the "race-locked and unlearnable" priest spells | **Removed** from that list. Lightwell is a trainer spell (B6). It stays off Shadow's cooldown list because `guide-cooldowns.json` does not name it, and the healing priests' rotation lists already show it as their prepull cast. Origin: landing 252455d15 (test only; no page changed). |
| L3 | the same file's "never lists a spell the live tree removed (Voice of Truth, Lightwell)" | **Kept, retitled**: "leaves a long baseline cooldown no curated row names off". Same assertions, true premise. |
| L4 | `data/curated/guide-cooldowns.json` (allowlist) | **Kept.** It decides which baseline cooldowns a page shows. Whether Holy should list Voice of Truth (a 3-minute silence and interrupt immunity, baseline at 20) is a content call for the class lane (see For other lanes). |
| L5 | `data/curated/class-spell-limits.json` notLearnable Shadowfiend (401977) | **Kept.** Its SLA row has AcquireMethod 0 at level 1, so the pipeline lists it. It is a SoD rune id that Wowhead's live spellbook leaves out, and no beta priest cast it. The raceLocked rows are RaceMasks facts; kept. |
| L6 | `data/curated/extra-spell-icons.json` 1310897 (Voice of Truth's icon) | **Kept.** No page names Voice of Truth today, but L4 may add it. Pruning the file (other ids are stale too) would also force a `spellicons.json` regen, which the nightly owns. |

## 4. Fresh check: does the sim agree with the verdicts? (`$R/base-engine` at 2606bc05e)

| # | finding | evidence | confirmed? |
|---|---|---|---|
| E1 | Aimed Shot is registered by level alone. | `sim/hunter/aimed_shot.go:76-89`; conformance rows 19434-20904 match. | Aligned |
| E2 | Consecration is registered by level (Era ids 26573/20116/..., damage ids 1280345+). | `sim/paladin/consecration.go:37-90` | Aligned |
| E3 | Kings is a raid buff at 10%. | `sim/core/buffs.go:283,464-470`; SpellEffect 696790 base 10 | Aligned |
| E4 | Divine Spirit is a raid buff by level. | `sim/core/buffs.go:277`, `DivineSpiritRanks.At` | Aligned |
| E5 | Omen of Clarity applies from level 20; Wrath does not consume Clearcasting. | `sim/druid/talents.go:316,369-375`; Forever tooltip "not consumed by Wrath" | Aligned |
| E6 | Lightwell is a trainer spell at 40/50/60, a prepull cast in both healing priest APLs; a renew of 5 ticks of 160/233/320. | `sim/priest/lightwell.go:50-70`; tooltip "800 / 1600 health over 10 sec" (5 x 160, 5 x 320) | Aligned |
| E7 | Binding Heal is trainable at 25-56 in the engine and in Discipline's APL; Voice of Truth is unregistered utility (paladin golden lists it). | `sim/core/testdata/conformance/paladin.golden.md:1546` | Aligned |
| E8 | **The coach's Holy priest persona tells players to use Shadowfiend.** "Shadowfiend returns a share of the whole mana bar over its fifteen seconds", but Shadowfiend is not learnable (L5) and the Holy APL does not cast it. | `data/curated/coach/priest-holy.md:26` | CONFIRMED (copy) for the platform lane; old (6174dc19c) |
| E9 | The class landing copy still lists Voice of Truth as a Holy talent. It has been a trainer spell since the 70707 push, and the node is Oath of Light. | `data/curated/classes.json:146` | CONFIRMED (copy) for the guides lane; new-data (70707 push) |

## 5. The round-7 audit list on 70753

| # | item | 70753 | status |
|---|---|---|---|
| A1 | Frostfire Bolt | SLA ranks 401502, 1237312 and 1237313 at 40/50/60; registered, and conformance matches | Aligned |
| A2 | Unstable Affliction | No learn row; `class_spell` 427717/1242972/1242971. Wowhead's live spellbook leaves it out too; the fork's disposition says it is not learnable (`sim/conformance/dispositions.go:90`). | Aligned (not learnable; the memory note's "trainer row" is out of date) |
| A3 | Curse of the Elements rework | 440892, 1311676, 1311677 and 1311680 at 20-50; registered | Aligned |
| A4 | Hydra Shot | No learn row; `class_spell` 1293020. Not in Wowhead's spellbook; the fork's NoLearnRow lists it (`hunter.golden.md:252`). | Aligned (not learnable) |
| A5 | Lacerate (hunter) | 24118-24120 and 1299332 at 30/40/50/60, baseline; registered and matching. Druid Lacerate is a Feral talent (1328617) with trainer ranks; now tagged `talent`. | Aligned |
| A6 | Pet abilities (Savage Rend, Tendon Rip, Dismember, Pinch, Sonic Blast) | Pet lines are not player trainables. The engine has all five; Sonic Blast stays uncast with no SLA row (`sim/hunter/pet_family_abilities.go:18-20`). | Aligned |
| A7 | Dark Sacrifice, Victory Rush | Ranks present; registered | Aligned |

## Before/after

No engine change and no APL change, so no spec moves. The pipeline change reaches only generated data (below).

## Files to regenerate

- `data/builds/1.60.1.70753/trainables/*.json` (`python -m pipeline trainables --build 1.60.1.70753`; data.yml, nightly-owned). The diff is the table above; `manifest.json` re-hashes on the next simconst or levels run.
- Downstream readers of trainables: `data/builds/1.60.1.70753/addon-data.json` (`pipeline.addonrotation`; Consecration and Immolate rank lists) and the coach rule files (`pipeline.coachrules`). Both pass their tests against the regen.
- Engine follow: the fork's `sim/core/testdata/conformance/client/trainables/*.json` copies and the conformance goldens (`SUMMARY.md`, the per-class `*.golden.md`). The source column reads `talent` for 124 rows, and Active rises by up to 16 per the T3 additions, less any talents the engine registers talent-gated.

## Measure at refresh

None (no ranker or sim effect).

## For other lanes

- **guides:** `data/curated/classes.json:146` lists Voice of Truth among Holy's new talents. Since the 70707 push it is a trainer spell at 20, and the Holy node is Oath of Light (E9). The same sentence names Instrument of Law, which round 11 found removed from the client.
- **platform (coach catalogue):** `data/curated/coach/priest-holy.md:26` coaches Shadowfiend, which no Forever priest can learn (E8, L5).
- **hybrids (paladin) / healers:** Voice of Truth is baseline at 20 with a 3-minute cooldown. Decide whether `guide-cooldowns.json` should name it for paladin-holy, where an interrupt or silence immunity matters on a caster boss. If yes, the icon entry already exists (L6).
- **healers:** Binding Heal rank 1 (401937) rests on the client alone (B8). Wowhead's book has ranks 2-6 only. Watch the first launch logs for a level-25 to 31 priest casting 401937.
- **physical (Aimed Shot), hybrids/auras (Kings, Consecration, Omen), healers/auras (Divine Spirit):** nothing to change. Each is baseline, level-gated, in the engine (E1-E5).

## Still open

- B6 Lightwell and B8 Binding Heal rank 1 have no cast in any log (the beta stopped at about level 30, before the push). Two sources for Lightwell (client, Wowhead live spellbook); one for Binding Heal rank 1.
- T6: six Era talent passives stay listed at level 1. Whether Forever folds them into the class is unconfirmed, and they reach no output.
- Other junk in the trainables that is outside this lane's question: SoD rune ids with AcquireMethod 0 at level 1 (Heart of the Lion 409580, Redirect 438040, Wyvern Strike) and non-Classic spells (Sacred Shield 53601 at 80, Devastate 20243, Earth Shield 974). None reaches a guide list (allowlist) or a rotation. A pipeline rule for them would need its own evidence pass.
- `forever-trainables-source.md` (memory) still calls UA and Hydra Shot trainer-row finds; 70753 and the fork both treat them as not learnable.

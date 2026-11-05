# Sim-accuracy round 22: platform lane

Site branch `wt/r22-platform` (from int/r22 d5cc4d6b4). Fork branch `lane/r22-platform` (from int/r22 170a59c39). The lane owns sim/core, the web, the addon and shared spell math. Client 1.60.1.70790.

Method: the five seed items in the brief's order, then a fresh re-audit of round 21's shared changes. Every sim went through `$R/sim.sh` (frozen BiS, guide raid builds, 10,000 iterations at seeds 1 and 50001). Every CONFIRMED-BUG has a test that was red before its fix. One wow-player pass, on the rendered Fire guide page.

**Confirmed and fixed: 4** (P2 `old`, P3 `old`, P5 `old`, plus the dead Feral option inside P3). **Cross-class DPS movers: none.** P3 moves druids only: raid feral +0.85%, bear threat +0.5%. P2 changes code in three classes (rogue, mage, shaman) and moves no published number. P5 moves nothing.

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | **There are still no live Forever logs.** api.foreversixty.gg `/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct). foreverlogs.gg `/api/reports/public` now answers `403 {"error":"Forbidden","message":"API access not allowed"}`, with and without `phase=2`, so the public list is closed to us. Warcraft Logs v1 (key from GSM, not stored or printed) lists no Forever zone on www (66 zones), classic (57), fresh (21), vanilla (29) or sod (12). The owner's Mac has only the `_classic_beta_` client and no combat log. The beta caches that rounds 8 to 11 fitted (`scratchpad/fl/rep`, `r10/healers/fl/tl`) have been wiped. | last looked 07:26 UTC 5 Nov | OPEN (no data) | - | - |

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | **The sub-level-20 spell penalty (guides r22 shaman, "Engine questions"): Forever does not have it, so nothing changes.** Classic Era cuts the spell power coefficient of a rank whose spell level is under 20 by (20 - level) x 3.75%. The client has no field for that cut: it is server code in Era, and the 70790 SpellLevels/SpellEffect rows carry no penalty. That leaves Forever evidence. Three independent sources say the cut is gone, at least for heals:<br>(1) **Beta logs, several characters** (round 11, E 074edfadf). Pairs of ranks of one heal from one character fit with no cut, and fit nonsensically with it. Rejuvenation 3 (level 16) against 5 (28): talent multiplier 1.16/1.20 without the cut, 0.77/0.81 with it. Healing Touch 2 against 5: 1.24 without, 0.67 with. Healing Wave 1 against 4: 1.24 without, 0.58 with. Holy Light 3 against 4: 0.94 without, 0.50 with.<br>(2) **In-game tooltips on a level-30 priest at 16 and 151 healing power** (HealRankAdvisor README and source comment, 24 Oct). "The original game's below-level-20 penalty is not used. It does not exist in WoW Forever: Heal rank 1 (learned at 16) and Renew rank 2 (learned at 14) keep their full coefficient." The tooltips include the downrank penalty, so this is the client's own arithmetic.<br>(3) **ForeverChanges' beta measurement.** Healing Touch rank 1 (level 1) keeps 94% at 18. With Era's cut it would keep 28.75%. Ozgar's Forever tool also drops the sub-20 rule.<br>**Damage spells** have no direct measurement. Blizzard's 1 October note puts Spell Damage and Spell Healing under one level-gap rule. One mage's Arcane Explosion ranks leaned the same way (round 11). The only web source that applies the cut to Forever damage is the SpellCoeffsForever addon (Admoney73, 11 Oct), and its own page says it models Classic Era ("Classic Era does not have the level-based downranking penalty"), so it does not count as Forever evidence.<br>The engine's `SpellLevelPenalty` (`sim/core/spell.go:802-829`) is the level-gap rule alone, heals and damage alike, and `TestLevelPenaltyIsForeversLevelGap` pins sub-20 ranks at full share ({16,30}=1, {13,30}=1). No class or rank moves. | https://github.com/mendeleevforever/HealRankAdvisor (README "How +healing is handled", HealRankAdvisor.lua:737-748); https://foreverchanges.pro/downranking (read 5 Nov); https://www.curseforge.com/wow/addons/spellcoeffsforever (read 5 Nov); round 11 healers fit (E 074edfadf); forever-bugs searches ("downrank", "spell power coefficient", "below level 20") turned up nothing relevant | Aligned for heals (three sources). OPEN-low for damage: the settling read is a level 25 to 30 shaman's or mage's tooltip for a sub-20 damage rank (Lightning Bolt 3, Frostbolt 3) at two spell damage values 20 or more apart | - | - |
| P2 | **Four spells registered before the trainer teaches them, with no cast-time check.** Feint (trainables 1966, level 16), Vanish (1856, 22), Counterspell (2139, 24) and Tremor Totem (8143, 18) registered at every level. Nothing checks `RequiredLevel` at cast time, so `spellIsKnown` read true and a rotation could cast them below their level. No shipped card does. Each now registers from its trainer level, the Evocation and warlock-summon pattern. The Nightslayer 3-piece (Vanish -30 s) and Jewel of Kaj'aro (Counterspell -2 s) skip a spell that is not learned yet. Conformance goldens: the rows for these spells at rungs below their trainer level were removed by hand (rogue Feint 10, Vanish 10/20; mage Counterspell 10/20; shaman Tremor Totem 10). A local rewrite on arm64 also moves unrelated last-digit rows, as abb4e6b8e notes, so those were left alone. | fork `sim/rogue/feint.go`, `vanish.go`, `items_sets_pve.go`, `sim/mage/counterspell.go`, `items.go`, `sim/shaman/earth_totems.go`; tests `TestFeintAndVanishAreNotKnownBelowTheirTrainerLevel`, `TestCounterspellIsNotKnownBelowLevel24`, `TestTremorTotemIsNotKnownBelowLevel18` (red, then green); `go test -tags with_db` on sim/rogue/..., sim/mage/..., sim/shaman/... passes except TestElemental, which also fails on base-engine (stat weights, pre-existing per the integrator notes) | CONFIRMED-BUG (engine, latent: no shipped card casts them) | old (round 21 P9 listed them; round 8 R6 listed Vanish) | E f5a419228 |
| P2b | **Warrior and druid spells that register early: the cast-gated ones stay, and two druid spells are not gated at all.** The warrior's Berserker Stance, Demoralizing Shout, Revenge, Shield Block, Shield Wall, Whirlwind and Cleave really are gated at cast time (`stances.go:67`, `whirlwind.go:65`, and others). For those, the only gap is that `spellIsKnown` reads true below the level, and no shipped `spellIsKnown` line names one (checked: the ids every curated APL asks about). They are left as they are. Revenge is under the warrior lane's seed this round, and making these registration-gated means rewriting registrations that other code holds pointers to. **Correction to round 21 P9:** druid Cat Form 768 (trainer 20) and caster Faerie Fire 9907 have no cast gate. `druid.RegisterSpell` checks form only. Faerie Fire registers its rank 4 (level 54, full armor value) at every level. A feral druid below 20 also starts the fight in Cat Form (`druid.go:354`, StartingForm). At `-band 20 -level 15` the Cat Form aura procs once at the pull and drops on the first Wrath. No band or card is affected: the round 21 ladder probe saw no early cast of either spell. | engine read; spec-breakdown druid-feral `-band 20 -level 15` (Cat Form 1.00 procs, 0.0% uptime) | OPEN-small, handed to the druid lane | old | - |
| P3 | **Rend and Tear never counted another player's bleed, and its option did nothing.** Rend and Tear reads `Druid.AssumeBleedActive` ("the raid keeps a bleed up"). `NewFeralDruid` never copied the Feral option's `assume_bleed_active` into it, so the option was dead. The bear's options had no such field, and the site never set either one. In a raid, a warrior's Deep Wounds or a rogue's Rupture is on the boss from the opening seconds, and the druid lane decided in round 21 (C2) that any bleed counts. Now both constructors copy the option, `FeralTankDruid.Options` gains `assume_bleed_active` (field 3; `druid.pb.go` regenerated with the pinned protoc 36.2 and protoc-gen-go v1.36.6, and nothing else changed), and the site sets it for the raid preset only (`sim/request/request.go` `assumeRaidBleed`). The bare fight counts only the druid's own bleeds. | fork `sim/druid/feral/feral.go`, `sim/druid/tank/tank.go`, `proto/druid.proto`; tests `TestAssumeBleedActiveOptionReachesTheDruid` (red: false for true), `TestBearAssumeBleedActiveOptionReachesTheDruid`; site test `TestTheRaidFightAssumesAnotherPlayersBleed` (red for both raid rows, then green) | CONFIRMED-BUG (engine option dead; site never set it) | old (druid round 21 C2 found the gap but not the dead option) | E c3c358267, S 26e4e7d33 |
| P4 | **Arcanite Champion's proc rate and Skyborne launch offsets: nothing new.** forever-bugs #414 has not changed since 21 Oct (closed, 2 comments). Searches of forever-bugs for "Skyborne", "Arcanite Champion" and "Strength of the Champion" find no stat sheet and no rate. Web searches turn up no Forever Skyborne stat table and no Arcanite Champion PPM (warcraft.wiki.gg and the Blizzard Classic thread "not sure of the PPM", as in round 21). The engine keeps 1 PPM and the beta Skyborne row. | https://github.com/ClassicWoWCommunity/forever-bugs/issues/414 ; searches 07:10 UTC 5 Nov | OPEN (unchanged: a 30-minute Arcanite Champion log; one naked Skyborne of each variant) | - | - |
| P5 | **The Fire raid card printed "Fireball RANK 11" beside "Below 30, Fireball is the filler".** It was confirmed on the rendered page (`shots/fire-card-0.png`, row 7). The row cannot fire at 60, because the Scorch filler above it fires whenever Heating Up is known, and it carried the raid rank. Now the Fireball filler carries the complement of the Scorch filler's gate, `not auraIsKnown(Heating Up)`, which is the Arcane Frostbolt pattern. The card's build filter (`buildRunsGates`, whose raid build takes Heating Up) therefore leaves the row off, and the leveling views keep it. The two rows still cover every case. Casts and DPS are identical: raid 60 horde 721.0 / 721.4 at seeds 1 and 50001 (equal to before.tsv), bare 30 76.0 and bare 40 159.4, before and after. The fork copy and `sim/request/apl/mage-fire.apl.json` were synced by hand, and `make apl-check` passes. `TestFireScorchFillerWaitsForHeatingUp` now asks for the complement gate. **wow-player: SHIP** on the re-capture (`shots2/fire-card-0.png`, full page `shots2/fire-1440.png`). | site `data/curated/apl/mage-fire.json`; vitest (31 files under lib/bis and lib/guides, 626 tests) pass; `data/tests/test_apl.py` 24 pass | CONFIRMED-BUG (web card) | R21 (mage F-S1 made the row an unconditional filler; the wow-player flagged the label) | S 5241620e1, E 69c6325fb |

## Fresh re-audit of round 21's shared changes

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **`RangedShootsWhileCasting` (E 723a3ccaf).** Only the hunter sets it. `shootDuringCast` runs `applyEffects` with no `Cast`, so no OnCastComplete fires for an Auto Shot loosed during a cast. Every OnCastComplete consumer was checked: Hand of Edward the Odd, Aura of the Blue Dragon, Ascendance, the client spellcast set rows (spell masks only), Dormant Heart, and Poised (class mask). None matches a ranged auto, so nothing is lost. Cast counts still increment, through `applyEffects`. Multi-Shot is instant (CastingTimeIndex 0). It cancels and then re-enables the auto in the same cast, and `EnableAutoSwing` keeps `swingAt`, so no shot is lost. Extra attacks are not taken during a cast (ranged has none). | `sim/core/attack.go:345-350, 921-939`; `sim/hunter/multi_shot.go:58-105` | Aligned | - | - |
| A2 | **The `spellIsKnown` revert (E 5ea8e0f85).** `APLValueSpellIsKnown.GetBool` is `spell != nil` (`apl_values_spell.go:24`). Its `String()` dereferences a nil spell. That is only reached through fmt, which recovers it, so no crash. P2's new unregistered spells use the same path as Evocation. | engine read | Aligned | - | - |
| A3 | **Demon registration (E 92720e6dd).** Summons register from 1/10/20/30. An unlearned default falls back to the Voidwalker from 10, or to the Imp below 10, and the caller's options are cloned. Its four tests pass on this head. | `go test -tags with_db ./sim/warlock -run 'Summon\|DefaultDemon'` | Aligned | - | - |
| A4 | **Spelltip overrides (S 5d4c1a1dd).** Only Defensive Stance 71 and 7376 carry overrides (30% to 80%), and 13 pipeline tests pass. Blizzard's notes thread has no post after 8 Oct; its first post was last edited 3 Nov, and that edit is the change already handled. None of round 21's other adoptions prints a number the client text contradicts: Bear Form 115%, Maelstrom Weapon, and Hand of Justice 1% (an item proc, outside spelltips). | `data/tests/test_spelltips.py`; https://us.forums.blizzard.com/en/wow/t/2360696 | Aligned | - | - |

## Follow-up (integrator, after int/r22 took druid, hunter, mage, paladin and warrior)

Both lane branches merged int/r22 (fast-forward: site 2cafd03c7, fork 83f1c08c4). Every sim and sim-running test here went through `$R/sim.sh`. The Fury run is `warrior-fury`, horde, raid 60, frozen BiS, guide raid build, 3,000 iterations, seed 1: 877.6 DPS (`$R/platform/fury.txt`, `fury.json`).

**Confirmed and fixed: 2** (F6 `old`, B1 `old`). Neither moves a published number, so there is **no cross-class DPS mover**.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| F1 | **Off-hand white misses 8.5% against the main hand's 18.6%.** This is Furious Precision, a Forever Fury talent: "Increases your chance to hit with off-hand attacks by 4/7/10%" (node 105953, hotfix_only; `sim/warrior/talents.go` applyFuriousPrecision). The guide raid build takes it 3/3: field 22, the fifth digit of the Fury segment of `30323013002-050530025122100231-2`. The gap is 10.1 points, which is the talent's 10%. Both hands use the same white table with the 19% dual-wield penalty (`spell_outcome.go:743`). | spec-breakdown rows OtherActionAttack#1 (main hand) and #2 (off hand) | Aligned (Forever talent) | - | - |
| F2 | **Boss armor 471 from 3,731.** Sunder Armor 5 x 450 = 2,250, Faerie Fire rank 4 505 and Curse of Recklessness rank 4 505 (Forever's row; Era's was 640) come to 3,260, and 3,731 - 3,260 = 471. Faerie Fire is the minor category and Curse of Recklessness its own, and both stack with the major one, as in Classic. | `sim/core/raid_debuff_ranks.go:13-46`, `debuffs.go:543-650` | Aligned | - | - |
| F3 | **Berserking and Death Wish uptimes.** Forever's Berserking lasts 15 s on a 3 min cooldown (client; `races.json`), not Era's 10 s. The run casts it 1.49 times a fight because a varied fight runs past 180 s about half the time, which allows a second cast. Its 11.6% uptime is 20.9 s, inside 15 s x 1.49. Death Wish is 1.09 casts with 18.1% uptime, 32.6 s, inside 30 s x 1.09. | spec-breakdown aura and cast rows | Aligned | - | - |
| F4 | **Heart of the Lion and Hunter's Mark in the raid preset.** Heart of the Lion is a Forever baseline hunter spell from level 1 (trainables 409580; Wowhead Forever "increasing total stats by 10% for all nearby allies"). Its 280/276 attack power value stays pending with the owner. Hunter's Mark is ranged attack power only (Wowhead Forever 14325: "increasing the Ranged Attack Power of all attackers ... by 143"). The engine puts it on the ranged attack table only (`debuffs.go:716-731`), so it gives a Fury warrior nothing. Listing it is right; a page that counts it as helping melee would be wrong (sim view lane). | https://nether.wowhead.com/forever/tooltip/spell/409583 , /14325 | Aligned (engine) | - | - |
| F5 | **Sunder Armor and Expose Armor both applied.** They are one exclusive category (`majorArmorReductionEffectCategory`), so only the larger applies. The published Fury preset does not list Expose Armor at all; the sim view's fixture did. | `debuffs.go:563, 590` | Aligned | - | - |
| F6 | **The raid buffs' Scroll of Strength and Scroll of Agility did nothing.** The site's buff list offers `scroll_of_strength` and `scroll_of_agility` as RaidBuffs. The engine stopped reading both in eb0cea267, an upstream commit from February 2024 that turned Strength of Earth and Grace of Air into auras and dropped their `else if scroll` branches. Now a scroll gives its +17 when no totem of the same stat is up, the rule the Intellect, Stamina and Spirit scrolls already follow. It is not counted twice when the consume slot holds the same scroll. The no-stack-with-the-totem part is upstream's rule and single-source; the client gives scrolls no category that would decide it. No preset or published run sets these flags. | fork `sim/core/buffs.go`; test `TestRaidBuffScrollsOfStrengthAndAgilityApply` (red: +0/+0, then green); `go test -tags with_db ./sim/core` passes | CONFIRMED-BUG (engine) | old (eb0cea267, upstream) | E 145b1fdd3 |
| B1 | **The BiS Play It panel ignored race-keyed rank rows.** The addon reads `race_spell_ids`; the web never did. No web or app surface renders levels 26 to 28: the guide and app cards are level 60, and the BiS band panel shows its top level. So the reported Survival 26 to 28 case shows only in the addon, which is already right. The same gap is real at band tops, though: Survival band 40 (top 49) on Horde printed Raptor Strike rank 7 where an Orc, Troll or Tauren keeps rank 5. The other band-top rows keyed by race are Holy priest Heal at 29 and 39 (Undead), Restoration druid Rejuvenation at 29 (Night Elf, Tauren, Skyborne) and Shadow's Devouring Plague at 28 (Troll, Gnome). `rotationLinesFor` now takes a race, and the BiS page passes each faction band's published race. Captures: `$R/platform/shots3/sv-band40-{horde,alliance}-playit.png` (Horde rank 5, Alliance rank 7). **wow-player: SHIP WITH FIXES, no MUSTs.** | site `web/src/lib/bis/rotation-view.ts`, `pages/bis/[class]/[spec].astro`; vitest `rotationLinesFor with a race-keyed rank row` (red: 14265 for 14263, then green); `src/lib/bis` and `src/lib/guides` (63 files, 1,095 tests) pass | CONFIRMED-BUG (web) | old (race rows predate this round; the web never read them) | S a88d07689 |
| C1 | **An APL value for mana income (paladin lane R2 and P2): specified here, not added.** It is not small. It needs a proto field and regenerated `apl.pb.go`, a trailing-window record of mana gains in `Unit.AddMana` (`sim/core/mana.go:79`), and the site's condition vocabulary to word it on the cards (`web/src/lib/bis/apl-gates.ts` CONDITION_NODES, the condition-words module, `data/pipeline/models.py` if it lists nodes), plus a pin. **Spec:** `APLValueManaIncomePerSecond { double window_seconds = 1; }` (default 10 s). It returns the mana gained over the trailing window divided by the window, net of overflow, from every source: regeneration ticks, Judgement of Wisdom procs, potions, runes and Innervate. A ring buffer of (time, amount) appended in `AddMana`, read lazily, pruned on read. Early in the fight it divides by the elapsed time instead of the window. Pair it with the existing `currentMana`, `spellCurrentCost` and `remainingTime` values: "cast the filler when currentMana + income x remainingTime >= core cost x casts left". That expresses both paladin rules without a preset switch. Test: a unit with a fixed 10 mana-per-second tick reads 10 after the window fills and 0 after the ticks stop for a window. | paladin.md "For other lanes" | OPEN (spec for the next platform pass) | - | - |

**For other lanes (follow-up):**
- **Hunter:** the wow-player SHOULD on the Survival panel: "Rank 5" on a Horde hunter at 48 to 49 needs its reason in the Raptor Strike note's first sentence ("Horde keeps rank 5 for the mana traps need"); your note text.
- **Component lane:** the Play It panel says "40 to 49" while its ranks are the level-49 ones, and it does not name the race a race-keyed row is for (wow-player SHOULDs 1 and 3). The panel should say "ranks for level 49" and add the race where a line used a race row.
- **Sim view lane:** do not count Hunter's Mark as a melee buff (F4). The Fury fixture's Expose Armor row applies nothing beside Sunder (F5).

**Files to regenerate (follow-up):** none. F6 and B1 move no published number and no golden.

## Second follow-up (shaman X2: talent-taught learn levels; the raid build legend)

Both branches merged int/r22 again (site 48c46c8ee, fork 95223c008).

**Confirmed and fixed: 2** (T1 `70753` with a guides handoff, T2 `old`). T1 is data generation only: it moves no number until the integrator regenerates. Once regenerated it changes the addon's rotation rows and the site sim's rank tables for talent spells (below). It moves no published raid-60 DPS: every spell it moves is learned below 60 either way.

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **A talent's spell took its client spell level as its learn level.** Every consumer reads its learn level from the generated `spellranks.json`, and the trainers' view from `trainables/`. That covers the addon rows (`addonrotation._family_ranks`), the site sim's rank rewriter and ladder (embedded spellranks), and the web's `knownAtFor` (first level of a family). Both files took `SpellLevels.SpellLevel`. Since 70753 that row gives Stormstrike 17364 level 40, a tier-3 talent a shaman takes at 25 and the engine casts from 25 (shaman N5). So the Enhancement rows dropped it from 25 to 39. The talent fallback at tier level in `_family_ranks` only covered spells spellranks did not track. Now `pipeline/talent_levels.py` holds the tier ladder (10 + 5 x tier, moved out of addonrotation). A talent's own spells are learned at their tier's first level, and a trainer rank above one no earlier. `spellranks` and `trainables` both apply it. The decision to leave out a single level-0 row is still made on the client's levels, so no new names appear (Shadowform stays out). | `data/pipeline/talent_levels.py`, `spellranks.py`, `trainables.py`, `addonrotation.py`; tests `test_a_talent_taught_spell_is_learned_at_its_tiers_level`, `test_talent_levels_open_each_tier_five_levels_apart` (red: unexpected keyword, missing module), `test_a_talent_taught_rank_is_learned_when_its_talent_can_first_be_taken` (red: 40 for 25), `test_a_talent_level_never_adds_a_name_the_client_levels_leave_out`; `tests/test_spellranks.py`, `test_trainables.py`, `test_addonrotation.py`, `test_addondata.py` pass (122 passed, 1 skipped) | CONFIRMED-BUG (data generation) | 70753 (the client's level moved under the old rule) + guides-handoff (guides r20 to r24 carried "RequiredLevel 40") | S b4a82d5ad |
| T2 | **The card legend read backwards to a raider:** "On the level 60 build, skip that row", read by a level-60 raider. The page's build names stay. "Level 60 build" was chosen in the 22 Oct review (cc11b1035), and the guides' prose uses it. The legend now names the raid build alone: "Raid build only: needs a talent only the raid build takes. Not on the raid build? Skip that row." Captures: `$R/platform/shots4/{mage-arcane,warrior-arms}-card.png`. | `web/src/lib/guides/copy.ts`; vitest `raid-build-legend.test.ts` (red on the old text, then green); `src/lib/guides` 30 files, 549 tests pass | CONFIRMED (copy); wow-player SHIP with one SHOULD (name the Level 60 build chip in the legend: not taken, it was the original confusion) | old | S d0a8bc90e |

**What T1 moves at the next regeneration**, simulated against the committed 70790 files without writing them:
- **spellranks, spells with a cost, cast or GCD (14 families):**
  - Talent spells that move earlier: Stormstrike 40 to 25; Counterattack 30 to 25; Vampiric Embrace 30 to 25 (and its 1237618 0 to 25).
  - Talent spells that move later: Insect Swarm rank 1 20 to 25; Repentance 20 to 30; Ice Block 4 to 25.
  - Talent actives that sat at level 0 or 1, now at their tier's level: Shifting Power 40, Swiftmend 25, Blade Flurry 30, Ghostly Strike 25, Mana Tide Totem 25, Death Wish 30, Reckless Swings 25, Spearing Strike 25.
  - 131 passive talent families move from level 0 or 1 to their tier's level.
- **trainables (active abilities, 38):** the same moves, plus talent actives the trainers' list had at level 0 or 1:
  - Druid: Nature's Swiftness 30.
  - Mage: Arcane Power 40, Presence of Mind 30, Cold Snap 30.
  - Paladin: Divine Favor 20, Oath of Light 25, Templar's Bulwark 20, Swift Judgement 25.
  - Priest: Inner Focus 20, Shadowform 40, Silence 30.
  - Rogue: Adrenaline Rush 40, Cold Blood 25, Preparation 30, Riposte 20.
  - Shaman: Lightning Mastery 40, Nature's Swiftness 30, Rage of the Farseer 40.
  - Warlock: Amplify Curse 20, Bane of Havoc 30, Curse of Exhaustion 25, Fel Domination 25, Soul Link 30.
  - Warrior: Last Stand 20.
  - Turn Evil (paladin) does not move. Its trainer ranks 1328372 to 1328374 are not the talent's own spell (1328370), so they stay at 38 and 52. Improved Mend Pet is passive.
- **Addon rotation rows (cards):**
  - Rows gained: Enhancement Stormstrike 25 to 39; Survival Counterattack 25 to 29; Shadow Vampiric Embrace 25 to 29.
  - Rows dropped before their talent's tier (a level-10 character showed them):
    - Druid: Balance Insect Swarm 20 to 24; Feral Shifting Power 10 to 39; Restoration Nature's Swiftness 10 to 29 and Swiftmend 10 to 24.
    - Mage: Arcane Power 10 to 39, Presence of Mind 10 to 29.
    - Paladin: Protection Swift Judgement 10 to 24.
    - Rogue: Subtlety Ghostly Strike 10 to 24.
    - Shaman: Restoration Mana Tide 10 to 24, Nature's Swiftness 10 to 29.
    - Warrior: Arms Reckless Swings and Spearing Strike 10 to 24; Fury Death Wish 10 to 29, Reckless Swings 10 to 24; Protection Last Stand 10 to 19.
- **Web:** the BiS Play It panels and the level-60 guide cards are unchanged at their band tops. `knownAtFor` reads talent families as open from their first level, now the tier's.

**Files to regenerate (second follow-up), in order, by the integrator (the data job's steps):** `python -m pipeline trainables --build 1.60.1.70790`, then `spellranks`, then the addon data step (`addon-data.json` rotations and `ability_levels`, and the addon's `Data.lua`), then `make simdb` (embeds spellranks). After the regen, run the touched sim/request tests, which read spellranks: rank registration, required-level floors and the ladder. Their goldens follow (TestRotationLadder for the specs listed above). `test_the_committed_rotations_never_show_an_ability_before_its_trainer_level` passes only once trainables and addon-data are regenerated together.

## Third follow-up (talent rows at the build's point; unknown auras in APL conditions)

Both branches merged int/r22 again: site 366e2a61b, fork f0cdaffee. Every sim and sim-running test here went through `$R/sim.sh`.

**Confirmed and fixed: 2.** U1 is `R22` (it corrects my own T1). U2 is `old`. **U2 is cross-class code, but no published number moves.**

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| U1 | **Card rows for a talent's spell started at the tier's first level, not where the published build takes the talent.** T1 floored them at 10 + 5 x tier. A card describes a published build, though, and Survival's order (`addon-builds.json`, leveling and raid) spends Counterattack at 32, so the row from 25 showed the button seven levels early (hunter F3). `build_rotation` now floors each talent family at the level the spec's leveling order spends the talent's first point. The raid order's level applies where the leveling order never takes the talent. A family is never earlier than its spellranks level. An order point is [tree, tier, column], 1-based, and the i-th point is spent at 10 + i. The tier level stays the floor for spellranks, trainables and the addon's `abilityLevels` (the addon's own "learned at" table). That table is keyed by spell across specs, and Arms and Fury take the same talents at different levels. | `data/pipeline/talent_levels.py` (`build_point_levels`, `load_spec_point_levels`, `floor_pairs`), `addonrotation.py`; test `test_a_talent_row_starts_where_the_published_build_takes_the_talent`, parametrized: Survival Counterattack 32 (red: 30), Enhancement Stormstrike 25 (red: 40), Fury Death Wish 30. It runs on spellranks as T1 writes them, built in the test because the committed file waits for the regen. 175 data tests pass. | CONFIRMED-BUG (card data) | R22 (S b4a82d5ad) | S 6c639ea7f |
| U2 | **An APL condition on an unknown aura or dot ran with no condition.** `newValueAuraIsActive`, `...WithReactionTime`, `...NumStacks`, `...RemainingTime`, `newValueDotIsActive` and `...DotRemainingTime` returned nil for an aura or dot the character never registered. `newValueAnd` and `newValueOr` filter nil operands, `newValueNot` and `newValueCompare` return nil, and `APLAction.IsReady` reads a nil condition as always true. So a gate on an untaken talent's aura vanished, and the line ran unconditionally. The zero-talent Fire smoke build's Scorch recast every global this way (`sim/request/rotations_smoke_test.go`). Now an unknown aura reads inactive, 0 stacks and 0 remaining, and an unknown dot inactive and 0 remaining. `GetAPLAura` still warns. `auraIsKnown` and `spellIsKnown` already read false. The ICD and should-refresh values keep nil, because no curated condition uses them. **Survey and measurement.** Rather than guess which ids each build registers, I ran every curated APL before and after, on frozen BiS and guide builds: raid 60 for 28 specs (horde, seed 1, 3,000 iterations), every bare band floor 20/30/40/50 for 28 specs (2,000 iterations), and the whole ladder (TestRotationLadder written with each engine, diffed, then reverted). **Every number is identical.** No published build or band reads an unknown aura in a gate that decides a cast. The fix guards custom builds and future lines. The fork's TestP1Mage fails the same 31 subtests before and after; that is already stale on int/r22, not this change. | fork `sim/core/apl_values_aura.go`, `apl_values_core.go`; test `TestAnUnknownAuraReadsInactiveInsteadOfDroppingTheGate` (red: auraIsActive(unknown) = nil); `go test -tags with_db ./sim/core` passes; `sim/request` smoke tests pass; screens in `$R/platform/aura-screen.tsv` and `aura-bands.tsv` | CONFIRMED-BUG (engine; cross-class code, no published mover) | old (upstream wowsims behaviour) | E 30bf8bdb0, S c2fe5d894 (smoke note) |

**Card rows after the regen, simulated without writing anything: T1 alone compared with T1 plus U1.** "Committed" is today's first level.

| spec | row | committed | T1 | now |
|---|---|---|---|---|
| druid-balance | Insect Swarm | 20 | 25 | 27 |
| druid-feral | Shifting Power | 10 | 40 | 40 |
| druid-restoration | Nature's Swiftness | 10 | 30 | 30 |
| druid-restoration | Swiftmend | 10 | 25 | 43 |
| hunter-survival | Counterattack | 30 | 25 | 32 |
| mage-arcane | Arcane Power | 10 | 40 | 40 |
| mage-arcane | Arcane Wrath | 25 | 25 | 31 |
| mage-arcane | Presence of Mind | 10 | 30 | 31 |
| mage-arcane | Pyroblast | 20 | 20 | 54 (only the raid order takes it) |
| mage-frost | Ice Lance | 20 | 20 | 30 |
| paladin-holy | Divine Favor | 20 | 20 | 21 |
| paladin-holy | Holy Shock | 30 | 30 | 32 |
| paladin-protection | Swift Judgement | 10 | 25 | 29 |
| paladin-retribution | Seal of Command | 20 | 20 | 26 |
| priest-shadow | Inner Focus | 20 | 20 | 57 |
| priest-shadow | Mind Flay | 20 | 20 | 28 |
| priest-shadow | Vampiric Embrace | 30 | 25 | 33 |
| rogue-assassination | Cold Blood | 25 | 25 | 33 |
| rogue-subtlety | Ghostly Strike | 10 | 25 | 26 |
| shaman-enhancement | Stormstrike | 40 | 25 | 25 |
| shaman-restoration | Mana Tide Totem | 10 | 25 | 25 |
| shaman-restoration | Nature's Swiftness | 10 | 30 | 44 |
| shaman-restoration | Riptide | 40 | 40 | 50 |
| warlock-affliction | Amplify Curse | 20 | 20 | 29 |
| warlock-destruction | Conflagrate | 25 | 25 | 34 |
| warlock-destruction | Shadowburn | 20 | 20 | 30 |
| warrior-arms | Reckless Swings | 10 | 25 | 57 |
| warrior-arms | Spearing Strike | 10 | 25 | 25 |
| warrior-fury | Death Wish | 10 | 30 | 30 |
| warrior-fury | Reckless Swings | 10 | 25 | 31 |
| warrior-protection | Last Stand | 10 | 20 | 28 |

**Regen order (the integrator; the data job's steps, nothing nightly-owned committed here):**
1. `python -m pipeline trainables --build 1.60.1.70790`
2. `python -m pipeline spellranks --build 1.60.1.70790`
3. The addon data step: `addon-data.json` rotations and `ability_levels`, which now read addon-builds.json too. Then the addon's `Data.lua`.
4. `make simdb`.
5. Run the sim/request tests that read spellranks: rank registration, required-level floors and `TestRotationLadder`. Refresh their goldens with the pinned engine.

Since `addon-builds.json` now feeds step 3, regenerate it first (`sim/cmd/addon-builds`) if any lane changed a guide build. The pin needs fork 30bf8bdb0 or later. Still owed from the first pass: ladder goldens for druid-feral, druid-feral-bear and mage-fire.

## Before/after

Raid 60, frozen BiS, guide raid builds, 10,000 iterations. Before is `$R/sb-before` from site-before; after is this branch.

| spec | faction | seed 1 | seed 50001 |
|---|---|---|---|
| druid-feral | alliance | 886.1 -> 893.6 (+0.85%) | 885.6 -> 893.2 (+0.86%) |
| druid-feral | horde | 870.6 -> 878.0 (+0.85%) | 870.5 -> 877.8 (+0.84%) |
| druid-feral-bear (threat/s) | alliance | 1756.3 -> 1765.2 (+0.51%) | 1757.3 -> 1766.1 (+0.50%) |
| druid-feral-bear (threat/s) | horde | 1750.6 -> 1759.1 (+0.49%) | 1750.7 -> 1759.2 (+0.49%) |
| mage-fire | horde | 721.0 -> 721.0 | 721.4 -> 721.4 |

Bare bands are unchanged. P3 applies to the raid preset only, and P5 casts the same at bare 30 and 40. P2 moves nothing at any published band.

## Files to regenerate

- Ladder goldens. These were seen by running the ladder for the three specs and then reverting, so nothing is committed.
  - `sim/request/testdata/ladder/druid-feral.golden.md`: the 60 rung, 224.8 -> 227.7.
  - `sim/request/testdata/ladder/druid-feral-bear.golden.md`: the 60 rung, 208.2 -> 208.9.
  - `sim/request/testdata/ladder/mage-fire.golden.md`: the 10 rung's Unresolved column gains `{SpellID: 400625}`, with the same casts.
- Fork conformance goldens: the explained rows are removed by hand and committed. Any full amd64 rewrite still belongs to CI.
- The published `addon-data.json` rotations still carry the Fireball line until the nightly. The card already drops it, through the curated gate.
- Pin: the site needs fork c3c358267 or later for the bear's `AssumeBleedActive` field. S 26e4e7d33 does not compile against 170a59c39.

## Measure at refresh

- Feral and bear BiS and weights under the raid bleed: +0.85% and +0.5% across the board. Rend and Tear is a flat multiplier, so the picks should not move much.

## For other lanes

- **Druid:** Cat Form 768 and caster Faerie Fire 9907 have no level gate (P2b). The fix: register Cat Form from `CatFormLevel[0]`, with a Humanoid StartingForm below it (the Moonkin precedent at `druid.go:366`), and give Faerie Fire its learned rank, or register it from 18. Check `innervate.go:32` and `items.go:124`, which dereference `druid.CatForm`. Also, your raid feral and bear numbers now include P3. If you edit `feral.go` or `tank.go` constructors, keep the `AssumeBleedActive` copy.
- **Mage:** the wow-player SHOULD on the Fire raid card: row 6's printed sentence "From Heating Up at 30, Scorch is the filler, not Fireball" reads as leveling advice on a level-60 card. Something like "Scorch is the filler, not Fireball, once you have Heating Up" would fit both views. Your file; not changed here. Also, if you edit the Fire APL this round, keep the Fireball row's `not auraIsKnown(400625)`.
- **Rogue, mage, shaman:** Feint, Vanish, Counterspell and Tremor Totem are now nil below their trainer level (E f5a419228). Guard any new reference.
- **Warrior:** your stance and ability gates are cast-time only; `spellIsKnown` reads true below the level. That does not matter while no line asks about them.
- **Everyone running fork tests locally:** `go test` needs `-tags with_db`. Without it every test that loads gear panics with "No item with id".

## For the guide lane

- None. P1 moves no page value (the guides r22 shaman report agrees). P3 and P5 change no build, rank or prose.

## Still open

- P1 for damage spells: a sub-20 damage-rank tooltip read on a level 25 to 30 character.
- P2b (druid gates, handed over), P4 (Arcanite Champion PPM, Skyborne offsets), L1 (no live logs; last looked 07:26 UTC 5 Nov).
- From round 21, unchanged: P1b (Eureka! and Execute since 17 Oct), P6 (bear threat constants), and the round-20 web SHOULDs (relic row colour, "Hit is capped" wording, the Fury card's "with 30 seconds left").

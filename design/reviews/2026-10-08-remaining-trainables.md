# The remaining trainables: Flametongue Totem, Seal and Judgement of Light and Wisdom, Lightwell (2026-10-08)

Lane `trainables`. Fork branch `trainables` (four commits on `forever` at f3a673d10, head 6cc1e6fa2: the engine work, the synced priest rotations, a no-roll-at-full-health fix, a SUMMARY sentence), site branch `trainables` (rotations, vocabulary, rank pins, ladder goldens, this review). Both are unmerged. The site's `sim/go.mod` replace pointed at the fork worktree during the work and is not committed, so the site branch compiles and passes only against a fork commit that carries `flametongue_totem` (`make engine-pin`, then `python -m pipeline simproto --engine <fork>` at merge).

These were the last trainables the conformance report listed as affecting a simmed result and left unregistered (SUMMARY.md, "Trainable abilities the engine does not register"). Client rows are read from `data/builds/1.60.1.70009/raw`; each is quoted in the fork test that pins it.

## 1. Flametongue Totem (shaman)

### Client rows

- Cast spells 8227 / 8249 / 10526 / 16387 (ranks 1 to 4): `SpellLevels` 28 / 38 / 48 / 58, `SpellPower` 90 / 140 / 200 / 275 mana, `SpellCooldowns` StartRecoveryTime 1000, duration index 5 (300000 ms); effect 28 (summon) misc 5950 / 6012 / 7423 / 10557. Rank 1 text: "enchants all party members' main-hand weapons with fire if they are within 8230a1 yards"; ranks 2 to 4: "enhances the melee attacks of all party members within `$8250a1` yards. Each main hand hit causes `${$8248m1/77*$<mult>-1}` to `${$8248M1/25*$<mult>}` additional Fire damage, based on the speed of the weapon. Slower weapons cause more fire damage per swing."
- Area auras 8230 / 8250 / 10521 / 15036: effect 35 (apply party aura), aura 42 (proc trigger), `SpellAuraOptions` ProcChance 100, ProcTypeMask_0 4 (a melee auto attack), radius index 10 (30 yards).
- Proc spells 8253 / 8248 / 10523 / 16389: effect 3 (dummy), EffectBasePointsF **548 / 781 / 1061 / 1363**, no per-level growth, no variance, no coefficient.
- Flametongue Weapon (8024 to 16342): "When applied to main hand, disables any benefit you personally receive from Flametongue Totem."

### The weapon-speed scaling

The conformance summary said the client rows state the amount but not the weapon-speed scaling, so it was not guessed. It is stated, in the description: `m1/77` to `M1/25` is the amount per second of weapon speed over 1.3 to 4.0 seconds (100/77 and 100/25), so a hit deals **base points x weapon speed / 100**. Flametongue Weapon's proc states the identical range, and the engine's weapon imbue already agrees with it: the client's rank 6 row is 2498 + 78 per level above 56 = 2810 at level 60, 28.1 a second, against the imbue's 112 / 4 = 28. `$<mult>` is the Enhancing Totems talent (spell 29192), which is not in the live trees, so 1.

### Model

- `core.FlametongueTotemRanks` (`sim/core/flametongue_totem.go`), `core.FlametongueTotemAura`: a permanent aura on every ally in the raid buff `flametongue_totem` (RaidBuffs field 42, `NextIndex` 43). Each landed main-hand auto attack casts the proc spell on the target: fire, magic hit and crit, `Amount x SwingSpeed / 100`, no coefficient (the client states none). The highest rank the character's level has learned applies.
- Reach: not a shaman with Flametongue Weapon on the main hand (the weapon's text), not a feral druid (no main-hand weapon), the same two rules as Windfury Totem (`weaponTotemReaches` now serves both).
- Non-stacking: the client's beta notes (2026-10-08-beta-evidence.md, section 4) say "Windfury no longer stacks with ... Flametongue totem". The two auras share the exclusive category `Melee Weapon Totem`; Windfury has priority 2 and Flametongue 1, the order the preset already chose. A request that names both receives Windfury on every melee class and Flametongue only where Windfury cannot reach (an Enhancement shaman with Windfury Weapon). Pinned by `TestFlametongueTotemDoesNotProcBesideTheWindfuryTotem`.
- The shaman's cast (`sim/shaman/flametongue_totem.go`): a fire totem, 1000 ms GCD, 5 minutes, the client's mana. Casting it activates the same aura for the totem's life. Searing and Magma Totem and Flametongue Totem now share `endStandingFireTotem`, so a buff totem and a damage totem replace each other (a Searing cast after a Flametongue cast used to have nothing to cancel and would have panicked on the nil dot).
- Conformance: the four ranks read `match`, and the four proc spells read **declared, matches** (548 / 781 / 1061 / 1363). Effect 3 is a script hook in general, so the report reads a dummy amount only for these four ids and only when the engine declares the base points (`weaponSpeedDummySpells`, the absorb precedent); reading every dummy turned Mutilate's unrelated declaration into a `differs`, which is why it is a list.

### Decision: does the raid preset carry it beside Windfury?

Fixed Phase 1 BiS entries, the raid preset, 10,000 iterations (the ranker's `spec-breakdown`, three seeds identical: the sim is deterministic per iteration count). Three variants of the preset's totem: neither, Windfury (the preset as it stands), Flametongue instead.

| Spec | neither | Windfury | Flametongue | Windfury over Flametongue |
|---|---|---|---|---|
| rogue-assassination | 724.8 | 777.3 (+7.2%) | 744.8 (+2.8%) | +4.4% |
| rogue-combat | 697.3 | 745.7 (+6.9%) | 714.7 (+2.5%) | +4.3% |
| rogue-subtlety | 605.9 | 656.0 (+8.3%) | 625.0 (+3.2%) | +5.0% |
| warrior-arms | 622.5 | 795.4 (+27.8%) | 641.5 (+3.1%) | +24.0% |
| warrior-fury | 823.0 | 912.4 (+10.9%) | 844.8 (+2.6%) | +8.0% |
| warrior-protection | 269.7 | 306.0 (+13.5%) | 279.5 (+3.6%) | +9.5% |
| paladin-retribution | 607.2 | 688.8 (+13.4%) | 625.9 (+3.1%) | +10.0% |
| paladin-protection | 393.3 | 427.1 (+8.6%) | 405.6 (+3.1%) | +5.3% |
| hunter-survival | 682.3 | 833.2 (+22.1%) | 707.9 (+3.8%) | +17.7% |
| shaman-enhancement | 692.8 | 692.8 (+0.0%) | 716.2 (+3.4%) | -3.3% |

**Verdict: keep Windfury Totem, do not add Flametongue Totem to the preset.** Windfury is worth 4 to 24% more than Flametongue to every melee class that can receive it (the errors are 0.1 to 0.3 DPS). The one exception is the Enhancement shaman, who cannot receive Windfury (his own Windfury Weapon disables it) and gains 3.4% from Flametongue. A raid shaman drops both, one per slot, so the real raid is "both": every other melee class on Windfury, the Enhancement shaman on Flametongue; adding the entry to the raid preset reproduces exactly that (the "both" run equals the Windfury run for every spec but Enhancement, which reads 716.2). It is not added because the preset describes what every spec receives and an Enhancement shaman's own totems are his rotation's business, which is the next measurement.

Enhancement's rotation: replacing the Searing Totem line with Flametongue Totem (rank 4, refreshed at 1.5 s, one prepull cast) reads 692.9 +-0.55 against 692.8 +-0.54, Flametongue Totem's procs adding 23.5 DPS where Searing's bolts did. A tie, so no rotation change; the group value of Flametongue is nil beside the Windfury every other melee class takes.

## 2. Seal of Light, Seal of Wisdom, Judgement of Light and Wisdom (paladin)

### Client rows

- Seal of Light, ranks 1 to 4: cast spells 20165 / 20347 / 20348 / 20349, `SpellLevels` 30 / 40 / 50 / 60, `SpellPower` 110 / 140 / 180 / 210 mana, StartRecoveryTime 1500, duration index 9 (30 s). Effect 0 is aura 42 (proc trigger) naming 20167 / 20333 / 20334 / 20340, ProcChance 100, ProcTypeMask_0 **20** (a melee auto attack or a melee ability); effect 2 is aura 4 (dummy) naming the judgement 20185 / 20344 / 20345 / 20346. The heal spells: effect 10, EffectBasePointsF **39 / 53 / 76 / 94**, no coefficient.
- Seal of Wisdom, ranks 1 to 3: 20166 / 20356 / 20357, levels 38 / 48 / 58, 135 / 170 / 200 mana, the same aura 42 naming 20168 / 20350 / 20351 (effect 30, energize, EffectBasePointsF **50 / 71 / 90**); judgement 20186 / 20354 / 20355.
- Judgement of Light 20185 / 20344 / 20345 / 20346: duration index 31 (**40000 ms**), aura 42 with ProcTypeMask_0 40 (a melee hit taken), trigger 5373 ("Judgement of Light Intermediate", effect 3); the heals 20267 / 20341 / 20342 / 20343, effect 10, **25 / 34 / 49 / 61**, no coefficient. Improved Judgement of Light (23564): "Increases the chance of triggering a Judgement of Light heal by 10%". Redemption Armor's 2-piece (28775): "Increases the amount healed by your Judgement of Light by 20".
- Judgement of Wisdom 20186 / 20354 / 20355: 40000 ms, aura 42 with ProcTypeMask_0 139944 (every melee, ranged and spell attack taken), trigger 1826; the mana 20268 / 20352 / 20353, effect 30, **33 / 46 / 59**.

### What the rows do not say

The seal and judgement auras state ProcChance 100 and no procs-per-minute row, exactly as Seal of Command does (its row says 100 and the engine gives it 7 a minute from the server's script). So two numbers are **assumptions**, named constants with the reason beside them:

- `paladin.UtilitySealProcsPerMinute` = 15, the rate vanilla's seals ran on. A seal heals or restores 15 times a minute, not on every hit.
- `core.JudgementProcChance` = 0.5, vanilla's chance for Judgement of Light and Wisdom (the engine's Judgement of Wisdom already used 0.5).

Both want beta evidence; neither moves a DPS number except the mana users' Judgement of Wisdom, whose size scales with the chance.

### Model

- `sim/paladin/utility_seals.go`: both seals share one registration (`utilitySeal`, `registerUtilitySeal`): the cast at every learned rank (client cost, 1.5 s GCD, 30 s aura as the spell's own buff, Twist of Light's mask), a PPM proc on any landed melee hit, a judgement spell per rank that rolls the spell hit table and lays the judgement for 40 s. Seal of Light's trigger casts the client heal spell 20167 to 20340 on the paladin (declared base points); Seal of Wisdom adds the mana filed under the client's energize spell id.
- `core.JudgementOfWisdomAura` and `core.JudgementOfLightAura` last 40 s (the engine kept vanilla's 10); Wisdom's mana is `JudgementOfWisdomRanks.At(attacker level)` (59 at 60, 33 at 38, where it was a flat 59); Light is new: a landed melee hit on the target heals a hurt attacker `JudgementOfLightRanks.At(level)` (61 at 60). **The `judgement_of_light` raid debuff did nothing before and now does this.** A heal on a full-health unit is neither rolled nor recorded, so a damage sim whose melee attackers are never hurt draws no random numbers for it.
- Conformance: the seal casts (7 rows a level), the seven judgement spells and the four Seal of Light heals read `match` / `declared, matches`. The heal rows' duration reads `client-scripted` (the client gives the heal spell none, the engine's sibling-id match finds the seal's 30 s).

### Decisions

**Judgement of Wisdom stays in the raid preset's debuffs.** Raid preset, with and without the debuff:

| Spec | without | with | gain |
|---|---|---|---|
| hunter-marksmanship | 759.4 | 810.7 | +6.8% |
| hunter-beast-mastery | 783.0 | 829.7 | +6.0% |
| shaman-elemental | 463.8 | 488.0 | +5.2% |
| paladin-retribution | 661.3 | 688.8 | +4.2% |
| shaman-enhancement | 681.1 | 692.8 | +1.7% |
| warlock-demonology / destruction / affliction | 859.2 / 826.1 / 908.3 | 873.4 / 839.4 / 922.3 | +1.7% / +1.6% / +1.5% |
| mage-fire | 746.2 | 758.2 | +1.6% |
| hunter-survival | 822.7 | 833.2 | +1.3% |
| mage-frost | 567.9 | 570.7 | +0.5% |
| druid-balance, mage-arcane, priest-shadow | | | within 0.2% |

It moves every mana-limited attacker by 1 to 7%, so a raid paladin judging it is worth a slot. Healers gain nothing (their spells are not used against the judged target), and Judgement of Light moves no damage number: it heals the melee attacker, which only the healing metrics see.

**Retribution and Protection keep judging their damage seal (Righteousness, Fury).** Retribution under the preset without Judgement of Wisdom, three rotations: the curated one (661.3); the curated one with a Seal of Wisdom cast whenever the debuff has 10 s or less and Judgement is ready, then a recast of Righteousness (678.3, +2.6%: 1.1 Wisdom twists a fight); Seal of Wisdom held permanently and judged on cooldown (514.1, -22%). So when the ret is the raid's only paladin, a twist once every 40 seconds recovers 17 of the 27.5 DPS the debuff is worth, and that is a statement about a raid with one paladin: the preset's premise is that a raid paladin supplies Judgement of Wisdom, under which the twist is idle (the debuff is permanent, the Wisdom line never fires; the +0.4% that run shows, 691.3 against 688.8, is the changed seal-recast line, not Wisdom, and is not adopted). Protection's Judgement is a threat source (its own review: 7% less threat without it), so it keeps the Fury judgement. No paladin rotation changes.

## 3. Lightwell (priest)

### Client rows

Spells 724 / 27870 / 27871 (ranks 1 to 3): `SpellLevels` 40 / 50 / 60, `SpellPower` 225 / 295 / 365 mana, `SpellCastTimes` 1500 ms, StartRecoveryTime 1500, CategoryRecoveryTime **600000**, duration index 26 (**180000 ms**); effect 50 (summon object) 181102 / 181105 / 181106. "Creates a holy Lightwell near the priest. Members of your raid or party can click the Lightwell to restore `$7001o1` health over `$7001d`. Being attacked cancels the effect. Lightwell lasts for `$d` or 5 charges." The renew spells 7001 / 27873 / 27874: effect 6 aura 8 (periodic heal), EffectBasePointsF **160 / 233 / 320** every 2000 ms, duration index 8 (10000 ms), no coefficient. SkillLineAbility 13838 / 13850 / 13851 (Holy, class mask 16).

### Model (`sim/priest/lightwell.go`, healing specs only)

- The cast leaves a "Lightwell" aura: 180 s, 5 stacks. The renew is a registered heal-over-time (5 ticks of 2 s, flat per tick, the caster's healing multiplier, no spell-power term).
- The fake raid members use it. The raid damage model gained `Raid.OnRaidDamage` (a listener told each hit: the unit, the damage after absorbs, whether it is the tank), which Prayer of Mending's comment said a healer lacked. A member the model hurts first **loses any Lightwell renew it carries (being attacked cancels it)**, then, if it is not the tank (who is attacked every swing) and is under 90% health and the well has a charge, **clicks at once**: one charge, one renew.
- Assumptions: the 90% click threshold, and that a click takes no time. The renew spells 7001 / 27873 / 27874 sit outside the spellconst extract (they carry no class family), so the conformance report has no damage row for them; `TestEachLightwellChargeHealsFiveTicksOf320` pins 25 ticks of 320 (8000) for five uninterrupted clicks, and `TestBeingAttackedCancelsALightwellRenew` pins the cancel.
- Conformance: the three ranks read `match` (cost, 1500 ms cast, 1500 ms GCD, 600000 ms cooldown, level, 180000 ms aura).

### Offered to the two priest rotations

The healer rotations are not searched by `rotation-search` (it skips every non-DPS role); a healer's measure is the healing fight. `spec-breakdown -heal`, 10,000 iterations (the figures repeat to 0.1 at five different iteration counts), the priest's Lightwell as an opening line against the list as written, and as a prepull cast 2.5 s before the pull:

| Spec | effective HPS | mana lasts (s) |
|---|---|---|
| priest-holy, as written | 691.9 | 249.2 |
| priest-holy, Lightwell first line | 694.4 (+0.36%) | 252.8 |
| priest-holy, Lightwell prepull | 694.9 (+0.43%) | 253.0 |
| priest-discipline, as written | 604.9 | 205.0 |
| priest-discipline, Lightwell first line | 613.3 (+1.4%) | 213.2 |
| priest-discipline, Lightwell prepull | 613.0 (+1.3%) | 212.8 |

Both gain beyond the noise and both last longer on mana (the well's healing replaces heals that would have been cast). **Adopted: a prepull cast of rank 3 at -2.5 s in both lists** (the prepull and the opener tie; a player places the well before the pull, and it costs no combat GCD). It is one cast a fight: the cooldown is 10 minutes. Ladder goldens `priest-holy` and `priest-discipline` regenerated (level 60 bare healing 131.5 to 139.4 and the matching Discipline move; levels 40 and 50 gain where the well is learned).

## 4. Tests, goldens, ranker

- Fork: `go test --tags=with_db ./sim/...` (39 packages) passes. **No `.results` moved**: none of the changes touch an existing rotation's casts, and the debuffs read the same at level 60. Conformance goldens regenerated (`shaman`, `paladin`, `priest`, `SUMMARY.md`): the damage table's declared rows 770 to 786, all matching and none differing, and the unregistered trainables 194 to 190 (Flametongue Totem, Seal of Light, Seal of Wisdom, Lightwell leave the list).
- Site: `make apl-sync` / `apl-check` against the fork worktree (the two healing priest rotations), `go test ./...` under `sim/` passes after the two priest ladder goldens were regenerated (they were the only ones to differ), `sim/leveling/buff_ranks_client_test.go` pins the Flametongue and Judgement of Light tables to the client, `sim/request/IDS.md` lists `flametongue_totem`.

### Ranker, band 60 (before and after)

`sim/cmd/leveling-bis -spec <spec> -bands 60`, headline `set_dps` of each entry (for a healer the ranker's own figure for the set), gear re-ranked in both runs. Before is the fork at f3a673d10 with the site at a3108a84; after is both `trainables` branches. Bare is the class kit alone, raid the Phase 1 raid preset.

| Spec | Faction | Preset | Before | After | Change |
|---|---|---|---|---|---|
| priest-holy | alliance | bare | 331.2 | 342.2 | +3.31% |
| priest-holy | horde | bare | 323.2 | 333.2 | +3.11% |
| priest-holy | alliance | raid | 707.6 | 710.0 | +0.35% |
| priest-holy | horde | raid | 682.9 | 684.7 | +0.26% |
| priest-discipline | alliance | bare | 373.4 | 387.7 | +3.83% |
| priest-discipline | horde | bare | 367.7 | 372.3 | +1.25% |
| priest-discipline | alliance | raid | 628.3 | 642.5 | +2.26% |
| priest-discipline | horde | raid | 622.6 | 639.0 | +2.64% |
| paladin-protection | alliance | raid | 442.2 | 440.1 | -0.49% |
| paladin-protection | horde | raid | 484.5 | 484.5 | +0.01% |

Every other entry of the nine specs is identical to the decimal: shaman-enhancement, shaman-elemental, shaman-restoration, paladin-retribution, paladin-holy and priest-shadow in both factions and both presets, and paladin-protection bare. The two healing priests gain from the Lightwell (and the gear it lets the ranker choose: 4 to 12 slots differ by entry); the raid Protection paladin moves because Judgement of Light now heals the one melee attacker who is hurt, the tank, which changes what the fight's damage model sees, and 2 to 3 slots differ in the raid entries, within half a percent. Nothing else moved: the new totem, seals and well are not in any other rotation, and the full-health rule keeps the debuff from drawing random numbers in a sim nobody is hurt in.

## 5. Follow-ups

- **Evidence wanted for the two assumed rates**: Seal of Light / Wisdom at 15 a minute and Judgement of Light / Wisdom at 50% (one beta log of either settles them); the Lightwell click threshold.
- **Not modelled**: Improved Judgement of Light (23564), the Redemption Armor 2-piece (28775, +20 Judgement of Light heal; the set report lists it "registered, not verified"), the Lightwell for the damage-taken side of a tank (the well is clicked only by non-tank members), Spiritual Healing and the other heal multipliers reaching the Lightwell renew (it takes the caster's multiplier, not the spell-mask talents).
- **Beta notes say Windfury also no longer stacks with Tranquil Air / Grace of Air**: the raid preset still carries Windfury Totem and Grace of Air Totem together and the engine applies both. This lane only added the Flametongue half of the rule.
- Flametongue Weapon's engine proc carries a 0.1 spell-power coefficient the client does not state (the client states none on 8026 to 16344 either); outside this lane.
- `ui/core/components/inputs/totem_inputs.ts` and the generated TypeScript proto do not know `flametongue_totem` yet (`make proto` needs node_modules).

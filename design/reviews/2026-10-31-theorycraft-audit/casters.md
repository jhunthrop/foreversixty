# Theorycraft audit: casters (2026-10-31)

Group: Mage (Arcane, Fire, Frost), Warlock (Affliction, Demonology, Destruction), Priest (Shadow; Discipline and Holy are healer-ranked, see the end).

## Method

- **Our side.** I read the guide `raidBuild` codes (decoded against `data/builds/1.60.1.70658/talents/*.json`), the APLs in `data/curated/apl/`, the band-60 raid entries in `data/builds/1.60.1.70658/bis/*.json`, and the pinned engine `816ec57c0` (`git archive` of `/Users/jh/code/wowsims-forever`; the fork checkout itself sits on an older sha).
- **Sim runs.** `sim/cmd/spec-breakdown` with a scratch `go.work` on the pinned engine: raid preset, band 60, horde (the published entries are horde), 10,000 iterations. Every gain called confirmed was re-run at seed 50001 and agreed to within 0.3 DPS. Errors are about ±0.3 DPS. Variant rotations change only the line named.
- **Engine fixes.** Where an engine bug inflates a number, I patched a scratch copy of the engine only to measure it. Nothing in the repo or the fork was changed.
- **Client truth.** Talent text comes from `talents/*.json` (client 70658). Learnability comes from `SkillLineAbility.csv` (70532, the newest raw copy) and `trainables/*.json`.
- **Web sources.** Level-60 Forever theorycraft is thin: the beta capped at 30, and most guides say "level 60 at launch". Icy Veins and Mobalytics returned 403. Wherever a source exists, it is cited per row. Era-only material is used only for mechanics that the Forever client text keeps.

Baselines (pinned engine, seed 1): Fury 797; Destruction 772; Affliction 737; Fire 722; Demonology 689; Arcane 562; Shadow 519; Frost 460.

## Cross-cutting engine bugs

### X1. CONFIRMED-BUG: a wand Shoot cast inside a GCD shortens that GCD

- **Mechanism.** `core/wand.go` registers Shoot with `GCD: 0` and a cast time equal to the wand's speed (1.2 s on the BiS wands). `core/cast.go:213` then calls `SetGCDTimer(now + effectiveTime)`. It only takes the max with the GCD still running when the spell carries `SpellFlagCastTimeNoGCD`, and Shoot does not.
- **What happens.** After every instant (Corruption, Wrack, Life Tap, Shadow Word: Pain), the APL's Shoot line fires in the same instant. It deals a free wand hit and lets the next spell start 0.3 s early.
- **Log evidence.** Affliction, one iteration of a debug run:
  - Wrack at 27.05, Shoot 27.05 to 28.25, Shadow Bolt at 28.25, before the 1.5 s GCD ends at 28.55.
  - Life Tap at 39.65, Shoot, Wrack at 40.85.
- **Side effect.** Wand time also restarts spirit regen, so Life Tap casts fall from about 25 to 18 a fight.
- **Size of the effect.** Measured by deleting the Shoot line, which is a lower bound because legitimate wanding is worth a little:

| Spec | With Shoot | Without | Inflation |
|---|---|---|---|
| Destruction | 772.2 | 706.5 | +9.3% |
| Affliction | 736.6 | 675.0 | +9.1% |
| Shadow | 519.2 | 501.0 | +3.6% |
| Demonology | 688.5 | 674.8 | +2.0% |
| Fire | 722.2 | 717.9 | +0.6% |
| Frost | 459.9 | 458.2 | +0.4% |
| Arcane | 561.5 | 561.5 | 0 |

- **Knock-on effects.**
  - Destruction's #2 and Affliction's #3 tier positions are partly this artifact.
  - Wand Specialization 2/2 sits in the Frost, Fire and Shadow raid builds; the talent search valued it on inflated wand uptime.
  - It distorts rotation choices. With Shoot, a Shadow Bolt filler beats Searing Pain for Demonology by 1.4%; without Shoot, Searing Pain wins by 1.6%.
  - It is the likely source of the significant negative haste weights (OPEN, below).
- **Fix.** Give Shoot `SpellFlagCastTimeNoGCD`, or refuse to start it while the GCD runs. Then re-run the ranker and the talent searches. Whatever Shoot's own GCD rule in the client, a 1.2 s wand shot can never end a 1.5 s GCD early.
- **Applies beyond this group.** The bug applies to every class with a Shoot line (hunters and rogues fire their ranged weapon through `EnableAutoAttacks`, not this path).

### X2. CONFIRMED-BUG: Missile Barrage is consumed by the second Arcane Missiles, not the first

- **Mechanism.** In `mage/missile_barrage.go:74`, `OnCastComplete` skips the consume when `RemainingDuration == Duration`. It is meant to protect the cast that granted the buff. But Arcane Missiles never grants Missile Barrage, and an Arcane Missiles started at the same timestamp as the proc completes in that same instant (channel cast time 0). That first Arcane Missiles keeps the buff, and the next one is free as well.
- **Evidence.** Frost debug log: a single Frostbolt proc at 6.82 is followed by free channels at 6.82 and at 9.32. Arcane casts 16.15 Arcane Missiles a fight; the expected number of procs is about 8.3 (20.3 Arcane Blasts × 20% + 42.4 Frostbolts × 10%).
- **Size of the effect.** With the guard removed (scratch engine), Arcane drops from **561.5 to 513.5 (−8.5%)** and Arcane Missiles casts fall to 9.0.
- **Related.** The Clearcasting guard in `applyArcaneConcentration` has the same pattern. It is worth checking for instants (Ice Lance) cast at the proc timestamp.

## Mage: Arcane

Raid build: Arcane 32 (Wand Specialization 1, Arcane Focus 5, Improved Channeling 3, Missile Barrage 1/2, Arcane Concentration 5, Arcane Impact 3, Arcane Blast, Presence of Mind, Arcane Mind 5, Arcane Instability 3, Arcane Power, …) / Fire 3 / Frost 16.

APL: Arcane Power and Presence of Mind on cooldown > Arcane Missiles on Missile Barrage > Arcane Blast to 2 stacks > Frostbolt rank 10 > wand.

| # | Difference | Class | Evidence |
|---|---|---|---|
| A1 | Missile Barrage is 1/2, with 1 point in Wand Specialization. Missile Barrage 2/2 in its place gains **+10.1%** (513.5 → 565.2, scratch engine with X2 fixed). On the shipped engine it reads +13.7%, inflated by X2. | CONFIRMED-BUG (build) | Client text: "Arcane Blast 40%, … Frostbolt … 20% chance". Warcraft Tavern Mage (beta) calls Missile Barrage the Arcane engine. Talent string `0532005113100101531-03-0535000300`. |
| A2 | The published DPS is inflated about 8.5% by X2. | CONFIRMED-BUG (engine) | X2 |
| A3 | Spell haste weighs −13.7 SP per 1% (error 0.8), significant and negative. Theorycraft lists haste above intellect. | OPEN | Probably mana compounding from the 175% Arcane Blast cost and X2's free channels. To settle: re-measure after X1 and X2. |
| A4 | The Arcane Blast stack cap is 2 in the APL; the client allows 4 stacks (+15% of intellect as spell damage each, +175% cost each). | OPEN | No Forever source gives a stack target. The APL note records a measured choice; re-run rotation-search after X2. |
| A5 | Frostbolt rank 11 (25304) is a book rank, so the filler uses rank 10. | FOREVER-CHANGE (as data says) | `book-ranks.json` and `core.IncludeAQ` |

## Mage: Fire

Raid build: Arcane 14 (Wand Specialization 2, Improved Channeling 5, Arcane Concentration 5, …) / Fire 34 (Ignite 5, Pyroblast, Improved Scorch 3, Heating Up, Master of Elements 3, Critical Mass 5, Fire Power 3, Combustion, …) / Frost 3.

APL: Combustion at 5 Scorch stacks > Scorch upkeep > Pyroblast at 3 Heating Up stacks > Fire Blast > **Scorch filler** > wand.

| # | Difference | Class | Evidence |
|---|---|---|---|
| F1 | Theorycraft uses Fireball as the filler, with Fire Blast weaved and Scorch only for the debuff. We fill with Scorch. A Fireball rank 11 filler measures **−8.6%** (660.2) because it runs out of mana for 11.6 s of 180. | OPEN | Warcraft Tavern Mage: "Incorporate instant casts such as Fire Blast after each Fireball cast"; Scorch "to debuff the target in long fights". Single Forever source; Era theorycraft agrees. To settle: a Forever level-60 fire parse's mana curve, or check that the Fireball DoT counts toward Fire Power (+3% per Fire effect) and that Master of Elements refunds are modelled. The engine has both, so this is likely a real Forever mana difference, not a bug. |
| F2 | The APL's Fireball line names 25306 (rank 12, a book rank, unregistered), so the line is inert. | CONFIRMED-BUG (APL, no DPS effect) | `spellranks.json`: 25306 `book: true`. The variant with only that line moved measured exactly the baseline, 722.2. |
| F3 | Frostfire Bolt is trainable (40/50/60, SkillLineAbility row 53024) but unused. As the filler it measures −12.8%. | Aligned (correctly unused) | Run `mage-fire-ffb` |
| F4 | Wand Specialization 2/2 sits in a raid build. | Follows X1 | Re-run the talent search after X1. |
| F5 | Tier: Fire is the top mage at about 91% of Fury. | Aligned | Matches the Era ratio; no Forever ranking exists yet. |

## Mage: Frost (re-checked after caster rounds 1 and 2)

Raid build: Arcane 10 (Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5) / Fire 7 / Frost 34 (… Fingers of Frost 2, **Winter's Chill 5**, Ice Barrier).

APL: Ice Barrier upkeep > Ice Lance on Fingers of Frost > Cone of Cold (in reach only) > Frostbolt > wand.

| # | Difference | Class | Evidence |
|---|---|---|---|
| R1 | Theorycraft takes Missile Barrage (it calls it "mandatory", alongside Arcane Focus and Arcane Impact) and casts Arcane Missiles on every proc. Our build has no Missile Barrage and the APL has no Arcane Missiles line. Missile Barrage 2/2 plus an "Arcane Missiles on proc" line gains **+6.7%**. Adding Winter's Chill at 1 point and Arcane Focus with the freed points (talents `0532005-13102-253510130101301211`) gains **+8.3% (459.9 → 497.9)**. These figures are on the X2-fixed scratch engine; on the shipped engine they read +11.7%. | CONFIRMED-BUG (build + APL) | Warcraft Tavern Frost Mage: "Use Arcane Missiles whenever Missile Barrage procs"; Winter's Chill "at 1 point"; Missile Barrage, Arcane Focus and Arcane Impact "mandatory". Client text gives Frostbolt a 20% proc chance. |
| R2 | Winter's Chill 5/5 has no value at 30 yd: it only feeds the mage's own Cone of Cold, which round 2 made out of reach. Four points are spent on nothing. | CONFIRMED-BUG (build) | Client text: "increasing the target's damage taken from your next Cone of Cold". Round-2 review. Ice Barrier still has its 30 tree points with Winter's Chill at 1. |
| R3 | The Ice Lance on Fingers of Frost, Ice Barrier upkeep and Frostbolt filler match theorycraft. | Aligned | Warcraft Tavern; wowtbc Frost Mage ("Ice Lance casts on those procs, then Frostbolt") |
| R4 | Frostfire Bolt as the filler measures −15%. | Aligned | Run |
| R5 | Haste weighs +6.9 SP per 1%; theorycraft lists haste second. | Aligned | |

## Warlock: Affliction

Raid build: Affliction 33 (Improved Life Tap 2, Suppression 5, Improved Corruption 5, Grim Reach 2, Improved Drains 3, Amplify Curse, **Pandemic 3**, Nightfall 2, **Siphon Life 1**, Soul Siphon 3, Shadow Mastery 5, Wrack) / Demonology 7 / Destruction 11 (Improved Shadow Bolt 5, Bane 5, Ruin 1). Succubus out.

APL: Life Tap below 10% > Corruption > **Bane of Doom** > Wrack > Immolate > Shadow Bolt > wand below 50% mana.

| # | Difference | Class | Evidence |
|---|---|---|---|
| W1 | The build takes Siphon Life, but the APL never casts it. Theorycraft maintains it. Siphon Life is an Affliction effect, so it lifts Soul Siphon on Wrack from 24% to 36%. Adding Siphon Life upkeep gains **+2.9%** (758.1; +2.0% without the wand line). | CONFIRMED-BUG (APL) | wowtbc Affliction: "Maintain Siphon Life". Engine `soulSiphonMultiplier` (warlock/talents.go:1033) counts Affliction DoTs. |
| W2 | Theorycraft keeps Bane of Agony; we keep Bane of Doom. With W1 in place, Bane of Agony gains **+0.6%** more (762.6). Wrack's "+10% damage from your other Shadow damage-over-time effects" favours the ticking curse. | CONFIRMED-BUG (APL, small) | wowtbc Affliction: "Maintain Bane of Agony"; client Wrack text. |
| W3 | Pandemic 3/3 adds crit damage to Corruption, Bane of Agony, Bane of Doom, Siphon Life and Wrack, but the engine ticks those with `OutcomeTick` (0.0% crit in every breakdown). Pandemic is inert, and Shadowform's crit bonus on Shadow Word: Pain and Mind Flay is also unreachable. Forever text implies periodic crits exist: Pandemic would be useless otherwise. | OPEN (likely bug) | Client Pandemic text. wow.gg Affliction (2026-09-19): "In Forever, periodic damage learned to crit." That is the only source. To settle: a Forever combat log showing a Corruption or Shadow Word: Pain tick crit. |
| W4 | Malevolence 5/5 (+5% Shadow crit) is not taken. | OPEN | Re-run the search after W3 and X1. |
| W5 | Haunt (403501/1293693/1293694) and Unstable Affliction (427717/1242971/1242972) appear in `trainables` as `class_spell`, but have no SkillLineAbility row in any build (70009 to 70532). | FOREVER-CHANGE (correctly unused) | `SkillLineAbility.csv` grep |
| W6 | The tier position (#3) is about +9% inflated by X1. Without the wand line, plus W1 and W2: 690.9. | CONFIRMED-BUG (via X1) | |

## Warlock: Demonology

Raid build: Affliction 10 / Demonology 30 (… Demonic Sacrifice, Decimation 2, Soul Link, Demonic Knowledge 3, Master Demonologist 5) / Destruction 11. Succubus out, Soul Link before the pull.

APL: Soul Fire on cooldown (Decimation) > Life Tap > Corruption > Bane of Doom > Immolate > **Searing Pain** filler > wand.

| # | Difference | Class | Evidence |
|---|---|---|---|
| D1 | Searing Pain is the filler, while the Succubus' Master Demonologist buffs Shadow. Without X1, Searing Pain beats Shadow Bolt by 1.6% (674.8 against 664.4, both seeds). With X1 the order flips. | Aligned (after X1) | Runs `demo-sb-noshoot` and `warlock-demonology-noshoot` |
| D2 | Demonic Pact (capstone) is unreachable at 30 tree points. Warcraft Tavern's Forever guide describes the intended play: sacrifice the Imp (+15% Shadow) and then summon the Succubus (+10% Shadow from Master Demonologist). The engine has the Pact hook (`talents.go`, the `DemonicPact` check in `applyDemonicSacrifice`), but the APL cannot sacrifice one pet and summon another, and the request fixes one summon. | OPEN | Warcraft Tavern Warlock. To settle: an APL summon action plus a talent-search run with 31 Demonology points. |
| D3 | Inflation from X1 is about 2%. | Via X1 | |

## Warlock: Destruction

Raid build: Affliction 12 / Demonology 5 / Destruction 34 (Bane 5, Aftermath 4, Ruin 5, Shadowburn, Agonizing Flames 3, Conflagrate, Bane of Havoc, Fire and Brimstone 3, Shadow and Flame 5, Incinerate). Succubus out.

APL: Corruption > Life Tap > Bane of Doom > Immolate > Conflagrate > Shadowburn on cooldown > Incinerate > wand.

| # | Difference | Class | Evidence |
|---|---|---|---|
| S1 | Shadowburn is cast on cooldown with no shard limit. | FOREVER-CHANGE (correct) | Client Shadow and Flame: "Shadowburn has a 100% chance to instantly refund a Soul Shard" |
| S2 | An Era-style sacrifice build (Succubus sacrifice +15% Fire; 6/11/34, talents `15-0050203001-2053045103101351`) loses 14% to the pet-out build. | Aligned | Run `warlock-destruction-ds`; Forever demons scale with gear (Warcraft Tavern Warlock). |
| S3 | The #2 tier position is inflated by X1 (+9.3%); without the wand line Destruction reads 706.5, below Fury, Beast Mastery, Marksmanship and Fire. | CONFIRMED-BUG (via X1) | |
| S4 | Haste weighs +0.72/1%; no Forever source. | Aligned | |

## Priest: Shadow

Raid build: Discipline 18 (Power in Light 3, Wand Specialization 2, Twin Disciplines 5, Mental Agility 3, Inner Focus, Improved Power Word: Shield 1, Meditation 3) / Shadow 33 (Shadow Focus 5, Blackout 2, Spirit Tap 3, Improved Mind Blast 5, Improved Shadow Word: Pain 2, Shadow Reach 1, Mind Flay, Improved Mind Flay 2, Vampiric Embrace, Shadow Weaving 3, Early Demise 2, Darkness 5, Shadowform).

APL: Shadow Word: Pain > Shadow Word: Death (execute) > Mind Blast (execute) > Inner Focus + Devouring Plague > **Mind Flay > Mind Blast** > wand.

| # | Difference | Class | Evidence |
|---|---|---|---|
| P1 | Mind Flay sits above Mind Blast. Theorycraft orders Shadow Word: Pain > Devouring Plague > Mind Blast > Shadow Word: Death > Mind Flay (two independent sources). Moving Mind Blast above Mind Flay gains **+6.1%** (519.2 → 550.7, both seeds). | CONFIRMED-BUG (APL) | Warcraft Tavern Shadow Priest (October 15 patch build); wowtbc Shadow Priest. Mind Blast casts go from 5.3 to 21.8 a fight. |
| P2 | Devouring Plague is cast only behind Inner Focus (1.45 casts a fight). The client has a 24 s cooldown, so it should be kept up (7.5 casts). The full theorycraft order gains **+14.7%** (595.4; +14.0% without the wand line, 501.0 → 570.9). | CONFIRMED-BUG (APL) | Trainables: Devouring Plague `cooldown_ms 24000`, engine `devouring_plague.go:64`. Both sources say "reapplied when the duration expires". |
| P3 | Mind Blast has no `ClassSpellMask`, so Shadowform's `ClassSpellsOnly` mods skip it: −50% mana cost and +100% crit damage. Debug log: Mind Blast costs 350 (full) while Mind Flay, Shadow Word: Pain and Shadow Word: Death cost 40% of base. Estimated +1 to 2% once Mind Blast is a 20% share. | CONFIRMED-BUG (engine) | Client Shadowform text: "all Shadow spells"; `priest/mind_blast.go` (no ClassSpellMask); `priest/talents.go:519,525`; `core/spell_mod.go:160` |
| P4 | Spell haste weighs −8.3 SP per 1% (error 0.36), significant. `mind_flay.go` has `AffectedByCastSpeed: false`. Theorycraft does not cover haste. | OPEN | Likely X1 interacting with the GCD. To settle: re-measure after X1; then check that Forever channels scale with haste. |
| P5 | Crit is worth 0.74 SP per 1%. It follows from non-critting damage over time and channels (see W3) and the Shadowform crit miss (P3). | OPEN (follows W3, P3) | |
| P6 | Shadowfiend (401977) has a SkillLineAbility row (line 78, AcquireMethod 0) but SpellLevels BaseLevel 0. | OPEN | It could be a trainer spell or a dead Season of Discovery copy. To settle: check Wowhead Forever or the trainer in game. |
| P7 | Tier: Shadow at 65% of Fury. With P1 and P2 it rises to about 75%. | Follows P1/P2 | |

## Priest: Discipline and Holy

Both are healer-ranked (heal profile), not DPS-tiered. Their weights (MP5 > spirit/intellect > healing power) are plausible for a mana-limited healing fight. No Forever level-60 healer theorycraft was found. I did not audit them further: OPEN.

## Ranked confirmed bugs (by DPS impact)

1. **X1, wand Shoot shortens the GCD** (engine). Inflates Destruction by 9.3%, Affliction 9.1%, Shadow 3.6% and Demonology 2.0%; it skews the tier order, Wand Specialization picks, rotation choices and probably the haste weights. `core/wand.go` plus `core/cast.go:213`.
2. **P1 + P2, Shadow APL priority** (Mind Blast below Mind Flay; Devouring Plague only behind Inner Focus): +14.7%.
3. **A1 + X2, Missile Barrage**: the engine double-consumes procs (Arcane −8.5% once fixed), and Arcane needs 2/2 (+10.1%).
4. **R1 + R2, Frost**: no Missile Barrage or Arcane Missiles line, and Winter's Chill 5/5 wasted at range: +8.3%.
5. **W1 + W2, Affliction**: Siphon Life never cast; Bane of Agony over Doom: +3.5%.
6. **P3, Shadowform misses Mind Blast** (engine): about +1 to 2%.
7. **F2, inert Fireball rank-12 line**: no DPS effect.

## Sources

- Warcraft Tavern, Shadow Priest Guide (Forever; build image for the October 15 patch): https://www.warcrafttavern.com/forever/guides/shadow-priest/
- wowtbc.gg, Shadow Priest PvE Guide, WoW Forever (undated, 2026): https://wowtbc.gg/warcraftforever/class-guides/shadow-priest/
- wowtbc.gg, Affliction Warlock PvE Guide, WoW Forever (undated, 2026): https://wowtbc.gg/warcraftforever/class-guides/affliction-warlock/
- Warcraft Tavern, Warlock Guide (beta 1.60.1.69913, level-20 cap): https://www.warcrafttavern.com/forever/guides/warlock/
- Warcraft Tavern, Affliction Warlock Guide (level-30 beta build only): https://www.warcrafttavern.com/forever/guides/affliction-warlock/
- wow.gg, Affliction Warlock specialization changes (2026-09-19): https://wow.gg/guides/warlock-affliction-forever-overview
- Warcraft Tavern, Frost Mage Guide (beta, images October 2026): https://www.warcrafttavern.com/forever/guides/frost-mage/
- Warcraft Tavern, Mage Guide (beta): https://www.warcrafttavern.com/forever/guides/mage/
- Icy Veins Forever PvE guides (Arcane, Fire, Frost, Affliction): 403 to the fetcher, not used.

Scratch rotation variants and the engine copies are in `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/tc/`. That directory is temporary; this file records every number.

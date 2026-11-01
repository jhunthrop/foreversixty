# Sim audit round 2: casters (Priest and Mage), 2026-11-01

Specs: priest-shadow, priest-discipline, priest-holy, mage-arcane, mage-fire, mage-frost. Warlocks are another lane.

## Method

- **Engine.** The lane branch `lane/r2-casters` starts at the pin `17d31ecfa`. Every "after" figure runs on the lane engine through a scratch `go.work`.
- **Client data.**
  - Client 1.60.1.70707. Raw tables were fetched from `wago.tools/db2/<table>/csv?build=1.60.1.70707` into the lane scratch directory: SkillLineAbility, SpellEffect, SpellMisc, SpellLevels, SpellName, SpellCooldowns and SpellClassOptions.
  - Talent text comes from `data/builds/1.60.1.70707/talents/*.json`.
  - Server text comes from the Wowhead Forever tooltips (`nether.wowhead.com/forever/tooltip/spell/<id>`).
- **Runs.**
  - `sim/cmd/spec-breakdown`: horde, raid preset, band 60, 10,000 iterations at seeds 1 and 50001. Both seeds agree to within 0.3%.
  - Healers run with `-heal`.
  - Talent searches run `sim/cmd/talent-search -preset raid -keep <raid-keeps.json>`.
- **Tests.** Every confirmed engine bug has a failing test written first. The test was confirmed to fail, then pass after the fix.
  - Fork packages: `go test -tags=with_db ./sim/mage/ ./sim/priest/...`. The item database needs `with_db`.
  - Site: `sim/request` (except TestRotationLadder), `make apl-check`, the data APL pytest, and the guide vitest.

## Cross-cutting

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| C1 | Three Shadow checks need re-measuring: the round-1 haste weights (P4, A3), whether channels scale with haste, and the Missile Barrage and Clearcasting guards (X2). | **Haste.** No trained channel carries a haste attribute: Mind Flay 18807 and Arcane Missiles 25345 both have SpellMisc Attributes_5 = 0, with neither 0x2000 nor 0x20000 set. The engine does not haste channels either (`AffectedByCastSpeed: false`). Nightly weights on 17d31ecfa, raid band 60, per 1% haste: Shadow +0.96 ± 0.21 and Arcane +3.54 ± 0.80. Both are positive and small, because Mind Blast's and Arcane Blast's 1.5 s casts already sit at the GCD floor. **Clearcasting guard.** The Missile Barrage guard is gone since c99fb94b4. The Clearcasting guard (`mage/talents.go:185`) only skips the cast that granted the proc, and no off-GCD mage instant can fire at that timestamp. | Aligned (P4, A3, X2 closed) | - |
| C2 | Shadowfiend 401977 (round-1 P6). Its SkillLineAbility row (49886, line 78 Shadow Magic, AcquireMethod 0) has no SpellLevels BaseLevel. | **For learnable:** the row is a class-line trainer row, not a rune row. The SoD rune is 402861 "Engrave Pants - Shadowfiend" on line 2851, and SoD class-line runes carry AcquireMethod 3. Heart of the Lion 409580 has the same shape (line 51, base 0) and is live in Forever. Wowhead Forever and wow-forever.gg both list it at "Requires level 1". wow-forever.gg logs a Forever-build text change on 20 September 2026 (6% to 5%), so Blizzard edited the spell for Forever. **Against:** no Forever guide mentions it; Warcraft Tavern says Shadow "has no self-power cooldowns". The healing lists already cast it. | Treated as learnable; still OPEN until someone sees it at the trainer or in a spellbook | S f5a3613e0 (Shadow APL) |

## Priest: Shadow

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | Darkness reached only part of Shadow damage. The engine copied Era's class-masked mods: base damage only on Mind Blast and Devouring Plague, full damage on Shadow Word: Pain and Mind Flay. That left out Shadow Word: Death entirely, and the spell-power part of Mind Blast and Devouring Plague. | **Forever 70707:** SpellEffect for 15259 is one aura-79 MOD_DAMAGE_PERCENT_DONE effect, misc 32 (Shadow). The Wowhead tooltip reads "Increases your Shadow damage done by 10%". **Era 1.15.9:** three aura-108 class-masked effects. Code: `priest/talents.go` applyDarkness. Test: `priest/shadow/darkness_test.go` TestDarknessRaisesAllShadowDamageDone. | CONFIRMED-BUG | E 7b4485455 |
| P2 | Shadow Word: Death added no Shadow Weaving stack, although every other Shadow damage spell does. | Talent text: "Your Shadow damage spells have a 100% chance…". `priest/shadow_word_death.go`. Test: TestShadowWordDeathAddsShadowWeaving. No DPS change: the stack's uptime is already 99.7%. | CONFIRMED-BUG (no DPS) | E 7b4485455 |
| P3 | Shadow Word: Death's backlash was the damage dealt (Era). Forever's is 30% of maximum health. | Wowhead Forever 1309636: "you take backlash damage equal to 30% of your maximum health". Test: TestShadowWordDeathBacklashDamagesTheCaster. No DPS change. | CONFIRMED-BUG (no DPS) | E 7b4485455 |
| P4 | The APL cast Shadow Word: Death only in the execute window. Both Forever guides cast it on cooldown, ahead of Mind Flay. | Warcraft Tavern Shadow Priest, rotation item 4: "Cast Shadow Word: Death when the cooldown is available". wowtbc Shadow Priest, item 6, above Mind Flay. Gain: +2.5% (653.0 to 669.2, both seeds). | CONFIRMED-BUG (APL) | S 0b9fef7a5, ef8825c51, E 6b6e50863 |
| P5 | Shadowfiend was not in the APL (see C2). | Cast below 30% mana: +13.7% on the bare preset (360.2 to 409.6, where Shadow ran dry for 17 s), and 0.0% on the raid preset. A 50% gate cost 0.1% on the raid preset, so the gate is 30%. | Follows C2 | S f5a3613e0, E 5b6e903bf |
| P6 | Shadowform's crit bonus on DoT ticks after 09ca529f3. | Shadowform 15473 effect 4: aura 108 misc 15 (crit damage bonus), whose class mask covers Mind Blast, Mind Flay, Shadow Word: Pain, Devouring Plague and Shadow Word: Death. Ticks resolve through `Dot.outcomeCritTick`, which takes the parent spell's crit multiplier, 1 + 0.5 × 2 = 2.0x. Inner Focus stays off periodic ticks (`innerFocusAddsCrit`). | Aligned | - |
| P7 | Shadow Weaving is self-stacking: a self buff of +2% Shadow damage a stack, five stacks, 15 s. | 70707 text reads "the Shadow damage the target takes from you": a target debuff, but from you only. On one target that equals the engine's self buff. 71 applications a fight, 99.7% uptime. | FOREVER-CHANGE (equivalent single-target) | - |
| P8 | Vampiric Embrace is never cast. | Tooltip (15286): 30 s cooldown, a party heal with no damage. wowtbc: "during party-wide damage". A DPS sim leaves out utility casts. | Aligned (policy) | - |
| P9 | Devouring Plague, Mind Blast and Shadow Word: Pain. | Devouring Plague: 24 s cooldown (SpellCooldowns 55344), 848 over 24 s, ticks crit. Mind Blast: 8 s cooldown minus Improved Mind Blast's 0.5 s a rank (54580). Shadow Word: Pain: plus 3 s a rank of Improved Shadow Word: Pain. Twin Disciplines' client masks (1225132) cover Mind Flay, Devouring Plague, Shadow Word: Pain and Shadow Word: Death, which matches `PriestSpellMaskInstantShadowDamage`. | Aligned | - |
| P10 | Raid talent search, with `-keep`. | Best variant is +0.4% (Wand Specialization 2 to Mental Strength): under 1%. | Hold | - |

## Priest: Discipline and Holy (first audit)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| H1 | Dark Sacrifice was registered for every race. The troll Holy healer and the human Discipline healer drew about 2,000 mana a fight from it. | SkillLineAbility rows for 1277324 to 1277328 carry RaceMasks_0 16 (Undead). Test: TestDarkSacrificeIsUndeadOnly. The healing tests now build an Undead healer for Dark Sacrifice. | CONFIRMED-BUG | E 72bfac1af |
| H2 | Prayer of Mending only spent a charge on a heal. Its other trigger, the holder taking damage, was unmodeled: the code comment said the fake raid had no damage hook, but `Raid.OnRaidDamage` exists (Lightwell uses it). Holy spent about 2 of the 5 charges per cast. | Wowhead Forever 1240827: "heals them … the next time they take damage or receive non-periodic healing". Test: TestPrayerOfMendingSpendsAChargeWhenTheHolderTakesDamage. | CONFIRMED-BUG | E c2c29ccf5 |
| H3 | Talent rates against the 70707 max-rank text. | All match `priest/healing_talents.go`: Mental Strength 10%; Spiritual Guidance 25% of Spirit; Improved Renew 15%; Improved Healing -15%; Divine Fury -0.5 s; Soul Warding; Divine Aegis 15%, capped at 20 × level; Renewed Hope +10% crit and -5 s; Litany of Light 10%; Improved Power Word: Shield 20%; Mental Agility 10%; Spiritual Healing 10%; Inspiration 25%; Meditation 50%; Holy Specialization 5%. | Aligned | - |
| H4 | Twin Disciplines on heals. | Client masks (1225132, effects 0 and 1) match Renew, Penance, Prayer of Mending, Holy Nova and Desperate Prayer. They do not match the player Flash Heal ranks (0x800) or Power Word: Shield. Same as `PriestSpellMaskInstantHealing`. | Aligned | - |
| H5 | Ranks and downranking. | The top ranks Greater Heal 5 (25314), Renew 10 (25315) and Prayer of Healing 5 (25316) are `book: true` in `spellranks.json`. The engine stops one rank short, as round 1 ruled for Frostbolt 11. The lists downrank (Heal 4, Renew 3, Power Word: Shield 3747). Tooltips agree on the trained ranks: Heal 6064 634-713 for 305 mana, Flash Heal 10917 778-914 for 380, Prayer of Healing 10961 581-612 for 1030. | FOREVER-CHANGE (as data says) | - |
| H6 | Penance, Power Word: Shield, Lightwell and Spirit of Redemption. | Penance 1316995: 425 a pulse, three pulses, 12 s cooldown, matching `penance.go`. Power Word: Shield 10901 absorbs 950, 4 s cooldown, 15 s Weakened Soul. Lightwell is baseline at 40/50/60 (SkillLineAbility 724/27870/27871). Spirit of Redemption cannot fire, because the healer never dies. Penance's healing coefficient and Power Word: Shield's 0.1 are the vanilla convention ("unconfirmed" in constants_auto_gen.go). | Aligned; the coefficients are OPEN | - |
| H7 | Mana model. | Priest Spirit regen is 12.5 + Spirit/4 per 2 s tick, plus the five-second rule and Meditation's 17/33/50% (`priest/priest.go:206`, `core/mana.go`). This is the Era formula; no Forever source differs. | Aligned | - |
| H8 | Power Infusion (10060: 20% healing for 15 s, 3 min cooldown, castable on yourself) is in neither build. | Its value depends on raid policy (who gets Power Infusion). No search was run. | OPEN | - |

## Mage: Arcane

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| A1 | Arcane Blast on 70707. | Talent text and tooltip 400574: 1.5 s cast, +12% of Intellect as spell damage a stack, +100% cost a stack, 4 stacks, 18 s from the first application. `mage/arcane_blast.go` has 12%, 100%, 4, 18 s, and only the first cast activates the aura (no refresh). | Aligned | - |
| A2 | Haste weight (round-1 A3). | See C1: +3.5 ± 0.8 per 1%. | Aligned | - |
| A3 | Raid talent search, with `-keep`. | Best variant is +0.9% (Frost Channeling 2 to Incineration): under 1%. | Hold | - |

## Mage: Fire

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| F1 | Ignite's ticks took the target's percentage modifiers again. The tick multiplier was x1.265 with 5 Improved Scorch stacks and Curse of the Elements, on a burn that is 40% of a crit which already took them. The earlier fix removed only the attacker half. | Blizzard's beta development notes: "Ignite no longer double dips on % damage increase modifiers" (us.forums.blizzard.com/t/2360696). `mage/ignite.go`: igniteTick now carries `SpellFlagIgnoreTargetModifiers`. Test: TestIgniteTicksSkipTheTargetsDamageTakenModifiers (failed at 1.265, passes at 1). Effect: -4.6%. Ignite is 22% of Fire's damage. | CONFIRMED-BUG | E 384cc1899 |
| F2 | Ignite tick crits (owner rule). | Ignite is a non-channelled DoT, so it crits, on the magic table at the spell's own 1.5x. Its tick spell has ProcMaskSpellProc and its ticks go to OnPeriodicDamageDealt, so a tick crit never feeds Ignite, Combustion, Heating Up or Master of Elements (TestIgniteTicksCritWithoutFeedingIgnite). | Aligned | - |
| F3 | Fireball and Pyroblast DoT crits, Master of Elements and Combustion. | Both DoTs crit (non-channelled). Master of Elements refunds 10% a rank of base cost on Fire and Frost direct crits, matching 29074's text. Combustion: 3 non-periodic crits, +10% a Fire hit, Ignite excluded (`talents.go` registerCombustionCD), matching 11129. Fire Power counts Fire DoTs and the Improved Scorch debuff, matching 11124. | Aligned | - |
| F4 | Fireball filler after F1. | Fireball measures -10.5% against Scorch (624.7 and 626.1 against 698.3 and 699.9), with 13 s out of mana. Scorch stays. | Aligned | - |
| F5 | Raid talent search, with `-keep`. | Best variant is +0.9% (re-spend of non-damage points): under 1%. | Hold | - |

## Mage: Frost (70707 re-verification)

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| R1 | Ice Barrier. | Tooltip 11426/13033: +15% Frost damage for 1 min, +5% more while shielded, 30 s cooldown. `mage/ice_barrier.go` has 15 + 5, 1 min and 30 s. The absorb is unmodeled, so the +5% holds the whole minute. Nothing hits a mage standing at 30 yd. It is cast 3.48 times a fight. | Aligned | - |
| R2 | Winter's Chill: 20% a stack. | 70707 text (rank 5, 11180): 100% chance, "+20% damage taken from your next Cone of Cold and an additional 7% of your spell damage", stacks to 5. It only feeds Cone of Cold, which is out of reach at 30 yd (round-1 R2). No DPS effect. | Aligned | - |
| R3 | Ice Lance does 300% more against Frozen targets. | Tooltips 1240047/1312002 and `ice_lance_multiplier_test.go`. Fingers of Frost gives 2 charges and Shatter adds 50% crit (71% crit observed). | Aligned | - |
| R4 | Ice Lance cooldown. In the client tables only rank 1312002 has one (CategoryRecoveryTime 8000; ranks 400640 and 1240044-1240047 have no row). Wowhead Forever shows "8 sec cooldown" on every rank. | The owner's DBCache.bin is build 70338 and holds no Ice Lance hotfix. A server-side 8 s on every rank would block the second Fingers of Frost charge. | OPEN (settle with an in-game cast of two Ice Lances inside 8 s, or a 70707 DBCache) | - |
| R5 | Frost Blast (1326053-1326055, levels 24/36/48) is in trainables as `class_spell`, but has no SkillLineAbility row in 70707. | Same shape as round-1 W5 (Haunt, Unstable Affliction). It would also be a weak filler: 80-88 damage, 1.5 s cast, 90 mana. | FOREVER-CHANGE (correctly unregistered) | - |
| R6 | Raid talent search, with `-keep`. | No variant beats the guide beyond error. | Hold | - |

## Real logs

- **Warcraft Logs.** No Forever priest or mage parses exist at any level. The beta was capped at 30, and the WCL search returns only Era, TBC and WotLK pages.
- **Our API** (`GET /v1/reports/recent`, public).
  - The three recent reports are 2026-10-20 to 2026-10-21. Two have 0 fights.
  - z2opt2x2gs5u has 42 trash fights, 0 bosses and an empty zone, the same beta log round 1 found.
  - No level-matched boss fight exists to compare a sim band against.
- **The owner's beta client** has no combat logs under `Logs/`.

## Before/after (relative, lane engine against the pin)

| Spec | Preset | Before (seed 1 / 50001) | After | Change | Cause |
|---|---|---|---|---|---|
| Shadow, horde | raid | 645.1 / 645.3 | 669.2 / 669.2 | +3.7% | P1 +1.2%, P4 +2.5% |
| Shadow, alliance | raid | 644.2 / 644.3 | 668.1 / 668.4 | +3.7% | same |
| Shadow, horde | bare | 360.2 (lane, before P5) | 409.6 / 409.5 | +13.7% | P5 |
| Fire, horde | raid | 732.3 / 734.0 | 698.3 / 699.9 | -4.6% | F1 |
| Arcane, horde | raid | 652.9 / 653.7 | 652.9 / 653.7 | 0 | - |
| Frost, horde | raid | 562.0 / 561.7 | 562.0 / 561.7 | 0 | - |
| Holy, horde (troll), effective HPS | raid | 657.3 / 657.4 | 760.1 / 760.7 | +15.6% | H2 up, H1 down |
| Holy, alliance (gnome), effective HPS | raid | 673.5 / 673.5 | 771.5 / 771.8 | +14.6% | H2 up, H1 down |
| Discipline, horde (undead), effective HPS | raid | 586.2 / 586.3 | 586.2 / 586.3 | 0 | - |
| Discipline, alliance (human), effective HPS | raid | 592.9 / 592.9 | 568.6 / 568.9 | -4.1% | H1 |

## Still open

- **Shadowfiend (C2).** Settle it at the trainer or in a spellbook at launch. If it is not learnable, revert f5a3613e0's line and remove it from both healing lists. Its pet's own damage is unmodeled.
- **Ice Lance's 8 s cooldown on every rank (R4).**
- **Penance and Power Word: Shield coefficients (H6).**
- **Power Infusion on yourself for healers (H8).**
- **The heal profile.** With H2, the Holy raid sets' mana now lasts 316 to 368 s of a 300 s fight. The retune rule in `heal-profile.json` wants mana to be a live constraint for at least one raid set, and Discipline still runs dry at 168 to 183 s, so the rule holds. Holy's numbers moved a lot, though; the integrator may want to re-check the 60-85% coverage band.
- **Hunter's Lacerating Strikes** (sim/hunter, physical lane) uses the same "share of a finished hit" shape as Ignite. If it also takes target modifiers again, it has F1's bug. Not checked here.
- **The ladder goldens.** TestRotationLadder goldens are not regenerated, per the brief. The Shadow rotation moved (Shadow Word: Death on cooldown, Shadowfiend), and the engine moved Fire, Shadow and the healers.

## Sources

- Warcraft Tavern, Shadow Priest Guide (Forever): https://www.warcrafttavern.com/forever/guides/shadow-priest/
- wowtbc.gg, Shadow Priest PvE Guide, WoW Forever: https://wowtbc.gg/warcraftforever/class-guides/shadow-priest/
- wow-forever.gg, Shadowfiend 401977: https://wow-forever.gg/db/spells/401977-shadowfiend/
- Blizzard, WoW Forever Beta Development Notes (Ignite): https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes/2360696
- Wowhead Forever tooltips: https://nether.wowhead.com/forever/tooltip/spell/ with the ids 1309636, 15259, 15257, 19280, 15286, 10947, 401977, 6064, 10917, 10961, 1316995, 1240827, 10901, 10060, 11426, 13033, 400640, 1240044, 1240047, 1312002, 400574, 11129, 29074 and 1326053/1326055

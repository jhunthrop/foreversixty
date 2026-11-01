# Sim-accuracy ledger

The owner's goal (2026-10-31): repeat audit and fix rounds over every spec until a whole round finds zero confirmed sim bugs. Each round has three or more audit groups, and every confirmed bug is fixed end to end before the next round starts.

| Round | Dates | Reports | Confirmed bugs | Engine pin after | Zero-bug round? |
|---|---|---|---|---|---|
| 1 | 2026-10-31 | design/reviews/2026-10-31-theorycraft-audit/ (casters, hybrids, physical) | every row in the reports' Fixed columns; the rest held under the adoption rule or open | 17d31ecfa (via 31fbe10e2, c99fb94b4, f923e1ad1, 09ca529f3 and the lane follows) | no |
| 2 | 2026-10-31 to 2026-11-01 | design/reviews/2026-11-01-sim-audit-round2/ (warlock, casters, hybrids, physical) | about 35 fixed; 2 held (warlock demons refusing player buffs, and the Destruction build that depends on it) | c8d29ea33 | no |
| 3 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round3/ (warlock, physical, hybrids, casters) | 22 fixed (12 from the fresh re-audit, 1 from the wow-player review); round 2's two holds landed | 640da215f | no |
| 4 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round4/ (warlock, physical, hybrids, casters, integration) | 33 fixed (warlock 6, physical 6, hybrids 8, casters 7, integrator 6); 1 from the wow-player review | f8a5f94ce | no |

## Round 4 at a glance

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (6).**
  - Demons take Forever's 17% attack power and 10% spell damage shares, not Classic's 56.5% and 15%. Sources: the Forever UI's `WARLOCK_PET_BONUS`, wowsims/forever's beta readings and the beta log fits.
  - Demon melee takes 8/17 of the warlock's hit and spells all of it. This rests on one Blizzard statement, forever-bugs #394.
  - Lash of Pain takes Malevolence's crit.
  - Shadow Vulnerability raises only the warlock's own Shadow damage.
  - Bane of Doom needs a minute left to be cast.
  - Eureka! follows each Gnome class's own client row: 8 s, no charges, class masks. A new generator, `make eureka`, writes it.
- **Physical (6).**
  - Weapon damage: the item file was Era's `db.json`, and the client curve missed QualityModifier and the caster-weapon bit. 1,027 of 1,036 melee weapons now match Wowhead Forever.
  - Items now carry ranged attack power (stat 39) and "+N Weapon Damage" (stat 83, Molten Heart).
  - Hunter pets inherit hit.
  - Wound Poison is outside Venom and Vile Poisons.
- **Hybrids (8).**
  - Heals over time crit; the client flags every trained one.
  - Touch of the Grave, Orc Savage and Tauren Endurance follow the 70707 client.
  - Maul spends Clearcasting.
  - Bloom's heal rolls no cast procs.
  - Seal of Righteousness takes spell power once. The Retribution list now keeps Seal of Command and twists to Righteousness, gated on mana and attack power.
- **Casters (7).**
  - Arcane Concentration's Clearcasting is one charge.
  - Arcane Blast's stacks and Missile Barrage keep the other cost mods.
  - Arcane Missiles rank 8 is an Ahn'Qiraj book.
  - Fire casts Evocation below 5% mana. It also adopts a raid build with Arcane Meditation, +5.2%.
  - Items keep holy, nature, frost, shadow and arcane spell damage and spell penetration (ItemSparse 84 to 89 and 47).
  - Wowhead-only items keep their flat crit, hit, dodge, parry, block and defense.
- **Integrator (6).**
  - The /sim page defaults to the published numbers' type-neutral target.
  - Its Raid-buffed buffs and debuffs come from the curated raid preset.
  - The coach's sim reference fights the type-neutral target.
  - A rage potion on a class with no rage bar no longer panics the engine.
  - The /sim armor field names the armor of the target the character fights; found by the wow-player review.
  - The BiS ranker kept main-hand-only weapons out of the off hand. The bis.yml run after round 3's weapon-DPS fix had put Teebu's Blazing Longsword in Survival's off hand, where the engine swings nothing: 653.7 against 768.6 on the re-rank. Otherwise that run's picks are sane: Feral takes Hammer of the Titans.

The heal profile is retuned to ×1.65 by its own rule. The Restoration druid covered 91% after heal-over-time crits; raid sets now cover 60.4 to 84.0%.

Before and after, pin 640da215f against f8a5f94ce on today's published data. Raid preset, band 60, 10,000 iterations, alliance / horde, seed 1 (seed 50001 within 0.2):
- Arcane -9.3 / -9.0.
- Fire -12.4 / -12.9 on the published build. The new build adds +5.2.
- Frost -2.2 / -2.5.
- Retribution +7.3 / -10.2. Horde is on Blade of Hanna; on Arcanite Champion it is +10.1.
- Protection paladin threat 0.0 / -4.5.
- Restoration shaman HPS +4.7 / +4.7.
- Restoration druid HPS +2.8 / +2.4.
- Holy paladin HPS +2.0 / +2.0.
- Holy priest HPS -2.1 / -1.4.
- Discipline HPS +0.7 / +0.5.
- Affliction -3.6 / -3.1.
- Demonology -1.4 / -0.7.
- Destruction -1.7 / -1.2.
- Balance 0.0 / -1.0.
- Feral 0.0 / -0.6.
- Feral bear threat -0.9 / -1.1.
- Beast Mastery +0.7 / +0.6.
- Marksmanship +0.5 / +0.7.
- Shadow -0.4 / -0.1.
- Elemental 0.0 / +0.8.
- Enhancement 0.0 / +0.9.
- Arms 0.0 / +0.8.
- Protection warrior threat 0.0 / -1.2.
- Every other spec 0.0.

The item-data fixes reach published numbers only after data.yml and bis.yml regenerate simdb.bin and the BiS sets. integration.md lists the lane measurements.

Open going into round 5:
- **Owner decisions.**
  - /tiers/healer ranks healers side by side. The ×1.65 retune puts Restoration shaman 0.8% above Holy priest, where ×1.5 had the priest 7% ahead. The wow-player review asks that healer scores stay out of a cross-class order.
  - Power Infusion for healers is policy: it goes to a DPS. One self-cast would add at most about 1.7%.
- **The /sim page's consumables** are its own list, not the per-spec curated groups plus the class kit. Combat reads 7.6% low on the page with no poisons, and Survival 4.1% high (integration I6). The page needs per-spec resolution.
- **The BiS ranker's verify losses.** After the item-stat fixes, the ranker's weights pick sets its own verification scores lower: Survival, Arcane (Kindling Stave), Enhancement, Destruction (Elemental Mage Staff) and the tanks (defense items). Rhok'delar is not among Marksmanship's candidates.
- **Blocked on level-60 logs** (launch is 2026-11-04):
  - Demon crit (Blizzard's stated intent, not live; +1.7 to +2.6%) and haste.
  - The demon level-60 sheets.
  - Pet haste for hunters.
  - The rage-crit factor and Sunder's AP threat.
  - Holy Strike's flat bonus and Unsealed Wrath on special attacks.
  - Omen's 4% at 60.
- **Single source, one in-game look each:**
  - Ice Lance's 8 s on ranks 2 to 6.
  - Ice Barrier's +15% / +5%. The shipped tables say 3% / 3%; worth 9% of Frost.
  - Touch of the Grave's cooldown, 25 s or 15 s.
  - Blood Fury, 2 or 3 min.
  - Endurance's 2% crit.
  - The Human Spirit, 5 or 10%.
  - Pet hit from forever-bugs #394.
- **Carried over:**
  - Kings with Heart of the Lion (moot for the preset and the page now).
  - Shadowfiend at launch.
  - Pet family.
  - Demonic Brand unmodeled.
  - Siphon Life crit held for the owner.
  - Healing Stream tick crit.
  - Master of Elements on a Clearcasting cast.
  - Nine weapons on hotfix rows newer than the seed.
  - Expertise (37) and haste (36) item stats unconfirmed.
  - The fork's Settings suite stands melee at 30 yd.
- **Pre-existing:** five data tests (`test_addondata`, `test_audit_addon`) already fail on origin/main.

## Round 3 at a glance

Confirmed and fixed, by group:
- **Warlock (7).**
  - Demons refuse player buffs (round 2's held dfbb63ae3, as E e3bdac097).
  - Demons inherit Classic's shares: 75% Stamina, 30% Intellect, 35% armor, 56.5% of spell power as attack power and 15% as spell damage (E 94a45a77b). Two sources each: wowsims/sod for the same client spell 416189, and warcraft.wiki.gg or TBC. Resistances (40%) and hit (100%) are single-source. Crit is not inherited: the 29 low-level logs show pet crit far below the warlock's.
  - Lash of Pain takes Forever's 50, not Era's 99. Firebolt uses the client's range. Unholy Power is one pet-wide multiplier (E f6d77f02a).
  - Demons at bands 20 and 30 had no stats or weapon. They now have level-interpolated stats and their top legal rank (E d1e2536b9).
  - Destruction adopts Pandemic and Cataclysm for Unholy Power: +1.55% with inheritance (round 2's held build, S e4db33928).
- **Physical (6).**
  - Fury holds Death Wish and Recklessness for execute, unless that costs a use. It gains at every fight length from 30 to 300 s. Round 2's 60 s loss was the fork suite standing melee at 30 yd.
  - talent-search and rotation-search fought a Humanoid, while every other tool fights a target with no type. That was the Marksmanship gap (647.0 against 619.5).
  - Improved Tracking gave nothing on the published target, which has no creature type. Every Phase 1 boss is trackable, so it now counts (Marksmanship +4.2%, Survival +5.0%).
  - Murder joins the rogue raid keeps.
- **Hybrids (5).**
  - Shred was Era's 225% weapon damage; Forever's is 155% (client, Wowhead Forever, three beta log fits).
  - Claw is 110%.
  - Savage Fury follows its client mask.
  - Omen of Clarity is a flat 4% a spell cast, doubled in Moonkin Form, measured against 12,402 logged casts.
  - Judgement of the Crusader is in the raid preset (owner ruling).
- **Casters (3).**
  - No rotation casts Shadowfiend. A Forever priest cannot learn it: about 25 beta priests over 1,103 encounters never cast it, and its skill-line row has no learn level.
  - Every Ice Lance rank waits Fire Blast's 8 s category cooldown (client 70685 change, plus Wowhead Forever's text).
- **Ranker (1, from the wow-player review).** The BiS ranker priced weapon DPS from `items/<class>.json`, while the sim swings `simdb.bin`'s damage. About 200 weapons disagree. Elemental Mage Staff is 57.5 DPS in the item file and 42.7 in the engine, so it won the cat's main hand on damage the sim never deals. Candidates now take the engine's weapon row before scoring (S 425147a6f). In a local cat run, Hammer of the Titans replaces the staff: +4.8%. Every melee, hunter and hybrid weapon pick can move at the next bis.yml run.

Settled without a code change:
- Penance (0.19 a pulse) and Power Word: Shield (0.10): the client tables match the engine.
- Revenge rank 6 is an Ahn'Qiraj book rank, so not learnable in Phase 1.
- Venom with Vile Poisons stays additive.
- Elemental keeps Improved Reincarnation (owner ruling).
- Siphon Life crit is held for the owner, unchanged.

Before and after, pin c8d29ea33 against 640da215f on the same published data. Raid preset, band 60, mean of seeds 1 and 50001, 10000 iterations, alliance / horde:
- Feral -11.3 / -11.3.
- Restoration druid HPS -1.4 / -1.4.
- Marksmanship +4.2 / +4.2.
- Survival +5.0 / +5.0.
- Fury +0.8 / +0.4.
- Frost -0.9 / -0.9.
- Discipline HPS -8.1 / -7.8.
- Holy priest HPS -15.6 / -17.3.
- Retribution +9.3 / +11.3.
- Protection paladin threat +9.8 / +9.3.
- Affliction -7.0 / -7.0.
- Demonology -6.5 / -6.3.
- Destruction -6.4 / -6.3.
- Every other spec 0.0.

The published BiS files still carry the old Destruction build and the old weapon picks, and bis.yml regenerates them; the lane measured Destruction at -4.8% with its new build. The new order is Destruction, then Demonology, then Affliction.

Open going into round 4:
- Pet crit, haste, hit with Suppression, and Lash of Pain's own spell-power share need a level-60 warlock log with its sheet. The demon stats at bands 20 and 30 are interpolated.
- Kings with Heart of the Lion: the client cannot tell. A Forever character sheet with both buffs settles it.
- Ice Lance's 8 s on every rank: one source. Two Ice Lances inside 8 s in game settle it.
- Shadowfiend: if a launch trainer shows it, revert S 09848d599 and E 72aefbb89.
- Which weapon damage is true for Forever, `items/<class>.json` or `simdb.bin`? The tooltips probably show the item file's numbers.
- The /sim page defaults to a Humanoid target, so rogues show 4% above their published number. Its "raid-buffed" list (`web/src/lib/sim/settings.ts`) carries Kings, Sanctity Aura, Winter's Chill and Shadow Weaving and no Heart of the Lion, against the curated preset.
- The fork's Settings suite stands melee at 30 yd, so it cannot judge a melee rotation.
- Single-source rows:
  - Sunder's 5% AP threat.
  - The rage-on-crit factor.
  - Omen's 4%, measured at levels 20 to 30 only.
- Holy Strike's flat bonus under its weapon share.
- Power Infusion for healers.
- The Restoration druid covers 91% of incoming damage, above the heal profile's 85% ceiling.
- Demonic Brand is unmodeled.
- Wound Poison takes Venom and Vile Poisons against its client masks; no rotation casts it, so it moves no number today.
- Unsealed Wrath's reading.
- Wild Growth under a party-damage profile.
- Pet family.

## Round 3 wow-player review

SHIP: Balance, Feral bear, Restoration druid, the three hunters, Arcane, Fire, Holy paladin, Protection paladin, Retribution, Discipline, Holy priest, Shadow, Combat, Subtlety, Enhancement, Restoration shaman, and the three warriors. The reviewer held six, and each was settled before the push:
- **Feral, "586 DPS on a caster staff".** This is the ranker bug above (S 425147a6f). The next bis.yml run puts a real weapon on the band.
- **Assassination, "about 260 DPS on a mace and sword".** That came from a stale lane note. The published set is already on daggers after the nightly re-rank: 601 alliance, 615 horde.
- **Affliction, "the published build is the pre-round-2 build".** Also stale. The published band-60 build is round 2's Malevolence build.
- **Frost, "the 8 s Ice Lance rests on one source".** It is kept, because the client table is the truth, and it is listed under Open above. It costs 0.9% if wrong.
- **Demonology, "label pet scaling as an estimate".** The guide now says so.
- **Destruction, "-6.4% against the lane's -4.8%".** Both are right. -6.4% is measured on the published build, and -4.8% includes the new build, which bis.yml publishes.

The reviewer also flagged the Fury guide, which still said Death Wish goes on cooldown; it now describes the held cooldowns. Survival's 62 Wing Clips a fight are by design: the line fires only under a Windfury Totem, as a cheap special that rolls the totem's extra attacks, and its note records the measurement.

## Round 2 at a glance

Confirmed and fixed, by group:
- **Warlock (8).**
  - Wrack is a 6 s channel, and its +10% reaches only Corruption and Agony.
  - Nightfall and Shadow Mastery follow the client masks.
  - Affliction adopts Malevolence and Ruin.
  - Decimation applies only below 35% health, its cast cut is 40% of 4 s, and it no longer cuts mana.
  - Blood Fury is off the global cooldown.
- **Priest and mage (7, plus Shadowfiend).**
  - Darkness covers all Shadow damage.
  - Shadow Word: Death adds Shadow Weaving, takes its 30% backlash, and is cast on cooldown.
  - Dark Sacrifice is Undead-only.
  - Prayer of Mending spends a charge when its target is hit.
  - Ignite no longer double-dips target modifiers.
  - Shadowfiend is added below 30% mana.
- **Hybrids (5).**
  - The Feral form attack uses the weapon's DPS, for cat and bear.
  - The Ravage opener counts.
  - Fire Nova uses the Forever totem damage.
  - Retribution takes Unsealed Wrath and swaps seals.
  - Two client amounts are corrected: Heart of the Lion is 280 AP at 60, and Judgement of the Crusader is 161.
- **Physical (11).**
  - Mutilate needs daggers in the BiS ranker.
  - A stone sits beside a poison on rogues.
  - Revenge and Sunder scale with attack power.
  - The pre-pull cast lands at 0 and counts.
  - Lacerating Strikes no longer double-dips.
  - Subtlety adopts Quietus; Beast Mastery adopts Careful Aim 5.

Held:
- Demons refusing player buffs (lane E dfbb63ae3). Hunter pets already refuse them, and the client's warlock pet scaling spell matches the hunter one, but demons inherit no owner stats yet, so landing it alone would push pet damage further from the logs. It lands with a pet-inheritance model.
- Destruction's Pandemic and Cataclysm build (lane S 5d013a9c0). It passes the adoption rule only once demons refuse buffs.

Open going into round 3 (each report's "Still open" has the detail):
- Warlock pet inheritance.
- Siphon Life never crit in 772 low-level logged ticks, against the owner's DoT crit rule.
- Kings stacking with Heart of the Lion.
- Judgement of the Crusader in the raid preset.
- Omen of Clarity's rate.
- Fury cooldowns held for the execute phase.
- Venom with Vile Poisons.
- Murder's target type.
- Revenge rank 6.
- The rage-on-crit factor.
- Penance and Power Word: Shield coefficients.
- Ice Lance's cooldown on lower ranks.
- Shadowfiend's learnability at launch.
- talent-search and spec-breakdown give different baselines for Marksmanship.
- Elemental's Improved Reincarnation hold.

Real logs: no level-60 Forever logs exist before launch (2026-11-04). The 29 warlock logs at levels 13 to 30 on foreverlogs.gg agree that DoT ticks crit and that Life Tap scales with Spirit. Our API holds one public beta trash report. Round 3 starts on the first level-60 logs after launch, and warlocks are checked against real logs every round.

## Round 2 wow-player review

SHIP: Demonology, Destruction, Fire, Arcane, Frost, Protection warrior, Combat, Subtlety, the hunters, Feral, Elemental, Enhancement, Retribution and Protection paladin. The reviewer held three, and each was settled before the push:
- **Affliction, "Wrack is never cast".** The client data settles it. Wrack (1316697) is 36 a tick at 0.143 a tick over a 6 s channel, 0.143 coefficient a second, against Shadow Bolt's 0.343. Every Wrack rotation measured below the shipped one, wowtbc's Forever rotation verbatim included (-2.4%). Row A14 of warlock.md has the numbers.
- **Shadow, "Shadow Word: Death on cooldown kills the priest".** Two Forever guides (Warcraft Tavern Forever, wowtbc Forever) rank Shadow Word: Death above Mind Flay. The 30% backlash costs the healers, not the priest's damage. The rotation stays.
- **Assassination, "about 260 DPS on the page".** The ranker now needs daggers for Mutilate. The bis.yml run dispatched after the push re-ranks the set (576 on daggers in the lane run).

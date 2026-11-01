# Sim-accuracy ledger

The owner's goal (2026-10-31): repeat audit and fix rounds over every spec until a whole round finds zero confirmed sim bugs. Each round has three or more audit groups, and every confirmed bug is fixed end to end before the next round starts.

| Round | Dates | Reports | Confirmed bugs | Engine pin after | Zero-bug round? |
|---|---|---|---|---|---|
| 1 | 2026-10-31 | design/reviews/2026-10-31-theorycraft-audit/ (casters, hybrids, physical) | every row in the reports' Fixed columns; the rest held under the adoption rule or open | 17d31ecfa (via 31fbe10e2, c99fb94b4, f923e1ad1, 09ca529f3 and the lane follows) | no |
| 2 | 2026-10-31 to 2026-11-01 | design/reviews/2026-11-01-sim-audit-round2/ (warlock, casters, hybrids, physical) | about 35 fixed; 2 held (warlock demons refusing player buffs, and the Destruction build that depends on it) | c8d29ea33 | no |
| 3 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round3/ (warlock, physical, hybrids, casters) | 22 fixed (12 from the fresh re-audit, 1 from the wow-player review); round 2's two holds landed | 640da215f | no |
| 4 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round4/ (warlock, physical, hybrids, casters, integration) | 33 fixed (warlock 6, physical 6, hybrids 8, casters 7, integrator 6); 1 from the wow-player review | f8a5f94ce | no |
| 5 | 2026-11-01 | design/reviews/2026-11-01-sim-audit-round5/ (warlock, physical, hybrids, casters, ranker, page, integration) | 30 fixed (warlock 1, casters 8, hybrids 4, physical 6, ranker 3, page 5, integrator 1); 2 more from the wow-player review | 8baffe5fc | no |
| 6 | 2026-11-01 to 2026-11-02 | design/reviews/2026-11-02-sim-audit-round6/ (warlock, casters, physical, hybrids, healers, ranker, pipeline, wasm, page, integration) | 71 fixed (warlock 5, casters 11, physical 10, hybrids 17, healers 10, pipeline 5, ranker 4, wasm 3, page 5, integrator 1) | b57765fb9 | no |

## Round 6 at a glance

The owner's rulings this round:
- **/sim below 60 defaults to the preset the band publishes** (bare below 60, raid at 60). The ranker's rule lives once in `sim/request.PublishedPreset`; Raid-buffed stays a choice.
- **Horde race picks within 1% stay as published.** The adoption rule's bar applies to race.
- **Restoration druid's lead stays,** blocked on launch logs.

The owner's priorities:
- **Protection paladin +50% stands.** Righteous Fury alone is +50.7%, and every input has two sources. This round's two threat bugs take 3.4% back off, leaving the paladin at 96 to 98% of a Protection warrior.
- **Torch of Light leaves the raid sets.** The ranker credited wand DPS that raid warlocks never shoot, and wand shots were Physical.
- **The hotfix push is layered into every generator** and re-fetched when it moves.
- **sim.wasm was at 94% of Cloudflare's limit.** It ships gzipped now, under a 20 MiB budget check.
- **The header layout shift is fixed at its two causes,** with the limit unchanged.

The owner's questions, answered:
- **Survival's Improved Concussive Shot did nothing.** It now goes to Deadly Aspects (+4% leveling, +8% new raid build). Efficiency is worth nothing because Resourcefulness already cuts melee mana by 60%.
- **Wild Growth is modelled right but loses to Rejuvenation on mana.** Raid damage now reaches the healer.
- **The bear was in form all along.** Maul and Swipe lacked Classic's 1.75x threat (+29.5% now).

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (5).**
  - Wand shots were Physical.
  - Classic-db items lost single-school spell damage.
  - Forever's dynamic DoTs for warlock, druid and shaman.
  - Balance and Feral registered Moonfire and Wrath twice, casting the talent-less copies.
- **Casters (11).**
  - Ignite rolls into one burn and is registered once.
  - Improved Scorch benefits only its mage, so it is out of the raid preset.
  - Ice Lance's Frozen x4 covers spell power.
  - Arcane Blast lands on cast.
  - Robe of the Archmage's use.
  - Mage and priest DoTs are dynamic.
  - The Mental Agility and Spiritual Healing masks.
  - The addon's stale default builds.
- **Physical (10).**
  - The build search protected Improved Concussive Shot.
  - The Resourcefulness proc mask.
  - Aspect of the Beast.
  - Bestial Discipline's mana half.
  - Claw ranks.
  - Defensive Stance and Defiance multiply.
  - Mana regen reads its rates each tick.
  - Rune of the Guard Captain is 14 attack power, and terrain riders are dropped.
  - Classic-db shield block.
  - Half-up stat rounding.
- **Hybrids (17).**
  - Maul and Swipe threat.
  - Moonkin crit counted twice, Vengeance on Insect Swarm, Nature's Reach.
  - Natural Instinct on Rip.
  - Genesis off Tranquility.
  - Illumination's casting regen.
  - Benediction on Holy Shock.
  - Shield Specialization mana threat, and mana-gain threat multipliers.
  - Judgements on the melee table, Hammer of Wrath's hit, the Benediction mask.
  - Windfury's Elemental Weapons, Stormstrike's charge, Flurry's 15 s.
  - Purification off Healing Stream.
  - Blessing of Light and Salvation removed.
- **Healers (10).**
  - Raid damage reaches the healer.
  - HoTs are dynamic.
  - Discipline casts Greater Heal; Binding Heal is trainable.
  - The Greater Blessing of Light precast is removed.
  - Chain Heal pacing.
  - Mana Tide rank 4.
  - Coach book ranks and priest rank notes.
  - The rotations heal the healer.
- **Pipeline (5).**
  - Generators read only the shipped tables.
  - The hand dump was incomplete.
  - Removed spells stayed as nameless constants.
  - Nothing re-fetched on a new push.
  - The follow hash ignored the push.
- **Ranker (4).**
  - Wand credit by Shoot time.
  - Verify's flat 1% gate.
  - Empty trinket slots.
  - Doubled ranged attack power.
- **Wasm (3).**
  - The deployed module was at 94% of the limit.
  - The local toolchain differed from CI's.
  - The dev replace built the fork checkout silently.
- **Page (5).**
  - /sim defaulted to Raid-buffed below 60.
  - Header layout shift.
  - Stale real-engine e2e specs.
  - From the wow-player review: an off hand under a two-hander drew an empty socket.
  - From the wow-player review: /sim's Damage tab opened collapsed on a nameless "Build" row.
- **Integrator (1).** web.yml had been red on prettier since f95ff8e77, so round 5's pin never deployed. It was pushed alone (bef2cb599), and web has deployed since.

Before and after, pin 8baffe5fc against b57765fb9, raid preset, band 60, 10,000 iterations. Alliance / horde, seed 1; seed 50001 is within 0.2:
- **Bear threat:** +29.5 / +29.7.
- **Discipline:** +11.1 / +11.0 HPS.
- **Frost:** +7.8 / +7.9.
- **Destruction:** -7.8 / -7.8.
- **Restoration shaman:** +7.6 / +7.4 HPS.
- **Holy priest:** +6.1 / +6.3 HPS.
- **Fire:** -5.7 / -6.1.
- **Elemental:** -5.5 / -5.5.
- **Balance:** +5.3 / +5.0.
- **Holy paladin:** +4.5 / +3.7 HPS.
- **Enhancement:** -3.5 / -3.4.
- **Protection paladin threat:** -3.4 / -3.4.
- **Protection warrior threat:** +3.1 / +3.1.
- **Demonology:** -2.9 / -2.8.
- **Retribution:** +2.1 / +1.6.
- **Beast Mastery:** +1.6 / +1.7.
- **Affliction:** -1.6 / -1.6.
- **Feral:** +1.4 / +1.3.
- **Every other spec:** within 0.5.

The item-data fixes and the re-ranks reach published numbers only after data.yml's fetch and bis.yml run on the pin. Those are:
- wand school
- classic-db spell damage
- the doubled ranged attack power
- Guard Captain
- classic-db shield block
- rounding
- Torch out of the raid sets
- trinket slots
- the new Survival and Beast Mastery builds

Open going into round 7:
- **Owner decisions.**
  - Hawk Eye to Deadly Aspects for Survival: +5% bare, at the cost of range utility, the same class of choice as Grim Reach.
  - The gzipped sim.wasm transfer is at 95% of its 4.5 MB budget.
- **After this round's nightly, check:**
  - The first client.yml push job re-fetches 70707. Its fetch-wowhead step may go red: a local run got 397 of 689 items, under the 90% floor.
  - /tiers/healer reorders, with Discipline passing Holy paladin.
  - The warlock and raid caster ranged slots leave Torch of Light.
  - The new Survival and Beast Mastery builds publish.
  - Horde hunters drop about 1.2% from the ranged attack power fix.
- **Blocked on level-60 logs** (launch is 2026-11-04):
  - Ice Lance's coefficient: the client says 0, the engine 0.4286. At 0, Frost is 9.7% lower.
  - Illumination's casting regen, client rows only.
  - Lacerate and Primal Bite's high-threat factor.
  - Holy Shield's flat threat.
  - Elemental Mastery: one charge, or 15 s.
  - Windfury swings: white or not.
  - Holy Power's crit-heal bonus.
  - Arcanite Champion's proc rate.
  - Ignite's logged 0.73 ratio.
  - Missile Barrage on a miss.
  - Arcane Concentration per missile.
  - Master of Elements on Clearcasting.
  - Prayer of Mending's bounce mask.
  - Binding Heal's crit rolls.
  - Wild Growth under burst damage.
  - Revenge, Shield Slam and Thunder Clap threat.
  - Venom and Rupture attack power.
  - Improved Corruption: add or multiply.
  - Serpent Sting's 0.2 spell power, where the client row has none.
  - Righteous Fury's multiplication.
  - Round 5's demon, pet and rage items.
- **Single source or open:**
  - Shadowfiend at launch.
  - The heal profile's coverage sits outside its 60 to 85% rule: druid about 86%, paladin about 57%.
- **Held under the adoption rule:**
  - Marksmanship's search winner, +0.9%.
  - Fire dropping its Improved Scorch upkeep line, +0.4%.
  - Affliction, +0.7%.
- **Pipeline and page follow-ups:**
  - Items still use the useHotfixes CSV without removals.
  - ChrRaceRacialAbility and the other unfetched tables are hand-written.
  - 24 aura-99-only spells still mirror into ranged attack power.
  - Benediction's planner row lacks its healing line.
  - The real-engine e2e specs need a wasm build in CI (build-real has none).
  - sim-island has 12 bytes of headroom.
  - Raid-buffed doesn't add the class kit (Protection paladin's Devotion Aura at 60).
  - Copy: why Torch loses in raid on BiS.
  - 36 action ids read "An unnamed spell" on /sim's Damage tab: simnames pruning, plus engine ids like Ignite 12654 that the client's SpellName lacks.
  - Ignite and Shoot show as Physical: the result carries no school for them. The engine deals Ignite as Fire, and Shoot's school arrives with damage_school at the data run.

## Round 5 at a glance

The owner's rulings this round:
- **/tiers/healer orders healers only by gaps beyond error at every tuning of the heal profile.** The ranker publishes `hps_error` and `hps_at_bounds` (×1.5 and ×1.65, now configuration in heal-profile.json). With no error published, the page shows no ranks.
- **Power Infusion stays out of every default.** No default carried it; tests now hold that.
- **/sim consumables are each spec's published set,** resolved by level, cut to the character's RequiredLevel and loaded lazily.
- **The ranker's guard covers every role,** and the root causes are fixed.
- **Rhok'delar** is not obtainable in Forever's phase 1. Its leaf drops only in Molten Core, and its sinew comes from Onyxia, which opens 9 December.
- **Ice Barrier's 15% / 5% has three sources:** the 31 October hotfix push on wago.tools, the hotfix cache wowforevertalents.com reads, and Wowhead Forever. Frigid Power carries the 15%, Ice Barrier's own effects the 5%, so the engine stands.

Confirmed and fixed, by group (rows and evidence in each report):
- **Warlock (1).** Healing on items gave a hidden third as spell damage, on top of the damage half the client states. The sources that grant both now state both halves. This is the largest cross-class move of the round.
- **Casters (8).**
  - Mages wore no armor: Molten Armor is an SoD rune a Forever mage cannot learn, so they now wear Mage Armor. Fire also adopts Improved Fire Blast 2/2, +3.3%.
  - Racials follow Blizzard's 26-27 October notes and the hotfix push: Orc Command (Savage removed), Berserking 15 s, Touch of the Grave every 15 s, The Human Spirit 10%, Wind Blessed 2%.
  - Holy Binding is modeled.
  - A Demonic Rune costs 800 health.
  - The race text follows the client.
- **Hybrids (4).**
  - Righteous Fury is +33% all threat and +90% Holy threat, multiplied.
  - Holy Shield costs 100/150/195 mana.
  - Champion of the Light is spell damage only.
  - Healing Stream ticks crit, as periodic heals.
- **Physical (6).**
  - Hunter's Mark rank 4 is 14325 at 143 ranged attack power, not SoD's 1213268 at 110.
  - Deadly Aspects procs Quick Strikes from melee.
  - The item hotfix dump was frozen at build 70009: 201 rows wrong and 569 items missing at 70707. data.yml now layers wago.tools' hotfix pushes.
  - Shields carry block value.
  - Client items take Wowhead's on-equip stats where the client lost the link (Benediction among 153).
  - Item stats 36 and 37 are haste and expertise rating.
- **Ranker (3).**
  - `weight_stats` missed stats Forever's talents turn into damage: Intellect for Enhancement, the hunters and Retribution; spirit and mp5 for mana casters.
  - The stat runner-up never reached verify.
  - Only healers had the published-set guard.
- **Page (5).**
  - The /sim consumables were the page's own lists.
  - They were the level-60 set at every level.
  - They included items above the character's level.
  - A rogue's stone was filed with the potions.
  - /tiers/healer ordered healers by any margin and by the tuned profile.
- **wow-player review (2).** A run in flight when another character loaded wrote its result onto the new character, and the after-sim line read an unowned Droptimizer upgrade.
- **Integrator (1).** The addon showed 20 rotation lines bands before the player can train them, from client copies with no SkillLineAbility row (Recklessness at 30 for Fury, among others).

Also fixed, not counted as sim bugs:
- **data.yml had been red on every push since round 4.** The addon data derived from curated rotations that only the nightly regenerated. A push job now owns it.
- **Two data tests pinned generated state.**
- **The bis.yml conflict handler,** which my push job broke for run 41332442067: that run's picks were lost.

Before and after, pin f8a5f94ce against 8baffe5fc on today's published data. Raid preset, band 60, 10,000 iterations, alliance / horde, seed 1 (seed 50001 within 0.3):
- Protection paladin threat +50.0 / +51.1.
- Fire +5.4 / +5.2 on the published build; the new raid build adds +3.3 at the bis run.
- Balance -4.6 / -3.9.
- Demonology -4.2 / -2.6.
- Elemental -3.4 / -4.1.
- Arcane -2.7 / -3.4.
- Affliction -3.0 / -2.3.
- Destruction -3.0 / -2.3.
- Frost -2.6 / -1.6.
- Shadow -2.2 / -1.8.
- Discipline HPS +1.8 / +0.5.
- Enhancement -0.5 / -1.4.
- Marksmanship +0.9 / +1.1.
- Beast Mastery +0.8 / +1.0.
- Restoration druid HPS +1.0 / +1.0.
- Arms 0.0 / -0.8.
- Retribution 0.0 / +0.8.
- Every other spec within 0.5.

The item-data fixes, the ranker fixes and the re-ranks reach published numbers only after data.yml's fetch and bis.yml run on the pin. The integrator dispatches both after the push. The ranker lane's local re-ranks (verify against the published set, alliance / horde):
- Survival +19.6 / +20.1 (Heartseeker in the off hand).
- Retribution horde +11.6 (Arcanite Champion).
- Enhancement +2.7 / +2.3.

Open going into round 6:
- **Owner decisions.**
  - Horde race picks. Orc beats the published Troll by about 1.0% on all three warlocks (Command). Undead beats the published Orc by 0.5 to 0.7% on Arcane and Fire.
  - /sim defaults to Raid-buffed below 60 while the published bands below 60 are bare. At band 40, buffs give +65 to +181%.
  - The Restoration druid's lead over the other healers rests on the heal profile.
- **After this round's nightly, check:**
  - /tiers/healer shows 1, =2, =2, 4, 5 on the published error.
  - Survival's off hand is Heartseeker; re-shoot /sim.
  - The warlock sets drop the healing-hybrid pieces.
  - Retribution horde is off The Needler.
  - Fire publishes its new build.
- **Blocked on level-60 logs** (launch is 2026-11-04):
  - Demon crit, haste and hit.
  - The demon level-60 sheets.
  - Pet haste.
  - The rage-crit factor and Sunder's AP threat.
  - Holy Strike's flat bonus.
  - Omen's 4% at 60.
  - Master of Elements on a Clearcasting cast.
  - Righteous Fury's two threat bonuses, multiplied or added (about 12% of Holy threat).
  - Fire traps scaling with attack power.
  - Rogue Armor Energize.
  - Raptor Strike lifting the dual-wield penalty.
- **Single source or open:**
  - Tauren row 1349372, +2% critical healing, which survives the push while the notes and tooltip say health only (0.2 to 0.35% HPS).
  - Forever DoTs may read spell power each tick (forever-bugs #398).
  - Demonic Brand: its client rows carry no amounts.
  - Shadowfiend at launch.
  - Kings with Heart of the Lion.
  - Pet family.
  - Siphon Life crit, held for the owner.
- **Pipeline follow-ups:**
  - The spell-table generators still read only the shipped tables, not the hotfix push, so the lanes read the push by hand.
  - About 17 items keep a stat gap Wowhead lists.
  - Keen Machete's quality.
  - 24 rotation lines carry placeholder icons until the next icons run.
  - The local sim.wasm build is 40.4 MB, over Cloudflare's 25 MiB, and sim/go.mod points at the owner's fork checkout rather than the pin.
  - header-finished-sim's CLS check is flaky on origin/main (2 failures in 6).
  - The comment-only fork commit e203aca7e merges with the next pin.
- **Held under the adoption rule:**
  - Affliction +0.7%.
  - Arcane +0.9%.
  - Survival's Deadly Aspects build, +8.6% on a legal dual-wield set, which needs Hawk Eye and Improved Concussive Shot freed from the raid keeps.

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

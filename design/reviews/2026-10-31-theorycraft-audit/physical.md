# Theorycraft audit: physical group (2026-10-31)

Warrior (Arms, Fury, Protection), Rogue (Assassination, Combat, Subtlety), and Hunter (Beast Mastery, Marksmanship, Survival). Survival is melee in Forever. Public theorycraft is compared with the site's level-60 raid preset: client 1.60.1.70658 bis bands, `data/curated/apl/*.json`, guide builds and the engine fork HEAD. Each difference is classified as CONFIRMED-BUG (our sim is wrong), FOREVER-CHANGE (the web source is Era-only and the client shows Forever differs) or OPEN. Some impacts were measured with `sim/cmd/spec-breakdown` and are marked as measured. The rest are hand estimates. This review changes no code, data or APLs.

## Top confirmed bugs across the group, by DPS impact

| Rank | Spec(s) | Bug | Est. impact |
|---|---|---|---|
| 1 | BM, MM, SV | Pets receive every raid buff and inherit nothing from the hunter. In Forever a pet inherits 10% of the hunter's AP and 100% of their crit, and gets no raid buffs. | BM -8 to -15%, MM -6 to -10%, SV -5 to -9% (unrun) |
| 2 | Assassination | Mutilate does not require daggers, so the BiS pairs a mace with a sword. | -8.5% (measured) |
| 3 | MM, BM | Multi-Shot and Aimed Shot run separate cooldowns, but the client shares one category cooldown. | MM -7 to -9%, BM -3 to -4% |
| 4 | Arms | Never casts Slam, and the build lacks Improved Slam. | +5 to +8% |
| 5 | Protection | The raid build is the engine's test input, with Bastion 1/5. | +7 to +8% DPS, +5 to +7% threat |
| 6 | Assassination | Seal Fate never procs off Mutilate. | +4 to +8% for a Seal Fate build |
| 7 | MM | Careful Aim adds melee AP only. | +4 to +6% |
| 8 | Arms | No Recklessness in the execute phase, because the stance gate blocks it. | +2 to +4% |
| 9 | SV | The rotation casts no aspect. | +2.5% (measured, Aspect of the Falcon) |
| 10 | All rogues | No Thistle Tea in the rogue preset. | +1.9 to +2.7% |
| 11 | Rogues | Deadly Poison ticks for 27, not Forever's 18. Mutilate's flat bonus is 67, not 50. | -1.8 to -2.7%, -2.1% |
| 12 | SV | Mortal Shots applies to melee strikes. Trueshot Aura adds melee AP, also for every melee spec. Rapid Fire grants no melee haste. | -2.4% (measured), -1.5%, +1.5% |

Read the group's tier order with this in mind. Assassination's lead over Combat comes from rank 2. Fixed, it lands about level with Combat. The hunter specs, BM above all, sit too high until rank 1 lands. Arms at about 91% of Fury is plausible: public sims put it near 95% with Slam in.


## Warrior (Arms, Fury, Protection)

Audited 2026-10-31 against origin/main in `/Users/jh/code/forever/.claude/worktrees/tc-physical` (4e5dde099) and the engine fork `/Users/jh/code/wowsims-forever` (HEAD 1f542def3). Client truth read from `data/builds/1.60.1.70707` (spellconst, trainables, talents, spellranks) and the 70658 bis files; Wowhead Forever tooltips fetched live for the ids cited. No level-60 Forever log exists yet (the beta closed at the level-30 cap on 23 October), so every public level-60 claim below is theorycraft or another sim, not a measurement.

Cross-spec findings that touch all three warriors are listed once under Fury (rage, racials, Tactical Mastery, book ranks) and referenced from the other sections.

### Warrior — Fury

(a) Public TC (Icy Veins, leprestore, wowfdb, sixtysim, wowtbc.gg) agrees with our core loop: Bloodthirst on cooldown, Whirlwind while it is down, Execute below 20%, Heroic Strike as the rage dump, Death Wish and Recklessness as the burst, hit > crit > strength in weights. Our Fury models the 1 October rebuild correctly (Bloodthirst 45% AP + 48, Furious Precision 4/7/10% off-hand hit, Dual Wield Spec +25% off-hand damage and +10-50% off-hand rage, Flurry behind Death Wish at 5% a point, Lingering Rage and Gore Drinker as no-DPS talents, Whirlwind hitting with both weapons baseline). Three things are wrong: the engine still gives a queued Heroic Strike the old off-hand hit bonus that Blizzard removed on 1 October, the Human racial is the pre-rework 2% sword crit instead of Ruthless 1% crit on everything, and the build spends two points on Enrage (dead in a raid DPS fight) while skipping Raging Blows. Tier: we have Fury tied first of 20 (775); sixtysim's independent WoWSims-Forever run has it 5th of 25 on a 3-minute single boss and 2nd on 1-minute fights, so "top tier" agrees and the exact spot depends on fight length.

(b)

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| F1 | Heroic Strike queue and off-hand hit | Since 1 Oct a queued Heroic Strike no longer lowers off-hand miss chance ("the HS queue trick is gone"). | The HS/Cleave queue aura sets `DisableDWMissPenalty = true` for as long as HS is queued, so the off-hand loses the 19% dual-wield penalty in every queue window. | `sim/warrior/heroic_strike_cleave.go:196` and `:201`; consumed at `sim/core/spell_outcome.go:701`; Fury APL queues HS at >= 40 rage (`data/curated/apl/warrior-fury.json`, last line) | wowfdb changes log (1 Oct entry); leprestore Fury guide "The Heroic Strike Queue Trick Is Gone"; Blizzard EU forum thread on its removal (9 Oct) | CONFIRMED-BUG | Fury -1 to -1.5% when fixed. Derivation: off-hand white miss about 12% outside the queue vs about 0% inside; assume HS queued on ~40% of main-hand swings, off-hand white about 15% of damage, plus the rage those extra hits pay into more HS. Also lowers Fury hit weight slightly. |
| F2 | Human racial | Racial rework: Human Ruthless is +2% crit on every attack (15 Oct), cut to +1% (26 Oct), weapon-independent. | `racialsByRace[Human]` applies Sword Specialization: +2% crit with a sword only. The raid preset is Human with Ironwood Blade (Fury) and Arcanite Champion, a two-hand sword (Arms), so it grants +2% where the client grants +1%. | Client spell 20597 is named "Ruthless" in `data/builds/1.60.1.70707/spells.json`; tooltip `nether.wowhead.com/forever/tooltip/spell/20597`: "Increases your critical strike chance with all spells and attacks by 1%"; engine `sim/core/racials.go:210-216` | Wowhead Forever tooltip (client text); wowfdb (15 and 26 Oct entries); sixtysim race tables ("Ruthless's 1% crit lasts the whole fight", "No racial cares what you swing") | CONFIRMED-BUG | Fury -1.0%, Arms -0.85% (1% crit is 36.3 / 34.6 AP of weight at 0.211 / 0.172 DPS per AP). It affects every Human sword user of every class, and Dwarf has the same problem: Mace Spec in the engine vs Artisan's Precision +1% hit in the client (spell 1259719 tooltip). |
| F3 | Talent build: dead points | Public 60 builds take Raging Blows (Whirlwind and Cleave cost 3 less); Icy Veins also recommends Improved Berserker Rage 2 and Improved Bloodrage 2 for rage. Nobody spends DPS points on Enrage for raids. | Band talents `34320003002000000-05153105022011501-2...`: Enrage 2, Lingering Rage 1, Piercing Howl 1, Improved Intercept 1 (all zero DPS in our raid fight: Enrage needs damage taken, and the DPS warrior takes none), Raging Blows 0, Improved Berserker Rage 0. | Raging Blows is modelled (`sim/warrior/talents.go:294-300`); Enrage needs `OnSpellHitTaken` (`talents.go:617-631`); live tree `data/builds/1.60.1.70707/talents/warrior.json` (Fury index 8 Raging Blows, tier 3, no prerequisite). A valid swap: Enrage 2 + Improved Intercept 1 out, Raging Blows 1 + Improved Berserker Rage 2 in; tier gates still hold (21 points in tiers 0-3, 25 in 0-4, 30 in 0-5). | leprestore (Raging Blows 1/1 and IBR 2/2 in its 60 route); Icy Veins Fury guide (RB 1, IBR 2, Improved Bloodrage 2 named as the rage talents) | CONFIRMED-BUG (Raging Blows); OPEN (Improved Berserker Rage) | Raging Blows: about +0.3-0.5% (3 rage x ~18 Whirlwinds = ~54 rage per 180 s, at ~8 damage per rage). IBR: +10 rage per 30 s but needs an APL line, because Berserker Rage is registered as a Survival cooldown, which `autocastOtherCooldowns` skips (`sim/warrior/berserker_rage.go:71-74`, `sim/core/major_cooldown.go:154`). Settle with a talent and rotation search run. |
| F4 | Cooldown timing | One Death Wish (3 min) and one Recklessness (30 min) per 3-minute fight, ideally stacked with the potion on the execute phase or the last 30 s (Execute is 17-23% of Fury damage). | Death Wish "on cooldown" fires at the pull; Recklessness goes out through `autocastOtherCooldowns` at the pull. Neither overlaps Execute. | `data/curated/apl/warrior-fury.json` (Death Wish line), `sim/warrior/recklessness.go:54-57` | sixtysim Fury and Arms pages ("Death Wish, Recklessness and the potion stacked on its last 30 seconds"); wowtbc.gg Forever Fury ("Ideally used during the execute phase") | OPEN | Probably +1-2% from 20% / crit on the highest-DPS window. Settle with a rotation search that gates Death Wish, Recklessness and the potion on `isExecutePhase` or remaining time <= 30 s. |
| F5 | Normalized rage: crit multiplier | Crit white swings pay 100% more rage (2 Oct, revised from 75% on 1 Oct). sixtysim states 1.77x. | `normalizedRageCritMultiplier = 2.0`. | `sim/core/rage.go:85`; Blizzard notes quoted at `rage.go:64-67` | wowfdb (1 and 2 Oct entries); Icy Veins revision (per memory note); sixtysim disagrees (1.77) | Aligned (OPEN on sixtysim's figure) | None now. The owner's combat log on live servers settles it. |
| F6 | Normalized rage: off-hand factor | No public number. Dual Wield Spec's off-hand rage clause is itself disputed (client says +10-50%; Blizzard's note read like a removal). | Off-hand base is 0.5x the one-hand rate, and Dual Wield Spec 5 multiplies off-hand rage by 1.5. | `sim/core/rage.go:84`, `:99-101`; `sim/warrior/talents.go:525-527`, `:566` | wowfdb "Still open" list; none settle the factor | OPEN | Large if wrong: off-hand rage is about a quarter of Fury's white rage. Settle by measuring off-hand rage per swing with and without Dual Wield Spec on a live dummy. |
| F7 | Normalized rage: per-speed constants | One-hand weapon speed x 3.46, two-hand speed x 4.50 ("a two-hander gives 4.5 rage per second swinging"). | 3.46 / 4.50. | `sim/core/rage.go:82-83` | Icy Veins measurement (25 Sep, via the rage.go comment); sixtysim Arms page | Aligned | None. |
| F8 | Bloodthirst | 45% of AP + 48 at rank 4 (2 Oct). | 0.45 x AP + 48. | Client 23894 effects (2: 48, 3: 45); tooltip "45% of your Attack Power plus 48"; `sim/warrior/bloodthirst.go:30`, `:96` | wowfdb; leprestore; Wowhead tooltip | Aligned | None. |
| F9 | Flurry, Death Wish, Dual Wield Spec, Furious Precision, Lingering Rage, Gore Drinker | Flurry 25% at 5/5 behind Death Wish; DW +20% physical / +5% taken; DWS +25% off-hand damage; FP 10% off-hand hit; LR and GD have no DPS effect. | As described. | `talents.go:180`, `:635-703`, `:729-780`, `:543-583`, `:363-375` | wowfdb; leprestore; wowclassicforever.info | Aligned | None. |
| F10 | Recklessness numbers | Client: 100% crit, 15 s, 30 min, +20% damage taken. | Code matches the client; the doc comment above it still describes the SoD version (50%, 12 s, 5 min). | Client 1719 (aura 290 = 100, dur 15000, cd 1800000); tooltip text; `sim/warrior/recklessness.go:10` (stale comment only) | Wowhead tooltip; stillsixty Arms ("100% extra crit chance for 15 seconds on a 30-minute cooldown") | Aligned | None (stale comment only). |
| F11 | Book ranks at 60 (HS r9 157, Battle Shout r7 139/140 AP, Revenge r6) | Split: foreverdb and stillsixty list Battle Shout rank 7 and Overpower rank 4 "at 60"; sixtysim says rank 6 (115 AP) is "the highest a trainer teaches". | `core.IncludeAQ = false`: HS rank 8, Battle Shout rank 6 (115.8 AP at 60), Revenge rank 5. | `sim/core/config.go:12`; `data/builds/1.60.1.70707/spellranks.json` marks 25286/25288/25289 `"book": true`, but `trainables/warrior.json` lists them at level 60 with source skill_line_ability | foreverdb; stillsixty; sixtysim | OPEN | If trainable: HS +19 per cast, Battle Shout +24 AP for the whole raid (~+1% for every physical spec), Revenge r6 for Prot. Settle with the launch trainer list or wago SkillLineAbility AcquireMethod for 25286. |
| F12 | Weights | Hit > crit > attack speed per 1% (sixtysim); expertise and strength top the per-point list. | Hit 1.87 > crit 1.18 > str 1.00 > agi 0.91 > AP 0.45; haste per 1% 10.15. | `data/builds/1.60.1.70658/bis/warrior-fury.json` band 60 raid | sixtysim; Icy Veins (27% dual-wield white miss before hit, our cap reads 28) | Aligned | F1 and F2 shift hit and crit slightly once fixed. |
| F13 | Fury stance dance | sixtysim's tuned Fury (18/31/2) dances into Battle Stance for Overpower and Rend, keeping up to 25 rage. | No stance dance. Also blocked by the missing baseline Tactical Mastery (A4). | `sim/warrior/stances.go:60` | sixtysim; wowtbc.gg Forever Fury priority list | OPEN | Unknown, probably small. Re-test after A4. |

(c) Verdict: 2 confirmed bugs (F1 HS queue off-hand hit, F2 Human/Dwarf racials) plus 1 confirmed build gap (F3 Raging Blows); 5 open items (F4, F6, F11, F13, and the IBR half of F3).

### Warrior — Arms

(a) Every Forever-specific Arms source (brokenmeta, stillsixty, foreverdb, wowcalc, sixtysim, mmonster) puts Slam in the level-60 rotation with Improved Slam 2/2: 1.0 s cast and GCD, no swing reset, 15 s cooldown, used between Mortal Strikes. Our Arms takes no Improved Slam (Deflection 2 sits where it would go) and its APL never casts Slam. The Bloodthrill/Rend/Overpower loop, Mortal Strike, Weaponmaster and Execute match public TC, and Sweeping Strikes and Whirlwind are correctly left out on a single target. Public sources also have Arms swap to Berserker Stance for Recklessness at the execute phase; our APL dropped Recklessness because the stance dance could not be expressed, and the engine is missing Forever's baseline Tactical Mastery (keep 10 rage). Arms at 91% of Fury is plausible: sixtysim has 94.6% on a 3-minute boss (870.8 / 920.2), with Slam in, and adding Slam should close most of the gap.

(b)

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| A1 | Slam in the rotation, Improved Slam in the build | Mortal Strike > Overpower > Slam (when ready, after a swing) > Rend upkeep > HS; Improved Slam 2/2 in every raid build (36-38 Arms points). | Band talents `02305213032515001-...`: Improved Slam 0, Deflection 2 (zero DPS). The APL has no Slam line. | Client Slam 11605: 15 rage, 1.5 s cast, 18 s category cooldown, weapon + 87; Improved Slam 12862 tooltip; engine Slam is modelled and keeps the swing with any Improved Slam point (`sim/warrior/slam.go:71-90`); APL `data/curated/apl/warrior-arms.json` | brokenmeta Arms guide and ranking page; stillsixty Arms; foreverdb Arms 60 builds; wowcalc Arms; sixtysim Arms ("Mortal Strike, Overpower and Slam") | CONFIRMED-BUG | About +5-8%. Derivation: Arcanite Champion average ~162 + AP 1700/14 x 3.5 ~ 425, + 87 = ~674 per hit, x ~1.36 for 30% crit at 2.2x = ~916 per Slam every 15 s = ~61 DPS (8.7% of 704), less the rage it takes from Overpower and Hamstring. |
| A2 | Improved Slam cooldown cut | The tooltip says "reduced by 3.0 sec" at rank 2; public guides say 18 s goes to 15 s. wowfdb asks whether rank 1 is already 3 s. | 3 s a point: 18 - 6 = 12 s at rank 2, based on Blizzard's 1 Oct note. | `sim/warrior/slam.go:27-30`; tooltip `nether.wowhead.com/forever/tooltip/spell/12862`; talents.json Arms index 14 rank-2 text | Wowhead tooltip; brokenmeta ("15 s cooldown with Improved Slam"); stillsixty ("from 18 to 15 seconds"); foreverdb; wowfdb open question | CONFIRMED-BUG (latent: only matters once A1 lands) | After A1, a 12 s Slam over-credits by 25% more casts: about +1.5-2% Arms DPS too high. Use 15 s at rank 2 (OPEN: 1.5 s or 3 s at rank 1). |
| A3 | Recklessness / execute stance swap | At 20%: Executes in Battle Stance, then swap to Berserker Stance keeping 25 rage, press Recklessness and keep Executing (+3% crit in Berserker). | No Recklessness. The APL note says it was removed because the cast silently fails in Battle Stance. | `data/curated/apl/warrior-arms.json` notes; `sim/warrior/recklessness.go:37` (Berserker-only, correct) | stillsixty Arms (cooldown section); wowforeverbuilds Arms ("Recklessness for burn phases"); sixtysim Arms (Recklessness in its burst stack) | CONFIRMED-BUG (APL) | About +2-4% (15 s at 100% crit on Execute-heavy damage, about +50% damage for 15 of 180 s, minus the swap cost). Needs A4 to be right. |
| A4 | Baseline Tactical Mastery | Tactical Mastery is a baseline trainer spell at 14 that keeps 10 rage on a stance swap; Improved Tactical Mastery adds 15, for 25. | Retained rage is only 3 x Improved Tactical Mastery points (0 without the talent, 15 at 5/5). | Client 1310185 "Tactical Mastery" (effect amount 10, trainable at 14 in `trainables/warrior.json`); 12295 ITM amount 15; `sim/warrior/stances.go:60` | Icy Veins Fury ("retain up to 10 Rage"); stillsixty Arms and Protection ("Baseline (10 Rage kept)", "raises it to 25"); sixtysim Fury ("Tactical Mastery is now baseline ... up to 25 rage") | CONFIRMED-BUG (latent for today's APLs, none of which swap stance) | 0% today; it gates A3 and F13, and every stance swap would lose 10 rage. |
| A5 | Rotation order otherwise | Rend upkeep for Bloodthrill, MS on cooldown, Overpower when lit, Execute under 20%, HS only with spare rage. | Matches (HS dump removed by rotation search: +1.4%). | `warrior-arms.json`; Bloodthrill `sim/warrior/bloodthrill.go` | brokenmeta; wowcalc; mmonster; Warcraft Tavern (Era) | Aligned | None. |
| A6 | Sweeping Strikes / Whirlwind on a single target | Not used on one boss (Sweeping Strikes costs 30 rage in Forever and copies 5 attacks); Whirlwind only on packs. | Not used. Sweeping Strikes 1 is only the Mortal Strike prerequisite. | Client 12292 (cost 300, 30 s cd); talents.json prerequisite Mortal Strike <- Sweeping Strikes 1 | stillsixty; sixtysim (multi-boss builds only) | Aligned | None. |
| A7 | Victory Rush | Trainable at 20; usable only within 20 s of killing a non-trivial enemy. | Not registered. | Client 402927; tooltip "(1 + AP x 15/100) ... Only useable within 20 sec after you kill"; trainables | stillsixty; wowtbc.gg | Aligned (never usable on a single boss) | None for raids. |
| A8 | Spearing Strike | 40% weapon damage, +80% vs Giants, Dragonkin and mounted targets; 15 rage, 20 s; Battle Stance. | Not in the build or APL. | Client 1310222; `sim/warrior/spearing_strike.go` | sixtysim (counts it because its benchmark boss is Dragonkin); mmonster | OPEN | Depends on boss type: Onyxia (9 Dec) is Dragonkin. Worth a boss-type switch on the raid preset. |
| A9 | Build otherwise | stillsixty: 36 Arms (MS, Weaponmaster, Improved Slam, bleeds, Overpower package) + 15 Fury (Cruelty, Unbridled Wrath, Booming Voice). sixtysim: 36/13/2. | 33/16/2: the same core with Booming Voice 5 as the Fury filler, Improved Execute 1, Improved Bloodrage 2. | band talents | stillsixty; sixtysim; foreverdb | Aligned apart from A1 | Fold into the A1 talent search. |
| A10 | Weights | Per 1%: crit > attack speed; strength > agility per point; hit capped at 8% for a two-hander. | Hit 1.93 > crit 1.12 > str 1.00 > agi 0.85 > AP 0.45; haste 7.07 per 1%; specials cap 9. | `bis/warrior-arms.json` band 60 raid | sixtysim | Aligned | F2 lowers crit slightly once fixed. |
| A11 | Rage, racials, book ranks | See F2, F5-F7, F11. | | | | | Arms -0.85% from F2. |

(c) Verdict: 4 confirmed bugs (A1 Slam and Improved Slam missing, A2 Improved Slam cooldown, A3 Recklessness, A4 baseline Tactical Mastery) plus F2; 1 open (A8). A1 is the largest DPS gap in the class.

### Warrior — Protection

(a) Public TC (stillsixty 6/0/45, brokenmeta 11/0/40, sixtysim 0/20/31) has Bastion at 5/5 in every raid tank build ("+10% damage with a shield, which means +10% threat"), along with Improved Bloodrage, Focused Rage, Defiance, Improved Revenge and Shield Slam. Our simmed Protection band is not a searched build: it is byte-for-byte the engine's regression input `ForeverProtectionTalents`, which its own comment says "is not advice". It spends 6 points on Improved Tactical Mastery and Anger Management and only 1 on Bastion. The rotation matches public TC (Shield Block up, Revenge, Shield Slam, Thunder Clap, Sunder filler, Heroic Strike dump). Sunder Armor is missing the attack-power threat term Blizzard fixed on 24 September. Tier among tanks agrees: we rank Prot warrior first of three on both threat (734 TPS vs paladin 627, bear 675) and survival (TMI 39 vs 46 / 53), as do sixtysim (#1 of 3 threat tanks) and brokenmeta (#1 of 3).

(b)

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| P1 | Talent build: Bastion | Bastion 5/5 in every raid build. Arms points go to Deflection and Improved Heroic Strike, or to Cruelty and Unbridled Wrath (sixtysim's threat build). | `35305013000000000-00000000000000000-050533120330001311`: Bastion 1 (2%), Improved Tactical Mastery 5 + Anger Management 1 (6 points for 1 rage every 3 s), Improved Bloodrage 0. This is exactly the engine's regression constant. | Band in `data/builds/1.60.1.70658/bis/warrior-protection.json` equals `ForeverProtectionTalents`, `sim/warrior/talents.go:61-83` ("It is not advice, it is a fixed input"); Bastion modelled at `talents.go:325-327`; Anger Management requires ITM 5 in the live tree | stillsixty Protection (Bastion 5/5, Improved Bloodrage 2/2); brokenmeta Protection (Bastion 5/5 listed among the "talents that weigh the most"); sixtysim Protection (0/20/31) | CONFIRMED-BUG (build) | Moving ITM 5 + AM 1 into Bastion +4 and Improved Bloodrage 2: about +7-8% DPS (287 to ~310) and +5-7% TPS (Bastion multiplies damage threat, while Sunder, Shield Slam and HS flat threat do not scale). The lost AM rage is close to free: the APL notes say a Forever tank is cooldown-limited, not rage-limited. Survival unchanged. Run the talent search for the tank role. |
| P2 | Sunder Armor threat | Sunder threat scales with attack power since the 24 Sep fix; sixtysim models it as 206 + 5% of AP. | Flat 206 only (x stance and Defiance). | `sim/warrior/sunder_armor.go:13`, `:60`; client 11597 effect 63 = 206 with ap_coefficient 0 (the fix is a server hotfix, not in the CDN tables) | wowfdb (24 Sep: "Sunder Armor threat scaling with attack power was fixed on every rank"); sixtysim tank method ("Sunder Armor adds 5% of your attack power to its 206"). The 5% figure is single-source. | CONFIRMED-BUG (coefficient OPEN) | About +2% TPS (AP ~1200 x 5% = +60 per Sunder, x1.45 stance and Defiance, on ~30 filler Sunders per 180 s, out of 734 TPS). No DPS effect. |
| P3 | Revenge | 70707 made Revenge scale with AP: rank 6 tooltip reads (137.7 + 0.25 x AP) to (168.3 + 0.25 x AP), "unannounced". | Rank 5 (book rank 6 gated by IncludeAQ) with flat client damage, no AP term. | Tooltip `nether.wowhead.com/forever/tooltip/spell/25288`; `sim/warrior/revenge.go:22-24`, `:88`; spellconst 70707 still shows 25288 effect 2 = 153 with ap_coefficient 0 | wowfdb (30 Oct 23:20 UTC entry); Wowhead tooltip | OPEN (needs the 70707 regen, and the rank-5 tooltip checked for the same AP term) | If rank 5 also gains 0.25 AP: about +300 per Revenge at 1200 AP, roughly +5-8% TPS. Settle by fetching tooltip 11601 and regenerating spellconst from 70707. |
| P4 | Rotation | Shield Slam on cooldown, Revenge when lit, Shield Block (2 charges, 7 s) kept up, Thunder Clap upkeep, Heroic Strike dump (33% of sixtysim's damage). | Shield Block when down > Revenge > Shield Slam > Thunder Clap refresh > Sunder filler > HS at >= 40 rage. Measured on the tank harness. | `data/curated/apl/warrior-protection.json`; client 2565 (10 rage, 5 s cd, 7 s, 75%) | sixtysim; brokenmeta; stillsixty | Aligned (Revenge before Shield Slam is harmless: separate cooldowns, and Revenge costs 2 rage) | None. |
| P5 | Threat constants | HS +142, Revenge 2.25x damage, Shield Slam flat 508 (SoD's value, unmeasured), Defensive Stance 1.3, Defiance +15%. | HS 145, Revenge 2.25x, Shield Slam 508, stance 1.3 + 0.05 per Defiance rank. | `heroic_strike_cleave.go:77`, `revenge.go:19`, `shield_slam.go:22`, `stances.go:126-130`, `talents_protection.go:35` | sixtysim tank method | Aligned (all equally unmeasured) | None. |
| P6 | Tank weights | sixtysim (pure threat score): defense and dodge weigh below zero ("-0.40 strength a point of defense"), because an avoided swing pays no damage rage. | Composite tank score: expertise 106 > defense 17.9 > dodge 12.6 > agi 10.4 > parry 10.2 > stamina 9.48 > block value 8.31 > hit 1.94 > crit 1.09 > str 1.00 > armor 0.63 > block 0.39 > AP 0.22. | `bis/warrior-protection.json` band 60 raid (`score_unit: tank_score`) | sixtysim | OPEN (a scoring choice, not a mechanic: survival-weighted vs threat-only) | None on mechanics. Worth stating the objective on the weights page. |
| P7 | Dwarf racial | Dwarf: Artisan's Precision +1% hit replaced Earth Fury on 26 Oct (Mace Spec was replaced on 15 Oct). | Mace Specialization crit. | Tooltip 1259719; `sim/core/racials.go` Dwarf block | wowfdb; Wowhead tooltip | CONFIRMED-BUG (same as F2) | Small for a tank (+1% hit vs +crit with a mace). |
| P8 | Baseline Tactical Mastery | See A4. | | | | | None in today's APL. |

(c) Verdict: 2 confirmed bugs (P1 build, P2 Sunder AP threat) plus the shared racial (P7); 2 open (P3 Revenge AP scaling, P6 weight objective).

### Ranked confirmed bugs (by DPS/TPS impact)

1. A1 Arms: Slam and Improved Slam missing from APL and build, about +5-8% Arms DPS.
2. P1 Protection: band is the regression input; Bastion 1 instead of 5, about +7-8% Prot DPS and +5-7% TPS.
3. A3 Arms: no Recklessness on the execute stance swap, about +2-4% (needs A4).
4. P2 Protection: Sunder Armor missing its AP threat term, about +2% TPS.
5. F1 Fury: queued Heroic Strike still removes the dual-wield miss penalty, about -1 to -1.5% Fury.
6. F2 (and P7): Human racial is 2% sword crit instead of Ruthless 1% all-weapon crit; Dwarf has Mace Spec instead of Artisan's Precision. About -1% Fury and -0.85% Arms for Humans; it affects every class.
7. A2 Arms: Improved Slam takes 6 s off the cooldown instead of 3 s; latent until A1, then about -1.5-2%.
8. F3 Fury: Enrage 2 and Improved Intercept are dead points while Raging Blows is skipped, about +0.3-0.5%.
9. A4 all: baseline Tactical Mastery (10 rage kept) missing; 0% today, gates A3 and F13.

### Sources

- https://wowfdb.com/classes/warrior/changes/ (checked 2026-10-30, fetched 2026-10-31): dated change log. HS queue off-hand hit removed (1 Oct), crit rage 75% to 100% (1-2 Oct), Fury rebuild details, Sunder AP threat fix (24 Sep), Ruthless 2% to 1% and Artisan's Precision (15 and 26 Oct), Revenge AP scaling in 70707 (30 Oct), open questions on off-hand rage, Dual Wield Spec rage and Improved Slam rank 1.
- https://nether.wowhead.com/forever/tooltip/spell/{20597,1259719,25286,25288,25289,12862,402927,1719,20572,11597,23925,2565,1680,23894,12292,20662} (fetched 2026-10-31): client tooltip text for Ruthless, Artisan's Precision, the book ranks, Improved Slam, Victory Rush, Recklessness, Blood Fury, Sunder, Shield Slam, Shield Block, Whirlwind, Bloodthirst, Sweeping Strikes, Execute.
- https://www.icy-veins.com/wow-forever/fury-warrior-melee-dps-pve-guide (updated 2026-10-27): level-30 Fury build, 19% dual-wield penalty / 27% white miss, Tactical Mastery keeps 10 rage, Raging Blows / IBR / Improved Bloodrage as the rage talents, HS swing makes no rage.
- https://leprestore.com/guides/world-of-warcraft-forever/wow-forever-fury-warrior-guide-best-builds-race-professions/ (updated 2026-10-21): 15/36/0 Fury PvE plan, "The Heroic Strike Queue Trick Is Gone", Raging Blows and IBR in the 60 route, Bloodthirst 45% AP.
- https://eu.forums.blizzard.com/en/wow/t/one-of-warriors-most-satisfying-mechanics-is-disappearing-in-forever/634526 (2026-10-09): players confirm HS queueing no longer helps off-hand hit.
- https://sixtysim.com/guides/warrior/arms/, /fury/, /protection/ (engine eb48c35, updated 2026-10-28): independent WoWSims-Forever sim. Ranks (Fury #2/#5, Arms #11/#10 of 25, Prot #1 of 3 tanks), Arms 94.6% of Fury at 3 min, Slam in Arms, cooldowns stacked late, Ruthless 1% weapon-independent, Battle Shout rank 6 = 115, crit rage 1.77x, Sunder 206 + 5% AP, Shield Slam 508.
- https://brokenmeta.gg/en/wow-forever/guides/warrior/arms/ and /classement/warrior-arms/ (no date; data build 70425 or later): Arms rotation MS > Slam (15 s with Improved Slam) > Overpower > Rend; Slam's 18 s client cooldown.
- https://brokenmeta.gg/en/wow-forever/guides/warrior/protection/ (updated 2026-10-28): 11/0/40 Prot with Bastion 5/5 and Improved Bloodrage; Prot #1 of 3 tanks.
- https://www.stillsixty.com/classes/warrior/arms and /protection (no date; post 1 Oct trees): Arms raid build with Improved Slam, rotation with Slam, Recklessness via Berserker swap at execute, Tactical Mastery baseline 10 / 25 with ITM; Prot 6/0/45 with Bastion 5/5.
- https://foreverdb.net/class/Warrior/builds (no date): Arms 60 builds with Improved Slam 2/2 ("cooldown 3 sec shorter"); claims Battle Shout rank 7 at 60.
- https://wowcalc.net/en/forever/guides/classes/warrior/arms (no date): Arms priority Execute > MS > Overpower > Slam (Improved Slam).
- https://wowforeverbuilds.com/guide/arms-warrior-pve-guide (2026-09-17, pre-rebuild): Arms priority; Recklessness for burn phases.
- https://mmonster.co/wow-forever/tools/rotation-trainer/warrior (no date): Arms action bar (Slam 15 rage, 18 s, 1.5 s cast), Bloodthrill 20%.
- https://wowtbc.gg/warcraftforever/class-guides/fury-warrior/ (no date): Fury priority, cooldowns "ideally used during the execute phase", Hamstring as a rage dump.
- https://mobalytics.gg/wow-forever/classes/fury-warrior-guide (updated 2026-10-28): level-30 Fury only; no level-60 data.
- https://wowclassicforever.info/classes/warrior/fury/ (2026-10-06): October-hotfix Fury talent values (DWS 25% / 50% off-hand rage, FP 10%).

## Rogue (Assassination, Combat, Subtlety)

Date: 2026-10-31. Site worktree `tc-physical` at 4e5dde099 (origin/main), engine fork `/Users/jh/code/wowsims-forever` HEAD 1f542def3, data build 1.60.1.70658. No code, data or APL was changed. The measurements below come from `sim/cmd/spec-breakdown` (built to the scratchpad; `make simdb` was run once, which writes only git-ignored embed copies). Unless noted: alliance night elf, raid preset, 5000 iterations, seed 1. The baselines reproduce the published set_dps (Combat 610.3, Assassination 644.5/645.4, Subtlety 524.0). Rotations and talent strings used for the experiments are in `/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/rogue/`.

### Headline: why Assassination sits above Combat

Our Assassination is ahead of Combat (645 against 610) for one main reason, and it is a bug. **The engine lets Mutilate fire from any pair of one-handers, and the gear search gave Assassination a mace main hand and a sword off hand** (Greenhammer 279261, subclass 4 Mace; Ironwood Blade 279259, subclass 7 Sword). The client and Blizzard both require daggers. With daggers (Heartseeker plus Scarlet Kris, the same pair Combat wears) the published Assassination build sims **589.6, 8.5% below its 644.5** and below Combat. Mutilate's normalized-weapon term also uses the 2.4 one-hander speed instead of the dagger's 1.7, which is most of that gap.

Two more Mutilate bugs pull in opposite directions:
- The flat bonus is over-credited by a third (+67 per hand where the client gives +50).
- Seal Fate never fires off Mutilate, because the Mutilate hand spells don't carry the builder flag.

With all three fixed, plus the shared fixes below, my estimate puts Assassination and Combat level at about 615-620. Forever's reworked trees do not justify a 6% Assassination lead. Combat is not missing a talent: Hack and Slash, Blade Flurry, Adrenaline Rush, Weapon Expertise, Aggression and Dual Wield Specialization all match the client text. Combat Swords losing to Combat Daggers is a real Forever change, not under-modelling.

Shared by every rogue spec:
- Deadly Poison uses the SoD damage spell: 27 per tick where Forever's enchant applies 18.
- The rogues get no Thistle Tea, though every other DPS class's preset carries its class potion.
- No sharpening stones are actually applied: the engine has one imbue slot per hand, and the poison takes it.

### Rogue — Assassination

(a) Public TC for Forever: Warcraft Tavern's level-60 build goes deep Assassination with Malice, Ruthlessness, Improved SnD, Relentless Strikes, Murder, Lethality 5, Improved Poisons 5, Vile Poisons 5, Seal Fate 5, Venom, Mutilate, plus Opportunity, Puncturing Wounds and Precision. Its boss rotation keeps up Expose Armor (when no warrior sunders), Slice and Dice and Venom, and spends the spare points on Eviscerate. Blizzard's Deep Dive says the same thing: Assassination spends its finishers on Venom and Slice and Dice, with Eviscerate and Rupture behind them. Mutilate needs daggers in both hands, and it fires Seal Fate exactly once per cast.

We sim a 28/21/2 build (no Seal Fate, no Venom, Vile Poisons 2) on a mace plus sword. Our rotation is Cold Blood and Eviscerate at 4+ combo points, Slice and Dice ahead of Eviscerate, and Mutilate, with Venom never cast. The weights look fine: agi 1.00, hit 1.75, crit 0.96, str 0.50, AP 0.45, haste 7.53 per 1%. Most of the build and rotation gap comes from engine bugs, not from the APL.

(b)

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| 1 | Mutilate weapon requirement | Mutilate requires daggers in both hands | Any dual-wield pair. The BiS band equips Greenhammer (Mace) + Ironwood Blade (Sword) | Client: wowhead Forever tooltip 1241584 "Requires Daggers"; items.json 279261 subclass 4, 279259 subclass 7. Engine `sim/rogue/mutilate.go:132-136`: ExtraCastCondition checks only `IsDualWielding`. Normalized damage at line 154-155 then uses the 2.4 one-hander speed | Blizzard Deep Dive 2026-10-07 ("both must be daggers"); classicwowforever.com; talentsforever.com | CONFIRMED-BUG | **-8.5%** measured: 644.5 -> 589.6 with `-gear main_hand:12783,off_hand:5267`. A gear re-search may claw back about 1% |
| 2 | Seal Fate from Mutilate | Each Mutilate that crits with either hand adds 1 combo point (once per cast) | Never procs. A Seal Fate build shows no Seal Fate combo-point row at all (Mutilate 70.9 CP, Ruthlessness 10.3, nothing from 14195) | Engine `sim/rogue/talents.go:249`: proc requires `SpellFlagBuilder` on the spell that dealt the hit. Mutilate's hits come from MutilateMH/OH, flagged at `mutilate.go:91,103` without the builder flag; only the parent (line 116) has it, and the parent never deals a hit | Blizzard Deep Dive; Warcraft Tavern Forever Assassination guide; classicwowforever.com rank text | CONFIRMED-BUG | Seal Fate is worth 0 today, so the build search drops it. Measured: Seal Fate 5 for Lethality 4 + Remorseless 1 = -8 DPS (589.6 -> 581.6). Fixed, estimate **+4 to +8%** for a Seal Fate build. Derivation: about 66% of Mutilates crit at least once (42% per hand), so +0.66 CP per cast, about +23 CP on 81 per fight, into Eviscerate at about 300 per CP, so +38 DPS, less the 8 DPS of displaced talents |
| 3 | Mutilate flat damage | Rank 4: 75% weapon damage "plus an additional 50 with each weapon" | +67 per hand, added after the 75% (flat + 0.75 x weapon) | Client SpellEffect 1241586/1241590: effect 121 base 67 with effect 31 75%. The tooltip's 50 = 67 x 0.75, and rank 1's 17 = 23 x 0.75, so the flat sits inside the percentage, as Backstab's does (25300: (weapon+150) x 150% = "plus 225"). Engine `mutilate.go:154-155` vs `backstab.go:30,70`, which multiplies the flat | wowhead Forever tooltips 1241584 / 1310707; endgametools.com (17 at rank 1, 50 at rank 4) | CONFIRMED-BUG | **-2.1%** est.: 17 x 1.2 poisoned x 1.10 Opportunity x 1.55 crit factor, about 35 per hit, x 64 hits / 180 s, about 12 DPS on dagger gear |
| 4 | Deadly Poison damage | DP IV: "72 Nature damage over 12 sec" (18 per tick). DP V: 92 (23 per tick) | 27 per tick (SoD damage spell 434315), stacked x5 | Client: SpellItemEnchantment 627 (DP IV) triggers 11354; SpellEffect 11354 base 18, period 3000; 2630 -> 25349 base 23. Engine `sim/rogue/client_damage.go:79-90` reads 434312-434316 (9/13/20/27/34), the SoD rune-era family, not the spells Forever's enchants apply | wowhead Forever tooltips 11356, 11354, 25347 vs classic tooltip 11356 (108 over 12 s) | CONFIRMED-BUG (base). Whether Forever adds AP scaling or tick crits on top is OPEN (row 9) | **-2.2%** est.: DP is 42.4 DPS of 645, x 1/3 |
| 5 | Thistle Tea | Standard rogue raid consumable (+100 Energy on a 5 min timer) | Not in the rogue raid preset. Warriors get Mighty Rage Potion, casters a mana potion | `data/builds/1.60.1.70658/bis/rogue-assassination.json` presets.raid.consumes; engine supports it (`sim/rogue/thistle_tea.go`, id `conjured_rogue_thistle_tea`) | Warcraft Tavern Forever Combat guide (Energy section); Classic Era practice | CONFIRMED-BUG (preset parity) | **+2.7%** measured: 645.4 -> 662.8 |
| 6 | Sharpening stones with poisons | Forever lets poisons sit alongside stones and weightstones | Preset lists Elemental + Dense stones, but the poison kit overwrites them. One imbue enum per hand (`proto/common.proto:725-726`), so stones never apply (the printed consume list has no stone) | spec-breakdown consume line; IDS.md "Both weapon imbue slots hold one enum" | Blizzard Deep Dive (poisons alongside temporary enchantments); classicwowforever.com | CONFIRMED-BUG (engine/preset) | About **+1.2%** (Combat measured +7.4 DPS for stones over no imbue; Assassination similar) |
| 7 | Talent build | Deep Assassination with Seal Fate 5, Venom, Vile Poisons 5 (WT; classicwowforever "Venom Assassination 38/11/2", "Vile Poisons 26/25/0") | 28/21/2: Seal Fate 0, Venom 0, Vile Poisons 2, Lethality 4, Precision + Dual Wield Specialization from Combat | `bis/rogue-assassination.json` band 60 raid talents `01532310421501000-31530300001500000-0020000000000000000` | WT guide; classicwowforever.com builds | OPEN (downstream of rows 2-4). Re-run the talent search after those fixes | n/a |
| 8 | Venom in the rotation | Keep Venom up (Blizzard's stated design) | Not cast. The APL note says it measured as a loss and the engine multiplies Venom with Vile Poisons where the client adds them | `data/curated/apl/rogue-assassination.json` notes; `sim/rogue/venom.go` `scalePoisonDamage` multiplies | Blizzard Deep Dive; WT | OPEN. Re-measure after rows 2 and 4 (DP inflation over-values poisons; Seal Fate adds the combo points Venom needs) | n/a |
| 9 | Poison AP scaling / crits | WT: "Poisons ... scale with Attack Power" and "can now deal critical strikes" | Instant Poison crits; DP ticks never crit; no AP term on either | Client DB2 has no AP coefficient for any rogue effect, including Eviscerate, whose Forever tooltip does say "increased by Attack Power", so client data cannot settle it. Malice text: "critical strike chance with all attacks and Poisons" | Warcraft Tavern Forever Assassination + Combat guides (single source family) | OPEN. A beta combat log of DP tick sizes at two AP values would settle it | Unknown. Could offset row 4 |
| 10 | Murder on the target | Only vs Humanoid/Giant (Forever dropped Beasts/Dragonkin) | Raid target defaults to Humanoid, so +4% always | `sim/api/envelope.go:273` (engine default MobTypeHumanoid); `sim/rogue/talents.go:140-157` | Blizzard Deep Dive (Murder change) | OPEN (target-type policy; MC/Onyxia bosses are mostly elementals, giants, beasts, dragonkin) | Up to -4% on non-humanoid bosses, same for all three specs |

(c) Verdict: 6 confirmed bugs (Mutilate dagger requirement, Seal Fate via Mutilate, Mutilate flat, Deadly Poison base, no Thistle Tea, no stones). The tier position is not justified as it stands.

### Rogue — Combat

(a) Public Forever TC treats Combat as the standard raid rogue: 18/33/0 or 20/31/0 with Improved SnD 3, Malice, Ruthlessness, Relentless Strikes, Lethality, Flawless Execution, Dual Wield Specialization, Blade Flurry, Hack and Slash, Weapon Expertise, Aggression and Adrenaline Rush. Warcraft Tavern also lists a dagger-only variant with Puncturing Wounds and Opportunity. Every guide makes Slice and Dice the top priority, kept up all fight. Eviscerate takes the spare points and Rupture is never used.

We sim a 16/33/2 dagger build (no Improved SnD) and spend combo points on Eviscerate at 5. Slice and Dice is cast only in the first 15 s, for **8.9% uptime**. The engine values Forever's buffed Eviscerate highly: Flawless Execution, Improved Eviscerate 20%, Aggression and AP scaling give 1721 per 5-point cast. Keeping SnD up measures within error of the shipped rotation, so the uptime gap costs little DPS, even though it reads wrong. Swords lose to daggers in the engine, and that holds up as a Forever change. The weights fit public priorities: agi 1.00, hit 1.50, crit 0.91, str 0.53, AP 0.48, haste 10.81 per 1%.

(b)

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| 1 | Slice and Dice uptime | Keep SnD up all fight; it is the top-priority finisher | SnD only while `currentTime < 15s`; uptime 8.9%, 1 cast per fight | `data/curated/apl/rogue-combat.json` line 2 condition; spec-breakdown aura table | WT Forever Combat guide; WT Forever Assassination guide; Era consensus (mechanic unchanged: client 6774 +30% attack speed, 9+3 s/CP) | OPEN (rotation reads wrong, model says it doesn't matter). Measured: SnD at >=3 CP whenever down (variant A) 612.1 vs 610.3; SnD-first variants 605.9-610.5 at 77-95% uptime. White damage rises 25% but Eviscerate falls by the same amount | +0.3% at best |
| 2 | Second Blade Flurry lost | BF on cooldown (2 casts in 180 s) | 1 cast. The engine's last-BF gate waits for SnD >= 1 s remaining, and SnD is never up after 15 s | Engine `sim/rogue/talents.go:456-465`; breakdown BF casts 1.00 (2.00 with variant A) | WT Combat guide ("Blade Flurry on cooldown") | CONFIRMED-BUG (APL/engine interaction; folded into row 1's measurement) | Included in row 1's +1.8 DPS |
| 3 | Swords vs daggers | Era: Combat Swords dominates. Forever: Hack and Slash gives swords/axes 5% extra attacks, daggers/fists 5% crit, maces 15% armor ignore; Weapon Expertise is dodge/parry only (no weapon skill) | Daggers chosen. Best sword setups sim 571.6-590.9 against 610.3 for daggers | Client talents 105727 (Hack and Slash), Weapon Expertise rank text; engine `talents.go:630-712` matches the ranks (200 ms ICD on the extra attack is an engine choice) | Blizzard Deep Dive (Combat supports Backstab builds; Aggression and Puncturing Wounds buff Backstab); endgametools.com; classicwowforever.com | FOREVER-CHANGE (sword weapon-skill and Sword Spec advantages are gone; Backstab got Puncturing Wounds +30% crit and Aggression) | none |
| 4 | Deadly Poison damage | 18 per tick at DP IV | 27 per tick | Same as Assassination row 4 | same | CONFIRMED-BUG | **-1.8%** (DP 33.4 DPS x 1/3) |
| 5 | Thistle Tea | Used | Absent from preset | Same as Assassination row 5 | same | CONFIRMED-BUG | **+1.9%** measured: 610.3 -> 622.1 |
| 6 | Sharpening stones | Stones plus poisons | Poisons only | Same as Assassination row 6 | same | CONFIRMED-BUG | **+1.2%** (stones over no imbue: 566.5 -> 573.9) |
| 7 | Talent build | Improved SnD 3, often Cold Blood or Improved Expose Armor; Murder optional | Improved SnD 0, Murder 2, Lightning Reflexes 5 as filler, Opportunity 2 | `bis/rogue-combat.json` talents `00532010500000000-31530300001515231-0020000000000000000` | wowfdb.com 18/33/0; WT 60 builds | OPEN (follows from the engine's low SnD value; consistent with row 1) | n/a |
| 8 | Talent values | Aggression 6%, Improved Eviscerate 20%, Dual Wield Specialization 25%, Flawless Execution -10, Adrenaline Rush 15 s/5 min, Blade Flurry +20% speed/2 min | Same | `sim/rogue/client_values.go:14-71`, `talents.go:373-522`; client talents/rogue.json and spellconst 13750/13877 | Deep Dive; WT (WT's "DWS up to 50%" is the stale Era figure) | aligned | none |
| 9 | Energy model | Continuous 10 Energy/s (WT: 1 per 0.1 s) | Continuous 10/s | `sim/core/energy.go` per design/reviews/2026-10-08-rogue-energy.md | WT Combat guide; Blizzard 2026-09-30 class highlights | aligned | none |
| 10 | Expose Armor vs Sunder | Forever EA equals 5 Sunders and does not stack with Sunder | Not cast; Sunder in raid preset | Deep Dive Expose Armor text; preset debuffs | Blizzard Deep Dive | aligned | none |
| 11 | Tier position | Guides call Combat the strongest PvE rogue; DistillNote's own sim had Combat at 707 (2026-10-28) | 610, about 21% below Fury 775 | n/a | WT, DistillNote | OPEN (no Forever logs at 60 yet; level 60 is not testable on beta) | n/a |

(c) Verdict: 4 confirmed bugs (Deadly Poison base, Thistle Tea, stones, second Blade Flurry). The Slice and Dice uptime is open: wrong-looking but worth under 0.5% in this model. Combat is not under-modelled in its talents.

### Rogue — Subtlety

(a) Public Forever Subtlety builds (classicwowforever "Preparation and Hemorrhage 22/3/26", "Thousand Cuts 11/8/32"; DistillNote) are PvP-leaning, and no source rates Subtlety a raid spec. The Deep Dive frames it as Bleeds and combat tricks.

We sim 20/0/31 Hemorrhage, Rupture, Ghostly Strike with a Stealth, Premeditation, Ambush opener. Three engine or APL faults cost it DPS:
- The pre-pull Ambush deals no counted damage but spends 60 Energy.
- Ghostly Strike is starved to 1.5 casts a fight.
- The forced filler includes Cutthroat 3, which does nothing with a Hemorrhage builder, where Quietus would add damage.

The weights are agi 1.00, hit 1.41, crit 0.85, str 0.55, AP 0.50, haste 6.60 per 1%.

(b)

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| 1 | Pre-pull Ambush | Ambush opener deals full damage | Cast at -1 s: 0.96 hits, **0 damage recorded**, 60 Energy spent before the pull | `data/curated/apl/rogue-subtlety.json` prepullActions (11269 at -1s). Measured: dropping it 527.6; casting it at 0 s (still stealthed) 531.0 with a 1450 average hit, against 524.0 shipped | Era mechanic (unchanged) | CONFIRMED-BUG (pre-pull damage not counted; worth checking every spec with damaging pre-pull actions) | **+1.3%** measured |
| 2 | Quietus vs filler | n/a (no Forever raid Subtlety TC) | Cutthroat 3 (procs only from Backstab, which this APL never casts) where Quietus fits in the same tier | Client talents 110867 (Quietus) and 105750 (Cutthroat); engine `talents.go:716-745` (Cutthroat keyed on rogue.Backstab), `talents.go:822-865` (Quietus implemented). Talent `...-5323200310013011501` (Dirty Tricks 2 + Cutthroat 3 -> Quietus 5) measured 530.5 vs 524.0 | client rank text | CONFIRMED-BUG (build-search miss; tier gating allows it) | **+1.2%** measured |
| 3 | Ghostly Strike usage | On cooldown (20 s) | 1.52 casts per fight: Hemorrhage at 35 Energy always fires before Energy reaches GS's 40 | APL line 1 has no pooling; breakdown GS casts 1.52 | APL notes say "on cooldown" | CONFIRMED-BUG (APL) | small (GS 556 at 40 Energy vs Hemorrhage 527 at 35): under 0.5% |
| 4 | Deadly Poison damage | 18 per tick | 27 per tick | Same as Assassination row 4 | same | CONFIRMED-BUG | **-2.7%** (DP 42.0 of 524, x 1/3) |
| 5 | Thistle Tea | Used | Absent | Same | same | CONFIRMED-BUG | **+2.3%** measured: 524.0 -> 536.1 |
| 6 | Sharpening stones | Stones plus poisons | Poisons only | Same | same | CONFIRMED-BUG | about +1% |
| 7 | Hemorrhage, Ghostly Strike, Serrated Blades, Thousand Cuts values | 145%/180% dagger, +15% Rupture taken, 9% armor + 30% Rupture, -3 Energy x5 | Same | `client_values.go:44-52`, `talents.go:752-811`, `rupture.go:59` | client rank text | aligned | none |
| 8 | Tier position | Not a raid spec in any source | Last rogue spec, 524 | n/a | DistillNote; classicwowforever | aligned | none |

(c) Verdict: 6 confirmed bugs. Subtlety's last place is right. The fixes move it about +3% net, and it stays last.

### Fix-lane priority (ranked by DPS impact)

1. Mutilate requires daggers in both hands (`mutilate.go:132`): Assassination -8.5%.
2. Seal Fate must proc off MutilateMH/OH, once per cast (`talents.go:249` / `mutilate.go:91,103`): +4 to +8% for a Seal Fate Assassination build. Then re-run the Assassination talent search, with Venom and Vile Poisons.
3. Thistle Tea in the rogue raid preset: +1.9 to +2.7% for all three specs.
4. Deadly Poison damage from 11354/25349 (18/23 per tick), not 434315/434316 (`client_damage.go:83-90`): -1.8 to -2.7%.
5. Mutilate flat inside the 75% (`mutilate.go:154-155`): -2.1%.
6. Subtlety: pre-pull Ambush to 0 s (+1.3%), Quietus over Cutthroat (+1.2%), Ghostly Strike pooling.
7. Poison and stone in the same hand (proto has one imbue per hand): +1.2%.
8. Combat: second Blade Flurry and Slice and Dice upkeep (+0.3%; also makes the rotation read correctly).

Net, my estimate: Assassination about 615-620, Combat about 619, Subtlety about 540. Assassination and Combat end up level, not 6% apart.

### Sources

- Blizzard "World of Warcraft: Forever" Rogue & Warlock class deep dive, 2026-10-07, reproduced in full at https://www.warcrafttavern.com/forever/news/poison-pets-rogue-warlock-deep-dive-for-wow-forever/. Supports: Mutilate needs daggers, Seal Fate once per Mutilate, Assassination's Venom/SnD finisher focus, EA = 5 Sunders, poisons alongside stones, Malice/Precision cover poisons, Murder target types, Hack and Slash design, Aggression on Backstab.
- Warcraft Tavern, Assassination Rogue Guide (Forever), https://www.warcrafttavern.com/forever/guides/assassination-rogue/ (fetched 2026-10-31). Supports: level-60 build with Seal Fate/Venom/Vile 5, boss rotation, poison AP/crit claims.
- Warcraft Tavern, Combat Rogue Guide (Forever), https://www.warcrafttavern.com/forever/guides/combat-rogue/ (fetched 2026-10-31). Supports: SnD top priority, no Rupture, dagger variant, continuous 10 Energy/s, Thistle Tea.
- classicwowforever.com Rogue talents / guide / class changes, https://classicwowforever.com/talents/rogue/ (updated 2026-10-19, build 70450). Supports: community builds (Venom 38/11/2, Vile 26/25/0, Combat 18/33 and 22/29, Subtlety 22/3/26 and 11/8/32), Mutilate requires daggers, stones plus poisons.
- wowfdb.com Combat talents, https://wowfdb.com/classes/rogue/combat/talents/ (build 70707). Supports: Combat 18/33/0 projected raid build.
- endgametools.com, "WoW Forever Rogue: Mutilate, Hack and Slash and new talents", 2026-09-27 (updated for build 70532), https://endgametools.com/en/wow-forever/news/wow-forever-rogue. Supports: Mutilate rank text (17 at rank 1, 50 at rank 4), talent list.
- wowclassicforever.info rogue talents, https://wowclassicforever.info/talent/rogue/. Supports: tree changes, Hack and Slash merger.
- DistillNote Rogue page, https://distillnote.com/classes/rogue (2026-10-28 run). Supports: Combat recommended for raids, their sim at 707.
- Wowhead Forever tooltips (client text), fetched 2026-10-31: spell 11300 (Eviscerate "increased by Attack Power"), 6774, 13750, 1241584, 11356, 11354, 25347, 11340; items 279261, 279259, 12783, 5267. Classic Era tooltips 11356/25347 for the 108/136 comparison.
- wago.tools DB2 for 1.60.1.70658: SpellItemEnchantment (627 -> 11354, 2630 -> 25349) and SpellEffect (11354 = 18, 434315 = 27, 1241586 effect 121 = 67 with effect 31 = 75%), fetched 2026-10-31.

## Hunter (Beast Mastery, Marksmanship, Survival)

Audited 2026-10-31 against site worktree `tc-physical` (origin/main 4e5dde099), data build 1.60.1.70658, engine fork HEAD 1f542def3 (`/Users/jh/code/wowsims-forever`). Survival is treated as a melee spec (owner rule). The measured figures come from `sim/cmd/spec-breakdown`: raid preset, band 60, alliance, 3000 iterations, seed 1, run from the month-test worktree. Its hunter APLs and bis entries are byte-identical to tc-physical's. The baselines it reproduced were BM 710.0, MM 702.2 and SV 688.3, against published values of 712 / 703 / 687. Rotation variants were passed with `-rotation` and talent variants with `-talents`. No code or data was changed.

Damage shares at baseline. Pet DPS is the total minus the hunter's own action rows.

| Spec | Total | Hunter | Pet (share) | Top hunter rows |
|---|---|---|---|---|
| BM | 710.0 | 464.8 | 245 (34%) | Auto Shot 283, Aimed 69.6 (17.5 casts), Multi 66.6 (21.2 casts), Serpent 25.2, Arcane 20.4; out of mana 32 s of 180 |
| MM | 702.2 | 540.5 | 162 (23%) | Auto Shot 315.8, Aimed 107.5 (21.8 casts), Multi 87.6 (22.7 casts), Arcane 29.6 |
| SV | 688.3 | 553.5 | 135 (20%) | White 250, Raptor 94.3, Mongoose 68.9, Strider Kick 57.1, Immolation Trap 25.7, Wing Clip 23.6, Lacerate 18.8, Lacerating Strikes 14.7. No aspect active. |

### Shared hunter findings (all three specs)

These come up in each spec table below and are described once here.

- **Pet model (H-PET).** Forever reworked hunter pets. Blizzard's Deep Dive, as reported by Warcraft Tavern on 2026-09-30, says pets "gain stats based on the Hunter's gear, and they can no longer receive player buffs that boost their stats". Warcraft Tavern's Forever hunter guide and its pet guide, plus the MMOGAH BM guide, give the numbers: the pet inherits 10% of the hunter's Attack Power (whichever of melee or ranged is higher) and 100% of the hunter's crit chance.
  - The engine does neither. `sim/hunter/pet.go:206-210` `makeStatInheritance` returns an empty `stats.Stats{}` ("No stat inheritance in classic").
  - `sim/core/buffs.go:436-461` `applyPetBuffEffects` hands the pet every raid buff. That includes Battle Shout, Blessing of Might, Strength of Earth, Mark of the Wild and Heart of the Lion's +10% and 276 AP. It also includes Trueshot Aura, which the engine applies to melee AP (see H-TSA).
  - Rough size: the pet's base AP is 252 (`pet.go`: Str 136 x 2 - 20), and the raid buffs add roughly 900 AP. The rework would leave 252 plus 10% of about 2000 hunter AP, so about 450. That roughly halves the pet's white damage, partly offset by crit inheriting from about 20% to about 26%.
  - That puts the pet at about -40 to -50% of its current DPS: BM about -8 to -15%, MM about -6 to -10%, SV about -5 to -9%. This is a hand estimate, not a run. The rework also makes hunter crit and AP raise pet damage, which the weights currently miss.
- **Multi-Shot and Aimed Shot share one cooldown (H-SHOT).** The Forever client tooltip for Multi-Shot 2643 reads "Multi-Shot shares its cooldown with Aimed Shot." Aimed Shot 20904 reads "Aimed Shot shares its cooldown with Multi-Shot." The Deep Dive says the same, and that Arcane Shot no longer shares its cooldown.
  - The engine gives each shot its own timer: `sim/hunter/hunter.go:201-207` (`multiShotTimer`, `arcaneShotTimer`, `aimedShotTimer`), `aimed_shot.go:38-41` and `multi_shot.go:71-74`. The BM and MM APLs fire both on cooldown, which is impossible in the client.
  - I measured two lower bounds by removing one line. BM drops to 683.5 without Aimed Shot (-3.7%) and 673.4 without Multi-Shot. MM drops to 636.2 without Aimed Shot (-9.4%) and 621.3 without Multi-Shot.
  - A shared-timer priority (Aimed Shot when the auto-shot timing allows, Multi-Shot otherwise) will land a little above those. Estimate BM -3 to -4%, MM -7 to -9%.
- **Trueshot Aura gives melee AP in the engine (H-TSA).** Every Trueshot Aura rank in the client is a single effect, aura 124 (ranged attack power): 1299346 at 30, 1299348 at 40, 19506 at 50, 20905 at 75, 20906 at 50. The talent text and the 20906 tooltip both say "Ranged Attack Power". Warcraft Tavern's Forever guide says it "now affects only Ranged Attack Power".
  - `sim/core/buffs.go:1422-1447` (`TrueshotAura`) adds the amount to `stats.AttackPower` as well, under the comment "melee and ranged attack power alike as the client's single effect does". That comment is wrong: melee AP would be aura 99, as Heart of the Lion 409583 shows.
  - Every melee spec in the raid preset, and the hunter pet, gets +75 melee AP it should not have. SV: 75 x 0.1407 DPS/AP = about 10.6 DPS (-1.5%). Warrior, rogue, feral, enhancement and ret should each lose a similar 1 to 1.5%; worth telling the other class lanes.
- **Hydra Shot (H-HYDRA).** The client has Hydra Shot 1293020: level 60, 250 mana, instant, weapon damage + 260 split across 5 targets. The Wowhead Forever tooltip says it shares its cooldown with Arcane Shot. `trainables/hunter.json` lists it as `source: class_spell` with no skill line, the same tag Widow Bite (a pet spell) and Wyvern Strike (an SoD rune) carry.
  - `design/reviews/2026-10-07-hunter-hydra-lacerate-pets.md` found no SkillLineAbility learn row and removed it from the engine. Only the constants remain, at `constants_auto_gen.go:1602-1611`.
  - No web source I found says how a player gets it. OPEN: a trainer row in a newer build, or a beta report of it in a spellbook, would settle it.
  - If it is learnable it replaces Arcane Shot on the same timer. Weapon damage + 260 against Arcane's 217 (291 average hit) is roughly +10 to +15 DPS for BM and MM (+1.5 to 2%).

### Hunter — Beast Mastery

**(a) Summary.**
- **Public TC:** Forever guides call BM the strongest hunter spec, specifically because of the new pet scaling and Forever's BM talents. Era TC puts MM-leaning builds ahead of BM at 60, but that does not carry over: the trees are Forever's own, with no Lethal Shots, Killer Instinct or The Beast Within.
- **Our build** is 16/10/4: Deadly Aspects 5, Focused Fire 2, Unleashed Fury 5, Ferocity 5, Frenzy 5, Bestial Wrath, plus Lethal Attacks 5, Efficiency 5, Rapid Killing 2 and Trueshot Aura. Every node is in the live tree and every talent rate matches the client text.
- **The rotation** casts Aimed Shot and Multi-Shot on separate timers, which the client forbids (H-SHOT).
- **The pet** carries every raid stat buff and inherits nothing from the hunter, both against Forever's rework (H-PET).
- **Weights** are internally consistent. Agility is worth 2.42 RAP after Heart of the Lion's 1.21, plus 1/53% crit, which Warcraft Tavern confirms at 60. Both routes value 1% crit at 2.9 to 3.0 DPS. Crit at 0.50 is therefore not a rating-conversion error: it is low because the pet does not inherit hunter crit.

**(b) Differences.**

| # | Area | Public TC | Our sim | Evidence (client row / engine file:line / APL path) | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| 1 | Aimed / Multi cooldown | Multi-Shot and Aimed Shot share one 6 s cooldown | Separate timers; both fire every 6 s (17.5 + 21.2 casts) | Tooltips 2643 and 20904; `hunter.go:201-207`; `data/curated/apl/hunter-beast-mastery.json` lines 3-4 | Wowhead Forever tooltips; WT Deep Dive 2026-09-30 | CONFIRMED-BUG | -3 to -4% (measured lower bound 683.5, -3.7%) |
| 2 | Pet stats | Pet takes 10% of hunter AP and 100% of hunter crit, and no player stat buffs | No inheritance; every raid buff | `pet.go:206-210`, `core/buffs.go:436-461` | WT Deep Dive; WT Forever hunter guide; MMOGAH BM guide | CONFIRMED-BUG (mechanic); size OPEN | -8 to -15% (hand estimate, H-PET) |
| 3 | Crit weight 0.50 | Hunter crit also raises pet crit, so it is worth more than its personal value | Crit counts only the hunter's own hits | Weights in `bis/hunter-beast-mastery.json`; agility→crit at 53 per 1% matches WT | WT Forever guide | Consequence of #2 | Weight order only: crit moves to about 0.7-0.8 of agility with the pet's 34% share |
| 4 | Trueshot Aura | Ranged AP only | Melee and ranged AP; the pet's melee AP also gains 75 | Client 20905 aura 124 only; `core/buffs.go:1435-1446` | WT Forever guide; talent text | CONFIRMED-BUG | BM hunter none; pet small (inside #2) |
| 5 | Hydra Shot | n/a (no source on how it is learned) | Not registered | Trainables `class_spell` only | Wowhead tooltip 1293020 | OPEN | +1.5 to 2% if learnable |
| 6 | Serpent Sting / mana | Era raid BM skips Serpent Sting | Casts it; goes out of mana for 32 s | Measured without Serpent Sting: 710.2 vs 710.0 | WT Classic BM rotation | Aligned (neutral) | 0% |
| 7 | Hit cap display | Ranged hit cap 9% against a boss (WT) | `hit_to_cap.white` = 28 (melee dual-wield cap) published for a ranged spec | `sim/cmd/leveling-bis/hitcap.go:37-40` keys on DualWielding, not on ranged | WT Forever guide | CONFIRMED-BUG (display) | 0 DPS |
| 8 | Stat label | Haste speeds ranged attacks | Guide `statPriority` says "Melee haste" for a ranged spec | `web/src/content/guides/hunter/beast-mastery.md` frontmatter | — | CONFIRMED-BUG (copy) | 0 DPS |
| 9 | The Beast Within, Kill Command, Steady Shot | Era/SoD BM tools | Absent | Not in the live BM tree; Kill Command and Steady Shot are SoD runes with no trainer row | Client talents/trainables | FOREVER-CHANGE (we're right) | — |
| 10 | Pet family | No Forever 60 raid pet consensus found | Cat (Bite/Claw), no family ability | `sim/request/request.go:680` | WT Deep Dive (family abilities exist) | OPEN | Small (Savage Rend etc. are modelled) |

**(c) Verdict:** 2 confirmed sim bugs (#1 measured, #2 mechanic confirmed with size open), plus 2 display/copy bugs. BM's 712 is overstated by about 3 to 4% from shared shots alone, and likely more once the pet is corrected.

### Hunter — Marksmanship

**(a) Summary.**
- **Public TC:** Forever guides call MM "very competitive" in beta but give no level-60 raid build.
- **Our build** is 10/20 plus 5 Improved Tracking, all in the live tree: Lethal Attacks, Careful Aim 5, Mortal Shots 5, Barrage 3, Ranged Weapon Specialization 5 and Trueshot Aura.
- **Shared shots:** MM is the spec most hurt by H-SHOT, because its APL alternates Aimed Shot and Multi-Shot.
- **Careful Aim:** the engine gives only melee AP, but the client and the Deep Dive both say melee and ranged AP. So MM currently gets nothing from 5 points.

**(b) Differences.**

| # | Area | Public TC | Our sim | Evidence | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| 1 | Aimed / Multi cooldown | One shared 6 s cooldown | Separate timers (21.8 + 22.7 casts) | Tooltips 2643 and 20904; `hunter.go:201-207`; `data/curated/apl/hunter-marksmanship.json` lines 2-3 | Wowhead tooltips; WT Deep Dive | CONFIRMED-BUG | -7 to -9% (measured lower bound 636.2, -9.4%) |
| 2 | Careful Aim | "increases melee and ranged Attack Power by 20-100% of Intellect" | Melee AP only | Client 1223984 has two effects, aura 580 and aura 598, both misc 3 (Intellect). Aura 598 alone is the melee-AP-from-stat aura (Cat Form, Mental Dexterity). `talents.go:499-505` adds only `stats.AttackPower` | WT Deep Dive 2026-09-30; client row | CONFIRMED-BUG | +4 to 6% (about 230 raid-buffed Int x 0.169 DPS per RAP = about 39 DPS; Int estimated, not read from the sim) |
| 3 | Pet stats | As H-PET | As H-PET | As BM #2 | As BM #2 | CONFIRMED-BUG (mechanic); size OPEN | -6 to -10% (pet share 23%) |
| 4 | Trueshot Aura | Ranged AP only | Melee and ranged | As BM #4 | As BM #4 | CONFIRMED-BUG | Pet only, inside #3 |
| 5 | Mortal Shots | Ranged crit damage +30% | `mortalShots()` 0.06 per rank on shots and Auto Shot | `talents.go:306-308`; tooltip 19485 | Tooltip | Aligned for MM | — |
| 6 | Hydra Shot | Unknown | Not registered | As H-HYDRA | Tooltip | OPEN | +1.5 to 2% if learnable |
| 7 | Hit cap display / stat label | 9% ranged cap; haste is ranged | `white` = 28; "Melee haste" in guide | As BM #7-8 | WT | CONFIRMED-BUG (display/copy) | 0 DPS |
| 8 | Aimed Shot cast and clipping | 2 s cast (Forever tooltip; Era was 3 s) | 2000 ms, scaled by ranged speed; cancels and resets the auto swing; APL starts it only with more than 1 s to the next auto | `aimed_shot.go:35-47`; APL line 2 | Tooltip 20904 | Aligned | — |

**(c) Verdict:** 2 confirmed sim bugs with size (#1 about -8%, #2 about +5%), the shared pet bug, and display/copy bugs. Net, MM settles roughly 3 to 5% below its current 703 before the pet fix.

### Hunter — Survival (melee)

**(a) Summary.**
- **Public TC:** Forever reworked Survival into melee: Savage Strikes, Predator's Edge, Expose Prey, Strider Kick, Lacerating Strikes and Resourcefulness. No source gives a settled level-60 raid build or priority.
  - wowtbc.gg's Forever SV page weaves Aspect of the Hawk, Serpent Sting and Arcane/Multi-Shot with the melee strikes. That is a single source, leveling-oriented, and the owner's melee rule governs.
- **Our rotation** (Raptor queue, Mongoose Bite in its window, Lacerate, Strider Kick, Immolation Trap, Wing Clip under Windfury) fits the client kit.
- Four engine or APL defects move SV, two up and two down:
  - **Mortal Shots** is credited to melee strikes, though it is a ranged talent.
  - **No aspect:** the APL casts none, though Aspect of the Falcon (120 melee and ranged AP) is trainable and registered.
  - **Trueshot Aura** gives SV melee AP it should not.
  - **Rapid Fire** misses its melee haste half.
- **The talent build** spends 10 of its 20 Marksmanship points (Hawk Eye 3, Improved Concussive Shot 2, Mortal Shots 5) on talents a melee hunter gets nothing from. The Mortal Shots bug is what makes that look right.

**(b) Differences.**

| # | Area | Public TC | Our sim | Evidence | Sources | Class | Est. impact |
|---|---|---|---|---|---|---|---|
| 1 | Mortal Shots on melee | "critical strike damage bonus on all ranged abilities by 30%"; Predator's Edge is the melee counterpart | `mortalShots()` added to Raptor Strike, Mongoose Bite, Counterattack, Strider Kick and Wing Clip | `raptor_strike.go:97`, `mongoose_bite.go:58`, `strider_kick.go:67`, `counterattack.go:72`, `wing_clip.go:51`; tooltip 19485 | Wowhead tooltip; talent text in `talents/hunter.json` | CONFIRMED-BUG | -2.4% (measured: Mortal Shots 0 gives 672.0 vs 688.3) |
| 2 | Aspect | A melee hunter keeps an AP aspect up. Falcon: "melee and ranged Attack Power by the same amount as ... Aspect of the Hawk" (60, trainable). Beast: 110 melee AP | No aspect cast; the aura list shows none | `data/curated/apl/hunter-survival.json` prepullActions null; `sim/hunter/aspect_of_the_falcon.go` registers 469145 | Tooltips 469145 and 1299447; trainables (Falcon at 60); WT Deep Dive ("Aspect of the Beast also grants melee Attack Power") | CONFIRMED-BUG (APL) | +2.5% (measured: Falcon prepull 705.3) |
| 3 | Trueshot Aura | Ranged AP only | +75 melee AP | As H-TSA | WT guide; client 20905 | CONFIRMED-BUG | -1.5% (75 x 0.1407 DPS/AP) |
| 4 | Rapid Fire | "Increases ranged and melee attack speed by 40% for 15 sec" | Ranged speed only; SV casts it for nothing (1 cast, 8.3% uptime) | Client 3045: aura 140 (ranged haste) 40 plus aura 319 (melee haste) 40; `rapid_fire.go:20-25` calls only `MultiplyRangedSpeed` | Tooltip 3045 | CONFIRMED-BUG | +1.5% (swing-bound damage about 345 DPS x 40% x 15/180) |
| 5 | Talent build | Melee SV gets nothing from Hawk Eye, Improved Concussive Shot or Mortal Shots | Guide/raid build 0/20/31 spends 10 points there | `bis/hunter-survival.json` band 60 talents `3250050000500000` (MM tree) | Talent text | CONFIRMED-BUG (consequence of #1) | Re-run talent-search after #1. BM-tree alternatives: Focused Fire +2% to all damage, Unleashed Fury/Ferocity for the pet, Deadly Aspects' melee half (needs Aspect of the Beast) |
| 6 | Pet stats | As H-PET | As H-PET | As BM #2 | As BM #2 | CONFIRMED-BUG (mechanic); size OPEN | -5 to -9% (pet share 20%) |
| 7 | Savage Strikes | Deep Dive says "+2/4% melee damage"; client says "+4% crit chance of all your melee abilities" | 2% crit per rank | `talents.go:56-58`; `talents/hunter.json` | WT Deep Dive vs client text | FOREVER-CHANGE (client text beats the article; we're right) | — |
| 8 | Lacerate / Lacerating Strikes | Deep Dive calls Lacerate a druid ability. The client has a hunter Lacerate trainer line (24118-24120, 1299332, Survival skill line) | Lacerate cast when the bleed is down (18.8 DPS); Lacerating Strikes 14.7 DPS | Trainables; `lacerate.go` | Client trainables | Aligned (client beats the article) | — |
| 9 | Ranged weaving | One guide weaves shots between strikes | Pure melee at 0 yd (owner rule; the engine has one static distance) | `hunter-survival.json` notes | wowtbc.gg Forever SV page (single source) | OPEN (owner rule decides) | Unknown |
| 10 | Efficiency on melee | Text says "Shots, Stings, and melee abilities" | Engine excludes melee | `talents.go:365-375` | Talent text | CONFIRMED-BUG, no current impact (SV takes 0 Efficiency) | 0% with the current build |
| 11 | Hit cap | Dual-wield white cap; Surefooted +3% | baseline 3, specials 6, white 25; white miss 18-20% in the breakdown | `bis/hunter-survival.json` | — | Aligned | — |

**(c) Verdict:** 4 confirmed defects with size: #1 -2.4%, #2 +2.5% (APL), #3 -1.5%, #4 +1.5%. They roughly cancel. The build (#5) and the pet (#6) should be re-measured after the fixes.

### Tier position after fixes (projection, not a run)

| Spec | Now | Shot / talent / APL fixes | Then the pet fix |
|---|---|---|---|
| BM | 712 | about 685 | about 590-630 |
| MM | 703 | about 670 (-8% shared shots, +5% Careful Aim) | about 615-640 |
| SV | 687 | about 690 | about 630-655 |

On this projection BM's lead over MM holds or narrows. SV may end up the top hunter once the pet is modelled correctly, because it leans least on the pet. The spread is too uncertain to rank until H-PET is built and measured.

### Ranked confirmed bugs (by estimated DPS impact)

1. **Pet stat model** (H-PET; all specs): BM -8 to -15%, MM -6 to -10%, SV -5 to -9%. Size needs a model run.
2. **Multi-Shot and Aimed Shot share a cooldown** (H-SHOT): MM -7 to -9%, BM -3 to -4%.
3. **Careful Aim missing ranged AP:** MM +4 to 6%.
4. **Survival APL casts no aspect** (Falcon): SV +2.5%, measured.
5. **Mortal Shots credited to melee strikes:** SV -2.4%, measured. The SV talent build needs re-searching afterwards.
6. **Trueshot Aura gives melee AP** (H-TSA): SV -1.5%. Also about -1 to -1.5% for every other melee spec in the raid preset; cross-class.
7. **Rapid Fire missing melee haste:** SV +1.5%.
8. **Display/copy, no DPS impact:** ranged specs publish a dual-wield `hit_to_cap.white`; BM and MM guides list "Melee haste".

### Sources

- Wowhead Forever tooltips, fetched 2026-10-31, client text. Spell pages use https://nether.wowhead.com/forever/tooltip/spell/<id> with the id shown:
  - Multi-Shot 2643 and Aimed Shot 20904: shared cooldown, 2 s cast, 13.9% base mana.
  - Hydra Shot 1293020: shares a cooldown with Arcane Shot.
  - Arcane Shot 14287.
  - Summon Hawk 1293527.
  - Careful Aim 1223984.
  - Mortal Shots 19485: "all ranged abilities".
  - Trueshot Aura 20906: "Ranged Attack Power ... by 50".
  - Rapid Fire 3045: "ranged and melee attack speed by 40%".
  - Bestial Wrath 19574.
  - Aspect of the Beast 1299447 and Aspect of the Falcon 469145.
- Warcraft Tavern, "Hunter & Druid Deep Dive for WoW Forever", 2026-09-30, https://www.warcrafttavern.com/forever/news/powershifting-is-back-sort-of-hunter-druid-deep-dive-for-wow-forever/. Supports: pets gain stats from the hunter's gear and get no player buffs; Multi-Shot shares Aimed Shot's cooldown; Arcane Shot no longer does; Careful Aim is melee and ranged AP; Aspect of the Beast gives melee AP; the new Survival melee talents.
- Warcraft Tavern, Forever Hunter Guide, undated (2026, pre-launch), https://warcrafttavern.com/forever/guides/hunter/. Supports: pets take 10% of AP and 100% of crit; Trueshot Aura is ranged-only and reduced; 53 Agility per 1% crit at 60; 9% hit cap; BM strongest; MM "very competitive" in beta.
- Warcraft Tavern, Forever Hunter Pets Guide, https://www.warcrafttavern.com/forever/guides/hunter-pets/, and MMOGAH, "WoW Forever Beast Mastery hunter leveling guide: pet scaling reworked", https://www.mmogah.com/es/news/wow-forever/wow-forever-beast-mastery-hunter-leveling-guide-pet-scaling-reworked-hunters-get-their-strongest-version-yet. Search snippets, 2026-10-31: the 10% AP and 100% crit figures as a second source.
- Mobalytics Forever hunter guides (search snippets only; the pages return 403), https://mobalytics.gg/wow-forever/hunter-guides. Supports: BM is "untouchable due to the new pet scaling"; Survival is a melee spec.
- wowtbc.gg, Forever Survival Hunter PvE Guide, undated, https://wowtbc.gg/warcraftforever/class-guides/survival-hunter/. A weave priority (Hawk, Serpent, Arcane, Multi with melee strikes); single source.
- wow.gg, "WoW Forever Survival Hunter: Specialization Changes", 2026-09-19, https://wow.gg/guides/hunter-survival-forever-overview. A pre-rework tree (old talent names); used only for "melee-focused, ranged slot is a stat stick".
- warcraft.wiki.gg, Hunter abilities (Forever), edited 2026-10-19, https://warcraft.wiki.gg/wiki/Hunter_abilities_(Forever). Ability list; nothing on Hydra Shot.
- Warcraft Tavern Classic BM/MM rotation guides (Era), as cited in the APL sources. Used only for Serpent Sting practice; Forever changed the shot cooldown sharing.
- Client data: `data/builds/1.60.1.70658/spellconst/hunter.json` (Trueshot Aura ranks, Careful Aim 580/598, Rapid Fire 3045, Heart of the Lion 409583, Mortal Shots 19485), `talents/hunter.json`, `trainables/hunter.json` (also checked in 1.60.1.70707).

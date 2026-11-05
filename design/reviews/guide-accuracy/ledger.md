# Guide-accuracy ledger

Rounds run until one finds zero confirmed errors. Each round: one expert auditor per class on the active client's data (hard evidence only), class fix lanes, a wow-player SHIP per class on rendered pages, one landing. Reports: `round<N>/<class>.md` (audit, fix log, fix passes).

Columns: **Audit** = CONFIRMED findings in the audit (a finding that touches several pages counts on each); **Review** = MUSTs the rendered-page reviews found and the fix passes fixed; **Fixed** = Audit + Review fixed this round (plus LIKELY findings the lanes proved with data, in the class report); **Left** = what waits on the owner or the next sim round; **sha** = the commit carrying the round's fixes (land/guides-r1 9acd0a78c, landed on main with the branch).

## Round 1 (client 1.60.1.70755, 2026-11-04)

| Spec                 | Audit | Review | Fixed | Left                                                                                          | sha       |
| -------------------- | ----- | ------ | ----- | --------------------------------------------------------------------------------------------- | --------- |
| warrior class page   | 0     | 0      | 0     | planner links (owner, kept by R1)                                                             | 9acd0a78c |
| warrior-arms         | 1     | 0      | 1     | Arms opener Bloodrage vs Charge (owner)                                                       | 9acd0a78c |
| warrior-fury         | 2     | 1      | 3     | Improved Berserker Rage (sim)                                                                 | 9acd0a78c |
| warrior-protection   | 1     | 2      | 3     | Improved Shield Wall (owner); Expertise cap footnote (component)                              | 9acd0a78c |
| shaman class page    | 0     | 0      | 0     |                                                                                               | 9acd0a78c |
| shaman-elemental     | 3     | 0      | 3     |                                                                                               | 9acd0a78c |
| shaman-enhancement   | 1     | 3      | 4     | Air totem action (sim); Stormstrike level gate (sim)                                          | 9acd0a78c |
| shaman-restoration   | 0     | 3      | 3     | NS into rank 2 Healing Wave (sim); Chain Heal party or raid (owner); low-rank tooltip (owner) | 9acd0a78c |
| rogue class page     | 0     | 0      | 0     |                                                                                               | 9acd0a78c |
| rogue-assassination  | 4     | 1      | 5     | Mutilate off-hand penalty (owner)                                                             | 9acd0a78c |
| rogue-combat         | 2     | 2      | 4     | Efficient Strikes on a Backstab raid build (sim)                                              | 9acd0a78c |
| rogue-subtlety       | 3     | 1      | 4     | Shank/Ghostly Strike/Eviscerate gating, Shank rank 1 (sim)                                    | 9acd0a78c |
| priest class page    | 1     | 1      | 2     |                                                                                               | 9acd0a78c |
| priest-discipline    | 2     | 1      | 3     | low-rank heal ranks (owner tooltip, sim)                                                      | 9acd0a78c |
| priest-holy          | 1     | 1      | 2     | low-rank heal ranks (owner tooltip, sim)                                                      | 9acd0a78c |
| priest-shadow        | 5     | 3      | 8     | leveling order and Spirit Tap 5/5 (sim); SW:P cutoff (sim)                                    | 9acd0a78c |
| mage class page      | 2     | 0      | 2     |                                                                                               | 9acd0a78c |
| mage-arcane          | 5     | 2      | 7     | Arcane Power timing (sim)                                                                     | 9acd0a78c |
| mage-fire            | 3     | 2      | 5     | Accelerant/Impact vs Burning Soul (sim); haste weight (sim)                                   | 9acd0a78c |
| mage-frost           | 2     | 2      | 4     | Accelerant point (sim)                                                                        | 9acd0a78c |
| warlock class page   | 1     | 0      | 1     |                                                                                               | 9acd0a78c |
| warlock-affliction   | 4     | 0      | 4     | Wrack after 70755, Agony vs Doom (sim)                                                        | 9acd0a78c |
| warlock-demonology   | 3     | 1      | 4     | Horde race (owner)                                                                            | 9acd0a78c |
| warlock-destruction  | 4     | 1      | 5     | Molten Skin/Destructive Reach points (sim); Curse of the Elements order (owner)               | 9acd0a78c |
| hunter class page    | 1     | 0      | 1     |                                                                                               | 9acd0a78c |
| hunter-beast-mastery | 2     | 1      | 3     | Focused Fire before Endurance Training (ladder-first, sim/owner)                              | 9acd0a78c |
| hunter-marksmanship  | 5     | 3      | 8     | Sniper Shot, Rapid Killing (sim)                                                              | 9acd0a78c |
| hunter-survival      | 4     | 1      | 5     | Deadly Aspects under Falcon (owner)                                                           | 9acd0a78c |
| druid class page     | 1     | 0      | 1     |                                                                                               | 9acd0a78c |
| druid-balance        | 1     | 0      | 1     | Rebirth in Moonkin Form (owner)                                                               | 9acd0a78c |
| druid-feral          | 5     | 3      | 8     | cat leveling build's 8 Nature's Focus/Subtlety points (sim, SIM-D1)                           | 9acd0a78c |
| druid-feral-bear     | 6     | 0      | 6     | Rend and Tear 2% in the engine (sim)                                                          | 9acd0a78c |
| druid-restoration    | 2     | 1      | 3     | Rejuvenation rank 1 tooltip (owner)                                                           | 9acd0a78c |
| paladin class page   | 0     | 0      | 0     |                                                                                               | 9acd0a78c |
| paladin-holy         | 2     | 2      | 4     | Oath of Light at band 30 (sim)                                                                | 9acd0a78c |
| paladin-protection   | 4     | 2      | 6     | Lay on Hands mana (owner)                                                                     | 9acd0a78c |
| paladin-retribution  | 5     | 1      | 6     | Dwarf vs Human (sim); Exorcism crit (owner)                                                   | 9acd0a78c |

Round 1 totals: 72 audit CONFIRMED by class (the per-spec sums above count multi-page findings once per page), about 40 rendered-review MUSTs, all fixed; two ad-review claims checked and rejected (Flame Shock 21 s, Fury opener). Verdict: NOT zero, round 2 needed.

Out of scope this round, for the site and sim lanes: stat-table Expertise cap footnote on melee specs (no Forever item carries expertise; engine weights still list it), hit-cap split footnotes, generated "raid build moves N points" summaries that drop talents (web/src/lib/guides/build-diff.ts), generated recommended-race line duplicating prose, the top Build card showing the leveling build, cooldown cards ignoring talent cuts (Sacred Duty, Rapid Killing), Discipline card "?" icon and duplicate rows, site-voice footers ("Numbers from client build ...", "Raid loot is not itemized ...", "enchants are not listed yet"), repo paths in source labels.

### Round 1, research pass (owner rule 2026-11-04: research on the web when the sim ties or cannot decide)

Every class's open owner questions were researched (Forever sources first, Classic/SoD only where Forever kept the mechanic) and resolved on strong-player consensus; sources are cited per class in `round1/<class>.md`, "Owner questions resolved by research". Resolved and applied to the prose: warrior 6/6, rogue 5/5, druid 5/5, mage 5/6, warlock 5/6, priest 5/7, hunter 7/8, shaman 5/6, paladin 5/6. Left for the owner (taste or untestable): Molten Skin and the Affliction Alliance race (warlock), Shadowfiend learnability and planner links (priest), Orc vs Troll and the Alliance race (hunter), Dwarf vs Human (Retribution), source labels (mage), Lightning Overload rows (needs a level 60 log). The final rendered review after the pass: all nine SHIP once each class's MUSTs were applied.

Biggest research answer: Forever replaced Era's level-20 downranking rule with a level-gap rule (full healing bonus until a rank is 17 levels behind you, then 5% a level), so the rank 1 and 2 heals the healer cards cast are wrong at 60; prose is fixed, the card ranks and the engine curve are the first sim-round items.

Hand-set leveling order: data/curated/ladder-order.json (sim/leveling/ladder_order.go), rows that reorder points only between sim checkpoints (10, 20, 25, 30, 35, 38, 40, 45, 50, 55, 60; TestLadderOrderMovesNoCheckpoint). First row: Beast Mastery takes Focused Fire at 15-16 ahead of Endurance Training (owner; Mobalytics and Icy Veins Forever cited). Researched orders that DO move a checkpoint (Frost Improved Frostbolt first and Ice Lance at 20, Arcane Blast at 20, Enhancement Stormstrike at 25, Shadow Mind Flay at 20) need a measured ladder-first row and go to the sim round; their prose follows the shipped order until then.

## Round 2 (client 1.60.1.70755, 2026-11-04)

Same loop on main ab42101f7, plus the owner rules added in round 1: web research when the sim cannot decide (Forever sources first), racial and build sims where they apply, nothing left to the owner as a taste call, and a client spell description beats an inferred flag. Reports: `round2/<class>.md`; sims: `round2/race-sims.md`. Every class's rendered pages got a wow-player SHIP. Round 2 fixes are on land/guides-r2 (commits 384717fd4, 647736df5, e547daf4b), landing only on the coordinator's go.

| Class   | Audit CONFIRMED | Rendered-review MUSTs | Fixed        | Decisions made (with sources or sims)                                                                                       |
| ------- | --------------- | --------------------- | ------------ | --------------------------------------------------------------------------------------------------------------------------- |
| warrior | 3               | 2                     | 5 + 2 likely | Arms and Fury: Night Elf and Undead, best free races beyond error (Skyborne leads by a hair but needs the paid Heroic Pack) |
| shaman  | 4               | 1                     | 5 + 4 likely | Lightning Overload is its own smaller bolt (client rows); no logs yet                                                       |
| rogue   | 5               | 0                     | 5            | Combat openers left unranked until measured                                                                                 |
| priest  | 2               | 1                     | 3 + 2 likely | Shadowfiend is not in Forever; planner links stay in one form                                                               |
| mage    | 6               | 0                     | 6 + 1 likely | Source labels in player words, no paths (_prose.test.ts filePaths rule)                                                     |
| warlock | 6               | 2                     | 8 + 4 likely | Human with Orc for all three specs; Destruction raid build Pandemic 3 + Malevolence 3 (+1.7%, adopted)                      |
| hunter  | 5               | 1                     | 6 + 1 likely | Night Elf with Orc (BM, SV) or Troll (MM); Survival Rapid Killing loses 1.9%                                                |
| druid   | 5               | 1                     | 6 + 5 likely | (none open)                                                                                                                 |
| paladin | 3               | 1                     | 4 + 1 likely | Retribution stays Human (raid tie, bare +1.2%)                                                                              |

Round 2 totals: 39 audit CONFIRMED plus 9 rendered-review MUSTs, all fixed. Verdict: NOT zero, round 3 needed.

Out of scope, handed on: hover tooltips (spelltips) still carry pre-hotfix text for rogue, hunter and likely warrior/shaman (data regen with the 70755 hotfix layer); BiS headers name the old default race; the Fury BiS entry's talent string is illegal on 70755; card actions for low-rank heals, Agony vs Doom, Curse of the Elements order, Soul Fire, Subtlety gating, Arcane Wrath (sim round).

## Round 3 (client 1.60.1.70755, 2026-11-04)

Same loop on main 7a2165cb0. New this round: race picks measured for every spec on the guide raid builds (round3/race-sims.md: DPS specs, Disc/Holy on the healer fight, Protection warrior on the tank fight), and paid races (Skyborne) never recommended. Every class got a wow-player SHIP on rendered pages. Fixes on land/guides-r3 (c2cece8b7, f085c45f9).

| Class   | Audit CONFIRMED | Likely fixed | Rendered-review MUSTs      | Race picks now                                                        |
| ------- | --------------- | ------------ | -------------------------- | --------------------------------------------------------------------- |
| warrior | 3               | 0            | 1 (Demoralizing Shout 196) | Arms/Fury Night Elf + Undead; Protection Night Elf + Tauren           |
| shaman  | 2               | 1            | 0                          | Dwarf + Orc (Elemental, Enhancement); Dwarf + Tauren (Restoration)    |
| rogue   | 5               | 2            | 0                          | Night Elf + Undead (Assassination, Subtlety); Gnome + Undead (Combat) |
| priest  | 3               | 1            | 0                          | Human + Undead (Shadow, Holy); Gnome + Undead (Discipline)            |
| mage    | 9               | 1            | 1 (Ice Block)              | Gnome + Undead (Arcane, Frost); Human + Undead (Fire)                 |
| warlock | 1               | 4            | 0                          | Human + Orc                                                           |
| hunter  | 1               | 3            | 0                          | Night Elf + Orc (BM, SV) / Troll (MM)                                 |
| druid   | 4               | 1            | 0                          | Night Elf + Tauren                                                    |
| paladin | 5               | 4            | 1 (Reverence wording)      | Human (Ret, Holy), Dwarf (Protection)                                 |

Round 3 totals: 33 audit CONFIRMED and 3 rendered-review MUSTs, all fixed. Verdict: NOT zero, round 4 needed. The count fell 72, then 39, then 33, and the round 3 errors are smaller (wording, single numbers, a round 2 edit that overstated).

Not adopted: the Restoration druid Improved Moonfire to Improved Regrowth swap (+0.5%, under the bar, and it breaks the Restoration ladder-first invariant). For the sim round: Starshards cooldown and a Shadow line; Eureka! masks (Incinerate in, periodic ticks out; re-run Gnome vs Human for Destruction, Assassination and Discipline); Stoneform rotation line for Dwarf tanks; Crusader opener for Retribution; Elune's Light outside Recklessness; Pounce behind-target check; Sniper Shot and Summon Hawk builds; Spellblasting potion on the Elemental pull; Destruction weapon enchant.

## Round 4 (client 1.60.1.70755, 2026-11-04)

On top of round 3 (main 85ba627ab). New this round, from the day-1 walkthrough: a "Levels 1 to 19" subsection opens every spec guide's Leveling section (first abilities and trainer levels, the first talent points as the shipped order takes them, gear before 20 with wands for casters, the class quests that matter; notes in round4/<class>-levels.md), the rest sits under "Levels 20 to 60", and every source link was swept. Every class got a wow-player SHIP on rendered pages. Fixes on land/guides-r4 (752c02e43, 77bd7c12b, a901bae80).

| Class   | Audit CONFIRMED | Rendered-review MUSTs |
| ------- | --------------- | --------------------- |
| warrior | 4               | 3 (source entries)    |
| shaman  | 2               | 0                     |
| rogue   | 2               | 1                     |
| priest  | 3               | 2                     |
| mage    | 3               | 4 (2 source entries)  |
| warlock | 5               | 1 (source entry)      |
| hunter  | 2               | 0                     |
| druid   | 2               | 2 (source entries)    |
| paladin | 2               | 0                     |

Round 4 totals: 25 audit CONFIRMED and 13 rendered-review MUSTs, all fixed. Verdict: NOT zero, round 5 needed. Trend 72, 39, 33, 25; round 4's errors are mostly source labels and wording, with one level-scaled value (Demoralizing Shout 204 at 60) reverted.

Source-link rule now in force: no guide links /sources; stat weights link /sim/weights on damage specs only; boss-drop entries read "which boss drops are upgrades" (/sim/drops) on damage specs, and healer and tank pages link their own best in slot page; no "how <spec> plays" entries (the specs page is the simulator's coverage page); talent-tree sources open the planner; Talents Forever is labelled as beta client data; wago.tools links carry the build. Race picks re-measured this round (round4/race-sims.md) all held.

For the sim round, new: haste shortens the global cooldown in the engine (Classic does not), Summon Hawk unmodeled, no Sniper Shot line, Feral Charge (Cat) unmodeled, non-stat pet buffs blocked.

## Round 5 (client 1.60.1.70755, 2026-11-04)

On top of round 4 (main 42f10dda9). Every class got a wow-player SHIP on rendered pages. A second fix pass was needed after the review: the nightly best-in-slot re-rank (de044125c) moved the stat tables and BiS rows after the audit, so prose that named a stat order or gear pieces contradicted the table on the same page. That prose now gives the table's current order with stable reasons, and it names gear only for reasons that do not move nightly (bind on pickup, quests, rules like "a staff usually wins").

Layout: the Leveling section now reads Levels 1 to 19, then Levels 20 to 60 with the talent strip under it, so the character home's #leveling link opens on the 1 to 19 text. Below lg, headings jumped to clear the sticky jump bar (scroll-margin 60 px). ux-designer and wow-player SHIP at 1440 and 390.

| Class   | Audit CONFIRMED | Rendered-review MUSTs                                                      |
| ------- | --------------- | -------------------------------------------------------------------------- |
| warrior | 3               | 0                                                                          |
| shaman  | 1               | 3 (BiS and stat prose against re-ranked tables)                            |
| rogue   | 1               | 3 (Subtlety weapons, hit)                                                  |
| priest  | 0               | 3 page-text (stat prose, Mind Blast note), 2 data (icon, BiS race caption) |
| mage    | 3               | 0                                                                          |
| warlock | 1               | 4 (stat prose, raid-build lines)                                           |
| hunter  | 0               | 2 (hit prose, Improved Stings row)                                         |
| druid   | 2               | 2 (Furor wording, Defense line)                                            |
| paladin | 1               | 3 (stat prose; a Dwarf MUST was withdrawn on sims)                         |

Round 5 totals: 12 audit CONFIRMED and 20 page-text review MUSTs, all fixed. Most of the MUSTs come from the nightly re-rank, not from guide errors. Verdict: NOT zero, so round 6 is needed. Trend: 72, 39, 33, 25, 12.

Race decisions (round5/race-sims.md): Restoration shaman goes to [dwarf, orc]. Blood Fury's spell power effect (aura 317, misc 0) covers heals, and Orc leads Horde healing on both seeds. Retribution Human holds on the re-ranked gear (Dwarf even in a raid, Human ahead solo).

For the data lane: Alliance BiS rows show the Horde source Frostwolf Clan for items whose Alliance source is Stormpike Guard; the "N% built in" hit baseline includes BiS gear hit; the BiS race captions for Shadow and Discipline disagree with recommendedRaces; Lesser Heal rank 1 (2050) has no icon; and the spelltips from the hotfix layer (f83b247a1) dropped icons, so the guide-tips vitest fails on main.
For the sim round: Subtlety Hemorrhage dagger check is main-hand only in the engine (SIM-R5-1), Stormstrike trainer level 40 against the talent at 32, healer haste weights now large and positive, Demonology Soul Fire card against the prose.

## Round 6 (client 1.60.1.70755, 2026-11-04)

Built on round 5 (main 2b869f01d). Audit CONFIRMED by class: warrior 4, druid 3, rogue 3, warlock 2, hunter 2, paladin 2, priest 1, mage 1, shaman 1, for 21 in total. Trend: 72, 39, 33, 25, 12, 21. The rise is almost all stat-priority and gear prose that the nightly re-rank made false, on pages round 5's review pass had not reached. The rest are small:

- Edgemaster's Handguards on Arms do nothing for a two-hander.
- The goggles need Engineering to wear.
- The "below" and "above" references were wrong.
- Corruption was called cheap.
- Master Demonologist was credited before 35.
- Libram: Holy Light IX has no known source in Forever yet.
- A few leveling readings were wrong under the new layout.

New rule (owner-approved; it follows "prose never quotes nightly numbers" and "one job per page"): Stat priority prose no longer ranks stats. The table above it renders the order from the nightly scale factors. The prose gives each stat's lasting reason and the cap rule. _prose.test.ts now enforces this inside Stat priority sections only (STAT_RANKING: first, next, last, then, leads, ahead of, worth more than, and "second" as a rank but not as a duration). Gear prose names a piece only for a reason that does not move nightly. Every guide touched this round carries updated: 2026-11-04.

Waiting for the final best-in-slot re-rank: race re-runs on the re-ranked gear (mage, hunter), and a check of where Retribution Dwarf's raid edge comes from.

Rendered review, round 6: every class has a wow-player SHIP. The page-text review MUSTs, all fixed:

- warrior 1: "Stamina and armor first".
- shaman 1: Elemental "first stat".
- rogue 1: Combat "no bleeds".
- mage 2: Fire "nothing but Fire" and the Frost Skyborne line. A third MUST, the numbered stat pills, went to the component lane: they render the table's own data order.
- priest 1: "From 20".
  Verdict: NOT zero, so round 7 is needed.

## Round 7 (client 1.60.1.70755, 2026-11-04)

Built on round 6 (main 394562f20). Audit CONFIRMED by class: mage 4, warrior 3, paladin 2, druid 2, warlock 1, shaman 1, rogue 1, priest 0, hunter 0. Total 14. Trend: 72, 39, 33, 25, 12, 21, 14. Errors fixed:

- Impale makes a crit 220% of a normal hit.
- Intellect also gives spell crit to mages and warlocks.
- Seal of Righteousness scales with weapon speed and spell damage, not attack power.
- Rend and Tear covers melee abilities only.
- Shifting Power costs 30% less with Natural Shapeshifter.
- Ambush needs a main-hand dagger, so Subtlety's sword set opens with Cheap Shot or Garrote.
- Improved Bloodrage is now in the Arms and Fury Bloodrage notes.
- Fire power on Demonology covers Rain of Fire, not Soul Fire.
- Shaman BiS prose keeps only bind facts.
- Two round 6 rewrites were wrong and are fixed.

Rendered review: every class got a wow-player SHIP. The page-text MUSTs, all fixed:

- shaman 1 (Restoration Professions)
- warlock 2 (Destruction setup line, weapon oil)
- rogue 1 (Eureka! covers only non-periodic damaging abilities)
- mage 2 (Arcane Blast's Intellect, the "cheapest" claims)

Class index pages now carry the date of their last content change. Verdict: NOT zero, round 8 needed.

Open sims, none of them claimed in prose:

- Retribution Human against Dwarf on a Beast target. No race sim uses one, so Big Game Hunter never counts.
- Fire Scorch ranks.
- Arcane Frostbolt ranks.
- Protection threat by ability.
- Elemental Earth Shock filler.
- Warrior Dwarf against Human breakdown.
- Race re-runs after the final best-in-slot re-rank.

## Round 8 (client 1.60.1.70755, 2026-11-04)

Built on round 7 (main 09b98d9e4). Audit CONFIRMED by class: hunter 3, warrior 2, warlock 1, druid 1, and 0 for mage, paladin, priest, rogue and shaman. Total 7. Trend 72, 39, 33, 25, 12, 21, 14, 7. Five classes reached zero confirmed this round.

The hunter errors repeat the mistake fixed for Impale in round 7: Mortal Shots and Predator's Edge raise the crit bonus, so a crit does 230% of a normal hit, and 212% or 224% on the Beast Mastery build. The other fixes:

- Eureka! covers non-periodic damaging abilities only (warrior, warlock).
- Hurricane is a Nature spell.
- Thunder Clap's slow is not the biggest cut to damage taken; Demoralizing Shout cuts more.
- Battle Shout reaches party and raid members within 20 yards for 3 minutes.
- Cheap Shot trains at 26.
- Frostfire Bolt slows from 40.
- Benediction does not discount Hammer of Wrath.
- The Vindication row reason was wrong.

New sim finding: the engine's Eureka! client lists cover Rain of Fire and Hellfire, and give Drain Life and Drain Soul the boost, while the tooltip says "non-periodic". The guide follows the tooltip.

Verdict: NOT zero, so round 9 is needed.
Rendered review: wow-player SHIP for all eight changed classes, with no MUSTs. Fixes are on land/guides-r8.

## Round 9 (client 1.60.1.70755, 2026-11-04)

Built on round 8 (main 9b0ff4666). Audit CONFIRMED by class: priest 3, warrior 1, mage 1. Every other class has 0 (hunter, rogue, shaman, warlock, paladin, druid). Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5.

All five errors are about how far Eureka! reaches and how crits are worded:

- Priest Eureka! covered Mind Flay and every heal; Renew is now excluded.
- Warrior and mage lines dropped "damaging" or said "spells".
- "Shadowform doubles every crit" now says it doubles the crit bonus.
- The Holy paladin Benediction line now names what it discounts, so it holds whichever reading of the talent text is right.

Rule now applied everywhere: Eureka! follows the tooltip, covering non-periodic damaging or healing abilities only. That leaves out DoTs, channels and HoTs.

Big engine finding for the sim lane (druid S13): the pinned engine carries 70753 values for Shred, Claw, Ravage, Rip, Rake, Ferocious Bite and Starfire. The 70755 values are higher: Shred 200% against 155%, Starfire rank 6 471 against 337. Every feral and balance weight, best in slot and race sim is affected until the engine regenerates. The Decimation damage effect is 0 in the client, against 6% in the talent text; a log is needed to settle it.

Expertise was restored on the three tank pages in round 8 (bb4d616e0). The guide-cards gate now hides its table row until a band's BiS carries it; the prose holds either way.

Verdict: NOT zero, so round 10 is needed.
Rendered review: wow-player SHIP for priest, warrior, mage and paladin, with no MUSTs.

## Round 10 (client 1.60.1.70755, 2026-11-04)

Built on main 84460cec8 (sim round 15, engine pin 9f7f1f6e4), with r15c and rotation-gates merged in. Audit CONFIRMED by class: druid 11, priest 11, warlock 10, mage 8, warrior 7, rogue 6, hunter 6, shaman 3, paladin 3, for 65 in total. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65.

Almost every finding came from one cause. Sim round 15 changed builds, card ranks and rotations but updated only the frontmatter build lines, so prose and card notes still described the old builds. The new builds the prose now describes:

- Arcane raid is 32/16/3 with a Fireball filler.
- Affliction runs a Wrack raid build.
- Arms takes Spearing Strike, Reckless Swings and Improved Execute.
- Fury takes Improved Berserker Rage.
- Marksmanship takes Sniper Shot.
- Shadow levels with Wand Specialization, then Spirit Tap.
- The bear takes 1 point of Furor and puts the rest into Naturalist.
- Restoration druid ranks changed.
- Subtlety from 40 builds every finisher to five points for Death Mark.
- Retribution opens with Seal of the Crusader.

Prose and card notes now match the shipped builds, point orders and cards. The new APL lines have notes.

Rendered review: wow-player SHIP for every class. The page-text MUSTs were:

- paladin race tie
- Marksmanship leveling line
- Subtlety Death Mark
- Rupture with Blood Splatter only
- Shadow pre-cast timing
- Holy Greater Heal step-down
- Reckless Swings' 8 seconds and health cost
- Chain Heal's party-only jumps
- Arcane Impact and Heating Up wording

Handed to the component lane: the Arcane, Fire, Feral, Balance and Marksmanship headers load the leveling `build` next to raid-oriented cards, the key-talent chips show leveling ranks, and the Arcane opener timing.

Process note: a sim round that changes a build: line should ship the matching prose or flag the guide lane.

Verdict: NOT zero, so round 11 is needed.

Open sims: race re-runs on the new builds for most specs, Rip against Bite, Bite pooling, Crusader opener on short mobs, Consecration on packs, and Wrack's mask against its tooltip.

## Round 11 (client 1.60.1.70755, 2026-11-04)

Built on main d54de48aa (round 10 landed, engine pin 9f7f1f6e4). Audit CONFIRMED by class: rogue 2, druid 1, shaman 1, mage 1, paladin 1, priest 1. Warrior, warlock and hunter had 0. Total 7. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7.

Almost every finding was a slip in a round 10 rewrite:

- Natural Shapeshifter and Feral Instinct reasons
- Chain Heal's party-only jumps, in the card note and the grouping line
- Arcane Meditation at 2 ranks keeps most regeneration running, not all
- Consecration: the newest rank on packs from 40, rank 1 for single targets
- The Shadow opener puts Shadow Word: Death before the second Mind Blast
- Venom adds 20 points of poison chance, which lands Instant Poison on half your hits
- Subtlety keeps Rupture ticking and marks at five while it ticks

Verdict: NOT zero, round 12 needed.
Rendered review: wow-player SHIP for druid, shaman, mage, paladin, priest and rogue. Two page-text MUSTs are fixed: the Enhancement summary now says imbue, not Windfury, and Venom's 20 percentage points now has its unit.

## Round 12 (client 1.60.1.70755, 2026-11-04)

Built on main a8b147a6f (eleven guide rounds, guide-header-build, engine pin 9f7f1f6e4).

Audit CONFIRMED by class: druid 2, mage 2, paladin 1. Warrior, warlock, shaman, hunter, priest and rogue had 0 confirmed. Rogue had 1 likely finding, also fixed. Total 5. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5.

All five are refinements of rewrites from rounds 10 and 11:

- Consecration ranks, once spell power is counted
- Feral Instinct is the Brutal Impact row filler
- the Restoration tank Rejuvenation note is "from 28 to 49"
- the Arcane leveling loop has no Arcane Wrath or Pyroblast
- Frost Channeling only arrives at 58

Production on a8b147a6f shows rounds 10 and 11. Arcane and Feral prose agree with the new raid-build header and rank chips.

Verdict: NOT zero, so round 13 is needed. Six of nine classes are at zero.

Open: DRUID-R12-3 (Feral raid build spends 2 points on Feral Instinct; sim it against Furor 2 or Thick Hide 2), race re-runs on the new builds, and the BiS re-rank on the shipped raid builds.
Rendered review: wow-player SHIP for paladin, druid, mage, rogue and hunter. No MUSTs; two SHOULDs applied.

## Round 13 (client 1.60.1.70755, 2026-11-04)

Built on main 604b78a80 (twelve guide rounds, engine pin 9f7f1f6e4).

Audit CONFIRMED by class:

- paladin 1
- warrior, warlock, shaman, hunter, priest, rogue, mage and druid 0

That is 1 in total. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1. Eight of nine classes are at zero, and warrior, warlock and hunter have been at zero for three rounds running.

The one error: a round 8 wording said Hammer of Wrath "still costs full price" after Benediction. Holy Conduit 2/2 takes 40% off it, as the same page says. The line now says Benediction does not discount it.

Verdict: NOT zero, so round 14 is needed.

Still open, all with the sim lane and none claimed in prose: the BiS re-rank on the shipped raid builds, race re-runs on the new builds, and the Feral Instinct points against Furor or Thick Hide.
Rendered review: wow-player SHIP for paladin after checking Hammer of Wrath's family mask against Benediction's class mask.

## Round 14 (client 1.60.1.70755, 2026-11-04, launch state)

Built on main d7fc51fe7, engine pin f637ec32c. These are the live launch BiS files (de044125c picks), because the final re-rank failed. From this round, prose is held to the live tables and is no longer treated as provisional.

Audit CONFIRMED by class: shaman 4, druid 3, hunter 3, warrior 3, mage 2, rogue 2, priest 1, warlock 0, paladin 0. Paladin had 1 likely finding, which is fixed. Total 18. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1, 18.

Most findings are Professions sections that did not name the bind-on-pickup crafts in their own BiS sets:

- Heart of the Mountain needs Enchanting.
- Legionite Glaive needs Blacksmithing.
- Robe of the Archmage needs Tailoring.
- The goggles need Engineering 270.

Other fixes:

- The Arms Alliance race order: Dwarf measured first; Night Elf stays the pick.
- The Enhancement Hearts are now worded by faction.
- Shaman rank ladders give their starting level.
- The Restoration druid tank Rejuvenation note matches the round 16 rank rows.
- Two round 16 rogue fixes that a sim-branch merge undid are back in.

The Arcane "before you have the Fire points" lines are left to sim round 17.

For the data lane:

- No single character can wear the whole Fire or Frost BiS set.
- Frozen Heart is listed as BiS for Arcane, Fire and Retribution, where its on-use adds nothing.
- Arcane, Shadow and Combat BiS were ranked on old talent strings.

Verdict: NOT zero, so round 15 is needed.
Rendered review: wow-player SHIP for all eight changed classes after one druid MUST (Feral leather binding) was fixed.

## Round 15 (client 1.60.1.70755, 2026-11-04, launch state)

Built on main 4600b7d82 (round 14 plus sim round 17), engine pin b2c4a2d05. The audits were relaunched after an API rate limit and run five at a time on the shared quota.

Audit CONFIRMED by class: rogue 2, priest 1, shaman 1, warlock 1, hunter 1, druid 1, paladin 1, warrior 0, mage 0. Total 8. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1, 18, 8.

Almost every finding is new text from sim round 17 or a round 14 fix:

- Aimed Shot rank 1: the reason is now damage per mana across a chain of pulls, not "almost as hard" or "a quarter of the mana".
- Demon raid buffs: protective party auras do reach the demon.
- Backstab card notes no longer credit Opportunity or Puncturing Wounds before the build takes them.
- Shadow healing-gear rule: Spirit is now in it.
- Enhancement hit: the two caps (9% melee, 16% spells) are named.
- Bear Professions: no longer points at the cat line.
- Protection paladin Elixir of Holy Power: it does nothing for Hammer of the Righteous.
- Reckoning: the reason it is left out is now the measured one.

Open, with the sim lane:

- Marksmanship Aimed Shot ranks at 53 to 59.
- Survival Lone Wolf against the pet at 10,000 iterations and two seeds.
- Affliction races with Heart of the Lion on the demon.
- Arcane "from 22" Missiles.
- Reckoning against Conviction plus Deflection.
- Seal of Righteousness rank 8 has a 0.0 spell-power coefficient in the client.

Open, with the data lane:

- Subtlety level 20 mace.
- Frost level 60 BiS on the old build.
- Enhancement hit weight.

Verdict: NOT zero, so round 16 is needed.
Rendered review: wow-player SHIP for all seven changed classes after four page-text MUSTs were fixed (Combat dagger from 20, demon auras by name, bear professions scoped to the raid set, Reckoning trade).

## Round 16 (client 1.60.1.70755, 2026-11-04, launch state)

Built on land/guides-r15 (2c3135a2e), engine pin b2c4a2d05, five auditors at a time.

Audit CONFIRMED by class: rogue 2, warlock 1, hunter 1 (plus 3 likely). The other six classes were 0. Total 4. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1, 18, 8, 4.

Fixes:

- **Rogue:** the Subtlety Ghostly Strike and Garrote card notes are now true at every level they show.
- **Warlock:** Retribution Aura reaches the demon.
- **Hunter:** low-mana order now keeps Serpent Sting over Multi-Shot (better damage per mana). "Mail from 40" lines now say leather stays in the best gear.

Folded in from the racials audit:

- **Eureka! now follows each class's client row, not the tooltip.** The warlock and priest rows add 10% periodic damage to Wrack, the drains, Mind Flay and Penance. Prose that sits next to the "non-periodic" racial tooltip says so. Incinerate is not covered.
- **Attuned** is removed on entering a raid.
- **Arms Dwarf edge** is put down to Artisan's Precision; Night Elf stays the pick.

Verdict: NOT zero, so round 17 is needed. Six classes audit at zero this round.
Rendered review: wow-player SHIP for all seven changed classes after three page-text MUSTs were fixed (Incinerate under Eureka!, tooltip wording beside the periodic claims, the Arms Heroic Strike sentence).

## Round 17 (client 1.60.1.70755, 2026-11-04, launch state)

Built on main 919f7e8a7 (guide round 16 plus legal copy), engine pin 6ea675594, five auditors at a time, the live client 1.60.1.70790 identical in talents, items, racials and numbers.

Audit CONFIRMED on main by class: priest 8, paladin 2, rogue 1, shaman 1, warrior 1, warlock 1; mage, hunter and druid 0. Total 14. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1, 18, 8, 4, 14. Most of the priest and paladin findings are sim round 18's own prose and card notes (rank reasons and card notes that were not true at every level they render on).

Fixes:

- **Priest:** Greater Heal and Renew rank reasons on Holy and Discipline (rank 1 keeps the whole bonus to 57; the bigger ranks spend mana a dungeon tank rarely needs; a small heal wastes less on a small hole); Devouring Plague rank 1 from 30; Human is level with Gnome on Discipline and the index, either race a fine pick (Gnome stays, the 0.8% lead is under the bar); Professions names Truefaith Vestments and the Dormant Heart of the Mountain; the Heal, Greater Heal filler and party Greater Heal card notes are true at every level.
- **Paladin:** the Retribution opener on the page follows its own card (Righteousness rank 7 after the Crusader Judgement, then Command on the first landed swing in a raid); Consecration is the most expensive button in the priority, not on the bars.
- **Rogue:** the Subtlety Rupture note credits Serrated Blades from 21 and Hemorrhage from 30, not at 20.
- **Shaman:** Enhancement's raid mana line no longer says "topped up"; you never run dry, so shock down to the last quarter.
- **Warrior:** Flurry's three charges go to white swings; a Heroic Strike comes out at the faster speed without using one.
- **Warlock:** Undead sits close behind Orc for Affliction, as on the other two pages (measured 0.5%).

Sim round 19 landed (04d7da663) while this round ran; its guide text carried five errors of its own, fixed in this landing: Destruction's Corruption/Immolate rank levels at 40 to 42; the priest "Keep Renew rank 7 or higher" note (false at 40 to 59), the "top trainer rank" filler note (round 17 fixed main's version: rank 1 until 57) and "rank 8 heals barely more" (38% more base); Restoration's "Nature's Swiftness makes the cast free" (instant, not free). Round 19 also rewrote lines this round touched (fury.md:64, holy.md:122, discipline.md:124, retribution.md:95, enhancement.md:87): the merge keeps round 17's text and round 19's own new clauses (Lesser Heal, rank 3 and 4 at 58, the Maelstrom half-mana rule, Consecration's rank levels).

Sims routed to the sim lane: Retribution fresh-boss opener (Righteousness first against Command first on the raid build); Devouring Plague rank 1 against 2 at 28 and 29; Frost races on the shipped build with Ice Lance at 0.107; Enhancement's 25% gate against no gate at raid 60; Gnome against Human for Discipline at the re-ranked BiS.

Verdict: NOT zero, so round 18 is needed. Three classes audit at zero this round.
Rendered review: wow-player SHIP for all six changed classes (rogue, shaman, warrior, warlock, priest, paladin), every tile of every page read.

## Round 18 (client 1.60.1.70755, 2026-11-04, launch state)

Built on main 588a3ce2b (guide round 17 plus sim round 19), engine pin da54b7b06, five auditors at a time, the live client 1.60.1.70790 identical in talents, items, racials and numbers.

Audit CONFIRMED by class: warrior 6, hunter 6, priest 4, warlock 4 (plus 2 likely, taken), shaman 3, mage 3, paladin 1, rogue 1, druid 0. Total 28. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1, 18, 8, 4, 14, 28. Page prose was clean on seven of nine classes; the rise is the every-level rule applied to every card note's later sentences (a note that credits a talent, pet ability or trained button before the shipped order gives it is false on the bands it renders on), which round 17 applied only to the priest and rogue notes.

Fixes:

- **Card notes, six classes (warrior, hunter, warlock, mage, priest, rogue):** every note now credits Bloodthrill, Mortal Strike, Recklessness, Execute, Whirlwind, Resourcefulness, Expose Prey, Clever Traps, Multi-Shot, Shadow and Flame, Bane, Improved Corruption, Amplify Curse, the Imp sacrifice, Heating Up, Fingers of Frost, Missile Barrage, Renewed Hope, Divine Aegis and Improved Mind Flay only from the level the player has them; Wing Clip no longer quotes the level 60 rank's damage.
- **Priest:** the Holy Greater Heal filler note is true at 59 (rank 3 or the top rank); Shadow's Devouring Plague line follows the Human rows at 50 and 51.
- **Shaman:** the Enhancement half-mana reserve runs from the first point of Maelstrom Weapon to the fifth, a quarter before the talent and at five points; the five-stack bolt is instant and free only with all five points (page and notes).
- **Paladin:** Consecration keeps its measured rank levels without the per-cast cost reason the client rows contradict; the Seal of Righteousness and rank 1 sentences read plainly.
- **Rogue:** Subtlety's weakness is its bleed dependence, not a damage ranking the current engine reverses.

Sims routed to the sim lane: Consecration ranks 1 to 4 at 50 to 53 without the pacing line; Fire Blast before Heating Up at 20 and 28; Arcane Missiles-first at 12 and 16; Enhancement half against a quarter at 37 to 41; Subtlety's Eviscerate fallback below 20 (SIM-R11-1); the cross-tree rogue order after the re-rank; Fury Alliance Skyborne against Night Elf. Component lane: the addon rows carry no talent gate (Arcane Fire Blast renders on the leveling build); the Survival cooldown card shows the trap's base cooldown beside a talented Deterrence.

Verdict: NOT zero, so round 19 is needed. Druid audits at zero; seven classes have clean page prose.
Rendered review: wow-player SHIP for all eight changed classes, every tile of every page read.

## Round 19 (client 1.60.1.70790, 2026-11-05, launch state)

Built on main 6beb919c2 (guide round 18 and the 1.60.1.70790 activation), engine pin abb4e6b8e, three auditors at a time, audited against data/builds/1.60.1.70790 (talents, items, racials and numbers identical to 70755 apart from two Sulfuras rings no guide names).

Audit CONFIRMED by class: paladin 11, mage 8, druid 7, hunter 3, priest 3, shaman 3 (plus 1 likely, taken), warrior 2, rogue 2, warlock 1 (plus 1 likely, taken). Total 40. Trend: 72, 39, 33, 25, 12, 21, 14, 7, 5, 65, 7, 5, 1, 18, 8, 4, 14, 28, 40. Every confirmed item but two (priest holy.md, the likely shaman restoration.md line) is a card note. The rise is round 18's every-level rule carried into each note's later sentences on the classes round 18 read only by first sentence (paladin, druid, mage), plus level 60 numbers quoted on bands where a lower rank renders (Mana Spring, Holy Strike, Divine Shield, Seal mana, Ferocity, Improved Fire Blast, Fingers of Frost, Missile Barrage).

Fixes:

- **Card notes, every class:** each note now credits Improved Berserker Rage and Berserker Stance (36), Improved Bloodrage's two ranks, Bestial Wrath, Rapid Killing, Aimed Shot, Eclipse (and its one-point value), Innervate, Starfire, Rend and Tear, Ferocity, Lacerate, Primal Bite, Bloom, Heating Up, Improved Scorch, Missile Barrage's ranks on both orders, Presence of Mind targets, Fingers of Frost, Improved Fire Blast, the mana gem and Major Mana Potion, Mutilate, Seal Fate, Boot Knife, Holy Shield, Divine Favor, Oath of Light, Sanctified Judgement, Shieldsworn, Divine Shield, Hammer of the Righteous, Vampiric Embrace, higher Greater Heal ranks, Magma Totem, Stormstrike, Conflagrate and Incinerate only from the level the player has them, and states values by rank where the rows cast lower ranks.
- **Priest:** Holy's leveling line teaches the Renew downranking the card casts instead of a Heal rank 2 line the card contradicts.
- **Shaman:** Restoration keeps the top rank for Nature's Swiftness from 44.
- **Warlock:** the four client-data source links name build 70790.
- addon-data.json and Data.lua regenerated for 1.60.1.70790 (note text only).

Sims routed to the sim lane (round 20 integrator): Fire Scorch against Fireball as the filler at 22 to 30 (decides fire.md:116, MAGE-R19-9); Destruction Conflagrate on cooldown against holding until Immolate has 4 seconds or less at 34 to 43 (decides destruction.md:128 or the card, WARLOCK-R19-3); Balance leveling build's Eclipse Wrath on and off at 36, 40 and 50; Affliction Amplify Curse saved for the last minute against before each Bane of Agony below 60; Subtlety Shank gate without Boot Knife at 24 to 32 (optional); Holy paladin emergency Holy Light rank at 47 to 59 (optional); Holy priest Heal rank 2 against 3 at 28 to 33 (optional); Consecration ranks 1 to 4 at 50 to 53 (carried). Sim findings: the Balance card has nothing to cast at 10 to 16 when mana is ahead and Moonfire is ticking; the Elemental card drops no Fire totem below 26 on two or more targets; the engine does not track Soul Shards (Destruction Shadowburn at 30 to 43).

Verdict: NOT zero, so round 20 is needed. Page prose was clean on seven of nine classes.
Rendered review: wow-player SHIP for all nine changed classes after three musts (Multi-Shot leads with the raid rule; Subtlety Shank's Boot Knife wording; the Protection Holy Shield notes kept whole in the page clause, which cuts at a colon or semicolon). Every changed line was read rendered; hunter and druid tiles were all read, and on the other classes the image limit left some unchanged tiles read as live page text instead.

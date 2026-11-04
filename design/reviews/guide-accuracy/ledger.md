# Guide-accuracy ledger

Rounds run until one finds zero confirmed errors. Each round: one expert auditor per class on the active client's data (hard evidence only), class fix lanes, a wow-player SHIP per class on rendered pages, one landing. Reports: `round<N>/<class>.md` (audit, fix log, fix passes).

Columns: **Audit** = CONFIRMED findings in the audit (a finding that touches several pages counts on each); **Review** = MUSTs the rendered-page reviews found and the fix passes fixed; **Fixed** = Audit + Review fixed this round (plus LIKELY findings the lanes proved with data, in the class report); **Left** = what waits on the owner or the next sim round; **sha** = the commit carrying the round's fixes (land/guides-r1 9acd0a78c, landed on main with the branch).

## Round 1 (client 1.60.1.70755, 2026-11-04)

| Spec | Audit | Review | Fixed | Left | sha |
|---|---|---|---|---|---|
| warrior class page | 0 | 0 | 0 | planner links (owner, kept by R1) | 9acd0a78c |
| warrior-arms | 1 | 0 | 1 | Arms opener Bloodrage vs Charge (owner) | 9acd0a78c |
| warrior-fury | 2 | 1 | 3 | Improved Berserker Rage (sim) | 9acd0a78c |
| warrior-protection | 1 | 2 | 3 | Improved Shield Wall (owner); Expertise cap footnote (component) | 9acd0a78c |
| shaman class page | 0 | 0 | 0 | | 9acd0a78c |
| shaman-elemental | 3 | 0 | 3 | | 9acd0a78c |
| shaman-enhancement | 1 | 3 | 4 | Air totem action (sim); Stormstrike level gate (sim) | 9acd0a78c |
| shaman-restoration | 0 | 3 | 3 | NS into rank 2 Healing Wave (sim); Chain Heal party or raid (owner); low-rank tooltip (owner) | 9acd0a78c |
| rogue class page | 0 | 0 | 0 | | 9acd0a78c |
| rogue-assassination | 4 | 1 | 5 | Mutilate off-hand penalty (owner) | 9acd0a78c |
| rogue-combat | 2 | 2 | 4 | Efficient Strikes on a Backstab raid build (sim) | 9acd0a78c |
| rogue-subtlety | 3 | 1 | 4 | Shank/Ghostly Strike/Eviscerate gating, Shank rank 1 (sim) | 9acd0a78c |
| priest class page | 1 | 1 | 2 | | 9acd0a78c |
| priest-discipline | 2 | 1 | 3 | low-rank heal ranks (owner tooltip, sim) | 9acd0a78c |
| priest-holy | 1 | 1 | 2 | low-rank heal ranks (owner tooltip, sim) | 9acd0a78c |
| priest-shadow | 5 | 3 | 8 | leveling order and Spirit Tap 5/5 (sim); SW:P cutoff (sim) | 9acd0a78c |
| mage class page | 2 | 0 | 2 | | 9acd0a78c |
| mage-arcane | 5 | 2 | 7 | Arcane Power timing (sim) | 9acd0a78c |
| mage-fire | 3 | 2 | 5 | Accelerant/Impact vs Burning Soul (sim); haste weight (sim) | 9acd0a78c |
| mage-frost | 2 | 2 | 4 | Accelerant point (sim) | 9acd0a78c |
| warlock class page | 1 | 0 | 1 | | 9acd0a78c |
| warlock-affliction | 4 | 0 | 4 | Wrack after 70755, Agony vs Doom (sim) | 9acd0a78c |
| warlock-demonology | 3 | 1 | 4 | Horde race (owner) | 9acd0a78c |
| warlock-destruction | 4 | 1 | 5 | Molten Skin/Destructive Reach points (sim); Curse of the Elements order (owner) | 9acd0a78c |
| hunter class page | 1 | 0 | 1 | | 9acd0a78c |
| hunter-beast-mastery | 2 | 1 | 3 | Focused Fire before Endurance Training (ladder-first, sim/owner) | 9acd0a78c |
| hunter-marksmanship | 5 | 3 | 8 | Sniper Shot, Rapid Killing (sim) | 9acd0a78c |
| hunter-survival | 4 | 1 | 5 | Deadly Aspects under Falcon (owner) | 9acd0a78c |
| druid class page | 1 | 0 | 1 | | 9acd0a78c |
| druid-balance | 1 | 0 | 1 | Rebirth in Moonkin Form (owner) | 9acd0a78c |
| druid-feral | 5 | 3 | 8 | cat leveling build's 8 Nature's Focus/Subtlety points (sim, SIM-D1) | 9acd0a78c |
| druid-feral-bear | 6 | 0 | 6 | Rend and Tear 2% in the engine (sim) | 9acd0a78c |
| druid-restoration | 2 | 1 | 3 | Rejuvenation rank 1 tooltip (owner) | 9acd0a78c |
| paladin class page | 0 | 0 | 0 | | 9acd0a78c |
| paladin-holy | 2 | 2 | 4 | Oath of Light at band 30 (sim) | 9acd0a78c |
| paladin-protection | 4 | 2 | 6 | Lay on Hands mana (owner) | 9acd0a78c |
| paladin-retribution | 5 | 1 | 6 | Dwarf vs Human (sim); Exorcism crit (owner) | 9acd0a78c |

Round 1 totals: 72 audit CONFIRMED by class (the per-spec sums above count multi-page findings once per page), about 40 rendered-review MUSTs, all fixed; two ad-review claims checked and rejected (Flame Shock 21 s, Fury opener). Verdict: NOT zero, round 2 needed.

Out of scope this round, for the site and sim lanes: stat-table Expertise cap footnote on melee specs (no Forever item carries expertise; engine weights still list it), hit-cap split footnotes, generated "raid build moves N points" summaries that drop talents (web/src/lib/guides/build-diff.ts), generated recommended-race line duplicating prose, the top Build card showing the leveling build, cooldown cards ignoring talent cuts (Sacred Duty, Rapid Killing), Discipline card "?" icon and duplicate rows, site-voice footers ("Numbers from client build ...", "Raid loot is not itemized ...", "enchants are not listed yet"), repo paths in source labels.

### Round 1, research pass (owner rule 2026-11-04: research on the web when the sim ties or cannot decide)

Every class's open owner questions were researched (Forever sources first, Classic/SoD only where Forever kept the mechanic) and resolved on strong-player consensus; sources are cited per class in `round1/<class>.md`, "Owner questions resolved by research". Resolved and applied to the prose: warrior 6/6, rogue 5/5, druid 5/5, mage 5/6, warlock 5/6, priest 5/7, hunter 7/8, shaman 5/6, paladin 5/6. Left for the owner (taste or untestable): Molten Skin and the Affliction Alliance race (warlock), Shadowfiend learnability and planner links (priest), Orc vs Troll and the Alliance race (hunter), Dwarf vs Human (Retribution), source labels (mage), Lightning Overload rows (needs a level 60 log). The final rendered review after the pass: all nine SHIP once each class's MUSTs were applied.

Biggest research answer: Forever replaced Era's level-20 downranking rule with a level-gap rule (full healing bonus until a rank is 17 levels behind you, then 5% a level), so the rank 1 and 2 heals the healer cards cast are wrong at 60; prose is fixed, the card ranks and the engine curve are the first sim-round items.

Hand-set leveling order: data/curated/ladder-order.json (sim/leveling/ladder_order.go), rows that reorder points only between sim checkpoints (10, 20, 25, 30, 35, 38, 40, 45, 50, 55, 60; TestLadderOrderMovesNoCheckpoint). First row: Beast Mastery takes Focused Fire at 15-16 ahead of Endurance Training (owner; Mobalytics and Icy Veins Forever cited). Researched orders that DO move a checkpoint (Frost Improved Frostbolt first and Ice Lance at 20, Arcane Blast at 20, Enhancement Stormstrike at 25, Shadow Mind Flay at 20) need a measured ladder-first row and go to the sim round; their prose follows the shipped order until then.

## Round 2 (client 1.60.1.70755, 2026-11-04)

Same loop on main ab42101f7, plus the owner rules added in round 1: web research when the sim cannot decide (Forever sources first), racial and build sims where they apply, nothing left to the owner as a taste call, and a client spell description beats an inferred flag. Reports: `round2/<class>.md`; sims: `round2/race-sims.md`. Every class's rendered pages got a wow-player SHIP. Round 2 fixes are on land/guides-r2 (commits 384717fd4, 647736df5, e547daf4b), landing only on the coordinator's go.

| Class | Audit CONFIRMED | Rendered-review MUSTs | Fixed | Decisions made (with sources or sims) |
|---|---|---|---|---|
| warrior | 3 | 2 | 5 + 2 likely | Arms and Fury: Night Elf and Undead, best free races beyond error (Skyborne leads by a hair but needs the paid Heroic Pack) |
| shaman | 4 | 1 | 5 + 4 likely | Lightning Overload is its own smaller bolt (client rows); no logs yet |
| rogue | 5 | 0 | 5 | Combat openers left unranked until measured |
| priest | 2 | 1 | 3 + 2 likely | Shadowfiend is not in Forever; planner links stay in one form |
| mage | 6 | 0 | 6 + 1 likely | Source labels in player words, no paths (_prose.test.ts filePaths rule) |
| warlock | 6 | 2 | 8 + 4 likely | Human with Orc for all three specs; Destruction raid build Pandemic 3 + Malevolence 3 (+1.7%, adopted) |
| hunter | 5 | 1 | 6 + 1 likely | Night Elf with Orc (BM, SV) or Troll (MM); Survival Rapid Killing loses 1.9% |
| druid | 5 | 1 | 6 + 5 likely | (none open) |
| paladin | 3 | 1 | 4 + 1 likely | Retribution stays Human (raid tie, bare +1.2%) |

Round 2 totals: 39 audit CONFIRMED plus 9 rendered-review MUSTs, all fixed. Verdict: NOT zero, round 3 needed.

Out of scope, handed on: hover tooltips (spelltips) still carry pre-hotfix text for rogue, hunter and likely warrior/shaman (data regen with the 70755 hotfix layer); BiS headers name the old default race; the Fury BiS entry's talent string is illegal on 70755; card actions for low-rank heals, Agony vs Doom, Curse of the Elements order, Soul Fire, Subtlety gating, Arcane Wrath (sim round).

## Round 3 (client 1.60.1.70755, 2026-11-04)

Same loop on main 7a2165cb0. New this round: race picks measured for every spec on the guide raid builds (round3/race-sims.md: DPS specs, Disc/Holy on the healer fight, Protection warrior on the tank fight), and paid races (Skyborne) never recommended. Every class got a wow-player SHIP on rendered pages. Fixes on land/guides-r3 (c2cece8b7, f085c45f9).

| Class | Audit CONFIRMED | Likely fixed | Rendered-review MUSTs | Race picks now |
|---|---|---|---|---|
| warrior | 3 | 0 | 1 (Demoralizing Shout 196) | Arms/Fury Night Elf + Undead; Protection Night Elf + Tauren |
| shaman | 2 | 1 | 0 | Dwarf + Orc (Elemental, Enhancement); Dwarf + Tauren (Restoration) |
| rogue | 5 | 2 | 0 | Night Elf + Undead (Assassination, Subtlety); Gnome + Undead (Combat) |
| priest | 3 | 1 | 0 | Human + Undead (Shadow, Holy); Gnome + Undead (Discipline) |
| mage | 9 | 1 | 1 (Ice Block) | Gnome + Undead (Arcane, Frost); Human + Undead (Fire) |
| warlock | 1 | 4 | 0 | Human + Orc |
| hunter | 1 | 3 | 0 | Night Elf + Orc (BM, SV) / Troll (MM) |
| druid | 4 | 1 | 0 | Night Elf + Tauren |
| paladin | 5 | 4 | 1 (Reverence wording) | Human (Ret, Holy), Dwarf (Protection) |

Round 3 totals: 33 audit CONFIRMED and 3 rendered-review MUSTs, all fixed. Verdict: NOT zero, round 4 needed. The count fell 72, then 39, then 33, and the round 3 errors are smaller (wording, single numbers, a round 2 edit that overstated).

Not adopted: the Restoration druid Improved Moonfire to Improved Regrowth swap (+0.5%, under the bar, and it breaks the Restoration ladder-first invariant). For the sim round: Starshards cooldown and a Shadow line; Eureka! masks (Incinerate in, periodic ticks out; re-run Gnome vs Human for Destruction, Assassination and Discipline); Stoneform rotation line for Dwarf tanks; Crusader opener for Retribution; Elune's Light outside Recklessness; Pounce behind-target check; Sniper Shot and Summon Hawk builds; Spellblasting potion on the Elemental pull; Destruction weapon enchant.

## Round 4 (client 1.60.1.70755, 2026-11-04)

On top of round 3 (main 85ba627ab). New this round, from the day-1 walkthrough: a "Levels 1 to 19" subsection opens every spec guide's Leveling section (first abilities and trainer levels, the first talent points as the shipped order takes them, gear before 20 with wands for casters, the class quests that matter; notes in round4/<class>-levels.md), the rest sits under "Levels 20 to 60", and every source link was swept. Every class got a wow-player SHIP on rendered pages. Fixes on land/guides-r4 (752c02e43, 77bd7c12b, a901bae80).

| Class | Audit CONFIRMED | Rendered-review MUSTs |
|---|---|---|
| warrior | 4 | 3 (source entries) |
| shaman | 2 | 0 |
| rogue | 2 | 1 |
| priest | 3 | 2 |
| mage | 3 | 4 (2 source entries) |
| warlock | 5 | 1 (source entry) |
| hunter | 2 | 0 |
| druid | 2 | 2 (source entries) |
| paladin | 2 | 0 |

Round 4 totals: 25 audit CONFIRMED and 13 rendered-review MUSTs, all fixed. Verdict: NOT zero, round 5 needed. Trend 72, 39, 33, 25; round 4's errors are mostly source labels and wording, with one level-scaled value (Demoralizing Shout 204 at 60) reverted.

Source-link rule now in force: no guide links /sources; stat weights link /sim/weights on damage specs only; boss-drop entries read "which boss drops are upgrades" (/sim/drops) on damage specs, and healer and tank pages link their own best in slot page; no "how <spec> plays" entries (the specs page is the simulator's coverage page); talent-tree sources open the planner; Talents Forever is labelled as beta client data; wago.tools links carry the build. Race picks re-measured this round (round4/race-sims.md) all held.

For the sim round, new: haste shortens the global cooldown in the engine (Classic does not), Summon Hawk unmodeled, no Sniper Shot line, Feral Charge (Cat) unmodeled, non-stat pet buffs blocked.

## Round 5 (client 1.60.1.70755, 2026-11-04)

On top of round 4 (main 42f10dda9). Every class got a wow-player SHIP on rendered pages. A second fix pass was needed after the review: the nightly best-in-slot re-rank (de044125c) moved the stat tables and BiS rows after the audit, so prose that named a stat order or gear pieces contradicted the table on the same page. That prose now gives the table's current order with stable reasons, and it names gear only for reasons that do not move nightly (bind on pickup, quests, rules like "a staff usually wins").

Layout: the Leveling section now reads Levels 1 to 19, then Levels 20 to 60 with the talent strip under it, so the character home's #leveling link opens on the 1 to 19 text. Below lg, headings jumped to clear the sticky jump bar (scroll-margin 60 px). ux-designer and wow-player SHIP at 1440 and 390.

| Class | Audit CONFIRMED | Rendered-review MUSTs |
|---|---|---|
| warrior | 3 | 0 |
| shaman | 1 | 3 (BiS and stat prose against re-ranked tables) |
| rogue | 1 | 3 (Subtlety weapons, hit) |
| priest | 0 | 3 page-text (stat prose, Mind Blast note), 2 data (icon, BiS race caption) |
| mage | 3 | 0 |
| warlock | 1 | 4 (stat prose, raid-build lines) |
| hunter | 0 | 2 (hit prose, Improved Stings row) |
| druid | 2 | 2 (Furor wording, Defense line) |
| paladin | 1 | 3 (stat prose; a Dwarf MUST was withdrawn on sims) |

Round 5 totals: 12 audit CONFIRMED and 20 page-text review MUSTs, all fixed. Most of the MUSTs come from the nightly re-rank, not from guide errors. Verdict: NOT zero, so round 6 is needed. Trend: 72, 39, 33, 25, 12.

Race decisions (round5/race-sims.md): Restoration shaman goes to [dwarf, orc]. Blood Fury's spell power effect (aura 317, misc 0) covers heals, and Orc leads Horde healing on both seeds. Retribution Human holds on the re-ranked gear (Dwarf even in a raid, Human ahead solo).

For the data lane: Alliance BiS rows show the Horde source Frostwolf Clan for items whose Alliance source is Stormpike Guard; the "N% built in" hit baseline includes BiS gear hit; the BiS race captions for Shadow and Discipline disagree with recommendedRaces; Lesser Heal rank 1 (2050) has no icon; and the spelltips from the hotfix layer (f83b247a1) dropped icons, so the guide-tips vitest fails on main.
For the sim round: Subtlety Hemorrhage dagger check is main-hand only in the engine (SIM-R5-1), Stormstrike trainer level 40 against the talent at 32, healer haste weights now large and positive, Demonology Soul Fire card against the prose.

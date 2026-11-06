# Guide accuracy round 26: Warrior

Audited on the round 26 tree (head ca06416cd: main after guide round 25, which merged sim round 22 at fd758eefc). Engine pin bff826207 (`sim/enginever/version.go`). Active client 1.60.1.70790.

The integration tree (`gr26/int`, HEAD 22783f43d, clean) carries two warrior changes made after the tree was cut. I audited both as the current state:

- `data/curated/coach/warrior-protection.json` and `coachrules.py`: the Demoralizing Shout aura is back, with `covered_by` = Demoralizing Roar 99, 1735, 9490, 9747, 9898 (every rank in `trainables/druid.json`). `f2_uptime.go` `withCover` counts that aura's time as the slot filled.
- `data/curated/coach/warrior-arms.md:57`: drill.ranks drops "Sometimes the right one is a lower, cheaper rank" (the rogue fix to `starter_kit.ranks_drill` for flat-cost resources). That is true for rage: every warrior rank row is strictly better than the one below.

Files audited:

- `web/src/content/guides/warrior/{index,arms,fury,protection}.md`.
- `data/curated/apl/warrior-{arms,fury,protection}.json`, including Fury's new `level_notes`.
- Every rendered band in `addon-data.json` `rotations/warrior-*`.
- The coach files `data/curated/coach/warrior-*.{md,json}`, the warrior entries of `coachrules.py`, and the coach code that judges them: `f1_core_casts.go`, `f2_uptime.go`, `f6_preparation.go`, `f7_death.go`, `copy.go`, `rules.go`.
- Against `addon-builds.json` (decoded below), `trainables/warrior.json`, `spelltips` and `spellconst/warrior.json`, `talents/warrior.json`, `sim/request/raid-consumes.json`, `items.json` and `catalogue.json`.

Read first: round25/warrior.md, the ledger's Round 25 section, `2026-11-05-sim-audit-round22/warrior.md` and integration.md, and the other round 26 class reports for the shared roots (HUNTER-R26-6, DRUID-R26-6, WARLOCK-R26-11).

## New work since round 25's tree (438f8ba15)

Sim round 22 (`git diff 438f8ba15 ca06416cd`):

- **`warrior-fury.json`.**
  - Six main notes were reworded: Bloodrage, Death Wish, Recklessness, Battle Shout, Execute, Reckless Swings and Berserker Rage, plus Heroic Strike.
  - Six `level_notes` were added at 60.
  - Every one was checked. One is wrong: the Heroic Strike note (R26-1).
- **`trainables/warrior.json`.** Talent spells now learn at their tier: Reckless Swings 25, Spearing Strike 25, Death Wish 30, Last Stand 20, Sweeping Strikes 30 and others.
  - F1 and F7 skip a talent until it is cast (`assessAbility`, `defensivesAt`), so the tier level never reaches a warrior coach line.
  - The rendered rows follow the orders: Last Stand at 28, Death Wish at 30, Reckless Swings at 31 on Fury and 57 on Arms raid.
- **Round 25's page fixes.**
  - arms.md:105 matches the Arms table from 50.
  - protection.md:109 matches the rendered preset.
- **Round 25's coach rewrite.** Re-read from scratch (Coach line sweep). The checks of round 25's own fixes are listed under Decisions.

Note-revert scan: for every current sentence or note in the ten warrior files, I looked back along first-parent history for a version that had it, then dropped it, then brought it back.

- The only recent hit is the Fury Whirlwind `level_note`, which the ledger accepts.
- The fury.md and Arms coach hits are from rounds 15 to 16 and 11. Each line has been audited since.

## Orders (index 0 = level 10)

| Spec | Order |
|---|---|
| Arms leveling | Improved Heroic Strike 10; Deflection 11-13; Improved Rend 14-16; Improved Charge 17-18; Improved Tactical Mastery 19-21; Improved Overpower 22-23; Anger Management 24; Deep Wounds 25-27; Two-Handed Weapon Specialization 28-30; Impale 31-32; Bloodthrill 33-34; Sweeping Strikes 35; Weaponmaster 36-39, 41; Mortal Strike 40; Improved Slam 42-43; Cruelty 44-48; Unbridled Wrath 49-53; Improved Bloodrage 54-55; Iron Will 56-60 |
| Arms raid | Improved Heroic Strike 10-12; Improved Rend 13-15; Improved Tactical Mastery 16-18; Improved Overpower 19-20; Anger Management 21; Deep Wounds 22-24; Spearing Strike 25; Two-Handed Weapon Specialization 26-28; Impale 29-30; Bloodthrill 31-32; Sweeping Strikes 33; Weaponmaster 34-38; Improved Slam 39, 41; Mortal Strike 40; Cruelty 42-46; Unbridled Wrath 47-51; Enrage 52-56; Reckless Swings 57; Improved Execute 58-59; Improved Bloodrage 60 |
| Fury (both orders identical) | Cruelty 10-14; Unbridled Wrath 15-19; Furious Precision 20-22; Enrage 23-24; Dual Wield Specialization 25-29; Death Wish 30; Reckless Swings 31; Improved Execute 32-33; Improved Berserker Rage 34-35; Improved Whirlwind 36-37; Flurry 38-39, 41; Bloodthirst 40; Improved Heroic Strike 42-44; Improved Rend 45-47; Improved Charge 48-49; Improved Tactical Mastery 50-52; Anger Management 53; Deep Wounds 54-56; Impale 57-58; Improved Bloodrage 59-60 |
| Protection (both orders identical) | Improved Bloodrage 10-11; Shield Specialization 12-16; Anticipation 17-21; Improved Revenge 22-24; Improved Thunder Clap 25-27; Last Stand 28; Master of Defense 29-30; Defiance 31-33; Improved Sunder Armor 34-36; Concussion Blow 37; Focused Rage 38-39, 41; Shield Slam 40; Bastion 42-46; Improved Heroic Strike 47-49; Deflection 50-54; Improved Rend 55-57; Deep Wounds 58-60. No Improved Demoralizing Shout, no Improved Shield Wall. |

Trainer levels: Thunder Clap 6, Hamstring 8, Bloodrage 10, Overpower 12, Revenge 14, Demoralizing Shout 14, Shield Block 16, Cleave 20, Slam 20, Victory Rush 20, Execute 24, Shield Wall 28, Berserker Stance and Berserker Rage 30, Whirlwind 36, Recklessness 50.

## Findings

Counts, guide pages and cards: CONFIRMED 1 (R26-1), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

Counts, coach: CONFIRMED 8 (R26-2 to R26-9), LIKELY 2 (R26-10, R26-11), NEEDS-SIM 0, NEEDS-OWNER 0.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| WARRIOR-R26-1 | CONFIRMED | card note: contradicts its own row condition, the page and the coach (sim round 22 text, new work) | `warrior-fury.json` row 11 (Heroic Strike) main note, rendered 10 to 60, and its `level_notes` step 11 at 60 | "Only above 60 rage, and, once you have Execute, never in the Execute phase: hold that rage for it." / at 60: "... never in the Execute phase: hold that rage for Execute." | The row's condition is `not isExecutePhase E25`: no Heroic Strike from **25%** down. On the same card, "the Execute phase" means E20: the Death Wish, Recklessness and Execute rows all use `isExecutePhase E20` and call it "the Execute phase". fury.md:90 says "From 25%, stop queueing Heroic Strike and let the bar fill", and coach/warrior-fury.md:21 says "never from twenty-five percent down". A player who follows the note queues Heroic Strike from 25% to 20%, which is exactly the rage the card holds for the first Execute. Before sim round 22 the note read "never from 25% down", which was right. | Main note: "Only above 60 rage, and, once you have Execute, never from 25% down: hold that rage for it. It is queued on your next swing and costs no global cooldown." `level_notes` step 11: "Only above 60 rage, and never from 25% down: hold that rage for Execute. It is queued on your next swing and costs no global cooldown." Regenerate `addon-data.json` (note text only). Note for the note-revert check: this restores pre-sim-22 wording on purpose. |
| WARRIOR-R26-2 | CONFIRMED | coach: the persona does not cover an aura the restored rule judges | `coach/warrior-protection.md:29` (mistakes, uptime) and `:44` (drill.uptime); the rule is `warrior-protection.json` auras, Demoralizing Shout with `covered_by` (int 1aaa72106) | ":29 uptime: Shield Block down between casts, or Thunder Clap falling off the boss."; ":44 Watch the Shield Block buff, not its cooldown. The moment it drops, press it again. Check Thunder Clap on the boss each time you do." | F2 now judges three tank auras: Shield Block, Demoralizing Shout and Thunder Clap. A low Demoralizing Shout uptime opens an uptime finding whose drill is :44, and :44 never mentions the Shout. The persona's own priority (:21) and adherence drill (:50) carry Demoralizing Shout with the druid's Roar as the exception. These agree with the rule: the Roar counts as covered (`withCover`), and another warrior's Shout is "not yours to judge" (`appliedByOthers`). Only :29 and :44 lag behind. | Hand edit:<br>- :29 "uptime: Shield Block down between casts, or Thunder Clap or Demoralizing Shout falling off the boss."<br>- :44 "... Check Thunder Clap and Demoralizing Shout on the boss each time you do; a druid's Roar holding the slot counts as your Shout." |
| WARRIOR-R26-3 | CONFIRMED | coach: two superlatives the card's own measurements break | `coach/warrior-protection.md:21` | "Shield Block whenever its buff is down; of everything you press, it is the one whose removal costs the most damage taken"; "Thunder Clap ... because its attack-speed slow is the largest cut to damage taken after Shield Block." | The Protection card's notes give three measurements:<br>- Shield Block: about 18% less damage taken.<br>- Thunder Clap upkeep: about 16%.<br>- Demoralizing Shout: "damage taken about 22 percent lower", because "the shout's 204 takes about 23 percent off every swing".<br>The arithmetic agrees. Thunder Clap's 20% longer swing timer is 1 − 1/1.2 = 16.7% fewer swings, and the Shout's 204 of 268 boss attack power takes about 23% off each swing. So Thunder Clap is not the largest cut after Shield Block, and on the card's figures Shield Block is not the largest either. The persona was written on 27 October, before the round 14 Demoralizing Shout ruling. | Hand edit: "Shield Block whenever its buff is down: it is your mitigation and your rage engine at once, and recasting over a live buff only wastes rage." and "Thunder Clap kept on the boss, refreshed in its last seconds: its attack-speed slow is a straight cut to the damage you take." (protection.md:91's own reason). Sim requested below to settle the order on one harness, but the fix does not wait for it. |
| WARRIOR-R26-4 | CONFIRMED | coach: buttons named without the level they arrive (the round 26 lens; R25-15 levelled only Shield Slam) | `coach/warrior-protection.md:17, :21, :23, :28, :43, :45, :49, :50` | ":21 Shield Block whenever its buff is down ... Revenge the moment it is available ... Demoralizing Shout next ... Before forty, Revenge and Sunder Armor carry your threat."; ":23 Last Stand and Shield Wall are for a fight the healers are losing ..."; ":28 dying with Shield Wall or Last Stand still ready"; ":49 Put Shield Wall on a key ..." | Levels: Revenge 14, Demoralizing Shout 14 and Shield Block 16 (trainer), Shield Wall 28 (trainer), and Last Stand at 28 on both orders. The rendered card shows each from those bands (Revenge and Demoralizing Shout 14, Shield Block 16, Last Stand and Shield Wall 28). The coach reads dungeon bosses from the low teens, and Ragefire Chasm and the Deadmines are tanked at 13 to 18. A tank at 13 is told to lead with Shield Block and Revenge. A tank at 20 to 27 is told to hold Shield Wall and Last Stand for the bad moment. He has neither. | Hand edit:<br>- :21 "Shield Block from sixteen whenever its buff is down", "Revenge from fourteen the moment it is available", "Demoralizing Shout from fourteen next", "Before forty, Revenge and Sunder Armor carry your threat, and before fourteen Sunder Armor alone".<br>- :23 open with "From twenty-eight, Last Stand and Shield Wall ...".<br>- :49 "From twenty-eight, put Shield Wall ...".<br>- :50 leave as is (a dummy drill). |
| WARRIOR-R26-5 | CONFIRMED | coach: talent and trainer buttons with no level (R25-13 fix incomplete) | `coach/warrior-fury.md:21, :23` | ":21 Death Wish and Recklessness for the execute phase, as the burst window. ... In execute range, Execute comes first ... Reckless Swings on cooldown right behind Whirlwind ..."; ":23 Recklessness is a once-a-fight button and goes in execute range too. On a short pull, press both in the closing seconds" | Levels: Death Wish 30 and Reckless Swings 31 (both orders), Recklessness 50 (trainer, rendered from 50), Execute 24. R25-13 named these levels, but the fix only added "Before thirty-six ..." and "from forty" to Bloodthirst. A Fury at 31 to 35 is told to press Reckless Swings "right behind Whirlwind" with no Whirlwind; the card says "once you have it". A Fury at 30 to 49 is told to hold Recklessness for execute. | Hand edit:<br>- :21 "From thirty, Death Wish for the execute phase as the burst window, and from fifty Recklessness with it."<br>- "From twenty-four, in execute range, Execute comes first ..."<br>- "From thirty-one, Reckless Swings on cooldown, right behind Whirlwind once you have it at thirty-six."<br>- :23 "From fifty, Recklessness is a once-a-fight button ...". |
| WARRIOR-R26-6 | CONFIRMED | coach: R25-4's fix is incomplete; a multi-target line still stands alone, and lines miss their level | `coach/warrior-arms.md:36, :37, :30` (generated, int head) | ":36 Whirlwind hits up to four targets and every Sweeping Strikes charge copies a hit onto a second one."; ":37 Use Cleave instead of Heroic Strike when two mobs are in front of you."; ":30 Spearing Strike in the raid build, on any free global in Battle Stance." | The list's header says "Your buttons, in the order you press them". The R25-4 override folded Whirlwind into the Sweeping Strikes line (:35, "only with two or more targets, from thirty-six"). The generator still emits the page's next sentence as its own button, with no target condition and no level. On the page it sits under "**Two or more targets.**" (arms.md:93). The Arms card never casts Whirlwind on one target. Cleave trains at 20. Spearing Strike is the raid order's 25th point (level 25 there). | `starter_overrides`:<br>- `EXCLUDED["warrior-arms"]` "Whirlwind". Its fact is already in the Sweeping Strikes line, or append ": Whirlwind hits up to four targets." there.<br>- `GUIDE_LINES["warrior-arms"]["Cleave"]` "From twenty, use Cleave instead of Heroic Strike when two mobs are in front of you."<br>- `["Spearing Strike"]` "From twenty-five in the raid build, Spearing Strike on any free global in Battle Stance. It stops at the execute swap."<br>Regenerate with `persona_starter.py --spec warrior-arms`. |
| WARRIOR-R26-7 | CONFIRMED | coach: consumables named at the top band for every level (the R25-9 shape; Fury was levelled in round 25, Arms and Protection were not) | Arms `:45, :49, :58` (generated: `starter_kit.consumable_paragraph`, `preparation_mistake`, `preparation_drill`); Protection `:32, :37, :48` (hand) | Arms: "Before a boss you want food (Smoked Desert Dumplings); a weapon enchant (Elemental Sharpening Stone); a combat potion (Mighty Rage Potion) ..."; Protection: "Flask of the Titans ..., Smoked Desert Dumplings, an Elemental Sharpening Stone, and the Mighty Rage Potion"; ":48 check flask, food buff, sharpening stone, a Mighty Rage Potion ..." | Item levels (`items.json`): Smoked Desert Dumplings 45, Mighty Rage Potion 46, Elemental Sharpening Stone 50, Flask of the Titans 50. The presets (`raid-consumes.json`) carry nothing that F6 checks below 35. From 35 they carry Poached Sunscale Salmon and a Dense stone, and from 50 (Protection) a flask. Since DRUID-R26-6, F6 expects only the band's items, so the persona and the check now disagree below 50. A level-40 tank in Scarlet Monastery is told to check a flask and a Mighty Rage Potion he cannot use. | Generator (shared root, every starter): `kit()` gives each item the first `from_level` it appears at in the spec's preset, and `consumable_paragraph` / `preparation_drill` write "from N" when N is above the earliest dungeon level, as the Fury hand edit does. Protection hand edit:<br>- :37 "From fifty, Flask of the Titans (or Flask of Chromatic Resistance where the fight calls for it); Smoked Desert Dumplings from forty-five, Poached Sunscale Salmon before; a sharpening stone, Elemental from fifty; from forty-six the Mighty Rage Potion."<br>- :32 and :48 "from fifty a flask", "from forty-six a Mighty Rage Potion". |
| WARRIOR-R26-8 | CONFIRMED | coach: "What changed for Forever" lists things that did not change (the R25-10 shape) | `coach/warrior-fury.md:25`; `coach/warrior-protection.md:25` | Fury and Protection: "Warriors wear plate from forty."; Protection: "Shield Block and Shield Slam need a shield in the off hand. Devastate is not learnable in Forever." | In Classic, warriors also learned plate at 40 (Warcraft Wiki, Plate armor: "Prior to Legion, warriors and paladins could not equip plate armor until level 40"). Shield Block and Shield Slam needed a shield in 1.12 too. Devastate was never a Classic warrior ability; it is a TBC talent. The Forever client lists it only as a `class_spell` with no SkillLineAbility row, so nothing changed for a warrior. Under this heading, each line tells the player something is new that is not. | Hand edit: drop "Warriors wear plate from forty." from both, and drop the two Protection sentences. If plate is worth keeping, move it out of the heading: "You wear mail until plate at forty." |
| WARRIOR-R26-9 | CONFIRMED | coach F1: a cooldown the shipped build shortens is judged and quoted at base (HUNTER-R26-6 shape; errs lenient) | `coachrules.py` `warrior-fury` core "Whirlwind"; `f1_core_casts.go` (`sc.CooldownMS`); `copy.go` `BasisCooldown` | Finding basis: "Measured against Whirlwind's 10-second cooldown and your own rage bar." | Improved Whirlwind is at 36 to 37 on both Fury orders: 9 s at 36 and 8 s from 37 (client: "Reduces the cooldown of your Whirlwind ability by 2 sec"; sim round 22 F1, "16.3 (8 s with Improved Whirlwind 2)"). Whirlwind trains at 36. So every Fury warrior who can be judged on Whirlwind has the talent, and the basis line states a cooldown that is false for all of them from 37. The "possible" count is low for the same reason. | The shared code fix routed from HUNTER-R26-6 (talent cooldown modifiers read from the build). Nothing to change in `coachrules.py`. |
| WARRIOR-R26-10 | LIKELY | coach F1: core abilities the card holds through the execute phase (the WARLOCK-R26-11 shape) | `coachrules.py` `warrior-fury` core "Bloodthirst", "Whirlwind"; `f1_core_casts.go` (`freeWindows` excludes only downtime) | F1 text: "{lost} of a possible {possible} went unused while you had the rage." | The card, fury.md:90 and coach :21 all put Execute ahead of Bloodthirst and Whirlwind below 20%. Execute costs 10 with Improved Execute (32 to 33) and empties the bar, so neither button is cast in the execute phase. F1 has no execute window. Every second with 30 (or 25) rage in the bar counts as ready time, and the per-second reading is the last power value in that second, which can come before the next Execute. Recklessness makes every swing crit, so the bar refills fast. Sim round 22 F1 has Bloodthirst at 22.1 of a possible 30 in 180 s, and the execute phase alone is the gap. Whether it reaches CoreFiresLost 2 and 10% depends on the log, hence LIKELY. When it fires, the sentence calls the guide's own play a loss. | Coach code: give a core entry an optional "held below 20% for Execute" flag that F1 reads as an excluded window from the boss's first sub-20% reading. Or drop Whirlwind from the Fury core list, since it also sits behind Bloodthirst, and flag Bloodthirst. Test first with a Fury fixture whose last 20% has Executes only. |
| WARRIOR-R26-11 | LIKELY | coach F1: Bloodrage held above a rage ceiling the coach cannot see | `coachrules.py` core "Bloodrage" in all three specs; `f1_core_casts.go` `newResourceCheck` (cost type health, so no resource check) | (rule entries) | Each card holds Bloodrage at a ceiling: Arms and Fury under 80, Protection under 60. Its cost is health, so F1 checks no bar, and every second at or above the ceiling counts as ready. A gap of 121.5 s (two 60 s cooldowns plus grace) inside one inter-cast interval makes 2 lost. Protection is most exposed. The card's own note says "a Forever tank is cooldown-limited, not rage-limited, at this gear", and the persona's :25 says "at raid gear a Forever tank is limited by cooldowns, not rage". A tank who follows the card and sits above 60 rage for stretches is told Bloodrage was lost. Round 25 checked only Arms ("gaps far below one cooldown"). | Coach code or rules: either give the core entry a rage ceiling (F1 counts ready time only below it), or drop Bloodrage from the Protection core list, where F4 already reads rage waste. See Sims requested. |

## Decisions

- **D1. R26-1 is a guide finding.** It is the only page or card statement on this tree that is false at a band it renders. Every other sim round 22 Fury note holds at every band:
  - Improved Berserker Rage 5 or 10 matches the client ranks.
  - Bloodrage "to 15 and 15" matches Improved Bloodrage's "all the Rage ... by 50%".
  - Reckless Swings "four white swings that land within 8 seconds" matches the client ("Your next 4 melee attacks ... Lasts 8 sec") and sim round 22 S6.
  - Battle Shout "every party and raid member within 20 yards" matches the tooltip of every rank.
- **D2. R26-2 is CONFIRMED although nothing in it is false on its own.** The restored rule and the persona disagree on the same pull: the coach fires on an aura, and the drill it hands the player names two others. The Roar and other-warrior cases agree between rule, card and persona (checked in `withCover` and `judgeAura`).
- **D3. R26-3 is CONFIRMED on the card's own figures and arithmetic, not left for a sim.** The figures come from two harness runs, rounds 11 and 14. The direction holds on the Shout's per-swing cut alone.
- **D4. R26-7 matches the Fury fix round 25 already made.** The druid auditor accepted top-band kit paragraphs (druid.md "Coach consumables paragraphs"). Round 25 made levels the rule for Fury, and F6 now reads bands, so the generated and hand-written kits need the same levels.
- **D5. R26-9 is CONFIRMED although it errs lenient,** as HUNTER-R26-6 D4: the basis line states a false cooldown.
- **D6. R26-10 and R26-11 are LIKELY.** Each needs a log shape (rage readings in the execute phase; long stretches above the ceiling) to fire. When it fires, the sentence is false.
- **D7. Round 25 fixes re-checked:**
  - Applied and right: R25-1, R25-2, R25-3, R25-5, R25-6, R25-8, R25-10, R25-11, R25-12 and R25-14.
  - R25-7: applied. Arms :18 has no meta talk; the Protection "Devastate" line falls under R26-8.
  - R25-16: code, `notYetLearned`.
  - R25-17: now `covered_by`.
  - Shoulds S1 to S7 applied.
  - Incomplete: R25-4 (R26-6), R25-9 Arms and Protection (R26-7), R25-13 (R26-5) and R25-15 (R26-4).
- **D8. Not re-filed:** W23-1 book ranks (Revenge 6, Battle Shout 7, Heroic Strike 9; sim round 22 S1 and S2, owner); S3 Defiance; S4 threat constants; S5 Deep Wounds; carried Shoulds ("biggest hit" for Execute and Shield Slam, Reckless Swings "two or three swings").

## Shoulds (wording, not findings)

| id | where | note | suggested |
|---|---|---|---|
| WARRIOR-R26-S1 | `coach/warrior-fury.md:47` drill.ranks | Whirlwind has one rank (1680 only, `trainables`), so "check ... Whirlwind ... against the rank your level calls for" asks for a check that cannot fail. | Drop Whirlwind from the list. |
| WARRIOR-R26-S2 | `coach/warrior-fury.md:34` mistakes, preparation | "a missing food, elixir or sharpening stone". F6 never checks elixirs: its categories are flask, food, weapon enchant, permanent enchants and potion. | "a missing food or sharpening stone, or a combat potion never used". |
| WARRIOR-R26-S3 | `coach/warrior-arms.md:64-66` Voice | Two healer lines ("For a healer ...", "Never tell a healer to heal more") in a damage persona; template residue. | The generator writes the healer lines only for role healer. |
| WARRIOR-R26-S4 | protection.md:109 | "the rum and Elixir of Fortitude are your health bar". From 50 the table's Flask of the Titans is the biggest health item. Incomplete, not false. | "... the flask, the rum and Elixir of Fortitude are your health bar ...". |
| WARRIOR-R26-S5 | `coach/warrior-protection.md:50` | The adherence list leaves out Battle Shout, which :21 asks for. On a three-minute dummy drill, one pre-drill shout lasts the whole drill, so nothing is false. | Optional: add "and Battle Shout if it drops". |

## Sims requested

- **Protection mitigation order on one harness (R26-3):**
  - Setup: raid 60, the tank fight, 10,000 iterations, seeds 1 and 50001.
  - Runs: remove Shield Block; remove Thunder Clap; remove Demoralizing Shout.
  - Report: damage taken against the shipped card.
  - Purpose: settles which cut is largest, if any line wants to keep a superlative.
- **Bloodrage ceiling (R26-11):**
  - Setup: raid 60 Protection (tank fight) and Fury, seeds 1 and 50001.
  - Report: the share of the fight above the card's ceiling (60 and 80) and the longest stretch between Bloodrage casts. A run of the coach's F1 on that log is better still.
- **Fury execute hold (R26-10):** run F1 on a sim Fury raid-60 log (the coach fixture path) and report Bloodthirst and Whirlwind lost casts inside the execute phase.
- Carried from round 25 if the wording stays: Protection Revenge against Shield Slam per landed cast; Protection race TMI on one shared preset; Arms rage per second, raid against bare.

## Engine questions

- None new. Carried: S3 Defiance 7/13/20, S4 threat constants, S5 Deep Wounds' double multiplier, W23-1 trainer book ranks (owner).

## Sim findings

- None.

## Sources

- Client 1.60.1.70790:
  - `spelltips/warrior.json`: 1160 and 11556 Demoralizing Shout, 1680 Whirlwind, 1719 Recklessness, 2565 Shield Block, 2687 Bloodrage, 11581 Thunder Clap, 18499 Berserker Rage, 1329555 Reckless Swings, 20662 Execute, 23925 Shield Slam.
  - `spellconst/warrior.json`: 11556, with −196 and −1.4 a level, so 204 at 60.
  - `talents/warrior.json`: Improved Whirlwind, Improved Berserker Rage, Improved Bloodrage, Improved Slam, Shield Specialization, Master of Defense.
  - `trainables/warrior.json` and `druid.json`; `addon-builds.json`; `addon-data.json` `rotations/warrior-*`; `items.json` (consumable required levels).
- Coach:
  - `coachrules.py` and `warrior-*.json` (int 22783f43d).
  - `api/internal/coach/{f1_core_casts,f2_uptime,f6_preparation,f7_death,copy,rules}.go`.
  - `catalogue.json`, `sim/request/raid-consumes.json`.
  - `tools/coach/*`. Scratch regen: `scratchpad/gr26/warrior/regen-audit/warrior-arms.md` and `regen-int/warrior-arms.md`, both byte-identical to their trees' files.
- Reviews: round25/warrior.md, ledger Round 25, `2026-11-05-sim-audit-round22/warrior.md` and integration.md, round26/{hunter,druid,warlock}.md.
- Web: Warcraft Wiki, Plate armor, https://warcraft.wiki.gg/wiki/Plate_armor ("Prior to Legion, warriors and paladins could not equip plate armor until level 40").

## Note sweep

One line per row, against both orders (Arms raid and leveling; Fury and Protection orders are identical) and the trainer. "Bands" are the rendered `addon-data.json` bands.

Arms:

- Berserker Stance swap, bands 30 to 60: Recklessness 50 ("From 50"), Improved Tactical Mastery (raid 16 to 18, leveling 19 to 21). OK.
- Bloodrage, 10 to 60: Improved Bloodrage leveling 54 and 55, raid 60, as stated. OK.
- Battle Shout, 10 to 60: 10 rage, 20 yards, raid. OK.
- Rend, 10 to 60: Bloodthrill "once ... in". OK.
- Overpower, 12 to 60: Mortal Strike "from 40"; Improved Overpower 2/2 (raid 20, leveling 23) given as a condition. OK.
- Mortal Strike, 40 to 60: 6 s cooldown, 50% healing cut. OK.
- Reckless Swings, 57 to 60: "raid build only", 8 s, 5 rage, a quarter of base health. OK.
- Execute, 24 to 60: "from 40" for Mortal Strike. OK.
- Slam, 20 to 60: Improved Slam condition (raid 39, leveling 42), Mortal Strike "from 40". OK.
- Spearing Strike, 25 to 60: "raid build only", 75% and twice that on dragons and giants. OK.
- Hamstring, 10 to 60: Windfury, 25%. OK.
- Heroic Strike, 10 to 60: Overpower 12, Execute 24, Mortal Strike 40. OK.

Fury:

- Bloodrage, 10 to 60: Improved Bloodrage at 1/2 and 2/2 given as conditions (59, 60); 15 and 15 at 2/2. OK.
- Death Wish, 30 to 60: E20, last 30 s. OK.
- Recklessness, 50 to 60: once a fight (30 min cooldown). OK.
- Battle Shout, 10 to 60: OK.
- Execute, 24 to 60: "from 40" for Bloodthirst. OK. The 60 `level_note` is OK.
- Bloodthirst, 40 to 60: OK.
- Whirlwind, 36 to 60: four targets, both weapons, "from 40". OK. The 60 `level_note` is OK.
- Reckless Swings, 31 to 60: "once you have it" (Whirlwind 36). OK. The 60 `level_note` is OK.
- Berserker Rage, 36 to 60: Improved Berserker Rage 34 to 35, 5 or 10. OK. The 60 `level_note` is OK.
- Overpower, 12 to 35: OK. The 60 `level_note` is OK.
- Heroic Strike, 10 to 60: **R26-1** (main and 60 note).

Protection:

- Defensive Stance, 10 to 60: OK.
- Last Stand, 28 to 60: 30% for 20 s. OK.
- Shield Wall, 28 to 60: 60%, 12 s, 15 min (no Improved Shield Wall in the order). OK.
- Bloodrage, 10 to 60: Improved Bloodrage at 10 and 11. OK.
- Shield Block, 16 to 60: 75%, 7 s, off the global cooldown. OK.
- Revenge, 14 to 60: most threat per rage (5, or 2 with Focused Rage). OK.
- Shield Slam, 40 to 60: OK.
- Thunder Clap, 10 to 60: OK.
- Demoralizing Shout, 14 to 60: "does not stack with a druid's Demoralizing Roar". OK, and it agrees with the `covered_by` rule.
- Battle Shout, 10 to 60: OK.
- Sunder Armor, 10 to 60: OK.
- Heroic Strike, 10 to 60: above 40. OK.

## Coach line sweep

warrior-arms.md (generated; int head):

| Line | Verdict |
|---|---|
| :18 | OK |
| :24 Bloodrage | OK |
| :25 the pull | OK |
| :26 Rend | OK |
| :27 Overpower | OK |
| :28 Mortal Strike, Reckless Swings | OK |
| :29 Slam | OK |
| :30 Spearing Strike | R26-6 |
| :31 Heroic Strike | OK |
| :32 Hamstring | OK |
| :33 Execute | OK ("biggest hit" Should carried) |
| :34 Recklessness | OK |
| :35 Sweeping Strikes | OK |
| :36 Whirlwind | R26-6 |
| :37 Cleave | R26-6 |
| :40 to :44, :46, :47 mistakes | OK |
| :45 | R26-7 |
| :49 | R26-7 |
| :53 to :57, :59, :60 drills | OK |
| :58 | R26-7 |
| :64 to :73 voice | S3 |

warrior-fury.md (hand):

| Line | Verdict |
|---|---|
| :17 | OK |
| :21 | R26-5 (Death Wish, Recklessness, Reckless Swings, Execute). Bloodthirst "from forty" and the Heroic Strike 25% stop are OK |
| :23 | R26-5 (Recklessness); the rest OK (Death Wish 3 min and 5% damage taken; Bloodrage 60 s at 20% of base health) |
| :25 | R26-8 (plate). Whirlwind both weapons, Unbridled Wrath 60%, Bloodthirst 45% plus a little and 10% speed, Reckless Swings, Dual Wield Specialization are OK |
| :28 to :33, :35 | OK |
| :34 | S2 |
| :37 | OK (levels as round 25) |
| :39 | OK |
| :43 to :46, :48 to :50 | OK |
| :47 | S1 |

warrior-protection.md (hand):

| Line | Verdict |
|---|---|
| :17 | R26-4 (Shield Block, Revenge) |
| :21 | R26-3, R26-4. Demoralizing Shout and Battle Shout lines agree with card and rule; Revenge per rage OK |
| :23 | R26-4 (Last Stand, Shield Wall). Demoralizing Shout 204 at 60 OK (196 + 1.4 × 6) |
| :25 | R26-8. Thunder Clap in Defensive Stance, Shield Specialization and Master of Defense rage OK |
| :28 | R26-4 |
| :29 | R26-2 |
| :30, :31, :33 to :35 | OK |
| :32 | R26-7 |
| :37 | R26-7 |
| :39 | OK |
| :43, :45, :49 | R26-4 |
| :44 | R26-2 |
| :46, :47 | OK |
| :48 | R26-7 |
| :50 | S5 |

Rule files: `warrior-arms.json` (Bloodrage and Mortal Strike core; Battle Shout and Rend auras; no defensives) is OK except R26-11. `warrior-fury.json`: R26-9, R26-10, R26-11. `warrior-protection.json`: R26-11; the Demoralizing Shout aura with `covered_by` is OK against card and persona :21 and :50, and against the persona's :29 and :44 only after R26-2.

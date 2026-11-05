# Guide accuracy round 23: Priest

Audited on the round 23 audit tree (24ee457d0: main after guide round 22; the log has no sim round 21 commit and there is no `*sim-audit-round21*` report, so sim round 21 had not landed). Engine pin d5c708834 (`sim/enginever/version.go`; source read at `scratchpad/gr23/engine`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited, line by line:

- `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`.
- Every note and every row condition in `data/curated/apl/priest-{discipline,holy,shadow}.json`.
- Every rendered row of `addon-data.json` (`rotations/priest-discipline` 25 bands, `priest-holy` 22, `priest-shadow` 22, with `race_spell_ids`).
- All six priest orders in `addon-builds.json`, re-decoded against `talents/priest.json`.
- `trainables/priest.json`, `spelltips/priest.json`, `spelltips/shared.json` (racials, consumables), `items/*.json`, `bis/priest-*.json`.
- Engine: `sim/core/spell.go` (level gap), `spell_max_level_auto_gen.go`, `sim/priest/heal_ranks.go`, `prayer_of_mending.go`, `healing_talents.go` (Litany of Light), `racials.go`.

Read first: round 22 priest, the ledger's Round 19 to 22 sections, the Round 23 reports already written (druid, hunter, mage, paladin, shaman, warrior) for how the new shapes were filed, and the older priest reports on the items below. The round 1 and round 2 briefs named by the round 23 brief are no longer on disk; statuses and report structure follow round 22's priest report and the round 23 notes.

Lenses:

- Round 20's note sweep (appendix).
- Round 21's sweep of every superlative, comparison and quoted cost, duration and cooldown.
- Round 22's shapes.
- Round 23's additions: the stated reason of each note and line, "whenever" and "every time" against the row's own condition, absolute claims a far-below-level rank breaks, per-mana claims, both sides of every comparison over all bands.

## Changes since round 22's tree (987d08a7c)

None for priest. `git diff` over the four guides and the three APLs is empty, and the three priest entries of `rotations` in `addon-data.json` are identical to round 22's tree (compared in code). Round 22 made no priest fix. Everything below was there at round 22 and earlier, and it surfaces under this round's lenses.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R23-1 | CONFIRMED | card note and page prose ("whenever" against the row's own condition) | `priest-discipline.json` priorityList[7], first sentence, Discipline bands 10 to 60; discipline.md:68 | Note: "On the tank at the top rank whenever Weakened Soul is off him". Page: "Shield the tank at the top rank whenever Weakened Soul is off him. In Forever a shield no longer costs a warrior any rage, so there is nothing to hold back" | The row's condition has three parts: `NOT auraIsActive(6788)` on the tank AND tank health below 90% AND `currentManaPercent >= remainingTimePercent x 0.9`. So the card does not shield a tank at 90% or more, and it holds the shield whenever mana falls behind the fight's pace. The APL's own header says so ("the tank is shielded when he is hurt and Weakened Soul is off", "behind that pace the filler waits ... and the shield waits"). "Nothing to hold back" is what the card contradicts: rage is no reason to hold it, but mana is. At 60 the shield is the third dearest thing the row list paces: 375 mana (Soul Warding and Mental Agility) for about 1,113 plus its bonus. The same page's Mana paragraph (:84) tells a priest behind the fight to cut back, so the page argues with itself. | Note: "On the tank at the top rank once he is hurt and Weakened Soul is off him, while your mana keeps pace with the fight: it soaks the next hit and costs a warrior no rage." (later sentences unchanged). :68: "Shield the tank at the top rank when he is hurt, Weakened Soul is off him and your mana is keeping pace with the fight. In Forever a shield no longer costs a warrior any rage, so rage is no reason to hold it back: it soaks the next hit, and the Weakened Soul it leaves is what Renewed Hope needs." (rest unchanged) |
| PRIEST-R23-2 | CONFIRMED | card note (names a spell on bands before the player has it; round 22 shape, missed by the sweeps since the wording arrived) | `priest-shadow.json` priorityList[6], first clause, the second Mind Blast line, Shadow bands 10, 16, 18, 20 on both orders, and 22 and 26 on the leveling order | "Mind Blast on cooldown, ahead of Mind Flay" | Mind Flay is a talent: level 28 on the leveling order, level 21 on the raid order (`addon-builds.json`, index 0 = 10). The note renders from band 10. So at 10 to 20 no priest on either order has it, and at 21 to 27 the leveling priest does not. shadow.md:139 says the same ("then Mind Flay at 28. Until 28 the wand is your filler"). This is the shape round 22 filed on Multi-Shot before 18 and Hemorrhage before 30. Round 20's rewrite and every sweep since read only the rest of this sentence. The Shadow Word: Death note's "ahead of Mind Flay" is fine: that row starts at 32. | "Mind Blast on cooldown, ahead of Mind Flay once you have it: a big hit for its short cast, at half cost in Shadowform." |
| PRIEST-R23-3 | CONFIRMED | stat prose (absolute a far-below-level rank breaks; contradicts the same page; the shape of DRUID-R23-2) | holy.md:88; discipline.md:88 | Holy: "Healing power is the size of every heal." Discipline: "Healing power is the size of every shield, Penance and heal." | Each page's Low ranks paragraph says otherwise: "At 60, Lesser Heal and the first ranks of Heal and Renew ... heal for little more than their base, and Renew rank 5 and 6 and Heal rank 4 keep only part of it". Client level gap (`SpellLevelPenalty`; MaxLevel = learn level + 5, grace 12, 5% a level): at 60, Lesser Heal rank 3 keeps none (from 47), Renew ranks 1 to 3 none, Heal ranks 1 and 2 none, Heal rank 4 55%, Renew rank 5 45%, rank 6 75%. The Holy card keeps Heal rank 4 as its last row at 60 (55%). The Discipline card casts Lesser Heal rank 3 at 10 to 38, at 45% of its bonus by 38. | Holy: "Healing power is the size of every heal, short of a rank far below your level." Discipline: "Healing power is the size of every shield, Penance and heal, short of a rank far below your level." |
| PRIEST-R23-4 | LIKELY | card note (stated reason backward; dearest holds only in some gear) | `priest-holy.json` priorityList[4] and [5], Holy bands 20 to 60 | [4] "it is your dearest heal, so spend it where the size matters." [5] "when someone is about to die, the size matters more than the mana." | Size is the one thing Flash Heal never has, on either side of the comparison at any band. Its top rank against the slow heal the player has (base): 20, 194 against Heal rank 1's 270; 26, 249 against 375; 38, 383 against 651; 44, 507 against Greater Heal rank 1's 844; 60, 825 against rank 4's 1,782. Its coefficient is half too (0.429 against 0.857). At 60 raid (554 healing, Spiritual Healing), Flash Heal rank 7 heals about 1,188 and Greater Heal rank 4 about 2,482, so even per second of casting the slow heal wins with Divine Fury (about 790 against 990). What a player pays Flash Heal's mana for is that it lands a second sooner. The Discipline card already says so ("the fast heal, and the dear one"). The reader of [4] and [5] is told the opposite of why the spell exists at under 30%. "Top rank rather than a low rank" is a defensible reading of "size", hence LIKELY. "Dearest" also fails with little healing gear (R22-S3): from 33, Renew heals less per mana until about 40 healing at 33 and about 110 at 50. Round 22 held that on the band presets (D2); the rewording below retires it. holy.md:74 is in the raid section and stays. | [4] "it heals less for its mana than your slower heals, so spend it where the speed matters." [5] "when someone is about to die, the speed matters more than the mana." |
| PRIEST-R23-5 | LIKELY | card note and page prose (a flat count where the client says "up to", beside a recast that takes the leftovers) | `priest-holy.json` priorityList[2], first sentence, Holy bands 40 to 60; holy.md:30 | Note: "Prayer of Mending on the tank every time it is ready: five heals for one cast on a 10 second cooldown". Page: "five heals for one cast" | Client (401859, 1240826, 1240827): "Jumps up to 5 times ... This spell can only be placed on one target at a time per caster." So a cast is up to five heals, and a recast takes the unspent charges off the last holder. The engine does the same (`prayer_of_mending.go` `place`: "taking them off the previous holder"). The same note says to recast every 10 seconds. After its first jump the charge sits on whoever it jumped to until that player is hit or takes a direct heal, so five inside 10 seconds is the best case, not the rule. "The cheapest healing you can cast all fight" holds at the full count (about 12.6 a point of mana against Renew rank 7's 5.5 at 60). One heal is about 2.2 a point of mana, so it needs about 3 heals a cast at 60 (and about 2.3 at 40) to stay ahead of Renew. That number is under Sims requested. | Note and page: "up to five heals for one cast". Keep "cheapest" unless the sim below lands under 3 heals a cast. |

Counts: CONFIRMED 3, LIKELY 2, NEEDS-SIM 0 as findings (the "cheapest" half of R23-5 and the items below are under Sims requested), NEEDS-OWNER 0.

## Shoulds (wording, not findings)

New this round:

- **PRIEST-R23-S1, shadow.md:36 "a second big instant".** Mind Blast has a 1.5 s cast, so there is no first big instant. The plain reading is "a second big hit, and instant", which is true. Optional: "a second big hit, instant, on an 18 second cooldown".
- **PRIEST-R23-S2, shadow.md:89 "until the last 6 seconds" against the card's 10.** The Shadow Word: Pain row gates on `remainingTime >= 10`; the note says only "the closing seconds". The page's own arithmetic for 6 holds (rank 8 is 127 a tick against half a rank 6 Mind Flay at about 237 base, and coefficients 0.366 for two ticks against about 0.273). So nothing stated is false; the card is more conservative. See Sims requested.
- **PRIEST-R23-S3, holy.md:74 "While leveling, a lower rank keeps your bar alive."** Every Holy Flash Heal row casts the top rank at 20 to 59, and the note says "at the top rank". Solo and dungeon can differ, so not false. Optional: "Solo, a lower rank keeps your bar alive."
- **PRIEST-R23-S4, discipline.md:124 Renew ranks by race.** "Rank 4, then rank 5 from 45": the Human rows cast rank 7 at 44 and 45. "From 56, take Renew rank 9": the Undead rows (the Horde pick) cast rank 7 at 56 and 58, and Troll to 59. The page already names race exceptions for Shadow's Devouring Plague. These are race-measured rows inside noise.
- **PRIEST-R23-S5, holy.md:122 "all or nearly all of your healing bonus".** Renew rank 6 keeps all of it to 55, then 85% at 58 and 80% at 59 (rows to 59). Optional: "most or all".
- **PRIEST-R23-S6, `priest-discipline.json` Heal note "your biggest heal until then".** Against a Human's once-a-fight Divine Grace at 26 and 27: rank 3 is 458 + 0.429 x healing, Heal rank 2 is 387 + 0.857 x healing, equal at 166 healing. On the band 20 preset (154) that is 524 against 519, a tie. Divine Grace needs a target under 50%. Round 22's D1 reads comparisons as all-fight spells. Not filed: two levels, one race, a tie.
- **PRIEST-R23-S7, `priest-shadow.json` Shadow Word: Death note "30% of your health".** The tooltip and shadow.md:83 say maximum health. Optional: "30% of your maximum health".
- **PRIEST-R23-S8, holy.md:70 against holy.md:72 and index.md:29.** "Heal rank 4 ... the last step down" sits beside "downrank only within the ranks you learned in your forties and fifties"; rank 4 is learned at 34. The sentence names its own limit.
- **PRIEST-R23-S9, shadow.md:103 "every spell you cast grows with them".** True at the ranks the 60 card casts. A low rank still grows, though less (Devouring Plague rank 2 keeps 35% at 58). Not the shape of R23-3, because no line on the page says otherwise.
- **PRIEST-R23-S10, discipline.md:76 "on anyone about to die".** Omits the party rows' mana gate (mana at least 70% of the fight left). The note has it. The druid report's restoration.md:73 is the same shape, filed as a Should.
- **PRIEST-R23-S11, discipline.md:31 and the note, "soaks the next hit".** A top-rank shield (about 1,100 plus its bonus) soaks part of a raid boss's hit. Optional: "soaks the next hit, or most of it".
- **PRIEST-R23-S12, Holy rank 1 row against "your mana runs short" (note [20], holy.md:70).** Row [20] carries the same mana-pace gate as the top rank, and behind pace rows [15] to [19] still cast the top rank on a deep hole. So at 60 the card reaches rank 1 only when it cannot pay for the top rank. Literal, not false. See Sims requested.

Carried, unchanged and optional: R22-S1 (discipline.md:49 "your best heal"), R22-S2 (shadow.md:60 "the cheapest hit you will ever get"), R21-S1 to S4 (R21-S3, Undead and Dwarf Mind Blast rank 6 at 50 and 51 against shadow.md:137 "back to the top rank"), R20-S1, R20-S2, R18-S1 to S5, R17-S2. R22-S3 is folded into R23-4.

## Checked and right this round

- **Builds and orders** decoded by tier and column. Discipline 29/22 (the round 22 report's "29/18" was a typo; the strings total 51). Holy 17/34. Shadow leveling 20/31 and raid 18/33. Discipline and Holy raid orders equal their leveling orders. The raid swap is 9 for 9 as shadow.md states. The Power Infusion swap is legal (30 Discipline points with Renewed Hope 5, Penance prereq 105834; Holy still clears tier 3 with Improved Renew at 1).
- **Prereqs**: Spirit of Redemption (105854) is Prayer of Mending's step.
- **Every talent value** quoted on the three pages against rank text: Shadowform, Shadow Weaving, Darkness, Improved Mind Flay, Improved Mind Blast, Improved Shadow Word: Pain, Shadow Focus, Spirit Tap, Blackout, Early Demise, Devouring Contagion, Silence, Shadow Reach, Dark Inspiration, Silent Resolve, Twin Disciplines, Mental Agility, Mental Strength, Meditation, Renewed Hope 4, Divine Aegis, Soul Warding, Improved Power Word: Shield, Power Infusion, Holy Specialization, Divine Fury, Improved Healing, Improved Renew, Twilight Focus, Spiritual Healing, Spiritual Guidance, Litany of Light (engine: by SpellCode, so "even at different ranks" holds), Holy Binding, Spirit of Redemption.
- **Row conditions against "on cooldown", "every time" and "whenever"**:
  - Mind Blast, Inner Focus and Shadow Word: Death are unconditional.
  - Devouring Plague `not dotIsActive` (24 s duration = cooldown); Vampiric Embrace known, not active, 10 s left (the note says each).
  - Shadow Word: Pain `not dotIsActive` and 10 s left (R23-S2).
  - Penance tank under 95%, party under 70% ("on whoever is hurt"); Prayer of Mending unconditional.
  - Renew `not dotIsActive`; Flash Heal thresholds match each note.
  - The Discipline shield is R23-1.
- **Opener arithmetic** on the raid order: Mind Blast lands, then Shadow Word: Pain, Vampiric Embrace, Devouring Plague, Shadow Word: Death at 4.5 s, and Mind Blast is back at 5.5 s (5 points); five damage spells make five Shadow Weaving stacks without Vampiric Embrace.
- **Both sides of every comparison over every band**, with talents later in the order:
  - Greater Heal rank 1 "far less": 370 against 655 (314.5 against 557 once Improved Healing lands, Holy 31, Discipline 60).
  - Heal rank 4 "cheaper still": 305 against 370.
  - Renew ranks "far less than the top rank": 140 against 205 and 250, 170 against 250, 205 against 305 and 365.
  - Renew rank 8 "costs more for it" and rank 7 per mana (crossing at about 266 healing, 464 on the band 50 preset).
  - Greater Heal ranks 3 and 4 "about the same for their mana": at 58 with 500 healing, 3.95 against 3.97 before Spiritual Healing.
  - Rank 1 "whole until then" (to 57); "a little" 85% at 60.
  - Holy Power Word: Shield "costs more mana than it saves": about 450 mana for about 1,000, against Greater Heal's 4.4 a point.
  - Flash Heal "costs the most mana for what it heals" in the raid section at raid gear (Discipline :76, Holy :74).
- **Low ranks**: Renew rank 1 keeps 65% at 32 ("most"); ranks 4, 5, 6 keep all at 40, 45, 52; Lesser Heal rank 3 none from 47; Heal rank 4 55% at 60. Heals learned below 20 keep their coefficient in Forever (round 11 C1), so the ledger's Round 22 engine question on the sub-20 cut changes no priest line.
- **Every quoted cost, duration and cooldown**: Shadowform 40% of base mana; Weakened Soul 15 s; Lightwell 3 min or 5 clicks; Divine Grace, Dark Sacrifice, Contingency Plan and Desperate Prayer 10 min; Inner Focus 3 min; Mind Blast 8 s; Shadow Word: Death 18 s and 30%; Devouring Plague 24 s; Vampiric Embrace 30 s and 20%; Fade 10 s on 30 s; Silence 5 s; Flash Heal 1.5 s; Penance 12 s, three pulses; Eureka! 8 s on 2 min; Elune's Light 15 s on 3 min; Berserking 10% on 3 min; Touch of the Grave 5% and 15 s (priest row 30%).
- **Race lines** against `races.json`, `recommendedRaces` and the last measurements (round 3 race-sims; Discipline Human against Gnome from sim round 19, "level", ledger Round 17). Each race's two priest spells on index.md match `trainables`. Skyborne's class list matches `combos.json`. No paid race is recommended. Eureka! lines match `eureka_auto_gen.go` (round 17).
- **Consumables and items**: Flask of Supreme Power, Elixir of Shadow Power and Greater Arcane Elixir are not exclusive in their tooltips (only flasks say "one at a time"); Major Spellblasting Potion, Major Mender's Potion (30 s), Sage's Tea, Greater Mageblood, Greater Cleric's, Elixir of Sages, Distilled Wisdom, Brilliant oils. Wands: Lesser Magic Wand level 5, Greater 13, Elven Wand 7, Cinder Wand 9. Truefaith Vestments bind on pickup and sit in every 60 raid healer set. Dormant Heart of the Mountain binds on pickup and is in the Horde Discipline raid set only, as the page says.
- **Trainer levels and leveling text** in all three "Levels 1 to 19" paragraphs and the "Levels 20 to 60" paragraphs match the orders and `trainables`. Every spell a leveling line tells the player to use is available by the line's level.
- **Source links**: Blizzard recap, Talents Forever data.json, Power Up Gaming, Forever Sixty weights and planner all return 200. Icy Veins returns 403 (treated as unknown, per the brief). No priest page or rendered note names 70755.

## Decisions

- **D1. R23-1 is CONFIRMED, not a Should.** The druid report filed a similar omission ("whenever he is under 75%") as a Should, because that line only left out a gate. Discipline's line goes further: "there is nothing to hold back" states the opposite of a gate the card has, and two conditions are missing.
- **D2. R23-3 is CONFIRMED on the same reasoning as DRUID-R23-2's D2:** an "every heal" stat line that the page's own Low ranks paragraph contradicts, with a rank on the card that keeps only part of the bonus. Shadow's line stays a Should because that page states no limit and its 60 card casts top ranks.
- **D3. R23-4 is LIKELY.** The numbers are hard (Flash Heal is smaller at every band). Whether "size" reads as "the Flash Heal's rank" or "the spell" is judgment. Take it: the fix also retires R22-S3.
- **D4. R23-5 is LIKELY on the count and left to a sim on "cheapest".** "Up to" is the client's word, so that half costs nothing to fix. Round 22's D1 stands: "all fight" excludes the 10 minute spells.
- **D5. Whole-row cases stay out of findings.** The Inner Focus row at 20 to 56 and the Mind Flay row at 20 to 26 render before the talent, and those card actions are accepted by the brief. The text on them is conditional or true once castable, except R23-2's clause, which names Mind Flay on a different row.

## Sims requested

1. **Holy, raid 60, shipped build and BiS, heal profile, 10,000 iterations, seeds 1 and 50001:** mean Prayer of Mending heals per cast with the shipped "every time it is ready" row, and healing per mana by spell against Renew rank 7. Decides "cheapest" (R23-5): below about 3 heals a cast it falls behind Renew. Repeat at bands 40 and 50.
2. **Shadow, raid 60:** the Shadow Word: Pain row at `remainingTime >= 6` against the shipped 10 (decides whether the card should match shadow.md:89; R23-S2).
3. **Discipline, raid 60 (optional):** the shield with no health gate against the shipped 90% gate (a shield before the hit, as "soaks the next hit" implies).
4. **Holy, raid 60 (optional):** Greater Heal rank 1 instead of the top rank on rows [15] to [19] when behind the mana pace (R23-S12).
5. Carried (optional): Shadow Word: Death ahead of Mind Blast at 32 to 51; Discipline Gnome against Human at the re-rank; every 0.1% to 0.5% race row at the re-rank.

## Engine questions

- **Prayer of Mending's jump target.** The engine jumps to the most hurt other raid member (`mostHurtOther`); the client says "a party or raid member within 20 yards", with no ranking stated. The choice sets how fast charges spend, so it bears on sim 1.
- Known and not re-filed: Lightwell has no spell-power coefficient in the engine and clicks below 90% health. No page claim depends on it.

## Sim findings

None new. Still open and not re-filed: P2 target-level healing, P3 Eureka! on a live Gnome Shadow log, P4 Lightwell Renew's crit, S2 Dark Inspiration's Spirit base.

## Sources

- Client data 1.60.1.70790: `spelltips/priest.json`, `spelltips/shared.json`, `talents/priest.json`, `trainables/priest.json`, `addon-builds.json`, `addon-data.json`, `bis/priest-*.json` (band healing summed from `items/*.json`: Discipline Human 154, 234, 363, 464, 490 bare at 20 to 60 and 576 raid), `races.json`, `combos.json`, `itemnames.json`.
- Curated: `data/curated/apl/priest-{discipline,holy,shadow}.json`, `data/curated/races.json`.
- Engine d5c708834: `sim/core/spell.go` (`SpellLevelPenalty`), `sim/core/spell_max_level_auto_gen.go`, `sim/core/eureka_auto_gen.go`, `sim/core/racials.go`, `sim/priest/{heal_ranks,prayer_of_mending,healing_talents,racials}.go`.
- Reviews: round 22 priest, ledger Round 19 to 22, round 23 druid, mage, paladin, shaman, warrior, sim round 13 to 15 healer notes (the low-rank rule, W1).
- No web research needed: the client data decided every call.

## Appendix: Note sweep

One line per row note as `addon-data.json` renders it, every sentence of the curated note, every band, both orders. Levels are when the player has each named item on the shipped order or from the trainer. Ranks per band are the rendered rows.

### Discipline (bands 10, 12, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 45, 48, 50, 54, 56, 58, 59, 60; raid order equals leveling)

- **Power Word: Shield** (10 to 60; ranks 1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 9, 9, 10): top rank at every band. First sentence "whenever Weakened Soul is off him" against the condition: **R23-1**. "costs a warrior no rage" (R10). Renewed Hope 32 and Divine Aegis 36, both "once you have". 
- **Renew** (10 to 60; ranks 1 to 28, 2 at 30 to 32, 3 at 34 to 38, 4 at 40 to 42, Human 7 and others 4 at 44, Human 7 and others 5 at 45, 5 at 48, 7 from 50, 9 for four races at 56 to 59): "at 60 that means rank 7 or higher". OK (R23-S4 on prose).
- **Lesser Heal rank 3** (10 to 38): "small top-up for little mana before 40" (75); "by your late forties it has lost your healing bonus" (none from 47); "costs regeneration". OK.
- **Lesser Heal** (party rows without a note, 10 to 38): nothing to check.
- **Heal** tank (16 to 60; ranks 1, 1, 1, 2, 2, 2, 3, 3, 3, then 4): "until you learn Greater Heal at 40"; "your biggest heal until then". Self-qualified. OK (R23-S6).
- **Heal** (no note): nothing to check.
- **Dark Sacrifice** (20 to 60; ranks 1 to 5 at 20, 30, 40, 50, 60): under 70% matches the condition; 15 s; Spirit; 10 min. OK.
- **Flash Heal** tank (20 to 60; top rank at every band): "the fast heal, and the dear one". OK.
- **Flash Heal** party (20 to 60): "as long as your mana is not far behind" (0.7 gate). OK.
- **Flash Heal** (no note, three party rows): nothing to check.
- **Penance** tank (30 to 60; ranks 1 at 30, 2 at 40, 3 at 50, 4 at 60): 12 s; three pulses; "on whoever is hurt" matches the conditions; "cheap for all it heals" (3.6 a point of mana at 30 on the band preset, about 5 at 60). OK.
- **Penance** (no note, party lines): nothing to check.
- **Prayer of Healing** (30 to 60; ranks 1, 3 from 50, 4 at 60): "three of your four party members under 85%"; "target and whole party"; "Rank 4 is the top rank a trainer teaches" (rank 5 is the book rank, `book-ranks.json`). OK.
- **Greater Heal** tank (40 to 60; rank 1 to 56, rank 4 from 58): "at the top rank ... at 60". OK.
- **Greater Heal rank 1** party (40 to 60): "once you have higher ranks" (46); keeps most (whole to 57, 85% at 60); far less mana. OK.
- **Greater Heal** (no note): nothing to check.
- **Greater Heal** deeper hole tank (40 to 60): "a tank under 50%". OK.
- **Greater Heal rank 1** deeper hole (40 to 60): "a party member under 35%". OK.

### Holy (bands 10, 16, 20, 22, 26, 28, 30, 32, 33, 34, 38, 40, 44, 45, 50, 51, 52, 53, 56, 58, 59, 60; raid order equals leveling)

- **Renew** (10 to 60): "at 60 that means rank 7 or higher". OK.
- **Heal** (16 to 60; top rank to 38, then rank 4): "top rank before Greater Heal at 40, rank 4 after"; "cheaper"; "about half your bonus at 60" (55%). OK.
- **Heal** (no note): nothing to check.
- **Heal** deeper hole (16 to 60): thresholds match. OK.
- **Dark Sacrifice** (20 to 60): as Discipline. OK.
- **Flash Heal** tank (20 to 60; top rank): "at the top rank" OK; "your dearest heal, so spend it where the size matters": **R23-4**.
- **Flash Heal** party (20 to 60): "the size matters more than the mana": **R23-4**.
- **Flash Heal** (no note): nothing to check.
- **Prayer of Healing** (30 to 60; rank 1 to 59, rank 4 at 60): as Discipline's first sentence. OK.
- **Prayer of Mending** (40 to 60; ranks 1, 2 from 50, 3 at 60): "every time it is ready" (unconditional row) OK; "five heals for one cast": **R23-5**; 10 s OK; "cheapest ... all fight" OK at the full count, sim 1; "next hit or direct heal" OK; "the tank is hit and healed the most" OK; "not held back by the mana pace" (no gate) OK.
- **Greater Heal** filler (40 to 60; rank 1 to 56, 4 at 58, 3 at 59, 4 at 60): "rank 1 until 57"; "rank 3 or your top trainer rank ... about the same for their mana". OK.
- **Greater Heal** (no note): nothing to check.
- **Greater Heal** deeper hole (40 to 60): thresholds match. OK.
- **Greater Heal rank 1** (40 to 60): "once you have higher ranks" (46); "costs far less"; "most of your healing bonus at 60" (85%). OK (R23-S12).
- **Greater Heal rank 1** deeper hole (40 to 60): OK.

### Shadow (bands 10, 16, 18, 20, 22, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 51, 52, 58, 60; leveling and raid orders)

- **Shadow Word: Pain** (10 to 60; ranks 2, 2, 3, 3, 3, 4, 4, 3 at 30 to 38, 5, 6 at 42 to 48, 7 at 50 to 52, 8): "Always up"; "closing seconds" (10 s gate). OK (R23-S2).
- **Mind Blast**, Inner Focus line (10 to 60): Inner Focus 57 on both orders, "whenever it is up"; free; 25% on a non-periodic spell. OK as text.
- **Inner Focus** row (20 to 60): same note. Whole-row case (D5). OK as text.
- **Mind Blast**, second line (10 to 60): "ahead of Mind Flay" at 10 to 27 leveling and 10 to 20 raid: **R23-2**. "a big hit for its short cast" OK; "half cost in Shadowform" (42, conditional) OK.
- **Dark Sacrifice** (20 to 60): below a fifth and behind the fight matches the condition; 15 s; Spirit; 10 min; "Only Undead priests have it" (`race` trainables). OK.
- **Devouring Plague** (20 to 60): 24 s at every rank; "every priest has it"; ticks crit (owner rule); heals you. OK.
- **Mind Flay** (20 to 60): "with every free global" on a whole-row case at 20 to 26 (D5); "clipped channel loses the rest"; Improved Mind Flay 29 / 22, conditional. OK.
- **Vampiric Embrace** (30 to 60): talent 33 / 28; 30 s; 20%; "Skip it in the last 10 seconds, and until you have the talent" matches the condition. OK.
- **Shadow Word: Death** (32 to 60; ranks 1, 2 at 40 to 46, 3 at 48, 2 at 50 to 58, 4 at 60): "ahead of Mind Flay" (from 32, after 28 and 21) OK; 30% (R23-S7); Vampiric Embrace conditional. OK.
- **Shoot** (curated note, no rendered row): no rank, talent or value. OK.

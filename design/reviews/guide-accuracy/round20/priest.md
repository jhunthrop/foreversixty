# Guide accuracy round 20: Priest

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20, engine pin d5c708834), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`. Files: `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`; every note in `data/curated/apl/priest-{discipline,holy,shadow}.json` (full text, every sentence); every rendered row of `addon-data.json` (`rotations/priest-discipline` 25 bands, `priest-holy` 22 bands, `priest-shadow` 22 bands, read band by band with spell ids and `race_spell_ids`); the shipped orders in `addon-builds.json` (leveling and raid, decoded against `talents/priest.json`, index 0 = level 10); `trainables/priest.json`, `spelltips/priest.json`, `spelltips/shared.json`. Engine at d5c708834: `sim/core/spell.go` (`SpellLevelPenalty`: MaxLevel + 12 grace, 5% a level), `sim/core/spell_max_level_auto_gen.go`, `sim/core/eureka_auto_gen.go` (the priest row 1259823). Read first: round 19 priest, sim round 20 `priest.md` and `integration.md`, the ledger's Round 19 section.

## Changes since round 19's tree (6beb919c2)

- Round 19's fixes: the Shadow Word: Death note ("once you have Vampiric Embrace"), the Holy Greater Heal rank 1 note (now "Once you have higher ranks, when the top rank would overheal ..."), holy.md:122 (Renew rank 1 through 32, rank 3 from 33).
- Sim round 20 (P-LH, P-DP, D1): the Discipline Lesser Heal gate is `currentLevel < 40` on all eight lines and its note is new ("Lesser Heal is the party's small top-up for little mana before 40. From 40 on, leave it off your bars for good: ..."); discipline.md:72 ("Greater Heal rank 1 on everyone else"), :74 (Lesser Heal off the bars for good), :100 (Gnome line), :124 (Renew rank 9 clause, Lesser Heal from 40); a Shadow Devouring Plague rank 1 row at 28 and 29 for Gnome and Troll (`race_spell_ids` gnome and troll 2944 at band 28); the engine's Renewed Hope cut is per cast.
- `addon-data.json` carries all of these (the Lesser Heal row renders 10 to 38 only; the band-28 Devouring Plague row has the Gnome and Troll overrides).

## Builds and orders

Decoded from `addon-builds.json`. Discipline (leveling and raid identical): Twin Disciplines 10-14, Mental Strength 15-19, Meditation 20, Mental Agility 21-23, Inner Focus 24, Improved Power Word: Shield 25-27, Meditation 28-29, Penance 30, Soul Warding 31, Renewed Hope 32-35, Divine Aegis 36-38, Twilight Focus 39-41, Improved Renew 42-44, Holy Specialization 45-49, Divine Fury 50-54, Inspiration 55-57, Improved Healing 58-60. Holy (identical): Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization 16-20, Divine Fury 21-25, Inspiration 26-28, Improved Healing 29-31, Holy Binding 32, Litany of Light 33-34, Spiritual Healing 35-37, Spirit of Redemption 38, Spiritual Guidance 39 and 41-43, Prayer of Mending 40, then Discipline, Inner Focus 57. Shadow leveling: Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus 17-21, Improved Mind Blast 22-25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Vampiric Embrace 33, Shadow Weaving 34-36, Darkness 37-41, Shadowform 42, Inner Focus 57. Shadow raid: Mind Flay 21, Improved Mind Flay 22-23, Vampiric Embrace 28, Shadowform 42, Inner Focus 57.

Trainer levels used: Shadow Word: Death 32 / 40 / 48 / 56 (295 / 370 / 403 / 448); Mind Blast 10 / 16 / 22 / 28 / 34 / 40 / 46 / 52 / 58 (39 / 68 / 101 / 151 / 196 / 258 / 324 / 405 / 485, all 1.5 s casts); Devouring Plague 20 / 28 / 36 / 44 / 52 / 60; Renew 8 to 56 by sixes; Heal 16 / 22 / 28 / 34; Greater Heal 40 / 46 / 52 / 58; Flash Heal 20 / 26 / 32 / 38 / 44 / 50 / 56; Penance 30 / 40 / 50 / 60; Prayer of Mending 40 / 50 / 60; Lesser Heal rank 3 at 10; Dark Sacrifice 20 / 30 / 40 / 50 / 60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R20-1 | CONFIRMED | card note (superlative by band) | `priest-shadow.json` priorityList[6], the second Mind Blast note, rendered on the Mind Blast row at Shadow bands 10 to 60; false at 32, 34, 36, 38, 40, 42, 44, 46, 48, 50 and 51 | "Mind Blast on cooldown, ahead of Mind Flay: your biggest hit for its cast time, at half cost in Shadowform." | From 32 the same card casts Shadow Word: Death (row from band 32), an instant that costs the same 1.5 s global Mind Blast's cast takes and hits far harder at every band to 51 on the rows the card shows: 32 rank 1 295 against Mind Blast rank 4 151; 34 to 38 295 against rank 5 196; 40 to 46 rank 2 370 against rank 6 258 (Human and the other overridden races rank 4, 151, at 40 and 42) and rank 7 324 at 46; 48 rank 3 403 against 324; 50 and 51 rank 2 370 against 324 (Undead rank 6, 258). Both carry the same 0.429 bonus coefficient (sim round 20 S1: Mind Blast 485 + 0.429, Shadow Word: Death 448 + 0.429) and the same Shadow talents, so gear never closes the gap. Mind Blast is the bigger hit only before 32 and from 52 (rank 8 405 against Shadow Word: Death rank 2 370; 485 against 448 at 60). The note is one sentence, so the addon row and the band cards carry all of it. Same shape as round 19's by-band rule: a superlative true at 60 quoted on bands where the rows say otherwise. | "Mind Blast on cooldown, ahead of Mind Flay: a big hit for its short cast, at half cost in Shadowform." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. The fix needs `addon-data.json` and Data.lua regenerated (note text only). No prose line on any of the four pages calls Mind Blast the biggest hit.

## Shoulds (wording, not findings)

- **PRIEST-R20-S1, shadow.md:137 and the guides-handoff (Devouring Plague rank 1 at 28 and 29).** The page reads "Devouring Plague rank 1 from 30 to 43, then rank 2 from 44 to 59; a Human drops back to rank 1 at 50 and 51." For the page's Human and the Horde pick Undead the card casts rank 2 at 28 (19276, learned that level) and rank 1 from 30 to 42, so the sentence is true for both picks. Only a Gnome or a Troll (`race_spell_ids` gnome and troll 2944 at band 28, sim round 20 P-DP, +0.2% to +0.6%) starts rank 1 at 28. The card note names no rank, so it is right for every race. Not filed (round 17 D6, round 18 S1/S2: race exceptions stay in the rows when the page's race matches its card). Optional, matching the Mind Blast clause beside it: "Devouring Plague rank 1 from 30 to 43 (a Gnome or Troll from 28), then rank 2 ..."
- **PRIEST-R20-S2, shadow.md:85 and the Shadow Word: Death note.** "Every cast that does not kill the target hits you for 30% of your maximum health" is the client tooltip word for word ("If your target is not killed by Shadow Word: Death, you take backlash damage equal to 30% of your maximum health", 1309595 to 1309636), so it is not filed. Blizzard's 3 November note (sim round 20 S1) says a miss takes no backlash. Optional: "every cast that lands and does not kill".
- **Penance "your best heal per mana"** (R19 Should, carried): true against every direct heal at every band from 30 (rank 1 126 x 3 for 150 = 2.52 a point of mana; rank 2 2.73; rank 3 2.98); against the tank's Renew at 60 it depends on healing gear.
- Carried, unchanged and optional: R18-S1 (Discipline Renew rows for Human and Undead, including "From 56, take Renew rank 9" against an Undead's rank 7 at 56 to 58), R18-S2, R18-S3, R18-S4 (Inspiration missing from both talent lists), R18-S5 / holy.md:122 "keep all or nearly all" (rank 6 keeps 80% at 59), R17-S2 ("well ahead of Troll"), and the rest listed in round 19.

## Checked and right this round

- **Sim round 20's prose, line by line.** discipline.md:72 "Greater Heal rank 1 on everyone else": the party filler rows at 60 are 2060 (priorityList 17, 19, 21, 23) and the deeper-hole party rows 2060; no rank 2 line. discipline.md:74: Lesser Heal rank 3 keeps 0% at 60 (MaxLevel 15 by the +5 rule, 12 grace, 5% a level: 85% at 30, 35% at 40, 0 from 47); Renew rank 5 45%, rank 6 75%, Heal rank 4 55% at 60, "only part of it"; "keep Lesser Heal off your bars for good" matches the card (no Lesser Heal row from 40). discipline.md:100 Gnome: the priest Eureka! row 1259823 puts Penance in both the +10% and the -10% cost masks, so "and so are Penance's heals" (cheaper and stronger) is right, and Renew is in neither. discipline.md:102 Human "level with Gnome, a hair ahead on a long tank fight": sim round 20 P-RACE (raid 60 +0.1%, bare 60 +0.7%). discipline.md:124: Renew rank 4 at 40 and 42 (6076, 100% to 95% of the bonus), rank 5 from 45 for the page's Gnome (6077, whole to 49), rank 7 from 50 (10927), Gnome rank 9 at 56 to 59 (10929) and rank 7 at 60; Greater Heal rank 1 on the tank 40 to 56 at 100% (whole to 57), rank 4 (10965) from 58; "by your late forties Lesser Heal has lost your healing bonus" (0 from 47); "Greater Heal rank 1 takes its place" (the party rows from 40 are 2060).
- **Round 19's fixes.** Shadow Word: Death note at 32 (Vampiric Embrace now conditional; the talent at 33 on the leveling order, 28 on the raid order). Holy Greater Heal rank 1 note "Once you have higher ranks" at 40 to 45 (only rank 1 learned) and from 46. holy.md:122: Holy Renew rows 139 at 10 to 32 (rank 1 keeps 75% at 30 and 65% at 32, "most"), 6075 at 33 to 38, 6076 at 40 and 44, 6077 at 45 and 50, Human 6078 from 52.
- **Every other prose line** of index.md, discipline.md, holy.md and shadow.md against the decoded orders, the trainer levels, the talent and spell descriptions above and the rendered rows: the Levels 1 to 19 paragraphs (Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Shadow Word: Pain 4, Psychic Scream 14, Dispel Magic 18, Holy Fire 20; Twin Disciplines 10-14, Mental Strength 15-19; Twilight Focus 10-12, Improved Renew 13-15, Holy Specialization from 16; Wand Specialization 10-11, Spirit Tap 12-16, Shadow Focus from 17); the Shadow 20 to 60 order (Shadow Focus done at 21, Improved Mind Blast to 25, Improved Shadow Word: Pain 26-27, Mind Flay 28, Improved Mind Flay 29-30, Blackout 31-32, Shadowform 42); the shadow.md:137 rank list against the Shadow rows (Shadow Word: Pain 970 at 30 to 38, Mind Blast 8104 overrides at 40 and 42 with no Gnome override, Shadow Word: Death 1309633 at 50 to 58); talent bullets (Renewed Hope 4 points 8% and 4 s; Divine Aegis 15%; Soul Warding removes the 4 s cooldown; Improved Power Word: Shield 20%; Blackout 2 points 4%; Dark Inspiration 1 point 3%; Early Demise 30%; Improved Mind Blast 0.5 s a point; Twilight Focus 70%; Litany of Light; Holy Binding 50%; Spiritual Healing 10%; Spirit of Redemption 15 s); the Power Infusion swap (Discipline 29 + Renewed Hope's fifth point = 30 for tier 7; Divine Aegis is the prerequisite); Lightwell 3 min or 5 charges; Fade 10 s on 30 s; Fear Ward; Divine Grace 10 min, "cannot be cast on self"; Desperate Prayer cost 0; Shadowform's +100% crit bonus making a Shadow crit twice a normal hit; the Eureka! lines (Mind Flay and Penance in the periodic mask; Shadow Word: Pain, Devouring Plague, Renew and Power Word: Shield in none). No em or en dashes; no site or sim talk outside the planner and weights links.
- **Card notes:** every row note at every band, in the appendix.

## Decisions

- **D1. PRIEST-R20-1 is CONFIRMED.** The comparison is per cast on the client rows with equal coefficients, on the ranks the card itself casts at each band; no gear or talent changes the order before 52. The fix drops the superlative instead of adding a band clause, so it holds at every band the row renders.
- **D2. The Inner Focus and Shadowform clauses on the two Mind Blast notes stay accepted.** "with Inner Focus first whenever it is up" (Inner Focus at 57 on both Shadow orders; rendered from band 10) and "at half cost in Shadowform" (42) are conditional on the player having the ability, as rounds 18 and 19 ruled; the same shape as the mage "with Arcane Power once you have it". The Inner Focus row itself rendering from band 20 is the component's whole-row case (round 19). Not re-filed.
- **D3. The Devouring Plague handoff (sim round 20 P-DP) is not a guide finding.** The card note names no rank; the page's rank sentence is true for both recommended races; the Gnome and Troll exception lives in the rows (R20-S1).
- **D4. The Discipline and Holy Heal rows rendered at 40 to 60** ("Heal on the tank until you learn Greater Heal at 40") are gated `not spellIsKnown(Greater Heal)` in the curated list and never fire there; the note qualifies itself. Component whole-row case, not filed.

## Sims requested

- Optional, for the card order (not the note): **Shadow Word: Death ahead of Mind Blast at 32 to 51** (the two on-cooldown lines swapped on bands where Shadow Word: Death is the bigger hit; ladder leveling build at each band, 10,000 iterations, seeds 1 and 50001, Human and Undead). Both are cast on cooldown, so the gain is only on collisions; the note fix does not wait for it.
- Carried: Discipline Gnome against Human at the re-ranked BiS; the Devouring Plague 28-29 row on re-ranked band-20 gear; every 0.1% to 0.5% race row at the re-rank.

## Engine questions

- Shadow Word: Death backlash on a miss: the engine follows Blizzard's 3 November note (no backlash on a miss); the client tooltip still reads "If your target is not killed". No guide change needed either way (R20-S2); for the sim round 21 integrator only if a log shows otherwise.

## Sim findings

None new. Still open from sim round 20 (not re-filed): P2 target-level healing, P3 Eureka! on a live Gnome Shadow log, P4 Lightwell Renew's crit, S2 Dark Inspiration's Spirit base.

## Sources

- Client data 1.60.1.70790: `talents/priest.json`, `spelltips/priest.json` (8092, 8102 to 8106, 10945 to 10947, 1309595, 1309633, 1309635, 1309636, 2944, 19276, 19280, 14751, 15286, 15407, 17311, 18807, 594, 970, 992, 2767, 10892, 10894, 1277324 to 1277328, 401859, 1240826, 1240827, 402174, 1240720, 1240721, 1316995, 2053, 2054, 2055, 6063, 6064, 2060, 10963, 10964, 10965, 139, 6074 to 6078, 10927, 10929, 2061, 9472 to 9474, 10915 to 10917, 596, 996, 10960, 10961, 17, 10901, 10060, 724, 27870, 27871, 6346, 10942, 586, 13908 to 19243, 1277370 to 1277378, 1277462), `spelltips/shared.json`, `trainables/priest.json`, `addon-builds.json`, `addon-data.json` (`rotations/priest-*`, every band).
- Curated: `data/curated/apl/priest-{discipline,holy,shadow}.json` (notes and conditions).
- Engine d5c708834: `sim/core/spell.go` (SpellLevelPenalty, constants 12 / 0.05 / 5), `sim/core/spell_max_level_auto_gen.go`, `sim/core/eureka_auto_gen.go` (ClassPriest, 1259823).
- Web: `web/src/lib/guides/rotation-lists.ts` (clauseOf), `web/src/lib/bis/rotation-view.ts` (alignNoteRank, band rows).
- Reviews: round19/priest.md, round18/priest.md (S1 to S5), round1/priest.md (the 6 s Shadow Word: Pain cutoff), design/reviews/2026-11-04-sim-audit-round20/priest.md and integration.md, ledger.md (Round 19).

## Appendix: Note sweep

One line per row note as `addon-data.json` renders it (first sentence; later sentences of the curated note checked too and named where they exist). Levels are when the player has each named item on the shipped orders or from the trainer.

### Discipline (bands 10 to 60)

- **Power Word: Shield** (10 to 60): "On the tank at the top rank whenever Weakened Soul is off him ..." Top rank at every band (17 at 10, 592 at 12-16, 600 at 18-22, 3747 at 24-28, 6065 at 30-34, 6066 at 36-40, 10898 at 42-45, 10899 at 48-50, 10900 at 54-59, 10901 at 60); Weakened Soul from rank 1. Later sentence: Renewed Hope (32) and Divine Aegis (36) both "once you have" / "with"; Divine Aegis "non-periodic critical heals" covers every Renewed Hope heal. OK.
- **Renew** (10 to 60): "Keep Renew on the tank all fight; at 60 that means rank 7 or higher." Renew from 8; 10927 at 60. OK.
- **Lesser Heal** (10 to 38): "Lesser Heal is the party's small top-up for little mana before 40." Rank 3 from 10, 75 mana; the row ends at 38 (gate below 40). Later sentence: "by your late forties it has lost your healing bonus" (0 from 47). OK.
- **Lesser Heal** (no note, the other party and deeper-hole lines, 10 to 38): nothing to check.
- **Heal** tank (16 to 60): "Heal on the tank until you learn Greater Heal at 40: it is your biggest heal until then." 270 / 375 / 523 / 651 at 16 / 22 / 28 / 34 against Flash Heal 194 to 383, Lesser Heal 163, Penance 3 x 126; never fires from 40 (D4). OK.
- **Heal** (no note, the deeper-hole line, 16 to 60): nothing to check.
- **Dark Sacrifice** (20 to 60): "Undead only, once you are under 70% mana and still healthy: ... over 15 seconds, plus your Spirit, ... 10 minute cooldown". Ranks 1 to 5 all 15 s, "+$SPI", 600 s; condition below 70% mana with health above 70%. OK.
- **Flash Heal** tank (20 to 60): "at the top rank on a tank who is really low: the fast heal, and the dear one." Top rank at every band (2061, 9472 from 26, 9473 from 32, 9474 from 38, 10915 from 44, 10916 from 50, 10917 from 56). OK.
- **Flash Heal** party (20 to 60): "on a party member about to die, as long as your mana is not far behind the fight." No rank or talent. OK.
- **Flash Heal** (no note, 20 to 60): nothing to check.
- **Penance** (30 to 60): "on whoever is hurt: three heals one second apart for one cast every 12 seconds, your best heal per mana." Talent at 30; every rank 12 s and "instantly and every 1 sec for 2 sec"; per mana see Shoulds. OK.
- **Penance** (no note, 30 to 60): nothing to check.
- **Prayer of Healing** (30 to 60): "once three of your four party members are under 85%: it heals the target and the target's whole party." Trainer 30. Later sentence "Rank 4 is the top rank a trainer teaches" (10961 at 60; rank 5 is the book). OK.
- **Greater Heal** tank top rank (40 to 60): "Greater Heal at the top rank as the tank's filler at 60 ..." Qualified "at 60"; 10965 from 58. OK.
- **Greater Heal rank 1** party (40 to 60): "... once you have higher ranks it keeps most of your healing bonus for far less mana." Rank 2 at 46; rank 1 100% to 57, 85% at 60; 370 against 455 / 545 / 655. OK.
- **Greater Heal** (no note, 40 to 60): nothing to check.
- **Greater Heal** deeper hole (40 to 60): "waits for a deeper hole once your mana falls behind the fight: a tank under 50%." Condition tank below 50%. OK.
- **Greater Heal rank 1** deeper hole (40 to 60): "... a party member under 35%." Condition below 35%. OK.

### Holy (bands 10 to 60)

- **Renew** (10 to 60): "Keep Renew on the tank all fight; at 60 that means rank 7 or higher." Renew from 8; 10927 at 60. OK.
- **Heal** (16 to 60): "Heal at your top rank before you learn Greater Heal at 40, and rank 4 after that when even Greater Heal rank 1 is out of reach: cheaper, though it keeps only about half your bonus at 60." Top rank at 16 to 38 (2054, 2055 from 22, 6063 from 28, 6064 from 34); 6064 from 40; 305 against 370 mana; 55% at 60. OK.
- **Heal** (no note, 16 to 60): nothing to check.
- **Heal** deeper hole (16 to 60): "waits for a deeper hole once your mana falls behind the fight ..." No rank or talent. OK.
- **Dark Sacrifice** (20 to 60): as Discipline. OK.
- **Flash Heal** tank (20 to 60): "at the top rank on a tank who is really low or on anyone about to die: it is your dearest heal ..." Top rank at every band; per mana below Heal, Greater Heal (Improved Healing from 29) and Prayer of Mending at every band. OK.
- **Flash Heal** party (20 to 60): "at the top rank on a party member under 30% ..." Top rank at every band; condition below 30%. OK.
- **Flash Heal** (no note, 20 to 60): nothing to check.
- **Prayer of Healing** (30 to 60): as Discipline (first sentence). OK.
- **Prayer of Mending** (40 to 60): "on the tank every time it is ready: five heals for one cast on a 10 second cooldown, the cheapest healing you have." Talent at 40; ranks 1 to 3 all "Jumps up to 5 times", 10 s; 172 x 5 for 210 at 40. Later sentences: "next hit or direct heal" (non-periodic healing), not paced. OK.
- **Greater Heal** filler (40 to 60): "rank 1 until 57, then rank 3 or your top trainer rank, which heal about the same for their mana." 2060 at 40 to 56, 10965 at 58 and 60, 10964 at 59; 1403 / 545 against 1782 / 655. OK.
- **Greater Heal** (no note, 40 to 60): nothing to check.
- **Greater Heal** deeper hole (40 to 60): "waits for a deeper hole ... a tank under 50% or a party member under 35%." OK.
- **Greater Heal rank 1** (40 to 60): "Once you have higher ranks, when the top rank would overheal or your mana runs short: it costs far less and still keeps most of your healing bonus at 60." Rank 2 at 46; 85% at 60. OK.
- **Greater Heal rank 1** deeper hole (40 to 60): "waits for a deeper hole too once your mana falls behind the fight." OK.

### Shadow (bands 10 to 60)

- **Shadow Word: Pain** (10 to 60): "Always up; stop refreshing it in the closing seconds of a fight." Trainer 4. OK.
- **Mind Blast**, Inner Focus line (10 to 60): "Mind Blast on cooldown, with Inner Focus first whenever it is up: Inner Focus makes that Mind Blast free, with 25% more crit chance." Mind Blast from 10; Inner Focus 57 on both orders, conditional (D2); "25% ... if it is a non-periodic spell". OK.
- **Inner Focus** row (20 to 60), same note: the row before 57 is the component's whole-row case (D2). OK as text.
- **Mind Blast**, second line (10 to 60): "Mind Blast on cooldown, ahead of Mind Flay: your biggest hit for its cast time, at half cost in Shadowform." Shadowform 42, conditional. "your biggest hit": false at 32 to 51 against Shadow Word: Death. **PRIEST-R20-1.**
- **Dark Sacrifice** (20 to 60): "(Undead) when mana is below a fifth and running behind the fight: ... over fifteen seconds and adds your Spirit ... ten minute cooldown". Condition below 20% and behind pace; every rank 15 s, 600 s. Later sentence "Only Undead priests have it." OK.
- **Devouring Plague** (20 to 60): "every time its 24 second cooldown is up: every priest has it in Forever, its ticks crit, and it heals you." Trainer 20; every rank 24 s and "heals the caster"; DoTs crit (owner rule); no rank named, so the band-28 Gnome and Troll rank 1 override reads right. OK.
- **Mind Flay** (20 to 60): "let all three ticks land, since a clipped channel loses the rest, and once you have Improved Mind Flay the last tick is the biggest." Talent 28 (leveling) / 21 (raid), rank 2 trains at 28; Improved Mind Flay 29 / 22, conditional; rows at 20 to 26 are the whole-row case. OK.
- **Vampiric Embrace** (30 to 60): "Vampiric Embrace on the boss whenever it is down." Talent 33 / 28; rows at 30 and 32 the whole-row case. Later sentences: 30 s cooldown equal to duration (15286: 30 s, "for 30 sec"), 20%, "until you have the talent". OK.
- **Shadow Word: Death** (32 to 60): "On cooldown, ahead of Mind Flay: each cast that does not kill costs you 30% of your health, and once you have Vampiric Embrace keep it up to heal that back." Trainer 32; tooltip 30% on every rank; Vampiric Embrace conditional (33 / 28). OK (see R20-S2).

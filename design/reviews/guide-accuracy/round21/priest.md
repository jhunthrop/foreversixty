# Guide accuracy round 21: Priest

Audited on the round 21 audit tree (e039e1490: main after guide round 20; no sim round 21 commit in the log, engine pin d5c708834 in `sim/enginever/version.go`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`. Files: `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`; every note in `data/curated/apl/priest-{discipline,holy,shadow}.json`; every rendered row of `addon-data.json` (`rotations/priest-discipline` 25 bands, `priest-holy` 22 bands, `priest-shadow` 22 bands, with `race_spell_ids`); `addon-builds.json` (orders as decoded in round 20, unchanged); `trainables/priest.json`, `spelltips/priest.json`, `talents/priest.json`, `races.json`, `bis/priest-discipline.md` (band 60 raid preset items) and `items/*.json` for their healing. Engine at d5c708834: `sim/priest/lightwell.go`, `penance.go`, `renew.go`, `heal_ranks.go`. Read first: round 20 priest, ledger Round 20 section.

## Changes since round 20's tree (4234643f5)

- One priest change: round 20's fix of PRIEST-R20-1. The second Shadow Mind Blast note now reads "Mind Blast on cooldown, ahead of Mind Flay: a big hit for its short cast, at half cost in Shadowform." (`priest-shadow.json` priorityList[6]; `addon-data.json` carries it at every Shadow band). No priest guide prose changed.
- Re-check of the fix: no superlative remains. "a big hit for its short cast" holds at every band 10 to 60 (Mind Blast ranks 1 to 9, 39 to 485, all 1.5 s casts with the 0.429 coefficient). "at half cost in Shadowform" is conditional on the form (42 on both Shadow orders), as round 20 D2 ruled. OK.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R21-1 | CONFIRMED | card note (superlative against another spell) | `priest-holy.json` priorityList[2], the Prayer of Mending note, rendered on the Prayer of Mending row at Holy bands 40 to 60 (first and third sentences) | "five heals for one cast on a 10 second cooldown, the cheapest healing you have." ... "the cheapest healing you have is the last thing to give up when mana runs short." | Lightwell is a trainer spell from 40, learned with Prayer of Mending (`trainables/priest.json`: 724 at 40, 27870 at 50, 27871 at 60), the same card casts it before the pull (prepullActions[0], 27871 rank 3), and holy.md:66 tells the player to drop it every pull. Client: Lightwell rank 1 225 mana for 5 charges of 800 health (4,000), rank 2 295 for 5 x 1,165 (5,825), rank 3 365 for 5 x 1,600 (8,000); the engine reads the same ticks (`lightwell.go`: 160 / 233 / 320 every 2 s for 10 s). That is 17.8, 19.7 and 21.9 health a point of mana. Prayer of Mending on the rows: rank 1 (401859) 172 x 5 = 860 for 210 (178.5 with Improved Healing from 31) = 4.8 at 40; rank 2 (1240826) 298 x 5 for 305 (259) = 5.8 at 50; rank 3 (1240827) 413 x 5 for 390 (331.5) = 6.2 at 60. With the 0.429 coefficient each jump at 800 healing and Spiritual Healing, rank 3 reaches about 12.5 a point of mana, still well under Lightwell's 21.9 (13.2 with only three of the five charges clicked). False at every band the row renders. | "Prayer of Mending on the tank every time it is ready: five heals for one cast on a 10 second cooldown, the cheapest healing you can cast all fight. ... It is not held back by the mana pace the costlier heals follow: it is the last heal to give up when mana runs short." |
| PRIEST-R21-2 | CONFIRMED | prose (superlative against another spell and every class) | holy.md:30 (Strengths) | "Prayer of Mending is the cheapest healing in the game: five heals for one cast" | Same numbers as PRIEST-R21-1: the priest's own Lightwell heals about two to four times as much per point of mana as Prayer of Mending at every rank (17.8 / 19.7 / 21.9 against 4.8 / 5.8 / 6.2 base), so "in the game" fails before it ever reaches another class. | "Prayer of Mending is the cheapest healing you can cast all fight: five heals for one cast" |
| PRIEST-R21-3 | CONFIRMED | card note (superlative against another spell) | `priest-discipline.json` priorityList[2], the Penance note, rendered on the Penance row at Discipline bands 30 to 60; false from 40 | "three heals one second apart for one cast every 12 seconds, your best heal per mana." | (a) Lightwell from 40, cast by the same card before the pull (prepullActions[0], 27871) and on the page (discipline.md:64): 800 health a click for 225 mana at 40, so a single click (3.6 a point of mana) already beats Penance rank 2 (1240720: 3 x 200 = 600 x 1.05 Twin Disciplines for 220 = 2.9, about 3.1 with the 0.19-a-pulse coefficient on band-40 gear), and five clicks (17.8) beat it many times over; the same holds at 50 (rank 2 well, Penance rank 3 3 x 268 for 270) and 60 (rank 3 well 21.9, Penance rank 4 3 x 425 x 1.05 for 327 with Improved Healing = 4.1 base). (b) At 60 in the page's own raid set the tank's Renew rank 7 (the row's 10927) catches it: raid-preset items (`bis/priest-discipline.md`, band 60 raid) carry about 576 healing on the healing-only line alone; with Twin Disciplines, Improved Renew and Mental Agility, Renew rank 7 heals (370 + 1.0 x H) x 1.2075 for 225 and Penance rank 4 (425 + 0.19 x H) x 3 x 1.05 for 327 (`heal_ranks.go`: Renew 0.2 a tick x 5, Penance 0.19 a pulse): level at about H = 600, Renew ahead above it (6.4 against 5.6 at H = 830 with the set's spell power counted). True at 30 to 38 (Penance rank 1 2.6 against Lesser Heal rank 3 2.2, Renew rank 2 or 3 at most 1.75). | "Penance on cooldown, on whoever is hurt: three heals one second apart for one cast every 12 seconds, cheap for all it heals." |
| PRIEST-R21-4 | CONFIRMED | prose (superlative against another spell) | discipline.md:76 (Emergencies) | "Flash Heal at the top rank on anyone about to die. It is your fastest heal and your dearest" | Penance "heals ... instantly and every 1 sec for 2 sec" (402174 to 1316995); its first heal lands the moment you press it, and by one second it has healed more than Flash Heal heals at 1.5 s on the ranks the card casts: 40, Penance rank 2 400 against Flash Heal rank 4 383; 50, rank 3 536 against rank 6 644 at 1.5 s but 268 at once; 60, rank 4 850 by 1 s against rank 7 825 at 1.5 s. From 30 the page's own healer has a faster heal than Flash Heal (and the page's Human co-pick has the instant Divine Grace). "Dearest" is right at every band (checked in the sweep). The claim is true only as cast time. | "Flash Heal at the top rank on anyone about to die. It takes a second and a half and it is your dearest heal, so it never becomes the filler." |

Counts: CONFIRMED 4, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. R21-1 and R21-3 need `addon-data.json` and Data.lua regenerated (note text only). No other note or prose line fails.

## Shoulds (wording, not findings)

- **PRIEST-R21-S1, holy.md:80 "a Prayer of Healing, your most expensive heal".** Holy learns Inner Focus at 57; the Holy card casts Prayer of Healing rank 1 (596, 410 mana) at 50 to 59 and Greater Heal rank 4 (10965, 655, 557 with Improved Healing) at 58, rank 3 (10964, 463) at 59. At 58 and 59 the Prayer of Healing the rows cast is cheaper than the Greater Heal they cast. The Rotation section is the raid section and at 60 it is right (10961, 1030), so not filed. Optional: "on a top-rank Prayer of Healing, your most expensive heal".
- **PRIEST-R21-S2, holy.md:58 Twilight Focus "worth more than any throughput point".** An unmeasured comparison. Optional: "On raid-wide damage this keeps your heals on time."
- **PRIEST-R21-S3, shadow.md:137 "then back to the top rank".** The Horde pick Undead casts Mind Blast rank 6 (8106) at 50 and 51 where the default row is rank 7 (10945), and Dwarf at 51; the page's Human is on the top rank, so by round 17 D6 the race exception stays in the rows. Optional: "(an Undead keeps rank 6 at 50 and 51)".
- **PRIEST-R21-S4, the Discipline Flash Heal tank note "the fast heal, and the dear one".** Not a superlative, true at every band; if R21-4 is taken, "the quick cast, and the dear one" reads the same way.
- Carried, unchanged and optional: R20-S1 (Devouring Plague rank 1 at 28 and 29 for Gnome and Troll), R20-S2 (Shadow Word: Death backlash on a miss), R18-S1 to S5, R17-S2.

## Checked and right this round

- **Every superlative, comparison and quoted value on the four pages**, by band, on the ranks the rows cast:
  - Discipline Heal note "your biggest heal until then" (16 to 38): Heal 270 / 375 / 523 / 651 against Flash Heal 194 to 383, Lesser Heal 163, Penance rank 1 3 x 126, Prayer of Healing rank 1 178 a target, Renew 45 to 125, Power Word: Shield at most 291 x 1.2. OK.
  - "the dear one" / "your dearest heal" (Discipline and Holy Flash Heal, discipline.md:76, holy.md:74): Flash Heal 1.55 (rank 1) to 2.17 (rank 7) health a point of mana base, below Heal (1.74 to 2.13), Greater Heal (2.28 to 2.72 base, more with Improved Healing), Prayer of Mending, Lesser Heal rank 3 (2.17) and Prayer of Healing on a full party; above only Renew rank 1 for Holy at 20 to 32 (45 x 1.15 for 30 = 1.73 against Flash Heal rank 1 1.55): Flash Heal is the dearer, as the note says. OK.
  - discipline.md:82 "a Greater Heal crit leaves the biggest Divine Aegis you can make": Greater Heal is the largest single heal at every band from 40 (Divine Aegis cap 20 x 60). OK.
  - holy.md:70 "Heal rank 4 is cheaper still" (305 against 370); "lost a little of your healing bonus by 60" (Greater Heal rank 1 85%); "keeps only about half" (55%); "every one of them is a 2.5 second cast" (Heal and Greater Heal 3.0 s, Divine Fury 0.5 s). OK.
  - Holy Greater Heal note "rank 3 or your top trainer rank, which heal about the same for their mana" (1403 / 545 against 1782 / 655). OK.
  - shadow.md:36 Shadow Word: Death 18 s (every rank 18000); :58 Mind Blast 8 s and 0.5 s a point; :59 Improved Shadow Word: Pain 6 s; :69 and :99 Inner Focus 3 min; :81 and the note, Devouring Plague 24 s on every rank; :87 and the note, Vampiric Embrace 30 s cooldown and 30 s duration, 20%; :89 Shadow Word: Pain 18 s in six 3 s ticks; :95 Fade 10 s on 30 s on every rank; :71 Silent Resolve 10% a point; :66 Silence 5 s; :65 Devouring Contagion 50% and 2 targets in 10 yards; :62 Blackout 2% a point; :68 Dark Inspiration 3% a point; :54 Shadowform (and "you may not cast healing spells", so :43 and :141 hold, the shield being an absorb). OK.
  - holy.md:66 Lightwell "3 minutes or 5 clicks", :68 Weakened Soul 15 s (Holy has no Renewed Hope), :100 Divine Grace 10 min; discipline.md:60 Power Infusion 20% for 15 s on 3 min; discipline.md:104 and holy.md:102 "well over a thousand mana" (Dark Sacrifice rank 5 1600 + Spirit). OK.
  - Race superlatives ("the strongest Shadow priest of all", "the best Holy priest of all", "level with Gnome, a hair ahead", "close behind", "clearly less", "well ahead of Troll") against the measured race results recorded in round 4 and sim round 20 P-RACE; racial values against `races.json`. OK.
  - The "Levels 1 to 19" paragraphs on all three specs: no superlative fails; Twin Disciplines at 10 to 14 touches only Shadow Word: Pain and Renew among the priest's damaging or healing instants; Wand Specialization 25% at two points; Spirit Tap 100% for 15 s with half regeneration at five points; Twilight Focus 70% at three; Improved Renew 15% at three. OK.
- **Round 20's fix** (above). **Shadow raid order**: every Shadow note that names a talent is conditional ("once you have Improved Mind Flay", "once you have Vampiric Embrace", "until you have the talent", "in Shadowform", Inner Focus "whenever it is up"), so nothing is credited early on the raid order (Mind Flay 21, Improved Mind Flay 22-23, Vampiric Embrace 28). Discipline and Holy ship no raidBuild.

## Decisions

- **D1. PRIEST-R21-1 to R21-3 compare against Lightwell.** Lightwell is a priest heal from the same trainer level as Prayer of Mending, cast by both healer cards before the pull and named on both pages; its per-mana figure needs only a fraction of its charges clicked to win, so the comparison does not hang on how often a raid clicks it. The fixes drop the superlative rather than add "after Lightwell", so they hold at every band.
- **D2. PRIEST-R21-3 is not a round 20 re-file.** Round 20 carried "your best heal per mana" as a Should on the Renew comparison alone ("depends on healing gear"). Lightwell is new evidence that decides it at every band from 40 without gear.
- **D3. PRIEST-R21-4 reads "fastest heal" as the time to the healing landing**, which is what an emergency line means; Penance lands first at every band from 30. Kept the true half ("your dearest").

## Sims requested

- None required: every fix removes a superlative instead of asserting a measured order.
- Carried (optional): Shadow Word: Death ahead of Mind Blast at 32 to 51 (round 20); Discipline Gnome against Human at the re-ranked BiS; every 0.1% to 0.5% race row at the re-rank.

## Engine questions

- None new. The engine's Lightwell has no spell-power coefficient and clicks below 90% health (`lightwell.go`, an assumption it names); neither changes R21-1 to R21-3.

## Sim findings

None new. Still open from sim round 20 (not re-filed): P2 target-level healing, P3 Eureka! on a live Gnome Shadow log, P4 Lightwell Renew's crit, S2 Dark Inspiration's Spirit base.

## Sources

- Client data 1.60.1.70790: `spelltips/priest.json` (724, 27870, 27871, 7001, 401859, 1240826, 1240827, 402174, 1240720, 1240721, 1316995, 139, 6074 to 6078, 10927 to 10929, 2050 to 2053, 2054, 2055, 6063, 6064, 2060, 10963 to 10965, 25314, 2061, 9472 to 9474, 10915 to 10917, 596, 996, 10960, 10961, 25316, 8092 to 10947, 1309595 to 1309636, 2944, 19276 to 19280, 589 to 10894, 15407 to 18807, 15286, 15473, 14751, 586 to 10942, 1277324 to 1277328, 1277370 to 1277378, 10060, 17 to 10901), `talents/priest.json`, `trainables/priest.json`, `races.json`, `addon-data.json` (`rotations/priest-*`, every band), `addon-builds.json`, `bis/priest-discipline.md` (band 60 raid preset), `items/druid.json` and `items/priest.json` (stats of those items).
- Curated: `data/curated/apl/priest-{discipline,holy,shadow}.json` (notes, conditions and prepull actions).
- Engine d5c708834: `sim/priest/lightwell.go`, `penance.go`, `renew.go`, `heal_ranks.go` (Renew 0.2 a tick, Penance 0.19 a pulse, Prayer of Mending 0.429).
- Reviews: round20/priest.md, round4/priest.md (race results), ledger.md (Round 20).

## Appendix: Note sweep

One line per row note as `addon-data.json` renders it, every sentence of the curated note, at every band. Levels are when the player has each named item on the shipped order or from the trainer.

### Discipline (bands 10, 12, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 45, 48, 50, 54, 56, 58, 59, 60)

- **Power Word: Shield** (10 to 60): "On the tank at the top rank whenever Weakened Soul is off him: it soaks the next hit and costs a warrior no rage." Top rank on the row at every band (17, 592, 600, 3747, 6065, 6066, 10898, 10899, 10900, 10901). Second sentence: "Once you have Renewed Hope" (32) and "with Divine Aegis" (36), both conditional. OK.
- **Renew** (10 to 60): "Keep Renew on the tank all fight; at 60 that means rank 7 or higher." 10927 at 60. OK.
- **Lesser Heal** (10 to 38): "the party's small top-up for little mana before 40." Rank 3 (75 mana) from 10. Second sentence "by your late forties it has lost your healing bonus" (0 from 47). OK.
- **Lesser Heal** (no note, the other party and deeper-hole lines, 10 to 38): nothing to check.
- **Heal** tank (16 to 60): "until you learn Greater Heal at 40: it is your biggest heal until then." Largest heal at 16 to 38 (above); gated off from 40 (round 20 D4). OK.
- **Heal** (no note, deeper hole, 16 to 60): nothing to check.
- **Dark Sacrifice** (20 to 60): "Undead only, ... over 15 seconds, plus your Spirit, ... 10 minute cooldown". Ranks 1 to 5 (1277324 at 20, 1277325 at 30, 1277326 at 40, 1277327 at 50, 1277328 at 60), all 15 s, +Spirit, 600 s. OK.
- **Flash Heal** tank (20 to 60): "at the top rank on a tank who is really low: the fast heal, and the dear one." Top rank at every band (2061, 9472, 9473, 9474, 10915, 10916, 10917); dearest per mana at every band (above). OK (R21-S4).
- **Flash Heal** party (20 to 60): "on a party member about to die, as long as your mana is not far behind the fight." OK.
- **Flash Heal** (no note, 20 to 60): nothing to check.
- **Penance** (30 to 60): "three heals one second apart for one cast every 12 seconds, your best heal per mana." Talent at 30, ranks at 30 / 40 / 50 / 60 on the rows, 12 s, three pulses. "your best heal per mana": true at 30 to 38, false from 40. **PRIEST-R21-3.**
- **Penance** (no note, the party lines, 30 to 60): nothing to check.
- **Prayer of Healing** (30 to 60): "it heals the target and the target's whole party. Rank 4 is the top rank a trainer teaches." 596 at 30 to 48, 10960 at 50 to 59, 10961 at 60; rank 5 is the book. OK.
- **Greater Heal** tank (40 to 60): "at the top rank as the tank's filler at 60". Qualified "at 60"; 10965 from 58. OK.
- **Greater Heal rank 1** party (40 to 60): "once you have higher ranks it keeps most of your healing bonus for far less mana." Rank 2 at 46; 100% to 57, 85% at 60. OK.
- **Greater Heal** (no note, 40 to 60): nothing to check.
- **Greater Heal** deeper hole (40 to 60): "a tank under 50%." OK.
- **Greater Heal rank 1** deeper hole (40 to 60): "a party member under 35%." OK.

### Holy (bands 10, 16, 20, 22, 26, 28, 30, 32, 33, 34, 38, 40, 44, 45, 50, 51, 52, 53, 56, 58, 59, 60)

- **Renew** (10 to 60): "at 60 that means rank 7 or higher." 10927 at 60. OK.
- **Heal** (16 to 60): "Heal at your top rank before you learn Greater Heal at 40, and rank 4 after that ...: cheaper, though it keeps only about half your bonus at 60." Top rank 16 to 38 (2054, 2055, 6063, 6064); 6064 from 40; 305 against 370; 55% at 60. OK.
- **Heal** (no note, 16 to 60): nothing to check.
- **Heal** deeper hole (40 to 60 lines; 16 to 38 rows of the same note): "a tank under 50% or a party member under 35%." OK.
- **Dark Sacrifice** (20 to 60): as Discipline. OK.
- **Flash Heal** tank (20 to 60): "at the top rank on a tank who is really low or on anyone about to die: it is your dearest heal, so spend it where the size matters." Top rank at every band; dearest at every band, including against Renew rank 1 at 20 to 32 (above). OK.
- **Flash Heal** party (20 to 60): "at the top rank on a party member under 30%: ... the size matters more than the mana." OK.
- **Flash Heal** (no note, 20 to 60): nothing to check.
- **Prayer of Healing** (30 to 60): "it heals the target and the target's whole party." OK.
- **Prayer of Mending** (40 to 60): "five heals for one cast on a 10 second cooldown, the cheapest healing you have." Talent 40, every rank 5 jumps and 10 s. "the cheapest healing you have" (first and third sentences): false at every band against Lightwell. **PRIEST-R21-1.** Second sentence "next hit or direct heal" (non-periodic healing). OK.
- **Greater Heal** filler (40 to 60): "rank 1 until 57, then rank 3 or your top trainer rank, which heal about the same for their mana." 2060 at 40 to 56, 10965 at 58 and 60, 10964 at 59. OK.
- **Greater Heal** (no note, 40 to 60): nothing to check.
- **Greater Heal** deeper hole (40 to 60): OK.
- **Greater Heal rank 1** (40 to 60): "Once you have higher ranks, when the top rank would overheal or your mana runs short: it costs far less and still keeps most of your healing bonus at 60." Rank 2 at 46; 85% at 60. OK.
- **Greater Heal rank 1** deeper hole (40 to 60): OK.

### Shadow (bands 10, 16, 18, 20, 22, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 51, 52, 58, 60; leveling and raid orders)

- **Shadow Word: Pain** (10 to 60): "Always up; stop refreshing it in the closing seconds of a fight." OK.
- **Mind Blast**, Inner Focus line (10 to 60): "with Inner Focus first whenever it is up: Inner Focus makes that Mind Blast free, with 25% more crit chance." Inner Focus 57 on both orders, conditional; 25% non-periodic. OK.
- **Inner Focus** row (20 to 60), same note: whole-row case (round 19). OK as text.
- **Mind Blast**, second line (10 to 60): "a big hit for its short cast, at half cost in Shadowform." Round 20's fix; no superlative; Shadowform 42 on both orders, conditional. OK.
- **Dark Sacrifice** (20 to 60): "when mana is below a fifth ... over fifteen seconds and adds your Spirit ... ten minute cooldown ... Only Undead priests have it." OK.
- **Devouring Plague** (20 to 60): "every time its 24 second cooldown is up: every priest has it in Forever, its ticks crit, and it heals you." Every rank 24 s and "heals the caster"; DoTs crit (owner rule); no rank named. OK.
- **Mind Flay** (20 to 60): "let all three ticks land ... once you have Improved Mind Flay the last tick is the biggest." Talent 28 leveling / 21 raid, Improved Mind Flay 29 / 22, conditional. OK.
- **Vampiric Embrace** (30 to 60): "on the boss whenever it is down." Talent 33 leveling / 28 raid. Later sentences: 30 s cooldown equal to the 30 s duration, 20%, "until you have the talent". OK.
- **Shadow Word: Death** (32 to 60): "each cast that does not kill costs you 30% of your health, and once you have Vampiric Embrace keep it up to heal that back." Trainer 32; 30% on every rank; Vampiric Embrace conditional. OK (R20-S2 carried).
- **Shoot** (curated note; the wand is declared inert and renders no row): "a wand shot cannot start during the global cooldown and holds you for the wand's full speed". No rank, talent or value. OK.

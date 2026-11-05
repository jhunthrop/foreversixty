# Guide accuracy round 22: Priest

Audited on the round 22 audit tree (987d08a7c: main after guide round 21; no sim round 21 commit or report in the tree, engine pin d5c708834 in `sim/enginever/version.go`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`. Files: `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`; every note in `data/curated/apl/priest-{discipline,holy,shadow}.json`; every rendered row of `addon-data.json` (`rotations/priest-discipline` 25 bands, `priest-holy` 22 bands, `priest-shadow` 22 bands, with `race_spell_ids`); `addon-builds.json` (all six priest orders re-decoded this round, unchanged); `trainables/priest.json`, `spelltips/priest.json`, `talents/priest.json`, `bis/priest-{discipline,holy}.json` (band presets and their item healing, summed from `items/*.json`). Engine at d5c708834: `sim/priest/heal_ranks.go` (rank tables, PerLevel growth, coefficients), `renew.go`, `penance.go`, `lightwell.go`, `sim/core/spell.go` (`SpellLevelPenalty`: whole to learn level + 17, then 5% a level). Read first: round 21 priest, ledger Round 19 to 21 sections.

## Changes since round 21's tree (e039e1490)

Priest changes (`git diff e039e1490 HEAD`): af6c72839 (round 21 fixes) and e932b6b43 (round 21 review wording); `addon-data.json` regenerated (note text only; no priest row, rank or band changed).

- `priest-holy.json` priorityList[2] (Prayer of Mending, Holy 40 to 60): "the cheapest healing you can cast all fight" and "it is the last heal to give up when mana runs short". Re-checked: OK (below).
- `priest-discipline.json` priorityList[2] (Penance, Discipline 30 to 60): "cheap for all it heals". Re-checked: OK.
- holy.md:30: "Prayer of Mending is the cheapest healing you can cast all fight". OK.
- discipline.md:76: "It takes a second and a half and it costs the most mana for what it heals, so it never becomes the filler." OK (raid section, raid gear).

## Findings

None. Every note at every band on both orders, and every prose line, holds on the 70790 client data, the shipped point orders and the trainer levels.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Re-check of round 21's fixes

- **PRIEST-R21-1 / R21-2 (Prayer of Mending "cheapest you can cast all fight").** Lightwell, the spell that broke round 21's wording, has a 10 minute cooldown (`spelltips` 724, 27870, 27871: `cooldown_ms` 600000), so it is a once-a-fight cast and "all fight" excludes it, as it does Divine Grace (1277370 to 1277378, 10 min), Desperate Prayer (13908 to 19243, 10 min, Dwarf, self only), Contingency Plan (10 min) and Inner Focus (3 min). Against everything Holy casts repeatedly, at the band presets' healing (`bis/priest-holy.json`: about 354 at 40, 507 at 50, 554 in the 60 raid set): 40, Prayer of Mending rank 1 (172 + 0.429 x 354) x 1.1 Spiritual Healing x 5 = 1,781 for 178.5 (Improved Healing) = 10.0 a point of mana, against Renew rank 4 (160 + 354) x 1.15 x 1.1 for 140 = 4.6, Greater Heal rank 1 (844 + 5.1 x 0 + 0.857 x 354) x 1.1 for 314.5 = 4.0; 60 raid, rank 3 about 12.6 against Renew rank 7 5.5 and Greater Heal 4.4 or less. Holy has no Holy Nova (not in the build: 33505003030121431). True at every band 40 to 60. The note's third sentence is advice the card follows (the row has no mana-pace gate). OK.
- **PRIEST-R21-3 (Penance "cheap for all it heals").** No comparison; at 30 with the band preset's 234 healing, 3 x (126 + 0.19 x 234) x 1.05 = 536 for 150 (3.6 a point of mana) against Heal rank 3's 2.8; at 60 raid about 5 against Greater Heal's 4.4. OK at every band.
- **PRIEST-R21-4 (Flash Heal "costs the most mana for what it heals").** The Emergencies paragraph is in the raid Rotation section. At 60 in the raid preset (about 576 healing): Flash Heal rank 7 (842 + 0.429 x 576) / 380 = 2.87; Greater Heal rank 5 (1,960 + 0.857 x 576) / 603.5 = 4.1; Heal rank 4 3.7; Renew rank 7 (370 + 576) x 1.2075 / 225 = 5.1; Penance about 5; Prayer of Healing rank 4 on a party 5 x 748 / 1,030 = 3.6; Lesser Heal rank 3 (no bonus at 60) 2.56 base for 63.75. Flash Heal is the dearest. "takes a second and a half": 1500 ms. OK.

## Shoulds (wording, not findings)

- **PRIEST-R22-S1, discipline.md:49 "Penance (1 point): your best heal."** No axis, so it cannot be shown false, but it reads as the "best per mana" claim round 21 removed from the card (Greater Heal heals more per cast and per second at 60; Renew rank 7 heals more per mana in raid gear). Optional: "Penance (1 point): the heal you never let sit."
- **PRIEST-R22-S2, shadow.md:60 Shadow Focus "the cheapest hit you will ever get".** A value judgement, not a measured order; a Dwarf's Artisan's Precision is free. Optional: "5% Shadow hit for five points, so your gear needs far less of it."
- **PRIEST-R22-S3, `priest-holy.json` Flash Heal note "it is your dearest heal" (Holy 20 to 60).** True with any real healing gear, but on base numbers alone the Renew ranks the Holy rows cast from 33 heal less per point of mana than Flash Heal (33, Renew rank 3 125 x 1.15 for 105 = 1.37 against Flash Heal rank 3 1.70; 50, Renew rank 5 1.60 against Flash Heal rank 6 2.25). Renew's full 1.0 coefficient overtakes Flash Heal's 0.429 from about 40 to 130 healing depending on the band, and the band presets carry 233 to 554 healing, so the note holds on the gear the card is measured with. Not filed. Optional, if a gearless reader matters: "it heals less for its mana than your slower heals, so spend it where the size matters", which is true at every band 20 to 60 on base numbers too (Heal, Greater Heal, Prayer of Healing and Prayer of Mending all beat it).
- Carried, unchanged and optional: R21-S1 (Prayer of Healing "your most expensive heal" at 58 and 59 against Greater Heal rank 4; Rotation is the raid section), R21-S2 (Twilight Focus "worth more than any throughput point"), R21-S3 (Undead and Dwarf Mind Blast rank 6 at 50 and 51), R21-S4 (Discipline Flash Heal note "the fast heal": an epithet, true, but Penance lands first from 30), R20-S1, R20-S2, R18-S1 to S5, R17-S2.

## Checked and right this round

- **Builds** re-decoded from the frontmatter against `talents/priest.json` (tier, column order): Discipline 29/18 (Twin Disciplines 5, Mental Strength 5, Mental Agility 3, Inner Focus 1, Improved Power Word: Shield 3, Meditation 3, Soul Warding 1, Penance 1, Renewed Hope 4, Divine Aegis 3; Twilight Focus 3, Improved Renew 3, Holy Specialization 5, Divine Fury 5, Inspiration 3, Improved Healing 3); Holy 17/34 with Holy Binding, Litany of Light 2, Spirit of Redemption, Spiritual Guidance 4, Spiritual Healing 3, Prayer of Mending; Shadow leveling and raid as the page lists, the raid swap 9 for 9. The Power Infusion swap (Renewed Hope 5 makes 30 Discipline points; Power Infusion needs Penance) is legal. Every talent value quoted on the three spec pages matches its rank text.
- **Leveling orders** (`addon-builds.json`, index 0 = 10): Discipline Twin Disciplines 10 to 14, Mental Strength 15 to 19, Penance 30; Holy Twilight Focus 10 to 12, Improved Renew 13 to 15, Holy Specialization 16 to 20, Prayer of Mending 40; Shadow Wand Specialization 10 and 11, Spirit Tap 12 to 16, Shadow Focus 17 to 21, Improved Mind Blast 22 to 25, Improved Shadow Word: Pain 26 and 27, Mind Flay 28, Improved Mind Flay 29 and 30, Blackout 31 and 32, Vampiric Embrace 33, Shadowform 42. All three "Levels 1 to 19" paragraphs and shadow.md:139 match.
- **Trainer levels** quoted: Shadow Word: Pain 4, Power Word: Shield 6, Renew 8, Mind Blast 10, Inner Fire 12, Psychic Scream 14, Dispel Magic 18, Holy Fire 20, Greater Heal 40 and rank 4 at 58, Shadow Word: Death 32 (rank 2 at 40, rank 3 at 48, rank 4 at 56: "rank 2 through your fifties" matches the 50 to 58 rows).
- **Every superlative and comparison on the four pages**, at the ranks the rows cast and the band presets' gear: Discipline Heal note "your biggest heal until then" (16 to 38); "a Greater Heal crit leaves the biggest Divine Aegis you can make"; Holy "Heal rank 4 is cheaper still", "lost a little of your healing bonus by 60" (85%), "about half" (55%), "a Prayer of Healing, your most expensive heal" (60: 1,030 or 1,070 against 710); Shadow "Improved Mind Flay makes the third tick the biggest", Shadow Word: Pain "takes two ticks to out-damage the Mind Flay" (rank 8 127 a tick x 1.05 against half a rank 6 Mind Flay with Improved Mind Flay, about 237 base; one tick short, two ahead, also with the coefficients); race superlatives against the recorded race results. OK.
- **Every quoted cost, duration and cooldown**, with the build's own talents: Weakened Soul 15 s (Holy has no Renewed Hope); Soul Warding removes the 4 s cooldown; Lightwell 3 min or 5 clicks; Divine Grace and Dark Sacrifice 10 min; Inner Focus 3 min; Mind Blast 8 s and 0.5 s a point; Shadow Word: Death 18 s and 30%; Devouring Plague 24 s; Devouring Contagion half cost; Vampiric Embrace 30 s and 20%; Fade 10 s on 30 s; Silence 5 s; Blackout 4% at two points; Flash Heal 1.5 s; Penance 12 s, three pulses a second apart. OK.
- **Low-rank healing bonus** (`SpellLevelPenalty`): Renew rank 1 keeps 65% at 32 ("most"); ranks 4, 5 and 6 keep all of it at 40, 45 and 52; ranks 5 and 6 keep 45% and 75% at 60; Heal rank 4 55%; Lesser Heal rank 3 none from 47 ("late forties"); Greater Heal rank 1 whole to 57, 85% at 60. OK.
- **Shadow raid order**: every note naming a talent is conditional, so nothing is credited early at 21 to 28 on the raid order. Discipline and Holy ship no raidBuild (their raid orders equal the leveling orders).
- Gear and professions: Truefaith Vestments (Tailoring) in every 60 healer set; Dormant Heart of the Mountain (Enchanting) in the Horde Discipline raid set (`bis/priest-discipline.md`:362). OK.
- No page or rendered note names 70755 (the one hit is the Discipline APL's engine header, which does not render). No em or en dashes.

## Decisions

- **D1. "All fight" excludes the 10 minute spells.** Lightwell, Divine Grace, Desperate Prayer and Contingency Plan share a 600000 ms cooldown, so a player cannot cast them all fight; round 21's fix holds against every spell the priest has at 40 to 60, once-a-fight spells included.
- **D2. Per-mana superlatives on leveling bands are judged on the band presets' gear,** the gear the card is measured with (233 to 507 healing at 30 to 50). On that gear Flash Heal is the dearest heal at every Holy and Discipline band; the gearless case goes in R22-S3.
- **D3. R21-S4 stays a should.** "the fast heal" is an epithet like the Greater Heal note's "the big heal", not an exclusive claim, and Flash Heal is fast.

## Sims requested

- None required.
- Carried (optional): Shadow Word: Death ahead of Mind Blast at 32 to 51 (round 20); Discipline Gnome against Human at the re-ranked BiS; every 0.1% to 0.5% race row at the re-rank.

## Engine questions

- None new. Known: Lightwell has no spell-power coefficient in the engine and clicks below 90% health (`lightwell.go`), which does not change any page claim.

## Sim findings

None new. Still open from sim round 20 (not re-filed): P2 target-level healing, P3 Eureka! on a live Gnome Shadow log, P4 Lightwell Renew's crit, S2 Dark Inspiration's Spirit base.

## Sources

- Client data 1.60.1.70790: `spelltips/priest.json` (every heal, shield and Shadow rank named above, Lightwell 724/27870/27871, Divine Grace 1277370 to 1277378, Desperate Prayer 13908 to 19243, Contingency Plan 1277462 to 1277640, Holy Nova 15237 to 27801), `talents/priest.json`, `trainables/priest.json`, `addon-builds.json`, `addon-data.json` (`rotations/priest-*`, every band), `bis/priest-discipline.json`, `bis/priest-holy.json` and `.md`, `items/*.json`.
- Curated: `data/curated/apl/priest-{discipline,holy,shadow}.json` (notes, prepull actions).
- Engine d5c708834: `sim/priest/heal_ranks.go`, `renew.go`, `penance.go`, `lightwell.go`, `sim/core/spell.go` (`SpellLevelPenalty`).
- Reviews: round21/priest.md, ledger.md (Round 19 to 21).
- No web research needed: the client data decided every call.

## Appendix: Note sweep

One line per row note as `addon-data.json` renders it, every sentence of the curated note, at every band. Levels are when the player has each named item on the shipped order or from the trainer. Rows and ranks are unchanged from round 21; the two changed notes are re-derived in full.

### Discipline (bands 10, 12, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 45, 48, 50, 54, 56, 58, 59, 60; one order)

- **Power Word: Shield** (10 to 60): "On the tank at the top rank whenever Weakened Soul is off him: it soaks the next hit and costs a warrior no rage." Top rank on the row at every band (17, 592, 600, 3747, 6065, 6066, 10898, 10899, 10900, 10901). "Once you have Renewed Hope" (32), "with Divine Aegis" (36): conditional. OK.
- **Renew** (10 to 60): "at 60 that means rank 7 or higher." 10927 at 60. OK.
- **Lesser Heal** (10 to 38): "small top-up for little mana before 40" (rank 3, 75); "by your late forties it has lost your healing bonus" (none from 47). OK.
- **Lesser Heal** (no note lines): nothing to check.
- **Heal** tank (16 to 60): "until you learn Greater Heal at 40: it is your biggest heal until then." Self-qualified; largest heal at 16 to 38. OK.
- **Heal** (no note): nothing to check.
- **Dark Sacrifice** (20 to 60): ranks 1 to 5 at 20/30/40/50/60, 15 s, plus Spirit, 10 min. OK.
- **Flash Heal** tank (20 to 60): "the fast heal, and the dear one." Top rank at every band; not a superlative (D3, R21-S4). OK.
- **Flash Heal** party (20 to 60): "as long as your mana is not far behind the fight." OK.
- **Penance** (30 to 60): "three heals one second apart for one cast every 12 seconds, cheap for all it heals." Talent 30; ranks 1 to 4 at 30/40/50/60 on the rows; 12 s; three pulses; cheap at every band (3.6 at 30, about 5 at 60). OK (round 21 fix).
- **Penance** (no note, party lines): nothing to check.
- **Prayer of Healing** (30 to 60): "heals the target and the target's whole party. Rank 4 is the top rank a trainer teaches." 596, 10960, 10961 on the rows. OK.
- **Greater Heal** tank (40 to 60): "at the top rank as the tank's filler at 60". Qualified. OK.
- **Greater Heal rank 1** party (40 to 60): "once you have higher ranks it keeps most of your healing bonus for far less mana." Rank 2 at 46; 100% to 57, 85% at 60. OK.
- **Greater Heal** / **Greater Heal rank 1** deeper hole (40 to 60): thresholds only. OK.

### Holy (bands 10, 16, 20, 22, 26, 28, 30, 32, 33, 34, 38, 40, 44, 45, 50, 51, 52, 53, 56, 58, 59, 60; one order)

- **Renew** (10 to 60): "at 60 that means rank 7 or higher." OK.
- **Heal** (16 to 60): "top rank before Greater Heal at 40, and rank 4 after that ...: cheaper, though it keeps only about half your bonus at 60." 6064 from 40; 305 against 370; 55% at 60. OK.
- **Heal** deeper hole (16 to 60): OK.
- **Dark Sacrifice** (20 to 60): as Discipline. OK.
- **Flash Heal** tank (20 to 60): "it is your dearest heal, so spend it where the size matters." Top rank at every band; dearest on the band presets' gear at every band (D2; base-number case in R22-S3). OK.
- **Flash Heal** party (20 to 60): "the size matters more than the mana." OK.
- **Prayer of Healing** (30 to 60): OK.
- **Prayer of Mending** (40 to 60): "five heals for one cast on a 10 second cooldown, the cheapest healing you can cast all fight." Talent 40; ranks 1 to 3 at 40/50/60, 5 jumps, 10 s; cheapest repeatable heal at every band (D1). "next hit or direct heal": non-periodic healing. "the last heal to give up when mana runs short": the row has no pace gate. OK (round 21 fix).
- **Greater Heal** filler (40 to 60): "rank 1 until 57, then rank 3 or your top trainer rank, which heal about the same for their mana." 2060 to 56, 10965 at 58 and 60, 10964 at 59. OK.
- **Greater Heal rank 1** (40 to 60): "costs far less and still keeps most of your healing bonus at 60." 85%. OK.
- **Greater Heal** / **rank 1** deeper hole (40 to 60): OK.

### Shadow (bands 10, 16, 18, 20, 22, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 51, 52, 58, 60; leveling and raid orders)

- **Shadow Word: Pain** (10 to 60): "Always up; stop refreshing it in the closing seconds of a fight." OK.
- **Mind Blast**, Inner Focus line (10 to 60): Inner Focus 57 on both orders, "whenever it is up"; free and 25% on a non-periodic spell. OK.
- **Inner Focus** row (20 to 60): same note, conditional. OK.
- **Mind Blast**, second line (10 to 60): "a big hit for its short cast, at half cost in Shadowform." Shadowform 42 on both orders, conditional. OK.
- **Dark Sacrifice** (20 to 60): OK.
- **Devouring Plague** (20 to 60): 24 s on every rank; heals the caster; DoTs crit (owner rule). OK.
- **Mind Flay** (20 to 60): Improved Mind Flay 29 leveling / 22 raid, conditional. OK.
- **Vampiric Embrace** (30 to 60): 30 s cooldown and duration, 20%, "until you have the talent" (33 leveling / 28 raid). OK.
- **Shadow Word: Death** (32 to 60): 30% on every rank; Vampiric Embrace conditional. OK.
- **Shoot** (curated note, no rendered row): no rank, talent or value. OK.

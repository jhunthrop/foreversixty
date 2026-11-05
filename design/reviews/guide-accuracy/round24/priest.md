# Guide accuracy round 24: Priest

Audited on the round 24 tree (ea9adea1c: main after guide round 23 landed, which includes sim round 21, range 24ee457d0..d5cc4d6b4). Engine pin 170a59c39 (`sim/enginever/version.go`; source read at `scratchpad/gr24/engine`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited, line by line:

- `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`.
- Every note and every row condition in `data/curated/apl/priest-{discipline,holy,shadow}.json`.
- Every rendered row of `addon-data.json`, with `race_spell_ids`: `rotations/priest-holy` (now 23 bands), `priest-discipline` (25), `priest-shadow` (22). The addon shows a note's first sentence; the site card shows the whole curated note, so every sentence was read at every band.
- `data/curated/rank-choices.json` (the Heal rows sim round 21 added for Holy and Discipline).
- All six priest orders in `addon-builds.json`, decoded against `talents/priest.json`.
- `trainables/priest.json`, `spelltips/priest.json`, `bis/priest-*.json`, with band healing summed from `items/*.json`.
- Engine: `sim/priest/heal_ranks.go` (Heal and Penance rank rows: Amount, PerLevel, SpellLevel, MaxLevel, coefficient), `penance.go`.

Read first: round 23 priest, the ledger's Round 23 section, sim round 21's `priest.md` and `integration.md`, and the coordinator notes.

Order of work, as the brief asks:

1. Sim round 21's priest changes. 8dcb3404d, 2e3a7d7a7 and e7623bfe9 made these changes:
   - Holy: about 650 lines in `priest-holy.json`, ten Lesser Heal rank 3 lines below 16, the Heal note rewritten, and Heal rank rows at 28, 29 and 34 to 39.
   - Discipline: Heal rank rows at 30 to 39.
   - Prose: holy.md:114 and :122, and discipline.md:124.
2. Round 23's five fixes (619a5dda0, then the wow-player wording in 9577f4bd3).
3. The full sweep with every lens from rounds 20 to 23.

## Changes since round 23's tree (24ee457d0)

- **Holy card:**
  - Band 10 used to show Renew alone. It now shows Renew rank 1 and ten Lesser Heal rank 3 lines, gated to `currentLevel < 16`, with a new note.
  - New bands 29 and 39.
  - The Heal rows take race overrides: rank 2 at 28 for Human and Undead, and at 29 for Undead; rank 3 at band 33 (34 to 37) for every race, at 38 for Troll and Undead, and at 39 for Undead.
  - The Heal note was rewritten.
- **Discipline card:**
  - Heal rank 3 at bands 34, 36 and 38 for every race, where it used to be rank 4.
  - Heal rank 2 at bands 30 and 32 for an Undead.
  - No note changed.
- **Shadow card:** no change from sim round 21. Round 23's Mind Blast note fix is in place.
- **Round 23's fixes** (Discipline shield, Mind Blast before Mind Flay, Flash Heal for speed, the two stat lines, "up to five" Prayer of Mending heals): all present and all re-checked (see below). The stat lines now read "though a rank learned far below your level keeps less of it, or none", the wow-player wording.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R24-1 | CONFIRMED | card note (contradicts the rows it sits on, by race and band; new in sim round 21) | `priest-holy.json` priorityList[30], second sentence, Holy bands 16 to 60 | "A new rank costs more mana than its extra healing is worth at first, so keep rank 2 at 28 and rank 3 from 34 to 37." | The rows this note sits on contradict it in three places (`addon-data.json` `race_spell_ids`; `rank-choices.json` Heal 6064):<br>(a) **Band 28.** Dwarf, Gnome, Night Elf and Troll rows cast rank 3. Only Human and Undead keep rank 2 (sim round 21 P-H23: Dwarf, Night Elf and Troll lose 0.6% to 1.5% on rank 2, and Gnome ties).<br>(b) **Bands 38 and 39.** Troll and Undead rows keep rank 3 at 38, and Undead rows at 39. The note puts the switch at 37.<br>(c) **Band 22.** The general rule "a new rank ... at first" does not hold here. Rank 2 is cast at 22, the level it is learned, for every race, and rank 3 is cast at 28 for four races.<br>A player of those races reads one rank on the row and another in its note. Sim round 21 itself kept "rank 3 at 29" out of this note because "the card rows follow race" (its "For the guide lane"). The rule it kept still goes beyond the rows. | "Heal is your filler before you learn Greater Heal at 40. Rank 4 costs more mana than its extra healing is worth in its first levels, so keep rank 3 from 34 to 37, to 38 as a Troll and to 39 as an Undead. A Human or Undead also keeps rank 2 at 28, and an Undead at 29. From 40 ..." (rest unchanged). R24-3 explains why the 28 clause carries no reason. |
| PRIEST-R24-2 | CONFIRMED | page prose (rank by level, false for the page's Horde pick; new in sim round 21) | holy.md:122 | "A new Heal rank costs more mana than its extra healing is worth in its first levels: keep rank 2 at 28 and move to rank 3 at 29, then keep rank 3 from 34 to 37 before you move to rank 4." | Undead is this page's Horde pick (`recommendedRaces: [human, undead]`, :102), and the card keeps it on rank 2 at 29 and on rank 3 at 38 and 39. Troll also keeps rank 3 at 38, and four races cast rank 3 at 28. The same general rule fails at 22, as in R24-1(c). These are measured rows, beyond the sim's 0.2 HPS bar at both seeds: Undead gains +1.5% at 29, +0.5% at 38 and +0.4% at 39. So "move to rank 3 at 29" and "move to rank 4" at 38 tell an Undead the opposite of the card. See D1 for why this is a finding while R23-S4 stayed a Should. | "Heal rank 4 costs more mana than its extra healing is worth in its first levels, so keep rank 3 from 34 to 37 before you move to rank 4; a Troll keeps it to 38 and an Undead to 39. A Human or Undead also keeps rank 2 at 28, and an Undead at 29. Renew steps down too: through 32 ..." ("the same way" goes, since Renew's reason differs; see R24-S2). |
| PRIEST-R24-3 | CONFIRMED | stated reason false per cast on the client numbers (round 18 rule; new in sim round 21) | `priest-holy.json` priorityList[30] (Holy bands 28 and 29); holy.md:122 | "A new rank costs more mana than its extra healing is worth at first, so keep rank 2 at 28" / "costs more mana than its extra healing is worth in its first levels: keep rank 2 at 28 and move to rank 3 at 29" | Engine rows (`heal_ranks.go`), which match the client:<br>- Rank 2: 375 at 22, +2.9 a level, MaxLevel 27, so 389.5 at 28. Cost 205.<br>- Rank 3: 523 at 28, +4 a level, so 527 at 29. Cost 255.<br>- Both use coefficient 0.857 and keep their whole bonus here.<br>Healing per point of mana is equal at 184 healing at 28 and 203 at 29. Below that, rank 3 heals more for its mana. At 28 and 29 Holy wears band-20 gear (154 healing for a Human, 144 for an Undead; `bis/priest-holy.json` summed from `items/*.json`) and has no Spiritual Guidance until 39. So rank 3 gives 1.0% (Human) and 1.3% (Undead) more healing per mana at 28: 2.57 against 2.54, and 2.54 against 2.50. Its extra 133.5 healing costs 50 mana, 2.67 a point, more than rank 2's average. Whatever wins rank 2 the measured +0.4% / +1.2%, it is not mana per heal. (Spiritual Healing, Holy Specialization and Improved Healing scale both ranks alike.) The reason does hold for rank 4: rank 3 heals more per mana at 34 to 39 above 9 to 116 healing, and Holy wears 233 to 244 then. | Keep the mana reason on rank 4 only, and give the 28 and 29 hold without a reason, as in R24-1 and R24-2. Sim 2 names the real cause if the owner wants a reason back. |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. The three share one sentence on each surface, so one rewrite of the note and one of the prose sentence closes all three.

## Shoulds (wording, not findings)

New this round:

- **PRIEST-R24-S1, holy.md:114 "Healing a group before you learn Heal at 16, keep Renew on the tank and use Lesser Heal rank 3 as your filler."** Rank 3 trains at 10 and Renew at 8, so the line is true from 10, where the card's band 10 starts. A group before 10 is rare, so this is not filed. Optional: "Healing a group from 10, before you learn Heal at 16, ...".
- **PRIEST-R24-S2, holy.md:122 "Renew steps down the same way".** Heal's step (hold the old rank a level or a few after the new one arrives) and Renew's (keep rank 1 through 32, a cheap rank that keeps most of its bonus) are different rules for different reasons. With R24-2's fix this becomes "Renew steps down too".
- **PRIEST-R24-S3, discipline.md:124.** The sentence repeats itself: "Heal is the tank's filler until you learn Greater Heal at 40; when rank 4 arrives at 34, keep casting rank 3 until you learn Greater Heal at 40, since ...". Optional: "Heal is the tank's filler until you learn Greater Heal at 40; when rank 4 arrives at 34, keep casting rank 3, since rank 4 costs more mana than its extra healing is worth." The reason is true per mana: rank 3 is ahead above 9 healing, and Discipline wears 234 to 251.
- **PRIEST-R24-S4, Discipline Heal rank 2 for an Undead at 30 to 33.** These rows are new in sim round 21. Neither the page nor the note mentions them, and neither says anything they contradict ("Heal is the tank's filler"; the note's "biggest heal until then" holds for rank 2 too: 389.5 + 0.857 x 251 = 605, against Penance rank 1's 3 x (126 + 0.19 x 251) x 1.05 = 547). Optional, alongside R23-S4.

Carried, unchanged and optional: R23-S1 to S3, R23-S5 to S12, R22-S1, R22-S2, R21-S1 to S4, R20-S1, R20-S2, R18-S1 to S5, R17-S2. R23-S4 (Discipline Renew by race) stays a Should (D1).

## Checked and right this round

- **Sim round 21, Holy band 10:**
  - The Lesser Heal note says "your filler until you learn Heal at 16". Heal rank 1 trains at 16, and every Lesser Heal row carries `currentLevel < 16`.
  - "rank 3": 2053 trains at 10, 163 healing for 75 mana in a 2.5 second cast.
  - "on the tank first and on the party once they are low": the tank rows (under 80%) come before the party rows (under 65%), and the deeper rows (under 50% and 35%) follow.
  - "Pace your mana the way you will with Heal": the same 0.9 pace gate.
  - Renew rank 1 at 10 to 13 and at 14 (rank-choices 14 to 32) matches "Keep Renew on the tank".
- **Sim round 21, Holy Heal note, the rest:**
  - "filler before ... 40": Greater Heal rank 1 trains at 40.
  - "From 40 Greater Heal is your main heal; fall back to Heal rank 4 only when you are short of mana for Greater Heal rank 1": row [30] carries the same gate and threshold as Greater Heal rows [10] and [20] above it, so it fires only when neither is affordable.
  - "cheaper": 305 against 370, and 259 against 314.5 once Improved Healing is in.
  - "about half your bonus at 60": 55%, since MaxLevel 39 and 21 levels over less 12 of grace is 45% lost.
- **Sim round 21, Discipline prose** "rank 4 costs more mana than its extra healing is worth" at 34 to 39:
  - Per mana, rank 3 (543 at 33+, 255 mana) beats rank 4 (651 at 34, +4.5 a level, 305 mana) above 9 healing at 34, and above 116 at 38.
  - Discipline wears 234 to 251 on band-30 gear.
  - The rows show rank 3 for every race at bands 34, 36 and 38; band 38 covers 39.
- **Holy prose :122, rest of the sentence:**
  - "Heal as the filler from level 16" holds.
  - Renew rank 1 through 32 and rank 3 from 33 (rank-choices 14 to 32 and 33 to 39) hold.
  - Rank 4 from 40, rank 5 from 45 and rank 6 from about 52 hold: band 51 overrides every race to 5 or 6, band 52 gives Dwarf, Gnome and Night Elf rank 5, and band 53 gives rank 6 to all.
  - "rank 1 as the filler until you learn rank 4 at 58" holds.
- **Round 23's fixes:**
  - **Discipline shield (note and :68):** "once he is hurt" is the under-90% gate. Weakened Soul and the 0.9 mana pace are both stated.
  - **Shadow Mind Blast note:** "once you have it" covers Mind Flay at 28 (leveling) and 21 (raid).
  - **Flash Heal notes:** "heals less for its mana than your slower heals" holds at every band, on band gear:
    - 20: Flash Heal rank 1 is 2.08 a point of mana against Heal rank 1's 2.65.
    - 26: 2.03 against 2.53.
    - 32: 2.23 against 2.90.
    - 56: Greater Heal rank 1 is ahead even before Improved Healing.
    - 60: about 2.87 against Heal rank 4's 3.71 with its 55% bonus, and further behind Greater Heal.
  - **Stat lines:** "keeps less of it, or none" matches the Low ranks paragraphs.
  - **Prayer of Mending:** "up to five" matches the tooltip.
- **Shadow card and page:** unchanged since round 23. The rows were re-read band by band against shadow.md:137:
  - Shadow Word: Pain rank 3 at 30 to 38.
  - Devouring Plague rank 1 at 30 to 42, rank 2 at 44 to 58, and a Human on rank 1 at 50 and 51.
  - Mind Blast rank 4 at 40 and 42 for every race but Gnome.
  - Shadow Word: Death rank 2 at 50 to 58.
- **Builds:** decoded again and unchanged. Discipline is 29/22, Holy 17/34, Shadow leveling 20/31 and raid 18/33. The healer raid orders equal their leveling orders.
- **Every talent value quoted on the three pages** was re-read against the rank text, including Improved Healing (Lesser Heal, Heal, Greater Heal, Penance and Prayer of Mending; each page names the one its spec has), Renewed Hope 4, Divine Fury, Spiritual Guidance and Holy Binding.
- **No priest page or rendered note names 70755.**

## Decisions

- **D1. R24-2 is CONFIRMED, while R23-S4 (Discipline Renew by race) stays a Should.** R23-S4's prose names no rank as a switch point a race should take, and its margins sat inside the round 20 noise reading. Holy :122 is different. It is a sentence written this round from a race-split measurement, and it gives switch levels ("move to rank 3 at 29", "move to rank 4" after 37). The card reverses those for the page's own Horde pick, at margins sim round 21 accepted at both seeds. R24-1 contradicts the row it sits on, the plainest case.
- **D2. R24-3 is CONFIRMED on the round 18 rule.** A rank reason must hold per cast on the client numbers, and at 28 and 29 the band gear's healing (154 and 144) is under the 184 and 203 crossings. The measured win stands as a card row. Only its stated cause is wrong, so the fix drops the cause for 28 and 29 rather than the row.
- **D3. The Heal-at-22 point goes into R24-1 and R24-2, not into a finding of its own.** Per mana, rank 1 beats rank 2 at 22 above 17 healing, so "a new rank costs more than it is worth" may well be true there. The card casts rank 2 because the ranker never tries a rank learned below 20 (the APL header's rank rule). The contradiction is between the general rule and the card, and the fix scopes the rule to rank 4. Sim 1 decides whether the card should change.
- **D4. Holy band 10 card actions** are accepted, per the brief's carve-out for card actions. That covers the Lesser Heal lines, and Discipline having no tank filler before 16. Only their notes were filed against; they were OK.

## Sims requested

1. **Holy and Discipline, Heal rank 1 against rank 2 at 22 to 27** (band-20 gear, every race, heal profile, 10,000 iterations, seeds 1 and 50001). The ranker excludes ranks learned below 20, so this was never measured, and per mana rank 1 leads above 17 healing. It decides D3 and whether the card should hold rank 1.
2. **Holy, Heal rank 2 against rank 3 at 28 and 29:** overheal share, mana left at the end, and casts by rank for Human and Undead. Per mana, rank 3 leads (R24-3), so this names the real reason for the measured rank 2 win, if the owner wants a reason in the note.
3. **Discipline, bands 10 to 15 (optional):** a tank Lesser Heal rank 3 line under 80% with the 0.9 pace, against the shipped shield-and-Renew-only tank. Holy got this filler in sim round 21 (+190% healing on band-20 gear); Discipline has none on the tank before Heal at 16.
4. Carried from round 23 and already routed: Prayer of Mending heals per cast; Shadow Word: Pain at 6 against 10 seconds; the Discipline shield with no health gate; Holy Greater Heal rank 1 behind the pace.

## Engine questions

None new. The open P2 (heals scaling with the target's level) is the one engine item that could explain R24-3's gap between per-mana efficiency and the measured rank 2 win at 28. It is already routed.

## Sim findings

None new. Still open and not re-filed: P2, P3 (Eureka! on a live Gnome Shadow log), P4 (Lightwell Renew's crit), S2 (Dark Inspiration's Spirit base).

## For other lanes

- **Coach starter `data/curated/coach/priest-holy.md`** (outside this audit's scope; noted so the right lane sees it):
  - :26 and :28 say Shadowfiend is in the Holy priest's kit. index.md:28 and the APL header say no Forever priest can learn it.
  - :22 and :47 make "a mid rank of Heal" the filler.
  - :22 says Prayer of Healing "heals your own party", but the Forever spell heals the target's party.

## Sources

- **Client data 1.60.1.70790:**
  - `spelltips/priest.json`: Lesser Heal, Heal, Flash Heal and Greater Heal costs and amounts.
  - `talents/priest.json`, `trainables/priest.json`, `addon-builds.json`.
  - `addon-data.json` rotations with `race_spell_ids`.
  - `bis/priest-{holy,discipline}.json` with `items/*.json`. Healing per band: Holy 154 and 144 at 20, 233 and 244 at 30, 554 raid; Discipline 154 and 146 at 20, 234 and 251 at 30.
- **Curated:** `data/curated/apl/priest-{discipline,holy,shadow}.json`, `data/curated/rank-choices.json`.
- **Engine 170a59c39:** `sim/priest/heal_ranks.go` (Heal 2054, 2055, 6063, 6064; Penance 402174), `sim/priest/penance.go`.
- **Reviews:** `design/reviews/2026-11-05-sim-audit-round21/priest.md` and `integration.md`; round 23 priest; ledger Round 23.
- **No web research needed:** the client data and the engine rows decided every call.

## Appendix: Note sweep

One line per row note, as `addon-data.json` renders it, with every sentence of the curated note, every band, and both orders (the healer raid orders equal their leveling orders). Levels are when the player has each named item on the shipped order or from the trainer. Ranks per band are the rendered rows, with race overrides in braces.

### Holy (bands 10, 16, 20, 22, 26, 28, 29, 30, 32, 33, 38, 39, 40, 44, 45, 50, 51, 52, 53, 56, 58, 59, 60)

- **Renew** (10 to 60): rank 1 from 10 to 32, 3 at 33 to 39, 4 at 40 and 44, 5 at 45 and 50, 7 at 51 and 52 (with every race overridden to 5 or 6), 6 at 53 to 59, and 7 at 60. "at 60 that means rank 7 or higher". OK.
- **Lesser Heal rank 3, tank** (band 10, which covers 10 to 15; new): "filler until you learn Heal at 16" (16); "rank 3" (10); tank before party; "once they are low" (65%); mana pace as Heal's (0.9). OK.
- **Lesser Heal** (party and deeper rows at band 10, no note): nothing to check.
- **Heal, tank** (16 to 60; ranks 1 at 16 and 20, 2 at 22 and 26, 3 {Human 2, Undead 2} at 28, 3 {Undead 2} at 29, 3 at 30 to 33, 4 {Troll 3, Undead 3} at 38, 4 {Undead 3} at 39, and 4 from 40):
  - "filler before ... 40": OK.
  - "A new rank costs more mana ... at first": **R24-1(c)** (rank 2 at 22, rank 3 at 28 for four races) and **R24-3** (false per mana at 28 and 29).
  - "keep rank 2 at 28": **R24-1(a)** (Dwarf, Gnome, Night Elf and Troll rows cast rank 3).
  - "rank 3 from 34 to 37": **R24-1(b)** (Troll and Undead at 38, Undead at 39).
  - "From 40 ... fall back to Heal rank 4 only when short of mana for Greater Heal rank 1": OK.
  - "cheaper": 305 against 370. OK.
  - "about half your bonus at 60": 55%. OK.
- **Heal** (no note): nothing to check.
- **Heal, deeper hole** (16 to 60): "a tank under 50% or a party member under 35%" matches rows [35] to [39]. OK.
- **Dark Sacrifice** (20 to 60; ranks 1 to 5 at 20, 30, 40, 50 and 60): "Undead only"; under 70% mana with health above 70% matches the condition; 15 seconds; Spirit; 10 minutes. OK.
- **Flash Heal, tank** (20 to 60; top rank at every band): "at the top rank"; "heals less for its mana than your slower heals" holds at every band, against Heal to 39 and Greater Heal from 40 (per-mana figures above). OK.
- **Flash Heal, party** (20 to 60): "under 30%" matches the condition; "speed matters more than the mana". OK.
- **Flash Heal** (no note): nothing to check.
- **Prayer of Healing** (30 to 60; rank 1 to 59, rank 4 at 60): "three of your four party members under 85%" and "target and whole party". OK.
- **Prayer of Mending** (40 to 60; ranks 1, 2 from 50, 3 at 60):
  - "every time it is ready": the row is unconditional. OK.
  - "up to five heals": OK.
  - 10 second cooldown: OK.
  - "cheapest ... all fight": routed in round 23, sim 1.
  - "next hit or direct heal": OK.
  - "tank is hit and healed the most": OK.
  - "not held back by the mana pace": the row has no gate. OK.
- **Greater Heal, filler** (40 to 60; rank 1 at 40 to 56, 4 at 58, 3 at 59, 4 at 60): "while your mana keeps pace" (0.9 gate); "rank 1 until 57, then rank 3 or your top trainer rank ... about the same for their mana". OK.
- **Greater Heal, deeper hole** (40 to 60): "a tank under 50% or a party member under 35%". OK.
- **Greater Heal rank 1** (40 to 60): "once you have higher ranks" (46); "costs far less" (370 against 455 to 655); "most of your healing bonus at 60" (85%). OK (R23-S12).
- **Greater Heal rank 1, deeper hole** (40 to 60): "waits for a deeper hole too". OK.

### Discipline (bands 10, 12, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 45, 48, 50, 54, 56, 58, 59, 60)

- **Power Word: Shield** (10 to 60; ranks 1, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 9, 9, 9, 9, 10):
  - "once he is hurt", "Weakened Soul is off him" and "mana keeps pace" match the under-90% gate, the aura check and the 0.9 pace (round 23 fix). OK.
  - "costs a warrior no rage": OK.
  - Renewed Hope (32) and Divine Aegis (36): both "once you have". OK.
- **Renew** (10 to 60; rank 1 to 28, 2 at 30 and 32, 3 at 34 to 38, 4 at 40 and 42, Human 7 and others 4 at 44, Human 7 and others 5 at 45, 5 at 48, 7 from 50, and 9 for four races at 56 to 59 with Undead also at 59): "at 60 that means rank 7 or higher". OK (R23-S4 on prose).
- **Lesser Heal rank 3, party** (10 to 38): "small top-up for little mana before 40" (75 mana); "by your late forties it has lost your healing bonus" (none from 47); "costs regeneration". OK.
- **Lesser Heal** (other party rows, 10 to 38): nothing to check.
- **Heal, tank** (16 to 60; ranks 1 at 16 to 20, 2 at 22 to 26, 3 at 28, 3 {Undead 2} at 30 and 32, 3 at 34 to 38, and 4 from 40, where the row is gated off by `spellIsKnown` Greater Heal): "until you learn Greater Heal at 40"; "your biggest heal until then", including the Undead rank 2 at 30 to 33 against Penance rank 1 (605 against 547 at 251 healing). OK (R24-S4, R23-S6).
- **Heal** (no note): nothing to check.
- **Dark Sacrifice** (20 to 60): as Holy. OK.
- **Flash Heal, tank** (20 to 60; top rank): "the fast heal, and the dear one". OK.
- **Flash Heal, party** (20 to 60): "as long as your mana is not far behind" (0.7 gate). OK.
- **Flash Heal** (no note): nothing to check.
- **Penance, tank** (30 to 60; ranks 1 at 30, 2 at 40, 3 at 50, 4 at 60): 12 seconds; three pulses one second apart ("instantly and every 1 sec for 2 sec"); "on whoever is hurt" (tank under 95%, party under 70%); "cheap for all it heals". OK.
- **Penance** (party rows, no note): nothing to check.
- **Prayer of Healing** (30 to 60; ranks 1, 3 from 50, 4 at 60): "three of your four party members under 85%"; "target and whole party"; "Rank 4 is the top rank a trainer teaches" (rank 5 is the book rank). OK.
- **Greater Heal, tank** (40 to 60; rank 1 to 56, rank 4 from 58): "at the top rank ... at 60". OK.
- **Greater Heal rank 1, party** (40 to 60): "missing a third" (under 65%); "once you have higher ranks" (46); keeps most of the bonus (all of it to 57, 85% at 60); "far less mana". OK.
- **Greater Heal** (no note): nothing to check.
- **Greater Heal, deeper hole, tank** (40 to 60): "a tank under 50%". OK.
- **Greater Heal rank 1, deeper hole** (40 to 60): "a party member under 35%". OK.

### Shadow (bands 10, 16, 18, 20, 22, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 51, 52, 58, 60; leveling and raid orders)

- **Shadow Word: Pain** (10 to 60; ranks 2, 2, 3, 3, 3, 4, 4, 3 at 30 to 38, 5, 6 at 42 to 48, 7 at 50 to 52, 8): "Always up"; "closing seconds" (10 second gate). OK (R23-S2).
- **Mind Blast with Inner Focus** (10 to 60): Inner Focus at 57 on both orders, "whenever it is up"; free; 25% on a non-periodic spell. OK as text.
- **Inner Focus row** (20 to 60): same note. A whole-row case (round 23 D5). OK as text.
- **Mind Blast, second line** (10 to 60; ranks 1, 2, 2, 2, 3, 3, 4, 4, 4, 5, 5, 5, 6 {four races on 4} at 40 and 42, 6, 7, 7, 7 {Undead 6} at 50, 7 {Dwarf and Undead 6} at 51, 8, 9, 9):
  - "ahead of Mind Flay once you have it": Mind Flay at 28 (leveling) and 21 (raid). OK.
  - "big hit for its short cast": OK.
  - "half cost in Shadowform": 42, and conditional. OK.
- **Dark Sacrifice** (20 to 60): "below a fifth and behind the fight" matches the condition; 15 seconds; Spirit; 10 minutes; "Only Undead priests have it". OK.
- **Devouring Plague** (20 to 60; ranks 1, 1, 1, 2 {Gnome and Troll 1} at 28, 1 at 30 to 42, 2 at 44 to 48, 4 {every race overridden; Human 1, others 2} at 50 and 51, 2 at 52 and 58, 6 at 60): 24 seconds at every rank; "every priest has it"; ticks crit; heals you. OK.
- **Mind Flay** (20 to 60): "with every free global" (a whole-row case at 20 to 27 on the leveling order); "a clipped channel loses the rest"; Improved Mind Flay at 29 (leveling) and 22 (raid), stated as conditional. OK.
- **Vampiric Embrace** (30 to 60): the talent is at 33 (leveling) and 28 (raid); 30 seconds; 20%; "Skip it in the last 10 seconds, and until you have the talent" matches the condition. OK.
- **Shadow Word: Death** (32 to 60; ranks 1, 2 at 40 to 46, 3 at 48, 2 at 50 to 58, 4 at 60):
  - "ahead of Mind Flay": the row starts at 32, after Mind Flay at 28 and 21. OK.
  - 30%: OK (R23-S7).
  - Vampiric Embrace: stated as conditional. OK.
- **Shoot** (curated note, no rendered row): no rank, talent or value. OK.

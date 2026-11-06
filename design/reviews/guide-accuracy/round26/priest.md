# Guide accuracy round 26: Priest

Audited on the round 26 tree (head ca06416cd: main after guide round 25 landed, with sim round 22 merged in through fd758eefc). Engine pin bff826207 (`sim/enginever/version.go`); engine source read at that pin. Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited, line by line:

- `web/src/content/guides/priest/{index,discipline,holy,shadow}.md`.
- Every note and row condition in `data/curated/apl/priest-{discipline,holy,shadow}.json`, and every rendered row of `addon-data.json` `rotations/priest-*` with `race_spell_ids`: Holy 25 bands, Discipline 23, Shadow 24.
- The coach files, read from scratch:
  - `data/curated/coach/priest-{discipline,holy,shadow}.md` (Holy is a hand-written persona; Discipline and Shadow are generated starters).
  - The rule files `data/curated/coach/priest-*.json` and the priest entries in `data/pipeline/coachrules.py`.
  - The coach checks that read them: `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go` and `f7_death.go`.
- `addon-builds.json` (all six priest orders, decoded against `talents/priest.json`), `trainables/priest.json`, the engine's `sim/priest/heal_ranks.go` and `sim/core/spell.go` (`SpellLevelPenalty`: whole to MaxLevel + 12, then 5% a level), and client `SkillLineAbility` 70790 (race masks).

Read first: round 25 priest, the ledger's Round 25 section, and sim round 22 (`design/reviews/2026-11-05-sim-audit-round22/priest.md` and `integration.md`).

## New work since round 25's tree (438f8ba15)

Sim round 22 was not in round 25's tree, so its priest changes are new work this round. All of them were audited first.

**Sim round 22 (e7093d9b9, 419f70c94, 37d5082fd, 5dbffb0bf, 928ddcf14):**
- New Heal rank rows for Holy (rank 1 through the 20s and early 30s, plus race rows at 27 to 39) and Discipline (rank 1 at 22 to 35, rank 3 at 36 to 39).
- Discipline tank Lesser Heal rows below 16.
- The Holy Heal note, the Discipline tank Heal note and the new Discipline tank Lesser Heal note.
- holy.md Levels 20 to 60 (now four paragraphs) and discipline.md Levels 1 to 19 and 20 to 60.
- "Healer card notes read right at 60" (5dbffb0bf): the Holy mana-short notes now start with their condition, and the Greater Heal rank 1 note lost "Once you have higher ranks" (R26-1).

**The round 22 row filter.** Shadow rows now follow the leveling order:
- Mind Flay from 28.
- Vampiric Embrace from 33.
- The Inner Focus row from 57.

**Round 25 coach rewrite:** every coach line below was re-read from scratch.

Coach regeneration: `persona_starter.py --spec priest-discipline` and `--spec priest-shadow` were run into the scratchpad (`gr26p/regen/`). Both match the committed files byte for byte, so a regen fixes nothing below. The tree is unchanged (`git status` clean).

## Sim round 22's work, checked

- **Holy Heal rows** (`race_spell_ids`, Heal 2054/2055/6063/6064). The prose and the note match every row at every band and race:
  - 16 to 26: rank 1.
  - 27: Human, Gnome and Night Elf rank 2; Dwarf, Troll and Undead rank 1.
  - 28: Dwarf, Gnome, Night Elf and Troll rank 3; Human rank 2; Undead rank 1.
  - 29: rank 1.
  - 30 to 31: Alliance rank 1, Horde rank 3.
  - 32: Dwarf rank 2, Troll rank 1, the others rank 3.
  - 33: rank 1.
  - 34: rank 3, Troll rank 1.
  - 35 to 37: rank 3.
  - 38: Alliance rank 4, Undead rank 3, Troll rank 2.
  - 39: Alliance rank 4, Troll and Undead rank 2.
- **"Keeps your whole healing bonus until 33".** Heal rank 1 has MaxLevel 21, and the grace runs to MaxLevel + 12 = 33. True.
- **"Heals more for its mana than any later rank" (holy.md:124, discipline.md:124, both notes).** Engine rows, all with coefficient 0.857:
  - Rank 1: 280 for 155 mana from 21.
  - Rank 2: 389.5 for 205 at 27+.
  - Rank 3: 523 to 543 for 255.
  - Rank 4: 651 + 4.5 a level for 305.
  - Per point of mana, rank 1 leads rank 2 above 69 healing, rank 3 above 113 (at 28) to 149 (at 33), and rank 4 above less.
  - Band gear carries 144 to 154 healing at 22 to 29 and 233 to 251 at 30 to 39, so the claim holds wherever rank 1 is cast.
  - At 34 and 35, rank 1 keeps 95% and 90% of the bonus and still leads rank 3: 3.00 against 2.94 at 240 healing.
- **Discipline Heal rows**: rank 1 at bands 16 to 34 (16 to 35), rank 3 at 36 and 38. OK.
- **Discipline tank Lesser Heal** (bands 10 and 12; gated below 16; gone at 16): OK. The discipline.md:116 sentence matches it.
- **Holy and Discipline mana-short notes** ("When your mana falls behind the fight, only on a tank under 50% ..."): those rows carry no pace gate and sit under the paced filler, so the wording describes when they fire. OK.
- **Holy Greater Heal rank 1 note**: R26-1.
- **Shadow rows after the filter**: every "once you have" clause still holds on the leveling order:
  - Mind Flay from 28, Improved Mind Flay at 29.
  - Vampiric Embrace at 33, so the Shadow Word: Death note at 32 is conditional.
  - The Inner Focus Mind Blast row: S3.

## Findings

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PRIEST-R26-1 | CONFIRMED | card note (a round 19 fix reverted; rank compared before it is trained) | `priest-holy.json` priorityList[20], Greater Heal rank 1 row, Holy bands 40 to 60 (rendered on row 19) | "When the top rank would overheal or your mana runs short, since it costs far less and still keeps most of your healing bonus at 60." | 5dbffb0bf (sim round 22, "healer card notes read right at 60") dropped round 19's opening "Once you have higher ranks,". Rounds 20 to 25 checked that clause as what made the note true at 40 to 45.<br>- At 40 to 45, Greater Heal rank 1 is the only rank (rank 2 trains at 46). The filler row above casts it too, so "costs far less" compares rank 1 with itself.<br>- At 46 to 57, the filler row also casts rank 1 (its own note: "rank 1 until 57"; holy.md:128). "When the top rank would overheal" describes no row the player has.<br>The note is true only from 58. | "Once you have higher ranks, when the top rank would overheal or your mana runs short: it costs far less and still keeps most of your healing bonus at 60." This is round 19's text, which keeps the condition first, as 5dbffb0bf wanted. Regenerate `addon-data.json` and Data.lua (note text only). |
| PRIEST-R26-2 | CONFIRMED | coach rule file: a racial spell judged for every race | `coach/priest-{discipline,holy,shadow}.json` "defensives"; `coachrules.py:82, :190, :197` (`defensives=["Desperate Prayer"]`) | Desperate Prayer (13908 to 19243), `"talent": false`, as each priest spec's only defensive | In Forever, Desperate Prayer is the Dwarf priest spell. Client `SkillLineAbility` 70790 gives 13908 and 19243 `RaceMasks_0` 4 (Dwarf). index.md:43 says "Dwarf: Desperate Prayer and Chastise".<br>F7 (`f7_death.go` `defensivesAt`) counts a baseline defensive that was never cast from its trainable level. `trainables/priest.json` lists it at 10 with no race. So on a reviewable death, the coach tells a Human, Gnome, Night Elf, Troll or Undead priest "Died at {t}, Desperate Prayer not used" (`TitleDiedNotUsed`), or that it was ready. Five of the six priest races, and every recommended race, do not have the spell. | Rule-file change, `coachrules.py`: priest `defensives=[]` for all three specs, with a comment ("Desperate Prayer is the Dwarf priest spell, SkillLineAbility race mask 4; no defensive every priest has"). F7 then says "Couldn't check" rather than inventing one. Then regenerate the three rule files (`python -m pipeline coach-rules --build 1.60.1.70790`) and update `specrules_data_test.go`. A race gate in F7 is the alternative, and it would also serve the other classes' racial spells. |
| PRIEST-R26-3 | CONFIRMED | coach: spell named before it is trained (the R25-14 shape) | `coach/priest-discipline.md:24`; `coach/priest-holy.md:26` and `:48` (drill.preparation) | Disc ":24 Drop Lightwell and tell the raid it is up."; Holy ":26 Lightwell goes down before the pull"; Holy ":48 check flask, mana food, a Major Mana Potion on your bars and Lightwell placed." | Lightwell trains at 40 (`trainables/priest.json`, Lightwell 1 at 40). The coach reviews any boss pull at the player's level (round 25 D1), so on a dungeon boss at 10 to 39 these lines ask for a pre-pull spell the player does not have. Round 25 filed and fixed the same shape for Shadowform at 40 and 41 ("Shadowform from forty-two"). | Disc: `starter_overrides.py` GUIDE_LINES `priest-discipline`, "Lightwell": "From forty, drop Lightwell before the pull and tell the raid it is up." Optional: "Fear Ward" with "from twenty".<br>Holy hand edits:<br>- :26 "From forty, Lightwell goes down before the pull ..."<br>- :48 "... a Major Mana Potion on your bars and, from forty, Lightwell placed." |
| PRIEST-R26-4 | CONFIRMED | coach: false comparison at every band it applies to (a round 25 override) | `coach/priest-discipline.md:34` | "Inner Focus on your most expensive heal, from forty a Greater Heal. Its crit leaves the biggest Divine Aegis you can make." | From 40, Prayer of Healing costs more than the Greater Heal the card casts, at every band:<br>- 40 to 49: Prayer of Healing rank 1, 410, against Greater Heal rank 1, 370 (the tank filler is rank 1 to 57; Improved Healing only from 58).<br>- 50 to 57: Prayer of Healing rank 3, 770, against 370.<br>- 58 and 59: 770 against Greater Heal rank 4, 655 less 5% to 10%.<br>- 60: Prayer of Healing rank 4, 1030, against 655 less 15%.<br>So "your most expensive heal, from forty a Greater Heal" is false at every level it names. discipline.md:82 says "Inner Focus on a Greater Heal or a Prayer of Healing". Round 25's proposed text kept "or a Prayer of Healing", but the shipped override dropped it. "Its crit" reads as Inner Focus's. | `starter_overrides.py` GUIDE_LINES `priest-discipline` "Inner Focus": "Inner Focus on a Prayer of Healing from thirty or a Greater Heal from forty: it makes the spell free and adds twenty-five percent crit chance, and a Greater Heal crit leaves the biggest Divine Aegis you can make." |
| PRIEST-R26-5 | CONFIRMED | coach: drill contradicts the card, the page and its own line (a round 25 fix) | `coach/priest-discipline.md:51` (drill.resource) | "If you are behind it, keep Penance and the shield on the tank, let the filler wait for a deeper hole, and save Flash Heal for the emergency." | The card's shield row ([7]) fires only while `currentManaPercent >= 0.9 x remainingTimePercent`, so behind the pace it stops shielding. discipline.md:68 says the same: "Shield the tank ... when he is hurt, Weakened Soul is off him and your mana is keeping pace with the fight". The coach's own :27 line says "Your mana has to keep pace with the fight."<br>The drill tells a player who is behind the pace to keep shielding, which is the opposite. Penance has no pace gate, so that half is right. | `starter_overrides.py` HEALER drills `priest-discipline` "resource": "Before the pull, note the mana you should have at the halfway mark. If you are behind it, keep Penance on cooldown, hold the shield until you are back on pace, let the filler wait for a deeper hole, and save Flash Heal for the emergency." |
| PRIEST-R26-6 | CONFIRMED | coach: rank drill false at 46 to 57 (a round 25 fix) | `coach/priest-discipline.md:52` (drill.ranks) | "From forty, bind Greater Heal rank one for the party's small holes, and keep the top rank for the tank, who takes the big hits." | discipline.md:124: "From 40, Greater Heal on the tank stays at rank 1 until you learn rank 4 at 58 ... the bigger ranks spend mana on healing a dungeon tank rarely needs." The card agrees:<br>- Bands 40 to 56: tank filler row [15] and deeper-hole row [26] cast rank 1.<br>- From 58: rank 4.<br>At 46 to 57, the top rank is 2 or 3, and the drill moves the tank onto a rank the page and card hold back. | `starter_overrides.py` HEALER drills `priest-discipline` "ranks": "Before forty, bind the Heal rank your level calls for. From forty, bind Greater Heal rank one: it is the party's small-hole heal, and the tank's filler too until you learn rank four at fifty-eight. From then the tank gets the top rank." |
| PRIEST-R26-7 | CONFIRMED | coach: filler rank false at 46 to 57 | `coach/priest-holy.md:22` | "Heal is the filler until Greater Heal at forty. From then Greater Heal is the filler at your top rank. Step down to rank one when it would overheal or your mana runs short" | holy.md:128: "keep rank 1 as the filler until you learn rank 4 at 58 ... a dungeon tank rarely needs the bigger ranks". The card note ([10]) says "rank 1 until 57, then rank 3 or your top trainer rank", and bands 40 to 56 cast rank 1 on every Greater Heal row. "From then ... at your top rank" puts a level 46 to 57 Holy priest on rank 2 or 3 as the filler, with rank 1 only as the step down. | Hand edit: "Heal is the filler until Greater Heal at forty. Greater Heal rank one is then the filler until you learn rank four at fifty-eight. From then the filler is the top rank, stepped down to rank one when it would overheal or your mana runs short ..." The rest of the sentence is unchanged. |
| PRIEST-R26-8 | CONFIRMED | coach: rank drill contradicts the card before 40 and at 46 to 57 | `coach/priest-holy.md:47` (drill.ranks) | "Bind rank one of your filler, Heal before forty and Greater Heal from forty, and use it for every small hole. Save the top rank for a hole as deep as it heals." | The card casts one rank per band on the filler and the deeper-hole rows alike:<br>- Heal rank 3 at 34 to 37 for every race but a Troll at 34. holy.md:124: "From 34 rank 1 has started to lose its bonus ... keep rank 3 from 34 to 37".<br>- Rank 4, 3 or 2 at 38 and 39.<br>- Rank 2 or 3 at 27, 28 and 30 to 32 for the races listed.<br>- Greater Heal rank 1 on the deeper-hole rows at 46 to 57, where "a dungeon tank rarely needs the bigger ranks" (holy.md:128).<br>The drill sends the player to Heal rank 1 where the page has left it, and to the top rank for deep holes where the page keeps rank 1. | Hand edit: "Bind the filler rank your level calls for: Heal before forty, Greater Heal rank one from forty, and from fifty-eight Greater Heal rank one for the small holes, with the top rank saved for a hole as deep as it heals." |
| PRIEST-R26-9 | CONFIRMED | coach: half-mana cooldown drill against the card and its own :26 | `coach/priest-holy.md:44` (drill.core_casts, second sentence) | "When the mana bar drops below half, check your Major Mana Potion and Demonic Rune, and Dark Sacrifice if you are Undead, before your next cast." | The card fires the potion and the rune "whenever it is ready and the bar has room for it" (APL header, the autocast line). It fires Dark Sacrifice under 70% mana with health above 70% (row [1]). The persona's own :26 says "go out whenever the bar has room for their full return", and holy.md:84 says "once you are down enough that neither wastes".<br>Waiting for half delays all three past when the card and page spend them. The ledger's Round 25 section lists half-mana cooldowns among the drills removed. This hand-written drill kept it. The Discipline starter's drill now reads right. | Hand edit: "When the mana bar drops, check your Major Mana Potion and Demonic Rune, and Dark Sacrifice if you are Undead, and press each as soon as the bar has room for its full return." |

Counts, guide pages and cards: CONFIRMED 1 (R26-1), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

Counts, coach: CONFIRMED 8 (R26-2 to R26-9), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

- One rule-file change: R26-2 (`coachrules.py`, then a regen of the three rule files).
- Four generator overrides in `starter_overrides.py`, none fixed by a plain regen: R26-3 (Discipline part), R26-4, R26-5, R26-6.
- Five hand edits to `priest-holy.md`: R26-3 (Holy part), R26-7, R26-8, R26-9.
- R26-4, R26-5 and R26-6 are round 25's own override text. R26-1 reverts a round 19 fix. Run the note-revert check on 5dbffb0bf: it was a sim-lane edit, so the push guard took it as deliberate.

## Shoulds (wording, not findings)

- **PRIEST-R26-S1, Holy Heal note first sentence (priorityList[30], bands 16 to 39).** "Heal is the step down when Greater Heal rank 1 costs more than you can spend, and your filler before Greater Heal at 40." On bands 16 to 39 the addon's sentence opens with a spell the player will not have until 40. It is not false, since it names the level. Optional: "Heal is your filler before Greater Heal at 40, and then the step down when Greater Heal rank 1 costs more than you can spend."
- **PRIEST-R26-S2, Discipline Heal rows at bands 40 to 60 (data/platform lane).** The two tank Heal rows ([16], [27]) are gated by `not spellIsKnown(Greater Heal)`, so they never fire from 40. The card still renders them as Heal rank 4, under the note "Heal rank 1 on the tank through 35 and rank 3 from 36 ... rank 4 costs more mana than its extra healing is worth". The note's "until you learn Greater Heal at 40" keeps it true, but a rank 4 icon under a sentence against rank 4 reads oddly. Hiding rows whose `spellIsKnown` gate is false at the band would fix it.
- **PRIEST-R26-S3, Shadow "with Inner Focus" Mind Blast row at bands 10 to 52 (data/platform lane).** The row filter now drops the Inner Focus half of the strictSequence row until 57, but its Mind Blast half still renders with "with Inner Focus first whenever it is up". So the card shows two Mind Blast rows and names a talent learned at 57. "Whenever it is up" keeps it true (as round 25 accepted). Optional: drop the paired row where Inner Focus is filtered out.
- **PRIEST-R26-S4, Holy coach :38 "the Major Mana Potion, up to its use limit in a fight".** Classic potions have a 2 minute cooldown, not a per-fight limit. The starters say "as often as its timer allows"; the Holy persona could say the same. The same text is in `warrior-protection.md:37`, for the warrior lane.
- **PRIEST-R26-S5, Holy coach :38 and :48.** These omit the Brilliant Mana Oil that holy.md:96 names, as R25-S5 did for Discipline. This is an omission, not a false line.
- **PRIEST-R26-S6, Discipline coach list (:24 to :35).** The list leaves out the party heals the card casts: Lesser Heal rank 3 before 40, Greater Heal rank 1 from 40 (the drill names it), and party Flash Heal. A healer has no adherence check, so nothing is judged against the gap.
- **PRIEST-R26-S7, APL header notes (not player-facing).** `priest-holy.json` and `priest-discipline.json` still state the ranker rule "never one learned below level 20". Sim round 22 dropped that rule (P1). The Discipline header also says "Renew is left out on purpose", but row [8] casts it.
- **PRIEST-R26-S8, Shadow coach :50 (drill.core_casts, shared kit).** "Press the first button of the priority list the moment it is ready, before any filler": the first button on this list is Shadowform, a form you hold.

Carried, unchanged and optional: R25-S1 (Holy pace "at least as full" against the card's 90%), R25-S4 (healer lines in the Shadow Voice), R25-S5 (Discipline oil), R24-S1, R24-S3, R24-S4, R23-S1 to S12 (R23-S4: Discipline Renew by race, including a Human on rank 7 at 44 to 47 and the Horde on rank 7 at 56 to 58), R22-S1, R22-S2, R21-S1 to S4 (R21-S3: Undead and Dwarf Mind Blast rank 6 at 50 and 51 against shadow.md:137), R20-S1, R20-S2, R18-S1 to S5, R17-S2. Round 25's R25-S2 (Holy front matter build) and R25-S3 (Shadow Shoot line) are fixed.

## Checked and right this round

**Guide pages.**
- discipline.md and holy.md: the lines sim round 22 rewrote are checked above.
- The four pages are otherwise unchanged since round 25. Each comparison, cost, cooldown and rank claim that rounds 20 to 25 verified still holds on identical data, read again against the current rows:
  - Shadow Word: Death 18 s, 30% backlash.
  - Mind Blast 8 s.
  - Fade 10 s on a 30 s cooldown.
  - Lightwell 3 minutes or 5 clicks.
  - Penance 12 s, a pulse every 1 s.
  - Inner Focus 3 minutes, 25%.
  - Devouring Plague and Vampiric Embrace cooldowns equal to their durations.
- The leveling paragraphs match the shipped orders:
  - Holy: Twilight Focus 10 to 12, Improved Renew 13 to 15, Holy Specialization from 16.
  - Discipline: Twin Disciplines 10 to 14, Mental Strength 15 to 19, Penance 30.
  - Shadow: Wand Specialization 10 and 11, Spirit Tap 12 to 16, Shadow Focus 17 to 21, Improved Mind Blast to 25, Improved Shadow Word: Pain 26 and 27, Mind Flay 28, Improved Mind Flay 29 and 30, Blackout 31 and 32, Vampiric Embrace 33, Shadowform 42.
- shadow.md:137 rank lines match the Shadow rows:
  - Shadow Word: Pain rank 3 at 30 to 39.
  - Devouring Plague rank 1 at 30 to 43, rank 2 at 44 to 59, Human rank 1 at 50 and 51.
  - Mind Blast rank 4 at 40 to 43, Gnome excepted.
  - Shadow Word: Death rank 2 at 50 to 58.
- Sim round 22's Vampiric Embrace change (it cannot miss) moves no page line: no page claims it can miss.
- No page or rendered note names 70755.

**Coach lines OK at every level they apply to.** The Coach sweep appendix gives one line each.

## Decisions

- **D1. Coach scope by level** stands as in round 25 D1: a coach line must be true at every level a boss pull can be logged. A button line that holds itself back ("every time it is up", "whenever it is down", "once you have") is self-scoping. A pre-pull instruction or a rank instruction is not. So:
  - Lightwell before the pull (R26-3) is a finding, like round 25's Shadowform line.
  - Vampiric Embrace "whenever it is down" and Shadow's Inner Focus "every time it is up" are not.
- **D2. R26-2 is CONFIRMED** although no file names Desperate Prayer in prose. The rule file is in this round's coach scope, and F7 turns it into a player-facing sentence ("Desperate Prayer not used"). The client race mask decides it: Dwarf only.
- **D3. R26-1 is a card finding (the trend column), not a coach finding.** It reverts round 19's fix, which every round since kept as the condition that made the note true at 40 to 45.
- **D4. R26-S2 and R26-S3 stay Shoulds.** Each note stays true at its band. Only the rendered row beside it is odd, and the rows are the data lane's.

## Sims requested

None new. Already routed: Prayer of Mending heals per cast; Holy Greater Heal rank 1 behind the pace; the Discipline shield with no health gate; Shadow Word: Pain at 6 against 10 seconds. Sim round 22 closed Holy and Discipline Heal rank 1 against 2 at 22 to 27, Holy rank 2 against 3 at 28 and 29, and the Discipline tank Lesser Heal at 10 to 15.

## Engine questions

None new. P2 (heals scaling with the target's level), P3, P4 and S2 stay with the sim lane.

## Sim findings

None new.

## Sources

- **Client data 1.60.1.70790:** `trainables/priest.json` (Lightwell 40, Fear Ward 20, Greater Heal 40/46/52/58, Heal 16/22/28/34, Prayer of Healing 30/40/50/60), `talents/priest.json`, `addon-builds.json`, `addon-data.json` (rows and `race_spell_ids`, diffed against 438f8ba15), `spellranks.json`.
- **Client race masks:** client `SkillLineAbility.csv` 70790, from the scratchpad's `b70790/70790/`. Desperate Prayer 13908 and 19243 have RaceMasks_0 4 (Dwarf); Dark Sacrifice 1277328 has 16 (Undead); Devouring Plague 2944 has -1 (all).
- **Engine at bff826207:** `sim/priest/heal_ranks.go` (ranks, costs, coefficients and MaxLevels), `sim/core/spell.go` (`levelPenaltyGrace` 12, 5% a level), `data/curated/book-ranks.json` (Greater Heal 25314 is a book rank).
- **Curated:** `data/curated/apl/priest-{discipline,holy,shadow}.json`, `data/curated/coach/priest-*.{md,json}`, `data/pipeline/coachrules.py`.
- **Generator:** `tools/coach/persona_starter.py`, `starter_overrides.py` (GUIDE_LINES and GUIDE_ORDER, and the healer drills for `priest-discipline`). Regen output: `scratchpad/gr26p/regen/`, identical to the tree.
- **Coach checks:** `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go` and `f7_death.go` (`defensivesAt`, `TitleDiedNotUsed` in `copy.go`).
- **Reviews:** round 25 priest; ledger Rounds 19 to 25; sim round 22 priest and integration.
- **No web research needed:** the client tables, the cards and the pages decided every call.

## Appendix: Note sweep

Bands are as rendered. Levels are when the player has each named item, on the shipped order or from the trainer. The healer raid orders equal their leveling orders. For Shadow, both orders were checked (Mind Flay 28 and 21, Improved Mind Flay 29 and 22, Vampiric Embrace 33 and 28, Shadowform 42 and 42, Inner Focus 57 and 57).

### Holy (bands 10, 16, 20, 26, 27, 28, 29, 30, 32, 33, 34, 35, 38, 39, 40, 44, 45, 50, 51, 52, 53, 56, 58, 59, 60)

- Renew (10 to 60): "at 60 that means rank 7 or higher". OK.
- Lesser Heal rank 3, tank (band 10): "until you learn Heal at 16", rank 3 (10), pace. OK.
- Heal, tank (16 to 60), sim round 22's note:
  - First sentence: S1.
  - "Below 34 cast rank 1 ... some races cast rank 2 or 3": OK.
  - "whole healing bonus until then" (whole through 33): OK.
  - "heals the most for its mana" (band gear 144 and up, above the 69 to 149 crossings): OK.
  - "28 and 32 for most races, at 27 for a Human, Gnome or Night Elf, at 30 and 31 for an Undead or Troll": OK on every row.
  - Rank 4 "in its first levels", "rank 3 from 34 to 37 and to 38 as an Undead", "a Troll stays on rank 1 at 34 and casts rank 2 at 38 and 39, an Undead rank 2 at 39": OK.
  - "From 40 ... Heal rank 4 ... cheaper (305 against 370) ... about half your bonus at 60 (55%)": OK.
- Heal, deeper hole (16 to 60): "When your mana falls behind ... a tank under 50% or a party member under 35%". OK.
- Dark Sacrifice (20 to 60): Undead only, the 70% gates, 15 s, Spirit, 10 minutes. OK.
- Flash Heal, tank (20 to 60): top rank; "heals less for its mana than your slower heals". OK at every band, including against Heal rank 1 at 22 to 33 (2.6 against 2.1).
- Flash Heal, party (20 to 60): under 30%. OK.
- Prayer of Healing (30 to 60): three of four under 85%; the target's party. OK.
- Prayer of Mending (40 to 60): no gate, 10 s, up to five, not paced. OK; "cheapest" is routed.
- Greater Heal filler (40 to 60): "rank 1 until 57, then rank 3 or your top trainer rank" (58 rank 4, 59 rank 3, 60 rank 4). OK.
- Greater Heal, deeper hole (40 to 60): OK.
- Greater Heal rank 1 (40 to 60): **R26-1**.
- Greater Heal rank 1, deeper hole (40 to 60): "When your mana falls behind ... so it waits for a deeper hole too". OK.

### Discipline (bands 10, 12, 16, 18, 20, 24, 26, 30, 32, 34, 36, 38, 40, 42, 44, 45, 48, 50, 54, 56, 58, 59, 60)

- Power Word: Shield (10 to 60): hurt, Weakened Soul off, pace, no rage, Renewed Hope (32), Divine Aegis (36). OK.
- Renew (10 to 60): rank 7 or higher at 60. OK.
- Lesser Heal rank 3, tank (10 and 12, new): "until you learn Heal at 16"; "shield and Renew alone will not keep him up". OK.
- Lesser Heal, party (10 to 38): "before 40"; "late forties ... lost your healing bonus" (MaxLevel 15, gone by 47). OK.
- Heal, tank (16 to 38, new note): rank 1 to 35, rank 3 from 36; whole bonus to 33; per mana; rank 4. OK. The same row at 40 to 60 is S2.
- Dark Sacrifice (20 to 60): as Holy. OK.
- Flash Heal, tank (20 to 60): "the fast heal, and the dear one". OK.
- Flash Heal, party (20 to 60): the 0.7 pace gate. OK.
- Penance (30 to 60): 12 s; three heals 1 s apart; whoever is hurt; "cheap for all it heals". OK.
- Prayer of Healing (30 to 60): "Rank 4 is the top rank a trainer teaches". OK.
- Greater Heal, tank (40 to 60): "at the top rank ... at 60". OK.
- Greater Heal rank 1, party (40 to 60): a third; "once you have higher ranks" (46); far less mana. OK.
- Greater Heal deeper-hole rows, tank and party (40 to 60): OK.

### Shadow (bands 10, 16, 18, 20, 22, 26, 28, 30, 32, 33, 34, 36, 38, 40, 42, 44, 46, 48, 50, 51, 52, 57, 58, 60)

- Shadow Word: Pain (10 to 60): "Always up; stop refreshing it in the closing seconds". OK.
- Mind Blast with Inner Focus (10 to 60; the Inner Focus row only from 57): "whenever it is up"; free; 25%. OK (S3).
- Mind Blast (10 to 60): "ahead of Mind Flay once you have it" (28); "half cost in Shadowform" (42). OK.
- Dark Sacrifice (20 to 60): below a fifth and behind; Undead only. OK.
- Devouring Plague (20 to 60): 24 s at every rank; every priest; ticks crit; heals. OK.
- Mind Flay (28 to 60): a clipped channel loses the rest; "once you have Improved Mind Flay" (29). OK.
- Vampiric Embrace (33 to 60): "whenever it is down"; 30 s; 20%; "until you have the talent". OK.
- Shadow Word: Death (32 to 60): 30%; "once you have Vampiric Embrace" (33). OK.
- Shoot (curated note, no rendered row): OK.

## Appendix: Coach sweep

One line per coach line, at every level a boss pull can be logged (D1).

### priest-discipline.md (generated; regen identical)

- :18 overview (discipline.md:26): OK.
- :22 lead-in: OK (no meta talk).
- :24 Lightwell: **R26-3**.
- :25 Fear Ward (20; "if the boss fears"): OK.
- :26 Penance from thirty: OK.
- :27 Power Word: Shield, pace: OK.
- :28 Renew: OK.
- :29 Flash Heal on the tank: OK.
- :30 Prayer of Healing from thirty: OK.
- :31 Greater Heal top rank at sixty: OK.
- :32 Heal ranks (matches the note): OK.
- :33 Lesser Heal on the tank to 16: OK.
- :34 Inner Focus: **R26-4**.
- :35 Dark Sacrifice: OK.
- Missing party heals: S6.
- :38 to :43 mistakes: OK.
- :45 consumables (Distilled Wisdom, Sage's Tea, Major Mana Potion, as in the preset): OK.
- :49 drill.core_casts: OK.
- :50 drill.uptime: OK.
- :51 drill.resource: **R26-5**.
- :52 drill.ranks: **R26-6**.
- :53 drill.preparation: OK (R25-S5).
- :54 drill.death: OK.
- :56 to :67 voice: OK.

### priest-shadow.md (generated; regen identical)

- :18 overview (shadow.md:30): OK.
- :22 lead-in: OK.
- :24 Shadowform from forty-two: OK.
- :25 Shadow Word: Pain: OK.
- :26 Vampiric Embrace (self-scoping): OK.
- :27 Devouring Plague: OK.
- :28 Inner Focus (57, "every time it is up"): OK.
- :29 Mind Blast: OK.
- :30 Shadow Word: Death, "once you have Vampiric Embrace": OK.
- :31 Mind Flay, "once you have Improved Mind Flay": OK.
- :32 Shoot: OK.
- :33 Fade 10 s on a 30 s cooldown: OK.
- :34 Dark Sacrifice: OK.
- :37 to :44 mistakes: OK.
- :46 consumables (Major Spellblasting Potion, Supreme Power, Baked Salmon, Brilliant Wizard Oil, as in the preset): OK.
- :50 drill.core_casts: S8.
- :51 to :57 the other drills: OK.
- :61 to :70 voice: R25-S4.

### priest-holy.md (hand-written)

- :18 who they are (the level 60 picture): OK.
- :22 the priority, in words:
  - The order, tank first; Prayer of Mending; Renew rank 7 or higher; Flash Heal; Prayer of Healing on the target's party: OK.
  - Greater Heal at the top rank from forty: **R26-7**.
  - Heal rank 4 the last step down: OK.
- :24 pacing: R25-S1. Never paced: OK.
- :26 cooldowns:
  - Inner Focus (57, self-scoping): OK.
  - Dark Sacrifice (Undead): OK.
  - Potion and Rune with room: OK.
  - Lightwell: **R26-3**.
- :28 what changed for Forever: OK.
- :31 to :36 mistakes: OK.
- :38 consumables: S4, S5.
- :40 ten against forty: OK.
- :44 drill.core_casts: Inner Focus from fifty-seven, OK; the half-mana sentence is **R26-9**.
- :45 drill.uptime: OK.
- :46 drill.resource: OK.
- :47 drill.ranks: **R26-8**.
- :48 drill.preparation: **R26-3** (Lightwell).
- :49 drill.death: OK.
- :53 to :62 voice: OK.
- Front matter: 70790, OK.

### Rule files

- `priest-discipline.json`:
  - core [] (healer, judged on efficiency): OK.
  - efficiency Inner Focus (talent, skipped if never cast): OK.
  - aura Inner Fire (12): OK.
  - downranks: unused by the healer F5 branch.
  - defensives Desperate Prayer: **R26-2**.
- `priest-holy.json`: as Discipline. Downranks Renew is unused; defensives **R26-2**.
- `priest-shadow.json`:
  - core Shadow Word: Death (32, guarded by `notYetLearned`), Devouring Plague (20), Inner Focus (talent): OK.
  - auras Shadow Word: Pain (dot), Shadowform (talent, never judged unless it was up): OK.
  - defensives **R26-2**.

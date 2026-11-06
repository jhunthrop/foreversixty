# Guide accuracy round 26: Paladin

Audited on the round 26 tree (`scratchpad/gr26/audit`, head ca06416cd: main after guide round 25, which carries sim round 22 through fd758eefc). No sim round 23 in the log. Engine pin bff826207 (`sim/enginever/version.go`); engine source `scratchpad/gr26/engine`. Active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).

Files audited:
- `web/src/content/guides/paladin/{index,holy,protection,retribution}.md`, line by line.
- Every note in `data/curated/apl/paladin-{holy,protection,retribution}.json`, against every rendered band of `addon-data.json` `rotations/paladin-*`, both shipped orders in `addon-builds.json` (decoded by level), `trainables/paladin.json`, `spellconst/paladin.json`, `spelltips/paladin.json`, `talents/paladin.json` and `rank-choices.json`.
- Coach, read again from scratch: `data/curated/coach/paladin-{holy,protection,retribution}.md` and `.json`; the paladin entries of `data/pipeline/coachrules.py`; the generator (`tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`, `starter_priority.py`, `starter_guide.py`); the checks that read the rule files (`api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go`, `f7_death.go`, `builddata.go`, `copy.go`); `sim/request/raid-consumes.json`.
- Engine: `sim/core/spell.go` (`SpellLevelPenalty`) and `sim/core/spell_max_level_auto_gen.go`, for the healing-power claims.

Read first: round 25 paladin, the ledger's Round 25 section, sim round 22 `paladin.md` and `integration.md`, round 26 druid (for the percent-of-base-mana and cover lenses).

## Changes since round 25's tree (438f8ba15)

None of this was audited by a guide round:
- **Sim round 22 (H1, H3):** holy.md step 5 (party Flash of Light rank 5 in a raid) and the Levels 1 to 19 paragraph (Holy Light rank 1 from 18); Holy card row [4] gains "While you level, Holy Light rank 1 from 18 to 24 and rank 2 from 25 to 29 keep all of your healing power for a fraction of the top rank's mana"; row [12] gains "and rank 5 at 60 in a raid"; `rank-choices.json` Holy Light 18 to 19 rank 1 and party Flash of Light raid 60 rank 5.
- **928ddcf14 (guide round 25 integration):** Holy row [3] note drops "(from 21)" after Divine Favor. **5dbffb0bf:** the self-target Holy Shock row note folds into row [5].
- **Round 25 coach fixes:** all three starters regenerated; `paladin-protection.json` drops Swift Judgement from core; `paladin-retribution.json` drops the Seal of Righteousness aura; F1 `notYetLearned` (trainer spells never cast below their level are not judged).
- **`trainables/paladin.json`:** sim round 22's platform fix gives six talent rows a level (Divine Favor 20, Templar's Bulwark 20, Oath of Light 25, Swift Judgement 25, Anticipation 25, Repentance 30). These are tier levels; the rendered rows still follow the shipped orders (Divine Favor 21, Swift Judgement 29), and every coach line follows the order.
- Protection and Retribution pages, cards and rank rows: unchanged.

**Coach regeneration check.** `persona_starter.py --spec paladin-<spec> --out scratchpad/gr26/paladin/regen/` for all three: byte-identical to the committed files. `test_persona_starter.py` (run with no cache, read-only): 55 passed. No paladin persona is hand-written.

## Round 25 fixes, rechecked

- R25-1 to R25-4, R25-6, R25-10, R25-11, R25-13, R25-14, R25-16: each is in the regenerated file or rule file and holds.
- R25-5 and R25-12: the named lines are scoped. The same lens finds lines the fix did not reach (R26-9, R26-10).
- R25-7: `assessAbility` now skips a never-cast trainer spell below its level (`f1_trainer_level_test.go`). Holds.
- R25-8: Swift Judgement is out of Protection's core. Holds.
- **R25-9: incomplete.** The proposed override had two sentences; the shipped one keeps only the first, so it now tells an Unsealed Wrath raider to stay in Righteousness (R26-12).
- **R25-15: incomplete.** It scoped the Holy Shock entry; Holy Shock is still named unscoped in the Holy Light line (R26-9).
- **Round 25's SHOULD "Protection Judgement and Hammer of the Righteous are judged at base cost only" was wrong.** Both cost a percent of base mana (`cost: 0`, `cost_pct: 6.0`), so `ParseSpellConst` marks the cost unknown and F1 does not read mana for them at all (R26-5, R26-6).

## Findings

Counts:
- **Guide pages and cards:** CONFIRMED 4, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.
- **Coach:** CONFIRMED 8, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| PALADIN-R26-1 | CONFIRMED | page and card note: false value at every level it names (sim round 22 text) | holy.md:113; `paladin-holy.json` row [4] note (renders 10 to 60) | holy.md: "From 18 ... put Holy Light (Rank 1) on its own key and heal with it: at 35 mana against rank 3's 110 it carries all of your healing power"; row [4]: "Holy Light rank 1 from 18 to 24 and rank 2 from 25 to 29 keep all of your healing power" | Forever's low-rank rule as the engine has it (`sim/core/spell.go` `SpellLevelPenalty`: whole until 12 levels past the rank's client MaxLevel, then 5% less a level; sourced in that comment to ForeverChanges' downranking measurements and HealRankAdvisor). Holy Light rank 1 (635) has client MaxLevel 5 (`spell_max_level_auto_gen.go`, `spellconst` max_level 5), the comment's own "Holy Light rank 1 running one level early": it keeps 95% of your healing power at 18, 90% at 19, 85% at 20 and 65% at 24. Rank 2 (639, MaxLevel 11) keeps 90% at 25 and 70% at 29. "All" is false at every level both lines name. The sim result itself (rank 1 +11 to +16% at 18 and 19) is right: it is the mana that wins, not a whole coefficient. Same paragraph and step 4 say "Forever takes the healing power off ranks learned far below your level", so the page also contradicts itself. | holy.md:113: "...put Holy Light (Rank 1) on its own key and heal with it: at 35 mana against rank 3's 110 it heals far more per point of mana, even though it keeps a little less of your healing power each level from 18, and over a whole dungeon you heal more before you drink." Row [4], second sentence: "While you level, Holy Light rank 1 from 18 to 24 and rank 2 from 25 to 29 heal far more per point of mana than your top rank, though each keeps less of your healing power every level." Regenerate addon-data (note text only). |
| PALADIN-R26-2 | CONFIRMED | page: contradicts the rendered card | holy.md:113 | "From 30 your newest rank takes over." | The card casts Holy Light rank 4 on both Holy Light rows at 40 to 42 and at 47 to 49 (`rank-choices.json` paladin-holy Holy Light 40-42 and 47-49 rank 4; `addon-data.json` bands 40 and 47 show rank 4, where the newest is rank 6 at 40 and rank 7 at 47). | "From 30 your newest rank takes over for most of your leveling; at 40 to 42 and 47 to 49 the card keeps rank 4, which heals as much for less mana." |
| PALADIN-R26-3 | CONFIRMED | card note: credits a talent before the shipped order (regression of a round 19 fix) | `paladin-holy.json` row [3] (Holy Light, tank below 40%), renders 10 to 60 | "On the tank below 40%, right behind Divine Favor." | Divine Favor is the 21st-level point on the Holy order (decoded `addon-builds.json`); the Divine Favor row first renders at band 21 (`addon-data.json`). At bands 10, 14, 18 and 20 the row is shown "right behind" a button the player does not have. Round 19 scoped it "(from 21)"; 928ddcf14 dropped the scope as "a leveling aside". | "On the tank below 40%, right behind Divine Favor once you have it at 21." Or move "right behind Divine Favor" into a `level_notes` entry from 21. Regenerate addon-data. |
| PALADIN-R26-4 | CONFIRMED | page: contradicts the rendered card (sim round 22 text) | holy.md:113 | "Rank 3 still heals more per cast, so keep it for the hit the tank must survive." | The card's emergency row is row [3] (tank below 40%, and from 21 straight after Divine Favor). The 18 to 24 rank row (authored id 10329) applies to both Holy Light rows, so the card casts rank 1 for that hit at 18 to 24 (`addon-data.json` bands 18, 20, 21: Holy Light rank 1 on both rows), including the Divine Favor crit from 21. The page tells the player to do what the card does not. Sim round 22 H3 forced the rank on every Holy Light cast; it never measured the two rows apart. | Sim first (Sims requested S1). If rank 3 (rank 4 from 22) on row [3] wins, the card needs row [3] authored on its own id so a rank row can split them, and the prose stands. If not, drop the sentence: "Rank 3 heals more per cast, but over a fight rank 1 heals more, even when the tank is low." |
| PALADIN-R26-5 | CONFIRMED | coach rule file: core abilities judged as lost "while you had the mana" when the page holds them for mana, and the coach cannot read their cost | `coachrules.py` `paladin-protection` core "Hammer of the Righteous", "Judgement"; `f1_core_casts.go` `newResourceCheck`; `builddata.go` `ParseSpellConst` | F1 detail: "{lost} of a possible {possible} went unused while you had the mana." + "Couldn't check whether you had the mana." | Both spells cost 6% of base mana with no flat cost (`spellconst` 407632 and 20271: `cost: 0`, `cost_pct: 6.0`). `ParseSpellConst` leaves `CostKnown` false, `newResourceCheck` returns `unverified`, and `readyMSInGap` counts every second off cooldown as ready. protection.md:95 and :97 cast both only "once Holy Shield's mana is safe", :101 "If you are going dry, cut Consecration first, then Hammer of the Righteous"; card rows 8 and 10 gate both on mana less Holy Shield's cost. A tank who holds them, as the page says, is told the casts were lost while he had the mana, and the appended "Couldn't check" contradicts that sentence. Sim round 22 P2 measured that mana limits Protection at every leveling band. The round 22 lens applies too: even with a known cost, F1 checks base cost, not the Holy Shield reserve. | Rule file: `paladin-protection` core = ["Holy Shield", "Holy Strike"], comment "Hammer of the Righteous and Judgement cost a percent of base mana and are held for Holy Shield's reserve (protection.md); F1 can read neither (round 26)." Rerun `python -m pipeline coach-rules`. The coach-code alternative (read `cost_pct` against base mana at `Row.Level`, shared with DRUID-R26-3) still would not see the reserve, so drop them here. |
| PALADIN-R26-6 | CONFIRMED | coach rule file: a core ability whose cost the coach cannot read | `coachrules.py` `paladin-retribution` core "Judgement"; same code path as R26-5 | as R26-5 | Judgement 20271 costs 6% of base mana (cost 0). Retribution is the leveling spec and runs dry on long pulls (retribution.md:47, :97; sim round 22 R2: the paced spend overspends at bare 60 and at most leveling bands). A paladin who could not Judge for lack of mana is told the casts went unused "while you had the mana". Judgement is the page's first button, so the check is worth keeping once the cost is readable. | Coach code (shared with DRUID-R26-3): F1 reads `cost_pct` times the class's base mana at `Row.Level` (gametables) as the cost, and checks the mana track against it. Until then, drop "Judgement" from `paladin-retribution` core (keep Holy Strike, flat cost) and rerun `coach-rules`. |
| PALADIN-R26-7 | CONFIRMED | coach check: a defensive claimed ready while Forbearance locks it out (all three specs) | `f7_death.go` `defensivesAt`; rule files' defensives Divine Shield, Divine Protection (and Templar's Bulwark on Protection) | F7 title: "Died at {t} with {Spell} ready"; detail "{Spell} had been ready since {since}." | Divine Shield, Divine Protection, Templar's Bulwark and Blessing of Protection each "Applies Forbearance for 1 min. Cannot be cast while Forbearance is active" (`spelltips/paladin.json`; Forbearance 25771, 60 s; 50 s with Sacred Duty). `defensivesAt` reads only each spell's own cooldown: a never-cast Divine Protection is "proven ready" from 5:00 into the pull, and nothing reads the Forbearance aura. A Retribution paladin who bubbles at 5:00 and dies at 5:40, or a tank who uses Templar's Bulwark and dies inside the minute, is told Divine Protection (or Divine Shield) was ready. The Protection persona (:36), protection.md:107 and card rows 2 and 3 all say they wait for Forbearance. | Coach code: in `defensivesAt`, a paladin defensive whose tooltip applies Forbearance is not ready while Forbearance (25771) is on the player at `atMS` (read from `in.Fight.Auras`), nor within its duration of the player's last Forbearance-applying cast or a Blessing of Protection received. Test first: a Retribution fixture with Divine Shield at 5:00 and a death at 5:40. |
| PALADIN-R26-8 | CONFIRMED | coach: drill contradicts the page and card after sim round 22; the list keeps an orphan line | `coach/paladin-holy.md:51` (drill.ranks, `DRILL_LINES["paladin-holy"]["ranks"]`); `coach/paladin-holy.md:29` | :51 "at sixty, Flash of Light rank five on the tank, rank four on everyone else"; :29 "In a raid, Blessing of Wisdom, Mana Spring and the raid's Intellect buffs pay for the bigger heal." | holy.md:79 (sim round 22 H1): "Flash of Light rank 5 on everyone else below 80% ... In a raid ... pay for the bigger heal. Without them, rank 4"; `rank-choices.json` party Flash of Light raid 60 rank 5; card row [12] "rank 5 at 60 in a raid". The coach reads raid logs, so the drill teaches the rank the card left. The persona never states the rule itself: the generator took step 5's second sentence as a "Blessing of Wisdom" button line, which points at "the bigger heal" that no line names. | `DRILL_LINES["paladin-holy"]["ranks"]`: "Bind the ranks you use: at sixty, Flash of Light rank five on the tank, and in a raid on everyone else too; without raid buffs, rank four on the party. Holy Light at its top rank. Rank one is a trap at sixty." GUIDE_LINES `paladin-holy` "Flash of Light", add: "At sixty in a raid that is rank five on everyone; without raid buffs, rank four on the party." EXCLUDED `paladin-holy` += "Blessing of Wisdom" (the orphan). Regenerate. |
| PALADIN-R26-9 | CONFIRMED | coach: spells named before the player has them (round 22 lens) | `coach/paladin-holy.md:27`, :25, :32 | :27 "Holy Light on the tank below seventy percent, and Holy Shock on anyone below seventy percent, but only with mana to spare"; :25 "Concentration Aura if nobody else in your group runs it."; :32 "Voice of Truth makes you immune to silences and interrupts..." | Holy Shock is the 32nd-level point (Holy order); Concentration Aura trains at 22, Voice of Truth at 20 (`trainables`). holy.md:117 levels Holy from 10 (Spiritual Focus), so dungeon boss pulls at 10 to 31 are coached against these lines. R25-15 scoped the Holy Shock entry (:30) but not this one; the card's Holy Shock rows render from 32. | GUIDE_LINES `paladin-holy`: "Holy Light": "Holy Light on the tank below seventy percent, and from thirty-two Holy Shock on anyone below seventy percent, but only with mana to spare: keep your mana percent no more than five points under the percent of the fight left."; "Concentration Aura": "Concentration Aura from twenty-two if nobody else in your group runs it."; "Voice of Truth": "From twenty, Voice of Truth makes you immune to silences and interrupts for six seconds, every three minutes: save it for the caster pack or boss that would lock you out." Regenerate. |
| PALADIN-R26-10 | CONFIRMED | coach: spells named before the player has them (round 22 lens) | `coach/paladin-protection.md:31` to :36, :59 (drill.death) | ":31 Exorcism on Undead and Demons."; ":32 Consecration only above half mana..."; ":33 ... Blessing of Kings before then."; ":35 Templar's Bulwark first, at about a third of your health"; ":36 Divine Shield and Divine Protection come after Templar's Bulwark" | Exorcism, Consecration and Blessing of Kings train at 20, Divine Shield at 34 (`trainables`); Templar's Bulwark is the 20th-level point (Protection order). protection.md:133 levels Protection from 10, so pulls at 10 to 19 meet every one of them, and 20 to 33 meet Divine Shield. Card row 1 scopes "Divine Shield from 34"; the persona does not. The Retribution persona scopes the same Exorcism and Consecration (R25-12). | GUIDE_LINES `paladin-protection`: "Exorcism": "Exorcism from twenty, on Undead and Demons."; "Consecration": "Consecration from twenty, only above half mana, to spend mana you will not need."; "Greater Blessing of Kings": "...Blessing of Kings from twenty before then..."; "Templar's Bulwark": "Templar's Bulwark from twenty, first, at about a third of your health: a full health bar of absorb for eight seconds."; "Divine Shield": "Divine Shield from thirty-four and Divine Protection come after Templar's Bulwark, once Forbearance is gone." DRILL_LINES death: "From twenty, put Templar's Bulwark on a key...". Regenerate. |
| PALADIN-R26-11 | CONFIRMED | coach: the core button's line teaches only the pull (R25-11 shape) | `coach/paladin-protection.md:30` | "Pull with Judgement: under Seal of Fury it taunts, so the boss comes straight to you." | protection.md:97 item 6 "Judgement on cooldown, once Holy Shield's mana is safe"; card row 10 "On cooldown, and once you have Holy Shield at 30, only while its mana is safe". The persona's only Judgement line is the opener sentence (:88), so the list never says to press it on cooldown or to hold it for Holy Shield, and Judgement is a core entry F1 judges (R26-5). Round 25 filed the same shape for Retribution and passed this line. | GUIDE_LINES `paladin-protection` "Judgement": "Pull with Judgement: under Seal of Fury it taunts, so the boss comes straight to you. After that, Judgement on cooldown, and from thirty only while Holy Shield's mana is safe." Regenerate. |
| PALADIN-R26-12 | CONFIRMED | coach: override drops the page's "while leveling" scope (R25-9 fix incomplete) | `coach/paladin-retribution.md:29`; `GUIDE_LINES["paladin-retribution"]["Seal of Righteousness"]` | "Seal of Righteousness at the top rank you have trained, straight after the opening Crusader Judgement. Recast it in its last second." | retribution.md:78 scopes this to "while leveling", then "Once you have Unsealed Wrath, go straight to Seal of Command"; card row 2 (with Unsealed Wrath) goes to Seal of Command after the Crusader Judgement, rows 6 to 8 keep Seal of Command up and use Righteousness for one swing. The raid order takes Unsealed Wrath at 40, and raid logs are where the coach is used most. An Unsealed Wrath paladin reads :29 and :30 as two different Seals for the same moment; :29 tells him to keep Righteousness up and recast it, which is the page's mistake "Twisting ... " turned around. The persona also never states the twist's release or its mana pace (:93, :139, :140). | GUIDE_LINES "Seal of Righteousness": "Seal of Righteousness while leveling, at the top rank you have trained, straight after the opening Crusader Judgement. Recast it in its last second. With Unsealed Wrath it is the one swing that releases three stored Commands, and your Seal whenever your mana falls behind the fight." "Seal of Command": add "Twist only while your mana is at least four fifths of the fight left." Regenerate. |

## Shared roots (filed elsewhere; named once for paladin, not counted)

- **F6 flask by role (DRUID-R26-6):** Retribution's preset has no flask at any band and the persona names none, but F6 expects one for a dps role; Holy and Protection get a flask only from 50, and every sub-50 paladin log is "missing" one.
- **Percent-of-base-mana costs (DRUID-R26-3):** the code half of R26-5 and R26-6. Other paladin buttons with a percent-only cost: Righteous Fury (30%), Blessing of Kings (8%), Divine Favor (4%), Oath of Light (8%). None is in a paladin rule file, so nothing else misfires today.
- **Cover (the f2_uptime.go `CoveredBy` fix in the int tree):** no paladin rule judges an aura another class can hold (Protection's only aura is Holy Shield on the player; Retribution and Holy have none). Nothing to add.

## SHOULDs (wording, not findings)

- **holy.md:113** "Until 17, rank 3 is your Holy Light for everything": rank 3 trains at 14 and the card casts rank 2 at 10 to 13; "From 14 to 17" would match.
- **coach/paladin-holy.md:18** "cheap Flash of Lights" (overview copy; sim round 22's wow-player said the same of holy.md:29).
- **coach/paladin-holy.md:27 and :30** repeat Holy Shock's mana rule; **:33** Lay on Hands drops "It drains all of your own mana" (carried).
- **coach/paladin-holy.md** drill.uptime "Recast your self buff" and the "Key buff and debuff uptime" and "Core abilities under-used" lines; **coach/paladin-retribution.md** "Key buff and debuff uptime" and drill.uptime: both rule files now have empty `auras`, so F2 can only report that it cannot check. The lines describe a check that never runs for these specs.
- **coach/paladin-protection.md** drill.uptime "a stacking debuff that falls off takes its stacks with it": the Protection rule judges no debuff (carried).
- **Kit:** Protection (Brilliant Wizard Oil from 45) and Holy (Brilliant Mana Oil from 45) list no weapon enchant, though their presets carry one.
- **Retribution "Who they are"** "Judge every eight seconds" (from 16, carried); **Protection "Who they are"** "Holy Strike every six seconds" (from 40, carried).
- **coach/paladin-retribution.md:32** "Do not let your mana fall far behind the fight" (carried).
- Carried from rounds 24 and 25, still true as SHOULDs: index.md:29 "for raiders"; Retribution row 11 parenthetical; "in a dungeon or a raid" (row 2, :78, coach :30); retribution.md:135 "when" clause; Human "from about 55"; retribution.md:43 "real damage"; row 6 "180 to 190"; retribution.md:72, :123, :36, :133; Holy row 12 and holy.md:81; protection.md:111 block chance and the Holy Shield ranks 2 and 3 tooltips (data lane).

## Checked and right this round

- **Sim round 22's other edits.** holy.md:79 step 5: rank 5 at raid 60 (row), rank 4 from 50 while leveling (default rule under authored rank 4, rank 5 trains at 50), "even once you have rank 5" holds. Row [12]'s new clause holds at every band. `rank-choices.json` Holy Light 18 to 19 and 20 to 24 rank 1, 25 to 29 rank 2 match holy.md:113's rank levels and the rendered bands 18, 20, 21, 25.
- **Page prose, full sweep.** Protection, Retribution and index are unchanged since round 25; re-read against the unchanged rows and the 70790 tooltips (Benediction all instants 10%, Improved Judgement 8 s from 16, Sanctified Judgement 60% at 22, Sacred Duty 60 s and 10 s off Forbearance, Shieldsworn 4 s, Iron Creed 6% for 6 s, Holy Shield four charges, Templar's Bulwark 100% for 8 s, Exorcism instant on a 15 s category cooldown, Consecration rank 1 at about 121 after Benediction and its level penalty: whole to 37, 50% at 47, 10% at 55). Holy: the rest of step 4, the Ranks, Mana, Light's Vigil, Movement, Silences and Emergencies paragraphs, Levels 20 to 60 against the Holy order (Oath 25, Infusion 27, Divine Precision 31).
- **Coach lines that hold** (Appendix B): Protection :24 to :29, :34, :37, consumables; Retribution :24 to :28, :30 to :33, consumables; Holy :24, :26, :28, :31, :33, :34, consumables.
- **Coach rule files:** Protection core Holy Shield (talent, flat cost, 10 s category cooldown) and Holy Strike (flat cost; 10 s in data, 6 s with Shieldsworn, so F1 undercounts and never overclaims); Retribution Holy Strike; Holy downranks; all three defensive lists are level-gated in F7 apart from R26-7. F5 honours `rank-choices.json` (Protection Holy Strike rank 1 at 40 to 42, Consecration rank 1; Retribution Consecration rows).

## Decisions

- **D1. R26-1 is CONFIRMED on the engine's rule**, not a sim: the claim is a statement about the coefficient, and the engine's level penalty with the client MaxLevel (5 for rank 1) gives 95% at 18 and 65% at 24. The page's own step 4 and stat paragraph say ranks far below your level lose healing power.
- **D2. R26-4 is CONFIRMED, its fix waits on a sim.** The contradiction is on the page today whichever side moves.
- **D3. R26-5 drops the core entries rather than asking for the cost fix.** Even with a readable cost, F1 checks base cost, not "Holy Shield's mana is safe".
- **D4. R26-9 and R26-10 include short windows** (Concentration Aura at 20 and 21) because round 25 set the rule that a coach line names no spell before the player has it, and the card scopes the same spells. Round 25 passed Concentration Aura as "a group-aura line"; that is reversed here.
- **D5. Not re-filed:** round 24's D4 list (rows rendering before their talent, card action ranks, the rank 9 Holy Light book, Exorcism never crits, Benediction and Hammer of Wrath, BiS race captions); the site card's per-race rank issue; sim round 22's open R2 and P2 (mana pacing, platform value).

## Sims requested

- **S1 (R26-4).** Holy at 18 to 24 (both factions, seeds 1 and 50001): row [3] (tank below 40%, behind Divine Favor from 21) at rank 3 (18 to 21) and rank 4 (22 to 24) against rank 1, with row [4] kept at rank 1. Needs row [3] authored on its own id.
- Round 24's optional runs stand.

## Engine questions

- None.

## Sim findings

- None.

## Sources

No web research was needed: the client data, the engine's low-rank rule, the card rows and the coach code decided every call. The engine's own citations for the low-rank rule (ForeverChanges, foreverchanges.pro/downranking; HealRankAdvisor, github.com/mendeleevforever/HealRankAdvisor) are in `sim/core/spell.go` at the pin.

- Client 1.60.1.70790: `spellconst/paladin.json` (Judgement 20271 and Hammer of the Righteous 407632 at `cost_pct` 6.0; Holy Light 635 max_level 5, 639 max_level 11; Forbearance 25771 60 s; defensive cooldowns), `spelltips/paladin.json` (Forbearance text on Divine Shield, Divine Protection, Templar's Bulwark, Blessing of Protection; Lay on Hands 550 at rank 3; Voice of Truth 6 s, 3 min), `talents/paladin.json`, `trainables/paladin.json`, `addon-builds.json` (six orders, decoded), `addon-data.json` (every paladin band).
- Engine bff826207: `sim/core/spell.go` (`SpellLevelPenalty`, grace 12, 5% a level), `sim/core/spell_max_level_auto_gen.go` (635: 5).
- Curated: `data/curated/apl/paladin-*.json`, `data/curated/rank-choices.json`, `data/curated/coach/paladin-*.{md,json}`, `data/pipeline/coachrules.py`, `sim/request/raid-consumes.json`.
- Coach: `api/internal/coach/f1_core_casts.go`, `builddata.go`, `f2_uptime.go`, `f5_ranks.go`, `f7_death.go`, `copy.go`; `tools/coach/*` (scratch regen at `scratchpad/gr26/paladin/regen/`).
- Reports: `design/reviews/2026-11-05-sim-audit-round22/paladin.md` (H1, H3, P2, R2); `round25/paladin.md`; `round26/druid.md`.
- Git: 438f8ba15..ca06416cd (928ddcf14, 5dbffb0bf, f2084a5bb, c05f93f85, 7b8c7acbb, 0cb781bca, 210c63d93).

## Appendix A: Note sweep

Bands from `addon-data.json`. Holy and Protection ship one order (leveling and raid identical); Retribution's raid order differs from 40 (Unsealed Wrath at 40, everything after shifts a level).

**Holy** (bands 10, 14, 18, 20, 21, 25, 30, 32, 38, 40, 43, 44, 46, 47, 50, 54, 55, 56, 60)
- Row 0 auto cooldowns, 10-60: potions and Demonic Runes; Divine Favor kept off the list. OK.
- Row 1 Oath of Light, 25-60: talent 25; "rank close to your level"; 5 minutes. OK.
- Row 2 Divine Favor, 21-60: talent 21; 40%; 2 minutes; sure crit. OK.
- Row 3 Holy Light below 40%, 10-60: "right behind Divine Favor" at 10 to 20 (Divine Favor 21). R26-3. Rank 8 at 60; "From 25 ... keep casting" (Holy Shock from 32 is on a cooldown). OK otherwise.
- Row 4 Holy Light paced, 10-60: 5-point pace, 45% at halfway OK; "keep all of your healing power" at 18 to 29. R26-1.
- Row 5 Holy Shock, 32-60: Boundless Light 30% from 34 (33, 34); Illumination 5/5 from 26. OK.
- Row 11 Flash of Light tank below 90%, 20-60: Oath from 25; rank 4 at bare 60 (row). OK.
- Row 12 Flash of Light others below 80%, 20-60: rank 1 most of leveling (2 at 44 to 49), rank 4 from 50, rank 5 at raid 60 (rows). OK.

**Protection** (bands 10, 12, 18, 20, 26, 28, 29, 30, 34, 36, 38, 40, 42, 43, 44, 50, 52, 58, 60)
- Row 0 auto cooldowns, 10-60: Lay on Hands below 10%, trains 10. OK.
- Row 1 Templar's Bulwark, 20-60: talent 20; 35%; 8 s; Forbearance; Divine Shield from 34. OK.
- Row 2 Divine Shield, 34-60: 25%; 10 s, 12 s from rank 2 at 50. OK.
- Row 3 Divine Protection, 10-60: 20%; Forbearance. OK.
- Row 4 Stoneform (Dwarf): 70%; 10% physical, 8 s, 3 minutes. OK.
- Row 5 Seal of Fury, 10-60: last 2 s; Holy Shield from 30; 30 s. OK.
- Row 6 Holy Shield, 30-60: 10 s category cooldown and duration; four charges. OK.
- Row 7 Holy Strike, 10-60: "twenty at 60" (rank 8); Iron Creed conditional (39); Shieldsworn from 40. OK.
- Row 8 Hammer of the Righteous, 40-60: reserve; target plus three. OK.
- Row 9 Swift Judgement, 29-60: one minute, free, off the global; hold for a swap. OK.
- Row 10 Judgement, 10-60: Holy Shield from 30; 6% of base mana; taunt. OK.
- Row 11 Exorcism, 20-60: Holy Shield from 30; Undead and Demons. OK.
- Row 12 Consecration, 20-60: above half mana; Holy Shield from 30; Hammer of the Righteous from 40. OK.

**Retribution** (bands 10, 12, 18, 20, 26, 28, 30, 34, 36, 38, 40, 42, 44, 50, 52, 53, 54, 58, 60; both orders)
- Row 1 Judgement, 10-60: Judgement 4; off the global; 6% of base mana; Benediction full at 14 (both orders); Sanctified Judgement 60% at 22. OK.
- Row 2 Seal of Command opener, 26-60: Unsealed Wrath conditional (leveling none, raid 40); Crusader 6; three stacks. OK (SHOULD "dungeon or raid").
- Row 3 Seal of Righteousness opener, 10-60: top rank trained; without Unsealed Wrath. OK.
- Row 4 Exorcism, 20-60: Undead and Demons. OK.
- Row 5 Hammer of Wrath, 44-60: 6 s all ranks; Champion of the Light 5/5 at 39 (both orders). OK.
- Row 6 Seal of Righteousness release, 40-60 raid order: four-fifths pace; 180 and 189 at 60. OK (SHOULD).
- Row 7 Seal of Command back, 40-60 raid order: 15 s reset. OK.
- Row 8 Seal of Command keep up, 40-60 raid order (needs rank 5 known): pace; last second. OK.
- Row 9 Seal of Righteousness otherwise, 10-60: matches rows 6 to 8. OK.
- Row 10 Holy Strike, 10-60: 18 at 60 with Benediction; Sacred Arbiter 27, conditional. OK.
- Row 11 Consecration, 20-60: reserve and pace; rank 4 at bare 60; 40% at halfway. OK (SHOULD placement).

## Appendix B: Coach sweep

**paladin-holy.md**
- :18 Who they are: Oath (25), Flash of Light (20), Divine Favor (21); "cheap Flash of Lights". SHOULD.
- :24 Oath of Light from 25. OK.
- :25 Concentration Aura (22), unscoped. R26-9.
- :26 Divine Favor from 21, then Holy Light. OK.
- :27 Holy Light paced; Holy Shock (32) unscoped. R26-9.
- :28 Flash of Light staple from 20; party below 80%; catch-up rule. OK (rank rule missing, R26-8).
- :29 orphan "Blessing of Wisdom" line. R26-8.
- :30 Holy Shock from 32. OK (SHOULD repeat).
- :31 Light's Vigil from 40. OK.
- :32 Voice of Truth (20), 6 s, 3 min, unscoped. R26-9.
- :33 Lay on Hands, 550 at rank 3, 20 min. OK (SHOULD).
- :34 Blessing of Protection (10). OK.
- :37 to :42 finding list: core and uptime lines unreachable (SHOULD); ranks line matches F5's overheal check.
- :44 consumables: Flask of Distilled Wisdom, Sage's Tea, Major Mana Potion match the preset; no oil (SHOULD).
- :48 core_casts OK. :49 uptime SHOULD. :50 resource OK. :51 ranks R26-8. :52 preparation OK (flask from 50: shared root). :53 death OK.
- Rule file: core and efficiency empty; downranks Holy Shock, Flash of Light OK; defensives R26-7.

**paladin-protection.md**
- :18 Who they are: Holy Shield (30), Holy Strike every 6 s (40). SHOULD.
- :24 Righteous Fury (10), 30 minutes. OK.
- :25 Seal of Fury at about -3 s (10). OK.
- :26 Holy Shield from 30. OK.
- :27 Holy Strike, twenty at 60; Shieldsworn and Iron Creed from 40. OK.
- :28 Hammer of the Righteous from 40, reserve. OK.
- :29 Swift Judgement from 29, hold for a swap. OK.
- :30 Judgement, pull only. R26-11.
- :31 Exorcism (20). R26-10.
- :32 Consecration (20). R26-10.
- :33 Greater Blessing of Kings (60), Blessing of Kings (20). R26-10.
- :34 Devotion Aura (1). OK.
- :35 Templar's Bulwark (20). R26-10.
- :36 Divine Shield (34), Divine Protection after Forbearance. R26-10.
- :37 Lay on Hands, full, 20 min, all your mana. OK.
- :40 to :47 finding list: core line R26-5; ranks line OK with F5's rank rows.
- :49 consumables: Flask of the Titans (50), Baked Salmon, Major Spellblasting Potion (55) match the preset; no oil (SHOULD).
- :53 core_casts OK. :54 uptime SHOULD. :55 idle OK. :56 resource OK. :57 ranks OK (rank 1 Consecration and Holy Strike rows). :58 preparation OK. :59 death R26-10 (Bulwark from 20). :60 adherence OK.
- Rule file: core Hammer of the Righteous and Judgement R26-5; Holy Shield aura OK; defensives R26-7.

**paladin-retribution.md**
- :18 Who they are: Judge every 8 s (16), Hammer of Wrath (44). SHOULD.
- :24 Blessing of Might (4). OK.
- :25 Seal of the Crusader (6), skip if another paladin's Crusader is up. OK.
- :26 Judgement on cooldown, off the global. OK.
- :27 Exorcism from 20. OK.
- :28 Hammer of Wrath from 44, highest rank (the card casts the newest rank at 44, 52, 60). OK.
- :29 Seal of Righteousness, no "while leveling". R26-12.
- :30 Seal of Command with Unsealed Wrath (raid 40). OK (pace missing, R26-12).
- :31 Holy Strike every 10 s. OK.
- :32 Consecration from 20, reserve. OK (SHOULD).
- :33 Holy Wrath from 50, Undead and Demons. OK.
- :36 to :43 finding list: core line R26-6; uptime line unreachable (SHOULD).
- :45 consumables: Smoked Desert Dumplings, Elemental Sharpening Stone (50), Major Mana Potion (49) match the preset; no flask (F6 shared root).
- :49 core_casts OK. :50 uptime SHOULD. :51 idle OK. :52 resource OK. :53 ranks OK (rank 1 Consecration rows). :54 preparation OK. :55 death OK. :56 adherence OK.
- Rule file: core Judgement R26-6; Holy Strike OK; auras empty; defensives R26-7.

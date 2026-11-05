# Guide accuracy round 26: Hunter

Audited on the round 26 audit tree: main after guide round 25 landed (head ca06416cd). It includes sim round 22 (merged in fd758eefc, engine pin bff826207 in `sim/enginever/version.go`), which no guide round had audited. Round 25's hunter audit ran on 438f8ba15, before that merge. Checked against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).

Files read:
- `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`
- the row notes in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`
- every rendered row of `addon-data.json` `rotations/hunter-*`, with `race_spell_ids`
- `sim/rankchoices/rank-choices.json` (hunter rows)
- the coach starters `data/curated/coach/hunter-*.md` and rule files `hunter-*.json`
- the hunter entries in `data/pipeline/coachrules.py`
- the coach checks that read those rule files: `api/internal/coach/f1_core_casts.go`, `f2_uptime.go`, `f5_ranks.go`, `f6_preparation.go`, `f7_death.go`, `builddata.go`
- the generator `tools/coach/persona_starter.py` and `starter_overrides.py`

Read first: the round 25 hunter report, the ledger's Round 25 section, `design/reviews/2026-11-05-sim-audit-round22/hunter.md` and integration.md, and `git diff 438f8ba15..HEAD` on every hunter file.

The generator was run read-only into the scratchpad (`persona_starter.py --spec hunter-<spec> --out gr26/hunter-coach/hunter-<spec>.md`). All three outputs are byte-identical to the committed starters. So every coach line below survives a regen unless its source (a card note, an override or the rule file) changes.

## New work: sim round 22 (hunter)

| change | where | verdict |
|---|---|---|
| B1 and M2 Arcane Shot and Serpent Sting rank rows | `rank-choices.json`; rendered rows BM 26 to 59, MM 20 to 59 | The rows match the prose's levels. They break the card notes' per-mana comparison at 40 and up: **HUNTER-R26-2** |
| BM:161 and MM:167 rank paragraphs | page prose | The levels match the rows. The reason given ("a low rank keeps most of the damage") is false: **HUNTER-R26-1** |
| M1 MM Aimed Shot rank 1 at 50 and 51 | MM:165; rows 28 to 59, bare 60 | OK. Ranks 28, 36, 44, 52, 60 (`trainables`). "most of what the newer ranks do" is under Sims requested (S3) |
| F2 Survival Raptor Strike note | SV card [2] | Two errors: **HUNTER-R26-3** (trap costs quoted without Resourcefulness) and **HUNTER-R26-4** (48 and 49 against the Horde rows and SV:148) |
| F3 Counterattack row from 32 | SV rows | OK. The row renders from band 32, which matches the point at 32 on both orders |
| N1 pet crit from a scope | engine only | No page line affected |

## Round 25's fixes

| item | verdict |
|---|---|
| R25-1 Aimed Shot / Multi-Shot override lines | OK (BM coach:28-29, MM coach:27,29) |
| R25-2 Survival potion | OK. Major Frenzy Potion matches the preset from 55 |
| R25-3 Counterattack "free" | OK, removed |
| R25-4, R25-11 orders | OK. SV is Mark, Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, Wing Clip. BM and MM put the mark first |
| R25-5 racial line | OK (SV coach:32) |
| R25-6 Feign Death line | OK as text. It names Feign Death without its level: R26-7 |
| R25-7, R25-8 drills | OK |
| R25-9 BM Rapid Fire "from forty with Bestial Wrath" | OK. Rapid Fire itself is still unscoped (trains at 26): R26-7 |
| R25-10 Explosive Trap reason | OK. Cost 520 against Volley 490 at 60 |
| R25-12 Growl | OK, gone from the BM list |
| R25-13 voice | OK |
| coachrules.py: Multi-Shot dropped from the BM and MM core lists | Correct in itself, but Aimed Shot is still judged as if its cooldown were its own: **HUNTER-R26-5** |

## Findings

Counts, guide pages and cards: CONFIRMED 4 (R26-1 to R26-4), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

Counts, coach: CONFIRMED 4 (R26-5 to R26-8), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R26-1 | CONFIRMED | page: the reason a sim change gives is false at most of the levels it covers (sim round 22, new work) | `beast-mastery.md:161`; `marksmanship.md:167` | "Both add a share of your ranged attack power on every rank, so a low rank keeps most of the damage for a fraction of the mana." | **The formula.** The engine's damage is base + 0.15 x RAP for Arcane Shot (`arcane_shot.go`, `arcaneShotAPCoeff`), and five ticks of base + 0.05 x RAP for Serpent Sting (`serpent_sting.go`). Base values come from the client: Arcane Shot 25 / 50 / 85 / 120 / 175 / 220 at ranks 1, 3, 4, 5, 6, 7; Serpent Sting 3 / 14 / 26 / 40 per tick at ranks 1, 3, 4, 5.<br>**What "most" needs.** For rank 1 to keep half of a newer rank, ranged attack power must reach:<br>- Arcane Shot against rank 5 (36 to 43): 467<br>- Arcane Shot against rank 6 (44 to 49): 833<br>- Arcane Shot rank 3 against rank 7 (BM 52 to 59): 800<br>- Sting rank 1 against rank 4 (26 to 33): 400<br>- Sting rank 1 against rank 5 (34 to 39): 680<br>A bare or quest-geared hunter at those levels is far below that: about level x 2 + Agility x 2 + Hawk, roughly 250 to 600.<br>**What it actually keeps.** Rank 1 keeps a third to a half of the damage. The rows win because rank 1 does two to three times the damage per point of mana. Sim round 22 B1 says it: "a leveling hunter who fires both on cooldown was running dry for over a minute". | Both pages: "Both add a share of your ranged attack power on every rank, so a low rank does two to three times the damage for every point of mana, and out in the world, where you run dry, that is worth more than the bigger hit." Hand edit; the coach is not affected. |
| HUNTER-R26-2 | CONFIRMED | card note: a comparison that sim round 22's rank rows broke (round 20 shape 3) | `hunter-beast-mastery.json` [6] and `hunter-marksmanship.json` [6], Serpent Sting note; renders 10 to 60 | "Every tick carries 5% of your ranged attack power on top of its own damage, so it does more for its mana than Arcane Shot, and from 18 more than Multi-Shot too" | **The rows.** From sim round 22, the Arcane Shot row casts rank 1 (25 mana) at BM 26 to 49 and MM 40 to 59. The sting row casts its newest rank from BM 40 and MM 30 (`addon-data.json`).<br>**Per mana.** Sting rank r (base T a tick, cost C) beats Arcane Shot rank 1 only if 25 x (5T + 0.25 RAP) > C x (25 + 0.15 RAP). That needs RAP below:<br>- rank 5 (40, 115): 193<br>- rank 6 (58, 150): 215<br>- rank 7 (75, 190): 208<br>- rank 8 (100, 230): 239<br>Level x 2 plus Agility and Hawk alone puts a level 40 hunter above 250. Before the rows, M3 measured BM 40 Arcane Shot at 2.22 per mana with rank 5, which implies about 600 RAP.<br>**Where it fails.** The claim is false at BM 40 to 49 and MM 40 to 59. It holds at 10 to 39 on both specs, and in a raid at 60, where both cast top ranks. At BM 50 to 59 (sting rank 7 or 8 against Arcane Shot rank 3) it holds only below about 580 to 610 RAP: Sims requested S2. | Notes (both specs): "...on top of its own damage, so from 18 it does more for its mana than Multi-Shot: keep it rolling for the whole fight." BM keeps its last sentence. Regenerate `addon-data.json` and Data.lua (note text only). |
| HUNTER-R26-3 | CONFIRMED | card note: a cost quoted without the build's own discount (round 21 shape) | `hunter-survival.json` [2] Raptor Strike note, later sentence; renders 10 to 60 | "...and rank 5 at 46 and 47, when Immolation Trap rank 4 arrives at 190 mana against 135." | 190 and 135 are the base costs of 14304 and 14303. Resourcefulness, taken at 33 and 34 on both Survival orders, reads "Reduces the Mana cost of your Trap and melee abilities by 60%" (`talents/hunter.json`). So at 46 and 47 the trap costs 76 against 54. The same note says Resourcefulness makes Raptor Strike "cost 60% less" from 34, and SV:53 says "60% off your melee and trap costs". The reason still holds (rank 5 saves Raptor Strike mana for a dearer trap), but the numbers are false. | "...and rank 5 at 46 and 47, when Immolation Trap rank 4 arrives at 76 mana against 54 with Resourcefulness." Or drop the numbers: "when the new Immolation Trap rank costs more". Regenerate addon-data (note text). |
| HUNTER-R26-4 | CONFIRMED | card note: false at bands and races it renders, and against the page (sim round 22 F2, new work) | `hunter-survival.json` [2] first sentence and 48 sentence; rendered rows 28, 48, 49 | First sentence (the only one the addon shows): "Raptor Strike on every swing it is ready, with an older rank on your bar at 26 to 29 and at 46 and 47." Later: "At 48 move to rank 7, unless your mana still runs out over a chain of pulls; then rank 5 stays the better button until 50." | **The rows.** At 48 and 49 the card renders rank 5 (14263) for Orc, Troll and Tauren, and at 48 for Windshaper Skyborne (`race_spell_ids`, from `rank-choices.json`).<br>**The page.** SV:148 says: "At 48 put rank 7 on the bar. On a Horde character rank 5 is still the better button: an Orc, Troll or Tauren at 48 and 49, a Windshaper Skyborne at 48."<br>**What the note says instead.** Its first sentence limits the older rank to 46 and 47. Its 48 sentence swaps the race rule for a mana rule. That sends an Orc with mana to spare to rank 7 beside a rank 5 icon, and an Alliance hunter who runs short to rank 5, which no row casts.<br>**A smaller miss.** The note's "rank 2 for most Horde hunters at 26 and 27" leaves out the Tauren row at 28 (rank 2), which SV:144 names. | First sentence: "Raptor Strike on every swing it is ready, with an older rank on your bar at 26 to 29 and at 46 and 47, and to 49 on most Horde characters." 48 sentence: "At 48 move to rank 7; an Orc, Troll or Tauren keeps rank 5 through 49, a Windshaper Skyborne at 48." Parenthesis: "(rank 2 for most Horde hunters at 26 and 27, and a Tauren at 28)". The card line must stay under its length limit. Regenerate addon-data and the coach starter (R26-8). |
| HUNTER-R26-5 | CONFIRMED | coach rule file and F1: a core cooldown the coach cannot read, because another spell shares it | `coachrules.py` `hunter-beast-mastery` and `hunter-marksmanship` `core=[..., "Aimed Shot"]`, read by `f1_core_casts.go` (`castsOf`, `coreGaps`) | F1 title "Aimed Shot: {n} casts lost"; detail "{lost} of a possible {possible} went unused while you had the mana." | **The shared cooldown.** Multi-Shot and Aimed Shot share category 6000. The client tooltip of 2643 says "Multi-Shot shares its cooldown with Aimed Shot".<br>**What F1 reads.** `castsOf` reads only the rule's Aimed Shot ids. After each cast, `coreGaps` marks Aimed Shot ready one cooldown later until the next Aimed Shot. No code in `api/internal/coach` knows about shared categories (grep "categor").<br>**What a Multi-Shot costs.** Aimed Shot succeeds at t; Multi-Shot goes at t+6, which the pages sanction (BM:91, MM:96: "Multi-Shot is for packs and for when you have to move"); the next Aimed Shot starts at t+12 and lands at t+14. The gap runs t+6 to t+14, which is 8 s. Less the 1.5 s grace, that leaves 6.5 s, so one lost cast. Every Multi-Shot the page asks for on a pack, or on the move, is reported as a lost Aimed Shot.<br>**Why it matters.** Round 25 removed Multi-Shot from core "so a ready Multi-Shot is never a lost cast". The mirror case was missed: a cast Multi-Shot is a "lost" Aimed Shot. | F1 code: a rule entry's cooldown also restarts on the casts of the spells that share its category. Read `category_cooldown_ms` partners from spellconst, or add a `shares` id list to the rule entry (2643 for Aimed Shot) and include those ids in `coreGaps`' cast sequence. Add a test that fails first: Aimed Shot, Multi-Shot, Aimed Shot at 0 / 6 / 14 s loses 0. Until then, drop Aimed Shot from the BM and MM core lists. Rule-file and code change. |
| HUNTER-R26-6 | CONFIRMED | coach F1: cooldowns that the shipped builds shorten are judged at base | `coachrules.py` `hunter-survival` core "Immolation Trap"; BM and MM core "Rapid Fire"; `f1_core_casts.go` (`sc.CooldownMS`), `copy.go` BasisCooldown | "Measured against Immolation Trap's 30-second cooldown"; "{lost} of a possible {possible}" | **Immolation Trap.** Survivalist's Discipline ("Reduces the cooldown of your Trap and Deterrence abilities by 40%") is points 37 and 38 on both Survival orders. From 38 every Survival hunter's Immolation Trap is on an 18 s cooldown. F1 states 30 s and counts "possible" casts on 30 s.<br>**Rapid Fire.** Rapid Killing takes 2 minutes off at 25 and 26 on the Marksmanship leveling order (3 minutes, which the MM card note and MM:161 teach). On the Beast Mastery leveling order it takes 1 minute off at 57 (4 minutes).<br>**Effect.** Whenever F1 fires on these spells, the cooldown it names and the "possible" count are false: too lenient, but stated as fact. The raid orders carry no Rapid Killing, so Rapid Fire at 300 s is right for a raider. | Survival: give F1 the build's cooldown (18 s from 38; both orders agree, so a level-gated value in the rule file is enough), or drop Immolation Trap from core and leave it to adherence. Rapid Fire: if the log can't say which build, word the basis as "at most every five minutes", or skip Rapid Fire below 60. Rule-file and code change. |
| HUNTER-R26-7 | CONFIRMED | coach: a spell or talent named without the level it arrives (the generator's own rule; priest R26-3 shape) | `coach/hunter-beast-mastery.md:27, 32, 33, 34, 35`; `coach/hunter-marksmanship.md:26, 28, 32, 33, 34, 35`; `coach/hunter-survival.md:26, 28, 29, 32, 33, 34, 35, 36, 37` | **Rotation lines:**<br>- Rapid Fire "on the pull" (all three; trains at 26)<br>- Sniper Shot "whenever it is ready" (MM; the raid order's 40th point, never on the leveling order)<br>- Lacerate "whenever its bleed has fallen off" and Strider Kick "on its own short cooldown" (SV; both 30)<br>- Counterattack (SV; point 32)<br>**Threat and utility lines:**<br>- Feign Death (30)<br>- Deterrence (SV; point 25)<br>- Disengage (20)<br>- Freezing Trap "laid before the pull" (20)<br>- Explosive Trap "is your AoE" (34)<br>- Volley (40) | `starter_overrides.py` states the rule in GUIDE_LINES' header: "A line that names a spell, talent or rank the player gains during leveling says from when (round 25: the coach reads a boss pull at whatever level the log records, dungeon bosses from the low teens included)". The same files already do this for Aimed Shot ("from twenty"), Bestial Wrath ("from forty"), Multi-Shot ("at eighteen") and Expose Prey ("From thirty-five").<br>**Levels.** From `trainables/hunter.json` and `addon-builds.json`.<br>**Example.** A level 22 to 25 hunter on a dungeon boss is told to press Rapid Fire on the pull, and a Survival hunter at 20 to 29 to keep Lacerate and Strider Kick going.<br>**Precedent.** Priest R26-3 (Lightwell at 40) and round 25's Shadowform line. Tranquilizing Shot (60, "a raid job") and Distracting Shot (12) need no level. | `starter_overrides.py` GUIDE_LINES, for example:<br>- "Rapid Fire from twenty-six on the pull..."<br>- "With the raid build, from forty, Sniper Shot..."<br>- "Lacerate from thirty whenever its bleed has fallen off."<br>- "Strider Kick from thirty on its own short cooldown, ahead of Immolation Trap."<br>- "Counterattack from thirty-two the moment it lights up..."<br>- "Feign Death, from thirty, before you reach the tank's threat..."<br>- "Deterrence, from twenty-five, is your emergency button..."<br>- "Explosive Trap, from thirty-four, ..."<br>- "Volley, from forty, ..."<br>- "Freezing Trap, from twenty, ..."<br>- "Disengage, from twenty, ..."<br>Override, then regenerate. |
| HUNTER-R26-8 | CONFIRMED | coach: the Raptor Strike line inherits R26-4 | `coach/hunter-survival.md:25` | "Raptor Strike on every swing it is ready, with an older rank on your bar at twenty-six to twenty-nine and at forty-six and forty-seven." | The line is the card note's first sentence (R26-4). For an Orc, Troll or Tauren at 48 and 49 (Windshaper at 48), the rows and SV:148 keep rank 5. The coach's line says the top rank. | Fixed by R26-4's first-sentence edit and a regen (`persona_starter.py --spec hunter-survival`). No override needed. |

## Shared root, filed by another class this round (not counted here)

- DRUID-R26-6 (F6 judges the flask by role, not by the spec's preset) reaches every hunter. No hunter preset carries a flask at any band (`sim/request/raid-consumes.json`), and the hunter personas name none, but the dps catalogue marks flasks complete. Every hunter at raid size gets "no flask" as a major finding. The druid fix covers it; nothing hunter-specific is needed.

## Shoulds (wording, not findings)

- **S1. Feign Death as the death-review defensive** (`hunter-*.json` defensives, F7). It drops threat. It does nothing against raid damage. "Died at {t} with Feign Death ready" after a cleave or a fire death reads as a missed save. It is true (Feign Death was ready), and it does save a hunter whom an add or the boss turned on. Consider limiting F7's Feign Death line to deaths from the target's melee, or naming it for threat only.
- **S2. Deterrence in the BM and MM rule files.** Neither build takes it (`addon-builds.json`). It is talent-gated by cast, so it never fires, but the entry is dead.
- **S3. MM rule file has no Serpent Sting aura**, though MM:96 and the card keep it rolling. This is an omission, not a false statement.
- **S4. BM:99 and MM:105** "the sting is the best damage for its mana you have and Arcane Shot is next". This is the raid rotation paragraph, where top ranks are cast; Sims requested S1 checks it at 60. Read by a leveling player, it is false at BM 40 to 49 and MM 40 to 59, where rank 1 Arcane Shot beats the sting per mana (R26-2), and at MM bare 60, where rank 1 Aimed Shot measured 7.68 against the sting's 5.37 (sim round 22 M3). "In a raid," at its start would scope it.
- **S5. BM:159 "well over half" and MM:165 "most"** for rank 1 Aimed Shot. The two pages disagree (sim round 22's handed SHOULD 6). The engine's shot is normalized weapon damage + ammo + flat 40 / 70 / 115 / 185 / 260 / 350. At 52 and bare 60 the share depends on the band's bow (Sims requested S3).
- **S6. BM coach:29 and MM coach:29.** The quarter-mana clause attaches only to "From twenty", but the card holds Multi-Shot under a quarter at 18 and 19 too.
- **Carried from round 25:** S1 (index:33 Survival scope), S2 (MM coach:18 and MM:48 "Aimed Shot and Multi-Shot on their shared cooldown"), S3 (SV:94 "either"), S6 (SV coach Mongoose Bite wording, now "after a dodge. From thirty-five ..." which is fine), S8 (BM coach:30 sting line drops the card's hold), and round 23 and 24's carried Shoulds.

## Superlative, comparison and quoted value sweep (new and changed lines)

| line | claim | verdict |
|---|---|---|
| BM:161, MM:167 | "a low rank keeps most of the damage" | **R26-1** |
| BM:161 | ranks at 26, 40, 50 | OK against `rank-choices.json` and rows |
| MM:165 | Aimed ranks at 28, 36, 44, 52, 60 | OK (`trainables`) |
| MM:165 | "the expensive ranks drain you fastest of all" | OK: Aimed Shot takes the shared cooldown every 6 s, and nothing else the card fires that often costs more |
| MM:167 | 20 to 29 rank 1, 30 newest sting and Arcane rank 3, 40 Arcane rank 1 | OK against rows |
| BM and MM sting notes | more per mana than Arcane Shot | **R26-2** (40 and up) |
| BM and MM Arcane notes | far more per mana than Multi-Shot | OK at every band (rank 1 or 3 only widens it) |
| BM and MM Aimed notes | hits harder than / outhits Multi-Shot | OK (flat 40 or more over the same normalized weapon hit) |
| SV Raptor note | 190 against 135 | **R26-3** |
| SV Raptor note | 26 to 29, 46 and 47, 48 rule | **R26-4** |
| SV Raptor note | "does not use the global cooldown" | OK (`gcd_ms` 0, all ranks) |
| SV Raptor note | Resourcefulness at 33 and 34, 60% | OK (both orders) |
| coach, all three | every cooldown, cost and duration | unchanged from round 25's sweep. OK, apart from R26-6's F1 basis values |

## Decisions

- **D1.** R26-1 and R26-2 are CONFIRMED on arithmetic, not left for a sim. They use the engine's own formula (client base values, 0.15 and 0.05 attack power shares) and thresholds no level 26 to 59 hunter's ranged attack power reaches (R26-1), or that every one exceeds (R26-2, at 40 and up). The one band the arithmetic cannot settle, BM 50 to 59 for R26-2, is left to Sims requested S2 and kept out of the finding.
- **D2.** R26-4 follows the rows and SV:148 (by race), not the note's mana rule. Sim round 22 F2 says the split is the faction's band gear. A player can check neither "faction band gear" nor the sim's race result, but they know their race. The rendered row is per race, so the text must match it.
- **D3.** R26-5 is the hunter instance of the round's "core ability the coach cannot read" lens. Here the cost can be read; the cooldown can't, because Multi-Shot casts restart it. It is the worst coach finding: it accuses a player of losing casts for doing what both pages tell them.
- **D4.** R26-6 is CONFIRMED although it errs lenient. The finding text states a cooldown and a "possible" count that are false for every Survival hunter from 38.
- **D5.** R26-7 counts as one finding across the three specs because it is one rule (the generator's) broken by one omission. The priest report filed the same shape as one item.
- **D6.** R26-8 is counted in the coach column although R26-4's fix regenerates it, because the coach file is its own player-facing text (round 25 counted the coach separately).
- **D7.** Not re-filed: anything routed in the ledger's Round 19 to 25 sections; DRUID-R26-6 (shared root); the Leveling strip and cooldown cards (component lane).
- **D8.** No web research was needed. The client data, the engine source, the cards, the rank rows and the coach code decide every call.

## Sims requested

- **S1.** Beast Mastery and Marksmanship, raid 60, guide raid builds, frozen BiS, 10000 iterations: damage per mana of Serpent Sting rank 9, Arcane Shot rank 8, Aimed Shot rank 6, and Sniper Shot rank 3 (MM). Decides BM:99 and MM:105 ("the sting is the best damage for its mana you have and Arcane Shot is next"; Should S4).
- **S2.** Beast Mastery bare 50 to 59 on the shipped rows: damage per mana of Serpent Sting rank 7 and 8 against Arcane Shot rank 3. Decides whether R26-2's fix must also cover BM 50 to 59. The proposed fix drops the comparison, so it is safe either way.
- **S3.** Beast Mastery and Marksmanship bare 28, 36, 44, 52 and 60: average hit of Aimed Shot rank 1 against the newest rank. Decides BM:159 "well over half" against MM:165 "most" (Should S5).

## Sim findings

None new.

## Engine questions

None new. Checked: the engine reads Aimed Shot rank 6's flat 350 from 20904, not the 600 duplicate 27632 (`client_damage.go` overrides the generated row). Standing: round 25's list, and sim round 22's open items (Immolation Trap 2% a tick; the attack power share reaching a low rank in full).

## Sources

**Client data, 1.60.1.70790:**
- `spellconst/hunter.json`:
  - Arcane Shot 3044, 14281 to 14287
  - Serpent Sting 1978, 13549 to 13555, 25295
  - Aimed Shot 19434, 20900 to 20904, 27632
  - Multi-Shot 2643 (category 6000, cost 13.9% of base mana)
  - Immolation Trap 13795, 14302 to 14305
  - Raptor Strike 2973, 14260 to 14266 (gcd 0)
  - Strider Kick 1317257 (cost 5.81% of base mana)
  - Rapid Fire 3045; Bestial Wrath 19574
  - Lacerate 24118 to 1299332; Feign Death 5384; Deterrence 19263
- `spelltips/hunter.json` (2643 "shares its cooldown with Aimed Shot")
- `talents/hunter.json` (Resourcefulness, Survivalist's Discipline, Rapid Killing, Improved Stings, Deterrence)
- `trainables/hunter.json`, `addon-builds.json` (decoded orders), `addon-data.json` rotations with `race_spell_ids`

**Engine bff826207:** `sim/hunter/arcane_shot.go`, `serpent_sting.go`, `attack_power_scaling.go` (0.15, 0.05), `aimed_shot.go`, `client_damage.go`.

**Repository:**
- `sim/rankchoices/rank-choices.json`
- `sim/request/raid-consumes.json`
- `data/pipeline/coachrules.py`
- `api/internal/coach/{f1_core_casts,f2_uptime,f5_ranks,f6_preparation,f7_death,builddata,copy,rules}.go`
- `tools/coach/{persona_starter,starter_overrides}.py`
- `design/reviews/2026-11-05-sim-audit-round22/hunter.md`
- round 25 hunter, druid R26 and priest R26 reports
- ledger Round 25
- read-only regen in `scratchpad/gr26/hunter-coach/` (identical to the committed files)

**Web:** none consulted (D8).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent, spell, rank or value with the level the player has it (leveling and raid orders), verdict.

- BM Hunter's Mark (10 to 60): ranks r1 10 to 20, r2 22 to 38, r3 40 to 52, r4 58 to 60; one mark per target. OK.
- BM Bestial Wrath (40 to 60): talent 40 on both orders; 50%, 18 s, fear, two minutes; Rapid Fire 26. OK.
- BM Rapid Fire (26 to 60): spell 26; "from 40 with Bestial Wrath"; shared cooldown is Multi-Shot's from 18; five minutes, four from 57 on the leveling order (Rapid Killing 1), five on the raid order; damage racial guarded. OK.
- BM Aimed Shot (20 to 60): spell 20; rank 1 rows 20 to 58, rank 6 at 60; autos through the cast; outhits Multi-Shot at every rank. OK.
- BM Multi-Shot (18 to 60): spell 18; quarter-mana gate; "From 20 Aimed Shot"; worst per mana at every band (Arcane Shot and sting rank 1 rows only widen it). OK.
- BM Serpent Sting (10 to 60): rows r2 10 to 12, r3 18 to 22, r1 26 to 38, r5 40, r6 42, r7 50, r8 58 to 60. Arcane Shot rows r1 10, r2 12 to 18, r3 20 to 22, r1 26 to 49, r3 50 to 58, r8 60.
  - "more for its mana than Arcane Shot": OK at 10 to 39 and 60. **HUNTER-R26-2** at 40 to 49; 50 to 59 open (S2).
  - "from 18 more than Multi-Shot": OK.
  - Last sentence (Multi-Shot from 18, Aimed Shot from 20, quarter mana): OK.
- BM Arcane Shot (10 to 60): 15% RAP; "from 18 far more per mana than Multi-Shot": OK at every band.
- MM Hunter's Mark (10 to 60): as BM. OK.
- MM Rapid Fire (26 to 60): five minutes on the raid order; three from 26 on the leveling order (Rapid Killing 25, 26); damage racial guarded. OK.
- MM Aimed Shot (20 to 60): rank 1 rows 20 to 58, rank 6 at 60; outhits Multi-Shot. OK.
- MM Multi-Shot (18 to 60): as BM. OK.
- MM Sniper Shot (40 to 60): raid talent 40; ranks r1 40 to 42, r2 48 to 50, r3 58 to 60; 4 s cast, 15 s cooldown; hardest hit. OK.
- MM Serpent Sting (10 to 60): rows r2 10 to 12, r3 18, r1 20 to 26, r4 30, r5 34 to 40, r6 42 to 48, r7 50, r8 58 to 60. Arcane Shot rows r1 10, r2 12 to 18, r1 20 to 26, r3 30 to 38, r1 40 to 58, r8 60.
  - "more for its mana than Arcane Shot": OK at 10 to 39 (rank 4 always wins; rank 5 below 895 RAP) and 60. **HUNTER-R26-2** at 40 to 59.
- MM Arcane Shot (10 to 60): as BM. OK.
- SV Hunter's Mark (10 to 60): Expose Prey from 35 (35, 36, both orders). OK.
- SV Raptor Strike (10 to 60): rows r2 10 to 12, r3 16 to 22 and 29, r4 24 and 30, per race at 26 to 28, r5 32 to 38 and 46, r6 40 to 44, r7 48 to 54 (r5 for Orc, Troll, Tauren at 48 and 49, Windshaper at 48), r8 56 to 60.
  - "older rank at 26 to 29 and 46 and 47": **HUNTER-R26-4** at 48 and 49 (Horde rows).
  - "rank 2 for most Horde hunters at 26 and 27": misses Tauren at 28 (R26-4).
  - "190 mana against 135": **HUNTER-R26-3**.
  - Resourcefulness 33 and 34, 60%: OK. gcd 0: OK.
- SV Counterattack (32 to 60): talent 32 on both orders; ranks r2 32 to 40, r3 42 to 52, r4 54 to 60; cannot be dodged, blocked or parried. OK.
- SV Mongoose Bite (16 to 60): spell 16; 5 s dodge window; Expose Prey at 35 and 36, 10 s, any hunter's mark. OK.
- SV Lacerate (30 to 60): spell 30; names Mongoose Bite (16). OK.
- SV Strider Kick (30 to 60): talent 30; 8 s; "the bleed" is Lacerate (30); 5.81% of base mana. OK ("sliver" carried Should).
- SV Immolation Trap (16 to 60): ranks r1 16 to 24, r2 26 to 32, r3 36 to 44, r4 46 to 54, r5 56 to 60; Clever Traps 20 and 21, Survivalist's Discipline 37 and 38 in a "once ... are in" clause; attack power share. OK.
- SV Wing Clip (12 to 60): ranks r1 12 to 36, r2 38 to 58, r3 60; Windfury Totem conditional. OK ("cheap" carried Should).

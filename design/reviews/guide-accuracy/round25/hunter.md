# Guide accuracy round 25: Hunter

Audited on the round 25 audit tree (main after guide round 24, head 438f8ba15; nothing newer than round 24 for hunter, no sim round 22 in the log; engine pin 170a59c39 in `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, the row notes in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`, every rendered row of `addon-data.json` `rotations/hunter-*` (BM 18 bands, MM 19, SV 25 with `race_spell_ids`), and, new this round, the coach starters `data/curated/coach/hunter-{beast-mastery,marksmanship,survival}.md` with their generator `tools/coach/persona_starter.py` (and `starter_overrides.py`, `starter_kit.py`, `starter_priority.py`, `starter_text.py`, `starter_food.py`). Read first: round 24 hunter report and the ledger's Round 24 section; `git diff ea9adea1c..HEAD` on the hunter pages, cards and coach files.

To say for each coach finding whether regenerating fixes it, the generator was run read-only into the scratchpad (`persona_starter.py --spec hunter-<spec> --out <scratch>`) and diffed against the committed files. Regenerating changes exactly two lines (BM Rapid Fire, BM and MM Explosive Trap); every other coach line below survives a regen and needs an override or generator change. The generator's own test is red for four tests; for hunter, `test_survival_never_lists_the_ranged_aspect` fails because the Survival list carries the Falcon line (its content is true at 60, see Shoulds).

## Round 24's fixes

| item | where | verdict |
|---|---|---|
| R24-1 class page weave | `index.md:33` | OK. Aimed Shot 19434 cast 2000 ms, category 6000 shared with 2643 (both tooltips say so); Aimed Shot adds 40 to 350 over weapon damage, Multi-Shot adds 0, so it outhits at every rank; Multi-Shot is instant (cast 0), so "for when you have to move" holds. Matches BM:91 and MM:96 and the ungated Aimed Shot row. The line does not say it is a ranged-spec rule (Survival fires no shots): Should S1 |
| R24-2 MM Rapid Fire note | MM card, bands 26 to 60 | OK. Raid order has no Rapid Killing: five minutes (3045 category 300000). Leveling order Rapid Killing 25 and 26 ("by 2 min"): three minutes from 26, the level Rapid Fire trains. "damage racial, if your race has one" OK |
| R24-3 BM Rapid Fire note | BM card, bands 26 to 60 | OK. Bestial Wrath named "from 40" (talent 40 on both orders, row from 40); five minutes, four from 57 on the leveling order (Rapid Killing rank 1 "by 1 min"), none on the raid order; a four minute cooldown still gives one in a normal fight |
| R24-4 BM sting note last sentence | BM card, bands 10 to 60 | OK. At 18 and 19 the shared cooldown is Multi-Shot's (Aimed Shot 20); the "mana above a quarter" clause matches Multi-Shot's `currentManaPercent > 25%` gate and the sting's `<= 25%` bypass |
| R24-5 Barrage reason dropped | MM Aimed note; `marksmanship.md:96` | OK. "it outhits Multi-Shot" with no talent; "by a wide margin" sits in the raid paragraph (rank 6, +350) |
| R24-6 Survival Raptor Strike ranks by race | `survival.md:144` | OK against the rows: rank 2 for Orc at 26, Troll and Windshaper at 26 and 27, Tauren at 26 to 28; Alliance rank 3 at 26 to 28; base rank 3 at 29; rank 4 at 30 |
| R24 S4 carried into the fix | `survival.md:148` | OK: rank 5 at 46 (base row), rank 7 from 48, Horde rank 5 for Orc, Troll, Tauren at 48 and 49 and Windshaper at 48 |
| Damage racials | BM:89, MM:94, SV:88 | OK on the pages: each names the three racials in a parenthetical. The Survival coach drops it: **HUNTER-R25-5** |

## Findings

All thirteen are in the coach starters. No guide page or card note is false this round.

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R25-1 | CONFIRMED | coach: rotation reversed (stale override) | `coach/hunter-beast-mastery.md:24,27`; `coach/hunter-marksmanship.md:24,29` | "Aimed Shot only as the opener, cast before the pull so it lands as the fight opens." and "Multi-Shot every time the cooldown it shares with Aimed Shot comes up." | Both cards cast Aimed Shot whenever the shared cooldown is up (row 20904, no condition, above Multi-Shot), and Multi-Shot only above a quarter of mana; BM:91 and MM:96 "cast Aimed Shot every time the shared six second cooldown comes up ... Multi-Shot is for packs and for when you have to move". The opener line also contradicts BM:87 and MM:92: the opening Aimed Shot is timed "so it lands just after the tank's first hit. Started any earlier, it is the first threat on the boss". From 20, at every band, the coach teaches the opposite rotation. Round 24 S1 and the ledger already named it. | Regenerating does NOT fix it: the lines are hand overrides. `tools/coach/starter_overrides.py` GUIDE_LINES `hunter-beast-mastery` and `hunter-marksmanship`: "Aimed Shot": "Aimed Shot every time the cooldown it shares with Multi-Shot comes up. Your auto shots keep firing while you cast it.", "Multi-Shot": "Multi-Shot for packs and when you have to move, and never under a quarter of your mana." Update the stale guide quotes in the comments above them, and GUIDE_ORDER (R25-11). |
| HUNTER-R25-2 | CONFIRMED | coach: contradicts the page | `coach/hunter-survival.md:49` | "a combat potion (Major Mana Potion)" | SV:112: "Survival's mana holds with Resourcefulness, so it drinks a Major Frenzy Potion on the pull instead." The Survival raid preset (`sim/request/raid-consumes.json`, top band) carries `major_frenzy_potion`, not a mana potion. The coach would hold a Frenzy Potion against the player as a missing mana potion. | Generator: `tools/coach/starter_kit.py` `kit()` picks the potion from `Traits.resource` (`mana` gives Major Mana Potion). Read the combat potion from the spec's raid preset the way `starter_food.preset_food` reads the food (Major Frenzy Potion for Survival; Major Mana Potion stays for BM and MM, whose presets carry it). Regenerating alone does not fix it. |
| HUNTER-R25-3 | CONFIRMED | coach: false client value | `coach/hunter-survival.md:27` | "Counterattack only after you parry: it is free and cannot be blocked, dodged or parried." | Counterattack costs mana at every rank: 30 (19306, rank 1), 45 (1242634), 65 (20909), 85 (20910, the card's rank 4). Resourcefulness takes 60% off melee costs, not all of it. The card note says nothing about cost. | Override `starter_overrides.py` GUIDE_LINES `hunter-survival` "Counterattack": "Counterattack the moment it lights up after you parry: it cannot be blocked, dodged or parried." Regenerating does not fix it. |
| HUNTER-R25-4 | CONFIRMED | coach: order against the page and card | `coach/hunter-survival.md:26 to 32` | List order Raptor Strike, Counterattack, Strider Kick, Immolation Trap, Mongoose Bite, Lacerate, under "in the order it gives them". | SV:90: "Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap"; the card's rows run Mark, Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, Wing Clip. The coach puts Mongoose Bite, the spec's core proc button, and Lacerate below the trap, so its "core ability under-used" check would read a Mongoose Bite pressed ahead of a ready trap as a mistake. | Add `"hunter-survival": ("Hunter's Mark", "Raptor Strike", "Counterattack", "Mongoose Bite", "Lacerate", "Strider Kick", "Immolation Trap", "Wing Clip")` to GUIDE_ORDER in `starter_overrides.py`. Regenerating does not fix it. |
| HUNTER-R25-5 | CONFIRMED | coach: "your racial" for races with none | `coach/hunter-survival.md:25` | "On the pull, press Rapid Fire, your trinkets and your racial together." | Round 24's shape (ledger: "your racial" lines for races with no damage racial). A Human (Ruthless, passive), Dwarf (Stoneform) or Tauren (War Stomp) Survival hunter has nothing to press there. The page line SV:88 carries "(Elune's Light, Blood Fury or Berserking)"; `starter_text.PARENTHETICAL` strips it. | Override GUIDE_LINES `hunter-survival` "Rapid Fire": "Rapid Fire on the pull with your trinkets, and with Elune's Light, Blood Fury or Berserking if your race has one." Regenerating does not fix it. |
| HUNTER-R25-6 | CONFIRMED | coach: garbled line, misleading | `coach/hunter-beast-mastery.md:35`; `coach/hunter-marksmanship.md:34` | "Feign Death: Started any earlier, it is the first threat on the boss and the boss comes for you; on a fresh tank, give him a second or two more, and keep Feign Death ready if the boss turns anyway." | The sentence is BM:87 and MM:92 about the opening Aimed Shot ("Time Aimed Shot so it lands just after the tank's first hit. Started any earlier, it ..."). Under a "Feign Death:" label "it" reads as Feign Death, telling the player that Feign Death is the first threat on the boss. The Feign Death rule is BM:103 and MM:109: Feign Death before you reach the tank's threat; if it says Resisted you still have all your threat. | Override GUIDE_LINES "Feign Death" for both specs: "Feign Death before you reach the tank's threat, not after the boss turns. If it says Resisted, you still have all your threat." Generator: `starter_priority.priority_entries` should not label a sentence with an ability whose name appears only in a later clause and whose subject is a pronoun for another spell. Regenerating does not fix it. |
| HUNTER-R25-7 | CONFIRMED | coach: drill contradicts the rank advice | `coach/hunter-*.md`, drill.ranks (BM:60, MM:59, SV:57) and the "Ranks" mistake line | "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank." | Every hunter page tells the leveling player to keep a lower rank on the bar: BM:159 Aimed Shot rank 1 "all the way to 60"; MM:165 rank 1 through 49 and from 52 to 59; SV:144 and SV:148 Raptor Strike rank 3 (or rank 2 by race) at 26 to 29 and rank 5 at 46 to 49. The card rows show the same ranks (BM and MM rank 1 at 20 to 58; SV rows above). The drill, framed per level, tells the player to undo it, and the "Ranks" finding would mark it as a skipped trainer visit. | Generator: `tools/coach/starter_kit.py` DRILLS dps "ranks": "Visit your trainer every time you level, then check every button on the priority list shows the rank the guide gives for your level: sometimes that is a lower one." Regenerating does not fix it. |
| HUNTER-R25-8 | CONFIRMED | coach: drill contradicts the card | `coach/hunter-*.md`, drill.uptime (BM:57, MM:56, SV:54) | "Recast each with 3 seconds left" | BM and MM sting notes: "Serpent Sting when it has under 1.5 seconds left ... Refreshing any earlier throws away a tick." Serpent Sting ticks every 3 seconds over 15 (13555: 500 over 15 s, 100 a tick), so a recast at 3 seconds left clips the last tick. Survival's Lacerate row casts only when the bleed is gone (`not dotIsActive`); Lacerate ticks every 3 s over 21 (24118: 119 over 21 s, 17 a tick). | Generator: `starter_kit.py` DRILLS dps "uptime": "... Recast each as the card says, no earlier, and again the moment you are back up after a death or a phase change." (or a per-spec drill). Regenerating does not fix it. |
| HUNTER-R25-9 | CONFIRMED | coach: talent named before the order gives it | `coach/hunter-beast-mastery.md:26` | "Rapid Fire on the pull with Bestial Wrath, as an auto shot leaves and the shared cooldown is about to come back." | R24-3's shape: Rapid Fire trains at 26, Bestial Wrath is the level 40 talent on both orders. The card note now reads "and from 40 with Bestial Wrath". | Regenerating fixes it: the scratch regen gives "Rapid Fire on the pull as an auto shot leaves and the shared cooldown is about to come back, and from forty with Bestial Wrath." |
| HUNTER-R25-10 | CONFIRMED | coach: reason the page dropped | `coach/hunter-beast-mastery.md:38`; `coach/hunter-marksmanship.md:37` | "Explosive Trap under a pack the tank is holding is good damage for its mana." | Round 23 replaced this reason on the pages: BM:105 and MM:111 "at 60 it costs even more than Volley, so save it for the same big packs" (14317 costs 520, Volley 14295 costs 490). The coach still sells it as mana-efficient. | Regenerating fixes it (scratch regen: "... hits everything within ten yards, and at sixty it costs even more than Volley, so save it for the same big packs."). |
| HUNTER-R25-11 | CONFIRMED | coach: order against the page and card | `coach/hunter-beast-mastery.md:23 to 30`; `coach/hunter-marksmanship.md:23 to 25` | BM lists Hunter's Mark eighth, after Arcane Shot; MM third, after the Aimed Shot opener; both under "in the order it gives them". | BM:87 and MM:92: "mark the boss with Hunter's Mark first, so the mana is back by the pull ... then put up Aspect of the Hawk"; the prepull casts the mark at -20 s ahead of the aspect and the mark is row 1 on both cards. GUIDE_ORDER pins Aspect of the Hawk and Aimed Shot first. | `starter_overrides.py` GUIDE_ORDER `hunter-beast-mastery`: ("Hunter's Mark", "Aspect of the Hawk", "Bestial Wrath", "Rapid Fire", "Aimed Shot", "Multi-Shot", "Serpent Sting", "Arcane Shot"); `hunter-marksmanship`: ("Hunter's Mark", "Aspect of the Hawk", "Rapid Fire", "Aimed Shot", "Sniper Shot", "Multi-Shot", "Serpent Sting", "Arcane Shot"). Regenerating does not fix it. |
| HUNTER-R25-12 | CONFIRMED | coach: a button the page turns off | `coach/hunter-beast-mastery.md:33` | "Growl." listed as a priority button. | BM:101: "turn off Growl in raids (right-click it on the pet bar) or it will pull the boss off your tank." Listing it bare among "the buttons the guide's rotation names" tells the coach its absence is an under-used core button. Marksmanship's Growl line (MM coach:33) carries the "turn off" clause and is fine. "Mend Pet." (BM coach:34) is a between-pulls button (BM:101), not a rotation one: Should S4. | Override GUIDE_LINES `hunter-beast-mastery` "Growl": "Growl turned off on the pet bar in a raid, or the cat pulls the boss off your tank." (or exclude it in `starter_guide.excluded_abilities`). Regenerating does not fix it. |
| HUNTER-R25-13 | CONFIRMED | coach: site meta talk (voice rule) | `coach/hunter-*.md:17,21,52` (BM), 17, 21, 51 (MM), 17, 21, 49 (SV) | "This is the Beast Mastery Hunter as the site's guide describes it. ... The best players' habits for this spec are not written down here yet. This file says what the guide says and leaves the rest for a later full write."; "These are the buttons the guide's rotation names ... Where the guide or the rotation data gives a reason"; "For this spec the catalogue looks for ..." | Owner rule (guide voice): no site or sim meta talk; the round 25 brief applies it to the coach starters. Every hunter starter opens and frames its lists in talk about the site, the file and the data. Class-wide template, not hunter-specific. | Generator: `persona_starter.py` `who_section` and the priority lead-in, `starter_kit.consumable_paragraph`. Speak as the player ("The Beast Mastery hunter: the cat does a big share of the damage ..."; "Your buttons, in order:"; "Before a boss you want a flask, food and ..."). Regenerating does not fix it. |

Counts: CONFIRMED 13, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. Regenerating fixes 2 (R25-9, R25-10); 6 need `starter_overrides.py` (R25-1, 3, 4, 5, 11, 12, plus R25-6's override); 5 need generator changes (R25-2 `starter_kit.kit`, R25-6 `starter_priority`, R25-7 and R25-8 `starter_kit.DRILLS`, R25-13 `persona_starter.who_section` and `starter_kit.consumable_paragraph`).

## Shoulds (wording, not findings)

- **S1. `index.md:33`, Survival.** The class page's "so it takes that cooldown every time it comes up" is a ranged-spec rule on a page that covers Survival, which fires no shots (SV:92). "... so Beast Mastery and Marksmanship give it that cooldown every time it comes up" scopes it.
- **S2. MM coach:17 and `marksmanship.md:48`.** "Aimed Shot and Multi-Shot on their shared cooldown" (round 24 S3 carried). Fixing MM:48 to "Aimed Shot on the shared cooldown, Multi-Shot on packs" regenerates the coach line too.
- **S3. SV coach:34.** "check the message, because either can miss or be resisted": "either" means Disengage and Feign Death in SV:94, but the coach splits them into two lines. "check the message, because it can be resisted."
- **S4. BM coach:34, "Mend Pet."** A between-pulls button in BM:101, listed bare in the rotation.
- **S5. SV coach:24, Falcon line.** True at 60 (Falcon matches the highest Hawk rank, 110 for rank 6 at 58, the same as Beast rank 4; Hawk rank 7 at 120 is the Ahn'Qiraj rank the page names). The generator's own test (`test_survival_never_lists_the_ranged_aspect`) wants it out of the Survival list; dropping it also matches "From 30 to 59 Beast is the only melee aspect you have".
- **S6. SV coach:30, Mongoose Bite "while Expose Prey is up: it is a counterattack".** Expose Prey is a talent whose proc opens the window from 35, not a buff that is "up". "Mongoose Bite the moment it lights up: after a dodge, or from 35 when Expose Prey opens it off your hits on a marked target."
- **S7. MM coach:28, Sniper Shot.** A raid-build talent at 40; the leveling build never has it. "With the raid build, Sniper Shot ..." as MM:96 says.
- **S8. BM coach:28, "Serpent Sting when it is about to drop."** True, but drops the card's "never while the shared cooldown is about to come back, unless under a quarter of your mana".
- **Carried from round 24:** S2 (Multi-Shot row's first sentence at 20 to 60), S5 (BM:99 and MM:105 "let Aimed Shot have the shared cooldown"), S6 (BM:61); and round 23's carried Shoulds (SV:142 "Throw Immolation Trap", SV Wing Clip "cheap", Falcon wording, Strider Kick "a third melee special", the instants bullet, index:28 Survival pet).

## Superlative, comparison and quoted value sweep

Pages and cards: every round 24 line rechecked above; the rest is unchanged from round 24's sweep (Sniper Shot hardest hit, 590 at rank 3, 4 s cast, 15 s cooldown; Multi-Shot worst per mana from 18; sting and Arcane Shot per mana; Explosive Trap 520 against Volley 490; mark at -20 s, two minute duration; Rapid Fire five minutes raid, three MM leveling from 26, four BM leveling from 57; Bestial Wrath two minutes, 18 s, 50%). All hold at every band on both orders.

Coach lines:

| line | claim | verdict |
|---|---|---|
| BM coach:25 | Bestial Wrath every two minutes | OK (19574 cooldown 120000) |
| BM/MM coach:29/31 | Arcane Shot "more than a global away or under a quarter of your mana" | OK (APL `spellTimeToReady(2643) > 1.5s` or `<= 25%`) |
| BM/MM coach:36/35 | Volley a mana-heavy channel | OK (350 to 490, channeled) |
| SV coach:23 | mark first so the mana is back by the pull | OK (prepull -20 s) |
| SV coach:24 | Falcon the same melee attack power as Beast at 60 | OK (S5) |
| SV coach:27 | Counterattack free | **HUNTER-R25-3** |
| SV coach:28 | Strider Kick on its own short cooldown | OK (8000 ms) |
| SV coach:33 | Disengage five second cooldown, turns off auto attack | OK (781, 14272, 14273 category 5000; "Disables auto-attack") |
| SV coach:49 | combat potion Major Mana Potion | **HUNTER-R25-2** |
| all, drill.uptime | recast with 3 seconds left | **HUNTER-R25-8** |
| all, drill.ranks | top rank on every button | **HUNTER-R25-7** |

## Decisions

- **D1.** The coach findings are CONFIRMED, not Shoulds: each is a statement a player reads (or the coach acts on) that is false at the levels it covers or contradicts its own page and card. R25-13 is filed under the owner's voice rule the brief names for the coach files; it is class-wide and will appear in every class report.
- **D2.** R25-12 (Growl) is a finding and Mend Pet a Should: Growl is a button the page tells a raider to switch off, Mend Pet is merely out of place.
- **D3.** "Regen fixes it" was decided by running the generator read-only into the scratchpad, not by inference. Only R25-9 and R25-10 change.
- **D4.** No guide-page or card-note finding this round: every round 24 fix holds against the client rows, the shipped orders and the rank rows, and no other line changed.
- **D5.** Not re-filed: anything routed in the ledger's Round 19 to 24 sections; the site card's per-race rank rows (Survival 26 to 28, sent to sim round 22); the Leveling strip and cooldown cards (component lane).
- **D6.** No web research needed: the client data, the cards, the presets and the pages decide every call.

## Sims requested

None new. Carried: Marksmanship raid 60 Auto Shot share after A1/A2 (round 21, still routed).

## Sim findings

None new.

## Engine questions

None new. Standing: auto shots through Sniper Shot's cast inferred from Aimed Shot logs (R24 E1); Immolation Trap attack power share; sting and Arcane Shot share read at cast against per tick; Falcon's flat 5%; Explosive Trap and Volley attack power shares.

## Sources

**Client data, 1.60.1.70790:** `spellconst/hunter.json`, `spelltips/hunter.json` (19306, 1242634, 20909, 20910 Counterattack; 781, 14272, 14273 Disengage; 3045 Rapid Fire; 19574 Bestial Wrath; 13555 Serpent Sting; 24118, 1299332 Lacerate; 19434, 20900 to 20904 Aimed Shot; 2643 Multi-Shot; 13813, 14316, 14317 Explosive Trap; 1510, 14294, 14295 Volley; 469145 Falcon; 13165 to 14322, 25296 Hawk; 13161, 1299445 to 1299447 Beast; 1317257 Strider Kick), `talents/hunter.json` (Rapid Killing, Resourcefulness, Expose Prey, Lacerating Strikes), `addon-data.json` rotations, `addon-builds.json`.

**Repository:** `sim/request/raid-consumes.json` (Survival top band `major_frenzy_potion`; BM and MM `major_mana_potion`), `tools/coach/persona_starter.py`, `starter_overrides.py` (GUIDE_LINES and GUIDE_ORDER for the three hunter specs), `starter_kit.py` (TRAITS, `kit`, DRILLS), `starter_text.py` (PARENTHETICAL), `starter_priority.py`, `test_persona_starter.py` (four red tests, one hunter), `design/reviews/guide-accuracy/round24/hunter.md`, `ledger.md` Round 24, `git diff ea9adea1c..HEAD`. Read-only regen output in the scratchpad (`gr25/hunter-coach/`).

**Web:** none consulted (D6).

## Note sweep

One line per row note: spec, row, bands, each named talent, spell, rank or value with the level the player has it (leveling and raid orders), verdict.

- BM Hunter's Mark (10 to 60): spell 6; ranks 1 to 4 by trainer (rows r1 10 to 20, r2 22 to 38, r3 40 to 52, r4 58 to 60); one mark per target. OK.
- BM Bestial Wrath (40 to 60): talent 40 both orders; Rapid Fire 26; 50%, 18 s, fear, two minutes. OK.
- BM Rapid Fire (26 to 60): spell 26; "from 40 with Bestial Wrath"; shared cooldown Multi-Shot 18, Aimed Shot 20; five minutes, four at 57 to 60 on the leveling order; damage racial guarded. OK (R24-3 fix).
- BM Aimed Shot (20 to 60): spell 20; rank 1 rows 20 to 58, rank 6 at 60; auto through the cast (engine A1); hits harder than Multi-Shot at every rank. OK.
- BM Multi-Shot (18 to 60): spell 18; quarter-mana gate; "From 20 Aimed Shot"; worst per mana among single-target shots at every band. OK (S2 carried).
- BM Serpent Sting (10 to 60): ranks 2 to 8 by trainer; "from 18" gates; quarter-mana clause; 5% RAP a tick; last sentence names Multi-Shot from 18 and Aimed Shot from 20, above a quarter of mana. OK (R24-4 fix).
- BM Arcane Shot (10 to 60): ranks 1 to 8; "from 18" clauses; quarter-mana clause; 15% RAP; instant. OK.
- MM Hunter's Mark (10 to 60): as BM. OK.
- MM Rapid Fire (26 to 60): spell 26; five minutes raid order; three minutes leveling order from 26 (Rapid Killing 25, 26); damage racial guarded. OK (R24-2 fix).
- MM Aimed Shot (20 to 60): spell 20; rank 1, rank 4 at 50, rank 1 at 52 to 58, rank 6 at 60; no talent named; outhits Multi-Shot at every rank. OK (R24-5 fix).
- MM Multi-Shot (18 to 60): as BM. OK.
- MM Sniper Shot (40 to 60): raid talent 40; ranks 40/48/58; 4 s cast, 15 s cooldown; auto through the cast (E1); hardest hit. OK.
- MM Serpent Sting (10 to 60): as BM without the last sentence; Improved Stings not named. OK.
- MM Arcane Shot (10 to 60): as BM. OK.
- SV Hunter's Mark (10 to 60): spell 6; Expose Prey "from 35" (35, 36); mark before the pull. OK.
- SV Raptor Strike (10 to 60): ranks 2 to 8 with race rows 26 to 28 and 48 to 49; replaces the next white hit; Resourcefulness "at 33 and 34", 60%. OK.
- SV Counterattack (30 to 60): talent 32 (rows at 30, 31: component lane); after a parry; cannot be dodged, blocked or parried. OK.
- SV Mongoose Bite (16 to 60): spell 16; 5 s dodge window; Expose Prey "at 35 and 36", 10 s, any hunter's mark. OK.
- SV Lacerate (30 to 60): spell 30; names Mongoose Bite (16). OK.
- SV Strider Kick (30 to 60): talent 30; 8 s; "the bleed" Lacerate 30; 5.81% base mana. OK (carried Should on "sliver").
- SV Immolation Trap (16 to 60): ranks 16/26/36/46/56; on cooldown; Clever Traps 20 and 21, Survivalist's Discipline 37 and 38 as a "once ... are in" clause. OK.
- SV Wing Clip (12 to 60): ranks 12/38/60; Windfury Totem conditional; main hand hit with procs; little damage of its own. OK (carried Should on "cheap").

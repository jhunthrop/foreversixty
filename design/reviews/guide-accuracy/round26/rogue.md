# Guide accuracy round 26: Rogue

Audited on the round 26 audit tree: main after guide round 25 landed (head ca06416cd). It includes sim round 22 (lane commits a0a2818ed, dd8f2408f, 9f0adec67; merged into int/r22 at 83a03b22b and into the guide lane at fd758eefc), engine pin bff826207 (`sim/enginever/version.go`). Round 25's rogue audit ran on 438f8ba15, before that merge, so sim round 22's rogue work is new here. Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files read, line by line:

- `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`.
- Every note in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`, at every band `addon-data.json` renders (Assassination 23 bands, Combat 22, Subtlety 23), against both orders in `addon-builds.json` (decoded with `scratchpad/gr25/rogue_orders.py`, output `scratchpad/gr26/rogue_orders.txt`) and `trainables/rogue.json`.
- The coach, from scratch: `data/curated/coach/rogue-{assassination,combat,subtlety}.md` and `.json`, the rogue entries of `data/pipeline/coachrules.py`, `tools/coach/starter_overrides.py` (rogue GUIDE_LINES and GUIDE_ORDER), `starter_kit.py` (RANKS_DRILL, DRILLS), and the checks that read the rule files: `api/internal/coach/f1_core_casts.go`, `f7_death.go`, `copy.go`, `rules.go`.
- Generator, read-only: `persona_starter.py --spec rogue-<spec> --out scratchpad/gr26/rogue-coach/rogue-<spec>.md`. All three outputs are byte-identical to the committed starters, so every coach line below survives a regen unless its source (page line, card note, override or kit) changes.

Read first: the round 25 rogue report, the ledger's Round 19 to 25 sections, `design/reviews/2026-11-05-sim-audit-round22/rogue.md`, `integration.md` and `platform.md` (T1 and U1 talent-row floors), round 23 and 24 rogue reports (for the fixes reverted below), round 26 hunter, druid, priest, paladin, mage and warlock (for the shared roots).

## New work: sim round 22 (rogue)

`git diff 438f8ba15..HEAD` on every rogue file:

| change | where | verdict |
|---|---|---|
| K3 Assassination raid build 31/20/0 | assassination.md :7, :37, :52, :61 to :66 | Build line, key talents, toll and Composure match the raidBuild string (decoded: Malice 5, Ruthlessness 3, Improved Slice and Dice 3, Vigor 1, Blood Splatter 1, Lethality 2, Vile Poisons 5, Cold Blood, Relentless Strikes 5, Venom, Seal Fate 3, Mutilate; Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 3, Endurance 2, Composure 2, Flawless Execution, Weapon Expertise 2, Aggression 2). The trade list at :52 is exact (12 points each way). But the build drops Improved Poisons, which left the Venom "half your main-hand hits" lines behind: **ROGUE-R26-6** |
| K3 three-point Slice and Dice | Assassination card row 2 note; assassination.md :80 | The rule matches the condition. The reason ("over 20 seconds") holds only with all three points of Improved Slice and Dice, which no shipped order and not the level 60 build has, and the note names Mutilate on bands 10 to 39: **ROGUE-R26-5** |
| S1 Hemorrhage never at five | Subtlety card Hemorrhage note; subtlety.md :80 | OK (5 is the cap; the row renders from 30). The coach was not updated: **ROGUE-R26-8** |
| K1 one-hand axes | index.md :26 ("One-handed axes are new for rogues"); combat.md :56, :104, :128, :132 | OK (ClassMask 79 against Era's 71, per K1) |
| Raid order in addon-builds.json equal to the leveling order | `addon-builds.json` rogue-assassination raid | Routed to sim round 23 (not re-filed). Dependent lines in D2 |
| The same lane commits, merged from a pre-round-23 base | subtlety.md :76, :86, :90, :105, :135; assassination.md :86 | Six lines fixed in rounds 23 and 24 are back to their old text: **ROGUE-R26-1 to R26-4** |
| T1 / U1 talent rows (platform) | Assassination Cold Blood row now from 33; Subtlety Ghostly Strike row from 26 | OK. This closes round 25's "rows 25 to 32 early" and "rows 10 to 24 early" carries |
| trainables talent levels (Cold Blood 25, Blade Flurry 30, Preparation 30, Adrenaline Rush 40, Seal Fate 35, Puncturing Wounds 15, Riposte 20, Ghostly Strike 25) | data | Tier floors; no page or coach line quotes them |

How the reverts happened: the lane wrote dd8f2408f and 9f0adec67 on a base from before guide rounds 23 and 24, and touched the same Subtlety paragraphs. The merge 83a03b22b took the lane's paragraphs whole. `git show` confirms: `fd758eefc^1` (guide lane) has "any damage racial" twice in subtlety.md, while `fd758eefc^2` and HEAD have it zero times. The push-time note-revert check reads card notes only, so it did not see these prose lines.

## Round 25's fixes

- Coach overrides (R25-1 to R25-8) are in `starter_overrides.py` and in the committed starters. Rupture and the Death Mark condition are back for Subtlety. Garrote and Premeditation now cover the fight, and Backstab and Sinister Strike are there before 30. The levels are on Cold Blood, Mutilate, Adrenaline Rush, Blade Flurry, Ghostly Strike and Death Mark, and Kidney Shot is right by rank. The poison hands follow the level. All OK, except the Hemorrhage override (R26-8), which predates sim round 22.
- R25-9 to R25-12 (racial wording, uptime drill, consumables from the preset, meta talk): OK in all three starters.
- Cold Blood "from thirty-three": true on both shipped orders today. See D2 for the routed raid order.

## Findings

Counts, guide pages and cards: CONFIRMED 6 (R26-1 to R26-6), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

Counts, coach: CONFIRMED 5 (R26-7 to R26-11), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R26-1 | CONFIRMED | page prose (round 23 fix R23-3 reverted by the sim round 22 merge) | subtlety.md:76, subtlety.md:90, assassination.md:86 | ":76 Shank on cooldown, right behind your racial and trinket"; ":90 line your racial, trinket and potion up at the pull"; "A:86 Line up Cold Blood with your racial, trinket and potion." | Undead is the Horde pick on both pages (`recommendedRaces`, both Races sections). Its actives are Will of the Forsaken and Cannibalize (`races.json`), neither a damage button, and its damage racial Touch of the Grave is passive. Human and Dwarf have no damage active either. So the line tells an Undead rogue to burn a fear break for nothing. Round 23 fixed all three lines. The same pages still carry the fixed wording elsewhere (assassination.md:80 "any damage racial"; the Combat note "your damage racial if your race has one"). So the reverted lines also contradict their own page. | Restore round 23's text: ":76 right behind any damage racial and your trinket"; ":90 line any damage racial, your trinket and potion up at the pull"; "A:86 Line up Cold Blood with any damage racial, your trinket and potion." Extend the note-revert check to guide prose (diff each merged guide line against the guide lane's parent). |
| ROGUE-R26-2 | CONFIRMED | page prose (R23-5 reverted) | subtlety.md:86 | "so a Shank left on cooldown costs you energy and points at once" | A Shank on cooldown cannot be pressed. What costs energy and points is a Shank left sitting ready, which is how assassination.md:86 ("Never let Shank sit ready") and combat.md:81 say it. :76 on the same page uses "on cooldown" to mean the opposite. | "so a Shank left sitting ready costs you energy and points at once." |
| ROGUE-R26-3 | CONFIRMED | page prose (R23-2 reverted; reason backward on the page's own weapons) | subtlety.md:105 | "Hemorrhage and Ghostly Strike hit harder with a dagger, but most of your damage comes from ..." | The tooltip percent is higher with a dagger (14278: 125%, 180% with a dagger; 16511: 100%, 145%). The hit is not higher. The page's 60 table still holds Teebu's Blazing Longsword (`bis/rogue-subtlety.json`, unchanged). On that sword both strikes hit harder than on the other trees' ilvl 65 dagger at every attack power (round 23's arithmetic: Ghostly Strike 637 against 554 at AP 1800; Hemorrhage 446 against 428). index.md:26 still carries the fixed wording, so the two pages now disagree. | Restore: "Hemorrhage and Ghostly Strike pay a bigger weapon percentage with a dagger, but a slow sword's bigger hits make up for it, and most of your damage comes from Rupture, Garrote, Death Mark, Shank and your white swings, and none of them care what kind of weapon you hold." |
| ROGUE-R26-4 | CONFIRMED | page prose (R24-1 reverted; "on cooldown" against the card's own hold) | subtlety.md:135 | "Shank comes at 24 for every rogue: press it on cooldown." | The paragraph runs from 24 to 60. Boot Knife comes at 35 (5/5 at 39 on the one order; 1329286, 100% for two points at 5/5). The card casts Shank only at three combo points or fewer, or under 50 energy (row 1), and :76 says "At four or five with 50 energy or more, finish first and Shank right after". At five points a Shank spills both Boot Knife points. Sim round 21 K1 measured the hold at +0.2 to 0.4% at 50 to 59. index.md:30 keeps the qualified wording. | Restore round 24's text: "Shank comes at 24 for every rogue: press it on cooldown, and once Boot Knife comes at 35, finish first when you sit on four or five combo points with 50 energy or more." |
| ROGUE-R26-5 | CONFIRMED | card note and page prose (sim round 22's reason holds only on a build no order ships; a spell named before it is trained) | `rogue-assassination.json` Slice and Dice note, second and third sentences (rows 10 to 60); assassination.md:80 | Note: "With Improved Slice and Dice a three-point cast still runs over 20 seconds, and getting the attack speed back beats waiting a Mutilate for the fourth point." Page: "with Improved Slice and Dice even three points keep it up for over 20 seconds." | Slice and Dice runs 15 s at three points (5171 and 6774: `${15*$<mult>}`). Improved Slice and Dice is 15% a point (14165), as the engine has it (`slice_and_dice.go`:29, 1.15 / 1.3 / 1.45). Three points of it give 21.75 s, but only the raidBuild string takes three.<br>- The level 60 build takes two (23 and 24 on its order): 19.5 s.<br>- Both shipped orders in `addon-builds.json` take two, so no rendered band has three.<br>- Before 23 there is none: 15 s.<br>So "over 20 seconds" is false on the level 60 build and on every rendered band, and the page's priority covers both builds (:82 scopes Rupture alone to the raid build). The card's three-point row has no talent gate and renders from 10. Its later clause names Mutilate on bands 10 to 39, where Backstab or Sinister Strike builds. The rule itself is measured (sim round 22 K3: +0.4% raid, ties or gains at bare 20 to 60). | Note: "... and at three once it has fallen off. Even three points keep it up for 15 seconds, over 20 with all three points of Improved Slice and Dice in the raid build, and getting the attack speed back beats waiting a builder for the fourth point." Page :80: "a three-point Slice and Dice once it has dropped: even three points keep it up for 15 seconds or more, over 20 with the raid build's Improved Slice and Dice." Regenerate addon-data.json and Data.lua. |
| ROGUE-R26-6 | CONFIRMED | page prose (left behind by sim round 22's build change) | assassination.md:56 (key talents of the raid build), assassination.md:80 | ":56 Venom ... and 20 percentage points to each poison's chance to land for 6 seconds, so Instant Poison procs on half your main-hand hits at 60."; ":80 and Instant Poison lands on half your main-hand hits at 60" | Instant Poison's base chance is 20% (8679 to 11340). Venom adds 20 points (1329328; `venom.go` `additivePoisonBonusChance`, added to the base in `poisons.go`:38). "Half" was 20 + 20 + 10 from five points of Improved Poisons (14113, 2% a point). The new raid build (31/20/0) takes no Improved Poisons. :66 even says it gives Improved Poisons up. So in the raid, Instant Poison lands on 40% of main-hand hits under Venom. Half is true only on the level 60 build, and :56 sits in the raid build's talent list. | ":56 ... so Instant Poison procs on 40% of your main-hand hits at 60, and half in the level 60 build with Improved Poisons." ":80 ... and Instant Poison lands on 40% of your main-hand hits at 60 (half with the level 60 build's Improved Poisons)." |
| ROGUE-R26-7 | CONFIRMED | coach (from page :80; generated) | `coach/rogue-assassination.md:28` | "a three-point Slice and Dice once it has dropped: with Improved Slice and Dice even three points keep it up for over twenty seconds." | As R26-5: 19.5 s on the level 60 build and the shipped orders, 15 s before 23. A boss pull at any level the coach reads falls on one of those. | Regen after R26-5's page fix (the generator reads assassination.md:80; no override exists for this line). If the page wording does not regenerate cleanly, add `GUIDE_LINES["rogue-assassination"]["Slice and Dice"]`: "At four or more combo points, Slice and Dice first if it is down or about to fall off, and at three once it has dropped: three points still last fifteen seconds, over twenty with the raid build's Improved Slice and Dice." |
| ROGUE-R26-8 | CONFIRMED | coach (round 25 override predates sim round 22) | `coach/rogue-subtlety.md:33` (`starter_overrides.GUIDE_LINES["rogue-subtlety"]["Hemorrhage"]`) | "Hemorrhage from thirty builds. It waits a moment when Ghostly Strike is about to come back, unless you are close to full energy." | Sim round 22 S1 gated Hemorrhage at four combo points or fewer (+0.6% raid 60). The card's note now opens "Hemorrhage at four combo points or fewer ... At five the extra point is wasted", and subtlety.md:80 says "never at five combo points, where its point is wasted". The persona's priority promises "Where a button has a reason or a condition, it follows the name", but it gives Hemorrhage without the new condition, so the coach reads a fifth-point Hemorrhage as on-list. Same shape as R25-1, a line that drops its condition. | Override: "Hemorrhage from thirty builds, at four combo points or fewer: at five the point is wasted. It waits a moment when Ghostly Strike is about to come back, unless you are close to full energy, and after a Ghostly Strike until the Garrote is in." |
| ROGUE-R26-9 | CONFIRMED | coach (spells named with no level, told to the player at boss pulls before they arrive; R25-8's shape on trainer spells and the talents it missed) | `coach/rogue-assassination.md:27, :37, :38, :40`; `rogue-combat.md:27, :34, :35, :38`; `rogue-subtlety.md:24, :31, :35, :36, :37, :38, :40` | e.g. "Shank on cooldown."; S "Premeditation a couple of seconds before the pull ... Press it again every two minutes"; S "Preparation finishes every other rogue cooldown, once every ten minutes."; "Kidney Shot costs twenty-five energy ... Use it on casters"; "Kick, Gouge and Blind handle the rest"; S "Cheap Shot opens on trash that will die fast" | The coach reads boss pulls from the low teens (round 25 D1, PRIEST-R25 D1). Each of these is told as a button to press with no level, and none of them is there at 13. Levels from `trainables/rogue.json` and the orders:<br>- Shank: 24 (all three specs).<br>- Premeditation: talent at 20.<br>- Preparation: talent at 34.<br>- Kidney Shot: 30 (all three specs).<br>- Cheap Shot: 26 (Subtlety's positive line; the Assassination and Combat lines only rule it out).<br>- Blind: 34 (Assassination and Subtlety).<br>- Feint: 16 (all three specs).<br>- Kick: 12 (Combat).<br>Round 25 gave levels to Cold Blood, Mutilate, Adrenaline Rush, Blade Flurry, Ghostly Strike, Death Mark, Hemorrhage and Backstab, so these lines stand out against their own file. F1 does not judge a never-cast spell below its level (`notYetLearned`), so the code holds. The persona text is still wrong at those levels. | Overrides, one per line, e.g. "Shank from twenty-four on cooldown ..."; "Premeditation from twenty, a couple of seconds before the pull ..."; "Preparation from thirty-four finishes every other rogue cooldown ..."; "Kidney Shot from thirty costs twenty-five energy ..."; "Cheap Shot from twenty-six opens on trash ..."; "Kick, Gouge and, from thirty-four, Blind handle the rest". Generator option (round 25 R25-8): prefix "From N," from `trainables` and `addon-builds.json` to any positive line whose spell is learned above the lowest dungeon level. |
| ROGUE-R26-10 | CONFIRMED | coach (kit drill false for an energy class) | `coach/rogue-*.md` drill.ranks (A:60, C:58, S:60); `starter_kit.RANKS_DRILL` | "Sometimes the right one is a lower, cheaper rank." | No rogue rank is cheaper. Every rank of every rogue spell costs the same energy (`spelltips/rogue.json`):<br>- Sinister Strike 45 at ranks 1 to 8; Backstab 60 at 1 to 9; Eviscerate 35 at 1 to 9.<br>- Garrote 50, Rupture 25, Slice and Dice 25, Kick 25, Gouge 45 at every rank; Expose Armor 25 at every rank.<br>The rogue rule files carry `downranks: []`, and every rendered rogue row casts the top trained rank. So the sentence is false for a rogue at every level, and it invites a downrank that only loses damage. Druid filed the same text as a Should (DRUID-R26-S2), because druid mana ranks can be cheaper. For rogue it is false outright. | Kit: drop the second sentence of RANKS_DRILL when the spec's rule file has no downranks, or when the resource is energy or rage. Rogue reads: "Visit your trainer every time you level, then check each button on the priority list against the rank your level calls for." |
| ROGUE-R26-11 | CONFIRMED | coach rule file and F7 (a defensive offered at deaths it cannot prevent) | `coach/rogue-*.json` defensives (Evasion first); `coachrules.py:88, :94, :100`; `f7_death.go` `defensivesAt`, `pick` | F7 output, for any rogue death: "Died at {t}, Evasion not used" / "{Killer} killed you at {t}. Evasion was not used this pull." (or "with Evasion ready") | Evasion is "dodge chance increased by 50% for 15 sec" (5277). All three pages say it "does nothing against spells" (assassination.md:97, combat.md:96, subtlety.md:97 "Neither helps against spells"). F7 picks the first defensive proven ready, or else the first never cast this pull (`pick`). It never looks at the killing hit (`deathHit` carries only SpellID, Name, Amount). Evasion is listed first and is rarely pressed on a raid boss, so nearly every rogue death names it, including deaths to fire, breath and shadow volleys. The coach then calls Evasion the missed save, which the page says it cannot be. HUNTER-R26-S1 (Feign Death) is the milder version: Feign Death at least drops threat. | Rule file: tag Evasion melee-only (a `melee_only` field in coachrules.py and the rule JSON). F7: skip a melee-only defensive unless the killing hit is a melee swing (`KillerMelee`) or a physical ability. Vanish stays (the pages keep it as the emergency button). Test first: a rogue death to a fire spell names no Evasion. |

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R26-S1 | coach/rogue-combat.md:33 (Blade Flurry) | Omits combat.md:87 "If a fight brings adds in the first minute, hold Blade Flurry for them" and :89 (never next to a crowd-controlled mob). It is not false, and F1's lost-cast threshold (2 casts, 10%) does not fire on a hold of under a minute. | "Blade Flurry from thirty-three, at any energy, for its twenty-five; hold it for adds coming in the first minute, and never beside a Sapped or Polymorphed mob." |
| ROGUE-R26-S2 | coach/rogue-subtlety.md:29 (Slice and Dice) | Gives only the Rupture-ticking condition. Before Rupture at 20 the card casts it whenever it is down, which :30 (Eviscerate) half covers. | "Slice and Dice at three or more combo points while Rupture ticks, or whenever it is down before Rupture trains at twenty." |
| ROGUE-R26-S3 | coach/rogue-*.md Garrote opener (A:24, C:24, S:25) | Garrote trains at 14, so a level 13 Ragefire Chasm pull has none. Marginal: one level. | "From fourteen, open with Garrote ..." |
| ROGUE-R26-S4 | Subtlety Death Mark row | Sim round 22 integration.md:110 hands "the Subtlety Death Mark row wording" to the guide lane as a Should, without the text. The note is true at every band (sweep line 19). | Ask the integrator for the reviewer's line. |
| ROGUE-R26-S5 | carried | R25-S1 (Subtlety Shank "on cooldown" at 24 to 34), R25-S2 (Assassination coach "lands into unshredded armor" now reads whole: closed), R25-S4 (world buffs: closed by the round 25 wording), R24-S1, R24-S2, R24-S4, R23-S1, R23-S3, R23-S4, R22-S1 to S4, R21-S1, R21-S3, R21-S4, R19-S6. | As filed. |

## Superlative, comparison and quoted value sweep (new and changed lines)

| line | claim | check | verdict |
|---|---|---|---|
| A:52 | "31/20/0" | 31 + 20 + 0 from the string | OK |
| A:52 | trades: Vigor, Blood Splatter, one Improved Slice and Dice, Combat past Precision, for Improved Gouge, Remorseless Attacks, Improved Poisons, Master of Deception 1, Opportunity 1 | 12 points each way | OK |
| A:61 | Aggression 2: Eviscerate 10% harder | 18427 at 5% a point | OK |
| A:62 | Flawless Execution: 15 less | 1310711 | OK |
| A:63 | Weapon Expertise 2: 2% less | 30919 | OK |
| A:66 | "beat more poison procs and 20% on Mutilate" | Opportunity 5 is 4% a point, on Mutilate (14057) | OK |
| A:66 | "Two points each of Lightning Reflexes and Endurance are the toll to reach the Aggression row, and Composure on the way" | Tier 4 needs 15: CdG 3, LR 2, PW 3, Precision 3, Endurance 2, Composure 2 | OK |
| A:66 | "Composure ... another 10 energy off your opener" | 1329265 at 2/2: +10 on Poised's 10 | OK |
| A:56, A:80 | Instant Poison on "half" the main-hand hits | 20 + 20 on the raid build | R26-6 |
| A:80, card | three-point Slice and Dice "over 20 seconds" | 21.75 s only at 3/3 | R26-5 |
| S:80, card | Hemorrhage "never at five combo points, where its point is wasted" | five is the cap | OK |
| S:105 | "hit harder with a dagger" | sword hits harder on the listed weapons | R26-3 |
| S:86 | "a Shank left on cooldown costs you" | reversed | R26-2 |
| coach S:38 | "Blind comes back every three and a half minutes" with Elusiveness | 300 s less 90 at 2/2 (33; Blind 34) | OK |
| coach S:28 | Death Mark "comes back every ten seconds" | 10000 ms at every rank | OK |
| coach A:37, C:34, S:37 | Kidney Shot five at five, six from fifty | 408 and 8643 | OK |

## Decisions

- **D1. The six reverted lines are four findings (R26-1 to R26-4), all CONFIRMED.** Each was a confirmed error in round 23 or 24 and is false again on this tree. Their cause is a merge, not new wording. The fix restores the guide lane's text. The guard that would have caught them, the note-revert check, covers card notes only, so the fix column asks to extend it to guide prose.
- **D2. The routed Assassination raid order.** Not re-filed. Lines that depend on it:
  - Cold Blood "from thirty-three" (coach A:30): true on both orders today. If sim round 23 builds a raid order with Cold Blood earlier (the pre-round-22 one had 31), the line is still true for the leveling order and the persona has no raid-only player below 60 to mislead. Leave it.
  - The Rupture row's Blood Splatter gate and its note, "(the raid build)": the shipped raid order takes no Blood Splatter, so no rendered raid band casts it. Routed.
  - Improved Slice and Dice: neither shipped order takes the third point that R26-5's reason needs. R26-5 is filed anyway, because the level 60 build (two points) is false on its own.
  - Seal Fate "from 45" (Eviscerate note): true on the leveling order. On the raid order it waits for sim round 23.
- **D3. R26-6 is a finding, not a build question.** The raid build is adopted and measured (sim round 22 K3). The page states a value its own build no longer gives. The level 60 build keeps "half", so the fix names both.
- **D4. R26-10 is CONFIRMED for rogue, though druid filed the same text as a Should.** No rogue rank costs less than another. The sentence is false at every level for this class, and the rule files list no downranks.
- **D5. R26-11 is CONFIRMED, though hunter filed Feign Death as a Should.** Evasion has no effect on a spell death, and the pages say so three times. F7 names it first at nearly every rogue death, so the coach points at the one button the guide says would not have helped.
- **D6. Core lists (F1).** Cold Blood (3 min), Adrenaline Rush (5 min), Blade Flurry (2 min) and Ghostly Strike (20 s) are judged at base cooldowns, and no talent in any rogue build shortens them. Preparation only adds Ghostly Strike casts. The holds the guides ask for are short beside each cooldown, and F1 fires only at two or more lost casts and 10%:
  - Cold Blood for a four-point Eviscerate.
  - Ghostly Strike never at five.
  - Blade Flurry for early adds.
  None of them reaches F1's threshold. No rogue core ability shares a cooldown. Not filed.
- **D7. Defensive cooldowns quoted at base.** Endurance (Assassination raid, Evasion and Sprint 60% shorter) and Elusiveness (Subtlety, Vanish 90 s shorter) make F7 prove "ready" later than the truth, never earlier. That is conservative, not false.
- **D8. Race lines** are unmeasured on the new Assassination raid build (round 11 rule). None contradicts a measurement or the client. Not filed.

## Sims requested

- Optional: Assassination's three-point Slice and Dice row on the level 60 (leveling) build at bare 60 and at 40 and 50, against the old four-point-only gate. The aim is to confirm that sim round 22's "ties or gains at bare 20 to 60" was measured on the leveling order, where three points give 19.5 s. The integration notes say the scratch `rt` helper falls back to the leveling build for an empty third tree, so it probably was. This decides only whether R26-5's fix keeps the rule for both builds. The wording fix stands either way.
- Carried, already routed: the Assassination raid order (sim round 23); rogue BiS with one-hand axes and race re-runs at the re-rank; Subtlety band 20 weapon; the 50-energy Shank threshold on the re-ranked BiS.

## Sim findings

None new. Still open, not re-filed: Death Mark storage; the Mutilate and Shank off-hand damage on a 40+ log; Backstab's flat bonus (sim round 22 X2) on a 60 log.

## Engine questions

None new. Routed earlier and still open: trainer ranks at 60 that the cards do not cast (Backstab 9, Eviscerate 9, Feint 5, Deadly Poison V). Sim round 22 K2 found engine, cards and client aligned, and left the trainer question with the owner.

## Sources

- **Client 1.60.1.70790:**
  - `spelltips/rogue.json`: Instant Poison 8679 to 11340 (20%); Deadly Poison 2823 to 25351 (30%); Slice and Dice 5171 and 6774; Venom 1329328; Shank 1329283 and 1329549; Kidney Shot 408 and 8643; Evasion 5277; Blind 2094; Ghostly Strike 14278; Hemorrhage 16511; every rank's energy cost.
  - `talents/rogue.json`: Improved Slice and Dice 14165, Improved Poisons 14113, Composure 1329265, Endurance 13742, Aggression 18427, Flawless Execution 1310711, Weapon Expertise 30919, Opportunity 14057, Boot Knife 1329286, Elusiveness 13981.
  - `trainables/rogue.json`, `addon-builds.json` (decoded: `scratchpad/gr26/rogue_orders.txt`), `addon-data.json` rotations, `races.json`, `bis/rogue-subtlety.json`.
- **Engine bff826207:** `sim/rogue/slice_and_dice.go`:29, `poisons.go`:37-50, `venom.go`.
- **Curated:** `data/curated/apl/rogue-*.json`, `data/curated/coach/rogue-*.{md,json}`, `data/pipeline/coachrules.py`, `sim/request/raid-consumes.json`, `data/curated/presets.json`.
- **Coach code:** `api/internal/coach/f1_core_casts.go` (coreGaps, notYetLearned), `f7_death.go` (defensivesAt, pick, deathHit), `copy.go`, `rules.go` (CoreFiresLost 2, CoreFiresPct 10).
- **Generator:** `tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`. Read-only regen: `scratchpad/gr26/rogue-coach/` (identical to the committed files).
- **History:** `git diff 438f8ba15..HEAD`; commits dd8f2408f, 9f0adec67, 83a03b22b, fd758eefc.
- **Reviews:** round 23, 24 and 25 rogue; ledger Round 19 to 25; sim round 22 rogue, integration and platform; round 26 hunter (S1), druid (S2), priest, paladin, mage and warlock.
- Web research: none needed. Client data, engine code and the committed measurements decide every call.

## Note sweep

L = leveling order, R = raid order (shipped `addon-builds.json`; Assassination's R equals L, routed). Bands are the rendered ones.

**Assassination** (bands 10, 12, 14, 16, 20, 22, 24, 28, 30, 32, 33, 36, 38, 40 to 60)

1. Shank (24-60; 1329283, then 1329549 from 44): cannot miss; pays back more than it costs (15 or 25 for 5); Venom from 30 sits behind "with Venom"; Vile Poisons from 27. OK.
2. Slice and Dice (10-60; 5171, 6774 from 42): the first sentence matches the condition. The 20-second reason and "a Mutilate" at 10-39 fail: **R26-5**.
3. Rupture (20-60): "only with Blood Splatter (the raid build)"; +6 s (1329256); ticks crit. OK (the gate is routed, D2).
4. Cold Blood (33-60; row floor now 33): off the global cooldown, costs 0 (14177). OK.
5. Eviscerate (10-60): Mutilate 40; Seal Fate "from 45" (L45). OK.
6. Mutilate (40-60; 399956, 1241582, 1241584): two points, +20% poisoned. OK.
7. Backstab (10-60): builder until Mutilate; "far harder than Sinister Strike" (round 21's fix). OK.
8. Sinister Strike (10-60): 60-energy gate with a dagger. OK.

**Combat** (bands 10 to 60, 22)

9. Shank (24-60): every 10, 8 at five points of Flash of Steel (R 43; L four points, the clause is conditional). OK.
10. Eviscerate (10-60): at five. OK.
11. Slice and Dice (10-60): one or two points. OK.
12. Backstab (10-60): Puncturing Wounds L19 / R17; "in the raid build Lethality from 54 and Opportunity from 56" (R54, R56). OK.
13. Adrenaline Rush (40-60): damage racial if your race has one. OK.
14. Sinister Strike (10-60): only when Backstab cannot be cast. OK.

**Subtlety** (bands 10 to 60, 23; one order)

15. Shank (24-60): hold at four or five with 50 or more; Boot Knife from 35, 5/5 by 39. OK.
16. Ghostly Strike (26-60; row floor now 26): dagger against Sinister Strike; from 30 against Hemorrhage, 40 against 35 (30 at 57). OK.
17. Garrote (14-60): Ghostly Strike 26; Opportunity from 13; Hemorrhage and Death Mark "once you have them". OK.
18. Rupture (20-60): Serrated Blades from 21, Hemorrhage from 30. OK.
19. Death Mark (40-60; four ranks): 6 s, twice the rate, 10 s at every rank. OK.
20. Eviscerate (10-60): before 20 with Slice and Dice up; to 39. OK.
21. Slice and Dice (10-60): matches the condition. OK.
22. Hemorrhage (30-60): four or fewer (new, sim round 22); 15% to Rupture and Garrote; the holds match. OK.
23. Backstab (10-60): until 30; Opportunity from 13, 4% a point. OK.
24. Sinister Strike (10-60): before Hemorrhage, 60-energy gate. OK.

24 notes swept: 23 OK, 1 finding (R26-5).

## Coach sweep

Every coach line at every level a boss pull can be logged (round 25 D1).

### rogue-assassination.md

- :18 Who: OK.
- :22 lead-in: OK.
- :24 Garrote opener: OK (S3).
- :25 Ambush and Lethality: OK.
- :26 Cheap Shot trash only: OK (it rules out).
- :27 Shank: R26-9 (24). Venom conditional: OK.
- :28 Slice and Dice: R26-7.
- :29 Rupture only with Blood Splatter: OK.
- :30 Cold Blood from thirty-three: OK (D2).
- :31 Eviscerate at four or more: OK.
- :32 Mutilate from forty: OK.
- :33 Expose Armor: OK.
- :34 poisons at sixty with stones: OK.
- :35 Deadly MH 30 to 59, swap at 60: OK.
- :36 Backstab until Mutilate: OK.
- :37 Kidney Shot: ranks OK; R26-9 (30).
- :38 Kick, Gouge, Blind: R26-9 (Blind 34).
- :39 Redirect (1): OK.
- :40 Feint: R26-9 (16).
- :43 to :50 mistakes: OK (preparation by preset is handled in code this round).
- :52 consumables and world buffs: OK.
- :56 core_casts: OK.
- :57 uptime: OK.
- :58 idle: OK.
- :59 resource: OK.
- :60 ranks: R26-10.
- :61 preparation: OK.
- :62 death: OK.
- :63 adherence: OK.
- :67 to :76 voice: OK.
- Rule file: core Cold Blood OK (D6); aura Slice and Dice OK; defensives R26-11.

### rogue-combat.md

- :18 Who: OK.
- :24 Garrote: OK (S3).
- :25 Ambush and Backstab from stealth: OK.
- :26 Cheap Shot: OK.
- :27 Shank: R26-9 (24); "pays back more than it costs" OK.
- :28 Slice and Dice at one or two: OK.
- :29 Eviscerate at five: OK.
- :30 Backstab or Sinister Strike: OK.
- :31 Sinister Strike: OK.
- :32 Adrenaline Rush from forty: OK.
- :33 Blade Flurry from thirty-three: OK (S1).
- :34 Kidney Shot: R26-9 (30).
- :35 Kick: R26-9 (12).
- :36 Gouge (6): OK.
- :37 Redirect: OK.
- :38 Feint: R26-9 (16).
- :41 to :48: OK.
- :50: OK.
- :54 to :61 drills: OK, except :58 (R26-10).
- Voice: OK.
- Rule file: core Adrenaline Rush and Blade Flurry OK (D6); aura Slice and Dice OK; defensives R26-11.

### rogue-subtlety.md

- :18 Who: OK.
- :24 Premeditation: R26-9 (20).
- :25 Garrote: OK (S3).
- :26 Ghostly Strike from twenty-six: OK.
- :27 Rupture: OK.
- :28 Death Mark from forty: OK.
- :29 Slice and Dice: OK (S2).
- :30 Eviscerate: OK.
- :31 Shank: R26-9 (24); the condition matches the card.
- :32 Backstab and Sinister Strike: OK.
- :33 Hemorrhage: R26-8.
- :34 Ambush: OK.
- :35 Cheap Shot: R26-9 (26); Initiative full at 32 OK.
- :36 Preparation: R26-9 (34).
- :37 Kidney Shot: R26-9 (30).
- :38 Kick, Gouge, Blind with Elusiveness: R26-9 (Blind 34); the 3.5 minutes is OK.
- :39 Redirect: OK.
- :40 Feint: R26-9 (16).
- :43 to :50: OK.
- :52: OK.
- :56 to :63 drills: OK, except :60 (R26-10).
- Voice: OK.
- Rule file: core Ghostly Strike OK (D6); auras Rupture and Slice and Dice OK; defensives R26-11.

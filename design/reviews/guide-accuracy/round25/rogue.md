# Guide accuracy round 25: Rogue

Audited on the round 25 tree (head 438f8ba15: main after guide round 24; no sim round 22 in the log). Engine pin 170a59c39 (`sim/enginever/version.go`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited, line by line:

- `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`.
- Every note and row condition in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`, at every band `addon-data.json` renders (`rotations/rogue-assassination` 23 bands, `rogue-combat` 22, `rogue-subtlety` 22), against both shipped orders in `addon-builds.json` and the trainer levels in `trainables/rogue.json`.
- New this round: the coach starters `data/curated/coach/rogue-{assassination,combat,subtlety}.md`, the rogue entries in `data/pipeline/coachrules.py` (core, auras, defensives), and the generator `tools/coach/persona_starter.py` with `starter_guide.py`, `starter_priority.py`, `starter_overrides.py`, `starter_kit.py`, `starter_food.py`. Inputs: `sim/request/raid-consumes.json`, `data/curated/presets.json`.
- Coach scope: the coach reviews any logged boss pull at the player's level (PRIEST-R25 D1; `api/internal/coach/f5_ranks.go`). A coach line must hold at every level a boss pull can be logged, dungeon bosses from the low teens included.

Read first: round 24 rogue, the ledger's Round 24 section, round 25 hunter, priest, mage, druid and warlock (for the shared coach roots).

## Changes since round 24's tree (ea9adea1c)

- One rogue edit: subtlety.md:135, round 24's R24-1 fix ("press it on cooldown, and once Boot Knife comes at 35, finish first when you sit on four or five combo points with 50 energy or more").
- Cards, `addon-builds.json`, `trainables/rogue.json`, `spelltips/rogue.json`, the rogue BiS files, `raid-consumes.json` and `presets.json` are unchanged. The three rogue rotations in `addon-data.json` are identical to ea9adea1c (compared as parsed JSON).

## Round 24 fix, re-checked

- **subtlety.md:135 (R24-1).** True. Shank (1329283, 1329549) awards no combo point of its own: "dealing 50% weapon damage and generating 15 Energy" (25 at rank 2). So before Boot Knife (35 on the one order; 20% a point, 100% at 5/5 by 39) a Shank at five points spills nothing. Keying the hold to Boot Knife matches round 24 K1: no change at 24 to 34. It matches the card's hold from 35 and :76. The card still holds Shank at 24 to 34. That is a one-global ordering behind a ready finisher, worth nothing per K1, so "press it on cooldown" there is a practical rule, not a false one (D2; S1).

## Coach regeneration check

I ran `persona_starter.py --spec rogue-<spec> --out <scratch>/rogue-<spec>.md` for all three specs into `scratchpad/gr25/rogue-coach/`, then diffed against the committed files:

- Assassination and Subtlety: no difference.
- Combat: two lines differ. Both are round 23's "any damage racial" wording, which the committed starter predates (R25-9).

A regen fixes R25-9 only. The generator's own test (`test_persona_starter.py`) is red on main for four reasons the coordinator lists; none of the four is a rogue spec.

## Findings

No guide page or card note is false this round. All twelve findings are in the coach files.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R25-1 | CONFIRMED | coach: the main finisher is missing; the mark loses its condition | `coach/rogue-subtlety.md:21-36` (no Rupture line), `:28` | The priority has no Rupture line at all. ":28 Death Mark is your only damage cooldown and it comes back every ten seconds, so there is nothing to hold." | subtlety.md:80: "Rupture at four or more combo points whenever it is not ticking. **Death Mark** at four or more points while Rupture ticks." :70: "Then Ghostly Strike for a fifth point, Rupture right away". The card's row 4 is Rupture whenever it is not ticking, and row 5 is Death Mark only with `dotIsActive(Rupture)`. Rupture is one of the two uptime auras `coachrules.py` judges for this spec (`("Rupture", DOT)`). Yet the persona never names it, and its Death Mark line drops the Rupture condition that :83 insists on ("never press it into a lapsed bleed"). Cause: `starter_guide.RULED_OUT_NEXT` reads "A mark **with no** Rupture under it" (:83) as a rule-out of Rupture. `excluded_abilities` returns `{'Rupture'}` for Subtlety (traced in the scratchpad). Combat's exclusion of Rupture, from "Rupture is not worth a global on a boss", is correct. | Override and generator. `starter_overrides.GUIDE_LINES["rogue-subtlety"]`: "Rupture": "Rupture at four or more combo points whenever it is not ticking: it is the main finisher, and the bleed every mark feeds on."; "Death Mark": "Death Mark at four or more combo points while Rupture ticks, never into a lapsed bleed. It comes back every ten seconds, so a mark now beats a bigger one later." Place Rupture and Death Mark after Ghostly Strike and Garrote with a `GUIDE_ORDER["rogue-subtlety"]`. Generator: `RULED_OUT_NEXT` should rule out an ability after "with no" only when that phrase opens its clause, not when it follows another noun ("a mark with no Rupture"). Regenerating alone does not fix it. |
| ROGUE-R25-2 | CONFIRMED | coach: Garrote and Premeditation taught only as openers (the round 24 "Aimed Shot only the opener" shape) | `coach/rogue-subtlety.md:23, :32` | ":23 Open with Garrote once the tank has the boss ..."; ":32 Premeditation: A couple of seconds before the pull, press Premeditation for two combo points. They last twenty seconds, so do not press it early." | The card's row 3 casts Garrote after every Ghostly Strike at four points or fewer, all fight long. Its note: "Ghostly Strike lets it in without Stealth". Page :78: "Every Ghostly Strike that lands is followed by a Garrote at four points or fewer, even with the old one still ticking ... At exactly four with a Ghostly Strike just landed, Garrote first." Page :70: Premeditation "needs no Stealth in Forever, so press it again every 2 minutes in the fight whenever you have three combo points or fewer". :90 has the same, with Preparation. The persona gives both only as the pull. The coach would then read the in-fight Garrotes and Premeditations as off-list or wrong. | Override, `GUIDE_LINES["rogue-subtlety"]`: "Garrote": "Garrote opens from Stealth, and from twenty-six it follows every Ghostly Strike at four points or fewer, even with the old one still ticking: the strike lets it in without Stealth."; "Premeditation": "Premeditation a couple of seconds before the pull for two combo points, and again every two minutes in the fight whenever you have three or fewer. Spend them inside twenty seconds." |
| ROGUE-R25-3 | CONFIRMED | coach: no builder below 30 | `coach/rogue-subtlety.md:30` (and no Backstab or Sinister Strike line) | ":30 Hemorrhage builds, and waits a moment when Ghostly Strike is about to come off cooldown unless you are close to full energy." | Hemorrhage is the leveling and raid order's point at 30. The card builds with Backstab (row 9, "the builder until Hemorrhage") and Sinister Strike (row 10, only when Backstab cannot be cast, at 60 energy or more with a dagger) at every band from 10 to 29. subtlety.md:123 and :135 say the same. On a boss pull at 13 to 29 (Ragefire Chasm to Shadowfang Keep), the persona's only builder is a talent the player does not have. The generator reads only the Rotation section, which never names Backstab. | Override, `GUIDE_LINES["rogue-subtlety"]`: "Backstab": "Backstab from behind with a dagger in your main hand is the builder until Hemorrhage at thirty; Sinister Strike only when you cannot get behind or hold no dagger." "Hemorrhage": "From thirty, Hemorrhage builds, and waits a moment ..." (the rest unchanged). |
| ROGUE-R25-4 | CONFIRMED | coach: a level 60 raid value stated at every level and on the leveling build | `coach/rogue-combat.md:26` | "Shank on cooldown ...: an off-hand stab for five energy that cannot miss and gives forty back with Flash of Steel." | Shank gives 15 energy at 24 (1329283) and 25 from 44 (1329549). Flash of Steel adds 3 a point (1329284). The raid order reaches 5/5 at 43, so 40 holds from 44 on the raid build only. The leveling order takes 4 points (39, 41 to 43), and the page runs dungeons at 60 on it (combat.md:52), so the yield is 37. At 24 to 38 it is 15. The card note scopes it correctly ("every 8 with five points of Flash of Steel"). combat.md:77 states 40 inside the raid section, where it is true. | Override, `GUIDE_LINES["rogue-combat"]`: "Shank": "Shank on cooldown, right behind any damage racial and your trinket: an off-hand stab for five energy that cannot miss and pays back more than it costs, forty energy at sixty with all five points of Flash of Steel." |
| ROGUE-R25-5 | CONFIRMED | coach: a rank 2 value stated at every level | `coach/rogue-assassination.md:36`, `coach/rogue-combat.md:33`, `coach/rogue-subtlety.md:33` | "Kidney Shot costs twenty-five energy, sits on a twenty second cooldown and stuns longer per combo point, up to six seconds at five" | Kidney Shot rank 1 (408) trains at 30: "1 point: 1 second ... 5 points: 5 seconds". Rank 2 (8643) trains at 50: 2 to 6 seconds. From 30 to 49 the stun at five points is 5 seconds. The pages state 6 in their level 60 utility sections, where it is true. | Override, `GUIDE_LINES` for all three rogue specs: "Kidney Shot": "Kidney Shot costs twenty-five energy, sits on a twenty second cooldown and stuns a second longer per combo point: five seconds at five, six once rank two trains at fifty." Keep Assassination's "use it on casters and on whatever is hitting your healer". |
| ROGUE-R25-6 | CONFIRMED | coach: Initiative credited before the order gives it | `coach/rogue-subtlety.md:23, :26` | ":23 Initiative makes it a fourth combo point"; ":26 open with Cheap Shot, which Initiative makes three combo points" | Initiative (13976) is a 33%, 67% or 100% chance for one more point on Ambush, Garrote or Cheap Shot. It is taken at 29, 31 and 32 on both Subtlety orders, so the extra point is certain only from 32. The "fourth" also needs Premeditation (20). At 14 to 19 the opening Garrote is one point, at 20 to 28 three with Premeditation, and at 29 to 31 three or four by chance. Cheap Shot (26) is two points at 26 to 28. The page lines sit in the level 60 pull section. | Override, `GUIDE_LINES["rogue-subtlety"]`: "Garrote" as in R25-2, opening with "Open with Garrote once the tank has the boss: from thirty-two, with Premeditation and all three points of Initiative, it is your fourth combo point, and a bleed does not care that the boss has no Sunders on it yet."; "Cheap Shot": "On trash that will die fast, open with Cheap Shot, three combo points once Initiative is full, and four seconds of stun, then Eviscerate." |
| ROGUE-R25-7 | CONFIRMED | coach: poison hands given only for 60; the reason reads as a rule for every level | `coach/rogue-assassination.md:33, :34` | ":33 Instant Poison on the main hand and Deadly Poison on the off hand at sixty ..."; ":34 Deadly Poison: Your off hand lands more than enough hits to hold Deadly Poison at five stacks, and every extra proc on the main hand is worth more as an Instant hit." | assassination.md:88: "From 30 to 59, keep Deadly on the main hand and Instant on the off hand: at 30 to 39 Backstab stacks Deadly faster ... Swap them at 60." The raid preset agrees (`raid-consumes.json` `rogue-assassination`: `main_hand_imbue:deadly_poison` from 30 to 55, swapped at 60). The persona drops the 30 to 59 rule. Its :34 reason, with no level, tells a level 30 to 59 player that Instant belongs on the main hand: the reverse of the page. | Override, `GUIDE_LINES["rogue-assassination"]`: "Deadly Poison": "From thirty to fifty-nine, Deadly Poison on the main hand and Instant on the off hand. At sixty swap them: your off hand then lands more than enough hits to hold Deadly at five stacks, and every extra main-hand proc is worth more as an Instant hit." |
| ROGUE-R25-8 | CONFIRMED | coach: talent buttons told to the player with no level | `coach/rogue-assassination.md:29`; `coach/rogue-combat.md:31, :32`; `coach/rogue-subtlety.md:25, :28` | A:29 "Cold Blood at four or more combo points, right before the Eviscerate."; C:31 "Adrenaline Rush on cooldown ..."; C:32 "Blade Flurry costs twenty-five and can go at any energy."; S:25 "Ghostly Strike goes on cooldown too, just never at five points."; S:28 Death Mark "nothing to hold" | Each is a talent the shipped orders give late:<br>- Cold Blood: leveling order 33, raid 31.<br>- Adrenaline Rush: 40 on both orders.<br>- Blade Flurry: 33 on both.<br>- Ghostly Strike: 26.<br>- Death Mark: 40.<br>Boss pulls are logged from the low teens, and each line tells the player to press the button then. F1 does not assess a talent the log never casts (`f1_core_casts.go`). So the text is wrong, but no false finding follows from it. The other level-bound lines carry their level ("Backstab ... until Mutilate", "with Venom", "only with Blood Splatter"). | Override per line, adding the level, e.g. "From thirty-one in the raid build, thirty-three in the level sixty build, Cold Blood at four or more combo points, right before the Eviscerate."; "From forty, Adrenaline Rush ..."; "From thirty-three, Blade Flurry ..."; "From twenty-six, Ghostly Strike on cooldown, never at five points."; "From forty, Death Mark ..." (with R25-1's wording). Generator option: prefix "From N," from `trainables` and `addon-builds.json` to any line whose spell is learned above the earliest dungeon level. |
| ROGUE-R25-9 | CONFIRMED | coach: "your racial" given to every race | `coach/rogue-combat.md:26, :31` | ":26 right behind your racial and trinket"; ":31 Adrenaline Rush on cooldown, with your racial, a trinket and a potion." | Round 23 (R23-3) changed the page and card to "any damage racial" / "your damage racial if your race has one": Undead, Human and Dwarf have no damage racial (`races.json`), and sending them to a fear break or Stoneform is wrong. The committed starter predates it. | Fixed by regen: `persona_starter.py --spec rogue-combat` gives "right behind any damage racial and your trinket" and "with your damage racial if your race has one" (scratch diff). |
| ROGUE-R25-10 | CONFIRMED | coach: the uptime drill contradicts the cards (shared root, HUNTER-R25-8) | `coach/rogue-*.md` drill.uptime (A:56, C:54, S:53) | "Recast each with 3 seconds left" | Each spec's coach aura has its own refresh rule:<br>- Combat Slice and Dice: the card casts it only when it is down (`not auraIsActive`). combat.md:77 says "Do not refresh Slice and Dice early with five points. Let it fall, put one or two points back into it".<br>- Subtlety Rupture: the card casts it only when it is not ticking (row 4), and a recast clips the bleed Death Mark stores.<br>- Assassination Rupture: only when not ticking.<br>Slice and Dice under 3 s (Assassination, Subtlety) happens to match. | Generator, the shared root: `starter_kit.DRILLS["dps"]["uptime"]`, "Recast each as your priority says, no earlier ...". Regenerating does not fix it. |
| ROGUE-R25-11 | CONFIRMED | coach: the consumables kit by trait, not the raid preset (shared root, HUNTER-R25-2, DRUID-R25-17) | `coach/rogue-*.md` consumables (A:51, C:49, S:48), "Pull preparation" mistake (A:47, C:45, S:44), drill.preparation (A:60, C:58, S:57) | "the catalogue looks for a flask (Flask of the Titans); food (Poached Sunscale Salmon)"; "Pull preparation: a missing flask and food." | The rogue raid preset (`raid-consumes.json`, top band of each rogue spec) carries no flask. It carries Ground Scorpok Assay, Thistle Tea, Elixir of the Mongoose, Juju Power and Might, Major Frenzy Potion (from 55; "a rogue drinks no other potion", `presets.json` by_class rogue) and a sharpening stone in each hand. `kit()` picks Flask of the Titans from `Traits.style` "melee" and no potion from `Traits.resource` "energy". So the coach marks a rogue without a tank flask as unprepared and never looks for the potion or the stones. Assassination's own :33 says "with a sharpening stone beside each". The food matches the preset. | Generator, the shared root: `starter_kit.kit()` reads flask, potion and weapon enchant from the spec's top-band preset, the way `starter_food.preset_food` reads the food. |
| ROGUE-R25-12 | CONFIRMED | coach: site, file and data meta talk (voice rule; shared root, HUNTER-R25-13) | `coach/rogue-*.md:17, :21`, consumables paragraph | "This is the Assassination Rogue as the site's guide describes it. ... The best players' habits for this spec are not written down here yet. This file says what the guide says and leaves the rest for a later full write."; "These are the buttons the guide's rotation names ... Where the guide or the rotation data gives a reason"; "the catalogue looks for" | The owner's guide-voice rule (no site or sim meta talk), which the round 25 brief applies to the coach files. | Generator, the shared root: `persona_starter.who_section`, the `know_section` lead-in, `starter_kit.consumable_paragraph`. |

Counts: CONFIRMED 12, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

By fix route:

- Fixed by regen: 1 (R25-9).
- Needs an override in `starter_overrides.py`: 7 (R25-2 to R25-8), plus R25-1's lines.
- Needs a generator or kit change: 4. R25-1 (`starter_guide.RULED_OUT_NEXT`), R25-10 (`starter_kit.DRILLS`), R25-11 (`starter_kit.kit`), R25-12 (`persona_starter.who_section`, `starter_kit.consumable_paragraph`).
- Needs a hand edit: none. All three rogue starters are generated.

## Coach rules (`coachrules.py`)

- Assassination core Cold Blood: the page uses it ("press Cold Blood if it is ready and Eviscerate straight after").
- Combat core Adrenaline Rush and Blade Flurry: combat.md:87 "On a 3 minute kill you get two Blade Flurries and one Adrenaline Rush, so do not hold them". Holding Blade Flurry for early adds is a condition, not "do not use".
- Subtlety core Ghostly Strike: on cooldown at four points or fewer.
- Auras:
  - Slice and Dice for all three specs.
  - Rupture for Subtlety: the page keeps it ticking.
- Defensives: Evasion and Vanish.

No core ability is one the guide says not to use. Not filed.

## Shared roots: which apply to rogue

- **Template meta talk:** R25-12.
- **Recast with 3 seconds left:** R25-10.
- **Combat potion by resource:** R25-11, with the flask and stones.
- **"Top rank" drill:** not filed for rogue. No rogue page or row keeps a lower rank: every rogue row at every band casts the top trained rank. The one exception is Eviscerate rank 8 (11300) at 60 beside trainer rank 9 (31016), which is round 24's routed engine question on book ranks. The drill is true for this class as the pages stand (D3).
- **Tank uptime drill, leveling-only "no X" clauses, "only with N targets", a pet spell as the player's button:** none occurs in the rogue files. Subtlety's Rupture exclusion is a different parse ("with no", R25-1).
- **Short or missing spell names:** Subtlety's bare "Slice and Dice." (:29) is the generator's bare-name line. It is not false, so it is a Should (S3).

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R25-S1 | subtlety.md:135; subtlety.md:9 (description "Ghostly Strike and Shank on cooldown") | The card holds Shank at four or five points with 50 energy or more from 24, not only from 35. Before Boot Knife the hold is one global behind a ready finisher and changes nothing (K1), so "on cooldown" is the practical rule. | "press it on cooldown; at four or five with 50 energy or more, finish first, which matters once Boot Knife comes at 35" |
| ROGUE-R25-S2 | coach/rogue-assassination.md:24; coach/rogue-subtlety.md:25 | The generator split a sentence from its antecedent: "lands into that armor" (the armor sentence before it was not carried) and "goes on cooldown too". | "Ambush hits harder on paper, but it lands into unshredded armor and gets none of Lethality's bonus."; drop "too" (or use R25-8's wording). |
| ROGUE-R25-S3 | coach/rogue-subtlety.md:29 | A bare "Slice and Dice." with no rule. The page's rule is three or more points while Rupture ticks, when it is down or about to fall off. | Override: "Slice and Dice at three or more combo points while Rupture ticks, when it is down or about to fall off." |
| ROGUE-R25-S4 | coach/rogue-*.md world buffs | Same as MAGE-R25 D4: the preset says Forever raids have none, and the coach never holds them against anyone. | "In dungeons and raids Forever turns world buffs off, so none are expected." |
| ROGUE-R25-S5 | carried | R24-S1, R24-S2, R24-S4, R23-S1 (Blade Flurry "at any energy", now also at combat coach :32), R23-S3, R23-S4, R22-S1 to S4, R21-S1, R21-S3, R21-S4, R19-S6. R24-S3 is closed by :127's place under "Slice and Dice ... at 10" (unchanged). | As filed. |

## Prose, line by line (all four pages)

The pages are unchanged since round 24 except subtlety.md:135 (re-checked above). Round 24's line-by-line sweep was re-run on this tree and stands:

- index.md :22 to :37.
- assassination.md :30 to :133.
- combat.md :30 to :132.
- subtlety.md :29 to :135.

Its superlative and quoted-value table also stands, re-run over the R24-1 sentence: Boot Knife "at 35" is exact, and "four or five combo points with 50 energy or more" is the card's condition. Voice: no em or en dashes, and no site, sim or band talk on the pages. No page names 70755.

## Decisions

- **D1. Coach scope by level** follows PRIEST-R25 D1: boss pulls are logged from the low teens. R25-3 to R25-8 are findings although each line is true at 60.
- **D2. R24-1's fix stands.** Shank awards no combo point without Boot Knife (client text), so the card's hold at 24 to 34 orders it one global behind a finisher and spills nothing. The wording point is S1, not a finding.
- **D3. The top-rank drill is not filed for rogue.** No rogue page advises a lower rank, and every rendered rogue row is the top trained rank but Eviscerate 9 at 60, which is routed (engine question, round 24).
- **D4. R25-8 is one finding for five lines.** All five share one shape (a talent button with no level), and none makes F1 misjudge a log (talents never cast are not assessed). They are misleading persona text, so CONFIRMED, not a Should.
- **D5. R25-1 is not a Combat issue.** Combat's Rupture exclusion is right (combat.md:77). Only Subtlety's "with no Rupture" parse is wrong. Assassination lists Rupture with its Blood Splatter condition. That condition comes from the card note, which the generator correctly did not read as a group condition.
- **D6. Nothing routed in the ledger's Rounds 19 to 24 is re-filed.** The site card's per-race rank issue does not touch rogue (no rogue rank rows).

## Sims requested

None new. Carried and already routed: round 24's optional Subtlety Hemorrhage stop at four points; Subtlety band 20 weapon; rogue BiS and race re-runs at the re-rank; K1's 50-energy threshold on the re-ranked BiS.

## Sim findings

None new. Still open, not re-filed: Death Mark storage, and the Mutilate and Shank off-hand damage on a level 40+ log.

## Engine questions

None new. Round 24's is still routed: trainer ranks at 60 the cards do not cast (Backstab 9, Eviscerate 9, Feint 5, Deadly Poison V).

## Sources

- **Client 1.60.1.70790:**
  - `spelltips/rogue.json`: Shank 1329283 and 1329549, Kidney Shot 408 and 8643, Cheap Shot 1833, Garrote 703 to 11290, Ambush, Blade Flurry 13877, Cold Blood 14177, Premeditation 14183, Redirect 438040, Feint, Blind 2094, Gouge, Kick.
  - `talents/rogue.json`: Flash of Steel 1329284, Initiative 13976, Boot Knife 1329286, Venom 1329328, Vile Poisons, Lethality 14128, Elusiveness 13981, Opportunity, Relentless Strikes.
  - `trainables/rogue.json`, `addon-builds.json` (all six rogue orders decoded against `talents/rogue.json`; scratch decoder `gr25/rogue_orders.py`), `addon-data.json` (rogue rotations, compared to ea9adea1c), `races.json`.
- **Curated:** `data/curated/apl/rogue-*.json`, `data/curated/coach/rogue-*.md`, `data/curated/presets.json` (by_class rogue), `sim/request/raid-consumes.json`, `data/pipeline/coachrules.py`.
- **Generator:**
  - `tools/coach/persona_starter.py`, `starter_guide.py` (`RULED_OUT_NEXT`, `excluded_abilities`, `guide_order`, traced per rogue spec), `starter_priority.py`, `starter_overrides.py` (no rogue entries today), `starter_kit.py` (`kit`, `DRILLS`, `TRAITS`), `starter_food.py`.
  - Regen output (read-only): `scratchpad/gr25/rogue-coach/`.
- **Coach:** `api/internal/coach/f1_core_casts.go` (an uncast talent is not assessed); PRIEST-R25 D1.
- **Reviews:** round 24 rogue, ledger Round 24, round 25 hunter, priest, mage (D4 world buffs, with its web sources), druid (D5), warlock.
- Web research: none needed. Every call is decided by the client data, the cards, the preset and the committed measurements.

## Appendix A: Note sweep

The cards and their rendered rows are unchanged since round 24 (byte-identical rows). I re-checked every note against both orders and the trainer, with the same verdicts as round 24's appendix. L = leveling order, R = raid order.

**Assassination** (23 bands, 10 to 60)

1. Shank (24-60): rank 1 at 24 (15), rank 2 at 44 (25). It cannot miss. Venom L30 / R30 sits behind "with Venom"; Vile Poisons L27 / R25. 30%, 20 points, 6 s. OK.
2. Slice and Dice (10-60): four or more, under 3 s. OK.
3. Rupture (20-60): Blood Splatter is R22 only (`auraIsKnown` gate); +6 s. OK.
4. Cold Blood (25-60): L33 / R31. Rows 25 to 32 are early (component lane). Off the global cooldown, free. OK.
5. Eviscerate (10-60): Mutilate 40; Seal Fate "from 45" (L45; R43, carried S). OK.
6. Mutilate (40-60): two points, +20% poisoned. OK.
7. Backstab (10-60): until Mutilate, dagger. OK.
8. Sinister Strike (10-60): 60-energy gate (carried S). OK.

**Combat** (22 bands, 10 to 60)

9. Shank (24-60): every 8 with five Flash of Steel (R by 43; L 4 points, conditional clause). OK.
10. Eviscerate at five. OK.
11. Slice and Dice at one or two. OK.
12. Backstab: Puncturing Wounds L19 / R17; Lethality R54, Opportunity R56. OK.
13. Adrenaline Rush (40-60): "your damage racial if your race has one". OK.
14. Sinister Strike: only when Backstab cannot be cast. OK.

**Subtlety** (22 bands, 10 to 60; one order)

15. Shank (24-60): Boot Knife from 35, 5/5 by 39; the hold matches the condition. OK.
16. Ghostly Strike (10-60): talent 26 (rows 10 to 24 early, component lane). Dagger and Hemorrhage comparisons: R24-S4 carried. OK.
17. Garrote (14-60): Ghostly Strike 26; Opportunity 13; Hemorrhage 30 and Death Mark 40 "once you have them". OK.
18. Rupture (20-60): Serrated Blades 21, Hemorrhage 30. OK.
19. Death Mark (40-60): 6 s, twice the rate, 10 s cooldown at every rank. OK.
20. Eviscerate (10-60): before 20 with Slice and Dice up, while Rupture ticks to 39. OK.
21. Slice and Dice (10-60): matches the condition. OK.
22. Hemorrhage (30-60): 15% to Rupture and Garrote; the holds match. OK.
23. Backstab (10-60): until 30; Opportunity from 13 at 4% a point. OK.
24. Sinister Strike (10-60): before Hemorrhage, 60-energy gate. OK.

24 notes swept: 24 OK.

## Appendix B: Coach sweep

One line per coach line, at every level a boss pull can be logged (D1).

### rogue-assassination.md

- :17 Who: R25-12. The overview matches assassination.md:30.
- :21 lead-in: R25-12.
- :23 Garrote opener (14): OK.
- :24 Ambush and Lethality: Lethality's client list excludes Ambush. OK; the dangling "that armor" is S2.
- :25 Cheap Shot trash only (26): OK.
- :26 Shank with Venom (24; Venom 30 conditional; Vile Poisons 25/27): OK.
- :27 Slice and Dice at four or more: OK.
- :28 Rupture only with Blood Splatter: OK.
- :29 Cold Blood: R25-8.
- :30 Eviscerate at four or more: OK.
- :31 Mutilate is the builder: OK, read with :35 "until Mutilate".
- :32 Expose Armor: OK.
- :33 poisons at sixty: OK as scoped (stones: R25-11).
- :34 Deadly reason: R25-7.
- :35 Backstab until Mutilate: OK.
- :36 Kidney Shot: R25-5.
- :37 Kick, Gouge and Blind: OK.
- :38 Redirect (1; 15 s): OK.
- :39 Feint (16): OK.
- :42 to :49 mistakes: OK, except :47 (R25-11).
- :51 consumables: R25-11, R25-12, S4.
- :55 core_casts: OK.
- :56 uptime: R25-10.
- :57 idle: OK.
- :58 resource: OK.
- :59 ranks: OK for rogue (D3).
- :60 preparation: R25-11.
- :61 death: OK.
- :62 adherence: OK.
- :66 to :75 voice: OK.

### rogue-combat.md

- :17 Who: R25-12. Matches combat.md:30.
- :21: R25-12.
- :23 Garrote: OK.
- :24 Ambush and Backstab from stealth: OK.
- :25 Cheap Shot: OK.
- :26 Shank: R25-4, R25-9.
- :27 Slice and Dice at one or two: OK.
- :28 Eviscerate at five: OK.
- :29 Backstab and Sinister Strike: OK.
- :30 Sinister Strike: OK.
- :31 Adrenaline Rush: R25-8, R25-9.
- :32 Blade Flurry: R25-8 (S5 for "any energy").
- :33 Kidney Shot: R25-5.
- :34 Kick 25: OK.
- :35 Gouge: OK.
- :36 Redirect: OK.
- :37 Feint: OK.
- :40 to :47: OK, except :45 (R25-11).
- :49: R25-11, R25-12, S4.
- :53: OK.
- :54: R25-10.
- :55: OK.
- :56: OK.
- :57: OK (D3).
- :58: R25-11.
- :59: OK.
- :60: OK.
- :64 to :73: OK.

### rogue-subtlety.md

- :17 Who: R25-12. Matches subtlety.md:29.
- :21: R25-12.
- :23 Garrote opener: R25-2, R25-6.
- :24 Ambush needs a dagger: OK.
- :25 Ghostly Strike: R25-8, S2.
- :26 Cheap Shot: R25-6.
- :27 Shank (card note's first sentence): OK.
- :28 Death Mark: R25-1, R25-8.
- (missing) Rupture: R25-1.
- (missing) Backstab and Sinister Strike: R25-3.
- :29 Slice and Dice: S3.
- :30 Hemorrhage: R25-3.
- :31 Preparation (34; 10 min): OK.
- :32 Premeditation: R25-2.
- :33 Kidney Shot: R25-5.
- :34 Blind with Elusiveness (Blind 34; Elusiveness 2/2 at 33; 5 min less 90 s): OK.
- :35 Redirect: OK.
- :36 Feint: OK.
- :39 to :46: OK, except :44 (R25-11).
- :48: R25-11, R25-12, S4.
- :52: OK.
- :53: R25-10.
- :54: OK.
- :55: OK.
- :56: OK (D3).
- :57: R25-11.
- :58: OK.
- :59: OK.
- :63 to :72: OK.

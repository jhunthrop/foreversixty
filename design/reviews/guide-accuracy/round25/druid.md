# Guide accuracy round 25: Druid

Audited on the round 25 tree (`scratchpad/gr25/audit`, head 438f8ba15: main after guide round 24; no sim round 22 in the log). Engine pin 170a59c39 (`sim/enginever/version.go`). Active client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files audited:

- `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`.
- Every note and row condition in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered band of `addon-data.json` `rotations/druid-*` with `race_spell_ids`, against both shipped orders in `addon-builds.json` (decoded against `talents/druid.json`; unchanged since round 23), `trainables/druid.json`, `spelltips/druid.json`, `spellconst/druid.json`.
- New this round: the coach starters `data/curated/coach/druid-{balance,feral,feral-bear,restoration}.md` (and their `.json` data files), with the generator `tools/coach/persona_starter.py`, `starter_overrides.py`, `starter_kit.py`, `starter_guide.py`, `starter_priority.py`, `test_persona_starter.py`, and `sim/request/raid-consumes.json` (the raid preset). Coach scope: `docs/superpowers/specs/2026-10-27-log-coach-design.md` §6.1 (level from `AdvLevel`) and F5.

Read first: round 24 druid, the ledger's Round 24 section, round 25 hunter, mage and priest (shared coach roots).

## Changes since round 24's tree (ea9adea1c)

`git diff ea9adea1c HEAD` on the druid files: balance card row [5] note, bear rows [7] and [8] notes, balance.md :139 and :141, feral-bear.md :80. These are round 24's four fixes (DRUID-R24-1 to R24-4). Nothing else in the druid pages, cards, orders or BiS files moved. All four druid coach starters are unchanged by the round.

## Coach regeneration check

`persona_starter.py --spec druid-<spec> --out scratchpad/gr25/druid-coach/druid-<spec>.md` for all four specs, diffed against the committed files: **no difference for any of them**. A regen fixes none of the findings below. No druid coach file is a full persona (`FULL_SPECS` is warrior-fury, priest-holy, warrior-protection), so every druid coach fix goes through the generator or its overrides.

The generator test (`test_persona_starter.py`) has four red tests on this tree. The druid one is `test_every_correction_restates_its_guide`: `GUIDE_ORDER["druid-feral-bear"]` lists Swipe, and Swipe is not among the bear's priority names. Cause: `starter_guide.CONDITION` matches `only with(?! (?:spare )?mana\b)`, so feral-bear.md:81 "**Swipe** only with two or more targets in front of you" is read as a group condition (the "only with a Windfury Totem" shape) and Swipe is excluded spec-wide. DRUID-R25-7 gives the fix.

## Round 24 fixes, rechecked

- **Balance [5] note (R24-1).** "With one point the quicker Starfires do not pay for the Wrath's cast" re-derived on every one-point band the card renders:
  - Leveling 35 to 39: Wrath rank 4 (rendered) against Starfire rank 3, 0.17 s a charge. It does not pay.
  - Raid 39 and 40: Starfire 3.0 s (Improved Starfire 5/5) against Wrath rank 4 and 6. It does not pay.
  - "Half a second with all three" matches the talent text.
  - The clause "the loop still wins" is DRUID-R25-1.
- **balance.md :139 and :141 (R24-2).** Now consistent with the card's 35 start: "from then, cast one Wrath whenever no Eclipse charge is banked" matches row [5] (gate `auraIsKnown(408248)`, Eclipse 35 on the leveling order). "Until Moonkin Form and Innervate at 40" matches the order (Moonkin Form 40 both orders) and the trainer (Innervate 40). "The loop starts to pay at 40" matches the lane's measurement (identical at 36 and 39, +0.9 to +1.0% at bare 40 to 60). The mana reason for 40 on holds: Wrath rank 6 at half cost (40) against Starfire rank 3 (180), both under Moonglow. OK.
- **feral-bear.md:80 and bear [7] (R24-3).** "Seldom waits" and "nearly every time" match the lane's 29.45 of 30 Bites at 43. OK.
- **Bear [8] (R24-4).** "Put the rage into Maul instead, and into Lacerate from 25 and Primal Bite from 40": Lacerate is the 25th-level point and Primal Bite the 40th on both orders. Swipe renders from 16 (trainer). It makes no "does more" claim at the measured tie at 20. OK.

## Findings

Seventeen CONFIRMED: one card note, sixteen in the coach files.

| id | status | kind | where | quote | evidence | fix |
|---|---|---|---|---|---|---|
| DRUID-R25-1 | CONFIRMED | note contradicts the page at 35 to 39 (side effect of the R24-1 fix) | `druid-balance.json` row [5] note, second clause (renders 10 to 60; fires from 35 leveling, 39 raid) | "With one point the quicker Starfires do not pay for the Wrath's cast; the loop still wins because the Wrath costs a fraction of a Starfire's mana." | The leveling order has one Eclipse point from 35 to 60, so the clause covers 35 to 39. balance.md:141, rewritten in the same round: "Until Moonkin Form and Innervate at 40 your mana keeps you on Wrath most of the time anyway, so the loop starts to pay at 40." Sim round 21 (S1, quoted in DRUID-R24-2) measured the loop on and off as identical at 36 and 39, because the mana line casts Wrath nearly every global there. So at 35 to 39 the note says the loop wins and the page says it does not pay yet; the measurement agrees with the page. | Note: "...With one point the quicker Starfires do not pay for the Wrath's cast; from Moonkin Form at 40 the loop still wins, because the Wrath costs a fraction of a Starfire's mana." Regenerate `addon-data.json` (note text only). The raid-order bands 39 to 41 stay with the routed sim. |
| DRUID-R25-2 | CONFIRMED | coach: the spec's core Wrath missing; "spending the charges" with nothing banking them | `coach/druid-balance.md:24, :27` | ":24 Start a Wrath so it finishes as the tank pulls; it lands after the tank's first hit." ":27 Starfire everything else, spending the charges." | The Wrath entry carries only the pre-pull sentence (balance.md:81). The page's loop step 2 is "With no Eclipse charge banked, cast **Wrath**" (:88), and the Mana paragraph says "If you fall behind, lean on Wrath" (:95). The card has three Wrath rows: [5] the Eclipse Wrath, [6] the filler below 20 and the mana line, [2] the end of the fight. The list never says when to cast Wrath in the fight, yet it says to spend charges only Wrath banks.<br>The Starfire line is also false by level:<br>- Below 20 there is no Starfire (trainer 20), and Wrath is the filler.<br>- From 20 to 34 there is no Eclipse (35 leveling, 39 raid), so there are no charges, and the card casts Starfire only while the mana keeps pace (row [6] gate). | `starter_overrides.py`, GUIDE_LINES `druid-balance`:<br>- "Wrath": "Wrath whenever no Eclipse charge is banked, from thirty-five: one Wrath, then Starfires. Before twenty it is your filler, and when your mana falls behind the fight, lean on it."<br>- "Starfire": "Starfire everything else from twenty, spending the charges once you have Eclipse."<br>GUIDE_ORDER `druid-balance`: ("Moonkin Form", "Moonfire", "Insect Swarm", "Wrath", "Starfire", "Innervate").<br>The pre-pull Wrath sentence can stay in the Wrath line if the sentence cap allows. |
| DRUID-R25-3 | CONFIRMED | coach: talent named before the order gives it | `coach/druid-balance.md:23` | "Shift into Moonkin Form and stay there." | Moonkin Form is the 40th-level point on both orders (`addon-builds.json`). The coach reviews any logged boss pull at the player's level (design §6.1). From 10 to 39 the first line of the list names a form the player does not have. This is the shape of PRIEST-R25-14 (Shadowform at 40 to 41). | GUIDE_LINES `druid-balance` "Moonkin Form": "Moonkin Form from forty: shift into it before the pull and stay there." |
| DRUID-R25-4 | CONFIRMED | coach: page's condition dropped | `coach/druid-balance.md:29` | "Faerie Fire." | balance.md:97: "If no feral druid has it on the boss, put yours up. In Moonkin Form it is not free: it costs a global and a little mana, so only do it when no one else will." The Balance card has no Faerie Fire row. A bare entry in "the buttons the guide's rotation names" tells the moonkin to keep it up whoever else does. The generator found no reason because the page names the condition in a bold run-in heading, not a lead clause. | GUIDE_LINES `druid-balance` "Faerie Fire": "Faerie Fire only when no feral druid has it on the boss: in Moonkin Form it costs a global and mana." |
| DRUID-R25-5 | CONFIRMED | coach: list order against the card and page | `coach/druid-feral.md:23 to 31` | Order: Shifting Power, Rip, Shred, Ravage, Pounce, Ferocious Bite, Faerie Fire, Claw, Cower. | The card runs Ravage [1], Faerie Fire [2], Shifting Power [5], Rip [6], Ferocious Bite [7], Shred [8], Claw [9]. feral.md:71 to :75 says the same: opener, then finishers (Rip, then the Bite "for 5 points while Rip still has 5 seconds or more") before "Shred is everything else". The coach puts Shred above Ferocious Bite and Faerie Fire. The "core abilities under-used" finding reads list order ("a button the priority list wants sitting ready while something lower went first"), so a five-point Bite, or a Faerie Fire cast while Shred was ready, reads as a mistake although the card casts it that way. The GUIDE_ORDER comment quotes a feral.md sentence ("sits ahead of every finisher and builder") that the page no longer has. | `starter_overrides.py` GUIDE_ORDER `druid-feral`: ("Ravage", "Faerie Fire", "Shifting Power", "Rip", "Ferocious Bite", "Shred", "Claw", "Pounce", "Cower"). Update the comment to quote feral.md's current rotation. `test_guide_order_statements_win` asserts `priority_names("druid-feral")[0] == "Shifting Power"`, so it needs the same change: Shifting Power ahead of Rip, Bite and Shred. |
| DRUID-R25-6 | CONFIRMED | coach: lines false at 20 to 39 | `coach/druid-feral.md:23, :25, :26` | ":25 Shred fills every other global, building combo points."; ":26 Prowl, walk behind the boss and Ravage once the tank has two or three globals on it."; ":23 Shifting Power every time it comes off cooldown..." | Cat Form and Shred train at 20, but the card's builder at 20 is Claw (row [9]: "Your builder at 20, before the second point of Shredding Attacks"). Shred takes over at 21 (row [8], feral.md:119). Ravage trains at 32 (`trainables`, ranks 32/42/50/58), so from 20 to 31 there is no opener to Prowl for. Shifting Power is the 40th-level point on both orders, and it heads the list. A level 20 to 39 boss pull (Wailing Caverns, Shadowfang Keep, Scarlet Monastery) is logged against all three lines. | GUIDE_LINES `druid-feral`:<br>- "Shred": "Shred fills every other global from twenty-one, building combo points. At twenty, Claw is the builder."<br>- "Ravage": "Ravage from thirty-two: Prowl, walk behind the boss and open once the tank has two or three globals on it."<br>- "Shifting Power": prefix "From forty, ".<br>Keep each sentence within the word cap. |
| DRUID-R25-7 | CONFIRMED | coach: a priority button missing, and list order against the card | `coach/druid-feral-bear.md:23 to 32`; `starter_guide.CONDITION` | List: Dire Bear Form, Enrage, Primal Bite, Lacerate, Faerie Fire, Maul, Demoralizing Roar, Challenging Roar, Barkskin, Frenzied Regeneration. | feral-bear.md:76 to :81 numbers six: Primal Bite, Lacerate, Faerie Fire, Demoralizing Roar, Maul, Swipe. The card agrees: rows [3] to [8]. The coach drops Swipe, the page's sixth item. It also puts Maul above Demoralizing Roar, so a Demoralizing Roar refresh while Maul could be queued reads as "something lower went first". The lead-in says "These are the buttons the guide's rotation names, in the order it gives them", which is false on both counts. Cause of the missing Swipe: `CONDITION` takes "Swipe only with two or more targets" as a group condition. This is also the druid test failure. | `starter_overrides.py`:<br>- GUIDE_LINES `druid-feral-bear`, add "Swipe": "Swipe only with two or more targets in front of you, above forty rage. On a single boss, put the rage into Maul, Lacerate and Primal Bite."<br>- GUIDE_ORDER `druid-feral-bear`: ("Dire Bear Form", "Enrage", "Primal Bite", "Lacerate", "Faerie Fire", "Demoralizing Roar", "Maul", "Swipe").<br>`test_order_follows_the_guide` (Primal Bite < Lacerate < Maul < Barkskin) still passes. The test goes green once Swipe is in GUIDE_LINES. Optional generator change: exclude a target count ("only with two or more targets") from `CONDITION`. |
| DRUID-R25-8 | CONFIRMED | coach: rule against the card at 40 to 53 | `coach/druid-feral-bear.md:28` | "Maul queued on your next swing whenever you have fifteen rage spare." | Card row [7] gate: rage ≥ 15 and (level < 40 or level > 53 or rage ≥ 30). The note and feral-bear.md:80 say the same: "At levels 40 to 53 rage is short: queue Maul only at 30 rage, so Primal Bite seldom waits on rage." At 40 to 53 the coach teaches the queue the card withholds, and the page's reason (keep rage for Primal Bite) is gone. | GUIDE_LINES `druid-feral-bear` "Maul": "Maul queued on your next swing whenever you have fifteen rage spare. From forty to fifty-three, wait for thirty rage, so Primal Bite seldom waits on rage." |
| DRUID-R25-9 | CONFIRMED | coach: form and talent named before the player has them | `coach/druid-feral-bear.md:23, :25` | ":23 Be in Dire Bear Form."; ":25 Then Primal Bite, Lacerate, and Maul queued on your swing." | Dire Bear Form trains at 40, and Bear Form is the form from 10 to 39 (`trainables`: 5487 at 10, 9634 at 40). The Primal Bite entry is the pull sentence (feral-bear.md:72), not priority item 1. It names Primal Bite (the 40th-level point on both orders) and Lacerate (the 25th) as the first presses on every pull. From 10 to 24 the player has neither, and from 25 to 39 no Primal Bite. The page's priority item 1 ("Primal Bite on cooldown, every 6 seconds") is the line the list needs. | GUIDE_LINES `druid-feral-bear`:<br>- "Dire Bear Form": "Be in Bear Form, and in Dire Bear Form from forty."<br>- "Primal Bite": "Primal Bite on cooldown from forty, every six seconds."<br>The "biggest threat hit" clause stays out (routed). |
| DRUID-R25-10 | CONFIRMED | coach: tank drill contradicts the card's refreshes | `coach/druid-feral-bear.md:49` (drill.uptime) | "Watch the buff on you and the debuffs on the boss, not the cooldowns. The moment one drops, press it again." | Card row [4] recasts Lacerate when it has under 3 seconds left; feral-bear.md:77 says the same. If Lacerate drops, its stacks go with it: five casts to rebuild, and Rend and Tear is off meanwhile. Row [6] refreshes Demoralizing Roar "in its last 3 seconds" (:79). So the drill's "the moment one drops" is the move the page avoids. This is the tank text in `starter_kit.DRILLS["tank"]["uptime"]`, a different string from the DPS "recast with 3 seconds left" root, and it applies to every generated tank starter. | Generator/kit: `starter_kit.DRILLS["tank"]["uptime"]`: "Watch the debuffs on the boss and refresh each where the priority list says, before it drops: a stacking bleed that falls off takes its stacks with it." Or a per-spec drill for the bear. |
| DRUID-R25-11 | CONFIRMED | coach: spell named before it is trained | `coach/druid-restoration.md:24` | "Keep Gift of the Wild up on the raid." | Gift of the Wild trains at 50 (`trainables`: 21849 at 50, 21850 at 60). From 10 to 49 a resto druid's group buff is Mark of the Wild, one target at a time. restoration.md:69 is in the level 60 rotation; the coach applies it at every logged level. | GUIDE_LINES `druid-restoration` "Gift of the Wild": "Gift of the Wild up on the raid from fifty, and Mark of the Wild on the group before then." |
| DRUID-R25-12 | CONFIRMED | coach: site, file and data meta talk (shared root, HUNTER-R25-13, PRIEST-R25-16) | `coach/druid-{balance,feral,feral-bear,restoration}.md:17, :21` and the consumables paragraph (:42, :43, :44, :40) | "This is the Balance Druid as the site's guide describes it ... This file says what the guide says and leaves the rest for a later full write."; "These are the buttons the guide's rotation names ... Where the guide or the rotation data gives a reason"; "For this spec the catalogue looks for ..."; bear and feral :17 also carry a site link ("Cat DPS has [its own guide](/guides/druid/feral)"). | Owner rule (guide voice), which the round 25 brief applies to the coach files: no site or sim meta talk. | Generator, the shared root: `persona_starter.who_section` and the `know_section` lead-in, `starter_kit.consumable_paragraph`. Drop the overview's guide link from the persona. |
| DRUID-R25-13 | CONFIRMED | coach: top-rank drill against the rank rows (shared root, HUNTER-R25-7, PRIEST-R25-11) | `coach/druid-balance.md:50` (drill.ranks) | "Visit your trainer every time you level and read your spellbook after: every button on the priority list should show its top rank." | The Balance rows cast below the newest rank at:<br>- Moonfire rank 2 at 20 to 28, rank 3 at 29 to 35, rank 4 at 36 to 39.<br>- Wrath rank 4 at 30 to 39, rank 6 at 46 to 48.<br>- Insect Swarm rank 1 at 30 to 38, rank 3 at 54 to 58.<br>balance.md:137 and :139 teach the same ranks. The drill tells the player to undo them. Feral and bear cast the newest trainer rank at every band, so their drills are true. | Generator: `starter_kit.DRILLS["dps"]["ranks"]`, as PRIEST-R25-11: "...check every button on the priority list shows the rank the guide gives for your level: sometimes that is a lower one." |
| DRUID-R25-14 | CONFIRMED | coach: drill contradicts the card (shared root, HUNTER-R25-8, PRIEST-R25-10) | `coach/druid-balance.md:47`, `coach/druid-feral.md:48` (drill.uptime) | "Recast each with 3 seconds left" | Balance: rows [3] and [4] recast Moonfire and Insect Swarm on `not dotIsActive`, and balance.md:87 says "Recast each one when it falls off". Feral: row [6] recasts Rip under 2 seconds, and feral.md:73 says "A new Rip replaces the old one and throws away the ticks it had left, so let it run to its last tick". | Generator, the shared root: `starter_kit.DRILLS["dps"]["uptime"]`: "Recast each as the priority list says, no earlier, ..." |
| DRUID-R25-15 | CONFIRMED | coach: healer rank drill contradicts the page at 60 (shared root, PRIEST-R25-12) | `coach/druid-restoration.md:47` (drill.ranks) | "Bind a mid rank of your filler heal and use it for every small hole on the tank. Save the top rank for a hole as deep as it heals." | restoration.md:81: "In a raid, keep the top rank on the tank and put rank 7 on a hurt party member." The card agrees at 60: tank Rejuvenation rank 10 ([8]), tank Healing Touch rank 10 ([9]), party Rejuvenation rank 7 ([10] to [14]). The drill moves the tank off the top rank. Leveling has its own ranks by level (:83, :85). | Generator: per-spec healer drill. For Restoration: "Keep the top rank on the tank in a raid, and bind Rejuvenation rank seven for a hurt party member. While leveling, use the ranks the guide gives for your level." |
| DRUID-R25-16 | CONFIRMED | coach: healer behind-pace drill against the page (shared root, PRIEST-R25-13) | `coach/druid-restoration.md:46` (drill.resource) | "If you are behind it, lean on your cheapest heal and save the fast, dear heal for the real emergency." | restoration.md:75: "If your mana is behind, let the HoTs and Swiftmend carry it." The card's Healing Touch filler ([9]) waits for mana ahead of the fight. At 60 the heal cheapest by cost is Rejuvenation rank 1, which :81 says keeps none of your healing power and belongs only on a long five-man with no raid behind you. In a raid it goes on the party at rank 7. A druid has no "fast, dear heal" apart from Nature's Swiftness, which the card keeps for a tank under 35%. | Generator: per-spec `HEALER_RESOURCE_DRILL`. For Restoration: "If you are behind it, let the HoTs on the tank carry the fight, hold Healing Touch for a tank under half, and keep Nature's Swiftness for the save." |
| DRUID-R25-17 | CONFIRMED | coach: consumables kit by resource, not the raid preset (shared root, HUNTER-R25-2, PRIEST-R25-9) | `coach/druid-balance.md:37, :42, :51`; `coach/druid-feral.md:39, :43, :52`; `coach/druid-feral-bear.md:44` | Balance "a combat potion (Major Mana Potion)"; Feral "a flask (Flask of the Titans); food (Smoked Desert Dumplings)" and "Pull preparation: a missing flask and food"; Bear lists no weapon enchant. | The top raid band of each spec in `sim/request/raid-consumes.json`:<br>- Balance (from 55): `major_spellblasting_potion`, no mana potion.<br>- Feral (from 55): no flask at all; `elixir_of_the_mongoose`, `major_frenzy_potion`, `main_hand_imbue:elemental_sharpening_stone`.<br>- Bear (from 50): Flask of the Titans, Mighty Rage Potion and `main_hand_imbue:elemental_sharpening_stone`.<br>So the feral coach marks a cat without Flask of the Titans as unprepared, which nothing on the page or in the preset asks for, and misses the stone and potion. feral.md:97 says "the cat drinks and sharpens like a rogue", and feral-bear.md:99 says "sharpen your weapon". Restoration matches its preset (Distilled Wisdom, Sage's Tea, Major Mana Potion); its oil is omitted, as PRIEST-R25-S5. | Generator, the shared root: `starter_kit.kit()` picks the flask by `Traits.style`, the stone by `Traits.stone` and the potion by `Traits.resource`. Read all three from the spec's top-band preset, the way `starter_food.preset_food` reads the food. |

Counts: CONFIRMED 17, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

- Guide pages and card notes: 1 (R25-1, a side effect of round 24's two Eclipse fixes).
- Coach: 16. Regenerating fixes none.
  - Ten need `starter_overrides.py`:
    - Balance: R25-2 (Wrath and Starfire), R25-3 (Moonkin Form) and R25-4 (Faerie Fire).
    - Feral: R25-5 (order) and R25-6 (levels).
    - Bear: R25-7 (Swipe and order; this is the druid test failure), R25-8 (the Maul gate) and R25-9 (Dire Bear Form, Primal Bite).
    - Restoration: R25-11 (Gift of the Wild).
  - Seven need a generator or kit change:
    - R25-10, the tank uptime drill: a new root, shared with every generated tank starter.
    - R25-12 to R25-17, the shared roots the brief names, each filed once for druid: meta talk, the top-rank drill, the recast-at-3-seconds drill, the healer mid-rank filler drill, the "cheapest heal" drill, and the combat potion by resource (with the feral flask and the missing stones).
- Shared roots that do not apply to druid:
  - Leveling-only "no X" clauses: the one druid exclusion, Rake, comes from the level 60 rotation sentence and is right.
  - Short spell names not matched, and full names missing where the page uses a short form: every druid page names its buttons in full.

## Shoulds (wording, not findings)

- **DRUID-R25-S1, Balance coach :25, :26 and Feral coach :24, :29.** "When Insect Swarm is not ticking on the target." and the like are card-note fragments with no verb. They are true, but they read oddly as list items.
- **DRUID-R25-S2, Balance coach :30, "Hurricane has no cooldown in Forever."** True (16914 to 17402: cooldown 0). It is a fact, not an instruction, and the card has no AoE row. Optional: "Hurricane on a pack that lives through the channel; it has no cooldown in Forever."
- **DRUID-R25-S3, Feral coach :27 Pounce** sits in the boss priority list while its own line says it is for trash. A Pounce on a boss is then not "outside the priority list". Optional: leave it out of the list.
- **DRUID-R25-S4, Bear coach :31 and :32.** Barkskin and Frenzied Regeneration sit at the bottom. The card has them at rows [1] and [2]. Both are off the global, and `test_order_follows_the_guide` pins Barkskin below Maul on purpose, so this is a Should.
- **DRUID-R25-S5, Restoration coach.** The list omits the party Rejuvenation (restoration.md:79, card rows [10] to [14]) and gives Innervate no use ("Innervate is yours from level forty"). The card uses it below 30% mana. Omissions, not false statements.
- **DRUID-R25-S6, override comments quote sentences the pages no longer have.**
  - `GUIDE_LINES["druid-restoration"]["Swiftmend"]` quotes "and your mana is keeping pace"; restoration.md's Swiftmend line no longer has it.
  - The `GUIDE_ORDER["druid-feral"]` comment quotes "sits ahead of every finisher and builder".
  - The coach Swiftmend line matches the card's mana gate. The page's missing gate is the carried restoration :73 Should.
- **DRUID-R25-S7, coach Voice on the three non-healers.** The voice section carries the healer clauses (PRIEST-R25-S4, `tools/coach/voice.md`).
- **DRUID-R25-S8, the balance page names no opening potion** while the preset opens with Major Spellblasting Potion (shadow.md names it for Shadow). This is an omission. R25-17 covers the coach side.
- **DRUID-R25-S9, `coach/druid-balance.json` `auras`** lists Insect Swarm only, not Moonfire. It is not player-facing.

Carried, unchanged and optional: every Should in round 24's list (none applied), including restoration :73 mana gate, balance :139 Starfire "for the opener" before 35 against the card's mana-ahead Starfire at 20 to 34 (leveling prose on questing, as before), feral-bear :83 "Swipe moves to the top on a pack", feral-bear :61 Feral Instinct.

## Checked and right this round

- **Page lines unchanged since round 24.** Every value round 24 re-derived still holds on identical data:
  - Eclipse 0.17 / 0.33 / 0.50 s.
  - Improved Wrath 0.5 s and 50%; Moonglow 25%.
  - Feral Instinct 10 / 20 / 30% (bear :115, "at 18 and 19 adds 20%", is two points).
  - Feral Swiftness 4%; Ferocity 5.
  - Enrage 10 rage plus 20 over 10 s ("2 a second").
  - Barkskin 20%, 15 s, 60 s; Frenzied Regeneration 1% health a rage point, 10 a second, 10 s.
  - Primal Bite 6 s; Swipe 3 targets; Hurricane cooldown 0.
  - Pounce and Ravage "Must be prowling and behind the target".
  - Tranquility 5 min, Improved Tranquility -60%; Shifting Power 55% of base mana.
  - Gift of the Wild 50.
- **Coach lines OK at every level they apply to** (one line each in the Coach sweep):
  - Bear: Enrage, Lacerate, Faerie Fire, Demoralizing Roar, Challenging Roar, Barkskin, Frenzied Regeneration.
  - Feral: Rip, Pounce, Ferocious Bite, Faerie Fire, Claw, Cower.
  - Balance: Insect Swarm, Moonfire, Innervate.
  - Restoration: Rejuvenation, Regrowth, Swiftmend, Healing Touch, Nature's Swiftness, Tranquility, Innervate.
- **Level 60 book ranks** (Starfire 7 25298, Healing Touch 11 25297, Rejuvenation 11 25299, Ferocious Bite 31018) are still not cast by the cards. This is round 2's "book" ranks and sim round 18 R5, pending owner. Not re-filed. balance.md:139's "cast the newest rank" and restoration.md:81's "top rank" read as the trainer ranks, as decided before.

## Decisions

- **D1. Coach lines are judged at every level a boss pull can be logged** (PRIEST-R25 D1; design §6.1 reads the level from `AdvLevel`). That is why R25-3, R25-6, R25-9 and R25-11 are findings, although each line is true at 60. A line that only names a button the player lacks, with nothing false in it (Barkskin before 44, Innervate before 40), is not filed. A line is filed when its instruction or its order is wrong at that level: "be in", "shift into and stay", the pull sequence, the builder.
- **D2. R25-1 is CONFIRMED although the note is generic.** The clause speaks of every one-point band. On the leveling order that starts at 35, where the page, rewritten in the same round, says the loop does not pay yet and the measurement shows a tie.
- **D3. R25-7's fix is an override, not the generator,** because GUIDE_LINES already re-adds a button the generator left out (`with_lines`). The `CONDITION` regex change is offered as optional. It would also touch every other spec's "only with" reading.
- **D4. R25-10 is a new shared root,** not the DPS "recast with 3 seconds left" root: it is the tank string. It is filed once here, and it affects the Protection paladin starter too.
- **D5. R25-17 bundles the feral flask and the missing stones with the potion root.** All three come from `kit()` choosing by trait instead of the preset, so they are one fix.

## Sims requested

None new. Carried and routed: the Balance raid order Eclipse loop on and off at bare 39, 40 and 41 (it also decides the raid-order half of R25-1's fixed note); the bear pack fight with Swipe at the top.

## Engine questions

None new. Carried and held: `BloomSpellCoeff` 0.2143 against client 0; Eclipse on a resisted Wrath; Maul, Swipe and Lacerate threat multipliers; Rend and Tear on others' bleeds; level 60 book ranks (pending owner).

## Sim findings

None new.

## Routed items, not re-filed

These are as in round 24's list:

- Row-before-talent rendering: Shifting Power 10 to 39, Insect Swarm 20 to 26, Nature's Swiftness 10 to 29, Swiftmend 10 to 42, Balance Eclipse Wrath 10 to 34, the Feral "Below level 20" rows past 20.
- "Biggest threat hit" and the bear threat model.
- `balance.md:105`, a resisted Wrath and Eclipse.

## Sources

- **Client 1.60.1.70790** (`data/builds/1.60.1.70790/`):
  - `addon-builds.json`, all eight druid orders decoded against `talents/druid.json`.
  - `addon-data.json`, `rotations/druid-*`, every band, with `race_spell_ids`.
  - `trainables/druid.json`: Cat Form, Dire Bear Form, Moonkin Form, Innervate, Gift of the Wild, Shred, Claw, Ravage, Pounce, Barkskin, Frenzied Regeneration, Challenging Roar, Enrage, Starfire, Hurricane, Tranquility, Healing Touch, Rejuvenation, Regrowth.
  - `spelltips/druid.json`: 5229, 9005, 9827, 9867, 22812, 22842, 29166, 1322605, 5209, 740, 9863, 21849, 16914, 17402, 1238070, 1235827, 9881, 9908, 24858, 22829, 31018, 25298; all Wrath, Starfire, Moonfire and Insect Swarm ranks.
  - `talents/druid.json`: Eclipse, Improved Wrath, Moonglow, Ferocity, Furor, Feral Swiftness, Feral Instinct, Shredding Attacks, Natural Reaction, Improved Tranquility.
- **Curated:** `data/curated/apl/druid-*.json`, `data/curated/coach/druid-*.md` and `.json`, `sim/request/raid-consumes.json` (druid top bands).
- **Generator:** `tools/coach/persona_starter.py`, `starter_overrides.py` (GUIDE_LINES and GUIDE_ORDER for the four druid specs), `starter_guide.py` (`CONDITION`, `excluded_abilities`), `starter_kit.py` (`kit`, `DRILLS`, `HEALER_RESOURCE_DRILL`), `test_persona_starter.py` (4 red; the druid one is `test_every_correction_restates_its_guide`). Read-only regen output: `scratchpad/gr25/druid-coach/` (identical to the committed files).
- **Coach design:** `docs/superpowers/specs/2026-10-27-log-coach-design.md` §6.1 and F5.
- **Reviews:** round 24 druid, the ledger's Round 24 section, round 25 hunter, mage and priest; `2026-11-05-sim-audit-round21/druid.md` (S1); `2026-11-04-sim-audit-round18/druid.md` (R5, book ranks); round 2 druid (book ranks).
- **No web research needed:** the client data, the cards, the pages and the preset decided every call.

## Shipped orders used

Unchanged from round 23 and round 24:

- **Balance leveling:** Improved Wrath 10 to 14, Genesis 15 and 16, Moonglow 17 to 19, Improved Moonfire 20 and 21, Insect Swarm 27, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonkin Form 40.
- **Balance raid:** as leveling to 32, then Improved Starfire 33 to 37, Nature's Grace 38, Eclipse 39, 41 and 42, Moonkin Form 40.
- **Feral (both orders):** Ferocity 10 to 14, Shredding Attacks 20 to 22, Predatory Discipline 31 (leveling) or 30 (raid), Rend and Tear from 37 (leveling) or 36 (raid), Shifting Power 40.
- **Bear (both orders):** Ferocity 10 to 14, Furor 15, Feral Swiftness 16 and 17, Feral Instinct 18, 19 and 22, Lacerate 25, Primal Bite 40, Rend and Tear 41 to 45.
- **Restoration (both orders):** Bloom 20, Nature's Swiftness 30, Swiftmend 43, Improved Tranquility 47 and 48.

## Appendix A: Note sweep

The card rows in `addon-data.json` are those of round 24's tree with round 24's three note edits. Each line gives: spec, row and action, the bands it renders on, each named talent, spell, rank or value with the level the player has it (leveling / raid where they differ), and the verdict.

**Balance**

- **[1] Moonfire, end of fight. Bands 10 to 60.** Wrath's cast time; Moonfire is the best instant: Insect Swarm has no hit, Faerie Fire no damage, Hurricane is channelled. OK.
- **[2] Wrath, end of fight. Bands 10 to 60.** Starfire from 20 (scoped); before 20, the last 3.5 s (gate verified). OK.
- **[3] Moonfire upkeep. Bands 10 to 60.** Improved Moonfire 20 and 21 and Nature's Splendor 26 are conditional. OK (carried Should).
- **[4] Insect Swarm. Bands 20 to 60 (talent 27, routed row).** One global. OK.
- **[5] Eclipse Wrath. Bands 10 to 60.**
  - Eclipse 35 / 39 is scoped; two charges, 0.17 / 0.50 s (talent text); "do not pay" holds at every one-point band.
  - "The loop still wins" at 35 to 39: **R25-1**.
- **[6] Wrath filler and mana line. Bands 10 to 60.** Starfire 20 (gate verified); Innervate 40; Improved Wrath 10 to 14; "a fraction of a Starfire's mana for half the cast" holds from 20 (rank 3 at 20 against Starfire rank 1, both under Moonglow). OK (band 10 Should carried).
- **[7] Starfire. Bands 20 to 60.** Eclipse 35 / 39 scoped. OK.

**Feral**

- **[1] Ravage. Bands 32 to 60.** Prowl 20, Poised 32, fires once. OK.
- **[2] Faerie Fire. Bands 18 to 60.** Cat Form 20; free, 6 s. OK (band 18 Should carried).
- **[3] Moonfire, below 20. Bands 10 to 60 (routed row).** OK.
- **[4] Wrath, below 20. Bands 10 to 60 (routed row).** OK.
- **[5] Shifting Power. Bands 10 to 60 (talent 40, routed row).** 40 energy, mana, 8 s. OK.
- **[6] Rip. Bands 20 to 60.** 5 points, under 2 s; Rend and Tear 37 / 36 scoped. OK.
- **[7] Ferocious Bite. Bands 24 to 60.** Predatory Discipline 31 / 30 is conditional; Rip 5 s. OK.
- **[8] Shred. Bands 21 to 60.** From 21; Shredding Attacks' second point at 21 on both orders; "only at 20"; "harder at every rank" re-derived in round 24 on the same data. OK.
- **[9] Claw. Band 20 only.** 40 against 54 at 20 (Ferocity 5, Shredding Attacks 1). OK.

**Feral bear**

- **[0] Cooldowns. Page only.** Mighty Rage Potion in form. OK.
- **[1] Barkskin. Bands 44 to 60.** Trainer 44; 20%, 15 s, 60 s. OK.
- **[2] Frenzied Regeneration. Bands 36 to 60.** Trainer 36; half health, 20 rage. OK.
- **[3] Primal Bite. Bands 40 to 60.** Talent 40; 6 s; "biggest threat hit" is routed. Not re-filed.
- **[4] Lacerate. Bands 25 to 60.** Talent 25; five stacks; 3 s; Rend and Tear "from 41" (both orders). OK.
- **[5] Faerie Fire. Bands 18 to 60.** Bear Form from 10; free, 6 s; threat, ranged pull. OK (band 18 Should carried).
- **[6] Demoralizing Roar. Bands 10 to 60.** "204 at 60" scoped; last 3 s; does not stack with Demoralizing Shout. OK (carried Should).
- **[7] Maul. Bands 10 to 60.** 15 rage; 30 at 40 to 53 (gate verified); Ferocity 10 by 14; Lacerate from 25 scoped; "nearly every time" (R24-3 fix). OK.
- **[8] Swipe. Bands 16 to 60.** Trainer 16; two or more targets and 40 rage (gate verified); up to three (tooltip); "Lacerate from 25 and Primal Bite from 40" scoped (R24-4 fix). OK.

**Restoration**

- **[1] Innervate. Bands 40 to 60.** Trainer 40; under 30%. OK.
- **[2] Read Leyline. Alliance Skyborne, no addon row.** 50%, 15 s; not under Innervate. OK.
- **[3] Nature's Swiftness. Bands 10 to 60 (talent 30, routed row).** Tank under 35%. OK.
- **[4] Healing Touch save. Bands 10 to 60.** The newest trainer rank at every band (rank 11 is a book). OK.
- **[5] Swiftmend. Bands 10 to 60 (talent 43, routed row).** Under 75%, mana pace; keeps the HoT; 15 s; 20% of base mana. OK.
- **[6] Tranquility. Bands 30 to 60.** Trainer 30; two under 70%; every 2 s. OK.
- **[7] Regrowth, tank. Bands 12 to 60.** Top rank at every band; Bloom from 20 scoped. OK.
- **[8] Rejuvenation, tank. Bands 10 to 60.**
  - Rank 2 from 20 to 39, with a Night Elf on rank 3 at 29 only (band 29: base 2091 overridden to 1430 for Night Elf, 1058 for Tauren and both Skyborne).
  - Ranks 4 / 5 / 6 / 8 at 40 to 43 / 44 to 48 / 49 to 56 / 57 to 59.
  - Rank 10 at 60.
  - All match the rows. OK.
- **[9] Healing Touch filler. Bands 10 to 60.**
  - Rank 2 at 20, with a Tauren keeping it to 23 (band 21: base 5188, Tauren 5186).
  - Ranks 3 / 4 / 5 / 6 at 30 to 33 / 34 to 39 / 40 to 45 / 46 to 49.
  - The top rank from 50.
  - All match the rows. OK.
- **[10] Rejuvenation, party. Bands 10 to 60.** Under 70%; no mana-pace gate. OK.
- **[11] to [14].** No note.

## Appendix B: Coach sweep

One line per coach line, judged at every level a boss pull can be logged.

**druid-balance.md**

- :17 overview. Meta talk: R25-12. The identity sentence is the page's overview. OK.
- :23 Moonkin Form: R25-3.
- :24 Wrath (pre-pull only): R25-2.
- :25 Insect Swarm. Talent 27; the fragment is true. OK (S1).
- :26 Moonfire. OK (S1).
- :27 Starfire, "spending the charges": R25-2.
- :28 Innervate. Trainer 40; matches balance.md:95. OK.
- :29 Faerie Fire, bare: R25-4.
- :30 Hurricane. Cooldown 0 (client). OK (S2).
- :33 to :40 findings list. Generic. OK.
- :42 consumables. Potion: R25-17. Flask, food and oil match the preset. Meta talk: R25-12.
- :46 core_casts. OK. :48 idle. OK. :49 resource. OK. :53 adherence. OK.
- :47 uptime: R25-14.
- :50 ranks: R25-13.
- :51 preparation: R25-17.
- :52 death. OK.

**druid-feral.md**

- :17 overview. Site link and meta talk: R25-12.
- :23 Shifting Power: R25-5 (order) and R25-6 (40).
- :24 Rip. OK.
- :25 Shred: R25-5 and R25-6.
- :26 Ravage: R25-6.
- :27 Pounce. The tooltip agrees. OK (S3).
- :28 Ferocious Bite: R25-5 (order). The text is OK.
- :29 Faerie Fire: R25-5 (order). The text is OK.
- :30 Claw. OK.
- :31 Cower. Bare, matches feral.md:85. OK.
- :39 and :43 flask, no stone or potion: R25-17.
- :47 core_casts. OK.
- :48 uptime: R25-14.
- :49 idle. OK.
- :50 resource. OK.
- :51 ranks. True: feral casts the newest trainer rank at every band.
- :52 preparation: R25-17.
- :53 death. OK. :54 adherence. OK.

**druid-feral-bear.md**

- :17 overview. Site link and meta talk: R25-12.
- :23 Dire Bear Form: R25-9.
- :24 Enrage. Trainer 12; 10 plus 20 over 10 s. OK.
- :25 Primal Bite entry: R25-9.
- :26 Lacerate. OK from 25.
- :27 Faerie Fire. OK from 18.
- :28 Maul: R25-8; its order against Demoralizing Roar is R25-7.
- :29 Demoralizing Roar. OK.
- Swipe missing: R25-7.
- :30 Challenging Roar. Trainer 28. OK.
- :31 Barkskin. 20%, 15 s, 60 s; trainer 44. OK (S4).
- :32 Frenzied Regeneration. 1% a rage point, 10 a second, 10 s; trainer 36. OK (S4).
- :44 consumables. No stone: R25-17.
- :48 core_casts. OK.
- :49 uptime: R25-10.
- :50 idle. OK.
- :51 resource. OK (the Maul gate keeps rage back; "never let a button wait" is the aim).
- :52 ranks. True: the bear casts the newest rank at every band.
- :53 preparation. OK apart from the stone (R25-17).
- :54 death. Frenzied Regeneration under half. OK.
- :55 adherence. OK.

**druid-restoration.md**

- :17 overview. Meta talk: R25-12.
- :23 Rejuvenation pre-pull. OK.
- :24 Gift of the Wild: R25-11.
- :25 Regrowth and Rejuvenation. OK.
- :26 Swiftmend. Matches the card (S6).
- :27 Healing Touch. "Under about half" against the card's 55%. OK.
- :28 Nature's Swiftness. Talent 30. OK.
- :29 Tranquility. Trainer 30; Improved Tranquility at 47 and 48 is conditional. OK.
- :30 Innervate. Trainer 40. OK (S5).
- :33 to :38 findings list. OK.
- :40 consumables. Matches the preset. Meta talk: R25-12.
- :44 core_casts. OK.
- :45 uptime. OK.
- :46 resource: R25-16.
- :47 ranks: R25-15.
- :48 preparation. OK.
- :49 death. OK.

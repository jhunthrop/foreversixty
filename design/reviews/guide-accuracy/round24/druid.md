# Guide accuracy round 24: Druid

Audited on the round 24 audit tree (`scratchpad/gr24/audit`, head ea9adea1c: main after guide round 23 landed, which merged sim round 21 at 617d4e9c4). Engine pin 170a59c39 (`sim/enginever/version.go`), source read from `scratchpad/gr24/engine`. Active client 1.60.1.70790.

Files: `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`, every row note in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered band of `addon-data.json` `rotations/druid-*`, against both shipped orders in `addon-builds.json` (unchanged since round 23), `trainables/druid.json`, `spelltips/druid.json`, `spellconst/druid.json`, `talents/druid.json`, `items/druid.json` and `bis/druid-*.json` (unchanged since round 23).

Read first: the round 1, 2 and 24 briefs, `round23/druid.md`, the ledger's Round 19 to 23 sections, `design/reviews/2026-11-05-sim-audit-round21/druid.md` and `integration.md`.

## New work: sim round 21 (commits b2ddafb03, 0dc6a8fbf, 1df6ca3a3, 96b829d1f, 4392cb7bf, bf29c62a7, b8cb907ec, fbde7096e, 0a0cdbf2e)

`git diff 24ee457d0 HEAD` on the druid files: Balance card rows [5] and [6] (note and gate), Feral rows [8] and [9] (note and gate) and the Claw `expected_idle` reason, bear rows [7] and [8] (note and gate); balance.md :127, :139, :141; feral.md :119; feral-bear.md :35, :80, :81 (and round 23's :107); restoration.md :91 (round 23). Each is checked below and in the sweep.

- **Balance [6] gate.** `not spellIsKnown(9876 rank 6)` reads "no Starfire known": `sim/request/rotation_ranks.go` rewrites a ranked id to the highest learned rank and nils it when none is learned. So the filler line fires exactly below 20, as the note says. OK.
- **Balance :127**, Wrath filler before 20 from 4: matches rows [3] and [6] at bands 10 to 16. OK.
- **Balance :139 and :141 and row [5]**: DRUID-R24-1 and DRUID-R24-2.
- **Feral [8], [9], :119**, Shred from 21: Shredding Attacks takes 20, 21, 22 on both orders; Shred 60 less 6 a point (54 at 20, 48 at 21); Claw 45 less Ferocity's 5 (40, full by 14). Rendered: Claw only at band 20, Shred from 21. Per energy, Claw beats Shred at 20 until a normal hit of about 131 (far above a level 20 cat); at 21 the break-even is about 32, where the lane measured Shred +1.0%. "Shred hits harder at every rank" re-derived at bands 21 to 60. OK.
- **Bear :35**, "Bear Form adds 115% to all the threat you generate": Blizzard's 3 November note as the sim lane read it; the client tooltips for 5487 and 9634 state no threat figure, so nothing on the client contradicts it. OK.
- **Bear :80 and [7]**, Maul at 30 rage from 40 to 53: gate matches (`currentLevel` < 40 or > 53 or rage >= 30). The reason it gives is DRUID-R24-3.
- **Bear :81 and [8]**, Swipe on two or more targets: gate matches. "Each hit grows with your attack power" is the 1 October note (3% of attack power); client rows carry 0 and the tooltip is silent, so no contradiction. The note's later clause is DRUID-R24-4.
- **Class index page.** It teaches no rotation, names no Swipe, Shred or Wrath rule, and nothing on it was left behind by sim round 21. The Furor, Shifting Power, Faerie Fire, paw swing (1.0 and 2.5 s) and Swiftmend lines agree with the spec pages. OK.

## Round 23's fixes, rechecked

- **feral-bear.md:107 "No craft in your raid set binds on pickup".** Holds. Band 60 raid, both factions: the crafts are Outlaw's Collar, Warbear Harness, Hivethrasher's Girdle (Leatherworking) and Fel Cape (Tailoring), all `equip` in `items/druid.json`; the ranged slot is empty (Idol of Swiftness only as an alternative). BiS file unchanged since round 23.
- **restoration.md:91 "every heal except Bloom and a rank far below your level".** Holds, and agrees with :81 (rank 1 keeps none at 60, rank 7 most). No other druid heal has a zero healing coefficient in `spellconst/druid.json`; Swiftmend's heal is the HoT's total, which carries healing power.

## Findings

### DRUID-R24-1: Balance card [5] and balance.md:141, the one-point Eclipse reason (CONFIRMED, kind: reason, cast-time arithmetic)

- **Card [5], second sentence (renders 10-60; Eclipse 35 leveling, 39 raid):** "Each Wrath banks two charges, each taking time off one Starfire: 0.17 seconds a charge with one point, half a second with all three, and even one point makes one Wrath and two quicker Starfires beat three plain ones."
- **balance.md:141:** "The single Eclipse point you level with takes only 0.17 seconds off a Starfire per charge, and that is still enough for one Wrath and two quicker Starfires to beat three plain ones."
- **Evidence (client 70790, `spellconst/druid.json` centres, `talents/druid.json`):** leveling order at 40, bare: Wrath rank 6 (6780, rendered) 54 + 0.9 x 2 = 55.8 at 1.5 s (Improved Wrath 5/5); Starfire rank 3 (8950, rendered 34 to 41) 201 + 2.2 x 6 = 214.2 at 3.4 s (Improved Starfire 1/5 at 33). One Wrath and two Starfires at 3.23 s: 484 damage in 7.96 s = 60.8 a second. Three plain Starfires: 643 in 10.2 s = 63.0 a second. The plain Starfires win by about 3.5% per second of casting. Break-even needs Wrath at 34% of a Starfire; rank 6 against rank 3 is 26%. At leveling 60 (Wrath 97 + 0.571 SP against Starfire rank 6, 477 + SP) the loop needs about 290 spell power to pay on cast time. On the raid order at 39 to 40 (one point, Starfire 3.0 s) the bar is 39%, at 41 (two points) 28%: both fail at bare. Raid 60 with all three points (bar 17%, ratio about 41% in the raid set) holds, which is what the old note said.
- **Why the measurement still shows a gain:** the sim lane's +0.9 to +1.0% at bare 40 to 60 (S1) comes from mana, not cast time. A Wrath at half cost (40 for rank 6) does 1.4 damage a mana point against Starfire rank 3's 1.2 under Moonglow, and from 20 the card's mana line shows the moonkin is mana-bound. So the fact (the loop wins from 40 on the leveling order) is measured; the reason both lines give (the 0.17 s pays for the Wrath) is false at leveling spell power, and the card note also states it for the raid order's one-point bands 39 to 41, which nothing measured.
- **Fix, note:** "...0.17 seconds a charge with one point, half a second with all three. With one point the quicker Starfires do not pay for the Wrath's cast; the loop still wins because the Wrath costs a fraction of a Starfire's mana."
- **Fix, :141:** "The single Eclipse point you level with takes only 0.17 seconds off a Starfire per charge, too little to pay for the Wrath on its own; the loop still wins because the Wrath is cheap and leaves more mana for Starfires."

### DRUID-R24-2: balance.md:139 and :141 against card [5] and [7] at 35 to 39 (CONFIRMED, kind: page contradicts its card)

- **Lines:** :139 "Until 40, Starfire is for the opener on a mob that is not yet on you." :141 "Eclipse is yours at 35, but wait for Moonkin Form and Innervate at 40 before you run the raid loop on long fights."
- **Card:** row [5] note "From 35 on the leveling order and 39 on the raid order, whenever no Eclipse charge is banked", gated only by `auraIsKnown(408248)` (no level, Innervate or fight-length condition); row [7] "Once you have Eclipse (35 leveling, 39 raid order) it spends the charges Wrath banked." Both render on bands 35 to 39 of the leveling order.
- **Why it counts:** at 35 to 39 the card tells the player to run the loop and the prose tells them to wait; "on long fights" is a condition the card does not have. The lane measured the two lines identical at 36 and 39 (the mana line casts Wrath nearly every global there), so the instruction "wait" is harmless in damage, but the page states two opposite instructions at the same levels.
- **Fix, :141 first sentence:** "Eclipse is yours at 35, and the card runs the loop from then: one Wrath whenever no Eclipse charge is banked, then Starfires. Until Moonkin Form and Innervate at 40 your mana keeps you on Wrath most of the time anyway, so the loop starts to pay at 40." **:139:** "Until 35, Starfire is for the opener on a mob that is not yet on you."

### DRUID-R24-3: feral-bear.md:80 and bear card [7], "never waits" (CONFIRMED, kind: absolute claim against the mechanism)

- **:80:** "queue Maul only at 30 rage, so Primal Bite never waits on rage when it comes off cooldown." **[7] (bands 10-60, clause about 40-53):** "the 30 makes sure Primal Bite fires every time it comes off cooldown."
- **Evidence:** costs from `spelltips/druid.json` less the bear build's discounts: Maul 15 - 5 = 10, Primal Bite 20 - 5 = 15, Lacerate 15 - 3 = 12, Demoralizing Roar 10. Engine (`sim/druid/maul.go`): queuing Maul checks the rage on hand, and the rage is spent when the swing lands; the Maul swing earns no rage (`core/rage.go`). Card order puts Lacerate and Demoralizing Roar above Maul, so at 30 rage: queue Maul, Lacerate (18 left), Maul lands (8 left), and a Primal Bite coming off cooldown waits. The lane's own measure agrees: 29.45 Bites of a possible 30 at 43 with the gate (28.35 without).
- **Fix, :80:** "...so Primal Bite seldom waits on rage when it comes off cooldown." Note: "...and the 30 keeps rage for Primal Bite nearly every time it comes off cooldown."

### DRUID-R24-4: bear card [8] Swipe, later clause names Lacerate and Primal Bite before they are trained (CONFIRMED, kind: spell named before trained; comparison at a measured tie)

- **Note (renders 16-60):** "...on a single boss the same rage does more in Maul, Lacerate and Primal Bite."
- **Evidence:** Lacerate is the bear order's point at 25, Primal Bite at 40 (both orders the same). Bands 16 to 24 render the row with neither, bands 25 to 39 without Primal Bite. The lane measured single-target Swipe off against on as "a tie at 20", so "does more" is not true at 20 either.
- **Fix:** "...on a single boss the rage is better spent on Maul, and on Lacerate from 25 and Primal Bite from 40." (feral-bear.md:81 sits in the level 60 rotation section and can stay.)

Counts: CONFIRMED 4, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. All four are sim round 21 lines. R24-2 and R24-1 tell a leveling moonkin two opposite things and a wrong reason at 35 to 41; R24-3 and R24-4 are bear note wording. None moves a button, gate or rank on a card.

## Shoulds (wording, not findings)

- **New: feral-bear.md:83** "Swipe moves to the top on a pack" against card [8], which now gates Swipe on two or more targets yet keeps it last, under Maul and Lacerate. The card is single-target and the pack order is unmeasured (Sims requested). Bear [8] "on a pack of three it is your main threat" is unmeasured too.
- **New: feral-bear.md:61** Feral Instinct's "more threat from Swipe" does nothing on a single boss now that the card never Swipes one target (lane decision B5 keeps the points); the line is true for packs.
- **New: Feral Claw `expected_idle` reason** says "Claw is the builder only at level 20": fine; the lane's open Should (the Claw card names 54 where the page gives 48) is already satisfied (the page names 48 for 21 and the card 54 for 20).
- **Carried, none applied:** restoration :73 mana gate; Feral [8] "at every rank" at 28-29 on a weak weapon; Balance [6] and :95 "half the cast"; balance :105 "every cast"; balance :91 "always lands on a Starfire or a DoT"; feral-bear :76 and bear [3] "biggest threat hit" (routed); restoration :121 tauren 21 to 23; Balance [6] "With Improved Wrath" at band 10; restoration :85 "Below 20, from 21"; feral :67 "The raid build skips Furor"; Feral Bite and Balance Moonfire conditionals; Bear Demoralizing Roar "204 ... at 60" on rank 1 bands; Feral and bear Faerie Fire "free" on band 18; "for good" at balance :115 and restoration :101; restoration :57 "next Nature spell"; feral :77 Wolfshead.

## Superlative, comparison and quoted value sweep (new and changed lines)

- Balance :127 "Never stand idle ... with mana to spare": row [6] fires unconditionally below 20. OK.
- Balance [6] "a fraction of a Starfire's mana for half the cast": as round 23 (band 20: 40 at half against 95). OK.
- Balance [5] / :141 "0.17 seconds a charge", "half a second with all three": Eclipse ranks 0.17/0.33/0.50 (talents). OK. "Beat": R24-1.
- Feral [8] "only at 20", "at every rank"; [9] "40 to Shred's 54": OK as above.
- Bear :35 "115% to all the threat": Blizzard note, form-wide. OK.
- Bear [7] "10 rage by level 14": Ferocity 10-14. "Lacerate from level 25": OK. "every time": R24-3.
- Bear [8] / :81 "up to three": Swipe tooltip "3 nearby enemies". "Two or more": gate. "Does more": R24-4.
- Restoration :91 "every heal except": OK (round 23 fix).

## Decisions

- **D1. R24-1 is CONFIRMED although the loop's gain is measured.** Round 23's lens is the reason each line gives. Both lines say the 0.17 s is what pays; the client values say it does not below raid spell power, and the gain the lane measured has a different cause.
- **D2. R24-2 is CONFIRMED although the lane measured no damage difference at 36 and 39.** The page and its card give opposite instructions on the same bands; that is the "contradicts the page" bar from round 8.
- **D3. R24-3 is CONFIRMED on the mechanism.** The sequence Maul queued, Lacerate, Maul lands leaves 8 rage with the card's own costs and order; the lane's 29.45 of 30 is consistent with it. Not routed as a sim: the wording, not the gate, is wrong.
- **D4. Bear Form 115% stays.** A Blizzard patch note for a server-side value the cached table predates, with the warrior Defensive Stance precedent visible in Wowhead's Forever tooltips (sim lane B1). The client tooltip states no figure.
- **D5. Swipe's attack power share stays.** Blizzard's 1 October note; the client tooltip states only the flat amount, which the page does not contradict.

## Sims requested

- **Bear pack fight, three targets, bare 30, 50 and raid 60:** Swipe at its current row against Swipe first (above Primal Bite) on two or more targets. Decides feral-bear.md:83 "Swipe moves to the top on a pack" and the [8] "main threat" clause.
- **Balance raid order, bare 39, 40 and 41:** Eclipse Wrath row on against off (one and two points, Starfire 3.0 s). Decides whether the [5] note's raid-order clause can keep "the loop still wins" after the R24-1 fix.
- Optional: bear at 43 with the Maul 30 gate, Primal Bite casts delayed by rage against delayed by the global, to word R24-3's fix ("seldom" against "nearly every time").

## Sim findings

None new.

## Engine questions

None new. Carried and held: `BloomSpellCoeff` 0.2143 against client 0; Eclipse on a resisted Wrath; Maul, Swipe and Lacerate threat multipliers (sim round 21 B4); Rend and Tear on others' bleeds (C2).

## Routed items, not re-filed

Row-before-talent rendering (Shifting Power 10-39, Insect Swarm 20-26, Nature's Swiftness 10-29, Swiftmend 10-42, Balance Eclipse Wrath 10-34, Feral "Below level 20" rows past 20); "biggest threat hit" and the bear threat model; `balance.md:105` resisted Wrath and Eclipse (launch log).

## Sources

- **Client 1.60.1.70790** (`data/builds/1.60.1.70790/`): `spellconst/druid.json` (5179, 5180, 6780, 8905, 9912, 2912, 8949, 8950, 8951, 9875, 9876, 25298), `spelltips/druid.json` (Bear Form 5487, Dire Bear Form 9634, Maul, Swipe, Lacerate, Primal Bite, Claw, Shred, Demoralizing Roar), `talents/druid.json` (Eclipse 104935, Improved Starfire 104933, Feral Instinct 104940), `addon-builds.json`, `addon-data.json` (all rendered druid bands), `bis/druid-feral-bear.json`, `items/druid.json`.
- **Curated:** `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`.
- **Engine 170a59c39:** `sim/druid/talents.go` (applyEclipse), `sim/druid/wrath.go`, `sim/druid/maul.go`, `sim/core/apl_values_spell.go`; site `sim/request/rotation_ranks.go` (ranked id rewrite in conditions).
- **Reports:** `2026-11-05-sim-audit-round21/druid.md` (S1, S2, C1, B1 to B7) and `integration.md`; `round23/druid.md`; `ledger.md` Rounds 19 to 23.
- **Web (as read and cited by the sim lane, not re-fetched this round):** Blizzard, World of Warcraft Forever patch notes thread, 3 November and 1 October sections (thread 2360696, cited in sim round 21 B1/B2); Wowhead Forever tooltips for Defensive Stance, https://nether.wowhead.com/forever/tooltip/spell/71.

## Shipped orders used

Unchanged from round 23 (`addon-builds.json` did not move). Balance leveling: Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, ..., Insect Swarm 27, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonkin 40. Balance raid: Improved Starfire 33-37, Nature's Grace 38, Eclipse 39, 41, 42, Moonkin 40. Feral (both): Shredding Attacks 20-22, Feral Charge 25 (leveling), Rend and Tear 37 / 36, Shifting Power 40. Bear (both): Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19 and 22, Lacerate 25, Primal Bite 40, Rend and Tear 41-45. Restoration (both): Bloom 20, Nature's Swiftness 30, Swiftmend 43. Full lists in `round23/druid.md`.

## Note sweep

One line per card row note: spec, row and action, bands it renders, each named talent, spell, rank or value with the level the player has it (leveling / raid order where they differ), verdict.

- **Balance [1] Moonfire, end of fight. 10-60, ranks 1 to 10.** Wrath's cast; best instant (Insect Swarm no hit, Faerie Fire no damage, Hurricane channelled). OK.
- **Balance [2] Wrath, end of fight. 10-60.** Starfire 20, scoped; before 20 last 3.5 s. OK.
- **Balance [3] Moonfire upkeep. 10-60.** Improved Moonfire 20-21, Nature's Splendor 26, conditional. OK (Should carried).
- **Balance [4] Insect Swarm. 20-60.** Talent 27 both orders (routed row). One global. OK.
- **Balance [5] Eclipse Wrath. 10-60.** Eclipse 35 / 39 scoped; two charges; 0.17 s / 0.50 s (talent text); "even one point ... beat": **R24-1**; leveling 35-39 against :141: **R24-2**.
- **Balance [6] Wrath filler / mana. 10-60.** Starfire 20 scoped (gate verified); Innervate 40; Improved Wrath 10-14; fraction of the mana, half the cast at 20-60. OK (band 10 Should carried).
- **Balance [7] Starfire. 20-60, ranks 1 to 6.** Eclipse 35 / 39 scoped. OK; see R24-2 for the prose side.
- **Feral [1] Ravage. 32-60.** Prowl 20, Poised 32. OK.
- **Feral [2] Faerie Fire. 18-60.** Cat Form 20; free, 6 s. OK (band 18-19 Should carried).
- **Feral [3] Moonfire, below 20. 10-60.** OK (routed row).
- **Feral [4] Wrath, below 20. 10-60.** OK (routed row).
- **Feral [5] Shifting Power. 10-60.** Talent 40 (routed row); 40 energy, mana, 8 s. OK.
- **Feral [6] Rip. 20-60.** 5 points, under 2 s; Rend and Tear 37 / 36 scoped. OK.
- **Feral [7] Ferocious Bite. 24-60.** Predatory Discipline 31 / 30, conditional. OK.
- **Feral [8] Shred. 21-60, ranks 1 to 5.** "From level 21"; Shredding Attacks second point 21 both orders; "only at 20"; "harder at every rank" re-derived. OK.
- **Feral [9] Claw. 20 only.** 40 against 54 at 20; "later" clause is advice for a target you cannot get behind. OK.
- **Bear [0] Cooldowns. Page only.** Mighty Rage Potion in form. OK.
- **Bear [1] Barkskin. 44-60.** Trainer 44; 20%, 15 s, 60 s. OK.
- **Bear [2] Frenzied Regeneration. 36-60.** Trainer 36; half health, 20 rage. OK.
- **Bear [3] Primal Bite. 40-60, ranks 1 to 3.** Talent 40; 6 s (client); "biggest threat hit" routed. Not re-filed.
- **Bear [4] Lacerate. 25-60.** Talent 25; five stacks; 3 s; Rend and Tear 41 scoped. OK.
- **Bear [5] Faerie Fire. 18-60.** Bear Form 10; free, 6 s; threat. OK (band 18 Should carried).
- **Bear [6] Demoralizing Roar. 10-60, ranks 1 to 5.** 204 at 60 scoped; last 3 s; no stack with Demoralizing Shout. OK (Should carried).
- **Bear [7] Maul. 10-60, ranks 1 to 7.** 15 rage; 30 at 40-53 (gate verified); Ferocity 10 by 14; Lacerate 25 scoped; "makes sure ... every time": **R24-3**.
- **Bear [8] Swipe. 16-60, ranks 1 to 5.** Trainer 16; two or more targets and 40 rage (gate verified); up to three (tooltip); "Lacerate and Primal Bite" on 16-39: **R24-4**.
- **Restoration [1] Innervate. 40-60.** Trainer 40; under 30%. OK.
- **Restoration [2] Read Leyline. Alliance Skyborne, no addon row.** 50%, 15 s; not under Innervate. OK.
- **Restoration [3] Nature's Swiftness. 10-60.** Talent 30 (routed row); tank under 35%. OK.
- **Restoration [4] Healing Touch save. 10-60.** Newest trainer rank at every band. OK.
- **Restoration [5] Swiftmend. 10-60.** Talent 43 (routed row); under 75%, mana pace; keeps the HoT; 15 s; 20% base mana. OK.
- **Restoration [6] Tranquility. 30-60.** Trainer 30; two under 70%; every 2 s. OK.
- **Restoration [7] Regrowth, tank. 12-60.** Top rank every band; Bloom 20 scoped. OK.
- **Restoration [8] Rejuvenation, tank. 10-60.** Rank ladder matches rows (as round 23); Bloom 20 scoped. OK.
- **Restoration [9] Healing Touch filler. 10-60.** Rank ladder matches rows (as round 23). OK.
- **Restoration [10] Rejuvenation, party. 10-60.** Under 70%; no mana-pace gate. OK. Rows [11] to [14] carry no note.

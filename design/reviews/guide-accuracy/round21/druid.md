# Guide accuracy round 21: Druid

Audited on the round 21 audit tree (`scratchpad/gr21/audit`, head e039e1490: main after guide round 20; no sim round 21 commit in the log, so there is no sim round 21 druid report to read). Engine pin d5c708834 (`sim/enginever/version.go`), source read from `scratchpad/gr21/engine`. Active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).

Files: `web/src/content/guides/druid/{index,balance,feral,feral-bear,restoration}.md`, every row note in `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`, and every rendered row of `addon-data.json`: `rotations/druid-balance` (22 bands), `druid-feral` (24), `druid-feral-bear` (20), `druid-restoration` (28). Each row was read band by band against both shipped point orders in `addon-builds.json` (index 0 = level 10) and the trainer levels in `trainables/druid.json`.

Read first: the round 1, round 2 and round 21 briefs; `round20/druid.md`; the ledger's Round 19 and Round 20 sections; `round21/rogue.md` (the Levels 1 to 19 superlative, R21-1); earlier druid reports wherever a claim had been settled before (Gift of Nature and Tranquility, the downranking shares, Feral Charge 414924, the quest items).

## Changes since round 20's tree (4234643f5)

- The only druid changes are card notes (8589ac5f0, a591b2f0a): Balance Wrath [5] and Starfire [7] now give Eclipse by both orders; Feral Rip [6] gives Rend and Tear by both orders; Feral Shred [8] says "from level 22". `addon-data.json` and Data.lua were regenerated (a54761d7b, 9832b5feb). The first sentences of rows [5] and [8] on the addon match the new notes.
- No druid page prose, build string, `addon-builds.json` order, rank-choice row, BiS file or races.json row changed.

## Shipped orders (re-decoded from `addon-builds.json`; identical to round 20)

- **Balance leveling:** Improved Wrath 10-14, Genesis 15-16, Moonglow 17-19, Improved Moonfire 20-21, Nature's Majesty 22-23, Nature's Reach 24-25, Nature's Splendor 26, Insect Swarm 27, Vengeance 28-32, Improved Starfire 33, Nature's Grace 34, Eclipse 35, Moonfury 36-39 and 41, Moonkin Form 40, Nature's Focus 42-46, Heart of the Wild 47-51, Naturalist 52-56, Reflection 57-59, Subtlety 60.
- **Balance raid:** the same to 32, then Improved Starfire 33-37, Nature's Grace 38, Eclipse 39 (1 point), Moonkin Form 40, Eclipse 41-42 (3 points), Moonfury 43-47, Nature's Focus 48-52, Naturalist 53-57, Subtlety 58-60.
- **Feral leveling:** Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17, Brutal Impact 18-19, Shredding Attacks 20-22, Savage Fury 23-24, Feral Charge 25, Sharpened Claws 26-27, Predatory Strikes 28-30, Predatory Discipline 31, Blood Frenzy 32-33, Leader of the Pack 34, Natural Instinct 35-36, Rend and Tear 37-39, Shifting Power 40, Genesis 41-45, Nature's Majesty 46-47, Nature's Reach 48-49, Heart of the Wild 50-54, Naturalist 55-59, Natural Shapeshifter 60.
- **Feral raid:** Ferocity 10-14, Feral Swiftness 15-16, Feral Instinct 17-19, Shredding Attacks 20-22, Savage Fury 23-24, Sharpened Claws 25-26, Predatory Strikes 27-29, Predatory Discipline 30, Blood Frenzy 31-32, Leader of the Pack 33, Natural Instinct 34-35, Rend and Tear 36-39 and 41, Shifting Power 40, Genesis 42-46, Nature's Majesty 47-48, Nature's Reach 49-50, Heart of the Wild 51-55, Naturalist 56-60.
- **Bear (leveling = raid):** Ferocity 10-14, Furor 15, Feral Swiftness 16-17, Feral Instinct 18-19 and 22, Savage Fury 20-21, Thick Hide 23-24 and 26, Lacerate 25, Shredding Attacks 27-29, Leader of the Pack 30, Sharpened Claws 31-32, Predatory Strikes 33-35, Blood Frenzy 36-37, Natural Instinct 38-39, Primal Bite 40, Rend and Tear 41-45, Natural Reaction 46-50, Heart of the Wild 51-55, Naturalist 56-60.
- **Restoration (leveling = raid):** Heart of the Wild 10-14, Reflection 15-17, Nature's Focus 18-19 and 21-23, Bloom 20, Naturalist 24-28, Gift of Nature 29 and 31-34, Nature's Swiftness 30, Tranquil Spirit 35-39, Improved Rejuvenation 40-42, Swiftmend 43, Living Spirit 44-46, Improved Tranquility 47-48, Improved Regrowth 49, Genesis 50-54, Moonglow 55-57, Nature's Majesty 58-59, Nature's Splendor 60.

Frontmatter builds decode to these totals. Balance leveling: 32 Balance and 19 Restoration. Balance raid: 38 and 13. Feral leveling: 6 Balance, 39 Feral and 6 Restoration. Feral raid: 6, 39 and 10. Bear: 41 Feral and 10 Restoration. Restoration: 11 Balance and 40 Restoration. Every sum is 51.

## Findings

None. No page line or card note states something false on 70790 or contradicts the page, at any band it renders on, on either order.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 20's fixes, re-checked as new work

- **DRUID-R20-1** (Balance [5] and [7]). "From 35 on the leveling order and 39 on the raid order" matches the decoded orders: Eclipse is `[1,5,3]` at index 25 on the leveling order and at index 29 on the raid order. On the raid order the second and third points come at 41 and 42.
  - The second sentence's values match talents/druid.json: 0.17 s with one point and 0.50 s with three. "In the raid build" scopes the half second correctly. Two points (0.33 s at 41) are not quoted, so nothing is wrong there.
  - The Starfire note "(35 leveling, 39 raid order)" matches. The addon first sentence of row [5] is the new wording on all 22 bands.
- **DRUID-R20-2** (Feral Shred [8]). "Your builder from level 22" now agrees with the Claw note, "Your builder at 20 and 21", and with feral.md:119.
  - Shredding Attacks is 20-22 on both orders, and Ferocity is full by 14 on both.
  - Costs: Shred is 54 at 20, 48 at 21 and 42 at 22 (60 less 6 a point); Claw is 40 (45 less 5).
  - "8 to 14 more energy" is 48-40 and 54-40.
  - "Shred hits harder than Claw at every rank" holds on the client tooltips (Shred 200% plus 22/32/44/64/80, Claw 140% plus 27/39/57/88/115) for every pairing the bands cast. The closest pairing is Shred rank 1 against Claw rank 2 at 28-29, where Shred is ahead once the paw hits for more than 18.
  - The card's Ferocious Bite gate (24) is still routed, so it is not re-filed.
- **DRUID-R20-3** (Feral Rip [6]). "(37 leveling, 36 raid order)" matches: Rend and Tear is `[2,6,1]` at index 27 on the leveling order and index 26 on the raid order. The review wording "keeps Rend and Tear on" reads cleanly.

## Round 21 shapes, checked on every note and line

- **(1) Talent credited early on the raid order.** Only Balance and Feral ship a `raidBuild`.
  - Every level a Balance or Feral note names is checked against both orders: Eclipse, Rend and Tear, Shredding Attacks, Ferocity, Improved Wrath, Improved Moonfire and Nature's Splendor (conditional), and Predatory Discipline (conditional, 31 leveling and 30 raid). Each holds on both.
  - The Levels 1 to 19 paragraphs follow the leveling order by design, and the Talents sections send the raid order through the respec at 60. Balance's two orders are identical to 32. Feral's raid order trades Brutal Impact 18-19 for Feral Instinct 18-19, and feral.md:113 is leveling prose.
- **(2) Superlatives and comparisons.** A grep for best, biggest, hardest, only, most, beats, harder, always, never and every over all five pages and every note turned up these:
  - Balance [1] "Moonfire's hit is the best instant you have": true on bands 10 to 60. Insect Swarm has no direct hit, Faerie Fire does no damage and Hurricane is channelled.
  - feral.md:71 "Ravage hits much harder": against Pounce, Ravage is 435% plus 147/343 and Pounce is a 114/186 bleed over 18 s.
  - Feral [8] "harder than Claw at every rank": holds (see above).
  - Bear [3] and feral-bear.md:76 "your biggest threat hit": not provable either way on the client data. The engine scores it as Maul ahead per hit from 42, but both engine multipliers are its own stated assumptions. See Shoulds and Engine questions.
  - restoration.md:55 "the best single point you take": a judgement call, not a data claim.
  - No superlative appears in any Levels 1 to 19 paragraph. The rogue R21-1 shape (a span claim broken by a spell trained inside the span) is absent; see the Levels 1 to 19 notes below.
- **(3) Quoted values a talent changes, and lines a sim change left behind.**
  - Shifting Power "55% of your base mana": cost_pct 55. Natural Shapeshifter would cut it, but only the leveling build takes one point, and only at 60; the rotation text is the raid build.
  - Primal Bite "15 rage" with Ferocity at 60.
  - Rake "35 energy" (40 less Ferocity).
  - Tranquility "5 minutes to 2" (Improved Tranquility 60%).
  - Starfire "3 second cast" (raid Improved Starfire 5).
  - Wild Growth "2.5 second cast" (3.0 s less Naturalist 0.5) on its 6 s cooldown, over 7 s.
  - Swiftmend "20% of your base mana" (18562 cost_pct 20). No talent in either build changes it.
  - Barkskin "20% ... physical ... 15 seconds ... every minute" (22812).
  - Enrage "rage up front and 2 a second for 10 seconds" (5229: 10 now and 20 over 10 s).
  - Frenzied Regeneration "up to 10 rage a second ... 100 rage heals your whole health bar" (22842: 1% health a rage).
  - Demoralizing Roar "204 ... at 60" (9898: 193, -1.4 a level from 52, which is 204 at 60).
  - No druid line is left behind by a sim change: no sim round 21 has landed, and sim round 20's restoration edits were checked last round and are unchanged.

## Levels 1 to 19 (all four pages, line by line)

- **balance.md:125-131.** Trainer levels match: Wrath 1/6/14, Moonfire 4, Entangling Roots 8, Nature's Grasp 10, Faerie Fire 18.
  - Improved Wrath 10-14 is 0.5 s and 50% at five points. Genesis 15-16 is 4% periodic damage and healing. Moonglow 17-19 is 25% at three points.
  - Heeding the Call and Bear Form at 10 are right. Druids get no wands.
  - The Aquatic Form quest at 16 pays Aquarius Belt 16608 or Idol of Shifting Tides 263411.
  - Nothing here contradicts Levels 20 to 60 or the strip.
- **feral.md:111-115.** Trainer levels match: Moonfire 4, Bear Form, Maul and Demoralizing Roar 10, Bash 14, Swipe 16, Cat Form 20.
  - Ferocity 10-14 is 5 less rage for Maul and Swipe now, and less energy for Claw and Rake later. Feral Swiftness 15-16 is 4% dodge, with Cat Form run speed from 20. Feral Instinct 17 is 10% Swipe.
  - Brutal Impact 18-19 lengthens the Bash stun and cuts its cooldown by 30 s.
  - Weapon DPS on the paw is right, and the Idol of Shifting Tides exists.
- **feral-bear.md:113-117.** Trainer levels match: Bear Form, Maul, Growl and Demoralizing Roar 10, Enrage 12, Bash 14, Swipe 16, Maul rank 2 at 18.
  - Talents: Ferocity 10-14, Furor 15 (20% for 10 rage), Feral Swiftness 16-17 (4% dodge), Feral Instinct 18-19 (20% Swipe).
- **restoration.md:111-117.** Trainer levels match: Healing Touch 1, Rejuvenation 4, Regrowth 12, Revive 12, Cure Poison 14.
  - Heart of the Wild 10-14 is 10% Intellect and 20% Bear Form Stamina. Reflection 15-17 is 50%.
  - Nature's Focus 18-19 is 28%, which is "more than a quarter". Healing Touch, Rejuvenation and Regrowth are Nature spells.

## Page prose (all five pages, line by line)

Re-read in full: balance.md:43-137, feral.md:31-119, feral-bear.md:28-121, restoration.md:27-121 and index.md. Every talent number matches talents/druid.json. Every trainer level matches trainables/druid.json. Rank lines match the rendered rows:

- **Balance ranks:** Moonfire rank 2 at 20-28, rank 3 at 29-35, rank 4 at 36-39. Wrath rank 3 at 20-23, rank 4 at 24-39, rank 6 at 46-48. Insect Swarm rank 3 at 54-58.
- **Restoration tank Rejuvenation:** 20-39 rank 2 (night elf rank 3 at 29), 40-43 rank 4, 44-48 rank 5, 49-56 rank 6, 57-59 rank 8, 60 rank 10.
- **Restoration Healing Touch filler:** 20 rank 2 (tauren to 23), 30-33 rank 3, 34-39 rank 4, 40-45 rank 5, 46-49 rank 6, and the top rank otherwise.
- **Restoration party Rejuvenation:** rank 1, with rank 2 at 27-29 and 40-42, and rank 7 at 60.

Other claims checked:

- Races against races.json: Elune's Light 10% for 15 s on 3 min; Quickness 1% dodge; Endurance 5%; War Stomp 5 targets, 8 yd, 2 s; Wind Blessed 2%; Read Leyline 50% for 15 s on 2 min. Skyborne is named only as the paid option.
- Professions and gear: Molten, Frozen and Dormant Heart of the Mountain, Mender's Leather Gauntlets and Wolfshead Helm (item level 45, required level 40) exist. These lines and the BiS files are unchanged since round 20 checked them.
- Clearcasting text matches 16864 ("not consumed by Wrath"). Poised matches 1328829 (full energy, out of combat, 10 off).
- Voice: no em or en dashes, no nightly numbers, no site or sim talk.

## Shoulds (wording, not findings)

- **New: feral-bear.md:76 and bear note [3] "your biggest threat hit".**
  - The client says only that Primal Bite generates "a high amount of threat"; Maul's tooltip says nothing about threat. On the engine's own multipliers, Maul out-threats Primal Bite per hit from 42. Maul is (weapon + 71/101/128) x 1.1 Savage Fury x 1.75; Primal Bite is (weapon + 42/62/77) x 2.0, so Maul leads unless the paw hits for 700 or more.
  - Primal Bite is still the bigger threat per global and per rage, because Maul replaces a white swing.
  - Suggested wording that holds under any threat model: "On cooldown: your top threat button, every 6 seconds." In the prose: "It is your biggest threat button."
- **New: restoration.md:121** "at 20 and from 30 to 49 a lower Healing Touch is the better hole filler" leaves out the tauren's 21 to 23 that :85 and note [9] give. Suggested: "at 20 (to 23 for a tauren) and from 30 to 49".
- **New: Balance note [6], band 10.** "With Improved Wrath it is quick and cheap" reads at one point (0.1 s and 10%) on band 10. Improved Wrath is full by 14, so this is soft, not false.
- **Carried from round 20 (none applied):**
  - restoration.md:85 "Below 20, from 21 ...".
  - feral-bear.md:70: Thorns before Dire Bear Form.
  - feral.md:67: "Both builds skip Furor".
  - Feral Ferocious Bite "With Predatory Discipline" (31 leveling, 30 raid).
  - Balance Moonfire "With Improved Moonfire and Nature's Splendor" (from 26).
  - Bear Demoralizing Roar "204 ... at 60" on rank 1 bands.
  - Feral and bear Faerie Fire "In Cat/Bear Form it is free" (Cat Form row at 18).
  - balance.md:135 "until 39" should read "through 38".
  - "for good" at balance.md:115 and restoration.md:101.
  - restoration.md:57 "your next Healing Touch" should be "your next Nature spell".
  - Lines carried from round 19: feral.md:89, restoration.md:105, feral.md:73, balance.md:119, balance.md:105 (twice), feral.md:77 (55 with Wolfshead Helm), feral.md:85, balance.md:91.

## Decisions

- **D1. "Biggest threat hit" is a Should, not a finding.** No 70790 client value states a threat number for Maul or Primal Bite, and the engine's 1.75 and 2.0 are its own labelled assumptions (`sim/druid/threat.go`, "unconfirmed"). A search of the Forever sources found guides that put Primal Bite and Lacerate at the centre of bear threat, with Maul as spare-rage use; none gives numbers. The page claim is therefore not shown false. The suggested wording removes the per-hit reading.
- **D2. No re-file of routed items.** These stay with the sim lane and are not re-filed:
  - Feral Claw against Shred at 21 to 23. On client tooltip arithmetic Shred is ahead per energy at 21 once the paw hits for more than 4, while Claw is cheaper per combo point; only a Rip-cycle sim decides it, and that sim decides the Claw note's "does more for its energy" at 21.
  - The Shred row's Ferocious Bite gate.
  - Balance leveling Eclipse on against off.
  - The Balance empty filler at 10 to 16.
  - Rend and Tear on any bleeding target.
- **D3.** Rows that render before their own talent or trainer level remain the component lane's whole-row issue, as in rounds 19 and 20: Feral Shifting Power 10-39, Balance Insect Swarm 20-26, Restoration Nature's Swiftness 10-29 and Swiftmend 10-42, and the Feral "Below level 20" rows past 20.

## Sims requested

- None new. Carried and already routed: Feral Claw against Shred as the builder at 21 to 23 (round 20); Balance leveling Eclipse Wrath on and off at 36, 40 and 50 (round 19).

## Sim findings

- None new. The Feral Shred row's Ferocious Bite gate (24) against Shredding Attacks (22) stays routed from round 20.

## Engine questions

- **Bear threat multipliers.** Maul is scored at 1.75x its damage (Classic table, with Savage Fury's 10% on top) and Primal Bite at 2.0x (`HighThreatMultiplier`, labelled unconfirmed). Together they make a queued Maul a larger single threat event than Primal Bite at every band from 42 (from 40 once the paw hits for more than 137). cmangos Classic instead gives Maul 9881 a flat 207 bonus threat, under which Primal Bite leads.
  - The bear card ranks Primal Bite first either way, since it is off the global and Maul is on the swing, so the card's order is unaffected.
  - The question is which model Forever uses. It decides the threat-per-second figures and whether "biggest threat hit" can stand as written.
- Carried: Rend and Tear on any bleeding target against the druid's own bleeds only (round 20).

## Sources

- **Client 1.60.1.70790** (`data/builds/1.60.1.70790/`):
  - `addon-builds.json` (all four specs, both orders), `addon-data.json` (`rotations/druid-*`, every band and race key)
  - `talents/druid.json`, `trainables/druid.json`
  - `spelltips/druid.json`: 99, 9898, 5229, 22812, 22842, 407995, 1238069, 1238070, 6807, 8972, 9745, 9881, 1235827, 1328617, 5221, 9830, 1082, 9850, 6785, 9867, 9005, 9827, 1328829, 16864, 17116, 18562, 29166, 1322605, 9863, 9912, 9876
  - `spellconst/druid.json` (9898 points per level, Maul and Primal Bite effects), `races.json`, `items.json`
- **Curated:** `data/curated/apl/druid-{balance,feral,feral-bear,restoration}.json`.
- **Engine d5c708834:** `sim/druid/threat.go`, `maul.go`, `primal_bite.go`, `bear_form.go`, `talents_restoration.go`, `constants_auto_gen.go` (Maul and Primal Bite flats).
- **Reports:** `round20/druid.md`, `round21/rogue.md`, `ledger.md` Rounds 19 and 20, `round13`, `round14` and `round18` druid (rank shares, Primal Bite cost), and the `round4` decision on the quest items.
- **Web (Forever sources first; neither settles threat numbers):**
  - Mobalytics, "WoW Forever Tanking Tier List": https://mobalytics.gg/wow-forever/tier-lists/tanking
  - RankedBoost, "WoW Forever Druid Feral Combat Macros": https://rankedboost.com/world-of-warcraft/forever/macros/druid/feral-combat/

## Note sweep

One line per card row note: spec, row index and action, the bands it renders on, each named talent, spell, rank or value with the level the player has it, and the verdict.

- **Balance [1] Moonfire, end of fight. Bands 10-60.** Moonfire from 4; "best instant" holds at every band. OK.
- **Balance [2] Wrath, end of fight. Bands 10-60.** Starfire from 20; the 3.5 s branch covers bands 10-16. OK.
- **Balance [3] Moonfire upkeep. Bands 10-60.** Improved Moonfire 20-21 and Nature's Splendor 26 on both orders, as a conditional. OK (Should carried).
- **Balance [4] Insect Swarm. Bands 20-60.** Talent at 27 on both orders (whole-row issue); one global. OK.
- **Balance [5] Eclipse Wrath. Bands 10-60.** Eclipse at 35 leveling and 39 raid; 0.17 s at one point; half a second at three, reached at 42 on the raid order. OK.
- **Balance [6] Wrath while behind on mana. Bands 10-60.** Innervate from 40, Improved Wrath 10-14, Starfire from 20; half the cast is exact at raid 60 and close at 43% before that. OK (Should on band 10).
- **Balance [7] Starfire. Bands 20-60.** Eclipse at 35 leveling and 39 raid. OK.
- **Feral [1] Ravage. Bands 32-60.** Prowl from 20, Poised from 32 (10 off at full energy out of combat). OK.
- **Feral [2] Faerie Fire. Bands 18-60.** Cat Form from 20; free on a 6 s cooldown in Cat Form (3025). OK (band 18 Should carried).
- **Feral [3] Moonfire, below 20. Bands 10-60.** Cat Form at 20. OK (the row past 20 is the component issue).
- **Feral [4] Wrath, below 20. Bands 10-60.** As [3]. OK.
- **Feral [5] Shifting Power. Bands 10-60.** Talent at 40 on both orders (whole-row issue); 40 energy, costs mana, 8 s cooldown. OK.
- **Feral [6] Rip. Bands 20-60.** Rip from 20; 5 combo points, under 2 s; Rend and Tear at 37 leveling and 36 raid. OK.
- **Feral [7] Ferocious Bite. Bands 24-60.** Bite from 24; the 5 s and 2 s windows; Shred is live on the card from 24; Predatory Discipline at 31 leveling and 30 raid, as a conditional. OK (Should carried).
- **Feral [8] Shred. Bands 20-60.** Shred from 20; Shredding Attacks 20-22 on both orders; costs 54, 48 and 42 against Claw's 40 (Ferocity full by 14 on both); "harder at every rank" holds on every pairing the bands cast. OK. The card gate is routed.
- **Feral [9] Claw. Bands 20-60.** Claw from 20 at 40 energy; Shred 54 at 20 and 48 at 21; "does more for its energy" at 21 is in the routed sim. OK pending that sim.
- **Bear [0] cooldowns. Page only.** Mighty Rage Potion is usable in form. OK.
- **Bear [1] Barkskin. Bands 44-60.** Trainer at 44; 20% physical damage, 15 s, 60 s cooldown (22812). OK.
- **Bear [2] Frenzied Regeneration. Bands 36-60.** Trainer at 36; under 50% health with 20 rage. OK.
- **Bear [3] Primal Bite. Bands 40-60.** Talent at 40; 6 s cooldown (category cooldown 6000); "biggest threat hit" is a Should and an Engine question.
- **Bear [4] Lacerate. Bands 25-60.** Talent at 25; five stacks; 3 s refresh; 10% weapon damage per existing stack; Rend and Tear from 41. OK.
- **Bear [5] Faerie Fire. Bands 18-60.** Bear Form from 10; free on a 6 s cooldown; carries threat. OK.
- **Bear [6] Demoralizing Roar. Bands 10-60.** 204 at 60 (193 + 1.4 x 8 on rank 5); last 3 s of the 30 s; does not stack with Demoralizing Shout. OK (Should carried on rank 1 bands).
- **Bear [7] Maul. Bands 10-60.** 15 rage; Ferocity brings it to 10 by 14; Lacerate from 25, Primal Bite from 40. OK.
- **Bear [8] Swipe. Bands 16-60.** Trainer at 16; above 40 rage; three targets. OK.
- **Restoration [1] Innervate. Bands 40-60.** Trainer at 40; under 30% mana; multiplies Spirit regeneration (29166: +400%). OK.
- **Restoration [2] Read Leyline. Alliance Skyborne only, no addon row.** 50% for 15 s; does not cast under Innervate. OK.
- **Restoration [3] Nature's Swiftness. Bands 10-60.** Talent at 30 (whole-row issue); tank under 35%. OK.
- **Restoration [4] Healing Touch save. Bands 10-60.** The newest rank at every band (5186 to 9889). OK.
- **Restoration [5] Swiftmend. Bands 10-60.** Talent at 43 (whole-row issue); tank under 75%; keeps the HoT; 15 s cooldown; 20% of base mana. OK.
- **Restoration [6] Tranquility. Bands 30-60.** Trainer at 30; two of the party under 70%; ticks every 2 s. OK.
- **Restoration [7] Regrowth, tank. Bands 12-60.** Trainer at 12; top rank at every band; Bloom from 20. OK.
- **Restoration [8] Rejuvenation, tank. Bands 10-60.** Rank 10 at 60; Bloom from 20; rank 2 at 20-39 (night elf rank 3 at 29); rank 4 at 40-43; rank 5 at 44-48; rank 6 at 49-56; rank 8 at 57-59. All match the rows. OK.
- **Restoration [9] Healing Touch filler. Bands 10-60.** Tank under 55% while mana is ahead; rank 2 at 20 (tauren 21-23); rank 3 at 30-33; rank 4 at 34-39; rank 5 at 40-45; rank 6 at 46-49. All match the rows. OK.
- **Restoration [10] Rejuvenation, party. Bands 10-60.** Under 70%; does not wait on mana pace. OK. Rows [11] to [14] carry no note.

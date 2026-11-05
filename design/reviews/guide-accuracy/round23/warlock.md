# Guide accuracy round 23: Warlock

Audited on the round 23 audit tree (24ee457d0: main after guide round 22; no sim round 21 commit or report in its log; engine pin d5c708834 in `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warlock/{index,affliction,demonology,destruction}.md`, line by line; every row note in `data/curated/apl/warlock-{affliction,demonology,destruction}.json`, sentence by sentence; every rendered row in `addon-data.json` `rotations/warlock-*`, re-derived band by band (Affliction 23 bands, Demonology 25, Destruction 27); all six orders in `addon-builds.json` decoded level by level; `talents/warlock.json`, `spellconst/warlock.json`, `spelltips/warlock.json`, `trainables/warlock.json`, `items/warlock.json`, `combos.json`, `races.json`; engine `sim/core/eureka_auto_gen.go`, `sim/core/periodic_crit_auto_gen.go`, `sim/warlock/wrack.go`, `sim/warlock/talents.go`. Read first: round 22 warlock report, the ledger's round 19 to 22 sections, this round's other reports for the new shapes, and sim rounds 14 to 20's warlock reports for decoded client masks.

The round 1 and round 2 briefs are gone from the scratchpad (wiped). Method, statuses and structure follow the earlier reports, which quote them.

## Changes since round 22's tree (987d08a7c)

`git diff 987d08a7c HEAD` touches warlock in two lines, and no card, talent, spell, trainer or build file:

- demonology.md:124, "from 13 the Greater wand out-damages your bolts". Holds, narrowly (D1).
- destruction.md:128, "Drain Life, and Death Coil from 42, are your only heals". Holds (D1).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARLOCK-R23-1 | CONFIRMED | talent reach stated wider than the client, in the page's own loop | `affliction.md`:57, :71; `warlock-affliction.json` priority line 9 (Wrack, 40 to 60) | :57 "makes the target take 10% more from your other Shadow DoTs"; :71 "while it channels your other Shadow DoTs take 10% more"; note "your other Shadow DoTs take 10% more while it runs" | Client row: 1316697 effect 1 (record 1349557, aura 271, base 10) has class mask word 0 = 0x402, words 1 to 3 = 0, decoded on 70755 with the live push in sim round 15 G5 and round 16 S7, and unchanged since. Bit 0x2 is Corruption and 0x400 is Bane of Agony. `spellconst/warlock.json` family masks confirm the other side: Siphon Life [0, 1], Bane of Doom [0x80000000, 2]; Immolate's school is Fire anyway. The engine agrees (`wrack.go:37`, `wrackVulnerableSpellCodes`). So in the raid loop these same lines describe (Corruption, Doom, Siphon Life) only Corruption takes the 10% until the last-minute Agony. The tooltip's "other Shadow damage over time effects" is broader than the row. Same shape and precedent as Eureka! (ledger round 16: the client row decides and the page says so beside the tooltip) and MAGE-R23-2. See D2 for why this overturns R2-9 and G5's "may stay". | :57 "...makes the target take 10% more from your Corruption and Bane of Agony; whatever the tooltip says about your other Shadow DoTs, Siphon Life and Bane of Doom get nothing from it." :71 "...and while it channels your Corruption and Bane of Agony take 10% more." Note: "...and your Corruption and Bane of Agony take 10% more while it runs." Regenerate the note. |
| WARLOCK-R23-2 | CONFIRMED | talent reach stated wider than the client | `affliction.md`:32, :85 | :32 "Every DoT tick can crit, and Pandemic doubles the bonus on those crits"; :85 "Crit is good for Affliction because DoT ticks crit too, and Pandemic doubles the bonus on them." | Pandemic (427712): "Corruption, Bane of Agony, Bane of Doom, Drain Soul, Drain Life, Siphon Life, and Wrack" (`talents/warlock.json`; mask decoded in sim rounds 14 A2 and 16 A3; `talents.go`, `pandemicAffectedSpellCodes`). Immolate is not on it. It is one of the "three DoTs" of the page's own first line (:27), its burn does crit (`periodic_crit_auto_gen.go`: 348, 11668), and neither Affliction build takes Ruin. So "every" and "those" include a DoT in the loop that gets no doubled bonus. index.md:25 ("on Affliction's") and :52, :54 are exact. | :32 "Every DoT tick can crit, and Pandemic doubles the bonus on all of them but Immolate's burn". :85 "...because DoT ticks crit too, and Pandemic doubles the bonus on all but Immolate's." |
| WARLOCK-R23-3 | CONFIRMED | race: stated reason that fails on the numbers; "save" advice against a two minute cooldown | `demonology.md`:110 | "Eureka! gives 8 seconds of 10% more damage and 10% lower cost on your direct damage spells every two minutes, so save it for the last 35%, when every Shadow Bolt hits 6% harder" | Row 1259821: 8 s, 120 s cooldown. A window in the last 35% is worth 1.06 of a normal window. Holding from the pull gives up whole windows. Fight of 5 minutes: on cooldown 0:00, 2:00, 4:00 = 3 windows, the last already inside the 35% (from 3:15), 3.06 in all; held to 3:15 = 1 window, 1.06 (the next at 5:15 is too late). Fight of 3 minutes: on cooldown 0:00, 2:00 = 2.06; held = 1.06. Held wins only below about 2 minutes (at 2:00 flat, on cooldown 1.0 and held 1.06). The engine registers Eureka! as a DPS major cooldown and the card's `autocastOtherCooldowns` fires it on cooldown, so the page's "a step behind" gap was measured on-cooldown too. The same paragraph tells Orc to press Blood Fury with trinkets "and again for the last 35% if it is back", which is the right shape. | "...every two minutes. Press it with your trinkets, and again for the last 35% if it is back, when every Shadow Bolt hits 6% harder, and Expansive Mind adds 5% mana." |
| WARLOCK-R23-4 | NEEDS-SIM | stated reason that may hold only in some gear | `warlock-demonology.json` priority line 1, third sentence (48 to 60); `demonology.md`:84 | note "In the raid build, skip it: Soul Fire is Fire, and your Shadow bonuses make Shadow Bolt stronger."; :84 "so it does no more than the Shadow Bolt it replaces" | No committed sim casts Soul Fire on the raid build: the Pact gate only keeps it out (sim rounds 15 G3, 16 S4, 18 D1), and round 18 D2's sweep is ranks on the bare leveling build. Client arithmetic, raid 60, per second inside Decimation, spell power x, Imp 15%, Master Demonologist 10%, Soul Link 3%, Decimation 6%: Soul Fire (431 + x) x 1.03 over 2.4 s against Shadow Bolt rank 9 (251 + 0.857x) x 1.381 x (1 + 0.2u) over 2.5 s, u = Improved Shadow Bolt uptime. At u = 0.5 they cross near 350 spell power; at 600 Shadow Bolt leads (about 465 against 442) and the line holds; at 100 Soul Fire leads (228 against about 205). On the raid order at 48 (rank 1, no Bane yet, Improved Shadow Bolt 3 of 5) Soul Fire also leads on a level 48 character's spell power, per cast clearly and per second narrowly. So the reason is true at raid power and may be false below it, on bands the note renders. Fire-only and Shadow-only gear moves it further. | Decide by the sim below. If Soul Fire wins below raid power, scope both lines: "at raid spell power your Shadow bonuses make Shadow Bolt stronger". demonology.md:98's "since this build skips Soul Fire" follows the result. |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 1, NEEDS-OWNER 0. Two card notes (R23-1 is prose plus a note; R23-4 is a note and prose), the rest page prose. R23-1's note change needs `addon-data.json` and Data.lua regenerated, note text only.

## Shoulds (wording, not findings)

New:
- **affliction.md Races, "the damage half also reaches Wrack and your drains".** True, but it reads as if the cost half does not. Row 1259821's cost mask (word 1 8650880, word 0 541165) covers Wrack and both drains as well, so Wrack costs 180 in the window. Suggested: "both halves also reach Wrack and your drains".
- **destruction.md:128, "so as a finisher it is free".** The sentence is about the shard (it says so just before), and Shadowburn still costs 105 to 365 mana. "so as a finisher it costs no shard" is exact.
- **destruction.md:128, "Until Shadow and Flame reaches 5 points, Conflagrate can eat your Immolate"** follows "at 32" with no level for Conflagrate (34 leveling). It tells the player what to do "once" they cast it, so it does not credit it early; part of the routed R19-3 wording anyway.

Carried: affliction.md:115 "the Shadow Bolts it replaces"; demonology.md:90 "every wipe or dead Succubus" applies to all three specs; R18-S1 (Affliction Corruption "Instant with Improved Corruption" on 10 to 20); the Demonology Corruption note written for the raid build; the Demonology and Destruction Agony "new Doom" tail; the Destruction Shadowburn shard below Shadow and Flame 5; Shadow Bolt rows rendering past 40 (component lane); the round 17 list.

## Decisions

- **D1. Round 22's two fixes hold.**
  - demonology.md:124: Greater Magic Wand (11288) 22 to 41 over 1.8 s = 17.5 a second, +2 spell power, required level 13. Shadow Bolt rank 3 (705): spellconst amount 40 is the minimum (variance 0.154, as Shadowburn's 66 against its "65 to 74" shows), plus 0.6 a level to 17, so about 43.7 at 13 and 46.3 from 17; with the wand's 2 spell power at 0.8, 16.2 to 17.1 a second over 2.8 s. The wand leads at every level from 13 to 19 on a character with no other spell power, narrowly at 17 to 19, and the below-20 spell power cut only helps it. A piece or two with spell power could tie it, so it is not a wide margin, but the line states no size. Rank 4 trains at 20, outside the paragraph. Not filed.
  - destruction.md:128: Death Coil 6789 `spell_level` 42; Drain Life from 14; the leveling order has no Siphon Life and no Improved Drains at any level (Soul Harvest at 59 and 60 is mana). Healthstones are items and Demon Armor's regeneration is not a heal button. "Only" holds.
- **D2. R23-1 overturns round 2's R2-9 and sim round 15 G5's "may stay".** The round 1 lesson (a client spell description beats an inferred flag) does not apply: 0x402 is not inferred, it is the effect's own class mask, decoded bit by bit in two sim rounds and matched by a passing engine test. The page already follows the row over the tooltip for Eureka!, and this round's lens names exactly this shape. It is a page finding, not an engine question: the engine already follows the row. A launch log would confirm it (sim round 20 L1 already names a 40+ Affliction log), and the fix wording stays true of the row today.
- **D3. Eureka! lines against row 1259821, re-derived from `eureka_auto_gen.go`.** Damage: Shadow Bolt, Immolate's hit, Searing Pain, Soul Fire, Shadowburn, Conflagrate, Death Coil, Rain of Fire, Hellfire. Periodic: Drain Life, Drain Soul, Wrack. Cost: all of those; not Incinerate, Corruption, the banes or Siphon Life. Affliction's and Destruction's statements hold (Should above aside), including "Incinerate gets nothing from it" and "press it as you start a Wrack" (6 s channel in an 8 s window). Eureka! on Incinerate stays with L1.
- **D4. Other masks checked:** Suppression 5% on every spell (mask 0); Ruin the Destruction list, not Corruption; Agonizing Flames the Destruction nukes; Shadow Mastery all Shadow spells including Shadowburn; Malevolence every warlock Shadow spell, ticks and Wrack included (Destruction's "Corruption ticks, Doom and Shadowburn" is exact); Improved Bane of Agony on 0x400; Aftermath the initial hit only; Bane's four cast times; Destruction's Pandemic line is limited to Corruption and both banes and is exact. Soul Siphon counts Corruption, the bane and Siphon Life, as stated.
- **D5. Every superlative, "only" and comparison, both sides, every band:**
  - Wrack "out-damages Shadow Bolt for a fraction of the mana" (:57, :71, note, 40 to 60): Wrack at level 40 is 36 a tick x 1.36 = 49 a second, at 60 is 86 x 1.36 = 117; 0.389 spell power a second. Shadow Bolt: rank 6 at 40 is 34 a second over 3 s; rank 9 at 60 is 100 a second and 0.343 spell power a second over 2.5 s (Bane 56 L / 52 R). Wrack leads on both terms on every band, for 33 mana a second against 70 to 148.
  - Malevolence line, "a Nightfall bolt is instant but still costs full mana, while Wrack does more for less" (:54): per cast and per mana true (Wrack 702 + 2.33 spell power for 200, bolt 251 + 0.857 for 370). An instant bolt leads per second of global, which is why the proc is tempting, but sim round 20 A4 measured Nightfall with a Trance bolt at -0.3 to -0.4%, so the conclusion stands.
  - Doom against Agony per minute (:80): 1742 + 4.0 against 552 x 2.5 = 1380 + 4.0 at 60. Mana 300 against 537. Holds, 60 only.
  - Searing Pain against Incinerate (destruction.md:84, 40 to 60), rank for rank per cast, per second and per mana with Bane: holds. Demonology's Searing Pain is R20-4 (routed).
  - "Wands do most of your killing", "carry you to 20": no per-second claim.
  - "The most Shadow bonuses of any warlock build": holds (Imp, Master Demonologist, Soul Link and Improved Shadow Bolt, against Shadow Mastery alone or Shadow and Flame).
  - "only a new Imp ends it", "Every warlock build takes all five" (Suppression is in all six orders: 16, 16, 46, 45, 51, 49), "Affliction has no execute spell" (Drain Soul's tips have no execute clause in Forever), "Spirit only matters through Life Tap" (raid stat paragraphs, no in-combat regeneration in those builds): hold.
  - "Conflagrate and Shadowburn on cooldown burn mana fast" (destruction.md:40): with Cataclysm's 10%, 153 and 219 mana a second of global against Incinerate's 146 with Bane. Holds.
- **D6. "On cooldown" and "whenever" lines against the card order.** Destruction's "on cooldown" lines sit under the stated priority (Corruption, Doom, Immolate first), which the page lists and the card matches. Blood Fury is never "on cooldown": it is tied to trinkets. Potion and Rune "on cooldown" is qualified by low mana. "Press it as Conflagrate and Shadowburn come off cooldown" is one damage racial at the two spells it covers, and Berserking, Perception and the like carry no such line. R23-3 is the one line that gets a cooldown wrong.
- **D7. Quoted costs, durations, cooldowns:** Curse of Recklessness 505; Curse of the Elements 10% / 75, 5 minutes every rank; Amplify Curse 3 min, 30 s, 50% / 20%; Siphon Life 30 s; Agony 24 s, 552 total at rank 6, "builds up" every rank; Doom 1 min, 60 s cooldown; Wrack 6 s; Improved Corruption 0.4 s a point, instant at 21; Corruption 18 s from rank 3; Bane 1.5 / 2 / 2.5 / 4 s; Conflagrate 10 s; Shadowburn 15 s and 8 s window; Soul Fire 60 s x 0.1 = 6 s, 6 s - 2 s = 4 s, x 0.6 = 2.4 s; Improved Shadow Bolt 12 s; Shadow and Flame 20 s; summons 10 s - 4 s = 6 s; Portal five shards and 1 hour, caster and 2 more, 2 to summon; Firestone 2% crit and Fire damage on the weapon (17953); Spellstone on the weapon (17728, items 5522, 13602, 13603 inventory type 0); Greater wand 22 to 41 at 1.8; Lesser wand 12 to 22 at 1.5. All hold at the ranks and talents the builds have. No mana cost is quoted as exact on any page.
- **D8. Not re-filed (routed or open):** WARLOCK-R20-4; R19-3 and R20-3 (Conflagrate on cooldown against held, and the Immolate wording that follows it); Affliction Amplify Curse below 60; the engine's missing Soul Shard; sim round 20 L1 items; race re-runs; BiS race captions; rows rendering before their talent (Amplify Curse 20 to 28 L, Shadowburn 20 to 29, Conflagrate 25 to 33).

## Sims requested

- **New (decides R23-4): Demonology raid build, Soul Fire inside Decimation (the row without its `not auraIsKnown 425464` gate) against the shipped skip.** At 48, 52 and 56 on the raid order with band gear, and at 60 bare, half and full raid gear. Report Shadow Bolt and Soul Fire shares and Improved Shadow Bolt uptime.
- Optional (confirms R23-3): Gnome Demonology raid 60, Eureka! on cooldown against held until 35%, at the standard fight length. The arithmetic already decides it.
- Carried with the sim lane: WARLOCK-R20-4; R19-3; Affliction Amplify Curse timing below 60; the Demonology Horde race re-run.

## Sim findings

- None new. Still open: the engine models no Soul Shard (Shadowburn 30 to 43, Soul Fire outside Decimation), and the sim round 20 L1 items.

## Engine questions

- None new. The engine follows the client rows on every mask this report quotes. Decimation's 0% client effect stays the known gap: the page's "6% harder" follows the talent text, and R23-3's arithmetic takes it at face value.

## Checked and right this round

- **index.md**, every line: slots; Curse of the Elements; DoT crits and "Pandemic ... on Affliction's"; new talents; swapped sacrifices; Improved Shadow Bolt yours alone; pet buffs; Portal; no Unstable Affliction or Haunt; races against `combos.json` (Human, Orc, Undead, Gnome, Troll new).
- **affliction.md**: FS1 strings decoded against the orders (leveling 35 / 11 / 5, raid 35 / 7 / 9); every talent value; globals; rotation against the card; Levels 1 to 19 and 20 to 60 against the order and trainer (Corruption 4, Life Tap 6, Agony 8, Drain Life 14, Improved Corruption 17 to 21, Soul Harvest 24, Siphon Life 30, Soul Siphon 36 to 38, Wrack 40, Succubus 20, Felhunter 30).
- **demonology.md**: FS1 strings (30 leveling, 31 raid with Pact at 40); prereqs (Soul Link needs Sacrifice, Pact needs Soul Link); the raid build's swaps; Decimation's text; setup against the prepull; Levels 1 to 19 and 20 to 60.
- **destruction.md**: FS1 strings (35 / 0 / 16 and 33 / 0 / 18), the six points moved into Pandemic and Malevolence; Incinerate's prereq on Bane of Havoc; opener and priority against the card; rank rows (Shadow Bolt 1106 at 28 and 32, 1088 at 30; Immolate 2941 at 40, 11665 at 41; Corruption 6223 at 41, 7648 at 42; Conflagrate rank 3 at 40); Levels 1 to 19 and 20 to 60.
- **Voice:** no em or en dashes, no nightly numbers, no site or sim talk outside source URLs.

## Sources

- Client 1.60.1.70790: `talents/warlock.json`, `spellconst/warlock.json`, `spelltips/warlock.json`, `trainables/warlock.json`, `items/warlock.json`, `items.json`, `combos.json`, `races.json`, `addon-builds.json`, `addon-data.json` (`rotations/warlock-*`, every band).
- Decoded client class masks: `design/reviews/2026-11-04-sim-audit-round15/warlock.md` (A5, G5, E1), `round16/warlock.md` (S7, A3), `round14/warlock.md` (A2, X1), `round18/warlock.md` (X2), `round20/warlock.md` (A4, A5, X3). A direct re-read of wago.tools SpellEffect (https://wago.tools/db2/SpellEffect?build=1.60.1.70790) returned 403, so unknown, not contradicted.
- Engine at d5c708834: `sim/core/eureka.go`, `sim/core/eureka_auto_gen.go`, `sim/core/periodic_crit_auto_gen.go`, `sim/warlock/wrack.go`, `sim/warlock/talents.go`, `sim/warlock/spellstone_test.go`.
- Curated: `data/curated/apl/warlock-{affliction,demonology,destruction}.json`. Earlier rounds: `round22/warlock.md`, `ledger.md` rounds 19 to 22, `round23/mage.md` (the same mask shape).
- No other web research was needed: client data and arithmetic decide every call but R23-4, which goes to the sim.

## Note sweep

Format: spec, row (action), bands it renders, each named talent, spell, rank or value with the level the player has it on the leveling and raid orders (L / R), then the verdict. Rows and bands re-derived from `addon-data.json` this round; all note text unchanged from round 22.

**Affliction**
1. Life Tap (1454 to 11689), 10-60. "a point for every point of Spirit": every rank base plus Spirit x 1; Healthstone (Create Healthstone 10). OK.
2. Curse of the Elements (440892 to 1311680), 20-60. Five minutes every rank; own slot; one per target. OK.
3. Corruption (172 to 11672), 10-60. "Instant with Improved Corruption": 17-21 L / R, conditional. OK (Should R18-S1 on 10 to 20).
4. Bane of Doom (603), 60. One global; big hit after a minute; Pandemic 31-33 L / 27-29 R; on its list. OK.
5. Amplify Curse (18288), 20-60. Talent 29 L / 26 R (earlier rows component lane); 50% on Agony; no global. OK.
6. Bane of Agony (980 to 11713), 10-60. "all fight before 60" (Doom 60); "biggest ticks come last" (every rank "builds up"); Amplify "once you have" 29 L / 26 R. OK.
7. Siphon Life (18265 to 18881), 30-60. Talent 30 L / R; 30 s every rank. OK.
8. Immolate (707 to 11668), 10-60. Direct hit above one tick every rank; 2.0 s cast (1.5 s with Bane 56-60 L / 52-56 R) lands before the 3 s mark. OK.
9. Wrack (1316697), 40-60. Talent 40 L / R; Soul Siphon full 38 L / R; out-damages Shadow Bolt on both terms at every band; "your other Shadow DoTs take 10% more": **WARLOCK-R23-1**.
10. Shadow Bolt (695 to 11661), 10-60. "until Wrack at 40": 40 L / R. OK (renders past 40: component lane).
11. Shoot (5019), drawer only. Wand shot not during a global. OK.

**Demonology**
12. Soul Fire (6353, 17924), 48-60. Trained 48 / 56; Decimation 2 by 31 L / R (6 s, 40%, 10 s, no shard); "2.4 seconds once Bane is full": Bane 56-60 L / 51-55 R, conditional; "never worth its shard" outside Decimation OK. "your Shadow bonuses make Shadow Bolt stronger": **WARLOCK-R23-4**.
13. Life Tap, 10-60. As row 1. OK.
14. Curse of the Elements, 20-60. As row 2. OK.
15. Corruption, 10-60. "A 2 second cast without Improved Corruption": raid never takes it; leveling 47-51. OK (Should, carried).
16. Bane of Doom, 60. One minute; "never lands" in the last minute. OK.
17. Bane of Agony, 10-60. First sentence scoped to 60; "new Doom" tail carried Should. OK.
18. Immolate, 10-60. As row 8 (Bane 56-60 L / 51-55 R). OK.
19. Shadow Bolt, 10-60. "beats Searing Pain" (Searing Pain 18): WARLOCK-R20-4, routed. Master Demonologist 35-39 L / R, conditional. Imp's sacrifice 15% "with Demonic Pact at 40 in the raid build": Pact 40 R, sacrifice 26 R. OK.
20. Shoot, drawer only. OK.

**Destruction**
21. Curse of the Elements, 20-60. As row 2. OK.
22. Corruption, 10-60. "until Improved Corruption makes it instant": 52-56 L / 50-54 R; 2 s every rank; one cast for the whole run. OK.
23. Life Tap, 10-60. As row 1. OK.
24. Bane of Doom, 60. OK.
25. Bane of Agony, 10-60. As row 17. OK.
26. Immolate, 10-60. Refresh at 3 s or less; Conflagrate 34 L / 32 R and Incinerate 40 L / R behind "once you have"; "until Shadow and Flame is full" 44 L / 42 R. OK (wording follows R19-3).
27. Conflagrate (1293817 to 18932), 25-60 (talent 34 L / 32 R; earlier rows component lane). Shadow and Flame 20 s, 4% a point to 20%, never eats Immolate at 5: 39 and 41-44 L / 37-39 and 41-42 R. OK.
28. Shadowburn (17877 to 18871), 20-60 (talent 30 L / 28 R). Any health; Fire buff and refund at 5 points as row 27. OK (shard below 5 points carried).
29. Incinerate (412758 to 1293813), 40-60. Talent 40 L / R; 25% with Immolate every rank. OK.
30. Shadow Bolt (695 to 11661), 10-60. "before you learn Incinerate at 40": 40 L / R. OK (renders past 40: component lane).
31. Shoot, drawer only. OK.

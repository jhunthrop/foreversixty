# Guide accuracy round 22: Warlock

Audited on the round 22 audit tree (987d08a7c: main after guide round 21; no sim round 21 commit in its log; engine pin d5c708834 in `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warlock/{index,affliction,demonology,destruction}.md`, line by line; every row note in `data/curated/apl/warlock-{affliction,demonology,destruction}.json`, sentence by sentence; every rendered row in `addon-data.json` `rotations/warlock-*`, band by band (Affliction 23 bands, Demonology 25, Destruction 27); both shipped orders in `addon-builds.json`; `talents/warlock.json`, `spellconst/warlock.json`, `trainables/warlock.json`, `items/warlock.json`. Read first: round 21 warlock report, the ledger's round 19, 20 and 21 sections, and the round 1, 2 and 22 briefs.

## Changes since round 21's tree (e039e1490)

- `git diff e039e1490 HEAD` touches no warlock card (`warlock-*.json` unchanged), no warlock line in `addon-data.json` (the regenerated note text is other classes only), and no talent, spell, trainer, item or build file. The rendered rows and both orders are exactly round 21's (round 20's decoded orders stand).
- Round 21's three warlock fixes (92c480d96, plus c532a59a4 rewording the Soul Harvest line):
  - demonology.md:124: "a Lesser Magic Wand at 5 and a Greater Magic Wand at 13 cost no mana, and from 13 the wand out-damages your bolts."
  - affliction.md:115: "Cast Drain Soul on the last few percent of a mob's health whenever you need Soul Shards; from 24, Soul Harvest turns each of those kills into mana regeneration, so you drink less."
  - destruction.md:128: "on level 30 gear your mana cannot carry rank 5 as well as the cheaper rank 4".

All three hold (D1).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARLOCK-R22-1 | CONFIRMED | leveling prose (trained spell named with no level on levels the trainer has not given it; round 21's R21-3 shape) | `destruction.md`:128 | "no Siphon Life and no talented drains, so Drain Life and Death Coil are your only heals" | The paragraph is "Levels 20 to 60" and gives a level for everything else it names ("Before 40", "from 30, once you take the talent", "rank 5 arrives at 28", "Shadowburn joins at 30", "Shadow and Flame is full at 44", "Incinerate at 40", "At 41", "from 42"). Death Coil rank 1 (6789) is `spell_level` 42 (`spellconst`), with ranks at 50 and 58; the trainer gives nothing earlier. So from 20 to 41, more than half the range, the player's only heal is Drain Life and the line credits a spell they do not have. Round 21 D3 made exactly this shape CONFIRMED for Soul Harvest in affliction.md:115 ("a 20 to 60 range that gives levels for everything else"); this line was carried as a Should since round 19 and is the same shape, so it is CONFIRMED for consistency. | "...no Siphon Life and no talented drains, so Drain Life, and Death Coil from 42, are your only heals." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. Page prose; no card note is wrong this round.

## Shoulds (wording, not findings)

- **affliction.md:115, Wrack against "the Shadow Bolts it replaces"** (carried from round 21, not applied). The sentence says "Drain Life while the DoTs run, and from 40 channel Wrack instead", so Wrack replaces Drain Life, and then that it costs less than "the Shadow Bolts it replaces"; the paragraph never names Shadow Bolt. Both are defensible (the card and :54 make Shadow Bolt the pre-40 filler), and the cost claim is true (Wrack 200 per 6 s channel against two rank 6 bolts), but a reader cannot tell what Wrack replaces. Against Drain Life rank 4 (7651: 185 for 5 s) Wrack would not be "far less". Suggested: "...three of your effects on the target add 36% to both, and Wrack does more than Shadow Bolt for about half the mana per second."
- **demonology.md:124**: "a Lesser Magic Wand at 5 and a Greater Magic Wand at 13 cost no mana, and from 13 the wand out-damages your bolts." True (D1), but "the wand" after two wands reads as either; "from 13 the Greater wand out-damages your bolts" is exact.
- **demonology.md:90** (carried): "every wipe or dead Succubus is another summon" applies to all three specs; the difference is Master Summoner's 6 second resummon.
- **Carried**: R18-S1 (Affliction Corruption "Instant with Improved Corruption" on 10 to 20); the Demonology Corruption note written for the raid build; the Demonology and Destruction Agony "new Doom" tail; the Destruction Shadowburn shard below Shadow and Flame 5; Destruction and Affliction Shadow Bolt rows rendering past 40 beside Incinerate / Wrack (component lane; both notes are true and scoped to before 40); the round 17 list.

## Decisions

- **D1. Round 21's three fixes hold.**
  - demonology.md:124: Greater Magic Wand (11288) 22 to 41 over 1.8 s = 17.5 a second, required level 13 (`items/warlock.json`); Shadow Bolt rank 3 (705) 40 at 0.8 over 2.8 s = 14.3 a second (about 17 with 8 spell power, more than a 13 to 19 character carries); rank 4 (1088) is 20, outside the paragraph; the Demonology order to 19 is Improved Health Funnel, Demonic Embrace and Unholy Power, nothing that touches Shadow Bolt. "cost no mana" true of both. No longer claims anything at 12.
  - affliction.md:115: Soul Harvest is the leveling order's 24 and 25; rank 1 "increases your Mana regeneration by 50% for 10 sec", rank 2 100%, both on "Killing a non-trivial target afflicted by your Drain Soul" (`talents/warlock.json`). "each of those kills" refers back to the Drain Soul kills. Holds at both ranks.
  - destruction.md:128: no number left; rank 4 (1088, 110) is cheaper than rank 5 (1106, 160) with or without Cataclysm's 6%. Holds.
- **D2. R22-1 is CONFIRMED**, not carried again as a Should: it is the R21-3 shape in a paragraph that dates every other spell, and the brief's lens is "credit no trained spell before the trainer gives it". The Healthstone is an item, not a heal spell; not filed.
- **D3. Superlatives and comparisons, every one on all four pages, at the ranks the rows cast and every spell the player has at that level:**
  - "Doom does more damage per minute than Agony for one global instead of two and a half" (demonology.md:80): 60 only (Doom 60). Holds (round 21 arithmetic, unchanged data).
  - Wrack "out-damages Shadow Bolt for a fraction of the mana" (affliction.md:57, :71, Wrack note, 40 to 60): Wrack 6 x 36 at 0.286 for 200 (1316697) against Shadow Bolt 0.857 per 3.0 s (2.5 s with Bane 56 L / 52 R); Soul Siphon full 38. Equal coefficient per second, higher base, 33 mana a second against 70 to 148. Holds every band.
  - "Wrack costs far less mana than the Shadow Bolts it replaces" (affliction.md:115, 40 to 60): true per second and per channel against two bolts (see Should).
  - Malevolence "about level with Nightfall" 34 to 39, ahead from 40 (affliction.md:54): measured in earlier rounds; consistent with the card's Shadow Bolt filler before 40.
  - "Soul Fire ... does no more than the Shadow Bolt it replaces" (demonology.md:84, raid 60): Soul Fire rank 2 (17924) 431 at 1.0 over 2.4 s with only Soul Link's 3%; Shadow Bolt 251 at 0.857 over 2.5 s with the Imp's sacrifice 15%, Master Demonologist 10%, Soul Link 3%, Decimation's 6% and Improved Shadow Bolt's 20% at its uptime. At raid spell power (about 600 to 700) Shadow Bolt is ahead per second of cast, by a margin that depends on Improved Shadow Bolt's uptime; the costs (370 against 335) are close. Holds; the raid skip is the card's measured choice.
  - "Searing Pain ... Incinerate does more" (destruction.md:84): holds 40 to 60 (round 21 arithmetic).
  - "The most Shadow bonuses of any warlock build at once" (demonology.md:32): holds.
  - "her melee and Lash of Pain out-damage the Imp's Firebolt" (affliction.md:63, destruction.md:67): accepted since early rounds; no client value contradicts it.
  - "Conflagrate and Shadowburn ... most of your damage" / "Fire power feeds Incinerate, Immolate and Conflagrate, which are most of your damage" (destruction.md:98): explains, does not rank stats; accepted.
  - "Wands do most of your killing before 20" (affliction.md:111) and "carry you to 20" (destruction.md:124): no per-second claim. Checked anyway: Destruction with Bane 5 at 16 casts rank 3 in 2.3 s (17.4 a second), level with the Greater wand; no comparison is stated, so nothing to file.
  - "Orc fits Demonology best of any Horde race", "a step behind", "close behind": race lines with the sim lane (race re-runs); not re-filed.
  - "only a new Imp ends it" (Pact), "only earns a cast on fights with two targets" (Havoc), "Spirit only matters through Life Tap", "your only heals" (R22-1): the first three hold on the client text.
- **D4. Quoted costs, durations and cooldowns, every one, with the build's discounts:** Curse of Recklessness 505 armor (rank 4); Curse of the Elements 10% / 75, five minutes every rank; Amplify Curse 3 min, 30 s, 50% / 20%, no global; Siphon Life 30 s; Bane of Agony 24 s; Doom 1 min; Wrack 6 s channel; Improved Corruption 0.4 s a point, instant at 21 (Affliction 17 to 21); Bane 1.5 / 2.0 / 2.5 / 4 s; Bane "half a second off Shadow Bolt and Immolate" at 16 (Destruction 12 to 16); Destructive Reach 20% at 11; Molten Skin 6% at 19; Decimation 6 s cooldown and 2.4 s cast (6 s x 0.6 with Bane full; 60 s x 0.1); Soul Fire 4 s with Bane outside Decimation; summons 6 s with Master Summoner; Improved Shadow Bolt 12 s; Shadow and Flame 20 s; Conflagrate 10 s; Shadowburn 15 s and the 8 s refund window; Corruption 18 s; Portal of Summoning five shards, one hour; Firestone 2%; Spellstone and Firestone one shard; Death Coil (R22-1). No cost number is quoted any more; nothing else is changed by a talent the build holds at that level.
- **D5. Notes that compare to a spell, checked against its trainer level on every band:** Affliction Wrack (40 to 60) names Shadow Bolt (1). Affliction Shadow Bolt names Wrack "at 40", dated. Demonology Shadow Bolt (10 to 60) names Searing Pain (18): bands 10 to 16 are folded into WARLOCK-R20-4 (round 20 Carried line), routed; not re-filed. Demonology Soul Fire (48 to 60) names Shadow Bolt and Bane, Bane conditional ("once Bane is full"). Destruction Immolate names Conflagrate and Incinerate behind "once you have". Destruction Shadow Bolt names Incinerate "at 40", dated. Life Tap notes name a Healthstone (Create Healthstone 10; rows from 10). All OK.
- **D6. Not re-filed (routed or open):** WARLOCK-R20-4; R19-3 and R20-3 (Conflagrate on cooldown against held, and the Immolate wording that follows it, including destruction.md:128 "cast it once the DoT has done most of its work"); Affliction Amplify Curse below 60; the engine's missing Soul Shard; sim round 20 L1 items; race re-runs; BiS race captions; rows rendering before their talent (Amplify Curse 20 to 28 L, Shadowburn 20 to 29, Conflagrate 25 to 33).

## Sims requested

- None new. Carried with the sim lane: WARLOCK-R20-4 (Demonology Shadow Bolt against Searing Pain at 18, 20, 26, 34, 42, 50 and 58); R19-3 (Destruction Conflagrate held against on cooldown at 34 to 43 L / 32 to 41 R); Affliction Amplify Curse timing below 60; the Demonology Horde race re-run.

## Sim findings

- None new. Still open: the engine models no Soul Shard (Shadowburn 30 to 43, Soul Fire outside Decimation), and the sim round 20 L1 items.

## Engine questions

- None. Every value above is a client `spellconst`, `talents` or `items` value; nothing here depends on engine behaviour the client contradicts.

## Checked and right this round

- **index.md**, every line: bane and curse slots; Curse of the Elements' schools and rank 4 values; DoT crits and Pandemic; the new-talent list; swapped sacrifices; Improved Shadow Bolt yours alone; pet buffs; Portal of Summoning; no Unstable Affliction or Haunt; races.
- **affliction.md**: builds (leveling 35 / 11 / 5, raid 35 / 7 / 9); every talent value; first four globals; rotation against the card; Levels 1 to 19 (Corruption 4, Life Tap 6, Agony 8, Drain Life 14, Improved Life Tap 10 to 11, Suppression 12 to 16, Improved Corruption 17 to 21, wands 5 and 13); Levels 20 to 60 (Siphon Life 30, Soul Siphon 36 to 38, Wrack 40, Soul Harvest 24, Succubus 20, Felhunter 30).
- **demonology.md**: leveling 30 Demonology points without Pact, raid 31 with Pact at 40; the raid build's swaps (gives up Improved Corruption, Improved Life Tap, Grim Reach, Soul Harvest, Demonic Energies, Improved Voidwalker); Decimation's rank 2 text (90%, 6%, 10 s, 40%, no shard; Decimation 29 and 31); setup order against the Pact-gated prepull; Levels 1 to 19 and 20 to 60 (Soul Link 30, Demonic Knowledge 32 to 34, Master Demonologist 35 to 39 with the Voidwalker's 10% physical, Soul Fire 48).
- **destruction.md**: leveling 35 / 0 / 16 and raid 33 / 0 / 18; the raid build's moves; opener and priority against the card; Levels 1 to 19 (Immolate 1, Destructive Reach 10 to 11, Bane 12 to 16, Molten Skin 17 to 19); Levels 20 to 60 rank rows (Shadow Bolt 1106 at 28 and 32, 1088 at 30; Conflagrate rank 3 is `spell_level` 40; Immolate 2941 at 40, 11665 at 41; Corruption 6223 at 41, 7648 at 42) and the post-44 Affliction points (Improved Life Tap 45 to 46, Suppression 47 to 51, Improved Corruption 52 to 56, then Grim Reach and Soul Harvest, which the line does not name and does not need to).
- **Voice:** no em or en dashes, no nightly numbers, no site or sim talk outside source URLs.

## Sources

- Client 1.60.1.70790: `talents/warlock.json` (Soul Harvest, Improved Drains, Decimation, Master Demonologist, Destructive Reach and every rank the pages quote), `spellconst/warlock.json` (705, 1088, 1106, 11661, 6353, 17924, 1316697, 699, 709, 7651, 11699, 11700, 6223, 7648, 2941, 11665, 5676, 6789, 17925, 17926, 1293817, 1293818, 17962, 18930 to 18932, 412758, 1293812, 1293813, 1120, 8288, 8289, 11675), `trainables/warlock.json`, `items/warlock.json` (11287, 11288), `addon-builds.json`, `addon-data.json` (`rotations/warlock-*`, every band).
- Curated: `data/curated/apl/warlock-{affliction,demonology,destruction}.json`.
- Earlier rounds: `round21/warlock.md`, `round20/warlock.md` (decoded orders, R20-4 scope), `round19/warlock.md`, `ledger.md` rounds 19 to 21.
- No web research was needed: the client data decides every call.

## Note sweep

Format: spec, row (action), bands it renders, each named talent, spell, rank or value with the level the player has it on the leveling and raid orders (L / R), then the verdict. Rows and bands re-derived from `addon-data.json` this round (identical to round 21).

**Affliction**
1. Life Tap (1454 to 11689), 10-60. "a point for every point of Spirit": every rank base plus Spirit x 1; Healthstone (Create Healthstone 10). OK.
2. Curse of the Elements (440892 to 1311680), 20-60. Five minutes at every rank; own slot; one per target. OK.
3. Corruption (172 to 11672), 10-60. "Instant with Improved Corruption": IC 17-21 L / R, conditional. OK (Should R18-S1 on 10 to 20).
4. Bane of Doom (603), 60. One global; big hit after a minute; Pandemic 31-33 L / 27-29 R. OK.
5. Amplify Curse (18288), 20-60. Talent 29 L / 26 R (earlier rows component lane); 50% on Agony; no global. OK.
6. Bane of Agony (980 to 11713), 10-60. "all fight before 60" (Doom 60); "biggest ticks come last" (every rank "builds up"); Amplify Curse "once you have" 29 L / 26 R. OK.
7. Siphon Life (18265 to 18881), 30-60. Talent 30 L / R; 30 s every rank. OK.
8. Immolate (707 to 11668), 10-60. Direct hit above one tick every rank; 2.0 s cast (1.5 s with Bane 56-60 L / 52-56 R) lands before the 3 s mark. OK.
9. Wrack (1316697), 40-60. Talent 40 L / R; Soul Siphon full 38 L / R; 10% to other Shadow DoTs; out-damages Shadow Bolt per second for under half the mana per second at every band. OK.
10. Shadow Bolt (695 to 11661), 10-60. "until Wrack at 40": 40 L / R. OK (renders past 40: component lane).
11. Shoot (5019), drawer only. Wand shot not during a global. OK.

**Demonology**
12. Soul Fire (6353, 17924), 48-60. Trained 48 / 56; Decimation 2 by 31 L / R (6 s, 40%, 10 s, no shard); "2.4 seconds once Bane is full": Bane 56-60 L / 51-55 R, conditional; raid skip gated on Pact 40 R; "never worth its shard" outside Decimation. OK.
13. Life Tap, 10-60. As row 1. OK.
14. Curse of the Elements, 20-60. As row 2. OK.
15. Corruption, 10-60. "A 2 second cast without Improved Corruption": raid never takes it; leveling 47-51. OK (Should, carried).
16. Bane of Doom, 60. One minute; "never lands" in the last minute. OK.
17. Bane of Agony, 10-60. First sentence scoped to 60; "new Doom" tail carried Should. OK.
18. Immolate, 10-60. As row 8 (Bane 56-60 L / 51-55 R). OK.
19. Shadow Bolt, 10-60. "beats Searing Pain" (Searing Pain 18, 10 to 16 folded in): WARLOCK-R20-4, routed. Master Demonologist with the Succubus, 10% Shadow: 35-39 L / R, conditional. Imp's sacrifice 15% "with Demonic Pact at 40 in the raid build": Pact 40 R, sacrifice 26 R. OK.
20. Shoot, drawer only. OK.

**Destruction**
21. Curse of the Elements, 20-60. As row 2. OK.
22. Corruption, 10-60. "until Improved Corruption makes it instant": IC 52-56 L / 50-54 R; 2 s every rank; one cast for the whole run. OK.
23. Life Tap, 10-60. As row 1. OK.
24. Bane of Doom, 60. OK.
25. Bane of Agony, 10-60. As row 17. OK.
26. Immolate, 10-60. Refresh at 3 s or less; Conflagrate 34 L / 32 R and Incinerate 40 L / R behind "once you have"; "until Shadow and Flame is full" 44 L / 42 R, and the condition recasts an eaten Immolate. OK (wording follows R19-3).
27. Conflagrate (1293817 to 18932), 25-60 (talent 34 L / 32 R; earlier rows component lane). Shadow and Flame 20 s, 4% a point to 20%, never eats Immolate at 5: 39 and 41-44 L / 37-39 and 41-42 R. OK.
28. Shadowburn (17877 to 18871), 20-60 (talent 30 L / 28 R). Any health; Fire buff and shard refund at 5 points as row 27. OK (shard below 5 points carried).
29. Incinerate (412758 to 1293813), 40-60. Talent 40 L / R; 25% with Immolate at every rank. OK.
30. Shadow Bolt (695 to 11661), 10-60. "before you learn Incinerate at 40": 40 L / R. OK (renders past 40: component lane).
31. Shoot, drawer only. OK.

# Guide accuracy round 21: Warlock

Audited on the round 21 audit tree (e039e1490: main after guide round 20; no sim round 21 commit in its log; engine pin d5c708834 in `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warlock/{index,affliction,demonology,destruction}.md`; every row note in `data/curated/apl/warlock-{affliction,demonology,destruction}.json`, sentence by sentence; every rendered row in `addon-data.json` `rotations/warlock-*`, band by band with its spell id; both shipped orders per spec in `addon-builds.json` (unchanged since round 20, so round 20's decoded orders stand); `talents/warlock.json` (every rank's text), `spelltips/warlock.json`, `spellconst/warlock.json` (coefficients, variance, points per level), `trainables/warlock.json`, `items/warlock.json` (wands). Read first: round 20 warlock report, the ledger's round 19 and round 20 sections, and the round 1, 2 and 21 briefs.

## Changes since round 20's tree (4234643f5)

Only round 20's warlock fix (bebce08c2) touched this class:
- `warlock-demonology.json` Shadow Bolt note: "...once Master Demonologist gives you more Shadow damage with the Succubus out it pulls further ahead. With Demonic Pact at 40 in the raid build, the Imp's sacrifice adds 15% more on top."
- `warlock-destruction.json` Immolate note: "...so it never runs out; once you have Conflagrate and Incinerate, both need it too, and until Shadow and Flame is full, recast it whenever Conflagrate eats it."
- `affliction.md`:115: "Drain Life while the DoTs run, and from 40 channel Wrack instead: once Soul Siphon is full at 38, three of your effects on the target add 36% to both, and Wrack costs far less mana than the Shadow Bolts it replaces."

All three are checked below and hold (one wording Should on the third).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARLOCK-R21-1 | CONFIRMED | leveling prose (quoted cost a talent in the shipped order changes; sim round 20 edit) | `destruction.md`:128 | "on level 30 gear your mana cannot carry the 160-mana rank 5 as well as the 110-mana rank 4" | 160 and 110 are the base costs of Shadow Bolt rank 5 (1106) and rank 4 (1088) (`spelltips`). The Destruction leveling order puts Cataclysm at 20 and 21, two points: "Reduces the Mana cost of your Destruction spells by 6%" (`talents/warlock.json`, Cataclysm rank 2). Shadow Bolt is a Destruction spell, so at 30 and 31 the player's tooltips read 150 (160 x 0.94 = 150.4) and 103 (110 x 0.94 = 103.4). Round 20's rule: a quoted cost is stated with the talents the build has at that level (the Benediction fixes). | "...on level 30 gear your mana cannot carry the 150-mana rank 5 as well as the 103-mana rank 4, so cast rank 4 at 30 and 31 and go back to rank 5 at 32." (or drop the numbers: "cannot carry rank 5 as well as the cheaper rank 4") |
| WARLOCK-R21-2 | CONFIRMED | leveling prose (comparison that fails at one level of its range) | `demonology.md`:124 | "a Lesser Magic Wand at 5 and a Greater Magic Wand at 13 do more than your bolts" | Lesser Magic Wand (11287): 12 to 22 over 1.5 s, 11.33 a second, required level 5; Greater Magic Wand (11288): 22 to 41 over 1.8 s, 17.5 a second, required level 13 (`items/warlock.json`). Shadow Bolt rank 3 (705) is trained at 12: 40 damage, 2.8 s (`spelltips`, `spellconst`), 14.3 a second, and the Demonology leveling order has no Bane or other Shadow Bolt talent before 56. So at level 12, with the Lesser wand the line names and the Greater one not yet wearable, your bolt does about a quarter more per second than the wand. From 5 to 11 the line holds (rank 1, 7.6 a second; rank 2 at 6 to 11, 23 plus 0.3 a level over 2.2 s, 10.5 to 11.1 a second), and from 13 the Greater wand is ahead of rank 3. | "Let the Voidwalker hold the mob and wand it down: a Lesser Magic Wand at 5 and a Greater Magic Wand at 13 cost no mana, and from 13 the wand out-damages your bolts. An enchanter makes both, or check the Auction House." |
| WARLOCK-R21-3 | CONFIRMED | leveling prose (talent named with no level on levels the order has not given it; round 20's R20-2 shape) | `affliction.md`:115 | "Cast Drain Soul on the last few percent of a mob's health whenever you need Soul Shards. Soul Harvest's mana regeneration then lets you drink less." | The paragraph covers 20 to 60 and gives levels for its other talents and pets ("Siphon Life from 30", "from 40 channel Wrack", "Soul Siphon is full at 38", "The Succubus comes at 20 and the Felhunter at 30"). Soul Harvest is the leveling order's 24 and 25 (`addon-builds.json`), so from 20 to 23 the player has no Soul Harvest. Round 20 D2 made the same shape CONFIRMED for Soul Siphon and Wrack in this paragraph; round 20 did not reach this sentence. Soul Harvest's text also needs the Drain Soul kill ("Killing a non-trivial target afflicted by your Drain Soul increases your Mana regeneration by 100% for 10 sec"), which the line says correctly. | "Cast Drain Soul on the last few percent of a mob's health whenever you need Soul Shards; from 24, Soul Harvest turns each of those kills into ten seconds of mana regeneration, so you drink less." |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. All three are page prose; no card note is wrong this round.

## Shoulds (wording, not findings)

- **affliction.md:115, Wrack against "the Shadow Bolts it replaces"** (round 20's new sentence). The sentence now has Wrack replacing Drain Life ("Drain Life while the DoTs run, and from 40 channel Wrack instead"), then says it is cheaper than "the Shadow Bolts it replaces"; the paragraph never names Shadow Bolt as the leveling filler. The claim itself is true (Wrack 200 mana for a 6 s channel against Shadow Bolt rank 6 at 210 per 3 s cast), and against Drain Life rank 4 (185 for 5 s) it would not be "far less". Suggested: "...three of your effects on the target add 36% to both, and Wrack does more than Shadow Bolt for about half the mana per second."
- **demonology.md:90, "Demonology spends more of them than the other specs: every wipe or dead Succubus is another summon".** Affliction and Destruction also keep a Succubus out and resummon after every wipe, so the stated reason applies to all three. The real difference is Master Summoner's 6 second summon, which makes a mid-fight resummon worth it only for Demonology. Suggested: "Demonology spends more of them than the other specs, since Master Summoner makes a mid-fight resummon worth it: every dead Succubus is another shard, so carry twice what you think you need."
- **Destruction Shadow Bolt row renders on every band from 40 to 60**, not 40 to 42 as round 20 wrote (`addon-data.json`: 44:11659 through 60:11661, in the round 20 tree too), beside Incinerate. The note ("Shadow Bolt before you learn Incinerate at 40") is true and the card gates the row on `spellIsKnown(412758)`; the rendering is the component lane's.
- **Carried**: R18-S1 (Affliction Corruption "Instant with Improved Corruption" on 10 to 20); the Demonology Corruption note written for the raid build (round 20 Should); the Demonology and Destruction Agony "new Doom" tail; the Destruction Shadowburn shard below Shadow and Flame 5; destruction.md:128 "Drain Life and Death Coil are your only heals" (Death Coil 42); the round 17 list.

## Decisions

- **D1. R21-1 is CONFIRMED.** Cataclysm's rank 2 text is a flat 6% off Destruction spells, the order gives two points at 20 and 21, and Shadow Bolt is a Destruction spell. The line is a sim round 20 edit and quotes the untalented tooltip.
- **D2. R21-2 is CONFIRMED.** Every number comes from the client: the wand item rows and Shadow Bolt's amount, variance-free base and cast time. The miss is one level, but it is the brief's "comparison that fails at some level" shape, and the fix costs nothing.
- **D3. R21-3 is CONFIRMED** for consistency with round 20 D2 (the same paragraph, the same shape, a 20 to 60 range that gives levels for everything else).
- **D4. Round 20's three fixes hold.** Demonology Shadow Bolt note: Master Demonologist (Succubus) is "Increases Shadow damage done by 10%", leveling and raid 35 to 39; the Imp's sacrifice is "Increases your Shadow damage by 15%"; Demonic Pact is the raid order's 40 and the card's prepull sacrifice is gated on it (`auraIsKnown 425464`); Searing Pain is Fire, so neither bonus reaches it. Destruction Immolate note: the card's `dotRemainingTime <= 3s` condition does recast an eaten Immolate, and Shadow and Flame is full at 44 L / 42 R. affliction.md:115: Soul Siphon 36 to 38, Wrack 40, Siphon Life 30 gives the third effect, and Soul Siphon's text names Drain Life and Wrack ("to both" holds).
- **D5. Superlatives and comparisons, every one on all four pages:**
  - "Doom does more damage per minute than Agony for one global instead of two and a half" (demonology.md:80): Doom 1742 plus 4.0 x spell power per minute; Agony rank 6 is 12 ticks of 46 at 0.133 each, 552 plus 1.6 x spell power per cast, so 2.5 casts a minute is 1380 plus 4.0 x spell power. Same coefficient per minute, Doom ahead on base. Holds.
  - Wrack "out-damages Shadow Bolt for a fraction of the mana" (affliction.md:57, :71, Wrack note): Wrack is 6 ticks at 0.286 (1.716 a channel) against Shadow Bolt's 0.857 per 3 s cast (1.714 per 6 s): equal coefficients; base 216 at 40 against 202 for two rank 6 bolts, before Soul Siphon's 36%. Mana per second 33 against 70 (148 with Bane at 60). Holds on every band from 40.
  - "Searing Pain is a threat spell ... and Incinerate does more" (destruction.md:84): rank 3 Incinerate 217 x 1.25 at 0.714 over 2.0 s against Searing Pain 114 at 0.429 over 1.5 s; at 40 to 49, rank 1 at 97 x 1.25 over 2.0 s (61 a second) against Searing Pain rank 4 at 62 over 1.5 s (41). Holds.
  - "Bane of Agony saves its biggest ticks for the end" / "its biggest ticks come last": every rank's text says it "builds up". Holds.
  - "The most Shadow bonuses of any warlock build at once" (demonology.md:32): raid Demonology stacks the sacrifice (15%), Master Demonologist (10%), Soul Link (3%) and Improved Shadow Bolt (20%); Affliction has Shadow Mastery (10%) and Wrack's 10% on DoTs, Destruction Shadow and Flame (20%). Holds.
  - "the fresh direct hit is worth the tick you give up" (Immolate, every rank): direct hit above one of five ticks at every rank (8 against 3 at rank 1, 158 against 55 at rank 8). Holds.
  - "the Succubus ... out-damage the Imp's Firebolt" (affliction.md:63, destruction.md:67): accepted in earlier rounds; no client value contradicts it.
  - "Orc fits Demonology best of any Horde race": with the sim lane (race re-runs, round 14); not re-filed.
  - "Affliction is the easiest warlock to level": opinion, not a value.
- **D6. Quoted costs and durations, every one on all four pages:** Curse of Recklessness 505 (rank 4, 56); Curse of the Elements 10% / 75 and five minutes (rank 4 at 50; five minutes at every rank); Amplify Curse 3 min, 30 s, 50% / 20%, gcd 0; Siphon Life 30 s every rank; Bane of Agony 24 s every rank; Doom 1 min; Wrack 6 s channel (`channeled: true`); Improved Corruption 0.4 s a point to instant at 21 (Affliction); Bane's 1.5 / 2.0 / 2.5 / 4 s; Decimation 2.4 s and 6 s (6 s x 0.6 with Bane; 60 s x 0.1); Soul Fire 4 s outside Decimation with Bane; summons 6 s with Master Summoner (10 s base); Improved Shadow Bolt 12 s; Shadow and Flame 20 s; Conflagrate 10 s; Shadowburn 15 s and 8 s; Corruption 18 s (rank 3 up); Portal of Summoning one hour (`cooldown_ms` 3600000); Firestone 2% crit (Greater and Major, 46 and 56); Spellstone shard. Only R21-1 is changed by a talent the build holds at that level.
- **D7. Not re-filed (routed or open):** WARLOCK-R20-4 (Demonology Shadow Bolt against Searing Pain, routed per the brief); R19-3 and R20-3 (Conflagrate on cooldown against held, and the Immolate wording that follows it); Affliction Amplify Curse below 60; the engine's missing Soul Shard; sim round 20 L1 items; race re-runs; BiS race captions; rows rendering before their talent (Amplify Curse 20 to 28, Shadowburn 20 to 29, Conflagrate 25 to 33).

## Sims requested

- None new. Carried with the sim lane: WARLOCK-R20-4 (Demonology filler at 18 to 21, 26, 34, 42, 50, 58); R19-3 (Destruction Conflagrate held against on cooldown at 34 to 43 L / 32 to 41 R), which also decides the R20-3 wording; Affliction Amplify Curse timing below 60; the Demonology Horde race re-run.

## Sim findings

- None new. Still open: the engine models no Soul Shard (Shadowburn 30 to 43, Soul Fire outside Decimation), and the sim round 20 L1 items.

## Engine questions

- None. Coefficients used above are client `spellconst` values; nothing in this audit depends on an engine behaviour the client contradicts.

## Checked and right this round

- **index.md**, every line: two slots for banes and curses; Curse of the Elements' schools and values at rank 4; DoT crits and Pandemic on Affliction's DoTs; the new-talent list against `talents/warlock.json` (Demonic Knowledge "up to 100% of your level"); the swapped sacrifices (Imp Shadow, Succubus Fire); Improved Shadow Bolt "from your attacks"; Portal of Summoning (caster plus two, two users, one hour); no Unstable Affliction or Haunt.
- **affliction.md**: both builds decode to 51 points (leveling 35 / 11 / 5, raid 35 / 7 / 9); every talent value and point count in the list; first four globals (Corruption, Doom, Siphon Life, Immolate; Amplify Curse has no global); rotation order against the card; Improved Bane of Agony 10%; Aftermath 40% at four; Bane 1.5 s; Life Tap and Improved Life Tap 20% (raid keeps both points); Levels 1 to 19 (Corruption 4, Life Tap 6, Agony 8, Drain Life 14, Improved Life Tap 10 to 11, Suppression 12 to 16, Improved Corruption 17 to 21; wands at required level 5 and 13).
- **demonology.md**: leveling 30 Demonology points, raid 31 with Pact the 31st; every swap between the builds; Demonic Embrace 9%, Demonic Aegis 15%, Improved Sayaad 20%; the setup order against the Pact-gated prepull; Decimation's text; Levels 1 to 19 (Health Funnel trained at 12, Improved Health Funnel 10 to 11, Demonic Embrace 12 to 14, Unholy Power 15 to 19); Levels 20 to 60 (Soul Link 30, Demonic Knowledge 32 to 34, Master Demonologist 35 to 39, Soul Fire 48, Decimation by 31).
- **destruction.md**: leveling 35 / 0 / 16 and raid 33 / 0 / 18 points; the raid build's moves (Pandemic 3 and Malevolence 3 from Grim Reach, Soul Harvest, Destructive Reach and Molten Skin, less Aftermath plus 2 and Cataclysm plus 1); Molten Skin 6%, Destructive Reach 20%; opener and priority against the card; Levels 1 to 19 (Destructive Reach 10 to 11, Bane 12 to 16, Molten Skin 17 to 19); Levels 20 to 60 rank rows (Shadow Bolt 1106 at 28 and 32, 1088 at 30; Immolate 2941 at 40, 11665 at 41; Corruption 6223 at 41, 7648 at 42).
- **Voice:** no em or en dashes, no nightly numbers, no site or sim talk outside source URLs.

## Sources

- Client 1.60.1.70790: `talents/warlock.json` (every rank of Cataclysm, Bane, Soul Siphon, Soul Harvest, Decimation, Shadow and Flame, Malevolence, Master Demonologist, Demonic Sacrifice, Demonic Pact, Demonic Knowledge, Improved Shadow Bolt, Aftermath, Molten Skin), `spelltips/warlock.json` (686 to 25307, 5676 to 17923, 980 to 11713, 603, 172 to 25311, 348 to 25309, 1316697, 18265 to 18881, 689 to 11700, 1120 to 11675, 412758 to 1293813, 1293817 to 18932, 17877 to 18871, 6353, 17924, 440892 to 1311680, 704 to 11717, 755, 712, 437169, 6366 to 17953, 2362 to 17728), `spellconst/warlock.json` (603, 11713, 1316697, 11661, 17924, 17923, 1293813, 11668, 18932, 5676, 695, 705, 1088), `trainables/warlock.json`, `items/warlock.json` (11287, 11288), `addon-builds.json`, `addon-data.json` (`rotations/warlock-*`, every band).
- Curated: `data/curated/apl/warlock-{affliction,demonology,destruction}.json`.
- Earlier rounds: `round20/warlock.md`, `round14/warlock.md` (race lines), `ledger.md` rounds 19 and 20.
- No web research was needed: the client data decides every call.

## Note sweep

Format: spec, row (action), bands it renders, each named talent, spell, rank or value with the level the player has it on the leveling order and the raid order (L / R), then the verdict.

**Affliction**
1. Life Tap (1454 to 11689), 10-60. "a point for every point of Spirit": every rank's formula is base plus Spirit x 1; Healthstone from 10. OK.
2. Curse of the Elements (440892 to 1311680), 20-60. Five minutes at ranks 1 to 4; separate slot from banes; one per target. OK.
3. Corruption, 10-60. "Instant with Improved Corruption": IC 17-21 L / R. Conditional. OK (Should R18-S1 on 10 to 20).
4. Bane of Doom (603), 60. One global; 1742 after 1 min; Pandemic 31-33 L / 27-29 R names Doom. OK.
5. Amplify Curse (18288), 20-60. Talent 29 L / 26 R (earlier rows component lane); 50% on Agony; gcd 0. OK.
6. Bane of Agony (980 to 11713), 10-60. "builds up" in every rank; Amplify Curse 29 L / 26 R named conditionally. OK.
7. Siphon Life (18265 to 18881), 30-60. Talent 30 L / R; 30 s every rank. OK.
8. Immolate (707 to 11668), 10-60. Direct hit above one tick at every rank; lands before expiry at 2.0 s (1.5 s with Bane 56-60 L / 52-56 R). OK.
9. Wrack (1316697), 40-60. Talent 40 L / R; Soul Siphon full 38 L / R; 10% to other Shadow DoTs; equal coefficients to Shadow Bolt per second, higher base, half the mana per second or less. OK.
10. Shadow Bolt (695 to 11661), 10-60. "until Wrack at 40": 40 L / R. OK.
11. Shoot (5019), drawer only. Wand on the global. OK.

**Demonology**
12. Soul Fire (6353, 17924), 48-60. Trained 48 / 56; Decimation 2 by 31 L / R (6 s, 40%, 10 s, no shard); "2.4 seconds once Bane is full": Bane 56-60 L / 51-55 R; raid skip gated on Pact 40 R. OK.
13. Life Tap, 10-60. As row 1. OK.
14. Curse of the Elements, 20-60. As row 2. OK.
15. Corruption, 10-60. "A 2 second cast without Improved Corruption": raid never takes it; leveling 47-51. Conditional. OK (Should, carried).
16. Bane of Doom, 60. 1 min; "never lands" in the last minute. OK.
17. Bane of Agony, 10-60. First sentence scoped to 60; "new Doom" tail carried Should. OK.
18. Immolate, 10-60. As row 8 (Bane 56-60 L / 51-55 R). OK.
19. Shadow Bolt, 10-60. "beats Searing Pain" (Searing Pain 18+): WARLOCK-R20-4, routed. Master Demonologist with the Succubus, Shadow 10%: 35-39 L / R, conditional on the Succubus. Imp's sacrifice 15% Shadow "with Demonic Pact at 40 in the raid build": Pact 40 R, sacrifice 26 R. OK.
20. Shoot, drawer only. OK.

**Destruction**
21. Curse of the Elements, 20-60. As row 2. OK.
22. Corruption, 10-60. "until Improved Corruption makes it instant": IC 52-56 L / 50-54 R; 2 s every rank. OK.
23. Life Tap, 10-60. OK.
24. Bane of Doom, 60. OK.
25. Bane of Agony, 10-60. As row 17. OK.
26. Immolate, 10-60. "never runs out" (refresh at 3 s or less); Conflagrate 34 L / 32 R, Incinerate 40 L / R; "until Shadow and Flame is full" 44 L / 42 R, and the row's condition recasts an eaten Immolate. OK (wording follows R19-3 when it lands).
27. Conflagrate (1293817 to 18932), 25-60 (talent 34 L / 32 R; earlier rows component lane). Shadow and Flame 20 s, 4% a point to 20%, never eats Immolate at 5: 39 and 41-44 L / 37-39 and 41-42 R. OK.
28. Shadowburn (17877 to 18871), 20-60 (talent 30 L / 28 R). Any health in every rank's text; Fire buff and shard refund at 5 points as row 27. OK (shard below 5 points carried).
29. Incinerate (412758 to 1293813), 40-60. Talent 40 L / R; 25% with Immolate in every rank. OK.
30. Shadow Bolt (695 to 11661), 10-60 (renders past 40 beside Incinerate: component-lane Should). "before you learn Incinerate at 40": 40 L / R. OK.
31. Shoot, drawer only. OK.

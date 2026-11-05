# Guide accuracy round 20: Warlock

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20; engine pin d5c708834), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warlock/{index,affliction,demonology,destruction}.md`; every row note in `data/curated/apl/warlock-{affliction,demonology,destruction}.json`, read sentence by sentence; every rendered row in `addon-data.json` `rotations/warlock-*`, band by band (Affliction 23 bands, Demonology 25, Destruction 27, with the spell id on each band); both shipped orders per spec in `addon-builds.json`, decoded by tier and column against `talents/warlock.json`; `spelltips/warlock.json`, `spellconst/warlock.json`, `trainables/warlock.json`. Read first: round 19 warlock report, the ledger's round 19 section, the sim round 20 warlock report and `integration.md`, and the round 1 and 2 briefs.

## Changes since round 19's tree (6beb919c2)

- `warlock-destruction.json` Immolate note (round 19's R19-1 fix): "It lands a moment before the old one ends, so it never drops; once you have Conflagrate and Incinerate, both need it too."
- All four pages: the wago.tools source names build 70790 (R19-2 fix).
- destruction.md:128 (sim round 20 X1): "Shadow Bolt rank 5 arrives at 28 and is worth casting right away. Shadowburn joins at 30, and on level 30 gear your mana cannot carry the 160-mana rank 5 as well as the 110-mana rank 4, so cast rank 4 at 30 and 31 and go back to rank 5 at 32." The `rank-choices.json` row renders as expected: Destruction Shadow Bolt is 1088 (rank 4) at band 30 and 1106 (rank 5) at 28 and 32.
- `addon-builds.json`: the warlock orders have not changed (the diff touches only mage and shaman).

## Shipped orders (index 0 = level 10)

- Affliction leveling: Improved Life Tap 10-11, Suppression 12-16, Improved Corruption 17-21, Grim Reach 22-23, Soul Harvest 24-25, Improved Drains 26-28, Amplify Curse 29, Siphon Life 30, Pandemic 31-33, Malevolence 34-35, Soul Siphon 36-38, Shadow Mastery 39 and 41-44, Wrack 40, Improved Health Funnel 45-46, Unholy Power 47-51, Demonic Aegis 52-53, Improved Voidwalker 54-55, Bane 56-60. Raid: the same to 23, then Improved Bane of Agony 24-25, Amplify Curse 26, Pandemic 27-29, Siphon Life 30, Malevolence 31-35, Soul Siphon 36-38, Shadow Mastery 39 and 41-44, Wrack 40, Unholy Power 45-49, Demonic Aegis 50-51, Bane 52-56, Aftermath 57-60.
- Demonology leveling: Improved Health Funnel 10-11, Demonic Embrace 12-14, Unholy Power 15-19, Demonic Aegis 20, Improved Voidwalker 21, Fel Vitality 22-24, Demonic Energies 25, Demonic Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, Improved Life Tap 40-41, Suppression 42-46, Improved Corruption 47-51, Grim Reach 52-53, Soul Harvest 54-55, Bane 56-60. **No Demonic Pact.** Raid: Fel Vitality 21-23, Improved Sayaad 24-25, Demonic Sacrifice 26, Master Summoner 27-28, Decimation 29 and 31, Soul Link 30, Demonic Knowledge 32-34, Master Demonologist 35-39, **Demonic Pact 40**, Suppression 41-45, Improved Shadow Bolt 46-50, Bane 51-55, Ruin 56-60.
- Destruction leveling: Destructive Reach 10-11, Bane 12-16, Molten Skin 17-19, Cataclysm 20-21, Aftermath 22-24, Ruin 25-29, Shadowburn 30, Agonizing Flames 31-33, Conflagrate 34, Bane of Havoc 35, Fire and Brimstone 36-38, Shadow and Flame 39 and 41-44, Incinerate 40, Improved Life Tap 45-46, Suppression 47-51, Improved Corruption 52-56, Grim Reach 57-58, Soul Harvest 59-60. Raid: Bane 10-14, Cataclysm 15-17, Aftermath 18-22, Ruin 23-27, Shadowburn 28, Agonizing Flames 29-31, Conflagrate 32, Bane of Havoc 33, Fire and Brimstone 34-36, Shadow and Flame 37-39 and 41-42, Incinerate 40, Improved Life Tap 43-44, Suppression 45-49, Improved Corruption 50-54, Pandemic 55-57, Malevolence 58-60.

Trainer levels (`trainables/warlock.json`): Immolate 1, Corruption 4 (ranks at 14/24/34/44/54/60), Life Tap 6, Bane of Agony 8, Create Healthstone 10, Searing Pain 18 / 26 / 34 / 42 / 50 / 58, Curse of the Elements 20 / 30 / 40 / 50, Shadow Bolt 1 / 6 / 12 / 20 / 28 / 36 / 44 / 52 / 60, Death Coil 42, Soul Fire 48 / 56, Bane of Doom 60.

## Findings

Rendering follows round 18 D2: the guide card and the addon row print a note's first sentence on every band where the row renders. The whole note shows in the sim page drawer and on /sim/specs, at whatever level the sim is set to.

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARLOCK-R20-1 | CONFIRMED | card note (credits two buffs at once that the leveling order never lets you hold together, and contradicts the page) | `data/curated/apl/warlock-demonology.json` Shadow Bolt notes (priorityList[8]); the row renders 10 to 60 | "...and once you have the Imp's sacrifice and Master Demonologist it pulls further ahead." | Demonic Sacrifice (`talents/warlock.json`) says "The effect is canceled if any Demon is summoned", and Master Demonologist works only "as long as that demon is active". Only Demonic Pact lets you keep both ("Your Demonic Sacrifice effect is no longer cancelled by summoning a different Demon pet"). The leveling order never takes Pact; the raid order takes it at 40. So on every leveling band from 35 (Master Demonologist) to 60, and on raid bands 35 to 39, the player cannot have the sacrifice and Master Demonologist at the same time. The card agrees: the prepull Imp summon and sacrifice are gated on `auraIsKnown 425464` (Demonic Pact), so below Pact the card never sacrifices. demonology.md:128 says the same: "The leveling build skips Demonic Pact, so any summon cancels your sacrifice. Keep a demon out instead." The note's only reason for "pulls further ahead" is a combination the leveling order never gives. Round 19 passed it ("sacrifice 26, Master Demonologist 35 to 39 on both orders") without checking Pact, so its pass is wrong. | "Shadow Bolt for every other cast. It beats Searing Pain, which also carries far too much threat, and once Master Demonologist gives you more Shadow damage with the Succubus out it pulls further ahead. With Demonic Pact at 40 in the raid build, the Imp's sacrifice adds 15% more on top." (The first clause waits on the Searing Pain sim in R20-4.) |
| WARLOCK-R20-2 | CONFIRMED | leveling prose (quotes a talent's full value, and a spell, with no level, on levels where the order has not given them) | `affliction.md`:115 | "Drain Life or Wrack while the DoTs run: with three of your effects on the target, Soul Siphon adds 36%, and Wrack costs far less mana than the Shadow Bolts it replaces." | This sentence opens the "Levels 20 to 60" paragraph, and the same paragraph gives levels for everything else ("Siphon Life from 30", "Succubus comes at 20 and the Felhunter at 30"). Soul Siphon is the leveling order's 36 to 38, so it is 12% a point and 36% only from 38. Wrack is the order's 40 (`addon-builds.json`; `trainables` Wrack level 40). A third Affliction effect needs Siphon Life at 30. So from 20 to 35 Soul Siphon does not exist, and from 20 to 39 Wrack does not exist. This is the level-60 figure on lower levels that the brief rules out. Round 19 passed this line ("affliction.md:109 and :115"), so its pass is wrong. | "Drain Life while the DoTs run, and from 40 channel Wrack instead: once Soul Siphon is full at 38, three of your effects on the target add 36% to both, and Wrack costs far less mana than the Shadow Bolts it replaces." |
| WARLOCK-R20-3 | LIKELY | card note (round 19 fix; its first clause is false on the levels where Conflagrate eats Immolate) | `warlock-destruction.json` Immolate notes (priorityList[6]); the row renders 10 to 60 | "It lands a moment before the old one ends, so it never drops; once you have Conflagrate and Incinerate, both need it too." | Conflagrate "consum[es] your Immolate effect" (1293818, 17962, 18932). Shadow and Flame keeps Immolate only 20% a point (ranks 1 to 4 keep it 20 to 80% of the time), and the engine consumes it on the rest (`conflagrate.go`). The card fires Conflagrate on cooldown while Immolate is up. So on leveling bands 34 to 43 (raid 32 to 41), Immolate drops after most Conflagrates no matter when you refresh it. The page says so itself (destruction.md:128: "Until Shadow and Flame reaches 5 points, Conflagrate can eat your Immolate"). The refresh timing stops Immolate from expiring, which is what "never drops" means to a careful reader, but on those levels the note still promises something the next row breaks. Filed LIKELY because "drops" can be read as "expires". | "Immolate with 3 seconds or less left on it. It lands a moment before the old one ends, so it never runs out; once you have Conflagrate and Incinerate, both need it too, and until Shadow and Flame is full, recast it whenever Conflagrate eats it." (Adjust to whatever the R19-3 Conflagrate measurement decides.) |
| WARLOCK-R20-4 | NEEDS-SIM | card note (comparison on bands where a new Searing Pain rank meets an older Shadow Bolt rank, with no Shadow bonus on the leveling build) | `warlock-demonology.json` Shadow Bolt notes (priorityList[8]); renders 10 to 60 | "It beats Searing Pain, which also carries far too much threat" | The client coefficients per second of cast are the same for both spells: Shadow Bolt 0.857 over 3.0 s and Searing Pain 0.429 over 1.5 s (`spellconst`). So per second of casting, the comparison comes down to base damage, Shadow Bolt's base over 3 s against Searing Pain's over 1.5 s. On the levels where Searing Pain has just gained a rank and Shadow Bolt has not, Searing Pain is ahead on base damage per second: at 18 to 21 the rendered row is Shadow Bolt rank 3 (705, 40 over 2.8 s) against Searing Pain rank 1 (5676, 23 over 1.5 s), 14.3 to 15.3 a second; at 26 to 27, rank 4 (58) against rank 2 (33), 19.3 to 22.0; at 34 to 35, 26.0 to 29.3; at 42 to 43, 33.7 to 41.3; at 50 to 51, 47.0 to 56.7; at 58 to 59, 63.7 to 76.0. The leveling build has no Shadow bonus before Master Demonologist at 35 (R20-1: no sacrifice without Pact). Shadow Bolt is ahead per point of mana on every band (for example 40/70 against 23/45 at 18), and the band 20 rank row (rank 3 over rank 4) shows those bands run short of mana. So "beats" may hold in the sim. The APL history says Searing Pain once out-measured Shadow Bolt as the filler on a non-Pact build (2026-10-08). Which way it goes is a measurement. | If Shadow Bolt wins on every band: no change. If Searing Pain wins on some bands: "Shadow Bolt for every other cast once your Shadow bonuses are up; Searing Pain does more per second at <levels> but costs more mana per hit and carries far too much threat." |

Counts: CONFIRMED 2, LIKELY 1, NEEDS-SIM 1, NEEDS-OWNER 0. Page prose: 1 confirmed error (R20-2). The rest are card notes.

## Shoulds (wording, not findings)

- **Demonology Corruption note** (priorityList[4], renders 10 to 60): "A 2 second cast without Improved Corruption, and still worth it." This is written for the raid build. The leveling order takes Improved Corruption at 47 to 51, so from 52 a leveling Demonology Corruption is instant. The sentence is conditional and true, the same shape as R18-S1 (round 19 D4). Suggested: "Corruption whenever it has fallen off. A 2 second cast until Improved Corruption makes it instant (the raid build skips it), and still worth it."
- **Destruction Shadow Bolt row at 40, 41 and 42**: the row and its note ("Shadow Bolt before you learn Incinerate at 40") render on bands where Incinerate is already known. The note itself is true. This is a rendering question for the component lane, not a note error.
- **Carried**: R18-S1 (Affliction Corruption "Instant with Improved Corruption" on 10 to 20), the Demonology and Destruction Agony "new Doom" tail (round 19 Should), the Demonology Shadow Bolt Searing Pain comparison on 10 to 16 (now folded into R20-4), the Destruction Shadowburn note's shard cost below Shadow and Flame 5 (round 19 Should and Sim finding), destruction.md:128 "Drain Life and Death Coil are your only heals" (Death Coil 42), and the round 17 list.

## Decisions

- **D1. R20-1 is CONFIRMED.** The client text decides it, not an inference: "The effect is canceled if any Demon is summoned", Pact's text is the only exception, and the leveling order has no Pact. The page's own line 128 and the card's Pact-gated prepull agree.
- **D2. R20-2 is CONFIRMED** under this round's rule for prose ("against the shipped talent order and the trainer levels"). The paragraph covers 20 to 60, gives levels for its other spells, and gives none for Soul Siphon or Wrack.
- **D3. R20-3 is LIKELY.** It is the same mechanic as R19-3, which is already routed. This finding concerns the round 19 note wording only, and its fix should follow the R19-3 result.
- **D4. R20-4 is NEEDS-SIM.** Per second of casting, Searing Pain is ahead on base damage at the rank-up levels; per point of mana, Shadow Bolt is ahead on every band. The client data cannot decide which matters more on a mana-short leveling fight.
- **D5. Every value quoted in a note matches the client at every rank and on both orders.** Checked: Amplify Curse 50%, no global (`spellconst` gcd 0); Wrack 10% and a 6 s channel; Siphon Life 30 s; Curse of the Elements 5 min (rank 1 to 4); Soul Fire 90% and 40% from Decimation 2 at 31 on both orders; Bane 2.4 s (6 s minus 2, times 0.6); Shadow and Flame 20 s at 4% a point; Incinerate 25%; Doom 1 min. No "biggest hit" or "best per mana" claim fails on a lower rank, apart from R20-4.
- **D6. Not re-filed (routed or open):** R19-3 Conflagrate on cooldown against held (ledger round 19 "Sims routed"; sim round 20 did not report it); Affliction Amplify Curse below 60 (routed); the engine models no Soul Shard (routed); sim round 20 L1 log items; BiS race captions; rows rendering before their talent (Amplify Curse 20 to 28, Shadowburn 20 to 29, Conflagrate 25 to 33).

## Sims requested

- **Demonology filler, Searing Pain against Shadow Bolt** (R20-4): leveling build, bands 18, 20, 26, 34, 42, 50 and 58, seeds 1 and 50001, the ladder preset. Run the card as it is against Shadow Bolt replaced by Searing Pain at each band's top rank. Report DPS and seconds out of mana.
- **Carried, not re-filed:** R19-3 (Destruction Conflagrate on cooldown against held until Immolate has 4 s or less, bands 34 to 43 leveling and 32 to 41 raid order); it also decides the R20-3 wording. Affliction Amplify Curse saved for the last minute against used before each Agony, bands 26 to 58.

## Sim findings

- None new. Still open: the engine models no Soul Shard (Shadowburn 30 to 43, Soul Fire outside Decimation), plus the sim round 20 L1 items (Eureka! on Incinerate, demon crit at 60, Soul Siphon's count, Wrack and Doom at 60, the level-50 Succubus weapon).

## Engine questions

- None. The engine matches the client on every point this audit touched: Amplify Curse has no global in `curses.go` and gcd 0 in `spellconst`; the Demonology prepull sacrifice is gated on Pact; Conflagrate consumes Immolate except on Shadow and Flame's chance to keep it.

## Checked and right this round

- **Round 19's fixes:** the 70790 source links on all four pages. The Immolate note's second clause ("once you have Conflagrate and Incinerate, both need it too") is true on both orders (Conflagrate 34 / 32, Incinerate 40); for the first clause see R20-3.
- **Sim round 20's guide edit** (destruction.md:128): Shadow Bolt rank 5 (1106) is trained at 28 and renders at band 28; Shadowburn is the leveling order's 30; costs are 160 for rank 5 and 110 for rank 4 (`spelltips`); the row renders rank 4 at 30 and rank 5 at 32. The rest of the paragraph is unchanged: Immolate 2941 at 40 and 11665 at 41, Corruption 6223 at 41 and 7648 at 42, Conflagrate rank 3 (17962) and Immolate rank 5 (11665) trained at 40.
- **Page prose, every line on all four pages**, against the decoded builds, `talents`, `spelltips` and `trainables`. Affliction: build totals (35 Affliction points in the raid build), Amplify Curse, Pandemic, Malevolence, the raid build's swaps, Aftermath 40%, Bane 1.5 s, Curse of Recklessness 505, Intellect 60 a point of crit, levels 4 / 6 / 8 / 10 / 14, Improved Corruption 17 to 21. Demonology: 30 points before Pact, Demonic Embrace 9% (3 of 5), Demonic Aegis 15% (1 point), Improved Sayaad 20%, Master Summoner 6 s, Decimation 6% / 10 s / 2.4 s / 6 s, Soul Fire 4 s outside Decimation, the raid build's swaps, leveling 30 / 32 / 35 / 39 / 48. Destruction: the raid and leveling Destruction and Affliction point moves, Bane's four cast times, Aftermath 30% / 50%, Molten Skin 6%, Destructive Reach 20%, Pandemic and Malevolence 3 each. Index: Haunt and Unstable Affliction are `class_spell` rows only.
- **Voice:** no em or en dashes, no nightly numbers, no site or sim talk outside source URLs.

## Sources

- Client 1.60.1.70790: `talents/warlock.json` (Demonic Sacrifice, Demonic Pact, Master Demonologist, Shadow and Flame, Decimation, Bane, Soul Siphon, Pandemic, Amplify Curse, Improved Corruption, Demonic Embrace, Demonic Aegis, Improved Sayaad, Master Summoner, Molten Skin, Aftermath), `spelltips/warlock.json` (686 to 11661, 5676 to 17923, 18288, 603, 980, 11713, 707, 1094, 2941, 11668, 1316697, 18265, 18881, 17877, 18867, 18871, 1293817, 1293818, 17962, 18932, 412758, 1293813, 6353, 17924, 1454, 11689, 172, 11672, 440892, 1311680, 11717), `spellconst/warlock.json` (gcd and coefficients for 18288, 705, 1088, 5676, 17919, 11661, 1316697), `trainables/warlock.json`, `addon-builds.json`, `addon-data.json` (`rotations/warlock-*`, every band).
- Curated: `data/curated/apl/warlock-{affliction,demonology,destruction}.json` (the notes, the conditions, the Demonology top-level `notes` history on Searing Pain and the Pact-gated prepull).
- Earlier rounds: `round19/warlock.md`, `ledger.md` round 19, `design/reviews/2026-11-04-sim-audit-round20/warlock.md` and `integration.md`.
- No web research was needed: the client text decides every call, and R20-4 is a measurement.

## Note sweep

Format: spec, row (action), bands it renders, each named talent, spell, rank or value with the level the player has it on the leveling order and the raid order (L / R), then the verdict.

**Affliction**
1. Life Tap (11689 at 60, ranks from 1454), 10-60. Spirit term in every rank's text; Healthstone at 10. OK.
2. Curse of the Elements (1311680), 20-60. 5 min at ranks 1 to 4 (20 / 30 / 40 / 50); separate slot from banes; no stacking. OK.
3. Corruption, 10-60. "Instant with Improved Corruption": IC 17-21 L / R, instant from 21; conditional. OK (Should R18-S1, 10 to 20).
4. Bane of Doom (603), 60. Trained at 60; 1 min; one global; Pandemic 31-33 L / 27-29 R names Doom. OK.
5. Amplify Curse (18288), 20-60. Talent 29 L / 26 R (the row before that is component lane); 50% on Agony, gcd 0, 3 min. OK.
6. Bane of Agony, 10-60. Agony ranks 8 / 18 / 28 / 38 / 48 / 58 "builds up"; "once you have Amplify Curse" 29 L / 26 R. OK.
7. Siphon Life, 30-60. Talent 30 L / R; 30 s at every rank. OK.
8. Immolate, 10-60. 2.0 s at every rank (1.5 s with Bane 56-60 L / 52-56 R), lands before expiry; direct hit larger than one tick at every rank (17 against 6 at 707, 146 against 52 at 11668). OK.
9. Wrack (1316697), 40-60. Talent 40 L / R; Soul Siphon full by 38 L / R; 10% to other Shadow DoTs; beats Shadow Bolt per time at 40, 44, 52 and 60 with Soul Siphon, for 200 mana a channel. OK.
10. Shadow Bolt, 10-60. "until Wrack at 40": 40 L / R. OK.
11. Shoot (5019), not on the card (drawer only). Wand on the global. OK.

**Demonology**
12. Soul Fire (6353, 17924), 48-60. Trained 48 / 56; Decimation 2 by 31 L / R (90%, so 6 s; 40%; 10 s; no shard); "2.4 seconds once Bane is full", Bane 56-60 L / 51-55 R; "In the raid build, skip it": raid Shadow bonuses (sacrifice with Pact 40, Master Demonologist 39, Improved Shadow Bolt 46-50). OK.
13. Life Tap, 10-60. As row 1. OK.
14. Curse of the Elements, 20-60. As row 2. OK.
15. Corruption, 10-60. "without Improved Corruption": raid never takes it; leveling 47-51. Conditional. OK (Should).
16. Bane of Doom, 60. 1 min, "never lands" in the last minute. OK.
17. Bane of Agony, 10-60. First sentence scoped to 60; the "new Doom" tail is a Should (round 19). OK.
18. Immolate, 10-60. As row 8 (Bane 56-60 L / 51-55 R). OK.
19. Shadow Bolt, 10-60. "beats Searing Pain" (Searing Pain 18+): **WARLOCK-R20-4**. "the Imp's sacrifice and Master Demonologist" (sacrifice 26, MD 35-39, together only with Pact, which is 40 R and never in L): **WARLOCK-R20-1**.
20. Shoot, drawer only. OK.

**Destruction**
21. Curse of the Elements, 20-60. As row 2. OK.
22. Corruption, 10-60. "until Improved Corruption makes it instant": IC 52-56 L / 50-54 R; 2 s at every rank. OK.
23. Life Tap, 10-60. OK.
24. Bane of Doom, 60. OK.
25. Bane of Agony, 10-60. As row 17. OK.
26. Immolate, 10-60. "never drops" on 34-43 L / 32-41 R while Conflagrate eats it: **WARLOCK-R20-3**. "once you have Conflagrate and Incinerate": 34 L / 32 R and 40. OK.
27. Conflagrate, 25-60 (talent 34 L / 32 R; the row before that is component lane). First sentence against destruction.md:128: R19-3 (routed). Shadow and Flame 20 s, 4% a point, kept 20% a point: 39 and 41-44 L / 37-39 and 41-42 R. OK.
28. Shadowburn, 20-60 (talent 30 L / 28 R). No health limit in any rank's text; Shadow and Flame Fire buff and shard refund as row 27. OK (Should and Sim finding on the shard below 44 L / 42 R).
29. Incinerate, 40-60. Talent 40 L / R; 25% in every rank's text. OK.
30. Shadow Bolt, 10-42. "before you learn Incinerate at 40": 40 L / R; the card gates it on `spellIsKnown(412758)`. OK (the rows at 40 to 42 are a component-lane Should).
31. Shoot, drawer only. OK.

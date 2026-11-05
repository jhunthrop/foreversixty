# Guide accuracy round 24: Warlock

Audited on the round 24 audit tree (ea9adea1c: main after guide round 23, which merged sim round 21 at 617d4e9c4; engine pin 170a59c39 in `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/warlock/{index,affliction,demonology,destruction}.md`, line by line; every row note in `data/curated/apl/warlock-{affliction,demonology,destruction}.json`, sentence by sentence; every rendered row and clause in `addon-data.json` `rotations/warlock-*` (Affliction 23 bands, Demonology 25, Destruction 27); all six orders in `addon-builds.json` decoded level by level; `talents/warlock.json`, `spellconst/warlock.json`, `spelltips/warlock.json`; `web/src/lib/guides/rotation-lists.ts` (`clauseOf`). Read first: round 23 warlock report, the ledger's round 19 to 23 sections, `design/reviews/2026-11-05-sim-audit-round21/warlock.md` and `integration.md`.

## New work: sim round 21 and round 23 (git diff 24ee457d0..HEAD)

Sim round 21 (8c345ae68, e61642ae0, 648d87244, 2784c90b0):
- Destruction card: Conflagrate split into a held row (Immolate at 6 s or less, before `auraIsKnown 426316`) and an on-cooldown row (after); new notes on both, on Shadowburn (shard and refund) and on Immolate. destruction.md :58 (Shadow and Flame per point), :76 (priority 4), :128 (Conflagrate from 34 and 39; Shadowburn on a dungeon boss).
- Affliction card: Amplify Curse fires before each Agony when Bane of Doom is not known, in the last minute when it is. New note. affliction.md :115 (from 29).

Round 23 fixes (3221fddef): affliction.md :32, :57, :71, :85 and the Wrack note (R23-1, R23-2); demonology.md :110 (R23-3). R23-4 is routed (ledger round 23).

Checked against the client: Shadow and Flame text (426316: 4% a point for 20 s, 20% a point Conflagrate keeps Immolate, 20% a point Shadowburn refunds its shard) matches :58 and both notes; Conflagrate "consuming your Immolate effect" every rank; Amplify Curse 18288 `gcd_ms` 0, 180 s cooldown, 30 s; orders: Conflagrate 34 L / 32 R, Shadow and Flame 39, 41 to 44 L / 37 to 39, 41, 42 R, Incinerate 40, Amplify Curse 29 L / 26 R, Bane of Doom 60. Round 23's fixes hold (D1).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| WARLOCK-R24-1 | CONFIRMED | rendered clause false at 60; contradicts the page | `warlock-affliction.json` priority line 5 (Amplify Curse) first sentence, as rendered by `clauseOf` on the guide's level 60 lists and as the addon's band-60 `condition`; against affliction.md :52, :67 | Rendered at 60: "Amplify Curse before each Bane of Agony, whenever it is off cooldown." | The row's condition is `not dot 603 and not dot 11713 and (not spellIsKnown 603 or remainingTime < 60s)`: at 60 Doom is known, so it fires only in the last minute. The note's second sentence carries the 60 rule, but `clauseOf` (rotation-lists.ts:81) keeps only the first sentence, and `addon-data.json` band 60 shows exactly that sentence in game. So the guide's level 60 list and the addon at 60 tell the player to press it on cooldown, while the same page says "In a raid it goes on the Bane of Agony you put up in a boss's last minute" (:52) and "once the last one has fired, press Amplify Curse" (:67). The integrator fixed the identical shape on Destruction's Conflagrate this round (one sentence, no semicolon) but not here. | Make the first sentence carry both: "Amplify Curse before each Bane of Agony whenever it is ready below 60, and at 60 only before the final Agony, once your last Doom has fired. It makes that Agony 50% stronger and costs no global." (No semicolon, under 160 characters after the name.) Regenerate addon-data.json and Data.lua. |
| WARLOCK-R24-2 | CONFIRMED | measured gain explained by the wrong mechanism | destruction.md :128; `warlock-destruction.json` priority line 7 (Conflagrate on cooldown), second sentence (renders 25 to 60; applies 39 to 60 L / 37 to 60 R) | :128 "From 39, with the first point of Shadow and Flame, cast it on cooldown: the Shadow damage it gives you for 20 seconds is worth more than the burn it may eat." Note: "Each hit then gives you 20 seconds of more Shadow damage, 4% a point, which is worth more than the burn it may eat" | The comparison the line answers is held (Immolate at 6 s or less) against on cooldown. Under the hold, Conflagrate still fires at least once per Immolate: Immolate lasts 15 s (11668 and every rank, `duration_ms` 15000), so the held cycle is about 12 s after an eaten Immolate (1.5 s global, 1.5 s Immolate with Bane, then 9 s to the 6 s mark; Conflagrate's 10 s cooldown is already back) and about 13.5 s after a kept one. The Shadow buff lasts 20 s, so it never drops under either rule: the 4% a point is the same in both and cannot be what the sim measured. What the first point changes is Conflagrate's 20% chance a point to keep Immolate (426316), which makes the extra on-cooldown Conflagrates (one every 10 s against one every 12 to 13.5 s) cost less burn; at 38, with no point, the two rules tie (sim round 21 S1: -0.1 to +0.1%), and at 39 on cooldown leads by 1.0%. From 40 the filler is Incinerate (Fire), so the Shadow buff matters even less. | :128 "From 39 cast it on cooldown: each point of Shadow and Flame gives it a 20% chance to keep your Immolate, so the extra Conflagrates cost less of the burn than they deal." Note: "...then on cooldown while Immolate is up. Each point of Shadow and Flame gives it a 20% chance to keep your Immolate, so the extra hits are worth the burn they may eat, and at five points it never eats it." The Shadow buff can stay as a separate clause if wanted, but not as the reason. |
| WARLOCK-R24-3 | CONFIRMED | class index teaches a rotation the cards no longer use | `index.md` :19 | "Every spec plays out of the same kit: Corruption, a bane, Immolate and Shadow Bolt, with your talents deciding what fills the gaps." | From 40 the Affliction card replaces Shadow Bolt with Wrack ("until Wrack at 40 takes over as your filler") and the Destruction card with Incinerate ("before you learn Incinerate at 40"); at 60 neither card casts Shadow Bolt, and affliction.md :27 and destruction.md :28 describe their rotations without it. Only Demonology plays Shadow Bolt at 60. So "every spec plays out of ... Shadow Bolt" is false for two of three specs on every band from 40, and the "gaps" Shadow Bolt is said to share with talents are exactly the casts the talents take over. | "Every spec plays out of the same kit: Corruption, a bane and Immolate, with Shadow Bolt as the filler until your talents replace it: Wrack for Affliction and Incinerate for Destruction from 40." |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. Two card notes (R24-1, R24-2; note text only, then regenerate `addon-data.json` and Data.lua), the rest page prose.

## Shoulds (wording, not findings)

New:
- **affliction.md :115, "press Amplify Curse before every Bane of Agony as soon as it is ready".** With a 3 minute cooldown and a 24 second Agony, "before every" is true only through "as soon as it is ready"; "before your next Bane of Agony whenever it is ready" reads exactly.
- **destruction.md :58, "a 20% chance that Conflagrate keeps your Immolate and Shadowburn refunds its shard".** Two separate rolls on two spells; "a 20% chance that Conflagrate keeps your Immolate, and the same for Shadowburn's shard" avoids reading as one joint roll.
- **Sim note, not prose:** sim round 21 S1 measured on cooldown ahead of the hold at 36 on the raid order (hold -0.3%) with no Shadow and Flame point; the card holds there. No line states otherwise, so nothing is filed.

Carried: R23 Shoulds ("both halves also reach Wrack and your drains"; "as a finisher it is free" vs the mana cost); affliction.md :115 "the Shadow Bolts it replaces"; demonology.md :90 "every wipe or dead Succubus"; R18-S1 (Affliction Corruption "Instant with Improved Corruption" on 10 to 20); the Demonology Corruption note written for the raid build; the Demonology and Destruction Agony "new Doom" tail; Shadow Bolt rows rendering past 40 (component lane).

## Decisions

- **D1. Round 23's fixes hold.** R23-1: :57, :71 and the Wrack note now name Corruption and Bane of Agony (mask 0x402); :57's "Siphon Life and Bane of Doom get nothing" is a true partial list (the drains get nothing either, but Wrack's channel excludes them anyway). R23-2: :32 and :85 except Immolate's burn, which is the only DoT in the loop outside Pandemic's list (427712). R23-3: demonology.md :110 now presses Eureka! with trinkets and again for the last 35%, matching the card's on-cooldown use. R23-4 is routed (ledger round 23), not re-filed.
- **D2. Sim round 21's other edits hold.** destruction.md :58 matches 426316 at every rank. :76 matches the two rows. :128's "Conflagrate joins at 34 and eats your Immolate every time until you take Shadow and Flame" is exact (Conflagrate 34 L, Shadow and Flame 39 L; every rank's tooltip says "consuming"). "on a dungeon boss, cast it on cooldown for as long as you have shards to spend" matches the card and S3. The held row's note ("Before then it always eats your Immolate, so let the burn do most of its work first": 9 of 15 s, three of five ticks) and the Shadowburn note (shard, 20% a point, 4% a point) hold on every band they render. The Immolate note's "until Shadow and Flame is full, recast it whenever Conflagrate eats it" holds. Affliction :115 "From 29" matches the leveling order; "costs no global" matches `gcd_ms` 0; "until Bane of Doom at 60 there is nothing to hold it for" matches the condition (`spellIsKnown 603`) and S2's measured tie-to-gain.
- **D3. Why R24-1 is a page finding and not component lane.** The clause is the note's own first sentence; the component renders it faithfully. The note text is what makes the rendered line false at 60, and round 21's integrator treated the same shape on Conflagrate as a note fix.
- **D4. Why R24-2 is CONFIRMED without a sim.** The falsity is arithmetic on client durations: a 20 s buff refreshed every 12 to 13.5 s under the hold is never down, so the buff is not a difference between the two rules. The replacement reason (the keep chance) is the only other thing the first point changes on the Conflagrate side; a confirming sim is listed as optional.
- **D5. "On cooldown", "every", "only", "never" sweep, both pages and notes, against the card and every race.** Destruction :28, :40, :77, :80 are level 60 lines and hold with Shadow and Flame 5. :75 "so it never drops" is the 60 priority (at 5 points Conflagrate never eats it). Affliction :73 potions "on cooldown" are qualified by low mana. Eureka! lines (Affliction "as you start a Wrack", Destruction "as Conflagrate and Shadowburn come off cooldown", Demonology with trinkets) are race-specific and hold. Blood Fury is tied to trinkets on every page. Demonology :128 Soul Link 30, Demonic Knowledge 32 to 34, Master Demonologist 35 to 39 L hold.
- **D6. Superlatives and quoted values.** Unchanged from round 23 D5 and D7 apart from the new Shadow and Flame per-point values (4%, 20%), Immolate 15 s, Conflagrate 10 s, Amplify Curse 3 min / 30 s / no global: all hold at the ranks and talents the builds have.
- **D7. Not re-filed (routed or open):** WARLOCK-R23-4 (Demonology Soul Fire below raid spell power); R20-4 (answered by sim round 21 D2: the Shadow Bolt note holds); Amplify Curse's client mask (sim round 21 A2, tooltip and Blizzard replies decide); Decimation's 0% client effect; the engine's missing Soul Shard (sim round 21 S3 decision); sim round 20 L1 log items; race re-runs; BiS race captions; rows rendering before their talent (Amplify Curse 20 to 28 L / 20 to 24 R, Shadowburn 20 to 29 L, Conflagrate 25 to 33 L).

## Sims requested

- Optional (confirms R24-2's replacement reason): Destruction leveling order at 39, 42 and 44, bare band gear, held Conflagrate against on cooldown, with (a) Shadow and Flame's Shadow buff zeroed and (b) its keep-Immolate chance zeroed. Expect the on-cooldown lead to survive (a) and vanish under (b).
- Optional (card, not prose): Destruction raid order at 36, the hold against on cooldown with no Shadow and Flame point (sim round 21 S1 measured hold -0.3%); if it holds, the held row's gate could be a level or Fire and Brimstone gate rather than Shadow and Flame.
- Carried with the sim lane: WARLOCK-R23-4; the Demonology Horde race re-run.

## Sim findings

- None new.

## Engine questions

- None new. The engine keeps Immolate at 20% a point of Shadow and Flame (sim round 21 integrator), which is the client text.

## Checked and right this round

- **index.md**: every line but :19 (R24-3). :25 "Pandemic doubles the bonus on Affliction's" and :31 the DoT list hold.
- **affliction.md**: the new :115 sentence; R23 fixes; opener (:65) against the card (Curse of the Elements, Corruption, Doom, Siphon Life, Immolate, Wrack; Amplify Curse off the global); raid build paragraph (:59: Improved Bane of Agony 24 to 25 R, Grim Reach, four Aftermath 57 to 60 R, Unholy Power 5, Demonic Aegis 2).
- **demonology.md**: :110 R23-3 fix; every other line unchanged since round 23 and re-read against the card.
- **destruction.md**: :58, :76, :128 new text against 426316, the two Conflagrate rows and the leveling order; Levels 1 to 19 unchanged.
- **Voice:** no em or en dashes, no nightly numbers, no site or sim talk.

## Sources

- Client 1.60.1.70790: `talents/warlock.json` (426316, 18288, 427712, 1316697), `spellconst/warlock.json` and `spelltips/warlock.json` (18288, 707 to 11668, 1293817 to 18932), `addon-builds.json`, `addon-data.json` (`rotations/warlock-*`, every band).
- `web/src/lib/guides/rotation-lists.ts` (`clauseOf`, first sentence; cooldown floor 30 s).
- `design/reviews/2026-11-05-sim-audit-round21/warlock.md` (S1 numbers, S2, S3, integrator follow-up) and `integration.md`; `design/reviews/guide-accuracy/round23/warlock.md`; `ledger.md` rounds 19 to 23.
- No web research was needed: the client rows, the card conditions and the sim round 21 numbers decide every call.

## Note sweep

Format: spec, row (action), bands it renders, each named talent, spell, rank or value with the level the player has it on the leveling and raid orders (L / R), then the verdict. Bands from `addon-data.json`; Destruction now has 32 rows (two Conflagrate rows).

**Affliction**
1. Life Tap (1454 to 11689), 10-60. Spirit point for point every rank; Healthstone. OK.
2. Curse of the Elements (440892 to 1311680), 20-60. Five minutes every rank; own slot; one per target. OK.
3. Corruption (172 to 11672), 10-60. "Instant with Improved Corruption": 17-21 L / R, conditional. OK (Should R18-S1).
4. Bane of Doom (603), 60. Pandemic 31-33 L / 27-29 R. OK.
5. Amplify Curse (18288), 20-60. Talent 29 L / 26 R (earlier rows component lane). "before each Bane of Agony, whenever it is off cooldown" true 26/29-58, false as the rendered clause at 60: **WARLOCK-R24-1**. "At 60, keep it for the fight's final Agony" matches `spellIsKnown 603`; 50% on Agony (tooltip, A2); no global (`gcd_ms` 0). Rest OK.
6. Bane of Agony (980 to 11713), 10-60. "all fight before 60"; "biggest ticks come last"; Amplify "once you have" 29 L / 26 R. OK.
7. Siphon Life (18265 to 18881), 30-60. Talent 30 L / R; 30 s. OK.
8. Immolate (707 to 11668), 10-60. Refresh at 3 s with 2.0 s cast (1.5 s with Bane 56-60 L / 52-56 R). OK.
9. Wrack (1316697), 40-60. Talent 40 L / R; Soul Siphon full 38; out-damages Shadow Bolt on both terms every band (round 23 D5); "your Corruption and Bane of Agony take 10% more" matches 0x402. OK.
10. Shadow Bolt (695 to 11661), 10-60. "until Wrack at 40". OK.
11. Shoot (5019), drawer only. OK.

**Demonology**
12. Soul Fire (6353, 17924), 48-60. Trained 48 / 56; Decimation 2 by 31; "2.4 seconds once Bane is full" 56-60 L / 51-55 R conditional; "In the raid build, skip it ... Shadow bonuses": WARLOCK-R23-4, routed. Rest OK.
13. Life Tap, 10-60. As row 1. OK.
14. Curse of the Elements, 20-60. As row 2. OK.
15. Corruption, 10-60. "A 2 second cast without Improved Corruption": raid never takes it; leveling 47-51. OK (Should, carried).
16. Bane of Doom, 60. OK.
17. Bane of Agony, 10-60. Scoped to 60; "new Doom" tail carried Should. OK.
18. Immolate, 10-60. As row 8 (Bane 56-60 L / 51-55 R). OK.
19. Shadow Bolt, 10-60. "beats Searing Pain": measured by sim round 21 D2 at every level, holds. Master Demonologist 35-39 L / R conditional; Imp's 15% "with Demonic Pact at 40 in the raid build": Pact 40 R, sacrifice 26 R. OK.
20. Shoot, drawer only. OK.

**Destruction**
21. Curse of the Elements, 20-60. As row 2. OK.
22. Corruption, 10-60. "until Improved Corruption makes it instant": 52-56 L / 50-54 R. OK.
23. Life Tap, 10-60. As row 1. OK.
24. Bane of Doom, 60. OK.
25. Bane of Agony, 10-60. As row 17. OK.
26. Immolate, 10-60. Refresh at 3 s; Conflagrate 34 L / 32 R and Incinerate 40 behind "once you have"; "until Shadow and Flame is full, recast it whenever Conflagrate eats it": full 44 L / 42 R. OK.
27. Conflagrate, on cooldown (1293817 to 18932, gate 426316 known), 25-60 (talent 34 L / 32 R; acts from 39 L / 37 R). First sentence (hold until Shadow and Flame, then on cooldown) matches both rows. "20 seconds of more Shadow damage, 4% a point" matches 426316. "which is worth more than the burn it may eat": **WARLOCK-R24-2**. "at five points it never eats your Immolate": 44 L / 42 R. 
28. Conflagrate, held (same ids, gate 426316 not known, Immolate 6 s or less), 25-60 (acts 34-38 L / 32-36 R). "always eats your Immolate" before Shadow and Flame: every rank's tooltip. "most of its work": 9 of 15 s. OK.
29. Shadowburn (17877 to 18871), 20-60 (talent 30 L / 28 R). Any health; costs a shard (SpellReagents, S3); 20% a point refund and 4% a point Fire buff, conditional on the talent (39 L / 37 R). OK.
30. Incinerate (412758 to 1293813), 40-60. Talent 40; 25% with Immolate every rank. OK.
31. Shadow Bolt (695 to 11661), 10-60. "before you learn Incinerate at 40". OK.
32. Shoot, drawer only. OK.

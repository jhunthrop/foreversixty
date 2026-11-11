# Sim accuracy 2026-11-11: warlock (Affliction, Demonology, Destruction)

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main 450eddafa, fork from master 0e04b1284 (the pin). Client on the site: 1.60.1.70790. Live client: 1.60.1.70888 (wago `wowf`, 10 Nov); the site's move to it is blocked (issue #21).

**Totals: 6 confirmed bugs, all new data (Blizzard's 6 and 10 Nov hotfixes, which the site's 70790 tables and push do not carry), all fixed. 1 rotation change adopted on them. Engine pin 0e04b1284 to 19063086a.**

## Method

1. **Real fights first.**
   - Our API holds one new public report since round 29: gxrgvz4yzsi2 (8 Nov). It is an open-world PvP-realm log with 17 warlocks and a handful of hits each, too thin for rates.
   - foreverlogs.gg's report API now answers browser-like requests. It has 23 public reports since launch: Ragefire Chasm to Zul'Farrak and the new low dungeons, no level-60 content.
   - From those 23 reports, 48,902 damage events from warlock rows and their pets, plus Life Tap energizes and the armory captures (gear, stats, tree points).
2. **Blizzard since 4 Nov.** The notes thread (t/2360696), "WoW: Forever Hotfixes" (t/2403040: 6, 9 and 10 Nov posts), "Class Tuning Incoming" (t/2404092), forever-bugs issues, and player threads.
3. **Client and push.**
   - Every warlock spell, talent and pet row of SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellAuraOptions, SpellCategories, SpellClassOptions and the trait tables, diffed between 70790 and 70888.
   - wago's hotfix entries for 70862 and 70888, read record by record (push 113766 is the 10 Nov class push).
4. **Sims** on the Windows box (`tools/remote-test.sh`, new `--engine-sha` for the unpinned fork branch).
   - Raid 60, guide raid builds, both factions, seeds 1 and 50001, 10,000 iterations.
   - Affliction bare 40, 50 and 60 (horde, seed 1).
   - Band 30 level 32 breakdowns for the log comparison.

## Logs against the engine

| # | finding | evidence | confirmed? |
|---|---|---|---|
| L1 | **Life Tap is exact.** Freohr (Demonology 23 points, level 32 by the ranks he casts and his points, Spirit 125) taps for 261 on all 10 casts. Rank 3: 130, plus 6 for the levels above its own, plus his 125 Spirit, is 261. Valpuri taps for 174 on all 91 casts. The pet gets the same amount back (Demonic Energies), capped at its missing mana. | foreverlogs 2847, `resource_gain_mana` 31818 | Aligned |
| L2 | **Per-hit formulas agree with one another.** Searing Pain 84-94 (median 90), Corruption ticks 43-47, Immolate's hit 82-87 and Demonic Brand 35-41 non-crit, all from Freohr at level 32. Each one alone implies the same effective spell power (about 100 to 120, with Curse of the Elements' 10%) through the engine's base and coefficient. The four spells have very different base-to-coefficient ratios, so the engine's numbers for them are consistent. His gear shows 9; the armory leaves two Forever items unresolved, and Demonic Knowledge adds 32. | 2847 meta_events; armory 196479 | Aligned (spell power not in the log) |
| L3 | **Demonic Brand does damage in logs.** This was round 22 S8's open item. Fire consumes from the Imp land at 35-41 non-crit and 52-61 crit (157 and 15 hits, 1.5x crits), on a stable band at one level, in line with L2. | 2847, spell 1293698 | Aligned (S8 closed) |
| L4 | DoT ticks crit at 5 to 6% on Corruption and Agony across 14,000 low-level ticks, at 1.5x. The DoT-crit rule holds. | all 23 reports | Aligned |
| L5 | Boss-only pet share: Freohr's Imp 28% (a Searing Pain and Demonic Brand player at 32), Valpuri's Imp 10.5%. Our bands put the Succubus out at 30 (13 to 15% share). Different demons, so this is not a like-for-like check. | 2847 boss-stats | Not comparable |

## Confirmed and fixed

| # | finding | source | origin | fix |
|---|---|---|---|---|
| W1 | **Wrack hit 30% too hard.** 10 Nov: "Wrack damage reduced by 30%." Push 113766, SpellEffect 1349556: 25.2 a tick at 40, +1.75 a level to 60 (60.2 at 60), coefficient 0.2002 (was 36, 2.5, 0.286). | t/2403040 post 6; wago push 113766 | new data | E 63ad57cdd (`wrack.go`; hand row until the client move) |
| W2 | **Wrack still made Corruption and Agony take 10% more.** 10 Nov: "Wrack no longer increases the damage of the Warlock's periodic effects on the target." The same push turns effect 1 from aura 271 at 10 into a dummy at 100. The vulnerability aura is gone. | same | new data | E 63ad57cdd |
| W3 | **Wrack did not speed up the DoTs.** 10 Nov: "While channeling, your Siphon Life, Corruption, and Bane of Agony effects on the target deal their periodic damage 100% faster", and "the time it skips is kept". The new core `Dot.SetClockRate` runs a DoT's clock at a rate: next tick, later ticks and expiry rescale, on the tick grid. Wrack sets the three to 2 for the channel and back to 1 when it ends or is cut. Tests pin Corruption under a 6 s Wrack from 1 s (ticks at 2, 3.5, 5, 6.5, 9, 12 s) and the hotfix note's own 1 s example. | same | new data | E 63ad57cdd, e3c89a504 (`core/dot.go`, `wrack.go`) |
| W4 | **Soul Siphon counted any Affliction DoT**, so Bane of Doom counted. 10 Nov: "per each of your Corruption, Siphon Life, and Bane of Agony spells active on the target, up to a maximum of 36% (was any Affliction spell)". Push 113766 rewrites 17804's text to the same. | same | new data | E 63ad57cdd (`talents.go`) |
| W5 | **Imp Firebolt ranks 1 to 4.** 10 Nov: "Imp base damage of the first 4 Ranks of Imp's Firebolt increased". Push 113766: centres 6, 11, 15 and 18 (were 4, 7, 12 and 17), 0.3333, 0.1818, 0.2353 and 0.2222 wide, growing 0.3, 0.4, 0.4 and 0.4 a level. | same | new data | E 63ad57cdd (`imp.go`) |
| W6 | **Orc Command counted twice on Firebolt and Lash of Pain.** 6 Nov: "Fixes an issue where some Warlock pet abilities were benefitting too much from the Orc Command racial ability." In client 70888 and push 113702, 20575 effect 1 (5% on class mask 0x3000) is 0. | t/2403040 post 4; client 70888 SpellEffect 697943 | new data | E 63ad57cdd (`core/racial_command.go`) |

Not modelled, no damage path: pet health (+15% on the Imp, Voidwalker, Succubus and Incubus, +40 on the Voidwalker), Bane of Havoc's new global cooldown (no card casts it), and Incinerate for non-Destruction specs (no other card casts it). The client diff from 70790 to 70888 shows no other warlock change.

## Rotation (adoption rule)

On the fixed engine, the shipped Affliction card (Bane of Doom while 60 s or more remain, then Agony; DoTs recast only once they fall off) lost 15.5% at raid 60. Two candidates were measured, plus a Shadow Bolt filler control (-24%, rejected):

| variant | raid 60 horde s1 / s50001 | alliance s1 / s50001 | bare 40 / 50 / 60 |
|---|---|---|---|
| shipped card, new engine | 623.0 / 622.9 | 633.9 / 633.6 | 177.6 / 256.9 / 390.5 |
| Agony all fight, no Doom | 651.9 / 651.5 | 663.8 / 663.5 | 177.6 / 256.9 / 390.1 |
| **Agony all fight; Corruption, Siphon Life, Agony recast under 3 s (adopted)** | **654.9 / 655.7** | **667.2 / 668.0** | **178.4 / 257.6 / 391.1** |
| the same at 6 s | 652.0 / 652.1 | 664.6 / 664.5 | 177.3 / 254.8 / 388.1 |

The adopted card beats the shipped one beyond error at every measured band: +5.1 to +5.4% at raid 60, +0.15 to +0.45% bare 40 to 60. Errors are 0.1 to 0.3 DPS. Before Wrack, the under-3 s recast loses (bare 20: 60.6 to 60.0, bare 30: 103.2 to 102.3, both beyond error), so the card gates that clause on knowing Wrack. Below 40 the card is unchanged, and the ladder rungs at 10 to 38 are unchanged too. Nothing it casts is dropped except Bane of Doom, which the new Soul Siphon and the Wrack speed-up leave behind. Amplify Curse now goes on the next Agony whenever it is ready.

## Before and after

Raid 60, guide raid builds, published BiS, 10,000 iterations, seeds 1 / 50001, pin 0e04b1284 against the fixed engine (e3c89a504 to 278876977; only the APL copy differs):

| spec | horde | alliance |
|---|---|---|
| Affliction | 737.3 / 737.0 to 654.9 / 655.7 (-11.2 / -11.0%) | 750.3 / 750.0 to 667.2 / 668.0 (-11.1 / -10.9%) |
| Demonology | 726.8 / 727.5 to 726.1 / 726.8 (-0.10%, Command) | 736.9 / 736.8, unchanged |
| Destruction | 791.0 / 790.8 to 790.7 / 790.4 (-0.04%, Command) | 801.2 / 801.2, unchanged |

Affliction bare (horde, seed 1): band 40 -6.8%, 50 -8.1%, 60 -11.4%. Below 40 (no Wrack) nothing moves. Blizzard's tuning post (t/2404092) meant to keep Affliction's total damage where it was. On the client's numbers, Wrack's 30% cut outweighs the faster DoTs.

## Sim changes in this landing

- Fork `lane/wl-1111`: e1f48e674 (tests, red on the box), 63ad57cdd (fix), e3c89a504 (test fix), 19063086a and 278876977 (Affliction APL synced). Site pin 278876977. `tools/remote-test.sh --engine-sha`.
- Shared code: `core/dot.go` adds `SetClockRate`; only Wrack calls it. At rate 1, `RecomputeAuraDuration` is unchanged. No DPS moves outside warlocks.
- Goldens: `sim/request/testdata/ladder/warlock-*.golden.md` regenerated on the box.

## Guide and card

- affliction.md: the talents (Wrack, Soul Siphon, Amplify Curse, Improved Bane of Agony), DoTs, Execute, the 20 to 60 leveling paragraph, and the date.
- The curated card: Corruption, Amplify Curse, Agony, Siphon Life and Wrack notes, and the panel lines. Bane of Doom's row is removed.
- Demonology and Destruction prose is still true (Command still "makes your demon hit harder").
- Below 40 the card is unchanged.

## Guide review

wow-player read the whole Affliction page and the card.

- **First pass: NO-SHIP**, on four MUSTs:
  - the Curses paragraph's leftover "last Bane of Agony";
  - the stated reason for the 3 s rule;
  - the Shadow Trance card line ("from 34", "the moment");
  - the Malevolence bullet's "34 and 35" against the talent list's 26 and 27.
- SHOULDs taken:
  - Amplify Curse never clips a running Agony;
  - the Wrack clause on the DoT notes is gated at 40, in the card and in the rotation;
  - "Wrack and the faster DoTs out-damage Shadow Bolt".
- Not taken, as later sim work: the Execute idea (Wrack to cash in the DoT time left in the last seconds), the full level-by-level talent list, and the reviewer's doubts on Curse of the Elements' wording, Curse of Recklessness' 505 and the Improved Shadow Bolt pick. Life Tap's Spirit term is confirmed by L1.
- **Second pass:** see the end of this file.

## Measure at refresh (ledger MR39 to MR41)

- Affliction, every band from 40 and raid 60: BiS, weights, races and the raid build's talents on the new Wrack. Soul Siphon, Siphon Life and Improved Bane of Agony gain; Bane of Doom is out of the rotation. Also the leveling order's Wrack-at-40 path.
- Orc warlocks: race margins for Demonology and Destruction, from Command's second 5% leaving.
- When the site moves to 70888 (issue #21): the generated `Wrack*` and `Firebolt*` tables will carry the push. `wrack.go` and `imp.go` can then read them again, and their tests stay as they are.

## Logs

Looked 2026-11-11 13:00 to 14:30 UTC.

- API: two public reports. z2opt2x2gs5u (20 Oct) and gxrgvz4yzsi2 (8 Nov, open world).
- foreverlogs.gg: 23 public reports since launch, all level 15 to 45 dungeons. Its report API answers with browser headers.
- Warcraft Logs: no Forever zone on any host.
- Still needed: any level-60 Affliction log after 10 Nov (Wrack's tick and the DoT speed-up), and a level-60 Succubus log.

## For the owner

- The site serves 1.60.1.70790 while the live client is 1.60.1.70888. The data move is blocked (issue #21, data.yml's fetch of 70862 failed). Until it moves, this lane's hand rows carry the 10 Nov push; no other class was checked against 70888.
- Affliction's published numbers, BiS and card are stale by about 11% at raid 60 until the refresh (the refresh hold).

## Sources

- Blizzard, WoW: Forever Hotfixes, https://us.forums.blizzard.com/en/wow/t/2403040 (posts 4 to 6: 6, 9 and 10 Nov).
- Blizzard, Class Tuning Incoming, https://us.forums.blizzard.com/en/wow/t/2404092 (6 Nov).
- Blizzard, development notes, https://us.forums.blizzard.com/en/wow/t/2360696 (post 9).
- wago.tools DB2 CSVs for 70790 and 70888; wago.tools/hotfixes pushes 113702 and 113766.
- foreverlogs.gg reports 2815 to 2849.
- foreversixty API report gxrgvz4yzsi2.

### Second pass

wow-player: **SHIP**, with no MUSTs. Left open for later lanes:
- the level-by-level leveling talent list (28 and 29, and 34 to 38, are not listed);
- the Execute idea;
- Curse of the Elements' wording, Curse of Recklessness' 505 armor and the Improved Shadow Bolt pick, each a sim or client check.

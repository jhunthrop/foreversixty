# Sim-accuracy round 24: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70790. The site branch is wt/r24-shaman (off int/r24 a5ae23579). The fork branch is lane/r24-shaman (off int/r24 3b4d6b054). Gear is the frozen BiS ($R/bis-frozen). Raid 60 runs the guide raid build ($R/rt). Bands run the leveling order cut to the level (`$R/shaman/tal.py`). Build mode: touched tests only, no wow-player gate.

Sources read:
- the round 23 shaman report and its integration "Open, for round 24";
- the guide-accuracy round 26 shaman report (gr26 scratch). Its one "Sims requested" item is the Berserking timing;
- Blizzard's Forever notes thread https://us.forums.blizzard.com/en/wow/t/2360696 (JSON read 00:05 UTC 6 Nov). It still has 10 posts; the last is 4 Nov 19:35 UTC and post 9's last edit is the same minute. No post since round 23. A web search for Forever shaman hotfixes in November found only retail Midnight threads.

**Six confirmed bugs, all fixed:**
- one rotation bug, with a shared engine root (Berserking);
- four rank rows (the sweep at the first two levels after each new rank);
- one card row that starts too early (handed to platform, not fixed here).

No published raid-60 number moves.

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| B1 | **A Troll Elemental pressed Berserking together with Lightning Mastery on the pull.**<br>**The card and guide:** card row 1 and elemental.md's Races section say a Troll presses Berserking after Mastery ends. Under Mastery, rank 10 Lightning Bolt is a 1.5 s cast (2.5 s, minus Elemental Alacrity's 0.5 and Mastery's 0.5). It already sits on the 1.5 s global cooldown, and casting speed never shortens that cooldown (`sim/core/cast.go:256`, no GCD haste in Vanilla).<br>**What the rotation did:** it left Berserking to "autocast other cooldowns", which fired it on the pull.<br>**Engine root:** `racials.go` registered six Berserking major cooldowns on one timer: the client spell plus five custom-percentage tags. An APL line that casts Berserking takes only tag 0 out of autocast, so tag 2 (the same 10%) still fired on the pull. No rotation could time it.<br>**Fix (engine):** only tag 0 is a major cooldown now. The tags stay castable from an APL.<br>**Fix (rotation):** a Berserking line after the cooldowns, gated on `not auraIsActive(16166)`.<br>**Measured, Troll, 10,000 iterations:**<br>- raid 60: 644.4 to 648.4 (seed 1) and 644.8 to 649.1 (seed 50001), +0.6% and +0.7%;<br>- bare 60: 236.6 to 236.5 and 237.0 to 237.5 (a tie).<br>So the card is right and the sim now follows it. | `sim/core/racials.go:540-583`; `TestOnlyTheClientBerserkingIsAMajorCooldown` (red: tags 2-6 left in autocast); `sim/shaman/elemental/berserking_hold_test.go` (red on the old APL: Berserking up inside Mastery at 0 s); `$R/shaman/out/berserk.txt` | CONFIRMED-BUG | old (the six tagged cooldowns predate the fork follow; the card note was written in round 24 of the guides) + guides-handoff (guide round 26) | E 248e1bdad, E 57403942d, S 37b4a0907 |
| B2 | **Guide text for B1.** elemental.md:76 told a Troll to press "any damage racial (Blood Fury or Berserking)" with Mastery, saying "the racial comes back sooner". Both are false for Berserking: its cooldown is 3 minutes, the same as Mastery's. The paragraph now names Blood Fury, and a Troll presses Berserking the moment Mastery ends. This is guide round 26's SHAMAN-R26-1 fix. The coach line (R26-2) is generated (tools/coach) and left to the guide lane. | `web/src/content/guides/shaman/elemental.md:76` | done | guides-handoff | S 37b4a0907 |
| B3 | Adding the line shifts the Lightning Bolt filler from priority line 8 to line 9. The rank row (`lines: [8]`) followed the wrong line at first, and bare 60 fell 6.6% before it was moved to `[9]`. `TestElementalLightningBoltRanksByBand` guards it. sim/rankchoices was copied by hand (a plain `cp`, `make rank-choices-sync`). | `data/curated/rank-choices.json` | fixed in the same commit | - | S 37b4a0907 |
| B4 | The other optional measurements in the round 26 report: there are none. Its only carried item, the Enhancement Earth Shock gate, was done in round 22. | gr26 shaman.md "Sims requested" | - | - | - |
| B5 | Round 23 open items:<br>- Maelstrom Weapon flat chance against 20 a minute: still OPEN in principle and aligned for every published band (every band wears a two-hander of 3.3 s or slower);<br>- sub-20 damage penalty: OPEN-low;<br>- Lightning Overload rank spells and level-60 book ranks: pending owner. | round 23 S1, S2 | OPEN / pending | - | - |
| B6 | Logs, 00:09 UTC 6 Nov:<br>- `api.foreversixty.gg/v1/reports/recent` lists only the 20 October report;<br>- foreverlogs.gg `/api/reports/public` answers 403;<br>- Warcraft Logs v1 classic `/zones` answers 401 with the stored key. | curl | no logs | - | - |

## Rank sweep (every ranked spell the cards cast, at the first two levels after each new rank)

Method: `$R/shaman/sweep.sh`. It runs the newest rank against the one below at L and L+1 for every new rank from 10 to 60, on bare band gear with the leveling order at the level. Bare and raid 60 are included. The screen ran both factions at seed 1 with 3,000 iterations (`out/sweep-horde.tsv`, `out/sweep-alliance.tsv`). Each candidate was then confirmed at 10,000 iterations, seeds 1 and 50001, both factions, plus troll and tauren where the factions split (`out/confirm1.txt`, `out/race-rows.txt`).

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Elemental Lightning Bolt filler at 20 to 25 cast the newest rank 4.** Rank 3 does more: +0.3% to +4.3% at every level, both factions, both seeds (for example 21 horde 30.3 to 31.4, 23 alliance 32.7 to 34.1). Round 19 found a tie at 20 to 25 and started the row at 26. The row now runs 20 to 27. The top-rank line (gated on mana ahead of the fight) ties its own downrank at every new rank (`out/lbtop-h1.txt`), so it stays on the newest rank. | `out/confirm1.txt` | CONFIRMED-BUG (rank row) | old (round 19 judged 20-25 a tie) | S 08434b177 |
| K2 | **An orc's Elemental Flame Shock at 19 cast the new rank 2.** Rank 1 does more: 16.90 to 16.96 and 16.92 to 16.99 (+0.36%, +0.41%, error 0.013). Other races keep rank 2: troll -0.3%, tauren -0.7% to -0.8%, dwarf -2.6%. At 18 every race ties. The row is race-keyed to orc and covers 19 only. | `out/confirm1.txt`, `out/race-rows.txt` | CONFIRMED-BUG (rank row) | old | S 08434b177 |
| K3 | **Enhancement Earth Shock at 25 (the level Stormstrike arrives) cast rank 4.** Rank 3 does more for an orc (58.68 to 59.12 and 58.64 to 59.06, +0.7%) and a troll (+0.3%). A tauren keeps rank 4 (-0.3%) and a dwarf ties. From 26, rank 4 wins again. The row is race-keyed to orc and troll and covers 25 only. | same | CONFIRMED-BUG (rank row) | old | S 08434b177 |
| K4 | **Restoration's tank Healing Wave at 18 and 19 cast the new rank 4.** Rank 3 heals 3.4% to 5.4% more (both factions, both seeds). It also beats rank 2 there (18: rank 2 30.6, rank 3 32.8, rank 4 31.2). From 20 the existing rank 2 row wins (20: rank 2 68.7, rank 3 65.8). The row now starts at 18. | `out/confirm1.txt`, `out/hw-low.txt` | CONFIRMED-BUG (rank row) | old | S 08434b177 |
| K5 | Every other rank point holds with the newest rank or matches an existing row:<br>- Flame Shock 28, 40, 52 and 60;<br>- Searing Totem at every rank, on both specs;<br>- Earth Shock 14, 24, 36, 48 and 60;<br>- Lesser Healing Wave;<br>- Healing Stream Totem;<br>- Riptide 3 at 60;<br>- Chain Heal (its rank 1 row holds at 46 to 47);<br>- Healing Wave from 20 (existing rows);<br>- the bolt filler from 26 (existing rows).<br>Not swept, because they are pack-only on these cards: Chain Lightning, Magma Totem and Elemental's Fire Nova. Also not swept: the buff totems (Strength of Earth, Grace of Air, Mana Spring), where a higher rank only adds. Enhancement's five-stack bolt and Clearcasting Fire Nova are free, so the top rank always wins. | sweep files | Aligned | - | - |

Guide text for K1 and K4 shipped in the same commit:
- the Elemental filler note and elemental.md:118 now read "3 from 20 to 27";
- the Restoration tank Healing Wave note and restoration.md:116 now say rank 3 at 18 and 19.

K2 and K3 change no note claim, since neither note names a rank; the card carries them as race ids.

## Re-audit

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| A1 | **Enhancement's "Lightning Bolt at five Maelstrom Weapon stacks" row prints from 37.** It can fire only at 5/5 Maelstrom Weapon: 42 on the leveling order and 39 on the raid order. The condition is `auraIsKnown(408498, tag 5)`. The engine registers the tag as the talent's exact rank (`sim/shaman/talents.go:883`), so 5/5 is the only rank that knows tag 5. Round 23's row-floor rule (`_required_known_ids`, `addonrotation.py:283`) ignores the tag and floors the row at the first point. The note does say "once you have all five points". The same gap reaches mage Arcane's Missile Barrage tag 2. Handed to platform (shared pipeline code). | `data/pipeline/addonrotation.py:283-306`, `talent_levels.build_point_levels` | CONFIRMED-BUG (card data), not fixed here | R23 20d45f26b | - |
| A2 | Round 24 platform C1: Fire Nova on Clearcasting now starts at 52, where the raid order takes Elemental Focus. The leveling order never takes it, and the note says "in the raid build". | card rebuilt locally from `build_rotations` | Aligned | - | - |
| A3 | Talent points that do nothing, checked on all six orders against the client texts:<br>- Call of Flame adds 5% a point to Fire totems, Flame Shock and Fire Nova, and 3 s a point to Flame Shock (`talents.go:112`, `flame_shock.go:58`). So Enhancement's three leveling points (51 to 53) do work, on Searing Totem;<br>- Concussion covers Lightning Bolt, Chain Lightning and Earth Shock;<br>- Elemental Devastation is modeled;<br>- Mindfulness is modeled;<br>- Totemic Focus is modeled (`talents.go:556`);<br>- Tidal Focus is modeled;<br>- Elemental Warding, Improved Reincarnation, Improved Ghost Wolf, Healing Focus and Ancestral Healing are the standing, accepted tolls. | `data/builds/1.60.1.70790/talents/shaman.json` | Aligned | - | - |
| A4 | Every Troll spec: B1's engine change moves no rotation that leaves Berserking to autocast, because tag 0 always sorted first among the six. No other curated APL casts 20554. | grep of data/curated/apl | Aligned | - | - |

## Before/after

**Raid 60:** the published races (orc, dwarf, tauren) show no change at either seed. Seed 1 / 50001:

| spec | alliance | horde |
|---|---|---|
| Elemental | 653.8 / 654.1 | 650.5 / 651.1 |
| Enhancement | 833.7 / 833.9 | 839.5 / 841.1 |
| Restoration | 742.5 / 742.6 | 740.6 / 740.5 |

Source: `out/after.tsv`, matching `$R/before.tsv`.

**Troll Elemental, seed 1:**
- raid 60: 644.4 to 648.4 (+0.6%);
- bare 60: a tie.

**Bands (seed 1, before = `$R/sb-before`):**
- Elemental 21 horde: 30.3 to 31.4 (+3.6%)
- Elemental 23 alliance: 32.7 to 34.1 (+4.3%)
- Elemental 19 orc: 16.9 to 17.0 (+0.4%)
- Enhancement 25 orc: 58.7 to 59.1 (+0.7%)
- Enhancement 25 troll: 58.3 to 58.5 (+0.3%)
- Restoration 18 horde: 31.2 to 32.8 HPS (+5.1%)
- Restoration 19 alliance: 32.5 to 34.0 HPS (+4.6%)

## Files to regenerate

- data/builds/1.60.1.70790/addon-data.json and addon/ForeverSixty/Data.lua (rotations). The shaman card bands change at:
  - Elemental 19 to 22;
  - Enhancement 25 and 26;
  - Restoration 18 and 19 (the Healing Wave rank).
  The new Berserking line does not render (a racial, as Stoneform).
- sim/request/testdata/ladder/shaman-elemental.golden.md: the ladder's non-Troll build now reports Berserking 20554 as unresolved, and the rows move the bands.
- Coach starters, if the guide lane's R26-2 override lands (not touched here).
- The ui/elemental_shaman fork APL copy is synced by hand (E 248e1bdad). The restoration fork copy's notes were already out of step and were not touched.

## Measure at refresh

- Carried from round 23:
  - Enhancement band 30 with Stormstrike;
  - Restoration Chain Heal rows at bands 40, 50 and bare 60;
  - Elemental band 20 and 30 talent strings;
  - Thunderbrew's Boot Flask.
- Troll Elemental race standing on the new number (orc remains the horde pick: 650.5 against 648.4).

## For other lanes

- **Integrator / platform (FLAG: shared engine code):** E 248e1bdad changes `sim/core/racials.go` for every Troll. Only the client Berserking is a major cooldown now. DPS is neutral for any APL that does not cast Berserking itself (none besides Elemental does), and the priest healing cooldown tests pass.
- **Platform:** A1. Make `_talents_allow` read an `auraIsKnown`/`spellIsKnown` tag on a ranked talent as "the Nth point". That needs a rank-level map from `build_point_levels`. Expected moves:
  - Enhancement's five-stack bolt row 37 to 42;
  - mage Arcane's Missile Barrage tag 2 row.
- `data/pipeline/apl.py` now lists Berserking 20554 beside Stoneform and Read Leyline as a racial id that spellconst never carries.

## For the guide lane

- Coach R26-2 (shaman-elemental Lightning Mastery line, "any damage racial") through `starter_overrides.py`. The page side is done (B2).
- Round 26 S1 (the Chain Lightning "until Lightning Mastery" wording) and the coach findings R26-3 to R26-6 are untouched.

## Still open

- Maelstrom Weapon's flat chance (needs a log of a 37+ Enhancement shaman on a weapon faster than 3.0 s).
- Sub-20 damage penalty (needs a level 25 to 30 shaman's Lightning Bolt 3 tooltip at two spell damage values).
- Live logs: none anywhere at 00:09 UTC 6 Nov (B6).

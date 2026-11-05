# Saved sim result spec and mock: wow-player review (2026-11-04)

Reviewed: `design/specs/2026-11-04-sim-results.md`; the boards `result-1440/1024/2000/390.png`, `timeline-1440/390.png` and `states-1440.png`; the real run `design/mocks/data/simview/fury-result.json`; the competitor captures (`desk.png`, `phone.png`); and our current page (`shots/before/saved-1440.png`).

## Verdict: NOT SHIPPABLE as drawn

The layout is close to SHIP. It is a big step up from today's page, which shows "An unnamed spell #1680" rows and hides everything behind tabs, and on hierarchy it beats the competitor's page. What sinks this board is the data. The run behind it is a Fury warrior who **never pressed Bloodthirst**, wearing two Blizzard QA test items, and the page says nothing about either. Post this in guild chat and the first reply is "where's Bloodthirst lol", the second is "90 Green Rogue Ring?". The spec also has a gap that lets exactly that run through without a warning. Fix the two blockers and the must-fix list below, and the layout itself is SHIP WITH FIXES.

**The player's question:** "How much DPS does this setup do, and where does it come from?" On the 1440 board the answer is the first thing on screen: 402.3 DPS in gold with "± 0.5 · 0.13% error", and the top three ability rows sit above the fold. On the 390 board the first screen shows the figure, the error and the button. That part is right.

**What is genuinely good, and better than the competitor:** there is one figure, not seven at the same size. Damage by ability sits above the timeline. The outcome split sits on the row. Every ability has its icon and rank ("Execute Rank 5", "Heroic Strike Rank 8", "Battle Shout Rank 6" are all the right launch ranks). The rage lane in Close up shows the warrior pooling to 100 at 2:33 to 2:42 before the Execute wash, which is exactly what I open WCL to find. Cooldowns line up readably: Death Wish and Blood Fury at the pull, Recklessness at the Execute phase. The W3 Recklessness line is correct and in plain words, and better than the competitor quietly turning Recklessness off.

## Findings

### Blockers

**1. The page shows a Fury warrior with 0 Bloodthirst casts and raises no warning.**
- The run has no `spell:23894` anywhere in `summary.casts` or `damage_done`.
- On the board, Rotation line 7 reads "Bloodthirst Rank 4, Not cast in this run". Line 6 (Execute) and line 8 (Whirlwind) both refer to Bloodthirst in their notes.
- The timeline has no Bloodthirst lane, and between Whirlwinds rage idles at 50 to 70.
- 402 DPS for Lionheart Helm, Onyxia Tooth Pendant, Crusader on both weapons and a full raid buff list reads as broken. The owner will hold it against the competitor's 993.
- §4.7 suppresses W2 exactly when the talents do not fit, which is this run. The likeliest root cause is open question 1: a mis-decoded talent string that drops the 31-point talent.

Fix:
- (a) Add a warning that does not depend on `talentsFit`. When a rotation line that names the spec's capstone talent ability has 0 casts, show: "{Bloodthirst} was never pressed in this run. Check the talents in Run this yourself." This rule goes in `warnings.ts` above W3.
- (b) Answer open question 1 before any fixture is pinned.
- (c) Regenerate the mock and `result-full.json` from a run whose talents decode and that casts Bloodthirst.

**2. The fixture's gear includes Blizzard test items.**
- "90 Green Warrior Waistband" (Waist · 90) and "90 Green Rogue Ring" (Finger 1 · 90) are internal QA items. They inflate the "Average item level 66.6", and any Classic player knows them on sight.
- Trinket 2 is "Nifty Stopwatch", a run-speed Use trinket, in a raid DPS sim.

Fix:
- Build the mock and the fixture from a real kit: the published Fury pre-raid list on /bis, or a real imported character.
- For the accuracy lane: the sim's gear source should never offer "90 Green …" or other test items.

### Must fix

**3. The enchant line uses recipe names, not the client's green line.**
- The spec (§3.8) promises "the client's green enchant line", then prints the recipe name after " - ".
- The client tooltip prints the enchant's own name. On the board that means:
  - "Greater Strength" should read "+7 Strength";
  - "Superior Strength" should read "+9 Strength";
  - "Greater Stats" should read "+4 All Stats";
  - "Lesser Agility" should read "+3 Agility";
  - "Minor Speed" should read "Minor Speed Increase";
  - "Lesser Arcanum of Voracity" should read "+8 Strength";
  - "Crusader" stays "Crusader".

Fix: use the SpellItemEnchantment name the client shows, not the recipe name.

**4. The "Average" column must be per landed hit.**
- "Main hand … 228" is total damage divided by attempts, so it counts the 24.6% misses and the glances. Details! and WCL ("Avg Hit") both average landed hits, so a player reads 228 as "my Runesword hits for 228" and thinks something is broken.
- Crit has the same problem: "26.2%" is 16 of 61 attempts.

Fix:
- Average = damage / landed hits (hits plus crits, glances included). Header "Avg hit".
- Crit = crits / landed hits, as WCL does it.
- Pin both in the §10 accuracy checks.
- Once the engine returns `max`, add a "Biggest hit" column. Players check it, and the field exists but is 0 today.

**5. The Execute threshold says three different things.**
- Settings: "Execute phase: Last 25%".
- The Execute note: "below 20%".
- The Heroic Strike note: "never from 25% down".
- Execute in the client is "target below 20% health". A player reads "Last 25%" as a wrong threshold.

Fix:
- Change the Settings value to "Last 25% of the fight (target under 20% health)", or move the default to 20%. The accuracy lane should say why 25% is the default.
- Make the guide notes use one wording: "in the Execute phase".

**6. There are no character stats.** This is the top thing the competitor has that players want. Their gear block ends with:
> Melee hit 11% · specials capped at 8% · white swings (dual wield) need 27%
> Strength · Agility · Attack power · Crit

Half the reason a warrior sims is "am I hit capped".

Fix: in the Gear card foot, or the paperdoll centre under item level, add one row: Attack power, Hit % against "8% for specials · 27% for dual-wield white swings" against a level 63 target, Crit %, and Weapon skill per hand. All of these are computed from the request.

**7. Settings lacks the target.**
- There is no target level or armor.
- The competitor shows "Level 60 (level 63) · 3,731 armor".
- The request stacks Sunder Armor, Expose Armor, Faerie Fire and Curse of Recklessness. Sunder and Expose do not stack in 1.12.

Fix:
- Add "Target: Level 63 boss · {armor} armor, {after} after debuffs" to Settings.
- In the On the target chips, mark a debuff the game suppresses ("Sunder Armor · replaced by Expose Armor") instead of listing both as if both apply.

**8. Overpower as a buff (open question 5): leave it out for Fury.**
- "Overpower 35.7% 16×" sits second in On you.
- The page's own Rotation line 11 says Overpower "never comes up" once you live in Berserker Stance.
- A Fury player reads "Overpower 35.7%" as a buff they never see. The client has no such aura; the client lights the Overpower button.

Fix:
- Never show the Overpower window in Buffs and debuffs.
- When the rotation actually casts Overpower (Arms, a stance-dance Fury APL), draw it on the timeline as a thin lane under Casts named "Overpower window", using the WeakAuras wording players already use. Hide it when the rotation never casts Overpower.
- The 16 windows do match the run's ~16 dodges, so the data is right, just irrelevant here.

**9. Off hand misses (open question 2): the direction is real Classic; the size needs checking; and the page should say why.**
- While Heroic Strike (or Cleave) is queued, off-hand white swings use the one-hand hit table with no dual-wield penalty. This is the well-known 1.12 "HS queue" behaviour that Classic Era kept, and the reason Fury queues Heroic Strike.
- Main-hand white swings exist only when Heroic Strike is not queued, so the main hand always pays the full dual-wield penalty. The off hand will always miss less.
- The tables agree on everything else:
  - dodge is 6.6% main hand vs 6.5% off hand;
  - glance is 39.3% vs 39.8%;
  - both weapons are swords, so the Orc axe racial plays no part.

  Only the miss chance differs, which is the right shape.
- The size is the problem. The queue track is up 46.7 s of 179.9 s (26%). That predicts the off hand at roughly 24.6 − 0.26 × 19 ≈ 19.7%, not 14.8%. The 14.8% implies about half the off-hand swings land under the queue. The accuracy loop should check whether the engine applies the one-hand table past the queue window.

Fix:
- After the accuracy check, add a sub-line on the Off hand row: "Misses less while Heroic Strike is queued". It is true, and it teaches the player something (tenet 14).

**10. Phone timeline: icon-only lanes cannot be told apart for a warrior.**
- Bloodrage, Death Wish, Berserker Rage, Recklessness and Execute are all red-orange icons at 18px. On `timeline-390.png` I could not tell the On you rows apart without the desktop board.

Fix: print the name at 11px muted inside each lane, sticky at the plot's left edge, and let tapping the icon open the spell tooltip. Keep the `aria-label`.

**11. "Run this yourself" has no destination in the spec.**

Fix:
- Specify that it opens `/sim` with this run's race, gear and enchants, talents, buffs, consumables and fight loaded and editable, ready to press Sim.
- Rename it "Open in the sim". A friend reading your link wants "load this", and "Run this yourself" reads as "type it all in".

**12. Extra attacks have no rule.**
- Windfury Totem, Hand of Justice, Ironfoe, Thrash Blade and Flurry Axe all add extra swings. The competitor shows them as "Melee (extra attack)".
- The spec only knows "Main hand" and "Off hand".

Fix: specify one row, "Extra attacks", using the main-hand weapon's icon, with a sub-line naming the sources ("Hand of Justice 6 · Windfury Totem 21"). Do not fold them silently into Main hand.

**13. Before-the-pull rows ignore whether they happened.**
- "-1.5 s Battle Shout Rank 6" is listed, but the sample has no Battle Shout: the raid buff covers it.
- Line 5 is dimmed for the same reason, but the pre-pull row is not.

Fix: pre-pull rows get the same dimmed "Not cast in this run" treatment as priority lines.

**14. The spread caption overclaims.** "Your pulls will spread the same way." Real pulls do not: they add movement, buff timing and latency.

Fix: keep "Two runs in three land between 371 and 431 DPS." and drop the second sentence (tenet 8).

**15. Illustrative numbers that contradict the run's own data.** Labelled or not, these become the hand fixture lanes A to C build against.
- Deep Wounds "75.0% · 32×", and the gappy On the target bar, disagree with the run's 58 ticks at about 3 s each. That is roughly 97% uptime, and the competitor shows 99.2%.
- "Gained from" lists Unbridled Wrath 86 but no Berserker Rage. Yet the run casts Berserker Rage 6 times, rotation line 10 says it gives 10 rage, and the sample shows rage going from 3 to 11 at 0:01.5.

Fix: derive both from the run (tick count × tick period; Berserker Rage's rage), or leave them out of the hand fixture.

**16. Two talent splits on one screen.**
- The band reads "20/31/0 talents", and the Talents card two screens down reads "17 / 32 / 2".
- §4.7 says this run shows the fallback, so the default board should show that.

Fix: once blocker 1 is fixed, the board shows the run's own trees. Until then, the 1440 and 390 boards show §4.7.

### Polish

17. **Heroic Strike in Whole fight.** The queue press and the landing draw as a pair of ticks, so the lane reads like 56 Heroic Strikes. Draw only landings in Whole fight; show the hollow queue ticks in Close up.

18. **Phone outcome split.** On phone the split breaks inside a pair ("Glance / 39.3%"). Keep each "Label value" pair together.

19. **Settings clutter.** "Ran on: Our server" means nothing to a player, so drop it. Keep the engine sha muted (tenet 7).

20. **Consumables in the band.** Consumables move a warrior's DPS more than anything else on the page. Add "Consumables: None" or "{n}" to the facts row next to Buffs. Actions per minute can move to Settings.

21. **Buffs that do nothing for you.**
- On the target: Demoralizing Shout and Roar, Thunder Clap, Curse of Weakness, Insect Swarm, Scorpid Sting, Judgement of Light and Wisdom do nothing for your damage.
- On you: the three resistance totems and two resistance auras, Mana Spring, Sanctity Aura and Thorns do nothing for a warrior.

  Sort what changes your damage first and group the rest under a muted "No effect on your damage".

22. **Holy Strength labels.** "Holy Strength (main hand)" and "(off hand)" are good once open question 4 is confirmed. The client shows two plain "Holy Strength" buffs, so the tag is a helpful addition, not a wrong name.

23. **"Whirlwind (off hand)".** Use the client's own name for spell 462891 if the 1.60 client has one. If it does not, "(off hand)" is fine. Confirm the Forever Whirlwind tooltip says both weapons, as rotation line 8 claims.

## Data I do not believe (for the accuracy reviewer)

- **No Bloodthirst cast and 402.3 DPS** (blocker 1). The likely cause is the talent decode in open question 1.
- **The raid buff preset has no Windfury Totem.** It is the largest melee buff a raid gives, shamans exist on both factions in Forever, and the competitor models it. Yet the same preset carries Fire, Frost and Nature Resistance Totems and Mana Spring Totem.
- **Sunder Armor and Expose Armor are both applied.** They do not stack in 1.12.
- **The off-hand miss gap** is about twice what the Heroic Strike queue uptime predicts (finding 9).
- **Berserker Rage giving rage**: 6 casts, a 3 to 11 jump, and a note crediting Improved Berserker Rage. That is a TBC talent. It is fine if the Forever 1.60 Fury tree (18 talents) really has it. Verify against the client's talent table.
- **Deep Wounds 75% uptime** against 58 ticks (finding 15).
- **Deep Wounds tick crit 32.8%** is consistent with the owner's fact that DoTs crit in Forever. The competitor's "Ticks don't crit" is theirs, not ours; I am not calling this an error.

## The open questions, from a player's view

1. **Talent string.** This is the most important question on the list. It is the likely cause of the missing Bloodthirst. Answer it before any fixture is pinned.
2. **Off hand misses.** The direction is real (the Heroic Strike queue removes the off hand's dual-wield penalty); the size needs a check. Explain it on the row (finding 9).
3. **`rotation_sha`.** Yes. "Today's rotation" should show exactly when the rotation differs, not only when the engine is stale. Players re-open old links after APL changes.
4. **Holy Strength hand tags.** Confirm from the adapter, then ship the "(main hand)" and "(off hand)" labels.
5. **Overpower.** Never "Overpower ready" in Buffs, and not shown at all for Fury. When the rotation casts Overpower, show a timeline lane "Overpower window" (finding 8).
6. **`created_at`.** Add it. "Saved 18 Sep 2026" is how a player tells an old link from a fresh one, and it is the one date the band needs.

## Round 2 (13596808d): verdict SHIP WITH FIXES

Reviewed: the spec's §13 and §14; every new 1440 board, plus `timeline-390`, the top of `result-390` and `states-1440`; and the new run, `fury-request.json` and `fury-result.json`. The run is the published Fury raid entry: troll, 10,000 runs, 904.4 DPS.

**No blockers remain.** This now reads like a real Fury parse:
- Bloodthirst Rank 4 has 22 casts.
- Windfury Totem procs show as their own "Extra attacks" row.
- Death Wish and Recklessness go up at 2:10 to 2:12, just as the Execute phase opens.
- Rage pools before Execute and then sits near 0 while Execute is spammed.
- The W2 state ("Bloodthirst was never pressed in this run.") is correct and in plain words.

All 23 round 1 findings are answered in §14. The items below are what is left. None of them changes the layout.

### Must fix

**R1. A missing icon ships as an empty box.**
- "Chthonic Power" (1249110) shows an empty box in the Casts lane, the cast list, the On you list and on the phone. "Thrash" (21919) shows an empty box in Up all fight.
- §11.4's icon test exempts "ids on this list while their icon is pending". That lets an empty box reach players, and an empty box is the missing-icon flash the tenets forbid.

Fix:
- No exemptions at merge: the test fails for any key without an icon.
- Chthonic Power comes from a use effect on worn gear. Show the source item's icon and name its item on hover ("Chthonic Power · {item}"), the way the in-game cast bar shows a trinket you clicked. Today a player has no idea what "Chthonic Power" is.
- Thrash stays off the page until lane E names its source (§11.4 already says so). Remove it from the board's Up all fight row until then.

**R2. The Buffs card and Settings disagree about what helps your damage.**
- Settings puts Curse of the Elements under "No effect on your damage". That is correct: it does nothing for physical damage.
- The Buffs card's On the target still lists Curse of the Elements as "Up all fight".
- Faerie Fire, which lowers armor and so helps a warrior, is missing from the Buffs card entirely.

Fix: the Buffs card's On the target uses the same `simbuffs.json` `stats` rule and the same full list as Settings. Debuffs that affect your damage go up top; the rest go in a muted "No effect on your damage" row.

**R3. Character stats show dashes for the two numbers that matter.**
- "Attack power –" and "Haste –". Attack power is the first stat a warrior reads. A dash there reads as broken, not as "illustrative".

Fix:
- The hand fixture fills Attack power.
- Haste reads "0%" when gear gives none. If Flurry is meant to count, write "30% with Flurry".
- If lane E cannot supply `final_stats` by merge, hide the whole block. Do not ship dashes.

**R4. The Target row has no armor.**
- §3.12 and §14 item 7 promise "armor and armor after debuffs". The board's Target row reads only "Level 63 boss".

Fix: draw "Level 63 boss · {armor} armor, {after} after Sunder Armor, Faerie Fire and Curse of Recklessness" from `target_stats` in the board and the fixture. Without it the round 1 finding is only half done.

**R5. The fixture's rage sources contradict the rotation note.**
- "Spent on: Reckless Swings 106". But rotation line 9 says Reckless Swings "costs a quarter of your base health" and makes white swings bring more rage.
- So it should appear under "Gained from", or nowhere, not under "Spent on".
- 106 is also exactly Bloodrage's gain, which looks copied.

Fix: lane E derives both lists from the engine's per-action rage. Add the case to the §10 check 8 fixture test.

**R6. The enchant line has never been drawn.**
- The new run has "no enchants", so the round 1 fix ("+7 Strength", "Crusader") appears on no board.
- A published level 60 raid list with no Crusader and no +Strength gloves or bracers is also something a warrior notices straight away ("BiS with no Crusader?").

Fix:
- Add a gear board from a fixture variant with real enchants (Crusader on both weapons, +7 Strength gloves, +9 Strength bracers, +8 Strength head and legs), at 1440 and 390.
- Send "the published Fury raid entry carries no enchants" to the data and accuracy lane. A player will compare this number against sims that enchant.

**R7. Open question 7 (crit on white rows): keep crits over landed hits, glances included, as Warcraft Logs counts them.**
- Never use crits over non-glancing hits. "90.9%" on Main hand and "100.0%" elsewhere would be screenshotted as a bug.
- The WCL convention still puts white rows (Main hand 42.6%, Extra attacks 43.8%) well above "Crit on the boss 30.9%" in the stats block. A player who compares the two will think one of them is wrong.

Fix: give the Crit header a tooltip, "Crits out of the hits that landed, as Warcraft Logs counts them; misses and dodges are left out."

### Polish

- **Phone lane names.** The names sit in dark boxes over the first ~130px of the plot, so on a scrolled plot they hide ticks and bar starts (Death Wish, Mighty Rage Potion). Fade the box to about 60% or sit the name on the lane's top edge.
- **The talents meta** reads "17/32/2" on the board but "17 / 32 / 2" in the copy table. Pick one.

### Data I do not believe (for the accuracy reviewer, not the page)

- **Berserking (20554)** shows 1 cast and 20.8 s up (11.6%). In 1.12 the Troll Berserking lasts 10 s. Verify the Forever duration.
- **Heart of the Lion (409583)** is in the raid preset. That id and buff are a Season of Discovery hunter rune. Confirm it exists in the Forever 1.60 client before it counts as a raid buff.
- **Hunter's Mark** is grouped as helping a warrior's damage. In 1.12 it raises ranged attack power only. Verify the Forever tooltip; if it matches 1.12, it belongs in "No effect on your damage".
- **The published raid list's own picks** look odd at 60: Skibi's Pendant (Neck · 49), Rune of the Guard Captain (green, Trinket · 51), Blackstone Ring (Finger · 54), and no enchants at all. Forever stats may justify them. The BiS reviewer should confirm.
- **The off-hand miss gap** is now 8.5% against 19.0%, wider than round 1's. It is already with the accuracy loop. I agree with the ruling that the page carries no note until the loop signs off.

### The open questions, round 2

1. **Talent strings:** still lane E's. W2 now covers the dangerous case, so it no longer blocks.
2. **Off-hand misses:** agree with the ruling (accuracy loop, no row note).
3. **`rotation_sha`:** good.
4. **Holy Strength hand tags:** good; they don't arise on this run.
5. **Overpower window:** settled. "Overpower window" as a timeline lane only when the rotation casts it is right.
6. **`created_at`:** good.
7. **Crit definition:** see R7.
8. **`final_stats`:** see R3. If the engine call is not there by merge, hide the block rather than show dashes.
9. **Fixture contradictions:** add R5's rage list to lane E's list.

## Build review (2026-11-05): verdict NOT SHIPPABLE

Reviewed rendered captures only: `shots/int/sbs-1440*`, `sbs-390*`, `full-{390,1024,1440,1920,2000}`, `older-{390,1440}`, and every capture in `shots/lane-a`, `lane-b` and `lane-c`. My crops are in `scratchpad/simview/review/`. I checked numbers against `web/src/fixtures/sim/result-full.json`.

**The player's question:** "How much DPS does this setup do, and what produced it?" At 1440 the answer is the first thing on screen. "904.4 DPS ± 0.5 · 0.05% error" sits top right, and Execute, Off hand and Main hand sit above the fold. On phone the first screen shows the figure, the error and Open in the sim. Every number I checked against the fixture is right: 904.4 mean, 49.5 spread, 731 to 1,116, 10,000 runs, attack power 1,775, hit 9.5%, crit on the boss 29.3%, armor 3,731, the histogram bin "910 to 919 DPS: 795 runs", Deep Wounds 59 ticks at 98.9%, and Bloodthirst 22 casts at 664 rage (30 each). Heroic Strike's 381 rage over 32 casts is 12 each, which is right with Improved Heroic Strike 3/3.

**What is good:** the 1440 desktop page reads like a real Fury parse.
- In Whole fight, Bloodthirst sits on its 6 s cooldown with the dodges drawn hollow.
- Heroic Strike stops at the Execute wash, and Death Wish and Recklessness go up at 2:10 to 2:12.
- Rage pools to 100 before Execute and then sits on the floor.
- Close up's readout and tick tooltip ("2:21.5 · rage after 0", "Next Execute 3.0 s later", "Crit · 2,439 damage") are what I open WCL to find.
- Round 2's fixes are in:
  - Frozen Heart of the Mountain is named, with its icon, in place of "Chthonic Power".
  - Faerie Fire is in the Buffs card.
  - Curse of the Elements sits under "No effect on your damage".
  - The Target row reads "3,731 armor, 471 after debuffs".
  - Haste reads "0%".
  - Extra attacks has its own row, "From Windfury Totem 21".
  - The enchanted variant prints "+7 Strength", "+9 Strength", "+4 All Stats", "Minor Speed Increase" and "Crusader".

Two things sink it:
- **Phone timeline:** it cannot show the pull.
- **Older-result page:** it calls a warrior's rage "Energy".

Either one gets screenshotted.

### Blockers

**B1. Phone timeline: the sticky lane names hide the first ~20 seconds of every lane, including the whole opener.** (`int/full-390.png`, crop `review/p390-timeline.png`; also `lane-b/A7-390-close.png` and `A8-390-whole-scrolled-end.png`)
- The name boxes are opaque, so on phone:
  - The Berserking, Frozen Heart of the Mountain and Mighty Rage Potion bars at 0:00 are invisible. Berserking shows a 4 px stub.
  - The Bloodthirst, Heroic Strike and Whirlwind ticks at the pull sit under "Bloodthirst", "Heroic Strike" and "Whirlwind".
  - Every On the target bar seems to start at a different time, because each starts where its label ends: Faerie Fire looks applied before Judgement of the Crusader.
- Scrolling does not help, because the label stays sticky over whatever is at the left edge. At the end scroll it covers part of Execute, Death Wish and Mighty Rage Potion.
- Checking the cooldown stack at the pull is the first thing a Fury player does with this panel, and on phone it shows nothing. I flagged this as polish in round 2. It is now a blocker, because the panel draws a false picture.

Fix:
- Take the name out of the plot. Print it on its own row above the lane (11px, muted, the lane's top edge), or keep the 44px icon column and show the name in a tap tooltip.
- Nothing may sit opaque over plot content.
- Recapture A7 and A8 at 390 with the pull visible.

**B2. The older-result page calls a warrior's resource "Energy".** (`int/older-1440.png`, `older-390.png`; crop `review/old-2.png`)
- "AVERAGE FIGHT · ENERGY", tiles "515 Gained · 509 Spent", foot "Where the energy came from and went needs a newer run.", and the nav item "Energy", all on "Orc Fury Warrior".
- This is the guild-chat screenshot: "since when do warriors use energy".

Fix:
- Take the resource name from the class and spec, not from a missing or misread `power_type` on old results. A warrior is always Rage, a rogue Energy, everyone else Mana.
- Add a test with an old warrior result.

### Must fix

**M1. The Crit header tooltip contradicts the table.**
- The copy is "Share of landed hits that crit. White swings count glancing blows as landed, which is why they read lower than your crit chance on the boss."
- On this page they read higher: Main hand 42.6%, Off hand 38.6% and Extra attacks 43.8%, against "Crit on the boss 29.3%".

Fix: "Crits out of the hits that landed, as Warcraft Logs counts them. Misses and dodges are left out, so a row can read above your crit chance on the boss." The definition stays as the coordinator decided; only the explanation changes.

**M2. The hit reference disagrees with the page's own miss rate.**
- The stats block reads "Hit 9.5%" against "8% specials · 27% dual-wield swings". On that reference a player expects white swings to miss 27 − 9.5 = 17.5%. The Main hand and Extra attacks rows both say "Miss 19.0%".
- At 300 weapon skill against a level 63 boss, the Classic Era numbers players use are 9% for specials and 28% for dual-wield white swings (the extra 1% is hit suppression). Those match the engine's 19.0% within rounding.

Fix:
- Draw the reference from the engine's own attack table at the character's weapon skill ("9% specials · 28% dual-wield swings" at 300).
- Put the tick on the bar at that value.
- Send the exact figures to the accuracy loop.
- Do not hard-code 8% and 27%.

**M3. Item tooltips are not the client's tooltip.** (`lane-a/A6-weapon-tooltip-1440.png`, `lane-c/A10-tooltip-*`)
- Persuader, equipped in the off hand, reads "Main hand" on the left and "One-Handed Weapon" on the right. The client prints the slot ("One-Hand", or "Main Hand" for main-hand-only weapons) on the left and the weapon type ("Sword", "Mace", "Axe") on the right. "One-Handed Weapon" is not a client line. "Main hand" on an off-hand item reads as an illegal equip. Ironwood Blade has the same problem.
- Persuader's "89 - 167 Damage" and "Speed 2.80" are purple, but Ironwood Blade's are white. The client prints damage and speed in white.
- Persuader's "Binds when equipped" is purple, but Raider Gauntlets' "Binds when picked up" is white. The client uses white.
- The enchant is missing from the tooltip. Raider Gauntlets shows "+7 Strength" on the paperdoll but not in its tooltip, and the same goes for Ironwood Blade and "Crusader". The client prints the enchant as a green line under the stats.

Fix: render slot and type, white damage, speed and binding lines, and the enchant's green line, exactly as the client tooltip does.

**M4. The Berserker Rage note claims rage the Rage card never shows.**
- Rotation line 10 reads "Improved Berserker Rage makes it 10 rage", with 6 casts. That should be about 60 rage.
- "Gained from" lists Main hand, Off hand, Reckless Swings, Mighty Rage Potion, Bloodrage, Unbridled Wrath, Refund and Anger Management. There is no Berserker Rage.
- Either the note is wrong for this client (Improved Berserker Rage is a TBC talent, and I cannot find it in the run's Fury tree on the Talents card), or the rage list drops a source.

Fix: the data lane checks the Forever Fury tree. If the talent is absent, the note becomes the 1.12 wording ("Use it on cooldown for the rage from the hits you take, and hold it for the fears on Magmadar, Onyxia and Nefarian") or whatever the tree supports. If it is present, add its gain to "Gained from".

**M5. The Execute note breaks round 1's ruling.**
- Line 6 reads "First, below 20%: …". §14 item 5 ruled one wording, "in the Execute phase", because Settings says "Last 25% of the fight".
- Line 12 still reads "never from 25% down". Line 3 and line 4 say "execute phase" in lower case, but the wash says "EXECUTE PHASE".

Fix: every note uses "in the Execute phase", with a capital E as the spell name is written.

**M6. "Battle Shout … raid included" is wrong for Classic.**
- Line 5 reads "Whenever it drops: 10 rage for the attack power of everyone within 20 yards, raid included."
- In 1.12 Battle Shout hits party members only.

Fix: "10 rage for the attack power of your party within 20 yards", unless the Forever client tooltip says otherwise. If it does, quote the tooltip.

**M7. Trueshot Aura is under "No effect on your damage".** It appears that way in both Settings and the Buffs card.
- In 1.12 Trueshot Aura raises the attack power of party members, melee included. A warrior in the hunter's group gets it.
- If the engine does not apply it as melee attack power, that is an accuracy bug. If it does, the label is wrong. Either way the page states something false.

Fix:
- Move it to the group that helps your damage.
- Send it to the accuracy loop with the tooltip.

**M8. The Close up readout leaves out buffs that are on.** (`lane-b/A7-1440-pinned.png`)
- At 0:16.5 the readout says "On you: Mighty Rage Potion". The lanes underneath show Berserker Rage (0:08 to 0:18) and Flurry (3 charges) both up at that moment.
- At 2:21.5 it says "Death Wish, Recklessness" while Flurry is up.

Fix: the readout lists every On you lane that is active at the crosshair, cooldowns first, as the spec's example does.

**M9. The older-result page contradicts itself.** (`int/older-1440.png`, crop `review/old-rot.png`)
- The damage table shows "Execute Rank 5 · 6 casts", but Rotation line 6 reads "Execute Rank 5 · Not cast in this run". The rotation's cast match is missing the run's Execute.
- Bloodthirst has 0 casts on a Fury warrior ("0/24/0 talents"), yet Worth knowing shows only the Recklessness note. W2 ("Bloodthirst was never pressed in this run.") was ruled independent of talents, so it must show here.

Fix:
- Match rotation lines to casts by spell family across ranks.
- Fire W2 on older results.
- Add both to the A15 captures.

**M10. Dead space between cards.**
- At 1440 there is about 220px of empty page between Timeline and Gear (`full-1440`, crop `review/f1440-03.png`).
- The older-result page has 300 to 600px holes between almost every card (`older-1440.png`).
- It reads as a page that failed to load.

Fix: release each lazy section's reserved height once it renders, or once it is known to be empty or old. Recapture at every width.

**M11. The phone and 1024 full-page captures are cut off.**
- `int/full-390.png` goes black from just under the Rotation header (about 3,500 CSS px of nothing) to the footer.
- `int/full-1024.png` cuts the Settings card off mid-row, and `lane-a/A1-full-390.png` is the same.
- This is probably the screenshot tool's height limit, but I cannot approve phone Rotation and Settings as integrated from these. I have only the pre-integration `lane-c/card-rotation-390.png`.

Fix: recapture in viewport-height slices, or confirm in the browser that the page renders there.

**M12. "Refund" is not a word a player has for it.**
- "Gained from" lists "Refund 78" with a generic arrows icon.
- Players know this as the rage you get back when Bloodthirst, Whirlwind or Execute is dodged or missed.

Fix: "Refunded on dodges and misses". Use the icon of the ability that refunded most, or drop the icon. Never use a placeholder glyph in a list where everything else has a game icon.

**M13. The Bloodrage note is garbled and talks about levels on a level 60 result.**
- The note reads "… back every minute, a quarter more from your first point of Improved Bloodrage at 59 and half again from 60."
- A player reads "at 59" as a character level and cannot work out what "a quarter more … half again" adds up to.
- The same leveling language is in "From 36, under 80 rage" (line 10), "from 40 it sits right behind Bloodthirst" (line 8) and "Only before level 36" (line 11).

Fix: on a result, print the note for the run's level in plain figures, for example "Whenever rage is under 80: 15 rage at once with Improved Bloodrage 2/2 and 10 more over 10 seconds, back every minute." Keep the leveling clauses on the guide, not here.

### Polish

1. **Copy for the game:** it has no icon and a thinner style than "Copy JSON" and "Download JSON" beside it. Make the three buttons match.
2. **Weapon skill:** "300 · 300" does not say which hand is which. Write "Main hand 300 · Off hand 300".
3. **Cast list:** "Casts: 120 of 120" sits above five rows and a "Show all 120" button, so the count reads as if all 120 are shown. Use "120 casts. Showing the first 5."
4. **Timeline overview strip:** the small rage strip under the Rage lane in Whole fight has no label and does nothing there. Show it only in Close up, where it is the brush.
5. **Item level on a partial kit:** the older page shows "ITEM LEVEL 62.0" from 2 of 17 slots. Show "2 items" or hide it under 17.
6. **Item tooltip header:** "Item Level 65" under the name. The Classic Era client tooltip does not print item level. Keep it muted, or move it to the paperdoll line, which already has it.
7. **Set line (enchanted variant):** "Set: The Gladiator 3 of 6 · (2) …" runs as one sentence. Print it the client's way: "The Gladiator (3/6)", then one line per bonus, with active bonuses green and the rest grey.
8. **Pre-integration icon gaps:** the empty boxes in the lane captures (Deep Wounds in `lane-a/A6-*`, Frozen Heart in `lane-b`, Power Word: Fortitude in `lane-a/A16-lazy-failed`) are filled in the integrated build. Keep the merge test that forbids a key without an icon.

### Data I do not believe (for the accuracy loop, not the page)

- **Armor after debuffs, 471.** Sunder Armor ×5 (2,250), Faerie Fire (505) and Curse of Recklessness (640) take 3,395 off 3,731, which leaves 336. 471 means 3,260 was removed. One of the three is short by 135.
- **Death Wish 18.1% from 1 cast** is about 32.6 s for a 30 s buff. It is the same shape as the Berserking overrun already with the loop.
- **Moonkin Aura** is grouped as helping a warrior. In 1.12 it is spell crit only. If unified crit in Forever makes it apply to melee, fine, but confirm it from the tooltip.
- **Flurry 86.5% uptime** at 29.3% crit on the boss is high for three charges. It is plausible with Execute's 52% crit, but worth one check.

### Verdict detail

I would post the 1440 page in guild chat today. Clear B1 and B2 and M1 to M13, recapture phone and older at every width, and send me the new captures. The layout and the numbers are already there.

## Build review round 2 (land/simview 2a93bf7b6)

**Page: SHIP WITH FIXES** (one must fix, the rest polish). **Rotation note text: SHIP.**

Reviewed:
- `shots/int2/` side-by-sides and the stitched full pages: both fixtures at 390, 1024, 1440, 1920 and 2000.
- Every capture in `shots/fix-layout/` and `shots/fix-page/`.
- `data/curated/apl/warrior-fury.json`, the step notes and `level_notes`.
- The client evidence: `design/reviews/2026-11-03-sim-audit-round12/auras.md` B2, and `design/reviews/guide-accuracy/round19/warrior.md`, lines 48 and 49.

My crops are in `scratchpad/simview/review/r2/`.

### Round 1 findings, as built

| # | Status | What I see |
|---|---|---|
| B1 phone pull | Fixed | Each lane name sits on its own row above its lane (`A7-390-pull`, `A7-360-pull`, `-close`). Berserking, Frozen Heart of the Mountain and Mighty Rage Potion bars all start at Pull, and the pull-time Bloodthirst, Heroic Strike and Whirlwind presses are visible. Nothing covers the plot at 360 or 390. |
| B2 Energy on a warrior | Fixed | The older warrior reads "RAGE", "Where the rage came from and went needs a newer run.", and the nav says "Rage". |
| M1 crit tooltip | Fixed | "Crits out of the hits that landed, glancing blows included. Misses, dodges and parries are left out, so a row can read above or below your crit chance on the boss." This agrees with the table. |
| M2 hit line | Fixed | "Missed in this run: 0% of specials · 13.9% of white swings". 13.9% is the swing-weighted average of 19.0% (63 + 21 swings) and 8.5% (82). It cannot disagree with the rows, and "0% of specials" answers "am I capped" plainly. |
| M3 tooltips | Fixed | Persuader reads "One-Hand / Mace", with white damage, speed and binding lines and the green "Crusader". Raider Gauntlets reads "Hands / Plate" with the green enchant line. See the note on enchant wording below. |
| M4 Berserker Rage | Fixed | Accepted. Improved Berserker Rage is in the Forever Fury tree (round 19 talent text: 5 and 10 rage), and the note states the condition. The run takes 0 points, so no rage in "Gained from" is consistent. |
| M5 Execute wording | Fixed | Every note says "in the Execute phase". |
| M6 Battle Shout | Withdrawn | The Forever client says "all party or raid members within 20 yards" (round 19, spell text). My objection came from 1.12 memory, and Forever is raid-wide. The note is right. |
| M7 Trueshot Aura | Withdrawn | In Forever it is aura 124, ranged attack power only (round 12, B2). "No effect on your damage" is right for a warrior. Moonkin Aura as crit is right too, under Forever's one crit stat. |
| M8 readout | Fixed | At 0:16.5 the readout reads "On you: Mighty Rage Potion, Flurry · On the target: Deep Wounds". I mismeasured Berserker Rage in round 1: its bar runs 0:04.0 to 0:14.0. At 2:21.5 it reads "Death Wish, Recklessness, Flurry". |
| M9 older contradictions | Fixed | "Execute Rank 5 · 6 casts" matches the table. W2 ("Bloodthirst was never pressed in this run. Check the talents in Open in the sim.") now shows on the older warrior. |
| M10 dead space | Fixed | No band over 106 CSS px at any width on either fixture, and that one sits above the footer. |
| M11 captures | Fixed | The stitched captures render Rotation and Settings at every width. |
| M12 Refund | Fixed | "Refunded on a miss, dodge or parry". The arrows glyph is still there (polish 3). |
| M13 Bloodrage and level text | Fixed | `level_notes` prints the level 60 text. The Bloodrage note reads cleanly and matches the Forever talent ("all the Rage … by 25%" per rank). It also matches the run: about 3.5 casts at 30 rage is roughly the 105 in "Gained from". |

**Enchant wording.** In round 1 I asked for "+7 Strength". The build now prints "Strength +7", "Agility +3", "All Stats +4" and "Strength +9" on the paperdoll and in the tooltip. If this is the Forever client's own SpellItemEnchantment text, keep it. Classic-era tooltips do use the "Stat +N" order, and the client's wording beats my round 1 example. If it is generated by the site, use the client's string instead. Either way, it is not a defect once the source is the client.

### Must fix

**R2-1. An icon is missing on the older result's Rotation line 1.**
- "Other cooldowns as they come up · 2 casts · This run: Blood Fury." has an empty icon slot (`fix-page/older-rotation-1440.png`).
- The main run shows the Mighty Rage Potion icon in the same place, so the rule exists but fails when the only cooldown is a racial.

Fix: use the icon of the first cooldown in "This run" (Blood Fury here). Extend the icon test to the older fixture's rotation card.

### Polish

1. **Weapon skill** "300 · 300". Label the hands: "Main hand 300 · Off hand 300".
2. **Talent split format.** The band says "17/32/2", but the Talents card on phone says "17 / 32 / 2". Pick one (open since round 2 of the spec).
3. **Refund icon.** The refund row still uses a generic arrows glyph among game icons. Use the icon of the ability that refunded most, or no icon.
4. **The IN PROGRESS note sits in the Rotation card.** "Being corrected against real parses. Treat the number as a direction, not a figure." It talks about the DPS figure, which is three cards up. Either word it about the rotation ("This rotation is still being checked against real parses, so read the DPS as a direction.") or put it under the figure.
5. **Older pre-pull rows.** They carry no "Cast" or "Not cast in this run" label, while the main run's do.
6. **Carried over from round 1:**
   - "Copy for the game" has no icon and a different style.
   - Cast list "Casts: 120 of 120" sits over 5 rows.
   - "Item Level" sits in the tooltip header.
   - "ITEM LEVEL 62.0" is shown from 2 of 17 slots.

### Rotation notes: SHIP

Every line is true for Forever at 60 and reads as a Fury player would say it. Each one agrees with the run on the page:
- Bloodthirst never casts in the Execute wash.
- Heroic Strike stops at the wash.
- Death Wish and Recklessness go up at 2:10 to 2:12.
- Overpower is dimmed with "Never at 60".

Optional wording:
- **Reckless Swings:** the client says the next four *landed* white swings *within 8 seconds*. Write "your next four white swings that land within 8 seconds each bring 5 more rage".
- **Heroic Strike, general note:** add a comma: "and, once you have Execute, never in the Execute phase".
- **Berserker Rage:** "back every 30 seconds" and "Hold it for the fears" pull two ways. Write "On cooldown under 80 rage, except hold it for the fears on Magmadar, Onyxia and Nefarian: …".

### Still with the accuracy loop (not this page)

These are unchanged from round 1 and still unanswered:
- Armor after debuffs, 471 against my 336.
- Death Wish 18.1% from one cast.
- Flurry 86.5%.

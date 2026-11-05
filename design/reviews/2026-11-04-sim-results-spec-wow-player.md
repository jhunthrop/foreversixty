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

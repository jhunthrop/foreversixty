SHIP WITH FIXES

Surface: /tiers Healers tab, lead under "Tier list at 60", both states. Seen rendered: healer-raid-1440-full, healer-raid-390-full, healer-gear-1440-full (the 1024 and 2000 captures not opened; the lead is the same text). Not seen: the "What the number is / What it leaves out / Ties and ranks" columns, which are not in any capture; I read their copy only, so that part is unjudged on screen.

1. The question: "Which healer should I play for a raid, and what does the number mean?" The table answers the first part at once. The lead answers the second, but not well enough yet: the raid state never says which raid.

2. What a raider would object to, most damaging first:
- Raid state: "Ranked on one healer's share of a logged raid kill, 110 seconds long." Which kill? A raider reads "a logged raid kill" and asks Molten Core, BWL, what boss. Molten Core or Onyxia is only in the hidden notes. "110 seconds long" also sounds like a short kill to anyone who has done Garr or Geddon; nothing says it is a long-kill figure. The sentence is the first thing under the title, so it must carry the answer.
- Raid state: the number is never said to be a share. 131.6 HPS next to a Details! meter that reads 300 will draw "this is wrong" in guild chat. The lead needs "its share" next to the unit, and the header only says "HPS".
- Gear state: "Ranked on the BiS fight, 300 seconds of heavy damage, until the next ranking adds a logged raid kill." "BiS fight" is our word, no raider has met it, and "the next ranking adds" is site talk. It should say it is a stand-in, not a boss, and that it will be replaced.
- Gear state shows "1,028.3" HPS for Restoration Druid. No healer has ever seen four figures of effective HPS; it is the drain-fight number. The sentence under the title has to make clear this is not a raid number, or a screenshot of it will be mocked. Not fixable on the table here, so the sentence must carry it.
- Both states: "5 min" in the meta line against "300 seconds" in the sentence, and "110 s" against "110 seconds". Pick one unit per line.
- Hidden notes, copy only: "as a raid of eleven healers meets it" reads oddly (a 40 man with eleven healers is fine, "meets it" is not how anyone talks). "the group takes the rest" is ambiguous: whose group? Your own five. "fitted to Classic Era Hardcore Fresh logs" is correct but a Forever player may not know the ruleset name; one clause "(Classic Era logs)" is enough in the lead and the full name can stay in the notes. "Phase 1 gear, no world buffs" beside the meta "raid buffs" is not a contradiction but will look like one; say "raid buffs, no world buffs" in the notes.

3. Missing: which raid, that the figure is one healer's share, and in the gear state that it is not a boss. No faction or phase issues on this page.

4. Data I do not believe:
- Raid order Shaman 131.6, Holy Priest 125.6, Disc 119.2, Druid 111.9, Paladin 109.6: the top matches what logs show (Chain Heal in Molten Core). The Paladin miss is the known one, left alone as ruled. Note for accuracy: logs put Druid "well behind", yet the list puts Druid 4th, 2.3 HPS (2.1%) ahead of Paladin and 7 points under Disc. A Druid raider will not complain; a Paladin raider reading "-16.7%" will. Reviewer of accuracy to decide whether the Druid-Paladin gap is outside the error.
- In the raid state there are no "≈ tie" rows. With the Druid and Paladin 2.3 HPS apart and the notes saying a tie is "within the measured error", check that no pair here is inside the error without showing a tie.
- Gear-state order Druid, Shaman, Priest, Paladin, Disc flips to Shaman, Priest, Disc, Druid, Paladin in the raid state. Fine, but the sentence should say the fight changes the order, or a player who saw yesterday's page thinks it is a bug.

5. Verdict: SHIP WITH FIXES. Exact replacement copy (no dashes):

Fix 1, raid lead (healerFightSentence, raid branch):
"Ranked on one healer's share of a Molten Core or Onyxia kill, 110 seconds, fitted to Classic Era logs."
Full lead then reads: "Ranked on one healer's share of a Molten Core or Onyxia kill, 110 seconds, fitted to Classic Era logs. HPS is effective healing, overheal left out. A shield counts for what it soaks, not for the death it stopped by being up before the big hit, and buffs on others, like Power Infusion, are not counted." At 390 that is about 7 lines; the table still starts well above the fold, accepted.

Fix 2, gear lead (healerFightSentence, gear branch):
"Ranked on a stand-in fight, not a boss: 300 seconds of heavy damage that runs every healer dry. A Molten Core or Onyxia kill replaces it with the next update."
Keep the numbers on the page; the sentence now says why they are in the four figures.

Fix 3, one unit per line: make the meta line "110 s" and "300 s" to match the sentence, or the sentence "five minutes"; do not mix.

Fix 4, hidden notes, "What the number is" (raid): 
"Effective healing per second over 110 seconds: one healer's share of a Molten Core or Onyxia kill. The tank's share lands in spikes and the healer's own group of five takes the rest. Fitted to Classic Era Hardcore Fresh logs, Phase 1 gear, raid buffs, no world buffs. Overhealing is not counted."
(drops "as a raid of eleven healers meets it"; keep ${gearSentence} after it.)

Fix 5, add to the raid lead or the notes, one clause a raider will ask for: the figure is a share, so it is lower than what your meter shows. Notes line: "It is a share of the fight, so it reads lower than your own meter."

Re-review after fixes: look at 390 and 1440 raid state again, and the notes columns open (they were not in the captures).

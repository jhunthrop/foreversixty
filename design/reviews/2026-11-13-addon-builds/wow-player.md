# wow-player review: addon Builds pane and panel Talents row (2026-11-13)

Boards seen: builds-hunter-42, builds-hunter-60-raid, panel-talents-42 (rendered). Data: addon-builds.json, every spec's why read.

Verdict: FIX (the layout is right; the whys and two wording points are not).

## What works
- A level 42 hunter sees at once: Beast Mastery, "Level 60 build", gold "Following", "Spend 1 now: Intimidation". The next point is the first thing under the followed row. Good.
- Splits read right: BM 31/16/4 and 31/20 raid, MM 31/15/5 and 33/13/5, Survival one row "Level 60 and raid build". Biggest tree first, as the talent frame lists it.
- Talents row: "Intimidation · Spend 1 now" on the full line, buttons below, nothing cut, "Change" has its label. The old blank button is gone.

## Must-fixes
1. "Level 60 build" is the wrong word for a level 42 player. Players say "leveling build" and "raid build". A 42 reading "Level 60 build" does not know it is the one he follows while leveling. Use "Leveling build" (the site's own planner label, if it says otherwise, fix the planner too), or "Leveling build (to 60)". Keep "Raid build".
2. Panel says "27 of 33 points · 5 points off plan" but the pane's heading says only "Level 42 · 27 of 33 points". That reads as six points unspent; it is 27 on plan, 5 off plan, 1 to spend. Show the off-plan count in the pane too ("27 of 33 on plan · 5 off plan"), or say "27 of 33 points on plan".
3. The why line is cut where the news is. Hunter BM: "all five points of Efficie...", Marksmanship: "take t...", Survival: "take every melee damage talent, ...". Rewrite whys to lead with the point of the build, short enough for one line (about 60 characters), full text stays in the tooltip. Whys must not be a first-sentence grab of prose.
4. Dangling "it/they/also/the rest" (a player reads each row alone). Rewrite these:
   - druid-feral-bear (both): "It tanks from level 10 to 60 with no respec."
   - druid-restoration (both): "It is built for a tank-healing fight that runs your mana down."
   - paladin-holy (both): "The rest is utility. Guardian's Favor takes 2 minutes off..." starts mid-thought.
   - paladin-retribution leveling: "The leveling build never casts it:" (casts what?)
   - mage-fire raid: "Combustion asks for 30 points above it. The Fire damage talents above it..." (two "it", no verdict).
   - mage-frost leveling: "...puts three of them into Arcane Meditation" (three of what?); mage-frost raid: "In a raid they go into Fire..." (they?).
   - warlock-destruction leveling: "The leveling build also takes..." (also).
   - warrior-fury (both): "The build also puts two points in Improved Berserker Rage... comes free" (also, and "comes free" is nonsense out of context). Raid row repeats the leveling text.
   - warrior-protection (both): "The Arms points go to Deflection..." (which Arms points).
   - shaman-restoration (both): "Thundering Strikes is 5% crit on your heals, because..." opens on a talent with no build named.
5. Useless whys (say nothing a player does not see from the split): mage-arcane raid "puts 32 points in Arcane and takes its second tree from Fire"; paladin-protection "Toughness (10% armor from items) and a point of Divine Intellect fill the tree" (filler talents, not the point of the build); rogue-assassination leveling opens "The level 60 build is the one you level with"; priest-discipline/holy "Deep Discipline for..." fragment with no verb, same line on both kinds.
6. Raid vs leveling whys that do not say what changed: druid-feral raid explains Furor mechanics, not the build; rogue-combat/assassination raid lines restate the split ("goes 31/20/0", "12/34/5") directly under the split line. Lead with what the raid build gives up or gains ("Drops Iron Will and Deflection for ...").
7. Single-build specs repeat one sentence for both kinds (fine, one row), but the sentence must stand alone: rogue-subtlety "One build does everything:" is good; survival hunter should open with the spec, not "Thirty-three points in Survival...".

## Missing
- Nothing says which row is "mine to follow at 60". A 42 following Leveling should see once "Raid build at 60" is a click away; the chat line covers it, the pane does not.
- No faction/race note is needed here; fine.

## Data I do not believe (for the accuracy reviewer)
- Priest "Penance", "Divine Aegis", "Prayer of Mending" are not Era talents. Confirm they are in the Forever client's priest tree. Same for hunter "Sniper Shot", "Lone Wolf", warlock "Wrack", "Malevolence", "Soul Siphon", "Bane and Aftermath for a faster Immolate", druid "Eclipse", "Nature's Grace once you reach it". If any are not in the client, the why is wrong.
- Hunter Marksmanship playstyle "Barrage and Sniper Shot" (heading line): confirm both are MM talents here.

## Panel Talents row
Fine as drawn. One note: "5 points off plan" in red with "Spend 1 now" in gold on the same row needs the pane to agree (must-fix 2).

## Round 2

Verdict: FIX (three whys say the wrong thing; layout and labels are now right).

Seen: builds-hunter-42.png, panel-talents-42.png. Round 1 items 1, 2, 3, 4 and 5 are closed: "Leveling build"/"Raid build", the pane and the panel now agree ("27 of 33 points · 5 points off plan"), "Intimidation · Spend 1 now" reads like the client, every why fits one line, none dangle. Checked all 45 whys against the build orders in addon-builds.json, the Forever talent trees and each guide's talents prose.

Must-fix (wrong against the build code):
1. hunter-beast-mastery raidBuildWhy "Trades Efficiency for Improved Stings and Mortal Shots". The raid order takes Improved Tracking 4 -> 0, Efficiency only 5 -> 3, Mortal Shots 1 -> 5, Improved Stings 0 -> 2 (guide prose agrees). Replace: "Drops Improved Tracking for Mortal Shots and Improved Stings" (60).
2. hunter-beast-mastery buildWhy "Efficiency and Improved Tracking keep you killing between drinks". Improved Tracking is a damage bonus on the tracked type, not mana. Replace: "Efficiency keeps you shooting between drinks, Tracking adds damage" (64; or "Efficiency saves mana, Improved Tracking adds damage").
3. warrior-arms raidBuildWhy "Trades Iron Will and Improved Charge for Reckless Swings". The raid order puts Enrage 0 -> 5, Improved Execute +2, Improved Heroic Strike +2, Reckless Swings only +1; the guide says Enrage is the way to Reckless Swings and Execute. Replace: "Iron Will and Charge go to Enrage, Execute and Reckless Swings" (61).

Should-fix:
4. druid-restoration (both): "Bloom and Reflection to outlast the fight". Bloom is a heal (your Spirit when a HoT expires), not mana. Replace: "HoT tank healer: Bloom heals on expiry, Reflection saves mana" (61).
5. hunter-marksmanship leveling: "Rapid Recuperation refills mana" is 50% regen while casting, not a refill. "Rapid Recuperation keeps mana up, Rapid Killing chains kills" (59).

Checked and correct: all other whys match their orders (Champion of the Light lands at 35-39, Lightning Mastery at 40, Wrack at 40, Barrage 3/3, Aftermath/Cataclysm full, Feral Charge matches the guide's Cat Form leap, Demonic Sacrifice on the Imp matches the guide, Bastion is 10% damage and the guide calls it threat). No Classic-memory disagreements with the Forever tree data.

## Round 3

SHIP. All five replacements are in data/builds/1.60.1.70790/addon-builds.json, each 60-62 chars (under 64). Boards builds-hunter-42, builds-druid-24-site, builds-paladin-13-ask and builds-hunter-60-raid render clean: no truncation, Following/Spend 1 now/off-plan states read correctly, the paladin question heads the list, and "and 3 more" fits. The hunter and druid why lines now read as real play advice. Not checked: whether Improved Tracking and Efficiency are in the BM talent tree at those counts; the accuracy reviewer owns that.

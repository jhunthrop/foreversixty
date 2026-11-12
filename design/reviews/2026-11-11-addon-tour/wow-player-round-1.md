# wow-player, round 1 (boards in this folder)

**Verdict: SHIP WITH FIXES.** The fix for the owner's complaint is right: a 2 Holy / 2 Retribution paladin is asked "Are you Holy or Retribution?" and nothing is pre-picked. The leads read like a player wrote them. A few false or unverified claims first.

## Must fix

1. `specPlaystyle["paladin-holy"]` "Plate healer, Holy Shock and Holy Light" under "Level 13 Paladin": paladins wear plate from 40 in Forever. Drop "Plate".
2. "Twist of Light" (paladin-retribution) and "Berserk" (druid-feral) are not in Data.lua's talents (the greps find only Berserker Rage, Berserker Stance and items). Twist of Light is a real Forever talent in the research notes and the sim, so the addon's talent data is behind it or it lives only in the hotfix overlay. Do not ship either name until the accuracy reviewer confirms it against the client data.
3. `builds-dialog-druid`, the pushed build row carries the paladin crest on a druid's dialog.
4. `builds-dialog-60`, "Paladin raid builds": the rows total 50 points (31/11/8, 33/10/7, 31/12/7), and the druid rows too, where a 60 has 51. Accuracy should check the data.

## Should fix

1. Own-build rows have no icon and no role; on the druid board "Feral build" and "Feral Bear build" differ only by split.
2. The totals on the level 24 boards are end-of-leveling targets under "leveling builds"; a unit ("50 points") would help.
3. "Next talent: Deflection" for a Retribution follower at 24 is plausible only if the build takes Protection early; the accuracy reviewer should check, and the line lacks the tree.
4. The question names two trees above three tiles; right for 2/2, and at level 10 with no points "Are you Holy, Protection or Retribution?" is exactly right.
5. "path" is not a Classic word ("Pick a path", "follow your path"); say spec or build, the same way everywhere.
6. "loads its leveling build" reads as a fragment.
7. The chip card's three solid yellow bars read as all selected.
8. The page 1 lead promises the arrow now while the panel for the same character says picks start at 20.
9. The page 3 lead repeats its first row.

## What reads right

The leads ("Your next talent gets a gold ring and a better bag item gets a green arrow", "your parse on each boss, what cost you most and what to fix first", "the line by your minimap says Tonight"); one gold Following; the heading and the hidden Enter hint; the panel's Pick rows; every other playstyle line checked against Data.lua, the sim and the client tables (Holy Shock, Holy Shield, Consecration, Leader of the Pack, Moonkin Form, Wild Growth, Bestial Wrath, Barrage, Lacerating Strikes, Blast Wave, Fingers of Frost, Prayer of Mending, Penance, Mutilate, Seal Fate, Blade Flurry, Lightning Overload, Lightning Mastery, Maelstrom Weapon, Siphon Life, Soul Link, Fire and Brimstone, Sweeping Strikes, Bloodthirst); the shaman line carries no second weapon; the melee survival line matches Forever.

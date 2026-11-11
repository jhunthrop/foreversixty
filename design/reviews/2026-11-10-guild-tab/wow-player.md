# wow-player verdict: the Guild tab (apps), 2026-11-11

Seen: iOS parity full renders (at 60, officer, under 60), the at-60 and no-guild top renders. Not opened one by one: the android/ and light/ walk captures (the interactions file says the light set is the dark app unchanged; the Android parity renders are listed as the same views). Judged on what I saw.

## Verdict: NOT YET (close; the fixes are small and mostly data and wording)

The tab answers the question a raider has the night before a pull: am I in, what is my job, who is missing, what do I bring. Your job is the first thing under the sign-up and the words are the addon's. The class colours, the Fort / WF icons and the red/green counts read like a raid tool. The items below keep it from shipping.

## Findings

1. Class colours disagree between sections. Kaelra is a Fury Warrior (warrior tan) in Roster, but light blue in the Damage line and in Assignments (whelps west). Grunvald is a Protection Warrior (tan) in Roster and orange (druid) in the Tanks line. A raider notices a name changing class colour. Fix: one colour lookup from the roster's class for every name on the tab; if it is fixture data, correct the fixture so the parity goldens prove it.

2. Flask of the Titans under "Before you pull" for a Damage viewer ("As Thandrel, Damage"). The Titans flask is the +health tank flask; a damage dealer would want Mongoose or a damage elixir. It also shows as a red X, so the tab nags a DPS about the wrong item. Fix: the consumable list follows the viewer's role and spec (tank: Titans, Fire Protection; melee: Mongoose, Juju; caster: their own); never show a red X for an item the role does not use.

3. The invite note contradicts Your job. The note says "Whelps go to the left wall", Your job says "west cave". On Onyxia there are two caves, not a left wall. Fixture text, but it is the first thing a player compares. Fix: fixture note should read like a real leader's note ("Whelps west and east, tanks call Deep Breath on voice"). Also the note cuts off mid-sentence ("the tanks call the deep breath in voice;..." / "cons…"). Fix: tap to expand, or a "More" at the cut.

4. Thornwick is listed twice: in "Recently on" (57 Feral Druid) and again in "Close to 60". Fix: a character appears once; Close to 60 skips those already in Recently on.

5. "Raids open at 60 · you're 36" sits under a live In / Late / Bench / Out bar. The ruling says sign-up is allowed and the line is quiet, so this is not a block. But the line says what is true and does not say what the leveler wants: how far from the Onyxia attunement. Fix: under 60, add one line "Onyxia attunement starts at 50ish: <quest giver, zone for the faction>" or the Drakefire Amulet state; at 60 the Attuned row is already there. (Check the exact chain level in the Forever client, I am not sure of the Forever number.)

6. "Missing Fort · WF 1 of 2" is shorthand the addon uses; a veteran reads it, but "1 of 2" has no unit. Fix: "Fortitude 1 of 2 groups"? Keep the icons and state what the count is (casters present / needed, or groups covered), same in the tooltip on long-press. A bare "WF" for Windfury Totem needs "Windfury" in the accessibility label at least.

7. "Preview mention" (blue button, Officer / Needs you) is web jargon. A raid leader says "Ping the 5 who have not answered". Fix: label "Ping the raid" with the line above it saying who; keep Preview as the sheet's title.

8. Officer Assignments lists only filled slots (5 of 10 or so). A leader wants the empty ones: Off tank healers, Raid healers, East cave whelps, Deep Breath callers do not show. Fix: show every template slot, empty ones as "0 of 2" in red, like Who is coming does.

9. Loot is a bare line: "Last raid: 4 pieces", and the at-60 page has "Loot history, Last raid: 4 pieces" in Pages. A looter wants the items, with quality colours and icons, and who got them. Fix: show the last raid's drops as item rows (icon, name in quality colour, winner); the count line is not content.

10. "Progression: First tier, 1 of 1 down · 6 pulls this tier" says nothing in game words. Fix: "Onyxia's Lair 1/1" with the boss kill, and the next target at launch (Molten Core when its phase opens). Also "Guild sims" (at 60) and "Loot history" (under 60) swap places in Pages for no reason a player can see; pick one set.

11. Recent logs data I do not believe: "Scholomance · 9/9 kills" (Scholomance has more bosses than nine in Classic), "UBRS · 7/8 kills", "Barrow Deeps Sat" in Calendar (not a Classic instance I know; confirm it is a real Forever instance or change the fixture to a real one such as Blackrock Spire or Zul'Gurub). Fixtures are what the owner screenshots, so use real instances and boss counts.

12. Onyxia template facts for the accuracy reviewer: "Off tank ... takes her when Bellowing Roar or Knock Away drops the main tank" (Roar is a fear, it does not drop the tank; Knock Away drops threat). "Fear Ward on the main tank" is a Dwarf priest racial in Classic; Forever may be different for other races, so check which Mirethil can cast it before the page names a priest for Fear cover. Fear Ward must be recast, the Purpose line should say so.

13. Missing that a raider always has: reset or lockout (Onyxia is a 5-day lockout in Era), a one-tap "ready check" status (are you buffed, repaired), and the gear gap to raid BiS at 60, which is the site's whole job. Suggest "Your gear vs raid BiS: N slots" under Before you pull linking to BiS.

14. Smaller: "Roster ›" appears twice at 60 (Who is coming and Roster) with the same label and different targets; label the first "Raid roster". "36 you" in the spread reads as a score; use "you · 36". No-guild "Claim your guild: Claim it: a guild master..." repeats the verb and the line is long. "Applied Mon" sits inside a nested outlined box that Dawnbreak does not have; the two recruiting rows look different. Mirelle (applicant) and Mirethil (healer) are one letter apart in the same fixture; rename one.

## What is right
Four answers In / Late / Bench / Out, Your job in the addon's words ("Whelps, west cave with Kaelra and Thalrin", "Onyxia · purpose"), Alliance and realm type on the guild head, class-coloured names, Tanks 2 of 2 / Healers 5 of 8 / Damage 11 of 30 split for a 40-man (adds to 40), Drakefire Amulet as the Onyxia attunement item, Missing buff icons, quiet under-60 line.

## Ship when
Findings 1, 2, 3, 4 and 8 are fixed (colours, role-correct consumables, note/job mismatch and truncation, duplicate row, empty slots), and 11 and 12 pass the accuracy reviewer. The rest are SHIP WITH FIXES material.

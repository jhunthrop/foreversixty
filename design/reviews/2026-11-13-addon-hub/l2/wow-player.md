
## Round 1

Reviewer: wow-player. Seen rendered: me-raid-day, me-no-pack, me-no-companion, me-stale-no-guild, me-new-character, me-sim-queued.

**Question the tab answers:** "How am I doing, and what do I do before the next pull?" The score and the last-raid table are the first things on screen, and Right now sits beside them. Good.

**What reads right.** The parse colours match what a Classic player knows from logs: 58 and 64 blue, 77 to 92 purple. The maths holds: the five score parts add to 1184, the median of the eight parses is 74, the one death matches the Died column, and best/worst match the table. Tue 8 Dec and Mon 7 Dec are the right weekdays. Hunter at 60 with 51 of 51 talent points is right. "No flask" in red, "All up" in grey and "Ratings 8 days old" in red are all clear. Every Right now line carries its tab name in yellow (Guild, Gear, Rotation, Builds). Gain icons give the client tooltip and shift-click.

### Verdict: FIX

### Must-fix
1. **Level 8 hunter is told "Pick your spec" (me-new-character).** Talents open at 10, so a level-8 hunter has no spec to pick. The line, and the yellow dot on Builds, send a player to a tab with nothing to do. Say what the leveling build is instead, or show no line until 10. The row also has no icon, so the text sits indented with a hole where the icon goes.
2. **A level-8 character is shown "Last raid / No raid yet. The companion app brings it in."** That is a raid pitch aimed at someone with no raid. Below level 60, or with no raid history, drop the section, or use the space for something that helps leveling (the Builds line is enough).
3. **Gain icons carry no names and no slot.** "+11.2 DPS", "+9.0 DPS", "+7.9 DPS" next to icons only. Wowhead and Sixty Upgrades always print the item name. The first gain duplicates Gear's Right now line (gloves) but the other two are never named without hovering. Print the slot or item name under each icon, or give each icon its tooltip name inline.
4. **"Experience 380 of 700" is the wrong word.** A player reads Experience as XP. Rename to "Raid experience" or "Progression" (whatever the score actually measures; the site name must match). Same for "Preparation 46 of 100": the only visible reason is "2 flasks missed", which cannot explain 54 points lost out of 100 with 2 flasks missed in 8 kills. Either show every reason or state the rule.
5. **Rotation line "Garr: 74% followed, 31 s downtime" is unclear.** Followed what? Say "74% of your rotation". The second line "Your weakest boss on Tuesday" says nothing a player can act on; name the one fix ("Hold Multi-Shot during the Firesworn phase" style), or link straight to the Rotation tab's fix.
6. **Footer says "No companion" but the sim section silently disappears (me-no-companion).** The right column is half empty and the player is never told that Sim again needs the companion. Add one line under Right now or where Your sim would be.

### Data I do not believe (accuracy reviewer to confirm)
- "Aimed Shot late 9 s" on a Beast Mastery hunter. Aimed Shot has a 6 s cooldown, so "late 9 s" is more than one full cooldown, and BM does not lead with Aimed Shot in most Classic raid builds. Check the engine's BM rotation uses it, and check the unit "late" means.
- "Hunter's Mark late 6 s" and "Serpent Sting off 4 s" are fine as phrases (applied late, dropped for 4 s). Keep.
- "Gloves of Holy Might" as a +11.2 DPS pick for a hunter, shown as "Hands · in your bags". Confirm it is a real item that a hunter can wear (armour type and proficiency), has agility/attack power, and is sourced. The name reads like a paladin or healer item.
- 941.2 DPS for a hunter in Molten Core gear looks high next to what players actually log; fine if it is single-target sim DPS, but label it that way (see 7).

### Polish (optional)
7. Say what "941.2 DPS ± 3.1" is: "Sim DPS, single target, level 60 raid build". A number without its basis will draw questions in guild chat.
8. Onyxia line: "20:00, in 3 h 12 min" has no timezone ("20:00 server time"). "In" alone reads oddly; "You're in" or "Signed up" is clearer.
9. me-stale-no-guild: header drops "Beast Mastery raid build" while Right now still names it. Show the spec in both or neither.
10. After "Queued" the player has no idea when it finishes or that the companion must be running. Add "Companion picks this up" or similar.
11. me-no-pack: "No ratings pack. Forever Sixty Data, on CurseForge or Wago." Say plainly "Install the ForeverSixty_Data addon" (use the real addon name).
12. Score has no cap shown (parts add to 2100, Epic at 1500). "1184 of 2100" in the sub-line would help.

## Round 2

Seen rendered: me-raid-day, me-new-character, me-no-companion (re-read).

**Verdict: SHIP.** Round 1 items 1, 3, 5 and 6 are fixed or properly handed off.

- 1 Level 12 hunter: "Pick your spec" is now right (talents open at 10).
- 3 Biggest gains names each item in its quality colour (Gloves of Holy Might, Forest Stalker's Bracers, Windreaver Greaves) with its DPS. This is what Wowhead and Sixty Upgrades do. "Simmed Mon 7 Dec" is clear.
- 5 Rotation line now reads "Tue 8 Dec · parse 58". Accepted as the Rotation tab's own line (L4). "74% followed" is still vague, so pass that to L4.
- 6 No companion: the Last raid section carries the app address. Acceptable per spec 8.
- 2 Below 10 the Last raid section is gone, accepted. At 12, though, the render still shows "No raid yet. The companion app brings it in." (see polish A).
- 4 Accepted: Experience and Preparation are the site's own score labels. I disagree with "Experience" as a word (it reads as XP). Raise it site-wide to the owner, not as a blocker here.

### Polish (not blockers)
A. Hide Last raid until there is something to show, or at least below 60. A level 12 does not raid, and the line is an app pitch.
B. The "Pick your spec" row has no icon, unlike every other Right now line.
C. Still open from round 1: label the 941.2 DPS ("single target"), add a timezone to "20:00", and carry the spec in the header on the stale view.
D. Accuracy reviewer: "Aimed Shot late 9 s" on BM (6 s cooldown) and "Gloves of Holy Might" legality for a hunter are still unconfirmed; the gear names now on screen are real items.

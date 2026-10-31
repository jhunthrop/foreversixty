# Player verdict: mobile Raids integration (integration C), 2026-10-31

Reviewed: Android and iOS captures of the raid composition, the character selector on Home, Raids and BiS (Android at 1.0, 1.3 and 2.0 text, iOS at default), and the officer and reminder-ask goldens in both snapshot folders.

Question this answers: "Which of my characters am I looking at, and is my guild's raid sorted for Saturday?" The selector is the first thing on screen on all three tabs, and the raid card answers the sign-up question right under it. That part works.

## Verdict: HOLD

## What is fixed (checked on screen)
- A missing buff no longer reads as a dark hole. In the integrated captures "Bless" sits dimmed beside bright, green-edged real icons, and "Missing: Blessings" underneath says it in words. A leader gets it at a glance.
- Role and class counts add up. In the officer golden I counted the roster rows by hand: Warrior 5, Hunter 2, Rogue 2, Priest 4, Shaman 4, Mage 2, Warlock 2, Druid 3 = 24, and 3 tanks, 6 healers, 14 DPS matches the rows. Paladin 0 is dimmed, which is right for a Horde guild with no Undead Paladin.
- Sentence-case buttons ("Lock sign-ups", "Copy the list", "Turn on reminders") read fine and match the rest of the app.
- Android and iOS composition cards match each other. Only the wrap point of the buff row differs (8 then 2 on Android, 7 then 3 on iOS).

## Blockers
1. Nothing shows that picking a character changes which guild's raids appear. Every capture has the same single character (Obnoxious Yell, Olympus XXVII). There is no capture of the open selector list and none of a second character in another guild. The raid leader's whole question on this build is "I switched to my alt in the other guild, which raid am I looking at?". Capture the open selector with two characters from two guilds, then the Raids tab after the switch, on both platforms.
2. Home and BiS tell the same level-58 Fury Warrior two different things. Home says "Best in slot for 60, together" and labels the list "60 · raid". The BiS tab for the same character opens on "50 to 59". The Home list is not raid gear either: Legionite Glaive and Hyper Deluxe Sniper Rifle are crafted, Beaststalker's Pants and Boots are Stratholme, so "raid" is wrong. A player will ask "am I gearing for 59 or for 60?" Make Home's band and its tag the same as the band BiS opens on, or say plainly that it is a look-ahead.

## Shoulds
- The live buff row is ten icons with no names. In the no-network goldens every tile carries a label (Fort, MotW, Int, Bless, SPro, SS, Rez, WF, SoE, MS), but on a live phone present buffs are bare icons and only MotW shows text. Nobody can tell Fortitude from Shadow Protection except by memory. Tap or long-press for the name, or a short label under each icon. Check that the 7th and 8th icons (red flame figure, pink torch) match "Rez" and "WF"; I could not match them by eye.
- MotW showing a text tile among real icons reads like a missing buff. It is bright with a green edge, so it is distinguishable, but until the icon is exported it looks like a mistake. Ship the icon, or dim-vs-bright is the only cue.
- The integrated Android composition card shows "Tanks 3  Healers 6  DPS 14". The officer golden shows "Tanks 3/4  Healers 6/8  DPS 14/28". The targets are what a leader reads (are we short a healer?). Show the targets on the live card too, or drop them from the golden. They should not differ.
- Buff gaps cover only the Blessings case in the live captures. A leader also checks Battle Shout, Leader of the Pack and curses, but that is a later addition.
- iOS: the floating tab bar sits over the content and the rows show through it as ghost text ("Insp...", "Neck, Shoulder, Ring, Trinket"). Give it a solid backing.
- "NOT NOW" in the reminder ask is still all caps beside "Turn on reminders". Make it "Not now". The same goes for the "NOT SIGNED UP" and "PLANNED" pills if you want strict sentence case, though pills as caps is defensible.
- Android Home at 1.3 and 2.0 text drops the "Fury · 58" line under the name; Raids and BiS and iOS Home keep it. It is the same selector everywhere only at 1.0. At large text keep the level line or truncate, but do the same on every tab.
- At 2.0 text Home wraps "Orc Fury Warrior · Level 58 · OLYMPUS XXVII" over two lines and the "UPGRADES FOR YOUR LEVEL RANGE" heading over two lines. It is readable but the first screen is mostly header.
- The guild header icon differs by platform on Raids: a star on Android, a shield on iOS. Pick one.
- Monospace creeps into body text: "11  signed up" (Android) and "11 signed up" (iOS) on the Barrow Deeps card are set in a code font. The times "6:45 pm" are a heavier mono than "pm". Use the numeric font for the digits only.
- Fixture slip in the reminder-ask goldens: "20/40 signed up" with every class count at 0 except Warrior 2, so the buff line says all ten buffs are missing. A player seeing this would think the card is broken. Use a roster that matches the count.
- Spec icon beside "Fury · DPS" is an empty dark square in the offline goldens. Expected offline, but show the class icon as a fallback.

## Data I do not believe (for the accuracy reviewer)
- "Hyper Deluxe Sniper Rifle Mk XVII", "Replaces Dark Iron Rifle, ilvl 55", +20.2 DPS on a Fury Warrior. Warriors can use guns, but is it a Forever item, and is Dark Iron Rifle really what is replaced?
- "Legionite Glaive", "Replaces Enchanted Azsharite Felbane Sword, ilvl 60" for a level 58: is the Felbane Sword equipped at 58, and is Legionite Glaive a real Forever blacksmithing craft?
- "Dawn Armor" (Leatherworking chest, ilvl 63) replacing "Shadowcraft Tunic" at +9.8 DPS for a warrior: leather below warrior's armor type is legal, but a leather chest beating plate for Fury needs a check.
- "Beaststalker's Pants/Boots" as mail for a warrior at 58 is legal (mail before plate at 40), but check the DPS (+16.1, +10.1) against a plate alternative.
- BiS tab: "Stone Guard's Plate Armor ... Vendor: Lady Palanseer" and the helm "Ornate Mithril Helm" in a 50 to 59 band, with Skibi's Pendant and Dark Phantom Cape as "World drop". I did not check them.
- "182.0 DPS at best in slot" on BiS against "+102.5 DPS" of upgrades on Home are different quantities; make sure they are both from the same sim.

## Bottom line
The buff row fix is real and readable, the counts add up, and the buttons are consistent. Two things stand in the way: there is no evidence the selector switches guilds, and Home and BiS disagree about which gear band a level-58 warrior is in. Fix those two and re-capture, and this ships with the shoulds queued.

## Re-check (after selector-open-raids and selector-raids-switched captures)

Blocker 1, selector changes the guild: cleared.
- Both platforms show the "Your characters" sheet opened from the Raids tab, with Obnoxious Yell (Orc Fury Warrior, 58, OLYMPUS XXVII, ticked) and Yellfury (Orc Hunter, 31, No guild).
- After picking Yellfury the chip changes to the hunter and the page becomes "NO GUILD YET" with "Yellfury is not in a guild on Forever Sixty." A raid leader sees the raid view follow the character. It would be nicer to see a second guild's raids, but the fixture has one guild, so this is accepted.
- The sheet reads like the client's own character list: crest, class-coloured name, "Orc Hunter · Level 31", guild or "No guild".
- New should: on iOS the chip truncates the second line to "Hunt..." while Android shows "Hunter · 31". Shorten it or let the chip grow.

Blocker 2, Home "60 · raid" versus BiS "50 to 59": moved off this integration. It lives on Home (F2 and B) and BiS (F3), not in the Raids merge, so it does not hold this ship. It must stay on the Home and BiS owners' list, with two parts: the band must match what BiS opens on, and the "raid" tag is wrong for crafted and Stratholme gear. A player will meet it on the first screen after install.

"Not now" sentence case: claimed fixed on Android. I did not see the re-recorded goldens, and the iOS golden is still being recorded. Check both once they land.

All earlier shoulds stand. The buff names on live icons, the missing role targets on the live card, and the iOS tab bar showing content through it are the ones I would fix first.

## Final verdict: SHIP (Raids integration only), with the shoulds above and the Home/BiS band mismatch tracked separately

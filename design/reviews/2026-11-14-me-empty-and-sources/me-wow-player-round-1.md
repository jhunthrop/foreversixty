# wow-player, round 1 (Me with no site activity)

Verdict: FIX

I judged from the five captures and the before shot only. I did not read the spec or code, so anything about clickability or copyability is from the screenshots.

**Must-fix**
1. **"Your sim" with a "Sim again" button, but no result** (l25-point, l60). This is the no-sim state, so nothing has run yet. "Again" is a lie, and a heading with only a button reads as broken. Say "Run a sim", or hide the block until a number exists. If the companion can't run one in-game, point the button at the site instead.
2. **Red "?" icon on step 2, "Hover an item you wear"** (l5-setup). That is the client's missing-texture icon, and players will mock it in guild chat. Use a real item icon, such as the chest slot or any equipped piece. The step also doesn't say what the hover shows.
3. **Talent plan count is off by one** (l25, l26). Level 25 means 16 talent points, but the Me tab says "15 of 15 on plan". At 26 a point is waiting ("Next point: Now"), so there should be 17 points against a plan of 15. Either the plan is short or the label is wrong. The "Next point: Level 26" row only reads right if the 15 counts correctly.
4. **"Level 60 / Max level"** (l60). Every player knows this, and the header already says Level 60. Drop the row.
5. **"Your score and raid logs show here after your first logged raid"** shows on levels 5, 25, 26 and 60. Level-25 characters aren't raiding, so it is noise there. At 60 it should name the one action: "Log a raid with the companion". Keep it off levels 5 to 59, or tie it to the single call to the site.
6. **The call to the site reads vaguely.** "Once the companion app brings them in" doesn't say what the app does. Write something like "Get the companion app: foreversixty.gg/app". Also check the URL can be copied, since a plain line of text can't be clicked in the game client.

**Shoulds**
7. **"8 of 14 best in slot" / "9 of 16 best in slot"** is the wrong phrase. At level 25 this is the band's gear list, not best in slot. Use "8 of 14 on your level 20 list" or similar. A level-25 or level-60 character also has more than 14 or 16 slots (ranged, ammo, rings and so on). Say what is counted.
8. **"Level 30, a new gear list"** is vague. Write "Level 30: gear list for 30", or drop the row and let Gear speak.
9. **"Rested 40%" and the XP bar** repeat what the client's own XP bar tooltip shows. They don't move the player toward a better character. I'd drop Rested. The "Set up" view already hides it while the same character shows "Rested 25%" in the other view, which is inconsistent.
10. **"2 enchants missing Chest, Wrist"** has no separator, so the red text and the slots run together. Add a colon. At level 25 an enchant nag is also noise, since few hunters have an enchanter. Hide it below about 40, or say where to get it.
11. **Gear rows have no source.** Wowhead always shows where an item comes from. A bare "Ranger Bow +3.9 DPS" doesn't say whether it is a quest, a drop or a vendor item, or what level it needs. The hover tooltip may cover this, but I can't see it on a still capture.
12. **"No companion" in the footer** reads like an error. "Companion not running" is clearer.
13. **Level 5 "Pick your spec" with a gold dot on Builds.** A hunter has no talents before level 10, so the prompt is early. Reword it to "Pick a build to follow", or hold it until a point exists.
14. **Mixed pronouns and tone.** "Where you stand" and "Right now" are fine. "Sim again" and "Builds" are fine. "Set up" step 3 uses a scroll icon for the app, which is a stretch but acceptable.

**Data I don't believe (for the accuracy reviewer)**
- **Gloves of Holy Might on a Beast Mastery hunter at 60, +11.2 DPS.** The name reads like a plate or caster glove. Confirm armor type and stats. The +11.2 is also the biggest gain on screen for a fresh 60.
- **Forest Stalker's Bracers as epic.** Check the rarity against the client.
- **Brawler's Leather Hood for a level-25 hunter.** Leather is legal before 40, but confirm it beats the mail-free options at that band. +1.5 DPS from a hood is plausible.
- **Improved Aspect of the Monkey as the next point after 15 on plan.** Check it against the actual leveling order.
- **Onyxia's Lair "Phase 2 ranged, west side" on the fresh-60 capture.** Confirm Onyxia is open in the raid schedule shown. Raids open 12/9 and the launch is 11/4.

**What works**
- The empty states are honest and short.
- "Level 10, your first talent point" at level 5 is exactly what a player wants to know.
- The gold dot on Builds when a point is waiting is a good nudge.
- The DPS gains use "+X DPS" with a unit, and the item colours match quality.

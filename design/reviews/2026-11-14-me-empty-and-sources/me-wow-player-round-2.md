# wow-player, round 2 (Me with no site activity)

Verdict: FIX

I re-read the four re-rendered captures. The earlier must-fixes are closed and the frame is much better. Four small defects remain, all visible in the captures.

**Must-fix**
1. **Truncated text in the enchant rows.** "No enchant: Enchant Chest - Minor Stats · Che…" and "…Minor Agility · …" cut off the slot name, which is the part that tells you where to put it. The Cloak row at 60 ("Greater Defense …") has the same problem. The repeated "No enchant: Enchant" is wasted width. Write "Chest: Enchant Chest - Minor Stats" in red, or drop the "No enchant:" prefix and let the red colour carry it.
2. **"Spend 1 point: Improved Aspect …" is cut off** under "Right now" at level 26. Shorten it or let it wrap. It also repeats the "Next point: Improved Aspect of the Monkey, Now" row to its left, which is the obvious thing explained twice. Either say "Spend 1 talent point" with the Builds link, or drop the name.
3. **"Captain Green…" is truncated on Impaling Harpoon.** "Main hand · The Deadmines · Captain Green…" is the one source line players will check, and it lost the name. Drop "Main hand" (the item says it) so the boss fits, or use "Captain Greenskin" in full.
4. **A level-5 hunter sees "Pick your spec" under "Right now" with the gold dot on Builds.** Hunters can't spend a point before 10, and the left side already says "Level 10: your first talent point". Reword it to "Pick a build to follow", or hold the prompt until a point exists. If this is a deliberate pre-pick, say so.

**Shoulds**
5. **"8 of 14 best in slot" is kept by decision, but the website's word is the wrong word in a WoW frame.** Players read BiS as the best item you could ever have. It is cheap to say "8 of 14 on your list". At 60 it is also odd that fewer than 17 slots are counted. A single "ammo, ranged, rings" gap would be fine, but say what is counted.
6. **"Rested 40%" stays.** Fine, but it is the client's own XP bar tooltip. If room is needed, it goes first.
7. **Professions at level 25, "Skinning 112, Leatherworking 105".** Skill levels are plausible. Skinning at 112 with a level-25 character is a touch high but legal. Leatherworking 105 for a Brawler's Leather Hood craft makes sense and matches the hood's source.
8. **"No companion" in the footer** still reads like an error to some. "Companion not running" is clearer. This is your call.

**Data to check (accuracy reviewer)**
- **Gloves of Holy Might, "Hands · World drop", +11.2 DPS** on a Beast Mastery hunter. Check that it is a real world drop and a hunter-usable glove, and that the +11.2 is the list's own gain. It is the largest gain on any capture.
- **Forest Stalker's Bracers, "Silverwing Sentinels", epic.** A Silverwing reputation reward is plausible. Check the rarity and the standing it requires.
- **Windreaver Greaves, "Scholomance · Kirtonos the Herald".** Check the boss. Kirtonos drops the Windreaver Greaves in Classic, so this looks right.
- **Impaling Harpoon, "The Deadmines · Captain Green…".** Checks out for a Deadmines spear if it is Captain Greenskin's drop. Verify.
- **"Enchant Bracer - Minor Agility" and "Enchant Chest - Minor Stats" at level 25.** Both are real vanilla enchants. "Enchant Cloak - Greater Defense" at 60 for a hunter is odd. Agility is the hunter stat, so check that the enchant plan isn't a defence pick for a ranged DPS.

**Good**
- "Not simmed yet" with "Run a sim" is honest.
- The one site line reads right.
- The sources under each item fix my earlier complaint.
- The red "?" icon is gone.
- The talent counts now agree (16 of 16, then 16 of 17).

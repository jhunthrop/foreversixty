# Round 9 ai-slop: guides (1440 and 390)

Tell 23 = definitional blurbs. f-content handed items count. Blank spell/class squares inside lists are fixture artefacts unless noted. Spec 1440 read in 5 slices, 390 in 8 slices side by side.

## guides: PASS
Nine class plates, class-coloured names, spec rows with icons, one "Updated 7 Nov". Same at 390. Zero tells.

## guides-you: PASS
Same page plus the Felsnap block and a highlighted Beast Mastery row. The "Your guide:" prefix is gone. At 390 the pill no longer wraps. Zero tells.

## guides-class (Hunter): FAIL (one tell, one-line fix)
- Tell 22, Sources footer (1440 and 390): the pill says "Blizzard" and the link text starts "Blizzard, Class Deep Dives..."; same for "Community" + "Warcraft Tavern, ..." and "Datamined" + "Forever Sixty, race and class pairs". The pill and the prefix say the same thing. Drop the "Blizzard, " / "Forever Sixty, " prefix from the title (the pill carries it). Same on the spec page, plus "This site" + "Forever Sixty, stat weights".
- The dead tan under the changes plate is gone, the "Mail from 40." bullet is gone. Races rail ok. Nothing else.

## guides-spec (Survival Hunter): FAIL
Handed items fixed: Recommended pills, "1 point of agility = 0.63" footer, "Numbers from client build", "DPS · Updated", "Your guide:", Humans/Skyborne rail restatements, the Strengths/Weaknesses caps eyebrows (now plain headings), broken-image glyphs (now bordered placeholders).

Still present:
1. Tell 23 / 19, tooltip restatements in Key talents: "Predator's Edge 5/5 / Your melee special crits hit for 230% of a normal hit instead of double, and your off hand does 50% more damage" (Overview and Stat priority prose already say it); the Lone Wolf blurb is a 9-line paragraph repeating the Strengths bullet and the Talents prose. Keep Expose Prey only.
2. Tell 23, handed lines still in prose: "You fire no shots in this spec. Your bow or gun is there for its stats." (Rotation) and "Plain attack power gives only itself." (Stat priority). "Levels 20 to 60" still opens "Survival levels in melee." (the same opener, moved).
3. Tell 22: weights header still reads "DPS per point or rating" ("or rating" qualifies nothing, then a footnote explains it). "51 points spent, all placed" under "Level 60: Lone Wolf, rank 1" explains the obvious. "Pre-raid 60" on the doll plate next to a "Level 60" list label. "Weight (Agility = 1.00)" in the band and again in the table.
4. Repetition (worst structural problem): the Hit cap block appears word for word in the band and again in Stat priority; weights appear in the band, in the Stat priority rail and, at 390, a third time as a bar list; Rotation rows appear in the band, then as Opener/Priority; Build/Load/Sim in the band, then the talent grid again in Talents and Leveling. At 390 the band is ~1,300 px before Overview. Best in slot shows the 17 items twice (paperdoll and list) and repeats the BiS page.
5. Weights table split: at 1440 the Stat priority section has the full table in the right rail and a second orphan table in the left column holding only "Attack speed / Not on any item" with its own header row. Fold the row into the rail table or delete it.
6. Races split: at 1440 Troll, Tauren and Windshaper Skyborne sit under the prose in the left column while Night Elf, Human, Dwarf, Skyborne and Orc sit in the right rail; one list cut in two for no reason. Race squares are empty, and Enchants, Consumables and Extras rows have an empty icon gutter (handed, still present) at both widths.
7. Dead plate: Best in slot leaves ~250 px empty under the paperdoll at 1440; Levels 20 to 60 has ~600 px empty right rail beside it.
8. Tell 6 (weaker): small grey caps labels "Build", "Rotation", "Stat priority", "Spec", "Jump to", "Alliance", "Horde", "Food, oil and potions". Fine on the slot list, cut on the band and rail.
9. Tell 21/19: Rotation and Leveling are walls of 6 to 9 line paragraphs; "Mail at 40 is the biggest armor upgrade..." sits mid-list between rank notes and repeats "Wear leather until 40. From 40 you can wear mail" from two paragraphs earlier and from Best in slot.
10. 390 sticky bar reads "Survival Hunter · Jump to" + "Change"; it is a jump menu, so the control should not say "Change" (ui/Toc.svelte).

Authentic: spell links with rank, "20 s before the pull", item source lines ("1% from Warchief Rend Blackhand"), the paperdoll, real Forever rotation rules, talents-by-level slider.

Fix order: 4, 5, 6, 1, 2, 3, then 7, 10.

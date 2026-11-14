# ai-slop-reviewer, round 1 (Me with no site activity)

Verdict: FAIL

Checked all five captures and Locale.lua lines 1243-1260. There are no visual tells. The slab is flat and the gear rows use real item icons, quality colours and DPS gains. This reads as a player-built addon, and the failures below are all in copy and state design.

**Must-fix**
1. **Tell 13: the empty state pretends to be content (l60).**
   - "Your sim" is a heading with only a "Sim again" button under it, and there are no numbers. It offers to re-run something that has never run.
   - It is also inconsistent with `meSiteFirstRaid`. The l60 capture shows "Companion synced 14 min ago" in the footer, yet the sim is empty.
   - Fix: drop the "Your sim" block until a sim result exists. Show a plain "Sim" button only if the action is useful with no result.
2. **Tell 22: qualifiers the context already gives (l25 and l26).**
   - "8 of 14 best in slot": "in slot" is redundant next to a gear list. Use "8/14 best".
   - "15 of 15 on plan" shows twice, once as "15 of 15 on plan" in the Build row and once as "15 of 15 points on plan" in Right now. Keep one, since it is the same number stated twice.
   - "Level 10, your first talent point" and "Level 30, a new gear list" explain the obvious to a veteran. Use "10, first talent point" and "30, gear list".
   - "Max level" beside Level 60 is redundant.
   - The DPS-gain rows are fine because they take a unit and nothing else.
3. **Tells 18 and 19: filler copy.**
   - `meSiteApp` ("Your score, raid logs and sim show here once the companion app brings them in.") is a pitch sentence with vague nouns ("score" is not defined).
   - It also sits as a grey paragraph over the same link on both l25 and l5, so it reads as a templated explainer.
   - `meSiteFirstRaid` is the same pattern.
   - Fix: replace the paragraph with one specific line, for example "Raid logs: foreversixty.gg/app", or cut it and keep only the URL.
   - Cut "Rested 40%" unless it is actionable.
4. **Tell 13: dead space (l25 and l5).**
   - The right column is mostly empty at l5 ("Right now" has one line) and below the gear list at l25.
   - The left panel has a 60 px gap above the filler paragraph.
   - Fix: fill with real data, such as the next trainable spells, quest zone level ranges, or the nearest upgrade source. If there is none, collapse the panel height.
5. **Tell 17: the l5-setup "Set up" list is a numbered pseudo-sequence (tell 10).**
   - Pick spec, hover an item and get the app are not an ordered sequence.
   - Step 2 uses a red question-mark icon as a placeholder for "an item".
   - "Hover an item you wear" explains the obvious.
   - Fix: drop the numbers, use real icons, and cut step 2.

**Shoulds**
6. The footer "No companion" at l25 and l5 against "Companion synced 14 min ago" at l60 and "/fs tour" at l5-setup is inconsistent. Pick one footer-right role.
7. "Where you stand" and "Right now" are both friendly-pitch headings. Plain "Character" and "Now" are what an addon would use. This is the only sentence-case copy tone on the screen.
8. The "Builds" link in yellow on the right of "Pick your spec" duplicates the Builds tab dot. Keep one.
9. The same Heading + rows block repeats in three sections ("Where you stand", "Right now", "Gear"). Vary it by dropping the "Your sim" heading, which already fails per item 1.
10. Em-dash check: none found in the captured copy. Keep it that way.

**Authentic, keep**
- Real item names such as "Brawler's Leather Hood" and "Impaling Harpoon".
- Quality-coloured links and DPS deltas.
- Real talent icon and name for "Improved Aspect of the Monkey".
- Real zone and timing ("Onyxia's Lair · 20:00, in 3 h 12 min").
- Uneven hierarchy: the XP bar and build lead.
- Footer sync-age detail.

VERDICT: SHIP

First impression: the data inside the windows reads like a raider wrote it (real Onyxia phase 2 job text, Molten Core bosses, sums and medians that actually add up), but the plate wrapper (fogged sky-to-sand backdrop, identical caption band, thin serif titles) and a few placeholder-looking regions read as generated.

Checked and correct, so nobody wastes time on them: Forever Score parts sum on plates 9 and 10 (410+352+268+150+60 = 1240; 410+380+228+120+46 = 1184; 1500-1184 = 316); median parse 74; downtime 92 s on plate 5 equals the per-boss sum; sign-up counts 11+1+2+1+16 = 31 and 12 coming; 9th of the shown scores; weekdays match 2026 dates; the +21.2 total on plate 6 equals the slot gains. Plate 3 is judged only as a neighbour.

## Findings

1. BLOCKER (tell 2, 12). All plates, full backdrop. The blue sky fading into mottled tan ground is a decorative gradient plus painted noise, identical on all 11, and it carries no information about the addon. Fix: one flat colour from the site ink token behind every plate, no mask fade, no ground texture. Let the window and its own header band be the only art.

2. BLOCKER (tell 8). Plates 1, 2, 9, 11, the dark rounded "stage" panel around the tooltips. It is a card holding cards (stage, then tooltip frame, then rows), and plates 4 to 8 and 10 do not use it, so the set is inconsistent. Fix: drop the stage on those four plates and place the frames straight on the flat backdrop, as the window plates already do.

3. BLOCKER (tell 22, owner rule). Plate 2, tooltip line "Beast Mastery, your sim". "your sim" says nothing a player needs; the spec is already implied by the number. Same family: plate 1 tunic tooltip "Upgrade +2.0 DPS" under a green triangle that already says upgrade, then "Best for 20 to 29" again, which is already on the reward button beside it. Fix: delete the "Beast Mastery, your sim" line; render the gain as "+2.0 DPS" next to the triangle, and show "Best for 20 to 29" once.

4. SHOULD (tell 13). Plate 8, "Master loot" panel: "Opens here when you loot as master looter." is an instruction pretending to be content, and the Roster panel under the list is about 200 px of empty box. Fix: replace the master loot text with the 7-item list the line already counts, or cut the panel; crop the roster panel to its rows so the window ends where the data ends.

5. SHOULD (tell 13). Plate 4, "Up next: Every point is placed." and the empty "Paste a build code (Ctrl+V)" field. Two of the four right-hand panels show nothing. Fix: capture a state with points to spend (the talent plate already has one) so "Up next" names a talent, and leave the paste field out of the plate or fill it.

6. SHOULD (tell 20, 13). Plate 4, export string "FS1:70938:hunter:60:Thornmaw:Thornwake Company:Forever:1,2:a1b5c2d7e3f9g1h2i5j0k3l2...". The tail walks the alphabet a, b, c, d, e, f, g, h, i, j, k, l. Any dev spots it as filler. It also breaks mid-word ("Thornwake Co" / "mpany"). Fix: use a real encoder output from the fixture character, and let the box truncate at a separator or with a single ellipsis, not mid-word.

7. SHOULD (tell 21). Plate 4, helper copy: "Talents, gear and professions in one code for the planner's Import from addon box.", "Click the box, then press Ctrl+C.", "14 items · 2 professions · your guild and realm are included". Veterans know how to copy a box. Also tell 7 on each spec row: grey explainer after the spec name ("Your pet fights, Bestial Wrath") followed by another grey explainer line per build. Fix: keep the build lines (talent split plus the one differing talent), drop the spec blurb and the how-to-copy sentences.

8. SHOULD (tell 22). Captions that restate the plate: plate 11 quotes "Onyxia's Lair 20:00 · In" which is the chip right above it; plate 9 quotes "Forever Score 1240" which is in the tooltip. Plate 1 caption wraps to an orphan "use." and plate 8 to "everyone not signed up." Fix: drop the quoted strings from captions, shorten plate 1 to one line ("Green arrow on the reward or roll that beats what you wear. Not for you on a roll your spec can't use" fits if the quote marks go).

9. SHOULD (tell 5, 8, 14). All plates, the footer band. Same crest, "Forever Sixty" wordmark, thin large serif title and grey Open Sans line, stamped 11 times, in addition to "Forever Sixty" in the tab strip and in every tooltip, and the upload steps already set the same title and description as store fields, so the store shows each caption twice. Fix: remove the caption band from the PNGs (the store fields carry it) and let the window fill the frame; if a band stays, drop the wordmark and set the title at body size, not display size.

10. SHOULD (tell 20, 18 adjacent). Roster and realm names. Velisandre, Aldrenna, Kestrin, Orlaith, Brannoc, Ysolde, Torvald, Kethra, Mirelle, Nessaly and guilds Thornwake Company, Ashen Covenant, realm Thornmaw all share one elegant-fantasy register. Real raid rosters carry Dravik and Pell next to Garrow, Quenna and joke names. Fix: swap about a third for blunt or odd names (short, lowercase-ish, one with a number-free pun), and name the guild the way players do ("Late Night Wipes").

11. SHOULD (fidelity). Plates 1, 2, 9, 11. The spec says tooltips stay native with the client font; these are drawn in Marcellus as a stand-in, so the numerals read oddly ("4O / 4O", "Level 2O", "1O" like "IO") and the whole plate is visibly a mock beside real item tooltips. Fix: capture the four tooltip plates from the real client (item, quest, unit, minimap), or at least use the client tooltip face in the render.

12. SHOULD. Plate 11, no minimap. Title says "By the minimap" but the chip sits on the edge of an empty dark box, with a blank left half and a lot of unused plate. Fix: show a real minimap corner with the chip attached, or retitle "Chip" and crop to the chip plus tooltip.

13. SHOULD. Plates 5 and 6 repeat the identical Stat weights table (same five rows, same numbers), and plates 6 and 10 repeat the same three upgrades. The plate 5 caption is about rotation, yet a quarter of it is a table the next plate shows again. Fix: on plate 5 show the second half of the priority (the cooldown row or the level 40 priority) instead of the stat table, or fill the dead right-hand space under "See parses and deaths".

14. NICE (tell 2 mild). Plates 4, 5, 6, 10 header band is a teal gradient wash, plates 7 and 8 use dark lava art. If the band is the spec's art veil it is fine, but the two looks side by side read as two skins. Fix: use the same art veil on all window headers, or the same flat colour.

15. NICE. Plate 9, two tooltips of unequal height with empty space under the left one and large margins; plate 3 keeps the gold Blizzard frame next to six teal-and-brown plates (known, unchanged). Fix: align both tooltips to the top and crop the stage tight; re-render 03 in the new skin when the talent window is next touched.

16. NICE (tell 22). Plate 8 legend "Gap in DPS", plate 6 "Better than the list", plate 10 "Died" column of 0 on seven of eight rows. Fix: "Gap" is enough as a column name; "Better than the list" means nothing without the list, use "Off list"; hide a column that is zero on every row but one.

17. NICE. Plate 10, Attendance part 120 of 200 (60 percent) next to roster attendance 88 percent on plates 7 and 8. A raider might reconcile them; if the part weights recent nights, say so nowhere, but make sure the fixture cannot read as a contradiction.

## What reads as authentic (keep)

- The numbers hold together across plates: same character, same raid, same Gloves of Holy Might +2.7 DPS on plates 2, 6, 10, 11, boss-by-boss downtime matching the totals.
- Real content: Drakefire Amulet attunement, Greater Fire Protection Potion missing, "Phase 2 ranged, west side. Fire on the whelps from the west wall when she lifts off. With Mirelle." is exactly how an officer writes an assignment. Real boss and item names, Westfall quest set, "Don Julio's Band".
- Officer plate density: Gap, Ench., iLvl columns, readiness lines with names, "Remind 16 in chat" matching 16 not signed up, footer hint "Click a name to whisper · Shift-click to invite".
- Uneven emphasis: the 1184 Rare score leads plate 10; green arrow and gold ring leads plates 1 to 3.
- Class-coloured names and class crests on the guild lists, tab dots only where something waits, plain rectangular buttons with no glow, no emoji, no em dashes in captions.

## Verdict

FIX. Findings 1 and 2 are visual tells (2, 12, 8) that block a pass on their own, and finding 3 breaks the owner's explicit rule on qualifiers. Clear 1 to 3, then 4 to 9 and 11 to 12 before upload; 10 and 13 to 17 can ride along.

## Round 2

VERDICT: FIX

First impression: the plates now read as a raider's screenshots (real Onyxia job text, Molten Core parse tables, sums that add up) and the opening paragraph sounds like an officer talking, but the pitch sells raid sign-ups and a roster as plain addon features while the page itself says they need the companion app, and one bullet claims enchant slots the addon does not check.

Re-checked and correct: 9th of 15 raiders (1184 sits under Garrow 1188), sign-ups 11+1+2+1+16 = 31, weekdays for Dec 2026 (Tue 15, Thu 17, Sun 20, Tue 22), downtime 92 s, mean followed 85%, Kestrin parts sum to 1240, +21.2 total equals the slot gains. Round 1 findings 2 (stage panel), 11 (digits) and the builds export code are fixed. Name and summary: acceptable. They are plain nouns with no "unlock" verbs, and the summary says what it does. "Guild Raids, Sim and BiS" is a keyword list a store needs, not a line a raid leader would say, which is fine for a title field. No em dashes anywhere in the file.

### Findings

1. BLOCKER (text, accuracy). Opening paragraph and bullets 2 to 4: "Forever Sixty runs your guild's raid sign-ups" reads as an addon feature, but the Guild and Officer raid parts need the companion app (the foot section says so, and heads it "optional"). An officer installs it for tonight's raid and gets an empty Guild tab. Fix: say it in the opening in one clause ("Sign-ups and jobs come through the companion app on your PC") and retitle the foot section "The companion app", naming the raid parts as needing it.

2. BLOCKER (text, accuracy). Bullet 3: "whether chest, back, boots and weapon are enchanted". The addon checks the slots the spec's BiS plan enchants (addon/ForeverSixty/EnchantPlan.lua), and plate 5 shows "No enchant: Chest, Feet", not back or weapon. Fix: "whether the slots your spec enchants carry an enchant".

3. SHOULD (text). The sim claim appears three times: summary sentence 2, the opening's second sentence and bullet 1. Fix: keep it in the opening, and make bullet 1 "Every upgrade shows its DPS gain", placed after the raid bullets, since the pitch is raid first and this bullet leads the list.

4. SHOULD (text, voice). Guild and Officer sections narrate the window panel by panel ("Under it", "Below it"), and Officer is one paragraph of about 130 words. That is the memo voice returning. Fix: bullets of nouns a leader scans (sign-ups by role, what you are missing for the fight, top five scores, readiness counts, Remind, Send to guild, master loot), no layout directions.

5. SHOULD, BLOCKER if generated (plate backdrop, tell 15/12). Kept as the live brand band, accepted as the site's look: it is painted art, not a CSS gradient or glow. But it is the largest thing in frame (windows fill under half of most plates, under a third of 3, 4, 10, 11) and the soft cloud and mottled sand still read airbrushed. Confirm the art is hand-made or licensed (if image-model output, this becomes a tell 15 blocker), and scale the windows up so the backdrop is margin, not subject.

6. SHOULD (plates, tell 5/14). The caption band on every plate still stamps the crest, the "Forever Sixty" wordmark and a thin display serif title 11 times, beside the store's own title and description fields (round 1 finding 9, not addressed). Fix: drop the wordmark, set the title at body weight, or remove the band.

7. SHOULD (captions). Plate 5 wraps to the orphan "from."; plate 10 to "parts on Shift."; plate 9 runs two long lines. Fix: 5 "Best in slot at your level against what you wear. The top upgrades say where they drop." 10 "Forever Score under the crest in the tier's colour, greyed when old. Shift shows the five parts."

8. SHOULD (plates, tell 13). Plate 2: the Roster panel ends in about 200 px of empty box. Plate 8: "Up next: Every point is placed." and an empty "Paste a build code (Ctrl+V)" field. Fix: crop the roster panel to its rows; for 8 capture a state with a point to spend, or leave the paste field out.

9. SHOULD (plate 5). Sources shows "Raid 0" and Upgrades names no raid drop, in a gallery pitched at raids. Fix: a state with at least one Molten Core piece in Upgrades.

10. SHOULD (shipped strings, kept by decision). "Upgrade +2.0 DPS" and "Beast Mastery, your sim" (plates 3 and 4), "Simmed Mon 14 Dec" (7), "Better than the list" (5), "Gap in DPS" (2) are tell 22 and the owner's rule. Accepted on the plates because they are what the addon draws, but schedule the addon wording change and a re-shoot. Not a gate blocker.

11. NICE (plates 10, 11). Tooltip stand-in pairs a serif face for letters with Open Sans digits on one line (1563, 410). Plate 10's tooltips differ in height on a mostly empty plate. Plate 11 shows a chip with no minimap, fine now it is retitled.

12. NICE (text). Seven categories dilute a raid-first pitch; keep Guild, Boss Encounters, Tooltip, Bags & Inventory.

13. NICE (text). "Raid night for Forever guilds." as a heading is a tagline; start at the sentence under it. "Forever's rules" appears in the summary, opening and tooltip section. "Our sim measures what one point of each stat is worth" explains stat weights to veterans. "What's new in 2.0.55" says the window "looks like the website", which is brand talk; say what a player can now do.

14. Round 1 findings 4 to 7, 10, 12 to 17: none is a blocker for the store gallery.

### What reads as authentic (keep)
- Opening: "When a raider asks whether the new chest is worth it" is how a leader talks.
- "In, Late, Bench or Out" and the Phase 2 job text with "With Mirelle".
- Officer readiness list with names and counts, Remind 16 matching 16 not signed up.
- Phone app marked "(in beta)", `/fs diag` as the support answer, Commands and Settings tables that match the addon.

### Verdict
FIX. Findings 1 and 2 are accuracy blockers in the new raid-first text; 5 becomes a blocker only if the backdrop art is generated. No plate carries tells 1, 2, 8 or 12 as it stands. Clear 1 and 2, confirm 5, then 3, 4 and 6 to 9 before upload; 10 to 14 can ride along.

## Round 3

VERDICT: SHIP

First impression: reads as a raid leader talking to the guild. The opening, the bullets and the Guild and Officer lists now say what a raider or officer sees, and the plates read as screenshots of a real addon.

Round 2 blockers, checked against the file and the plates:
- Finding 1 (companion app): fixed. The opening says the guild parts in game come through the companion app and that upgrades, marks, Builds and Gear work without it; the Guild section says how a guild gets on it; the foot section is titled "The companion app (needed for the guild tabs)". The tab no longer promises more than it delivers.
- Finding 2 (enchant slots): fixed. The bullet now credits the site's readiness board with chest, back, boots and weapon and describes the in-game Officer line separately. Plate 5 shows "No enchant: Chest, Feet", which matches the addon's own plan check.
- Finding 3 (sim claim): fixed, now in the summary and bullet 1 only. Summary and bullet wording kept as the owner's choice; they are plain and specific, no tell.
- Finding 4 (Guild and Officer voice): fixed. Short bullets with nouns, Officer leads with Remind. "Phase 2 ranged, west side" and "Greater Fire Protection Potion" in the raider paragraph are what a leader would write.
- Finding 7 (captions 5 and 10): fixed. Both are one line, no orphan word. No em dashes in the file.

### Remaining findings (none blocks upload)

1. SHOULD (plates, backdrop, tell 15/12). Kept as the site's painted band over ground. Accepted as the live brand on this basis, so the one open question is provenance: it must be hand-made or licensed art, not image-model output. If it is generated, this becomes a BLOCKER. One line from the owner closes it.
2. SHOULD (plates, tell 5/14). The caption band (crest, wordmark, thin display-serif title) is still stamped on all 11 plates and repeats the store's own title and description fields. Not a slogan hero, so not an automatic fail, but the first thing to cut when the generator is next touched: drop the wordmark or the band.
3. SHOULD (plate 2, tell 13). The roster panel ends in a large empty area; plate 8 shows "Every point is placed." and an empty paste field. Both are real states of the shipped fixture, not fake content, so they stay SHOULD. Crop plate 2's panel to its rows when the generator allows.
4. SHOULD (plate 5). "Raid 0" in Sources and no raid drop in Upgrades, on a raid-first gallery.
5. SHOULD (shipped strings). "Upgrade +2.0 DPS", "Beast Mastery, your sim", "Simmed Mon 14 Dec", "Better than the list", "Gap in DPS": tell 22 in the addon itself. Change in the addon, then re-shoot.
6. NICE. Seven categories (keep Guild, Boss Encounters, Tooltip, Bags & Inventory); the "Raid night for Forever guilds." heading is a tagline; the title field "Guild Raids, Sim and BiS" is a keyword list, fine for a store.

### Verdict
SHIP. No plate carries tells 1, 2, 8, 12 or 13 in a form that misleads, the accuracy blockers are gone, and the text voice is a person's. Ship on the condition that the owner confirms the backdrop art is not generated (finding 1).

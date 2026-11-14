# UX designer verdict, addon growth lane (2026-11-14)

Reviewed: addons-list-tooltip.png (after pane), addon/ForeverSixty/ForeverSixty.toc, design/listing/2026-11-14-growth-listing.md against design/listing/curseforge-description.md. Reference: the client AddOns list tooltip and CurseForge search rows. No em dashes in either file (grepped).

## Verdict: FIX

### Must fix
1. Featured image and upload order contradict the approved file. The proposal leads with the bag arrow (plate 2) "because the first image is the one search shows", but the approved steps star plate 1 as featured and the Wago reverse upload is keyed to plate numbers. Either renumber the gallery (plate 02 becomes 01, captions move with it, step 3 star and step 4 reverse order restated) or keep plate 1 featured. One rule, both stores, written in the file.
2. The one line drops the unit. Approved: "with the gain in DPS". Proposed: "with the gain." A veteran reads "the gain" as nothing. Use "with the gain in DPS". Count: the proposed line is 191 characters, so "in DPS" (+7) hits 198 of CurseForge's 200. Cut "every" before "bag item" (-6) to keep a margin.
3. Categories column "Live" is wrong. It says Tooltip, Guild, Map & Minimap; the approved file lists Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat. Correct the table, and say what is removed: the proposal also drops Boss Encounters and Combat, which are where the rotation grading and raid sign-up are found. Keep Combat at least, or state why not.

### Should fix
1. TOC Notes: "your next talent marked" is a fragment tacked on a comma. Suggest: "Gear upgrades, BiS, talents and rotation for your class in Forever. A green arrow on gear that beats yours and your next talent marked. /fs opens it." Wraps to four lines in the capture, the same as the proposal; the before text was three. Do not add words.
2. `## Category: Class` cannot be judged from the capture: the mock shows no grouped list. Confirm in the Forever client (1.60.1) that the field is read and the addon groups under Class, and attach that screenshot. Remove the line if the client ignores it.
3. Description first line is install logistics above the value statement ("The arrow is right for Forever."). The approved Install section already says the apps add Forever Sixty Data. Do not add the line on top; extend the Install sentence with "Installing by hand? Get both folders" instead.
4. Name: the proposed "Forever Sixty: Gear Upgrades, BiS & Talents" is the right one of the three (the player's search words, 43 characters). Owner call stands. Also change the Website field note: unchanged, fine.
5. Summary order: lead with what the player searches (BiS, talents), then the arrow. The name already carries them, so the arrow first is acceptable once item 1 is settled.

### Checklist
- [x] No em dashes, game words only
- [x] TOC Notes match the listing in-game line word for word
- [x] Name searches a word a player types
- [ ] Featured plate rule consistent (must 1)
- [ ] Unit in the one line (must 2)
- [ ] Category table accurate (must 3)
- [ ] Category field proven in client (should 2)

## Round 2: SHIP

Must fixes 1 to 3 are closed: gallery order and featured plate stand as approved, "in DPS" is back (192 of 200 characters), the categories table now shows Live, Approved and Proposed, and Boss Encounters and Combat are kept. No em dashes in the TOC or the listing file.

Shoulds, none blocking:
1. The re-rendered tooltip leaves "it." alone on the fourth line. Drop "raid" ("rotation and sign-ups") or move "/fs opens it." so the last line holds two or more words.
2. Confirm CurseForge accepts seven categories (primary plus six) before the dashboard session. The approved list was also seven and was never published.
3. `## Category: Class` is backed only by Questie's Forever TOC. Attach one in-game AddOns list screenshot after the first login on 1.60.1, or remove the line.

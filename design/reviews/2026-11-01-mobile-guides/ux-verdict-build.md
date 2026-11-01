# F9 Guides tab, build review (UX)

Verdict: HOLD

Reviewed iOS and Android captures against the boards and spec §5.1 F9, §4, §9. Layout, hierarchy, crests, class colours, Your guide card, class page, offline states, new-data strip and the changelog match the boards.

## MUST
1. Mana Spring Totem prints a duration. ios/f9-guide-rotation-ios.png, and/f9-guide-rotation-and.png show "Restores 10 mana every 2 seconds for its five-minute life". Spec F9.3: no duration is printed for it (client reads 300 s, Era 60 s, unsettled); the line is "Down first. It returns mana for the rest of the fight." Remove the duration until the accuracy reviewer settles it.
2. Missing icons in the rotation. ios/f9-guide-rotation-ios.png and and/f9-guide-rotation-and.png: Fire Nova draws the red question-mark placeholder; on iOS Astral Recall and "Trinkets, racials and potions" draw empty dark tiles. Tenet 3 and the no-missing-icon rule: every row has a real spell icon; for the trinkets row use a real generic icon, and map Fire Nova and Astral Recall to their spell icons.
3. iOS largest text, numbered chips. ios/f9-guide-large-ios.png (also ios/f9-guide-talents-large-ios.png if chips appear): the 22 pt number disc does not scale, so the numbers 1 to 7 render at about 8 pt next to 30 pt names (Android scales them correctly, and/f9-guide-large-and.png), and the "Spell penetration" chip is cut at the right edge. Spec §4: chips wrap, never cut. Scale the disc and number with the text and let the label wrap inside the pill.
4. iOS largest text, tier chip. ios/f9-tiers-large-ios.png: the tier letter chip (S/A/B/C/D) shrinks to about 10 pt and is unreadable; the letter is the row's answer. Scale it with the text as Android does (and/f9-tiers-large-and.png). The 28 pt chip must grow, not shrink, at AX5.

## SHOULD
1. ios/f9-changes-ios.png: the Name reservation row prints "through Nov 3" and, in gold, "Name reservation open through Nov 3". Same fact twice; keep the gold running sentence only.
2. ios/f9-tiers-large-ios.png: the three segments stack to two lines on Android (and/f9-tiers-large-and.png "What changed" wraps) and the iOS segment control crushes to a 3-row stack; at AX5 allow the segments to scroll or stack as one list with 44 pt rows.
3. Tier chip shape: iOS and Android captures draw a rounded square, spec §4 says pill (28 pt, `--radius-pill`); confirm against the board.
4. ios/f9-guide-top-ios.png: with the share button absent in simulator the bar is empty; confirm on device that the bar title carries the spec in the class colour (Android shows it, and/f9-guide-top-and.png).
5. Offline class page: add the board line "This guide is not on your phone yet. Connect once to save it." (announced in the brief); verify in ios/f9-class-offline-ios.png and and/f9-class-offline-and.png on the next capture.

## Checklist
- Answer visible without scroll (Your guide card; stat priority chips): yes
- Section strip, sticky, gold current chip: yes
- Numbered chips at default: yes; at AX5 iOS: no (MUST 3)
- Rotation rows with icon, rank, condition: yes; duration rule: no (MUST 1); icons: no (MUST 2)
- Tier rows, Yours edge, tie mark, Top, gaps: yes
- What changed key dates and changelog rows: yes
- New-data strip over a guide: yes
- Class page Saved / Not saved: yes

## Round 2

Verdict: SHIP

Checked ios/f9-guide-rotation-ios.png, ios/f9-guide-large-ios.png, ios/f9-tiers-large-ios.png against round 1.

- MUST 2 (icons): fixed. Fire Nova draws a neutral bordered tile, no question mark; the Trinkets row has no tile; Astral Recall is gone. Accepted. The empty tile is still a placeholder; if the client ever ships a Fire Nova rank icon, use it.
- MUST 3 (AX5 chips): fixed. The number disc scales to match the label, labels wrap inside the pill, nothing is cut.
- MUST 4 (AX5 tier chip): fixed. The letter chip is a pill and scales with the text, legible.
- MUST 1 (Mana Spring Totem duration): withdrawn as a UX blocker. Spec F9.2 allows a duration when the client data and the guide agree, and the developer states both now say five minutes. Condition: the accuracy reviewer must confirm 300 s is what Forever ships (Era players know 60 s). If they do not, the line reverts to "Down first. It returns mana for the rest of the fight."
- SHOULD items 1, 2, 3, 5 from round 1 are reported fixed; the running key date and segment stack at AX5 look correct in the captures I opened. Not re-opened: ios/f9-changes-ios.png and the Android set.

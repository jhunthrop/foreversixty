# Paperdoll spec and mocks: second-designer verdict (2026-10-31)

Reviewed: design/specs/2026-10-31-paperdoll.md and the boards raid-360/390/1024/1280/2000, beta-390/1280, states.png. Loading and empty boards not opened.

## Verdict: HOLD

The direction is right and the desktop boards are good. Two boards and one phone rule do not hold up, and they cover exactly the cases the owner will test: beta-unknown items, a two-hander, the phone.

## Blockers

1. Phone character page (beta-390) loses the answer. With a gain line, the name clamps to 1 line, so "Knight-Lieutenant's Plate Helm" and "Knight-Lieutenant's Plate Pauldrons" both render as "Knight-...". The two slots cannot be told apart. The picks truncate to "Lionhe...", "Darksp...", "Forest S...". The gain is the next step (tenet 14) and the pick is what the player acts on, yet the spec makes the pick untappable below 640 and sends the player to the BiS link. Fix before build. Either (a) on phone the gain line becomes its own full-width row under the pair of columns, or (b) the phone cell stacks name (2 lines) over "+18.4 pts" only, and the pick shows in the tooltip, with the tooltip body naming the pick and its gain. Pick one, draw it, and cap truncation so no two slots can collapse to the same visible string.

2. Phone clamping has no ellipsis and cuts mid-phrase (raid-360). "Darkmoon Card:" and "Frozen Heart of the" end with nothing to show text was dropped; "Signet Ring of the Bronze" drops "Dragonflight" silently. Spec §3.3 says 2-line clamp, but the mock shows clamp without `text-overflow` marks and at 360 drops to a state where 2 lines are not enough. Specify `-webkit-line-clamp: 2` with ellipsis, and decide the 360 rule: either socket 36 px and 12/15 type at under 380, or one column's width gets the larger share. The spec's own estimate ("about 15 characters a line") means a third of epic names cannot fit; say so and show the worst case, not the best.

3. The two-hander board contradicts the spec (beta-1280, beta-390, states.png). Main hand shows "Blight", a one-hand-looking blade, under "Two-hand in main hand". The ghost icon in the off-hand is a different, spear-like icon, not the main hand's icon as §4 says. A reviewer cannot sign off the state the owner called out when the mock shows a mismatched pair. Regenerate with a real two-hander (an Arms set) and the same icon in both. Also verify that the Classic client greys the main hand icon into the off-hand slot; I believe the Classic frame leaves the off hand empty. If it does, the spec cites a reference that does not do this, and the state should be the plain empty socket with the label "Two-hand in main hand". Cite what the client does, or drop the claim.

4. The tooltip board is broken (states.png, bottom). The tooltip overlaps its trigger: the socket and icon sit on top of the tooltip title and clip it ("nheart Helm", "h Level 61"). Placement for the left column, the right column (which should flip toward the centre) and the phone is not shown. Spec item 7 and the tooltip judgement cannot be passed on this. Redraw it anchored beside the cell, with 8 px offset, for left, right and phone, with the edge flip rule written down.

## Shoulds

Reads as the character screen
- Desktop reads as the client frame: correct mirror, sockets outward, names toward the crest, weapons bottom, ranged centred. Good. It does not yet read as an "item window": the doll floats on the card ground with no inset stage. Add a stage behind the 960 frame (ground `#0b0e14`, 1 px `--color-line-soft`, radius 12) so it reads as one object, the way the client frames its model. Keep it flat; no new asset.
- Left rows 7 and 8 are empty, giving about 190 px of dead space left of the crest at 1280. It is faithful to the client and sheet. Keep it, but state in the spec that it is deliberate (§5 does) and check it against the card head at 1024 where the dead zone is the most visible.
- Spec §3.2 says the sim head uses the `user` tile. Every mock draws a shield glyph. Make one of them true.

Fit with /account and site alignment
- Card, compact head, gold-tint top edge and secondary button match the /account card. Good.
- On the character page the player-helper job is the gain. The spec puts "+57.1 points across 6 slots" in 13 px muted in the head and each gain in a 12 px detail line. One level of emphasis is on item quality, so the 12 gains the page exists for are quieter than 17 item names. Raise the meta to `--color-gold` in the head, and the gain figure to 13 px 600, so the gains are the second thing seen.
- Ranged shows "+4.0 pts · The Purifier" under an empty "Ranged". That is correct, but the empty label muted next to a gold gain reads as two different things. Keep, and add it to states.png.

Slot states
- Unknown: the red question mark at 55% is still the loudest thing in the left column of beta-1280. Take it to 40% and drop the red saturation, or the three unknowns outshout the real items. The one-note rule (phrase once under the frame) is good. Make the note sit in a `--color-line-soft` strip so it is not mistaken for the footer.
- Unknown items count against item level ("12 of 15 items"): show it on desktop too, it does, good. Say what 15 is (the worn slots), or the figure is confusing next to 17 slots.
- Add the unknown-plus-gain case: can an unknown worn item have a gain line? The spec is silent. Say yes (the gain is for the pick, not the worn item) or no.
- Phone tap targets: the cell is 48 px but the active/tap ground is drawn only on the cell; make sure two columns at 8 px gap do not allow a tap in the gap to hit the wrong cell.

Phone choice (two mirrored columns)
- The choice is right: half the height, matches the client and sheet. It fails only where names are long (blockers 1 and 2). If blocker 1 is solved with a gain row, keep the columns for the plain sim card.
- Item level on one line with "Level 60" and "12 of 15 items" wraps at 360 on the beta page. Draw 360 for the beta page too (only 390 exists).

Widths
- 1024: good, C 240 with names in one line. The longest name (Signet Ring of the Bronze Dragonflight) has about 40 px to spare; a longer name will clamp. Draw one clamp case at 1024.
- 2000: centred 960 frame inside a 1344 card leaves about 140 px dead each side, as the spec says. Acceptable, and the crest is on the screen centre. The board is rendered at a different scale from the others; redo it at the same scale as raid-1280 for the side-by-side rule.

Tooltip
- Everything in §4 on hover/focus/tap is correct on paper; see blocker 4 for the board. Add: tooltip never covers the crest or item level.

Call sites
- CharacterStrip, CharacterGearPanel and GearList are the right three. Add to the audit list, with a one-line reason each: the home "your character" strip if it draws gear, the guild member row/hover card (if it shows a gear set, not a count), and the OG/share image for a saved sim. If none draw a full set, say so after a grep.
- The `gearCopy.slotLabels` deletion changes "Ring 1" to "Finger 1" across the site. Grep the e2e tests and the Discord copy for "Ring 1" before the build lane starts, and list them in the spec.

## Not reviewed
loading-390/1280 and empty-1280 boards were not opened; the build-time reviewer should check them against §4.

## Re-check (2026-10-31, after fix round 1)

Reviewed: beta-360, longest-360, tip-flip-1280, tip-phone-390, tip-phoneflip-390 and spec sections 3.3, 4.2, 10. Verdict: SHIP. No blocker remains.

Blockers
1. Phone gain row: CLOSED. beta-360 draws the upgrade as a full-width 44 px row under each pair, slot word first, pick name on 2 lines, gain right-aligned in gold. "Head / Lionheart Helm / +60.0" and "Shoulder / Truestrike Shoulders / +24.2" cannot be confused. The head figure (+148.8 across 9 slots) matches the 9 gain rows drawn.
2. Clamping at 360: CLOSED on the board. longest-360 uses the worst launch-legal names, shows 3 lines with no name cut, and the rule is in the spec table. See must-fix A for one residual.
3. Two-hander: CLOSED. Arcanite Champion in main hand, the off hand is a plain dashed empty socket with "Two-hand in main hand", no ghost icon. This matches the client behaviour the spec now cites (PaperDollFrame.lua).
4. Tooltip boards: CLOSED. Desktop right column flips to align-end (panel right edge on cell right edge, 8 px gap, flips above near the bottom), panel is about 358 px, clear of the crest and item level. Phone panel is 358 px between 16 px gutters, 8 px below the tapped cell, and flips above near the bottom without covering its trigger.

New findings
- A (must fix at build, not a blocker): longest-360 breaks "Electromagne|tic" with no hyphen mark. Spec 3.3 says a long word is hyphenated, never cut without a mark. The headless render lacks `hyphens: auto`. Build must insert soft hyphens (or ship a hyphenation that works in Chromium on macOS and WebKit) and the 360 acceptance capture must show the "-".
- B (polish): the open desktop tooltip crops the row beside it ("Plaguehound Leggings" is clipped by the panel top). Transient and accepted by 4.2; no change needed.
- C (polish): the phone tooltip hides the rows below the tapped cell. Expected. Confirm that tapping outside or Escape closes it, as in 4.1, in the acceptance capture.

Acceptance: the implementer attaches longest-360 with the hyphen visible and tip-phone at 360 as built.

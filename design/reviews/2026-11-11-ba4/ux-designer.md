# UX review, companion after lane BA4 (2026-11-11)

Reviewed on the 2x captures in this folder, region by region. Reference surfaces: Details! window (status then rows), Windows/macOS tray menus, Finder DMG window.

Verdict: SHIP WITH FIXES (no MUST that blocks the type-scale move; one MUST on the DMG).

## What held
- Type-scale move broke nothing. Status (1600, 1100, 560), Reports, Device, Settings, sign-in, install steps all render with no clipping, wrap or overlap. Nav at 15 item stays on one row at 560; the wordmark drops, the crest stays.
- Hierarchy is right per region. Logging card: 18/display state word first, "3" stat second (28 stat, right-aligned), yellow link third. Device card: "Open your reports" is the single filled primary; "Forget this device" is text-only red, visibly secondary. Settings: labels 13 small are legible against 15 inputs, the label/control step reads clearly.
- Code in sign-in-waiting is the largest element in its card (28 stat mono), with the countdown and bar beneath it. Good.
- "Open your reports" is one name in tray (both OS), Reports link and Device button. Tray label is not truncated.
- "...Advanced Combat Logging, then relog." ends correctly and fits two lines at 1600.
- Install steps: failed state keeps the panel, the red message sits under the button, focus ring on Install is visible (yellow 2 px).

## MUST
1. DMG window (dmg-window.png): the Applications target is a flat blue gradient rectangle with no folder glyph, and the lower third (about 250 px of 864) is empty cream with no instruction. If the capture is a mock, re-shoot from a real mounted DMG so the real Finder folder icon shows. Either way, the painted ground needs a one-line hint ("Drag Forever Sixty to Applications", 15 item, dark ink on cream) between icons and bottom, or the empty third must be cut by shortening the window. A first-time Mac user otherwise has no instruction.

## SHOULD
2. Dim captions on the dark card (Status: "Log written just now", "Addon 2.0.16 . Data 2026.11.09 . up to date", Reports row meta) are the faintest text on the page at 13 small. Lift them one step so they hold 4.5:1 on the card ground; measure and attach the ratio.
3. update-ready: the banner covers the Game folders card mid-row (the "Last played" line is sliced). Pad the scroll area by the banner height so nothing is cut at rest; same for the Settings save bar (WoW Forever row is clipped under it in settings.png).
4. Reports row with "not on the site yet . uploads when the site answers" runs long at 560; confirm it wraps under the title and does not push the Copy link button (reports-1100 and a 560 reports capture are missing; attach the 560).
5. Settings "No changes" disabled Save is dark teal on dark teal; text contrast looks under 3:1. Disabled should read as disabled but stay legible.
6. Tray-notice banner and update banner use the same style; give the Dismiss link a visible focus-visible state in one capture so keyboard parity is evidenced.
7. Missing 560 captures for Reports, Device connected and Settings; only Status and sign-in were shot at 560.

## Checklist
- [x] Type scale applied: no off-scale sizes visible, no regression in wrap.
- [x] One primary per region.
- [x] Reports job named "Open your reports" everywhere captured.
- [x] Advanced-logging help ends "then relog."
- [x] Tray menus legible on both OS themes.
- [ ] DMG labels on a light ground: labels are dark on cream (pass), but folder icon and instruction missing (MUST 1).
- [ ] 560 coverage of all pages (SHOULD 7).
- [ ] Contrast evidence for dim captions (SHOULD 2).

The previous-round comparison file companion-reports.png under 2026-11-10-unify-server/after was not found at that path; compared against 2026-11-10-companion-installers/companion-*.png by listing only (not opened), so no regression claim is made against it beyond the above.

## Round 2 (2026-11-11)

Verdict: SHIP.

Reviewed dmg-window.png, reports.png and advanced-logging-off.png at 2x.

- MUST 1 closed. dmg-window.png is a real Finder window: the app icon, the real blue Applications folder with its alias arrow, dark labels on cream. "Drag Forever Sixty to Applications" sits centred under the icons in dark ink at about 15 item size. The empty lower third is now roughly 100 px of 864 and reads as margin. The arrow between the icons points at the target. Nothing is truncated.
- Reports: the link now reads "Open on the site", yellow, right-aligned on the title line. It no longer repeats "Your reports". Copy link buttons and the row meta are unchanged and nothing wraps or overlaps at 1600. The long Zul'Gurub meta line fits at this width; the 560 wrap stays under SHOULD 4.
- Advanced logging: the card reads "Reports will miss health at death, item level, resources and positions. Turn it on in Options > System > Network > Advanced Combat Logging, then relog." It fits two lines, ends with "then relog.", and the state word stays the largest element in the card. No regression.

Carried follow-ups (not regressions from this lane): SHOULD 2 (caption contrast), 3 (banner and save-bar padding), 4 (560 Reports wrap), 5 (disabled Save contrast), 6 (Dismiss focus ring), 7 (560 captures). Track them in the follow-up list.

Checklist update
- [x] DMG labels on a light ground, folder icon and instruction present.
- [ ] 560 coverage of all pages (SHOULD 7, carried).
- [ ] Contrast evidence for dim captions (SHOULD 2, carried).

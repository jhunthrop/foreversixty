# F5 Sims mock review (design reviewer), 2026-10-31

Verdict: HOLD. Reviewed: f5-sims-ios/and, f5-states-ios/and, xl-ios/and (XL.5, XL.6) against spec 5.5.
Scale note: measurements are read off the PNGs and approximate.

## Must fix
1. XL.5 iOS (AX5): "Started 3 min ago" is cut at the card edge ("agc"). Spec: titles and state wrap whole, age stacks under state. Also the first-row state wraps while the age stays right.
2. XL.6 Android (2.0): the "Full breakdown on foreversixty.gg" ghost label stays about 9 pt while every other string scales. It must scale and keep a 48 dp hit area.
3. XL.6 iOS: caption says "6-up grid holds" but the grid is below the fold, so nothing proves it. XL.5 caption says "Run again grows" but no failed row is shown (either platform). Re-render both.
4. Run again ladder is not drawn: Starting, Started, 5xx failure, offline-disabled (offline only on one row). It is the only new interactive behaviour; add one board per platform.
5. F5.3 states not drawn: loading skeleton, sims.resultFailed + Try again, offline strip, no-gear, unknown-item question mark.
6. Footer copy is inconsistent and unpinned. F5.2 free reads "Start a sim ... and it shows here."; F5.6-F5.8 read "... Server runs keep going after you close the app."; Premium F5.1 shows no footer. F5.7 then says the same sentence three times (serverEmpty, empty, footer). Spec lists keys only: pin verbatim text for footer.premium and footer.free, and drop the repeat.
7. Run again button measures about 41 pt/dp in both F5.1 boards. Needs 44 pt / 48 dp.
8. F5.2b "See plans" renders wipe-red; spec says gold underlined, 44 pt tall hit area.

## Should fix
- iOS F5.1 row 1: state line wraps "takes about 4 / minutes" while row 2 fits. Spec: age drops under state, state never wraps first.
- iOS offline strip sits above the large title; spec says under the title.
- Floating iOS tab bar has no backing: row text and footer ("S...") bleed beside and under it. Check against the Home/BiS boards.
- F5.3b headline in uppercase condensed makes "+41 DPS FROM VIS'KAG" hard to read; check the stat token's case.
- Paging spinner row and failed-page row (sims.failedLoad as last row) not drawn.

## Passes
First-screen answer on both widths; spec/Kind meta rules; unseen dot; class-coloured meta; no rerun or Premium line on F5.3; bar only on bulk; 36/32/40 icons; Android/iOS title patterns match XL.2-XL.4.

## Re-check, 2026-10-31 (after the fix round)

Verdict: SHIP, with the build checks below.

Closed: XL.5 iOS age no longer clipped, state and age stack, Run again visible and growing; XL.6 iOS grid and ghost shown, Android ghost label scales; f5-rerun boards cover Starting, Started, 5xx, lapsed (link and other regions), offline; f5-result-states cover loading, did not load, offline, no gear, unknown item; footers pinned (F5.7 none, F5.1 and F5.8 premium text, F5.2 free text); See plans gold underlined; iOS offline strip now under the title; headline in sentence case; paging spinner and failed-page rows drawn; first-row state line no longer wraps at 402.

Build checks (not mock blockers):
1. Run again reads about 41 pt tall in the renders. Hit area must be 44 pt iOS / 48 dp Android (pad the touch target if the drawn button stays shorter). Same for See plans.
2. iOS floating tab bar: row text shows beside it at the foot (F5.2b, F5.8b). Match the Home and BiS treatment; attach screenshots at 402 and 360.
3. Attach the AX5 and 2.0 captures of a failed row at the real build.

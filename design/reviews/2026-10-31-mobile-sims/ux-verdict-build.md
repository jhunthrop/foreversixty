# F5 Sims build review (design reviewer), 2026-10-31

Reviewed: sims-{premium,premium-large,lapsed,rerun-states,offline,empty}-{ios,and}, result-{run,run-large,gear,failed}-{ios,and} against the SHIP mock and spec 5.5. Other captures (free, loading, error, page-failed, result-loading/offline) were not opened in depth; attach them if they changed.

Verdict: HOLD.

## Must fix
1. sims-rerun-states-and, state 1: the button label renders "," instead of "Starting...". iOS is correct.
2. Android sims-premium, -lapsed, -offline: the first row wraps the state ("Running . takes about 2 / minutes") while the age stays right. Spec F5.1 row 3: the age drops under the state; the state never wraps first. iOS does this correctly.
3. iOS page title is system "Sims" (SF) on the list and on the result ("Tuesday"). Spec and mock: Barlow Condensed uppercase `title`, as Android draws it. Fix in all iOS captures.
4. iOS dates read "Dec 14" and "Dec 12". Spec: web rule, "14 Dec" (Android is right).
5. Android captures show no back arrow (list and result) and no pushed-screen app bar; the mock has it. Re-capture in the real push, with back control, iOS and Android.

## Should fix
- Android AX result (result-run-large-and): the external-link glyph is missing on the wrapped ghost; iOS large puts the glyph at the far right, detached from the label. Keep it after the last word.
- Large iOS: "combina-/tions" hyphenates a mono line; prefer wrapping at spaces ("combinations" on its own line).
- iOS gear icons carry a glow ring that Android does not; match the mock (no glow).
- Run again measures about 43 pt (iOS) and 46 dp (Android) drawn; confirm the 48 dp hit area in code.

## Passes
Premium list order and states; bulk bar and "40 of 64 combinations"; unseen dot; lapsed gold underlined See plans and the sentence alone elsewhere; Could not start it again; offline strip under the title with Run again disabled and sims.offline; pinned footers (premium and free); empty states without a footer; result figure, class-coloured meta, mono error line, 6-up grid, gold corner and caption, sentence-case headline; large-text stacking of state and age; Run again full width and growing.

## Re-check after fixes, 2026-10-31

Verdict: SHIP.

Checked in the refreshed captures: Android Starting... label correct; age drops under the state on the first row, both platforms; iOS titles are Barlow Condensed caps with an inline nav title and back chevron (list and result); iOS dates read "14 Dec". Android has no top bar by the Your gear precedent (system back); accepted. Run again min height is 44 pt / 48 dp in code (SimsLayout.runAgainHeight); accepted on that evidence.

Left as polish, not blocking: iOS gear icons keep a soft glow ring that Android lacks; the iOS inline nav title repeats the big title (same as Your gear and the raid page, so consistent).

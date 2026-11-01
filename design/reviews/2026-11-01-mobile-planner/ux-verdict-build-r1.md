# F13 talent planner, build review round 1 (ux-designer)

Date: 2026-11-01. Spec §5.5, §4, §9. Boards `f13-planner-{ios,and}.png`.

## Verdict: HOLD

Close to the boards on both platforms: crest, chips, points left (gold, kill green at 0), tab totals summed from the
build, F6's cell states, refusal in kill red in the strip with the refused talent's icon, one gold Share, the share
confirm verbatim, the viewer's head and read-only trees, Your builds rows, footnote, empty copy and snackbar.
`Compare with mine` hidden does not block (follow-up).

## MUST
1. **The stepper strip is not 96 tall and resizes as you tap** (106 to 132 across captures), so the tree's viewport
   moves under the thumb; on iOS the bottom row hides behind it. Fix the strip at 96, clamp the rank text to 3 lines and
   the refusal to 2, full text on the tooltip path. Re-capture with a long description.
2. **Largest text: the strip is absent** (Android) or not delivered (iOS). Need the editor with the strip visible and a
   refused tap with the reason in view.
3. **The Reset alert has no capture** on either platform.
4. **F13.4 system share has no capture**; or state in the lane report that it is deferred.

## SHOULD
1. iOS not-found and offline snapshots render a narrow column; confirm the bar.
2. Offline is not an error: muted colour, Try again on both platforms.
3. At 0 left Android keeps gold borders on cells you cannot afford; iOS dims them. Use one rule.
4. Android Your builds at 2.0: the date splits; views beside a 4-line meta.
5. Android share sheet bottom inset; no card.png preview; iOS sheet has no grab handle.
6. `Compare with mine`: add the copy and a status variant for a shared build.
7. Android Undo looks enabled with no history.
8. The Share glyph does not scale at 2.0.
9. Remaining acceptance captures: draft restored after a force quit, F13.5 with Compare and with an old tree_version,
   the swipe states, "over a legal 51", VoiceOver and TalkBack recordings.

# F13 talent planner, build review round 2 (ux-designer)

Date: 2026-11-01. Round 1: `ux-verdict-build-r1.md`.

## Verdict: HOLD, on two missing iOS captures only

No rendering defect needs a code change. Round 1 status:
- MUST 1 (strip height): fixed on both. The strip's top rule sits at one y across plan, refusal and zero; rank text
  clamps to 3 lines, the refusal to 2, verbatim `Needs 20 points in Elemental. You have 14.`; iOS shows all 7 shaman
  rows above the strip.
- MUST 2 (largest text): Android fixed (strip above the tabs and tree, refusal in view, Reset and Share pinned); iOS
  strip placement fixed; the iOS refused tap at AX5 not yet delivered.
- MUST 3 (Reset alert): Android fixed (`Clear all 49 points?`, red Clear, Cancel, dimmed planner); iOS not delivered.
- MUST 4 (F13.4 system share): accepted as deferred to the on-device pass, stated in the lane report.
- SHOULDs 1 to 5 and 8 fixed (full-width muted offline with Try again, one rule at 0 left, Android Your builds at 2.0,
  share sheet insets and grab handle, Share glyph scales).

## MUST
1. iOS refusal at AX5 (planner-refusal.ax5): the reason in view, 2 lines, kill red, Reset and Share pinned.
2. iOS Reset alert capture: `Clear all {n} points?` with Cancel and a destructive Clear over the dimmed planner.

## SHOULD
1. iOS AX5 header takes 38% of the screen; keep `2 left` with the crest and let the pickers share a row; add a
   capture scrolled to the tabs.
2. Undo state evidence: a fresh planner with Undo at 40%, then 100% after one tap, on each platform.
3. iOS pill hairlines at the pills' ends: a renderer artefact or a real stroke; check before the device pass.
4. Android Reset dialog: tighten the space between title and buttons to the Material 24 dp rule.
5. Android builds-large: the meta's `·` dangles at a line end; keep it with the date.
6. Carried over: Compare with mine copy, draft after a force quit, F13.5 with Compare and an old tree_version, swipe
   states, over a legal 51, VoiceOver and TalkBack recordings, F13.4 on device: the on-device acceptance pass.

# F13 talent planner, build review round 3 (ux-designer)

Date: 2026-11-01. Spec §5.5, §4, §9. Rounds 1 and 2: `ux-verdict-build-r1.md`, `ux-verdict-build-r2.md`.

## Verdict: SHIP

Both round 2 MUSTs are delivered and match the spec; no regression, no blocker.

1. **iOS refusal at AX5, closed** (`build/ios/f13-2-refusal-large-ios`): the strip above the tabs and tree, Reset and
   Share pinned, the refused talent's icon, `Elemental Reach 0/2`, the full `Needs 20 points in Elemental. You have
   14.` in kill red; `37 left` on the crest row, Orc and Level 60 on one row.
2. **iOS Reset alert, closed** (`build/ios/f13-1-reset-ios`, a UI-test screenshot of the running app):
   `Clear all 51 points?`, Cancel and a red Clear over the dimmed planner.

Round 2 SHOULDs closed: Android Undo at 40% and 100% (`f13-1-fresh-and`, `f13-1-undo-and`); the Android Reset dialog's
spacing; the Android dangling separator; the iOS pill hairlines (a snapshot artefact, absent from the real screenshot);
Race and Level at AX5.

## SHOULD, for the follow-up pass
1. iOS Your builds at AX5 still wraps the meta as `Orc Shaman ·` / `30 / 3 / 18 · 4 Nov`; the separator should travel
   with the following token, as on Android.
2. Confirm iOS Undo goes to full after the first tap (on-device pass).
3. On a guide-link planner the loaded-build note replaces the selected talent's text in the strip: clear it after the
   first edit, and keep the full text on the strip tap.
4. On-device acceptance pass: F13.4 system share; draft after a force quit; F13.5 with Compare and an old
   tree_version; the swipe states; over a legal 51; VoiceOver and TalkBack recordings; the Compare with mine copy.

# F13 talent planner, build review round 1 (wow-player)

Date: 2026-11-01. Captures: `build/and/*.png` and the iOS `PlannerSnapshotTests` goldens of round 1.

## Verdict: HOLD

The numbers are right and the page does its job ("Where do my 51 points go, and what can I take next?"). It holds on
one sentence a player reads as wrong, and on largest text hiding the only place the refusal reason appears.

**Numbers checked (all pass):** 30 + 3 + 16 + 2 left = 51 (Android plan); 14 + 3 + 18 + 16 left = 51 (Android
refusal); 14 + 37 left (iOS refusal); 30 + 3 + 18 + 0 left, green (iOS zero); shared build and Your builds totals all
51; the Elemental badges sum to 30; tier gating matches (25 points above Elemental Fury, 14 points lock the 15 and 20
rows); prerequisite arrows run the right way and dim when locked; Elemental Fury's rank text is the client's.

## MUST
1. **Tier wording, both platforms and the site.** "Tier 4 of Elemental needs 20 points in Elemental first" names the
   fifth row tier 4: a 0-based index shown to a player. Fix at the source (`web/src/lib/planner/rules.ts`), with the
   spec's sentence `Needs 20 points in Elemental. You have 14.`
2. **Android `f13-1-plan-large-and`:** the stepper strip is not visible at the largest text, and it is the only place a
   refusal reason appears and the TalkBack route. Re-capture with the strip visible; the iOS AX5 capture is still owed.
3. **iOS plan, zero and refusal:** the tree's last row is hidden under the strip and its arrow runs into the strip.

## SHOULD
4. iOS share sheet over black; record it over the planner.
5. "Talent order" is an implementation word to a player ("Talent points" would read better); the title bullet repeats the field.
6. Your builds: an untitled row repeats its totals in the meta line.
7. An untouched, unlocked talent in the strip should show the rank-1 text.
8. Confirm VoiceOver and TalkBack say a locked talent's requirement.
9. iOS shared-build: a blank band above the title in the snapshot; check it is not on device.

**Data I do not believe:** none.

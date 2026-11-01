# F13 talent planner, build review round 2 (wow-player)

Date: 2026-11-01. Round 1: `player-verdict-build-r1.md`.

## Verdict: HOLD, on one missing capture only

Round-1 blockers fixed wherever visible:
- MUST 1, tier wording: `Needs 20 points in Elemental. You have 14.` on Android, Android large and iOS; the 14 is
  true on all three. It reads like the client.
- MUST 2, Android large text: the strip sits under the header with the red sentence, readable without scrolling;
  Reset and Share pinned. iOS AX5 plan has the strip at the top; the AX5 refusal capture is missing.
- MUST 3, iOS row pitch: all 7 rows show; the last row clears the strip and its arrow ends on the icon.
- SHOULD 4 (iOS share over the planner) fixed; SHOULD 5 ("Talent order") accepted; SHOULD 6 (meta totals) fixed.

Numbers rechecked, all pass: 30+3+16 with 2 left; `Clear all 49 points?`; 14+3+18 with 16 left; 30+3+18 with 0 left in
green; iOS 14+0+0 with 37 left; the shared build header equals its tabs; every Your builds total is legal; Elemental's
badges sum to its tab; tier gating at 14 points dims the 15, 20 and 25 rows; arrows run the right way.

## MUST
1. iOS AX5 refusal capture: the red sentence visible without scrolling, points left on screen.

## SHOULD
2. Elemental Mastery's text reads `... by 0.` at 0/1; check the talent data (see below).
3. Untitled Your builds rows still say the class twice (`Orc Shaman 31 / 20 / 0` over `Orc Shaman · 29 Oct`).
4. `Share anyway` gives no reason when points are unspent.
5. The + button looks live on a maxed or unaffordable talent; dim it and keep it tappable.
6. iOS AX5 plan: stack the pickers on one line so the tree shows sooner.
7. `Reset` (Android) and `RESET` (iOS): one case.
8. The refusal wraps `... You / have 14.` at default size.
9. Confirm VoiceOver and TalkBack speak a locked talent's requirement.
10. iOS shared-build blank band: closed (status bar and island space).

## Data I do not believe
- Elemental Mastery's rank text (`... reducing the cast time of Lightning Bolt and Lava Burst by 0.`): may be Forever's
  rework, but the `0` looks like an unresolved placeholder in the talent data; for the accuracy reviewer.
- The fixture's `Saved 4 Nov` is after the capture date (invented fixture data).

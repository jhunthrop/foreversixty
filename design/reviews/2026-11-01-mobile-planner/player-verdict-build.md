# F13 talent planner, build review round 3 (wow-player)

Date: 2026-11-01. Rounds 1 and 2: `player-verdict-build-r1.md`, `player-verdict-build-r2.md`.

## Verdict: SHIP

The one MUST is cleared and nothing regressed.

**iOS AX5 refusal (`build/ios/f13-2-refusal-large-ios.png`):** the strip is the first thing under the header:
`Elemental Reach 0/2`, the red `Needs 20 points in Elemental. You have 14.`, then − and +; `37 left` on screen; the
Elemental tab total 14 visible; Reset and Share pinned; 14 + 0 + 0 + 37 = 51. A player gets the reason without
scrolling, and it is the right sentence.

**Round-2 SHOULDs rechecked:** untitled Your builds rows show only the date (`Orc Shaman 30 / 3 / 7` over `4 Nov`);
− dims at rank 0 and + stays live on a refusal; a fresh tree shows rank-1 text (Convection 0/5); at AX5 Race and Level
share a row.

**New captures:** Android fresh: `51 left`, 0 / 0 / 0, Undo and Reset dimmed, only the first row gold. Android undo:
one point in Convection, `50 left`, Elemental 1. Android refusal regression check: 14 + 3 + 18 with 16 left, the 15 and
20 rows dim, arrows gold into the locked 0/3 and grey below.

## SHOULD (none block)
1. VoiceOver and TalkBack for a locked talent: confirm the spoken line includes `Needs 20 points in Elemental`.
2. `Share anyway` with points unspent: say `2 points unspent.`
3. `Reset` / `Cancel` / `Clear` casing differs between the platforms (each platform's button style).
4. iOS AX5 plan: the tree is below the fold at the largest size, acceptable at AX5.

## Data I do not believe
- Lava Burst in Convection's and Elemental Mastery's text, and Elemental Mastery's cast-time line, differ from Era.
  If these are the Forever client's own strings (Lava Burst is new in Forever), the app is faithful. For the accuracy
  reviewer.
- `Saved 4 Nov` in the fixture is launch day, after the capture date (invented fixture data).

# /setup steps: the companion second, the addon seen by the companion (2026-11-10)

Lane `lane/setup-steps`. Captures in this folder, at 1440 and 390, full page and the plate alone:
`signed-out`, `no-companion` (signed in from Battle.net, no companion), `companion-no-addon`, `companion-addon`,
`addon-outdated`. Taken by `web/tests/captures/setup-steps/capture.spec.ts` on the setup fixtures at the mocks' clock.
Spec: `design/specs/2026-11-07-setup.md` section 0B.

## Final verdicts

| Reviewer | Verdict |
| --- | --- |
| ux-designer | SHIP |
| ai-slop-reviewer | PASS (round 2 "PASS WITH FIXES"; its should and wording nit applied in round 3) |
| wow-player | SHIP |

## Round 1

- **ux-designer: SHIP.** Shoulds: the gold behind line read like a link; "Not on JHUNT-PC yet" was a third muted line under the note; the spec was stale. Nit: a string-equality branch standing in for a flag.
- **ai-slop-reviewer: FAIL.** Musts: the behind line had no action and its gold clashed with the green tick; the "What it does in the game" grid shot empty. Should: the scanned-none line doubled up with the note.
- **wow-player: NO-SHIP.** Must: before any scan, the main audience (Battle.net players who installed the addon first) was still told the addon was undone, and the gold step was one the page called "Optional". Shoulds: name the `_forever_` folder; the behind state read as done and behind at once.

## Fixes

- Before any companion has looked: the note, then "Already installed? The companion sees it." (wow-player's must; kept over the slop reviewer's later nit to cut it).
- "Optional." is gone from the companion's note, so the gold step is never called skippable.
- A scan that found none: one line, "Not in _forever_ on JHUNT-PC yet", replacing the note.
- Behind the feed: still ticked (owner ruling), a plain "2.0.37 is out" line, then quiet gold links "Update from CurseForge ›" and "Update from Wago Addons ›".
- `addonLine` returns a `scanned` flag; the spec has the 0B amendment; captures load lazy pictures before the full-page shot.

## Round 2 and 3

- ux-designer: SHIP holds.
- wow-player: SHIP.
- ai-slop-reviewer: PASS WITH FIXES. Should: make the Update buttons quieter. Nits: cut "Already installed?" (declined: wow-player's must) and avoid "JHUNT-PC's _forever_" (now "Not in _forever_ on JHUNT-PC yet"). Both applied in round 3.

## Out of scope (other lanes)

- Three "Download for Windows" placements (head, plate, The companion) and "Other platforms" wrapping at 390: lane/companion-installers.
- `parity-setup` at 390 differs only in `[site-footer]` (mock 257 px, build 277 px): main's operator footer, not this change.

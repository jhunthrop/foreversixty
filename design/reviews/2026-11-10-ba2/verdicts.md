# BA2 identity and listing: review verdicts (2026-11-10)

Lane `lane/ba2-identity`. The evidence is in this folder:
- `identity-sheet.png`: every identity image, the Discord sidebar mock at 240x135 with the name overlay, the 20 px role icon at 4x, and the live link previews in `og-live/`.
- `store-app-store.png` and `store-play.png`: the store screenshots.
- `gallery-plate-preview-stale-capture.png`: the addon gallery plate design. The capture in it is an old JPEG on purpose.
- `compose_sheets.py` draws the sheets.

There were two rounds. Round 1 fixes are in `ffa53a78f`, with polish after it.

## Verdicts by asset set (round 2)

| Asset set | brand-consistency | ai-slop | ux-designer | wow-player |
|---|---|---|---|---|
| Identity (Discord, store icons, feature graphic, OG) | CONSISTENT | PASS | SHIP | n/a |
| App Store screenshots | CONSISTENT | PASS, gated | SHIP | HOLD, app content |
| Play screenshots | DRIFT, empty portrait ring | PASS, gated | HOLD 01 and 04 until recaptured | HOLD, app content |
| Store text (App Store, Play) | n/a | PASS | SHIP | HOLD on the plates only; text fixes landed |
| Addon listing text | n/a | PASS | SHIP | SHIP |
| Addon gallery pipeline | CONSISTENT, held | PASS, gated | SHIP | n/a |

## What still holds, and why it is outside this lane

Every remaining HOLD or DRIFT comes from the apps or the addon themselves, not from the generated
images or the text. `design/listing/APPLY.md` makes each one a gate before upload, and
`capture-apps.sh` plus `build-store.mjs` rebuild the set when they are fixed.

1. **Android portrait.** It draws an empty ring when the portrait does not load: Play 01 and 04, the profile button and the gear medallion. `GearRoute.kt` passes a portrait painter with no fallback. This is apps work (lane F).
2. **Retired forced caps on the raid page.** Lane F.
3. **App copy.**
   - Home's row says "Sign up In, Late or Out." with Bench missing.
   - "2 levels to band 60".
4. **Fixture character.** A level 58 Fury Warrior is offered a gun and hunter pants, and has "0 of 17 slots on the list". A level 60 character with warrior picks and a part-done list is needed. The picks also go to the accuracy reviewer.
5. **Addon gallery.** It waits on BA1's new theme and the owner's twelve lossless in-game shots (`RECAPTURE.md`). Each plate must fill the stage. The old navy plates are removed.

## Owner rulings the reviewers accepted

- "Play your class better." stays, but only in the App Store subtitle and the addon's one-line summary.
- The tavern and road art are the owner's own commissioned art (2026-11-09), used as background only. The dark right edge of the invite background is the tavern's beam in that art, not a render fault.
- Captions are centred Marcellus, following the site's page head.

## Follow-ups noted by reviewers

- A clean lock-screen raid notification shot: the headline feature. The simulator's home screen shows test-runner apps, so it needs a lock-screen capture.
- The link previews are about 1.1 MB PNG each (`web/src/lib/og.ts`). A JPEG card would be about a tenth of that. This is a web change, outside this lane.
- The "Fan-made player tool" sub-line on every OG card is small at preview size (ai-slop, minor).

# Companion round 3 captures (15 Oct 2026)

Source: main at 5a6da634 (the round-2 fix lane merged). Same rig as round 2 (`r2/captures.md`): the headless companion against the fake API in Chrome at 565x400 (screenshots 800x566), with a no-op `xdg-open` so the ordinary waiting state shows. Only the surfaces the fix round touched were re-taken; everything else in `r2/` still stands. "Companion 0.0.0" is the dev build string.

## Re-taken by the controller on the merged build

| file | state | what changed |
|---|---|---|
| device-not-connected.jpg | Device, not paired | text-only "Sign in with Battle.net" button (no glyph) |
| status-not-connected.jpg | Status, not paired | button now "Sign in with Battle.net"; Addon row muted + reworded |
| reports-not-connected.jpg | Reports, not paired | "Sign in first." + the same button |
| device-waiting.jpg | Device, ordinary waiting (browser opened) | "Waiting for you to sign in..." |
| status-advanced-off.jpg / status-advanced-off-folders.jpg | Status, advanced logging OFF | the instruction sentence only in the headline; the folder row reads "Advanced logging: OFF" |
| status-logging.jpg / status-logging-folders.jpg | Status, logging | Current report row shows the zone as plain text; the headline keeps the link |

## From the fix round (lane build, same code)

| file | state |
|---|---|
| device-forget-confirm-focus.jpg | Forget two-step by keyboard, ring on Cancel, buttons on one row |
| device-forget-confirm-360.jpg | the same at 360 px |
| device-forget-after.jpg | after Forget (revoked on the site) |
| poll-failure-bar.jpg | the companion stopped while the page was open |
| footer-update-ready.jpg | footer with "Update ready" |
| reports-empty-paired.jpg | Reports, paired, no reports yet |
| site-connect-<state>-<360|1280>.png | the site page: not-signed-in, pending (with the code-check line), success, expired, cancelled; the 470 px card minimum removed; the expired/cancelled button font fixed |

Not capturable headless: the native folder dialog and its invalid-folder error, the tray, the first-close bar.

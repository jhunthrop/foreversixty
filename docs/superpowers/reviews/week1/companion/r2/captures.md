# Companion round 2 captures (15 Oct 2026)

Source: main at 84b66582 (the three lanes merged), the companion built with `-tags nogui` and run `-headless` in WSL against the test fake API (`internal/fakeapi`) with a scratch home; the UI opened in Chrome at 565x400 (screenshots are 800x566, scaled 1.42x) and one at 900x700. Because WSL has no browser, every sign-in lands in the "browser blocked" variant of the waiting state (the ordinary waiting state is the same minus the link box and with "Open the browser again" in place of "Copy link"). Linux has no keychain here, so "Saved in" shows the file fallback in amber. "Companion 0.0.0" is the dev build's version. Not captured (need the real Windows window): the tray menu/tooltip, the first-close bar, the native folder dialog, the Windows Credential Manager wording.

## Companion window

| file | page / state | notes |
|---|---|---|
| device-not-connected.jpg | Device, not paired | first launch opens here |
| status-not-connected.jpg | Status, not paired | red headline + Connect button |
| reports-not-connected.jpg | Reports, not paired | |
| settings-no-folder.jpg | Settings, no folder | |
| device-waiting-browser-blocked.jpg | Device, signing in, browser could not open | code, countdown, link box, Copy link / Cancel |
| device-failed-denied.jpg | Device, failed (denied on the site) | red bar, Try again, code fallback below |
| device-connected-flash.jpg | Device, just connected (4 s flash) | then switches to Status |
| status-no-folder.jpg | Status, paired, no game folder | |
| device-connected.jpg | Device, paired, idle | |
| device-forget-confirm.jpg | Device, Forget two-step | Escape cancels |
| device-details-open.jpg | Device, Details disclosure open | |
| status-advanced-off.jpg / status-advanced-off-folders.jpg | Status, folder with advanced logging OFF | amber headline; the folder row marked |
| status-logging.jpg / status-logging-folders.jpg | Status, logging a report (3 fights) | green headline, report link, Watching tag |
| status-logging-900x700.jpg | the same at 900x700 | |
| settings-folder-list.jpg / settings-folder-list-bottom.jpg | Settings with one folder | Save disabled until a change |
| reports-rows.jpg | Reports with one in-progress report | |
| reports-focus-after-5s.jpg | the report link focused, 5 s later | focus survives the poll |
| nav-focus-ring.jpg | keyboard focus on Settings while Device is current | ring vs underline |

## Site /connect page (Playwright, mocked API)

`site-connect-<state>-<width>.png` for not-signed-in, pending, success, expired at 360 and 1280 px.

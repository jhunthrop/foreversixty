# Companion window: every screen and state (batch G)

This lists every screen and state of the desktop companion. The window opens at 800 x 620 and can be resized down to 560 x 480 (`internal/shell`). It has four pages: Status, Reports, Device and Settings. Next to the window are the tray icon, the tray menu and the app icon.

- **The state behind each capture:** the snapshot in `companion/ui/testdata` that the Go tests pin, plus the patch or action in `design/mocks/companion/shoot_ui.cjs`.
- **Words:** `companion/ui/copy.js` holds the window's copy, and `internal/app/copy.go` holds the headline, addon and error copy.
- **Captures:** `round-N/<state>.png` at 800 x 620. The states marked "+1100" are also captured at 1100 x 800, and "+560" at the minimum 560 x 480 (`<state>-560.png`).

## First run and sign-in

| State | What shows | Capture |
|---|---|---|
| First launch, not connected | The window opens on Device. The lead panel shows a red dot, "Not connected", the intro sentence, the teal **Sign in with Battle.net** button with the account glyph, the password note, and "Having trouble? Paste a code from the site". | `first-run-device` |
| Status while not connected | The Status headline says "Not connected" with a red dot and a Sign in button. | `first-run-status` |
| Paste a code (open) | The code field, its help line and **Connect with code**. | `pair-paste-code` |
| Paste a code, refused | The error sits under the form in the bad tone. | `pair-paste-code-bad` |
| Waiting for the browser (+560) | A spinner and "Waiting for you to sign in…", the browser line, the code plate, the countdown and its bar, **Open the browser again** and Cancel. | `sign-in-waiting` |
| Last minute | The countdown turns amber. | `sign-in-last-minute` |
| Site unreachable while waiting | "Having trouble reaching the site. Still trying..." | `sign-in-trouble` |
| Browser blocked | The sign-in link in a read-only field, and **Copy link**. | `sign-in-browser-blocked` |
| Sign-in refused or expired | The reason in the bad well, and **Try again**. | `sign-in-failed` |
| Status while signing in | The headline shows a spinner, "Waiting for you to sign in..." and **Show the sign-in**. | `sign-in-connecting-status` |
| Signed out by the site | Status shows "Signed out by the site" with a red dot. Device shows the reason well. | `signed-out-by-site`, `signed-out-device` |
| Just connected | "Connected. The companion is ready." shows for four seconds, then the window goes to Status. | `connected-flash` |

## Device (connected)

| State | What shows | Capture |
|---|---|---|
| Connected | A green dot and "Connected as Name#1234", the device and API on one mono line, then Saved in, Uploaded from this PC, Addon last seen in game and Log parser. **Open my reports**, and **Forget this device** in the wipe colour. | `device-connected` |
| Forget, confirming | The question, the red outline **Forget** and **Cancel**. Focus sits on Cancel and Esc closes it. | `device-forget-confirm` |
| Forgotten, site unreachable | Back to Not connected, with the line "Forgotten on this PC ... Account > Devices". | `device-forgotten-local` |

## Choosing the WoW folder

| State | What shows | Capture |
|---|---|---|
| No install found | Status shows a red headline, "No Classic beta install found", and **Open Settings**. | `no-folder-status` |
| Settings with no folder | "No Classic beta install found yet." and **Add a folder...** | `no-folder-settings` |
| Folder added | The new folder row, plus "Added." in the good tone. | `folder-added` |
| Not a WoW folder | The error in the bad tone under the list. | `folder-not-wow` |

## Watching and uploading

| State | What shows | Capture |
|---|---|---|
| Ready, not logging | A grey dot and "Ready" with the /combatlog line. | `ready` |
| Advanced logging off | An amber headline. The folder row reads "Advanced logging: OFF" with the amber dot. | `advanced-logging-off` |
| Logging, credited | A green live dot and "Logging" with who is logging, the fights the site holds as the headline's figure, and "Open this report on the site". The rows show the report (zone, start, visibility, when the log was last written), the addon (one line when current, **Fetch my builds from the site** under its name) and characters sent. The game folders follow. | `logging-credited` +1100 +560 |
| Logging, name not read yet | "Your name is read from the first fight." | `logging-no-name` |
| Logging, not credited | An amber headline, "Logging · not credited to you yet". | `logging-not-credited` |
| First fights on their way | "Uploading 3 fights…" | `uploading-first` |
| Uploading fights | The figure and the Queued row. | `uploading` |
| Upload held (a fight failing holds the rest) | The headline and the tray turn amber, "Logging · an upload is retrying"; the Queued row in amber: "Fight 4 failed 2 times and retries on its own; the rest wait behind it." (`pipeline.held`, new in this batch). | `uploading-held` |

## Reports (upload done and failed)

| State | What shows | Capture |
|---|---|---|
| Reports list | A live dot on the open report, then zone, when and fights. The usual character is named once, by the Status headline; a row names only another character, or "not credited yet" in amber. **Copy link** is a hairline ghost. A row the site does not have yet shows a ring and "not on the site yet · uploads when the site answers". | `reports` +1100, `reports-560` |
| Link copied | The button reads "Copied." with a check for two seconds. | `reports-copied` |
| No reports yet | The logs glyph, "No reports yet." and the /combatlog line. | `reports-empty` |
| Signed out | "Sign in first." and the Sign in button. | `reports-signed-out` |

## The addon (Status's Addon row)

| State | Capture |
|---|---|
| Addon missing: the **Install the addon into Classic beta** button | `addon-missing` |
| Installing (spinner, "Installing…") | `addon-installing` |
| Install refused while the game runs | `addon-install-game-running` |
| Installed, takes effect at next login | `addon-installed` |
| Ready, installs when the game closes | `addon-ready` |
| Up to date | `addon-current` |
| A build is waiting for a character | `addon-build-waiting` |
| Installed, no character yet | `addon-no-character` |
| Sync over a day old (amber) | `addon-stale` |
| Never synced | `addon-never` |
| Check for builds now, checking | `addon-checking-builds` |

## Settings

| State | Capture |
|---|---|
| The form: visibility, keep the addon up to date, credited to, game folders, and the Save strip | `settings` +1100 |
| Unsaved changes (Save enabled) | `settings-unsaved` +560 |
| Saved ("Saved. New reports will be guild only...") | `settings-saved` |

## Errors and offline

| State | What shows | Capture |
|---|---|---|
| The companion's service stops answering | The red bar under the band, "Status unavailable", and the last numbers dimmed under "Last known". | `offline-status` |
| No answer at first launch | The red bar alone. | `offline-first-launch` |
| First close of the window | The tray notice bar, "Still running in the tray...", and Dismiss. | `tray-notice` |

## The updater

| State | What shows | Capture |
|---|---|---|
| Update downloaded | The bar above the footer, "Update ready. Quit the companion from the tray icon and open it again to install it.", with Dismiss. | `update-ready` |
| Just updated | The footer reads "Companion 0.2.13 · just updated" for half a minute. | `just-updated` |

## Tray and app icons

| Surface | What shows | Capture |
|---|---|---|
| Tray icon, four tones | Seal v3, with a green, amber or red dot in the window's tone tokens. "Off" is the plain seal. Windows uses ICOs from 16 to 48 px, and macOS and Linux use a 32 px PNG (16 pt at 2x). | `icons` |
| Tray menu | The headline title (disabled), Open Forever Sixty, Open my reports on the site, and Quit. The tooltip reads "Forever Sixty · Logging". | `tray-menu` |
| App icon | macOS: `.icns` on Apple's 824/1024 grid, Dock and Finder. Windows: `.ico` from 16 to 256 px, taskbar and Explorer. Linux: PNGs from 16 to 512 px. | `icons` |

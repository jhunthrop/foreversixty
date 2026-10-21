# Companion changelog

## 0.2.12

- Settings never hides a field under the Save bar; the sign-in screen says the browser has opened; the Status title says when a log is not credited to you yet.

## 0.2.11

- The companion looks like the site: its own typefaces, a dark title bar that never flashes white, a fixed frame with a scrolling page, a Status headline you can read from across the room, and a tray icon that shows green, amber or red.

## 0.2.10

- While a log is being followed, the companion tells the site it is alive once a minute, so a report page can say when it has gone quiet.

## 0.2.9

- Status says plainly whether the current report is credited, and to whom; installs you are not logging no longer warn about advanced logging; the companion checks for builds right after you log out, and the Addon row has a Check-for-builds-now button.

## 0.2.8

- A report opened before this update is credited after the update too: the stored logger name is re-read the new way.

## 0.2.7

- Reports are credited to two-word Forever names: the log writes "Bow-Jackzon-ClassicBetaPvP-", and the companion now reads it as Bow Jackzon. When a log cannot be credited, Status says whether the name or the realm is what the addon has not sent.
- The update check finds the newest companion release even when the repository's latest release is the nightly data addon.

## 0.2.6

- After an update, the companion starts the new build itself instead of exiting; the log says what it launched.
- The Status headline and the tray warn about advanced logging only for the game folder the companion is watching, not for every install on the disk.

## 0.2.5

- The companion waits for the game to finish writing its saved variables before it delivers the inbox, and
  writes the file atomically.
- Status shows your last sync and who each report is credited to; the stopped-responding note says
  where the tray icon is.

## 0.2.4

- When the site no longer accepts this computer's sign-in (the device was removed, or the account
  reset), the companion says so and asks you to sign in again instead of showing Connected.
- Launching the companion again just opens the running one, instead of starting a second copy that
  tails the same log and uploads every fight twice.

## 0.2.3

- Reports are credited to the character that wrote the log, read from the combat log itself; pick
  one in Settings only to override.

## 0.2.2

- The companion now watches the game install you played most recently, not the first one it
  found. The Status page tags it Watching and shows when each folder was last played.
- The Forever Sixty seal is the app's icon: window, taskbar, tray, Explorer, and a macOS app bundle.

## 0.2.0

- Sign in with Battle.net: the Device page opens your browser on the site, shows a code to
  compare, and connects by itself once you press Connect there. Paste-a-code stays under
  "Having trouble? Paste a code from the site".
- Status opens with one headline (Not connected, No game folder found, Advanced logging is off,
  Ready, Logging), repeated word for word in the tray. Unpaired, the window opens on Device.
- Rows renamed (Current report, Fights uploaded, Queued), Engine dropped, the addon row says
  "Last synced 4 minutes ago" and warns in amber after a day. A footer shows the version, an
  update that is ready, and the site link.
- Game folders: a list with the flavour name and a Watching tag; Settings adds folders through
  the system folder dialog and removes only the ones you added.
- Device page: five states, a two-step "Forget this device" that also removes the device from your account, "Saved in" names
  the keychain (or says a private file is the fallback).
- Reports: relative dates, Copy link, a link to all your reports on the site.
- Closing the window keeps the companion in the tray (Windows only in this release); the first
  close says so. On macOS and Linux, closing the window quits the companion; keep it open while
  logging.
- Every external link opens in your system browser.
- Errors are sentences, never exception text; the page no longer loses focus on its two-second
  refresh.

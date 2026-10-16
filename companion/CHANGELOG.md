# Companion changelog

## 0.2.2

- The companion now watches the game install you played most recently, not the first one it
  found. The Status page tags it Watching and shows when each folder was last played.

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

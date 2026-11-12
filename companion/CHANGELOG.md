# Companion changelog

## Unreleased

- A report that begins partway through a combat log the game was already writing (the companion started, or the tail moved to this game folder, mid-session) is now parsed with the header the game wrote when logging began: the companion reads the last `COMBAT_LOG_VERSION` line behind the report's start out of the file, instead of leaving the engine to infer the layout from field counts. A night started that way on 12 Nov had a quarter of its lines unreadable.
- Every call to the site names the companion and its version (`foreversixty-companion/<version>`), so the site's logs can tell a companion's uploads from the browser's.
- The companion reads and uploads only the live game, `_forever_`. The Classic beta ended at launch, and its `_classic_beta_` folder is now treated like any other game: never found, never offered, and refused by name if you pick it, so beta characters no longer reach the site beside your live ones. A stored beta folder is dropped at start.

## 0.2.44

- The window takes the site's new look: the painted blue band with the seal and the four pages, the brown ground, the framed headline, the teal buttons, and the site's own type and icons.
- Status leads with the fights the site holds. The Current report row reads zone, start and visibility, with when the log was last written.
- An upload that keeps failing now says which fight is holding the queue and how often it has failed. The headline and the tray turn amber while it does.
- The app and tray icons are the new seal on Windows, macOS and Linux, at a fraction of the old icons' size. Each build now carries only its own platform's tray icons.
- After this update the new icon shows straight away. On a Mac the app's bundle is touched, so the Dock and Finder read the icon again. On Windows, Explorer is told that icons changed, so shortcuts and the taskbar refresh without a reboot. Both steps are best effort and never hold up an update.
- A game folder is now recognised by the `.flavor.info` file the game writes in it (`wowf` is the live game), so a renamed or custom install path works. A `_forever_` folder holding another game is refused, and the refusal names that game.

## 0.2.43

- WoW Forever went live in its own game folder, `_forever_`, and the companion now finds it: it lists, watches and keeps the addon up to date in `_forever_` first, and still in the Classic beta's `_classic_beta_` for testers. With both, the live game is watched unless the beta was played more recently. Every message that asks for a folder now names `_forever_`, and a picked `_forever_` folder is accepted instead of refused.

## 0.2.39

- Raid answers from the game reach the site: when you answer a raid with the addon panel's Tonight buttons, the companion sends the answer with your next sync, and once the site has taken it (or refused it, because you answered later on the site, the app or Discord, or the raid closed) it takes the answer out of the addon's saved file. Nothing else in that file is touched. The inbox tells the addon this companion carries raid answers, so the panel offers them only with a companion that sends them.

## 0.2.37

- A Mac companion: one app for Apple silicon and Intel Macs on macOS 12 or later, signed with Commish LLC's Developer ID and notarized by Apple, shipped as a disk image to drag into Applications. It finds the game in /Applications/World of Warcraft, sits in the menu bar, hides when its window is closed, has the Edit menu that makes Command-V paste a pairing code, keeps the device token in the macOS Keychain, waits for the Mac client to close before it updates the addon, and updates itself by replacing the whole signed app.

## 0.2.22

- The Windows build is signed by Commish LLC, the company behind Forever Sixty, through Azure Artifact Signing: Windows names that publisher in its prompts, and SmartScreen's "unknown publisher" warning goes with it. The README lists the two platforms that ship, Windows and Linux.

## 0.2.21

- When the game closes the companion reads the releases again and installs the newest signed addon, not the one it staged hours earlier; it reads the installed version from the addon folder first and never installs an older or equal one, saying so once in the log. The Status line names the version that will go in and follows a newer release.

## 0.2.19

- The Addon row no longer tells you to install an addon it shows as installed: with the addon in place and no export yet it says "No character yet" and to log out once in game. A first run on a machine that still holds an old sign-in reads as a first run, not as a sign-in that lapsed. The log says each thing once: an addon update waiting for the game to close, and a step that keeps failing, are written when they start and not every minute or second.

## 0.2.18

- A build sent from the site while you are in game is in the addon's inbox seconds after you log out, not a minute: the companion writes it as soon as the game has finished saving, so logging straight back in shows it. The Status card says a build is waiting, and for whom, until a logout shows the addon held it.

## 0.2.17

- The companion knows only Forever variants: the Classic beta is listed, watched and written to, and Retail, Anniversary and other folders are not; a picked folder of another game is refused with a sentence that says why; the addon updater logs one line for unsigned releases.

## 0.2.16

- The companion keeps the Forever Sixty addon and its data up to date in every game folder, when the game is closed, and can install the addon where it is missing.

## 0.2.14

- A log is one row in Reports; the counters never say uploaded and queued for the same fights; Settings stays legible in a small window.

## 0.2.13

- The Status card says what to do when a log is not credited to you; Settings shows whole folder names and a Save strip that says No changes; the update bar says uploads pause during the restart.

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

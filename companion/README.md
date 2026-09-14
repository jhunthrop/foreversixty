# Forever Sixty companion

A desktop app that watches World of Warcraft's combat log, uploads each fight to
[foreversixty.gg](https://foreversixty.gg) as it ends, and keeps the ForeverSixty addon in step
with the builds you chose on the site. One static binary per platform, a tray icon, and a small
window. It never modifies the game's log and never sees your Battle.net password.

## Install

Download the binary for your machine from the
[latest release](https://github.com/jhunthrop/foreversixty/releases?q=companion) and put it
anywhere you like.

| Platform | Asset |
|---|---|
| macOS, Apple silicon | `foreversixty-companion_darwin_arm64` |
| macOS, Intel | `foreversixty-companion_darwin_amd64` |
| Windows | `foreversixty-companion_windows_amd64.exe` |
| Linux | `foreversixty-companion_linux_amd64` |

Every asset is published beside a `.minisig` signature. To check one:

```sh
minisign -Vm foreversixty-companion_darwin_arm64 -P "$(cat MINISIGN_PUBLIC_KEY.txt)"
```

The public key is in `MINISIGN_PUBLIC_KEY.txt` beside the assets on every release the project has
signed. The companion verifies the same signature itself before it installs an update, so the
check above is for the first download.

macOS and Windows builds are signed and notarized when the project's signing credentials are in
place. If a build is unsigned, macOS will refuse it on the first run: right-click the binary,
choose Open, and confirm once.

Linux needs WebKitGTK for the window: `sudo apt install libgtk-3-0 libwebkit2gtk-4.1-0`. Without
it, run with `-headless` and open the printed URL in a browser.

## Pair it

1. Sign in at [foreversixty.gg/logs](https://foreversixty.gg/logs) with Battle.net or an email
   link.
2. Press **Pair a device**. The site shows a code that is good for ten minutes.
3. Open the companion, go to **Device**, paste the code, press **Pair this device**.

The companion receives an upload token, stores it in your operating system's keychain (macOS
Keychain, Windows Credential Manager, the Linux Secret Service) and shows which one it used on
the Device page. Where no keychain is available it falls back to `config.json`, which is written
readable only by you. Revoke a device any time from **Account → Devices** on the site.

## Log a raid

Turn advanced combat logging on once, in **Options → Network → Advanced Combat Logging**. The
companion's status page tells you if it is off.

Then type `/combatlog` in game at the start of the night and `/combatlog` again at the end — or
leave it on. That is the whole workflow. The companion notices the file growing, uploads each
fight within a few seconds of it ending, streams the raw log up in the background so the server
can check our arithmetic, and closes the report fifteen minutes after the log goes quiet. A gap
longer than thirty minutes starts a new report, so two raids in one night are two reports.

The **Reports** page links every report the moment it has an id, live ones included.

## Settings

| Setting | What it does |
|---|---|
| New reports are | public, unlisted, private or guild — the visibility a report starts with |
| Game folders | one per line; the companion also looks in the usual install paths on its own |
| Logging character | who reports are attributed to, picked from the characters the addon has seen |

Where the companion keeps its files:

| Platform | Directory |
|---|---|
| macOS | `~/Library/Application Support/ForeverSixty` |
| Windows | `%APPDATA%\ForeverSixty` |
| Linux | `~/.config/foreversixty` |

Inside it: `config.json`, `queue/` for uploads that have not landed yet, `state/` for one file
per report, `logs/companion.log`, and `update/` for a downloaded update.

## The addon

If the ForeverSixty addon is installed, the companion reads its saved variables after you log
out and uploads each character's export string, so your real talents and gear reach the planner
without a copy and paste. It writes `ForeverSixtyInbox.lua` beside them every ten minutes with
the builds you chose on the site, which the addon loads the next time the game starts or you
`/reload`. The game rewrites its own saved variables at logout, so an inbox written mid-session
may be replaced; the next pass puts it back.

## Troubleshooting

**"This device is not paired."** Nothing uploads until it is. Fights are still parsed and queued
while you are unpaired, and they go up as soon as you pair, so a missed pairing costs nothing.

**"No World of Warcraft folder was found."** Add the flavour folder — the one containing `WTF`
and `Logs`, such as `_classic_era_` — on the Settings page. The folder above it works too.

**"Advanced logging: OFF."** The report will be accepted, marked, and will lose the views that
need the advanced fields: deaths with health, item level, resources, and positions. Turn it on
in **Options → Network → Advanced Combat Logging** and restart the game.

**Fights are queued and not going up.** Open `logs/companion.log`. Every upload failure is one
line with the report, the fight and the reason. Uploads are strictly in order, so one failing
fight holds the rest behind it deliberately; they go as soon as it does. A fight the server
refuses outright is dropped with an `ERROR` line rather than retried forever. Run the companion
with `-debug` for a more detailed log if the reason isn't clear.

**Nothing appears while the raid is going.** The engine emits a fight when the line after its
last one arrives, which in a live raid is immediate. A fight that is genuinely the last thing in
the file closes when the report does.

**The window will not open.** Run with `-headless` and open the URL it prints. The interface is
the same; it is served on loopback with a per-session token in the path.

**Start over.** Quit the companion and delete the directory above. Reports already uploaded stay
on the site.

## Build it yourself

```sh
cd companion
go test -tags nogui ./...      # the whole suite, no desktop libraries needed
go build ./cmd/foreversixty-companion
```

The window needs cgo and a platform toolkit: WebKit on macOS (Xcode command line tools),
WebView2 on Windows (a C compiler such as mingw-w64), and `libgtk-3-dev libwebkit2gtk-4.1-dev`
on Linux. `-tags nogui` builds and tests everything else without them.

The companion depends on the `logs/` module in this repository through a `replace`, so the
engine that parses your log in the companion is the same code that parses it on the server.

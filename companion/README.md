# Forever Sixty companion

A desktop app that watches World of Warcraft's combat log, uploads each fight to
[foreversixty.gg](https://foreversixty.gg) as it ends, and keeps the ForeverSixty addon in step
with the builds you chose on the site. One binary per platform, a tray icon, and a small window.
It never modifies the game's log and never sees your Battle.net password.

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

The Windows build is signed through Azure Artifact Signing in the name of Commish LLC, the
company behind Forever Sixty, so that is the publisher Windows names in its prompts. A brand-new
signing identity has no SmartScreen reputation yet, so the first installs may still see "Windows
protected your PC": choose **More info**, then **Run anyway**. That warning retires on its own as
installs accumulate. The macOS build is not signed yet: right-click the binary, choose Open, and
confirm once.

Linux needs WebKitGTK for the window: `sudo apt install libgtk-3-0 libwebkit2gtk-4.1-0`. Without
it, run with `-headless` and open the printed URL in a browser.

## Pair it

1. Open the companion. Until it is connected it opens on **Device**.
2. Press **Sign in with Battle.net**. Your browser opens on foreversixty.gg with a short code
   (for example `KQ4-7TM`); check it matches the one in the companion, sign in if you are not
   already, and press **Connect**. The companion notices by itself, says **Connected as
   Name#1234** and switches to Status. The link works for ten minutes; **Open the browser
   again** reopens the same one and **Cancel** stops waiting.
3. If your browser will not open, the companion shows the link to copy into any browser, on this
   PC or another.

Your Battle.net password never reaches the companion: you sign in on Blizzard's own page through
the site, and the companion only ever receives an upload token.

**Paste a code instead.** Under **Having trouble? Paste a code from the site** on the Device page
there is the older path: sign in at [foreversixty.gg/logs](https://foreversixty.gg/logs), press
**Pair a device**, paste the code (good for ten minutes) and press **Connect with code**.

The companion stores the token in your operating system's keychain (macOS Keychain, Windows
Credential Manager, the Linux Secret Service) and names which one it used on the Device page under
**Saved in**. Where no keychain is available it falls back to `config.json`, which is written
readable only by you, and says so. **Forget this device** (it asks first) revokes the device on
the site (it disappears from **Account → Devices**) and removes the token from this PC. If the site
cannot be reached, the token is still removed here and the window says to remove the device from
**Account → Devices** yourself. Your reports on the site stay.

## Close it and it keeps running

Closing the window does not stop the companion: it stays in the tray (Windows) and keeps logging.
The first time you close the window it says so, then hides. The tray icon's menu shows what the
companion is doing in a line at the top (the same words as the Status headline, which is also the
tooltip), **Open Forever Sixty**, **Open my reports on the site** and **Quit**. Quit is the only way
to stop it. On macOS and Linux, closing the window quits the companion; keep it open while logging. The
close-to-tray behaviour is Windows-only in the first release.
Launching the companion again just opens the running one.

## Log a raid

Turn advanced combat logging on once, in **Options → Network → Advanced Combat Logging**. The
companion's status page tells you if it is off.

Then type `/combatlog` in game at the start of the night and `/combatlog` again at the end — or
leave it on. That is the whole workflow. The companion notices the file growing, uploads each
fight within a few seconds of it ending, streams the raw log up in the background so the server
can check our arithmetic, and closes the report fifteen minutes after the log goes quiet. A gap
longer than thirty minutes starts a new report, so two raids in one night are two reports.

The **Reports** page lists the newest reports the moment they have an id, live ones included, each with a **Copy link** button, and links to all your reports on the site.

## Settings

| Setting | What it does |
|---|---|
| New reports are | public, unlisted, private or guild — the visibility a report starts with |
| Keep the addon up to date | on by default. Installs new Forever Sixty addon and data versions into every game folder here, when the game is closed (see below). CurseForge and Wago users can leave it on; they get the same files |
| Game folders | the folders it found on its own, and any you add with **Add a folder...** (the system folder dialog); only the ones you add are saved, and **Remove** takes one back out. Pick a flavour folder such as `_classic_era_` or the folder above it |
| Reports are credited to | Automatic by default: the companion reads who is logging from the combat log itself (the game marks the logging player's own units in every line), then finds that character among the ones the addon has exported, so a report is credited to the right region and ruleset with nothing to pick. If the addon has not exported the character yet, the report is credited the moment it has (log out once with the addon loaded). Pick a character from the list only if Automatic is wrong; the Status page and this setting both show who it currently reads |

The companion watches the install you played most recently; the Status page tags it Watching.

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
without a copy and paste. With them it sends what the addon left for the site at logout (a "Sim
me" request, your guild tab's roster), once; `companion.log` says in one line each what the site
did with it ("Sim request queued"). It writes `ForeverSixtyInbox.lua` beside them
(`WTF/Account/<account>/SavedVariables/ForeverSixtyInbox.lua`) every ten minutes with the builds
you chose on the site, your newest sim, your last raid and what the site last saw, which the addon
loads the next time the game starts or you `/reload`. The game reads that file only for the small
`ForeverSixtyInbox` addon that ships in the same download as ForeverSixty (both folders go into
`Interface/AddOns`); without it the inbox is never seen. The status page shows "Seen by the site
14:03" once the site has your export. The companion only ever writes `ForeverSixtyInbox.lua`;
`ForeverSixty.lua` is the game's. The game rewrites its own saved variables at logout, so an inbox written mid-session
may be replaced; the next pass puts it back.

### Keeping the addon up to date

At launch and then every six hours the companion reads the repository's
releases for a newer `addon-v*` release (the `ForeverSixty` and `ForeverSixtyInbox` folders) and
the newest `data-v*` release (the nightly `ForeverSixtyData` folder), and installs them into every
game folder that already has the addon. Both zips are signed with the same minisign key as the
companion's own releases, and the companion installs nothing it cannot verify against the public
key built into it: a release without a `.minisig` beside its zip is not a candidate, and a bad
signature is logged and skipped. The installed version is the `## Version` line of
`Interface/AddOns/<folder>/<folder>.toc`.

The companion never writes into `Interface/AddOns` while that install's game is running. It asks
the system for the running programs (on Windows the process list is the truth) and treats a
`Wow*.exe` whose path is under the install, or whose path it cannot read, as the game. A release
that arrives while the game is up waits in `update/addons/` and goes in a minute after the game
closes. Each folder is swapped by rename: the new one is copied beside the old as
`<folder>.fs-new`, the old becomes `<folder>.fs-previous`, the new takes its place, the new TOC
must read back with the version just installed, and only then is the previous folder deleted; on
any failure the previous folder is put back and `companion.log` says so. Only `ForeverSixty`,
`ForeverSixtyInbox` and `ForeverSixtyData` (and those two suffixes) are ever written or removed; a
folder that is a link or holds a `.git` is a developer's and is left alone, and saved variables
are never touched.

An install with no `ForeverSixty` folder gets nothing automatically. The Addon row on Status says
"Addon not found in Classic Era" and offers **Install the addon into Classic Era**, which installs
the addon and the data addon from the signed releases (it refuses while that game is running).

## Troubleshooting

**"Not connected."** Nothing uploads until you sign in on the Device page. Fights are still parsed and queued
while you are not connected, and they go up as soon as you connect, so a missed sign-in costs nothing.

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
the same; it is served on loopback with a per-session token in the path. The Windows build has
no console of its own — it is a windowed application, so starting it never opens a `cmd` box —
so redirect what it prints to a file and read it from there:

```
foreversixty-companion.exe -headless > url.txt
foreversixty-companion.exe -version > version.txt
```

The URL is not written to `companion.log`: its token is the whole authentication for the local
interface, and the log is a file you may well paste somewhere.

**Start over.** Quit the companion and delete the directory above. Reports already uploaded stay
on the site.

## Icons

The Forever Sixty seal (`web/public/favicon.svg`) is the app's icon everywhere the OS shows one:
the window and taskbar, the tray, Explorer (the Windows exe carries it as a resource, with the
product and company fields), and the dock on macOS. The rendered files live in
`internal/icon/assets/` and are committed; to regenerate them after the seal changes:

```sh
node tools/icons/render.mjs   # needs `pnpm install` in web/ for resvg; WEB_NODE_MODULES=... if it lives elsewhere
```

The Windows resource is built from `cmd/foreversixty-companion/winres/winres.json` by the release
workflow (go-winres), not committed. The macOS release also publishes
`foreversixty-companion_darwin_<arch>.app.zip`: unzip it and put **Forever Sixty.app** in
Applications. The updater replaces the binary inside the bundle, so the app keeps its icon across
updates; the bare binary stays published beside it because that is what the updater downloads.

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

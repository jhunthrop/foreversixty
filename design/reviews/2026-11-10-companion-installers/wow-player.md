# Round 1

# wow-player review: companion install experience

**Verdict: FIX.** The Mac window and the Windows installer both read well. I would trust either in one glance. Two things stop me from saying SHIP: the /setup capture is a build error page rather than the page, and the first-run screen has no way to confirm the install worked.

## What a player sees

**Mac move window** (mac-1-move-window.png). The question "Move Forever Sixty to Applications?" is the first thing on screen. It has one toggle and one button, and the crest and parchment frame match the companion. I would click it without thinking. "Open at login" is on by default, which is what an uploader wants. It is a single honest toggle, so it is not naggy.

**Windows setup** (win-1-setup.png). It is a single page with two ticks and Install. "The companion for your Windows account. No admin needed." is the right reassurance for a veteran who has been burned by installers. The crest is in the corner and nothing is bundled. "Start with Windows" is on and "Desktop shortcut" is off, which is the correct split.

**First run** (win-2-first-run.png). "Not connected" with one Battle.net button is clear. "Opens your browser. Your Battle.net password never reaches this app." is a good trust line.

## Numbered fixes

1. **/setup was not rendered, so I cannot approve it.** setup-windows-1440.png is an Astro "FailedToLoadModuleSSR: Could not import ../../data/generated/combos.json" error page. This is the generated data missing in the worktree, not a copy problem. Re-capture `web/src/pages/setup.astro` for Windows at 1440 after the data is generated, and re-run it. setup-mac-1440.png did render, and that is the only /setup surface I can judge.

2. **The Mac line under the button is a pitch and reads as a sequence.** File: `web/src/data/companion.json`, mac `firstRun`. It currently says "Signed by Commish LLC and notarized by Apple. Open the download and Forever Sixty moves itself to Applications, then sets itself up. macOS 12 or later." On screen it sits in grey under the line "The companion logs live from your desktop." Change it to: "Signed by Commish LLC and notarized by Apple. Open the download and it moves itself to Applications. macOS 12 or later." Players do not need "then sets itself up", because step 1 already covers the sign-in. Also drop "The companion logs live from your desktop." It tells a veteran nothing the header does not.

3. **The Windows line promises something the screen does not show.** File: `web/src/data/companion.json`, windows `firstRun`. It says "Open the download and Forever Sixty sets itself up" and then explains SmartScreen. The player actually meets a Setup window with two ticks and an Install button. Change it to: "Signed by Commish LLC. Run the download, leave the ticks as you like, press Install. Your Windows account only, no admin. If SmartScreen asks, choose More info, then Run anyway." Keep the SmartScreen line, because that is the thing that makes people quit.

4. **The first-run lede still says "up to date" in two senses.** File: `companion/ui/copy.js`, `install.lede`: "It runs from there and keeps itself up to date. The download goes in the Trash." The first half is fine. "The download goes in the Trash." is the line a player will stop on. A veteran asks whether the zip gets deleted. Change it to: "It runs from there and keeps itself up to date. The downloaded copy goes to the Trash." Using "copy" makes clear that nothing else is touched.

5. **The Windows title bar and header repeat the version without a reason.** File: `companion/packaging/windows/companion.iss`. The title "Setup - Forever Sixty 0.2.45" is Inno's default and is fine. The header "Install Forever Sixty" is also fine. However, the empty lower half of the window, roughly 60 percent of its height, looks unfinished. Either shrink the wizard height in `[Setup]` (`WizardSizePercent` or `WizardStyle=modern` with a smaller size) or add one line of what Install does, for example "Installs to your user folder and opens the companion." Choose the shrink and do not add copy.

6. **There is no confirmation that the move or install happened.** File: `companion/ui/copy.js`. After Move or Install the player lands on "Not connected", which is the same screen as a first launch from anywhere. A player cannot tell it is now running from Applications or its install folder. This is a minor gap. Add a one-line toast or a line in the Device footer such as "Installed. Opens at login." (`install.done`), shown once. Otherwise it is fine to ship without.

7. **The failure line is a dead end.** File: `companion/ui/copy.js`, `install.failed`: "That did not work. Close this window and open Forever Sixty again to try once more." "That did not work" is vague. Say what to check: "Could not move it. Drag Forever Sixty into Applications yourself, then open it from there." A veteran can always drag an app. The retry advice may fail repeatedly.

8. **The uninstall question is good but one clause is odd.** File: `companion/packaging/windows/companion.iss`, `[Code]` message. "Also delete the companion's settings, logs and any fights still waiting to upload? Your reports on foreversixty.gg stay either way." Defaulting to No (`MB_DEFBUTTON2`) is correct. "logs" may be read as WoW combat logs. Change "logs" to "its own logs", or drop it, so no player fears it deletes their WoWCombatLog files.

## Missing

- No way to see where it installs on Windows before pressing Install, since the one-screen setup hides the folder. That is acceptable for a per-user install, but "Your Windows account only" in the web line should say where it lives only if a player asks. I would not add it.
- The Mac toggle gives no hint that it can be changed later. If Settings has "Open at login", the lede could say so, but it is not needed.

## Data and claims I do not believe

- "notarized by Apple" and "No admin needed" are claims about the build, not the screens. I cannot verify them from screenshots, so whoever checks accuracy should confirm the notarization ticket and the `PrivilegesRequired=lowest` setting.
- "macOS 12 or later" should match the build's `LSMinimumSystemVersion`.

## Summary

Items 1 to 3 are needed before shipping. Items 4 and 7 are copy fixes worth making now. Items 5, 6 and 8 are small and can follow. No em dashes or slop phrasing appeared in the copy I read. The tone matches how a veteran would write it.

# Round 2

# Round 2 verdict: SHIP WITH FIXES

All round-1 blockers are closed. I re-viewed mac-1-move-window.png, mac-state-failed.png, win-1-setup.png and setup-windows-1440.png. I did not re-open the Mac /setup captures or the 360 and 2000 versions, so the Mac /setup line is judged on the copy you quoted, not on a render.

## Round-1 items

1. **/setup not rendered: closed.** setup-windows-1440.png is now the real page. The companion card reads "Signed by Commish LLC. Open the download and press Install. It installs for your account only, no admin needed. If SmartScreen asks: More info, then Run anyway." That matches what the player sees, and the SmartScreen line is still there.
2. **Mac pitch line: closed.** "Notarized by Apple. Open the download and Forever Sixty moves itself to Applications. macOS 12 or later." is plain and true. Dropping "Signed by Commish LLC" on the Mac is fine, since notarization is the thing Gatekeeper shows.
3. **Windows line promises vs. screen: closed.** "Press Install" matches the one button.
4. **Lede: closed.** "The download goes in the Trash." is short and honest.
5. **Empty window: closed.** The Setup window is visibly tighter with 110px trimmed, and "Installs for your Windows account." is enough.
6. **No done toast: accepted.** The first run follows at once, so the player sees it worked. I agree.
7. **Failure line: accepted.** "Move failed. Reopen Forever Sixty to try again." reads fine in the red strip, and the button stays live. I accept the owner's no-drag rule.
8. **Uninstall "its own logs": closed.** This is not a screenshot item. I take it from your message, and it removes the fear that WoWCombatLog files get deleted.

## Remaining fixes

1. **`web/src/pages/setup.astro` (the /setup header, step 3 and companion card): three "Download for Windows" buttons plus two "Other platforms" links on one screen.** A player scrolling the Windows /setup at 1440 sees the same button in the header, in step 3 and in the card. Keep the header one and the card one and make step 3 a plain "Install the companion below" row. This is minor and was there before this lane.
2. **`web/src/data/companion.json`, both `firstRun` lines: low contrast.** The grey text under the button on the dark card is hard to read on the Windows capture, and that is the line with the SmartScreen instruction. Raise it to the body text colour, or at least the same grey as the step text, so the "More info, then Run anyway" part is readable. This is the line players need when SmartScreen blocks them.
3. **"The companion logs live from your desktop." still sits between that line and step 1.** I accept that it belongs to another lane (addonCopy), but it still tells a veteran nothing. Log it as a follow-up for that lane.

## Not verified

"Notarized by Apple", "no admin needed" and "macOS 12 or later" remain build claims for the accuracy reviewer, as in round 1.

Fix 2 is worth doing before release. Fixes 1 and 3 can follow.

# After round 2

Applied fix 2: the first-run line is in body ink, not muted. Fixes 1 and 3 belong to the /setup page and addonCopy lanes and are left as follow-ups.

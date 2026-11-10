# Round 1

# AI-slop review: companion installers

## 1. First impression

These read as made by a person who plays WoW. The Mac window and the Windows first run use the Warcraft-style framing: the LX seal, the parchment backdrop and the ornate-bordered panel. The Windows setup is Inno's native wizard with short copy. Nothing here looks generated. The worst tells are three small copy items, and one of them is on a surface that was meant to be fixed.

## 2. Tells found, most damaging first

**Tell 22 (explaining the obvious or adding a qualifier), `web/src/data/companion.json`, both `firstRun` strings, shown on `/setup` (setup-mac-1440.png).**

The grey line is a four-clause pitch: "Signed by Commish LLC and notarized by Apple. Open the download and Forever Sixty moves itself to Applications, then sets itself up. macOS 12 or later." The Windows string is the same, plus "for your Windows account only, with no admin needed."

- It narrates the install mechanics, which a veteran player would do without being told. "Open the download" and "then sets itself up" explain the obvious.
- "Signed by Commish LLC" is trust-signalling copy, and it sits in the same grey run as the SmartScreen instruction.
- Change the Mac string to: `Notarized by Apple. macOS 12 or later.`
- Change the Windows string to: `Signed by Commish LLC. SmartScreen: More info, then Run anyway.`
- Drop "for your Windows account only, with no admin needed" here. The setup screen already says it.

**Tell 22 and tell 7 (an explainer sentence under a heading), `companion/ui/copy.js`, `install.lede`.**

"It runs from there and keeps itself up to date. The download goes in the Trash." This is a grey explainer under the heading.

- "It runs from there" restates the title.
- "Keeps itself up to date" is a benefit pitch that has nothing to do with the question being asked.
- Keep only the consequence the user needs: `The download goes in the Trash.`

**Tell 22, `companion/packaging/windows/companion.iss`, `[Messages]` `SelectTasksDesc` and `WizardSelectTasks`.**

The subtitle reads "The companion for your Windows account. No admin needed."

- The first half is a product description on an installer that is already open.
- "No admin needed" is reassurance nobody asked for, because the wizard shows no UAC prompt.
- Delete `SelectTasksDesc`, or shorten it to `Installs for your Windows account.`
- Title and subtitle currently say the same thing twice. "Install Forever Sixty" sits directly above a line about the companion.

**Tell 22 (minor), `companion/ui/copy.js`, `install.failed`.**

"That did not work. Close this window and open Forever Sixty again to try once more." This is padded.

- Change it to `Move failed. Reopen Forever Sixty to try again.`

**Tell 7 (minor), `companion/ui/install.html`, and the overall composition.**

The Mac prompt is a title, a grey sentence, a toggle and a button inside a double-bordered panel. It is sparse, but it is a real, single-purpose dialog. The panel has a border inside a border, a frame-in-frame look. Check that this is the house panel style from the main companion window and not a one-off. It is not a card stamp, so it is not tell 8, provided the main window uses the same panel.

## 3. What already reads as authentic

- **Win-1 setup:** the buttons are the OS-native "Install" and "Cancel", there is no marketing text, and the checkboxes are "Start with Windows" and "Desktop shortcut" (unchecked by default). The corner seal is the real crest. This is the right amount of restraint.
- **Mac window:** the Marcellus title, the Open Sans body, the real LX crest and the parchment backdrop. They match the app and the game, and there is no gradient, glow or glass.
- **Toggle and button:** both are flat teal. There is no glow on the button, and the toggle is a normal toggle.
- **Title:** "Move Forever Sixty to Applications?" is a plain question in a human voice.
- **Win-2:** the existing first run, "Not connected / Sign in with Battle.net", is concrete and its copy is short. The yellow "Having trouble? Paste a code from the site" link is a bit loud but is not a tell. I saw no glaring issues.
- **/setup:** the numbered steps are a real sequence. They use real menu paths (`Options > Network > Advanced Combat Logging`) and a real command (`/combatlog`).

## 4. Verdict: PASS WITH FIXES

None of the visual tells that block a pass (1, 2, 5, 8, 12, 13, 15) are present. Fix the copy items before shipping, in this order:

1. `web/src/data/companion.json`: cut both `firstRun` strings as in item 1 above.
2. `companion/ui/copy.js`: cut `install.lede` to `The download goes in the Trash.` and shorten `install.failed`.
3. `companion/packaging/windows/companion.iss`: remove or shorten `SelectTasksDesc`.

Files reviewed, all under `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296`:
- `design/reviews/2026-11-10-companion-installers/` (the captures)
- `companion/ui/install.html`
- `companion/ui/install.js`
- `companion/ui/copy.js`
- `companion/packaging/windows/companion.iss`
- `web/src/data/companion.json`


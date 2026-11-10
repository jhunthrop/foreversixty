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

# Round 2

# Round 2 re-review: companion installers

## Verdict: PASS WITH FIXES

None of the blocking visual tells (1, 2, 5, 8, 12, 13, 15) are present. The remaining fixes are copy only.

## What cleared since round 1

- **Mac window** (`companion/ui/copy.js`, `install.lede`): "The download goes in the Trash." is one concrete fact. The pitch is gone.
- **Mac window** (`install.failed`): "Move failed. Reopen Forever Sixty to try again." is short and plain.
- **Windows setup** (`companion.iss`, `SelectTasksDesc`): "Installs for your Windows account." is fine. It states the per-user install honestly, and "No admin needed" no longer appears in the wizard.
- **Double-border panel**: accepted as the main window's house `.panel.lead` style, so it is not tell 8.
- **Mac button focus**: the yellow focus ring on "Move to Applications" is a normal keyboard focus ring, not a glow. Keep it flat.

## Remaining items

1. **Tell 22, `web/src/data/companion.json`, Windows `firstRun`.**
   - Current text: "Signed by Commish LLC. Open the download and press Install. It installs for your account only, no admin needed. If SmartScreen asks: More info, then Run anyway."
   - "Open the download and press Install" explains the obvious. The Install button is the only button on the setup screen.
   - The account-only and no-admin line stays, because the owner asked for it. It is the one clause that changes what the user expects.
   - Change it to: `Signed by Commish LLC. Installs for your account only, no admin needed. If SmartScreen asks: More info, then Run anyway.`

2. **Tell 22, `web/src/data/companion.json`, Mac `firstRun`.**
   - Current text: "Notarized by Apple. Open the download and Forever Sixty moves itself to Applications. macOS 12 or later."
   - "Open the download" is the same filler. The owner wanted "open the download, it sets itself up", so keep one sentence for it.
   - Change it to: `Notarized by Apple. Open the download and it moves itself to Applications. macOS 12 or later.`
   - Keep it only if the owner insists on the verb. Otherwise drop the middle sentence.

3. **Tell 7, `/setup`.**
   - In setup-windows-1440.png the companion section has two grey lines under the heading: the "Optional. Sends bags, bank and live logs…" caption and the `firstRun` text.
   - The same "Optional. Sends bags, bank…" sentence also appears in the top steps list, so it is repeated within the page.
   - This was not changed in this lane, so I am not blocking on it. Drop the heading caption on the card, since the steps row at the top already says it.

## Keep

- The Inno wizard is native: one "Install" button, "Start with Windows" checked, "Desktop shortcut" unchecked, and the real crest in the corner.
- The Mac window uses Marcellus and Open Sans, the LX seal, flat teal controls and no glow.
- `/setup`:
  - Its numbered steps are a real sequence.
  - They use real menu paths and `/combatlog`.
  - The "Example" tag marks the screenshot honestly.

I did not check the mac-state-*.png captures or the 360 and 2000 width captures. Only the Windows 1440 `/setup` capture was reviewed.

Files, all under `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296`:
- `web/src/data/companion.json`
- `companion/ui/copy.js`
- `companion/packaging/windows/companion.iss`

# After round 2

Applied: Mac line "Notarized by Apple. Open the download and it moves itself to Applications. macOS 12 or later." Windows keeps "Open the download and press Install" (the owner's wording) and drops "It installs": "For your account only, no admin needed." Item 3 (the heading caption) belongs to the /setup lane.

# Round 3 (DMG window)

**Verdict: PASS WITH FIXES.** None of the fixes is a blocking visual tell (1, 2, 5, 8, 12, 13, 15). I judged this from the three captures and `tools/art/dmg_background.py`. I did not open `companion/tools/icons/render.py` or `tools/package-mac-companion.sh`.

**First impression:** It reads as made on purpose. It mirrors Discord's layout (small wordmark top left, icon, arrow, Applications) but in our own materials, and it has no glow, gradient or card chrome.

**Tells found, most damaging first**

1. **Tell 4, cream ground, serif display face, warm accent (mild).** The ground is cream, the wordmark is Marcellus and the arrow is brass. This is the combination I flag by default.
   - It is defensible here. `dmg_background.py` documents why the ground must be light: Finder draws icon labels in dark ink in both appearances. The cream, ink and brass are all `design/tokens/tokens.json` values. The wordmark is 19pt, not a hero slogan, so tell 5 does not apply.
   - Change: none required. Keep the docstring's reasoning. Do not add any further serif or gold elsewhere in the window.

2. **Tell 21/22, the name is stated three times.** "Forever Sixty" appears in the title bar, in the Marcellus wordmark and in the icon label. The seal also appears twice, in the wordmark and in the app icon, which is itself the seal.
   - Change in `tools/art/dmg_background.py`: drop the `seal` paste (and `SEAL_PT` / `NAME_GAP`) so the top-left is the name alone, with `MARK_AT` unchanged.
   - Alternatively keep the seal because the owner's Discord reference does the same. In that case say so in the docstring.
   - The title-bar and label repeats come from Finder, so leave them.

3. **Arrow proportion.** The shaft runs about 172pt between 128pt icons with a 3.5pt stroke. That reads as a long thin UI-library arrow. Discord's is short and close to the icons.
   - Change in `dmg_background.py`: raise `ARROW_AIR` from 34 to about 52 and lower `ARROW_HEAD` from 13 to 11. This gives a shorter arrow with the same weight.

4. **Tell 15, watch only.** The app icon is the owner's painted teal and blue medallion with a lit rim. It reads somewhat airbrushed, but it is the owner's own brand art, so it is exempt.
   - Change: none. Just make sure `mac_seal` in `companion/tools/icons/render.py` does not add its own gloss, highlight or drop shadow on top of the art. The captures show only the rim Finder draws, which is correct.

**What reads as authentic (keep)**
- A single flat brass arrow with no glow or gradient, used as the only accent.
- The light ground, chosen for a documented Finder legibility reason (checked on macOS 26), not for style.
- A wordmark that is small and sits top left, so the icon and arrow carry the window. The hierarchy is uneven on purpose.
- Real tokens, a real brand typeface and the real seal asset. The build is deterministic.
- The window holds nothing but the install action: no tagline, no instruction text, no badges.
- Both appearances hold up. The cream ground works under the dark and the light title bar.

# Round 4 (both platforms, the Windows setup in our own window)

# Companion installer review: FAIL

**First impression:** The copy is player-written and specific, and the Mac move prompt and the Windows setup are the same window. The visuals are the problem. Cream, serif and brass on the DMG is the template, and the painted medallion, blotchy parchment and teal-fade header band read as generated.

I did not open the full `design/reviews/.../page-*.png` set or the dark and light DMG variants beyond `mac-dmg-window-dark.png`. I read the Mac DMG, the move window, the Windows setup, installing and uninstall states, the Discord reference, `/setup` at 1440, and the copy and style source.

## Tells, most damaging first

1. **Tell 15, the LX medallion.** It is the app icon on `mac-dmg-window-dark.png`, the Mac tile, and the header seal in `win-1-setup.png` and `mac-1-move-window.png`. Teal marble, over-lit bronze and a glossy rim read as an airbrushed generated illustration.
   - Fix: use the faithful flat HD crest per the owner's icon rule, or a hand-drawn mark with a flat fill and no lighting.
   - This is a brand-wide asset, so the owner may decide it separately, but it cannot pass while it is the installer's hero.

2. **Tell 4, DMG background.** In `tools/art/dmg_background.py` and `mac-dmg-window-dark.png`, the cream ground, the Marcellus serif "Forever Sixty" wordmark and the brass arrow are the cream, serif and warm-accent combination.
   - Fix: copy the reference. Use a plain white or neutral light ground, drop the wordmark or set it as the app's own title, and use a short chevron in the one teal accent instead of a long brass arrow.
   - The cream was chosen so Finder's dark labels stay readable. Plain white solves that equally well.

3. **Tell 12, window ground.** The blotchy painted sand texture sits behind the prompt in every Windows and Mac capture (`win-1-setup.png`, `mac-1-move-window.png`, `page-setup-installing.png`). It reads as noise and a generated texture.
   - Fix: a flat `--color-bg` fill for `.prompt`, with no texture.

4. **Tell 2, header band.** The teal band at the top of the Windows window (`win-1-setup.png`) and the Mac window (`mac-1-move-window.png`) is a dark-to-lighter teal fade with no information in it. It also repeats the title bar: "Forever Sixty" appears in the title bar and again in the band.
   - Fix: delete the band in the prompt windows and let the h1 carry the identity.

5. **Tell 8, card inside a window.** The dark rounded card has an outer border plus an inset ring (`page-setup-installing.png`) and holds one heading, one line, one toggle and one button. It sits inside a window that is already a container.
   - Fix: in `companion/ui/styles.css`, drop the card background, border and inset ring from `.prompt`. Set the text straight on the window ground, left-aligned with the window padding.

6. **Tell 13, blank window.** `win-2-installing.png` is an empty brown rectangle with no header, no card and no bar. If that is what a player sees between launch and render, it is an empty state posing as a window.
   - Fix: paint the window background and the h1 in the first frame, or show the window only once the content is ready. If it is only the capture timing, retake it.

7. **Tell 22, qualifiers.** `copy.js` `install.setup.lede` is "For your Windows account. No admin needed." The companion.json Windows `firstRun` says the same thing again: "For your account only, no admin needed." A veteran knows it is per-user from the lack of a UAC prompt.
   - Fix: `lede` becomes empty on Windows, since the Mac prompt gets by with one useful line. Keep "Replaces the copy already installed. Your settings stay." for upgrades.

8. **Tell 21, explaining the obvious.** In `web/src/data/companion.json` the Windows `firstRun` has "Open the download and press Install."
   - Fix: "Signed by Commish LLC. SmartScreen: More info, Run anyway." The Mac line can lose "it moves itself to Applications", because the prompt says so.

9. **Tell 21, heading repeats the button.** "Install Forever Sixty?" over "Install", "Remove Forever Sixty?" over "Remove", and "Move to Applications?" over "Move to Applications" say the same thing twice.
   - Fix: make the window's single line the useful fact, for example "Installs to %LocalAppData%" or the version number. Keep the heading only if it states something the button does not.

10. **Minor inconsistency, accent colour.** In `.progress-track span`, the bar is cyan (`--color-accent`). The button and toggle are teal-green. Two accents appear in one 360px column (`page-setup-installing.png`).
    - Fix: use the button's fill colour for the bar.

11. **Tell 7 and duplication on `/setup`.** `setup-windows-1440.png` shows the line "Optional. Sends bags, bank and live logs after every session." twice, once in the step row and again beside "The companion". Grey sub-copy sits under nearly every heading on the page.
    - Fix: say it once.

## What already reads as authentic (keep)
- "The download goes in the Trash." and "Your reports on foreversixty.gg stay." are concrete lines a developer writes.
- Every state is named: upgrade, uninstall, "Keep my settings" and a `Start with Windows` toggle that defaults on. Each has its own words, including `install.steps.*` such as "Adding it to Apps & Features".
- The signed and notarized lines in `companion.json` name Commish LLC and Apple instead of saying "secure".
- The uninstall button is red text on dark. That is by role, not a glowing primary.
- The DMG layout follows the owner's Discord reference, and the icon-centring review rounds are visible.
- There are no em dashes, glows, fake stats or emoji in the copy.

## Verdict: FAIL
Tells 4, 12, 13, 15 and 8 are present, and 13, 15, 12 and 8 each block a pass on their own. After fixes 1 to 6, items 7 to 11 are copy trims that would otherwise be PASS WITH FIXES. The paths involved are `companion/ui/styles.css`, `companion/ui/copy.js`, `web/src/data/companion.json`, `tools/art/dmg_background.py`, and the captures in `/Users/jh/code/forever/.claude/worktrees/agent-a401ca0c4adf37296/design/reviews/2026-11-10-companion-installers/`.

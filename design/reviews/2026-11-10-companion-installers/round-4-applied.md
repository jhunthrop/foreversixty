# Round 4: what was applied, and what was not

Applied (ux-designer):
1. Progress is held for at least 1.2 s. The bar shows the steps already done and fills only at `fsDone`.
2. The title turns to "Installing…", "Updating…" or "Removing…" while the steps run, and the lede is hidden.
3. An update keeps the player's Start with Windows choice (`WinSystem.StartupEnabled`, tested).
4. An update reads "Update Forever Sixty?" with the button "Update".
5. Failure lines name the button on screen.
6. The uninstall focuses the toggle, not the red Remove.
10. The step lines are "Copying files" and "Removing files".
11. The done state keeps the lede.
12. The setup and uninstall windows lose minimize and maximize (`dialogChrome`).

Applied (wow-player):
- The Mac lede reads "This download goes in the Trash."
- The Windows /setup line opens with "Verified publisher: Commish LLC", SmartScreen's own words.

Applied (ai-slop):
- The progress bar uses the button's teal, not the cyan accent.

Not applied, with reasons:
- **ai-slop 1, 3, 4, 5 (the painted seal, the painted ground, the band, the lead panel).** These are the owner's round-5 brand: one logo, the painted art, and the house panel. The companion window and the site share them, and the earlier rounds passed this same window. Changing them belongs to the owner, not to this lane.
- **ai-slop 2 and ux 9 (the DMG's cream, serif and brass; a teal band on the DMG).** The cream stays because Finder's labels are dark in both appearances. The owner's spec asks for the seal and wordmark top left; a band was left out to keep the window plain, as the reference is.
- **ux 8 ("drag" on /setup).** The owner's later ruling: the Mac line is "open it, then open Forever Sixty, which offers to move itself".
- **wow-player 2 and 3 (the hand-install path, the Advanced Combat Logging menu path).** These are /setup copy owned by another lane (`addonCopy`); logged as follow-ups.
- **wow-player 5 ("change this later in Settings").** The companion has no such setting yet.
- **ai-slop 13 and ux 1 (win-2 blank).** The capture caught the companion's own window opening after a 0.5 s install. The setup now holds its progress for 1.2 s, and the capture is retaken mid-run.

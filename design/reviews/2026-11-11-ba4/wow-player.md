# wow-player review: desktop companion, 2026-11-11

Question the surface answers: "Is my log reaching the site, and is anything stopping it?" The Status page answers it first (Logging / Ready / Advanced logging is off, big title, coloured dot). Good.

## Requested changes
- "Open your reports" is identical in the Reports header, the tray menu (Windows and macOS shots) and tray tooltips. Good.
- "Turn it on in Options > System > Network > Advanced Combat Logging, then relog." is right: that is where the client puts the checkbox, and "relog" is what a player says. Good.

## Verdict: SHIP (no MUST). Numbered SHOULD list follows.

MUST
- none

SHOULD
1. "Reports will miss deaths with health, item level, resources and positions." reads wrong ("deaths with health"). Say what is lost: "Reports lose health, resources and positions at every death, and item levels." Players know what a death recap is.
2. companion.headline.advancedRowHelp ("Turn it on in Options > System > Network > Advanced Combat Logging.") has no "then relog"; the headline version does. Same sentence, one ending, from one string.
3. "not credited yet" / "credited to {name}" / "Reports are credited to" is site jargon, not game words. A raider reads "Blackwing Lair ... not credited yet" and does not know whose credit. Prefer "not linked to a character yet" and "Reports are logged as". The help line already says it needs a logout.
4. Two ways to say the same thing: "Log out once in game with the addon loaded" (credit help) vs "Log out to the character screen and back in" (builds waiting) vs "relog" (advanced logging). Pick "relog" or the long form everywhere; the long form on a veteran's screen is explaining the obvious.
5. Explains the obvious: "The first fight uploads when it ends." (the headline already says it uploads fights as they end), "Only you." / "Members of your guild." help lines (the option names say it), "CurseForge and Wago users can leave this on; they get the same files." (cut after "leave this on"), "A code works once and expires after 10 minutes." Cut or shorten.
6. "No character yet · log out once in game and the addon's save arrives here." puts a middle dot in the middle of a sentence and says "save" for SavedVariables. Use a full stop: "No character yet. Relog once and the addon's data arrives here."
7. companion.footer.updateReady: "Update ready. Quit the companion from the tray icon and open it again to install it." Every addon manager a Classic player has used (CurseForge, WowUp) just updates. If this must stay manual, shorten to "Update ready. Quit from the tray and reopen." If it can update itself, do that.
8. DMG window (dmg-window.png): the Applications target is a plain blue rounded square with no folder glyph and the lower 40% of the window is empty. If that is how the real DMG renders, it looks unfinished next to the rest of the polish; give the target the stock folder icon and trim the window height.

## Checks that passed
- No em dashes anywhere in companion.json or the screenshots. Middle dots are used as field separators only (except item 6).
- Game words are right: relog, addon, fights, report, raid names (Molten Core, Blackwing Lair, Onyxia's Lair, Ruins of Ahn'Qiraj, Zul'Gurub, Naxxramas, Scholomance), /combatlog, "Close World of Warcraft first", "WoW Forever", _forever_ folder, Classic flavor names.
- Empty/ready states give a game action, not a pitch: "Turn /combatlog on in game and your raids land here."
- No marketing hero, slogan or feature list anywhere in the window, tray or installers.
- Install steps and failures say what happened and what to press. "Start with Windows" and "Open at login" use each OS's own words.
- Counts carry units ("3 fights uploaded", "2 characters, 4 min ago"); Addon 2.0.16 / Data 2026.11.09 show versions with state.
- Tray: "Logging" status line, Open Forever Sixty, Open your reports, Quit. Matches what a tray app does.

## Data I do not question
Setting path, /combatlog, and the "relog" requirement match the client. Nothing here claims a game stat.

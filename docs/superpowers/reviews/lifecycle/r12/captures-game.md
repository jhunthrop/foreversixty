# In-game captures, 22 Oct (last beta day), addon 2.0.16 on Bow Jackzon (level 24 High Order Skyborne hunter, Darnassus)

| File | What it shows |
|---|---|
| game-00-welcome.jpg | The first-login welcome, three lines in General chat |
| game-00-help.jpg | `/fs help`, the last lines (…/fs options, /fs diag, /fs help, the two "For testers" lines) |
| game-00-unit-tooltip.jpg | Hovering another player: "Forever Score: no logs yet" and "Target them, then /fs link: site page" under the class line |
| game-00-chip-next.jpg | The Forever Sixty line beside the minimap after A.1: "Next: Endurance Training (1 of 4)" with the class icon |
| game-00-build-waiting.jpg | `/fs` with a build delivered from the site: "Builds · 1 waiting", the row "Beast Mastery build · Hunter · Beast Mastery 15 · 15 points · For Bow Jackzon · Classic Beta PvP · from the site", "Enter to preview" |
| game-01-build-preview.jpg | Enter on the row: the preview with "Gear in this build · at 60 · 14 of 17 slots", three gear rows then empty space, "Replaces Beast Mastery build · first differs at point 6: Focused Fire (Beast Mastery)", the gold "Use this build", "Enter to use" |
| game-01-build-preview-full.jpg | The same at screen scale (0.5) |
| game-02-chip-build-complete.jpg | After "Use this build": the line reads "✓ Build complete" (the sent build equals the character's 15 spent points) |
| game-03-chat-now-following.jpg | Chat: "Forever Sixty: Now following Beast Mastery build." between "[SERVER] Shutdown in 6:00" and trade chat |
| game-04-talent-window.jpg | The talent window (N) with the addon's marks on the Beast Mastery tree and its strip at the bottom |
| game-05-talent-strip.jpg | The strip: "Beast Mastery build · 15 of 15 points · complete · 0 off plan" and a greyed "Apply Changes" |
| game-06-copy-my-character.jpg | The Copy my character tab: "Selected. Press Ctrl+C, then paste it on foreversixty.gg.", the FS1 export in the box, "Seen by the site 02:21 pm", "Tab to switch" |
| game-07-chat-ready-loot.jpg | `/fs ready`: "Ready check (only you see this): nothing missing at your level."; `/fs loot` with no link: "Type /fs loot and an item link or id to see what it is worth to you." |
| game-08-options.jpg | Options → AddOns → Forever Sixty: nine switches with one-line explanations, "Reset positions", and a Commands list cut off by the panel's bottom edge |

Observed, for the reviews: the preview says "at 60" for a level-24 build; the preview's gear list stops after three rows with the rest of the panel empty; the Options page's Commands list is clipped; the "Build complete" line names no next step (the next point comes at level 25); "Seen by the site 02:21 pm" is the local clock while the game shows 03:56.

## After A.2 (four commits, merged `44fc76a5`), verified at a normal login

| File | What it shows |
|---|---|
| game-09-chip-build-complete-next-level.jpg | The line beside the minimap: "Build complete · next point at level 25" with the class icon, placed under the buff row |
| game-10-options-commands-first.jpg | Options: Commands first (/fs, /fs copy, /fs build, /fs chip, "/fs help: all the commands"), then the switches at body size, the page scrolling |
| game-11-help-no-tester-lines.jpg | `/fs help` without the tester lines; "/fs diag: check what is wrong" |

Also seen: `/fs diag` "Site: has your save from 1 h 20 min ago (6 characters)"; the talent strip "Beast Mastery build · 15 of 15 points · complete · next point at level 25". Three A.2 rounds were needed for the line: the buff-frame anchor, then `SetWidth(0)` on the font string, then the finished-plan idle state that A.1 had left as the icon alone.

## Solo-viewable surfaces on the owner's alt (Obnoxious Yell, level 21 Arms warrior, OLYMPUS XXVII), addon 2.0.17 files

| File | What it shows |
|---|---|
| game-15-chip-warrior-equip.jpg | The line for the warrior: "Equip Earthfury Gauntlets, +3.5 DPS" with the item icon |
| game-13-guild-window-tab-button.jpg | The client's Guild & Communities window with the addon's "Forever Sixty" tab button under its bottom-right corner (hover: "Your guild's Forever Scores, gear and live builds. Click to open or close.") |
| game-12-guild-tab.jpg, game-12-guild-tab-full.jpg | The addon's guild window: "Forever Sixty · Guild", "Online only" ticked, "14 of 773 shown · 1 live of 14 online · as of 22 Oct", columns Name · Level · Spec · Gear · Score↓ · Attendance · Gap; every row "— · no logs · —" except the player's own ("Arms", gap "0.7 DPS"); two level-7 rows (Carla Delianco, Fear Dapriestt) with no cells at all |
| game-14-social-raid-tab.jpg | The Social window's Raid tab solo: the client's raid text and the addon's "Forever Sixty" button under the bottom-right corner |
| game-16-chat-group-hint.jpg | `/fs group` solo: "The group panel shows on the Raid tab of the Social window (O)." |

Also: `/fs panel` printed nothing and opened nothing visible while solo (`/fs diag` says "Panel: closed, 2 targets, rotation 5 rows").

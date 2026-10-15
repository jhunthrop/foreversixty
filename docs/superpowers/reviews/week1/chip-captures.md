# Chip captures, round 1 (2026-10-11, L3.1 at ea8ffd60)

Client: WoW Classic beta 1.60.1.70291. Character: Obnoxious Yell, level 21 Gnome Warrior, Arms 8 / Fury 4 / Protection 0, 0 unspent points (the client's talent window says "Unspent Talents 0"), guild OLYMPUS XXVII, standing in Stormwind's Trade District. Build loaded: the 51-point Arms/Fury build pasted in round 2 (12 of 51 matched; the plan's next point is #13, taken at level 22).

| File | What it shows |
|---|---|
| chip-screen-after-reload.jpg | Whole screen after `/reload`: the chip sits left of the minimap (top right), chat holds the `/fs` hint and `/fs diag` output |
| chip-reload-chat.jpg | Chat zoom: `/fs` prints the command list; `/fs diag` prints "What this client turned out not to have: The Forever Sixty Data addon could not be read. Update both addons." |
| chip-point.jpg | The chip, zoomed: talent icon, "Take Improved Tactical Mast…", a small "+1" badge top right |
| chip-tooltip.jpg | Hovering the chip: "Take Improved Tactical Mastery" (body), "Equip Burnished Shield, +0.9 DPS" (muted), "Right-click: Hide until next level", "Shift-drag to move" |
| chip-click-talent-window.jpg | After clicking the chip: the client's talent window opened (Arms 8, Fury 4, Unspent Talents 0); the chip's tooltip still showing |
| chip-hidden-after-right-click.jpg | After right-clicking the chip: the chip is gone |
| chip-restored-after-reload.jpg | After clearing the hide pref and `/reload`: the chip is back |

Not captured this round: idle (the character always has an upgrade or a point waiting), combat (no hostile mob in the city), "here for you" (needs The Deadmines), the path card (needs a level 1–9 alt), raid check and last night (no readers yet).

Known facts the reviewers should weigh:
- The character has 0 unspent points, yet the chip's top state is "Take Improved Tactical Mastery" with the talent window as the click target. The plan's next point is taken at level 22, one level away.
- The character wields a two-handed weapon (Arms). "Equip Burnished Shield, +0.9 DPS" is the top bag upgrade the chip offers.

# Chip captures, round 2 (2026-10-11, L3.1 fix round 1 at 25aca89f)

Same character (level 21, 12 points spent, 0 unspent, two-hander worn). After the fix the chip has nothing to say, so it is idle.

| File | What it shows |
|---|---|
| chip2-screen-after-reload.jpg | Whole screen after `/reload`, `/fs`, `/fs diag`: the chip is the idle class icon left of the minimap |
| chip2-reload-chat.jpg | Chat zoom: `/fs` lists "/fs chip: show or hide the chip", "/fs inspect: look at your target", "/fs diag: what is not working", "/fs help: every command"; `/fs diag` prints "Not working:", "The Forever Sixty Data addon is missing or too old. Install or update it from the site." (the data addon is installed here and is the old format, so this is right) and "Chip: anchored, 28 of 280 wide, not truncated" |
| chip2-tooltip-idle.jpg | Hovering the idle chip: "Forever Sixty" (gold), "Next at 16: Improved Tactical Mastery", blank, "Right-click: Hide until next level", "Shift-drag: Move" |
| chip2-idle.jpg | The idle chip zoomed: the class icon at 40 % alpha over a bright roof; hard to find |
| chip2-hidden-chat-line.jpg | After right-click: "Chip hidden until level 22. /fs chip brings it back." |
| chip2-fs-chip-restore.jpg | After `/fs chip`: "Chip shown." |

Known facts: the character's next talent point arrives at level 22 (12 spent, level 21), so "Next at 16" is wrong (it is the plan's order index of the first unmet cell, 9 + 7). The idle chip has no ground or border, only the faint icon.

## Round 2, the level-1 alt (Shorty Small, Gnome Rogue, Coldridge Valley)

| File | What it shows |
|---|---|
| chip2-alt-screen-pick-a-path.jpg | Whole screen right after login: the chip reads "Pick a path" with a "+1" badge |
| chip2-alt-path-chip.jpg | The chip zoomed: class icon, "Pick a path", "+1" |
| chip2-alt-path-tooltip.jpg | Hover: "Forever Sixty", "Pick a path", "Next at 10: Improved Gouge", blank, "Click: Choose a spec", "Right-click: Hide until next level", "Shift-drag: Move" |
| chip2-alt-screen-path-card.jpg | After clicking: the card under the chip, "Gnome Rogue. Pick a path:" with Assassination / Combat / Subtlety gold buttons and a playstyle line under each |
| chip2-alt-path-card.jpg | The card zoomed |
| chip2-alt-after-choice.jpg | After choosing Combat: the chip reads "Next at 10: Improved Eviscerate" with the talent's icon |
| chip2-alt-after-choice-tooltip.jpg | Hover after the choice: title, the next line, blank, the two hints |

Known facts: the "+1" on the path chip counts only the informational "Next at 10" line (nothing else waits). The playstyle lines are the addon's own copy. "Next at 10" is right for a level-1 character.

# Chip captures, round 3 (2026-10-11, L3.1 fix round 2 at 81812c99)

## The level-1 alt (Shorty Small, Gnome Rogue, Combat path chosen)

| File | What it shows |
|---|---|
| chip3-alt-screen.jpg | Whole screen after `/reload`: the chip reads "Next at 10: Improved Eviscerate" with no badge |
| chip3-alt-next-line.jpg | The chip zoomed: talent icon, the next line, no "+N" |
| chip3-alt-diag.jpg | `/fs diag`: "Not working:", the data-addon line (installed, old format), "Chip: anchored, 229 of 280 wide, not truncated" |
| chip3-alt-tooltip.jpg | Hover: "Forever Sixty", "Next at 10: Improved Eviscerate", blank, the two hints |

## The level-21 warrior again (Obnoxious Yell, Stormwind, 12 points spent, 0 unspent, two-hander)

| File | What it shows |
|---|---|
| chip3-main-screen-idle.jpg | Whole screen after login: the idle chip (dark ground, faint class icon) left of the minimap |
| chip3-main-idle.jpg | The idle chip zoomed: 28×28 dark ground with the class icon at 40 % |
| chip3-main-tooltip.jpg | Hover: "Forever Sixty", "Next at 22: Improved Tactical Mastery", blank, the two hints |

Facts: 22 is right (12 spent at level 21). No point state, no shield, no badge. Still not captured: a real unspent point with the ring, the upgrade state with a legitimate item, combat, "here for you" in The Deadmines, raid check and last night.

# Chip captures, round 4 (2026-10-11, L3.1 fix round 3 at 44c60d1b), the level-21 warrior

| File | What it shows |
|---|---|
| chip4-main-diag.jpg | `/fs diag`: "Not working:", "The Forever Sixty Data addon is too old. Update it from the site.", then "Chip: shown, nothing cut off" outside the heading |
| chip4-main-idle.jpg | The idle chip: dark ground, the client's round class icon at 40 %, muted border |
| chip4-main-tooltip.jpg | Hover: "Forever Sixty" (gold), "Next talent at 22: Improved Tactical Mastery" (muted), blank, "Click: Open talents", "Right-click: Hide until next level", "Shift-drag: Move" (dimmer) |

# Chip round 5 and talent marks round 1 (2026-10-11, L3.2 at 2e7f2b59), the level-21 warrior with all 12 points unlearned

The user unlearned the warrior's talents, so 12 points are unspent and the plan's first unmet cell is Improved Heroic Strike (the pasted build's first point).

| File | What it shows |
|---|---|
| chip5-point-state.jpg | The chip in the point state: talent icon with a gold ring, "Take Improved Heroic Strike", no badge |
| chip5-screen-point.jpg | Whole screen with the chip in the point state |
| chip5-point-tooltip.jpg | Hover: title, "Take Improved Heroic Strike" (body), blank, "Click: Open talents", the two hints |
| chip5-point-after-6s.jpg | The chip about six seconds later (the ring still looks the same in stills; the pulse could not be told from three 0.6 s stills) |
| marks1-screen-talent-window.jpg | After clicking the chip: the client's talent window, Unspent Talents 12; NO ring, NO badges, NO strip |
| marks1-trees.jpg / marks1-trees-all.jpg | The trees zoomed: every tier-1 node has the client's own green "available" border; nothing from the addon |
| marks1-window-bottom.jpg | The window's bottom edge: only the client's Apply Changes button, no strip |
| marks1-diag.jpg | `/fs diag`: "No talent button matched the next point; the chip still names it." and "18 of 18 talent badges found no button; badges are off" |

Facts: in round 1 (B12/B13, before the chip) the ring rendered on this same client and character. After unlearning, with the trait window showing every tier-1 node available, neither the ring nor any badge found a button, and the strip did not draw at all.
| marks1-second-open-trees.jpg | Closing and reopening the window: the gold ring now draws on Improved Heroic Strike (top-left Arms node); still no badges (latched off by the first miss) and still no strip |
| marks1-second-open-diag.jpg | `/fs diag` after the second open: "Talent window found: PlayerSpellsFrame (trait mapping)", "Glow on button Improved Heroic Strike, width 39.99, shown true" |

So the first open after a reload runs the marks before the trait buttons exist (the lookup fails for every node and the session latch turns badges off), the second open finds the buttons; and the strip never draws even when the frame is found.

# Talent marks round 2 (2026-10-11, L3.2 fix round 1 at 33aa9c13), the level-21 warrior, 12 points unspent

First open after `/reload`, two seconds after the click on the chip.

| File | What it shows |
|---|---|
| marks2-screen.jpg | Whole screen: the talent window with the gold ring on Improved Heroic Strike (Arms, top left), a number badge on every planned node across Arms, Fury and one Protection node, and the strip pill at the window's bottom left; Unspent Talents 12 |
| marks2-trees.jpg | The three trees zoomed: badges 3, 6, 11, 13, 14, 17, 20, 22, 23, 28, 30, 31 (Arms), 36, 41, 42, 45, 50, 51 (Fury); the ring on the top-left Arms node |
| marks2-strip.jpg | The strip pill: "Warrior build · 0 of 51 · next at level 21 · 0 off plan", left-aligned, clear of the client's Apply Changes button |
| marks2-diag.jpg | `/fs diag`: the window and glow notes, and "Talent strip: edge anchor on TalentsFrame, at 466, 459, 274 by 24, strata HIGH level 110, visible true; talents frame level 100" |

Facts: the loaded build is the pasted 51-point Arms/Fury build with no name (so the strip says "Warrior build"); 0 of 51 is right after the unlearn; the character is level 21 with 12 points to spend now, so "next at level 21" means "now". All 18 badges found their buttons on the first open after the fix.

# Talent marks round 3 (2026-10-11, L3.2 fix round 2 at 70e61914), the level-21 warrior, 12 points unspent

| File | What it shows |
|---|---|
| chip6-point-rank.jpg | The chip: "Take Improved Heroic Strike (1 of 3)" with the ring on the icon |
| marks3-screen.jpg | The talent window on first open: badges on dark plates with gold digits, the next node's badge inverted (gold plate), unreachable nodes' plates dimmer; strip "Arms / Fury build · 0 of 51 points · spend 12 now · 0 off build" |
| marks3-trees.jpg | The trees zoomed |
| marks3-strip.jpg | The strip pill at 12 px, clear of Apply Changes |
| marks3-node-tooltip.jpg | Hovering Improved Tactical Mastery: the client's tooltip ends with "Forever Sixty: points 7 to 11 of your build" |
| marks3-staged-on-build-trees.jpg | After one click on Improved Heroic Strike (staged, not applied): the node shows rank 1, the ring stays on it (the plan wants 3), Apply Changes lit |
| marks3-staged-on-build-strip.jpg | The strip: "Arms / Fury build · 1 of 51 points · spend 11 now · 0 off build · click Apply Chan…" (cut by the pill's cap at the Apply button) |
| marks3-staged-off-build-strip.jpg | After one click on a Protection node the plan does not want: "… spend 10 now · 1 off build" with the count in red, before Apply |
| marks3-staged-off-build-node.jpg | That Protection node at rank 1 |

Facts: both staged clicks were discarded with the client's reset button afterwards (nothing applied). The only defect seen: the "click Apply Changes" segment is cut mid-word by the pill's width cap.

# Talent marks round 4 (2026-10-11, L3.2 fix round 3 at 45fc3475), the level-21 warrior, 12 points unspent

| File | What it shows |
|---|---|
| marks4-screen.jpg | First open: the chip is gone from beside the minimap while the window is open; plates with digits readable on dim plates; node "13" bright; strip "Arms / Fury build · 0 of 51 points · spend 12 now · 0 off plan" |
| marks4-trees.jpg | The trees zoomed |
| marks4-strip.jpg | The strip pill |
| marks4-node-tooltip.jpg | Hovering Improved Tactical Mastery: a blank line, then "Forever Sixty:" in gold and "Your build: 5 ranks, points 7 to 11" in white, after the client's red requirement line |
| marks4-offplan-node-tooltip.jpg | After one staged click on a Protection node the plan does not want, hovering it: "Forever Sixty: not in your build" in red at the tooltip's end |
| marks4-offplan-strip.jpg | The strip: "0 of 51 points · spend 11 now · 1 off plan" (red) with Apply Changes lit; one staged click, 12 − 1 = 11 |
| marks4-chip-after-close.jpg | The chip back after the reset and Escape: "Take Improved Heroic Strike (1 of 3)" |

Facts: the staged point was discarded with the client's reset button. Not captured: the pulse frames, an applied point, the 0-unspent "next point at level N" state, other UI scales.

# Talent marks round 5 (2026-10-11, L3.2 fix round 4 at ec979b2f), the level-21 warrior, 12 points unspent

| File | What it shows |
|---|---|
| marks5-screen.jpg | First open: range badges on opaque plates ("1–3" inverted on the next node, "4–6", "7–11", "12–13", "14", "15–17" … "32–36", "37–41", "46–50", "51"); the strip leads with "Next: Improved Heroic Strike (1 of 3) · 0 of 51 points · 0 off plan" (the "spend 12 now" and name segments were dropped to fit the pill's cap at Apply Changes) |
| marks5-trees.jpg | The trees zoomed |
| marks5-arms-badges.jpg | Arms top rows zoomed: the en dash renders; dim plates are dark and their digits readable |
| marks5-fury-badge.jpg | The Fury top-right node's "32–36" badge, now readable on the bright icon |
| marks5-strip.jpg | The strip pill |

# Item tooltip round 1 (2026-10-11, L3.3 at da64c179), the level-21 warrior (two-hander worn, bags open at a trainer)

| File | What it shows |
|---|---|
| tip1-bags-open.jpg | The bags window |
| tip1-shield-under-two-hander.jpg | Hovering a bag Burnished Shield: the client's two compare tooltips (Goblin Hammer worn, Smite's Mighty Hammer) and the shield's tooltip ending with "Forever Sixty · Default weights" and "Fills an empty slot, +0.9 DPS" — an off-hand offered under a worn two-hander |
| tip1-mace-no-section.jpg | Hovering Cookie's Tenderizer (one-hand mace, +4 Spirit): NO Forever Sixty section at all |
| tip1-hearthstone.jpg | Hovering the Hearthstone: no section (right) |

Facts: the worn main hand is a two-hander (Smite's Mighty Hammer) with Goblin Hammer in the compare; quest items and food show no section (right). Not captured yet: a worn item ("Equipped"), a vendor window, a quest reward, a plate item ("until 40"), the "Your sim" tag with the inbox fixture, a downgrade's minus sign.
| tip1-worn-chest.jpg | Character pane, hovering the worn chest (Abomination Bones): "Forever Sixty · Default weights", "Equipped", "Next: Mutant Scale Breastplate · Wailing Caverns · Mutanus the Devourer · 33%", "Also: Veteran's Chain Shirt, Defender's Leather Armor"; the compare tooltip for the Mutant Scale Breastplate beside it |
| tip1-empty-finger.jpg | Hovering the empty finger slot: the old empty-slot section (the band's pick row with a "new at 20" tag and its reputation source) |
| tip1-worn-mainhand.jpg | Hovering the worn two-hander: "Equipped", "Next: … · Shadowfang Keep · Sever", "Also: Living Root, Duskbringer", with the client's own "Retrieving item information" line above |

# Item tooltip round 2 (2026-10-11, L3.3 fix round 1 at b4a50d8a), the level-21 warrior (two-hander worn, bags open at a trainer), no inbox message loaded

| File | What it shows |
|---|---|
| tip2-shield.jpg | Hovering the bag Burnished Shield: the section ends with "Forever Sixty · Arms · Estimate" and a muted "Not with your two-hander"; no verdict, no compare tooltip from us |
| tip2-mace.jpg | Hovering Cookie's Tenderizer (one-hand mace): "Forever Sixty · Arms · Estimate" then "Downgrade −5.7 DPS over Smite's Mighty Hammer" in red (the worn two-hander is the comparison) |
| tip2-chest.jpg | Character pane, hovering the worn chest (Abomination Bones): "Forever Sixty · Arms · Estimate" with "Equipped" muted at the right; "Next upgrade: Mutant Scale Breastplate" (name in its quality colour); "Wailing Caverns · Mutanus the Devourer · 33% drop" on its own muted line; "Also good: Veteran's Chain Shirt"; the client's compare tooltip to the right titled "Next upgrade" |

Facts: the "Your sim" tag is still not reachable in the client because the inbox file is never loaded (root cause recorded in the ledger; the stub addon lands in fix round 2). The Next line in tip2-chest.jpg shows the name only; no "+X DPS" gain is visible after it. Not captured yet: a vendor window, a quest reward, a plate item ("until 40"), "Your sim", an Upgrade in a bag, the hearthstone with the addon disabled.

# Item tooltip round 3 (2026-10-11, L3.3 fix round 2 at 5f6ff33b), the level-21 warrior, the stub inbox addon installed and the controller's sim-weights fixture loaded (Str 2.2, Agi 1.1, AP 1, dps_per_point 0.42)

| File | What it shows |
|---|---|
| tip3-shield.jpg | Bag Burnished Shield, uncropped: "Forever Sixty · Arms · Your sim" then "Not with your two-hander" (muted, brighter than round 2); no compare tooltip from us |
| tip3-mace.jpg | Cookie's Tenderizer, uncropped: "Forever Sixty · Arms · Your sim", "Downgrade −12.0 DPS over Smite's Mighty Hammer" in red (11 Str × 2.2 + 4 Agi × 1.1 = 28.6 AP × 0.42; weapon DPS is not yet counted) |
| tip3-chest.jpg | Worn Abomination Bones: "Forever Sixty · Arms · Your sim" + "Equipped"; "Next upgrade: Mutant Scale Breastplate +1.8 DPS"; "Wailing Caverns · Mutanus the Devourer · 33% drop"; "Also good: Veteran's Chain Shirt" (label muted); compare titled "Next upgrade" |
| tip3-worn-mainhand.jpg | Worn Smite's Mighty Hammer (two-hander): header + "Equipped" and NO primary line (neither "Next upgrade" nor "Best for now"); no compare; no "Retrieving item information" |
| tip3-worn-ranged.jpg | Worn Dull Sawblade (thrown): header + "Equipped" and "Best for now" |
| tip3-empty-finger.jpg | Empty finger slot: the old empty-slot section ("Finger", "Forever Sixty", "[Demon Band] new at 20", "World Drop · World drop") — L4.1 replaces it |

Facts: the stub addon was detected at the character screen (AddOn list shows "Forever Sixty Inbox" nested under Forever Sixty); after one /reload the client rewrote ForeverSixty.lua without the inbox variable and kept the fixture in ForeverSixtyInbox.lua intact. Defect seen: the worn two-hander shows neither a Next line nor "Best for now". Not captured yet: a vendor window, a quest reward, a plate item, an Upgrade/Sidegrade in a bag, a long-name compare, the hearthstone.

Note (L3.3 fix round 3, uncommitted at the time of writing): weapon DPS now counts at 14 attack power per DPS for melee and hunter specs (Progress.itemScore), shared by the chip's bag upgrades and the item tooltip, so the chip's upgrade figures and the ring's "Take" lines may differ from the captures above.

# Item tooltip round 4 (2026-10-11, L3.3 fix round 3 at cdadee86), the level-21 warrior, sim-weights fixture loaded (Str 2.2, Agi 1.1, AP 1, dps_per_point 0.42 — a controller fixture value, about six times the 14-AP-per-DPS rule)

| File | What it shows |
|---|---|
| tip4-mace.jpg | Cookie's Tenderizer: "Forever Sixty · Arms · Your sim" ("Your sim" brighter), "Downgrade −46.0 DPS over Smite's …" — the comparison name ellipsised to the tooltip's width; weapon DPS now counted (28.6 AP + 5.8 × 14 = 109.8 AP × 0.42) |
| tip4-shield.jpg | Burnished Shield: "Not with your two-hander" (brighter muted) |
| tip4-worn-mainhand.jpg | Worn Smite's Mighty Hammer: header + "Equipped" + "Best until 30" (the band pick is client-unconfirmed; its alternatives did not beat the worn two-hander by 1.0 with weapon DPS counted) |
| tip4-chest.jpg | Worn Abomination Bones: "Next upgrade: Mutant Scale Breastplate +1.8 DPS (enchant not counted)" with the gain green and the note muted; source line; "Also good: Veteran's Chain Shirt, Defender's Leather Armor" (both names) — but the whole tooltip is now visibly wider than the client's own widest line ("Press F6 …"): our Next line widened it and the alternatives then used the new width |
| tip4-worn-ranged.jpg | Worn Dull Sawblade: "Best until 30" |

Facts: the chip still reads "Take Improved Heroic Strike (1 of 3)". Defects seen: the chest tooltip widens past the client's own lines (the Next line with the enchant note is not width-fitted); the mace's comparison name is reduced to "Smite's …" because that tooltip is narrow. Not captured yet: a vendor window, a quest reward, a plate item, an Upgrade/Sidegrade in a bag, the hearthstone.

# Item tooltip round 5 (2026-10-11, L3.3 fix round 4 at dd64ab7f), the level-21 warrior, sim-weights fixture loaded (dps_per_point 0.42, a fixture value)

| File | What it shows |
|---|---|
| tip5-mace.jpg | Cookie's Tenderizer with the client's own "Equipped" compare beside it: "Forever Sixty · Arms · Your sim", "Downgrade −52.8 DPS over Smite's Mighty Hammer" on one line (the client's own "Chance on hit" line is wide, so the full name fits); Smite's Weapon Damage +4 enchant now netted |
| tip5-shield.jpg | Burnished Shield: "Not with your two-hander"; the tooltip at the client's natural width |
| tip5-chest.jpg | Worn Abomination Bones: header with "Equipped" on the header row; "Best until 30" (the Mutant Scale Breastplate fell under the 1.0 floor once the worn +4 AP enchant was netted); "At 30: Avenger's Armor" (name in quality colour) — NO source after the name; the tooltip at the client's natural width |
| tip5-worn-mainhand.jpg | Worn Smite's Mighty Hammer: "Best until 30", "At 30: Manual Crowd Pummeler" — no source |
| tip5-worn-ranged.jpg | Worn Dull Sawblade: "Best until 30", "At 30: Glass Shooter" — no source |

Facts: no tooltip is widened by our lines any more; the chip still reads "Take Improved Heroic Strike (1 of 3)". Defect seen: every "At 30:" line lacks its "· <source>" part. Not captured yet: a vendor window, a quest reward, a plate item, an Upgrade/Sidegrade in a bag, the hearthstone, a wrapped Next line.
| tip5b-mace.jpg | Same hover after the controller set the fixture's dps_per_point to the shipped sim value 0.2342: "Downgrade −29.4 DPS over Smite's Mighty Hammer" (28.6 AP + (5.8 + 4/3.5) × 14 AP = 125.6 AP × 0.2342); the chest is unchanged ("Best until 30 / At 30: Avenger's Armor") |

# Item tooltip round 6 (2026-10-11, L3.3 fix round 5 at b595d5c6), the level-21 warrior, sim-weights fixture at the shipped dps_per_point 0.2342

| File | What it shows |
|---|---|
| tip6-chest.jpg | Worn Abomination Bones: "Forever Sixty · Arms · Your sim" + "Equipped"; "Best until level 30" (brighter); "At level 30: Avenger's Armor" then "· Razorfen Kraul" on its own muted line (branch b of the fit rule); natural width |
| tip6-worn-mainhand.jpg | Worn Smite's Mighty Hammer: the header is ELLIPSISED to "Forever Sixty · Arms · Yo…" beside "Equipped" (the fitter cut the "Your sim" tag on this narrower tooltip); "Best until level 30"; "At level 30: Manual Crowd Pummeler" / "· Gnomeregan" |
| tip6-worn-ranged.jpg | Worn Dull Sawblade: the same header ellipsis "… · Yo…"; "At level 30: Glass Shooter" / "· Gnomeregan" |
| tip6-mace.jpg | Cookie's Tenderizer: "Downgrade −29.4 DPS over Smite's Mighty Hammer" (unchanged) |
| tip6-shield.jpg | Burnished Shield: "Not with your two-hander" (unchanged) |

Facts: the At-line source now shows on every worn item; the tooltips stay at the client's width. Defect seen: fitting the header by ellipsis cuts the "Your sim"/"Estimate" tag on narrow worn-item tooltips (the header plus "Equipped" exceeds the client's width). Still not captured: a vendor window, a quest reward, a plate item, an Upgrade/Sidegrade in a bag, the hearthstone, a wrapped Next line (none exists for this character's worn slots per the fix 5 report).

# Item tooltip round 7 (2026-10-11, L3.3 fix round 6 at 6fece3a6), the level-21 warrior, sim-weights fixture at the shipped dps_per_point 0.2342

| File | What it shows |
|---|---|
| tip7-chest.jpg | Worn Abomination Bones: full header "Forever Sixty · Arms · Your sim" + "Equipped"; "Best until level 30"; "At level 30: Avenger's Armor" / "· Razorfen Kraul" |
| tip7-worn-mainhand.jpg | Worn Smite's Mighty Hammer (narrower tooltip): header step b "Forever Sixty · Your sim" + "Equipped" on the header row, tag whole; "Best until level 30"; "At level 30: Manual Crowd Pummeler" / "· Gnomeregan" |
| tip7-worn-ranged.jpg | Worn Dull Sawblade: the same step-b header; "At level 30: Glass Shooter" / "· Gnomeregan" |
| tip7-mace.jpg | Cookie's Tenderizer: "Forever Sixty · Arms · Your sim"; "Downgrade −29.4 DPS over Smite's Mighty Hammer" |
| tip7-shield.jpg | Burnished Shield: "Not with your two-hander" |

Facts: no header is ellipsised; no tooltip is widened; every worn item ends with exactly one primary line. Still not captured on this character: a vendor window, a quest reward, a plate item, an Upgrade/Sidegrade in a bag, the hearthstone, a wrapped Next line (no Next exists for the worn slots per the fix 5 report).

# Item tooltip round 8 (2026-10-11, L3.3 fix round 7 at fe90fa89), the level-21 warrior, sim-weights fixture at the shipped dps_per_point 0.2342

| File | What it shows |
|---|---|
| tip8-mace.jpg | Cookie's Tenderizer: "Downgrade −29.4 DPS over Smite's Mighty Hammer" with the muted breakdown "(−11 Str, −4 Agi, −6.9 weapon DPS)" under it |
| tip8-chest.jpg | Worn Abomination Bones: "Best until level 30"; "At level 30: Avenger's Armor +2.7 DPS" (gain green) then "· Razorfen Kraul" |
| tip8-worn-mainhand.jpg | Worn Smite's Mighty Hammer: "At level 30: Manual Crowd Pummeler" with no gain (the pick's stats were cold or the gain is under the floor) then "· Gnomeregan" |

Facts: the breakdown reconciles the figure with the client's own lines; widths natural; the header step b still holds on the hammer.

# Build dialog round 1 (2026-10-11, L3.4 at 59cd4983), the level-21 warrior, no inbox fixture messages of type build/ack

| File | What it shows |
|---|---|
| dialog-1-screen.jpg | `/fs`: the dialog centred on screen (whole screen at 0.6) |
| dialog-1-empty.jpg | The dialog: title "Forever Sixty", tabs "Paste a build" (gold, underlined) and "Copy my character", the box with the muted "Paste a build code"; keyboard in the box |
| dialog-2-preview.jpg | After typing the Arms code: the box shrunk to two visible lines of code; warrior icon, "Arms / Fury build" in gold, "Warrior · Arms 33 / Fury 16 / Protection 2 · 51 points" muted, three rows (Lionheart Helm, Beads of Ogre Might, Highlander's Plate Spaulders) with icons, the gold "Use this build" button |
| dialog-3-chat.jpg | After Enter: chat "Forever Sixty: Now following Arms / Fury build." |
| dialog-3-chip.jpg | The chip now reads "Take Deflection (1 of 2)" (the raid default's first point) |
| dialog-4-refused.jpg | The 1.99.0.99999 code: the code stays in the box; below it muted "That code is for a newer patch than this addon. Update the addon."; no preview, no button |

Facts: Escape closed the dialog on the first press with the box focused, and W moved the character afterwards (keyboard free). No Lua error with scriptErrors on. Not yet captured: copy mode, the ack/Seen line and the inbox row (fixture at character select), Tab, drag/reset, minimap/compartment/keybind, /fs help.
| dialog-5-copy.jpg | `/fs copy`: the one-line box with the selected FS1 export, muted "Selected: press Ctrl+C to copy it.", the gold "Copy my character" button, and at the bottom "The companion, the small app on your PC, sends this for you when you log out." |
| dialog-5-copied.jpg | After clicking the button: "Copied 21:53, paste it on the site" under the button; gone after 3 s |
| dialog-10-tab-paste.jpg | Tab from copy mode: paste mode with the box focused; Tab again returns to copy mode |
| dialog-13-help.jpg | `/fs help`: the command list (the tail shown: diag, inbox, hello, seen, record, help) |

Facts: the first Tab press after a missed click did nothing (the keyboard was not in the box); with the box focused Tab switches both ways. Capture 6/7 (ack and inbox build row) follow after the fixture edit at character select.
| dialog-6-seen.jpg | With the ack fixture (export_seen_at 03:30Z, sim queued): under the button "Seen by the site 22:30" and "Sim queued" (local time) |
| dialog-6-copied-then-seen.jpg | After clicking the button: "Copied 21:55, paste it on the site" with "Sim queued" under it; after 3 s "Seen by the site 22:30" returns |
| dialog-7-inbox-row.jpg | With the build fixture: above the box "Arms raid · from the site" with a "Use" button at the right |
| dialog-7-inbox-chat.jpg | `/fs inbox`: "Forever Sixty: 1 build(s) from the site waiting. Type /fs to use one." |
| dialog-7-use-chat.jpg | After clicking Use: "Forever Sixty: Now following Arms raid."; the chip still reads "Take Deflection (1 of 2)" (the same order as the pasted default) |
| dialog-7-after-use.jpg | `/fs` again: no "from the site" row (dismissed) |

Facts: nothing loaded on its own at login with the build message present (the chip's follow stayed as pasted until Use). Not captured: drag/reset (11), minimap button/compartment/keybind (12), the top of `/fs help` (13 shows the tail).

# Build dialog round 2 (2026-10-11, L3.4 fix round 1 at 1e5ded18), the level-21 warrior, fixture with the ack (03:30Z, sim queued); the build message was used in round 1 so no row shows

| File | What it shows |
|---|---|
| dialog2-1-empty.jpg | `/fs`: tabs with "Tab switches" at the strip's right end; "Following: Arms raid · Warrior · 51 points"; the box with a gold focus border and the placeholder "Paste a build code from the site (Ctrl+V)"; "Next talent: Take Deflection (1 of 2)" |
| dialog2-2-preview.jpg | After the Arms code: the box collapsed to one dim line ending in "…"; "Arms / Fury build"; "Warrior · Arms 33 / Fury 16 / Protection 2 · 51 points · you can place 12 now"; "Gear in this build · at 60"; three rows with names in quality colour (epic, rare, epic) and "Head / Neck / Shoulders" right-aligned; "Use this build" |
| dialog2-4-refused.jpg | "hello there": red border; "That is not a Forever Sixty build code." in red; "You are still following Arms raid." muted |
| dialog2-5-copy.jpg | Copy mode: "Your character is selected. Press Ctrl+C, then paste it on the site."; the export wrapped over five lines, selected; the secondary "Copy my character"; "Seen by the site 22:30"; "Sim queued. The result arrives when you next log in." |
| dialog2-5-copied.jpg | After Ctrl+C: "You copied 22:21." replaces the lead line |

Facts: the copy box shows a partly clipped sixth line of glyph tops under "|level=21" (the paste box's clipped-line fix did not reach the copy box). Not captured this round: the inbox row as a list entry feeding the preview (the fixture's build was already used; a fresh build message is needed at character select), the wrong-class refusal, the no-ack companion lines, drag/reset, minimap/compartment/keybind.

# Build dialog round 3 (2026-10-11, L3.4 fix round 2 at 0596792d), the level-21 warrior; the fixture's ack (03:30Z) is now OLDER than the export saved at this reload

| File | What it shows |
|---|---|
| dialog3-1-empty.jpg | `/fs`: "Following: Arms raid · Warrior · 51 points" and "Next talent: Take Deflection (1 of 2)" together above a two-line box with the placeholder "Paste a build code from foreversixty.gg (Ctrl+V)"; "Tab to switch" dim at the bottom-left; no hint in the tab strip |
| dialog3-2-preview.jpg | After the Arms code: the collapsed code line is an unreadable grey bar (no text visible); "Arms / Fury build"; the muted spec line; "You can place 12 now." in body colour; "Gear in this build · at 60"; slot-first rows (Head / Neck / Shoulders, icon, quality-coloured name); "Replaces Arms raid"; "Next talent would be: Take Deflection (1 of 2)"; "Use this build"; the panel is full but nothing overflows |
| dialog3-4-wrong-class.jpg | A mage code: red border; "That code is for a Mage, not a Warrior." in red; "You are still following Arms raid." muted; no advice line |
| dialog3-5-copy.jpg | Copy mode: the lead line wraps to two lines ("…paste it on / foreversixty.gg."); the export over ten whole lines (the last visible line is cut at the right edge of the box, not clipped vertically); no button; "The site has your character from 22:30. Yours has changed since."; the sim line |
| dialog3-5-copied.jpg | After Ctrl+C: "You copied 22:43. Now paste it on foreversixty.gg." in green |

Facts: the "changed since" line is correct here because the reload saved a newer export than the fixture's ack. Defects: the collapsed code line renders as a solid grey bar; the class refusal has no advice sentence; a long export line is cut at the box's right edge (the box wraps at word boundaries only, and the export has none). Not captured: the inbox row feeding the preview, the no-ack companion lines, the older-ack "yesterday" wording, control states, drag/reset, minimap/compartment/keybind.

# Build dialog round 4 (2026-10-11, L3.4 fix round 3 at 99ca69f3), the level-21 warrior; fixture: ack at yesterday 20:00 local (sim queued), a fresh site build "Arms raid (old patch)" (1.60.1.69893 code)

| File | What it shows |
|---|---|
| dialog4-1-row.jpg | `/fs`: "Next talent: Take Deflection (Arms, 1 of 2)"; the row "Arms raid (old patch) · Arms 33 / Fury 16 / Prote… · from the site" (the split ellipsised, the name whole) above the two-line box |
| dialog4-2-row-preview.jpg | After clicking the row: the row stays in the code line's place; the preview "Arms raid (old patch)", the spec line, "You can place 12 now.", slot-first gear rows at a 20px pitch, "Replaces Arms raid · same point order", "Next talent would be: Take Deflection (Arms, 1 of 2) · same as n…" (cut), "Use this build", "Tab to switch" and "Enter to use" at the bottom corners |
| dialog4-3-use-hover.jpg, dialog4-3-use-pressed.jpg | The Use button at rest/hover and held mid-click: no visible difference in either zoom |
| dialog4-3-close-hover.jpg | The close x hovered: a darker square ground with the x in body colour |
| dialog4-4-refused-advice.jpg | "hello there": red sentence, the body-colour advice "Copy the whole code from foreversixty.gg. It starts with FSB1.", the muted "still following" line |
| dialog4-5-copy.jpg | Copy mode: the Paste tab reads "Paste a build  1"; the one-line lead; the export box is a GREY BLOCK with only fragments readable ("=21", "89") — the invisible box's selection highlight sits over the wrapping label; "The site has your character from yesterday 20:00. Yours has changed since."; "Sim queued for the yesterday 20:00 copy. The result arrives when you next log in." |
| dialog4-5-copied.jpg | After Ctrl+C: the green copied line; "Waiting for the site to see this copy."; the box still a grey block |

Facts: the "yesterday" wording and the post-copy line work. Defects: the export is unreadable (the highlight of the invisible edit box is drawn over the label); the Use button shows no hover or pressed change; "· same as now" is cut on the Next-would-be line; "for the yesterday 20:00 copy" reads oddly. Not captured: the newer-patch refusal with advice (round 1 showed the old form), drag/reset, minimap/compartment/keybind.

# Build dialog round 5 (2026-10-11, L3.4 fix round 4 at 93192c82), the level-21 warrior; fixture as round 4

| File | What it shows |
|---|---|
| dialog5-2-row-preview.jpg | The row preview: "Replaces Arms raid · same point order"; "Next talent would be: Take Deflection (Arms, 1 of 2)" (whole); "Use this build"; "Tab to switch" / "Enter to use" |
| dialog5-3-use-hover.jpg, dialog5-3-use-pressed.jpg | The Use button with the pointer over it and held mid-click: no visible change from rest in either zoom (the frame spec pins the wiring; the capture tool's synthetic pointer may not raise OnEnter/OnMouseDown in this client — a human hover must confirm) |
| dialog5-5-copy.jpg | Copy mode: the export readable, wrapped per character over ten lines with a selection tint and a scrollbar thumb at the right; "The site has your character from yesterday 20:00. Yours has changed since."; "Sim queued for yesterday's 20:00 copy." |
| dialog5-5-copied.jpg | After Ctrl+C: the green copied line; the box unchanged; "Waiting for the site to see this copy."; the sim line |

Facts: the grey block is gone; the sim wording reads right; nothing overflows. Open: the Use button's hover/pressed states are not demonstrable with the capture tool's pointer. Not captured: the newer-patch refusal with advice, drag/reset, minimap/compartment/keybind.

# Build dialog round 6 (2026-10-12, L3.4 fix round 5 at d5e146d4), the level-21 warrior after a fresh login; fixture as round 4 (the ack is now two days old)

| File | What it shows |
|---|---|
| dialog6-1-row.jpg | `/fs`: the row "Arms raid (old patch) · 33/16/2 · from the site" |
| dialog6-2-row-preview.jpg | The row selected (gold left edge) feeding the preview: "Gear in this build · at 60 · 3 of 17 slots"; "Replaces Arms raid · same points and gear"; "Next talent would be: Take Deflection (Arms, 1 of 2)"; "Use this build"; corner hints brighter |
| dialog6-5-copy.jpg | Copy mode: the Paste tab reads "Paste a build  1 waiting"; the one-line lead; the export readable and selected; "The site has your character from 10 Oct 20:00. Yours is newer."; "Sim done · your sim weights are in use" (no best-next-upgrade line: the chip's next item is a talent) |
| dialog6-5-copied.jpg | After Ctrl+C: "You copied 17:10. Now paste it on foreversixty.gg."; the box unchanged (selected, gold border); "Waiting for the site to see this copy. It answers at your next login."; the sim-done line |

Facts: no layout shift between the copy states; the date form appears once the ack is older than yesterday. Still unverified by a human pointer: the Use button's hover and pressed looks. Not captured: drag/reset, minimap/compartment/keybind.

# Build dialog round 7 (2026-10-12, L3.4 fix round 6 at 4e2ab5c4), the level-21 warrior; fixture as round 4

| File | What it shows |
|---|---|
| dialog7-1-row.jpg | `/fs`: the tab strip "Paste a build  1 waiting  Copy my character" with the badge muted in its own slot; the row "Arms raid (old patch) · 33/16/2 · from the site" with "from the site" muted, starting at the content margin |
| dialog7-2-row-preview.jpg | The row selected (gold edge) feeding the preview; "Replaces Arms raid · same points and gear · nothing changes"; the Next-would-be line; the corner hints |
| dialog7-5-copy.jpg | Copy mode: the tab strip at the same positions as in paste mode; "The site has your character from 10 Oct 20:00. Yours is newer."; "Sim done from an earlier copy · a new sim is queued" (the fixture's ack is queued while weights are in use); "Your next step: Take Deflection (1 of 2)" |
| dialog7-tabs-paste.jpg | Back in paste mode: the strip unchanged (no jump) |
| dialog7-escape-empty.jpg | Escape with a row previewing: back to the empty state with the row listed; a second Escape closes |

Facts: the personas' one shared must-fix (the tab strip) is fixed in pixels; everything else from the final passes is on screen. Owed to L9.1: a human hover/press on Use and on a row, the no-ack lines, the other Replaces variants, the newer-patch refusal, drag/reset, minimap/compartment/keybind, "Best next upgrade" on a character whose next step is an item.

# Build dialog round 8 (2026-10-12, L3.4 fix round 7 at d6b6f01f)

| File | What it shows |
|---|---|
| dialog8-5-sim-line.jpg | Copy mode's status block: "The site has your character from 10 Oct 20:00. Yours is newer."; "Sim done · your sim weights are in use" (no copy named, no queued claim); "Your next step: Take Deflection (1 of 2)" |

Facts: the first reload after the commit did not take (the old line showed once); the second reload shows the new line. L3.4 is closed.

# Settings page round 1 (2026-10-12, L3.5 at c23c8c03), the level-21 warrior

| File | What it shows |
|---|---|
| settings-1-screen.jpg | Right-clicking the Forever Sixty minimap button opens the client's Options window on AddOns › Forever Sixty (whole screen at 0.6) |
| settings-1-page.jpg | The page: "Forever Sixty" in gold; six rows with gold ticked squares, body labels and one muted hint each (Show the chip; Hide the chip in combat; Show the minimap button; Gear tips on item tooltips; Share my build with my guild; Save my character for the companion, the small app on your PC); "Commands" with six muted lines; "Reset positions" with its sentence; "Data for this client: current". The whole page fits the canvas with room below |
| settings-6-minimap-tooltip.jpg | The minimap button's tooltip: "Forever Sixty / Arms raid / 0 of 51 points / Deflection (Arms, tier 1) / Left-click: paste or copy a build / Right-click: settings" (the gold "(1x)"-style button at the minimap's lower left; the blue one above it is RestedXP's) |
| settings-4-chip-off-switch.jpg | "Show the chip" clicked off: an empty 1px-bordered square |
| settings-4-chip-gone.jpg | The chip gone from beside the minimap at once |
| settings-4-chip-back.jpg | Clicked on: the chip back ("Take Deflection (1 of 2)") |
| settings-10-help.jpg | `/fs help`: the six command lines |

Facts: no scroll needed on this 1456×819 client; hover and pressed looks on the rows are not demonstrable with the capture tool's pointer (same limit as the dialog). Not captured: Reset positions moving all three surfaces, the compartment entry, the "older patch" data line, the other four switches' effects.

# Settings page round 2 (2026-10-12, L3.5 fix 1 at 39ac50a6, merged tip fe560357), the level-21 warrior

| File | What it shows |
|---|---|
| settings2-1-screen.jpg | Right-click on the minimap button: Options › AddOns › Forever Sixty. The page sits on its own dark opaque panel inside the canvas; nothing of the world shows through (whole screen at 0.6) |
| settings2-1-page-top.jpg | The page at the top: gold "Forever Sixty"; seven rows (Show the chip; Hide the chip in combat; Marks on the talent window; Show the minimap button; Gear tips on item tooltips; Share my build with my guild; Save my character for the companion app), each with its hint in muted grey; Reset positions with its sentence; a gold "Commands" heading over a 1 px line; the seven `/fs …:` lines; the foot's first line cut by the panel edge |
| settings2-3-page-scrolled.jpg | The page after one wheel scroll (about 70 px): the foot reads "Addon data: current for this patch" (muted) and "Ratings data: too old. Update it from the site." (label muted, state in the warning red; this beta character's ForeverSixtyData is old). Reset stays visible before and after the scroll |
| settings2-2-minimap-hint-glyph.jpg | The minimap-button hint at full scale: "…Off: type /fs options or use Options · AddOns." — the "›" draws as a small mid-height mark in this client's font, closer to a middle dot than a chevron |
| settings2-4-off-switch.jpg | "Gear tips" clicked off: an empty square with a grey 1 px outline (round 1's dark-blue border is gone) |
| settings2-5-chip-hidden.jpg / settings2-5-switch-reads-off.jpg / settings2-5-chip-back.jpg | Right-click hides the chip; the page then shows "Show the chip" off; clicking it on brings "Take Deflection (1 of 2)" back at once |
| settings2-6-marks-switch-off.jpg / settings2-6-talents-no-marks.jpg | "Marks on the talent window" off: the talent window opens with no ring, no number badges and no strip; the chip stays beside the minimap instead of hiding |
| settings2-6-talents-marks-back.jpg | The switch on again: badges, ring and the strip ("Next: Deflection (1 of 2) · Arms raid · 0 of 51 points · spend 12 now · 0 off plan") are back on the next open |
| settings2-7-button-moved.jpg / settings2-7-reset-pressed.jpg / settings2-7-button-back.jpg / settings2-7-reset-after.jpg | The minimap button dragged to the ring's upper left; Reset positions clicked: the label reads "Positions reset" greyed, the button is back at its default spot on the ring, and after about 2 s the label reads "Reset positions" again |
| settings2-8-fs-options.jpg | `/fs options` with the Options window closed opens the page (keeping the previous scroll position) |
| settings2-8-fs-help.jpg | `/fs help`: the seven player lines in body colour, then the muted "For testers: /fs diag verbose, /fs inbox, /fs hello, /fs seen, /fs record on, off, status or clear"; the client wrapped that line after the "/" of "/fs seen" |
| settings2-9-older-patch-page.jpg / settings2-9-older-patch-foot.jpg | Data.lua's build set to 1.59.0.1 and reloaded: "Addon data:" muted, "an older patch. Update the addon." in the warning red, above the ratings line (reverted after the capture) |
| settings2-10-share-off.jpg / settings2-10-hello-refused.jpg | "Share my build with my guild" off, then `/fs hello`: "Not sent: "Share my build with my guild" is off in the settings." (restored to on afterwards) |

Facts: the panel needs one wheel tick to reach the foot on this 1456×819 client (Reset never leaves view); the chip could not be shift-dragged with the capture tool (shift cannot be held through a synthetic drag), so Reset was exercised on the minimap button only, the chip and the dialog resets are pinned by specs; hover and pressed looks remain undemonstrable with the synthetic pointer; the hello refusal line appeared once and scrolled under the help output in the first attempt, so it was retaken alone.

# Settings page round 3 (2026-10-12, L3.5 fix 2 at c42b375a; client loads the merged tip fb996995), the level-21 warrior

| File | What it shows |
|---|---|
| settings3-1-screen.jpg / settings3-1-page-top.jpg | The page after `/reload`: "Forever Sixty", then the warning line "Player ratings: too old. Get the new Forever Sixty Data from the site." in the warning red under the title (this character's data addon is old), the seven rows with the round-2 copy: "…Off: type /fs options or use Options > AddOns." (ASCII, draws correctly), the Share hint on one line "Guildmates with Forever Sixty see your level, spec, build and gear gap. Off: they see nothing.", the Save hint on two lines "…to pick up, so the site can read your gear and talents." with a full row gap before Reset; Reset; Commands |
| settings3-4-reset-pressed.jpg / settings3-4-reset-after.jpg | Reset clicked: "Positions reset" in the muted grey with a 1 px grey border, readable; after about 2 s "Reset positions" in gold again |
| settings3-3-page-scrolled.jpg | One wheel tick down: the foot's first line "Addon data: current for this patch" at the bottom edge |
| settings3-3-foot-max-scroll.jpg | Ten wheel ticks down (the scroll's end): the page stops at "Addon data: current for this patch"; the "Player ratings: …" foot line and the 12 px pad are below the reachable range (not visible at any scroll position while the title warning line is shown) |
| settings3-5-title-warning.jpg | Scrolled back to the top: the title and the warning line at full scale |
| settings3-6-talents-marks-on-screen.jpg / settings3-6-chip-hidden-under-talents.jpg / settings3-6-chip-back-after-close.jpg | Marks on, talent window opened with the N keybind: the strip reads "Next: Deflection (1 of 2) · Arms raid · 0 of 51 points · spend 12 now · 0 off plan" and the chip is NOT beside the minimap (round 2's note that it stayed was wrong); closing the window brings "Take Deflection (1 of 2)" back |
| settings3-8-fs-help.jpg | `/fs help`: the seven player lines, then the muted "For testers: /fs diag verbose, inbox, hello, seen, record." wrapping once at a space ("…seen," / "record.") |

Facts: the title warning line is the whole sentence in the warning colour (label included); the foot's last line is cut off by the scroll range while the warning line shows (a defect for fix round 3); the addon-data "older patch" title line was not re-shot with the 1.59 trick (same mechanism as the ratings line shown here); the chip's shift-drag and hover/pressed states remain undemonstrable with the capture tool.

# Settings page round 4 (2026-10-12, L3.5 fix 3 at 6f989280), the level-21 warrior

| File | What it shows |
|---|---|
| settings4-1-page-top.jpg | After `/reload`: "Forever Sixty", the warning line "Player ratings (on player tooltips): too old. Get the new Forever Sixty Data at foreversixty.gg." in red, the seven rows, the Save hint on ONE line ("Written at logout for the companion app on your PC, so the site can read your gear and talents."), Reset, Commands |
| settings4-2-foot-reached.jpg | Five single wheel notches (one per action): the thumb at the bottom of its track, the foot's two lines fully visible — "Addon data: current for this patch" and "Player ratings (on player tooltips): too old. …" — with panel below them and the Close button under the panel. Round 3's "unreachable" was the capture tool sending its multi-tick scroll as one notch |
| settings4-2-scroll-diagnostic.jpg | The report's `/run` check printed "585 646.55554199219 61.555526733398": clip 585, page 646.56, offset 61.56 (clip + offset = page at the end of the scroll) |
| settings4-3-reopened-at-top.jpg | Closing and reopening the page: it opens at the top again (title and warning line in view) |
| settings4-4-share-off.jpg / settings4-4-hello-refused.jpg | Share off, then `/fs hello`: "Not sent: "Share my build with my guild" is off. Type /fs options to turn it on." (restored to on afterwards) |
| settings4-5-fs-help.jpg | `/fs help`: the seven player lines, then two muted lines "For testers: /fs diag verbose" and "Also /fs inbox, /fs hello, /fs seen, /fs record" |
| settings4-6-addon-data-warning-title.jpg | Data.lua's build set to 1.59.0.1 and reloaded: the title line reads "Addon data: an older patch. Update the addon at foreversixty.gg." (addon data wins over ratings when both are stale); reverted after the capture |
| settings4-6-both-warnings-foot.jpg | Same state at the end of the scroll: both foot lines in red, reachable, with the pad and Close below |

Facts: "not installed" foot-only is spec-only on this client (the data addon is installed and old); hover/pressed looks and the chip's shift-drag remain undemonstrable with the synthetic pointer.

# Hunter spot check (2026-10-13, client restarted on the user's level-24 hunter "Bow Jackzon", Stonetalon; addon at d3bc356c)

| File | What it shows |
|---|---|
| hunter-1-screen.jpg / hunter-1-chip-icon-only.jpg / hunter-1-chip-tooltip.jpg | The chip beside the minimap is an icon with NO text line; its tooltip reads "Forever Sixty / Click: Open talents / Right-click: Hide until next level / Shift-drag: Move" (no state line, no build name) |
| hunter-2-diag.jpg | `/fs diag`: "Not working:" · "The Forever Sixty Data addon is too old. Update it from the site." · "The paper-doll and bag marks failed once and turned themselves off: Interface/AddOns/ForeverSixty/surfaces/SlotMarks.lua:101: bad argument #2 to 'tonumber' (number expected, got string)" · "Chip: shown, nothing cut off" · "Bag marks: off in Baganator's bags. The marks on your character panel still show." |
| hunter-3-bag-probe.jpg | The implementer's `/dump` probe: hooksecurefunc true; `ContainerFrame_Update` FALSE (absent on 1.60); `ContainerFrameItemButtonMixin.UpdateNewItem` true; `ContainerFrame1.UpdateItems` true; `ContainerFrame1.Items` true |
| hunter-4-talents-no-build.jpg | The talent window: 0 unspent, points spread over Beast Mastery, Marksmanship and Survival; no badges, ring or strip (the character has no build) |
| hunter-5-paperdoll-no-marks.jpg | The character pane: no arrow or target corner on any slot (the marks disabled themselves after the crash; the back slot's tooltip says "Best until level 30", which by the implementer's rule draws no corner anyway) |
| hunter-5-back-tooltip.jpg | Hovering the worn Grave Shroud (Back): "Forever Sixty · Estimate   Equipped / Best until level 30 / At level 30: Tigerstrike Mantle / · World Drop" — the tooltip works on a hunter (band 20, Marksmanship pick) |

Facts: the Lua error window that appeared once came from the controller's `/run ToggleCharacter("PaperDollFrame")` (tainted execution in Blizzard_TextStatusBar), not from the addon; the `c` keybind did not open the pane on this client. Defects for the next rounds: the SlotMarks crash (L4.1 fix 1); the empty chip line on a character with no build and no unspent points (L3.1 carry, folded into L4.1 fix 1 if the chip model owns it).
| hunter-6-empty-trinket-hover.jpg | Hovering the hunter's empty trinket slots (both empty): the client's own "Trinket" tooltip and NO addon section — either band 20 has no hunter trinket pick (then the empty-slot section has nothing to say, and the brief did not define that case) or `emptySlotSection` did not fire; the reviewer decides from the code, the fix round pins the no-pick case ("No pick for your band yet" or silence, by ruling) |
| hunter-7-lua-error-on-c.jpg / hunter-7-lua-error-show-stack.jpg / hunter-7-lua-error-hide-stack.jpg | Pressing the client's own `c` keybind (no addon command involved): the Lua error window, four errors in the session — on every Show AND Hide of the character pane: "Interface/AddOns/Blizzard_TextStatusBar/TextStatusBar.lua:110: attempt to compare a secret number value (execution tainted by '***ForceTaint_Strong***')", stack UpdateTextStringWithValues ← UpdateTextString ← ShowStatusBarText/HideStatusBarText ← CharacterFrame.lua:747/792 ← Show/Hide ← SetUIPanel ← ShowUIPanel/HideUIPanel ← SetAttribute ← ToggleCharacter. The user reports it on every open. Earlier rounds (through 6f989280, before the slot marks) opened the paper doll with scriptErrors on and no error, so the taint is new with L4.1's hooks on PaperDollFrame/CharacterFrame OnShow and the textures created on the client's slot buttons |
| hunter-8-pane-open-no-error.jpg / hunter-8-pane-and-bags-no-error.jpg / hunter-8-paperdoll-target-corners.jpg | After `/reload` with the overlay fix (758f0d68) and the user rebinding the pane to `c`: seven opens/closes of the character pane plus the bags with scriptErrors on → NO Lua error window (the taint is gone). The paper doll now shows the 12 px gold target corner at the bottom-right of the Waist slot and of the Ranged weapon slot (the slots whose hover prints "Next upgrade"); no arrows (no bag item beats a worn piece) |
| hunter-9-right-column-corners.jpg | The right column at full scale: gold target corners at the bottom-right of the Hands slot and the first Finger slot (plus the Ranged slot below) |
| hunter-9-hands-corner-tooltip.jpg | Hovering the Hands slot (Gloves of the Fang, a corner on it): "Forever Sixty · Beast Mastery · Estimate   Equipped / Best until level 30 / At level 30: Pilferer's Gloves · Crafted / Planned for Beast Mastery at 20" — the corner sits on a slot whose hover says the worn item is best for the band (the corner follows the At-30 target, not a "Next upgrade" line) |
| hunter-9-finger-corner-tooltip.jpg | Hovering the first Finger slot (Ancient Heirloom, a corner): "Best until level 30 / At level 30: Ironspine's Eye +1.1 DPS · Scarlet Monastery" — same mismatch |
| hunter-9-mainhand-bare-equipped.jpg | Hovering the main hand (Blackwater Cutlass): the addon section is a bare "Forever Sixty · Estimate   Equipped" with no verdict, no Best-until and no Next line (the silence the personas rejected in L3.3) |
| hunter-10-gloves-in-bag-screen.jpg / hunter-10-bag-gloves-tooltip.jpg | The worn Gloves of the Fang dragged into an empty bag cell: the bag item's tooltip reads "Forever Sixty · Beast Mastery · Estimate / Fills an empty slot, +0.8 DPS / (+4 Str, +6 Agi) / Planned for Beast Mastery at 20 / Wailing Caverns · Druid of the Fang · 2% drop" |
| hunter-10-paperdoll-empty-hands-no-arrow.jpg | The paper doll with the Hands slot empty: NO arrow on the empty slot (the +0.8 DPS filler is under the chip's 1.0 floor, so the chip and the marks ignore it) while the Waist/Finger/Ranged corners stay; Baganator's bag cell shows no mark (as designed) |
| hunter-10-reequipped.jpg | Right-click on the bag item re-equips it; the pane back to its previous state, no error |

# Hunter after the L3.8 merge and L4.1 fix 3 (2026-10-13, client at 4f30639c)

| File | What it shows |
|---|---|
| hunter-11-talents-saved-build.jpg | The talent window: the saved 31/20 build is followed — numbered badges on Beast Mastery and Marksmanship, the ring on the next talent, the strip "Next: Lethal Attacks (1 of 5) · 12 of 51 points · 3 off plan"; "Unspent Talents 0" |
| hunter-11-chip-tooltip-build.jpg | The chip's tooltip now leads with "Forever Sixty · Beast Mastery / Marksmanship build / Next talent at 25: Lethal Attacks", then the click/right-click/shift-drag lines |
| hunter-11-diag-build-line.jpg | `/fs diag`: "Build: Beast Mastery / Marksmanship build (saved)" · "Chip: shown, nothing cut off" · "Talent badges: 17 drawn, 0 missing" · "Bag marks: not checked yet (open your bags)" (the data-addon-too-old line stays) |
| hunter-11-paperdoll-after-reload.jpg | The paper doll right after `/reload`: gold target corners on the Hands, Waist and Finger slots (hover check below) |
| hunter-12-hands-corner-still-best-until.jpg / hunter-12-finger-corner-still-best-until.jpg | After the fix-3 reload and with the item data long loaded: the Hands and Finger slots still carry the target corner while their hover says "Best until level 30 / At level 30: …" — the stale-data theory does not hold; the corner follows the At-line's later pick |
| hunter-12-mainhand-no-pick-line.jpg | The main hand now reads "Forever Sixty · Estimate   Equipped / No pick for your band yet / Planned for Beast Mastery / Marksma…" — the "No pick" line works; the header lacks the spec name the armour tooltips carry ("Beast Mastery · Estimate"), and the "Planned for <build name>" line is cut at the tooltip's width |
| hunter-12-diag-after-bags.jpg | `/fs diag` after opening the bags once: "Bag marks: off in Baganator's bags. The marks on your character panel still show." (the "not checked yet" line before) |
| hunter-13-lethal-attacks-badge-8-36.jpg / hunter-13-lethal-attacks-tooltip.jpg | User-reported: the Lethal Attacks badge reads "8–36" and its tooltip "Your build: 5 ranks, points 8 to 36". The saved build's order (decoded from the SavedVariables code, patch 1.60.1.70009) takes Lethal Attacks at points 8, 9, 10, 11 and 36: four ranks early, the fifth after the Beast Mastery run. The badge collapses a split run into first–last, hiding the gap; the other Marksmanship badges ("37–39", "40–44", "45–46", "48–51", "47") are consecutive runs |
| hunter-14-right-column-corners-fullres.jpg | CORRECTION of the hunter-9/12 notes: at full resolution the gold corners sit on the Waist slot (2nd) and the second Finger slot (6th), not on Hands or the first Finger; the slots I hovered earlier carried no corner (the fix-3 reviewer spotted this from the captures) |
| hunter-14-waist-next-upgrade.jpg | Hovering the Waist (a corner): "Forever Sixty · Estimate   Equipped / Next upgrade: Blackened Defias Belt +1.1 DPS / (−2 Agi, +18 Attack Power) / Planned for Beast Mastery / Marksma… / The Deadmines · Captain … · 30% drop / Also good: Deviate Scale Belt +0.4, Dark Leather Belt +0.2" — the corner rule holds |
| hunter-14-ranged-next-upgrade.jpg | Hovering the Ranged slot (a corner): "Next upgrade: Ranger Bow +3.5 DPS / (+1 Agi, +3.4 weapon DPS) / Planned for … / World drop / Also good: Lil Timmy's Peasho… +2.2, Cracked Blacksmith Hammer +0.8" with the compare tooltip titled "Next upgrade" beside it — the corner rule holds |
| hunter-15-lethal-attacks-badge-8-11.jpg / hunter-15-lethal-attacks-tooltip-runs.jpg | After the badge fix (client at eb02eb0b): the Lethal Attacks badge reads "8–11" and its tooltip "Forever Sixty: Your build: 5 ranks, points 8–11 and 36"; the other Marksmanship badges unchanged ("37–39", "40–44", "45–46", "48–51", "47") |
| hunter-15-corners-1s-after-reload.jpg / hunter-15-corners-6s-after-reload.jpg | The right column one second and six seconds after `/reload`: the gold corners on the Waist and the second Finger slot in both — nothing appears and then vanishes while the item data arrives |

# Paper-doll and bag marks, persona round 1 (2026-10-13, client at eb02eb0b: L4.1 fix 4 + L3.8 + the badge fix), the level-24 hunter following its saved build

| File | What it shows |
|---|---|
| hunter-16-gloves-in-bag-chip-equip-line.jpg | The worn Gloves of the Fang dragged into a bag cell: the chip beside the minimap now reads "Equip Gloves of the Fang, +0.8 DPS" (the empty-slot filler counts at any gain); the bag item's tooltip "Fills an empty slot, +0.8 DPS / (+4 Str, +6 Agi) / Planned for Beast Mastery at 20 / Wailing Caverns · Druid of the Fang · 2% drop" (whole screen at 0.6) |
| hunter-16-empty-hands-arrow.jpg | The right column with the Hands slot empty: the green arrow at the top-right of the empty Hands slot; the gold corners on Waist and the second Finger unchanged |
| hunter-16-empty-hands-hover-next.jpg | Hovering the empty Hands slot: "Hands / Forever Sixty · Beast Mastery · Estimate / Next upgrade: Gloves of the Fang / Wailing Caverns · Druid of the Fang · 2% drop" with the compare tooltip titled "Next upgrade" beside it |
| hunter-16-reequipped-arrow-gone.jpg | Right-click on the bag item re-equips it: the arrow is gone, the corners stay |
| hunter-16-empty-trinket-no-pick.jpg | Hovering an empty Trinket slot: "Trinket / Forever Sixty · Beast Mastery · Estimate / No pick for your band yet" (muted) |
| hunter-16-diag.jpg | `/fs diag`: "Build: Beast Mastery / Marksmanship build (saved)" · "Chip: shown, nothing cut off" · "Talent badges: 17 drawn, 0 missing" · "Bag marks: off in Baganator's bags. The marks on your character panel still show." |

Facts: Baganator's bag cells carry no mark (the hooks do not reach its buttons; the diag line says so); the corner slots' hovers print "Next upgrade" (hunter-14); the chip's line with the gloves bagged is the same list the marks read; the chip showed its line with a build in force (the earlier idle icon was the no-build state, since fixed by L3.8); hover and pressed states remain undemonstrable with the synthetic pointer; the user moved the character to Astranaar during the capture, so the client is left alone from here.
| hunter-17-finger2-corner-crop.jpg / hunter-17-finger2-hover.jpg | CORRECTION (the designer was right): the second Finger slot (Clergy Ring) has NO corner — the yellow at its lower-left is the ring's icon art — and its hover reads "Best until level 30 / At level 30: Monkey Ring · Quest". The corners are on the Waist and the Ranged slot only; every earlier note naming a Finger corner is wrong |

# Bag marks in the client's own bags (2026-10-13, Baganator disabled by the user, client at eb02eb0b)

| File | What it shows |
|---|---|
| hunter-18-blizzard-bag-screen.jpg | The Combined Backpack open beside the character pane with the gloves dragged into it: the bagged gloves' tooltip ("Fills an empty slot, +0.8 DPS …"), the chip "Equip Gloves of the Fang, +0.8 DPS" (whole screen at 0.6) |
| hunter-18-blizzard-bag-mark.jpg | The backpack at full scale: the green mark at the top-right of the bagged gloves' cell (row 2, column 2); no mark on any other item |
| hunter-18-empty-hands-arrow-fullres.jpg | The right column at full scale: the green mark at the top-right of the empty Hands slot; the gold corner on Waist; nothing on the rings |
| hunter-18-switch-off-pane.jpg / hunter-18-switch-off-bag.jpg | "Gear tips on item tooltips" switched off: the Hands slot's mark and the backpack cell's mark are gone (the Waist corner too); the chip line stays (the chip is not under that switch) |
| hunter-18-switch-on-pane.jpg / hunter-18-switch-on-bag.jpg | Switched back on: both marks and the corner are back at once |
| hunter-18-reequipped.jpg / hunter-18-chip-idle-after.jpg | Right-click on the backpack cell re-equips the gloves: no mark anywhere on Hands, the chip back to its idle icon |

Facts: the green mark is the ~6 px glyph the personas flagged (fix round 5 makes it a full 12 px arrow); opening the Options window hides the pane and the bags, so the switch captures were taken after closing it (Close, then Return to Game); the gloves drop needed the pointer on a backpack cell, not between cells.

# Paper-doll and bag marks after fix round 5 (2026-10-13, client at 84b4f966, Blizzard bags)

| File | What it shows |
|---|---|
| hunter-19-arrow-4x-still-small.jpg / hunter-19-bag-mark-after-fix5.jpg | The green mark on the empty Hands slot (4× crop) and on the bagged gloves is STILL the small ~6 px glyph, although the client has the `bags-greenarrow` atlas (hunter-19-atlas-probe: `GetAtlasInfo` answers a table — height 22, file 969828, leftTexCoord 0.1289, bottomTexCoord 0.4492) |
| hunter-19-in-your-bags-hover.jpg | Hovering the empty Hands slot with the gloves bagged: "Hands / Forever Sixty · Beast Mastery · Estimate / In your bags / Gloves of the Fang +0.8 DPS / Right-click it in your bags to equip." with the compare tooltip showing the gloves — as ruled |
| hunter-19-waist-hover-equipped-collision.jpg | Waist: "Forever Sixty · Beast Mastery · Estimate / Next upgrade: Blackene…   Equipped / +1.1 DPS / (−2 Agi, +18 Attack Power) / Planned for your build / The Deadmines · Captain Greenskin / Also good: Deviate Scale Belt +0.4, Dark Leather Belt +0.2" — the header reorder pushed "Equipped" onto the Next line's row, truncating the pick's name and splitting the gain onto its own line; the boss is whole and the chance dropped as ruled; "Planned for your build" fallback as ruled |
| hunter-19-ranged-hover-equipped-collision.jpg | Ranged: the same collision ("Next upgrade: Ranger Bo…   Equipped"); "World drop · check the AH" and "Lil Ti… +2.2 (needs bullets)" as ruled (the alternative's name cut by the same squeeze) |
| hunter-19-mainhand-header-fits.jpg | The main hand (now Thief's Blade): the header with the spec word and "Equipped" on one row where it fits; "Best until level 30 / At level 30: Alliance Outrunner's Sword · Thora Feathermoon · World…" (the At line now carries the source) ; "Planned for Beast Mastery / Marksmanship build" fits |
| hunter-19-chip-legend.jpg | The chip tooltip: "Forever Sixty · Beast Mastery / Marksmanship build / Next talent at 25: Lethal Attacks / Green arrow: a bag item beats what you wear. Gold corner: next upgrade to chase." |
| hunter-19-diag-working-heading.jpg | `/fs diag`: "Not working:" (the data addon line) then "Working:" with the Build/Chip/… lines |

## hunter-20 — slot marks fix round 6 (1d1becfe), Bow Jackzon L24, Astranaar, Blizzard bags (Baganator off)

- `hunter-20-hands-arrow-4x.jpg` — 4× crop of the Hands slot after dragging the worn gloves to the backpack: the addon-drawn 12 px green arrow with a dark outline at the slot's top-right. Reads as an arrow now, not a dot.
- `hunter-20-bag-arrow.jpg` — the Combined Backpack: the same arrow on the gloves' bag cell (row 2, cell 2).
- `hunter-20-waist-tooltip.jpg` — Waist hover: header "Forever Sixty · Beast Mastery · Estimate", then "Equipped" right-aligned on its own row, then "Next upgrade: Blackened Defias Belt" whole on one row and "+1.1 DPS" on the next row (never-widen rule; the tooltip's width is the client's "Press F6…" line), "(−2 Agi, +18 Attack Power)", "Planned for your build", "The Deadmines · Captain Greenskin", "Also good: Deviate Scale Belt +0.4, Dark Leather Belt +0.2".
- `hunter-20-ranged-tooltip.jpg` — Ranged hover (with the compare tooltip for Ranger Bow): "Equipped" own row; "Next upgrade: Ranger Bow +3.5 DPS" fits one row; "(+1 Agi, +3.4 weapon DPS)", "World drop · check the AH", "Also good: Lil Ti… +2.2 (needs bullets), Cracked Blacksmith Hammer +0.8".
- `hunter-20-chip-legend.jpg` — chip tooltip: "Next talent at 25: Lethal Attacks" then the two legend lines "Green arrow: a bag item beats what you wear." / "Gold corner: next upgrade to chase." then the click hints.
- The `/run` overlay probe printed nothing visible in the chat log at the time of the capture (the chat had scrolled to live General traffic); the geometry is pinned by the spec instead.

## hunter-21 — chip tooltip with "Gear tips on item tooltips" off (designer round-2 must-have 2)

- `hunter-21-chip-legend-off.jpg` — the switch unticked in Options → AddOns → Forever Sixty: the chip tooltip reads "Forever Sixty · Beast Mastery / Marksmanship build", "Next talent at 25: Lethal Attacks", one blank row, then the three click hints. No legend lines and no gap where they were. Switch re-ticked afterwards (verified in the Options zoom).

## hunter-22 — slot marks fix round 7 (36410dc3), after /reload

- `hunter-22-ranged-also-good-rows.jpg` — Ranged hover: "Also good:" alone on its row, then "Lil Timmy's Peashooter +2.2" (whole name, quality blue, gain muted) with "(needs bullets)" on a muted row under it, then "Cracked Blacksmith Hammer +0.8". No name cut, no trailing commas. The rest unchanged ("Equipped" own row, "Next upgrade: Ranger Bow +3.5 DPS", "World drop · check the AH").
- `hunter-22-waist-also-good-rows.jpg` — Waist hover: "Also good:" / "Deviate Scale Belt +0.4" / "Dark Leather Belt +0.2" each on its own row (the pair is wider than the tooltip). This time the client anchored the tooltip lower so nothing is clipped at the screen top; the compare tooltip (Blackened Defias Belt) sits beside it whole.
- `hunter-22-chip-legend-colours.jpg` — chip tooltip: "Green arrow" in green and "Gold corner" in gold, the rest of each legend line muted; the hints unchanged.

## hunter-23 — character panel tab round 1 (a53a3005), Bow Jackzon L24, after /reload

- `hunter-23-panel-button-screen.jpg` / `-crop.jpg` — the pane open; `/dump` shows `CharacterFrameTab1` absent and `CharacterFrame.numTabs` nil on this client (no bottom tabs), `ForeverSixtyPanelTab` present, the flyout nil before the first click. The button is a small dark plate reading "Forever Sixty" just under the pane's bottom-right edge (the fallback anchor), not styled like the pane's chrome.
- `hunter-23-flyout-screen.jpg` / `-crop.jpg` — after clicking: a narrow flyout at the pane's top-right titled "Forever Sixty" with an "x". Regions: Next targets (Ranger Bow +3.5 DPS, "Ranged · World drop · check the AH"; Blackened Defias Belt +1.1 DPS, "Waist · The Deadmines · Captain Greenskin · 30% drop" wrapping to a second line), Rotation (Aimed Shot Rank 1, Multi-Shot, Serpent Sting Rank 3, Arcane Shot Rank 3, each with a two-line plain condition ending in an ellipsis where cut), Build ("Beast Mastery 31 / Marksmanship 20", "12 of 51 points on plan", "3 off plan" in red, buttons "Use a different build" / "Copy my character"). Defects visible: the item and ability icons draw far larger than their rows (≈40 px) and sit over the region titles and neighbouring rows ("Next targets", "Rotation" and "Build" titles are covered by icons; the target names start under the icon); the flyout is narrow (~220 px) so the source line wraps and overlaps the next row; the rotation icons overlap the rows above them.
- `hunter-23-diag-panel.jpg` — `/fs diag`: "Panel: open, 2 targets, rotation 4 rows" under "Working:". Also visible: "Not working: The Forever Sixty Data addon is too old. Update it from the site." (the client's ForeverSixtyData nightly addon is behind the shipped Data.lua; not an L4.3 matter, carried).

## hunter-24 — character panel tab fix round 2 (970f2e14), Bow Jackzon L24, after /reload

- `hunter-24-panel-button-screen.jpg` / `-crop.jpg` — the pane open (no Lua error): the "Forever Sixty" button under the pane's bottom-right edge, gold label and 1 px gold edge over the panel fill, 112 × 24 (probe: left 535, top 600).
- `hunter-24-flyout-screen.jpg` / `-crop.jpg` — after one click: the 320 × 414 flyout opens right of the pane's vertical mode-tab column (flyout left 711 = `CharacterFrameModeTabs` right 711). Regions: "Next targets" (Ranger Bow +3.5 DPS, "Ranged · World drop · check the AH"; Blackened Defias Belt +1.1 DPS, "Waist · The Deadmines · Captain Greenskin · 30% drop", each with its 16 px icon), "Rotation" (1. Aimed Shot Rank 1, 2. Multi-Shot, 3. Serpent Sting Rank 3, 4. Arcane Shot Rank 3, each with a two-line plain condition cut with an ellipsis where needed; then "Rapid Fire at level 26" muted, last), "Build" ("Beast Mastery 31 / Marksmanship 20", "12 of 51 points on plan", "3 off plan" in red, the two buttons). No icon overlaps, every line on its own row.
- `hunter-24-flyout-row-hover.jpg` — hovering the Ranger Bow row shows the item tooltip with its Forever Sixty section ("Upgrade +3.5 DPS over Short Ash Bow", "(+1 Agi, +3.4 weapon DPS)", "Planned for Beast Mastery at 20", "World drop · check the AH") and the client's compare tooltip for the worn Short Ash Bow.
- `hunter-24-flyout-after-reopen.jpg` — the pane closed and reopened with the flyout open: the flyout is still open with current content (the stale-redraw path), no error.
- `hunter-24-diag-panel.jpg` — `/fs diag`: "Panel: open, 2 targets, rotation 5 rows".

## hunter-25 — panel tab capture-only asks from the personas' round 1 (still 970f2e14)

- `hunter-25-font-scale-probe.jpg` — probe: `GameFontNormalSmall` is 10 px; `UIParent` and the flyout both at effective scale 0.64; the screen is 2133 × 1200 UI units (a 16 px icon is ~10 physical px at this scale; the capture tool's frame is 1456 × 819, so no crop can be 1:1 — the zooms are upscaled from the real pixels).
- `hunter-25-build-paste-dialog-screen.jpg` / `-crop.jpg` — "Use a different build" opens the build dialog on its "Paste a build" tab ("Following: Beast Mastery / Marksmanship build · Hunter · 51 po…", "Next talent: Take Lethal Attacks (Marksmanship, 1 of 5)", the paste box, "Tab to switch"); the dialog layers above the flyout; the flyout stays open.
- `hunter-25-copy-dialog-screen.jpg` / `-crop.jpg` — "Copy my character" opens the dialog on "Copy my character" with the FS1 export code selected ("Selected. Press Ctrl+C, then paste it on foreversixty.gg.", "No companion found. …"); the flyout stays open behind it.
- `hunter-25-button-closed-flyout.jpg` — a second click on the Forever Sixty button closes the flyout; the button stays under the pane.

## hunter-26 — character panel tab fix round 3 (04c42093), after /reload

- `hunter-26-button-hover-hint-screen.jpg` / `-crop.jpg` — hovering the button: a tooltip "Forever Sixty" (gold) / "Your next targets, rotation and build." / "Click to open or close." (muted); the button brightens.
- `hunter-26-button-open-state.jpg` — with the flyout open the label is in the body white on the inset ground (closed: gold label).
- `hunter-26-flyout-crop.jpg` — the flyout: "Next targets" (the two rows as before) then the muted "Other slots are best for your level."; a 1 px rule above "Rotation" and above "Build"; Arcane Shot's condition now cuts at a clause ("…will not be ready within a global…"); "Rapid Fire at level 26" and its condition dimmed; Build reads "Plan: Beast Mastery 31 / Marksmanship 20", "12 of 15 points on plan", "3 off plan" (red).
- `hunter-26-offplan-hover.jpg` — hovering "3 off plan": a tooltip "3 off plan" (red) / "Pathfinding · Beast Mastery · Rank 2/2" / "Improved Revive Pet · Beast Mastery · Rank 1/2" (3 points).
- `hunter-26-rotation-hover.jpg` — hovering the Arcane Shot row: the row's hover ground and a tooltip "Arcane Shot" with the whole condition wrapped ("…mana is above a 25% reserve: it is instant, so it never clips an auto shot … runs dry without the floor.").
- `hunter-26-target-row-hover.jpg` — hovering the Ranger Bow row: the hover ground on the row, the item tooltip with its Forever Sixty section and the compare tooltip. The "· over Short Ash Bow" suffix did not appear on this row's source line ("Ranged · World drop · check the AH") — the fit rule dropped it; the personas may rule on whether the suffix should win over "check the AH".
- `hunter-26-rapid-fire-hover.jpg` — hovering the dim "Rapid Fire at level 26" row: the row's hover ground and a tooltip "Rapid Fire" (gold) / the full condition wrapped ("Rapid Fire timed to land with an incoming auto shot just as Aimed Shot is about to come off cooldown, so both benefit.") / "Learned at level 26".

## hunter-27 — player tooltip score and chat badges (L4.4 merged at 8ca99638), the ForeverSixtyData-l44 fixture junctioned in, after /reload

- `hunter-27-diag-fixture.jpg` — `/fs diag`: "Everything is working." (the "data addon too old" line gone with the fixture dated today).
- `hunter-27-own-tooltip-score.jpg` — hovering the player frame: after the client's lines, "Forever Score 1240 · Rare · as of 13 Oct" in the rare blue, then the muted "Shift: score details · /fs link: site page".
- (Superseded by the round-4 correction under hunter-28: the unlogged-player capture is `hunter-28-unlogged-game-shot.jpg`.) `hunter-27-unlogged-tooltip.jpg` — hovering a nearby player not in the record (Jet Dankus, level 19 rogue): the muted "Forever Score: no logs yet".
- `hunter-27-shift-components.jpg` — with Shift reported held (a temporary `/run` override of `IsShiftKeyDown`, restored right after): the score line then "Gear 380 / Experience 410 / Parses 300 / Attendance 100 / Discipline 50" (sum 1240), the hint line gone.
- `hunter-27-guild-badge.jpg` — one guild line from the hunter ("testing an addon, ignore me"): the chat reads "[Guild] [Bow Jackzon]: [1240] testing an addon, ignore me" with the badge in the rare blue before the message; the name link untouched.
- `hunter-27-fs-link-screen.jpg` (+ crop) — Jet Dankus targeted, `/fs link`: the build dialog on "Copy my character" reads "Jet Dankus's page on foreversixty.gg. Press Ctrl+C, then paste it in your browser." with `foreversixty.gg/character/us/pvp/jet-dankus` selected.
- `hunter-27-settings-eight-switches.jpg` / `hunter-27-settings-page-crop.jpg` — the Settings page with the eighth switch "Score badges in chat" ("Shows a player's Forever Score before their message in LFG, Trade, guild chat and whispers.") between "Gear tips on item tooltips" and "Share my build with my guild"; the command list gains "/fs link: copy a player's site page".
- `hunter-27-guild-badge-off.jpg` — the switch off, a second guild line "(addon test 2, badges off)": no badge; the earlier line keeps its "[1240]". Switch re-ticked afterwards.
- Not captured: a guildmate with the addon (the live line) and Trade with two badged names — no guildmate with Forever Sixty was online and the hunter is not in a city; Smoke Gouda (fixture, 1610 epic) was out of `/target` range. The guild-chat badge on the hunter's own line covers the badge path; the whisper and Trade paths are spec-covered.

## hunter-28 — player tooltip and badges fix round 2 (1fde8b18), game screenshots (`/run C_Timer.After(2, function() Screenshot() end)`, no OS overlay; 3840 × 2160)

- `hunter-28-own-tooltip-game-shot.jpg` — the player frame hovered: after the client's lines, "Forever Score 1240 · Rare · as of 13 Oct" in the rare blue, then the muted "Shift: score details · target them, then /fs link".
- `hunter-28-shift-real-game-shot.jpg` — the real Shift path: Shift held with the physical key (hold_key) while hovering, the screenshot fired during the hold: the tooltip unmoved and without duplicated lines; the score line, then five double lines "Gear 380 / Experience 410 / Parses 300 / Attendance 100 / Discipline 50" (labels muted left, values right-aligned) and the muted right-aligned "= 1240"; the hint line gone.
- (Superseded: this file was replaced in round 3 and again in round 4, see the round-4 correction below; the notes here describe the first version.) `hunter-28-unlogged-game-shot.jpg` — a passing player not in the record (a level 18 dwarf shaman): "Forever Score: no logs yet" and the muted "Target them, then /fs link: site page".
- `hunter-28-guild-badge-fs.jpg` — a third guild line "(addon test 3)": "[Guild] [Bow Jackzon]: FS 1240 · (addon test 3)" — "FS 1240" in the rare blue, a muted "·", no brackets.
- Still not captured: the live line (no guildmate with the addon online), Trade/LFG and an incoming whisper, the epic colour (Smoke Gouda out of range).
- Round-3 re-shoots (after the total-line fix c90f8700, the empty left text replaced by a space): `hunter-28-shift-real-game-shot.jpg` replaced — Shift held with the physical key, the timed screenshot fired during the hold: five double lines then "= 1240" on its own row, one line pitch below "Discipline 50", no overlap, the tooltip unmoved. (Superseded by the round-4 correction below:) `hunter-28-unlogged-game-shot.jpg` replaced (the earlier file was a duplicate of the Shift shot because its timer had not fired) — the targeted level 18 dwarf shaman "Olympus XXVI": "Forever Score: no logs yet" and the muted "Target them, then /fs link: site page". `hunter-28-settings-gear-tips-hint.jpg` — the Settings row "Gear tips on item tooltips" with the hint ending "Off: also hides the score on player tooltips." wrapped to a second line, the "Score badges in chat" row under it.
- Round-4 correction: `hunter-28-unlogged-game-shot.jpg` replaced again (the round-3 file had caught the tooltip fading; the target frame on this UI draws no unit tooltip, so the world model was hovered instead). Verified by eye: Jet Dankus (level 19 night elf rogue, Knights Of Olympus), then "Forever Score: no logs yet" and the muted "Target them, then /fs link: site page" at bottom right; no OS overlay.

## hunter-29 — guild tab round 1 (ba9c39d7), Bow Jackzon (member of "Employed Millennials", 97 members / 3 online), the real ForeverSixtyData (too old → every row "no logs")

- `hunter-29-guild-frame-button.jpg` — the Guild & Communities frame open (`j`): the addon's gold-edged "Forever Sixty" button under the frame's bottom-right edge, beneath the client's red "Invite Member".
- `hunter-29-guild-flyout-screen.jpg` / `-online-crop.jpg` — after one click: the 480-wide flyout "Forever Sixty · Guild" opens right of the frame's side-tab column; "Online only" ticked; "3 of 97 shown · 0 live"; headers Name · Level · Spec · Gear · Score ▾ · Attendance · Gap (Score carries the sort arrow); rows Bow Jackzon 24, Onearm Lopen 30, Sanna Widebraid 9 in class colours, "no logs" muted, the other cells "-".
- `hunter-29-diag-guild.jpg` — `/fs diag`: "Guild tab: open, 3 rows (0 live), officer no".
- `hunter-29-guild-roster-all.jpg` — "Online only" unticked: "97 of 97 shown · 0 live", 20 rows visible with a scrollbar; online members bright, offline dimmed; long names cut with an ellipsis ("Bloodveil Whisperb…").
- `hunter-29-guild-roster-scrolled.jpg` — after one wheel notch: the list moved three rows.
- `hunter-29-guild-roster-sorted-level.jpg` — the Level header clicked: sorted by level descending (30s first), the arrow now on Level, the sort header gold.
- Not yet captured: rows with scores (the real data addon is behind; the L4.4 fixture is being junctioned in for a second pass), a live row (no guildmate with the addon online), the officer region and the push confirm (need the user's guild-master character).
- `hunter-29-guild-flyout-scored-crop.jpg` — the L4.4 data fixture junctioned in and `/reload`: the hunter's own row now reads "Bow Jackzon · 24 · Marksmanship · 380 · 1240 (rare blue) · 50% · -"; the two other online members stay "no logs". (The fixture's spec word comes from the record, not the live build.) Junction restored to the real data addon afterwards.

## hunter-30 — guild tab fix round 1 (d2c901d1), the real ForeverSixtyData (too old), after /reload

- `hunter-30-guild-flyout-screen.jpg` / `-default-crop.jpg` — the flyout now 4 units clear of the side tabs; count line "3 of 97 shown · 1 of 3 online run Forever Sixty"; one muted status line "Player ratings: too old. Get the new Forever Sixty Data at foreversixty.gg."; headers right-aligned over the numeric columns, the sort arrow (client art) on Score; the hunter's own row live from local data: "Bow Jackzon · 24 · Beast Mastery · (blank) · (blank) · (blank) · 0 DPS"; the two other online rows with blank cells (no "no logs" while the record is unreadable).
- `hunter-30-guild-live-row.jpg` — after `/run ForeverSixtyGuildTab.seedLive("Onearm Lopen", {...gearGap=14...})`: "Onearm Lopen · 30 · Arms · … · 14 DPS"; the count line "4 of 97 shown · 2 of 4 online run Forever Sixty" (a fourth member, Belee Dai, came online).
- `hunter-30-guild-header-hover.jpg` — the Spec header hovered: text brightened to the theme white.
- `hunter-30-guild-sorted-spec.jpg` — Spec clicked: sorted ascending with the arrow flipped, the gold header moved to Spec, blanks last.
- `hunter-30-guild-roster-all.jpg` — "Online only" off: "97 of 97 shown · 2 of 4 online run Forever Sixty · 1–20 of 97" (fits the line); offline rows dimmed once, class colours lifted (shaman/druid names readable); online rows bright; the 8-unit scroll thumb.
- `hunter-30-guild-roster-scrolled.jpg` — one wheel notch: "4–23 of 97".
- Open from the controller: the own row's Gap reads "0 DPS" while the tooltips show +3.5 and +1.1 DPS upgrades — sent to the code re-review to settle which quantity `gearGap` is.
- Fixture pass (the L4.4 fixture extended with four guild members, 1b7143da + 1a8f6cab, junctioned in, `/reload`): `hunter-30-guild-fixture-tiers-crop.jpg` (online only: Onearm Lopen · Fury · 480 · 1720 purple · 90%; Bow Jackzon · Beast Mastery · 380 · 1240 blue · 50% · 0 DPS; Sanna Widebraid · Holy · 200 · 640 green · 50%; Belee Dat "no logs" before the name fix), `hunter-30-guild-fixture-all.jpg` (97 rows, "no logs" per row with blanks elsewhere), `hunter-30-guild-fixture-live-all.jpg` (Onearm Lopen seeded live: "Arms · 14 DPS", the live spec word over the record's "Fury"; "2 of 4 online run Forever Sixty"), `hunter-30-guild-fixture-mixed-tiers.jpg` — Score sort over five scored rows: Belee Dat (legendary orange 2310, 100%, online), Onearm Lopen (epic, live 14 DPS), Bow Jackzon (rare, 0 DPS), Aelith Moonclaw (offline, dimmed rare 1180 · 60%), Sanna Widebraid (offline, uncommon 640 · 50%), then the "no logs" rows; offline scored rows readable. Junction restored to the real data addon afterwards.

## hunter-31 — guild tab fix round 2 (d4a5676e), real data addon (too old), after /reload

- `hunter-31-guild-default-crop.jpg` — count line "3 of 97 shown · 1 live of 3 online"; the status line; the self row "Bow Jackzon · 24 · Beast Mastery · … · 0 DPS" (muted zero) — still 0 on the live client (sent to the re-review: what the gap treats as the planned item for a default build with no gear list).
- `hunter-31-guild-gear-tooltip.jpg` — the Gear header hovered: "Gear score from the nightly record".
- `hunter-31-guild-gap-tooltip.jpg` — the Gap header hovered: "Gear gap reported live by the member's addon." / "Biggest slot behind their build, in DPS or points; 0 means on plan."
- `hunter-31-guild-live-tooltip.jpg` — the count line hovered: "Live: guildmates running Forever Sixty share their spec and gear gap."
- `hunter-31-guild-score-desc-game-shot.jpg` / `-spec-asc-game-shot.jpg` — 3840 × 2160 game screenshots with the flyout open: the client's sort arrow on Score (descending) and on Spec (ascending) at native resolution.
- `hunter-31-guild-scrolled-count.jpg` — "Online only" off and one wheel notch: "4–23 of 97 · 1 live of 3 online" (the range replaces the first part while scrolled).

## hunter-32 — guild tab fix round 3 (5ce75588), after /reload

- `hunter-32-guild-self-row-1dps.jpg` — the self row now "Bow Jackzon · 24 · Beast Mastery · … · 1 DPS" (body colour, no longer muted zero); the probe prints `OWN hunter-beast-mastery 1 dps 1`. The tooltips still name Ranger Bow +3.5 DPS and the belt +1.1 DPS, so the ranged slot is not scoring its weapon term in the gap path (sent to the re-review).
- `hunter-32-ranged-tooltip-regressed.jpg` — after fix 3 the worn Short Ash Bow's tooltip reads "Best until level 30 / At level 30: Glass Shooter +9.7 DPS · Gnomeregan · Dark Iron Ambassador / Planned for your build" instead of "Next upgrade: Ranger Bow +3.5 DPS": the active saved build's gear list (the exported equipment, `ranged=3039`) was taken as the plan. The probe prints "BUILD Beast Mastery / Marksmanship build 1 dps". Reverted by fix round 4 (band picks only). (Glass Shooter is a rare-spawn drop the L1.7 data removes; that data is not merged into the addon yet.)

## hunter-33 — guild tab fix round 4 (5294f4bc), after /reload

- `hunter-33-ranged-tooltip-restored.jpg` — the worn Short Ash Bow's tooltip reads "Next upgrade: Ranger Bow +3.5 DPS (+1 Agi, +3.4 weapon DPS) / Planned for your build / World drop · check the AH / Also good: …" again (the band's pick; the saved build's gear list no longer taken as a plan).
- `hunter-33-guild-self-row-4dps.jpg` — the guild flyout's self row: "Bow Jackzon · 24 · Beast Mastery · … · 4 DPS" (body colour), "4 of 97 shown · 1 live of 4 online": the row's gap now equals the tooltip's biggest next upgrade, rounded.

## hunter-34 — group panel round 1 (102f314d), the Social window's Raid tab (solo)

- `hunter-34-raid-tab-no-button.jpg` — the Raid tab open ("Raids are groups of more than 5 people…", "Convert To Raid"): no "Forever Sixty" button appears. Probe: `GB true nil nil RaidFrame true nil` — the button exists and is shown under `RaidFrame`, but `RaidFrame` has no rect on this client (`GetRight`/`GetBottom` nil), so the button has no position; `/fs group` prints "The group panel shows on the Raid tab of the Social window (O)." even with the tab open.

## hunter-35 — group panel fix round 1 (d4d2689c), solo, after /reload

- `hunter-35-group-button-raid-tab.jpg` — the Social window's Raid tab: the gold-edged "Forever Sixty" button under the window's bottom-right edge (anchored to `SocialUIFrame`).
- `hunter-35-group-flyout-solo.jpg` — one click: the flyout "Forever Sixty · Group" at the window's top-right with the single muted line "Not in a group."
- `hunter-35-chip-ready-check.jpg` — `/fs ready` (dry run, solo): the chip reads "✗ Ready check: 1 missing (Bow Jackzon)" with the client's not-ready cross.
- `hunter-35-diag-group.jpg` — `/fs diag`: "Guild tab: closed, 4 rows (1 live), officer no" and "Group panel: open, 0 members (0 inspected, 0 live), ready check 19:43: 1 missing" (solo: 0 members yet 1 missing — the dry run counts the player; wording to settle).
- `hunter-35-button-friends-tab.jpg` — after switching the window to the Friends tab the button stays fully drawn (the accepted limit: a tab switch is not observed until the next refresh; the sentinel frame ruled for fix 2 fixes it).

## hunter-36 — group panel fix 2 on the live client (probes, 2026-10-13 20:00)

No jpgs: the state was probed from chat; the button rested/woke nowhere because the tab test never changes on 1.60.

- `RaidFrame:IsShown()` is true on the Friends tab and on the Raid tab; `RaidFrame:IsVisible()` false on both; its parent is `RaidParentFrame` (338 wide, parent `UIParent`, hidden). The legacy raid frame, not the Social window's Raid tab.
- The Social window's Raid tab content is `SocialUIFrame.RaidFrame` (unnamed, 409 wide, parent `SocialUIFrame`): shown and visible only on the Raid tab. `SocialUIFrame.activeContentFrame` is the active tab's content; `SocialUIFrame.FriendsList` the Friends tab. `SocialUIFrame` has 22 children, mostly unnamed; `GetMouseFoci` over the Raid tab body gives `SocialUIFrame` itself.
- Observed: the "Forever Sixty" button drawn under the window on the Friends tab right after `/reload` + `o`; still drawn after switching to Raid; the flyout "Not in a group." right of the window on the Raid tab.
- `/fs diag`: "Group panel: open, 1 member (you), ready check 20:01: you are missing flask, food, imbue" — as ruled.
- Chip: "✗ Ready check: you are missing flask, foo…" — shortened by the chip width (fix 3: "Ready check: missing flask, food, imbue").

## hunter-37 — group panel fix 3 on the live client (game screenshots, 2026-10-13 20:20)

Build 4d644860 (+ the data merge). Bow Jackzon, level 24 hunter, solo in Astranaar. Game screenshots at 3840×2160 (no OS overlay).

- `hunter-37-1-raid-tab-flyout-hint.jpg`: the Social window on the Raid tab with the "Forever Sixty" button under its bottom-right, the flyout "Forever Sixty · Group / Not in a group." right of the window's top, and the button's hover hint ("Forever Sixty / Your group's specs, gear, scores and buffs. / Click to open or close.").
- `hunter-37-2-friends-tab-no-button.jpg`: after clicking the Friends tab — the flyout closed and the button rested (invisible). Opening the window fresh on Friends after `/reload` shows no button either (probed on screen before the capture).
- `hunter-37-3-chip-ready-check.jpg`: after `/fs ready` with the window closed — the chip reads "✗ Ready check: missing flask, food, imbue" (fits, no shortening); chat: "Ready check (only you see this):" / "Bow Jackzon: flask, food, imbue".
- Diag (hunter-36, same wording at this build): "Group panel: open, 1 member (you), ready check HH:MM: you are missing flask, food, imbue".
- `hunter-37-4-open-on-friends.jpg` / `hunter-37-5-raid-tab-wake.jpg` (20:26): the window opened on the Friends tab (no button), then the Raid tab clicked (the button wakes under the window) — the Friends → Raid route that failed in fix 2.

## hunter-38 — group panel fix 4 on the live client (game screenshots, 2026-10-13 20:49)

Build 98a1fe08. Bow Jackzon, level 24, solo.

- `hunter-38-1-solo-flyout-clear-of-tabs.jpg`: the flyout now sits right of the Social window's side-tab column with a visible gap; the solo text is "Not in a group." plus the muted second line "A party or raid shows each member's spec, gear gap, score and buffs. /fs ready tries the check on you." (wrapped onto two lines).
- (no file — the timed game screenshot did not fire; seen on screen only) after `/fs ready` and `/fs diag` at level 24 — diag "Group panel: closed, not in a group, ready check 20:49: nothing missing"; chat "Ready check (only you see this): nothing missing."; the chip "Ready check: nothing missing" drawn with the red ✗ (wrong for a clean result; fixed in fix 5, see hunter-39-2).

## hunter-39 — group panel fix 5 on the live client (game screenshots, 2026-10-13 21:08)

Build 22a5fc23. Bow Jackzon, level 24, solo.

- `hunter-39-1-solo-flyout-59-units.jpg`: the Social window on the Raid tab, the flyout now 59 units right of the window's edge (clear of the four side tabs by 4 units), the solo text and its two-line hint.
- `hunter-39-2-chip-clean-tick.jpg`: after `/fs ready` at level 24 — the chip reads "✓ Ready check: nothing missing" with the green tick (the ✗ of hunter-38 is gone).

## hunter-40 — loot council helper (L5.3) solo on the live client (2026-10-14, client build 1.60.1.70374)

Build e3cfc47d. Bow Jackzon, level 24, solo.

- `hunter-40-fs-loot-solo.jpg`: chat after `/fs loot 3021` and `/fs loot 3039` — "Ranger Bow: +3.5 DPS over Short Ash Bow, planned" (the band's ranged pick; the same +3.5 as the item tooltip's Next line) and "Short Ash Bow: +0.0 DPS over Short Ash Bow, not planned" (the worn bow against itself), each followed by "Only you see this: nothing was sent (only the master looter asks the group)."
- Seen on screen: `/fs loot 17069` → "…: can't score: the client has not loaded the item yet; try again in a moment." on every try (sent to the review); `/fs help` lists twelve commands; `/fs diag` ends "Loot: none".
- `hunter-40-last-raid-chip-tooltip.jpg` (the L5.3 inbox fixture junctioned in, after `/reload`): the chip reads "Magmadar: 38th percentile, biggest lo…" (shortened by the chip width; diag says "Chip: shown, the line was shortened") and its hover tooltip's second line reads in full "Magmadar: 38th percentile, biggest loss standing in fire"; clicking the chip printed the full line in chat. The fixture link was restored afterwards (the client's ForeverSixtyInbox is now a junction to addon/ForeverSixtyInbox).

## hunter-41 — loot council fix 1 on the live client (2026-10-14, client build 1.60.1.70374)

Build 204ac3e1. Bow Jackzon, level 24, solo.

- `hunter-41-fs-loot-uncached-and-planned.jpg`: chat after `/fs loot 17069` then `/fs loot 3021` — "Item 17069: loading…", "Only you see this: nothing was sent (only the master looter asks the group).", then after five seconds "Item 17069: the client could not load the item (it exists but the server did not answer). Try again after you have seen it drop or linked it."; then "Ranger Bow: +3.5 DPS over Short Ash Bow, planned". Seen on screen before the capture: `/fs loot 3039` → "Short Ash Bow: already worn".

## hunter-42 — the build dialog's Copy tab on the hunter (for the website's addon page)

- `hunter-42-copy-dialog-hunter.jpg` (2026-10-14, build 204ac3e1): `/fs copy` on Bow Jackzon — the "Forever Sixty · Copy my character" dialog with the hunter's own FS1 export string in the box, "Selected. Press Ctrl+C, then paste it on foreversixty.gg" above, and the honest "No companion found. The companion is a small app on your computer that sends this when you log out. Or paste it on the site at foreversixty.gg." below (no "Seen by the site" line on this character).

## hunter-43 — loot council fix 2 on the live client (2026-10-14)

Build 575795fb. Bow Jackzon, level 24, solo; the L5.3 inbox fixture loaded (restored afterwards).

- `hunter-43-loot-lines-and-chip-short-form.jpg`: the chip's last-night short form "Magmadar 38th · standing in fire" with the hover (the mouse is on the chip): "Magmadar: 38th percentile, biggest loss: standing in fire" and a muted "as of 12 Oct" line (no "for Beast Mastery" — sent to the re-review); chat: "Ranger Bow (Ranged): +3.5 DPS over Short Ash Bow, on your list" and "Short Ash Bow: already worn (Ranged)". `/fs diag`: "Chip: shown, short form".

## hunter-44 — loot council fix 3 on the live client (2026-10-14)

Build af24d073; the L5.3 inbox fixture loaded (restored afterwards).

- `hunter-44-last-raid-hover-with-spec.jpg`: the chip "Magmadar 38th · standing in fire" with its hover reading "Magmadar: 38th percentile for Beast Mastery, biggest loss: standing in fire" (the saved build's spec now derived) and the muted "as of 12 Oct" line.

## hunter-45 — polish round A on the live client (2026-10-14)

Build 2f9963be. Bow Jackzon, level 24.

- `hunter-45-ranged-tooltip-thrown-tag.jpg`: the worn Short Ash Bow's tooltip — "Next upgrade: Ranger Bow +3.5 DPS (+1 Agi, +3.4 weapon DPS) / Planned for your build / World drop - check the AH / Also good: Lil Timmy's Peashooter +2.2 (needs bullets), Cracked Blacksmith Hammer +0.8 (thrown)" — the "(thrown)" tag and the tag rows aligned under the names. No Lua error opening the pane.
- `hunter-46-panel-flyout-round-a.jpg`: the character-panel flyout after round A — "Ranger Bow · over Short Ash Bow +3.5 DPS / Ranged · World drop - check the AH", "Blackened Defias Belt · over Dusty Belt +1.1 DPS / Waist · The Deadmines · Captain Greenskin · 30% drop", "Nothing better for your level in other slots."; Rotation rows with "Rank 1"/"Rank 3" muted after the names (Multi-Shot shows none — the client gives no subtext for it); the dim "Rapid Fire at level 26" row's name legible; Build "Plan: Beast Mastery 31 / Marksmanship 20 · 12 of 15 points on plan · 3 off plan".
- `hunter-47-panel-flyout-build-region.jpg`: the flyout's Build region ("Plan: Beast Mastery 31 / Marksmanship 20", "12 of 15 points on plan", "3 off plan" in red) — the off-plan hover did not open at the cursor position tried (README row 83 stays owed).
- `hunter-47-settings-round-a.jpg`: the settings page after round A — the title warning "Player ratings (on player tooltips): too old (patch 1.60.1). Get the new Forever Sixty Data at foreversixty.gg." (the "patch" is the build truncated to three fields, so it reads 1.60.1 — the same as the client's; a wording doubt for the personas) and the Gear-tips hint "Adds a line to item tooltips: in your build, and better than what you wear. Off: hides player scores too." with no orphaned word.
- `hunter-48-off-plan-hover.jpg` (round A, README row 83): the "3 off plan" hover, anchored beside the flyout (top-right of it): "3 off plan / Pathfinding · Beast Mastery · 2 off plan (Rank 2/2) / Improved Revive Pet · Beast Mastery · 1 off plan (Rank 1/2) / Not in your plan. A respec moves them. / Next point at level 25: Lethal Attacks". (Round B rewords the action line to "Respec to fix this.")

## hunter-49 — polish round B on the live client (2026-10-14)

Build 837f8d9b. Bow Jackzon, level 24.

- `hunter-49-talent-marks-overlays.jpg`: the talent window after `/reload` with the marks now drawn on addon-owned overlays (no hooks on the client's window): order numbers and level ranges on the planned nodes, the gold rings on taken ones, the green ring on the next, the strip "Next: Lethal Attacks (1 of 5) · 12 of 51 points · 3 off plan". No Lua error.
- `hunter-49-settings-nine-switches.jpg`: the settings page with the ninth switch "Ask the group about rare drops too" (off; hint "As the master looter you ask the group about epics. On: rares too, whatever the loot threshold."), the reworded Gear-tips hint, the title warning in its undated form ("too old." — this data record carries no generated date), and the Commands list gaining "/fs build: switch between saved builds". `/fs help` printed thirteen commands plus the two tester lines.
- `hunter-49-guild-flyout-round-b.jpg`: the guild flyout (Online only, "2 of 98 shown · 1 live of 2 online", the "Player ratings: too old" line): the self row "Bow Jackzon · 24 · Beast Mastery · … · 3.5 DPS" (one decimal under 10) and "Cryptic Stem · 24" with the no-record cells (a dim "—" was ruled — the personas judge whether it shows); no record-date line (this record carries no date).

## hunter-50 — round B with the L4.4 data fixture junctioned in (2026-10-14; restored afterwards)

- `hunter-50-guild-flyout-dash-and-date.jpg`: the guild flyout with a readable record — the count line "2 of 98 shown · 1 live of 2 online · as of 13 Oct" (the record date), the self row "380 · 1240 · 50% · 3.5 DPS", and "Cryptic Stem · 24 · — · no logs · —" (the dim dashes for the no-record cells, "no logs" in the Score cell).
- `hunter-50-talent-node-hover-plan-line.jpg`: the talent window after the overlay refactor, hovering the Beast Mastery "1-5" node (Deadly Aspects): the client's tooltip with the addon's line "Forever Sixty: Your build: 5 ranks, points 1-5" once (not doubled) after a `/reload`.

## hunter-51 — round B fix 1 on the live client (2026-10-14)

Build b1f27429. Bow Jackzon, level 24.

- `hunter-51-talent-strip-level-points.jpg`: the talent window's strip now reads "Next: Lethal Attacks (1 of 5) · 12 of 15 points · 3 off plan" (the level's points, matching the flyout), the marks unchanged.

## hunter-52 — the integrated release tree on the live client (2026-10-14, build 377e68c1, client 1.60.1.70374)

- `hunter-52-ranged-tooltip-no-thrown.jpg`: the worn Short Ash Bow's tooltip with the 70374 data and the hunter-thrown fix — "Next upgrade: Ranger Bow +3.5 DPS … Also good: Lil Timmy's Peashooter +2.2 (needs bullets), Lovingly Crafted Boomstick +0.5 (needs bullets)" — the Cracked Blacksmith Hammer is gone from the hunter's ranged list.
- Seen on screen: `/fs diag verbose` now lists the addon's frames ("Chip: login, 0 s in", "Slot mark x2: a bag or paper doll slot first drawn, 0 s in", "Minimap button: login", "Settings page: login", …) and ends "memory: 5694 KB" right after the reload; after `collectgarbage` the figure is 4,850 KB (≈2 MB shipped data + Locale 158 KB + ~2.7 MB code and state).

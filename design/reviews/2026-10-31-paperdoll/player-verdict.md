# Paperdoll: wow-player verdict (2026-10-31)

Seen: raid-1280, beta-1280, raid-390, states. Read but not looked at: raid-1024, raid-2000, raid-360, empty, loading, beta-390.

## Verdict: HOLD

## Blockers

1. Hands slot: worn "Raider Gauntlets" with the gain line "+4.8 pts · Raider Gauntlets" (beta-1280). The pick is the item already on the character, yet it claims a gain. A player screenshots that and mocks it. Either it is a different copy or enchant (then say so) or the gain is wrong. Never offer an item you already wear as an upgrade.
2. Item level rule. A gold "53.3" is the mean of the 12 known items of 15, with the 3 beta items left out. Those unknown items are the newest ones and probably the highest level, so the headline number is wrong in the one case the board exists for, and two characters cannot be compared. Classic's character frame also has no average item level, so players will read it as a score. Show the figure only when every worn item is known (otherwise show "12 of 15 items" and a dash), say which slots count (shirt, tabard and a two-hander's empty off hand), and name the number in the spec.
3. The tooltip on states.png is broken. The icon sits on top of the name ("nheart Helm") and the Item Level line ("n Level 61"). The last line reads a bare "Blacksmithing", where the client says "Requires Blacksmithing (300)". It also shows Item Level 61 and Requires Level 56 on a helm players know as a level-60 piece, so check it. The board is the spec's proof of the tooltip, and it shows a defect.
4. Two-hander: beta-1280 shows Arms "Blight" in main hand with "Two-hand in main hand" in the off hand, and the greyed icon is a dagger. If Blight is a one-hand dagger, the state is wrong for that item. Confirm the item is two-handed in the build data, or change the example.

## Shoulds

- Slot order matches the client (head to wrist left, hands to trinket right). The client has Shirt and Tabard sockets, so Wrist sits level with Trinket 2 there. Here Wrist floats at row 6 beside Finger 2, with a hole under the left column. Show empty Shirt and Tabard sockets, or accept that a player will say "where is my shirt".
- Weapons: in the client, main hand, off hand and ranged sit together at the bottom centre. Here main hand is under the far-left column and off hand under the far-right one, about 1000 px apart, with ranged on its own row below. It reads as three unrelated items. Group them in the centre, as the client does.
- Unknown items show "Item 288887" with no slot word. You cannot tell if it is waist or neck. Put the slot label as a small caption on unknown and empty slots, and on filled slots when two rings or two trinkets are shown. "Finger 1" and "Trinket 2" are right, but they only appear on empty slots, so the words are barely used.
- No item level on slots in the sim or character cards, only in the report detail. Players expect to read it in the tooltip at least; show it in the tooltip on every card.
- The gain line says "+18.4 pts · Lionheart Helm" with no verb, and "pts" is undefined here. Use "Upgrade: Lionheart Helm +18.4 points" or follow the weights convention, with the unit explained once on the page.
- Phone (raid-390): "Signet Ring of the Bronze Dragonflight" clamps to "Signet Ring of the Bronze". Names cut at about 15 characters give a wrong name. Allow 3 lines or reduce font size on long names. Main hand is also cut, and ranged sits left of centre, not centred as specced.
- Quality colours read right. The 2 px game-hue border and the name text colour agree. Purple and blue on dark are legible. Gold item level is fine.
- Data to check for accuracy: the Frozen Heart of the Mountain trinket on Arms and Fury (its Use is not modelled yet), "Ravencrest's Legacy" and "Shadowsong's Sorrow" as a Fury raid pair, and The Purifier as the ranged pick. Check that the Knight-Lieutenant and Sergeant Major PvP items are the rank 7 and rank 5 sets (both under rank 10).

## Re-check (fix round 1)

Seen: states-1280, beta-1280, raid-1280, tip-right-1280, tip-phone-390. Read only: the other tip boards, raid-390/1024/2000.

### Verdict: SHIP (no blockers left)

### My four blockers
1. Gains: closed. Hands on beta-1280 now reads "Devilsaur Gauntlets" with "Timbermaw Brawlers +4.4", a different item. Worn picks read "Star of Mystaria / Best in slot" and "Frozen Heart of the Mountain / Best in slot", muted, with no gain. The header "+148.8 points across 9 slots" matches nine gain lines. The board shows nine, so it checks out.
2. Item level: closed. beta-1280 has no figure, and the footer note says "Item level shows once every item is known." raid-1280 shows 62.0 with all slots known. The counted slots are named in §3.5 (every slot except shirt and tabard). One small point: the two-hander's empty off hand must not count as unknown. §3.5 should say so in one clause.
3. Tooltip board: closed. The icon no longer sits on the name, the panel hangs from the right edge on the right column, and on phone it sits below the tapped cell. The numbers (Timbermaw Brawlers ilvl 61, Requires Level 56, 107 Armor, Plaguehound Leggings ilvl 62, Requires Level 57) come from the record.
   - "Requires Blacksmithing (300)": I agree with the decline. The client shows a profession requirement on the recipe, not on the crafted item. My request was wrong. "Crafted: Leatherworking" is the site's own line, and a reasonable one.
4. Two-hander: closed. Arcanite Champion has the empty off-hand socket and "Two-hand in main hand". That matches the client, where the off hand is simply empty.

### Declined suggestions
- Weapons centred, as the client has them: I accept it for now. The site and the Discord sheet should show one picture, and a player will not mistake the layout for something else. It is a parity decision, so the owner should rule on it together with the shirt and tabard sockets. Do not leave it as an unstated default. If the owner wants the client layout, the sheet changes too.
- Empty shirt and tabard sockets: I accept it, with the same owner ruling. The hole under the left column at 1280 is visible, but it is not a defect.

### Remaining shoulds, none blocking
- The tooltip lacks "Binds when picked up" and the armour type word beside the slot ("Legs" / "Leather"; "Hands" / "Leather"), which the client shows on every armour tooltip. Add both.
- The unknown-item slots now read "Waist · Item 288887", which fixes my slot-word complaint.
- Gain lines are now "Lionheart Helm +60.0" with no unit word. The header says "points", but each row does not. Acceptable.
- The raid-1280 board's "Hyper Deluxe Sniper Rifle Mk XVII" wraps to two lines. That is fine, and the full name shows.

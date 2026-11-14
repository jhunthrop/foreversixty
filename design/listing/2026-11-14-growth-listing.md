# Store listing changes from the growth research (2026-11-14)

Prepared, not published. The evidence is design/marketing/2026-11-14-addon-growth.md. This file
changes three fields of `curseforge-description.md` (name, one line, categories) and leaves its
description, Install section, gallery order and captions as approved. The name waits on the
owner's yes.

## Project display name (owner's call)

CurseForge's search matches the project name and the author, not the summary
(https://docs.curseforge.com/rest-api/, "free text search in the mod name and author"). Today
"Forever Sixty" matches no word a player types. The TOC Title and the folder stay "Forever Sixty".

| | Name | Length | Searches it answers |
|---|---|---|---|
| Proposed | **Forever Sixty: Gear Upgrades, BiS and Talents** | 45 | gear, upgrades, bis, talents |
| Alternative 1 | Forever Sixty: BiS, Talents and Rotation Helper | 47 | bis, talents, rotation |
| Alternative 2 | Forever Sixty: BiS Gear and Talent Planner | 42 | bis, gear, talent, planner |

Use the same name on Wago.

## One line (CurseForge summary, Wago tagline)

The live summary (read through the CurseForge API on 2026-11-14) is "Your next upgrade, your next
talent and your Forever Score, on the game's own frames. No setup, no network, nothing in combat."
The approved line in `curseforge-description.md` was never published. Proposed, 192 characters
(CurseForge's cap is 200):

> A green arrow on any bag item, quest reward or roll that beats what you wear, with the gain in DPS. Your next talent marked, BiS for each slot, what to press, and raid sign-ups for your guild.

The arrow leads and the score is not in the line: a score up front invites the gear-score
argument (the report's ranked change 9). The description still covers the score.

## Categories

| | Categories |
|---|---|
| Live (CurseForge API, 2026-11-14) | Class, Tooltip, Guild, Map & Minimap |
| Approved, never published | Boss Encounters, Class, Guild, Tooltip, Minimap, Quests & Leveling, Combat |
| Proposed | Class (primary), Quests & Leveling, Tooltip, Bags & Inventory, Guild, Boss Encounters, Combat |

The approved list, with Bags & Inventory in place of Minimap: the arrow lives on the bags, and
the minimap button (off on a new install) is not a reason to install. Quests & Leveling is where
DungeonJournal, GearQuest Forever and Talents Forever all sit.

## In game (shipped in the addon, not the store)

The TOC Notes the client's AddOns list shows now say what the addon does, and the TOC carries
`## Category: Class` (the field Questie's Forever TOC uses; not yet seen in our client):

> Gear upgrades, BiS, talents, rotation and raid sign-ups. A green arrow on gear that beats yours and your next talent marked. /fs opens it.

## Dashboard clicks

CurseForge (authors.curseforge.com, project 1706689): General > Project name, Summary; Categories
> pick the primary, then the others. Wago (addons.wago.io, XKqEdoKy): Settings > Name, Summary,
Categories (react-select). Quirks are in the store-listings memory note. Then copy the chosen name
and line into the table of `curseforge-description.md` so it stays the one source.

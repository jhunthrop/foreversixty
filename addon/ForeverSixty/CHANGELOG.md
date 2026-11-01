# Forever Sixty

## 2.0.21

Data for client build 1.60.1.70707: the item, talent and spell tables the addon ships were regenerated for the new client by the client watcher. Nothing else changed.

## 2.0.20

Data for client build 1.60.1.70658: the item, talent and spell tables the addon ships were regenerated for the new client by the client watcher. Nothing else changed.

## 2.0.19

Data for client build 1.60.1.70630: the item, talent and spell tables the addon ships were regenerated for the new client by the client watcher. Nothing else changed.

## 2.0.18

- The line, the Gear page's bag rows and the minimap count never offer a piece you cannot wear: its level, its class line and your armor and weapon training are checked first (a level-21 warrior was offered the Shaman tier-1 gauntlets). /fs loot says why when you ask: "Not for you: requires level 60", "Not for you: Shaman only".
- The line's hover names the slot and what the piece replaces ("Hands: Earthfury Gauntlets replaces Gloves of the Fang, +3.5 DPS"), and the guild window's gap and that number come from one baseline: the hover reads "Gap to your planned gear: 0.7 DPS".
- The guild window with nobody scored says so once ("Nobody here has logs on the site yet. Scores appear after a member's logs are uploaded."), draws one dash in every cell it has nothing for (the under-level-10 rows too), sorts by Level until a score exists, and counts "14 online of 773 · 1 running Forever Sixty" with no date. Its title is the guild's name ("OLYMPUS XXVII · Forever Sixty") and its button sits inside the guild window beside Invite Member.
- The Raid tab's button reads "Forever Sixty · solo" alone and opens the panel with your own row ("Invite someone and their row appears here."); /fs group and /fs panel print where the panel is and when it fills when you are solo.

## 2.0.17

- The build preview lists every gear slot it counts, at the build's own level ("Gear in this build · level 24 · 14 of 17 slots"), each with where it comes from in full; the build window and the line say which build is which ("Beast Mastery build · from the site, 22 Oct · 15 points", "Replaces Beast Mastery build (51 points)").
- The line, the chat line after Use and the talent strip name the next point ("Build complete · next point at level 25", "next at level 25: Focused Fire"); the line keeps one class icon and sits under the buff row when the buffs are in its way.
- "Seen by the site 02:21 pm" is "Site has your save from 1 h 35 min ago" (and "Not sent to the site yet · log out to send"), the same age in /fs diag and the login line.
- The ready check says the same words on the line and in chat, with no "(test)" and what it checked; /fs help lists the player's commands and /fs help all adds the testers'; the welcome names the addon once and says what the companion does; the Options page lists four commands and /fs help for the rest, its explanations at full contrast; the player tooltip's /fs link hint shows three times, then goes; Copy my character says what the export holds.
- Data for client build 1.60.1.70532: the item, talent and spell tables the addon ships were regenerated for the new client by the client watcher.

## 2.0.16

- The Forever Sixty line beside the minimap is on screen at a first login and says your next step ("Next: Endurance Training (1 of 5)", or "Pick a build: /fs"); a saved spot that is off screen puts it back beside the minimap, it sits above the buff icons, and the welcome names the line only when it is drawn.

## 2.0.15

- The first-login welcome says what you can see and how your character reaches the site; 'chip' is gone from the words you read.

## 2.0.14

- A first login says what the addon does and where its commands are; Escape closes the build dialog whatever has focus, and typing to chat is never swallowed by it.

## 2.0.13

- The build preview never cuts a talent or tree name short, and the dialog's hint only says Enter to use when that works.

## 2.0.12

- Builds from the site are rows you can pick in /fs, with the character they are for and what they change; the preview never cuts a sentence; /fs diag says when the inbox was read and that builds arrive at login, not on /reload; the addon's times match the chat's clock.

## 2.0.11

- Builds sent from the site, the site's acknowledgements and guild notes now reach characters on realms whose name is not the ruleset word (the beta's Classic Beta PvP and PvE realms): /fs diag's Site line, the login line and the inbox read them.

## 2.0.10

- /fs diag now says whether the site has seen this character and what is waiting in the inbox; at login one line tells you when a build is waiting, when the site has not seen you yet, or when the companion has gone quiet.

## 2.0.9

- Your gear, talents and guild reach the site again. The export is now kept up to date while you play and written at logout from that copy, because the 1.60 client answers empty gear, talents and guild at the moment of logout.

## 2.0.8

- Old one-word copies of your characters (saved before the addon learned Forever surnames) are dropped on the next save, so the site stops listing them twice.

## 2.0.7

- Your characters reach the site again through the companion. The 1.60 beta client does not name its region, so every saved character carried an empty one and the site refused the whole list ("region must be one of us, eu, kr, tw, cn"; the companion's log said "The site refused the character list"). The addon now records "us" when the client cannot say (the beta realms are one world, hosted in the US) and moves the records it saved before under that region, so the next logout syncs everything.
## 2.0.6

Data for client build 1.60.1.70450: the item, talent and spell tables the addon ships were regenerated for the new client by the client watcher. Nothing else changed.

## 2.0.5

- The default builds for client 1.60.1.70425 are rebuilt from the guides, carried onto the new talent trees by talent name. 2.0.4 derived them from codes carried by node id, which put ranks on the wrong talents where the client had moved them (Tidal and Totemic Focus, Elemental, the mage and priest reshuffles), and every mage, priest, warlock and shaman default refused to load.

## 2.0.4

Data for client build 1.60.1.70425: the item, talent and spell tables the addon ships were regenerated for the new client by the client watcher. Nothing else changed.

## 2.0.3

- Clicking the chip opens the talent window on its Talents tab again. The 1.60 client opens the player-spells window on whichever tab it showed last, so after a visit to the spellbook the chip landed there; the addon now asks for the Talents tab by name.

## 2.0.2

Two fixes for characters just starting out, seen on a level-18 shaman and a level-6 paladin.

- Below level 20, before the first set of gear picks at level 20, the gear surfaces no longer go dark. An item's hover says "Picks start at level 20" and shows what the level-20 pick is for that slot and where it comes from ("At level 20: ..."); a bag item the plan names says "Planned for your build at level 20"; the character panel lists those picks under "Next targets" with "Picks start at level 20." beneath; the chip's hover says "Gear picks start at level 20."; `/fs diag` reads "Panel: closed, picks start at level 20, rotation N rows". Nothing is called an upgrade until the character reaches the band, so no gold corner and no gear gap yet.
- When you follow a build (the path you picked on the chip, a default or a saved build), every surface now reads that build's spec: a level-6 retribution paladin with no talent points sees "Forever Sixty · Retribution" on item tooltips, not "Holy". With no build followed the talents decide, as before.

## 2.0.1

Two fixes reported on a fresh level-1 rogue right after 2.0.0.

- The character panel's rotation rows show the rank you have learned (the spellbook's highest), not the rank of the spell the level-10 band names; an unlearned "at level N" row keeps the band's rank.
- Without a guild the "Forever Sixty" guild button stays hidden (and appears on the next refresh after you join); a flyout opened anyway reads "Not in a guild." instead of promising a roster, and `/fs diag` says "Guild tab: not in a guild".

## 2.0.0

Forever Sixty 2.0.0 replaces the six-page window with surfaces on the client's own frames,
so the answer to "what makes my character better right now?" is where you would act on it.
Nothing needs setting up: a character with no saved build follows the default build for its
spec, and every surface with nothing to say stays out of the way.

### What changed for you

- **`/fs` opens one small build dialog, not a window.** Paste a build in or copy your
  character out for the site; the old six pages (Overview, Talents, Gear, Guild, Export,
  Settings) are gone, and their content lives on the surfaces below and in the client's
  Settings page. The dialog shows "Following: <build> · 51 points", previews what a pasted
  code would replace, and never overwrites a raid build you restored by pasting a leveling
  one. `/fs build <slot>` switches between your saved builds (Raid, Leveling).
- **Zero setup.** A character with no saved build follows the default leveling build for its
  spec (the raid build from 60), taken from the simulator's guide builds. At level 1 to 9
  with no spec yet, the chip offers a path to pick. The followed build, the chip's hide and
  the chosen path are kept per character: an alt never sees the main's.
- **The chip.** One small icon and line beside the minimap says the single next thing: a
  talent point to spend, an upgrade waiting in your bags (click to equip), loot in the
  dungeon you just entered, what a ready check is missing, last raid night's biggest loss,
  or nothing at all. It fades in combat, hides with the UI, and right-click hides it until
  your next level.
- **Marks on the talent window.** Your build's talents are numbered in order, the next one
  is ringed, and a strip counts the points your level allows ("12 of 15 points · 3 off plan").
  Hover a planned node for its place in the plan; hover "off plan" for the talents you
  spent outside it and the next point to take. The marks sit on frames of the addon's own,
  never on the client's talent buttons.
- **Item tooltips say whether it is better.** "Upgrade +4.2 DPS over <worn item>" (healers
  read HPS), the next upgrade for the slot and the
  level it lasts until, where it drops with the boss and its chance, "Also good:" runners-up,
  a thrown weapon tagged "(thrown)" and ammunition users "(needs bullets)", and a pair
  verdict for two one-handers against a worn two-hander. An item the client has not cached
  yet shows "stats pending" and redraws the moment it arrives.
- **Paper doll and bag marks.** A green arrow on a character slot and on the bag item that
  beats it (or fills an empty slot), a gold corner on a slot whose next upgrade is worth
  chasing, and a hover on the slot that names the pick for your spec, faction and level. The
  chip's hover carries the legend. Another bag addon's bags are left alone, and `/fs diag`
  says so.
- **Quest rewards and group rolls.** "Best for <Spec>" in gold on the reward that is an
  upgrade (or "Vendor" on each when none is), and a one-line verdict beside the roll frame
  that matches the item's tooltip.
- **A "Forever Sixty" tab on the character panel.** A button opens a flyout with your next
  targets (up to five, with the source and drop chance of each and the worn item each
  replaces), your rotation at your level (the rank to use, and "Rank N at level M" for what
  you have not learned yet), and your build (the split, "12 of 15 points on plan", the
  off-plan count, a way to use a different build, "Copy my character"). At 60 it adds your
  simulation (DPS, your top stat weights in DPS per point, "Sim me again") and your last
  raid (each boss's parse percentile, biggest loss, deaths, consumables missed and rotation
  adherence).
- **Player tooltips show the Forever Score.** "Forever Score 1240 · Rare · as of 13 Oct" in
  the tier colour, a Shift breakdown of the five parts (Gear, Experience, Parses,
  Attendance, Preparation) and their total, or "Forever Score: no logs yet" and never a
  zero. Guildmates running the addon also show their level, spec and gear gap live. `/fs
  link` copies a player's site page.
- **Score badges in chat.** "FS 1240 ·" before a player's message in LFG, Trade, guild chat
  and whispers, in the tier colour (muted when the record is more than three days old).
  Switch it off in Settings.
- **A "Forever Sixty" tab on the guild frame.** Your guild's roster with level, spec, gear
  score, Forever Score, attendance and gear gap, sortable, with an "Online only" filter and
  the "as of" date of the records; members without the addon show the nightly record, members
  with it show their live snapshot. For officers: the approvals waiting on the site, and "Push the raid build to guild" behind a
  confirmation. A pushed build is announced once in chat and listed in `/fs` as "pushed by
  <officer>"; nobody is made to use it.
- **A "Forever Sixty" tab on the Social window's Raid tab.** Each member's spec, item level,
  gear gap, score and buffs for a party or raid; members without the addon are read from
  inspect, one inspect every two seconds, never in combat. `/fs group` opens it, `/fs ready`
  runs the ready check on demand (missing flask, food, weapon imbue and, inside Molten Core,
  Onyxia's Lair or Blackwing Lair, class coverage for the bosses), and `/fs floor N` tags
  members scored under N. It informs; it never kicks, whispers or announces for you.
- **Loot you can ask the group about.** `/fs loot <item link>` says what an item is worth to
  you ("+6.3 DPS over <worn item>, on your list"). As master looter, an epic drop asks the
  group which of them gain from it and by how much, with a 10 second countdown in the title
  and nothing asked about an item no member's class can wear; a ninth setting, "Ask the group
  about rare drops too", adds rares.
- **Settings are the client's own page.** Nine switches (the chip, hide in combat, marks on
  the talent window, minimap button, gear tips on item tooltips, score badges in chat, ask
  the group about rare drops, share my build with my guild, save my character for the
  companion app), the thirteen commands, and plain lines for the data: "Player ratings (on
  player tooltips): too old (made 9 Oct)", "broken or missing", "not installed". Reset
  positions puts the chip, the minimap button and the build dialog back.
- **Thirteen slash commands**, all listed by `/fs help`: `/fs`, `/fs copy`, `/fs build`,
  `/fs chip`, `/fs inspect`, `/fs link`, `/fs group`, `/fs ready`, `/fs loot`, `/fs floor`,
  `/fs options`, `/fs diag` and `/fs help`. `/fs diag` states what is not working, and
  `/fs diag verbose` is for testers.

### Companion, data and the site

- **The companion app** sends what the addon leaves for the site at logout (your character,
  a "Sim me" request, your guild tab's roster, once each) and writes back, through the small
  `ForeverSixtyInbox` addon that ships in the same download, the builds you chose on the
  site, your newest sim with its upgrades, your last raid and what the site last saw. Both
  folders go into `Interface/AddOns`.
- **The nightly `ForeverSixtyData` addon** carries each rated character's Forever Score, its
  five parts, tier and date, and each guild's standing, so tooltips, chat badges and the guild
  tab work with no companion.
- **Best-in-slot data is for client build 1.60.1.70374.** Leveling picks come from the
  simulator once per spec and band, deterministically (the same input always gives the same
  picks), with a rare spawn's drop kept out of a slot something else can fill and items the
  client has not confirmed never ranked. Default builds, DPS per point, the rotation lists
  and each raid's needs ship in the same data.
- The site's /setup page now has a "What it does in the game" section for these surfaces and
  lists the same thirteen commands, and the Forever Score pages call the fifth part
  "Preparation", the word the tooltip uses.

### Under the hood

- Every number is computed once and read by every surface that shows it, so the chip, the
  tooltip, the marks, the panel and the guild tab cannot disagree. Surfaces measure what they
  draw and turn themselves off, with the reason in `/fs diag`, if they cannot fit or find
  their anchor.
- The addon writes nothing to the client's frames: marks live on frames the addon owns,
  parented to the client's, with no scripts hooked onto the client's buttons.
- Wire formats are pinned by shared golden fixtures (`docs/contracts`) checked by the addon,
  the companion and the API.

### Measured on the 1.60.1.70374 beta client

- Memory after `collectgarbage`, idle in town on a level-24 hunter: about 4.85 MB, of which the shipped `Data.lua` is about 2 MB and Locale 0.16 MB. The tooltip cache is capped at 200 sections and the inspect cache at 100 records; `/fs diag verbose` lists every frame the addon created and ends with a `memory:` line.


## Before 2.0.0

These notes were kept while the six-page window existed; 2.0.0 replaces that window with the
surfaces above, so some of what they describe (the Overview, Talents, Gear and Export pages) is gone.

- Added `/fs record on|off|status|clear`, a measurement recorder for the simulator: energy
  ticks, white-swing outcomes, Eviscerate and other yellow hits with their combo points,
  Windfury procs on you and your pet, mana regeneration and Shadowfiend returns, kept as
  a capped ring buffer in SavedVariables. Off by default, silent, and free while off.
- Fixed the Overview's build-arrived banner: the Load button could land past the window's
  own right edge because it duplicated Dismiss's own width-and-gap arithmetic against a
  frame width that could disagree with it; Load now anchors directly off Dismiss instead,
  so the two can never drift apart, with a layout test pinning every banner child inside
  the page's own width.
- The Talents page's Top Gear upgrade rows now show the item's own icon and its quality-
  coloured link (once the client has it cached) beside the slot's own icon, instead of a
  plain "<slot>: <item name>" line; hovering a row shows the item's real tooltip. A message
  from an older companion build with no item id still falls back to the plain name.
- Redesigned the Overview's rotation card: each priority row now carries the ability's own
  icon, its name, its rank, and the condition on its own muted line beneath, numbered
  top to bottom; a compact header strip names your spec (with its icon), your level band
  and which data build it is from; the row for whatever you just learned this level glows,
  off the same event as the toast; novice mode's "+N more" expands the card in place
  instead of sending you to Settings, and advanced detail adds the spell id and cooldown.
- The best-in-slot hover now shows the recommended item properly: its own icon inline, its
  name in its real quality colour, and its item level, plus a second tooltip beside it
  showing the item's own full details, the way shift-to-compare already does. An item the
  client has not cached yet is asked for once and the hover redraws itself in place the
  moment the data arrives, rather than showing a bare "item:12345" until you move the mouse
  away and back.
- Hovering an equipment slot on the character frame now names the best-in-slot item for
  your spec, faction and level, marked equipped when it is what you already wear or newly
  the pick since your last level; an empty slot shows it directly, since it has no item
  tooltip of its own to add to. "Show advanced detail" adds where it comes from (quest,
  dungeon, crafted...).
- The Talents page now shows the site's Top Gear upgrade queue for this character (slot,
  item, source and delta), newest first, three at a time until "Show advanced detail" is
  on; item tooltips call out a capped stat when the site's saved weights say more of it is
  wasted. Both read the companion's inbox, addressed per character like a queued build.
- Turning on "Show advanced detail" in Settings (off by default) widens the personal
  rating card with its top components, alongside the existing gear and rotation detail.
- A personal rating card on the Overview reads your rating from the Forever Sixty Data
  addon, with an empty state when you are not rated yet.
- A rotation card on the Overview shows your spec's priority list at your current level
  (the top four abilities, or all of them with advanced detail on), and a toast names a
  new ability the moment it enters your rotation on level-up.
- Named build slots (Raid, Leveling, PvP) on the Talents page: switch with one click, and
  the tracker and talent glow follow whichever is active. Your existing followed build
  moved into the Raid slot.
- A one-time banner on the Overview announces a companion build queued for this
  character, with a one-line summary of how it differs from what is loaded and a
  Load it / Dismiss choice.
- The followed build and the companion's inbox are now per character: logging in on a
  second character no longer shows the first character's talent tracker, glow and Follow
  tab, and an inbox build queued on the site for one character is no longer offered to
  another. The first character to log in after this update keeps the account's old build.
- The Overview's Send-to-the-site card shows the start of your code and its length under
  the title, and the Export page's code box shows the start of the code rather than its
  tail, so there is visibly something to copy before you press Copy code.
- The export carries your character's level, so the planner and simulator know it
  instead of always assuming 60.
- Characters are exported under their full name, first and last, on the Forever client
  (its UnitName answers the first name alone from build 1.60.1.70009).
- First release: import your character into the planner, follow a build in game,
  and score what is in your bags against it.
- A window (`/fs`, or the minimap button) with Export, Follow, Gear and Settings
  tabs, a draggable tracker for the next talent point, a glow on that talent in
  the talent window, and a page in the game's own Settings panel.

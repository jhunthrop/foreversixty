# Forever Sixty

## Unreleased

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

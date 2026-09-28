# Forever Sixty

## Unreleased

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

# Task L8.3 fix round 1 captures (Chrome, fixture build, preview on 4371)

Setup page (images scrolled into view, every `img.complete` true with no forcing; box sizes in `_setup-notes.txt`):
- setup-1280-full.png, setup-360-full.png: whole page after scrolling.
- setup-1280-first-paint.png, setup-360-first-paint.png: first paint at DOMContentLoaded, lazy loading honoured.

Roster tab with the snapshot on (Zulmara/Kraggor matched, Duskstrider consent-withheld "seen Mon 23:12", "per addon" on Zulmara; measurements in `_roster-notes.txt`):
- roster-1280-snapshot.png: fixed In game slot, header over it; rows 67 px high whatever the addon knew.
- roster-1024-snapshot.png, roster-768-snapshot.png: below xl the In game text is its own line; no "—" on rows without data.
- roster-360-snapshot.png, roster-360-header-wrap.png: phone, header wrap.
- roster-1280-sort-focus-visible.png: sort header with keyboard focus (2 px gold outline, :focus-visible true).
- The In game hovers are native title tooltips, which headless captures do not draw; their text is listed in `_roster-notes.txt` ("hovers:").

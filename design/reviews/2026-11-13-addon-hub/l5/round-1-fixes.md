# L5 Guild tab: round 1 fixes and rulings

Fixed in the build (renders redrawn):
- Guild line on one line: standing and raid nights; each raider's attendance stays the table's column.
- `you are 9th of 15 raiders` once officers name raider ranks; `you are 9th of 15 at level 60` while they have not (the fallback says so).
- 4 px above your own standings row; the ringed class crest replaces the square spec icon on standings rows.
- `1 of 2 jobs filled` dropped. Roles coloured against the comp target: short in warning, met in success (the web's Short mark).
- The time reads `Thu 10 Dec, 20:00 local`: the shared clock is the client's local clock, never the server's, so `server` would be false.
- `Change it on the site` under an answer the companion cannot send; `Sign up on the site` only with no answer.
- The job line names the boss when it is not the raid's own title (`Your job: Main tank · Magmadar`).
- Your prep: item icons with the client tooltip; an `In your bags` row with the supplies carried by kind (`Flask 2 · Food 20`).
- Next raids: a row's hover gives that raid's sign-ups (`12 of 40 coming`) when its raid_roster came.
- Officers sent: `from Brannoc, Tue` (the push is now stored with its time).

Ruled, not changed (owner-approved spec copy and colours, design/specs/2026-11-13-addon-hub.md 3.4, 3.5, 3.7, 8, 11):
- `Your sign-up:`, `Your job:`, `Your prep`, `for Onyxia`, `12 of 40 coming`, `No raid from the site yet.` are the spec's verbatim copy.
- The job label and `Not signed up` in client gold; Late in client gold (spec 3.5 and the web's raid page).
- The footer `Sends when the companion runs` is spec 3.7's busy state.
- Section headings in the title face, as the shipped Builds tab after the 2026-11-13 brand review; brand-consistency passed that.
- Raid leader, invite time, guild progression per raid and a resistance target have no data (no field in raid_signup, raid_roster or the ratings pack); they are left out, never faked. Spec 9 follow-ups.
- Molten Core and Blackwing Lair have no door check in the data (only Onyxia's amulet is tracked); the prep shows only checks it can read.

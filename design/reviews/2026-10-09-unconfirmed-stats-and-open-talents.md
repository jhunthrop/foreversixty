# What is left unverified, and why no table can settle it (2026-10-09)

## The 1.12-stats items

126 of the level-60 raid picks (publish 27) wear items whose stats come only
from the 1.12 database supplement: their id is in the client's base Item
table, but their ItemSparse row is not in the CDN export (it arrives as a
runtime hotfix), the hotfix dump taken from a 70009 play session does not
hold it, and Wowhead's Forever planner payload no longer lists it (12,060
items on 2026-09-27, 10,678 on 2026-10-08). Wowhead's per-item Forever page
returns 404 for such an item (Battleborn Armbraces, 12936) while a confirmed
item (Boots of Heroism, 21995) resolves, so Wowhead carries no Forever
record of them at all.

What the site does about it: the picker publishes a 1.12-stats item in a
slot only when it beats the best confirmed-stats item by more than the 1%
adoption margin, and lists it as the alternative otherwise
(`unconfirmed_within_error_held` per band); the card says "1.12 stats, not
yet seen in Forever's client".

What settles it: a hotfix dump from a play session that sees the items.
`python -m pipeline hotfixes --build 1.60.1.70291 <DBCache.bin>` reads the
client's cache; visiting Blackrock Spire, Blackrock Depths, Dire Maul,
Stratholme and Scholomance and the PvP and faction vendors pulls their rows
into it. The next fetch then merges them (SEEDED_FROM is not written for a
dump of the build's own session, so it merges fully).

## The two talent texts the rows cannot settle

Improved Slam (one spell, 12862, for both ranks: cast time −500 ms,
cooldown −3,000 ms) and Thick Hide (one spell, 16929, for three ranks:
base 3 with 3 per level, 180 at 60). Neither row carries a per-rank term.
They do not touch a published list: the Arms raid build takes Improved
Slam 0/2, and both feral builds take Thick Hide 3/3, where the engine's
180 armor at 60 is the client's.

## The client-side measurements that remain

design/reviews/2026-10-08-beta-evidence.md §1–§9: the hybrid crit sum
(character sheet, one minute), the white-hit inputs (energy rate,
dual-wield miss, glancing, Eviscerate's AP term, spirit regen, Shadowfiend
swing: the addon recorder, `python -m pipeline recorder`), and the hotfix
dump above.

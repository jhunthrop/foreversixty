# Pre-beta Forever data

Snapshots of the only real Forever game data that exists before the beta client on
2026-09-17. Committed raw and dated, never edited: the pipeline reads them, and after the
17th they become the record of what we believed beforehand.

## wowhead-talents-2026-09-14.json

Wowhead's Forever talent payload, fetched from
`https://nether.wowhead.com/forever/data/talents-classic?dv=100` on 2026-09-14 and unwrapped
from its `WH.setPageData(...)` envelope. Wowhead's internal name for the environment is
`classicplus`.

What it contains, measured:

| | |
|---|---|
| Trees | 27, three per class, named e.g. `WarriorArms`, `MageFire` |
| Talents | 470 |
| Talents with tooltip text for **every** rank | 470 of 470 |
| Spell id range | 104724 to 113398, all new Forever ids |
| Rows used | 0 to 6, so seven rows, matching Blizzard's statement |
| Talents per tree | 16 to 19 |
| Rank caps | 1, 2, 3 or 5 points |
| Talents with a prerequisite | 70 |

This supersedes every community talent dataset. Those were built by reading BlizzCon stream
frames with a vision model, which only ever saw rank 1, so their higher ranks are arithmetic
extrapolated from the Classic talent of the same name and are wrong wherever Forever retuned
one. This payload has Blizzard's own text for all four ranks of Pandemic, all five of
Malevolence, and so on.

## What this source does NOT give us

The same environment's gear planner
(`https://nether.wowhead.com/forever/data/gear-planner?dv=100`, 5.0 MB) is **still Classic
Era data** as of this snapshot, so it is not committed. Checked on 2026-09-14:

- Items carry `versionNum: 11300`, patch 1.13.0.
- Lionheart Helm still has the split `mlehitpct`, `rgdhitpct`, `mlecritstrkpct` and
  `rgdcritstrkpct`, not Forever's unified hit and crit.
- Edgemaster's Handguards still grant `skillBuff` 7 per weapon, where Forever's panel card
  showed 1.
- Hide of the Wild has `splheal: 42` and no spell damage, where Forever's card showed 42 and 14.
- None of the genuinely new items exists: no Graverobber's Shovel, Worgenbane Talisman,
  Ankle Slicers, Ladimore Heirloom Ring or Sharpened Cutlery.

So: **talents are Forever, items are not.** Anything gear-shaped waits for the beta client.
`statToRating` in that payload is Wowhead's own generic stat-to-rating map, present in Era
too, and is not evidence that Forever uses a rating system; see `research/08-stats.md` §2.

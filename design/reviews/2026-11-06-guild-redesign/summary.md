# Guild redesign, 5-6 November 2026: what shipped

Spec: `design/specs/2026-11-05-web-layout.md` §7. Mocks: `design/mocks/web/web-guild-*.html`. Shaped by three persona reviews (guild master, raid leader and loot officer, member), all built on the shared `ui/` primitives.

## The first raid tier

Barrow Deeps (10), Hyjal Summit (20) and Onyxia's Lair (40) open on 9 December 2026 (`data/curated/raids/first-tier.json`, `forever-raid-phases.json`, Blizzard's What's Next recap). Molten Core isn't in that tier and doesn't appear anywhere in guild code. The 70790 client carries Onyxia only (map 249, encounter 1084) and has no Dungeon Journal. Barrow Deeps and Hyjal Summit join the boss table once a client build names them (`pipeline raid-bosses`, re-run by `tools/activate-build.sh`).

## Slices

The deployed API applied every migration at startup. No tool touched the production database or Discord.

| Slice | Main | Migration | What |
|---|---|---|---|
| 1 Overview | ca9f0c710 | 0072 | Role-aware Overview: the next raid with your own In / Late / Bench, officer Action needed with Ping on Discord, comp against the target, buff and debuff cover with icons, the last four raids' attendance, the member's private checklist; signed in, `/guild` opens your guild; Segmented, Grid, Tally and Checklist primitives |
| 2 Raid planner | be1cd0a40 | none | Grouped sign-ups, officer in-row state, Not signed with Ping, comp target editing, bench and swaps |
| 3 Loot council | 26abdb297 | 0073 | Votes, MS/OS, award with a reason, decision record, history export |
| 4 Raid report | 70c4045b3 | 0074 | Kills, wipes, attempts and deaths; death causes; DPS and HPS; consumables by pull; officer attendance fix; Discord summary; the 360 px sign-up row fix |
| raid-route | 556ed120f | none | Raid pages answered 404 at the edge, including every Discord raid-post link; now fixed, with a test over every guild link |
| guild-migrations | 09dac9eb8 | 0070, 0071 | The invite link can be shown again (sealed token, inviter named); raid-less loot awards no longer overwrite each other |
| 5 Recruitment | b0c3bab9f | 0075 | Recruiting card, applications with votes, accept and decline with the Discord role, member profile, officer notes |
| 6 Data, export, Raid-ready | 49357b28c | none | Boss table from the client; the addon raid-prep export (resistances, supplies, keyring, attunement quests, Drakefire Amulet); officer Raid-ready counts; the 22 guild e2e failures that were on main |
| 7 Drops, assignments | c4e4a5de2 | 0076, 0077 | The addon records each boss corpse's loot and its winner; manual assignments with Onyxia's 12-slot template; Post to Discord, Copy as raid warning; `/fs job`, the fourteenth command in `/fs help` |

## Decisions the owner delegated (5-6 Nov)

- Sign-up words stay In / Late / Bench.
- A member's readiness is a private checklist: Ready / Almost / Getting started. There is no public worst-first list.
- New stores were approved.
- The addon export gained the raid-prep sections.
- Buff cover comes from a curated catalogue and shows HD icons.
- The bot @mentions non-signers in the raid channel. DMs are opt-in only.
- Assignments are manual per-encounter templates, with no auto-solver.

## Owner checks in the beta client (`addon/README.md` spike rows)

- Rows 98-102: the raid-prep export sections.
- Rows 103-105: the drops capture (`/fs drops`, `GetLootSourceInfo`, the `LOOT_ITEM*` strings, and two raiders on one corpse giving one row).
- Assignments: `/fs job` after a companion sync, the line when zoning into Onyxia's Lair, and Post to Discord arriving as a reply with no pings.

These reach players with the 2.0.29 addon tag. Until then the site reads 2.0.28 exports as "not shared yet", never as an error.

## Open

- Barrow Deeps and Hyjal Summit bosses, loot and assignment templates wait for a client build that carries them.
- `RaidComposition.targets` overlaps `comp_target`. Retiring it is on the mobile list.
- Refresh hold: while `data.yml` is disabled, a new client build still stops at that job's other outputs.
- Not built because there's no source: "Officers asked you to", the guild's wanted count per class, a character switcher on the profile, a picker for the recruit Discord role (officers type the id).
- The companion's `internal/app` tests read the real WoW install on the owner's Mac and fail there only. CI is unaffected.

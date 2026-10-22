# Round 10 walk — one character across the site (21 Oct 2026, local stack, hunthrop#1894, Bow Jackzon chosen)

Dev server 4321 on main at `07d34f5d` with the island scripts rebuilt against the local API (`PUBLIC_API_BASE_URL=http://localhost:8080 npm run build:island`, copied into `web/public/`, not committed); API 8080; browser pane at desktop width. Pointer read from `localStorage['fs.currentCharacter']` after each stop.

| # | Stop | Read | Pointer after | Verdict |
|---|---|---|---|---|
| 1 | Header selector → Bow Jackzon | chip "Bow Jackzon · Survival · 60", toast "Now using Bow Jackzon." | `armory us/pvp/bow-jackzon`, label "Bow Jackzon · Survival hunter" (lowercase class) | OK; label wording → W.11 ruling 4 |
| 2 | `/planner` (bare) | class Hunter, race Dwarf, level 60, "Your character" card right; **0 of 51 points, trees empty**; "Sim this build" → `/sim?code=FS1:…:hunter:dwarf:0/0/0:` | unchanged | the defect (rulings 2, 3) |
| 3 | that "Sim this build" link | the simulator's landing (paste box); "Your characters: sign in" | unchanged here (the island had the production API base until rebuilt) | ruling 3; on production this writes the "Hunter" placeholder (ruling 1) |
| 4 | `/sim?source=armory&ref=us/pvp/bow-jackzon` | card "Bow Jackzon · Dwarf Survival Hunter · 60 · 25 Talent Points", gear filled, badge "Addon export, 13 hours ago", strip "Bow Jackzon · PvP · US · Plan · Sim · Logs · Rankings for Hunter · Forget" | label rewritten to "Bow Jackzon · Survival Hunter" | OK — the one path that works |
| 5 | sim strip "Plan" / card "Open in planner" | `/planner?class=hunter` / `/planner?code=FS1…` (gear + talents) | — | rulings 2 and 1 |
| 6 | `/bis` | "Your character: Bow Jackzon, Dwarf Survival Hunter, Level 60, Alliance, Addon · 13 hours ago; 3 of 17 slots" | unchanged | OK |
| 7 | `/rankings` | card Bow Jackzon; "Not ranked yet on Onyxia"; board "Survival Hunters" | unchanged | OK |
| 8 | `/tiers` | "Survival Hunter is 6th of 20 DPS specs" | unchanged | OK |
| 9 | `/character/us/pvp/bow-jackzon`, `/reports/<id>` | "No page here" on the dev server (served through `_redirects` in production) | — | dev only; capture on a preview |
| 10 | guild roster "Gear in the planner" / "Open in simulator" | `/planner?code=FS1…` / `/sim?code=FS1…` (the member's export) | — | ruling 1: must not replace the pointer |
| 11 | `/account` current character | "Bow Jackzon · Dwarf Hunter · Level 60 · PvP US"; "Open in simulator" and "Open in planner" are `?code=` links with the export | — | ruling 3: the export should load as what it is (addon export), never as a manual code |

Dev-server facts for anyone walking this: `web/src/lib/items/lookup.ts` imported node:fs and crashed every sim/planner island on `npm run dev` (fixed `07d34f5d`); the sim pages need `npm run build:island` output served (and built with the local API base); dynamic character and report routes 404 on the dev server.

## After W.11 (`59ed53cd`) and the seed fix (`04d8585a`), same stack, Bow Jackzon chosen

| # | Stop | Read | Pointer after | Verdict |
|---|---|---|---|---|
| 12 | sign in (dev) with no pointer | the first `/v1/me` answer settles the pointer to the main (Pipsqueak) | `armory us/pvp/pipsqueak` | OK (ruling 7) |
| 13 | selector → Bow Jackzon → `/planner` (bare) | "Hunter talents · High Order Skyborne · Level 60 · Beast Mastery Hunter", race select `high-order-skyborne`, **51/51** (31/16/4), DPS 102, card "High Order Skyborne Beast Mastery Hunter · Level 60 · Alliance · Addon · just now", button "Send this build to Bow Jackzon", note "Talents loaded from a character…" | `armory us/pvp/bow-jackzon`, label "Bow Jackzon · Beast Mastery Hunter" | OK |
| 14 | planner "Sim this build" | `/sim?code=FS1:1.60.1.70450:hunter:high-order-skyborne:<talents>:<gear>\|level=60\|who=Bow Jackzon:`; the simulator card "Bow Jackzon · High Order Skyborne Beast Mastery Hunter · 60 · Build from planner · Bow Jackzon's gear", the production gear, "Talents 51 points"; chip unchanged | unchanged (armory) | OK — the owner's defect is gone |

The two earlier oddities (Orc, 0 of 51) were the seed: every seeded export declared data build 1.60.1.70009 (whose race table has no Skyborne) and its talent digits were not legal ranks. Production exports carry 1.60.1.70450 and per-cell ranks, so production was never affected by those two; the seed now carries the owner's production head with the level-60 Beast Mastery build.

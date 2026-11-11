# Round 3 findings: what we change in round 4

Round 3: guild-officer WORTH PAYING FOR (passes); info-density DENSE WITH ADDITIONS; brand DRIFT;
ux-designer NO SHIP; ai-slop FAIL. Codes: ID = info-density, GO = guild-officer, BR = brand, UX =
ux-designer, AS = ai-slop (round-3 sections of each verdict file).

## Both apps
| # | Change | From |
|---|---|---|
| 1 | Covered line: never an empty icon frame or gap. A name with a spell icon draws it; a name without one draws none and the text closes up. Short names from the data's `label` (the fixture gains `Kings`, `Might`, `Wisdom` labels and the right spell icons, see s2). One separator ` · ` on both apps from one builder; no mid-name wrap (each name is non-breaking) | UX2, BR N3, ID F1, AS4 |
| 2 | Stale tone: the `Synced … ago` stamp under Next upgrades and Gear's `synced … ago` take the warning tone past 7 days on both apps, as the drawer's stamp does; one `isStale` rule | UX1, BR N2 |
| 3 | Gear: a slot that is neither Equipped nor has a gain reads `Not modelled` (when `modelled` is false) or `+0.0 DPS`; the header keeps `· N gains not modelled yet` when N > 0 | ID F2 |
| 4 | The answer control at large text stacks 2 x 2 on both (iOS from accessibility sizes, Android from 1.5x), the stacked form at radius 4 | BR N5 |
| 5 | Who is coming's Damage row: `Damage 15 · 7 melee, 8 ranged` on one line, the names under it, on both (the Android form) | BR N6 |
| 6 | iOS Home tally sits 8 pt above the Segmented, as Android | BR N7 |
| 7 | `Show older` left-aligned under its list on both, 44 pt / 48 dp target | BR N8, UX polish |
| 8 | Hero line at large text breaks at the dot: `Level 58 Fury Warrior` then `2 levels to the 60 list` (never strands `· 2`) | UX polish, AS polish |
| 9 | Activity filter chips: the selected chip fills `teal` with a `tealEdge` hairline on iOS too (as Android and the Segmented); a trailing fade where the row scrolls | BR N1, UX polish |
| 10 | Drawer: the pressed state of the stamp line is on the `Refresh` link only (48 dp / 44 pt), not the whole row | UX polish |
| 11 | Copy: cut `home.answer.noGearNote`; `raid.noRaid` and any other no-raid pitch variant on these screens read `No raid posted yet.` or are deleted when unused | AS polish |

## iOS
| # | Change | From |
|---|---|---|
| i1 | Type residue on the tab roots: `GearListView.swift:97, 107` on roles | BR N4 |
| i2 | Captures: the officer's Raid tab (Needs you, not-synced, Covered), the filtered roster after tapping the not-synced line, the drawer's `Refreshing…` frame, the raid push tap-through landing (`xcrun simctl push` with the fixture), and the officer pushes (raid.short with Mention, the result; guild.applicant with Accept and Decline) | GO N9, P2, P4, UX evidence |

## Android
| # | Change | From |
|---|---|---|
| a1 | Type residue on the tab roots: `FlowFigures.kt:45, 55`, `FlowLayout.kt:304, 307` (ACTION, PANEL_TITLE), `FlowRows.kt:338, 345` line heights, `ForeverTheme.kt:52` TOP_BAR_TITLE (barTitle), `PaperdollHeader.kt:301-304` on roles (add `panelTitle` if no role fits) | BR N4 |
| a2 | A debug intent that posts any push fixture (generalise `RaidDebugPush` to take the fixture name), so the walk captures raid.short with its Mention action and its result, and guild.applicant with Accept and Decline; the filtered roster after the not-synced tap; the drawer `Refreshing…` frame | GO P2, P4, UX evidence |

## Server and fixtures
| # | Change | From |
|---|---|---|
| s1 | One short-role rule for the Mention action and Needs you: the first short role in tank, healer, dps order (the role an officer can fill), not the largest raw gap; the app's Needs you line and the push name the same role | GO P1 |
| s2 | Fixtures: `guild.raid*.json` cover rows and `composition.buffs` carry the right spell icons and short labels for the blessings (Kings, Might, Wisdom; the HD icon names the site uses); the raid push's body names the same raid and time as its title (`Onyxia's Lair`, `Sat 19 Dec, 7:00 pm CST`); the feed's Molten Core row dates a raid after the fixture clock | AS4, UX3, AS polish |
| s3 | One DPS scale in the fixtures: the sim rows and the guild-sim row sit on the scale of the Gear header's band list (Gorvash, Fury 58: about 190 DPS, not 1,184); the warrior home's gains total about +18 DPS; `guildsim.passed` reads `+12.2 over your 191.4` (formatted like the figure, no class parenthetical) | AS2 |
| s4 | `raid.short` body uses the actual lead (not a hard-coded `24 hours`); `raid.starting` has a singular (`1 minute`) | AS polish |
| s5 | `Not synced this week` must be true: if `logged_recently` is a game login and not an addon sync, compute the line from the addon's sync time (`logged_at` if that is the addon's, else the roster's `addon` / `sim_export` fields) or rename the line to what the field means; say which | GO P3 |

## KEEP (with the ruling)
- **Panels on the list blocks** (AS round-3 must-fix 1, tell 8): the round-5 system (design/specs/2026-11-09-directions.md and the redesign plan, Wave 3) puts each block in a dark panel on the earth ground, on the web and the apps alike; the web home and guild pages draw the same panels (design/reviews/2026-11-10-mobile-backfill/web/), so bare rows would be drift from the site the brand reviewer holds us to; the same panels passed ai-slop on these screens in M1a to M5a (design/reviews/2026-11-10-m1a ... m4a-m5a/ai-slop-reviewer.md). The one framed Plate stays the only double border.
- **`OLYMPUS XXVII` in the loot body** (BR carried N-1): the guild's own name as its officers typed it; the app never restyles a name.
- **Home's Damage without the split** (ID F3): the card is the short form; the Raid tab carries the split.
- **Android's three notification actions** (GO P5): platform limit (rule 8); Out is one tap away on the raid page.

# Round 2 findings: what we change in round 3

Sources: round 2 of info-density (DENSE WITH ADDITIONS), guild-officer (CLOSE), brand-consistency
(DRIFT), ux-designer (NO SHIP), ai-slop (FAIL); each verdict file beside this one. Codes in brackets
name the finding (ID = info-density, GO = guild-officer, BR = brand, UX = ux-designer, AS = ai-slop).

## Both apps
| # | Change | From |
|---|---|---|
| 1 | One count model: `Damage N` equals melee plus ranged on every surface (Home card, Raid tab); one derivation shared by Home and the Raid tab; a test asserts melee + ranged = total and Home equals Raid | UX1, BR D4, AS4 |
| 2 | Home's raid card shows the tally on iOS too (`Healers 6` without a target, `Healers 5 of 8 · 3 short` with one); Home's compact form carries `· 3 short` | UX2, BR D2, ID N4 |
| 3 | Home's New for you ages use Activity's bucket formatter (`Sat`), on its own meta line, never `· 3 d ago` wrapping | UX3, BR D3, UX polish |
| 4 | Covered line: neutral label, `Missing` once, buff short names from the data's `label` (`Fort`, `MotW`, `Kings`...) else the name, at most three, ordered missing before short, then `and N more`; the spell icons beside the names (ID N1). Tone: wipe red only on the names, never a red wall | GO N7, AS polish, ID N1, UX polish |
| 5 | The officers' not-synced line names them: `Not synced this week: Kaelra, Thalrin, Brakka, Mirethil` (three then `and N more`), and a tap opens the roster filtered to them (Roster is reachable from the tab already) | GO N4 |
| 6 | Counts agree: the Needs you `5 have not answered` and the foot `N not answered` count the same population (the raid's not-answered list); if the fixture disagrees, fix the fixture | GO N5 |
| 7 | One answer state per surface: the Raid tab head drops the answer chip beside the title (the Segmented is the state); Home's card drops the `You are In` line when the Segmented shows it | AS5 |
| 8 | Synced stamp sits under the section meta line on both (`13 upgrades · +157.6 DPS` then `Synced 2 h ago`), outside the inner frame | BR D7, UX, AS2 |
| 9 | The raid notes (`Bring flasks`) sit inside the raid card on both | BR D8, UX |
| 10 | Gear: the counter, the Equipped cells, the header's upgrade count and the gain cells agree (reconcile against `upgrade_count` and `already_best[]`); slot labels from `slot_label` (`Finger 1`, `Trinket 2`) in cells and sheet; the separator joins its node for TalkBack | ID N2, N5, N6 |
| 11 | Activity rows: the figure shows once (the server stops writing it in the body where `figure` carries it; the app formats it like the rest, `1,196 DPS`) | ID N9, AS polish |
| 12 | Drawer Refresh shows an in-flight state on both (`Refreshing…`, the link disabled), then the new stamp or the error; the stamp turns to the warning tone past 7 days with `Refresh` as the action | UX5 |
| 13 | Copy, ai-slop AS6/AS7 and polish: delete unused keys (`menu.tool.*.line` if not drawn, the duplicate pairs listed in the AS verdict); cut `home.answer.note`, `gear.intro` if drawn to the bare value; `gear.whereTonight` reads `Where to go`; one no-guild line instead of four (`home.findGuild.line`, `raid.tab.outLine`, `raid.tab.withGuildLine`, `raid.tab.noRaid`, `home.findGuild.firstRaids`: keep one sentence per surface, no pitch); `home.head.toBand` says `2 levels to the 60 list`, never `band`; `Saving…` with the ellipsis character. KEEP with ruling: `Updated 2 h ago from Battle.net · Refresh` (owner wording, spec R15); `From your last log-out, 9 h ago` (says where the data came from, which the player needs to trust Before you pull) |
| 14 | Plan row shows the founder price end date when the account has one (fixture me.premium.json); if the API lacks the field, add `founder_until` on the server | ID N3 |

## iOS
| # | Change | From |
|---|---|---|
| i1 | Activity gets `Show older` (same copy as Android, 44 pt), loading the next page | UX6 |
| i2 | Countdown, drawer and other residue on type roles: `NextRaidBlock.swift:82` uses `countdown`; `ProfileMenu.swift` inline TypeRoles use `drawerName` / `drawerSection` / `tag`; `NextRaidBlock.swift:257`, `HomeFlowView.swift:198` use roles | BR D6 |
| i3 | Content gets a bottom inset so nothing scrolls under the tab bar unreadably | UX polish |

## Android
| # | Change | From |
|---|---|---|
| a1 | The chosen segment fills `FsColor.teal` with a `tealEdge` edge, as iOS and the web primary | BR D1 |
| a2 | The answer control stacks (2 x 2) from font scale 1.5, as iOS does at accessibility sizes; the Late chips carry `How late?` on the Raid tab as on Home | BR D5, UX polish |
| a3 | Type residue: `ProfileButton.kt` BAR_TITLE, `Paperdoll.kt:167-168`, `FlowRows.kt:251, 356, 367` on roles (add `barTitle`, `paperdollName`, `paperdollDetail` to tokens.json if no role fits; `countdown` for `FlowCountdown.kt` CLOCK) | BR D6 |
| a4 | Drawer `Refresh` on its own row at font scale >= 1.5 (48 dp) | UX polish |

## Server (api)
| # | Change | From |
|---|---|---|
| s1 | `raid.short` carries one action, `Mention <role>` (the shortest role below target), posting the existing raid ping (`RaidPingResult`), honouring `next_allowed_at` (no action when inside the cooldown); category `RAID_SHORT`; both apps handle it | GO N1 |
| s2 | `guild.applicant` offers Accept beside Decline (`application_accept`), or neither | GO N2 |
| s3 | Push bodies lead with the count: `24 of 40 signed. Thu 5 Nov, 8:00 pm CST.`; `raid.short` leads with the short roles; the apps set big-text style so the expanded view shows all | GO N3, UX polish |
| s4 | The `answers` blob uses the same time layout as the body; `moved` has a title and no leading space; `cancelled` reads `<Raid> was cancelled.`; `push.applicant.failed` reads `Open the application.` | GO N6, AS8 |
| s5 | Activity bodies stop repeating the figure the row now carries (`sim.finished`, `guildsim.passed`) | ID N9 |

## KEEP (contested by ai-slop round 2; the ruling)
- **The iOS tab bar's Liquid Glass** (AS1): iOS 26's system `TabView` draws it; spec 4.1 asks for a standard
  TabView, and rule 8 (brand) keeps native navigation. A custom opaque bar would be the web forced onto
  native. Change i3 removes the unreadable overlap, which is the part a player feels.
- **The iOS large title band** (AS3): ruling N3 and the redesign plan's "iOS large titles"; the band art
  is the round-5 look the owner approved (Wave 3). It collapses to the inline title on scroll.
- **The Plate frame on Next upgrades** (AS2): the round-5 PanelFrame is the one framed panel per page
  (spec Plate row: one per screen). Change 8 takes the stamp out of it; the frame stays.
- **Gear's `182.0 DPS` against `+157.6 DPS`** (AS4): the header figure is the band list's own DPS
  (`set_dps` for 50 to 59), not the character's; the upgrades' gains came from a hand-made fixture and
  are scaled down in round 3 so the numbers read true (fixture fix, both apps).

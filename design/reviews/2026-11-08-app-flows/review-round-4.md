# App flows review, round 4 (tabs Home, Raid, Gear, Activity)

wow-player as four personas on the recaptured mocks, with the spec's purpose audit (appendix A).

| Persona | Verdict |
|---|---|
| Raider (60 Fury Warrior, raid day at work) | HOLD |
| Officer (fills healers for a 40-man) | HOLD |
| Leveler (34 Beast Mastery Hunter) | SHIP |
| Casual (twice a week, pushes mostly off) | SHIP |

Tap counts: every counted flow is within budget. Late from launch was not mocked yet.

## Must-fix and disposition
| # | Finding | Disposition |
|---|---|---|
| R1 | Answered state is a dead end (no Change) | Answers stay with the current one selected; `home-raid-day-late` mocks Late's chips in the app |
| R2 | Contradicting numbers: 87th vs 94th; widget +48.1 summed | One number (87th, your best on Onyxia); the widget shows the next upgrade, never a sum |
| R3 | Tonight line on raid day, and it names a smaller upgrade | Left out on raid day; elsewhere `Best run: ..., Blackcrow +1.9 DPS, 8%` with the gain |
| O1 | The screens do not help fill healers | Mention the 2 healers first, Mention all 5 second; Bench and Out filters; Healers on the bench with Mark In |
| O2 | What Accept does is unclear | `Accept into guild`, the role line, `Accepted into Olympus XXVII` |
| O3 | No change path after Mark In | `Marked In · Undo` on the row |

## Never look at, and disposition
| Element | Disposition |
|---|---|
| Home raid day: Tonight line | cut on raid day |
| Home raid day: Healers short row in the preview | cut (the card says it) |
| Home day after: preview repeating Last night | the preview leaves Last night's kinds out (`exclude_kinds`, API) |
| Home raid live: "starts in 60 min" row | replaced with a live change |
| Gear: Next upgrades block | cut (Home holds the top three) |
| Activity: group label per row | cut |
| Menu: Plan, Devices, Offline data, version | kept: account chores (F10), out of every flow |
| Activity filters (casual) | kept: All preselected; raider and officer narrow to Raids |
| Raid tab guild pages (casual, leveler) | kept for raider and officer; a leveler gets `Raids start at 60` instead |

## Shoulds taken
E7 on the next-raid line; Ready check with boots and the export stamp; Stormpike `needs revered`; item sheet without item level, with the gap; Pick your main points at the name on Home; notification settings say Gear and sims and Tiers, Officer only for officers; Home's preview renamed New for you, with unread dots and the count on the link; the menu tool renamed Patch notes; the chip row fades at its edge; scenario date moved to Wed 6 Jan so history lines fit.

## Shoulds left for the build lanes
Leveler dungeon level fit in Best run, quest zones on upgrade rows, the Forever Score scale on the applicant, the healer comparison for an applicant's HPS, a digest option for casuals, verifying the item names Greenhammer, Ironwood Blade, Skibi's Pendant, Timbermaw Tunic and Howler's Furs.

# Round 5

| Persona | Verdict |
|---|---|
| Raider | SHIP |
| Officer | SHIP |
| Leveler | SHIP |
| Casual | SHIP |

Every flow within budget (Late with minutes from launch: 2). No must-fixes.

Shoulds taken after the verdicts: Home's `All activity (N)` is the Activity badge's number; on the Not answered page the counts follow Grunvald's Mark In (Not answered 4, In 19, Mention all 4); the roster says 31 at 60, matching the raid's 31; a line says the bench fills the healer gap if all three answer In.

Never-look-at items left in on purpose: the menu's account rows (F10); the read third row of New for you (it shows the casual the list is current); the Best run line for a raider on other days (a leveler and a casual use it); Activity filters for a casual (All preselected).

Left for the build lanes, as round 4: leveler Best run level fit and chance, quest zones, Forever Score scale, an applicant's HPS against the current healers, a notification digest, the item-name check, talents for levelers, the leveler's Gear capture.

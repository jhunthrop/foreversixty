# ai-slop-reviewer, Me tab (L2)

## Round 1

Verdict: FIX (3 small copy items; no visual tells 1, 2, 5, 8, 12, 13 or 15 found)

First impression: reads as a WoW player's addon. Real boss names, parse colours by bracket, real item icons, one big number (1184) that is the point of the tab, flat bordered panel, no glow, gradient, glass, pills or icon grids. Not generated-looking.

Must fix
1. Tell 22/21, Right now Rotation row (me-raid-day, me-stale-no-guild): the sub-line "Your weakest boss on Tuesday" explains what the title already says (Garr, 74% followed, 31 s downtime). Replace with data, e.g. "Tue 8 Dec · parse 58", or drop the line.
2. Tell 22, Last raid table, Consumables column: "All up" repeated on six of eight rows is filler. Show nothing (or a dim dash) when fine; keep red "No flask" as the only thing that speaks. The Died column already does this with a plain 0.
3. Tell 21, Locale.lua meLogLead / meAppLead (and the Copied variants): "Press Ctrl+C, then paste it in your browser." explains the obvious to a veteran. Cut to the URL plus "Ctrl+C to copy"; the Copied line can be "Copied."

Non-blocking notes (not slop, fix if cheap)
- me-new-character: the "Pick your spec" row has no icon and sits indented where the icon would be, so it looks like a layout gap. Give it an icon or left-align it.
- "Mon 7 Dec" under Your sim has no label; it reads as an orphan. "Simmed Mon 7 Dec" is a qualifier trap, so prefer placing the date on the header line right of "Your sim", as Forever Score does with "as of".
- The top gain (+11.2 DPS gloves) appears in both Right now and Your sim. Acceptable on a summary tab; just do not add a third restatement.
- Right column has a large dead area under the sim block on raid day. Fine, but do not fill it with decoration.
- The renders use a Noto/Open Sans-style face; confirm the shipped addon uses the in-game fonts, since the default-safe font tell (14) would apply if this is what players see.

Authentic, keep
- 1184 Rare with "316 to Epic (1500)" and the five bars with "of N" values; the bars are coloured by role, not decoration.
- Last raid table: per-boss parse in bracket colours, "Hunter's Mark late 6 s", "Aimed Shot late 9 s", "Serpent Sting off 4 s". Class-specific and concrete.
- Right now rows with Onyxia's Lair 20:00 "Phase 2 ranged, west side", Gloves of Holy Might "in your bags", routed by a plain yellow tab name.
- Stale state turns the date and footer red ("as of 2 Dec", "Ratings 8 days old") instead of adding a banner.
- No-pack and no-companion states are one short line each, with no skeletons or fake content.
- Footer states "No companion" and "Companion synced 14 min ago" are real state; tab dots only appear with something behind them.
- No em dashes in the me* copy.

## Round 2

Verdict: PASS

1. Rotation sub-line: now "Tue 8 Dec · parse 58", data not explanation. Resolved.
2. "All up": accepted. It is muted, the red "No flask" rows still carry the column, and the owner-approved hub-me.png and spec 4.1 name it. A blank cell would read as missing data, which is a fair argument. Keep it muted.
3. Copy-box leads: accepted. One copy box speaking one way matches the shipped /fs link lines. If the shared dialog copy is ever revisited, drop the "Press Ctrl+C, then paste it in your browser" sentence everywhere at once.
4. Non-blocking items fixed: the "Pick your spec" row now starts at the column edge, Biggest gains names each item in quality colour, fonts confirmed as shipped files.

Minor, not blocking
- "Simmed Mon 7 Dec": the word "Simmed" is the qualifier the context already gives (it sits under Your sim next to Sim again). "Mon 7 Dec" alone would do. Take or leave.
- "51 of 51 points on plan": "points" is new and the context makes it clear; "51 of 51 on plan" was cleaner.

# ai-slop-reviewer: Guild tab (iOS + Android), 2026-11-11

Verdict: FAIL (tell 8 is present; it blocks a pass). Copy is mostly clean and the fixes are small.

First impression: the data and words read as written by a player (Whelps west cave, Missing Fort / WF, Bring flasks, Drakefire Amulet, Mirethil healers in class colour). The layout reads as a card stack: every block is the same near-black rounded bordered box on a parchment backdrop.

## Findings, most damaging first

1. Tell 8, same card on every block. Both apps, every state: the guild head, raid head, Needs you, each Officer group, Assignments, Who is coming, Before you pull, the guild banner, Roster, Loot, Recent logs, Progression and Pages are all one dark rounded-rect with the same border and radius. A one-line block (Loot "Last raid: 4 pieces", Progression "First tier / 1 of 1 down") gets a full card of its own. Fix: put the list sections (Officer, Assignments, Who is coming, Recent logs, Pages) on the page as hairline-divided rows with no box; keep a fill only for the one thing that matters (the raid head with the In/Late/Bench/Out control). Single-line sections (Loot, Progression) become a heading row with the value on the right.
2. Tell 22, qualifiers the context already gives:
   - `Seen just now`, `Seen 12 min ago` under the heading "Recently on": drop "Seen".
   - `6 open` beside "Needs you": the rows say what is open; drop it or make it the badge number only.
   - `6 pulls this tier` under "Progression": "6 pulls".
   - `From your last log-out, 9 h ago` under "Before you pull": "9 h ago".
   - Assignments card repeats its own title inside it (`Onyxia assignments` under the "Assignments" heading): keep the boss, drop the word, or drop the section heading.
   - `Healers short by 3` (Needs you) and `Healers 5 of 8 · 3 short` (Who is coming) say the same thing twice on one screen. Keep it in Who is coming.
3. Tell 7, grey explainer lines. `Claim it: a guild master or an officer confirms with the addon or Battle.net.` under "Claim your guild" (the title already says claim; "Claim it:" is a stutter and the sentence explains the flow to a veteran). `By name or realm` under "Search guilds on the site". `Assigning is on the site for now.` is an apology line, and "for now" is a roadmap note. Fix: delete the three; if the Assignments rows are read-only the Open in planner link says enough. Rename the first row "Search guilds" (the browser hand-off needs no label).
4. Tell 19/vocabulary: `Preview mention` on Needs you is the only button on the page and a player cannot tell what it does (mention who, where). Fix: name the noun, e.g. `Ping 5 on Discord`, or whatever it actually sends.
5. Tell 16, the 30 / 5 / 58 spread in the guild head: three equal columns of big numbers. It earns its place only under 60 (it is the point of the page for a leveler). At 60 the head should not carry it. Verify the at-60 render shows no stat row; if it shows one, drop it.
6. Tell 11, minor: a `Roster ›` link appears twice (Who is coming, and the Roster section). Pick one entry point. Text chevrons (`Roster ›`, `Loot history ›`, `Open in planner ›`) sit next to row chevrons in the same screen; use one convention.
7. Tell 2, minor and not this lane: the sky-to-sand gradient backdrop with texture and the glossy brown gradient-bordered tab capsule on iOS carry no information. It is the app-wide brand backdrop, so route it to the brand owner; it does not decide this verdict.

No em dashes found in the captures or in guild.tab.* / shared.assign.*. No glow buttons, no uppercase eyebrows, no emoji, no pill badges beyond real count badges, no fake social proof.

## Authentic, keep
- Real game nouns everywhere: Onyxia's Lair, Whelps west cave, Fear cover, Main tank healers, Missing Fort / WF with spell icons, Drakefire Amulet, Flask of the Titans with Auction house link.
- Class-coloured names in rosters and assignments; "Tanks 2 of 2 / Healers 5 of 8 · 3 short" is how an officer talks.
- Uneven hierarchy: the raid head with In / Late / Bench / Out dominates; Officer rows show their count at the right (`1 waiting`, `4 of 8 done`) with the steps still to do.
- `Raids open at 60 · you're 58` is short and exact. `Bring flasks` with the noun bold.
- Copy keys are terse (`None planned`, `None waiting`, `{n} of {total}`).
- Fixtures: `Test Applicant` appears in the captures; fine as a fixture, keep it out of any store screenshots.

## Required to pass
Findings 1 to 4 (1 is the blocker). 5 to 6 are fixes on the way. Re-capture iOS and Android officer, leveler and no-guild after.

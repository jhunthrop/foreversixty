# Guild tab: response to review round 3 (2026-11-11)

Verdicts: guild-officer WORTH PAYING FOR (conditions: full-page captures, Ping after-state); ux-designer NOT YET (one
must); wow-player NOT YET (could not see the round-3 sections); ai-slop FAIL (four required items). Round 2's
info-density DENSE and brand-consistency CONSISTENT stand: round 3 adds rows inside existing sections and one tone.

## Fixed in round 4

| Item | Fix |
|---|---|
| All four: the "full-page" renders were one screen tall | `parity/ios-full-*.png` and `parity/android-full-*.png` are now whole pages (ImageRenderer at 2x on iOS, a 4600 dp Robolectric window on Android), every section from the raid head to Pages |
| ux must 1: the leveler render showed targets (`5 of 8 · 3 short`) under 60 | Under 60 Who is coming shows the counts alone (`Healers 5`), both apps (`countOnly`), with tests; the leveler mock follows. The device walk had no targets only because its fixture raid has none |
| guild-officer condition 2: Ping on Discord after-state | Both walks now tap Ping on Discord (`officer-07-ping-list`) and Mention (`officer-08-mentioned`): the result line `Mentioned 2 healers in Discord.`, and a repeat too soon reads `They were mentioned a moment ago. Try again later.` |
| ai-slop 4: `Gear gap 41 DPS` under a `Gear gaps` label | The row reads `41 DPS` |
| wow-player 7: copy and inventory disagreed | The inventory quotes the shipped `Assigning is on the site.` |

## Kept, with the reason

| Item | Why |
|---|---|
| ai-slop 1: delete `Assigning is on the site.` | The coordinator's brief for this lane: "Assigning stays on the web for now; say so plainly." The line is that statement, one short sentence beside `Open in planner ›`; removing it would break the brief |
| ai-slop 2 and 3: `6 open`, `Healers short by 3` in Needs you | Needs you is the at-60 raid block, kept unchanged by this lane's brief ("At 60: the raid first, unchanged"); it is shared with Home and owned by the raid lanes. Routed, as in round 1 |
| ai-slop: the log-out stamp wording | Before you pull is the same unchanged at-60 block; it is gone under 60 |
| wow-player 1, 5: Barrow Deeps, Mirelle and Mirethil | Fixture and contract data (`guild.raids`, `guild.applications`), shared with every raid surface's tests; not product copy |
| wow-player 2, 6, 8: Fort/WF unit, the note truncation, the whelps phase text | The raid block's own rows and the planner template's text, routed in round 1 to the raid and planner owners |
| wow-player 3: the guild head moves | Spec 2.1/2.2, owner-approved: under 60 the guild leads, at 60 the raid does |
| wow-player 4, 9: an attunement line for a leveler | Needs data (the attunement chain per character is not in the app contract); listed in the spec's needs-data list |
| guild-officer should 3: Missing raids and Gear gaps above Recently on | Follow-up; one order on both apps for every viewer keeps the section one shape. Raised for the next guild pass |
| ux shoulds 2 to 5 (echo keeps the rest of the roster, the `Saved.` shift, Android chevrons and hand-off glyphs, Calendar) | Follow-ups. The echo moves your own answer as the server does; the canned sign-up answer's raid has fewer sign-ups than the GET fixture, which is fixture data. `Saved.`, chevrons and Calendar are the shared raid flow components, unchanged by this lane |

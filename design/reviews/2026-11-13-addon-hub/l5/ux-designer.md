# L5 Guild tab: ux-designer verdict

Verdict: FIX (two must-fix, then SHIP). Judged on the five renders at 740 px against spec 3, 4.5, 8, 9, 11 and hub-guild.png. Code read only to confirm hover, click and Follow behaviour exist (GuildHubTab.lua: item tooltips and shift-link, name whisper/invite, Follow button); the renders do not show them.

## Must fix

1. Section headings are in the title face (Marcellus 14), on Tonight, Your prep, Sign-ups, Next raids, the guild name and Officers sent. Spec 3.5: "Marcellus only for the window title"; headings are semibold Open Sans in the strong colour, 18 px rows. The board shows bold sans headings. The mock generator's TITLE_CSS copies this from the Builds tab, so the Builds tab has the same defect: fix both in one place (the shared section-heading helper, Theme.titleFont at windowSection), not in L5 alone. Acceptance: l5-raid-day.png headings match the board's weight and face; Builds tab matches.
2. The guild line wraps to two lines in every render that has the pack ("3rd of all guilds · 41 raid nights in 90 days · 91% attendance", the orphan word "attendance" on line 2). Board and spec: one muted line. Shorten the pieces so the line fits the 288 px column (for example `3rd of all guilds · 41 nights · 91%`, or drop the "of all guilds" tail), or cut with an ellipsis and put the whole line in the hover. A wrapped line pushes the table down and makes the right column taller than the left in l5-no-companion.png. Acceptance: one line at 288 px in l5-raid-day, l5-queued, l5-old-companion (where `Ratings 8 days old` also sits on the line: keep it, shorten the rest).

## Should

1. `Late` renders in client gold in Next raids (`Late as DPS`) and in Your sign-up. Client gold means act now (spec 3.5). An answered raid should not ask for action. Use the body colour for Late and Bench, green for In, warning for Out, as in the Sign-ups tally.
2. `you are 9th of 15`: the spec's 31 is the guild size, built has the count of members in the pack (15). Say what it counts or use the real roster size once the ranks ship (spec section 9, last row). Until then keep `of 15` only if 15 is the count of ranked raiders; confirm the number is not read as the guild's size.
3. The own-place row (9, Velisandre) follows the top five with no gap. Add a 4 px extra space above it so "under them" reads as a separate row, and keep the green class colour. The board does the same without a gap, so this is polish.
4. `1 of 2 jobs filled` under the role line is in neither spec nor board. Keep it only if it is data-backed (it is), but then add it to the spec copy list; otherwise drop it. Not a blocker.
5. Tonight lacks raid leader and invite time (the build's own listed omissions: no message carries them). Accepted for ship; file the contract-3 addition so the board's line can be completed. Do not fake it.
6. Not rendered, so unverified: the not-following state of Officers sent (the `Follow` button at 64 px, its right alignment and the shortened `from` line), the hover ground on Next raids rows and the click that puts a row in Tonight, focus-visible on the four answers. Attach one capture of the not-following Follow state and one of a Next raids row hover before merge.

## Checklist (what the renders can show)

- Raid day: Guild tab dot, four answers blue, job in client gold, prep rows in the 132 px label column, Sign-ups with tally and roles: pass.
- Queued: Late teal, `not sent yet`, footer `Sends when the companion runs`, no dot: pass.
- Older companion: Sign-ups left out, `Sign up on the site`, answers hidden: pass (spec 3.7, 8).
- No companion: `No raid from the site yet.`, guild block from the pack, Officers sent kept: pass.
- No pack / unknown guild: `No raids from Thornwake Company on the site yet.`, top five and place left out: pass.
- Class colours, tier-coloured scores, spec icons, item quality colour on Drakefire Amulet: pass.
- Fixed column widths, no table crossing the vertical rule, no boxes or cards: pass.
- Section headings per 3.5: FAIL (must fix 1).
- Guild line on one row: FAIL (must fix 2).

## Round 2

Verdict: FIX (one must-fix). Judged on the redrawn renders at 740 px.

Closed: guild line on one row in raid-day (standing and nights only); `of 15 raiders` says what it counts; 4 px above the own row with the crest; `1 of 2 jobs` dropped; roles coloured against target; `local` clock; `Change it on the site`; boss on the job line; prep icons, bags row; `from Brannoc, Tue`. Late in gold stays as ruled (spec 3.5 web parity).

Must fix:
1. l5-old-companion.png: the guild line still wraps. `Ratings 8 days old` drops to line 2 as `8 days old` with the orphan, pushing the table 15 px below the left column. Round 1 acceptance named this render. Put the pack age in the footer only (spec 3.5 already shows `Ratings 8 days old` there on Me) or shorten to `Ratings 8 d old`; the line must be one row at 288 px.

Ruled, accepted: section headings stay in the title face. Spec line 111 says Marcellus only for the window title, but the Builds tab shipped the same after the 2026-11-13 brand review. Consistency across tabs wins; amend spec line 111 to list section headings so the spec stops contradicting the build. Not a blocker.

Still open, not blocking: attach the not-following Follow capture and a Next raids row hover capture before merge (round 1 should 6).

## Round 3

Verdict: SHIP. Judged on l5-old-companion.png and l5-follow-hover.png at 740 px.

Closed: round 2 must-fix. In l5-old-companion the guild line is one row, `3rd of all guilds · Ratings 8 days old`, and the table starts level with the left column. The Follow capture shows the 64 px `Follow` button right-aligned beside the pushed build with the shortened `from Brannoc, Tue` line, and the Guild tab dot is on. The row hover shows the tooltip `Molten Core / 22 of 40 coming`, matching the Sign-ups tally. Round 1 should 6 is met.

Polish, not blocking:
1. In the hover capture no hover ground is visible on the Sun 13 Dec row. Confirm the row takes the 8 percent white ground the spec gives rows; the tooltip alone does not show the row is clickable.
2. The tooltip sits about 55 px below the hovered row and covers the guild line and the Score column header. Anchor it to the row's lower edge, or to the cursor, so it does not cover the guild line.
3. Spec 3.5 line 111 now records section headings in the title face: confirmed, no further action.

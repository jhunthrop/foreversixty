# AI-slop review, L5 Guild tab

Verdict: FIX (copy and one empty state; no visual tells 1, 2, 5, 8, 12, 13 or 15 found)

First impression: reads as a player's tool. Real item names (Drakefire Amulet, Greater Fire Protection Potion), class-coloured roster, a rank jump 5 to 9 for "you", a flat blue button row, no glow, no gradient, no cards. Not generated-looking.

## Fixes

1. Tell 22, "for Onyxia" on the "Your prep" header (all raid-day renders). Tonight directly above is Onyxia's Lair. Drop the qualifier.
2. Tell 22, "Your" prefix on three labels: "Your sign-up:", "Your job:", "Your prep". The tab is the player's own; "Sign-up:", "Job:", "Prep" say the same. Keep "you are 9th of 15" and "you follow it" only if kept short.
3. Tell 22 / 21, footer "Sends when the companion runs" (l5-queued). It explains plumbing to a veteran. The line "Late as DPS, not sent yet" already says the state. Cut the footer, or make it "Queued".
4. Tell 22, redundant counts in Sign-ups: "12 of 40 coming" then "11 In · 1 Late ... 16 not signed up" then "1 of 2 jobs filled". Three restatements of one headcount. Keep the In/Late/Bench/Out line and the role line; drop "1 of 2 jobs filled" (it is also the one vague line, jobs are never defined on this tab, and it is not on the approved board).
5. Tell 17 / colour by habit: yellow is used for "Not signed up" (four times on raid day), the job text "Phase 2 ranged, west side", and the Guild dot. One colour carrying warning and plain data. Make the job text white (job is information, not a warning) and keep yellow for the sign-up nag. Mute the two repeat "Not signed up" rows in Next raids to grey, since only tonight needs the nag.
6. Copy accuracy (l5-no-companion): Tonight says "No raid from the site yet." but the cause is no companion, not no raid. It reads as a real empty state and it is false. The left half is also a dead column. Say "Companion not running" (or the true cause) and move the guild block full-width or let the right column fill the space.
7. Tell 14, light: section headings are the display face at regular weight and body is one weight, so the hierarchy comes from size only. The board used bold headings. Restore weight difference so "Onyxia's Lair" and the section titles read over the labels.
8. Layout nit: "attendance" wraps as an orphan in the guild subline (raid-day, queued, no-companion). Shorten ("91% att.") or let the line wrap at a middle dot.
9. Board drift: the build drops "Raid leader Brannoc", "invites 19:45", "server" after the time and the dates on Officers sent. These were the specific, idiosyncratic details. Restore invites time and raid leader if the data exists; "20:00 server" is the one qualifier that does change meaning, so keep it. Guild line on the board ("Molten Core 10/10 · Onyxia 1/1") was real content; the built "3rd of all guilds · 41 raid nights in 90 days" is fine too, but the progress line is what a raider checks.

## Authentic, keep
- Uneven hierarchy: raid name with its boss icon is the one big thing.
- Class-coloured names, quality-coloured item, real class crest icons in the roster.
- "you are 9th of 15" row pinned in the list with its true rank.
- Short player copy: "In as DPS", "Bench", "Late", "Att.". No em dashes, no pitch, no eyebrows, no icon grids, no pills, no glow, no card chrome.
- The tab dot has state behind it, and degraded states (old companion, no pack) say what is missing in the roster ("Ratings 8 days old").

PASS once 1 to 6 land; 7 to 9 are recommended.

## Round 2

Verdict: PASS

Checked l5-raid-day, l5-queued, l5-no-companion against round-1-fixes.md.

- 4 (redundant count): "1 of 2 jobs filled" is gone; roles now coloured against comp target. Fixed.
- 8 (orphan "attendance"): guild line is one line now. Fixed.
- 9 (board drift): "Flask 2 · Food 20", "from Brannoc, Tue" and the boss named on the job line restore specifics; the missing fields are absent in the data and not faked. Accepted.
- 1, 2, 3, 5, 6: owner-approved spec copy and colours (3.5, 3.7, 8, 11). Judged as the owner's words: short, player-voiced, no pitch, no em dashes, not generic-AI marketing. Not tells under the ruling. Gold carrying both the job and the sign-up nag is the client's own gold, not a decorative accent.
- 7 (heading weight): title face is the shipped Builds tab convention; reads as a deliberate face change, not a one-weight default. Accepted.
- No visual tells 1, 2, 5, 8, 12, 13 or 15 present; no glow, gradient, card chrome, pills or blur.

Non-blocking note: in l5-no-companion the left column is empty under "No raid from the site yet." That is spec copy, but it names the wrong cause; consider "Companion not running" in a later spec revision.

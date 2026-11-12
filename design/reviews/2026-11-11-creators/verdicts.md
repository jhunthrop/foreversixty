# Creator program landing: persona verdicts

Captures: `build/` (full page, 1440 and 390, the fixtures' clock 2026-12-09), made by `capture.mjs` from the built site with every API call answered from `web/tests/e2e/fixtures/creators`. Spec: `design/specs/2026-11-08-creator-program.md` (section 0G is this landing). Each persona reviewed the round-1 captures; the dispositions below say what changed for round 2 and what was held, and why.

## Round 1 (commit e2f8b6a73)

### ux-designer: SHIP WITH FIXES

MUST
1. Month names spelled out everywhere (`monthLabel`). Fixed: `Nov 2026`, `Dec 2026 · of $500 cap`, the Select `Dec 2026 (open)`, the dialog `Pay Grimbrew $168.60 for Nov 2026?`.
2. The public page's "On this page" rail is not in spec 3e. Held: the public page keeps main's redesigned layout (P14 of the redesign, owner-reviewed); the coordinator's brief for this landing says so. The spec's 3e is superseded by the shipped page.
3. The lead sits in the panel, not the PageHeader. Held, same reason.
4. Requests gone in the edit capture. A capture artifact (the edit state served no requests); the capture now serves them.

SHOULD
5. The payouts-off line sat on the selected Row. Fixed: 12 px under it.
6. Lines wrap on the year. Fixed: a month view's lines read `10 Nov`.
7. The balance line is not in 3b. It is in spec 0E (the balance check); copy.ts holds it.
8. No current marker on the creator being edited. Fixed: the Row is `selected` in edit mode.
9. The date picker's glyph is invisible. Fixed in the shared field recipe (`scheme-dark`); the typed format is the browser's.
10. Week rows all bold. Fixed: the week is the title, the counts are three Figures in the trail, `still filling` under a week less than 14 days old.
11. "Your figures" as a titled panel. Kept: the panels grammar gives every list its head; the meta reads `all time`.
12. Prose first at 390 on the public page. Held: main's page.
13. Request buttons above the text at 390. Held: Row's wrap order is the shared one.

### wow-player: SHIP WITH FIXES

MUST
1. Stats: the Payouts rail promised a first-week approval the capture had already missed. Fixed: the rail says `Nov 2026 closed on the 1st. We approve it by 7 Dec.` from the API's open month.
2. `$27 after 6 Jan 2027` unexplained; "Before Stripe fees." contradicts the share paragraph. Fixed: `$27 at list price after 6 Jan 2027`; `Before Stripe’s fees; you are paid on what is left after them.`
3. "use the site again" is vague. Fixed: `sign in again on a later day within 14 days of signing up`.
4. "Paid: first week" promises payment. Fixed: the lane reads `Approved · by the 7th`; the paragraph: `We approve it by the 7th, and the moment we do Stripe moves it to a payout account you set up with them.`
5. The cap's overflow unsaid. Fixed: `anything over the cap is not carried over`.
6. "ends your place" reads automatic. Fixed: `can end your place`; the earned-months sentence was already in the payouts paragraph.

SHOULD
7. The code sentence. Fixed as suggested.
8. "Each account: once" and the link-wins rule. Fixed: `once, for one creator`; `a player who came through your link and later types another creator’s code stays yours`.
9. `$2 a player` in the head. Fixed: `$2 a returning player`.
10. The lead. Held: main's page.
11. The 14-day note on Returning. Fixed: `62% of players · counts after 14 days`.
12. Week rows as figures with a filling mark. Fixed (see ux 10).
13. The approved state's first step. Fixed: `Your link and code are on your stats page.` with the link.
14. The reapply date as body text. Fixed.
15. Reply time promise. Held: no promise we cannot keep; the consent line now reads `We review every application and reply on your Forever Sixty account.`
16. The year in a month view's lines. Fixed.

### ai-slop-reviewer: PASS WITH FIXES

1. `all time` under every figure. Fixed: once, as the list's meta (`3 creators · counts all time`, `Your figures · all time`); the stats captions now carry each count's rate instead.
2. Qualifiers. Fixed: `Dec 2026 so far · of $500 cap`; `by week · latest weeks still filling · no names`; lines without the year.
3. The job lines under the headings. Held: `/admin/links` carries the same lead; house pattern.
4. Hints that restate the label (Name, Channel, Monthly cap). Fixed: removed; the link, code and Starts-on hints stay (each states a rule).
5. The terms lane repeats the figures. Held: main's page; the lanes are its approved structure.
6. The pitch lead. Held: main's page.
7. Apply three times. Held: main's page.
8. The apply form's empty right column. Held: `Page layout="form"` is the shared form layout.
9. Button weight. Fixed in part: Decline and Lines are ghost; Edit stays secondary like Copy on `/admin/links`; Approve is the one primary in a row.
10. Week rows. Fixed (see ux 10).
11. The consent line. Fixed (see wow 15).

### info-density-reviewer: PASS WITH FIXES

1. Weeks as prose. Fixed: three Figures per week in the trail.
2. No cap headroom or rates. Fixed: `Dec 2026 so far · of $500 cap`, `62% of players`, `28% of returning`.
3. The Payouts rail. Fixed: the Stripe status Tag in its head, the closed month and its approval day, `paid so far` under the months.
4. Week-over-week deltas. Held: the figures per week show the trend; a delta mark is a later pass.
5. Sort, channel link, cap bar, drill-down on the creators list. Fixed in part: the request Rows link the channel; the rest held (three creators; the statements page is one link away).
6. Conversion per click on the admin list. Held: `/admin/links` shows the funnel per link already.
7. The suggested code on a request. Held: the form prefills it on Approve, where the owner can change it.
8. A paid statement capture and a paid/ready/blocked tally. Held: the fixtures hold no paid month; the status line per Row carries it.
9. The Lines rail as columns with a filter. Held: four lines in the fixture; the amount column lines up.
10. The public page's rails. Held: main's page.
Needs data: clicks per creator on `/v1/creators/me` is a follow-up for the API.

### security-reviewer: no critical, no high

Fixed: approving a request refuses an applicant who is already a creator before anything is made (409 `already_creator`); strict JSON bodies (unknown fields and trailing data refused); an application's text takes no control characters; the application budget is five an hour per account as well as per address. Held with a note in spec 0G: the browser's click time on `/v1/me/arrival` is the client's claim, like the slug itself; the second-day and character rules, the cap and the owner's Approve bound what a false claim can earn.

## Round 2 (commit a7871d703, fresh captures)

- ux-designer: SHIP WITH FIXES. MUST: the running week's lone figure sat in the Paying column. Fixed: every week keeps three cells; a running week reads `–` in Returning and Paying. Polish taken: Cancel in the shared ConfirmDialog is ghost so Pay is the louder button; Copy onboarding invite and Refresh Stripe status are ghost. Held: the Lines button's edge on a Row with no Tag (the shared trail's gap); the Row wrap order at 390.
- wow-player: SHIP WITH FIXES. MUST: `We approve it by 7 Dec` read as a missed promise on the 9th. Fixed: once the 7th has passed the line reads `Nov 2026 closed on the 1st. We are approving it now.` Minors: `No payouts yet…` replaced by `$0.00 paid so far`; the `Needs info` Tag is Stripe's own state (spec 3a) and stays.
- ai-slop-reviewer: PASS WITH FIXES. Fixed: the empty payouts line is the paid-so-far figure; `all time` is gone from Your figures (kept once on the owner's list); the captions read `62% of players` and `28% of returning`; the players meta reads `no names, by design`.
- info-density-reviewer: PASS WITH FIXES. Fixed: `–` in the running week's empty cells; the paid-so-far line always shows. Held: a paid fixture and a payout history capture (the fixtures hold no paid month); clicks on `/v1/creators/me` stays an API follow-up.

## Round 3 (the fixes above, captures in `build/`)

Checked by the integrator against each MUST above on the final captures; every MUST is on screen, and the round-2 polish items are in. Verdicts stand at SHIP and PASS.

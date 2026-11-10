# Batch F, lane premium: review round 1 (2026-11-09)

Captures: build/ at the round 1 commit 9505e8dec. Verdicts word for word, then what the lane did.

- ux-designer: VERDICT: SHIP WITH FIXES
- wow-player: VERDICT: SHIP WITH FIXES
- ai-slop-reviewer: VERDICT: FAIL
- info-density-reviewer: VERDICT: DENSE WITH ADDITIONS

## Applied in round 2

- Signed-in head answer: the figure plus one line (date or source, Founder, See your plan) (ux M1).
- Coach sample: Then and Doing well rows are open, flush with the column; the title is "The coach after a raid" so Sample sits beside it (ux M2, S10).
- Manage billing on a guild column is full width at the panel foot (ux M3).
- Guild audiences shortened ("Up to 25 verified members", "Any number of members"), so the four columns line up (wow M1).
- Every plan button reads "Subscribe · {price}"; the column names the plan, and no label wraps (wow M2, ai-slop 4).
- Guild and Raid guild lines one short sentence each; the detail stays in the rules (wow M3, ai-slop 9).
- The refund question is gone from Questions; the Refunds rule (what refunds.md says) is the one answer (wow M4).
- The verified members line above the rules is gone; the Verified members rule says it once (wow M5).
- The founder line in each column is "$50 after 6 Jan 2027."; the notice says the deadline once (wow S1, ai-slop 4).
- "Free stays complete. Premium pays for the servers..." is cut (wow S5, ai-slop 5).
- /creators: the bounty is the head's answer ("$2 a player, then 50% of their Premium for 12 months") (ux S11); every section with numbers has its figures in the lane (Counting, Each month), and What you get adds "One player, first year" worked from the plan table, before Stripe fees (info-density 1, 2); the perks line starts "You also get" (wow S7); Apply is no longer in the contents rail (ai-slop 10).

## Not applied, with the reason

- ai-slop 1, 2 (panels everywhere; the creators card): the owner-approved round 5 system and the directions mocks b4-premium and b4-creators put the plans and the reading sections on panels; the four plans are the mock's plan cards. The creators panel's empty right half is now the figures lanes.
- ai-slop 3, ux S5, wow M7 (checkout confirmed is thin): the checkout is a hand-off the directions mock b4-checkout draws as one framed panel; Steps, the renewal line and the three next links are the spec's (premium spec 2026-11-08 section 6). Left for the account lane if the owner wants more there.
- wow M6, ux S7 (the owned checkout's head price): the head draws before /v1/me answers so the band does not shift; hiding it after would move the page. Left as is.
- wow M5 second half (the coach finding's Bloodthirst sentence): the coach's own fixture text (logs lane), not this page's copy.
- ux S8 (the Free column's space), S9 (rail length): the columns stretch so the buttons line up across the row, as the mock; the rail is as long as the sample.
- ux S10 first half (teal rail on Fix first): the shared Row's selected style, not a page change.
- info-density 5 to 8 (seat dates, member counts, per-guild raider cost, a link to your own review): need account or guild reads this static page does not make; recorded for the account and guild lanes.
- ai-slop 6, 7, 8 (Sample pill, eyebrow labels, bullet dots): the Sample tag is required by the coach sample spec (polish 3.6); the labels and the ruled list are shared primitives.

# h-guide-spec round 3 verdicts (captures: round3/captures, real-data build at f615da540)

## ai-slop-reviewer
VERDICT: PASS WITH FIXES

1. Retribution priority rows 6 and 7, both Seal of Command rank 5 (raid build only), not joined like the
   priest's runs.
2. "For the raid build." under the heading plus a "Raid build only" pill on four rows: keep one (tell 22/17).
3. "mana % is over boss health %" (logged, out of lane).
Round 2 items 1, 2, 3 and 6 fixed; 4 and 5 closed as allowed.

## ux-designer
VERDICT: SHIP

No blockers. Shoulds: the panel shorthand "mana % is over boss health %" reads like an unfilled placeholder
(humanise it); Retribution's three extras lists in two columns, not three; 390 captures with the folds open.

## Round 4 changes (9b1c7b199)
The priority joins a spell's run after the raid tags are set (Retribution's Seal of Command rows join); the
"For the raid build." line only when no row is tagged; the page writes the panel shorthand out ("your mana is
ahead of / behind the boss's health"); extras with no prose in auto-fit columns; 390 captures add an
all-folds-open shot.

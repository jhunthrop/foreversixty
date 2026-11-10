# Slop sweep 2026-11-09: reviewer verdicts

Captures at 1440 sit in `before/` (production, the head stubbed to the owner's view) and `after/` (this branch, fixture build), taken by `design/mocks/shoot_slop_sweep.mjs`. Pages: character, guide-fury, guide-holy-priest, guide-beast-mastery, guide-subtlety, guide-retribution, premium, changelog, privacy, terms, refunds.

## ai-slop-reviewer

- Round 1: PASS WITH FIXES.
  - Fixed in 1edf39af1:
    - refunds said one thing three times;
    - the score card and gear note explained the obvious;
    - Premium carried a sample caption pitch, "Free stays complete", a repeated cancel line, "Nothing switches off" twice and "Members never pay to sign up".
- Round 2: PASS WITH FIXES.
  - Fixed in 190b46901:
    - the Beast Mastery opener tricolon;
    - the terms pledge "even if it did not have to";
    - the privacy slogan "for anyone, ever";
    - "usually much sooner" and "nothing to opt out of".
- Round 3: **PASS**. "Prices in US dollars." and the Premium FAQ refund line are accepted.

## wow-player

- Round 1: SHIP WITH FIXES.
  - Fixed in 1edf39af1:
    - "members sign up with one tap" on Premium;
    - "Holy tier D" read as a raid set, so it is now "Holy: D on the tier list" on one row with the other facts;
    - "world buffs at the pull" (Forever has none);
    - "Guild and Raid guild are further down" is now "Guild plans".
- Round 2: SHIP WITH FIXES. The refunds rewrite had dropped the owner's manual cancel step; it is restored in one clause in 190b46901. The blank spell icons are a fixture-build artefact, not copy.
- Round 3: **SHIP**. The Beast Mastery opener matches the guide's rotation section.

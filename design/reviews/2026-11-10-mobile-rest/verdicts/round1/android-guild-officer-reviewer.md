# guild-officer-reviewer, Android guild and raid sub-screens, round 1 (2026-11-10)

**VERDICT guild-raids: NOT WORTH IT**

The member and officer flows would pass on their own. The contrast failure is a regression the restyle introduced onto the page's most-read facts, so it holds the gate.

## Restyle findings
- **BLOCKER:** secondary text on the tan ground is nearly unreadable, at about 2:1 by eye. This hits the invite, pull and end labels, the guild name, the comp line, the section labels and the calendar labels. Fix: a ground text token at 4.5:1 or better, cream for the time labels, and a contrast assertion in the tests.
- **SHOULD:** "Forever Score" breaks sentence case on the applicant screen (copy).
- **SHOULD:** the ground is mixed. Applicant, unanswered and live raid night are flat dark while the raid page is tan. Unify them.
- **SHOULD:** empty icon squares in the buff row, the loot item, the raid crest and the spec sheet. Is this a harness stub or real?
- **SHOULD:** the roster sort menu is clipped and see-through.
- **SHOULD:** the loading roster shows an empty stats card.
- **NICE:** two columns for the composition at extra-large type. A label above each figure on raid night.

## What works
- The member raid page.
- The officer comp against targets.
- The unanswered sheet, the best screen in the set.
- The applicant's checklist.

## Follow-ups (product, not restyle)
- Comp targets for members.
- Raid-ready flags on roster rows.
- Applicant actions pinned at the bottom.
- Context for the sim figure.
- A loot fairness view (needs data).
- Assignments and the boss list (needs data).

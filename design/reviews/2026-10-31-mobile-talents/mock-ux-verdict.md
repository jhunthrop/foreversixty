# F6 Talents mock: ux-designer verdict (2026-10-31), after the fix round

Reviewed: boards `f6-talents`, `f6-rank`, `f6-states`, `f6-xl` (`-ios`, `-and`) against `design/specs/2026-10-31-mobile-talents.md` (Revision 1) and `mock-player-verdict.md`. I viewed f6-talents-ios, f6-rank-ios, f6-rank-and, f6-xl-ios and f6-states-and after the last regeneration; f6-states-ios, f6-talents-and and f6-xl-and were regenerated from the same code and not re-viewed after the final edit.

## Blockers from the player

1. **Dashes hide in closed trees.** Fixed. Each tab whose tree holds a dashed cell carries an 8 pt wipe dot (F6.2 shows it on Elemental and Restoration; the vector now holds a Restoration difference, Tidal Focus 5 to 4). Spoken value `talents.a11yTabDiffers`. The dot is also drawn in the stacked tabs.
2. **Gold sends the player to the wrong talents.** Fixed by dropping gold on Your build. Cells are green (points), open (line border, full face) or locked (line border, 40% face). The guide's wants show as `−n` gold chips on the cells where you are under, and `+n` wipe chips where you are over. Partly filled talents are no longer a special case.

## Shoulds

| Should | Ruling |
|---|---|
| Dash direction | Adopted: signed chip plus ring colour; the sentence stays the planner's `diffCount`, and a `talents.missing` line states the guide's shortfall |
| Maxed green vs in-game gold | Rejected, with the reason in Revision 1 item 4: one system with the planner; every cell prints its rank |
| Requirement lines | Adopted (`Requires 30 points in Elemental Talents` shown on Elemental Mastery, F6.4), unmet only |
| Arrowheads | Adopted (drawn on every line); planner follow-up flagged |
| `pickSpec` wording | Reworded; the planner button is removed |
| iOS title bleed | Fixed: inline title, solid bar ground. No bleed in the final boards |
| Capstone under the tab bar | Visible on Matches (iPhone 17 and Pixel 9); half covered on Differs because two extra meta lines sit above. Accepted as the first scroll |
| Nits | Mixed-case names, `guide build` wording adopted; Matches with unspent stated; tag legibility left (tag is the shared `Tag`) |

## Data doubts

- **Band length.** Confirmed from the files: a band string holds `band - 9` points. The planner counts the whole string with no level cut, so a level 58 reads `Differs` on the site too. The app matches the planner and explains it with `talents.beyond`. A fairer rule needs the owner's ruling and a change to `diffAgainstBand` plus both apps; not decided here.
- **Elemental Mastery 0 in the guide.** The data's. Not ours to change.
- **Talent texts** are the build 1.60.1.70532 `ranks[].description`, copied verbatim into the mock's `data.json`.

## Open items (not mock defects)

- `home.next.talentsDiffer` and `home.answer.pickSpec` need the same wording changes on Home (flagged to the coordinator).
- The planner has no arrowheads; the apps do.
- No board for the beyond-the-band line or the stale-export state.

VERDICT: SHIP

# info-density-reviewer, round 1 (boards in this folder)

**Verdict: THIN WITH FIXES.** The tour, panel and chip structure is sound and uses its data on the pages that already commit to a build. The undecided state and the Builds rows are the thin spots, and the data to fix them ships in the addon: each `defaultBuilds[spec][kind].order` is a full ordered list of points, and the live ranks are read.

## Builds dialog

1. Own rows, line 2: the end-state split alone, the same at 24 and at 60. Read the character's position in the order instead: the next point's talent and the matched count; keep the split as a tail.
2. Own rows, line 1: only Following sits in the right gutter; put the next point or count there on the other rows, gold reserved for Following.
3. Above the first row, when undecided: show the split beside the question ("2 Holy · 2 Retribution").
4. Druid board: Feral and Feral Bear share first points, so a next-talent line alone may not tell them apart; the matched count or the first divergence would.
5. The companion explainer and the empty row above Custom build: an honest "No builds sent" would say more than copy.

## Question block

6. Add the split to the right of the question where "loads its leveling build" sits; that line is copy every tile click makes obvious.
7. Tiles: add the next point for that spec's build to the line ("Next: Divine Strength"), within the two-line cell.
8. A bare point count per tile.

## Panel

9. Pick rows: the points in that tree and the next talent beside the playstyle.
10. "Picks start at level 20" is the thinnest block; honest empty may be correct.
11. The rotation block at 13 while undecided: the generic low-level one, if any.

## Chip card

12. Delete the closing explainer line; fold the split into the heading. The yellow fills on all three rows suggest all selected.
13. The same figures as the panel, right-aligned only.

## Tour pages

14. Page 4's answer lines carry no counts; sign-up totals need data the addon may not have.
15. The "Next point at level 15: Deflection" pattern is right; add the tree for a hybrid build, and "of N" at the cap.
16. Page 1 is dense and finished.

**Ranked:** 7 and 9; 1 and 2; 6 and 8; 12 and 13; 14 and 4. **Needs data:** guild sign-up totals; per-build upgrade counts under 20; a "points match this order" figure (derivable, a new pure function).

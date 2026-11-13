# Builds window, wow-player review, round 2 (2026-11-13)

**Verdict: SHIP WITH FIXES**

Round 1 fixes confirmed on the re-rendered boards:
- The druid Up next now matches the data (Feral Charge at 25, Savage Fury 26 to 27, Predatory Strikes 28 to 30, Leader of the Pack 31, Blood Frenzy 32, Natural Instinct 33 to 34). "26 more points after these" is right (51 - 15 - 10).
- The hunter count checks out: 19 points remain (Now plus levels 43 to 60), six rows cover 18, "1 more point after these" is right.
- Off plan talents are named (Hawk Eye is Marksmanship, 3 ranks; Deflection is Survival, 5 ranks, both match the client data).
- The build name reads "level 60 build" everywhere, matching the row label. The raw `\194\183` is gone. "Levels 44 to 48" reads right.
- The wrong-code message now says what to copy and that it starts with FSB1.

## Remaining must-fix

1. The Off plan lines read badly: "Hawk Eye · Marksmanship · 3 off plan (Rank 3/3)". The player has 3 of 3 ranks, so "3 off plan" and "Rank 3/3" say the same thing twice. Use "Hawk Eye 3/3 · Marksmanship" and "Deflection 2/5 · Survival", and the red heading already says it is off plan.
2. Header "27 of 33 points · 5 points off plan · 1 to spend" still adds up only if the reader does the sum (27 + 5 + 1 = 33). Say "32 of 33 points spent" or keep the line and drop "of 33". Pick one reading.

## Should-fix, not blocking

- Spec "Level 60 build" against a player's word "leveling build" is now consistent, but "Level 60 and raid build" on a one-row spec is clumsy next to the two-row specs. "Level 60 / Raid build" would read cleaner.
- No respec cost line when Use this build replaces 15 points. A druid at 24 pays gold, and the red Replaces line should say so.
- Gear rows in the paste preview have no icons or sources (already on the follow-up list; fine for now).

---

# Round 3

**Verdict: SHIP WITH FIXES** (one small must-fix)

Round 2 items confirmed:
- Off plan lines read "Hawk Eye 3/3 · Marksmanship" and "Deflection 2/5 · Survival". Clean.
- The header "Level 42 · 27 on plan · 5 off plan · 1 to spend" adds to 33. Clear.
- The druid header "Level 24 · 15 on plan" is right for the data.

New in this round: the preview's "Against your talents" list (Improved Wrath 0/5 · Balance · build 5, and so on, with "and 17 more").

## Remaining must-fix

1. The row format is ambiguous. "Improved Wrath 0/5 · Balance · build 5" shows the player's rank as a rank fraction (0/5) and the build's rank as a bare number ("build 5"). A player reads "0/5" as the talent's own cap. Write "Improved Wrath · Balance · you 0, build 5", or "Improved Wrath 0 to 5 · Balance". Use the same shape the Off plan lines use.

## Should-fix, not blocking

- The list shows only what the build adds. A Feral druid switching to Moonkin also loses Ferocity 5/5 and the rest, which is the respec the player pays for. Add a "You lose" line, or say "and 17 more" covers both.
- Still no respec cost on the red Replaces line.
- "and 17 more" cannot be expanded (already on the follow-up list).

---

# Round 4

**Verdict: SHIP**

The round 3 must-fix is done. "Improved Wrath · you 0, build 5 · Balance" reads unambiguously, and the tree-order list puts the drops (Ferocity, you 5, build 0) behind "and 17 more". The earlier should-fix notes (respec cost on the Replaces line, an expandable "and N more") are follow-ups, not blockers.

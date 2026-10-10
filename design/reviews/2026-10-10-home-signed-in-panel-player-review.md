# Home signed-in panel without the Switch column: WoW player review (2026-10-10)

Spec: design/specs/2026-10-10-home-signed-in-panel.md. Renders: design/mocks/renders/home-panel*.png.
Canvas: https://claude.ai/artifact/3PPFhXNS84DYqnKCkZ5Akr

**Verdict: SHIP**, with small fixes in the same pass.

Findings, most serious first:
1. A two-line name ("Sir Obnoxious Yellington") grows the hero ~36 px and
   pushes the table on a header switch; reserve the height or cap the name.
2. Simulator card: "37.5 in band best in slot" has no unit; "SIM · 26 DPS"
   does not say it is yours. Write "26 DPS now · 37.5 DPS at band best in
   slot".
3. Phone: the dates timeline sits between the answer and the table, with
   Dec 9 orphaned and past dates (Sept 13, 17) at full weight; move it below
   the table on phone, mute passed dates at every width.
4. The hint "Change character in the header" reads as a footnote; on phone
   the header is only a crest; "Change character: top right"; plain text,
   never a button; revisit in two weeks.
5. Which character the page is on: clear (eyebrow, 34 px name, descriptor,
   header chip agree; realm kept).
6. The re-flow beats narrowing; the ~130 px gap between name and cards at
   1440 is acceptable; no fourth card.
7. Faction emblems in the descriptor are a blob and a generic shield; use
   the real crests.
8. At 2000 the footer logo and legal line pin to the far edges while the
   content is centred.
9. Phone: "FULL LIST FOR FURY 20 TO 29" wraps beside "YOUR UPGRADES"; the
   name wrapping beside the crest is fine; first screen holds name and the
   Best in slot card.
10. "you wear this" on every worn upgrade row is filler (pre-existing).
11. One character: "Your character" with no hint is right.
12. Nothing reads like an account settings page.

Data the player doubts (for the accuracy loop): a level-24 hunter at 105 DPS
against a level-27 warrior at 26 DPS (pet included?); leather Defias set over
mail for a warrior at 27 (+2.7 to +2.9 DPS); a melee weapon swap worth +0.6
DPS for a Beast Mastery hunter.

Questions for the designer: (1) how to stop the table jumping on a one- to
two-line name; (2) the Simulator card's two unlabelled numbers and what "in
band best in slot" compares against; (3) how a phone player finds "the
header" selector from the hint.

# Navigation bar with a character selector: WoW player review (2026-10-09, round 1)

Spec: design/specs/2026-10-09-nav-character-selector.md. Renders: design/mocks/renders/nav*.png.
Canvas: https://claude.ai/artifact/SDThco2iDDHhBgm5kostYG

**Verdict: NOT YET.** The 1440 bar and open list are the best character
switcher the player has seen on a Classic site; the structure is right.

Findings, most serious first:
1. Below 1280 the bar shows only a class crest (name and spec hidden below
   1440): two warriors or two mages show the same crest, so the bar answers
   "which class", not "which character".
2. "Link Battle.net" shown to an account that already has a Battle.net
   character, and the spec itself says Blizzard serves no Forever data, so
   the row promises something that does not work.
3. Signed out: "Pick your character" over an empty panel; the one gold
   button is Battle.net sign-in while the real way in is the addon or a
   paste, which is the quiet row below.
4. The fifth row is cut mid-line with no fade; five characters should not
   scroll; make room for six.
5. Phone sheet drops the fifth character with no cue; dark strip at the
   bottom.
6. The Discord glyph reads as a bell or headset below 1440; use the real
   mark with a tooltip.
7. Setup rows (Paste, Link Battle.net, Get the addon) plus gold MANAGE and
   "Your account" take 40% of the panel: an integrations menu. Keep Paste;
   show the other two only when missing; Manage quiet; gold is for the
   current row only.
8. "Log in to update" reads as log in to the site; say "Log in to the game
   to update"; ember orange on a druid's orange name.
9. The closed bar never shows that the current character is stale or failed.
10. Faction marks ~10 px, unreadable; Alliance not blue.
11. Missing renders: 768, 1023, 1280, 1439; phone Menu grid with "BiS";
    signed out at 1024 and phone; signed out with a pasted export; expired
    session; before hydration.
12. Order "current, then newest sync" reshuffles on every sync; order by
    last chosen, then sync.
13. Race is never shown (optional).

Answers to the designer's doubts: stale = more than 3× the player's own
median sync gap, clamped 3–14 days, always show the plain age; two-tap
sign-in accepted once finding 3 is fixed, and add "Sign in" to the top of
the phone Menu; Discord demotion fine above 1440, below only with the real
mark and a tooltip; tablets on the phone bar right in principle, unproven
without a 768/1023 render.

Door order proposal: Planner, BiS, Simulator, Logs, Rankings, Tier List,
Guides, Get set up (BiS is the gear-gap door; Tier List beside Rankings
confuses class rankings with log rankings).

Three questions for the designer: (1) below 1440, how does a player with
two same-class characters know which one the site is on; (2) why is
Battle.net the single gold sign-in button when it returns no Forever
characters, and why show Link Battle.net to a linked account; (3) what
does the bar show for a signed-out player with a pasted character or an
expired session, and where does a stale or failed current character show
without opening the list.

## Round 2 (2026-10-10)

Owner rule added between rounds: the closed selector shows the character's
full name as the game shows it ("Obnoxious Yell") at every width above the
phone bar, ellipsis only when it cannot fit; phone bar crest-only with the
full name as the Menu's first row. Door order now Planner, BiS, Simulator,
Logs, Rankings, Tier List, Guides, Get set up.

Round-one findings: 1–9 fixed; 10 mostly (marks bigger, still read as dots);
11 mostly (all captures exist, pre-hydration phone cells broken); 12 open
or unproven (order rule not visible on screen); 13 optional, open.

New problems:
- 1024 closed: the spec line drops to "Level 60", class and spec vanish;
  keep "Fury Warrior" (or "Fury · 60") and drop the level.
- 768 and 1023: chevron badge over the crest's lower right; name touches
  the crest with no gap; reads as a cramped phone bar.
- Signed out at 1024: the closed selector reads "Sign..."; use "Sign in"
  with a smaller icon, or icon only.
- Pre-hydration phone cells: desktop door row drawn over the phone bar,
  unreadable; must be redone.
- nav-open-many is identical to nav-open; the seven-or-more case (scroll or
  "Show all") is not covered.
- "Obnoxious Yell" has a space (player's point; the owner's rule that
  Forever names can be two words stands, so the example name is kept).
- Expired banner fine; the ring is the only closed-bar signal, acceptable.

**Verdict: NOT YET.** Deciders: (1) redo the pre-hydration phone render and
add a seven-or-more capture; (2) fix 768/1023 (no chevron over the crest,
a real gap) and 1024 (class and spec stay; signed out reads "Sign in").
Everything else fixed or minor; with those four done the player calls it
SHIP.

## Round 3 (2026-10-10)

All four deciders fixed: pre-hydration phone render real (before/after
stacked at 390); seven-or-more covered (nine characters, filter input, the
seventh row fading under the footer, scroll under a fade chosen over a
"Show all" row); 768/1023 clean crest, chevron clear, real gap, Menu
separate; 1024 line two "Fury · 60" and signed out reads "Sign in" in full.

Left, all minor: the ellipsis-before-Menu behaviour is untested because
"Obnoxious Yell" fits; 768/1023 show no second line until the popover
opens; the pre-hydration phone render has no signed-out state; the sort
rule is visible only through the example data.

**Verdict: SHIP.**

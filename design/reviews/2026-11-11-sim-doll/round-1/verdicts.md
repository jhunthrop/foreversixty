# Round 1 verdicts (captures at c2aa168e8, in captures/)

Lane sim-doll. Owner ruling 2026-11-11: the sim's Gear section shows the shared paperdoll frame as BiS, the character page and the reports do; the saved-runs strip keeps the list form. Reviewed: the Gear section after a run and on the landing at 1440, 2000 and 390, beside the BiS doll (bis-doll-1440.png) and the character page's frame (2026-11-09-p9-character/build-character-owner-60-premium-1440.png).

## ux-designer
The frame, slots, quality borders, selected ring and tooltip are the BiS and character page's; the title bar is the spec and the level; the 640 px cap holds; the phone shows the list form row for row as before.

Shoulds:
1. Centre the 640 px doll: 220 px of empty panel at 1440, 600 px at 2000.
2. Keep Upgrades on Best in slot in the section head after a run too (it moved to the foot).
3. A capture with the stats sheet present.
4. Out of scope: the phone Timeline's labels wrap one word per line (sim-gear-page-390.png).

VERDICT: SHIP

## wow-player
The doll answers "what am I wearing" and matches the other dolls on frame, crest, title bar, borders, colours and tooltip; the 390 list keeps the enchants on the green line.

Musts: the dead panel beside the left-aligned doll; the link's two positions; the capture's enchants (Presence of Might on a Fury warrior's head and legs, a Naxxramas shoulder enchant, a TBC glove enchant).
Shoulds: a two-handed cue on the off hand socket; Requires Level in red when unmet (the capture's level 29 character wears level 56 gear).

VERDICT: FIX

## ai-slop-reviewer
The doll reads as the game's; the surface around it does not: tell 13, the empty right half of the Gear card at 1440 and 2000. Minor: the Level 29 repeated from the page head; the "Your stats in the fight show after a rerun." line floats alone.

VERDICT: FAIL

## info-density-reviewer
The doll matches the other three; the 390 list is dense and complete. Blocking: the left-aligned doll with 240 to 590 px of dead panel beside it. Centre it, or put the post-run stats and set lines beside it at wide widths; before a run centring is the only fix.

VERDICT: THIN

## Response (round 2, ab003dadc)
- The frame is centred when nothing sits beside it (the landing, a running sim); from a 760 px panel the run's five figures, the measured miss line and the set lines take the column beside it, as the report's Gear at the pull pairs its frame and list; under 760 px they sit under the frame.
- Upgrades on Best in slot is the section's action in every state; the footer link is gone.
- The capture set wears Fury enchants (Strength +8 head and legs, Strength +10 gloves, no Naxxramas shoulder). The stats sheet still depends on the capture's browser run; the "after a rerun" line predates this lane.
- Routed, outside this lane: the off hand socket's two-handed cue (the frame's socket is named "two-handed" for assistive tech, the list says it; a visible cue is a PaperdollFrame change for every doll), the unmet Requires Level in red (the shared ItemTooltip), the phone Timeline's wrapping labels, the "Level" after the spec on the title bar (the character page's bar carries it the same way).

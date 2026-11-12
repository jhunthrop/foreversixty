# Lane sim-doll: final verdicts

Owner ruling 2026-11-11: the sim page's Gear section shows the shared paperdoll frame like BiS, the character page and the reports do, instead of the list form at every width; the saved-runs strip keeps the list form. Final captures in after/ (0959b3592): the Gear section on the landing, after a browser run and after a server run at 1440, 2000 and 390, beside the BiS doll.

Each reviewer's last verdict on this lane's surface:
- ux-designer, round 3 (9054a48c4): VERDICT: SHIP. Shoulds: the list outruns the frame by about 300 px at 1440; one figures grid at every width; the hovered slot's row highlighted; the figures first at 390.
- wow-player, round 5 (0959b3592): VERDICT: SHIP. Blockers: none.
- ai-slop-reviewer, round 3 (9054a48c4): VERDICT: PASS.
- info-density-reviewer, round 3 (9054a48c4): VERDICT: DENSE.

Rounds 4 and 5 changed only the capture set (dual wield, a level 60 spend), not the product. Round-by-round verdicts and responses: round-1/, round-2/, round-3/, round-4/ (their captures keep the section crops; the full pages are in after/ only).

## What shipped
- SimGearSection on the shared frame (form auto): the frame from a 760 px panel beside its column, under 760 px the frame centred with the column under it, under 423 px the frame's own list form. The title bar is the spec and the level; the page head names the character.
- The column: after a run the five figures and the measured miss line, then the set lines, then the worn list in every state. The worn list is one component (character/WornList) drawn by the report's Gear at the pull and the sim.
- Upgrades on Best in slot is the section's action in every state.

## Routed to the coordinator (outside this lane)
- Shared ItemTooltip: "+28 Crit" and "+20 Hit" without a rating unit (also in the h-doll final verdicts); the unmet Requires Level not red.
- Sim accuracy: the fixture's measured white miss (13.9% at 9.5% hit, dual wielding) in src/fixtures/sim/result-full.json, an engine figure.
- Data: Enchant Gloves - Greater Strength at +10 in the build's enchants.json beside the bracer's +7; whether Savage Gladiator Chain is a launch source.
- The stats block (sim results spec): the figures' order on a phone (after the list), the lone Haste cell, the "Your stats in the fight show after a rerun." wording.
- PaperdollFrame: a visible two-handed cue on the covered off hand socket (the socket is named for assistive tech; the list says it); the hovered slot's row highlighted in the worn list.
- The phone Timeline's labels wrapping one word per line (sim-gear-page-390.png).

# Server runs: Classic player review (2026-10-24)

Question the surface answers: "Did the sim I sent finish, and what did it say?" On /sim the list answers it first. Good: crest, spec, kind, result headline ("+41 DPS from Vis'kag" reads like a real upgrade), no fake progress bar, no queue talk.

Verdict: SHIP WITH FIXES

## FIX list
1. sim-empty-1440 / sim-payer-states-1440: the empty line says "Choose Run on our servers when you start a sim", but the block below shows one RUN button and "Runs on your machine; nothing is uploaded until you choose to save." A payer cannot find the choice and the line below contradicts the promise. Show the "Run on our servers" choice beside RUN for Premium, and swap the privacy line for Premium to say what is uploaded.
2. sim-payer-states-1440: two rows read "Frost Mage / Sim" with nothing to tell them apart (three of five are identical). Add the thing a player recognises: character name or gear set, and fight type (for example "Frost Mage · Sim · Raid boss, 5 min").
3. sim-payer-states-390: the unread gold dot and the "Running" dot are almost the same muted gold. Unread must be unmistakable: brighter, larger dot, plus an "New" word or bold headline. Running dot should be a different, calmer colour.
4. home-line-three-1440 / home-line-one-390: the line is about 11px text on a dark strip under the character card, and the dismiss "x" is a few pixels wide. Make it a real 44px target with a visible border, and put a label ("Dismiss") on focus/hover.
5. home-line-one-390: on the phone "Open result" is gone, so the line reads "Your Frost Mage sim finished: 1,244 DPS ->". Keep the words "Open result".
6. home-line-three-1440: "Your Frost Mage sim finished: 1,324 DPS" while the band shows a Holy Priest. Say which character when it differs from the band ("Your Frost Mage (Mirelle) sim..."). Also the 1,324 here and 1,204 on the sim page for the same run must match.
7. sim-failed-runagain-couldnotstart-390: "Could not start it again, Try again." is a comma splice with a stray capital. Say "Could not start the run. Try again." Also "Failed: stopped before completion" tells nothing; say "Failed: the server stopped the run" and say Run again reuses the same gear and settings.
8. sim-payer-states-1440: "Running" has an estimate on one row and none on another. Show "takes about N minutes" on every running row, or say "taking longer than usual" honestly when past it. "Starting" is fine only if it lasts seconds.
9. sim-payer-states-1440: "1,204 DPS" needs its context (fight length, single target) like the Details! header does, at least on hover. "Top Gear" and "Sim" as kind labels are OK.
10. account-runs-4-1440: "On · 4 runs this month" is a link with no hint where it goes, and in the same panel the plan says "Not on Premium." Fix the fixture so it cannot contradict, and make clear the link goes to the run list.
11. sim-error-390: "Your server runs did not load." is fine, but the Run block below does not say whether runs can still be sent. Add "You can still run on your machine."

Not an issue: no queue, no cancel, no percentage. That is honest and matches the promise.

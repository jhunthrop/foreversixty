# Round 3: the final gate

Round 3 closes the three items left open after round 2:

- **Pages sized to their content.** Each page is 240 to 450 tall (`Onboarding.frameHeight`) and anchored by its top.
- **A real tooltip on page 2.** It is drawn in `surfaces/TourExample.lua`. It shows the character's first panel upgrade, or failing that the Data BiS pick. The capture shows Impaling Harpoon, "Best in slot at level 20" and its Deadmines source.
- **A hover on Tooltips only.** Its title is gold, followed by the setting's hint and "Ends the tour. /fs help opens it again."

The mock now draws "coming" as the plain gold word that the Lua draws.

## Verdicts on round-3/

| Reviewer | Verdict | Blocking items |
|---|---|---|
| wow-player | SHIP | none |
| ai-slop-reviewer | PASS | none |
| info-density-reviewer | DENSE | none |

## Follow-ups for the next addon pass

None of these block shipping. They were left out so the shipped state stays exactly the one that was reviewed.

**Page 1**
- "tiers" becomes "tier lists".
- The no-guild loop gets a tidy bottom row.
- The "Welcome to Forever Sixty" subtitle is dropped on pages 2 to 5.

**Page 2**
- The ring row shows the real next point ("Level 25: Lethal Attacks") when a build is followed.
- The rotation row shows the panel's first button.
- The example gets a "Requires Level" line.

**Page 3**
- Show the character's own score when one exists.

**Page 4**
- Show the guild's real posted raid and job when one exists.
- Make "coming" muted rather than gold.
- Cut "on raid day" from the minimap note.
- Shorten the lead of the intro.

**Data check for the accuracy reviewer**
- Sniper Shot and Lacerating Strikes in `L.specPlaystyle`.
- Impaling Harpoon as the Beast Mastery band-20 main_hand.

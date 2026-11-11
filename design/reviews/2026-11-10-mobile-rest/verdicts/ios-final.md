# iOS secondary screens: final gate (2026-11-10)

Goldens recorded on CI (ios.yml record, run 42940733894).

| Group | ux-designer | wow-player | ai-slop | info-density | brand-consistency | guild-officer |
|---|---|---|---|---|---|---|
| guides | SHIP (r3) | SHIP (r2) | PASS (r3) | DENSE (r1) | CONSISTENT (r3) | |
| character | SHIP (r3) | SHIP (r2) | PASS (r3) | DENSE (r1) | CONSISTENT (r2) | |
| bis-search | SHIP (r3) | SHIP (r2) | PASS (r3) | DENSE (r1) | CONSISTENT (r2) | |
| account | SHIP (r3) | SHIP (r2) | PASS (r3) | DENSE (r1) | CONSISTENT (r2) | |
| guild-raids | SHIP (r3) | SHIP (r2) | PASS (r3) | DENSE (r1) | CONSISTENT (r3) | WORTH PAYING FOR (r2) |
| sims-reports-talents | SHIP (r3) | SHIP (r2) | PASS (r3) | DENSE (r1) | CONSISTENT (r3) | |

## What iOS needed beyond Android's port

**One N3 bar system on every pushed screen.**
- The bar is clear at rest with no title.
- The title fades in on scroll.
- A shared `BandHead` runs the band under the status bar with the page title.
- The snapshot hosts render pushed screens inside a NavigationStack, as the app does.

**The answer control.** One segmented pill with no answer filled when unanswered (the guild officer's blocker), shared with Home.

**Test hosts paint the repo's game icons.**

## Follow-ups (none block a verdict)

**From review:**
- The BiS set line wording "{dps} {unit} · Updated {date}" (shared copy, under another lane's copy lock).
- Guide loading and error heads should name the guide once it is known.
- The skeletons draw the real head.
- The note field becomes a bottom-rule input.
- The delete screens get a short line plus a web link.
- Forever Score bars coloured by role.

**Carried from Android's data and product list:**
- stat weights on the guide;
- the signed-out paperdoll;
- Fire Nova's icon.

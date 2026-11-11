# iOS secondary screens, round 1 verdicts (2026-11-10)

Goldens recorded on CI (ios.yml record, run 42901961850); sheets in `ios/`.

| Group | ux-designer | wow-player | ai-slop | info-density | brand-consistency | guild-officer |
|---|---|---|---|---|---|---|
| guides | NO SHIP | NO SHIP | FAIL | DENSE | DRIFT | |
| character | NO SHIP | SHIP | FAIL | DENSE | DRIFT | |
| bis-search | NO SHIP | SHIP | FAIL | DENSE | DRIFT | |
| account | NO SHIP | NO SHIP | FAIL | DENSE | DRIFT | |
| guild-raids | NO SHIP | NO SHIP | FAIL | DENSE | DRIFT | NOT WORTH IT |
| sims-reports-talents | NO SHIP | NO SHIP | FAIL | DENSE | DRIFT | |

All the reviewers agree that the components inside the screens already match Android:
- tier anatomy and crests;
- LoadError Plates;
- the sim Figure;
- kill and wipe dots;
- the delete Plate;
- signed out with one primary;
- the Plan Figure;
- the outline Reset and Cancel.

## The shared cause
The navigation bar has four treatments across pushed screens:
- band with no title;
- band with an inline title plus a large title below;
- transparent, with the band starting below the bar;
- black.

N3 rules that the bar has no title at rest and the screen's head continues the band and carries the title. Part of the mismatch is the snapshot hosts, which do not render the app's real bar.

## Round-2 work (iOS)
**One bar system:**
- the bar is transparent over the band and has no title at rest;
- a shared band head carries the title on every pushed screen;
- there are no black bars and no capsule back buttons;
- the snapshot hosts render inside a NavigationStack with BarAppearance.

**The answer control.** One FlowSegmented look that shows state, with nothing selected when unsigned. This was the guild officer's blocker.

**Test hosts** paint the repo's game icons: the buffs, loot and spec sheet.

**Layout:**
- the offline and new-data strips sit below the band;
- the guides root wraps at the largest text;
- the character offline state shows its strip text.

**Components:**
- text roles as Android (SectionTitle over the ground, Label12 inside panels);
- the picker chips on the control radius;
- secondary buttons as outline;
- the raid head's tags and guild name;
- mutedGround on the remaining ground lines.

**Copy:** the explainers Android cut (the sims footers, "Ranks are the ones you have at 60."), and the Strengths bullets.

**Captures:** sheet grabbers; scrolled delete captures; Not answered's chips scroll.

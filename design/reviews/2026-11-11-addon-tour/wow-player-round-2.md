# wow-player, round 2

**Verdict: SHIP WITH FIXES.** All four round-1 must-fix items are answered or handed off: "Healer, Holy Shock and Holy Light" (Holy Shock is Holy tier 5); "Two-handed weapons, Seal of Command and Vengeance" (Retribution tiers 3 and 5); "Cat form, bleeds and Savage Fury" (Feral tier 3); Twist of Light and Berserk gone; the druid crest on the pushed row; the 50-point raid defaults handed to the data lane (do not call the 60 boards final until the data has 51).

## Must fix

1. `panel-ask`: the split breaks mid-name ("2 Holy · 2" / "Retribution") beside a question that also wraps. Its own muted line under the question.

## Should fix

1. `panel-ask`: "Next: Spiritual" / "Focus" breaks a talent name across lines; put "Next: …" on its own line as the tiles do.
2. "Pick a path" stays "path" (recorded as the owner's word; "Pick a spec" would cost nothing).
3. The set-up lead promises the bag arrows while the panel says picks start at 20.
4. Protection's first point in Divine Intellect reads oddly to a tank, but the data is right (one point, then five in Redoubt).
5. `builds-dialog-60`: "Tab to switch" sits at the frame's bottom edge, dimmer than the other hints.
6. `builds-dialog-druid`: "Next talent: Feral Charge" at the top and "Next: Ferocity" on the Feral Bear row are two things with near-identical labels; the top line follows the build you have.

## Data checked

Every "Next:" on the boards matches the first entry of its order in Data.lua (Improved Wrath, Ferocity, Spiritual Focus, Divine Intellect, Benediction). Level 13 with four points is right.

## What reads right

The question and the split; the new leads ("Your next talent gets a gold ring and a better bag item gets a green arrow.", "What you pick on the site shows up on your bags, in your talent window and in your panel.", "Each boss gets a parse and a biggest loss once the companion uploads the fight."); "From the site"; Following as the only gold row; the chip card's plain client buttons; Feral and Feral Bear told apart; no plate or mail claim before 40; the shaman line carries no second weapon.

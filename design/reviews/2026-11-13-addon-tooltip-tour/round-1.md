# Addon tooltip and tour, round 1 (2026-11-13)

Renders: tooltip-waist-42-before, tooltip-waist-42-after, tooltip-waist-pick, tooltip-band-edges, onboarding-1-loop, onboarding-5-set-up-done.

| Reviewer | Verdict | Must fix |
|---|---|---|
| ux-designer | FIX | 1. a worn item the pick beats (case 16) showed no next band line; 2. one pick reads two spans (30 to 49 at level 35, 40 to 49 at 42); 3. the bare Shift "−0.3 DPS" under "Better than" is ambiguous |
| wow-player | FIX | 1. Shift figure unlabelled, sign reads backwards; 2. "Better than" gives nothing to do; 3. "Best at 60" vs "Best for 50 to 59"; 4. tour page 1 subtitle "Welcome to Forever Sixty" |
| ai-slop-reviewer | FIX | 1. Shift figure contradicts "Better than"; 2. tour page 1 subtitle restates the title; 3. "/fs tour opens this again." explains a command; 4. node sub-lines on tour page 1 |
| info-density-reviewer | FIX | 1. Shift figure unlabelled; 2. default and Shift differ on where the figure lives; 3. no figure or source for the At pick under Shift; 4. move "/fs tour" to the Done button's tooltip |

## Fixed

- Shift figure (all four): reads "Pick: −0.3 DPS" (Locale tooltipPickFigure), the figure said as the pick's.
- Case 16 (ux 1): "Next upgrade: <pick> +1.7 DPS" is followed by "At level 50: <pick>", the next band whose pick differs; the upgrade's source and the At pick's source and gain move under Shift. Swept over every hunter slot and band edge in band_tooltip_spec.
- Top band (wow-player 3): "Best at level 60".
- The tour's page 2 example still read the removed "Best in slot" key (nil for a character inside a band); it now says the tooltip's own gold line, "Best for 40 to 49".

## Held, with the reason

- Two spans (ux 2): the span runs from the player's band to the band that changes the pick. At 35 the belt is best for 30 to 49; at 42 the 30s are behind the player, so 40 to 49. One rule, read from where the player stands.
- "Better than" gives nothing to do (wow-player 2): keeping what you wear is the action; a figure on the default line is case 16's alone by spec 2.2 (figures under Shift for a hold).
- Default vs Shift figure (info-density 2): case 16 is an upgrade and carries its gain on the line by spec 2.2; a hold keeps its sub-floor figure under Shift. Two cases, one rule each.
- At pick's figure (info-density 3): the At line's source and gain are under Shift already (atDetail); the gain shows only at or above the 1.0 floor, which a 50 pick over a 42 belt reaches only at 50.
- Tour page 1 subtitle, node sub-lines, footer "/fs tour opens this again." (wow-player 4, ai-slop 2 to 4, info-density 4): the tour's frames passed all four reviewers in 2026-11-11-addon-tour; this lane changes only the footer's command (/fs help to /fs tour), which the owner asked to be findable. The page 1 subtitle was ruled in that tour's round 2 (Onboarding.head: on page 1 only, the same line on every page carried nothing), and page 1's node lines passed its round 3 as nouns.

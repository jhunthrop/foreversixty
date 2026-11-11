# ux-designer: BA1, rounds 1 to 3

## Round 1: FIX

Panel, onboarding, buttons (one primary per view), rarity and Android colours matched the spec. Raised: the item tooltip in Marcellus with old-style digits, a navy frame and neon green (#1 to #4); the title-band score in Marcellus (#5); kill green on gains (#6, polish); the Android Home caps (#7, lane F); no hover or selected captures.

## Round 2: FIX

Withdrawn: #1 to #4. The tooltip is the client's own GameTooltip (Friz, its navy ground and grey rim, GREEN_FONT_COLOR), kept in the client's style by spec G2; Marcellus is only the mock's stand-in for Friz, which ships inside the client. Fixed: the score is an Open Sans figure; `button-states.png` shows secondary rest, hover and held, primary rest and hover, close rest and hover. New: the title band's meta line cut the guild name ("Employ…").

## Round 3: SHIP

- The meta line shows "Level 60 Hunter · Employed Millennials" in full; the number and its unit sit on the title's line (`Widgets.titleBand`, `70e853db9`).
- The title stays the primary element in the band; the boxed close does not collide with the score.
- `onboarding-5-set-up.png`: Done is the one teal primary.
- Hairline plates, Marcellus headings, Open Sans body, rarity text colours, button recipes and the generated Android colours hold.

Polish, not blocking: close hover is barely brighter than rest; kill green on gains is the site's convention; the Android Home caps belong to lane F; the client tooltip needs an in-game capture for BA5.

VERDICT: SHIP

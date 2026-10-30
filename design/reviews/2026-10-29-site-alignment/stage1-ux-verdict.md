# Stage 1 UX verdict (site alignment + home reorder)

Reviewed on captures: stage1-* vs before-*, home-order-*, against design/specs/2026-10-29-site-alignment.md and the polish spec home order. Full-page captures at 1440 and 390 were read; 2000 captures not opened.

## Verdict: SHIP

## Blockers (must fix before push)

None. /account at 1440 and 390 is unchanged (the report's pixel-diff claim is consistent with what renders). Cards on /bis, the guide rail, /setup, /ratings, /sim tabs and the guild tabs take the /account surface (gold top edge, tile head, 18 px head type) with no layout break, and the home order matches the ruling (dates line, then upgrades, then look-up, no empty band above Latest).

## Shoulds (can land in stage 2)

1. /bis at 1440, Stat weights: the Scale bar column collapses to a 2 px dot (it renders as a full bar at 390). The same collapse is in before-*, so it is not a regression, but the bar is the table's one visual. Give the bar cell a min width at three-up.
2. /bis at 1440: This set is stretched to the height of Stat weights and is about 60 percent empty. Align to start, or move a second fact in. Same for Play it (about 170 px dead).
3. /bis: the spec-tab icons (Arms, Fury, Protection) and the four Play it spell icons render as broken-image glyphs in the captures. before-* shows the same, so this is likely fixture data, but confirm on a real build; tenet 2 and the no-missing-icon-flash rule make it a blocker if it reproduces live.
4. The navy band still sits on /bis, /guide, /guild and home. This is the stage 2 work under ruling D1. On /guild the TabStrip fade and the red faction art read faintly, as the report says. Stage 2 removes both.
5. Home signed in: two filled gold buttons (See your upgrades, Look up) still compete. Polish spec rule is one primary per region; make Look up secondary.
6. The /bis phone head wraps "Alliance" onto its own line after the title ("20 to 29 ·" dangles). Keep the sub-title on one line or drop the trailing dot.
7. /guild head: "Guild settings" is a bare gold text link above the invite row. Make it a proper action in stage 2.

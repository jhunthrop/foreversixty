# BiS landing redesign: WoW player review (2026-10-10, round 1)

Spec: design/specs/2026-10-10-bis-landing.md. Renders: design/mocks/renders/bis-landing*.png.
Canvas: https://claude.ai/artifact/WgxPbByTdqjkJUsD5zis4K

**Verdict: NOT YET.** Far better than today; the signed-in half answers the
question and the 28 DPS strings are gone. The fixes are cheap.

Findings, most serious first:
1. Signed out, the strip defaults to 60 ("Spec links open at level 60" in 12 px
   muted text): a level-27 warrior who clicks Arms lands on the 60 list. Default
   to 20 to 29.
2. "Top sources: dungeons, crafting." under three bands is filler; caption only
   where it differs, or cut.
3. The strip looks like a filter but only rewrites links; nothing visible
   changes on click; two gold highlights at once (YOU pill and selected pill);
   say "Spec links open at this band" in every state.
4. Signed out, the strip is the first thing on screen and the grid starts ~450
   px down; grid first or a compact one-row strip.
5. States board contradicts the spec on hover (spec: nothing moves; board:
   lift); missing active, selected-pill hover, and a character under 20
   ("Your best in slot at 20 to 29" would be false for a level 12).
6. Stale copy "Log in to the game to update" is wrong for addons: SavedVariables
   write on logout or /reload, then the site needs the sync step.
7. "6 of 15" / "14 of 17": a player counts 17 slots and asks where two went
   (empty trinkets below 41); "picks" is not a player word; "slots" is.
8. The one action is a small outlined all-caps button far right; name the spec:
   "Fury gear for 40 to 49 →", bigger; the identity block should link too.
9. Faction mark is a tiny grey shield; stray gap; warrior-tan name on a
   gold-tinted panel does not stand out.
10. Phone: only signed-in rendered; pills ~40 px not 44; 11 px mono tiny; YOU
    badge crowds; cards run ~2.5 screens not ~2.0.
11. Spec icons are first-talent icons, not the spec tab icons players know;
    Mage and Warlock cards end in a blank strip beside Druid's fourth row;
    tier-list pointer tucked at 13 px.
12. Spec contradiction: 60 called "raid-ready" while the leveling rule says no
    raid drops as default picks at launch.

Fine: tier-list material gone; game class order; "Yours" row mark; signed-out
choice faster than today; 2000 centred; no pitch.

Questions for the architect: (1) why 60 signed out, and what does the 60
list contain; (2) does the band change the lead's button and count or only
the links, and if only the links why does it sit above the lead; (3) what is
the 15 / 17, which slots are excluded, and what does a stale player do.

## Round 2

Round-one findings: 1, 5, 6, 7, 8 fixed; 3 mostly (label explains the strip);
2, 4, 9, 10, 11, 12 partly: caption gaps under 30 to 49 unexplained; blank
strip beside Druid's fourth row on Mage and Warlock; faction glyph still
tiny; phone pills wrap "YOU" at 360/390; the 60 note never names Molten Core
and Onyxia.

New: the "Example character … invented for this board" footnote is mock
scaffolding and must not ship; the strip's explanation is 11 px muted text
wrapping to three lines at 1440; at 360 the pills squeeze to ~11 px.

**Verdict: SHIP WITH FIXES.** Deciders: (1) remove the example footnote and
confirm it never renders live; (2) phone pills with no wrapped "YOU" and a
minimum touch size. Follow-ons: explain or remove the caption gaps; the blank
strip beside Mage and Warlock; the strip label at a readable size.

## Round 3 (architect's fixes)

Example footnote out of every render; phone pills 44 px, one-line label, gold
dot for the character's band below 480; "Same sources as …" captions so no
pill is bare; cards size to rows, grid aligned by top; strip explanation 13 px,
two lines. Not done: naming Molten Core and Onyxia in the 60 note, because no
band-60 pick or runner-up in the data comes from either raid (an owner-side
question: what "raid-ready, Phase 1" contains).

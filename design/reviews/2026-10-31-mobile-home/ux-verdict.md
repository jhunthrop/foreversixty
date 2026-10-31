# UX verdict: mobile Home (F2), iOS and Android

Reviewed: the 33 captures in this folder against spec section 5.2 (F2.1 to F2.7), the f2-* boards and the tooltip section. Known gaps in the lane report (selector, next raid, C6 fields, iOS tab bar) are not counted.

## Verdict: HOLD

One blocker, a one-line layout fix. Everything else is a should.

## Blockers

1. **Android, largest text: the tab bar breaks words mid-word** (`home-default-xl-and.png`, bottom bar). `Character` renders as `Chara / cter` and `Account` as `Accou / nt`, and the bar grows to about 100 dp. Spec says nothing is ever cut mid-word and that the largest size stays legible. Fix: single-line labels with ellipsis-free shortening (`Home`, `You`, as on the boards), or cap the nav label scale at 1.3 and let the icon-only fallback carry the rest. If the shell belongs to another lane, that lane must land the fix before Home ships, because the F2 capture exposes it.

## Shoulds

Parity
- **Next rows use different icons per platform.** iOS: sparkles (Talents) and flask (Sims). Android: star and play triangle. The boards use the spec icon (crest or spec icon, 36) for Talents and a flask for Sims. Use one icon set on both platforms, from the boards.
- **Tab bar labels and order differ from the boards.** Android shows Character, BiS, Raids, Sims, Account (five tabs). The boards show Home, Raids, BiS, You (four). Align with the spec's tab list. iOS could not be checked because its captures omit the bar.
- **Top of screen differs.** iOS draws a 168 pt empty bar with a centred `Velaren` before the hero. Android draws no bar. The selector (F1) will fill both, but until then iOS spends the first 19% of the screen on dead space. Left-align the name at the bar height of the selector chip (about 56 pt) in the interim.

Hierarchy and density
- **iOS default capture stops at the card footer.** The Already line and the Next rows are not shown, so they cannot be verified on iOS at the default size. Attach a full-length capture. Android shows them and they match the boards.
- **Band label reads `60 · raid`.** The boards read `60 · Phase 1 raid`. The lane report documents it. Ask the home endpoint for the label (C1) so it matches before launch.
- **Mono figures use the system mono.** JetBrains Mono is not bundled. On Android the gain column (`+11.1`) is a wide fallback with loose tracking, visibly different from the boards. Bundle the font on both platforms.
- **Quality-coloured names in the gear card truncate to two or three lines with the arrow pick below** (Android `gear-default-and.png`, Neck, Wrist rows). Rows grow to 110 dp. Acceptable, but the pick icon column could take 8 dp less to give the name a line back.

States
- **Gear row for an unknown worn item** (Android Wrist) draws a solid black 36 tile with no border, which reads as a broken icon. Use the dashed `--color-line` box (as for Empty) or a neutral question-mark tile.
- **Loading** reads `Loading your upgrades...` with three dots. The spec and boards read the ellipsis character. Use `…`.
- **Android loading** has no Next skeleton. Fine, the Next rows are not specified as skeletons. iOS matches.

Tooltip sheet
- **No `Equip:` / `Use:` prefix on the green effect line** (known data gap, lane report concern 4). A Use: trinket will read like a passive. Ship the trigger in the data lane; until then do not show the line for items whose trigger is unknown on a trinket or consumable.
- **Panel order, quality colours, Item Level in `#e5cc80`, set bonus layout, reputation source (`The Defilers (Exalted)`), armor type on the right, no armor type for Back**: all match the spec on both platforms.
- **iOS large-xl tooltip:** the header icon is centred against a three-line name. Top-align it to the first line.
- **Medium detent on iOS** opens at about 55% of the screen and cuts the set bonuses mid-line, which is expected for a medium detent; the footer Where it comes from is only reachable by dragging up. Acceptable.

Brand
- Tokens, band gradient, gold primary button, kill and gold tags, quality borders with inner line and glow, condensed display figures: match the boards on both platforms.
- Crests: none on Home yet (selector pending). When F1 lands, the chip must carry the circular crest, never a letter square.

## Checklist (F2)

- [x] F2.1 hero, two figures, note, sync line, card of five, footer, Already line (Android verified; iOS tail not captured)
- [x] F2.2 loading skeletons at real heights
- [x] F2.3 offline strip with stamp
- [x] F2.4 no gear export copy and two buttons
- [x] under 20, no spec, no list, all match captures present
- [x] F2.5 error in wipe colour with Try again, Next rows still drawn
- [x] F2.6 tooltip at medium and large, rep and weapon variants
- [x] F2.7 gear: 17 rows in frame order, BiS tag, empty box
- [ ] Large text: Android nav bar breaks (blocker 1)
- [ ] Parity: Next icons, tab labels (shoulds)

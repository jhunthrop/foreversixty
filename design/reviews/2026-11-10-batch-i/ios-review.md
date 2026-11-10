# Batch I: iOS after M0, the first look

The iOS goldens were recorded on CI (ios.yml record run 42797843107, Xcode 27.1, iOS 26.5), because the laptop's Xcode 27.2 still needs its first-launch install. The captures in `ios/` come from that run. iOS keeps its old tab layouts until M1i to M5i; this review asks one question: did M0's design system (round-5 tokens, Open Sans and Marcellus, the DesignSystem primitives) break iOS?

Reviewed: the design-system sheet, Home (60, raid day, 34), Raid (tab, tab at XL, member page), Gear (default, XL) and Activity (default, ax5), next to `design/mocks/directions/png/m-ios-*.png`.

## Verdicts

- ux-designer: **SHIP**. M0 broke nothing. Contrast holds everywhere: cream and gold on brown-black, the quality colours, the muted meta text and the DPS green. Crests, icons and both faces render with no missing glyphs.
- ai-slop-reviewer: **PASS**. The tokens and primitives read as one system with no gradients, glows, glass or blur. Quality-bordered items, crests, figures with units and the hairline lists on Home all read as authentic.

## M0 notes (non-blocking)

1. Button roles differ only by colour. The sheet's teal, blue and outlined buttons are told apart by hue alone, and Home's In, Late, Bench and Out are four identical blue buttons, so the answered state does not show there. The selected answer needs a mark as well as a colour (the selected ground plus the accent bar, as on the web).
2. Marcellus is a thin face on cream over near-black. Keep it to titles and section heads, as `Primitives.swift` already says.
3. On the design-system sheet, everything below the tab bar is unfilled canvas: the frame is 2700 pt and the content is shorter. Nothing is cut off. The sheet could fit its height to its content.

## For the iOS tab lanes (M1i to M5i)

- **All tabs:** the old teal nav band with a centred small title (Gear, Raid) gives way to the round-5 header. Gear also repeats its title as "YOUR GEAR", and Raid shows a double hairline under the band.
- **Tab bar:** a floating pill sits inside a black box. The restyle replaces it with the mock's bar.
- **Raid:**
  - The In/Late/Bench/Out segment is dim grey with a dark-green In, which is off the round-5 palette.
  - The member page has a large gap above "Note for the officers".
  - "Next raid" is a card inside a bordered card.
  - At XL, "OLYMPUS XXVII" and the invite line wrap.
  - The spec icon in "Your sign-up" draws its empty placeholder in the snapshot. That was already the case before M0, because the fixture has no icon for the test bundle.
- **Gear:**
  - "Wrist · Item 999999" leaks from a fixture.
  - The dashed empty slots are old states.
  - The grey intro line and the "Gold dot: ..." footnote explain the obvious.
  - The two-column grid truncates "Main Hand · +5.9 D…" and should let the line wrap.
- **Home:** "you have not answered, as Thandrel, Damage" and "Best run: ..." wrap awkwardly. Rewrite them to the round-5 Home copy.
- **Activity at ax5:**
  - The icon sits bottom-aligned to a three-line title.
  - The filter chips clip at the right edge.
  - "Mark all read" wraps.
  - The chips are the old blue pills.
- **Eyebrows:** "NEW FOR YOU" and "YOUR SIGN-UP" are uppercase letter-spaced labels. Keep them rare.

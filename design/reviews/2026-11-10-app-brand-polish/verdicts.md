# App brand polish: review verdicts (2026-11-10)

Lane `lane/app-brand-polish`. Evidence in this folder:

- `icon-preview.png`: the icon before and after, at the iOS home-screen sizes and Android 48 dp
  at every density, beside stand-in neighbours. `render-icon-preview.mjs` draws it.
- `launcher-home-android.png`: the real Pixel launcher. The light ring is the launcher's
  suggested-app ring.
- `side-by-side-*.png`: the website at 390 px, iOS before, iOS after, and Android after.
- `fade-*-{ios,android}.png`: simulator and emulator captures, light and dark, at rest and
  scrolled.
- `web-*-390.png`: the live site.

The scripts are in `design/mocks/2026-11-10-app-brand-polish/`.

| Reviewer | Verdict |
| --- | --- |
| ux-designer | **SHIP** (no must-fix items) |
| ai-slop-reviewer | **PASS** (no must-fix items) |
| brand-consistency-reviewer | **CONSISTENT** (no must-fix items) |

## ux-designer: SHIP

- **Icon.** The painted ground and brass seal read well, and the icon stands apart from flat
  neighbours. The seal holds 78% of the visible tile under the circle mask and under the iOS
  squircle, with an even margin. At 60 pt @2x and at Android mdpi the LX is soft but still
  recognisable. On a dark wallpaper the brass ring carries the silhouette.
- **Head fade.** The band dissolves into the ground with no hairline on Home, Raid, Gear and
  Activity, matching the site's 390 px treatment and its eased ramp. Text over the fade stays
  legible. The iOS 26 glass tab bar is unchanged.
- **Note, acted on.** The Android status bar was a solid black strip above the band. Fixed after
  the review: `FlowBandPage` now draws the band under the status bar (edge to edge). Once the page
  scrolls, a strip of the band holds the status bar. See `fade-raid-dark-android.png` and
  `fade-raid-scrolled-dark-android.png`. The ux-designer re-checked the new captures and **SHIP
  holds**: at rest the band runs under the status bar and the clock stays readable; scrolled, the
  band strip holds the status bar.
- **Optional.**
  - A touch more seal contrast for sizes under 60 px.
  - The Android emulator's large font scale wraps the Home hero. This is not in this lane's diff.

## ai-slop-reviewer: PASS

- The fade is not a decorative gradient: it removes the band's hard edge, as the site does.
- The icon replaces the gold-on-black placeholder, which was itself a tell, with the owner's
  painted art.
- No tells are introduced.
- Out of scope: Android's large emulator font scale.

## brand-consistency-reviewer: CONSISTENT

- **Icon.** The painted ground plus seal v4 matches the site's seal and the companion's macOS
  tile. The flat seal stays on Dark, Tinted and monochrome.
- **Head fade.** `size.bandFade` (64) and `gradient.bandFade` are generated into both apps.
  `tokens.test.mjs` holds them to `band.css`, and the Android stops were checked by hand. The
  foot rule is dropped, as on the web.
- **NICE items, not gating.**
  - The emulator's font scale.
  - A long KDoc line in `FlowBandPage`, since rewrapped.

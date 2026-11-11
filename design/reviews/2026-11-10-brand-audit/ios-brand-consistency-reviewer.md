# brand-consistency-reviewer: iOS app (2026-11-10)

Reviewer read 8 captures against the web references (m1i, m3i, restyle-app, unify-format/after) and grepped apps/ios (App, Sources, Widgets; tests and DesignTokens.swift excluded). It did not open every m2i/m4i/m5i image or the widget captures.

## DRIFT

**BLOCKER:** none. The home and gear side-by-sides (m1i, m3i) read as the same product: display type, dark plates, parchment ground, ringed crest, quality-bordered item icons. The floating tab bar in `restyle-app/gear-ios.png` is native idiom (ALLOWED).

**SHOULD**
1. **Forced caps on titles are still live.** There are 38 `uppercased()` calls, against the retirement in `design/specs/2026-11-07-mobile-redesign.md:165`.
   - Worst, all on the raid page: `RaidPageView.swift:228` (`raid.title.uppercased()`), `:173`, `:216`, `:309`; `RosterViews.swift:71`; `SignupViews.swift:252`; `RaidComponents.swift:47`, `:198`.
   - `RaidClock.swift:42` uppercases the weekday and month ("WED", "JAN").
   - The web (`unify-format/after/guild-raids-390.jpg`) is sentence case.
   - Fix: drop `uppercased()` and use the `Copy.*` strings as written.
2. **"Equipped" vs "Worn".** `apps/shared/copy/en.json:703` sets `gear.worn: "Equipped"`, while `gear.sheet.overWorn` says "over your worn". Capture: `m3i/side-by-side-gear.png`. Fix: pick one word, the one the web paperdoll uses.
3. **Glow and shadow the web does not use.**
   - Quality glow: `ItemIcon.swift:113` and `RaidNightView.swift:222` (`.shadow(color: quality.opacity(0.22), radius: 5)`).
   - Crest glow: `BrandImages.swift:43`.
   - `PanelShadow`: `HomeFlowView.swift:115` and `NextRaidBlock.swift:29`.
   - 7 non-token `.shadow(` calls in total. The web quality border has no glow (`bis-390.jpg`).
4. **SF Symbols where the web has generated icons.**
   - 148 `systemName`/`systemImage` uses, and 0 uses of the generated Vectors set (`gen-ui-icons.mjs` already writes `apps/shared/brand/ios/Brand.xcassets/Vectors`).
   - Top: `chevron.right` ×17, `chevron.down` ×5, `arrow.up.right.square` ×5, `magnifyingglass` (header search), `person` (avatar placeholder, `restyle-app/gear-ios.png`), `flask`, `lock`.
   - Share and calendar menus are native idiom (ALLOWED).
5. **Type and spacing literals.**
   - 100 `size: N` font literals and 97 `.font(.system` / `.custom(` uses outside the tokens.
   - 783 padding, frame and spacing literals.
   - Examples: `HomeFlowView.swift:183`, `NextRaidBlock.swift:215`, `RaidFlowParts.swift:36` and `:77`, `ApplicantView.swift:87`, `ItemPageView.swift:81`, `:321`, `:429`.
   - Fix: use `FSType` / `FSSpace` roles. `HomeFlowView.swift:292` already shows the pattern.
6. **Avatar placeholder.** The avatar placeholder is a person glyph, not the LX crest or the class crest (`restyle-app/gear-ios.png`).

**NICE**
- Token-formatted literals: `ReportView.swift:219-222` (`Color(tokenHex: 0x3D94F0 / 0xB866F5 / 0xF08CC0)`).
- Gold hardcoded as a `CGColor`: `EventKitCalendarAdapter.swift:36`.
- `widget.noGain: "—"` at `en.json:1183` uses an em dash as the no-value glyph.

## Shared-source consumption

| Source | Status | Evidence |
|---|---|---|
| Tokens | Generated: `DesignTokens.swift`, via `FSColor` / `Theme.Palette` / `GameColors` | 5 raw colours, 97 raw fonts, 100 size literals, 783 spacing literals |
| Icons | Tab bar uses the brand glyphs; the rest are SF Symbols | 148 `systemName`, 0 generated Vectors in use |
| Copy | Generated: `App/Generated/Copy.swift` from `gen-mobile-copy.mjs` | 1835 `Copy.` references, no literal sentences found |
| Formats | Not consumed: 0 references to `apps/shared/formats.json` | 47 `String(format:` / `DateFormatter` / `.formatted` calls, e.g. `RaidClock.swift:42` |

## Components duplicated with the web

- **Item line:** quality border and colour match the web. iOS adds an "Equipped" state word and keeps DPS in green.
- **Crest:** ringed, except the avatar placeholder (see SHOULD 6).
- **Raid row:** a big countdown here, against date block, chip and fill bar on the web. The order is acceptable; the caps are drift (SHOULD 1).
- **Sign-up:** four tinted buttons on iOS against a segmented control on Android, two different controls for one job.
- **Figures:** match the web, with tabular digits.

## Stale captures

- `design/reviews/2026-10-31-mobile-wave3-widgets/build/ios-*.png` (widgets, Live Activity) predate the 11-07 redesign.
- `m5i/ios-live-raid-night.png` predates the restyle.
- `m1i` side-by-sides predate `restyle-app`.

Re-shoot all of these on the iPhone 17 Pro simulator.

## Notifications

`notif.*` in `en.json` holds only the settings labels, which are fine. Push titles and bodies are composed by the server, so they are not judged here.

VERDICT: DRIFT

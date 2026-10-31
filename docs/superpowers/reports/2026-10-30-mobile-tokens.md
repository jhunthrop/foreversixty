# Mobile lane T: design tokens and brand assets

- Generator: `web/scripts/gen-mobile-tokens.mjs` (reads `web/src/styles/tokens.css`), outputs `apps/ios/Shared/DesignTokens.swift` (typechecked with swiftc) and `apps/android/.../designsystem/Tokens.kt`.
- Guard: `web/src/styles/mobile-tokens.test.ts` regenerates and compares both files.
- Fonts: `apps/shared/fonts/` six TTFs (Barlow 400/600/700, Barlow Condensed 700/800, Cinzel 700) from the site's Latin subsets, plus OFL licences.
- Brand: `web/scripts/build-app-icon.mjs` writes `apps/shared/brand/` (iOS AppIcon set, Android adaptive layers).
- Crests: `apps/shared/README.md`.

Notes: the site is dark only, so no dark-mode variants. Spacing is Tailwind's 4px scale (the site has no spacing tokens). No Android monochrome (themed icon) layer yet. JetBrains Mono is not bundled (apps use the system monospace).

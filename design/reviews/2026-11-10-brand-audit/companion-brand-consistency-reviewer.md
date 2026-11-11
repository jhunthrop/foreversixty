# brand-consistency-reviewer: companion and installers (2026-11-10)

Sources read:
- `companion/ui` (styles.css, tokens.css, icons.js, copy.js, format.js) and `internal/app/copy.go`.
- The tray code.
- The captures in `2026-11-09-companion/round-4`, `2026-11-10-unify-server/after/companion-*`, `2026-11-10-seal-flat/tray-*` and `2026-11-10-companion-installers/*`.

Not read in detail: the Mac move window, the uninstall pages, the settings and device captures.

## DRIFT

**BLOCKER:** none. Colours, fonts, icons and art come from the shared pipeline. `assets_test.go:95` fails on any hex literal in `styles.css`.

**SHOULD**
1. **The tray menu wording is hardcoded in the GUI layer.**
   - `companion/internal/shell/shell_gui.go:94-97` hardcodes "Open Forever Sixty", "Open my reports on the site" and "Quit", plus tooltips.
   - Elsewhere the same job has other names. The window says "Open all my reports on the site". The web nav says "Your reports" (`web/src/lib/account/connect-copy.ts`).
   - Fix: move the tray strings into `internal/app/copy.go` and use one name for the job.
2. **Formats are implemented twice.**
   - `copy.go:78-88` and `format.js:94,122-130` each reimplement `apps/shared/formats.json`, and neither reads it.
   - The round-4 `addon-ready.png` still shows "Today 18:40" and "1 minute ago". That capture is stale; the code now prints "6:40 pm" and "5 min ago".
   - Fix: replay `formats.json` cases in `format.test.js` and the `copy.go` tests, as `web/src/lib/formats.test.ts` does.
3. **The copy is not keyed to `apps/shared/copy`.**
   - Two stores: `copy.js` (265 lines) and `copy.go` (307 lines).
   - 0 keys come from `en.json` or `shared.json`.
   - Fix: key the shared nouns (reports, credited, device, addon) to `shared.json`.
4. **The DMG window is a light baked bitmap** with a muted teal arrow (`mac-dmg-window-dark.png`, `mac-dmg-window-light.png`, `packaging/macos/dmg-background{,@2x}.png`). The web and the window are dark-ground.
   - Fix: have `tools/art/dmg_background.py` render from `--color-ground` and `--color-bg`. Finder backgrounds cannot follow the system appearance, so ship one dark ground.
5. **Windows setup.** The web-view install page is on-brand (`win-1-setup.png` against `page-setup.png`). The light native title bar belongs to the OS and is allowed.

**NICE**
- 24 `font`/`font-size` literals in `styles.css` are not `var(--text-*)`: lines 93, 126, 132, 203, 213, 273, 283, 349 and 460, among others. Examples: `21px/1`, `22px/28px`, and `28px` for the pair code.
- Raw radii at `styles.css:306, 309, 461, 563`.
- Untokened veils and shadows:
  - `rgba(255,255,255,…)` at `styles.css:29, 110, 537, 563`.
  - `text-shadow rgba(0,0,0,.45)` at 97 and 129.
  - Fix: add `--shadow-heading` and `--color-veil-*` to tokens.json.
- `border: 1px solid black` at `styles.css:447`.
- `install.html:19` draws the seal at 32 px. Confirm that `seal.svg` is the flat variant. The identity sheet shows `companion/ui/seal.svg` as the flat turquoise seal, which is correct for that size.

## Shared-source consumption

| Source | Status |
|---|---|
| Colour tokens | Generated: `ui/tokens.css` from `build.mjs:320`; 0 hex literals |
| Fonts | The site's Marcellus, Open Sans, Archivo and JetBrains Mono woff2 files |
| Type scale | Mixed: about 24 literal sizes |
| Radii | Mostly tokens; 4 literals |
| Icons | Generated: `ui/icons.js` from `gen-ui-icons.mjs` (6 glyphs plus the corner) |
| Tray icons | `tools/icons/render.py`. The dot tones match the `--tone-*` CSS by convention only, with no test |
| Copy | Hand-written: `copy.js`, `copy.go` and tray literals |
| Formats | Reimplemented in `format.js` (360 lines) and `copy.go` |
| Installer art | Baked bitmaps (`dmg_background.py`) |

## Components

- **Buttons:** primary, secondary and quiet, in sentence case with the gold focus ring. Matches the web.
- **Rows:** a muted label and a strong value with hairlines, as on the web data plates. Matches.
- **Reports list:** the dot marker, name, meta and "Copy link". Matches the web account reports.
- **Lead panel:** carries the corner filigree. It is the most decorative element in the product and the web's dense plates have nothing like it. The owner should look at it.
- **Tray menu:** native, so allowed.

## Stale captures

- `2026-11-09-companion/round-4/*` predates `80e3da9f4` and `1b8b19ac0`.
- `2026-11-10-unify-server/after/companion-*` is current.

VERDICT: DRIFT

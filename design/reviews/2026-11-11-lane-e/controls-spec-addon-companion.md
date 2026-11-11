# Controls and states outside the web: the addon and the companion

Lane E, 11 Nov 2026. The brand audit's hand-off to E (`design/reviews/2026-11-10-brand-audit/report.md`, "Hand-offs: E") asks for the web's control and state spec carried to the addon's buttons and frames, and a call on the companion's lead-panel filigree. This is that spec. It changes no addon or companion code.

The web source of truth is `web/src/lib/ui/button.ts` (button recipes), `web/src/lib/ui/copy.ts` (every retry, copy and busy word), `web/src/components/ui/{EmptyState,LoadError,ConfirmButtons,CopyButton}.svelte` and `web/src/components/ui/panel.css` (plate and lead).

## 1. Buttons

| Web recipe | Web use | Addon today | Companion today | Rule |
|---|---|---|---|---|
| primary: teal ground, teal edge, white label | the view's one main action | none: `CONTROL_STATES.siteButton.selected` (`Theme.lua:775`) is a pressed toggle's look, not a primary | `button.primary` and `button.account` (`styles.css:347-349`), two classes for one recipe | One primary per frame. Teal, never client red, never gold. |
| secondary: blue ground, blue edge, white label | every other action | `Widgets.clientButton` at rest | `button.action` | The default button. |
| ghost: gold text, no box | a link-weight action | `Widgets.textButton` | `button.ghost` | Gold text only; hover to gold-hover. |
| danger: wipe ink, 40 % wipe hairline, never filled; hover takes the full wipe border | the confirm of a destructive ask | none | `button.danger` (`styles.css:350-351`): full wipe border, fills on hover; `danger-text` is a ghost variant the web lacks | Pair with a secondary Cancel that takes focus (web `ui/ConfirmButtons`). |
| a pressed toggle: the selected-mark edge with its ring | a choice among a set | `siteButton.selected` | none | Keep apart from primary. |

Sizes: web 36 px (sm, fine pointer) and 44 px (md), 20 px side padding. The addon's 20 to 22 px client buttons follow the client's UI scale and stay. The companion takes the web's sizes.

Addon gaps, for the addon lane:
- Add `CONTROL_STATES.siteButton.primary = { ground = "teal", edge = "tealEdge", hover = "tealEdge" }` and draw the frame's one main action with it; `selected` stays the toggle state.
- `Widgets.button` (raised ground, hairline, gold label) is a fourth look the web does not have. It has four callers: `surfaces/Settings.lua:435` (Reset), `surfaces/GuildTab.lua:1027` (Push), and `Widgets.confirmFrame` (`Widgets.lua:949`, `:957`), which `GuildTab.lua:1139` and `:1484` use for the push confirm. All move to `Widgets.clientButton`: Cancel secondary and first, the confirm secondary (the push destroys nothing); then `Widgets.button` is deleted.
- `Widgets.edgedButton` (panel fill under a 1 px gold edge) stays: it is the tab hung on the client's Character, Guild and Group frames, and the spec treats it as the web's tab.
- Busy: no addon button waits on the network today. One that does follows owner ruling 5, a spinner plus a busy word ("Sending…").

Companion gaps, for the companion lane:
- Fold `button.account` into `button.primary`.
- The Pair submit (`app.js:584`, `action`) is its view's one main action: make it `action primary`. Check `install.html` / `install.js` against the same rule.
- Drop the 40 px / 15 px `big` step: 36 px for small, 44 px for the main action, 20 px padding.
- `button.danger`: a 40 % wipe border at rest, the full wipe border on hover, no fill.

## 2. Frames

| Web | Addon | Companion | Rule |
|---|---|---|---|
| plate: flat `--color-raised` ground, 1 px black edge plus a 1 px panel-edge ring, plate radius | `Widgets.umberFrame` (`Widgets.lua:991-1000`): an `umberTop` to `umber` gradient inside one `umberEdge` hairline | `.panel` matches | Every window and panel is a plate: flat ground, two hairlines. No client bevel (no `BackdropTemplate` dialog border) on any frame we own. |
| framed: the plate plus the cream line inset 6 px (`.ui-framed`); lead: framed plus the four filigree corners (`.ui-lead`, `ui/FrameCorners.svelte`, 64 px, 44 px under 640 px) | none | `.panel.lead` merges framed and lead, corners fixed at 44 px | At most one lead per page or window. |
| overlay: tooltips, menus, sheets (`ui/Overlay.svelte`) | the client's GameTooltip, restyled by `Theme.compareTooltip` | none | The game tooltip stays the game's; the web's tooltip kind copies it. |

Addon gap: `umberFrame` drops its gradient for the flat raised ground and draws the second hairline (black outside, panel edge inside). No bevelled frame remains in the addon's own windows; the client art left is on frames we hang tabs from.

## 3. States

The words come from one place per surface and match the web's `lib/ui/copy.ts`:

| State | Web | Addon (`Locale.lua`) | Companion (`ui/copy.js`) |
|---|---|---|---|
| retry | "Try again" | none needed (no network) | `tryAgain: 'Try again'` matches |
| copied | "Copied", clears after about 2 s | "Selected. Press Ctrl+C" (the client has no clipboard API; correct, stays) | clears after 2 s already (`app.js:157-161`); drop the full stop in `copy.js:117` and `:149` ("Copied") |
| other confirmations | n/a | n/a | `note()` clears success text after 4 s (`NOTE_MS`); that is a notice, not the copy confirmation, and may stay |
| loading | "Loading…" | `gearItemLoading = "Loading item..."` (`Locale.lua:261`): three dots become "…" | matches |
| empty | one muted sentence, then the one action that fills it (`EmptyState`) | the `*None` / `*No*` strings already read as one sentence | same |
| error | the muted sentence and a secondary "Try again" (`LoadError`) | n/a | the bar's ghost button after the sentence: matches |

## 4. The companion's lead-panel filigree (owner call)

The brand audit called the companion's lead-panel corners "the most decorative element in the product" and said the web's dense plates have nothing like it. That is no longer true: the web's lead plate (`.ui-lead`, `ui/FrameCorners.svelte`) carries the same filigree, the same path (traced from `design/assets/ornaments/corner-filigree.svg`), the same cream at 80 % and the same 4 px inset. The companion draws it at 44 px, the web's size under 640 px, and right for a 480 px window.

Options for the owner:
1. **Keep it (recommended).** The companion's Status headline and Device account are each the window's one lead, as the web's home lead is the page's. Same ornament, same rule: one lead per view.
2. **Drop it in the companion only.** The window becomes plates only. The web keeps its lead, so the two surfaces would then differ.
3. **Drop it everywhere.** Remove `FrameCorners` from the web lead and `frameCorners` from the companion. The lead keeps its inset cream line alone.

Whatever the call, the addon does not take the filigree: its frames sit on the game's own art, and a second ornament there competes with Blizzard's.

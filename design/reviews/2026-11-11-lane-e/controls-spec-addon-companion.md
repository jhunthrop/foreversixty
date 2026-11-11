# Controls and states outside the web: the addon and the companion

Lane E, 11 Nov 2026. The brand audit's hand-off to E (`design/reviews/2026-11-10-brand-audit/report.md`, "Hand-offs: E") asks for the web's control and state spec carried to the addon's buttons and frames, and a call on the companion's lead-panel filigree. This is that spec. It changes no addon or companion code.

The web source of truth is `web/src/lib/ui/button.ts` (button recipes), `web/src/lib/ui/copy.ts` (every retry, copy and busy word), `web/src/components/ui/{EmptyState,LoadError,ConfirmButtons,CopyButton}.svelte` and `web/src/components/ui/panel.css` (plate and lead).

## 1. Buttons

| Web recipe | Web use | Addon today | Companion today | Rule |
|---|---|---|---|---|
| primary: teal in its teal edge, white label | the view's one main action | `Widgets.clientButton` selected (`CONTROL_STATES.siteButton.selected`) | none: `button.action` is drawn blue | One primary per frame. Teal, never client red, never gold. |
| secondary: blue in its blue edge, white label | every other action | `Widgets.clientButton` at rest | `button.action`, `button.account` | The default button. |
| ghost: gold text, no box | a link-weight action | `Widgets.textButton` | `button.ghost` | Gold text only; hover to gold-hover. |
| danger: wipe ink, 40 % wipe hairline, never filled | the confirm of a destructive ask | none | none | Use when the addon or companion gains a destructive confirm; pair with a secondary Cancel that takes focus. |

Addon gaps, for the addon lane to close:
- `Widgets.button` (raised ground, hairline, gold label) is a fourth look the web does not have. Its two callers (`surfaces/Settings.lua:435` Reset, `surfaces/GuildTab.lua:1027` Push) move to `Widgets.clientButton` at rest (secondary), and `Widgets.button` is deleted.
- `Widgets.edgedButton` (the panel fill under a 1 px gold edge) stays: it is the tab hung on the client's own Character, Guild and Group frames, where it must read as ours on Blizzard's art. It is a tab, not a button, and the spec treats it as the web's tab.
- Sizes: the client button is 20 to 22 px tall against the web's 36 / 44 px. That is right for the client's UI scale; keep the proportions (6 px radius equivalent is square in Lua, which is accepted), not the pixels.
- Busy: the addon has no request buttons that wait on the network. If one is added, it follows owner ruling 5: a spinner plus a busy word ("Sending…"), from shared copy.

Companion gap, for the companion lane: the window's one main action (Sign in on the signed-out Status, Install on `install.html`) takes the primary recipe, a `button.action.primary` in teal on the teal edge; every other `button.action` stays secondary.

## 2. Frames

| Web | Addon | Companion | Rule |
|---|---|---|---|
| plate: raised ground, 1 px black edge plus 1 px panel-edge ring, plate radius | `Widgets.umberFrame` (umber gradient inside one umberEdge hairline, no bevel) | `.panel` | Every window and panel is a plate. No client bevel (`BackdropTemplate` with the dialog border) on any frame we own. |
| lead: the plate plus the cream line inset 6 px and the four filigree corners (`ui/FrameCorners.svelte`, 64 px, 44 px under 768) | none | `.panel.lead` with `app.js frameCorners` at 44 px | At most one lead per page or window. |
| overlay: tooltips, menus, sheets (`ui/Overlay.svelte`) | the client's GameTooltip, restyled by `Theme.compareTooltip` | none | The game tooltip stays the game's; that is the in-game frame the web's tooltip kind copies. |

The addon's frames already follow this since BA1 (`umberFrame`, `clientCloseButton`). No bevelled frame remains in the addon's own windows; the remaining client art is on frames we hang tabs from.

## 3. States

The words come from one place per surface and match the web's `lib/ui/copy.ts`:

| State | Web | Addon (`Locale.lua`) | Companion (`ui/copy.js`) |
|---|---|---|---|
| retry | "Try again" | none needed (no network) | `tryAgain: 'Try again'` matches |
| copied | "Copied", clears after about 2 s | "Selected. Press Ctrl+C" (the client has no clipboard API; this is correct and stays) | `copied: 'Copied.'`: drop the full stop to match the web, and clear it after 2 s |
| loading | "Loading…" | `gearItemLoading = "Loading item..."`: three dots become the ellipsis "…" | matches |
| empty | one muted sentence, then the one action that fills it (`EmptyState`) | the `*None` / `*No*` strings already read as one sentence; keep them sentence case, with the fix where the client allows it | same |
| error | the muted sentence and a secondary "Try again" button (`LoadError`) | n/a | the bar's ghost button after the sentence: matches |

## 4. The companion's lead-panel filigree (owner call)

The brand audit called the companion's lead-panel corners "the most decorative element in the product" and said the web's dense plates have nothing like it. That is no longer true: the web's lead plate (`.ui-lead`, `ui/FrameCorners.svelte`) carries the same filigree, the same path (traced from `design/assets/ornaments/corner-filigree.svg`), the same cream at 80 % and the same 4 px inset. The companion draws it at 44 px, which is the web's size under 768 px, and right for a 480 px window.

Options for the owner:
1. **Keep it (recommended).** The companion's Status headline and Device account are each the window's one lead, as the web's home lead is the page's. Same ornament, same rule: one lead per view.
2. **Drop it in the companion only.** The window becomes plates only. The web keeps its lead, so the two surfaces would then differ.
3. **Drop it everywhere.** Remove `FrameCorners` from the web lead and `frameCorners` from the companion. The lead keeps its inset cream line alone.

Whatever the call, the addon does not take the filigree: its frames sit on the game's own art, and a second ornament there competes with Blizzard's.

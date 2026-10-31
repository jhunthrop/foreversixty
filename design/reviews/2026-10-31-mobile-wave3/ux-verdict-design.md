# Wave 3 mobile design: ux-designer verdict

Reviewer: ux-designer, 2026-10-31. Reviewed what renders: the 16 boards in `design/mocks/2026-10-31-mobile-wave3/renders/` (8 features, iPhone 17 and Pixel 9) against `design/specs/2026-10-31-mobile-wave3.md`. This is a design-stage review: it checks that the boards say what the spec says and that the spec leaves no engineer a wrong reading. Build captures are reviewed later, lane by lane, against the same boards (spec §12).

## Method

Every board was opened and read region by region. The spec's tables (§4 pieces, §5 per feature, §6 labels, §7 states, §9 collapses) were checked against the pixels. Findings from the first pass were fixed in the generator and the board re-rendered before this verdict was written.

## Deviations found and fixed in this pass

| Element | Spec said | What first rendered | Severity | Fixed |
|---|---|---|---|---|
| F13.2 refusal line | `Needs {n} points in {Tree}. You have {m}.` must be true | `Needs 20 points in Elemental. You have 21.`: a refusal contradicting its own count | blocker | The low-points build now has 16 points; the line reads `You have 16.` |
| F11.3 pinned row | rank 204 with a plausible figure | `-873 DPS` | blocker | Figures now fall off as a power of the rank; 204th reads 203 DPS |
| F10.3 meters | one number with its unit (tenet 7: a number without its unit is a defect) | values with no unit | must fix | `DPS` in `micro` beside each value; spec §5.2 F10.3 states it |
| F11.4 recent fights, F14.1 medium upgrades | unit on every figure | `541`, `+5.3` bare | must fix | `DPS` added |
| F16.3 history rows | `You were there` visible, meta on one line | tag crushed at the right, `3 h 12 m` broke across two lines | must fix | Tag moved under the meta, meta set no-wrap, rows top-aligned |
| F16.5 access off | the switch is off when access is off | switch drawn on | must fix | Switch off in that frame |
| F14.1 medium upgrades widget | footer line `{n} upgrades · +{total} DPS` fully visible | footer clipped by the card's edge | must fix | Row icons 28, tighter rows |
| F9.3 and F12.2 | `1 rank`, not `1 ranks` | `1 ranks` | polish | Plural rule applied |
| F14.4 Dynamic Island | compact pill and expanded island drawn apart, each labelled | the two overlapped | must fix | Expanded island lower, labels added |
| F15 Siri and Spotlight text | the app's type (Barlow fallback) | system-font fallback drew serif | polish | Font stack fixed |
| F10.4 scrolled report (Android) | content stops at the bar | scrolled content showed behind the status bar | polish | Status area covered |
| v1 screen changes (§1) | Home gains `Logs` and `Rankings`; Raids gains `Roster` and `Raid history`; search on root bars | not on any board | must fix | F10.8 and F16.7 added to the boards and the spec's board table |

## Region by region, after the fixes

| Board | Verdict on the pixels |
|---|---|
| F9 Guides | The answer is first: `Your guide` is the one card with the warm surface and sits above the fold with the class grid (F9.1). The guide's at-a-glance card (role, seven numbered stat chips) is the first thing under the head (F9.2). Every talent and spell carries its icon; rotation lines carry the rank in mono (F9.3, tenet 3). The tier list leads with the callout, not the list; your row is tagged `Yours`; ties show `≈` (F9.4). Offline shows saved and not-saved guides distinctly (F9.6). |
| F10 Reports | `Your parses` is the one surface card on the report; trash is one entry (`Trash · Onyxia's Lair`, `9 pulls · 14:20`); the fight sheet is Details!-shaped with your row marked (F10.3). Empty, processing, failed, not-yours and offline all have words (F10.5, F10.6). |
| F11 Rankings, character | The answer card carries `14th of 212 Elemental Shaman characters`; the filter collapses to one summary line plus a 44 pt button; your row is marked in place and pinned only when off the page (F11.1, F11.3). The score card keeps one primary and the Paperdoll follows the in-game frame with the gold dot (F11.4, F11.5). `No logs yet.` is not a zero (F11.6). |
| F12 Search, item | Empty search is useful (for your character, watching, recent). Results are grouped and capped. The item page opens with the name in quality colour and `Your Back pick`, then the tooltip, then sources and the specs that list it with `Pick` or `Runner-up` (F12.4, F12.5). |
| F13 Planner | The tree is the primary; `Share` is the only gold button; the stepper strip tells you what the last tap did; refusals are in words (F13.1, F13.2). The share confirm is the web's text verbatim (F13.3). The viewer shows the build's own totals (F13.5). |
| F14 Widgets | Every small state has words; no spinner, no empty box (F14.2). The medium widget answers In, Late, Out without opening the app. Lock screen widgets are monochrome as the system requires (F14.3). The Live Activity and ongoing notification carry the countdown and the answer. |
| F15 System | Siri confirms once and answers with the count; the upgrade answer carries the item icon and quality colour (F15.1, F15.2). Offline data has no build ids or counts (F15.4). The three strips (offline, stale, new data) are distinct components of one strip (F15.5). |
| F16 Guild | Counts card, role chips, rows with spec icon, rank and attendance (F16.1). History rows say `You were there` (F16.3). Calendar sync and link are separate parts with one switch each (F16.4). |

## Checklist (spec §13)

- [x] Five tabs; search on Home, Raids, BiS, Guides (F10.8, F16.7, F9.1).
- [x] Order and letters come from data, not the app (spec §5.1, §10 W3-C1).
- [x] Every spell, talent and item has its icon; rank on rotation lines; totals summed from the build.
- [x] New data never moves a screen under the reader (F15.5).
- [x] Reports: trash one entry; parses first; coach only with `log_coach`, no upsell (F10.2, F10.4).
- [x] Rankings: pin only off the page; no per-row crest on a single-class board.
- [x] `No logs yet.`, never zero.
- [x] Search offline for items, guides, talents; characters say they need a connection (F12.6).
- [x] No `0%` drop chance; unknown prints nothing.
- [x] Planner: tap, long press, refusals in words, confirm text verbatim, shared build is copied not edited.
- [x] Widgets: every state has words; answers without opening the app; stale says so.
- [x] No mobile-only Premium (spec §1.1).
- [x] Share links only for things with a web page.
- [x] Offline, stale, new-data are distinct.
- [x] Every new element has the §7 states; touch targets and labels per §6.
- [ ] Planner rules against golden vectors, budgets on device, contract tests: build-time, checked on each lane's captures.

## Remaining, for the lanes (none blocks the design)

1. **Planner class label.** F13.1 shows the class crest only. The spec gives the crest an accessibility label; a lane may add the class name to the chip if it fits at 360 dp. Polish.
2. **Tooltip `Binds` line on the item page.** The boards omit it because the v1 data file has no `bonding` for the Cape. Spec §5.4 requires the tooltip to print it when W3-C8 and the pending C6 fields publish. The F12 lane's capture must show a Binds line.
3. **Widget on the lock screen** is drawn white-on-grey by design (the system tints it). The lane must verify the tinted rendering on device.
4. **Largest-text boards** for Wave 3 are not drawn. Spec §9 states the collapses; each lane attaches AX5 and 2.0 captures, and the ux-designer reviews them before merge. If a lane cannot make §9 hold, it asks for a board.

## Verdict

**SHIP.** The boards match the spec, every blocker and must-fix from the first pass is fixed, and the spec leaves no state or copy line undefined. The four items above are tracked in the lanes' own captures.

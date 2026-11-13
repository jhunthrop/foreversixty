# Addon source filter, review round 1 (2026-11-13)

Boards: this folder, from `design/mocks/addon_sources.py`. Spec: `design/specs/2026-11-13-addon-source-filter.md`.

| Reviewer | Verdict |
|---|---|
| ux-designer | FIX: the rep-standing board contradicted the leveling board; set-aside caption named dps_delta; spec copy for the set-aside line; empty board's age stamp; mixed fonts across boards; a test for combat close |
| wow-player | FIX: the boards did not show the owner's case (world drops) being solved; the rep-standing board read as the ring's own source; gains looked invented; cut the flyout foot line; mute zero-count rows |
| ai-slop-reviewer | FIX: cut the Settings hint and the flyout foot line (explainers); shorten the set-aside line to the Next line's format |
| info-density-reviewer | FIX: cut the flyout foot line; show what each kind brings (its best item per slot) on the row; mute zero-count rows; no filtered set DPS (heap and one job per page) |

What changed for round 2:

- Copy: no flyout foot line, no Settings hint; the set-aside line reads `All sources: <item> +1.1 DPS`, priced in game from the two items' stats.
- A kind with nothing at the band is muted, still a switch.
- Each Sources row's hover lists the best item that kind gives each slot (`Sources.kindItems`), from data the addon already carries.
- The set-aside line sits directly under the source line in both tooltip layouts.
- A spec pins that combat closes the flyout with the panel.
- Boards: one font set, real gains from item stats and the band's weights, a World / BoE off board (the owner's case), a row-hover board, an unambiguous rep-standing board.

Kept, with reasons:

- Default stays the web's (every kind but Reputation above Friendly below 60): the brief says to match the web. World / BoE is one click off, and the new board shows it.
- `No rep above Friendly`, `No %s from these sources at this level` and `Standings up to Friendly still count.` are the web's words.
- The `-` count shows only when the band has no counts at all; a real zero reads `0`.
- `Change` is disabled in combat in code, but the panel closes when combat starts, so the state never shows.
- The sim's age stays on the empty Upgrades heading: it dates the weights, which still hold.
- Re-picked dungeon items carry no drop chance (the data does not carry a chance per by-kind entry, for the heap).

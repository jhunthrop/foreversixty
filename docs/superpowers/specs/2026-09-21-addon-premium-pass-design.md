# Addon premium pass

**Date:** 2026-09-21. **Status:** approved by the owner's instruction ("review the addon as
the world's best WoW addon UI builder and make it feel premium, really showing the value of
our platform"). Builds on `2026-09-20-addon-ui-design.md`; where they differ, this wins.

## Why

The first in-game screenshots showed a working addon that reads as a settings dialog: it
opens on an unreadable export string, every line has the same weight, there are no icons,
colours of meaning or progress, and nothing of the platform's value (the next talent to
take, bag upgrades scored by our weights, sync status) is visible without hunting through
tabs. Nothing exists outside the window.

## Rules that do not change

- No network, ever (`no_network_spec.lua`). No protected actions in combat. Every client
  API is guarded or goes through `Compat`/`Theme`: nobody can run the game here, and the 1.60
  client (Interface 16001) has already been missing three functions we assumed.
- No stock templates (`Theme.TEMPLATES` is empty): the addon draws its own chrome.
- An edit box never holds the keyboard unless the player pressed Copy or clicked into it.
- Honest copy, no exclamation marks, no emoji. Strings live in `Locale.lua`.
- Pure models (tables and strings) are separate from frames and are what the specs test;
  frame code is thin. Files stay under 400 lines.

## Shell (coordinator builds)

Window 720 x 500. Title bar: mark, wordmark, addon version, close. Left sidebar (150 wide):
Overview, Talents, Gear, Export, Settings; the active item has a gold bar on its left edge
and a raised ground; the data build and `foreversixty.gg` sit at the sidebar's foot. Content
header: character name in class colour, "Level N Race Class · Spec", and a status pill (data
build, or a warning when the game build differs). Drop shadow, hairline border, 0.15 s fade
in, the client's own open and close sounds. One primary button style (gold ground, dark
label) per page; everything else is secondary.

**Overview** (the landing page), four cards:
1. *Your build*: name, progress bar, "Next: <talent>, rank x of y"; or how to load one.
2. *Gear*: slots filled, planned pieces equipped, upgrades waiting in bags.
3. *Talent points*: one bar per tree, and points unspent when there are any.
4. *Send to the site*: the Copy button, where to paste, when the export last saved.

**Talents**: build name and progress bar; rows with the talent's icon (from the client's
trait or talent API when it has one, a neutral tile when not), rank as "2 / 5", the next
talent marked with a gold edge, finished ones dimmed. Loading a build lives in a section at
the top when nothing is loaded and at the foot when something is.

**Gear**: slot rows with item icon and quality colour for planned and equipped, a tick
where they match; bag upgrades with the score gain and Equip.

**Export**: three numbered steps, the primary Copy button, the code in a one-line field.

## Outside the window (parallel lane builds; owns only the files named)

Owns: new `Tooltip.lua`, `Toast.lua`, `Bindings.xml` (and the `Locale.lua` strings and TOC
lines they need), `Tracker.lua`, `Minimap.lua`, `Options.lua` (slash help only), and their
specs. Must not edit `Window.lua`, `Widgets.lua`, `Theme.lua` or anything under `views/`;
uses only Theme's existing drawing API (`texture`, `outline`, `paint`, `fontString`,
`rgb`, `classColor`).

1. *Item tooltips*: on any item tooltip, at most two lines under a "Forever Sixty" heading:
   "Planned for your <slot>" when the loaded build wants that item, and "Upgrade for <slot>:
   +N by our weights" or "Not an upgrade" for equippable items when weights exist for the
   spec, using `Gear`'s existing scoring (no second scorer). A setting turns it off. Hook
   through `TooltipDataProcessor` when the client has it, else `OnTooltipSetItem`; never
   both; never error inside a tooltip hook (pcall, and disable the hook after a failure
   with a diagnostic).
2. *Level-up toast*: on `PLAYER_LEVEL_UP` and whenever an unspent talent point appears while
   a build is loaded: a small top-centre panel, "Level 12. Take Improved Rend, rank 2 of 3",
   which fades after a few seconds, never in combat (queue it until combat ends), click to
   open the Talents page. A setting turns it off.
3. *Tracker*: the same information as a compact bar: next talent, rank, a thin build
   progress bar; hidden when no build is loaded; still draggable and lockable.
4. *Keybind*: `Bindings.xml` with "Toggle Forever Sixty" under its own header.
5. *Minimap button*: tooltip shows build progress and upgrades waiting; left click toggles,
   right click opens Settings; supports the addon compartment when the client has it.
6. *Slash help*: `/fs` alone opens the window; `/fs help` prints the commands in gold.

## Verification

Specs for every model and every guard. The in-game checklist in `addon/README.md` gains a
row per visible feature. Nothing is claimed as verified in game until the owner tests it.

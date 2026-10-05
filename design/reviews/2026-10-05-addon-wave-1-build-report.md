# Forever Sixty addon — wave 1 build report

Branch `addon-wave-1`. Built against `design/specs/2026-10-04-addon.md` and the five boards
in `design/mocks/renders/`.

## Shipped

**§10 rulings (all five):**

1. BiS row now draws for any item whose equip location maps to a tracked slot
   (`Tooltip.bisSlotFor`), not only an exact-worn match — `Tooltip.sections` calls it in place of
   the old `equippedSlotFor`-only gate.
2. Overview's arrival-banner gap collapses to `Theme.SIZES.padding` when nothing is pending
   (`OverviewView.gridTop`); the 2×2 grid, rating card and rotation card re-anchor on every
   `apply()` via a new `reposition()`.
3. Build-name fallback is title-cased (`Talents.titleCase`, shared by `Window.specLabel` too —
   consolidated, was duplicated).
4. Tiers are 1-indexed wherever shown to a player (`Talents.displayTier`): `Follow.line`,
   `Toast.model`'s unknown-cell fallback, `FollowView.rows`.
5. `GearView`'s "UPGRADES IN YOUR BAGS" shows `L.gearNone` when the list is empty.

**Window / boards:**

- Circular ringed class crest: `addon/tools/make_class_crests.py` renders the nine class crests
  (already-circular source webps, just resized) plus one shared ring texture, tinted per class at
  runtime (`Theme.buildCrest`/`Theme.paintCrest`). No `MaskTexture` dependency — the source asset
  is pre-cropped round, so there is no square card left to mask. Wired into the window header and
  the Overview's "Send to the site" card (28px).
- Overview cards: `Your Build`/`Gear` now lead with one large gold figure (`card.title` repurposed
  via `Theme.applyFont(..., "large")`), caption beside it, detail line under. Raw "N of 17 slots
  filled" figure removed per the round-2 ruling; empty state shows a plain `-` placeholder, never
  an invented count.
- Talent points' three bars are class-coloured (`Theme.paintRGB`, new), not plain gold.
- Send-to-the-site card always shows crest + class-coloured name + level, Copy enabled except on
  `Export.string`'s genuine class-unknown refusal.
- Personal rating's empty state (installed, not yet rated) reads "Upload a log or run the
  companion to get one." (`overviewRatingNone`); dead `ratingsNotRated`/`overviewSyncPreview`/
  `OverviewView.preview` removed.
- Overview Gear card and the Gear page now read the exact same model: `model.gear.upgrades` is
  `#rows.upgrades` from `GearView.rows`, not a second, independently-filtered count
  (`OverviewView.waiting` deleted).

**Gear page (§4.5.4):** slot-label column (`L.gearSlotLabels`), equal 232px Planned/Equipped name
columns, green tick glyph (`Theme.MEDIA.gearTick`, the client's own ready-check texture) replacing
the "As planned" text tag (tooltip carries the words instead), "No plan for this slot" for an
unplanned-but-equipped slot (new `GearView.noteFor` branch) with its own wider 96px strip so the
prose never overlaps the item name. Item hyperlink tooltips on every Planned/Equipped cell were
already wired (`Widgets.itemRow`'s existing `attachTooltip`) — confirmed, not re-built.

**Guild tab:** officer/leader-only gating of the whole private state block (`GuildView`'s
`OFFICER_RANKS`) — a member now sees none of it, not even the claim-state half. New member
standing line (`GuildView.standingLine`), shown first on every panel: the real top gear gap
(`Gear.upgrades`, same model the Gear page uses) plus an honest "not in your data addon yet" for
the item-level rank half, since `ForeverSixtyData`'s schema carries no `item_level`/`spec` per
member today. Roster rows showing crest/spec/item-level are **not built** — blocked on that same
missing data-addon schema (different repo area; noted, not invented).

**Companion inbox upgrade line:** `OverviewView.upgradeLine` reads the newest "upgrade" inbox
message, renders "Your top upgrade: `<link>` for `<slot>`, +N DPS, from your sim (`<date>`)"
(`L.inboxUpgradeLine`, previously dead, now wired); absent message renders nothing, page height
collapses to meet it.

## Tests

```
716 successes / 0 failures / 0 errors / 0 pending
```

```
luacheck ForeverSixty tests: 0 warnings / 0 errors in 61 files
```

Every pure-model change above has a spec: `talents_spec.lua` (titleCase/displayTier),
`theme_spec.lua` (crestPath), `window_spec.lua` (crest+ring paint), `tooltip_bis_spec.lua`
(any-tracked-slot BiS row), `overview_view_spec.lua` (gridTop collapse, upgradeLine, big-figure
fields, class-coloured bars, sync card), `gear_view_spec.lua` (slot label, tick glyph, no-plan
width, empty-state copy), `guild_view_spec.lua` (officer gating, standingLine).

Parse-checked shipped files with the host's Lua 5.5 `luac -p` as a best-effort proxy (clean); the
real CI job (`syntax-lua51`, Lua 5.1.5) was not run locally — no 5.1 interpreter installed in this
environment — but no 5.1-incompatible syntax was introduced (no goto, bitwise ops, integer
division, `<close>`).

## Needs an in-game spike check

Added as README rows 29–33: crest/ring render round with a correct class-colour ring; the
big-figure font (`GameFontNormalLarge`) actually reads as dominant; the gear tick glyph
(`Interface\RaidFrame\ReadyCheck-Ready`) exists and looks right at 14px; Gear page column widths
clear the longest band-20 name at 1x/1.15 without truncation; the guild message's real `rank`
field spelling matches `OFFICER_RANKS` (`officer`/`leader`).

## Files changed

```
addon/ForeverSixty/.pkgmeta
addon/ForeverSixty/Follow.lua
addon/ForeverSixty/Locale.lua
addon/ForeverSixty/Talents.lua
addon/ForeverSixty/Theme.lua
addon/ForeverSixty/Toast.lua
addon/ForeverSixty/Tooltip.lua
addon/ForeverSixty/Window.lua
addon/ForeverSixty/views/FollowView.lua
addon/ForeverSixty/views/GearView.lua
addon/ForeverSixty/views/GuildView.lua
addon/ForeverSixty/views/OverviewView.lua
addon/ForeverSixty/media/crests/{druid,hunter,mage,paladin,priest,ring,rogue,shaman,warlock,warrior}.tga  (new)
addon/tools/make_class_crests.py  (new)
addon/README.md
addon/tests/follow_spec.lua
addon/tests/follow_view_spec.lua
addon/tests/gear_view_spec.lua
addon/tests/guild_view_spec.lua
addon/tests/overview_view_spec.lua
addon/tests/talents_spec.lua
addon/tests/theme_spec.lua
addon/tests/toast_spec.lua
addon/tests/toc_spec.lua
addon/tests/tooltip_bis_spec.lua
addon/tests/window_spec.lua
```

No `.toc` line changes needed (no new `.lua` files — media is unpackaged by path per `.pkgmeta`,
confirmed by `toc_spec.lua`'s new disk-presence check).

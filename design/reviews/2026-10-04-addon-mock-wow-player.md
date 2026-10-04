# Addon mock review — wow-player — 2026-10-04

Reviewed: `addon-tooltip.png`, `addon-overview.png`, `addon-talents.png`, `addon-gear.png` against
`design/specs/2026-10-04-addon.md` and `docs/tenets.md`. Worked character: Obnoxious Yell, lvl 23
Human Fury Warrior, warrior-fury 20-29 alliance band.

**Would I run this?** Yes, both leveling and at 60. The tracker+toast are exactly the "glance and
go back to playing" WeakAuras discipline I already run, the tooltip tells me in one line whether to
loot/buy/equip, and the window is a real inspect panel, not a dashboard I'd never open. Nothing
here would make me `/fs` it off like I eventually do to half my bag addons.

## Verdicts

| Surface | Verdict |
|---|---|
| Tooltip | SHIP |
| Tracker + toast | SHIP |
| Window — Overview | FIX FIRST |
| Window — Talents | SHIP |
| Window — Gear | FIX FIRST |

## Findings

1. **[blocker]** Overview's "TALENT POINTS" card reads `Fury 11` right next to the "Your build"
   card's own `7 of 11 points taken` on the same screen. 11 is the build's *target* allocation; the
   character has actually spent 7. The real character pane's talent tab never shows a plan, only
   what you've spent — this is the one number on the board I'd screenshot to guild chat as "the
   addon is lying to me," because the contradiction is visible without even changing pages. **Fix:**
   wire this card to live spent-per-tree counts (`0 / 7 / 0`), not the build's target total.
2. **[must fix]** Gear page item-name cells (Planned and Equipped columns) are icon + quality-
   coloured text, which is right, but nothing in §4.5.4 confirms they carry a real hyperlink with
   hover-tooltip. Tenet 2 is explicit: *wherever* an item appears, hover shows the real tooltip. The
   main tooltip region and the Talents page's BiS row both do this; the Gear list must too, or it's
   the one place on the addon where an item is "just a name." Confirm `GearView` rows call
   `SetHyperlink`/`SetItemByID` before ship.
3. **[must fix]** The spec's own ruling 1 (§10.1) — the BiS row only draws on a byte-exact worn-item
   match, never on a bag item being compared against the same slot — is the single biggest gap
   against Pawn on the signature surface, and it's the one thing Pawn *always* does (any item,
   any slot, instant verdict + "get this instead"). The board already draws the fixed behaviour; the
   real hook does not yet. This must land before the in-game build matches what's been approved here.
4. **[polish]** Overview's Gear card shows two different denominators 6px apart — `14 of 17 slots
   filled` up top, `11 of 15 planned pieces equipped` in the foot bar — with no word explaining why
   (2 unsourced trinkets + the unworn Legs piece). Internally consistent once you do the math, but a
   player reading fast will wonder why the numbers disagree. A one-word label ("of planned") on the
   second line would fix it cheaply.
5. **[polish]** The rotation card — one of only six Overview regions — sits below the fold on first
   open (measured 420px of content in a 416px box). The soft-fade hint is a reasonable call per
   tenet 4 (scroll, not a disclosure click), but WeakAuras/ElvUI panels pair a fade with a visible
   scrollbar thumb; a fade alone under-signals that there's a 7th line of content.
6. **[polish]** Two fixes the spec itself already flagged as owner rulings — the 84px dead gap above
   the Overview card grid (ruling 2) and the empty "UPGRADES IN YOUR BAGS" section growing no copy
   (ruling 5) — are drawn *already fixed* on the boards. Good self-QA, but both are still pending
   owner sign-off and code changes, not shipped. Don't let "the mock looks right" stand in for "the
   build matches it."

## Data checked against `addon-items.json` / `guides-warrior-fury.json`

- `Upgrade for head: +20 by our weights` — correct. Head pick has 10 strength (band weight
  2.0089520256251725/str, no stamina weight this band) = 20.09, rounds to `+20` under `%+.0f`.
- `Downgrade -28.1%` — correct. Worn Mutant Scale Breastplate (10 str) scores 20.09; the alt
  Veteran's Chain Shirt (7 str, 4 agi, 5 stam — stam unweighted) scores 14.45; delta = -28.05%,
  rounds to -28.1%.
- `7 of 11` — correct. Band's 11-point Fury build is Cruelty 5/5 + Booming Voice 3/3 + Iron Will
  0/3; shown mid-progress (Cruelty 5/5 done, Booming Voice 2/3 not yet taken, Iron Will 0/3) = 7
  spent. Consistent with the toast firing to prompt the 3rd Booming Voice point.
- `14 of 17 slots filled` / `11 of 15 planned pieces equipped` — internally consistent (17 total
  slots − 2 unsourced trinkets = 15 "planned"; Legs is planned but unworn), but not independently
  checkable from the two data files alone since bag/trinket contents aren't in them — flagged as
  finding 4 for the UI confusion, not as a wrong number.
- Tooltip anatomy, stat order, quality colours (blue=rare on all three variants, matching
  `quality: 3` in the data) and the UNRATED variant's correct silence (no BiS row, no sourced
  trinket) all check out.

## Tenet-specific answers

- **Pawn comparison:** at a glance, yes, as good — colour-coded single verdict line, readable in
  under a second. Beyond the glance this beats Pawn: real icon, quality-coloured BiS link, a source
  line, a live second tooltip for the recommended item. What Pawn still does that this doesn't
  (until finding 3 lands): score *every* item you hover against a tracked slot, not only the one
  already worn.
- **720×500 earns its space:** mostly yes — it's the ElvUI config-shell reference, not a Details!
  HUD, so dense rows of small text are the right idiom; the Talents and Gear pages back every row
  with an icon and colour. Overview is the weakest of the three: no item icons at all on that page,
  and every number is the same weight, so nothing reads "from across the room" the way `14/17` bold
  almost does but doesn't quite commit to.
- **Overview honesty:** real, not filler — Send to the Site (`There is nothing to send yet.`) and
  Personal Rating (`No rating yet for this character.`) are both genuine empty states for this
  character, not placeholder copy.
- **Character header / crest:** correctly absent. The spec carves this out explicitly — the addon
  has no web asset pipeline, so it uses the client's own class-colour-on-name convention
  (`#c79c6e` on "Obnoxious Yell"), the same idiom every real addon (Details!, Recount, unit frames)
  uses in-game. This is not a tenet 7 violation; tenet 7's crest is a web-only rule.
- **Clipped/truncated/hidden:** nothing clipped on any of the four boards. The rotation card is the
  one thing below the fold — see finding 5.
- **Tracker + toast:** the two-line tracker is exactly right for questing — title, progress, bar,
  nothing else, draggable, lockable. Toast copy `"Level 23. Take Booming Voice, rank 3 of 3."` is
  verbatim-correct against the locale format and the data.

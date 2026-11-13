# Builds window, round 1 (2026-11-13)

Verdicts on the first boards: ux-designer SHIP WITH FIXES, wow-player SHIP WITH FIXES, ai-slop-reviewer FAIL, info-density-reviewer FAIL, brand-consistency-reviewer PASS with fixes.

## Fixed for round 2

- The boards printed `\194\183` raw: the mock's Locale reader now decodes Lua byte escapes (the game always did). All boards re-rendered.
- The column rule is drawn on the boards (the addon draws it in `BuildDialog.buildInto`).
- The good-code board collapses the paste box to one muted line; the hunter board shows the paste box holding the keyboard (gold edge).
- The druid board's Up next follows the data: 15 points on plan, Feral Charge at 25 (wow-player 1).
- The header shows the points waiting in the client gold, `1 to spend` (wow-player 2, ux polish).
- The followed build's name reads `Beast Mastery level 60 build` everywhere, as `Follow.buildName` says it in game; the boards had a stale `leveling` (wow-player 3).
- The off-plan talents show over the list without a hover: `Off plan` in the warning colour, a line a talent, three at most then `and N more` (wow-player 4, info-density 1).
- Up next says what is past its six rows: `26 more points after these` (wow-player 5, info-density 2).
- Section headings in the title face at 14 px, as the planner's panel titles (brand 1, ux polish).
- `Levels 44 to 48`, the plural as the planner's bands (brand 3).
- Shorter lines: `Read: <name>.`, `Every point is placed.`, the two transfer explanations tightened, the planner's apostrophe (ai-slop 5, brand 6).
- Placeholder names: the pusher is `Varnok`; every character and guild on a board is made up.

## Kept, with the ruling

- The two transfer explanations stay (one line each): the owner asked for import and export "with their own explanation".
- `your guild and realm are included` stays: it is the privacy line of the A.2 ruling, the one place a player learns the export carries them.
- The Replaces line stays in the warning colour: A.2, "the only warning before a plan changes".
- The split stays named by tree, biggest first (spec 2026-11-13 builds pane, section 1.1); the planner's rail shows the trees in its own header row, the addon's rows have none.
- The Follow buttons stay the client's blue (`Widgets.clientButton`, builds-pane spec and the dialogs lane's button factory).

## Follow-ups for later lanes (outside the Builds tab's job or needing new data)

- A talent tooltip on the Up next rows (needs the client's talent index per cell).
- Each row's distance from the character's spend ("12 points change") and per-talent diffs in the preview.
- Item icons and sources on the preview's gear rows; an expandable `and N more` for the site rows.

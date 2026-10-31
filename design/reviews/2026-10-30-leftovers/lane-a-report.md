# Lane A report (2026-10-30)

1. /sim strip: now the first row of the Your characters card on the landing (SettingsRow-style row: tile or crest, doors and Forget as buttons). On a loaded sim and the tool pages it is a card of its own. New STRIP_HEIGHT (h-[150px] md:h-[58px], measured 148-150 / 56-58) replaces CHIP_HEIGHT for the strip's slot; sim.astro, specs, [id], gear, drops, talents, weights min-h rose +62 phone / +14 desktop. Test sim-landing updated (strip inside the card).
2. Planner: the "gear not read" card moved out of the character column to a full-width row under the head; the 200 px hole is closed. Trees reserve unchanged.
3. Weights "REQUEST": it is the raw-request JSON drawer (RequestDrawer). Summary now reads "Edit the raw request" plus the hint "The JSON this run sends to the engine". Test ids unchanged.
4. Guides: Content's measure-aside centring removed on the guide page (global style in [spec].astro); prose now starts at x=96 like the head, TOC rail at the column's right edge.

Captures: a-<route>-1440/390.png (sim, sim-weights, planner, guide-warrior-fury).
Note: lane B's or another sync swapped public/data to real data mid-run; I re-synced fixture and reran.

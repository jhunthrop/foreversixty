# Raid-ready preset on the web (2026-10-07)

Branch `raid-preset-web`. Web only; builds against the contract (`preset` on every band entry, a top-level `presets.raid` block) and keeps today's files working (no `preset` means bare, no `presets` block means no label).

## What changed

- `web/src/lib/bis/presets.ts`: the one reader. `selectBand(file, band, faction, preset?)` returns the raid entry when the file has one and no preset is named, bare otherwise; an explicit preset the file lacks falls back to bare. Also `bandPresets` (what one band offers), `filePresets` (what a file offers), `presetMeta`, `parsePresetId`, `presetLabelFor`. Node-free, so `load.ts` `bandEntry` and the planner island's `hover.ts` `bandEntryFor` both go through it.
- BiS page (`/bis/<class>/<spec>`): a level-60 panel that offers both presets gets a "Conditions: Raid-ready | Bare" segmented control, a one-line caption built from `presets.raid` (label plus counts), and a "What is included" disclosure with the buff, debuff and consumable labels and the notes. Both variants are rendered at build time and swapped with CSS on `.page-body[data-preset]`, so weights, set DPS, the list, slot deltas, where-to-get-it and the stat-weight table all follow. `?preset=bare` or `?preset=raid` is read on load and written (replaceState, no reload) on click; the hash is kept. Below 60 there is no control. The panel body moved to `BandPanelBody.astro` so the two variants share one template; the bare copy's test ids carry a `-bare` suffix.
- Like-with-like deltas: the "set DPS up vs previous band" figure is only computed when the previous band is the same preset, so raid 60 is never compared with bare 50.
- Other readers, all through the shared reader (so they show raid when present): guides' stat-weight tables and rail card (`band60Weights`, the "Updated" line now ends with the preset label), class landing set DPS, BiS index cards (preset label as tooltip), home hero sim card and planner band compare. The planner talents/gear hover read the default entry.
- Fixture: `web/src/data/fixtures/bis/warrior-protection.json` gained raid entries for both factions and a `presets.raid` block (invented numbers) so the e2e is hermetic.

## Tests

- Unit: `presets.test.ts` (legacy file, file with raid entries, listings, labels, caption), two cases in `panel-view.test.ts`.
- E2E: `tests/e2e/bis-preset.spec.ts` (default raid, caption counts, disclosure, click writes `?preset=bare` with no reload, `?preset=bare` deep link incl. weights and list, no control at level 20). Desktop and mobile projects pass; `bis.spec.ts` and `guides-redesign.spec.ts` still pass.
- `format:check`, `check`, full `npm test` clean.

## Screenshots

`design/reviews/raid-preset-web/bis-preset-{raid,bare}-{1440,2000}.png` (captured by the e2e with `PRESET_SHOTS_DIR` set; they open scrolled to the band panel, so the control and caption sit at the top).

## Open questions

1. The home hero sim card compares the user's own sim with the band's set DPS. A user sim is a bare-ish run, so against the raid number it reads low; the preset label is only a tooltip there. Should that card compare against bare, or should the user's sim gain a raid option?
2. "Open in simulator" still links `/sim?spec=&band=60` with no preset; the sim has no raid-preset parameter yet.
3. The spec-tab and faction links on the BiS header do not carry `?preset=`; switching spec resets to the default (raid).
4. The control and caption are full-bleed with the existing 48px gutters; the page's panels were not max-width containers before and still are not (2000px shows them stretched). Left alone as out of scope.
5. Caption counts come from the array lengths; if the data lane wants a hand-written one-liner, `notes` is the natural field.

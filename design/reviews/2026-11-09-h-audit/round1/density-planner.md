# Density review: talent planner (round 1)

Job: pick and order talent points for a class. Captures: first/planner, planner-you, planner-guide at 1440 and 390.
Reachable data: public/data/<build>/talents/<class>.json (per-rank spell_id + description, tier, column, prereq), trees art, /data/leveling-plans.json (ladder order), guide builds (planner.astro guides collection: build, raidBuild), /v1/me characters, class and race tables.

## Verdict: THIN

The main column ends at the Point order plate while the rail runs 300 to 480 px longer at 1440 (planner-you, planner-guide: empty ground under Point order). The default capture (0 of 51) is two big plates of blanks: a Point order grid of 51 dashed squares and trees at 0/n. Rail cards (In the game, Share, Import) are each a heading plus one or two controls.

## MUST
1. Dead space under Point order (1440, left column, all three captures). Fill with a per-spec "Build vs guide" table: one row per guide build for this class (Beast Mastery 31/20/0 etc.), columns points differing from the loaded build, first differing talent, Load button. Source: guideBuilds in planner.astro + decoded build codes (lib/planner/guide-builds.ts); replaces the rail Guide builds list (which only links). Files: components/planner/GuideBuilds.svelte, Planner.svelte.
2. Trees carry no tier gate or spec context (all captures). Add the points-to-unlock gutter (5/10/15...) per row and a "n more to tier k" line in each tree footer. Source: talent.tier and column in talents/<class>.json. Files: TreeGrid.svelte, TalentCell.svelte.
3. Empty default state is blank (planner-1440). Show the class default: pre-select the spec guide build faded, or list the three guide builds with spec split and first 10 points as icons in Point order. Source: guideBuilds + leveling-plans.json order. Files: PlannerOrder.svelte, Planner.svelte.
4. Point order cells have no level and no name (all). Add the exact level on hover and a level stamp on the first cell of each row; tooltip names the talent and rank text. Source: leveling-plans.json order, talents ranks[].description. File: PlannerOrder.svelte.
5. Rail "Level 60 build" says only Matches or Covers 21 of 25 (you, guide). List the differing talents inline (name, you n, guide m) with click to flash the cell. Source: BandCompare diff already computed (Planner.svelte onband). File: BandCompare.svelte.
6. Rail cards are thin (In the game, Share, Import): merge into one "Take it with you" plate (Copy, Send to Felsnap, Share with title, Import field as a row). Frees the rail for MUST 1. Files: InGame.svelte, SharePanel.svelte, ImportBox.svelte.

## Should (ranked)
- Head: show spec split as 5/13/33 beside "51 of 51 points" and the build's guide name; the default head reads "Level 9" with 0 of 51, mixed signal. Source: store tree totals. File: SummaryBar.svelte.
- Character mode: "Talents loaded from a character." line is a grey explainer; replace with the character crest and Reset. File: Planner.svelte.
- Guide builds: mark which row the loaded build equals (only "Loaded" pill exists on one); add points-differing count per row.
- Class tab row: spec icons under the selected class with guide-build counts; race select only matters for racials, show nothing else.
- Share toggle "Include a simmed DPS on the card" lives on a talents-only page; fine as a toggle but never show the number here (sim owns it); link "Sim this build" is correct.

## Phone (390)
Order today: head, class, trees (spec tabs), Level 60 build, In the game, Guide builds, Point order, Share, Import. Move Point order directly under trees; fold In the game, Share, Import into one "Take it with you" plate with the guide table as a drawer. Point order wraps at 10 columns; keep 10 per row with the level stamp.

## Needs data
- Respec cost by level (gold): no table in repo.
- Talent pick rates among players: none.
- Per-talent sim value: sim owns it; a link only, no number here.

# h-audit: findings handed to other lanes

These sit in files another in-flight lane owns tonight, so h-audit lists them instead of fixing them.

## h-report-tail (score/RatingsPage)
- /ratings (ai-slop FAIL, ux NOT SHIPPABLE):
  - the empty 48 px icon gutter in the Forever Score rows, the weights table and the component list;
  - at 390 the weights print "35 / 30 / 10" with no role labels, and each score row shows its number twice ("412", then "412 of 500") with no bar;
  - the flat fake "Last 6 fights" sparkline, and the "Sample · Forever Score" / "The sample" labels;
  - a grey explainer under every part, and the lines "There is no leaderboard of ratings", "Topping the meter while dying…" and "The tiers take the colours of item quality";
  - six identical stacked plates;
  - tabs that are really in-page anchors;
  - about 4,500 px of page at 390, with the Tiers legend last;
  - "2100" against "2,100", and the subtitle repeating the tabs.

## f-selected
- Rankings.svelte and rank-list.css:
  - the Characters/Guilds tabs collapse to "Characters Change" at 390;
  - the stranded "Comparing specs, not characters? Tier list" card;
  - the rankings-board dead space at 390;
  - the Report column missing on rankings-you;
  - Date on phone rows.
- segmented.css: the BiS spec control band wraps "All sources" onto a second row at 1440.
- GuideSpecRow.astro: the "Your spec" / "Yours" pill wraps at 390.
- ui/PointOrder.svelte: the yellow dash marker before "Felsnap is level 34…".
- ui/TalentEditor.svelte: "Client build 1.60.1.70790" in the talent plate footer.

- ui/Toc.svelte: the guide spec page's 390 sticky bar reads "On this page" + "Change"; label it "Jump to".

## f-content
- ui/SpellIcon.svelte: a missing spell icon shows the browser's broken-image glyph on the guide spec page (Lightning Reflexes, Resourcefulness, Strider Kick, Aspect of the Beast, Rapid Fire, tree nodes); swap to the bordered placeholder on error.
- planner/TalentCell.svelte (via ui/SpellIcon): talent-tree nodes on the guide spec page show the broken-image glyph on lit nodes and a grey glyph on dimmed ones.
- content/guides/hunter/survival.md: the opener "Survival levels in melee."; the race prose "Humans can be Hunters for the first time in Forever" and "Skyborne can be Hunters on either faction:" restate the rail; "Enchant Weapon - Agility" listed twice.
- Guide spec page (round 4):
  - The footer line "1 point of agility = 0.63 DPS at level 60" repeats the Level 60 label and the DPS column (lib/guides/copy.ts).
  - The weights header says "DPS per point or rating", but "or rating" qualifies nothing.
  - The Lone Wolf key-talent blurb restates the Overview prose (GuideBuildsLane).
  - The survival.md blurbs "You fire no shots in this spec" and "Plain attack power gives only itself".
  - [spec].astro keeps its own `.guide-page` 40 px bottom padding on top of the body's footer gap. Its mock keeps a `.g-page` wrapper to match; drop both together.
- GuideRacesLane: "Recommended" pills beside coloured race names on Night Elf, Orc and Troll.
- The guide spec footer "Numbers from client build 1.60.1.70790, checked against launch logs…" is site meta talk.
- planner/TreeGrid.svelte and TalentCell.svelte: a tier-gate gutter and a "n more to tier k" line.
- SimTimeline.svelte: the timeline "Off / On" toggle has no label.
- lib/guides/copy.ts:
  - the "Your guide:" prefix (signedInCallout);
  - "DPS · Updated" in the spec band (headerMeta).
- content/guides/hunter/survival.md: the lede "Survival is a melee spec in Forever: …" (tell 23).
- GuideRacesLane and race-table.ts: racial text pasted verbatim from the tooltip.
- GuideBisLane, GuideExtrasLane, GuideBuildsLane and [spec].astro:
  - dead plate beside the prose;
  - Enchants and consumables rows with no icon;
  - Best in slot repeats the BiS page's 17-slot list;
  - Load and Sim are repeated as text links.

## f-account (footer gap in the mocks)
- Main's e4f6f961d and de47037db added a page-to-footer gap: global.css `[data-page-body]` is 40 px, and 64 px from 1024. design/mocks/ws_lib.py `_BODY_CSS` doesn't carry it, so every parity spec's site-footer box is ~64 px off on main. f-account's ws_lib.py change fixes it. Until that lands, this lane's parity is green only with that change applied (it was applied uncommitted when this lane was verified).

## f-content (sim-phone)
- web/tests/e2e/sim-phone.spec.ts reads the This run section. h-audit keeps that section unfolded on phones so the spec stays green. If This run should fold too, the spec needs `openSimSection(page, 'settings')` first.

## h-doll (later)
- ui/Paperdoll: item level and enchant state at 1440 on the character doll, and the floating Upgrades card.
- SimGearSection and CharacterStrip: the sim gear stat strip and the BiS DPS target on the sim landing; "Item level 65" repeated on every gear row (use the bare number or "ilvl"); a full-gear (16 slot) capture of Top Gear.

## API (h-api-html or the coordinator)
- The feed line "(Fury Warrior, 60 raid)" on home is generated by the server.

## Needs data (no lane)
- The spec median and rank per boss on the character payload.
- Item level and enchant/gem flags per slot.
- Per-ability DPS for the BiS benchmark.
- Respec cost by level, and talent pick rates.

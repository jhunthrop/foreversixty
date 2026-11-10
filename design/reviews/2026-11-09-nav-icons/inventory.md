# UI icon inventory, 2026-11-09

Every UI glyph web/src drew before this lane, what it means, where it came from, and where it
comes from now. One source after: `web/src/lib/ui/ui-icons.ts` (generated from
`design/assets/icons/ui/*.svg`, the owner's custom set, and `design/assets/icons/ui/line/*.svg`,
the line family: 24 px grid, 2 px round stroke), drawn by `ui/Icon.svelte`, `ui/Icon.astro` or
`lib/ui/icon.ts`. `src/lib/ui/icon-source.test.ts` fails on any other raw svg.

The rule between the two families: the custom (filled) set marks a place or an object (doors,
panel rows, account tiles, settings rows) and is drawn at 20 px and up; the line family marks a
control or a state (chevrons, close, check, menu, spinner, refresh, chart markers).

No icon library (lucide, iconify, heroicons, phosphor) was in use.

## Nav

| File | Glyph | Before | After |
| --- | --- | --- | --- |
| Header.astro | door icons: Best in slot, Guides, Tier list, Sim, Planner, Logs, Guild, Premium | none (text only) | custom bis, guides, tier-list, sim, planner, logs, guild, premium |
| Header.astro | door chevron (panel toggle) | inline svg, 2.5 stroke | line chevron-down |
| Header.astro | phone Menu button | inline svg | line menu |
| HeaderPanel.astro + nav-row-icons.ts | Upload a log | hand-drawn path table | custom upload |
| same | Your reports | hand-drawn | custom logs |
| same | Rankings | hand-drawn | custom rankings |
| same | Your guild | hand-drawn | custom guild |
| same | Roster (Find a guild for a guildless visitor) | hand-drawn | custom character (two heads stays the account menu's Account) |
| same | Loot council | hand-drawn | custom bis |
| same | Guild logs | hand-drawn | custom logs |
| same | Planner, BiS, Simulator, Get set up, Tier list rows | hand-drawn, no row used them | deleted (dead) |
| HeaderSearch.astro | search button, bar panel field, Menu field | 3 inline magnifiers | custom search |
| nav/CharacterSelector.svelte | account chevron (x5), crest badge chevron | inline svg, 2.4 and 4 stroke | line chevron-down |
| nav/CharacterSelector.svelte | signed-out ghost (x3) | inline person svg | custom account |
| nav/SelectorMenuRow.svelte | chevron, ghost | inline svg | line chevron-down, custom account |
| nav/SelectorPanel.svelte | account menu tiles | custom set (name lookup) | custom set (per-icon import) |
| nav/SelectorPanel.svelte | row chevrons | Glyph chevron | line chevron-right |
| nav/SelectorPanel.svelte | sheet close | inline svg | line close |
| nav/SelectorRow.svelte | current-character check | Glyph check | line check |

## Where the door icons show

Panel rows, the account menu and every phone Menu row carry their icons at every width. The bar's
doors carry theirs only where the bar is at least 1640 px wide (the wide column from about a
1930 px viewport): measured signed in with an 18-letter name, the icons need about 220 px and the
1440 bar has 24 px free. The owner's rule keeps the labels whole, so below that line the doors
are labels alone. This is the measured limit, not a regression.

## Actions

| File | Glyph | Before | After |
| --- | --- | --- | --- |
| account/discord/DiscordActions.svelte | send test | Glyph send | line send |
| account/discord/DiscordReplace.svelte, sim/ActionMark.svelte | refresh / refund | Glyph refresh | line refresh |
| ui/UiSheet.svelte (design sheet) | sample button glyph | Glyph card | none (no meaning to carry) |
| ui/UiSheet.svelte | sample settings rows | Glyph server, chart, star, bell, lock | custom sim, rankings, premium, link, logs, manage |
| coach/CoachReviewBody.svelte | bullet arrow | inline 12 px svg, 1.5 stroke | line arrow-right |

## Status

| File | Glyph | Before | After |
| --- | --- | --- | --- |
| ui/ItemIcon.svelte | upgrade badge, worn check | Glyph upgrade (filled), check | line upgrade (filled), check |
| checkout/GuildStepCard.svelte | step done | inline 14 px svg | line check |
| ui/Spinner.svelte | busy ring | inline svg, 3 stroke | line spinner |
| chart/TimeSeriesPlot.svelte, ChartTooltip.svelte | death, target change, cast marker | Glyph skull, target, bolt | line skull, target, bolt |
| chart/ChartLegend.svelte, LegendSheet.svelte | pinned line | Glyph pin | line pin |

## Chevrons and carets

| File | Before | After |
| --- | --- | --- |
| ui/ActionLink.svelte, ui/SettingsRow.svelte, account/discord/DiscordConnect + DiscordPasteForm | Glyph chevron | line chevron-right |
| ui/Disclosure.astro (an Astro file importing the Svelte Glyph) | Glyph chevron | line chevron-right via Icon.astro |
| ui/ControlBar.svelte | Glyph chevron rotated 90 | line chevron-down |
| ui/LevelSlider.svelte, character/ChartReadout.svelte | Glyph chevron / inline 16 px svg (1.75 stroke) | line chevron-right (rotated for back) |
| chart/TimeSeriesPlot.svelte | Glyph chevron | line chevron-right |
| ui/Select.svelte, rankings/RankingsMoreFilters.svelte, bis/BisSources.svelte | inline svg, 2.5 stroke | line chevron-down |
| guild/RaidDisclosure.svelte | inline 16 px grid svg, 1.5 stroke | line chevron-down |
| ui/TabStripMore.svelte | text caret "▾" | line chevron-down |

## Close and search

| File | Before | After |
| --- | --- | --- |
| report/ThreatChart.svelte, sim/TimelineReadout.svelte | Glyph close | line close |
| character/FinishedSimLine.svelte | inline 16 px grid svg, 1.75 stroke | line close |
| lib/ui/listbox-render.ts | glyphs.ts CHECK, CLOSE, SEARCH strings | line check, close; custom search |
| CharacterSearchForm.astro | search field clear button: svg data URI in CSS, 2.5 stroke | line close as a CSS mask (lib/ui/icon.ts iconMaskUrl) |

## Removed as dead

- `lib/ui/glyphs.ts` + `ui/Glyph.svelte`: the old line table. Its glyphs with no caller left
  with it: download, share, layers, trophy, news, enchant, play, link, alert, spark, gear, copy,
  clock, crown, hash, external, device, signout, the shield, user, discord, scale, list, sword,
  tree, book, upload. The ones in use moved to the line family; the design sheet's samples moved
  to custom glyphs.
- `bis/GlyphSheet.astro`, `bis/SourceKindIcon.astro`, `lib/bis/glyphs.ts`: the BiS slot and
  source-kind sprite. No page referenced its symbols any more; the guide page still mounted the
  empty sheet, which is gone too.
- `components/nav-row-icons.ts`.

## Allowed raw svg (not UI glyphs)

Charts: ui/Spark, character/ChartPlot, chart/OverviewBrush, report/ResourceGraphs,
sim/SimTimeline, sim/TimelineBrush, lib/sim/timeline-draw.ts, planner/TreeGrid and ui/TalentEditor
(talent connectors). Art: ui/FrameCorners (filigree), character/ClassCrestRing and
character/CharacterPortrait (the class-less crest silhouette). Brand: DiscordMark.astro. Game
icons (item, spell, class crests) are images and out of scope.

## Left as text, not icons

- The `›` after inline text links (CheckoutLauncher, Upload, Notice, Section, Checklist, FlatSheet,
  BisSide, BisSpecPage, CharacterGear, InGame, and the home lane's HomeFeed, HomeGuildSims,
  HomeCharacter, HomeRaidPlaque, HomeUpgrades) and `←` in GuildShell: typography in the link's own
  text, not a glyph slot. ActionLink (chevron icon) is the primitive for a link that wants one;
  moving them is a copy/style change for each page's lane.
- Steps.svelte and lib/ui/cells.ts status marks (`✓`, `!`, `×` set in a coloured disc): text marks
  in the font, same follow-up.

## Other lanes' files

No icon sits in PageHeader.svelte or the page head bands. Home lane files with `›` text links:
HomeFeed, HomeGuildSims, HomeCharacter, HomeRaidPlaque, HomeUpgrades (left untouched).

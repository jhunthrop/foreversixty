# L3 Gear tab: brand-consistency-reviewer

Method: read the renders (gear-hunter-60.png, gear-warrior-34.png) against web/src/lib/bis (source-filter.ts, source-choice.ts, page-state.ts, copy.ts, scale-view.ts), design/tokens/tokens.json, addon Tokens.lua/Theme.lua, Sources.lua, GearTabView.lua and Locale.lua. The live page was not screenshotted; the web side is read from code.

## Round 1: PASS

### Checked and matching the web
- Source filter kinds, order, words: Dungeon, Raid, World / BoE, Quest, Crafted, Reputation, Vendor, PvP, exactly `SOURCE_FILTER_KINDS` and `sourceKindLabel`. Locale keys sourcesKind* carry the same words. World drops and BoE fold into one kind (B and W), as the web's FILTER_KIND_OF_SOURCE does.
- Default: below 60 every kind but Reputation (standings up to Friendly kept), at 60 every kind (`Sources.defaultOn`, `Sources.REP_CAP`, `isLevelingDefault` mirror `defaultSources`/`isLevelingDefault`). The field text matches `sourcesFieldLabel`: "All sources" at 60, "No rep above Friendly" at band 30 (warrior render). Counts are the web's "slots with a ranked option from that source"; a kind with none is dimmed (Raid and PvP at 0 verified in the render's pixels).
- Re-pick, pair rule (second slot steers clear of the first, no off hand beside a two-hander) and the pre-raid kind set are twinned from source-filter.ts and pre-raid.ts.
- Wording: "Best in slot" (sentence case, never "BiS" in the list), "Upgrades", "Stat weights", "Level", gains as "+4.9 DPS" (one decimal, unit, as `oneDecimal` and togetherDpsValue). Header "9 of 16 best in slot" is the web's term.
- Colour: no hex literal in GearTab.lua, GearTabView.lua or Sources.lua; gains use `success` (= outcome.kill 7fd48a), "No enchant" uses `warning` (= outcome.wipe f0736b), item names take the quality colour, icons carry a quality edge, muted/body from tokens. Tokens.lua is the generated file. Ground is the web's brown-black raised plate; hairline rule between columns; flat, no gradients or glows.
- Casing: sentence case throughout, tab underline in gold as the site's active state.

### Must
None.

### Should
1. Upgrade source lines truncate the information the player came for: "Bouldefist Enforc... · World · 1% drop", "Scarlet Mon... · Interrogator Vishas", "Razorfen Kr... · Charlga Razorfla...". The web's source line (source-line.ts) wraps rather than elides. Give the slot and instance a fixed share and let the boss and the drop chance keep theirs, or drop the slot word (the doll and the icon already say it), or wrap to a second line; the full line stays in the hover.
2. "Beats the pick" is not a web word. The web says "Best in slot" for the list's top item and shows the worn item as "equipped" (`equippedTickLabel`) with no jargon for "your gear out-scores the list". Reword with the web's vocabulary, for example "Better than the list" or simply "Best in slot" with the worn item shown, and add the key to the shared copy so the Hub, the tooltip and the web say one thing. (Spec 4.3 may own the phrase; if so, add it to the web's copy too.)
3. Weights note: the web states the rule (`weightsScaleNote`: "Per point of stat, normalized to Hit = 1.00, with DPS per point and the sim error"); the addon says only "shipped, per point" and drops the error figure the spec lists ("error"). Either show the error column (the web does, and greys "Not significant" rows via `weightsNotSignificant`) or put the normalisation in the section's hover so a regular can read why Hit is 1.00. Also use the web's "Not significant" and "No gain" words for rows the data marks so, if the addon shows any.
4. Reputation rule line: web `defaultRuleLine` is "Reputation above Friendly is off below level 60, so those slots are re-picked without it. Turn it on in Sources to include it." The addon's "Off: Reputation above Friendly" is a shortened form. Acceptable in the narrow column; put the full web sentence in its hover so the words stay one source.

### Nice
5. Level control: the web's group is "Level range" with the band label ("20 to 29"); the addon shows "Level 20 30 40 50 60". Fine for the window, but the band hover could say "20 to 29" via `bandRangeLabel`'s words. The unselected band buttons are the client's blue button, not the site's plate; the selected one is teal. Per the addon idiom this is allowed, but the unselected state reads louder than the web's quiet pills.
6. Upgrades total: "+17.1 DPS in all" versus the web's "Together +N DPS" (`togetherLabel`). Two phrases for one figure; the addon's reads better at the heading, but name it once in the shared copy.

Everything above is polish; nothing makes the tab read as a different product. The hunter and warrior renders hold the web BiS page's order of emphasis: the doll first, then upgrades, weights, sources.

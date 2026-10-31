# Wave 3 mobile design: wow-player verdict

Reviewer: the wow-player persona (`.claude/agents/wow-player.md`), applied by the ux-designer on the rendered boards in `design/mocks/2026-10-31-mobile-wave3/renders/`, 2026-10-31. Classic since 2019, every class to 60, Naxx, loot council; lives in Details!, WeakAuras, ElvUI, Wowhead and Sixty Upgrades. I read the boards as a player, not the diff. I judge each against what I already trust. Everything on the boards that is a figure, a source or an order is the build's own (1.60.1.70658) or marked invented in the spec's Mock section.

The site's job still applies: help me play my class better, not a reference site. No board opens on a slogan or a feature pitch. Checked: Guides opens on `Your guide`, rankings on `Your rank`, the item page on `Your Back pick`, the widgets on my raid and my next upgrade.

## F9 Guides, tier lists, what changed
1. **Question.** "How do I play Elemental, and where does it stand?" The answer is first: role, then the seven stats as numbered chips, then the overview.
2. **Wrong or amateur.** Nothing I would screenshot to mock. The rotation lines have the icon, the rank and the condition in words, the way a WeakAura names a spell. The tier list leads with my own spec, not with the top of the list.
3. **Missing.** A guide has no talent point order for levelling (the planner has it). Fine: one job per page.
4. **Data I do not believe.** **Elemental Shaman is 20th of 20 DPS specs, 51% behind Fury Warrior, tier D.** No Classic player thinks that. It is what `bis/shaman-elemental.json` says today for the 60 raid band, so the number is the data's, not the layout's. I will not call the board wrong, but I will not sign the tier list as fact. **Gate for lane F9: the accuracy reviewer confirms the Elemental band and the tier list's spread before `tiers.json` ships (tenet 8).** The same goes for Frost Mage (44% behind) and Balance Druid (48%).
5. The `≈` tie mark and the tier letters in quality colours read like the site's own tier list. Good.

## F10 Your reports
1. **Question.** "How did I do last raid?" `Your parses` is the first card, with the DPS and the parse colour. Good.
2. **Wrong.** Nothing. Trash is one line per zone (`Trash · Onyxia's Lair`, `9 pulls · 14:20`), as Warcraft Logs does it. The meters are Details!-shaped with my row marked and the unit beside every number (this was missing in the first pass and is fixed).
3. **Missing.** Damage taken and deaths are on the web only. I accept that: a phone is not where I read a death log. The meter has no `%` of total; Details! has it. Not needed on 5 inches.
4. **Believe.** A 3:12 Onyxia kill with a 78 parse is plausible. Coach text is marked as an invented example.
5. Free players see no coach card and no prompt. Good: no nag.

## F11 Rankings, character page
1. **Question.** "Where do I stand?" `14th of 212 Elemental Shaman characters` is first, in words I recognise.
2. **Wrong.** Nothing. The pinned row at 204th and the in-place gold edge are what I expect. Parse colours follow the Warcraft Logs ladder (grey, green, blue, purple, orange, pink), which is the culture's own. The Paperdoll looks like the character frame with names beside the sockets and a gold dot where my list has a better item.
3. **Missing.** The character page does not show talents. The `Builds seen` row with `Open in planner` covers it.
4. **Believe.** The Forever Score is the site's own number and the card links to the explanation.
5. `No logs yet.` instead of a zero is right.

## F12 Search and the item page
1. **Question.** "Where does it drop, and does my spec want it?" The item page opens on the name in quality colour and `Your Back pick`, then the tooltip, then `Where it comes from`. That is Wowhead's page, shorter, with the answer for my spec on top.
2. **Wrong.** The Spritecaster Cape tooltip has no `Binds when picked up` or `Binds when equipped` line. A cloak from BRD must say which. The board omits it because the build file has no `bonding` for this item yet (contract C6, pending). **A lane capture without that line is a blocker: the F12 lane does not ship until the tooltip prints it** (spec §5.4, the improvements list in the contract).
3. **Missing.** The drop chance is not shown for the cape. The spec says a chance prints only when the data has one, never `0%`. I would rather see `Drop chance not known` than nothing, but a blank is honest, and Wowhead shows blanks too.
4. **Believe.** The two BRD bosses that drop the cape (Gerstahn, Grebmar) are right. Six specs listed with `Pick` for the cape is plausible for a caster cloak at 60.
5. Search by `devout` returns the five Devout pieces with slot, item level and required level. I would use it.

## F13 Planner
1. **Question.** "Let me build it and send it." The tree is the primary; the strip tells me what my last tap did. Tap adds, long press removes, plus and minus for precision, undo.
2. **Wrong.** Nothing I would mock. The refusal says why in words (`Needs 20 points in Elemental. You have 16.`); the first pass had a count that contradicted its sentence, which is fixed. The share step shows exactly what goes into a public link: the web's own text.
3. **Missing.** The class is shown by its crest only. I know a shaman crest; a new player might not. Polish, left to the lane.
4. **Believe.** Elemental Fury's rank text (`by 60%` at 3 of 5) and Lava Burst's text (`114 to 145 Fire damage`, `20% increased damage`) are the client's own words from the talents file.

## F14 Widgets and live surfaces
1. **Question.** "Am I in tonight?" The small widget says the raid, the time, the countdown and my answer. The medium one lets me change it without opening the app. This is the feature I would install the app for.
2. **Wrong.** Nothing. Locked and cancelled raids drop the buttons. `Updated 14 h ago` says when the data is old instead of lying.
3. **Missing.** The widget shows my phone's time only. My guild's server time is on the raid page. I accept it: a raid leader says "8 pm server" and I can see both on the raid page one tap away.
4. The sim Live Activity exists only for a server run, so it cannot nag a free player. Right.

## F15 System integration
1. **Question.** "Do it by voice, and work with no signal." Siri confirms once and answers with the count. The next-upgrade answer has the item icon, its colour, the source and the gain.
2. **Wrong.** Nothing. I like that Android does not pretend to have voice sign-up.
3. **Missing.** Nothing a player asked for.
4. Offline data lists sizes and no build ids or hashes. Right.

## F16 Guild extras
1. **Question.** "Who is in my guild, and who turns up?" Counts, role chips, spec icon, rank and `6 of 8`. That is the roster I check before inviting.
2. **Wrong.** Nothing. The history says `You were there`.
3. **Missing.** Per-night attendance dots would help an officer, but the API gives only a count, and the screen does not pretend otherwise.
4. The calendar sync and the link are what my raid leader's spouse will use.

## Things that must be true in the build (not design defects)
- The Elemental tier band confirmed by the accuracy reviewer before the tier list ships.
- A `Binds` line on every item tooltip the app shows, from the client's ItemSparse rows.
- No board shows a figure I cannot verify as fact: parses, guild names, people and sim DPS are invented and marked so.

## Verdict

**SHIP.** The surfaces answer the question I would have, in the words and colours I know, with the unit beside every number. Two gates travel with the lanes: the tier list's data (lane F9) and the tooltip's Binds line (lane F12). Neither is a layout defect.

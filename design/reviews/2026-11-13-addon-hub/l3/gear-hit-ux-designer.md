# Gear tab, Conditions row and hit lines: ux-designer

Reviewed: l3/gear-hunter-60-raid.png and l3/gear-warrior-60-bare.png against hub-gear.png, same scale. Spec 3.5, 4.3, 13 (Hit wording, Hit cap).

## Verdict: SHIP

No musts. The earlier l3 acceptance of "no Conditions row" is now closed: the row exists, only where a Raid-ready list exists.

## Checked, region by region

- Conditions row: sits directly under Level, same label column (x 90) and same first-button edge (x 230) as Level. Client buttons only; the chosen one is teal (Raid-ready on the hunter, Bare on the warrior), the other blue. Matches the spec's control table (line 129) and the hub-gear.png recipe.
- Header, doll, Upgrades, Sources counts follow the chosen list: hunter 8 of 16 and +21.2 DPS in all on Raid-ready; warrior 17 of 17, Upgrades empty line "Nothing on this list beats what you wear.", source counts differ from the Raid-ready board. Consistent.
- Hit lines: muted, under the weights, same pitch and position as the approved Haste line in hub-gear.png, so one pattern for "a note under the weights". Hunter shows Hit and Spell hit; dual-wield warrior shows Hit and White swings. Warrior's 3% / 22% pair is internally consistent (6% in hand against 9% and 28%).
- Hovers: Raid-ready tooltip (Buffs, Debuffs, Consumables, grey group heads), Bare tooltip, Hit cap tooltip all use the client tooltip recipe. The Hit cap tooltip quotes the client's own line, as section 13 requires.

## Shoulds

1. The hit lines have no shown hover state. They are muted text with a tooltip, so nothing says they answer. Draw the hover: the line lifts to body colour while the tooltip is up. Same for the Conditions buttons' hover (blue lift), which the boards do not show.
2. "2% to the 9% cap" can be read as "you have 2% of 9%". The spec's wording is "N% to cap", so it ships. If the tooltip has room, open it with "Hit still needed:" or add "more" in the line ("2% more to the 9% cap"). Owner's call; wow-player to weigh in.
3. The hunter board has no Haste line under the weights, where hub-gear.png has one. If the shipped weights carry no haste figure, say so in the notes; if they do, the line is missing.
4. The Raid-ready tooltip is 640 px tall at this scale: three paragraph blocks of comma lists. Acceptable as the web's own content, but break it into one line per buff only if the client wrap makes it taller than the window; check in game.
5. The hunter's Spell hit line (9% to the 16% cap) is for wow-player to confirm against the hunter's spell-table specials (Arcane Shot, Serpent Sting); layout is fine either way.

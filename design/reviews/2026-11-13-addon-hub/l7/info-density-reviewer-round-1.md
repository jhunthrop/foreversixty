# L7 Officer tab: information density, round 1

Checked: the nine renders in `l7/`, spec 4.6, 8 and 9, `Officer.lua`, `OfficerLoot.lua`, `OfficerTab.lua`, the `ReadyCheck.lua` function list, and the old GuildTab officer region.

## Verdict: PASS

The raid-day tab is dense and uses most of what we have. Nothing blocks. Several data sources we already hold are still unused.

1. SHOULD: Readiness skips the attunement row in the render; show `Not attuned N` with names when above 0.
2. SHOULD: The Raid build row says "Dravik off" but not what each person follows; put it in the hover.
3. SHOULD: Buff coverage from `ReadyCheck` (coverageEntries) has no home; add a row or confirm `/fs ready` is its only home.
4. SHOULD: No visible sort state on the default heading.
5. SHOULD: The roster hover should carry the bare slots, best parse and the report's age.
6. SHOULD: The session's Wears is cut; hover with the item.
7. SHOULD: The right column could carry the Give time.
8. NICE: The night's line could name the top gainer.
9. NICE: Click a role to filter the roster to it.
10. NICE: Copy the roster; a whisper per row.

Needs data: durability and repair state; the guild-wide loot ledger (`loot_awards`); parses against the spec median.

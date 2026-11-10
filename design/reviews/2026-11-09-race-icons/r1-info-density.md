Verdict: PASS

Page jobs: the class landing says who can play the class and where its specs go; the spec Races lane says which race to pick for this spec and what each racial does.

Reachable and used: races.json, combos.json and raceicons.json drive the landing list (faction order, faction colour, client portrait). The lane keeps its racials, tooltips, Recommended tag and prose. The change adds the missing portrait and no invented data.

Checked captures (1440 and 390: hunter, mage, hunter-bm, warrior-fury):
- Landing rail: the list fills what was a prose-only panel. Alliance then Horde, 24 px portraits, faction colours. At 390 it folds to two columns and long names (High Order Skyborne) wrap cleanly. No dead space added.
- Lane rows: the 28 px portrait leads each race, with 20 px racial icons below. The size step reads as hierarchy. Tag, tooltip and wrap behaviour hold at 390.
- Missing-portrait fallback is name only, honest.

Optional additions (not required):
1. Landing list: mark the races the spec guides recommend (Orc, Night Elf for hunter), data from race-table.ts rows. Each spec differs, so only a link or per-spec marks would do; skip if it blurs the one-job rule.
2. Landing list: link each race to the race's row in a spec Races lane (anchor #races) so the list drills down.
3. Landing prose duplicates the list's names (hunter paragraph restates Dwarf, Night Elf, Orc...). The prose now only carries the Skyborne and new-in-Forever facts; trim it in the intro/prose lane, not here.

Needs data: none.

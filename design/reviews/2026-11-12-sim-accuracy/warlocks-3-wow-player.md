# wow-player review: Affliction warlock guide, card and coach (warlocks pass 3, 2026-11-12)

Surfaces: `web/src/content/guides/warlock/affliction.md` (the leveling build from 42, the Talents bullets, the leveling list), `data/curated/apl/warlock-affliction.json` (the Shadow Trance note), `data/curated/coach/warlock-affliction.md` (brought to the 11 Nov Agony rotation). Text read in the worktree, not rendered: the page renders the same components with new text, so no Playwright spec touches it.

## First pass: NO-SHIP

MUSTs, all taken:
1. The Improved Drains sentence contradicted the Drain Life and Drain Soul advice and talked like a patch note ("no longer reaches Wrack", "rotation"). Now: "Neither build takes Improved Drains: it does nothing for Wrack, and Wrack replaces Drain Life from 40."
2. The coach card's food level contradicted itself (forty-five in the consumables line, thirty-five in the preparation drill). Baked Salmon is a level 45 food (items.json 13935); the drill now says forty-five. The other 27 coach cards carry the same "food from thirty-five" drill line against their own food levels: open, for a coach lane.
3. The stray "Portal of Summoning:" label on the shards bullet (pre-existing). Now "Soul Shards:".

SHOULDs taken: "from 47 spends its points on Bane and Aftermath" (the last points are Unholy Power); the leveling list names every level from 22 to 60 (Grim Reach 28 and 29, Siphon Life 30, Soul Siphon 31 to 33, Soul Harvest 34, Shadow Mastery 35 to 39, Wrack 40, the second Soul Harvest 41, then the new stretch); Amplify Curse "whenever it is ready" instead of "before every Bane of Agony"; the coach's reviewed date. Not taken: the curated file's top-level `notes` and `engineering_notes` (internal history no page renders; the reviewer asked to confirm, confirmed).

Checked for the reviewer: Ruthless is the Forever Human racial (races.json: 1% critical strike chance with all spells and attacks); Improved Bane of Agony is 5/10% in the client's talent text.

## Second pass: SHIP on the text, no MUSTs

SHOULDs taken: the DoT paragraph's lapse timing (a channel burns 12 seconds of the DoT clock, so one with a little over 3 seconds left lapses partway through, not in the last seconds); Immolate in the levels 20 to 60 opener; "after Wrack" (Soul Harvest's second point sits at 41); the Shadow Trance line in the coach card and the card note says Improved Shadow Bolt is the raid build's.

Not taken (pre-existing text the client data confirms, for the record): Curse of the Elements rank 1 trains at 20 (440892 spell_level 20) and Demon Armor rank 1 at 20 (706), both Forever's own levels; Curse of Recklessness rank 4 is 505 armor (11717 effect 1); Demonic Sacrifice lasts 2 hours (the client's talent text); Portal of Summoning is Forever's name for the ritual (the class guide index); Forever gear carries spell haste; the Undead line in Races stands with Human and Orc as the picks. Left open for a later read: whether every demon but the Imp costs a shard in Forever (the shards bullet names the Succubus), the Life Tap note's phrasing, the Eureka! sentence's wording.

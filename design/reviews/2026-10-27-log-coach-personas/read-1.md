# Task 23 WoW player read: Fury, Holy, Protection personas (2026-10-27)

## warrior-fury: FIX (1 correction)
1. "the catalogue looks for Flask of the Titans or Flask of Supreme Power" -- Supreme Power is the +150 spell damage flask; no Fury warrior drinks it. The shared dps catalogue lists it (logs/engine/mechanics/consumables/catalogue.json roles.dps.flask), but the persona should name only Flask of the Titans for a melee spec (and say the coach accepts either).
Verified OK: Death Wish on cooldown (180s cd, data/curated/apl/warrior-fury.json), Bloodrage <80 rage, Execute first below 20%, HS at 40+ rage, Whirlwind both weapons / Unbridled Wrath any weapon / Bloodthirst 45% AP (web/src/content/guides/warrior/fury.md).

## priest-holy: FIX (2 corrections)
1. core_casts: "Inner Focus or Shadowfiend sitting ready for most of the fight" -- contradicts the file's own pacing line: Shadowfiend (5 min cd, spellconst/priest.json) is gated below half mana in data/curated/apl/priest-holy.json and is correctly held until then. Only Inner Focus (3 min cd) should be named; Shadowfiend should read "not used by the end of a fight that drained the bar".
2. uptime: "It lasts ten minutes, so it drops only after a death or a long fight." -- Classic Inner Fire is 20 charges spent when you are hit; spellconst shows duration 600000 ms but no charge data, and the APL never casts it. Do not state a drop-off rule the repo cannot back; say "recast when the buff is gone".
Verified OK: Inner Focus (180s cd), Shadowfiend and Dark Sacrifice (spell rows; Dark Sacrifice trainable id 1277324 in data/tests/test_trainables.py; APL gates below half / below 70% mana with health above 70%), Lightwell 600s cd prepull, Prayer of Mending rank 1 at level 40, Litany of Light, Spiritual Guidance, Spirit of Redemption prerequisite. No "heal more" outside the Voice rule line (grep clean).

## warrior-protection: FIX (1 correction)
1. uptime: "Thunder Clap and Demoralizing Shout falling off the boss" -- Demoralizing Shout is not in the rotation (data/curated/apl/warrior-protection.json notes; web/src/content/guides/warrior/protection.md) and drill.adherence says everything off the six-button list waits. Drop Demoralizing Shout from the finding (or from "earns its place on the bar") so the file does not tell a player to do what its own drill forbids.
Verified OK: Shield Wall below 40% before Last Stand below 35%; Shield Wall 15 min / 60% / 12s; Last Stand 30% max health, 20s; Shield Block 18%, Thunder Clap upkeep 16%, refresh in last 3s; Bloodrage <60; HS 40+ rage; Thunder Clap in Defensive Stance (classes.json); Devastate not learnable.

## All three
Nobody named: no real player, guild or ranking appears in any file.

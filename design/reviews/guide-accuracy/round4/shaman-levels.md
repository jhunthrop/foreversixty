# Shaman: Levels 1 to 19 (round 4)

Each spec guide's `## Leveling` now opens with `### Levels 1 to 19`; the existing text sits under a new `### Levels 20 to 60` so it does not render under the low-level heading. Where the old text already covered 1 to 19 it was moved into the new subsection, not repeated.

## Per spec

- **Elemental**: Lightning Bolt pull, Earth Shock 4, melee finish (no wand); Call of Earth 4 (Stoneskin Totem), Call of Fire 10 (Searing Totem); Flame Shock 10, Fire Nova 12 around the Fire totem; talents Convection 10 to 14, Concussion 15 to 19; leather or cloth with Intellect, quest rewards, quest staff or mace with spell damage. Removed "Searing Totem adds damage to every pull" from the later text.
- **Enhancement**: Rockbiter 1, Earth Shock 4, Lightning Shield 8; Rockbiter stays on until Windfury 30; totem quests; Searing and Strength of Earth from 10; talents Thundering Strikes 10 to 14, Ancestral Knowledge 15 to 19; first slow two-hander from a quest (staves default, two-handed maces and axes from a weapon master); leather to 40; never dual wield. Moved the Lightning Shield, imbue and 10-totem sentences and the talent opening out of the later text.
- **Restoration**: quests with Lightning Bolt, shocks, Searing Totem; no wand; totem quests; talents Improved Healing Wave 10 to 14, Mindfulness 15 and 16, Tidal Focus from 17; Healing Wave ranks 6, 12, 18; Tremor Totem 18 for Wailing Caverns sleeps; quest staff or mace with spell damage. Later text now starts at Water Shield 20 (Tidal Focus full 22, Mindfulness 23, both from the order).

## Data checks (1.60.1.70755)

- trainables/shaman.json: Lightning Bolt 1, Rockbiter 1, Earth Shock 4, Stoneskin 4, Earthbind 6, Healing Wave r2 6, Lightning Shield 8, Rockbiter r2 8, Flame Shock 10, Flametongue 10, Searing 10, Strength of Earth 10, Fire Nova (408341) 12, Healing Wave r3 12, Rockbiter r3 16, Flametongue r2 18, Healing Wave r4 18, Tremor 18, Windfury Weapon 30.
- addon-builds.json leveling orders mapped to talents/shaman.json (tier/column 0-based in file): elemental Convection x5, Concussion x5, Elemental Devastation at 20; enhancement Thundering Strikes x5, Ancestral Knowledge x5, Mental Dexterity from 20; restoration Improved Healing Wave x5, Mindfulness x2, Tidal Focus x3, Water Shield at 20, Tidal Focus 21 and 22, Mindfulness 23.
- Per-rank text from talents/shaman.json: Convection 2%, Concussion 1% (Lightning Bolt, Chain Lightning, Earth Shock), Thundering Strikes 1% crit, Ancestral Knowledge 2% Intellect, Improved Healing Wave 0.1 s, Tidal Focus 1%.
- Imbue claim from spelltips/shaman.json: Rockbiter r2 +58 AP (about 4.1 DPS) and r3 +88 AP (about 6.3 DPS) against Flametongue r1 about 3.3 and r2 about 4.8 fire DPS (per-hit damage scales with speed, so DPS is speed independent). Rockbiter leads at every level below 20.
- Fire Nova 408341: "Instantly inflicts 33 fire damage to enemies within 10 yd of your active Fire totem."
- No wand: no Forever source gives shamans wands; Classic rule kept.

## Sources

- Warcraft Tavern, Shaman Leveling Guide (Forever): https://www.warcrafttavern.com/forever/guides/shaman-leveling/ (Call of Earth/Fire/Water levels, abilities by level)
- Warcraft Tavern, Dwarf Shaman Totem Quests Guide (Forever): https://www.warcrafttavern.com/forever/guides/dwarf-shaman-totem-quests/ (Call of Earth 4 Stoneskin, Call of Fire 10 Searing, Call of Water 20 Healing Stream)
- Warcraft Tavern, Forever shaman guide: https://www.warcrafttavern.com/forever/guides/shaman/ (totems at 4/10/20/30, leather to 40, staves and maces default, other weapons from trainers, Tremor Totem for Druids of the Fang in Wailing Caverns)

## Open

- Call of Earth start level: the leveling guide's table lists 3 for some starts and 4 for others; the data teaches Stoneskin at 4, so the text says 4.

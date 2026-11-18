"""The client DB2 tables the pipeline reads, and where each lives in the client.

`TABLES` is what `pipeline fetch` writes into `builds/<build>/raw/<Table>.csv`;
`FILE_DATA_IDS` also covers the tables single steps read on their own
(`pipeline.guild_events`, `pipeline.gametables`).
"""

TABLES = [
    "Map",
    "AreaTable",
    "JournalInstance",
    # A raid's bosses and their journal loot (pipeline.raid_bosses, the guild loot
    # picker's first-tier table). DungeonEncounter names each boss and its map;
    # the Journal tables tie a boss to its loot where the client ships them.
    "DungeonEncounter",
    "MapDifficulty",
    "JournalEncounter",
    "JournalEncounterItem",
    "ItemSparse",
    "Item",
    "ItemEffect",
    "ItemXItemEffect",
    "SpellName",
    "Spell",
    "SpellEffect",
    # A craft spell's reagents (pipeline.client_crafted): a Forever-new
    # craft is gated by its reagents' gates, as a classic-db one is.
    "SpellReagents",
    "SpellDuration",
    # The proc, radius and range tokens in a description ($h, $n, $u, $a1, $r).
    "SpellAuraOptions",
    "SpellRadius",
    "SpellRange",
    "SpellMisc",
    "SpellCooldowns",
    "SpellCastTimes",
    "SpellPower",
    "SpellClassOptions",
    "SpellLevels",
    # A spell's form gate: "Attack Power - Feral" (ShapeshiftMask 145, Cat and
    # Bear forms) is feral attack power, not generic (pipeline/simdb/equip.py).
    "SpellShapeshift",
    "SpellItemEnchantment",
    # A spell's target cap, the "$i" in "up to $i enemies" (pipeline/spelltext.py).
    "SpellTargetRestrictions",
    # A description's `$<name>` variables (pipeline/spelltext_grammar.py): Frostbolt's
    # `$<frostdamage>`, Arcane Barrage's `$<power>`.
    "SpellDescriptionVariables",
    "SpellXDescriptionVariables",
    # A spell's weapon restriction: a scope's equip spell reaches ranged
    # weapons only (pipeline/simdb/enchants.py's ranged_weapon_only_spells).
    "SpellEquippedItems",
    "ManifestInterfaceData",
    "ItemSet",
    "ItemSetSpell",
    "ChrClasses",
    "ChrRaces",
    # Reputation names and sides (pipeline.normalize.factions): an item's
    # ItemSparse.MinFactionID names its faction by id only.
    "Faction",
    "Talent",
    "TalentTab",
    # Weapon proficiency (pipeline.proficiency.configure): SkillLine's
    # CategoryID 6 rows are the "Weapon Skills" lines (Axes, Swords, ...);
    # SkillLineAbility's ClassMask says which classes may train each one.
    "SkillLineAbility",
    # The level a class learns a weapon or armour skill at (mail and plate at 40, Dual
    # Wield at 10 or 20): pipeline.equip_skills, for the guild Loot tab's candidates.
    "SkillRaceClassInfo",
    # The 1.60 client's talent trees live in the modern trait tables rather than
    # the legacy Talent rows above, which it still ships for its own UI chrome.
    # pipeline/normalize/traits.py reads these; Classic Era has no class trait
    # trees and 404s on five of them (see OPTIONAL_TABLES).
    "SkillLine",
    "SkillLineXTraitTree",
    "TraitNode",
    "TraitNodeEntry",
    "TraitNodeXTraitNodeEntry",
    "TraitDefinition",
    "TraitEdge",
    "TraitCond",
    "TraitNodeGroup",
    "TraitNodeGroupXTraitNode",
    "TraitCurrency",
    "TraitDefinitionEffectPoints",
    "CurvePoint",
    "ItemArmorTotal",
    "ItemArmorQuality",
    "ItemArmorShield",
    "ArmorLocation",
    "RandPropPoints",
    # The simulator's own inputs. ItemEffect + ItemXItemEffect resolve an item's
    # on-equip spells (pipeline/simdb/equip.py); the five spell tables are what
    # `simconst` reads; SpellItemEnchantment is `simdb`'s enchant source.
    # SpellCategories is deliberately absent: nothing reads Category or
    # StartRecoveryCategory, and a spell's global cooldown is already
    # SpellCooldowns.StartRecoveryTime.
    # ItemDamage* resolve weapon damage for a client whose ItemSparse states no
    # MinDamage_<n> (pipeline/simdb/weapons.py), the same way ItemArmor* resolve
    # armour. SpellScaling is deliberately absent: it 404s on this lineage.
    "ItemDamageOneHand",
    "ItemDamageTwoHand",
    "ItemDamageRanged",
    "ItemDamageWand",
    "ItemDamageThrown",
]
# Tables allowed to be missing (404) for a given product/build without aborting the
# fetch. JournalInstance (the Dungeon Journal) predates Classic Era's client, so it
# doesn't exist there. The five curve tables resolve armour and stat amounts for a
# client whose ItemSparse states no literal amounts (see
# pipeline/normalize/item_curves.py); both wow_classic_era and wow_classic_beta
# 1.60.1.69893 have them today, but they are listed here defensively in case a future
# product truly lacks them -- normalize_build then falls back to no armour/stats for
# that build's curve-shaped rows, the same honest gap as before curve support existed,
# rather than aborting the fetch. Any other 404 is treated as a real failure (e.g. a
# typo).
OPTIONAL_TABLES = frozenset(
    {
        # Description variables: a build without them leaves its `$<name>` tokens raw.
        "SpellDescriptionVariables",
        "SpellXDescriptionVariables",
        "JournalInstance",
        # The rest of the Dungeon Journal: 1.60.1.70790 answers 400 "Table not
        # found" for both, as it does for JournalInstance.
        "JournalEncounter",
        "JournalEncounterItem",
        "ItemArmorTotal",
        "ItemArmorQuality",
        "ItemArmorShield",
        "ArmorLocation",
        "RandPropPoints",
        # Classic Era 1.15.9.69722 has no class trait trees: it serves the other
        # eight trait tables (all of whose rows belong to non-class trees) and
        # 404s on these five. A build without them falls back to the legacy
        # Talent reader, which is exactly what Era wants.
        "SkillLineXTraitTree",
        "TraitEdge",
        "TraitCond",
        "TraitCurrency",
        "TraitDefinitionEffectPoints",
        # The 1.60 client's ItemEffect has no ParentItemID: the item-to-effect
        # link moved to its own table, which Classic Era 404s on. A build
        # without it uses ParentItemID instead (pipeline/simdb/equip.py).
        "ItemXItemEffect",
        # Weapon damage curves, optional for the same reason the armour curves
        # above are: a product that truly lacks them emits no weapon damage
        # rather than aborting the fetch.
        "ItemDamageOneHand",
        "ItemDamageTwoHand",
        "ItemDamageRanged",
        "ItemDamageWand",
        "ItemDamageThrown",
    }
)

#: Each table's file data id in the client (the community listfile,
#: wowdev/wow-listfile `DBFilesClient/<Table>.db2`). A DB2's id never changes
#: between builds; pipeline.cdn_tables reads the tables by these.
FILE_DATA_IDS: dict[str, int] = {
    "Map": 1349477,
    "AreaTable": 1353545,
    "JournalInstance": 1237438,
    "DungeonEncounter": 1347279,
    "MapDifficulty": 1367868,
    "JournalEncounter": 1240336,
    "JournalEncounterItem": 1344467,
    "ItemSparse": 1572924,
    "Item": 841626,
    "ItemEffect": 969941,
    "ItemXItemEffect": 3177687,
    "SpellName": 1990283,
    "Spell": 1140089,
    "SpellEffect": 1140088,
    "SpellReagents": 841946,
    "SpellDuration": 1137828,
    "SpellAuraOptions": 1139952,
    "SpellRadius": 1134584,
    "SpellRange": 1146820,
    "SpellMisc": 1003144,
    "SpellCooldowns": 1139924,
    "SpellCastTimes": 1134089,
    "SpellPower": 982806,
    "SpellClassOptions": 979663,
    "SpellLevels": 1140079,
    "SpellShapeshift": 1139929,
    "SpellItemEnchantment": 1362771,
    "SpellTargetRestrictions": 1139993,
    "SpellDescriptionVariables": 1140004,
    "SpellXDescriptionVariables": 1724949,
    "SpellEquippedItems": 1140011,
    "ManifestInterfaceData": 1375801,
    "ItemSet": 1343609,
    "ItemSetSpell": 1314689,
    "ChrClasses": 1361031,
    "ChrRaces": 1305311,
    "Faction": 1361972,
    "Talent": 1369062,
    "TalentTab": 2178102,
    "SkillLineAbility": 1266278,
    "SkillRaceClassInfo": 1240406,
    "SkillLine": 1240935,
    "SkillLineXTraitTree": 4505477,
    "TraitNode": 4420297,
    "TraitNodeEntry": 4420298,
    "TraitNodeXTraitNodeEntry": 4420304,
    "TraitDefinition": 4420327,
    "TraitEdge": 4420308,
    "TraitCond": 4543085,
    "TraitNodeGroup": 4420300,
    "TraitNodeGroupXTraitNode": 4420302,
    "TraitCurrency": 4524216,
    "TraitDefinitionEffectPoints": 4536895,
    "CurvePoint": 892586,
    "ItemArmorTotal": 1283022,
    "ItemArmorQuality": 1283021,
    "ItemArmorShield": 1277741,
    "ArmorLocation": 1284818,
    "RandPropPoints": 1310245,
    "ItemDamageOneHand": 1277743,
    "ItemDamageTwoHand": 1277738,
    "ItemDamageRanged": 6156256,
    "ItemDamageWand": 6156258,
    "ItemDamageThrown": 6156257,
    "GroupFinderActivity": 974813,
    "GroupFinderCategory": 974812,
    "LFGDungeons": 1361033,
    "Lock": 1343608,
    "PlayerExpectedStat": 7667204,
}

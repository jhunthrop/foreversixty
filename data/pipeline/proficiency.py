"""Which armour and weapon subclasses each class can equip.

The client does not export this: a character learns proficiency from trainer
spells. These are the World of Warcraft Classic 1.x proficiencies, keyed by
ChrClasses.ID. Revisit them when the Forever client ships its own class data;
a wrong entry shows up as an item missing from, or wrongly offered in, one
class's gear picker.

Item.ClassID 4 is armour with subclasses
    0 miscellaneous (rings, necks, trinkets)  1 cloth (includes cloaks)
    2 leather  3 mail  4 plate  6 shield  7 libram  8 idol  9 totem
Item.ClassID 2 is weapons with subclasses
    0 one-hand axe   1 two-hand axe   2 bow          3 gun
    4 one-hand mace  5 two-hand mace  6 polearm      7 one-hand sword
    8 two-hand sword 10 staff         13 fist        15 dagger
    16 thrown        18 crossbow      19 wand
Subclasses 14 (monster items), 17 (deprecated spear) and 20 (fishing pole)
are deliberately absent: they are not player gear.
"""

from __future__ import annotations

WEAPON = 2
ARMOR = 4

_MISC = 0
_CLOTH = 1
_LEATHER = 2
_MAIL = 3
_PLATE = 4
_SHIELD = 6

ARMOR_SUBCLASSES: dict[int, frozenset[int]] = {
    1: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL, _PLATE, _SHIELD}),  # warrior
    2: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL, _PLATE, _SHIELD}),  # paladin
    3: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL}),  # hunter
    4: frozenset({_MISC, _CLOTH, _LEATHER}),  # rogue
    5: frozenset({_MISC, _CLOTH}),  # priest
    7: frozenset({_MISC, _CLOTH, _LEATHER, _MAIL, _SHIELD}),  # shaman
    8: frozenset({_MISC, _CLOTH}),  # mage
    9: frozenset({_MISC, _CLOTH}),  # warlock
    11: frozenset({_MISC, _CLOTH, _LEATHER}),  # druid
}

WEAPON_SUBCLASSES: dict[int, frozenset[int]] = {
    1: frozenset({0, 1, 2, 3, 4, 5, 6, 7, 8, 10, 13, 15, 16, 18}),  # warrior
    2: frozenset({0, 1, 4, 5, 6, 7, 8}),  # paladin
    3: frozenset({0, 1, 2, 3, 6, 7, 8, 10, 13, 15, 16, 18}),  # hunter
    4: frozenset({2, 3, 4, 7, 13, 15, 16, 18}),  # rogue
    5: frozenset({4, 10, 15, 19}),  # priest
    7: frozenset({0, 1, 4, 5, 10, 13, 15}),  # shaman
    8: frozenset({7, 10, 15, 19}),  # mage
    9: frozenset({7, 10, 15, 19}),  # warlock
    11: frozenset({4, 5, 6, 10, 13, 15}),  # druid
}


def can_equip(class_id: int, item_class_id: int, subclass_id: int) -> bool:
    if item_class_id == ARMOR:
        return subclass_id in ARMOR_SUBCLASSES.get(class_id, frozenset())
    if item_class_id == WEAPON:
        return subclass_id in WEAPON_SUBCLASSES.get(class_id, frozenset())
    return False

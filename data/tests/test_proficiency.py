from pipeline.proficiency import ARMOR, WEAPON, can_equip


def test_plate_is_warrior_and_paladin_only():
    assert can_equip(1, ARMOR, 4)
    assert can_equip(2, ARMOR, 4)
    assert not can_equip(8, ARMOR, 4)


def test_every_class_can_wear_cloth_and_miscellaneous_armour():
    for class_id in (1, 2, 3, 4, 5, 7, 8, 9, 11):
        assert can_equip(class_id, ARMOR, 0)
        assert can_equip(class_id, ARMOR, 1)


def test_shields_exclude_the_classes_that_cannot_use_them():
    assert can_equip(1, ARMOR, 6)
    assert can_equip(7, ARMOR, 6)
    assert not can_equip(4, ARMOR, 6)


def test_wands_are_for_the_three_wand_classes():
    assert can_equip(5, WEAPON, 19)
    assert can_equip(8, WEAPON, 19)
    assert can_equip(9, WEAPON, 19)
    assert not can_equip(1, WEAPON, 19)


def test_two_handed_swords_exclude_rogues():
    assert can_equip(1, WEAPON, 8)
    assert not can_equip(4, WEAPON, 8)


def test_monster_and_fishing_subclasses_belong_to_nobody():
    for class_id in (1, 4, 8):
        assert not can_equip(class_id, WEAPON, 14)
        assert not can_equip(class_id, WEAPON, 17)
        assert not can_equip(class_id, WEAPON, 20)


def test_non_equipment_item_classes_are_never_equippable():
    assert not can_equip(1, 0, 0)
    assert not can_equip(1, 7, 5)

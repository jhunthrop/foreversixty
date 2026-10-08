"""The merge-only night keeps the committed buff names and names a new engine id."""

import pytest

from pipeline.loot.buffs import BuffError, merge_simbuffs, normalise
from pipeline.models import BuffOverride, SimBuffEntry

FORK = {normalise("Windfury Totem"): ("Windfury Totem", "spell_nature_windfury")}
KEPT = {"battle_shout": SimBuffEntry(name="Battle Shout", icon="ability_warrior_battleshout")}


def test_a_committed_entry_is_kept_and_a_new_id_is_named_from_the_fork():
    out = merge_simbuffs(["battle_shout", "windfury_totem"], [FORK], KEPT, {})
    assert out.entries["battle_shout"] == KEPT["battle_shout"]
    assert out.entries["windfury_totem"] == SimBuffEntry(
        name="Windfury Totem", icon="spell_nature_windfury"
    )


def test_an_id_ids_md_dropped_is_dropped():
    out = merge_simbuffs(["windfury_totem"], [FORK], KEPT, {})
    assert list(out.entries) == ["windfury_totem"]


def test_an_override_names_the_fork_row_when_no_entry_is_committed():
    override = BuffOverride(name="Windfury Totem", spell_id=8512)
    out = merge_simbuffs(
        ["windfury_totem:improved"], [FORK], {}, {"windfury_totem:improved": override}
    )
    assert out.entries["windfury_totem:improved"].name == "Windfury Totem"


def test_a_new_id_with_no_name_anywhere_is_a_hard_error():
    with pytest.raises(BuffError, match="new since the last full loot rebuild"):
        merge_simbuffs(["eureka"], [FORK], KEPT, {})

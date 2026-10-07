import json
from pathlib import Path

import pytest

from pipeline.manifest import refresh_manifest, write_manifest
from pipeline.models import ClassTalents, TalentEntry, TalentRank, TalentTree
from pipeline.normalize.forever_talents import ForeverTalentError
from pipeline.normalize.wowhead_overlay import (
    clean_description,
    load_overlay_payload,
    overlay_wowhead_talents,
)

KNOWN = {100, 200}


def old_talent(talent_id: int, spell_id: int, tier: int, column: int) -> TalentEntry:
    return TalentEntry(
        id=talent_id,
        name="Old",
        icon="old_icon",
        max_rank=1,
        tier=tier,
        column=column,
        prereq_talent_id=None,
        prereq_rank=None,
        ranks=[TalentRank(spell_id=spell_id, description="old text")],
        spell_id=spell_id,
    )


def records() -> list[ClassTalents]:
    tree = TalentTree(
        id=164,
        name="Fury",
        position=1,
        talents=[old_talent(1, 100, 0, 0)],
        background="warriorfury",
        icon="tab_icon",
    )
    return [ClassTalents(build="b", class_id=1, class_slug="warrior", trees=[tree])]


def talent(talent_id: int, ranks: list[int], row: int, col: int, **extra) -> dict:
    return {
        "id": talent_id,
        "row": row,
        "col": col,
        "icon": "new_icon",
        "name": f"New {talent_id}",
        "ranks": ranks,
        "descriptions": {str(i): f"text {i}<br />more" for i in range(1, len(ranks) + 1)},
        "requires": [],
        **extra,
    }


def payload(*talents: dict) -> dict:
    return {"talents": {"164": {str(t["id"]): t for t in talents}}}


def test_overlay_replaces_the_talents_and_keeps_the_tree_identity():
    out = overlay_wowhead_talents(
        records(),
        payload(
            talent(2, [200, 200], 1, 1, requires=[{"id": 1, "qty": 1}]), talent(1, [100], 0, 0)
        ),
        KNOWN,
    )
    tree = out[0].trees[0]
    assert (tree.id, tree.name, tree.position, tree.background, tree.icon) == (
        164,
        "Fury",
        1,
        "warriorfury",
        "tab_icon",
    )
    assert [t.id for t in tree.talents] == [1, 2]
    second = tree.talents[1]
    assert (second.name, second.icon, second.max_rank, second.tier, second.column) == (
        "New 2",
        "new_icon",
        2,
        1,
        1,
    )
    assert (second.prereq_talent_id, second.prereq_rank) == (1, 1)
    assert [r.spell_id for r in second.ranks] == [200, 200]
    assert second.ranks[1].description == "text 2\nmore"
    assert second.spell_id == 200


def test_overlay_does_not_mutate_its_input():
    original = records()
    overlay_wowhead_talents(original, payload(talent(5, [100], 0, 0)), KNOWN)
    assert [t.id for t in original[0].trees[0].talents] == [1]


def test_a_spell_missing_from_the_client_is_kept_and_flagged():
    out = overlay_wowhead_talents(
        records(), payload(talent(7, [999], 0, 0), talent(8, [100], 0, 1)), KNOWN
    )
    by_id = {t.id: t for t in out[0].trees[0].talents}
    assert by_id[7].hotfix_only is True
    assert by_id[7].ranks[0].spell_id == 999
    assert by_id[8].hotfix_only is False


def test_hotfix_only_is_emitted_only_when_true_so_old_files_stay_byte_identical():
    out = overlay_wowhead_talents(
        records(), payload(talent(7, [999], 0, 0), talent(8, [100], 0, 1)), KNOWN
    )
    dumped = {t["id"]: t for t in out[0].model_dump()["trees"][0]["talents"]}
    assert dumped[7]["hotfix_only"] is True
    assert "hotfix_only" not in dumped[8]
    assert list(dumped[8]) == [
        "id",
        "name",
        "icon",
        "max_rank",
        "tier",
        "column",
        "prereq_talent_id",
        "prereq_rank",
        "ranks",
        "spell_id",
    ]


def test_a_payload_with_a_tree_we_lack_is_refused():
    bad = {"talents": {"164": {}, "999": {}}}
    with pytest.raises(ForeverTalentError, match="not ours"):
        overlay_wowhead_talents(records(), bad, KNOWN)


def test_a_payload_missing_one_of_our_trees_is_refused():
    with pytest.raises(ForeverTalentError, match="missing from the payload"):
        overlay_wowhead_talents(records(), {"talents": {}}, KNOWN)


def test_two_talents_in_one_cell_are_refused():
    with pytest.raises(ForeverTalentError, match="one cell"):
        overlay_wowhead_talents(
            records(), payload(talent(1, [100], 0, 0), talent(2, [100], 0, 0)), KNOWN
        )


@pytest.mark.parametrize(
    ("html", "plain"),
    [
        ("Deals <!--ppl20:26:28:35-->28 to 32 Fire damage.", "Deals 28 to 32 Fire damage."),
        ("a<br />b<br /><br /><br />c", "a\nb\n\nc"),
        ('<span style="color: #FF2020">Requires Cat Form</span><br />Go', "Requires Cat Form\nGo"),
        (
            "as if 1 <!--singular:level:levels-->levels<!--singular--> higher",
            "as if 1 level higher",
        ),
        ("for 2 <!--singular:Point:Points-->Points<!--singular-->", "for 2 Points"),
        (
            "Charge.<br /><table><tr><td><b>x</b><table><tr><td>y</td></tr></table></td></tr></table>",  # noqa: E501
            "Charge.",
        ),
        ("lasts <!--sp1:0-->12 sec<!--sp1-->.  ", "lasts 12 sec."),
    ],
)
def test_clean_description(html: str, plain: str):
    assert clean_description(html) == plain


def test_load_overlay_payload_is_none_without_a_snapshot_and_drops_meta(tmp_path: Path):
    assert load_overlay_payload(tmp_path) is None
    raw = tmp_path / "raw"
    raw.mkdir()
    (raw / "wowhead-talents.json").write_text(json.dumps({"talents": {}, "_meta": {"x": 1}}))
    assert load_overlay_payload(tmp_path) == {"talents": {}}


def test_manifest_extra_survives_a_refresh(tmp_path: Path):
    (tmp_path / "a.json").write_text("[]\n")
    write_manifest(tmp_path, "b", "p", "t", extra={"talent_source": "raw/wowhead-talents.json"})
    assert refresh_manifest(tmp_path)["talent_source"] == "raw/wowhead-talents.json"

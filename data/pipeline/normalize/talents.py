from pipeline.models import TalentNode


class TalentDataError(ValueError):
    """A talent tab names more or fewer than one class; this module will not guess."""


def _class_id_from_mask(tab_id: int, mask: int) -> int:
    """The one class a `ClassMask` bit names, or TalentDataError if it names none or many.

    `ClassMask` is `1 << (class_id - 1)`, same as `AllowableClass` on an item -- a
    single bit for every tab this pipeline has seen. A tab this pipeline cannot single
    out a class from is exactly the "no guessing" case `build_talent_trees` already
    handles by requiring a single-bit match; this mirrors that for the flat talent list.
    """
    if mask <= 0 or mask & (mask - 1) != 0:
        raise TalentDataError(f"talent tab {tab_id} has ClassMask {mask}, not a single class")
    return mask.bit_length()


def normalize_talents(
    talent_rows: list[dict[str, str]], tab_rows: list[dict[str, str]]
) -> list[TalentNode]:
    tabs = {int(t["ID"]): t["Name_lang"] for t in tab_rows}
    # The 1.60 client (Forever beta) leaves Talent.ClassID at 0 for every row -- class
    # ownership moved to TalentTab.ClassMask, which build_talent_trees already reads for
    # the per-class files. Classic Era's Talent.ClassID is still the real, correct
    # class id, so it is used when the row states one; a build that leaves it 0 falls
    # back to the tab's mask. The mask is only resolved for a tab actually needed as a
    # fallback, so a tab no talent falls back to can carry any mask shape.
    tab_masks = {int(t["ID"]): int(t["ClassMask"]) for t in tab_rows}
    out: list[TalentNode] = []
    for t in talent_rows:
        spell_ids = [
            int(t[f"SpellRank_{i}"])
            for i in range(9)
            if t.get(f"SpellRank_{i}", "0").strip() not in ("", "0")
        ]
        prereq = int(t.get("PrereqTalent_0", "0") or 0)
        tab_id = int(t["TabID"])
        class_id = int(t["ClassID"]) or _class_id_from_mask(tab_id, tab_masks.get(tab_id, 0))
        out.append(
            TalentNode(
                id=int(t["ID"]),
                tab_id=tab_id,
                tab_name=tabs.get(tab_id, ""),
                class_id=class_id,
                tier=int(t["TierID"]),
                column=int(t["ColumnIndex"]),
                spell_ids=spell_ids,
                prereq_talent_id=prereq or None,
            )
        )
    return out

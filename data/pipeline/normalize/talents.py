from pipeline.models import TalentNode


def normalize_talents(
    talent_rows: list[dict[str, str]], tab_rows: list[dict[str, str]]
) -> list[TalentNode]:
    tabs = {int(t["ID"]): t["Name_lang"] for t in tab_rows}
    out: list[TalentNode] = []
    for t in talent_rows:
        spell_ids = [
            int(t[f"SpellRank_{i}"])
            for i in range(5)
            if t.get(f"SpellRank_{i}", "0").strip() not in ("", "0")
        ]
        prereq = int(t.get("PrereqTalent_0", "0") or 0)
        out.append(
            TalentNode(
                id=int(t["ID"]),
                tab_id=int(t["TabID"]),
                tab_name=tabs.get(int(t["TabID"]), ""),
                class_id=int(t["ClassID"]),
                tier=int(t["TierID"]),
                column=int(t["ColumnIndex"]),
                spell_ids=spell_ids,
                prereq_talent_id=prereq or None,
            )
        )
    return out

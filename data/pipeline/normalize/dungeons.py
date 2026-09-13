from pipeline.models import Dungeon


def normalize_dungeons(journal_rows: list[dict[str, str]]) -> list[Dungeon]:
    out: list[Dungeon] = []
    for j in journal_rows:
        map_id = int(j["MapID"])
        out.append(
            Dungeon(
                id=int(j["ID"]),
                name=j["Name_lang"],
                map_id=map_id,
                description=j.get("Description_lang", ""),
                order=int(j.get("OrderIndex", "0") or 0),
            )
        )
    return out

from pipeline.models import Dungeon


def normalize_dungeons(
    journal_rows: list[dict[str, str]], map_rows: list[dict[str, str]]
) -> list[Dungeon]:
    known_maps = {int(m["ID"]) for m in map_rows}
    out: list[Dungeon] = []
    for j in journal_rows:
        map_id = int(j["MapID"])
        # keep entries even if Map fixture lacks them; Map is used only for
        # validation logging in Phase 0
        _ = map_id in known_maps
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

from pipeline.models import Zone


def _int_or_none(v: str) -> int | None:
    v = v.strip()
    return int(v) if v not in ("", "0") else None


def normalize_zones(area_rows: list[dict[str, str]], map_rows: list[dict[str, str]]) -> list[Zone]:
    maps = {int(m["ID"]): m["MapName_lang"] for m in map_rows}
    zones: list[Zone] = []
    for a in area_rows:
        map_id = int(a["ContinentID"])
        level = int(a["ExplorationLevel"]) if a["ExplorationLevel"].strip() else 0
        zones.append(
            Zone(
                id=int(a["ID"]),
                name=a["AreaName_lang"],
                map_id=map_id,
                map_name=maps.get(map_id, ""),
                parent_id=_int_or_none(a["ParentAreaID"]),
                level_min=level if level > 0 else None,
                level_max=None,
            )
        )
    return zones

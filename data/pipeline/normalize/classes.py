from pipeline.models import PlayableClass, PlayableRace

CLASS_COLORS = {
    "warrior": "#c69b6d",
    "paladin": "#f48cba",
    "hunter": "#aad372",
    "rogue": "#fff468",
    "priest": "#ffffff",
    "shaman": "#0070dd",
    "mage": "#3fc7eb",
    "warlock": "#8788ee",
    "druid": "#ff7c0a",
}


def _slug(name: str) -> str:
    return name.lower().replace(" ", "-").replace("'", "")


def normalize_classes(rows: list[dict[str, str]]) -> list[PlayableClass]:
    out: list[PlayableClass] = []
    for r in rows:
        slug = _slug(r["Name_lang"])
        out.append(
            PlayableClass(
                id=int(r["ID"]),
                name=r["Name_lang"],
                slug=slug,
                color=CLASS_COLORS.get(slug, "#ffffff"),
            )
        )
    return out


def normalize_races(rows: list[dict[str, str]]) -> list[PlayableRace]:
    faction_by_flag = {"0": "alliance", "1": "horde", "2": "neutral"}
    return [
        PlayableRace(
            id=int(r["ID"]),
            name=r["Name_lang"],
            slug=_slug(r["Name_lang"]),
            faction=faction_by_flag.get(r["Alliance"].strip(), "neutral"),
        )
        for r in rows
    ]

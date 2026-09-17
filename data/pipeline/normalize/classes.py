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


def slugify(name: str) -> str:
    return name.lower().replace(" ", "-").replace("'", "")


def normalize_classes(rows: list[dict[str, str]]) -> list[PlayableClass]:
    out: list[PlayableClass] = []
    for r in rows:
        slug = slugify(r["Name_lang"])
        out.append(
            PlayableClass(
                id=int(r["ID"]),
                name=r["Name_lang"],
                slug=slug,
                color=CLASS_COLORS.get(slug, "#ffffff"),
            )
        )
    return out


def _is_selectable(row: dict[str, str]) -> bool:
    """False for a row the client itself marks as not a selectable race.

    The 1.60 client (Forever beta) exports `PlayableRaceBit`, and some rows carry -1 in
    it: an internal duplicate the client never lets a player pick (for instance id 33
    "Human", an unused "ThinHuman" body-type row with placeholder flavour text, sharing
    its name with the real playable Human at id 1). Two rows sharing a slug would
    collide in `curated.merge_curated`'s by-slug lookup, so these are dropped before a
    slug is ever computed. Classic Era's ChrRaces has no `PlayableRaceBit` column at
    all, and every one of its rows is selectable, so a missing column keeps every row.
    """
    bit = row.get("PlayableRaceBit")
    return bit is None or bit.strip() != "-1"


def normalize_races(rows: list[dict[str, str]]) -> list[PlayableRace]:
    faction_by_flag = {"0": "alliance", "1": "horde", "2": "neutral"}
    return [
        PlayableRace(
            id=int(r["ID"]),
            name=r["Name_lang"],
            slug=slugify(r["Name_lang"]),
            faction=faction_by_flag.get(r["Alliance"].strip(), "neutral"),
        )
        for r in rows
        if _is_selectable(r)
    ]

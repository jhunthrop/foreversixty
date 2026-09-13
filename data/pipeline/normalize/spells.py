from pipeline.models import Spell


def normalize_spells(rows: list[dict[str, str]]) -> list[Spell]:
    return [Spell(id=int(r["ID"]), name=r["Name_lang"]) for r in rows if r["Name_lang"].strip()]

"""Writes data.json and assets/ for the mobile boards from a checkout with the synced static data
(web/public/data/<build>) and the raw client tables (data/builds/<build>/raw). Run from the repo root:
python3 design/mocks/2026-10-30-mobile-apps/extract_data.py [repo root holding the data]
The boards then run from data.json alone, so a worktree without the synced data still renders."""
import csv
import json
import shutil
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
SRC = Path(sys.argv[1]) if len(sys.argv) > 1 else HERE.parents[2]
BUILD = "1.60.1.70532"
PUB = SRC / "web/public/data" / BUILD
RAW = SRC / "data/builds" / BUILD / "raw"
SPEC_ICONS = ["spell_nature_starfall", "ability_racial_bearform", "spell_nature_healingtouch", "ability_hunter_beasttaming",
              "ability_marksmanship", "ability_hunter_swiftstrike", "spell_holy_magicalsentry", "spell_fire_firebolt02",
              "spell_frost_frostbolt02", "spell_holy_holybolt", "spell_holy_devotionaura", "spell_holy_auraoflight",
              "spell_holy_wordfortitude", "spell_holy_guardianspirit", "spell_shadow_shadowwordpain", "ability_rogue_eviscerate",
              "ability_backstab", "ability_stealth", "spell_nature_lightning", "spell_nature_lightningshield",
              "spell_nature_magicimmunity", "spell_shadow_deathcoil", "spell_shadow_metamorphosis", "spell_shadow_rainoffire",
              "ability_warrior_innerrage", "inv_shield_06"]
ONYXIA_LOOT = [17078]  # Sapphiron Drape, an Onyxia drop in loot.json (raid:onyxias-lair, boss Onyxia)
ARMOR_SUBCLASS = {1: "Cloth", 2: "Leather", 3: "Mail", 4: "Plate", 6: "Shield"}
BONDING = {1: "Binds when picked up", 2: "Binds when equipped", 3: "Binds when used"}
STANDING = {4: "Friendly", 5: "Honored", 6: "Revered", 7: "Exalted"}


def client_rows(ids: set[int]) -> dict[int, dict]:
    """Bonding, uniqueness and reputation from ItemSparse (hotfixes first), armor type from Item."""
    out: dict[int, dict] = {}
    for name in ("hotfixes/ItemSparse.csv", "ItemSparse.csv"):
        for r in csv.DictReader(open(RAW / name)):
            i = int(r["ID"])
            if i in ids and i not in out:
                out[i] = {"bonding": BONDING.get(int(r["Bonding"])), "unique": r["MaxCount"] == "1",
                          "min_standing": STANDING.get(int(r["MinReputation"]))}
    for r in csv.DictReader(open(RAW / "Item.csv")):
        i = int(r["ID"])
        if i in ids and r["ClassID"] == "4":
            out.setdefault(i, {})["armor_type"] = ARMOR_SUBCLASS.get(int(r["SubclassID"]))
    return out


def main() -> None:
    items = {i["id"]: i for i in json.load(open(PUB / "items/shaman.json"))["items"]}
    loot = json.load(open(PUB / "loot.json"))
    bis = json.load(open(PUB / "bis/shaman-elemental.json"))
    out = {"build": BUILD, "bands": {}}
    for preset in ("raid", "bare"):
        b = next(b for b in bis["bands"] if b["band"] == 60 and b["faction"] == "horde" and b["preset"] == preset)
        out["bands"][preset] = {k: b.get(k) for k in ("set_dps", "talents", "talent_points", "reference_dps_per_point", "slots")}
    ids = set(ONYXIA_LOOT)
    for band in out["bands"].values():
        for s in band["slots"]:
            if "item_id" in s:
                ids.add(s["item_id"])
            ids |= {a["item_id"] for a in s.get("alternatives", [])[:3]}
    keys = ["name", "icon", "quality", "item_level", "required_level", "slot", "armor", "stats", "effect_text", "weapon_type",
            "dps", "speed", "damage_min", "damage_max", "two_hand"]
    client = client_rows(ids)
    out["items"] = {str(i): {**{k: items[i].get(k) for k in keys}, **client.get(i, {})} for i in ids if i in items}
    rep = {}
    for src in loot["sources"]:
        if src.get("kind") == "rep":
            for i in src.get("items", []):
                if i in ids:
                    rep[str(i)] = f"{src['name']} ({src['standing'].title()})"
    out["rep_sources"] = rep
    talents = json.load(open(PUB / "talents/shaman.json"))
    out["trees"] = [{"name": t["name"], "icon": t["icon"],
                     "talents": [{**{k: x[k] for k in ("id", "name", "icon", "max_rank", "tier", "column", "prereq_talent_id", "prereq_rank")},
                                 "ranks": [r["description"] for r in x["ranks"]]} for x in t["talents"]]}
                    for t in talents["trees"]]
    (HERE / "data.json").write_text(json.dumps(out, indent=1, sort_keys=True))
    icons = {v["icon"] for v in out["items"].values()} | set(SPEC_ICONS)
    icons |= {t["icon"] for t in out["trees"]} | {x["icon"] for t in out["trees"] for x in t["talents"]}
    (HERE / "assets").mkdir(exist_ok=True)
    for stem in sorted(icons):
        shutil.copy(PUB / "icons" / f"{stem}.webp", HERE / "assets" / f"{stem}.webp")
    print(len(out["items"]), "items,", len(icons), "icons")


if __name__ == "__main__":
    main()

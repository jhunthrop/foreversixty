"""Rewrites w3.json and assets/ for the Wave 3 boards from a checkout holding data/builds/<build>/ (run from the repo root):
python3 design/mocks/2026-10-31-mobile-wave3/extract_w3.py
Real: the tier list (the DPS tier rules of web/src/lib/tiers/tier-rules.ts over the BiS bands), spell ranks, the Spritecaster Cape
item page (loot.json drops, every BiS file that lists it), the key dates, the changelog titles. Invented: guild, people, parses."""
import glob
import json
import re
import shutil
from pathlib import Path

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[2]
BUILD = ROOT / "data/builds/1.60.1.70658"
ICONS = BUILD / "icons"
TIER_FROM = [("S", 0), ("A", 5), ("B", 10), ("C", 20), ("D", 30)]
NEED_ICONS = {
    "spell_nature_chainlightning", "spell_fire_searingtotem", "spell_fire_selfdestruct", "spell_nature_manaregentotem", "spell_shaman_lavaburst",
    "spell_nature_lightning", "spell_fire_flameshock", "spell_nature_earthshock", "spell_nature_lightningshield", "inv_misc_questionmark",
    "inv_misc_cape_05", "spell_nature_wispsplode", "ability_warrior_innerrage", "ability_rogue_eviscerate", "inv_shield_06", "ability_marksmanship",
    "spell_holy_guardianspirit", "spell_shadow_shadowwordpain", "spell_frost_frostbolt02", "spell_fire_firebolt02", "spell_shadow_rainoffire",
    "spell_nature_starfall", "spell_nature_healingtouch", "ability_hunter_swiftstrike", "ability_backstab", "ability_stealth", "spell_holy_magicalsentry",
    "spell_shadow_deathcoil", "spell_shadow_metamorphosis", "ability_hunter_beasttaming", "spell_holy_wordfortitude", "spell_nature_magicimmunity",
    "ability_racial_bearform", "spell_shaman_improvedstormstrike", "spell_nature_lightningoverload", "inv_misc_book_15",
}
CLASS_OF = lambda slug: slug.split("-")[0]  # noqa: E731
SPEC_NAME = {"feral-bear": "Feral Bear", "beast-mastery": "Beast Mastery"}


def spec_name(slug: str) -> str:
    rest = slug.split("-", 1)[1]
    return SPEC_NAME.get(rest, rest.replace("-", " ").title())


def tier_for(gap: float) -> str:
    return [letter for letter, floor in TIER_FROM if gap >= floor][-1]


def tier_rows() -> list[dict]:
    rows = []
    for f in sorted(glob.glob(str(BUILD / "bis/*.json"))):
        slug = Path(f).stem
        d = json.load(open(f))
        band = next((b for b in d["bands"] if b["band"] == 60 and b["preset"] == "raid" and b["faction"] == "horde"), None)
        if band is None or band.get("role", "dps") != "dps":
            continue
        rows.append({"slug": slug, "class": CLASS_OF(slug), "name": spec_name(slug), "race": band.get("race"), "metric": band["set_dps"]})
    rows.sort(key=lambda r: -r["metric"])
    top = rows[0]["metric"]
    for i, r in enumerate(rows):
        r["rank"] = i + 1
        r["gap"] = round((top - r["metric"]) / top * 100, 1)
        r["fraction"] = round(r["metric"] / top, 4)
        r["tier"] = tier_for(r["gap"])
        r["tie"] = (i > 0 and abs(rows[i - 1]["metric"] - r["metric"]) / top * 100 <= 1) or (i + 1 < len(rows) and abs(r["metric"] - rows[i + 1]["metric"]) / top * 100 <= 1)
    return rows


def ranks() -> dict:
    sh = json.load(open(BUILD / "spellranks.json"))["classes"]["shaman"]
    out = {}
    for name in ["Lightning Bolt", "Chain Lightning", "Lava Burst", "Searing Totem", "Magma Totem", "Mana Spring Totem", "Flame Shock", "Earth Shock"]:
        rows = sh.get(name, [])
        out[name] = {"max_rank": max(r["rank"] for r in rows), "level": max(r["level"] for r in rows)} if rows else None
    return out


def item_page(item_id: int) -> dict:
    loot = json.load(open(BUILD / "loot.json"))
    drops = []
    for s in loot["sources"]:
        for b in s.get("bosses", []):
            if item_id in b["items"]:
                drops.append({"zone": s["name"], "boss": b["name"].strip(), "chance": b.get("item_chances", {}).get(str(item_id)), "kind": s["kind"]})
    uses = []
    for f in sorted(glob.glob(str(BUILD / "bis/*.json"))):
        slug = Path(f).stem
        d = json.load(open(f))
        best = None
        for b in d["bands"]:
            if b["faction"] != "horde":
                continue
            for s in b["slots"]:
                place = 0 if s.get("item_id") == item_id else next((i + 1 for i, a in enumerate(s.get("alternatives", [])) if a["item_id"] == item_id), None)
                if place is None:
                    continue
                key = (place, -b["band"], 0 if b["preset"] == "raid" else 1)
                if best is None or key < best[0]:
                    best = (key, {"spec": slug, "class": CLASS_OF(slug), "name": spec_name(slug), "band": b["band"], "preset": b["preset"], "place": place, "slot": s["slot"]})
        if best:
            uses.append(best[1])
    uses.sort(key=lambda u: (u["place"], u["class"], u["name"]))
    return {"drops": drops, "uses": uses}


def main() -> None:
    dates = json.load(open(ROOT / "web/src/data/dates.json"))
    changelog = []
    for p in sorted((ROOT / "web/src/content/changelog").glob("*.md"), reverse=True)[:6]:
        head = p.read_text().split("---")[1]
        title = re.search(r"title: '?(.+?)'?\n", head).group(1)
        kind = re.search(r"kind: (\w+)", head).group(1)
        changelog.append({"date": p.name[:10], "title": title.replace("''", "'"), "kind": kind})
    out = {"build": BUILD.name, "tiers": tier_rows(), "ranks": ranks(), "cape": item_page(11623), "dates": dates, "changelog": changelog}
    (HERE / "w3.json").write_text(json.dumps(out, indent=1))
    for stem in NEED_ICONS:
        src = ICONS / f"{stem}.webp"
        if src.exists():
            shutil.copy(src, HERE / "assets" / src.name)
        else:
            print("missing icon", stem)
    print(len(out["tiers"]), "specs;", out["cape"]["drops"], len(out["cape"]["uses"]), "uses")


if __name__ == "__main__":
    main()

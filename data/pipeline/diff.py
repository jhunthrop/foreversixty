import json
from pathlib import Path

ENTITIES = ["zones", "dungeons", "items", "spells", "classes", "races", "talents", "sets"]


def diff_entities(before: list[dict], after: list[dict]) -> dict:
    b = {r["id"]: r for r in before}
    a = {r["id"]: r for r in after}
    added = [a[i] for i in sorted(a.keys() - b.keys())]
    removed = [b[i] for i in sorted(b.keys() - a.keys())]
    changed = []
    for i in sorted(a.keys() & b.keys()):
        fields = sorted(k for k in set(a[i]) | set(b[i]) if a[i].get(k) != b[i].get(k))
        if fields:
            changed.append({"id": i, "before": b[i], "after": a[i], "fields": fields})
    return {"added": added, "removed": removed, "changed": changed}


def diff_builds(
    from_build: str,
    to_build: str,
    root: Path = Path("builds"),
    out: Path = Path("diffs"),
) -> Path:
    result = {"from": from_build, "to": to_build, "entities": {}}
    for name in ENTITIES:
        fb, tb = root / from_build / f"{name}.json", root / to_build / f"{name}.json"
        if not fb.exists() or not tb.exists():
            continue
        d = diff_entities(json.loads(fb.read_text()), json.loads(tb.read_text()))
        result["entities"][name] = d
        print(f"{name}: +{len(d['added'])} -{len(d['removed'])} ~{len(d['changed'])}")
    out.mkdir(parents=True, exist_ok=True)
    path = out / f"{from_build}__{to_build}.json"
    path.write_text(json.dumps(result, indent=2, ensure_ascii=False) + "\n", encoding="utf-8")
    return path

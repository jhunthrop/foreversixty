"""The F13 tree grid and build maths the Wave 3 boards drew with. They lived in ../2026-10-30-mobile-apps/screens_more.py until F6 moved the talents screens to screens_talents.py; copied here unchanged so the boards render again."""
from kit import DATA, asset

ICON, PX, PY = 46, 92, 64
LEVEL_POINTS = 51


def ranks_of(code: str) -> list[list[int]]:
    trees = code.split("-")
    return [[int(c, 36) for c in tree] + [0] * (len(DATA["trees"][i]["talents"]) - len(tree)) for i, tree in enumerate(trees)]


def your_build(guide: list[list[int]], state: str) -> list[list[int]]:
    if state == "same":
        return [list(t) for t in guide]
    yours = [list(t) for t in guide]
    yours[0][3] += 2  # Reverberation 0 -> 2, where the guide spends none
    yours[0][5] -= 2  # Call of Flame 3 -> 1; nothing depends on it, so the build stays legal
    return yours


def diff_count(yours: list[list[int]], guide: list[list[int]]) -> tuple[int, set[tuple[int, int]]]:
    count, cells = 0, set()
    for t, (ys, gs) in enumerate(zip(yours, guide)):
        for i, (y, g) in enumerate(zip(ys, gs)):
            count += max(y - g, 0)
            if y != g:
                cells.add((t, i))
    return count, cells


def tree_grid(ranks: list[int], differs: set[int]) -> str:
    talents = DATA["trees"][0]["talents"]
    spent = sum(ranks)
    by_id = {t["id"]: (i, t) for i, t in enumerate(talents)}
    left = (402 - (3 * PX + ICON)) // 2
    pos = lambda t: (left + t["column"] * PX, t["tier"] * PY)  # noqa: E731
    lines = ""
    for i, t in enumerate(talents):
        if t["prereq_talent_id"] is None:
            continue
        j, pre = by_id[t["prereq_talent_id"]]
        (ax, ay), (bx, by) = pos(pre), pos(t)
        h = ICON / 2
        if pre["column"] == t["column"]:
            d = f"M {ax + h} {ay + ICON} V {by}"
        else:
            side = 1 if t["column"] > pre["column"] else -1
            d = f"M {ax + h} {ay + ICON} V {by + h} H {bx + h - side * h}"
        met = ranks[j] >= (t["prereq_rank"] or 0)
        lines += (f"<path d=\"{d}\" fill=\"none\" stroke=\"#07090d\" stroke-width=\"5\" stroke-linecap=\"round\"/>"
                  f"<path d=\"{d}\" fill=\"none\" stroke=\"{'#e5b955' if met else '#262e40'}\" stroke-width=\"2\" stroke-linecap=\"round\"/>")
    cells = ""
    for i, t in enumerate(talents):
        rank = ranks[i]
        prereq_ok = t["prereq_talent_id"] is None or ranks[by_id[t["prereq_talent_id"]][0]] >= (t["prereq_rank"] or 0)
        available = spent >= t["tier"] * 5 and prereq_ok
        if rank:
            border, pill, face = "#7fd48a", ("#7fd48a" if rank >= t["max_rank"] else "#e9e4d8"), ""
        elif available:
            border, pill, face = "#e5b955", "#e5b955", ""
        else:
            border, pill, face = "#262e40", "#9a9484", "opacity:.4;"
        ring = "outline:2px dashed #f0736b;outline-offset:3px;" if i in differs else ""
        x, y = pos(t)
        cells += (f"<span style=\"position:absolute;left:{x}px;top:{y}px\"><img src=\"{_icon(t['icon'])}\" alt=\"\" style=\"display:block;width:{ICON}px;height:{ICON}px;border-radius:8px;border:2px solid {border};{face}{ring}\">"
                  f"<span class=\"mono\" style=\"position:absolute;right:-7px;bottom:-7px;font-size:11px;padding:1px 5px;border-radius:6px;background:#07090d;border:1px solid {border};color:{pill}\">{rank}/{t['max_rank']}</span></span>")
    tiers = max(t["tier"] for t in talents) + 1
    height = (tiers - 1) * PY + ICON + 10
    return (f"<div style=\"position:relative;height:{height}px;margin-top:18px\"><svg width=\"402\" height=\"{height}\" style=\"position:absolute;left:0;top:0\" aria-hidden=\"true\">{lines}</svg>{cells}</div>")


def _icon(stem: str) -> str:
    return asset(stem)


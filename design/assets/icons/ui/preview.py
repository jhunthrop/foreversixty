"""Write preview.html: the site's one icon set (the custom set and the line family) at 16, 20 and
24 px on the dark ground, in the cream and gold the site draws them in. preview.png is its capture."""

from pathlib import Path

HERE = Path(__file__).resolve().parent
SIZES = [(16, "#f3eee2"), (20, "#f3eee2"), (24, "#ffec71")]
FAMILIES = [("Custom set", HERE), ("Line family", HERE / "line")]


def cell(path: Path) -> str:
    svg = path.read_text().replace("<svg ", '<svg width="{s}" height="{s}" ', 1)
    sizes = "".join(f'<span style="color:{c}">{svg.format(s=s)}</span>' for s, c in SIZES)
    return f'<div class="c">{sizes}<b>{path.stem}</b></div>'


def family(title: str, folder: Path) -> str:
    cells = "".join(cell(p) for p in sorted(folder.glob("*.svg")))
    return f"<h2>{title}</h2><div class=\"g\">{cells}</div>"


def main() -> None:
    style = ("body{background:#170e09;color:#f3eee2;font:13px 'Open Sans',sans-serif;padding:24px;margin:0}"
             "h1,h2{font-weight:400}h2{font-size:16px;color:#b9ab95;margin:22px 0 10px}"
             ".g{display:grid;grid-template-columns:repeat(4,1fr);gap:12px}"
             ".c{background:#1f140d;border:1px solid #3a2a1c;border-radius:8px;padding:12px 14px;"
             "display:flex;align-items:center;gap:14px}b{margin-left:auto;font-weight:600}")
    body = "".join(family(title, folder) for title, folder in FAMILIES)
    (HERE / "preview.html").write_text(
        f'<!doctype html><meta charset="utf-8"><title>UI icons</title><style>{style}</style>'
        f'<h1>UI icons at 16, 20 and 24 px</h1>{body}\n')


if __name__ == "__main__":
    main()

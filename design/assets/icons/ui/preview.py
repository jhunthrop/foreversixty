"""Write preview.html: every traced icon at 16, 20, 24 and 48 px in the site's colours."""

from pathlib import Path

HERE = Path(__file__).resolve().parent
SIZES = [(16, "#f3eee2"), (20, "#f3eee2"), (24, "#ffec71"), (48, "#0093b2")]


def cell(name: str) -> str:
    svg = (HERE / f"{name}.svg").read_text().replace("<svg ", '<svg width="{s}" height="{s}" ')
    sizes = "".join(f'<span style="color:{c}">{svg.format(s=s)}</span>' for s, c in SIZES)
    return f'<div class="c">{sizes}<b>{name}</b></div>'


def main() -> None:
    names = sorted(p.stem for p in HERE.glob("*.svg"))
    style = ("body{background:#170e09;color:#f3eee2;font:13px 'Open Sans',sans-serif;padding:24px}"
             ".g{display:grid;grid-template-columns:repeat(4,1fr);gap:18px}"
             ".c{background:#1f140d;border:1px solid #3a2a1c;border-radius:10px;padding:14px;"
             "display:flex;align-items:center;gap:14px}b{margin-left:auto;font-weight:600}")
    body = "".join(cell(n) for n in names)
    (HERE / "preview.html").write_text(
        f'<!doctype html><meta charset="utf-8"><title>UI icons</title><style>{style}</style>'
        f'<h1 style="font-weight:400">UI icons at 16, 20, 24 and 48 px</h1><div class="g">{body}</div>\n')


if __name__ == "__main__":
    main()

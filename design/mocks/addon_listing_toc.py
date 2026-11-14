"""Boards for the addon's own line in the client's AddOns list (addon growth lane, 2026-11-14).

    python3 design/mocks/addon_listing_toc.py            # every board
    python3 design/mocks/addon_listing_toc.py addons-list-tooltip

Writes design/reviews/2026-11-14-addon-growth/<board>.png at device scale 2 through addon_ux.py's
client tooltip look. Nothing is typed here: the after side reads Title and Notes from
addon/ForeverSixty/ForeverSixty.toc as it stands, the before side from the same file at origin/main
(BEFORE_REF overrides), and the icon is the TOC's IconTexture (addon/ForeverSixty/icon.png).
The client draws the AddOns list tooltip as the title, then the notes wrapped in white; the row is the
list's checkbox, the TOC icon and the title.
"""
import os
import subprocess
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
import addon_ux as ux  # noqa: E402
from addon_ux import l, page, tooltip  # noqa: E402

ROOT = Path(__file__).resolve().parents[2]
OUT = Path(os.environ['MOCK_OUT']) if os.environ.get('MOCK_OUT') else ROOT / 'design/reviews/2026-11-14-addon-growth'
TOC = 'addon/ForeverSixty/ForeverSixty.toc'
BEFORE_REF = os.environ.get('BEFORE_REF', 'origin/main')
ICON = ux.b64((ROOT / 'addon/ForeverSixty/icon.png').read_bytes(), 'image/png')
GOLD = '#ffd100'


def fields(text):
    out = {}
    for line in text.splitlines():
        if line.startswith('## ') and ':' in line:
            key, _, value = line[3:].partition(':')
            out[key.strip()] = value.strip()
    return out


def toc_now():
    return fields((ROOT / TOC).read_text())


def toc_before():
    return fields(subprocess.run(['git', 'show', f'{BEFORE_REF}:{TOC}'], cwd=ROOT, check=True,
                                 capture_output=True, text=True).stdout)


def row(toc):
    return (f'<div style="display:flex;align-items:center;gap:8px;padding:4px 8px;background:#0b0b0b;'
            f'border:1px solid #333;width:300px;color:{GOLD};font-size:14px">'
            f'<span style="width:14px;height:14px;border:1px solid #888;background:#222;color:#ffd100;'
            f'font-size:12px;line-height:14px;text-align:center">&#10003;</span>'
            f'<img src="{ICON}" style="width:18px;height:18px">{toc["Title"]}</div>')


def side(caption, toc):
    tip = tooltip([l(toc['Title'], color=GOLD, cls='big'), f'<div style="color:#fff;white-space:normal;margin-top:2px">{toc["Notes"]}</div>'], width=300)
    return (f'<div style="display:flex;flex-direction:column;gap:8px">'
            f'<div style="color:#aaa;font-size:13px">{caption}</div>{row(toc)}{tip}</div>')


def board_list_tooltip():
    return page(f'<div style="display:flex;gap:40px;padding:24px;background:#141414">'
                f'{side("Before (origin/main)", toc_before())}{side("After", toc_now())}</div>')


BOARDS = {'addons-list-tooltip': board_list_tooltip}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1200, 'height': 600})
        pg = ctx.new_page()
        for name in names:
            pg.set_content(BOARDS[name]())
            pg.evaluate('document.fonts.ready')
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

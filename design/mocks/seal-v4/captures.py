"""Before/after captures of the seal swap: site header and footer, a favicon tab mock, the addon icons.

    python3 design/mocks/seal-v4/captures.py <before-dir>

<before-dir> holds the outgoing files (git show origin/main:<path> > <before-dir>/<name>): seal-88.webp,
favicon-32.png, favicon.ico, icon.png, minimap.tga. "After" is the working tree. The header and footer come from
the home mock (design/mocks/web/web-home-60.html), which draws web/public/art/seal-88.webp; the before run swaps
that one request for the outgoing bytes. Writes design/reviews/2026-11-09-seal-v4/<surface>-{before,after}.png.
"""
import base64
import io
import os
import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont
from playwright.sync_api import sync_playwright

REPO = Path(__file__).resolve().parents[3]
OUT = REPO / 'design/reviews/2026-11-09-seal-v4'
MOCK = REPO / 'design/mocks/web/web-home-60.html'
FONT = REPO / 'api/internal/brand/fonts/OpenSans-SemiBold.ttf'
VIEWPORT = {'width': 1440, 'height': 900}
TAB_GROUNDS = (('dark', '#35363a', '#202124', '#e8eaed'), ('light', '#ffffff', '#dee1e6', '#202124'))
ADDON_GROUND = (14, 12, 10)
ZOOM = 4


def tab_html(favicon: str) -> str:
    """Two browser tab strips, dark and light, each tab with the 16 px favicon and the site title."""
    strips = ''.join(
        f'<div style="background:{strip};padding:8px 8px 0;display:flex;gap:2px">'
        f'<div style="background:{tab};color:{ink};border-radius:8px 8px 0 0;padding:8px 12px;width:240px;'
        f'display:flex;align-items:center;gap:8px;font:12px system-ui;white-space:nowrap">'
        f'<img src="{favicon}" width="16" height="16">Forever Sixty: Play your class better</div>'
        f'<div style="color:{ink};opacity:.6;padding:8px 12px;font:12px system-ui">Wowhead</div></div>'
        for _, tab, strip, ink in TAB_GROUNDS)
    return f'<body style="margin:0"><div id="tabs" style="width:420px">{strips}</div></body>'


def capture_site(page, before: Path | None, label: str) -> None:
    if before:
        page.route('**/seal-88.webp', lambda route: route.fulfill(body=(before / 'seal-88.webp').read_bytes(),
                                                                   content_type='image/webp'))
    page.goto(MOCK.as_uri())
    page.wait_for_load_state('networkidle')
    page.locator('header').first.screenshot(path=str(OUT / f'header-{label}.png'))
    page.locator('footer').first.screenshot(path=str(OUT / f'footer-{label}.png'))
    page.unroute('**/seal-88.webp')


def capture_tab(page, favicon: bytes, label: str) -> None:
    uri = 'data:image/png;base64,' + base64.b64encode(favicon).decode()
    page.set_content(tab_html(uri))
    page.wait_for_load_state('networkidle')
    page.locator('#tabs').screenshot(path=str(OUT / f'favicon-tab-{label}.png'))


def addon_sheet(icon: Path, minimap: Path, label: str) -> None:
    """The list icon (64) and the minimap button (32), each at 1:1 and at 4x nearest, on a dark frame."""
    font = ImageFont.truetype(str(FONT), 16)
    icons = [(Image.open(icon).convert('RGBA'), 'icon.png 64'), (Image.open(minimap).convert('RGBA'), 'minimap.tga 32')]
    width = sum(im.width + im.width * ZOOM + 72 for im, _ in icons) + 24
    sheet = Image.new('RGB', (width, 64 * ZOOM + 64), ADDON_GROUND)
    draw = ImageDraw.Draw(sheet)
    x = 24
    for im, name in icons:
        draw.text((x, 8), name, fill=(220, 210, 190), font=font)
        sheet.paste(im, (x, 40), im)
        big = im.resize((im.width * ZOOM, im.height * ZOOM), Image.NEAREST)
        sheet.paste(big, (x + im.width + 24, 40), big)
        x += im.width + big.width + 72
    sheet.save(OUT / f'addon-icons-{label}.png')


def favicon16(ico: Path) -> bytes:
    """The ico's 16 px frame as PNG bytes, which the tab mock draws at 1:1."""
    out = io.BytesIO()
    Image.open(io.BytesIO(ico.read_bytes())).ico.getimage((16, 16)).convert('RGBA').save(out, 'PNG')
    return out.getvalue()


def main() -> None:
    before = Path(sys.argv[1]).resolve()
    after = {'favicon.ico': REPO / 'web/public/favicon.ico', 'favicon-32.png': REPO / 'web/public/favicon-32.png',
             'icon.png': REPO / 'addon/ForeverSixty/icon.png', 'minimap.tga': REPO / 'addon/ForeverSixty/media/minimap.tga'}
    outgoing = {name: before / name for name in after}
    with sync_playwright() as p:
        browser = p.chromium.launch(executable_path=os.environ.get('CHROMIUM'))  # any installed Chrome for Testing
        for scale, suffix in ((1, ''), (2, '@2x')):
            page = browser.new_page(viewport=VIEWPORT, device_scale_factor=scale)
            capture_site(page, before, f'before{suffix}')
            capture_site(page, None, f'after{suffix}')
            for label, files in (('before', outgoing), ('after', after)):
                tab_icon = files['favicon-32.png'].read_bytes() if scale == 2 else favicon16(files['favicon.ico'])
                capture_tab(page, tab_icon, f'{label}{suffix}')
            page.close()
        browser.close()
    for label, files in (('before', outgoing), ('after', after)):
        addon_sheet(files['icon.png'], files['minimap.tga'], label)
    print(f'wrote captures to {OUT}')


if __name__ == '__main__':
    main()

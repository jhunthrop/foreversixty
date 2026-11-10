"""Before/after captures of the flat-seal swap (owner 2026-11-10): site header, favicon tab, addon button, tray.

    python3 design/mocks/seal-flat/captures.py <before-dir> [site-url]

<before-dir> holds the outgoing files (git show HEAD:<path> > <before-dir>/<name>): favicon.ico, favicon-32.png,
minimap.tga, seal-tray-off.ico, seal-tray-on.ico, seal-tray-off.png, seal-tray-warn.png, seal-64.png. "After" is the
working tree. The header comes from the built site served at site-url (default http://localhost:4399, `npm run
preview -- --port 4399` in web/); the before run serves the painted seal-88.webp in place of the flat SVG.
Writes design/reviews/2026-11-10-seal-flat/<surface>-{before,after}*.png. Needs Pillow and Playwright.
"""
import base64
import io
import os
import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont
from playwright.sync_api import sync_playwright

REPO = Path(__file__).resolve().parents[3]
OUT = REPO / 'design/reviews/2026-11-10-seal-flat'
FONT = REPO / 'api/internal/brand/fonts/OpenSans-SemiBold.ttf'
PAINTED_88 = REPO / 'web/public/art/seal-88.webp'
HEADER_WIDTHS = (1440, 390, 2000)
TAB_GROUNDS = (('dark', '#35363a', '#202124', '#e8eaed'), ('light', '#ffffff', '#dee1e6', '#202124'))
TASKBARS = (('Windows dark', (32, 32, 32)), ('Windows light', (243, 243, 243)))
MENUBARS = (('macOS dark', (40, 40, 44)), ('macOS light', (236, 236, 236)))
ADDON_GROUND = (14, 12, 10)
MINIMAP_SHOWN_PX = 20  # Theme.SIZES.minimapIcon
ZOOM = 4
INK = (220, 210, 190)


def label_font() -> ImageFont.FreeTypeFont:
    return ImageFont.truetype(str(FONT), 14)


def capture_headers(browser, site: str, label: str, before: bool) -> None:
    for scale in (1, 2):
        for width in HEADER_WIDTHS:
            page = browser.new_page(viewport={'width': width, 'height': 700}, device_scale_factor=scale)
            if before:
                page.route('**/seal-flat-small.svg',
                           lambda route: route.fulfill(body=PAINTED_88.read_bytes(), content_type='image/webp'))
            page.goto(f'{site}/')
            page.wait_for_load_state('networkidle')
            page.locator('header.site-header').first.screenshot(path=str(OUT / f'header-{width}-dpr{scale}-{label}.png'))
            page.close()


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


def png_uri(data: bytes) -> str:
    return 'data:image/png;base64,' + base64.b64encode(data).decode()


def ico_frame(ico: bytes, px: int) -> Image.Image:
    return Image.open(io.BytesIO(ico)).ico.getimage((px, px)).convert('RGBA')


def to_png(im: Image.Image) -> bytes:
    out = io.BytesIO()
    im.save(out, 'PNG')
    return out.getvalue()


def capture_tabs(browser, site: str, before_dir: Path) -> None:
    """The tab at 1x and 2x: before, the ico's 16 frame and favicon-32.png; after, favicon.svg and the new ico's 16."""
    sources = {
        'before': {1: png_uri(to_png(ico_frame((before_dir / 'favicon.ico').read_bytes(), 16))),
                   2: png_uri((before_dir / 'favicon-32.png').read_bytes())},
        'after': {1: f'{site}/favicon.svg', 2: f'{site}/favicon.svg'},
        'after-ico': {1: png_uri(to_png(ico_frame((REPO / 'web/public/favicon.ico').read_bytes(), 16))),
                      2: png_uri(to_png(ico_frame((REPO / 'web/public/favicon.ico').read_bytes(), 32)))},
    }
    for label, by_scale in sources.items():
        for scale, uri in by_scale.items():
            page = browser.new_page(viewport={'width': 440, 'height': 200}, device_scale_factor=scale)
            page.set_content(tab_html(uri))
            page.wait_for_load_state('networkidle')
            page.locator('#tabs').screenshot(path=str(OUT / f'favicon-tab-dpr{scale}-{label}.png'))
            page.close()


def pair_sheet(rows: list[tuple[str, Image.Image]], ground: tuple[int, int, int], name: str) -> None:
    """Each image at 1:1 and at 4x nearest, labelled, one row each."""
    font = label_font()
    height = sum(max(im.height * ZOOM, 24) + 40 for _, im in rows) + 16
    width = 560
    sheet = Image.new('RGB', (width, height), ground)
    draw = ImageDraw.Draw(sheet)
    ink = INK if sum(ground) < 384 else (40, 36, 30)
    y = 12
    for text, im in rows:
        draw.text((16, y), text, fill=ink, font=font)
        sheet.paste(im, (16, y + 24), im)
        big = im.resize((im.width * ZOOM, im.height * ZOOM), Image.NEAREST)
        sheet.paste(big, (16 + im.width + 32, y + 24), big)
        y += max(im.height * ZOOM, 24) + 40
    sheet.save(OUT / name)


def addon_sheet(before_dir: Path) -> None:
    """The minimap texture as the button draws it (32 px texture at 20 px) before and after, on a dark ground."""
    rows = []
    for label, path in (('before', before_dir / 'minimap.tga'),
                        ('after', REPO / 'addon/ForeverSixty/media/minimap.tga')):
        texture = Image.open(path).convert('RGBA')
        rows.append((f'{label}: minimap.tga drawn at {MINIMAP_SHOWN_PX} px',
                     texture.resize((MINIMAP_SHOWN_PX, MINIMAP_SHOWN_PX), Image.BILINEAR)))
        rows.append((f'{label}: minimap.tga at its 32 px', texture))
    pair_sheet(rows, ADDON_GROUND, 'addon-minimap-button.png')


def tray_sheet(before_dir: Path) -> None:
    """The Windows tray ico frames at 100/150/200 % and the macOS menu-bar PNG at 16 pt, on dark and light bars."""
    after = REPO / 'companion/internal/icon/assets'
    for bar, ground in TASKBARS:
        rows = []
        for label, base in (('before', before_dir), ('after', after)):
            for tone in ('off', 'on'):
                ico = (base / f'seal-tray-{tone}.ico').read_bytes()
                for px in (16, 24, 32):
                    rows.append((f'{label} {tone} {px} px ({px * 100 // 16} %)', ico_frame(ico, px)))
        pair_sheet(rows, ground, f'tray-{bar.split()[1]}-windows.png')
    for bar, ground in MENUBARS:
        rows = []
        for label, base in (('before', before_dir), ('after', after)):
            for tone in ('off', 'warn'):
                im = Image.open(base / f'seal-tray-{tone}.png').convert('RGBA')
                rows.append((f'{label} {tone}: 32 px, 16 pt on Retina', im))
                rows.append((f'{label} {tone}: on a 1x screen', im.resize((16, 16), Image.LANCZOS)))
        pair_sheet(rows, ground, f'tray-{bar.split()[1]}-macos.png')


def capture_companion_header(browser, before_dir: Path) -> None:
    """The companion window's band wordmark (32 px seal) before and after, at 1x and 2x."""
    index = REPO / 'companion/ui/index.html'
    for scale in (1, 2):
        for label in ('before', 'after'):
            page = browser.new_page(viewport={'width': 800, 'height': 200}, device_scale_factor=scale)
            if label == 'before':
                page.route('**/seal.svg', lambda route: route.fulfill(body=(before_dir / 'seal-64.png').read_bytes(),
                                                                      content_type='image/png'))
            page.goto(index.as_uri())
            page.wait_for_timeout(500)
            page.locator('#wordmark').screenshot(path=str(OUT / f'companion-wordmark-dpr{scale}-{label}.png'))
            page.close()


def main() -> None:
    before_dir = Path(sys.argv[1]).resolve()
    site = sys.argv[2] if len(sys.argv) > 2 else 'http://localhost:4399'
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as p:
        browser = p.chromium.launch(executable_path=os.environ.get('CHROMIUM'))
        capture_headers(browser, site, 'before', before=True)
        capture_headers(browser, site, 'after', before=False)
        capture_tabs(browser, site, before_dir)
        capture_companion_header(browser, before_dir)
        browser.close()
    addon_sheet(before_dir)
    tray_sheet(before_dir)
    print(f'wrote captures to {OUT}')


if __name__ == '__main__':
    main()

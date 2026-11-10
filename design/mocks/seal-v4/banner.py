"""The Discord banner (design/assets/art/discord-banner.html: tavern art, the seal, the name) with the seal v4.

    python3 design/mocks/seal-v4/banner.py <out-dir>

Renders the 960 x 540 layout at 1x and 2x into <out-dir>/discord-banner-960.jpg and discord-banner.jpg (1920 x
1080, the size the 2026-11-09 upload had), and copies the 2x to design/assets/art/discord-banner-seal-v4.jpg.
Set CHROMIUM to a Chrome for Testing binary when Playwright's own browsers are not installed.
"""
import os
import shutil
import sys
import tempfile
from pathlib import Path

from playwright.sync_api import sync_playwright

REPO = Path(__file__).resolve().parents[3]
ART = REPO / 'design/assets/art'
LAYOUT = ART / 'discord-banner.html'
FILES = {'tavern.jpeg': ART / 'discord-banner-tavern.jpeg', 'seal.png': REPO / 'design/assets/logo/lx-seal-v4-512.png',
         'marcellus.ttf': REPO / 'apps/shared/fonts/Marcellus-Regular.ttf'}
SIZE = {'width': 960, 'height': 540}
JPEG_QUALITY = 90


def main() -> None:
    out = Path(sys.argv[1]).resolve()
    out.mkdir(parents=True, exist_ok=True)
    with tempfile.TemporaryDirectory() as tmp:
        work = Path(tmp)
        for name, src in FILES.items():
            shutil.copyfile(src, work / name)
        (work / 'banner.html').write_text(LAYOUT.read_text().replace('marcellus.woff2', 'marcellus.ttf'))
        with sync_playwright() as p:
            browser = p.chromium.launch(executable_path=os.environ.get('CHROMIUM'))
            for scale, name in ((1, 'discord-banner-960.jpg'), (2, 'discord-banner.jpg')):
                page = browser.new_page(viewport=SIZE, device_scale_factor=scale)
                page.goto((work / 'banner.html').as_uri())
                page.wait_for_load_state('networkidle')
                page.evaluate('document.fonts.ready')
                page.locator('.b').screenshot(path=str(out / name), type='jpeg', quality=JPEG_QUALITY)
                page.close()
            browser.close()
    shutil.copyfile(out / 'discord-banner.jpg', ART / 'discord-banner-seal-v4.jpg')
    print(f'wrote {out}/discord-banner.jpg and discord-banner-960.jpg')


if __name__ == '__main__':
    main()

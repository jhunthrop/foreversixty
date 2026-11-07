"""Before/after sheets for the setup area (design/specs/2026-11-07-setup.md section 0): the live capture and the mock at
the same scale, then the tracked 256-colour copies in design/reviews/2026-11-07-setup/{before,after}.
  python3 design/mocks/ws_setup_compare.py <before dir with the captures>
The before captures are production (foreversixty.gg, 2026-11-07) with the account reads answered from the same fixtures
the mocks read (home/me-34, me-new, setup/devices, setup/connect-device); see the spec, section 0."""
import sys
from pathlib import Path

from PIL import Image

import ws_compose

ROOT = Path(__file__).resolve().parents[2]
RENDERS = Path(__file__).resolve().parent / 'renders' / 'web'
REVIEW = ROOT / 'design/reviews/2026-11-07-setup'
ALL, TWO = (390, 1440, 1920, 2000), (390, 1440)
PAIRS = [('setup', 'setup-out', 'web-setup', ALL), ('setup-none', 'setup-new', 'web-setup-none', TWO),
         ('setup-bnet', 'setup-bnet', 'web-setup-bnet', TWO), ('setup-done', 'setup-done', 'web-setup-done', (390, 1440, 2000)),
         ('connect', 'connect-out', 'web-connect', ALL), ('connect-pending', 'connect-pending', 'web-connect-pending', ALL),
         ('connect-success', 'connect-success', 'web-connect-success', TWO), ('connect-used', 'connect-used', 'web-connect-used', TWO),
         ('connect-failed', 'connect-failed', 'web-connect-failed', TWO), ('connect-cancelled', 'connect-cancelled', 'web-connect-cancelled', TWO),
         ('connect-expired', 'connect-expired', 'web-connect-expired', TWO)]
SCALE = {390: .9, 1440: .6, 1920: .5, 2000: .5}
TRACK_BEFORE = ['setup-out-390', 'setup-out-1440', 'setup-out-2000', 'setup-new-1440', 'setup-bnet-390', 'setup-done-1440',
                'connect-out-390', 'connect-out-1440', 'connect-pending-390', 'connect-pending-1440', 'connect-success-1440', 'connect-expired-1440']
KEY_2000 = {'web-setup-2000-fold.png', 'web-setup-bnet-2000-fold.png', 'web-setup-done-2000-fold.png', 'web-connect-pending-2000-fold.png'}


def pal(src: Path, dst: Path, max_h: int = 4000) -> None:
    im = Image.open(src).convert('RGB')
    im = im.crop((0, 0, im.width, min(im.height, max_h)))
    dst.parent.mkdir(parents=True, exist_ok=True)
    im.quantize(256).save(dst, optimize=True)


def main(before: Path) -> None:
    ws_compose.BEFORE = before
    for page, b, a, widths in PAIRS:
        for w in widths:
            out = ws_compose.compose(f'setup-{page}' if not page.startswith('setup') else page, b, a, w, SCALE[w], max_h=3000,
                                     before_label='live' if page in ('setup', 'connect') else 'live, account reads stubbed')
            pal(out, REVIEW / 'after' / out.name, 3000)
    for n in TRACK_BEFORE:
        pal(before / f'{n}.png', REVIEW / 'before' / f'{n}.png', 3200)
    for p in sorted(RENDERS.glob('web-*-fold.png')):
        if not p.name.startswith(('web-setup', 'web-connect')):
            continue
        w = p.stem.split('-fold')[0].rsplit('-', 1)[1]
        if w in ('1440', '390') or p.name in KEY_2000:
            pal(p, REVIEW / 'after' / p.name)
    for full in ('web-setup-1440', 'web-setup-bnet-1440', 'web-setup-390', 'web-connect-pending-390'):
        pal(RENDERS / f'{full}.png', REVIEW / 'after' / f'{full}-full.png', 6000)


if __name__ == '__main__':
    main(Path(sys.argv[1]))

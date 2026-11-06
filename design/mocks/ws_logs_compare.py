"""Before/after sheets for the logs area (design/specs/2026-11-07-logs-reports.md section 0): the live or fixture capture
and the mock at the same scale, then the tracked 256-colour copies in design/reviews/2026-11-07-logs/{before,after}.
  python3 design/mocks/ws_logs_compare.py <before dir with the captures>"""
import sys
from pathlib import Path

from PIL import Image

import ws_compose

ROOT = Path(__file__).resolve().parents[2]
RENDERS = Path(__file__).resolve().parent / 'renders' / 'web'
REVIEW = ROOT / 'design/reviews/2026-11-07-logs'
# (sheet name, before capture, mock, widths). Production where the page is public; the fixture build (signed-in stubs,
# the e2e report) for the rest, captured at 1440 and 390.
PAIRS = [('logs', 'logs', 'web-logs', (390, 1440, 1920, 2000)), ('reports', 'reports', 'web-reports', (390, 1440, 1920, 2000)),
         ('report-sample', 'sample', 'web-report-sample', (390, 1440, 1920, 2000)), ('rankings', 'rankings', 'web-rankings-empty', (390, 1440, 1920, 2000)),
         ('ratings', 'ratings', 'web-ratings', (390, 1440, 1920, 2000)), ('logs-you', 'logs-you', 'web-logs-you', (390, 1440)),
         ('reports-list', 'reports-list', 'web-reports-list', (390, 1440)), ('reports-empty', 'reports-empty', 'web-reports-empty', (390, 1440)),
         ('rankings-you', 'rankings-you', 'web-rankings-you', (390, 1440)), ('rankings-board', 'rankings-slug', 'web-rankings-board', (390, 1440)),
         ('report-summary', 'report-summary', 'web-report-summary', (390, 1440)), ('report-damage-done', 'report-damage-done', 'web-report-damage-done', (390, 1440)),
         ('report-deaths', 'report-deaths', 'web-report-deaths', (1440,)), ('report-buffs', 'report-buffs', 'web-report-buffs', (1440,)),
         ('report-threat', 'report-threat', 'web-report-threat', (1440,)), ('report-casts', 'report-casts', 'web-report-casts', (1440,)),
         ('report-mechanics', 'report-mode-mechanics', 'web-report-mechanics', (1440,)), ('report-timeline', 'report-view-timelines', 'web-report-timeline', (1440,)),
         ('report-events', 'report-view-events', 'web-report-events', (1440,))]
SCALE = {390: .9, 1440: .6, 1920: .5, 2000: .5}
TRACK_BEFORE = ['logs-1440', 'logs-390', 'logs-2000', 'reports-1440', 'sample-1440', 'sample-390', 'sample-2000', 'rankings-1440', 'ratings-1440',
                'logs-you-1440', 'reports-list-1440', 'rankings-you-1440', 'rankings-you-390', 'report-summary-1440', 'report-summary-390',
                'report-damage-done-1440', 'report-deaths-1440', 'report-mode-mechanics-1440', 'report-view-timelines-1440']


def pal(src: Path, dst: Path, max_h: int = 4000) -> None:
    im = Image.open(src).convert('RGB')
    im = im.crop((0, 0, im.width, min(im.height, max_h)))
    dst.parent.mkdir(parents=True, exist_ok=True)
    im.quantize(256).save(dst, optimize=True)


def main(before: Path) -> None:
    ws_compose.BEFORE = before
    for page, b, a, widths in PAIRS:
        for w in widths:
            out = ws_compose.compose(f'logs-{page}', b, a, w, SCALE[w], max_h=3600, before_label='live' if w > 1440 or page in ('logs', 'reports', 'report-sample', 'rankings', 'ratings') else 'fixture build')
            pal(out, REVIEW / 'after' / out.name, 6000)
    for n in TRACK_BEFORE:
        pal(before / f'{n}.png', REVIEW / 'before' / f'{n}.png')
    for p in sorted(RENDERS.glob('web-*-fold.png')):
        w = p.stem.split('-fold')[0].rsplit('-', 1)[1]
        key2000 = any(p.name.startswith(x) for x in ('web-logs-you-', 'web-report-summary-', 'web-rankings-you-', 'web-ratings-'))
        if (w in ('1440', '390') or (w == '2000' and key2000)) and any(p.name.startswith(x) for x in ('web-logs', 'web-reports', 'web-report-', 'web-rankings', 'web-ratings')):
            pal(p, REVIEW / 'after' / p.name)


if __name__ == '__main__':
    main(Path(sys.argv[1]))

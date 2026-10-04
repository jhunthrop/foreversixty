"""Screenshot every board under design/mocks/renders with Playwright (run from web/ so the
browser is found): python3 ../design/mocks/render.py [name ...]. Phone boards (name
contains 'phone') render at 390 wide, others at 1440; full page, after fonts load."""
import subprocess, sys
from pathlib import Path

RENDERS = Path(__file__).resolve().parent / 'renders'
names = sys.argv[1:] or [p.stem for p in RENDERS.glob('*.html')]
for name in names:
    width = 390 if 'phone' in name else 1440
    src = RENDERS / f'{name}.html'
    out = RENDERS / f'{name}.png'
    subprocess.run(['npx', 'playwright', 'screenshot', '--browser=chromium', f'--viewport-size={width},900', '--full-page',
                    '--wait-for-timeout=3000', src.as_uri(), str(out)], check=True, capture_output=True)
    print(out)

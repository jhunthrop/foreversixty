"""Mock boards for design/specs/2026-10-31-tier-column.md: before (shipped) and after (spec) of the
tier column, rendered from the real tier-list mock (gen_tier_list.py) with the spec's CSS laid over it.

  python3 design/mocks/2026-10-31-tier-column/gen.py   -> {before,after}-{1440,2000,360}.png, plus -hover.png
"""
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent))
import mocklib
import gen_tier_list as g
from playwright.sync_api import sync_playwright

mocklib.RENDERS = HERE / 'html'
CHROME = next(Path.home().glob('Library/Caches/ms-playwright/chromium_headless_shell-*/*/chrome-headless-shell')).as_posix()

# The spec's CSS, token for token (mock vars: --raised = --color-raised; card-top = #131824).
AFTER = f'''<style>
.tgroup{{grid-template-columns:148px minmax(0,1fr)}}.tgroup.ph{{grid-template-columns:minmax(0,1fr)}}
.tchip{{background:color-mix(in srgb,var(--t) 3%,{g.RAISED});border-left:3px solid var(--t);border-right:0}}
.trow{{border-left:0;padding-left:16px}}
.trow::before{{content:"";position:absolute;top:0;bottom:0;left:-148px;width:3px;background:transparent;transition:background-color 120ms ease-out}}
.trow:hover,.trow.is-hover{{background:#131824}}
.trow:hover::before,.trow.is-hover::before{{background:var(--c)}}
.trow.is-you,.trow.is-you:hover{{background:color-mix(in srgb,{g.GOLD} 6%,transparent)}}
.trow.is-you::before,.trow.is-you:hover::before{{background:{g.GOLD}}}
.trow.is-you:hover{{background:color-mix(in srgb,{g.GOLD} 10%,transparent)}}
</style>'''
PHONE_AFTER = ''  # the phone rows are inline-styled; patched below


def build(tag: str, after: bool) -> dict[str, Path]:
    base = g.STYLE
    g.STYLE = base + (AFTER if after else '')
    g.CHIP_W = 148 if after else 132
    out = {}
    out['1440'] = Path(g.board('dps', 'alliance', False, 1440, f'{tag}-1440.html'))
    out['2000'] = Path(g.board('dps', 'alliance', False, 2000, f'{tag}-2000.html'))
    ph = Path(g.board('dps', 'alliance', True, 360, f'{tag}-360.html'))
    if after:
        ph.write_text(ph.read_text().replace(f'border-left:2px solid {g.GOLD}', f'border-left:0;box-shadow:inset 3px 0 0 {g.GOLD}').replace('border-left:2px solid transparent', 'border-left:0'))
    out['360'] = ph
    g.STYLE = base
    g.CHIP_W = 132
    return out


def shoot(p, html: Path, png: Path, w: int, clip=None, hover_you=False) -> None:
    b = p.chromium.launch(executable_path=CHROME)
    pg = b.new_page(viewport={'width': w, 'height': 900})
    pg.goto(html.as_uri()); pg.wait_for_timeout(600)
    if hover_you:
        pg.locator('.trow:not(.is-you)').nth(1).hover()
    pg.screenshot(path=str(png), full_page=True, clip=clip)
    b.close()


if __name__ == '__main__':
    with sync_playwright() as p:
        for tag in ('before', 'after'):
            files = build(tag, tag == 'after')
            shoot(p, files['1440'], HERE / f'{tag}-1440.png', 1440, clip={'x': 0, 'y': 480, 'width': 1440, 'height': 1500})
            shoot(p, files['2000'], HERE / f'{tag}-2000.png', 2000, clip={'x': 0, 'y': 440, 'width': 2000, 'height': 1500})
            shoot(p, files['360'], HERE / f'{tag}-360.png', 360, clip={'x': 0, 'y': 560, 'width': 360, 'height': 700})
        shoot(p, files['1440'], HERE / 'after-1440-hover.png', 1440, clip={'x': 0, 'y': 480, 'width': 1440, 'height': 1500}, hover_you=True)
    # Horde is the faction whose sim result has an A tier today: S, A, B and a you-row-free list.
    with sync_playwright() as p:
        base = g.STYLE
        g.STYLE = base + AFTER
        g.CHIP_W = 148
        h = Path(g.board('dps', 'horde', False, 1440, 'after-horde-1440.html'))
        g.STYLE = base
        shoot(p, h, HERE / 'after-horde-1440.png', 1440, clip={'x': 0, 'y': 480, 'width': 1440, 'height': 1100})

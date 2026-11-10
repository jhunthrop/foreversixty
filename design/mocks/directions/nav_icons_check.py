"""A check sheet for the owner's traced tab icons at their real size: python3 design/mocks/directions/nav_icons_check.py
-> html/nav-icons-check.html (24 px and 48 px, selected gold and muted, on the tab bar's colour), shot by shoot_sheet.cjs.
"""
import json

from assets import owner_asset
from common import write

NAV_ICONS = ('home', 'guild', 'gear', 'activity')
# Guild is the site's own Guild door, not a nav trace (spec 2026-11-10-guild-tab.md section 1).
SOURCES = {'guild': 'icons/ui/guild.svg'}
SELECTED, MUTED, BAR = '#ffd36b', 'rgba(255,255,255,.7)', '#0b2f3b'
# At 24 px the swords' grips and the horn's strap break up; a same-colour stroke of 22 units (viewBox 794, about
# two-thirds of a pixel at 24) holds them, applied at embed time so the owner's files stay as traced. One source with
# the apps' vectors (web/scripts/gen-ui-icons.mjs bakes it in): design/assets/icons/nav/thicken.json.
THICKEN = {k: v for k, v in json.loads(owner_asset('icons/nav/thicken.json').read_text()).items() if not k.startswith('$')}


def nav_svg(name: str, px: int) -> str:
    """The owner's traced icon, tintable (currentColor), sized, thickened where it breaks up at small sizes."""
    body = owner_asset(SOURCES.get(name, f'icons/nav/{name}.svg')).read_text()
    if name in THICKEN:
        body = body.replace('fill="currentColor"', f'fill="currentColor" stroke="currentColor" stroke-width="{THICKEN[name]}" stroke-linejoin="round"')
    return body.replace('<svg ', f'<svg width="{px}" height="{px}" aria-hidden="true" ', 1)


def svg(name: str, px: int, colour: str) -> str:
    return f'<span style="display:inline-block;width:{px}px;height:{px}px;color:{colour}">{nav_svg(name, px)}</span>'


def build() -> None:
    rows = ''.join(f'<div style="display:flex;gap:28px;align-items:center;padding:10px 16px">'
                   + ''.join(svg(n, px, col) for n in NAV_ICONS for col in (SELECTED, MUTED)) + '</div>' for px in (24, 48))
    write('nav-icons-check', f'<!doctype html><meta charset="utf-8"><body style="margin:0;background:{BAR}">{rows}</body>')


if __name__ == '__main__':
    build()

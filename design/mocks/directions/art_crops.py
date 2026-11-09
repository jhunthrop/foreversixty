"""The owner's home art at the three widths, plus a link-preview page (design/specs/2026-11-09-directions.md section 6).

    python3 design/mocks/directions/art_crops.py
    node design/mocks/directions/shoot_sheet.cjs <node_modules> html/og-home.html png/og-home-full.png
    python3 design/mocks/directions/art_crops.py og        -> png/og-home.png (1200 x 630)

The art is generated in the Classic style (the owner's choice for art assets): background only, never captioned or
implied to be a screenshot. The preview carries our seal, the wordmark and the fan-made line, so a shared link can't
read as an official promo (ux-designer), under the page's veil, cropped clear of the smeared right-edge pine.
"""
import base64
import sys
from io import BytesIO
from pathlib import Path

from PIL import Image

from common import write
from assets import logo_uri, owner_asset
from style import FONTS

HERE = Path(__file__).resolve().parent
BAND_HEIGHT = {1440: 312, 2000: 312, 390: 270}  # the home hero band's rendered height at each width
SOCIAL = 'art/social-banner-valley.jpeg'  # the owner's link-preview art (byte-identical to hero-road-dusk.jpeg)
OG_SIZE, OG_BOX = (1200, 630), (1180, 480, 2600, 1225)  # glow over the pass, the road's head
OG_VEIL = 'rgba(7,29,46,.42)'
APP_ICON_PX, APP_SEAL_SHARE = 512, .62  # the flat seal fills 62% of the icon, inside the hammered-gold frame
DISCORD_SIZE = (960, 540)


def app_icon() -> Path:
    """A sample app icon: the owner's blue field with its hammered-gold frame, the flat seal centred inside it."""
    field = Image.open(owner_asset('art/app-icon-background.jpeg')).convert('RGBA').resize((APP_ICON_PX, APP_ICON_PX), Image.LANCZOS)
    s = round(APP_ICON_PX * APP_SEAL_SHARE)
    seal = Image.open(owner_asset('logo/lx-seal-flat-512.png')).convert('RGBA').resize((s, s), Image.LANCZOS)
    field.alpha_composite(seal, ((APP_ICON_PX - s) // 2, (APP_ICON_PX - s) // 2))
    out = HERE / 'png' / 'app-icon.png'
    field.convert('RGB').save(out)
    return out


def discord_banner() -> Path:
    out = HERE / 'png' / 'discord-banner.png'
    Image.open(owner_asset('art/discord-banner-tavern.jpeg')).convert('RGB').resize(DISCORD_SIZE, Image.LANCZOS).save(out)
    return out


def band_crops() -> None:
    for w, h in BAND_HEIGHT.items():
        Image.open(HERE / 'png' / f'g-home-{w}.png').crop((0, 0, w, h)).save(HERE / 'png' / f'art-band-{w}.png')


def og_page() -> Path:
    art = Image.open(owner_asset(SOCIAL)).convert('RGB').crop(OG_BOX).resize(OG_SIZE, Image.LANCZOS)
    buf = BytesIO()
    art.save(buf, 'JPEG', quality=86)
    uri = 'data:image/jpeg;base64,' + base64.b64encode(buf.getvalue()).decode()
    w, h = OG_SIZE
    html = (f'<!doctype html><meta charset="utf-8"><link rel="stylesheet" href="https://fonts.googleapis.com/css2?{FONTS}&display=swap">'
            f'<style>body{{margin:0}}.og{{width:{w}px;height:{h}px;position:relative;background:linear-gradient({OG_VEIL},{OG_VEIL}),url({uri}) 0 0/cover}}'
            f'.mk{{position:absolute;left:56px;bottom:52px;display:flex;align-items:center;gap:22px}}.mk img{{width:120px;height:120px}}'
            f'.mk b{{display:block;font:400 60px/1 Marcellus,serif;color:#f3eee2;text-shadow:0 2px 4px rgba(0,0,0,.6)}}'
            f'.mk span{{display:block;margin-top:10px;font:600 22px "Open Sans",Arial;color:rgba(255,255,255,.86);text-shadow:0 1px 3px rgba(0,0,0,.7)}}</style>'
            f'<div class="og"><div class="mk"><img src="{logo_uri(256)}" alt=""><div><b>Forever Sixty</b><span>Fan-made player tool for WoW: Forever</span></div></div></div>')
    return write('og-home', html)


def og_crop() -> Path:
    out = HERE / 'png' / 'og-home.png'
    Image.open(HERE / 'png' / 'og-home-full.png').crop((0, 0, *OG_SIZE)).save(out)
    (HERE / 'png' / 'og-home-full.png').unlink()
    return out


if __name__ == '__main__':
    if sys.argv[1:] == ['og']:
        print(og_crop())
    else:
        band_crops()
        print(app_icon(), discord_banner())
        print(og_page())

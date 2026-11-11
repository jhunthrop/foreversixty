"""Review sheets for lane BA2: every identity image, and each store's screenshots, on one sheet each.

    python3 design/reviews/2026-11-10-ba2/compose_sheets.py <og-dir>

<og-dir> holds the live link previews (curl https://foreversixty.gg/og/<page>.png). Writes
identity-sheet.png, store-app-store.png and store-play.png beside this file.
"""
import sys
from pathlib import Path

from PIL import Image, ImageDraw, ImageFont

HERE = Path(__file__).resolve().parent
REPO = HERE.parents[2]
DISCORD = REPO / 'design/community/discord/assets'
LISTING = REPO / 'design/listing'
BG = (16, 10, 6)
INK = (227, 223, 217)
GAP = 24
MIN_CELL = 70  # a label never runs into the next cell
FONT = ImageFont.truetype(str(REPO / 'apps/shared/fonts/OpenSans-Regular.ttf'), 18)


def circle(img: Image.Image) -> Image.Image:
    mask = Image.new('L', img.size, 0)
    ImageDraw.Draw(mask).ellipse((0, 0, *img.size), fill=255)
    out = img.convert('RGBA')
    out.putalpha(mask)
    return out


def tile(path: Path, size: int, shape: str = 'square') -> Image.Image:
    img = Image.open(path).convert('RGBA').resize((size, size), Image.LANCZOS)
    return circle(img) if shape == 'circle' else img


def fit(path: Path, width: int) -> Image.Image:
    img = Image.open(path).convert('RGB')
    return img.resize((width, round(img.height * width / img.width)), Image.LANCZOS)


def row(items: list[tuple[Image.Image, str]], min_cell: int = MIN_CELL) -> Image.Image:
    height = max(i.height for i, _ in items) + 40
    width = sum(max(i.width, min_cell) for i, _ in items) + GAP * (len(items) + 1)
    out = Image.new('RGB', (width, height), BG)
    draw = ImageDraw.Draw(out)
    x = GAP
    for img, label in items:
        out.paste(img, (x, 0), img if img.mode == 'RGBA' else None)
        draw.text((x, img.height + 8), label, font=FONT, fill=INK)
        x += max(img.width, min_cell) + GAP
    return out


def stack(rows: list[Image.Image], out: Path) -> None:
    width = max(r.width for r in rows)
    sheet = Image.new('RGB', (width, sum(r.height for r in rows) + GAP * (len(rows) + 1)), BG)
    y = GAP
    for r in rows:
        sheet.paste(r, (0, y))
        y += r.height + GAP
    sheet.save(out, optimize=True)


def discord_sidebar() -> Image.Image:
    """The banner as Discord's channel sidebar shows it: 240x135 at 2x, the server name drawn over
    its top-left on a shade, as the client does."""
    scale = 2
    banner = Image.open(DISCORD / 'banner.jpg').convert('RGB').resize((240 * scale, 135 * scale), Image.LANCZOS)
    shade = Image.new('L', banner.size, 0)
    draw = ImageDraw.Draw(shade)
    for y in range(48 * scale):
        draw.line([(0, y), (banner.width, y)], fill=int(150 * (1 - y / (48 * scale))))
    banner.paste((0, 0, 0), (0, 0, *banner.size), shade)
    name = ImageFont.truetype(str(REPO / 'apps/shared/fonts/OpenSans-Bold.ttf'), 16 * scale)
    ImageDraw.Draw(banner).text((16 * scale, 12 * scale), 'Forever Sixty', font=name, fill=(255, 255, 255))
    return banner


def identity(og: Path) -> None:
    icons = row([
        (tile(DISCORD / 'icon.png', 160, 'circle'), 'Discord server icon'),
        (tile(DISCORD / 'icon.png', 48, 'circle'), 'at 48'),
        (tile(DISCORD / 'avatar.png', 160, 'circle'), 'bot avatar'),
        (tile(DISCORD / 'premium.png', 20), 'Premium role 20'),
        (tile(DISCORD / 'premium.png', 64), 'Premium file 64'),
        (tile(LISTING / 'icons/listing-icon-400.png', 160), 'CurseForge, Wago'),
        (tile(LISTING / 'app-store/icon-1024.png', 160), 'App Store icon'),
        (tile(LISTING / 'play/icon-512.png', 160, 'circle'), 'Play icon, masked'),
    ], min_cell=190)
    crests = row([(tile(p, 64), p.stem.replace('crest-', '')) for p in sorted(DISCORD.glob('crest-*.png'))]
                 + [(tile(DISCORD / f'{f}.png', 64), f) for f in ('alliance', 'horde')])
    discord = row([
        (discord_sidebar(), 'banner in the sidebar, 240x135 at 2x, name drawn by Discord'),
        (tile(DISCORD / 'premium.png', 20).resize((80, 80), Image.NEAREST), 'Premium, 20 px at 4x'),
    ], min_cell=200)
    art = row([(fit(DISCORD / 'splash.jpg', 960), 'Discord invite background (1920x1080); banner.jpg is the same without the name'),
               (fit(LISTING / 'play/feature-graphic.png', 820), 'Play feature graphic (1024x500)')])
    ogs = row([(fit(og / f'{name}.png', 600), f'live link preview /og/{name}.png') for name in ('index', 'bis', 'app')])
    stack([icons, crests, discord, art, ogs], HERE / 'identity-sheet.png')


def store(name: str, width: int) -> None:
    shots = sorted((LISTING / name / 'screenshots').glob('*.jpg'))
    stack([row([(fit(p, width), p.name) for p in shots])], HERE / f'store-{name}.png')


if __name__ == '__main__':
    identity(Path(sys.argv[1]))
    store('app-store', 440)
    store('play', 450)
    print('sheets written')

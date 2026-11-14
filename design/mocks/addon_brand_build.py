"""Captures of the built addon brand skin (2.0.55), for the persona gate.

    python3 design/mocks/addon_brand_build.py

Writes design/reviews/2026-11-14-addon-brand/build/: build-<tab>.png (Me, Gear, Builds, Rotation, Guild) at
device scale 2, tooltip.png, and side-by-side.png / side-by-side-2.png (the shipped 2.0.52 window, the build
and the live site page at the same 1x CSS scale, the colour key under each).

This is addon_brand.py's pipeline (the store fixture's boards, one evening of one made-up character) with
everything the skin draws read from the build instead of from the spec:
- every texture is the committed file the addon loads (addon/ForeverSixty/media: the ground tile, the band
  per tab, the cave, the seal, the filigree, the glyphs, the 128 px crests and the ring), never the site's
  source art;
- every measure is the Lua's own (Theme.SIZES and Hub.LAYOUT, read from Theme.lua and surfaces/Hub.lua):
  the earth's margins, the plates' padding and gap, the header's crest, the strip's seal, wordmark and tab
  padding, the corner size, the ground's repeat;
- every colour and alpha is the token the Lua paints with (Tokens.lua, through addon_tokens).
The layout follows the Lua: the first plate of each column carries the larger top padding, the lead plate
draws the two top corners, the tabs are as wide as their glyph, label and dot slot, the wordmark shows while
the tabs fit beside it, and the item tooltip's block opens on one blank row and its header line.
The client's own fonts are not in the repo: Open Sans and Marcellus (the addon's own media/fonts) draw the
window, and the tooltip's client face is stood in for by Marcellus, labelled on the board.
"""
import base64
import io
import re
import sys
from pathlib import Path

from PIL import Image

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE))
import addon_brand as mock  # noqa: E402  (pins Data.lua to the store captures' commit)
from addon_brand import ROOT, T, marks, ux  # noqa: E402
from addon_tokens import color  # noqa: E402

ADDON = ROOT / 'addon/ForeverSixty'
MEDIA = ADDON / 'media'
OUT = ROOT / 'design/reviews/2026-11-14-addon-brand/build'


def lua_table(path, name):
    """The `key = number` entries of the Lua table assigned to `name` in `path`."""
    with open(ADDON / path, encoding='utf-8') as file:  # the build's own Lua, not the pinned captures' commit
        text = file.read()
    start = re.search(rf'^{re.escape(name)} = \{{', text, re.M)
    if start is None:
        raise RuntimeError(f'{name} not in {path}')
    depth, end = 0, start.end() - 1
    for index in range(end, len(text)):
        depth += {'{': 1, '}': -1}.get(text[index], 0)
        if depth == 0:
            end = index
            break
    return {k: int(v) for k, v in re.findall(r'\b(\w+) = (-?\d+)\b', text[start.end():end])}


S = lua_table('Theme.lua', 'Theme.SIZES')
G = lua_table('surfaces/Hub.lua', 'Hub.LAYOUT')
A = {k: color(f'addon.{k}') for k in ('plate', 'footer', 'zebra', 'rowHover', 'selectedActive', 'tabHover',
                                      'tabActive', 'headerFade', 'headerJoin', 'corner', 'caveFadeFrom',
                                      'caveFadeTo', 'titleRule', 'headerNote', 'tealHover', 'blueHover')}


def tga_png(path, tint=None):
    image = Image.open(path).convert('RGBA')
    if tint is not None:
        ink = Image.new('RGBA', image.size, tint)
        ink.putalpha(image.getchannel('A'))
        image = ink
    buf = io.BytesIO()
    image.save(buf, 'PNG')
    return 'data:image/png;base64,' + base64.b64encode(buf.getvalue()).decode()


def media(name):
    return tga_png(MEDIA / f'{name}.tga')


# --- The build's assets in place of the spec's ---------------------------------------------------------------

def band_image(mood):
    return Image.open(MEDIA / f'band-{mood}.tga').convert('RGB')


def glyph(name, size, tint, extra=''):
    return (f'<span class="gl" style="width:{size}px;height:{size}px;background:{tint};'
            f'-webkit-mask:url({media("icons/" + name)}) center/contain no-repeat;{extra}"></span>')


def crest_hd(cls, size):
    return (f'<span class="crest" style="width:{size}px;height:{size}px"><img src="{media("crests/" + cls)}">'
            f'<span style="position:absolute;inset:0;background:{color("class." + cls)};-webkit-mask:url('
            f'{media("crests/ring")}) center/100% no-repeat"></span></span>')


def build_css():
    edge = S['platePad'] + S['border']
    left_plate = S['windowLeft'] + 2 * edge
    right_plate = S['windowRight'] + 2 * edge
    return f'''
.win {{ background: {T['ground']} url({media('ground')}) 0 0/{S['groundRepeat']}px {S['groundRepeat']}px repeat !important;
  border: 0 !important; box-shadow: 0 8px 28px rgba(0,0,0,.7) !important; }}
.win::after {{ content: ''; position: absolute; inset: 0; pointer-events: none; z-index: 5;
  box-shadow: inset 0 0 0 1px {T['edge']}, inset 0 0 0 2px {T['panelEdge']}; }}
.win .hd {{ position: relative; height: {G['header']}px !important; padding: 0 {S['hubSide']}px !important;
  gap: {G['crestGap']}px !important; border-bottom: 1px solid {A['headerJoin']} !important; box-sizing: border-box;
  background-size: 100% 100% !important; }}
.win .hd.warm {{ border-bottom-color: {T['ruleWarm']} !important; }}
.win .hd::before {{ content: ''; position: absolute; inset: 0; background: var(--veil); }}
.win .hd::after {{ content: ''; position: absolute; left: 0; top: 0; bottom: 0; width: {S['headerFade']}px;
  background: linear-gradient(90deg, {A['headerFade']}, transparent); }}
.win .hd > * {{ position: relative; z-index: 1; }}
.win .hd .tt2, .win .hd .htt {{ font-family: Title, serif !important; font-size: {S['windowTitle']}px !important;
  line-height: 22px !important; color: {T['strong']} !important; }}
.win .hd .m, .win .hd .nt {{ color: {A['headerNote']} !important; font-size: {S['dialogText']}px !important; }}
.win > .x {{ z-index: 2; top: {(G['header'] - S['closeButton']) // 2}px !important; right: {S['hubSide']}px !important;
  width: {S['closeButton']}px !important; height: {S['closeButton']}px !important; border: 0 !important;
  font-size: 0 !important; display: flex; align-items: center; justify-content: center; }}
.win .strip {{ height: {G['strip']}px !important; padding: 0 !important; gap: {G['tabGap']}px !important;
  position: relative; background: {T['navBar']} !important; border-bottom: 1px solid {T['panelEdge']} !important;
  box-sizing: border-box; align-items: stretch; }}
.win .strip .seal {{ flex: none; width: {G['seal']}px; height: {G['seal']}px; margin: auto 0 auto {G['sealLeft']}px;
  background: url({media('seal')}) center/contain no-repeat; }}
.win .strip .wm {{ flex: none; align-self: center; margin: 0 {G['tabsGap']}px 0 {G['wordmarkGap']}px;
  font-family: Title, serif; font-size: {G['wordmark']}px; color: {T['cream']}; white-space: nowrap; }}
.win .strip .tab {{ min-width: 0 !important; flex: none; justify-content: flex-start !important;
  padding: 0 {S['tabPad'] + G['dot'] + G['dotGap']}px 0 {S['tabPad']}px !important; gap: {S['tabGlyphGap']}px;
  color: {T['tabOff']} !important; font-weight: 600 !important; font-size: 12px;
  border-bottom: {S['tabUnderline']}px solid transparent !important; margin-bottom: -1px; position: relative; }}
.win .strip .tab .dot {{ position: absolute; right: {S['tabPad']}px; margin: 0 !important;
  background: {T['gold']} !important; width: {G['dot']}px; height: {G['dot']}px; }}
.win .strip .tab.on {{ color: {T['strong']} !important; border-color: {T['accent']} !important;
  background: {A['tabActive']}; }}
.win .bd {{ padding: {S['hubMargin']}px {S['hubSide']}px !important; gap: {S['plateGap']}px; align-items: stretch; }}
.win .vr {{ display: none !important; }}
.win .lc {{ width: {left_plate}px !important; margin: 0 !important; display: flex; flex-direction: column;
  gap: {S['plateGap']}px; }}
.win .rc {{ width: {right_plate}px !important; margin: 0 !important; display: flex; flex-direction: column;
  gap: {S['plateGap']}px; }}
.win .bd.one > .pl {{ margin-bottom: {S['plateGap']}px; }}
.pl {{ position: relative; background: {A['plate']}; border: 1px solid {T['panelEdge']};
  padding: {S['platePad']}px !important; margin: 0 !important; box-sizing: border-box; }}
.lc > .pl:first-child, .rc > .pl:first-child, .bd.one > .pl:first-child {{ padding-top: {S['platePadFirst']}px !important; }}
.lc > .pl:last-child, .rc > .pl:last-child {{ flex: 1; }}
.fc {{ position: absolute; width: {S['plateCorner']}px; height: {S['plateCorner']}px; background: {A['corner']};
  pointer-events: none; -webkit-mask: url({media('filigree')}) center/contain no-repeat; z-index: 2; }}
.fc.tl {{ top: {S['plateCornerInset']}px; left: {S['plateCornerInset']}px; }}
.fc.tr {{ top: {S['plateCornerInset']}px; right: {S['plateCornerInset']}px; transform: scaleX(-1); }}
.fc.br, .fc.bl {{ display: none; }}
.pl.cave {{ background: linear-gradient(90deg, {A['caveFadeFrom']}, {A['caveFadeTo']}),
  linear-gradient({T['veilCave']}, {T['veilCave']}), var(--cave) center/100% 100%; }}
.pl .sh {{ height: auto !important; padding-bottom: 4px; margin-bottom: 4px !important;
  border-bottom: 1px solid {A['titleRule']}; align-items: baseline !important; }}
.pl .sh > span:first-child, .pl .sh > span:first-child span {{ font-family: Title, serif !important;
  font-size: {S['windowSection']}px !important; font-weight: 400 !important; color: {T['strong']} !important;
  line-height: 19px; white-space: nowrap; }}
.pl .sh .r {{ color: {T['muted']} !important; }}
.win .btn {{ background: {T['blue']} !important; border-color: {T['blueEdge']} !important; color: {T['strong']} !important;
  font-weight: 700 !important; }}
.win .btn.sel, .win .btn.armed {{ background: {T['teal']} !important; border-color: {T['tealEdge']} !important; }}
.win .link, .win .lk {{ color: {T['accent']} !important; }}
.win .ft {{ background: {A['footer']} !important; border-top: 1px solid {T['panelEdge']} !important;
  padding: 0 {S['hubSide']}px !important; }}
.win .ft span {{ opacity: 1 !important; color: {T['muted']}; }}
.gl {{ display: inline-block; flex: none; }}
.zebra > .r:nth-of-type(even) {{ background: {A['zebra']}; }}
.win .brow.followed::before {{ background: {A['selectedActive']} !important; right: 0 !important; }}
.win .brow.followed::after {{ background: {T['accent']} !important; }}
.win .brow .fol {{ color: {T['accent']} !important; }}
.pl > .sec {{ margin-top: 0 !important; }}
.pulse-bar {{ position: relative; }}
.pulse-bar::before {{ content: ''; position: absolute; left: -{S['platePad']}px; top: 0; bottom: 0;
  width: {S['tabUnderline']}px; background: {T['accent']}; opacity: .6; }}
.trk, .part .tr {{ background: {T['lineSoft']} !important; }}
.trk i, .part .tr i {{ background: {T['accent']} !important; }}
'''


def tooltip_board():
    """The item tooltip as ItemTooltip.drawEntry adds to it: the client's lines, one blank row, our block's
    header line, the verdict in the gain colour (the worn item's name in its quality), the attribution in
    the dim grey and the Alt-click hint in the client gold; and a player's tooltip as PlayerTooltip adds to
    it: the crest line, then the score in its tier's colour."""
    ch = marks.velisandre()
    hands = ch.doll()['hands']
    bag, worn, gain = hands['pick']['item_id'], hands['worn'], hands['gain']
    g, w = ux.item(bag), ux.item(worn)
    l, esc = ux.l, marks.esc
    shipped = ux.tooltip(marks.item_lines(g) + [
        '<div class="gap"></div>',
        l(f'{ux.arrow()}Upgrade +{gain} DPS over {esc(w["name"])}', color=ux.C['green']),
        marks.dim('Forever Sixty · Beast Mastery, your sim'), marks.dim('Alt-click: open Gear')], width=340)
    name_w = f'<span style="color:{ux.QUALITY[w["quality"]]}">{esc(w["name"])}</span>'
    built = ux.tooltip(marks.item_lines(g) + ux.item_block([
        l(f'{ux.inline_arrow()}Upgrade +{gain} DPS over {name_w}', color=ux.TOOLTIP_BLOCK['gain']),
        l('Beast Mastery, your sim', color=ux.TOOLTIP_BLOCK['dim']),
        l('Alt-click: open Gear', color=ux.TOOLTIP_BLOCK['next'])]), width=340)
    player = ux.tooltip([
        f'<div class="t" style="color:{color("class.hunter")}">Velisandre</div>', l('Thornwake Company'),
        l('Level 60 Night Elf Hunter'), ux.player_header('hunter'),
        l('Forever Score 1184', color=color('rarity.rareText'))], width=240)
    wrap = '<style>.tt .l, .tt .l span { white-space: normal !important; }</style>'
    return ux.page(f'{wrap}<div class="row" style="gap:40px;padding-top:22px">'
                   f'<div class="col"><div class="note"><b>2.0.52</b></div>{shipped}</div>'
                   f'<div class="col"><div class="note"><b>2.0.55</b> · Shift held, Hands in the bags</div>{built}</div>'
                   f'<div class="col"><div class="note"><b>2.0.55</b> · a player\'s tooltip</div>{player}</div>'
                   f'<div class="col"><div class="note" style="max-width:260px">The client draws tooltips in its own '
                   f'face, which is not in the repo; Marcellus stands in for it on both.</div></div></div>')


def install():
    """Point addon_brand's renderer at the build (once; a second call changes nothing)."""
    if mock.OUT == OUT:
        return
    mock.OUT = OUT
    mock.band_image = band_image
    mock.glyph = glyph
    mock.crest_hd = crest_hd
    mock.skin_css = build_css
    mock.tooltip_board = tooltip_board
    mock.SEAL = media('seal')
    mock.FILIGREE = media('filigree')
    real_chrome = mock.chrome
    mock.chrome = lambda *a, **k: chrome_with(real_chrome, *a, **k)


def chrome_with(real, html, tab, cls='hunter', lead='rc', cave=False):
    html = real(html, tab, cls, lead, cave)
    html = html.replace('<span class="seal"></span>', '<span class="seal"></span><span class="wm">Forever Sixty</span>', 1)
    if cave:
        html = re.sub(r'--cave:url\([^)]*\)', f'--cave:url({media("cave")})', html, count=1)
    return html


def main():
    install()
    OUT.mkdir(parents=True, exist_ok=True)
    for key in mock.TABS:
        old = OUT / f'brand-{key}.png'
        if old.exists():
            old.unlink()
    mock.main()
    for key in mock.TABS:
        (OUT / f'brand-{key}.png').rename(OUT / f'build-{key}.png')
    for mood in mock.BAND_CROPS:
        (OUT / f'band-{mood}.png').unlink(missing_ok=True)


if __name__ == '__main__':
    main()

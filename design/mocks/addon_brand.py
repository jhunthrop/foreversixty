"""Mock of the addon brand skin (design/specs/2026-11-14-addon-brand.md).

    python3 design/mocks/addon_brand.py

Writes design/reviews/2026-11-14-addon-brand/mock/:
- brand-<tab>.png: the hub on Me, Gear, Builds, Rotation and Guild in the spec's skin, at the real 740 px
  width, device scale 2.
- tooltip.png: the shipped item tooltip beside the branded one.
- side-by-side.png: per tab, the shipped window, the mock and the live site page, all at 1x CSS scale (one
  CSS px is one image px), so the three read at the same scale.
- band-<mood>.png: the five header band crops (BAND_CROPS below; the build lane freezes these).

Nothing here draws content of its own. Every board is the store fixture's own (design/listing/gen_captures.py:
the lane generators the 2.0.48 frames were reviewed against, one evening of one made-up character); this
script swaps the chrome, exactly as the build lane will: ground, header band, nav strip, plates, rows,
buttons, icons and crests. Every colour comes from design/tokens/tokens.json through addon_tokens, the new
addon tokens included (ADDON_NEW below lists the values the spec adds to the addon group; the build lane
writes them to tokens.json). The art and icons are the site's own: web/public/art, web/public/icons/hd/crests,
design/assets/icons/ui, design/assets/ornaments.

The lane generators read the addon's Lua (Data.lua's packed picks, MeFrame.lua's layout), which has moved
on since the store captures were made, so this script reads it as of lua_pin.CAPTURE_COMMIT (the store
fixture's commit): the shipped window and the mock then show the same evening.
"""
import base64
import io
import os
import re
import sys
from pathlib import Path

from PIL import Image

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
OUT = ROOT / 'design/reviews/2026-11-14-addon-brand/mock'
SITE = ROOT / 'design/reviews/2026-11-14-addon-brand/site'
SHIPPED = ROOT / 'design/listing/captures'

sys.path.insert(0, str(HERE))
sys.path.insert(0, str(ROOT / 'design/listing'))
import lua_pin  # noqa: E402,F401  (pins the lanes' Lua to the store captures' commit)
import addon_hub_l10 as marks  # noqa: E402
import addon_ux as ux  # noqa: E402
import gen_captures as store  # noqa: E402
from addon_tokens import color  # noqa: E402

# --- Tokens -------------------------------------------------------------------------------------------------

T = {k: color(v) for k, v in {
    'ground': 'forever.ground', 'band': 'forever.band', 'navBar': 'forever.navBar', 'accent': 'forever.accent',
    'panelEdge': 'forever.panelEdge', 'veilRule': 'forever.veilRule', 'veilBand': 'forever.veilBand',
    'veilNight': 'forever.veilNight', 'veilCave': 'forever.veilCave', 'ruleWarm': 'forever.ruleWarm',
    'tabOff': 'forever.tabOff', 'inkSoft': 'forever.inkSoft', 'foot': 'forever.foot', 'edge': 'forever.edge',
    'teal': 'forever.teal', 'tealEdge': 'forever.tealEdge', 'blue': 'forever.blue', 'blueEdge': 'forever.blueEdge',
    'gain': 'forever.gain', 'warn': 'forever.warn', 'tooltipYellow': 'forever.tooltipYellow', 'seal': 'forever.seal',
    'frameLine': 'forever.frameLine', 'cream': 'forever.cream',
    'raised': 'base.raised', 'hover': 'base.hover', 'pressed': 'base.pressed', 'selected': 'base.selected',
    'line': 'base.line', 'lineSoft': 'base.lineSoft', 'text': 'base.text', 'strong': 'base.strong',
    'muted': 'base.muted', 'gold': 'base.gold', 'inset': 'addon.inset',
    'tooltipMuted': 'addon.tooltipMuted', 'tooltipText': 'addon.tooltipText',
}.items()}
# The spec's new addon tokens (section 8), until the build lane adds them to tokens.json.
ADDON_NEW = {'plate': 'rgba(23, 14, 9, 0.84)', 'zebra': 'rgba(255, 255, 255, 0.03)',
             'tabHover': 'rgba(255, 255, 255, 0.06)', 'selectedActive': 'rgba(0, 147, 178, 0.24)',
             'tealHover': '#12727a', 'blueHover': '#17608a', 'headerFade': 'rgba(11, 7, 5, 0.35)'}

# --- Art: the site's files, cut the way the generator will cut them -------------------------------------------

ART = ROOT / 'web/public/art'
# mood -> (source, crop box as fractions (left, top, right, bottom), mirrored). Exported 512 x 64, drawn 740 x 56.
BAND_CROPS = {
    'me': ('band-2400.webp', (0.0, 0.30, 1.0, 0.52), False),
    'builds': ('band-2400.webp', (0.0, 0.48, 1.0, 0.70), True),
    'gear': ('paperdoll-frame@2x.webp', (0.22, 0.42, 0.78, 0.50), False),
    'rotation': ('band-2400.webp', (0.0, 0.62, 1.0, 0.84), False),
    'cave': ('cave-1600.webp', (0.3, 0.12, 1.0, 0.40), False),
}
TAB_BAND = {'Me': 'me', 'Builds': 'builds', 'Gear': 'gear', 'Rotation': 'rotation', 'Guild': 'cave',
            'Officer': 'cave'}
VEIL = {'me': 'veilBand', 'builds': 'veilBand', 'gear': 'veilNight', 'rotation': 'veilBand', 'cave': 'veilNight'}
TAB_ICON = {'Me': 'character', 'Builds': 'planner', 'Gear': 'bis', 'Rotation': 'sim', 'Guild': 'guild',
            'Officer': 'manage'}


def png_uri(im):
    buf = io.BytesIO()
    im.save(buf, 'PNG')
    return 'data:image/png;base64,' + base64.b64encode(buf.getvalue()).decode()


def band_image(mood):
    src, (l, t, r, b), mirror = BAND_CROPS[mood]
    im = Image.open(ART / src).convert('RGB')
    w, h = im.size
    im = im.crop((int(l * w), int(t * h), int(r * w), int(b * h))).resize((512, 64), Image.LANCZOS)
    return im.transpose(Image.FLIP_LEFT_RIGHT) if mirror else im


def ground_tile():
    """ground.webp to the addon's 256 px tile: a 640 px window of the 1600 px tile, so the grain keeps the
    site's scale at the addon's size (the whole tile at 256 px reads as noise)."""
    return Image.open(ART / 'ground.webp').convert('RGB').crop((0, 0, 640, 640)).resize((256, 256), Image.LANCZOS)


def svg_uri(svg):
    return 'data:image/svg+xml;base64,' + base64.b64encode(svg.encode()).decode()


def ui_icon(name):
    for path in (ROOT / f'design/assets/icons/ui/{name}.svg', ROOT / f'design/assets/icons/ui/line/{name}.svg'):
        if path.exists():
            return svg_uri(path.read_text())
    raise FileNotFoundError(name)


def glyph(name, size, tint, extra=''):
    """A white generator glyph tinted by SetVertexColor: a mask in the token colour."""
    return (f'<span class="gl" style="width:{size}px;height:{size}px;background:{tint};'
            f'-webkit-mask:url({ui_icon(name)}) center/contain no-repeat;{extra}"></span>')


SEAL = svg_uri((ROOT / 'web/public/art/seal-flat-small.svg').read_text())
FILIGREE = svg_uri((ROOT / 'design/assets/ornaments/corner-filigree.svg').read_text())
RING = ux.tga_uri(ROOT / 'addon/ForeverSixty/media/crests/ring.tga')


def crest_hd(cls, size):
    face = 'data:image/webp;base64,' + base64.b64encode(
        (ROOT / f'web/public/icons/hd/crests/{cls}.webp').read_bytes()).decode()
    return (f'<span class="crest" style="width:{size}px;height:{size}px"><img src="{face}">'
            f'<span style="position:absolute;inset:0;background:{color("class." + cls)};-webkit-mask:url({RING}) '
            f'center/100% no-repeat"></span></span>')


# --- The skin ---------------------------------------------------------------------------------------------------

def skin_css():
    ground = png_uri(ground_tile())
    return f'''
.win {{ background: {T['ground']} url({ground}) 0 0/256px 256px repeat !important; border: 1px solid {T['edge']} !important;
  box-shadow: inset 0 0 0 1px {T['panelEdge']}, 0 8px 28px rgba(0,0,0,.7) !important; }}
.win .hd {{ position: relative; height: 56px !important; padding: 0 12px !important; gap: 10px !important;
  border-bottom: 1px solid rgba(0,147,178,.5) !important; background-size: 100% 100% !important; }}
.win .hd.warm {{ border-bottom-color: {T['ruleWarm']} !important; }}
.win .hd::before {{ content: ''; position: absolute; inset: 0; background: var(--veil); }}
.win .hd::after {{ content: ''; position: absolute; left: 0; top: 0; bottom: 0; width: 260px;
  background: linear-gradient(90deg, {ADDON_NEW['headerFade']}, transparent); }}
.win .hd > * {{ position: relative; z-index: 1; }}
.win .hd .tt2, .win .hd .htt {{ font-family: Title, serif !important; font-size: 18px !important; line-height: 22px !important;
  color: {T['strong']} !important; }}
.win .hd .m, .win .hd .nt {{ color: {T['inkSoft']} !important; font-size: 11px !important; }}
.win > .x {{ z-index: 2; top: 17px !important; right: 10px !important; width: 22px !important; height: 22px !important;
  border: 0 !important; font-size: 0 !important; display: flex; align-items: center; justify-content: center; }}
.win .strip {{ height: 28px !important; padding: 0 0 0 36px !important; gap: 0 !important; position: relative;
  background: {T['navBar']} !important; border-bottom: 1px solid {T['panelEdge']} !important; }}
.win .strip .seal {{ position: absolute; left: 8px; top: 4px; width: 20px; height: 20px;
  background: url({SEAL}) center/contain no-repeat; }}
.win .strip .tab {{ min-width: 0 !important; width: 117px; flex: none; padding: 0 !important; justify-content: center; gap: 6px; color: {T['tabOff']} !important;
  font-weight: 600; border-bottom: 2px solid transparent !important; margin-bottom: -1px; }}
.win .strip .tab.on {{ color: {T['strong']} !important; font-weight: 700; border-color: {T['accent']} !important;
  background: rgba(5, 69, 98, .55); }}
.win .strip .tab.hov {{ color: {T['strong']} !important; background: {ADDON_NEW['tabHover']}; }}
.win .strip .tab .dot {{ background: {T['gold']} !important; }}
.win .strip .tab .dot.pulse {{ box-shadow: 0 0 0 3px rgba(255,236,113,.25); }}
.win .bd {{ padding: 6px 10px !important; gap: 6px; align-items: stretch; }}
.win .vr {{ display: none !important; }}
.win .lc {{ width: 405px !important; margin: 0 !important; display: flex; flex-direction: column; gap: 6px; }}
.win .rc {{ width: 307px !important; margin: 0 !important; display: flex; flex-direction: column; gap: 6px; }}
.win .bd.one > .pl {{ margin-bottom: 6px; }}
.pl {{ position: relative; background: {ADDON_NEW['plate']}; border: 1px solid {T['panelEdge']}; padding: 8px !important;
  margin: 0 !important; }}
.lc > .pl:last-child, .rc > .pl:last-child {{ flex: 1; }}
.fc {{ position: absolute; width: 28px; height: 28px; opacity: .8; background: {T['cream']}; pointer-events: none;
  -webkit-mask: url({FILIGREE}) center/contain no-repeat; z-index: 2; }}
.fc.tl {{ top: 3px; left: 3px; }}
.fc.tr {{ top: 3px; right: 3px; transform: scaleX(-1); }}
.fc.br {{ bottom: 3px; right: 3px; transform: scale(-1, -1); }}
.fc.bl {{ bottom: 3px; left: 3px; transform: scaleY(-1); }}
.pl.lead {{ padding: 15px 8px 8px 8px !important; }}
.pl.cave {{ background: linear-gradient(90deg, rgba(23,14,9,.9) 40%, rgba(23,14,9,.5)),
  linear-gradient({T['veilCave']}, {T['veilCave']}), var(--cave) 70% 40%/cover; }}
.pl .sh {{ height: auto !important; padding-bottom: 4px; margin-bottom: 4px !important; border-bottom: 1px solid {T['veilRule']};
  align-items: baseline !important; }}
.pl .sh > span:first-child, .pl .sh > span:first-child span {{ font-family: Title, serif !important; font-size: 15px !important;
  font-weight: 400 !important; color: {T['strong']} !important; line-height: 19px; }}
.pl .sh .r {{ font-size: 11px; color: {T['muted']} !important; }}
.win .btn {{ background: {T['blue']} !important; border-color: {T['blueEdge']} !important; color: {T['strong']} !important;
  font-weight: 700 !important; }}
.win .btn.sel, .win .btn.armed {{ background: {T['teal']} !important; border-color: {T['tealEdge']} !important; }}
.win .btn.hov {{ background: {ADDON_NEW['blueHover']} !important; }}
.win .link, .win .lk {{ color: {T['accent']} !important; }}
.win .g, .win .cg {{ color: {T['gold']}; }}
.win .ft {{ background: rgba(23,14,9,.85) !important; border-top: 1px solid {T['panelEdge']} !important; }}
.win .ft span {{ opacity: 1 !important; color: {T['muted']}; font-size: 11px; }}
.gl {{ display: inline-block; flex: none; }}
.zebra > :nth-child(odd) {{ background: {ADDON_NEW['zebra']}; }}
.hovrow {{ background: rgba(36,27,21,.9) !important; }}
.selrow {{ background: {T['selected']} !important; box-shadow: inset 2px 0 0 {T['accent']} !important;
  border-left: 0 !important; }}
.selrow.active {{ background: {ADDON_NEW['selectedActive']} !important; }}
.win .brow.followed::before {{ background: {ADDON_NEW['selectedActive']} !important; right: 0 !important; }}
.win .brow.followed::after {{ background: {T['accent']} !important; }}
.win .brow .fol {{ color: {T['accent']} !important; }}
.pl > .sec {{ margin-top: 0 !important; }}
.pulse-bar {{ box-shadow: inset 2px 0 0 {T['accent']}; padding-left: 9px; margin-left: -9px; }}
.chev {{ margin-left: 4px; vertical-align: -1px; }}
.trk, .part .tr {{ background: {T['lineSoft']} !important; box-shadow: inset 0 1px 0 {T['inset']}; }}
.trk i, .trk > span, .part .tr i {{ box-shadow: inset 0 1px 0 rgba(243,238,226,.21); }}
'''


# --- Plates: each column becomes a stack of plates ---------------------------------------------------------------

VOID = {'img', 'br', 'input', 'hr', 'col', 'wbr'}
TAG = re.compile(r'<(/?)([a-zA-Z0-9]+)([^>]*?)(/?)>')


def children(html):
    """The top-level elements of an HTML fragment, as strings."""
    out, depth, start = [], 0, None
    for m in TAG.finditer(html):
        closing, name = m.group(1), m.group(2).lower()
        if name in VOID or m.group(4):
            if depth == 0:
                out.append(m.group(0))
            continue
        if not closing:
            if depth == 0:
                start = m.start()
            depth += 1
        else:
            depth -= 1
            if depth == 0:
                out.append(html[start:m.end()])
    return out


def element(html, cls):
    """(start, end) of the first element whose class list starts with `cls`, and its inner span."""
    m = re.search(rf'<div class="{cls}[" ]', html)
    if m is None:
        raise RuntimeError(f'no .{cls}')
    depth = 0
    for t in TAG.finditer(html, m.start()):
        if t.group(2).lower() in VOID or t.group(4):
            continue
        depth += -1 if t.group(1) else 1
        if depth == 0:
            open_end = html.index('>', m.start()) + 1
            return m.start(), t.end(), open_end, t.start()
    raise RuntimeError(f'.{cls} not closed')


def is_break(child):
    return ('class="gap"' in child or re.match(r'<div style="height:\s*8px">\s*</div>', child)
            or re.match(r'<div[^>]*style="height:8px"[^>]*>\s*</div>', child))


def starts_plate(child):
    return re.match(r'<div class="(sec|spec)\b', child) is not None


def plates(inner, lead=False, extra=''):
    kids = children(inner)
    wrapper = ''
    if len(kids) == 1 and re.match(r'<div class="me"', kids[0]):
        wrapper = 'me '
        kids = children(kids[0][kids[0].index('>') + 1:kids[0].rindex('</div>')])
    groups, cur = [], []
    for k in kids:
        if is_break(k):
            groups.append(cur)
            cur = []
            continue
        if starts_plate(k) and cur:
            groups.append(cur)
            cur = []
        cur.append(k)
    groups.append(cur)
    groups = [g for g in groups if g]
    out = []
    for i, g in enumerate(groups):
        cls = f'{wrapper}pl' + (' lead' if lead and i == 0 else '') + (f' {extra}' if extra and i == 0 else '')
        corners = ''.join(f'<span class="fc {c}"></span>' for c in ('tl', 'tr')) if 'lead' in cls else ''
        out.append(f'<div class="{cls}">{corners}{"".join(g)}</div>')
    return ''.join(out)


def column(html, cls, **kw):
    s, e, os_, ie = element(html, cls)
    return html[:os_] + plates(html[os_:ie], **kw) + html[ie:]


def chrome(html, tab, cls='hunter', lead='rc', cave=False):
    band = band_image(TAB_BAND[tab])
    veil = T[VEIL[TAB_BAND[tab]]]
    # Officer shares Guild's cave band; its header join is the warm rule (Hub.BANDS officer.join).
    hd = 'hd warm' if tab == 'Officer' else 'hd'
    html = html.replace('<div class="hd">', f'<div class="{hd}" style="background:url({png_uri(band)});--veil:{veil}">', 1)
    html = re.sub(r'<span class="crest" style="width:32px;height:32px">.*?</span></span>', crest_hd(cls, 40), html,
                  count=1, flags=re.S)
    html = html.replace('<div class="x">×</div>', f'<div class="x">{glyph("close", 12, T["tabOff"])}</div>', 1)
    html = html.replace('<div class="strip">', '<div class="strip"><span class="seal"></span>', 1)
    for name, icon_name in TAB_ICON.items():
        html = re.sub(rf'(<div class="tab( on)?">){name}(?=<|$)',
                      lambda m, n=name, i=icon_name: f'{m.group(1)}{glyph(i, 14, "currentColor")}{n}', html, count=1)
    if '<div class="lc">' in html:
        html = column(html, 'lc', lead=lead == 'lc', extra='cave' if cave and lead == 'lc' else '')
        html = column(html, 'rc', lead=lead == 'rc', extra='cave' if cave and lead == 'rc' else '')
    else:
        s, e, os_, ie = element(html, 'bd one')
        html = html[:os_] + plates(html[os_:ie], lead=True) + html[ie:]
    if cave:
        cave = Image.open(ART / 'cave-1600.webp').convert('RGB')
        cave_uri = png_uri(cave.crop((480, 60, 1600, 620)).resize((256, 128), Image.LANCZOS))
        html = html.replace('<div class="win"', f'<div class="win" style="--cave:url({cave_uri})"', 1)
    return html


# --- The tabs -----------------------------------------------------------------------------------------------------

def me():
    html = chrome(store.me_raid_day(), 'Me', lead='rc')
    html = html.replace('<div class="tb">', '<div class="tb zebra">', 1)
    # The next action: the first Right now row carries the accent bar (the pulse's steady frame).
    html = html.replace('<div class="now">', '<div class="now pulse-bar">', 1)
    html = html.replace('<span class="dot"></span>', '<span class="dot pulse"></span>', 1)
    html = re.sub(r'(<span class="lk link"[^>]*>)([^<]*)</span>',
                  lambda m: f'{m.group(1)}{m.group(2)}{glyph("chevron-right", 10, T["accent"], "display:inline-block;margin-left:2px")}</span>',
                  html)
    return html


def gear():
    html = chrome(store.gear_raid_ready(), 'Gear', lead='lc')
    gold_tick = f'background:{color("base.gold")};display:inline-flex;align-items:center;justify-content:center;color:{color("base.bg")}'
    if gold_tick not in html:
        raise RuntimeError('the Sources tick moved')
    html = html.replace(gold_tick, f'background:{T["accent"]};display:inline-flex;align-items:center;'
                                   f'justify-content:center;color:{T["strong"]}')
    # One hovered slot cell and the selected one (Hands, the bag upgrade the tooltip below is about).
    cells = [m.start() for m in re.finditer(r'<div class="cell r"', html)]
    html = html[:cells[0]] + '<div class="cell r selrow active"' + html[cells[0] + len('<div class="cell r"'):]
    first = html.index('<div class="up">')
    html = html[:first] + '<div class="up hovrow">' + html[first + len('<div class="up">'):]
    return html


def builds():
    html = chrome(store.builds_60(), 'Builds', lead='lc')
    return html


def rotation():
    return chrome(store.rotation_60(), 'Rotation', lead='lc')


def guild():
    return chrome(store.guild_signed_in(), 'Guild', lead='lc', cave=True)


def officer():
    """The raid day: the roster is the lead plate over the cave art (OfficerTab.build's plateLooks)."""
    return chrome(store.officer_raid_day(), 'Officer', lead='lc', cave=True)


TABS = {'me': me, 'gear': gear, 'builds': builds, 'rotation': rotation, 'guild': guild, 'officer': officer}
SITE_PAGE = {'me': 'char', 'gear': 'bis', 'builds': 'guide', 'rotation': 'guide', 'guild': 'home'}
SHIPPED_FILE = {'me': 'window-me', 'gear': 'window-gear', 'builds': 'window-builds', 'rotation': 'window-rotation',
                'guild': 'window-guild'}

# --- The item tooltip ---------------------------------------------------------------------------------------------


def tooltip_board():
    ch = marks.velisandre()
    hands = ch.doll()['hands']
    bag, worn, gain = hands['pick']['item_id'], hands['worn'], hands['gain']
    g, w = ux.item(bag), ux.item(worn)
    l, esc = ux.l, marks.esc
    shipped = ux.tooltip(marks.item_lines(g) + [
        '<div class="gap"></div>',
        l(f'{ux.arrow()}Upgrade +{gain} DPS over {esc(w["name"])}', color=ux.C['green']),
        marks.dim('Forever Sixty · Beast Mastery, your sim'), marks.dim('Alt-click: open Gear')], width=340)
    seal = f'<img src="{SEAL}" style="width:14px;height:14px;vertical-align:-2px;margin-right:4px">'
    name_w = f'<span style="color:{ux.QUALITY[w["quality"]]}">{esc(w["name"])}</span>'
    branded = ux.tooltip(marks.item_lines(g) + [
        '<div class="gap"></div>',
        l(f'{seal}Forever Sixty', color=T['accent']),
        l(f'{ux.arrow().replace("#35e04a", T["gain"]).replace("#b6ffbf", T["cream"])}'
          f'<span style="color:{T["tooltipText"]}">Upgrade</span> '
          f'<span style="color:{T["gain"]}">+{gain} DPS</span> '
          f'<span style="color:{T["tooltipMuted"]}">over</span> {name_w}'),
        l(f'<span style="color:{T["tooltipMuted"]}">Beast Mastery, your sim</span>'),
        l('Alt-click: open Gear', color=T['tooltipYellow'])], width=340)
    wrap = '<style>.tt .l, .tt .l span { white-space: normal !important; }</style>'
    return ux.page(f'{wrap}<div class="row" style="gap:40px;padding-top:22px">'
                   f'<div class="col"><div class="note"><b>Shipped</b></div>{shipped}</div>'
                   f'<div class="col"><div class="note"><b>Branded</b> · the native frame, our block</div>{branded}</div>'
                   f'<div class="col"><div class="note" style="max-width:260px">Both draw the client\'s own tooltip face; '
                   f'the mock has no copy of it, so Marcellus stands in. In game the lines are the client font.</div></div>'
                   f'</div>')


# --- The colour key under each sheet (spec 14.5) ------------------------------------------------------------------

KEY = [
    ('Window ground (tile, 10 px side margins, 6 px gaps)', 'forever.ground', T['ground']),
    ('Header band veil (Me, Builds, Rotation)', 'forever.veilBand', T['veilBand']),
    ('Header band veil (Gear, Guild, Officer)', 'forever.veilNight', T['veilNight']),
    ('Header left fade', 'addon.headerFade (new)', ADDON_NEW['headerFade']),
    ('Header join line', 'forever.accent at 0.5', T['accent']),
    ('Tab strip plate', 'forever.navBar', T['navBar']),
    ('Active tab plate', 'forever.band at 0.55', T['band']),
    ('Active tab underline, selected bar, links, ticks', 'forever.accent', T['accent']),
    ('Tab label at rest', 'forever.tabOff', T['tabOff']),
    ('Waiting dot, score', 'base.gold', T['gold']),
    ('Plates', 'addon.plate (new)', ADDON_NEW['plate']),
    ('Plate edge', 'forever.panelEdge', T['panelEdge']),
    ('Title rule', 'forever.veilRule', T['veilRule']),
    ('Lead plate corners', 'forever.cream at 0.8', T['cream']),
    ('Zebra rows', 'addon.zebra (new)', ADDON_NEW['zebra']),
    ('Hovered row', 'base.hover at 0.9', T['hover']),
    ('Selected row', 'base.selected', T['selected']),
    ('Selected row, detail open', 'addon.selectedActive (new)', ADDON_NEW['selectedActive']),
    ('Primary button', 'forever.teal / tealEdge', T['teal']),
    ('Secondary button', 'forever.blue / blueEdge', T['blue']),
    ('Progress track', 'base.lineSoft', T['lineSoft']),
    ('Footer plate', 'forever.foot at 0.85', T['foot']),
    ('Tooltip header line', 'forever.accent', T['accent']),
    ('Tooltip gain', 'forever.gain', T['gain']),
    ('Tooltip next step', 'forever.tooltipYellow', T['tooltipYellow']),
]


def key_html():
    rows = ''.join(f'<div class="k"><span class="sw" style="background:{hex_}"></span><span class="rg">{region}</span>'
                   f'<span class="tk">{token}</span></div>' for region, token, hex_ in KEY)
    css = (f'#board {{ background: {T["raised"]} !important; padding: 20px 24px !important; }} '
           '.keys { display: grid; grid-template-columns: repeat(3, 560px); column-gap: 32px; row-gap: 6px; '
           f'font-family: Site, sans-serif; font-size: 13px; color: {T["text"]}; }} '
           '.k { display: flex; gap: 10px; align-items: center; } '
           f'.sw {{ width: 22px; height: 16px; flex: none; box-shadow: 0 0 0 1px {T["line"]}; }} '
           '.rg { width: 330px; } '
           f'.tk {{ color: {T["muted"]}; font-family: monospace; font-size: 12px; }} '
           f'.kt {{ font-family: Title, serif; font-size: 18px; color: {T["strong"]}; margin-bottom: 12px; }}')
    return ux.page(f'<style>{css}</style><div class="kt">Colour key: every region and its token</div>'
                   f'<div class="keys">{rows}</div>')


# --- Rendering ----------------------------------------------------------------------------------------------------

def shoot(browser, html, path, selector, scale):
    ctx = browser.new_context(device_scale_factor=scale, viewport={'width': 1800, 'height': 1400})
    page = ctx.new_page()
    page.set_content(html.replace('</style>', store.HIDE_NOTES + skin_css() + '</style>', 1))
    page.evaluate('document.fonts.ready')
    page.wait_for_timeout(200)
    box = page.evaluate(store.BOX_JS, selector)
    page.screenshot(path=str(path), clip=box, full_page=True)
    ctx.close()
    return box


def side_by_side(key):
    """Shipped, mock and site at 1x CSS scale, top aligned, on the site's own dark ground."""
    shipped = Image.open(SHIPPED / f'{SHIPPED_FILE[key]}.png').convert('RGB')
    mock = Image.open(OUT / f'brand-{key}.png').convert('RGB')
    shipped = shipped.resize((740, round(shipped.height * 740 / shipped.width)), Image.LANCZOS)
    mock = mock.resize((740, round(mock.height * 740 / mock.width)), Image.LANCZOS)
    site = Image.open(SITE / f'{SITE_PAGE[key]}-1440.png').convert('RGB')
    return [shipped, mock, site]


def sheet(rows):
    gap = 32
    width = max(sum(im.width for im in r) + gap * (len(r) + 1) for r in rows)
    height = sum(max(im.height for im in r) for r in rows) + gap * (len(rows) + 1)
    out = Image.new('RGB', (width, height), (16, 10, 6))
    y = gap
    for r in rows:
        x = gap
        for im in r:
            out.paste(im, (x, y))
            x += im.width + gap
        y += max(im.height for im in r) + gap
    return out


def main():
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    for mood in BAND_CROPS:
        band_image(mood).save(OUT / f'band-{mood}.png')
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        for key, build in TABS.items():
            box = shoot(browser, build(), OUT / f'brand-{key}.png', '.win', 2)
            print(f'brand-{key}.png {box["width"]}x{box["height"]} css')
        shoot(browser, tooltip_board(), OUT / 'tooltip.png', '#board', 2)
        shoot(browser, key_html(), OUT / 'key.png', '#board', 1)
        browser.close()
    key = [Image.open(OUT / 'key.png').convert('RGB')]
    sheet([side_by_side(k) for k in ('me', 'gear', 'builds')] + [key]).save(OUT / 'side-by-side.png', optimize=True)
    sheet([side_by_side(k) for k in ('rotation', 'guild')] + [key]).save(OUT / 'side-by-side-2.png', optimize=True)
    print('side-by-side.png, side-by-side-2.png')


if __name__ == '__main__':
    main()

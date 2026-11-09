"""The in-game paperdoll on the owner's frame (design/specs/2026-11-09-directions.md, Paperdoll).

The frame is paperdoll/frame@{1x,2x} (cleaned by paperdoll_asset.py); every slot, the medallion socket, the title bar and
the centre panel are positioned from paperdoll/slots.json, as percentages of the frame, so it scales with its width.
- Each slot shows the item's real icon inside the slot's inner rectangle with a quality border; an empty slot shows the
  frame's own glyph (no overlay); a two-hander shows greyed in the off hand, as the client draws it.
- The medallion holds the character's Battle.net head render, or the Forever Sixty seal without one; the title bar the spec, band and faction on one line with the set DPS at its right.
- The centre: the selected slot's client tooltip and Open in the sim; hover or tap on a slot updates it. Signed in, the
  character's Battle.net render (characters.render_url, migration 0026) stands behind it, as the client's pane shows the model.
- Under MIN_SLOT_PX the frame gives way to its header plus a two-column slot list on the same materials.
"""
import base64
import json
from functools import lru_cache
from html import escape
from pathlib import Path

import data as D
from assets import logo_uri

HERE = Path(__file__).resolve().parent
SPEC = json.loads((HERE / 'paperdoll/slots.json').read_text())
FRAME_W, FRAME_H = SPEC['size']
MIN_SLOT_PX = 36
MEDALLION_SEAL_PX = 256  # the painted seal's export that covers the socket at 2x on the widest frame
SLOT_PCT_W = min(v['w'] for v in SPEC['slots'].values())
# The frame stays a frame while its slots are at least MIN_SLOT_PX wide: below this width, the list.
MIN_FRAME_PX = round(MIN_SLOT_PX * 100 / SLOT_PCT_W)
HEADER_SHARE = .152  # the frame's header strip (medallion, title bar, parchment) as a share of its height
ICON_INSET_PCT = 6  # the icon sits inside the slot's teal by this share of the slot, as the client insets it
QUALITY_BORDER = {'poor': '#9d9d9d', 'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#0070dd', 'epic': '#a335ee', 'legendary': '#ff8000'}


@lru_cache(maxsize=None)
def frame_uri(scale: str = '2x') -> str:
    return 'data:image/webp;base64,' + base64.b64encode((HERE / f'paperdoll/frame@{scale}.webp').read_bytes()).decode()


def _box(v: dict) -> str:
    return f'left:{v["x"]}%;top:{v["y"]}%;width:{v["w"]}%;height:{v["h"]}%'


CSS = '''
/* the frame: the owner's cleaned paperdoll art; everything on it positioned in % of the frame (paperdoll/slots.json) */
.pd{position:relative;width:100%;max-width:var(--pd-max,680px);aspect-ratio:''' + f'{FRAME_W}/{FRAME_H}' + ''';background:var(--frame) 0 0/100% 100% no-repeat;
container-type:inline-size}
.pd .sl{position:absolute;display:block}
.pd .sl img{position:absolute;inset:''' + f'{ICON_INSET_PCT}%' + ''';width:''' + f'{100 - 2 * ICON_INSET_PCT}%' + ''';height:''' + f'{100 - 2 * ICON_INSET_PCT}%' + ''';
box-sizing:border-box;border:2px solid var(--qb,#000);border-radius:3px;box-shadow:0 0 0 1px #000}
.pd .sl.two img{filter:grayscale(1) brightness(.55);border-color:#555}
.pd .sl.sel::after{content:"";position:absolute;inset:-4%;border:2px solid #f3eee2;border-radius:6px;box-shadow:0 0 0 1px #000}
.pd .med{position:absolute;display:flex;align-items:center;justify-content:center}.pd .med img{width:92%;height:92%;border-radius:50%}
.pd .tb{position:absolute;display:flex;align-items:center;gap:2cqw;padding:0 2.2cqw;color:#f3eee2;font:400 3.4cqw/1 var(--title);
letter-spacing:-.01em;text-shadow:0 1px 2px #000;white-space:nowrap}
.pd .tb small{font:600 2.3cqw var(--body);color:rgba(255,255,255,.8);text-shadow:0 1px 2px #000}
.pd .tb b{margin-left:auto;font:600 3.2cqw var(--body);color:#fff}.pd .tb b small{font-size:2.1cqw;color:rgba(255,255,255,.75)}
.pd .ctr{position:absolute;display:flex;flex-direction:column;align-items:center;justify-content:center;gap:3cqw;padding:4cqw}
.pd .ctr .tip{width:100%;max-width:340px;font-size:13.5px}
.pd .ctr .lk{font:600 14px var(--body)}
.pd .ctr .who{position:absolute;left:0;right:0;bottom:3cqw;text-align:center;font:400 3.6cqw var(--title);color:#f3eee2;text-shadow:0 1px 3px #000}
.pd .ctr .who small{display:block;font:600 2.3cqw var(--body);color:rgba(255,255,255,.8)}
/* the list form: the frame's header strip, then two columns of slots on its bronze and teal */
.pdl{display:none}
.pdl .hd{position:relative;width:100%;aspect-ratio:''' + f'{FRAME_W}/{round(FRAME_H * HEADER_SHARE)}' + ''';background:var(--frame) 0 0/100% auto no-repeat;container-type:inline-size}
.pdl .hd .med{position:absolute;display:flex;align-items:center;justify-content:center}.pdl .hd .med img{width:92%;height:92%;border-radius:50%}
.pdl .hd .tb{position:absolute;display:flex;align-items:center;gap:2cqw;padding:0 2.4cqw;margin:0;color:#f3eee2;font:400 5.4cqw/1 var(--title);
text-shadow:0 1px 2px #000;white-space:nowrap}
.pdl .hd .sub{position:absolute;left:21%;right:2.5%;top:38%;display:flex;align-items:center;padding:0 2cqw;font:600 4.2cqw var(--body);white-space:nowrap;color:rgba(255,255,255,.85);
text-shadow:0 1px 2px #000}.pdl .hd .sub b{margin-left:auto;color:#fff;font-size:4.4cqw}
.pdl .ls{display:grid;grid-template-columns:1fr 1fr;gap:6px 10px;padding:12px;background:#2b1e12;border:2px solid #8a6a3a;border-top:0;
border-radius:0 0 10px 10px;box-shadow:inset 0 0 0 1px #000}
.pdl .it{display:grid;grid-template-columns:44px 1fr;gap:8px;align-items:center;min-width:0}
.pdl .cell{position:relative;width:44px;height:44px;background:#1d6f7d;border:2px solid #c9a46a;border-radius:4px;box-shadow:inset 0 0 0 1px #000,0 0 0 1px #000}
.pdl .cell img{position:absolute;inset:3px;width:34px;height:34px;border:2px solid var(--qb,#000);border-radius:2px;box-sizing:border-box}
.pdl .it.two .cell img{filter:grayscale(1) brightness(.55);border-color:#555}.pdl .it.sel .cell{outline:2px solid #f3eee2;outline-offset:1px}
.pdl .it .n{display:block;font-weight:600;font-size:13px;line-height:16px}.pdl .it .s{display:block;font-size:12px;color:rgba(255,255,255,.6);white-space:nowrap;overflow:hidden;text-overflow:ellipsis}.pdl .hd .tb b{margin-left:auto;font:600 5cqw var(--body);color:#fff}.pdl .hd .tb b small{font-size:3.4cqw;color:rgba(255,255,255,.75)}.pdl .it .g{color:var(--gain,#32ae88);font-weight:700}.pdl .lead{padding:10px 12px;background:#2b1e12;border:2px solid #8a6a3a;border-top:0;font-size:14px}
.pdl .tip{margin-top:10px}.pdl.nohd .ls{border-top:2px solid #8a6a3a;border-radius:var(--r-sm)}
'''

LIST_SLOTS = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2',
              'main_hand', 'off_hand', 'ranged']
SLOT_LABEL = {**D.SLOT_NAME, 'shirt': 'Shirt', 'tabard': 'Tabard'}


def _icon(cell: dict) -> str:
    return f'<img src="{D.icon_uri(cell["icon"])}" alt="" style="--qb:{QUALITY_BORDER[cell["quality"]]}">'


def slot_html(name: str, cells: dict, sel: str, two_hand: bool) -> str:
    v = SPEC['slots'][name]
    c = cells.get(name)
    if name == 'off_hand' and two_hand and cells.get('main_hand'):
        return f'<a class="sl two" href="#" style="{_box(v)}" aria-label="Off Hand: two-hander">{_icon(cells["main_hand"])}</a>'
    if not c or c.get('state') != 'item':
        return f'<span class="sl" style="{_box(v)}" aria-label="{SLOT_LABEL[name]}: empty"></span>'  # the frame's glyph shows
    return f'<a class="sl{" sel" if name == sel else ""}" href="#" style="{_box(v)}" aria-label="{SLOT_LABEL[name]}: {escape(c["name"])}">{_icon(c)}</a>'


def _medallion_src(avatar: str) -> str:
    """The socket holds the character's Battle.net head render (characters.avatar_url, the -avatar.jpg media) when there is
    one; with none, signed out included, the Forever Sixty seal (owner, 2026-11-08), never the class crest."""
    return escape(avatar) if avatar else logo_uri(MEDALLION_SEAL_PX)


def frame(cells: dict, sel: str, two_hand: bool, title: str, sub: str, dps: str, centre: str,
          avatar: str = '', render: str = '', unit: str = 'DPS') -> str:
    """avatar and render: the character's Battle.net media URLs (referenced, never copied into the repo). The render stands
    behind the centre as the client's pane shows the model."""
    model = f'<i class="rd" style="--render:url({escape(render)})" aria-hidden="true"></i>' if render else ''
    slots = ''.join(slot_html(n, cells, sel, two_hand) for n in SPEC['slots'])
    return (f'<div class="pd" role="group" aria-label="{escape(title)} best in slot">{slots}'
            f'<span class="med" style="{_box(SPEC["medallion"])}"><img src="{_medallion_src(avatar)}" alt=""></span>'
            f'<h1 class="tb" style="{_box(SPEC["title"])}">{escape(title)} <small>{sub}</small>{f"<b>{dps} <small>{unit}</small></b>" if dps else ""}</h1>'
            f'<div class="ctr{" model" if render else ""}" style="{_box(SPEC["centre"])}">{model}{centre}</div></div>')


def _gain(c: dict) -> str:
    """A slot's gain over what is worn, or Worn (the Gear tab); nothing on the public list."""
    g = c.get('vs_worn')
    if g is None:
        return ''
    return ' &middot; Worn' if g == 'worn' else f' &middot; <span class="g">{g} DPS</span>'


def list_form(cells: dict, sel: str, two_hand: bool, title: str, dps: str, tip: str, lead: str = '', avatar: str = '', unit: str = 'DPS',
              head: bool = True) -> str:
    """Under MIN_FRAME_PX: the frame's header strip (medallion and title bar) and the slots as a two-column list."""
    items = []
    for n in LIST_SLOTS:
        c = cells.get(n)
        if n == 'off_hand' and two_hand and cells.get('main_hand'):
            items.append(f'<div class="it two"><span class="cell">{_icon(cells["main_hand"])}</span><span><span class="s">Off Hand</span>'
                         f'<span class="s">Two-hander</span></span></div>')
        elif c and c.get('state') == 'item':
            label = 'Two-Hand' if n == 'main_hand' and two_hand else SLOT_LABEL[n]
            items.append(f'<div class="it{" sel" if n == sel else ""}"><span class="cell">{_icon(c)}</span><span><span class="n" style="color:var(--q-{c["quality"]})">'
                         f'{escape(c["name"])}</span><span class="s">{label}{_gain(c)}</span></span></div>')
        else:
            items.append(f'<div class="it"><span class="cell"></span><span><span class="s">{SLOT_LABEL[n]}</span></span></div>')
    m = SPEC['medallion']
    hd_med = f'left:{m["x"]}%;top:{m["y"] / HEADER_SHARE:.2f}%;width:{m["w"]}%;height:{m["h"] / HEADER_SHARE:.2f}%'
    t = SPEC['title']  # the engraved bar, from the measured frame, in the strip's coordinates
    hd_tb = f'left:{t["x"]}%;top:{t["y"] / HEADER_SHARE:.2f}%;width:{t["w"]}%;height:{t["h"] / HEADER_SHARE:.2f}%'
    # one line: the band and faction are the tabs above (reviews); head=False drops the strip where the page band already
    # names the spec and its figure (the sim's run)
    hd = (f'<div class="hd"><span class="med" style="{hd_med}"><img src="{_medallion_src(avatar)}" alt=""></span>'
          f'<h1 class="tb" style="{hd_tb}">{escape(title)}{f"<b>{dps} <small>{unit}</small></b>" if dps else ""}</h1></div>') if head else ''
    return (f'<div class="pdl{"" if head else " nohd"}">{hd}'
            f'{f"<p class=lead>{lead}</p>" if lead else ""}<div class="ls">{"".join(items)}</div>{tip}</div>')

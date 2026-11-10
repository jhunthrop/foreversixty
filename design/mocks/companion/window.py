"""The desktop companion's window in the round-5 look (batch G, docs/superpowers/plans/2026-11-09-redesign.md).

    python3 design/mocks/companion/window.py      # -> design/mocks/companion/html/*.html
    node design/mocks/companion/shoot.cjs <web/node_modules>   # -> design/mocks/companion/png/*.png

One 800 x 620 window (internal/shell: the starting size; it resizes down to 560 x 480). Reuses the round-5
style (design/mocks/directions/style.py): the tokens, the painted band and brown ground, the panel and its frames,
the teal primary. The window has no floating nav bar of its own: the band carries the seal, the name and the four
page tabs as the site's underlined tab group (`.ctl`). The Status headline is the page's one lead panel (ornate);
Device's sign-in is Device's. Icons are the site's set (design/assets/icons/ui), read from their source files.
Copy is the companion's own (companion/ui/copy.js, internal/app/copy.go), sample values from ui/testdata.
"""
import re
import sys
from html import escape
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parents[2]
sys.path.insert(0, str(HERE.parent / 'directions'))
import style as S  # noqa: E402
from assets import png_uri  # noqa: E402

OUT = HERE / 'html'
ICONS = REPO / 'design/assets/icons/ui'
FONTS = REPO / 'companion/ui/fonts'
W, H = 800, 620


def ui_icon(name: str, size: int = 18) -> str:
    """One glyph of the site's set, drawn in currentColor (custom: filled even-odd; line: 2 px stroke)."""
    custom, line = ICONS / f'{name}.svg', ICONS / 'line' / f'{name}.svg'
    if custom.exists():
        svg = custom.read_text()
        box = re.search(r'viewBox="0 0 (\d+) \d+"', svg).group(1)
        d = re.search(r'\sd="([^"]+)"', svg).group(1)
        return (f'<svg class="gl" width="{size}" height="{size}" viewBox="0 0 {box} {box}" aria-hidden="true">'
                f'<path fill="currentColor" fill-rule="evenodd" d="{d}"/></svg>')
    body = re.search(r'<svg\b[^>]*>([\s\S]*)</svg>', line.read_text()).group(1)
    return (f'<svg class="gl" width="{size}" height="{size}" viewBox="0 0 24 24" fill="none" stroke="currentColor" '
            f'stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">{body}</svg>')


FACES = ''.join(
    f"@font-face{{font-family:'{fam}';font-weight:{wt};src:url('{(FONTS / f).as_uri()}') format('woff2')}}"
    for fam, wt, f in (('Open Sans', 400, 'open-sans-400-subset.woff2'), ('Open Sans', 600, 'open-sans-600-subset.woff2'),
                       ('Open Sans', 700, 'open-sans-700-subset.woff2'), ('Marcellus', 400, 'marcellus-400-subset.woff2'),
                       ('Archivo', 600, 'archivo-600-subset.woff2'), ('JetBrains Mono', 500, 'jetbrains-mono-500-subset.woff2')))

WINDOW_CSS = FACES + '''
html,body{height:100%;overflow:hidden}body{display:flex;flex-direction:column;min-height:0;font-size:14px;line-height:1.5}
.mono{font:500 12px/1.5 'JetBrains Mono',monospace}
.top{position:relative;flex:none}.top .in{display:flex;align-items:center;gap:18px;max-width:752px;margin:0 auto;padding:0 24px;height:60px}
.wm{display:flex;align-items:center;gap:10px;font:400 21px/1 var(--title);letter-spacing:var(--title-track);color:var(--cream)}.wm img{width:32px;height:32px}
.tabs{display:flex;gap:20px;margin-left:auto;align-self:stretch}.tabs a{display:flex;align-items:center;gap:7px;font:600 15px var(--ui);color:rgba(255,255,255,.72)}
.tabs a.on{color:#fff;box-shadow:inset 0 -3px 0 var(--blue)}.tabs .gl{opacity:.9}.gl{display:inline-block;vertical-align:-3px;flex:none}
body>main{flex:1 1 auto;min-height:0;overflow:hidden}.col{max-width:752px;margin:0 auto;padding:16px 24px 20px}
.card{margin-bottom:14px}.lead{padding:20px 24px}
.framed.ornate::after{background:none}
.corner{position:absolute;width:44px;height:44px;color:var(--cream);opacity:.8;z-index:1}
.corner.tl{top:4px;left:4px}.corner.tr{top:4px;right:4px;transform:scaleX(-1)}.corner.br{bottom:4px;right:4px;transform:scale(-1,-1)}.corner.bl{bottom:4px;left:4px;transform:scaleY(-1)}
.hl{display:flex;gap:18px;align-items:flex-start;position:relative;z-index:2}.hl .t{min-width:0;flex:1}
.hl h1{font-size:26px;line-height:30px;display:flex;align-items:center;gap:12px}
.dot{width:12px;height:12px;border-radius:50%;flex:none}.dot.on{background:#7fd48a}.dot.warn{background:#f7b90a}.dot.bad{background:#f0736b}.dot.off{background:rgba(255,255,255,.5)}
.hl p{margin:6px 0 0 24px;color:var(--text)}.hl .fig{text-align:right;white-space:nowrap}.hl .fig b{display:block;font:600 26px/30px var(--body);color:#fff}
.hl .fig span{font-size:13px;color:var(--muted)}
.acts{display:flex;gap:16px;align-items:center;margin:14px 0 0 24px}
.btn,.btn2{display:inline-flex;align-items:center;gap:8px;height:40px;padding:0 18px}.btn.sm,.btn2.sm{height:32px;padding:0 12px;font-size:14px}
.ghost{color:var(--gold);font-weight:600;display:inline-flex;align-items:center;gap:6px}
.rows .r{display:flex;justify-content:space-between;gap:16px;padding:9px 20px;border-bottom:1px solid rgba(255,255,255,.08)}.rows .r:last-child{border:0}
.rows .k{color:var(--muted)}.rows .v{color:#fff;text-align:right}.rows .v small{display:block;color:var(--muted);font-size:12px}
h2{font-size:20px;line-height:26px;margin:4px 0 10px}
.inst{padding:12px 20px;border-bottom:1px solid rgba(255,255,255,.08)}.inst:last-child{border:0}
.inst .l1{display:flex;justify-content:space-between;align-items:baseline}.inst b{color:#fff;font-weight:600}
.tag{display:inline-block;font:600 12px/18px var(--body);padding:0 7px;border-radius:var(--r-sm);background:rgba(15,96,100,.55);border:1px solid var(--teal-edge);color:#fff;margin-left:8px}
.ok{color:#7fd48a}.inst .path{display:block;color:var(--muted);overflow:hidden;text-overflow:ellipsis;white-space:nowrap}
.foot{flex:none;margin:0;padding:0;height:32px;font-size:12px;border-top:1px solid #000}.foot .w{display:flex;justify-content:space-between;align-items:center;height:32px;max-width:752px;padding:0 24px}
.rep{display:grid;grid-template-columns:10px 1fr auto;gap:12px;align-items:center;padding:11px 20px;border-bottom:1px solid rgba(255,255,255,.08)}.rep:last-child{border:0}
.rep .z{color:#fff;font-weight:600;font-size:15px}.rep .m{font-size:12px;color:var(--muted)}.rep .m b{color:var(--text);font-weight:600}
.mk{width:8px;height:8px;border-radius:50%}.mk.live{background:#7fd48a}.mk.done{border:1.5px solid rgba(255,255,255,.4)}
.h2l{display:flex;justify-content:space-between;align-items:baseline}
.grp{padding:14px 20px;border-bottom:1px solid rgba(255,255,255,.08)}.grp:last-child{border:0}
label,.lab{display:block;font:600 14px var(--body);color:#fff;margin-bottom:6px}
.sel{height:38px;border:1px solid #4a3f34;border-radius:var(--r-sm);background:#100a06;display:flex;align-items:center;justify-content:space-between;padding:0 12px;color:var(--text)}
.help{font-size:13px;color:var(--muted);margin-top:4px}
.sw{display:flex;align-items:center;gap:12px;font-weight:600;color:#fff}.sw i{width:38px;height:22px;border-radius:11px;background:var(--teal);border:1px solid var(--teal-edge);position:relative}
.sw i::after{content:"";position:absolute;right:3px;top:3px;width:14px;height:14px;border-radius:50%;background:#fff}
.fl{border:1px solid rgba(255,255,255,.1);border-radius:var(--r-sm)}.fl .f{display:flex;justify-content:space-between;align-items:center;padding:9px 12px;border-bottom:1px solid rgba(255,255,255,.08)}.fl .f:last-child{border:0}
.save{position:absolute;left:0;right:0;bottom:32px;background:var(--foot);border-top:1px solid #000}.save .w{display:flex;align-items:center;justify-content:space-between;height:60px;max-width:752px;padding:0 24px}
.code{display:inline-block;margin-top:12px;padding:10px 18px;background:#100a06;border:1px solid #4a3f34;border-radius:var(--r-sm);font:500 28px/1.2 'JetBrains Mono',monospace;letter-spacing:.16em;color:#fff}
.bar{height:3px;width:200px;background:rgba(255,255,255,.1);border-radius:2px;margin-top:10px}.bar i{display:block;height:100%;width:72%;background:var(--gold);border-radius:2px}
'''

CORNER_D = re.search(r'\sd="([^"]+)"', (REPO / 'design/assets/ornaments/corner-filigree.svg').read_text()).group(1)
CORNERS = ''.join(f'<svg class="corner {c}" viewBox="0 0 1761 1761" aria-hidden="true"><path fill="currentColor" fill-rule="evenodd" d="{CORNER_D}"/></svg>'
                  for c in ('tl', 'tr', 'br', 'bl'))
TABS = (('Status', 'upload'), ('Reports', 'logs'), ('Device', 'account'), ('Settings', 'setup'))


def window(cur: str, body: str, extra: str = '') -> str:
    tabs = ''.join(f'<a class="{"on" if name == cur else ""}">{ui_icon(ic, 18)}{name}</a>' for name, ic in TABS)
    top = (f'<div class="band top"><div class="in"><span class="wm"><img src="{png_uri("logo/lx-seal-v3-64.png")}" alt="">'
           f'Forever Sixty</span><nav class="tabs">{tabs}</nav></div></div>')
    foot = ('<footer class="foot"><div class="w"><span>Companion 0.2.13</span><a style="color:rgba(255,255,255,.72);font-weight:400">'
            'foreversixty.gg</a></div></footer>')
    return S.page(f'Companion: {cur}', WINDOW_CSS, f'{top}<main><div class="col">{body}</div></main>{extra}{foot}')


def status() -> str:
    lead = (f'<section class="card framed ornate lead">{CORNERS}<div class="hl"><div class="t"><h1><span class="dot on"></span>Logging</h1>'
            '<p>Uploading your fights as they end · credited to Thandrel</p>'
            f'<div class="acts"><a class="ghost">{ui_icon("link", 16)}Open this report on the site</a></div></div>'
            '<div class="fig"><b class="num">14</b><span>fights uploaded</span></div></div></section>')
    rows = [('Current report', 'Onyxia\'s Lair<small>Started 20:41 · credited to Thandrel</small>'), ('Queued', '0'),
            ('Addon', 'Addon 2.0.16 · up to date<small>Data 2026.11.09 · up to date</small>'),
            ('Characters sent', '2 characters, 4 minutes ago')]
    panel = '<section class="card rows">' + ''.join(f'<div class="r"><span class="k">{k}</span><span class="v">{v}</span></div>' for k, v in rows) + '</section>'
    inst = ('<h2>Game folders</h2><section class="card"><div class="inst"><div class="l1"><span><b>Classic beta</b><span class="tag">Watching</span></span>'
            f'<span class="ok">{ui_icon("check", 14)} Advanced logging on</span></div><span class="muted" style="font-size:13px">Last played today at 20:41</span>'
            '<span class="path mono">C:\\Program Files (x86)\\World of Warcraft\\_classic_beta_</span></div></section>')
    return window('Status', lead + panel + inst)


def device() -> str:
    lead = (f'<section class="card framed ornate lead">{CORNERS}<div class="hl"><div class="t"><h1><span class="dot bad"></span>Not connected</h1>'
            '<p>Sign in on foreversixty.gg so the companion can upload your fights and keep the addon up to date.</p>'
            f'<div class="acts"><a class="btn">{ui_icon("account", 18)}Sign in with Battle.net</a></div>'
            '<p class="help" style="margin-left:24px">Opens your browser. Your Battle.net password never reaches this app.</p>'
            f'<p style="margin:16px 0 0 24px" class="muted">{ui_icon("chevron-right", 14)} Having trouble? Paste a code from the site</p></div></div></section>')
    wait = (f'<section class="card pad"><div class="hl"><div class="t"><h1 style="font-size:22px"><span class="dot off"></span>Waiting for you to sign in...</h1>'
            '<p>Check that your browser shows this code:</p><div style="margin-left:24px"><span class="code">KQ4-7TM</span><div class="bar"><i></i></div>'
            '<p class="help">Expires in 7:12</p><div class="acts" style="margin-left:0"><a class="btn2 sm">Open the browser again</a><a class="ghost">Cancel</a></div></div></div></div></section>')
    return window('Device', lead + wait)


def reports() -> str:
    rows = [('live', "Onyxia's Lair", '<b>in progress</b> · 14 fights · credited to Thandrel'),
            ('done', 'Molten Core', 'Today 19:02 · 38 fights · credited to Thandrel'),
            ('done', 'Blackwing Lair', 'Yesterday · 27 fights · credited to Thandrel'),
            ('done', 'Zul\'Gurub', 'Tue 3 Nov · 22 fights · credited to Kelsharra'),
            ('done', 'Molten Core', 'Sun 1 Nov · 41 fights · credited to Thandrel')]
    body = ('<div class="h2l"><h2>Your reports</h2><a>Open all my reports on the site</a></div><section class="card">' +
            ''.join(f'<div class="rep"><span class="mk {m}"></span><span><span class="z">{escape(z)}</span><br><span class="m">{meta}</span></span>'
                    f'<a class="btn2 sm">{ui_icon("link", 14)}Copy link</a></div>' for m, z, meta in rows) + '</section>')
    return window('Reports', body)


def settings() -> str:
    body = ('<section class="card">'
            '<div class="grp"><label>New reports are</label><div class="sel">Public' + ui_icon('chevron-down', 16) + '</div><p class="help">Anyone can find it in rankings and search.</p></div>'
            '<div class="grp"><span class="sw"><i></i>Keep the addon up to date</span><p class="help">Installs new addon and data versions into every game folder here, when the game is closed.</p></div>'
            '<div class="grp"><label>Reports are credited to</label><div class="sel">Automatic (currently Thandrel)' + ui_icon('chevron-down', 16) + '</div></div>'
            '<div class="grp"><div class="h2l"><span class="lab">Game folders</span><a class="ghost">Add a folder...</a></div><div class="fl">'
            '<div class="f"><span><b style="color:#fff;font-weight:600">Classic beta</b> <span class="muted" style="font-size:12px">Found automatically</span><br>'
            '<span class="mono muted">C:\\Program Files (x86)\\World of Warcraft\\_classic_beta_</span></span></div></div></div></section>')
    save = '<div class="save"><div class="w"><span class="muted">No changes</span><a class="btn" style="opacity:.45">Save changes</a></div></div>'
    return window('Settings', body, save)


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    for name, fn in (('status', status), ('device', device), ('reports', reports), ('settings', settings)):
        (OUT / f'{name}.html').write_text(fn())
        print(f'html/{name}.html')


if __name__ == '__main__':
    main()

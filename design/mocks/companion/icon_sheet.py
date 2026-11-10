"""The companion's app and tray icons as each OS shows them, plus the tray menu (batch G).

    python3 design/mocks/companion/icon_sheet.py   # -> design/mocks/companion/html/icons.html, tray-menu.html

Every image is read out of the BUILT assets (companion/internal/icon/assets: seal.icns, seal.ico, the tray ICOs and
PNGs), never re-rendered, so the sheet shows exactly what ships. The tray menu is the native menu drawn as each
OS draws it, with the companion's own items (internal/shell/shell_gui.go) and the headline titles
(internal/app/copy.go); the native menu cannot be screenshotted from a webview, so this is its stand-in.
"""
import base64
import struct
from pathlib import Path

HERE = Path(__file__).resolve().parent
REPO = HERE.parents[2]
ASSETS = REPO / 'companion/internal/icon/assets'
OUT = HERE / 'html'


def uri(blob: bytes) -> str:
    return 'data:image/png;base64,' + base64.b64encode(blob).decode()


def icns_entries() -> dict[str, bytes]:
    b = (ASSETS / 'seal.icns').read_bytes()
    out, p = {}, 8
    while p < len(b):
        kind, n = b[p:p + 4].decode(), struct.unpack('>I', b[p + 4:p + 8])[0]
        out[kind] = b[p + 8:p + n]
        p += n
    return out


def ico_entries(name: str) -> dict[int, bytes]:
    b = (ASSETS / name).read_bytes()
    count = struct.unpack('<H', b[4:6])[0]
    out = {}
    for i in range(count):
        e = b[6 + 16 * i:22 + 16 * i]
        size = e[0] or 256
        length, offset = struct.unpack('<II', e[8:16])
        out[size] = b[offset:offset + length]
    return out


CSS = '''body{margin:0;background:#1b1b1b;color:#eee;font:14px/1.4 -apple-system,"Segoe UI",sans-serif;padding:24px}
h1{font-size:20px;margin:0 0 16px}h2{font-size:15px;margin:22px 0 8px;color:#bbb}
.row{display:flex;gap:28px;align-items:flex-end;flex-wrap:wrap}.cell{display:flex;flex-direction:column;align-items:center;gap:6px}
.cell span{font-size:11px;color:#999}.pad{padding:14px 18px;border-radius:12px;display:flex;gap:22px;align-items:flex-end}
.dock-dark{background:rgba(40,40,40,.85);border:1px solid #444}.dock-light{background:rgba(235,235,235,.8);border:1px solid #ccc}
.task-dark{background:#202020;border-radius:0}.task-light{background:#f3f3f3;border-radius:0}.linux{background:#2c2c2c}
.mb-light{background:#ececec;height:24px;padding:0 10px;align-items:center;border-radius:0}.mb-dark{background:#2a2a2a;height:24px;padding:0 10px;align-items:center;border-radius:0}
img{image-rendering:auto;display:block}
'''


def img(blob: bytes, css_px: int) -> str:
    return f'<img src="{uri(blob)}" width="{css_px}" height="{css_px}" alt="">'


def cell(inner: str, label: str) -> str:
    return f'<div class="cell">{inner}<span>{label}</span></div>'


def icons_page() -> str:
    mac, ico = icns_entries(), ico_entries('seal.ico')
    # Each entry at the size the OS draws it on a Retina screen: 128 pt from 256 px, 64 pt from 128, and so on.
    dock = lambda bg: f'<div class="pad {bg}">' + ''.join(cell(img(mac[k], pt), f'{pt} pt') for k, pt in (('ic08', 128), ('icp6', 32), ('icp5', 16), ('ic07', 64))) + '</div>'
    task = lambda bg: f'<div class="pad {bg}">' + ''.join(cell(img(ico[s], s // 2 if s > 32 else s), f'{s} px') for s in (16, 24, 32, 48)) + '</div>'
    explorer = ''.join(cell(img(ico[s], px), f'{px} px') for s, px in ((48, 48), (64, 64), (256, 128), (256, 256)))
    tones = ('off', 'on', 'warn', 'bad')
    tray_ico = {t: ico_entries(f'seal-tray-{t}.ico') for t in tones}
    tray_png = {t: (ASSETS / f'seal-tray-{t}.png').read_bytes() for t in tones}
    win_tray = lambda bg: f'<div class="pad {bg}">' + ''.join(cell(img(tray_ico[t][s], s // (s // 16)), f'{t} {s}px') for t in tones for s in (16, 32)) + '</div>'
    mac_bar = lambda bg: f'<div class="pad {bg}">' + ''.join(img(tray_png[t], 16) for t in tones) + '</div>'
    return (f'<!doctype html><html><head><meta charset="utf-8"><title>Companion icons</title><style>{CSS}</style></head><body>'
            '<h1>Forever Sixty companion: app and tray icons (seal v3), as shipped</h1>'
            '<h2>macOS Dock, dark and light (seal.icns: one quantized PNG per pixel size, 16 to 512, on Apple\'s 824/1024 grid; drawn at Retina 2x)</h2>'
            f'<div class="row">{dock("dock-dark")}{dock("dock-light")}</div>'
            '<h2>Windows taskbar at 100 % and 200 %, dark and light (seal.ico)</h2>'
            f'<div class="row">{task("task-dark")}{task("task-light")}</div>'
            f'<h2>Windows Explorer (seal.ico; 128 is Windows scaling the 256 entry)</h2><div class="row">{explorer}</div>'
            '<h2>Windows tray, the four tones, dark and light (seal-tray-*.ico); Linux takes the 32 px PNGs below</h2>'
            f'<div class="row">{win_tray("task-dark")}{win_tray("task-light")}</div>'
            '<h2>macOS menu bar, the four tones at 16 pt (seal-tray-*.png at 2x), light and dark</h2>'
            f'<div class="row">{mac_bar("mb-light")}{mac_bar("mb-dark")}</div>'
            '<h2>The tray tones enlarged: the 32 px PNG shown at 64 px, so the corner badge can be judged</h2>'
            f'<div class="row"><div class="pad task-dark">' + ''.join(cell(img(tray_png[t], 64), t) for t in tones) + '</div></div>'
            '</body></html>')


MENU_CSS = CSS + '''.menus{display:flex;gap:48px;align-items:flex-start}
.win{background:#2b2b2b;border:1px solid #444;border-radius:8px;padding:4px;width:260px;font:14px "Segoe UI",sans-serif;box-shadow:0 8px 24px rgba(0,0,0,.5)}
.win div{padding:7px 12px;border-radius:4px}.win .dis{color:#8a8a8a}.win .hov{background:#3a3a3a}.win hr{border:0;border-top:1px solid #444;margin:4px 0}
.mac{background:rgba(236,236,236,.96);color:#111;border-radius:7px;padding:5px;width:250px;font:13px -apple-system,sans-serif;box-shadow:0 8px 24px rgba(0,0,0,.4)}
.mac div{padding:3px 10px;border-radius:4px}.mac .dis{color:#999}.mac .hov{background:#2f6fd6;color:#fff}.mac hr{border:0;border-top:1px solid #cfcfcf;margin:5px 8px}
.bar{display:flex;gap:10px;align-items:center;justify-content:flex-end;background:#ececec;height:24px;padding:0 12px;width:250px;margin-bottom:4px;border-radius:0}
.tip{display:inline-block;background:#2b2b2b;border:1px solid #555;padding:4px 8px;font:12px "Segoe UI",sans-serif;margin-top:10px}
'''


def menu_page() -> str:
    on = (ASSETS / 'seal-tray-on.png').read_bytes()
    items = ('Open Forever Sixty', 'Open my reports on the site')
    win = ('<div class="win"><div class="dis">Logging</div><hr>' + ''.join(f'<div class="{"hov" if i == 0 else ""}">{t}</div>' for i, t in enumerate(items))
           + '<hr><div>Quit</div></div><div class="tip">Forever Sixty · Logging</div>')
    mac = (f'<div class="bar">{img(on, 16)}</div><div class="mac"><div class="dis">Logging</div><hr>'
           + ''.join(f'<div class="{"hov" if i == 1 else ""}">{t}</div>' for i, t in enumerate(items)) + '<hr><div>Quit</div></div>')
    return (f'<!doctype html><html><head><meta charset="utf-8"><title>Tray menu</title><style>{MENU_CSS}</style></head><body>'
            '<h1>The tray menu, as Windows 11 and macOS draw it (the native menu; items from internal/shell/shell_gui.go)</h1>'
            f'<div class="menus"><div><h2>Windows</h2>{win}</div><div><h2>macOS</h2>{mac}</div></div></body></html>')


def main() -> None:
    OUT.mkdir(parents=True, exist_ok=True)
    (OUT / 'icons.html').write_text(icons_page())
    (OUT / 'tray-menu.html').write_text(menu_page())
    print('html/icons.html html/tray-menu.html')


if __name__ == '__main__':
    main()

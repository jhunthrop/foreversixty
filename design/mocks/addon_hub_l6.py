"""Review boards for lane L6, the chip as the one launcher, as built (design/specs/2026-11-13-addon-hub.md 6).

    python3 design/mocks/addon_hub_l6.py              # every board
    python3 design/mocks/addon_hub_l6.py l6-states

Writes design/reviews/2026-11-13-addon-hub/l6/<board>.png at device scale 2, every frame at its real size,
drawn to the built code's own rules, not the approved board's (hub-chip.png):

- The chip: addon/ForeverSixty/surfaces/Chip.lua. 28 high, the 22 px ringed crest 3 in from the left edge,
  the line 8 past it and 10 clear of the right edge (Theme.SIZES chipCrest, chipCrestInset, chipTextGap,
  chipTextInset 43), at most 280 wide; the raised ground with the 1 px line edge; 12 px text in the body
  colour, muted with nothing waiting and in combat; "+N" muted after the line (L.chipMore).
- The cut: Chip.shownText. The answer after " · " or a trailing figure (Chip.figureOf) and the "+N" stay
  whole; the part before them gives way with an ellipsis. Measured here in the browser with the same font.
- Each line is the owning tab's own `waiting` line (Hub.urgent's first item): Guild (GuildRaid.waitingOf:
  "<raid> <clock> · Sign up" in the client gold, "<raid> <clock> · <job>" with the job in the client gold),
  tonight's raid once answered (GuildRaid.tonightLine: the answer in its colour, no dot), Officer
  (OfficerTab.waiting: "<item> · N of M answered"), Builds (BuildDialog.waiting: "Spend 1 point: <talent>"
  with "Spend 1 point" in the client gold, "2 builds from the site"), Gear (GearTab.equipLine: "Equip <item>
  +N.N", the gain in the success colour), Me (MeTab.waitingItems: "<weekday>: median N, best <boss>", the
  median in its parse colour). The clock is Compat.clockText on a 24-hour client.
- In combat (the "Hide in combat" switch off): "After combat", muted, the chip at 25 % (Chip.COMBAT_ALPHA).
- The hover: the client's GameTooltip (Chip.hoverLines): "Forever Sixty", every item in order, then the
  hints in the tooltip's muted grey. The minimap button (Minimap.lua) is the chip's twin: the same lines,
  its own move ("Drag: move").

Every character here is made up: Velisandre, a level 60 Night Elf hunter in Thornwake Company, and Corvane,
her level 42 hunter alt. Items and talents are the game's own (items/hunter.json, Data.lua).
"""
import os
import sys

import addon_hub as hub
from addon_hub import ADDON, CLIENT_GOLD, esc
from addon_ux import ROOT, l, tga_uri, tooltip

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub/l6'
TMUTED = ADDON['tmuted']
ROOM = 280 - 43

CSS = f'''
.scene {{ position: relative; background: radial-gradient(ellipse at 70% 0%, #3c4a33 0%, #1d2419 60%, #12160f 100%);
  padding: 18px 22px; }}
.lch {{ display: inline-flex; align-items: center; height: 28px; box-sizing: border-box; max-width: 280px;
  background: {ADDON['raised']}; border: 1px solid {ADDON['line']}; font-family: Site, sans-serif; font-size: 12px;
  color: {ADDON['body']}; white-space: nowrap; padding: 0 9px 0 2px; }}
.lch .crest {{ margin-right: 8px; }}
.lch.muted .tx {{ color: {ADDON['muted']}; }}
.lch .tx {{ display: inline-block; }}
.st {{ display: grid; grid-template-columns: 230px 300px 1fr; align-items: center; row-gap: 12px; column-gap: 18px; }}
.st .k {{ color: #d2cbbb; font-family: Site, sans-serif; font-size: 13px; }}
.st .o {{ color: #9a9484; font-family: Site, sans-serif; font-size: 13px; }}
.mm {{ position: absolute; width: 150px; height: 150px; border-radius: 50%;
  background: radial-gradient(circle at 45% 40%, #5b6b3f, #33402a 55%, #1d2617);
  box-shadow: 0 0 0 4px #2b2416, 0 0 0 6px #6b5626, 0 4px 14px #000; }}
.mmb {{ position: absolute; width: 32px; height: 32px; }}
.mmb .g {{ position: absolute; left: 7px; top: 6px; width: 20px; height: 20px; border-radius: 50%; background: #000; }}
.mmb img.i {{ position: absolute; left: 7px; top: 6px; width: 20px; height: 20px; border-radius: 50%; }}
.mmb .r {{ position: absolute; left: 0; top: 0; width: 32px; height: 32px; border-radius: 50%;
  box-shadow: inset 0 0 0 3px #b8a061, inset 0 0 0 5px #4d3f1e; }}
.cap2 {{ font-family: Site, sans-serif; font-size: 13px; color: #c9c2b2; margin: 0 0 8px; max-width: 560px; }}
'''

# The fit (Chip.fit with Chip.shownText's keep) in the browser, on the font the chip draws in.
FIT_JS = '''
() => {
  for (const tx of document.querySelectorAll('.lch .tx[data-keep]')) {
    const cnt = tx.parentElement.querySelector('.cnt');
    const room = Number(tx.dataset.room) - (cnt ? cnt.getBoundingClientRect().width + 6 : 0);
    const keep = tx.querySelector('.keep');
    const body = tx.querySelector('.body');
    let text = body.textContent;
    if (tx.getBoundingClientRect().width <= room) continue;
    while (text.length > 0) {
      text = text.slice(0, -1).replace(/[\\s,]+$/, '');
      body.textContent = text + '\\u2026';
      if (tx.getBoundingClientRect().width <= room) break;
    }
  }
}
'''


def span(hex_, text):
    return f'<span style="color:{hex_}">{esc(text)}</span>'


def gold(text):
    return span(CLIENT_GOLD, text)


def more(count):
    return f'<span class="m cnt" style="margin-left:6px">+{count}</span>' if count else ''


def chip(body, keep='', count=0, cls='hunter', muted=False, alpha=1, lead=''):
    """The chip as Chip.lua draws it: `lead` (HTML) never cut, `body` the part that may be cut, `keep` (HTML)
    the figure, `count`."""
    tx = (f'<span class="tx" data-keep="1" data-room="{ROOM}">{lead}<span class="body">{esc(body)}</span>'
          f'<span class="keep">{keep}</span></span>{more(count)}')
    return (f'<span class="lch{" muted" if muted else ""}" style="opacity:{alpha}">'
            f'{hub.crest(cls, 22)}{tx}</span>')


RAID = "Onyxia's Lair 20:00"
STATES = [
    ('Raid day, signed up', chip(RAID, ' · ' + span(ADDON['success'], 'In'), 2), 'Opens Guild (no dot: nothing to do)'),
    ('Raid day, not signed up', chip(RAID, ' · ' + gold('Sign up'), 1), 'Opens Guild'),
    ('Raid day, a job', chip(RAID, ' · ' + gold('Phase 2 ranged, west side')), 'Opens Guild'),
    ('Master loot', chip("Striker's Mark", ' · 2 of 2 answered', 1), 'Opens Officer'),
    ('Corvane, level 42, a point waiting', chip(': Intimidation', lead=gold('Spend 1 point')),
     'Opens Builds and the talent window'),
    ('Builds from the site', chip('2 builds from the site'), 'Opens Builds'),
    ('An upgrade in the bags', chip('Equip Gloves of Holy Might', ' ' + span(ADDON['success'], '+11.2 DPS')), 'Opens Gear'),
    ('A long name, cut', chip("Equip Rhok'delar, Longbow of the Ancient Keepers", ' ' + span(ADDON['success'], '+24.6 DPS'),
                              1), 'Opens Gear'),
    ('After a raid night', chip(', best Golemagg', lead='Tuesday: median ' + span(hub.parse_hex(74), '74')), 'Opens Me'),
    ('In combat (Hide in combat off)', chip('After combat', muted=True, alpha=0.25), 'Does not open'),
    ('Nothing waiting', chip('Forever Sixty', muted=True), 'Opens the last tab'),
]


def board_states():
    rows = ''.join(f'<div class="k">{a}</div><div>{c}</div><div class="o">{o}</div>' for a, c, o in STATES)
    body = f'<div class="scene"><div class="st">{rows}</div></div>'
    return hub.board2('<b>L6 · the chip in every urgent state</b> · Velisandre, level 60 hunter, unless named. '
                      'Each line is the owning tab\'s own waiting line; a click opens the tab it names.',
                      f'<style>{CSS}</style>{body}')


def hover_lines(drag):
    hints = ['Click: open on Guild', 'Right-click: settings'] + ([drag] if drag else [])
    return ([f'<div class="t" style="color:#fff">Forever Sixty</div>',
             l(f"Onyxia's Lair 20:00 · {span(ADDON['success'], 'In')}", 'Guild', rcolor=TMUTED),
             l('Starts in 1 h 12 min', color=TMUTED),
             l('Equip Gloves of Holy Might ' + span(ADDON['success'], '+11.2 DPS'), 'Gear', rcolor=TMUTED),
             l('2 builds from the site', 'Builds', rcolor=TMUTED)] + [l(h, color=TMUTED) for h in hints])


def spend_lines():
    return [f'<div class="t" style="color:#fff">Forever Sixty</div>',
            l(gold('Spend 1 point') + ': Intimidation', 'Builds', rcolor=TMUTED),
            l('Click: open on Builds and the talent window', color=TMUTED),
            l('Right-click: settings', color=TMUTED), l('Shift-drag: move', color=TMUTED)]


def board_hover():
    icon = tga_uri(ROOT / 'addon/ForeverSixty/media/minimap.tga')
    chip_scene = (f'<div class="scene" style="width:470px;height:300px">'
                  f'<div class="mm" style="right:24px;top:18px"></div>'
                  f'<div style="position:absolute;right:186px;top:18px">'
                  f'{chip(RAID, " · " + span(ADDON["success"], "In"), 2)}</div>'
                  f'<div style="position:absolute;right:186px;top:52px">'
                  f'{tooltip(hover_lines("Shift-drag: move"), width=300)}</div></div>')
    spend_scene = (f'<div class="scene" style="width:470px;height:230px">'
                   f'<div class="mm" style="right:24px;top:18px"></div>'
                   f'<div style="position:absolute;right:186px;top:18px">'
                   f'{chip(": Intimidation", lead=gold("Spend 1 point"))}</div>'
                   f'<div style="position:absolute;right:186px;top:52px">'
                   f'{tooltip(spend_lines(), width=300)}</div></div>')
    twin_scene = (f'<div class="scene" style="width:470px;height:300px">'
                  f'<div class="mm" style="right:24px;top:18px"></div>'
                  f'<div class="mmb" style="right:150px;top:118px"><span class="g"></span>'
                  f'<img class="i" src="{icon}"><span class="r"></span></div>'
                  f'<div style="position:absolute;right:190px;top:140px">'
                  f'{tooltip(hover_lines("Drag: move"), width=300)}</div></div>')
    html = (f'<div class="row" style="align-items:flex-start;gap:28px;flex-wrap:wrap;max-width:1000px">'
            f'<div><p class="cap2">The chip\'s hover: everything waiting, in order, each with the tab it opens; how '
            f'long until the pull under tonight\'s raid; then what a click does.</p>{chip_scene}</div>'
            f'<div><p class="cap2">The minimap button, off by default, the chip\'s twin: the same click and '
            f'hover, its own move. Shown here with the chip switched off.</p>{twin_scene}</div>'
            f'<div><p class="cap2">Corvane, level 42, a point to spend: the click opens the talent window as well '
            f'as Builds, and the hover says so.</p>{spend_scene}</div></div>')
    return hub.board2('<b>L6 · the hover and the minimap twin</b>', f'<style>{CSS}</style>{html}')


BOARDS = {
    'l6-states': board_states,
    'l6-hover': board_hover,
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000})
        pg = ctx.new_page()
        for name in names:
            pg.set_content(BOARDS[name]())
            pg.evaluate('document.fonts.ready')
            pg.evaluate(FIT_JS)
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

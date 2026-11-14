"""Review boards for the built Me tab (lane L2 of docs/superpowers/plans/2026-11-13-addon-hub.md).

    python3 design/mocks/addon_hub_me.py              # every board
    python3 design/mocks/addon_hub_me.py me-raid-day

Writes design/reviews/2026-11-13-addon-hub/l2/<board>.png at device scale 2, the window at its real 740 px width.
Unlike the acceptance boards (addon_hub.py), every figure, word and colour inside the tab comes from the addon's
own model: this script hands each state's inputs to addon/ForeverSixty/surfaces/MeTab.lua through
design/mocks/addon_me_model.lua (plain Lua with the specs' client mock) and draws the JSON it returns on the
layout of surfaces/MeFrame.lua (MeFrame.LAYOUT and MeFrame.COLUMNS). The chrome is the Builds window's
(addon_builds_window.py), the frame the tab sits in.

Made up, as on the acceptance boards: Velisandre, a level 60 Beast Mastery hunter in Thornwake Company, her
Molten Core night and her sim. The Guild, Gear and Rotation lines in Right now are those tabs' own leads, drawn
with the acceptance board's words until their lanes land; the Builds line is BuildDialog.lead's shape.
"""
import json
import os
import re
import subprocess
import sys
import tempfile

import addon_builds_window as bw
import addon_hub as hub
from addon_hub import ADDON, ROOT, esc, icon, item_icon

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub/l2'
SEP = ' · '
# MeFrame.LAYOUT and MeFrame.COLUMNS, read from the file so the boards follow the frame.
FRAME = (ROOT / 'addon/ForeverSixty/surfaces/MeFrame.lua').read_text()
LAYOUT = {k: int(v) for k, v in re.findall(r'(\w+) = (\d+)', FRAME[FRAME.index('MeFrame.LAYOUT = {'):
                                                                  FRAME.index('}', FRAME.index('MeFrame.LAYOUT = {'))])}
COLUMNS = [dict(key=k, x=int(x), width=int(w), right=bool(r)) for k, x, w, r in re.findall(
    r'key = "(\w+)", head = "\w+", x = (\d+), width = (\d+)(, right = true)?', FRAME)]
HEADS = {'name': 'Boss', 'parse': 'Parse', 'loss': 'Biggest loss', 'died': 'Died', 'consumables': 'Consumables'}

CSS = f'''
.me .row {{ position: relative; height: {LAYOUT['partRow']}px; }}
.me .row > * {{ position: absolute; top: {(LAYOUT['partRow'] - LAYOUT['line']) // 2}px; white-space: nowrap; }}
.me .trk {{ position: absolute; top: {(LAYOUT['partRow'] - LAYOUT['track']) // 2}px !important; height: {LAYOUT['track']}px;
  background: {hub.color('base.lineSoft')}; }}
.me .trk i {{ display: block; height: 100%; }}
.me .big {{ height: {LAYOUT['big']}px; display: flex; align-items: baseline; gap: {LAYOUT['reasonGap']}px; }}
.me .big .n {{ font-size: {LAYOUT['bigSize']}px; font-weight: 600; line-height: {LAYOUT['big']}px; }}
.me .tb {{ position: relative; }}
.me .tb .r {{ position: relative; height: {LAYOUT['bossRow']}px; }}
.me .tb .r span {{ position: absolute; top: {(LAYOUT['bossRow'] - LAYOUT['line']) // 2}px; white-space: nowrap;
  overflow: hidden; text-overflow: ellipsis; }}
.me .sum {{ margin-top: {LAYOUT['summaryGap']}px; white-space: nowrap; }}
.me .link {{ color: {ADDON['gold']}; }}
.me .now {{ position: relative; height: {LAYOUT['nowRow']}px; }}
.me .now .ni {{ position: absolute; left: 0; top: 2px; line-height: 0; }}
.me .now .tx {{ position: absolute; left: {LAYOUT['nowIcon'] + LAYOUT['nowGap']}px;
  width: {288 - LAYOUT['nowIcon'] - LAYOUT['nowGap'] - LAYOUT['nowLink']}px; }}
.me .now .tx div {{ white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.me .now .lk {{ position: absolute; right: 0; top: 0; }}
.me .dps {{ font-size: {LAYOUT['simSize']}px; font-weight: 600; color: {ADDON['strong']}; margin-top: 2px; }}
.me .gain {{ height: {LAYOUT['gainRow']}px; display: flex; align-items: center; gap: {LAYOUT['gainGap']}px; }}
.me .gain .nm {{ flex: 1; min-width: 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.me .gain .v {{ width: {LAYOUT['gainValue']}px; text-align: right; }}
.me .sh {{ margin-bottom: {LAYOUT['underHead']}px; }}
.me .simft {{ display: flex; justify-content: space-between; align-items: center; margin-top: 4px; height: 20px; }}
.me .btn {{ display: inline-flex; align-items: center; height: 20px; padding: 0 7px; font-weight: 600;
  color: {ADDON['strong']}; background: {ADDON['blue']}; border: 1px solid {ADDON['blue_edge']}; }}
.me .btn.off {{ opacity: .4; }}
.me .gap {{ height: {LAYOUT['sectionGap']}px; }}
.me .ok {{ color: {ADDON['success']}; }}
'''

# --- Velisandre's made-up state, as MeTab.model's inputs ---------------------------------------------------

RECORD = {'score': 1184, 'tier': 'rare', 'as_of': '2026-12-09', 'fights': 22,
          'components': {'gear': 410, 'experience': 380, 'parses': 228, 'attendance': 120, 'discipline': 46}}
BOSSES = [(name, pct, loss, secs, deaths, missed) for name, pct, loss, secs, deaths, missed, _ in hub.LAST_RAID['bosses']]
LAST_RAID = {'report_id': 'r_velisandre_mc', 'date': '2026-12-08', 'bosses': [
    {'name': name, 'percentile': pct, 'biggest_loss': loss.lower() if loss == 'Downtime' else loss,
     'loss_seconds': secs, 'deaths': deaths, 'consumables_missed': [m.lower() for m in missed]}
    for name, pct, loss, secs, deaths, missed in BOSSES]}
TOP = sorted((r for r in hub.slot_rows().values() if r['gain'] > 0), key=lambda r: -r['gain'])[:4]
SIM = {'dps': 941.2, 'error': 3.1, 'generated_at': '2026-12-07T20:55:00Z',
       'slots': [{'slot': r['slot'], 'item_id': r['pick'], 'delta': round(r['gain'], 1)} for r in TOP]}
SUCCESS = 'ff' + ADDON['success'].lstrip('#')
GLOVES = 'ff' + hub.QUALITY[hub.ITEMS[hub.BAG_UPGRADE]['quality']].lstrip('#')
LEADS = {
    'guild': {'tab': 'guild', 'label': 'Guild', 'waiting': [{'key': 'tonight', 'line': "Onyxia's Lair 20:00 · In"}],
              'lead': {'icon': 'inv_misc_head_dragon_black', 'line': "Onyxia's Lair · 20:00, in 3 h 12 min",
                       'sub': f'|c{SUCCESS}In|r · Phase 2 ranged, west side'}},
    'gear': {'tab': 'gear', 'label': 'Gear', 'waiting': [{'key': 'bag', 'line': 'Equip Gloves of Holy Might +11.2'}],
             'lead': {'itemId': hub.BAG_UPGRADE, 'line': f'|c{GLOVES}Gloves of Holy Might|r |c{SUCCESS}+11.2 DPS|r',
                      'sub': 'Hands · in your bags'}},
    'rotation': {'tab': 'rotation', 'label': 'Rotation', 'waiting': [],
                 'lead': {'icon': hub.SPELL_ICONS['2643'], 'line': 'Garr: 74% followed, 31 s downtime',
                          'sub': 'Tue 8 Dec · parse 58'}},
    'builds': {'tab': 'builds', 'label': 'Builds', 'waiting': [],
               'lead': {'icon': 'ability_hunter_beasttaming', 'line': 'Beast Mastery raid build',
                        'sub': f'|c{SUCCESS}51 of 51 points on plan|r'}},
}
URGENT = ['guild', 'officer', 'builds', 'gear', 'me']
BASE = {'level': 60, 'record': RECORD, 'pack': {'available': True, 'generated': '2026-12-09', 'stale': False},
        'packDays': 1, 'lastRaid': LAST_RAID, 'raidName': 'Molten Core', 'companion': True, 'syncedAgo': '14 min ago',
        'sim': SIM, 'upgrades': [], 'spec': 'hunter-beast-mastery',
        'leads': [LEADS[k] for k in ('builds', 'gear', 'rotation', 'guild')], 'urgentOrder': URGENT,
        'guild': 'Thornwake Company', 'buildName': 'Beast Mastery raid build'}


def state(**changes):
    out = dict(BASE)
    out.update(changes)
    return {k: v for k, v in out.items() if v is not None}


STATES = {
    'raid-day': state(),
    'no-pack': state(record=None, pack={'available': False, 'absent': True}),
    'no-companion': state(companion=False, lastRaid=None, sim=None, syncedAgo=None, raidName=None,
                          leads=[LEADS['gear'], LEADS['builds']]),
    'stale-no-guild': state(pack={'available': True, 'generated': '2026-12-02', 'stale': True}, packDays=8,
                            record=dict(RECORD, as_of='2026-12-02'), guild=None,
                            leads=[LEADS['gear'], LEADS['rotation'], LEADS['builds']]),
    'new-character': state(level=12, record=None, companion=False, lastRaid=None, sim=None, syncedAgo=None,
                           raidName=None, guild=None, buildName=None, spec=None, leads=[
                               {'tab': 'builds', 'label': 'Builds', 'waiting': [{'key': 'spec', 'line': 'Pick your spec'}]}]),
    'sim-queued': state(queued=True),
}


def lua(value, indent=''):
    if value is None:
        return 'nil'
    if isinstance(value, bool):
        return 'true' if value else 'false'
    if isinstance(value, (int, float)):
        return repr(value)
    if isinstance(value, str):
        return '"' + value.replace('\\', '\\\\').replace('"', '\\"') + '"'
    if isinstance(value, list):
        return '{ ' + ', '.join(lua(v) for v in value) + ' }'
    return '{ ' + ', '.join(f'["{k}"] = {lua(v)}' for k, v in value.items()) + ' }'


def models():
    with tempfile.NamedTemporaryFile('w', suffix='.lua', delete=False) as handle:
        handle.write('return ' + lua(STATES) + '\n')
        path = handle.name
    # LuaJIT, as busted runs the addon (Lua 5.5 refuses the addon's assignments to loop variables).
    run = subprocess.run([os.environ.get('LUA', 'luajit'), str(ROOT / 'design/mocks/addon_me_model.lua'), path], cwd=ROOT,
                         capture_output=True, text=True, check=True)
    return json.loads(run.stdout)


# --- Drawing the model on MeFrame's layout -----------------------------------------------------------------

def wow(text):
    """The client's colour escapes as HTML: |cffRRGGBB...|r."""
    out = esc(text or '')
    out = re.sub(r'\|c[0-9a-fA-F]{2}([0-9a-fA-F]{6})', r'<span style="color:#\1">', out)
    return out.replace('|r', '</span>')


def section(heading, note=None):
    note_html = f'<span class="r">{wow(note)}</span>' if note else ''
    return f'<div class="sh"><span>{esc(heading)}</span>{note_html}</div>'


def score_html(m):
    if m is None:
        return ''
    if m.get('line'):
        return section('Forever Score') + f'<div class="m" style="margin-top:2px">{esc(m["line"])}</div><div class="gap"></div>'
    note = m.get('note') or []
    if m.get('noteWarn') and note:
        note = [f'|cff{ADDON["warning"].lstrip("#")}{note[0]}|r'] + note[1:]
    out = section('Forever Score', SEP.join(note))
    out += (f'<div class="big"><span class="n" style="color:#{m["tierHex"]}">{m["total"]}</span>'
            f'<span style="color:#{m["tierHex"]}">{esc(m["tierLabel"])}</span><span class="m">· {esc(m["toNext"])}</span></div>')
    for p in m['parts']:
        length = round(LAYOUT['trackMax'] * p['track'])
        fill = round(length * p['fill'])
        reason = (f'<span class="m" style="left:{LAYOUT["partLabel"] + length + LAYOUT["reasonGap"]}px">{esc(p["reason"])}</span>'
                  if p.get('reason') else '')
        out += (f'<div class="row"><span class="m" style="left:0">{esc(p["label"])}</span>'
                f'<span class="trk" style="left:{LAYOUT["partLabel"]}px;width:{length}px"><i style="width:{fill}px;'
                f'background:#{p["hex"]}"></i></span>{reason}'
                f'<span style="right:0;text-align:right">{p["points"]} <span class="m">of {p["max"]}</span></span></div>')
    return out + '<div class="gap"></div>'


def cell(col, text, color=None):
    style = f'left:{col["x"]}px;width:{col["width"]}px;' + ('text-align:right;' if col['right'] else '')
    if color:
        style += f'color:{color};'
    return f'<span style="{style}">{text}</span>'


def raid_html(m):
    if m is None:
        return ''
    out = section('Last raid', m.get('note'))
    if m.get('empty'):
        return out + (f'<div class="m" style="margin-top:2px">{esc(m["empty"])}</div>'
                      f'<div class="link" style="margin-top:4px">{esc(m["link"])}</div>')
    cols = {c['key']: c for c in COLUMNS}
    out += '<div class="tb"><div class="r">' + ''.join(cell(c, HEADS[c['key']], ADDON['muted']) for c in COLUMNS) + '</div>'
    for r in m['rows']:
        out += ('<div class="r">' + cell(cols['name'], esc(r['name']))
                + cell(cols['parse'], r['parse'], '#' + r['parseHex'])
                + cell(cols['loss'], esc(r.get('loss') or '-'), ADDON['muted'])
                + cell(cols['died'], r['died'], ADDON['warning'] if r['died'] else ADDON['muted'])
                + cell(cols['consumables'], esc(r['consumables']), ADDON['warning'] if r['missed'] else ADDON['muted'])
                + '</div>')
    s = m['summary']
    full = f'<span class="m">{SEP}</span><span class="link">Full log</span>' if m.get('url') else ''
    out += (f'</div><div class="sum">Median <span style="color:#{s["medianHex"]}">{s["median"]}</span>'
            f'<span class="m">{SEP}{esc(s["rest"])}</span>{full}</div>')
    return out


def now_html(rows):
    if not rows:
        return ''
    out = section('Right now')
    for r in rows:
        ic = (item_icon(r['itemId'], LAYOUT['nowIcon']) if r.get('itemId') else
              icon(r['icon'], LAYOUT['nowIcon'], 'box-shadow:0 0 0 1px #000') if r.get('icon') else '')
        line = wow(r['line'])
        if r.get('urgent') and not r.get('sub'):
            line = f'<span style="color:{hub.CLIENT_GOLD}">{line}</span>'
        sub = f'<div class="m">{wow(r["sub"])}</div>' if r.get('sub') else ''
        flush = '' if ic else ' style="left:0"'
        out += (f'<div class="now"><span class="ni">{ic}</span><div class="tx"{flush}><div>{line}</div>{sub}</div>'
                f'<span class="lk link">{esc(r["label"])}</span></div>')
    return out + f'<div style="height:{LAYOUT["sectionGap"] - (LAYOUT["nowRow"] - 2 * LAYOUT["line"])}px"></div>'


def sim_html(m):
    if m is None:
        return ''
    out = section('Your sim')
    if m.get('dps'):
        err = f' <span class="m" style="font-size:12px;font-weight:400">{esc(m["error"])}</span>' if m.get('error') else ''
        out += f'<div class="dps">{esc(m["dps"])}{err}</div>'
    if m['gains']:
        rows = ''.join(f'<div class="gain">{item_icon(g["itemId"], LAYOUT["gainIcon"])}'
                       f'<span class="nm">{hub.item_name(g["itemId"])}</span><span class="v ok">+{g["gain"]:.1f} DPS</span></div>'
                       for g in m['gains'])
        out += f'<div class="m" style="margin-top:4px">Biggest gains</div><div style="margin-top:2px">{rows}</div>'
    if m.get('date') or m.get('button'):
        btn = f'<span class="btn{" off" if m.get("queued") else ""}">{esc(m["button"])}</span>' if m.get('button') else ''
        out += f'<div class="simft"><span class="m">{esc(m.get("date") or "")}</span>{btn}</div>'
    return out


def board(name, caption, model, dots, tabs=None):
    figures = SEP.join(['Level 60 Hunter' if name != 'new-character' else 'Level 12 Hunter'] + model['figures'])
    left = score_html(model.get('score')) + raid_html(model.get('lastRaid'))
    right = now_html(model.get('rightNow')) + sim_html(model.get('sim'))
    saved = bw.TABS
    bw.TABS = tabs or saved
    try:
        win = bw.window('hunter', 'Velisandre', esc(figures), f'<div class="me">{left}</div>',
                        f'<div class="me">{right}</div>', 'Esc to close', wow(model.get('footer') or ''),
                        strip=('Me', dots))
    finally:
        bw.TABS = saved
    return bw.ux.page(f'<style>{bw.CSS}{CSS}</style><div title="{esc(caption)}">{win}</div>')


NO_GUILD = ['Me', 'Builds', 'Gear', 'Rotation']
BOARDS = {
    'me-raid-day': ('raid-day', 'Velisandre on raid day: the score, the Molten Core night, what waits, the sim.',
                    ('Gear', 'Guild'), None),
    'me-no-pack': ('no-pack', 'No ratings pack installed: the score says where it comes from.', ('Gear', 'Guild'), None),
    'me-no-companion': ('no-companion', 'No companion: the raid says how to bring one in, no sim.', ('Gear',), None),
    'me-stale-no-guild': ('stale-no-guild', 'Not in a guild, the ratings pack 8 days old.', ('Gear',), NO_GUILD),
    'me-new-character': ('new-character', 'A level 12 character with no logs, no companion, no spec yet.', ('Builds',),
                         ['Me', 'Builds']),
    'me-sim-queued': ('sim-queued', 'Sim again pressed: Queued until the companion takes it.', ('Gear', 'Guild'), None),
}


def main(names):
    from playwright.sync_api import sync_playwright
    import os
    data = models()
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000})
        page = ctx.new_page()
        for name in names:
            key, caption, dots, tabs = BOARDS[name]
            page.set_content(board(key, caption, data[key], dots, tabs))
            page.evaluate('document.fonts.ready')
            page.wait_for_timeout(150)
            page.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

"""Review boards for the hub's Rotation tab as built (lane L4, design/specs/2026-11-13-addon-hub.md 4.4).

    python3 design/mocks/addon_hub_rotation.py                 # every board
    python3 design/mocks/addon_hub_rotation.py rotation-60

Writes design/reviews/2026-11-13-addon-hub/l4/<board>.png at device scale 2, the window at its real 740 px
width in the frame the Builds tab draws (addon_builds_window.py): the header, the strip, the left column at
388 and the right at 288, the footer. Each board draws what surfaces/RotationTab.lua draws for its inputs, by
the same rules (RotationTab.stops, priority, played, weightRows), so the review sees the build.

Every character, guild and raid is made up: Velisandre, a level 60 Night Elf hunter, and Torvin, a level 34
dwarf hunter. No real player.

Sources, so nothing is invented: every string from addon/ForeverSixty/Locale.lua; the priority, its levels
and each spell's learned level from Data.lua (rotations, abilityLevels); the stat weights from Data.lua
(weights, dpsPerPoint) and Rules.RATING_PER_PERCENT; spell icons from spellicons.json; the last raid in the
contract's last_raid shape (docs/contracts/addon-v2.md) with the API's own loss phrases
(api/internal/rating/last_raid.go).
"""
import datetime
import os
import re
import sys

import addon_builds as ab
import addon_builds_window as bw
import addon_hub as ah
import addon_ux as ux
from addon_builds import L, esc
from addon_ux import ADDON, ROOT

OUT = ROOT / 'design/reviews/2026-11-13-addon-hub/l4'
DATA_LUA = (ROOT / 'addon/ForeverSixty/Data.lua').read_text()
SPEC = 'hunter-beast-mastery'
RATING_PER_PERCENT = {'hit': 10, 'crit': 14, 'expertise': 10, 'dodge': 12, 'parry': 15, 'block': 5, 'defense': 1,
                      'haste': 10}
HASTE = {'haste', 'melee_haste', 'spell_haste'}
STAT_NAMES = dict(re.findall(r'(\w+) = "([^"]+)"', re.search(r'L\.rotStatNames = \{(.*?)\}', ab.LOCALE, re.S).group(1)))
WEEKDAYS = re.findall(r'"(\w+day)"', re.search(r'L\.rotWeekdays = \{([^}]*)\}', ab.LOCALE).group(1))
MAX_BOSSES, MAX_WEIGHTS = 12, 6
SAME_AS = {'attack_power': 'ranged_attack_power', 'ranged_attack_power': 'attack_power'}
COND_WIDTH = 388 - 16 - 24 - 16

CSS = bw.CSS + f'''
.pr {{ position: relative; display: flex; gap: 8px; min-height: 30px; margin-bottom: 6px; }}
.pr .no {{ width: 16px; flex: none; text-align: center; color: {ADDON['muted']}; line-height: 24px; }}
.pr .tx {{ flex: 1; min-width: 0; }}
.pr .nm {{ display: flex; justify-content: space-between; font-weight: 600; color: {ADDON['strong']}; }}
.pr .nm .at {{ font-weight: 400; color: {ADDON['muted']}; }}
.pr .cd {{ color: {ADDON['muted']}; display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical;
  overflow: hidden; width: {COND_WIDTH}px; }}
.pr.dim .no, .pr.dim .ic, .pr.dim .nm > span:first-child {{ opacity: .55; }}
.step {{ display: flex; align-items: center; gap: 4px; font-family: Site, sans-serif; font-size: 12px; }}
.step .btn {{ width: 20px; height: 20px; padding: 0; justify-content: center; }}
.step .btn.off {{ opacity: .4; background: {ADDON['blue']}; border-color: {ADDON['blue_edge']}; color: {ADDON['strong']}; }}
.step .lvl {{ width: 64px; text-align: center; font-weight: 600; color: {ADDON['strong']}; }}
.tb {{ width: 100%; border-collapse: collapse; table-layout: fixed; }}
.tb td {{ height: 18px; padding: 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.tb .n {{ text-align: right; font-variant-numeric: tabular-nums; }}
.tb .h td {{ color: {ADDON['muted']}; }}
.wa {{ color: {ADDON['warning']}; }}
.s {{ color: {ADDON['strong']}; }}
.link {{ color: {ADDON['gold']}; margin-top: 4px; }}
.btn {{ display: inline-flex; align-items: center; height: 20px; padding: 0 8px; font-weight: 600; flex: none;
  color: {ADDON['strong']}; background: {ADDON['blue']}; border: 1px solid {ADDON['blue_edge']}; white-space: nowrap; }}
'''


def fmt(pattern, *args):
    """string.format as the addon calls it: %d, %s, %.2f and %% read the same in Python."""
    return pattern % args


# --- Data.lua, read the way the addon reads it ------------------------------------------------------

def rotation_lines():
    table = re.search(r'^local rotationLines = \{\n(.*?)^\}', DATA_LUA, re.M | re.S).group(1)
    return [m.groups() for m in re.finditer(r'\{ spellId = (\d+), name = "([^"]*)", condition = "([^"]*)"', table)]


def bands(spec):
    lines = rotation_lines()
    start = DATA_LUA.index('rotations = {')
    block = re.search(rf'\["{re.escape(spec)}"\] = \{{\n(.*?)\n\t\t\}},', DATA_LUA[start:], re.S).group(1)
    out = []
    for level, row in re.findall(r'\{ level = (\d+), lines = \{ ([^}]*) \} \}', block):
        out.append((int(level), [lines[int(i) - 1] for i in re.findall(r'rotationLines\[(\d+)\]', row)]))
    return out


ABILITY_LEVELS = {int(k): int(v) for k, v in re.findall(
    r'\[(\d+)\] = (\d+),', DATA_LUA[DATA_LUA.index('abilityLevels = {'):])}


def shipped_weights(spec):
    row = re.search(rf'\["{re.escape(spec)}"\] = \{{ (\["[^}}]*)\}}', DATA_LUA[DATA_LUA.index('weights = {'):]).group(1)
    return {k: float(v) for k, v in re.findall(r'\["(\w+)"\] = ([\d.]+)', row)}


def dps_per_point(spec, band):
    block = DATA_LUA[DATA_LUA.index('dpsPerPoint = {'):]
    spec_block = re.search(rf'\["{re.escape(spec)}"\] = \{{(.*?)\n\t\t\}},', block, re.S).group(1)
    return float(re.search(rf'\[{band}\] = \{{ reference = "\w+", dps = ([\d.]+)', spec_block).group(1))


# --- RotationTab's rules ----------------------------------------------------------------------------

def stops(spec, level):
    levels, previous = [], None
    for band_level, lines in bands(spec):
        key = [(n, c) for _, n, c in lines]
        if key != previous:
            levels.append(band_level)
        previous = key
    return sorted(set(levels + [level]))


def band_at(spec, level):
    chosen = bands(spec)[0]
    for band in bands(spec):
        if band[0] <= level:
            chosen = band
    return chosen[1]


def condition(text):
    plain = re.sub(r'\s*\([^)]*\)', '', text).strip().rstrip('.')
    return plain[0].upper() + plain[1:]


def priority(spec, shown, level):
    rows = []
    for index, (sid, name, cond) in enumerate(band_at(spec, shown), 1):
        learned = ABILITY_LEVELS.get(int(sid))
        rows.append({'no': index, 'sid': sid, 'name': name, 'cond': condition(cond),
                     'at': learned if learned is not None and learned > level else None})
    return rows


def weekday(date):
    y, m, d = map(int, date.split('-'))
    return WEEKDAYS[(datetime.date(y, m, d).isoweekday()) % 7]


def short_date(date):
    months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
    _, m, d = map(int, date.split('-'))
    return f'{d} {months[m - 1]}'


def loss_cell(boss):
    phrase = boss.get('biggest_loss')
    if not phrase:
        return None
    label = phrase[0].upper() + phrase[1:]
    return fmt(L('meLossSeconds'), label, boss['loss_seconds']) if boss['loss_seconds'] > 0 else label


def played(message):
    rows = []
    for b in message['bosses']:
        a = b.get('adherence')
        rows.append({'name': b['name'], 'followed': round(a * 100) if a is not None else None, 'adh': a,
                     'loss': loss_cell(b)})
    counted = [r['adh'] for r in rows if r['adh'] is not None]
    rows.sort(key=lambda r: (r['adh'] is None, r['adh'] or 0))
    kinds = {}
    for b in message['bosses']:
        if b.get('biggest_loss'):
            k = kinds.setdefault(b['biggest_loss'], {'label': b['biggest_loss'][0].upper() + b['biggest_loss'][1:],
                                                     'seconds': 0, 'bosses': 0})
            k['seconds'] += b['loss_seconds']
            k['bosses'] += 1
    lost = sorted((k for k in kinds.values() if k['seconds'] > 0), key=lambda k: (-k['seconds'], -k['bosses']))
    for k in lost:
        k['counted'] = k['seconds'] == 0
        k['value'] = (fmt(L('rotSeconds'), k['seconds']) if k['seconds'] > 0 else
                      L('rotOneBoss') if k['bosses'] == 1 else fmt(L('rotBosses'), k['bosses']))
    day = weekday(message['date'])
    return {'note': f'{day[:3]} {short_date(message["date"])}', 'day': day,
            'mean': round(sum(counted) / len(counted) * 100) if counted else None, 'counted': len(counted),
            'bosses': rows[:MAX_BOSSES], 'more': max(0, len(rows) - MAX_BOSSES), 'lost': lost}


def weight_rows(weights, dps, note, reference='ranged_attack_power'):
    rows, haste = [], None
    for stat, w in weights.items():
        twin = SAME_AS.get(stat)
        if w <= 0 or (twin and weights.get(twin) == w and stat != reference):
            continue
        if stat in HASTE:
            haste = max(haste or 0, w)
        else:
            rows.append((stat, w / RATING_PER_PERCENT.get(stat, 1)))
    rows.sort(key=lambda r: (-r[1], r[0]))
    top = rows[0][1]
    return {'stats': [(s, STAT_NAMES.get(s, s)) for s, _ in rows[:MAX_WEIGHTS]],
            'rows': [(fmt(L('rotRating'), STAT_NAMES.get(s, s)) if s in RATING_PER_PERCENT else STAT_NAMES.get(s, s),
                      p / top, p * dps if dps else None) for s, p in rows[:MAX_WEIGHTS]],
            'haste': haste * dps if haste and dps else None, 'note': note, 'unit': L('rotDps')}


# --- Drawing -----------------------------------------------------------------------------------------

def section(heading, body, right=''):
    return (f'<div class="sec"><div class="sh"><span>{heading}</span><span class="r">{right}</span></div>'
            f'{body}</div>')


def stepper(shown, level, spec):
    s = stops(spec, level)
    prev_off = ' off' if shown == s[0] else ''
    next_off = ' off' if shown == s[-1] else ''
    return (f'<span class="step"><span class="btn{prev_off}">{esc(L("rotPrevious"))}</span>'
            f'<span class="lvl">{fmt(L("rotLevel"), shown)}</span>'
            f'<span class="btn{next_off}">{esc(L("rotNext"))}</span></span>')


def priority_html(spec, shown, level, followed):
    rows = ''
    for r in priority(spec, shown, level):
        at = f'<span class="at">{fmt(L("rotAtLevel"), r["at"])}</span>' if r['at'] else ''
        rows += (f'<div class="pr{" dim" if r["at"] else ""}"><span class="no">{r["no"]}</span>'
                 f'{ah.icon(ah.SPELL_ICONS[r["sid"]], 24)}<div class="tx"><div class="nm"><span>{esc(r["name"])}</span>'
                 f'{at}</div><div class="cd">{esc(r["cond"])}</div></div></div>')
    title = L('rotPriority')
    if not followed:
        title += f' <span class="m" style="font-family:Site,sans-serif;font-size:12px">{esc(fmt(L("rotFor"), "Beast Mastery"))}</span>'
    head = (f'<div class="sh" style="height:22px"><span>{title}</span>{stepper(shown, level, spec)}</div>'
            f'<div style="height:4px"></div>')
    return f'<div class="sec">{head}{rows}</div>'


def weights_html(w):
    rows = (f'<tr class="h"><td style="width:232px">{L("rotStat")}</td><td class="n" style="width:60px">'
            f'{L("rotWeight")}</td><td class="n">{fmt(L("rotPerPoint"), w["unit"])}</td></tr>')
    for label, scale, value in w['rows']:
        rows += (f'<tr><td>{esc(label)}</td><td class="n s">{scale:.2f}</td>'
                 f'<td class="n m">{f"{value:.2f}" if value is not None else L("rotNone")}</td></tr>')
    haste = (f'<div class="m" style="margin-top:2px">{esc(fmt(L("rotHaste"), w["haste"], w["unit"]))}</div>'
             if w['haste'] else '')
    return section(L('rotWeights'), f'<table class="tb">{rows}</table>{haste}', esc(w['note']))


def played_html(p, me_link=True, companion=True):
    if p is None:
        return section(L('rotPlayed'), f'<div>{esc(L("rotNoRaidLogged" if companion else "meNoRaid"))}</div>')
    rows = (f'<tr class="h"><td style="width:136px">{L("rotBoss")}</td><td class="n" style="width:56px">'
            f'{L("rotFollowed")}</td><td style="padding-left:8px">{L("rotBiggestLoss")}</td></tr>')
    for b in p['bosses']:
        followed = fmt(L('rotPercent'), b['followed']) if b['followed'] is not None else L('rotNone')
        short = re.sub(r' the [A-Z].*$', '', b['name'])
        rows += (f'<tr><td>{esc(short)}</td><td class="n {"s" if b["followed"] is not None else "m"}">{followed}'
                 f'</td><td class="m" style="padding-left:8px">{esc(b["loss"] or L("rotNone"))}</td></tr>')
    rows += (f'<tr><td class="m">{esc(fmt(L("rotAllBosses"), p["counted"]))}</td>'
             f'<td class="n s">{fmt(L("rotPercent"), p["mean"])}</td><td></td></tr>')
    more = f'<div class="m">{fmt(L("rotMore"), p["more"])}</div>' if p['more'] else ''
    out = section(L('rotPlayed'), f'<div style="height:4px"></div><table class="tb">{rows}</table>{more}', esc(p['note']))
    lost = ''.join(f'<tr><td>{esc(k["label"])}</td><td class="n wa" style="width:60px">'
                   f'{esc(k["value"])}</td></tr>' for k in p['lost'])
    if lost:
        out += section(L('rotLost'), f'<table class="tb">{lost}</table>')
    if me_link:
        out += f'<div class="link">{esc(L("rotOnMe"))}</div>'
    return out


def board(name, level, figures, left, right, dots=('Gear', 'Guild')):
    parts = [fmt(L('hubIdentity'), level, 'Hunter')] + list(figures)
    html = bw.window('hunter', name, esc(L('dialogSeparator').join(parts)), left, right, esc(L('windowEscClose')),
                     strip=('Rotation', dots))
    return ux.page(f'<style>{CSS}</style><div>{html}</div>')


# --- The made-up nights --------------------------------------------------------------------------------

def b(name, adherence, loss, seconds, deaths=0):
    return {'name': name, 'percentile': 70, 'adherence': adherence, 'biggest_loss': loss, 'loss_seconds': seconds,
            'deaths': deaths, 'consumables_missed': []}


MOLTEN_CORE = {'date': '2026-12-15', 'bosses': [
    b('Lucifron', 0.82, 'downtime', 18), b('Magmadar', 0.78, 'downtime', 24), b('Gehennas', 0.91, 'low output', 0),
    b('Garr', 0.74, 'downtime', 31), b('Baron Geddon', 0.88, 'dead early', 42, 1), b('Shazzrah', 0.93, 'downtime', 4),
    b('Sulfuron Harbinger', 0.80, 'downtime', 15), b('Golemagg the Incinerator', 0.95, 'standing in fire', 0),
]}


def board_60():
    p = played(MOLTEN_CORE)
    weights = weight_rows(shipped_weights(SPEC), dps_per_point(SPEC, 60), '')
    left = priority_html(SPEC, 60, 60, True) + weights_html(weights)
    figures = (bw.name_of('Beast Mastery', 'raid'), fmt(L('rotFiguresFollowed'), p['mean'], p['day']))
    return board('Velisandre', 60, figures, left, played_html(p))


def board_34_stepped():
    weights = weight_rows(shipped_weights(SPEC), dps_per_point(SPEC, 40), '')
    left = priority_html(SPEC, 40, 34, True) + weights_html(weights)
    figures = (bw.name_of('Beast Mastery', 'leveling'),)
    return board('Torvin', 34, figures, left, played_html(None), dots=('Builds',))


def board_no_build():
    p = played({'date': MOLTEN_CORE['date'], 'bosses': MOLTEN_CORE['bosses'][:5]})
    sim = dict(shipped_weights(SPEC), hit=27.4, agility=1.72, crit=16.9)
    weights = weight_rows(sim, 0.1412, fmt(L('rotWeightsSimOn'), '14 Dec'))
    left = priority_html(SPEC, 60, 60, False) + weights_html(weights)
    figures = (fmt(L('rotFiguresFollowed'), p['mean'], p['day']),)
    return board('Velisandre', 60, figures, left, played_html(p, me_link=False), dots=('Builds',))


BOARDS = {
    'rotation-60': board_60,
    'rotation-34-stepped': board_34_stepped,
    'rotation-no-build': board_no_build,
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
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

"""Mock boards for the signed-in home panel without the Switch character column
(design/specs/2026-10-10-home-signed-in-panel.md).

  python3 design/mocks/gen_home_panel.py [name ...]   -> renders/home-panel*.png

  home-panel          1440, signed in, with alts: the new layout (identity left, three cards right)
  home-panel-one      1440, one character: no "Change character" line, eyebrow "Your character"
  home-panel-2000     2000: the 1344 container stays centred
  home-panel-1024     1024: the stacked layout (identity, then three cards)
  home-panel-phone    390: identity, cards 1-up, timeline, upgrades
  home-panel-before   1440, today's layout (hero 7/12 + Switch character 5/12) for the side-by-side
  home-panel-states   identity, card and upgrade states at the 1440 widths

Every item, source, gain, set DPS and talent-point total is read from data/builds/<active build>
(bis/<spec>.json band entries and items/<class>.json); nothing numeric is typed except the example
characters' own facts (level, how many talent points differ, sync age), which are examples by
definition. The worn item in each upgrade row is a real runner-up from the band's own alternatives
list, and the gain is the ranker's own sim-verified delta, the same rule lib/home/upgrades.ts
applies. The header is gen_nav.py's shipped bar CSS and markup, so the page renders under the real
bar. Screenshots run from web/ so Playwright is found.
"""
import json
import subprocess
import sys
import tempfile
from pathlib import Path

from mocklib import (
    ROOT, WEB, RENDERS, DATA, GOLD, MUTED, TEXT, BODY, CLASS_COLOR, HEAD, SEAL, crest, emblem, data_uri, footer,
)
import gen_nav
gen_nav.CURRENT_DOOR = ''   # the home page is not a door
from gen_nav import example, bar as nav_bar, CSS as NAV_CSS, SPEC_NAMES, CLASS_NAME

BUILD = json.load(open(WEB / 'src/data/active-build.json'))['build']
BUILD_DIR = ROOT / 'data/builds' / BUILD
DATES = json.load(open(WEB / 'src/data/dates.json'))
CATALOG = {s['spec']: s for s in json.load(open(ROOT / 'data/curated/specs.json'))}
TODAY = '2026-10-10'

ALLIANCE = '#6fb1ff'
HORDE = '#ff6b5c'
KILL = '#7fd48a'
RARITY_TEXT = {0: '#9d9d9d', 1: '#ffffff', 2: '#1eff00', 3: '#3d94f0', 4: '#b866f5', 5: '#ff8000'}
RARITY_BORDER = {0: '#9d9d9d', 1: '#ffffff', 2: '#1eff00', 3: '#0070dd', 4: '#a335ee', 5: '#ff8000'}
SLOT_LABEL = {
    'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist',
    'hands': 'Hands', 'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Ring', 'finger2': 'Ring',
    'trinket1': 'Trinket', 'trinket2': 'Trinket', 'main_hand': 'Main hand', 'off_hand': 'Off hand', 'ranged': 'Ranged',
}
ALREADY_SHOWN = 3
MIN_GAIN = 0.095   # a gain that prints as +0.0 is a tie, not an upgrade

COPY = {   # verbatim from the spec's copy table
    'eyebrow_one': 'Your character', 'eyebrow_alts': 'Current character',
    'hint': 'Change character: top right',
    'sync_tail': 'from the addon · gear and talents in sync', 'stale': 'Reopen the addon to refresh your gear.',
    'no_bnet': 'Blizzard serves no data for this realm type yet.',
    'bis': 'Best in slot', 'talents': 'Talents', 'sim': 'Simulator',
    'heading': 'Your upgrades', 'th': ['Slot', 'You wear', 'Best in slot', 'Gain'],
}


# ---------------------------------------------------------------------------- data
def fx(name, spec, level, realm, faction, guild=None, differ=2, synced='12 minutes ago', hours=0.2, build=True, sim=True):
    """An example character. Facts that no data file can supply (level, guild, how many talent points
    differ, sync age) are typed here; everything derived from them is read from the build."""
    return dict(name=name, spec=spec, cls=spec.split('-')[0], level=level, realm=realm, faction=faction, guild=guild,
                differ=differ, synced=synced, hours=hours, build=build, sim=sim)


OBNOX = fx('Obnoxious Yell', 'warrior-fury', 27, 'Living Flame', 'alliance', guild='Midnight Oil')
FROST = fx('Frostbyte', 'mage-frost', 28, 'Stonehearth', 'horde')
TREE = fx('Treewalker', 'druid-balance', 24, 'Living Flame', 'alliance', hours=456, synced='19 days ago')
OAK = fx('Oakheart', 'paladin-protection', 25, 'Living Flame', 'alliance')
QUICK = fx('Quickshot', 'hunter-beast-mastery', 24, 'Skyborne', 'horde', guild='Night Watch', differ=0, synced='3 hours ago', hours=3)
ALTS = [OBNOX, FROST, TREE, OAK]


def band_for(ch):
    f = json.load(open(BUILD_DIR / 'bis' / f"{ch['spec']}.json"))
    lvl = max(20, min(60, ch['level'])) // 10 * 10
    cands = [b for b in f['bands'] if b['band'] == lvl and b['faction'] == ch['faction']]
    return (next((b for b in cands if b['preset'] == 'raid'), None) or next((b for b in cands if b['preset'] == 'bare'), None) or cands[0]), f


_ITEMS = {}


def item(cls, item_id):
    if cls not in _ITEMS:
        _ITEMS[cls] = {i['id']: i for i in json.load(open(BUILD_DIR / 'items' / f'{cls}.json'))['items']}
    return _ITEMS[cls].get(item_id)


def compute(ch):
    """The worn-vs-best-in-slot comparison for an example character: upgrade rows sorted by gain, and the
    slots already best in slot. Mirrors lib/home/upgrades.ts: a worn runner-up with a nonzero sim-verified
    delta is an upgrade of -dps_delta; a delta of 0 is a tie and counts as already best in slot."""
    band, file = band_for(ch)
    ups, already = [], []
    for s in band['slots']:
        if 'item_id' not in s:
            continue
        pick = item(ch['cls'], s['item_id'])
        alt = next((a for a in s.get('alternatives', []) if a.get('dps_delta', 0) <= -MIN_GAIN and item(ch['cls'], a['item_id'])), None)
        if alt and pick:
            ups.append(dict(slot=s['slot'], pick=s, pick_item=pick, worn=item(ch['cls'], alt['item_id']), worn_src=alt, gain=-alt['dps_delta']))
        elif pick:
            already.append(dict(slot=s['slot'], item=pick))
    ups.sort(key=lambda u: -u['gain'])
    total = sum(u['gain'] for u in ups)
    return dict(band=band, file=file, ups=ups, already=already, total=total)


def band_label(b):
    return '60' if b == 60 else f'{b} to {b + 9}'


def race(band):
    return band['race'].replace('-', ' ').title()


def descriptor_line(ch, band=None):
    band = band or band_for(ch)[0]
    return f"Level {ch['level']} {race(band)} {SPEC_NAMES[ch['spec']]} {CLASS_NAME[ch['cls']]}"


def bar_char(ch, **kw):
    return example(ch['name'], ch['spec'], ch['level'], ch['realm'], ch['faction'], 'Addon', ch['synced'], current=True, **kw)


# ---------------------------------------------------------------------------- css
CSS = '''<style>
:root{--ember:#d66e28;--c:#fff}
.page{display:flex;flex-direction:column;background:var(--bg)}
.sky{position:relative;overflow:hidden;background:
 radial-gradient(1px 1px at 12% 18%,rgba(255,214,140,.9),transparent 60%),radial-gradient(1px 1px at 28% 42%,rgba(255,190,110,.8),transparent 60%),
 radial-gradient(1.5px 1.5px at 61% 22%,rgba(255,214,140,.9),transparent 60%),radial-gradient(1px 1px at 74% 55%,rgba(255,190,110,.7),transparent 60%),
 radial-gradient(1px 1px at 88% 30%,rgba(255,214,140,.8),transparent 60%),radial-gradient(1px 1px at 45% 65%,rgba(255,190,110,.6),transparent 60%),
 radial-gradient(ellipse 90% 55% at 50% 100%,rgba(214,110,40,.5),transparent 70%),radial-gradient(ellipse 70% 60% at 50% 0%,rgba(38,64,92,.9),transparent 70%),
 linear-gradient(180deg,#070b12 0%,#0d1522 55%,#1b1410 100%)}
.sky svg{position:absolute;left:0;bottom:0;width:100%;height:160px}
.sky .fade{position:absolute;inset:0;background:linear-gradient(180deg,transparent 55%,rgba(7,9,13,.9))}
.wrap{position:relative;width:100%;max-width:1344px;margin:0 auto;padding:32px 48px}
.eyebrow{display:flex;align-items:center;gap:10px;font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--gold)}
.eyebrow i{display:inline-block;width:28px;height:1px;background:var(--gold)}
.hero{display:grid;grid-template-columns:repeat(12,minmax(0,1fr));gap:32px;align-items:end}
.hero.before{align-items:end}
.id{grid-column:span 5;display:flex;flex-direction:column;gap:16px;min-width:0}
.idrow{display:flex;align-items:center;gap:18px}
.idrow .crest{width:84px;height:84px}
.idtx{display:flex;flex-direction:column;gap:2px;min-width:0}
.nm{margin:0;font-family:var(--font-display);font-weight:700;font-size:34px;line-height:1.1;color:var(--c)}
.nmw{min-height:75px;display:flex;align-items:flex-end}.nm{display:-webkit-box;-webkit-line-clamp:2;-webkit-box-orient:vertical;overflow:hidden}
.nm.old{font-weight:600;font-size:22px}.nmw.old{min-height:0}
.desc{display:flex;flex-direction:column;gap:2px;margin-top:4px}.d1{font-size:15px;color:#c9c2b2}.d2{display:flex;flex-wrap:wrap;align-items:center;column-gap:4px;font-size:14px;color:#c9c2b2}
.desc img{width:16px;height:16px;object-fit:contain}
.desc a{color:#b9b3a4}
.sync{font-size:12px;color:var(--muted);margin:0}
.hint{font-size:12px;color:var(--muted);margin:0}
.nb{font-size:13px;color:var(--muted);margin:0}
.cards{grid-column:span 7;display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:12px}
.card{display:flex;flex-direction:column;gap:6px;padding:14px 16px;border-radius:6px;border:1px solid var(--line);background:linear-gradient(180deg,#131824,#0d111a);box-shadow:inset 0 -1px 0 rgba(229,185,85,.35);color:var(--strong);transition:border-color .12s;min-width:0}
.card:hover,.card.is-hover{border-color:#a8762a}
.card.is-focus{outline:2px solid var(--gold);outline-offset:2px}
.card.is-active{border-color:#a8762a;background:linear-gradient(180deg,#0f131d,#0a0d15)}
.card .lab{font-size:11px;font-weight:700;letter-spacing:.14em;text-transform:uppercase;color:var(--gold)}
.fig{font-family:var(--font-display);font-size:22px;font-weight:700;text-transform:uppercase;letter-spacing:.01em;color:var(--strong);line-height:1.2}
.sub{font-size:12px;color:var(--muted)}.sub.m{font-family:var(--font-mono)}
.card .t15{font-size:15px;font-weight:600;color:var(--strong)}.card .t13{font-size:13px;color:var(--muted)}.card .run{font-size:12px;font-weight:600;color:#b9b3a4}
.sk{background:linear-gradient(90deg,#161c2a,#1c2433,#161c2a);border-radius:3px;display:block}
.tl{display:flex;align-items:stretch;overflow:hidden;background:rgba(13,17,26,.85);border:1px solid var(--line);border-radius:6px;padding:14px 18px}
.tl .c{display:flex;flex-direction:column;gap:4px;min-width:150px;flex:1;padding:4px 12px}
.tl .c:not(:last-of-type){border-right:1px solid var(--line-soft)}.tl .c.past{opacity:.5}.tl .c.past .v{color:var(--muted)}
.tl .k{display:flex;align-items:center;gap:8px;font-family:var(--font-mono);font-size:12px;color:var(--muted)}
.tl .dot{width:8px;height:8px;border-radius:999px;background:var(--gold);box-shadow:0 0 10px rgba(229,185,85,.8)}
.tl .v{font-size:13px;font-weight:600;color:var(--strong)}.tl .n{font-size:12px;color:var(--muted)}
.tl .up{margin-left:auto;display:flex;align-items:center;font-size:12px;color:var(--muted);padding-left:4px}
.content{display:flex;flex-direction:column;gap:48px;width:100%;max-width:1344px;margin:0 auto;padding:32px 48px}
.sec{display:flex;flex-direction:column;gap:14px}
.h2{display:flex;align-items:baseline;justify-content:space-between}
.h2 h2{font-family:var(--font-display);font-size:18px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:var(--strong)}
.h2 a{font-size:12px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:var(--gold)}
.up-body{background:var(--raised);border:1px solid var(--line);border-radius:6px;padding:18px 20px}
.uth,.urow{display:grid;grid-template-columns:90px 1fr 20px 1fr 90px;column-gap:12px;align-items:center}
.uth{padding-bottom:8px;border-bottom:1px solid var(--line-soft);font-size:11px;font-weight:700;text-transform:uppercase;color:var(--muted)}
.urow{padding:10px 0;border-bottom:1px solid var(--line-soft)}.urow:last-child{border-bottom:0}
.slt{font-size:12px;font-weight:700;text-transform:uppercase;color:var(--muted)}
.ucell{display:flex;align-items:center;gap:10px;min-width:0}
.ico{width:36px;height:36px;flex-shrink:0;border:1px solid;border-radius:4px;object-fit:cover}
.itx{display:flex;flex-direction:column;gap:1px;min-width:0;font-size:13px;font-weight:600}
.itx small{font-size:12px;font-weight:400;color:var(--muted)}
.arrow{text-align:center;color:var(--gold)}
.gain{text-align:right;font-family:var(--font-mono);font-size:13px;font-weight:700;color:var(--kill,#7fd48a)}
.already{display:flex;flex-wrap:wrap;align-items:center;gap:8px;margin-top:12px;padding-top:12px;border-top:1px solid var(--line-soft);font-size:13px}
.already b{color:var(--muted);font-weight:600}.already .ico18{width:18px;height:18px;border:1px solid;border-radius:4px;vertical-align:-4px;margin-right:6px}
.dcard{border:1px solid var(--line);border-radius:6px;background:var(--raised);padding:16px;display:flex;flex-direction:column;gap:6px;box-shadow:inset 0 1px 0 rgba(229,185,85,.35)}
.dcard b{font-family:var(--font-display);font-size:15px;color:var(--strong)}.dcard span{font-size:13px;color:var(--muted)}
.swp{display:flex;flex-direction:column;gap:4px;min-height:168px;padding:14px 18px;border-radius:6px;border:1px solid var(--line);background:rgba(13,17,26,.85)}
.swp .hd2{display:flex;align-items:center;justify-content:space-between;margin-bottom:4px}
.swp .hd2 a{font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:#b9b3a4}
.swr{display:flex;align-items:center;gap:12px;min-height:44px;padding:8px 0;border-bottom:1px solid var(--line-soft)}.swr:last-child{border-bottom:0}
.swr .crest{width:36px;height:36px}
.swr .tx{display:flex;flex-direction:column;flex:1;min-width:0}
.swr .n1{font-family:var(--font-display);font-size:14px;font-weight:600;color:var(--c)}
.swr .n2{font-size:12px;color:var(--muted);white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.swr .st{font-family:var(--font-mono);font-size:12px;color:var(--muted)}
.curp{border:1px solid var(--gold);border-radius:999px;padding:2px 8px;font-size:11px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;color:var(--gold)}
.cell{display:flex;flex-direction:column;gap:8px;min-width:0}.cell .cap{font-size:11px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.cell .cnote{font-size:12px;color:var(--muted)}
.stage{background:#0b0f18;border:1px dashed #262e40;border-radius:6px;padding:16px}
@media (max-width:1279px){.id,.cards{grid-column:1/-1}.hero{align-items:stretch;gap:22px}}
@media (max-width:1023px){.idrow .crest{width:64px;height:64px}}
@media (max-width:767px){
 .wrap{padding:28px 18px}.content{padding:32px 18px;gap:32px}
 .cards{grid-template-columns:1fr}
 .idrow{gap:14px}
 .tl{display:grid;grid-template-columns:1fr 1fr;gap:12px}.tl .c{min-width:0;padding:4px 12px;border-right:0!important}
 .tl .up{align-items:flex-end;margin-left:0}
 .uth{display:none}.urow{grid-template-columns:1fr;row-gap:6px}.urow .arrow{transform:rotate(90deg)}.gain{text-align:left}
 .h2{flex-direction:column;align-items:flex-start;gap:4px}.h2 a{font-size:12px}
 .dgrid{grid-template-columns:1fr!important}
}
</style>'''


# ---------------------------------------------------------------------------- pieces
def sky(inner):
    return ('<div class="sky"><svg viewBox="0 0 1440 160" preserveAspectRatio="none" aria-hidden="true">'
            '<path d="M0 160 L0 110 L120 85 L240 120 L380 55 L500 95 L640 45 L760 90 L900 30 L1040 80 L1180 60 L1300 105 L1440 70 L1440 160 Z" fill="#0b0e14"/>'
            '<path d="M0 160 L0 135 L180 120 L320 140 L480 112 L660 132 L840 104 L1000 128 L1200 108 L1440 132 L1440 160 Z" fill="#07090d"/></svg>'
            f'<div class="fade"></div><div class="wrap">{inner}</div></div>')


def logo(faction):
    return data_uri(WEB / 'public/icons/hd/faction' / f'{faction}-logo-512.webp')


def faction_html(f):
    col = ALLIANCE if f == 'alliance' else HORDE
    return f'<span>·</span><img src="{emblem(f)}" alt=""><span style="font-weight:600;color:{col}">{f.title()}</span>'


def identity(ch, alts, old=False, stale=False, no_build=False, no_bnet=False, show_spec=True, show_guild=True, name=None):
    band = band_for(ch)[0]
    color = CLASS_COLOR[ch['cls']]
    line_a = f"Level {ch['level']} {race(band)} {SPEC_NAMES[ch['spec']] + ' ' if show_spec else ''}{CLASS_NAME[ch['cls']]}"
    col = ALLIANCE if ch['faction'] == 'alliance' else HORDE
    mark = emblem(ch['faction']) if old else logo(ch['faction'])   # old = today's unit-frame emblem, for the before board
    line_b = f'<img src="{mark}" alt="" width="16" height="16"><span style="font-weight:600;color:{col}">{ch["faction"].title()}</span>'
    if ch['guild'] and show_guild:
        line_b += f'<span>·</span><span>&lt;<a href="#guild">{ch["guild"]}</a>&gt;</span>'
    line_b += f'<span>·</span><span>{ch["realm"]}-US</span>'
    if old:   # today: one line in the 7/12 column
        desc = f'<span class="d2" style="font-size:15px"><span>{line_a}</span><span>·</span>{line_b}</span>'
    else:
        desc = f'<span class="d1">{line_a}</span><span class="d2">{line_b}</span>'
    sync = ''
    if not no_build:
        sync = f'<p class="sync"><span class="mono">{ch["synced"]}</span> {COPY["sync_tail"]}' + (f'<br>{COPY["stale"]}' if stale else '') + '</p>'
    nb = f'<p class="nb">{COPY["no_bnet"]}</p>' if no_bnet else ''
    hint = f'<p class="hint">{COPY["hint"]}</p>' if (alts and not old) else ''
    eyebrow = COPY['eyebrow_alts'] if alts and not old else COPY['eyebrow_one']
    name_cls = 'nm old' if old else 'nm'
    return (f'<div class="id"><p class="eyebrow"><i></i>{eyebrow}</p>'
            f'<div class="idrow"><img class="crest" src="{crest(ch["cls"])}" alt="" style="--c:{color}">'
            f'<div class="idtx"><div class="nmw{' old' if old else ''}"><h1 class="{name_cls}" style="--c:{color}">{name or ch["name"]}</h1></div><span class="desc">{desc}</span>{sync}{nb}{hint}</div></div></div>')


def card(label, body, cls=''):
    return f'<a class="card {cls}" href="#x"><span class="lab">{label}</span>{body}</a>'


def sk_lines(n, h=14):
    ws = ['75%', '50%', '83%']
    return ''.join(f'<span class="sk" style="height:{h}px;width:{ws[i % 3]}"></span>' for i in range(n))


def bis_card(ch, state='ready', cls=''):
    d = compute(ch)
    n = len(d['ups'])
    lab = COPY['bis']
    if state == 'ready':
        unit = 'score' if d['band'].get('score_unit') == 'tank_score' else 'DPS'
        body = (f'<span class="fig">{n} upgrade{"s" if n != 1 else ""}</span>'
                f'<span class="sub">in your band, {band_label(d["band"]["band"])} · +{d["total"]:.1f} DPS together</span>')
    elif state == 'match':
        body = f'<span class="fig">Best in slot</span><span class="sub">every slot matches the {band_label(d["band"]["band"])} list</span>'
    elif state == 'loading':
        body = sk_lines(3, 16)
    elif state == 'noexport':
        body = '<span class="t13">Not available yet: no gear export for this character. Open the addon once to send it.</span>'
    elif state == 'nospec':
        body = '<span class="t15">Pick a spec</span><span class="sub">Set a spec in the planner to see this.</span>'
    else:
        body = '<span class="t13">No best in slot list published yet for this spec and band.</span>'
    return card(lab, body, cls)


def talents_card(ch, state='ready', cls=''):
    d = compute(ch)
    pts = d['band']['talent_points']
    if state == 'ready':
        fig = 'Optimized' if ch['differ'] == 0 else 'Unoptimized'
        body = f'<span class="fig">{fig}</span><span class="sub">{ch["differ"]} of the {pts} points in the {band_label(d["band"]["band"])} build differ · compare in the planner</span>'
    elif state == 'loading':
        body = sk_lines(3, 16)
    elif state == 'noexport':
        body = '<span class="t13">Not available yet: no talent export for this character. Open the addon once to send it.</span>'
    elif state == 'nospec':
        body = '<span class="t15">Pick a spec</span><span class="sub">Set a spec in the planner to see this.</span>'
    else:
        body = '<span class="t13">No talent build published yet for this spec and band.</span>'
    return card(COPY['talents'], body, cls)


def sim_card(ch, state='ready', cls=''):
    d = compute(ch)
    if state == 'ready':
        dps = round(d['band']['set_dps'] - d['total'])
        body = (f'<span class="fig">{dps} DPS now</span><span class="sub"><span class="mono">{d["band"]["set_dps"]:.1f} DPS</span> at band best in slot</span>')
    elif state == 'loading':
        body = sk_lines(2, 12)
    elif state == 'none':
        body = '<span class="t13">No sim yet.</span><span class="run">Run</span>'
    else:
        body = '<span class="sub">No sim yet.</span>'
    return card(COPY['sim'], body, cls)


def cards(ch):
    return f'<div class="cards" data-region="next-actions">{bis_card(ch)}{talents_card(ch)}{sim_card(ch)}</div>'


def timeline():
    nxt = next(i for i, d in enumerate(DATES) if d['iso'] >= TODAY)
    cells = ''
    for i, d in enumerate(DATES):
        status = 'past' if i < nxt else ''
        dot = '<span class="dot"></span>' if i == nxt else ''
        note = f'<span class="n">{d["note"]}</span>' if d.get('note') else ''
        cells += f'<div class="c {status}"><span class="k">{dot}{d["key"]}</span><span class="v">{d["value"]}</span>{note}</div>'
    return f'<div class="tl">{cells}<div class="up">Updated Oct 10</div></div>'


def icon(it, size_cls='ico'):
    return f'<img class="{size_cls}" src="{data_uri(BUILD_DIR / "icons" / (it["icon"] + ".webp"))}" alt="" style="border-color:{RARITY_BORDER[it["quality"]]}">'


def source_line(pick):
    return f"{pick['source']} · crafted" if pick.get('source_kind') == 'crafted' else pick['source']


def upgrade_row(u):
    w, p = u['worn'], u['pick_item']
    worn_sub = f'<small>{source_line(u["worn_src"])}</small>' if u['worn_src'].get('source') else ''
    return (f'<li class="urow" style="list-style:none"><span class="slt">{SLOT_LABEL[u["slot"]]}</span>'
            f'<span class="ucell">{icon(w)}<span class="itx"><span style="color:{RARITY_TEXT[w["quality"]]}">{w["name"]}</span>{worn_sub}</span></span>'
            f'<span class="arrow">→</span>'
            f'<span class="ucell">{icon(p)}<span class="itx"><span style="color:{RARITY_TEXT[p["quality"]]}">{p["name"]}</span><small>{source_line(u["pick"])}</small></span></span>'
            f'<span class="gain">+{u["gain"]:.1f} DPS</span></li>')


def upgrades(ch, state='ready'):
    d = compute(ch)
    spec_name = SPEC_NAMES[ch['spec']]
    band = band_label(d['band']['band'])
    link = f'<a href="#full">Full list for {spec_name} {band}</a>'
    if state == 'loading':
        body = ''.join(f'<span class="sk" style="height:20px;width:{w}%;margin-bottom:12px"></span>' for w in (75, 50, 83, 66))
    elif state == 'noexport':
        body = '<p style="margin:0;min-height:44px;display:flex;align-items:center;color:var(--muted);font-size:14px">Not available yet: no gear export for this character. Open the addon once to send it.</p>'
    else:
        head = ''.join(f'<span class="{"" if i != 3 else ""}" style="{"text-align:right" if i == 3 else ""}">{t}</span>' if i != 2 else '<span></span>' for i, t in enumerate(['Slot', 'You wear', 'x', 'Best in slot', 'Gain'][:5]))
        head = f'<span>Slot</span><span>You wear</span><span></span><span>Best in slot</span><span style="text-align:right">Gain</span>'
        rows = ''.join(upgrade_row(u) for u in d['ups'])
        shown = d['already'][:ALREADY_SHOWN]
        more = len(d['already']) - len(shown)
        names = ' <span style="color:var(--muted)">·</span> '.join(
            f'<span>{icon(a["item"], "ico18")}<span style="color:{RARITY_TEXT[a["item"]["quality"]]}">{a["item"]["name"]}</span></span>' for a in shown)
        more_html = f' <span style="color:var(--muted)">· {more} more slot{"s" if more != 1 else ""}</span>' if more > 0 else ''
        body = (f'<div class="uth">{head}</div><ul style="margin:0;padding:0">{rows}</ul>'
                f'<p class="already" style="margin-bottom:0"><b>Already best in slot:</b> {names}{more_html}</p>')
    return (f'<section class="sec" id="upgrades"><div class="h2"><h2>{COPY["heading"]}</h2>{link}</div><div class="up-body">{body}</div></section>')


DOORS = [('Planner', 'Every talent tree for 1.60, every race and class, shared by link.'),
         ('Simulator', 'Your DPS, upgrades from any loot table, stat weights.'),
         ('Logs', 'Upload a night or log live with the companion; every fight, every parse.'),
         ('Rankings', 'Guild progression and character parses, per boss.'),
         ('Guides', '28 spec guides, talent builds and rotations for every class.')]


def doors():
    cards_html = ''.join(f'<div class="dcard"><b>{t}</b><span>{s}</span></div>' for t, s in DOORS)
    return (f'<section class="sec"><div class="h2"><h2>Five doors</h2></div>'
            f'<div class="dgrid" style="display:grid;grid-template-columns:repeat(5,minmax(0,1fr));gap:14px">{cards_html}</div></section>')


def switch_panel(current, alts):
    rows = ''
    for ch in [current] + [a for a in alts if a is not current]:
        col = CLASS_COLOR[ch['cls']]
        if ch is current:
            tail = '<span class="curp">Current</span>'
        else:
            n = len(compute(ch)['ups'])
            tail = f'<span class="st">{n} upgrade{"s" if n != 1 else ""}</span>'
        rows += (f'<li class="swr" style="list-style:none;--c:{col}"><img class="crest" src="{crest(ch["cls"])}" alt="" style="--c:{col}">'
                 f'<span class="tx"><span class="n1">{ch["name"]}</span><span class="n2">{descriptor_line(ch)}</span></span>{tail}</li>')
    return (f'<div class="swp"><div class="hd2"><span class="label">Switch character</span><a href="#add">Add one</a></div>'
            f'<ul style="margin:0;padding:0">{rows}</ul></div>')


def header(ch, phone=False):
    return nav_bar(bar_char(ch))


def page_html(ch, alts, width, before=False, phone=False):
    d_ = ch
    if before:
        idb = identity(ch, False, old=True).replace('class="id"', 'class="id" style="grid-column:span 7"')
        left = idb.replace('</div></div></div>', '</div></div>', 1)
        # the old column stacked the three cards under the identity inside the same 7/12 cell
        left = (f'<div style="grid-column:span 7;display:flex;flex-direction:column;gap:16px">'
                f'{identity(ch, False, old=True).replace(" class=\"id\"", " class=\"id\" style=\"gap:16px\"")}'
                f'{cards(ch).replace("grid-column:span 7", "")}</div>')
        hero = (f'<div class="hero before">{left}<div style="grid-column:span 5">{switch_panel(ch, alts)}</div></div>')
        hero = hero.replace('class="cards" data-region', 'class="cards" style="grid-column:auto" data-region')
    else:
        hero = f'<div class="hero">{identity(ch, len(alts) > 1)}{cards(ch)}</div>'
    strip = f'<div style="max-width:1344px;margin:-8px auto 0;width:100%;padding:0 {"18px" if width < 768 else "48px"}">{timeline()}</div>'
    if width < 768 and not before:   # signed in, phone: the dates strip sits under the upgrades table
        body = (header(ch) + sky(hero) + f'<div class="content">{upgrades(ch)}<div style="margin:0 -0px">{timeline()}</div>{doors()}</div>' + footer(phone))
    else:
        body = (header(ch) + sky(hero) + strip + f'<div class="content">{upgrades(ch)}{doors()}</div>' + footer(phone))
    return body


def doc(body, width, extra_css=''):
    head = HEAD.replace('</head>', NAV_CSS + CSS + extra_css + '</head>')
    return head + f'<div class="page" style="width:{width}px;min-height:100vh">{body}</div></body></html>'


# ---------------------------------------------------------------------------- states sheet
def cell(cap, inner, note='', width=None):
    w = f'style="width:{width}px"' if width else ''
    return f'<div class="cell" {w}><span class="cap">{cap}</span><div class="stage">{inner}</div><span class="cnote">{note}</span></div>'


def idwrap(html):
    return f'<div style="display:grid;grid-template-columns:1fr">{html.replace("class=\"id\"", "class=\"id\" style=\"grid-column:auto\"")}</div>'


def states_sheet():
    stale_ch = {**OBNOX, 'synced': '2 days ago', 'hours': 48}
    long_ch = {**OBNOX, 'name': 'Sir Obnoxious Yellington'}
    no_spec_ch = {**fx('Dwarfling', 'hunter-marksmanship', 14, 'Stonehearth', 'alliance')}
    skeleton_id = ('<div class="id"><p class="eyebrow"><i></i>&nbsp;</p><div class="idrow"><span class="sk" style="width:84px;height:84px;border-radius:999px"></span>'
                   '<div class="idtx" style="gap:8px"><div class="nmw"><span class="sk" style="width:240px;height:34px"></span></div><span class="sk" style="width:300px;height:15px"></span><span class="sk" style="width:200px;height:12px"></span></div></div></div>')
    cells = [
        cell('Identity, alts (2 or more characters)', idwrap(identity(OBNOX, True)), 'Eyebrow "Current character" and the one quiet line. Plain text, not a link.', 501),
        cell('Identity, one character', idwrap(identity(QUICK, False)), 'Eyebrow "Your character", no line: nothing to change.', 501),
        cell('Identity, sync older than 24 hours', idwrap(identity(stale_ch, True, stale=True)), 'One added sentence, no alarm colour.', 501),
        cell('Identity, no build yet and Blizzard has no data', idwrap(identity(no_spec_ch, True, no_build=True, no_bnet=True, show_spec=False, show_guild=False)), 'No sync line (never invented). A spec the addon has not sent is left out of the descriptor.', 501),
        cell('Identity, long name', idwrap(identity(OBNOX, True, name='Sir Obnoxious Yellington')), 'The full name always. It wraps to a second line at 34 px; it is never cut with an ellipsis.', 501),
        cell('Identity, session hinted, /v1/me not answered', idwrap(skeleton_id), 'Skeleton at the ready height (shared Skeleton blocks), so nothing moves when the character lands. Painted from the session snapshot instead when one exists.', 501),
        cell('Card, rest', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX)}</div>', 'Border #262e40, gold hairline along the bottom.', 281),
        cell('Card, hover', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX, cls="is-hover")}</div>', 'Border to --gold-deep #a8762a, 120 ms.', 281),
        cell('Card, focus-visible', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX, cls="is-focus")}</div>', '2 px gold outline, 2 px offset. Enter follows the link.', 281),
        cell('Card, active (pressed)', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX, cls="is-active")}</div>', 'Border stays --gold-deep, fill one step darker.', 281),
        cell('Best in slot, loading', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX, "loading")}</div>', 'Three 16 px Skeleton lines, the ready figure\'s height.', 281),
        cell('Best in slot, every slot matches', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX, "match")}</div>', 'A win is shown, not hidden.', 281),
        cell('Best in slot, no gear export', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{bis_card(OBNOX, "noexport")}</div>', 'The real reason, in the character\'s own terms.', 281),
        cell('Talents, optimized', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{talents_card(QUICK)}</div>', '', 281),
        cell('Talents, no spec learned', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{talents_card(OBNOX, "nospec")}</div>', 'Best in slot shows the same two lines.', 281),
        cell('Talents, no list published', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{talents_card(OBNOX, "nolist")}</div>', '', 281),
        cell('Simulator, no sim yet', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{sim_card(OBNOX, "none")}</div>', '"Run" is plain text inside the link, not a button.', 281),
        cell('Simulator, loading', f'<div class="cards" style="grid-template-columns:230px;grid-column:auto">{sim_card(OBNOX, "loading")}</div>', '', 281),
    ]
    grid = ('<div style="display:grid;grid-template-columns:repeat(3,auto);justify-content:start;gap:28px 36px;padding:28px 48px">' + ''.join(cells) + '</div>')
    intro = ('<div style="padding:28px 48px 0"><div class="label" style="color:var(--gold)">Home signed-in panel</div>'
             '<div class="display" style="font-size:24px;font-weight:700;color:var(--strong);margin-top:4px">States</div>'
             '<div style="font-size:12px;color:var(--muted);margin-top:6px">Example characters only. Every cell is the real markup at the 1440 widths: identity 5/12 (about 500 px), a card about 230 px.</div></div>')
    upsec = f'<div style="padding:0 48px 36px;max-width:1344px">{upgrades(OBNOX, "loading")}<div style="height:24px"></div>{upgrades(OBNOX, "noexport")}</div>'
    cap = '<div style="padding:0 48px"><span class="label">Your upgrades: loading (top) and no gear export (bottom). The ready state is in home-panel.png.</span></div>'
    return doc(intro + grid + cap + upsec, 1440)


# ---------------------------------------------------------------------------- boards
BOARDS = {
    'home-panel': (lambda: doc(page_html(OBNOX, ALTS, 1440), 1440), 1440, 900),
    'home-panel-one': (lambda: doc(page_html(QUICK, [QUICK], 1440), 1440), 1440, 900),
    'home-panel-2000': (lambda: doc(page_html(OBNOX, ALTS, 2000), 2000), 2000, 900),
    'home-panel-1024': (lambda: doc(page_html(OBNOX, ALTS, 1024), 1024), 1024, 900),
    'home-panel-phone': (lambda: doc(page_html(OBNOX, ALTS, 390, phone=True), 390), 390, 844),
    'home-panel-before': (lambda: doc(page_html(OBNOX, ALTS, 1440, before=True), 1440), 1440, 900),
    'home-panel-states': (states_sheet, 1440, 900),
}


def main(names):
    scratch = Path(tempfile.mkdtemp(prefix='gen_home_panel_'))
    RENDERS.mkdir(exist_ok=True)
    for name in names:
        build, width, height = BOARDS[name]
        src = scratch / f'{name}.html'
        src.write_text(build())
        out = RENDERS / f'{name}.png'
        subprocess.run(['npx', 'playwright', 'screenshot', '--browser=chromium', f'--viewport-size={width},{height}', '--full-page',
                        '--wait-for-timeout=3000', src.as_uri(), str(out)], check=True, capture_output=True, cwd=WEB)
        print(out)


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

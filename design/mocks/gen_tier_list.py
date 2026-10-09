"""Mock boards for the tier list (design/specs/2026-10-09-tier-list.md).

  python3 design/mocks/gen_tier_list.py   -> renders/tier-list{,-2000,-phone,-phone-tank,-tank,-healer,-states}.html

Every number is read from data/builds/<active build>/bis/<spec>.json at render time (the band
with band == 60, preset == "raid", the chosen faction); spec names, classes and roles come from
data/curated/specs.json. No figure in this file is typed. The tier cut-offs and the tank sort
key are the spec's own rules (constants below), applied to that data.
"""
import json
from datetime import datetime

from mocklib import (
    ROOT, GOLD, MUTED, TEXT, BODY, BORDER, SOFT, RAISED, CLASS_COLOR, PARSE,
    crest, emblem, nav, nav_phone, footer, page, write,
)

BUILD = json.load(open(ROOT / 'web/src/data/active-build.json'))['build']
BIS_DIR = ROOT / 'data/builds' / BUILD / 'bis'
CATALOG = json.load(open(ROOT / 'data/curated/specs.json'))
CLASS_NAME = {c['slug']: c['name'] for c in json.load(open(ROOT / 'data/builds' / BUILD / 'classes.json'))}

ALLIANCE = '#6fb1ff'; HORDE = '#ff6b5c'
YOU = {'class': 'warrior', 'faction': 'alliance', 'battletag': 'Obnoxious Yell', 'spec': 'warrior-fury'}

# Round 2 (player review): no letters. DPS and healer sort on the sim number; tanks on damage taken.
RULERS = (10, 20, 30)        # DPS only: measurement lines, not grades
TIE_PCT = 1.0                # the site's own 1% adoption margin: within it, specs tie
BEST = '#7fd48a'             # the lit "best of this column" colour (mocklib GREEN)
ROLE_LABEL = {'dps': 'DPS', 'tank': 'Tank', 'healer': 'Healer'}


# ---------------------------------------------------------------------------- data
def load_rows(faction: str) -> tuple[list[dict], str, str]:
    rows, stamps, preset_label = [], [], ''
    for s in CATALOG:
        f = json.load(open(BIS_DIR / f"{s['spec']}.json"))
        band = next(b for b in f['bands'] if int(b['band']) == 60 and b['preset'] == 'raid' and b['faction'] == faction)
        stamps.append(f['generated_at'])
        preset_label = f['presets']['raid']['label']
        m = band.get('metrics') or {}
        rows.append({
            'spec': s['spec'], 'class': s['class_slug'], 'slug': s['spec_slug'], 'name': s['name'],
            'role': band.get('role') or 'dps', 'race': band['race'].replace('-', ' ').title(), 'value': band['set_dps'], 'talents': band['talents'],
            'low': bool(band.get('weights_low_confidence')),
            'eh': m.get('effective_health'), 'dtps': m.get('dtps'), 'tps': m.get('tps'),
        })
    stamp = max(stamps)
    return rows, datetime.strptime(stamp, '%Y-%m-%dT%H:%M:%SZ').strftime('%b %-d'), preset_label


def ranked(rows: list[dict], role: str) -> list[dict]:
    pool = [r for r in rows if r['role'] == role]
    if role == 'tank':   # lower damage taken is better: the bar is best / value, the gap is the extra taken
        pool.sort(key=lambda r: (r['dtps'], r['name']))
        top = pool[0]['dtps']
        out = [{**r, 'metric': r['dtps'], 'gap': (r['dtps'] / top - 1) * 100, 'frac': top / r['dtps']} for r in pool]
    else:
        pool.sort(key=lambda r: (-r['value'], r['name']))
        top = pool[0]['value']
        out = [{**r, 'metric': r['value'], 'gap': (1 - r['value'] / top) * 100, 'frac': r['value'] / top} for r in pool]
    for i, r in enumerate(out):
        r['rank'] = i + 1
        prev = out[i - 1]['metric'] if i else None
        r['tie'] = role != 'tank' and prev is not None and abs(1 - r['metric'] / prev) * 100 <= TIE_PCT
    return out


# ---------------------------------------------------------------------------- bits
def crest_img(slug: str, px: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" width="{px}" height="{px}" style="width:{px}px;height:{px}px;--c:{CLASS_COLOR[slug]}">'


def fnum(v: float, role: str) -> str:
    return f'{v:,.0f}' if role == 'tank' else f'{v:,.1f}'


def unit_of(role: str) -> str:
    return {'dps': 'DPS', 'healer': 'HPS', 'tank': 'taken per sec'}[role]


def gap_text(r: dict) -> str:
    if r['rank'] == 1:
        return 'Top' if r['role'] != 'tank' else 'Least'
    return f"+{r['gap']:.1f}%" if r['role'] == 'tank' else f"&minus;{r['gap']:.1f}%"


def hrefs(r: dict) -> tuple[str, str, str]:
    return (f"#bis/{r['class']}/{r['slug']}", f"#guides/{r['class']}/{r['slug']}", f"#planner?spec={r['spec']}")


STYLE = f'''<style>
.trow{{position:relative;display:grid;align-items:center;column-gap:14px;min-height:60px;padding:0 16px 0 14px;border-left:2px solid transparent;border-bottom:1px solid {SOFT}}}
.trow:last-child{{border-bottom:0}}
.trow:hover,.trow.is-hover{{background:#101624;border-left-color:var(--c)}}
.trow:hover .fill,.trow.is-hover .fill{{filter:brightness(1.18)}}
.trow.is-you{{background:color-mix(in srgb,{GOLD} 6%,transparent);border-left-color:{GOLD}}}
.fill{{height:100%;border-radius:3px;background:var(--c);opacity:.9}}
.track{{height:10px;border-radius:3px;background:{SOFT};overflow:hidden}}
.lnk{{display:inline-flex;align-items:center;height:44px;padding:0 6px;font-size:12px;font-weight:700;letter-spacing:.08em;text-transform:uppercase;color:{BODY};border-radius:4px}}
.nm:hover{{text-decoration:underline;text-underline-offset:4px}}.nm:focus-visible{{outline:2px solid {GOLD};outline-offset:2px;border-radius:3px}}
.ruler{{display:flex;align-items:center;gap:10px;height:28px;padding:0 16px;background:#0a0d15;border-bottom:1px solid {SOFT};border-top:1px solid {SOFT}}}
.ruler i{{flex:1;height:1px;background:{BORDER}}}
.lnk.p{{color:{GOLD}}}
.lnk:hover,.lnk.is-hover{{color:{TEXT};text-decoration:underline;text-underline-offset:4px}}
.lnk.p:hover{{color:#f5d27f}}
.lnk:focus-visible,.lnk.is-focus{{outline:2px solid {GOLD};outline-offset:2px}}
.tab{{display:inline-flex;align-items:center;justify-content:center;gap:8px;height:44px;padding:0 14px;border:1px solid {BORDER};border-radius:4px;font-size:13px;color:#b9b3a4;background:transparent}}
.tab .n{{font-size:11px;color:{MUTED}}}
.tab.cur{{border-color:{GOLD};color:{GOLD};background:color-mix(in srgb,{GOLD} 8%,transparent)}}.tab.cur .n{{color:{GOLD}}}
.fac{{font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase}}
.fac.a.cur{{border-color:{ALLIANCE};color:{ALLIANCE};background:color-mix(in srgb,#2f6fd6 15%,transparent)}}
.fac.h.cur{{border-color:{HORDE};color:{HORDE};background:color-mix(in srgb,#c0392b 15%,transparent)}}
.pill-low{{color:{MUTED};background:rgba(154,148,132,.12);border-color:rgba(154,148,132,.3);height:20px;font-size:10px}}
.pill-you{{color:{GOLD};background:rgba(229,185,85,.14);border-color:rgba(229,185,85,.35);height:20px;font-size:10px}}
.sk{{background:linear-gradient(90deg,{SOFT},#232b3d,{SOFT});border-radius:3px}}
</style>'''


def low_note() -> str:
    return (f'<span style="font-size:11px;color:{MUTED};border-bottom:1px dotted {MUTED};align-self:flex-start" '
            f'title="This spec\'s stat weights did not settle cleanly, so its gear pick and number are less firm.">Stat weights less certain</span>')


def you_pill() -> str:
    return '<span class="pill pill-you">Your spec</span>'


def name_cell(r: dict, you: bool, phone: bool = False) -> str:
    col = CLASS_COLOR[r['class']]
    badges = you_pill() if you else ''
    return (f'<div style="display:flex;flex-direction:column;gap:2px;min-width:0"><span style="display:flex;align-items:center;gap:8px;flex-wrap:nowrap">'
            f'<a class="nm display" href="{hrefs(r)[0]}" style="font-size:15px;font-weight:700;color:{col};white-space:nowrap">{r["name"]}</a>{badges}</span>'
            f'<span style="font-size:12px;color:{MUTED}">{CLASS_NAME[r["class"]]} &middot; {r["race"]}</span>{low_note() if r["low"] else ""}</div>')


def links(r: dict, cls: str = '') -> str:
    b, g, p = hrefs(r)
    return (f'<span style="display:flex;gap:10px;justify-content:flex-end"><a class="lnk p" href="{b}">BiS</a>'
            f'<a class="lnk" href="{g}">Guide</a><a class="lnk" href="{p}">Planner</a></span>')


def number_cell(r: dict, role: str) -> str:
    return (f'<span style="display:flex;flex-direction:column;align-items:flex-end;line-height:1.2"><span class="mono" style="font-size:16px;color:{BEST if role == 'tank' and r['rank'] == 1 else TEXT}">{fnum(r["metric"], role)}</span>'
            f'<span class="label" style="font-size:10px;letter-spacing:.08em">{unit_of(role)}</span></span>')


def gap_cell(r: dict) -> str:
    top = r['rank'] == 1
    tie = f'<span style="font-size:10px;color:{MUTED}" title="Within 1% of the spec above: a tie">&asymp; tie</span>' if r['tie'] else ''
    return (f'<span style="display:flex;flex-direction:column;align-items:flex-end;line-height:1.2"><span class="mono" style="font-size:13px;color:{GOLD if top else MUTED};font-weight:{700 if top else 500}">{gap_text(r)}</span>{tie}</span>')


def tank_cells(r: dict, best: dict) -> str:
    def cell(key: str) -> str:
        c, w = (BEST, 700) if r[key] == best[key] else (MUTED, 500)
        return f'<span class="mono" style="font-size:14px;text-align:right;color:{c};font-weight:{w}">{r[key]:,.0f}</span>'
    return cell('eh') + cell('tps')


def cols_for(role: str) -> str:
    base = '32px 36px 248px minmax(120px,1fr) 108px'
    return base + (' 104px 104px 84px 210px' if role == 'tank' else ' 84px 210px')


def desk_row(r: dict, role: str, you: bool, best: dict | None, hover: bool = False, cols: str | None = None) -> str:
    col = CLASS_COLOR[r['class']]
    extra = tank_cells(r, best) if role == 'tank' and best else ''
    cls = 'trow' + (' is-you' if you else '') + (' is-hover' if hover else '')
    return (f'<div class="{cls}" style="grid-template-columns:{cols or cols_for(role)};--c:{col}">'
            f'<span class="mono" style="font-size:13px;color:{MUTED}">{r["rank"]}</span>{crest_img(r["class"], 36)}{name_cell(r, you)}'
            f'<span class="track"><span class="fill" style="display:block;width:{r["frac"] * 100:.1f}%"></span></span>'
            f'{number_cell(r, role)}{extra}{gap_cell(r)}{links(r)}</div>')


def col_header(role: str, cols: str | None = None) -> str:
    heads = ['#', '', 'Spec', {'dps': 'Bar: share of the top spec\'s damage', 'healer': 'Bar: share of the top healer\'s healing', 'tank': 'Bar: how close to the least damage taken'}[role],
             {'dps': 'Damage per second', 'healer': 'Healing per second', 'tank': 'Damage taken per second, lower is better'}[role]]
    heads += ['Effective health (HP)', 'Threat per second', 'More taken than best', ''] if role == 'tank' else ['Gap to the top, %', '']
    cells = ''.join(f'<span class="label" style="font-size:10px;{"text-align:right" if i >= 4 else ""}">{h}</span>' for i, h in enumerate(heads))
    return f'<div style="display:grid;grid-template-columns:{cols or cols_for(role)};column-gap:14px;padding:0 16px 8px 14px;align-items:end">{cells}</div>'


def ordinal(n: int) -> str:
    return f'{n}{"th" if 10 <= n % 100 <= 20 else {1: "st", 2: "nd", 3: "rd"}.get(n % 10, "th")}'


def header(role: str, faction: str, rows: list[dict], preset: str, phone: bool) -> str:
    pad = '22px 18px 20px 18px' if phone else '28px 48px 24px 48px'
    pool = ranked(rows, role)
    me = next((r for r in pool if r['spec'] == YOU['spec']), None)
    mine = next((r for r in rows if r['spec'] == YOU['spec']), None)
    you = ''
    if me:
        top = pool[0]
        word = 'taking the least damage' if role == 'tank' else 'level with the top spec'
        behind = (f'Level with the top spec.' if me['rank'] == 1 and role != 'tank' else 'Takes the least damage.' if me['rank'] == 1
                  else (f'{me["gap"]:.1f}% more damage taken than {top["name"]} {CLASS_NAME[top["class"]]}.' if role == 'tank'
                        else f'{me["gap"]:.1f}% behind {top["name"]} {CLASS_NAME[top["class"]]}.'))
        you = (f'<div style="display:flex;{"flex-direction:column;align-items:stretch" if phone else "align-items:center"};gap:{12 if phone else 16}px;padding:14px 18px;border:1px solid {GOLD};border-radius:6px;'
               f'background:color-mix(in srgb,{GOLD} 6%,transparent);max-width:{"none" if phone else "820px"}">'
               f'<span style="display:flex;align-items:center;gap:12px;flex:1;min-width:0">{crest_img(me["class"], 36)}<span style="display:flex;flex-direction:column;gap:2px">'
               f'<span style="font-size:15px;color:{TEXT}"><b class="display" style="color:{CLASS_COLOR[me["class"]]}">{me["name"]} {CLASS_NAME[me["class"]]}</b> is {ordinal(me["rank"])} of {len(pool)} {ROLE_LABEL[role]} specs.</span>'
               f'<span style="font-size:13px;color:{BODY}">{behind}</span></span></span>'
               f'<a class="btn btn-gold" href="{hrefs(me)[0]}" style="flex-shrink:0">Your best in slot &rarr;</a></div>')
    elif mine:
        you = (f'<a href="#tiers/{mine["role"] if mine["role"] != "dps" else ""}" style="display:flex;align-items:center;gap:10px;min-height:44px;font-size:14px;color:{BODY}">'
               f'{crest_img(mine["class"], 36)}<span><b class="display" style="color:{CLASS_COLOR[mine["class"]]}">{mine["name"]} {CLASS_NAME[mine["class"]]}</b> is on the {ROLE_LABEL[mine["role"]]} list. <span style="color:{GOLD};font-weight:700">See where it stands &rarr;</span></span></a>')
    roles = ''.join(f'<a class="tab{" cur" if k == role else ""}" href="#tiers/{k}" style="{"flex:1" if phone else ""}">{ROLE_LABEL[k]} <span class="n mono">{sum(1 for r in rows if r["role"] == k)}</span></a>' for k in ['dps', 'tank', 'healer'])
    facs = ''.join(f'<a class="tab fac {k[0]}{" cur" if k == faction else ""}" href="#?faction={k}" style="{"flex:1" if phone else ""}"><img src="{emblem(k)}" alt="" style="width:20px;height:20px;object-fit:contain">{k.capitalize()}</a>' for k in ['alliance', 'horde'])
    controls = (f'<div style="display:flex;flex-direction:column;gap:8px">'
                f'<div style="display:flex;gap:6px" role="tablist" aria-label="Role">{roles}</div><div style="display:flex;gap:6px" role="group" aria-label="Faction">{facs}</div></div>'
                if phone else
                f'<div style="display:flex;align-items:center;gap:18px"><div style="display:flex;gap:6px" role="tablist" aria-label="Role">{roles}</div>'
                f'<span style="width:1px;height:32px;background:{BORDER}"></span><div style="display:flex;gap:6px" role="group" aria-label="Faction">{facs}</div></div>')
    return (f'<div style="max-width:1344px;margin:0 auto;width:100%;display:flex;flex-direction:column;gap:{14 if phone else 16}px;padding:{pad}">'
            f'<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Level 60 &middot; {preset}</span>'
            f'<h1 class="display" style="font-size:{26 if phone else 30}px;font-weight:700;color:{TEXT};line-height:1.1">Tier list</h1>'
            f'<p style="margin:0;font-size:14px;color:{BODY};max-width:62ch">Where every spec stands at level 60, from our own sim. Open yours for its best gear, guide and talents.</p>'
            f'{you}{controls}</div>')


def notes(role: str, faction: str, preset: str) -> list[str]:
    """The short lines above the list (round 2): one fact each, in the order a raider asks."""
    fresh = f'This is full {preset} gear. At a fresh 60 the order differs: <a href="#bis">see the leveling bands</a>.'
    race = f'Each spec is simmed as its best {faction.capitalize()} race, named under it.'
    if role == 'dps':
        return ['Damage per second on one target for 180 seconds, with raid buffs and consumables. Cleave, adds, movement and what a spec brings the raid are not counted.',
                fresh, f'{race} <b>&asymp; tie</b> marks a spec within {TIE_PCT:.0f}% of the one above.']
    if role == 'healer':
        return ['Effective healing per second over 300 seconds against a stand-in Phase 1 fight: tank hits and raid-wide pulses, not a named boss. Overhealing is not counted. Mana and raid utility are not ranked here.',
                fresh, f'{race} <b>Gap to the top</b> is less healing in this fight, not a verdict on the class.']
    return ['Sorted by damage taken per second, lower is better: it is the sim\'s outcome against the boss, where effective health is only the hit-point pool going in. The best of each column is green.',
            'The boss is a stand-in level 63 target with one melee swing every 2 seconds, not a named boss; the tank wears raid-ready gear with tank consumables.',
            f'{fresh} {race}']


def honesty(role: str, stamp: str, preset: str, phone: bool) -> str:
    lines = [f'Every spec is simmed in its best {preset} gear with raid buffs and consumables.',
             'One target for 180 seconds. Cleave, add fights, movement and what a spec brings the raid are not counted, so a spec can sit lower here than it plays in your raid.' if role == 'dps'
             else 'Same boss profile as the BiS pages, so the number here is the number on the spec\'s page.']
    items = ''.join(f'<li style="font-size:13px;color:{BODY};line-height:1.5">{l}</li>' for l in lines)
    return (f'<div class="panel" style="padding:18px 22px;display:flex;flex-direction:column;gap:10px;max-width:{"none" if phone else "820px"}"><span class="label">How to read this</span>'
            f'<ul style="margin:0;padding-left:18px;display:flex;flex-direction:column;gap:6px">{items}</ul>'
            f'<span style="display:flex;gap:16px;align-items:center;flex-wrap:wrap;font-size:12px;color:{MUTED}">Updated {stamp}<a href="#sim/specs" style="font-size:12px;font-weight:700">How we check the sim &rarr;</a></span></div>')


def ruler(pct: int) -> str:
    return f'<div class="ruler"><span class="label" style="font-size:10px">{pct}% or more behind the top</span><i></i></div>'


def with_rulers(pool: list[dict], render) -> str:
    out, shown = [], set()
    for r in pool:
        for pct in RULERS:
            if pct not in shown and r['gap'] >= pct:
                shown.add(pct); out.append(ruler(pct))
        out.append(render(r))
    return ''.join(out)


def best_of(pool: list[dict], role: str) -> dict | None:
    if role != 'tank':
        return None
    return {'dtps': min(r['dtps'] for r in pool), 'eh': max(r['eh'] for r in pool), 'tps': max(r['tps'] for r in pool)}


def desktop_list(role: str, pool: list[dict]) -> str:
    best = best_of(pool, role)
    render = lambda r: desk_row(r, role, r['spec'] == YOU['spec'], best)
    rows = with_rulers(pool, render) if role == 'dps' else ''.join(render(r) for r in pool)
    return f'<div>{col_header(role)}<div class="panel" style="overflow:hidden">{rows}</div></div>'


def phone_row(r: dict, role: str, you: bool, best: dict | None) -> str:
    col = CLASS_COLOR[r['class']]
    b, g, p = hrefs(r)
    extra = ''
    if role == 'tank' and best:
        def fig(key: str, label: str) -> str:
            c, w = (BEST, 700) if r[key] == best[key] else (MUTED, 500)
            return f'<span style="font-size:12px;color:{MUTED}">{label} <b class="mono" style="color:{c};font-weight:{w}">{r[key]:,.0f}</b></span>'
        extra = f'<div style="display:flex;gap:18px;padding-left:28px">{fig("eh", "Effective health")}{fig("tps", "Threat/s")}</div>'
    tie = f'<span style="font-size:10px;color:{MUTED}">&asymp; tie</span>' if r['tie'] else ''
    return (f'<div style="--c:{col};padding:12px 14px 0 12px;border-left:2px solid {GOLD if you else "transparent"};border-bottom:1px solid {SOFT};{"background:color-mix(in srgb," + GOLD + " 6%,transparent);" if you else ""}display:flex;flex-direction:column;gap:8px">'
            f'<div style="display:flex;align-items:center;gap:10px"><span class="mono" style="font-size:12px;color:{MUTED};width:18px">{r["rank"]}</span>{crest_img(r["class"], 36)}'
            f'<div style="flex:1;min-width:0">{name_cell(r, you)}</div>{number_cell(r, role)}</div>'
            f'<div style="display:flex;align-items:center;gap:10px;padding-left:28px"><span class="track" style="flex:1"><span class="fill" style="display:block;width:{r["frac"] * 100:.1f}%"></span></span>'
            f'<span style="display:flex;flex-direction:column;align-items:flex-end;line-height:1.2;min-width:56px"><span class="mono" style="font-size:12px;color:{GOLD if r["rank"] == 1 else MUTED}">{gap_text(r)}</span>{tie}</span></div>{extra}'
            f'<div style="display:grid;grid-template-columns:repeat(3,1fr);margin:0 -14px 0 -12px;border-top:1px solid {SOFT}"><a class="lnk p" href="{b}" style="justify-content:center">BiS</a>'
            f'<a class="lnk" href="{g}" style="justify-content:center;border-left:1px solid {SOFT}">Guide</a><a class="lnk" href="{p}" style="justify-content:center;border-left:1px solid {SOFT}">Planner</a></div></div>')


def phone_list(role: str, pool: list[dict]) -> str:
    best = best_of(pool, role)
    render = lambda r: phone_row(r, role, r['spec'] == YOU['spec'], best)
    rows = with_rulers(pool, render) if role == 'dps' else ''.join(render(r) for r in pool)
    return f'<div class="panel" style="overflow:hidden">{rows}</div>'


def board(role: str, faction: str, phone: bool, width: int, file: str, signed_in: bool = True) -> str:
    global YOU
    rows, stamp, preset = load_rows(faction)
    pool = ranked(rows, role)
    saved = YOU
    if not signed_in:
        YOU = {**YOU, 'spec': ''}
    gut = '18px' if phone else '48px'
    top = nav_phone(signed_in, saved) if phone else nav('', signed_in, saved)
    note_html = ''.join(f'<p style="margin:0;font-size:13px;color:{BODY};max-width:{"none" if phone else "112ch"}">{n}</p>' for n in notes(role, faction, preset))
    body = (f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{top}{header(role, faction, rows, preset, phone)}</div>'
            f'<div style="max-width:1344px;margin:0 auto;width:100%;padding:{"20px" if phone else "28px"} {gut} 36px {gut};display:flex;flex-direction:column;gap:{16 if phone else 20}px">'
            f'<div style="display:flex;flex-direction:column;gap:6px">{note_html}</div>'
            f'{phone_list(role, pool) if phone else desktop_list(role, pool)}{honesty(role, stamp, preset, phone)}</div>{footer(phone)}')
    YOU = saved
    return write(file, page(STYLE + body, width)).as_posix()


def states_sheet() -> str:
    rows, stamp, _ = load_rows('alliance')
    pool = ranked(rows, 'dps')
    pick = lambda spec: next(r for r in pool if r['spec'] == spec)
    def cap(t: str) -> str:
        return f'<span class="label" style="color:{GOLD}">{t}</span>'
    def one(t: str, r: dict, hover=False, you=False, focus=False) -> str:
        html = desk_row(r, 'dps', you, None, hover=hover, cols=cols_for('dps'))
        if focus:
            html = html.replace('class="lnk p"', 'class="lnk p is-focus"', 1)
        return f'<div style="display:flex;flex-direction:column;gap:6px">{cap(t)}<div class="panel" style="overflow:hidden">{html}</div></div>'
    def skel() -> str:
        cells = ''.join(f'<div class="trow" style="grid-template-columns:32px 36px 248px 1fr 108px 84px 210px"><span class="sk" style="width:16px;height:12px"></span><span class="sk" style="width:36px;height:36px;border-radius:999px"></span><span class="sk" style="width:120px;height:14px"></span><span class="sk" style="height:10px"></span><span class="sk" style="height:14px"></span><span class="sk" style="height:12px"></span><span class="sk" style="height:14px"></span></div>' for _ in range(6))
        return f'<div style="display:flex;flex-direction:column;gap:6px">{cap("Loading: one 60px skeleton row per spec on this tab (20, 3 or 5), so nothing moves when data lands")}<div class="panel">{cells}</div></div>'
    empty = (f'<div style="display:flex;flex-direction:column;gap:6px">{cap("Error and empty: the list panel keeps its place")}'
             f'<div class="panel" style="min-height:120px;padding:24px;display:flex;flex-direction:column;gap:10px;align-items:flex-start;justify-content:center"><span class="display" style="font-size:16px;color:{TEXT}">The tier list did not load.</span>'
             f'<span style="font-size:13px;color:{BODY}">Check your connection and try again. The BiS pages still work.</span><a class="btn btn-gold" href="#retry">Try again</a></div></div>')
    body = (f'<div style="max-width:1344px;margin:0 auto;padding:36px 48px;display:flex;flex-direction:column;gap:18px">'
            f'<h1 class="display" style="font-size:22px;color:{TEXT}">Tier list: row states</h1>'
            f'{one("Rest", pick("warlock-destruction"))}{one("Hover: row lifts, class-colour edge, bar brightens; spec name and links underline", pick("warlock-destruction"), hover=True)}'
            f'{one("Your spec (signed in): gold edge and tint", pick("warrior-fury"), you=True)}{one("Tie: within 1% of the spec above", pick("warlock-demonology"))}'
            f'{one("Focus-visible on the BiS link: 2px gold outline, 2px offset", pick("mage-fire"), focus=True)}'
            f'{one("Stat weights less certain (weights_low_confidence): plain text under the name, never on the headline", pick("warlock-affliction"))}{skel()}{empty}</div>')
    return write('tier-list-states.html', page(STYLE + body, 1440)).as_posix()


if __name__ == '__main__':
    print(board('dps', 'alliance', False, 1440, 'tier-list.html'))
    print(board('dps', 'horde', False, 1440, 'tier-list-horde.html'))
    print(board('dps', 'alliance', False, 2000, 'tier-list-2000.html', signed_in=False))
    print(board('dps', 'alliance', True, 390, 'tier-list-phone.html'))
    print(board('tank', 'alliance', True, 390, 'tier-list-phone-tank.html'))
    print(board('tank', 'alliance', False, 1440, 'tier-list-tank.html'))
    print(board('healer', 'alliance', False, 1440, 'tier-list-healer.html'))
    print(states_sheet())

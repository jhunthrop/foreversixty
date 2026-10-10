"""Mock boards for the /bis landing page (design/specs/2026-10-10-bis-landing.md).

  python3 design/mocks/gen_bis_landing.py   -> renders/bis-landing{,-signed-out,-2000,-phone,-states,-before}.png

Real data: spec names, roles and classes come from data/curated/specs.json; spec icons are each
tree's first talent icon (web/src/lib/bis/spec-icon.ts); the band story (what each band is mostly
made of, how many picks change between bands), the item counts and the Updated stamp are computed from
data/builds/<active build>/bis/*.json at render time. The "before" board's per-spec figures are the
same band-60 raid figures today's index prints. The CHARACTER (Obnoxious Yell, the picks it wears,
the sync age) is an EXAMPLE, as in gen_nav.py: the board labels it. No figure in this file is a
literal except those example worn counts. HTML goes to a scratch dir; only PNGs land in renders/.
"""
import json
import statistics
import subprocess
import sys
import tempfile
from collections import Counter
from datetime import datetime
from pathlib import Path

from mocklib import (
    ROOT, WEB, RENDERS, GOLD, MUTED, TEXT, BODY, BORDER, SOFT, RAISED, CLASS_COLOR,
    crest, data_uri, nav, nav_phone, footer, page,
)

BUILD = json.load(open(ROOT / 'web/src/data/active-build.json'))['build']
BIS_DIR = ROOT / 'data/builds' / BUILD / 'bis'
ICONS = ROOT / 'data/builds' / BUILD / 'icons'
CATALOG = json.load(open(ROOT / 'data/curated/specs.json'))
CLASSES = json.load(open(WEB / 'src/data/classes.json'))     # the site's own class order
EMBER = '#d66e28'
BANDS = (20, 30, 40, 50, 60)
DEFAULT_BAND = 20            # signed out: launch players are leveling
PAGE_MAX = 1344
EXAMPLE = {'class': 'warrior', 'faction': 'alliance', 'battletag': 'Obnoxious Yell', 'spec': 'warrior-fury', 'level': 47}
WORN = {40: 6, 60: 14}
UNDER_20 = 12                 # the character-under-20 state's level
SLOT_WORD = {'trinket': 'Trinkets', 'finger': 'Rings', 'main_hand': 'Main hand', 'off_hand': 'Off hand', 'ranged': 'Ranged'}            # EXAMPLE gear counts for the example character (labelled on the board)
KIND_WORDS = {'dungeon': 'dungeons', 'crafted': 'crafting', 'vendor': 'vendors', 'quest': 'quests',
              'world_drop': 'world drops', 'rep': 'reputation rewards', 'pvp': 'PvP rewards', 'world': 'world drops'}
ROLE_TAG = {'tank': 'Tank', 'healer': 'Healer'}               # DPS is the unlabelled default

FLAT = {f: data_uri(WEB / f'public/icons/hd/faction/{f}-logo-512.webp') for f in ('alliance', 'horde')}


# ---------------------------------------------------------------------------- data
def load_file(spec: str) -> dict:
    return json.load(open(BIS_DIR / f'{spec}.json'))


FILES = {s['spec']: load_file(s['spec']) for s in CATALOG}


def band_entry(spec: str, band: int, faction: str = 'alliance') -> dict:
    """The headline band: raid-ready where the file has it (level 60), bare below (presets.ts selectBand)."""
    cands = [b for b in FILES[spec]['bands'] if int(b['band']) == band and b['faction'] == faction]
    return next((b for b in cands if b['preset'] == 'raid'), cands[0])


def sourced(spec: str, band: int) -> int:
    return sum(1 for s in band_entry(spec, band)['slots'] if s.get('source_kind'))


def band_label(band: int) -> str:
    return '60' if band == 60 else f'{band} to {band + 9}'


def band_of_level(level: int) -> int:
    return min(60, max(20, level // 10 * 10))


def band_story() -> dict[int, dict]:
    """Per band, across every written spec: the two source kinds the picks mostly come from, and the
    median number of picks that change from the band before (the file's own new_at_band)."""
    out = {}
    for band in BANDS:
        kinds, news, totals = Counter(), [], []
        for s in CATALOG:
            b = band_entry(s['spec'], band)
            kinds.update(w for w in (KIND_WORDS.get(x.get('source_kind', '')) for x in b['slots']) if w)
            news.append(len(b['new_at_band']))
            totals.append(sourced(s['spec'], band))
        top = [k for k, _ in kinds.most_common(2)]
        out[band] = {'sources': f'{top[0]}, {top[1]}', 'new': round(statistics.median(news)), 'of': round(statistics.median(totals))}
    previous, previous_band = None, None
    for band in BANDS:       # "Top sources" only where they differ from the band before (the first band always); else "Same sources as {previous band}"
        same = out[band]['sources'] == previous
        out[band]['caption'] = None if same else out[band]['sources']
        out[band]['same_as'] = band_label(previous_band) if same else None
        previous, previous_band = out[band]['sources'], band
    return out


STORY = band_story()
# "From one band to the next, about N of M picks change": the median across the 30, 40 and 50 bands
# (bare to bare, so the raid preset's own band-60 list does not distort it).
TRANSITION = (round(statistics.median(STORY[b]['new'] for b in (30, 40, 50))), round(statistics.median(STORY[b]['of'] for b in (30, 40, 50))))
UPDATED = datetime.strptime(max(f['generated_at'] for f in FILES.values()), '%Y-%m-%dT%H:%M:%SZ').strftime('%b %-d')
CLASS_NAME = {c['slug']: c['name'] for c in CLASSES}
RAID_LABEL = next(iter(FILES.values()))['presets']['raid']['label']


def spec_icon_stem(s: dict) -> str:
    """The planner's tab icon: TreeTabs.svelte reads data/curated/specs.json `icon` by class and tree."""
    return s['icon']


# ---------------------------------------------------------------------------- copy (verbatim in the spec, section 6)
def kinds_caption(sources: str) -> str:
    return f'Top sources: {sources}.'


COPY = {
    'eyebrow': 'Levels 20 to 60',
    'h1': 'Best in slot',
    'job': 'The best gear you can wear at your level, for every spec. Open yours.',
    'updated': f'Updated {UPDATED}',
    'tier_q': 'Which spec is strongest?',
    'tier_link': 'Tier list →',
    'lead_label': 'Your character',
    'lead_cta': lambda spec, band: f'{spec} gear for {band_label(band)} →',
    'worn_label': 'Best-in-slot gear you wear',
    'worn_label_stale': 'Gear you wore at the last sync',
    'count': lambda n, total: (n, total),
    'missing_note': lambda words: f'{words} have no pick at this band.',
    'nosync': 'Add the addon to see how many of these slots you already wear.',
    'nosync_link': 'Get the addon →',
    'stale_tail': 'Log out or /reload in the game; the companion then syncs it',
    'under20': 'Best in slot starts at 20.',
    'grid_out': 'Pick a class and spec',
    'grid_in': 'Another class or spec',
    'strip_head': 'Spec links open at this band',
    'strip_note': 'Each list is the best you can wear inside one band; gear changes about every ten levels. 60: full ' + RAID_LABEL[0].lower() + RAID_LABEL[1:] + ' gear.',
    'same': lambda band: f'Same sources as {band}',
    'band_you': 'you',
    'caption': kinds_caption,
    'yours': 'Yours',
}

# ---------------------------------------------------------------------------- bits
STYLE = f'''<style>
.crest{{--c:{GOLD}}}
.card{{position:relative;background:linear-gradient(180deg,color-mix(in srgb,var(--c) 7%,#0d111a) 0%,#0d111a 120px);border:1px solid {BORDER};border-radius:6px;overflow:hidden}}
.card::before{{content:"";position:absolute;left:0;right:0;top:0;height:3px;background:var(--c)}}
.spr{{display:flex;align-items:center;gap:12px;min-height:48px;padding:0 16px 0 14px;border-left:2px solid transparent;border-top:1px solid {SOFT};color:{TEXT};font-size:15px;font-weight:600}}
.spr.tight{{min-height:44px}}
.spr .nm{{flex:1;min-width:0}}
.spr .tag{{font-size:11px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:{MUTED}}}
.spr .go{{color:{MUTED};opacity:0;transition:opacity .12s;flex-shrink:0}}
.spr:hover,.spr.is-hover{{background:#101624;border-left-color:var(--c)}}
.spr:hover .go,.spr.is-hover .go{{opacity:1;color:{GOLD}}}
.spr:hover .nm,.spr.is-hover .nm{{text-decoration:underline;text-underline-offset:4px}}
.spr:focus-visible,.spr.is-focus{{outline:2px solid {GOLD};outline-offset:-2px;background:#101624}}
.spr:active,.spr.is-active{{background:#0a0d15;border-left-color:var(--c)}}
.spr:active .nm,.spr.is-active .nm{{text-decoration:underline;text-underline-offset:4px;color:#fff}}
.spr.is-you{{background:color-mix(in srgb,{GOLD} 6%,transparent);border-left-color:{GOLD}}}
.pb{{display:flex;align-items:center;justify-content:center;gap:6px;position:relative;height:44px;border:1px solid {BORDER};border-radius:4px;font-family:var(--font-mono);font-size:13px;font-weight:500;color:#b9b3a4;background:transparent;white-space:nowrap}}
.pb .yw{{font-family:var(--font-body);font-size:10px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;opacity:.75}}
.pb .dot{{position:absolute;top:5px;right:5px;width:7px;height:7px;border-radius:999px;background:{GOLD}}}
.pb.sm{{font-size:12px;letter-spacing:-.04em;padding:0}}
.pb.sm.tabl{{letter-spacing:0}}
.band{{min-width:0;overflow:hidden}}
.pb:hover,.pb.is-hover{{color:{TEXT};border-color:#4a4030}}
.pb:focus-visible,.pb.is-focus{{outline:2px solid {GOLD};outline-offset:2px}}
.pb:active,.pb.is-active{{background:#0a0d15;color:#fff;border-color:#4a4030}}
.pb.cur{{border-color:{GOLD};color:{GOLD};background:color-mix(in srgb,{GOLD} 8%,transparent)}}
.pb.cur:hover,.pb.cur.is-hover{{border-color:#f5d27a;color:#f5d27a;background:color-mix(in srgb,{GOLD} 14%,transparent)}}
.pb.cur:active,.pb.cur.is-active{{background:color-mix(in srgb,{GOLD} 4%,transparent)}}
.cta{{height:48px;background:linear-gradient(180deg,#f0cc6c,#c99a3a);border:1px solid #a8762a;color:#1a1408;font-size:13px;box-shadow:0 0 18px rgba(229,185,85,.18)}}
.cta:hover,.cta.is-hover{{background:linear-gradient(180deg,#f5d27a,#d4a444);color:#1a1408}}
.cta:focus-visible,.cta.is-focus{{outline:2px solid {GOLD};outline-offset:2px}}
.cta:active,.cta.is-active{{background:linear-gradient(180deg,#d9b55c,#b88a32)}}
.idl{{display:flex;align-items:center;gap:16px;color:inherit;min-width:0}}
.idl:hover .idn,.idl.is-hover .idn{{text-decoration:underline;text-underline-offset:4px}}
.idl:focus-visible{{outline:2px solid {GOLD};outline-offset:4px;border-radius:6px}}
.tl{{color:{GOLD};font-weight:700}}.tl:hover{{text-decoration:underline;text-underline-offset:4px;color:#f5d27a}}
.track{{height:6px;border-radius:3px;background:{SOFT};overflow:hidden}}.fillg{{display:block;height:100%;background:{GOLD};border-radius:3px}}
.sk{{background:linear-gradient(90deg,{SOFT},#232b3d,{SOFT});border-radius:3px;display:block}}
.pill-you{{color:{GOLD};background:rgba(229,185,85,.14);border-color:rgba(229,185,85,.35);height:20px;font-size:10px}}
</style>'''


def crest_img(slug: str, px: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" width="{px}" height="{px}" style="width:{px}px;height:{px}px;--c:{CLASS_COLOR[slug]}">'


def icon_img(stem: str, px: int) -> str:
    return (f'<img src="{data_uri(ICONS / (stem + ".webp"))}" alt="" width="{px}" height="{px}" '
            f'style="width:{px}px;height:{px}px;border-radius:4px;border:1px solid {BORDER};flex-shrink:0;object-fit:cover">')


def arrow() -> str:
    return '<svg class="go" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12h14M13 6l6 6-6 6"/></svg>'


def gutter(width: int) -> int:
    return 18 if width < 600 else 32 if width < 1024 else 48


def inner(width: int, pad_y: str) -> str:
    return f'max-width:{PAGE_MAX}px;margin:0 auto;width:100%;padding:{pad_y};padding-inline:{gutter(width)}px'


# ---------------------------------------------------------------------------- regions
def hero(width: int) -> str:
    phone = width < 600
    stamp = f'<span style="font-size:12px;color:{MUTED}">{COPY["updated"]}</span>'
    h1 = f'<h1 class="display" style="font-size:{26 if phone else 30}px;font-weight:700;color:{TEXT};line-height:1.1">{COPY["h1"]}</h1>'
    head = f'<div style="display:flex;align-items:baseline;justify-content:space-between;gap:16px">{h1}{"" if phone else stamp}</div>'
    tier = (f'<span style="font-size:14px;color:{MUTED}">{COPY["tier_q"]} <a class="tl" href="#tiers">{COPY["tier_link"]}</a></span>')
    return (f'<div style="{inner(width, "22px 0 20px" if phone else "28px 0 24px")};display:flex;flex-direction:column;gap:{10 if phone else 12}px">'
            f'<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>{COPY["eyebrow"]}</span>'
            f'{head}<p style="margin:0;font-size:14px;color:{BODY};max-width:62ch">{COPY["job"]}</p>{tier}{stamp if phone else ""}</div>')


def missing_note(spec: str, band: int) -> str:
    names = []
    for slot in (s['slot'] for s in band_entry(spec, band)['slots'] if not s.get('source_kind')):
        word = SLOT_WORD.get(slot.rstrip('12'), slot.replace('_', ' ').capitalize())
        if word not in names:
            names.append(word)
    return COPY['missing_note'](' and '.join(names)) if names else ''


def lead(state: str, width: int) -> str:
    """The signed-in lead. state: synced | nosync | stale | l60 | under20 | loading."""
    phone = width < 1024
    s = next(x for x in CATALOG if x['spec'] == EXAMPLE['spec'])
    cls = EXAMPLE['class']; col = CLASS_COLOR[cls]
    level = 60 if state == 'l60' else UNDER_20 if state == 'under20' else EXAMPLE['level']
    band = band_of_level(level)
    total = sourced(s['spec'], band)
    worn = WORN.get(band, 0)
    cols = '' if phone else 'grid-template-columns:minmax(0,1.15fr) minmax(0,1fr) auto;'
    box = (f'display:grid;gap:{16 if phone else 28}px;align-items:center;padding:{"16px" if phone else "18px 22px"};border:1px solid {GOLD};border-radius:6px;'
           f'background:{RAISED};min-height:{0 if phone else 108}px;{cols}')
    if state == 'loading':
        return (f'<div style="{box}"><span style="display:flex;align-items:center;gap:16px"><span class="sk" style="width:64px;height:64px;border-radius:999px"></span>'
                f'<span style="display:flex;flex-direction:column;gap:8px"><span class="sk" style="width:90px;height:11px"></span><span class="sk" style="width:200px;height:22px"></span><span class="sk" style="width:240px;height:13px"></span></span></span>'
                f'<span style="display:flex;flex-direction:column;gap:10px"><span class="sk" style="width:150px;height:11px"></span><span class="sk" style="width:110px;height:26px"></span><span class="sk" style="width:100%;height:6px"></span></span>'
                f'<span class="sk" style="width:{"100%" if phone else "250px"};height:48px;border-radius:4px"></span></div>')
    age = '19 days ago' if state == 'stale' else '12 minutes ago'
    tail = f' &middot; <span style="color:{EMBER}">{COPY["stale_tail"]}</span>' if state == 'stale' else ''
    sync = '' if state == 'nosync' else f'<span style="font-size:12px;color:{MUTED}">Addon &middot; {age}{tail}</span>'
    faction = (f'<img src="{FLAT[EXAMPLE["faction"]]}" alt="" style="width:16px;height:16px;object-fit:contain;flex-shrink:0">')
    desc = (f'<span style="display:flex;align-items:center;gap:6px;font-size:13px;color:{BODY};flex-wrap:wrap">{s["name"]} {CLASS_NAME[cls]}'
            f'<span style="color:{MUTED}">&middot;</span>Level {level}<span style="color:{MUTED}">&middot;</span>{faction}Alliance</span>')
    href = f'#bis/{cls}/{s["spec_slug"]}'
    crest_px = 56 if phone else 64
    ident = (f'<span style="display:flex;flex-direction:column;gap:3px;min-width:0"><span class="label" style="color:{GOLD}">{COPY["lead_label"]}</span>'
             f'<span class="display idn" style="font-size:22px;font-weight:700;color:{col};line-height:1.15">{EXAMPLE["battletag"]}</span></span>')
    identity = (f'<span style="display:flex;flex-direction:column;gap:6px;min-width:0"><a class="idl" href="{href}">{crest_img(cls, crest_px)}{ident}</a>'
                f'<span style="display:flex;flex-direction:column;gap:3px;padding-left:{crest_px + 16}px">{desc}{sync}</span></span>')
    if state == 'under20':
        meter = f'<span style="font-size:14px;color:{TEXT}">{COPY["under20"]}</span>'
    elif state == 'nosync':
        meter = (f'<span style="display:flex;flex-direction:column;align-items:flex-start;gap:2px;font-size:13px;color:{BODY}">{COPY["nosync"]}'
                 f'<a class="tl" href="#setup" style="display:inline-flex;align-items:center;min-height:44px;font-size:13px">{COPY["nosync_link"]}</a></span>')
    else:
        label = COPY['worn_label_stale'] if state == 'stale' else COPY['worn_label']
        note = missing_note(s['spec'], band)
        note_html = f'<span style="font-size:12px;color:{MUTED}">{note}</span>' if note else ''
        meter = (f'<span style="display:flex;flex-direction:column;gap:6px"><span class="label">{label}</span>'
                 f'<span class="mono" style="font-size:26px;font-weight:500;color:{TEXT};line-height:1">{worn}<span style="font-size:15px;color:{MUTED}"> of {total} slots</span></span>'
                 f'<span class="track"><span class="fillg" style="width:{worn / total * 100:.1f}%"></span></span>{note_html}</span>')
    cta = f'<a class="btn cta" href="{href}" style="{"width:100%" if phone else "min-width:250px"}">{COPY["lead_cta"](s["name"], band)}</a>'
    return f'<div style="{box}">{identity}{meter}{cta}</div>'


def pill(b: int, selected: int, you_band: int | None, width: int, state: str = '') -> str:
    """The character's band: the small word `you` inside the pill at 480 and up, a small gold dot in the corner below 480."""
    mark = ''
    if b == you_band:
        mark = '<i class="dot"></i>' if width < 480 else f'<span class="yw">{COPY["band_you"]}</span>'
    cls = ' '.join(x for x in ('pb', 'cur' if b == selected else '', 'sm' if width < 900 else '', state) if x)
    return f'<a class="{cls}" href="#band-{b}" aria-current="{"true" if b == selected else "false"}"><span>{band_label(b)}</span>{mark}</a>'


def caption_html(b: int, size: int = 12) -> str:
    """Every pill has a caption: its top sources where they differ from the band before, else a muted `Same sources as`."""
    story = STORY[b]
    if story['caption']:
        return f'<span style="font-size:{size}px;color:{BODY};line-height:1.35">{COPY["caption"](story["caption"])}</span>'
    return f'<span style="font-size:{size}px;color:{MUTED};line-height:1.35">{COPY["same"](story["same_as"])}</span>'


def strip(width: int, selected: int, you_band: int | None) -> str:
    """The grid's own header row: one compact row at 900 and up; the stacked form below."""
    stack = width < 900
    head = f'<span style="font-size:13px;font-weight:600;color:{TEXT}">{COPY["strip_head"]}</span>'
    note = f'<span style="font-size:13px;color:{MUTED};line-height:1.45">{COPY["strip_note"]}</span>'
    if stack:
        pad = '14px 10px' if width < 480 else '14px 16px'
        return (f'<div class="panel" style="padding:{pad};display:flex;flex-direction:column;gap:10px">{head}'
                f'<div style="display:grid;grid-template-columns:repeat(5,1fr);gap:3px">{"".join(pill(b, selected, you_band, width) for b in BANDS)}</div>{caption_html(selected, 13)}{note}</div>')
    cols = ''.join(f'<div style="display:flex;flex-direction:column;gap:6px;min-width:0">{pill(b, selected, you_band, width)}{caption_html(b)}</div>' for b in BANDS)
    return (f'<div class="panel" style="padding:14px 18px;display:grid;grid-template-columns:440px minmax(0,1fr);gap:28px;align-items:start">'
            f'<div style="display:flex;flex-direction:column;gap:4px">{head}{note}</div>'
            f'<div style="display:grid;grid-template-columns:repeat(5,1fr);gap:12px">{cols}</div></div>')


def spec_row(s: dict, you: bool, tight: bool, state: str = '') -> str:
    col = CLASS_COLOR[s['class_slug']]
    tag = (f'<span class="pill pill-you">{COPY["yours"]}</span>' if you else
           (f'<span class="tag">{ROLE_TAG[s["role"]]}</span>' if s['role'] in ROLE_TAG else ''))
    cls = ' '.join(x for x in ('spr', 'tight' if tight else '', 'is-you' if you else '', state) if x)
    return (f'<a class="{cls}" href="#bis/{s["class_slug"]}/{s["spec_slug"]}" style="--c:{col}">{icon_img(spec_icon_stem(s), 28)}'
            f'<span class="nm">{s["name"]}</span>{tag}{arrow()}</a>')


def class_card(c: dict, signed_in: bool, phone: bool) -> str:
    slug = c['slug']; col = CLASS_COLOR[slug]
    rows = ''.join(spec_row(s, signed_in and s['spec'] == EXAMPLE['spec'], phone) for s in CATALOG if s['class_slug'] == slug)
    head = (f'<div style="display:flex;align-items:center;gap:14px;padding:{"12px 16px 10px" if phone else "18px 18px 16px"}">{crest_img(slug, 44 if phone else 56)}'
            f'<h2 class="display" style="font-size:18px;font-weight:700;color:{col};letter-spacing:.02em">{c["name"]}</h2></div>')
    return f'<section class="card" style="--c:{col}">{head}{rows}</section>'


def grid(width: int, signed_in: bool, selected: int, you_band: int | None) -> str:
    phone = width < 600
    title = COPY['grid_in'] if signed_in else COPY['grid_out']
    heading = f'<h2 class="display" style="font-size:18px;font-weight:700;letter-spacing:.1em;text-transform:uppercase;color:{TEXT}">{title}</h2>'
    cols = 1 if phone else 2 if width < 1024 else 3
    cards = ''.join(class_card(c, signed_in, phone) for c in CLASSES if any(s['class_slug'] == c['slug'] for s in CATALOG))
    return (f'<div style="display:flex;flex-direction:column;gap:12px">{heading}{strip(width, selected, you_band)}'
            f'<div style="display:grid;grid-template-columns:repeat({cols},minmax(0,1fr));gap:{12 if phone else 16}px;align-items:start">{cards}</div></div>')


# ---------------------------------------------------------------------------- boards
def wide_chrome(html: str) -> str:
    """Nav and footer keep their content in the 1344 column above 1440 (nav spec 2.4)."""
    return html.replace('padding:18px 48px', f'padding:18px max(48px,calc((100% - {PAGE_MAX}px)/2 + 48px))').replace(
        'padding:28px 48px', f'padding:28px max(48px,calc((100% - {PAGE_MAX}px)/2 + 48px))')


def board(width: int, signed_in: bool) -> str:
    phone_bar = width < 1024
    top = nav_phone(signed_in, EXAMPLE) if phone_bar else wide_chrome(nav('BiS', signed_in, EXAMPLE))
    if not phone_bar and width < 1280:   # the mocklib bar is a stand-in; the shipped one shrinks per nav spec 2.4
        top = f'<style>.nav{{gap:12px!important;padding:14px 32px!important}}.nav a{{font-size:11px!important;letter-spacing:.06em!important;white-space:nowrap}}.nav span{{white-space:nowrap}}</style>' + top
    you_band = band_of_level(EXAMPLE['level']) if signed_in else None
    selected = you_band or DEFAULT_BAND
    body = (lead('synced', width) if signed_in else '') + grid(width, signed_in, selected, you_band)
    main = f'<div style="{inner(width, "4px 0 40px")};display:flex;flex-direction:column;gap:{16 if width < 600 else 22}px">{body}</div>'
    foot = footer(True) if phone_bar else wide_chrome(footer(False))
    return page(STYLE + f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{top}{hero(width)}</div>{main}{foot}', width)


def states_board() -> str:
    cap = lambda t: f'<span class="label" style="color:{GOLD}">{t}</span>'
    row = lambda t, html, w=None: f'<div style="display:flex;flex-direction:column;gap:8px;{f"width:{w}px;" if w else ""}">{cap(t)}{html}</div>'
    pick = lambda spec: next(s for s in CATALOG if s['spec'] == spec)
    wrap = lambda html: f'<div class="card" style="--c:{CLASS_COLOR["warrior"]};padding-top:3px">{html}</div>'
    fury, bear, resto = pick('warrior-fury'), pick('druid-feral-bear'), pick('druid-restoration')
    rows = ''.join(row(t, wrap(spec_row(fury, y, False, st)), 420) for t, y, st in (
        ('Spec row, rest', False, ''),
        ('Spec row, hover: background and underline only, arrow lights gold; nothing moves', False, 'is-hover'),
        ('Spec row, focus-visible: 2 px gold outline, inset', False, 'is-focus'),
        ('Spec row, active (pressed): darker ground, name white', False, 'is-active'),
        ('Spec row, the character\'s own spec', True, ''),
        ('Spec row, the character\'s own spec, hover', True, 'is-hover')))
    rows2 = ''.join(row(t, wrap(spec_row(s, False, False)), 420) for t, s in (('Tank word (druid)', bear), ('Healer word (druid)', resto)))
    P = lambda b, sel, you, st='', w=1440: f'<div style="width:{130 if w >= 480 else 66}px;padding-top:2px">{pill(b, sel, you, w, st)}</div>'
    pills = ''.join(row(t, h, 130 if 'below 480' not in t else 90) for t, h in (
        ('Band, rest', P(30, 20, None)), ('Band, hover', P(30, 20, None, 'is-hover')), ('Band, focus-visible', P(30, 20, None, 'is-focus')),
        ('Band, active', P(30, 20, None, 'is-active')), ('Band, selected', P(20, 20, None)),
        ('Band, selected, hover', P(20, 20, None, 'is-hover')), ('Band, selected, active', P(20, 20, None, 'is-active')),
        ('Band, the character\'s: a small "you" word inside, no second highlight', P(40, 20, 40)),
        ('Band, the character\'s and selected: one gold pill, "you" inside', P(40, 40, 40)),
        ('Band, below 480 px: the character\'s band is a small gold dot in the corner, label on one line', P(40, 20, 40, '', 390)),
        ('Band, below 480 px, selected and the character\'s', P(40, 40, 40, '', 390))))
    btn = lambda c: f'<a class="btn cta {c}" href="#x" style="min-width:250px">{COPY["lead_cta"]("Fury", 40)}</a>'
    ctas = ''.join(row(t, btn(c), 270) for t, c in (('Lead button, rest', ''), ('Lead button, hover', 'is-hover'),
                                                     ('Lead button, focus-visible', 'is-focus'), ('Lead button, active', 'is-active')))
    idl = row('Lead identity link (crest and name), hover: name underlines; same page as the button',
              f'<a class="idl is-hover" href="#x"><span class="display idn" style="font-size:22px;font-weight:700;color:{CLASS_COLOR["warrior"]}">{EXAMPLE["battletag"]}</span></a>', 300)
    leads = ''.join(row(t, lead(s, 1440)) for t, s in (
        ('Lead: addon synced (the reference, 12 minutes ago; trinkets have no pick at this band, so 15 slots)', 'synced'),
        ('Lead: no addon sync (character known, no gear on file). The count is replaced by one sentence and the link that fixes it; the button still works', 'nosync'),
        ('Lead: stale sync (19 days; the words past the age in ember, the label says last sync)', 'stale'),
        ('Lead: level 60 (band is plain 60, all 17 slots, no note)', 'l60'),
        ('Lead: a character under 20 (level 12). No count; the button opens the first band', 'under20'),
        ('Lead: loading (same 108 px box, so nothing moves when the character resolves; the grid and strip never wait for it)', 'loading')))
    err = row('Lead: error or no character. Omitted entirely, and the page is the signed-out page. Nothing says "failed"; the header selector carries the stale or signed-out mark.',
              f'<div class="panel" style="padding:14px 18px;font-size:13px;color:{BODY}">Signed-out page, from the hero down: Pick a class and spec.</div>')
    body = (f'<div style="{inner(1440, "36px 0")};display:flex;flex-direction:column;gap:22px">'
            f'<h1 class="display" style="font-size:22px;color:{TEXT}">BiS landing: states</h1>'
            f'<div style="display:flex;flex-wrap:wrap;gap:18px 24px;align-items:flex-start">{rows}{rows2}</div>'
            f'<div style="display:flex;flex-wrap:wrap;gap:18px 20px;align-items:flex-start">{pills}</div>'
            f'<div style="display:flex;flex-wrap:wrap;gap:18px 20px;align-items:flex-start">{ctas}{idl}</div>{leads}{err}</div>')
    return page(STYLE + body, 1440)


def before_board() -> str:
    """Today's /bis (web/src/pages/bis/index.astro + BisClassCard.astro), reproduced from its markup."""
    def fig(spec: str, role: str) -> str:
        b = band_entry(spec, 60)
        if role == 'tank':
            m = b['metrics']
            return f"Level 60: {m['effective_health']:,.0f} effective health · {m['dtps']:,.0f} damage taken per second"
        return f"Level 60: {b['set_dps']:.1f} {'HPS' if role == 'healer' else 'DPS'}"
    cards = ''
    for c in CLASSES:
        specs = [s for s in CATALOG if s['class_slug'] == c['slug']]
        if not specs:
            continue
        rows = ''.join(f'<li style="display:flex;align-items:center;justify-content:space-between;gap:12px"><a href="#x">{s["name"]}</a>'
                       f'<span class="mono" style="font-size:12px;color:{MUTED}">{fig(s["spec"], s["role"])}</span></li>' for s in specs)
        cards += (f'<section style="border:1px solid {BORDER};background:{RAISED};border-radius:4px;padding:16px"><div style="display:grid;grid-template-columns:56px 1fr;gap:16px;align-items:start">'
                  f'<div style="display:flex;flex-direction:column;align-items:center;gap:4px">{crest_img(c["slug"], 44)}<span style="font-size:11px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:{TEXT}">{c["name"]}</span></div>'
                  f'<ul style="margin:0;padding:0;list-style:none;display:flex;flex-direction:column;gap:4px;font-size:13px">{rows}</ul></div></section>')
    intro = 'A band is "the best you can wear at that level" from a named source -- a quest, a vendor, a dungeon, a drop you can farm. Pick a class to see its specs.'
    main = (f'<div style="max-width:900px;margin:0 auto;padding:40px 0;display:flex;flex-direction:column;gap:20px">'
            f'<div style="display:flex;flex-direction:column;gap:8px"><h1 style="font-size:28px;font-weight:700;letter-spacing:.02em;color:{TEXT}">BiS</h1><span style="font-size:12px;color:{MUTED}">Updated {UPDATED}</span></div>'
            f'<p style="margin:0;font-size:14px;color:{MUTED}">{intro}</p>'
            f'<p style="margin:0;font-size:14px"><span style="color:{MUTED}">Where each spec stands at level 60:</span> <a href="#tiers" style="font-weight:700">Tier list →</a></p>'
            f'<div style="display:flex;flex-direction:column;gap:12px">{cards}</div></div>')
    return page(STYLE + nav('BiS', True, EXAMPLE) + main + footer(False), 1440)


# ---------------------------------------------------------------------------- render
BOARDS = {   # name: (html, viewport width)
    'bis-landing': (board(1440, True), 1440),
    'bis-landing-signed-out': (board(1440, False), 1440),
    'bis-landing-2000': (board(2000, True), 2000),
    'bis-landing-1024': (board(1024, True), 1024),
    'bis-landing-768': (board(768, True), 768),
    'bis-landing-phone': (board(390, True), 390),
    'bis-landing-phone-signed-out': (board(390, False), 390),
    'bis-landing-360': (board(360, True), 360),
    'bis-landing-states': (states_board(), 1440),
    'bis-landing-before': (before_board(), 1440),
}


def main(names: list[str]) -> None:
    scratch = Path(tempfile.mkdtemp(prefix='gen_bis_landing_'))
    RENDERS.mkdir(exist_ok=True)
    for name in names:
        html, width = BOARDS[name]
        src = scratch / f'{name}.html'
        src.write_text(html)
        out = RENDERS / f'{name}.png'
        subprocess.run(['npx', 'playwright', 'screenshot', '--browser=chromium', f'--viewport-size={width},900', '--full-page',
                        '--wait-for-timeout=3000', src.as_uri(), str(out)], check=True, capture_output=True, cwd=WEB)
        print(out)


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

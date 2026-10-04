"""Mock boards for the /guides rebuild (design/specs/2026-10-04-guides.md).

  python3 design/mocks/gen_guides.py                       -> renders/guides-index.html (1440, signed in)
  GUIDES_PAGE=class python3 design/mocks/gen_guides.py      -> renders/guides-class.html (1440, warrior)
  GUIDES_PAGE=spec python3 design/mocks/gen_guides.py       -> renders/guides-spec.html (1440, fury)
  GUIDES_PAGE=spec GUIDES_PHONE=1 python3 ...               -> renders/guides-spec-phone.html (390, fury)

Numbers come from design/mocks/data/guides-warrior-fury.json and
data/builds/1.60.1.70009/bis/warrior-arms.json (the live BiS files for warrior-fury and
warrior-arms), design/mocks/data/guides-warrior.json (the warrior talent file), and
data/builds/1.60.1.70009/addon-data.json's rotations["warrior-fury"] for the rail's rotation
lines. Icons are read from data/builds/1.60.1.70009/icons/ (the full client icon tree), not
the smaller web/public mirror. No DPS figure anywhere in this file is a literal number: every
one is read off a band's own `set_dps`/`scale_factor` field at render time (round-1 review
fix, design/reviews/2026-10-04-guides-mock-wow-player.md finding 1).
"""
import json
import os

from mocklib import (
    ROOT,
    DATA,
    GOLD,
    MUTED,
    TEXT,
    BODY,
    BORDER,
    SOFT,
    CLASS_COLOR,
    crest,
    data_uri,
    nav,
    nav_phone,
    footer,
    page,
    write,
)

PAGE = os.environ.get('GUIDES_PAGE', 'index')  # index | class | spec
PHONE = os.environ.get('GUIDES_PHONE') == '1'
BUILD = '1.60.1.70009'
ICONS = ROOT / 'data/builds' / BUILD / 'icons'
BIS_DIR = ROOT / 'data/builds' / BUILD / 'bis'

FURY = json.load(open(DATA / 'guides-warrior-fury.json'))
ARMS = json.load(open(BIS_DIR / 'warrior-arms.json'))
TALENTS = json.load(open(DATA / 'guides-warrior.json'))
ADDON = json.load(open(ROOT / 'data/builds' / BUILD / 'addon-data.json'))

OBNOXIOUS = {'class': 'warrior', 'faction': 'alliance', 'battletag': 'Obnoxious Yell'}
WARRIOR = CLASS_COLOR['warrior']


def icon_uri(stem: str) -> str:
    return data_uri(ICONS / f'{stem}.webp')


def icon_img(stem: str, size: int, title: str = '') -> str:
    return (f'<img src="{icon_uri(stem)}" alt="" width="{size}" height="{size}" title="{title}" '
            f'style="width:{size}px;height:{size}px;border-radius:4px;border:1px solid {BORDER};'
            f'flex-shrink:0;object-fit:cover">')


def crest_img(slug: str, size: int) -> str:
    return f'<img class="crest" src="{crest(slug)}" alt="" style="width:{size}px;height:{size}px;--c:{CLASS_COLOR[slug]}">'


def btn(lbl: str, gold: bool = False, h: int = 40) -> str:
    cls = 'btn' + (' btn-gold' if gold else '')
    return f'<a class="{cls}" href="#{lbl.lower().replace(" ", "-")}" style="height:{h}px">{lbl}</a>'


def label(text: str) -> str:
    return f'<span class="label">{text}</span>'


def panel(body: str, extra: str = '') -> str:
    return f'<div class="panel" style="padding:16px 20px;display:flex;flex-direction:column;gap:10px;{extra}">{body}</div>'


# ---------------------------------------------------------------------------
# Spec identity: tab icons (specTabIcon.ts's own rule -- tree's own first talent,
# sorted (tier, column), this build has no dedicated tab icon), roles, descriptions and each
# guide's own `recommendedRaces` frontmatter (verbatim content, not a computed number).
SPEC_TAB_ICON = {'Arms': 'ability_rogue_ambush', 'Fury': 'spell_nature_purge', 'Protection': 'inv_shield_06'}
SPEC_HREF_SLUG = {'Arms': 'arms', 'Fury': 'fury', 'Protection': 'protection'}
SPEC_ROLE = {'Arms': 'DPS', 'Fury': 'DPS', 'Protection': 'Tank'}
SPEC_DESC = {
    'Arms': "Talents, rotation, stat priority, and race picks for Arms Warrior in Forever, with beta-versus-projection called out.",
    'Fury': "Talents, rotation, stat priority, and race picks for Fury Warrior in Forever, with beta-versus-projection called out.",
    'Protection': "Talents, tanking priority, stat priority, and race picks for Protection Warrior in Forever, with beta-versus-projection called out.",
}
# recommendedRaces frontmatter, first entry -- warrior/{arms,fury,protection}.md, verbatim.
SPEC_FIRST_RACE = {'Arms': 'human', 'Fury': 'human', 'Protection': 'dwarf'}
RACE_FACTION = {
    'human': 'alliance', 'dwarf': 'alliance', 'night-elf': 'alliance', 'gnome': 'alliance',
    'orc': 'horde', 'undead': 'horde', 'tauren': 'horde', 'troll': 'horde', 'skyborne': 'neutral',
}
SPEC_BIS = {'Arms': ARMS, 'Fury': FURY, 'Protection': None}  # Protection has no ranked BiS file


def set_dps_for(spec_name: str, band: int = 60) -> float | None:
    """Spec §4.D/§4.B's own ruling: the guide's own first `recommendedRaces` entry decides
    the faction, read off that spec's real BiS file -- never a literal DPS number in this
    module (round-1 review finding 1)."""
    bis = SPEC_BIS[spec_name]
    if bis is None:
        return None
    faction = RACE_FACTION[SPEC_FIRST_RACE[spec_name]]
    row = next((b for b in bis['bands'] if b['band'] == band and b['faction'] == faction), None)
    return row['set_dps'] if row else None


def spec_tabs(active: str, compact: bool = False) -> str:
    """`ClassHeader`'s own `SpecTabs` row. `compact` is the phone-only sizing rule from
    round-1 review finding 2 (spec §7.C): icon 20px (desktop 28px), tighter padding and an
    11px label so all three tabs -- including "Protection", the longest name -- fit inside
    390px without clipping the third tab at the viewport edge."""
    icon_px = 20 if compact else 28
    pad = '0 10px 0 6px' if compact else '0 14px 0 8px'
    gap = '6px' if compact else '8px'
    font_px = 11 if compact else 12
    out = []
    for name in ['Arms', 'Fury', 'Protection']:
        is_active = name == active
        st = (f'border-color:{WARRIOR};background:color-mix(in srgb, {WARRIOR} 8%, transparent)'
              if is_active else f'border-color:{BORDER}')
        name_color = f'color:{WARRIOR}' if is_active else f'color:{TEXT}'
        out.append(
            f'<a href="#{SPEC_HREF_SLUG[name]}" style="display:flex;align-items:center;gap:{gap};height:44px;'
            f'padding:{pad};border-radius:4px;border:1px solid;flex-shrink:0;{st}">'
            f'{icon_img(SPEC_TAB_ICON[name], icon_px)}'
            f'<span class="display" style="font-size:{font_px}px;font-weight:700;white-space:nowrap;{name_color}">{name}</span></a>'
        )
    row_gap = '6px' if compact else '8px'
    return f'<div style="display:flex;gap:{row_gap};overflow-x:auto">{"".join(out)}</div>'


def class_header(summary: str | None, active_spec: str, phone: bool = False) -> str:
    suffix = f'<p style="margin:0;font-size:14px;color:#c9c2b2;max-width:62ch">{summary}</p>' if summary else ''
    h1_text = f'{active_spec} Warrior' if summary else 'Warrior in Forever'
    # ClassHeader.astro's own `@media (max-width:1023px)` rule: padding drops to 22px 18px.
    pad = '22px 18px 0 18px' if phone else '28px 48px 0 48px'
    crest_px = 56 if phone else 64
    return f'''<div style="display:flex;flex-direction:column;gap:16px;padding:{pad};max-width:1344px">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guides</span>
<div style="display:flex;align-items:center;gap:18px">{crest_img("warrior", crest_px)}
<div style="display:flex;flex-direction:column;gap:6px;min-width:0">
<h1 class="display" style="margin:0;font-size:22px;line-height:1.1;font-weight:700;color:{WARRIOR}">{h1_text}</h1>
{suffix}</div></div>
{spec_tabs(active_spec, compact=phone)}
</div>'''


# ---------------------------------------------------------------------------
# INDEX
# ---------------------------------------------------------------------------
CLASS_SPECS = {
    'warrior': ['Arms Warrior in Forever', 'Fury Warrior in Forever', 'Protection Warrior in Forever'],
    'paladin': ['Holy Paladin in Forever', 'Protection Paladin in Forever', 'Retribution Paladin in Forever'],
    'hunter': ['Beast Mastery Hunter in Forever', 'Marksmanship Hunter in Forever', 'Survival Hunter in Forever'],
    'rogue': ['Assassination Rogue in Forever', 'Combat Rogue in Forever', 'Subtlety Rogue in Forever'],
    'priest': ['Discipline Priest in Forever', 'Holy Priest in Forever', 'Shadow Priest in Forever'],
    'shaman': ['Elemental Shaman in Forever', 'Enhancement Shaman in Forever', 'Restoration Shaman in Forever'],
    'mage': ['Arcane Mage in Forever', 'Fire Mage in Forever', 'Frost Mage in Forever'],
    'warlock': ['Affliction Warlock in Forever', 'Demonology Warlock in Forever', 'Destruction Warlock in Forever'],
    'druid': ['Balance Druid in Forever', 'Feral Druid in Forever', 'Restoration Druid in Forever'],
}


def build_index() -> str:
    cards = []
    for slug, specs in CLASS_SPECS.items():
        color = CLASS_COLOR[slug]
        name = slug.capitalize()
        links = ''.join(
            f'<li><a href="#{t.lower().replace(" ", "-")}" style="font-size:13px;color:{MUTED}">{t}</a></li>'
            for t in specs
        )
        cards.append(
            f'<div class="panel" style="display:flex;gap:16px;padding:16px 20px;align-items:flex-start">'
            f'{crest_img(slug, 56)}'
            f'<div style="display:flex;flex-direction:column;gap:6px;min-width:0">'
            f'<span class="display" style="font-size:16px;font-weight:700;color:{color}">{name}</span>'
            f'<ul style="list-style:none;margin:0;padding:0;display:flex;flex-direction:column;gap:3px">{links}</ul>'
            f'</div></div>'
        )
    grid = f'<div style="display:grid;grid-template-columns:repeat(3,1fr);gap:14px">{"".join(cards)}</div>'

    callout = (
        f'<div style="display:flex;align-items:center;gap:10px;padding:14px 18px;border:1px solid {GOLD};border-radius:6px;'
        f'background:color-mix(in srgb, {GOLD} 6%, transparent)">'
        f'{crest_img("warrior", 36)}'
        f'<a href="#fury-warrior" style="font-size:14px;font-weight:700;color:{GOLD}">Your guide: Fury Warrior &rarr;</a></div>'
    )

    header = f'''<div style="display:flex;flex-direction:column;gap:10px;padding:28px 48px 0 48px;max-width:1344px">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guides</span>
<h1 class="display" style="margin:0;font-size:22px;color:{TEXT}">Guides</h1>
<span style="font-size:12px;color:{MUTED}">Updated Oct 4</span>
<p style="margin:0;font-size:14px;color:{BODY};max-width:70ch">The beta caps at level 30. Level 60 play is a projection until we can confirm it against a live raid.</p>
{callout}
</div>'''

    body = f'''<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav("Guides", True, OBNOXIOUS)}{header}</div>
<div style="display:flex;flex-direction:column;gap:24px;padding:28px 48px 32px 48px">{grid}</div>{footer()}'''
    return write('guides-index.html', page(body, 1440)).as_posix()


# ---------------------------------------------------------------------------
# CLASS LANDING
# ---------------------------------------------------------------------------
WARRIOR_IDENTITY = (
    "Warrior is Forever's pure melee class: no ranged attack rotation and no spellcasting kit, "
    "just weapon damage built on Rage that generates from swinging and from taking hits. It "
    "covers both halves of the melee role inside one class &mdash; two damage specs and the "
    "game's shield-and-plate tank &mdash; rather than splitting tanking off into a separate class."
)


def spec_card(name: str) -> str:
    role = SPEC_ROLE[name]
    desc = SPEC_DESC[name]
    dps = set_dps_for(name)
    dps_line = (
        f'<span class="mono" style="font-size:13px;color:{TEXT}">{dps:.1f} DPS '
        f'<span style="color:{MUTED};font-weight:400">at 60, this set</span></span>'
        if dps is not None else ''
    )
    your_spec_style = (
        f'margin-left:8px;color:{GOLD};background:rgba(229,185,85,.14);'
        f'border-color:rgba(229,185,85,.35)'
    )
    your_spec = f'<span class="pill" style="{your_spec_style}">Your spec</span>' if name == 'Fury' else ''
    return f'''<div class="panel" style="display:flex;flex-direction:column;gap:10px;padding:18px 20px">
<div style="display:flex;align-items:center;gap:10px">{icon_img(SPEC_TAB_ICON[name], 40)}
<div style="display:flex;flex-direction:column"><span class="label" style="color:{MUTED}">{role}{your_spec}</span>
<span class="display" style="font-size:16px;font-weight:700;color:{WARRIOR}">{name} Warrior</span></div></div>
<p style="margin:0;font-size:13px;color:{BODY}">{desc}</p>
{dps_line}
{btn('Read the guide', gold=True)}
</div>'''


def build_class() -> str:
    header = f'''<div style="display:flex;flex-direction:column;gap:16px;padding:28px 48px 0 48px;max-width:1344px">
<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Guides</span>
<div style="display:flex;align-items:center;gap:18px">{crest_img("warrior", 64)}
<h1 class="display" style="margin:0;font-size:22px;color:{WARRIOR}">Warrior in Forever</h1></div>
{spec_tabs("")}
</div>'''

    body = f'''<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav("Guides", False)}{header}</div>
<div style="display:flex;flex-direction:column;gap:28px;padding:28px 48px 32px 48px;max-width:1344px">
<p style="margin:0;font-size:16px;line-height:1.6;color:{BODY};max-width:68ch">{WARRIOR_IDENTITY}</p>
<div style="display:grid;grid-template-columns:repeat(3,1fr);gap:16px">{spec_card('Arms')}{spec_card('Fury')}{spec_card('Protection')}</div>
</div>{footer()}'''
    return write('guides-class.html', page(body, 1440)).as_posix()


# ---------------------------------------------------------------------------
# SPEC GUIDE -- real talent decode, any band's talent string (the format is constant-width
# per tree regardless of band: 17 Arms / 18 Fury / 18 Protection digits, sorted tier/column)
# ---------------------------------------------------------------------------
def decode_tree(tree_name: str, code: str) -> list[dict]:
    tree = next(t for t in TALENTS['trees'] if t['name'] == tree_name)
    talents = sorted(tree['talents'], key=lambda t: (t['tier'], t['column']))
    return [{**t, 'rank': int(code[i])} for i, t in enumerate(talents)]


def decode_band(band_entry: dict) -> dict[str, list[dict]]:
    arms_c, fury_c, prot_c = band_entry['talents'].split('-')
    return {
        'Arms': decode_tree('Arms', arms_c),
        'Fury': decode_tree('Fury', fury_c),
        'Protection': decode_tree('Protection', prot_c),
    }


BAND60_ALLIANCE = next(b for b in FURY['bands'] if b['band'] == 60 and b['faction'] == 'alliance')
TREE_RANKS = decode_band(BAND60_ALLIANCE)


def tree_column(tree_name: str, cell_px: int) -> str:
    ranks = TREE_RANKS[tree_name]
    by_tier: dict[int, list[dict]] = {}
    for t in ranks:
        by_tier.setdefault(t['tier'], []).append(t)
    rows = []
    for tier in sorted(by_tier):
        cells = []
        for t in sorted(by_tier[tier], key=lambda x: x['column']):
            lit = t['rank'] > 0
            op = '1' if lit else '0.32'
            cells.append(
                f'<span style="position:relative;opacity:{op}">{icon_img(t["icon"], cell_px, t["name"])}'
                + (f'<span class="mono" style="position:absolute;bottom:-3px;right:-3px;font-size:8px;'
                   f'background:{SOFT};color:{GOLD};border-radius:2px;padding:0 2px;line-height:1.3">{t["rank"]}</span>'
                   if lit else '')
                + '</span>'
            )
        rows.append(f'<div style="display:flex;gap:3px">{"".join(cells)}</div>')
    pts = sum(t['rank'] for t in ranks)
    return (f'<div style="display:flex;flex-direction:column;gap:5px">'
            f'<span class="label" style="font-size:11px">{tree_name} &middot; {pts}</span>'
            f'<div style="display:flex;flex-direction:column;gap:3px">{"".join(rows)}</div></div>')


def full_tree() -> str:
    cols = ''.join(tree_column(n, 22) for n in ['Arms', 'Fury', 'Protection'])
    return (f'<div style="display:flex;gap:28px;padding:16px;background:{SOFT};border-radius:6px;overflow-x:auto">{cols}</div>'
            f'<div style="display:flex;gap:8px;margin-top:10px">{btn("Load this build")}{btn("Sim this build")}</div>')


def tree_thumbnail() -> str:
    by_tier: dict[int, list[dict]] = {}
    for t in TREE_RANKS['Fury']:
        by_tier.setdefault(t['tier'], []).append(t)
    rows = []
    for tier in sorted(by_tier):
        cells = []
        for t in sorted(by_tier[tier], key=lambda x: x['column']):
            lit = t['rank'] > 0
            op = '1' if lit else '0.32'
            cells.append(f'<span style="opacity:{op}">{icon_img(t["icon"], 18, t["name"])}</span>')
        rows.append(f'<div style="display:flex;gap:3px">{"".join(cells)}</div>')
    return (f'<div style="display:flex;flex-direction:column;gap:3px;padding:8px;background:{SOFT};'
            f'border-radius:4px;min-height:132px;justify-content:center">{"".join(rows)}</div>')


# ---------------------------------------------------------------------------
# SPEC GUIDE -- rail's rotation card (level-60 entry, first 4 lines)
# ---------------------------------------------------------------------------
ROTATION_60 = next(e for e in ADDON['rotations']['warrior-fury'] if e['level'] == 60)


def rotation_lines(n: int) -> str:
    out = []
    for line in ROTATION_60['lines'][:n]:
        out.append(
            f'<div style="display:flex;gap:10px;align-items:flex-start">{icon_img(line["icon"], 28)}'
            f'<span style="display:flex;flex-direction:column;gap:2px;min-width:0">'
            f'<span class="display" style="font-size:13px;font-weight:700;color:{TEXT}">{line["name"]}</span>'
            f'<span style="font-size:13px;color:{TEXT}">{line["condition"]}</span></span></div>'
        )
    return f'<div style="display:flex;flex-direction:column;gap:10px">{"".join(out)}</div>'


# ---------------------------------------------------------------------------
# SPEC GUIDE -- stat priority (prose-section pills, unchanged order list) and the rail's own
# weighted version (round-1 review finding 4: real DPS-scale numbers, not a 2/3-empty card)
# ---------------------------------------------------------------------------
STAT_PRIORITY = ['Attack power', 'Strength', 'Agility', 'Critical strike', 'Hit', 'Melee haste']
STAT_KEY = {
    'Attack power': 'attack_power', 'Strength': 'strength', 'Agility': 'agility',
    'Critical strike': 'crit', 'Hit': 'hit', 'Melee haste': 'melee_haste',
}
WEIGHTS_BY_STAT = {w['stat']: w for w in BAND60_ALLIANCE['weights']}


def stat_pills() -> str:
    out = []
    for i, s in enumerate(STAT_PRIORITY, 1):
        out.append(
            f'<li class="pill" style="border-color:{BORDER};background:{SOFT};color:{TEXT};height:26px;padding:0 10px">'
            f'<span class="mono" style="color:{GOLD};font-size:11px">{i}</span>'
            f'<span style="text-transform:none;font-weight:600;font-size:13px">{s}</span></li>'
        )
    return f'<ul style="list-style:none;margin:0;padding:0;display:flex;flex-wrap:wrap;gap:8px">{"".join(out)}</ul>'


def rail_stat_rows() -> str:
    """BiS's own `StatWeightsPanel` convention (scale factor normalized so the top
    *significant* stat = 1.00), trimmed to one line per stat for the rail card -- real numbers
    from this band's own `weights[]`, normalized here (not trusted pre-normalized from the
    file, since the top-scaling stat differs band to band)."""
    sig = [WEIGHTS_BY_STAT[STAT_KEY[s]] for s in STAT_PRIORITY if not WEIGHTS_BY_STAT[STAT_KEY[s]]['insignificant']]
    top = max(w['scale_factor'] for w in sig) if sig else 1.0
    rows = []
    for i, s in enumerate(STAT_PRIORITY, 1):
        w = WEIGHTS_BY_STAT[STAT_KEY[s]]
        if w['insignificant']:
            value = f'<span style="font-size:11px;color:{MUTED}">Not significant</span>'
        else:
            value = f'<span class="mono" style="font-size:12px;color:{TEXT}">{w["scale_factor"] / top:.2f}</span>'
        rows.append(
            f'<div style="display:flex;align-items:center;justify-content:space-between;gap:10px">'
            f'<span style="display:flex;align-items:center;gap:6px;font-size:13px;color:{TEXT}">'
            f'<span class="mono" style="color:{GOLD};font-size:11px">{i}</span>{s}</span>{value}</div>'
        )
    return f'<div style="display:flex;flex-direction:column;gap:8px">{"".join(rows)}</div>'


# ---------------------------------------------------------------------------
# SPEC GUIDE -- leveling band strip (real per-band talent data, alliance rows -- human is
# this guide's first recommended race). Round-1 review finding 5: each desktop row also gets
# a compact, non-interactive thumbnail of that band's own lit talents (all three trees,
# real decode) so a reader sees *which* points before clicking through; phone drops the
# thumbnail to avoid crowding the strip at 390px (spec §7.C says so explicitly).
# ---------------------------------------------------------------------------
def band_label(band: int) -> str:
    return '60' if band == 60 else f'{band} to {band + 9}'


def band_thumbnail(band_entry: dict) -> str:
    ranks = decode_band(band_entry)
    lit = [t for tree in ('Arms', 'Fury', 'Protection') for t in ranks[tree] if t['rank'] > 0]
    cells = []
    for t in lit:
        badge = (
            f'<span class="mono" style="position:absolute;bottom:-2px;right:-2px;font-size:7px;'
            f'background:{SOFT};color:{GOLD};border-radius:2px;padding:0 1px;line-height:1.2">{t["rank"]}</span>'
            if t['rank'] > 1 else ''
        )
        cells.append(f'<span style="position:relative">{icon_img(t["icon"], 16, t["name"])}{badge}</span>')
    return f'<div style="display:flex;flex-wrap:wrap;gap:3px;padding-top:2px">{"".join(cells)}</div>'


def leveling_strip() -> str:
    bands = [b for b in FURY['bands'] if b['faction'] == 'alliance']
    rows = []
    for i, b in enumerate(bands):
        border = '' if i == len(bands) - 1 else f'border-bottom:1px solid {SOFT};'
        thumb = '' if PHONE else band_thumbnail(b)
        rows.append(
            f'<div style="display:flex;flex-direction:column;gap:6px;padding:10px 0;{border}">'
            f'<div style="display:grid;grid-template-columns:90px 1fr auto;gap:16px;align-items:center">'
            f'<span class="mono" style="font-size:13px;color:{TEXT}">{band_label(b["band"])}</span>'
            f'<span style="font-size:13px;color:{MUTED}">{b["talent_points"]} points</span>'
            f'<a href="#load-in-planner" style="font-size:12px;font-weight:700;letter-spacing:.06em;'
            f'text-transform:uppercase">Load in planner</a></div>{thumb}</div>'
        )
    return f'<div class="panel" style="padding:4px 20px 6px 20px">{"".join(rows)}</div>'


# ---------------------------------------------------------------------------
# SPEC GUIDE -- races, sources (real data)
# ---------------------------------------------------------------------------
RACES = [
    ('Human', 'alliance', True), ('Dwarf', 'alliance', False), ('Night Elf', 'alliance', False), ('Gnome', 'alliance', False),
    ('Orc', 'horde', False), ('Undead', 'horde', False), ('Tauren', 'horde', False), ('Troll', 'horde', True),
    ('Skyborne', 'neutral', False),
]
FACTION_COLOR = {'alliance': '#6fb1ff', 'horde': '#ff6b5c', 'neutral': MUTED}


def race_pills() -> str:
    out = []
    for name, faction, rec in RACES:
        color = FACTION_COLOR[faction]
        st = (f'background:color-mix(in srgb, {color} 18%, transparent);border-color:{color};color:{color}'
              if rec else f'border-color:{BORDER};color:{MUTED}')
        out.append(
            f'<li class="pill" style="{st};height:28px;padding:0 10px;text-transform:none;font-weight:600;font-size:13px">{name}</li>'
        )
    return f'<ul style="list-style:none;margin:0;padding:0;display:flex;flex-wrap:wrap;gap:8px">{"".join(out)}</ul>'


SOURCES = [
    ('BLIZZARD', '#6fb1ff', 'rgba(0,112,221,.18)', 'rgba(0,112,221,.35)', 'Blizzard, Deep Dive panel recap'),
    ('COMMUNITY', '#7bff5c', 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)', 'Talents Forever (demo transcription)'),
    ('COMMUNITY', '#7bff5c', 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)', 'Warcraft Tavern: Classic Fury Warrior DPS rotation'),
    ('COMMUNITY', '#7bff5c', 'rgba(30,255,0,.10)', 'rgba(30,255,0,.25)', 'Icy Veins, the racial rework'),
    ('THIS SITE', GOLD, 'rgba(229,185,85,.14)', 'rgba(229,185,85,.35)', 'Forever Sixty stat weights, data/curated/specs.json'),
    ('THIS SITE', GOLD, 'rgba(229,185,85,.14)', 'rgba(229,185,85,.35)', 'Forever Sixty rotation data, data/curated/apl/warrior-fury.json'),
    ('THIS SITE', GOLD, 'rgba(229,185,85,.14)', 'rgba(229,185,85,.35)', 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'),
    ('THIS SITE', GOLD, 'rgba(229,185,85,.14)', 'rgba(229,185,85,.35)', 'This site, character planner'),
]


def sources_footer() -> str:
    rows = []
    for kind, fg, bg, bc, lbl in SOURCES:
        rows.append(
            f'<li style="display:flex;align-items:center;gap:10px;font-size:14px">'
            f'<span class="pill" style="color:{fg};background:{bg};border-color:{bc};height:22px;padding:0 8px">{kind}</span>'
            f'<a href="#source">{lbl}</a></li>'
        )
    return (f'<footer style="border-top:1px solid {SOFT};padding-top:20px;display:flex;flex-direction:column;gap:12px">'
            f'<span class="display" style="font-size:14px;font-weight:700;color:{TEXT}">Sources</span>'
            f'<ul style="list-style:none;margin:0;padding:0;display:flex;flex-direction:column;gap:8px">{"".join(rows)}</ul></footer>')


# ---------------------------------------------------------------------------
# SPEC GUIDE -- the 3-up action rail (new region, directly under the header)
# ---------------------------------------------------------------------------
def rail(cols: str) -> str:
    build_card = panel(
        f'{label("Build")}{tree_thumbnail()}<div style="display:flex;gap:8px">{btn("Load this build")}{btn("Sim this build")}</div>'
    )
    rotation_card = panel(
        f'{label("Rotation &middot; 60")}{rotation_lines(4)}'
        f'<a href="#rotation-and-priority" style="font-size:12px;font-weight:700;letter-spacing:.06em;'
        f'text-transform:uppercase;margin-top:auto">Full priority &darr;</a>'
    )
    stat_card = panel(f'{label("Stat priority")}{rail_stat_rows()}')
    return f'<div style="display:grid;grid-template-columns:{cols};gap:16px">{build_card}{rotation_card}{stat_card}</div>'


# ---------------------------------------------------------------------------
# SPEC GUIDE -- "On this page", the inline/boxed placement SpecTableOfContents already
# renders below 1024px (round-1 review finding 3: restore it on the phone board -- the live
# phone page already has it directly under the rail, closed by default).
# ---------------------------------------------------------------------------
def inline_toc_box() -> str:
    return (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:12px;'
            f'min-height:44px;padding:0 18px;border-bottom:1px solid {SOFT};background:var(--bg)">'
            f'<span class="label" style="color:{MUTED}">On this page</span>'
            f'<span style="font-size:12px;color:{MUTED};text-transform:uppercase;letter-spacing:.08em">Show</span></div>')


# ---------------------------------------------------------------------------
# SPEC GUIDE -- prose sections (real excerpts, content/guides/warrior/fury.md)
# ---------------------------------------------------------------------------
def h2(text: str) -> str:
    return f'<h2 class="display" style="font-size:18px;letter-spacing:.06em;text-transform:uppercase;margin:0;color:{TEXT}">{text}</h2>'


def p(text: str, width: str = '68ch') -> str:
    return f'<p style="margin:0;font-size:16px;line-height:1.6;color:{BODY};max-width:{width}">{text}</p>'


OVERVIEW = (
    "Fury is Forever's dual-wield DPS spec, built around Bloodthirst's frequent big hit and "
    "Whirlwind's multi-target filler rather than Arms's single hard-hitting finisher. The beta "
    "only reaches level 30, so nothing below about level 60 play, raiding, or best-in-slot gear "
    "has actually been tested &mdash; it's a projection from the demo trees and 1.12 knowledge, "
    "not confirmed play."
)
TALENTS_PROSE = (
    "This build reaching Bloodthirst spends 32 points in Fury, with the remaining 19 in Arms for "
    "Deep Wounds and Impale &mdash; Fury's high crit rate keeps both the bleed and the "
    "crit-damage talent relevant &mdash; rather than Protection, which has little to offer a "
    "dual-wielding damage build."
)
ROTATION_PROSE = (
    "Bloodthirst is the priority hit whenever its six-second cooldown is up, since it's the "
    "highest damage-per-Rage button in the kit; Whirlwind fills the gap whenever Bloodthirst "
    "isn't ready. Below 20% target health, Execute replaces the rest of the priority."
)
STAT_PROSE = (
    "Attack power is the reference stat: every weapon swing and special attack scales off it "
    "directly. Strength is the primary source of attack power, split between two weapons instead "
    "of one."
)
GEAR_PROSE = (
    "Look for attack power and strength first, then crit, then hit, following the order above. "
    "Specific pre-raid or raid-tier item picks can't be named with confidence yet: raid loot "
    "tables are incomplete until the first raid tier opens 9 December 2026."
)
ENCHANTS_PROSE = (
    "Enchant Weapon - Strength and Enchant Weapon - Crusader are both verified in this build's "
    "enchant data. For consumables, Mighty Rage Potion and Elixir of Ogre Strength are verified "
    "options in this build's consumable list."
)
RACES_PROSE = (
    "For Alliance, Human is a strong pick if you're dual-wielding swords. For Horde, Troll pairs "
    "well through Berserking, a direct swing-speed boost for a spec that already leans on attack "
    "speed for Flurry uptime."
)
PROFESSIONS_PROSE = (
    "Blacksmithing and Mining remain the standard Classic-era combination for a melee DPS "
    "Warrior. This is general Classic-community convention, not confirmed for Forever "
    "specifically."
)
LEVELING_PROSE = (
    "Fury is generally considered the strongest leveling spec for Warrior of the three, since "
    "dual-wielding two fast weapons together with Bloodthirst clears trash quickly."
)


def toc() -> str:
    items = ['Overview', 'Talents and builds', 'Rotation and priority', 'Stat priority', 'Gear',
             'Enchants and consumables', 'Races', 'Professions', 'Leveling']
    rows = ''.join(
        f'<li><a href="#{t.lower().replace(" ", "-")}" style="font-size:13px;color:{MUTED}">{t}</a></li>'
        for t in items
    )
    return (f'<nav style="display:flex;flex-direction:column;gap:8px;min-width:200px">'
            f'<span class="display" style="font-size:13px;font-weight:700;color:{TEXT};text-transform:uppercase;'
            f'letter-spacing:.08em">On this page</span>'
            f'<ol style="list-style:none;margin:0;padding:0;display:flex;flex-direction:column;gap:6px">{rows}</ol></nav>')


def prose_sections() -> str:
    return f'''<div style="display:flex;flex-direction:column;gap:28px">
{h2('Overview')}{p(OVERVIEW)}
{h2('Talents and builds')}{p(TALENTS_PROSE)}{full_tree()}
{h2('Rotation and priority')}{p(ROTATION_PROSE)}
{h2('Stat priority')}{stat_pills()}{p(STAT_PROSE)}
{h2('Gear')}{p(GEAR_PROSE)}
{h2('Enchants and consumables')}{p(ENCHANTS_PROSE)}
{h2('Races')}{race_pills()}{p(RACES_PROSE)}
{h2('Professions')}{p(PROFESSIONS_PROSE)}
{h2('Leveling')}{leveling_strip()}{p(LEVELING_PROSE)}
<p style="margin:0;font-size:13px;color:{MUTED}">Our inference from what has been said; may change</p>
{sources_footer()}
</div>'''


def build_spec() -> str:
    header = class_header(SPEC_DESC['Fury'], 'Fury', phone=PHONE)

    if PHONE:
        top_rail = f'<div style="padding:18px 18px 0 18px">{rail("1fr")}</div>'
        article = f'<div style="display:flex;flex-direction:column;gap:22px;padding:22px 18px 24px 18px">{prose_sections()}</div>'
        body = (f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav_phone(False)}{header}</div>'
                f'{top_rail}{inline_toc_box()}{article}{footer(phone=True)}')
        return write('guides-spec-phone.html', page(body, 390)).as_posix()

    top_rail = f'<div style="padding:24px 48px 0 48px;max-width:1344px">{rail("repeat(3,1fr)")}</div>'
    article = (f'<div style="display:grid;grid-template-columns:1fr 220px;gap:48px;align-items:start;'
               f'padding:32px 48px 40px 48px;max-width:1344px">'
               f'<div style="display:flex;flex-direction:column;gap:28px;max-width:820px">{prose_sections()}</div>'
               f'<div style="position:sticky;top:24px">{toc()}</div></div>')
    body = (f'<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav("Guides", False)}{header}</div>'
            f'{top_rail}{article}{footer()}')
    return write('guides-spec.html', page(body, 1440)).as_posix()


# ---------------------------------------------------------------------------
if PAGE == 'index':
    out = build_index()
elif PAGE == 'class':
    out = build_class()
else:
    out = build_spec()
print(out)

"""The character page, the reports list and the report screens of the flows-first app, app-composed in the flows
primitives (ws_app_flows.py, ws_app_secondary.py): what the app's character page and reports show today, and where the
site has more, one row at the end opens it. Data: the app's character fixtures (Elyra Duskvale, Holy Priest 60; Felsnap,
Beast Mastery Hunter 34; apps/android rankings/CharacterFixtures.kt) and the report contract fixtures
(docs/contracts/fixtures/app reports.*, coach.fight); item icons from the active build's items files.

The coordinator's additions over today's fixtures (2026-11-08): recent fights carry wipes, the Gear gap card shows, the
Forever Score has its sparkline and the Rating figure, and the reports list shows This week and Earlier.
"""
from __future__ import annotations

import json
from functools import cache

from mocklib import CLASS_COLOR, ROOT
from ws_app_frame import app_page
from ws_app_secondary import data, error_block, footer, screen, site_row, skeleton_block
from ws_prims import (Button, Disclosure, Figure, ItemRow, Meter, Notice, PageHeader, Row, Rows, Section, Skeleton, Spark, TabStrip, Tag, Tally,
                      crest_img, icon)

BUILD = json.loads((ROOT / 'web/src/data/active-build.json').read_text())['build']
PRIEST, HUNTER, WARRIOR = CLASS_COLOR['priest'], CLASS_COLOR['hunter'], CLASS_COLOR['warrior']
DOT = ' &middot; '
QUALITY = {'common': 1, 'uncommon': 2, 'rare': 3, 'epic': 4}
PARSE = [(95, '#ff8000'), (75, '#a335ee'), (50, '#0070dd'), (25, '#1eff00'), (0, '#9d9d9d')]

CHARACTER_CSS = '''
.gold-link{color:var(--gold);font-weight:600}
.answer-col{display:flex;flex-direction:column;gap:8px}
.answer-col .links{display:flex;gap:16px;flex-wrap:wrap}
.score-line{display:flex;align-items:center;gap:12px;margin-top:8px}.score-line .t-small{white-space:nowrap}
.blk.score-bars .tally{grid-template-columns:minmax(0,1fr)}
/* Best in slot in the small face, so a long worn name (Dormant Heart of the Mountain) sits well clear of its column's edge */
.irow .g .v .bis{font:600 13px/18px var(--font-body)}
.ck{display:inline-flex;align-items:center;justify-content:center;width:24px;height:24px;border-radius:50%;font-weight:700}
.ck.ok{color:var(--kill)}.ck.miss{color:var(--wipe)}
.dot{display:inline-block;width:10px;height:10px;border-radius:50%;background:var(--muted)}.dot.kill{background:var(--kill)}.dot.wipe{background:var(--wipe)}
.vis{display:inline-flex;color:var(--muted);margin-left:8px;vertical-align:middle}
.me-row{box-shadow:inset 3px 0 0 var(--gold)}
.coach .lbl{margin-top:16px}.coach .lbl:first-child{margin-top:4px}
.coach .lbl.muted{color:var(--muted)}
'''


def link(text: str) -> str:
    return f'<a class="link gold-link" href="#">{text} &rsaquo;</a>'


def page_html(body: str, title: str) -> str:
    return screen(body + f'<style>{CHARACTER_CSS}</style>', title)


@cache
def items(cls: str) -> dict[str, dict]:
    rows = json.loads((ROOT / f'data/builds/{BUILD}/items/{cls}.json').read_text())['items']
    return {r['name']: r for r in rows}


def item_icon(cls: str, name: str, quality: int) -> str:
    found = items(cls).get(name.replace('&rsquo;', "'"))
    return icon(found['icon'] if found else 'inv_misc_questionmark', 40, f'q{quality}')


def parse_colour(pct: int) -> str:
    return next(c for floor, c in PARSE if pct >= floor)


# ------------------------------------------------------------------------------------------------ the character page
ELYRA_GEAR = [  # slot, worn, quality, item level, gain or None (best in slot), pick, pick quality
    ('Head', 'Mooncloth Circlet', 'rare', 62, '+18.4', 'Devout Crown', 'rare'),
    ('Neck', 'Lady Maye&rsquo;s Pendant', 'epic', 64, '+2.9', 'Glowing Eye of Mordresh', 'rare'),
    ('Shoulder', 'Earthenweave Mantle', 'uncommon', 57, '+5.5', 'Mooncloth Shoulders', 'rare'),
    ('Back', 'Caretaker&rsquo;s Cape', 'rare', 63, '+3.9', 'Drape of Recovery', 'rare'),
    ('Chest', 'Field Marshal&rsquo;s Satin Tunic', 'epic', 74, '+0.8', 'Truefaith Vestments', 'epic'),
    ('Wrist', 'Devout Bracers', 'rare', 57, None, '', ''),
    ('Main hand', 'Resurgence Rod', 'rare', 53, '+21.1', 'Glowing Brightwood Staff', 'epic'),
    ('Ranged', 'Brilliant Wand', 'rare', 60, '+5.3', 'Torch of Light', 'epic'),
    ('Hands', 'Hands of the Exalted Herald', 'rare', 59, '+15.7', 'Devout Gloves', 'rare'),
    ('Waist', 'Penitent&rsquo;s Cinch', 'epic', 60, '+3.4', 'Wisdom of the Timbermaw', 'rare'),
    ('Legs', 'Earthenweave Leggings', 'uncommon', 58, '+25.7', 'Devout Skirt', 'rare'),
    ('Feet', 'Faith Healer&rsquo;s Boots', 'rare', 63, '+16.1', 'Devout Sandals', 'rare'),
    ('Finger', 'Rosewine Circle', 'rare', 60, '+0.5', 'Blessed Band of Light', 'rare'),
    ('Finger', 'Rosewine Circle', 'rare', 60, '+1.4', 'Band of Mending', 'rare'),
    ('Trinket', 'Weakness Analyzer', 'rare', 65, '+8.3', 'Serenity Field', 'rare'),
    ('Trinket', 'Dormant Heart of the Mountain', 'rare', 55, None, '', ''),
]
FELSNAP_GEAR = [
    ('Head', 'Brawler&rsquo;s Leather Hood', 'rare', 25, '+0.6', 'Brawler&rsquo;s Leather Helm', 'rare'),
    ('Neck', 'Archmagister&rsquo;s Faceted Pendant', 'rare', 35, '+1.1', 'Sentinel&rsquo;s Medallion', 'rare'),
    ('Shoulder', 'Mantle of Thieves', 'rare', 30, '+0.3', 'Forest Tracker Epaulets', 'rare'),
    ('Back', 'Hawkeye&rsquo;s Cloak', 'uncommon', 35, '+0.3', 'Tigerstrike Mantle', 'rare'),
    ('Chest', 'Tunic of Westfall', 'rare', 24, 'Same', 'Garb of Florid Feathers', 'rare'),
    ('Wrist', 'Madwolf Bracers', 'rare', 29, '+0.3', 'Hawkeye&rsquo;s Bracers', 'uncommon'),
    ('Main hand', 'Armor Piercer', 'rare', 29, '+1.3', 'Felbough Arm', 'rare'),
    ('Ranged', 'Ironweaver', 'rare', 34, '+0.7', 'Golemsight Long Gun', 'rare'),
    ('Hands', 'Pilferer&rsquo;s Gloves', 'rare', 28, '+0.8', 'Cartilage Shapers', 'rare'),
    ('Waist', 'Highlander&rsquo;s Chain Girdle', 'rare', 33, '+0.7', 'Stalker&rsquo;s Leather Belt', 'rare'),
    ('Legs', 'Petrolspill Leggings', 'rare', 30, '+0.2', 'Troll&rsquo;s Bane Leggings', 'rare'),
    ('Feet', 'Blackened Defias Boots', 'rare', 20, '+0.7', 'Insignia Boots', 'rare'),
    ('Finger', 'Monkey Ring', 'rare', 32, '+0.3', 'Ring of Precision', 'rare'),
    ('Finger', 'Monkey Ring', 'rare', 32, '+0.3', 'Protector&rsquo;s Band', 'rare'),
]
BEST = [('Lucifron', '3 Nov', 98, '1,012'), ('Magmadar', '3 Nov', 91, '948'), ('Gehennas', '3 Nov', 76, '902'),
        ('Garr', '27 Oct', 64, '866'), ('Baron Geddon', '27 Oct', 42, '790')]
RECENT = [('Lucifron', True, '2 d ago', '1,000'), ('Magmadar', False, '2 d ago', '612'), ('Magmadar', True, '2 d ago', '948'),
          ('Gehennas', True, '2 d ago', '902'), ('Garr', False, '9 d ago', '540')]
BUILDS = [('15/31/5', 'Holy &middot; first seen 3 Nov 2026'), ('21/30/0', 'Holy &middot; first seen 27 Oct 2026'),
          ('31/20/0', 'Discipline &middot; first seen 20 Oct 2026')]
ACTIVITY = [('Elyra Duskvale simmed 912 HPS, #3 of 14 healers in Cx66 Olympus (Holy, 60 raid)', '30 min ago', 'Your guild'),
            ('Elyra Duskvale uploaded Molten Core, 10 bosses in 92 min', '1 h ago', 'Logs'),
            ('<b>Holy, 60:</b> Elyra Duskvale climbed 6 places to <b>#4</b>', '21 h ago', 'Rankings'),
            ('Elyra Duskvale downed Lucifron in Molten Core, the first kill on Forever', '2 d ago', 'Logs')]
SCORE_PARTS = [('Gear', 410, 500), ('Experience', 520, 700), ('Parses', 470, 600), ('Attendance', 200, 200), ('Preparation', 100, 100)]
SCORE_HISTORY = [1610, 1640, 1655, 1700, 1690, 1735, 1760, 1790, 1810, 1830]


def answer(score: str, upgrades: tuple[str, str]) -> str:
    fig = Figure(score, 'Forever Score') if score else ''
    return (f'<div class="answer-col">{fig}{Figure(upgrades[0], "Upgrades available", upgrades[1], "gold", "big")}'
            f'<div class="links">{link("Open in simulator")}{link("Talents in planner")}</div></div>')


def elyra_head(owner: bool) -> str:
    meta = ['Level 60 Night Elf', 'Holy Priest', '<a class="link" href="#">15/31/5</a>', '<a class="link" href="#">&lt;Cx66 Olympus&gt;</a>',
            'Doomhowl', 'Hardcore US', Tag('Supporter')] + (['Addon &middot; 1 d ago'] if owner else [])
    return PageHeader('Elyra Duskvale', PRIEST, crest_img('priest', 56), meta, answer('1,830', ('+128.9', 'HPS')))


def felsnap_head(owner: bool) -> str:
    meta = ['Level 34 Dwarf', 'Beast Mastery Hunter'] + (['<a class="link" href="#">23/0/0</a>'] if owner else []) + ['Skyborne', 'Normal US'] \
        + (['Addon &middot; 1 d ago'] if owner else [])
    return PageHeader('Felsnap', HUNTER, crest_img('hunter', 56), meta, answer('', ('+7.7', 'DPS')) if owner else '')


def summary(next_line: str, action: str, standing: list[str]) -> str:
    rows = Rows([Row(f'{s} &rsaquo;', '', slim=True) for s in standing], open_=True)
    return f'<div class="blk" data-region="summary">{Notice(next_line, link(action))}{rows}</div>'


def gear_cell(cls: str, owner: bool, unit: str, cell: tuple) -> str:
    slot, name, q, ilvl, gain, pick, pq = cell
    if not owner:
        g = None
    elif gain is None:
        g = ('<span class="bis">Best in slot</span>', '', '')
    else:
        g = (gain, '' if gain == 'Same' else unit, f'vs <span style="color:var(--q{QUALITY[pq]},inherit)">{pick}</span>')
    return ItemRow(name, QUALITY[q], item_icon(cls, name, QUALITY[q]), f'{slot}{DOT}Item level {ilvl}', gain=g)


def gear(cls: str, cells: list[tuple], owner: bool, unit: str, band: str, gap: int | None, trinkets_empty: bool = False) -> str:
    rows = [gear_cell(cls, owner, unit, c) for c in cells]
    rows += [Row('Off hand', 'Two-hander in main hand', link=False, slim=True)] if cls == 'priest' else [Row('Off hand', '', link=False, slim=True)]
    if trinkets_empty:
        rows += [Row('Empty', 'Trinket', link=False, slim=True)] * 2
    gap_card = (f'<div class="blk" data-region="section-gear-gap">{Rows([Row("Gear gap", "", trail=f"<b>{gap} points</b>", link=False)], open_=True)}</div>'
                if gap is not None else '')
    return Section('Gear', '<div data-region="slots">' + ''.join(rows) + '</div>' + gap_card, action=link('Best in slot'), meta=band)


def score_section() -> str:
    figs = ('<div class="figs" style="--c:3">' + Figure('1,830', 'Score') + Figure('Epic', 'Tier', tone='epic')
            + Figure('65', 'Rating', ' of 100') + '</div>')
    line = f'<div class="score-line"><span class="t-small">As of 5 Nov</span>{Spark(SCORE_HISTORY, "#a335ee")}</div>'
    bars = Tally([('', label, have, want) for label, have, want in SCORE_PARTS])
    most = f'<p class="t-small" style="margin-top:12px">Most to gain: Experience, 180 points to the top. {link("How it is scored")}</p>'
    parts = Disclosure('Score parts', Rows([Row(n, f'{h} of {w}', link=False, slim=True) for n, h, w in SCORE_PARTS], open_=True))
    return Section('Forever Score', figs + line + f'<div class="blk score-bars">{bars}</div>' + most + parts, meta='6 rated fights')


def best_section(unit: str) -> str:
    rows = [Row(boss, f'{when}{DOT}Holy', trail=f'<b style="color:{parse_colour(p)}">{p}</b> <span class="t-small">{v} {unit}</span>', link=False)
            for boss, when, p, v in BEST]
    return Section('Best per encounter', Rows(rows, open_=True), action=link('All ranked fights (7)'))


def recent_section(unit: str) -> str:
    rows = [Row(boss, f'{"Kill" if kill else "Wipe"}{DOT}{age}', f'<span class="ck {"ok" if kill else "miss"}" role="img" aria-label="{"Kill" if kill else "Wipe"}">'
                f'{"&#10003;" if kill else "&times;"}</span>', f'<span class="t-small">{v} {unit}</span>', link=False) for boss, kill, age, v in RECENT]
    return Section('Recent fights', Rows(rows, open_=True))


def activity_section(owner: bool) -> str:
    shown = ACTIVITY if owner else ACTIVITY[1:]
    rows = [Row(text, f'{age}{DOT}<b>{tag}</b>', crest_img('priest', 40), link=True) for text, age, tag in shown]
    return Section('Activity', Rows(rows, open_=True))


def history_section() -> str:
    return Section('History', Rows([Row('Score over time, then parse per boss &rsaquo;', 'Every rated fight, kept with Premium'),
                                    Row('Coach history &rsaquo;', 'Every review the coach wrote for this character')], open_=True))


def elyra(owner: bool) -> str:
    head = elyra_head(owner)
    summ = summary('Next: get Devout Skirt from Stratholme: Baron Rivendare.', 'Open best in slot',
                   ['#3 of 14 healers in Cx66 Olympus, 60 raid', 'Holy is graded A this week', 'Onyxia&rsquo;s Lair in 33 days 8 hours, 18 in']) if owner else ''
    builds = Section('Builds seen', Rows([Row(s, sub, link=False, slim=True) for s, sub in BUILDS], open_=True))
    return (head + summ + gear('priest', ELYRA_GEAR, owner, 'HPS', 'Level 60 raid', 90 if owner else None) + score_section() + best_section('HPS')
            + recent_section('HPS') + builds + activity_section(owner) + (history_section() if owner else ''))


def felsnap(owner: bool) -> str:
    summ = summary('Next: spend your 2 talent points, then get Felbough Arm.', 'Open the planner', ['Beast Mastery is graded B this week']) if owner else ''
    score = Section('Forever Score', f'<p class="t-small" style="padding-top:8px">{"No logs yet. Upload a log to get a rating." if owner else "No logs yet."}</p>')
    return felsnap_head(owner) + summ + gear('hunter', FELSNAP_GEAR, owner, 'DPS', 'Levels 30 to 39', None, trinkets_empty=True) + score


def not_found() -> str:
    look = (Rows([Row('Name', 'Thalrin', link=False)], open_=True) + '<div class="vstack">' + Button('Look up', primary=True) + '</div>')
    add = Rows([Row('Sign in to see your characters', 'Battle.net lists every character on your account.'),
                Row('Get the addon and companion', 'Export with the addon and the companion to add this character.')], open_=True)
    return (PageHeader('Thalrin') + Notice('No character named Thalrin on US Hardcore.', tone='neutral')
            + Section('Look up a character', look) + Section('Add this character', add))


def loading() -> str:
    head = f'<div class="blk" data-region="head"><div class="score-line">{Skeleton(56, 56, True)}{Skeleton(280, 32)}</div></div>'
    cells = ''.join(f'<div style="margin-top:8px">{Skeleton(361, 60)}</div>' for _ in range(17))
    return head + Section('Gear', f'<div data-region="slots">{cells}</div>')


def character(state: str):
    def b(platform: str) -> str:
        bodies = {
            'owner-60': (lambda: elyra(True), 'Elyra Duskvale', 'Home', True),
            'visitor-60': (lambda: elyra(False), 'Elyra Duskvale', 'Search', True),
            'owner-34': (lambda: felsnap(True), 'Felsnap', 'Home', True),
            'no-score': (lambda: felsnap(False), 'Felsnap', 'Home', False),
            'not-found': (not_found, 'Thalrin', 'Search', False),
            'error': (lambda: PageHeader('Elyra Duskvale', PRIEST, crest_img('priest', 56)) + error_block('We could not load this character.'),
                      'Elyra Duskvale', 'Home', False),
            'loading': (loading, '', 'Home', False),
        }
        build, title, back, share = bodies[state]
        return app_page(platform, page_html(build(), title or 'Character'), 'home', back=back, title=title, share=share)
    return b


# ------------------------------------------------------------------------------------------------ the reports list
VIS = {  # the visibility glyph: muted, 16 px, its word the accessible name
    'public': '<circle cx="12" cy="12" r="9"/><path d="M3 12h18M12 3c3 3 3 15 0 18M12 3c-3 3-3 15 0 18"/>',
    'unlisted': '<path d="M10 14a4 4 0 0 0 6 0l3-3a4 4 0 0 0-6-6l-1 1"/><path d="M14 10a4 4 0 0 0-6 0l-3 3a4 4 0 0 0 6 6l1-1"/>',
    'guild': '<path d="M6 3v18M6 4h12l-3 4 3 4H6"/>',
    'private': '<rect x="5" y="11" width="14" height="10" rx="2"/><path d="M8 11V7a4 4 0 0 1 8 0v4"/>',
}
REPORTS = [  # section, title, sub, date, state tag (text, tone) or None, visibility
    ('This week', 'Onyxia&rsquo;s Lair', '1 boss &middot; 1 kill', '14 Dec', None, 'guild'),
    ('This week', 'Blackrock Spire', 'Updated 40 s ago', '15 Dec', ('Live', 'kill'), 'private'),
    ('Earlier', 'Molten Core', '10 bosses &middot; 10 kills', '12 Dec', ('Processing', 'gold'), 'public'),
    ('Earlier', 'Stratholme', 'Could not be read', '29 Nov', ('Failed', 'wipe'), 'unlisted'),
]


def vis_glyph(kind: str) -> str:
    return (f'<span class="vis" role="img" aria-label="{kind.capitalize()}"><svg width="16" height="16" viewBox="0 0 24 24" fill="none" '
            f'stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round">{VIS[kind]}</svg></span>')


def reports_head(count: int) -> str:
    meta = ['Every fight you logged, newest first.'] + ([f'{count} reports'] if count else [])
    return PageHeader('Your reports', meta=meta, actions=Button('Upload a log', primary=True))


def report_row(title: str, sub: str, date: str, state: tuple | None, vis: str) -> str:
    trail = f'<span class="t-small">{date}</span>' + (Tag(*state) if state else '') + vis_glyph(vis)
    return Row(title, sub, crest_img('warrior', 40), trail)


def reports(state: str):
    def b(platform: str) -> str:
        if state == 'empty':
            body = (reports_head(0) + Section('Reports', '<p class="t-small" style="padding-top:8px">No reports yet. Upload a log, or let the companion send your next raid.</p>'
                                              + '<div class="vstack">' + Button('Upload a log', primary=True) + '</div>'
                                              + f'<p style="margin-top:12px">{link("Open the sample report")}</p>'))
        elif state == 'failed':
            body = reports_head(0) + error_block('Your reports did not load.')
        else:
            body = reports_head(len(REPORTS)) + ''.join(
                Section(sec, Rows([report_row(*r[1:]) for r in REPORTS if r[0] == sec], open_=True)) for sec in ('This week', 'Earlier'))
        return app_page(platform, page_html(body, 'Your reports'), 'home', back='Home', title='Your reports')
    return b


# ------------------------------------------------------------------------------------------------ the report
FIGHTS = [('wipe', 'Onyxia &middot; pull 1 of 3', 'Wipe &middot; 1:12 &middot; boss at 61%'),
          ('wipe', 'Onyxia &middot; pull 2 of 3', 'Wipe &middot; 4:52 &middot; boss at 22%'),
          ('kill', 'Onyxia &middot; pull 3 of 3', 'Kill &middot; 4:28 &middot; 40 players &middot; 3 deaths'),
          ('', 'Trash &middot; Onyxia&rsquo;s Lair', '6 pulls &middot; 7:40')]
METERS = [('Zulgrin', 'rogue', 788), ('Vexara', 'mage', 762), ('Grimtusk', 'warrior', 745), ('Shadowfen', 'warlock', 731),
          ('Kelthar', 'rogue', 727), ('Obnoxious Yell', 'warrior', 712), ('Mordak', 'mage', 699), ('Thrallin', 'hunter', 684),
          ('Ashvane', 'warlock', 672), ('Rokhan', 'warrior', 663), ('Sylvari', 'mage', 655)]
ME = 'Obnoxious Yell'
SITE_REPORT = ('Full report on the site', 'Meters, casts, deaths and the timeline for every fight')


def report_head() -> str:
    return (PageHeader('Onyxia&rsquo;s Lair', meta=['Monday 14 December 2026 &middot; logged by Obnoxious Yell', Tag('Guild', 'muted')])
            + f'<div data-region="controls">{TabStrip(["Summary", "More &#9662;"], "Summary")}</div>')


def parses() -> str:
    row = Row('Onyxia', 'Kill 4:28 &middot; Obnoxious Yell, Fury', crest_img('warrior', 40),
              f'<span class="t-small">712 DPS &middot; <b style="color:{parse_colour(91)}">91</b> percentile</span>{Tag("New best")}', link=False)
    return Section('Your parses', Rows([row], open_=True))


def fights() -> str:
    rows = [Row(t, s, f'<span class="dot {k}" aria-hidden="true"></span>') for k, t, s in FIGHTS]
    return Section('Fights', Rows(rows, open_=True), meta='3 boss pulls')


def coach() -> str:
    c = data('coach.fight')
    r = c['review']
    then = ''.join(f'<p class="t-h4" style="margin-top:8px"><b>{x["title"]}</b></p><p class="prose">{x["why"]}</p>' for x in r['thenFix'])
    well = '<ul class="blist">' + ''.join(f'<li>{x["text"]}</li>' for x in r['doingWell']) + '</ul>'
    cant = '<ul class="blist">' + ''.join(f'<li>{x}</li>' for x in r['couldntCheck']) + '</ul>'
    basis = '; '.join(x.rstrip('.') for x in c['basis_labels'])
    body = (f'<div class="coach"><p class="lbl">{c["persona"]["label"]} {Tag("Premium")}</p>'
            f'<p class="lbl">Fix first</p><p><b>{r["fixFirst"]["title"]}</b></p><p class="prose">{r["fixFirst"]["why"]}</p>'
            f'<p class="lbl">Practice</p><p class="prose">{r["fixFirst"]["drillText"]}</p>'
            f'<p class="lbl">Then</p>{then}<p class="lbl">Doing well</p>{well}<p class="lbl muted">Couldn&rsquo;t check</p>{cant}'
            f'<p class="t-small" style="margin-top:12px">Reviewed 14 Dec &middot; Based on: {basis}.</p></div>')
    return Section('Coach', body)


def meters_sheet() -> str:
    """The meters sheet over a fight: the kill's damage done, your row edged in gold, then Show all 40."""
    top = max(v for _, _, v in METERS)

    def meter_row(i: int, name: str, cls: str, v: int) -> str:
        row = Row(name, '', crest_img(cls, 32), f'<b>{v}</b> <span class="t-small">DPS</span>', CLASS_COLOR[cls], rank=str(i), link=False,
                  foot=Meter(v / top))
        return row.replace('class="row', 'class="me-row row', 1) if name == ME else row
    rows = [meter_row(i, *m) for i, m in enumerate(METERS, 1)]
    seg = '<div class="chiprow" data-region="controls">' + ''.join(Button(x, primary=(x == 'Damage'), small=True) for x in ('Damage', 'Healing', 'Damage taken')) + '</div>'
    return (PageHeader('Damage done', meta=['<span class="dot kill"></span> Onyxia &middot; pull 3 of 3', 'Kill &middot; 4:28 &middot; 40 players'])
            + seg + Section('Damage', Rows(rows, open_=True)) + '<div class="vstack">' + Button('Show all 40') + '</div>')


def report(state: str):
    def b(platform: str) -> str:
        if state == 'loading':
            body = (PageHeader('Onyxia&rsquo;s Lair') + skeleton_block([(361, 96)], 'section-your-parses') + skeleton_block([(361, 220)], 'section-fights'))
            return app_page(platform, page_html(body, 'Report'), 'home', back='Your reports')
        if state == 'player':
            return app_page(platform, page_html(meters_sheet(), 'Damage done'), 'home', back='Report', title='Damage done')
        tail = (coach() if state == 'coach' else '')
        body = (report_head() + parses() + fights() + tail
                + footer('Meters, casts and deaths for every fight are under More, on foreversixty.gg.') + site_row(*SITE_REPORT))
        return app_page(platform, page_html(body, 'Onyxia&rsquo;s Lair'), 'home', back='Your reports', title='Onyxia&rsquo;s Lair', share=True)
    return b


SCREENS = {
    'character-owner-34': character('owner-34'),
    'character-owner-60': character('owner-60'),
    'character-visitor-60': character('visitor-60'),
    'character-no-score': character('no-score'),
    'character-not-found': character('not-found'),
    'character-loading': character('loading'),
    'character-error': character('error'),
    'reports': reports('list'),
    'reports-empty': reports('empty'),
    'reports-failed': reports('failed'),
    'report': report('summary'),
    'report-player': report('player'),
    'report-coach': report('coach'),
    'report-loading': report('loading'),
}

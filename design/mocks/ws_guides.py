"""Web-layout mock: the Guides area (/guides, /guides/<class>, /guides/<class>/<spec>), composed ONLY from ws_prims primitives.

Spec: design/specs/2026-11-08-guides.md. Every word, figure, talent, spell, item and race is read from
design/mocks/data/guides-facts.json, which design/mocks/guides_facts.mjs writes by running the pages' own models
(lib/guides/*, the copy module, the content files) on build 1.60.1.70790. Nothing is typed in except the strings this
round proposes (PROPOSED below, listed in spec section 8.2) and the layout. The signed-in characters are the home
fixtures (me-34 Felsnap, me-60 Thandrel); the level 34 Survival hunter and the level 60 Fire mage are derived from
them as the committed fixtures web/tests/e2e/fixtures/guides (README.md there says how).

  node design/mocks/guides_facts.mjs && python3 design/mocks/ws_guides.py
  python3 design/mocks/ws_guides.py [mock ...]                 (every mock, or just those named)
  node design/mocks/ws_shoot.mjs <mock> 390,1440,1920,2000   (each name in `python3 design/mocks/ws_guides.py --list`)
  python3 design/mocks/ws_guides.py compare <before dir>      (before/after sheets; before: shoot_guides_before.mjs)
FS_GUIDES_FACTS and FS_MOCK_OUT point the generator elsewhere (a parity spec's test-time render).
"""
import json
import os
import sys
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_prims import (Button, ClassSpecPicker, Facts, ItemRow, LaneBlock, LevelSlider, Link, Menu, Page, PageHeader, Plate, PlateColumns,
                      Prose, ReadingPanel, ReadingSection, Row, RowGrid, Rows, RowsSlim, Section, Skeleton, StatRow, StatTable, Tag, TalentTrees, Toc,
                      build_page, control_bar, crest_img, icon, seg, sel, write, REGION, Meter, QCOL, ItemTip, PaperdollFrame, reading_wrap_script)

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_GUIDES_FACTS') or HERE / 'data' / 'guides-facts.json').read_text())
C = F['copy']
SEC = {s['id']: s['title'] for s in F['sections']}   # SPEC_SECTIONS (lib/guides/sections.ts)
TA = F['tier_answer']
FACTION = {'alliance': '#6fb1ff', 'horde': '#ff6b5c'}
QUAL = {2: 'q2', 3: 'q3', 4: 'q4', 5: 'q5'}

# The strings this round proposed (spec 6.2) now live in lib/guides/copy.ts; read from there.
PROPOSED = {k: C[k] for k in ('answerLabel', 'noCharacter', 'glance', 'builds', 'sources')}


def e(text: str) -> str:
    return escape(str(text), quote=False)


def rsaquo(text: str) -> str:
    return f'{e(text)} &rsaquo;'


# ------------------------------------------------------------------------------------------------ pieces
def trees_of(decoded: list[dict]) -> list[tuple]:
    return [(t['name'], t['points'], [(x['tier'], x['column'], icon(x['icon'], 28), x['rank'], x['max_rank']) for x in t['talents']])
            for t in decoded]


def points_line(decoded: list[dict]) -> str:
    return ' &middot; '.join(f"{e(t['name'])} {t['points']}" for t in decoded)


def answer(kind: str, me: dict | None = None, page: dict | None = None) -> str:
    """The signed-in answer (PageHeader.answer). kind: same | other | loading | unknown | error | nochar."""
    lab = f'<span class="sr-only">{PROPOSED["answerLabel"]}</span>'   # a screen reader's label (ui/CharacterAnswer)
    if kind == 'loading':
        return f'<div class="ans-box">{lab}<span class="ans-row">{Skeleton(40, 40, True)}<span class="ans-who">{Skeleton(150, 22)}{Skeleton(190, 16)}</span></span>{Skeleton(220, 20)}</div>'
    if kind == 'unknown':
        return f'<div class="ans-box">{lab}<span class="t-small">{e(F["fns"]["unknown"])}</span>{Link(rsaquo(TA["unknownLink"]))}</div>'
    if kind == 'error':
        return f'<div class="ans-box">{lab}<span class="t-small" role="alert">{e(TA["error"])}</span>{Link(e(TA["retry"]))}</div>'
    if kind == 'nochar':
        return f'<div class="ans-box">{lab}<span class="t-small">{PROPOSED["noCharacter"]}</span>{Link(rsaquo(F["setup_title"]))}</div>'
    cls = me['class'].lower()
    who = (f'<span class="ans-row">{crest_img(cls, 40)}<span class="ans-who"><b class="ans-name" style="color:{CLASS_COLOR[cls]}">{e(me["name"])}</b>'
           f'<span class="t-small">{e(me["detail"])}</span></span></span>')
    if kind == 'same':
        first = me['talents_at'] if me['level'] < 60 else TA['bis']
        second = TA['bis'] if me['level'] < 60 else TA['sim']
        links = Link(rsaquo(first)) + Link(rsaquo(second))
        if me['level'] < 60 and me.get('next'):   # Q5: the ladder's next point, beside the level link
            links = f'<span class="t-small ans-next">{e(me["next"])}</span>' + links
    else:
        key = {'Beast Mastery': 'bm', 'Survival': 'survival', 'Fury': 'fury'}[me['spec']]
        links = Link(e(F['fns']['signedInCallout'][key]))
    return f'<div class="ans-box">{lab}{who}<span class="ans-links">{links}</span></div>'


def spell_row(item: dict, n: str = '') -> str:
    rank = f' <span class="t-small">{e(item["rankText"])}</span>' if item.get('rankText') else ''
    tag = ' ' + Tag(e(C['raidBuildTag']), 'muted') if item.get('raidBuildOnly') else ''
    # web GuideSpellRow: a line that names no spell has no lead at all, never an empty socket
    return Row(f'{e(item["name"])}{rank}{tag}', e(item.get('clause', '')), lead=icon(item['icon'], 28) if item.get('icon') else '', rank=n, link=False)


def stat_rows(rows: list[dict], numbered: bool, per_point: bool) -> list[str]:
    out, n = [], 0
    for r in rows:
        valued = r['reading'] == 'value' and 'perPercent' not in r
        if valued:
            n += 1
        lead = f'<span class="rk num muted">{n}</span>' if numbered and valued else ''
        figs = [e(r['valueText'])]
        if per_point:
            pp = r.get('perPoint')
            figs.append(e(r['perPointText']) or '&ndash;')
        out.append(StatRow(lead, e(r['label']), '', figs, meter=Meter(r['value']) if valued else ''))
    return out


# ------------------------------------------------------------------------------------------------ the spec guide
def band(g: dict) -> str:
    """web GuideGlance: the build alone, its trees with Load and Sim (the sections own Rotation and Stat priority)."""
    b = g['band']
    if not b['trees']:
        return ''
    col = (e(C['raidBuild'] if b['isRaidBuild'] else C['buildRailLabel']), '', TalentTrees(trees_of(b['trees'])),
           Button(e(C['loadThisBuild']), primary=True) + Button(e(C['simThisBuild'])))
    return f'<div class="g-glance">{PlateColumns([col], PROPOSED["glance"], frame="framed")}</div>'

def overview(g: dict) -> str:
    o = g['overview']
    prose = (f'<div class="sw"><div><h3 class="sw-h">{e(C["strengths"])}</h3>{Prose(o["strengths"], "plus")}</div>'
             f'<div><h3 class="sw-h">{e(C["weaknesses"])}</h3>{Prose(o["weaknesses"], "minus")}</div></div>')
    return ReadingSection(SEC['overview'], prose, LaneBlock(e(C['inARaid']), Prose(o['raid'], 'lane')), 'overview')


def talents(g: dict) -> str:
    t = g['talents']
    label = {'both': C['levelingAndRaidBuild'], 'leveling': C['levelingBuild'], 'raid': C['raidBuild']}
    rows = [Row(e(label[b['kind']]), points_line(b['trees']), trail=Link(e(C['loadThisBuild'])) + Link(e(C['simThisBuild'])), link=False) for b in t['blocks']]
    diff = f'<p class="t-small">{e(t["raidDiffLine"])}</p>' if t.get('raidDiffLine') else ''
    keys = []
    for k in t['keyTalents']:
        names = ' &middot; '.join(f'{e(x["name"])} <span class="t-small">{" &middot; ".join(e(r) for r in x["rankText"])}</span>' for x in k['talents'])
        keys.append(Row(names, e(k['why']), lead=icon(k['talents'][0].get('icon') or '', 32), link=False))
    # web GuideBuildsLane: the Builds block only when the raid build differs (one build is the band's own)
    builds = LaneBlock(PROPOSED['builds'], RowsSlim(rows) + diff) if len(t['blocks']) > 1 else ''
    lane = builds + LaneBlock(e(C['keyTalents']), wrap_list(Rows(keys, True)))
    return ReadingSection(SEC['talents-and-builds'], Prose(t['prose']), lane, 'talents-and-builds')


def rotation(g: dict, lists: dict | None = None) -> str:
    r = g['rotation']
    lists = r['lists'] if lists is None else lists
    lane = ''
    if lists:
        # web GuideRotationLane: the row's own Raid build only tag, no legend; a whole list wraps under the prose
        notes = (f'<p class="t-small">{e(C["rotationForRaidBuild"])}</p>' if r['forRaidBuild'] else '')
        blocks = [(C['opener'], lists['opener'], True), (C['priority'], lists['priority'], True), (C['cooldowns'], lists['cooldowns'], False)]
        lists_html = ''.join(LaneBlock(e(t), RowsSlim([spell_row(x, str(i + 1) if num else '') for i, x in enumerate(items)])) for t, items, num in blocks if items)
        lane = notes + f'<div class="gblocks" data-wrap-list>{lists_html}</div>'
    return ReadingSection(SEC['rotation-and-priority'], Prose(r['prose']), lane, 'rotation-and-priority')


def stats(g: dict, rows: list[dict] | None = None) -> str:
    s = g['stats']
    t = s['table']
    rows = t['rows'] if rows is None else rows
    rows = [r for r in rows if r['reading'] != 'not-on-items']   # web GuideStatLane: only what gear offers
    if rows:
        cap = ''
        c = t.get('hitCap')
        if c:
            cap += f'<p class="t-small"><b class="strong">{e(c["cap"])}</b>' + (f' &middot; {e(c["detail"])}' if c.get('detail') else '') + '</p>'
        if t.get('expertiseCap'):
            cap += f'<p class="t-small">{e(t["expertiseCap"]["text"])}</p>'
        table = StatTable([('', False), (e(F['markup']['statHead']), False), ('', False), (e(t['weightHead']), True), (e(t['perPointHead']), True)],
                          stat_rows(rows, False, True), 2, open_=True)
        # the table is one block: it never wraps under the prose (web GuideStatLane, no StatTable `wrap`)
        notes = ''   # web GuideStatLane: no footnotes under the table
        notes = f'<div class="gnotes" data-wrap-list>{notes}</div>' if notes else ''
        lane = f'<div class="gstat"><div class="lane-blk">{cap + table + notes}</div></div>'   # no title: the section heading says it
    else:   # no weights for the spec: the guide's own order, unnumbered weights absent (the shipped fallback)
        lane = f'<div class="lane-blk">{wrap_list(RowsSlim([Row(e(x), rank=str(i + 1), link=False) for i, x in enumerate(g["statPriority"])]))}</div>'
    return ReadingSection(SEC['stat-priority'], Prose(s['prose']), lane, 'stat-priority')


def gear(g: dict, band_no: int = 60, faction: str = 'alliance', data: bool = True) -> str:
    b = g['bis']
    lane = ''
    if data and b['data']:
        panel = next(p for p in b['data']['panels'] if p['band'] == band_no and p['faction'] == faction)
        items, weapons = [], []
        for r in panel['rows']:
            out = weapons if r['slot'] in ('main_hand', 'off_hand', 'ranged') else items
            if r.get('empty'):
                # an empty slot names no item: its line is not a tooltip trigger (the built ItemRow's empty state)
                out.append(ItemRow(f'<span class="muted">{e(r.get("emptyCopy") or "")}</span>', 1, '<span class="ic" style="width:36px;height:36px"></span>', '', e(r['slotLabel']))
                           .replace(' role="button" tabindex="0"', '', 1))
                continue
            out.append(ItemRow(e(r['itemName']), r['quality'], icon(r['icon'], 36, QUAL.get(r['quality'], '')), e(r['sourceDetail']), e(r['slotLabel'])))
        caption = panel['caption']
        note = ''   # at 60 the doll's title bar says pre-raid and whose picks they are
        lane = LaneBlock(e(F['fns']['levelRangeLabel'][str(band_no)]), f'<div class="items" data-wrap-list><div class="gblk">{"".join(items)}</div>' + (f'<div class="gblk">{"".join(weapons)}</div>' if weapons else '') + '</div>' + note, '' if band_no == 60 else e(caption))
    lane += f'<p>{Link(rsaquo(C["fullListLink"]))}</p>'
    # Every band's own BiS page and the role's tier list (SEO lane 2026-11-08), one quiet block under the full list.
    bands = ''.join(Link(e(s)) for s in F['fns']['bandSlugs'])
    lane += f'<p class="t-small gbands">{e(C["bisByLevelLabel"])}{bands}</p><p class="gtier">{Link(rsaquo(C["tierListLink"]))}</p>'
    # one block, as the built GuideBisLane: its rows and links 28 px apart, and the links may wrap under the prose
    lane = f'<div class="gbis" data-wrap-list>{lane}</div>'
    return ReadingSection(SEC['best-in-slot'], Prose(b['prose']) + guide_doll(g, faction) if data and b['data'] else Prose(b['prose']), lane, 'best-in-slot',
                          prose_first=True, fold=e(g['bisFold']))


def wrap_list(html: str) -> str:
    """Marks the list whose rows may wrap under the prose (the built lanes' `data-wrap-list`)."""
    i = html.index('>')
    return html[:i] + ' data-wrap-list' + html[i:]


QUALITY_NAME = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}


def guide_doll(g: dict, faction: str) -> str:
    """The Best in slot doll (web guides/GuideBisDoll.svelte): the level 60 picks on the shared PaperdollFrame in the prose
    column, the first pick's client tooltip in its centre; under the frame's 423 px it steps aside (the lane lists the picks)."""
    panel = next(p for p in g['bis']['data']['panels'] if p['band'] == 60 and p['faction'] == faction)
    picks = [r for r in panel['rows'] if not r.get('empty')]
    slots = [{'slot': r['slot'], 'label': f"{r['slotLabel']}: {r['itemName']}", 'covered': False, 'selected': i == 0,
              'icon': r['icon'], 'quality': r['quality']} for i, r in enumerate(picks)]
    # A two-hander shows greyed in the empty off hand (the built doll's coveredBy): the off hand's own line names it.
    main = next((r for r in picks if r['slot'] == 'main_hand'), None)
    off = next((r for r in panel['rows'] if r['slot'] == 'off_hand' and r.get('empty')), None)
    if main and off and (off.get('emptyCopy') or '').startswith(f"{main['itemName']} is a two-hander"):
        slots.append({'slot': 'off_hand', 'label': f"{off['slotLabel']}: {main['itemName']}, two-handed", 'covered': True, 'selected': False,
                      'icon': main['icon'], 'quality': main['quality']})
    first = picks[0]
    centre = ItemTip(first['itemName'], QUALITY_NAME.get(first['quality'], 'common'), [], [], first['sourceDetail'])
    meta = panel['raceName']   # web GuideBisDoll: the race; the list says level 60
    doll = PaperdollFrame(slots, g['title'], meta, None, centre, f"{g['title']} best in slot, level 60", heading='h2')
    return f'<div class="gdoll"><div class="gdoll-f">{doll}</div></div>'


def extras(g: dict, data: bool = True) -> str:
    x = g['extras']
    lane = ''
    if data and x['data']:
        ench = RowsSlim([Row(e(r['name']), e(r['slotLabel']), lead=icon(r['icon'], 28) if r.get('icon') else '', link=False)
                         for r in x['data']['enchants']])   # MUST 6: the same icon rows as Consumables
        # web GuideExtrasLane: the flasks and elixirs under Consumables, each header's rows as a block of its own
        groups: list[tuple[str, list[str]]] = []
        for line in x['data']['lines']:
            if 'header' in line:
                groups.append((line['header'], []))
                continue
            if not groups:
                groups.append((C['consumables'], []))
            r = line['row']
            # web lib/guides/gear-extras.ts consumableNote: only the slot a stone or oil goes on
            note = r['gives'].split(', ', 1)[1] if ', ' in r['gives'] else ''
            groups[-1][1].append(Row(e(r['name']), e(note[:1].upper() + note[1:]), lead=icon(r['icon'], 28) if r.get('icon') else '', link=False))
        blocks = LaneBlock(e(C['enchants']), ench) + ''.join(LaneBlock(e(label), RowsSlim(rows)) for label, rows in groups)
        lane = f'<div class="gblocks" data-wrap-list>{blocks}</div>'   # a whole block wraps under the prose, with its label
    return ReadingSection(SEC['enchants-and-consumables'], Prose(x['prose']), lane, 'enchants-and-consumables', fold=e(x['fold']))


def races(g: dict) -> str:
    r = g['races']
    body = ''
    for grp in r['groups']:   # Alliance then Horde, Recommended first (glance.ts racesByFaction)
        rows = []
        for x in grp['rows']:
            foot = ''.join(f'<span class="racial">{icon(z["icon"], 20)}<span><b>{e(z["name"])}</b></span></span>' for z in x['racials'])   # name only: the prose says the numbers
            portrait = r.get('portraits', {}).get(x.get('slug', ''))
            rows.append(Row(e(x['name']), title_color=FACTION[x['faction']], lead=icon(portrait, 28) if portrait else '',
                            link=False, foot=foot))
        head = C['alliancePicksHeading'] if grp['faction'] == 'alliance' else C['hordePicksHeading']
        body += f'<div class="grace"><h4 class="sub-h">{e(head)}</h4>' + Rows(rows, True) + '</div>'   # a faction wraps whole
    lines = ''.join(f'<p class="t-small">{e(t)}</p>' for t in r['lines'])
    # web GuideRacesLane: no lane title (the section is Races); the factions head their groups
    return ReadingSection(SEC['races'], Prose(r['prose']), f'<div class="graces" data-wrap-list>{body + lines}</div>', 'races',
                          fold=e(F['fns']['foldShow'].replace('%s', f"{len(r['rows'])} races")))


def leveling(g: dict, level: int = 60, data: bool = True) -> str:
    lv = g['leveling']
    prose = Prose(lv['early'] + (f'<h3>{e(lv["stripHeading"])}</h3>' if lv['stripHeading'] else '') + lv['late'])
    lane = ''
    at = lv['at'].get(str(level))
    if data and at:
        lane = LaneBlock(e(C['talentsByLevel']),
                         f'<p class="lvl-l"><b class="strong">{e(at["point"])}</b><span class="t-small">{e(at["spent"])}</span></p>'
                         + LevelSlider(level, 10, 60, e(C['levelSliderLabel'])) + TalentTrees(trees_of(at['trees'])) + f'<p>{Link(rsaquo(C["loadInPlanner"]))}</p>')
    # web: the tree sits above its prose in one column (no rail beside nine paragraphs)
    return ReadingSection(SEC['leveling'], lane + prose, '', 'leveling', fold=e(F['fns']['foldShow'].replace('%s', 'leveling')))


def sources(g: dict) -> str:
    seen, rows = set(), []
    for s in g['sources']:
        if s['url'] in seen:
            continue
        seen.add(s['url'])
        rows.append(Row(f'<span class="gold link">{e(source_title(s["label"]))}</span>', lead=Tag(e(F['source_labels'][s['kind']]), 'muted')))
    return Section(PROPOSED['sources'], RowsSlim(rows))


def reading_page(top: str, toc: str, body: str) -> str:
    """ui/Page layout 'reading' with tocFrom 1024: the full-width top (the lede and the plate), the contents rail at the left
    from 1024, the body beside it."""
    return (f'<div class="pc g-read"><div class="layout-reading toc-lg"><div class="g-top">{top}</div>'
            f'<nav class="toc" aria-label="{e(C["tocBarLabel"])}">{toc}</nav><div class="body">{body}</div></div></div>')


def spec_bar(g: dict, open_m: bool = False) -> str:
    items = [(e(t['name']), t['active'], None, icon(t['icon'], 22)) for t in g['tabs']]
    toc_m = '<span class="t-label" style="display:block;margin:8px 0 4px">' + e(C['tocBarLabel']) + '</span>' + RowsSlim(
        [Row(e(s['title'])) for s in F['sections']])
    end = f'<span class="toc-sel">{sel(e(C["tocBarLabel"]))}</span>'
    spec = e(C['specGroupLabel'])
    return control_bar([(spec, seg(items, label=spec), seg(items, label=spec) + toc_m)],
                       f'{e(g["specName"])} {e(g["className"])}', end=end, open_m=open_m, label=e(C['barLabel']), labels_from=1920)


def spec_page(gid: str, kind: str = 'out', me_key: str = '', level: int = 60, open_m: bool = False, partial: bool = False) -> str:
    g = F['guides'][gid]
    me = F['me'].get(me_key) if me_key else None
    ans = '' if kind == 'out' else answer(kind, me, g)
    head = PageHeader(e(g['title']), CLASS_COLOR[g['classSlug']], crest_img(g['classSlug'], 56), [e(g['meta'])], ans)
    gp = dict(g)
    if partial:   # FORCED (spec 5.3): no rotation entry, no weights, no ranked BiS file, no extras, no ladder (none occurs at 70790)
        gp = {**g, 'band': {**g['band'], 'rotation': [], 'stats': []}}
    band_no = 60 if level >= 60 else max(20, level // 10 * 10)
    faction = me['faction'] if me and kind == 'same' else 'alliance'
    sections = [overview(gp), talents(gp), rotation(gp, {} if partial else None), stats(gp, [] if partial else None),
                gear(gp, band_no, faction, not partial), extras(gp, not partial), races(gp),
                ReadingSection(SEC['professions'], Prose(g['professions']), '', 'professions', fold=e(F['fns']['foldShow'].replace('%s', 'professions'))), leveling(gp, 60, not partial)]   # the server render: the tree's island (client:visible) moves to the character's level only once scrolled to
    body = (ReadingPanel(''.join(sections)) + sources(g)
            )
    toc = Toc(e(C['tocBarLabel']), [(e(s['title']), s['id']) for s in F['sections']], 'overview')
    html = f'<div class="g-page">{head + spec_bar(g, open_m) + reading_page(band(gp), toc, body)}</div>' + PAGE_CSS
    who = (me['name'], me['class'].lower(), f"Level {me['level']}") if me else None
    note = (f'Guide mock (design/specs/2026-11-08-guides.md), built only from ws_prims, read from guides-facts.json (build {F["build"]}). '
            + ('Partial-data state: rotation, weights, BiS and extras removed (FORCED; no spec lacks them at this build). ' if partial else '')
            + (f'Signed in: {me["name"]} ({me["identity"]}, {me["level_line"]}), a home fixture' + (' (derived, see guides_facts.mjs).' if me_key in ('survival34', 'mage60') else '.') if me else ''))
    return build_page(f'{g["title"]} guide', html + reading_wrap_script(), 'Guides', who, note)


def source_title(label: str) -> str:
    """web lib/sources.ts sourceTitle: the publisher the kind pill already names is not said again."""
    for prefix in ('Blizzard, ', 'Forever Sixty, '):
        if label.startswith(prefix):
            rest = label[len(prefix):]
            return rest[:1].upper() + rest[1:]
    return label


# ------------------------------------------------------------------------------------------------ the class landing
FACTION_COLOR = {'alliance': '#6fb1ff', 'horde': '#ff6b5c'}  # web tokens.css --color-alliance, --color-horde


def race_list(groups: list[dict]) -> str:
    """web guides/GuideRaceList: per faction its label, then each race (portrait 24, name in the faction colour, "new")."""
    out = []
    for g in groups:
        items = ''.join(
            f'<li style="color:{FACTION_COLOR[g["faction"]]}">'
            + (f'<span aria-hidden="true">{icon(r["icon"], 24)}</span>' if r['icon'] else '')
            + f'<span>{e(r["name"])}</span>'
            + ('<span class="t-small" style="color:var(--muted);font-weight:400">new</span>' if r['fresh'] else '')
            + '</li>' for r in g['races'])
        out.append(f'<div class="grl-g"><h4 class="t-label grl-h">{e(g["heading"])}</h4><ul class="grl-items">{items}</ul></div>')
    return f'<div class="grl">{"".join(out)}</div>'


def class_page(cls: str, kind: str = 'out', me_key: str = '') -> str:
    L = F['landings'][cls]
    me = F['me'].get(me_key) if me_key else None
    mine = me and me['class'].lower() == cls
    cards = []
    for c in L['cards']:
        you = bool(mine and me['spec'] == c['name'])
        title = e(c['name'])   # the accent rule marks yours, no pill
        # the whole row is the link to the guide (no "Read the guide" line)
        cards.append(Plate(Row(title, e(c['line']), lead=icon(c['icon'], 40), title_color=CLASS_COLOR[cls], selected=you)))
    specs = (f'<div class="g-specs" role="region" aria-label="{e(C["classLandingSpecsHeading"])}" style="--specs:{len(cards)}">'
             f'{"".join(cards)}</div>')
    main = ''.join(Section(e(s['heading']), Prose(s['html']), boxed=True) for s in L['sections'] if not s['panel'])
    rail = ''.join(Section(e(s['heading']), (race_list(L['races']) if s['heading'] == 'Races' else '')
                           + (Prose(s['html'], 'lane') if s['html'] else ''), boxed=True) for s in L['sections'] if s['panel'])
    srcs = Section(PROPOSED['sources'], RowsSlim([Row(f'<span class="gold link">{e(source_title(s["label"]))}</span>', lead=Tag(e(F['source_labels'][s['kind']]), 'muted')) for s in L['sources']]))
    ans = '' if kind == 'out' else answer(kind, me)
    head = PageHeader(e(L['title']), CLASS_COLOR[cls], crest_img(cls, 56), [f'Updated {e(L["updated"])}'], ans)
    html = head + Page('ledger', f'<div class="g-ledger">{specs}</div>') + Page('split', main, rail) + Page('ledger', srcs) + PAGE_CSS
    who = (me['name'], me['class'].lower(), f"Level {me['level']}") if me else None
    return build_page(f'{L["title"]} guides', html, 'Guides', who, f'Class landing mock (design/specs/2026-11-08-guides.md); data guides-facts.json, build {F["build"]}.')


# ------------------------------------------------------------------------------------------------ /guides
def index_page(kind: str = 'out', me_key: str = '') -> str:
    me = F['me'].get(me_key) if me_key else None
    yours = f"{me['class'].lower()}/{me['spec'].lower().replace(' ', '-')}" if me else ''
    picks = [(c['slug'], e(c['name']), c['color'], [(f"{c['slug']}/{s['slug']}", e(s['name']), icon(s['icon'], 28)) for s in c['specs']])
             for c in F['index']['classes']]
    picker = ClassSpecPicker(picks, F['markup']['indexTitle'], class_links=True, yours=yours, yours_label=e(C['yourSpecPill']))
    ans = '' if kind == 'out' else answer(kind, me)
    head = PageHeader(e(F['markup']['indexTitle']), '', '', [f'Updated {e(F["index"]["updated"])}'], ans)
    who = (me['name'], me['class'].lower(), f"Level {me['level']}") if me else None
    return build_page('Guides', head + Page('ledger', picker) + PAGE_CSS, 'Guides', who, f'Guides index mock (design/specs/2026-11-08-guides.md); data guides-facts.json, build {F["build"]}.')


# ------------------------------------------------------------------------------------------------ every element state
def states_page() -> str:
    g = F['guides']['hunter/survival']
    me = F['me']
    blocks = []
    def block(title, html):
        blocks.append(f'<section class="st-b"><h2 class="t-h3">{title}</h2>{html}</section>')
    for t, k, m in (('Answer: loading (reserved at its final size; the session is hinted)', 'loading', None),
                    ('Answer: this spec, level 34', 'same', me['survival34']), ('Answer: this spec, level 60', 'same', me['survival60']),
                    ('Answer: another spec or class', 'other', me['warrior60']), ('Answer: the spec is not known yet', 'unknown', None),
                    ('Answer: no character on the account', 'nochar', None), ('Answer: the account read failed', 'error', None)):
        block(t, PageHeader(e(g['title']), CLASS_COLOR['hunter'], crest_img('hunter', 56), [e(g['meta'])], answer(k, m, g)))
    rot = g['rotation']['lists']['priority']   # the Rotation section's rows (the band shows the build alone)
    block('Rotation row: default, hover, a raid-build-only row', RowsSlim([spell_row(rot[0], '1'),
          spell_row(rot[1], '2').replace('class="row', 'class="row hov', 1), spell_row({**rot[2], 'raidBuildOnly': True}, '3')]))
    block('Spec row on /guides: default, hover, your spec, focus-visible', RowsSlim([
        Row('Beast Mastery', lead=icon(F['index']['classes'][2]['specs'][0]['icon'], 28)),
        Row('Marksmanship', lead=icon(F['index']['classes'][2]['specs'][1]['icon'], 28)).replace('class="row', 'class="row hov', 1),
        Row('Survival', lead=icon(F['index']['classes'][2]['specs'][2]['icon'], 28), selected=True),
        Row('Survival', lead=icon(F['index']['classes'][2]['specs'][2]['icon'], 28)).replace('class="row', 'class="row fv', 1)]))
    block('Contents: the bar&rsquo;s On this page menu open (below 1760); the rail with Rotation current (from 1760)',
          '<div class="st-pair"><div class="st-menu">' + Menu(e(C['tocBarLabel']), [('', [(e(s['title']), '', '', '', s['id'] == 'rotation-and-priority') for s in F['sections']])])
          + '</div><nav class="toc st-toc">' + Toc(e(C['tocBarLabel']), [(e(s['title']), s['id']) for s in F['sections']], 'rotation-and-priority') + '</nav></div>')
    block('Talents by level: at 10 (previous off), 34, 60 (next off)', ''.join(f'<div class="st-lv">{LevelSlider(v, 10, 60, "Level")}</div>' for v in (10, 34, 60)))
    block('Talents by level: loading (the island hydrates when visible; the slider row and the trees are reserved)',
          f'<div class="st-lv">{Skeleton(360, 44)}</div><div class="st-lv">{Skeleton(390, 190)}</div>')
    block('Best in slot lane: an empty slot and an item whose icon is missing', '<div class="items">'
          + ItemRow('<span class="muted">Nothing better yet</span>', 1, '<span class="ic" style="width:36px;height:36px"></span>', '', 'Trinket')
          + ItemRow('Mark of Fordring', 3, '<span class="ic q3" style="width:36px;height:36px"></span>', 'Quest: In Dreams', 'Neck') + '</div>')
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Guides: element states</h1></div>' + Page('ledger', f'<div class="statesheet">{"".join(blocks)}</div>') + PAGE_CSS
    return build_page('Guide states', body, 'Guides', ('Felsnap', 'hunter', 'Level 34'), 'Guides states sheet; data as the guide mocks.')


PAGE_CSS = '''<style>
.gnotes{display:flex;flex-direction:column;gap:8px}.gnotes p{margin:0}.graces,.gblocks{display:flex;flex-direction:column;gap:28px;min-width:0}.grace{display:flex;flex-direction:column;gap:8px}.grace>.sub-h{padding:0}.graces>p{margin:0}
.gstat .st .stc{grid-template-columns:0 var(--st-name) var(--st-bar) repeat(var(--n),var(--st-fig))}   /* GuideStatLane: no icon track */
@media(max-width:1023px){.gstat .st .stc{grid-template-columns:0 minmax(0,1fr) auto}}
@media(min-width:1280px){#leveling .rd.solo{grid-template-columns:minmax(0,1fr)}#leveling .rd.solo>.rd-prose{columns:2 320px;column-gap:var(--gap)}#leveling .lane-blk{break-inside:avoid;margin-bottom:24px}#leveling li,#leveling p{break-inside:avoid}#leveling h3{break-after:avoid}}   /* [spec].astro: Leveling in two columns */.gblk{display:flex;flex-direction:column}.gbis{display:flex;flex-direction:column;gap:28px}.gbis>p:not(.gbands,.gtier){margin:0}.rprose:empty+.gdoll .gdoll-f{margin-top:0}
.gdoll{container-type:inline-size}.gdoll-f{max-width:640px;margin-top:24px}@container (max-width:422.98px){.gdoll-f{display:none}}
/* page-level layout only (no surfaces): the answer box's inner rows, column splits, spacing */
.ans-row{display:flex;align-items:center;gap:12px}.ans-who{display:flex;flex-direction:column;min-width:0}.ans-name{font:700 18px/22px var(--font-body)}
.ans-links{display:flex;gap:4px 16px;flex-wrap:wrap}
.grl{display:flex;flex-direction:column;gap:10px;margin-bottom:12px}.grl-g{display:flex;flex-direction:column;gap:6px}.grl-h{margin:0;color:var(--muted);font-weight:700}
.grl-items{display:grid;grid-template-columns:repeat(auto-fill,minmax(150px,1fr));gap:6px 12px;margin:0;padding:0;list-style:none}
.grl-items li{display:flex;gap:8px;align-items:center;min-height:24px;font-size:14px;font-weight:700;line-height:18px}
.grl-items li>span[aria-hidden]{display:flex;flex:none}.grl-items li .ic{display:block}
.gbands{display:flex;flex-wrap:wrap;gap:4px 12px;margin-top:-16px}.gtier{margin-top:-16px}
@media(min-width:900px){.ans-row{flex-direction:row-reverse}.ans-who{align-items:flex-end}.ans-links{justify-content:flex-end}}
.g-read{padding-top:28px}.layout-reading{row-gap:var(--section-gap)}.g-top{grid-column:1/-1;display:flex;flex-direction:column;gap:22px;min-width:0}

.g-ledger{display:flex;flex-direction:column;gap:22px}
.g-specs{display:grid;grid-template-columns:minmax(0,1fr);gap:18px}
@media(min-width:1024px){.g-specs{grid-template-columns:repeat(var(--specs),minmax(0,1fr))}}
.g-page{padding-bottom:40px}  /* [spec].astro .guide-page's own bottom space, under the page body's gap (ws_lib main) */
.sw{display:grid;gap:20px var(--gap);grid-template-columns:minmax(0,1fr)}@media(min-width:768px){.sw{grid-template-columns:repeat(2,minmax(0,1fr))}}
.rd-lane .rows:not(.slim) .row:not(.slim){min-height:58px}  /* web tokens.css --row-h: a name-only key talent row is the token height */
.sw-h{margin:0;padding-bottom:8px;font:600 15px/20px var(--font-body);color:#f3eee2}
.lvl-l{display:flex;flex-direction:column}
.racial{display:flex;gap:8px;align-items:flex-start;font-size:13px;line-height:18px;color:var(--muted);width:100%}.racial b{color:var(--text);font-weight:600}
.sub-h{padding:12px 0 2px;margin:0;font:600 15px/20px var(--font-body);color:#f3eee2;letter-spacing:0;text-transform:none}
.g-glance .pcol-h h2{font:600 15px/20px var(--font-body);color:#f3eee2;letter-spacing:0;text-transform:none}
.rd-lane .rows.slim .row .ic:not(img){border:1px solid #82705b;background:var(--bg)}
.lane-blk .st{--st-name:minmax(max-content,2fr)}   /* GuideStatLane: a stat's name on one line */
.items{display:flex;flex-direction:column}
.foot-note{padding-top:16px}
section>.sec-h+.rows,section>.sec-h+.rprose{margin-top:12px}.gl-fill{padding-top:8px}.ans-next{align-self:center}
@media(min-width:900px){.ans-next{flex-basis:100%;text-align:right}}.ic-spacer{display:inline-block;flex:none}
.cgrp{padding-bottom:20px;min-width:0}
@media(max-width:767px){.cbar .pc{padding-block:0;min-height:0}}   /* the phone bar is its 48 px line, as the built ControlBar */
.toc-sel{display:inline-flex}@media(min-width:1024px){.toc-sel{display:none}}
.statesheet{display:grid;gap:40px}.st-b>h2{margin-bottom:12px}.st-pair{display:flex;gap:48px;flex-wrap:wrap;align-items:flex-start}.st-menu{position:relative;width:360px;height:440px}.st-menu .menu{position:static}
.st-toc{display:block;width:200px}.st-lv{max-width:480px;padding-bottom:12px}
.row.fv{outline:2px solid var(--gold);outline-offset:-2px}
</style>'''


PAGES = {
    'web-guides': lambda: index_page(),
    'web-guides-you': lambda: index_page('other', 'hunter34'),
    'web-guides-hunter': lambda: class_page('hunter'),
    'web-guides-hunter-you': lambda: class_page('hunter', 'other', 'hunter34'),
    'web-guides-mage': lambda: class_page('mage'),
    'web-guides-priest': lambda: class_page('priest'),
    'web-guides-priest-other': lambda: class_page('priest', 'other', 'warrior60'),
    'web-guides-hunter-survival': lambda: spec_page('hunter/survival'),
    'web-guides-hunter-survival-you': lambda: spec_page('hunter/survival', 'same', 'survival34', 34),
    'web-guides-hunter-survival-other': lambda: spec_page('hunter/survival', 'other', 'warrior60'),
    'web-guides-hunter-survival-loading': lambda: spec_page('hunter/survival', 'loading'),
    'web-guides-hunter-survival-error': lambda: spec_page('hunter/survival', 'error'),
    'web-guides-mage-fire': lambda: spec_page('mage/fire'),
    'web-guides-mage-fire-you': lambda: spec_page('mage/fire', 'same', 'mage60', 60),
    'web-guides-priest-holy': lambda: spec_page('priest/holy'),
    'web-guides-priest-holy-partial': lambda: spec_page('priest/holy', partial=True),
    'web-guides-bar': lambda: spec_page('hunter/survival', open_m=True),
    'web-guides-states': states_page,
}


def compare(before: Path) -> None:
    """Before/after at the same scale: live captures (shoot_guides_before.mjs) beside the mock renders."""
    import ws_compose
    ws_compose.BEFORE = before
    pairs = [('guides', 'guides', 'web-guides'), ('guides-hunter', 'guides-hunter', 'web-guides-hunter'),
             ('guide-survival', 'guide-survival', 'web-guides-hunter-survival')]
    for name, live, mock in pairs:
        for w, sc in ((1440, .5), (1920, .42), (2000, .4), (390, .8)):
            print(ws_compose.compose(name, live, mock, w, sc, max_h=6000))


if __name__ == '__main__':
    if sys.argv[1:2] == ['--list']:
        print('\n'.join(PAGES))
    elif sys.argv[1:2] == ['compare']:
        compare(Path(sys.argv[2]))
    else:
        only = sys.argv[1:]
        unknown = [name for name in only if name not in PAGES]
        if unknown:
            sys.exit(f'ws_guides.py: no such mock: {" ".join(unknown)} (see --list)')
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))

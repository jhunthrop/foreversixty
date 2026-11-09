"""Mock batch 2: guides (index, class landing, spec guide) in the directions system (plan 2026-11-09-redesign).

Data, never invented: design/mocks/data/guides-facts.json, which design/mocks/guides_facts.mjs writes by running the
pages' own models (lib/guides/*, the copy module, the content files) on the active build: every word, talent, spell,
item, weight and race. The spec shown is Survival Hunter (melee in Forever). Drawn by prims.py: panels, the talent grid,
reading sections, the table of contents, item rows, row lists and tables.
"""
import json
import re
from html import escape
from pathlib import Path

import data as D
from common import write
from prims import PRIMS_CSS, facts, item_row, meter, panel, reading_section, row_list, table, talent_trees, toc
from style import footer, header, medallion, page, titlebar

F = json.loads((Path(__file__).resolve().parent.parent / 'data/guides-facts.json').read_text())
C = F['copy']
GID = 'hunter/survival'
VIEWER = ('Thandrel', 'warrior', 60)
QUALITY_NAME = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}
GTIP = re.compile(r'<span class="gtip"[^>]*>(.*?)</span>')  # the live page's tooltip spans: here, a plain gold term

GUIDE_CSS = PRIMS_CSS + '''
.gd{padding-top:6px}.gl{display:grid;grid-template-columns:minmax(0,1.25fr) minmax(0,1fr) minmax(0,.9fr);gap:0;margin-top:22px}
.gl>div{padding:18px 20px}.gl>div+div{border-left:1px solid rgba(255,255,255,.1)}.gl h2{font-size:20px;margin-bottom:10px}
.gl .tt{grid-template-columns:repeat(3,auto);gap:10px}.gl .tg{grid-template-columns:repeat(4,22px);grid-auto-rows:22px;gap:5px 6px;padding:7px}
.gl .tc,.gl .tc img{width:22px;height:22px}.gl .tc b{display:none}.gl .tt h3{font-size:14px;margin-bottom:4px}
.acts{display:flex;gap:10px;flex-wrap:wrap;margin-top:14px}.gold{color:var(--gold)}
.page2{display:grid;grid-template-columns:200px minmax(0,1fr);gap:28px;margin-top:22px;align-items:start}.page2 .toc{position:sticky;top:16px}
.card.rd{padding:0}.sw{display:grid;grid-template-columns:1fr 1fr;gap:18px}.sw h3{font:600 14px var(--body);color:#fff;margin-bottom:4px}
.lede{font-size:16px;color:var(--text);max-width:80ch;padding-top:4px}.pn .prose{max-width:none}
.specs{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:18px;margin-top:22px}.specs .card{padding:18px 20px}
.specs h2{font-size:22px;display:flex;align-items:center;gap:10px}.specs p{margin:10px 0 12px;font-size:15px}
.idx{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:18px;margin-top:22px}.idx .card{padding:16px 18px}
.idx h2{font-size:21px;display:flex;align-items:center;gap:10px;margin-bottom:6px}
@media(max-width:1100px){.gl,.specs,.idx{grid-template-columns:1fr}.gl>div+div{border-left:0;border-top:1px solid rgba(255,255,255,.1)}
.page2{grid-template-columns:minmax(0,1fr)}.page2 .toc{display:none}}
@media(max-width:700px){.sw{grid-template-columns:1fr}
/* the glance keeps its three small trees side by side on a phone, as the guide shows the whole split */
.gl .tt{grid-template-columns:repeat(3,auto);gap:6px}.gl .tg{grid-template-columns:repeat(4,16px);grid-auto-rows:16px;gap:4px;padding:5px}
.gl .tc,.gl .tc img{width:16px;height:16px}.gl .tt h3{font-size:12px}}
'''


def prose(html: str) -> str:
    """The live page's tooltip terms: gold at a term's first mention in the block, plain after (the reviewers: every term in
    gold made the paragraphs heavy). The tooltip still opens on every mention in the build."""
    seen: set[str] = set()

    def term(m: re.Match) -> str:
        word = m.group(1)
        first = word not in seen
        seen.add(word)
        return f'<span class="gold">{word}</span>' if first else word

    return GTIP.sub(term, html)


def hit_cap(text: str) -> str:
    """'Hit cap: reached (special attacks), 28% (white swings), 16% (spells)' as a facts list: what, then the figure."""
    items = []
    for part in text.split(':', 1)[-1].split(','):
        m = re.match(r'\s*(.+?)\s*\((.+)\)\s*$', part)
        if m:
            items.append((m.group(2).capitalize(), escape(m.group(1))))
    return f'<h3 style="font:600 13px var(--body);color:var(--muted);margin:12px 0 4px">Hit cap</h3>{facts(items)}' if items else ''


def trees(decoded: list[dict]) -> str:
    return talent_trees([(t['name'], t['points'], [(x['tier'], x['column'], D.icon_uri(x['icon']), x['rank'], x['max_rank']) for x in t['talents']])
                         for t in decoded])


def spell_rows(items: list[dict], numbered: bool) -> str:
    rows = []
    for i, x in enumerate(items):
        ic = f'<img class="ai" src="{D.icon_uri(x["icon"])}" alt="" width="28" height="28">' if x.get('icon') else ''
        lead = f'<span class="num muted" style="width:16px;display:inline-block">{i + 1}</span>{ic}' if numbered else ic
        rows.append((f'<span style="display:flex;gap:8px;align-items:center">{lead}</span>', escape(x['name']), escape(x.get('clause', '')), ''))
    return row_list(rows)


def weights_table(rows: list[dict], per_point: bool, weight_head: str, per_point_head: str = '') -> str:
    valued = [r for r in rows if r['reading'] == 'value' and 'perPercent' not in r]
    top = max((r['value'] for r in valued), default=1) or 1
    out = []
    # with the per-point column the table sits in a side column: no bar, so its four columns fit
    bars = not per_point
    for r in rows:
        figs = [escape(r['valueText'])] + ([escape(r.get('perPointText') or '') or '&ndash;'] if per_point else [])
        bar = [meter(r['value'] / top) if r in valued else ''] if bars else []
        out.append([f'<span class="strong">{escape(r["label"])}</span>'] + bar + figs)
    heads = ([('Stat', False)] + ([('', False)] if bars else []) + [(escape(weight_head), True)]
             + ([(escape(per_point_head), True)] if per_point else []))
    return table(heads, out, hide_phone={1} if bars else set())


def glance(g: dict) -> str:
    b = g['band']
    build = (f'<div><h2>{escape(C["buildRailLabel"])}</h2>{trees(b["trees"])}<div class="acts"><a class="btn" href="#">{escape(C["loadThisBuild"])}</a>'
             f'<a class="btn2" href="#">{escape(C["simThisBuild"])}</a></div></div>')
    rot = (f'<div><h2>{escape(b["rotationTitle"])}</h2>{spell_rows(b["rotation"], False)}'
           + (f'<p class="muted" style="font-size:13px;margin-top:8px">{escape(b["filler"])}</p>' if b.get('filler') else '')
           + f'<p style="margin-top:10px"><a href="#">{escape(C["fullPriorityLink"])}</a></p></div>')
    cap = b.get('hitCap')
    stats = (f'<div><h2>{escape(C["statPriorityRailLabel"])}</h2>{weights_table(b["stats"], False, g["stats"]["table"]["weightHead"])}'
             + (hit_cap(cap["cap"]) if cap else "")
             + (f'<p class="muted" style="font-size:13px">{escape(b["unitLine"])}</p>' if b.get('unitLine') else '') + '</div>')
    return f'<section class="card framed gl">{build}{rot}{stats}</section>'


def sections(g: dict) -> str:
    sec = {s['id']: s['title'] for s in F['sections']}
    o, t, r, s, b, x, rc = g['overview'], g['talents'], g['rotation'], g['stats'], g['bis'], g['extras'], g['races']
    sw = (f'<div class="sw"><div><h3>{escape(C["strengths"])}</h3>{prose(o["strengths"])}</div>'
          f'<div><h3>{escape(C["weaknesses"])}</h3>{prose(o["weaknesses"])}</div></div>')
    out = [reading_section(sec['overview'], sw, [(C['inARaid'], f'<div class="prose" style="font-size:14.5px">{prose(o["raid"])}</div>')], 'overview')]
    label = {'both': C['levelingAndRaidBuild'], 'leveling': C['levelingBuild'], 'raid': C['raidBuild']}
    blocks = row_list([('', escape(label[k['kind']]), ' &middot; '.join(f'{escape(tr["name"])} {tr["points"]}' for tr in k['trees']),
                        f'<a href="#">{escape(C["loadThisBuild"])}</a>') for k in t['blocks']])
    keys = row_list([(f'<img class="ai" src="{D.icon_uri(k["talents"][0].get("icon") or "")}" alt="" width="28" height="28">',
                      ' &middot; '.join(escape(z['name']) for z in k['talents']), escape(k['why']), '') for k in t['keyTalents']])
    out.append(reading_section(sec['talents-and-builds'], prose(t['prose']), [('Builds', blocks), (C['keyTalents'], keys)], 'talents-and-builds'))
    lists = r['lists']
    rot_sides = [(C[k], spell_rows(lists[k], k != 'cooldowns')) for k in ('opener', 'priority', 'cooldowns') if lists.get(k)]
    out.append(reading_section(sec['rotation-and-priority'], prose(r['prose']), rot_sides, 'rotation-and-priority'))
    tb = s['table']
    out.append(reading_section(sec['stat-priority'], prose(s['prose']),
                               [(C['statPriorityRailLabel'], weights_table(tb['rows'], True, tb['weightHead'], tb['perPointHead'])
                                 + (f'<p class="muted" style="font-size:13px;margin-top:8px">{escape(tb["unitLine"])}</p>' if tb.get('unitLine') else ''))],
                               'stat-priority'))
    pnl = next(p for p in b['data']['panels'] if p['band'] == 60 and p['faction'] == 'alliance')
    items = ''.join(item_row(D.icon_uri(rw['icon']), rw['itemName'], QUALITY_NAME.get(rw['quality'], rw['quality']) if isinstance(rw['quality'], int) else rw['quality'],
                             escape(rw['sourceDetail']), escape(rw['slotLabel'])) for rw in pnl['rows'] if not rw.get('empty'))
    out.append(reading_section(sec['best-in-slot'], prose(b['prose']), [(F['fns']['levelRangeLabel']['60'], items + f'<p style="margin-top:10px"><a href="#">{escape(C["fullListLink"])} &rsaquo;</a></p>')],
                               'best-in-slot'))
    ench = row_list([(f'<img class="ai" src="{D.icon_uri(e.get("icon") or "")}" alt="" width="28" height="28">', escape(e['name']),
                      f'{escape(e["slotLabel"])} &middot; {escape(e["gives"])}', '') for e in x['data']['enchants']])
    cons = row_list([(f'<img class="ai" src="{D.icon_uri(ln["row"].get("icon") or "")}" alt="" width="28" height="28">', escape(ln['row']['name']),
                      escape(ln['row']['gives']), '') for ln in x['data']['lines'] if 'row' in ln])
    out.append(reading_section(sec['enchants-and-consumables'], prose(x['prose']), [(C['enchants'], ench), (C['consumables'], cons)], 'enchants-and-consumables'))
    races = []
    for grp in rc['groups']:
        rows = [('', f'<span style="color:{"#3d94f0" if r["faction"] == "alliance" else "#e0524a"}">{escape(r["name"])}</span>',
                 ' &middot; '.join(escape(z['name']) for z in r['racials']), escape(C['recommendedTag']) if r['recommended'] else '') for r in grp['rows']]
        races.append((C['alliancePicksHeading'] if grp['faction'] == 'alliance' else C['hordePicksHeading'], row_list(rows)))
    out.append(reading_section(sec['races'], prose(rc['prose']), races, 'races'))
    out.append(reading_section(sec['professions'], prose(g['professions']), [], 'professions'))
    lv = g['leveling']
    at = lv['at'].get('60')
    lv_side = [(C['talentsByLevel'], f'<p style="font-size:14px"><b class="strong">{escape(at["point"])}</b> <span class="muted">{escape(at["spent"])}</span></p>'
                + f'<p style="margin-top:10px"><a href="#">{escape(C["loadInPlanner"])} &rsaquo;</a></p>')] if at else []
    out.append(reading_section(sec['leveling'], prose(lv['early'] + lv['late']), lv_side, 'leveling'))
    return ''.join(out)


def spec_page() -> str:
    g = F['guides'][GID]
    tabs = [[(t['name'], t['active']) for t in g['tabs']]]
    band = titlebar(g['classSlug'], g['title'], escape(g['meta']), title_colour=D.CLASS_COLOR[g['classSlug']], ctl=tabs, top=header(*VIEWER, 'Guides'))
    lede = f'<p class="lede" style="margin-top:22px">{escape(g["summary"])}</p>'
    contents = toc([(s['title'], s['id']) for s in F['sections']], 'overview')
    srcs = row_list([('', f'<a href="#">{escape(s["label"])}</a>', escape(F['source_labels'][s['kind']]), '') for s in {s['url']: s for s in g['sources']}.values()])
    body = (band + f'<main class="w gd">{lede}{glance(g)}<div class="page2"><aside>{contents}</aside><div><section class="card rd">{sections(g)}</section>'
            f'{panel("Sources", srcs)}</div></div></main>' + footer())
    return page(f'Forever Sixty, {g["title"]} guide', GUIDE_CSS, body)


def class_page() -> str:
    cls = GID.split('/')[0]
    L = F['landings'][cls]
    cards = ''.join(f'<section class="card"><h2 style="color:{D.CLASS_COLOR[cls]}"><img class="ai" src="{D.icon_uri(c["icon"])}" alt="" width="36" height="36">{escape(c["name"])}</h2>'
                    f'<p>{escape(c["line"])}</p><a href="#">{escape(C["readTheGuide"])} &rsaquo;</a></section>' for c in L['cards'])
    main = ''.join(panel(s['heading'], f'<div class="prose">{prose(s["html"])}</div>') for s in L['sections'] if not s['panel'])
    rail = ''.join(panel(s['heading'], f'<div class="prose" style="font-size:14.5px">{prose(s["html"])}</div>') for s in L['sections'] if s['panel'])
    band = titlebar(cls, L['title'], f'Updated {escape(L["updated"])}', title_colour=D.CLASS_COLOR[cls], top=header(*VIEWER, 'Guides'))
    body = band + f'<main class="w gd"><p class="lede" style="margin-top:22px">{escape(L["identity"])}</p><div class="specs">{cards}</div><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page(f'Forever Sixty, {L["title"]} guides', GUIDE_CSS, body)


def index_page() -> str:
    cards = []
    for c in F['index']['classes']:
        specs = row_list([(f'<img class="ai" src="{D.icon_uri(s["icon"])}" alt="" width="28" height="28">', f'<a href="#" style="color:#fff">{escape(s["name"])}</a>', '', '')
                          for s in c['specs']])
        cards.append(f'<section class="card"><h2 style="color:{c["color"]}">{medallion(c["slug"], 34)}{escape(c["name"])}</h2>{specs}</section>')
    band = titlebar(None, F['markup']['indexTitle'], '', top=header(*VIEWER, 'Guides'))
    return page('Forever Sixty, guides', GUIDE_CSS, band + f'<main class="w gd"><div class="idx">{"".join(cards)}</div></main>' + footer())


def build() -> None:
    write('b2-guide-spec', spec_page())
    write('b2-guide-class', class_page())
    write('b2-guides', index_page())


if __name__ == '__main__':
    build()

"""Mock batch 4b: setup, connect, the app sign-in hand-off, the content and legal pages, the changelog, 404, the
build-unavailable page, the guild landing and the creators pages, in the directions system (plan 2026-11-09-redesign).

Words and data are read at render time, never retyped:
- design/mocks/data/setup-facts.json (setup and connect: the steps, the companion downloads, the pairing code)
- design/mocks/data/pages-facts.json (about and the legal pages as rendered, the changelog's entries, the 404's copy)
- web/src/lib/app-links/handoff-copy.ts (the /app/auth hand-off) and web/src/pages/b-unavailable.astro (its sentence)
- design/mocks/guild_directory_fixtures.py (the guild landing's top guilds)
- design/mocks/data/creators-facts.json (the program and stats; node design/mocks/creators_facts.mjs)
Drawn by prims.py.
"""
import json
import re
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))  # design/mocks: the guild directory fixtures

import guild_directory_fixtures as GD  # noqa: E402
from mocklib import emblem  # noqa: E402
import data as D  # noqa: E402
from common import write  # noqa: E402
from prims import PRIMS_CSS, field, meter, panel, reading_section, row_list, settings_bar, table, tag, toc  # noqa: E402
from style import footer, header, page, titlebar  # noqa: E402

MOCKS = Path(__file__).resolve().parent.parent
SETUP = json.loads((MOCKS / 'data/setup-facts.json').read_text())
PAGES = json.loads((MOCKS / 'data/pages-facts.json').read_text())
HANDOFF = dict(re.findall(r"^\s+(\w+): ['\"](.+?)['\"],$", (D.ROOT / 'web/src/lib/app-links/handoff-copy.ts').read_text(), re.M))
UNAVAILABLE = ''.join(re.findall(r"'([^']+)'", (D.ROOT / 'web/src/pages/b-unavailable.astro').read_text().split('const sentence =', 1)[1].split("';", 1)[0] + "'"))
VIEWER = ('Thandrel', 'warrior', 60)
SIGNED_OUT = ('', 'warrior', None)


def creators_facts() -> dict:
    """The creator program's facts (creators_facts.mjs over lib/creators/copy.ts and the e2e fixtures)."""
    return json.loads((MOCKS / 'data/creators-facts.json').read_text())


SITE_CSS = PRIMS_CSS + '''
.st{padding-top:6px}.steps li{display:grid;grid-template-columns:40px minmax(0,1fr) auto;gap:14px;align-items:center;padding:14px 0;border-bottom:1px solid rgba(255,255,255,.08)}
.steps li:last-child{border-bottom:0}.steps .n{width:32px;height:32px;border-radius:50%;display:flex;align-items:center;justify-content:center;font:600 15px var(--body);
border:1px solid var(--rule);color:var(--text)}.steps .n.next{background:var(--teal);border-color:var(--teal-edge);color:#fff}.steps .n.done{background:var(--gain);border-color:var(--gain);color:#000}
.steps b{display:block;color:#fff;font-size:16px}.steps span.l{display:block;font-size:14px;color:var(--muted)}
.msg{max-width:620px;margin:22px 0 0}.msg .bd{padding-top:4px}.msg p+p{margin-top:10px}
.code{font:600 30px ui-monospace,Menlo,monospace;letter-spacing:.12em;color:#fff;margin:6px 0 10px}
.cl .ent{padding:14px 0;border-bottom:1px solid rgba(255,255,255,.08)}.cl .ent:last-child{border-bottom:0}.cl .ent h3{font:600 16px var(--body);color:#fff;margin:2px 0 4px}
.cl .d{font-size:12.5px;color:var(--muted)}.cl .ent p{font-size:14.5px}.cl .ent a{font-size:14px}.page3{display:grid;grid-template-columns:200px minmax(0,1fr);gap:28px;margin-top:22px;align-items:start}
.page3 .toc{position:sticky;top:16px}.lg .prose{max-width:76ch}.page3>section{max-width:780px}.lg h2{font-size:24px;margin:22px 0 8px}.lg h2:first-child{margin-top:0}
.figs{display:flex;gap:32px;flex-wrap:wrap}.figs div{display:flex;flex-direction:column}.figs b{font:600 28px var(--body);color:#fff}.figs span{font-size:12.5px;color:var(--muted)}
.figs .lead b{color:var(--gold)}.figs{flex-wrap:nowrap;gap:24px}.lede{margin-top:22px;max-width:80ch}
.pick .setb .f b{border:1px solid var(--rule);border-radius:var(--r-sm);padding:5px 10px;margin-top:3px;display:inline-block}
@media(max-width:1100px){.page3{grid-template-columns:minmax(0,1fr)}.page3 .toc{display:none}}
@media(max-width:700px){.steps li{grid-template-columns:32px minmax(0,1fr)}.steps li>:last-child{grid-column:2}}
'''


def site_page(title: str, body: str) -> str:
    return page(f'Forever Sixty, {title}', SITE_CSS, body + footer())


# ------------------------------------------------------------------------------------------------ setup, connect, /app
def setup_page() -> str:
    """Signed out: the three steps, the first one next. The companion's download for this machine and its first-run line."""
    marks = {'next': 'next', 'todo': '', 'done': 'done'}
    steps = ''.join(f'<li><span class="n {marks[r["mark"]]}">{r["n"]}</span><span><b>{escape(r["title"])}</b>'
                    + ''.join(f'<span class="l">{escape(ln["text"])}</span>' for ln in r['lines'])
                    + f'</span><a class="{"btn" if r["action"]["primary"] else "btn2"}" href="#">{escape(r["action"]["label"])}</a></li>'
                    for r in SETUP['rows']['out'])
    own = SETUP['download_for']['windows']['own']
    others = SETUP['download_for']['windows']['others']
    comp = (f'<p class="muted" style="font-size:13px;margin-bottom:8px">{escape(own["meta"])}</p><p style="font-size:14px">{escape(own["firstRun"])}</p>'
            + ''.join(f'<p style="margin-top:10px"><a href="#">{escape(o["platform"])} &rsaquo;</a></p>' for o in others))
    main = panel('', f'<ul class="steps">{steps}</ul>', framed=True)
    rail = panel('The companion', comp)
    body = titlebar(None, PAGES['copy']['setupTitle'], '', top=header(*SIGNED_OUT, '')) + f'<main class="w st"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return site_page('setup', body)


def connect_page() -> str:
    c = SETUP['connect']
    inner = (f'<p class="muted">{escape(c["codeCheck"])}</p><p class="code">{escape(SETUP["code"])}</p>'
             f'<p style="font-size:14px">{escape(c["consent"])}</p><p style="margin-top:16px"><a class="btn" href="#">Connect</a></p>')
    card = panel('', inner, framed=True).replace('class="card pn', 'class="card pn msg', 1)
    body = titlebar(None, c['title'], '', top=header(*VIEWER, '')) + f'<main class="w st">{card}</main>'
    return site_page('connect', body)


def app_auth_page() -> str:
    body = (titlebar(None, HANDOFF['heading'], '', top=header(*VIEWER, ''))
            + '<main class="w st">' + panel('', f'<p>{escape(HANDOFF["body"])}</p><p style="margin-top:16px"><a class="btn" href="#">{escape(HANDOFF["button"])}</a></p>'
                                             f'<p class="muted" style="font-size:13.5px;margin-top:14px">{escape(HANDOFF["notYou"])}</p>', framed=True).replace('class="card pn', 'class="card pn msg', 1)
            + '</main>')
    return site_page('open the app', body)


# ------------------------------------------------------------------------------------------------ content, legal, changelog
def content_page(key: str, legal: bool) -> str:
    """About and the legal pages: the text exactly as rendered (never edited), in the reading column with its contents."""
    pg = PAGES['pages'][key]
    heads = re.findall(r'<h2[^>]*id="([^"]+)"[^>]*>(.*?)</h2>', pg['html'])
    tabs = [[(lk['label'], lk['label'].lower() == key) for lk in PAGES['copy']['legalLinks']]] if legal else None
    text = f'<section class="card pad lg" style="max-width:780px;margin-top:22px"><div class="prose">{pg["html"]}</div></section>'
    if heads:
        contents = toc([(re.sub('<[^>]+>', '', t), a) for a, t in heads], heads[0][0])
        text = f'<div class="page3"><aside>{contents}</aside>{text.replace("margin-top:22px", "margin-top:0")}</div>'
    body = titlebar(None, pg['title'], f'Updated {escape(pg["updated"])}', ctl=tabs, top=header(*VIEWER, '')) + f'<main class="w st">{text}</main>'
    return site_page(pg['title'], body)


def changelog_page() -> str:
    cl = PAGES['changelog']
    cc = PAGES['copy']['changelog']
    months = ''
    for m in cl['months'][:2]:
        shown = [e for e in m['entries'] if not e['folded']]
        if not shown:  # a month whose entries are all folded (pre-launch) is the fold's, not an empty panel
            continue
        ents = ''.join(f'<div class="ent"><span class="d">{escape(e["date"])}</span> {tag(e["kindLabel"], "gold") if e["show"] == "game" else ""}'
                       f'<h3>{escape(e["title"])}</h3>{e["html"]}' + ''.join(f'<a href="#">{escape(a["label"])}</a> ' for a in e['actions']) + '</div>'
                       for e in shown)
        months += panel(m['title'], f'<div class="cl">{ents}</div>')
    show = [[(cc['show'][k], k == 'all') for k in ('all', 'game', 'site')]]
    body = titlebar(None, cc['title'], f'Updated {escape(cl["updated"])}', ctl=show, top=header(*VIEWER, '')) + f'<main class="w st">{months}</main>'
    return site_page('changelog', body)


def not_found_page() -> str:
    """Did you mean (framed), then Pick your class: each class in its colour over its specs (each spec its guide), then the
    character lookup the shipped page keeps (P15 review round 1: the designer signs the specs and the lookup in)."""
    nf = PAGES['copy']['notFound']
    sug = PAGES['suggest']
    classes = PAGES['classes']
    look = PAGES['copy']['lookup']
    picks = row_list([(f'<img class="ai" src="{D.crest_uri(c["slug"])}" width="36" height="36" alt="" style="border-radius:50%">',
                       f'<a href="#" style="color:{D.CLASS_COLOR[c["slug"]]}">{escape(c["name"])}</a>',
                       ', '.join(f'<a href="#" class="muted">{escape(s["name"])}</a>' for s in c['specs']), '') for c in classes])
    did = panel(nf['didYouMean'], f'<p><a href="#">{escape(sug["page"][1])} &rsaquo;</a></p>', framed=True)
    lookup = (f'<section style="margin-top:28px"><h2 style="font:400 22px/28px var(--title);color:var(--cream);margin:0 0 12px">'
              f'{escape(look["heading"])}</h2>{field(look["namePlaceholder"], f"<a class=btn href=#>{escape(look['submit'])}</a>", look["nameLabel"])}</section>')
    body = (titlebar(None, nf['title'], escape(sug['from']), top=header(*VIEWER, ''))
            + f'<main class="w st">{did}{panel(nf["pick"], "<div style=columns:3;column-gap:24px>" + picks + "</div>")}{lookup}</main>')
    return site_page('page not found', body)


def unavailable_page() -> str:
    body = (titlebar(None, 'Build unavailable', '', top=header(*VIEWER, ''))
            + '<main class="w st">' + panel('', f'<p>{escape(UNAVAILABLE)}</p><p style="margin-top:16px"><a class="btn" href="#">Open the build planner</a>'
                                             f'<a href="#" style="margin-left:18px">Talent builds in the class guides &rsaquo;</a></p>', framed=True).replace('class="card pn', 'class="card pn msg', 1)
            + '</main>')
    return site_page('build unavailable', body)


# ------------------------------------------------------------------------------------------------ guild landing
def guild_landing_page() -> str:
    top = GD.TOP
    rows = row_list([(f'<img src="{emblem("alliance")}" width="30" height="30" alt="">', f'<a href="#" style="color:#fff">{escape(n)}</a>',
                      f'{reg.upper()} {GD.RULESET_WORD[rs]}', f'<span class="num">{k}</span> boss{"" if k == 1 else "es"} down')
                     for n, reg, rs, k in top])
    find = field('Guild name', '<a class="btn" href="#">Open</a>') + settings_bar([('Region', 'US'), ('Ruleset', 'Any')]).replace('card setb', 'setb" style="margin-top:6px;padding:8px 2px 0')
    main = panel('Open a guild', f'<div class="pick">{find}</div>', framed=True) + panel('Top guilds', rows)
    rail = panel('Your guilds', row_list([(f'<img src="{emblem("alliance")}" width="30" height="30" alt="">', '<a href="#" style="color:#fff">Olympus XXVII</a>', 'US PvP', tag('Officer', 'gold'))]))
    body = titlebar(None, 'Guilds', '', top=header(*VIEWER, 'Guild')) + f'<main class="w st"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return site_page('guilds', body)


# ------------------------------------------------------------------------------------------------ creators
def creators_page() -> str:
    pr = creators_facts()['program']
    secs = ''.join(reading_section(s['title'], ''.join(f'<p>{escape(p)}</p>' for p in s['paragraphs']), [], s['id']) for s in pr['sections'])
    contents = toc([(s['title'], s['id']) for s in pr['sections']], pr['sections'][0]['id'])
    body = (titlebar(None, pr['pageTitle'], '', right=f'<a class="btn" href="#">{escape(pr["apply"])}</a>', top=header(*VIEWER, ''))
            + f'<main class="w st"><div class="page3"><aside>{contents}</aside><section class="card" style="max-width:760px">'
            f'<p class="lede" style="margin:0;padding:22px 24px 0">{escape(pr["lead"])}</p>{secs}</section></div></main>')
    return site_page('creators', body)


def weeks_table(lines: list[str]) -> str:
    """'Week of 30 Nov: 61 signed up · 38 returning · 11 paying' as a table row; the newest week fills in over 14 days."""
    rows = []
    parsed = []
    for ln in lines:
        week, rest = ln.split(':', 1)
        nums = {k: int(v) for v, k in re.findall(r'(\d+) (signed up|returning|paying)', rest)}
        parsed.append((week, nums, 'so far' in rest))
    top = max((n.get('signed up', 0) for _, n, _ in parsed), default=1) or 1
    for week, n, partial in parsed:
        rows.append([escape(week) + (' <span class="muted">so far</span>' if partial else ''), meter(n.get('signed up', 0) / top),
                     str(n.get('signed up', 0)), str(n['returning']) if 'returning' in n else '&ndash;', str(n['paying']) if 'paying' in n else '&ndash;'])
    return table([('Week', False), ('', False), ('Signed up', True), ('Returning', True), ('Paying', True)], rows, hide_phone={1})


def creator_stats_page() -> str:
    st = creators_facts()['stats']
    figs = '<div class="figs">' + ''.join(f'<div class="{"lead" if f["lead"] else ""}"><b class="num">{escape(f["value"])}</b><span>{escape(f["label"])}'
                                          + (' &middot; estimate' if f['lead'] else '') + '</span></div>' for f in st['figures']) + '</div>'
    weeks = weeks_table(st['weeks'])
    c = st['copy']
    main = panel(c['summaryLabel'], figs, framed=True) + panel(c['playersTitle'], weeks, flush=True)
    rail = panel(c['payoutsTitle'], f'<p class="muted" style="font-size:14px">{escape(st["payoutsEmpty"])}</p><p style="margin-top:12px"><a class="btn2" href="#">{escape(st["dashboard"])}</a></p>')
    body = titlebar(None, st['name'], escape(st['meta']), top=header(*VIEWER, '')) + f'<main class="w st"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>'
    return site_page('creator stats', body)


def build() -> None:
    write('b4-setup', setup_page())
    write('b4-connect', connect_page())
    write('b4-app-auth', app_auth_page())
    write('b4-about', content_page('about', False))
    write('b4-privacy', content_page('privacy', True))
    write('b4-terms', content_page('terms', True))
    write('b4-refunds', content_page('refunds', True))
    write('b4-sources', content_page('sources', False))
    write('b4-changelog', changelog_page())
    write('b4-404', not_found_page())
    write('b4-unavailable', unavailable_page())
    write('b4-guilds', guild_landing_page())
    write('b4-creators', creators_page())
    write('b4-creator-stats', creator_stats_page())


if __name__ == '__main__':
    build()

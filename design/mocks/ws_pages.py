"""Web-layout mocks: the changelog (/changelog), the content pages (/terms, /privacy, /refunds, /about, /sources, /roadmap), the
404 page and the build-unavailable page, composed ONLY from ws_prims primitives.

Spec: design/specs/2026-11-08-pages.md. Every word is read from design/mocks/data/pages-facts.json, which
design/mocks/pages_facts.mjs writes from the collections the pages read (the markdown rendered by the processor Astro uses),
the copy modules and the pages' own markup literals. Nothing is typed in except the strings this round proposes (PROPOSED below,
listed in spec section 6.2) and the layout. The legal text is never edited: only its frame.

  node design/mocks/pages_facts.mjs && python3 design/mocks/ws_pages.py
  python3 design/mocks/ws_pages.py [mock ...]                 (every mock, or just those named)
  node design/mocks/ws_shoot.mjs <mock> 390,1440,1920,2000   (each name in `python3 design/mocks/ws_pages.py --list`)
  python3 design/mocks/ws_pages.py compare <before dir>      (before/after sheets; before: shoot_pages_before.mjs)
FS_PAGES_FACTS and FS_MOCK_OUT point the generator elsewhere (a parity spec's test-time render).
"""
import json
import os
import re
import sys
from datetime import date
from html import escape
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_prims import (Button, CharacterAnswer, Dated, DatedGroup, Disclosure, Field, FormRow, LaneBlock, Link, Notice, Page, PageHeader,
                      Plate, Prose, ReadingSection, Row, RowGrid, RowsSlim, Section, Select, Skeleton, TabStrip, Tag, Toc, build_page,
                      control_bar, crest_img, seg, sel, write, REGION)

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_PAGES_FACTS') or HERE / 'data' / 'pages-facts.json').read_text())
C = F['copy']
M = F['markup']
LOG = F['changelog']
ME = F['me']
SPEC_DOC = 'design/specs/2026-11-08-pages.md'

# The strings this round proposes (spec 6.2). The build adds them to lib/changelog/copy.ts and lib/not-found/copy.ts.
PROPOSED = {
    'showLabel': 'Show',
    'all': 'All', 'game': 'Game', 'site': 'Site',
    'summary': {'all': 'All changes', 'game': 'Game changes', 'site': 'Site changes'},
    'barLabel': 'Changelog choices',
    'changes': lambda n: f'{n} change' if n == 1 else f'{n} changes',
    'new': 'New',
    'empty': {'game': 'No game changes yet.', 'site': 'No site changes yet.'},
    'sections': lambda n: f'{n} sections',
    'didYouMean': 'Did you mean',
    'legalTabsLabel': 'Legal pages',
}
MONTHS = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
GAME = {'blizzard', 'community', 'datamined'}


def e(text: str) -> str:
    return escape(str(text), quote=False)


def rsaquo(text: str) -> str:
    return f'{e(text)} &rsaquo;'


# ------------------------------------------------------------------------------------------------ the changelog
def month_of(iso: str) -> tuple[str, str]:
    d = date.fromisoformat(iso)
    return f'{MONTHS[d.month - 1]} {d.year}', f'{d.year}-{d.month:02d}'


def entry_html(x: dict, new: bool = False) -> str:
    tags = Tag(e(x['kindLabel']), 'muted') + (Tag(PROPOSED['new']) if new else '')
    body = x['html'].strip()
    # An entry with no body reads its note as its text; an entry with a body keeps the note as the caveat beside its sources.
    text = Prose(body) if body else (f'<p class="dated-n">{e(x["note"])}</p>' if x['note'] else '')
    seen, rows = set(), []
    for s in x['sources']:
        if s['url'] in seen:
            continue
        seen.add(s['url'])
        # a source of the entry's own kind carries no tag (the date column already says it)
        tag = Tag(e(s['kindLabel']), 'muted') if s['kind'] != x['kind'] else ''
        rows.append(Row(f'<span class="gold link">{e(s["label"])}</span>', lead=tag, slim=True))
    caveat = f'<p class="dated-c">{e(x["note"])}</p>' if body and x['note'] else ''
    lane = LaneBlock(e(C['sources']), RowsSlim(rows)) + caveat
    return Dated(x['id'], x['iso'], e(x['date']), tags, e(x['title']), text, lane)


def months(entries: list[dict], new_after: str = '') -> list[tuple[str, str, list[dict]]]:
    out: list[tuple[str, str, list[dict]]] = []
    for x in entries:
        title, anchor = month_of(x['iso'])
        if not out or out[-1][1] != anchor:
            out.append((title, anchor, []))
        out[-1][2].append(x)
    return out


def changelog_bar(show: str, groups: list, counts: dict, open_m: bool = False) -> str:
    items = [(PROPOSED[k], k == show, counts[k], '') for k in ('all', 'game', 'site')]
    lab = PROPOSED['showLabel']
    toc_m = ('<span class="t-label" style="display:block;margin:8px 0 4px">' + e(C['tocBarLabel']) + '</span>'
             + RowsSlim([Row(e(t)) for t, _, _ in groups]))
    end = f'<span class="toc-sel">{sel(e(C["tocBarLabel"]))}</span>' if groups else ''
    return control_bar([(lab, seg(items, label=lab), seg(items, label=lab) + (toc_m if groups else ''))],
                       f'{PROPOSED["summary"][show]} &middot; {counts[show]}', end=end, open_m=open_m, label=PROPOSED['barLabel'])


def changelog_page(show: str = 'all', new_after: str = '', open_m: bool = False, force_empty: bool = False, who: tuple | None = None) -> str:
    entries = LOG['entries']
    counts = {'all': len(entries), 'game': sum(x['kind'] in GAME for x in entries), 'site': sum(x['kind'] == 'site' for x in entries)}
    shown = [x for x in entries if show == 'all' or (x['kind'] in GAME) == (show == 'game')]
    if force_empty:   # FORCED (spec 5.1): no filter is empty at this content; the state exists for a kind with no entry yet
        shown, counts = [], {**counts, show: 0}
    groups = months(shown)
    head = PageHeader(e(M['changelog']['h1']), meta=[f'Updated {e(LOG["updated"])}'], lead=e(M['changelog']['description']))
    if groups:
        body = ''.join(DatedGroup(t, PROPOSED['changes'](len(xs)), a, ''.join(entry_html(x, bool(new_after) and x['iso'] > new_after) for x in xs))
                       for t, a, xs in groups)
    else:
        body = f'<p class="empty cl-empty" data-region="{REGION["notes"]}">{PROPOSED["empty"][show]}</p>'
    toc = Toc(e(C['tocBarLabel']), [(t, a) for t, a, _ in groups], groups[0][1] if groups else '') if groups else ''
    html = head + changelog_bar(show, groups, counts, open_m) + Page('reading', body, toc, toc_label=C['tocBarLabel']) + PAGE_CSS
    note = (f'Changelog mock ({SPEC_DOC}), built only from ws_prims; every entry is web/src/content/changelog read through pages-facts.json. '
            f'Filter: {show}.' + (f' Returning visitor: entries after {new_after} carry New (localStorage, proposed CL4).' if new_after else '')
            + (' Empty state FORCED.' if force_empty else ''))
    return build_page(M['changelog']['title'], html, '', who, note)


# ------------------------------------------------------------------------------------------------ the content pages
H2 = re.compile(r'<h2(?: id="([^"]*)")?>(.*?)</h2>', re.S)
COACH_ANCHOR = 'log-coach'   # PROPOSED (gap PG3): CoachPrivacySection's h2 carries no id; the build adds this one


def split_sections(html: str) -> tuple[str, list[tuple[str, str, str]]]:
    parts = H2.split(html)
    intro, rest = parts[0], parts[1:]
    out = []
    for i in range(0, len(rest), 3):
        anchor, title, inner = rest[i], rest[i + 1], rest[i + 2]
        out.append((anchor or re.sub(r'[^a-z0-9]+', '-', re.sub(r'<[^>]+>', '', title).lower()).strip('-') or COACH_ANCHOR, title, inner))
    return intro.strip(), out


LEGAL = ('terms', 'privacy', 'refunds')


def content_page(pid: str, current: str = '', open_toc: bool = False, scroll_to: str = '') -> str:
    p = F['pages'][pid]
    intro, secs = split_sections(p['html'])
    tabs = ''
    if pid in LEGAL:
        labels = [l['label'] for l in C['legalLinks']]
        active = next(l['label'] for l in C['legalLinks'] if l['href'] == f'/{pid}')
        tabs = TabStrip(labels, active).replace('<nav class="strip-tabs ph-tabs">', f'<nav class="strip-tabs ph-tabs" aria-label="{PROPOSED["legalTabsLabel"]}">', 1)
    # The description is the lead only when the body opens on a heading (the legal pages); a page that opens on a paragraph
    # leads with it (About, Data sources, Roadmap), so the head never says the same thing twice (decision L3).
    head = PageHeader(e(p['title']), meta=[f'Updated {e(p["updated"])}'], tabs=tabs, lead='' if intro else e(p['description']))
    body = f'<div class="doc-intro" data-region="{REGION["intro"]}">{Prose(intro)}</div>' if intro else ''
    contents = ''
    if len(secs) > 1:
        contents = ('<div class="doc-toc-m">' + Disclosure(e(C['tocBarLabel']), RowsSlim([Row(t) for _, t, _ in secs]), open_toc,
                                                          PROPOSED['sections'](len(secs))) + '</div>')
    body = contents + body + ''.join(ReadingSection(t, Prose(inner), '', a) for a, t, inner in secs)
    toc_from = 1024 if len(secs) > 1 else 0
    toc = Toc(e(C['tocBarLabel']), [(t, a) for a, t, _ in secs], current or (secs[0][0] if secs else '')) if len(secs) > 1 else ''
    jump = (f'<script>document.getElementById({json.dumps(scroll_to)}).scrollIntoView()</script>' if scroll_to else '')
    html = head + Page('reading', body, toc, toc_from=toc_from, toc_label=C['tocBarLabel']) + PAGE_CSS + jump
    note = (f'Content page mock ({SPEC_DOC}): web/src/content/pages/{pid}.md rendered by the processor Astro uses (pages-facts.json); '
            'the text is unchanged, only its frame.' + (f' Arrived at #{scroll_to}.' if scroll_to else ''))
    return build_page(p['title'], html, '', None, note)


# ------------------------------------------------------------------------------------------------ 404 and build unavailable
def answer(kind: str) -> str:
    if kind == 'loading':
        return CharacterAnswer(C['answerLabel'], 'loading')
    if kind == 'error':
        return CharacterAnswer(C['answerLabel'], 'error', text=e(C['answer']['error']), action=Link(e(C['answer']['retry'])))
    if kind == 'nochar':
        return CharacterAnswer(C['answerLabel'], 'message', text=e(C['noCharacter']), action=Link(rsaquo(C['setupTitle'])))
    cls = ME['class'].lower()
    return CharacterAnswer(C['answerLabel'], 'who', crest_img(cls, 40), e(ME['name']), CLASS_COLOR[cls], f'{e(ME["level_line"])} &middot; {e(ME["identity"])}',
                           [Link(e(C['callout'])), Link(rsaquo(C['answer']['bis']))])


def class_picker() -> str:
    rows = [Row(e(c['name']), ', '.join(e(s['name']) for s in c['specs']), lead=crest_img(c['slug'], 44), title_color=c['color']) for c in F['classes']]
    return Section(e(M['notFound']['pick']), Plate(RowGrid(rows, 3), e(M['notFound']['pick'])))


def lookup() -> str:
    L = C['lookup']
    ruleset = next(r['label'] for r in C['rulesets'] if r['id'] == 'normal') if isinstance(C['rulesets'][0], dict) else 'Normal'
    form = FormRow([(e(L['nameLabel']), Field(e(L['namePlaceholder']), True), True), (e(L['regionLabel']), Select(C['regions'][0].upper()), False),
                    (e(L['rulesetLabel']), Select(e(ruleset)), False)], Button(e(L['submit']), primary=True))
    return Section(e(L['heading']), f'<p class="t-small nf-aside">{e(L["aside"])}</p>{form}<p class="t-small nf-hint">{e(L["hint"])}</p>')


SUGGEST = ('/guides/hunter/beastmastery', 'Beast Mastery Hunter guide')   # SAMPLE (spec 4.3 N3): a mistyped guide address


def not_found_page(kind: str = 'out', suggest: bool = False) -> str:
    nf = M['notFound']
    path = SUGGEST[0] if suggest else '/this-page-does-not-exist'
    head = PageHeader(e(nf['h1']), meta=[e(path)], answer='' if kind == 'out' else answer(kind), lead=e(nf['job']))
    did = (f'<div data-region="{REGION["notes"]}">' + Notice(PROPOSED['didYouMean'], Link(rsaquo(SUGGEST[1])), 'neutral') + '</div>') if suggest else ''
    main = class_picker() + f'<div class="section">{lookup()}</div>'
    who = (ME['name'], ME['class'].lower(), ME['level_line']) if kind != 'out' else None
    html = head + did + Page('ledger', main) + PAGE_CSS
    note = f'404 mock ({SPEC_DOC}); copy from 404.astro and the copy modules (pages-facts.json). State: {kind}' + (', did-you-mean (proposed N3).' if suggest else '.')
    return build_page(nf['docTitle'], html, '', who, note)


def unavailable_page() -> str:
    u = M['unavailable']
    head = PageHeader(e(u['h1']), lead=e(u['paragraphs'][0]))
    body = f'<div data-region="{REGION["intro"]}">{Prose("<p>" + u["paragraphs"][1] + "</p>")}</div>'
    return build_page(u['docTitle'], head + Page('ledger', body) + PAGE_CSS, '', None, f'Build unavailable mock ({SPEC_DOC}); copy verbatim from b-unavailable.astro.')


# ------------------------------------------------------------------------------------------------ every element state
def states_page() -> str:
    blocks = []

    def block(title: str, html: str) -> None:
        blocks.append(f'<section class="st-b"><h2 class="t-h3">{title}</h2>{html}</section>')
    c = {'all': 38, 'game': 24, 'site': 14}
    segs = [seg([(PROPOSED[k], k == 'all', c[k], '') for k in c], label='Show'),
            seg([(PROPOSED[k], k == 'all', c[k], '') for k in c], label='Show').replace('<a href="#" role="radio" aria-checked="false" class="">', '<a href="#" role="radio" aria-checked="false" class="hov">', 1),
            seg([(PROPOSED[k], k == 'all', c[k], '') for k in c], label='Show').replace('<a href="#" role="radio" aria-checked="false" class="">', '<a href="#" role="radio" aria-checked="false" class="fv">', 1)]
    block('Show (changelog bar): selected All; hover on Game; focus-visible on Game', '<div class="st-row">' + ''.join(segs) + '</div>')
    x = LOG['entries']
    block('Entry: default; a returning visitor&rsquo;s New entry; the permalink date hovered (gold)', '<div class="st-entries">'
          + entry_html(x[2]) + entry_html(x[0], True) + entry_html(x[1]).replace('class="d"', 'class="d st-hov"', 1) + '</div>')
    block('Entry: a game entry with three sources and its caveat; an entry whose only text is its note',
          '<div class="st-entries">' + entry_html(next(y for y in x if y['id'] == '2026-09-13-item-and-loot-pass'))
          + entry_html(next(y for y in x if y['id'] == '2026-09-12-forever-announced')) + '</div>')
    block('Filtered to a kind with no entry (FORCED)', f'<p class="empty cl-empty">{PROPOSED["empty"]["site"]}</p>')
    privacy = F['pages']['privacy']
    _, secs = split_sections(privacy['html'])
    block('Contents rail: the current section (gold, aria-current); a hovered link', '<nav class="toc st-toc">'
          + Toc(e(C['tocBarLabel']), [(t, a) for a, t, _ in secs], 'cookies').replace('<a href="#data-retention"', '<a class="st-hov2" href="#data-retention"', 1) + '</nav>')
    block('Contents on a phone: closed; open', '<div class="st-row st-col">' + Disclosure(e(C['tocBarLabel']), '', False, PROPOSED['sections'](len(secs)))
          + Disclosure(e(C['tocBarLabel']), RowsSlim([Row(t) for _, t, _ in secs]), True, PROPOSED['sections'](len(secs))) + '</div>')
    labels = [l['label'] for l in C['legalLinks']]
    block('Legal tabs: Privacy current; Terms hovered; Refunds focus-visible', TabStrip(labels, 'Privacy').replace('<a href="#" class="">', '<a href="#" class="st-tab-h">', 1)
          .replace('<a href="#" class="">', '<a href="#" class="st-tab-f">', 1))
    for t, k in (('Your character: loading (reserved at its final size; the session is hinted)', 'loading'), ('Your character: signed in', 'who'),
                 ('Your character: no character on the account', 'nochar'), ('Your character: the account read failed', 'error')):
        block(t, '<div class="st-ans">' + answer(k) + '</div>')
    L = C['lookup']
    block('Look up: the empty name submitted; a name no character can have', f'<p class="t-small wipe" role="alert">{e(L["empty"])}</p><p class="t-small wipe" role="alert">{e(L["invalid"])}</p>')
    block('Did you mean (404, proposed N3)', Notice(PROPOSED['didYouMean'], Link(rsaquo(SUGGEST[1])), 'neutral'))
    block('Class row (404 plate): focus-visible; hover; default', '<div class="st-rows">' + RowGrid([
        Row(e(cc['name']), ', '.join(e(s['name']) for s in cc['specs']), lead=crest_img(cc['slug'], 44), title_color=cc['color']) for cc in F['classes'][:3]], 3)
        .replace('<a class="row"', '<a class="row hov"', 2).replace('<a class="row hov"', '<a class="row"', 1).replace('<a class="row"', '<a class="row fv"', 1) + '</div>')
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Changelog, content pages and 404: element states</h1></div>' + Page('ledger', f'<div class="statesheet">{"".join(blocks)}</div>') + PAGE_CSS
    return build_page('Pages states', body, '', None, f'States sheet ({SPEC_DOC}); data as the page mocks.')


PAGE_CSS = '''<style>
/* page-level layout only (no surfaces): spacing, column splits, what shows at which width */
main{padding-bottom:var(--section-gap)}
.toc-sel{display:inline-flex}@media(min-width:1760px){.toc-sel{display:none}}
@media(max-width:767px){.cbar .pc{padding-block:0;min-height:0}}
.cl-empty{padding:24px 0}
.doc-toc-m{margin-bottom:8px}@media(min-width:1024px){.doc-toc-m{display:none}}
.doc-intro+.rsec{margin-top:0}
.nf-aside{margin:12px 0 16px}.nf-hint{margin-top:10px}
.statesheet{display:grid;gap:40px}.st-b>h2{margin-bottom:12px}.st-row{display:flex;gap:24px;flex-wrap:wrap;align-items:flex-start}.st-col{flex-direction:column;max-width:420px}
.st-entries{max-width:1248px}.st-toc{display:block;width:220px}.st-ans{max-width:520px}.st-rows{max-width:1248px}
.seg a.hov{background:var(--hover);color:var(--strong)}.seg a.fv{outline:2px solid var(--gold);outline-offset:-2px}
.st-hov{color:var(--gold-hover)!important}.toc a.st-hov2{color:var(--strong)}
.strip-tabs a.st-tab-h{color:var(--strong)}.strip-tabs a.st-tab-f{outline:2px solid var(--gold);outline-offset:-2px}
.row.fv{outline:2px solid var(--gold);outline-offset:-2px}
</style>'''


PAGES = {
    'web-changelog': lambda: changelog_page(),
    'web-changelog-game': lambda: changelog_page('game'),
    'web-changelog-site': lambda: changelog_page('site'),
    'web-changelog-new': lambda: changelog_page(new_after='2026-10-27', who=(ME['name'], ME['class'].lower(), ME['level_line'])),
    'web-changelog-empty': lambda: changelog_page('site', force_empty=True),
    'web-changelog-bar': lambda: changelog_page(open_m=True),
    'web-legal-terms': lambda: content_page('terms'),
    'web-legal-privacy': lambda: content_page('privacy'),
    'web-legal-privacy-cookies': lambda: content_page('privacy', current='cookies', scroll_to='cookies'),
    'web-legal-privacy-toc': lambda: content_page('privacy', open_toc=True),
    'web-legal-refunds': lambda: content_page('refunds'),
    'web-about': lambda: content_page('about'),
    'web-about-sources': lambda: content_page('sources'),
    'web-about-roadmap': lambda: content_page('roadmap'),
    'web-404': lambda: not_found_page(),
    'web-404-in': lambda: not_found_page('who'),
    'web-404-loading': lambda: not_found_page('loading'),
    'web-404-error': lambda: not_found_page('error'),
    'web-404-nochar': lambda: not_found_page('nochar'),
    'web-404-suggest': lambda: not_found_page(suggest=True),
    'web-b-unavailable': unavailable_page,
    'web-pages-states': states_page,
}


def compare(before: Path) -> None:
    """Before/after at the same scale: live captures (shoot_pages_before.mjs) beside the mock renders."""
    import ws_compose
    ws_compose.BEFORE = before
    pairs = [('changelog', 'changelog', 'web-changelog'), ('privacy', 'privacy', 'web-legal-privacy'), ('about', 'about', 'web-about'),
             ('404', '404', 'web-404')]
    for name, live, mock in pairs:
        for w, sc in ((1440, .5), (1920, .42), (2000, .4), (390, .8)):
            print(ws_compose.compose(name, live, mock, w, sc, max_h=6000))
    for name, live, mock in (('404-in', '404-in', 'web-404-in'), ('terms', 'terms', 'web-legal-terms'), ('refunds', 'refunds', 'web-legal-refunds'),
                             ('sources', 'sources', 'web-about-sources'), ('roadmap', 'roadmap', 'web-about-roadmap'), ('b-unavailable', 'b-unavailable', 'web-b-unavailable')):
        for w, sc in ((1440, .5), (390, .8)):
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
            sys.exit(f'ws_pages.py: no such mock: {" ".join(unknown)} (see --list)')
        for name, fn in PAGES.items():
            if not only or name in only:
                print(write(f'{name}.html', fn()))

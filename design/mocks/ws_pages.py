"""Web-layout mocks: the changelog (/changelog), the content pages (/terms, /privacy, /refunds, /about, /roadmap), the
404 page and the build-unavailable page, composed ONLY from ws_prims primitives.

Spec: design/specs/2026-11-08-pages.md (section 0A: the rulings). Every word is read from design/mocks/data/pages-facts.json,
which design/mocks/pages_facts.mjs writes from the collections the pages read (the markdown rendered by the processor Astro
uses), the pages' copy modules (lib/changelog/copy.ts, lib/not-found/copy.ts) and their models (lib/changelog/entries.ts: the
order, months, counts, folded pre-launch entries and the site entries' action links). Nothing is typed in except the SAMPLE
returning visitor's last visit and the layout. The legal text is never edited: only its frame.

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
from ws_lib import main_foot_padding
from ws_prims import (Button, CharacterAnswer, Dated, DatedGroup, Disclosure, Field, FormRow, LaneBlock, Link, Notice, Page, PageHeader,
                      PageForm, Plate, Prose, ReadingPanel, ReadingSection, Row, RowGrid, RowsSlim, Section, Select, Skeleton, TabStrip, Tag, Toc, build_page,
                      control_bar, crest_img, seg, sel, write, REGION)

HERE = Path(__file__).resolve().parent
F = json.loads(Path(os.environ.get('FS_PAGES_FACTS') or HERE / 'data' / 'pages-facts.json').read_text())
C = F['copy']
CL = C['changelog']
NF = C['notFound']
LOG = F['changelog']
ME = F['me']
SPEC_DOC = 'design/specs/2026-11-08-pages.md'


def e(text: str) -> str:
    return escape(str(text), quote=False)


def rsaquo(text: str) -> str:
    return f'{e(text)} &rsaquo;'


# ------------------------------------------------------------------------------------------------ the changelog
def lane_html(x: dict) -> str:
    """The sources (a source tagged only when its kind differs from the entry's), this site's pages as action links (ruling
    Q5), then the caveat (the note of an entry with a body)."""
    parts = []
    if x['sources']:
        rows = [Row(f'<span class="gold link">{e(s["label"])}</span>', lead=Tag(e(s['tag']), 'muted') if s.get('tag') else '', slim=True)
                for s in x['sources']]
        parts.append(LaneBlock(e(C['sources']), RowsSlim(rows)))
    parts += [f'<p class="dated-a">{Link(e(a["label"]))}</p>' for a in x['actions']]
    if x.get('caveat'):
        parts.append(f'<p class="dated-c">{e(x["caveat"])}</p>')
    return ''.join(parts)


def entry_html(x: dict, new: bool = False) -> str:
    # P15 round 2: a pill only where the source changes the reading; this site's own entries carry none
    tags = (Tag(e(x['kindLabel']), 'muted') if x['show'] != 'site' else '') + (Tag(e(CL['newTag'])) if new else '')
    text = Prose(x['html']) if x['html'] else f'<p class="dated-n">{e(x.get("noteAsBody") or "")}</p>'
    return Dated(x['id'], x['iso'], e(x['date']), tags, e(x['title']), text, lane_html(x), x['folded'])


def shown_months(show: str) -> list[tuple[str, str, list[dict], str]]:
    return [(m['title'], m['anchor'], [x for x in m['entries'] if show == 'all' or x['show'] == show], m['meta'][show])
            for m in LOG['months'] if m['counts'][show] > 0]


def coming_up() -> str:
    """Roadmap's dates still ahead (ruling Q7), its table in Prose, the source in the lane."""
    cu = LOG['comingUp']
    rows = ''.join(f'<tr><td>{e(r["when"])}</td><td>{e(r["what"])}</td></tr>' for r in cu['rows'])
    table = f'<table><thead><tr>{"".join(f"<th>{e(h)}</th>" for h in cu["heads"])}</tr></thead><tbody>{rows}</tbody></table>'
    lane = LaneBlock(e(C['sources']), RowsSlim([Row(f'<span class="gold link">{e(cu["source"]["label"])}</span>', slim=True)]))
    return ReadingSection(e(CL['comingUp']), Prose(table), lane, 'coming-up')


def changelog_bar(show: str, groups: list, counts: dict, open_m: bool = False) -> str:
    items = [(e(CL['show'][k]), k == show, counts[k], '') for k in ('all', 'game', 'site')]
    lab = e(CL['showLabel'])
    # the phone stack's months carry their count in the filter (ruling SHOULD)
    toc_m = ('<span class="t-label" style="display:block;margin:8px 0 4px">' + e(C['tocBarLabel']) + '</span>'
             + RowsSlim([Row(e(t), trail=f'<span class="t-small">{e(meta)}</span>') for t, _, _, meta in groups]))
    end = f'<span class="toc-sel">{sel(e(C["tocBarLabel"]))}</span>' if groups else ''
    summary = CL['summary'][show] if counts[show] else f'{CL["summaryWords"][show]} &middot; 0'
    return control_bar([(lab, seg(items, label=lab), seg(items, label=lab) + (toc_m if groups else ''))],
                       e(summary).replace(' · ', ' &middot; '), end=end, open_m=open_m, label=e(CL['barLabel']))


def changelog_page(show: str = 'all', new_after: str = '', open_m: bool = False, force_empty: bool = False, who: tuple | None = None) -> str:
    counts = dict(LOG['counts'])
    groups = shown_months(show)
    if force_empty:   # FORCED (spec 5.1): no filter is empty at this content; the state exists for a kind with no entry yet
        groups, counts = [], {**counts, show: 0}
    head = PageHeader(e(CL['title']), meta=[f'Updated {e(LOG["updated"])}'])
    # Round 5 (directions mock b4-changelog): each month, and Coming up, on its own reading panel.
    upcoming = ReadingPanel(coming_up()) if show != 'site' else ''
    if groups:
        body = upcoming + ''.join(ReadingPanel(DatedGroup(e(t), e(meta), a, ''.join(entry_html(x, bool(new_after) and x['iso'] > new_after) for x in xs)))
                                  for t, a, xs, meta in groups)
    else:
        body = upcoming + f'<p class="empty cl-empty" data-region="{REGION["notes"]}">{e(CL["empty"][show])}</p>'
    toc = Toc(e(C['tocBarLabel']), [(t, a) for t, a, _, _ in groups], groups[0][1] if groups else '') if groups else ''
    html = head + changelog_bar(show, groups, counts, open_m) + Page('reading', body, toc, toc_label=C['tocBarLabel']) + PAGE_CSS
    note = (f'Changelog mock ({SPEC_DOC}), built only from ws_prims; every entry is web/src/content/changelog read through pages-facts.json '
            f'(lib/changelog/entries.ts). Filter: {show}. Entries before launch (4 Nov) are title rows that open in place (ruling CL1).'
            + (f' Returning visitor (SAMPLE last visit {new_after}): later entries carry New (localStorage, CL4).' if new_after else '')
            + (' Empty state FORCED.' if force_empty else ''))
    return build_page(CL['docTitle'], html, '', who, note)


# ------------------------------------------------------------------------------------------------ the content pages
H2 = re.compile(r'<h2 id="([^"]*)">(.*?)</h2>', re.S)


def split_sections(html: str) -> tuple[str, list[tuple[str, str, str]]]:
    """lib/content-pages/sections.ts splitDocument: every h2 carries its id (the Log coach section's too, gap PG3 closed)."""
    parts = H2.split(html)
    intro, rest = parts[0], parts[1:]
    return intro.strip(), [(rest[i], rest[i + 1], rest[i + 2].strip()) for i in range(0, len(rest), 3)]


LEGAL = ('terms', 'privacy', 'refunds')


def content_page(pid: str, current: str = '', open_toc: bool = False, scroll_to: str = '') -> str:
    p = F['pages'][pid]
    intro, secs = split_sections(p['html'])
    tabs = ''
    if pid in LEGAL:
        labels = [l['label'] for l in C['legalLinks']]
        active = next(l['label'] for l in C['legalLinks'] if l['href'] == f'/{pid}')
        tabs = TabStrip(labels, active).replace('<nav class="strip-tabs ph-tabs">', f'<nav class="strip-tabs ph-tabs" aria-label="{e(C["legalTabsLabel"])}">', 1)
    # The description is the lead only when the body opens on a heading (the legal pages); a page that opens on a paragraph
    # leads with it (About, Roadmap), so the head never says the same thing twice (decision L3).
    head = PageHeader(e(p['title']), meta=[f'Updated {e(p["updated"])}'], tabs=tabs)
    body = f'<div class="doc-intro" data-region="{REGION["intro"]}">{Prose(intro)}</div>' if intro else ''
    contents = ''
    if len(secs) > 1:
        contents = ('<div class="doc-toc-m">' + Disclosure(e(C['tocBarLabel']), RowsSlim([Row(t) for _, t, _ in secs]), open_toc,
                                                          e(p['sectionsMeta'])) + '</div>')
    # Round 5 (directions mocks b4-about, b4-privacy): the text on one reading panel at the reading measure.
    body = contents + '<div class="doc-panel">' + ReadingPanel(body + ''.join(ReadingSection(t, Prose(inner), '', a) for a, t, inner in secs)) + '</div>'
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
    # your guide and your best in slot, then the smaller Your character (ruling Q6)
    return CharacterAnswer(C['answerLabel'], 'who', crest_img(cls, 40), e(ME['name']), CLASS_COLOR[cls], f'{e(ME["level_line"])} &middot; {e(ME["identity"])}',
                           [Link(e(C['callout'])), Link(rsaquo(C['answer']['bis'])), f'<a class="link cans-sm" href="#">{rsaquo(NF["yourCharacter"])}</a>'])


def class_row(c: dict) -> str:
    """A class row of the 404 picker: crest, the class (its landing) in its colour, each spec a link to its guide (SHOULD)."""
    specs = '<span class="nf-specs">' + ', '.join(f'<a class="nf-spec" href="#">{e(s["name"])}</a>' for s in c['specs']) + '</span>'
    return Row(f'<a class="nf-cls" href="#" style="color:{c["color"]}">{e(c["name"])}</a>', specs, lead=crest_img(c['slug'], 44), link=False)


def class_picker() -> str:
    return Section(e(NF['pick']), RowGrid([class_row(c) for c in F['classes']], 3), panel=True, flush=True)


def lookup() -> str:
    L = C['lookup']
    ruleset = next(r['label'] for r in C['rulesets'] if r['id'] == 'normal') if isinstance(C['rulesets'][0], dict) else 'Normal'
    # FormRow as the shipped lookup is built: a form, the name a combobox (the typeahead), a submit button, then
    # the status line its two validation sentences are written into (empty until a submit)
    name = f'<span class="field grow" role="combobox" aria-label="{e(L["nameLabel"])}">{e(L["namePlaceholder"])}</span>'
    form = (FormRow([(e(L['nameLabel']), name, True), (e(L['regionLabel']), Select(C['regions'][0].upper()), False),
                     (e(L['rulesetLabel']), Select(e(ruleset)), False)], Button(e(L['submit']), primary=True))
            .replace('<div class="formrow">', f'<form class="formrow" aria-label="{e(L["heading"])}">', 1).removesuffix('</div>') + '</form>')
    form = form.replace('<a class="btn primary" href="#">', '<button type="submit" class="btn primary">', 1).replace(f'{e(L["submit"])}</a>', f'{e(L["submit"])}</button>', 1)
    return Section(e(L['heading']), f'<p class="t-small nf-aside">{e(L["aside"])}</p>{form}'
                   '<p class="t-small nf-status" role="status"></p>')


SUGGEST = (F['suggest']['from'], F['suggest']['page'][1])   # SAMPLE address (pages_facts.mjs); its match is lib/not-found/suggest.ts's
CHARACTER_PATH = '/character/us/normal'   # SAMPLE: a dead character address (no name), signed out


def not_found_page(kind: str = 'out', suggest: bool = False, character: bool = False) -> str:
    path = SUGGEST[0] if suggest else (CHARACTER_PATH if character else '/this-page-does-not-exist')
    head = PageHeader(e(NF['title']), meta=[e(path)], answer='' if kind == 'out' else answer(kind))
    did = (f'<div data-region="{REGION["notes"]}"><div class="plate framed nf-did" data-region="{REGION["plate"]}">'
           f'<h2 class="nf-suggest-title">{e(NF["didYouMean"])}</h2>{Link(rsaquo(SUGGEST[1]))}</div></div>') if suggest else ''
    # signed out on a dead /character/ address the lookup comes first (SHOULD)
    first_lookup = character and kind == 'out'
    main = (lookup() + f'<div class="section">{class_picker()}</div>') if first_lookup else (class_picker() + f'<div class="section">{lookup()}</div>')
    who = (ME['name'], ME['class'].lower(), ME['level_line']) if kind != 'out' else None
    html = head + Page('ledger', did + main) + PAGE_CSS
    note = (f'404 mock ({SPEC_DOC}); copy from lib/not-found/copy.ts and the shared copy modules (pages-facts.json). State: {kind}'
            + (', did-you-mean (N3).' if suggest else '') + (', a dead /character/ address: the lookup first.' if character else '.'))
    return build_page(NF['docTitle'], html, '', who, note)


# ------------------------------------------------------------------------------------------------ sign in
def lookup_panel() -> str:
    """/characters' lookup (round 5, directions mock b2-characters; web CharacterSearch `panel`): the page's framed Plate under
    the title band, no heading of its own, no aside or hint: the name (its label for assistive tech) and Look up on one row,
    Region and Ruleset under them, then the status line."""
    L = C['lookup']
    ruleset = next(r['label'] for r in C['rulesets'] if r['id'] == 'normal') if isinstance(C['rulesets'][0], dict) else 'Normal'
    name = f'<span class="field grow" role="combobox" aria-label="{e(L["nameLabel"])}">{e(L["namePlaceholder"])}</span>'
    field = lambda lab, html: f'<label class="f"><span class="t-label">{lab}</span>{html}</label>'  # noqa: E731
    form = (f'<form class="cs-p" aria-label="{e(L["heading"])}"><label class="f grow cs-n">{name}</label>'
            f'<button type="submit" class="btn primary cs-b">{e(L["submit"])}</button>'
            f'<div class="cs-s">{field(e(L["regionLabel"]), Select(C["regions"][0].upper()))}{field(e(L["rulesetLabel"]), Select(e(ruleset)))}</div></form>')
    inner = f'<div class="cs-panel">{form}<p class="t-small nf-status" role="status"></p></div>'
    return Plate(inner, e(L['heading']))


def characters_page() -> str:
    """/characters, signed out (round 5, directions mock b2-characters): the title band names the page; main holds the lookup's
    framed panel. Signed in, the rail lists the account's characters (character/YourCharacters); signed out it draws nothing."""
    body = PageHeader(e(C['lookup']['pageTitle'])) + Page('split', lookup_panel(), panels=True) + PAGE_CSS
    body += ('<style>.cs-panel{padding:22px 24px}.cs-panel .nf-status{min-height:20px;margin:0}.cs-p{display:grid;grid-template-columns:minmax(0,1fr) auto;'
             'grid-template-areas:"n b" "s s";gap:12px;align-items:end}.cs-n{grid-area:n}.cs-b{grid-area:b}'
             '.cs-s{grid-area:s;display:flex;gap:12px}.cs-s .f{display:flex;flex-direction:column;gap:6px}'
             '.cs-p .btn,.cs-s .sel{min-height:44px;box-sizing:border-box}</style>')
    return build_page(e(C['lookup']['heading']), body, '', None, '/characters mock; data pages-facts.json.')


def login_page() -> str:
    """/login signed out: the form page as /account signed out draws it (ws_premium account_page('out')), one Plate holding the
    reason, the Battle.net button, what Battle.net shares, and the reserved status line. Copy: lib/account/signin-copy.ts."""
    L = C['login']
    inner = (f'<div class="lg-in"><p class="lede">{e(L["reason"])}</p>{Button(e(L["action"]), primary=True)}'
             f'<p class="t-small">{e(L["privacy"])}</p><p class="lg-status" role="status"></p></div>')
    html = PageHeader(e(L['title'])) + PageForm(Plate(inner, e(L['title']))) + PAGE_CSS
    return build_page(L['title'], html, '', None, '/login mock, signed out. Copy: lib/account/signin-copy.ts, my-reports-copy.ts signInAction.')


# ------------------------------------------------------------------------------------------------ every element state
def states_page() -> str:
    blocks = []

    def block(title: str, html: str) -> None:
        blocks.append(f'<section class="st-b"><h2 class="t-h3">{title}</h2>{html}</section>')
    c = LOG['counts']
    segs = [seg([(e(CL['show'][k]), k == 'all', c[k], '') for k in c], label='Show'),
            seg([(e(CL['show'][k]), k == 'all', c[k], '') for k in c], label='Show').replace('<a href="#" role="radio" aria-checked="false" class="">', '<a href="#" role="radio" aria-checked="false" class="hov">', 1),
            seg([(e(CL['show'][k]), k == 'all', c[k], '') for k in c], label='Show').replace('<a href="#" role="radio" aria-checked="false" class="">', '<a href="#" role="radio" aria-checked="false" class="fv">', 1)]
    block('Show (changelog bar): selected All; hover on Game; focus-visible on Game', '<div class="st-row">' + ''.join(segs) + '</div>')
    x = [y for m in LOG['months'] for y in m['entries']]
    block('Entry: default; a returning visitor&rsquo;s New entry; the permalink date hovered (gold)', '<div class="st-entries">'
          + entry_html(x[2]) + entry_html(x[0], True) + entry_html(x[1]).replace('class="d"', 'class="d st-hov"', 1) + '</div>')
    block('Entry: a game entry with three sources and its caveat; an entry whose only text is its note',
          '<div class="st-entries">' + entry_html(next(y for y in x if y['id'] == '2026-09-13-item-and-loot-pass'))
          + entry_html(next(y for y in x if y['id'] == '2026-09-12-forever-announced')) + '</div>')
    block('Filtered to a kind with no entry (FORCED)', f'<p class="empty cl-empty">{e(CL["empty"]["site"])}</p>')
    privacy = F['pages']['privacy']
    _, secs = split_sections(privacy['html'])
    block('Contents rail: the current section (gold, aria-current); a hovered link', '<nav class="toc st-toc">'
          + Toc(e(C['tocBarLabel']), [(t, a) for a, t, _ in secs], 'cookies').replace('<a href="#data-retention"', '<a class="st-hov2" href="#data-retention"', 1) + '</nav>')
    block('Contents on a phone: closed; open', '<div class="st-row st-col">' + Disclosure(e(C['tocBarLabel']), '', False, e(privacy['sectionsMeta']))
          + Disclosure(e(C['tocBarLabel']), RowsSlim([Row(t) for _, t, _ in secs]), True, e(privacy['sectionsMeta'])) + '</div>')
    labels = [l['label'] for l in C['legalLinks']]
    block('Legal tabs: Privacy current; Terms hovered; Refunds focus-visible', TabStrip(labels, 'Privacy').replace('<a href="#" class="">', '<a href="#" class="st-tab-h">', 1)
          .replace('<a href="#" class="">', '<a href="#" class="st-tab-f">', 1))
    for t, k in (('Your character: loading (reserved at its final size; the session is hinted)', 'loading'), ('Your character: signed in', 'who'),
                 ('Your character: no character on the account', 'nochar'), ('Your character: the account read failed', 'error')):
        block(t, '<div class="st-ans">' + answer(k) + '</div>')
    L = C['lookup']
    block('Look up: the empty name submitted; a name no character can have', f'<p class="t-small wipe" role="alert">{e(L["empty"])}</p><p class="t-small wipe" role="alert">{e(L["invalid"])}</p>')
    block('Did you mean (404, proposed N3)', Notice(e(NF['didYouMean']), Link(rsaquo(SUGGEST[1])), 'neutral'))
    block('Class row (404 plate): the class link focus-visible; a spec link hovered; default', '<div class="st-rows">' + RowGrid([class_row(cc) for cc in F['classes'][:3]], 3)
        .replace('class="nf-cls"', 'class="nf-cls fv"', 1).replace('class="nf-spec"', 'class="nf-spec"', 3).replace('<a class="nf-spec" href="#">', '<a class="nf-spec hov" href="#">', 4)
        .replace('<a class="nf-spec hov" href="#">', '<a class="nf-spec" href="#">', 3) + '</div>')
    block('Entry before launch: folded (a title row); opened in place', '<div class="st-entries">' + entry_html(x[-1])
          + entry_html(x[-2]).replace('<details class="dated-f"', '<details open class="dated-f"', 1) + '</div>')
    body = '<div class="pc" style="padding-top:24px"><h1 class="t-title">Changelog, content pages and 404: element states</h1></div>' + Page('ledger', f'<div class="statesheet">{"".join(blocks)}</div>') + PAGE_CSS
    return build_page('Pages states', body, '', None, f'States sheet ({SPEC_DOC}); data as the page mocks.')


PAGE_CSS = '''<style>
/* page-level layout only (no surfaces): spacing, column splits, what shows at which width */
''' + main_foot_padding('var(--section-gap)') + '''
.toc-sel{display:inline-flex}@media(min-width:1760px){.toc-sel{display:none}}
@media(max-width:767px){.cbar .pc{padding-block:0;min-height:0}}
.cl-empty{padding:24px 0}
.doc-toc-m{margin:calc(12px - var(--section-gap)) 0 8px}@media(min-width:1024px){.doc-toc-m{display:none}}
.doc-intro+.rsec{margin-top:0}
.doc-panel{max-width:calc(var(--measure-reading, 720px) + 48px)}.rpanel>.doc-intro{padding:22px 24px}.rpanel>.doc-intro+.rsec{border-top:1px solid var(--line-soft)}
@media(max-width:767px){.rpanel>.doc-intro{padding:18px 16px}}
.nf-did{padding:22px 24px}.nf-suggest-title{margin:0 0 6px;font:400 22px/28px var(--font-title);letter-spacing:-.015em;color:var(--cream)}
.nf-aside{margin:12px 0 16px}.nf-hint{margin-top:10px}.nf-status{min-height:20px;margin:0}
@media(max-width:767px){.formrow .sel,.formrow .field{height:44px}}
.nf-cls{font:700 15px/24px var(--font-body);display:inline-block}.nf-spec{display:inline-block;line-height:24px}.nf-cls:hover{text-decoration:underline}
.nf-specs{display:block;font-size:13px;line-height:18px;color:var(--muted)}.nf-spec{color:var(--muted)}.nf-spec:hover,.nf-spec.hov{color:var(--gold-hover)}
.nf-cls.fv,.nf-spec.fv{outline:2px solid var(--gold);outline-offset:2px}
.dated-a{margin:0}
.lg-in{display:flex;flex-direction:column;align-items:flex-start;gap:12px;padding:16px 20px 4px}.lg-in .lede{font-size:15px;line-height:24px;max-width:72ch}
.lg-status{min-height:20px;margin:-12px 0 0}
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
    'web-404': lambda: not_found_page(),
    'web-404-in': lambda: not_found_page('who'),
    'web-404-loading': lambda: not_found_page('loading'),
    'web-404-error': lambda: not_found_page('error'),
    'web-404-nochar': lambda: not_found_page('nochar'),
    'web-404-suggest': lambda: not_found_page(suggest=True),
    'web-404-character': lambda: not_found_page(character=True),
    'web-characters': characters_page,
    'web-login': login_page,
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
    for name, live, mock in (('404-in', '404-in', 'web-404-in'), ('terms', 'terms', 'web-legal-terms'), ('refunds', 'refunds', 'web-legal-refunds')):
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

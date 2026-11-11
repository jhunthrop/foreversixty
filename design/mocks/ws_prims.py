"""The shared primitives of the web layout system (design/specs/2026-11-05-web-layout.md section 4), as the mock library.

Every page mock after BiS is composed ONLY from these functions. Each emits classes defined once in PRIM_CSS; a page
generator never writes a border, background, radius or shadow itself (`assert_no_surface_styles` fails the build of the mock
if it does). The names, props and token reads here are the contract the Svelte/Astro components implement.
"""
import re
from html import escape

from mocklib import emblem
from ws_lib import CSS as BASE_CSS, page as _page, control_bar, seg, sel, crest_img, icon, QCOL, write, REGION, region_slug

PRIM_CSS = '''
/* ---- Page: one column, four layouts ---- */
.layout-split{display:grid;gap:var(--gap);grid-template-columns:minmax(0,1fr);align-items:start}
@media(min-width:1024px){.layout-split{grid-template-columns:minmax(0,1fr) 340px}}
@media(min-width:1760px){.layout-split{grid-template-columns:minmax(0,1fr) var(--rail-w)}}
.layout-split>.main,.layout-split>.rail{display:flex;flex-direction:column;gap:var(--section-gap);min-width:0}
.layout-split>.rail{gap:32px}
@media(min-width:1024px){.layout-split>.rail.sticky{position:sticky;top:76px}}
/* ---- PageHeader ---- */
.ph-wrap{padding-bottom:0}
.ph-title{display:flex;gap:16px;align-items:center;min-width:0}
/* below 900 the side is no box: the answer, then the actions, each a row of the head's grid (16 px apart); a hidden
   answer leaves no gap (web ui/PageHeader) */
@media(max-width:899px){.ph-side{display:contents}.ph-side>.ans{order:1}.ph-side>.ph-actions{order:2}}
@media(min-width:900px){.ph-side{display:flex;flex-direction:row;align-items:flex-end;gap:24px;min-width:0}.ph-side .ph-actions{justify-content:flex-end}.ph-side:has(>.ans)>.ph-actions{margin-bottom:-6px}.head .ans{max-width:26rem}}
.ph-title .t-title{overflow-wrap:anywhere}
/* the crumb (a way back) on the band above the title: it takes the row's top padding (web PageHeader .ui-page-header-crumb) */
.ph-crumb{display:flex;align-items:center;min-height:32px;padding-top:4px}.ph-crumb+.head{padding-top:0}
.ph-crumb a{display:inline-flex;align-items:center;min-height:32px;font-size:14px;color:var(--muted);text-decoration:none}
/* the notice line under the row, one style (web .ui-page-header-notice: 10 px under the row, the head's 14 px under it) */
.head:has(+.ph-notice){padding-bottom:0}.ph-notice{margin-top:10px;padding-bottom:14px}
/* a pending title: the skeleton holds the title row, the heading is for a screen reader only */
.ph-title-sk{display:block;width:280px;max-width:100%;height:30px}@media(max-width:767px){.ph-title-sk{height:25px}}
.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap}
.ph-meta{display:flex;gap:6px 16px;flex-wrap:wrap;margin-top:6px}
.ph-actions{display:flex;gap:10px;flex-wrap:wrap;align-items:center}
.ph-tabs{margin-top:8px}
/* ---- Section: a text heading over a hairline, no box ---- */
.sec-h{display:flex;align-items:baseline;gap:12px;padding-bottom:10px;border-bottom:1px solid var(--line-soft);margin-bottom:0}
.sec-h .act{margin-left:auto;font-weight:600;font-size:14px}
.sec-h+.plate{margin-top:12px}.sec-h+.rows{margin-top:0}
/* ---- Row: lead | body | trail ---- */
.rows{display:flex;flex-direction:column}
.row{display:grid;grid-template-columns:auto minmax(0,1fr) auto;column-gap:12px;align-items:center;padding:10px 16px;min-height:56px;border-bottom:1px solid var(--line-soft);position:relative}
.row:last-child{border-bottom:0}
.row.nolead{grid-template-columns:minmax(0,1fr) auto}
.row .b{min-width:0}.row .ttl{font:700 15px/20px var(--font-body);color:var(--strong)}
.row .sub{font-size:13px;line-height:18px;color:var(--muted)}
.row .trail{display:flex;gap:12px;align-items:center;justify-self:end}
.row .foot{grid-column:1/-1;display:flex;flex-wrap:wrap;gap:8px 12px;align-items:center;padding-top:8px}
a.row:hover,.row.hov{background:var(--hover)}
.row.cur{background:#e5b9551a;box-shadow:inset 3px 0 0 var(--gold)}
.row .rank{width:24px;text-align:right;font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--muted)}
.plate>.rows>.row{padding-inline:16px}
.rows.open .row{padding-inline:0}
/* ---- Figure and Meter ---- */
.fig{display:flex;flex-direction:column;gap:2px;min-width:0}
.fig .v{font:700 28px/32px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}
.fig .v small{font-size:13px;font-weight:600;color:var(--muted);margin-left:4px}
.fig.big .v{font-size:36px;line-height:40px}.fig.sm .v{font-size:18px;line-height:24px}
.fig.gold .v{color:var(--gold)}.fig.kill .v{color:var(--kill)}.fig.wipe .v{color:var(--wipe)}
.figs{display:grid;gap:20px 24px;grid-template-columns:repeat(var(--c,3),minmax(0,1fr))}
.meter{height:8px;border-radius:3px;background:var(--line-soft);overflow:hidden;min-width:60px}.meter>i{display:block;height:100%;border-radius:3px;background:var(--gold)}
.meter.kill>i{background:var(--kill)}.meter.muted>i{background:#5a5a52}
.meter.seg{display:flex;gap:2px;background:none}.meter.seg>i{border-radius:0}
/* ---- Tag, Field ---- */
/* web ui/Tag: 22 px, 0 8 px (site-alignment 1: "A state on a thing is a Tag: 22 px") */
.pill.tag{height:22px;padding:0 8px}
.tag.kill{color:var(--kill);border-color:#7fd48a66}.tag.muted{color:var(--muted);border-color:var(--line)}.tag.wipe{color:var(--wipe);border-color:#f0736b66}
.field{display:flex;align-items:center;gap:8px;height:40px;padding:0 12px;border:1px solid var(--line);border-radius:4px;background:#0a0d14;color:var(--muted);font-size:14px;min-width:0}
.field.grow{flex:1;min-width:200px}
.empty{padding:16px;color:var(--muted);font-size:14px}
.strip-tabs{display:flex;gap:4px;border-bottom:1px solid var(--line);overflow-x:auto}
.strip-tabs a{padding:14px 14px;font:600 14px/16px var(--font-body);color:var(--nav);border-bottom:2px solid transparent;white-space:nowrap}
.strip-tabs a.cur{color:var(--strong);border-bottom-color:var(--gold)}
.strip-tabs a .n{font-size:12px;color:var(--muted);margin-left:6px}
'''

_BANNED = re.compile(r'style="[^"]*(border|background|box-shadow|border-radius)', re.I)


def assert_no_surface_styles(html: str) -> None:
    bad = _BANNED.findall(html)
    if bad:
        raise AssertionError(f'a page drew its own surface ({len(bad)} inline border/background styles); use a primitive')


# ----------------------------------------------------------------- primitives
def Button(label: str, primary: bool = False, small: bool = False) -> str:
    return f'<a class="btn{" primary" if primary else ""}{" sm" if small else ""}" href="#">{label}</a>'


def Link(label: str) -> str:
    return f'<a class="link" href="#">{label}</a>'


def Tag(label: str, tone: str = 'gold') -> str:
    return f'<span class="pill tag {tone}">{label}</span>'


def Figure(value: str, label: str = '', unit: str = '', tone: str = '', size: str = '') -> str:
    u = f'<small>{unit}</small>' if unit else ''
    lab = f'<span class="t-label">{label}</span>' if label else ''
    return f'<div class="fig {size} {tone}">{lab}<span class="v">{value}{u}</span></div>'


def Figures(items: list[tuple]) -> str:
    return f'<div class="figs" style="--c:{len(items) if len(items) != 3 else 3}">' + ''.join(Figure(*i) for i in items) + '</div>'


def Meter(frac: float, tone: str = '') -> str:
    return f'<div class="meter {tone}"><i style="width:{max(0, min(1, frac)) * 100:.0f}%"></i></div>'


def Field(label: str, grow: bool = False) -> str:
    return f'<span class="field{" grow" if grow else ""}">{label}</span>'


def Row(title: str, sub: str = '', lead: str = '', trail: str = '', title_color: str = '', rank: str = '', link: bool = True, selected: bool = False, foot: str = '', slim: bool = False) -> str:
    cls = 'row' + ('' if lead or rank else ' nolead') + (' cur' if selected else '') + (' slim' if slim else '')
    tag = 'a' if link else 'div'
    href = ' href="#"' if link else ''
    color = f' style="color:{title_color}"' if title_color else ''
    lead_html = (f'<span class="rank">{rank}</span>' if rank else '') + lead
    lead_cell = f'<span style="display:flex;gap:12px;align-items:center">{lead_html}</span>' if (lead or rank) else ''
    return (f'<{tag} class="{cls}"{href}>{lead_cell}<div class="b"><div class="ttl"{color}>{title}</div>' + (f'<div class="sub">{sub}</div>' if sub else '') + f'</div>'
            + (f'<div class="trail">{trail}</div>' if trail else '<span></span>')
            + (f'<div class="foot">{foot}</div>' if foot else '') + f'</{tag}>')


def Rows(rows: list[str], open_: bool = False) -> str:
    return f'<div class="rows{" open" if open_ else ""}">' + ''.join(rows) + '</div>'


def Plate(inner: str, label: str = '') -> str:
    """label: the Plate's accessible name (role region), as the built Plate's `label` prop."""
    name = f' role="region" aria-label="{label}"' if label else ''
    return f'<div class="plate" data-region="{REGION["plate"]}"{name}>{inner}</div>'


def Section(title: str, body: str, action: str = '', meta: str = '', boxed: bool = False) -> str:
    a = f'<span class="act">{action}</span>' if action else ''
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    # round 5 (web ui/Section `boxed`): the body on the panel 12 px under the heading, unless it brings its own panel
    inner = f'<div class="sbox">{body}</div>' if boxed else body
    return f'<section data-region="{region_slug(title)}"><div class="sec-h"><h2 class="t-h3">{title}</h2>{m}{a}</div>{inner}</section>'


def TabStrip(tabs: list[str], active: str, counts: dict | None = None) -> str:
    return '<nav class="strip-tabs ph-tabs">' + ''.join(f'<a href="#" class="{"cur" if t == active else ""}">{t}{f"<span class=\"n\">{counts[t]}</span>" if counts and t in counts else ""}</a>' for t in tabs) + '</nav>'


def PageHeader(title: str, title_color: str = '', crest: str = '', meta: list[str] | None = None, answer: str = '', actions: str = '', tabs: str = '',
               crumb: str = '', notice: str = '', pending: bool = False, second_head: bool = False) -> str:
    """web ui/PageHeader: `crumb` a way back on the band above the title, `notice` one line under the row, `pending` a title
    not known yet (a skeleton holds the row, the h1 is sr-only). `second_head`: a pending head beside the page's own
    server-rendered head (home's signed-in skeleton), its title a sr-only paragraph, never a second h1."""
    m = f'<div class="ph-meta" data-region="{REGION["head_meta"]}">{"".join(f"<span class=\"t-small\">{x}</span>" for x in meta)}</div>' if meta else ''
    color = f' style="color:{title_color}"' if title_color else ''
    ans = f'<div class="ans" data-region="{REGION["head_answer"]}">{answer}</div>' if answer else ''
    acts = f'<div class="ph-actions">{actions}</div>' if actions else ''
    right = f'<div class="ph-side">{acts}{ans}</div>' if (ans or acts) else ''
    # round 5: the title band (web ui/PageHeader): the painted band, full bleed, the head back in the column
    sr_tag = 'p' if second_head else 'h1'
    h1 = (f'<span class="skel ph-title-sk" aria-hidden="true"></span><{sr_tag} class="sr-only">{title}</{sr_tag}>' if pending
          else f'<h1 class="t-title"{color}>{title}</h1>')
    cr = f'<div class="ph-crumb t-small">{crumb}</div>' if crumb else ''
    nt = f'<div class="ph-notice t-small">{notice}</div>' if notice else ''
    busy = ' aria-busy="true"' if pending else ''
    return (f'<div class="band-r5 full-bleed"><div class="pc ph-wrap"><div data-region="{REGION["head"]}"{busy}>{cr}<div class="head"><div class="ph-title" data-region="{REGION["head_title"]}">{crest}<div>{h1}{m}</div></div>{right}</div>{nt}{tabs}</div></div></div>')


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False) -> str:
    if layout == 'ledger':
        return f'<div class="pc" style="padding-top:20px">{main}</div>'
    if layout == 'split':
        return f'<div class="pc" style="padding-top:28px"><div class="layout-split"><div class="main">{main}</div><aside class="rail{" sticky" if sticky_rail else ""}">{rail}</aside></div></div>'
    raise ValueError(layout)


def build_page(title: str, body: str, nav: str = '', who: tuple | None = None, note: str = '') -> str:
    assert_no_surface_styles(body)
    return _page(title, body + f'<style>{PRIM_CSS}</style>', nav, who, note)


PRIM_CSS += '''
.notice{display:flex;gap:6px 14px;flex-wrap:wrap;align-items:baseline;padding:14px 0 4px;font-size:15px;line-height:24px}.notice.gold{color:var(--gold)}.notice .link{font-weight:600}
.row .trail{flex-wrap:wrap;justify-content:flex-end}@media(max-width:600px){.row,.row.nolead{grid-template-columns:auto minmax(0,1fr)}.row.nolead{grid-template-columns:minmax(0,1fr)}.row .trail{grid-column:1/-1;justify-self:stretch;justify-content:flex-start;margin-top:8px}.figs{grid-template-columns:repeat(2,minmax(0,1fr))}}
@media(max-width:767px){.strip-tabs{flex-wrap:wrap;overflow:visible;gap:0 2px}.strip-tabs a{padding:12px 12px}}
.formrow{display:flex;gap:12px;flex-wrap:wrap;align-items:flex-end}
.formrow .f{display:flex;flex-direction:column;gap:6px;min-width:0}.formrow .f.grow{flex:1;min-width:220px}
.formrow .sel{height:40px}.formrow .field.grow{flex:none;width:100%}
@media(max-width:767px){.formrow .f,.formrow .btn{width:100%}}
'''


def FormRow(fields: list[tuple[str, str, bool]], submit: str) -> str:
    """Labelled Fields and Selects in one wrapping row with one submit Button (a form is not a filter: it is not a ControlBar)."""
    f = ''.join(f'<label class="f{" grow" if grow else ""}"><span class="t-label">{lab}</span>{html}</label>' for lab, html, grow in fields)
    return f'<div class="formrow">{f}{submit}</div>'


def Select(label: str) -> str:
    return sel(label)


def Emblem(faction: str, size: int = 56) -> str:
    """A guild's or character's faction mark in the one circular crest frame (FactionCrest)."""
    return f'<img class="crest" src="{emblem(faction)}" alt="" style="width:{size}px;height:{size}px;--c:{"#6fb1ff" if faction == "alliance" else "#ff6b5c"}">'


def Notice(text: str, action: str = '', tone: str = 'gold') -> str:
    """One sentence under the header (a standing line, a state): text, optional ActionLink, no box."""
    return f'<div class="pc"><p class="notice {tone}">{text}{action}</p></div>'


# ------------------------------------------------ primitives added by the guild persona round (reusable, not guild-only)
PRIM_CSS += '''
.gridm{display:grid;grid-template-columns:minmax(120px,1.5fr) repeat(var(--n),minmax(48px,1fr));align-items:center}
.gridm .h{font:700 12px/16px var(--font-body);letter-spacing:.1em;text-transform:uppercase;color:var(--muted);padding:8px 6px;border-bottom:1px solid var(--line-soft)}
.gridm .l{padding:10px 6px;border-bottom:1px solid var(--line-soft);min-width:0}.gridm .l .ttl{font:700 14px/18px var(--font-body);color:var(--strong)}.gridm .l .sub{font-size:12px;line-height:16px;color:var(--muted)}
.gridm .c{padding:10px 6px;border-bottom:1px solid var(--line-soft);text-align:center;font:700 14px/18px var(--font-body);font-variant-numeric:tabular-nums}
.gridm .c.ok{color:var(--kill)}.gridm .c.warn{color:var(--gold)}.gridm .c.miss{color:var(--wipe)}.gridm .c.none{color:var(--muted);font-weight:400}
.tally{display:grid;gap:12px 16px;grid-template-columns:repeat(auto-fill,minmax(150px,1fr))}
.tally .t{display:flex;flex-direction:column;gap:6px;min-width:0}.tally .t .top{display:flex;align-items:center;gap:8px}.tally .t .top b{font:700 15px/20px var(--font-body);color:var(--strong)}.tally .t .top .n{margin-left:auto;font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums}
.tally .t .n.none{color:var(--muted);font-weight:400}.tally .t .n.ok{color:var(--kill)}.tally .t .n.miss{color:var(--wipe)}.tally .t .n.warn{color:var(--gold)}
.ck{width:22px;height:22px;border-radius:50%;display:inline-flex;align-items:center;justify-content:center;font:700 13px/22px var(--font-body);color:#07090d;flex:none}
.ck.ok{background:var(--kill)}.ck.warn{background:var(--gold)}.ck.miss{background:var(--wipe)}.ck.gap{background:var(--muted)}
.gap-tag{font-size:12px;color:var(--muted)}
.seg.sm a{height:32px;padding:0 10px;font-size:13px}
.es-line{display:flex;min-height:44px;flex-wrap:wrap;align-items:center;gap:12px;font-size:14px}.es-line .msg{color:var(--muted)}
'''


def Segmented(items: list[tuple[str, bool]], small: bool = False) -> str:
    """An inline choice of one among few (sign-up In / Late / Bench). Same look as the ControlBar's segmented control."""
    return seg([(a, on, None, '') for a, on in items], 'sm' if small else '')


def EmptyState(message: str, action: str | None = None) -> str:
    """The shared empty line (ui/EmptyState.svelte, component audit cluster 10): a region whose data loaded with nothing in it,
    the muted sentence at 44 px and at most one gold link that leads to what would fill it."""
    act = Link(f'{action} &rsaquo;') if action else ''
    return f'<div class="es-line"><span class="msg">{message}</span>{act}</div>'


def Checklist(items: list[tuple]) -> str:
    """Private pre-raid checks as the built Checklist: a Row per check, its mark an image, its fix a link in the trail.
    items: (state ok|warn|miss, title, sub, fix_label|None)."""
    rows = []
    for st, title, sub, fix in items:
        mark = {'ok': '&#10003;', 'warn': '!', 'miss': '&times;'}[st]
        word = {'ok': 'Done', 'warn': 'To do', 'miss': 'Missing'}[st]
        rows.append(Row(title, sub, lead=f'<span class="ck {st}" role="img" aria-label="{word}">{mark}</span>', trail=Link(fix) if fix else '', link=False))
    return Rows(rows, open_=True)

def Votes(yes: int, no: int, total: int) -> str:
    return f'<div style="display:flex;gap:6px;align-items:center">{Tag(f"{yes} yes", "kill")}{Tag(f"{no} no", "wipe" if no else "muted")}<span class="t-small">of {total} officers</span></div>'


def Gap(label: str) -> str:
    """A module whose data source is not confirmed: marked in the mock, named in the spec."""
    return f'<span class="gap-tag">&#9888; {label}</span>'


# ------------------------------------------------ primitives added by the tiers round (2026-11-06; design/specs/2026-11-06-tiers.md section 4)
# RankList / RankGroup / RankRow / ColumnHeads: a ranked list on the page's one Plate. RankRow is a Row preset (as ItemRow is),
# reused next by the log rankings and the guild Top guilds list; RankGroup is the letter column (and later BiS "Where to go"
# groups and the raid planner's role groups). Meter gains a class tone; Skeleton is the shipped `Skeleton`.
PRIM_CSS += '''
.rl{--grp-w:0px;--rr-name:minmax(180px,220px);--rr-bar:minmax(120px,1fr);--rr-fig:112px;--rr-gap:88px;--rr-links:206px;--n:1}
.rl.links3{--rr-links:156px}
/* tank figures share the list's 112 px columns: at 124 the built heads clip at 1280 (web e2e tiers.spec.ts, "every column head is one short line") */
@media(min-width:1100px){.rl.grouped{--grp-w:184px}}
@media(min-width:1760px){.rl{--rr-name:minmax(240px,1fr);--rr-bar:560px}}
.rl .cols{display:grid;column-gap:16px;grid-template-columns:28px 36px var(--rr-name) var(--rr-bar) repeat(var(--n),var(--rr-fig)) var(--rr-gap) var(--rr-links)}
.rr-heads{align-items:end;padding:12px 16px 10px calc(16px + var(--grp-w));border-bottom:1px solid var(--line-soft)}
.rr-heads .r{text-align:right}
.rg{display:grid;grid-template-columns:var(--grp-w) minmax(0,1fr);border-top:1px solid var(--line)}
.rg:first-child{border-top:0}
.rg-l{display:grid;grid-template-columns:22px minmax(0,1fr);grid-template-areas:"lt bd" "lt ct";align-content:start;align-items:center;column-gap:10px;padding:10px 12px 10px 14px;border-left:3px solid var(--t);border-right:1px solid var(--line-soft)}
.rg-l .lt{grid-area:lt;font:700 32px/36px var(--font-body);color:var(--t)}
.rg-l .ct{grid-area:ct;font-size:13px;line-height:18px;color:var(--muted);font-variant-numeric:tabular-nums}
.rg-l .bd{grid-area:bd;font-size:13px;line-height:18px;color:var(--text)}
.rg-rows{min-width:0}
.rr{position:relative;align-items:center;min-height:var(--row-h);padding:8px 16px;border-bottom:1px solid var(--line-soft)}
.rg-rows>.rr:last-child,.rl>.rr:last-child,.rl>[role=list]>.rr:last-child{border-bottom:0}
.rr::before{content:"";position:absolute;top:0;bottom:0;left:calc(-1 * var(--grp-w));width:3px;background:transparent;z-index:1}
.rr:hover,.rr.hov{background:var(--hover)}.rr:hover::before,.rr.hov::before{background:var(--c)}.rr:hover .meter>i,.rr.hov .meter>i{filter:brightness(1.18)}
.rr.cur{background:#e5b9551a}.rr.cur::before{background:var(--gold)}
.rr .rk{text-align:right;font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--muted)}
.rr .ld{display:flex}
.rr .nm{min-width:0;display:flex;flex-direction:column}
.rr .nm .ttl{display:flex;gap:8px;align-items:center;flex-wrap:wrap}
.rr .nm .ttl a{font:700 16px/20px var(--font-body);color:var(--c)}.rr .nm .ttl a:hover{text-decoration:underline;text-underline-offset:4px}
.rr .nm .sub{font-size:13px;line-height:18px;color:var(--muted);white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.rr .nm .low{font-size:13px;line-height:18px;color:var(--muted);text-decoration:underline dotted;text-underline-offset:3px;align-self:flex-start}
.rr .br{display:flex;flex-direction:column;gap:4px;min-width:0}.rr .br .tie{font:700 12px/16px var(--font-body);color:var(--strong)}
.rr .x{display:contents}
.rr .fg{text-align:right;font:700 16px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong);white-space:nowrap}
.rr .fg small{font-size:12px;font-weight:600;color:var(--muted);margin-left:4px}.rr .fg.best{color:var(--kill)}.rr .fg .lbl{display:none}
.rr .gp{text-align:right;font:600 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--text);white-space:nowrap}
.rr .gp.top{color:var(--strong);font-weight:700}
.rr .lk{display:flex;justify-content:flex-end;gap:16px}
/* the shipped RankRow (ui/rank-list.css): the token row height (tokens.css --row-h 58) and 44 px link targets */
.rl .rr{min-height:58px}.rl .rr-heads .t-label{white-space:nowrap;overflow:hidden;text-overflow:ellipsis}
.rr .lk a{display:inline-flex;align-items:center;min-height:44px;font:600 14px/20px var(--font-body);color:var(--nav)}.rr .lk a.p{color:var(--gold)}
.rr .lk a:hover,.rr .lk a.hov{color:var(--gold-hover)}.rr .lk a.fv{outline:2px solid var(--gold);outline-offset:2px;border-radius:3px}
.tag.you{margin-left:2px}
.meter.cls>i{background:var(--c)}
.skel{display:block;border-radius:4px;background:linear-gradient(90deg,#131824,#1a2030,#131824)}
@media(max-width:1099px){.rg{grid-template-columns:minmax(0,1fr)}.rg-l{display:flex;align-items:baseline;gap:12px;padding:8px 14px 8px 12px;border-right:0;border-bottom:1px solid var(--line-soft)}.rg-l .lt{font-size:22px;line-height:24px}.rg-l .bd{order:1}.rg-l .ct{order:2;margin-left:auto}}
@media(max-width:1023px){.rl .rr-heads{display:none}
 .rl .rr.cols{grid-template-columns:22px 32px minmax(0,1fr) auto;grid-template-areas:"rk ld nm f1" "rk ld nm gp" ". . xx xx" ". . br br" ". . lk lk";column-gap:12px;row-gap:0;padding:10px 14px 2px}
 .rr .rk{grid-area:rk}.rr .ld{grid-area:ld}.rr .nm{grid-area:nm}.rr .br{grid-area:br}.rr .gp{grid-area:gp}.rr .lk{grid-area:lk}.rr .fg.f1{grid-area:f1}.rr .x{grid-area:xx;display:flex;gap:0 16px;flex-wrap:wrap}.rr .x .fg{text-align:left;font-size:14px}.rr .x .fg .lbl{display:inline;font-weight:400;color:var(--muted);margin-right:6px}
 .rr .br{padding-top:8px}.rr .gp{align-self:start}.rr .fg.f1{align-self:end}.rr .x:not(:empty){padding-top:6px}.rr .lk{justify-content:flex-start;gap:20px}.rr .lk a{min-height:44px}}
'''


def MeterClass(frac: float, color: str) -> str:
    """Meter with tone 'class': the fill in the class colour, set by the primitive from a prop (Details! bars)."""
    return f'<div class="meter cls" style="--c:{color}"><i style="width:{max(0, min(1, frac)) * 100:.1f}%"></i></div>'


def Skeleton(w: int, h: int, round_: bool = False) -> str:
    """The shipped `Skeleton`: a block at the final geometry, so nothing shifts when the content lands."""
    return f'<span class="skel{" round" if round_ else ""}" style="width:{w}px;height:{h}px"></span>'


PRIM_CSS += '.skel.round{border-radius:999px}\n'


def ColumnHeads(cells: list[tuple[str, bool]]) -> str:
    """The Plate's one column-head row (12 px caps labels), on the RankList's own grid. cells: (label, right_aligned)."""
    return '<div class="rr-heads cols" aria-hidden="true">' + ''.join(f'<span class="t-label{" r" if r else ""}">{t}</span>' for t, r in cells) + '</div>'


def RankRow(rank: str, crest: str, color: str, name: str, sub: str, frac: float, figures: list[tuple], gap: str,
            links: list[tuple[str, bool]], gap_top: bool = False, tie: str = '', low: str = '', you: str = '',
            selected: bool = False, state: str = '') -> str:
    """One ranked row. figures: (value, unit, phone_label, best); the first is the sorted figure. links: (label, primary)."""
    def fig(i, f):
        v, u, lbl, best = f
        return f'<span class="fg{" f1" if i == 0 else ""}{" best" if best else ""}"><span class="lbl">{lbl}</span>{v}<small>{u}</small></span>'
    first = fig(0, figures[0])
    extra = ''.join(fig(i, f) for i, f in enumerate(figures[1:], 1))
    # the built RankRow (ui/RankRow.svelte): the tie line sits under the bar ("≈ tie with Arms")
    tie_html = f'<span class="tie">{tie}</span>' if tie else ''
    low_html = f'<span class="low">{low}</span>' if low else ''
    you_html = f'<span class="pill tag you">{you}</span>' if you else ''
    lk = ''.join(f'<a href="#" class="{"p" if p else ""}{" " + state.split(":")[1] if state.startswith("link:") and i == 0 else ""}">{t}</a>' for i, (t, p) in enumerate(links))
    cls = 'rr cols' + (' cur' if selected else '') + (' hov' if state == 'hover' else '')
    return (f'<div class="{cls}" role="listitem" style="--c:{color}"><span class="rk">{rank}</span><span class="ld">{crest}</span>'
            f'<div class="nm"><span class="ttl"><a href="#">{name}</a>{you_html}</span><span class="sub">{sub}</span>{low_html}</div>'
            f'<span class="br">{MeterClass(frac, color)}{tie_html}</span>{first}<span class="x">{extra}</span>'
            f'<span class="gp{" top" if gap_top else ""}">{gap}</span><span class="lk">{lk}</span></div>')


def RankGroup(letter: str, color: str, band: str, count: str, rows: list[str]) -> str:
    """A letter group: the 148 px label column (letter, count, band) with the tier rail at x=0; a header strip below 1100."""
    # the built RankGroup's semantics: a list item of the ranked list, holding its own list named "Tier {letter}"
    return (f'<div class="rg" role="listitem" style="--t:{color}"><div class="rg-l"><span class="lt">{letter}</span><span class="ct">{count}</span>'
            f'<span class="bd">{band}</span></div><div class="rg-rows" role="list" aria-label="Tier {letter}">{"".join(rows)}</div></div>')


def RankList(heads: str, body: str, n_figs: int = 1, grouped: bool = False, links3: bool = False, tank: bool = False) -> str:
    cls = 'rl' + (' grouped' if grouped else '') + (' links3' if links3 else '') + (' tank' if tank else '')
    # the built RankList's semantics: the heads are hidden from assistive tech, the rows an ordered list
    return f'<div class="{cls}" style="--n:{n_figs}">{heads}<div role="list">{body}</div></div>'


# ------------------------------------------------ primitives added by the sim round (2026-11-06; design/specs/2026-11-06-sim.md section 4)
# Additive only. Figure gains size 'xl' (the sim's 56 px answer). StatTable is a numeric table on the Plate (lead icon, name and
# sub line, an optional Meter column, right-aligned figure columns), the shape the shipped report/AbilityTable.svelte adopts and
# the logs report reuses; in open mode (no Plate) it is the resource and aura lists. Facts is a key/value list. IconList is a
# wrapping list of icon + name with no chip border (the old 30 px buff chips drew a box each). The shared Paperdoll (below) is the worn doll in
# `worn` mode (no gain column: the sim shows inputs only), plated or open. Histogram, Lanes and Trees mirror the shipped chart
# components (DpsDistribution, SimTimeline, SimTalentsCard over GuideTreeColumns): SVG and icons, no surfaces.
PRIM_CSS += '''
.fig.xl .v{font-size:56px;line-height:60px;letter-spacing:-.01em}.fig.xl .v small{font-size:18px;margin-left:8px}
@media(max-width:767px){.fig.xl .v{font-size:44px;line-height:48px}}
.fig .cap{font-size:13px;line-height:20px;color:var(--muted)}.fig .cap b{color:var(--text);font-weight:600}.fig .cap.kill{color:var(--kill)}.fig .cap.wipe{color:var(--wipe)}
/* StatTable tracks (fixed 2026-11-06, the buffs grid overlap): the name and the bar share the free space 2 : 1 and may shrink to 0,
   so no fixed minimum can push a column past the table's own edge; a figure column is never narrower than its widest figure
   (minmax(max-content, n) resolves to max-content when that is wider), so a figure never spills into the next column. The table
   clips nothing: it is as wide as its container, at every container width. */
.st{--st-name:minmax(0,2fr);--st-bar:minmax(0,1fr);--st-fig:minmax(max-content,62px);--n:5;min-width:0;max-width:100%}
@media(min-width:1760px){.st{--st-name:minmax(0,1fr);--st-bar:minmax(0,280px);--st-fig:minmax(max-content,80px)}}
.st .br{min-width:0}.st .br .meter{max-width:280px;min-width:0}.st .nm b{overflow-wrap:anywhere}
.st-h .t-label{letter-spacing:.06em;min-width:0}
.st .stc{display:grid;column-gap:12px;align-items:center;grid-template-columns:28px var(--st-name) var(--st-bar) repeat(var(--n),var(--st-fig))}
.st.nobar .stc{grid-template-columns:28px var(--st-name) repeat(var(--n),var(--st-fig))}
.st-h{padding:12px 16px 10px;border-bottom:1px solid var(--line-soft)}.st-h .r{text-align:right}.st-h .on{color:var(--strong)}.st-h .on::after{content:" \\25BE";color:var(--gold)}
.st-r{min-height:56px;padding:8px 16px;border-bottom:1px solid var(--line-soft)}.st-r:last-child{border-bottom:0}.st-r:hover,.st-r.hov{background:var(--hover)}
.st-r .nm{min-width:0}.st-r .nm b{font:700 15px/20px var(--font-body);color:var(--strong)}.st-r .nm .rk{font-size:13px;color:var(--muted);font-weight:400;margin-left:4px}
.st-r .nm .wpn{font-weight:600;font-size:14px}.st-r .nm .sub{font-size:13px;line-height:18px;color:var(--muted)}.st-r .nm .sub span{white-space:nowrap}.st-r .nm .sub .m{display:none}
.st-r .f{text-align:right;font:600 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--text);white-space:nowrap}.st-r .f.p{font-weight:700;color:var(--strong)}.st-r .f.d{color:var(--muted);font-weight:400}
.st-r .f.p .m{display:none}
.st-r.tot{min-height:44px}.st-r.tot .f,.st-r.tot .nm{font-size:13px;color:var(--muted);font-weight:400}
.st.open .st-r,.st.open .st-h{padding-inline:0}
@media(max-width:1023px){.st .st-h{display:none}.st .stc,.st.nobar .stc{grid-template-columns:28px minmax(0,1fr) auto}.st-r .br,.st-r .f:not(.p){display:none}.st-r .f.p .m{display:block;font-size:13px;font-weight:400;color:var(--muted)}.st-r .nm .sub .m{display:inline}.st.open .st-r .br{display:block;grid-column:2/-1;padding-top:6px}.st.open .st-r .f.p{grid-row:1;grid-column:3}.st.open .stc{row-gap:0}}
.kv{display:grid;grid-template-columns:minmax(0,auto) minmax(0,1fr);margin:0}
.kv dt,.kv dd{padding:10px 0;border-bottom:1px solid var(--line-soft);font-size:14px;line-height:20px;margin:0}
.kv dt{color:var(--muted);padding-right:16px}.kv dd{color:var(--strong);font-weight:600;text-align:right;font-variant-numeric:tabular-nums}.kv dd small{display:block;font-weight:400;color:var(--muted);font-size:13px}
.il{display:flex;flex-wrap:wrap;gap:8px 18px;padding-top:10px}.il span{display:inline-flex;align-items:center;gap:8px;font-size:14px;line-height:20px;color:var(--text);min-height:28px}.il span.dim{color:var(--muted)}.il span.dim img,.il span.dim .ic{opacity:.62}
.il-l{margin-top:16px}
.hist svg{display:block;width:100%;height:168px}.hist .ax{display:flex;justify-content:space-between;font-size:13px;color:var(--muted);padding-top:6px;font-variant-numeric:tabular-nums}
.lanes{overflow-x:auto}.lanes svg{display:block}.lanes .lg{display:flex;flex-wrap:wrap;gap:6px 18px;font-size:13px;color:var(--muted);padding:10px 0}.lanes .lg i{display:inline-block;width:14px;height:8px;margin-right:6px;border-radius:2px;vertical-align:middle}
.trees{display:grid;gap:var(--gap);grid-template-columns:repeat(3,minmax(0,1fr))}.trees .tr h3{display:flex;gap:8px;align-items:baseline;font:700 15px/20px var(--font-body);color:var(--strong);padding-bottom:8px}.trees .tr h3 b{color:var(--gold)}
.trees .tg{display:grid;grid-template-columns:repeat(4,36px);gap:8px}.trees .tc{position:relative;width:36px;height:36px}.trees .tc img,.trees .tc .ic{width:36px!important;height:36px!important}.trees .tc.off img{filter:grayscale(1) brightness(.45)}
.trees .tc.full img{box-shadow:0 0 0 2px var(--gold)}.trees .tc.part img{box-shadow:0 0 0 2px var(--kill)}.trees .tc b{position:absolute;right:-4px;bottom:-6px;font:700 11px/14px var(--font-body);color:var(--strong);text-shadow:0 0 3px #000,0 0 3px #000}
@media(max-width:767px){.trees{grid-template-columns:minmax(0,1fr)}}
'''


def StatTable(heads: list[tuple[str, bool]], rows: list[str], n_figs: int, bar: bool = True, open_: bool = False, sort: str = '',
              sortable: tuple[str, ...] = ()) -> str:
    """A numeric table. heads: (label, right_aligned) for icon, name, [bar], figures; `sort` names the sorted head and
    `sortable` the heads that are sort buttons. An ARIA table (as the built ui/StatTable): the heads are columnheaders, a
    sortable one holds a button and carries aria-sort; the icon and bar cells are decorative (aria-hidden). A table whose
    heads are all empty (the open Rage and aura lists) draws no head row."""
    def head(i, t, r):
        if i == 0 or not t:
            return f'<span class="t-label" aria-hidden="true">{t}</span>'
        srt = f' aria-sort="{"descending" if t == sort else "none"}"' if t in sortable else ''
        inner = f'<button type="button" class="st-sort">{t}</button>' if t in sortable else t
        return f'<span class="t-label{" r" if r else ""}{" on" if t == sort and t else ""}" role="columnheader"{srt}>{inner}</span>'
    h = ''.join(head(i, t, r) for i, (t, r) in enumerate(heads))
    head_row = f'<div class="st-h stc" role="row">{h}</div>' if any(t for t, _ in heads) else ''
    cls = 'st' + ('' if bar else ' nobar') + (' open' if open_ else '')
    return f'<div class="{cls}" style="--n:{n_figs}" data-n="{n_figs}" role="table">{head_row}{"".join(rows)}</div>'


def StatRow(icon_html: str, name: str, sub: str, figures: list[str], meter: str = '', primary_phone: str = '', state: str = '', total: bool = False,
            bar: bool = True) -> str:
    """One table row. figures: html of each figure cell; the first is the sorted, primary figure (stays on phone with
    `primary_phone` under it). `sub` may carry `<span class="m">` parts shown only on phone. `bar`: the table has a bar
    column (the Total row keeps its empty cell so its figures stay under their heads)."""
    figs = ''.join(f'<span class="f{" p" if i == 0 else ""}{" d" if f in ("&ndash;", "") else ""}" role="cell">{f}{f"<span class=m>{primary_phone}</span>" if i == 0 and primary_phone else ""}</span>' for i, f in enumerate(figures))
    br = f'<span class="br" aria-hidden="true">{meter}</span>' if (meter or bar) else ''
    return (f'<div class="st-r stc{" hov" if state == "hover" else ""}{" tot" if total else ""}" role="row"><span aria-hidden="true">{icon_html}</span>'
            f'<div class="nm" role="cell">{name}' + (f'<div class="sub">{sub}</div>' if sub else '') + f'</div>{br}{figs}</div>')


def Facts(items: list[tuple[str, str]]) -> str:
    return '<dl class="kv">' + ''.join(f'<dt>{k}</dt><dd>{v}</dd>' for k, v in items) + '</dl>'


def IconList(items: list[tuple[str, str, bool]], label: str = '') -> str:
    """items: (icon_html, name, dim). A wrapping list; `dim` items are the ones with no effect on this spec."""
    lab = f'<h3 class="t-label il-l">{label}</h3>' if label else ''
    return lab + '<div class="il">' + ''.join(f'<span class="{"dim" if d else ""}">{ic}{n}</span>' for ic, n, d in items) + '</div>'


def Histogram(counts: list[int], start: float, bin_dps: float, mean: float, lo: float, hi: float, label: str, axis: tuple[str, str, str]) -> str:
    """DpsDistribution: at most 40 bars, the 16th to 84th percentile in gold, a mean line; real engine bins."""
    n, top = len(counts), max(counts) or 1
    w = 100 / n
    bars = ''
    for i, c in enumerate(counts):
        mid = start + (i + .5) * bin_dps
        fill = '#e5b955b3' if lo <= mid <= hi else '#262e40'
        h = 150 * c / top
        bars += f'<rect x="{i * w + .15:.2f}" y="{168 - h:.1f}" width="{w - .3:.2f}" height="{h:.1f}" fill="{fill}"/>'
    mx = (mean - start) / (n * bin_dps) * 100
    bars += f'<rect x="{mx - .15:.2f}" y="10" width=".3" height="158" fill="#f5d27a"/>'
    ax = ''.join(f'<span>{a}</span>' for a in axis)
    return f'<div class="hist"><svg viewBox="0 0 100 168" preserveAspectRatio="none" role="img" aria-label="{label}">{bars}</svg><div class="ax">{ax}</div></div>'


def LanesLegend(legend: list[tuple]) -> str:
    """The lanes' legend, a list (the shipped SimTimeline's `ul`): (colour, label) or (colour, label, width, height) for a
    swatch drawn at the mark's own size."""
    def item(c, t, *size):
        dim = f';width:{size[0]}px;height:{size[1]}px' if size else ''
        return f'<li><i style="--k:{c}{dim}" class="k"></i>{t}</li>'
    return '<ul class="lg">' + ''.join(item(*x) for x in legend) + '</ul>'


def Lanes(svg: str, legend: list[tuple[str, str]] | None = None) -> str:
    """The caller's plot carries the shipped plot's role. No legend: the caller places LanesLegend itself (the sim's tool line)."""
    return f'<div class="lanes">{LanesLegend(legend) if legend else ""}{svg}</div>'


PRIM_CSS += '.lanes .lg i.k{background:var(--k)}\n'


def Trees(trees: list[tuple[str, int, list[tuple]]]) -> str:
    """Read-only talent trees: (name, points, cells (tier, column, icon_html, rank, max)). The planner's tooltips on hover.
    The built TreeGrid's semantics: a grid per tree, a row per tier, a gridcell per column, a group per talent."""
    out = ''
    for name, pts, cells in trees:
        grid = ''
        for tier in range(7):
            row = ''
            for col in range(4):
                c = next((x for x in cells if x[0] == tier and x[1] == col), None)
                if c is None:
                    row += '<span role="gridcell"></span>'
                    continue
                _, _, ic, rk, mx = c
                st = 'full' if rk == mx else ('part' if rk else 'off')
                row += f'<span role="gridcell"><span class="tc {st}" role="group">{ic}' + (f'<b>{rk}/{mx}</b>' if rk else '') + '</span></span>'
            grid += f'<div role="row">{row}</div>'
        out += f'<div class="tr"><h3>{name} <b class="num">{pts}</b></h3><div class="tg" role="grid" aria-label="{name} talents">{grid}</div></div>'
    return f'<div class="trees">{out}</div>'
PRIM_CSS += '''
.tl{display:grid;grid-template-columns:200px minmax(0,1fr);min-width:0}
.tl .lbl{display:flex;flex-direction:column}.tl .lbl>*{display:flex;align-items:center;gap:8px;font-size:14px;line-height:16px;color:var(--text);white-space:nowrap;overflow:hidden}
.tl .lbl .g{font:600 13px/16px var(--font-body);color:var(--muted)}.tl .lbl .res{flex-direction:column;align-items:flex-start;justify-content:center;gap:0}.tl .lbl .res small{font-size:13px;color:var(--muted)}
.tl .plot{position:relative;min-width:0}.tl .plot svg{display:block;width:100%}
.tl .axis{position:relative;height:30px}.tl .axis span{position:absolute;top:6px;transform:translateX(-50%);font:600 13px/16px var(--font-body);color:var(--muted);font-variant-numeric:tabular-nums}.tl .axis span.pull{color:var(--gold)}
.tl .wash{position:absolute;top:34px;font:700 11px/14px var(--font-body);letter-spacing:.14em;color:#d66e28;padding-left:6px}
.tl .nmi{display:none}
@media(max-width:767px){.lanes .lg{padding-inline:var(--page-inset)}.tl .lbl .g{overflow:visible;font-size:12px}}
@media(max-width:767px){.lanes{margin-inline:calc(-1 * var(--page-inset))}.tl{grid-template-columns:44px minmax(696px,1fr);width:max-content;min-width:100%}.tl .lbl{position:sticky;left:0;z-index:2;background:var(--bg);padding-left:var(--page-inset)}.tl .lbl .t{display:none}.tl .lbl .res small{display:none}}
'''


# ------------------------------------------------ Sheet (added by the sim round, ruling S7, 2026-11-06)
# The shipped SettingsSheet as a primitive: an overlay panel over a scrim. Desktop (768 and up): a 480 px panel on the right,
# full height, the page dimmed behind it. Phone: a bottom sheet, 88vh at most, the top 16 px rounded. Its head (title, close)
# and foot (actions) stay put; the body scrolls. It is the one overlay surface: the page under it never moves (no layout shift).
PRIM_CSS += '''
.sheet-scrim{position:fixed;inset:0;background:#05070bcc;z-index:60}
.sheet{position:fixed;z-index:61;display:flex;flex-direction:column;background:var(--raised);border:1px solid var(--line);
right:0;left:0;bottom:0;max-height:88vh;border-radius:16px 16px 0 0}
@media(min-width:768px){.sheet{top:0;left:auto;width:480px;max-height:none;border-radius:0;border-width:0 0 0 1px}}
.sheet-h{display:flex;align-items:center;gap:12px;padding:16px 24px;border-bottom:1px solid var(--line-soft)}
.sheet-h h2{font:700 20px/24px var(--font-body);color:var(--strong);margin:0;flex:1}
.sheet-h .x{width:44px;height:44px;display:inline-flex;align-items:center;justify-content:center;color:var(--nav);font-size:22px}
.sheet-b{flex:1;overflow:auto;padding:8px 24px 16px}
.sheet-f{display:flex;gap:12px;justify-content:flex-end;padding:16px 24px;border-top:1px solid var(--line-soft)}
.sheet-f .btn{flex:none}@media(max-width:767px){.sheet-f .btn{flex:1}}
'''


def Sheet(title: str, body: str, actions: str) -> str:
    """An overlay panel: title and a close control, a scrolling body, a foot of actions. Rendered after the page body."""
    return (f'<div class="sheet-scrim" aria-hidden="true"></div><div class="sheet" role="dialog" aria-modal="true" aria-label="{title}">'
            f'<div class="sheet-h"><h2>{title}</h2><a class="x" href="#" aria-label="Close">&times;</a></div>'
            f'<div class="sheet-b">{body}</div><div class="sheet-f">{actions}</div></div>')
# ------------------------------------------------ Dialog (added by the account-delete lane, 2026-11-07)
# The shipped confirm dialog (ui/Overlay kind 'sheet' over a scrim; AccountDeleteDialog.svelte): from 768 a 560 px box centred
# in the viewport with the overlay's full edge and radius; below, a full-width bottom sheet of at most 88vh, top corners
# rounded, no bottom edge. Its head (title, one line) and foot (actions) stay put; the body scrolls. Region `dialog`.
PRIM_CSS += '''
.dlg-scrim{position:fixed;inset:0;background:#05070bcc;z-index:60}
.dlg{position:fixed;z-index:61;left:0;right:0;bottom:0;display:flex;flex-direction:column;max-height:88vh;outline:none}
.dlg>.dlg-box{display:flex;flex-direction:column;min-height:0;flex:1;background:var(--raised);border:1px solid var(--line);border-bottom:0;
 border-radius:6px 6px 0 0;box-shadow:0 12px 32px rgba(0,0,0,.6)}
@media(min-width:768px){.dlg{left:50%;right:auto;bottom:auto;top:50%;transform:translate(-50%,-50%);width:560px;max-width:calc(100vw - 48px);max-height:calc(100vh - 64px)}
 .dlg>.dlg-box{border-bottom:1px solid var(--line);border-radius:6px}}
.dlg-form{display:contents}.dlg-h{display:flex;flex-direction:column;gap:4px;padding:20px 20px 4px}
.dlg-h h2{margin:0;font:700 20px/24px var(--font-body);color:var(--strong)}
.dlg-h p{margin:0;font-size:15px;line-height:22px}
.dlg-b{flex:1;min-height:0;overflow:auto;padding:8px 20px 20px;display:flex;flex-direction:column}
.dlg-b h3{margin:16px 0 6px;font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.dlg-b h3.wipe{color:var(--wipe)}
.dlg-b ul{margin:0;padding-left:20px;list-style:disc;display:flex;flex-direction:column;gap:6px;font-size:15px;line-height:22px}
.dlg-b li::marker{color:var(--muted)}
.dlg-f{display:flex;flex-direction:column;gap:6px;margin-top:20px}.dlg-f label{font-size:15px;line-height:22px}
.dlg-f input{box-sizing:border-box;width:100%;height:44px;padding:0 12px;border:1px solid var(--line);border-radius:4px;background:var(--bg);color:var(--strong);font:14px var(--font-body)}
@media(min-width:768px){.dlg-f input{height:40px}}
.dlg-f .err{font-size:13px;line-height:20px;color:var(--wipe)}
.dlg-alert{margin:12px 0 0;font-size:14px;line-height:20px;color:var(--wipe)}
.dlg-a{display:flex;gap:12px;padding:16px 20px;border-top:1px solid var(--line-soft)}
@media(max-width:767px){.dlg-a{padding-bottom:calc(16px + env(safe-area-inset-bottom))}.dlg-a .btn{flex:1}}
'''


def Dialog(title: str, line: str, body: str, actions: str, title_id: str = 'dlg-title') -> str:
    """A modal confirm: title and one line, a scrolling body, a foot of actions; rendered after the page body."""
    return (f'<div class="dlg-scrim" aria-hidden="true"></div>'
            f'<div class="dlg" role="dialog" aria-modal="true" aria-labelledby="{title_id}" tabindex="-1" data-region="{REGION["dialog"]}">'
            f'<div class="dlg-box" role="none"><form class="dlg-form"><div class="dlg-h"><h2 id="{title_id}">{title}</h2><p>{line}</p></div>'
            f'<div class="dlg-b">{body}</div><div class="dlg-a">{actions}</div></form></div></div>')
# ------------------------------------------------ Paperdoll (moved from ws_bis.py 2026-11-06; the shipped `Paperdoll`, web/src/components/ui)
# The slot cells on the page's one Plate: two columns of slots ("By slot"), or the slots grouped by where they drop
# ("Where to go"). A cell is line 1 slot label and level flag, icon, name and short stats, the gain figure; line 2 the
# full source alone (spec 5.13 option C). BiS is the first user; the character gear tab and the sim's gear panel next.
PRIM_CSS += '''
.doll-cols{display:grid;grid-template-columns:minmax(0,1fr)}
@media(min-width:1024px){.doll-cols{grid-template-columns:minmax(0,1fr) minmax(0,1fr)}.doll-col+.doll-col{border-left:1px solid var(--line-soft)}
 .doll[data-view="where"] .doll-cols{display:block;columns:2;column-gap:0}}
.doll-grp{break-inside:avoid}.doll-grp-h{display:flex;align-items:baseline;gap:12px;padding:14px 16px 6px}
.doll-grp-h b{font:700 18px/24px var(--font-body);color:var(--strong)}.doll-grp-h span{margin-left:auto;font-size:13px;color:var(--muted)}
.cell{display:grid;grid-template-columns:64px 40px minmax(0,1fr) auto;column-gap:12px;align-items:start;min-height:58px;padding:9px 16px;border-bottom:1px solid var(--line-soft)}
.cell .lb{display:flex;flex-direction:column;padding-top:1px;font:600 13px/18px var(--font-body);color:var(--muted)}
.cell .lb .fl{font-size:12px;line-height:16px;font-weight:700}
.cell .bd{display:flex;flex-direction:column;gap:1px;min-width:0}
.cell .l1{display:flex;align-items:baseline;column-gap:10px;min-width:0}.cell .l1 .nm{flex:none;max-width:100%;font:700 15px/20px var(--font-body)}
.cell .l1 .st{min-width:0;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:13px;line-height:20px;color:var(--muted)}
.cell .l2{font-size:13px;line-height:18px;color:var(--muted);overflow-wrap:anywhere}
.cell .l2 .fact{display:inline-block;max-width:100%}  /* Paperdoll.svelte .ui-doll-fact: the line wraps between facts */
.cell .gn{display:flex;flex-direction:column;align-items:flex-end;width:144px;min-width:0;justify-self:end}
.cell .gn .v{font:700 16px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}.cell .gn .v small{font-size:13px;font-weight:600;color:var(--muted);margin-left:4px}
.cell .gn .bs{max-width:100%;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:12px;line-height:16px;color:var(--muted);text-align:right}
.cell .ic0{width:40px;height:40px;border-radius:5px;border:1px dashed var(--line-soft)}
'''

RARITY = {'poor': '#9d9d9d', 'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#3d94f0', 'epic': '#b866f5', 'legendary': '#ff8000'}
_RARITY_ICON = {'poor': 'q0', 'uncommon': 'q2', 'rare': 'q3', 'epic': 'q4', 'legendary': 'q5'}


def _colored(text: str, rarity: str) -> str:
    return f'<span style="color:{RARITY.get(rarity, RARITY["common"])}">{escape(text)}</span>'


def facts(line: str) -> str:
    """A source line as Paperdoll.svelte draws it: each fact whole, the trailing dot kept with it (factSegments)."""
    parts = line.split(' · ')
    return ' '.join(f'<span class="fact">{escape(p + (" ·" if i < len(parts) - 1 else ""))}</span>' for i, p in enumerate(parts))


def PaperdollCell(c: dict) -> str:
    """One slot. c: slot facts (label, flag, state, name, quality, icon, stats, line2, gain, basis), as bis_facts.mjs reads them."""
    flag = f'<span class="fl">{escape(c["flag"]["text"])}</span>' if c['flag'] else ''
    lead = f'<span class="lb">{escape(c["label"])}{flag}</span>'
    if c['state'] == 'empty':
        return (f'<div class="cell" role="option" aria-disabled="true">{lead}<span class="ic0"></span><span class="bd">'
                f'<span class="l1"><span class="nm muted" style="font-weight:400">{escape(c["name"])}</span></span><span class="l2">{facts(c["line2"])}</span></span><span></span></div>')
    ic = icon(c['icon'], 40, _RARITY_ICON.get(c['quality'], '')) if c['icon'] else '<span class="ic0"></span>'
    stats = f'<span class="st">{escape(c["stats"])}</span>' if c['stats'] else ''
    if c['gain'] is None and c['basis'] is None:   # `worn` mode (the sim): no gain column, as the built cell's empty auto track
        return (f'<div class="cell" role="option" aria-selected="false">{lead}{ic}<span class="bd"><span class="l1">'
                f'<span class="nm" style="color:{RARITY.get(c["quality"], RARITY["common"])}">{escape(c["name"])}</span>{stats}</span>'
                f'<span class="l2">{facts(c["line2"])}</span></span><span class="gn0"></span></div>')
    g = c['gain'] or {'value': '', 'unit': ''}
    unit = f'<small>{escape(g["unit"])}</small>' if g['unit'] else ''
    b = c['basis']
    basis = f'<span class="bs">{escape(b["prefix"])} {_colored(b["item"], b["quality"]) if b["item"] else ""}</span>' if b else ''
    return (f'<div class="cell" role="option" aria-selected="false">{lead}{ic}<span class="bd"><span class="l1">'
            f'<span class="nm" style="color:{RARITY.get(c["quality"], RARITY["common"])}">{escape(c["name"])}</span>{stats}</span>'
            f'<span class="l2">{facts(c["line2"])}</span></span><span class="gn"><span class="v">{escape(g["value"])}{unit}</span>{basis}</span></div>')


def Paperdoll(label: str, view: str, columns: tuple[list[dict], list[dict]] = ([], []), groups: list[dict] | None = None) -> str:
    """The slot listbox. view 'where' renders groups (title, count, cells); any other view the two columns (left, right)."""
    if view == 'where':
        inner = ''.join(f'<div class="doll-grp" role="group" aria-label="{escape(g["title"])}"><div class="doll-grp-h"><b>{escape(g["title"])}</b>'
                        f'<span>{escape(g["count"])}</span></div>{"".join(PaperdollCell(c) for c in g["cells"])}</div>' for g in groups or [])
    else:
        inner = ''.join(f'<div class="doll-col">{"".join(PaperdollCell(c) for c in side)}</div>' for side in columns)
    return (f'<div class="doll" data-view="{view}" data-region="{REGION["slots"]}">'
            f'<div class="doll-cols" role="listbox" aria-label="{escape(label)}">{inner}</div></div>')


# ------------------------------------------------ primitives added by the home round (2026-11-06; design/specs/2026-11-06-home.md section 2)
# ItemRow is the shipped ui/ItemRow (slot label | icon | name and source | the gain Figure), without a hairline (the shipped
# row's box makes every row its own last child). RowGrid is the shipped ui/RowGrid: open Rows in columns (one below 768,
# two to 1023, `cols` from 1024), the class picker first. Additive only: no earlier mock changes.
PRIM_CSS += '''
.irow{display:grid;grid-template-columns:auto minmax(0,1fr) auto;column-gap:12px;align-items:center;min-height:58px;padding:8px 16px}
.irow.slot{grid-template-columns:60px auto minmax(0,1fr) auto}
.irow .lbl{font:600 13px/18px var(--font-body);color:var(--muted)}
.irow .bd{min-width:0}.irow .nm{display:block;font:700 15px/20px var(--font-body)}.irow .src{display:block;font-size:13px;line-height:18px;color:var(--muted)}
.irow .tr{display:flex;gap:12px;align-items:center;justify-self:end}
.irow .g{display:inline-flex;flex-direction:column;align-items:flex-end;gap:2px;text-align:right}
.irow .g .v{font:700 16px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}.irow .g .v small{font-size:13px;font-weight:600;color:var(--muted);margin-left:4px}
.irow .g .c{font-size:12px;line-height:16px;color:var(--muted)}
.irow-box{container-type:inline-size;min-width:0}
@container (max-width:399px){.irow.slot{grid-template-columns:auto minmax(0,1fr) auto}.irow.slot .lbl{grid-column:2/-1;grid-row:1;font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase}.irow.slot>:nth-child(2){grid-row:1/span 2;grid-column:1}}
.rgrid{display:grid;grid-template-columns:minmax(0,1fr);column-gap:var(--gap)}
@media(min-width:768px){.rgrid{grid-template-columns:repeat(2,minmax(0,1fr))}}
@media(min-width:1024px){.rgrid{grid-template-columns:repeat(var(--cols),minmax(0,1fr))}}
.rgrid>.row{border-bottom:1px solid var(--line-soft)}
'''


def ItemRow(name: str, quality: int, icon_html: str, source: str, slot: str = '', gain: tuple[str, str, str] | None = None) -> str:
    """One item: slot label, icon, the name in its quality colour (a tooltip trigger), its source, the gain (value, unit, caption)."""
    lbl = f'<span class="lbl">{slot}</span>' if slot else ''
    g = ''
    if gain:
        v, u, c = gain
        g = f'<span class="g"><span class="v">{v}{f"<small>{u}</small>" if u else ""}</span>' + (f'<span class="c">{c}</span>' if c else '') + '</span>'
    return (f'<div class="irow-box"><div class="irow{" slot" if slot else ""}">{lbl}<span>{icon_html}</span><span class="bd">'
            f'<span class="nm" style="color:{QCOL.get(quality, QCOL[1])}" role="button" tabindex="0">{name}</span>'
            f'<span class="src">{source}</span></span><span class="tr">{g}</span></div></div>')


def RowGrid(rows: list[str], cols: int = 3) -> str:
    return f'<div class="rgrid" style="--cols:{cols}">' + ''.join(rows) + '</div>'


# ------------------------------------------------ primitives added by the logs round (2026-11-07; design/specs/2026-11-07-logs-reports.md section 3)
# Additive only. TabStrip gains overflow (`TabStripMore`: the shipped ModeBar "More" menu as a TabStrip mode: the first `md`
# tabs below 1760, the first `xl` from 1760, the rest behind "More"; the More item names the active tab when it is one of
# them). `LineChart` is the shipped report TimeChart (and the threat chart) as SVG on S0: a legend list, phase bands, death
# marks, an x axis; no surface. `Spark` is a 1-line series in a table cell (resources, a cast sequence). `Menu` is the
# Overlay of kind 'menu' (the fight picker): a panel anchored under its control, grouped rows. `DropZone` is the upload
# target (dashed socket, the one dashed surface, like the empty doll socket). `Disclosure` is the shipped one (summary +
# body), used for the phone chart and the death recap.
PRIM_CSS += '''
.strip-tabs .tm{display:none}.strip-tabs .more{display:inline-flex;gap:6px}
.strip-tabs button.more{padding:14px;font:600 14px/16px var(--font-body);color:var(--nav);background:none;border:0;border-bottom:2px solid transparent;white-space:nowrap;cursor:pointer}
@media(max-width:1759px){.strip-tabs a.xmd{display:none}}
@media(min-width:1760px){.strip-tabs a.xxl{display:none}.strip-tabs .more.nmd{display:none}}
.strip-tabs .more.cur{color:var(--strong);border-bottom-color:var(--gold)}
@media(max-width:767px){.strip-tabs a.xmd{display:none}}
.lc{min-width:0}.lc .lg{display:flex;flex-wrap:wrap;gap:4px 18px;padding:2px 0 8px;font-size:13px;line-height:20px;color:var(--muted)}
.lc .lg li{display:inline-flex;align-items:center;gap:6px}.lc .lg b{font-weight:600;color:var(--c)}.lc .lg i{display:inline-block;width:14px;height:3px;border-radius:2px;background:var(--c)}
.lc svg{display:block;width:100%;overflow:visible}.lc .ax{display:flex;justify-content:space-between;font:600 12px/16px var(--font-body);color:var(--muted);font-variant-numeric:tabular-nums;padding-top:4px}
.spark{display:block;width:100%;height:24px}
.menu{position:absolute;z-index:40;width:min(480px,calc(100vw - 24px));max-height:70vh;overflow:auto;background:var(--raised);border:1px solid var(--line);border-radius:6px;box-shadow:0 18px 40px #000a}
.menu .mg{padding:10px 16px 4px;font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.menu .mi{display:grid;grid-template-columns:minmax(0,1fr) auto;gap:2px 12px;padding:8px 16px;border-bottom:1px solid var(--line-soft)}
.menu .mi b{font:700 15px/20px var(--font-body);color:var(--strong)}.menu .mi span{font-size:13px;line-height:18px;color:var(--muted)}
.menu .mi .r{text-align:right}.menu .mi.cur{background:#e5b9551a;box-shadow:inset 3px 0 0 var(--gold)}.menu .mi:hover{background:var(--hover)}
.drop{display:flex;flex-direction:column;align-items:center;justify-content:center;gap:6px;min-height:96px;padding:16px 20px;border:1px dashed var(--line);border-radius:6px;text-align:center;color:var(--muted);font-size:14px}
.drop.off{opacity:.55}
.disc>summary{display:flex;align-items:center;gap:10px;min-height:44px;cursor:pointer;list-style:none;font:600 14px/20px var(--font-body);color:var(--strong)}
.disc>summary::-webkit-details-marker{display:none}.disc>summary .chg{margin-left:auto;color:var(--gold);font-size:13px}
.crest.none{display:inline-block;box-shadow:0 0 0 2px var(--line)}
.st-h .st-sort{all:unset;cursor:pointer}.st-h .st-sort:focus-visible{outline:2px solid var(--gold);outline-offset:2px}
.skel{max-width:100%}.ph-title>div{min-width:0}
.lc-only-d{display:none}@media(min-width:768px){.lc-only-d{display:block}.lc-only-m{display:none}}
'''


def TabStripMore(tabs: list[str], active: str, md: int, xl: int) -> str:
    """TabStrip with the shipped "More" overflow: tabs[:md] always (below 1760), tabs[:xl] from 1760, the rest in More."""
    out = ''
    for i, t in enumerate(tabs[:xl]):
        cls = ('cur' if t == active else '') + (' xmd' if i >= md else '')
        out += f'<a href="#" class="{cls.strip()}"{" aria-current=\"page\"" if t == active else ""}>{t}</a>'
    rest_md, rest_xl = tabs[md:], tabs[xl:]
    def more(rest, cls):
        on = active in rest
        lab = f'More: {active}' if on else 'More'
        return f'<button type="button" class="more {cls}{" cur" if on else ""}" aria-haspopup="menu" aria-expanded="false">{lab} &#9662;</button>'
    return '<nav class="strip-tabs ph-tabs" aria-label="Report tabs">' + out + more(rest_md, 'nmd') + more(rest_xl, 'nxl') + '</nav>'


PRIM_CSS += '@media(max-width:1759px){.strip-tabs .more.nxl{display:none}}\n'


def LineChart(series: list[tuple[str, str, list[float], str]], duration_s: float, height: int = 180, phases: list[tuple] = (),
              marks: list[tuple[float, str]] = (), ymax: float | None = None, unit: str = '', label: str = 'Chart', step_ms: float = 0) -> str:
    """series: (name, colour, values per bucket, legend figure). phases: (name, start_s, end_s). marks: (at_s, colour) death marks."""
    n = max((len(v) for _, _, v, _ in series), default=1)
    top = ymax or max((max(v) for _, _, v, _ in series if v), default=1) or 1
    W = 1000
    def x(i, cnt): return i * W / max(cnt - 1, 1)
    svg = ''
    for _, start, end in [(p[0], p[1], p[2]) for p in phases]:
        svg += f'<line x1="{start / duration_s * W:.1f}" x2="{start / duration_s * W:.1f}" y1="0" y2="{height}" stroke="#3a4256" stroke-dasharray="3 4"/>'
    for i in range(1, 4):
        svg += f'<line x1="0" x2="{W}" y1="{height * i / 4:.0f}" y2="{height * i / 4:.0f}" stroke="#1c2230"/>'
    for name, col, vals, _ in series:
        if not vals:
            continue
        # step_ms: the summary's series are one value per second from the pull (lib/report/types.ts), so x is time, not index
        xs = (lambda i: i * step_ms / 1000 / duration_s * W) if step_ms else (lambda i: x(i, len(vals)))
        pts = ' '.join(f'{xs(i):.1f},{height - v / top * (height - 8):.1f}' for i, v in enumerate(vals))
        svg += f'<polyline fill="none" stroke="{col}" stroke-width="2" vector-effect="non-scaling-stroke" points="{pts}"/>'
    for at, col in marks:
        svg += f'<line x1="{at / duration_s * W:.1f}" x2="{at / duration_s * W:.1f}" y1="0" y2="{height}" stroke="{col}" stroke-width="2" vector-effect="non-scaling-stroke"/>'
    ph = ''.join(f'<text x="{p[1] / duration_s * W + 6:.0f}" y="14" fill="#9a9484" font-size="11" font-weight="700" letter-spacing="1.5">{escape(p[0].upper())}</text>' for p in phases)
    lg = '<ul class="lg">' + ''.join(f'<li style="--c:{c}"><i></i><b>{escape(nm)}</b> {fig}</li>' for nm, c, _, fig in series) + '</ul>'
    m = int(duration_s // 60)
    ticks = [f'{int(t // 60)}:{int(t % 60):02d}' for t in (0, duration_s / 4, duration_s / 2, duration_s * 3 / 4, duration_s)]
    return (f'<div class="lc">{lg}<svg viewBox="0 0 {W} {height}" preserveAspectRatio="none" style="height:{height}px" role="img" aria-label="{escape(label)}">{svg}{ph}</svg>'
            f'<div class="ax">' + ''.join(f'<span>{t}</span>' for t in ticks) + '</div></div>')


def Spark(values: list[float], color: str = '#e5b955', marks: list[float] | None = None, total: float | None = None) -> str:
    """A one-line series (resources) or, with `marks` (0..1 positions), a tick strip (a cast sequence)."""
    if marks is not None:
        ticks = ''.join(f'<rect x="{p * 996:.1f}" y="4" width="4" height="16" rx="1" fill="{color}"/>' for p in marks)
        return f'<svg class="spark" viewBox="0 0 1000 24" preserveAspectRatio="none" aria-hidden="true"><rect x="0" y="11" width="1000" height="2" fill="#1c2230"/>{ticks}</svg>'
    top = total or max(values or [1]) or 1
    pts = ' '.join(f'{i * 1000 / max(len(values) - 1, 1):.1f},{22 - v / top * 20:.1f}' for i, v in enumerate(values))
    return f'<svg class="spark" viewBox="0 0 1000 24" preserveAspectRatio="none" aria-hidden="true"><polyline fill="none" stroke="{color}" stroke-width="2" vector-effect="non-scaling-stroke" points="{pts}"/></svg>'


def Menu(label: str, groups: list[tuple[str, list[tuple[str, str, str, str, bool]]]]) -> str:
    """Overlay kind 'menu': groups of (title, sub, right_top, right_sub, current) rows; one is current."""
    body = ''
    for g, rows in groups:
        body += f'<div class="mg" role="presentation">{g}</div>' + ''.join(
            f'<a href="#" class="mi{" cur" if cur else ""}" role="menuitemradio" aria-checked="{"true" if cur else "false"}"><b>{t}</b><span class="r">{rt}</span><span>{s}</span><span class="r">{rs}</span></a>'
            for t, s, rt, rs, cur in rows)
    return f'<div class="menu" role="menu" aria-label="{escape(label)}">{body}</div>'


def DropZone(lines: str, action: str, off: bool = False) -> str:
    return f'<div class="drop{" off" if off else ""}" role="group" aria-label="Choose a combat log">{lines}{action}</div>'


def Disclosure(summary: str, body: str, open_: bool = False, cls: str = '') -> str:
    return f'<details class="disc {cls}"{" open" if open_ else ""}><summary>{summary}<span class="chg">{"Hide" if open_ else "Show"} &#9662;</span></summary>{body}</details>'


def EmptyCrest(px: int) -> str:
    """A character whose class the log does not carry: the ringed empty disc of the one crest language."""
    return f'<span class="crest none" style="width:{px}px;height:{px}px"></span>'

# ------------------------------------------------ primitives added by the character round (2026-11-06; design/specs/2026-11-06-character.md)
# Disclosure is the shipped ui/Disclosure.svelte: a summary line (the body face 600, a caret) that opens its body in place,
# no box. Additive only.
PRIM_CSS += '''
.disc>summary{display:flex;align-items:center;gap:8px;min-height:44px;cursor:pointer;list-style:none;font:600 14px/20px var(--font-body);color:var(--nav)}
.disc>summary::-webkit-details-marker{display:none}.disc>summary::before{content:'\\25B8';font-size:11px;color:var(--muted)}
.disc[open]>summary::before{content:'\\25BE'}.disc>summary .t-small{margin-left:4px}
.disc>.disc-b{padding:4px 0 8px}
'''


def Disclosure(summary: str, body: str, open_: bool = False, meta: str = '') -> str:
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    return f'<details class="disc"{" open" if open_ else ""}><summary>{summary}{m}</summary><div class="disc-b">{body}</div></details>'

# ------------------------------------------------ primitives added by the setup round (2026-11-07; design/specs/2026-11-07-setup.md section 3A)
# Additive only. `Steps` is the shipped ui/Checklist with numbered marks and an action in the trail (a setup or onboarding
# list: the first step not done carries the page's one primary Button). Marks: ok (done, the kill check), next (the first
# step not done: gold disc with its number), todo (not done, optional or later: an outlined ring with its number), warn
# (needs the player: gold "!"), miss (failed: wipe "×"), busy (in progress: a gold ring turning). `Capture` is a client
# screenshot at its own aspect in a quiet frame (the one image frame on the site: /setup's surfaces and the companion,
# later the guides' in-game shots), with an optional muted Tag under it. `Code` is literal text a player copies or matches
# (slash commands, pairing and device codes): the one use of a mono face. `PageForm` is Page layout 'form': one column
# at --measure-form (560 px), left-aligned in the page column, for a page that is one task (connect, sign in, checkout).
PRIM_CSS += '''
.steps .row{min-height:64px;padding-block:12px}
.smk{width:26px;height:26px;border-radius:50%;display:inline-flex;align-items:center;justify-content:center;flex:none;font:700 13px/26px var(--font-body);font-variant-numeric:tabular-nums}
.smk.ok{background:var(--kill);color:#07090d}.smk.next{background:var(--gold);color:#17110a}.smk.todo{box-shadow:inset 0 0 0 1.5px var(--line);color:var(--muted)}
.smk.warn{background:var(--gold);color:#17110a}.smk.miss{background:var(--wipe);color:#07090d}
.smk.busy{box-shadow:inset 0 0 0 2px var(--line);color:var(--muted);position:relative}
.smk.busy::after{content:"";position:absolute;inset:0;border-radius:50%;border:2px solid transparent;border-top-color:var(--gold);animation:sm-spin 1s linear infinite}
@keyframes sm-spin{to{transform:rotate(360deg)}}@media(prefers-reduced-motion:reduce){.smk.busy::after{animation:none}}
.steps .row .sub{font-size:14px;line-height:20px;margin-top:2px}.steps .row .sub .warn{color:var(--gold)}.steps .row .sub .miss{color:var(--wipe)}
.steps .row .sub .ok{color:var(--text)}
.steps .row .trail .btn{min-width:0}
.btn.is-off{opacity:.45;pointer-events:none}
.cap{display:block;position:relative;width:100%;max-width:var(--w);max-height:var(--mh,420px);aspect-ratio:var(--ar);border:1px solid var(--line-soft);border-radius:4px;background:var(--raised);overflow:hidden}
.cap img{display:block;width:100%;height:100%;object-fit:contain}
.cap picture{display:contents}@media(max-width:639px){.cap.has-narrow{aspect-ratio:var(--ar-n)}}
.cap-wrap{display:flex;flex-direction:column;gap:8px;align-items:flex-start;min-width:0}
code.c{font:600 14px/20px ui-monospace,'SF Mono',Menlo,monospace;color:var(--strong);letter-spacing:.02em}
code.c.big{font-size:24px;line-height:30px;letter-spacing:.06em}
.layout-form{max-width:560px}
'''


def Steps(items: list[tuple]) -> str:
    """items: (mark, number, title, sub_html, action_html). mark: ok | next | todo | warn | miss | busy."""
    word = {'ok': 'Done', 'next': 'Next', 'todo': 'Not done', 'warn': 'To do', 'miss': 'Failed', 'busy': 'In progress'}
    rows = []
    for mark, n, title, sub, action in items:
        glyph = {'ok': '&#10003;', 'warn': '!', 'miss': '&times;', 'busy': ''}.get(mark, str(n))
        lead = f'<span class="smk {mark}" role="img" aria-label="Step {n}: {word[mark]}">{glyph}</span>'
        rows.append(Row(title, sub, lead=lead, trail=action, link=False))
    return f'<div class="steps">{Rows(rows)}</div>'


def Capture(src: str, w: int, h: int, alt: str, tag: str = '', tall: bool = False, lazy: bool = True,
            narrow: dict | None = None) -> str:
    """A client screenshot at its own aspect, capped at its native width and 420 px high (900 for a portrait capture).
    `narrow` ({src, width, height}: a tighter crop of the same picture) replaces it under 640 px, as ui/Capture.svelte's does."""
    t = f'{Tag(tag, "muted")}' if tag else ''
    lz = ' loading="lazy"' if lazy else ''
    image = f'<img src="{src}" width="{w}" height="{h}" alt="{escape(alt, quote=True)}"{lz} decoding="async">'
    cls, ar_narrow = 'cap', ''
    if narrow is not None:
        cls, ar_narrow = 'cap has-narrow', f';--ar-n:{narrow["width"]} / {narrow["height"]}'
        image = (f'<picture><source media="(max-width: 639px)" srcset="{narrow["src"]}" width="{narrow["width"]}" '
                 f'height="{narrow["height"]}">{image}</picture>')
    # The picture links to its own file (spec 3A), as the built Capture does.
    return (f'<div class="cap-wrap"><a class="{cls}" href="{src}" '
            f'style="--w:{w}px;--ar:{w} / {h};--mh:{900 if tall else 420}px{ar_narrow}">{image}</a>{t}</div>')


def Code(text: str, big: bool = False) -> str:
    return f'<code class="c{" big" if big else ""}">{escape(text)}</code>'


def PageForm(main: str) -> str:
    """Page layout 'form': one column at 560 px under the page header, left-aligned in the page column."""
    return f'<div class="pc" style="padding-top:20px"><div class="layout-form">{main}</div></div>'

# ------------------------------------------------ primitives added by the guides round (2026-11-08; design/specs/2026-11-08-guides.md section 3A)
# Additive only: no earlier mock changes (PageHeader and Page are redefined as supersets; with the new arguments left at their
# defaults they emit exactly what they emitted before). `PageHeader(lead=)` is the Lead type under the title (one sentence,
# 17/26, max 62ch: the system's own type role, 3.3). `Page('reading')` is L2 Reading (3.7): a 200 px contents rail from 1760,
# the sections beside it. `ReadingSection` is a Section on a reading page: its head (22/28), then the prose (720) and its lane
# (440 to 640) side by side from 1280, the lane first in one column below. `Prose` (class rprose: logs owns .prose) is the rendered guide markdown's typography
# (paragraphs, lists, h3, the game-tooltip names, the ruled fact list). `Toc` is the contents rail. `PlateColumns` is a Plate
# split into labelled columns by hairlines (the guide's At a glance band; next: the character summary and the sim result band).
# `TalentTrees` is `Trees` at a compact cell (24 px, 28 from 1760). `LevelSlider` is the shipped GuideLevelingTree's control
# row (previous, the level range, next). Rows gain the contract's `density='slim'` (44 px).
PRIM_CSS += '''
.ph-lead{margin-top:10px;font-size:17px;line-height:26px;max-width:62ch;color:var(--text)}
@media(max-width:767px){.ph-lead{font-size:16px;line-height:24px}}
.layout-reading{display:grid;grid-template-columns:minmax(0,1fr);column-gap:var(--gap);align-items:start}
.layout-reading>.toc{display:none}
@media(min-width:1760px){.layout-reading{grid-template-columns:var(--toc-w) minmax(0,1fr)}.layout-reading>.toc{display:block;position:sticky;top:76px}}
.layout-reading>.body{display:flex;flex-direction:column;gap:var(--section-gap);min-width:0}
.rsec>.sec-h h2{font:700 22px/28px var(--font-body);color:var(--strong)}
@media(max-width:767px){.rsec>.sec-h h2{font-size:20px;line-height:26px}}
.rd{display:grid;grid-template-columns:minmax(0,var(--measure-reading));grid-template-areas:"lane" "prose";gap:24px var(--gap);padding-top:20px}
.rd.solo{grid-template-areas:"prose"}
@media(min-width:1280px){.rd{grid-template-columns:var(--measure-reading) minmax(0,640px);grid-template-areas:"prose lane"}.rd.solo{grid-template-columns:var(--measure-reading)}}
.rd>.rd-prose{grid-area:prose;min-width:0}.rd>.rd-lane{grid-area:lane;min-width:0;display:flex;flex-direction:column;gap:28px}
.lane-blk{display:flex;flex-direction:column;gap:8px;min-width:0}.lane-blk>.t-label{padding-bottom:2px}
.rsec-fold-sum{display:none}
@media(max-width:767px){.rsec-fold-sum{display:flex;align-items:center;min-height:44px;color:var(--gold);font-size:15px;font-weight:600;line-height:24px;list-style:none}.rsec-fold-sum::-webkit-details-marker{display:none}.rsec-fold::details-content{content-visibility:hidden}}
.lane-lbl{margin:0;font:600 15px/20px var(--font-body);color:#f3eee2;letter-spacing:0;text-transform:none}
.rprose{font-size:16px;line-height:26px;color:var(--text);max-width:var(--measure-reading);overflow-wrap:break-word}
.rprose.lane{font-size:15px;line-height:24px}
.rprose p{margin:0 0 14px}.rprose>:last-child{margin-bottom:0}
.rprose ul,.rprose ol{margin:0 0 14px;padding-left:20px;list-style:disc}.rprose ol{list-style:decimal}.rprose li{margin:0 0 8px}.rprose li::marker{color:var(--muted)}
.rprose h3{font:700 18px/24px var(--font-body);color:var(--strong);margin:28px 0 10px}.rprose>h3:first-child{margin-top:0}
.rprose strong{color:var(--strong);font-weight:700}.rprose a{font-weight:600}
.rprose .gtip{text-decoration:underline dotted;text-decoration-color:color-mix(in srgb,currentColor 40%,transparent);text-underline-offset:3px;cursor:pointer}
.rprose span.gtip:not(.gtip-item){color:var(--strong);font-weight:500}
.rprose .glabel{font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted);margin:18px 0 4px}.rprose .glabel:first-child{margin-top:0}
.rprose.ruled ul{list-style:none;padding:0;margin:0}.rprose.ruled li{position:relative;margin:0;padding:12px 0 12px 20px;border-bottom:1px solid var(--line-soft);break-inside:avoid}
.rprose.ruled li::before{content:"";position:absolute;left:0;top:21px;width:6px;height:6px;border-radius:50%;background:var(--gold)}
@media(min-width:1024px){.rprose.cols{max-width:none}.rprose.cols ul{column-count:2;column-gap:48px}}
.rprose.plus ul,.rprose.minus ul{list-style:none;padding:0}.rprose.plus li,.rprose.minus li{position:relative;padding-left:20px}
.rprose.plus li::before,.rprose.minus li::before{position:absolute;left:0;font-weight:700}.rprose.plus li::before{content:"+";color:var(--kill)}.rprose.minus li::before{content:"\\2212";color:var(--wipe)}
.toc .t-label{display:block;padding-bottom:10px}.toc ol{list-style:none;margin:0;padding:0;display:flex;flex-direction:column}
.toc a{display:flex;align-items:center;min-height:32px;padding-left:12px;font:600 14px/18px var(--font-body);color:var(--nav);box-shadow:inset 2px 0 0 var(--line-soft)}
.toc a:hover{color:var(--strong)}.toc a[aria-current]{color:var(--gold);box-shadow:inset 2px 0 0 var(--gold)}
.pcols{display:grid;grid-template-columns:minmax(0,1fr)}
@media(min-width:1024px){.pcols{grid-template-columns:repeat(var(--n),minmax(0,1fr))}.pcols.n1{grid-template-columns:minmax(0,1fr)}}
.pcol{display:flex;flex-direction:column;gap:12px;padding:16px 20px 20px;min-width:0}
.pcol+.pcol{border-top:1px solid var(--line-soft)}
@media(min-width:1024px){.pcol+.pcol{border-top:0;border-left:1px solid var(--line-soft)}}
.pcol-h{display:flex;align-items:baseline;gap:12px}.pcol-h .t-small{margin-left:auto;text-align:right}
.pcol .st-r{min-height:40px;padding-block:6px}.pcol .st-r .nm b{white-space:nowrap}
@media(max-width:1023px){.pcol .st .stc{grid-template-columns:28px minmax(0,1fr) minmax(0,96px) auto}.pcol .st.open .st-r .br{display:block;grid-column:auto;padding-top:0}.pcol .st.open .st-r .f.p{grid-row:auto;grid-column:auto}}.pcol .rows.slim .row{padding-block:7px}
.pcol-f{margin-top:auto;display:flex;gap:12px;flex-wrap:wrap;align-items:center;padding-top:4px}
.tbox{--cg:4px;--tg:16px;container-type:inline-size}
.trees.cmp{--cell:clamp(18px,calc((100cqw - 2 * var(--tg) - 9 * var(--cg)) / 12 - 0.5px),32px);gap:var(--tg);grid-template-columns:repeat(3,calc(4 * var(--cell) + 3 * var(--cg)))}.trees.cmp .tg>[role=row]{display:contents}.trees.cmp .tc{display:block}.trees.cmp .tr h3{font-size:13px;line-height:18px;padding-bottom:6px;white-space:nowrap}
.trees.cmp .tg{grid-template-columns:repeat(4,var(--cell));gap:var(--cg)}.trees.cmp .tc,.trees.cmp .tc img,.trees.cmp .tc .ic{width:var(--cell)!important;height:var(--cell)!important}
.trees.cmp .tc b{font-size:10px;line-height:12px;right:-3px;bottom:-4px}
@media(max-width:767px){.trees.cmp .tr h3{white-space:normal;flex-wrap:wrap;column-gap:6px;row-gap:0;min-height:42px;align-content:flex-end}.tbox{--cg:3px;--tg:10px}}
@container (max-width:270px){.trees.cmp{grid-template-columns:calc(4 * var(--cell) + 3 * var(--cg))}}
.lvl{display:flex;align-items:center;gap:10px}.lvl .bt{all:unset;box-sizing:border-box;width:44px;height:44px;display:inline-flex;align-items:center;justify-content:center;border:1px solid var(--line);border-radius:4px;color:var(--nav);font-size:18px}
.lvl .bt:hover{color:var(--strong);border-color:var(--gold)}.lvl .bt.off{opacity:.45}
.lvl .rg{position:relative;flex:1;height:44px;display:flex;align-items:center}.lvl .rg .tk{height:4px;width:100%;border-radius:2px;background:var(--line-soft)}
.lvl .rg .fl{position:absolute;left:0;height:4px;border-radius:2px;background:var(--gold)}
.lvl .rg .th{position:absolute;width:20px;height:20px;margin-left:-10px;border-radius:50%;background:var(--gold);box-shadow:0 0 0 4px #e5b95533}
.rows.slim .row{min-height:44px;padding-block:6px}
/* MUST 5 (guides): the open phone bar wraps a segmented control that does not fit, so the current choice is never cut off */
.bar-m .seg{flex-wrap:wrap}.bar-m .seg a{min-width:max-content}
.ans-box{display:flex;flex-direction:column;gap:6px;min-width:0}
@media(min-width:900px){.ans-box{min-width:380px;min-height:72px;align-items:flex-end;text-align:right}}
'''


_PageHeader0 = PageHeader
_Page0 = Page


def PageHeader(title: str, title_color: str = '', crest: str = '', meta: list[str] | None = None, answer: str = '', actions: str = '', tabs: str = '',
               lead: str = '', crumb: str = '', notice: str = '', pending: bool = False, second_head: bool = False) -> str:
    """PageHeader with the optional Lead sentence directly under the title, the meta line under it (inside head-title,
    so the regions are unchanged; seo-fix 2026-11-10: the lead is the page's first sentence, never under a caveat)."""
    html = _PageHeader0(title, title_color, crest, meta, answer, actions, tabs, crumb=crumb, notice=notice, pending=pending,
                        second_head=second_head)
    if not lead:
        return html
    cut = html.index('</h1>') + len('</h1>') if '</h1>' in html else html.index('</p>') + len('</p>')
    return html[:cut] + f'<p class="ph-lead">{lead}</p>' + html[cut:]


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False) -> str:
    """Page with layout 'reading' (L2): `rail` is the Toc (shown from 1760), `main` the ReadingSections."""
    if layout != 'reading':
        return _Page0(layout, main, rail, sticky_rail)
    return f'<div class="pc" style="padding-top:var(--section-gap)"><div class="layout-reading"><nav class="toc" aria-label="Sections in this guide">{rail}</nav><div class="body">{main}</div></div></div>'


def Prose(html: str, *mods: str) -> str:
    """Rendered guide markdown. mods: 'lane' (15/24, in a lane), 'ruled' (a fact list with gold marks and hairlines), 'cols' (two
    columns from 1024), 'plus' / 'minus' (Strengths and Weaknesses marks)."""
    return f'<div class="rprose {" ".join(mods)}">{html}</div>'


def ReadingSection(title: str, prose: str, lane: str = '', anchor: str = '', action: str = '', prose_first: bool = False,
                   fold: str = '') -> str:
    """A section of a reading page: the Section head (22/28), then prose and lane (lane first in the DOM: one column reads it first).
    `fold` (web ReadingSection `fold`): below 768 the body folds behind that line, from 768 it reads open."""
    a = f'<span class="act">{action}</span>' if action else ''
    lane_html = f'<div class="rd-lane" data-reading-lane>{lane}</div>' if lane else ''
    more = '<div class="rd-more" data-reading-more></div>' if lane else ''
    solo = '' if lane else ' data-solo'
    body = (f'<div class="rd{"" if lane else " solo"}" data-reading{solo}{prose_first and " data-prose-first" or ""}>{lane_html}'
            f'<div class="rd-prose" data-reading-prose>{prose}{more}</div></div>')
    if fold:
        # open, as the build serves it; on a phone the body is hidden as the build's script closes it
        body = f'<details class="rsec-fold" open><summary class="rsec-fold-sum">{fold}</summary>{body}</details>'
    return f'<section class="rsec" data-region="{region_slug(title)}" id="{anchor}"><div class="sec-h"><h2>{title}</h2>{a}</div>{body}</section>'


def LaneBlock(label: str, body: str, meta: str = '') -> str:
    """One labelled block in a lane (web ui/LaneBlock: a 15 px cream heading, then its rows or table)."""
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    return f'<div class="lane-blk"><div class="pcol-h"><h3 class="lane-lbl">{label}</h3>{m}</div>{body}</div>'


def Toc(label: str, items: list[tuple[str, str]], current: str = '') -> str:
    """The contents rail: (title, anchor); the current section gold with a 2 px bar (aria-current="location")."""
    lis = ''.join(f'<li><a href="#{a}"{" aria-current=\"location\"" if a == current else ""}>{t}</a></li>' for t, a in items)
    return f'<span class="t-label">{label}</span><ol>{lis}</ol>'


def PlateColumns(cols: list[tuple[str, str, str, str]], label: str = '') -> str:
    """A Plate in labelled columns: (label, meta, body, foot). Hairlines between, stacked below 1024."""
    inner = ''.join(f'<div class="pcol"><div class="pcol-h"><h2 class="t-label">{lab}</h2>{f"<span class=t-small>{meta}</span>" if meta else ""}</div>{body}'
                    + (f'<div class="pcol-f">{foot}</div>' if foot else '') + '</div>' for lab, meta, body, foot in cols)
    return Plate(f'<div class="pcols{" n1" if len(cols) == 1 else ""}" style="--n:{len(cols)}">{inner}</div>', label)


def TalentTrees(trees: list[tuple[str, int, list[tuple]]]) -> str:
    """`Trees` at the compact cell: three trees side by side in a 400 px column (the At a glance band, the leveling lane)."""
    # the built TalentTrees: the cells size themselves from the box (a container query, 18 to 32 px)
    return '<div class="tbox">' + Trees(trees).replace('<div class="trees">', '<div class="trees cmp">', 1) + '</div>'


def LevelSlider(level: int, lo: int, hi: int, label: str) -> str:
    """The shipped GuideLevelingTree control row: previous level, the level range (a slider), next level."""
    frac = (level - lo) / (hi - lo)
    return (f'<div class="lvl"><button type="button" class="bt{" off" if level == lo else ""}" aria-label="Previous level"{" disabled" if level == lo else ""}>&lsaquo;</button>'
            f'<span class="rg" role="slider" aria-label="{label}" aria-valuemin="{lo}" aria-valuemax="{hi}" aria-valuenow="{level}" aria-valuetext="Level {level}" tabindex="0">'
            f'<span class="tk"></span><span class="fl" style="width:{frac * 100:.1f}%"></span><span class="th" style="left:{frac * 100:.1f}%"></span></span>'
            f'<button type="button" class="bt{" off" if level == hi else ""}" aria-label="Next level"{" disabled" if level == hi else ""}>&rsaquo;</button></div>')


def RowsSlim(rows: list[str], open_: bool = True) -> str:
    """Rows at the contract's density 'slim' (44 px)."""
    return Rows(rows, open_).replace('<div class="rows', '<div class="rows slim', 1)


# ------------------------------------------------ primitives added by the home B round (2026-11-07; design/specs/2026-11-07-home-b.md)
# Countdown (ui/Countdown.svelte): a panel with a caps eyebrow (live dot), a days/hours/min clock read as one timer, and its
# content at the right. FeedRow (ui/FeedRow.svelte): one feed line as a link: icon or crest, the sentence, the time over a caps
# kind. FeedSkeleton: the feed while it loads. GradeTile: a tier letter on its tier colour. Row gains `slim` (44 px).
PRIM_CSS += '''
.row.slim{min-height:44px;padding-block:6px}
.cdn{position:relative;min-width:0;padding:20px 24px;border:1px solid var(--line);border-radius:8px;background:var(--raised)}
.cdn.raid{border-color:#3b2a4d;background:radial-gradient(500px 200px at 100% 0,#5a1d7a33,transparent 70%),var(--raised)}
.cdn.guild{background:radial-gradient(400px 200px at 100% 0,#e5b9551a,transparent 70%),var(--raised)}
.cdn .eb{margin:0;font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.cdn .live{display:inline-block;width:8px;height:8px;margin-right:8px;border-radius:50%;background:var(--kill);vertical-align:1px}
.cdn .rw{display:flex;flex-wrap:wrap;align-items:flex-end;gap:16px 28px;margin-top:10px}
.cdn .clk{display:flex;gap:14px;font-variant-numeric:tabular-nums}
.cdn .clk>span{display:flex;flex-direction:column;align-items:center}
.cdn .clk b{font:700 40px/42px var(--font-body);color:var(--strong)}
.cdn .clk>span>span{font-size:11px;line-height:16px;letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.cdn .as{flex:1 1 200px;min-width:0;margin-left:auto}
.frow{display:grid;grid-template-columns:40px minmax(0,1fr) auto;gap:14px;align-items:center;padding:12px 0;border-bottom:1px solid var(--line-soft);font-size:15px;line-height:22px;color:inherit}
.frow .ld{display:flex;line-height:0}.frow .tx{min-width:0}.frow .tx b{font-weight:700}
.frow .wh{text-align:right;font-size:13px;line-height:20px;white-space:nowrap;color:var(--muted)}
.frow .wh .k{display:block;font-size:11px;line-height:16px;letter-spacing:.1em;text-transform:uppercase;font-weight:700;color:#7f7864}
.fsk{display:flex;flex-direction:column}.fsk>span{display:flex;height:65px;align-items:center;gap:14px;border-bottom:1px solid var(--line-soft)}
.grade{display:inline-flex;flex:none;align-items:center;justify-content:center;width:56px;height:56px;border-radius:8px;font:700 32px/1 var(--font-body);color:#17110a}
'''

TIER_TONES = {'S': '#ff8000', 'A': '#b866f5', 'B': '#3d94f0', 'C': '#1eff00', 'D': '#9d9d9d'}
PRIM_CSS += ''.join(f'.grade.t{k}{{background:{v}}}.grade.bare.t{k}{{background:none;color:{v}}}' for k, v in TIER_TONES.items()) + '\n'
PRIM_CSS += '.grade.bare{font-size:40px}\n'


def Countdown(eyebrow: str, clock: list[tuple[str, str]] | None, aside: str = '', tone: str = 'neutral', live: bool = False) -> str:
    """clock: (value, unit) parts, or None for no clock. The clock is one timer read aloud; the parts are hidden from it."""
    dot = '<span class="live" aria-hidden="true"></span>' if live else ''
    clk = ''
    if clock:
        label = ', '.join(f'{int(v)} {u}' for v, u in clock)
        clk = f'<div class="clk" role="timer" aria-label="{label}">' + ''.join(f'<span aria-hidden="true"><b>{v}</b><span>{u}</span></span>' for v, u in clock) + '</div>'
    side = f'<div class="as">{aside}</div>' if aside else ''
    return f'<div class="cdn {tone}"><p class="eb">{dot}{eyebrow}</p><div class="rw">{clk}{side}</div></div>'


def FeedRow(lead: str, sentence: str, when: str, kind: str) -> str:
    return f'<a class="frow" href="#"><span class="ld">{lead}</span><span class="tx">{sentence}</span><span class="wh">{when}<span class="k">{kind}</span></span></a>'


def FeedSkeleton(rows: int = 5) -> str:
    return '<div class="fsk" role="status" aria-busy="true">' + ''.join(f'<span aria-hidden="true">{Skeleton(40, 40)}{Skeleton(480, 16)}</span>' for _ in range(rows)) + '</div>'


def GradeTile(letter: str, bare: bool = False) -> str:
    """web ui/GradeTile: the letter on its tier colour, or `bare`, the letter in its colour with no tile."""
    return f'<span class="grade{" bare" if bare else ""} t{letter}" role="img" aria-label="Tier {letter}">{letter}</span>'

# ------------------------------------------------ primitives added by the planner round (2026-11-08; design/specs/2026-11-08-planner.md section 3A)
# Additive only: no earlier mock changes. `Page('workbench')` is L3 Workbench (3.7) with the split at 1280 (main | 340 rail;
# 400 from 1760), one column below 1280 where main and rail give their regions to one flow ordered by `data-order`.
# `TalentEditor` is the planner's editable trees on the page's one Plate: three tree columns split by hairlines (one tree
# behind tree tabs below 768), each a head (spec icon, tree name, points), the client's tree art behind a grid of talent
# buttons, the prerequisite arrows, a rank pill on every cell; the cell colours are the game's (TalentFrame): a talent you
# can spend in has a green ring and a green rank, a partial one green, a maxed one gold, a locked one grey and dimmed.
# `TalentTip` is the talent tooltip (Overlay kind 'tooltip'). `PointOrder` is the build's spend order in rows of ten levels
# with a socket per level. `Switch` is the shipped ui/Switch. `CodeField` is the paste box for a code a player copies out
# of the game (mono, the one face for literal codes). `ClassPicker` is the ControlBar's class choice: nine crests, the
# class names from 1280 (ruling P10) and in the phone stack.
PRIM_CSS += '''
.layout-wb{display:grid;grid-template-columns:minmax(0,1fr);gap:var(--section-gap) var(--gap);align-items:start}
.layout-wb>.main,.layout-wb>.rail{display:contents}
.layout-wb [data-order]{order:var(--o,0);min-width:0}
@media(min-width:1024px) and (max-width:1279px){.layout-wb{grid-template-columns:repeat(2,minmax(0,1fr))}.layout-wb [data-span]{grid-column:1/-1}}
@media(min-width:1280px){.layout-wb{grid-template-columns:minmax(0,1fr) 340px}.layout-wb>.main,.layout-wb>.rail{display:flex;flex-direction:column;gap:var(--section-gap);min-width:0}.layout-wb>.rail{gap:40px}}
@media(min-width:1760px){.layout-wb{grid-template-columns:minmax(0,1fr) var(--rail-w)}}
.tedit{--cell:44px;--cgap:8px}
@media(min-width:768px){.tedit{--cell:48px}}@media(min-width:1760px){.tedit{--cell:56px;--cgap:10px}}
.tedit-tabs{display:flex;border:1px solid var(--line);border-radius:4px;overflow:hidden;margin:12px 12px 0}
.tedit-tabs button{all:unset;box-sizing:border-box;flex:1;display:flex;flex-direction:column;align-items:center;gap:2px;min-height:60px;padding:8px 4px;border-left:1px solid var(--line-soft);cursor:pointer;color:var(--nav);font:600 12px/16px var(--font-body);text-align:center}
.tedit-tabs button:first-child{border-left:0}.tedit-tabs button>span{display:flex;align-items:center;gap:6px}
.tedit-tabs button b{font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}
.tedit-tabs button[aria-selected=true]{background:#e5b9551a;color:var(--gold);box-shadow:inset 0 -2px 0 var(--gold)}.tedit-tabs button[aria-selected=true] b{color:var(--gold)}
@media(min-width:768px){.tedit-tabs{display:none}}
.tedit-cols{display:grid;grid-template-columns:minmax(0,1fr)}
@media(min-width:768px){.tedit-cols{grid-template-columns:repeat(var(--n),minmax(0,1fr))}.tedit-tree+.tedit-tree{border-left:1px solid var(--line-soft)}}
@media(max-width:767px){.tedit-tree:not(.on){display:none}}
.tedit-tree{display:flex;flex-direction:column;align-items:center;gap:14px;padding:14px 12px 18px;min-width:0}
@media(min-width:1280px){.tedit-tree{padding:16px 20px 20px}}
.tedit-h{align-self:stretch;display:flex;align-items:center;gap:10px;min-width:0}
.tedit-h h2{font:700 15px/20px var(--font-body);color:var(--strong);margin:0;min-width:0}.tedit-h b{margin-left:auto;font:700 18px/24px var(--font-body);font-variant-numeric:tabular-nums;color:var(--gold)}
.tedit-h b.zero{color:var(--muted)}
@media(max-width:767px){.tedit-h{display:none}}
.tedit-g{position:relative;display:grid;grid-template-columns:repeat(var(--cols),var(--cell));gap:var(--cgap);padding:6px}
.tedit-g>.art{position:absolute;inset:0;width:100%;height:100%;object-fit:cover;border-radius:4px;opacity:.5;pointer-events:none}
.tedit-g>svg{position:absolute;inset:6px;width:calc(100% - 12px);height:calc(100% - 12px);pointer-events:none;overflow:visible}
.tedit-g>[role=row]{display:contents}
.tedit-g [role=gridcell]{position:relative;width:var(--cell);height:var(--cell)}
.tc2{all:unset;box-sizing:border-box;position:relative;display:block;width:var(--cell);height:var(--cell);border-radius:5px;cursor:pointer;background:#10141d}
.tc2 img,.tc2 .ic{display:block;width:100%!important;height:100%!important;border-radius:5px;border:0}
.tc2::after{content:"";position:absolute;inset:0;border-radius:5px;box-shadow:inset 0 0 0 2px var(--ring,var(--line))}
.tc2 .rk{position:absolute;right:-5px;bottom:-5px;z-index:1;padding:0 4px;border-radius:999px;border:1px solid var(--ring,var(--line));background:var(--bg);font:700 12px/14px var(--font-body);font-variant-numeric:tabular-nums;color:var(--rk,var(--muted));white-space:nowrap}
.tc2.avail{--ring:#4fae5c;--rk:var(--kill)}.tc2.part{--ring:var(--kill);--rk:var(--kill)}.tc2.max{--ring:var(--gold);--rk:var(--gold)}
.tc2.lock img,.tc2.lock .ic{filter:grayscale(1) brightness(.55)}.tc2.lock{cursor:default}
.tc2.diff::before{content:"";position:absolute;left:-4px;top:-4px;z-index:1;width:9px;height:9px;border-radius:50%;background:var(--gold);box-shadow:0 0 0 2px var(--bg)}
.tc2:hover:not(.lock),.tc2.hov{filter:brightness(1.15)}.tc2.hov::after{box-shadow:inset 0 0 0 2px var(--gold-hover)}
.tc2:focus-visible,.tc2.fv{outline:2px solid var(--gold);outline-offset:2px}
.tc2.act{transform:translateY(1px)}
.tc2.plan .rk{opacity:.6;border-style:dashed}
.tedit-h .trs{all:unset;cursor:pointer;font:600 13px/18px var(--font-body);color:var(--gold)}.tedit-h .trs:hover{color:var(--gold-hover)}
button.link{all:unset;cursor:pointer;font:600 14px/20px var(--font-body);color:var(--gold)}input.field{box-sizing:border-box;width:100%;font-family:var(--font-body)}input.field::placeholder{color:var(--muted)}
.porder-m{display:grid;grid-template-columns:120px minmax(0,1fr);column-gap:16px;padding-top:4px}
.porder-m .porder-s>span{height:auto;text-align:center;font-size:11px;line-height:14px;color:#5d6577;font-variant-numeric:tabular-nums}
@media(max-width:767px){.porder-m{grid-template-columns:minmax(0,1fr)}.porder-m>span:first-child{display:none}}
.txt-s{display:none}@media(max-width:767px){.txt-l{display:none}.txt-s{display:inline}}
.tedit-g path{fill:none;stroke-linecap:round;vector-effect:non-scaling-stroke}.tedit-g path.halo{stroke:var(--bg);stroke-width:5}.tedit-g path.on{stroke:var(--gold);stroke-width:2}.tedit-g path.off{stroke:#3a4256;stroke-width:2}
.tedit-f{display:flex;flex-wrap:wrap;align-items:center;gap:6px 16px;padding:12px 16px;border-top:1px solid var(--line-soft)}
.tedit-f .st2{flex-basis:100%;min-height:20px;font-size:14px;line-height:20px;color:var(--strong)}.tedit-f .st2.wipe{color:var(--wipe)}
.tedit-f .rs{margin-left:auto;display:flex;gap:10px;align-items:center;flex-wrap:wrap}
.ttip{position:relative;width:300px;max-width:100%;padding:12px 14px;background:var(--raised);border:1px solid var(--line);border-radius:6px;box-shadow:0 18px 40px #000a;display:flex;flex-direction:column;gap:6px}
.ttip b{font:700 15px/20px var(--font-body);color:var(--strong)}.ttip .r{font-size:13px;line-height:18px;color:var(--muted);font-variant-numeric:tabular-nums}
.ttip p{font-size:13px;line-height:19px;color:var(--text);margin:0}.ttip p.nx{color:var(--muted)}.ttip p.nx .t-label{display:block;margin-bottom:2px}
.ttip .req{color:var(--wipe);font-size:13px;line-height:18px}.ttip .bd{color:var(--gold);font:600 13px/18px var(--font-body)}
.porder{display:flex;flex-direction:column}
.porder-r{display:grid;grid-template-columns:120px minmax(0,1fr);align-items:center;column-gap:16px;padding:8px 0;border-bottom:1px solid var(--line-soft)}
.porder-r:last-child{border-bottom:0}.porder-r .t-small{font-variant-numeric:tabular-nums}
.porder-s{display:grid;grid-template-columns:repeat(10,36px);gap:6px}
.porder-s>span{position:relative;width:36px;height:36px}.porder-s img,.porder-s .ic{width:36px!important;height:36px!important}
.porder-s .em{display:block;width:36px;height:36px;border-radius:5px;border:1px dashed var(--line)}.porder-s .em.nx{border-color:var(--gold)}
.porder-s .em.past{border-style:solid;border-color:var(--line-soft);opacity:.5}
.porder-s .pl img,.porder-s .pl .ic{opacity:.55}.porder-s .pl::after{content:"";position:absolute;inset:0;border-radius:5px;box-shadow:inset 0 0 0 1px var(--line)}
.porder-s b{position:absolute;right:-4px;bottom:-5px;padding:0 3px;border-radius:999px;background:var(--bg);font:700 12px/14px var(--font-body);font-variant-numeric:tabular-nums;color:var(--muted)}
.porder-cap{display:flex;align-items:center;gap:10px;padding:10px 0;font-size:13px;line-height:18px;color:var(--gold);border-bottom:1px solid var(--line-soft)}
.porder-cap::before{content:"";flex:none;width:10px;height:2px;background:var(--gold)}
@media(max-width:767px){.porder-r{grid-template-columns:minmax(0,1fr);row-gap:6px}.porder-s{grid-template-columns:repeat(10,minmax(0,30px));gap:4px}.porder-r .porder-s>span,.porder-s img,.porder-s .ic,.porder-s .em{width:30px!important;height:30px!important}}
.swt{display:flex;align-items:center;justify-content:space-between;gap:12px;min-height:44px;font-size:14px;line-height:20px;color:var(--text)}
.swt i{flex:none;position:relative;width:40px;height:22px;border-radius:999px;background:var(--line)}.swt i::after{content:"";position:absolute;left:3px;top:3px;width:16px;height:16px;border-radius:50%;background:var(--nav)}
.swt.on i{background:var(--gold)}.swt.on i::after{left:21px;background:#17110a}.swt.off{opacity:.5}
.codef{display:block;width:100%;min-height:68px;padding:10px 12px;border:1px solid var(--line);border-radius:4px;background:#0a0d14;color:var(--strong);font:500 13px/20px ui-monospace,'SF Mono',Menlo,monospace;overflow-wrap:anywhere;white-space:pre-wrap}
.codef.ph{color:var(--muted)}.codef.bad{border-color:var(--wipe)}.codef.short{min-height:44px;display:flex;align-items:center}
.clsp a{padding:0 10px}.clsp a .lbl{display:none}
@media(min-width:1280px){.clsp a .lbl{display:inline}}@media(min-width:1760px){.clsp a{padding:0 12px}}
.bar-m .clsp a .lbl{display:inline}.bar-m .clsp a{flex:1 1 30%;justify-content:flex-start}
'''


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False, _prev=Page) -> str:
    """Page with layout 'workbench' (L3, split from 1280): main and rail; below 1280 one flow ordered by data-order."""
    if layout != 'workbench':
        return _prev(layout, main, rail, sticky_rail)
    return f'<div class="pc" style="padding-top:20px"><div class="layout-wb"><div class="main">{main}</div><aside class="rail">{rail}</aside></div></div>'


def Ordered(html: str, order: int, span: bool = False) -> str:
    """A region's place in the one-column flow of Page('workbench') below 1280 (and full width from 1024 when `span`)."""
    return f'<div data-order style="--o:{order}"{" data-span" if span else ""}>{html}</div>'


def TalentCell2(icon_html: str, name: str, rank: int, mx: int, state: str, differs: bool = False, extra: str = '') -> str:
    """One talent button. state: avail | part | max | lock (lib/planner/styles cellState: available | filled | maxed | locked)."""
    cls = f'tc2 {state}' + (' diff' if differs else '') + (f' {extra}' if extra else '')
    dis = ' aria-disabled="true"' if state == 'lock' else ''
    return f'<button type="button" class="{cls}" aria-label="{escape(name)}, rank {rank} of {mx}"{dis}>{icon_html}<span class="rk" aria-hidden="true">{rank}/{mx}</span></button>'


def TextFor(long: str, short: str) -> str:
    """One sentence in two lengths: the long one from 768, the short one on a phone (one line)."""
    return f'<span class="txt-l">{long}</span><span class="txt-s">{short}</span>'


def TalentEditor(trees: list[dict], active: int = 0, foot: str = '', label: str = 'Talent trees') -> str:
    """trees: {name, icon_html, points, art_src, cols, tiers, viewbox, connectors [(d, met)], cells [(tier, col, cell_html)]}.
    The Plate's content: tree tabs (below 768), the tree columns, then the foot (status line, source line, Reset)."""
    tabs = ''.join(f'<button type="button" role="tab" aria-selected="{"true" if i == active else "false"}"><span>{t["icon_html"]}<b>{t["points"]}</b></span>{escape(t["name"])}</button>'
                   for i, t in enumerate(trees))
    cols = ''
    for i, t in enumerate(trees):
        rows = ''
        for tier in range(t['tiers']):
            row = ''
            for col in range(t['cols']):
                c = next((x for x in t['cells'] if x[0] == tier and x[1] == col), None)
                row += f'<span role="gridcell">{c[2] if c else ""}</span>'
            rows += f'<div role="row">{row}</div>'
        paths = ''.join(f'<path class="halo" d="{d}"/><path class="{"on" if met else "off"}" d="{d}"/>' for d, met in t['connectors'])
        art = f'<img class="art" src="{t["art_src"]}" alt="" aria-hidden="true">' if t.get('art_src') else ''
        pts = f'<b class="num{" zero" if not t["points"] else ""}">{t["points"]}</b>'
        rs = t.get('reset', '')
        cols += (f'<div class="tedit-tree{" on" if i == active else ""}" role="tabpanel" aria-label="{escape(t["name"])}"><div class="tedit-h">{t["icon_html"]}<h2>{escape(t["name"])}</h2>{pts}{rs}</div>'
                 f'<div class="tedit-g" role="grid" aria-label="{escape(t["name"])} talents" style="--cols:{t["cols"]}">{art}<svg viewBox="{t["viewbox"]}" preserveAspectRatio="none" aria-hidden="true">{paths}</svg>{rows}</div></div>')
    f = f'<div class="tedit-f">{foot}</div>' if foot else ''
    return Plate(f'<div class="tedit"><div class="tedit-tabs" role="tablist" aria-label="Talent trees">{tabs}</div><div class="tedit-cols" style="--n:{len(trees)}">{cols}</div>{f}</div>', label)


def TalentTip(name: str, rank: int, mx: int, now: str, nxt: str, req: str = '', band: str = '', next_label: str = 'Next rank') -> str:
    n = f'<p>{escape(now)}</p>' if now else ''
    x = f'<p class="nx"><span class="t-label">{next_label}</span>{escape(nxt)}</p>' if nxt else ''
    r = f'<span class="req">{escape(req)}</span>' if req else ''
    bd = f'<span class="bd">{escape(band)}</span>' if band else ''
    return f'<div class="ttip" role="tooltip"><b>{escape(name)}</b><span class="r">Rank {rank} of {mx}</span>{n}{x}{r}{bd}</div>'


def PointOrder(rows: list[tuple[str, list[str]]], cap_after: int = -1, cap_text: str = '', markers: list[str] | None = None) -> str:
    """rows: (label, slots html). A slot: an icon with its rank, an empty socket, the next socket (gold). `cap_after`: the row
    index after which the character's level line sits (points past it are planned ahead). A faint 0 to 9 marker sits over
    the first row: the last digit of each socket's level."""
    heads = markers if markers is not None else [str(d) for d in range(10)]  # additive: the built `markers`
    out = '<div class="porder-m" aria-hidden="true"><span></span><div class="porder-s">' + ''.join(f'<span>{d}</span>' for d in heads) + '</div></div>'
    for i, (lab, slots) in enumerate(rows):
        out += f'<div class="porder-r"><span class="t-small">{lab}</span><div class="porder-s" role="list">{"".join(slots)}</div></div>'
        if i == cap_after and cap_text:
            out += f'<p class="porder-cap">{cap_text}</p>'
    return f'<div class="porder">{out}</div>'


def OrderSlot(icon_html: str = '', name: str = '', rank: str = '', kind: str = '') -> str:
    """kind: '' (spent), 'pl' (spent past the character's level), 'em' (empty), 'nx' (the next point), 'past' (beyond 60: none)."""
    if kind in ('em', 'nx', 'past'):
        return f'<span role="listitem" aria-label="{escape(name)}" title="{escape(name)}"><span class="em {kind if kind != "em" else ""}"></span></span>'
    return f'<span role="listitem" class="{kind}" aria-label="{escape(name)}" title="{escape(name)}">{icon_html}<b>{rank}</b></span>'


def Switch(label: str, on: bool = True, disabled: bool = False) -> str:
    return f'<label class="swt{" on" if on else ""}{" off" if disabled else ""}"><span>{label}</span><i role="switch" aria-checked="{"true" if on else "false"}"{" aria-disabled=\"true\"" if disabled else ""} tabindex="0"></i></label>'


def CodeField(text: str, placeholder: bool = False, bad: bool = False, label: str = '', short: bool = False) -> str:
    cls = 'codef' + (' ph' if placeholder else '') + (' bad' if bad else '') + (' short' if short else '')
    return f'<span class="{cls}" role="textbox" aria-multiline="{"false" if short else "true"}" aria-label="{escape(label)}">{escape(text)}</span>'


def ClassPicker(classes: list[tuple[str, str, bool]]) -> str:
    """The ControlBar's class choice: (crest_html, name, current). Crests always; names from 1280 and in the phone stack."""
    return seg([(f'<span class="lbl">{n}</span>', on, None, c) for c, n, on in classes], 'clsp', 'Class').replace('<a href="#" role="radio"', '<a href="#" role="radio" title=""', 0)

# ------------------------------------------------ primitives added by the pages round (2026-11-08; design/specs/2026-11-08-pages.md section 3A)
# Additive only: no earlier mock changes (Page is redefined as a superset; at its defaults it emits exactly what it did).
# `Page('reading', toc_from=1024)` is a document with no lanes (the legal and About pages): the contents rail at the left from
# 1024 instead of 1760; `toc_from=0` is a document with no contents (About). `Prose` gains a ruled table (the roadmap) and
# inline code (a cookie name) in the one mono face. `Dated` is one entry of a dated log (the changelog): the date (a permalink)
# and its tags in a 120 px column, the entry (title 18/24, body 16/26 at 720) and its lane (sources, a caveat) side by side
# from 1280, stacked below. `DatedGroup` is a reading section of entries (a month). `CharacterAnswer` is the guides' signed-in
# answer drawn as a primitive (the strings are lib/account/answer-copy.ts and guidesCopy; guides gap G2), reused by the 404.
PRIM_CSS += '''
@media(min-width:1024px){.layout-reading.toc-lg{grid-template-columns:var(--toc-w) minmax(0,1fr)}.layout-reading.toc-lg>.toc{display:block;position:sticky;top:24px}}
.layout-reading.toc-none>.toc{display:none}
@media(min-width:1760px){.layout-reading.toc-none{grid-template-columns:minmax(0,1fr)}.layout-reading.toc-none>.toc{display:none}}
.rsec{scroll-margin-top:24px}
.layout-reading.toc-lg .rsec>.sec-h,.layout-reading.toc-none .rsec>.sec-h{max-width:var(--measure-reading)}
.rprose table{width:100%;border-collapse:collapse;margin:0 0 18px;font-size:15px;line-height:22px}
.rprose th{text-align:left;font:700 12px/16px var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted);padding:0 16px 8px 0;border-bottom:1px solid var(--line-soft)}
.rprose td{padding:10px 16px 10px 0;border-bottom:1px solid var(--line-soft);vertical-align:top}
.rprose td:first-child{width:1%;white-space:nowrap;font-weight:600;color:var(--strong)}
@media(max-width:767px){.rprose thead{display:none}.rprose tr{display:grid;gap:2px;padding:10px 0;border-bottom:1px solid var(--line-soft)}.rprose td{padding:0;border:0}.rprose td:first-child{width:auto;white-space:normal}}
.rprose code{font:600 14px/20px ui-monospace,'SF Mono',Menlo,monospace;color:var(--strong)}
.dated{display:grid;grid-template-columns:minmax(0,1fr);gap:10px var(--gap);padding:24px 0;border-bottom:1px solid var(--line-soft);scroll-margin-top:72px}
.dated:first-child{padding-top:20px}.dated:last-child{border-bottom:0}
.dated-w{display:flex;flex-wrap:wrap;align-items:center;gap:6px 10px}
.dated-w .d{font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}.dated-w .d:hover{color:var(--gold-hover)}
.dated-b{min-width:0;max-width:var(--measure-reading)}
.dated-b>h3{font:700 18px/24px var(--font-body);color:var(--strong);margin:0 0 8px}
.dated-n{font-size:16px;line-height:26px;color:var(--text);margin:0}
.dated-l{min-width:0;display:flex;flex-direction:column;gap:10px}
.dated-c{font-size:13px;line-height:20px;color:var(--muted);margin:0}
@media(min-width:1280px){.dated{grid-template-columns:120px var(--measure-reading) minmax(0,1fr)}.dated-w{flex-direction:column;align-items:flex-start;gap:8px}.dated-l{padding-top:2px}}
.dated-g>.sec-h>.t-small{margin-left:auto}
.cans-row{display:flex;align-items:center;gap:12px}.cans-who{display:flex;flex-direction:column;min-width:0}.cans-name{font:700 18px/22px var(--font-body)}
.cans-links{display:flex;gap:4px 16px;flex-wrap:wrap}
@media(min-width:900px){.cans-row{justify-content:flex-end}.cans-who{align-items:flex-start;text-align:left}.cans-links{justify-content:flex-end}}
.cans-links .cans-sm{font-size:13px;line-height:20px;align-self:center}
.dated-f{padding:0;scroll-margin-top:72px;border-bottom:1px solid var(--line-soft)}.dated-f:last-child{border-bottom:0}
.dated-f>summary{list-style:none;cursor:pointer;display:grid;grid-template-columns:auto minmax(0,1fr);grid-template-areas:'d g' 'h h';align-items:center;gap:4px 10px;padding:14px 0;min-height:44px}
.dated-f>summary>.d{grid-area:d}.dated-f>summary>.dated-k{grid-area:g;display:flex;gap:6px;align-items:center}.dated-f>summary>h3{grid-area:h}
.dated-f>summary::-webkit-details-marker{display:none}
.dated-f>summary h3{font:700 16px/22px var(--font-body);color:var(--strong);margin:0;display:flex;gap:8px;align-items:baseline}
.dated-f>summary h3::before{content:'\\25B8';font-size:11px;color:var(--muted);flex:none}.dated-f[open]>summary h3::before{content:'\\25BE'}
.dated-f>summary:hover h3{color:var(--gold-hover)}
.dated-f>summary>.d{font:600 14px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--muted)}
.dated-o{display:grid;grid-template-columns:minmax(0,1fr);gap:10px var(--gap);padding:0 0 24px}
@media(min-width:1280px){.dated-f>summary,.dated-o{grid-template-columns:120px var(--measure-reading) minmax(0,1fr);column-gap:var(--gap)}.dated-f>summary{grid-template-areas:'d h g'}.dated-o>.dated-b{grid-column:2}}
.dated-g>.sec-h+.dated-list>.dated-f:first-child>summary{padding-top:20px}
'''

_Page3 = Page


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False, toc_from: int = 1760, toc_label: str = '') -> str:
    """Page with `toc_from` on the reading layout: 1760 (guides, the changelog), 1024 (a document: legal, Data sources), 0 (no contents)."""
    html = _Page3(layout, main, rail, sticky_rail)
    if layout != 'reading':
        return html
    if toc_from != 1760:
        html = html.replace('class="layout-reading"', f'class="layout-reading toc-{"lg" if toc_from == 1024 else "none"}"', 1)
    if toc_label:
        html = html.replace('aria-label="Sections in this guide"', f'aria-label="{escape(toc_label)}"', 1)
    return html


def Dated(anchor: str, iso: str, date: str, tags: str, title: str, body: str, lane: str = '', folded: bool = False) -> str:
    """One entry of a dated log: the date (a permalink to the entry) over its tags, the entry, its lane (sources, a caveat).
    folded (ruling CL1, entries before launch): a title row (date, tags, title) that opens in place onto the body and lane."""
    lane_html = f'<div class="dated-l">{lane}</div>' if lane else ''
    if folded:
        return (f'<details class="dated-f" id="{anchor}"><summary><time class="d" datetime="{iso}">{date}</time>'
                f'<h3>{title}</h3><span class="dated-k">{tags}</span></summary><div class="dated-o"><div class="dated-b">{body}</div>{lane_html}</div></details>')
    return (f'<article class="dated" id="{anchor}"><div class="dated-w"><a class="d" href="#{anchor}"><time datetime="{iso}">{date}</time></a>{tags}</div>'
            f'<div class="dated-b"><h3>{title}</h3>{body}</div>{lane_html}</article>')


def DatedGroup(title: str, meta: str, anchor: str, entries: str) -> str:
    """A reading section of dated entries (a month): the reading head (22/28) with its count at the right."""
    return (f'<section class="rsec dated-g" data-region="{region_slug(title)}" id="{anchor}"><div class="sec-h"><h2>{title}</h2>'
            f'<span class="t-small">{meta}</span></div><div class="dated-list">{entries}</div></section>')


def CharacterAnswer(label: str, state: str, crest: str = '', name: str = '', color: str = '', line: str = '', links: list[str] | None = None,
                    text: str = '', action: str = '') -> str:
    """PageHeader.answer for 'your character'. state: who | loading | message | error (message with role="alert")."""
    lab = f'<span class="sr-only">{label}</span>'   # the label is a screen reader's (ui/CharacterAnswer)
    if state == 'loading':
        return (f'<div class="ans-box" aria-busy="true">{lab}<span class="cans-row">{Skeleton(40, 40, True)}<span class="cans-who">{Skeleton(150, 22)}'
                f'{Skeleton(190, 16)}</span></span>{Skeleton(220, 20)}</div>')
    if state in ('message', 'error'):
        role = ' role="alert"' if state == 'error' else ''
        return f'<div class="ans-box">{lab}<span class="t-small"{role}>{text}</span>{action}</div>'
    who = (f'<span class="cans-row">{crest}<span class="cans-who"><b class="cans-name" style="color:{color}">{name}</b>'
           f'<span class="t-small">{line}</span></span></span>')
    return f'<div class="ans-box">{lab}{who}<span class="cans-links">{"".join(links or [])}</span></div>'


# ------------------------------------------------ primitives added by the premium and account round (2026-11-08; design/specs/2026-11-08-premium-account.md section 3A)
# Additive only. `Columns` is PlateColumns without the Plate: labelled columns split by hairlines, open on S0 (the guild plans
# under the player's Plate, the How it works and Questions grids). `Button(tone='danger')` is the shipped button('danger'):
# a destructive confirm (Leave guild, Revoke, Unlink), wipe text and a wipe hairline, never filled. `PlanHead` is the price
# block a plan column opens with: the name in Barlow 700 22/28, a Tag, the audience, the price Figure (gold on the one plan
# the page sells, strong elsewhere) and its two small lines; the column's one Button sits under it.
PRIM_CSS += '''
.pcols.open .pcol{padding:16px 0 0}.pcols.open .pcol+.pcol{padding-left:0}
@media(min-width:1024px){.pcols.open .pcol{padding:4px 24px 0}.pcols.open .pcol:first-child{padding-left:0}.pcols.open .pcol:last-child{padding-right:0}}
.btn.danger{color:var(--wipe);border-color:#f0736b66}.btn.danger:hover{border-color:var(--wipe);color:#ff8f88}
@media(max-width:1023px){.pcols.swap{display:flex;flex-direction:column}.pcols.swap>.pcol:nth-child(2){order:-1;border-top:0}.pcols.swap>.pcol:first-child{border-top:1px solid var(--line-soft)}}
.plh{display:flex;flex-direction:column;gap:4px;min-width:0}.plh .nm{display:flex;align-items:center;gap:10px;font:400 24px/30px var(--font-title);letter-spacing:-.015em;color:var(--strong)}.plh .nm>*{font-family:var(--font-body);letter-spacing:normal}
.plh .au{font-size:14px;line-height:20px;color:var(--muted)}.plh .fig{margin-top:8px}.plh .pl2{font-size:13px;line-height:20px;color:var(--muted)}.plh .pl2.gold{color:var(--gold)}
'''


def Columns(cols: list[tuple[str, str, str, str]], open_: bool = True, second_first: bool = False) -> str:
    """PlateColumns' columns: (label, meta, body, foot), hairlines between, stacked below 1024. open_ (default): no Plate, no
    side padding; open_=False: the padded columns a caller puts in its own Plate when the columns carry their own heads.
    second_first: below 1024 the second column stacks first (the plan a page sells leads on a phone)."""
    inner = ''.join(f'<div class="pcol">{f"<div class=pcol-h><h3 class=t-label>{lab}</h3>{meta}</div>" if lab or meta else ""}{body}'
                    + (f'<div class="pcol-f">{foot}</div>' if foot else '') + '</div>' for lab, meta, body, foot in cols)
    return f'<div class="pcols{" open" if open_ else ""}{" swap" if second_first else ""}" style="--n:{len(cols)}">{inner}</div>'


_Button0 = Button


def Button(label: str, primary: bool = False, small: bool = False, tone: str = '') -> str:
    """The system Button; tone 'danger' is the destructive confirm (button('danger'))."""
    html = _Button0(label, primary, small)
    return html.replace('class="btn', 'class="btn danger', 1) if tone == 'danger' else html


def PlanHead(name: str, audience: str, value: str, unit: str, lines: list[tuple[str, str]], tag: str = '', gold: bool = False, level: int = 3) -> str:
    """A plan column's head: name (+ Tag), audience, the price Figure, then (text, tone) small lines."""
    sm = ''.join(f'<span class="pl2{" " + t if t else ""}">{x}</span>' for x, t in lines)
    return (f'<div class="plh"><h{level} class="nm">{name}{tag}</h{level}><span class="au">{audience}</span>'
            f'{Figure(value, unit=unit, tone="gold" if gold else "", size="big")}{sm}</div>')


# ------------------------------------------------ the guild parity round (2026-11-07): the shipped markup of Button, Meter, Tally, Grid and
# Rows, so a mock's blocks read as the built ones (web/src/components/ui: Button is a <button> when it acts, Meter a labelled meter,
# Tally a list, Grid a table with a caption and an optional lead per row, Rows a list when the built rows are <li>s). RowGroup is a
# Plate's caps group label over its rows (the raid planner's Tanks, Healers ... groups).
PRIM_CSS += '''
button.btn{cursor:pointer;-webkit-appearance:none;appearance:none}button.btn.sm{padding:0 18px}
.row .foot .pill{height:22px}
.rprose.note{font-size:14px;line-height:1.6}
.btnrow{display:flex;flex-wrap:wrap;align-items:center;gap:8px;padding-top:12px}
.btn.ghost{border-color:transparent;color:var(--gold);text-transform:none;letter-spacing:.01em;font-size:13px;padding:0 18px}
.rows.open .row.cur{padding-inline:16px}
section>.sec-h+ul.rows,section>.sec-h+ul,section>.sec-h+table,section>.sec-h+.rprose{margin-top:12px}
ul.rows{list-style:none;margin:0;padding:0}ul.rows>li:not(:last-child)>.row{border-bottom:1px solid var(--line-soft)}
.rgrp{padding-bottom:4px}.rgrp>h3{display:flex;gap:8px;align-items:baseline;padding:12px 16px 4px;font:700 12px/16px var(--font-body);letter-spacing:.1em;text-transform:uppercase;color:var(--muted)}
.rgrp>h3 span{font-variant-numeric:tabular-nums}
.tally{list-style:none;margin:0;padding:0}.tally .t .note{font-size:12px;line-height:16px;color:var(--muted)}
table.gridm{display:table;width:100%;border-collapse:collapse}table.gridm caption{position:absolute;width:1px;height:1px;overflow:hidden;clip:rect(0 0 0 0)}
table.gridm thead th{height:33px;box-sizing:border-box;padding-block:8px}table.gridm th.h{text-align:left}table.gridm th.h.hc{text-align:center}table.gridm th.l{text-align:left;font-weight:400}
table.gridm .ld{display:flex;gap:8px;align-items:center;min-width:0}table.gridm .ld img{width:24px;height:24px;flex:none}
'''

_Button1 = Button


def Button(label: str, primary: bool = False, small: bool = False, tone: str = '', act: bool = False) -> str:
    """act: the built Button that does something (a <button>), not a link styled as one. tone 'ghost' is the quiet in-row Edit."""
    html = _Button1(label, primary, small, '' if tone == 'ghost' else tone)
    if tone == 'ghost':
        html = html.replace('class="btn', 'class="btn ghost', 1)
    return html.replace('<a ', '<button type="button" ', 1).replace(' href="#"', '', 1).replace('</a>', '</button>') if act else html


_Meter0 = Meter


def Meter(frac: float, tone: str = '', label: str = '') -> str:
    """label: the built Meter's accessible name (role meter); without one it is decoration."""
    html = _Meter0(frac, tone)
    return html.replace('<div class="meter', f'<div role="meter" aria-label="{label}" aria-valuenow="{round(max(0, min(1, frac)) * 100)}" aria-valuemin="0" aria-valuemax="100" class="meter', 1) if label else html


def Tally(items: list[tuple]) -> str:
    """Counts against a target, one labelled Meter each, as a list. items: (lead_html, label, have, want[, note[, tone]]); a tone
    (ok|warn|miss|none) overrides the target's, as the built Tally's item tone (`none`: a share, not a goal, reads muted)."""
    out = ''
    for lead, label, have, want, *rest in items:
        st = rest[1] if len(rest) > 1 and rest[1] else ('ok' if have >= want else ('warn' if have >= want * 0.75 else 'miss'))
        note = f'<span class="note">{rest[0]}</span>' if rest and rest[0] else ''
        bar = 'kill' if st == 'ok' else ('muted' if st == 'none' else '')
        out += (f'<li class="t"><div class="top">{lead}<b>{label}</b><span class="n {st}">{have}/{want}</span></div>'
                f'{Meter(have / want if want else 1, bar, f"{label}: {have} of {want}")}{note}</li>')
    return f'<ul class="tally">{out}</ul>'


def Grid(cols: list[str], rows: list[tuple], caption: str = '', label_header: str = '') -> str:
    """A coverage matrix as the built table. rows: (label, sub, [(state, text)][, lead_html]) with state ok|warn|miss|none."""
    head = f'<th scope="col" class="h">{label_header or ""}</th>' + ''.join(f'<th scope="col" class="h hc">{c}</th>' for c in cols)
    body = ''
    for lab, sub, cells, *rest in rows:
        lead = rest[0] if rest else ''
        text = f'<span><span class="ttl" style="display:block">{lab}</span>' + (f'<span class="sub" style="display:block">{sub}</span>' if sub else '') + '</span>'
        body += (f'<tr><th scope="row" class="l"><span class="ld">{lead}{text}</span></th>'
                 + ''.join(f'<td class="c {st}">{tx}</td>' for st, tx in cells) + '</tr>')
    cap = f'<caption>{caption}</caption>' if caption else ''
    return f'<table class="gridm" style="--n:{len(cols)}">{cap}<thead><tr>{head}</tr></thead><tbody>{body}</tbody></table>'


_Rows0 = Rows


def Rows(rows: list[str], open_: bool = False, list_: bool = False) -> str:
    """list_: the built rows are list items (a <ul> of <li>s)."""
    if not list_:
        return _Rows0(rows, open_)
    return f'<ul class="rows{" open" if open_ else ""}">' + ''.join(f'<li>{r}</li>' for r in rows) + '</ul>'


def ButtonRow(buttons: str) -> str:
    """Buttons that act on the block above, under it in one wrapping line (the assignments' Copy as raid warning and Post to Discord)."""
    return f'<div class="btnrow">{buttons}</div>'


def RowGroup(label: str, count: int, rows: list[str]) -> str:
    """A group inside a Plate: a caps label with its count (a level 3 heading, never a Section) over its rows as a list."""
    return f'<div class="rgrp"><h3>{label} <span>{count}</span></h3>{Rows(rows, list_=True)}</div>'


PRIM_CSS += '''
.sgr .st{min-height:20px}
.sgr{border-bottom:1px solid var(--line-soft);padding-top:12px;list-style:none}.sgr:last-child{border-bottom:0}
.sgr .g{display:grid;grid-template-columns:minmax(180px,1fr) auto minmax(200px,1.4fr);align-items:center;gap:4px 16px;min-height:68px}
.sgr .who{display:flex;gap:12px;align-items:center;min-width:0}.sgr .who p{margin:0}.sgr .nm{font:700 14px/20px var(--font-body)}.sgr .ws{font-size:12px;color:var(--muted)}
.sgr .g.spec{grid-template-columns:minmax(180px,1fr) auto 220px minmax(200px,1.4fr)}
.sgr input{height:44px;width:100%;min-width:0;border:1px solid var(--line-warm);border-radius:4px;background:var(--raised);color:var(--text);padding:0 12px;font-size:13px}
.sgr .seg a{height:44px}
'''

_Section0 = Section


def Section(title: str, body: str, action: str = '', meta: str = '', named: bool = False, boxed: bool = False) -> str:
    """named: the built Section with an id, a landmark named by its heading (role region). boxed: the body on the panel."""
    html = _Section0(title, body, action, meta, boxed=boxed)
    if not named:
        return html
    sid = region_slug(title)
    return html.replace('<section ', f'<section id="{sid}" aria-labelledby="{sid}-title" ', 1).replace('<h2 ', f'<h2 id="{sid}-title" ', 1)


def SignupRow(crest: str, name: str, color: str, who: str, segments: str, spec: str = '', note: str = 'Note for the officers (optional)') -> str:
    """The viewer's own sign-up (guild/SignupRow): crest, name and spec line, In / Late / Bench / Out, the spec pick when asked, the note."""
    pick = f'<span>{spec}</span>' if spec else ''
    return (f'<li class="sgr"><div class="g{" spec" if spec else ""}"><div class="who">{crest}<div><p class="nm" style="color:{color}">{name}</p><p class="ws">{who}</p></div></div>'
            f'{segments}{pick}<input type="text" aria-label="{note}" placeholder="{note}"></div><div class="st"></div></li>')


# Lane A (the same round): in-page tabs on one URL are a tablist, a spell icon carries its name, a disclosure Button its state.
_TabStrip0 = TabStrip


def TabStrip(tabs: list[str], active: str, counts: dict | None = None, tablist: bool = False) -> str:
    """tablist: in-page tabs on one URL (the guild's hash tabs) are a tablist of tabs, as the built TabStrip's `tablist` mode."""
    html = _TabStrip0(tabs, active, counts)
    if not tablist:
        return html
    html = html.replace('<nav class="strip-tabs ph-tabs">', '<div class="strip-tabs ph-tabs" role="tablist">', 1).replace('</nav>', '</div>')
    return re.sub(r'<a href="#" class="(cur)?"', lambda m: f'<a href="#" role="tab" aria-selected="{"true" if m.group(1) else "false"}" class="{m.group(1) or ""}"', html)


def SpellIcon(name: str, stem: str) -> str:
    """A spell's 24 px icon with its name as the title (the built SpellIcon)."""
    return icon(stem, 24).replace('class="ic"', 'class="sicon"', 1).replace('alt=""', f'alt="" title="{name}"', 1)


def DisclosureButton(label: str, expanded: bool = False) -> str:
    """A gold text button that opens a block under it in place (Export history): the built <button aria-expanded>."""
    return f'<button type="button" class="link tbtn" aria-expanded="{"true" if expanded else "false"}">{label}</button>'


def VoteButtons(yes: int, no: int, voters: int) -> str:
    """The built Votes an officer can press: each count a Tag inside a <button>, then who votes."""
    pill = ''.join(f'<button type="button" class="vote">{Tag(f"{n} {w}", tone if n else "muted")}</button>' for n, w, tone in ((yes, 'yes', 'kill'), (no, 'no', 'wipe')))
    return f'<span class="votes">{pill}<span class="t-small">of {voters} officers</span></span>'


PRIM_CSS += '''
.sicon{width:24px;height:24px;border:1px solid var(--line-soft);border-radius:3px;flex:none}
.strip-tabs[role=tablist] a{text-decoration:none}
button.tbtn{background:none;border:0;padding:0;font:600 14px/20px var(--font-body);cursor:pointer}
button.vote{background:none;border:0;padding:0;cursor:pointer}.votes{display:inline-flex;flex-wrap:wrap;align-items:center;gap:6px}
'''
# ------------------------------------------------ primitives added by the guild report round (2026-11-07; guild parity, lane C)
# Additive only. `Toggle` is a Row title that opens the row's detail (the shipped Row
# titleContent button, aria-expanded). `Inspector` is the shipped ui/Inspector at its side placement: an aside named
# "<selection> detail", 320 px, a hairline at its left, an optional head (title, sub) and caps-labelled sections; `Entries`
# is a section's ordered list (a line and a detail line each). `Live` is the polite status a copy button speaks through
# (empty until it speaks). `Split` lays a Plate's rows beside its Inspector.
PRIM_CSS += '''
button.tgl{all:unset;cursor:pointer;font:inherit;color:inherit}
.split-insp{display:flex;min-width:0}.split-insp>.rows{flex:1;min-width:0}
.insp{width:320px;flex:none;border-left:1px solid var(--line);align-self:start;min-width:0}
.insp-h,.insp-s{padding:14px 16px}.insp-s+.insp-s,.insp-h+.insp-s{border-top:1px solid var(--line-soft)}
.insp-h p{margin:0}.insp-h .t{font:700 15px/20px var(--font-body);color:var(--strong)}.insp-h .s{font-size:13px;line-height:18px;color:var(--muted)}
.insp-s h3{margin:0;font:700 12px/16px var(--font-body);letter-spacing:.1em;text-transform:uppercase;color:var(--muted)}
.insp-l{display:flex;flex-direction:column;gap:16px;margin:4px 0 0;padding:0;list-style:none;font-size:14px;line-height:22px}
.insp-l li{display:flex;flex-direction:column;gap:4px;min-width:0}.insp-l p{margin:0}
.live{position:absolute;width:1px;height:1px;overflow:hidden;clip:rect(0 0 0 0);white-space:nowrap}
'''


def Toggle(label: str, expanded: bool = False) -> str:
    return f'<button type="button" class="tgl" aria-expanded="{"true" if expanded else "false"}">{label}</button>'


def Entries(items: list[tuple[str, str]]) -> str:
    return '<ol class="insp-l">' + ''.join(f'<li><p>{line}</p><div>{detail}</div></li>' for line, detail in items) + '</ol>'


def Inspector(selection: str, sections: list[tuple[str, str]], title: str = '', sub: str = '') -> str:
    head = f'<div class="insp-h"><p class="t">{title}</p><p class="s">{sub}</p></div>' if title else ''
    body = ''.join(f'<section class="insp-s"><h3>{label}</h3>{html}</section>' for label, html in sections)
    return f'<aside class="insp" data-region="{REGION["inspector"]}" aria-label="{escape(selection)} detail">{head}{body}</aside>'


def Live() -> str:
    return '<span class="live" role="status" aria-live="polite"></span>'


def Split(rows: str, inspector: str) -> str:
    return f'<div class="split-insp">{rows}{inspector}</div>'


# ------------------------------------------------ Frame (round 5; web ui/Frame.svelte): the featured panel, two or three a page
# The panel (raised ground, black 1 px edge, 10 px radius, the warm outer ring) with the cream line inset 6 px; `ornate` marks the
# page's one lead panel (the built corners are decoration on the inset line, no box of their own). Padding 22/24, 18/16 below 768.
PRIM_CSS += '''
.frame{position:relative;min-width:0;overflow:hidden;padding:22px 24px;background:var(--raised);border:1px solid #000;border-radius:10px;
box-shadow:0 0 0 1px rgba(130,112,91,.55);--muted:#8b8784}
.frame::before{content:"";position:absolute;inset:6px;border:1px solid rgba(243,238,226,.42);border-radius:6px;pointer-events:none}
.frame.plain::before{content:none}
@media(max-width:767px){.frame{padding:18px 16px}}
'''


def Frame(inner: str, ornate: bool = False, region: str = '', plain: bool = False) -> str:
    """web ui/Frame: `region` is the parity region it carries, when it is one (REGION values); `plain` drops the inset line."""
    r = f' data-region="{region}"' if region else ''
    mod = " ornate" if ornate else " plain" if plain else ""
    return f'<div class="frame{mod}"{r}>{inner}</div>'


# The panel form (web ui/Section `panel`, `frame`, `flush`; ui/Page `panels`): the logs and report mocks' prims.panel, the
# heading row inside one Plate with the body. `frame` is the Plate's ('plain', 'framed', 'lead'); `flush` runs the body
# edge to edge (a table that pads its own cells). Page(panels=True) sets them 22 px apart, the columns 28.
PRIM_CSS += '''
.plate>.sec-h.pn-h{padding:16px 20px 10px;border-bottom:0}.pn-b{padding:0 20px 16px}.pn-b.flush{padding:0 0 6px}
.plate.framed{position:relative}.plate.framed::before{content:"";position:absolute;inset:6px;border:1px solid rgba(243,238,226,.42);border-radius:6px;pointer-events:none}
.layout-split.panels{column-gap:28px}.layout-split.panels>.main,.layout-split.panels>.rail{gap:22px}
@media(max-width:767px){.plate>.sec-h.pn-h{padding:14px 14px 8px}.pn-b{padding:0 14px 14px}}
'''

_Section1 = Section


def _named(html: str, title: str) -> str:
    sid = region_slug(title)
    return html.replace('<section ', f'<section id="{sid}" aria-labelledby="{sid}-title" ', 1).replace('<h2 ', f'<h2 id="{sid}-title" ', 1)


def Section(title: str, body: str, action: str = '', meta: str = '', named: bool = False, boxed: bool = False,
            panel: bool = False, frame: str = 'plain', flush: bool = False) -> str:
    """panel: the heading row inside one Plate (its frame) with the body, as the built Section's `panel`."""
    if not panel:
        if not boxed:
            return _Section1(title, body, action, meta, named)
        html = _Section0(title, body, action, meta, boxed=True)
        return _named(html, title) if named else html
    a = f'<span class="act">{action}</span>' if action else ''
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    cls = {'plain': '', 'framed': ' framed', 'lead': ' framed lead'}[frame]
    html = (f'<section data-region="{region_slug(title)}"><div class="plate{cls}" data-region="{REGION["plate"]}">'
            f'<div class="sec-h pn-h"><h2 class="t-h3">{title}</h2>{m}{a}</div>'
            f'<div class="pn-b{" flush" if flush else ""}">{body}</div></div></section>')
    return _named(html, title) if named else html


_Page1 = Page


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False, panels: bool = False, **kwargs: object) -> str:
    """panels: main and rail hold Section panels (web ui/Page `panels`); every other keyword goes to the Page before."""
    html = _Page1(layout, main, rail, sticky_rail, **kwargs)
    return html.replace('class="layout-split"', 'class="layout-split panels"', 1) if panels else html


# ------------------------------------------------ Grid edge (web ui/Grid `edge`): the grid on a flush panel, the label column
# taking the slack, the outer cells carrying the panel's 20 px inset (14 on a phone), the figure columns tight at the right.
PRIM_CSS += '''
table.gridm.edge :is(th,td){padding-inline:10px}table.gridm.edge :is(th,td):first-child{width:100%;padding-left:20px}
table.gridm.edge :is(th,td):last-child{padding-right:20px}
@media(max-width:767px){table.gridm.edge :is(th,td):first-child{padding-left:14px}table.gridm.edge :is(th,td):last-child{padding-right:14px}}
'''

_Grid0 = Grid


def Grid(cols: list[str], rows: list[tuple], caption: str = '', label_header: str = '', edge: bool = False) -> str:
    """edge: on a flush panel (web ui/Grid `edge`)."""
    html = _Grid0(cols, rows, caption, label_header)
    return html.replace('<table class="gridm"', '<table class="gridm edge"', 1) if edge else html
# ------------------------------------------------ the compact ranked list (round 5; web ui RankList `compact`, RankRow `barTone`)
# The tier list beside its rail (directions mock g-tiers): elastic name and bar tracks, 48 px rows, the letter column at every
# width (44 px, 64 px from 1280) with the letter alone in the title face, the accent bar and the accent tint on the selected row.
# Additive: RankList(compact=True), RankGroup with no band or count, RankRow(accent=True) with no sub line and no links.
PRIM_CSS += '''
.rl.compact{--rr-name:minmax(170px,1fr);--rr-bar:minmax(48px,160px);--rr-gap:72px;--rr-links:0px}
.rl.compact.grouped{--grp-w:44px}
@media(min-width:1280px){.rl.compact.grouped{--grp-w:64px}}
@media(min-width:1760px){.rl.compact{--rr-name:minmax(170px,1fr);--rr-bar:minmax(48px,160px)}}
@media(max-width:1023px){.rl.compact .rr-heads{display:flex;justify-content:flex-end;padding:7px 10px}.rl.compact .rr-heads>span:not(:nth-child(5)){display:none}}
.rl.compact .rg{grid-template-columns:var(--grp-w) minmax(0,1fr)}
.rl.compact .rg-l{display:block;padding:10px 0 0;text-align:center;border-left:0;border-right:1px solid var(--line-soft);border-bottom:0}
.rl.compact .rg-l .lt{font:400 40px/44px var(--font-title);color:var(--cream)}
.rl.compact .rr{min-height:48px;padding-block:6px}.rl.compact .rr .meter>i{background:var(--accent)}
.rl.compact .rr.cur{background:rgba(0,147,178,.16)}.rl.compact .rr.cur::before{background:var(--accent)}
@media(max-width:1279px){.rl.compact .rg-l .lt{font-size:30px;line-height:36px}.rl.compact .rr.cols{padding:8px 14px}}
@media(max-width:767px){.rl.compact .rr.cols{grid-template-areas:"rk ld nm f1" ". . xx xx";column-gap:8px;padding:6px 12px 6px 8px}
 .rl.compact .rr .br,.rl.compact .rr .gp{display:none}.rl.compact .rr .fg.f1{align-self:center}}
'''

_RankRow0, _RankGroup0, _RankList0 = RankRow, RankGroup, RankList


def RankRow(rank: str, crest: str, color: str, name: str, sub: str, frac: float, figures: list[tuple], gap: str,
            links: list[tuple[str, bool]], gap_top: bool = False, tie: str = '', low: str = '', you: str = '',
            selected: bool = False, state: str = '', accent: bool = False) -> str:
    """RankRow; `accent` is the compact row: the accent bar, no sub line, no link strip (the name is the row's one link)."""
    html = _RankRow0(rank, crest, color, name, sub, frac, figures, gap, links, gap_top, tie, low, you, selected, state)
    if not accent:
        return html
    html = html.replace(MeterClass(frac, color), Meter(frac)).replace('<span class="sub"></span>', '')
    return html.replace('<span class="lk"></span>', '')


def RankGroup(letter: str, color: str, band: str, count: str, rows: list[str]) -> str:
    """RankGroup; with no band and no count, the compact letter column (the letter alone)."""
    if band or count:
        return _RankGroup0(letter, color, band, count, rows)
    return (f'<div class="rg" role="listitem"><div class="rg-l"><span class="lt">{letter}</span></div>'
            f'<div class="rg-rows" role="list" aria-label="Tier {letter}">{"".join(rows)}</div></div>')


def RankList(heads: str, body: str, n_figs: int = 1, grouped: bool = False, links3: bool = False, tank: bool = False,
             compact: bool = False, fig_px: int = 0) -> str:
    """RankList; `compact` is web RankList `compact`, `fig_px` its `figureWidth`."""
    html = _RankList0(heads, body, n_figs, grouped, links3, tank)
    if compact:
        html = html.replace('<div class="rl', '<div class="rl compact', 1)
    if fig_px:
        html = html.replace(f'style="--n:{n_figs}"', f'style="--n:{n_figs};--rr-fig:{fig_px}px"', 1)
    return html


# ------------------------------------------------ PaperdollFrame and the in-flow client tooltip (lane P2, 2026-11-09)
# The shipped ui/PaperdollFrame (web/src/components/ui/PaperdollFrame.svelte) on the owner's frame: every socket, the
# medallion, the title bar and the centre are boxes in percent of the frame (web/src/lib/ui/paperdoll-slots.json, the
# file the build reads), so mock and build place them alike. Filled slots are buttons named "<Slot>: <item>"; the title
# bar is the page's heading; the centre holds the caller's html. Under 423 px it is the strip and the two-column list.
# `ItemTip` is ItemTooltip with floating={false}: the client tooltip in the flow, a group named by the item, its last
# lines (`tail`) in the client yellow, its source under the hairline.
import json as _json
from mocklib import WEB as _WEB, data_uri as _data_uri

_PD = _json.loads((_WEB / 'src/lib/ui/paperdoll-slots.json').read_text())
_PD_SLOTS = ['head', 'neck', 'shoulder', 'back', 'chest', 'shirt', 'tabard', 'wrist',
             'hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2',
             'main_hand', 'off_hand', 'ranged']
_PD_LIST = [s for s in _PD_SLOTS if s not in ('shirt', 'tabard')]
_PD_SEAL = _data_uri(_WEB / 'public/art/seal-128.webp')

PRIM_CSS += '''
.pdf{container-type:inline-size;width:100%}.pdf-frame,.pdf-list{display:none}
@container (min-width:423px){.pdf-frame{position:relative;display:block;width:100%;background:var(--pd-frame) 0 0/100% 100% no-repeat}}
@container (max-width:422.98px){.pdf-list{display:block}}
.pdf-slot{position:absolute;display:block;padding:0;border:0;background:none}
.pdf-icon{position:absolute;inset:6%;width:88%;height:88%;box-sizing:border-box;border:2px solid var(--qb);border-radius:3px;box-shadow:0 0 0 1px #000;background:#100a06}
.pdf-covered{filter:grayscale(1) brightness(.55)}
.pdf-slot.pdf-sel::after{content:'';position:absolute;inset:-4%;border:2px solid #f3eee2;border-radius:6px;box-shadow:0 0 0 1px #000}
.pdf-medal{position:absolute;display:flex;align-items:center;justify-content:center}.pdf-medal img{width:92%;height:92%;border-radius:50%;object-fit:cover}
.pdf-title{position:absolute;display:flex;align-items:center;gap:2cqw;margin:0;padding:0 2.2cqw;color:#f3eee2;font:400 3.4cqw/1 var(--font-title);
  letter-spacing:-.015em;text-shadow:0 1px 2px #000;white-space:nowrap}
.pdf-meta{font:600 2.3cqw var(--font-body);color:#e3dfd9}
.pdf-figure{margin-left:auto;font:600 3.2cqw var(--font-body);color:#fff;text-shadow:0 1px 2px #000;white-space:nowrap}.pdf-figure small{font-size:.66em;color:#e3dfd9}
.pdf-centre{position:absolute;display:flex;flex-direction:column;align-items:center;justify-content:center;gap:3cqw;padding:4cqw}
.pdf-strip{display:flex;align-items:center;gap:10px;padding:8px 12px;border:1px solid #000;border-radius:10px 10px 0 0;background:#0b2f3b}
.pdf-strip img{width:40px;height:40px;border-radius:50%}.pdf-strip :is(h1,h2){min-width:0;margin:0;font:400 19px/1.2 var(--font-title);color:#f3eee2}
.pdf-strip .pdf-figure{font-size:17px}
.pdf-items{display:grid;grid-template-columns:1fr 1fr;gap:6px 10px;margin:0;padding:12px;list-style:none;border:1px solid #000;border-top:0;border-radius:0 0 10px 10px;background:#170e09}
.pdf-item{display:grid;grid-template-columns:44px minmax(0,1fr);gap:8px;align-items:center;width:100%;min-height:44px;padding:0;border:0;background:none;color:inherit;font:inherit;text-align:left}
.pdf-cell{position:relative;width:44px;height:44px;border-radius:3px;background:#100a06}.pdf-cell .pdf-icon{inset:3px;width:38px;height:38px}
.pdf-name{display:block;font-size:13px;font-weight:600;line-height:16px;overflow-wrap:anywhere}.pdf-sub{display:block;font-size:12px;line-height:16px;color:#8b8784}
.pdf-list-centre{margin-top:10px}
/* The built frame's additive list props (lane h-doll): `form="list"` the list at every width, `head={false}` the slots straight on
   the page's panel, `facts` "Head · Item level 61" with the enchant green on its own line. */
.pdf-listonly .pdf-list{display:block}
.pdf-headless{padding:0;border:0;border-radius:0;background:none;box-shadow:none}
.pdf-facts .pdf-sub{overflow:hidden;text-overflow:ellipsis;white-space:nowrap}.pdf-facts .pdf-enchant{white-space:normal;color:#1eff00}
@container (min-width:360px) and (max-width:559.98px){.pdf-fact-slot{display:none}}
@container (max-width:359.98px){.pdf-facts{grid-template-columns:minmax(0,1fr)}}
.itip{display:flex;flex-direction:column;width:100%;font-size:13px;color:#fff;border:1px solid #8b8f9c;border-radius:4px;background:#060914;box-shadow:inset 0 0 0 1px #2b3040}
.itip .bd{display:flex;flex-direction:column;gap:8px;padding:12px 14px}.itip .nm{font-size:15px;font-weight:700}.itip .ln{font-size:12px}
.itip .tl{font-size:12px;color:#ffd100}.itip .ft{padding:8px 14px 12px;font-size:12px;line-height:16px;color:#8b8784;border-top:1px solid #2b3040}
'''


def _pd_box(name: str) -> str:
    b = _PD['slots'][name] if name in _PD['slots'] else _PD[name]
    return f'left:{b["x"]}%;top:{b["y"]}%;width:{b["w"]}%;height:{b["h"]}%'


def _pd_icon(cell: dict) -> str:
    """The item's icon in its quality border, as the built `.pdf-icon` (a socket when the checkout has no icon file)."""
    q = QCOL.get(cell.get('quality', 1), QCOL[1]) if isinstance(cell.get('quality'), int) else RARITY.get(cell.get('quality'), RARITY['common'])
    cov = ' pdf-covered' if cell.get('covered') else ''
    src = _icon_src(cell.get('icon'))
    if src:
        return f'<img class="pdf-icon{cov}" src="{src}" alt="" style="--qb:{q}">'
    return f'<span class="pdf-icon{cov}" style="--qb:{q}"></span>'


def _icon_src(name: str | None) -> str:
    html = icon(name, 40) if name else ''
    m = re.search(r'src="([^"]+)"', html)
    return m.group(1) if m else ''


def PaperdollFrame(slots: list[dict], title: str, meta: str, figure: tuple[str, str] | None, centre: str, label: str,
                   list_centre: str = '', heading: str = 'h1', form: str = 'auto', head: bool = True, facts: bool = False) -> str:
    """slots: the filled sockets (slot, label (the button's name), covered, selected, icon, quality; with `facts`, line2 and an
    optional enchant). centre: the centre's html.
    heading: the title bar's heading (the built `heading`): h1 where the bar is the page's title (BiS), h2 under a page title.
    form 'list': the list form at every width, the frame never drawn; head False: no strip, the slots on the page's panel; facts:
    each row's sub line is "Slot · line 2", the enchant green under it (the built `form`, `head`, `facts`)."""
    tag = {'h1': 'h1', 'h2': 'h2'}[heading]
    by = {s['slot']: s for s in slots}
    fig = f'<b class="pdf-figure">{escape(figure[0])} <small>{escape(figure[1])}</small></b>' if figure else ''
    buttons = ''.join(
        f'<button type="button" class="pdf-slot{" pdf-sel" if by[n]["selected"] else ""}" style="{_pd_box(n)}" '
        f'aria-label="{escape(by[n]["label"])}" aria-pressed="{"true" if by[n]["selected"] else "false"}">{_pd_icon(by[n])}</button>'
        for n in _PD_SLOTS if n in by)
    frame_bg = _data_uri(_WEB / 'public/art/paperdoll-frame@1x.webp')
    w, h = _PD['size']
    frame = (f'<div class="pdf-frame" style="aspect-ratio:{w}/{h};--pd-frame:url({frame_bg})" data-region="{REGION["slots"]}">{buttons}'
             f'<span class="pdf-medal" style="{_pd_box("medallion")}" data-region="{REGION["paperdoll_medallion"]}"><img src="{_PD_SEAL}" alt=""></span>'
             f'<div class="pdf-title" style="{_pd_box("title")}"><{tag} style="margin:0;font:inherit">{escape(title)}</{tag}><small class="pdf-meta">{escape(meta)}</small>{fig}</div>'
             f'<div class="pdf-centre" style="{_pd_box("centre")}" data-region="{REGION["paperdoll_centre"]}">{centre}</div></div>')
    items = []
    for n in _PD_LIST:
        s = by.get(n)
        if s is None:
            items.append('<li><span class="pdf-item"><span class="pdf-cell"></span><span class="pdf-sub"></span></span></li>')
            continue
        slot_word, _, item = s['label'].partition(': ')
        items.append(f'<li><button type="button" class="pdf-item" aria-pressed="{"true" if s["selected"] else "false"}">'
                     f'<span class="pdf-cell">{_pd_icon(s)}</span><span><span class="pdf-name">{escape(item)}</span>'
                     f'{_pd_facts(slot_word, s) if facts else f"<span class=pdf-sub>{escape(slot_word)}</span>"}</span></button></li>')
    strip = f'<div class="pdf-strip"><img src="{_PD_SEAL}" alt=""><{tag}>{escape(title)}</{tag}>{fig}</div>' if head else ''
    ul_class = 'pdf-items' + ('' if head else ' pdf-headless') + (' pdf-facts' if facts else '')
    lst = (f'<div class="pdf-list">{strip}'
           f'<ul class="{ul_class}">{"".join(items)}</ul>{f"<div class=pdf-list-centre>{list_centre}</div>" if list_centre else ""}</div>')
    list_only = form == 'list'
    return (f'<div class="pdf{" pdf-listonly" if list_only else ""}" role="group" aria-label="{escape(label)}">'
            f'{"" if list_only else frame}{lst}</div>')


def _pd_facts(slot_word: str, cell: dict) -> str:
    """A `facts` row's sub lines: "Head · Item level 61" (the slot word hides in a phone's half column), the enchant green."""
    line2 = cell.get('line2') or ''
    lead = (f'<span class="pdf-fact-slot">{escape(slot_word)} &middot; </span><span>{escape(line2)}</span>' if line2
            else escape(slot_word))
    enchant = f'<span class="pdf-sub pdf-enchant">{escape(cell["enchant"])}</span>' if cell.get('enchant') else ''
    return f'<span class="pdf-sub">{lead}</span>{enchant}'


def ItemTip(name: str, quality: str, lines: list[str], tail: list[str], foot: str) -> str:
    """The client tooltip in the flow (ItemTooltip floating={false}): its lines, the yellow tail paragraphs, the source."""
    body = ''.join(f'<span class="ln">{escape(x)}</span>' for x in lines[1:])
    t = f'<div class="tl">{"".join(f"<p>{escape(x)}</p>" for x in tail)}</div>' if tail else ''
    f = f'<div class="ft"><div>{escape(foot)}</div></div>' if foot else ''
    return (f'<div class="itip" role="group" aria-label="{escape(name)}"><div class="bd"><span class="nm" style="color:{RARITY.get(quality, RARITY["common"])}">'
            f'{escape(name)}</span>{body}{t}</div>{f}</div>')



# ------------------------------------------------ ReadingPanel (lane P8, 2026-11-09; web ui/ReadingPanel.svelte)
# A reading page's ReadingSections on one panel (directions mock b2-guide-spec): no head hairline, a hairline between sections,
# each padded 22/24 (18/16 below 768), the heading in Marcellus 24/30 cream; from 1280 the prose takes the room and the lane is a
# 380 px column behind a hairline. `PlateColumns(frame='framed')` is the built PlateColumns' `frame` (the cream inset line).
PRIM_CSS += """
.rpanel{min-width:0;background:var(--raised);border:1px solid #000;border-radius:10px;box-shadow:0 0 0 1px rgba(130,112,91,.55);--muted:#8b8784}
.rpanel>.rsec{padding:22px 24px}.rpanel>.rsec+.rsec{border-top:1px solid var(--line-soft)}
.rpanel>.rsec>.sec-h{padding-bottom:0;border-bottom:0}
.rpanel .rsec>.sec-h h2{font:400 24px/30px var(--font-title);letter-spacing:-.015em;color:var(--cream)}
.rpanel .rd{padding-top:12px}
@media(max-width:767px){.rpanel>.rsec{padding:18px 16px}}
@media(min-width:1280px){.rpanel .rd:not(.solo){grid-template-columns:minmax(0,var(--measure-reading,720px)) minmax(380px,1fr);column-gap:28px}
.rpanel .rd:not(.solo)>.rd-lane{padding-left:18px;border-left:1px solid var(--line-soft)}}
.plate.framed{position:relative}
.plate.framed::before{content:"";position:absolute;inset:6px;border:1px solid rgba(243,238,226,.42);border-radius:6px;pointer-events:none}
"""


def ReadingPanel(sections: str) -> str:
    return f'<div class="rpanel">{sections}</div>'


# The lane wraps under its prose, and a short lane is pinned, by the shipped code itself (web/src/lib/ui/reading-wrap.ts,
# bundled once per run with the web's esbuild), so the mock and the build place every row alike.
PRIM_CSS += '''
.rd-more{display:none}.rd-more:not(:empty){display:block;margin-top:24px;padding-top:16px;border-top:1px solid var(--line-soft);columns:340px;column-gap:28px}
.rd-more [data-wrapped]{break-inside:avoid}
@media(max-width:1279.98px){.rd[data-prose-first]{grid-template-areas:'prose' 'lane'}}
@media(min-width:1280px){.rd{align-items:start}.rd[data-lane-sticky]>.rd-lane{position:sticky;top:12px}}
'''
_READING_WRAP: list[str] = []


def reading_wrap_script() -> str:
    """The shipped reading wrap as an inline script, mounted on the page (a page with reading sections appends it)."""
    if not _READING_WRAP:
        import subprocess
        web = _WEB
        out = subprocess.run([str(web / 'node_modules/.bin/esbuild'), str(web / 'src/lib/ui/reading-wrap.ts'), '--bundle', '--format=iife',
                              '--global-name=FSReadingWrap', '--target=es2022', '--log-level=error'],
                             capture_output=True, text=True, check=True).stdout
        _READING_WRAP.append(out)
    return f'<script>{_READING_WRAP[0]}FSReadingWrap.mountReadingWrap();</script>'


_PlateColumns0 = PlateColumns


def PlateColumns(cols: list[tuple[str, str, str, str]], label: str = '', frame: str = 'plain') -> str:
    html = _PlateColumns0(cols, label)
    return html.replace('<div class="plate"', '<div class="plate framed"', 1) if frame == 'framed' else html



# ------------------------------------------------ ClassSpecPicker (lane P8, 2026-11-09; web components/guides/ClassSpecPicker.astro)
# The class and spec picker /guides and /bis share (directions b2-guides): one panel per class, three across from 1024 and two
# from 768: the class crest and name in the class colour (a link on /guides), then its specs as slim link rows, no role tag.
# classes: (slug, name, colour, [(key, name, icon_html)]); `yours` marks the header character's spec with `yours_label`.
PRIM_CSS += """
.cspick{display:grid;grid-template-columns:minmax(0,1fr);gap:18px;align-items:start}
@media(min-width:768px){.cspick{grid-template-columns:repeat(2,minmax(0,1fr))}}
@media(min-width:1024px){.cspick{grid-template-columns:repeat(3,minmax(0,1fr))}}
.cspick-c{padding-bottom:8px;min-width:0}
"""


def ClassSpecPicker(classes: list[tuple[str, str, str, list[tuple[str, str, str]]]], label: str, class_links: bool = False,
                    yours: str = '', yours_label: str = '') -> str:
    groups = []
    for slug, name, colour, specs in classes:
        rows = [Row(n, lead=ic, selected=key == yours) for key, n, ic in specs]   # the accent rule marks yours, no pill
        head = Row(name, lead=crest_img(slug, 40), title_color=colour, link=class_links)
        groups.append(Plate(f'<div class="cspick-c">{head}{RowsSlim(rows, open_=False)}</div>'))
    return f'<div class="cspick" role="region" aria-label="{label}">{"".join(groups)}</div>'


# ------------------------------------------------ the panelled workbench (lane P7; web ui/Page `panels` on 'workbench', ui/TalentEditor
# `title`, `frame`): the planner's panels 22 px apart and the columns 28 from 1280; the trees in a Section panel with its heading
# inside ("Talents", framed) instead of a bare Plate. Additive: Page(panels=False) and TalentEditor(title='') emit what they did.
PRIM_CSS += '''
.layout-wb.panels{--section-gap:22px}
.porder-s .ah img,.porder-s .ah .ic{opacity:.5}.porder-s .ah b{opacity:.7}
.porder-s .ah::after{content:"";position:absolute;inset:0;border-radius:5px;border:1px dashed var(--line)}
@media(min-width:1280px){.layout-wb.panels{column-gap:28px}.layout-wb.panels>.main,.layout-wb.panels>.rail{gap:22px}}
'''
# The built PointOrder's phone sockets: 30 px, shrinking with the column below 336 px so a row never runs past it (ui/PointOrder).
PRIM_CSS += '''
@media(max-width:767px){.porder-r .porder-s>span,.porder-r .porder-s img,.porder-r .porder-s .ic,.porder-r .porder-s .em{width:100%!important;max-width:30px;height:auto!important;aspect-ratio:1}.porder-r .porder-s img,.porder-r .porder-s .ic{display:block}}
'''

_Page2, _TalentEditor0 = Page, TalentEditor


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False, panels: bool = False, **kwargs: object) -> str:
    """Page; `panels` on 'workbench' sets the workbench's panels apart as the split's (web ui/Page `panels`)."""
    if layout != 'workbench' or not panels:
        return _Page2(layout, main, rail, sticky_rail, panels, **kwargs)
    html = _Page2(layout, main, rail, sticky_rail)
    return html.replace('style="padding-top:20px"><div class="layout-wb"', 'style="padding-top:28px"><div class="layout-wb panels"', 1)


def TalentEditor(trees: list[dict], active: int = 0, foot: str = '', label: str = 'Talent trees', title: str = '', frame: str = 'plain') -> str:
    """TalentEditor; `title` puts the trees in a Section panel (its heading inside, its `frame`), as the built `title`."""
    html = _TalentEditor0(trees, active, foot, label)
    if not title:
        return html
    inner = html[html.index('<div class="tedit">'):html.rindex('</div>')]
    return Section(title, inner, panel=True, frame=frame, flush=True)


# ------------------------------------------------ the workbench's row below (h-audit 2026-11-09; web ui/Page `below`): full width under
# main and the rail from 1280, its Ordered regions in the one flow below 1280. Additive: Page(below='') emits what it did.
_PageBelow0 = Page


def Page(layout: str, main: str = '', rail: str = '', sticky_rail: bool = False, panels: bool = False, below: str = '',
         **kwargs: object) -> str:
    """Page; `below` on 'workbench' is a full-width row under main and the rail from 1280."""
    html = _PageBelow0(layout, main, rail, sticky_rail, panels, **kwargs)
    if not below or layout != 'workbench':
        return html
    at = html.rindex('</aside>') + len('</aside>')
    return html[:at] + f'<div class="below">{below}</div>' + html[at:]


PRIM_CSS += '''
.layout-wb>.below{display:contents}
@media(min-width:1280px){.layout-wb>.below{grid-column:1/-1;display:flex;flex-direction:column;gap:var(--section-gap);min-width:0}}
'''


# ------------------------------------------------ the Subscribe button and the Plate frame (batch F, /premium): web lib/ui/button.ts
# BUTTON_SUBSCRIBE is Blizzard Forever's teal Subscribe exactly (flat teal-edge, white semibold label); ui/Plate `frame`.
# Additive: Button(tone='subscribe') and Plate(frame=...) are new; every other call emits what it did.
PRIM_CSS += '''
.btn.subscribe{background:var(--teal-edge);border-color:var(--teal-edge);color:var(--strong)}
'''

_Button3, _Plate1 = Button, Plate


def Button(label: str, primary: bool = False, small: bool = False, tone: str = '', act: bool = False) -> str:
    """tone 'subscribe': the Premium plan's teal Subscribe (web button('subscribe'))."""
    if tone != 'subscribe':
        return _Button3(label, primary, small, tone, act)
    return _Button3(label, False, small, '', act).replace('class="btn', 'class="btn subscribe', 1)


def Plate(inner: str, label: str = '', frame: str = 'plain') -> str:
    """frame: 'framed' draws the cream inset line (web ui/Plate `frame`)."""
    html = _Plate1(inner, label)
    return html.replace('class="plate"', 'class="plate framed"', 1) if frame == 'framed' else html


# ------------------------------------------------ one talent tree (talent-tree lane 2026-11-10; design/reviews/2026-11-10-talent-tree)
# Every tree on the site is web ui/TalentTrees in a mode (edit: the planner; view: the sim, /sim/talents; compact: a guide), so the
# mocks draw every tree with one primitive too. The CSS is the built stylesheet itself (web components/ui/talent-tree.css), its
# `--color-*` tokens read as the mock's own. `talent_tree_dicts` is web lib/planner/editor-trees.ts `viewTrees` over a class's
# talent file (the cells' states, the spec icon, the art, the arrows); `TalentEditor`, `Trees` and `TalentTrees` all draw through
# `TalentTreesAll`, the built markup.
import base64 as _b64
import json as _json
from mocklib import ROOT as _ROOT

_TT_CSS = (_ROOT / 'web/src/components/ui/talent-tree.css').read_text()
PRIM_CSS += ':root{--radius-field:4px;--selected:#0093b229;--shadow-selected-underline:inset 0 -3px 0 #0093b2}\n'
PRIM_CSS += re.sub(r'/\*.*?\*/', '', _TT_CSS, flags=re.S).replace('var(--color-', 'var(--') + '\n'
_TT_BUILD = _json.loads((_ROOT / 'web/src/data/active-build.json').read_text())['build']
_TT_SPECS = _json.loads((_ROOT / 'web/src/data/specs.json').read_text())
_tt_art: dict[str, str] = {}
_CELL, _GAP = 44, 8
_PITCH, _HALF = _CELL + _GAP, _CELL // 2


def _tt_img(icon_html: str) -> str:
    """A mock icon as the built cell's face: the img (or the dark socket) at the cell's size, its fixed size dropped."""
    html = re.sub(r'\s(?:style|width|height)="[^"]*"', '', icon_html)
    return html.replace('class="ic"', 'class="tt-ic"').replace('class="ic ', 'class="tt-ic ')


def _tt_art_src(bg: str) -> str:
    if bg not in _tt_art:
        p = _ROOT / 'data/builds' / _TT_BUILD / 'trees' / f'{bg}.webp'
        _tt_art[bg] = 'data:image/webp;base64,' + _b64.b64encode(p.read_bytes()).decode() if p.exists() else ''
    return _tt_art[bg]


def _tt_centre(tier: int, col: int) -> tuple[int, int]:
    return col * _PITCH + _HALF, tier * _PITCH + _HALF


def _tt_path(a: dict, b: dict) -> str:
    """web lib/planner/connectors.ts connectorPath: down a column, across a row, or the elbow."""
    ax, ay = _tt_centre(a['tier'], a['column'])
    bx, by = _tt_centre(b['tier'], b['column'])
    if a['column'] == b['column']:
        return f'M {ax} {ay + _HALF} V {by - _HALF}'
    side = 1 if b['column'] > a['column'] else -1
    if a['tier'] == b['tier']:
        return f'M {ax + side * _HALF} {ay} H {bx - side * _HALF}'
    return f'M {ax} {ay + _HALF} V {by} H {bx - side * _HALF}'


def _tt_file(class_slug: str) -> dict:
    return _json.loads((_ROOT / 'data/builds' / _TT_BUILD / 'talents' / f'{class_slug}.json').read_text())


def talent_tree_dicts(class_slug: str, ranks: dict[int, int]) -> list[dict]:
    """A read-only build's trees (web `viewTrees`): untaken greyed, partial green, maxed gold, arrows met in gold."""
    out = []
    for index, tree in enumerate(sorted(_tt_file(class_slug)['trees'], key=lambda t: t['position'])):
        talents = sorted(tree['talents'], key=lambda t: (t['tier'], t['column']))
        by_id = {t['id']: t for t in talents}
        cols = max(t['column'] for t in talents) + 1
        tiers = max(7, max(t['tier'] for t in talents) + 1)
        cells = []
        for t in talents:
            rk = min(t['max_rank'], ranks.get(t['id'], 0))
            state = 'max' if rk >= t['max_rank'] else ('part' if rk else 'lock')
            cells.append((t['tier'], t['column'], TalentCell2(icon(t['icon'], 48), t['name'], rk, t['max_rank'], state, editable=False)))
        conns = sorted((t['prereq_talent_id'], t['id']) for t in talents if t.get('prereq_talent_id') in by_id)
        connectors = [(_tt_path(by_id[a], by_id[b]), ranks.get(a, 0) >= (by_id[b].get('prereq_rank') or 0)) for a, b in conns]
        spec = next((s for s in _TT_SPECS if s['class_slug'] == class_slug and s['tree_index'] == index), None)
        out.append({'name': tree['name'], 'icon_html': icon(spec['icon'], 28) if spec else '',
                    'points': sum(min(t['max_rank'], ranks.get(t['id'], 0)) for t in talents),
                    'art_src': _tt_art_src(tree['background']) if tree.get('background') else '',
                    'cols': cols, 'tiers': tiers, 'viewbox': f'0 0 {cols * _PITCH - _GAP} {tiers * _PITCH - _GAP}',
                    'connectors': connectors, 'cells': cells})
    return out


def talent_ranks_from_strings(class_slug: str, per_tree: list[str]) -> dict[int, int]:
    """An engine talents string's trees (`05505`, `5530515`, ``) as talent id to rank: each digit in tier, column order."""
    ranks: dict[int, int] = {}
    for index, tree in enumerate(sorted(_tt_file(class_slug)['trees'], key=lambda t: t['position'])):
        digits = per_tree[index] if index < len(per_tree) else ''
        for p, t in enumerate(sorted(tree['talents'], key=lambda t: (t['tier'], t['column']))):
            ranks[t['id']] = min(t['max_rank'], int(digits[p]) if p < len(digits) else 0)
    return ranks


def TalentCell2(icon_html: str, name: str, rank: int, mx: int, state: str, differs: bool = False, extra: str = '',
                editable: bool = True) -> str:
    """One talent (web ui/TalentButton). state: avail | part | max | lock. Editable: wrapped in its own tooltip host."""
    cls = f'tt-cell {state}' + (' diff' if differs else '') + (f' {extra}' if extra else '')
    dis = ' aria-disabled="true"' if editable and state == 'lock' else ''
    beside = '' if editable else ' data-tip-beside'
    btn = (f'<button type="button" class="{cls}" aria-label="{escape(name)}, rank {rank} of {mx}"{dis} data-rank="{rank}" data-state="{state}"{beside}>'
           f'{_tt_img(icon_html)}<span class="tt-rk" aria-hidden="true">{rank}/{mx}</span></button>')
    return f'<span class="tt-wrap" data-tip-beside>{btn}</span>' if editable else btn


def _tt_head_icon(t: dict, px: int) -> str:
    html = re.sub(r'\s(?:style|width|height)="[^"]*"', '', t['icon_html'])
    return html.replace('<img ', f'<img width="{px}" height="{px}" ', 1)


def TalentTreesAll(trees: list[dict], mode: str = 'edit', active: int = 0, foot: str = '', label: str = 'Talent trees',
                   title: str = '', frame: str = 'plain') -> str:
    """web ui/TalentTrees. trees: {name, icon_html, points, art_src, cols, tiers, viewbox, connectors [(d, met)], cells
    [(tier, col, cell_html)], reset}. mode: edit (the Plate or a titled Section panel), view (bare, md), compact (bare, sm)."""
    size = {'edit': 'lg', 'view': 'md', 'compact': 'sm'}[mode]
    tabs = ''
    if size != 'sm':
        tabs = f'<div class="tt-tabs" role="tablist" aria-label="{label}">' + ''.join(
            f'<button type="button" role="tab" aria-selected="{"true" if i == active else "false"}"><span>{_tt_head_icon(t, 22)}<b>{t["points"]}</b></span>{escape(t["name"])}</button>'
            for i, t in enumerate(trees)) + '</div>'
    cols = ''
    for i, t in enumerate(trees):
        rows = ''
        for tier in range(t['tiers']):
            row = ''.join(f'<span role="gridcell">{next((c[2] for c in t["cells"] if c[0] == tier and c[1] == col), "")}</span>' for col in range(t['cols']))
            rows += f'<div role="row">{row}</div>'
        paths = ''.join(f'<path class="halo" d="{d}"/><path class="{"on" if met else "off"}" d="{d}"/>' for d, met in t['connectors'])
        art = f'<img class="tt-art" src="{t["art_src"]}" alt="" aria-hidden="true">' if t.get('art_src') else ''
        pts = f'<b class="num{" zero" if not t["points"] else ""}">{t["points"]}</b>'
        rs = t.get('reset', '').replace('class="trs"', 'class="tt-reset"') if mode == 'edit' else ''
        panel = f' role="tabpanel" aria-label="{escape(t["name"])}"' if size != 'sm' else f' aria-label="{escape(t["name"])}"'
        cols += (f'<div class="tt-col{" on" if i == active else ""}"{panel}><div class="tt"><div class="tt-h">{_tt_head_icon(t, 28)}<h3>{escape(t["name"])}</h3>{pts}{rs}</div>'
                 f'<div class="tt-g" role="grid" aria-label="{escape(t["name"])} talents" style="--cols:{t["cols"]}">{art}'
                 f'<svg class="tt-arrows" viewBox="{t["viewbox"]}" preserveAspectRatio="none" aria-hidden="true">{paths}</svg>{rows}</div></div></div>')
    f = f'<div class="tt-f">{foot}</div>' if foot else ''
    bare = '' if mode == 'edit' else ' data-bare'
    body = f'<div class="tt-trees" data-size="{size}" data-mode="{mode}"{bare}>{tabs}<div class="tt-cols" style="--n:{len(trees)}">{cols}</div>{f}</div>'
    if mode != 'edit':
        return f'<div>{body}</div>'
    if title:
        return Section(title, body, panel=True, frame=frame, flush=True)
    return Plate(body, label)


def TalentEditor(trees: list[dict], active: int = 0, foot: str = '', label: str = 'Talent trees', title: str = '', frame: str = 'plain') -> str:
    """The planner's trees: TalentTreesAll in edit mode."""
    return TalentTreesAll(trees, 'edit', active, foot, label, title, frame)


def Trees(trees: list[dict]) -> str:
    """A read-only build's trees (the sim's Talents, /sim/talents): TalentTreesAll in view mode."""
    return TalentTreesAll(trees, 'view')


def TalentTrees(trees: list[dict]) -> str:
    """A guide's trees (At a glance, the Leveling lane): TalentTreesAll in compact mode."""
    return TalentTreesAll(trees, 'compact')


def TalentTip(name: str, rank: int, mx: int, now: str, nxt: str, req: str = '', band: str = '', next_label: str = 'Next rank') -> str:
    """The talent's game tooltip (web GameTooltip over lib/ui/talent-tree cellTipModel): name and Rank n/m, the refusal in red,
    this rank (the first rank's text when untaken), the next rank under its label, the band line in gold."""
    desc = now or (nxt if rank == 0 else '')
    nx = nxt if rank > 0 else ''
    r = f'<p style="color:var(--wipe);margin:0;font-size:12px">{escape(req)}</p>' if req else ''
    d = f'<p style="color:#ffd100;margin:0;font-size:12px">{escape(desc)}</p>' if desc else ''
    x = (f'<p style="color:var(--muted);margin:0;font-size:12px"><span style="display:block;font-weight:600">{next_label}</span>'
         f'{escape(nx)}</p>') if nx else ''
    bd = f'<p style="color:var(--gold);margin:0;font-size:12px;font-weight:600">{escape(band)}</p>' if band else ''
    return (f'<div class="ttip" role="tooltip"><span style="display:flex;justify-content:space-between;gap:8px"><b>{escape(name)}</b>'
            f'<span class="r">Rank {rank}/{mx}</span></span>{r}{d}{x}{bd}</div>')


_tt_class_of: dict[int, str] = {}


def talent_class(talent_id: int) -> str:
    """The class whose talent file holds this talent id (a guide's facts name the talents, not the class)."""
    if not _tt_class_of:
        for f in sorted((_ROOT / 'data/builds' / _TT_BUILD / 'talents').glob('*.json')):
            for tree in _json.loads(f.read_text())['trees']:
                _tt_class_of.update({t['id']: f.stem for t in tree['talents']})
    return _tt_class_of[talent_id]

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
.ph-title .t-title{overflow-wrap:anywhere}
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


def Section(title: str, body: str, action: str = '', meta: str = '') -> str:
    a = f'<span class="act">{action}</span>' if action else ''
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    return f'<section data-region="{region_slug(title)}"><div class="sec-h"><h2 class="t-h3">{title}</h2>{m}{a}</div>{body}</section>'


def TabStrip(tabs: list[str], active: str, counts: dict | None = None) -> str:
    return '<nav class="strip-tabs ph-tabs">' + ''.join(f'<a href="#" class="{"cur" if t == active else ""}">{t}{f"<span class=\"n\">{counts[t]}</span>" if counts and t in counts else ""}</a>' for t in tabs) + '</nav>'


def PageHeader(title: str, title_color: str = '', crest: str = '', meta: list[str] | None = None, answer: str = '', actions: str = '', tabs: str = '') -> str:
    m = f'<div class="ph-meta" data-region="{REGION["head_meta"]}">{"".join(f"<span class=\"t-small\">{x}</span>" for x in meta)}</div>' if meta else ''
    color = f' style="color:{title_color}"' if title_color else ''
    right = f'<div class="ans" data-region="{REGION["head_answer"]}">{answer}</div>' if answer else (f'<div class="ph-actions">{actions}</div>' if actions else '')
    return (f'<div class="pc ph-wrap"><div data-region="{REGION["head"]}"><div class="head"><div class="ph-title" data-region="{REGION["head_title"]}">{crest}<div><h1 class="t-title"{color}>{title}</h1>{m}</div></div>{right}</div>{tabs}</div></div>')


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
.tally .t .n.ok{color:var(--kill)}.tally .t .n.miss{color:var(--wipe)}.tally .t .n.warn{color:var(--gold)}
.ck{width:22px;height:22px;border-radius:50%;display:inline-flex;align-items:center;justify-content:center;font:700 13px/22px var(--font-body);color:#07090d;flex:none}
.ck.ok{background:var(--kill)}.ck.warn{background:var(--gold)}.ck.miss{background:var(--wipe)}.ck.gap{background:var(--muted)}
.gap-tag{font-size:12px;color:var(--muted)}
.seg.sm a{height:32px;padding:0 10px;font-size:13px}
'''


def Segmented(items: list[tuple[str, bool]], small: bool = False) -> str:
    """An inline choice of one among few (sign-up In / Late / Bench). Same look as the ControlBar's segmented control."""
    return seg([(a, on, None, '') for a, on in items], 'sm' if small else '')


def Grid(cols: list[str], rows: list[tuple]) -> str:
    """A coverage matrix. rows: (label, sub, [(state, text)]) with state ok|warn|miss|none. Reused for buffs, attendance, readiness."""
    h = '<div class="h"></div>' + ''.join(f'<div class="h" style="text-align:center">{c}</div>' for c in cols)
    body = ''.join(f'<div class="l"><div class="ttl">{lab}</div>' + (f'<div class="sub">{sub}</div>' if sub else '') + '</div>' + ''.join(f'<div class="c {st}">{tx}</div>' for st, tx in cells) for lab, sub, cells in rows)
    return f'<div class="gridm" style="--n:{len(cols)}">{h}{body}</div>'


def Tally(items: list[tuple]) -> str:
    """Counts against a target, one Meter each. items: (lead_html, label, have, want)."""
    out = ''
    for lead, label, have, want in items:
        st = 'ok' if have >= want else ('warn' if have >= want * 0.75 else 'miss')
        out += f'<div class="t"><div class="top">{lead}<b>{label}</b><span class="n {st}">{have}/{want}</span></div>{Meter(have / want if want else 1, "kill" if st == "ok" else "")}</div>'
    return f'<div class="tally">{out}</div>'


def Checklist(items: list[tuple]) -> str:
    """Private pre-raid checks. items: (state ok|warn|miss|gap, title, sub, fix_label|None). A Row per check, the fix a link."""
    rows = []
    for st, title, sub, fix in items:
        mark = {'ok': '&#10003;', 'warn': '!', 'miss': '&times;', 'gap': '?'}[st]
        word = {'ok': 'Done', 'warn': 'To do', 'miss': 'Missing', 'gap': 'Not tracked'}[st]
        rows.append(Row(title, sub, lead=f'<span class="ck {st}" role="img" aria-label="{word}">{mark}</span>', trail=(f'<span class="link">{fix}</span>' if fix else ''), link=bool(fix)))
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
.rl.links3{--rr-links:156px}.rl.tank{--rr-fig:124px}
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
.rr .nm .sub{font-size:13px;line-height:18px;color:var(--muted)}
.rr .nm .low{font-size:13px;line-height:18px;color:var(--muted);text-decoration:underline dotted;text-underline-offset:3px;align-self:flex-start}
.rr .br{}
.rr .x{display:contents}
.rr .fg{text-align:right;font:700 16px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong);white-space:nowrap}
.rr .fg small{font-size:12px;font-weight:600;color:var(--muted);margin-left:4px}.rr .fg.best{color:var(--kill)}.rr .fg .lbl{display:none}
.rr .gp{text-align:right;font:600 15px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--text);white-space:nowrap}
.rr .gp.top{color:var(--strong);font-weight:700}.rr .gp .tie{display:block;font:700 12px/16px var(--font-body);color:var(--strong)}
.rr .lk{display:flex;justify-content:flex-end;gap:16px}
.rr .lk a{display:inline-flex;align-items:center;min-height:40px;font:600 14px/20px var(--font-body);color:var(--nav)}.rr .lk a.p{color:var(--gold)}
.rr .lk a:hover,.rr .lk a.hov{color:var(--gold-hover)}.rr .lk a.fv{outline:2px solid var(--gold);outline-offset:2px;border-radius:3px}
.tag.you{margin-left:2px}
.meter.cls>i{background:var(--c)}
.skel{display:block;border-radius:4px;background:linear-gradient(90deg,#131824,#1a2030,#131824)}
@media(max-width:1099px){.rg{grid-template-columns:minmax(0,1fr)}.rg-l{display:flex;align-items:baseline;gap:12px;padding:8px 14px 8px 12px;border-right:0;border-bottom:1px solid var(--line-soft)}.rg-l .lt{font-size:22px;line-height:24px}.rg-l .bd{order:1}.rg-l .ct{order:2;margin-left:auto}}
@media(max-width:1023px){.rl .rr-heads{display:none}
 .rl .rr.cols{grid-template-columns:22px 32px minmax(0,1fr) auto;grid-template-areas:"rk ld nm f1" "rk ld nm gp" ". . xx xx" ". . br br" ". . lk lk";column-gap:12px;row-gap:0;padding:10px 14px 2px}
 .rr .rk{grid-area:rk}.rr .ld{grid-area:ld}.rr .nm{grid-area:nm}.rr .br{grid-area:br}.rr .gp{grid-area:gp}.rr .lk{grid-area:lk}.rr .fg.f1{grid-area:f1}.rr .x{grid-area:xx;display:flex;gap:0 16px;flex-wrap:wrap}.rr .x .fg{text-align:left;font-size:14px}.rr .x .fg .lbl{display:inline;font-weight:400;color:var(--muted);margin-right:6px}
 .rr .br{padding-top:8px}.rr .gp{align-self:start}.rr .fg.f1{align-self:end}.rr .x:not(:empty){padding-top:6px}.rr .gp .tie{display:inline;margin-left:6px}.rr .lk{justify-content:flex-start;gap:20px}.rr .lk a{min-height:44px}}
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
    tie_html = f'<span class="tie">{tie}</span>' if tie else ''
    low_html = f'<span class="low">{low}</span>' if low else ''
    you_html = f'<span class="pill tag you">{you}</span>' if you else ''
    lk = ''.join(f'<a href="#" class="{"p" if p else ""}{" " + state.split(":")[1] if state.startswith("link:") and i == 0 else ""}">{t}</a>' for i, (t, p) in enumerate(links))
    cls = 'rr cols' + (' cur' if selected else '') + (' hov' if state == 'hover' else '')
    return (f'<div class="{cls}" role="listitem" style="--c:{color}"><span class="rk">{rank}</span><span class="ld">{crest}</span>'
            f'<div class="nm"><span class="ttl"><a href="#">{name}</a>{you_html}</span><span class="sub">{sub}</span>{low_html}</div>'
            f'<span class="br">{MeterClass(frac, color)}</span>{first}<span class="x">{extra}</span>'
            f'<span class="gp{" top" if gap_top else ""}">{gap}{tie_html}</span><span class="lk">{lk}</span></div>')


def RankGroup(letter: str, color: str, band: str, count: str, rows: list[str]) -> str:
    """A letter group: the 148 px label column (letter, count, band) with the tier rail at x=0; a header strip below 1100."""
    return (f'<div class="rg" style="--t:{color}"><div class="rg-l"><span class="lt">{letter}</span><span class="ct">{count}</span>'
            f'<span class="bd">{band}</span></div><div class="rg-rows">{"".join(rows)}</div></div>')


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
@media(max-width:1023px){.st .st-h{display:none}.st .stc,.st.nobar .stc{grid-template-columns:28px minmax(0,1fr) auto}.st-r .br,.st-r .f:not(.p){display:none}.st-r .f.p .m{display:block;font-size:13px;font-weight:400;color:var(--muted)}.st-r .nm .sub .m{display:inline}.st.open .st-r .br{display:block;grid-column:2/-1;padding-top:6px}.st.open .stc{row-gap:0}}
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
.cell .gn{display:flex;flex-direction:column;align-items:flex-end;width:144px;min-width:0;justify-self:end}
.cell .gn .v{font:700 16px/20px var(--font-body);font-variant-numeric:tabular-nums;color:var(--strong)}.cell .gn .v small{font-size:13px;font-weight:600;color:var(--muted);margin-left:4px}
.cell .gn .bs{max-width:100%;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;font-size:12px;line-height:16px;color:var(--muted);text-align:right}
.cell .ic0{width:40px;height:40px;border-radius:5px;border:1px dashed var(--line-soft)}
'''

RARITY = {'poor': '#9d9d9d', 'common': '#ffffff', 'uncommon': '#1eff00', 'rare': '#3d94f0', 'epic': '#b866f5', 'legendary': '#ff8000'}
_RARITY_ICON = {'poor': 'q0', 'uncommon': 'q2', 'rare': 'q3', 'epic': 'q4', 'legendary': 'q5'}


def _colored(text: str, rarity: str) -> str:
    return f'<span style="color:{RARITY.get(rarity, RARITY["common"])}">{escape(text)}</span>'


def PaperdollCell(c: dict) -> str:
    """One slot. c: slot facts (label, flag, state, name, quality, icon, stats, line2, gain, basis), as bis_facts.mjs reads them."""
    flag = f'<span class="fl">{escape(c["flag"]["text"])}</span>' if c['flag'] else ''
    lead = f'<span class="lb">{escape(c["label"])}{flag}</span>'
    if c['state'] == 'empty':
        return (f'<div class="cell" role="option" aria-disabled="true">{lead}<span class="ic0"></span><span class="bd">'
                f'<span class="l1"><span class="nm muted" style="font-weight:400">{escape(c["name"])}</span></span><span class="l2">{escape(c["line2"])}</span></span><span></span></div>')
    ic = icon(c['icon'], 40, _RARITY_ICON.get(c['quality'], '')) if c['icon'] else '<span class="ic0"></span>'
    stats = f'<span class="st">{escape(c["stats"])}</span>' if c['stats'] else ''
    if c['gain'] is None and c['basis'] is None:   # `worn` mode (the sim): no gain column, as the built cell's empty auto track
        return (f'<div class="cell" role="option" aria-selected="false">{lead}{ic}<span class="bd"><span class="l1">'
                f'<span class="nm" style="color:{RARITY.get(c["quality"], RARITY["common"])}">{escape(c["name"])}</span>{stats}</span>'
                f'<span class="l2">{escape(c["line2"])}</span></span><span class="gn0"></span></div>')
    g = c['gain'] or {'value': '', 'unit': ''}
    unit = f'<small>{escape(g["unit"])}</small>' if g['unit'] else ''
    b = c['basis']
    basis = f'<span class="bs">{escape(b["prefix"])} {_colored(b["item"], b["quality"]) if b["item"] else ""}</span>' if b else ''
    return (f'<div class="cell" role="option" aria-selected="false">{lead}{ic}<span class="bd"><span class="l1">'
            f'<span class="nm" style="color:{RARITY.get(c["quality"], RARITY["common"])}">{escape(c["name"])}</span>{stats}</span>'
            f'<span class="l2">{escape(c["line2"])}</span></span><span class="gn"><span class="v">{escape(g["value"])}{unit}</span>{basis}</span></div>')


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
.drop{display:flex;flex-direction:column;align-items:center;justify-content:center;gap:6px;min-height:168px;padding:20px;border:1px dashed var(--line);border-radius:6px;text-align:center;color:var(--muted);font-size:14px}
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
        return f'<a href="#" class="more {cls}{" cur" if on else ""}" aria-haspopup="menu" aria-expanded="false">{lab} &#9662;</a>'
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


def Capture(src: str, w: int, h: int, alt: str, tag: str = '', tall: bool = False, lazy: bool = True) -> str:
    """A client screenshot at its own aspect, capped at its native width and 420 px high (900 for a portrait capture)."""
    t = f'{Tag(tag, "muted")}' if tag else ''
    lz = ' loading="lazy"' if lazy else ''
    # The picture links to its own file (spec 3A), as the built Capture does.
    return (f'<div class="cap-wrap"><a class="cap" href="{src}" style="--w:{w}px;--ar:{w} / {h};--mh:{900 if tall else 420}px">'
            f'<img src="{src}" width="{w}" height="{h}" alt="{escape(alt, quote=True)}"{lz} decoding="async"></a>{t}</div>')


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
@media(max-width:1023px){.pcol .st .stc{grid-template-columns:28px minmax(0,1fr) minmax(0,96px) auto}.pcol .st.open .st-r .br{display:block;grid-column:auto;padding-top:0}}.pcol .rows.slim .row{padding-block:7px}
.pcol-f{margin-top:auto;display:flex;gap:12px;flex-wrap:wrap;align-items:center;padding-top:4px}
.trees.cmp{gap:16px;grid-template-columns:repeat(3,max-content)}.trees.cmp .tg>[role=row]{display:contents}.trees.cmp .tc{display:block}.trees.cmp .tr h3{font-size:13px;line-height:18px;padding-bottom:6px;white-space:nowrap}
.trees.cmp .tg{grid-template-columns:repeat(4,24px);gap:4px}.trees.cmp .tc,.trees.cmp .tc img,.trees.cmp .tc .ic{width:24px!important;height:24px!important}
.trees.cmp .tc b{font-size:10px;line-height:12px;right:-3px;bottom:-4px}
@media(min-width:1760px){.trees.cmp .tg{grid-template-columns:repeat(4,28px)}.trees.cmp .tc,.trees.cmp .tc img,.trees.cmp .tc .ic{width:28px!important;height:28px!important}}
@media(max-width:767px){.trees.cmp .tr h3{white-space:normal;flex-wrap:wrap;column-gap:6px;row-gap:0;min-height:42px;align-content:flex-end}.trees.cmp{grid-template-columns:repeat(3,minmax(0,max-content));gap:10px}.trees.cmp .tg{grid-template-columns:repeat(4,22px);gap:3px}.trees.cmp .tc,.trees.cmp .tc img,.trees.cmp .tc .ic{width:22px!important;height:22px!important}}
.lvl{display:flex;align-items:center;gap:10px}.lvl .bt{all:unset;box-sizing:border-box;width:44px;height:44px;display:inline-flex;align-items:center;justify-content:center;border:1px solid var(--line);border-radius:4px;color:var(--nav);font-size:18px}
.lvl .bt:hover{color:var(--strong);border-color:var(--gold)}.lvl .bt.off{opacity:.45}
.lvl .rg{position:relative;flex:1;height:44px;display:flex;align-items:center}.lvl .rg .tk{height:4px;width:100%;border-radius:2px;background:var(--line-soft)}
.lvl .rg .fl{position:absolute;left:0;height:4px;border-radius:2px;background:var(--gold)}
.lvl .rg .th{position:absolute;width:20px;height:20px;margin-left:-10px;border-radius:50%;background:var(--gold);box-shadow:0 0 0 4px #e5b95533}
.rows.slim .row{min-height:44px;padding-block:6px}
/* MUST 5 (guides): the open phone bar wraps a segmented control that does not fit, so the current choice is never cut off */
.bar-m .seg{flex-wrap:wrap}.bar-m .seg a{min-width:max-content}
.ans-box{display:flex;flex-direction:column;gap:6px;min-width:0}
@media(min-width:900px){.ans-box{min-width:380px;min-height:96px;justify-content:flex-end;align-items:flex-end;text-align:right}}
'''


_PageHeader0 = PageHeader
_Page0 = Page


def PageHeader(title: str, title_color: str = '', crest: str = '', meta: list[str] | None = None, answer: str = '', actions: str = '', tabs: str = '',
               lead: str = '') -> str:
    """PageHeader with the optional Lead sentence under the meta line (inside head-title, so the regions are unchanged)."""
    html = _PageHeader0(title, title_color, crest, meta, answer, actions, tabs)
    if not lead:
        return html
    at = re.search(r'<div class="ph-meta" data-region="head-meta">.*?</span></div>', html) if meta else None
    cut = at.end() if at else html.index('</h1>') + len('</h1>')
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


def ReadingSection(title: str, prose: str, lane: str = '', anchor: str = '', action: str = '') -> str:
    """A section of a reading page: the Section head (22/28), then prose and lane (lane first in the DOM: one column reads it first)."""
    a = f'<span class="act">{action}</span>' if action else ''
    lane_html = f'<div class="rd-lane">{lane}</div>' if lane else ''
    return (f'<section class="rsec" data-region="{region_slug(title)}" id="{anchor}"><div class="sec-h"><h2>{title}</h2>{a}</div>'
            f'<div class="rd{"" if lane else " solo"}">{lane_html}<div class="rd-prose">{prose}</div></div></section>')


def LaneBlock(label: str, body: str, meta: str = '') -> str:
    """One labelled block in a lane (12 px caps label, then its rows or table)."""
    m = f'<span class="t-small">{meta}</span>' if meta else ''
    return f'<div class="lane-blk"><div class="pcol-h"><h3 class="t-label">{label}</h3>{m}</div>{body}</div>'


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
    return Trees(trees).replace('<div class="trees">', '<div class="trees cmp">', 1)


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
PRIM_CSS += ''.join(f'.grade.t{k}{{background:{v}}}' for k, v in TIER_TONES.items()) + '\n'


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


def GradeTile(letter: str) -> str:
    return f'<span class="grade t{letter}" role="img" aria-label="Tier {letter}">{letter}</span>'

"""Round 4 mocks: home, best in slot and the tier list in Blizzard's Forever colours (design/specs/2026-11-09-directions.md).

Home leads with the guild's raid night; BiS opens on the paperdoll standing alone (the selected item's tooltip where
the character would be) with the itinerary by dungeon below; the tier list is the shipped ranking (ws_tiers.py).
The road map is gone (owner, round 4); its silhouette was tried in the home band and cut on review (read as decoration).
Copy rule (ai-slop-reviewer tell 22): a number takes its unit and nothing else.
    python3 design/mocks/directions/pages.py
"""
import re
import sys
from html import escape
from pathlib import Path

import data as D
from common import DOLL_BOTTOM, DOLL_LEFT, DOLL_RIGHT, doll_cells, next_best, sentence, write
from art import cave_uri
import paperdoll as PD
from style import footer, header, icon, medallion, page, titlebar, tooltip

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
import ws_tiers  # noqa: E402  the shipped tier rules

HOME_FEED_ROWS = 6
UPGRADES_SHOWN = 5
RANGE = re.compile(r'\((\d+)-(\d+)\)')
TIER_FACTS = 'Best in slot, raid buffs · level 63 · 3 min'
BAR_FLOOR = 0.9  # tier bars start at 90% of the lowest spec, so the field's spread shows


def dps_ranking() -> list[dict]:
    rows, _ = ws_tiers.load('alliance')
    out, _ = ws_tiers.ranked(rows, 'dps')
    return out


def spec_name(r: dict) -> str:
    return f'{r["name"]} {ws_tiers.CLASS_NAME[r["cls"]]}'


# ---------- home ----------

HOME_CSS = '''
.home{display:grid;grid-template-columns:minmax(0,1fr) 400px;grid-template-areas:"lead lead" "feed rail" "sims sims";gap:28px;padding-top:28px;align-items:start}
.rail{grid-area:rail}.feedsec{grid-area:feed}.lead{grid-area:lead}
/* the raid plaque: the owner's cave art behind it, cave mouth on the right, under a night veil so the type stays AA */
.raid{padding:26px 30px;background:linear-gradient(rgba(14,9,6,.46),rgba(14,9,6,.46)),var(--cave) 35% 40%/160% auto no-repeat,var(--card)}
@media(min-width:900px){.raid{display:grid;grid-template-columns:minmax(0,1fr) minmax(0,1fr);gap:0 34px}.raid .prep{grid-column:2;grid-row:1/span 5;
border-left:1px solid rgba(255,255,255,.12);padding-left:28px;margin-top:0;border-top:0;padding-top:0}}
.raid .when{font:400 30px/1.15 var(--title);color:var(--cream);text-shadow:0 1px 2px rgba(0,0,0,.45)}.raid .sub{color:var(--text);font-size:15px}
.count{display:flex;flex-wrap:wrap;gap:26px;margin-top:16px}.count b{font:600 44px/1 var(--body);color:#fff;display:block}.count span{font-size:13px;color:var(--muted)}
.count .sep{width:1px;background:rgba(255,255,255,.18);align-self:stretch}
.roles{display:flex;flex-wrap:wrap;gap:6px 18px;margin-top:14px;font-size:14px}.roles b{color:#fff}
.prep{margin-top:16px;border-top:1px solid rgba(255,255,255,.12);padding-top:12px}.prep h3{font-size:19px;margin-bottom:8px}
.prep li{display:flex;gap:10px;align-items:baseline;padding:3px 0}
.prep li::before{content:"";width:11px;height:11px;border:1px solid var(--rule-warm);flex:none;transform:translateY(1px)}
.feedsec .sh,.rail .sh{margin-bottom:10px}
.feed{padding:4px 20px}.feed li{display:grid;grid-template-columns:30px 1fr auto;gap:12px;align-items:center;padding:10px 0;border-bottom:1px solid rgba(255,255,255,.1)}
.feed li:last-child{border-bottom:0}.feed b{font-weight:700;color:#fff}.feed .t{font-size:13px;color:var(--muted);white-space:nowrap}
.feed .ic,.feed .crest{width:30px;height:30px}.crest{display:block;border-radius:50%}
.simsec{grid-area:sims}.simsec .lad{columns:2;column-gap:48px}.simsec .lad li{break-inside:avoid}.simsec .lad li.me{margin-top:0}
@media(max-width:700px){.simsec .lad{columns:1}}
.lad li{display:grid;grid-template-columns:20px 30px 1fr auto;gap:10px;align-items:center;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.1)}
.lad li:last-child{border-bottom:0}.lad li.me{border-top:1px solid rgba(255,255,255,.22);margin-top:4px}.lad .r{color:var(--muted)}.lad .v{color:#fff}
.spec{margin-top:12px}
@media(max-width:1100px){.home{grid-template-columns:1fr;grid-template-areas:"lead" "rail" "feed" "sims"}}
@media(max-width:700px){.raid{padding:20px 18px;background:linear-gradient(rgba(14,9,6,.62),rgba(14,9,6,.62)),var(--cave) 59% 40%/cover no-repeat,var(--card)}.raid .when{font-size:25px}.count b{font-size:26px}.count{gap:16px}.feed{padding:2px 14px}}
'''

GUILD_SIM_TAIL = re.compile(r', #(\d+) of (\d+) DPS in [^,(]+ \([^)]*\)$')


def feed_lead(row: dict) -> str:
    if row.get('crest'):
        return f'<img class="crest" src="{D.crest_uri(row["crest"])}" alt="">'
    return icon(row.get('icon') or '', 30)


def home_sentence(row: dict) -> str:
    """A guild sim on your own guild's home reads '#2 of 14': the guild, spec and band are the page's (tell 22)."""
    if row['kind'] != 'guild_sim':
        return sentence(row)
    segs = [dict(s, text=GUILD_SIM_TAIL.sub(r', #\1 of \2', s['text'])) if s.get('style') is None else s for s in row['segments']]
    return sentence({**row, 'segments': segs})


def raid_card(g: dict) -> str:
    r, s = g['next_raid'], g['standing']
    c = r['counts']
    days, hours, _ = D.countdown()
    roles = ''.join(f'<span><b>{c[k]}</b> {k}</span>' for k in ('tanks', 'healers', 'melee', 'ranged'))
    needs = ''.join(f'<li>{escape(n)}</li>' for n in s['needs_before_next_raid'])
    return (f'<section class="card framed ornate raid" style="--cave:url({cave_uri()})"><div class="when">{escape(r["zone_name"])}</div>'
            f'<div class="sub">{escape(g["guild"]["name"])} · Wednesday 9 December · 20:00</div>'
            f'<div class="count num"><div><b>{days}</b><span>days</span></div><div><b>{hours}</b><span>hours</span></div><i class="sep"></i>'
            f'<div><b>{c["in"]}</b><span>signed up</span></div><div><b>{c["late"]}</b><span>late</span></div></div>'
            f'<div class="roles">{roles}</div>'
            f'<div class="btnrow" style="margin-top:18px"><a class="btn" href="#">I\'m in</a><a class="btn2 on" href="#">Late</a><a class="btn2" href="#">Can\'t come</a></div>'
            f'<div class="prep"><h3>Before the pull</h3><ul>{needs}</ul></div></section>')


def home() -> str:
    h = D.home('60')
    c, home_ = h['char'], h['home']
    ranking = dps_ranking()
    fury = next(r for r in ranking if r['spec'] == 'warrior-fury')
    rows = [x for x in h['feed'] if x['kind'] != 'guild_raid'][:HOME_FEED_ROWS]  # the raid post is the raid card already
    feed = ''.join(f'<li>{feed_lead(x)}<span>{home_sentence(x)}</span><span class="t">{D.ago(x["at"])}</span></li>' for x in rows)
    ups = ''.join(f'<div class="gear">{icon(u["icon"])}<div><span class="n" style="color:var(--q-{u["quality"]})">{escape(u["name"])}</span>'
                  f'<span class="src">{D.SLOT_NAME[u["slot"]]} · {escape(u["source"])}</span></div><div class="d">+{u["gain"]:.1f} DPS</div></div>'
                  for u in h['upgrades'][:UPGRADES_SHOWN])
    gs = home_['guild_sims']
    lrow = lambda x, cls='': (f'<li class="{cls}"><span class="r num">{x["rank"]}</span>{medallion(x["class"], 30)}'
                              f'<span style="color:{D.CLASS_COLOR[x["class"]]};font-weight:600">{escape(x["name"])}</span><span class="v num">{x["value"]:.1f} DPS</span></li>')
    # Round 5: the hero art gets the band to itself; the raid plaque leads the page under it on the cave art (the
    # ornate frame); upgrades are the second featured panel (the plain frame); the feed and guild sims stay plain.
    rail = f'<aside class="rail card framed pad"><h2 class="sh"><a href="#" style="color:inherit;font-weight:400">Your upgrades</a></h2>{ups}</aside>'
    sims = (f'<section class="simsec"><h2 class="sh"><a href="#" style="color:inherit;font-weight:400">Guild sims</a></h2><div class="card pad"><ol class="lad">'
            f'{"".join(lrow(x) for x in gs["rows"])}{lrow(gs["you"], "me")}</ol>'
            f'<p class="spec">Tier {fury["tier"]} · #{fury["rank"]} of {len(ranking)} · <a href="#">Tier list</a></p></div></section>')
    main = (f'<div class="lead">{raid_card(h["guild"])}</div>{rail}'
            f'<section class="feedsec"><h2 class="sh">Latest on Forever</h2><ul class="card feed">{feed}</ul></section>{sims}')
    body = (titlebar(h['slug'], c['name'], f'Level 60 Human Fury Warrior · &lt;{escape(c["guild"]["name"])}&gt;',
                     f'<b class="num">{home_["import_sim"]["dps"]}</b> DPS', D.CLASS_COLOR[h['slug']],
                     top=header(c['name'], h['slug'], c['level'], ''), hero=True)
            + f'<main class="w home">{main}</main>' + footer())
    return page('Forever Sixty, home', HOME_CSS, body)


# ---------- best in slot ----------

BIS_CSS = '''
.doll{display:grid;grid-template-columns:minmax(0,1fr) 340px minmax(0,1fr);gap:8px 30px;padding:20px 22px;margin-top:28px}
.slot{display:grid;grid-template-columns:40px 1fr;gap:10px;align-items:center;min-height:50px}
.slot .n{display:block;font-weight:600;font-size:14px;line-height:18px}.slot .s{display:block;font-size:12px;line-height:16px;color:var(--muted)}
.slot.none .ic{opacity:.45}.slot.none .s{font-size:13px}
.dr .slot{grid-template-columns:1fr 40px;text-align:right}.dr .slot .ic{order:2}
.slot.sel .ic{box-shadow:0 0 0 2px var(--cream)}
.mid{align-self:start;padding-top:6px}.mid p{margin-top:12px;font-size:14px}.mid .lk{margin-top:10px}
/* weapons sit in the centre column under the tooltip, where the model would stand (ai-slop-reviewer, round 5: no dead centre) */
.dl,.dr{grid-row:1/span 2}.mid{grid-column:2;grid-row:1}
.wpn{grid-column:2;grid-row:2;align-self:end;display:flex;flex-direction:column;gap:6px;border-top:1px solid rgba(255,255,255,.1);padding-top:12px}
.wpn .slot{grid-template-columns:40px auto}
.stops{margin-top:30px;padding:8px 24px 18px}
.stop{padding-top:16px}.stop h2{font-size:22px;display:flex;align-items:baseline;gap:10px;padding-bottom:6px;border-bottom:1px solid rgba(255,255,255,.16)}
.stop h2 span{font:600 14px var(--body);color:var(--text)}
/* narrow: the doll packs into a grid of sockets, the selected item's tooltip right under it (ux-designer, round 4) */
@media(max-width:1100px){.doll{grid-template-columns:1fr;gap:10px;padding:16px}.doll .dl,.doll .dr,.doll .wpn,.doll .mid{grid-row:auto;grid-column:auto}.doll .wpn{flex-direction:row;align-self:auto}.doll .dl,.doll .dr,.doll .wpn{display:flex;flex-wrap:wrap;gap:8px;
justify-content:flex-start;border:0;padding:0}.doll .slot{display:block;min-height:0}.doll .slot>div{display:none}.doll .dr .slot .ic{order:0}
.doll .dl{order:1}.doll .dr{order:2}.doll .wpn{order:3}.doll .mid{order:4;border-top:1px solid rgba(255,255,255,.1);padding-top:14px}}
@media(max-width:700px){.doll .ic{width:44px;height:44px}.stops{padding:4px 14px 14px}}
'''


def place_range(group: dict) -> tuple[int, int] | None:
    m = RANGE.search(' '.join(c['line2'] for c in group['cells'] if f'{group["title"]} (' in c['line2']))
    return (int(m.group(1)), int(m.group(2))) if m else None


def group_heading(title: str) -> str:
    return 'Crafted' if title.startswith('Crafted') else title


PAIRED_SLOTS = {'finger1', 'finger2', 'trinket1', 'trinket2'}  # two of a kind: the gain must say which one it replaces


def gain_line(cell: dict) -> str:
    """The tooltip's gain: '+0.3 DPS', and 'over <item>' only for a ring or trinket (owner rule)."""
    gain = f'{cell["gain"]["value"]} DPS'
    other = (cell.get('basis') or {}).get('item', '')
    return f'{gain} over {other}' if cell['slot'] in PAIRED_SLOTS and other else gain


def two_handed(facts: dict, cells: dict) -> bool:
    """The set's main hand is a two-hander when the BiS notes say so (the facts carry it as a note, not a flag)."""
    main = cells.get('main_hand')
    return bool(main) and any(main['name'] in n and 'two-hander' in n for n in facts['notes'])


def slot_label(slot: str, two_hand: bool) -> str:
    return 'Two-Hand' if slot == 'main_hand' and two_hand else D.SLOT_NAME[slot]


def source_line(cell: dict, group_title: str, two_hand: bool) -> str:
    """What the row adds to its heading: the boss and chance for a dungeon, the profession, the standing, the quest name."""
    parts = []
    for p in cell['line2'].split(' · '):
        if p == group_title or p.startswith(f'{group_title} (') or p == 'Crafted' or (p == 'BoE' and 'BoE' in group_title):
            continue
        parts.append(p.removeprefix('Quest: '))
    return ' · '.join([slot_label(cell['slot'], two_hand), *parts])


def doll_slot(cells: dict, slot: str, sel: str, two_hand: bool) -> str:
    c = cells.get(slot)
    if slot == 'off_hand' and two_hand:  # a two-hander takes both hands: no off-hand socket (wow-player, round 4)
        return ''
    if not c or c['state'] != 'item':
        return f'<div class="slot none">{icon("")}<div><span class="s">{D.SLOT_NAME[slot]}</span></div></div>'
    return (f'<div class="slot{" sel" if slot == sel else ""}">{icon(c["icon"])}<div><span class="n" style="color:var(--q-{c["quality"]})">{escape(c["name"])}</span>'
            f'<span class="s">{slot_label(slot, two_hand)}</span></div></div>')


def gear_row(c: dict, group_title: str, sel: str, two_hand: bool) -> str:
    gain = (c.get('gain') or {}).get('value', '')
    return (f'<div class="gear{" sel" if c["slot"] == sel else ""}">{icon(c["icon"])}<div><span class="n" style="color:var(--q-{c["quality"]})">{escape(c["name"])}</span>'
            f'<span class="src">{escape(source_line(c, group_title, two_hand))}</span></div><div class="d">{gain + " DPS" if gain.startswith("+") else ""}</div></div>')


def bis() -> str:
    f = D.bis('low-20')
    cells = doll_cells(f)
    sel = 'back'
    groups = sorted(f['doll']['groups'], key=lambda g: (place_range(g) or (99, 99))[0])
    c = cells[sel]
    th = two_handed(f, cells)
    mid = f'<div class="mid">{tooltip("hunter", c["item"], gain_line(c))}<p class="lk"><a href="#">Open in the sim ›</a></p></div>'
    slots = lambda names: ''.join(doll_slot(cells, s, sel, th) for s in names)
    doll = (f'<section class="card framed ornate doll"><div class="dl">{slots(DOLL_LEFT)}</div>{mid}<div class="dr">{slots(DOLL_RIGHT)}</div>'
            f'<div class="wpn">{slots(DOLL_BOTTOM)}</div></section>')
    stops = ''.join(f'<section class="stop"><h2>{escape(group_heading(g["title"]))}{f"<span>{place_range(g)[0]} to {place_range(g)[1]}</span>" if place_range(g) else ""}</h2>'
                    + ''.join(gear_row(x, g['title'], sel, th) for x in g['cells']) + '</section>' for g in groups)
    ctl = [[(o['label'], o['on']) for o in g['options']] for g in f['controls']['groups'] if g['kind'] == 'segmented']
    a = f['answer']['figure']
    body = (titlebar('hunter', 'Beast Mastery Hunter', 'Best in slot', f'<b class="num">{a["value"]}</b> DPS', D.CLASS_COLOR['hunter'], ctl,
                     top=header('Thandrel', 'warrior', 60, 'Best in slot'))
            + f'<main class="w">{doll}<div class="card stops">{stops}</div></main>' + footer())
    return page('Forever Sixty, best in slot', BIS_CSS, body)


# ---------- best in slot on the owner's paperdoll frame ----------

FRAME_CSS = PD.CSS + '''
.pdpage{display:grid;grid-template-columns:minmax(0,640px) minmax(0,1fr);gap:36px;margin-top:28px;align-items:start}
.pdside .card+.card{margin-top:22px}.pdside h2{font-size:21px;margin-bottom:8px}.pdside p{font-size:14px}
.wts li{display:grid;grid-template-columns:1fr auto;padding:5px 0;border-bottom:1px solid rgba(255,255,255,.1)}.wts li:last-child{border-bottom:0}.wts b{color:#fff}
.unit{margin-top:10px;font-size:13px;color:var(--muted)}.idx span{color:var(--muted)}
.pd .ctr .det{max-width:340px;width:100%;font-size:13.5px;line-height:1.5;color:rgba(255,255,255,.88)}
/* signed in: the character's Battle.net render stands in the centre like the client's model; the tooltip sits at its foot.
   Sized from the render's own bounds (figure 678 of 1200 px tall, its centre at 50.6% across, its feet at 81.4% down) so the
   figure fills about 85% of the panel and stands near its foot. */
.pd .ctr.model{justify-content:flex-end}
.pd .ctr .rd{position:absolute;inset:0;background:var(--render) 51% 54%/auto 150% no-repeat}
.pd .ctr.model>*:not(.rd):not(.who){position:relative;z-index:1}
.pd .ctr.model .who{top:3cqw;bottom:auto;z-index:1}
.pd .ctr.model .who{text-shadow:0 1px 2px #000,0 0 6px #000}.pd .ctr .who .as{font-weight:400;font-style:italic;color:rgba(255,255,255,.65)}
.pdafter{display:flex;align-items:baseline;justify-content:space-between;gap:16px;margin-top:12px;padding:12px 16px;
  background:var(--card);border-radius:var(--r-sm);font-size:14px;color:rgba(255,255,255,.88)}
.pdafter .det{margin:0}
.pdafter .lk{font:600 14px var(--body);white-space:nowrap}
.stops{margin-top:30px;padding:8px 24px 18px}
.stop{padding-top:16px}.stop h2{font-size:22px;display:flex;align-items:baseline;gap:10px;padding-bottom:6px;border-bottom:1px solid rgba(255,255,255,.16)}
.stop h2 span{font:600 14px var(--body);color:var(--text)}
@media(max-width:1100px){.pdpage{grid-template-columns:1fr}.pd{margin:0 auto}}
/* the frame holds while its slots are at least 36 px; below that, its header and the two-column list */
@media(max-width:''' + str(PD.MIN_FRAME_PX + 28) + '''px){.pd,.pdafter{display:none}.pdl{display:block}.stops{padding:4px 14px 14px}}
'''


def controls_band(ctl: list, top: str) -> str:
    """The band without a title: the frame's title bar carries the page's h1, so the band holds the nav and the controls."""
    groups = ''.join('<div class="g">' + ''.join(f'<a class="{"on" if on else ""}" href="#">{escape(lb)}</a>' for lb, on in g) + '</div>' for g in ctl)
    return f'<div class="band titlebar">{top}<div class="w"><nav class="ctl" aria-label="Page controls" style="border-top:0;margin-top:10px">{groups}</nav></div></div>'


# The owner's own character (coordinator, 2026-11-08): Blizzard's render, referenced by URL, never committed.
OWNER_CHAR = 'Frickin Laserbeams'
OWNER_MEDIA = 'https://render.worldofwarcraft.com/forever-us/character/204/6928844'
OWNER_RENDER, OWNER_AVATAR = f'{OWNER_MEDIA}-main-raw.png', f'{OWNER_MEDIA}-avatar.jpg'
OWNER_LEVEL = 31  # production, per the coordinator


def bis_frame(signed_in: bool) -> str:
    """BiS on the owner's paperdoll frame (owner, 2026-11-09). The centre holds the selected slot's client tooltip, where
    it drops and what it replaces; hover or tap on a slot swaps it. Signed in, the character's Battle.net render
    (characters.render_url) stands behind the tooltip with the name above, as the client's pane shows the model. The mock
    uses the owner's own character, referenced by URL and never copied into the repo."""
    f = D.bis('low-20')
    cells = doll_cells(f)
    th = two_handed(f, cells)
    sel = 'back'
    a = f['answer']['figure']
    sub = '20 to 29 &middot; Alliance'
    c = cells[sel]
    detail = f'<p class="det">{escape(c["line2"].replace(" · ", ", "))}<br>{next_best(c)}</p>'
    who = f'<p class="who" style="color:{D.CLASS_COLOR["hunter"]}">{OWNER_CHAR}<small>Level {OWNER_LEVEL} Night Elf Hunter</small><small class="as">As last logged out</small></p>' if signed_in else ''
    tool_tip = tooltip("hunter", c["item"], gain_line(c))
    after = f'{detail}<a class="lk" href="#">Open in the sim &rsaquo;</a>'
    tip = tool_tip + after
    # Signed in, only the tooltip stands over the model; where it drops and the sim link go under the frame.
    centre = tool_tip + who if signed_in else tip
    avatar, render = (OWNER_AVATAR, OWNER_RENDER) if signed_in else ('', '')
    doll = PD.frame(cells, sel, th, 'Beast Mastery Hunter', sub, a['value'], centre, avatar, render)
    if signed_in:
        doll += f'<div class="pdafter">{after}</div>'
    listed = PD.list_form(cells, sel, th, 'Beast Mastery Hunter', a['value'], f'<div class="tip-wrap" style="margin-top:12px">{tip}</div>', avatar=avatar)
    w = f['weights']
    ref = w['rows'][0]
    weights = ''.join(f'<li><span>{escape(r["name"])}</span><b class="num">{escape(r["scale"])}</b></li>' for r in w['rows'][:4] if r['scale'])
    groups = sorted(f['doll']['groups'], key=lambda g: (place_range(g) or (99, 99))[0])
    index = ''.join(f'<li><a href="#">{escape(group_heading(g["title"]))}</a><span class="num">{len(g["cells"])}</span></li>' for g in groups)
    side = (f'<aside class="pdside"><section class="card pad"><h2>What to stack</h2><ol class="wts">{weights}</ol>'
            f'<p class="unit">1.00 = {escape(ref["line"])} of {escape(ref["name"])}</p><p class="unit">Impaling Harpoon is two-handed: the off hand stays empty.</p></section>'
            f'<section class="card pad"><h2>Where to go</h2><ol class="wts idx">{index}</ol></section>'
            f'</aside>')
    stops = ''.join(f'<section class="stop"><h2>{escape(group_heading(g["title"]))}{f"<span>{place_range(g)[0]} to {place_range(g)[1]}</span>" if place_range(g) else ""}</h2>'
                    + ''.join(gear_row(x, g['title'], sel, th) for x in g['cells']) + '</section>' for g in groups)
    ctl = [[(o['label'], o['on']) for o in g['options']] for g in f['controls']['groups'] if g['kind'] == 'segmented']
    who = (OWNER_CHAR, 'hunter', OWNER_LEVEL) if signed_in else ('Thandrel', 'warrior', 60)
    body = (controls_band(ctl, header(*who, 'Best in slot'))
            + f'<main class="w" style="--frame:url({PD.frame_uri()})"><div class="pdpage"><div>{doll}{listed}</div>{side}</div>'
            f'<div class="card stops">{stops}</div></main>' + footer())
    return page('Forever Sixty, best in slot', FRAME_CSS, body)


# ---------- tier list ----------

TIER_CSS = '''
.tiers{display:grid;grid-template-columns:minmax(0,1fr) 340px;gap:28px;padding-top:28px;align-items:start}
.cols{display:grid;grid-template-columns:64px 28px 34px 1fr 160px 76px 64px;gap:12px;padding:9px 18px 9px 0;font-size:12px;color:var(--muted);border-bottom:1px solid #000}
.cols span:nth-child(n+5){text-align:right}
.tier{display:grid;grid-template-columns:64px 1fr;border-bottom:1px solid #000}.tier:last-child{border-bottom:0}
.tier .L{font:400 40px/1 var(--title);color:var(--cream);padding:14px 0 0 18px;border-right:1px solid rgba(255,255,255,.1)}
.tier ol{padding:6px 18px 6px 12px}
.tier li{display:grid;grid-template-columns:28px 34px 1fr 160px 76px 64px;gap:12px;align-items:center;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.08)}
.tier li:last-child{border-bottom:0}.tier .r{color:var(--muted)}.tier .v{font-weight:600;text-align:right;color:#fff}.tier .gap{color:var(--muted);text-align:right;font-size:13px}
.rel{height:6px;background:rgba(255,255,255,.07)}.rel i{display:block;height:100%;background:var(--blue)}
.tier li.me{background:rgba(14,134,202,.14);box-shadow:-12px 0 0 rgba(14,134,202,.14),18px 0 0 rgba(14,134,202,.14)}
.yours{position:sticky;top:16px}.yours h2{font-size:22px;margin-bottom:6px}.yours .pos{color:#fff;font-weight:600}
.yours p{margin-top:8px}.yours .btnrow{margin-top:16px}
@media(max-width:1100px){.tiers{grid-template-columns:1fr}.tiers>aside{order:-1}.yours{position:relative;top:auto}}
@media(max-width:700px){.cols{display:flex;justify-content:flex-end;padding:7px 10px}.cols span{display:none}.cols .dps{display:inline}
.tier{grid-template-columns:44px 1fr}.tier .L{font-size:30px;padding-left:10px}.tier ol{padding:4px 10px 4px 8px}
.tier li{grid-template-columns:22px 30px 1fr auto;gap:8px}.tier .rel,.tier .gap{display:none}}
'''


def bar_width(r: dict, out: list[dict]) -> float:
    floor = min(x['metric'] for x in out) * BAR_FLOOR
    return (r['metric'] - floor) / (out[0]['metric'] - floor) * 100


def gap_text(r: dict) -> str:
    return f'−{r["gap"]:.1f}%' if r['gap'] else ''


def tiers() -> str:
    out = dps_ranking()
    you = next(r for r in out if r['spec'] == 'warrior-fury')
    groups = []
    for letter in ('S', 'A', 'B', 'C', 'D'):
        rows = [r for r in out if r['tier'] == letter]
        if not rows:
            continue
        lis = ''.join(f'<li class="{"me" if r is you else ""}"><span class="r num">{r["rank"]}</span>{medallion(r["cls"], 34)}'
                      f'<span style="color:{D.CLASS_COLOR[r["cls"]]};font-weight:600">{escape(spec_name(r))}</span>'
                      f'<span class="rel"><i style="width:{bar_width(r, out):.1f}%"></i></span>'
                      f'<span class="v num">{r["metric"]:.1f}</span><span class="gap num">{gap_text(r)}</span></li>' for r in rows)
        groups.append(f'<div class="tier"><div class="L">{letter}</div><ol>{lis}</ol></div>')
    heads = '<div class="cols"><span></span><span>#</span><span></span><span>Spec</span><span></span><span class="dps">DPS</span><span>vs #1</span></div>'
    # Your spec: where it sits and where to go next; the table already gives its distance to #1 (tell 22, round 4).
    yours = (f'<section class="card framed pad yours"><h2>Fury Warrior</h2><p class="pos">Tier {you["tier"]} · <span class="num">#{you["rank"]}</span> · '
             f'<span class="num">{you["metric"]:.1f}</span> DPS</p>'
             f'<div class="btnrow"><a class="btn" href="#">Best in slot</a><a class="btn2" href="#">Guide</a></div></section>')
    ctl = [[('Damage', True), ('Tanks', False), ('Healers', False)], [('Alliance', True), ('Horde', False)]]
    body = (titlebar(None, 'Tier list at 60', TIER_FACTS, ctl=ctl, top=header('Thandrel', 'warrior', 60, 'Tier list'))
            + f'<main class="w tiers"><section class="card">{heads}{"".join(groups)}</section><aside>{yours}</aside></main>' + footer())
    return page('Forever Sixty, tier list', TIER_CSS, body)


def build() -> None:
    write('g-home', home())
    write('g-bis', bis())
    write('g-tiers', tiers())
    write('g-bis-frame', bis_frame(False))
    write('g-bis-frame-signed', bis_frame(True))


if __name__ == '__main__':
    build()

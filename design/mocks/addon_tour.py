"""Mock boards for the addon-tour lane (design/specs/2026-11-08-addon-ux.md section 6, 2026-11-11).

    python3 design/mocks/addon_tour.py              # every board
    python3 design/mocks/addon_tour.py builds-dialog-ask onboarding-5-ask

Writes design/reviews/2026-11-11-addon-tour/<board>.png at device scale 2, through the shared recipes
of addon_ux.py (the tour's chrome, the panel plate, the client's tooltip and button looks, the tokens).
Every string a board shows is read from addon/ForeverSixty/Locale.lua (locale()), never typed here;
every build row is read from addon/ForeverSixty/Data.lua (default_builds()), so a data refresh redraws
it. The worked characters: a level 13 paladin with 2 Holy and 2 Retribution points (the owner's own
capture that started the lane), the same paladin once decided, a level 60, and a druid for the
four-tile case.
"""
import os
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
import addon_ux as ux  # noqa: E402
from addon_ux import (ADDON, C, CSS, OB_H, OB_W, PAGES, RING, NOTCH, TITLE_FONT, U, b64, flow_arrow,  # noqa: E402
                      icon_uri, qr_block, url_box, page, tga_uri)

ROOT = Path(__file__).resolve().parents[2]
# Onboarding.lua coming(): L.tourComing ("· coming"), muted (round 3), 4 px after the line.
COMING_GAP = 4
OUT = Path(os.environ['MOCK_OUT']) if os.environ.get('MOCK_OUT') else ROOT / 'design/reviews/2026-11-11-addon-tour'
LOCALE = (ROOT / 'addon/ForeverSixty/Locale.lua').read_text()
DATA = (ROOT / 'addon/ForeverSixty/Data.lua').read_text()


# --- Locale.lua, read, never typed ------------------------------------------------------------------
def locale(key: str) -> str:
    """L.<key> = "..." (string literals joined by `..`, continuation lines included, and the SITE and
    L.siteName names Locale.lua concatenates), or `<key> = "..."` inside the L table."""
    m = re.search(r'(?:^L\.%s|^\s*%s) = (.*(?:\n\s*\.\. .*)*)' % (key, key), LOCALE, re.M)
    if m is None:
        raise KeyError(key)
    parts = []
    for literal, name in re.findall(r'"((?:[^"\\]|\\.)*)"|\b(SITE|L\.siteName)\b', m.group(1)):
        parts.append('foreversixty.gg' if name else literal)
    return ''.join(parts).replace('\\"', '"').replace('\\194\\183', '·')


def playstyle(slug: str) -> str:
    m = re.search(r'\["%s"\] = "([^"]*)"' % re.escape(slug), LOCALE)
    return m.group(1)


def ask(trees):
    """Follow.specQuestion's wording: L.pathAsk over the trees, joined by L.pathAskComma and L.pathAskOr."""
    names = list(trees)
    last = names.pop()
    joined = (locale('pathAskComma').join(names) + locale('pathAskOr') + last) if names else last
    return locale('pathAsk') % joined


# --- Data.lua, read, never typed ---------------------------------------------------------------------
def tabs(class_slug: str):
    block = DATA[DATA.index('["%s"] = {' % class_slug):]
    return re.findall(r'\{ name = "([^"]+)", talents = \{', block)[:3]


#: A level 60 holds 51 talent points; every default build's order spends them all.
POINTS_AT_60 = 51


def order_triples(body: str, kind: str):
    """(name, [(tab, tier, column)]) of one default build in a spec's Data.lua block, or None.

    The list is matched triple by triple up to its own closing brace. A lazy
    `\\{(.*?)\\} \\}` stopped inside the last triple and dropped it, which is
    how the 2026-11-11 tour boards showed raid builds at 50 points."""
    k = re.search(r'%s = \{ name = "([^"]+)", order = \{((?:\s*\{\d+,\d+,\d+\},?)*)\s*\}' % kind, body)
    if k is None:
        return None
    order = [tuple(int(x) for x in t) for t in re.findall(r'\{(\d+),(\d+),(\d+)\}', k.group(2))]
    if len(order) != POINTS_AT_60:
        raise ValueError(f'{kind} build {k.group(1)!r} reads {len(order)} points, not {POINTS_AT_60}')
    return k.group(1), order


def default_builds(class_slug: str, kind: str):
    """[(spec slug, data name, {tab: points})] in tree order, as BuildDialog.ownRows lists them."""
    rows = []
    start = DATA.index('defaultBuilds = {')
    for m in re.finditer(r'\["(%s-[a-z-]+)"\] = \{(.*?)\n\t\t\},' % class_slug, DATA[start:], re.S):
        spec, body = m.group(1), m.group(2)
        found = order_triples(body, kind)
        if found is None:
            continue
        name, order = found
        split = {}
        for tab, _, _ in order:
            split[tab] = split.get(tab, 0) + 1
        rows.append((spec, name, split))
    names = tabs(class_slug)
    order = {name.split()[0].lower(): i for i, name in enumerate(names)}
    rows.sort(key=lambda r: (order.get(r[0].split('-')[1], 9), r[0]))
    return rows


def talent_names(class_slug: str):
    """{(tab, tier, column): name} from Data.classes, the names BuildDialog shows beside a row."""
    block = DATA[DATA.index('["%s"] = {' % class_slug):]
    names = {}
    for tab, (_, body) in enumerate(re.findall(r'\{ name = "([^"]+)", talents = \{(.*?)\} \},', block, re.S)[:3], 1):
        for name, tier, column in re.findall(r'\{ name = "([^"]+)", tier = (\d+), column = (\d+)', body):
            names[(tab, int(tier), int(column))] = name
    return names


def default_order(class_slug: str, spec: str, kind: str):
    start = DATA.index('defaultBuilds = {')
    body = re.search(r'\["%s"\] = \{(.*?)\n\t\t\},' % re.escape(spec), DATA[start:], re.S).group(1)
    return order_triples(body, kind)[1]


def next_talent(class_slug: str, spec: str, kind: str, taken: int):
    """Follow.defaultNext for a character whose first `taken` points follow the order: the next point's talent."""
    order = default_order(class_slug, spec, kind)
    if taken >= len(order):
        return None
    return talent_names(class_slug)[order[taken]]


def split_of(class_slug: str, points: dict) -> str:
    """Follow.specSplit: "2 Holy · 2 Retribution", the trees that hold points in tab order."""
    names = tabs(class_slug)
    return locale('exportTreeSeparator').join(locale('pathSplit') % (points[t], names[t - 1])
                                              for t in sorted(points) if points[t] > 0)


def build_name(data_name: str, kind: str) -> str:
    """Follow.buildName: the data's "Retribution leveling" reads "Retribution build" (L.followBuildName)."""
    stem = data_name[:-len(' ' + kind)] if data_name.endswith(' ' + kind) else data_name
    return locale('followBuildName') % stem


def split_text(class_slug: str, split: dict) -> str:
    """BuildDialog.split: "Retribution 31 / Holy 20", most points first."""
    names = tabs(class_slug)
    parts = sorted(split.items(), key=lambda kv: (-kv[1], kv[0]))
    return locale('dialogSplitSeparator').join(locale('dialogSplitPart') % (names[t - 1], n) for t, n in parts)


def spec_name(slug: str) -> str:
    return ' '.join(w.capitalize() for w in slug.split('-')[1:])


# --- The tour's chrome, with a lead (Onboarding.lua draw: lead at the top, the body 12 px under) ------
CRESTS = {c: b64((ROOT / f'web/public/icons/hd/crests/{c}.webp').read_bytes(), 'image/webp')
          for c in ('paladin', 'druid', 'hunter')}


def medallion(crest, faction='alliance', size=44):
    inset = 6
    return (f'<div style="position:relative;width:{size}px;height:{size}px;flex:none">'
            f'<div style="position:absolute;inset:{inset}px;border-radius:50%;background:#000 url({crest}) center/100%">'
            f'</div><img src="{RING}" style="position:absolute;inset:0;width:100%;height:100%">'
            f'<div style="position:absolute;inset:0;background:{U[faction]};-webkit-mask:url({NOTCH}) center/100%">'
            f'</div></div>')


def lead(text):
    return f'<div style="font-size:13.5px;line-height:17px;color:{U["ink"]};margin-bottom:12px">{text}</div>'


def tour(page_no, heading, content, lead_key, next_label=None, extra_right='', total=PAGES, crest='paladin',
         last=False):
    """Onboarding.lua: the title band, the lead, the body, the footer. The last page swaps Skip tour for
    L.tourAgain at the footer's left (spec 6.3)."""
    next_label = next_label or locale('tourNext')
    left_extra = (f'<span class="sbtn" style="padding:0 10px">{locale("tourJustTooltips")}</span>' if page_no == 1
                  else '')
    if last:
        left = f'<span style="font-size:12px;color:{U["muted"]}">{locale("tourAgain")}</span>'
    else:
        left = f'<span class="sbtn" style="padding:0 14px">{locale("tourSkip")}</span>{left_extra}'
    back = '' if page_no == 1 else f'<span class="sbtn">{locale("tourBackButton")}</span>'
    footer = (f'<div style="margin:24px 16px 0;padding-bottom:14px;display:flex;align-items:center;'
              f'justify-content:space-between"><span style="display:flex;gap:8px;align-items:center">{left}</span>'
              f'<span style="display:flex;gap:8px">{back}{extra_right}'
              f'<span class="sbtn primary" style="padding:0 20px">{next_label}</span></span></div>')
    band = (f'<div style="height:64px;display:flex;align-items:center;gap:12px;padding:0 42px 0 12px;'
            f'border-bottom:1px solid {U["line"]}">{medallion(CRESTS[crest])}'
            f'<div style="flex:1;min-width:0"><div style="font-family:Title,serif;font-size:19px;color:{U["bone"]};'
            f'text-shadow:1px 1px 0 #000;line-height:22px">{heading}</div>'
            + (f'<div style="font-size:12px;color:{U["muted"]};margin-top:3px">{locale("tourTitle")}</div>'
               if page_no == 1 else '') + '</div>'
            f'<div style="white-space:nowrap;align-self:flex-start;margin-top:12px"><span style="font-family:Site,'
            f'sans-serif;font-weight:600;font-size:17px;color:{U["ink"]}">{page_no}</span>'
            f'<span style="font-size:12px;color:{U["bone_dim"]}"> {locale("tourPageOf") % total}</span></div></div>')
    return (f'<style>@font-face {{ font-family: Title; src: url({TITLE_FONT}); }}</style>'
            f'<div class="tourframe" style="position:relative;width:{OB_W}px;min-height:240px;'
            f'box-sizing:border-box;color:{U["ink"]};font-family:Site,sans-serif;'
            f'background:{U["ground"]};border:1px solid {U["edge"]};box-shadow:0 8px 28px rgba(0,0,0,.7)">'
            f'{band}<div style="position:absolute;top:6px;right:8px;width:20px;height:20px;color:{U["ink"]};'
            f'border:1px solid {U["line"]};box-sizing:border-box;'
            f'font-size:14px;display:flex;align-items:center;justify-content:center">×</div>'
            f'<div style="padding:16px 26px 0 26px">{lead(locale(lead_key))}{content}</div>{footer}</div>')


def board(caption, frame):
    return page(f'<p class="cap">{caption}</p><div style="padding:18px 0 0 18px">{frame}</div>')


# --- Pages 1 to 4 with their leads (the bodies as addon_ux draws them, strings from Locale) ----------
def node_icon(uri, size=36, round_=False):
    return ux.node_icon(uri, size, round_)


def loop_cell(icon, title_key, line_key):
    return ux.loop_cell(icon, locale(title_key), locale(line_key))


def board_onboarding_1(guild=True):
    top = [loop_cell(node_icon(CRESTS['paladin'], round_=True), 'tourCharacter', 'tourCharacterLine'),
           loop_cell(node_icon(ux.ADDON_ICON), 'tourSite', 'tourSiteLine'),
           loop_cell(node_icon(icon_uri('inv_misc_bag_08')), 'tourGame', 'tourGameLine')]
    guild_cell = loop_cell(node_icon(ux.BANNER), 'tourGuild', 'tourGuildLine')
    bottom = [guild_cell if guild else '<div style="width:172px"></div>',
              loop_cell(node_icon(icon_uri('inv_misc_book_09')), 'tourBack', 'tourBackLine'),
              loop_cell(node_icon(icon_uri('ability_warrior_battleshout')), 'tourFights', 'tourFightsLine')]
    row1 = (f'<div style="display:flex;align-items:flex-start;gap:6px">{top[0]}{flow_arrow("right")}{top[1]}'
            f'{flow_arrow("right")}{top[2]}</div>')
    down = f'<div style="display:flex;justify-content:flex-end;padding-right:150px;margin:6px 0">{flow_arrow("down")}</div>'
    first_arrow = flow_arrow('left') if guild else '<div style="width:22px"></div>'
    row2 = (f'<div style="display:flex;align-items:flex-start;gap:6px">{bottom[0]}{first_arrow}{bottom[1]}'
            f'{flow_arrow("left")}{bottom[2]}</div>')
    diagram = f'<div style="width:586px;padding-top:2px">{row1}{down}{row2}</div>'
    total = PAGES if guild else PAGES - 1
    return board(f'<b>Page 1 of {total}</b> · level 13 paladin{"" if guild else ", no guild"}. The lead says what '
                 'the player gets before the loop explains how',
                 tour(1, locale('tourLoopHeading'), diagram, 'tourLoopLead', total=total))


def board_onboarding_1_solo():
    return board_onboarding_1(guild=False)


def icon_row(icon, title_key, line_key, source_key=None, extra=''):
    return ux.icon_row(icon, locale(title_key), locale(line_key), locale(source_key) if source_key else None, extra)


def board_onboarding_2():
    rows = (icon_row(ux.ARROW_ICON, 'tourArrow', 'tourArrowLine', 'tourSiteBis')
            + icon_row(icon_uri('inv_weapon_bow_02'), 'tourOneLine', 'tourOneLineText', 'tourSiteWeights',
                       ux.example_card())
            + icon_row(icon_uri('ability_searingarrow'), 'tourRing', 'tourRingLine', 'tourSiteBuild')
            + icon_row(icon_uri('inv_spear_07'), 'tourRotation', 'tourRotationLine', 'tourSiteRotation'))
    return board('<b>Page 2</b> · the one-line lead, then what the site puts in the game',
                 tour(2, locale('tourSiteHeading'), rows, 'tourSiteLead'))


def board_onboarding_3():
    rows = (icon_row(icon_uri('ability_warrior_battleshout'), 'tourUpload', 'tourUploadLine')
            + icon_row(icon_uri('inv_misc_book_09'), 'tourRated', 'tourRatedLine')
            + ux.icon_row(ux.ADDON_ICON, locale('tourScore'), locale('tourScoreLine') + ' ' + locale('tourScoreHide')))
    return board('<b>Page 3</b> · the lead, then what the game sends back',
                 tour(3, locale('tourBackHeading'), rows, 'tourBackLead'))


def board_onboarding_4():
    """Onboarding.lua drawGuild(), placed as it places them: two columns from the same top, the left
    at 0, the right at 320; "normal" text 13.5 px, the rest "small" 12 px; the chip's clock as
    Compat.clockText prints 20:00 on a 12-hour client (the site's "8 pm")."""
    normal, small = 'font-size:13.5px;line-height:16px', 'font-size:12px;line-height:15px'

    def at(x, y, style, color, body):
        return (f'<div style="position:absolute;left:{x}px;top:{y}px;white-space:nowrap;{style};'
                f'color:{color}">{body}</div>')

    clock = locale('timeTwelveHour') % (8, locale('timePm'))
    job = locale('panelJob') % f'<span style="color:{C["gold"]}">{locale("tourGuildJob")}</span>'
    coming = f'<span style="margin-left:{COMING_GAP}px">{locale("tourComing")}</span>'
    content = (f'<div style="position:relative;height:{62 + 15}px">'
               + at(0, 0, normal, U['ink'], locale('chipTonight') % (clock, locale('tourGuildRaid')))
               + at(0, 20, small, U['muted'], locale('tourGuildChipNote'))
               + at(0, 46, small, U['ink'], job)
               + at(0, 62, small, U['muted'], locale('tourGuildJobNote'))
               + at(320, 0, normal, U['ink'], locale('tourAnswer'))
               + at(320, 20, small, U['ink'], locale('tourAnswerSite'))
               + at(320, 36, small, U['ink'], locale('tourAnswerDiscord'))
               + at(320, 52, small, U['muted'], locale('tourAnswerLater') + coming)
               + '</div>')
    return board('<b>Page 4, in a guild</b> · the lead takes the old intro\'s slot and its claim; the minimap note '
                 'lost "on raid day" because the lead says it',
                 tour(4, locale('tourGuildHeading'), content, 'tourGuildLead'))


# --- The last page: the question and its tiles, or the path followed ---------------------------------
def tile_line(class_slug, spec, kind, taken):
    """Onboarding.tileLine: the playstyle, then "Next: <talent>" (Follow.defaultNext), one separator."""
    nxt = next_talent(class_slug, spec, kind, taken)
    parts = [playstyle(spec)] + ([locale('pathNext') % nxt] if nxt else [])
    return locale('exportTreeSeparator').join(parts)


def tiles(class_slug, points, kind='leveling', taken=None):
    """`points` is {tab: n}; `taken` how many points of each build's order the character already holds
    (0 for a character whose points do not follow any order)."""
    names = tabs(class_slug)
    trees = [names[t - 1] for t in sorted(points) if points[t] > 0]
    specs = [(s, spec_name(s)) for s, _, _ in default_builds(class_slug, kind)]
    n = len(specs)
    width = (588 - 8 * (n - 1)) // n
    heading = ask(trees if len(trees) >= 2 else names)
    split = split_of(class_slug, points)
    note = f'<span style="font-size:12px;color:{U["muted"]};margin-left:8px">{split}</span>' if split else ''
    html = ''.join(f'<div style="width:{width}px"><span class="sbtn" style="width:{width}px">{name}</span>'
                   f'<div style="font-size:12px;color:{U["muted"]};line-height:15px;margin-top:6px">'
                   f'{tile_line(class_slug, s, kind, (taken or {}).get(s, 0))}</div></div>' for s, name in specs)
    return (f'<div style="font-size:13.5px;color:{U["ink"]};margin-bottom:6px">{heading}{note}</div>'
            f'<div style="display:flex;gap:8px;margin-bottom:8px">{html}</div>')


def path_row(name, count, nxt):
    return (f'<div style="display:flex;justify-content:space-between;align-items:flex-start;margin-bottom:24px">'
            f'<div><span style="font-size:13.5px;color:{U["ink"]}">{locale("tourYourPath") % name}</span>'
            f'<span style="font-size:12px;color:{U["muted"]};margin-left:10px">{count}</span>'
            f'<div style="font-size:12px;color:{U["muted"]};margin-top:3px">{nxt}</div></div>'
            f'<span class="sbtn">{locale("panelChange")}</span></div>')


def connect_rows(seen=None, name='Uther'):
    if seen is not None:
        return (f'<div style="font-size:13.5px;color:{C["green"]};margin-bottom:2px">{locale("tourCompanionRan")}</div>'
                f'<div style="font-size:12px;color:{U["muted"]};margin-bottom:10px">{locale("tourSiteSaw") % (name, seen)}'
                f'</div>')
    rows = [(locale('tourSignIn'), locale('tourSignInLine'), url_box('foreversixty.gg')),
            (locale('tourCompanion'), locale('tourCompanionLine'), url_box('foreversixty.gg/setup'))]
    return ''.join(f'<div style="margin-bottom:6px"><div style="color:{U["ink"]};font-size:13.5px">{t}</div>'
                   f'<div style="font-size:12px;color:{U["muted"]};margin:1px 0 4px">{d}</div>{u}</div>'
                   for t, d, u in rows)


def setup_page(top, connect, page_no, total, boxes=True, crest='paladin'):
    links = f'<div style="font-size:12px;color:{U["muted"]}">{locale("tourLinks")}</div>' if boxes else ''
    content = (f'{top}<div style="display:flex;gap:20px"><div style="flex:1">{connect}{links}</div>{qr_block()}</div>')
    return tour(page_no, locale('tourSetupHeading'), content, 'tourSetupLead', next_label=locale('tourDone'),
                total=total, extra_right=f'<span class="sbtn">{locale("tourOpenPanel")}</span>', crest=crest,
                last=True)


def board_onboarding_5_ask():
    return board('<b>Last page, level 13 paladin, 2 Holy and 2 Retribution, no guild</b> · the points name no tree, so '
                 'the page asks, naming the two trees that hold points; every spec is a tile; nothing is '
                 'pre-picked; the split sits beside the question and each tile says the talent it would take next',
                 setup_page(tiles('paladin', {1: 2, 3: 2}), connect_rows(), 4, 4))


def board_onboarding_5_ask_zero():
    return board('<b>Last page, level 10 paladin, no points</b> · too few points to tell, so all three trees are '
                 'named, and no split', setup_page(tiles('paladin', {}), connect_rows(), 4, 4))


def board_onboarding_5_ask_druid():
    return board('<b>Last page, level 22 druid, Feral Combat 8 and Restoration 6</b> · the lead is 2 over 6, so it '
                 'asks; four tiles at 141 px, the playstyle lines wrap to two',
                 setup_page(tiles('druid', {2: 8, 3: 6}), connect_rows(name='Thornhide'), 4, 4, crest='druid'))


def board_onboarding_5_ask_60():
    return board('<b>Last page at 60, 25 Holy and 22 Retribution, nothing followed</b> · the lead is 3, so it asks; '
                 'the tiles load the raid builds, and say their next point',
                 setup_page(tiles('paladin', {1: 25, 3: 22}, kind='raid'), connect_rows(), 5, 5))


def board_onboarding_5_path():
    ret = next(r for r in default_builds('paladin', 'leveling') if r[0] == 'paladin-retribution')
    top = path_row(build_name(ret[1], 'leveling'), locale('panelOnPlan') % (5, 5),
                   locale('tourNextPoint') % (15, next_talent('paladin', 'paladin-retribution', 'leveling', 5)))
    return board('<b>Last page, level 14 paladin, 5 Retribution</b> · the points decide, so the page shows the '
                 'path followed, its count, the next point and Change (which opens the Builds dialog above the '
                 'tour)', setup_page(top, connect_rows(), 4, 4))


def board_onboarding_5_done():
    hol = next(r for r in default_builds('paladin', 'leveling') if r[0] == 'paladin-holy')
    top = path_row(build_name(hol[1], 'leveling'), locale('panelOnPlan') % (15, 15),
                   locale('tourNextPoint') % (25, next_talent('paladin', 'paladin-holy', 'leveling', 15)))
    return board('<b>Last page, a character already set up</b> · level 24, in a guild (page 5 of 5), the companion '
                 'has run; "' + locale('tourAgain') + '" now sits in the footer\'s empty left',
                 setup_page(top, connect_rows(seen='2 hours ago'), 5, 5, boxes=False))


# --- The Builds dialog (BuildDialog.lua on Widgets.panel: the window ground, the title bar, two tabs) ---
DIALOG_W = 420
PAD = 16


def dialog(content, hint_right=None, hint_left=True):
    """Widgets.panel at 420 wide (grows with its content), the title bar with "Forever Sixty" in gold and the
    close X, the Builds and Copy my character tabs (Widgets.tab: 96 x 24, muted label, the active one gold
    over a 2 px gold underline), then the content at 12 px, and the muted hints at the bottom corners."""
    # Widgets.tab is 96 wide; refreshTabs widens the copy tab to its label (BuildDialog.pasteTabWidth).
    tab = ('<span style="display:inline-flex;min-width:96px;height:24px;align-items:center;justify-content:center;'
           'padding:0 8px;white-space:nowrap;font-size:11.5px;color:{c};border-bottom:2px solid {u}">{t}</span>')
    tabs_html = (tab.format(c=ADDON['gold'], u=ADDON['gold'], t=locale('dialogTabPaste'))
                 + '<span style="width:8px"></span>'
                 + tab.format(c=ADDON['muted'], u='transparent', t=locale('dialogCopy')))
    left = (f'<span style="color:{ADDON["muted"]};opacity:.75">{locale("dialogTabHint")}</span>' if hint_left
            else '<span></span>')
    right = f'<span style="color:{ADDON["muted"]};opacity:.75">{hint_right}</span>' if hint_right else ''
    return (f'<div class="dialog" style="width:{DIALOG_W}px;background:rgba(16,10,6,.96);border:1px solid {ADDON["line"]};'
            f'box-shadow:0 8px 28px rgba(0,0,0,.7);font-family:Site,sans-serif;font-size:12px;line-height:15px;'
            f'color:{ADDON["body"]};position:relative">'
            f'<div style="height:28px;display:flex;align-items:center;padding:0 {PAD}px;color:{ADDON["gold"]}">'
            f'{locale("addonName")}</div>'
            f'<div style="position:absolute;top:4px;right:8px;width:20px;height:20px;border:1px solid {ADDON["line"]};'
            f'display:flex;align-items:center;justify-content:center;font-size:14px;color:{ADDON["body"]}">×</div>'
            f'<div style="padding:0 {PAD}px;display:flex">{tabs_html}</div>'
            f'<div style="padding:2px {PAD}px 0">{content}</div>'
            f'<div style="display:flex;justify-content:space-between;padding:10px {PAD}px 2px">{left}{right}</div>'
            f'</div>')


def crest(class_slug, size):
    """Theme.classIcon: the ringed crest, one treatment at every size (the heading's 20, a site row's 32)."""
    return (f'<div style="width:{size}px;height:{size}px;border-radius:50%;background:#000 url({CRESTS[class_slug]}) '
            f'center/100%;box-shadow:0 0 0 1px {ADDON["line"]};flex:none"></div>')


def own_heading(class_slug, text, split=None):
    note = (f'{locale("exportTreeSeparator")}<span style="color:{ADDON["muted"]}">{split}</span>' if split else '')
    return (f'<div style="display:flex;align-items:center;gap:8px;margin:6px 0 8px">{crest(class_slug, 20)}'
            f'<span>{text}{note}</span></div>')


def own_row(name, detail, following=False, cursor=False, failed=False, hover=False, pressed=False, nxt=None):
    ground = ADDON['inset'] if pressed else (ADDON['hover'] if (cursor or hover) else 'transparent')
    edge = (f'<div style="position:absolute;left:-6px;top:0;bottom:0;width:2px;background:{ADDON["body"]}"></div>'
            if cursor else '')
    if following:
        mark = f'<span style="color:{ADDON["gold"]};white-space:nowrap">{locale("dialogFollowingMark")}</span>'
    elif nxt:
        mark = f'<span style="color:{ADDON["muted"]};white-space:nowrap">{locale("pathNext") % nxt}</span>'
    else:
        mark = ''
    detail_html = (f'<span style="color:{ADDON["warning"]}">{locale("dialogOwnFailed")}</span>' if failed
                   else f'<span style="color:{ADDON["muted"]}">{detail}</span>')
    return (f'<div style="position:relative;height:34px;background:{ground};margin-bottom:6px">{edge}'
            f'<div style="display:flex;justify-content:space-between"><span>{name}</span>{mark}</div>'
            f'<div style="margin-top:4px">{detail_html}</div></div>')


def own_rows(class_slug, kind, following_spec=None, cursor_spec=None, taken=None):
    """`taken`: how many points of each build's order the character holds, for the Next talent at the right."""
    return ''.join(own_row(build_name(n, kind), split_text(class_slug, s), following=(spec == following_spec),
                           cursor=(spec == cursor_spec),
                           nxt=next_talent(class_slug, spec, kind, (taken or {}).get(spec, 0)))
                   for spec, n, s in default_builds(class_slug, kind))


def site_row(name, stamp, points, detail, target, cursor=False, class_slug='paladin'):
    ground = ADDON['hover'] if cursor else 'transparent'
    edge = (f'<div style="position:absolute;left:-6px;top:0;bottom:0;width:2px;background:{ADDON["body"]}"></div>'
            if cursor else '')
    return (f'<div style="position:relative;display:flex;gap:8px;margin-bottom:4px;background:{ground}">{edge}'
            f'{crest(class_slug, 32)}<div style="min-width:0">'
            f'<div style="white-space:nowrap">{name} <span style="color:{ADDON["muted"]}">{locale("dialogSeparator")}'
            f'{stamp}{locale("dialogSeparator")}{points}</span></div>'
            f'<div style="color:{ADDON["muted"]};margin-top:2px">{detail}</div>'
            f'<div style="color:{ADDON["muted"]};margin-top:2px">{target}</div></div></div>')


def paste_box():
    return (f'<div style="height:42px;background:{ADDON["inset"]};border:1px solid {ADDON["line"]};padding:5px 8px;'
            f'color:{ADDON["muted"]}">{locale("dialogPaste")}</div>')


def custom_heading():
    return f'<div style="margin:10px 0 6px">{locale("dialogSiteHeading")}</div>'


def companion_line():
    return f'<div style="color:{ADDON["muted"]};margin-top:8px">{locale("dialogPasteCompanion")}</div>'


def next_line(text):
    return f'<div style="color:{ADDON["muted"]};margin:4px 0 8px">{locale("dialogNextTalent") % text}</div>'


def following_line(name, cls, points):
    return (f'<div style="margin-top:4px">{locale("dialogFollowingLine") % (name + locale("dialogSeparator") + cls + locale("dialogSeparator") + (locale("dialogPoints") % points))}</div>')


def dialog_board(caption, content, hint_right=None, hint_left=True, extra=''):
    return page(f'<p class="cap">{caption}</p><div style="padding:18px 0 0 18px;display:flex;gap:28px;'
                f'align-items:flex-start">{dialog(content, hint_right, hint_left)}{extra}</div>')


def board_dialog_ask():
    content = (own_heading('paladin', ask(['Holy', 'Retribution']), split_of('paladin', {1: 2, 3: 2}))
               + own_rows('paladin', 'leveling')
               + custom_heading() + paste_box() + companion_line())
    return dialog_board('<b>Builds dialog, level 13 paladin, 2 Holy 2 Retribution, nothing followed, no site rows</b> '
                        '· the heading is the question with the split, three own rows each saying its next talent, '
                        'none marked; the cursor starts on nothing, so no Enter hint', content)


def board_dialog_own():
    content = (next_line(next_talent('paladin', 'paladin-retribution', 'leveling', 15))
               + own_heading('paladin', locale('dialogOwnLeveling') % 'Paladin')
               + own_rows('paladin', 'leveling', following_spec='paladin-retribution',
                          taken={'paladin-retribution': 15})
               + custom_heading() + paste_box() + companion_line())
    return dialog_board('<b>Builds dialog, level 24 paladin following Retribution, 15 points on plan</b> · the '
                        'points decide, so the heading names the class\'s builds; Following in gold on the row '
                        'followed, the only gold among the rows; the others say their next talent; no "Following:" '
                        'line because the build is one of the rows', content)


def board_dialog_site():
    content = (next_line(next_talent('paladin', 'paladin-retribution', 'leveling', 15))
               + own_heading('paladin', locale('dialogOwnLeveling') % 'Paladin')
               + own_rows('paladin', 'leveling', following_spec='paladin-retribution', taken={'paladin-retribution': 15})
               + custom_heading()
               + site_row('Raid night Ret', locale('dialogFromSiteOn') % '9 Nov', locale('dialogPoints') % 15,
                          'Paladin ' + locale('dialogSeparator') + 'Retribution 15',
                          (locale('dialogForCharacter') % 'Uther') + locale('dialogSeparator') + 'Ashbringer',
                          cursor=True)
               + site_row('Holy for dungeons', locale('dialogFromPush') % 'Mallory', locale('dialogPoints') % 15,
                          'Paladin ' + locale('dialogSeparator') + 'Holy 15', '')
               + f'<div style="color:{ADDON["muted"]};margin:2px 0 6px">{locale("dialogInboxMore") % 2}</div>'
               + paste_box())
    return dialog_board('<b>Builds dialog, the same paladin with four builds waiting from the site</b> · two site rows '
                        'and "and 2 more"; the cursor starts on the first site row (the body-coloured edge, never '
                        'gold), and Enter previews it; the companion line goes while builds wait',
                        content, hint_right=locale('dialogEnterPreview'), hint_left=False)


def board_dialog_60():
    content = (own_heading('paladin', locale('dialogOwnRaid') % 'Paladin')
               + own_rows('paladin', 'raid', following_spec='paladin-retribution')
               + custom_heading() + paste_box() + companion_line())
    return dialog_board('<b>Builds dialog at 60, following Retribution raid</b> · the raid build per spec, with its '
                        'split', content)


def board_dialog_druid():
    content = (next_line(next_talent('druid', 'druid-feral', 'leveling', 15))
               + own_heading('druid', locale('dialogOwnLeveling') % 'Druid')
               + own_rows('druid', 'leveling', following_spec='druid-feral', cursor_spec='druid-feral-bear',
                          taken={'druid-feral': 15})
               + custom_heading()
               + site_row('Bear for Deadmines', locale('dialogFromPush') % 'Mallory', locale('dialogPoints') % 12,
                          'Druid ' + locale('dialogSeparator') + 'Feral Combat 12', '', class_slug='druid')
               + f'<div style="color:{ADDON["muted"]};margin:2px 0 6px">{locale("dialogInboxMore") % 2}</div>'
               + paste_box())
    return dialog_board('<b>Builds dialog, a level 24 druid following Feral</b> · four own rows (Feral and Feral Bear '
                        'share a tree), so the site list cuts to one row and "and 2 more"; the cursor on Feral Bear '
                        'says Enter follows it', content, hint_right=locale('dialogEnterFollow'), hint_left=False)


def board_dialog_custom():
    content = (following_line('My Ret', 'Paladin', 31) + next_line('Vengeance')  # the paste's own next point
               + own_heading('paladin', locale('dialogOwnLeveling') % 'Paladin')
               + own_rows('paladin', 'leveling')
               + custom_heading() + paste_box() + companion_line())
    return dialog_board('<b>Builds dialog, following a pasted build</b> · "Following: My Ret · Paladin · 31 points" '
                        'above the list, no row marked', content)


def board_dialog_states():
    ret = next(r for r in default_builds('paladin', 'leveling') if r[0] == 'paladin-retribution')
    name, detail = build_name(ret[1], 'leveling'), split_text('paladin', ret[2])
    nxt = next_talent('paladin', 'paladin-retribution', 'leveling', 0)
    cases = [('default', {'nxt': nxt}), ('hover', {'hover': True, 'nxt': nxt}),
             ('cursor (keyboard)', {'cursor': True, 'nxt': nxt}), ('pressed', {'pressed': True, 'nxt': nxt}),
             ('Following', {'following': True}), ('Following, hovered', {'following': True, 'hover': True}),
             ('did not load', {'failed': True, 'nxt': nxt})]
    cells = ''.join(f'<div style="width:{DIALOG_W - 2 * PAD}px;font-family:Site,sans-serif;font-size:12px;'
                    f'color:{ADDON["body"]}"><div class="note" style="margin-bottom:4px">{label}</div>'
                    f'<div style="padding:0 0 0 6px">{own_row(name, detail, **kw)}</div></div>' for label, kw in cases)
    return page(f'<p class="cap"><b>Own-build row, every state</b> · gold appears only as Following; the next talent '
                f'muted at the right of the others; the cursor is a 2 px body edge over the hover ground</p><div style="display:flex;flex-direction:column;gap:10px;'
                f'padding:12px 0 0 18px;background:rgba(16,10,6,.96);border:1px solid {ADDON["line"]};'
                f'width:{DIALOG_W}px">{cells}</div>')


# --- The panel's talents block, and the chip's path card ----------------------------------------------
def board_panel_ask():
    specs = default_builds('paladin', 'leveling')
    rows = ''.join(f'<div class="irow"><div class="tx"><div class="n"><span>{spec_name(s)}</span></div>'
                   f'<div class="s" style="white-space:normal">{playstyle(s)}<br>'
                   f'{locale("pathNext") % next_talent("paladin", s, "leveling", 0)}</div></div>'
                   f'<span class="sbtn" style="{ux.SMALL_BTN}">{locale("panelPick")}</span></div>'
                   for s, _, _ in specs)
    path = (f'<div class="sect" style="white-space:normal">{ask(["Holy", "Retribution"])}</div>'
            f'<div class="small" style="margin:-2px 0 8px">{split_of("paladin", {1: 2, 3: 2})}</div>' + rows)
    gear = ('<div class="sect" style="margin-bottom:2px">Upgrades</div>'
            f'<div class="small">{locale("homePicksStart") % 20}</div>')
    plate = (ux.panel('Uther', 'Level 13 Paladin', [path, gear])
             .replace(ux.CREST, CRESTS['paladin']))
    return page(f'<p class="cap"><b>Panel, level 13 paladin, 2 Holy 2 Retribution, nothing followed</b> · the talents '
                f'block is the question with the split and a Pick per spec, each row naming its next talent; no '
                f'rotation while no build is followed, and no "First talent point" foot from level 10</p>'
                f'<div style="padding:16px 0 0 16px">{plate}</div>')


def board_chip_card_ask():
    specs = default_builds('paladin', 'leveling')
    buttons = ''.join(f'<div style="margin-top:8px"><span class="sbtn" style="width:216px;height:20px">{spec_name(s)}'
                      f'</span><div style="font-size:11px;color:{ADDON["muted"]};margin:4px 0 0 8px">{playstyle(s)}<br>'
                      f'{locale("pathNext") % next_talent("paladin", s, "leveling", 0)}</div></div>'
                      for s, _, _ in specs)
    card = (f'<div style="width:240px;background:rgba(16,10,6,.96);border:1px solid {ADDON["line"]};padding:12px;'
            f'font-family:Site,sans-serif;font-size:12.5px;color:{ADDON["body"]}">'
            f'<div style="color:{ADDON["gold"]};line-height:16px">{ask(["Holy", "Retribution"])}</div>{buttons}'
            f'<div style="font-size:11px;color:{ADDON["muted"]};margin-top:12px">{split_of("paladin", {1: 2, 3: 2})}'
            f'</div></div>')
    # The chip's path state draws the class crest (Chip.pathState: input.classIcon).
    chip = re.sub(r'url\([^)]*\)', f'url({CRESTS["paladin"]})', ux.chip(locale('chipPath')), count=1)
    return page(f'<p class="cap"><b>The chip\'s path card, level 13 paladin, 2 Holy 2 Retribution</b> · the line reads '
                f'"Pick a path" at any level while the points name no tree; the card heading is the question, the '
                f'buttons the client recipe with each row\'s next talent, the split at the foot</p>'
                f'<div style="padding:16px 0 0 16px;display:flex;flex-direction:column;gap:8px;align-items:flex-end;'
                f'width:260px">{chip}{card}</div>')


def board_guild_undecided():
    head = ''.join(f'<span style="color:{ADDON["muted"]}">{h}</span>' for h in
                   ('Name', 'Level', 'Spec', 'Gear', 'Build'))
    def row(name, level, spec, gap, build, own=False):
        color = ADDON['gold'] if own else ADDON['body']
        return (f'<span style="color:{color}">{name}</span><span>{level}</span><span>{spec}</span><span>{gap}</span>'
                f'<span style="color:{ADDON["muted"]}">{build}</span>')
    grid = (f'<div style="display:grid;grid-template-columns:120px 56px 120px 90px 1fr;row-gap:6px;font-size:12px;'
            f'font-family:Site,sans-serif;color:{ADDON["body"]}">{head}'
            f'{row("Uther", 13, "", "", "", own=True)}'
            f'{row("Mallory", 60, "Marksmanship", "≥ 3.5 DPS", "Marksmanship build")}'
            f'{row("Thornhide", 24, "", "", "")}'
            f'{row("Ashvale", 60, "", "", "")}</div>')
    plate = (f'<div style="width:560px;background:{ADDON["raised"]};border:1px solid {ADDON["line"]};padding:12px 14px">'
             f'{grid}</div>')
    return page(f'<p class="cap"><b>Guild tab: the own row and two inspected players (24 and 60) with undecided '
                f'points</b> · the spec cell is empty, never a guess or a dash</p>'
                f'<div style="padding:16px 0 0 16px">{plate}</div>')


def board_tour_aside():
    content = (own_heading('paladin', ask(['Holy', 'Retribution']), split_of('paladin', {1: 2, 3: 2}))
               + own_rows('paladin', 'leveling') + custom_heading() + paste_box() + companion_line())
    scene = (f'<div style="position:relative;width:700px;height:420px;background:radial-gradient(ellipse at 40% 30%,'
             f'#2a3a2c,#101812 70%)"><div style="position:absolute;left:120px;top:30px">{dialog(content)}</div></div>')
    return page('<p class="cap"><b>Change pressed on the tour\'s last page</b> · the tour steps aside (hidden, not '
                'ended) and the Builds dialog stands alone, opaque, over the world; closing it brings the tour back on '
                f'the same page</p><div style="padding:18px 0 0 18px">{scene}</div>')


BOARDS = {
    'tour-aside': board_tour_aside,
    'onboarding-1-loop': board_onboarding_1,
    'onboarding-1-loop-solo': board_onboarding_1_solo,
    'onboarding-2-site-in-game': board_onboarding_2,
    'onboarding-3-game-to-site': board_onboarding_3,
    'onboarding-4-guild': board_onboarding_4,
    'onboarding-5-ask': board_onboarding_5_ask,
    'onboarding-5-ask-zero': board_onboarding_5_ask_zero,
    'onboarding-5-ask-druid': board_onboarding_5_ask_druid,
    'onboarding-5-ask-60': board_onboarding_5_ask_60,
    'onboarding-5-path': board_onboarding_5_path,
    'onboarding-5-set-up-done': board_onboarding_5_done,
    'builds-dialog-ask': board_dialog_ask,
    'builds-dialog-own': board_dialog_own,
    'builds-dialog-site': board_dialog_site,
    'builds-dialog-60': board_dialog_60,
    'builds-dialog-druid': board_dialog_druid,
    'builds-dialog-custom': board_dialog_custom,
    'builds-dialog-states': board_dialog_states,
    'panel-ask': board_panel_ask,
    'chip-path-card-ask': board_chip_card_ask,
    'guild-tab-undecided': board_guild_undecided,
}


def main(names):
    from playwright.sync_api import sync_playwright
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1800, 'height': 1000})
        pg = ctx.new_page()
        for name in names:
            pg.set_content(BOARDS[name]())
            pg.evaluate('document.fonts.ready')
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(OUT / f'{name}.png'))
            # The tour frame's and the dialog's heights (CSS px), against the spec's 450 and 380 + chrome.
            heights = pg.evaluate('Array.from(document.querySelectorAll(".tourframe, .dialog")).map(e => e.offsetHeight)')
            print(OUT / f'{name}.png', heights)
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

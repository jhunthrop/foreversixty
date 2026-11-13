"""Acceptance boards for the Builds window (design/specs/2026-11-13-addon-builds-window.md section 8).

    python3 design/mocks/addon_builds_window.py                 # every board
    python3 design/mocks/addon_builds_window.py window-hunter-42

Writes design/reviews/2026-11-13-addon-builds-window/<board>.png at device scale 2, the window at its real
740 px width (Theme.SIZES.buildsWindowWidth): the header band, the build list at 388 on the left (the Builds
pane's rows, drawn by addon_builds.py's own functions), the right column at 288 with Up next, From the site
and To the site, and the footer.

Sources, so nothing is invented: every string from addon/ForeverSixty/Locale.lua (the spec's new keys are
read from there once they land, from NEW until then); builds, splits and whys from addon-builds.json; the
next points from the followed build's order over talents/<class>.json (name, icon, max rank); the level 42
hunter's figures from the owner's capture of 2026-11-12.
"""
import collections
import re
import os
import sys

import addon_builds as ab
import addon_ux as ux
from addon_builds import L, esc, lua_format
from addon_ux import ADDON, ROOT

OUT = ROOT / 'design/reviews/2026-11-13-addon-builds-window'
W, PAD, LEFT, GUTTER, RIGHT = 740, 16, 388, 32, 288
HEADER, FOOTER, UP_NEXT_ROWS = 56, 28, 6
CLIENT_GOLD = ab.CLIENT_GOLD

# The spec's new strings (section 6), until Locale.lua carries them.
NEW = {
    'windowUpNext': 'Up next', 'windowUpNextNow': 'Now', 'windowUpNextLevel': 'Level %d', 'windowUpNextLevels': 'Levels %d to %d',
    'windowUpNextRank': '%d/%d', 'windowUpNextNone': 'Follow a build to see its next points.',
    'windowUpNextDone': 'Every point is placed.', 'windowUpNextMore': '%d more points after these',
    'windowUpNextMoreOne': '1 more point after these', 'windowToSpend': '%d to spend', 'windowOffPlan': 'Off plan',
    'windowFromSiteWhat': 'On the planner, press Copy for the game, then paste the code here.',
    'windowPaste': 'Paste a build code (Ctrl+V)', 'windowRead': 'Read: %s. Check it on the left.',
    'windowToSite': 'To the site',
    'windowToSiteWhat': 'Talents, gear and professions in one code for the planner’s Import from addon box.',
    'windowClickToSelect': 'Click the box, then press Ctrl+C.', 'windowEscClose': 'Esc to close',
    'windowEscBack': 'Esc for your builds', 'windowNoBuilds': 'No builds for this class yet.',
}


def T(key):
    try:
        return L(key)
    except KeyError:
        return NEW[key]


CSS = ab.CSS + f'''
.win {{ position: relative; width: {W}px; background: {ADDON['raised']}; border: 1px solid {ADDON['line']};
  box-shadow: 0 8px 28px rgba(0,0,0,.7); font-family: Site, sans-serif; font-size: 12px; line-height: 15px;
  color: {ADDON['body']}; }}
.win .hd {{ height: {HEADER}px; display: flex; align-items: center; gap: 10px; padding: 0 {PAD}px;
  border-bottom: 1px solid {ADDON['line']}; }}
.win .hd .tt2 {{ font-family: Title, serif; font-size: 16px; line-height: 20px; color: {ADDON['cream']}; }}
.win .x {{ position: absolute; top: 18px; right: 12px; width: 20px; height: 20px; border: 1px solid {ADDON['line']};
  display: flex; align-items: center; justify-content: center; font-size: 14px; }}
.win .bd {{ display: flex; padding: 12px {PAD}px; }}
.win .lc {{ width: {LEFT}px; flex: none; }}
.win .vr {{ position: absolute; left: {PAD + LEFT + GUTTER // 2}px; top: {HEADER}px; bottom: {FOOTER}px;
  width: 1px; background: {ADDON['line']}; }}
.win .lc {{ margin-right: {GUTTER}px; }}
.win .brow::before {{ right: -{ab.G['groundLead']}px; }}
.win .rule {{ width: 1px; background: {ADDON['line']}; margin: -12px {GUTTER // 2 - 1}px -12px {GUTTER // 2}px; flex: none;
  align-self: stretch; }}
.win .rc {{ width: {RIGHT}px; flex: none; }}
.win .ft {{ height: {FOOTER}px; border-top: 1px solid {ADDON['line']}; display: flex; align-items: center;
  justify-content: space-between; padding: 0 {PAD}px; color: {ADDON['muted']}; }}
.win .ft span {{ opacity: .75; }}
.sec {{ margin-bottom: 20px; }}
.sec:last-child {{ margin-bottom: 0; }}
.sh {{ height: 18px; display: flex; align-items: center; justify-content: space-between; gap: 8px;
  font-family: Title, serif; font-size: 14px; color: {ADDON['cream']}; }}
.sh .r {{ font-family: Site, sans-serif; font-size: 12px; }}
.sh .r {{ font-weight: 400; color: {ADDON['muted']}; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.what {{ color: {ADDON['muted']}; margin: 2px 0 6px; }}
.un {{ height: 18px; display: flex; align-items: center; gap: 6px; }}
.un img {{ width: 16px; height: 16px; box-shadow: 0 0 0 1px #000; }}
.un .lv {{ width: 92px; white-space: nowrap; flex: none; color: {ADDON['muted']}; }}
.un .lv.now {{ color: {CLIENT_GOLD}; }}
.un .nm {{ flex: 1; min-width: 0; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }}
.un .rk {{ color: {ADDON['muted']}; }}
.rc .box {{ height: 42px; }}
.rc .box.warn {{ border-color: {ADDON['warning']}; }}
.rc .box.gold {{ border-color: {ADDON['gold']}; }}
.rc .box.one-line {{ height: 27px; }}
.rc .box.code {{ word-break: break-all; overflow: hidden; }}
.rc .box.code span {{ display: -webkit-box; -webkit-line-clamp: 2; -webkit-box-orient: vertical; overflow: hidden; }}
.rc .box.sel {{ border-color: {ADDON['gold']}; }}
.rc .box.sel span {{ background: {ADDON['hover']}; color: {ADDON['body']}; }}
.st {{ margin-top: 6px; }}
.ok {{ color: {ADDON['success']}; }}
.warn-t {{ color: {ADDON['warning']}; }}
.q {{ margin-bottom: 8px; }}
.pv .ph {{ display: flex; gap: 8px; margin-bottom: 6px; }}
.pv .gh {{ color: {ADDON['muted']}; margin: 6px 0 2px; }}
.pv .gr {{ display: flex; align-items: center; gap: 6px; height: 16px; margin-bottom: 2px; }}
.pv .gr .sl {{ width: 72px; color: {ADDON['muted']}; }}
.pv .gr img {{ width: 14px; height: 14px; box-shadow: 0 0 0 1px #000; }}
.gold-btn {{ display: inline-flex; align-items: center; height: 22px; padding: 0 12px; font-weight: 600;
  color: #1a1206; background: {ADDON['gold']}; border: 1px solid #f3d27a; }}
'''

# --- Up next: the followed build's next points (BuildsWindow.upNext) -----------------------------


def cells(cls):
    out = {}
    for tab, tree in enumerate(ab.trees(cls), start=1):
        for n in tree['talents']:
            out[(tab, n['tier'] + 1, n['column'] + 1)] = n
    return out


def up_next(spec, kind, start_name, unspent, level, start_index=None):
    """BuildsWindow.upNext: the order from the first `start_name` point on, a run of points on one talent as
    one row. A point waiting reads Now; the rest arrive one a level from level + 1. The row's rank is the
    talent's rank once the run is taken."""
    cls = spec.split('-')[0]
    order = [tuple(p) for p in ab.BUILDS[spec][kind]['order']]
    table = cells(cls)
    names = [table[p]['name'] for p in order]
    start = start_index if start_index is not None else (names.index(start_name) if start_name else len(order))
    groups = []
    for i in range(start, len(order)):
        k = i - start
        now = k < unspent
        at = level + k - unspent + 1
        last = groups[-1] if groups else None
        if last and last['cell'] == order[i] and last['now'] == now:
            last['to'], last['rank'] = at, last['rank'] + 1
            continue
        if len(groups) == UP_NEXT_ROWS:
            break
        rank = sum(1 for p in order[:i + 1] if p == order[i])
        groups.append({'cell': order[i], 'now': now, 'from': at, 'to': at, 'rank': rank})
    total = len([1 for i in range(start, len(order)) if level + (i - start) - unspent + 1 <= 60 or i - start < unspent])
    rows = []
    for g in groups:
        cell = table[g['cell']]
        if g['now']:
            when = T('windowUpNextNow')
        elif g['to'] > g['from']:
            when = lua_format(T('windowUpNextLevels'), g['from'], g['to'])
        else:
            when = lua_format(T('windowUpNextLevel'), g['from'])
        rows.append((cell['icon'], when, g['now'], cell['name'], g['rank'], cell['max_rank']))
    shown = sum(g['to'] - g['from'] + 1 if not g['now'] else 1 for g in groups)
    return rows, max(0, total - shown)


def up_next_html(build_name, rows_more, empty=None):
    rows, more = rows_more if rows_more else ([], 0)
    head = f'<div class="sh"><span>{T("windowUpNext")}</span><span class="r">{esc(build_name or "")}</span></div>'
    if empty:
        return f'<div class="sec">{head}<div class="what">{empty}</div></div>'
    body = ''
    for icon, when, now, name, rank, top in rows:
        rk = f' <span class="rk">{lua_format(T("windowUpNextRank"), rank, top)}</span>' if top > 1 else ''
        body += (f'<div class="un">{ux.img(icon, 16)}<span class="lv{" now" if now else ""}">{when}</span>'
                 f'<span class="nm">{esc(name)}{rk}</span></div>')
    if more:
        text = T('windowUpNextMoreOne') if more == 1 else lua_format(T('windowUpNextMore'), more)
        body += f'<div class="m" style="margin:2px 0 0 22px">{text}</div>'
    return f'<div class="sec">{head}<div style="height:4px"></div>{body}</div>'


# --- From the site and To the site --------------------------------------------------------------


def from_site(site='', box_cls='', box_text=None, after=''):
    text = box_text if box_text is not None else esc(T('windowPaste'))
    return (f'<div class="sec"><div class="sh"><span>{L("dialogSiteHeading")}</span></div>'
            f'<div class="what">{esc(T("windowFromSiteWhat"))}</div>{site}'
            f'<div class="box {box_cls}">{text}</div>{after}</div>')


def export(cls='hunter', level=42):
    """A made-up character's export in the FS1 shape (never a real player's name or guild)."""
    return (f'FS1:70938:{cls}:{level}:Thornmaw:Ashen Vigil:Forever:1,2:a1b5c2d7e3f9g1h2i5j0k3l2m8n1p4q7r2s6t9u3v'
            '8w5x1y4z7:head=10164,neck=12047,shoulder=14775,back=9947,chest=15055,wrist=14762,hands=15060')


def to_site(state='rest', facts=(), cls='hunter', level=42):
    box_cls = {'rest': '', 'selected': 'sel', 'copied': ''}[state]
    status = {'rest': f'<div class="st m">{esc(T("windowClickToSelect"))}</div>',
              'selected': f'<div class="st">{esc(L("dialogSelectedLead"))}</div>',
              'copied': f'<div class="st ok">{esc(lua_format(L("dialogYouCopied"), "22:30"))}</div>'}[state]
    facts_html = ''.join(f'<div class="m">{esc(f)}</div>' for f in facts)
    return (f'<div class="sec"><div class="sh"><span>{T("windowToSite")}</span></div>'
            f'<div class="what">{esc(T("windowToSiteWhat"))}</div>'
            f'<div class="box code {box_cls}"><span>{export(cls, level)}</span></div>{status}{facts_html}</div>')


def facts_hunter(level=42):
    return [L('dialogSeparator').join([lua_format(L('dialogSummaryItems'), 14),
                                       lua_format(L('dialogSummaryProfessions'), 2), L('dialogSummaryShared')]),
            lua_format(L('dialogSeen'), '1 h 35 min ago')]


# --- The window ---------------------------------------------------------------------------------


def header(cls, title, note_html):
    return (f'<div class="hd">{ab.crest(cls, 32)}<div style="flex:1;min-width:0"><div class="tt2">{esc(title)}</div>'
            f'<div class="m">{note_html}</div></div></div><div class="x">×</div>')


def window(cls, title, note_html, left, right, esc_hint, enter_hint=''):
    return (f'<div class="win">{header(cls, title, note_html)}<div class="bd"><div class="lc">{left}</div>'
            f'<div class="rc">{right}</div></div><div class="vr"></div>'
            f'<div class="ft"><span>{esc_hint}</span><span>{enter_hint}</span></div></div>')


def blocks_html(blocks):
    return f'<div style="height:{ab.G["blockGap"]}px"></div>'.join(blocks)


def board(caption, html):
    """The board is the window alone; `caption` (what the board shows) is kept as the board's title only."""
    return ux.page(f'<style>{CSS}</style><div title="{esc(caption)}">{html}</div>')


def name_of(spec_name, kind):
    """Follow.buildName for a default build: "Beast Mastery level 60 build", "Marksmanship raid build"."""
    word = re.search(rf'followKindWord = {{ leveling = "([^"]*)", raid = "([^"]*)" }}', ab.LOCALE)
    return lua_format(L('followKindBuildName'), spec_name, word.group(1 if kind == 'leveling' else 2))


def note(level, matched=None, allowed=None, off=0, spend=0):
    """BuildsPane.note: the level, then parts that add up to its points: on plan, off plan, to spend."""
    parts = [lua_format(L('dialogLevelPoints'), level)]
    if allowed:
        parts.append(lua_format(L('windowOnPlan'), matched))
    if off:
        parts.append(f'<span style="color:{ADDON["warning"]}">{lua_format(L("windowOffPlanCount"), off)}</span>')
    out = L('dialogSeparator').join(parts)
    if spend:
        out += L('dialogSeparator') + f'<span style="color:{CLIENT_GOLD}">{lua_format(T("windowToSpend"), spend)}</span>'
    return out


def off_plan(lines):
    """BuildDialog's drawOffPlan: "Off plan" in the warning colour, a muted line a talent."""
    body = ''.join(f'<div class="one m">{esc(line)}</div>' for line in lines)
    return f'<div class="warn-t">{T("windowOffPlan")}</div>{body}<div style="height:8px"></div>'


# The owner's capture named the count (5 off plan), not the talents: two of the hunter's own tier-1 talents
# stand in for them, in PanelTab.offPlanTalents' own line (L.panelOffPlanTalent).
HUNTER_OFF_PLAN = [lua_format(L('windowTalentHave'), *args)
                   for args in (('Hawk Eye', 3, 3, 'Marksmanship'), ('Deflection', 2, 5, 'Survival'))]


def board_hunter_42():
    nxt = ab.next_line('hunter', 'Intimidation', 1)
    blocks = ab.class_blocks('hunter', 42, 'hunter-beast-mastery', 'leveling', nxt,
                             states={'hunter-marksmanship': {1: 'cursor'}})
    rows = up_next('hunter-beast-mastery', 'leveling', 'Intimidation', 1, 42)
    right = (up_next_html(name_of('Beast Mastery', 'leveling'), rows) + from_site(box_cls='gold')
             + to_site('rest', facts_hunter()))
    return board("The Builds window · level 42 Hunter following Beast Mastery's level 60 build, one point to spend, "
                 "27 of 33 on plan and 5 off (the owner's case, 2026-11-12); the cursor on Marksmanship's raid build "
                 "after the arrow keys, the paste box holding the keyboard; 740 wide",
                 window('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), note(42, 27, 33, 5, 1),
                        off_plan(HUNTER_OFF_PLAN) + blocks_html(blocks), right, T('windowEscClose'),
                        L('dialogEnterFollow')))


def site_druid():
    return ab.site_rows([
        (f'Moonkin for Molten Core{L("dialogSeparator")}51 points', 'Balance 31 / Restoration 20',
         f'{lua_format(L("dialogForCharacter"), "Brindlefen")}{L("dialogSeparator")}12 Nov', 'druid'),
        (f'Bear for Onyxia{L("dialogSeparator")}51 points', 'Feral Combat 34 / Restoration 17',
         f'{lua_format(L("dialogFromPush"), "Varnok")}{L("dialogSeparator")}12 Nov', 'druid')], 2)


def druid_up_next():
    """Level 24 with the order's first 15 points taken: the next point is the 16th, at level 25."""
    return up_next('druid-feral', 'leveling', None, 0, 24, start_index=15)


def druid_next_line():
    order = ab.BUILDS['druid-feral']['leveling']['order']
    return ab.next_line('druid', cells('druid')[tuple(order[15])]['name'], 0, 25)


def board_druid_24():
    blocks = ab.class_blocks('druid', 24, 'druid-feral', 'leveling', druid_next_line(),
                             specs=['druid-balance', 'druid-feral', 'druid-feral-bear', 'druid-restoration'])
    facts = [L('dialogSeparator').join([lua_format(L('dialogSummaryItems'), 11), L('dialogSummarySharedRealm')]),
             L('dialogNotSent')]
    right = (up_next_html(name_of('Feral', 'leveling'), druid_up_next()) + from_site(site_druid())
             + to_site('selected', facts, 'druid', 24))
    return board('The Builds window · level 24 Druid following Feral, every point on plan, four builds waiting from '
                 'the site (two rows and "and 2 more"), the copy box selected after Tab',
                 window('druid', lua_format(L('dialogOwnBuilds'), 'Druid'), note(24, 15, 15),
                        blocks_html(blocks), right, T('windowEscClose')))


def board_hunter_60_copied():
    blocks = ab.class_blocks('hunter', 60, 'hunter-marksmanship', 'raid', None)
    sim = L('dialogSimDone')
    right = (up_next_html(name_of('Marksmanship', 'raid'), None, T('windowUpNextDone')) + from_site()
             + to_site('copied', facts_hunter(60)[:1] + [lua_format(L('dialogSeen'), '2 min ago'), sim], 'hunter', 60))
    return board('The Builds window · level 60 Hunter following the Marksmanship raid build, every point placed; '
                 'the export copied with Ctrl+C, the site\'s ack and the sim line under it',
                 window('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), note(60, 51, 51),
                        blocks_html(blocks), right, T('windowEscClose')))


def preview_changes():
    """BuildsWindow.talentDiff of the pasted Balance build against the level 24 druid's 15 Feral points, in tree
    order, CHANGE_ROWS lines then "and N more" (BuildsWindow.talentLines with L.windowTalentBuild)."""
    feral = ab.BUILDS['druid-feral']['leveling']['order'][:15]
    moonkin = ab.BUILDS['druid-balance']['raid']['order']
    table = cells('druid')
    have = collections.Counter(tuple(p) for p in feral)
    want = collections.Counter(tuple(p) for p in moonkin)
    trees = [t['name'] for t in ab.trees('druid')]
    keys = sorted(set(have) | set(want))
    entries = [k for k in keys if have[k] != want[k]]
    lines = []
    for k in entries[:6]:
        cell = table[k]
        if want[k] == 0:
            lines.append(lua_format(L('windowTalentHave'), cell['name'], have[k], cell['max_rank'], trees[k[0] - 1]))
        else:
            lines.append(lua_format(L('windowTalentBuild'), cell['name'], have[k], cell['max_rank'], trees[k[0] - 1],
                                    want[k]))
    if len(entries) > 6:
        lines.append(lua_format(L('dialogInboxMore'), len(entries) - 6))
    return lines


def preview_left():
    gear = [('Head', 'Cenarion Helm', 4), ('Chest', 'Robe of Volatile Power', 4),
            ('Neck', 'Choker of Enlightenment', 4), ('Hands', 'Cenarion Gloves', 4)]
    rows = ''.join(f'<div class="gr"><span class="sl">{slot}</span>'
                   f'<span style="color:{ux.QUALITY[q]}">{n}</span></div>' for slot, n, q in gear)
    replaces = lua_format(L('dialogReplacesDiffers'), name_of('Feral', 'leveling'), '15 points', 1, 'Improved Wrath')
    changes = ''.join(f'<div class="one m">{esc(line)}</div>' for line in preview_changes())
    return (f'<div class="pv"><div class="ph">{ab.crest("druid", 32)}<div><div>Moonkin for Molten Core'
            f'{L("dialogSeparator")}51 points</div><div class="m">Balance 31 / Restoration 20</div></div></div>'
            f'<div class="gh">{L("dialogGearHeading")}{L("dialogSeparator")}4 of 17 slots</div>{rows}'
            f'<div style="margin-top:8px">{T("windowChanges")}</div>{changes}'
            f'<div class="warn-t" style="margin-top:10px">{esc(replaces)}</div>'
            f'<div class="m" style="margin:2px 0 6px">{lua_format(L("dialogNextWould"), "Improved Wrath")}</div>'
            f'<span class="gold-btn">{L("dialogUse")}</span></div>')


def board_paste_preview():
    after = f'<div class="st ok">{esc(lua_format(T("windowRead"), "Moonkin for Molten Core"))}</div>'
    box = '<span class="one" style="display:block">FSB1:70938:druid:1a21a21a21b31b31b3…</span>'
    right = (up_next_html(name_of('Feral', 'leveling'), druid_up_next())
             + from_site(box_cls='gold one-line', box_text=box, after=after)
             + to_site('rest', facts_hunter(24)[1:], 'druid', 24))
    return board('The Builds window · a pasted Balance code previewing on the left; the box one line, "Read" under '
                 'it; Enter uses it',
                 window('druid', lua_format(L('dialogOwnBuilds'), 'Druid'), note(24, 15, 15),
                        preview_left(), right, T('windowEscBack'), L('dialogEnterHint')))


def board_bad_code():
    nxt = ab.next_line('hunter', 'Intimidation', 1)
    blocks = ab.class_blocks('hunter', 42, 'hunter-beast-mastery', 'leveling', nxt)
    rows = up_next('hunter-beast-mastery', 'leveling', 'Intimidation', 1, 42)
    lead, body = L('codecWrongPrefix').split('. ', 1)
    tail = re.search(r'codecWrongPrefix = .*\n\t\t\.\. "([^"]*)"', ab.LOCALE).group(1)
    after = f'<div class="st warn-t">{esc(lead)}.</div><div class="m">{esc(body + tail)}</div>'
    box = '<span class="one" style="display:block">https://www.wowhead.com/classic/talent-calc/hunter/5300052-0505</span>'
    right = (up_next_html(name_of('Beast Mastery', 'leveling'), rows)
             + from_site(box_cls='warn one-line', box_text=box, after=after) + to_site('rest', facts_hunter()))
    return board('The Builds window · a pasted link that is not a build code: the warning edge and the refusal, the '
                 'list kept; Esc clears the box',
                 window('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), note(42, 27, 33, 5, 1),
                        off_plan(HUNTER_OFF_PLAN) + blocks_html(blocks), right, T('windowEscClose')))


BOARDS = {
    'window-hunter-42': board_hunter_42,
    'window-druid-24-site': board_druid_24,
    'window-paste-preview': board_paste_preview,
    'window-bad-code': board_bad_code,
    'window-copied': board_hunter_60_copied,
}


def main(names, out=None):
    from playwright.sync_api import sync_playwright
    target = out or OUT
    target.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        ctx = browser.new_context(device_scale_factor=2, viewport={'width': 1400, 'height': 1000})
        pg = ctx.new_page()
        for name in names:
            pg.set_content(BOARDS[name]())
            pg.evaluate('document.fonts.ready')
            pg.wait_for_timeout(150)
            pg.locator('#board').screenshot(path=str(target / f'{name}.png'))
            print(target / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

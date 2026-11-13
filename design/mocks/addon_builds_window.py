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
    'windowUpNext': 'Up next', 'windowUpNextNow': 'Now', 'windowUpNextLevel': 'Level %d', 'windowUpNextLevels': 'Level %d to %d',
    'windowUpNextRank': '%d/%d', 'windowUpNextNone': 'Follow a build to see its next points.',
    'windowUpNextDone': 'Every point of this build is placed.',
    'windowFromSiteWhat': 'On the planner, press Copy for the game, then paste the code here.',
    'windowPaste': 'Paste a build code (Ctrl+V)', 'windowRead': 'Read: %s. Check it on the left.',
    'windowToSite': 'To the site',
    'windowToSiteWhat': "Your talents, gear and professions as one code, for the planner's Import from addon box.",
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
.win .rule {{ width: 1px; background: {ADDON['line']}; margin: -12px {GUTTER // 2 - 1}px -12px {GUTTER // 2}px; flex: none; }}
.win .rc {{ width: {RIGHT}px; flex: none; }}
.win .ft {{ height: {FOOTER}px; border-top: 1px solid {ADDON['line']}; display: flex; align-items: center;
  justify-content: space-between; padding: 0 {PAD}px; color: {ADDON['muted']}; }}
.win .ft span {{ opacity: .75; }}
.sec {{ margin-bottom: 20px; }}
.sec:last-child {{ margin-bottom: 0; }}
.sh {{ height: 18px; display: flex; align-items: center; justify-content: space-between; gap: 8px;
  font-weight: 600; color: {ADDON['strong']}; }}
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


def up_next(spec, kind, start_name, unspent, level):
    """BuildsWindow.upNext: the order from the first `start_name` point on, a run of points on one talent as
    one row. A point waiting reads Now; the rest arrive one a level from level + 1. The row's rank is the
    talent's rank once the run is taken."""
    cls = spec.split('-')[0]
    order = [tuple(p) for p in ab.BUILDS[spec][kind]['order']]
    table = cells(cls)
    names = [table[p]['name'] for p in order]
    start = names.index(start_name) if start_name else len(order)
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
    return rows


def up_next_html(build_name, rows, empty=None):
    head = f'<div class="sh"><span>{T("windowUpNext")}</span><span class="r">{esc(build_name or "")}</span></div>'
    if empty:
        return f'<div class="sec">{head}<div class="what">{empty}</div></div>'
    body = ''
    for icon, when, now, name, rank, top in rows:
        rk = f' <span class="rk">{lua_format(T("windowUpNextRank"), rank, top)}</span>' if top > 1 else ''
        body += (f'<div class="un">{ux.img(icon, 16)}<span class="lv{" now" if now else ""}">{when}</span>'
                 f'<span class="nm">{esc(name)}{rk}</span></div>')
    return f'<div class="sec">{head}<div style="height:4px"></div>{body}</div>'


# --- From the site and To the site --------------------------------------------------------------


def from_site(site='', box_cls='', box_text=None, after=''):
    text = box_text if box_text is not None else esc(T('windowPaste'))
    return (f'<div class="sec"><div class="sh"><span>{L("dialogSiteHeading")}</span></div>'
            f'<div class="what">{esc(T("windowFromSiteWhat"))}</div>{site}'
            f'<div class="box {box_cls}">{text}</div>{after}</div>')


EXPORT = ('FS1:70938:hunter:42:Thornmaw:Ashen Vigil:Forever:1,2:a1b5c2d7e3f9g1h2i5j0k3l2m8n1p4q7r2s6t9u3v'
          '8w5x1y4z7:head=10164,neck=12047,shoulder=14775,back=9947,chest=15055,wrist=14762,hands=15060')


def to_site(state='rest', facts=(), link=None):
    if state == 'none':
        body = f'<div class="what" style="margin:0">{esc(L("exportNoSaveYet")) if False else "Log out once and the addon saves your character here."}</div>'
        return f'<div class="sec"><div class="sh"><span>{T("windowToSite")}</span></div>{body}</div>'
    cls = {'rest': '', 'selected': 'sel', 'copied': ''}[state]
    status = {'rest': f'<div class="st m">{esc(T("windowClickToSelect"))}</div>',
              'selected': f'<div class="st">{esc(L("dialogSelectedLead"))}</div>',
              'copied': f'<div class="st ok">{esc(lua_format(L("dialogYouCopied"), "22:30"))}</div>'}[state]
    facts_html = ''.join(f'<div class="m">{esc(f)}</div>' for f in facts)
    return (f'<div class="sec"><div class="sh"><span>{T("windowToSite")}</span></div>'
            f'<div class="what">{esc(T("windowToSiteWhat"))}</div>'
            f'<div class="box code {cls}"><span>{EXPORT}</span></div>{status}{facts_html}</div>')


def facts_hunter(level=42):
    return [L('dialogSeparator').join([lua_format(L('dialogSummaryWho'), level, 'Hunter'),
                                       lua_format(L('dialogSummaryItems'), 14),
                                       lua_format(L('dialogSummaryProfessions'), 2), L('dialogSummaryShared')]),
            lua_format(L('dialogSeen'), '1 h 35 min ago')]


# --- The window ---------------------------------------------------------------------------------


def header(cls, title, note_html):
    return (f'<div class="hd">{ab.crest(cls, 32)}<div style="flex:1;min-width:0"><div class="tt2">{esc(title)}</div>'
            f'<div class="m">{note_html}</div></div></div><div class="x">×</div>')


def window(cls, title, note_html, left, right, esc_hint, enter_hint=''):
    return (f'<div class="win">{header(cls, title, note_html)}<div class="bd"><div class="lc">{left}</div>'
            f'<div class="rule"></div><div class="rc">{right}</div></div>'
            f'<div class="ft"><span>{esc_hint}</span><span>{enter_hint}</span></div></div>')


def blocks_html(blocks):
    return f'<div style="height:{ab.G["blockGap"]}px"></div>'.join(blocks)


def board(caption, html):
    """The board is the window alone; `caption` (what the board shows) is kept as the board's title only."""
    return ux.page(f'<style>{CSS}</style><div title="{esc(caption)}">{html}</div>')


def board_hunter_42():
    nxt = ab.next_line('hunter', 'Intimidation', 1)
    blocks = ab.class_blocks('hunter', 42, 'hunter-beast-mastery', 'leveling', nxt,
                             states={'hunter-marksmanship': {1: 'cursor'}})
    rows = up_next('hunter-beast-mastery', 'leveling', 'Intimidation', 1, 42)
    right = (up_next_html('Beast Mastery leveling build', rows) + from_site()
             + to_site('rest', facts_hunter()))
    return board("The Builds window · level 42 Hunter following Beast Mastery's leveling build, one point to spend, "
                 "27 of 33 on plan and 5 off (the owner's case, 2026-11-12); the cursor on Marksmanship's raid build; "
                 "740 wide",
                 window('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), ab.note(42, 27, 33, 5),
                        blocks_html(blocks), right, T('windowEscClose'), L('dialogEnterFollow')))


def site_druid():
    return ab.site_rows([
        (f'Moonkin for Molten Core{L("dialogSeparator")}51 points', 'Balance 31 / Restoration 20',
         f'{lua_format(L("dialogForCharacter"), "Leafcall")}{L("dialogSeparator")}12 Nov', 'druid'),
        (f'Bear for Onyxia{L("dialogSeparator")}51 points', 'Feral Combat 34 / Restoration 17',
         f'{lua_format(L("dialogFromPush"), "Boss")}{L("dialogSeparator")}12 Nov', 'druid')], 2)


def board_druid_24():
    nxt = ab.next_line('druid', 'Feral Instinct', 0, 25)
    blocks = ab.class_blocks('druid', 24, 'druid-feral', 'leveling', nxt,
                             specs=['druid-balance', 'druid-feral', 'druid-feral-bear', 'druid-restoration'])
    rows = up_next('druid-feral', 'leveling', 'Feral Instinct', 0, 24)
    facts = [L('dialogSeparator').join([lua_format(L('dialogSummaryWho'), 24, 'Druid'),
                                        lua_format(L('dialogSummaryItems'), 11), L('dialogSummarySharedRealm')]),
             L('dialogNotSent')]
    right = up_next_html('Feral leveling build', rows) + from_site(site_druid()) + to_site('selected', facts)
    return board('The Builds window · level 24 Druid following Feral, four builds waiting from the site (two rows and '
                 '"and 2 more"), the copy box selected after Tab',
                 window('druid', lua_format(L('dialogOwnBuilds'), 'Druid'), ab.note(24, 15, 15),
                        blocks_html(blocks), right, T('windowEscClose'), L('dialogEnterFollow')))


def board_hunter_60_copied():
    blocks = ab.class_blocks('hunter', 60, 'hunter-marksmanship', 'raid', None)
    sim = L('dialogSimDone')
    right = (up_next_html('Marksmanship raid build', [], T('windowUpNextDone')) + from_site()
             + to_site('copied', facts_hunter(60)[:1] + [lua_format(L('dialogSeen'), '2 min ago'), sim]))
    return board('The Builds window · level 60 Hunter following the Marksmanship raid build, every point placed; '
                 'the export copied with Ctrl+C, the site\'s ack and the sim line under it',
                 window('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), ab.note(60, 51, 51),
                        blocks_html(blocks), right, T('windowEscClose')))


def preview_left():
    gear = [('Head', 'inv_helmet_09', 'Cenarion Helm', 4), ('Chest', 'inv_chest_cloth_06', 'Robe of Volatile Power', 4),
            ('Neck', 'inv_jewelry_necklace_01', 'Choker of Enlightenment', 4),
            ('Hands', 'inv_gauntlets_14', 'Cenarion Gloves', 4)]
    rows = ''.join(f'<div class="gr"><span class="sl">{s}</span>{ux.img(i, 14) if False else ""}'
                   f'<span style="color:{ux.QUALITY[q]}">{n}</span></div>' for s, i, n, q in gear)
    replaces = lua_format(L('dialogReplacesDiffers'), 'Feral leveling build', '15 points', 1, 'Improved Wrath')
    return (f'<div class="pv"><div class="ph">{ab.crest("druid", 32)}<div><div>Moonkin for Molten Core'
            f'{L("dialogSeparator")}51 points</div><div class="m">Balance 31 / Restoration 20</div></div></div>'
            f'<div class="gh">{L("dialogGearHeading")}{L("dialogSeparator")}4 of 17 slots</div>{rows}'
            f'<div class="warn-t" style="margin-top:10px">{esc(replaces)}</div>'
            f'<div class="m" style="margin:2px 0 6px">{lua_format(L("dialogNextWould"), "Improved Wrath")}</div>'
            f'<span class="gold-btn">{L("dialogUse")}</span></div>')


def board_paste_preview():
    rows = up_next('druid-feral', 'leveling', 'Feral Instinct', 0, 24)
    after = f'<div class="st ok">{esc(lua_format(T("windowRead"), "Moonkin for Molten Core"))}</div>'
    box = '<span class="one" style="display:block">FSB1:70938:druid:1a21a21a21b31b31b3…</span>'
    right = (up_next_html('Feral leveling build', rows) + from_site(box_cls='gold', box_text=box, after=after)
             + to_site('rest', facts_hunter(24)[1:]))
    return board('The Builds window · a pasted Balance code previewing on the left; "Read" under the box; Enter uses it',
                 window('druid', lua_format(L('dialogOwnBuilds'), 'Druid'), ab.note(24, 15, 15),
                        preview_left(), right, T('windowEscBack'), L('dialogEnterHint')))


def board_bad_code():
    nxt = ab.next_line('hunter', 'Intimidation', 1)
    blocks = ab.class_blocks('hunter', 42, 'hunter-beast-mastery', 'leveling', nxt)
    rows = up_next('hunter-beast-mastery', 'leveling', 'Intimidation', 1, 42)
    reason = L('codecWrongPrefix') if False else 'That is not a Forever Sixty build code.'
    after = (f'<div class="st warn-t">{esc(reason)}</div>'
             f'<div class="m">{esc(lua_format(L("dialogStillFollowing"), "Beast Mastery leveling build"))}</div>')
    box = '<span>https://www.wowhead.com/classic/talent-calc/hunter/5300052-0505</span>'
    right = (up_next_html('Beast Mastery leveling build', rows) + from_site(box_cls='warn', box_text=box, after=after)
             + to_site('rest', facts_hunter()))
    return board('The Builds window · a pasted link that is not a build code: the warning edge and the refusal, the '
                 'list kept; Esc clears the box',
                 window('hunter', lua_format(L('dialogOwnBuilds'), 'Hunter'), ab.note(42, 27, 33, 5),
                        blocks_html(blocks), right, T('windowEscClose')))


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

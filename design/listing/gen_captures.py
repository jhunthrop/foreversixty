"""Store gallery captures from the addon mocks, for gen_gallery.mjs.

    python3 design/listing/gen_captures.py                 # every capture
    python3 design/listing/gen_captures.py window-me chip  # some

Writes design/listing/captures/<name>.png, which gen_gallery.mjs prefers over any JPEG. Each capture is a
board of the addon hub's lane generators (design/mocks/addon_hub_*.py, addon_builds_window.py), the renders
the 2.0.48 frames were reviewed and shipped against: every word and layout rule there comes from the addon's
own Locale.lua, Data.lua and models. This script renders those boards again without their review notes and
cuts out the frame; it never draws a frame of its own.

Each lane drew Velisandre from its own fixture, so side by side the plates disagreed (a gain, a countdown,
a sign-up, a raid date). The store fixture below makes one evening of one character out of them:

- Thu 17 Dec 2026 is raid day (Onyxia's Lair at 20:00, 3 h 12 min away) and she signed up In. Her last raid
  is Molten Core on Tue 15 Dec, after the raids opened on 9 Dec, with the Rotation lane's boss results (the
  site's loss phrases; it measures no per-spell lateness).
- Her gear is the Gear lane's model (Raid-ready at 60): Beaststalker's Gloves in her bags beat her Trueaim
  Gauntlets by 2.6 DPS (the marks lane's figure); her sim's biggest gains are the Gear tab's top three; her
  gap on the Officer roster is the Gear tab's total.
- Builds shows her at 60 following the Beast Mastery raid build, every point placed, as Me says; with no
  builds from the site there, the line's hover leaves that item out.
- The guild's wanted roles are one table (addon_hub.WANTED) on Guild and Officer.

The device scale is chosen per capture so the frame fills the plate's stage (gen_gallery's 1440x600 at
device scale 2) without the plate ever enlarging it, capped per capture so a small tooltip is not drawn
huge. Every character, guild and raid group here is made up (Velisandre in Thornwake Company and her raid
team, Kestrin and Orlaith on the player tooltips).
"""
import os
import sys
from pathlib import Path

HERE = Path(__file__).resolve().parent
sys.path.insert(0, str(HERE.parent / 'mocks'))

import addon_builds as ab  # noqa: E402
import addon_builds_window as builds  # noqa: E402
import addon_hub as hub  # noqa: E402
import addon_hub_l3 as gear  # noqa: E402
import addon_hub_l5 as guild  # noqa: E402
import addon_hub_l6 as chip  # noqa: E402
import addon_hub_l7 as officer  # noqa: E402
import addon_hub_l10 as marks  # noqa: E402
import addon_hub_me as me  # noqa: E402
import addon_hub_rotation as rotation  # noqa: E402

OUT = HERE / 'captures'
# gen_gallery.mjs: the stage is 1440x600 CSS px at device scale 2, inside a 1 px border.
STAGE_DEVICE = ((1440 - 2) * 2, (600 - 2) * 2)
MIN_SCALE = 1.5
WINDOW_MAX_SCALE, MARK_MAX_SCALE = 3, 3.5
# The boards' review notes and captions: hidden so only the addon's own frames are left.
HIDE_NOTES = '.cap, .cap2, .note { display: none !important; }'

# --- The store fixture: one evening of one character ---------------------------------------------------------

# Board dates moved a week on, so the last raid falls after the raids opened (9 Dec). Longest first.
DATES = [('Tue 15 Dec 19:30', 'Tue 22 Dec 19:30'), ('Thu 10 Dec', 'Thu 17 Dec'), ('Sun 13 Dec', 'Sun 20 Dec')]
LAST_RAID_ISO, RECORD_ISO, SIM_ISO = '2026-12-15', '2026-12-16', '2026-12-14T20:55:00Z'
BAG_ITEM, BAG_GAIN = 16676, 2.6  # Beaststalker's Gloves over the worn Trueaim Gauntlets (l10-gear-marks)
GEAR = gear.velisandre('raid')
GEAR_CELLS = [c for c in GEAR.doll().values() if c.get('pick') and not c.get('holds') and c.get('gain', 0) > 0]
GEAR_TOTAL = round(sum(c['gain'] for c in GEAR_CELLS), 1)
GEAR_TOP = sorted(GEAR_CELLS, key=lambda c: -c['gain'])[:3]
OFFICER_GAP_SHOWN = '40.4'  # Velisandre's gap on the Officer lane's roster (addon_hub.ROSTER)


def shift_dates(html):
    for old, new in DATES:
        html = html.replace(old, new)
    return html


def me_last_raid():
    """The Me lane's raid with the Rotation lane's boss results (the same night on both tabs)."""
    results = rotation.MOLTEN_CORE['bosses']
    return dict(me.LAST_RAID, date=LAST_RAID_ISO, bosses=[
        dict(boss, biggest_loss=result['biggest_loss'], loss_seconds=result['loss_seconds'], deaths=result['deaths'])
        for boss, result in zip(me.LAST_RAID['bosses'], results)])


def me_raid_day():
    quality = 'ff' + hub.QUALITY[hub.ITEMS[BAG_ITEM]['quality']].lstrip('#')
    name = hub.ITEMS[BAG_ITEM]['name']
    leads = dict(me.LEADS)
    leads['gear'] = {'tab': 'gear', 'label': 'Gear', 'waiting': [{'key': 'bag', 'line': f'Equip {name} +{BAG_GAIN}'}],
                     'lead': {'itemId': BAG_ITEM, 'line': f'|c{quality}{name}|r |c{me.SUCCESS}+{BAG_GAIN} DPS|r',
                              'sub': 'Hands · in your bags'}}
    leads['rotation'] = dict(me.LEADS['rotation'], lead=dict(me.LEADS['rotation']['lead'], sub='Tue 15 Dec · parse 58'))
    sim = dict(me.SIM, generated_at=SIM_ISO, slots=[
        {'slot': c['slot'], 'item_id': c['pick']['item_id'], 'delta': c['gain']} for c in GEAR_TOP])
    me.STATES = {'raid-day': me.state(
        record=dict(me.RECORD, as_of=RECORD_ISO), pack=dict(me.BASE['pack'], generated=RECORD_ISO),
        lastRaid=me_last_raid(), sim=sim, leads=[leads[k] for k in ('builds', 'gear', 'rotation', 'guild')])}
    key, caption, dots, tabs = me.BOARDS['me-raid-day']
    return me.board(key, caption, me.models()[key], ('Gear',), tabs)


def builds_60():
    """Builds at 60 following the Beast Mastery raid build, every point placed, inside the hub's strip."""
    blocks = ab.class_blocks('hunter', 60, 'hunter-beast-mastery', 'raid', None)
    right = (builds.up_next_html(builds.name_of('Beast Mastery', 'raid'), None, builds.T('windowUpNextDone'))
             + builds.from_site() + builds.to_site('rest', builds.facts_hunter(60), 'hunter', 60))
    html = builds.board('Builds at 60', builds.window(
        'hunter', builds.NAMES['hunter'], builds.note(60, 51, 51), builds.blocks_html(blocks), right,
        builds.T('windowEscClose'), strip=('Builds', ('Gear',))))
    return html.replace('Ashen Vigil', hub.GUILD)


def gear_raid_ready():
    return gear.board(GEAR, 'Gear, level 60, Raid-ready', 'Velisandre', 'Level 60 Hunter', 60, 'All sources')


def guild_signed_in():
    html = guild.win(guild.tonight('in') + guild.prep() + guild.signups(),
                     guild.next_raids(('In as DPS', 'ok')) + guild.guild_block() + guild.sent(), dots=())
    return shift_dates(hub.board2('Guild, signed up', html).replace('</style>', guild.TITLE_CSS + '</style>', 1))


def officer_raid_day():
    caption, build = officer.STATES['l7-raid-day']
    html = hub.board2(caption, build()).replace('</style>', officer.CSS + '</style>', 1)
    if html.count(OFFICER_GAP_SHOWN) != 1:
        raise RuntimeError(f'the Officer roster no longer shows one {OFFICER_GAP_SHOWN} gap')
    return shift_dates(html.replace(OFFICER_GAP_SHOWN, f'{GEAR_TOTAL:.1f}'))


def chip_hover():
    name = hub.ITEMS[BAG_ITEM]['name']
    html = chip.board_hover().replace('Starts in 1 h 12 min', 'Starts in 3 h 12 min')
    html = html.replace('Gloves of Holy Might', name).replace('+11.2 DPS', f'+{BAG_GAIN} DPS')
    # Builds shows no builds from the site, so the hover drops that line and the count drops to +1.
    site_line = ('<div class="l " style="color:#fff"><span>2 builds from the site</span>'
                 '<span style="color:#b8b8b8">Builds</span></div>')
    count = 'class="m cnt" style="margin-left:6px">+2<'
    first = html.index('class="scene"')
    scene = html[first:html.index('class="scene"', first + 1)]
    if scene.count(site_line) != 1 or scene.count(count) != 1:
        raise RuntimeError('the chip board no longer draws the builds line and the +2 count this script expects')
    return shift_dates(html[:first] + scene.replace(site_line, '').replace(count, count.replace('+2', '+1'))
                       + html[first + len(scene):])


def rotation_60():
    return rotation.board_60()


# name -> (board html, the frame to cut out, extra CSS for this board, the largest device scale)
CAPTURES = {
    # The line and its hover only: the board's stand-in minimap and its ground are left out.
    'chip': (chip_hover, '#board .row > div:first-child .scene > *',
             '#board .row > div:not(:first-child) { display: none !important; } '
             '.scene .mm { display: none !important; } '
             '.scene { background: none !important; box-shadow: none !important; height: 0 !important; }', 4.5),
    'window-me': (me_raid_day, '.win', '', WINDOW_MAX_SCALE),
    'window-builds': (builds_60, '.win', '', WINDOW_MAX_SCALE),
    'window-gear': (gear_raid_ready, '.win', '', WINDOW_MAX_SCALE),
    'window-rotation': (rotation_60, '.win', '', WINDOW_MAX_SCALE),
    'window-guild': (guild_signed_in, '.win', '', WINDOW_MAX_SCALE),
    'window-officer': (officer_raid_day, '.win', '', WINDOW_MAX_SCALE),
    # The Hands slot with its tooltip and the bags with the arrow; the bag item's tooltip repeats the gain.
    'bag-arrow': (marks.board_gear, '#board',
                  '#board .row > .col:nth-child(2) .tt { display: none !important; }', MARK_MAX_SCALE),
    # The upgrade reward and two rolls: the "nothing helps" quest frame and the repeated Striker's Mark go.
    'quest-reward-marks': (marks.board_rewards, '#board',
                           '#board .row > .col:first-child > .qf:last-child, '
                           '#board .row > .col:nth-child(2) > .rf:nth-of-type(4) { display: none !important; }',
                           MARK_MAX_SCALE),
    # The stale record and the Shift view; the fresh one is the Shift view's own first lines.
    'player-tooltip': (marks.board_player, '#board',
                       '#board > .row > .col:first-child { display: none !important; }', MARK_MAX_SCALE),
}

# The box of every element the selector matches, with everything drawn inside them (tooltips that hang past
# a frame's edge included).
BOX_JS = '''(selector) => {
  const roots = [...document.querySelectorAll(selector)];
  if (!roots.length) throw new Error(`no ${selector}`);
  let [l, t, r, b] = [Infinity, Infinity, -Infinity, -Infinity];
  for (const el of roots.flatMap((root) => [root, ...root.querySelectorAll('*')])) {
    const box = el.getBoundingClientRect();
    if (box.width === 0 || box.height === 0 || getComputedStyle(el).display === 'none') continue;
    if (el.closest('[style*="display: none"]')) continue;
    [l, t, r, b] = [Math.min(l, box.left), Math.min(t, box.top), Math.max(r, box.right), Math.max(b, box.bottom)];
  }
  return { x: Math.floor(l + scrollX), y: Math.floor(t + scrollY),
           width: Math.ceil(r - l), height: Math.ceil(b - t) };
}'''


def scale_for(box, most):
    fit = min(STAGE_DEVICE[0] / box['width'], STAGE_DEVICE[1] / box['height'])
    return max(MIN_SCALE, min(most, int(fit * 100) / 100))


def render(browser, name):
    build, selector, extra, most = CAPTURES[name]
    html = build().replace('</style>', HIDE_NOTES + extra + '</style>', 1)
    scale = MIN_SCALE
    for _ in range(2):  # measure at the minimum scale, then shoot at the scale that fills the stage
        ctx = browser.new_context(device_scale_factor=scale, viewport={'width': 1800, 'height': 1200})
        page = ctx.new_page()
        page.set_content(html)
        page.evaluate('document.fonts.ready')
        if name == 'chip':
            page.evaluate(chip.FIT_JS)
        page.wait_for_timeout(150)
        box = page.evaluate(BOX_JS, selector)
        if scale != scale_for(box, most):
            scale = scale_for(box, most)
            ctx.close()
            continue
        file = OUT / f'{name}.png'
        page.screenshot(path=str(file), clip=box, full_page=True)
        ctx.close()
        print(f'{file.relative_to(HERE.parent.parent)} | {box["width"]}x{box["height"]} css at {scale}x')
        return
    raise RuntimeError(f'{name}: scale did not settle')


def main(names):
    from playwright.sync_api import sync_playwright
    unknown = [n for n in names if n not in CAPTURES]
    if unknown:
        raise SystemExit(f'unknown capture {unknown}; one of {list(CAPTURES)}')
    OUT.mkdir(parents=True, exist_ok=True)
    with sync_playwright() as pw:
        exe = os.environ.get('PLAYWRIGHT_CHROMIUM')
        browser = pw.chromium.launch(executable_path=exe) if exe else pw.chromium.launch()
        for name in names:
            render(browser, name)
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(CAPTURES))

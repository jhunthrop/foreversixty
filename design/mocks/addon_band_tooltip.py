"""Mock boards for the worn item's band lines (owner's report 2026-11-13, a level 42 hunter).

    python3 design/mocks/addon_band_tooltip.py                 # every board
    python3 design/mocks/addon_band_tooltip.py tooltip-waist-42-after

Writes design/reviews/2026-11-13-addon-tooltip-tour/<board>.png at device scale 2, through addon_ux.py's
client tooltip look. Nothing is typed here: every addon line is read from addon/ForeverSixty/Locale.lua
(addon_tour.locale), the waist picks and their names from addon/ForeverSixty/Data.lua (the same packed
picks Gear.pick decodes), the spec weights and the band's DPS per point from Data.lua too, and the items'
own lines from data/builds/1.60.1.70790/items/hunter.json. The owner's belt is the capture's (Brigade
Girdle of the Bandit: 4 Agility, 9 Stamina, Equip +11 Attack Power, Requires Level 38). The figures are
worked exactly as the addon works them (Gear.currentWeights' shipped weights times the band's DPS per
point, one decimal), so the board shows what the client will.
"""
import os
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
import addon_ux as ux  # noqa: E402
from addon_tour import locale  # noqa: E402
from addon_ux import C, COPPER, ITEMS, QUALITY, SILVER, arrow, iname, l, mini, page, tooltip  # noqa: E402

ROOT = Path(__file__).resolve().parents[2]
OUT = Path(os.environ['MOCK_OUT']) if os.environ.get('MOCK_OUT') else ROOT / 'design/reviews/2026-11-13-addon-tooltip-tour'
DATA = (ROOT / 'addon/ForeverSixty/Data.lua').read_text()
SPEC = 'hunter-beast-mastery'
SPEC_WORD = 'Beast Mastery'
SLOT = 'waist'
BANDIT_ID, BANDIT_NAME = 9931, 'Brigade Girdle of the Bandit'
BANDIT_STATS = {'agility': 4, 'stamina': 9, 'attack_power': 11}


# --- Data.lua, read, never typed ----------------------------------------------------------------------
def spec_block(table: str) -> str:
    """The text of Data.lua's `<table>` entry for the spec."""
    start = DATA.index('\n\t%s = {' % table)
    begin = DATA.index('["%s"] = {' % SPEC, start)
    end = DATA.index('["hunter-marksmanship"] = {', begin)
    return DATA[begin:end]


def picks() -> dict:
    """Band -> the alliance waist pick's item id, from the packed `waist = "<id>|..."` strings."""
    out = {}
    block = spec_block('bis')
    for m in re.finditer(r'\[(\d+)\] = \{', block):
        band = int(m.group(1))
        alliance = block.index('alliance = {', m.end())
        waist = re.compile(r'%s = "(\d+)\|' % SLOT).search(block, alliance)
        out[band] = int(waist.group(1))
    return out


PICKS = picks()
NAMES = {int(i): n for i, n in re.findall(r'\[(\d+)\] = "((?:[^"\\]|\\.)*)"', DATA[DATA.index('\n\tbisNames = {'):])}
WEIGHTS = {k: float(v) for k, v in re.findall(r'\["(\w+)"\] = ([\d.]+)',
                                              re.search(r'\["%s"\] = \{([^\n]*)\}' % SPEC, DATA[DATA.index('\n\tweights = {'):]).group(1))}
DPS_PER_POINT = {int(b): float(d) for b, d in re.findall(r'\[(\d+)\] = \{ reference = "\w+", dps = ([\d.]+)',
                                                         spec_block('dpsPerPoint'))}


def band_for(level: int) -> int:
    return max(b for b in PICKS if b <= level)


def band_after(band: int, item_id: int):
    """ItemTooltip.bandAfter: the first later band whose waist pick is not `item_id`, else None."""
    for candidate in sorted(b for b in PICKS if b > band):
        if PICKS[candidate] != item_id:
            return candidate
    return None


def band_word(band: int, after) -> str:
    return locale('tooltipBandCap') % band if after is None else locale('tooltipBandSpan') % (band, after - 1)


def best_text(band: int, after) -> str:
    return locale('tooltipBestAtCap') % band if after is None else locale('tooltipBestFor') % band_word(band, after)


def hold_text(band: int, after, name_html: str, gain) -> str:
    key = 'tooltipBeatsPick' if gain is not None and gain < 0 else 'tooltipMatchesPick'
    return locale(key) % (band_word(band, after), name_html)


def at_text(band: int) -> str:
    return locale('tooltipAtBandName') % (band, iname(PICKS[band]))


def score(stats: dict) -> float:
    return sum(v * WEIGHTS.get(k, 0) for k, v in stats.items())


def gain(level: int, pick_id: int, worn_stats: dict) -> float:
    """The pick's gain over what is worn, in DPS, one decimal (Progress.pickGain's rounding)."""
    return round((score(ITEMS[pick_id]['stats']) - score(worn_stats)) * DPS_PER_POINT[band_for(level)], 1)


def figure(value: float) -> str:
    sign = locale('tooltipMinus') if value < 0 else '+'
    return locale('tooltipFigure') % (sign, abs(value), locale('tooltipUnitDps'))


def attribution() -> str:
    return locale('tooltipAttribution') % (locale('tooltipAttrEstimate') % SPEC_WORD)


def crafted_source() -> str:
    return locale('tooltipSourcePair') % ('Crafted', 'Leatherworking')


# --- The worn item's lines, as the addon prints them ---------------------------------------------------
def worn_lines(level: int, worn_id: int, worn_stats: dict, shift=False):
    """ItemTooltip.section's verdict and support line for a worn waist item, Shift's block when asked."""
    band = band_for(level)
    pick_id = PICKS[band]
    g, grey, gold, green = C['grey'], C['grey'], C['gold'], C['green']
    if worn_id == pick_id:
        after = band_after(band, worn_id)
        lines = [l(best_text(band, after), color=gold)]
        detail = []
        if after is not None:
            lines.append(l(at_text(after), color=grey))
            detail.append(l(crafted_source(), color=grey))
        return lines, detail
    delta = gain(level, pick_id, worn_stats)
    after = band_after(band, pick_id)
    if delta >= 1.0:
        lines = [l(f'{locale("tooltipNextLabel")} {iname(pick_id)} <span style="color:{green}">{figure(delta)}</span>'),
                 l(crafted_source(), color=grey)]
        return lines, []
    lines = [l(hold_text(band, after, iname(pick_id), delta))]
    # The pick's figure and source; the next band's pick is the same profession, said once.
    detail = [l(figure(delta), color=g), l(crafted_source(), color=g)]
    if after is not None:
        lines.append(l(at_text(after), color=grey))
    return lines, detail


def bandit_client_lines():
    """The capture's own lines for the owner's belt (not_showing_40_upgrades.png)."""
    return [
        f'<div class="t" style="color:{QUALITY[2]}">{BANDIT_NAME}</div>', l('Soulbound'), l('Waist', 'Mail'),
        l('136 Armor'), l('+4 Agility'), l('+9 Stamina'), l('Equip: +11 Attack Power.', color=C['green']),
        l('Durability 28 / 35'), l('Requires Level 38'), l(f'Sell Price: 41{SILVER}56{COPPER}'),
    ]


def item_client_lines(item_id: int):
    """An item's lines from items/hunter.json: name, binding, slot and type, armor, stats, level."""
    it = ITEMS[item_id]
    bound = {'pickup': 'Binds when picked up', 'equip': 'Binds when equipped'}.get(it.get('bonding'), 'Soulbound')
    armor = (it.get('armor_type') or 'Leather').capitalize()
    out = [f'<div class="t" style="color:{QUALITY[it["quality"]]}">{it["name"]}</div>', l(bound), l('Waist', armor),
           l(f'{it["armor"]} Armor')]
    for stat in ('agility', 'stamina', 'intellect', 'spirit'):
        if it['stats'].get(stat):
            out.append(l(f'+{it["stats"][stat]} {stat.capitalize()}'))
    out.append(l(f'Requires Level {it["required_level"]}'))
    return out


def board(caption: str, cards: str) -> str:
    return page(f'<p class="cap">{caption}</p><div class="row" style="padding-top:18px">{cards}</div>')


# --- Boards --------------------------------------------------------------------------------------------
def board_before():
    lines = bandit_client_lines() + [l(locale('tooltipNoUpgradeUntil') % 50), l(at_text(50), color=C['grey'])]
    return board('<b>Before</b> · the owner\'s capture: level 42 Beast Mastery, Brigade Girdle of the Bandit worn. '
                 '"No upgrade until level 50" reads as "you have the 40 band\'s best"; the 40 to 49 pick is never named.',
                 tooltip(lines, width=330))


def board_after():
    plain, detail = worn_lines(42, BANDIT_ID, BANDIT_STATS)
    shift = tooltip(bandit_client_lines() + plain + detail + [l(attribution(), color=C['grey'])], width=350)
    return board('<b>After</b> · the same hover. The band\'s pick is named and yours is said to beat it; the next '
                 'band\'s pick follows. Right: with Shift, the pick\'s own figure and where it comes from.',
                 f'<div class="col"><div class="note"><b>Default</b></div>{tooltip(bandit_client_lines() + plain, width=330)}</div>'
                 f'<div class="col"><div class="note"><b>Shift held</b></div>{shift}</div>')


def board_pick():
    pick = PICKS[40]
    plain, _ = worn_lines(42, pick, ITEMS[pick]['stats'])
    return board('<b>Worn and it is the pick</b> · the same hunter once the belt is crafted: gold says which '
                 'band it is best for, then what changes at 50.', tooltip(item_client_lines(pick) + plain, width=330))


def board_edges():
    # The highest belt a level 42 could wear that the band's pick beats by the floor or more (the data's
    # own candidates; the gain is worked as the addon works it).
    beaten = max((i for i in ITEMS.values() if i['slot'] == SLOT and i['required_level'] <= 42
                  and i['id'] != PICKS[40] and gain(42, PICKS[40], i['stats']) >= 1.0), key=lambda i: i['item_level'])
    cases = [
        (25, PICKS[20], ITEMS[PICKS[20]]['stats'], '<b>Level 25</b> · the 20 band\'s pick worn'),
        (35, PICKS[30], ITEMS[PICKS[30]]['stats'], '<b>Level 35</b> · a pick that holds for two bands'),
        (42, BANDIT_ID, BANDIT_STATS, '<b>Level 42</b> · the owner\'s belt beats the pick'),
        (42, beaten['id'], beaten['stats'], '<b>Level 42</b> · a belt the pick beats'),
        (55, PICKS[50], ITEMS[PICKS[50]]['stats'], '<b>Level 55</b> · the 50 band\'s pick worn'),
        (60, PICKS[60], ITEMS[PICKS[60]]['stats'], '<b>Level 60</b> · the top band, no span'),
    ]
    cards = []
    for level, worn_id, stats, caption in cases:
        plain, _ = worn_lines(level, worn_id, stats)
        title = iname(worn_id, BANDIT_NAME if worn_id == BANDIT_ID else None)
        cards.append(mini(title, plain, caption, width=345))
    rows = ''.join(f'<div class="row" style="margin-bottom:22px">{"".join(cards[i:i + 3])}</div>'
                   for i in range(0, len(cards), 3))
    return page(f'<p class="cap"><b>Every band edge</b> · the waist for a Beast Mastery hunter, alliance, from the '
                f'shipped picks: gold is yours for the band named, white names the pick yours holds against, '
                f'the second line is always the band that changes it.</p>{rows}')


BOARDS = {
    'tooltip-waist-42-before': board_before,
    'tooltip-waist-42-after': board_after,
    'tooltip-waist-pick': board_pick,
    'tooltip-band-edges': board_edges,
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
            print(OUT / f'{name}.png')
        browser.close()


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

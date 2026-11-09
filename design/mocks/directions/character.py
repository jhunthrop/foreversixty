"""Mock batch 2: the character page and /characters in the directions system (plan 2026-11-09-redesign).

Data, never invented: web/tests/e2e/fixtures/character and fixtures/home, the files the character parity spec routes,
read through design/mocks/ws_character.py's own model (Char: the band, the items, the build's one gain rule). The page
shown is the visitor's view of Elyra Duskvale, a level 60 Holy Priest. /characters reads search-thalrin.json. Drawn by
prims.py and the shared paperdoll.
"""
import sys
from html import escape
from pathlib import Path

sys.path.append(str(Path(__file__).resolve().parent.parent))  # design/mocks: the parity mocks' data readers

import ws_character as WC  # noqa: E402
import data as D  # noqa: E402
import paperdoll as PD  # noqa: E402
from common import write  # noqa: E402
from prims import PRIMS_CSS, facts, field, item_row, meter, panel, row_list, settings_bar, table  # noqa: E402
from style import footer, header, medallion, page, titlebar  # noqa: E402

VIEWER = ('Thandrel', 'warrior', 60)
SCORE_MAX = {'gear': 500, 'experience': 700, 'parses': 600, 'attendance': 200, 'discipline': 100}
UPGRADES_SHOWN = 4

CHAR_CSS = PRIMS_CSS + PD.CSS + '''
.ch{padding-top:6px}.dollwrap{display:flex;justify-content:center}.dollwrap>.pdl{width:100%}.dollwrap>.pd{width:100%;max-width:640px}
.pd .ctr .ups{width:100%;max-width:330px;background:rgba(6,9,20,.84);border-radius:4px;padding:10px 12px}
.pd .ctr .ups h3{font:600 13px var(--body);color:#fff;margin-bottom:4px}.pd .ctr .ups .ir{padding:5px 0}.pd .ctr .ups .ir img{width:30px;height:30px}
.scp .mt{min-width:0}.scp td:nth-child(2){width:46%}.gain{color:var(--gain);font-weight:700}.tier{font:400 22px var(--title)}.pct{font-weight:700}
@media(max-width:''' + str(PD.MIN_FRAME_PX + 28) + '''px){.pd{display:none}.pdl{display:block}}
'''


def gain_text(cell: dict) -> str:
    g = cell.get('gain')
    if not g or not g.get('unit'):
        return ''
    return f'<span class="gain">{g["value"]} {g["unit"]}</span>'


def gear_panel(ch: WC.Char) -> str:
    cells = {g['slot']: WC.cell(ch, g, True) for g in ch.page['gear']}
    upgrades = sorted((c for c in cells.values() if gain_text(c)), key=lambda c: -float(c['gain']['value'].lstrip('+').replace(',', '')))
    ups = ''.join(item_row(D.icon_uri(c['icon']), c['name'], c['quality'], escape(c['label']), gain_text(c)) for c in upgrades[:UPGRADES_SHOWN])
    centre = f'<div class="ups"><h3>Upgrades</h3>{ups}</div>' if ups else ''
    spec = f'Holy {ch.c["class"]}'
    total = f'{ch.total():,.1f}'
    # the bar names the spec and level; the band already carries the upgrades' total, so the bar carries no figure
    doll = PD.frame(cells, '', False, spec, f'Level {ch.c["level"]}', '', centre, unit=ch.unit)
    listed = PD.list_form(cells, '', False, spec, '', f'<div style="margin-top:12px">{ups}</div>', unit=ch.unit)
    return panel('Gear', f'<div class="dollwrap" style="--frame:url({PD.frame_uri()})">{doll}{listed}</div>', action='<a href="#">Best in slot &rsaquo;</a>', framed=True)


def score_panel(ch: WC.Char) -> str:
    p = ch.page
    tier = WC.LABELS['tiers'][p['tier']]
    rating = WC.J(WC.FIX / 'rating-60.json')
    label = {x['key']: x['label'] for x in WC.LABELS['components']}
    head = facts([('Score', f'{p["score"]:,}'), ('Tier', f'<span style="color:{tier["color"]}">{escape(tier["label"])}</span>')])
    # each part's bar in the quality colour of how full it is; the figure alone (the fixture's attendance, 330, runs past
    # the contract's 200 cap: a data question for the score lane, so no "of N" is drawn against it)
    # one scale for every part: each track is as long as the part's maximum, so a part's share of the score reads at a
    # glance and no "of N" is drawn (the fixture's attendance, 330, runs past the contract's 200 cap: a data question)
    top = max(SCORE_MAX.values())
    rows = [[escape(label[k]), f'<span style="display:block;width:{mx / top * 100:.0f}%">{meter(p["components"][k] / mx, quality_colour(p["components"][k] / mx))}</span>',
             str(p['components'][k])] for k, mx in SCORE_MAX.items()]
    return panel('Forever Score', head + '<div class="scp">' + table([('', False), ('', False), ('', True)], rows).replace('<thead>', '<thead hidden>') + '</div>')


def quality_colour(frac: float) -> str:
    """How full a score part is, on the item quality ladder players read: common, uncommon, rare, epic."""
    for floor, q in ((.9, 'epic'), (.75, 'rare'), (.5, 'uncommon')):
        if frac >= floor:
            return f'var(--q-{q})'
    return 'var(--q-common)'


def fights_panel(ch: WC.Char) -> str:
    rows = [('', escape(f['encounter']), f'{WC.short_date(f["fought_at"])} &middot; {escape(f["spec"])}',
             f'<span class="pct" style="color:{WC.parse_color(f["percentile"])}">{round(f["percentile"])}</span> &nbsp;{f["value"]:,} {ch.unit}')
            for f in ch.page['best'][:WC.FIGHTS_SHOWN]]
    return panel('Best per encounter', row_list(rows), action=f'<a href="#">All ranked fights ({len(ch.page["history"])}) &rsaquo;</a>')


def builds_panel(ch: WC.Char) -> str:
    rows = [('', escape(b['talent_split']), f'{escape(b["spec"])} &middot; first seen {WC.short_date(b["first_seen"], True)}', '') for b in ch.page['builds_seen'][:3]]
    return panel('Builds seen', row_list(rows))


def activity_panel() -> str:
    rows = WC.J(WC.FIX / 'feed-60.json')['rows']
    return panel('Activity', row_list([(medallion(r['crest'], 32), WC.sentence(r['segments']), '', f'<span class="muted">{WC.ago(r["at"])}</span>') for r in rows]))


def character_page() -> str:
    ch = WC.elyra()
    c = ch.c
    meta = (f'Level {c["level"]} {escape(c["race"])} &middot; Holy {escape(c["class"])} &middot; <a href="#">&lt;{escape(c["guild"]["name"])}&gt;</a> &middot; '
            f'{escape(c["realm"])} &middot; {"Hardcore" if c["ruleset"] == "hardcore" else "Normal"} {c["region"].upper()}')
    right = f'<b class="num">+{ch.total():,.1f}</b> {ch.unit} in upgrades'
    band = titlebar(ch.cls, c['name'], meta, right=right, title_colour=D.CLASS_COLOR[ch.cls], top=header(*VIEWER, ''))
    main = gear_panel(ch) + fights_panel(ch) + activity_panel()
    rail = score_panel(ch) + builds_panel(ch)
    body = band + f'<main class="w ch"><div class="twocol"><div>{main}</div><aside>{rail}</aside></div></main>' + footer()
    return page(f'Forever Sixty, {c["name"]}', CHAR_CSS, body)


def characters_page() -> str:
    rows = WC.J(WC.FIX / 'search-thalrin.json')['rows']
    found = row_list([(medallion(r['class'].lower(), 36), f'<span style="color:{D.CLASS_COLOR[r["class"].lower()]}">{escape(r["name"])}</span>',
                       f'Level {r["level"]} {escape(r["spec"])} {escape(r["class"])} &middot; {"Hardcore" if r["ruleset"] == "hardcore" else "Normal"} {r["region"].upper()}',
                       '<a href="#">Open &rsaquo;</a>') for r in rows])
    search = field('Character name', '<a class="btn" href="#">Look up</a>') + settings_bar([('Region', 'US'), ('Ruleset', 'Any')]).replace('card setb', 'setb" style="margin-top:6px;padding:8px 2px 0')
    band = titlebar(None, 'Characters', '', top=header(*VIEWER, ''))
    main = panel('', search, framed=True) + panel('Thalrin', found)
    mine = WC.J(WC.HOME / 'me-60.json')['characters']
    yours = row_list([(medallion(c['class'].lower(), 32), f'<span style="color:{D.CLASS_COLOR[c["class"].lower()]}">{escape(c["name"])}</span>',
                       f'Level {c["level"]} {escape(c["spec"])} {escape(c["class"])}', '') for c in mine])
    body = band + f'<main class="w ch"><div class="twocol"><div>{main}</div><aside>{panel("Your characters", yours)}</aside></div></main>' + footer()
    return page('Forever Sixty, characters', CHAR_CSS, body)


def build() -> None:
    write('b2-character', character_page())
    write('b2-characters', characters_page())


if __name__ == '__main__':
    build()

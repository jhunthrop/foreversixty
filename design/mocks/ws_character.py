"""Web-layout mock: the character page (/character/<region>/<ruleset>/<slug>), every state, composed ONLY from ws_prims
primitives. Spec: design/specs/2026-11-06-character.md.

States (design/mocks/web/character-parity.json): visitor-60 (Elyra Duskvale, a level 60 Holy Priest), owner-34 (Felsnap,
a level 34 leveler with no logs, the owner), owner-60-premium (Elyra, the owner on Premium), owner-60-coach (the same with Coach history open: coach-history-60.json,
four typed reviews whose titles are the coach's own Holy output), no-score (Felsnap seen by a
visitor), unread (Felsnap's export carried no gear), not-found, elsewhere (the same name on other realms), loading,
error. Every figure is read from web/tests/e2e/fixtures/character/ and fixtures/home (character_fixtures.py writes
them), the files the parity spec routes; every gain is a listed runner-up's sim delta or the pick's score times the
band's reference rate (lib/home/upgrades.ts `slotGainFor`, the build's one gain rule).

  python3 design/mocks/character_fixtures.py && python3 design/mocks/ws_character.py
  node design/mocks/ws_shoot.mjs web-character-visitor-60 390,1440,1920,2000
"""
import json
from datetime import datetime
from pathlib import Path

from mocklib import CLASS_COLOR
from ws_lib import REGION
from ws_prims import (Disclosure, FeedRow, Figure, Link, Tag, PageHeader, Paperdoll, Plate, Row, Rows, Section,
                      Skeleton, build_page, crest_img, write)

ROOT = Path(__file__).resolve().parents[2]
FIX = ROOT / 'web/tests/e2e/fixtures/character'
HOME = ROOT / 'web/tests/e2e/fixtures/home'
LABELS = json.loads((ROOT / 'docs/contracts/fixtures/score-labels.json').read_text())
NOW = datetime.fromisoformat('2026-11-06T18:00:00+00:00')
SLOT_LABELS = {'head': 'Head', 'neck': 'Neck', 'shoulder': 'Shoulder', 'back': 'Back', 'chest': 'Chest', 'wrist': 'Wrist', 'hands': 'Hands',
               'waist': 'Waist', 'legs': 'Legs', 'feet': 'Feet', 'finger1': 'Finger 1', 'finger2': 'Finger 2', 'trinket1': 'Trinket 1',
               'trinket2': 'Trinket 2', 'main_hand': 'Main hand', 'off_hand': 'Off hand', 'ranged': 'Ranged'}
RARITY = {0: 'poor', 1: 'common', 2: 'uncommon', 3: 'rare', 4: 'epic', 5: 'legendary'}
PARSE = [(100, '#e5cc80'), (99, '#e268a8'), (95, '#ff8000'), (75, '#a335ee'), (50, '#0070ff'), (25, '#1eff00'), (0, '#8c8c8c')]
FIGHTS_SHOWN = 5
TIE = 0.05
CHAR_CSS = '''
.cmain{display:flex;flex-direction:column;gap:var(--section-gap)}
.sec-b{padding-top:12px}
.ans-figs{display:flex;align-items:flex-end;gap:24px;justify-content:flex-end}
.ans-links{display:flex;gap:16px;justify-content:flex-end;font-size:14px}
.pct{font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums;width:32px;text-align:right}
.head-sk{display:flex;align-items:center;gap:16px;padding-block:20px 16px;min-height:96px}
.plate.pad{padding:16px}
.err{display:flex;flex-wrap:wrap;align-items:center;gap:12px;min-height:44px;font-size:14px;color:var(--muted)}
.cell.sk{align-items:center}.cell.sk .bd{gap:8px}
.line{font-size:15px;line-height:24px;color:var(--text);padding-block:6px}
.stand{display:flex;flex-wrap:wrap;gap:4px 20px;padding-bottom:4px;font-size:15px;line-height:24px}
.frow .st{color:var(--strong)}
'''


def J(path: Path) -> dict:
    return json.loads(path.read_text())


def short_date(iso: str, year: bool = False) -> str:
    d = datetime.fromisoformat(iso.replace('Z', '+00:00'))
    return f'{d.day} {d.strftime("%b")}' + (f' {d.year}' if year else '')


def sec(title: str, body: str, action: str = '', meta: str = '') -> str:
    return Section(title, f'<div class="sec-b">{body}</div>', Link(f'{action} &rsaquo;') if action else '', meta)


def summary(text: str, action: str = '', tone: str = 'gold', standing: list[str] | None = None) -> str:
    act = f'<a class="link" href="#">{action} &rsaquo;</a>' if action else ''
    stand = ('<p class="stand">' + ''.join(f'<a class="link" href="#">{t} &rsaquo;</a>' for t in standing) + '</p>') if standing else ''
    return f'<div class="pc"><div data-region="{REGION["summary"]}"><p class="notice {tone}">{text}{act}</p>{stand}</div></div>'


# ------------------------------------------------------------------------------------------------ the alive pass (section 6)
TAGS = {'your_guild': 'Your guild', 'logs': 'Logs', 'rankings': 'Rankings', 'your_spec': 'Your spec'}
ROLE_WORD = {'dps': 'DPS', 'healer': 'healers', 'tank': 'tanks'}


def ago(iso: str) -> str:
    minutes = int((NOW - datetime.fromisoformat(iso.replace('Z', '+00:00'))).total_seconds() // 60)
    if minutes < 60:
        return f'{minutes} min ago'
    return f'{minutes // 60} h ago' if minutes < 1440 else f'{minutes // 1440} d ago'


def sentence(segments: list[dict]) -> str:
    out = ''
    for seg in segments:
        if seg.get('style') == 'name':
            out += f'<b style="color:{CLASS_COLOR[seg["class"]]}">{seg["text"]}</b>'
        elif seg.get('style') == 'strong':
            out += f'<b class="st">{seg["text"]}</b>'
        else:
            out += seg['text']
    return out


def activity_section(feed: str | None) -> str:
    if not feed:
        return ''
    rows = J(FIX / feed)['rows']
    if not rows:
        return ''
    return sec('Activity', ''.join(FeedRow(crest_img(r['crest'], 40), sentence(r['segments']), ago(r['at']), TAGS[r['tag']]) for r in rows))


def standing(home: str | None, guild: bool, spec: str) -> list[str]:
    if not home:
        return []
    a = J(FIX / home)
    out = []
    gs = a.get('guild_sims')
    if gs:
        own = gs['you'] or next((r for r in gs['rows'] if r['you']), None)
        if own:
            band = '60' if gs['band'] >= 60 else f"{gs['band']} to {gs['band'] + 9}"
            out.append(f"#{own['rank']} of {gs['of']} {ROLE_WORD[gs['role']]} in {gs['guild_name']}, {band}{' raid' if gs['preset'] == 'raid' else ''}")
    w = a.get('spec_week')
    if w:
        out.append(f"{spec} is graded {w['grade']} this week" if w['grade'] else f"{spec} is #{w['place']} of {w['of']} this week")
    if guild:
        r = J(HOME / 'guild-home-60.json')['next_raid']
        left = datetime.fromisoformat(r['starts_at'].replace('Z', '+00:00')) - NOW
        d, rem = divmod(int(left.total_seconds()) // 60, 1440)
        out.append(f"{r['title']} in {d} days, {rem // 60} hours, {rem % 60} min, {r['counts']['in']} in")
    return out


def ledger(*sections: str) -> str:
    return f'<div class="pc" style="padding-top:20px"><div class="cmain">{"".join(s for s in sections if s)}</div></div>'


def figure_text(value: float) -> str:
    r = round(value, 1)
    return ('+' if r > 0 else '') + f'{abs(r):,.1f}' if r >= 0 else f'&minus;{abs(r):,.1f}'


# ------------------------------------------------------------------------------------------------ the data
class Char:
    """One character answer with its band and items, priced the build's way."""

    def __init__(self, answer: str, bis: Path, items: Path):
        self.page = J(FIX / answer)
        self.c = self.page['character']
        self.cls = self.c['class'].lower()
        self.band = J(bis)['bands'][0]
        self.items = {i['id']: i for i in J(items)['items']}
        self.unit = 'HPS' if self.band.get('role') == 'healer' else 'DPS'

    def gain(self, g: dict) -> tuple[str, float | None] | None:
        """('best' | 'gain' | 'tie', value) for a slot with a pick, as lib/gear/paperdoll.ts upgradeFor + slotGainFor."""
        pick, worn = g['pick'], g['worn']
        if pick is None:
            return None
        worn_ids = {x['slot']: (x['worn'] or {}).get('id') for x in self.page['gear']}
        pair = {'finger1': 'finger2', 'finger2': 'finger1', 'trinket1': 'trinket2', 'trinket2': 'trinket1'}.get(g['slot'])
        if worn and (worn['id'] == pick['id'] or (pair and worn_ids.get(pair) == pick['id'])):
            return ('best', None)
        if not g['gap'] or g['gap'] <= 0:
            return None
        slot = next(s for s in self.band['slots'] if s.get('item_id') == pick['id'])
        alt = next((a for a in slot.get('alternatives') or [] if worn and a['item_id'] == worn['id'] and a.get('dps_delta', 0) != 0), None)
        if alt:
            value = -alt['dps_delta']
        elif g['slot'] in ('main_hand', 'off_hand', 'ranged') or slot.get('score') is None or not self.band.get('reference_dps_per_point'):
            return ('gain', None)
        else:
            value = slot['score'] * self.band['reference_dps_per_point']
        return ('tie', value) if abs(value) < TIE else ('gain', value)

    def total(self) -> float:
        return sum(v for k, v in filter(None, (self.gain(g) for g in self.page['gear'])) if k == 'gain' and v is not None)

    def top(self) -> dict | None:
        priced = [(v, g) for g in self.page['gear'] for k, v in [self.gain(g) or ('', None)] if k == 'gain' and v is not None]
        return max(priced, key=lambda p: p[0])[1]['pick'] if priced else None


def cell(ch: Char, g: dict, gains: bool) -> dict:
    label = SLOT_LABELS[g['slot']]
    worn = g['worn']
    if worn is None:
        cover = next((x for x in ch.page['gear'] if x['slot'] == 'main_hand'), None) if g['slot'] == 'off_hand' else None
        two = cover and cover['worn'] and ch.items.get(cover['worn']['id'], {}).get('two_hand')
        return {'label': label, 'flag': None, 'state': 'empty', 'name': 'Two-hand in main hand' if two else 'Empty', 'line2': ''}
    it = ch.items.get(worn['id'], {})
    c = {'label': label, 'flag': None, 'state': 'item', 'name': worn['name'], 'quality': RARITY[it.get('quality', 1)], 'icon': it.get('icon', ''),
         'stats': '', 'line2': f"Item level {worn['item_level']}", 'gain': None, 'basis': None}
    up = ch.gain(g) if gains else None
    if up is None:
        return c
    kind, value = up
    pick = g['pick']
    basis = {'prefix': 'vs', 'item': pick['name'], 'quality': RARITY[ch.items.get(pick['id'], {}).get('quality', 1)]}
    if kind == 'best':
        return {**c, 'gain': {'value': 'Best in slot', 'unit': ''}}
    if kind == 'tie':
        return {**c, 'gain': {'value': 'Same', 'unit': ''}, 'basis': basis}
    if value is None:
        return {**c, 'gain': {'value': '&mdash;', 'unit': ''}, 'basis': {'prefix': 'Not sim-checked', 'item': '', 'quality': 'common'}}
    return {**c, 'gain': {'value': figure_text(value), 'unit': ch.unit}, 'basis': basis}


def pick_cell(ch: Char, g: dict) -> dict | None:
    """The unread doll: the band's pick and its source (no worn item, no gain)."""
    pick = g['pick']
    if pick is None:
        return None
    it = ch.items.get(pick['id'], {})
    return {'label': SLOT_LABELS[g['slot']], 'flag': None, 'state': 'item', 'name': pick['name'], 'quality': RARITY[it.get('quality', 1)],
            'icon': it.get('icon', ''), 'stats': '', 'line2': pick['source'], 'gain': None, 'basis': None}


LEFT = ['head', 'neck', 'shoulder', 'back', 'chest', 'wrist', 'main_hand', 'off_hand', 'ranged']
RIGHT = ['hands', 'waist', 'legs', 'feet', 'finger1', 'finger2', 'trinket1', 'trinket2']


def doll(cells: dict) -> str:
    return Paperdoll('Gear', 'slot', columns=([cells[s] for s in LEFT if s in cells], [cells[s] for s in RIGHT if s in cells]))


# ------------------------------------------------------------------------------------------------ the page
def head(ch: Char, owner: dict | None, answer: bool = True) -> str:
    c = ch.c
    spec = {'priest': 'Holy', 'hunter': 'Beast Mastery'}[ch.cls]
    split = (owner and '/'.join(map(str, owner['build']['talents']['points'])) if owner and any(owner['build']['talents']['points'])
             else (ch.page['builds_seen'][0]['talent_split'] if ch.page['builds_seen'] else None))
    meta = [f"Level {c['level']} {c['race']}", f"{spec} {c['class']}"]
    if split:
        meta.append(f'<a class="link" href="#">{split}</a>')
    if c.get('guild'):
        meta.append(f"<a class=\"link\" href=\"#\">&lt;{c['guild']['name']}&gt;</a>")
    meta += [c['realm'], f"{'Hardcore' if c['ruleset'] == 'hardcore' else 'Normal'} {c['region'].upper()}"]
    if ch.page.get('supporter'):
        meta.append(Tag('Supporter'))
    if owner:
        meta.append('Addon &middot; 1 day ago')
    figs = ''
    total = ch.total()
    if answer and total > 0:
        score = Figure(f"{ch.page['score']:,}", 'Forever Score', size='sm') if ch.page.get('score') else ''
        figs = (f'<div class="ans-figs">{score}{Figure(figure_text(total), "Upgrades available", ch.unit, size="big")}</div>'
                f'<div class="ans-links">{Link("Open in simulator &rsaquo;")}{Link("Talents in planner &rsaquo;")}</div>')
    return PageHeader(c['name'], CLASS_COLOR[ch.cls], crest_img(ch.cls, 56), meta, figs)


def gear_section(ch: Char, gains: bool) -> str:
    cells = {g['slot']: cell(ch, g, gains) for g in ch.page['gear']}
    band = ch.c['level'] // 10 * 10
    meta = 'Level 60 raid' if band >= 60 else f'Levels {band} to {band + 9}'
    return sec('Gear', Plate(doll(cells), 'Gear'), 'Best in slot', meta)


def score_section(ch: Char, owner: bool) -> str:
    p = ch.page
    if not p.get('score'):
        line = 'No logs yet. Upload a log to get a rating.' if owner else 'No logs yet.'
        return sec('Forever Score', f'<p class="line">{line}</p>')
    rating = J(FIX / 'rating-60.json')
    tier = LABELS['tiers'][p['tier']]
    figs = (f'<div class="figs" style="--c:3">{Figure(f"{p["score"]:,}", "Score")}'
            f'<div class="fig"><span class="t-label">Tier</span><span class="v" style="color:{tier["color"]}">{tier["label"]}</span></div>'
            f'{Figure(str(rating["latest"]["overall"]), "Rating", "of 100")}</div>')
    comps = p['components']
    maxes = {'gear': 500, 'experience': 700, 'parses': 600, 'attendance': 200, 'discipline': 100}
    gaps = sorted(((maxes[k] - comps[k], k) for k in maxes if maxes[k] - comps[k] > 0), key=lambda x: -x[0])
    label = {x['key']: x['label'] for x in LABELS['components']}
    most = f'<p class="line">Most to gain: {label[gaps[0][1]]}, {gaps[0][0]} points to the top. {Link("How it is scored &rsaquo;")}</p>'
    parts = Rows([Row(label[k], f'{comps[k]} of {maxes[k]}') for k in maxes], open_=True)
    return sec('Forever Score', figs + most + Disclosure('Score parts', parts), meta=f"{rating['sample_size']} rated fights")


def parse_color(pct: float) -> str:
    return next(c for floor, c in PARSE if round(pct) >= floor)


def fights_section(ch: Char) -> str:
    p = ch.page
    rows = []
    for f in p['best'][:FIGHTS_SHOWN]:
        trail = (f'<span class="pct" style="color:{parse_color(f["percentile"])}">{round(f["percentile"])}</span>'
                 + Figure(f"{f['value']:,}", '', ch.unit, size='sm'))
        rows.append(Row(f['encounter'], f"{short_date(f['fought_at'])} &middot; {f['spec']}", trail=trail))
    return sec('Best per encounter', Rows(rows, open_=True), f"All ranked fights ({len(p['history'])})")


def builds_section(ch: Char) -> str:
    p = ch.page
    rows = [Row(b['talent_split'], f"{b['spec']} &middot; first seen {short_date(b['first_seen'], True)}") for b in p['builds_seen'][:3]]
    return sec('Builds seen', Rows(rows, open_=True))


def coach_rows(file: str) -> str:
    """The opened Coach history (coach/CoachHistoryList): link Rows, newest first: the fix-first title, the pull's
    date, boss and spec muted under it, Wipe and Log only as Tags, a muted chevron."""
    rows = []
    for r in J(FIX / file)['reviews']:
        title = (r.get('fix_first') or {}).get('title') or 'Nothing to fix'
        tags = ('' if r['kill'] else Tag('Wipe', 'wipe')) + (Tag('Log only', 'muted') if r['source'] == 'rules' else '')
        trail = tags + '<span class="t-small" aria-hidden="true">&rsaquo;</span>'
        rows.append(Row(title, ' &middot; '.join([short_date(r['generated_at']), r['boss'], r['spec']]), trail=trail))
    return f'<section aria-label="Your coach">{Rows(rows, open_=True, list_=True)}</section>'


def owner_sections(premium: bool, coach: str | None = None) -> str:
    if not premium:
        return sec('History', Rows([Row('Score over time and parse per boss', 'Premium keeps every rated fight.', trail=Link('See plans &rsaquo;'), link=False)], open_=True))
    history = sec('History', Disclosure('Score over time, then parse per boss', ''))
    body = coach_rows(coach) if coach else ''
    return history + sec('Coach history', Disclosure('Every review the coach wrote for this character', body, open_=coach is not None))


def next_line(ch: Char, owner: dict) -> tuple[str, str]:
    earned = max(0, owner['level'] - 9)
    to_spend = earned - sum(owner['build']['talents']['points'])
    top = ch.top()
    if to_spend > 0:
        then = f", then get {top['name']}" if top else ''
        return f'Next: spend your {to_spend} talent points{then}.', 'Open the planner'
    if top:
        return f"Next: get {top['name']} from {top['source']}.", 'Open best in slot'
    return 'Nothing waiting.', ''


# ------------------------------------------------------------------------------------------------ the states
def elyra() -> Char:
    return Char('character-60.json', FIX / 'bis-priest-holy.json', FIX / 'items-priest.json')


def felsnap(answer: str = 'character-34.json') -> Char:
    return Char(answer, HOME / 'bis-hunter-beast-mastery.json', HOME / 'items-hunter.json')


def state_ready(ch: Char, me: str | None, premium: bool = False, feed: str | None = None, home: str | None = None,
                coach: str | None = None) -> str:
    owner = J(FIX / me)['characters'][0] if me else None
    body = head(ch, owner)
    if owner:
        spec = {'priest': 'Holy', 'hunter': 'Beast Mastery'}[ch.cls]
        body += summary(*next_line(ch, owner), standing=standing(home, 'guild' in owner, spec))
    logged = bool(ch.page.get('score'))
    body += ledger(gear_section(ch, owner is not None), score_section(ch, owner is not None),
                   fights_section(ch) if ch.page['best'] else '', builds_section(ch) if ch.page['builds_seen'] else '',
                   activity_section(feed), owner_sections(premium, coach) if owner and logged else '')
    return body


def state_unread() -> str:
    ch = felsnap('character-34-unread.json')
    owner = J(FIX / 'me-34-unread.json')['characters'][0]
    cells = {g['slot']: c for g in ch.page['gear'] if (c := pick_cell(ch, g))}
    return (head(ch, owner, answer=False)
            + summary('Your gear has not been read yet. In game, log in, wait for the world to load, then log out once; it appears here within a minute.', tone='neutral',
                      standing=standing('home-34.json', False, 'Beast Mastery'))
            + ledger(sec('Best in slot for level 30 to 39', Plate(doll(cells), 'Gear'), 'Best in slot'), score_section(ch, True)))


def lookup() -> str:
    """The built FormRow: a form holding the name Field (a textbox, filled) and the one submit Button."""
    form = ('<form class="formrow" aria-label="Look up a character"><label class="f grow"><span class="t-label">Name</span>'
            '<input class="field grow" value="Thalrin"></label><span><button type="submit" class="btn primary">Look up</button></span></form>')
    return sec('Look up a character', form)


def state_missing(elsewhere: bool) -> str:
    head_ = PageHeader('Thalrin')
    rows = ''
    if elsewhere:
        found = J(FIX / 'search-thalrin.json')['rows']
        rows = sec('Same name elsewhere', Rows([Row(r['name'], f"Level {r['level']} {r['spec']} {r['class']} &middot; {'Hardcore' if r['ruleset'] == 'hardcore' else 'Normal'} {r['region'].upper()}",
                                                    lead=crest_img(r['class'].lower(), 40), title_color=CLASS_COLOR[r['class'].lower()]) for r in found], open_=True))
    add = sec('Add this character', Rows([
        Row('Sign in to see your characters', 'Battle.net lists every character on your account.'),
        Row('Get the addon and companion', 'Export with the addon and the companion to add this character.'),
    ], open_=True))
    return head_ + summary('No character named Thalrin on US Hardcore.', tone='neutral') + ledger(rows, lookup(), add)


def skeleton_cell() -> str:
    return (f'<div class="cell sk" aria-hidden="true"><span class="lb">{Skeleton(48, 12)}</span>{Skeleton(40, 40)}'
            f'<span class="bd">{Skeleton(160, 16)}{Skeleton(224, 12)}</span>{Skeleton(40, 20)}</div>')


def state_loading() -> str:
    head_ = f'<div class="pc"><div data-region="{REGION["head"]}" class="head-sk" role="status" aria-busy="true">{Skeleton(56, 56, True)}{Skeleton(280, 32)}</div></div>'
    cells = ''.join(skeleton_cell() for _ in range(17))
    plate = Plate(f'<div class="doll" data-region="{REGION["slots"]}" aria-busy="true"><div class="doll-cols" role="presentation">{cells}</div></div>', 'Gear')
    return head_ + ledger(sec('Gear', plate))


def state_error() -> str:
    plate = Plate('<div class="err" role="alert"><span>We could not load this character.</span><button type="button" class="btn sm">Try again</button></div>', 'Character')
    return PageHeader('Elyra Duskvale') + ledger(plate.replace('class="plate"', 'class="plate pad"'))


STATES = {
    'visitor-60': (lambda: state_ready(elyra(), None, feed='feed-60.json'), None),
    'owner-34': (lambda: state_ready(felsnap(), 'me-34.json', home='home-34.json'), ('Felsnap', 'hunter', 'Level 34')),
    'owner-60-premium': (lambda: state_ready(elyra(), 'me-60.json', premium=True, feed='feed-60-owner.json', home='home-60.json'),
                         ('Elyra Duskvale', 'priest', 'Level 60')),
    'owner-60-coach': (lambda: state_ready(elyra(), 'me-60.json', premium=True, feed='feed-60-owner.json', home='home-60.json',
                                           coach='coach-history-60.json'), ('Elyra Duskvale', 'priest', 'Level 60')),
    'no-score': (lambda: state_ready(felsnap(), None), None),
    'unread': (state_unread, ('Felsnap', 'hunter', 'Level 34')),
    'not-found': (lambda: state_missing(False), None),
    'elsewhere': (lambda: state_missing(True), None),
    'loading': (state_loading, None),
    'error': (state_error, None),
}


def main() -> None:
    for state, (render, who) in STATES.items():
        write(f'web-character-{state}.html', build_page('Forever Sixty', render() + f'<style>{CHAR_CSS}</style>', who=who))


if __name__ == '__main__':
    main()

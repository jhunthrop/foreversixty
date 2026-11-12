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
from ws_lib import QCOL
from ws_prims import (Button, Disclosure, Facts, FeedRow, Figure, ItemRow, Link, MeterClass, Page, PageHeader, PaperdollFrame, Spark,
                      Ordered, Plate, Row, Rows, RowsSlim, Section, Skeleton, Tag, build_page, crest_img, icon, write)

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
UPGRADES_SHOWN = 6
TIE = 0.05
CHAR_CSS = '''
.cmain{display:flex;flex-direction:column;gap:var(--section-gap)}
.sec-b{padding-top:12px}
.ans-figs{display:flex;align-items:flex-end;gap:24px;justify-content:flex-end}
.ans-links{display:flex;gap:16px;justify-content:flex-end;font-size:14px}
.pct{font:700 15px/20px var(--font-body);font-variant-numeric:tabular-nums;width:32px;text-align:right}
.plate.pad{padding:16px}
.err{display:flex;flex-wrap:wrap;align-items:center;gap:12px;min-height:44px;font-size:14px;color:var(--muted)}
.cell.sk{align-items:center}.cell.sk .bd{gap:8px}
.line{font-size:15px;line-height:24px;color:var(--text);padding-block:6px}
@media(max-width:767px){.wait-pre{display:none}}.own-in{margin-top:8px;padding-top:4px;border-top:1px solid var(--line-soft)}.bcells{display:flex;align-items:center;gap:12px}.bcells.narrow{margin-top:4px}@media(min-width:768px){.bcells.narrow{display:none}}@media(max-width:767px){.bcells.wide{display:none}}.bsince{font-size:13px;color:var(--muted);white-space:nowrap}.dsum-end{margin-left:auto;display:inline-flex;align-items:center}
.dspk{display:inline-block;width:64px}.bspk{display:inline-block;width:48px}.wait-plate{margin-top:16px;width:100%;padding:0 16px}.disc.dlink>summary{color:var(--gold);font-size:15px}.disc.dlink>summary::before{display:none}.wait{display:block;overflow:hidden;white-space:nowrap;text-overflow:ellipsis;color:var(--gold);font-weight:600}
.sum-row{display:flex;flex-wrap:wrap;align-items:baseline;column-gap:20px}.notice.facts{gap:4px 20px}
.frow .st{color:var(--strong)}
/* round 5 (lane P9): the split from 1280, the frame centred at 640, the answer's sim link beside its figure, the parts' bars */
.char-pg{padding-bottom:48px}.char-pg .frow .k:empty{display:none}
.doll-c{display:flex;justify-content:center}.doll-w{width:100%;max-width:640px}
.ans-row{display:flex;align-items:center;gap:20px}
.itip.up-box{max-width:340px;text-align:left}.itip.up-box>.bd{gap:0;padding:10px 12px}
.up-h{margin:0 0 4px;font:600 13px/20px var(--font-body);color:var(--strong)}
.up-box .irow-box{padding:0}.itip.up-box .irow{min-height:0;padding:0;column-gap:12px}
.itip.up-box .irow .bd{display:block;padding:0}.up-box .irow .nm{font-size:15px;line-height:20px}.up-box .irow .src{font-size:13px;line-height:18px}
.up-box .irow .g .v{font-size:12px;line-height:16px;color:#32ae88}.up-box .irow .g .v small{font-size:12px;color:#32ae88;margin-left:3px}
.up-more{margin-top:4px;font-size:13px}
.parts{margin-top:12px}.part{display:flex;align-items:center;gap:16px}.part .trk{width:120px}
.part .pts{width:40px;text-align:right;font-size:14px;color:var(--strong);font-variant-numeric:tabular-nums}
.nb{font-weight:400}
.fws{display:inline-flex;align-items:center;justify-content:flex-end;gap:12px}.fsp{width:64px}.fsp .spark{display:block;width:100%;height:24px}
'''


def J(path: Path) -> dict:
    return json.loads(path.read_text())


def short_date(iso: str, year: bool = False) -> str:
    d = datetime.fromisoformat(iso.replace('Z', '+00:00'))
    return f'{d.day} {d.strftime("%b")}' + (f' {d.year}' if year else '')


def sec(title: str, body: str, action: str = '', meta: str = '') -> str:
    return Section(title, f'<div class="sec-b">{body}</div>', Link(f'{action} &rsaquo;') if action else '', meta)


def summary(text: str, action: str = '', tone: str = 'gold') -> str:
    act = f'<a class="link" href="#">{action} &rsaquo;</a>' if action else ''
    return f'<div class="pc"><div class="sum-row" data-region="{REGION["summary"]}"><p class="notice {tone}">{text}{act}</p></div></div>'


def waiting_summary(facts: list[tuple[str, str]]) -> str:
    """The owner's waiting facts (web characterWaiting) on a plate (h-audit 2026-11-09): a slim link Row each, the upgrade led by
    its item's icon, the text one gold line that truncates."""
    if not facts:
        return ''
    rows = RowsSlim([Row(f'<span class="wait">{t}</span>', lead=lead, slim=True) for t, lead in facts])
    return (f'<div class="pc"><div class="sum-row" data-region="{REGION["summary"]}">'
            f'<div class="plate wait-plate">{rows}</div></div></div>')


def standing_section(lines: list[tuple[str, str]]) -> str:
    """The owner's standing in the rail (character/CharacterStanding, h-audit 2026-11-09): a slim link Row a line, label left, value right."""
    return panel('Standing', RowsSlim([Row(f'<span class="nb">{label}</span>', trail=f'<b class="st">{value}</b>', slim=True)
                                       for label, value in lines])) if lines else ''


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


def own_segments(segments: list[dict], name: str) -> list[dict]:
    """The character's own name out of its own feed rows (web ownActivitySegments): the text after it opens with a capital."""
    at = next((i for i, s in enumerate(segments) if s.get('style') == 'name' and s['text'] == name), None)
    if at is None:
        return segments
    rest = segments[at + 1:]
    if not rest or rest[0].get('style'):
        return segments[:at] + rest
    text = rest[0]['text'].lstrip()
    opened = text[:1].upper() + text[1:] if at == 0 else text
    return segments[:at] + [{**rest[0], 'text': opened}] + rest[1:]


def activity_section(feed: str | None, name: str) -> str:
    if not feed:
        return ''
    rows = J(FIX / feed)['rows']
    if not rows:
        return ''
    return panel('Activity', ''.join(FeedRow(crest_img(r['crest'], 32), sentence(own_segments(r['segments'], name)), ago(r['at']), '')
                                     for r in rows))


def standing(home: str | None, guild: bool, spec: str) -> list[tuple[str, str]]:
    if not home:
        return []
    a = J(FIX / home)
    out = []
    gs = a.get('guild_sims')
    if gs:
        own = gs['you'] or next((r for r in gs['rows'] if r['you']), None)
        if own:
            band = '60' if gs['band'] >= 60 else f"{gs['band']} to {gs['band'] + 9}"
            out.append((f"{gs['guild_name']} &middot; {band}{' raid' if gs['preset'] == 'raid' else ''}", f"#{own['rank']} of {gs['of']} {ROLE_WORD[gs['role']]}"))
    w = a.get('spec_week')
    if w:
        out.append((f"{spec} tier", w['grade'] if w['grade'] else f"#{w['place']} of {w['of']}"))
    if guild:
        r = J(HOME / 'guild-home-60.json')['next_raid']
        left = datetime.fromisoformat(r['starts_at'].replace('Z', '+00:00')) - NOW
        d, rem = divmod(int(left.total_seconds()) // 60, 1440)
        out.append((r['title'], f"{d} days, {rem // 60} hours"))
    return out


def ledger(*sections: str) -> str:
    """Page layout ledger; `char-pg` carries the built main's section-gap pad under the page, as the split states do."""
    return f'<div class="char-pg"><div class="pc" style="padding-top:20px"><div class="cmain">{"".join(s for s in sections if s)}</div></div></div>'


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


def doll_slots(ch: Char, picks: bool = False) -> list[dict]:
    """The frame's filled sockets (lib/character/doll.ts): worn items, the off hand under a two-hander as covered; `picks`, the
    unread state, the band's pick for every slot that has one."""
    out = []
    worn = {g['slot']: g['worn'] for g in ch.page['gear']}
    for g in ch.page['gear']:
        item = g['pick'] if picks else g['worn']
        label = SLOT_LABELS[g['slot']]
        if item is None:
            main = worn.get('main_hand')
            if picks or g['slot'] != 'off_hand' or not main or not ch.items.get(main['id'], {}).get('two_hand'):
                continue
            it = ch.items.get(main['id'], {})
            out.append({'slot': 'off_hand', 'label': f"{label}: {main['name']}, two-handed", 'covered': True, 'selected': False,
                        'icon': it.get('icon'), 'quality': it.get('quality', 1)})
            continue
        it = ch.items.get(item['id'], {})
        out.append({'slot': g['slot'], 'label': f"{label}: {item['name']}", 'covered': False, 'selected': False,
                    'icon': it.get('icon'), 'quality': it.get('quality', 1)})
    return out


def upgrade_rows(ch: Char) -> list[tuple[dict, str, float]]:
    """The band's picks with a priced gain over what is worn, largest first (lib/character/upgrades.ts)."""
    rows = [(g['pick'], SLOT_LABELS[g['slot']], v) for g in ch.page['gear'] for k, v in [ch.gain(g) or ('', None)]
            if k == 'gain' and v is not None and v >= TIE]
    return sorted(rows, key=lambda r: -r[2])


def upgrades_box(ch: Char) -> str:
    """The frame's centre: the client tooltip ground holding the first six Upgrades, then what the rest add."""
    rows = upgrade_rows(ch)
    if not rows:
        return ''
    items = ''.join(ItemRow(pick['name'], ch.items.get(pick['id'], {}).get('quality', 1), icon(ch.items.get(pick['id'], {}).get('icon'), 28),
                            label, gain=(figure_text(v), ch.unit, '')) for pick, label, v in rows[:UPGRADES_SHOWN])
    rest = rows[UPGRADES_SHOWN:]
    more = (f'<a class="link up-more" href="#">{len(rest)} more, {figure_text(sum(v for _, _, v in rest))} {ch.unit} &rsaquo;</a>'
            if rest else '')
    return f'<div class="itip up-box" role="group" aria-label="Upgrades"><div class="bd"><h3 class="up-h">Upgrades</h3>{items}{more}</div></div>'


def doll(ch: Char, picks: bool = False) -> str:
    centre = '' if picks else upgrades_box(ch)
    frame = PaperdollFrame(doll_slots(ch, picks), f'{SPEC[ch.cls]} {ch.c["class"]}', f"Level {ch.c['level']}", None, centre,
                           'Gear by slot', list_centre=centre, heading='h2')
    return f'<div class="doll-c"><div class="doll-w">{frame}</div></div>'


# ------------------------------------------------------------------------------------------------ the page
SPEC = {'priest': 'Holy', 'hunter': 'Beast Mastery'}


def head(ch: Char, owner: dict | None, answer: bool = True) -> str:
    c = ch.c
    split_ = (owner and '/'.join(map(str, owner['build']['talents']['points'])) if owner and any(owner['build']['talents']['points'])
              else (ch.page['builds_seen'][0]['talent_split'] if ch.page['builds_seen'] else None))
    meta = [f"Level {c['level']} {c['race']}", f"{SPEC[ch.cls]} {c['class']}"]
    if split_:
        meta.append(f'<a class="link" href="#">{split_}</a>')
    if c.get('guild'):
        meta.append(f"<a class=\"link\" href=\"#\">&lt;{c['guild']['name']}&gt;</a>")
    meta += [c['realm'], f"{'Hardcore' if c['ruleset'] == 'hardcore' else 'Normal'} {c['region'].upper()}"]
    if owner:
        meta.append('Synced 1 day ago')
    total = ch.total()
    figure = Figure(figure_text(total), '', f'{ch.unit} from upgrades') if total > 0 else ''
    ans = f'<span class="ans-row">{Button("Sim this character", small=True)}{figure}</span>' if answer else ''
    return PageHeader(c['name'], CLASS_COLOR[ch.cls], crest_img(ch.cls, 56), meta, ans)


def panel(title: str, body: str, action: str = '', frame: str = 'plain') -> str:
    return Section(title, body, Link(f'{action} &rsaquo;') if action else '', panel=True, frame=frame)


def gear_section(ch: Char, title: str = 'Gear', picks: bool = False) -> str:
    return panel(title, doll(ch, picks), 'Best in slot', frame='framed')


PART_FLOORS = [(0.9, 4), (0.75, 3), (0.5, 2)]
PART_MAX = {'gear': 500, 'experience': 700, 'parses': 600, 'attendance': 200, 'discipline': 100}


def spark_fact(values: list[float], text: str) -> str:
    """A fact's value with its trend: a 64 px spark of the series over its lowest point, then the figure
    (character/CharacterScore `withSpark`, lib/character/score-parts.ts `fromLowest`)."""
    low = min(values) if values else 0
    spark = f'<span class="fsp">{Spark([v - low for v in values])}</span>' if len(values) > 1 else ''
    return f'<span class="fws">{spark}<span>{text}</span></span>'


def score_section(ch: Char, owner: bool) -> str:
    """The rail's Forever Score (character/CharacterScore): Score (in its tier's colour, with the tier word) and Rating as Facts, then the parts."""
    p = ch.page
    if not p.get('score'):
        line = 'No logs yet. Upload a log to get a rating.' if owner else 'No logs yet.'
        return panel('Forever Score', f'<p class="line">{line}</p>')
    rating = J(FIX / 'rating-60.json')
    tier = LABELS['tiers'][p['tier']]
    history = sorted(p.get('score_history') or [], key=lambda x: x['date'])
    trend = ([(f"Since {short_date(history[0]['date'] + 'T00:00:00Z')}",
               spark_fact([x['score'] for x in history], f"+{history[-1]['score'] - history[0]['score']:,}".replace(',', '')))]
             if len(history) > 1 else [])
    # The tier is the score's colour and its word on the Score row (h-audit 2026-11-09), never a row of its own.
    score = f'<span style="color:{tier["color"]}">{p["score"]:,} <span class="t-small" style="color:inherit;font-weight:600">{tier["label"]}</span></span>'
    facts = Facts([('Score', score), *trend, ('Rating', f"{rating['latest']['overall']} of 100")])
    comps = p['components']
    label = {x['key']: x['label'] for x in LABELS['components']}
    longest = max(PART_MAX.values())
    rows = []
    for k, mx in PART_MAX.items():
        fill = min(1, comps[k] / mx)
        q = next((qq for floor, qq in PART_FLOORS if fill >= floor), 1)
        trail = f'<span class="part"><span class="trk">{MeterClass(fill * mx / longest, QCOL[q])}</span><span class="pts">{comps[k]}</span></span>'
        rows.append(Row(f'<span class="nb">{label[k]}</span>', trail=trail, link=False, slim=True))
    return panel('Forever Score', facts + f'<div class="parts">{RowsSlim(rows)}</div>')


def parse_color(pct: float) -> str:
    return next(c for floor, c in PARSE if round(pct) >= floor)


def fights_section(ch: Char, extra: str = '') -> str:
    p = ch.page
    rows = []
    for f in p['best'][:FIGHTS_SHOWN]:
        # The boss kill by kill (web bestRows, h-audit 2026-11-09): the parse trend and the change since the kill before.
        kills = sorted((h for h in p['history'] if h['encounter_id'] == f['encounter_id'] and h['difficulty'] == f['difficulty']
                        and h['metric'] == f['metric'] and h.get('kill') is not False), key=lambda h: h['fought_at'])
        parses = [h['percentile'] for h in kills if h.get('percentile') is not None]
        spark = f'<span class="bspk" aria-hidden="true">{Spark(parses)}</span>' if len(parses) > 1 else ''
        before = [h for h in kills if h['fought_at'] < f['fought_at']]
        change = round(f['value'] - before[-1]['value']) if before else None
        # The change since the kill before sits on the parse line (h-audit 2026-11-09).
        since = (f"<span class=\"bsince\">{'+' if change > 0 else '&minus;' if change < 0 else ''}{abs(change)} vs "
                 f"{short_date(before[-1]['fought_at'])}</span>") if before else ''
        cells = since + spark
        # The delta and spark sit in the trail from 768 and on their own line under the boss below it.
        trail = ((f'<span class="bcells wide">{cells}</span>' if cells else '')
                 + f'<span class="pct" style="color:{parse_color(f["percentile"])}">{round(f["percentile"])}</span>'
                 + Figure(f"{f['value']:,}", '', ch.unit, size='sm'))
        phone = f'<span class="bcells narrow">{cells}</span>' if cells else ''
        ex = f" &middot; Execution {round(f['execution_score'] * 100)}%" if f.get('execution_score') is not None else ''
        spec = '' if f['spec'] == SPEC[ch.cls] else f" &middot; {f['spec']}"  # the page's own spec goes without saying
        rows.append(Row(f['encounter'], f"{short_date(f['fought_at'])}{spec}{ex}{phone}", trail=trail, slim=True))
    return panel('Best per encounter', RowsSlim(rows) + extra, f"All fights ({len(p['history'])})")


def builds_section(ch: Char) -> str:
    rows = [Row(b['talent_split'], f"{b['spec']} &middot; {short_date(b['first_seen'])}", slim=True)
            for b in ch.page['builds_seen'][:3]]
    return panel('Builds seen', RowsSlim(rows))


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


def owner_sections(premium: bool, fights: int, coach: str | None = None, latest: dict | None = None, in_panel: bool = False,
                   scores: list[float] | None = None) -> str:
    scores = scores or []
    """One History panel (h-audit 2026-11-09): the charts' Disclosure with the fight count, then the coach's reviews."""
    wrap = (lambda body: f'<div class="own-in">{body}</div>') if in_panel else (lambda body: panel('History', body))
    if not premium:
        return wrap(Rows([Row('Score over time and parse per boss', 'Premium keeps every rated fight.',
                              trail=Link('See plans &rsaquo;'), link=False)], open_=True))
    body = coach_rows(coach) if coach else ''
    summary = f"Score {latest['score']:,} on {short_date(latest['date'] + 'T00:00:00Z')}" if latest else 'Score and parses'
    # History and Coach as gold link rows that open in place (web Disclosure `link`, h-audit 2026-11-09).
    link = lambda html: html.replace('<details class="disc"', '<details class="disc dlink"', 1).replace(
        '</summary>', '<span aria-hidden="true">&rsaquo;</span></summary>', 1)
    history = link(Disclosure('History', '', meta=f"{fights} {'fight' if fights == 1 else 'fights'}"))
    return wrap(history + link(Disclosure('Coach', body, open_=coach is not None)))


def waiting(ch: Char, owner: dict) -> list[tuple[str, str]]:
    """(text, lead icon html): the upgrade carries its item's icon."""
    earned = max(0, owner['level'] - 9)
    to_spend = earned - sum(owner['build']['talents']['points'])
    top = ch.top()
    talents = [(f"{to_spend} talent {'point' if to_spend == 1 else 'points'} unspent", '')] if to_spend > 0 else []
    item = ch.items.get(top['id'], {}) if top else {}
    return talents + ([(f"<span class=\"wait-pre\">Upgrade: </span>{top['name']}, {top['source'].rsplit(': ', 1)[-1]}", icon(item.get('icon'), 28))] if top else [])


# ------------------------------------------------------------------------------------------------ the states
def elyra() -> Char:
    return Char('character-60.json', FIX / 'bis-priest-holy.json', FIX / 'items-priest.json')


def felsnap(answer: str = 'character-34.json') -> Char:
    return Char(answer, HOME / 'bis-hunter-beast-mastery.json', HOME / 'items-hunter.json')


def O(html: str, order: int, span: bool = False) -> str:
    """A block's place in the phone flow (web ui/Ordered); a block left out leaves no slot."""
    return Ordered(html, order, span) if html else ''


def split(main: str, rail: str) -> str:
    """Page layout workbench with panels (h-audit 2026-11-09): side by side from 1280, one ordered flow below it."""
    return f'<div class="char-pg">{Page("workbench", main, rail, panels=True)}</div>'


def state_ready(ch: Char, me: str | None, premium: bool = False, feed: str | None = None, home: str | None = None,
                coach: str | None = None) -> str:
    owner = J(FIX / me)['characters'][0] if me else None
    body = head(ch, owner)
    if owner:
        body += waiting_summary(waiting(ch, owner))
    logged = bool(ch.page.get('score'))
    latest = (ch.page.get('score_history') or [None])[-1]
    fights = len(ch.page['history'])
    # The owner's History and Coach rows sit under the best fights, in their panel (h-audit 2026-11-09).
    scores = [x['score'] for x in sorted(ch.page.get('score_history') or [], key=lambda x: x['date'])]
    inner = owner_sections(premium, fights, coach, latest, in_panel=True, scores=scores) if owner and logged and ch.page['best'] else ''
    alone = owner_sections(premium, fights, coach, latest) if owner and logged and not ch.page['best'] else ''
    main = (O(gear_section(ch), 1, True) + O(fights_section(ch, inner) if ch.page['best'] else '', 4, True) + O(alone, 5, True))
    stand = standing_section(standing(home, 'guild' in owner, SPEC[ch.cls])) if owner else ''
    rail = (O(score_section(ch, owner is not None), 2) + O(stand, 3) + O(builds_section(ch) if ch.page['builds_seen'] else '', 6)
            + O(activity_section(feed, ch.page['character']['name']), 7))
    return body + split(main, rail)


def state_unread() -> str:
    ch = felsnap('character-34-unread.json')
    owner = J(FIX / 'me-34-unread.json')['characters'][0]
    return (head(ch, owner, answer=False)
            + summary('Your gear has not been read yet. In game, log in, wait for the world to load, then log out once; it appears here within a minute.', tone='neutral')
            + split(O(gear_section(ch, 'Best in slot for level 30 to 39', picks=True), 1, True),
                    O(score_section(ch, True), 2) + O(standing_section(standing('home-34.json', False, 'Beast Mastery')), 3)))


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


def state_loading() -> str:
    # The loading head is the PageHeader (Character.svelte headSkeleton): the name from the URL, a 56 px crest skeleton and a
    # meta skeleton line, the whole head one status.
    head_ = (f'<div role="status" style="display:contents">'
             f'{PageHeader("Elyra Duskvale", crest=Skeleton(56, 56, True), meta=[Skeleton(320, 16)])}</div>')
    # The gear panel's body is one status too (Character.svelte, lane E's shared loading state): role=status, aria-busy, a
    # sr-only "Loading…" and the empty frame (its own group with the Gear h2) inside it.
    frame = PaperdollFrame([], '', '', None, '', 'Gear', heading='h2')
    body = f'<div class="doll-c" role="status" aria-busy="true"><span class="sr-only">Loading&hellip;</span><div class="doll-w">{frame}</div></div>'
    return head_ + split(panel('Gear', body, frame='framed'), '')


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

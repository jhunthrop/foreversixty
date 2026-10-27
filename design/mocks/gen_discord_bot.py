"""Mock boards for the Discord bot lane (design/specs/2026-10-27-discord-bot.md).

  python3 design/mocks/gen_discord_bot.py
    writes discord-bot-settings.html        (1440, guild settings · Discord, every state as cells)
    writes discord-bot-settings-phone.html  (390, the four main settings states stacked)
    writes discord-bot-account.html         (1440, the account page's You panel, every link state)
    writes discord-bot-account-phone.html   (390, unlinked, linked, conflict)
    writes discord-bot-post.html            (1440, the raid post in the desktop and phone clients, seven states)
    writes discord-bot-replies.html         (1440, the ephemeral replies to a button press)
    writes discord-bot-commands.html        (1440, every slash command's answer, desktop and phone)

Built on gen_guild_raids (G1): the same tokens, chips, SIGNUPS and Discord chrome values, imported,
never copied. Boards go through mocklib.write into design/mocks/renders/ (ignored); render.py
screenshots them (names with 'phone' at 390).

Mock data, labelled on every board: G1's Onyxia's Lair raid of Thu 10 Dec, its 29 sign-ups, the
guild Olympus XXVII, its Discord server "Olympus Discord". Real data: the /bis answers read the
active build's BiS file (warrior-fury, level 60, Phase 1 raid, Alliance) and the build's item
table for icon, quality and item level, so the names, sources and DPS deltas are the nightly's.
The /char upgrades' gains, the /report figures and the readiness rows are mock. Discord draws
timestamp markup in each reader's own zone; the boards show a reader in Central time.
"""
import json

from gen_guild_raids import (
    CSS, DOT, GUILD, ONY, SIGNUPS, KILL, WIPE, cell, note, gate_sentence, plural,
)
from mocklib import ROOT, SEAL, CLASS_COLOR, crest, data_uri, page, write

BUILD = json.load(open(ROOT / 'web/src/data/active-build.json'))['build']
BUILD_DIR = ROOT / 'data/builds' / BUILD
SERVER = 'Olympus Discord'
CHANNEL = 'raid-signups'
HOLDER = 'Mograine Raiders'
SIZE = 40
MUTED_HEX = '#9a9484'
GOLD_HEX = '#e5b955'
# The game's quality colours (tokens.css --color-rarity-*), with the word players read.
QUALITY = dict([(0, ('Poor', '#9d9d9d')), (1, ('Common', '#ffffff')), (2, ('Uncommon', '#1eff00')),
                (3, ('Rare', '#0070dd')), (4, ('Epic', '#a335ee')), (5, ('Legendary', '#ff8000'))])
SLOT_WORD = dict(head='Head', neck='Neck', shoulder='Shoulder', back='Back', chest='Chest', wrist='Wrist', hands='Hands',
                 waist='Waist', legs='Legs', feet='Feet', finger1='Ring', finger2='Ring', trinket1='Trinket',
                 trinket2='Trinket', main_hand='Main hand', off_hand='Off hand', ranged='Ranged')
TOP_SLOTS = 5
ALTERNATIVES = 2
ZUL = 'Zulmara'
BIS_FOOT = 'Fury Warrior · level 60 · Phase 1 raid · Forever Sixty'
ITEM_CLASSES = ('rogue', 'druid', 'hunter', 'warrior')
# Brace characters typed as U+2983/U+2984 in the CSS below, swapped at build time.
BRACES = str.maketrans('⦃⦄', '{}')

# Counts, computed once from G1's SIGNUPS (G1 count grammar: roles count In rows only).


def tally(rows: list) -> tuple:
    by_state = dict((s, [r for r in rows if r[4] == s]) for s in ('in', 'late', 'bench', 'out'))
    roles = dict((k, sum(1 for r in by_state['in'] if r[3] == k)) for k in ('Tank', 'Healer', 'DPS'))
    return by_state, roles


STATES, ROLES = tally(SIGNUPS)
N_IN, N_LATE = len(STATES['in']), len(STATES['late'])
N_BENCH, N_OUT = len(STATES['bench']), len(STATES['out'])
N_SIGNED = N_IN + N_LATE


def signed_value() -> str:
    """The Signed up field: line one the G1 count grammar, line two bench and out."""
    if N_SIGNED == 0:
        return 'None yet'
    parts = ['<b>' + f'{N_SIGNED} of {SIZE}' + '</b>', plural(ROLES['Tank'], 'tank'), plural(ROLES['Healer'], 'healer'),
             str(ROLES['DPS']) + ' DPS']
    if N_LATE:
        parts.append(f'{N_LATE} late')
    tail = [p for p in (f'Bench {N_BENCH}' if N_BENCH else '', f'Out {N_OUT}' if N_OUT else '') if p]
    second = DOT.join(tail)
    return DOT.join(parts) + ('<br>' + second if second else '')


# Real BiS data for the /bis answers.


def load_items() -> dict:
    found = dict()
    for cls in ITEM_CLASSES:
        for item in json.load(open(BUILD_DIR / 'items' / f'{cls}.json'))['items']:
            found.setdefault(item['name'], item)
    return found


def load_bis(spec: str) -> list:
    bis = json.load(open(BUILD_DIR / 'bis' / f'{spec}.json'))
    band = next(b for b in bis['bands'] if b['band'] == 60 and b['preset'] == 'raid' and b['faction'] == 'alliance')
    return [s for s in band['slots'] if 'item_id' in s]


BIS_SLOTS = load_bis('warrior-fury')
ITEMS = load_items()


def margin(slot: dict) -> float:
    """How much the pick beats the next best by, in DPS (0 when no alternative is listed)."""
    alts = slot.get('alternatives') or []
    return -alts[0].get('dps_delta', 0) if alts else 0


def top_slots() -> list:
    return sorted(BIS_SLOTS, key=lambda s: -margin(s))[:TOP_SLOTS]


def icon_uri(stem: str) -> str:
    return data_uri(BUILD_DIR / 'icons' / f'{stem}.webp')


def dps(delta: float) -> str:
    return ('%.1f DPS' % delta).replace('-', '−')


# Discord chrome. Values match gen_guild_raids' embed() and message() and the Discord client.

DCSS = '''<style>
.dmsg⦃background:#313338;border-radius:8px;padding:16px;display:flex;gap:16px;font-family:'gg sans','Noto Sans','Helvetica Neue',Arial,sans-serif;color:#dbdee1;font-size:16px;line-height:1.375⦄
.dmsg.ph⦃gap:12px;padding:12px⦄
.dav⦃width:40px;height:40px;border-radius:50%;background:#0d111a;display:grid;place-items:center;flex-shrink:0⦄
.dname⦃font-size:16px;font-weight:500;color:#f2f3f5⦄
.dapp⦃background:#5865f2;color:#fff;font-size:10px;font-weight:600;border-radius:3px;padding:1px 5px;margin:0 6px 0 4px;vertical-align:2px⦄
.dtime⦃font-size:12px;color:#949ba4⦄
.dused⦃display:flex;align-items:center;gap:6px;font-size:14px;color:#949ba4;margin:0 0 4px 56px;font-family:'gg sans','Noto Sans',Arial,sans-serif⦄
.dused i⦃width:16px;height:16px;border-radius:50%;background:#5c6068;display:inline-block⦄
.dused b⦃color:#dbdee1;font-weight:500⦄.dused em⦃color:#00a8fc;font-style:normal⦄
.demb⦃background:#2b2d31;border-left:4px solid;border-radius:4px;padding:8px 16px 16px 12px;display:flex;gap:16px;margin-top:4px⦄
.demb .col⦃min-width:0;flex:1⦄
.dauth⦃font-size:14px;font-weight:600;color:#f2f3f5;margin-top:8px⦄
.dtitle⦃font-size:16px;font-weight:600;color:#00a8fc;margin-top:8px⦄
.ddesc⦃font-size:14px;color:#dbdee1;margin-top:8px⦄
.dflds⦃display:grid;gap:8px 16px;margin-top:8px⦄
.dfn⦃font-size:14px;font-weight:600;color:#f2f3f5⦄.dfv⦃font-size:14px;color:#dbdee1⦄
.dfoot⦃display:flex;align-items:center;gap:8px;font-size:12px;color:#b5bac1;margin-top:8px⦄
.dthumb⦃width:80px;height:80px;border-radius:4px;object-fit:cover;flex-shrink:0;margin-top:8px⦄
.dts⦃background:#ffffff0f;border-radius:3px;padding:0 2px⦄
.dbtns⦃display:flex;flex-wrap:wrap;gap:8px;margin-top:8px⦄
.dbtn⦃display:inline-flex;align-items:center;gap:6px;height:32px;padding:0 16px;border-radius:8px;font-size:14px;font-weight:500;color:#fff;background:#4e5058;white-space:nowrap⦄
.ph .dbtn⦃height:40px;padding:0 18px⦄
.dbtn.suc⦃background:#248046⦄.dbtn.dis⦃opacity:.5⦄
.dsel⦃display:flex;align-items:center;height:40px;padding:0 12px;border-radius:8px;background:#1e1f22;color:#949ba4;font-size:14px;margin-top:8px;max-width:400px⦄
.dsel.open⦃border-radius:8px 8px 0 0⦄
.dsel .chev⦃margin-left:auto⦄
.dopts⦃background:#2b2d31;border:1px solid #1e1f22;border-top:0;border-radius:0 0 8px 8px;max-width:400px⦄
.dopt⦃padding:8px 12px;border-top:1px solid #1e1f22⦄.dopt.on⦃background:#404249⦄
.dopt b⦃display:block;font-size:16px;font-weight:500;color:#dbdee1⦄.dopt span⦃font-size:14px;color:#b5bac1⦄
.deph⦃display:flex;align-items:center;gap:6px;font-size:12px;color:#949ba4;margin-top:4px⦄.deph a⦃color:#00a8fc⦄
.dthink⦃font-size:16px;color:#949ba4;font-style:italic;margin-top:2px⦄
.dtext⦃font-size:16px;color:#dbdee1;margin-top:2px⦄
</style>'''.translate(BRACES)

EYE = ('<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="#949ba4" stroke-width="2">'
       '<path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8S1 12 1 12z"/><circle cx="12" cy="12" r="3"/></svg>')
OUT_ICON = ('<svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="#fff" stroke-width="2">'
            '<path d="M14 4h6v6M20 4l-9 9M18 14v6H4V6h6"/></svg>')
DESK_W, DESK_EMBED = 600, 520
PHONE_W, PHONE_EMBED = 354, 270
NOW = 'Today at 6:12 PM'
STAMP = '<span class="dts">Thursday, December 10, 2026 8:00 PM</span>'
WHEN = STAMP + ' (<span class="dts">in 2 days</span>)'
WHEN_PAST = STAMP + ' (<span class="dts">2 hours ago</span>)'


def btn(label: str, kind: str = 'sec', disabled: bool = False) -> str:
    cls = 'dbtn' + (' suc' if kind == 'suc' else '') + (' dis' if disabled else '')
    tail = OUT_ICON if kind == 'link' else ''
    return f'<span class="{cls}">{label}{tail}</span>'


def buttons(*items: str) -> str:
    return '<div class="dbtns">' + ''.join(items) + '</div>'


def signup_row(disabled: bool = False) -> str:
    """Row 1 of the raid post: In (success), Late, Bench, Out (secondary), then the link button."""
    return buttons(btn('In', 'suc', disabled), btn('Late', disabled=disabled), btn('Bench', disabled=disabled),
                   btn('Out', disabled=disabled), btn('Raid page', 'link'))


def footer_line(label: str) -> str:
    return f'<div class="dfoot"><img src="{SEAL}" width="20" height="20" alt="">{label}{DOT}{NOW}</div>'


def field_cell(name: str, value: str, inline: bool) -> str:
    span = '' if inline else 'grid-column:1/-1;'
    return f'<div style="{span}min-width:0"><div class="dfn">{name}</div><div class="dfv">{value}</div></div>'


def embed(title: str, colour: str, desc: str = '', fields: tuple = (), thumb: str = '', foot: str = '',
          phone: bool = False, author: str = '') -> str:
    cols = 1 if phone else 3
    cells = ''.join(field_cell(n, v, i) for n, v, i in fields)
    grid = 'grid-template-columns:repeat(' + str(cols) + ',minmax(0,1fr))'
    flds = f'<div class="dflds" style="{grid}">{cells}</div>' if cells else ''
    au = f'<div class="dauth">{author}</div>' if author else ''
    ti = f'<div class="dtitle">{title}</div>' if title else ''
    ds = f'<div class="ddesc">{desc}</div>' if desc else ''
    th = f'<img class="dthumb" src="{thumb}" alt="">' if thumb else ''
    ft = footer_line(foot) if foot else ''
    return f'<div class="demb" style="border-color:{colour}"><div class="col">{au}{ti}{ds}{flds}{ft}</div>{th}</div>'


def message(body: str, phone: bool = False, used: tuple = (), ephemeral: bool = False) -> str:
    width = PHONE_W if phone else DESK_W
    av = f'<span class="dav"><img src="{SEAL}" width="28" height="28" alt=""></span>'
    head = f'<div><span class="dname">Forever Sixty</span><span class="dapp">APP</span><span class="dtime">{NOW}</span></div>'
    eph = f'<div class="deph">{EYE}Only you can see this{DOT}<a>Dismiss message</a></div>' if ephemeral else ''
    reply = f'<div class="dused"><i></i><b>{used[0]}</b> used <em>{used[1]}</em></div>' if used else ''
    cls = 'dmsg ph' if phone else 'dmsg'
    return (f'<div style="width:{width}px">{reply}<div class="{cls}">{av}'
            f'<div style="min-width:0;flex:1">{head}{body}{eph}</div></div></div>')


def embeds(html: str, phone: bool) -> str:
    width = PHONE_EMBED if phone else DESK_EMBED
    return f'<div style="max-width:{width}px">{html}</div>'


def pair(builder, title: str) -> str:
    """One cell: the same message in the desktop client (600) and the phone client (354)."""
    return (f'<div class="cell"><span class="cell-t">{title}</span>'
            f'<div style="display:flex;gap:24px;align-items:flex-start">{builder(False)}{builder(True)}</div></div>')


def text(s: str) -> str:
    return f'<div class="dtext">{s}</div>'


# The raid post (spec 3.3.1).

RAID_TITLE = ONY + DOT + 'Thursday 10 December'
POST_FOOT = GUILD + DOT + 'Forever Sixty'


def post_parts(state: str) -> tuple:
    """(title, colour, description, fields, buttons) for one post state."""
    when, gt, signed = ('When', WHEN, True), ('Guild time', '8:00 pm CST', True), ('Signed up', signed_value(), False)
    past = ('When', WHEN_PAST, True)
    moved = ('When', '<span class="dts">Friday, December 11, 2026 8:00 PM</span> (<span class="dts">in 3 days</span>)', True)
    came = ('Came', f'22 of {N_SIGNED}', True)
    parts = dict(
        open=(RAID_TITLE, GOLD_HEX, 'Sign up with the buttons below.', (when, gt, signed), signup_row()),
        locked=(RAID_TITLE, MUTED_HEX, 'Sign-ups are locked. An officer can still change yours.', (when, gt, signed),
                signup_row(True)),
        moved=('Moved: ' + ONY + DOT + 'Friday 11 December', GOLD_HEX, 'Sign-ups carry over. Check you can still make it.',
               (moved, ('Was', STAMP, True), gt, signed), signup_row()),
        cancelled=('Cancelled: ' + RAID_TITLE, WIPE, 'This raid was cancelled.', (('When', '<s>' + STAMP + '</s>', True), signed),
                   buttons(btn('Raid page', 'link'))),
        ended=(RAID_TITLE, MUTED_HEX, 'This raid has ended.', (past, gt, signed), signup_row(True)),
        done=('Done: ' + RAID_TITLE, KILL, 'Logged as a raid night.', (past, came, ('Kills', 'Onyxia 1:42', True)),
              buttons(btn('Open the night', 'link'))),
        webhook=(RAID_TITLE, GOLD_HEX, 'Sign up on Forever Sixty.', (when, gt, signed), ''),
    )
    return parts[state]


def raid_post(state: str, phone: bool) -> str:
    title, colour, desc, fields, row = post_parts(state)
    return message(embeds(embed(title, colour, desc, fields, foot=POST_FOOT, phone=phone), phone) + row, phone)


POST_STATES = [
    ('open', 'Open: planned, buttons live (one post per raid, edited for life)'),
    ('locked', 'Locked: muted bar, the four buttons disabled, the link live'),
    ('moved', 'Moved: the same message edited; Was shows until the raid starts'),
    ('cancelled', 'Cancelled: wipe bar, When struck, only the link stays'),
    ('ended', 'Ended, no log yet: muted bar, buttons disabled'),
    ('done', 'Done: the Done green, Came and Kills, the link opens the night'),
    ('webhook', 'Webhook post (no bot): the same builder, no buttons, counts at posting'),
]


def post_board() -> str:
    cells = ''.join(pair(lambda ph, s=s: raid_post(s, ph), t) for s, t in POST_STATES)
    return cells + note(
        f'Counts from G1 SIGNUPS: {N_IN} In, {N_LATE} Late, {N_BENCH} Bench, {N_OUT} Out. Edits are coalesced to one per 2 s per post; '
        'the last state wins. Thumbnail: the guild crest when absolute (this mock guild has none).')


# Ephemeral replies to a button press (spec 3.3.2).


def reply_line(s: str, link: str = '', used: tuple = ()):
    def build(ph: bool) -> str:
        tail = buttons(btn(link, 'link')) if link else ''
        return message(text(s) + tail, ph, used=used, ephemeral=True)
    return build


SELECT_OPTIONS = [(ZUL, 'Marksmanship Hunter · DPS · your main'), ('Duskstrider', 'Marksmanship Hunter · DPS'),
                  ('Grimvow', 'Shadow Priest · DPS')]


def r_select(ph: bool, open_: bool = False) -> str:
    sel = '<div class="dsel' + (' open' if open_ else '') + '"><span>Choose a character</span><span class="chev">▾</span></div>'
    rows = ''.join('<div class="dopt' + (' on' if i == 0 else '') + f'"><b>{n}</b><span>{d}</span></div>'
                   for i, (n, d) in enumerate(SELECT_OPTIONS))
    lst = f'<div class="dopts">{rows}</div>' if open_ else ''
    lead = text(f'Sign up as <b>In</b> for {ONY}, <span class="dts">Thu, Dec 10, 8:00 PM</span>, with which character?')
    return message(lead + sel + lst, ph, ephemeral=True)


SUCCESS_IN = f'<b>{ZUL}</b> is In as DPS for {ONY}. {N_SIGNED} of {SIZE} signed up.'
SUCCESS_LATE = f'<b>{ZUL}</b> is Late as DPS for {ONY}. Add a note for the officers on the raid page.'

REPLIES = [
    (reply_line('Link your Discord account to Forever Sixty to sign up from here, then press the button again.', 'Link Discord'),
     'Unlinked: what to do, one link, nothing else'),
    (lambda ph: r_select(ph), 'Several verified characters: the select, your main first'),
    (lambda ph: r_select(ph, True), 'The select open (at most 25 options)'),
    (reply_line(SUCCESS_IN), 'Success: In'),
    (reply_line(SUCCESS_LATE, 'Add a note'), 'Success: Late'),
    (reply_line(f'<b>{ZUL}</b> is In as DPS for {ONY}. Duskstrider moved to Bench.'), 'Success: the other character moved to Bench (R3)'),
    (reply_line(f'<b>{ZUL}</b> is already In as DPS for {ONY}.'), 'The same state pressed again: no toggle'),
    (reply_line('Sign-ups are locked. An officer can still change yours.'), 'Refused: locked'),
    (reply_line('This raid was cancelled.'), 'Refused: cancelled'),
    (reply_line('This raid has ended.'), 'Refused: ended'),
    (reply_line('Duskstrider is waiting for approval. You can sign it up once an officer approves it.'),
     'Refused: the only character is pending'),
    (reply_line(f'Sign-ups are for members of {GUILD}. Save your character on Forever Sixty and an officer can approve it.',
                'Get set up'), 'Refused: no character in this guild'),
    (reply_line('Too many sign-up changes. Wait a minute, then try again.'), 'Refused: 30 a minute per account'),
    (reply_line('That did not save. Try again.'), 'Error: our side'),
    (lambda ph: message('<div class="dthink">Forever Sixty is thinking…</div>', ph, ephemeral=True),
     'Deferred: thinking (an answer that needs more than 3 s)'),
]


def replies_board() -> str:
    return ''.join(pair(b, t) for b, t in REPLIES) + note(
        'Every reply is ephemeral (flag 64). A select choice replaces the same ephemeral message; the raid post updates within 2 s.')


# Slash command answers (spec 3.3.3).

SHARE = 'Share in channel'


def item_embed(name: str, slot: str, source: str, phone: bool, lead: bool = False, delta: str = '', foot: str = '') -> str:
    """One item: thumbnail its own icon, bar its quality colour, the quality word in the text."""
    it = ITEMS.get(name, dict())
    qword, qcol = QUALITY[it.get('quality', 1)]
    level = 'item level ' + str(it['item_level']) if it else ''
    head = DOT.join(p for p in (SLOT_WORD.get(slot, ''), qword, level) if p)
    second = DOT.join(p for p in (delta, source) if p)
    thumb = icon_uri(it['icon']) if it.get('icon') else ''
    return embed(name, qcol, head + '<br>' + second, thumb=thumb, phone=phone,
                 author='Best in slot' if lead else '', foot=foot)


def last_foot(i: int, count: int, label: str) -> str:
    return label if i == count - 1 else ''


def a_bis_slot(ph: bool) -> str:
    pick = next(s for s in BIS_SLOTS if s['slot'] == 'head')
    alts = (pick.get('alternatives') or [])[:ALTERNATIVES]
    html = item_embed(pick['item_name'], 'head', pick['source'], ph, lead=True)
    html += ''.join(item_embed(a['item_name'], 'head', a['source'], ph, delta=dps(a.get('dps_delta', 0)),
                               foot=last_foot(i, len(alts), BIS_FOOT)) for i, a in enumerate(alts))
    return message(text('Head for a Fury Warrior, level 60 Phase 1 raid:') + embeds(html, ph)
                   + buttons(btn(SHARE), btn('Open on Forever Sixty', 'link')), ph,
                   used=('Kraggor', '/bis class:Warrior spec:Fury slot:Head'), ephemeral=True)


TOP_LEAD = 'Fury Warrior, level 60 Phase 1 raid. The five slots where the best item gains the most over the next best:'


def top_embeds(ph: bool, count: int = TOP_SLOTS) -> str:
    tops = top_slots()[:count]
    return ''.join(item_embed(s['item_name'], s['slot'], s['source'], ph, foot=last_foot(i, len(tops), BIS_FOOT))
                   for i, s in enumerate(tops))


def a_bis_top(ph: bool) -> str:
    return message(text(TOP_LEAD) + embeds(top_embeds(ph), ph) + buttons(btn(SHARE), btn('All 17 slots', 'link')), ph,
                   used=('Kraggor', '/bis class:Warrior spec:Fury'), ephemeral=True)


def a_shared(ph: bool) -> str:
    return message(text(TOP_LEAD) + embeds(top_embeds(ph, 2), ph) + buttons(btn('All 17 slots', 'link')), ph,
                   used=('Kraggor', '/bis class:Warrior spec:Fury'))


UPGRADES = [('legs', "Sentinel's Chain Leggings", '+9.2 DPS', 'Illiyana Moonblaze'),
            ('chest', 'Dawn Armor', '+6.6 DPS', 'Leatherworking'),
            ('finger1', 'Signet Ring of the Bronze Dragonflight', '+5.5 DPS', 'Anachronos')]


def a_char(ph: bool) -> str:
    card = embed(ZUL, CLASS_COLOR['hunter'], 'Level 60 Marksmanship Hunter' + DOT + GUILD,
                 fields=(('Item level', '61.4', True), ('Readiness', 'Short: no flask', True), ('Synced', '2 days ago', True),
                         ('Top 3 upgrades', 'By DPS gained, against the guide build.', False)),
                 thumb=crest('hunter'), phone=ph)
    ups = ''.join(item_embed(n, s, src, ph, delta=g, foot=last_foot(i, len(UPGRADES), 'Forever Sixty'))
                  for i, (s, n, g, src) in enumerate(UPGRADES))
    return message(embeds(card + ups, ph) + buttons(btn(SHARE), btn('Character page', 'link')), ph,
                   used=(ZUL, '/char name:Zulmara'), ephemeral=True)


def a_raid(ph: bool) -> str:
    e = embed(RAID_TITLE, GOLD_HEX, f'You are <b>In</b> on {ZUL} as DPS.',
              fields=(('When', WHEN, True), ('Guild time', '8:00 pm CST', True), ('Signed up', signed_value(), False)),
              foot=POST_FOOT, phone=ph)
    return message(embeds(e, ph) + signup_row(), ph, used=(ZUL, '/raid'), ephemeral=True)


LOOT = [("Vis'kag the Bloodletter", 'Vexlash'), ('Head of Onyxia', 'Kraggor'), ('Ring of Binding', 'Thornhide')]
NIGHT = f'22 of {N_SIGNED} came' + DOT + '1 kill' + DOT + '3 wipes' + DOT + '14 deaths' + DOT + '2 h 41 min'


def a_report(ph: bool, officer: bool = False) -> str:
    fields = [('Kill order', '1. Onyxia <b>1:42</b> after 3 wipes', False)]
    if officer:
        fields += [('Most deaths', 'Pyrewisp 4 · Windtalon 3 · Vexlash 2', False),
                   ('Consumables', '3 short before the raid · 2 with no flask in the log', False)]
    fields += [('Loot', '6 items. The first 3:', False)]
    main = embed('Report: ' + RAID_TITLE, KILL, NIGHT, fields=tuple(fields), phone=ph)
    loot = ''.join(item_embed(n, '', 'to ' + who + DOT + 'Onyxia', ph, foot=last_foot(i, len(LOOT), POST_FOOT))
                   for i, (n, who) in enumerate(LOOT))
    return message(embeds(main + loot, ph) + buttons(btn(SHARE), btn('Open the report', 'link')), ph,
                   used=('Kraggor' if officer else ZUL, '/report'), ephemeral=True)


LAST_FIVE = 'Thu 10 Dec Present<br>Sat 5 Dec Present<br>Thu 3 Dec Late, excused<br>Sat 28 Nov Not in the log<br>Thu 26 Nov Present'


def a_attendance(ph: bool, officer: bool = False) -> str:
    who, cls = ('Pyrewisp', 'mage') if officer else (ZUL, 'hunter')
    e = embed(who + DOT + 'attendance', CLASS_COLOR[cls], 'Came <b>7 of 8</b>' + DOT + '<b>26 of 30</b> raids',
              fields=(('Last 5 raids', LAST_FIVE, False),), foot=POST_FOOT, phone=ph)
    cmd = '/attendance name:Pyrewisp' if officer else '/attendance'
    return message(embeds(e, ph) + buttons(btn('Roster', 'link')), ph, used=('Kraggor' if officer else ZUL, cmd), ephemeral=True)


READY_ROWS = ('<b>Pyrewisp</b> · Short · Synced 9 days before the raid<br><b>Windtalon</b> · Short<br>'
              '<b>Felsnap</b> · No export yet<br><b>Hexbramble</b> · Not shared<br><b>Vexlash</b> · Gear gap 12%')


def a_readiness(ph: bool) -> str:
    e = embed('Readiness' + DOT + RAID_TITLE, GOLD_HEX, f'The {N_SIGNED} signed In or Late. 5 have something to fix.',
              fields=(('To fix, most checks first', READY_ROWS, False),), foot=POST_FOOT, phone=ph)
    return message(embeds(e, ph) + buttons(btn('Readiness board', 'link')), ph, used=('Kraggor', '/readiness'), ephemeral=True)


HELP = [('/raid', "Your guild's next raid. Sign up from the answer."), ('/signup', 'Sign up for a raid without the post.'),
        ('/char', 'A character: item level, readiness and the top 3 upgrades.'), ('/bis', 'Best in slot for a class and spec.'),
        ('/report', 'The last logged raid: kills and loot.'), ('/attendance', 'Your attendance. Officers can name anyone.'),
        ('/readiness', 'Officers: who is not ready for the next raid.'), ('/link', 'Link this Discord account to Forever Sixty.'),
        ('/help', 'This list.')]


def a_help(ph: bool) -> str:
    lines = '<br>'.join(f'<b>{c}</b> {d}' for c, d in HELP)
    return message(embeds(embed('Forever Sixty commands', GOLD_HEX, lines, foot='Forever Sixty', phone=ph), ph), ph,
                   used=(ZUL, '/help'), ephemeral=True)


COMMANDS = [
    (a_raid, '/raid: the next raid, your status and the buttons (no Share: a visitor sees no planned raid)'),
    (reply_line(SUCCESS_LATE, 'Add a note', used=(ZUL, '/signup status:Late')), "/signup: the button's own reply"),
    (a_char, '/char: the card (class colour, ringed crest), then three item embeds (icon, quality bar)'),
    (a_bis_slot, '/bis with a slot: the pick and two alternatives, each with its DPS delta'),
    (a_bis_top, '/bis without a slot: five item embeds'),
    (a_shared, 'Share in channel: the public repost (cut to two items on this board)'),
    (lambda ph: a_report(ph), '/report, member: night line, kill order, loot'),
    (lambda ph: a_report(ph, True), '/report, officer: adds deaths and consumables (a Share leaves them out)'),
    (lambda ph: a_attendance(ph), '/attendance: your record'),
    (lambda ph: a_attendance(ph, True), '/attendance name: (officers only)'),
    (a_readiness, '/readiness (officers): failing rows first'),
    (reply_line('Link this Discord account to Forever Sixty so sign-up buttons and commands know your characters.',
                'Link Discord', used=(ZUL, '/link')), '/link, unlinked'),
    (reply_line('This Discord account is linked to Forever Sixty. Unlink it on your account page.', 'Account page',
                used=(ZUL, '/link')), '/link, linked'),
    (a_help, '/help'),
    (reply_line('No public character by that name.'), '/char: private or unknown'),
    (reply_line('No raid logged yet. A raid shows here once its log is matched.'), '/report: none'),
    (reply_line(f'Readiness is for officers of {GUILD}.'), '/readiness, a member'),
    (reply_line('Planning and running raids are part of the guild plan. An officer can subscribe the guild.', 'Subscribe the guild'),
     '/readiness, plan lapsed'),
    (reply_line("This guild's claim is contested. Officer actions are frozen until a moderator resolves it."), '/readiness, claim frozen'),
    (reply_line("Only officers can look up another member's attendance."), '/attendance name:, a member'),
    (reply_line("Link your Discord account to see your guild's next raid.", 'Link Discord'), '/raid, unlinked'),
    (reply_line('No raid planned yet.'), '/raid: none planned'),
    (reply_line('Name a class and spec, or link Discord to use your character.', 'Link Discord'), '/bis with no options, unlinked'),
    (reply_line("This server is not linked to a Forever Sixty guild. An officer can add the bot from the guild's settings."),
     'Any guild command in an unbound server'),
    (reply_line('Could not load that. Try again.'), 'Any command: our error'),
]


def commands_board() -> str:
    return ''.join(pair(b, t) for b, t in COMMANDS) + note(
        f"BiS rows: build {BUILD}, warrior-fury level 60 Phase 1 raid, Alliance (the guild's faction). "
        "Item embeds: thumbnail the item's own icon, bar its quality colour, the quality word in the text. "
        '/char gains, /report figures and readiness rows are mock.')


# The site: guild settings · Discord (spec 3.1).

SITE_CSS = '''<style>
.sec⦃display:flex;flex-direction:column;gap:12px;max-width:720px⦄
.p13⦃margin:0;font-size:13px;color:#9a9484⦄.p14⦃margin:0;font-size:14px⦄
.st⦃min-height:21px;font-size:13px;margin:0⦄
.lb⦃display:inline-flex;align-items:center;min-height:44px;font-size:13px;color:#e9e4d8;text-decoration:underline;text-underline-offset:3px⦄
.strong⦃color:#f2eee4;font-weight:600⦄
.row⦃display:flex;flex-wrap:wrap;gap:12px;align-items:center⦄
.notice⦃border-left:2px solid;padding-left:12px;display:flex;flex-direction:column;gap:2px⦄
.sel⦃display:flex;align-items:center;height:44px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;background:#0d111a;color:#e9e4d8;font-size:14px;max-width:360px⦄
.sel.ph⦃color:#9a9484⦄.sel .chev⦃margin-left:auto;color:#9a9484;font-size:11px⦄
.opts⦃border:1px solid #3a3326;border-top:0;border-radius:0 0 4px 4px;background:#0d111a;max-width:360px⦄
.opts div⦃padding:0 12px;min-height:44px;display:flex;align-items:center;font-size:14px⦄
.opts .grp⦃font-size:11px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;color:#9a9484;min-height:28px⦄
.opts .on⦃background:#131824⦄
.focus⦃outline:2px solid #e5b955;outline-offset:2px⦄
.stack⦃display:flex;flex-direction:column;align-items:flex-start;gap:8px⦄
</style>'''.translate(BRACES)


def h2(label: str = 'Discord') -> str:
    return f'<span class="h2">{label}</span>'


def status(msg: str = '', bad: bool = False) -> str:
    cls = 'st wipe' if bad else 'st'
    return f'<p class="{cls}">{msg}</p>'


GATE = gate_sentence().replace('#settings', '#billing')
BOT_EXPLAIN = ('Raid posts get In, Late, Bench and Out buttons, and members can ask /raid, /bis and /char in the server. '
               'Discord asks which server; you need Manage Server in it.')
WEBHOOK_NOTE = '<p class="p13">Raids post there, and members sign up on Forever Sixty.</p>'


def s_none(msg: str = '', bad: bool = False, focus: bool = False, gated: bool = False) -> str:
    off = ' off' if gated else ''
    btn_cls = 'btnp' + (' focus' if focus else '') + off
    gate = GATE if gated else ''
    line = status(msg, bad)
    alt = '<div><span class="lb' + off + '">Connect a channel without the bot</span>' + WEBHOOK_NOTE + '</div>'
    return (f'<div class="sec">{h2()}{gate}<div class="stack"><span class="{btn_cls}">Add the Forever Sixty bot</span>'
            f'<p class="p13">{BOT_EXPLAIN}</p>{line}</div>{alt}</div>')


def s_webhook_offer() -> str:
    return (f'<div class="sec">{h2()}<div><p class="p14"><span style="color:#9a9484">Posts raids to</span> '
            '<span class="strong">a Discord channel</span></p>'
            '<p class="p13">Channel <span class="mono" style="color:#e9e4d8">1198…4568</span> · '
            '<span class="lb" style="min-height:0">Open in Discord</span></p></div>'
            '<div class="stack"><p class="p14">Add the bot and raid posts get In, Late, Bench and Out buttons. '
            'This channel stays as the fallback.</p><span class="btns">Add the Forever Sixty bot</span></div>'
            '<div class="row"><span class="btns">Send test message</span><span class="btns">Remove</span></div>' + status()
            + '<div><span class="lb">Connect a different channel</span>'
            '<p class="p13">Connecting again replaces this channel and removes its webhook from Discord.</p></div></div>')


PICKER_HEAD = f'<p class="p14"><span style="color:#9a9484">Bot added to</span> <span class="strong">{SERVER}</span></p>'
CH_LABEL = '<span class="label">Raid channel</span>'
CH_HELP = '<p class="p13">Text channels the bot can post in.</p>'
REMOVE_LINK = '<span class="lb">Remove the bot</span>'


def picker_shell(body: str, actions: str, msg: str = '', bad: bool = False) -> str:
    line = status(msg, bad)
    return (f'<div class="sec">{h2()}{PICKER_HEAD}<div style="display:flex;flex-direction:column;gap:8px">{body}</div>'
            f'{actions}{line}</div>')


def picker_actions(save_off: bool = False) -> str:
    cls = 'btnp off' if save_off else 'btnp'
    return f'<div class="row"><span class="{cls}">Save channel</span>{REMOVE_LINK}</div>'


def select_box(label: str, placeholder: bool = False, focus: bool = False) -> str:
    cls = 'sel' + (' ph' if placeholder else '') + (' focus' if focus else '')
    return f'<div class="{cls}"><span>{label}</span><span class="chev">▾</span></div>'


OPEN_LIST = ('<div class="opts"><div class="grp">Raids</div><div class="on">#raid-signups</div><div>#raid-chat</div>'
             '<div class="grp">General</div><div>#announcements</div><div>#general</div></div>')
EMPTY_LINE = (f'<p class="p14">The bot cannot post in any text channel of {SERVER}. In Discord, give the Forever Sixty role '
              'View Channel, Send Messages and Embed Links in your raid channel, then reload the list.</p>')
RACE_LINE = 'That channel is gone, or the bot cannot post there. Choose another; the list is fresh.'


def s_picker(state: str = 'ready') -> str:
    if state == 'loading':
        return picker_shell(CH_LABEL + select_box('Loading channels…', True) + CH_HELP, picker_actions(True))
    if state == 'empty':
        return picker_shell(EMPTY_LINE, '<div class="row"><span class="btns">Reload the list</span>' + REMOVE_LINK + '</div>')
    if state == 'error':
        body = CH_LABEL + '<p class="st wipe">Could not load the channels.</p>'
        return picker_shell(body, '<div class="row"><span class="btns">Try again</span>' + REMOVE_LINK + '</div>')
    if state == 'race':
        return picker_shell(CH_LABEL + select_box('Choose a channel', True) + CH_HELP, picker_actions(), RACE_LINE, True)
    if state == 'open':
        return picker_shell(CH_LABEL + select_box('Choose a channel', True, True) + OPEN_LIST + CH_HELP, picker_actions())
    return picker_shell(CH_LABEL + select_box('#' + CHANNEL, focus=True) + CH_HELP, picker_actions(),
                        'Bot added. Choose the channel it posts raids in.')


DEST = (f'<div><p class="p14"><span style="color:#9a9484">Posts raids to</span> <span class="strong">#{CHANNEL}</span> '
        f'<span style="color:#9a9484">in</span> <span class="strong">{SERVER}</span></p>'
        '<p class="p13">With sign-up buttons · <span class="lb" style="min-height:0">Open in Discord</span></p></div>')
LINKED_LINE = '<p class="p13">9 of 31 members have linked Discord. The others get a link when they press a button.</p>'
WARNING = ('<div><p class="p14" style="font-size:13px">The last 3 posts failed. At 5 in a row it switches off. '
           'Send a test message to check.</p><p class="p13">Last error from Discord: Missing Permissions.</p></div>')
FALLBACK = '<p class="p13">Your webhook channel stays as the fallback if the bot is removed.</p>'
STILL = '<p class="p13">Members can still sign up with the buttons on posted raids and use the commands.</p>'


def remove_question(webhook: bool) -> str:
    tail = 'Raids post through your webhook channel again.' if webhook else 'Raids stop posting to Discord.'
    return (f'<div><p class="p14" style="margin-bottom:8px">Remove the bot from {SERVER}? Raid posts lose their buttons '
            f'and the commands stop. {tail}</p><div class="row"><span class="btns">Remove</span>'
            '<span class="btns focus">Cancel</span></div></div>')


def installed_actions(gated: bool, frozen: bool) -> str:
    off = ' off' if (gated or frozen) else ''
    rm_off = ' off' if frozen else ''
    return (f'<div class="row"><span class="btns{off}">Send test message</span><span class="btns{off}">Change channel</span>'
            f'<span class="btns{rm_off}">Remove</span></div>')


def s_installed(warning: bool = False, confirm: bool = False, gated: bool = False, frozen: bool = False,
                webhook: bool = True, msg: str = '') -> str:
    actions = remove_question(webhook) if confirm else installed_actions(gated, frozen)
    parts = [h2(), GATE if gated else '', DEST, WARNING if warning else '', LINKED_LINE, STILL if gated else '',
             actions, status(msg), FALLBACK if webhook and not confirm else '']
    return '<div class="sec">' + ''.join(parts) + '</div>'


def s_kicked(webhook: bool = True) -> str:
    line = 'Raids post through your webhook channel again.' if webhook else 'Raids stop posting to Discord until the bot is added again.'
    return (f'<div class="sec">{h2()}<div class="notice" style="border-color:#e5b955">'
            f'<p class="p14 strong">The bot is no longer in {SERVER}.</p><p class="p13">It was removed in Discord. {line}</p></div>'
            '<div class="stack"><span class="btnp">Add the Forever Sixty bot</span></div>' + status() + '</div>')


TAKEN = (f'That Discord server is linked to {HOLDER}. One server works with one guild: ask their officers to remove the bot '
         'from their guild settings, or add it to another server.')
FAILED = 'Could not add the bot. You need Manage Server in that Discord server and to be an officer of this guild.'

SETTINGS_STATES = [
    ('Nothing connected: the bot is the one fill, the webhook one link away', s_none()),
    ('Nothing connected, officer-read-only', s_none(gated=True)),
    ('Webhook connected, no bot: the next step under the destination', s_webhook_offer()),
    ('Return: bot added, choose the channel (Raid channel focused)', s_picker()),
    ('Channel select open (grouped by category)', s_picker('open')),
    ('Channels loading', s_picker('loading')),
    ('No channel the bot can post in', s_picker('empty')),
    ('Channels did not load', s_picker('error')),
    ('Race: the channel was deleted before Save', s_picker('race')),
    ('Installed (webhook kept as the fallback), just saved', s_installed(msg='Saved. Raids post here with sign-up buttons.')),
    ('Installed with the 1 to 4 failure warning', s_installed(warning=True)),
    ('Remove question (focus on Cancel)', s_installed(confirm=True)),
    ('Refused: the server is linked to another guild (focus on the add button)', s_none(TAKEN, True, focus=True)),
    ('Plan lapsed: read-only, Remove live', s_installed(gated=True)),
    ('Claim frozen: every control disabled', s_installed(frozen=True)),
    ('The bot was removed in Discord', s_kicked()),
    ('Return: cancelled in Discord', s_none('Cancelled in Discord. Nothing was added.')),
    ('Return: failed', s_none(FAILED, True)),
]
PHONE_SETTINGS = (0, 3, 9, 12)


def settings_board(phone: bool = False) -> str:
    states = [SETTINGS_STATES[i] for i in PHONE_SETTINGS] if phone else SETTINGS_STATES
    width = 354 if phone else 754
    return ''.join(cell(t, b, width=width) for t, b in states)


def billing_then_discord() -> str:
    billing = '<div class="sec">' + h2('Billing') + '<p class="p14">Guild plan · renews 24 November 2026</p></div>'
    return cell('Section order: Billing, then Discord', billing + '<div style="height:32px"></div>' + s_none(), width=754)


# The site: account page, You panel (spec 3.2).

ANON = ('<div><label style="display:flex;align-items:center;gap:12px;min-height:44px;font-size:14px"><span class="cb no"></span>'
        'Show me under a pseudonym on public pages</label><p class="p13">(The existing pseudonym row, unchanged.)</p></div>')
PANEL_HEAD = ('<div style="display:flex;align-items:center;gap:8px"><span style="width:6px;height:6px;border-radius:50%;'
              'background:#e5b955"></span><span class="label" style="color:#e5b955">You</span></div>')
LINK_LEAD = ("<p class=\"p13\">Link your Discord account so your guild's sign-up buttons and Forever Sixty commands "
             'know your characters.</p>')
SUGGEST_LEAD = f'<p class="p14">{GUILD} signs up in Discord. Link your Discord account to use the buttons.</p>'


def you_panel(body: str, phone: bool = False) -> str:
    pad = '16px' if phone else '20px 24px'
    style = f'background:#0d111a;border:1px solid #262e40;border-radius:6px;padding:{pad};display:flex;flex-direction:column;gap:16px'
    return f'<div style="{style}">{PANEL_HEAD}{ANON}{body}</div>'


def l_unlinked(suggest: bool = False, msg: str = '', bad: bool = False, sub: str = '', busy: bool = False) -> str:
    lead = SUGGEST_LEAD if suggest else LINK_LEAD
    cls = 'btnp' if suggest else 'btns'
    label = 'Opening Discord…' if busy else 'Link Discord'
    extra = f'<p class="p13">{sub}</p>' if sub else ''
    line = status(msg, bad)
    return f'<div class="stack">{lead}<span class="{cls}">{label}</span><div>{line}{extra}</div></div>'


LINKED_HEAD = '<p class="p14"><span style="color:#9a9484">Linked to Discord as</span> <span class="strong">zulmara</span></p>'
UNLINK_Q = ('<div><p class="p14" style="margin-bottom:8px">Unlink Discord? Sign-up buttons and commands will ask you to link again.</p>'
            '<div class="row"><span class="btns">Unlink</span><span class="btns focus">Cancel</span></div></div>')


def l_linked(confirm: bool = False, msg: str = '') -> str:
    act = UNLINK_Q if confirm else '<span class="btns">Unlink</span>'
    return f'<div class="stack">{LINKED_HEAD}{act}{status(msg)}</div>'


CONFLICT = 'That Discord account is linked to another Forever Sixty account.'
CONFLICT_HELP = 'Sign in to that account and unlink it on its account page, then try again.'

ACCOUNT_STATES = [
    ('Unlinked (secondary)', l_unlinked()),
    ("Unlinked, a guild of yours has the bot (the panel's one fill)", l_unlinked(True)),
    ('Opening Discord…', l_unlinked(busy=True)),
    ('Return: linked (focus on Unlink)', l_linked(msg='Linked. In Discord, sign-up buttons and commands now know your characters.')),
    ('Linked at rest', l_linked()),
    ('Unlink question (focus on Cancel)', l_linked(True)),
    ('Unlinked (4 s), focus on Link Discord', l_unlinked(msg='Unlinked.')),
    ('Conflict: linked to another account', l_unlinked(msg=CONFLICT, bad=True, sub=CONFLICT_HELP)),
    ('Return: cancelled in Discord', l_unlinked(msg='Cancelled in Discord. Nothing was linked.')),
    ('Return: failed', l_unlinked(msg='Could not link. Try again.', bad=True)),
    ('Stuck: Discord did not open in 10 s', l_unlinked(msg='Discord did not open. Try again.')),
]
PHONE_ACCOUNT = (0, 4, 7)


def account_board(phone: bool = False) -> str:
    states = [ACCOUNT_STATES[i] for i in PHONE_ACCOUNT] if phone else ACCOUNT_STATES
    width = 354 if phone else 429
    return ''.join(cell(t, you_panel(b, phone), width=width) for t, b in states)


def caption() -> str:
    return note(f'Mock guild {GUILD}, server {SERVER}, roster from G1. BiS rows from build {BUILD}. Not real Discord captures: '
                'the captures from the private test server replace the Discord boards at acceptance.')


def sheet(name: str, title: str, body: str, width: int) -> str:
    gutter = 18 if width < 500 else 48
    content = (f'<div style="display:flex;flex-direction:column;gap:24px;padding:32px {gutter}px">'
               f'<h1 class="display" style="margin:0;font-size:22px;color:#f2eee4">{title}</h1>'
               f'<div style="display:flex;flex-wrap:wrap;gap:24px">{body}</div>{caption()}</div>')
    return write(name + '.html', page(CSS + SITE_CSS + DCSS + content, width)).as_posix()


BOARDS = [
    ('discord-bot-settings', 'Guild settings · Discord, with the bot', lambda: billing_then_discord() + settings_board(), 1440),
    ('discord-bot-settings-phone', 'Settings · Discord', lambda: settings_board(True), 390),
    ('discord-bot-account', 'Account · You panel, Discord link', account_board, 1440),
    ('discord-bot-account-phone', 'Account · Discord link', lambda: account_board(True), 390),
    ('discord-bot-post', 'Discord: the raid post', post_board, 1440),
    ('discord-bot-replies', 'Discord: replies to a button press', replies_board, 1440),
    ('discord-bot-commands', 'Discord: slash command answers', commands_board, 1440),
]

if __name__ == '__main__':
    for board_name, board_title, build, board_width in BOARDS:
        print(sheet(board_name, board_title, build(), board_width))

"""Mock boards for the launch pricing lane (design/specs/2026-10-27-launch-pricing.md).

  python3 design/mocks/gen_launch_pricing.py
    writes launch-pricing.html                  (1440, the plans page, founder open, Yearly)
    writes launch-pricing-2000.html             (2000, the same page, proving the centred 1344 column)
    writes launch-pricing-phone.html            (390, the plans page: Free summary row, then Premium)
    writes launch-pricing-states.html           (1440, Monthly, before and after the window, checkout chooser)
    writes launch-pricing-billing.html          (1440, guild settings #billing: Raid guild, seats, who pays, records)
    writes launch-pricing-billing-phone.html    (390, the same section stacked)
    writes launch-pricing-billing-states.html   (1440, every other billing state as labelled cells)
    writes launch-pricing-account.html          (1440, the account Plan panel: offer, accepted, seat, founder, ended)
    writes launch-pricing-account-phone.html    (390, the seat offer)
    writes launch-pricing-emails.html           (1440, the five emails: subject, first sentence, link)
    writes premium-value-1024/.html/-2000.html  (1024, 1440, 2000: the plans page as B3.5, Region A and Region B)

Built on gen_guild_raids (G1): its CSS, pills, crests, guild band and state cells are imported, never
copied. Boards go through mocklib.write into design/mocks/renders/; render.py screenshots them
(names with 'phone' at 390, names with '2000' at 2000, the rest at 1440).

Mock data, labelled on every board: the guild Olympus XXVII and G1's roster names; "now" is
Tue 8 Dec 2026, inside the founder window. Every price comes from PRICES below (the pricing
design's section 3) and every money string is computed from it, never typed.
"""
from datetime import date

from gen_guild_raids import CSS, DOT, GUILD, KILL, TEXT_CLASS, cell, crest_img, guild_band, pill, sk
from mocklib import footer, nav, nav_phone, page, write

NOW = date(2026, 12, 8)
FOUNDER_OPENS = date(2026, 11, 4)
FOUNDER_CLOSES = date(2027, 1, 6)
TRIAL_DAYS = 7
GUILD_CAP = 25
VERIFIED = 28
SEAT_LOCK_DAYS = 30

PRICES = dict(
    premium=dict(name='Premium', monthly=6, yearly=50, founder_monthly=5, founder_yearly=40, seats=0),
    guild=dict(name='Guild', monthly=15, yearly=150, founder_monthly=12, founder_yearly=120, seats=1),
    raid_guild=dict(name='Raid guild', monthly=25, yearly=250, founder_monthly=20, founder_yearly=200, seats=3),
)
WHO = dict(free='For every player', premium='For one player', guild='For a guild of up to 25',
           raid_guild='For a raiding guild of any size')
FEATURES = dict(
    free=['Addon export and import', 'BiS lists for every spec and level band', 'Talent planner', 'Browser sims',
          'Logs, live logging and parses', 'Rankings and the tier list', 'Class guides'],
    premium=['Server sims: 5,000 combinations, any precision, no tab left open', 'History charts of your score and parses',
             'Discord notifications when a server sim or a log finishes', 'A supporter mark on your characters'],
    guild=['Up to 25 verified members', 'Raids and sign-ups, on the site and in Discord',
           'Attendance and the raid-night report', "The guild's Discord channel and bot", '1 Premium seat to give to a member'],
    raid_guild=['Everything in Guild', 'Any number of members', 'Loot council', 'The readiness board, with consumables',
                '3 Premium seats to give to members'],
)
PROMISE = ('Everything a player needs is free, and the free list never shrinks. Premium and the guild plans pay for the '
           'site. There are no ads for anyone.')
FOUNDER_OPEN = 'Founder prices until 6 January. Buy by then and the price stays for as long as you stay subscribed.'
FOUNDER_BEFORE = 'Founder prices from 4 November until 6 January.'
FOUNDER_NOTE = 'Founder price, kept while you stay subscribed.'
TRIAL = 'First 7 days free, once per guild. Cancel before day 8 and nothing is charged.'
HISTORY = ('If a guild plan ends, its raids, attendance, loot and reports stay readable for good, and officers can '
           'export attendance and loot as CSV at any time.')
HOW = [
    ('Size', 'Guild counts verified members each day, the leader included. Above 25 the officer is told before renewal '
             'and offered Raid guild. Nothing switches off mid-plan.'),
    ('Seats', 'The subscriber or the guild leader gives each seat to any verified member. A seat holder has Premium '
              'while the plan lasts.'),
    ('Moving a seat', 'A seat can move once every 30 days. The old holder keeps Premium until it moves. If a holder '
                      'leaves the guild, the seat comes back at once.'),
    ('Already on Premium', 'A member who pays for Premium accepts a seat first. Their own Premium stops renewing and the '
                           'seat starts when it ends, so nobody pays twice. If the seat is taken away later, they get a '
                           'link to come back at the price they had.'),
    ('Trial', "A guild's first 7 days are free, once per guild. The card is added at the start; cancel before day 8 and "
              'nothing is charged.'),
    ('Renewals', 'We email 7 days before every yearly renewal. Cancel in one click; the plan runs to the end of the '
                 'period you paid for.'),
    ('A new guild leader', 'The subscriber hands the plan to another officer from guild settings. The founder price goes '
                           'with it.'),
    ('Refunds', 'Ask within 14 days of a first charge for a refund. After that there are no part-year refunds.'),
]
FOOT = 'Prices in US dollars. Cancel any time; you keep the plan to the end of the period you paid for.'
MOCK_NOTE = 'Mock guild Olympus XXVII and G1 roster names. Now: Tue 8 Dec 2026, inside the founder window.'

LP_CSS = '''<style>
.lp-col{width:100%;max-width:1344px;margin:0 auto}
.lp-card{background:#0d111a;border:1px solid #262e40;border-radius:6px;box-shadow:inset 0 1px 0 #e5b95559;padding:20px;display:flex;flex-direction:column}
.lp-name{font-family:var(--font-display);font-size:17px;font-weight:700;color:#f2eee4;line-height:24px}
.lp-who{font-size:13px;color:#9a9484;margin-top:2px}
.lp-price{display:flex;flex-direction:column;gap:2px;margin-top:14px}
.lp-fig{font-family:var(--font-mono);font-variant-numeric:tabular-nums;font-size:32px;color:#f2eee4;line-height:38px}
.lp-unit{font-size:15px;color:#9a9484;margin-left:2px}
.lp-was{font-family:var(--font-mono);font-size:15px;color:#9a9484;margin-left:8px}
.lp-sub{font-size:13px;color:#9a9484}
.lp-fnote{font-size:12px;color:#e5b955}
.lp-btn{display:flex;align-items:center;justify-content:center;height:44px;border:1px solid #e5b955;border-radius:4px;color:#e5b955;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap}
.lp-btn.hover{border-color:#f5d27a;color:#f5d27a;background:#e5b95514}
.lp-btn.focus{box-shadow:0 0 0 2px #07090d,0 0 0 4px #e5b955}
.lp-trial{font-size:13px;color:#e9e4d8;line-height:18px}
.lp-feat{list-style:none;margin:0;padding:0}
.lp-feat li{display:flex;align-items:center;gap:8px;min-height:32px;border-top:1px solid #1c2230;font-size:14px;color:#e9e4d8;line-height:18px;padding:6px 0}
.lp-tick{flex-shrink:0;width:14px;height:14px}
.lp-seg{display:inline-flex;border:1px solid #262e40;border-radius:4px;overflow:hidden;background:#0d111a}
.lp-seg span{display:inline-flex;align-items:center;justify-content:center;width:112px;height:44px;font-size:13px;font-weight:600;color:#9a9484}
.lp-seg span+span{border-left:1px solid #262e40}
.lp-seg span.on{color:#f2eee4;background:#131824;box-shadow:inset 0 -2px 0 #e5b955}
.lp-seg.focus span.on{outline:2px solid #e5b955;outline-offset:2px}
.lp-seg.full{display:flex;width:100%}.lp-seg.full span{flex:1;width:auto}
.lp-rule{border-left:2px solid #262e40;padding-left:14px;font-size:14px;color:#e9e4d8;max-width:720px;margin:0}
.lp-rule.gold{border-left-color:#e5b955}
.lp-how{display:grid;grid-template-columns:200px minmax(0,1fr);border-top:1px solid #1c2230}
.lp-how>div{min-height:44px;padding:12px 0;border-bottom:1px solid #1c2230;font-size:14px;color:#e9e4d8}
.lp-confirm{border:1px solid #1c2230;padding:16px;display:flex;flex-direction:column;gap:12px;border-radius:4px}
.lp-confirm.gold{border:0;border-left:2px solid #e5b955;background:#0d111a}
.lp-tbtn{display:inline-flex;align-items:center;min-height:44px;font-size:13px;color:#e9e4d8;text-decoration:underline;text-underline-offset:3px}
.lp-seat{display:grid;grid-template-columns:minmax(220px,1fr) minmax(0,1.4fr) auto;gap:16px;align-items:center;min-height:64px;border-bottom:1px solid #1c2230}
.lp-seat.phone{grid-template-columns:minmax(0,1fr) auto;row-gap:4px;padding:10px 0}
.lp-seat.phone .mid{grid-column:1 / -1;grid-row:2}
.lp-empty{width:32px;height:32px;border-radius:999px;border:1px dashed #262e40;flex-shrink:0}
.lp-opt{border:1px solid #262e40;border-radius:6px;padding:16px;min-height:96px;display:flex;gap:12px;background:#0d111a;flex:1}
.lp-opt.on{border-color:#e5b955;background:#131824}
.lp-radio{width:16px;height:16px;border-radius:999px;border:1px solid #9a9484;flex-shrink:0;margin-top:3px}
.lp-opt.on .lp-radio{border:5px solid #e5b955}
.lp-sel{display:flex;align-items:center;height:44px;padding:0 12px;border:1px solid #3a3326;border-radius:4px;background:#0d111a;color:#e9e4d8;font-size:14px}
.lp-sel .chev{margin-left:auto;color:#9a9484;font-size:11px}
.lp-mail{background:#ffffff;color:#1f1f1f;border-radius:6px;padding:20px 24px;font-family:Arial,Helvetica,sans-serif;font-size:14px;line-height:1.5}
.lp-mail .subj{font-size:18px;font-weight:700;margin-bottom:4px}
.lp-mail .from{font-size:12px;color:#5f6368;margin-bottom:14px}
.lp-mail a{color:#1a5fb4;text-decoration:underline}
.m{font-family:var(--font-mono);font-variant-numeric:tabular-nums;font-size:.9em}
.sr{font-size:11px;color:#9a9484;font-style:italic}
</style>'''


# Money and dates: one formatter each, the plans.ts formatPrice convention.

def money(amount: float) -> str:
    whole = abs(amount - round(amount)) < 0.005
    return f'${round(amount)}' if whole else f'${amount:.2f}'


def price_of(plan: str, interval: str, founder: bool) -> int:
    p = PRICES[plan]
    return p[f'founder_{interval}'] if founder else p[interval]


def price_words(plan: str, interval: str, founder: bool) -> str:
    return f'{money(price_of(plan, interval, founder))} a {"year" if interval == "yearly" else "month"}'


def day(d: date, year: bool = False) -> str:
    text = f'{d.day} {d.strftime("%b")}'
    return f'{text} {d.year}' if year else text


def long_day(d: date) -> str:
    return f'{d.day} {d.strftime("%B")} {d.year}'


def prorated(old: str, new: str, interval: str, renews: date) -> float:
    days_left = (renews - NOW).days
    period = 365 if interval == 'yearly' else 30
    return (price_of(new, interval, True) - price_of(old, interval, True)) * days_left / period


TICK = ('<svg class="lp-tick" viewBox="0 0 24 24" fill="none" stroke="' + KILL + '" stroke-width="3" '
        'stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M5 12l5 5 9-10"/></svg>')


def founder_pill() -> str:
    return pill('Founder', 'site')


# The plans page.

def band(founder: str, phone: bool, interval: str = 'yearly', focus: bool = False) -> str:
    seg_cls = 'lp-seg' + (' full' if phone else '') + (' focus' if focus else '')
    on_y, on_m = (' class="on"', '') if interval == 'yearly' else ('', ' class="on"')
    seg = f'<div class="{seg_cls}" role="radiogroup"><span{on_y}>Yearly</span><span{on_m}>Monthly</span></div>'
    line = dict(open=FOUNDER_OPEN, before=FOUNDER_BEFORE).get(founder)
    frow = (f'<p style="margin:0;font-size:13px;display:flex;align-items:center;gap:8px;flex-wrap:wrap">'
            f'{founder_pill()}<span>{line}</span></p>') if line else ''
    h1 = '<h1 class="display" style="font-size:22px;color:#f2eee4;margin:0">What is free, and what the plans add</h1>'
    head = (f'<div style="display:flex;flex-direction:column;gap:12px">{h1}{seg}</div>' if phone else
            f'<div style="display:flex;align-items:center;justify-content:space-between;gap:24px">{h1}{seg}</div>')
    pad = '22px 18px 24px' if phone else '28px 48px 24px'
    nav_html = nav_phone(False) if phone else nav('', False)
    return (f'<div class="band">{nav_html}<div class="lp-col" style="padding:{pad};display:flex;flex-direction:column;gap:10px">'
            f'<span class="label" style="color:#e5b955">Plans</span>{head}'
            f'<p style="margin:0;font-size:15px;max-width:720px">{PROMISE}</p>{frow}</div></div>')


def price_block(plan: str, interval: str, founder_open: bool) -> str:
    if plan == 'free':
        return ('<div class="lp-price" style="height:96px"><div><span class="lp-fig">$0</span></div>'
                '<span class="lp-sub">No card, no trial, no end date.</span></div>')
    unit = '/year' if interval == 'yearly' else '/month'
    fig = money(price_of(plan, interval, founder_open))
    was = f'<s class="lp-was">{money(price_of(plan, interval, False))}</s>' if founder_open else ''
    if interval == 'yearly':
        sub = f'{money(price_of(plan, "yearly", founder_open) / 12)} a month, billed yearly'
    else:
        sub = f'Or {price_words(plan, "yearly", founder_open)}.'
    fnote = f'<span class="lp-fnote">{FOUNDER_NOTE}</span>' if founder_open else ''
    return (f'<div class="lp-price" style="height:96px"><div><span class="lp-fig">{fig}</span><span class="lp-unit"> {unit}</span>{was}</div>'
            f'<span class="lp-sub">{sub}</span>{fnote}</div>')


def plan_card(plan: str, interval: str, founder_open: bool, btn_state: str = '') -> str:
    name = 'Free' if plan == 'free' else PRICES[plan]['name']
    if plan == 'free':
        action = ('<div style="height:44px;display:flex;align-items:center;font-size:13px;color:#9a9484">'
                  'Sign in with Battle.net from the top of any page.</div>')
    else:
        verb = 'Subscribe' if plan == 'premium' else 'Subscribe a guild'
        action = f'<span class="lp-btn {btn_state}">{verb} · {price_words(plan, interval, founder_open)}</span>'
    trial = (f'<p class="lp-trial" style="margin:0"><b style="color:#f2eee4;font-weight:600">First 7 days free</b>'
             f'{TRIAL[len("First 7 days free"):]}</p>') if plan in ('guild', 'raid_guild') else ''
    items = []
    for i, f in enumerate(FEATURES[plan]):
        strong = ' style="color:#f2eee4;font-weight:600"' if plan == 'raid_guild' and i == 0 else ''
        items.append(f'<li>{TICK}<span{strong}>{f}</span></li>')
    promise = ('<p style="margin:12px 0 0;font-size:13px;font-weight:600;color:#f2eee4">This list never shrinks.</p>'
               if plan == 'free' else '')
    return (f'<div class="lp-card"><span class="lp-name">{name}</span><span class="lp-who">{WHO[plan]}</span>'
            f'{price_block(plan, interval, founder_open)}<div style="margin-top:12px">{action}</div>'
            f'<div style="height:40px;margin-top:12px">{trial}</div>'
            f'<ul class="lp-feat" style="margin-top:8px">{"".join(items)}</ul>{promise}</div>')


def free_summary() -> str:
    return ('<a href="#free" class="lp-card" style="flex-direction:row;align-items:center;justify-content:space-between;'
            'min-height:64px;padding:12px 16px;color:inherit"><span style="display:flex;flex-direction:column">'
            '<span class="display" style="font-size:15px;font-weight:700;color:#f2eee4">Free</span>'
            '<span style="font-size:13px;color:#9a9484">Everything a player needs. This list never shrinks.</span></span>'
            '<span class="m" style="font-size:20px;color:#f2eee4">$0</span></a>')


def how_section(phone: bool) -> str:
    if phone:
        rows = ''.join(f'<div style="padding:12px 0;border-bottom:1px solid #1c2230"><span class="label">{t}</span>'
                       f'<p style="margin:4px 0 0;font-size:14px">{s}</p></div>' for t, s in HOW)
        body = f'<div style="border-top:1px solid #1c2230">{rows}</div>'
    else:
        cells = ''.join(f'<div class="label" style="display:flex;align-items:center">{t}</div><div>{s}</div>' for t, s in HOW)
        body = f'<div class="lp-how">{cells}</div>'
    return f'<section style="display:flex;flex-direction:column;gap:14px"><span class="h2">How the guild plans work</span>{body}</section>'


def plans_body(phone: bool, interval: str = 'yearly', founder: str = 'open') -> str:
    open_ = founder == 'open'
    order = ['premium', 'guild', 'raid_guild'] if phone else ['free', 'premium', 'guild', 'raid_guild']
    cards = [plan_card(p, interval, open_) for p in order]
    history = f'<p class="lp-rule">{HISTORY}</p>'
    foot = f'<p style="margin:0;font-size:13px;color:#9a9484">{FOOT}</p>'
    faq = ('<section style="display:flex;flex-direction:column"><span class="h2" style="margin-bottom:8px">FAQ</span>'
           + ''.join(f'<div style="padding:12px 0;border-bottom:1px solid #1c2230;font-size:15px;font-weight:600">{q}</div>'
                     for q in ('What happens if I cancel?', 'Can I get a refund?', 'Is anything from Battle.net ever paid?',
                               'Are there ads?', 'Who am I paying?')) + '</section>')
    note = f'<p class="ann">{MOCK_NOTE}</p>'
    if phone:
        free_full = f'<div id="free">{plan_card("free", interval, open_)}</div>'
        stack = free_summary() + ''.join(cards) + history + free_full
        return (f'<div style="padding:16px 18px 32px;display:flex;flex-direction:column;gap:16px">{stack}'
                f'<div style="height:6px"></div>{how_section(True)}{foot}{faq}{note}</div>')
    grid = f'<div style="display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:16px">{"".join(cards)}</div>'
    hist_row = (f'<div style="display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:16px">'
                f'<div></div><div></div><div style="grid-column:3 / span 2">{history}</div></div>')
    return (f'<div class="lp-col" style="padding:24px 48px 48px;display:flex;flex-direction:column;gap:32px">'
            f'<div style="display:flex;flex-direction:column;gap:16px">{grid}{hist_row}</div>'
            f'{how_section(False)}{foot}{faq}{note}</div>')


def plans_page(width: int, phone: bool) -> str:
    body = band('open', phone) + plans_body(phone) + footer(phone=phone)
    return page(CSS + LP_CSS + body, width)


def chooser_rows(rows: list) -> str:
    out = []
    for name, count, cls in rows:
        fit = 'Guild' if count <= GUILD_CAP else f'Raid guild fits {count} members'
        out.append(f'<div style="display:flex;align-items:center;gap:12px;min-height:56px;padding:0 12px;'
                   f'border-bottom:1px solid #1c2230">{crest_img(cls, 32)}<span style="display:flex;flex-direction:column">'
                   f'<span style="font-size:15px;color:#f2eee4">{name}</span><span style="font-size:13px;color:#9a9484">'
                   f'{count} verified members</span></span><span style="margin-left:auto;font-size:13px;color:#9a9484">'
                   f'{fit}</span></div>')
    return ('<div style="display:flex;flex-direction:column;gap:10px"><span class="sub">Which guild?</span>'
            f'<div style="border-top:1px solid #1c2230">{"".join(out)}</div></div>')


def states_page() -> str:
    monthly = ''.join(plan_card(p, 'monthly', True) for p in ('free', 'premium', 'guild', 'raid_guild'))
    cells = [
        cell('Monthly, founder open (heights equal to Yearly)',
             f'<div style="display:grid;grid-template-columns:repeat(4,minmax(0,1fr));gap:16px">{monthly}</div>', width=1344),
        cell('Before 4 November: list prices, the "from" founder line',
             band('before', False).replace(nav('', False), '') + '<div style="display:grid;grid-template-columns:repeat(2,1fr);gap:16px;margin-top:16px">'
             + plan_card('premium', 'yearly', False) + plan_card('guild', 'yearly', False) + '</div>', width=900),
        cell('After 6 January: no founder row, no strike',
             '<div style="display:grid;grid-template-columns:repeat(2,1fr);gap:16px">'
             + plan_card('premium', 'yearly', False) + plan_card('raid_guild', 'yearly', False) + '</div>', width=660),
        cell('Focus-visible on Yearly; hover and focus on a card button',
             band('open', False, focus=True).replace(nav('', False), '')
             + '<div style="display:flex;gap:16px;margin-top:16px;width:640px">'
             + '<span class="lp-btn hover" style="flex:1">Subscribe · ' + price_words('premium', 'yearly', True) + '</span>'
             + '<span class="lp-btn focus" style="flex:1">Subscribe a guild · ' + price_words('guild', 'yearly', True) + '</span></div>',
             width=900),
        cell('Checkout guild chooser (no guild id in the URL)',
             chooser_rows([(GUILD, VERIFIED, 'warrior'), ('Mograine Raiders', 19, 'paladin')]), width=640),
        cell('Checkout guild chooser, no billable guild',
             '<p style="margin:0;font-size:14px">You can subscribe a guild you lead or have claimed. '
             '<a href="#guild" style="text-decoration:underline">Find your guild</a></p>', width=640),
    ]
    title = '<span class="h2">Plans page: other states</span>'
    body = f'<div style="padding:32px 48px;display:flex;flex-wrap:wrap;gap:32px">{title}<div style="width:100%"></div>{"".join(cells)}</div>'
    return page(CSS + LP_CSS + body, 1440)


# Guild settings, the Billing section.

RENEWS = date(2027, 11, 18)
TRIAL_CHARGE = date(2026, 12, 11)
SUBSCRIBER = 'Kraggor'
NEW_PAYER = 'Bow Jackzon'


def btn(text: str, kind: str = 's', off: bool = False, full: bool = False) -> str:
    cls = 'btnp' if kind == 'p' else 'btns'
    parts = ['border-color:#e5b955;color:#e5b955'] if kind == 'g' else []
    if full:
        parts.append('width:100%;justify-content:center')
    style = f' style="{";".join(parts)}"' if parts else ''
    return f'<span class="{cls}{" off" if off else ""}"{style}>{text}</span>'


def tbtn(text: str) -> str:
    return f'<span class="lp-tbtn">{text}</span>'


def plan_line(text: str, line2: str, founder: bool, buttons: str = '', phone: bool = False) -> str:
    fp = f' {founder_pill()}' if founder else ''
    left = (f'<div style="display:flex;flex-direction:column;gap:4px"><p style="margin:0;font-size:14px;display:flex;'
            f'align-items:center;gap:8px;flex-wrap:wrap">{text}{fp}</p>'
            f'<p style="margin:0;font-size:13px;color:#9a9484">{line2}</p></div>')
    if not buttons:
        return left
    if phone:
        return (f'<div style="display:flex;flex-direction:column;gap:12px">{left}'
                f'<div style="display:flex;flex-direction:column;gap:8px">{buttons}</div></div>')
    return (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:16px">{left}'
            f'<div style="display:flex;gap:8px">{buttons}</div></div>')


def who(name: str, cls: str, sub: str, extra: str = '') -> str:
    sub_html = f'<span style="font-size:12px;color:#9a9484">{sub}</span>' if sub else ''
    return (f'<div style="display:flex;align-items:center;gap:10px">{crest_img(cls, 32)}<span style="display:flex;flex-direction:column">'
            f'<span style="display:flex;align-items:center;gap:8px"><span class="display" style="font-size:14px;font-weight:700;'
            f'color:{TEXT_CLASS[cls]}">{name}</span>{extra}</span>{sub_html}</span></div>')


def seat_row(state: str, phone: bool, **k) -> str:
    cls = 'lp-seat phone' if phone else 'lp-seat'
    if state == 'held':
        when = date.fromordinal(k['given'].toordinal() + SEAT_LOCK_DAYS)
        locked = when > NOW
        mid = (f'<span style="color:#9a9484">Can move from <span class="m">{day(when)}</span></span>' if locked
               else '<span>Can move now</span>')
        left, right = who(k['name'], k['cls'], f'Given {day(k["given"])}'), btn('Move seat', off=locked)
    elif state == 'offered':
        left = who(k['name'], k['cls'], '', pill('Waiting', 'sample'))
        mid = f'<span>Offered <span class="m">{day(k["at"])}</span>. They pay for Premium, so they accept it first.</span>'
        right = tbtn('Withdraw')
    elif state == 'accepted':
        left = who(k['name'], k['cls'], f'Accepted {day(k["at"])}')
        mid = f'<span>Accepted · starts <span class="m">{day(k["starts"])}</span>, when their own Premium ends</span>'
        right = btn('Move seat', off=True)
    else:
        left = ('<div style="display:flex;align-items:center;gap:10px"><span class="lp-empty"></span>'
                '<span style="font-size:14px;color:#9a9484">Not given</span></div>')
        ret = k.get('returned')
        mid = (f'<span style="color:#9a9484">Free since <span class="m">{day(ret[1])}</span>: {ret[0]} left the guild.</span>'
               if ret else '')
        right = btn('Give seat', 'g')
    return f'<div class="{cls}"><div>{left}</div><div class="mid" style="font-size:13px">{mid}</div><div>{right}</div></div>'


SEATS_MAIN = [
    ('held', dict(name='Lightbrand', cls='priest', given=date(2026, 11, 20))),
    ('offered', dict(name='Zulmara', cls='hunter', at=NOW)),
    ('empty', dict()),
]
SUBHEAD = 'class="sub" style="text-transform:uppercase;letter-spacing:.04em"'


def seats_block(phone: bool, rows: list = SEATS_MAIN, total: int = 3, extra: str = '') -> str:
    given = sum(1 for s, _ in rows if s != 'empty')
    head = (f'<div style="display:flex;align-items:baseline;gap:10px"><span {SUBHEAD}>Premium seats</span>'
            f'<span class="m" style="font-size:13px;color:#9a9484">{given} of {total} given</span></div>')
    helper = ('<p style="margin:0;font-size:13px;color:#9a9484">A seat gives one verified member Premium while the plan lasts. '
              'A seat can move once every 30 days.</p>')
    body = ''.join(seat_row(s, phone, **k) for s, k in rows)
    return (f'<div style="display:flex;flex-direction:column;gap:8px">{head}{helper}'
            f'<div style="border-top:1px solid #1c2230">{body}</div>{extra}</div>')


def who_pays(phone: bool, state: str = 'idle') -> str:
    head = f'<span {SUBHEAD}>Who pays</span>'
    helper = ('<p style="margin:0;font-size:13px;color:#9a9484">Hand the plan to another officer, for example when the guild '
              'leader changes. The founder price goes with it.</p>')
    if state == 'waiting':
        act = (f'<p style="margin:0;font-size:14px">Offered to {NEW_PAYER} on <span class="m">{day(NOW)}</span>. '
               f'Waiting for them. {tbtn("Withdraw")}</p>')
    elif state == 'confirm':
        price = price_words('raid_guild', 'yearly', True)
        act = (f'<div class="lp-confirm"><span class="label">Transfer to</span>'
               f'<div class="lp-sel">{NEW_PAYER} · Warrior<span class="chev">▾</span></div>'
               f'<p style="margin:0;font-size:14px">{NEW_PAYER} gets an offer to take over. When they add a card, they are billed '
               f'from <span class="m">{day(RENEWS, True)}</span> at <span class="m">{price}</span>, and your plan stops renewing. '
               f'Nothing changes until they accept.</p><div style="display:flex;gap:12px;align-items:center">'
               f'{btn("Send the offer", "p", full=phone)}{tbtn("Cancel")}</div></div>')
    else:
        act = f'<div>{btn("Transfer the plan", full=phone)}</div>'
    return f'<div style="display:flex;flex-direction:column;gap:8px">{head}{helper}{act}</div>'


def records(phone: bool) -> str:
    dirn = 'column' if phone else 'row'
    kept = "Your guild's raids, attendance and loot are kept for good."
    return (f'<div style="display:flex;flex-direction:column;gap:8px"><span {SUBHEAD}>Records</span>'
            f'<p style="margin:0;font-size:13px;color:#9a9484">{kept}</p><div style="display:flex;flex-direction:{dirn};gap:8px">'
            f'{btn("Download attendance (CSV)", full=phone)}{btn("Download loot (CSV)", full=phone)}</div></div>')


def chooser(phone: bool, trial_used: bool = False, lapsed: bool = False) -> str:
    founder = not lapsed
    pick = 'raid_guild' if VERIFIED > GUILD_CAP else 'guild'
    lines = dict(guild='Up to 25 verified members · 1 Premium seat',
                 raid_guild='Any size · loot council and readiness · 3 Premium seats')
    opts = []
    for plan in ('guild', 'raid_guild'):
        was = f'<s class="lp-was">{money(price_of(plan, "yearly", False))}</s>' if founder else ''
        opts.append(f'<div class="lp-opt{" on" if plan == pick else ""}"><span class="lp-radio"></span><span style="display:flex;'
                    f'flex-direction:column;gap:4px"><span class="display" style="font-size:15px;font-weight:700;color:#f2eee4">'
                    f'{PRICES[plan]["name"]}</span><span><span class="m" style="font-size:20px;color:#f2eee4">'
                    f'{money(price_of(plan, "yearly", founder))}</span><span class="lp-unit"> /year</span>{was}</span>'
                    f'<span style="font-size:13px;color:#9a9484">{lines[plan]}</span></span></div>')
    dirn = 'column' if phone else 'row'
    if lapsed:
        label = 'Subscribe again'
    elif trial_used:
        label = f'Subscribe · {price_words(pick, "yearly", founder)}'
    else:
        label = 'Start the 7-day trial'
    trial = '' if trial_used or lapsed else f'<p style="margin:0;font-size:13px">{TRIAL}</p>'
    fl = ('<p style="margin:0;font-size:12px;color:#e5b955">Founder price until 6 January, kept while the plan stays active.</p>'
          if founder else '')
    seg = f'<div class="lp-seg{" full" if phone else ""}"><span class="on">Yearly</span><span>Monthly</span></div>'
    return (f'<div style="display:flex;flex-direction:column;gap:12px">{seg}'
            f'<div style="display:flex;flex-direction:{dirn};gap:12px">{"".join(opts)}</div>'
            f'<p style="margin:0;font-size:13px;color:#9a9484">{VERIFIED} verified members today.</p>'
            f'<div>{btn(label, "p", full=phone)}</div>{trial}{fl}</div>')


def raid_guild_line(phone: bool) -> str:
    text = f'Raid guild · renews <span class="m">{day(RENEWS, True)}</span> · {price_words("raid_guild", "yearly", True)}'
    return plan_line(text, 'Billed to you.', True, btn('Manage billing', full=phone), phone)


def billing_section(phone: bool) -> str:
    gap = 20
    parts = [raid_guild_line(phone), seats_block(phone), who_pays(phone), records(phone)]
    return (f'<section style="display:flex;flex-direction:column;gap:{gap}px"><span class="h2">Billing</span>'
            + ''.join(parts) + '</section>')


def settings_page(phone: bool) -> str:
    band_html = guild_band('officer', phone, active='Settings')
    pad = '22px 18px 32px' if phone else '32px 48px 48px'
    width = '' if phone else 'max-width:880px;'
    note = f'<p class="ann">{MOCK_NOTE} Kraggor holds the plan (Raid guild, founder, yearly).</p>'
    body = (f'<div class="lp-col" style="padding:{pad}"><div style="{width}display:flex;flex-direction:column;gap:24px">'
            f'<h1 class="h2">Settings</h1>{billing_section(phone)}{note}</div></div>')
    return page(CSS + LP_CSS + band_html + body + footer(phone=phone), 390 if phone else 1440)


def size_and_confirm(loading: bool) -> str:
    guild_line = plan_line(f'Guild · renews <span class="m">{day(RENEWS, True)}</span> · {price_words("guild", "yearly", True)}',
                           'Billed to you.', True, btn('Manage billing'))
    size = (f'<div style="display:flex;align-items:center;justify-content:space-between;gap:16px"><p class="lp-rule gold">{VERIFIED} verified '
            f'members. Guild covers up to 25; nothing switches off. Raid guild covers any size.</p>{btn("Move to Raid guild", "g")}</div>')
    amount = (sk(14, '64px').replace('<div ', '<span ').replace('</div>', '</span>').replace('style="', 'style="display:inline-block;vertical-align:middle;')
              if loading else f'<span class="m">{money(prorated("guild", "raid_guild", "yearly", RENEWS))}</span>')
    confirm = (f'<div class="lp-confirm"><p style="margin:0;font-size:14px">Move to Raid guild? Today you pay {amount} for the rest '
               f'of this year. From <span class="m">{day(RENEWS, True)}</span> it is {price_words("raid_guild", "yearly", True)}.</p>'
               f'<div style="display:flex;gap:12px;align-items:center">{btn("Move to Raid guild", "p", off=loading)}'
               f'<span class="lp-tbtn" style="outline:2px solid #e5b955;outline-offset:2px">Keep Guild</span></div></div>')
    return f'<div style="display:flex;flex-direction:column;gap:16px">{guild_line}{size}{confirm}</div>'


def picker_open() -> str:
    rows = [('held', SEATS_MAIN[0][1]), ('offered', SEATS_MAIN[1][1]), ('empty', dict())]
    picker = ('<div class="lp-confirm" style="margin-top:8px"><span class="label">Give this seat to</span>'
              '<div class="lp-sel">Earthhoof · Druid · pays for Premium<span class="chev">▾</span></div>'
              '<p style="margin:0;font-size:13px;color:#9a9484">They pay for Premium, so they get an offer to accept. Their own '
              'Premium stops renewing and the seat starts when it ends.</p>'
              f'<div style="display:flex;gap:12px;align-items:center">{btn("Give seat", "p")}{tbtn("Cancel")}</div></div>')
    return seats_block(False, rows, extra=picker)


def receiver_offer() -> str:
    price = price_words('raid_guild', 'yearly', True)
    return (f'<div class="lp-confirm gold"><p style="margin:0;font-size:14px">{SUBSCRIBER} offers you the guild plan. You would be '
            f'billed from <span class="m">{day(RENEWS, True)}</span> at <span class="m">{price}</span>, the founder price. '
            f'Nothing is charged before then.</p><div style="display:flex;gap:12px;align-items:center">'
            f'{btn("Take over the plan", "p")}{tbtn("Not now")}</div></div>')


def billing_states_page() -> str:
    trial_line = plan_line(f'Guild · trial · first charge <span class="m">{day(TRIAL_CHARGE, True)}</span>: '
                           f'{money(price_of("guild", "yearly", True))}',
                           f'Cancel before <span class="m">{day(TRIAL_CHARGE)}</span> and nothing is charged. Billed to you.',
                           True, btn('Manage billing'))
    more_rows = [('accepted', dict(name='Tidecaller', cls='shaman', at=date(2026, 12, 1), starts=date(2026, 12, 22))),
                 ('empty', dict(returned=('Sagebloom', date(2026, 12, 2))))]
    lapsed = (plan_line(f'Guild plan ended <span class="m">{day(RENEWS, True)}</span>.',
                        'Officer tools are read-only. Raids, attendance, loot and reports stay readable, and the export below '
                        'always works.', False)
              + '<div style="height:16px"></div>' + chooser(False, lapsed=True) + '<div style="height:20px"></div>' + records(False))
    read_only_rows = ''.join(seat_row(s, False, **k).replace(btn('Move seat', off=True), '').replace(tbtn('Withdraw'), '')
                             .replace(btn('Give seat', 'g'), '') for s, k in SEATS_MAIN)
    read_only = (plan_line(f'Raid guild · renews <span class="m">{day(RENEWS, True)}</span> · '
                           f'{price_words("raid_guild", "yearly", True)}', f'Billed by {SUBSCRIBER}', True)
                 + f'<div style="height:16px"></div><div style="border-top:1px solid #1c2230">{read_only_rows}</div>'
                 + '<div style="height:20px"></div>' + records(False))
    gate = ('<p style="margin:0;font-size:13px">Loot council and the readiness board are part of Raid guild. '
            '<a href="#billing" style="text-decoration:underline;white-space:nowrap">Move to Raid guild</a>.</p>')
    cells = [
        cell('Not on a plan: chooser, trial available (28 members, Raid guild preselected)', chooser(False), width=660),
        cell('Not on a plan: trial already used', chooser(False, trial_used=True), width=660),
        cell('Trial', trial_line, width=660),
        cell('Guild with 28 members: size line, confirm loading (focus on Keep Guild)', size_and_confirm(True), width=660),
        cell('Confirm ready: prorated amount from the preview', size_and_confirm(False), width=660),
        cell('Seat rows: accepted (starts later), returned', seats_block(False, more_rows), width=660),
        cell('Give seat: picker open, a payer chosen', picker_open(), width=660),
        cell('Who pays: send confirm', who_pays(False, 'confirm'), width=660),
        cell('Who pays: waiting', who_pays(False, 'waiting'), width=660),
        cell("Receiver's offer (top of #billing for Bow Jackzon)", receiver_offer(), width=660),
        cell('Lapsed: read-only line, Subscribe again at list price, Records', lapsed, width=660),
        cell('Officer who may not bill: read-only seats, Records', read_only, width=660),
        cell('Gate sentence on a Guild-tier guild (Loot and Readiness tabs)', gate, width=660),
    ]
    title = '<span class="h2">Guild settings, Billing: every other state</span>'
    body = (f'<div style="padding:32px 48px;display:flex;flex-wrap:wrap;gap:32px">{title}<div style="width:100%"></div>'
            f'{"".join(cells)}<p class="ann" style="width:100%">{MOCK_NOTE}</p></div>')
    return page(CSS + LP_CSS + body, 1440)


# The account page, Plan panel. The member is Zulmara: founder Premium, monthly, renews 22 Dec 2026.

MEMBER = 'Zulmara'
OWN_RENEWS = date(2026, 12, 22)
OWN_PRICE = price_words('premium', 'monthly', True)
ROWS = [('Server sims', 'On · 4 runs this month'), ('History charts', 'Open on your main character'),
        ('Discord notifications', 'On'), ('Supporter mark', 'Shown on your characters')]


def state_panel(inner: str) -> str:
    dot = '<span style="width:8px;height:8px;border-radius:999px;background:#e5b955;box-shadow:0 0 10px rgba(229,185,85,.8)"></span>'
    return (f'<div class="panel" style="padding:20px 24px;display:flex;flex-direction:column;gap:16px">'
            f'<div style="display:flex;align-items:center;gap:10px">{dot}<span class="label" style="color:#e5b955">Plan</span></div>'
            f'{inner}</div>')


def premium_rows(phone: bool) -> str:
    cols = '8.5rem' if phone else '10rem'
    return ''.join(f'<div style="display:grid;grid-template-columns:{cols} minmax(0,1fr);gap:12px;min-height:28px;font-size:14px">'
                   f'<span style="color:#9a9484">{k}</span><span style="text-decoration:underline">{v}</span></div>' for k, v in ROWS)


def offer_block(phone: bool) -> str:
    sentence = (f'If you accept, your own Premium stops renewing and ends <span class="m">{day(OWN_RENEWS, True)}</span>, and the '
                f'seat starts then. You are not charged twice. If the seat is ever taken away, we email you a link to come back '
                f'at <span class="m">{OWN_PRICE}</span>, the price you pay now.')
    dirn = 'column' if phone else 'row'
    return (f'<div class="lp-confirm gold"><p style="margin:0;font-size:15px;color:#f2eee4">{GUILD} offers you a Premium seat.</p>'
            f'<p style="margin:-8px 0 0;font-size:13px;color:#9a9484">From {SUBSCRIBER}, <span class="m">{day(NOW)}</span>.</p>'
            f'<p style="margin:0;font-size:14px">{sentence}</p><div style="display:flex;flex-direction:{dirn};gap:12px;'
            f'align-items:{"stretch" if phone else "center"}">{btn("Accept the seat", "p", full=phone)}{tbtn("Decline")}</div></div>')


def own_line() -> str:
    return plan_line(f'Premium · renews <span class="m">{day(OWN_RENEWS, True)}</span> · {OWN_PRICE}', '', True).replace(
        '<p style="margin:0;font-size:13px;color:#9a9484"></p>', '')


def account_states(phone: bool) -> list:
    manage = f'<div>{btn("Manage billing", full=phone)}</div>'
    offer = state_panel(offer_block(phone) + own_line() + premium_rows(phone) + manage)
    accepted = state_panel(f'<p style="margin:0;font-size:13px;color:#9a9484">Accepted. Your Premium ends <span class="m">'
                           f'{day(OWN_RENEWS, True)}</span>; the seat starts then.</p>'
                           + plan_line(f'Premium · ends <span class="m">{day(OWN_RENEWS, True)}</span> · then a seat from {GUILD}',
                                       '', True).replace('<p style="margin:0;font-size:13px;color:#9a9484"></p>', '')
                           + premium_rows(phone) + manage)
    seat = state_panel(plan_line(f'Premium · a seat from {GUILD}',
                                 f"Given by {SUBSCRIBER}. It lasts while the guild's plan does; nothing is billed to you.", False)
                       + premium_rows(phone))
    founder = state_panel(own_line() + '<p class="sr">sr-only after the pill: Founder price, kept while your plan stays active.</p>'
                          + premium_rows(phone) + manage)
    ended = state_panel(f'<p style="margin:0;font-size:14px">Your seat from {GUILD} ended <span class="m">{day(date(2027, 2, 3))}</span>. '
                        f'<a href="#resubscribe" style="text-decoration:underline">Come back at {OWN_PRICE}</a></p>'
                        '<p style="margin:0;font-size:14px"><span style="color:#9a9484">Plan</span> · Not on Premium. '
                        '<span style="text-decoration:underline">See plans</span></p>')
    return [('Seat offer to a payer (the panel opens at #plan)', offer), ('Accepted', accepted),
            ('Holding a seat (no Manage billing)', seat), ('Own Premium at the founder price', founder),
            ('Seat ended, for a member who gave up their own Premium for it (30 days)', ended)]


def account_page() -> str:
    cells = ''.join(cell(t, b, width=640) for t, b in account_states(False))
    guild_row = (f'<p style="margin:0;font-size:13px;color:#9a9484">Raid guild plan · renews <span class="m">{day(RENEWS, True)}</span>'
                 ' · you hold a seat · <span style="text-decoration:underline">Guild settings</span></p>')
    cells += cell('Guild row on the account page (AccountGuildPlanLine)', guild_row, width=640)
    title = '<span class="h2">Account, Plan panel</span>'
    body = (f'<div style="padding:32px 48px;display:flex;flex-wrap:wrap;gap:32px">{title}<div style="width:100%"></div>{cells}'
            f'<p class="ann" style="width:100%">{MOCK_NOTE} {MEMBER} pays for Premium at the founder price, monthly.</p></div>')
    return page(CSS + LP_CSS + body, 1440)


def account_phone() -> str:
    _, offer = account_states(True)[0]
    body = (f'{nav_phone(True, None)}<div style="padding:22px 18px 32px;display:flex;flex-direction:column;gap:16px">'
            f'<h1 class="h2">Account</h1>{offer}<p class="ann">{MOCK_NOTE}</p></div>')
    return page(CSS + LP_CSS + body + footer(phone=True), 390)


# Emails: subject, first sentence, one link.

def emails() -> list:
    renew = price_words('raid_guild', 'yearly', True)
    given = (f"{SUBSCRIBER} gave you a Premium seat in {GUILD}, so you have Premium for as long as the guild's plan lasts, "
             "with nothing billed to you.")
    return [
        ('Renewal, 7 days before (11 Nov 2027)', f'Your Raid guild plan renews on {RENEWS.day} {RENEWS.strftime("%B")}',
         f'Your Raid guild plan for {GUILD} renews on {long_day(RENEWS)} for {renew}, your founder price; to stop it, cancel '
         'before then and nothing more is charged.', 'Manage or cancel'),
        ('Size over 25, 14 days before (4 Nov 2027)', f'{GUILD} has {VERIFIED} members; Guild covers 25',
         f'{GUILD} has {VERIFIED} verified members and its Guild plan covers up to 25, so when it renews on {long_day(RENEWS)} '
         f'you may want Raid guild at {renew}; nothing switches off either way.', 'Look at Raid guild'),
        ('Seat offered, to a payer', f'{GUILD} offers you a Premium seat',
         f'{SUBSCRIBER} offered you a Premium seat in {GUILD}; if you accept, your own Premium stops renewing and ends on '
         f'{long_day(OWN_RENEWS)}, the seat starts then, and you are not charged twice.', 'Accept or decline'),
        ('Seat given, to a non-payer', f'You have Premium from {GUILD}', given, 'See your plan'),
        ('Seat removed, to someone who gave up their own Premium', f'Your Premium seat from {GUILD} has ended',
         f'Your Premium seat in {GUILD} ended on {long_day(date(2027, 2, 3))}, and because you gave up your own Premium for it, '
         f'this link brings it back at {OWN_PRICE}, the price you had.', f'Come back at {OWN_PRICE}'),
    ]


def email_card(subject: str, first: str, link: str) -> str:
    return (f'<div class="lp-mail"><div class="subj">{subject}</div><div class="from">Forever Sixty, help at foreversixty.gg</div>'
            f'<p style="margin:0 0 14px">{first}</p><a href="#link">{link}</a></div>')


def emails_page() -> str:
    cells = ''.join(cell(t, email_card(s, f, l), width=640) for t, s, f, l in emails())
    title = '<span class="h2">Emails: subject, first sentence, link</span>'
    body = (f'<div style="padding:32px 48px;display:flex;flex-wrap:wrap;gap:32px">{title}<div style="width:100%"></div>{cells}'
            f'<p class="ann" style="width:100%">{MOCK_NOTE} Plain text first; the light card stands for a mail client.</p></div>')
    return page(CSS + LP_CSS + body, 1440)


# The premium-value board (spec B3.5, copy B5): the comparison board for capture row 22.

# Mirrors RAID_SIZES in web/src/lib/billing/plans.ts, the source of truth (display order: 40 leads).
RAID_SIZES = (40, 20, 10)
PV_PREMIUM_SUPPORT = 'Free stays complete. Premium pays for the servers and keeps the site going.'
PV_PREMIUM_FEATURES = [
    'Try up to 5,000 gear, talent or drop combinations in one sim and find your best DPS. Your browser stops at 400.',
    'Close the tab. Server sims keep running, and the site tells you when they finish.',
    'Runs go up to 100,000 iterations instead of 30,000, so close upgrades come out clear.',
    'See if you are improving: your Forever Score and every boss parse on a chart over time, with the fights behind each point.',
    'A message in your own Discord channel when a server sim finishes or a log is parsed.',
    "A Supporter mark beside your name on your characters and in your guild's roster.",
]
PV_PREMIUM_SCOPE = ('Sims cover damage specs; healing and tanking specs are not simulated yet. History, Discord messages '
                    'and the Supporter mark work for every spec.')
PV_EXAMPLE_LABEL = 'EXAMPLE · A REAL SERVER RUN'
PV_EXAMPLE_CAPTION = "A level-60 Fury Warrior in the site's preset gear: 4,812 combinations on our servers, 27 Oct 2026."
PV_EXAMPLE_LINK = 'Open this result'
PV_BAND_LEAD = 'Running a guild?'
PV_BAND_LINK = 'Guild and Raid guild are further down'
PV_GUILD_HEADING = 'For guild leaders'
PV_GUILD_LEDE = 'One officer pays, and the whole guild gets raids, sign-ups and attendance. Members never pay to sign up.'
PV_GUILD_MEMBERS = ('Plans count verified members: players with a Forever Sixty account whose character the guild has '
                    'confirmed. People in your Discord without an account do not count.')
PV_GUILD_FEATURES = [
    'Post raids on the site and in Discord; members sign up with one tap.',
    'See who came: attendance and a report for every raid night from its log.',
    'The Forever Sixty bot in your Discord: raid posts with buttons, plus /bis and /char lookups.',
    '1 Premium seat to give a member, such as your raid leader.',
    f'Up to {GUILD_CAP} verified members, so a guild that runs 10- or 20-person raids fits under Guild. '
    'Nothing switches off if you grow past it.',
]
PV_RAID_FEATURES = [
    'Loot council: each drop ranks the raiders it helps most, with what they wear now, their attendance and their loot '
    'this tier. Officers award it from the same screen.',
    'Readiness board: before the pull, one list of who is missing enchants, consumables, talent points or upgrades, worst first.',
    '3 Premium seats to give members.',
    'Any number of members.',
    'Everything in Guild: raids, sign-ups, attendance and the Discord bot.',
]
PV_TRIAL = ('First 7 days free, once per guild. Cancel before day 8 and nothing is charged; guild settings shows the '
            'date of the first charge.')
HOW_BY_TERM = dict(HOW)
PV_HOW = [
    ('Verified members', 'A verified member is a player with a Forever Sixty account whose character the guild has confirmed: '
                         "an officer approved it, they joined by the guild's invite link, or the character is in the guild's "
                         'logs. Each player counts once, however many characters they have in the guild. Discord members '
                         'without an account do not count.'),
    ('Size', HOW_BY_TERM['Size']),
    ('Paying together', "One person's card pays for the plan. Members cannot split the bill on the site."),
    ('Seats', HOW_BY_TERM['Seats']),
    ('Moving a seat', HOW_BY_TERM['Moving a seat']),
    ('Already on Premium', HOW_BY_TERM['Already on Premium']),
    ('Trial', "A guild's first 7 days are free, once per guild. The card is added at the start, and guild settings shows the "
              'date of the first charge from day one. Cancel before day 8 and nothing is charged.'),
    ('Renewals', HOW_BY_TERM['Renewals']),
    ('A new guild leader', HOW_BY_TERM['A new guild leader']),
    ('Refunds', HOW_BY_TERM['Refunds']),
]
PV_COST_BOX_PX = 54
PV_GUILD_AUDIENCE = f'For a guild of up to {GUILD_CAP} verified members'


def pv_cost_line(plan: str) -> str:
    amount = price_of(plan, 'yearly', True)
    if plan == 'guild':
        return f'{money(amount / GUILD_CAP)} per member a year at {GUILD_CAP} members.'
    a40, a20, a10 = (money(amount / n) for n in RAID_SIZES)
    return f'{a40} per raider a year in a {RAID_SIZES[0]}-person raid, {a20} in a {RAID_SIZES[1]}, {a10} in a {RAID_SIZES[2]}.'


def pv_ticks(lines: list) -> str:
    items = ''.join(f'<li>{TICK}<span>{f}</span></li>' for f in lines)
    return f'<ul class="lp-feat" style="margin-top:8px">{items}</ul>'


def pv_free_card() -> str:
    return ('<div class="lp-card" style="box-shadow:none"><span class="lp-name">Free</span>'
            f'<span class="lp-who">{WHO["free"]}</span>{price_block("free", "yearly", True)}'
            '<div style="margin-top:12px"><div style="height:44px;display:flex;align-items:center;font-size:13px;color:#9a9484">'
            'Sign in with Battle.net from the top of any page.</div></div><div style="height:40px;margin-top:12px"></div>'
            f'{pv_ticks(FEATURES["free"])}'
            '<p style="margin:12px 0 0;font-size:13px;font-weight:600;color:#f2eee4">This list never shrinks.</p></div>')


def pv_premium_card() -> str:
    button = f'<span class="lp-btn">Subscribe · {price_words("premium", "yearly", True)}</span>'
    return (f'<div class="lp-card"><span class="lp-name">Premium</span><span class="lp-who">{WHO["premium"]}</span>'
            f'{price_block("premium", "yearly", True)}<div style="margin-top:12px">{button}</div>'
            f'<p class="lp-trial" style="height:40px;margin:12px 0 0">{PV_PREMIUM_SUPPORT}</p>'
            f'{pv_ticks(PV_PREMIUM_FEATURES)}'
            f'<p style="margin:12px 0 0;font-size:13px;color:#9a9484">{PV_PREMIUM_SCOPE}</p></div>')


def pv_example_card() -> str:
    shot = ('<div style="aspect-ratio:4/3;width:100%;border:1px solid #1c2230;border-radius:4px;background:#131824;'
            'display:flex;flex-direction:column;gap:8px;padding:12px;box-sizing:border-box">'
            f'<div style="display:flex;align-items:center;gap:8px">{crest_img("warrior", 32)}'
            '<span style="display:flex;flex-direction:column"><span style="font-size:14px;color:#f2eee4">Fury Warrior</span>'
            '<span class="sub">Level 60, preset gear</span></span></div>'
            '<span class="sr" style="margin-top:auto">Mock capture of the real result page: combinations, precision, best set '
            'DPS with its error, first swaps with item icons.</span></div>')
    return ('<div class="lp-card" style="box-shadow:none;border-top-color:#262e40">'
            f'<span class="label" style="color:#9a9484;font-size:11px">{PV_EXAMPLE_LABEL}</span>'
            f'<div style="margin-top:8px">{shot}</div>'
            f'<p style="margin:8px 0 0;font-size:13px;color:#9a9484">{PV_EXAMPLE_CAPTION}</p>'
            f'<a href="#result" style="display:flex;align-items:center;min-height:44px;color:#e5b955;font-size:14px">'
            f'{PV_EXAMPLE_LINK}</a></div>')


def pv_guild_price(plan: str) -> str:
    fig = money(price_of(plan, 'yearly', True))
    was = money(price_of(plan, 'yearly', False))
    sub = f'{money(price_of(plan, "yearly", True) / 12)} a month, billed yearly'
    return ('<div class="lp-price" style="height:152px;overflow:hidden">'
            f'<div><span class="lp-fig" style="font-size:24px;line-height:30px">{fig}</span><span class="lp-unit"> /year</span>'
            f'<s class="lp-was">{was}</s></div><span class="lp-sub">{sub}</span><span class="lp-fnote">{FOUNDER_NOTE}</span>'
            f'<p style="margin:0;font-size:13px;line-height:18px;height:{PV_COST_BOX_PX}px;color:#e9e4d8">'
            f'{pv_cost_line(plan)}</p></div>')


def pv_guild_card(plan: str, features: list) -> str:
    button = ('<span class="lp-btn" style="border-color:#3a3326;color:#e9e4d8">'
              f'Subscribe a guild · {price_words(plan, "yearly", True)}</span>')
    trial = (f'<p class="lp-trial" style="margin:0"><b style="color:#f2eee4;font-weight:600">First 7 days free</b>'
             f'{PV_TRIAL[len("First 7 days free"):]}</p>')
    return (f'<div class="lp-card" style="box-shadow:none;border-top-color:#262e40">'
            f'<span class="lp-name">{PRICES[plan]["name"]}</span>'
            f'<span class="lp-who">{PV_GUILD_AUDIENCE if plan == "guild" else WHO[plan]}</span>{pv_guild_price(plan)}<div style="margin-top:12px">{button}</div>'
            f'<div style="min-height:40px;margin-top:12px">{trial}</div>{pv_ticks(features)}</div>')


def pv_gutter(width: int) -> int:
    return 32 if width < 1280 else 48


def pv_band(width: int) -> str:
    seg = '<div class="lp-seg" role="radiogroup"><span class="on">Yearly</span><span>Monthly</span></div>'
    h1 = '<h1 class="display" style="font-size:22px;color:#f2eee4;margin:0">What is free, and what the plans add</h1>'
    frow = (f'<p style="margin:0;font-size:13px;display:flex;align-items:center;gap:8px;flex-wrap:wrap">{founder_pill()}'
            f'<span>{FOUNDER_OPEN}</span></p>')
    guild_link = (f'<p style="margin:0;font-size:13px;display:flex;align-items:center;gap:6px;min-height:44px">{PV_BAND_LEAD} '
                  f'<a href="#guild-leaders" style="color:#e5b955">'
                  f'{PV_BAND_LINK}</a></p>')
    return (f'<div class="band">{nav("", False)}<div class="lp-col" style="padding:28px {pv_gutter(width)}px 24px;display:flex;'
            f'flex-direction:column;gap:10px"><span class="label" style="color:#e5b955">Plans</span>'
            f'<div style="display:flex;align-items:center;justify-content:space-between;gap:24px">{h1}{seg}</div>'
            f'<p style="margin:0;font-size:15px;max-width:720px">{PROMISE}</p>{frow}{guild_link}</div></div>')


def pv_how() -> str:
    cells = ''.join(f'<div class="label" style="display:flex;align-items:center">{t}</div><div>{s}</div>' for t, s in PV_HOW)
    return (f'<div style="display:flex;flex-direction:column;gap:14px"><span class="section-title" style="font-size:18px">'
            f'How the guild plans work</span><div class="lp-how">{cells}</div></div>')


def pv_body(width: int) -> str:
    grid = 'display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:16px;align-items:stretch'
    region_a = f'<div style="{grid}">{pv_free_card()}{pv_premium_card()}{pv_example_card()}</div>'
    history = f'<p class="lp-rule">{HISTORY}</p>'
    region_b = (
        '<section id="guild-leaders" style="display:flex;flex-direction:column;gap:4px">'
        f'<h2 class="section-title" style="font-size:18px;margin:0">{PV_GUILD_HEADING}</h2>'
        f'<p style="margin:4px 0 0;font-size:14px;max-width:720px">{PV_GUILD_LEDE}</p>'
        f'<p style="margin:0;font-size:13px;color:#9a9484;max-width:720px">{PV_GUILD_MEMBERS}</p>'
        f'<div style="{grid};margin-top:12px">{pv_guild_card("guild", PV_GUILD_FEATURES)}'
        f'{pv_guild_card("raid_guild", PV_RAID_FEATURES)}{history}</div></section>')
    note = (f'<p class="ann">{MOCK_NOTE} The example image is a labelled stand-in; the real capture replaces it. '
            f'Board width {width}.</p>')
    foot = f'<p style="margin:0;font-size:13px;color:#9a9484">{FOOT}</p>'
    return (f'<div class="lp-col" style="padding:24px {pv_gutter(width)}px 48px;display:flex;flex-direction:column;gap:40px">'
            f'{region_a}{region_b}{pv_how()}{foot}{note}</div>')


def premium_value_page(width: int) -> str:
    return page(CSS + LP_CSS + pv_band(width) + pv_body(width) + footer(), width)


def build_all() -> None:
    boards = [
        ('', plans_page(1440, False)),
        ('-2000', plans_page(2000, False)),
        ('-phone', plans_page(390, True)),
        ('-states', states_page()),
        ('-billing', settings_page(False)),
        ('-billing-phone', settings_page(True)),
        ('-billing-states', billing_states_page()),
        ('-account', account_page()),
        ('-account-phone', account_phone()),
        ('-emails', emails_page()),
    ]
    for suffix, html in boards:
        print(write('launch-pricing' + suffix + '.html', html).as_posix())
    for suffix, width in (('-1024', 1024), ('', 1440), ('-2000', 2000)):
        print(write('premium-value' + suffix + '.html', premium_value_page(width)).as_posix())


if __name__ == '__main__':
    build_all()

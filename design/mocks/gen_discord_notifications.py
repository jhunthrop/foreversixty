"""Discord notifications (design/specs/2026-10-25-discord-notifications.md).

  python3 design/mocks/gen_discord_notifications.py   -> renders/discord-notifications-states.png

One sheet at 1440 wide holding 1440 boards and 390 frames: the right rail of /account (the Plan panel with its
Premium rows, then the Discord block under it), every state of the block, the button states, the Premium row's values,
and the embeds drawn the way Discord renders them (a dark message card; Discord's own look is approximated, it is not
our UI). The MEMBER, the channel id, the report, the sim, every number and date are EXAMPLES. The webhook token is never
drawn: the block shows the channel id segment only, and the invalid-URL message never echoes what was typed.
Helpers (cell, h2, CSS, fonts) come from gen_character_history.py.
"""
import gen_character_header as CH
import gen_character_history as H
import gen_nav as N
from mocklib import RAISED, GOLD, MUTED, TEXT, BORDER, SOFT, SEAL, crest, nav_phone

WARM = '#3a3326'
WIPE = '#f0736b'
CHANNEL = '1098765432101234567'
MONO, BODYF = H.MONO, H.BODYF

W_DESK, W_PH = 395, 354          # the rail's panel width at 1440 (col-span-4) and at 390 (gutters 18)
EVENT_SIM = 'A server sim finishes'
EVENT_UPLOAD = 'A log file you upload finishes parsing'
EVENTS_NOTE = 'Posts every log you upload, including private ones; each message says who can open it.'
UPSELL = 'Discord messages tell you when a server sim or a log you uploaded is done.'
INTRO = 'Get a message in a Discord channel when a server sim finishes or a log finishes parsing.'
HOWTO = 'In Discord: Channel settings, Integrations, Webhooks, New Webhook, Copy Webhook URL.'
TRUST = 'We store the URL encrypted, never show it again, and delete it with your account.'

CSS_RAW = '''<style>
.dn-btn{display:inline-flex;align-items:center;justify-content:center;height:44px;padding:0 16px;border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap;border:1px solid WARM_;color:TEXT_;background:transparent;font-family:BODYF_}
.dn-btn.primary{background:GOLD_;border-color:GOLD_;color:#07090d;padding:0 20px}
.dn-btn.is-hover{border-color:GOLD_;color:#f5d27a}
.dn-btn.primary.is-hover{background:#f5d27a;border-color:#f5d27a;color:#07090d}
.dn-btn.is-active{background:#131824}
.dn-btn.primary.is-active{background:#cfa64a;border-color:#cfa64a}
.dn-btn.is-focus{outline:2px solid GOLD_;outline-offset:2px}
.dn-btn.is-off{opacity:.45}
.dn-link{display:inline-flex;align-items:center;min-height:44px;font-size:13px;color:TEXT_;text-decoration:underline;text-underline-offset:3px;background:none;border:0;padding:0}
.dn-link.is-hover{color:#f5d27a}
.dn-link.is-focus{outline:2px solid GOLD_;outline-offset:2px;border-radius:2px}
.dn-input{height:44px;width:100%;box-sizing:border-box;border:1px solid WARM_;border-radius:4px;background:RAISED_;color:TEXT_;padding:0 12px;font-size:15px;font-family:BODYF_}
.dn-input::placeholder{color:MUTED_}
.dn-input.is-focus{outline:2px solid GOLD_;outline-offset:2px}
.dn-input.is-bad{border-color:WIPE_}
.dn-check{display:flex;align-items:center;gap:12px;min-height:44px;font-size:14px;color:#e9e4d8}
.dn-check input{width:16px;height:16px;margin:0;color-scheme:dark}
.dn-check.is-off{color:MUTED_}
.dn-check.is-focus input{outline:2px solid GOLD_;outline-offset:2px}
.dn-ln{margin:0;min-height:21px;font-size:13px;line-height:1.45;color:MUTED_}
.dn-ln.bad{color:WIPE_}
.dn-ln.ok{color:TEXT_}
</style>'''
CSS = (CSS_RAW.replace('WARM_', WARM).replace('TEXT_', TEXT).replace('GOLD_', GOLD).replace('BODYF_', BODYF).replace('RAISED_', RAISED)
       .replace('MUTED_', MUTED).replace('WIPE_', WIPE))


# ------------------------------------------------------------------ parts
def btn(text, kind='', state=''):
    return f'<span class="dn-btn {kind} {state}">{text}</span>'


def link(text, state=''):
    return f'<span class="dn-link {state}">{text}</span>'


def line(text, kind=''):
    return f'<p class="dn-ln {kind}">{text}</p>'


def check(text, on=True, off=False, state=''):
    box = f'<input type="checkbox" {"checked" if on else ""} {"disabled" if off else ""}>'
    return f'<label class="dn-check {"is-off" if off else ""} {state}">{box}<span>{text}</span></label>'


def label(text):
    return f'<span class="label" style="color:{MUTED}">{text}</span>'


def small(text):
    return f'<p style="margin:0;font-size:13px;color:{MUTED};line-height:1.45">{text}</p>'


def panel(inner, width, phone=False, title='Discord'):
    pad = '16px' if phone else '20px 24px'
    dot = f'<span style="width:8px;height:8px;border-radius:50%;background:{GOLD};box-shadow:0 0 10px rgba(229,185,85,.8)"></span>'
    head = f'<div style="display:flex;align-items:center;gap:10px">{dot}<span class="label" style="color:{GOLD}">{title}</span></div>'
    return (f'<div style="width:{width}px;box-sizing:border-box;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:{pad};'
            f'display:flex;flex-direction:column;gap:16px;font-size:14px;color:{TEXT}">{head}{inner}</div>')


def posts_to():
    return (f'<p style="margin:0;font-size:14px;color:{TEXT};display:flex;flex-wrap:wrap;gap:4px 8px;align-items:baseline">'
            f'<span style="color:{MUTED}">Posts to</span><span>channel</span><span style="font-family:{MONO};font-size:13px;color:#f2eee4">{CHANNEL}</span></p>')


def events(sim=True, upload=True, off=False, note=True):
    n = small(EVENTS_NOTE) if note else ''
    return (f'<div style="display:flex;flex-direction:column;gap:0"><div style="min-height:24px;display:flex;align-items:center">{label("Post when")}</div>'
            f'{check(EVENT_SIM, sim, off)}{check(EVENT_UPLOAD, upload, off)}{n}</div>')


def row(*items):
    return f'<div style="display:flex;flex-wrap:wrap;gap:12px;align-items:center">{"".join(items)}</div>'


def notice(first, second='', accent=WIPE):
    sec = small(second) if second else ''
    return (f'<div style="border-left:2px solid {accent};padding:8px 0 8px 14px;display:flex;flex-direction:column;gap:4px">'
            f'<p style="margin:0;font-size:14px;color:#f2eee4;line-height:1.4">{first}</p>{sec}</div>')


# ------------------------------------------------------------------ the block's states
def paste_form(error='', cancel=False, focus=False, filled=''):
    ph = 'https://discord.com/api/webhooks/…'
    val = f' value="{filled}"' if filled else ''
    bad = 'is-bad' if error else ''
    foc = 'is-focus' if focus else ''
    err = line(error, 'bad') if error else line('')
    cancel_html = link('Cancel') if cancel else ''
    intro = '' if cancel else small(INTRO)
    return (f'{intro}<div style="display:flex;flex-direction:column;gap:8px"><label class="label" style="color:{MUTED}" for="u">Webhook URL</label>'
            f'<input id="u" class="dn-input {bad} {foc}" placeholder="{ph}"{val} autocomplete="off" spellcheck="false">{err}'
            f'{small(HOWTO)}{small(TRUST)}</div>{events()}{row(btn("Save webhook", "primary"), cancel_html)}')


def saved(status='', status_kind='', test_state='', test_label='Send test message', sim=True, upload=True, warn='', confirm=False):
    if confirm:
        act = (f'<div style="display:flex;flex-direction:column;gap:12px"><p style="margin:0;font-size:14px;line-height:1.4">Remove this webhook? Posts stop and we forget the URL.</p>'
               f'{row(btn("Remove"), btn("Cancel"))}</div>')
    else:
        act = row(btn(test_label, '', test_state), btn('Remove'))
    w = small(warn) if warn else ''
    return (f'{posts_to()}{w}{events(sim, upload)}<div style="display:flex;flex-direction:column;gap:6px">{act}{line(status, status_kind)}</div>'
            f'<div style="margin-top:-12px">{link("Use a different webhook")}</div>')


def disabled_off():
    return (f'{posts_to()}'
            f'{notice("Switched off after 5 failed posts in a row.", "Last error from Discord: Unknown Webhook")}'
            f'{small("Sends a test message. If Discord accepts it, posts resume.")}'
            f'{row(btn("Test and switch on", "primary"), btn("Remove"))}'
            f'<div style="margin-top:-8px">{link("Use a different webhook")}</div>{events()}')


def lapsed():
    return (f'{posts_to()}'
            f'{notice("Posts are paused because your plan ended.", "We kept your webhook. Posts start again when you are back on Premium.", GOLD)}'
            f'{row(link("See plans"))}{events(True, True, True, False)}{row(btn("Remove"))}')


def upsell_line():
    a = f'<a style="color:{TEXT};text-decoration:underline;text-underline-offset:2px">See plans</a>'
    return (f'<p style="margin:0;font-size:13px;color:{MUTED};line-height:1.45;display:flex;flex-wrap:wrap;align-items:center;gap:0 6px;min-height:44px">'
            f'{UPSELL} {a}</p>')


# ------------------------------------------------------------------ the Plan panel above it (for placement)
def prem_row(label_, value, link_=False, muted=False):
    color = MUTED if muted else '#e9e4d8'
    deco = 'text-decoration:underline;' if link_ else ''
    return (f'<div style="display:grid;grid-template-columns:10rem minmax(0,1fr);column-gap:12px;align-items:baseline;min-height:28px;font-size:14px">'
            f'<span style="color:{MUTED}">{label_}</span><span style="color:{color};{deco}">{value}</span></div>')


def plan_panel(width, phone, discord_value, link_=True, muted=False):
    pad = '16px' if phone else '20px 24px'
    dot = f'<span style="width:8px;height:8px;border-radius:50%;background:{GOLD};box-shadow:0 0 10px rgba(229,185,85,.8)"></span>'
    head = f'<div style="display:flex;align-items:center;gap:10px">{dot}<span class="label" style="color:{GOLD}">Plan</span></div>'
    rows = (prem_row('Server sims', 'On · 4 runs this month', True) + prem_row('History charts', 'Open on your main character', True)
            + prem_row('Discord notifications', discord_value, link_, muted) + prem_row('Supporter mark', 'Shown on your characters', True))
    plan = f'<p style="margin:0;font-size:14px"><span style="color:{MUTED}">Plan</span> · Premium — Renews 24 Nov 2026</p>'
    return (f'<div style="width:{width}px;box-sizing:border-box;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:{pad};display:flex;flex-direction:column;gap:16px;color:{TEXT}">'
            f'{head}<div style="display:flex;flex-direction:column;gap:12px">{plan}<div>{rows}</div></div></div>')


def rail(width, phone, value, block, link_=True, muted=False):
    return (f'<div style="width:{width}px;display:flex;flex-direction:column;gap:16px">{plan_panel(width, phone, value, link_, muted)}{block}</div>')


# ------------------------------------------------------------------ Discord's rendering of our embeds
DISC_BG, DISC_EMBED = '#313338', '#2b2d31'
DF = "'gg sans','Noto Sans','Helvetica Neue',Arial,sans-serif"


def embed(title, fields, footer, thumb='', desc='', width=520, phone=False):
    fcols = 1 if phone else 3
    flds = ''.join(f'<div style="min-width:0"><div style="font-size:14px;font-weight:600;color:#f2f3f5;line-height:1.3">{n}</div>'
                   f'<div style="font-size:14px;color:#dbdee1;line-height:1.35">{v}</div></div>' for n, v in fields)
    th = (f'<img src="{thumb}" width="80" height="80" style="width:80px;height:80px;border-radius:4px;flex-shrink:0;margin-left:16px;object-fit:cover" alt="">' if thumb else '')
    ds = f'<div style="font-size:14px;color:#dbdee1;line-height:1.4;margin-top:8px">{desc}</div>' if desc else ''
    fgrid = f'<div style="display:grid;grid-template-columns:repeat({fcols},minmax(0,1fr));gap:8px 16px;margin-top:8px">{flds}</div>' if fields else ''
    return (f'<div style="max-width:{width}px;background:{DISC_EMBED};border-left:4px solid {GOLD};border-radius:4px;padding:8px 16px 16px 12px;box-sizing:border-box">'
            f'<div style="display:flex;justify-content:space-between"><div style="min-width:0;flex:1">'
            f'<div style="font-size:16px;font-weight:600;color:#00a8fc;line-height:1.3;margin-top:8px">{title}</div>{ds}{fgrid}</div>{th}</div>'
            f'<div style="font-size:12px;color:#b5bac1;margin-top:12px">{footer}</div></div>')


def message(emb, width, time='Today at 9:41 PM'):
    av = (f'<span style="width:40px;height:40px;border-radius:50%;background:{RAISED};display:grid;place-items:center;flex-shrink:0">'
          f'<img src="{SEAL}" width="28" height="28" alt="" style="width:28px;height:28px"></span>')
    tag = '<span style="background:#5865f2;color:#fff;font-size:10px;font-weight:600;border-radius:3px;padding:1px 5px;margin:0 6px 0 4px;vertical-align:1px">APP</span>'
    return (f'<div style="width:{width}px;box-sizing:border-box;background:{DISC_BG};border-radius:8px;padding:16px;display:flex;gap:16px;font-family:{DF}">{av}'
            f'<div style="min-width:0;flex:1"><div style="line-height:1.3"><span style="font-size:16px;font-weight:500;color:#f2f3f5">Forever Sixty</span>{tag}'
            f'<span style="font-size:12px;color:#949ba4">{time}</span></div><div style="margin-top:4px">{emb}</div></div></div>')


FOOT = 'Forever Sixty · Today at 9:41 PM'


def sim_embed(width, phone=False):
    e = embed('Your Frost Mage sim finished', [('Result', '1,204 DPS'), ('Run', 'Phase 1 BiS check'), ('Took', '3 min 12 s')], FOOT, crest('mage'), width=width - 76, phone=phone)
    return message(e, width)


def upload_embed(width, phone=False):
    e = embed('Tuesday Molten Core finished parsing', [('Zone', 'Molten Core'), ('Fights', '9 fights · 7 kills'), ('Visible to', 'Your guild')], FOOT, width=width - 76, phone=phone)
    return message(e, width)


def test_embed(width, phone=False):
    e = embed('Test message from Forever Sixty', [], FOOT, desc='Your webhook works. You will get a message here when a server sim finishes or a log file you upload finishes parsing.',
              width=width - 76, phone=phone)
    return message(e, width)


# ------------------------------------------------------------------ boards
def d_(title, inner, note='', w=W_DESK):
    return H.cell(title, panel(inner, w), note, w)


def states_desktop():
    bad_url = 'That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.'
    cells = [
        d_('No webhook yet (the paste form)', paste_form(), 'Save is the one filled gold button. Both events start on.'),
        d_('Invalid URL', paste_form(bad_url, focus=True, filled='https://discord.com/channels/…'), 'Message is fixed copy; it never repeats what was typed.'),
        d_('Saved', saved('Saved. Send a test message to see it in your channel.', 'ok'), 'Channel id only; the token is never shown.'),
        d_('Test sending', saved('', '', 'is-off', 'Sending…'), 'Button disabled, aria-busy, label Sending…'),
        d_('Test sent', saved('Sent. Check the channel.', 'ok')),
        d_('Test failed', saved('Discord said: Unknown Webhook.', 'bad'),
           "Discord's own words, shortened to 140 characters, never the URL."),
        d_('Some posts failed, not yet off', saved(warn='The last 3 posts failed. At 5 in a row it switches off. Last error from Discord: Missing Permissions.'),
           'Quiet 13 px line, shown from the first failure.'),
        d_('Switched off after five failures', disabled_off(), 'The reason and the last error come from the API. Primary is the way back.'),
        d_('Plan lapsed (URL kept)', lapsed(), 'Toggles read-only, test hidden, Remove stays.'),
        d_('Remove, asked in place', saved(confirm=True), "The devices list's own confirm grammar."),
        d_('Use a different webhook', paste_form(cancel=True), 'Same form; the saved webhook keeps working until Save succeeds.'),
        d_('Event turned off, saved', saved('Saved.', 'ok', upload=False)),
    ]
    free = H.cell("Free member: the sentence in the block's place (no panel, no form)", upsell_line(), 'One sentence under the Plan panel.', W_DESK)
    return '<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">' + ''.join(cells) + free + '</div>'


def states_buttons():
    states = ['', 'is-hover', 'is-focus', 'is-active', 'is-off']
    names = ['default', 'hover', 'focus-visible', 'active', 'disabled']
    sec = ''.join(H.cell(n, btn('Send test message', '', s), '', 190) for n, s in zip(names, states))
    pri = ''.join(H.cell(n, btn('Save webhook', 'primary', s), '', 170) for n, s in zip(names, states))
    lk = ''.join(H.cell(n, link('Use a different webhook', s), '', 190) for n, s in zip(names[:3], states[:3]))
    ph = 'https://discord.com/api/webhooks/…'
    inp = (H.cell('input default', f'<input class="dn-input" placeholder="{ph}">', '', 300)
           + H.cell('input focus-visible', f'<input class="dn-input is-focus" placeholder="{ph}">', '', 300)
           + H.cell('input error', '<input class="dn-input is-bad" value="https://discord.com/channels/…">', '', 300)
           + H.cell('checkbox row, focus', check(EVENT_SIM, True, False, 'is-focus'), '', 300)
           + H.cell('checkbox row, read-only (lapsed)', check(EVENT_SIM, True, True), '', 300))
    wrap = lambda x: f'<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">{x}</div>'
    return wrap(sec) + wrap(pri) + wrap(lk) + wrap(inp)


def premium_rows():
    items = [
        ('Free', 'Off', False, True),
        ('Payer, nothing saved', 'Off · set up', True, False),
        ('Saved, two events on', 'On · 2 events', True, False),
        ('Saved, one event on', 'On · 1 event', True, False),
        ('Saved, no event chosen', 'Off · no events chosen', True, False),
        ('Switched off', 'Off · failed five times', True, False),
        ('Plan lapsed', 'Paused · plan lapsed', True, False),
    ]
    out = ''
    for name, val, lk, mt in items:
        box = (f'<div style="width:{W_DESK}px;box-sizing:border-box;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:10px 24px">'
               f'{prem_row("Discord notifications", val, lk, mt)}</div>')
        out += H.cell(name, box, '', W_DESK)
    return f'<div style="display:flex;flex-wrap:wrap;gap:16px;align-items:flex-start">{out}</div>'


def embeds():
    d = ('<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">'
         + H.cell('Server sim finished (desktop, 600 px message)', sim_embed(600), 'Title is the link to /sim/ID. The crest is the 80 px thumbnail.', 600)
         + H.cell('Log parsed (desktop)', upload_embed(600), 'Title is the link to /reports/ID. Fields: zone, fights, who can open it.', 600)
         + H.cell('Test message (desktop)', test_embed(600), 'Title links to the account page.', 600) + '</div>')
    p = ('<div style="display:flex;gap:24px;align-items:flex-start">'
         + H.cell('Server sim finished (phone, 354 px)', sim_embed(354, True), 'Fields stack in one column on a phone.', 354)
         + H.cell('Log parsed (phone)', upload_embed(354, True), '', 354)
         + H.cell('Test message (phone)', test_embed(354, True), '', 354) + '</div>')
    return d, p


def phone_frame(inner):
    bar = nav_phone(True, dict([('class', 'mage')]))
    return (f'<div style="width:390px;border:1px solid {SOFT};background:var(--bg);flex-shrink:0">{bar}'
            f'<div style="padding:22px 18px 28px;display:flex;flex-direction:column;gap:32px">{inner}</div></div>')


def sheet():
    gut = CH.gutter(1440)
    wrap = lambda b: f'<div style="max-width:{CH.PAGE_MAX}px;margin:0 auto;padding:36px {gut}px;display:flex;flex-direction:column;gap:24px">{b}</div>'
    saved_rail = rail(W_DESK, False, 'On · 2 events', panel(saved('', ''), W_DESK))
    free_rail = rail(W_DESK, False, 'Off', upsell_line(), False, True)
    off_rail = rail(W_DESK, False, 'Off · failed five times', panel(disabled_off(), W_DESK))
    rails = ('<div style="display:flex;gap:24px;align-items:flex-start">'
             + H.cell('1440: right rail of /account, payer with a saved webhook (Discord directly under Plan)', saved_rail, '', W_DESK)
             + H.cell('1440: free member (the row says Off; one sentence under it)', free_rail, '', W_DESK)
             + H.cell('1440: switched off (row and block agree)', off_rail, '', W_DESK) + '</div>')
    ph_saved = phone_frame(rail(W_PH, True, 'On · 2 events', panel(saved('', ''), W_PH, True)))
    ph_form = phone_frame(rail(W_PH, True, 'Off · set up', panel(paste_form(), W_PH, True)))
    ph_off = phone_frame(rail(W_PH, True, 'Off · failed five times', panel(disabled_off(), W_PH, True)))
    ph_lapsed = phone_frame(rail(W_PH, True, 'Paused · plan lapsed', panel(lapsed(), W_PH, True)))
    ph_free = phone_frame(rail(W_PH, True, 'Off', upsell_line(), False, True))
    phones = ('<div style="display:flex;gap:24px;align-items:flex-start;flex-wrap:wrap">'
              + H.cell('390: saved', ph_saved, '', 392) + H.cell('390: paste form', ph_form, '', 392)
              + H.cell('390: switched off', ph_off, '', 392) + H.cell('390: plan lapsed', ph_lapsed, '', 392)
              + H.cell('390: free member', ph_free, '', 392) + '</div>')
    e_d, e_p = embeds()
    body = (wrap(f'<h1 class="display" style="font-size:22px;color:{TEXT}">Discord notifications: every state</h1>'
                 + H.h2("Placement: the account page's right rail, under the Plan panel") + rails)
            + wrap(H.h2('Block states (395 px, the rail at 1440)') + states_desktop()
                   + H.h2('Controls') + states_buttons()
                   + H.h2("The Premium row's values (Plan panel)") + premium_rows())
            + wrap(H.h2('What Discord shows: the two embeds and the test message') + e_d + H.h2('The same on a phone (354 px message)') + e_p)
            + wrap(H.h2('Phone (390)') + phones))
    return N.doc(CH.CSS + H.CSS + CSS + body, 1440)


if __name__ == '__main__':
    print(H.shoot(sheet(), 'discord-notifications-states', 1440))

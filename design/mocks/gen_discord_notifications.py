"""Discord notifications (design/specs/2026-10-25-discord-notifications.md).

  python3 design/mocks/gen_discord_notifications.py   -> renders/discord-notifications-states.png

One sheet at 1440 wide holding 1440 boards and 390 frames: the right rail of /account (the Plan panel with its
Premium rows, then the Discord block under it), every state of the block, the button states, the Premium row's values,
and the embeds drawn the way Discord renders them (a dark message card; Discord's own look is approximated, it is not
our UI). The MEMBER, the webhook name, the channel id, the report, the sim, every number and date are EXAMPLES. The webhook token
is never drawn: the block shows Discord's webhook name and a shortened channel id only, and the invalid-URL message never echoes
what was typed. Amended 2026-10-25 (follow-up lane): steps, destination name, third checkbox, switched-off notice, warning,
Premium row values, embeds, the loading reserve and gold checkboxes.
Helpers (cell, h2, CSS, fonts) come from gen_character_history.py.
"""
import gen_character_header as CH
import gen_character_history as H
import gen_nav as N
from mocklib import RAISED, GOLD, MUTED, TEXT, BORDER, SOFT, SEAL, crest, nav_phone

WARM = '#3a3326'
WIPE = '#f0736b'
CHANNEL = '1198…4568'          # first four and last four digits of the channel id (spec 3.2)
NAME = 'Raid pings'            # Discord's own name for the webhook
NB = '\u00a0'                  # non-breaking space (spec 3.7)
MONO, BODYF = H.MONO, H.BODYF

W_DESK, W_PH = 395, 354          # the rail's panel width at 1440 (col-span-4) and at 390 (gutters 18)
EVENT_SIM = 'A server sim finishes'
EVENT_UPLOAD = 'A combat log you upload finishes parsing'
EVENT_PRIVATE = 'Include private logs'
EVENTS_NOTE = 'Off, private logs never post. On, everyone in the channel sees the log\'s name and \u201cOnly you\u201d.'
UPSELL = 'Discord messages tell you when a server sim or a log you uploaded is done.'
INTRO = 'Get a message in a Discord channel when a server sim finishes or a log finishes parsing.'
STEPS = ["In Discord, open the channel's settings, then Integrations, then Webhooks.",
         'Choose New Webhook, name it something you will recognize, then Copy Webhook URL.',
         'Paste it below.']
CONNECT = 'Connect a Discord channel'          # (amended 2026-10-25, connect lane)
CONNECT_NOTE = 'Discord asks which server and channel, then lets Forever Sixty post there. We keep only that webhook, encrypted.'
PASTE_INSTEAD = 'Paste a webhook URL instead'
CONNECT_INSTEAD = 'Connect a Discord channel instead'
RECONNECT = 'Connect a different channel'
RECONNECT_NOTE = 'Connecting again replaces this channel and removes its webhook from Discord.'
REMOVE_Q = {'pasted': 'Remove this webhook? Posts stop and we forget the URL. It stays in Discord until you delete it there.',
            'connected': 'Remove this connection? Posts stop and we delete its webhook from Discord.'}
CURE = 'Fixed it in Discord? Test and switch on. If the webhook is gone, make a new one.'
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
.dn-check input{width:16px;height:16px;margin:0;color-scheme:dark;accent-color:GOLD_}
.dn-check input:disabled{accent-color:MUTED_}
.dn-steps{margin:0;padding-left:20px;font-size:13px;line-height:1.45;color:MUTED_}
.dn-steps li{margin:0 0 4px}
.dn-steps li:last-child{margin-bottom:0}
.dn-ln.under{padding-top:4px;box-sizing:border-box}
.sk{border-radius:6px;background:#121826;border:1px dashed #262e40;display:grid;place-items:center;color:#9a9484;font-size:12px;box-sizing:border-box}
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


def nb(value):
    # Spec 3.7: U+00A0 after the state word and between a count and its noun.
    return value.replace(' \u00b7 ', NB + '\u00b7 ').replace('2 events', f'2{NB}events').replace('1 event', f'1{NB}event')


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


def posts_to(name=NAME, source='pasted'):
    if source == 'connected':
        one = f'<span style="color:{MUTED}">Connected to</span> <span style="color:#f2eee4;font-weight:600">a Discord channel</span>'
    elif name:
        one = f'<span style="color:{MUTED}">Posts through</span> <span style="color:#f2eee4;font-weight:600">\u201c{name}\u201d</span>'
    else:
        one = f'<span style="color:{MUTED}">Posts to a Discord channel</span>'
    return (f'<div style="display:flex;flex-direction:column;gap:2px"><p style="margin:0;font-size:14px;color:{TEXT};overflow-wrap:anywhere">{one}</p>'
            f'<p style="margin:0;font-size:13px;color:{MUTED}">Channel <span style="font-family:{MONO};font-size:13px;color:{TEXT}">{CHANNEL}</span></p></div>')


def events(sim=True, upload=True, off=False, note=True, private=False):
    n = small(EVENTS_NOTE) if note else ''
    third_off = off or not upload
    return (f'<div style="display:flex;flex-direction:column;gap:0"><div style="min-height:24px;display:flex;align-items:center">{label("Post when")}</div>'
            f'{check(EVENT_SIM, sim, off)}{check(EVENT_UPLOAD, upload, off)}'
            f'<div style="padding-left:28px">{check(EVENT_PRIVATE, private, third_off)}</div>{n}</div>')


def row(*items):
    return f'<div style="display:flex;flex-wrap:wrap;gap:12px;align-items:center">{"".join(items)}</div>'


def notice(first, second='', accent=WIPE):
    sec = small(second) if second else ''
    return (f'<div style="border-left:2px solid {accent};padding:8px 0 8px 14px;display:flex;flex-direction:column;gap:4px">'
            f'<p style="margin:0;font-size:14px;color:#f2eee4;line-height:1.4">{first}</p>{sec}</div>')


# ------------------------------------------------------------------ the block's states
def steps():
    return '<ol class="dn-steps">' + ''.join(f'<li>{t}</li>' for t in STEPS) + '</ol>'


def paste_form(error='', cancel=False, focus=False, filled='', back=False):
    ph = 'https://discord.com/api/webhooks/…'
    val = f' value="{filled}"' if filled else ''
    bad = 'is-bad' if error else ''
    foc = 'is-focus' if focus else ''
    err = line(error, 'bad under') if error else line('', 'under')
    cancel_html = link('Cancel') if cancel else ''
    intro = '' if cancel else small(INTRO)
    backlink = f'<div style="margin-top:-12px">{link(CONNECT_INSTEAD)}</div>' if back else ''
    return (f'{intro}{steps()}<div style="display:flex;flex-direction:column;gap:0"><label class="label" style="color:{MUTED};margin-bottom:8px" for="u">Webhook URL</label>'
            f'<input id="u" class="dn-input {bad} {foc}" placeholder="{ph}"{val} autocomplete="off" spellcheck="false">{err}'
            f'{small(TRUST)}</div>{events()}{row(btn("Save webhook", "primary"), cancel_html)}{backlink}')


def connect_view(status='', status_kind='', replace=False, btn_label=CONNECT, btn_state=''):
    # Spec 3.2 Connect state: button, one explanation, reserved status line, then the paste disclosure.
    intro = '' if replace else small(INTRO)
    cancel = f'<div style="margin-top:-12px">{link("Cancel")}</div>' if replace else ''
    return (f'{intro}<div style="display:flex;flex-direction:column;gap:8px">{row(btn(btn_label, "primary", btn_state))}{small(CONNECT_NOTE)}'
            f'{line(status, status_kind)}{link(PASTE_INSTEAD)}</div>{cancel}')


def warning(lead, detail):
    return (f'<div style="display:flex;flex-direction:column;gap:2px"><p style="margin:0;font-size:13px;color:{TEXT};line-height:1.45">{lead}</p>'
            f'{small(detail)}</div>')


def saved(status='', status_kind='', test_state='', test_label='Send test message', sim=True, upload=True, warn=None, confirm=False,
          private=False, name=NAME, source='pasted'):
    if confirm:
        act = (f'<div style="display:flex;flex-direction:column;gap:12px"><p style="margin:0;font-size:14px;line-height:1.4">{REMOVE_Q[source]}</p>'
               f'{row(btn("Remove"), btn("Cancel"))}</div>')
    else:
        act = row(btn(test_label, '', test_state), btn('Remove'))
    w = warning(*warn) if warn else ''
    if confirm:
        swap = ''
    elif source == 'connected':
        swap = f'<div style="margin-top:-12px">{link(RECONNECT)}{small(RECONNECT_NOTE)}</div>'
    else:
        swap = f'<div style="margin-top:-12px">{link("Use a different webhook")}</div>'
    return (f'{posts_to(name, source)}{w}{events(sim, upload, private=private)}<div style="display:flex;flex-direction:column;gap:6px">{act}{line(status, status_kind)}</div>'
            f'{swap}')


def disabled_off(error='Last error from Discord: Unknown Webhook.'):
    err = small(error) if error else ''
    box = (f'<div style="border-left:2px solid {WIPE};padding:8px 0 0 14px;display:flex;flex-direction:column;gap:4px">'
           f'<p style="margin:0;font-size:14px;color:#f2eee4;line-height:1.4">Switched off after 5 failed posts in a row.</p>{err}{small(CURE)}'
           f'{link("Use a different webhook")}</div>')
    return (f'{posts_to()}{box}'
            f'<div style="display:flex;flex-direction:column;gap:6px">{row(btn("Test and switch on", "primary"), btn("Remove"))}{line("")}</div>'
            f'{events()}')


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
            f'<span style="color:{MUTED}">{label_}</span><span style="color:{color};{deco}">{nb(value) if label_.startswith("Discord") else value}</span></div>')


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


def sim_embed(width, phone=False, named=True):
    title = "Frostbolt's Frost Mage sim finished" if named else 'Your Frost Mage sim finished'
    e = embed(title, [('Result', '1,204 DPS'), ('Run', 'Phase 1 BiS check'), ('Run time', '3 min 12 s')], FOOT, crest('mage'), width=width - 76, phone=phone)
    return message(e, width)


def upload_embed(width, phone=False):
    e = embed('Tuesday Molten Core finished parsing', [('Zone', 'Molten Core'), ('Boss fights', '9 · 7 kills'), ('Visible to', 'Your guild')], FOOT, width=width - 76, phone=phone)
    return message(e, width)


def test_embed(width, phone=False, private=False):
    tail = 'finishes parsing, private logs included.' if private else 'finishes parsing.'
    e = embed('Test message from Forever Sixty', [], FOOT, desc=f'Your webhook works. You will get a message here when a server sim finishes or a combat log you upload {tail}',
              width=width - 76, phone=phone)
    return message(e, width)


# ------------------------------------------------------------------ boards
def d_(title, inner, note='', w=W_DESK):
    return H.cell(title, panel(inner, w), note, w)


def states_desktop():
    bad_url = 'That is not a Discord webhook URL. It starts with https://discord.com/api/webhooks/.'
    cells = [
        d_('No webhook yet (the paste form)', paste_form(), 'Save is the one filled gold button. Sim and log start on; private logs start off.'),
        d_('Invalid URL', paste_form(bad_url, focus=True, filled='https://discord.com/channels/…'), 'Message is fixed copy; it never repeats what was typed.'),
        d_('Saved', saved('Saved. Send a test message to see it in your channel.', 'ok'), "Discord's webhook name and a shortened channel id; the token is never shown."),
        d_('Test sending', saved('', '', 'is-off', 'Sending…'), 'Button disabled, aria-busy, label Sending…'),
        d_('Test sent', saved('Sent. Check the channel.', 'ok')),
        d_('Test failed', saved('Discord said: Unknown Webhook.', 'bad'),
           "Discord's own words, shortened to 140 characters, never the URL."),
        d_('Some posts failed, not yet off', saved(warn=('The last 3 posts failed. At 5 in a row it switches off. Send a test message to check.',
                                                          'Last error from Discord: Missing Permissions.')),
           'Two paragraphs: the facts and the next step in --text, the error in --muted; no bar.'),
        d_('Switched off after five failures', disabled_off(), 'Cause, cure and the way to a new webhook sit in one notice; Test and switch on is the one gold fill.'),
        d_('Switched off, 140-character error', disabled_off('Last error from Discord: Invalid Webhook Token: the token in this webhook URL was changed, or the webhook was deleted from the channel settings by someone with\u2026'),
           'Cut at a word, one ellipsis, no full stop after it.'),
        d_('Saved, no webhook name (rows saved before the name column)', saved(name=None), 'Line one falls back to Posts to a Discord channel.'),
        d_('Saved, private logs on', saved(private=True), 'The third box is a modifier of the second: indented 28 px, off by default.'),
        d_('Log off: private row disabled, value kept', saved(upload=False, private=True), 'A disabled box takes no focus; its value is kept.'),
        d_('Plan lapsed (URL kept)', lapsed(), 'Toggles read-only, test hidden, Remove stays.'),
        d_('Remove, asked in place', saved(confirm=True), "The devices list's own confirm grammar. Use a different webhook is not drawn while the question is open."),
        d_('Use a different webhook', paste_form(cancel=True), 'Same form; the saved webhook keeps working until Save succeeds.'),
        d_('Event turned off, saved', saved('Saved.', 'ok', upload=False)),
    ]
    cells += connect_boards()
    free = H.cell("Free member: the sentence in the block's place (no panel, no form)", upsell_line(), 'One sentence under the Plan panel.', W_DESK)
    return '<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">' + ''.join(cells) + free + '</div>'


def connect_boards():
    # (amended 2026-10-25, connect lane) every state the Connect flow adds; the paste-first boards above are the connect=false state.
    return [
        d_('CONNECT ON: nothing saved (the new first state)', connect_view(), 'One gold button, one explanation, the paste path as a quiet link. No steps, no field, no checkboxes: defaults apply and the member edits them once connected.'),
        d_('Connect: button states, left to right in the capture: hover, focus, opening', connect_view(btn_state='is-hover', btn_label=CONNECT) + connect_view(replace=True, btn_state='is-focus') + connect_view(replace=True, btn_label='Opening Discord\u2026', btn_state='is-off'),
           'Hover gold-hover, focus 2 px ring, then Opening Discord\u2026 (aria-busy, not clickable) in the instant before the redirect.'),
        d_('Back from Discord: connected', saved('Connected. Send a test message to see it in your channel.', 'ok', source='connected'), 'Focus lands on Send test message. Line one never shows the Discord-side webhook name (it is "Forever Sixty").'),
        d_('Back from Discord: cancelled', connect_view('Cancelled in Discord. Nothing was connected.', 'ok'), 'Plain text colour, role=status; focus returns to Connect a Discord channel.'),
        d_('Back from Discord: failed', connect_view('Could not connect. Check that you can manage webhooks in that channel, then try again.', 'bad'), 'role=alert, --wipe. We know only that it failed.'),
        d_('Connect on, paste view (disclosed by Paste a webhook URL instead)', paste_form(back=True), 'The same form as today, plus one link back. The view swaps; nothing is hidden behind an accordion.'),
        d_('Saved, connected (reconnect in one click)', saved(source='connected'), 'Connected to a Discord channel + the shortened channel id. The link goes straight to Discord; its note says the old webhook is removed from Discord.'),
        d_('Remove, connected', saved(confirm=True, source='connected'), 'We also delete the webhook from Discord.'),
        d_('Remove, pasted (connect on)', saved(confirm=True, source='pasted'), 'We only forget it; the member made it, so the member deletes it.'),
        d_('Use a different webhook, pasted + connect on', connect_view(replace=True), 'Replace view for a pasted webhook is the connect view with Cancel; the old webhook keeps posting until a new one is saved.'),
        d_('Reconnect cancelled (connected, saved)', saved('Cancelled in Discord. Nothing was connected.', 'ok', source='connected'), 'The saved webhook is untouched and keeps posting.'),
    ]


def states_buttons():
    states = ['', 'is-hover', 'is-focus', 'is-active', 'is-off']
    names = ['default', 'hover', 'focus-visible', 'active', 'disabled']
    sec = ''.join(H.cell(n, btn('Send test message', '', s), '', 190) for n, s in zip(names, states))
    pri = ''.join(H.cell(n, btn('Save webhook', 'primary', s), '', 170) for n, s in zip(names, states))
    lk = ''.join(H.cell(n, link('Use a different webhook', s), '', 190) for n, s in zip(names[:3], states[:3]))
    ph = 'https://discord.com/api/webhooks/…'
    chk = ''.join(H.cell(nm, f'<div style="display:flex;align-items:center;gap:12px;min-height:44px;color:{TEXT};font-size:14px">'
                              f'<input type="checkbox" {a} style="width:16px;height:16px;margin:0;color-scheme:dark;accent-color:{GOLD}">{t}</div>', '', 190)
                  for nm, a, t in (('unchecked', '', 'Show a pseudonym'), ('checked (gold)', 'checked', 'Show a pseudonym'),
                                   ('disabled, checked', 'checked disabled', 'Read-only'), ('disabled, unchecked', 'disabled', 'Read-only')))
    inp = (H.cell('input default', f'<input class="dn-input" placeholder="{ph}">', '', 300)
           + H.cell('input focus-visible', f'<input class="dn-input is-focus" placeholder="{ph}">', '', 300)
           + H.cell('input error', '<input class="dn-input is-bad" value="https://discord.com/channels/…">', '', 300)
           + H.cell('checkbox row, focus', check(EVENT_SIM, True, False, 'is-focus'), '', 300)
           + H.cell('checkbox row, read-only (lapsed)', check(EVENT_SIM, True, True), '', 300))
    wrap = lambda x: f'<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">{x}</div>'
    return wrap(sec) + wrap(pri) + wrap(lk) + wrap(inp) + wrap(chk)


def premium_rows():
    items = [
        ('Free', 'Off · Premium', False, True),
        ('Payer, nothing saved', 'Off · set up', True, False),
        ('Saved, two events on', 'On · 2 events', True, False),
        ('Saved, one event on', 'On · 1 event', True, False),
        ('Saved, no event chosen', 'Off · no events chosen', True, False),
        ('Switched off', 'Off · needs attention', True, False),
        ('Plan lapsed', 'Paused · plan lapsed', True, False),
    ]
    out = ''
    for name, val, lk, mt in items:
        box = (f'<div style="width:{W_DESK}px;box-sizing:border-box;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:10px 24px">'
               f'{prem_row("Discord notifications", val, lk, mt)}</div>')
        out += H.cell(name, box, '', W_DESK)
    narrow = ''
    for name, val in (('1024 rail (288 px): the count never splits from its noun', 'On · 2 events'),
                      ('1024 rail (288 px): a wrap falls between ordinary words', 'Off · needs attention')):
        box = (f'<div style="width:288px;box-sizing:border-box;border:1px solid {BORDER};border-radius:6px;background:{RAISED};padding:10px 24px">'
               f'<div style="display:grid;grid-template-columns:7.5rem minmax(0,1fr);column-gap:12px;align-items:baseline;min-height:28px;font-size:14px">'
               f'<span style="color:{MUTED}">Discord notifications</span><span style="color:#e9e4d8;text-decoration:underline">{nb(val)}</span></div></div>')
        narrow += H.cell(name, box, '', 288)
    return (f'<div style="display:flex;flex-wrap:wrap;gap:16px;align-items:flex-start">{out}</div>'
            f'<div style="display:flex;flex-wrap:wrap;gap:16px;align-items:flex-start;margin-top:16px">{narrow}</div>')


def embeds():
    d = ('<div style="display:flex;flex-wrap:wrap;gap:24px;align-items:flex-start">'
         + H.cell('Server sim finished (desktop, 600 px message)', sim_embed(600), 'Title is the link to /sim/ID. The crest is the 80 px thumbnail.', 600)
         + H.cell('Server sim finished, no character name on the run', sim_embed(600, named=False), 'Falls back to Your Frost Mage sim finished.', 600)
         + H.cell('Log parsed (desktop)', upload_embed(600), 'Title is the link to /reports/ID. Fields: zone, boss fights (boss pulls only), who can open it.', 600)
         + H.cell('Test message (desktop)', test_embed(600), 'Title links to the account page.', 600)
         + H.cell('Test message, private logs on', test_embed(600, private=True), 'The clause is added only when Include private logs is on.', 600) + '</div>')
    p = ('<div style="display:flex;gap:24px;align-items:flex-start">'
         + H.cell('Server sim finished (phone, 354 px)', sim_embed(354, True), 'Fields stack in one column on a phone.', 354)
         + H.cell('Log parsed (phone)', upload_embed(354, True), '', 354)
         + H.cell('Test message (phone)', test_embed(354, True), '', 354) + '</div>')
    return d, p


def sk(height, width, text=''):
    return f'<div class="sk" style="height:{height}px;width:{width}px">{text}</div>'


DISCORD_FLOORS = (('none', 0, 0), ('free', 44, 44), ('connect', 320, 340), ('form', 610, 630), ('saved', 460, 475), ('off', 540, 565), ('lapsed', 475, 475))


def loading_rail(width, discord_h, phone=False):
    # Spec 9.1: the rail's loading skeleton. Heights are the 1440 floors; the phone floors are measured at build time.
    gap = 'display:flex;flex-direction:column'
    plan_wrap = (f'<div style="{gap};gap:16px">{sk(219, width, "Plan  219")}'
                 + (sk(discord_h, width, f"Discord  {discord_h}") if discord_h else '') + '</div>')
    return (f'<div style="{gap};gap:32px;width:{width}px">{sk(232, width, "Devices  232")}{sk(192, width, "You  192")}{plan_wrap}</div>')


def skeleton_boards():
    ready = rail(W_DESK, False, 'On · 2 events', panel(saved('', ''), W_DESK))
    pair = ('<div style="display:flex;gap:24px;align-items:flex-start">'
            + H.cell('1440: loading (saved state, cached /v1/me): the rail at its real positions', loading_rail(W_DESK, 460), 'Same width, same gaps (32, 32, 16). Devices, You and Plan sit under the ready panels; the footer does not move.', W_DESK)
            + H.cell('1440: ready (Plan and Discord shown; Devices and You omitted here)', ready, '', W_DESK) + '</div>')
    states = ''.join(H.cell(f'data-discord="{k}"  {a} / {b} px', sk(a, 200, f'{a} px') if a else '<div style="height:20px;color:#9a9484;font-size:12px">draws nothing</div>', '', 200)
                     for k, a, b in DISCORD_FLOORS)
    return pair + f'<div style="display:flex;gap:16px;align-items:flex-start;flex-wrap:wrap;margin-top:16px">{states}</div>'


def phone_frame(inner):
    bar = nav_phone(True, dict([('class', 'mage')]))
    return (f'<div style="width:390px;border:1px solid {SOFT};background:var(--bg);flex-shrink:0">{bar}'
            f'<div style="padding:22px 18px 28px;display:flex;flex-direction:column;gap:32px">{inner}</div></div>')


def sheet():
    gut = CH.gutter(1440)
    wrap = lambda b: f'<div style="max-width:{CH.PAGE_MAX}px;margin:0 auto;padding:36px {gut}px;display:flex;flex-direction:column;gap:24px">{b}</div>'
    saved_rail = rail(W_DESK, False, 'On · 2 events', panel(saved('', ''), W_DESK))
    free_rail = rail(W_DESK, False, 'Off · Premium', upsell_line(), False, True)
    off_rail = rail(W_DESK, False, 'Off · needs attention', panel(disabled_off(), W_DESK))
    rails = ('<div style="display:flex;gap:24px;align-items:flex-start">'
             + H.cell('1440: right rail of /account, payer with a saved webhook (Discord directly under Plan)', saved_rail, '', W_DESK)
             + H.cell('1440: free member (the row says Off · Premium; one sentence under it)', free_rail, '', W_DESK)
             + H.cell('1440: switched off (row and block agree)', off_rail, '', W_DESK) + '</div>')
    ph_saved = phone_frame(rail(W_PH, True, 'On · 2 events', panel(saved('', ''), W_PH, True)))
    ph_form = phone_frame(rail(W_PH, True, 'Off · set up', panel(paste_form(), W_PH, True)))
    ph_connect = phone_frame(rail(W_PH, True, 'Off · set up', panel(connect_view(), W_PH, True)))
    ph_conn_saved = phone_frame(rail(W_PH, True, 'On · 2 events', panel(saved(source='connected'), W_PH, True)))
    ph_off = phone_frame(rail(W_PH, True, 'Off · needs attention', panel(disabled_off(), W_PH, True)))
    ph_lapsed = phone_frame(rail(W_PH, True, 'Paused · plan lapsed', panel(lapsed(), W_PH, True)))
    ph_free = phone_frame(rail(W_PH, True, 'Off · Premium', upsell_line(), False, True))
    phones = ('<div style="display:flex;gap:24px;align-items:flex-start;flex-wrap:wrap">'
              + H.cell('390: saved', ph_saved, '', 392) + H.cell('390: connect (connect on)', ph_connect, '', 392) + H.cell('390: connected, saved', ph_conn_saved, '', 392)
              + H.cell('390: paste form (connect off)', ph_form, '', 392)
              + H.cell('390: switched off', ph_off, '', 392) + H.cell('390: plan lapsed', ph_lapsed, '', 392)
              + H.cell('390: free member', ph_free, '', 392) + '</div>')
    e_d, e_p = embeds()
    body = (wrap(f'<h1 class="display" style="font-size:22px;color:{TEXT}">Discord notifications: every state</h1>'
                 + H.h2("Placement: the account page's right rail, under the Plan panel") + rails)
            + wrap(H.h2('Block states (395 px, the rail at 1440)') + states_desktop()
                   + H.h2('Controls') + states_buttons()
                   + H.h2("The Premium row's values (Plan panel)") + premium_rows()
                   + H.h2('The account page loading reserve (spec 9.1): the rail skeleton, and the Discord floor by state (1440 / 390)') + skeleton_boards())
            + wrap(H.h2('What Discord shows: the two embeds and the test message') + e_d + H.h2('The same on a phone (354 px message)') + e_p)
            + wrap(H.h2('Phone (390)') + phones))
    return N.doc(CH.CSS + H.CSS + CSS + body, 1440)


if __name__ == '__main__':
    print(H.shoot(sheet(), 'discord-notifications-states', 1440))

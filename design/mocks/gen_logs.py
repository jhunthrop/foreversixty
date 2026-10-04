"""Mock boards for the /logs landing rebuild (design/specs/2026-10-04-logs-landing.md).

  python3 design/mocks/gen_logs.py            -> renders/logs-signed-in.html (1440)
  LOGS_STATE=signed-out python3 ...           -> renders/logs-signed-out.html (1440)
  LOGS_PHONE=1 python3 ...                    -> renders/logs-phone.html (390, signed in)

Numbers come from design/mocks/data (the real public sample report, refreshed by
fetch_logs_data.py). On the signed-in board that report's own numbers stand in for the
visitor's own newest report; the board caption says so.
"""
import json, os, html
from mocklib import (DATA, GOLD, MUTED, TEXT, BODY, BORDER, SOFT, GREEN, CLASS_COLOR, crest, emblem,
                     nav, nav_phone, footer, page, write)

PHONE = os.environ.get('LOGS_PHONE') == '1'
STATE = os.environ.get('LOGS_STATE', 'signed-in')
SIGNED_IN = STATE == 'signed-in'
L = json.load(open(DATA / 'logs-sample-report.json'))
R = L['report']
ZULMARA = {'class': 'hunter', 'faction': 'horde', 'battletag': 'Zulmara'}

hero_title = 'Sanguine Depths' if SIGNED_IN else R['title']
hero_day = R['created_at'][:10]
top = L['dps'][0]
fight = L['fight']
HUNTER = CLASS_COLOR['hunter']


def btn(label, href, gold=False, fill=False, h=44):
    cls = 'btn' + (' btn-gold' if gold else '') + (' btn-fill' if fill else '')
    return f'<a class="{cls}" href="{href}" style="height:{h}px">{label}</a>'


def tab(label, active=False):
    st = f'color:{GOLD};border-color:{GOLD};background:#e5b95514' if active else f'color:{TEXT};border-color:{BORDER}'
    return f'<span style="display:inline-flex;align-items:center;height:44px;padding:0 14px;border:1px solid;border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;white-space:nowrap;{st}">{label}</span>'


def mono(s, color=MUTED, size=14):
    return f'<span class="mono" style="font-size:{size}px;color:{color}">{s}</span>'


sample_pill = '' if SIGNED_IN else '<span class="pill pill-sample" style="margin-left:10px;vertical-align:4px">Sample</span>'
facts = mono(f'{hero_day}&nbsp;·&nbsp;{R["fight_count"]} fights&nbsp;·&nbsp;{R["kill_count"]} kills')
desc = (f'<span style="font-size:14px;color:{BODY};max-width:62ch">Every pull, ranked: damage, healing, deaths, buffs, casts and threat. '
        f'<a href="#companion" style="color:{BODY};text-decoration:underline">Log live with the companion</a> or '
        f'<a href="#upload" style="color:{BODY};text-decoration:underline">upload a log file</a>.</span>')
dps_text = f"{top['dps']:,} DPS"
if SIGNED_IN:
    hook = (f'<span style="font-size:13px;color:{MUTED};display:flex;align-items:center;gap:10px;flex-wrap:wrap">'
            f'<span>Last kill, {html.escape(fight["name"])}:</span>'
            f'<span class="display" style="font-size:14px;font-weight:700;color:{CLASS_COLOR.get(top["class"].lower(), TEXT)}">{top["name"]}</span>'
            f'{mono(dps_text, TEXT, 13)}'
            f'<a href="#planner" style="font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase">Gear in the planner</a>'
            f'<span style="color:{SOFT}">|</span><a href="#sim" style="font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase">Sim</a></span>')
else:
    hook = f'<span style="font-size:13px;color:{MUTED}">Open it to see every player\'s gear next to the planner and the simulator.</span>'

hero = f'''<span class="label" style="color:{GOLD};display:flex;align-items:center;gap:10px"><i style="width:28px;height:1px;background:{GOLD};display:inline-block"></i>Logs</span>
<h1 class="display" style="font-size:22px;line-height:1.1;font-weight:700;color:{TEXT}">{hero_title}{sample_pill}</h1>
<span>{facts}</span>
{desc}
{hook}
<span>{btn('Open this report', '#report', gold=True)}</span>'''

if SIGNED_IN:
    doors = ''.join(f'<a href="#{d.lower()}" style="font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:{GOLD if d == "Logs" else TEXT};{"border-bottom:2px solid " + GOLD if d == "Logs" else ""}">{d}</a>' for d in ['Planner', 'Simulator', 'Logs', 'Rankings'])
    spine = (f'<div style="display:flex;align-items:center;gap:18px;flex-wrap:wrap;min-height:44px;padding:8px 0;border-bottom:1px solid {SOFT}">'
             f'<span style="display:flex;align-items:center;gap:10px"><img class="crest" src="{crest("hunter")}" alt="" style="width:28px;height:28px;--c:{HUNTER}">'
             f'<span class="display" style="font-size:14px;font-weight:700;color:{HUNTER}">Zulmara</span>'
             f'<span style="font-size:12px;color:{MUTED}"><img src="{emblem("horde")}" alt="" style="width:14px;height:14px;object-fit:contain;vertical-align:-2px;margin-right:4px">Level 24 Troll Hunter</span></span>'
             f'<span style="display:flex;gap:16px">{doors}</span>'
             f'<span style="margin-left:auto;display:flex;gap:14px"><a href="#switch" style="font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:{MUTED}">Switch</a>'
             f'<a href="#forget" style="font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:{MUTED}">Forget</a></span></div>')
else:
    spine = (f'<div style="display:flex;align-items:center;gap:14px;flex-wrap:wrap;min-height:44px;padding:8px 0;border-bottom:1px solid {SOFT};font-size:13px;color:{MUTED}">'
             f'Sign in with Battle.net or paste an export to point the site at your character. <a href="#signin" style="text-decoration:underline">Sign in with Battle.net</a> <a href="#paste" style="text-decoration:underline">Paste an export</a></div>')


def panel(title, body, aside=''):
    a = f'<span style="font-size:12px;color:{MUTED}">{aside}</span>' if aside else ''
    return (f'<section class="panel" style="display:flex;flex-direction:column;gap:12px;padding:16px 20px">'
            f'<div style="display:flex;justify-content:space-between;align-items:baseline"><span class="display" style="font-size:15px;font-weight:700;color:{TEXT};letter-spacing:.02em">{title}</span>{a}</div>{body}</section>')


def report_row(title, day, fights, kills, guild=None):
    g = f'<span style="font-size:12px;color:{MUTED}">{guild}</span>' if guild else ''
    return (f'<div style="display:grid;grid-template-columns:1fr auto;gap:8px 16px;align-items:baseline;padding:10px 0;border-bottom:1px solid {SOFT}">'
            f'<span style="display:flex;flex-direction:column;gap:2px;min-width:0"><a href="#r" style="font-weight:600;color:{TEXT}">{title}</a>{g}</span>'
            f'{mono(f"{day} · {fights} fights · {kills} kills", MUTED, 13)}</div>')


if SIGNED_IN:
    tabs = f'<div style="display:flex;gap:6px">{tab("Mine", True)}{tab("Olympus XXVII")}</div>'
    rows = report_row('Sanguine Depths', hero_day, R['fight_count'], R['kill_count']) + report_row('Deadmines, first clear', '2026-09-28', 24, 6) + report_row('Wailing Caverns', '2026-09-24', 31, 7)
    mine = panel('Your reports', tabs + f'<div>{rows}</div><span style="font-size:12px;color:{MUTED}">Newest first. A guild tab shows the reports your guild shared with its members.</span>')
else:
    mine = panel('Your reports', f'<span style="font-size:14px;color:{BODY}">Sign in to see the reports you own. <a href="#email">Use an email link instead.</a></span><span>{btn("Sign in with Battle.net", "#signin", fill=True)}</span>')
recent = panel('Recent public reports', f'<span style="font-size:14px;color:{BODY}">No public reports yet. The first raid logs land in December; dungeon logs are welcome now.</span>')

status = (f'<div style="display:flex;align-items:center;gap:10px;font-size:13px;color:{BODY}"><i style="width:8px;height:8px;border-radius:50%;background:{GREEN};display:inline-block"></i>MacBook Pro · macOS · last seen {mono("4 min ago", BODY, 13)}</div>' if SIGNED_IN else '')
pointer = f'<span style="font-size:14px;color:{BODY}">Downloads and the in-game /combatlog step are on the setup page. <a href="#setup" style="text-decoration:underline">Get set up</a></span>'
if SIGNED_IN:
    pairing = (f'<div style="display:flex;flex-direction:column;gap:10px;padding-top:12px;border-top:1px solid {SOFT}"><span class="label">Pair another device</span>'
               f'<div style="display:flex;align-items:center;gap:16px;flex-wrap:wrap"><span class="mono" style="font-size:26px;letter-spacing:.18em;color:{TEXT};padding:8px 14px;border:1px dashed {BORDER};border-radius:4px">7K3Q-P2ND</span>'
               f'<span style="font-size:13px;color:{MUTED}">Type this into the companion within {mono("9 minutes", MUTED, 13)}. This panel confirms the moment it pairs.</span></div></div>')
else:
    pairing = (f'<div style="display:flex;flex-direction:column;gap:10px;padding-top:12px;border-top:1px solid {SOFT}"><span class="label">Pair it with your account</span>'
               f'<span style="font-size:14px;color:{BODY}">Sign in to pair the companion with your account. <a href="#signin" style="text-decoration:underline;color:{BODY}">Sign in with Battle.net</a> or <a href="#email">use an email link</a>.</span></div>')
companion = panel('The companion', status + pointer + pairing, aside='Live, while you raid')

vis = ''.join(f'<span style="display:inline-flex;align-items:center;height:36px;padding:0 12px;border:1px solid {GOLD if v == "Public" else BORDER};border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:{GOLD if v == "Public" else TEXT}">{v}</span>' for v in ['Public', 'Unlisted', 'Guild', 'Private'])
upload_body = (f'<span style="font-size:14px;color:{BODY}">A whole {mono("WoWCombatLog.txt", BODY, 13)}. The first fight is readable within ten seconds of the upload finishing, and the rest appear as they are parsed.</span>'
               f'<span style="font-size:13px;color:{MUTED}">The game writes it to its {mono("Logs", MUTED, 12)} folder, beside {mono("Interface", MUTED, 12)} and {mono("WTF", MUTED, 12)}. Up to 4 GB. Type {mono("/combatlog", MUTED, 12)} again to stop logging before you upload, so the file is whole.</span>')
if SIGNED_IN:
    upload_body += (f'<div style="display:flex;align-items:center;justify-content:center;gap:6px;flex-wrap:wrap;min-height:96px;border:1px dashed {BORDER};border-radius:6px;color:{MUTED};font-size:13px;text-align:center">Drop {mono("WoWCombatLog.txt", TEXT, 13)} here, or <span style="color:{GOLD};font-weight:700;letter-spacing:.06em;text-transform:uppercase;font-size:12px">choose a file</span></div>'
                    f'<span class="label">Title</span><span style="height:40px;border:1px solid {BORDER};border-radius:4px;padding:0 12px;display:flex;align-items:center;color:{MUTED};font-size:14px">Optional</span>'
                    f'<span class="label">Who can open it</span><div style="display:flex;gap:6px;flex-wrap:wrap">{vis}</div><span style="font-size:12px;color:{MUTED}">Listed, ranked, anyone can open it.</span>'
                    f'<span>{btn("Upload", "#go", gold=True, h=40)}</span>')
else:
    upload_body += f'<span style="font-size:14px;color:{BODY}">Sign in to upload. The report is filed under your account. <a href="#signin" style="text-decoration:underline;color:{BODY}">Sign in with Battle.net</a> or <a href="#email">use an email link</a>.</span>'
upload = panel('Upload a log', upload_body, aside='For a night already logged')
framing = f'<span style="font-size:13px;color:{MUTED}">Logs are for group content at any level: a dungeon run logs the same way a raid does.</span>'

if not PHONE:
    body = f'''<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav('Logs', SIGNED_IN, ZULMARA)}
<div style="display:flex;flex-direction:column;gap:14px;padding:28px 48px 32px 48px;max-width:1344px">{hero}</div></div>
<div style="display:flex;flex-direction:column;gap:24px;padding:0 48px 32px 48px">
{spine}{framing}
<div style="display:grid;grid-template-columns:7fr 5fr;gap:24px;align-items:start">{mine}{recent}</div>
<div style="display:grid;grid-template-columns:7fr 5fr;gap:24px;align-items:start">{companion}{upload}</div>
</div>{footer()}'''
    out = write(f'logs-{STATE}.html', page(body, 1440))
else:
    body = f'''<div class="band" style="display:flex;flex-direction:column;flex-shrink:0">{nav_phone(SIGNED_IN, ZULMARA)}
<div style="display:flex;flex-direction:column;gap:12px;padding:18px 18px 22px 18px">{hero}</div></div>
<div style="display:flex;flex-direction:column;gap:16px;padding:0 18px 24px 18px">{spine}{mine}{recent}{companion}{upload}{framing}</div>{footer(phone=True)}'''
    out = write(f'logs-phone-{STATE}.html', page(body, 390))
print(out)

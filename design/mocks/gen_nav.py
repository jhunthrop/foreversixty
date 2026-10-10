"""Mock boards for the site navigation bar and its character selector
(design/specs/2026-10-09-nav-character-selector.md).

  python3 design/mocks/gen_nav.py   -> renders/nav*.png (see BOARDS)

Real data: class colours, the circular HD crests and the faction emblems come from mocklib (the
shipped site assets); spec names come from data/curated/specs.json. The characters in the list are
EXAMPLES (invented names, realistic class/spec pairs); the board labels them as such. The HTML is
written to a scratch directory and only the PNGs land in renders/. Screenshots run from web/ so
Playwright is found, at each board's own viewport width.
"""
import json
import subprocess
import sys
import tempfile
from pathlib import Path

from mocklib import (
    ROOT, WEB, RENDERS, GOLD, MUTED, TEXT, BODY, CLASS_COLOR, HEAD, SEAL, crest, emblem,
)

SPECS = json.load(open(ROOT / 'data/curated/specs.json'))
SPEC_NAMES = {s['spec']: s['name'] for s in SPECS}
CLASS_NAME = {s['class_slug']: s['class_slug'].title() for s in SPECS}

WIPE = '#f0736b'      # --color-wipe: the site's failure colour
EMBER = '#d66e28'     # --color-ember: stale
KILL = '#7fd48a'

DOORS = ['Planner', 'BiS', 'Simulator', 'Logs', 'Rankings', 'Tier List', 'Guides', 'Get set up']
CURRENT_DOOR = 'Planner'


def example(name, spec, level, realm, faction, source, age, ruleset='Normal', status='ok', **flags):
    cls = spec.split('-')[0]
    return {'name': name, 'cls': cls, 'specname': SPEC_NAMES[spec], 'spec': f"{SPEC_NAMES[spec]} {CLASS_NAME[cls]}", 'level': level, 'realm': realm,
            'faction': faction, 'source': source, 'age': age, 'ruleset': ruleset, 'status': status, **flags}


# Ordered the way the selector orders them: current first, then by newest sync. EXAMPLES, not accounts.
CURRENT = example('Obnoxious Yell', 'warrior-fury', 60, 'Living Flame', 'alliance', 'Addon', '12 minutes ago', current=True)
STALE = example('Treewalker', 'druid-balance', 60, 'Living Flame', 'alliance', 'Addon', '19 days ago', status='stale')
BNET = example('Frostbyte', 'mage-frost', 42, 'Stonehearth', 'horde', 'Battle.net', '3 days ago', ruleset='Hardcore')
FAILED = example('Oakheart', 'paladin-protection', 60, 'Living Flame', 'alliance', 'Battle.net', '3 days ago', status='failed')
PASTED = example('Shadowmend', 'priest-shadow', 35, 'Living Flame', 'horde', 'Pasted export', '2 days ago')
QUICK = example('Quickshot', 'hunter-beast-mastery', 60, 'Living Flame', 'horde', 'Addon', '5 days ago')
NIGHT = example('Nightsnare', 'rogue-combat', 60, 'Stonehearth', 'horde', 'Addon', '6 days ago')
CIND = example('Cindermaw', 'warlock-affliction', 51, 'Living Flame', 'horde', 'Addon', '8 days ago')
# Ordered the way the selector orders them: the current one, then the rest by last chosen on this
# browser (most recent first), then newest sync for characters never chosen here.
DAWN = example('Dawnbringer', 'paladin-retribution', 60, 'Stonehearth', 'alliance', 'Addon', '12 days ago')
# Chosen on this browser, most recent first: Frostbyte (yesterday), Shadowmend (3 days ago). Never chosen here, newest sync first:
# Oakheart 3 d, Quickshot 5 d, Nightsnare 6 d, Cindermaw 8 d, Dawnbringer 12 d, Treewalker 19 d.
MANY = [CURRENT, BNET, PASTED, FAILED, QUICK, STALE]                 # six rows: fits the list with no scroll
MANY9 = [CURRENT, BNET, PASTED, FAILED, QUICK, NIGHT, CIND, DAWN, STALE]   # nine: filter row, the list scrolls under a fade

CSS = f'''<style>
body{{font-size:14px}}
.hd{{display:flex;align-items:center;justify-content:space-between;gap:24px;padding:18px 48px;border-bottom:1px solid #262e4099;background:var(--bg);position:relative;z-index:5}}
.wm{{display:inline-flex;align-items:center;gap:8px;min-height:44px;flex-shrink:0}}
.wm img{{width:30px;height:30px}}.wm span{{font-family:var(--font-display);font-weight:700;font-size:20px;letter-spacing:.1em;text-transform:uppercase;color:var(--strong);white-space:nowrap}}
.doors{{display:flex;gap:22px;font-size:13px;font-weight:700;letter-spacing:.12em;text-transform:uppercase;white-space:nowrap}}
.doors a{{color:#b9b3a4}}.doors a.cur{{color:var(--gold)}}
.right{{display:flex;align-items:center;gap:12px;flex-shrink:0}}
.slot{{position:relative}}
.sel{{display:flex;align-items:center;gap:10px;height:44px;width:232px;padding:0 12px 0 6px;border:1px solid var(--line);background:var(--raised);border-radius:4px;color:var(--strong);font:inherit;cursor:pointer;text-decoration:none}}
.sel.hover{{border-color:#4a4030;background:#101624}}.sel.hover .sch{{color:var(--strong)}}
.sel.focus{{outline:2px solid var(--gold);outline-offset:2px}}
.sel.open{{border-color:var(--gold);background:#101624}}.sel.open .sch{{transform:rotate(180deg);color:var(--strong)}}
.sel.pre .sch,.sel.pre .badge{{visibility:hidden}}
.cw{{position:relative;width:32px;height:32px;flex-shrink:0}}
.cw .crest{{width:32px;height:32px}}
.mk{{position:absolute;top:-3px;right:-3px;width:12px;height:12px;border-radius:999px;border:2px solid var(--mc);background:var(--bg)}}
.stx{{display:flex;flex:1;flex-direction:column;min-width:0;max-width:168px;line-height:1.15;text-align:left}}
.snm{{font-size:14px;font-weight:600;white-space:nowrap;overflow:hidden;text-overflow:ellipsis}}
.ssp{{font-size:12px;color:var(--muted);white-space:nowrap;overflow:hidden;text-overflow:ellipsis}}
.pn{{display:none}}.ssl{{display:none}}
.sch{{color:var(--muted);flex-shrink:0}}
.ghost{{width:32px;height:32px;border-radius:999px;border:1.5px dashed #5a5546;display:grid;place-items:center;color:var(--muted);flex-shrink:0}}
.badge{{display:none}}
.dc{{display:inline-flex;align-items:center;gap:8px;height:44px;padding:0 6px;color:#b9b3a4;font-weight:600;font-size:13px;position:relative}}
.dc .t{{display:inline}}
.menu{{display:none;background:none;font:inherit;width:44px;height:44px;align-items:center;justify-content:center;border:1px solid #3a3326;border-radius:4px;color:var(--text)}}
.menu.on{{border-color:var(--gold)}}
.tip{{position:absolute;z-index:30;background:#161c2a;border:1px solid #3a3326;border-radius:4px;padding:6px 9px;font-size:12px;color:var(--text);white-space:nowrap;box-shadow:0 6px 18px #000a}}
.pop{{position:absolute;right:0;top:calc(100% + 10px);width:392px;background:var(--raised);border:1px solid #3a3326;border-radius:6px;box-shadow:0 18px 40px #000000b3;z-index:20;text-align:left}}
.ph{{display:flex;align-items:center;justify-content:space-between;height:44px;padding:0 14px;border-bottom:1px solid var(--line-soft)}}
.ph a{{font-size:12px;font-weight:600;color:var(--muted);text-decoration:underline;text-decoration-color:#9a948466;text-underline-offset:4px}}
.list{{list-style:none;margin:0;padding:4px 0;max-height:440px;overflow:hidden;position:relative}}
.list.scroll{{max-height:476px}}
.list.scroll::after{{content:"";position:absolute;left:0;right:0;bottom:0;height:36px;background:linear-gradient(180deg,#0d111a00,#0d111a)}}
.row{{display:grid;grid-template-columns:36px 1fr auto;align-items:center;gap:12px;min-height:64px;padding:8px 14px;border-left:2px solid transparent;position:relative}}
.row.cur{{background:#e5b95510;border-left-color:var(--gold)}}
.row.hover{{background:#101624;border-left-color:var(--c)}}
.row.focus{{outline:2px solid var(--gold);outline-offset:-2px;background:#101624}}
.row .crest{{width:36px;height:36px}}
.rt{{display:flex;flex-direction:column;min-width:0;gap:1px}}
.r1{{display:flex;align-items:center;gap:7px;font-size:15px;font-weight:600;line-height:1.2}}
.fm{{display:inline-block;width:14px;height:14px;flex-shrink:0;background:var(--fc);-webkit-mask:var(--fe) center/contain no-repeat;mask:var(--fe) center/contain no-repeat}}
.r2{{font-size:13px;color:{BODY};white-space:nowrap;overflow:hidden;text-overflow:ellipsis}}
.r3{{font-size:12px;color:var(--muted);white-space:nowrap;overflow:hidden;text-overflow:ellipsis}}
.r3 em{{font-style:normal}}.r3 .stale{{color:{EMBER}}}.r3.failed{{color:{WIPE}}}
.chk{{color:var(--gold)}}
.retry{{height:44px;min-width:56px;padding:0 10px;border:1px solid #3a3326;border-radius:4px;color:var(--strong);font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;display:inline-flex;align-items:center;justify-content:center}}
.act{{display:grid;grid-template-columns:36px 1fr;gap:12px;align-items:center;min-height:52px;padding:6px 14px;border-top:1px solid var(--line-soft)}}
.act .ic{{width:36px;height:36px;border-radius:4px;border:1px solid var(--line);display:grid;place-items:center;color:#b9b3a4}}
.act b{{display:block;font-size:14px;font-weight:600;color:var(--strong)}}.act small{{font-size:12px;color:var(--muted)}}
.act.quiet b{{font-size:13px;color:#b9b3a4}}
.note{{padding:10px 14px;font-size:12px;color:var(--muted);border-top:1px solid var(--line-soft)}}
.sk{{background:linear-gradient(90deg,#161c2a,#1c2433,#161c2a);border-radius:3px}}
.body{{padding:36px 48px;max-width:1440px;margin:0 auto;width:100%}}
.btnn{{display:flex;align-items:center;justify-content:center;height:44px;border-radius:4px;border:1px solid #4a4030;color:var(--strong);font-weight:700;font-size:13px;letter-spacing:.06em;text-transform:uppercase}}
.ta{{height:76px;border:1px solid var(--line);border-radius:4px;background:#07090d;padding:8px 10px;color:var(--muted);font-family:var(--font-mono);font-size:12px}}
.cell{{display:flex;flex-direction:column;gap:10px}}
.cell .stage{{position:relative;min-height:60px}}
.cell .pop{{position:static;width:100%;box-shadow:none}}
.mp{{padding:12px 18px 14px;border-bottom:1px solid var(--line-soft);background:var(--bg)}}
.mchar{{display:grid;grid-template-columns:36px 1fr auto;gap:12px;align-items:center;min-height:56px;padding:6px 12px 6px 8px;border:1px solid var(--line);border-radius:6px;background:var(--raised);margin-bottom:12px}}
.mchar .crest{{width:36px;height:36px}}
.mgrid{{display:grid;grid-template-columns:repeat(3,1fr);gap:2px 12px;font-size:13px;font-weight:700;letter-spacing:.12em;text-transform:uppercase}}
.mgrid a{{display:flex;align-items:center;justify-content:center;min-height:44px;color:#b9b3a4;white-space:nowrap}}.mgrid a.cur{{color:var(--gold)}}
.mdc{{display:flex;align-items:center;justify-content:center;gap:8px;height:44px;margin-top:8px;border:1px solid #3a3326;border-radius:4px;font-weight:600;font-size:13px;color:var(--text)}}
@media (min-width:1441px){{.hd{{padding-left:calc((100% - 1344px)/2);padding-right:calc((100% - 1344px)/2)}}}}
@media (max-width:1439px){{
 .dc .t{{display:none}}.dc{{width:44px;justify-content:center;padding:0}}
 .hd{{gap:18px;padding:18px 32px}}.doors{{gap:20px;letter-spacing:.1em}}
}}
@media (max-width:1279px){{
 .hd{{gap:12px;padding:18px 20px}}.wm span{{font-size:17px;letter-spacing:.08em}}.wm img{{width:28px;height:28px}}.wm{{gap:7px}}
 .doors{{gap:14px;font-size:12px;letter-spacing:.08em}}.right{{gap:4px}}.dc{{width:40px}}
 .sel{{width:176px;padding:0 8px 0 5px;gap:6px}}.stx{{max-width:112px}}.ssp{{display:none}}.ssl{{display:block;font-size:12px;color:var(--muted)}}
 .sel.so{{padding:0 10px 0 8px;gap:6px;width:122px}}.sel.so .ghost{{width:24px;height:24px}}.sel.so .ghost svg{{width:13px;height:13px}}.sel.so .stx{{max-width:none}}
}}
@media (max-width:1099px){{.doors{{letter-spacing:.06em;gap:13px}}}}
@media (max-width:1023px){{
 .hd{{display:grid;grid-template-columns:minmax(0,1fr) auto auto;gap:8px;padding:6px 18px;height:56px}}
 .doors,.dc{{display:none}}.wm span{{font-size:16px;letter-spacing:.1em}}.wm img{{width:28px;height:28px}}.wm{{gap:8px}}
 .right{{display:contents}}
 .sel,.sel.so{{width:auto;min-width:44px;max-width:none;padding:0;gap:4px;justify-content:center;border:0;background:none;position:relative}}.sel .sch{{display:none}}
 .sel .stx{{display:none}}.sel .cw{{width:36px;height:36px}}.sel .cw .crest{{width:36px;height:36px}}.sel.so .ghost{{display:grid}}
 .pn{{display:none}}
  .badge{{display:grid;position:absolute;left:25px;bottom:3px;width:14px;height:14px;border-radius:999px;background:var(--bg);border:1px solid #3a3326;place-items:center;color:var(--strong)}}
 .menu{{display:inline-flex}}
}}
@media (min-width:560px) and (max-width:1023px){{.sel,.sel.so{{gap:10px;padding:0 6px 0 4px}}.badge{{display:none}}.sel .sch{{display:block}}.sel.so .sch{{display:none}}.pn{{display:block;font-size:14px;font-weight:600;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;max-width:min(200px,calc(100vw - 460px))}}}}
</style>'''

CHEV = '<svg class="sch" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.4" stroke-linecap="round" stroke-linejoin="round"><path d="M6 9l6 6 6-6"/></svg>'
CHECK = '<svg class="chk" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.6" stroke-linecap="round" stroke-linejoin="round"><path d="M5 12.5l4.5 4.5L19 7.5"/></svg>'
USER = '<svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><circle cx="12" cy="8" r="4"/><path d="M4 21c0-4 3.6-7 8-7s8 3 8 7"/></svg>'
# The real Discord mark (Simple Icons path, 24 px grid), filled with currentColor.
DISCORD = ('<svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true"><path d="M20.317 4.3698a19.7913 19.7913 0 00-4.8851-1.5152.0741.0741 0 00-.0785.0371c-.211.3753-.4447.8648-.6083 1.2495-1.8447-.2762-3.68-.2762-5.4868 0-.1636-.3933-.4058-.8742-.6177-1.2495a.077.077 0 00-.0785-.037 19.7363 19.7363 0 00-4.8852 1.515.0699.0699 0 00-.0321.0277C.5334 9.0458-.319 13.5799.0992 18.0578a.0824.0824 0 00.0312.0561c2.0528 1.5076 4.0413 2.4228 5.9929 3.0294a.0777.0777 0 00.0842-.0276c.4616-.6304.8731-1.2952 1.226-1.9942a.076.076 0 00-.0416-.1057c-.6528-.2476-1.2743-.5495-1.8722-.8923a.077.077 0 01-.0076-.1277c.1258-.0943.2517-.1923.3718-.2914a.0743.0743 0 01.0776-.0105c3.9278 1.7933 8.18 1.7933 12.0614 0a.0739.0739 0 01.0785.0095c.1202.099.246.1981.3728.2924a.077.077 0 01-.0066.1276 12.2986 12.2986 0 01-1.873.8914.0766.0766 0 00-.0407.1067c.3604.698.7719 1.3628 1.225 1.9932a.076.076 0 00.0842.0286c1.961-.6067 3.9495-1.5219 6.0023-3.0294a.077.077 0 00.0313-.0552c.5004-5.177-.8382-9.6739-3.5485-13.6604a.061.061 0 00-.0312-.0286zM8.02 15.3312c-1.1825 0-2.1569-1.0857-2.1569-2.419 0-1.3332.9555-2.4189 2.157-2.4189 1.2108 0 2.1757 1.0952 2.1568 2.419 0 1.3332-.9555 2.4189-2.1569 2.4189zm7.9748 0c-1.1825 0-2.1569-1.0857-2.1569-2.419 0-1.3332.9555-2.4189 2.1569-2.4189 1.2108 0 2.1757 1.0952 2.1568 2.419 0 1.3332-.946 2.4189-2.1568 2.4189Z"/></svg>')
MENU = '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M4 7h16M4 12h16M4 17h16"/></svg>'
CLOSE = '<svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"><path d="M6 6l12 12M18 6L6 18"/></svg>'
CLIP = '<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="6" y="4" width="12" height="17" rx="2"/><path d="M9 4h6v3H9zM9 12h6M9 16h4"/></svg>'
DL = '<svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 4v11M7 11l5 5 5-5M5 20h14"/></svg>'

FACTION_COLOR = {'alliance': '#6fb1ff', 'horde': '#ff6b5c'}     # --color-alliance, --color-horde
MARK_COLOR = {'stale': EMBER, 'failed': WIPE, 'session': EMBER}
MARK_TIP = {'stale': 'Last synced {age}. Log in to the game to update.',
            'failed': 'Battle.net refresh failed. Showing the last good data.',
            'session': 'Signed out. Sign in to see your other characters.'}


# ---------------------------------------------------------------------------- bar
def aria(ch):
    mark = {'stale': ' Sync is stale.', 'failed': ' Sync failed.', 'session': ' Signed out.'}.get(ch.get('mark'), '')
    return f"{ch['name']}, {ch['spec']}, level {ch['level']}.{mark} Change character."


def selector(ch, state='', tag='button', tip=False):
    """The closed selector. `ch` None is the signed-out slot. state: '', 'hover', 'focus', 'open', 'pre'."""
    if ch is None:
        return (f'<{tag} class="sel so {state}" aria-label="Sign in or choose a character"><span class="ghost">{USER}</span>'
                f'<span class="stx"><span class="snm">Sign in</span><span class="ssp">or paste an export</span></span>{CHEV}</{tag}>')
    color = CLASS_COLOR[ch['cls']]
    mark = ch.get('mark')
    mk = f'<i class="mk" style="--mc:{MARK_COLOR[mark]}"></i>' if mark else ''
    tipx = (f'<span class="tip" style="top:calc(100% + 6px);left:0">{MARK_TIP[mark].format(age=ch["age"])}</span>' if (mark and tip) else '')
    chev = '<svg width="8" height="8" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="4" stroke-linecap="round" stroke-linejoin="round"><path d="M6 9l6 6 6-6"/></svg>'
    return (f'<{tag} class="sel {state}" aria-label="{aria(ch)}" style="position:relative">'
            f'<span class="cw"><img class="crest" src="{crest(ch["cls"])}" alt="" style="--c:{color}">{mk}<span class="badge">{chev}</span></span>'
            f'<span class="pn" style="color:{color}">{ch["name"]}</span>'
            f'<span class="stx"><span class="snm" style="color:{color}">{ch["name"]}</span><span class="ssp">{ch["spec"]} · {ch["level"]}</span><span class="ssl">{ch["specname"]} · {ch["level"]}</span></span>'
            f'{CHEV}{tipx}</{tag}>')


def bar(ch, state='', popover='', tip=False, menu_on=False, dc_tip=False):
    doors = ''.join(f'<a href="#{d}" class="{"cur" if d == CURRENT_DOOR else ""}">{d}</a>' for d in DOORS)
    dtip = '<span class="tip" style="top:calc(100% + 4px);right:0">Discord</span>' if dc_tip else ''
    return (f'<header class="hd"><a class="wm" href="#home"><img src="{SEAL}" alt=""><span>Forever Sixty</span></a>'
            f'<nav class="doors" aria-label="Primary">{doors}</nav>'
            f'<div class="right"><div class="slot">{selector(ch, state, "a" if state == "pre" else "button", tip)}{popover}</div>'
            f'<a class="dc" href="#discord" aria-label="Discord">{DISCORD}<span class="t">Discord</span>{dtip}</a>'
            f'<button class="menu {"on" if menu_on else ""}" aria-label="Menu">{MENU}</button></div></header>')


# ---------------------------------------------------------------------------- list
def fmark(faction):
    return f'<i class="fm" style="--fc:{FACTION_COLOR[faction]};--fe:url({emblem(faction)})"></i>'


def row(ch, flag=''):
    color = CLASS_COLOR[ch['cls']]
    where = ch['realm'] if ch['ruleset'] == 'Normal' else f"{ch['realm']} · {ch['ruleset']}"
    if ch['status'] == 'stale':
        tail_txt = 'Log in to the game to update' if ch['source'] != 'Battle.net' else 'Refresh from your account'
        r3 = f'<span class="r3">{ch["source"]} · {ch["age"]} · <span class="stale">{tail_txt}</span></span>'
    elif ch['status'] == 'failed':
        r3 = f'<span class="r3 failed">Battle.net refresh failed · last good {ch["age"]}</span>'
    else:
        r3 = f'<span class="r3">{ch["source"]} · {ch["age"]}</span>'
    tail = CHECK if ch.get('current') else ('<span class="retry">Retry</span>' if ch['status'] == 'failed' else '<span></span>')
    cls = ' '.join(x for x in ['row', 'cur' if ch.get('current') else '', flag] if x)
    return (f'<li class="{cls}" style="--c:{color}"><img class="crest" src="{crest(ch["cls"])}" alt="" style="--c:{color}">'
            f'<span class="rt"><span class="r1" style="color:{color}">{ch["name"]}{fmark(ch["faction"])}</span>'
            f'<span class="r2">{ch["spec"]} · {ch["level"]} · {where}</span>{r3}</span>{tail}</li>')


def action(icon, title, sub, quiet=False):
    return f'<div class="act{" quiet" if quiet else ""}"><span class="ic">{icon}</span><span><b>{title}</b><small>{sub}</small></span></div>'


PASTE = action(CLIP, 'Paste an export', 'From the addon. Adds a character to this list.')
GETADDON = action(DL, 'Get the addon', 'Export a character from the game.')
SIGNIN = action(USER, 'Sign in', 'Email or Battle.net. Keeps your characters.', quiet=True)


def head(title='Your characters', manage=True):
    return f'<div class="ph"><span class="label">{title}</span>{"<a href=#account>Manage</a>" if manage else ""}</div>'


def pop(inner):
    return f'<div class="pop" role="dialog" aria-label="Choose a character">{inner}</div>'


def skeleton_rows(n=3):
    one = ('<li class="row" aria-hidden="true"><span class="sk" style="width:36px;height:36px;border-radius:999px"></span><span class="rt">'
           '<span class="sk" style="width:120px;height:14px"></span><span class="sk" style="width:210px;height:12px;margin-top:5px"></span>'
           '<span class="sk" style="width:150px;height:11px;margin-top:5px"></span></span><span></span></li>')
    return one * n


def msg(title, body, button=''):
    b = f'<div class="btnn" style="margin-top:12px">{button}</div>' if button else ''
    return f'<div style="padding:16px 14px"><div style="font-weight:600;color:var(--strong);font-size:14px">{title}</div><div style="color:{BODY};font-size:13px;margin-top:4px">{body}</div>{b}</div>'


def rows_html(rows, flags=None):
    flags = flags or {}
    return ''.join(row(r, flags.get(r['name'], '')) for r in rows)


def pop_many(rows=None, flags=None):
    rows = rows or MANY
    scroll = len(rows) > 6
    filt = ('<div style="padding:8px 14px;border-bottom:1px solid var(--line-soft)"><div class="ta" style="height:36px;padding:8px 10px;font-family:var(--font-body);font-size:13px">Filter by name</div></div>' if len(rows) > 8 else '')
    return pop(head() + filt + f'<ul class="list {"scroll" if scroll else ""}" role="list">{rows_html(rows, flags)}</ul>' + PASTE)


def pop_one():
    return pop(head() + f'<ul class="list">{row(CURRENT)}</ul>' + PASTE)


def pop_none():
    return pop(head() + msg('No characters yet.', 'Paste an export from the addon to add one.') + PASTE + GETADDON)


def pop_loading():
    return pop(head() + f'<ul class="list" aria-busy="true">{skeleton_rows(3)}</ul>' + PASTE)


def pop_error():
    return pop(head() + msg('Your characters did not load.', 'Check your connection and try again. You can still paste an export.', 'Try again') + PASTE)


def pop_paste():
    form = (f'<div style="padding:12px 14px 14px;border-top:1px solid var(--line-soft)"><div class="label" style="margin-bottom:8px">Addon export</div>'
            f'<div class="ta">Paste the export here</div><div style="display:flex;gap:8px;margin-top:10px"><div class="btnn" style="flex:1">Use this character</div>'
            f'<span class="retry" style="min-width:84px">Cancel</span></div></div>')
    return pop(head() + f'<ul class="list">{row(CURRENT)}</ul>' + form)


def pop_signed_out():
    return pop(head('Add your character', manage=False) + PASTE + GETADDON + SIGNIN)


def pop_signed_out_pasted():
    return pop(head() + f'<ul class="list">{row(PASTED_CUR)}</ul><div class="note">Not saved to an account. Sign in to keep it.</div>' + PASTE + SIGNIN)


def pop_expired():
    return pop(head() + f'<ul class="list">{row(CURRENT)}</ul><div class="note" style="color:{EMBER}">Your session ended. Sign in to see your other characters.</div>' + PASTE + SIGNIN)


# ---------------------------------------------------------------------------- pages
def body_stub(pad='36px 48px'):
    sk = lambda w, h=12, mt=0: f'<div class="sk" style="width:{w};height:{h}px;margin-top:{mt}px"></div>'
    return (f'<div class="body" style="padding:{pad}"><div class="label" style="color:var(--gold)">Planner</div>'
            f'<div class="display" style="font-size:30px;font-weight:700;color:var(--strong);margin:6px 0 14px">Planner</div>'
            f'<div class="panel" style="padding:18px;display:grid;gap:10px;opacity:.7">{sk("62%")}{sk("48%")}{sk("55%")}</div>'
            f'<div style="margin-top:14px;font-size:12px;color:var(--muted)">Page content stands in. Example characters only: the planner, simulator, BiS and tier list read the selected one.</div></div>')


def doc(body, width, height=None):
    h = f'height:{height}px;overflow:hidden;' if height else 'min-height:100vh;'
    return HEAD.replace('</head>', CSS + '</head>') + \
        f'<div style="width:{width}px;{h}position:relative;display:flex;flex-direction:column;background:var(--bg)">{body}</div></body></html>'


def desktop_page(width, ch=CURRENT, state='', popover='', height=None, **kw):
    return doc(bar(ch, state, popover, **kw) + body_stub(), width, height or (700 if popover else 360))


def sheet_html(inner_rows, extra_after='', title='Your characters', top=True):
    return (f'<div style="position:absolute;inset:0;background:#000000a6;z-index:10"></div>'
            f'<div role="dialog" aria-label="Choose a character" style="position:absolute;left:0;right:0;bottom:0;z-index:11;background:var(--raised);border:1px solid #3a3326;border-bottom:0;border-radius:12px 12px 0 0;max-height:86%;display:flex;flex-direction:column">'
            f'<div style="display:grid;place-items:center;height:20px"><span style="width:36px;height:4px;border-radius:2px;background:#3a3326"></span></div>'
            f'<div class="ph" style="height:48px"><span class="label">{title}</span><span style="width:44px;height:44px;display:grid;place-items:center;color:var(--text);margin-right:-12px">{CLOSE}</span></div>'
            f'{inner_rows}{extra_after}<div style="height:calc(14px + env(safe-area-inset-bottom))"></div></div>')


def phone_page(ch=CURRENT, sheet=False, menu=False):
    b = bar(ch, '', menu_on=menu)
    extra = ''
    if menu:
        if ch is None:
            top = (f'<div class="mchar"><span class="ghost" style="width:36px;height:36px">{USER}</span><span><div style="font-size:15px;font-weight:600;color:var(--strong)">Sign in</div>'
                   f'<div style="font-size:12px;color:var(--muted)">or paste an export</div></span>{CHEV}</div>')
        else:
            color = CLASS_COLOR[ch['cls']]
            top = (f'<div class="mchar"><img class="crest" src="{crest(ch["cls"])}" alt="" style="--c:{color}"><span><div style="font-size:15px;font-weight:600;color:{color}">{ch["name"]}</div>'
                   f'<div style="font-size:12px;color:var(--muted)">{ch["spec"]} · {ch["level"]}</div></span>{CHEV}</div>')
        grid = ''.join(f'<a href="#{d}" class="{"cur" if d == CURRENT_DOOR else ""}">{d}</a>' for d in DOORS)
        extra = f'<div class="mp">{top}<div class="mgrid">{grid}</div><div class="mdc">{DISCORD}Discord</div></div>'
    sh = ''
    if sheet:
        if ch is None:
            sh = sheet_html(PASTE + GETADDON + SIGNIN, title='Add your character')
        else:
            sh = sheet_html(f'<ul class="list" style="max-height:none;overflow:hidden;flex:1">{rows_html(MANY)}</ul>', PASTE)
    return doc(b + extra + body_stub('24px 18px') + sh, 390, 844)


def states_page():
    def cell(title, inner, note=''):
        return f'<div class="cell"><div class="label">{title}</div><div class="stage">{inner}</div><div style="font-size:12px;color:var(--muted)">{note}</div></div>'
    sel = lambda ch, st='', tip=False: f'<div style="display:inline-block;position:relative">{selector(ch, st, tip=tip)}</div>'
    cur_stale = {**CURRENT, 'mark': 'stale', 'age': '19 days ago', 'status': 'stale'}
    cur_failed = {**CURRENT, 'mark': 'failed', 'age': '3 days ago', 'status': 'failed'}
    cur_sess = {**CURRENT, 'mark': 'session'}
    dc = f'<div style="display:inline-block;position:relative"><span class="dc" style="width:40px;justify-content:center">{DISCORD}<span class="tip" style="top:calc(100% + 4px);left:-6px">Discord</span></span></div>'
    cells = [
        cell('Closed, default', sel(CURRENT), 'Border #262e40, chevron muted.'),
        cell('Closed, hover', sel(CURRENT, 'hover'), 'Border #4a4030, fill #101624, chevron to strong.'),
        cell('Closed, focus-visible', sel(CURRENT, 'focus'), '2px gold outline, 2px offset.'),
        cell('Closed, open', sel(CURRENT, 'open'), 'Gold border, chevron up, aria-expanded.'),
        cell('Closed, current is stale', sel(cur_stale, tip=True) + '<div style="height:34px"></div>', 'Ember ring mark on the crest; tooltip on hover and focus. The name never changes colour.'),
        cell('Closed, current sync failed', sel(cur_failed, tip=True) + '<div style="height:34px"></div>', 'Red ring mark, same place.'),
        cell('Closed, session expired', sel(cur_sess, tip=True) + '<div style="height:34px"></div>', 'The pointer stays; the ring mark says why the list is short.'),
        cell('Discord below 1440', dc + '<div style="height:30px"></div>', 'Real Discord mark; the label moves to a tooltip and the aria-label.'),
        cell('Row, hover', pop(f'<ul class="list">{row(PASTED, "hover")}</ul>'), '2px class-colour edge, fill #101624.'),
        cell('Row, focus-visible', pop(f'<ul class="list">{row(BNET, "focus")}</ul>'), 'Gold ring inside the row; arrow keys move it.'),
        cell('Stale sync (past the rule)', pop(f'<ul class="list">{row(STALE)}</ul>'), 'Plain age always; ember only on the words past it, never on the name.'),
        cell('Failed sync', pop(f'<ul class="list">{row(FAILED)}</ul>'), 'Row still selectable; Retry is its own 44 px button.'),
        cell('Signed in, one character', pop_one(), 'Still opens: the paste row lives here.'),
        cell('Signed in, none', pop_none(), 'Closed slot shows the signed-out button instead of a crest.'),
        cell('Loading', pop_loading(), 'Three 64 px skeleton rows; the closed slot keeps its width.'),
        cell('List did not load', pop_error(), 'Paste still works with no list.'),
        cell('Paste an export, open', pop_paste(), 'Inline: the list stays above, focus lands in the box.'),
    ]
    grid = '<div style="display:grid;grid-template-columns:repeat(3,392px);gap:28px 36px;padding:36px 48px">' + ''.join(cells) + '</div>'
    intro = ('<div style="padding:28px 48px 0"><div class="label" style="color:var(--gold)">Nav character selector</div>'
             '<div class="display" style="font-size:24px;font-weight:700;color:var(--strong);margin-top:4px">States</div>'
             '<div style="font-size:12px;color:var(--muted);margin-top:6px">Example characters only. Every cell is the real markup at the bar\'s 1440 width.</div></div>')
    return doc(intro + grid, 1440)


def prehydration_page():
    lab = lambda t: f'<div class="label" style="padding:18px 48px 8px;color:var(--gold)">{t}</div>'
    note = ('<div style="padding:0 48px 24px;font-size:12px;color:var(--muted);max-width:900px">Before hydration the slot is a link to /account#characters. '
            'Base.astro\'s pre-paint script has already set the crest (from data-pointer-class) and written the name and spec from the stored pointer, so the box, '
            'the crest and the text are where they will be. Hydration adds the chevron and the popover; nothing moves.</div>')
    return doc(lab('Desktop 1440, before hydration (HTML and the pre-paint script only)') + bar(CURRENT, 'pre') + note +
               lab('Desktop 1440, after hydration') + bar(CURRENT, '') + '<div style="height:28px"></div>', 1440)


def prehydration_phone_page():
    """A real 390 capture: the viewport is 390, so the phone bar's own media query applies."""
    lab = lambda t: f'<div class="label" style="padding:18px 18px 8px;color:var(--gold)">{t}</div>'
    note = ('<div style="padding:12px 18px 24px;font-size:12px;color:var(--muted)">Before hydration the crest is painted from the pointer and the slot is a link; '
            'hydration adds the chevron badge. The bar is 56 px in both.</div>')
    return doc(lab('Phone, before hydration') + bar(CURRENT, 'pre') + lab('Phone, after hydration') + bar(CURRENT, '') + note, 390, 420)


# ---------------------------------------------------------------------------- render
PASTED_CUR = {**PASTED, 'current': True}
SO_PASTE = desktop_page(1440, ch=PASTED, state='open', popover=pop_signed_out_pasted())
BOARDS = {   # name: (html, viewport width, viewport height)
    'nav': (desktop_page(1440), 1440, 900),
    'nav-open': (desktop_page(1440, state='open', popover=pop_many(), height=620), 1440, 900),
    'nav-open-many': (desktop_page(1440, state='open', popover=pop_many(MANY9), height=700), 1440, 900),
    'nav-2000': (desktop_page(2000), 2000, 900),
    'nav-1439': (desktop_page(1439), 1439, 900),
    'nav-1280': (desktop_page(1280), 1280, 900),
    'nav-1024': (desktop_page(1024), 1024, 900),
    'nav-1023': (desktop_page(1023), 1023, 900),
    'nav-768': (desktop_page(768), 768, 900),
    'nav-phone': (phone_page(), 390, 844),
    'nav-phone-open': (phone_page(sheet=True), 390, 844),
    'nav-phone-menu': (phone_page(menu=True), 390, 844),
    'nav-phone-360': (doc(bar(CURRENT) + body_stub('24px 18px'), 360, 300), 360, 300),
    'nav-signed-out': (desktop_page(1440, ch=None, state='open', popover=pop_signed_out()), 1440, 900),
    'nav-signed-out-1024': (desktop_page(1024, ch=None, state='open', popover=pop_signed_out()), 1024, 900),
    'nav-signed-out-phone': (phone_page(ch=None, menu=True), 390, 844),
    'nav-signed-out-phone-open': (phone_page(ch=None, sheet=True), 390, 844),
    'nav-signed-out-pasted': (SO_PASTE, 1440, 900),
    'nav-expired': (desktop_page(1440, ch={**CURRENT, 'mark': 'session'}, state='open', popover=pop_expired(), tip=False), 1440, 900),
    'nav-prehydration': (prehydration_page(), 1440, 900),
    'nav-prehydration-phone': (prehydration_phone_page(), 390, 420),
    'nav-states': (states_page(), 1440, 900),
}


def main(names):
    scratch = Path(tempfile.mkdtemp(prefix='gen_nav_'))
    RENDERS.mkdir(exist_ok=True)
    for name in names:
        html, width, height = BOARDS[name]
        src = scratch / f'{name}.html'
        src.write_text(html)
        out = RENDERS / f'{name}.png'
        subprocess.run(['npx', 'playwright', 'screenshot', '--browser=chromium', f'--viewport-size={width},{height}', '--full-page',
                        '--wait-for-timeout=3000', src.as_uri(), str(out)], check=True, capture_output=True, cwd=WEB)
        print(out)


if __name__ == '__main__':
    main(sys.argv[1:] or list(BOARDS))

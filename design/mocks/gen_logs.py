"""Mock boards for the /logs redesign (design/specs/2026-10-11-logs.md, round two; revises the 2026-10-04 board).

  python3 design/mocks/gen_logs.py   -> renders/logs{,-signed-out,-2000,-1024,-phone,-states}.png

The header is the shared component (gen_character_header.py). The sample row's counts (63 fights, 12 kills)
come from design/mocks/data/logs-sample-report.json, the public sample report. The character, its own
reports (zones, boss counts, lengths, dates), the device and the pairing code are EXAMPLES, labelled in
the spec only. HTML goes to a scratch dir; only PNGs land in renders/.
"""
import json
import sys

import gen_character_header as H
from mocklib import DATA, GOLD, MUTED, TEXT, BODY, BORDER, SOFT

R = json.load(open(DATA / 'logs-sample-report.json'))['report']
SAMPLE_DAY = 'Sep 15'     # the sample's created_at 2026-09-15, older than two weeks so it prints as a date

STYLE = f'''<style>
.rr{{display:grid;grid-template-columns:1fr auto;gap:2px 16px;align-items:baseline;min-height:56px;padding:10px 0;border-bottom:1px solid {SOFT}}}
.rr:hover .t{{text-decoration:underline;text-underline-offset:4px}}
.rr .t{{font-weight:600;color:{TEXT}}}
.tab{{display:inline-flex;align-items:center;height:44px;padding:0 14px;border:1px solid {BORDER};border-radius:4px;font-size:12px;font-weight:700;letter-spacing:.06em;text-transform:uppercase;color:{TEXT}}}
.tab.on{{border-color:{GOLD};color:{GOLD};background:color-mix(in srgb,{GOLD} 8%,transparent)}}
.drop{{display:flex;flex-direction:column;align-items:center;justify-content:center;gap:8px;min-height:150px;border:1px dashed #4a4030;border-radius:6px;color:{MUTED};font-size:13px;text-align:center;padding:16px}}
.drop.on{{border-color:{GOLD};background:color-mix(in srgb,{GOLD} 6%,transparent)}}
</style>'''


def mono(s: str, color: str = MUTED, size: int = 13) -> str:
    return f'<span class="mono" style="font-size:{size}px;color:{color}">{s}</span>'


def panel(title: str, body: str, aside: str = '', anchor: str = '') -> str:
    a = f'<span style="font-size:12px;color:{MUTED}">{aside}</span>' if aside else ''
    return (f'<section {anchor} class="panel" style="display:flex;flex-direction:column;gap:12px;padding:16px 20px"><div style="display:flex;justify-content:space-between;align-items:baseline;gap:12px">'
            f'<span class="display" style="font-size:15px;font-weight:700;color:{TEXT};letter-spacing:.02em">{title}</span>{a}</div>{body}</section>')


def report_row(title: str, meta: str, days: int | None, date: str = '') -> str:
    """Title and a relative date on the first line; zone, boss kills and length under the title."""
    when = date or H.rel_date(days or 0)
    return (f'<a class="rr" href="#r"><span class="t">{title}</span>{mono(when)}'
            f'<span style="font-size:12px;color:{MUTED};grid-column:1/-1">{meta}</span></a>')


def your_reports(signed_in: bool, state: str = 'list') -> str:
    if not signed_in:
        return panel('Your reports', f'<span style="font-size:14px;color:{BODY}">Reports you own appear here once you are signed in.</span>')
    tabs = '<div style="display:flex;gap:6px"><span class="tab on">Mine</span><span class="tab">Olympus XXVII</span></div>'
    if state == 'empty':
        return panel('Your reports', f'{tabs}<span style="font-size:14px;color:{BODY}">No reports yet. Upload a log, or log live with the companion.</span>')
    if state == 'loading':
        return panel('Your reports', tabs + ''.join(f'<div class="rr">{H.skel("50%", 14)}{H.skel(120, 13)}</div>' for _ in range(3)))
    rows = (report_row('Onyxia, wipes', 'Onyxia\'s Lair · 0 bosses · 1 h 05 m', 1) + report_row('Onyxia, first kill', 'Onyxia\'s Lair · 1 boss · 48 m', 3)
            + report_row('Sanguine Depths run', 'Sanguine Depths · 4 bosses · 2 h 10 m', 18))
    return panel('Your reports', tabs + f'<div>{rows}</div>')


def recent() -> str:
    sample = report_row('Sanguine Depths, sample log', f'Sanguine Depths · {R["kill_count"]} kills', None, SAMPLE_DAY)
    return panel('Recent public reports', f'<span style="font-size:14px;color:{BODY}">No public reports yet. The first raid logs land in December; dungeon logs are welcome now.</span>'
                 f'<div><span class="label">A sample report</span>{sample}</div>')


def upload(signed_in: bool, chosen: bool = False) -> str:
    intro = (f'<span style="font-size:14px;color:{BODY}">A whole {mono("WoWCombatLog.txt", BODY)}. The first fight is readable within ten seconds of the upload finishing, and the rest appear as they are parsed.</span>')
    folder = (f'<span style="font-size:13px;color:{MUTED}">The game writes it to its {mono("Logs", MUTED, 12)} folder, beside {mono("Interface", MUTED, 12)} and {mono("WTF", MUTED, 12)}. '
              f'Up to 4 GB. Type {mono("/combatlog", MUTED, 12)} again to stop logging first, so the file is whole.</span>')
    if not signed_in:
        return panel('Upload a log', intro + folder + f'<span style="font-size:14px;color:{BODY}">Uploading needs a signed-in account; the report is filed under it.</span>', 'For a night already logged', 'id="upload"')
    pick = f'<span style="color:{GOLD};font-weight:700;letter-spacing:.06em;text-transform:uppercase;font-size:12px'
    if chosen:
        zone = f'<div class="drop on"><span class="mono" style="color:{TEXT}">WoWCombatLog.txt</span><span>412,883,201 bytes</span>{pick}">Choose another file</span></div>'
        vis = ''.join(f'<span class="tab {"on" if v == "Public" else ""}" style="height:40px">{v}</span>' for v in ('Public', 'Unlisted', 'Guild', 'Private'))
        form = (f'<span class="label">Title</span><span style="height:44px;border:1px solid {BORDER};border-radius:4px;padding:0 12px;display:flex;align-items:center;color:{MUTED};font-size:14px">Optional</span>'
                f'<span class="label">Who can open it</span><div style="display:flex;gap:6px;flex-wrap:wrap">{vis}</div><span style="font-size:12px;color:{MUTED}">Listed, ranked, anyone can open it.</span>'
                f'<span><a class="btn btn-gold" href="#go" style="height:44px">Upload</a></span>')
    else:
        zone = f'<div class="drop"><span>Drop {mono("WoWCombatLog.txt", TEXT)} here, or</span>{pick}">choose a file</span></div>'
        form = ''
    return panel('Upload a log', intro + folder + zone + form, 'For a night already logged', 'id="upload"')


def companion(signed_in: bool, paired: bool = True) -> str:
    status = f'<span style="font-size:13px;color:{BODY}">MacBook Pro · macOS · last seen {mono("4 min ago", BODY)}</span>' if (signed_in and paired) else ''
    if not signed_in:
        pair = f'<span style="font-size:14px;color:{BODY}">Pairing needs a signed-in account.</span>'
    else:
        pair = (f'<div style="display:flex;align-items:center;gap:16px;flex-wrap:wrap"><span class="mono" style="font-size:26px;letter-spacing:.18em;color:{TEXT};padding:8px 14px;border:1px dashed {BORDER};border-radius:4px">7K3Q-P2ND</span>'
                f'<span style="font-size:13px;color:{MUTED}">Type this into the companion within {mono("9 minutes", MUTED)}.</span></div>')
    head = f'<span class="label">{"Pair another device" if paired else "Pair it with your account"}</span>' if signed_in else ''
    body = (status + f'<span style="font-size:14px;color:{BODY}">Log live while you raid. Downloads and the in-game /combatlog step are on the setup page. <a href="#setup" style="text-decoration:underline">Get set up</a></span>'
            + f'<div style="display:flex;flex-direction:column;gap:10px;padding-top:12px;border-top:1px solid {SOFT}">{head}{pair}</div>')
    return panel('The companion', body, 'Live, while you raid')


def board(width: int, signed_in: bool, paired: bool = True) -> str:
    phone = width < 600
    sample = ''   # the sample report appears once, as the row in Recent public reports
    hero = H.hero_block(width, 'Combat logs', 'Logs', 'Upload a combat log or log live, then read every pull.', sample)
    if signed_in:
        lead = H.character_header(width, 'logs', 'ready', action='Open your latest report →' if phone else None, page_answers=_answers(phone))
    else:
        lead = H.signed_out_band()
    # The companion sits above Upload when a device is paired (the live logger's route), below when none is.
    # Companion above Upload only when a device is paired; signed out matches signed in without a device (Upload first).
    ways = (companion(signed_in, paired) + upload(signed_in)) if (signed_in and paired) else (upload(signed_in) + companion(signed_in, paired))
    col = lambda inner: f'<div style="display:flex;flex-direction:column;gap:{16 if phone else 24}px;min-width:0">{inner}</div>'
    reading = your_reports(signed_in) + recent()
    if width < 1024:
        grid = col(reading + ways)
    else:
        grid = f'<div style="display:grid;grid-template-columns:7fr 5fr;gap:24px;align-items:start">{col(reading)}{col(ways)}</div>'
    return H.page_doc(width, 'Logs', signed_in, STYLE + hero, lead + grid)


def _answers(phone: bool) -> dict:
    a = {**H.PAGE_ANSWERS['logs']}
    if phone:
        a['href'] = '#report'
    return {'logs': a}


def states() -> str:
    from mocklib import page
    cap = lambda t: f'<span class="label" style="color:{GOLD}">{t}</span>'
    cell = lambda t, h, w=610: f'<div style="display:flex;flex-direction:column;gap:8px;width:{w}px">{cap(t)}{h}</div>'
    wide = 1248
    items = [
        cell('Header, ready: latest kill, the note says kill, and the newer report is named', H.character_header(1440, 'logs', 'ready'), wide),
        cell('Header, no logs yet', H.character_header(1440, 'logs', 'empty'), wide),
        cell('Header, phone: the action is Open your latest report; Upload lives only in the upload panel', f'<div style="width:390px">{H.character_header(390, "logs", "ready", layout="card", action="Open your latest report →", page_answers=_answers(True))}</div>', 390),
        cell('Your reports, list (zone, boss kills, length; relative dates inside two weeks, else a date)', your_reports(True)), cell('Your reports, empty', your_reports(True, 'empty')),
        cell('Your reports, loading', your_reports(True, 'loading')), cell('Your reports, signed out: plain words, no button', your_reports(False)),
        cell('Upload, before a file: the zone only; no Upload button yet', upload(True)),
        cell('Upload, a file chosen: title, who can open it, then the outline Upload button', upload(True, True)),
        cell('Upload, signed out: one sentence, no button', upload(False)),
        cell('Upload, drag over the page: gold edge and 6% tint, nothing moves', f'<div class="drop on"><span>Drop {mono("WoWCombatLog.txt", TEXT)} to upload</span></div>'),
        cell('Upload failed (generic line; the detail stays in the console)', panel('Upload a log', f'<span style="font-size:14px;color:{BODY}">The upload did not finish. Your file is untouched; try again.</span><span><a class="btn btn-gold" href="#x" style="height:44px">Try again</a></span>')),
        cell('Recent public reports (the sample is a different row from your own)', recent()),
        cell('Companion, no device yet: sits below Upload and shows the pairing code', companion(True, False)),
        cell('Companion, a device paired: sits above Upload', companion(True, True)),
    ]
    body = f'<div style="{H.inner(1440, "36px 0")};display:flex;flex-direction:column;gap:20px"><h1 class="display" style="font-size:22px;color:{TEXT}">Logs: states</h1><div style="display:flex;flex-wrap:wrap;gap:18px 24px;align-items:flex-start">{"".join(items)}</div></div>'
    return page(H.CSS + STYLE + body, 1440)


BOARDS = {
    'logs': (lambda: board(1440, True), 1440),
    'logs-signed-out': (lambda: board(1440, False), 1440),
    'logs-2000': (lambda: board(2000, True), 2000),
    'logs-1024': (lambda: board(1024, True), 1024),
    'logs-phone': (lambda: board(390, True), 390),
    'logs-states': (states, 1440),
}

if __name__ == '__main__':
    for name in sys.argv[1:] or list(BOARDS):
        build, width = BOARDS[name]
        print(H.shoot(build(), name, width))

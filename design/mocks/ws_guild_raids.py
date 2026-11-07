"""Guild mocks, Raids tab (list, plan form, log match) and the raid page (planner, assignments) (GUILD-ALL lane, 2026-11-07).
Composed only from ws_prims primitives; the data is guild_raids_fixtures.py's, which guild_fixtures.py writes for the build.
Registered in ws_guild2.all_pages().

- web-guild-raids: the officer's Raids tab. Two open Sections: Planned (its rows are the tab's one Plate, Plan an event at the
  head's right) and Raid nights (open rows, the first night open on its pulls).
- web-guild-raids-plan: the same tab with Plan an event's form open under the Planned head (no frame: open Fields).
- web-guild-assignments: the raid planner with the Off tank slot's picker open under its row.
"""
from datetime import datetime
from zoneinfo import ZoneInfo

import guild_overview_data as O
import guild_raids_fixtures as RF
from guild_data import NOW, ZONE
from ws_guild2 import _act, _view, band_line, hdr, raid_planner
from ws_prims import Button, Plate, Section, Tag, build_page

PRIM = '''<style>
.sr-only{position:absolute;width:1px;height:1px;padding:0;margin:-1px;overflow:hidden;clip:rect(0,0,0,0);white-space:nowrap}
.prow{display:grid;grid-template-columns:56px minmax(0,1.5fr) minmax(136px,.8fr) minmax(max-content,1.2fr) 176px 16px;align-items:center;gap:16px;
min-height:64px;padding:7px 14px;color:inherit;text-decoration:none}
.prow+.prow{border-top:1px solid var(--line-soft)}
.prow .dt{display:flex;flex-direction:column;align-items:center;justify-content:center;width:56px;height:56px}
.prow .dt .l{font:700 11px/1 var(--font-body);letter-spacing:.12em;text-transform:uppercase;color:var(--muted)}
.prow .dt .d{font:700 22px/1.25 var(--font-body);color:var(--strong);font-variant-numeric:tabular-nums}
.prow .tt{display:block;font:600 15px/22px var(--font-body);color:var(--strong)}
.prow .sm{display:block;font-size:12px;line-height:16px;color:var(--muted)}
.prow .tm{display:block;font-size:14px;line-height:20px;color:var(--text)}
.prow .cd{display:block;font-size:12px;line-height:16px}
.prow .ct{display:flex;align-items:center;gap:8px;font-size:15px;color:var(--strong)}
.prow .bar{display:block;width:96px;height:8px;border-radius:9px;overflow:hidden}
.prow .chips{display:flex;justify-content:flex-end;gap:6px;flex-wrap:wrap}
.prow .chev{color:var(--muted);font-size:16px}
.night{display:block}.night+.night{border-top:1px solid var(--line-soft)}
.night>button{all:unset;box-sizing:border-box;display:flex;flex-wrap:wrap;align-items:center;gap:4px 12px;width:100%;min-height:var(--row-h,58px);padding:10px 0;cursor:pointer}
.night .zt{display:block;font:700 15px/20px var(--font-body);color:var(--strong)}
.night .zs{display:block;font-size:13px;line-height:18px;color:var(--muted)}
.night .facts{margin-left:auto;font-size:13px;color:var(--muted)}
.night .open{display:flex;flex-direction:column;gap:12px;padding-bottom:16px}
.night .open>a{display:inline-flex;align-items:center;min-height:44px;align-self:flex-start;font-size:14px}
.fights{list-style:none;margin:0;padding:0;display:flex;flex-direction:column}
.fights li{display:flex;align-items:center;gap:12px;min-height:44px;padding:6px 0}.fights li+li{border-top:1px solid var(--line-soft)}
.fights .fl{min-width:160px;font-size:14px}.fights .fl.b{font-weight:600;color:var(--strong)}.fights .fl.q{font-size:13px;color:var(--muted)}
.fights .fd{font-size:13px;color:var(--muted);font-variant-numeric:tabular-nums}.fights .fe{margin-left:auto;font-size:13px;color:var(--muted)}
.rform{display:flex;flex-direction:column;gap:16px}
.rform .g2{display:grid;grid-template-columns:1fr 1fr;gap:16px}.rform .g3{display:grid;grid-template-columns:1fr 1fr 1fr;gap:16px}
.rform .f{display:flex;flex-direction:column;gap:8px}.rform .f>label{font:700 12px/16px var(--font-body);letter-spacing:.08em;text-transform:uppercase;color:var(--muted)}
.rform input,.rform textarea,.rform select{height:44px;width:100%;box-sizing:border-box;padding:0 12px;font:13px var(--font-body);color:var(--text)}
.rform textarea{height:auto;padding:10px 12px;line-height:20px}
.rform .help{font-size:13px;color:var(--muted)}.rform .help.s{font-size:12px}
.rform .acts{display:flex;align-items:center;gap:16px;flex-wrap:wrap}.rform .txt{all:unset;font-size:13px;text-decoration:underline;min-height:44px;cursor:pointer}
.rform .chk{display:flex;align-items:center;gap:12px;min-height:44px;font-size:14px}
</style>'''


def _local(iso: str) -> datetime:
    return datetime.fromisoformat(iso.replace('Z', '+00:00')).astimezone(ZoneInfo(ZONE))


def _countdown(at: datetime) -> str:
    left = at - _local(NOW)
    return f'in {left.days} days {left.seconds // 3600} h' if left.days else f'in {left.seconds // 3600} h'


def _clock(at: datetime) -> str:
    return f"{at.hour % 12 or 12}:{at.minute:02d} {'am' if at.hour < 12 else 'pm'}"


def planned_row(raid: dict) -> str:
    """One planned raid as the built PlannedRaidRow draws it from 768: one link, the date in words for a screen reader."""
    at = _local(raid['starts_at'])
    c = raid['counts']
    signed = c['in'] + c['late']
    you = raid['viewer']['characters'][0]
    chip = (f'<span class="pill">In &middot; DPS</span>' if you['state'] in ('in', 'late') else '<span class="pill">Not signed up</span>')
    tone = 'gold' if you['state'] is not None else 'muted'
    word = lambda n, one, many: f'{n} {one if n == 1 else many}'  # noqa: E731
    parts = [word(c['tanks'], 'tank', 'tanks'), word(c['healers'], 'healer', 'healers'), f"{c['melee']} melee", f"{c['ranged']} ranged"]
    parts += [f"{c['late']} late"] if c['late'] else []
    return (f'<a class="prow" href="#"><span class="dt" aria-hidden="true"><span class="l">{at:%a}</span><span class="d">{at.day}</span>'
            f'<span class="l">{at:%b}</span></span>'
            f'<span><span class="tt">{raid["title"]}</span></span>'
            f'<span><span class="sr-only">{at:%A} {at.day} {at:%B}, </span><span class="tm">{_clock(at)} CST</span>'
            f'<span class="cd {tone}">{_countdown(at)}</span></span>'
            f'<span style="display:flex;flex-direction:column;gap:4px"><span class="ct">{signed}/{raid["size"]}<span class="bar meter"><i style="width:{round(100 * signed / raid["size"])}%"></i></span></span>'
            f'<span class="sm">{" &middot; ".join(parts)}</span></span>'
            f'<span class="chips">{chip}</span><span class="chev" aria-hidden="true">&rsaquo;</span></a>')


def night_rows() -> str:
    page = RF.raids_page()
    out = []
    for i, night in enumerate(page['rows']):
        at = _local(night['created_at'])
        facts = f"1 boss (Onyxia) &middot; 1 kill &middot; 1 wipe"
        head = (f'<button type="button" aria-expanded="{"true" if i == 0 else "false"}"><span><span class="zt">{night["zone"]}</span>'
                f'<span class="zs">{at:%a} {at.day} {at:%b} &middot; {round(night["duration_ms"] / 60000)} min</span></span>'
                f'<span class="facts">{facts}</span></button>')
        body = ''
        if i == 0:
            fights, pulls = [], 0
            for n, f in enumerate(night['fights'], 1):
                secs = round(f['duration_ms'] / 1000)
                dur = f'{secs // 60}:{secs % 60:02d}'
                if f['name'] is None:
                    # No boss id for the pull: the build says so rather than inventing a name (guild spec 12.1).
                    fights.append(f'<li><span class="fl">Pull {n}</span><span class="fd">no named encounter published</span></li>')
                else:
                    pulls += 1
                    tag = Tag('Kill', 'kill') if f['kill'] else Tag('Wipe', 'wipe')
                    fights.append(f'<li><span class="fl b">{f["name"]} &middot; pull {pulls}</span><span class="fd">{dur}</span><span style="margin-left:auto">{tag}</span></li>')
            body = f'<div class="open"><a class="link" href="#">Open the raid</a><ul class="fights">{"".join(fights)}</ul></div>'
        out.append(f'<div class="night">{head}{body}</div>')
    return '<div>' + ''.join(out) + '</div>'


def raid_form() -> str:
    """Plan an event's form, open (RaidForm create, as the built form reads with the guild's Discord channel connected): open
    labelled fields, no frame; the defaults the build fills (the Tuesday after the last planned raid, 8:00 pm, 3 h, the guild's zone)."""
    def f(label, control, extra=''):
        return f'<div class="f"><label>{label}</label>{control}{extra}</div>'
    combo = lambda v: f'<button type="button" class="sel" role="combobox" aria-haspopup="listbox" aria-expanded="false">{v}</button>'  # noqa: E731
    switch = lambda label, on: (f'<label class="chk"><input type="checkbox" role="switch"{" checked" if on else ""} '  # noqa: E731
                                f'style="width:36px;height:20px">{label}</label>')
    return ('<form class="rform">'
            f'<div class="g2">{f("What", combo(""))}{f("Title", "<input type=text>")}</div>'
            f'<div class="g3">{f("Date", "<input type=text value=12/22/2026>")}{f("Time", "<input type=text value=08:00&nbsp;PM>")}{f("Duration", combo("3 h"))}</div>'
            + f("Guild time zone", '<button type="button" class="sel" aria-haspopup="listbox" aria-expanded="false">Central (Chicago) &middot; CST</button>',
                '<p class="help">For you: same time.</p>')
            + f("Notes", "<textarea rows=6></textarea>", "<p class='help s'>Markdown: bold, lists and links. Members see this on the raid page.</p>")
            + '<div class="f">' + switch('Also plan the next 3 weeks at this time', False) + '</div>'
            + '<div class="f">' + switch('Post to Discord', True) + '<p class="help">Posts to #raids on Tuesday 15 December, a week before the raid.</p>'
            + '<div style="padding-left:16px">' + switch('Remind the day before', True) + '<p class="help">Posts a reminder 24 hours before the start.</p></div></div>'
            + '<div style="display:flex;flex-direction:column;gap:4px"><div class="acts">' + _act('Post the event', primary=True, small=False, submit=True)
            + '<button type="button" class="txt">Cancel</button></div><div style="min-height:20px"></div></div>'
            + '</form>')


def raids_tab(form: bool = False) -> str:
    page = RF.raids_page()
    planned = ''.join(planned_row(r) for r in page['planned'])
    body = (raid_form() + '<div style="height:16px"></div>' if form else '') + Plate(planned)
    head = _act('Plan an event', primary=True) if not form else ''
    main = (Section('Planned', f'<div style="padding-top:12px">{body}</div>', head, meta=str(len(page['planned'])))
            + Section('Raid nights', f'<div style="padding-top:12px">{night_rows()}</div>'))
    ledger = f'<div class="pc" style="padding-top:20px"><div style="display:flex;flex-direction:column;gap:var(--section-gap)">{main}</div></div>'
    title = 'Raids, plan form' if form else 'Raids'
    return build_page(f'Guild, {title}', _view(hdr(tabs_active='Raids', actions=Button('Invite link', small=True)) + band_line(O.OFFICER) + ledger) + PRIM,
                      'Guild', ('Dunmaro', 'hunter', 'Beast Mastery Hunter &middot; 60'))


def assignments() -> str:
    """The raid planner with the Off tank slot's picker open (ws_guild2.raid_planner, edited=off_tank)."""
    return raid_planner(edited='off_tank')


PAGES: dict = {'web-guild-raids': raids_tab, 'web-guild-raids-plan': lambda: raids_tab(True), 'web-guild-assignments': assignments}

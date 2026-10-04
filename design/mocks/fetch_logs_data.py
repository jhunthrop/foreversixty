"""Pull the public sample report's real numbers for the logs landing mock.

Writes design/mocks/data/reports-recent.json (the live recent list) and
design/mocks/data/logs-sample-report.json (one boss kill's per-player DPS, healing,
deaths and mechanics). Re-run whenever the sample report changes; never hand-edit.
"""
import json, sys, urllib.request
from pathlib import Path

API = 'https://api.foreversixty.gg'
OUT = Path(__file__).resolve().parent / 'data'


def get(url: str):
    # The edge refuses requests with no User-Agent; name ourselves honestly.
    req = urllib.request.Request(url, headers={'User-Agent': 'foreversixty-design-mocks/1.0'})
    with urllib.request.urlopen(req) as r:
        return json.load(r)


def main() -> int:
    recent = get(f'{API}/v1/reports/recent')
    (OUT / 'reports-recent.json').write_text(json.dumps(recent, indent=1))
    rows = recent['data']['rows']
    if not rows:
        print('no public reports', file=sys.stderr)
        return 1
    rid = rows[0]['id']
    meta = get(f'{API}/v1/reports/{rid}')['data']
    kills = [f for f in meta['fights'] if f['kind'] != 'trash' and f['kill']]
    fight = kills[-1] if kills else meta['fights'][-1]
    s = get(f'{API}/v1/reports/{rid}/files/fights/{fight["index"]}/summary.json')
    dur = s['duration_ms'] / 1000
    roster = {x['guid']: x for x in s['roster']}
    dps = sorted(({'name': d['name'].split('-')[0], 'class': roster[d['guid']].get('class'), 'spec': roster[d['guid']].get('spec'),
                   'role': roster[d['guid']].get('role'), 'dps': round(d['effective'] / dur), 'deaths': roster[d['guid']].get('deaths', 0)}
                  for d in s['damage_done'] if d['guid'] in roster), key=lambda x: -x['dps'])
    hps = sorted(({'name': h['name'].split('-')[0], 'class': h.get('class'), 'hps': round(h['effective'] / dur)}
                  for h in s['healing'] if h['guid'] in roster), key=lambda x: -x['hps'])
    deaths = [{'name': d['name'].split('-')[0], 'at': round(d['at_ms'] / 1000), 'by': (d.get('killing_blow') or {}).get('spell_name')} for d in s['deaths']]
    mech = [{'name': m['name'], 'kind': m['kind'], 'note': m.get('note'), 'hits': sum(p.get('hits', 0) for p in m.get('players', []))} for m in s['mechanics'].get('rows', [])][:5]
    out = {'report': {'id': rid, 'title': meta['title'], 'created_at': meta['created_at'], 'fight_count': rows[0]['fight_count'], 'kill_count': rows[0]['kill_count']},
           'fight': {'index': fight['index'], 'name': fight['name'], 'duration_s': round(dur), 'kill': fight['kill']},
           'dps': dps, 'hps': hps, 'deaths': deaths, 'mechanics': mech}
    (OUT / 'logs-sample-report.json').write_text(json.dumps(out, indent=1))
    print(rid, fight['name'], dps[:2])
    return 0


if __name__ == '__main__':
    raise SystemExit(main())

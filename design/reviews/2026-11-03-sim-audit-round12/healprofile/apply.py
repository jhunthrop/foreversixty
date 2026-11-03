"""Write each band's logged tank share into data/curated/heal-profile.json, then its figures.

apply.py <site worktree> <go.work>. The share is the run tank's (split.run_tanks) over logged
trash fights only, median per fight, over the band's content dungeons, or over every logged
beta dungeon when the content has fewer than MIN_FIGHTS fights."""
import json, math, os, statistics, subprocess, sys
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
import split as S

MIN_FIGHTS = 30
CONTENT = {  # heal-profile content names -> the split tool's dungeon names
    'The Deadmines': 'Deadmines', 'Wailing Caverns': 'Wailing Caverns', 'Razorfen Kraul': 'Razorfen Kraul',
    'Gnomeregan': 'Gnomeregan', 'Scarlet Monastery Graveyard': 'SM Graveyard', 'Scarlet Monastery Armory': 'SM Armory',
}
SOURCE = ('foreverlogs.gg beta reports (the 119 public reports cached 2026-10-31; damage timelines fetched '
          '2026-11-03 with reports/{id}/encounters/{id}/timeline?event_types=damage): a seeded sample of up to 12 '
          'trash fights of 10 s or more a report; tools in design/reviews/2026-11-03-sim-audit-round12/healprofile/')

def goround(x, places):
    f = 10 ** places
    return math.floor(x * f + 0.5) / f

def evidence(rows, names):
    logged = [r for r in rows if r.get('run_share') is not None and r['kind'] == 'trash']
    content = [r for r in logged if r['dungeon'] in names]
    pooled = len(content) < MIN_FIGHTS
    used = logged if pooled else content
    sh = [r['run_share'] for r in used]
    q = statistics.quantiles(sh, n=4)
    return {
        'content_fights': len(content), 'pooled': pooled, 'fights': len(used),
        'dungeons': sorted({r['dungeon'] for r in used}),
        'median_tank_share': goround(statistics.median(sh), 4), 'q25': goround(q[0], 4), 'q75': goround(q[2], 4),
        'source': SOURCE,
    }

def main():
    site, gowork = sys.argv[1], sys.argv[2]
    rows = json.load(open(S.SP + '/r12/healprofile/fl/split.json'))
    S.run_tanks(rows)
    path = site + '/data/curated/heal-profile.json'
    hp = json.load(open(path))
    bands = hp['bands']
    bands['split_rule'] = {'min_fights': MIN_FIGHTS}
    tank = hp['tuning']['base_tank_hit_damage'] / hp['tank']['swing_seconds']
    group = hp['tuning']['base_pulse_damage'] * hp['pulse']['members'] / hp['pulse']['interval_seconds']
    new_rows = []
    for row in bands['rows']:
        out = {k: v for k, v in row.items() if k not in ('tank_share', 'split', 'tank_hit_damage', 'pulse_damage')}
        if row['level'] == bands['reference_level']:
            out['tank_share'] = goround(tank / (tank + group), 3)
        else:
            names = {CONTENT.get(c['name'], c['name']) for c in row['content']}
            ev = evidence(rows, names)
            out['tank_share'] = goround(ev['median_tank_share'], 3)
            out['split'] = ev
        out['tank_hit_damage'] = row['tank_hit_damage']; out['pulse_damage'] = row['pulse_damage']
        new_rows.append(out)
    bands['rows'] = new_rows
    tmp = S.SP + '/r12/healprofile/hp-tmp.json'
    json.dump(hp, open(tmp, 'w'), indent=2, ensure_ascii=False)
    res = subprocess.run(['go', 'test', './request', '-run', 'TestZZPrintBandFigures', '-count=1', '-v'], cwd=site + '/sim',
                         capture_output=True, text=True, env={**os.environ, 'GOWORK': gowork, 'HP': tmp})
    line = next(l for l in res.stdout.splitlines() if l.startswith('FIGURES '))
    figs = json.loads(line[len('FIGURES '):])
    for row in new_rows:
        f = figs[str(row['level'])]
        row['tank_hit_damage'] = f['TankHitDamage']; row['pulse_damage'] = f['PulseDamage']
        print(row['level'], row['tank_share'], row.get('split', {}).get('fights'), row.get('split', {}).get('pooled'),
              'hit', f['TankHitDamage'], 'pulse', f['PulseDamage'], 'tank health', f['TankHealth'])
    text = json.dumps(hp, indent=2, ensure_ascii=False) + '\n'
    open(path, 'w').write(text)

if __name__ == '__main__':
    main()

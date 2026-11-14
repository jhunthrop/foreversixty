"""The items whose ItemSparse stat list the hotfixes restate, with what the Wowhead fill gives them."""
import csv, json, sys
from pathlib import Path
# Usage: data/.venv/bin/python restated_fills.py <repo root> <dir with wago ItemSparse.csv (hotfixed) and ItemSparse_shipped.csv>
sys.path.insert(0, sys.argv[1] + '/data')
from pipeline.wowhead_items import load_items, planner_stats
W=sys.argv[2].rstrip('/') + '/'
B=Path(sys.argv[1]) / 'data/builds/1.60.1.70938'
def stats(row):
    return tuple((row[f'StatModifier_bonusStat_{i}'],row[f'StatPercentEditor_{i}']) for i in range(10) if row[f'StatModifier_bonusStat_{i}'] not in ('-1','0'))
hot={r['ID']:r for r in csv.DictReader(open(W+'ItemSparse.csv'))}
shp={r['ID']:r for r in csv.DictReader(open(W+'ItemSparse_shipped.csv'))}
restated={i for i in hot if i in shp and stats(hot[i])!=stats(shp[i])}
print('restated', len(restated))
wh={str(i.id):i for i in load_items(B/'raw/wowhead/items.json')}
items={}
for f in (B/'items').glob('*.json'):
    for it in json.load(open(f))['items']: items[str(it['id'])]=it
n=0
for i,it in items.items():
    if it.get('stats_source') is not None or i not in wh: continue
    miss={k:v for k,v in planner_stats(wh[i]).items() if v and not it['stats'].get(k)}
    # items json already filled; recompute what client states: unknown; approximate by restated set
    if i in restated:
        n+=1; print(i,it['name'],it['stats'],'wh',dict(wh[i].stats))
print(n)

"""scan_own.py <data/own>: every hunter (by Auto Shot / hunter spells) in every fight: adv_level range, top ranks, fight kind."""
import sys, os, json, glob, collections
import pyarrow.parquet as pq
HUNTER_SPELLS = {'Auto Shot', 'Aimed Shot', 'Multi-Shot', 'Arcane Shot', 'Serpent Sting', 'Raptor Strike', "Hunter's Mark"}
COLS = ['event', 'source_guid', 'source_name', 'spell_id', 'spell_name', 'adv_present', 'adv_info_guid', 'adv_level']
root = sys.argv[1]
out = collections.defaultdict(lambda: {'name': '', 'levels': collections.Counter(), 'fights': set(), 'spells': collections.Counter()})
for rep in sorted(d for d in os.listdir(root) if os.path.exists(f'{root}/{d}/report.json')):
    rj = json.load(open(f'{root}/{rep}/report.json'))
    hunters = {u['guid'] for u in rj['units'] if u.get('class') == 'Hunter'}
    for fp in sorted(glob.glob(f'{root}/{rep}/fights/*/events.parquet')):
        n = fp.split('/')[-2]
        t = pq.read_table(fp, columns=COLS).to_pylist()
        for e in t:
            g = e['source_guid']
            if g not in hunters: continue
            o = out[(rep, g)]; o['name'] = e['source_name']; o['fights'].add(int(n))
            if e['adv_present'] and e['adv_info_guid'] == g and e['adv_level']: o['levels'][e['adv_level']] += 1
            if e['event'] == 'SPELL_CAST_SUCCESS': o['spells'][(e['spell_name'], e['spell_id'])] += 1
for (rep, g), o in sorted(out.items(), key=lambda kv: -max(kv[1]['levels'] or {0: 0})):
    lv = sorted(o['levels']); 
    print(rep, o['name'], 'lvl', (lv[0], lv[-1]) if lv else None, 'fights', len(o['fights']), 'casts', sum(o['spells'].values()),
          [f'{k[0]}:{k[1]}x{v}' for k, v in o['spells'].most_common(8)])

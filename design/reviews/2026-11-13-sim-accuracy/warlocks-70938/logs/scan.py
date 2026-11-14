"""scan.py <dl dir> <spellconst warlock.json>: every unit casting a warlock spell in the downloaded fights, with its
highest advanced-log level, the highest spell level it casts (the real level floor) and its cast count."""
import glob, json, sys, collections
import pyarrow.parquet as pq
dl, sc = sys.argv[1], sys.argv[2]
spells = json.load(open(sc))['spells']
lvl = {int(k): v['spell_level'] for k, v in spells.items() if v['rank'] > 0}
name = {int(k): v['name'] for k, v in spells.items()}
cols = ['event', 'source_guid', 'source_name', 'spell_id', 'adv_present', 'adv_info_guid', 'adv_level']
stat = collections.defaultdict(lambda: {'casts': 0, 'adv': 0, 'rank': 0, 'top': '', 'reps': set(), 'name': ''})
for f in glob.glob(f'{dl}/*/f*.parquet'):
    rep = f.split('/')[-2]
    t = pq.read_table(f, columns=[c for c in cols if c in pq.read_schema(f).names]).to_pandas()
    w = t[t.spell_id.isin(lvl.keys())]
    for g, grp in w.groupby('source_guid'):
        if not str(g).startswith('Player'):
            continue
        s = stat[g]; s['reps'].add(rep); s['name'] = grp.source_name.iloc[0] if 'source_name' in grp else ''
        s['casts'] += int((grp.event == 'SPELL_CAST_SUCCESS').sum())
        sl = grp.spell_id.map(lvl).max()
        if sl > s['rank']:
            s['rank'] = int(sl); s['top'] = name[int(grp.spell_id[grp.spell_id.map(lvl) == sl].iloc[0])]
    a = t[t.adv_present & (t.adv_info_guid == t.source_guid)] if 'adv_present' in t else t.iloc[0:0]
    for g, v in a.groupby('source_guid').adv_level.max().items():
        if g in stat:
            stat[g]['adv'] = max(stat[g]['adv'], int(v))
for g, s in sorted(stat.items(), key=lambda kv: -kv[1]['rank']):
    if s['casts'] >= 5:
        print(f"{g[-8:]} rank-floor {s['rank']:2d} ({s['top']}) adv {s['adv']:2d} casts {s['casts']:5d} {','.join(sorted(s['reps']))}")

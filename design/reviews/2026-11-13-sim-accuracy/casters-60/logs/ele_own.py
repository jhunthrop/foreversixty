"""ele_own.py <dl root> <report> <min level>: the Lightning Bolt users at or above min level in one of our
reports: casts, landed bolts, crits, overload hits (the 4084xx rows), mana on their own cast events.
Players are numbered, never named."""
import sys, json, glob, collections
import pyarrow.parquet as pq
root, rep, minlv = sys.argv[1], sys.argv[2], int(sys.argv[3])
SC = json.load(open(__file__.rsplit('/', 1)[0] + '/../../../../../data/builds/1.60.1.70938/spellconst/shaman.json'))['spells']
LB = {int(k) for k, v in SC.items() if v['name'] == 'Lightning Bolt'}
OVER = {i for i in LB if i >= 400000}
rj = json.load(open(f'{root}/{rep}/report.json'))
sh = {u['guid'] for u in rj['units'] if u.get('class') == 'Shaman'}
st = collections.defaultdict(lambda: collections.Counter())
lv = collections.defaultdict(set); mana = collections.defaultdict(list)
for fp in glob.glob(f'{root}/{rep}/f*.parquet'):
    for e in pq.read_table(fp).to_pylist():
        g = e['source_guid']
        if g not in sh: continue
        if e['adv_present'] and e['adv_info_guid'] == g and e['adv_level']: lv[g].add(e['adv_level'])
        sid = e['spell_id']
        if sid not in LB: continue
        c = st[g]; kind = 'over' if sid in OVER else 'bolt'
        if e['event'] == 'SPELL_CAST_SUCCESS':
            c[kind + '_cast'] += 1
            if e['adv_info_guid'] == g and e['adv_max_power']: mana[g].append(e['adv_current_power'] / e['adv_max_power'])
        elif e['event'] == 'SPELL_DAMAGE':
            c[kind + '_hit'] += 1; c[kind + '_crit'] += bool(e['critical']); c[kind + '_dmg'] += e['amount']
            c[kind + '_id_%d' % sid] += 1
        elif e['event'] == 'SPELL_MISSED': c[kind + '_miss'] += 1
n = 0
for g, c in st.items():
    if not lv[g] or max(lv[g]) < minlv: continue
    n += 1; m = mana[g]
    print('shaman', n, 'levels', sorted(lv[g]), dict(c), 'mana frac at cast: min %.2f median %.2f' % (min(m), sorted(m)[len(m)//2]) if m else '')

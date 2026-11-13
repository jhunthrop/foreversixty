"""var.py <dl root> <spell names comma>: per player, per spell id: non-crit base_amount min/mean/max/n, plus white swing."""
import sys, glob, os, collections, statistics as st
import pyarrow.parquet as pq
root, names = sys.argv[1], set(sys.argv[2].split(','))
C=['event','source_guid','source_name','spell_id','spell_name','base_amount','critical','glancing','adv_present','adv_info_guid','adv_level']
v=collections.defaultdict(list); lv={}
for rep in os.listdir(root):
    for fp in glob.glob(f'{root}/{rep}/*.parquet'):
        for e in pq.read_table(fp,columns=C).to_pylist():
            g=e['source_guid']
            if not g.startswith('Player'): continue
            if e['adv_present'] and e['adv_info_guid']==g and e['adv_level']: lv[g]=e['adv_level']
            if e['base_amount'] is None or e['critical'] or e['glancing']: continue
            if e['spell_name'] in names: v[(g,e['source_name'][:16],e['event'][:10],e['spell_name'],e['spell_id'])].append(e['base_amount'])
cats={k[0] for k in v}
for g in cats:
    pass
for k,x in sorted(v.items(), key=lambda kv:(kv[0][1],kv[0][3])):
    print(k[1], lv.get(k[0]), k[2], k[3], k[4], 'n', len(x), 'min', min(x), 'mean', round(st.mean(x),1), 'max', max(x))

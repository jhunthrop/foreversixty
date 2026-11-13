import pyarrow.parquet as pq, glob, json, collections, sys
rep=sys.argv[1]
units={u['guid']:u for u in json.load(open(f'dl/{rep}/report.json'))['units']}
lv=collections.defaultdict(set); n=collections.Counter()
for f in glob.glob(f'dl/{rep}/f*.parquet'):
    try: t=pq.read_table(f,columns=['source_guid','adv_info_guid','adv_level','event','adv_present']).to_pandas()
    except Exception as e: print('bad',f,e); continue
    t=t[t.adv_present & (t.source_guid==t.adv_info_guid)]
    for g,l in zip(t.source_guid,t.adv_level):
        if units.get(g,{}).get('class')=='Rogue': lv[g].add(int(l)); n[g]+=1
for g in lv: print(g, sorted(lv[g])[:3], max(lv[g]), n[g])

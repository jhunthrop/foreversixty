import pyarrow.parquet as pq, glob, json, pandas as pd
reps=['ctxicu4cuyek','4q5el2xsgoy7','edcgw6rodqs2','eygg26pg4gr2','3sdrcwdgo2yi']
cols=['time_unix_nano','event','source_guid','dest_guid','dest_name','dest_flags','spell_id','spell_name','amount','base_amount','critical','glancing','off_hand','miss_type','power_type','adv_present','adv_info_guid','adv_current_power','adv_max_power','adv_power_cost','adv_attack_power','adv_level','stacks','aura_type','extra_spell_id','school','resisted','absorbed','blocked','overkill']
out=[]
for rep in reps:
    u=json.load(open(f'dl/{rep}/report.json'))['units']
    rog={x['guid'] for x in u if x.get('class')=='Rogue'}
    lvl={x['guid']:x.get('level') for x in u}
    for f in glob.glob(f'dl/{rep}/f*.parquet'):
        try: t=pq.read_table(f).to_pandas()
        except Exception as e: print('bad',f,e); continue
        t=t[t.source_guid.isin(rog)|t.dest_guid.isin(rog)][cols].copy()
        t['rep']=rep; t['fight']=int(f.split('/f')[-1].split('.')[0])
        out.append(t)
df=pd.concat(out); df.to_parquet('rogue_events.parquet')
print(len(df)); print(df.event.value_counts().head(25))

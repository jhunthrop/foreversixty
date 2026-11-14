import pyarrow.parquet as pq,glob,collections,pandas as pd,sys
S='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/lane/dl/'
REND={772,6546,6547,6548,11572,11573,11574}
cols=['time_unix_nano','event','source_guid','source_name','dest_guid','spell_id','spell_name','amount','critical','adv_info_guid','adv_attack_power','adv_level','resisted','absorbed','blocked']
rows=[]
for f in sorted(glob.glob(S+'*/f*.parquet')):
    t=pq.read_table(f,columns=cols).to_pandas()
    rend=t[(t.event=='SPELL_PERIODIC_DAMAGE')&t.spell_id.isin(REND)]
    if rend.empty: continue
    for g in rend.source_guid.unique():
        mine=t[(t.adv_info_guid==g)&(t.source_guid==g)&t.adv_attack_power.notna()]
        ap=mine.adv_attack_power.median() if len(mine) else None
        lv=mine.adv_level.max() if len(mine) else None
        r=rend[rend.source_guid==g]
        rows.append(dict(rep=f.split('/')[-2],fight=f.split('/')[-1],guid=g[-6:],lvl=lv,ap=ap,rank=sorted(r.spell_id.unique()),n=len(r),med=r[r.critical!=True].amount.median(),mx=r.amount.max()))
d=pd.DataFrame(rows)
print(d.groupby(['rep','guid']).agg(lvl=('lvl','max'),ap=('ap','median'),rank=('rank','first'),n=('n','sum'),med=('med','median')).to_string())

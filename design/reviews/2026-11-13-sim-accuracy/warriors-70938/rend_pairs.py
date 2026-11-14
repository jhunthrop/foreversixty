"""rend_pairs.py: each Rend application's cast-time attack power (the warrior's own advanced snapshot on
SPELL_CAST_SUCCESS) against the non-crit ticks of that application, every warrior in our reports."""
import pyarrow.parquet as pq,glob,pandas as pd
S='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/lane/dl/'
BASE={772:5,6546:7,6547:9,6548:11,11572:14,11573:18,11574:21}
cols=['time_unix_nano','event','source_guid','dest_guid','spell_id','amount','critical','adv_info_guid','adv_attack_power','adv_level']
out=[]
for f in sorted(glob.glob(S+'*/f*.parquet')):
    t=pq.read_table(f,columns=cols).to_pandas()
    r=t[t.spell_id.isin(list(BASE))]
    if r.empty: continue
    casts=r[(r.event=='SPELL_CAST_SUCCESS')&(r.adv_info_guid==r.source_guid)]
    ticks=r[(r.event=='SPELL_PERIODIC_DAMAGE')&(r.critical!=True)]
    for c in casts.itertuples():
        nxt=casts[(casts.source_guid==c.source_guid)&(casts.dest_guid==c.dest_guid)&(casts.time_unix_nano>c.time_unix_nano)].time_unix_nano.min()
        end=min(nxt if nxt==nxt else 1e30, c.time_unix_nano+22e9)
        tk=ticks[(ticks.source_guid==c.source_guid)&(ticks.dest_guid==c.dest_guid)&(ticks.time_unix_nano>c.time_unix_nano)&(ticks.time_unix_nano<=end)]
        if len(tk)==0: continue
        out.append(dict(rep=f.split('/')[-2],g=c.source_guid[-6:],spell=c.spell_id,base=BASE[c.spell_id],ap=c.adv_attack_power,lvl=c.adv_level,tick=tk.amount.median(),n=len(tk),spread=tk.amount.max()-tk.amount.min()))
d=pd.DataFrame(out); d.to_csv('rend_pairs.csv',index=False); print(len(d))
d['k']=(d.tick-d.base)/d.ap
print(d.groupby(['rep','g','spell']).agg(n=('n','size'),ap_lo=('ap','min'),ap_hi=('ap','max'),t_lo=('tick','min'),t_hi=('tick','max'),k=('k','median'),spread=('spread','max')).to_string())

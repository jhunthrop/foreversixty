"""dump.py <name prefix> [report]: one player's events in time order (casts, auras, damage, energize) to stdout tsv."""
import pyarrow.parquet as pq,glob,sys,pandas as pd
who=sys.argv[1]; rep=sys.argv[2] if len(sys.argv)>2 else '*'
cols=['time_unix_nano','event','source_guid','source_name','dest_name','spell_id','spell_name','amount','critical','stacks','adv_level','adv_info_guid','adv_current_power','adv_max_power','adv_power_cost','adv_spell_power','adv_power_type','aura_type','miss_type','over_energize','power_type','overheal','resisted','absorbed']
out=[]
for f in sorted(glob.glob(f'dl/{rep}/f*.parquet')):
    t=pq.read_table(f,columns=cols).to_pandas()
    m=t.source_name.str.startswith(who)|t.dest_name.str.startswith(who)
    t=t[m].copy(); t['fight']=f.split('/')[-1]; t['rep']=f.split('/')[1]
    out.append(t)
d=pd.concat(out).sort_values(['rep','fight','time_unix_nano'])
d.to_parquet(f'{who}.parquet')
print(len(d))

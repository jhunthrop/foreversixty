"""costdelta.py <who> <spell name>...: for each cast success of the player, mana in the cast event (before the cost) minus the
next snapshot of the player's own power within 0.3 s on any event that carries it (after the cost)."""
import pandas as pd,sys,collections
who=sys.argv[1]; names=set(sys.argv[2:])
d=pd.read_parquet(who+'.parquet')
me=d[d.source_name.str.startswith(who)].adv_info_guid
guid=me[me.str.startswith('Player')].mode()[0]
out=collections.defaultdict(list)
for fight,g in d.groupby(['rep','fight']):
    g=g.reset_index(drop=True)
    snap=g[(g.adv_info_guid==guid)&g.adv_current_power.notna()&(g.adv_max_power>0)]
    casts=snap[(snap.event=='SPELL_CAST_SUCCESS')&snap.spell_name.isin(names)&snap.source_name.str.startswith(who)]
    for i,r in casts.iterrows():
        nxt=snap[(snap.index>i)&(snap.time_unix_nano>r.time_unix_nano)&(snap.time_unix_nano<=r.time_unix_nano+3e8)]
        nxt=nxt[nxt.time_unix_nano>r.time_unix_nano]
        if len(nxt): out[(r.spell_name,int(r.spell_id),r.adv_power_cost)].append(int(r.adv_current_power-nxt.iloc[0].adv_current_power))
for k,v in sorted(out.items()):
    print(k, 'n',len(v), 'median',sorted(v)[len(v)//2], sorted(v)[:12])

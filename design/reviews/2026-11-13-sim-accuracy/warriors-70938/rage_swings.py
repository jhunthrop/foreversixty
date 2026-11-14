"""rage_swings.py: per warrior, rage gained on landed white swings (adv power on the swing minus the
warrior's previous adv power snapshot within 0.6 s with no cast or energize between), by hand and crit."""
import pyarrow.parquet as pq,glob,collections,pandas as pd,numpy as np
S='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/lane/dl/'
cols=['time_unix_nano','event','source_guid','dest_guid','spell_id','spell_name','amount','critical','glancing','off_hand','adv_info_guid','adv_current_power','adv_max_power','adv_power_type','adv_level','adv_attack_power','power_type']
WAR={2457,71,2458}  # stances
res=[]
for f in sorted(glob.glob(S+'*/f*.parquet')):
    t=pq.read_table(f,columns=cols).to_pandas()
    # warriors: rage power type 1 and adv max power 1000 (rage x10) on their own swing snapshots
    sw=t[(t.event=='SWING_DAMAGE')&(t.adv_info_guid==t.source_guid)&(t.adv_power_type==1)]
    for g in sw.source_guid.unique():
        mine=t[((t.adv_info_guid==g))|(t.source_guid==g)].sort_values('time_unix_nano')
        prev=None
        for r in mine.itertuples():
            own=(r.adv_info_guid==g) and not pd.isna(r.adv_current_power)
            if r.source_guid==g and r.event in ('SPELL_CAST_SUCCESS','SPELL_ENERGIZE','SPELL_PERIODIC_ENERGIZE','SPELL_CAST_START'):
                prev=None
            if r.dest_guid==g and r.event in ('SWING_DAMAGE','SPELL_DAMAGE','SWING_DAMAGE_LANDED','SPELL_ENERGIZE','SPELL_PERIODIC_ENERGIZE','SPELL_PERIODIC_DAMAGE') and r.source_guid!=g:
                prev=None  # rage from damage taken
            if r.event=='SWING_DAMAGE' and r.source_guid==g and own:
                if prev is not None and (r.time_unix_nano-prev[0])<0.6e9 and prev[2]==g:
                    pass
                if prev is not None:
                    res.append(dict(rep=f.split('/')[-2],g=g[-6:],t=r.time_unix_nano,dr=(r.adv_current_power-prev[1]),crit=bool(r.critical),oh=bool(r.off_hand),glance=bool(r.glancing),amt=r.amount,lvl=r.adv_level,ap=r.adv_attack_power,gap=(r.time_unix_nano-prev[0])/1e9))
            if own: prev=(r.time_unix_nano,r.adv_current_power,g)
d=pd.DataFrame(res); d.to_parquet('rage_swings.parquet'); print(len(d))

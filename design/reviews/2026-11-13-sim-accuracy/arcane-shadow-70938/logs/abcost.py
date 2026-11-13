import pandas as pd,sys
who=sys.argv[1]
d=pd.read_parquet(who+'.parquet')
s=d[d.source_name.str.startswith(who)|d.dest_name.str.startswith(who)].reset_index(drop=True)
stack=0; cc=False; mb=False; pom=False; ap=False
for _,r in s.iterrows():
    sid=r.spell_id; ev=r.event
    if sid==400573:
        if ev=='SPELL_AURA_APPLIED': stack=1
        elif ev=='SPELL_AURA_APPLIED_DOSE': stack=int(r.stacks)
        elif ev=='SPELL_AURA_REMOVED': stack=0
    if sid==12536: cc = ev in('SPELL_AURA_APPLIED','SPELL_AURA_REFRESH') or (cc and ev!='SPELL_AURA_REMOVED')
    if sid==12042: ap = ev in('SPELL_AURA_APPLIED','SPELL_AURA_REFRESH') or (ap and ev!='SPELL_AURA_REMOVED')
    if ev=='SPELL_CAST_SUCCESS' and r.source_name.startswith(who) and r.spell_name in('Arcane Blast','Arcane Wrath','Arcane Explosion','Fire Blast','Frostbolt'):
        print(r.rep, r.fight, r.spell_name, sid, 'stacks_before_or_after', stack, 'cc', cc, 'cost', r.adv_power_cost, 'mana', r.adv_current_power, '/', r.adv_max_power, 'lvl', r.adv_level, 'sp', r.adv_spell_power)

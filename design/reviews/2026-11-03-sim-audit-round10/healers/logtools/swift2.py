"""Every Swiftmend in the beta timelines: which HoTs were up on the target, which one its amount matches, and whether that HoT kept ticking."""
import collections,glob,json
SP='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
TL=SP+'/r10/healers/fl/tl'
TICKS={'Rejuvenation':4,'Regrowth':7}
stats=collections.Counter()
for f in sorted(glob.glob(TL+'/*.json')):
    rid,enc,cid=map(int,f.split('/')[-1].split('.')[:3])
    try: ev=json.load(open(f)).get('events') or []
    except Exception: continue
    if not any(e['spell_name']=='Swiftmend' for e in ev): continue
    over=collections.defaultdict(float)
    for e in ev:
        if e['event_type']=='overheal': over[(e['timestamp_ms'],e['target_id'],e['spell_id'])]+=float(e['amount'])
    raw=lambda e: float(e['amount'])+over.get((e['timestamp_ms'],e['target_id'],e['spell_id']),0)
    for i,e in enumerate(ev):
        if e['spell_name']!='Swiftmend' or e['event_type']!='heal' or e['source_id']!=cid: continue
        t=e['timestamp_ms']; tg=e['target_id']
        up={}
        for x in ev:
            if x['target_id']!=tg or x['spell_name'] not in TICKS or x['source_id']!=cid: continue
            if x['timestamp_ms']>t: break
            if x['event_type']=='buff_applied': up[x['spell_name']]={'at':x['timestamp_ms'],'tick':None}
            elif x['event_type']=='buff_removed': up.pop(x['spell_name'],None)
            elif x['event_type']=='heal' and x['spell_name'] in up and not x['is_critical'] and x['timestamp_ms']>up[x['spell_name']]['at']:
                up[x['spell_name']]['tick']=raw(x)
        after={n:[x for x in ev if x['target_id']==tg and x['spell_name']==n and x['source_id']==cid and t<x['timestamp_ms']<=t+3500 and x['event_type'] in('heal','buff_removed')] for n in up}
        kept={n:any(x['event_type']=='heal' for x in a) for n,a in after.items()}
        removed_at_sm={n:any(x['event_type']=='buff_removed' and x['timestamp_ms']-t<=60 for x in a) for n,a in after.items()}
        amt=raw(e)/(1.5 if e['is_critical'] else 1)
        match=[n for n,u in up.items() if u['tick'] and abs(amt-TICKS[n]*u['tick'])<=max(6,0.03*amt)]
        left={n:round(( {'Rejuvenation':12000,'Regrowth':21000}[n]-(t-u['at']))/1000,1) for n,u in up.items()}
        key=('both' if len(up)==2 else 'one')
        stats[(key,tuple(match))]+=1
        for n in up: stats[('kept' if kept[n] else 'stopped', n, 'removed@sm' if removed_at_sm[n] else '')]+=1
        print(rid,enc,cid,f"SM {raw(e):.0f}{' crit' if e['is_critical'] else ''}", 'up', {n:(u['tick'],left[n]) for n,u in up.items()}, 'match',match,'kept',kept, 'removed@SM',removed_at_sm)
for k,v in sorted(stats.items()): print(k,v)

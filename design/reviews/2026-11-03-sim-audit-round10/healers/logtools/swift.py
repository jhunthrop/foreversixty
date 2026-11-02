import collections,glob,json
SP='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
TL=SP+'/r10/healers/fl/tl'
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
        if e['spell_name']!='Swiftmend' or e['event_type']!='heal': continue
        t=e['timestamp_ms']; tg=e['target_id']
        prev=[x for x in ev[:i] if x['event_type']=='heal' and x['target_id']==tg and x['spell_name'] in('Rejuvenation','Regrowth') and t-x['timestamp_ms']<25000]
        rem=[x for x in ev if x['event_type']=='buff_removed' and x['target_id']==tg and abs(x['timestamp_ms']-t)<=50 and x['spell_name'] in ('Rejuvenation','Regrowth')]
        last={}
        for x in prev: last[x['spell_name']]=(x['spell_id'],raw(x),x['is_critical'],t-x['timestamp_ms'])
        # applications
        apps=[x for x in ev[:i] if x['event_type'] in ('buff_applied','cast') and x['target_id']==tg and x['spell_name'] in('Rejuvenation','Regrowth')]
        lastapp={x['spell_name']:t-x['timestamp_ms'] for x in apps}
        print(rid,enc,cid,'SM',raw(e),'crit' if e['is_critical'] else '', 'removed',[r['spell_name'] for r in rem],'last ticks',last,'since cast',lastapp)

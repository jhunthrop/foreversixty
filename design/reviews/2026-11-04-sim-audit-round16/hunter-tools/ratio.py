import json,glob,os,collections,statistics
SP='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
lvl={}
for f in glob.glob(SP+'/fl/rep/*.roster.json'):
    rid=os.path.basename(f).split('.')[0]
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            tabs=s['talents']['talent_groups'][str(s['talents']['active_group'])]['tabs']
            lvl[(rid,str(c['character_id']))]=sum(v['points'] for v in tabs.values())+9
rows=collections.defaultdict(lambda: collections.defaultdict(list))
for f in sorted(glob.glob('full/*.json')):
    rid,eid,cid=os.path.basename(f).split('.')[:3]
    try: ev=json.load(open(f))['events']
    except Exception: continue
    L=lvl.get((rid,cid)); ci=int(cid)
    for e in ev:
        if e.get('source_owner_id')==ci and e['spell_id']==-1 and e['event_type']=='damage' and not e['is_critical'] and not e.get('is_glancing'):
            who='hawk' if e['source_name']=='Hawk' else 'pet:'+e['source_name']
            rows[(rid,cid,L)][who].append(int(e['amount']))
        if e['source_id']==ci and e['spell_id']==75 and e['event_type']=='damage' and not e['is_critical']:
            rows[(rid,cid,L)]['autoshot'].append(int(e['amount']))
for k,v in sorted(rows.items(),key=lambda x:x[0][2] or 0):
    print(k, {w:(len(a),round(statistics.mean(a),1)) for w,a in v.items() if len(a)>=5})

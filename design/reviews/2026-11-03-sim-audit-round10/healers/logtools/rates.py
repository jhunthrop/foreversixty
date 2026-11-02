"""Per character: crit rate of each heal spell (direct vs tick unknown in events: use spell ids), with expected spell crit from Int."""
import collections,glob,json,csv,sys
SP='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
TL=SP+'/r10/healers/fl/tl'
want=sys.argv[1].upper() if len(sys.argv)>1 else None
CLS={'WARRIOR':1,'PALADIN':2,'HUNTER':3,'ROGUE':4,'PRIEST':5,'SHAMAN':7,'MAGE':8,'WARLOCK':9,'DRUID':11}
base={2:3.5,5:0.8,7:2.3,11:1.8}
rates={}
for r in csv.DictReader(open('/Users/jh/code/forever/.claude/worktrees/r10-healers/data/builds/1.60.1.70707/gametables/playerexpectedstat.csv')):
    rates[(int(r['ClassID']),int(r['Level']))]=float(r['SpellCritPerIntellect'])*100
snap={}
for f in glob.glob(SP+'/fl/rep/*.roster.json'):
    rid=int(f.split('/')[-1].split('.')[0])
    for c in json.load(open(f))['roster']:
        for s in (c.get('snapshots') or [])[:1]:
            tabs=s['talents']['talent_groups'][str(s['talents']['active_group'])]['tabs']
            pts={v['name']:v['points'] for v in tabs.values()}
            snap[(rid,c['character_id'])]=(s['player']['class'],sum(pts.values())+9,pts,s['stats'].get('intellect',0))
agg=collections.defaultdict(lambda:[0,0]); names={}
meta={}
for f in glob.glob(TL+'/*.json'):
    rid,enc,cid=map(int,f.split('/')[-1].split('.')[:3])
    info=snap.get((rid,cid))
    if not info or (want and info[0]!=want): continue
    try: d=json.load(open(f))
    except Exception: continue
    for e in d.get('events') or []:
        if e['event_type']!='heal' or e['source_id']!=cid: continue
        k=(cid,rid,e['spell_id']); names[e['spell_id']]=e['spell_name']
        agg[k][0]+=1; agg[k][1]+=bool(e['is_critical']); meta[(cid,rid)]=info
for (cid,rid),info in sorted(meta.items(), key=lambda x:(x[1][0],x[1][1])):
    cls,L,pts,intel=info
    exp=base.get(CLS[cls],0)+intel*rates.get((CLS[cls],L),0)
    rows=[(names[s],s,n,c) for (c2,r2,s),(n,c) in agg.items() if c2==cid and r2==rid and n>=20]
    if not rows: continue
    print(f"{cls} {cid} r{rid} L{L} {pts} int {intel} exp spell crit(no buffs/talents) {exp:.1f}")
    for nm,s,n,c in sorted(rows,key=lambda r:-r[2]): print(f"    {nm} {s}: {n} heals, crit {100*c/n:.1f}%")

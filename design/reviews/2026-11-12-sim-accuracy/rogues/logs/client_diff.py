import csv,sys,collections
W='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/wago'
N='wago'
def load(path,key='SpellID'):
    rows=collections.defaultdict(list)
    with open(path) as f:
        for r in csv.DictReader(f): rows[r.get(key)].append(r)
    return rows
ids=set()
for r in csv.DictReader(open(N+'/SpellClassOptions_70888.csv')):
    if r['SpellClassSet']=='8': ids.add(r['SpellID'])
for r in csv.DictReader(open(W+'/SpellClassOptions_70790h.csv')):
    if r['SpellClassSet']=='8': ids.add(r['SpellID'])
for r in csv.DictReader(open(N+'/Talent_70888.csv')):
    if r['ClassID']=='4':
        for i in range(9):
            if r['SpellRank_%d'%i]!='0': ids.add(r['SpellRank_%d'%i])
# triggered spells, two hops
eff=load(N+'/SpellEffect_70888.csv')
for _ in range(2):
    for s in list(ids):
        for e in eff.get(s,[]):
            if e['EffectTriggerSpell'] not in ('0',''): ids.add(e['EffectTriggerSpell'])
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(N+'/SpellName_70888.csv'))}
print(len(ids),'rogue spells',file=sys.stderr)
open('rogue_ids.txt','w').write('\n'.join(sorted(ids,key=int)))
tables=['SpellEffect','SpellMisc','SpellCooldowns','SpellPower','SpellAuraOptions','SpellCategories','SpellClassOptions','SpellLevels','SpellEquippedItems','SpellAuraRestrictions','SpellCastingRequirements']
for t in tables:
    old=load(W+'/%s_70790h.csv'%t) if __import__('os').path.exists(W+'/%s_70790h.csv'%t) else load(W+'/%s_70790_h1.csv'%t)
    new=load(N+'/%s_70888.csv'%t)
    for s in sorted(ids,key=int):
        o=sorted([{k:v for k,v in r.items()} for r in old.get(s,[])],key=lambda r:r['ID'])
        n=sorted([{k:v for k,v in r.items()} for r in new.get(s,[])],key=lambda r:r['ID'])
        if o==n: continue
        def norm(rows): return {r['ID']:r for r in rows}
        O,Nn=norm(o),norm(n)
        for i in sorted(set(O)|set(Nn)):
            a,b=O.get(i),Nn.get(i)
            if a==b: continue
            if a and b:
                ch={k:(a[k],b[k]) for k in a if a[k]!=b.get(k)}
                # ignore float repr noise
                ch2={}
                for k,(x,y) in ch.items():
                    try:
                        if abs(float(x)-float(y))<1e-6: continue
                    except: pass
                    ch2[k]=(x,y)
                if ch2: print(t,s,names.get(s),'row',i,ch2)
            else:
                print(t,s,names.get(s),'row',i,'ADDED' if b else 'REMOVED',(b or a))

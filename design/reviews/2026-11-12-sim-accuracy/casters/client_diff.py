import csv,sys,collections,os
W=sys.argv[1]; classset=sys.argv[2]; talentclass=sys.argv[3]; out=sys.argv[4]
def load(path,key='SpellID'):
    rows=collections.defaultdict(list)
    with open(path) as f:
        for r in csv.DictReader(f): rows[r.get(key)].append(r)
    return rows
ids=set()
for suf in ('70888h','70790h'):
    for r in csv.DictReader(open(W+'/SpellClassOptions_%s.csv'%suf)):
        if r['SpellClassSet']==classset: ids.add(r['SpellID'])
    for r in csv.DictReader(open(W+'/Talent_%s.csv'%suf)):
        if r['ClassID']==talentclass:
            for i in range(9):
                if r['SpellRank_%d'%i]!='0': ids.add(r['SpellRank_%d'%i])
eff=load(W+'/SpellEffect_70888h.csv')
for _ in range(2):
    for s in list(ids):
        for e in eff.get(s,[]):
            if e['EffectTriggerSpell'] not in ('0',''): ids.add(e['EffectTriggerSpell'])
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(W+'/SpellName_70888h.csv'))}
print(len(ids),'spells',file=sys.stderr)
tables=['SpellEffect','SpellMisc','SpellCooldowns','SpellPower','SpellAuraOptions','SpellCategories','SpellClassOptions','SpellLevels','SpellDuration','SpellCastTimes','SpellProcsPerMinute','SpellInterrupts']
o=open(out,'w')
for t in tables:
    old=load(W+'/%s_70790h.csv'%t); new=load(W+'/%s_70888h.csv'%t)
    for s in sorted(ids,key=int):
        O={r['ID']:r for r in old.get(s,[])}; Nn={r['ID']:r for r in new.get(s,[])}
        if O==Nn: continue
        for i in sorted(set(O)|set(Nn)):
            a,b=O.get(i),Nn.get(i)
            if a==b: continue
            if a and b:
                ch={k:(a[k],b[k]) for k in a if a.get(k)!=b.get(k)}
                o.write('%s %s %s row %s changed %s\n'%(t,s,names.get(s,'?'),i,ch))
            elif b: o.write('%s %s %s row %s NEW %s\n'%(t,s,names.get(s,'?'),i,{k:v for k,v in b.items() if v not in ('0','','0.0')}))
            else: o.write('%s %s %s row %s REMOVED\n'%(t,s,names.get(s,'?'),i))
o.close()

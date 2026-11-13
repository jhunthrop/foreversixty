import csv,collections
W='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/wago'
def load(p,key='SpellID'):
    d=collections.defaultdict(list)
    for r in csv.DictReader(open(p)): d[r[key]].append(r)
    return d
cls=set(r['SpellID'] for r in csv.DictReader(open(W+'/SpellClassOptions_70888h.csv')) if r['SpellClassSet']=='8')
eff=load(W+'/SpellEffect_70888h.csv')
ids=set(cls)
for t in csv.DictReader(open(W+'/Talent_70888h.csv')) if __import__('os').path.exists(W+'/Talent_70888h.csv') else []:
    if t['ClassID']=='4':
        for i in range(9):
            v=t.get('SpellRank_%d'%i,'0')
            if v not in ('0',''): ids.add(v)
for _ in range(3):
    for s in list(ids):
        for e in eff.get(s,[]):
            if e['EffectTriggerSpell'] not in ('0',''): ids.add(e['EffectTriggerSpell'])
misc=load(W+'/SpellMisc_70888h.csv')
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(W+'/SpellName_70888h.csv'))}
print(len(ids))
for s in sorted(ids,key=int):
    es=eff.get(s,[])
    per=[e for e in es if e['EffectAura'] in ('3','89','53','64','226') or (e['Effect']=='2')]
    if not per: continue
    m=misc.get(s,[{}])[0]
    a=[int(m.get('Attributes_%d'%i,0) or 0) for i in range(16)]
    cant=bool(a[2]&0x20000000); pcc=bool(a[8]&0x200) if len(a)>8 else False
    desc=';'.join('E%s/A%s bp%s amp%s mech%s' %(e['Effect'],e['EffectAura'],e['EffectBasePointsF'] if 'EffectBasePointsF' in e else e.get('EffectBasePoints'),e.get('EffectAuraPeriod'),e.get('EffectMechanic','')) for e in per)
    print(s,names.get(s,'?'),'school',m.get('SchoolMask'),'CANT_CRIT' if cant else '-', 'PCC' if pcc else '-', 'a2=%#x a3=%#x'%(a[2],a[3]), desc)

import csv,sys,collections
W=sys.argv[1]
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(W+'/SpellName_70938h.csv'))}
for t in ['SpellEffect','SpellMisc','SpellCooldowns','SpellPower','SpellAuraOptions','SpellCategories','SpellClassOptions','SpellLevels','SpellEquippedItems','SpellInterrupts']:
    O={r['ID']:r for r in csv.DictReader(open(W+'/%s_70888h.csv'%t))}
    N={r['ID']:r for r in csv.DictReader(open(W+'/%s_70938h.csv'%t))}
    for i in sorted(set(O)|set(N),key=lambda x:int(x) if x.isdigit() else 0):
        a,b=O.get(i),N.get(i)
        if a==b: continue
        r=b or a; s=r.get('SpellID','?')
        if a and b:
            ch={k:(a[k],b[k]) for k in a if a.get(k)!=b.get(k) and not (k.endswith('Amplitude') or k=='EffectBonusCoefficient' and abs(float(a[k])-float(b[k]))<1e-6)}
            if ch: print(t,s,names.get(s,'?'),'row',i,'changed',ch)
        elif b: print(t,s,names.get(s,'?'),'row',i,'NEW',{k:v for k,v in b.items() if v not in ('0','','0.0')})
        else: print(t,s,names.get(s,'?'),'row',i,'REMOVED')

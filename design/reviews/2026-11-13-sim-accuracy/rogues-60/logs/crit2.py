import csv,collections,sys
W='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/wago'
eff=collections.defaultdict(list)
for r in csv.DictReader(open(W+'/SpellEffect_70888h.csv')): eff[r['SpellID']].append(r)
misc={r['SpellID']:r for r in csv.DictReader(open(W+'/SpellMisc_70888h.csv'))}
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(W+'/SpellName_70888h.csv'))}
for s in sys.argv[1:]:
    todo=[s]; seen=set()
    while todo:
        x=todo.pop()
        if x in seen: continue
        seen.add(x); m=misc.get(x,{})
        a2=int(m.get('Attributes_2',0) or 0); a8=int(m.get('Attributes_8',0) or 0)
        print(x,names.get(x),'CANT' if a2&0x20000000 else '-','PCC' if a8&0x200 else '-',[(e['Effect'],e['EffectAura'],e['EffectBasePointsF'],e['EffectAuraPeriod'],e['EffectTriggerSpell']) for e in eff.get(x,[])])
        for e in eff.get(x,[]):
            if e['EffectTriggerSpell'] not in ('0',''): todo.append(e['EffectTriggerSpell'])

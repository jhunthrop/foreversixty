"""rows.py <class> <talent name>... | ids: dump 70938 hotfixed client rows."""
import json,csv,sys
W='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/wago/'
D='/Users/jh/code/forever-sim-arcshadow/data/builds/1.60.1.70938/'
cls=sys.argv[1]; names=sys.argv[2:]
t=json.load(open(D+f'talents/{cls}.json'))['trees']
ids=[]
for tree in t:
  for ta in tree['talents']:
    if ta['name'] in names:
      print('##',ta['name'], [r.get('spell_id') for r in ta['ranks']]); print('   ', ta['ranks'][-1]['description'])
      ids+=[r['spell_id'] for r in ta['ranks']]
ids+=[int(x) for x in names if x.isdigit()]
def load(t): return [r for r in csv.DictReader(open(W+t+'.csv')) if r.get('SpellID','').isdigit() and int(r['SpellID']) in ids]
for r in sorted(load('SpellEffect'),key=lambda r:(int(r['SpellID']),int(r['EffectIndex']))):
    print(r['SpellID'],'e',r['EffectIndex'],'eff',r['Effect'],'aura',r['EffectAura'],'bp',r['EffectBasePointsF'],'misc',r['EffectMiscValue_0'],'per',r['EffectAuraPeriod'],'coef',r['EffectBonusCoefficient'],'trig',r['EffectTriggerSpell'],'mask',r['EffectSpellClassMask_0'],r['EffectSpellClassMask_1'],r['EffectSpellClassMask_2'],r['EffectSpellClassMask_3'])
for tb in ['SpellAuraOptions','SpellCooldowns','SpellCategories','SpellPower','SpellMisc','SpellClassOptions']:
    for r in load(tb):
        print(tb, {k:v for k,v in r.items() if v not in ('0','','0.0')})

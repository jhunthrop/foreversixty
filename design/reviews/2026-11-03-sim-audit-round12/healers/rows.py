"""Print every client row for spell ids (70753 layered). Usage: rows.py id [id...]"""
import csv, sys, functools
C='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r11/c70753/builds/1.60.1.70753/raw'
T={'SpellName':'ID','Spell':'ID','SpellMisc':'SpellID','SpellEffect':'SpellID','SpellClassOptions':'SpellID','SpellLevels':'SpellID','SpellPower':'SpellID','SpellCooldowns':'SpellID','SpellAuraOptions':'SpellID','SpellCategories':'SpellID'}
@functools.lru_cache(None)
def load(t):
    try: f=open(f'{C}/{t}.csv')
    except FileNotFoundError: return {}
    out={}
    for r in csv.DictReader(f): out.setdefault(r[T[t]],[]).append(r)
    return out
@functools.lru_cache(None)
def lk(t,key='ID'):
    return {r[key]:r for r in csv.DictReader(open(f'{C}/{t}.csv'))}
def nz(r): return {k:v for k,v in r.items() if v not in ('0','','0.0','-1') or k in('EffectIndex',)}
def show(i):
    print(f'===== {i} {load("SpellName").get(i,[{}])[0].get("Name_lang")}')
    for t in T:
        if t=='SpellName': continue
        for r in load(t).get(i,[]):
            r=nz(r)
            if t=='SpellMisc':
                ct=lk('SpellCastTimes').get(r.get('CastingTimeIndex',''),{}).get('Base')
                du=lk('SpellDuration').get(r.get('DurationIndex',''),{}).get('Duration')
                r['_cast']=ct; r['_dur']=du
            if t=='Spell': r={k:v for k,v in r.items() if k in('ID','NameSubtext_lang')}
            print(' ',t,r)
if __name__=='__main__':
    for i in sys.argv[1:]: show(i)

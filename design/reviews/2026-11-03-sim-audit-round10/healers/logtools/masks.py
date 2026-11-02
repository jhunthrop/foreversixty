"""masks.py <classset> <hex m0> <hex m1> <hex m2> <hex m3>: list spells (named, ranked) in the class set matching any bit."""
import csv,sys,collections
RAW="/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad/r9/casters/raw/"
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(RAW+'SpellName.csv'))}
cs=sys.argv[1]; m=[int(x,16)&0xffffffff for x in sys.argv[2:6]]
hit=collections.defaultdict(list)
for r in csv.DictReader(open(RAW+'SpellClassOptions.csv')):
    if r['SpellClassSet']!=cs: continue
    sm=[int(r['SpellClassMask_%d'%i])&0xffffffff for i in range(4)]
    if any(a&b for a,b in zip(sm,m)):
        hit[names.get(r['SpellID'],'?')].append(r['SpellID'])
for n,ids in sorted(hit.items()): print(n, len(ids), ids[:6])

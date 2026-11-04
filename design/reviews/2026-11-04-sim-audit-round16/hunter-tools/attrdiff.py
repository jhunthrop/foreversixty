import csv,sys
def load(p):
    return {r['SpellID']:r for r in csv.DictReader(open(p))}
new=load(sys.argv[1]); old=load(sys.argv[2])
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open(sys.argv[3]))}
ids=sys.argv[4].split(',')
for i in ids:
    n=new.get(i); o=old.get(i)
    if not n: print(i,'missing new'); continue
    diffs=[]
    for k in n:
        if k.startswith('Attributes') or k in('RangeIndex','DurationIndex'):
            if o is None or n[k]!=o.get(k): diffs.append(f"{k}:{o.get(k) if o else None}->{n[k]}")
    print(i,names.get(i),diffs)

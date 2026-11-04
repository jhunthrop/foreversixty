import csv,sys
new={r['SpellID']:r for r in csv.DictReader(open('c755/SpellMisc.csv'))}
old={r['SpellID']:r for r in csv.DictReader(open('era/SpellMisc.csv'))}
names={r['ID']:r['Name_lang'] for r in csv.DictReader(open('c755/SpellName.csv'))}
for pat in sys.argv[1:]:
    for i,n in names.items():
        if n==pat and i in new:
            a=int(new[i]['Attributes_3']); b=int(old[i]['Attributes_3']) if i in old else None
            print(f"{pat:32} {i:>8} attr3 era={b} forever={a} onlyPlayers={'Y' if a&256 else '-'} era={'Y' if b and b&256 else '-'}")

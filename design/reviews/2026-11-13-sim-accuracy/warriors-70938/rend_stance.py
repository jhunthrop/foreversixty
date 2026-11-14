"""rend_stance.py <log> <player guid>: Rend 6548 applications by the stance cast before them, with the
warrior's attack power on the cast (advanced field 16) and the median non-crit tick (field 31; crit 38)."""
import sys,statistics,collections
stance={'2457':('Battle',1.0),'71':('Defensive',0.9),'2458':('Berserker',1.0)}
G=sys.argv[2]; cur=None; apps=[]; last={}
for line in open(sys.argv[1],encoding='utf-8'):
  if G not in line: continue
  f=line.split('  ',1)[1].rstrip().split(',')
  if f[0]=='SPELL_CAST_SUCCESS' and f[1]==G:
    if f[9] in stance: cur=f[9]
    if f[9]=='6548' and f[12]==G:
      apps.append(dict(stance=cur,ap=int(f[16]),ticks=[])); last[f[5]]=apps[-1]
  if f[0]=='SPELL_PERIODIC_DAMAGE' and f[1]==G and f[9]=='6548':
    a=last.get(f[5])
    if a is not None and f[38]!='1': a['ticks'].append(int(f[31]))
rows=collections.defaultdict(list)
for a in apps:
  if a['ticks'] and a['stance']:
    name,m=stance[a['stance']]; t=statistics.median(a['ticks'])
    rows[name].append((a['ap'],t,round((t/m-11)/a['ap'],5),len(a['ticks'])))
for s,v in rows.items():
  print(s,'apps',len(v),'ticks',sum(x[3] for x in v),'k median',statistics.median([x[2] for x in v]))
  print('  ',sorted(v))

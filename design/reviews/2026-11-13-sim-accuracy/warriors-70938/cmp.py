"""cmp.py <before> <after>...: DPS or threat per entry and the change, from bands.sh outputs."""
import re,sys
def load(f):
  d={}
  for l in open(f):
    m=re.match(r'^(\S+ \S+ \d+ \S+) dps ([\d.]+) ([\d.]+)',l)
    if m: d[m.group(1)]=float(m.group(2))
    m=re.match(r'^(\S+ raid 60 \S+) \| threat /s \| ([\d.]+)',l)
    if m: d[m.group(1)]=float(m.group(2))
  return d
b=load(sys.argv[1]); others=[load(f) for f in sys.argv[2:]]
for k in b:
  print(f"{k:42} {b[k]:8.1f} "+" ".join(f"{o.get(k,0):8.1f} {100*(o.get(k,0)/b[k]-1):+6.1f}%" for o in others))

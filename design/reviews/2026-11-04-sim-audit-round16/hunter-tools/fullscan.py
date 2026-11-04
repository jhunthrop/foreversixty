import json, subprocess, sys, time, os, glob, re
SP='/private/tmp/claude-501/-Users-jh-code-forever/54c69daf-a689-41fc-ad89-a21df2102a2f/scratchpad'
FL=SP+'/r3/casters/py/fl.sh'; IN=SP+'/r16/hunter/fl'; OUT=SP+'/r16/hunter/full'
os.makedirs(OUT,exist_ok=True)
per={}
for f in sorted(glob.glob(IN+'/2*.json')):
    if '"spell_id": 1293241' in open(f).read():
        rid,eid,cid=os.path.basename(f).split('.')[:3]
        per.setdefault((rid,cid),[]).append(eid)
todo=[]
for k,v in per.items():
    for eid in v[:int(sys.argv[1])]: todo.append((k[0],eid,k[1]))
print(len(todo),'encounters',file=sys.stderr)
for rid,eid,cid in todo:
    fn=f'{OUT}/{rid}.{eid}.{cid}.json'
    if os.path.exists(fn): continue
    for a in range(6):
        d=subprocess.run(['sh',FL,f'reports/{rid}/encounters/{eid}/timeline?event_types=damage,cast,miss,dodge,parry,buff_applied,buff_removed&limit=10000'],capture_output=True,text=True).stdout
        if '"success":true' in d: open(fn,'w').write(d); break
        time.sleep(30*(a+1))
    time.sleep(2)
print('done',file=sys.stderr)

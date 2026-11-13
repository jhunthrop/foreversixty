"""rakeap.py <dl root>: every non-crit Rake hit and tick, Claw hit and white swing with the caster's latest advanced attack power."""
import sys, glob, os
import pyarrow.parquet as pq
C=['time_unix_nano','event','source_guid','source_name','spell_id','spell_name','base_amount','critical','glancing','adv_present','adv_info_guid','adv_attack_power','adv_level']
root=sys.argv[1]
for rep in sorted(os.listdir(root)):
    for fp in sorted(glob.glob(f'{root}/{rep}/*.parquet'), key=lambda p:int(p.split('/f')[-1].split('.')[0])):
        ap={}
        for e in pq.read_table(fp,columns=C).to_pylist():
            g=e['source_guid']
            if not g.startswith('Player'): continue
            if e['adv_present'] and e['adv_info_guid']==g and e['adv_attack_power']: ap[g]=(e['adv_attack_power'],e['adv_level'])
            if e['spell_name']=='Rake' and e['base_amount'] and not e['critical']:
                print(rep, e['source_name'][:12], ap.get(g), e['event'][6:10], e['spell_id'], e['base_amount'])

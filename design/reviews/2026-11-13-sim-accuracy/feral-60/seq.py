"""seq.py <dir> <name prefix> [max]: event sequence for one player."""
import sys, glob
import pyarrow.parquet as pq
d, who = sys.argv[1], sys.argv[2]; mx = int(sys.argv[3]) if len(sys.argv) > 3 else 200
C = ['time_unix_nano','event','source_name','dest_name','spell_id','spell_name','amount','base_amount','critical','power_type','over_energize','adv_present','adv_info_guid','source_guid','adv_attack_power','adv_power_type','adv_current_power','adv_power_cost','adv_level','miss_type','stacks','glancing']
n = 0
for fp in sorted(glob.glob(f'{d}/*.parquet'), key=lambda p: int(p.split('/f')[-1].split('.')[0])):
    for e in pq.read_table(fp, columns=C).to_pylist():
        if not (e['source_name'].startswith(who) or e['dest_name'].startswith(who)): continue
        if e['source_name'].startswith(who) or e['event'] in ('SPELL_ENERGIZE','SPELL_AURA_APPLIED'):
            adv = (e['adv_attack_power'], e['adv_current_power'], e['adv_power_cost']) if e['adv_present'] and e['adv_info_guid'] == e['source_guid'] else ''
            print(fp.split('/')[-1], e['time_unix_nano']//10**6 % 10**7, e['event'], e['spell_id'], e['spell_name'], e['dest_name'][:14], e['amount'], e['base_amount'], 'C' if e['critical'] else '', e['power_type'], e['miss_type'], e['stacks'], adv)
            n += 1
            if n >= mx: sys.exit()

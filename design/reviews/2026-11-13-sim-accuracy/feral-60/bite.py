"""bite.py <dl root>: every Ferocious Bite and Rip in our reports with the combo points gained on that
target since the last finisher (SPELL_ENERGIZE power 4), the energy before the bite (last advanced
reading after a spend plus 10 a second, capped 100), base_amount and crit."""
import sys, os, glob, collections
import pyarrow.parquet as pq
C = ['time_unix_nano','event','source_guid','source_name','dest_guid','spell_id','spell_name','amount','base_amount','critical','power_type','adv_present','adv_info_guid','adv_current_power','adv_power_cost','adv_attack_power','adv_level','miss_type']
root = sys.argv[1]
for rep in sorted(os.listdir(root)):
    for fp in sorted(glob.glob(f'{root}/{rep}/*.parquet')):
        t = pq.read_table(fp, columns=C).to_pylist()
        cp = collections.defaultdict(lambda: collections.Counter())  # player -> target -> cp
        energy = {}  # player -> (time, value)
        last_tgt = {}
        pend = {}
        for i, e in enumerate(t):
            g = e['source_guid']
            if not g.startswith('Player'): continue
            ts = e['time_unix_nano'] / 1e9
            if e['adv_present'] and e['adv_info_guid'] == g and e['adv_current_power'] is not None and e['event'] != 'SPELL_CAST_SUCCESS':
                energy[g] = (ts, e['adv_current_power'], e['adv_attack_power'], e['adv_level'])
            if e['event'] == 'SPELL_CAST_SUCCESS' and e['spell_name'] in ('Claw','Rake','Shred','Ravage','Pounce'):
                last_tgt[g] = e['dest_guid']
            if e['event'] == 'SPELL_ENERGIZE' and e['power_type'] == 4:
                # a builder's point is logged just before its own cast: credit it to that cast's target
                tg = next((x['dest_guid'] for x in t[i:i+6] if x['source_guid'] == g and x['event'] == 'SPELL_CAST_SUCCESS'
                           and x['spell_name'] == e['spell_name'] and abs(x['time_unix_nano'] - e['time_unix_nano']) < 5e6), last_tgt.get(g))
                if tg: cp[g][tg] = min(5, cp[g][tg] + (e['amount'] or 0))
            if e['event'] == 'SPELL_CAST_SUCCESS' and e['spell_name'] in ('Ferocious Bite','Rip'):
                n = cp[g][e['dest_guid']]
                en = energy.get(g)
                est = min(100, en[1] + 10 * (ts - en[0])) if en else None
                pend[(g, e['spell_name'])] = (n, est, en[2] if en else None, en[3] if en else None)
                cp[g] = collections.Counter()
            if e['event'] in ('SPELL_DAMAGE','SPELL_PERIODIC_DAMAGE') and e['spell_name'] in ('Ferocious Bite','Rip') and (g, e['spell_name']) in pend:
                n, est, ap, lv = pend[(g, e['spell_name'])]
                print(rep, os.path.basename(fp), e['source_name'][:16], lv, e['spell_name'], e['spell_id'], e['event'][6:9], 'cp', n, 'E', None if est is None else round(est), 'advAP', ap, e['base_amount'], e['amount'], 'C' if e['critical'] else '-')
                if e['spell_name'] == 'Ferocious Bite': del pend[(g, e['spell_name'])]

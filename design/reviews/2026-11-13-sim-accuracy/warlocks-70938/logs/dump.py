"""dump.py <dl dir> <report> <guid suffix> [spell substring]: one unit's events (and its pets' by owner) in time order."""
import glob, sys
import pandas as pd, pyarrow.parquet as pq
dl, rep, suf = sys.argv[1:4]; filt = sys.argv[4] if len(sys.argv) > 4 else None
out = []
for f in sorted(glob.glob(f'{dl}/{rep}/f*.parquet'), key=lambda p: int(p.split('/f')[-1].split('.')[0])):
    t = pq.read_table(f).to_pandas()
    m = t.source_guid.astype(str).str.endswith(suf)
    if 'owner_guid' in t:
        m |= t.owner_guid.astype(str).str.endswith(suf)
    t = t[m].copy(); t['fight'] = int(f.split('/f')[-1].split('.')[0])
    out.append(t)
df = pd.concat(out)
if filt:
    df = df[df.spell_name.astype(str).str.contains(filt, case=False)]
df['t'] = (df.time_unix_nano - df.time_unix_nano.min()) / 1e9
keep = [c for c in ['fight', 't', 'event', 'source_name', 'spell_id', 'spell_name', 'dest_name', 'amount', 'base_amount', 'critical', 'miss_type', 'adv_current_power', 'adv_power_cost', 'adv_level', 'aura_type'] if c in df]
pd.set_option('display.width', 250); pd.set_option('display.max_rows', 5000); pd.set_option('display.max_columns', 30)
print(df[keep].to_string(index=False))

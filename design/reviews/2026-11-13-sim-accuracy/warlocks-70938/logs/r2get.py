"""r2get.py <out dir> <report...>: download each report's report.json and the events of every fight a warlock
(or a warlock's pet) is in, read only, from the logs bucket. Run through r2env.sh."""
import boto3, os, sys, json
s3 = boto3.client('s3', endpoint_url='https://%s.r2.cloudflarestorage.com' % os.environ['R2_ACCOUNT_ID'],
                  aws_access_key_id=os.environ['R2_ACCESS_KEY_ID'], aws_secret_access_key=os.environ['R2_SECRET_ACCESS_KEY'],
                  region_name='auto')
B = os.environ['R2_BUCKET']
out = sys.argv[1]
def get(key, dst):
    if os.path.exists(dst) and os.path.getsize(dst) > 0:
        return
    os.makedirs(os.path.dirname(dst), exist_ok=True)
    s3.download_file(B, key, dst)
for rep in sys.argv[2:]:
    get(f'reports/{rep}/report.json', f'{out}/{rep}/report.json')
    d = json.load(open(f'{out}/{rep}/report.json'))
    # A report without a unit list (units null) carries no classes: take every fight.
    wl = {u['guid'] for u in d['units'] or [] if u.get('class') == 'Warlock'}
    fs = [f['index'] for f in d['fights'] if d['units'] is None or wl & set(f.get('players', []))]
    print(rep, len(wl), 'warlocks', len(fs), 'of', len(d['fights']), 'fights', flush=True)
    for i in fs:
        get(f'reports/{rep}/fights/{i}/events.parquet', f'{out}/{rep}/f{i}.parquet')

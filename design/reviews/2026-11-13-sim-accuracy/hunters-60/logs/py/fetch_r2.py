"""fetch_r2.py <out_root> <report_id>...: read-only GETs of report.json and every fight's summary/events from the
foreversixty-logs R2 bucket (private reports). Credentials come from GSM at runtime and are never printed or written."""
import sys, os, json, subprocess
import boto3
from concurrent.futures import ThreadPoolExecutor
def secret(name):
    return subprocess.run(['gcloud', 'secrets', 'versions', 'access', 'latest', f'--secret={name}', '--project', 'foreversixty'],
                          check=True, capture_output=True, text=True).stdout.strip()
ACCOUNT = '1f1504a52d74d323c0a5648a766e0871'; BUCKET = 'foreversixty-logs'
s3 = boto3.client('s3', endpoint_url=f'https://{ACCOUNT}.r2.cloudflarestorage.com', region_name='auto',
                  aws_access_key_id=secret('R2_ACCESS_KEY_ID'), aws_secret_access_key=secret('R2_SECRET_ACCESS_KEY'))
def get(key, path):
    if os.path.exists(path) and os.path.getsize(path): return True
    os.makedirs(os.path.dirname(path), exist_ok=True)
    try: s3.download_file(BUCKET, key, path); return True
    except Exception as e: print('miss', key, type(e).__name__, file=sys.stderr); return False
out = sys.argv[1]
for rid in sys.argv[2:]:
    if not get(f'reports/{rid}/report.json', f'{out}/{rid}/report.json'): continue
    jobs = [(f'reports/{rid}/fights/{f["index"]}/{name}', f'{out}/{rid}/fights/{f["index"]}/{name}')
            for f in json.load(open(f'{out}/{rid}/report.json'))['fights'] for name in ('summary.json', 'events.parquet')]
    with ThreadPoolExecutor(12) as pool: list(pool.map(lambda j: get(*j), jobs))
    print(rid, 'done', flush=True)

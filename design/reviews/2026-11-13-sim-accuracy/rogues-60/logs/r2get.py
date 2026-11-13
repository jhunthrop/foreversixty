import boto3, os, sys, json
s3=boto3.client('s3',endpoint_url='https://%s.r2.cloudflarestorage.com'%os.environ['R2_ACCOUNT_ID'],aws_access_key_id=os.environ['R2_ACCESS_KEY_ID'],aws_secret_access_key=os.environ['R2_SECRET_ACCESS_KEY'],region_name='auto')
B=os.environ['R2_BUCKET']
def get(key,out):
    if os.path.exists(out) and os.path.getsize(out)>0: return
    os.makedirs(os.path.dirname(out),exist_ok=True)
    s3.download_file(B,key,out)
for rep in sys.argv[1:]:
    get(f'reports/{rep}/report.json',f'dl/{rep}/report.json')
    d=json.load(open(f'dl/{rep}/report.json'))
    rog={u['guid'] for u in d['units'] if u.get('class')=='Rogue'}
    fs=[f['index'] for f in d['fights'] if rog & set(f.get('players',[]))]
    print(rep,len(rog),'rogues',len(fs),'of',len(d['fights']),'fights',flush=True)
    for i in fs: get(f'reports/{rep}/fights/{i}/events.parquet',f'dl/{rep}/f{i}.parquet')

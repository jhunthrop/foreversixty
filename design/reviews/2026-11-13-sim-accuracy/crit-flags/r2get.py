"""Download every fight's events of the given reports from R2 (read only).

Usage: python3 -I r2get.py <out dir> <report id>...  (R2_* env vars set)
"""
import os
import sys
from concurrent.futures import ThreadPoolExecutor

import boto3

s3 = boto3.client("s3", endpoint_url=f"https://{os.environ['R2_ACCOUNT_ID']}.r2.cloudflarestorage.com",
                  aws_access_key_id=os.environ["R2_ACCESS_KEY_ID"],
                  aws_secret_access_key=os.environ["R2_SECRET_ACCESS_KEY"], region_name="auto")
BUCKET = os.environ["R2_BUCKET"]


def get(key, out):
    if os.path.exists(out) and os.path.getsize(out) > 0:
        return
    os.makedirs(os.path.dirname(out), exist_ok=True)
    try:
        s3.download_file(BUCKET, key, out)
    except Exception as exc:  # a fight without events is reported, not fatal
        print("missing", key, exc, flush=True)


def keys(prefix):
    pages = s3.get_paginator("list_objects_v2").paginate(Bucket=BUCKET, Prefix=prefix)
    return [o["Key"] for page in pages for o in page.get("Contents", [])]


def main(out_dir, reports):
    for rep in reports:
        found = keys(f"reports/{rep}/")
        jobs = [(k, f"{out_dir}/{rep}/report.json") for k in found if k.endswith("/report.json")]
        jobs += [(k, f"{out_dir}/{rep}/f{k.split('/')[-2]}.parquet") for k in found if k.endswith("/events.parquet")]
        with ThreadPoolExecutor(8) as pool:
            for key, out in jobs:
                pool.submit(get, key, out)
        print(rep, len(jobs), "files", flush=True)


if __name__ == "__main__":
    main(sys.argv[1], sys.argv[2:])

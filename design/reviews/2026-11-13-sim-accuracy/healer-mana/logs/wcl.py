"""Cached, rate-limited Warcraft Logs v1 reads (vanilla.warcraftlogs.com, Classic Era).

The key comes from the WCL_KEY environment variable only. It is passed as a request
parameter and never written to a cache file, a log line or an exception message.
"""
import hashlib
import json
import os
import sys
import time
import urllib.error
import urllib.parse
import urllib.request

BASE = "https://vanilla.warcraftlogs.com/v1/"
CACHE = os.environ.get("WCL_CACHE") or os.path.join(os.path.dirname(os.path.abspath(__file__)), "cache")
PAUSE_SECONDS = 1.2


def _key():
    key = os.environ.get("WCL_KEY", "")
    if not key:
        sys.exit("WCL_KEY is not set")
    return key


def get(path, **params):
    """GET BASE+path with params; the cache name is the path and params without the key."""
    shown = path + ("?" + urllib.parse.urlencode(sorted(params.items())) if params else "")
    name = hashlib.sha1(shown.encode()).hexdigest() + ".json"
    file = os.path.join(CACHE, name)
    if os.path.exists(file):
        with open(file) as f:
            return json.load(f)["body"]
    query = urllib.parse.urlencode({**params, "api_key": _key()})
    for attempt in range(4):
        time.sleep(PAUSE_SECONDS * (1 + 4 * attempt))
        try:
            with urllib.request.urlopen(BASE + path + "?" + query, timeout=60) as resp:
                body = json.load(resp)
            break
        except urllib.error.HTTPError as err:
            if err.code in (429, 500, 502, 503) and attempt < 3:
                print(f"retry {err.code} {shown}", file=sys.stderr)
                continue
            raise RuntimeError(f"HTTP {err.code} for {shown}") from None
        except urllib.error.URLError as err:
            if attempt < 3:
                continue
            raise RuntimeError(f"network error for {shown}: {err.reason}") from None
    os.makedirs(CACHE, exist_ok=True)
    with open(file, "w") as f:
        json.dump({"request": shown, "fetched": time.strftime("%Y-%m-%d"), "body": body}, f)
    return body

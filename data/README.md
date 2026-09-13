# data

Pipeline that turns a WoW client build into the JSON the site reads.

## Commands
```
uv sync
uv run python -m pipeline fetch --product wow_classic_era            # latest build
uv run python -m pipeline fetch --product <forever product> --build <build>
uv run python -m pipeline normalize --build <build>
uv run python -m pipeline diff --from <build> --to <build>
uv run pytest
```

## Layout
```
builds/<build>/raw/*.csv     downloaded DB2 exports (gitignored)
builds/<build>/*.json        normalized entities (committed)
builds/<build>/manifest.json build, product, fetched_at, sha256 per file
diffs/<from>__<to>.json      added / removed / changed per entity
```

## Sept 17 checklist
1. Find the Forever product key on https://wago.tools/builds (it will appear when the beta client is on the CDN).
2. Run the workflow_dispatch in GitHub Actions with that product, or run fetch + normalize locally.
3. If any table 404s or a column is missing, fix TABLES or the normalizer, add a fixture row, keep the golden tests green.
4. Commit the normalized JSON; the site build picks it up.

## Known gaps
The Classic Era client has no `JournalInstance` table, so `dungeons.json` is empty for `wow_classic_era` builds. `JournalInstance` is the only entry in `OPTIONAL_TABLES` in `pipeline/wago.py`; any other missing table fails the fetch on purpose.

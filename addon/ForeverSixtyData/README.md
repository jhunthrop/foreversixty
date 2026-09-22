# Forever Sixty Data

The nightly data pack the Forever Sixty addon reads for in-game ratings and guild standings.

- `Data.lua` here is the empty shape; the site's nightly job writes the real file and the release workflow publishes it to CurseForge (project 1706707) and Wago Addons (project ANz7R564) every night.
- No code lives here. Every rule about how the data is read belongs to the main addon.
- The listing copy is in `design/listing/data-addon-description.md`.

## Where Data.lua comes from

`api/internal/dataaddon` (the `api` module's `data-addon` Cloud Run job) reads the database
nightly, aggregates every public rated character's last 90 days and every guild's verified
roster, and writes the real `Data.lua` to a Cloud Storage bucket. `addon-data-release.yml`
downloads it and runs this package through the same release pipeline `addon-release.yml`
uses for the main addon. See `api/README.md`'s "The data-addon job" section for the
production setup, and `docs/superpowers/plans/2026-09-21-data-addon.md` for the exact
aggregation rule and the region/ruleset-as-realm key mapping.

# Data sources after the wago.tools block (2026-11-18)

Read-only research. Live probes were single requests: Blizzard's version service, the CDN configs and one HEAD, and one 7.7 MB DBCache.bin download. No logins and no credentials were used.

## Findings

**1. Blizzard CDN via TACT/CASC (DB2 tables): use this as the primary source for (a).**
- The product is `wowf`. `https://us.version.battle.net/wowf/versions` returns us/eu `1.60.1.70959`, build config `ad887adc...`, CDN config `4831465d...`, with an empty KeyRing. `wowf/cdns` serves `us.cdn.blizzard.com/tpr/wow` on plain HTTP(S) with no auth. The build config fetched from the CDN reads `build-uid = wowf`, `build-name = WOW-70959patch1.60.1_Retail`.
- Size for a CI run:
  - The encoding file is 141 MB (HEAD on the CDN).
  - The CDN config lists 1405 archives whose indexes total 125 MB. These are content-addressed, so cache them between runs and later builds fetch only new ones.
  - Then the DB2s themselves, fetched by range from the archives (tens of MB for our tables).
  - The first run is roughly 300 MB. Later builds are about 150 MB.
- Tools:
  - TACTSharp/TACTTool (MIT, .NET, CLI extraction by fdid/name, loads online installs; https://github.com/wowdev/TACTSharp) runs on Linux CI.
  - CascLib (C, online mode) works too.
  - Field layouts come from WoWDBDefs (https://github.com/wowdev/WoWDBDefs), which the repo's `pipeline/db2.py` reader can keep using.
- Encryption: the KeyRing is empty, and TACT-encrypted DB2 sections are only used for unreleased content. Item/Spell/Talent rows that are live need no keys. TACTSharp lists encrypted products as a TODO. That is a product-level gap (`wowf` is not encrypted), not a section-level one.
- Accuracy: these are exactly the shipped tables, so the CDN gives no hotfixes (b).
- Terms: this is the public CDN the launcher uses, the same path wago.tools and wow.export use, with no account involved. Low risk if we fetch once per build and cache.

**2. Hotfixes (DBCache.bin): use Raidbots' public cache as the primary source for (b).**
- Only a logged-in client receives the push. Blizzard offers no login-free endpoint for hotfixes.
- Marlamin's wow.tools.local reads DBCache.bin from a WoW folder. Since 0.9.1 (CHANGELOG, 2026-01-31) it downloads `https://storage.googleapis.com/raidbots-static/wow/{branch}/enUS/DBCache.bin` and has a `forever` branch (`Managers/HotfixManager.cs` and `wwwroot/dbc/hotfixes.html`, https://github.com/Marlamin/wow.tools.local).
- Tested:
  - The `forever` file is live: 7,673,299 bytes, last-modified 2026-11-18 01:22 GMT.
  - Its header reads XFTH v9, build 70959. The repo's `pipeline.hotfix_cache.read_cache` decodes it cleanly: 94,979 records, 6,179 pushes, 247 tables.
  - Issue #27 is fixed in this file. ItemSparse 870 (Fiery War Axe), push 16778086, is VALID and carries 3600 in the **enUS** cache, so we do not depend on other locales.
  - Item 2801 has a newer push, 16780017.
  - The `deDE` path returns 404, so the cache is enUS only.
- Coverage: it holds only what one Raidbots client session received. Removals and invalidations are present (12,534 and 36,917 records). Refresh cadence looks daily, not guaranteed.
- Terms: it is a public GCS bucket with no stated license. Ask Raidbots (Seriallos) before relying on it, and keep it to one GET per run with If-Modified-Since.
- Fallback: the owner's own client's DBCache.bin, which the existing `hotfixes` subcommand already reads. Automating a game login to harvest the push would breach Blizzard's EULA ban on bots and unauthorized automation (https://www.blizzard.com/en-us/legal/fba4d00f-c7e4-4883-b8b9-1b4500a402ea/blizzard-end-user-license-agreement), so do not do that.

**3. Blizzard Game Data API: a cross-check for items only.**
- The repo's profile import uses `profile-forever-us`. Responses link `static-1.60.1_70399-forever-us` (see `api/internal/bnetapi/static.go`; plain `static-forever-us` answers 403).
- The Item endpoint's `preview_item.weapon.attack_speed` and its stats come from Blizzard's server data, so it should reflect hotfixes. This is unverified for Forever: testing it needs a client token, which I did not use.
- As far as I know, the Classic Game Data docs list no Spell or Talent endpoint, so this source probably does not cover spells or talents. That comes from memory: the docs page did not render, so it is unconfirmed.
- Limits are 36,000 requests per hour and 100 per second (https://us.forums.blizzard.com/en/blizzard/t/api-access-clients-rate-limits/5602). The Developer API Terms allow use only through the API and forbid scraping elsewhere. Lowest ToS risk of all sources.

**4. Wowhead Forever: a spot-check only.**
- In round 5 its tooltips matched the server push (`pipeline/wago_hotfixes.py` doc), so it does include hotfixed values.
- It has no public API, and Wowhead discourages parsing (https://www.wowhead.com/ptr/forums/topic/wowhead-api-156427). Bulk crawling would repeat the wago mistake.
- Keep it for item and talent spot checks only.

**5. Other options.**
- wago.tools by agreement: email wago for permission or an API key. Its `/hotfixes` data is the same community DBCache feed, so it adds no authority over Raidbots.
- wow.tools.local in read-only/hosted mode packages the CDN and Raidbots pair, but it is a heavy .NET web app for CI.
- I did not probe Townlong Yak; it has no documented Forever feed.

## Ranked recommendation
1. **Primary for (a): the Blizzard CDN.** Run TACTTool or CascLib in CI on Linux, key on `wowf/versions` (clientwatch already polls it), cache archive indexes and the encoding file per build, and extract the DB2 fdids we already list.
2. **Primary for (b): Raidbots `forever/enUS/DBCache.bin`.** Decode it with the existing `pipeline.hotfix_cache` and layer it with the existing `layer_rows` rule. Gate on header build == target build, ask Raidbots for permission, and fall back to the owner's own DBCache.bin.
3. **Cross-check:** a nightly sample of hotfixed item ids against the Game Data Item API (`static-<build>-forever-us`), comparing speed, damage and stats. Wowhead stays a manual spot check for spells and talents.
4. Retire the wago crawls (`hotfix_push_store`, `item_push`, `wago_hotfixes`, `casc.py` via wago) once 1 and 2 land.

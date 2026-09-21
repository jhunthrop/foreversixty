# Guild-addon lane: FS1 gains a `guild` section — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development
> (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps
> use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Give the addon export a `guild` fact — name and rank index — as a new optional
FS1 version-2 section, read by both the Lua codec and the web decoder, with shared codec
vectors and spike-checklist rows for the human tester.

**Architecture:** The FS1 wire format already carries a fixed-order sequence of optional
pipe-delimited sections after the version-1 head (`bags`, `bank`, `sets`, `loadouts`,
`professions`); `guild` becomes the sixth, appended after `professions`. Its payload is
`<url-encoded-name>:<rank-index>` (one fact, not a named-entries collection). The web
decoder (`fs1.ts`) is the reference implementation the shared codec vectors are generated
from; the Lua codec (`Codec.lua`) is then tested against those same vectors, so the two
implementations cannot drift. `Export.lua` calls `GetGuildInfo("player")` once, through a
new `Export.guildInfo()`, and passes the result into the existing `Codec.encodeFS1` call.

**Tech Stack:** Lua 5.1 (WoW addon, tested with busted on Lua 5.4), TypeScript (web,
tested with vitest), Node (fixture generator script).

**Spec:** `docs/superpowers/specs/2026-09-21-guild-membership-design.md` — this plan
implements section 1 in full (the Addon+codec lane of section 6's lane split).

## Global Constraints

- **File ownership (this lane only):** `addon/ForeverSixty/Export.lua`, `Codec.lua`,
  `Locale.lua`; `web/src/lib/planner/fs1.ts` (the `guild` branch only — no other line of
  that file moves); `tools/gen-codec-vectors.mjs`; both `codec-vectors.json` fixture
  copies (`addon/tests/fixtures/codec-vectors.json`,
  `web/src/fixtures/addon/codec-vectors.json`); `addon/README.md` (spike rows + manual
  checklist line only); `addon/tests/codec_fs1_spec.lua`, `addon/tests/wow_mock.lua`;
  `web/src/lib/planner/fs1.test.ts`. One narrow addition outside that list is permitted
  and justified per-task below: `addon/.luacheckrc` gains `GetGuildInfo` to its
  `read_globals` list (Task 5) — required for `luacheck` to pass on this lane's own
  `Export.lua` change, touches no line any other lane depends on, and no other guild lane
  (API, Web) has any reason to touch it. Nothing under `api/` is touched. Nothing in
  `web/` other than the two files named above is touched — in particular, no addon UI
  file (`addon/ForeverSixty/views/`, `Window.lua`) changes.
- **`GetGuildInfo("player")`** returns `guildName, guildRankName, guildRankIndex[,
  guildRealm]`, or `nil` when unguilded. `guildRankIndex` is 0-based; 0 is always the
  guild master. Rank **names** are player free text and are never used — only the index.
  This return shape is unverified against the 1.60.1 beta client; the plan routes every
  use through `wow_mock.lua`'s stub and adds spike-checklist rows (Task 5) for a human to
  confirm it in-game. No step in this plan claims in-game behaviour was verified.
- **Wire syntax (FS1 v2, section 6, after `professions`):** payload
  `<url-encoded-guild-name>:<rank-index>`. The name is percent-encoded with the same
  `urlEncode`/`encodeURIComponent` functions `sets`/`loadouts` already use (which already
  escape `:`, so a literal colon in a guild name is always pre-escaped before this format
  ever splits on the first unescaped `:`). Decode: split the payload on the **first**
  `:`; URL-decode the left half with the existing `decodeName`/`urlDecode` helper (never
  throws); the right half must satisfy `isDigits` (Lua) / `/^\d+$/` (TS) or the whole code
  is refused with a new message — **not** silently dropped, because a malformed *known*
  section refuses the whole code (only an *unrecognised section name* is forgiven).
- **`FS1Build.guild?: { name: string; rankIndex: number }`** (TS) / `build.guild` (Lua,
  `nil` or a table) — always optional, **never** defaulted to an empty object when absent
  (unlike `bags`/`bank`/`sets`/`loadouts`/`professions`, which default to `[]`). "No
  guild" and "an empty guild" are different facts.
- **The encoder must never write a `guild=` section with an empty name.** This is
  guaranteed by `Export.guildInfo()` returning `nil` (not a table with an empty name) when
  `GetGuildInfo` returns `nil`.
- `MAX_CODE_LENGTH` (16,384 in both languages) is untouched; a guild section is under 100
  bytes.
- **Locale message**, mirroring `L.codecGearEntry`'s phrasing: `L.codecGuildRank = "That
  code has an unreadable guild rank: %s."` (Lua); the TS decoder uses the matching
  English text inline, the same way every other refusal message in `fs1.ts` does.
- **House rules:** functions under 50 lines, files under 800, no magic numbers, pure
  functions with explicit inputs, immutable updates (never mutate a passed-in table/array
  — build new ones), errors/refusals handled and named (never a generic "invalid code"),
  table-driven / vector-driven tests. DRY: the Lua and TS implementations share the
  fixture vectors so neither reimplements the other's test cases by hand.
- **Toolchains** (run before every commit that touches that language):
  - Addon, from `addon/`: `export PATH=$HOME/.luarocks/bin:$PATH` then
    `luacheck ForeverSixty tests` and `busted`. 341 specs pass on `main` today; the
    luacheckrc spec shells out to luacheck, so the `PATH` export is needed for `busted`
    too.
  - Web, from `web/`: `export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12`,
    then (once per session) `FOREVER_DATA=fixture npm run sync`, then
    `npx vitest run src/lib/planner src/fixtures/addon`, `npx astro check`,
    `npm run lint`, `npx prettier --check <paths touched>`.
  - The vector generator itself: from `web/`, `node ../tools/gen-codec-vectors.mjs` (it
    imports `web/src/lib/planner/fs1.ts` directly via a relative path and writes both
    fixture copies; run it from `web/` so the `tsx`/TypeScript loader config it depends on
    is in scope, exactly as the file's own header comment says).
  - `web/node_modules` is a symlink in this worktree: never run `npm install`, never
    stage it.
  - `addon/ForeverSixty/Data.lua` is generated: never hand-edit it (this plan never
    touches it).
- **Commits:** write the message to a file under this worktree's `.superpowers/` with
  `printf`, then run `git commit -F <file>` as its own Bash command with nothing else in
  it. Never a heredoc combined with `git commit`. Never `-n`. Never bare `git stash`.
  Conventional subjects. End every message with exactly:
  `Co-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>`
  `Claude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5`
- Never connect to the production database, never print a secret, never deploy.

---

## Task 1: `fs1.ts` gains the `guild` branch (encode + decode)

**Files:**
- Modify: `web/src/lib/planner/fs1.ts` (the `guild` branch only — `FS1Build.guild?`,
  `decodeFS1`'s section loop, `encodeFS1V2`'s section list; no other line moves)
- Test: `web/src/lib/planner/fs1.test.ts`

**Interfaces:**
- Consumes: nothing new — uses this file's own existing `Parsed<T>`, `decodeName`,
  `FS1Error` types and `SLOTS`/`Slot` imports, already present.
- Produces (for later tasks and other lanes): `FS1Build.guild?: { name: string;
  rankIndex: number }`; `decodeFS1` fills `build.guild` (present only when the code
  carries a `guild=` section, otherwise the key is absent so `JSON.stringify` — used by
  `tools/gen-codec-vectors.mjs` in Task 2 — omits it); `encodeFS1V2` emits `guild=` after
  `professions=` when `build.guild` is set.

- [ ] **Step 1: Write the failing tests**

Add to `web/src/lib/planner/fs1.test.ts`, inside the existing `describe('version 2
sections', ...)` block (after the `'reads the professions section'` test, so it sits next
to its sibling sections):

```typescript
  it('reads the guild section', () => {
    const decoded = decodeFS1(`${V1}|guild=Iron%20Vanguard:2`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.guild).toEqual({ name: 'Iron Vanguard', rankIndex: 2 });
  });

  it('leaves guild undefined, never an empty object, when the code carries no guild section', () => {
    const decoded = decodeFS1(V1);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.guild).toBeUndefined();
  });

  it('refuses a non-numeric guild rank rather than silently dropping the section', () => {
    const decoded = decodeFS1(`${V1}|guild=Iron%20Vanguard:officer`);
    expect(decoded.ok).toBe(false);
    if (decoded.ok) return;
    expect(decoded.message).toBe('That code has an unreadable guild rank: officer.');
  });

  it('splits the guild payload on the first colon, so a name containing one still reads (pre-escaped)', () => {
    // A literal colon in a guild name is always percent-encoded by encodeURIComponent
    // before it reaches this format (RFC 3986's unreserved set excludes ':'), so this
    // is exercising the decoder's own split rule, not a real un-escaped name.
    const decoded = decodeFS1(`${V1}|guild=A%3AB:3`);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(decoded.build.guild).toEqual({ name: 'A:B', rankIndex: 3 });
  });
```

Add to `describe('encodeFS1V2', ...)`, after the `'writes the sections in the contract’s
order and round-trips them'` test:

```typescript
  it('writes the guild section after professions and round-trips it', () => {
    const build = {
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: ['engineering'],
      guild: { name: 'Iron Vanguard', rankIndex: 2 },
      ignored: [],
    };
    const code = encodeFS1V2(build);
    expect(code.indexOf('|professions=')).toBeLessThan(code.indexOf('|guild='));
    expect(code).toContain('|guild=Iron%20Vanguard:2');

    const decoded = decodeFS1(code);
    expect(decoded.ok).toBe(true);
    if (decoded.ok) expect(decoded.build.guild).toEqual({ name: 'Iron Vanguard', rankIndex: 2 });
  });

  it('omits the guild section entirely when build.guild is absent', () => {
    const code = encodeFS1V2({
      dataBuild: '1.15.9',
      classSlug: 'warrior',
      raceSlug: 'orc',
      treeRanks: [[], [], []],
      gear: {},
      bags: [],
      bank: [],
      sets: [],
      loadouts: [],
      professions: [],
      ignored: [],
    });
    expect(code).not.toContain('guild=');
  });
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd web && export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12; npx
vitest run src/lib/planner/fs1.test.ts`
Expected: FAIL — `build.guild` does not exist on the decoded type / is never set (new
tests fail; every pre-existing test in the file still passes).

- [ ] **Step 3: Implement the minimal code to make the tests pass**

In `web/src/lib/planner/fs1.ts`:

1. Add `guild?` to the `FS1Build` interface, in the version-2 block (after
   `professions?: string[];`, before `ignored?: string[];`):

```typescript
  /**
   * The character's current guild, from `GetGuildInfo("player")` (section 1.1 of the
   * guild-membership design). Absent means unguilded -- never defaulted to `{}`, the
   * way `bags`/`bank`/`sets`/`loadouts`/`professions` default to `[]`, because "no
   * guild" and "an empty guild" are not the same fact.
   */
  guild?: { name: string; rankIndex: number };
```

2. Add a `parseGuild` helper, near `parseNamed` (after it, before `decodeFS1`):

```typescript
/**
 * `<name>:<rank-index>`, split on the first colon -- a guild is one fact, not the
 * `name=payload;…` collection grammar `sets`/`loadouts` use. The name side never throws
 * (decodeName); the rank side must be digits-only or the whole code is refused, since a
 * malformed *known* section refuses the whole code (only an unrecognised section name is
 * forgiven).
 */
function parseGuild(field: string): Parsed<{ name: string; rankIndex: number }> {
  const at = field.indexOf(':');
  const namePart = at === -1 ? field : field.slice(0, at);
  const rankPart = at === -1 ? '' : field.slice(at + 1);
  if (!/^\d+$/.test(rankPart)) {
    return { ok: false, message: `That code has an unreadable guild rank: ${rankPart}.` };
  }
  return { ok: true, value: { name: decodeName(namePart), rankIndex: Number.parseInt(rankPart, 10) } };
}
```

3. In `decodeFS1`, add `guild: undefined as { name: string; rankIndex: number } | undefined,`
   to the `build` object literal, directly after `professions: [] as string[],` and before
   `ignored: [] as string[],` (keeps the key present-but-undefined so `guild` can be
   assigned below without widening the inferred object type; `JSON.stringify` still omits
   an `undefined`-valued key, which is what keeps the generated fixture vector honest).

4. In the section-loop `if/else if` chain, add a branch for `'guild'` directly after the
   `'professions'` branch and before the final `else if (name !== '')` (unknown-section)
   branch:

```typescript
    } else if (name === 'guild') {
      const guild = parseGuild(field);
      if (!guild.ok) return guild;
      build.guild = guild.value;
```

5. In `encodeFS1V2`, add a `guild` line after the `professions` push (after `if
   (professions.length > 0) sections.push(...)`, before the `slots =
   build.gearSlots ?? gearSlotsFrom(build.gear)` line):

```typescript
  if (build.guild) sections.push(`guild=${encodeURIComponent(build.guild.name)}:${build.guild.rankIndex}`);
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd web && npx vitest run src/lib/planner/fs1.test.ts`
Expected: PASS, every test in the file (old and new).

- [ ] **Step 5: Type-check and lint this one file**

Run: `cd web && npx astro check` (whole-project check; must stay green) and `npm run
lint` and `npx prettier --check src/lib/planner/fs1.ts src/lib/planner/fs1.test.ts`
Expected: no new errors.

- [ ] **Step 6: Commit**

```bash
printf 'feat(web): FS1 gains an optional guild section\n\nAdds FS1Build.guild? and a guild branch to decodeFS1/encodeFS1V2 -- section 6\nof the wire format, after professions. Payload is <url-encoded-name>:<rank-index>;\na malformed rank refuses the whole code with a new message, matching every other\nknown-section refusal in this decoder.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > /Users/jh/code/forever/.worktrees/guild-addon/.superpowers/commit-msg-1.txt
cd /Users/jh/code/forever/.worktrees/guild-addon
git add web/src/lib/planner/fs1.ts web/src/lib/planner/fs1.test.ts
git commit -F .superpowers/commit-msg-1.txt
```

---

## Task 2: Shared codec vectors gain the two guild cases

**Files:**
- Modify: `tools/gen-codec-vectors.mjs`
- Generated (not hand-edited): `addon/tests/fixtures/codec-vectors.json`,
  `web/src/fixtures/addon/codec-vectors.json`

**Interfaces:**
- Consumes: `decodeFS1` from Task 1 (imported by the generator).
- Produces: two new fixture vectors, `fs1` entry named `'version 2, guild'` and
  `fs1Invalid` entry named `'a non-numeric guild rank'`, read by Task 3's Lua tests and
  already covered on the TS side by Task 1's own tests (the generator itself asserts the
  site's decoder accepts/refuses these, so this task's own run is its test).

- [ ] **Step 1: Add the two vectors**

In `tools/gen-codec-vectors.mjs`, append to the `FS1_CODES` array (after the
`'version 2, professions'` entry, before the `'version 2, an unknown section is ignored
and named'` entry, so `guild` sits next to its sibling version-2 sections):

```javascript
  ['version 2, guild', 'FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:2'],
```

Append to the `INVALID` array (after `'too few fields'`):

```javascript
  ['a non-numeric guild rank', 'FS1:1.60.1.69893:warrior:tauren:0/0/0:|guild=Iron%20Vanguard:officer'],
```

- [ ] **Step 2: Regenerate both fixture copies**

Run: `cd web && export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12; node
../tools/gen-codec-vectors.mjs`
Expected output: two `wrote ...` lines, one per fixture path, no thrown error (a thrown
error means Task 1's decoder rejected the valid vector or accepted the invalid one —
stop and fix Task 1 first, do not hand-edit the fixture).

- [ ] **Step 3: Verify the two copies are byte-identical**

Run: `diff /Users/jh/code/forever/.worktrees/guild-addon/addon/tests/fixtures/codec-vectors.json /Users/jh/code/forever/.worktrees/guild-addon/web/src/fixtures/addon/codec-vectors.json`
Expected: no output (files identical). This is also asserted by the existing
`addon/tests/vectors_spec.lua` "are byte-identical in both lanes" test, run in Task 3.

- [ ] **Step 4: Confirm the new vectors are present**

Run: `grep -c "version 2, guild" /Users/jh/code/forever/.worktrees/guild-addon/addon/tests/fixtures/codec-vectors.json`
Expected: `1` (the vector's `name` field appears once; `code` may also mention `guild=`
but the grep targets the exact name string).

- [ ] **Step 5: Commit**

```bash
printf 'feat: add guild vectors to the shared codec fixtures\n\nRegenerates both codec-vectors.json copies from the sites decoder (now\nreading FS1s guild section, prior commit): one valid guild vector, one\nwith a non-numeric rank that the decoder refuses.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > /Users/jh/code/forever/.worktrees/guild-addon/.superpowers/commit-msg-2.txt
cd /Users/jh/code/forever/.worktrees/guild-addon
git add tools/gen-codec-vectors.mjs addon/tests/fixtures/codec-vectors.json web/src/fixtures/addon/codec-vectors.json
git commit -F .superpowers/commit-msg-2.txt
```

---

## Task 3: `Codec.lua` gains the `guild` branch, tested against the shared vectors

**Files:**
- Modify: `addon/ForeverSixty/Locale.lua` (one new key, `codecGuildRank`)
- Modify: `addon/ForeverSixty/Codec.lua` (`encodeFS1` gains the `guild=` section;
  `decodeFS1` gains the `guild` branch)
- Modify: `addon/.luacheckrc` (add `GetGuildInfo` to `read_globals`'s "Character" group —
  see Global Constraints for why this one file outside the lane table is touched)
- Test: `addon/tests/codec_fs1_spec.lua`

**Interfaces:**
- Consumes: `L.codecGuildRank` (new Locale key, this task defines it); the two new shared
  vectors from Task 2, read through `helper.vectors()` (already wired — `codec_fs1_spec.lua`
  iterates `vectors.fs1` and `vectors.fs1Invalid` in existing tests, which need one more
  field asserted per vector: `guild`).
- Produces: `Codec.encodeFS1` accepts `build.guild = { name = <string>, rankIndex =
  <number> }` (optional, nil default) and appends `guild=<urlEncode(name)>:<rankIndex>`
  after `professions`; `Codec.decodeFS1` fills `build.guild` (nil unless a `guild=`
  section was present) for `addon/ForeverSixty/Export.lua` (Task 4) to eventually feed
  back in, and for anything on `main` that already calls `Codec.decodeFS1`/`Codec.loadBuild`
  (neither reads `build.guild` today, so this is purely additive).

- [ ] **Step 1: Write the failing tests**

In `addon/tests/codec_fs1_spec.lua`, add a new `REFUSALS` entry (the vector's name must
match `tools/gen-codec-vectors.mjs`'s `INVALID` array exactly):

```lua
	["a non-numeric guild rank"] = string.format(L.codecGuildRank, "officer"),
```

Add `guild` to both vector-comparison loops' assertions — in `describe("decode", ...)`'s
`"reads every shared vector"` test, directly after the `professions` line:

```lua
				assert.are.same(vector.build.guild, build.guild, vector.name)
```

And in `describe("encode", ...)`'s `"re-encodes every shared vector to a canonical fixed
point"` test, directly after its `professions` line:

```lua
				assert.are.same(decoded.guild, redecoded.guild, vector.name)
```

Add two hand-written tests inside the existing `describe("decode", ...)` block, directly
after its `"refuses a code past the length bound before parsing it"` test (so they join
their sibling decode-refusal tests; this keeps the block nesting unchanged, only adding
`it(...)` entries where tests already live) — the spec's own request for a case distinct
from "unknown section is ignored" (a malformed *known* section refuses the whole code):

```lua
		it("refuses a malformed guild rank with the new message, distinct from an ignored unknown section", function()
			local build, message = Codec.decodeFS1(
				"FS1:1:paladin:human:0/0/0:|guild=Iron%20Vanguard:officer"
			)
			assert.is_nil(build)
			assert.are.equal(string.format(L.codecGuildRank, "officer"), message)
		end)

		it("leaves guild nil, not an empty table, when no guild section is present", function()
			local build = assert(Codec.decodeFS1("FS1:1:paladin:human:0/0/0:"))
			assert.is_nil(build.guild)
		end)
```

Add two hand-written tests inside the existing `describe("encode", ...)` block, directly
after its `"url-encodes a set name that uses the grammar's own punctuation"` test (same
rule: joining sibling encode tests, no new nesting):

```lua
		it("appends a guild section after professions and round-trips it", function()
			local code = Codec.encodeFS1({
				dataBuild = "1",
				classSlug = "paladin",
				raceSlug = "human",
				treeRanks = { {}, {}, {} },
				gearSlots = {},
				professions = { "enchanting" },
				guild = { name = "Iron Vanguard", rankIndex = 2 },
			})
			assert.are.equal("FS1:1:paladin:human:0/0/0:|professions=enchanting|guild=Iron%20Vanguard:2", code)
			local build = assert(Codec.decodeFS1(code))
			assert.are.same({ name = "Iron Vanguard", rankIndex = 2 }, build.guild)
		end)

		it("writes no guild section when build.guild is nil", function()
			local code = Codec.encodeFS1({
				dataBuild = "1",
				classSlug = "paladin",
				raceSlug = "human",
				treeRanks = { {}, {}, {} },
				gearSlots = {},
			})
			assert.is_nil(code:find("guild=", 1, true))
		end)
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && busted tests/codec_fs1_spec.lua`
Expected: FAIL — `Codec.decodeFS1`/`Codec.encodeFS1` do not know `guild` yet (new tests
fail; the two loop assertions fail on the new vector; pre-existing tests still pass).

- [ ] **Step 3: Implement the minimal code to make the tests pass**

In `addon/ForeverSixty/Locale.lua`, add one key directly after `codecStatPair` (still
inside the "Codec refusals" group, before the FSB1-only `codecEmptyField` block, since
`guild` is an FS1 refusal like its neighbours):

```lua
	codecGuildRank = "That code has an unreadable guild rank: %s.",
```

In `addon/ForeverSixty/Codec.lua`, in `Codec.encodeFS1`, add a `guild` section after the
`professions` block (after the `if build.professions and #build.professions > 0 then ...
end` block, before `if #sections == 0 then`):

```lua
	if build.guild and build.guild.name then
		sections[#sections + 1] = "guild=" .. urlEncode(build.guild.name) .. ":" .. tostring(build.guild.rankIndex)
	end
```

In `Codec.decodeFS1`'s section loop, add a `guild` branch after the `elseif name ==
"professions" then ... end` block, before the final `elseif name ~= "" then` (unknown
section) branch:

```lua
		elseif name == "guild" then
			local at = field:find(":", 1, true)
			local namePart = at and field:sub(1, at - 1) or field
			local rankText = at and field:sub(at + 1) or ""
			if not isDigits(rankText) then
				return nil, refuse(L.codecGuildRank, rankText)
			end
			build.guild = { name = urlDecode(namePart), rankIndex = tonumber(rankText) }
```

Do **not** add `guild = {}` (or any default) to the `build` table literal earlier in
`decodeFS1` — leaving the key entirely unset is what makes `build.guild` read as `nil`
for a code with no `guild=` section, matching the "never defaulted to an empty object"
rule.

In `addon/.luacheckrc`, add `"GetGuildInfo"` to the `read_globals` table's "Character"
group (same line as `"UnitClass", "UnitRace", ...`):

```lua
	"UnitClass", "UnitRace", "UnitLevel", "UnitName", "GetRealmName", "GetCurrentRegion",
	"GetProfessions", "GetProfessionInfo", "GetBuildInfo", "GetGuildInfo",
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && busted tests/codec_fs1_spec.lua
tests/vectors_spec.lua`
Expected: PASS, every test in both files (the byte-identical fixture check in
`vectors_spec.lua` passes because Task 2 already regenerated both copies).

- [ ] **Step 5: Lint**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && luacheck ForeverSixty tests`
Expected: no new warnings (the `GetGuildInfo` global is now declared; Task 4 is what
actually calls it from `Export.lua`, so until that task lands `luacheck` may report it as
unused in `read_globals` — `read_globals` entries are never flagged as unused by
luacheck, only `globals`, so this is fine either order).

- [ ] **Step 6: Commit**

```bash
printf 'feat(addon): Codec.lua reads and writes the guild section\n\nCodec.encodeFS1 appends guild=<name>:<rank-index> after professions when\nbuild.guild is set; Codec.decodeFS1 gains the matching branch, refusing a\nnon-numeric rank by name (L.codecGuildRank) rather than dropping the\nsection silently. Tested against the shared codec vectors (prior commit).\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > /Users/jh/code/forever/.worktrees/guild-addon/.superpowers/commit-msg-3.txt
cd /Users/jh/code/forever/.worktrees/guild-addon
git add addon/ForeverSixty/Locale.lua addon/ForeverSixty/Codec.lua addon/.luacheckrc addon/tests/codec_fs1_spec.lua
git commit -F .superpowers/commit-msg-3.txt
```

---

## Task 4: `Export.lua` reads `GetGuildInfo`, `wow_mock.lua` gains the stub

**Files:**
- Modify: `addon/ForeverSixty/Export.lua` (`Export.guildInfo()`, new; `Export.string`
  passes `guild = Export.guildInfo()`)
- Modify: `addon/tests/wow_mock.lua` (`GetGuildInfo` stub, install + uninstall)
- Test: `addon/tests/export_spec.lua`

**Interfaces:**
- Consumes: `Codec.encodeFS1` (Task 3) now honouring a `guild` field.
- Produces: `Export.guildInfo(): { name: string, rankIndex: number } | nil`, callable by
  any other addon code later (none does yet); `wow_mock`'s `state.guild` (a table with
  `name`, `rankName`, `rankIndex`, and optionally `realm`, or left unset/`nil` for the
  unguilded case, which is the default — no spec sets it unless it opts in).

- [ ] **Step 1: Write the failing tests**

Add to `addon/tests/wow_mock.lua`'s doc comment at the top of `mock.install` (the
parameter list comment, currently `-- @param state table with any of: talents, traits,
equipped, bags, itemStats, class, race, realm, region, professions, build`) — append
`, guild` to that list, since this task adds the field it documents.

Add to `addon/tests/export_spec.lua`, after the `"does not drop a profession sitting
behind an unlearned earlier slot"` test:

```lua
	it("returns nil from guildInfo when the character has no guild", function()
		character({})
		assert.is_nil(Export.guildInfo())
	end)

	it("returns the name and rank index from guildInfo when guilded", function()
		character({ guild = { name = "Iron Vanguard", rankName = "Officer", rankIndex = 2 } })
		assert.are.same({ name = "Iron Vanguard", rankIndex = 2 }, Export.guildInfo())
	end)

	it("writes no guild section for an unguilded character", function()
		character({})
		assert.is_nil(assert(Export.string(DATA)):find("|guild=", 1, true))
	end)

	it("writes the guild section, after professions, for a guilded character", function()
		character({
			professions = { 1 },
			professionNames = { "Enchanting" },
			guild = { name = "Iron Vanguard", rankName = "Officer", rankIndex = 2 },
		})
		local code = assert(Export.string(DATA))
		assert.is_truthy(code:find("|professions=enchanting|guild=Iron%20Vanguard:2", 1, true))
	end)
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && busted tests/export_spec.lua`
Expected: FAIL — `Export.guildInfo` does not exist; `GetGuildInfo` is not stubbed (new
tests fail; pre-existing tests still pass).

- [ ] **Step 3: Implement the minimal code to make the tests pass**

In `addon/tests/wow_mock.lua`, add the `GetGuildInfo` stub next to `GetProfessionInfo`
(after its definition, before `_G.GetBuildInfo`):

```lua
	-- Returns guildName, guildRankName, guildRankIndex, guildRealm; nil when the unit is
	-- not in a guild (the addon's own default: no spec here sets state.guild). Rank
	-- index is 0-based; 0 is always the guild master -- server-authoritative, the client
	-- never lets a non-GM report 0. Unverified against the 1.60.1 beta client; see
	-- addon/README.md's spike checklist rows 23/23a.
	_G.GetGuildInfo = function()
		if state.guild == nil then
			return nil
		end
		return state.guild.name, state.guild.rankName, state.guild.rankIndex, state.guild.realm
	end
```

Add `"GetGuildInfo"` to `mock.uninstall`'s name list (same line as `"GetProfessions",
"GetProfessionInfo",`):

```lua
		"GetRealmName", "GetCurrentRegion", "GetProfessions", "GetProfessionInfo", "GetGuildInfo",
```

In `addon/ForeverSixty/Export.lua`, add `Export.guildInfo()` after `Export.professionSlugs()`
(before `--- The export string, or nil and the reason.`):

```lua
--- The character's current guild, or nil when unguilded. GetGuildInfo returns nil for
--- an unguilded character (also the shape a login before guild data has loaded would
--- produce, per the WoW API's own documented behaviour -- both read the same way here:
--- no guild= section this export). Rank name (the second return) is player-chosen free
--- text and is never read; only the index is trustworthy (spike check 23, README.md).
function Export.guildInfo()
	local name, _, rankIndex = GetGuildInfo("player")
	if name == nil then
		return nil
	end
	return { name = name, rankIndex = rankIndex }
end
```

In `Export.string`, add `guild = Export.guildInfo(),` to the table passed to
`Codec.encodeFS1`, after the `professions = Export.professionSlugs(),` line:

```lua
		professions = Export.professionSlugs(),
		guild = Export.guildInfo(),
	})
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && busted tests/export_spec.lua`
Expected: PASS, every test in the file (old and new).

- [ ] **Step 5: Full addon suite and lint**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && luacheck ForeverSixty tests &&
busted`
Expected: `luacheck` clean; `busted` reports more than the baseline 341 specs, all
passing (the new tests added across Tasks 3 and 4).

- [ ] **Step 6: Commit**

```bash
printf 'feat(addon): Export.lua reads GetGuildInfo into the guild section\n\nExport.guildInfo() wraps GetGuildInfo("player"), returning nil when\nunguilded so Export.string never writes an empty guild= section.\nwow_mock.lua gains the matching stub (state.guild, unset by default, i.e.\nunguilded) so this is exercised without a live client.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > /Users/jh/code/forever/.worktrees/guild-addon/.superpowers/commit-msg-4.txt
cd /Users/jh/code/forever/.worktrees/guild-addon
git add addon/ForeverSixty/Export.lua addon/tests/wow_mock.lua addon/tests/export_spec.lua
git commit -F .superpowers/commit-msg-4.txt
```

---

## Task 5: `addon/README.md` spike checklist and manual checklist

**Files:**
- Modify: `addon/README.md`

**Interfaces:**
- Consumes: nothing (documentation only).
- Produces: nothing another task reads; this is the human tester's own checklist.

- [ ] **Step 1: Add the three spike-checklist rows**

In `addon/README.md`'s "Beta spike checklist" table, insert three rows immediately after
row 22 (the `Talent button mapping` row) and before the `## Findings` heading:

```markdown
| 23 | Guild info shape while in a guild | `/dump GetGuildInfo("player")` — expect `name, rankName, rankIndex[, realm]`; confirm `rankIndex` is `0` for the guild master | `Export.guildInfo()` reads positions 1 and 3; if the shape differs, fix there, not a flag |
| 23a | Guild info while unguilded | `/dump GetGuildInfo("player")` on a character with no guild — expect `nil` | Confirms `Export.guildInfo()` returns `nil` and `Export.string` writes no `guild=` section |
| 23b | Round trip | `/fs export` while in a guild, paste the code into the planner's import box on the site, confirm the guild name and rank index shown there match what `/dump GetGuildInfo("player")` reported | End-to-end check that `Codec.encodeFS1` and `fs1.ts`'s `decodeFS1` agree, beyond the fixture vectors |
```

- [ ] **Step 2: Add the manual checklist line**

In the "Manual checklist (before each release)" section, add one line after the existing
"Logging out writes `ForeverSixtyDB.characters` and `savedAt`..." item:

```markdown
- [ ] The Export tab's export includes a `|guild=` section for a guilded character and
      none for an unguilded one (`/fs diag` or `/dump` the saved string).
```

- [ ] **Step 3: Verify the table still renders (no broken pipes)**

Run: `sed -n '20,50p' /Users/jh/code/forever/.worktrees/guild-addon/addon/README.md`
Expected: the table prints with consistent `|`-delimited columns; rows 23/23a/23b sit
between row 22 and the `## Findings` heading, in that order.

- [ ] **Step 4: Commit**

```bash
printf 'docs(addon): spike-checklist rows for the guild section\n\nThree new rows (23/23a/23b) for a human tester to confirm GetGuildInfo\'s\nreturn shape, the unguilded nil case, and an end-to-end export/import round\ntrip, since none of that can be verified without the live client. One new\nmanual-checklist line for the guild= export line itself.\n\nCo-Authored-By: Claude Fable 5.1 <noreply@anthropic.com>\nClaude-Session: https://claude.ai/code/session_01GsMkaNo8ZBb3gdSBSRTSj5\n' > /Users/jh/code/forever/.worktrees/guild-addon/.superpowers/commit-msg-5.txt
cd /Users/jh/code/forever/.worktrees/guild-addon
git add addon/README.md
git commit -F .superpowers/commit-msg-5.txt
```

---

## Task 6: Whole-lane verification

**Files:** none modified unless a check below surfaces a regression, in which case the
fix lands in whichever file the failure names, staying inside this lane's ownership.

- [ ] **Step 1: Full addon suite**

Run: `cd addon && export PATH=$HOME/.luarocks/bin:$PATH && luacheck ForeverSixty tests &&
busted`
Expected: `luacheck` clean; every `busted` spec passes.

- [ ] **Step 2: Full web checks for the touched areas**

Run: `cd web && export NVM_DIR="$HOME/.nvm"; . "$NVM_DIR/nvm.sh"; nvm use 22.12;
FOREVER_DATA=fixture npm run sync && npx vitest run src/lib/planner src/fixtures/addon &&
npx astro check && npm run lint && npx prettier --check src/lib/planner/fs1.ts
src/lib/planner/fs1.test.ts tools/gen-codec-vectors.mjs`
Expected: all green.

- [ ] **Step 3: Confirm no out-of-scope files changed**

Run: `cd /Users/jh/code/forever/.worktrees/guild-addon && git diff --stat main...HEAD`
Expected: every changed path is one of: `addon/ForeverSixty/Export.lua`,
`addon/ForeverSixty/Codec.lua`, `addon/ForeverSixty/Locale.lua`, `addon/.luacheckrc`,
`addon/tests/codec_fs1_spec.lua`, `addon/tests/wow_mock.lua`, `addon/tests/export_spec.lua`,
`addon/README.md`, `web/src/lib/planner/fs1.ts`, `web/src/lib/planner/fs1.test.ts`,
`tools/gen-codec-vectors.mjs`, `addon/tests/fixtures/codec-vectors.json`,
`web/src/fixtures/addon/codec-vectors.json`, plus this plan file itself and the
`.superpowers/` commit-message scratch files. Nothing under `api/`; nothing in `web/`
beyond the two named files and the fixture.

- [ ] **Step 4: No further commit needed if Steps 1–3 are clean**

If any step required a fix, commit it with a `fix:` subject following the same commit
protocol as the tasks above, naming what regressed and why.

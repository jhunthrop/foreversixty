# Premium lane 1: page truth and the Supporter mark

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Every line on `/premium` and the legal pages is true today, and a paying member's Supporter mark shows on their characters and in guild rosters.

**Architecture:** Copy changes in `web/src/lib/billing/copy.ts` and three Markdown legal pages; one new boolean (`supporter`) on two existing public API payloads, computed by the existing `entitlements.Store.IsSupporter`; one new Svelte chip rendered in the character header and the roster table; the account page's billing row grows into a Premium block that lists each benefit with its live state.

**Tech Stack:** Go API (`api/internal/rankings`, `api/internal/guilds`, `api/internal/entitlements`), Astro + Svelte 5 web (`web/src`), vitest with the Astro container, Playwright e2e on the fixture build.

**Spec:** `docs/superpowers/specs/2026-10-24-premium-experience-design.md` (§4.1, §4.2, §4.5, §4.7's copy line, §5, §6.1), arguing from `docs/superpowers/specs/2026-09-21-entitlements-and-payments-design.md` §1.3 and §5.2.

## Global Constraints

- Copy: no exclamation marks, no urgency, "Premium" is a plan name. Every visible string lives in a `copy.ts`, never inline in a component.
- Public surfaces never upsell: a visitor's character page shows the chip or nothing, never plan copy.
- The crest is never altered; the Supporter mark is a text chip beside the name.
- Gating is server-side: the web renders what `GET /v1/me` and the public payloads say.
- Benefits not yet live carry their month: `Character history charts over time (ships November).` and `Discord notifications (ships December).` Never "soon".
- Legal pages: no `OWNER:` marker, no draft banner; `legal-markers.test.ts` stays green and grows.
- Every task runs the touched vitest files; the full web suite and the full Playwright suite plus the lhci shard run once before the lane merges (the header and the account page are shared layout).
- Commit messages in the repo's conventional style with the Co-Authored-By trailer; the owner merges the lane into `main`.

## Review Focus

1. A character with no linked account (never synced, or account deleted) must render no chip and no error: `supporter` is `false`, never missing or null.
2. A guild-plan member who is not an officer is a supporter (September spec §1.3); the chip must not depend on rank. Task 2's roster test pins a `member`-rank supporter.
3. A canceled subscription whose period has not ended still shows the chip (status `active`, `cancel_at_period_end` true); a `canceled` status does not. Task 2's test covers both rows.
4. The account page with an older `/v1/me` shape (no `entitlements`) must still render the block in its free state, as the existing billing row does with optional chaining.
5. The legal pages must not keep any sentence that promises a retention window; the grep in Task 1's test covers "ninety", "two-year", "two years", "retention falls back".

---

### Task 1: The page and the legal pages tell the truth

**Files:**
- Modify: `web/src/lib/billing/copy.ts` (`premiumFeatures`, `guildFeatures`, `faq`)
- Modify: `web/src/content/pages/terms.md:26-32`, `web/src/content/pages/privacy.md:36-39`
- Modify: `web/src/content/pages/legal-markers.test.ts`
- Modify: `web/src/lib/billing/copy.test.ts`, `web/tests/e2e/premium.spec.ts`, `web/tests/e2e/legal-pages.spec.ts`
- Create: `web/src/content/changelog/2026-10-24-premium-page-truth.md` (follow the collection's existing frontmatter; read one neighbour first)

**Interfaces:**
- Produces: `premiumCopy.premiumFeatures` is exactly these four strings, in this order:
  1. `Run sims on our servers (5,000 combinations, any precision, no tab left open).`
  2. `Character history charts over time (ships November).`
  3. `Discord notifications (ships December).`
  4. `A supporter mark on your profile.`
- Produces: `premiumCopy.guildFeatures` is exactly:
  1. `Everything in Premium, for everyone in the guild.`
  2. `Officer tools: the loot council helper and the raid readiness board, with more as they ship.`
- Produces: the FAQ entry whose question is `What happens to my long-retention logs if I stop?` is removed; the other five stay verbatim.

- [ ] **Step 1: Write the failing copy test**

In `web/src/lib/billing/copy.test.ts` add:

```ts
it('promises only what exists, and dates what is coming', () => {
  expect(premiumCopy.premiumFeatures).toEqual([
    'Run sims on our servers (5,000 combinations, any precision, no tab left open).',
    'Character history charts over time (ships November).',
    'Discord notifications (ships December).',
    'A supporter mark on your profile.',
  ]);
  expect(premiumCopy.guildFeatures).toEqual([
    'Everything in Premium, for everyone in the guild.',
    'Officer tools: the loot council helper and the raid readiness board, with more as they ship.',
  ]);
  expect(premiumCopy.faq.map((f) => f.question)).not.toContain('What happens to my long-retention logs if I stop?');
  for (const text of [...premiumCopy.premiumFeatures, ...premiumCopy.guildFeatures, ...premiumCopy.faq.map((f) => f.answer)]) {
    expect(text).not.toMatch(/retention|ninety|two-year|two years|Compare more than two/);
  }
});
```

- [ ] **Step 2: Run it and watch it fail**

Run: `cd web && npx vitest run src/lib/billing/copy.test.ts`
Expected: FAIL on the `toEqual` of `premiumFeatures`.

- [ ] **Step 3: Change the copy**

Edit `premiumFeatures`, `guildFeatures` and `faq` in `web/src/lib/billing/copy.ts` to the Interfaces values. Leave every other export untouched.

- [ ] **Step 4: Run the copy test**

Run: `cd web && npx vitest run src/lib/billing/copy.test.ts`
Expected: PASS.

- [ ] **Step 5: Write the failing legal-pages test**

In `web/src/content/pages/legal-markers.test.ts`, inside the `for (const file of LEGAL_FILES)` loop, add:

```ts
it(`${file} promises no retention window`, () => {
  expect(text).not.toMatch(/ninety|two-year|two years|retention falls back|long-retention/i);
});
```

- [ ] **Step 6: Run it and watch it fail**

Run: `cd web && npx vitest run src/content/pages/legal-markers.test.ts`
Expected: FAIL for `terms.md` and `privacy.md`.

- [ ] **Step 7: Rewrite the two passages**

`terms.md`, the Cancelling paragraph: end it at "— nothing is deleted." and remove the two sentences about long-retention logs.

`privacy.md`, replace the `## Data retention` body with exactly:

```
Uploaded combat logs, the reports parsed from them, saved sims and the ratings derived from
them are kept until you delete them or your account; the rankings and a character's history
are built from that record and would be wrong without it. Session rows expire with the
session. If a storage limit ever becomes necessary it will be announced on the changelog
with dates before anything is removed.
```

- [ ] **Step 8: Run the legal test, the premium e2e and the legal e2e**

Update `web/tests/e2e/premium.spec.ts` so its feature assertions name the new four lines and assert the retention and compare lines are absent (`getByText('Two years of log retention', { exact: false })` has count 0). `web/tests/e2e/legal-pages.spec.ts` needs no change unless it names a removed sentence; check.

Run: `cd web && npx vitest run src/content/pages/legal-markers.test.ts && FOREVER_DATA=fixture npx playwright test tests/e2e/premium.spec.ts tests/e2e/legal-pages.spec.ts --project=desktop`
Expected: all PASS.

- [ ] **Step 9: Changelog entry**

Create the changelog entry in the past tense, dated 2026-10-24: the Premium page now lists only what is live, with the month beside what is coming; logs and saved sims are kept for everyone with no retention tier; the "compare more than two builds" line is retired.

- [ ] **Step 10: Commit**

```bash
git add web/src/lib/billing/copy.ts web/src/lib/billing/copy.test.ts web/src/content/pages/terms.md web/src/content/pages/privacy.md web/src/content/pages/legal-markers.test.ts web/tests/e2e/premium.spec.ts web/tests/e2e/legal-pages.spec.ts web/src/content/changelog/2026-10-24-premium-page-truth.md
git commit -m "feat(premium): the page and the legal pages promise only what exists; logs are kept for everyone"
```

---

### Task 2: `supporter` on the public character and roster payloads

**Files:**
- Modify: `api/internal/rankings/handler.go` (`Service` struct, `character` handler), `api/internal/rankings/query.go` (the character page DTO that carries `score_history`)
- Modify: `api/internal/guilds/home.go` (`RosterRow`, the roster query)
- Modify: `api/cmd/api/main.go` (wire `Supporters` into both services)
- Test: `api/internal/rankings/handler_test.go` (or the package's existing character test file), `api/internal/guilds/home_test.go`

**Interfaces:**
- Consumes: `entitlements.Store.IsSupporter(ctx, userID int64) (bool, error)` (exists).
- Consumes: the character's owning account. Migration `0024_bnet_characters` is where a character gained an account; use the same column `GET /v1/me` uses to list an account's characters (read `api/internal/auth/handler.go`'s characters query to confirm the column name before writing the join). A character with no owner is never a supporter.
- Produces, in both packages:

```go
// Supporters answers whether an account shows the supporter mark; nil means
// nobody does (the fixture build, and tests that do not care).
type Supporters interface {
	IsSupporter(ctx context.Context, userID int64) (bool, error)
}
```

  as a field `Supporters Supporters` on `rankings.Service` and on the guilds home service.
- Produces: the character page DTO gains `Supporter bool \`json:"supporter"\``; `guilds.RosterRow` gains `Supporter bool \`json:"supporter"\``. Both are `false` when `Supporters` is nil or the character has no owner; an `IsSupporter` error is logged and read as `false`, never a 500 (a cosmetic flag must not take the page down).

- [ ] **Step 1: Failing test, character page**

In the rankings handler test file add a test that seeds a character owned by a user with an `entitlements` row `(user_id, plan 'premium', source 'stripe', status 'active', cancel_at_period_end true)`, wires `Supporters: &entitlements.Store{Pool: pool}`, calls `GET /v1/characters/{region}/{ruleset}/{name}` and asserts `supporter` is `true`; then updates the row to `status = 'canceled'` and asserts `false`; then calls for a character with no owner and asserts `false` with status 200.

- [ ] **Step 2: Run it and watch it fail**

Run: `cd api && TEST_DATABASE_URL='postgres://forever:forever@localhost:5434/forever_test?sslmode=disable' go test -p 1 -count=1 ./internal/rankings/ -run TestCharacterPageCarriesTheSupporterMark`
Expected: FAIL, `supporter` missing from the body.

- [ ] **Step 3: Implement in rankings**

Add the `Supporters` interface and field; in `character`, after the DTO is built, resolve the owner id (the query already joins what it needs, or add the column to `Store.Character`'s select) and set `Supporter` through a small helper:

```go
func (s *Service) supporterOf(ctx context.Context, ownerID *int64) bool
```

  that returns false on nil service, nil owner, or error (logged at warn).

- [ ] **Step 4: Run the test**

Run: the same command. Expected: PASS.

- [ ] **Step 5: Failing test, roster**

In `api/internal/guilds/home_test.go` add a test with a claimed guild, two verified members (`leader` and `member`), a guild-plan `entitlements` row for the guild (`guild_id`, plan `guild`, status `active`), and assert both `RosterRow.Supporter` are `true`; with the row `canceled`, both `false`.

- [ ] **Step 6: Run it and watch it fail**

Run: `cd api && TEST_DATABASE_URL=... go test -p 1 -count=1 ./internal/guilds/ -run TestRosterRowsCarryTheSupporterMark`
Expected: FAIL.

- [ ] **Step 7: Implement in guilds**

Same interface and helper shape; set `Supporter` per row from the row's account id (the roster query already distinguishes account rows from name-only rows; name-only rows are `false`). Batch: call `IsSupporter` once per distinct account, not once per character.

- [ ] **Step 8: Wire main.go**

Both services get `Supporters: entitlementsStore` where `main.go` already constructs the entitlements store for billing. Run `go vet ./...` and the two packages' tests.

- [ ] **Step 9: Commit**

```bash
git add api/internal/rankings api/internal/guilds api/cmd/api/main.go
git commit -m "feat(api): the character page and guild rosters say who is a supporter"
```

---

### Task 3: The Supporter chip and the account page's Premium block

**Files:**
- Create: `web/src/components/SupporterChip.svelte`, `web/src/components/SupporterChip.test.ts`
- Modify: `web/src/lib/account/api.ts` (the character page type and the roster row type gain `supporter?: boolean`)
- Modify: `web/src/components/character/CharacterHeader.svelte` (chip beside the name when the loaded character's `supporter` is true), `web/src/components/guild/GuildRosterTable.svelte` (chip after the name in account rows)
- Modify: `web/src/components/Account.svelte` (the billing row at ~`:886` becomes the Premium block), `web/src/lib/billing/copy.ts` (`premiumBlockCopy`)
- Modify: `web/src/test-support/billing-fixtures.ts` (`supporter_mark` toggle helper), `web/src/fixtures/*` character fixture gains `supporter`
- Test: `web/src/components/character/CharacterHeader.test.ts`, `web/src/components/guild/GuildRosterTable.test.ts`, `web/src/components/Account.test.ts` (or the account page's existing test file), `web/tests/e2e/character-identity.spec.ts`, `web/tests/e2e/account-billing.spec.ts`

**Interfaces:**
- Produces: `SupporterChip.svelte` with no props, rendering `<span data-testid="supporter-chip" class="label ...gold chip classes from the design system">Supporter</span>`; the one word comes from `premiumBlockCopy.supporterChip`.
- Produces: `premiumBlockCopy` in `copy.ts`:

```ts
export const premiumBlockCopy = {
  heading: 'Premium',
  rows: {
    serverSims: { label: 'Server sims', on: 'On', off: 'Browser lane only' },
    history: { label: 'History charts', on: 'Ships November', off: 'Ships November' },
    notifications: { label: 'Discord notifications', on: 'Ships December', off: 'Ships December' },
    supporter: { label: 'Supporter mark', on: 'Shown on your characters', off: 'Not shown' },
  },
  supporterChip: 'Supporter',
} as const;
```

- Produces: the Premium block renders one row per key, label then state, state chosen by `entitlements.server_sims` for `serverSims`, `entitlements.supporter_mark` for `supporter`, and the fixed "Ships" text for the two coming rows; the existing plan line, past-due banner and `Manage billing` / `See plans` controls stay under it unchanged.

- [ ] **Step 1: Failing chip and header tests**

`SupporterChip.test.ts`: renders the word `Supporter` with the testid. `CharacterHeader.test.ts`: with a character fixture whose `supporter` is `true` the header contains the chip; with `false` or absent it does not.

- [ ] **Step 2: Run them and watch them fail**

Run: `cd web && npx vitest run src/components/SupporterChip.test.ts src/components/character/CharacterHeader.test.ts`
Expected: FAIL (module missing; chip absent).

- [ ] **Step 3: Build the chip and place it**

Create the chip. In `CharacterHeader.svelte`, render it immediately after the name element, inline, with the same gap the descriptor uses; in `GuildRosterTable.svelte`, after the name in account rows only.

- [ ] **Step 4: Run the component tests**

Run: the same command plus `src/components/guild/GuildRosterTable.test.ts` (add a row with `supporter: true` and assert the chip once).
Expected: PASS.

- [ ] **Step 5: Failing account block test**

In the account page's test file: with `mockEntitlements({ server_sims: true, supporter_mark: true })` the block shows `Server sims · On` and `Supporter mark · Shown on your characters`; with defaults it shows `Browser lane only` and `Not shown`; with no `entitlements` at all it still renders the block in the free state.

- [ ] **Step 6: Run it and watch it fail**

Run: `cd web && npx vitest run src/components/Account.test.ts` (use the file that already tests the billing row).
Expected: FAIL.

- [ ] **Step 7: Build the block**

Replace the single billing row's wrapper with the block: heading, the four rows, then the existing plan line and controls. Keep every existing testid.

- [ ] **Step 8: Run the account tests, then the e2e**

Extend `account-billing.spec.ts` with one test that routes `/v1/me` with `supporter_mark: true` and sees the chip text in the block; extend `character-identity.spec.ts` with a fixture character carrying `supporter: true` and assert the chip beside the name, and once more with `false` asserting its absence.

Run: `cd web && npx vitest run src/components/Account.test.ts && FOREVER_DATA=fixture npx playwright test tests/e2e/account-billing.spec.ts tests/e2e/character-identity.spec.ts --project=desktop`
Expected: PASS.

- [ ] **Step 9: Commit**

```bash
git add web/src/components/SupporterChip.svelte web/src/components/SupporterChip.test.ts web/src/components/character web/src/components/guild/GuildRosterTable.svelte web/src/components/guild/GuildRosterTable.test.ts web/src/components/Account.svelte web/src/lib/billing/copy.ts web/src/lib/account/api.ts web/src/test-support web/src/fixtures web/tests/e2e/account-billing.spec.ts web/tests/e2e/character-identity.spec.ts
git commit -m "feat(web): the Supporter chip on characters and rosters, and the account page's Premium block"
```

---

### Task 4: Lane close

- [ ] **Step 1: Full suites**

Run: `cd web && npm run check && npm run lint && npx vitest run && FOREVER_DATA=fixture npx playwright test` and `cd api && go vet ./... && TEST_DATABASE_URL=... go test -p 1 ./...`
Expected: all green. The lhci shard runs in CI on the lane's dispatch.

- [ ] **Step 2: Screenshots for review**

Capture the character page with and without the chip, the guild roster, and the account page's Premium block at 1440 and 2000, into `design/reviews/2026-10-24-premium-lane-1/`, and run the ux-designer and wow-player reviewers on them before the merge.

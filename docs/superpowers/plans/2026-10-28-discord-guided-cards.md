# Discord guided cards and raids that post themselves: Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Every bot command that used to answer a missing argument with a sentence now answers with an only-you card that asks for it (`/bis` class, spec and range; `/char`, `/signup`, `/attendance`, `/readiness`, `/report` per the one rule). Raids also post themselves: `Post to Discord` and `Remind the day before` are on by default and remembered per guild. A raid posts when saved if it starts within 7 days, otherwise 7 days before it starts. A reminder goes out 24 h before the start. It edits the post when the post is still the newest message, replies under the post when it is not, and is deleted again on a move, cancel or lock.

**Architecture:**
- **Part G: guided cards.** This lives entirely in `api/internal/discordbot`. A card's state rides in its custom id (`g{cmd}:1:{part}:{state…}`), so the server keeps nothing between presses. Every card press is a component interaction (type 3) answered with type 7, or type 6 then an edit. Two Discord paths are new: modals (response 9) and modal submits (interaction 5). A card that commits builds its answer with the same function the typed command uses.
- **Part P: posting schedule.** Each raid gets one row in a new table, `guild_raid_discord_schedule`, owned by `guilds`. The row says when the raid posts and whether it reminds. The destination (bot or webhook) is picked only when the post or the reminder actually goes out.
- **The clock** is a new Cloud Run job, `raid-posts`. Cloud Scheduler runs it every 10 minutes, the same way `billing-daily` is run. On each run it sends the due posts and reminders through `RaidEvents` (the bot's router, or the webhook), then runs bot drain passes.
- **Who does the Discord work for a bot reminder:** the job only flags the reminder as due. The bot's existing per-raid drainer then makes the Discord calls: it reads the newest channel message, edits the post into its `Tomorrow` state, and replies under it when needed.
- **Web:** the raid form's checkbox becomes a `role="switch"` row, with a second row under it for the reminder.

**Tech Stack:** Go (stdlib `net/http`, pgx v5), Postgres migrations (golang-migrate files in `api/internal/db/migrations`), Cloud Run jobs plus Cloud Scheduler, Astro with Svelte 5 runes, Vitest, Playwright.

**Spec:** `design/specs/2026-10-28-discord-guided-cards.md` is the authority: copy is verbatim from §G.7 and §P.7, and acceptance is its "Acceptance" section. It amends `design/specs/2026-10-27-discord-bot.md` (the bot spec) and `design/specs/2026-10-26-guild-raids.md` (G1). The existing bot design is `docs/superpowers/specs/2026-10-27-discord-bot-design.md`. The mock generator `design/mocks/2026-10-28-discord-guided-cards/gen_guided_cards.mjs` fixes the exact embed and row layout of every `/bis` frame and every post frame. Executors read the spec and this plan; where they differ, "Rulings" below decides.

## Global Constraints

- **Secrets.** The bot token, webhook URLs (sealed `token_box`), interaction tokens and OAuth codes are never printed, logged, written to a file, put in a custom id or returned in a response. Errors that wrap a Discord answer go through the existing `describe`/`sanitize`. No step in this plan echoes a secret, including in a shell.
- **Migrations.** Only the deployed API applies migrations. No task migrates a shared database. The `raid-posts` job migrates at start like every job, because it runs the deployed image.
- **API tests.** Run with `-p 1` against a private database per agent: `TEST_DATABASE_URL=postgres://forever:forever@localhost:5434/<private db>?sslmode=disable`, created with `psql postgres://forever:forever@localhost:5434/forever_test -c 'create database forever_gc_<agent>'`. A test that skips for a missing `TEST_DATABASE_URL` is not a pass. Below, `$TDB` stands for that variable assignment.
- **Scoped testing.** Each task runs the packages and files it touched (`go test ./internal/discordbot/...` plus the other touched packages; `cd web && npx vitest run <touched files>`). The full suite is CI's job. Never chain a piped test command into a commit; gate on the test command's own exit code.
- **Copy.** Strings come verbatim from the amendment.
  - Discord card strings live in `guideCopy` (`api/internal/discordbot/copy_guide.go`). Their test reads each one back out of §G.7 with `spectable.Cell` (path `../../../design/specs/2026-10-28-discord-guided-cards.md`, heading `## G.7 Copy, verbatim`, label the backticked key such as `` `bis.lead.main` ``).
  - Post and reminder strings join notify's `raidCopy`. They are read from `## P.7 Copy, verbatim`.
  - Site strings join `raidCopy.form` in `web/src/lib/guild/raid-copy.ts`, read with `web/src/test-support/spec-table.ts`.
  - No string uses `soon`, `!` or `unlock`. Member text (`{Class}`… `{typed}`, `{Title}`) goes through `memberText`/notify's escaper.
- **Discord limits**, asserted by a shared helper `assertCardLimits`:
  - Custom ids are at most 100 characters.
  - At most 5 action rows.
  - Every string select has 1 to 25 options.
  - Labels and option descriptions are at most 100 characters; placeholders at most 150.
  - `allowed_mentions` is `{"parse": []}` on every message.
  - Button styles are only Success (3, one per card, its forward action), Secondary (2) and Link (5). No emoji.
- **Privacy.** Every card is ephemeral (flag 64). Nothing new posts to a channel except the raid post and reminder (part P) and `Share in channel` (unchanged). No custom id or select value identifies another account, except officer `/attendance` values (character names), which are re-checked against rank on every pick. `/signup` card ids carry only the presser's own character key, and every press re-checks it.
- **One answer function per command.** A card's commit answer is produced by the same function as the typed command (`bisAnswer`, `answerChar`, `reportMessage`, the press flow), so the two cannot drift. The test compares them byte for byte.
- **File size.** Go files stay under 800 lines; new card code goes in new files (`card.go`, `bis_card.go`, `char_card.go`, `signup_card.go`, `officer_cards.go`). Svelte and TS files stay under 400 lines. `RaidForm.svelte` (379 lines) moves its Discord block into `RaidFormDiscord.svelte`.
- **Commits.** Conventional commits. Every commit ends with `Co-Authored-By: Claude Opus 5.5 <noreply@anthropic.com>`. Run `git reset -q` before `git add`, because the shared index can be stale, then `git add` only the task's paths. Never `git stash`.
- **Migration number:** `0057_discord_guided_posting` (the last one is `0056_coach_account_spend`).

## Rulings (spec silences and conflicts decided here)

1. **The schedule is destination-neutral and lives in `guilds`.** P.8 has the bot store gain `post_at`, `remind`, `reminder_message_id` and `reminded_at`. But a raid saved 8 days out must pick bot or webhook when it posts, not when it is saved. So `post_at`, `remind`, `posted_at`, `reminded_at` and the sent reminder's reference live in `guild_raid_discord_schedule`. The bot's `guild_raid_discord_posts` gains only what its drainer needs: `remind_due`, `reminder_message_id`, `reminded_at`. The raid view's `discord_post` (P.8) reads the schedule row.
2. **The clock is a job.** `raid-posts` runs every 10 minutes (`*/10 * * * *`), following the `billing-daily` pattern (`api/cmd/api/billing_daily.go`, the `main.go` switch, the `api.yml` image-update list, and the README job and scheduler commands).
   - It works whether or not the bot is configured, and whether or not the service is warm.
   - A deferred post or a reminder can land up to 10 minutes after its time.
   - After the schedule pass, the job runs up to 5 `Drainer.DrainOnce` passes, so a bot reminder or post goes out even if the service is scaled to zero. Leases keep the job and a warm service drainer off the same row.
3. **The 12-hour rule** (P.4) is a due condition, not a send-time check: a reminder is due only when `posted_at <= starts_at - 36h`. A skipped reminder is never sent, and the form keeps `Remind the day before` live until the start.
4. **Moved beats Tomorrow.** If a raid moves after its reminder, the drainer deletes the reminder message and clears `reminded_at`. The post shows `Moved:` (`PostState` already ranks moved above open), and the schedule re-arms the reminder for the new start.
5. **Existing guilds flip to on.** Migration 0057 sets `post_default = true` on every install and webhook row, and makes `true` the column default. The opt-in `false` they hold is the default the owner rejected, not an officer's choice. `remind_default` is added as `true`.
6. **Legacy posts.** Every existing `guild_raid_discord_posts` row is backfilled into the schedule as posted (`post_at = posted_at = created_at`, `remind = false`), so editing an old raid never posts it twice. Webhook-era posts cannot be known and are not backfilled.
7. **Card failure (G.6) echoes the pressed message.** The component interaction's `message` carries the card's own `embeds` and `components`. On a card route, a handler error answers type 7 with those, unchanged, and `card.failure`:
   - as the content line;
   - for `gbis`, as the first embed's description instead (the lead).
   Every pick is kept because the `default: true` options are echoed as sent.
8. **Part names** that §G.4.1 leaves open:
   - `ask`: a button that opens a modal.
   - `name`: a modal submit.
   - `pick`: the `/char` own-characters select and the `/attendance` officer select.
   - `raid`: the `/readiness`, `/report` and `/signup` raid selects.
   - `char`: the `/signup` character select.
   - The four state words: `/signup` status buttons.
   - The text input inside both modals has custom id `name`.
9. **`/char` answers reached through the card** (a bare `/char`, a select pick, a submitted name) carry `Someone else` and the own-characters select. A typed `/char name:` answers as today. `Share in channel` re-renders a guided answer with its controls, because the share row's options carry the hidden key `guided: "1"`, which `answerChar` ignores. A `/bis` answer always carries `Change picks` and a `/report` answer always carries its select (when there are 2 or more nights), including after Share.
10. **`/attendance`'s officer list** contains only members with at least one night, so every option has a chip.
11. **`{channel}` in the form helpers** is `#` plus `channel_name` cut to 60 characters followed by `…`; the 60 does not count the `#`.
12. **`card.footer.window`** is not built unless capture D13 (Task 15) shows Discord refusing a press 20 minutes after the card was drawn. In that case Task 15 adds it to every card's footer, and nothing else changes.

## Review Focus

1. **Forged or garbled card ids**, for example a class that does not exist, a spec of another class, a range outside `phaseChoices`, extra or missing state fields, a non-numeric raid, or another account's character key in `gsgn`. Expected: unknown fields read as empty and the card re-renders; `gsgn` with a key that is not the presser's answers `That did not save. Try again.`; nothing panics. Test: Task 3 `TestBisPicksIgnoreForgedFields`, Task 6 `TestSignupCardRefusesAForgedCharacter`.
2. **A raid moved, cancelled or locked between scheduling and its post or reminder.** Expected: a cancelled or locked unposted raid never posts; a moved unposted raid's `post_at` follows the new start; a sent reminder is deleted (bot and webhook) and, on a move, re-armed. Test: Task 12 `TestScheduleFollowsMovesCancelsAndLocks`, Task 11 `TestDrainerWithdrawsTheReminder`.
3. **The job and a warm service drainer at once,** or the job overlapping itself (a slow run). Expected: one post and one reminder per raid, never two. Test: Task 12 `TestTwoScheduleRunsSendOnce` (two concurrent `RunRaidSchedule`), Task 11 `TestReminderDueIsDrainedOnce`.
4. **Share in channel from a card answer.** Expected: the public follow-up never carries `Change picks`, a select or `Someone else`; the presser's own re-rendered answer keeps them. Test: Task 4 `TestShareFromACardKeepsTheControls`, Task 5 `TestGuidedCharShareKeepsTheControls`.
5. **Hostile and long member text in cards and reminders**: an 80-rune raid title, a 12-character name of markdown and `@everyone`, a 60-plus character channel name. Expected: rendered literally, pings no one, and every limit holds. Test: Task 6 `TestSignupCardSurvivesHostileTitles`, Task 10 `TestReminderLineSurvivesHostileTitles`, Task 14 `channel name is cut at 60`.

## Parallel waves

- **Wave A:** Task 1, then Task 2.
- **Wave B:** Task 9 runs alongside Wave A.
- **Wave C:** Tasks 3 to 8 (Task 3 before Task 4; Tasks 5, 6, 7 and 8 run in parallel after Task 2).
- **Wave D:** Task 10, then Task 11, then Task 12, then Task 13, alongside Wave C.
- **Wave E:** Task 14 (after Task 12) and Task 15 (last).
- `routes.go`, `commands.go` and `copy_commands.go` are touched by several Wave C tasks. Each adds its own lines; resolve conflicts by keeping both.

---

### Task 1: Card grammar and the guided copy table

**Files:**
- Create: `api/internal/discordbot/card.go`, `api/internal/discordbot/copy_guide.go`, `api/internal/discordbot/card_test.go`, `api/internal/discordbot/copy_guide_test.go`
- Modify: `api/internal/discordbot/message.go` (`SelectOption.Default`), `api/internal/discordbot/spectable_test.go` (a `guideString` helper beside `specString`)

**Interfaces:**
- Produces:
  - `const cardVersion = "1"`
  - Prefixes: `const (prefixBis = "gbis"; prefixChar = "gchar"; prefixSignup = "gsgn"; prefixAttendance = "gatt"; prefixReadiness = "grdy"; prefixReport = "grep")`
  - `func cardCustomID(prefix, part string, state ...string) string` builds `{prefix}:1:{part}:{state joined by ":"}`, keeping empty fields. With no state it ends in `:`, as `gchar:1:name:` does.
  - `type cardPress struct{ Prefix, Version, Part string; State []string }` and `func parseCardID(id string) (cardPress, bool)`: ok is false with fewer than 3 fields.
  - `func staleCard(command string) Reply`: Update true, content `card.stale` filled with the command name, no embeds, no components.
  - `func lockedSelect(customID, placeholder string) Component`: disabled, one option `{Label: placeholder, Value: "none"}`.
  - `type cardPart func(ctx context.Context, in Interaction, p cardPress) (Reply, error)`
  - `func (s *Service) cardHandler(command string, parts map[string]cardPart) Handler`: an unknown version answers `staleCard(command)`; an unknown part answers `failureReply()` with Update.
  - `var guideCopy = struct{...}` with the §G.7 keys as fields (for example `BisAuthor`, `BisTitleNone`, `BisLeadMain`, `BisNudgeSpec`, `BisRangeDesc [6]string`, `CardStale`, `CardFailure`). There is no `card.footer.window` field (Ruling 12).
  - Test helper `assertCardLimits(t *testing.T, m Message)`.

- [ ] **Step 1: Write the failing tests**

```go
// card_test.go
func TestCardIDRoundTrips(t *testing.T) {
	id := cardCustomID(prefixBis, "chg", "warrior", "feral-bear", "p1", "main_hand")
	if id != "gbis:1:chg:warrior:feral-bear:p1:main_hand" || len(id) > 100 { t.Fatal(id) }
	p, ok := parseCardID(id)
	if !ok || p.Prefix != "gbis" || p.Version != "1" || p.Part != "chg" || strings.Join(p.State, "|") != "warrior|feral-bear|p1|main_hand" { t.Fatalf("%+v", p) }
	if got := cardCustomID(prefixBis, "c", "", "", "p1", ""); got != "gbis:1:c:::p1:" { t.Fatal(got) }
	if customPrefix(id) != prefixBis { t.Fatal("router key") }
}
func TestUnknownCardVersionIsStale(t *testing.T) { /* cardHandler("bis", …) on "gbis:9:go:::p1:" → Update, Content == "This card is out of date. Run /bis again.", no rows */ }
func TestLockedSelectHasOneNoneOption(t *testing.T) { /* Disabled, len(Options)==1, Value "none", Label == placeholder */ }
func TestSelectDefaultMarshals(t *testing.T) { /* {"label":…,"value":…,"default":true}; Default false omits the key */ }
// copy_guide_test.go
func TestGuideCopyMatchesTheSpec(t *testing.T) { /* each field == guideString(t, "`key`", i), e.g. guideCopy.BisLeadMain == guideString(t,"`bis.lead.main`",0); BisRangeDesc[i] == guideString(t,"`bis.range.desc`",i) for i<6; no string contains "soon", "!" or "unlock" */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && go test ./internal/discordbot/ -run 'CardID|UnknownCardVersion|LockedSelect|SelectDefault|GuideCopy' -v`
Expected: FAIL to compile (undefined `cardCustomID`).

- [ ] **Step 3: Implement the Interfaces in `card.go` and `copy_guide.go`, and `SelectOption.Default bool `json:"default,omitempty"`` in `message.go`**

`guideCopy` stores bold markers as Discord markdown, exactly as the spec's "(bold)" notes say:
- `BisNudgeSpec = "**Choose a spec first.** Then Show BiS."`
- `BisLeadMain = "Picked from your main, **{Name}**. Show BiS, or change a pick."`

The test compares the spec's segments after stripping `**`.

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/card.go api/internal/discordbot/copy_guide.go api/internal/discordbot/card_test.go api/internal/discordbot/copy_guide_test.go api/internal/discordbot/message.go api/internal/discordbot/spectable_test.go
git commit -m "feat(discordbot): card custom ids, stale cards and the guided copy table"
```

---

### Task 2: Modals, modal submits, and the card failure echo

**Files:**
- Create: `api/internal/discordbot/modal.go`, `api/internal/discordbot/modal_test.go`
- Modify: `api/internal/discordbot/interaction.go` (`typeModalSubmit = 5`, the source message, modal values), `api/internal/discordbot/respond.go` (`respModal = 9`, dispatch of type 5, `replyOf` on card routes), `api/internal/discordbot/routes.go` (`Reply.Modal`, `cardFailureRoutes`), `api/internal/discordbot/message.go` (an echo edit body)

**Interfaces:**
- Consumes: `parseCardID`, `guideCopy.CardFailure` (Task 1).
- Produces:
  - `type Modal struct{ CustomID, Title, InputLabel, Placeholder string; MinLength, MaxLength int }`. Its `MarshalJSON` writes `{custom_id, title, components:[{type:1, components:[{type:4, custom_id:"name", style:1, label, placeholder, min_length, max_length, required:true}]}]}`.
  - `Reply.Modal *Modal`. `writeReply` sends `{type:9, data: modal}` when it is set; a modal is never deferred.
  - `func (i Interaction) ModalValue(inputID string) string` reads `data.components[].components[]` for the input's `value`.
  - `Interaction.Source sourceMessage`, which keeps the pressed message's raw `embeds` and `components` (`json.RawMessage`).
  - `func cardFailureReply(in Interaction) Reply`: an Update whose edit body is the source's embeds and components with content `card.failure`. For prefix `gbis` the content stays empty and the first embed's `description` is replaced by `card.failure`.
  - `func cardFailureRoutes() map[string]bool`: the six card prefixes. `replyOf` uses `cardFailureReply` for them.
  - `serveInteraction` routes type 5 through `dispatch` exactly like type 3, keyed by the modal's custom id. Opener handlers (`ask` parts) must do no I/O, so they answer inside the window.

- [ ] **Step 1: Write the failing tests**

```go
func TestModalOpensAtOnce(t *testing.T) { /* a component handler returning Reply{Modal:&Modal{…"gchar:1:name:"…}} → body {"type":9,"data":{"custom_id":"gchar:1:name:","title":…,"components":[{"type":1,"components":[{"type":4,"custom_id":"name","style":1,"required":true,"min_length":2,"max_length":12,…}]}]}} */ }
func TestModalSubmitRoutesByCustomID(t *testing.T) { /* signed type-5 body {"data":{"custom_id":"gchar:1:name:","components":[{"type":1,"components":[{"type":4,"custom_id":"name","value":"Zulmara"}]}]}} reaches the gchar handler; in.ModalValue("name") == "Zulmara" */ }
func TestCardFailureKeepsTheCard(t *testing.T) {
	// a gbis handler returning an error, the interaction's message carrying one embed (description "Ready. Show BiS, or change a pick.") and four rows
	// → type 7, embeds[0].description == guideCopy.CardFailure, components byte-equal to the source's, content ""
	// the same on a gsgn route → content == guideCopy.CardFailure, embeds and components unchanged
}
func TestNonCardFailureIsUnchanged(t *testing.T) { /* a pick: route error still answers failureReply() with Update (today's behaviour) */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/discordbot/ -run 'Modal|CardFailure|NonCardFailure' -v`
Expected: FAIL (undefined `Modal`).

- [ ] **Step 3: Implement the Interfaces**

The deferred path (`deliver`) sends the echo through `EditOriginal`. The echo is carried as `Message.echo *sourceMessage`, and `Message.edit()` writes the raw embeds and components when `echo` is set.

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Respond|Interactions|Share'` (regressions). Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/modal.go api/internal/discordbot/modal_test.go api/internal/discordbot/interaction.go api/internal/discordbot/respond.go api/internal/discordbot/routes.go api/internal/discordbot/message.go
git commit -m "feat(discordbot): modals, modal submits and the card failure that keeps the picks"
```

---

### Task 3: The `/bis` card: render and the bare or partial command

**Files:**
- Create: `api/internal/discordbot/bis_card.go`, `api/internal/discordbot/bis_card_test.go`, `api/internal/discordbot/card_harness_test.go`
- Modify: `api/internal/discordbot/cmd_bis.go` (entry), `api/internal/discordbot/commands.go` and `copy_commands.go` (§G.8 `/bis` descriptions), `api/internal/discordbot/copy_lookup.go` (retire `BisNeedsSpec`), `api/internal/discordbot/cmd_char_test.go` (`TestLookupCopyMatchesTheSpec` drops `BisNeedsSpec`)

**Interfaces:**
- Consumes: Task 1 (`cardCustomID`, `lockedSelect`, `guideCopy`, `assertCardLimits`).
- Produces:
  - `type bisPicks struct{ Class, Spec, Range, Slot string }`. `Spec` is the spec slug within the class, for example `beast-mastery`.
  - `func (p bisPicks) state() []string` returns `[class, spec, range, slot]`.
  - `func bisPicksOf(state []string) bisPicks`: forged values read as empty. A class outside `classChoices` reads `""`. A spec not of the class reads `""`. A range outside `bandLabels` reads `p1`. A slot outside `slotChoices` reads `""`. Missing fields read as empty.
  - `func (p bisPicks) specKey() (string, bool)` returns, for example, `hunter-beast-mastery`.
  - `type bisMain struct{ Name string; Picks bisPicks }`.
  - `func (s *Service) mainPicks(ctx context.Context, guildID int64, viewer *int64) (*bisMain, error)` wraps `mainOf`. It adds the main's name to `mainFacts` and returns nil without a main.
  - `type bisCard struct{ Picks bisPicks; Lead string; Main *bisMain; Linked bool }`.
  - `func (s *Service) bisCardMessage(c bisCard) Message`: one embed and four rows, ephemeral.
  - `func bisLeadFor(p bisPicks, main *bisMain, linked bool) string`. It returns:
    - `bis.lead.main` when the picks equal the main's;
    - `bis.lead.ready` with a spec;
    - `bis.lead.class` with a class;
    - `bis.lead.slot` with a slot;
    - `bis.lead.unlinked` when not linked;
    - `bis.lead.none` otherwise.
  - The card's rows (spec §G.4.2):
    1. Class select `gbis:1:c:{state}`. Nine options in `classChoices` order; each description is the class's spec names in `specs.All` order, joined by `, `.
    2. Spec select `gbis:1:s:{state}`. It is `lockedSelect` with `Choose a class first` when there is no class. Otherwise it lists the class's specs, each described by `roleWord(sp.Role)`.
    3. Range select `gbis:1:r:{state}`. The six `phaseChoices`, described by `BisRangeDesc`.
    4. Buttons:
       - `Show BiS` (Success, `gbis:1:go:{state}`);
       - `Use {Name}` (Secondary, `gbis:1:me:{state}`), only when `Main != nil` and a pick differs from `Main.Picks` on class, spec or range;
       - `s.linkButton()`, only when `!Linked`.
    - The current pick of each select carries `Default: true`.
  - The embed: author `Best in slot`.
    - Title: `bis.title.none`, `bis.title.class` or `bis.title.ready` (the last uses `bandLabels`).
    - Colour: `goldColor` until there is a class, then `classColors[class]`.
    - Thumbnail: `{PublicBaseURL}/icons/hd/crests/{class}.webp` once there is a class.
    - Footer `Forever Sixty`, with the timestamp `s.now()`.
  - `cmdBis` with no `spec` option answers the card. Its picks follow these rules:
    - **Class:** the typed `class` when valid. With no `class` typed, the main's class for a linked viewer with a main.
    - **Spec:** the main's spec when the class is the main's.
    - **Range:** the typed `phase`, else the main's band when the class is the main's (`phaseOfBand`), else `p1`.
    - **Slot:** the typed `slot`.
  - A typed `spec` that does not resolve opens the card with the typed valid class (or none), an empty spec, and the lead `bis.lead.unknownSpec`. That lead is the class form when there is a class and the plain form when there is not.
  - A typed resolvable `spec` answers at once, exactly as today.
  - Test harness in `card_harness_test.go`:
    - `func (w *lookupWorld) card(t *testing.T, discordUser, customID string, values ...string) Reply` runs the routed handler for a type-3 interaction.
    - `func (w *lookupWorld) modal(t *testing.T, discordUser, customID, typed string) Reply` does the same for a type-5 interaction.

- [ ] **Step 1: Write the failing tests**

```go
func TestBareBisUnlinked(t *testing.T) {
	m := w.command(t, "d-stranger", "bis", nil).Message
	assertCardLimits(t, m)
	// m.Ephemeral; Embeds[0].Title == "Choose a class"; Description == guideCopy.BisLeadUnlinked; Color == goldColor; Thumbnail == ""
	// 4 rows; row 2 Disabled with one option Value "none" Label "Choose a class first"
	// row 3 option "p1" Default; Label "Phase 1 raid", Description "Level 60, raid buffs and consumables"
	// row 1 Druid option Description == "Balance, Feral, Feral Bear, Restoration"
	// row 4 == [Show BiS Success "gbis:1:go:::p1:", Link Discord Link]
}
func TestBareBisLinkedWithAMain(t *testing.T) { /* main Zulmara (Beast Mastery Hunter, level 60): class/spec/range Default on hunter/beast-mastery/p1; Description == "Picked from your main, **Zulmara**. Show BiS, or change a pick."; Title "Beast Mastery Hunter · level 60 Phase 1 raid"; Color == classColors["hunter"]; Thumbnail ends "/icons/hd/crests/hunter.webp"; no Use button; no Link button */ }
func TestBareBisLinkedWithoutACharacter(t *testing.T) { /* Description == guideCopy.BisLeadNone; no Link Discord */ }
func TestBisPartials(t *testing.T) {
	// class only (warrior) → class Default warrior, spec empty, lead "Now choose a Warrior spec."
	// phase only "40" (unlinked) → range Default "40"
	// slot only "head" → lead "For Head only. Choose a class, a spec and a range, then Show BiS."; every id's state ends ":head"
	// spec "nonsense" with class warrior → lead "No Warrior spec by that name. Choose one."; without class → "No spec by that name. Choose one."
	// spec "warrior-fury" phase p1 → today's answer (Content == the existing five-slot lead)
}
func TestBisPicksIgnoreForgedFields(t *testing.T) { /* bisPicksOf([]string{"paladin","fury","99","boots","x"}) == {Class:"paladin", Range:"p1"}; bisPicksOf(nil) == {Range:"p1"} */ }
func TestBisCommandDescriptions(t *testing.T) { /* bisOption "class".Description == "The class. Leave out to pick it."; "spec" == "The spec. Leave out to pick it." */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/discordbot/ -run 'BareBis|BisPartials|BisPicksIgnore|BisCommandDescriptions' -v`
Expected: FAIL (undefined `bisPicks`).

- [ ] **Step 3: Implement the Interfaces in `bis_card.go` and the entry in `cmdBis`**

`resolveBis` keeps resolving the typed spec. When there is no spec, the card path replaces the old `BisNeedsSpec` return. Delete `lookupCopy.BisNeedsSpec` and its spec-test reference.

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Bis|LookupCopy|Commands'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/bis_card.go api/internal/discordbot/bis_card_test.go api/internal/discordbot/card_harness_test.go api/internal/discordbot/cmd_bis.go api/internal/discordbot/commands.go api/internal/discordbot/copy_commands.go api/internal/discordbot/copy_lookup.go api/internal/discordbot/cmd_char_test.go
git commit -m "feat(discordbot): a bare /bis answers with the class, spec and range card"
```

---

### Task 4: `/bis` card presses, `Change picks`, and Share

**Files:**
- Modify: `api/internal/discordbot/bis_card.go` (parts), `api/internal/discordbot/cmd_bis.go` (`bisAnswer`), `api/internal/discordbot/routes.go` (`gbis` in `componentRoutes` and `updateComponents`), `api/internal/discordbot/share.go` (re-render decoration), `api/internal/discordbot/bis_card_test.go`

**Interfaces:**
- Consumes: Task 3 (`bisPicks`, `bisCardMessage`, `mainPicks`), Task 2 (`cardHandler`, failure echo).
- Produces:
  - `func (s *Service) bisAnswer(ctx context.Context, in Interaction, guildID int64, viewer *int64, opts map[string]string, update bool) (Reply, error)`. It is `answerBis`, then the Share button (`offerShare`), then `Change picks` (Secondary, `gbis:1:chg:{class}:{spec}:{range}:{slot}`) between Share and the link button, made ephemeral. `cmdBis` and the `go`, `me` and `chg` parts all use it. When nothing is found it answers `failureReply()`.
  - `func withChangePicks(msg Message, p bisPicks) Message`.
  - `func (s *Service) bisPress() Handler` is `cardHandler("bis", …)` with these parts:
    - `c`: a class pick. The spec clears, unless the new class is the main's, in which case the main's spec is restored. The range is kept. The lead is `bis.lead.class`, or `bis.lead.main` when the main's spec was restored.
    - `s`: a spec pick. The lead is `bis.lead.ready`.
    - `r`: a range pick. The lead is `bisLeadFor` (so a nudge gives way).
    - `go`: with a spec, `bisAnswer(…, update=true)`. Without a spec, the nudge in the lead (`bis.nudge.spec`, or `bis.nudge.both` when there is no class) with every pick kept.
    - `me`: the main is re-read with `mainPicks`, never from the id. The answer is for the main's spec and band; with no main it answers `failureReply()` with Update.
    - `chg`: the card reopens with the id's picks.
  - Every pick and press answers with Update.
  - `sharedReply` gains `func decorateOwn(command string, opts map[string]string, msg Message) Message`. For `bis` it adds `withChangePicks`, using picks rebuilt from the stored spec key, phase and slot. The public follow-up (`visitorAnswer`) is never decorated.

- [ ] **Step 1: Write the failing tests**

```go
func TestBisCardPicks(t *testing.T) {
	// c on "gbis:1:c:::p1:" value "hunter" (unlinked) → Update, Title "Hunter · choose a spec", spec select enabled with 3 options described "DPS", range Default "p1", lead "Now choose a Hunter spec."
	// s value "beast-mastery" → lead "Ready. Show BiS, or change a pick."
	// go on "gbis:1:go:hunter::p1:" → lead "**Choose a spec first.** Then Show BiS.", class Default hunter, range Default p1
	// go on "gbis:1:go:::p1:" → "**Choose a class and a spec first.** Then Show BiS."
	// r value "50" after a nudge → lead back to the state's lead
	// linked main Zulmara: c "warrior" then c "hunter" → spec beast-mastery restored, lead the main lead
}
func TestShowBisEqualsTheTypedAnswer(t *testing.T) {
	// go on "gbis:1:go:warrior:fury:p1:" vs command bis spec:warrior-fury phase:p1
	// → same Content and Embeds (json-equal); rows equal apart from the Share token; card answer has Update true
	// row 1 labels == [Share in channel, Change picks, All 17 slots]; Change picks id "gbis:1:chg:warrior:fury:p1:"
	// slot in state "…:p1:head" → the slot form (lead "Head for a Fury Warrior, …", 3 embeds)
}
func TestUseMainReadsTheLinkNotTheID(t *testing.T) { /* me on "gbis:1:me:warrior:fury:50:" by the linked viewer → answer for hunter-beast-mastery p1; the same id pressed by "d-stranger" → failure with Update */ }
func TestChangePicksReopensTheCard(t *testing.T) { /* chg "gbis:1:chg:warrior:fury:p1:head" → card, class/spec/range Default, lead bis.lead.slot replaced by bisLeadFor (ready) */ }
func TestUseMainOnlyWhenAPickDiffers(t *testing.T) { /* linked: card at main picks → no Use; range "50" → row 4 == [Show BiS, Use Zulmara] */ }
func TestShareFromACardKeepsTheControls(t *testing.T) { /* press the answer's Share: the FollowUp public message has no "Change picks"; the Update edit (Shared) still carries Change picks */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/discordbot/ -run 'BisCardPicks|ShowBisEquals|UseMain|ChangePicks|ShareFromACard' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces; route `prefixBis` in `componentRoutes`, `updateComponents` and `cardFailureRoutes`**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Bis|Share|Router'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/bis_card.go api/internal/discordbot/bis_card_test.go api/internal/discordbot/cmd_bis.go api/internal/discordbot/routes.go api/internal/discordbot/share.go
git commit -m "feat(discordbot): /bis card picks, the nudge, Use your main and Change picks"
```

---

### Task 5: Guided `/char`

**Files:**
- Create: `api/internal/discordbot/char_card.go`, `api/internal/discordbot/char_card_test.go`
- Modify: `api/internal/discordbot/cmd_char.go`, `api/internal/discordbot/commands.go` and `copy_commands.go` (`name` optional, description `A character name. Yours if left out.`), `api/internal/discordbot/routes.go` (`gchar`), `api/internal/discordbot/share.go` (`decorateOwn` for guided char), `api/internal/discordbot/commands_wire_test.go` (`TestUnhandledCommandAnswersTheFailureSentence` uses a name no route knows, for example `nosuch`)

**Interfaces:**
- Consumes: Tasks 1 and 2 (`Modal`, `ModalValue`, `cardHandler`).
- Produces:
  - `const optGuided = "guided"`
  - `func (s *Service) guidedChar(ctx context.Context, in Interaction, guildID int64, viewer *int64, name string, update bool) (Reply, error)`. The answer is `answerChar` plus Share, with `optGuided: "1"` in the share options. Then:
    - Row 1 becomes `Share in channel`, `Someone else` (Secondary, `gchar:1:ask:`), `Character page`.
    - With 2 or more verified characters, row 2 is a select (`gchar:1:pick:`). Its placeholder is `Another of your characters`. It lists the `characterSelect` order with its 25 cap. Each option's value is the character name and its description is `{Spec} {Class} · Level {n}`. The shown character is `Default`.
  - `func (s *Service) charAskCard(content string, linked bool, own []guilds.SignupCandidate, update bool) Reply`: no embed. Content is `char.lead.unlinked`, `char.lead.none` or `char.unknown`. Row 1 is `Type a name` (Success, `gchar:1:ask:`) followed by `s.linkButton()` when not linked; with characters, the select row follows.
  - `func charModal() *Modal`: `gchar:1:name:`, `Look up a character`, `Character name`, `Zulmara`, 2 to 12 characters.
  - `cmdChar` behaviour:
    - With a `name` that resolves: as today.
    - With a `name` that does not resolve (unknown or private): `charAskCard` with `char.unknown` (§G.3.4), carrying the viewer's select when they have characters. This replaces the `lookupCopy.NoCharacter` sentence. Delete `NoCharacter` and its entry in `TestLookupCopyMatchesTheSpec` if no code still references it.
    - With no `name`, linked with characters: `guidedChar` for the main.
    - Otherwise: `charAskCard`.
  - Parts:
    - `ask`: `Reply{Modal: charModal()}`.
    - `pick`: `guidedChar` with Update.
    - `name`: `guidedChar` for the typed name with Update. An unknown or private name gives `charAskCard(fill(char.unknown,"{typed}",memberText(cut to 12)))` with Update.

- [ ] **Step 1: Write the failing tests**

```go
func TestBareCharLinkedAnswersTheMain(t *testing.T) { /* equals command char name:Zulmara on embeds; row1 labels [Share in channel, Someone else, Character page]; with Duskstrider verified: row 2 select placeholder "Another of your characters", Zulmara first and Default, description "Beast Mastery Hunter · Level 60" */ }
func TestBareCharUnlinked(t *testing.T) { /* Content == "Which character? Type a name, or link Discord to start from yours."; row [Type a name Success "gchar:1:ask:", Link Discord]; no embeds */ }
func TestTypeANameOpensTheModal(t *testing.T) { /* ask → Reply.Modal: CustomID "gchar:1:name:", Title "Look up a character", InputLabel "Character name", Placeholder "Zulmara", Min 2, Max 12 */ }
func TestModalNameAnswersInPlace(t *testing.T) { /* modal "Zulmara" → Update, the card embed; modal "Nobody" → Content "No public character named Nobody. Try another name.", Update */ }
func TestTypedUnknownCharOpensTheCard(t *testing.T) { /* command char name:Nobody → Content "No public character named Nobody. Try another name.", row [Type a name, …]; Update false */ }
func TestCharPickKeepsTheSelect(t *testing.T) { /* pick "Duskstrider" → Update, Duskstrider's card, select still present with Duskstrider Default */ }
func TestGuidedCharShareKeepsTheControls(t *testing.T) { /* Share on a guided answer → public follow-up has neither Someone else nor a select; the Shared edit keeps both */ }
func TestCharNameIsOptional(t *testing.T) { /* Commands() char option name: Required false, Description "A character name. Yours if left out." */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/discordbot/ -run 'BareChar|TypeAName|ModalName|TypedUnknownChar|CharPick|GuidedCharShare|CharNameIsOptional' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces; route `prefixChar` in `componentRoutes`, `updateComponents` and `cardFailureRoutes`**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Char|Share|Unhandled|Commands'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/char_card.go api/internal/discordbot/char_card_test.go api/internal/discordbot/cmd_char.go api/internal/discordbot/commands.go api/internal/discordbot/copy_commands.go api/internal/discordbot/routes.go api/internal/discordbot/share.go api/internal/discordbot/commands_wire_test.go
git commit -m "feat(discordbot): a bare /char answers your main, and Someone else opens a name field"
```

---

### Task 6: Guided `/signup`

**Files:**
- Create: `api/internal/discordbot/signup_card.go`, `api/internal/discordbot/signup_card_test.go`
- Modify: `api/internal/discordbot/cmd_signup.go`, `api/internal/discordbot/cmd_raid.go` (extract the embed), `api/internal/discordbot/guilds_port.go` (`MemberRaidView`), `api/internal/discordbot/commands.go` and `copy_commands.go` (`status` optional, description `In, Late, Bench or Out. Leave out to pick on a card.`), `api/internal/discordbot/routes.go` (`gsgn`)

**Interfaces:**
- Consumes: Task 1; `signupFlow`, `characterOption`, `signableRaids`, `raidLocation`.
- Produces:
  - `GuildService` gains `MemberRaidView(ctx context.Context, guildID, viewerID, raidID int64) (guilds.RaidView, bool, error)`. `*guilds.Service` already has it.
  - `func (s *Service) raidAnswerEmbed(ctx context.Context, facts guilds.RaidPostFacts, movedFrom *time.Time, own *guilds.NextRaidOwn, cands []guilds.SignupCandidate) notify.Embed`. It is extracted from `cmdRaid`, which uses it unchanged.
  - `type signupPicks struct{ RaidID int64; CharacterKey string }`.
  - `func (s *Service) signupCard(ctx context.Context, a signupAsk, raids []guilds.Raid, verified []guilds.SignupCandidate, p signupPicks, update bool) (Reply, error)`. The rows (spec §G.4.4):
    - The embed is `raidAnswerEmbed` for the chosen raid, with the chosen character's state read from `MemberRaidView(...).Viewer.Characters`.
    - With 2 or more raids, a raid select (`gsgn:1:raid:{raid}:{key}`). Its placeholder is `Choose a raid`. Options are labelled `{Title} · Ddd d Mmm` and described `{Guild time} · {signed} of {size} signed up`, or `{n} signed up` when the size is unknown.
    - With 2 or more characters, a character select (`gsgn:1:char:{raid}:{key}`), using `characterOption`, placeholder `Choose a character`.
    - The status row: `In` (Success), then `Late`, `Bench`, `Out` (Secondary), with ids `gsgn:1:{state}:{raid}:{key}`, then `Raid page` (Link).
  - `cmdSignup` with no `status`, after the member gate:
    - With no signable raid: `pressCopy.NoRaid`.
    - Otherwise the card. The raid defaults to the typed `raid`, else the next raid. The character defaults to the typed `character`, else the main.
  - Parts:
    - `raid` and `char`: re-render with Update.
    - Each state word: `runFlow` with `signupAsk{RaidID, State, CharacterKey, Update:true}`. A forged key gives `pressCopy.Failure` through `chosen`.

- [ ] **Step 1: Write the failing tests**

```go
func TestBareSignupCard(t *testing.T) { /* one raid, one character: no selects, row [In Success, Late, Bench, Out, Raid page]; ids "gsgn:1:in:{raid}:us/normal/zulmara"; two raids and two characters: raid select (Default the next, label "Onyxia's Lair · {Ddd d Mmm}") and character select (Default the main) */ }
func TestSignupCardPressSignsUp(t *testing.T) { /* press in → Update, the §3.3.2 outcome sentence verbatim, one signup row */ }
func TestSignupCardRefusesAForgedCharacter(t *testing.T) { /* "gsgn:1:in:{raid}:us/normal/someone-else" → Content == pressCopy.Failure ("That did not save. Try again."), no signup */ }
func TestSignupCardRaidLockedSinceDrawn(t *testing.T) { /* lock the raid, press → pressCopy.Locked, Update */ }
func TestSignupCardSurvivesHostileTitles(t *testing.T) { /* 80-rune title "@everyone **x**…" → assertCardLimits, labels ≤100, allowed_mentions empty */ }
func TestSignupStatusIsOptional(t *testing.T) { /* status Required false, Description "In, Late, Bench or Out. Leave out to pick on a card." */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/discordbot/ -run 'SignupCard|BareSignup|SignupStatusIsOptional' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces; route `prefixSignup` in `componentRoutes`, `updateComponents` and `cardFailureRoutes`**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Signup|Raid|Press'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/signup_card.go api/internal/discordbot/signup_card_test.go api/internal/discordbot/cmd_signup.go api/internal/discordbot/cmd_raid.go api/internal/discordbot/guilds_port.go api/internal/discordbot/commands.go api/internal/discordbot/copy_commands.go api/internal/discordbot/routes.go
git commit -m "feat(discordbot): a bare /signup answers with the raid, the character and four buttons"
```

---

### Task 7: `/attendance` for officers: the lowest-attendance select and a typed name

**Files:**
- Create: `api/internal/guilds/discord_reads_lowest.go`, `api/internal/guilds/discord_reads_lowest_test.go`, `api/internal/discordbot/officer_cards.go`, `api/internal/discordbot/officer_cards_test.go`
- Modify: `api/internal/discordbot/cmd_attendance.go`, `api/internal/discordbot/guilds_port_officer.go`, `api/internal/discordbot/routes.go` (`gatt`)

**Interfaces:**
- Consumes: Tasks 1 and 2.
- Produces:
  - `func (s *Service) LowestAttendance(ctx context.Context, guildID, viewerID int64, limit int) ([]AttendanceRecord, error)` in guilds:
    - Officers only, with `seesEveryAttendanceRow`'s rule; anyone else gets `ErrOfficerOnly`.
    - It covers the guild's verified characters that have a chip (Ruling 10), read through `attendanceRoster`.
    - Sorted by came over the last-8 window ascending, then by name. At most `limit`.
  - `OfficerReads` gains `LowestAttendance`.
  - `func (s *Service) attendanceOfficerRows(ctx context.Context, guildID int64, recs []guilds.AttendanceRecord, shown string) []ActionRow`:
    - A select `gatt:1:pick:`, placeholder `Look up a member`, labels the names, descriptions the chip verbatim, `Default` on `shown`.
    - The `Roster` row becomes `Type a name` (Secondary, `gatt:1:ask:`) then `Roster`.
  - `cmdAttendance` for an officer with no `name`: today's answer plus those rows. A member's answer is unchanged.
  - Parts:
    - `pick`: answers that member's record with the rows kept. Rank is re-checked through `AttendanceFor`, so a demoted officer gets `officerCopy.OfficersOnly`.
    - `ask`: a modal `gatt:1:name:` titled `Look up attendance`, otherwise the same input as `/char`'s.
    - `name`: answers the record; a name that is not a member gives the officer's own answer with content `No member of {Guild} named {typed}.`.

- [ ] **Step 1: Write the failing tests**

```go
// guilds
func TestLowestAttendanceIsOfficersOnlyLowestFirst(t *testing.T) { /* three members came 3/8, 7/8, 5/8 → order 3,5,7; one with no night left out; a member viewer → ErrOfficerOnly; limit 2 → two */ }
// discordbot
func TestOfficerAttendanceSelect(t *testing.T) { /* officer bare → select placeholder "Look up a member", first option Description "Came 3 of 8 · 10 of 30 raids", nothing Default; last row [Type a name Secondary, Roster Link]; member bare → no select */ }
func TestOfficerAttendancePickAndDemotion(t *testing.T) { /* pick "Thalwyn" → Update, Thalwyn's embed, select Default Thalwyn; demote the officer, pick → "Only officers can look up another member's attendance." */ }
func TestAttendanceModalUnknownName(t *testing.T) { /* ask → Modal title "Look up attendance"; name "Nobody" → Content "No member of Olympus XXVII named Nobody.", officer's own embeds below */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/guilds/ ./internal/discordbot/ -run 'LowestAttendance|OfficerAttendance|AttendanceModal' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces; route `prefixAttendance` in `componentRoutes`, `updateComponents` and `cardFailureRoutes`**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run Attendance` and `$TDB go test -p 1 ./internal/guilds/ -run Attendance`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/guilds/discord_reads_lowest.go api/internal/guilds/discord_reads_lowest_test.go api/internal/discordbot/officer_cards.go api/internal/discordbot/officer_cards_test.go api/internal/discordbot/cmd_attendance.go api/internal/discordbot/guilds_port_officer.go api/internal/discordbot/routes.go
git commit -m "feat(discordbot): officers look up the lowest attendance from a select or a typed name"
```

---

### Task 8: `/readiness` gains a raid; `/readiness` and `/report` gain their selects

**Files:**
- Modify: `api/internal/guilds/discord_reads_readiness.go` (+ test), `api/internal/discordbot/cmd_readiness.go`, `api/internal/discordbot/cmd_report.go`, `api/internal/discordbot/officer_cards.go`, `api/internal/discordbot/officer_cards_test.go`, `api/internal/discordbot/guilds_port_officer.go`, `api/internal/discordbot/autocomplete.go` and `routes.go` (`readiness` raid autocomplete, `grdy`, `grep`), `api/internal/discordbot/commands.go` (`/readiness` `raid` option, `descRaidOption`), `api/internal/discordbot/share.go` (`decorateOwn` for report)

**Interfaces:**
- Consumes: Tasks 1 and 2; `raidChoices`, `DoneRaids`, `nightLine`.
- Produces:
  - `func (s *Service) RaidReadiness(ctx context.Context, guildID, viewerID, raidID int64) (ReadinessAnswer, error)` in guilds. It uses `NextRaidReadiness`'s gate and refusals, and reads the named planned raid of the guild. Another guild's raid, or a raid that is not planned, gives `ReadinessAnswer{}` (nil Raid).
  - `OfficerReads` gains `RaidReadiness`.
  - `/readiness` gains `stringOption("raid", descRaidOption, false, true, nil)`, and `autocompleteRoutes` gains `"readiness": s.autocompleteReadiness`, which is `raidChoices` for an officer and empty for anyone else.
  - `func (s *Service) readinessRaidRow(raids []guilds.Raid, shown int64, loc *time.Location) []ActionRow` is present only with 2 or more raids. It is a select `grdy:1:raid:`, placeholder `Another raid`, labels `{Title} · Ddd d Mmm`, descriptions `{Guild time}`, `Default` on `shown`.
  - `func (s *Service) reportNightRow(ctx context.Context, guildID int64, viewer int64, shown int64) ([]ActionRow, error)` is present only with 2 or more done raids, from the last 10 newest first. It is a select `grep:1:raid:`, placeholder `Another night`, labels `{Title} · Ddd d Mmm`, descriptions the night line cut to 100 characters, `Default` on `shown`. It follows the Share and `Open the report` row, and the coach lines are unchanged.
  - Parts:
    - `grdy` `raid`: answers that raid's board, re-checking rank, plan and freeze on every pick (the bot spec's refusals).
    - `grep` `raid`: answers that night with a fresh Share when the log is public.
  - `decorateOwn("report", …)` re-adds the night row.

- [ ] **Step 1: Write the failing tests**

```go
func TestRaidReadinessNamesARaid(t *testing.T) { /* guilds: the second planned raid's board; another guild's raid → nil Raid; member → ErrNotOfficer */ }
func TestReadinessSelect(t *testing.T) { /* two planned raids → select Default the next, placeholder "Another raid"; pick the second → Update, that raid's board; one raid → no select; officer demoted then pick → the officers-only refusal */ }
func TestReadinessRaidOption(t *testing.T) { /* Commands(): readiness raid optional, autocomplete, Description "Which raid. The next one if left out." */ }
func TestReportNightSelect(t *testing.T) { /* two done raids → select after the Share row, Default the last night, description ≤100; pick the other → Update, that report; Share present only when that log is public */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/guilds/ ./internal/discordbot/ -run 'RaidReadiness|ReadinessSelect|ReadinessRaidOption|ReportNightSelect' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces; route `prefixReadiness` and `prefixReport` in `componentRoutes`, `updateComponents` and `cardFailureRoutes`**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Readiness|Report|Share|Commands'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/guilds/discord_reads_readiness.go api/internal/guilds/*readiness*_test.go api/internal/discordbot/cmd_readiness.go api/internal/discordbot/cmd_report.go api/internal/discordbot/officer_cards.go api/internal/discordbot/officer_cards_test.go api/internal/discordbot/guilds_port_officer.go api/internal/discordbot/autocomplete.go api/internal/discordbot/routes.go api/internal/discordbot/commands.go api/internal/discordbot/share.go
git commit -m "feat(discordbot): /readiness takes a raid; readiness and report answers change scope from a select"
```

- [ ] **Step 6: The shared sweep**

Add `TestEveryCardHoldsDiscordsLimits` to `card_test.go`. It:
- walks every card frame Tasks 3 to 8 produce through `assertCardLimits`;
- asserts each of the six card prefixes is in `componentRoutes`, `updateComponents` and `cardFailureRoutes`, and that `gchar` and `gatt` also route a type-5 submit;
- asserts `len(Commands())` is unchanged and `commandsHash(Commands())` differs from the hash of the pre-amendment command set (kept as a literal in the test).

Run `cd api && $TDB go test -p 1 ./internal/discordbot/`. Expected: PASS. Commit with `test(discordbot): every guided card holds Discord's limits and routes`.

---

### Task 9: Migration 0057, the schedule store, and the remembered reminder default

**Files:**
- Create: `api/internal/db/migrations/0057_discord_guided_posting.up.sql`, `.down.sql`, `api/internal/guilds/raids_schedule_store.go`, `api/internal/guilds/raids_schedule_store_test.go`
- Modify: `api/internal/notify/guild_store.go` (+ test: `RemindDefault`, `SetDefaults`), `api/internal/discordbot/store_install.go` (+ test: `RemindDefault`, `SetDiscordDefaults`, the install insert copies both defaults)

**Interfaces:**
- Produces:
  - Up migration:
    - `guild_discord_installs` and `guild_notification_webhooks`: `alter column post_default set default true`, `add column remind_default boolean not null default true`, and `update … set post_default = true` (Ruling 5).
    - `guild_raid_discord_posts` gains `remind_due boolean not null default false`, `reminder_message_id text`, `reminded_at timestamptz`.
    - New table:
      ```sql
      create table guild_raid_discord_schedule (
        raid_id bigint primary key references guild_raids(id) on delete cascade,
        guild_id bigint not null references guilds(id) on delete cascade,
        post_at timestamptz not null,
        remind boolean not null,
        posted_at timestamptz,
        reminded_at timestamptz,
        reminder_via text check (reminder_via in ('bot','webhook')),
        reminder_message_id text,
        lease_until timestamptz,
        attempts int not null default 0,
        next_attempt_at timestamptz not null default now(),
        last_error text,
        created_at timestamptz not null default now());
      ```
      An index on `(next_attempt_at)` where `posted_at is null or reminded_at is null`.
    - The Ruling 6 backfill from `guild_raid_discord_posts`.
  - The down migration reverses every step except the data flip.
  - In guilds:
    - Constants: `postHorizon = 7 * 24 * time.Hour`, `reminderLead = 24 * time.Hour`, `reminderFresh = 12 * time.Hour`, `scheduleLease = 2 * time.Minute`, `scheduleBatch = 50`.
    - `func postAtFor(startsAt, now time.Time) time.Time` returns `max(now, startsAt - postHorizon)`.
    - `type RaidSchedule struct{ RaidID, GuildID int64; PostAt time.Time; Remind bool; PostedAt, RemindedAt *time.Time; Reminder *notify.ReminderRef }`. `notify.ReminderRef` is created in this task, in the new file `notify/reminder.go`: `type ReminderRef struct{ Via, MessageID string }` with `const (ReminderViaBot = "bot"; ReminderViaWebhook = "webhook")`. Task 10 adds the reminder copy and functions to the same file.
    - `type dueKind int` with `dueKind{duePost, dueReminder}`, and `type DueSchedule struct{ RaidSchedule; Kind dueKind; Attempts int }`.
    - `(s *Store)` methods:
      - `UpsertSchedule(ctx, raidID, guildID int64, postAt time.Time, remind bool) error`: an unposted row's `post_at` and `remind` are replaced; a posted row keeps `post_at` and `posted_at` and takes `remind`.
      - `SetScheduleRemind(ctx, raidID int64, remind bool) error`
      - `DeleteUnpostedSchedule(ctx, raidID int64) error`
      - `MarkSchedulePosted(ctx, raidID int64, at time.Time) error`
      - `MarkScheduleReminded(ctx, raidID int64, ref notify.ReminderRef, at time.Time) error`
      - `ResetScheduleReminder(ctx, raidID int64) (*notify.ReminderRef, error)`: clears `reminded_at` and the reference and returns the old reference.
      - `SchedulesOf(ctx, raidIDs []int64) (map[int64]RaidSchedule, error)`
      - `ClaimDueSchedules(ctx, now time.Time, limit int) ([]DueSchedule, error)`: one `update … where raid_id in (select … for update skip locked) returning`, leasing for `scheduleLease`. Due means `not leased`, `next_attempt_at <= now`, the raid `status = 'planned'` and `starts_at > now`, and either `posted_at is null and post_at <= now` or `remind and posted_at is not null and reminded_at is null and starts_at - 24h <= now and posted_at <= starts_at - 36h` (Ruling 3).
      - `FailSchedule(ctx, raidID int64, retryAt time.Time, lastErr string) error`
      - `ReleaseSchedule(ctx, raidID int64) error`
  - In notify: the `GuildWebhook.RemindDefault` field, and `(s *GuildStore) SetDefaults(ctx, guildID int64, post, remind bool) error`. In discordbot: `Install.RemindDefault`, and `(s *Store) SetDiscordDefaults(ctx, guildID int64, post, remind bool) error`. Both are added beside the existing `SetPostDefault`, which stays until Task 12 switches the `guilds.DiscordBot` interface and deletes it, so the tree builds between the two tasks.

- [ ] **Step 1: Write the failing tests**

```go
func TestPostAtFor(t *testing.T) { /* start now+3d → now; start now+10d → start-7d */ }
func TestScheduleDueWindows(t *testing.T) {
	// post_at past, unposted, planned → one duePost; cancelled raid → none; locked → none
	// posted 3 days before start, remind on, now = start-23h → dueReminder; remind off → none
	// posted 30h before start, now = start-23h → none (fresh post, 12 h rule)
	// a claimed row is not claimed again inside its lease; two concurrent claims split rows (skip locked)
}
func TestUpsertKeepsAPostedRow(t *testing.T) { /* MarkSchedulePosted, then UpsertSchedule with a new post_at → post_at and posted_at unchanged, remind taken */ }
func TestResetReturnsTheOldReminder(t *testing.T) { /* MarkScheduleReminded(webhook, "m1") → Reset returns {webhook, m1}; second Reset returns nil */ }
func TestDefaultsStartOnAndAreRemembered(t *testing.T) { /* new install: PostDefault true, RemindDefault true; SetDiscordDefaults(false,true) read back; same for the webhook row */ }
func TestMigrationBackfillsBotPosts(t *testing.T) { /* a guild_raid_discord_posts row created before 0057 has a schedule row posted, remind false (run via the migration test harness the package already uses for 0053) */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/guilds/ ./internal/notify/ ./internal/discordbot/ ./internal/db/ -run 'PostAtFor|ScheduleDue|UpsertKeeps|ResetReturns|DefaultsStartOn|MigrationBackfills' -v`
Expected: FAIL.

- [ ] **Step 3: Write the migration and implement the Interfaces**

No service-level method changes in this task. Task 12 adds `notify.Service.SetGuildDiscordDefaults` and `discordbot.Service.SetDiscordDefaults`, and deletes the old `SetPostDefault`/`SetGuildPostDefault` path.

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/notify/ ./internal/discordbot/ -run 'Guild|Install|Store'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/db/migrations/0057_discord_guided_posting.up.sql api/internal/db/migrations/0057_discord_guided_posting.down.sql api/internal/guilds/raids_schedule_store.go api/internal/guilds/raids_schedule_store_test.go api/internal/notify/reminder.go api/internal/notify/guild_store.go api/internal/notify/guild_store_test.go api/internal/discordbot/store_install.go api/internal/discordbot/store_test.go
git commit -m "feat(api): migration 0057, the raid posting schedule and the remembered reminder default"
```

---

### Task 10: notify: the `Tomorrow` post, the reminder line, and webhook reminders

**Files:**
- Create: `api/internal/notify/reminder_test.go`
- Modify: `api/internal/notify/raid_embed.go` (`PostTomorrow`, copy), `api/internal/notify/reminder.go`, `api/internal/notify/discord.go` (`Poster.PostContent`, `Poster.DeleteMessage`, and the fake), `api/internal/notify/guild_deliver.go` (`RaidReminder`, `RaidReminderWithdrawn`)

**Interfaces:**
- Produces:
  - `const PostTomorrow RaidPostState = "tomorrow"`. Its title prefix is `Tomorrow: `, so the title reads `Tomorrow: {Title} · {Weekday} {d} {Month}`. Its description is `post.tomorrowDescription`. Its colour and fields are open's.
  - `raidCopy` gains `TomorrowPrefix`, `ReminderLead`, `ReminderCount`, `ReminderCountOpen`, `ReminderNobody`, `ReminderWebhookLink`.
  - `func ReminderLine(in RaidPosted) string`:
    - `**{Title} is tomorrow**, <t:{unix}:F> (<t:{unix}:R>). {signed} of {size} signed up so far.`
    - `{signed}` is `Counts.In + Counts.Late`, as in `SignedUpValue`.
    - With an unknown size: `{n} signed up so far.`
    - With nobody signed: `Nobody has signed up yet.`
    - The title is escaped.
  - `func WebhookReminderContent(in RaidPosted) string` is `ReminderLine` + ` ` + `[Sign up on Forever Sixty]({RaidURL})`.
  - `Poster` gains:
    - `PostContent(ctx context.Context, webhookURL, content string) (messageID string, err error)`: POST `?wait=true` with `{content, username, avatar_url, allowed_mentions:{parse:[]}}`, returning the answer's `id`.
    - `DeleteMessage(ctx context.Context, webhookURL, messageID string) error`: DELETE `{url}/messages/{id}`; a 404 is success.
    - Neither ever wraps the URL in an error.
  - `func (d *GuildDeliverer) RaidReminder(ctx context.Context, guildID int64, e RaidPosted) (ReminderRef, bool, error)`: `false, nil` with no channel or with the channel switched off; it counts failures with the five-failure rule.
  - `func (d *GuildDeliverer) RaidReminderWithdrawn(ctx context.Context, guildID int64, ref ReminderRef)`: webhook references only, failure logged.

- [ ] **Step 1: Write the failing tests**

```go
func TestReminderCopyMatchesTheSpec(t *testing.T) { /* spectable P.7: post.tomorrowTitle, post.tomorrowDescription, reminder.line (3 forms), reminder.webhookLink */ }
func TestTomorrowEmbed(t *testing.T) { /* RaidPostEmbed(state PostTomorrow) Title "Tomorrow: Onyxia's Lair · Thursday 10 December", Description "The raid is tomorrow. Sign up with the buttons below." */ }
func TestReminderLine(t *testing.T) { /* 33 of 40 → "**Onyxia's Lair is tomorrow**, <t:1796954400:F> (<t:1796954400:R>). 33 of 40 signed up so far."; size unknown 12 → "… 12 signed up so far."; nobody → "… Nobody has signed up yet." */ }
func TestReminderLineSurvivesHostileTitles(t *testing.T) { /* "@everyone **x**" escaped; ≤2000 chars with an 80-rune title */ }
func TestWebhookReminderPostsAndDeletes(t *testing.T) { /* httptest Discord: POST has wait=true, no "embeds", allowed_mentions parse []; returns id "m1"; RaidReminderWithdrawn DELETEs /messages/m1; a 404 on delete is success; the test's log buffer never contains the webhook token */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/notify/ -run 'ReminderCopy|TomorrowEmbed|ReminderLine|WebhookReminder' -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/notify/`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/notify/raid_embed.go api/internal/notify/reminder.go api/internal/notify/reminder_test.go api/internal/notify/discord.go api/internal/notify/guild_deliver.go
git commit -m "feat(notify): the Tomorrow post state, the reminder line and webhook reminders"
```

---

### Task 11: The bot's reminder: route, drain, and withdraw

**Files:**
- Create: `api/internal/discordbot/drain_reminder.go`, `api/internal/discordbot/drain_reminder_test.go`
- Modify: `api/internal/guilds/raids_discord.go` (`RaidEvents` gains two methods; update `recordingRaidEvents` in `raids_discord_test.go`), `api/internal/discordbot/router.go`, `api/internal/discordbot/rest.go` (+ `fakes_test.go`), `api/internal/discordbot/message.go` (`ReplyTo`), `api/internal/discordbot/store_posts.go`, `api/internal/discordbot/drain.go`

**Interfaces:**
- Consumes: Task 10 (`RaidPosted`, `ReminderLine`, `ReminderRef`, `GuildDeliverer.RaidReminder`), Task 9 (the columns).
- Produces:
  - `guilds.RaidEvents` gains:
    - `RaidReminder(ctx context.Context, guildID int64, e notify.RaidPosted) (notify.ReminderRef, bool, error)`
    - `RaidReminderWithdrawn(ctx context.Context, guildID int64, ref notify.ReminderRef)`
    - Both `*notify.GuildDeliverer` and `*discordbot.RaidRouter` satisfy it.
  - `RaidRouter.RaidReminder`: when the bot is active and the raid has a post row that is not gone, `Store.MarkReminderDue(ctx, raidID)` returns `{Via: bot}, true`. Otherwise it is `Webhook.RaidReminder`.
  - `RaidRouter.RaidReminderWithdrawn`: a bot reference is a no-op, because the drainer withdraws on the raid's own mark. A webhook reference goes to `Webhook.RaidReminderWithdrawn`.
  - `REST` gains:
    - `LatestMessageID(ctx context.Context, channelID string) (string, error)`: GET `/channels/{id}/messages?limit=1`, `""` when empty.
    - `DeleteMessage(ctx context.Context, channelID, messageID string) error`: Discord 10008 is success.
  - `Message.ReplyTo string` marshals as `message_reference:{message_id, fail_if_not_exists:false}`.
  - `Post` gains `RemindDue bool`, `ReminderMessageID *string`, `RemindedAt *time.Time`, which `ClaimDue` returns.
  - `(s *Store) MarkReminderDue(ctx, raidID int64) error` sets `remind_due` and marks the row dirty.
  - `(s *Store) FinishReminder(ctx, p Post, reminderID *string, now time.Time) error` is lease-guarded and sets `remind_due=false, reminded_at=now, reminder_message_id`.
  - `(s *Store) ClearReminder(ctx, p Post, rearm bool) error` is lease-guarded and nulls `reminder_message_id`; it also nulls `reminded_at` when `rearm`.
  - `func lookOf(state notify.RaidPostState, p Post, startsAt, now time.Time) notify.RaidPostState` returns `PostTomorrow` when the state is open, `p.RemindedAt != nil` and `now` is before the start; otherwise the state.
  - `wakeAt` gains a case: `PostTomorrow` wakes at `StartsAt`, so that at the start the post returns to its open look and from then on follows §3.3.1.
  - `func (d *Drainer) remind(ctx context.Context, p Post, facts guilds.RaidPostFacts, now time.Time) error`. It runs inside `drainRow` after the post's own send, when `p.RemindDue && state == open`:
    1. `LatestMessageID`.
    2. Edit the post in `PostTomorrow`.
    3. If the newest message is not the post: `CreateMessage` with `Content: notify.ReminderLine`, the post's own `postRows` (state open) and `ReplyTo` the post.
    4. `FinishReminder`.
    - When `RemindDue` and the state is not open, it is `FinishReminder(nil)` with no send.
  - `func (d *Drainer) withdraw(ctx context.Context, p Post, state notify.RaidPostState) error`: when `p.ReminderMessageID != nil` and the state is moved, locked or cancelled, it calls `DeleteMessage`, then `ClearReminder(rearm: state == moved)`.
  - Discord refusals on either call go through the existing `refused`.

- [ ] **Step 1: Write the failing tests**

```go
func TestReminderEditsTheNewestPost(t *testing.T) { /* LatestMessageID == post id → one EditMessage whose embed Title starts "Tomorrow: ", no CreateMessage; row reminded_at set, remind_due false */ }
func TestReminderRepliesWhenTheChannelMovedOn(t *testing.T) { /* LatestMessageID "other" → EditMessage(Tomorrow) + CreateMessage{Content: ReminderLine, ReplyTo: post id, 5 components with raid: ids, allowed_mentions parse []}; reminder_message_id stored */ }
func TestTomorrowStaysOnLaterEdits(t *testing.T) { /* after the reminder, a sign-up mark → the edit keeps the Tomorrow title; wake_at == start; drained at the start → the open title again */ }
func TestDrainerWithdrawsTheReminder(t *testing.T) { /* cancel → DeleteMessage(reminder) and Cancelled post; move → DeleteMessage, Moved post, reminded_at cleared; Discord 10008 on delete → still cleared */ }
func TestReminderDueIsDrainedOnce(t *testing.T) { /* two drainers, one due row → one CreateMessage */ }
func TestRouterSendsTheReminderToTheBotOrTheWebhook(t *testing.T) { /* active bot with a post row → MarkReminderDue, ref bot; no row → webhook RaidReminder called; Withdrawn(bot) → no webhook delete */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/discordbot/ ./internal/guilds/ -run 'Reminder|Tomorrow|Withdraws|RouterSends' -race -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/discordbot/ -run 'Drain|Router|Rest' -race` and `$TDB go test -p 1 ./internal/guilds/ -run Discord`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/discordbot/drain_reminder.go api/internal/discordbot/drain_reminder_test.go api/internal/discordbot/router.go api/internal/discordbot/rest.go api/internal/discordbot/fakes_test.go api/internal/discordbot/message.go api/internal/discordbot/store_posts.go api/internal/discordbot/drain.go api/internal/guilds/raids_discord.go api/internal/guilds/raids_discord_test.go
git commit -m "feat(discordbot): the day-before reminder edits the post or replies under it, and is withdrawn on a change"
```

---

### Task 12: guilds: scheduling on save, the remembered defaults, the raid view, and the schedule run

**Files:**
- Create: `api/internal/guilds/raids_schedule.go`, `api/internal/guilds/raids_schedule_test.go`, `api/internal/guilds/raid_url_test.go`
- Modify: `api/internal/guilds/raids_handler.go` (create, patch, status actions), `api/internal/guilds/raids_discord.go` (`rememberDiscordDefaults`), `api/internal/guilds/raids_validate.go` (`RemindDayBefore *bool` `json:"remind_day_before"`), `api/internal/guilds/raids_view.go` (`RaidDiscord.RemindDefault`, `ChannelName *string`; `RaidView.DiscordPost *RaidDiscordPost`), `api/internal/guilds/discord.go` (`DiscordBot` and `raidDiscord`, `raidViews`), `api/internal/discordbot/install.go` (implement the new `DiscordBot` method)

**Interfaces:**
- Consumes: Task 9 (store), Task 11 (`RaidEvents.RaidReminder`, `RaidReminderWithdrawn`).
- Produces:
  - `DiscordBot` interface: `Active(ctx, guildID) (active, postDefault bool, err error)` is replaced by `Destination(ctx context.Context, guildID int64) (BotDestination, error)` with `type BotDestination struct{ Active, PostDefault, RemindDefault bool; ChannelName string }`. `SetPostDefault` is replaced by `SetDiscordDefaults(ctx, guildID int64, post, remind bool) error`.
  - `RaidDiscord{Available, Connected, PostDefault, RemindDefault bool; ChannelName *string}`. JSON `remind_default`, and `channel_name` (null for a webhook or when not connected).
  - `type RaidDiscordPost struct{ PostedAt, PostAt, RemindedAt *time.Time; Remind bool }`. JSON `posted_at`, `post_at`, `reminded_at`, `remind` (P.8). `RaidView.DiscordPost` is `json:"discord_post"`, null without a schedule row.
  - `func (s *Service) rememberDiscordDefaults(ctx context.Context, guildID int64, post, remind *bool)` stores both on the webhook (`notify.Service.SetGuildDiscordDefaults`, new) and on the install (`discordbot.Service.SetDiscordDefaults`, new) when the save carried them. Delete `rememberPostDefault`, `SetGuildPostDefault`, both `SetPostDefault` methods and the two stores' `SetPostDefault`.
  - Files this also touches: `api/internal/notify/guild_deliver.go`, `api/internal/notify/guild_store.go`, `api/internal/discordbot/store_install.go`. Add them to the commit.
  - `func (s *Service) scheduleCreated(ctx context.Context, guildID int64, views []RaidView, remind bool, now time.Time)`. For each view it upserts `postAtFor(start, now)`. Those due now go through the existing `postCreated` loop and are marked posted. A repeat of four posts only the first raid within 7 days (P.4).
  - `func (s *Service) scheduleEdited(ctx context.Context, guildID int64, before, after Raid, req raidRequest, now time.Time)`:
    - **Posted row:** `remind_day_before`, when present, goes to `SetScheduleRemind`. On a moved start, reset the reminder, call `RaidReminderWithdrawn` for a webhook reference, and keep today's "post if the time changes" webhook path.
    - **Unposted or no row:** `post_to_discord` true upserts `postAtFor(after.StartsAt, now)` and posts at once when due. False deletes the unposted row.
  - `func (s *Service) withdrawOnStatus(ctx context.Context, guildID, raidID int64)`: cancel and lock reset the reminder and call `RaidReminderWithdrawn`.
  - `type ScheduleRun struct{ Posted, Reminded, Failed int }`.
  - `func (s *Service) RunRaidSchedule(ctx context.Context, now time.Time) (ScheduleRun, error)`:
    - It claims due rows in batches of `scheduleBatch` until none are left or `ctx` ends.
    - Posts go through `RaidPostFacts` then `RaidEvents.RaidPosted`, and are then marked posted.
    - Reminders go through `RaidEvents.RaidReminder`. `ok` marks them reminded. `!ok` (no channel) also marks them reminded with a nil reference, so they are not retried every run.
    - An error does `FailSchedule` with backoff (2 min doubling, cap 1 h) and counts `Failed`.
    - It returns an error only when every claimed row failed.
  - `create` and `patch` read `remind_day_before`. An absent field reads as the guild's `remind_default`.
  - Test helper `func parseGuildRaidPathGo(path string) bool` in `raid_url_test.go` mirrors `web/src/lib/characters.ts` `parseGuildRaidPath`: exactly six segments, `guild`, region, ruleset, slug, `raids`, digits only.

- [ ] **Step 1: Write the failing tests**

```go
func TestSaveWithinAWeekPostsAtOnce(t *testing.T) { /* POST raid starting in 3 days, post_to_discord true → one RaidPosted; view.discord_post.posted_at set, post_at == posted_at */ }
func TestSaveFurtherOutPostsAWeekBefore(t *testing.T) { /* start in 10 days → no RaidPosted; discord_post.post_at == start-7d, posted_at null */ }
func TestRepeatPostsOneWeekAtATime(t *testing.T) { /* repeat_weeks 3, first in 2 days → one RaidPosted; RunRaidSchedule at now+6d → the second posts; at now+13d the third */ }
func TestReminderRunsTheDayBefore(t *testing.T) { /* posted 5 days ahead, RunRaidSchedule at start-23h → one RaidReminder; again → none; remind off → none; posted at start-30h → none (12 h rule) */ }
func TestScheduleFollowsMovesCancelsAndLocks(t *testing.T) {
	// unposted, PATCH start to 3 days → posts at once; to 20 days → post_at follows
	// reminded (webhook ref m1), PATCH start +1h → RaidReminderWithdrawn({webhook,m1}), reminded_at null, re-sent by the next run at the new start-24h
	// cancel → Withdrawn called; RunRaidSchedule never posts or reminds a cancelled or locked raid
	// PATCH post_to_discord false on an unposted raid → row gone
}
func TestTwoScheduleRunsSendOnce(t *testing.T) { /* two RunRaidSchedule goroutines on one due row → one RaidPosted */ }
func TestDefaultsAreRememberedAndRead(t *testing.T) { /* save post false, remind false → next GET raids discord.post_default false, remind_default false; a fresh guild reads true, true; bot active → channel_name "raid-signups"; webhook → channel_name null */ }
func TestEveryPostedRaidURLIsTheSitePath(t *testing.T) { /* for each RaidPosted the recorder saw, and for a guild named "Olympus XXVII": the path of RaidURL passes parseGuildRaidPathGo; "https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/1" passes */ }
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd api && $TDB go test -p 1 ./internal/guilds/ -run 'SaveWithin|SaveFurther|RepeatPosts|ReminderRuns|ScheduleFollows|TwoScheduleRuns|DefaultsAreRemembered|EveryPostedRaidURL' -race -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces**

Keep `raids_handler.go` under 800 lines by putting the schedule decisions in `raids_schedule.go`.

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./internal/guilds/ -run 'Raid|Discord|Settings'` and `$TDB go test -p 1 ./internal/discordbot/ -run 'Install|Router'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/internal/guilds/raids_schedule.go api/internal/guilds/raids_schedule_test.go api/internal/guilds/raid_url_test.go api/internal/guilds/raids_handler.go api/internal/guilds/raids_discord.go api/internal/guilds/raids_validate.go api/internal/guilds/raids_view.go api/internal/guilds/discord.go api/internal/discordbot/install.go api/internal/discordbot/store_install.go api/internal/notify/guild_deliver.go api/internal/notify/guild_store.go
git commit -m "feat(guilds): raids post when saved or a week ahead, remind the day before, and remember both"
```

---

### Task 13: The `raid-posts` job

**Files:**
- Create: `api/cmd/api/raid_posts.go`, `api/cmd/api/raid_posts_test.go`
- Modify: `api/cmd/api/main.go` (the `raid-posts` case), `api/cmd/api/discordbot.go` (share the router building with the job), `.github/workflows/api.yml` (add `raid-posts` to the job loop), `api/README.md` (the job and its schedule)

**Interfaces:**
- Consumes: Task 12 `RunRaidSchedule`, Task 11 `RaidRouter`, `Drainer.DrainOnce`.
- Produces:
  - `func runRaidPosts(ctx context.Context, log *slog.Logger) error` mirrors `runBillingDaily`: `start()`, build, run.
  - `type raidPostsServices struct{ Guilds *guilds.Service; Drainer *discordbot.Drainer }`. `Drainer` is nil when the bot is not configured.
  - `func newRaidPostsServices(cfg config.Config, pool *pgxpool.Pool, log *slog.Logger) (raidPostsServices, error)`. It builds the same guilds, notify and discordbot wiring `serve` uses for the raid router, but never starts a drainer goroutine and never registers commands. Factor `func raidRouterFor(bot *discordbot.Service, deps server.Deps) guilds.RaidEvents` out of `newDiscordBot` so both use it.
  - `const raidPostsDrainPasses = 5`.
  - `func raidPosts(ctx context.Context, log *slog.Logger, svc raidPostsServices, now time.Time) error`:
    1. `RunRaidSchedule`.
    2. Up to `raidPostsDrainPasses` `DrainOnce` calls when `Drainer != nil`.
    3. Log `raid-posts` with `posted`, `reminded` and `failed` counts only.
    4. Return wrapped errors as `billingDaily` does.
  - README: under the jobs section, the `gcloud run jobs create raid-posts … --args raid-posts --cpu 1 --memory 512Mi --task-timeout 5m`, using the service's env and secrets by name. Never paste a value. Then `gcloud scheduler jobs create http raid-posts-every-10m --schedule "*/10 * * * *" …` in the `stripe-reconcile-nightly` shape. Add one line saying why: deferred posts and reminders, even when the service is cold.

- [ ] **Step 1: Write the failing test**

```go
func TestRaidPostsRunsTheScheduleThenDrains(t *testing.T) { /* seeded due post (webhook FakePoster) and a due bot reminder row with a FakeREST drainer → one webhook post, the drainer's EditMessage, log line has posted=1 reminded=1 and no webhook URL or token; Drainer nil → schedule only, no panic */ }
```

- [ ] **Step 2: Run the test and check it fails**

Run: `cd api && $TDB go test -p 1 ./cmd/api/ -run RaidPosts -v`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces, then the workflow and README lines**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `$TDB go test -p 1 ./cmd/api/ -run 'DiscordBot|BillingDaily'`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add api/cmd/api/raid_posts.go api/cmd/api/raid_posts_test.go api/cmd/api/main.go api/cmd/api/discordbot.go .github/workflows/api.yml api/README.md
git commit -m "feat(api): the raid-posts job sends scheduled raid posts and reminders every ten minutes"
```

After this task, report this owner action: the Cloud Run job and its scheduler entry are created once by hand, using the README commands. CI only moves an existing job's image.

---

### Task 14: Web: the raid form's `Post to Discord` switch and `Remind the day before`

**Files:**
- Create: `web/src/components/guild/RaidFormDiscord.svelte`, `web/src/components/guild/RaidFormDiscord.test.ts`
- Modify: `web/src/lib/guild/api.ts` (`GuildRaidDiscord.remind_default`, `channel_name: string | null`; `GuildRaidView.discord_post`), `web/src/lib/guild/raids-api.ts` (`RaidInput.remind_day_before: boolean`), `web/src/lib/guild/raid-form.ts` (+ test), `web/src/lib/guild/raid-copy.ts` (+ test), `web/src/components/guild/RaidForm.svelte` (+ test), `web/tests/e2e/guild-raids.spec.ts` (stubs and new cases)

**Interfaces:**
- Consumes: the Task 12 contract (`discord.remind_default`, `discord.channel_name`, `discord_post`, `remind_day_before`).
- Produces:
  - `raidCopy.form` gains:
    - plain strings: `remindDayBefore`, `postOff`, `remind`, `remindNeedsPost`, `reminded`;
    - functions: `postNow(channel: string | null)`, `postLater(channel: string | null, day: string)`, `postRepeat(firstWithinWeek: boolean)`, `posted(channel: string)`.
  - `RaidFormValues.remindDayBefore: boolean`. Create defaults read `discord.remind_default`. Edit defaults read `discord_post?.remind ?? discord.remind_default`.
  - `export type DiscordHelper = { post: string | null; remind: string; postRow: 'switch' | 'posted' | 'webhookMoved'; remindRow: 'live' | 'reminded' }`.
  - `export function discordHelper(args: { values: RaidFormValues; discord: GuildRaidDiscord; post: GuildRaidDiscordPost | null; mode: 'create' | 'edit'; startsAt: Date | null; now: Date; timezone: string }): DiscordHelper`. It implements P.3's choice of helper and row:
    - `postRow`:
      - `posted` for an edit of a bot post (`discord_post.posted_at` set and `discord.channel_name` not null);
      - `webhookMoved` for an edit of a webhook post (`posted_at` set and `channel_name` null), which keeps G1's `Post to Discord if the time changes`;
      - `switch` otherwise.
    - `post` helper:
      - `postOff` when the switch is off;
      - `postRepeat` when repeat is checked;
      - `postNow` within 7 days (168 h);
      - otherwise `postLater` with `{Weekday} {d} {Month}`, which is the start minus 7 days in the guild zone.
    - `remindRow` is `reminded` once `discord_post.reminded_at` is set.
    - `remind` is `remindNeedsPost` with the switch off, else `remind`.
  - `export function channelLabel(name: string | null): string | null` returns `#` plus the name cut to 60 characters followed by `…`.
  - `raidPayload(values, mode, discordConnected)` adds `remind_day_before`.
  - `RaidFormDiscord.svelte` props: `{ values, helper, busy, set }`. It renders:
    - the `CHECK_ROW` input with `role="switch"` and `aria-describedby` set to its 13 px `--muted` helper, 4 px under the row;
    - an 8 px gap;
    - the reminder row, indented 24 px (`pl-6`). It is native `disabled` with `DISABLED_LOOK` when the switch is off, and keeps its value.
    - Both rows are `min-h-11`. The block reserves its height at first paint (`min-h-[120px]` under 640 px, `sm:min-h-[96px]`).
    - Test ids: `raid-form-discord`, `raid-form-discord-helper`, `raid-form-remind`, `raid-form-remind-helper`, `raid-form-discord-posted`.

- [ ] **Step 1: Write the failing tests**

```ts
// raid-copy.test.ts — P.7 verbatim, SPEC2 = design/specs/2026-10-28-discord-guided-cards.md, HEADING '## P.7 Copy, verbatim'
it('form strings match P.7', () => {
  expect(raidCopy.form.remindDayBefore).toBe(str2('`form.remindDayBefore`'));
  expect(raidCopy.form.postNow('#raid-signups')).toBe(str2('`form.postNow`', 0, { channel: '#raid-signups' }));
  expect(raidCopy.form.postNow(null)).toBe(str2('`form.postNow`', 1));
  // postLater both forms, postRepeat both forms, postOff, remind, remindNeedsPost, posted, reminded
});
// raid-form.test.ts
it('helper: within a week posts now, further out names the day, off says so', () => { /* start now+3d → postNow; now+10d (guild America/Chicago) → postLater with "Thursday 3 December"; switch off → postOff and remind == remindNeedsPost; repeat → postRepeat(true|false) */ });
it('edit of a bot post shows the posted line; reminded shows Reminder sent', () => { /* postRow 'posted'; reminded_at set → remindRow 'reminded' */ });
it('channel name is cut at 60', () => { expect(channelLabel('x'.repeat(70))).toBe('#' + 'x'.repeat(60) + '…'); });
it('payload carries remind_day_before', () => { /* raidPayload(...).remind_day_before === true */ });
// RaidFormDiscord.test.ts
it('draws the switch on with its helper and the reminder under it', () => { /* role="switch", checked, aria-describedby → helper text; toggle off → reminder disabled, still checked, helper remindNeedsPost */ });
```

```ts
// guild-raids.spec.ts (API stubbed)
test('Post to Discord is on by default and the reminder under it', /* both checked, role switch, helpers per P.3; axe clean; both rows 44 px at 360 */);
test('turning Post to Discord off disables the reminder', /* … */);
test('the saved values return as the next raid’s defaults', /* the POST body carries post_to_discord false, remind_day_before false; the stubbed next GET returns them; the reopened form shows them */);
test('edit of a posted raid shows the posted line', /* discord_post.posted_at set, channel_name raid-signups → "Posted to #raid-signups. The post updates itself." */);
test('no layout shift as the defaults load', /* CLS 0 over the form at 360 and 1440 */);
test('a posted raid link opens the raid page', /* goto /guild/us/normal/olympus-xxvii/raids/1 (API stubbed) → the raid page heading (P.6) */);
```

- [ ] **Step 2: Run the tests and check they fail**

Run: `cd web && npx vitest run src/lib/guild/raid-copy.test.ts src/lib/guild/raid-form.test.ts src/components/guild/RaidFormDiscord.test.ts src/components/guild/RaidForm.test.ts`
Expected: FAIL.

- [ ] **Step 3: Implement the Interfaces; mount `RaidFormDiscord` in `RaidForm.svelte` in place of today's `discordPart === 'box'` label. The `line` branch is unchanged.**

- [ ] **Step 4: Run the tests and check they pass**

Run: the Step 2 command, then `cd web && npx playwright test tests/e2e/guild-raids.spec.ts tests/e2e/guild-raids-bot-only.spec.ts`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add web/src/components/guild/RaidFormDiscord.svelte web/src/components/guild/RaidFormDiscord.test.ts web/src/lib/guild/api.ts web/src/lib/guild/raids-api.ts web/src/lib/guild/raid-form.ts web/src/lib/guild/raid-form.test.ts web/src/lib/guild/raid-copy.ts web/src/lib/guild/raid-copy.test.ts web/src/components/guild/RaidForm.svelte web/src/components/guild/RaidForm.test.ts web/tests/e2e/guild-raids.spec.ts
git commit -m "feat(web): Post to Discord is a switch, on by default, with Remind the day before"
```

---

### Task 15: Real payloads beside the boards, Discord captures D13–D17, and the 15-minute question

**Files:**
- Modify: `web/tests/captures/discord-bot/payload-dump.go.txt` (frames), `web/tests/captures/discord-bot/payloads.json` (regenerated), `web/tests/captures/discord-bot/capture-messages.spec.ts` (the frame list and the three render additions from the mock: a chosen option, a disabled select, one list open). Also `design/specs/2026-10-27-discord-bot.md` §9 table (rows D13–D17 marked captured).
- Conditional (only if D13 shows expiry): `api/internal/discordbot/copy_guide.go`, `bis_card.go`, `char_card.go`, `signup_card.go`, `officer_cards.go` (footer), and their tests.

**Interfaces:**
- Consumes: every earlier task.
- Produces:
  - Payload frames with the board ids: `bis-card-1-class`, `bis-card-2-spec`, `bis-card-3-range`, `bis-card-4-ready`, `bis-card-nudge`, `bis-card-main`, `bis-card-main-changed`, `bis-card-stale`, `bis-result`, `post-scheduled`, `post-tomorrow`, `post-reminder`, `post-reminder-webhook`, plus `char-card-*`, `signup-card`, `attendance-select`, `readiness-select`, `report-select`.
  - Each frame comes from the real handlers or drainer: `go test -overlay` from a scratch directory, as the file's header says. The file is never added to `api/`.

- [ ] **Step 1: Write the frames into `payload-dump.go.txt` and regenerate `payloads.json`**

Run: the README's overlay command for `TestCaptureDump`. Expected: `payloads.json` holds every id above.

- [ ] **Step 2: Render and compare against the boards**

Run: `cd web && npx playwright test -c tests/captures/discord-bot/playwright.capture.config.ts capture-messages`. Expected: a 640 and a 390 PNG per frame under `design/reviews/2026-10-28-discord-guided-cards/`, each placed beside its mock board. Any difference in words, colour, button style, row order or pre-chosen values is a defect: fix it in the owning task's code, not in the board.

- [ ] **Step 3: Discord captures D13–D17 on the test server**

This needs the deployed build with the commands registered to `DISCORD_COMMANDS_GUILD_ID`, and the owner's or coordinator's Discord session. Capture desktop and phone for each, then save them in `design/reviews/2026-10-28-discord-guided-cards/discord/`:
- **D13:** bare `/bis`, unlinked and linked, every step, the nudge, the result and `Change picks`. Then one pick made 20 minutes after the card was drawn.
- **D14:** `/char` bare (linked and unlinked), the modal, an unknown name.
- **D15:** `/signup` bare with two raids and two characters, then a status press.
- **D16:** `/attendance` officer select, `/readiness` and `/report` with their selects.
- **D17:**
  - a raid saved within 7 days, posting at once;
  - the `Tomorrow` edit;
  - the reply reminder;
  - a webhook reminder;
  - a cancel deleting the reminder.

  For timing, run the `raid-posts` job by hand: `gcloud run jobs execute raid-posts --region us-east1 --wait`. Never echo its environment.

- [ ] **Step 4: Settle the 15-minute question (Ruling 12)**

- If the D13 late pick worked: record "D13: presses accepted at 20 min, no footer" in the bot spec's §9 row. No code changes.
- If Discord drew `This interaction failed`:
  1. Add `CardFooterWindow` (`card.footer.window`, read from §G.7) to `guideCopy`.
  2. Make every card embed's footer that string. For the `/char` and `/signup` cards that have no card embed of their own, append the string as a last content line.
  3. Add `TestCardsCarryTheWindowFooter`.
  4. Run `cd api && $TDB go test -p 1 ./internal/discordbot/`. Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git reset -q && git add web/tests/captures/discord-bot/payload-dump.go.txt web/tests/captures/discord-bot/payloads.json web/tests/captures/discord-bot/capture-messages.spec.ts design/reviews/2026-10-28-discord-guided-cards design/specs/2026-10-27-discord-bot.md
# plus the conditional Go files when Step 4 added the footer
git commit -m "docs(review): guided cards and raid reminders from the real handlers, captures D13-D17"
```

Then the boards go to the ux-designer and wow-player pair (`feedback-design-before-build-two-reviewers`): same-scale side-by-side with each mock board, 640 and 390.

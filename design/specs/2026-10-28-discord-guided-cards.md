# Discord bot amendment: guided cards for bare commands, and raids that post themselves: experience spec

Author: ux-designer, 2026-10-28; amended the same night with §P.9 (the roster on the post) and §P.10 (Which spec tonight?), then at 22:10 with the spec icon as the roster mark (§P.9.1, §P.9.3) and the raid page's spec ask (§P.10.7). Amends `design/specs/2026-10-27-discord-bot.md` (the bot spec): §3.3.2, §3.3.3, §4.2, §6.2, §8 and §9; and `design/specs/2026-10-26-guild-raids.md` (G1) §3.5 item 8. Tenets: `docs/tenets.md` 1, 2, 7, 10, 14.

Source, part G (the owner, 2026-10-28, after using the live bot, verbatim): "instead of making these long option input commands, can we respond with cards that guide the user? For example, if they just type /bis, we should come back with a card to select class spec and range." What the owner saw: `/bis` with no arguments answered only-you with `Name a class and spec, or link Discord to use your character.` and a `Link Discord` button; the option list read class, spec, slot, phase.

Source, part P (the owner the same night, through the coordinator): the raid form's `Post to Discord` checkbox should not be an opt-in; "it should default post to discord 5 days before the scheduled day or however guilds usually operate." The coordinator's ruling is binding and is §P.1.

Source, part P roster (the owner, 2026-10-28 20:10, on the review page): the raid post must show who is signed up, by class and spec, like existing Discord bots. §P.9 and §P.10.

Mock: `design/mocks/2026-10-28-discord-guided-cards/gen_guided_cards.mjs`. It renders through the reviewers' render kit (`web/tests/captures/discord-bot/discord-render.ts`) and extends it, inside the mock only, with three things the kit lacks: a select showing its chosen option (Discord's `default: true`), a disabled select, and one list open per frame. The result frame is the real handler payload `bis-top` from the kit's `payloads.json` plus this amendment's `Change picks`; the post frames start from the real `post-open` and `post-webhook` payloads, with the §P.9 roster drawn by the same rules the builder follows, and the spec-icon emoji drawn from the real icon files at 20 px. Every board is drawn at 640 (desktop client) and 390 (phone client), as `capture-messages.spec.ts` does. Boards, each `-desktop.png` and `-phone.png`:

| Board | What it shows |
|---|---|
| `bis-card-1-class` | Bare `/bis`, not linked: the card, the class list open, spec disabled, range pre-chosen `Phase 1 raid` |
| `bis-card-2-spec` | After the class pick: class colour bar, Hunter crest, the spec list open |
| `bis-card-3-range` | After the spec pick: the range list open |
| `bis-card-4-ready` | Every pick made: `Show BiS` |
| `bis-card-nudge` | `Show BiS` pressed before a spec: the nudge, picks kept |
| `bis-card-main` | Bare `/bis`, linked: the card picked from the main, one tap from the answer |
| `bis-card-main-changed` | Linked, a pick changed away from the main: `Use Zulmara` appears |
| `bis-card-stale` | A press on a card of an older shape |
| `bis-result` | `Show BiS`: the card has become the typed command's answer, in place |
| `post-scheduled` | The raid post the moment the raid is saved: `Signed up` `None yet`, no roster |
| `post-partial` | Partly signed, 8 of 40: the four groups, `Healers · 0` reading `None yet`, Late and Out |
| `post-tomorrow` | The day-before reminder as an edit of the post, 33 of 40 with the roster (the post is still the newest message) |
| `post-reminder` | The post with its roster, and the day-before reminder replying under it (the channel has moved on) |
| `post-reminder-webhook` | The day-before reminder through a webhook: the line, the link, the roster in text lines |
| `post-full` | 40 of 40: Tanks, Healers, Melee, Ranged, Late, Bench, Out |
| `spec-ask-card` | `In` pressed by a Paladin with no spec known: `Which spec tonight?`, the spec list open |
| `spec-ask-done` | The pick: the sign-up outcome in place, `Change spec` under it |

No new token and no new site control (tenet 10). The Discord side uses only Discord's own components (embeds, buttons, string selects, one modal); its look is Discord's, its words are ours.

# Part G. Guided cards

## G.1 The one question, and the answer on the first screen

| Viewer, surface | Question | Answer visible with no scroll, Discord phone client (390) and desktop client (640) |
|---|---|---|
| Anyone, a bare `/bis` | "Best in slot for which class, spec and range?" | The card: its title says what is picked so far; the three selects and `Show BiS` sit under it. On the phone the whole card, buttons included, is about 450 px tall (`bis-card-main-phone`) and fits one 640 px screen. A linked player with a main sees the card already picked from it: one tap to the answer. |
| Anyone, a bare `/char` | "Whose character?" | Linked with a character: the main's card at once, with the other characters one pick away. Otherwise: one sentence and `Type a name`. |
| Raider, a bare `/signup` | "Sign up for which raid, as which character?" | The next raid's embed, pre-picked raid and character, and the four status buttons; a press is the sign-up. |
| Officer, a bare `/attendance`, `/readiness`, `/report` | The command's question, then "and someone or something else?" | The answer as today, first; one select under it for the other scope. |

Tenet 14: a card leads with the player's own character wherever we know it. A card is never a menu of features.

## G.2 References, and what we borrow

- Wowhead's BiS guide filter bar (class, spec and phase dropdowns above the table). Borrowed: three dropdowns in that order, the table appearing under them. Not borrowed: the page reload per change; our card edits in place.
- Our own BiS landing (`design/specs/2026-10-10-bis-landing.md`): class, then spec, then band. Borrowed: the order and the band words (`level 60 Phase 1 raid`).
- The in-game character panel: the character comes first. Borrowed: a linked player's card starts picked from the main and names it.
- Discord's own pattern for ephemeral pickers (the bot spec's character select, §3.3.2 step 6): one ephemeral message, a pick replaces it with interaction type 7. Borrowed unchanged; this amendment generalises it.
- Raid-Helper's day-before reminder (part P): a short line in the channel the day before. Borrowed: the timing and the brevity. Not borrowed: the `@here` ping (the bot never pings, §3.3.0).

## G.3 The rule

1. A command that can take arguments never answers a missing or unusable argument with an error sentence. It answers with an only-you (flag 64) guiding card.
2. A command with several questions (`/bis`: class, spec, range; `/signup`: raid, character, status), or with one question and no default (`/char` for a viewer with no character), answers a bare invocation with a card that asks. The card commits on one button when it asks several questions, and on the pick itself when it asks one.
3. A command with one question whose default is known (`/report`: the last night; `/readiness`: the next raid; `/attendance`: your own record; `/char`: your main) answers at once, as today, and adds one select under the answer to change that one question. The answer is never held back behind a picker.
4. A partial invocation fills the card with what was typed and asks only for the rest. A typed argument that cannot be resolved (an unknown spec, an unknown name) opens the card with that argument's line (§G.7) instead of a sentence alone.
5. A complete invocation answers exactly as today. The arguments remain for players who type them; autocomplete is unchanged.
6. Refusals of permission or membership stay sentences (`Only officers can look up another member's attendance.`, `Readiness is for officers of {Guild}.`): there is nothing to pick.
7. Every card and every answer it becomes is ephemeral. Nothing in this amendment posts to a channel except `Share in channel` (§3.3.4, unchanged).

## G.4 The cards, command by command

### G.4.1 Shared grammar

- Interaction flow. A select pick or a button press on a card is a component interaction; the handler answers with type 7 (update the message) within 3 s, or defers with type 6 and edits in through the press's own token (respond.go `writeDeferral(update)`; the card prefixes join `updateComponents`). Discord draws its own spinner on the pressed component while deferred.
- State lives in the custom id, never on our side: `g{cmd}:{version}:{part}:{state}`. `{version}` is `1` for this amendment. `{part}` names the component (`c` class, `s` spec, `r` range, `go` show, `me` use my character, `chg` change picks, `raid`, `char`, a status word). `{state}` is the card's picks, colon-separated, empty fields kept. The router keys on the prefix before the first colon (`gbis`, `gchar`, `gsgn`, `gatt`, `grdy`, `grep`), as `customPrefix` does today. The longest id, `gbis:1:chg:warrior:feral-bear:p1:main_hand`, is 45 characters; a test asserts every card id under Discord's 100.
- Nothing in a custom id identifies another account or carries a credential (bot spec §9 contract). `/signup` card ids carry the presser's own character key, as the `pick:` select values do today; every press re-checks it against the presser's verified characters. Officer select values in `/attendance` name members of the guild and are re-checked against the presser's rank on every pick.
- A select keeps showing the chosen value only when the re-rendered option carries `default: true`. `SelectOption` gains `Default bool` (`json:"default,omitempty"`); every re-render sets it on the current pick.
- A disabled select still needs one option (Discord requires 1 to 25). It carries a single option whose label is its placeholder and whose value is `none`.
- Buttons follow the bot spec's styles (§3.3.0): the card's one forward action is Success (3, the one coloured button, as `In` is on the post); every other action is Secondary (2); links are Link (5). No Primary, no Danger, no emoji.
- A card's embed, where it has one, is the bot spec's embed grammar: footer `Forever Sixty` with the timestamp; gold `15055189` until a class is known, then the class colour (`--color-class-*` as an integer); thumbnail the class's ringed HD crest (`/icons/hd/crests/{class}.webp`, absolute; one crest language) once a class is known, none before.
- Row budget: Discord allows 5 action rows. The largest card, `/bis`, uses 4. No card uses a fifth.

### G.4.2 `/bis` (mock: the `bis-card-*` boards and `bis-result`)

When: no `spec` typed. That covers a bare `/bis`, `class` only, `phase` only, `slot` only, and any mix without `spec`. A typed `spec` (with or without `class`) answers at once, as today, `phase` defaulting as today.

Layout, top to bottom, one message:
1. Embed. Author `Best in slot`. Title: `Choose a class` (no class); `{Class} · choose a spec` (class, no spec); `{Spec} {Class} · {band label}` (class and spec; `{band label}` from `bandLabels`). No title link: the result links. Colour and thumbnail per §G.4.1. Description: the lead line (§G.7 keys `bis.lead.*`), or the nudge, or a failure line. Footer `Forever Sixty` and the timestamp.
2. Row 1, the class select. Placeholder `Choose a class`. Nine options in the game's order (`classChoices`): label the class, description its specs joined by `, ` in catalogue order (`Balance, Feral, Feral Bear, Restoration`, 38 characters, the longest). Value the class slug.
3. Row 2, the spec select. With no class: disabled, placeholder `Choose a class first`. With a class: placeholder `Choose a spec`; the class's specs from `specs.All` in catalogue order (two to four; Druid has four); label the spec name, description its role word (`DPS`, `Tank`, `Healer`). Value the spec slug.
4. Row 3, the range select. Placeholder `Choose a range`. Six options in this order, label then description: `Level 20` `Levels 20 to 29`; `Level 30` `Levels 30 to 39`; `Level 40` `Levels 40 to 49`; `Level 50` `Levels 50 to 59`; `Level 60` `Level 60, no buffs or consumables`; `Phase 1 raid` `Level 60, raid buffs and consumables`. Values are `phaseChoices`' (`20` to `60`, `p1`). Always pre-chosen: the typed `phase`; else the main's band when the class is the main's (`phaseOfBand`); else `Phase 1 raid`.
5. Row 4, the buttons, in order: `Show BiS` (Success, always live); `Use {Name}` (Secondary; only for a linked viewer with a main, and only when any pick differs from the main's class, spec and band); `Link Discord` (Link; only for a viewer who is not linked).

Picks:
- Class pick: the card re-renders with the class. The spec clears, except when the new class is the main's class, where the main's spec is pre-chosen. The range keeps its value. Lead `Now choose a {Class} spec.` (or the main lead when the main's spec was restored).
- Spec pick: the card re-renders; lead `Ready. Show BiS, or change a pick.`
- Range pick: the card re-renders; the lead is unchanged unless it was the nudge, which gives way to the state's lead.
- `Show BiS` with a spec: the card becomes the answer (§G.5). Without a spec: the nudge replaces the lead (`Choose a spec first. Then Show BiS.`, or with no class `Choose a class and a spec first. Then Show BiS.`, the first sentence in bold), every pick kept, nothing else moves.
- `Use {Name}`: the main is resolved afresh from the presser's link (never from the custom id); the card becomes the answer for the main's spec and band.
- A typed `slot` rides in the card's state. The lead says so (`bis.lead.slot`), and `Show BiS` answers the slot form of the answer (the pick and two alternatives).

Linked viewer with a main: a bare `/bis` opens the card fully picked from the main (class, spec, the main's band) with the lead `Picked from your main, {Name}. Show BiS, or change a pick.` (the name in bold). This replaces today's immediate answer (ruling G2). Linked with no verified character in the guild: the unlinked card without `Link Discord`, lead `bis.lead.none`.

Partial and unusable arguments:
- `class` only: the class pre-chosen; the main's spec pre-chosen when it is the main's class.
- `spec` that does not resolve: the class pre-chosen if `class` was typed and valid, the spec empty, lead `No {Class} spec by that name. Choose one.` (or `No spec by that name. Choose one.` without a class). This replaces `Name a class and spec, or link Discord to use your character.`, which is retired.

### G.4.3 `/char`

`name` becomes optional (§G.8).

- Linked, one or more verified characters in the guild, no `name`: the main's answer at once, exactly `/char name:{main}` (one-question rule, §G.3.3). Row 1 becomes `Share in channel`, `Someone else` (Secondary), `Character page` (Link, last). With two or more verified characters, row 2 is a select: placeholder `Another of your characters`, options the account's verified characters in the guild, main first then the game's sort (the `characterSelect` order and its 25 cap), label the name, description `{Spec} {Class} · Level {n}`, the shown character `default`. A pick replaces the message with that character's answer and keeps the select.
- Not linked, or linked with no verified character: a card with no embed. Content `char.lead.unlinked` or `char.lead.none`. One row: `Type a name` (Success), then `Link Discord` (Link; not linked only).
- `Someone else` and `Type a name` open a modal (interaction response type 9; it is answered at once, never deferred, because it reads nothing). Modal title `Look up a character`. One short text input (component type 4, style 1): label `Character name`, placeholder `Zulmara`, required, min length 2, max length 12. Submitting (interaction type 5, routed by the modal's custom id `gchar:1:name:`) answers with type 7: the answer replaces the card.
- A name that is unknown or private, typed or submitted: the card above, its content `No public character named {typed}. Try another name.` (`{typed}` escaped, at most 12 characters). With a linked viewer who has characters, the card also carries their select.
- Discord modals have no autocomplete. A player who wants suggestions types `/char name:` as today.

### G.4.4 `/signup`

`status` becomes optional (§G.8). With `status` typed the command commits as today.

Without `status`, after the bot spec's member gate (unlinked, not a member, pending: their §3.3.2 sentences; no planned raid: `No raid planned yet.`), the card:
1. Embed: the `/raid` answer's embed for the chosen raid (the post embed with the description replaced by the chosen character's state, bot spec §3.3.3 `/raid`).
2. Row 1, only with two or more planned raids: select, placeholder `Choose a raid`, the guild's planned raids soonest first, at most 25, label `{Title} · {Ddd d Mmm}` (`Onyxia's Lair · Thu 10 Dec`), description `{Guild time} · {signed} of {size} signed up`. Default: the typed `raid`, else the next raid.
3. Row 2, only with two or more verified characters: the bot spec's character select (same order, labels, descriptions, cap), placeholder `Choose a character`. Default: the typed `character`, else the main.
4. Last row: `In` (Success), `Late`, `Bench`, `Out` (Secondary), `Raid page` (Link). Their custom ids are `gsgn:1:{state}:{raid}:{character key}`, so a press needs no further select.

A raid or character pick re-renders the card (type 7). A status press runs the bot spec's flow with raid and character fixed and replaces the card with the outcome sentence (§3.3.2 outcome table, verbatim). A raid that locked, ended or was cancelled since the card was drawn answers its G1 sentence in place. With one raid and one character the card is the embed and the status row alone.

### G.4.5 `/attendance`

A member: unchanged (one embed per own character). An officer with no `name`: the same answer, then one more row, then the existing `Roster` row:
- Select, placeholder `Look up a member`: the guild's members with the lowest attendance over the last 8 raids, lowest first, at most 25; label the character name, description the roster chip verbatim (`Came 3 of 8 · 10 of 30 raids`). Nothing pre-chosen.
- The `Roster` row gains `Type a name` (Secondary) before `Roster`. It opens the modal of §G.4.3 titled `Look up attendance`, label `Character name`.

A pick or a submitted name replaces the message with that member's record; the select stays, the member `default`. A name that is not a member: the officer's own answer stays, the content line `No member of {Guild} named {typed}.` above it. Rank is re-checked on every pick (bot spec product §8.5); a former officer's pick answers `Only officers can look up another member's attendance.`

### G.4.6 `/readiness`

Officers only, as today. Gains an optional `raid` option (autocomplete of planned raids, as `/signup`'s). The answer is unchanged; with two or more planned raids a select follows the `Readiness board` row: placeholder `Another raid`, the planned raids soonest first, at most 25, label `{Title} · {Ddd d Mmm}`, description `{Guild time}`, the shown raid `default`. A pick replaces the answer with that raid's board. Rank, plan and freeze are checked on every pick; the refusals are the bot spec's sentences.

### G.4.7 `/report`

The answer is unchanged (the last done raid, or the typed one). A second row follows the `Share in channel` / `Open the report` row: select, placeholder `Another night`, the guild's last 10 done raids newest first (the autocomplete's list), label `{Title} · {Ddd d Mmm}`, description the G2 night line cut to 100 characters, the shown raid `default`. A pick replaces the message with that night's report; Share is offered again per §3.3.4 for that night (a fresh token; the log's visibility decides); coach lines follow as the coach amendment says.

### G.4.8 `/raid`, `/link`, `/help`

Unchanged. `/raid` takes no argument and its answer already carries its buttons. `/link` and `/help` take none. The `/help` lines stay the §6.2 strings.

## G.5 The in-place edit

1. Every pick and press on a card answers with type 7, so the member sees one message that changes, never a trail of replies.
2. The commit (`Show BiS`, `Use {Name}`, a `/char` pick or name, a `/signup` status, a scope pick) replaces the card with exactly the message the typed command produces for the same resolved arguments: same content line, same embeds, same rows. It is built by the same function (`answerBis`, `answerChar`, `reportMessage`, the press flow), so the two can never drift.
3. One addition on the `/bis` answer, typed or carded: `Change picks` (Secondary) between `Share in channel` and the link button. It reopens the card in place with the answer's class, spec, range and slot. It is never in the shared public follow-up.
4. The Share token is minted when the answer is drawn, so a card that never reaches its answer mints none.
5. The edit body writes content, embeds and components explicitly (`Message.edit`), so a card's selects never linger under an answer.

## G.6 States

| State | What the member sees |
|---|---|
| Not linked | `/bis`: the card with `Link Discord` and `bis.lead.unlinked`; BiS is public, so every pick works. `/char`: `char.lead.unlinked` with `Type a name` and `Link Discord`. `/signup`, `/attendance`, `/report`, `/readiness`: their bot spec unlinked sentences, unchanged (nothing to pick before a link). |
| Linked, no verified character in the guild | `/bis`: the card without `Link Discord`, `bis.lead.none`. `/char`: `char.lead.none` with `Type a name`. `/signup`: the bot spec's member sentences. `/attendance`: `Attendance is for members of {Guild}.` |
| Pick in flight | Discord's spinner on the component; over 2.5 s the press defers (type 6) and the edit follows. |
| `Show BiS` too early | The nudge in the lead (§G.4.2), picks kept. |
| A pick or commit fails | The card stays with every pick; its lead (or content line) becomes `Could not load that. Try again.` This changes `replyOf` for card routes: today a failed update replaces the message with the failure sentence and drops the selects. |
| An old card | Each press on a card is its own interaction with its own fresh 15-minute token, so a card answers for as long as Discord shows it; the slash command's own token is never used after the first answer, and the card's state is in its ids, so nothing on our side expires. An ephemeral card disappears when the member dismisses it or reloads the client. |
| A card of an older shape | A press whose id has an unknown `{version}` replaces the card in place with `This card is out of date. Run /{command} again.` (`bis-card-stale`). |
| Discord stops delivering presses | If the test server shows Discord refusing presses on an ephemeral card after 15 minutes (it draws its own `This interaction failed`), the builder adds `card.footer.window` to every card's footer. Our code cannot answer a press Discord does not deliver, so there is no other handling. Capture D13 decides. |
| Options over 25 | Classes 9; specs 2 to 4; ranges 6; report nights 10: all fit. Planned raids are capped at 12 by G1 (`raid_cap`). Own characters and the officer list cap at 25, the bot spec's `Showing 25 of {n}.` line rule applying to own characters. |
| Unbound server | `This server is not linked to a Forever Sixty guild. An officer can add the bot from the guild's settings.`, unchanged, before any card. |
| Someone else's card | Impossible: every card is ephemeral, so only its owner can press it. |

## G.7 Copy, verbatim

Discord strings live in a new copy struct `guideCopy` (`api/internal/discordbot/copy_guide.go`), keyed as below; the tests read each string back out of this table, as `lookupCopy` does from the bot spec. `{Class}`, `{Spec}`, `{Name}`, `{Guild}`, `{typed}`, `{Slot}`, `{command}` are escaped data. Bold is Discord markdown. No string uses `soon`, an exclamation mark or `unlock`.

| Key | String |
|---|---|
| `bis.author` | `Best in slot` |
| `bis.title.none` | `Choose a class` |
| `bis.title.class` | `{Class} · choose a spec` |
| `bis.title.ready` | `{Spec} {Class} · {band label}` |
| `bis.lead.unlinked` | `Choose a class, a spec and a range, then Show BiS. Link Discord and this starts from your character.` |
| `bis.lead.none` | `Choose a class, a spec and a range, then Show BiS.` |
| `bis.lead.main` | `Picked from your main, {Name}. Show BiS, or change a pick.` (`{Name}` in bold) |
| `bis.lead.class` | `Now choose a {Class} spec.` |
| `bis.lead.ready` | `Ready. Show BiS, or change a pick.` |
| `bis.lead.slot` | `For {Slot} only. Choose a class, a spec and a range, then Show BiS.` |
| `bis.lead.unknownSpec` | `No {Class} spec by that name. Choose one.`; without a class `No spec by that name. Choose one.` |
| `bis.nudge.spec` | `Choose a spec first.` (bold) + ` Then Show BiS.` |
| `bis.nudge.both` | `Choose a class and a spec first.` (bold) + ` Then Show BiS.` |
| `bis.placeholder.class` | `Choose a class` |
| `bis.placeholder.specLocked` | `Choose a class first` |
| `bis.placeholder.spec` | `Choose a spec` |
| `bis.placeholder.range` | `Choose a range` |
| `bis.range.desc` | `Levels 20 to 29`, `Levels 30 to 39`, `Levels 40 to 49`, `Levels 50 to 59`, `Level 60, no buffs or consumables`, `Level 60, raid buffs and consumables` (labels are `phaseChoices`' names) |
| `bis.role` | `DPS`, `Tank`, `Healer` |
| `bis.show` | `Show BiS` |
| `bis.useMain` | `Use {Name}` |
| `bis.change` | `Change picks` |
| `char.lead.unlinked` | `Which character? Type a name, or link Discord to start from yours.` |
| `char.lead.none` | `Which character? Type a name.` |
| `char.unknown` | `No public character named {typed}. Try another name.` |
| `char.someoneElse` | `Someone else` |
| `char.typeName` | `Type a name` |
| `char.select` | `Another of your characters`; option description `{Spec} {Class} · Level {n}` |
| `modal.char.title` | `Look up a character` |
| `modal.attendance.title` | `Look up attendance` |
| `modal.label` | `Character name` |
| `modal.placeholder` | `Zulmara` |
| `signup.placeholder.raid` | `Choose a raid` |
| `signup.raid.option` | label `{Title} · {Ddd d Mmm}`, description `{Guild time} · {signed} of {size} signed up` (unknown size `{n} signed up`) |
| `signup.placeholder.character` | `Choose a character` (the bot spec's `Press: select` placeholder, shared) |
| `attendance.placeholder` | `Look up a member` |
| `attendance.unknown` | `No member of {Guild} named {typed}.` |
| `readiness.placeholder` | `Another raid` |
| `report.placeholder` | `Another night` |
| `card.stale` | `This card is out of date. Run /{command} again.` |
| `card.failure` | `Could not load that. Try again.` (the bot spec's `Any failure`, shared) |
| `card.footer.window` | `Forever Sixty · Picks work for 15 minutes.` (only if capture D13 shows Discord refusing older presses) |
| `spec.title` | `Which spec tonight?` |
| `spec.lead` | `The pick signs **{Name}** up as {State} for {Title}, {Date}, and is kept for the next raid.` |
| `spec.lead.change` | `The pick changes the spec **{Name}** brings to {Title}, {Date}, and is kept for the next raid.` |
| `spec.placeholder` | `Choose a spec` (shared with `bis.placeholder.spec`) |
| `spec.kept` | ` {Spec} is kept for {Name}.` (appended to the outcome sentence; alone, without the leading space) |
| `spec.change` | `Change spec` |
| `spec.option` | The option label is the spec's full catalogue name (`Beast Mastery`, `Retribution`); never a short word (§P.9.3) |
| `spec.stale` | `This card is out of date. Press the button on the raid post again.` |

Retired: `Name a class and spec, or link Discord to use your character.` (`lookupCopy.BisNeedsSpec`).

## G.8 Command registration

Changes to `Commands()`; the commands hash changes, so every scope registers again on deploy (`commandsHash`):

| Command | Option | Was | Becomes |
|---|---|---|---|
| `/char` | `name` | required | optional; description `A character name. Yours if left out.` |
| `/signup` | `status` | required | optional; description `In, Late, Bench or Out. Leave out to pick on a card.` |
| `/bis` | `class` | `The class. Yours if left out.` | `The class. Leave out to pick it.` |
| `/bis` | `spec` | `The spec. Yours if left out.` | `The spec. Leave out to pick it.` |
| `/readiness` | `raid` | absent | optional, autocomplete; description `Which raid. The next one if left out.` (the existing `raid` string) |

Command descriptions and the `/help` lines are unchanged. All option descriptions stay under 100 characters.

# Part P. Raids that post themselves

## P.1 The coordinator's ruling (binding, 2026-10-28)

Posting is on by default when the guild has a Discord channel. The sign-up post goes out as soon as the raid is scheduled: guilds post the week's raids the moment they are set and want sign-ups early. A reminder goes out the day before: 24 hours before the start, in the same channel, and it edits the original when nothing changed. The form's control becomes a switch `Post to Discord` (on) with `Remind the day before` (on) under it; both are remembered per guild as the default for the next raid.

Reading of "edits the original when nothing changed" (this spec's, ruling P3): when the post is still the newest message in the channel, a new message would sit directly under it, so the reminder is an edit of the post. When the channel has moved on, the reminder is a short reply to the post, so it surfaces at the bottom where members read.

## P.2 The question, and the answer on the first screen

| Viewer, surface | Question | Answer, no scroll, 360 x 640 and 1280 x 800 |
|---|---|---|
| Officer, the raid form | "Will this go to Discord, and when?" | `Post to Discord`, on, with the helper `Posts to #raid-signups as soon as you save.` directly under it, above the form's actions. |
| Raider, the channel, the day before | "Is that raid tomorrow, and am I in?" | The post retitled `Tomorrow: Onyxia's Lair · Thursday 10 December`, or a one-line reply with the four buttons. One press answers their own state ephemerally, as on the post. |
| Raid leader, the channel | "Who is in, and what is the comp?" | `Signed up` `**33 of 40**` in the post's first row, then `Tanks`, `Healers`, `Melee` (and on desktop `Ranged`, `Late`, `Bench`) with the spec icon and the name on every line (§P.9). The desktop client shows the groups' headers and first lines on the first screen; the phone shows `Signed up`, Tanks and Healers. |

## P.3 The form (amends G1 §3.5 item 8)

Reference: the in-game Calendar's event creation, where an invite goes out on save with no separate send step. Borrowed: saving is sending.

Mount, order and the flat form grammar are G1's. Item 8 becomes, when the guild has a channel (a bot channel or a webhook; `discord.connected`):
1. `Post to Discord`: the form's existing `CHECK_ROW` checkbox markup (the `Also plan the next 3 weeks` row) with `role="switch"` on the input, so assistive technology announces on and off. The site has no switch control and this amendment invents none (tenet 10); the coordinator's "switch" is its role, not a new look. Default: the guild's `discord.post_default`, which starts `true`.
2. Its helper, 13 px `--muted`, `aria-describedby` from the input, 4 px under the row: `form.postNow` (the raid starts within 7 days), `form.postLater` (further out), `form.postRepeat` (the repeat box is checked), or with the switch off `form.postOff`.
3. 8 px gap, `Remind the day before`: the same markup, indented 24 px so its box aligns with the first label's text. Default: the guild's `discord.remind_default`, which starts `true`. With `Post to Discord` off it is native `disabled` in the shared `DISABLED_LOOK`, keeps its value, and is described by `form.remindNeedsPost`. On, its helper is `form.remind`.
4. The G1 actions row.

Without a channel: G1's `Connect a Discord channel in Settings to post raids there.` line, unchanged. With the deployment's connect off: nothing, unchanged.

Edit mode:
- Not yet posted (the switch was off, or the post is scheduled for later): both rows as on create, with their current values.
- Posted through the bot: the `Post to Discord` row is replaced by one 13 px `--muted` line `form.posted`; `Remind the day before` stays live until the reminder has gone, then reads `form.reminded` in its place.
- Posted through a webhook: G1's `Post to Discord if the time changes` checkbox, unchanged (a webhook post is never edited); `Remind the day before` as above.

Remembered: on a successful save, the two values become the guild's defaults (`post_default`, `remind_default`) for the next raid, as G1 remembers the box today.

States: the rows are 44 px tall at every width (the label is the hit area); focus-visible is the checkbox's gold ring; disabled is `DISABLED_LOOK`; while saving both are `disabled` with the form; errors are the form's own status line. Keyboard: Space toggles; tab order `Post to Discord`, `Remind the day before`, then the actions. Phone: one column, helpers wrap, nothing collapses. Height: the block is at most 120 px at 360 (two rows, two two-line helpers) and 96 px at 1440; the form reserves it on first paint, so loading the guild's defaults shifts nothing.

## P.4 When the bot posts

| Event | Bot channel | Webhook channel |
|---|---|---|
| A raid is saved with `Post to Discord` on and it starts within 7 days (168 h) | The post goes out at once (`post-scheduled`, `Signed up` reads `None yet`) | The webhook post goes out at once |
| Saved further out | The post is scheduled for 7 days before the start (`post_at`) and goes out then | Same |
| A repeat (`Also plan the next 3 weeks`) | The first raid as above; each later raid at 7 days before its start, so the channel holds one week of raids at a time | Same |
| 24 h before the start, `Remind the day before` on, the raid planned and not locked | The reminder (§P.5) | The webhook reminder (§P.5) |
| The post went out less than 12 h before the reminder time (a raid scheduled late) | No reminder: the post itself is fresh | Same |
| The raid is moved after the reminder | The reminder message, if one was posted, is deleted; the post reads `Moved:` (§3.3.1); a new reminder is scheduled 24 h before the new start | The old reminder is deleted (webhook message delete); a new one is scheduled |
| The raid is cancelled or locked before the start | The reminder message is deleted; the post shows its state (§3.3.1) | The reminder is deleted |
| `Post to Discord` turned on in edit for an unposted raid | As on save | As on save |

Ruling P2 explains the 7-day horizon. The owner's "5 days before" sits inside it: a raid scheduled at least 5 days ahead is always up at least 5 days ahead.

## P.5 The reminder

Bot channel, the post still the newest message (the bot reads one message of history; Read Message History is already one of its four permissions, bot spec §3.1): the post itself is edited into its `Tomorrow` state (`post-tomorrow`). Title `post.tomorrowTitle`; description `post.tomorrowDescription`; everything else is the open post. It stays in that state until the start, then follows §3.3.1 (ended, done).

Bot channel, the channel moved on: the post takes the same `Tomorrow` state, and a new message replies to it (`post-reminder`): a `message_reference` to the post with `fail_if_not_exists` false, and the empty mention list, so the reply pings no one. Content `reminder.line`, the clause before the first comma in bold. One row, the post's own five components (`In`, `Late`, `Bench`, `Out` with the post's `raid:` ids, then `Raid page`), so a press there is a press on the post. The reminder is never edited; §P.4 says when it is deleted.

Webhook channel: a new message, `reminder.line`, a space, then `reminder.webhookLink`, a masked link to the raid page; under it the webhook post's embed in the Tomorrow state with the roster in text lines (§P.9.7); no components, because a webhook cannot carry buttons (`post-reminder-webhook`).

The counts in `reminder.line` are the counts when it was sent; the living counts are on the post, one tap away.

## P.6 The raid link

Every link the bot or a webhook posts to a raid (embed title, `Raid page`, `Add a note`, the reminder's masked link) is the public base URL followed by `{guildPath}/raids/{id}`, where `{guildPath}` is `/guild/{region}/{ruleset}/{slug}` as `GuildPagePathOf` returns it, never rebuilt by hand. It is the path the site serves: `guildRaidHref` builds it and `parseGuildRaidPath` (`web/src/lib/characters.ts`) accepts exactly six segments with a digits-only id. An API test asserts every posted raid URL passes that rule (a Go mirror of `parseGuildRaidPath`), and the site's e2e suite opens the example `https://foreversixty.gg/guild/us/normal/olympus-xxvii/raids/1` to the raid page.

## P.7 Copy, verbatim

Site strings join `raidCopy.form` (beside `postToDiscord`); Discord strings join the API's post copy. `{channel}` is `#raid-signups` (escaped, at most 60 characters, then `…`).

| Key | String |
|---|---|
| `form.postToDiscord` | `Post to Discord` (unchanged) |
| `form.remindDayBefore` | `Remind the day before` |
| `form.postNow` | `Posts to {channel} as soon as you save.`; through a webhook `Posts to your Discord channel as soon as you save.` |
| `form.postLater` | `Posts to {channel} on {Weekday} {d} {Month}, a week before the raid.`; through a webhook `Posts to your Discord channel on {Weekday} {d} {Month}, a week before the raid.` |
| `form.postRepeat` | `Each raid posts a week before it starts; the first posts as soon as you save.` when the first raid is within 7 days, else `Each raid posts a week before it starts.` |
| `form.postOff` | `This raid will not post to Discord.` |
| `form.remind` | `Posts a reminder 24 hours before the start.` |
| `form.remindNeedsPost` | `Turn on Post to Discord to send a reminder.` |
| `form.posted` | `Posted to {channel}. The post updates itself.` |
| `form.reminded` | `Reminder sent.` |
| `post.tomorrowTitle` | `Tomorrow: {Title} · {Weekday} {d} {Month}` |
| `post.tomorrowDescription` | `The raid is tomorrow. Sign up with the buttons below.` |
| `reminder.line` | `{Title} is tomorrow, <t:{unix}:F> (<t:{unix}:R>). {signed} of {size} signed up so far.` Unknown size: `{n} signed up so far.` Nobody: `Nobody has signed up yet.` |
| `reminder.webhookLink` | `[Sign up on Forever Sixty]({raid url})` |
| `roster.signed` | `**{signed} of {size}**`; unknown size `**{n} signed up**`; nobody `None yet` (the `Signed up` field's value; the role clauses and the `Bench {n} · Out {n}` line leave the post) |
| `roster.tanks` | `Tanks · {n}` |
| `roster.healers` | `Healers · {n}` |
| `roster.melee` | `Melee · {n}` |
| `roster.ranged` | `Ranged · {n}` |
| `roster.late` | `Late · {n}` |
| `roster.bench` | `Bench · {n}` |
| `roster.out` | `Out · {n}` |
| `roster.continued` | U+200B (a continuation field's name) |
| `roster.empty` | `None yet` |
| `roster.line` | `{emoji} {Name}`; no spec `{Name} · {Class}`; no class `{Name}` |
| `roster.textLine` | `{Name} · {Spec} {Class}` (the full spec name and the class, `Aurelith · Holy Paladin`); no spec `{Name} · {Class}` |
| `roster.more` | `and {n} more` |
| `roster.outJoin` | `, ` |

## P.8 Contract

- The raid view's `discord` object gains `remind_default` (boolean) beside `post_default`, and `channel_name` (null for a webhook).
- The raid view gains `discord_post`: null, or an object with `posted_at`, `post_at`, `reminded_at` (each RFC 3339 or null) and `remind` (boolean).
- `POST …/raids` and `PATCH …/raids/{raid}` gain `remind_day_before` (boolean). A save stores it and `post_to_discord` as the guild's defaults.
- The bot store gains, per raid, `post_at`, `remind`, `reminder_message_id` and `reminded_at`. The drain sends the post at `post_at` and the reminder at the start minus 24 h, through the same per-raid worker as edits (bot spec §3.3.1).
- `guild_characters` gains `raid_spec` (text, null), migration `0058_raid_spec`. Written only by the spec card and `Change spec` (§P.10), after the key is re-checked against the presser's verified characters. The roster read prefers it over the newest logged fight's spec.
- The post builder reads, per raid, every sign-up row's name, class, spec, role and state (all already on `guild_raid_signups`); `RaidPosted` gains `Roster []RosterLine` beside `Counts`.
- The bot holds the 28 application emoji ids in memory (§P.9.1); no table stores them.

## P.9 The roster on the post

Source (the owner, 2026-10-28 20:10, on the review page's raid-post boards, verbatim): "we should show what classes and specs are signed up like existing discord bots". Then: "how do we know what someone is if they're not on the site yet?" The coordinator's answer is binding and is encoded here and in §P.10: sign-ups need a linked Discord account (today's `pressCopy.NotLinked` answer and its `Link Discord` button stay); a linked member's character gives the class (the guild roster export); the spec is the character's saved spec, else asked once (§P.10). Guest sign-ups for unlinked people are out of scope until the owner asks.

Reference: Raid-Helper's sign-up embed, the bot raid leaders already read. Borrowed: one column per role group, each headed with its count, one line per player with the spec icon and the name, Late, Bench and Absence listed under the columns. Not borrowed: a row of several emoji per player, the numbered sign-up order, the pings, and the "Absence" word (ours is `Out`, the button's word). Also borrowed, from the in-game raid frames: grouping a raid by what each player does, not by when they pressed.

### P.9.1 The mark: the spec's own icon as an application emoji

Source (the owner, 2026-10-28 22:10, on the roster boards, verbatim): "we should just show the spec icon beside the name like existing discord tools." This replaces the class crest as the mark. The line is the spec icon, then the name; no class crest, no short spec word.

Decision: each line starts with the spec's own talent-tab icon, uploaded once as an application-owned custom emoji. Reference: Raid-Helper and the in-game talent frame, where a spec is its icon (Starfall for Balance, the Bear Form icon for Feral).

- Why not the crest: a crest says the class and the line then needs a spec word; the spec icon says both, because every spec's icon is its own and the spec belongs to one class. Why not Unicode: there is no glyph per spec. Why not the word alone: it wraps lines in a 3-column embed.
- Why application emoji, not server emoji: an application emoji works in every server the bot posts in, needs no Manage Expressions permission, and uses none of a server's emoji slots. An application may own 2000; this uses 28.
- Source of the images: `data/curated/specs.json`, each spec's `icon` field (`spell_nature_starfall` for Balance). The file is `web/public/data/{build}/icons/{icon}.webp`, where `{build}` is the active Forever build's folder (today `1.60.1.70658`; every spec's icon exists there, and the older `forever-prebeta` folder lacks seven, so it is never the source). The files are 64 x 64, the game's own icons; the images are not redrawn or upscaled.
- Count: one emoji per row of `specs.json`, **28**: Druid 4 (Balance, Feral, Feral Bear, Restoration), Hunter 3, Mage 3, Paladin 3, Priest 3, Rogue 3, Shaman 3, Warlock 3, Warrior 3. Forever's new combos are counted in: Survival hunters, Enhancement and Elemental shamans with the Forever rules, and Feral Bear as its own spec. Feral and Feral Bear share one icon (`ability_racial_bearform`) and so share one image under two names; the owner may want a distinct mark (see the questions).
- Names: `fs_{class}_{spec}` from the row's `spec` slug with hyphens turned to underscores: `fs_druid_balance`, `fs_druid_feral`, `fs_druid_feral_bear`, `fs_druid_restoration`, `fs_hunter_beast_mastery`, `fs_hunter_marksmanship`, `fs_hunter_survival`, and so on for every row. The longest is `fs_hunter_beast_mastery` (23 characters, under Discord's 32). Markup in a field value: `<:fs_hunter_survival:{emoji id}>`. Discord draws it at the line's height (about 20 px in an embed field).
- Images: 28 PNGs, 64 x 64, transparent, converted from the webp by `make discord-emoji` into `api/internal/discordbot/emoji/{class}_{spec}.png`, committed and embedded with `go:embed`. Each stays under Discord's 256 KiB.
- Setup step: the subcommand `api discord-emoji` beside `stripe-setup`. It reads `GET /applications/{application id}/emojis`, uploads each missing `fs_{class}_{spec}` with `POST /applications/{application id}/emojis` (name and a `data:image/png;base64,...` image), and prints each name and id, never the token. It is idempotent: a second run uploads nothing. `--replace` deletes and re-uploads all 28 after an icon changes or a build adds a spec. It runs once per application: production, then the test application. A build that adds a spec adds its row to `specs.json`, and the next run uploads only that one.
- The bot at start reads the same list once and holds `fs_{class}_{spec}` to id in memory. While any of the 28 is missing (a failed read, or setup not yet run), it re-reads at most once an hour, logs `emoji_missing` with the names, and draws the text line for those specs.
- Text line (an emoji missing, or any webhook message, §P.9.7): `{Name} · {Spec} {Class}`, the full spec name and the class word (owner ruling 2026-10-28: Holy, Protection, Restoration and Frost each belong to two classes), no mark.

### P.9.2 Layout, field by field

The post embed's fields, in order. "Inline" fields sit three to a row on the desktop client, or two to a row when the embed has the guild crest thumbnail; each group has its own header, so the reading holds either way.

1. `When`, inline, unchanged.
2. `Guild time`, inline, unchanged (absent without a guild zone).
3. `Signed up`, inline (it was full width): only the bold count, `roster.signed` (`**33 of 40**`). The role clauses and the second `Bench 3 · Out 4` line leave the post; each group's header now carries its count.
4. The four groups, inline, in this order: `Tanks · {n}`, `Healers · {n}`, `Melee · {n}`, `Ranged · {n}`. They count and list In rows only. All four are drawn as soon as one person is In or Late; a group with nobody reads `None yet`, so a missing role is the first thing a raid leader sees (`post-partial`: `Healers · 0`).
5. `Late · {n}`, inline, only when someone is Late: lines as in a group.
6. `Bench · {n}`, inline, only when someone is on the Bench: lines as in a group.
7. `Out · {n}`, full width, only when someone is Out: names only, A to Z, joined `, `.

With nobody In or Late, fields 4 to 6 are absent and `Signed up` reads `None yet` (`post-scheduled`). Field 7 still shows any Out.

Desktop reading (`post-tomorrow-desktop`, `post-full-desktop`): row 1 is When, Guild time, Signed up; row 2 is Tanks, Healers, Melee; row 3 is Ranged, Late, Bench; Out under them. Phone reading: the kit stacks inline fields under 500 px, so the phone shows Signed up, then each group in order; capture D18 records what the phone client really does, and the spec holds either way.

### P.9.3 The line

`{emoji} {Name}`, one per line, nothing else. `{Name}` is the sign-up row's name through `memberText`. The emoji is the row's spec icon (§P.9.1). There is no short spec word anywhere: the icon says the spec in the post, and the text line writes the full spec name.

No spec on the row (a sign-up made before the spec ask existed, or by an officer): no mark exists, so the line is the text line with the class, `{Name} · {Class}`. No class on the row: `{Name}` alone, sorted after every class. The text line (§P.9.1) is `{Name} · {Spec} {Class}`, or `{Name} · {Class}` with no spec. In the emoji line the class word appears only when the spec is unknown.

Sort within a group, Late and Bench: class in the game's order (`classChoices`: Druid, Hunter, Mage, Paladin, Priest, Rogue, Shaman, Warlock, Warrior), then name, case-insensitive. A class's specs therefore sit together by name, and a raid leader counts five mages at a glance. Out sorts by name alone, because it has no mark to group by.

A long name may wrap inside a 3-column field on the desktop client; the icon stays at the start of the line. The icon is never the only carrier of the spec for a screen reader: Discord reads the emoji's name (`fs_druid_balance`), and the post's text fallback and the raid page carry the words.

### P.9.4 The groups

- Tanks: role `tank`. Healers: role `healer`.
- Melee and Ranged split role `dps` by spec. Melee specs: Feral, Survival (Survival fights in melee in Forever), Retribution, Assassination, Combat, Subtlety, Enhancement, Arms, Fury. Every other DPS spec is Ranged.
- A DPS row with no spec goes by class: Paladin, Rogue and Warrior are Melee; every other class is Ranged.
- The role is the sign-up row's (`guild_raid_signups.role`), so an officer's role change on the site moves the line.

### P.9.5 Limits and overflow

- A field holds at most 16 lines. The longest line is 60 characters (46 for the emoji markup, `<:fs_hunter_beast_mastery:` plus a 19-digit id plus `>`, a space, 12 for a WoW name, which has no markdown specials to escape, a newline), so 16 lines are at most 960 characters, under the 1024 field limit. The text line is at most 37 characters (12, ` · `, 13 for `Beast Mastery` or `Assassination`, a space, 7 for `Warlock` or `Paladin`, a newline).
- A group over 16 lines continues in one more inline field directly after it, named U+200B (zero-width space; Discord refuses an empty name), so it reads as the same list. A group over 32 lines draws 31 and then `and {n} more`.
- Late and Bench draw at most 10 lines; the tenth becomes `and {n} more`.
- Out: names joined by `, ` until the next name would pass 1024 characters; then `, and {n} more`.
- Field count: at most 3 + 8 + 3 = 14, under Discord's 25.
- Embed total: Discord allows 6000 characters across title, description, field names and values, and footer. A 40-player roster with 10 Bench and a full Out list is about 3,900. A guard measures the embed and, while it is over 5,900, turns the tail of the longest list (Out first, then Bench, then the largest group) into `and {n} more`. A test builds 120 sign-ups and asserts the guard holds.

### P.9.6 States

| Post state | Roster |
|---|---|
| Open, just posted | `Signed up` `None yet`, no groups (`post-scheduled`) |
| Open, partly signed | All four groups, empty ones `None yet` (`post-partial`) |
| Tomorrow | The roster as Open (`post-tomorrow`) |
| Full | `**40 of 40**`; no other change; more In than the size still lists everyone (`post-full`) |
| Locked, Moved, Ended | The roster stays: a locked post is the final comp |
| Cancelled | No roster: the struck `When` and `Signed up` count only |
| Done | Unchanged: `Came` and `Kills`, no roster |

The post already edits itself on every sign-up change through the per-raid worker (bot spec §3.3.1), which coalesces bursts; the roster rides those edits and adds none.

### P.9.7 Webhooks and the reminder

- A webhook message cannot draw an application emoji, so every webhook roster uses the text line.
- The webhook post (sent once, never edited) carries the roster as it stands when it posts.
- The webhook reminder carries, under its `reminder.line` and `reminder.webhookLink` content, the webhook post's embed in the Tomorrow state: title `post.tomorrowTitle`, description the webhook's `Sign up on Forever Sixty.`, the roster as it stands, no components (`post-reminder-webhook`). For a webhook guild this is the only up-to-date roster in Discord.
- The bot's reminder reply is unchanged: one line and the post's row. The post it replies to carries the living roster, one tap up through the reply chip (`post-reminder` shows both).

## P.10 Which spec tonight?

The question: "As what are you coming?", asked once, only when we do not know.

### P.10.1 Where the spec comes from

1. The character's saved raid spec: new column `guild_characters.raid_spec` (text, null when unset; the spec's catalogue name, as `fight_metrics.spec` stores it). Only this card, its `Change spec` button and the raid page's spec select (§P.10.7) write it. It is per guild membership, because a character can heal for one guild and tank for another.
2. Else the newest logged fight's spec (today's `raidRoster` read).
3. Else unknown.

When `raid_spec` is set, it wins over the logs (the player said so), and the sign-up's role is `roleForSpec(raid_spec)`. Otherwise the role is today's `defaultRole`. `roleForSpec` must read `Feral Bear` as a tank. The class always comes from the roster read (the newest logged fight, else `characters.class`, which the guild roster export and Battle.net import fill).

### P.10.2 Where it sits in the press flow

In `signupFlow` (bot spec §3.3.2), after the raid's refusals and after the character is known (the one verified character, the character select's pick, the `/signup` card's key), and before `signUp`:

- The state is In, Late or Bench, the class is known, and the spec is unknown: answer the spec card instead of signing up. It is a new ephemeral message (type 4) on a press of the post or the reminder, and an in-place update (type 7) after the character select or on the `/signup` card.
- The state is Out, or the class is unknown, or the spec is known: sign up as today, with no card.
- An unlinked presser never reaches this step: `pressCopy.NotLinked` and `Link Discord`, unchanged.

So a member's first In costs two taps (the button, the spec); every press after it costs one.

### P.10.3 The card (`spec-ask-card`)

Part G's grammar (§G.4.1) applies.

- Embed: no author. Title `spec.title`. Description `spec.lead` (`spec.lead.change` when opened from `Change spec`). Colour the class colour (`--color-class-*` as an integer). Thumbnail the class's ringed HD crest. Footer `Forever Sixty` and the timestamp.
- Row 1, the only row: one string select. Placeholder `spec.placeholder`. Options: the class's specs from `specs.All` in catalogue order (two to four); label the full spec name (`spec.option`), description the role word (`bis.role`); value the spec slug. With a saved spec (opened from `Change spec`), that option carries `default: true`.
- No buttons. One question with no default commits on the pick (§G.3 rule 2), so there is no Success button.
- Custom id: `gspc:1:ask:{raid}:{state}:{character key}` (`gspc:1:ask:1:in:us/normal/aurelith`, 34 characters; a test asserts every id under 100).

### P.10.4 The pick (`spec-ask-done`)

1. Re-check, in order: the raid's refusals (the bot spec's sentences, in place); the character key is one of the presser's verified characters (else `pressCopy.Failure` in place); the value is one of the class's specs (else `pressCopy.Failure`).
2. Save `raid_spec`.
3. Run `signUp` for the state in the id, with the new spec's role.
4. Answer type 7: the §3.3.2 outcome sentence, then `spec.kept` (`Aurelith is In as Healer for Onyxia's Lair, Thu 10 Dec. 34 of 40 signed up. Holy is kept for Aurelith.`). The outcome's row gains `Change spec` (Secondary) before any `Add a note` link.
5. The post's next edit draws the line with the new spec.

### P.10.5 Change spec

- `Change spec` (Secondary, custom id `gspc:1:chg:{raid}:{character key}`) sits on every In, Late and Bench outcome, and on their `Already` sentences. It is how a healer comes as Retribution tonight.
- A press opens the card in place (type 7) with `spec.lead.change` and the saved spec `default`.
- A pick saves `raid_spec` and updates the role of the character's current row for that raid when it is In, Late or Bench. It never changes the state, so a member who went Out since is not put back In. The answer is the §3.3.2 `Already` sentence for the row's current state, then `spec.kept`. With no row, or an Out row, the answer is `spec.kept` alone, without its leading space.

### P.10.6 States

| State | What the member sees |
|---|---|
| Pick in flight | Discord's spinner on the select; over 2.5 s the press defers (type 6) and the edit follows |
| Save or sign-up fails | The card stays with its select; the description becomes `pressCopy.Failure` (`That did not save. Try again.`) |
| Raid locked, cancelled or ended since the card | The bot spec's sentence in place |
| Character no longer verified | `pressCopy.WaitingOne` in place |
| A card of an older shape (unknown `{version}`) | `spec.stale` in place |
| Rate limited | `pressCopy.RateLimited` in place; nothing saved |
| Keyboard and touch | Discord's own select: arrow keys and Enter, or a tap; nothing of ours |

### P.10.7 The raid page asks too (`raid-signup-spec`)

Source (the owner, 2026-10-28 22:10): the raid page's own sign-up asks the spec as well, so every sign-up has one, whichever door it came through.

Reference: the same question as the Discord card, on the surface the member is already on. Component: the raid page's existing sign-up control, with one native select added under the status buttons; no new control (tenet 10).

| Part | Spec |
|---|---|
| Where | The raid page sign-up, for In, Late and Bench, after the character is known; shown only when `raid_spec` is unknown for that character in this guild. Out never asks. |
| Control | One select: label `spec.title` (`Which spec tonight?`), placeholder `spec.placeholder` (`Choose a spec`), options the class's specs by full name (`spec.option`) in catalogue order. A native select cannot draw an icon per option, so the chosen spec's own icon (`/data/{build}/icons/{icon}.webp`, 20 px, in a reserved box, no missing-icon flash) sits beside the control. With a saved spec the control collapses to a 44 px text button `{Spec} · Change spec` (`aria-expanded`) that reveals the select (the saved spec chosen) and the `spec.lead.change` line; a pick collapses it and returns focus to the button. |
| Copy | The Discord card's keys, verbatim: `spec.title`, `spec.lead` (`spec.lead.change` with a saved spec), `spec.placeholder`, `spec.kept`, `spec.change`. No new strings. |
| Save | One write path with the Discord card: `guild_characters.raid_spec` after the character is re-checked as the member's, then the sign-up with that spec's role. The pick commits the sign-up, as on the card; there is no separate save button. |
| Pending press | In, Late or Bench pressed with no saved spec saves nothing: the segment reads pending (dashed outline, `aria-busy`, not checked) until the pick. Escape, or 60 s out of focus, drops it; a reload shows the row unsigned. An unsigned row reserves the block's height (116 px at 360, 100 px from md, 52 px with a saved spec) so the press shifts nothing below it. |
| States | Default: closed select. Hover, focus-visible and active as the site's select. Disabled while saving. Error: `That did not save. Try again.` under the select; the sign-up is not made. Keyboard: native select, arrow keys and Enter; touch: the native picker. 44 px tall at 360. |
| Unlinked members | Unchanged: the link nudge (`pressCopy.NotLinked` and `Link Discord`); no spec ask until linked. |
| Roster | The next edit of the post draws the line with the new spec's icon. |

Acceptance: a site e2e signs up with no saved spec and sees the select; picking saves `raid_spec`, returns the outcome with `spec.kept`; the next sign-up on either surface asks nothing. The implementer attaches the raid page's sign-up at 360 and 1440 with the select closed, open and with a saved spec. This row's board is the Discord card (`spec-ask-card`); the site control follows the site's select and needs its own captures before merge.

# Rulings

G1. One rule for every command (§G.3): several questions or no default, a card that asks; one question with a known default, the answer first with one select. It keeps answers instant where they were, and removes every "name an argument" sentence.

G2. A linked player's bare `/bis` opens the card picked from the main instead of answering at once. The owner asked for a card on a bare `/bis`; picked from the main it costs one tap and shows the player which class, spec and range the answer is for before it lands. Rejected: an instant answer with three selects under it, because changing the class would leave a stale answer above a half-made pick.

G3. All three `/bis` selects are drawn from the first frame, the spec disabled until a class. The card never grows between steps, and the player sees every question at once (Wowhead's filter bar). Rejected: revealing one select per step, which shifts the buttons under the thumb on the phone.

G4. `Show BiS` is never disabled. A disabled button hides why (bot spec §4.1, `Save channel`); the nudge says what is missing and keeps the picks.

G5. The card asks class, spec and range only. The slot stays a typed option and the card's answer is the five-slot answer, whose `All 17 slots` link opens every slot. A fourth select would fill the card's fifth row for a question few ask in Discord.

G6. `Use {Name}` appears only when a pick differs from the main. Drawn always, it would duplicate `Show BiS` on the main's card.

G7. One Success button per card, its forward action. It extends ruling 3 of the bot spec (one coloured button per message) to cards.

G8. `/char` and `/attendance` take a typed name through a Discord modal. A select cannot hold a guild's public roster (over 25), and a modal is Discord's own text field. It has no autocomplete; the typed option keeps it.

G9. The officer list in `/attendance` is the lowest attendance first, because that is who an officer looks up.

G10. `Change picks` on every `/bis` answer, typed or carded, never on the public share. A player who typed the wrong range gets the card back without retyping.

P1. Posting defaults on, reminder on (the coordinator's ruling).

P2. A raid posts at once when it starts within 7 days, else 7 days before it starts. This honours "as soon as the raid is scheduled" for the week being planned, and stops a 4-week repeat from putting four posts in the channel at once; the owner's "5 days before" sits inside the window.

P3. "Edits the original when nothing changed" reads as: the post is still the newest message, so the reminder edits it into `Tomorrow`; otherwise a reply. This is the spec's reading of the coordinator's words; if the coordinator meant "when the raid itself did not change", the only difference is that a moved raid's reminder would always be a new message, and §P.4 already deletes and reschedules it.

P4. The reminder carries the post's own buttons and is deleted on cancel, lock or move, so no stale live button ever sits in the channel. It never pings.

P5. The switch is the existing checkbox with `role="switch"`. No new control (tenet 10).

P6. The mark is the spec's own talent-tab icon as an application emoji (§P.9.1), by the owner's direction at 22:10 ("like existing discord tools"). It says class and spec in one 20 px image, fits a 3-column field, needs no server permission, and degrades to `{Name} · {Spec} {Class}`. Rejected: the class crest plus a short spec word, the first draw, which doubled every line.

P7. Four groups, not three: DPS splits into Melee and Ranged by spec. The owner asked for the roster "like existing discord bots", and Raid-Helper's classic template splits them the same way, because a raid leader places melee and ranged differently. The first draw of a single DPS column needed four columns at 20 DPS, and the continuation under `Tanks` read as tanks. Rejected: one DPS column with continuations.

P8. Counts move into each group's header; `Signed up` keeps only the bold total and joins the first row beside `When` and `Guild time`. The old clause line would repeat every header.

P9. Late and Bench are lists with marks, beside the groups, because a raid leader fills from them. Out is names only, A to Z, full width, because nobody plans around who is out.

P10. The spec is asked once, on the first In, Late or Bench press whose spec is unknown, and kept per guild membership (§P.10). A card per press would tax every raid; never asking would leave site-only and log-less members as bare names.

P11. The webhook reminder carries the roster in text lines, because a webhook post is never edited and has no living roster to point at.

P12. Unlinked members cannot sign up from Discord (the coordinator's answer). Guest sign-ups wait for the owner.

P13. The raid page asks the spec too (§P.10.7): the same select and copy, saved to `guild_characters.raid_spec`, so every sign-up has a spec and every roster line an icon.

# Acceptance

## Handler tests (`api/internal/discordbot`, each string read back from §G.7 and §P.7)

1. Card routing: each `g{cmd}` prefix routes; each card prefix is in `updateComponents`; an unknown `{version}` answers `card.stale` in place (type 7).
2. `/bis` bare: unlinked draws title, lead, four rows, the spec select disabled with one `none` option, range `default` on `p1`, `Show BiS` Success, `Link Discord` Link; linked with a main draws class, spec and band pre-chosen with `default: true` and no `Use {Name}`; linked with no character draws `bis.lead.none` and no link.
3. `/bis` picks: class clears the spec (restores the main's on the main's class) and keeps the range; spec gives `bis.lead.ready`; `Show BiS` without a spec gives the nudge with every pick kept; with a spec the message equals the typed `/bis spec:… phase:…` answer byte for byte plus `Change picks`; `Use {Name}` answers the main's spec and band; `Change picks` reopens the card with the answer's picks.
4. `/bis` partials: `class` only, `phase` only, `slot` only (lead `bis.lead.slot`, result the slot form), an unknown spec (`bis.lead.unknownSpec` both forms); a typed `spec` still answers at once.
5. `/char`: bare linked answers the main with `Someone else` and the own-characters select (25 cap); bare unlinked draws `char.lead.unlinked`; `Type a name` answers type 9 with the modal's title, label, placeholder and 2 to 12 length; a modal submit answers type 7; an unknown name draws `char.unknown`.
6. `/signup` bare: raid select only with 2 or more planned raids, character select only with 2 or more verified characters; a status press saves with no further select and answers the §3.3.2 outcome in place; a forged character key in a `gsgn` id is refused with `That did not save. Try again.`; a raid locked since the card answers the lock sentence.
7. `/attendance` officer: the select lists the lowest attendance first, at most 25, with the chip descriptions; a pick answers that member; a demoted officer's pick answers the officers-only sentence; a member's answer has no select.
8. `/readiness` and `/report`: the select appears only with 2 or more raids; a pick answers that raid; Share on a report pick follows the log's visibility.
9. Failures: a handler error on a card press keeps the card and sets the lead to `card.failure`.
10. Limits: every card id at most 100 characters; at most 5 rows; every select 1 to 25 options; labels, descriptions and placeholders inside Discord's 100 and 150; no card id or value names another account outside the officer `/attendance` values.
11. Registration: the §G.8 option changes, and the commands hash changes.
12. Posting: a save within 7 days posts at once; further out sets `post_at` 7 days before; a repeat posts one at a time; the reminder at start minus 24 h edits the post into `Tomorrow` when it is the newest message, else replies with the five components and no mentions; no reminder when the post is under 12 h old; move, cancel and lock delete the reminder; a webhook reminder carries the masked link and no components; every posted raid URL passes the `parseGuildRaidPath` mirror.
13. The payload dump (`payload-dump.go.txt`) gains every card frame and both reminder forms, so the reviewers' boards render from the real handlers.
14. Roster: the post's fields are `When`, `Guild time`, `Signed up` (inline, `roster.signed`), then `Tanks`, `Healers`, `Melee`, `Ranged`, `Late`, `Bench` (inline) and `Out` (full width) per §P.9.2; nobody In or Late draws no groups; an empty group draws `None yet`; lines sort by class then name, Out by name; Melee and Ranged split per §P.9.4, a spec-less Rogue in Melee and a spec-less Mage in Ranged.
15. Roster limits: a 17-line group draws a U+200B continuation; a 40-line group draws 31 lines and `and 9 more`; Bench at 11 draws 9 and `and 2 more`; every field value at most 1024; at most 25 fields; 120 sign-ups keep the embed under 6000 through the guard.
16. Marks: with the emoji map loaded each line is `<:fs_{class}_{spec}:{id}> {Name}`; with a spec missing from the map, that spec's lines are the text line `{Name} · {Spec} {Class}`; a row with no spec is `{Name} · {Class}`; every webhook roster is text lines; the emoji names are exactly the 28 `specs.json` slugs with hyphens as underscores, each icon file exists in the active build's icon folder, and no short spec word exists in the code; `api discord-emoji` against a stub uploads only the missing names and prints no token.
17. Spec card: the options are the class's specs by full name (`Beast Mastery`, `Retribution`, never a short word); an In press with no spec known answers the card (type 4 from the post, type 7 after the character select), with title, lead, colour, crest and the class's specs; an Out press never asks; a pick saves `raid_spec`, signs up with that spec's role and answers the outcome plus `spec.kept` and `Change spec`; a forged key answers `pressCopy.Failure`; a locked raid answers its sentence; `Change spec` never changes the state; the next press with the spec saved signs up with no card.

## Site tests

`guild-raids.spec.ts` (API stubbed): the form draws `Post to Discord` checked with `role="switch"` and `Remind the day before` checked; each helper per §P.3; turning `Post to Discord` off disables the reminder with its helper; the saved values return as the next raid's defaults; edit of a posted raid shows `form.posted`; axe clean; both rows 44 px at 360; zero layout shift as the defaults load.

## Boards the wow-player reviews

The mock boards above, then the same frames rendered from the real handlers' `payloads.json` through `capture-messages.spec.ts` at 640 and 390, beside each mock board. A rendered frame that differs from its board in words, colour, button style, row order or pre-chosen values is a defect against this spec; Discord's own chrome is not.

## Discord captures (test server, production application), added to the bot spec's §9 table

| # | Message | Desktop client | Phone client |
|---|---|---|---|
| D13 | `/bis` bare, unlinked and linked: each step, the nudge, the result, `Change picks`; one pick made 20 minutes after the card was drawn (decides `card.footer.window`) | x | x |
| D14 | `/char` bare linked and unlinked, the modal open, an unknown name | x | x |
| D15 | `/signup` bare with two raids and two characters, a status press | x | x |
| D16 | `/attendance` officer select, `/readiness` and `/report` with their selects | x | x |
| D17 | A raid saved within 7 days posting at once; the `Tomorrow` edit; the reply reminder; a webhook reminder; a cancel deleting the reminder | x | x |
| D18 | The roster post at 8 of 40, 33 of 40 and 40 of 40, in a server where the bot has never posted (the application emoji render without any server permission); one spec's emoji deleted (the text line); the phone client's inline layout | x | x |
| D19 | `In` with no spec known: the card, the pick, the outcome, `Change spec`, the next press with no card | x | x |

Site captures, added to the bot spec's §9: the raid form's Discord block (create, create with repeat, switch off, edit posted) at 360, 390 and 1440.

## Questions only the owner can answer (part P roster)

1. Melee and Ranged in place of one DPS column (ruling P7). It follows Raid-Helper, but the brief said Tanks, Healers, DPS. Keep the split?
2. Feral and Feral Bear share one icon, so their two emoji look the same in the roster, and the Melee and Tanks groups tell them apart. Keep it (the game's own icon), or draw a distinct mark for Feral Bear?
3. Answered (owner, 2026-10-28): the text fallback carries the class word, `{Name} · {Spec} {Class}` (`Aurelith · Holy Paladin`), on every text line.
4. A row with no spec (only sign-ups made before the spec ask) draws `{Name} · {Class}` with no icon. Fine, or ask officers to set it?
3. The site has no control for a character's spec today. The coordinator's "spec from what the character has set on the site" therefore means only what the Discord card saves, or the logs. Should the raid page's own sign-up ask the spec too, so a site-only member is never a bare name?

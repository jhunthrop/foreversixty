<!-- web/src/components/guild/RaidBandLine.svelte -->
<!-- The raid line in the guild band (raids spec 3.4, 6.1, R8): the sentence, the current
     character's segments and the `{Zone} →` link, between the standing line and the needs line.
     From md one 44 px row; on phone the sentence and the link share the first row and the
     segments take the second (at most 112 px). A press rewrites the sentence in place; a refused
     press shows its line under the row for 6 s, and a lock or cancel race updates the band. `now`
     ticks each minute, so `Tonight` turns over at midnight. -->
<script lang="ts">
  import { tick } from 'svelte';
  import { minuteClock } from '../../lib/ui/minute-clock.svelte';
  import { guildRaidHref } from '../../lib/characters';
  import type { GuildRaidRole, GuildSignupState, NextRaidView } from '../../lib/guild/api';
  import {
    bandClause,
    bandLineParts,
    kindCountsRoles,
    kindOf,
    raidCopy,
    whenPhrase,
  } from '../../lib/guild/raid-copy';
  import { classTone } from '../../lib/ui/primitives';
  import Crest from '../ui/Crest.svelte';
  import Row from '../ui/Row.svelte';
  import { isRaceLine, pressSignup } from '../../lib/guild/signup-save';
  import { guildGateCopy } from '../../lib/guild/copy';
  import { gateAction, type GateFacts } from '../../lib/guild/plan-gate';
  import { viewerLocale, viewerZone } from '../../lib/dates';
  import GateActionInline from './GateActionInline.svelte';
  import SignupSegments from './SignupSegments.svelte';

  /** One of the account's characters on the roster; the row layout reads its class, spec and level. */
  interface OwnCharacter {
    character_key: string;
    name: string;
    class?: string;
    spec?: string;
    level?: number | null;
  }

  let {
    guildId,
    nextRaid,
    officer,
    officerTools = true,
    gateFacts = null,
    guildPath,
    currentKey,
    ownCharacters,
    defaultRole,
    onChange,
    onStatusLine = () => {},
    now: fixedNow,
    locale = viewerLocale(),
    viewerTimezone = viewerZone(),
    layout = 'band',
  }: {
    guildId: number;
    nextRaid: NextRaidView | null;
    officer: boolean;
    /** False when the guild's plan gates officer tools: the none-planned line drops its link. */
    officerTools?: boolean;
    /** The gate sentence's facts: a gated none-planned line carries the same plan link the Raids
     *  tab does (guild round 3). Null draws part 1 alone. */
    gateFacts?: GateFacts | null;
    guildPath: string;
    /** The header character, when it is in this guild; no segments without one. */
    currentKey: string | null;
    /** The account's characters on this guild's roster: two or more name the character. */
    ownCharacters: readonly OwnCharacter[];
    /** The current character's role when it has not signed up (the roster's server role). */
    defaultRole: GuildRaidRole;
    /** The band's raid after a press; `null` once a race says the raid was cancelled. */
    onChange: (next: NextRaidView | null) => void;
    /** Whether a status line shows under the row: the guild band saves its height only at rest. */
    onStatusLine?: (shown: boolean) => void;
    now?: Date;
    locale?: string;
    viewerTimezone?: string;
    /** `band`: the sentence, the segments and the raid link (the guild band on every tab but the
     *  Overview). `row`: the viewer's own Row in the Overview's next-raid Plate (web layout spec
     *  2026-11-05 §7.1): crest, character, spec and level, their state, the segments. */
    layout?: 'band' | 'row';
  } = $props();

  /** The space between the gate's two parts, an expression so block-edge trimming keeps it. */
  const WORD_SPACE = ' ';
  /** A refused press's line stays this long, then the band is back at rest. */
  const STATUS_LINE_MS = 6000;

  let pending = $state<GuildSignupState | null>(null);
  let busy = $state(false);
  let status = $state('');
  let statusLine = $state<HTMLParagraphElement | null>(null);
  let statusTimer: ReturnType<typeof setTimeout> | undefined;
  const clock = minuteClock(() => fixedNow);

  $effect(() => onStatusLine(status !== ''));
  $effect(() => () => clearTimeout(statusTimer));

  function showStatus(line: string): void {
    clearTimeout(statusTimer);
    status = line;
    statusTimer = setTimeout(() => (status = ''), STATUS_LINE_MS);
  }

  const own = $derived(nextRaid?.own ?? null);
  const locked = $derived(nextRaid?.status === 'locked');
  const lockedId = $derived(nextRaid === null ? '' : `guild-raid-line-locked-${nextRaid.id}`);
  /** The character the segments act on (spec 3.4): another of the account's characters the API
   *  has In or Late, which the sentence names, so a press never benches it unseen; else the
   *  header's current character. */
  const boundKey = $derived(
    own !== null && own.character_key !== currentKey && (own.state === 'in' || own.state === 'late')
      ? own.character_key
      : currentKey,
  );
  const bound = $derived(own !== null && own.character_key === boundKey ? own : null);
  const boundState = $derived(pending ?? bound?.state ?? null);
  /** What the sentence describes: a press in flight, else the API's own sign-up. */
  const shown = $derived(
    pending !== null && boundKey !== null
      ? { key: boundKey, state: pending, role: bound?.role ?? defaultRole }
      : own === null
        ? null
        : { key: own.character_key, state: own.state, role: own.role },
  );
  const boundName = $derived(ownCharacters.find((c) => c.character_key === boundKey)?.name ?? '');

  /** The sentence around its state word, which reads 600 `--strong` (spec 3.4). */
  const parts = $derived.by(() => {
    if (nextRaid === null) return { head: '', word: '', tail: '' };
    const startsAt = new Date(nextRaid.starts_at);
    const when = whenPhrase(
      startsAt,
      clock.now,
      locale,
      nextRaid.guild_timezone ?? viewerTimezone,
      viewerTimezone,
    );
    const named = ownCharacters.length >= 2 && shown !== null;
    return bandLineParts({
      when,
      state: shown?.state ?? null,
      // A guild event counts no roles: `you are In`, never `In as DPS` (guild round 3).
      role: kindCountsRoles(kindOf(nextRaid)) ? shown?.role : undefined,
      characterName: named ? ownCharacters.find((c) => c.character_key === shown.key)?.name : undefined,
      n: nextRaid.counts.in + nextRaid.counts.late,
      size: nextRaid.size,
    });
  });

  /** The row's character: the one the segments act on (its name is the row's title). */
  const rowCharacter = $derived(ownCharacters.find((c) => c.character_key === boundKey) ?? null);
  /** The row's sub line: `Holy Priest · level 60 · ` before the state clause. */
  const rowFacts = $derived.by(() => {
    if (rowCharacter === null) return '';
    const cls = (rowCharacter.class ?? '').charAt(0).toUpperCase() + (rowCharacter.class ?? '').slice(1);
    const what = [rowCharacter.spec ?? '', cls].filter(Boolean).join(' ');
    const level = rowCharacter.level == null ? '' : `level ${rowCharacter.level}`;
    const parts = [what, level].filter((part) => part !== '');
    return parts.length === 0 ? '' : `${parts.join(' · ')} · `;
  });
  const rowClause = $derived(
    bandClause({
      when: '',
      state: shown?.state ?? null,
      role: nextRaid !== null && kindCountsRoles(kindOf(nextRaid)) ? shown?.role : undefined,
      n: 0,
      size: null,
    }),
  );

  async function press(state: GuildSignupState): Promise<void> {
    if (nextRaid === null || boundKey === null || busy || locked) return;
    const prev = nextRaid;
    const key = boundKey;
    const mine = bound;
    pending = state;
    busy = true;
    clearTimeout(statusTimer);
    status = '';
    // The PUT replaces the note, so the saved one travels with the press.
    const body = { state, role: mine?.role ?? defaultRole, note: mine?.note ?? '' };
    try {
      const outcome = await pressSignup(guildId, prev, key, body, currentKey ?? key);
      if (outcome.next !== undefined) onChange(outcome.next);
      if (!outcome.ok) {
        showStatus(outcome.line);
        await tick();
        // Spec 4 "Races": the race line takes focus; the generic line is only announced.
        if (isRaceLine(outcome.line)) statusLine?.focus();
      }
    } finally {
      pending = null;
      busy = false;
    }
  }
</script>

{#snippet statusRow()}
  {#if status !== ''}
    <p
      bind:this={statusLine}
      class="text-wipe text-[13px]"
      role="alert"
      tabindex="-1"
      data-testid="guild-raid-line-status"
    >
      {status}
    </p>
  {/if}
{/snippet}

{#if nextRaid === null}
  <p class="text-[14px]" data-testid="guild-raid-line-none">
    {officer ? raidCopy.bandNone.officer : raidCopy.bandNone.member}
    {#if officer && !officerTools}
      <!-- The gate sentence, part 2 included: the guild master meets the trial link here too. -->
      <span data-testid="guild-raid-line-gate"
        >{guildGateCopy.lead.raids}{#if gateFacts !== null}{WORD_SPACE}<GateActionInline
            action={gateAction('raids', gateFacts)}
          />{/if}</span
      >
    {:else if officer}
      <a href="#raids" class="text-[13px] font-bold" data-testid="guild-raid-line-plan"
        >{raidCopy.bandNone.planLink}</a
      >
    {/if}
  </p>
  {@render statusRow()}
{:else if layout === 'row'}
  <div data-testid="guild-raid-line" aria-busy={busy ? 'true' : undefined}>
    <Row title={boundName} titleTone={classTone(rowCharacter?.class ?? '')} wrapTrail>
      {#snippet lead()}<Crest cls={rowCharacter?.class ?? ''} size={36} />{/snippet}
      {#snippet subContent()}<span data-testid="guild-raid-line-sentence"
          >{rowFacts}{rowClause.head}{#if rowClause.word !== ''}<strong class="text-strong font-semibold"
              >{rowClause.word}</strong
            >{/if}{rowClause.tail}</span
        >{/snippet}
      {#snippet trail()}
        {#if boundKey !== null}
          <SignupSegments
            value={boundState}
            name={boundName}
            size="band"
            disabled={locked}
            describedBy={locked ? lockedId : undefined}
            onChange={(state) => void press(state)}
            testid="guild-raid-line-segments"
          />
        {/if}
      {/snippet}
    </Row>
    {#if locked}
      <p id={lockedId} class="sr-only" data-testid="guild-raid-line-locked">{raidCopy.raidStates.locked}</p>
    {/if}
    <div class="px-4">{@render statusRow()}</div>
  </div>
{:else}
  <div class="flex flex-col gap-1" data-testid="guild-raid-line">
    <div
      class="grid grid-cols-[1fr_auto] items-center gap-x-4 gap-y-2 md:flex md:min-h-11 md:justify-between"
      aria-busy={busy ? 'true' : undefined}
    >
      <p class="text-text line-clamp-2 text-[14px] md:flex-1" data-testid="guild-raid-line-sentence">
        {parts.head}{#if parts.word !== ''}<strong class="text-strong font-semibold">{parts.word}</strong
          >{/if}{parts.tail}
      </p>
      {#if boundKey !== null}
        <div class="col-span-2 row-start-2 md:shrink-0">
          <SignupSegments
            value={boundState}
            name={boundName}
            size="band"
            disabled={locked}
            describedBy={locked ? lockedId : undefined}
            onChange={(state) => void press(state)}
            testid="guild-raid-line-segments"
          />
        </div>
      {/if}
      <a
        href={guildRaidHref(guildPath, nextRaid.id)}
        class="text-gold col-start-2 row-start-1 inline-flex h-11 items-center text-[13px] font-bold whitespace-nowrap md:h-auto"
        data-testid="guild-raid-line-link">Open the raid ›</a
      >
    </div>
    {#if locked}
      <!-- The disabled segments' reason, read with them; sighted viewers see the inert group and
           the raid page states it in full. Kept off screen so the resting block stays in its cap. -->
      <p id={lockedId} class="sr-only" data-testid="guild-raid-line-locked">{raidCopy.raidStates.locked}</p>
    {/if}
    {@render statusRow()}
  </div>
{/if}

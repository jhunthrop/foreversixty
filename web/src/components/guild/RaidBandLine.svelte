<!-- web/src/components/guild/RaidBandLine.svelte -->
<!-- The raid line in the guild band (raids spec 3.4, 6.1, R8): the sentence, the current
     character's segments and the `{Zone} →` link, between the standing line and the needs line.
     From md one 44 px row; on phone the sentence and the link share the first row and the
     segments take the second (at most 112 px). A press rewrites the sentence in place; a refused
     press shows its line under the row for 6 s, and a lock or cancel race updates the band. `now`
     ticks each minute, so `Tonight` turns over at midnight. -->
<script module lang="ts">
  import { GuildApiError, type GuildRaidView, type NextRaidView } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';

  /** The band's fields after a sign-up, from the PUT's raid: the API's own rule for `own` (the
   *  current character's sign-up, else the account's first, else none). */
  export function nextRaidAfter(prev: NextRaidView, view: GuildRaidView, currentKey: string): NextRaidView {
    const signed = view.viewer.characters.filter((c) => c.state !== null);
    const own = signed.find((c) => c.character_key === currentKey) ?? signed[0];
    const { tanks, healers, dps } = view.counts;
    const note = (key: string): string => view.signups.find((s) => s.character_key === key)?.note ?? '';
    return {
      ...prev,
      starts_at: view.starts_at,
      status: view.status === 'locked' ? 'locked' : 'planned',
      size: view.size,
      counts: { in: view.counts.in, late: view.counts.late, tanks, healers, dps },
      own:
        own === undefined || own.state === null
          ? null
          : {
              character_key: own.character_key,
              state: own.state,
              role: own.role,
              note: note(own.character_key),
            },
    };
  }

  /** The status line after a refused press (spec 4 "Races", 6.1 strip status). */
  export function pressFailedLine(thrown: unknown): string {
    if (thrown instanceof GuildApiError && thrown.code === 'raid_locked') return raidCopy.races.locked;
    if (thrown instanceof GuildApiError && thrown.code === 'raid_cancelled') return raidCopy.races.cancelled;
    return raidCopy.stripStatus.failed;
  }

  /** The band after a refused press: a lock locks it, a cancel drops the raid (the band then says
   *  none is planned); `undefined` leaves it as it was. */
  export function nextRaidAfterRefusal(prev: NextRaidView, thrown: unknown): NextRaidView | null | undefined {
    if (!(thrown instanceof GuildApiError)) return undefined;
    if (thrown.code === 'raid_locked') return { ...prev, status: 'locked' };
    if (thrown.code === 'raid_cancelled') return null;
    return undefined;
  }
</script>

<script lang="ts">
  import { viewerLocale } from '../../lib/guild/raid-time';
  import { tick } from 'svelte';
  import { minuteClock } from '../../lib/guild/minute-clock.svelte';
  import { guildRaidHref } from '../../lib/characters';
  import type { GuildRaidRole, GuildSignupState } from '../../lib/guild/api';
  import { bandLineParts, whenPhrase } from '../../lib/guild/raid-copy';
  import { putSignup } from '../../lib/guild/raids-api';
  import SignupSegments from './SignupSegments.svelte';

  interface OwnCharacter {
    character_key: string;
    name: string;
  }

  let {
    guildId,
    nextRaid,
    officer,
    guildPath,
    currentKey,
    ownCharacters,
    defaultRole,
    onChange,
    onStatusLine = () => {},
    now: fixedNow,
    locale = viewerLocale(),
    viewerTimezone = new Intl.DateTimeFormat().resolvedOptions().timeZone,
  }: {
    guildId: number;
    nextRaid: NextRaidView | null;
    officer: boolean;
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
  } = $props();

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
      role: shown?.role,
      characterName: named ? ownCharacters.find((c) => c.character_key === shown.key)?.name : undefined,
      n: nextRaid.counts.in + nextRaid.counts.late,
      size: nextRaid.size,
    });
  });

  async function press(state: GuildSignupState): Promise<void> {
    if (nextRaid === null || boundKey === null || busy || locked) return;
    const prev = nextRaid;
    const key = boundKey;
    const mine = bound;
    pending = state;
    busy = true;
    clearTimeout(statusTimer);
    status = '';
    try {
      // The PUT replaces the note, so the saved one travels with the press.
      const body = { state, role: mine?.role ?? defaultRole, note: mine?.note ?? '' };
      const { raid } = await putSignup(guildId, prev.id, key, body);
      onChange(nextRaidAfter(prev, raid, currentKey ?? key));
    } catch (thrown) {
      const after = nextRaidAfterRefusal(prev, thrown);
      if (after !== undefined) onChange(after);
      showStatus(pressFailedLine(thrown));
      await tick();
      // Spec 4 "Races": the race line takes focus; the generic line is only announced.
      if (status !== raidCopy.stripStatus.failed) statusLine?.focus();
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
    {#if officer}
      <a href="#raids" class="text-[13px] font-bold" data-testid="guild-raid-line-plan"
        >{raidCopy.bandNone.planLink}</a
      >
    {/if}
  </p>
  {@render statusRow()}
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
        data-testid="guild-raid-line-link">{nextRaid.zone_name} →</a
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

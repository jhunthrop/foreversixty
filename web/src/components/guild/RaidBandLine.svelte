<!-- web/src/components/guild/RaidBandLine.svelte -->
<!-- The raid line in the guild band (raids spec 3.4, 6.1, R8): the sentence, the current
     character's segments and the `{Zone} →` link, between the standing line and the needs line.
     From md one 44 px row; on phone the sentence and the link share the first row and the
     segments take the second (at most 112 px). A press rewrites the sentence in place. -->
<script module lang="ts">
  import { GuildApiError, type GuildRaidView, type NextRaidView } from '../../lib/guild/api';
  import { raidCopy } from '../../lib/guild/raid-copy';

  /** The band's fields after a sign-up, from the PUT's raid: the API's own rule for `own` (the
   *  current character's sign-up, else the account's first, else none). */
  export function nextRaidAfter(prev: NextRaidView, view: GuildRaidView, currentKey: string): NextRaidView {
    const signed = view.viewer.characters.filter((c) => c.state !== null);
    const own = signed.find((c) => c.character_key === currentKey) ?? signed[0];
    const { tanks, healers, dps } = view.counts;
    return {
      ...prev,
      starts_at: view.starts_at,
      size: view.size,
      counts: { in: view.counts.in, late: view.counts.late, tanks, healers, dps },
      own:
        own === undefined || own.state === null
          ? null
          : { character_key: own.character_key, state: own.state, role: own.role },
    };
  }

  /** The status slot's line after a refused press (spec 4 "Races", 6.1 strip status). */
  export function pressFailedLine(thrown: unknown): string {
    if (thrown instanceof GuildApiError && thrown.code === 'raid_locked') return raidCopy.races.locked;
    if (thrown instanceof GuildApiError && thrown.code === 'raid_cancelled') return raidCopy.races.cancelled;
    return raidCopy.stripStatus.failed;
  }
</script>

<script lang="ts">
  import { tick } from 'svelte';
  import { guildRaidHref } from '../../lib/characters';
  import type { GuildRaidRole, GuildSignupState } from '../../lib/guild/api';
  import { bandLine, whenPhrase } from '../../lib/guild/raid-copy';
  import { fetchGuildRaid, putSignup } from '../../lib/guild/raids-api';
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
    now = new Date(),
    locale = 'en-US',
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
    onChange: (next: NextRaidView) => void;
    now?: Date;
    locale?: string;
    viewerTimezone?: string;
  } = $props();

  let pending = $state<GuildSignupState | null>(null);
  let busy = $state(false);
  let status = $state('');
  let statusLine = $state<HTMLParagraphElement | null>(null);

  const own = $derived(nextRaid?.own ?? null);
  const currentState = $derived(
    pending ?? (own !== null && own.character_key === currentKey ? own.state : null),
  );
  /** What the sentence describes: a press in flight, else the API's own sign-up. */
  const shown = $derived(
    pending !== null && currentKey !== null
      ? { key: currentKey, state: pending, role: own?.character_key === currentKey ? own.role : defaultRole }
      : own === null
        ? null
        : { key: own.character_key, state: own.state, role: own.role },
  );
  const currentName = $derived(ownCharacters.find((c) => c.character_key === currentKey)?.name ?? '');

  const sentence = $derived.by(() => {
    if (nextRaid === null) return '';
    const startsAt = new Date(nextRaid.starts_at);
    const when = whenPhrase(startsAt, now, locale, nextRaid.guild_timezone ?? viewerTimezone, viewerTimezone);
    const named = ownCharacters.length >= 2 && shown !== null;
    return bandLine({
      when,
      state: shown?.state ?? null,
      role: shown?.role,
      characterName: named ? ownCharacters.find((c) => c.character_key === shown.key)?.name : undefined,
      n: nextRaid.counts.in + nextRaid.counts.late,
      size: nextRaid.size,
    });
  });
  /** The sentence around its state word, which reads 600 `--strong` (spec 3.4). */
  const parts = $derived.by(() => {
    if (shown === null) return { head: sentence, word: '', tail: '' };
    const word = raidCopy.segments[shown.state];
    const at = sentence.indexOf(` ${word}`, sentence.indexOf('you ')) + 1;
    return { head: sentence.slice(0, at), word, tail: sentence.slice(at + word.length) };
  });

  /** The PUT replaces the note, so a signed character's note is read from the raid first. */
  async function savedNote(raidId: number, key: string): Promise<string> {
    if (own === null) return '';
    const view = await fetchGuildRaid(guildId, raidId);
    return view.signups.find((s) => s.character_key === key)?.note ?? '';
  }

  async function press(state: GuildSignupState): Promise<void> {
    if (nextRaid === null || currentKey === null || busy) return;
    const prev = nextRaid;
    const key = currentKey;
    const role = own?.character_key === key ? own.role : defaultRole;
    pending = state;
    busy = true;
    status = '';
    try {
      const note = await savedNote(prev.id, key);
      const { raid } = await putSignup(guildId, prev.id, key, { state, role, note });
      onChange(nextRaidAfter(prev, raid, key));
    } catch (thrown) {
      status = pressFailedLine(thrown);
      await tick();
      // Spec 4 "Races": the race line takes focus; the generic line is only announced.
      if (status !== raidCopy.stripStatus.failed) statusLine?.focus();
    } finally {
      pending = null;
      busy = false;
    }
  }
</script>

{#if nextRaid === null}
  <p class="text-[14px]" data-testid="guild-raid-line-none">
    {officer ? raidCopy.bandNone.officer : raidCopy.bandNone.member}
    {#if officer}
      <a href="#raids" class="text-[13px] font-bold" data-testid="guild-raid-line-plan"
        >{raidCopy.bandNone.planLink}</a
      >
    {/if}
  </p>
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
      {#if currentKey !== null}
        <div class="col-span-2 row-start-2 md:shrink-0">
          <SignupSegments
            value={currentState}
            name={currentName}
            size="band"
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
    <p
      bind:this={statusLine}
      class="text-wipe min-h-5 text-[13px]"
      role="alert"
      tabindex="-1"
      data-testid="guild-raid-line-status"
    >
      {status}
    </p>
  </div>
{/if}

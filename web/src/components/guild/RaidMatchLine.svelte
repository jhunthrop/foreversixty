<!-- web/src/components/guild/RaidMatchLine.svelte -->
<!-- The match line under the night line (attendance spec 3.1): line 1 is `Logged as a raid night.`
     with `Open the night`, which opens the night expanded; line 2, muted, says who uploaded the log
     and how it was matched. Officers get `Not this log`, which asks `RaidClearConfirm` first. The
     gold is the link, once. A refusal's line sits under the row. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { RaidMatch } from '../../lib/guild/attendance-contract';
  import { attendanceCopy, uploadedLine } from '../../lib/guild/attendance-copy';
  import { PHONE_HIT_AREA } from '../../lib/guild/field-styles';
  import { focusOnMount } from '../../lib/guild/focus-now';
  import type { MatchOutcome } from '../../lib/guild/match-picker.svelte';
  import { RAID_TEXT_BUTTON } from '../../lib/guild/raid-buttons';
  import { nightHref } from '../../lib/guild/raids-hash';
  import RaidClearConfirm from './RaidClearConfirm.svelte';

  let {
    match,
    guildPath,
    viewerTimezone,
    guildTimezone,
    locale,
    officer,
    readOnly,
    gateId,
    onClear,
    confirming = $bindable(false),
  }: {
    match: RaidMatch;
    guildPath: string;
    viewerTimezone: string;
    guildTimezone: string;
    locale: string;
    officer: boolean;
    /** Officer-read-only: `Not this log` stays but is disabled (gate spec 3.2). */
    readOnly: boolean;
    /** The gate sentence's id, read while disabled. */
    gateId: string;
    /** Clears the match; the page answers instead of throwing. */
    onClear: () => Promise<MatchOutcome>;
    /** The clear confirm is open; a test seeds it. */
    confirming?: boolean;
  } = $props();

  const copy = attendanceCopy.match;
  let clearing = $state(false);
  let failure = $state<string | null>(null);
  let notThisLog = $state<HTMLButtonElement | null>(null);

  async function keep(): Promise<void> {
    confirming = false;
    failure = null;
    await tick();
    notThisLog?.focus();
  }

  async function clear(): Promise<void> {
    clearing = true;
    failure = null;
    const outcome = await onClear();
    clearing = false;
    // A cleared match swaps this whole line for the ended one, which takes focus.
    if (!outcome.ok) failure = outcome.message;
  }
</script>

<div class="flex flex-col text-[13px]" data-testid="raid-match-line">
  <p class="text-text">
    {copy.loggedAsNight}&#32;<a
      id="raid-open-night"
      href={nightHref(guildPath, match.report_id)}
      class="text-gold font-bold {PHONE_HIT_AREA}"
      data-testid="raid-open-night">{copy.openNight}</a
    >
  </p>
  <div class="flex flex-wrap items-center gap-x-3">
    <p class="text-muted" data-testid="raid-match-uploaded">
      {uploadedLine(match, viewerTimezone, guildTimezone, locale)}
    </p>
    {#if officer && !confirming}
      <button
        type="button"
        class={RAID_TEXT_BUTTON}
        bind:this={notThisLog}
        disabled={readOnly || clearing}
        aria-describedby={readOnly ? gateId : undefined}
        onclick={() => (confirming = true)}
        data-testid="raid-not-this-log">{clearing ? copy.clearing : copy.notThisLog}</button
      >
    {/if}
  </div>
  {#if confirming}
    <div class="mt-2">
      <RaidClearConfirm busy={clearing} onConfirm={() => void clear()} onKeep={() => void keep()} />
    </div>
  {/if}
  {#if failure !== null}
    {#key failure}
      <p
        class="text-wipe mt-1 text-[13px]"
        role="alert"
        tabindex="-1"
        use:focusOnMount
        data-testid="raid-match-error"
      >
        {failure}
      </p>
    {/key}
  {/if}
</div>

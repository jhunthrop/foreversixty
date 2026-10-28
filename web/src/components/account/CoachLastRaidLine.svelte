<!-- web/src/components/account/CoachLastRaidLine.svelte -->
<!-- The account's "last raid" coach line (experience spec 2026-10-28-log-coach section 3.3): the
     first child of the Characters panel. It loads GET /v1/me/coach/recent on its own once the
     panel is on screen, so a cold cache never slows the page, and holds one line's height
     (44 px desktop, 64 px phone) in every state so nothing below it moves. A "pending" answer
     (packs still being read) refreshes once after 5 s. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchCoachRecent } from '../../lib/coach/api';
  import { coachCopy } from '../../lib/coach/copy';
  import {
    LAST_RAID_LABEL,
    createLastRaidLoader,
    type LastRaidRole,
    type LastRaidState,
  } from '../../lib/coach/last-raid';
  import { uiCopy } from '../../lib/ui/copy';
  import type { CoachRecent } from '../../lib/coach/types';

  let {
    premium,
    fetchRecent = () => fetchCoachRecent(),
    refreshDelayMs = 5000,
  }: {
    /** `me.entitlements.log_coach`. */
    premium: boolean;
    fetchRecent?: () => Promise<CoachRecent>;
    refreshDelayMs?: number;
  } = $props();

  let state = $state<LastRaidState>({ kind: 'loading' });

  const ROLE_CLASS: Record<LastRaidRole, string> = {
    lead: 'text-muted',
    place: 'text-strong font-semibold',
    plain: 'text-text',
    link: 'text-gold underline hover:text-gold-hover focus-visible:outline-2 focus-visible:outline-gold',
    muted: 'text-muted',
  };

  const loader = createLastRaidLoader({
    fetchRecent: () => fetchRecent(),
    premium,
    refreshDelayMs,
    onState: (next) => (state = next),
  });

  onMount(() => {
    loader.start();
    return loader.stop;
  });

  const busy = $derived(state.kind === 'loading' || state.kind === 'retrying');
</script>

<div
  class="border-line-soft flex min-h-[77px] items-center border-b py-2 text-[13px] leading-5 md:min-h-[57px]"
  role="status"
  aria-label={LAST_RAID_LABEL}
  aria-live="polite"
  aria-busy={busy}
  data-testid="coach-last-raid"
  data-state={state.kind}
>
  {#if state.kind === 'loading'}
    <span class="skeleton-block h-4 w-3/4" aria-hidden="true" data-testid="coach-last-raid-skeleton"></span>
    <span class="sr-only">{uiCopy.loading}</span>
  {:else if state.kind === 'error' || state.kind === 'retrying'}
    <p class="text-muted">
      {coachCopy.lastRaidError}
      <button
        type="button"
        class="text-gold hover:text-gold-hover focus-visible:outline-gold underline focus-visible:outline-2 disabled:opacity-60"
        data-testid="coach-last-raid-retry"
        disabled={state.kind === 'retrying'}
        aria-busy={state.kind === 'retrying'}
        onclick={() => loader.retry()}>{coachCopy.rowRetry}</button
      >
    </p>
  {:else}
    <p>
      {#each state.view.segments as segment, index (index)}
        {#if index > 0}<span class="text-muted"> · </span>{/if}
        {#if segment.role === 'link' && state.view.href !== null}
          <a class={ROLE_CLASS.link} href={state.view.href} data-testid="coach-last-raid-link"
            >{segment.text}</a
          >
        {:else if state.view.stem !== null && segment.text.includes(state.view.stem)}
          {@const [before, after] = segment.text.split(state.view.stem)}
          <span class={ROLE_CLASS[segment.role]}
            >{before}<span class="inline-block max-w-[24ch] truncate align-bottom" title={state.view.stem}
              >{state.view.stem}</span
            >{after}</span
          >
        {:else}
          <span class={ROLE_CLASS[segment.role]}>{segment.text}</span>
        {/if}
      {/each}
    </p>
  {/if}
</div>

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
  import { settleLastRaid, type LastRaidRole, type LastRaidState } from '../../lib/coach/last-raid';
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

  async function load(): Promise<LastRaidState> {
    state = await settleLastRaid(fetchRecent, premium);
    return state;
  }

  onMount(() => {
    let timer: ReturnType<typeof setTimeout> | undefined;
    let alive = true;
    void load().then((settled) => {
      if (!alive || settled.kind !== 'ready' || !settled.pending) return;
      timer = setTimeout(() => void load(), refreshDelayMs);
    });
    return () => {
      alive = false;
      clearTimeout(timer);
    };
  });
</script>

<div
  class="border-line-soft flex min-h-16 items-center border-b py-2 text-[13px] leading-5 md:min-h-11"
  data-testid="coach-last-raid"
  data-state={state.kind}
  aria-busy={state.kind === 'loading'}
>
  {#if state.kind === 'error'}
    <p class="text-muted" role="status">
      {coachCopy.lastRaidError}
      <button
        type="button"
        class="text-gold hover:text-gold-hover focus-visible:outline-gold underline focus-visible:outline-2"
        data-testid="coach-last-raid-retry"
        onclick={() => void load()}>{coachCopy.rowRetry}</button
      >
    </p>
  {:else if state.kind === 'ready'}
    <p>
      {#each state.view.segments as segment, index (index)}
        {#if index > 0}<span class="text-muted"> · </span>{/if}
        {#if segment.role === 'link' && state.view.href !== null}
          <a class={ROLE_CLASS.link} href={state.view.href} data-testid="coach-last-raid-link"
            >{segment.text}</a
          >
        {:else}
          <span class={ROLE_CLASS[segment.role]}>{segment.text}</span>
        {/if}
      {/each}
    </p>
  {/if}
</div>

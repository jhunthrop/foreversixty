<!-- web/src/components/account/CoachLastRaidLine.svelte -->
<!-- The account's "last raid" coach line (experience spec 2026-10-28-log-coach section 3.3; premium
     spec 2026-11-08 section 7.2): the Characters Plate's foot row, under the character rows. It
     loads GET /v1/me/coach/recent on its own, so a cold cache never slows the page, and draws only
     when there is a last raid: while loading, on a failed read and with no raid it draws nothing
     (G5). A "pending" answer (packs still being read) refreshes once after 5 s. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchCoachRecent } from '../../lib/coach/api';
  import { COACH_SEPARATOR } from '../../lib/coach/copy';
  import {
    LAST_RAID_LABEL,
    cachedLastRaid,
    createLastRaidLoader,
    lastRaidLine,
    type LastRaidRole,
    type LastRaidState,
  } from '../../lib/coach/last-raid';
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

  // A remount (the account page re-creates this line when /v1/me refreshes) starts from the
  // page's last answer, so the skeleton shows once per page, never over a line already read.
  let state = $state<LastRaidState>(cachedLastRaid() ?? { kind: 'loading' });

  const ROLE_CLASS: Record<LastRaidRole, string> = {
    lead: 'text-muted',
    place: 'text-strong font-semibold',
    plain: 'text-text',
    // A 44 px hit area around the link (the touch floor), drawn by ::after so the line keeps its height.
    link: "text-gold relative underline hover:text-gold-hover focus-visible:outline-2 focus-visible:outline-gold after:absolute after:inset-x-0 after:top-1/2 after:h-11 after:-translate-y-1/2 after:content-['']",
    muted: 'text-muted',
  };

  // Getters, not captured values: `premium` follows /v1/me (stale-while-revalidate, so it can
  // flip after first paint), and the copy below is re-derived from the raw answer when it does.
  const loader = createLastRaidLoader({
    fetchRecent: () => fetchRecent(),
    get premium() {
      return premium;
    },
    get refreshDelayMs() {
      return refreshDelayMs;
    },
    onState: (next) => (state = next),
  });

  onMount(() => {
    loader.start();
    return loader.stop;
  });

  // Only a raid draws a line (spec 2026-11-08 G5): loading, a failed read and no raid draw nothing.
  const view = $derived(
    state.kind === 'ready' && state.recent.last_raid !== null ? lastRaidLine(state.recent, premium) : null,
  );
</script>

{#if view !== null}
  <p
    class="px-4 py-3 text-[13px] leading-5"
    role="status"
    aria-label={LAST_RAID_LABEL}
    data-testid="coach-last-raid"
    data-state={state.kind}
  >
    {#each view.segments as segment, index (index)}
      {#if index > 0}<span class="text-muted">{COACH_SEPARATOR}</span>{/if}
      {#if segment.role === 'link' && view.href !== null}
        <a class={ROLE_CLASS.link} href={view.href} data-testid="coach-last-raid-link">{segment.text}</a>
      {:else if view.stem !== null && segment.text.includes(view.stem)}
        {@const [before, after] = segment.text.split(view.stem)}
        <span class={ROLE_CLASS[segment.role]}
          >{before}<span class="inline-block max-w-[24ch] truncate align-bottom" title={view.stem}
            >{view.stem}</span
          >{after}</span
        >
      {:else}
        <span class={ROLE_CLASS[segment.role]}>{segment.text}</span>
      {/if}
    {/each}
  </p>
{/if}

<!-- web/src/components/sim/ServerRunRow.svelte -->
<!-- One row of "Your server runs" (experience spec 2026-10-24 section 3.3). Done rows are
     whole-row links to the result; running and failed rows are not links, and a failed row's
     one control is `Run again`. The row only draws: the parent owns the rerun call and says
     which of the five rerun states to show. -->
<script lang="ts">
  import { SECONDARY_BUTTON_FIXED } from '../../lib/planner/styles';
  import { simCopy } from '../../lib/sim/copy';
  import { headlineOf, kindOf } from '../../lib/sim/history';
  import { ageWords, elapsedSeconds, rowState } from '../../lib/sim/server-runs';
  import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
  import type { RerunStatus } from '../../lib/sim/server-runs-poll';
  import { classDisplayName, classOfSpec, specLabel } from '../../lib/sim/spec-label';
  import type { SimListRow } from '../../lib/sim/types';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';

  const CREST_SIZE = 32;
  const PREMIUM_PATH = '/premium';

  let {
    row,
    now,
    rerun,
    onrerun,
  }: { row: SimListRow; now: Date; rerun: RerunStatus; onrerun: (simId: string) => void } = $props();

  const state = $derived(rowState(row, now));
  const inFlight = $derived(state === 'queued' || state.startsWith('running'));
  const isDone = $derived(state === 'done' || state === 'done-unseen');
  const unseen = $derived(state === 'done-unseen');
  const spec = $derived(specLabel(row.spec));
  const kind = $derived(simCopy.kindLabel[kindOf(row)] ?? '');
  const title = $derived(row.title === '' ? spec : row.title);
  const meta = $derived(row.title === '' ? kind : `${spec} · ${kind}`);
  const className = $derived(classDisplayName(classOfSpec(row.spec)));

  const stampSource = $derived(inFlight || row.finished_at === undefined ? row.created_at : row.finished_at);
  const stampDate = $derived(new Date(stampSource));
  const age = $derived(ageWords(stampDate, now));
  const stampTitle = $derived(stampDate.toLocaleString('en-US', { dateStyle: 'medium', timeStyle: 'short' }));
  const ageLine = $derived(
    inFlight
      ? serverRunsCopy.startedLine(age)
      : state === 'error'
        ? serverRunsCopy.endedLine(age)
        : serverRunsCopy.finishedLine(age),
  );
  const mutedSuffix = $derived(
    state === 'queued' ? '' : serverRunsCopy.runningLine(row.estimate_seconds, elapsedSeconds(row, now)),
  );
  const rerunLabel = $derived(
    rerun === 'pending'
      ? serverRunsCopy.runAgainPending
      : rerun === 'started'
        ? serverRunsCopy.runAgainDone
        : serverRunsCopy.runAgain,
  );
  const rerunLocked = $derived(rerun === 'pending' || rerun === 'started');
</script>

<svelte:element
  this={isDone ? 'a' : 'div'}
  href={isDone ? `/sim/${row.sim_id}` : undefined}
  class="server-run-row border-line-soft relative grid grid-cols-[32px_minmax(0,1fr)] items-center gap-x-3.5 gap-y-1 border-b px-3.5 py-3 last:border-b-0 md:min-h-16 md:grid-cols-[32px_minmax(0,1.4fr)_minmax(0,1.2fr)_auto] md:py-2.5"
  class:server-run-link={isDone}
  class:server-run-unseen={unseen}
  data-testid={`server-run-${row.sim_id}`}
  data-state={state}
>
  <span class="row-span-3 self-start md:row-span-1 md:self-center">
    <ClassCrestRing characterClass={className} size={CREST_SIZE} />
  </span>

  <div class="col-start-2 flex min-w-0 flex-col gap-0.5 md:col-start-auto">
    <span class="text-strong flex min-w-0 items-center gap-2 text-[15px] font-semibold">
      {#if unseen}
        <span class="bg-gold inline-block h-2 w-2 shrink-0 rounded-full" aria-hidden="true"></span>
        <span class="sr-only">{serverRunsCopy.unreadSr}</span>
      {/if}
      <span class="server-run-title truncate" {title}>{title}</span>
    </span>
    <span class="text-muted truncate text-[13px]">{meta}</span>
  </div>

  <!-- Phone: the state and its age share one line, so the row stays within the stored
       px reserve (spec 3.7); desktop stacks them, right-aligned. -->
  <div
    class="col-start-2 flex min-w-0 flex-wrap items-baseline justify-between gap-x-3 gap-y-0.5 md:col-start-auto md:flex-col md:items-end md:justify-start md:text-right"
  >
    {#if inFlight}
      <span class="flex items-center gap-2 text-[14px]">
        <span
          class="bg-gold inline-block h-2 w-2 shrink-0 rounded-full"
          class:server-run-dot-pulse={state !== 'queued'}
          aria-hidden="true"
        ></span>
        <span class="text-strong">
          {state === 'queued' ? serverRunsCopy.queued : serverRunsCopy.running}<span class="text-muted"
            >{mutedSuffix}</span
          >
        </span>
      </span>
    {:else if isDone}
      <span class="text-strong numeric text-[15px]">{headlineOf(row)}</span>
    {:else}
      <span class="text-wipe text-[14px]">{serverRunsCopy.failedLine(row.failure_reason)}</span>
    {/if}
    <time class="text-muted numeric text-[12px]" datetime={stampSource} title={stampTitle}>{ageLine}</time>
    {#if rerun === 'failed'}
      <p class="text-wipe basis-full text-[13px]" role="alert">{serverRunsCopy.runAgainFailed}</p>
    {:else if rerun === 'lapsed'}
      <p class="text-wipe basis-full text-[13px]" role="alert">
        {serverRunsCopy.planLapsed}
        <a class="text-text inline-flex min-h-11 items-center underline md:min-h-0" href={PREMIUM_PATH}>
          {serverRunsCopy.seePlans}
        </a>
      </p>
    {/if}
  </div>

  {#if state === 'error'}
    <button
      type="button"
      class={`${SECONDARY_BUTTON_FIXED} col-start-2 w-full min-w-28 justify-center px-3 md:col-start-auto md:h-9 md:w-auto`}
      disabled={rerunLocked}
      aria-busy={rerun === 'pending' ? 'true' : undefined}
      onclick={() => onrerun(row.sim_id)}
      data-testid={`server-run-again-${row.sim_id}`}
    >
      {rerunLabel}
    </button>
  {:else if isDone}
    <svg
      class="server-run-arrow text-muted absolute top-1/2 right-3.5 -translate-y-1/2 md:static md:translate-y-0"
      width="16"
      height="16"
      viewBox="0 0 16 16"
      fill="none"
      stroke="currentColor"
      stroke-width="1.75"
      stroke-linecap="round"
      stroke-linejoin="round"
      aria-hidden="true"
    >
      <path d="M6 3l5 5-5 5" />
    </svg>
  {/if}
</svelte:element>

<style>
  .server-run-unseen {
    background: color-mix(in srgb, var(--color-gold) 8%, var(--color-raised));
  }
  .server-run-link:hover {
    background: var(--color-card-top);
  }
  .server-run-link:hover .server-run-title {
    text-decoration: underline;
    text-underline-offset: 4px;
  }
  .server-run-link:hover .server-run-arrow,
  .server-run-link:focus-visible .server-run-arrow {
    color: var(--color-gold);
  }
  .server-run-link:active {
    background: #171d2b;
  }
  .server-run-link:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  .server-run-dot-pulse {
    animation: server-run-pulse 1.6s ease-in-out infinite;
  }
  @keyframes server-run-pulse {
    50% {
      opacity: 0.35;
    }
  }
  @media (prefers-reduced-motion: reduce) {
    .server-run-dot-pulse {
      animation: none;
    }
  }
</style>

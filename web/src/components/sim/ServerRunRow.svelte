<!-- web/src/components/sim/ServerRunRow.svelte -->
<!-- One server run in the rail's Your sims (spec 2026-11-06-sim; experience spec 2026-10-24 3.6):
     an open ui/Row like every saved sim beside it, marked "On our servers" in its sub line with its
     age ("Finished 1 day ago"). A finished run is a link to its result and shows its headline; a
     run in flight says Starting or Running with its estimate; a failed run says why and offers
     Run again. An unopened result carries the gold dot. -->
<script lang="ts">
  import { PREMIUM_LINKS } from '../../lib/billing/premium-links';
  import { simCopy } from '../../lib/sim/copy';
  import { headlineOf, kindOf } from '../../lib/sim/history';
  import { ageWords, elapsedSeconds, rowState } from '../../lib/sim/server-runs';
  import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
  import type { RerunStatus } from '../../lib/sim/server-runs-poll';
  import { specLabel } from '../../lib/sim/spec-label';
  import type { SimListRow } from '../../lib/sim/types';
  import Button from '../ui/Button.svelte';
  import Row from '../ui/Row.svelte';

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
  const titleText = $derived(row.title === '' ? spec : row.title);
  const meta = $derived(row.title === '' ? kind : `${spec} · ${kind}`);

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

{#snippet title()}
  <span class="flex min-w-0 items-center gap-2">
    {#if unseen}
      <span class="bg-gold inline-block h-2 w-2 shrink-0 rounded-full" aria-hidden="true"></span>
      <span class="sr-only">{serverRunsCopy.unreadSr}</span>
    {/if}
    <span class="truncate">{titleText}</span>
  </span>
{/snippet}
{#snippet sub()}
  {meta === '' ? '' : `${meta} · `}{serverRunsCopy.serverMark} ·
  {#if inFlight}
    <span class="text-strong" data-testid={`server-run-state-${row.sim_id}`}
      ><span
        class="bg-gold mr-1 inline-block h-2 w-2 rounded-full align-middle"
        class:server-run-dot-pulse={state !== 'queued'}
        aria-hidden="true"
      ></span>{state === 'queued' ? serverRunsCopy.queued : serverRunsCopy.running}{mutedSuffix}</span
    > ·
  {:else if state === 'error'}
    <span class="text-wipe">{serverRunsCopy.failedLine(row.failure_reason)}</span> ·
  {/if}
  <time datetime={stampSource} title={stampTitle}>{ageLine}</time>
  {#if rerun === 'failed'}
    <span class="text-wipe block" role="alert">{serverRunsCopy.runAgainFailed}</span>
  {:else if rerun === 'lapsed'}
    <span class="text-wipe block" role="alert"
      >{serverRunsCopy.planLapsed}
      <a class="text-text underline" href={PREMIUM_LINKS.serverSims}>{serverRunsCopy.seePlans}</a></span
    >
  {/if}
{/snippet}
{#snippet trail()}
  {#if state === 'error'}
    <Button
      size="sm"
      disabled={rerunLocked}
      onclick={() => onrerun(row.sim_id)}
      testid={`server-run-again-${row.sim_id}`}>{rerunLabel}</Button
    >
  {:else if isDone}
    <span class="text-strong text-[16px] leading-5 font-bold tabular-nums">{headlineOf(row)}</span>
  {/if}
{/snippet}

<Row
  title={titleText}
  titleContent={title}
  subContent={sub}
  {trail}
  href={isDone ? `/sim/${row.sim_id}` : undefined}
  open
  testid={`server-run-${row.sim_id}`}
/>

<style>
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

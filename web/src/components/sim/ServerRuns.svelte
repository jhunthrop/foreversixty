<!-- web/src/components/sim/ServerRuns.svelte -->
<!-- "Your server runs" (experience spec 2026-10-24 section 3.2): the payer's newest five
     server runs, first under the page heading. Rendered by SimView only when the account has
     server sims. The parent reads and polls the list; this draws it, and owns the Run again
     calls because each row's button state belongs to the list, not to the page. -->
<script lang="ts">
  import Card from '../ui/Card.svelte';
  import { rerunSim } from '../../lib/sim/api';
  import { visibleRows } from '../../lib/sim/server-runs';
  import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
  import { rerunOutcome } from '../../lib/sim/server-runs-rerun';
  import { writeReserveHeight, type RerunStatus } from '../../lib/sim/server-runs-poll';
  import type { SimListRow } from '../../lib/sim/types';
  import EmptyState from '../ui/EmptyState.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import ServerRunRow from './ServerRunRow.svelte';

  let {
    rows,
    failed,
    now,
    onretry,
    onstarted,
  }: {
    rows: SimListRow[] | null;
    /** The first read failed and nothing is held to show. */
    failed: boolean;
    now: Date;
    onretry: () => void;
    /** A failed run was started again: the parent re-reads the list for the new row. */
    onstarted: () => void;
  } = $props();

  const SKELETON_ROWS = 3;
  let rerunStatus = $state<Record<string, RerunStatus>>({});

  const shown = $derived(rows === null ? null : visibleRows(rows));

  let panel = $state<HTMLElement | undefined>();

  // The panel's height without the stored reserve holding it open, so the empty state (which
  // carries the reserve class) can shrink the reserve instead of writing it back unchanged.
  function naturalHeight(element: HTMLElement): number {
    const held = element.style.minHeight;
    element.style.minHeight = '0';
    const height = element.offsetHeight;
    element.style.minHeight = held;
    return height;
  }

  // The panel's rendered height is next visit's reserve (spec 3.7): measured once rows or the
  // empty state are painted, never while loading or failed, and again whenever a row's state
  // or a Run again status changes (a poll flipping running to done, a failure line appearing).
  $effect(() => {
    void rerunStatus;
    if (panel !== undefined && shown !== null) {
      void shown.rows.map((row) => row.state);
      writeReserveHeight(naturalHeight(panel));
    }
  });

  function setStatus(simId: string, status: RerunStatus): void {
    rerunStatus = { ...rerunStatus, [simId]: status };
  }

  async function rerun(simId: string): Promise<void> {
    setStatus(simId, 'pending');
    const outcome = await rerunOutcome(() => rerunSim(simId));
    setStatus(simId, outcome);
    if (outcome === 'started') onstarted();
  }
</script>

<Card
  glyph="server"
  compact
  body="none"
  title={serverRunsCopy.heading}
  id="server-runs"
  testid="server-runs"
  class="mx-[18px] md:mx-0"
>
  {#if failed}
    <div class="border-line-soft server-runs-rows-reserve flex items-center border-t px-3.5">
      <LoadError message={serverRunsCopy.loadFailed} onRetry={onretry} testid="server-runs-error" />
    </div>
  {:else if shown === null}
    <div
      class="border-line-soft server-runs-rows-reserve overflow-hidden border-t"
      aria-busy="true"
      data-testid="server-runs-loading"
    >
      <span class="sr-only">{serverRunsCopy.loadingSr}</span>
      {#each Array.from({ length: SKELETON_ROWS }, (_, index) => index) as index (index)}
        <div
          class="border-line-soft flex items-center gap-3.5 border-b px-3.5 py-3 last:border-b-0"
          aria-hidden="true"
        >
          <span class="skeleton-block h-8 w-8 shrink-0 rounded-full"></span>
          <span class="flex min-w-0 flex-1 flex-col gap-1.5">
            <span class="skeleton-block h-[15px] w-[180px] max-w-full"></span>
            <span class="skeleton-block h-3 w-[120px] max-w-full"></span>
          </span>
          <span class="skeleton-block h-3.5 w-[140px] max-w-[30%] shrink-0"></span>
        </div>
      {/each}
    </div>
  {:else if shown.rows.length === 0}
    <div
      bind:this={panel}
      class="border-line-soft server-runs-rows-reserve flex items-center border-t px-3.5"
    >
      <EmptyState message={serverRunsCopy.empty} testid="server-runs-empty" />
    </div>
  {:else}
    <div bind:this={panel} class="border-line-soft overflow-hidden border-t" data-testid="server-runs-list">
      {#each shown.rows as row (row.sim_id)}
        <ServerRunRow
          {row}
          {now}
          rerun={rerunStatus[row.sim_id] ?? 'idle'}
          onrerun={(id) => void rerun(id)}
        />
      {/each}
      {#if shown.hasMore}
        <p class="text-muted px-3.5 py-2 text-[12px]" data-testid="server-runs-overflow">
          {serverRunsCopy.overflow}
        </p>
      {/if}
    </div>
  {/if}
</Card>

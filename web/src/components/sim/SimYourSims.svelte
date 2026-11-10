<!-- web/src/components/sim/SimYourSims.svelte -->
<!-- "Your sims" in the /sim rail (spec 2026-11-06-sim §4.3, §4.5, §5 "Rail rows"): one list of open
     Rows, newest first. A payer's server runs are in it too, each marked "On our servers" with its
     state and age (ServerRunRow; experience spec 2026-10-24 3.6), no panel of their own; every
     other saved sim is its title, the spec and day, and a run's DPS at the right. "All ›" opens the
     whole list on the account page. Three skeleton rows while loading, one line when empty, one
     line and Try again when a read failed. -->
<script lang="ts">
  import { rerunSim } from '../../lib/sim/api';
  import { KIND_TITLES, simCopy } from '../../lib/sim/copy';
  import { dayMonth } from '../../lib/sim/head-meta';
  import { kindOf, runDeltas, titleOf } from '../../lib/sim/history';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { grouped, oneDecimal } from '../../lib/sim/result-copy';
  import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
  import type { RerunStatus } from '../../lib/sim/server-runs-poll';
  import { rerunOutcome } from '../../lib/sim/server-runs-rerun';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import type { SimListRow } from '../../lib/sim/types';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import ServerRunRow from './ServerRunRow.svelte';

  let {
    rows,
    error,
    onretry,
    serverRows = null,
    serverFailed = false,
    now = new Date(),
    onserverretry,
    onstarted,
    pageSpec = '',
  }: {
    rows: readonly SimListRow[] | null;
    error: string | null;
    onretry: () => void;
    /** A payer's server runs (null: not a payer, or not read yet). */
    serverRows?: readonly SimListRow[] | null;
    serverFailed?: boolean;
    now?: Date;
    onserverretry?: () => void;
    /** A failed server run was started again. */
    onstarted?: () => void;
    /** The spec the page is about: a run of it says only its date. */
    pageSpec?: string;
  } = $props();

  /** The rail shows the newest few; the account page keeps them all. */
  const SHOWN = 5;
  const SKELETON_ROWS = 3;

  const serverIds = $derived(new Set((serverRows ?? []).map((row) => row.sim_id)));
  /** One list: the server runs (with their state) and every other saved sim, newest first. */
  const shown = $derived(
    [...(serverRows ?? []), ...(rows ?? []).filter((row) => !serverIds.has(row.sim_id))]
      .sort((a, b) => b.created_at.localeCompare(a.created_at))
      .slice(0, SHOWN),
  );

  /** Each run against the run of its setup before it, from every row read, not only the shown five. */
  const deltas = $derived(runDeltas(rows ?? []));
  const signed = (delta: number): string => `${delta >= 0 ? '+' : '−'}${oneDecimal(Math.abs(delta))}`;

  let rerunStatus = $state<Record<string, RerunStatus>>({});
  async function rerun(simId: string): Promise<void> {
    rerunStatus = { ...rerunStatus, [simId]: 'pending' };
    const outcome = await rerunOutcome(() => rerunSim(simId));
    rerunStatus = { ...rerunStatus, [simId]: outcome };
    if (outcome === 'started') onstarted?.();
  }
</script>

<!-- The header's "finished" line links here (/sim#server-runs). -->
<div id="server-runs">
  <Section
    panel
    title={copy.sections.yourSims}
    action={{ label: copy.sections.yourSimsAll, href: '/account#sims' }}
    testid="sim-history"
  >
    {#if error !== null}
      <p class="text-muted py-3 text-[14px]" role="alert" data-testid="sim-history-error">
        Your sims did not load. <button type="button" class="text-gold font-semibold" onclick={onretry}
          >Try again</button
        >
      </p>
    {:else if rows === null}
      <div aria-busy="true">
        {#each Array.from({ length: SKELETON_ROWS }, (_, index) => index) as index (index)}
          <div class="py-[9px]"><Skeleton lines={2} /></div>
        {/each}
      </div>
    {:else if shown.length === 0}
      <p class="text-muted py-3 text-[14px]" data-testid="sim-history-empty">{simCopy.historyEmpty}</p>
    {:else}
      <div data-testid="sim-history-list">
        {#each shown as row (row.sim_id)}
          {#if serverIds.has(row.sim_id)}
            <ServerRunRow
              {row}
              {now}
              {pageSpec}
              rerun={rerunStatus[row.sim_id] ?? 'idle'}
              onrerun={(id) => void rerun(id)}
            />
          {:else}
            {@const kind = kindOf(row)}
            <!-- An untitled run of the page's own spec is named by its day alone (h-audit 2026-11-09). -->
            {@const ownRun = kind === 'run' && row.title === '' && row.spec === pageSpec}
            <Row
              title={ownRun ? dayMonth(row.created_at) : titleOf(row)}
              sub={ownRun
                ? undefined
                : [
                    kind === 'run'
                      ? row.spec === pageSpec
                        ? ''
                        : specDisplayName(row.spec)
                      : ((KIND_TITLES as Record<string, string>)[kind] ?? kind),
                    dayMonth(row.created_at),
                  ]
                    .filter((part) => part !== null && part !== '')
                    .join(' · ')}
              href={`/sim/${row.sim_id}`}
              testid={`sim-history-${row.sim_id}`}
              open
            >
              {#snippet trail()}
                {#if kind === 'run' && row.dps > 0}<Figure
                    value={grouped(row.dps)}
                    unit="DPS"
                    size="sm"
                    align="end"
                    caption={deltas.has(row.sim_id) ? signed(deltas.get(row.sim_id) ?? 0) : undefined}
                  />{/if}
              {/snippet}
            </Row>
          {/if}
        {/each}
      </div>
    {/if}
    {#if serverFailed && onserverretry}
      <p class="text-muted py-3 text-[14px]" role="alert" data-testid="server-runs-error">
        {serverRunsCopy.loadFailed}
        <button type="button" class="text-gold font-semibold" onclick={onserverretry}>Try again</button>
      </p>
    {/if}
  </Section>
</div>

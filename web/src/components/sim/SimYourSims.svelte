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
  import { kindOf, titleOf } from '../../lib/sim/history';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { grouped } from '../../lib/sim/result-copy';
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
              rerun={rerunStatus[row.sim_id] ?? 'idle'}
              onrerun={(id) => void rerun(id)}
            />
          {:else}
            {@const kind = kindOf(row)}
            <Row
              title={titleOf(row)}
              sub={[
                kind === 'run'
                  ? specDisplayName(row.spec)
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

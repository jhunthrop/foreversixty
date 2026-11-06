<!-- web/src/components/sim/SimYourSims.svelte -->
<!-- "Your sims" in the /sim rail (spec 2026-11-06-sim §4.3, §4.5, §5 "Rail rows"): the signed-in
     player's newest saved sims as open link rows, each its title, the kind and date, and a run's
     DPS at the right; "All ›" opens the full list on the account page. Three skeleton rows while
     loading, one line when empty, one line and Try again when the read failed. -->
<script lang="ts">
  import { KIND_TITLES, simCopy } from '../../lib/sim/copy';
  import { dayMonth } from '../../lib/sim/head-meta';
  import { kindOf, titleOf } from '../../lib/sim/history';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import { grouped } from '../../lib/sim/result-copy';
  import type { SimListRow } from '../../lib/sim/types';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Skeleton from '../ui/Skeleton.svelte';

  let {
    rows,
    error,
    onretry,
  }: {
    rows: readonly SimListRow[] | null;
    error: string | null;
    onretry: () => void;
  } = $props();

  /** The rail shows the newest few; the account page keeps them all. */
  const SHOWN = 5;
  const SKELETON_ROWS = 3;
  const shown = $derived((rows ?? []).slice(0, SHOWN));
</script>

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
    <div>
      {#each shown as row (row.sim_id)}
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
            {#if kind === 'run' && row.dps > 0}<Figure value={grouped(row.dps)} unit="DPS" size="sm" />{/if}
          {/snippet}
        </Row>
      {/each}
    </div>
  {/if}
</Section>

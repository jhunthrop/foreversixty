<!-- web/src/components/sim/tools/WeightsShell.svelte -->
<!-- /sim/weights' server render (sim/weights.astro), the island's no-character state drawn
     without a store: the source switcher, the run line with its Run held off, the stats to weigh
     and the default spec's published weights. The island mounts the same parts over it
     (ToolsView: SourceSwitcher, then StatWeights with `defaultWeights`), so the table a search
     engine reads is the one a player sees, figure for figure. -->
<script lang="ts">
  import { BULK_PRECISIONS, type Precision } from '../../../lib/sim/bulk-types';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { PrecisionId } from '../../../lib/sim/precision';
  import type { PublishedWeights } from '../../../lib/sim/published-weights';
  import { defaultSettings } from '../../../lib/sim/settings';
  import { pickableStatsForSpec } from '../../../lib/sim/weights';
  import Page from '../../ui/Page.svelte';
  import SimRunBar from '../SimRunBar.svelte';
  import SourceSwitcher from '../SourceSwitcher.svelte';
  import WeightsPicks from './WeightsPicks.svelte';
  import WeightsPublishedSection from './WeightsPublishedSection.svelte';

  let { published }: { published: PublishedWeights } = $props();

  const nothing = (): void => {};
  /** The bulk store's opening precision (bulk-store.svelte.ts). */
  const OPENING_PRECISION: PrecisionId = 'fast';
  const PRECISION_LABELS: Record<Precision, string> = {
    fast: bulkCopy.precisionFast,
    normal: bulkCopy.precisionNormal,
    high: bulkCopy.precisionHigh,
  };
  const precisionOptions = BULK_PRECISIONS.map((id) => ({ value: id, label: PRECISION_LABELS[id] }));
</script>

<div data-testid="sim-tools-view">
  <!-- Held busy (every field and Load disabled) until the island mounts: a paste typed into the
       server's form would be dropped when the island replaces it. -->
  <SourceSwitcher
    busy={true}
    message={null}
    signedIn={false}
    hasCharacters={false}
    onaddon={nothing}
    onbuild={nothing}
    onfight={nothing}
    onsignin={nothing}
  />
  <Page layout="ledger" panels testid="sim-stat-weights">
    {#snippet bar()}
      <!-- BulkRunBar's opening state: the store's default settings and Fast, Run held off. -->
      <SimRunBar
        settings={defaultSettings()}
        precisionId={OPENING_PRECISION}
        spec=""
        level={60}
        primary="disabled"
        primaryLabel={bulkCopy.runBulk}
        primaryTestid="sim-run-bulk"
        testid="sim-run-bulk-bar"
        busy={false}
        {precisionOptions}
        onsettings={nothing}
        onprecision={nothing}
        onprimary={nothing}
        onmore={nothing}
      />
    {/snippet}
    <WeightsPicks pickable={pickableStatsForSpec(published.spec)} picked={[]} reference="" />
    <WeightsPublishedSection {published} isDefault />
  </Page>
</div>

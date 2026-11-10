<!-- web/src/components/sim/tools/WeightsPublishedSection.svelte -->
<!-- The published weights on /sim/weights' framed panel (lib/sim/published-weights.ts): the spec
     they belong to, the set they were measured on as the link, and the scale-factor table. The
     island draws it for the loaded character's band, or for the default spec before one is
     loaded; sim/weights.astro server-renders the default one through WeightsShell, so the two
     agree on every figure. -->
<script lang="ts">
  import { MAX_BAND } from '../../../lib/bis/band-levels';
  import { bulkCopy } from '../../../lib/sim/copy';
  import type { PublishedWeights } from '../../../lib/sim/published-weights';
  import { specLabel } from '../../../lib/sim/spec-label';
  import Section from '../../ui/Section.svelte';
  import WeightsScaleTable from '../WeightsScaleTable.svelte';

  let {
    published,
    isDefault = false,
  }: {
    published: PublishedWeights;
    /** No character is loaded: the line under the head says whose figures these are. */
    isDefault?: boolean;
  } = $props();
</script>

<Section
  panel
  frame="plain"
  title={bulkCopy.weightsPublishedTitle}
  meta={specLabel(published.spec)}
  action={{ label: published.label, href: published.href }}
  testid="sim-weights-published"
>
  {#if isDefault}
    <p class="text-small text-muted m-0 pb-3" data-testid="sim-weights-default-note">
      {bulkCopy.weightsDefaultNote(MAX_BAND)}
    </p>
  {/if}
  <WeightsScaleTable view={published.view} />
</Section>

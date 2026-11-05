<!-- web/src/components/sim/SimResourceCard.svelte -->
<!-- The lane contract's placeholder (design/specs/2026-11-04-sim-results-contract.md): the owning
     lane replaces this file. -->
<script lang="ts">
  import type { ActionNames } from '../../lib/sim/action-names';
  import type { SimResult } from '../../lib/sim/types';

  const props: { result: SimResult; names: ActionNames | null } = $props();
</script>

<div data-testid="resources-pending" data-spec={props.result.request.spec}></div>

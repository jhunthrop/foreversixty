<!-- web/src/components/sim/SimView.svelte -->
<!-- The simulator island's root. It owns the URL state and the page frame; every section
     below it is a component that takes data and raises events, so this file stays a
     composition and never a place where a layout decision hides. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import { simCopy } from '../../lib/sim/copy';
  import { defaultSimState, parseSimState, simIdFrom, type SimState } from '../../lib/sim/url';
  import { ENGINE_VERSION, engineLabel } from '../../lib/sim/version';
  import type { SimResult } from '../../lib/sim/types';

  let { simId = '', inlineResult = null }: { simId?: string; inlineResult?: SimResult | null } = $props();

  let state = $state<SimState>(defaultSimState());
  // untrack: simId and inlineResult are one-shot bootstrap values the shell inlined, not
  // bindings. Reading a prop while initialising state is exactly what Svelte's
  // state_referenced_locally warning is about, and ReportView.svelte does the same for the
  // same reason. Without it the island build prints a warning on every build.
  let openSimId = $state(untrack(() => simId));

  $effect(() => {
    state = parseSimState(window.location.search);
    if (openSimId === '') openSimId = simIdFrom(window.location.pathname);
  });

  const hasCharacter = $derived(state.ref !== '' || inlineResult !== null);
</script>

<div class="flex flex-col gap-[22px] md:gap-8" data-testid="sim-view">
  <div class="flex flex-wrap items-baseline gap-x-4 gap-y-1 px-[18px] md:px-0">
    <h1 class="section-title">Simulator</h1>
    <a
      class="tabular text-muted ml-auto font-mono text-[12px]"
      href="/sim/specs"
      data-testid="sim-engine-version">{engineLabel(ENGINE_VERSION)}</a
    >
  </div>

  {#if !hasCharacter}
    <p class="text-muted px-[18px] text-[14px] md:px-0" data-testid="sim-empty">
      {simCopy.emptyPrompt}
    </p>
  {/if}
</div>

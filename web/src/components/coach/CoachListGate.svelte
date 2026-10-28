<!-- web/src/components/coach/CoachListGate.svelte -->
<!-- Holds the combatant list until /v1/me resolves or 150 ms pass, whichever is first (spec 7), so
     the viewer's own row renders with its strip slot already in place and no other row moves. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { LIST_GATE_MS, listGate, listGateOpen } from '../../lib/coach/list-gate';
  import { createCoachViewerState } from '../../lib/coach/viewer-state.svelte';

  let { children, apiBase = undefined }: { children: Snippet; apiBase?: string } = $props();

  // svelte-ignore state_referenced_locally
  const viewerState = createCoachViewerState(apiBase);
  let timedOut = $state(false);
  $effect(() => {
    const timer = setTimeout(() => (timedOut = true), LIST_GATE_MS);
    return () => clearTimeout(timer);
  });

  const open = $derived(listGateOpen(listGate(viewerState.resolved, timedOut)));
</script>

{#if open}{@render children()}{/if}

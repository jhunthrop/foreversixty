<!-- web/src/components/planner/PlannerToolbar.svelte -->
<!-- The build's own controls: a read-only build shows only the Fork refusal; an editable one
     shows Reset (with its own confirm step) beside the Share panel. Split out of
     Planner.svelte (design loop, planner round) so that file stays under the project's
     file-size guideline as the responsive rail layout grew it. -->
<script lang="ts">
  import { plannerCopy } from '../../lib/planner/copy';
  import type { LiveDps } from '../../lib/planner/live-dps.svelte';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import { SECONDARY_BUTTON } from '../../lib/planner/styles';
  import SharePanel from './SharePanel.svelte';

  let {
    store,
    live,
    confirmingReset = $bindable(),
  }: {
    store: PlannerStore;
    live: LiveDps;
    /** Owned by the caller too: a class switch clears the build on its own (store.svelte.ts)
     *  and has to close a confirm left open across it, the same reset the Keep/Clear buttons
     *  below perform locally. */
    confirmingReset: boolean;
  } = $props();
</script>

{#if store.readOnly}
  <!-- A build opened from a share link. Every edit is refused until Fork, so the
     toolbar says so up front rather than leaving the refusal message to explain it
     after the first click. Reset and Share are gone with it: there is nothing of
     one's own to clear, and re-sharing someone else's build under a new id is the
     one thing Fork is for. -->
  <p class="text-muted text-[13px]">This build was shared as a link. Fork it to spend points of your own.</p>
  <button
    type="button"
    class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
    onclick={() => store.fork()}
  >
    Fork
  </button>
{:else}
  <!-- Nested rather than a third arm of the branch above, so `readOnly` is asked once:
     Reset and Share belong to the same half of that decision, and SharePanel has to
     sit outside the confirm to survive it -- it holds the title being typed and the
     link of the last save, and re-mounting it when the confirm opens would throw
     both away. -->
  <section
    class="border-line bg-raised rounded-panel flex w-full flex-col gap-3 border p-4"
    data-testid="planner-share-section"
  >
    <header class="flex flex-wrap items-center justify-between gap-3">
      <h2 class="section-title text-[15px]">{plannerCopy.shareTitle}</h2>
      <div class="flex flex-wrap items-center gap-3">
        {#if confirmingReset}
          <span class="text-muted text-[13px]">Clear every point in this build?</span>
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm-strong text-gold px-4"
            onclick={() => {
              store.reset();
              confirmingReset = false;
            }}
          >
            Clear all points
          </button>
          <!-- Reset leaves the DOM the moment it is pressed, so the keyboard lands on the
               question it just asked rather than back at the top of the document. It lands
               on the safe answer: a second Enter pressed out of habit keeps the build rather
               than clearing it, which is the only reason the second step exists. -->
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
            {@attach (node) => node.focus()}
            onclick={() => (confirmingReset = false)}
          >
            Keep the build
          </button>
        {:else}
          <button
            type="button"
            class="{SECONDARY_BUTTON} border-line-warm text-text px-4"
            onclick={() => (confirmingReset = true)}
          >
            Reset…
          </button>
        {/if}
      </div>
    </header>
    <SharePanel {store} {live} />
  </section>
{/if}

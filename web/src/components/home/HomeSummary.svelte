<!-- web/src/components/home/HomeSummary.svelte -->
<!-- The signed-in home's one next-step line (home spec 2026-11-06 H2), the page's `summary` region:
     a Notice under the head, or its skeleton while the session loads. -->
<script lang="ts">
  import { REGION } from '../../lib/ui/regions';
  import Notice from '../ui/Notice.svelte';

  let {
    text,
    action,
    loading = false,
  }: { text?: string; action?: { label: string; href: string } | null; loading?: boolean } = $props();
</script>

<div data-region={REGION.summary} class="home-summary" data-testid="home-next-step">
  {#if loading}
    <div class="pt-[14px] pb-1">
      <span class="skeleton-block block h-5 w-[520px] max-w-full" aria-hidden="true"></span>
    </div>
  {:else if text}
    <Notice {text} action={action ?? undefined} />
  {/if}
</div>

<style>
  .home-summary {
    padding-inline: var(--page-inset);
  }
</style>

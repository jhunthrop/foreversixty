<!-- web/src/components/logs/LogsPanel.svelte -->
<!-- One titled panel of the logs landing: the same shell as Panel.astro (title and a muted aside
     over a bordered raised box), for the Svelte island that composes the page. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    title,
    aside = '',
    id,
    children,
  }: { title: string; aside?: string; id?: string; children: Snippet } = $props();
</script>

<section {id} class="flex scroll-mt-6 flex-col gap-4">
  <div class="flex items-baseline justify-between gap-3">
    <h2 class="section-title">{title}</h2>
    {#if aside !== ''}<span class="text-muted text-[13px]">{aside}</span>{/if}
  </div>
  <div class="bg-raised border-line rounded-panel flex flex-col border">
    {@render children()}
  </div>
</section>

<!-- web/src/components/logs/LogsPanel.svelte -->
<!-- One titled panel of the logs landing, as the mock draws it (gen_logs.py `panel`): a bordered
     raised box whose title row (a 15 px display title and a muted aside) sits inside it, over the
     panel's own content. The box owns the padding; the content brings none. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    title,
    aside = '',
    id,
    children,
  }: { title: string; aside?: string; id?: string; children: Snippet } = $props();
</script>

<section {id} class="bg-raised border-line rounded-panel flex scroll-mt-6 flex-col gap-3 border p-4 md:px-5">
  <div class="flex items-baseline justify-between gap-3">
    <h2 class="section-title">{title}</h2>
    {#if aside !== ''}<span class="text-muted text-[12px]">{aside}</span>{/if}
  </div>
  {@render children()}
</section>

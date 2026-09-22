<!-- web/src/components/account/AccountPanel.svelte -->
<!-- Svelte twin of Panel.astro (bg-raised border border-line rounded-panel), for /account's
     two grouped panels (spec 2026-09-22 §2.1): Devices+You, and Guilds+Billing+Reports.
     Panel.astro is Astro-only (no <script>), so it cannot mount inside a Svelte island, and
     this lane does not own components/ui/ to add a shared one there -- this repeats the box
     treatment instead. Sub-sections inside are the caller's own <section> elements; the
     caller adds `border-line-soft border-t pt-4` to every one after the first for the
     divider the spec asks for. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let { testid, children }: { testid?: string; children: Snippet } = $props();
</script>

<div
  class="bg-raised border-line rounded-panel flex flex-col gap-4 border px-4 py-4 md:px-6 md:py-5"
  data-testid={testid}
>
  {@render children()}
</div>

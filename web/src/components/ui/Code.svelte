<!-- web/src/components/ui/Code.svelte -->
<!-- Literal text a player types or matches (setup spec 2026-11-07 section 3A): slash commands,
     /combatlog, pairing and connect codes, the export field. The site's one use of a mono face:
     ui-monospace 600 14/20 (`big`: 24/30, tracking .06em), --color-strong. Selectable, never a
     button. -->
<script lang="ts">
  let { text, big = false, testid }: { text: string; big?: boolean; testid?: string } = $props();
</script>

<code class="ui-code" data-big={big ? 'true' : undefined} data-testid={testid}>{text}</code>

<style>
  .ui-code {
    font:
      600 14px/20px ui-monospace,
      'SF Mono',
      Menlo,
      Consolas,
      monospace;
    color: var(--color-strong);
    letter-spacing: 0.02em;
    font-variant-numeric: tabular-nums;
  }
  .ui-code[data-big='true'] {
    font-size: 24px;
    line-height: 30px;
    letter-spacing: 0.06em;
  }
</style>

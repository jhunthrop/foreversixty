<!-- web/src/components/ui/StatTile.svelte -->
<!-- A headline figure (polish spec 4.2): the figure in the heading face at 28 px over a micro
     label, on a quiet tile. Replaces a mono headline number. The label names the figure for a
     screen reader; the figure carries tabular numerals. -->
<script lang="ts">
  let { value, label, testid }: { value: string; label: string; testid?: string } = $props();
</script>

<div
  class="rounded-control border-line-soft flex flex-col gap-1 border bg-white/[.03] px-[14px] py-3"
  data-testid={testid}
>
  <span class="figure-stat text-strong">{value}</span>
  <span class="text-muted text-[11px] font-bold tracking-[.08em] uppercase">{label}</span>
</div>

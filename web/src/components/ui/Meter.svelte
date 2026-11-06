<!-- web/src/components/ui/Meter.svelte -->
<!-- A bar (web layout spec 2026-11-05 §3A.2): an 8 px --color-line-soft track and a fill in gold,
     kill or muted, for weights, tier bars, progression and readiness. `segments` draws that many
     equal cells with 2 px gaps (a role mix). Decorative by default: the figure beside it carries the
     number; pass `label` to make it a meter for assistive tech. -->
<script lang="ts">
  let {
    value,
    tone = 'gold',
    segments,
    label,
    testid,
  }: {
    /** 0 to 1; clamped. */
    value: number;
    tone?: 'gold' | 'kill' | 'muted';
    segments?: number;
    label?: string;
    testid?: string;
  } = $props();

  const fraction = $derived(Math.min(1, Math.max(0, Number.isFinite(value) ? value : 0)));
  const FILL = { gold: 'bg-gold', kill: 'bg-kill', muted: 'bg-muted/50' } as const;
  const lit = $derived(segments ? Math.round(fraction * segments) : 0);
</script>

<span
  class={[
    'ui-meter flex h-2 min-w-[60px] overflow-hidden rounded-[3px]',
    segments ? 'gap-[2px]' : 'bg-line-soft',
  ]}
  role={label ? 'meter' : undefined}
  aria-label={label}
  aria-valuenow={label ? Math.round(fraction * 100) : undefined}
  aria-valuemin={label ? 0 : undefined}
  aria-valuemax={label ? 100 : undefined}
  aria-hidden={label ? undefined : 'true'}
  data-testid={testid}
>
  {#if segments}
    {#each Array.from({ length: segments }, (_, index) => index) as index (index)}
      <i class={['h-full flex-1', index < lit ? FILL[tone] : 'bg-line-soft']}></i>
    {/each}
  {:else}
    <i class={['h-full rounded-[3px]', FILL[tone]]} style:width="{(fraction * 100).toFixed(1)}%"></i>
  {/if}
</span>

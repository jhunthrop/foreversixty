<!-- web/src/components/RatingTrend.svelte -->
<!-- A small dedicated sparkline for a character's rating trend: a handful of unevenly
     -spaced points across weeks, not TimeChart.svelte's dense per-second in-fight series
     with a drag-brush -- see docs/superpowers/plans/2026-09-21-rating-web.md's Ruling 6 for
     why this is its own tiny component rather than a TimeChart reuse. No charting library:
     one inline <svg>, one <polyline>, matching the site's existing no-dependency-chart
     precedent (ResourceGraphs.svelte, TimeChart.svelte). -->
<script lang="ts">
  let { points }: { points: { fought_at: string; overall: number }[] } = $props();

  const WIDTH = 240;
  const HEIGHT = 48;
  const PAD = 4;

  const ordered = $derived([...points].sort((a, b) => a.fought_at.localeCompare(b.fought_at)));

  const coords = $derived.by(() => {
    if (ordered.length < 2) return [];
    const span = ordered.length - 1;
    return ordered.map((point, index) => {
      const x = PAD + (index / span) * (WIDTH - PAD * 2);
      const y = PAD + (1 - point.overall / 100) * (HEIGHT - PAD * 2);
      return `${x.toFixed(1)},${y.toFixed(1)}`;
    });
  });
</script>

{#if coords.length > 0}
  <svg
    viewBox={`0 0 ${WIDTH} ${HEIGHT}`}
    width={WIDTH}
    height={HEIGHT}
    role="img"
    aria-label={`Rating trend over the last ${ordered.length} fights, from ${Math.round(ordered[0].overall)} to ${Math.round(ordered[ordered.length - 1].overall)}`}
    data-testid="rating-trend"
  >
    <polyline
      points={coords.join(' ')}
      fill="none"
      stroke="var(--color-gold)"
      stroke-width="2"
      stroke-linejoin="round"
      stroke-linecap="round"
    />
  </svg>
{/if}

<!-- web/src/components/Sparkline.svelte -->
<!-- The profile's sparkline figure (Task L8.4 ruling 3, fix round 2 items 1 and 2, fix round 3
     item 12; shared with the Performance Rating card in L8.5 ruling 1). Plain model in
     (lib/score/sparkline.ts): a gold line stretched over the whole box at a fixed height -- 40px,
     48px from 1024px -- with an 8px end dot on the right edge that may overhang it; under it the
     caption ("Up 330 since 10 Nov") at 12px in the text colour leading the date row and the
     newest date ending it, mono muted; and an image named for a screen reader with the table
     beside it. The card says how wide the figure is (`class`); the ids name its parts. -->
<script lang="ts">
  import { SPARKLINE_HEIGHT, SPARKLINE_WIDTH, type Sparkline } from '../lib/score/sparkline';

  let {
    line,
    class: className = '',
    testids,
    tableCaption,
    dateHeader,
    valueHeader,
  }: {
    line: Sparkline;
    /** The figure's own classes: its width and its place in the card. */
    class?: string;
    /** Each part's data-testid (the score card keeps the ones its tests already read). */
    testids: {
      figure: string;
      plot: string;
      svg: string;
      last: string;
      dates: string;
      caption: string;
      lastLabel: string;
      table: string;
    };
    tableCaption: string;
    dateHeader: string;
    valueHeader: string;
  } = $props();
</script>

<figure class={`m-0 flex flex-col gap-1 ${className}`} data-testid={testids.figure}>
  <div class="relative h-10 lg:h-12" data-testid={testids.plot}>
    <svg
      class="absolute inset-0 block h-full w-full overflow-visible"
      viewBox={`0 0 ${SPARKLINE_WIDTH} ${SPARKLINE_HEIGHT}`}
      preserveAspectRatio="none"
      role="img"
      aria-label={line.ariaLabel}
      data-testid={testids.svg}
    >
      <path
        d={line.path}
        fill="none"
        stroke="var(--color-gold)"
        stroke-width="2"
        stroke-linejoin="round"
        stroke-linecap="round"
        vector-effect="non-scaling-stroke"
      />
    </svg>
    <span
      class="absolute size-2 -translate-x-1/2 -translate-y-1/2 rounded-full"
      style={`left: ${line.last.x}%; top: ${line.last.y}%; background-color: ${line.last.color ?? 'var(--color-gold)'}`}
      aria-hidden="true"
      data-testid={testids.last}
    ></span>
  </div>
  <div class="flex items-baseline justify-between gap-3" aria-hidden="true" data-testid={testids.dates}>
    <span class="text-text text-[12px]" data-testid={testids.caption}>{line.caption}</span>
    <span class="text-muted numeric text-[12px]" data-testid={testids.lastLabel}>{line.lastLabel}</span>
  </div>
  <table class="sr-only" data-testid={testids.table}>
    <caption>{tableCaption}</caption>
    <thead>
      <tr><th scope="col">{dateHeader}</th><th scope="col">{valueHeader}</th></tr>
    </thead>
    <tbody>
      {#each line.rows as row, index (index)}
        <tr><td>{row.date}</td><td>{row.value}</td></tr>
      {/each}
    </tbody>
  </table>
</figure>

<!-- web/src/components/report/Glossary.svelte -->
<!-- The words the tables use, in one place a phone can open. Every column heading carries
     the same text as a hover title, which a touch screen never shows; this is the tap
     version, a native disclosure so it costs no script and reads to a screen reader as
     what it is. -->
<script lang="ts">
  const TERMS: { term: string; meaning: string }[] = [
    {
      term: 'Parse',
      meaning:
        'Where a kill sits among every ranked kill of the same boss by the same spec, on this ruleset, as a percentile: 100 is the best, 0 the lowest. A wipe is not ranked, and a dash means nothing of this spec has been ranked on this boss yet.',
    },
    {
      term: 'Damage · DPS',
      meaning: 'Damage done in the window, and the same divided by the window’s length in seconds.',
    },
    {
      term: 'Healing · HPS',
      meaning:
        'Effective healing done: healing that landed on a missing health bar. Overhealing is the rest.',
    },
    { term: 'Taken · DTPS', meaning: 'Damage taken, and the same per second. A tank’s headline figure.' },
    {
      term: 'Active',
      meaning:
        'The share of the window the player spent casting or attacking. Time spent dead is not active.',
    },
    {
      term: 'Kill · Wipe · pull 1 of 2',
      meaning: 'How a boss pull ended, and which attempt it was when a boss took more than one.',
    },
    {
      term: 'Source',
      meaning:
        'Narrows every table to the friendlies, the enemies, or one player. Clicking a name in the summary does the same.',
    },
    {
      term: 'Window',
      meaning:
        'A slice of the fight. Drag across the chart, use the sliders, or pick a preset; every table follows.',
    },
    {
      term: '~',
      meaning:
        'A figure split in proportion to the window or a filter rather than measured directly. Totals are exact.',
    },
    { term: '†', meaning: 'A whole-fight figure the summary cannot cut down to a window.' },
    { term: 'Split', meaning: 'Talent points per tree, in tree order.' },
    {
      term: 'Threat',
      meaning:
        'Threat accumulated from damage and healing under the named model; indicative until every class’s modifiers are in.',
    },
  ];

  let open = $state(false);
</script>

<details class="border-line-soft rounded-panel border" bind:open data-testid="glossary">
  <summary
    class="text-nav inline-flex min-h-11 cursor-pointer items-center px-3 text-[12px] font-bold tracking-[0.06em] uppercase select-none md:min-h-9"
  >
    {open ? 'Hide the glossary' : 'What do these words mean?'}
  </summary>
  <dl class="grid grid-cols-1 gap-x-6 gap-y-2 px-3 pb-3 text-[13px] md:grid-cols-2">
    {#each TERMS as entry (entry.term)}
      <div class="flex flex-col gap-0.5">
        <dt class="text-strong font-semibold">{entry.term}</dt>
        <dd class="text-muted m-0">{entry.meaning}</dd>
      </div>
    {/each}
  </dl>
</details>

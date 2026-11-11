<!-- web/src/components/ui/BigFigures.svelte -->
<!-- The plaque's big figures (round 5, the raid plaque: "Days and hours at 44 px, the largest thing
     on the page, then in and late"): groups of a figure over its label, a hairline between
     groups. The first group may be a live clock (`timer`), read aloud as one label. -->
<script lang="ts" module>
  export interface BigFigure {
    value: string;
    label: string;
  }
</script>

<script lang="ts">
  let {
    groups,
    timer,
    testid,
  }: {
    groups: readonly (readonly BigFigure[])[];
    /** The first group is a clock: role timer, named by this label ("32 days, 5 hours"). */
    timer?: string;
    testid?: string;
  } = $props();

  const shown = $derived(groups.filter((group) => group.length > 0));
</script>

<div class="ui-big-figures" data-testid={testid}>
  {#each shown as group, index (index)}
    {#if index > 0}<span class="ui-big-figures-rule" aria-hidden="true"></span>{/if}
    <div
      class="ui-big-figures-group"
      role={index === 0 && timer !== undefined ? 'timer' : undefined}
      aria-label={index === 0 ? timer : undefined}
    >
      {#each group as figure (figure.label)}
        <span class="ui-big-figure" aria-hidden={index === 0 && timer !== undefined ? 'true' : undefined}>
          <b>{figure.value}</b><span>{figure.label}</span>
        </span>
      {/each}
    </div>
  {/each}
</div>

<style>
  .ui-big-figures {
    display: flex;
    flex-wrap: wrap;
    align-items: stretch;
    gap: 12px 26px;
  }
  .ui-big-figures-group {
    display: flex;
    gap: 26px;
  }
  .ui-big-figures-rule {
    width: 1px;
    background: var(--color-frame-line);
  }
  .ui-big-figure {
    display: flex;
    flex-direction: column;
  }
  .ui-big-figure b {
    font-weight: 600;
    font-size: 44px;
    line-height: 1;
    font-variant-numeric: tabular-nums;
    color: var(--color-strong);
  }
  .ui-big-figure span {
    margin-top: 4px;
    font-size: 13px;
    line-height: 18px;
    color: var(--color-muted);
  }
  @media (max-width: 767px) {
    .ui-big-figures,
    .ui-big-figures-group {
      gap: 10px 16px;
    }
    .ui-big-figure b {
      font-size: 26px;
    }
  }
</style>

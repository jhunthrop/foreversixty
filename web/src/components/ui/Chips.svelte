<!-- web/src/components/ui/Chips.svelte -->
<!-- A panel's own filter row (round 5, design/mocks/directions/prims.py `chips`): what a lane view,
     a log or a query shows, as toggle chips under the panel head. Each chip is a button with
     aria-pressed; a lit chip is white on the cyan tint. `swatch` gives a chip its mark's colour, so
     the row is its own legend (no separate key repeating it). The caller owns the state: several
     lit at once is a filter, one lit at a time is a pick. -->
<script lang="ts" module>
  export interface Chip {
    id: string;
    label: string;
    on: boolean;
    /** The mark's colour (a token var), drawn as a small square before the label. */
    swatch?: string;
    /** A tabular count after the label (the slots a source fills). */
    count?: number;
  }
</script>

<script lang="ts">
  import './chips.css';

  let {
    chips,
    label,
    onToggle,
    testid,
    chipTestid = testid,
  }: {
    chips: readonly Chip[];
    /** The group's accessible name ("Show", "Starting points"). */
    label: string;
    onToggle: (id: string) => void;
    testid?: string;
    /** Each chip's test id prefix (`<prefix>-<id>`), when the group's own id lives on a wrapper. */
    chipTestid?: string;
  } = $props();
</script>

<div class="ui-chips" role="group" aria-label={label} data-testid={testid}>
  {#each chips as chip (chip.id)}
    <button
      type="button"
      class="ui-chip"
      aria-pressed={chip.on}
      data-testid={chipTestid === undefined ? undefined : `${chipTestid}-${chip.id}`}
      onclick={() => onToggle(chip.id)}
      >{#if chip.swatch !== undefined}<i
          class="ui-chip-swatch"
          style:background={chip.swatch}
          aria-hidden="true"
        ></i>{/if}{chip.label}{#if chip.count !== undefined}<span class="ui-chip-count">{chip.count}</span
        >{/if}</button
    >
  {/each}
</div>

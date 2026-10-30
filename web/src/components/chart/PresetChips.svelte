<!-- web/src/components/chart/PresetChips.svelte -->
<!-- The Players control's presets (controls and charts spec §B5.3, fix round 2 ruling): Top 4 +
     bottom 4 (the damage chart's default), Top 8, Bottom 8, and Custom, which is lit while the
     reader's own set is drawn and is not itself a button. One row in the list's head and in the
     legend sheet; the chip look is select.css's `.chip-preset`, shared with the list's foot. -->
<script lang="ts">
  import { chartCopy } from '../../lib/chart/copy';
  import type { DrawPreset, PlayersChoice } from '../../lib/chart/drawn';

  let {
    presets,
    choice,
    onPreset,
    testid = 'players-presets',
  }: {
    presets: readonly DrawPreset[];
    choice: PlayersChoice;
    onPreset: (preset: DrawPreset) => void;
    testid?: string;
  } = $props();
</script>

<div
  class="presets flex flex-wrap items-center gap-1"
  role="group"
  aria-label={chartCopy.players}
  data-testid={testid}
>
  {#each presets as preset (preset)}
    <button
      type="button"
      class="chip-preset"
      aria-pressed={choice === preset}
      data-preset={preset}
      onclick={() => onPreset(preset)}>{chartCopy.presets[preset]}</button
    >
  {/each}
  <span
    class="chip-preset is-inert"
    class:lit={choice === 'custom'}
    data-lit={choice === 'custom' ? 'true' : undefined}
    data-testid="{testid}-custom">{chartCopy.presets.custom}</span
  >
</div>

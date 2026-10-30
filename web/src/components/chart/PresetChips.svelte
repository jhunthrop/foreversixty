<!-- web/src/components/chart/PresetChips.svelte -->
<!-- The Players control's presets (controls and charts spec §B5.3, fix round 2 ruling): Top 4 +
     bottom 4 (the damage chart's default), Top 8, Bottom 8, and Custom, which is lit while the
     reader's own set is drawn and is not itself a button. One row in the list's head and in the
     legend sheet. -->
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
      class="preset"
      aria-pressed={choice === preset}
      data-preset={preset}
      onclick={() => onPreset(preset)}>{chartCopy.presets[preset]}</button
    >
  {/each}
  <span class="preset custom" class:lit={choice === 'custom'} data-testid="{testid}-custom"
    >{chartCopy.presets.custom}</span
  >
</div>

<style>
  .preset {
    display: inline-flex;
    align-items: center;
    min-height: 32px;
    padding: 0 10px;
    border: 1px solid var(--color-line-card);
    border-radius: var(--radius-pill);
    color: var(--color-text);
    font-size: 12px;
    font-weight: 600;
    white-space: nowrap;
  }
  button.preset:hover {
    border-color: rgba(229, 185, 85, 0.45);
  }
  button.preset[aria-pressed='true'],
  .preset.custom.lit {
    border-color: var(--color-gold);
    color: var(--color-gold);
    background: rgba(229, 185, 85, 0.1);
  }
  button.preset:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .preset.custom:not(.lit) {
    color: var(--color-muted);
    border-style: dashed;
  }
  @media (pointer: coarse), (hover: none) {
    .preset {
      min-height: 44px;
    }
  }
</style>

<!-- web/src/components/ui/LevelSlider.svelte -->
<!-- A level control (guides spec 2026-11-08 §3A): previous (44 px, off at the first level), the
     level range (a slider whose value text is "Level 34"; arrows move one level, PageUp/PageDown
     ten, Home and End the ends), next (off at the last level). -->
<script lang="ts">
  import Glyph from './Glyph.svelte';

  let {
    level,
    min,
    max,
    label,
    valueText,
    previousLabel,
    nextLabel,
    disabled = false,
    onchange,
    onslide,
    testid = 'level-slider',
  }: {
    level: number;
    min: number;
    max: number;
    label: string;
    valueText: string;
    previousLabel: string;
    nextLabel: string;
    disabled?: boolean;
    onchange: (level: number) => void;
    /** A drag started (true) or ended (false): a live caption can stay quiet while it moves. */
    onslide?: (sliding: boolean) => void;
    testid?: string;
  } = $props();

  const PAGE_STEP = 10;
  const clamp = (next: number): number => Math.min(max, Math.max(min, next));
  const fill = $derived(`${((level - min) / Math.max(1, max - min)) * 100}%`);

  function onkeydown(event: KeyboardEvent): void {
    const moves: Record<string, number> = { PageUp: PAGE_STEP, PageDown: -PAGE_STEP };
    if (event.key in moves) {
      event.preventDefault();
      onchange(clamp(level + moves[event.key]!));
    } else if (event.key === 'Home' || event.key === 'End') {
      event.preventDefault();
      onchange(event.key === 'Home' ? min : max);
    }
  }
</script>

<div class="ui-lvl" data-testid={testid}>
  <button
    type="button"
    class="ui-lvl-step"
    aria-label={previousLabel}
    disabled={disabled || level <= min}
    onclick={() => onchange(clamp(level - 1))}
    data-testid={`${testid}-previous`}
  >
    <Glyph name="chevron" size={16} class="rotate-180" />
  </button>
  <input
    type="range"
    class="ui-lvl-range"
    {min}
    {max}
    step="1"
    value={level}
    aria-label={label}
    aria-valuetext={valueText}
    {disabled}
    style:--fill={fill}
    oninput={(event) => onchange(clamp(Number(event.currentTarget.value)))}
    {onkeydown}
    onpointerdown={() => onslide?.(true)}
    onpointerup={() => onslide?.(false)}
    onpointercancel={() => onslide?.(false)}
    data-testid={`${testid}-range`}
  />
  <button
    type="button"
    class="ui-lvl-step"
    aria-label={nextLabel}
    disabled={disabled || level >= max}
    onclick={() => onchange(clamp(level + 1))}
    data-testid={`${testid}-next`}
  >
    <Glyph name="chevron" size={16} />
  </button>
</div>

<style>
  .ui-lvl {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  .ui-lvl-step {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    flex: none;
    width: 44px;
    height: 44px;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-field);
    background: none;
    color: var(--color-nav);
  }
  .ui-lvl-step:not(:disabled):hover {
    color: var(--color-strong);
    border-color: var(--color-gold);
  }
  .ui-lvl-step:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .ui-lvl-step:disabled {
    opacity: 0.45;
    cursor: not-allowed;
  }
  .ui-lvl-range {
    flex: 1 1 auto;
    min-width: 0;
    height: 44px;
    margin: 0;
    background: transparent;
    appearance: none;
    cursor: pointer;
  }
  .ui-lvl-range:disabled {
    opacity: 0.45;
    cursor: not-allowed;
  }
  .ui-lvl-range::-webkit-slider-runnable-track {
    height: 4px;
    border-radius: 2px;
    background: linear-gradient(to right, var(--color-gold) var(--fill), var(--color-line-soft) var(--fill));
  }
  .ui-lvl-range::-moz-range-track {
    height: 4px;
    border-radius: 2px;
    background: var(--color-line-soft);
  }
  .ui-lvl-range::-moz-range-progress {
    height: 4px;
    border-radius: 2px;
    background: var(--color-gold);
  }
  .ui-lvl-range::-webkit-slider-thumb {
    appearance: none;
    width: 20px;
    height: 20px;
    margin-top: -8px;
    border-radius: 50%;
    border: 0;
    background: var(--color-gold);
    box-shadow: 0 0 0 4px color-mix(in srgb, var(--color-gold) 20%, transparent);
  }
  .ui-lvl-range::-moz-range-thumb {
    width: 20px;
    height: 20px;
    border-radius: 50%;
    border: 0;
    background: var(--color-gold);
  }
  .ui-lvl-range:focus-visible {
    outline: none;
  }
  .ui-lvl-range:focus-visible::-webkit-slider-thumb {
    box-shadow:
      0 0 0 2px var(--color-bg),
      0 0 0 4px var(--color-gold);
  }
  .ui-lvl-range:focus-visible::-moz-range-thumb {
    box-shadow:
      0 0 0 2px var(--color-bg),
      0 0 0 4px var(--color-gold);
  }
</style>

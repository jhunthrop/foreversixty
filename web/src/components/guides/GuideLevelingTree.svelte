<!-- web/src/components/guides/GuideLevelingTree.svelte -->
<!-- The Leveling section's tree (design/specs/2026-11-04-leveling-tree.md): the guide's own
     ladder from level 10 to 60, one point a level in the order the sim, the addon and the
     planner share (`guideLadderSteps`), drawn by the same read-only trees as the guide's other
     builds (GuideTreeColumns). A level control (slider and step buttons) fills the trees up to
     that level; the caption names the point taken there, and Load in planner follows it. On a
     phone the caption and the control stay in view while the stacked trees scroll. The
     server renders level 60 with the trees as a skeleton, so nothing moves when the island
     loads. -->
<script lang="ts">
  import { guidesCopy } from '../../lib/guides/copy';
  import type { LadderStep } from '../../lib/guides/leveling-builds';
  import { guideTreeStore, loadGuideTreeSource, type GuideTreeSource } from '../../lib/guides/tree-store';
  import ActionLink from '../ui/ActionLink.svelte';
  import Glyph from '../ui/Glyph.svelte';
  import LoadError from '../ui/LoadError.svelte';
  import Skeleton from '../ui/Skeleton.svelte';
  import GuideTreeColumns from './GuideTreeColumns.svelte';

  let {
    code,
    spec,
    steps,
  }: {
    /** The guide's level 60 FS1 code: the class and talent data the trees load. */
    code: string;
    /** The spec's full id (`rogue-combat`), for the planner link. */
    spec: string;
    /** One per level, 10 to 60 (`guideLadderSteps`). */
    steps: LadderStep[];
  } = $props();

  const first = $derived(steps[0]?.level ?? 10);
  const last = $derived(steps[steps.length - 1]?.level ?? 60);
  /** PageUp and PageDown move this many levels. */
  const PAGE_STEP = 10;

  let level = $state(60);
  let source = $state<GuideTreeSource | null>(null);
  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let attempt = $state(0);
  let dragging = $state(false);

  const step = $derived(steps.find((candidate) => candidate.level === level) ?? steps[steps.length - 1]);
  const order = $derived(
    steps
      .filter((candidate) => candidate.level <= level && candidate.talentId !== null)
      .map((candidate) => candidate.talentId!),
  );
  const store = $derived(source === null ? null : guideTreeStore(source, order));
  const plannerHref = $derived(
    `/planner?spec=${encodeURIComponent(spec)}&talents=${encodeURIComponent(step?.talents ?? '')}`,
  );
  const ready = $derived(status === 'ready');

  async function load(): Promise<void> {
    status = 'loading';
    try {
      source = await loadGuideTreeSource(code);
      status = 'ready';
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void attempt;
    void load();
  });

  function setLevel(next: number): void {
    level = Math.min(last, Math.max(first, next));
  }

  function onKeyDown(event: KeyboardEvent): void {
    const moves: Record<string, number> = { PageUp: PAGE_STEP, PageDown: -PAGE_STEP };
    if (event.key in moves) {
      event.preventDefault();
      setLevel(level + moves[event.key]!);
    } else if (event.key === 'Home' || event.key === 'End') {
      event.preventDefault();
      setLevel(event.key === 'Home' ? first : last);
    }
  }
</script>

<div class="leveling-tree" data-testid="guide-leveling-tree">
  <div class="leveling-head" data-testid="guide-leveling-head">
    <div
      class="leveling-caption"
      aria-live={dragging ? 'off' : 'polite'}
      aria-atomic="true"
      data-testid="guide-leveling-caption"
    >
      <p class="leveling-caption-point">
        {step === undefined ? '' : guidesCopy.levelingPoint(step.level, step.name, step.rank)}
      </p>
      <p class="leveling-caption-count">{guidesCopy.levelingPointsSpent(level - first + 1, last - level)}</p>
    </div>

    <div class="leveling-control">
      <button
        type="button"
        class="leveling-step"
        aria-label={guidesCopy.previousLevel}
        disabled={!ready || level <= first}
        onclick={() => setLevel(level - 1)}
        data-testid="guide-leveling-previous"
      >
        <Glyph name="chevron" size={16} class="rotate-180" />
      </button>
      <input
        type="range"
        class="leveling-range"
        min={first}
        max={last}
        step="1"
        value={level}
        aria-label={guidesCopy.levelSliderLabel}
        aria-valuetext={guidesCopy.levelValueText(level)}
        disabled={!ready}
        style={`--fill:${((level - first) / Math.max(1, last - first)) * 100}%`}
        oninput={(event) => setLevel(Number(event.currentTarget.value))}
        onkeydown={onKeyDown}
        onpointerdown={() => (dragging = true)}
        onpointerup={() => (dragging = false)}
        onpointercancel={() => (dragging = false)}
        data-testid="guide-leveling-slider"
      />
      <button
        type="button"
        class="leveling-step"
        aria-label={guidesCopy.nextLevel}
        disabled={!ready || level >= last}
        onclick={() => setLevel(level + 1)}
        data-testid="guide-leveling-next"
      >
        <Glyph name="chevron" size={16} />
      </button>
    </div>
  </div>

  <ActionLink
    href={plannerHref}
    aria-label={guidesCopy.loadLevelBuildAriaLabel(level)}
    data-testid="guide-leveling-load"
  >
    {guidesCopy.loadInPlanner}
  </ActionLink>

  <div class="leveling-trees" aria-busy={status === 'loading'}>
    {#if status === 'failed'}
      <LoadError
        message={guidesCopy.treeLoadFailed}
        onRetry={() => (attempt += 1)}
        testid="guide-tree-error"
      />
    {:else if store === null}
      <Skeleton lines={3} rowHeight="h-12" minHeight="min-h-[360px]" testid="guide-tree-skeleton" />
    {:else}
      <GuideTreeColumns {store} />
    {/if}
  </div>
</div>

<style>
  .leveling-caption {
    min-height: 64px;
  }
  /* On a phone the three trees stack about 1,200 px tall: the caption and the level control stay
     in view under the sticky contents bar (48 px, SpecTableOfContents), so a step always shows
     its tree and its words together. From md the trees sit side by side and it scrolls away. */
  @media (max-width: 767px) {
    .leveling-head {
      position: sticky;
      top: 48px;
      z-index: 10;
      background: var(--color-bg);
      padding-top: 8px;
      box-shadow: 0 8px 12px -8px var(--color-bg);
    }
  }
  @media (min-width: 640px) {
    .leveling-caption {
      min-height: 42px;
    }
  }
  .leveling-caption-point {
    margin: 0;
    font-size: 16px;
    line-height: 22px;
    font-weight: 600;
    color: var(--color-strong);
  }
  .leveling-caption-count {
    margin: 0;
    font-size: 13px;
    line-height: 20px;
    color: var(--color-nav);
  }
  .leveling-control {
    display: flex;
    align-items: center;
    gap: 8px;
    max-width: 560px;
    margin: 8px 0 4px;
  }
  .leveling-step {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 44px;
    height: 44px;
    flex-shrink: 0;
    border-radius: 8px;
    border: 1px solid var(--color-line);
    background: rgb(255 255 255 / 0.045);
    color: var(--color-strong);
    transition: border-color 120ms;
  }
  @media (hover: hover) and (pointer: fine) {
    .leveling-step:not(:disabled):hover {
      border-color: color-mix(in srgb, var(--color-gold) 30%, transparent);
    }
  }
  .leveling-step:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .leveling-step:disabled {
    opacity: 0.45;
    cursor: not-allowed;
  }
  .leveling-range {
    flex: 1 1 auto;
    min-width: 0;
    height: 44px;
    margin: 0;
    background: transparent;
    appearance: none;
    cursor: pointer;
  }
  .leveling-range:disabled {
    opacity: 0.45;
    cursor: not-allowed;
  }
  .leveling-range::-webkit-slider-runnable-track {
    height: 6px;
    border-radius: 3px;
    background: linear-gradient(to right, var(--color-gold) var(--fill), var(--color-line) var(--fill));
  }
  .leveling-range::-moz-range-track {
    height: 6px;
    border-radius: 3px;
    background: var(--color-line);
  }
  .leveling-range::-moz-range-progress {
    height: 6px;
    border-radius: 3px;
    background: var(--color-gold);
  }
  .leveling-range::-webkit-slider-thumb {
    appearance: none;
    width: 24px;
    height: 24px;
    margin-top: -9px;
    border-radius: 50%;
    border: 0;
    background: var(--color-gold);
    transition:
      width 120ms,
      height 120ms;
  }
  .leveling-range::-moz-range-thumb {
    width: 24px;
    height: 24px;
    border-radius: 50%;
    border: 0;
    background: var(--color-gold);
  }
  .leveling-range:focus-visible {
    outline: none;
  }
  .leveling-range:focus-visible::-webkit-slider-thumb {
    box-shadow:
      0 0 0 2px var(--color-bg),
      0 0 0 4px var(--color-gold);
  }
  .leveling-range:focus-visible::-moz-range-thumb {
    box-shadow:
      0 0 0 2px var(--color-bg),
      0 0 0 4px var(--color-gold);
  }
  .leveling-trees {
    margin-top: 12px;
  }
</style>

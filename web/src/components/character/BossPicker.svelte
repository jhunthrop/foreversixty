<!-- web/src/components/character/BossPicker.svelte -->
<!-- The parse column's boss picker (experience spec 2026-10-24-character-history 3.4): the flat
     system's underline tabs (TabStrip's look: sentence case, white text over the accent underline on
     the chosen boss, its kill count muted beside it), one per boss with a kill, newest kill first. One
     tab stop with a roving tabindex; the arrow keys move focus, Enter and Space select. The
     selected pill scrolls into view, and the row fades at its right edge while it overflows. -->
<script lang="ts">
  import { tick } from 'svelte';
  import type { BossPill } from '../../lib/character/history-chart';
  import { historyCopy } from '../../lib/character/history-copy';
  import { scrollBehavior } from '../../lib/character/history-layout';

  let {
    pills,
    activeId,
    onpick,
  }: { pills: readonly BossPill[]; activeId: string; onpick: (id: string) => void } = $props();

  let strip: HTMLDivElement | undefined = $state();
  let moreToRight = $state(false);

  const FADE_PX = 28;

  function measure(): void {
    if (strip === undefined) return;
    moreToRight = strip.scrollLeft + strip.clientWidth < strip.scrollWidth - 1;
  }

  async function revealActive(): Promise<void> {
    await tick();
    const pill = strip?.querySelector('[aria-selected="true"]');
    pill?.scrollIntoView({ block: 'nearest', inline: 'nearest', behavior: scrollBehavior() });
    measure();
  }

  $effect(() => {
    void activeId;
    void revealActive();
  });

  $effect(() => {
    const element = strip;
    if (element === undefined) return;
    const observer = new ResizeObserver(measure);
    observer.observe(element);
    return () => observer.disconnect();
  });

  function onkeydown(event: KeyboardEvent): void {
    const offset = event.key === 'ArrowRight' ? 1 : event.key === 'ArrowLeft' ? -1 : 0;
    if (offset === 0 || strip === undefined) return;
    event.preventDefault();
    const tabs = [...strip.querySelectorAll<HTMLButtonElement>('[role="tab"]')];
    const current = tabs.indexOf(document.activeElement as HTMLButtonElement);
    tabs[Math.min(Math.max(current + offset, 0), tabs.length - 1)]?.focus();
  }
</script>

<div
  bind:this={strip}
  class="picker -mx-3 flex min-h-11 overflow-x-auto"
  style={moreToRight
    ? `mask-image: linear-gradient(to right, black calc(100% - ${FADE_PX}px), transparent); -webkit-mask-image: linear-gradient(to right, black calc(100% - ${FADE_PX}px), transparent)`
    : undefined}
  role="tablist"
  aria-label={historyCopy.pickerLabel}
  tabindex="-1"
  data-testid="history-picker"
  onscroll={measure}
  {onkeydown}
>
  {#each pills as pill (pill.id)}
    {@const active = pill.id === activeId}
    <button
      type="button"
      role="tab"
      aria-selected={active}
      aria-label={pill.accessibleName}
      tabindex={active ? 0 : -1}
      class="focus-visible:outline-gold relative inline-flex h-11 shrink-0 items-center gap-2 px-3 text-[14px] font-semibold focus-visible:outline-2 focus-visible:-outline-offset-2 {active
        ? 'text-strong after:bg-selected-mark after:absolute after:inset-x-3 after:bottom-0 after:h-[2px]'
        : 'text-text hover:text-strong'}"
      data-testid="history-pill"
      onclick={() => onpick(pill.id)}
    >
      {pill.label}
      <span class="text-muted text-[13px] font-normal tabular-nums">{pill.count}</span>
    </button>
  {/each}
</div>

<style>
  .picker {
    scrollbar-width: none;
  }
  .picker::-webkit-scrollbar {
    display: none;
  }
</style>

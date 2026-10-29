<!-- web/src/components/character/BossPicker.svelte -->
<!-- The parse column's boss picker (experience spec 2026-10-24-character-history 3.4): tab
     pills in the MyReports strip's pattern, one per boss with a kill, newest kill first. One
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
  class="picker -mx-1 flex min-h-11 gap-2 overflow-x-auto px-1"
  style={moreToRight
    ? `mask-image: linear-gradient(to right, #000 calc(100% - ${FADE_PX}px), transparent); -webkit-mask-image: linear-gradient(to right, #000 calc(100% - ${FADE_PX}px), transparent)`
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
      class="rounded-control inline-flex h-11 shrink-0 items-center gap-2 border px-3.5 text-[12px] font-bold tracking-[0.06em] uppercase active:bg-[#ffffff0d] {active
        ? 'border-gold text-strong'
        : 'border-line text-text hover:border-line-warm-strong hover:text-strong'}"
      style={active ? 'background: #e5b95514' : undefined}
      data-testid="history-pill"
      onclick={() => onpick(pill.id)}
    >
      {pill.label}
      <span class="text-muted numeric text-[11px] font-normal tracking-normal">{pill.count}</span>
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

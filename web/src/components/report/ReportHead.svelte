<!-- web/src/components/report/ReportHead.svelte -->
<!-- The report's PageHeader (logs spec 2026-11-07 6.1): the title, one meta line (zone, date and
     fight count, then the Sample report or Not credited Tag, or "Logged by {name}", or Live /
     Quiet while the companion logs; the sample's note), the answer at the right (decision R1: a
     trash pull leads with its mobs and length, a boss pull with the DPS) with Sim this fight and
     Copy link, and the one tab strip under it. Below 768 the two actions sit behind one "···"
     menu button, and the note and the fight count are hidden so the meta stays one line. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { reportCopy } from '../../lib/report/copy';
  import type { Credit } from '../../lib/report/credit';
  import type { ReportAnswer } from '../../lib/report/frame';
  import { simCopy } from '../../lib/sim/copy';
  import { button as buttonClass } from '../../lib/ui/button';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';
  import Overlay from '../ui/Overlay.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Tag from '../ui/Tag.svelte';

  let {
    title,
    where,
    fightCount,
    sample,
    credit,
    live,
    answer,
    simHref,
    tabs,
  }: {
    title: string;
    /** "Zephras Isle · 17 Sept 2026". */
    where: string;
    /** "79 fights", hidden below 768. */
    fightCount: string;
    sample: boolean;
    credit: Credit;
    /** The fight on screen is still being logged: 'live', or 'quiet' once the companion went quiet. */
    live: 'live' | 'quiet' | null;
    answer: ReportAnswer | null;
    /** The simulator for this fight; null when no single fight is on screen. */
    simHref: string | null;
    tabs: Snippet;
  } = $props();

  let copied = $state('');
  let menuOpen = $state(false);

  async function copyLink(): Promise<void> {
    menuOpen = false;
    try {
      await navigator.clipboard.writeText(window.location.href);
      copied = reportCopy.copied;
    } catch {
      copied = reportCopy.copyFailed;
    }
    setTimeout(() => (copied = ''), 2000);
  }

  $effect(() => {
    if (!menuOpen) return;
    const close = (event: KeyboardEvent): void => {
      if (event.key === 'Escape') menuOpen = false;
    };
    document.addEventListener('keydown', close);
    return () => document.removeEventListener('keydown', close);
  });
</script>

{#snippet meta()}
  <span data-testid="report-subtitle"
    >{where}<span class="max-md:hidden">{where === '' ? '' : ' · '}{fightCount}</span></span
  >
  <span aria-live="polite" class="contents">
    {#if sample}
      <Tag tone="muted" testid="report-sample-badge">{reportCopy.sampleBadge}</Tag>
    {:else if live !== null}
      <Tag tone="success" testid="report-live">{live === 'live' ? reportCopy.liveTag : reportCopy.quietBadge}</Tag>
    {/if}
  </span>
  {#if !sample}
    {#if credit.credited}
      <span data-testid="report-credit"
        >Logged by {#if credit.href}<a class={LINK} href={credit.href}>{credit.name}</a>{:else}{credit.name}{/if}</span
      >
    {:else}
      <Tag tone="muted" title={credit.prefix} testid="report-not-credited">{reportCopy.notCreditedTag}</Tag>
    {/if}
  {/if}
  {#if sample}<span class="max-md:hidden" data-testid="report-sample-note">{reportCopy.sampleNote}</span>{/if}
{/snippet}

{#snippet answerSlot()}
  {#if answer !== null}
    <!-- The answer column holds 360 px from 900, as the mock's. -->
    <span class="hidden min-[900px]:block min-[900px]:w-[360px]" aria-hidden="true"></span>
    <span class="text-label tracking-label text-muted font-bold uppercase" data-testid="report-answer-label"
      >{answer.label}</span
    >
    <Figure value={answer.value} unit={answer.unit} size="big" tone={answer.tone} testid="report-answer" />
    <span class="text-small text-muted" data-testid="report-answer-caption">{answer.caption}</span>
  {/if}
  <span class="report-actions mt-[6px] flex gap-2">
    {#if simHref !== null}
      <a class={buttonClass('secondary', 'sm')} href={simHref} data-testid="report-sim-fight"
        >{simCopy.simThisFight}</a
      >
    {/if}
    <button
      type="button"
      class={buttonClass('secondary', 'sm')}
      title="Copy a link to exactly this view"
      data-testid="copy-link"
      onclick={() => void copyLink()}>{copied === '' ? reportCopy.copyLink : copied}</button
    >
  </span>
  <span class="report-more">
    <button
      type="button"
      class="{buttonClass('secondary', 'md')} w-11 px-0"
      aria-label={reportCopy.reportActions}
      aria-haspopup="menu"
      aria-expanded={menuOpen}
      data-testid="report-actions"
      onclick={() => (menuOpen = !menuOpen)}>···</button
    >
    {#if menuOpen}
      <span class="report-more-menu">
        <Overlay kind="menu" label={reportCopy.reportActions}>
          <span class="flex flex-col py-1">
            {#if simHref !== null}
              <a class="report-more-item" role="menuitem" href={simHref}>{simCopy.simThisFight}</a>
            {/if}
            <button type="button" class="report-more-item" role="menuitem" onclick={() => void copyLink()}
              >{copied === '' ? reportCopy.copyLink : copied}</button
            >
          </span>
        </Overlay>
      </span>
    {/if}
  </span>
{/snippet}

<PageHeader {title} {meta} answer={answerSlot} {tabs} titleTestid="report-title" testid="report-head" />

<style>
  .report-more {
    display: none;
  }
  @media (max-width: 767px) {
    .report-actions {
      display: none;
    }
    .report-more {
      display: block;
      position: absolute;
      top: 0;
      right: 0;
    }
    :global([data-testid='report-head'] [data-region='head-answer']) {
      position: relative;
    }
  }
  .report-more-menu {
    position: absolute;
    right: 0;
    top: calc(100% + 4px);
    z-index: 40;
    min-width: 200px;
  }
  .report-more-item {
    display: flex;
    align-items: center;
    width: 100%;
    min-height: 44px;
    padding: 0 16px;
    font-size: 14px;
    font-weight: 600;
    color: var(--color-strong);
    text-align: left;
  }
</style>

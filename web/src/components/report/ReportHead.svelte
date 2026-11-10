<!-- web/src/components/report/ReportHead.svelte -->
<!-- The report's title band (logs spec 2026-11-07 6.1; directions b1-report-*): the title, one meta
     line ("Zephras Isle · Pull 6, Trash · 12:43 · 177 mobs died", then Live or Quiet while the
     companion logs; the sample's Tag and its one line on why the numbers are small; who logged it, or Not
     credited), the one figure at the right that follows the tab (group DPS,
     HPS or DTPS, deaths, the player's DPS; none where no one figure answers the tab), and the tab
     strip under it. Sim this fight and Copy link sit in the rail's This pull. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { reportCopy } from '../../lib/report/copy';
  import type { Credit } from '../../lib/report/credit';
  import type { ReportFigure } from '../../lib/report/frame';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Tag from '../ui/Tag.svelte';

  let {
    title,
    line,
    lineTail = '',
    sample,
    credit,
    live,
    figure,
    tabs,
  }: {
    title: string;
    /** "Zephras Isle · Pull 6, Trash · 12:43 · 177 mobs died". */
    line: string;
    /** " · 40 players": shown from 768 only, so the phone's meta stays one line. */
    lineTail?: string;
    sample: boolean;
    credit: Credit;
    /** The fight on screen is still being logged: 'live', or 'quiet' once the companion went quiet. */
    live: 'live' | 'quiet' | null;
    figure: ReportFigure | null;
    tabs: Snippet;
  } = $props();
</script>

{#snippet meta()}
  <span data-testid="report-subtitle"
    >{line}{#if lineTail !== ''}<span class="max-md:hidden">{lineTail}</span>{/if}</span
  >
  <span aria-live="polite" class="contents">
    {#if sample}
      <Tag tone="muted" testid="report-sample-badge">{reportCopy.sampleBadge}</Tag>
    {:else if live !== null}
      <Tag tone="success" testid="report-live"
        >{live === 'live' ? reportCopy.liveTag : reportCopy.quietBadge}</Tag
      >
    {/if}
  </span>
  {#if sample}
    <span class="basis-full max-md:hidden" data-testid="report-sample-note">{reportCopy.sampleNote}</span>
  {:else if credit.credited}
    <span class="max-md:hidden" data-testid="report-credit"
      >Logged by {#if credit.href}<a class={LINK} href={credit.href}>{credit.name}</a
        >{:else}{credit.name}{/if}</span
    >
  {:else}
    <Tag tone="muted" title={credit.prefix} testid="report-not-credited">{reportCopy.notCreditedTag}</Tag>
  {/if}
{/snippet}

{#snippet answer()}
  {#if figure !== null}
    <Figure value={figure.value} unit={figure.unit} align="end" testid="report-answer" />
  {/if}
{/snippet}

<PageHeader {title} {meta} {answer} {tabs} titleTestid="report-title" testid="report-head" />

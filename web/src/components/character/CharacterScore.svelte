<!-- web/src/components/character/CharacterScore.svelte -->
<!-- The Forever Score panel in the rail (round 5, directions mock b2-character; character spec
     2026-11-06 C5, C11): the score and its tier in the tier's colour as Facts, then the parts, each a
     bar on one scale (lib/character/score-parts.ts) and its points. One headline: the per-fight
     rating lives on /ratings (review 2026-11-09: three overlapping numbers read as noise). No logs:
     one line (the owner's says how to get a rating). -->
<script lang="ts">
  import { characterPageCopy as copy } from '../../lib/character/page-view';
  import { scorePartRows } from '../../lib/character/score-parts';
  import { hasScore, tierLabel, tierTextColor, type ForeverScore } from '../../lib/score/score';
  import type { Fact } from '../../lib/ui/facts';
  import Facts from '../ui/Facts.svelte';
  import Meter from '../ui/Meter.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    forever,
    owned,
    fought = false,
  }: {
    forever: ForeverScore;
    owned: boolean;
    /** Ranked fights are on the page: the line says the score is not in yet, never "no logs". */
    fought?: boolean;
  } = $props();

  /** Every bar is one track on one scale, the longest part's maximum (Experience, 700 points). */
  const TRACK_PX = 120;

  const scored = $derived(hasScore(forever));
  const facts = $derived<Fact[]>([
    {
      id: 'score',
      label: copy.scoreFigure,
      value: (forever.score ?? 0).toLocaleString('en-US'),
      testid: 'character-score-figure',
    },
    { id: 'tier', label: copy.tier, content: tierValue, testid: 'character-tier' },
  ]);
  const parts = $derived(scorePartRows(forever.components));
</script>

{#snippet tierValue()}
  <span style:color={tierTextColor(forever.tier) ?? undefined}>{tierLabel(forever.tier) ?? ''}</span>
{/snippet}

<Section title={copy.score} panel testid="character-score">
  {#if !scored}
    <p class="text-text py-[6px] text-[15px] leading-6" data-testid="character-no-logs">
      {fought ? copy.noScoreYet : owned ? copy.noLogsOwner : copy.noLogs}
    </p>
  {:else}
    <Facts items={facts} />
    <div class="mt-3" data-testid="character-parts">
      {#each parts as part (part.key)}
        <Row title={part.label} density="slim" open testid={`character-part-${part.key}`}>
          {#snippet titleContent()}<span class="font-normal">{part.label}</span>{/snippet}
          {#snippet trail()}
            <span class="flex items-center gap-4">
              <span style:width={`${TRACK_PX}px`}
                ><Meter value={part.fill * part.track} tone="class" color={part.color} /></span
              >
              <span class="text-strong w-10 text-right text-[14px] tabular-nums">{part.points}</span>
            </span>
          {/snippet}
        </Row>
      {/each}
    </div>
  {/if}
</Section>

<!-- web/src/components/rankings/RankingsAnswer.svelte -->
<!-- The body of the board's Your rank panel in the rail (logs spec 2026-11-07 section 7; directions
     mock b1-rankings-board): the rank as a big Figure in its percentile colour with
     `of {total} {Spec} {Class}s` (the filter bar names the encounter and ruleset; the marked row
     carries the best kill's figures), and
     `See your parses ›`. Unranked: the sentence and `Upload a log ›`. Failed: `Your rank did not
     load.` The answer is read by the board's own spec (lib/rankings/own-rank.ts); a newer spec
     drops an older read. The block keeps its loading height so the rail never moves. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import type { AnswerSpec, HeaderAnswer } from '../../lib/character-header/types';
  import { ownParsesAction } from '../../lib/rankings/own-rank';
  import { rankingsCopy, rankingsEmptyCopy } from '../../lib/rankings/copy';
  import { LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';

  let {
    spec,
    character,
  }: {
    spec: AnswerSpec;
    character: MeCharacter;
  } = $props();

  let answer = $state<HeaderAnswer | null>(null);
  let failed = $state(false);

  $effect(() => {
    const current = spec;
    const who = character;
    let stale = false;
    answer = null;
    failed = false;
    current.load(who).then(
      (next) => {
        if (!stale) answer = next;
      },
      () => {
        if (!stale) failed = true;
      },
    );
    return () => {
      stale = true;
    };
  });

  const parses = $derived(ownParsesAction(character));
  const figure = $derived(answer?.kind === 'figures' ? answer.figures[0] : undefined);
</script>

<div class="flex min-h-[88px] flex-col gap-1" data-testid="rankings-answer">
  {#if failed}
    <p class="text-small text-muted">{rankingsCopy.answerFailed}</p>
  {:else if answer?.kind === 'sentence'}
    <p data-testid="rankings-answer-unranked">{answer.text}</p>
    <p class="mt-2">
      <a class={LINK} href={rankingsEmptyCopy.href}>{rankingsEmptyCopy.action} ›</a>
    </p>
  {:else if figure !== undefined}
    <Figure
      value={figure.value}
      unit={figure.unit}
      color={figure.color}
      size="big"
      testid="rankings-answer-rank"
    />
    <p class="mt-3">
      <a class={LINK} href={parses.href} data-testid="rankings-answer-parses">{rankingsCopy.parsesLink} ›</a>
    </p>
  {/if}
</div>

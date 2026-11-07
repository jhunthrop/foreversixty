<!-- web/src/components/rankings/RankingsAnswer.svelte -->
<!-- The board's signed-in answer in the PageHeader (logs spec 2026-11-07 section 7, the tier list's
     slot): `Your rank`, the character's crest and name in its class colour, the ordinal as a big
     gold Figure with `of {total} {Spec} {Class}s`, the scope line and `See your parses ›`.
     Unranked: the sentence and `Upload a log ›`. Failed: `Your rank did not load.` The answer is
     read by the board's own spec (lib/rankings/own-rank.ts); a newer spec drops an older read. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import type { AnswerSpec, HeaderAnswer } from '../../lib/character-header/types';
  import { classCrestSrc } from '../../lib/class-crest';
  import { ownParsesAction } from '../../lib/rankings/own-rank';
  import { rankingsCopy, rankingsEmptyCopy } from '../../lib/rankings/copy';
  import { titleColor, isClassSlug, LINK } from '../../lib/ui/primitives';
  import Figure from '../ui/Figure.svelte';

  let { spec, character }: { spec: AnswerSpec; character: MeCharacter } = $props();

  const CREST_PX = 40;

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

  const slug = $derived(character.class.toLowerCase());
  const color = $derived(isClassSlug(slug) ? titleColor({ classSlug: slug }) : undefined);
  const parses = $derived(ownParsesAction(character));
  const figure = $derived(answer?.kind === 'figures' ? answer.figures[0] : undefined);
</script>

<div
  class="flex flex-col gap-1 min-[900px]:min-w-[360px] min-[900px]:items-end"
  data-testid="rankings-answer"
>
  <span class="text-label tracking-label text-muted font-bold uppercase">{rankingsCopy.answerLabel}</span>
  <span class="flex flex-wrap items-center gap-x-3 gap-y-2 min-[900px]:justify-end">
    <img src={classCrestSrc(slug)} alt="" width={CREST_PX} height={CREST_PX} class="rounded-full" />
    <b class="text-[18px] leading-[22px] font-bold" style:color>{character.name}</b>
    {#if figure !== undefined}
      <Figure value={figure.value} unit={figure.unit} tone="gold" size="big" testid="rankings-answer-rank" />
    {/if}
  </span>
  {#if failed}
    <span class="text-small text-muted">{rankingsCopy.answerFailed}</span>
  {:else if answer?.kind === 'sentence'}
    <span class="text-small text-muted" data-testid="rankings-answer-unranked">{answer.text}</span>
    <span><a class={LINK} href={rankingsEmptyCopy.href}>{rankingsEmptyCopy.action} ›</a></span>
  {:else if answer?.kind === 'figures'}
    <span class="text-small text-muted">{answer.note?.text}</span>
    <span
      ><a class={LINK} href={parses.href} data-testid="rankings-answer-parses">{rankingsCopy.parsesLink} ›</a
      ></span
    >
  {:else}
    <span class="text-small text-muted" aria-hidden="true">&nbsp;</span>
    <span aria-hidden="true">&nbsp;</span>
  {/if}
</div>

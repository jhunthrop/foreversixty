<!-- web/src/components/character/CharacterScore.svelte -->
<!-- The Forever Score and the rating in one open Section (character spec 2026-11-06 C5, C11): the
     score, its tier and the rating as Figures, the one "Most to gain" line, the parts in place under
     a Disclosure. No logs: one line (the owner's says how to get a rating). The rating reads its own
     endpoint; until it answers, or with no rated fight, its Figure is left out. -->
<script lang="ts">
  import labels from '../../../../docs/contracts/fixtures/score-labels.json';
  import type { CharacterPath } from '../../lib/characters';
  import { characterHrefs, characterPageCopy as copy } from '../../lib/character/page-view';
  import { fetchCharacterRating } from '../../lib/rankings/api';
  import type { CharacterRating } from '../../lib/rating/types';
  import {
    COMPONENT_MAX,
    hasScore,
    mostToGain,
    tierLabel,
    tierTextColor,
    type ForeverScore,
    type ScoreComponentKey,
  } from '../../lib/score/score';
  import { LINK } from '../../lib/ui/primitives';
  import Disclosure from '../ui/Disclosure.svelte';
  import Figure from '../ui/Figure.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    path,
    forever,
    owned,
    fought = false,
  }: {
    path: CharacterPath;
    forever: ForeverScore;
    owned: boolean;
    /** Ranked fights are on the page: the line says the score is not in yet, never "no logs". */
    fought?: boolean;
  } = $props();

  const scored = $derived(hasScore(forever));
  let rating = $state<CharacterRating | null>(null);
  $effect(() => {
    const requested = path;
    rating = null;
    if (!scored) return;
    void fetchCharacterRating(requested)
      .then((answer) => {
        if (requested === path) rating = answer;
      })
      .catch(() => {
        rating = null;
      });
  });

  const PART_LABEL = Object.fromEntries(labels.components.map((part) => [part.key, part.label])) as Record<
    ScoreComponentKey,
    string
  >;
  const parts = Object.keys(COMPONENT_MAX) as ScoreComponentKey[];
  const most = $derived(mostToGain(forever.components));
  const overall = $derived(rating?.latest?.overall ?? null);
</script>

<Section
  title={copy.score}
  meta={rating !== null && rating.sample_size > 0 ? copy.ratedFights(rating.sample_size) : undefined}
  testid="character-score"
>
  {#if !scored}
    <p class="text-text py-[6px] text-[15px] leading-6" data-testid="character-no-logs">
      {fought ? copy.noScoreYet : owned ? copy.noLogsOwner : copy.noLogs}
    </p>
  {:else}
    <div class="grid grid-cols-2 gap-x-6 gap-y-5 sm:grid-cols-3">
      <Figure
        value={(forever.score ?? 0).toLocaleString('en-US')}
        label={copy.scoreFigure}
        testid="character-score-figure"
      />
      <Figure
        value={tierLabel(forever.tier) ?? ''}
        label={copy.tier}
        color={tierTextColor(forever.tier) ?? undefined}
        testid="character-tier"
      />
      {#if overall !== null}<Figure
          value={String(overall)}
          unit={copy.ratingUnit}
          label={copy.rating}
          testid="character-rating"
        />{/if}
    </div>
    {#if most !== null}
      <p class="text-text py-[6px] text-[15px] leading-6" data-testid="character-most-to-gain">
        {copy.mostToGain(PART_LABEL[most.key], most.gap)}
        <a class={LINK} href={characterHrefs.scoreHelp}>{copy.howScored} ›</a>
      </p>
    {/if}
    <Disclosure summary={copy.parts} testid="character-parts">
      {#each parts as part (part)}
        <Row
          title={PART_LABEL[part]}
          sub={copy.partSub(forever.components?.[part] ?? 0, COMPONENT_MAX[part])}
          open
        />
      {/each}
    </Disclosure>
  {/if}
</Section>

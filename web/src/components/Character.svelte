<!-- web/src/components/Character.svelte -->
<!-- Every ranked fight this character has, their best per encounter, and the builds they
     have been seen in at pull. The page exists for characters who have never signed in:
     a public report is enough to have one, which is the spec's position. -->
<script lang="ts">
  import CurrentCharacterBar from './CurrentCharacterBar.svelte';
  import CharacterPortrait from './character/CharacterPortrait.svelte';
  import { CHARACTER_LOADING_MIN_H } from '../lib/character-layout';
  import { parseCharacterPath, rulesetLabel, type CharacterPath } from '../lib/characters';
  import { classColorVar, formatAmount, percentileToken, rowLink } from '../lib/report/format';
  import { encounterSlug, fetchCharacter, type CharacterPage } from '../lib/rankings/api';
  import { RANKING_METRICS } from '../lib/rankings/url';
  import { executionHref, executionLabel, executionTitle } from '../lib/sim/execution';
  import CharacterHandoffLinks from './CharacterHandoffLinks.svelte';
  import CharacterRatingPanel from './CharacterRatingPanel.svelte';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let { path = null }: { path?: CharacterPath | null } = $props();

  const resolved = $derived(
    path ?? (typeof window === 'undefined' ? null : parseCharacterPath(window.location.pathname)),
  );

  /**
   * A percentile beside a formatted amount has no column heading at any breakpoint --
   * unlike the report island's tables, this list never had one to lose, so the two bare
   * numbers need their own labels rather than relying on position. `RANKING_METRICS`
   * already holds the one label table for `dps`/`hps`/`damage_taken`; echoed as-is if the
   * API ever returns an id this list has not heard of, the same fallback shape
   * `rulesetLabel` and `phaseLabel` use.
   */
  function metricLabel(id: string): string {
    return RANKING_METRICS.find((metric) => metric.id === id)?.label ?? id;
  }

  function percentileAriaLabel(percentile: number | undefined): string | undefined {
    return percentile === undefined ? undefined : `${Math.round(percentile)} percentile`;
  }

  let data = $state<CharacterPage | null>(null);
  let status = $state<'loading' | 'ready' | 'failed' | 'missing'>('loading');
  let error = $state('');
  let attempt = $state(0);

  /**
   * This page has no filter or fight-switching UI -- unlike Rankings.svelte or the report
   * island, `resolved` is derived once from a prop that never changes after mount (the
   * shell passes none) and never re-fires this effect. There is therefore no second,
   * differently-addressed fetch that can land after this one and need dropping. The guard
   * below is kept anyway, at the same cost as Rankings.svelte's `state !== requested`
   * check, so this effect stays correct if that ever stops being true (a future edit that
   * threads a live `path` prop through, say) rather than relying on today's absence of a
   * trigger.
   */
  $effect(() => {
    void attempt;
    const requested = resolved;
    if (requested === null) {
      status = 'missing';
      return;
    }
    status = 'loading';
    void fetchCharacter(requested)
      .then((result) => {
        if (resolved !== requested) return;
        data = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (resolved !== requested) return;
        status = 'failed';
        error = thrown instanceof Error ? thrown.message : 'That character did not load.';
      });
  });
</script>

{#if status === 'missing'}
  <p class="text-[14px]" data-testid="character-missing">
    That is not a character address. They look like <code class="font-mono"
      >/character/us/hardcore/elyra-duskvale</code
    >.
  </p>
{:else if status === 'loading'}
  <Skeleton lines={6} rowHeight="h-4" minHeight={CHARACTER_LOADING_MIN_H} testid="character-skeleton" />
{:else if status === 'failed'}
  <LoadError message={error} onRetry={() => (attempt += 1)} testid="character-error" />
{:else if data !== null && resolved !== null}
  <div class="reveal flex flex-col gap-[22px] md:gap-8" data-testid="character" id="character">
    <CurrentCharacterBar />
    <header class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
      <div class="flex flex-col gap-1">
        <div class="flex items-center gap-3">
          {#if data.character.render_url === undefined}
            <CharacterPortrait character={data.character} size="lg" testid="character" />
          {/if}
          <h1 class="section-title text-[18px]" style={`color: ${classColorVar(data.character.class)}`}>
            {data.character.name}
          </h1>
        </div>
        <p class="text-muted text-[13px]">
          {rulesetLabel(resolved.ruleset)}
          {resolved.region.toUpperCase()}
          {#if data.character.class}· {data.character.class}{/if}
          · <span class="tabular font-mono">{data.history.length}</span> ranked fights
        </p>
        <CharacterHandoffLinks path={resolved} />
      </div>
      {#if data.character.render_url !== undefined}
        <img
          class="hidden max-h-[280px] w-auto object-contain lg:block lg:max-h-[360px]"
          src={data.character.render_url}
          alt=""
          loading="lazy"
          data-testid="character-render"
        />
      {/if}
    </header>

    <CharacterRatingPanel path={resolved} />

    <section class="flex flex-col gap-2">
      <h2 class="section-title text-[18px]">Best per encounter</h2>
      {#if data.best.length === 0}
        <EmptyState message="Nothing ranked yet." testid="character-empty" />
      {:else}
        <ul class="flex flex-col" data-testid="character-best">
          {#each data.best as row (`${row.encounter_id}-${row.difficulty}-${row.metric}`)}
            <li
              class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto_auto_auto] items-center gap-3 border-b px-2 py-2 text-[14px]"
            >
              <a class="{rowLink} truncate" href={`/rankings/${encounterSlug(row.encounter)}`}>
                {row.encounter}
              </a>
              <span
                class="tabular text-right font-mono text-[13px]"
                style={row.percentile === undefined ? undefined : `color: ${percentileToken(row.percentile)}`}
                aria-label={percentileAriaLabel(row.percentile)}
              >
                {row.percentile === undefined ? '' : Math.round(row.percentile)}
              </span>
              <a
                class="{rowLink} tabular justify-end text-right font-mono"
                href={`/reports/${row.report_id}?fight=${row.fight_index}`}
                aria-label={`${formatAmount(Math.round(row.value))} ${metricLabel(row.metric)}`}
              >
                {formatAmount(Math.round(row.value))}
              </a>
              <!-- This table has no `md:` breakpoint at all -- percentile and value are always
                   visible, not desktop-gated -- so the execution cell is always visible too,
                   unlike the same cell on "Every ranked fight" below. -->
              {#if row.execution_score === null}
                <span
                  class="text-muted tabular text-right font-mono text-[13px]"
                  title={executionTitle(null)}
                  aria-label={executionTitle(null)}
                  data-testid="character-best-execution">{executionLabel(null)}</span
                >
              {:else}
                <a
                  class="{rowLink} tabular justify-end text-right font-mono text-[13px]"
                  href={executionHref(row.report_id, row.fight_index)}
                  title={executionTitle(row.execution_score)}
                  data-testid="character-best-execution">{executionLabel(row.execution_score)}</a
                >
              {/if}
            </li>
          {/each}
        </ul>
      {/if}
    </section>

    <section class="flex flex-col gap-2">
      <h2 class="section-title text-[18px]">Every ranked fight</h2>
      <ul class="flex flex-col" data-testid="character-history">
        {#each data.history as row, index (`${row.report_id}-${row.fight_index}-${index}`)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[minmax(0,1fr)_auto_auto] items-center gap-3 border-b px-2 py-2 text-[14px] md:grid-cols-[96px_minmax(0,1fr)_88px_72px_88px_72px]"
          >
            <span class="text-muted tabular hidden font-mono text-[13px] md:inline"
              >{row.fought_at.slice(0, 10)}</span
            >
            <a class="{rowLink} truncate" href={`/reports/${row.report_id}?fight=${row.fight_index}`}
              >{row.encounter}</a
            >
            <span class="text-muted hidden text-[13px] md:inline">{row.spec ?? ''}</span>
            <span
              class="tabular text-right font-mono text-[13px]"
              style={row.percentile === undefined ? undefined : `color: ${percentileToken(row.percentile)}`}
              aria-label={percentileAriaLabel(row.percentile)}
            >
              {row.percentile === undefined ? '' : Math.round(row.percentile)}
            </span>
            <span
              class="tabular text-right font-mono"
              aria-label={`${formatAmount(Math.round(row.value))} ${metricLabel(row.metric)}`}
            >
              {formatAmount(Math.round(row.value))}
            </span>
            {#if row.execution_score === null}
              <span
                class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
                title={executionTitle(null)}
                aria-label={executionTitle(null)}
                data-testid="character-execution">{executionLabel(null)}</span
              >
            {:else}
              <a
                class="{rowLink} tabular hidden text-right font-mono text-[13px] md:inline"
                href={executionHref(row.report_id, row.fight_index)}
                title={executionTitle(row.execution_score)}
                data-testid="character-execution">{executionLabel(row.execution_score)}</a
              >
            {/if}
          </li>
        {/each}
      </ul>
    </section>

    {#if data.builds_seen.length > 0}
      <section class="flex flex-col gap-2">
        <h2 class="section-title text-[18px]">Builds seen at pull</h2>
        <ul class="flex flex-col" data-testid="character-builds">
          {#each data.builds_seen as build (`${build.talent_split}-${build.first_seen}`)}
            <li
              class="border-line-soft flex min-h-11 flex-wrap items-center gap-3 border-b px-2 py-2 text-[14px]"
            >
              <span class="tabular font-mono font-semibold">{build.talent_split}</span>
              {#if build.spec}<span class="text-muted text-[13px]">{build.spec}</span>{/if}
              <span class="text-muted tabular font-mono text-[13px]"
                >first seen {build.first_seen.slice(0, 10)}</span
              >
            </li>
          {/each}
        </ul>
        <p class="text-muted text-[12px]">
          A split, not a planner link: a saved build records the order points were spent in, and the combat
          log does not.
        </p>
      </section>
    {/if}
  </div>
{/if}

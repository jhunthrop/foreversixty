<!-- web/src/components/sim/SimRotationSection.svelte -->
<!-- What the engine is told to press, in words (sim spec 2026-11-06 §4.3, §4.5; spec
     2026-11-04-sim-results 3.11): an open Section of open Rows, the curated APL's own notes, and
     after a run this run's casts per fight beside each line, a line the run never cast dimmed.
     Before a run (no `result`) it is the same lines with nothing counted. The guide is the
     section's action. The spec's fidelity note lives in the page's answer (§4.6), not here. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { loadActionNames, type ActionNames } from '../../lib/sim/action-names';
  import { createCardTables } from '../../lib/sim/card-tables.svelte';
  import { classSpecPath } from '../../lib/sim/gear-card';
  import { rotationCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { plannedRotationRows, rotationActionKeys, rotationRows } from '../../lib/sim/rotation-card';
  import { loadRotationLines, type RotationLines } from '../../lib/sim/rotations';
  import { specDisplayName } from '../../lib/sim/spec-label';
  import type { SimRequest, SimResult } from '../../lib/sim/types';
  import { isStale } from '../../lib/sim/version';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import ActionMark from './ActionMark.svelte';

  let {
    request,
    result = null,
    names = null,
  }: {
    request: SimRequest;
    /** The run whose casts the lines count; absent before a run. */
    result?: SimResult | null;
    /** The page's names when it has them; else the section reads the class's own. */
    names?: ActionNames | null;
  } = $props();

  let lines = $state.raw<RotationLines | null | undefined>(undefined);
  let ownNames = $state.raw<ActionNames | null>(null);
  $effect(() => {
    const { spec } = request;
    loadRotationLines().then(
      (all) => (lines = all[spec] ?? null),
      (error: unknown) => {
        console.error('sim rotation: the rotation did not load', error);
        lines = null;
      },
    );
  });
  $effect(() => {
    if (names !== null) return;
    loadActionNames(activeBuild.build, request.character.class).then(
      (loaded) => (ownNames = loaded),
      (error: unknown) => console.error('sim rotation: the names did not load', error),
    );
  });

  const tables = createCardTables(
    () => result ?? { request },
    () => names ?? ownNames,
    () => [
      ...(lines === null || lines === undefined ? [] : rotationActionKeys(lines)),
      ...(result?.summary.casts.map((row) => row.spell_name) ?? []),
    ],
  );
  const spec = $derived(specDisplayName(request.spec));
  const level = $derived(request.character.level);
  const rows = $derived.by(() => {
    if (lines === null || lines === undefined) return null;
    return result === null
      ? plannedRotationRows(lines, tables.names, level)
      : rotationRows(result, lines, tables.names);
  });
  const drifted = $derived(
    result !== null &&
      lines !== null &&
      lines !== undefined &&
      (result.rotation_sha === undefined
        ? isStale(result.engine_version)
        : lines.sha !== null && result.rotation_sha !== lines.sha),
  );
  const meta = $derived(
    `${result === null ? copy.meta(spec, level) : copy.metaRan(spec, level)}${drifted ? copy.today : ''}`,
  );

  function rankOf(action: string | null): number {
    return action === null ? 0 : tables.visuals(action).rank;
  }
</script>

{#snippet title(name: string, action: string | null)}
  <ActionMark {action} {tables} size={24}
    >{name}{#if rankOf(action) > 0}<span class="text-muted ml-1 text-[13px] font-normal whitespace-nowrap"
        >{copy.rank(rankOf(action))}</span
      >{/if}</ActionMark
  >
{/snippet}

{#snippet count(text: string)}
  <span class="text-muted text-[13px] font-semibold whitespace-nowrap">{text}</span>
{/snippet}

<Section
  title={copy.title}
  {meta}
  action={{ label: copy.guide(spec), href: `/guides/${classSpecPath(request.spec)}` }}
  testid="sim-rotation-saved-card"
>
  {#if rows === null}
    {#if lines === null}<p class="text-muted text-[14px]">{copy.noRotation(spec)}</p>{/if}
  {:else}
    {#if rows.prepull.length > 0}
      <h3 class="text-label text-muted pt-4 pb-1 font-bold">{copy.beforePull}</h3>
      <div data-testid="rotation-prepull">
        {#each rows.prepull as row (row.key)}
          <div
            class={row.cast === false ? 'opacity-50' : ''}
            data-testid="rotation-prepull-row"
            data-cast={row.cast}
          >
            <Row title={row.name} open>
              {#snippet lead()}<span class="text-muted w-10 text-[13px] tabular-nums">{row.at}</span
                >{/snippet}
              {#snippet titleContent()}{@render title(row.name, row.action)}{/snippet}
              {#snippet trail()}{#if row.cast !== null}{@render count(
                    row.cast ? copy.cast : copy.notCast,
                  )}{/if}{/snippet}
            </Row>
          </div>
        {/each}
      </div>
    {/if}
    <h3 class="text-label text-muted pt-4 pb-1 font-bold">{copy.thenPriority}</h3>
    <div data-testid="rotation-priority">
      {#each rows.priority as row (row.key)}
        <div class={row.dimmed ? 'opacity-50' : ''} data-testid="rotation-line" data-dimmed={row.dimmed}>
          <Row title={row.name} sub={row.note === '' ? undefined : row.note} rank={String(row.number)} open>
            {#snippet titleContent()}{@render title(row.name, row.action)}{/snippet}
            {#snippet trail()}{#if row.casts !== null}{@render count(
                  row.dimmed ? copy.notCast : copy.casts(row.casts),
                )}{/if}{/snippet}
          </Row>
        </div>
      {/each}
    </div>
    <p class="text-muted pt-[10px] text-[13px] leading-5">
      {result === null ? copy.foot(spec) : copy.footRan(spec)}
    </p>
  {/if}
</Section>

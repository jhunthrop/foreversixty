<!-- web/src/components/bis/BisSide.svelte -->
<!-- Beside the paperdoll (design/specs/2026-11-09-directions.md, Paperdoll): a tank's or a healer's
     set figures (the frame's bar holds one number), What to stack (the top weights, what 1.00 is
     worth, the band's standing notes, and the sim's weights page for the rest: one job per page)
     and Where to go (each stop of the itinerary below and its pieces). -->
<script lang="ts">
  import Section from '../ui/Section.svelte';
  import Facts from '../ui/Facts.svelte';
  import type { Fact } from '../../lib/ui/facts';
  import type { BisBandModel } from '../../lib/bis/page-model';
  import type { ItineraryStop, StackView } from '../../lib/bis/frame-view';
  import { bisCopy } from '../../lib/bis/copy';

  let {
    model,
    stack,
    stops,
    notes,
    weightsHref,
    whereAction,
  }: {
    model: BisBandModel;
    stack: StackView;
    stops: readonly ItineraryStop[];
    /** The band's standing notes (the source rule, the hit cap, the two-hander). */
    notes: readonly { text: string; title?: string; testid: string }[];
    weightsHref?: string;
    /** Where to go's one link (the pre-raid list). */
    whereAction?: { label: string; href: string; onclick?: (event: MouseEvent) => void };
  } = $props();

  const setFacts = $derived.by((): Fact[] => {
    const answer = model.answer;
    if (answer.tank) {
      return answer.tank.figures.map((figure) => ({
        id: figure.id,
        label: figure.label,
        value: figure.value,
        testid: `bis-tank-figure-${figure.id}`,
      }));
    }
    return (answer.healer?.metrics ?? []).map((metric) => ({
      id: metric.key,
      label: metric.label,
      value: metric.value,
      testid: `bis-healer-${metric.key}`,
    }));
  });

  const stackFacts = $derived(
    stack.rows.map((row): Fact => ({
      id: row.stat,
      label: row.label,
      value: row.scaleFactor.toFixed(2),
      testid: `bis-weight-row-${row.stat}`,
    })),
  );
</script>

<div class="bis-side">
  {#if setFacts.length > 0}
    <Section title={bisCopy.setTitle} boxed testid="bis-set">
      <div class="bis-side-body">
        <Facts items={setFacts} />
        {#if model.answer.tank}<p
            class="bis-side-line"
            title={model.answer.tank.secondaryTitle}
            data-testid="bis-tank-secondary"
          >
            {model.answer.tank.secondaryLine}
          </p>{/if}
      </div>
    </Section>
  {/if}
  <Section title={bisCopy.stackTitle} boxed testid="bis-weights">
    <div class="bis-side-body">
      {#if stackFacts.length > 0}<Facts items={stackFacts} />{/if}
      {#if stack.unitLine}<p class="bis-side-line" data-testid="bis-weights-unit">{stack.unitLine}</p>{/if}
      {#each notes as note (note.testid)}
        <p class="bis-side-line" title={note.title} data-testid={note.testid}>{note.text}</p>
      {/each}
      {#if weightsHref}<a class="bis-side-link" href={weightsHref} data-testid="bis-weights-sim-link"
          >{bisCopy.allWeightsLink} ›</a
        >{/if}
    </div>
  </Section>
  {#if stops.length > 0}
    <Section title={bisCopy.whereTitle} boxed testid="bis-where">
      <div class="bis-side-body bis-stops">
        {#each stops as stop (stop.key)}
          <a class="bis-stop-chip" href={`#${stop.anchor}`}>{stop.title}</a>
        {/each}
      </div>
      {#if whereAction}<p class="bis-side-after">
          <a
            class="bis-side-link"
            href={whereAction.href}
            onclick={whereAction.onclick}
            data-testid="bis-pre-raid-toggle">{whereAction.label} ›</a
          >
        </p>{/if}
    </Section>
  {/if}
</div>

<style>
  .bis-side {
    display: flex;
    flex-direction: column;
    gap: 28px;
    min-width: 0;
  }
  .bis-side-body {
    display: flex;
    flex-direction: column;
    padding-bottom: 8px;
  }
  .bis-side-line {
    margin-top: 10px;
    color: var(--color-muted);
    font-size: 13px;
    line-height: 20px;
  }
  .bis-side-link {
    margin-top: 12px;
    color: var(--color-gold);
    font-size: 14px;
    font-weight: 600;
  }
  .bis-side-after {
    margin-top: 10px;
  }
  /* Jump chips to the stops below: the list under the doll names each piece. */
  .bis-stops {
    flex-direction: row;
    flex-wrap: wrap;
    gap: 8px;
  }
  .bis-stop-chip {
    display: inline-flex;
    align-items: center;
    min-height: 36px;
    padding: 0 12px;
    border: 1px solid var(--color-line-warm-strong);
    border-radius: var(--radius-pill);
    color: var(--color-gold);
    font-size: 14px;
    font-weight: 600;
    line-height: 20px;
  }
  .bis-stop-chip:hover {
    border-color: var(--color-gold);
  }
  .bis-stop-chip:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
</style>

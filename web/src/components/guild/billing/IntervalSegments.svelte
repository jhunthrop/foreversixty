<!-- web/src/components/guild/billing/IntervalSegments.svelte -->
<!-- Yearly / Monthly on native radios: the G1 segment recipe (spec 3.2.2), 44 px, Yearly first.
     Arrow keys move and select with no script; the checked segment shows a gold underline. -->
<script lang="ts">
  import { INTERVALS, type Interval } from '../../../lib/billing/plans';
  import { plansPageCopy } from '../../../lib/billing/plans-page-copy';

  let { value, name, onChange }: { value: Interval; name: string; onChange: (interval: Interval) => void } =
    $props();

  const { interval } = plansPageCopy;
</script>

<div
  role="radiogroup"
  aria-label={interval.groupLabel}
  class="interval-control"
  data-testid="interval-segments"
>
  {#each INTERVALS as option (option)}
    <div class="seg">
      <input
        type="radio"
        {name}
        id={`${name}-${option}`}
        value={option}
        checked={value === option}
        onchange={() => onChange(option)}
        data-testid={`${name}-${option}`}
      />
      <label for={`${name}-${option}`}>{interval.labels[option]}</label>
    </div>
  {/each}
</div>

<style>
  .interval-control {
    display: flex;
    width: 100%;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-control);
    background: var(--color-raised);
  }
  .seg {
    position: relative;
    flex: 1;
    height: 44px;
  }
  .seg input {
    position: absolute;
    inset: 0;
    width: 100%;
    height: 100%;
    margin: 0;
    opacity: 0;
    cursor: pointer;
  }
  .seg label {
    display: flex;
    height: 100%;
    align-items: center;
    justify-content: center;
    font-size: 14px;
    font-weight: 600;
    color: var(--color-text);
    pointer-events: none;
    border-bottom: 2px solid transparent;
    border-radius: var(--radius-control);
  }
  .seg input:checked + label {
    background: var(--color-card-top);
    color: var(--color-strong);
    border-bottom-color: var(--color-gold);
  }
  .seg input:not(:checked):hover + label {
    background: var(--color-card-top);
  }
  .seg input:focus-visible + label {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  @media (min-width: 768px) {
    .interval-control {
      width: 224px;
    }
  }
</style>

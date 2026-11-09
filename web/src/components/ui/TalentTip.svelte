<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  let {
    id,
    name,
    rankLine,
    now = '',
    next = '',
    nextLabel,
    req = '',
    band = '',
    talentId,
  }: {
    /** The element id a button's aria-describedby points at. */
    id?: string;
    name: string;
    /** `Rank 2 of 5`. */
    rankLine: string;
    now?: string;
    next?: string;
    /** `Next rank`. */
    nextLabel: string;
    req?: string;
    band?: string;
    talentId?: number;
  } = $props();
</script>

<!-- web/src/components/ui/TalentTip.svelte -->
<!-- The talent tooltip (planner spec 2026-11-08 section 3A, the mock's TalentTip): the overlay
     ground, 300 px; the name, `Rank n of m`, this rank's text, the next rank under its label, the
     requirement in --wipe when the point is refused, the band line in gold after a band load.
     Draws the box only: TalentButton positions it. `talentId` names the test ids
     (`talent-tip-<id>`, `talent-band-line-<id>`). -->

<div
  class="ttip"
  role="tooltip"
  {id}
  data-testid={talentId === undefined ? undefined : `talent-tip-${talentId}`}
>
  <b>{name}</b>
  <span class="r">{rankLine}</span>
  {#if now}<p>{now}</p>{/if}
  {#if next}<p class="nx"><span class="t-label">{nextLabel}</span>{next}</p>{/if}
  {#if req}<span class="req">{req}</span>{/if}
  {#if band}<span class="bd" data-testid={talentId === undefined ? undefined : `talent-band-line-${talentId}`}
      >{band}</span
    >{/if}
</div>

<style>
  .ttip {
    position: relative;
    width: 300px;
    max-width: calc(100vw - 32px);
    padding: 12px 14px;
    background: var(--color-raised);
    border: 1px solid var(--color-line);
    border-radius: var(--radius-plate);
    box-shadow: 0 18px 40px #000a;
    display: flex;
    flex-direction: column;
    gap: 6px;
    text-align: left;
  }
  b {
    font: 700 15px/20px var(--font-body);
    color: var(--color-strong);
  }
  .r {
    font-size: 13px;
    line-height: 18px;
    color: var(--color-muted);
    font-variant-numeric: tabular-nums;
  }
  p {
    font-size: 13px;
    line-height: 19px;
    color: var(--color-text);
    margin: 0;
  }
  p.nx {
    color: var(--color-muted);
  }
  .t-label {
    display: block;
    margin-bottom: 2px;
    font: 700 12px/16px var(--font-body);
    letter-spacing: 0;
  }
  .req {
    color: var(--color-wipe);
    font-size: 13px;
    line-height: 18px;
  }
  .bd {
    color: var(--color-gold);
    font: 600 13px/18px var(--font-body);
  }
</style>

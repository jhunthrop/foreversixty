<!-- web/src/components/LogsHeroBand.svelte -->
<!-- Logs landing spec (2026-10-04) §4.A: the page's own header BAND -- full-bleed, no
     border, no card -- carrying `LogsHero`. It breaks out with the shared `.full-bleed`
     utility and its inner wrapper is the shared `.page-column`, the same site column every
     page and band uses, so the hero's content lines up with the page's own gutters at every
     width.

     One column, never the planner band's two: a report has no character card to put beside
     it (§4.B explains why the spine, not this band, keeps that job). §4.A: flat
     `--color-raised` floor, no art -- a report has no one class to paint, and principle 5
     keeps atmosphere out of a page with nothing to paint it from. -->
<script lang="ts">
  import LogsHero from './LogsHero.svelte';
</script>

<div class="logs-hero-band full-bleed" data-testid="logs-hero-band">
  <div class="logs-hero-band-inner page-column">
    <LogsHero />
  </div>
</div>

<style>
  .logs-hero-band {
    background: var(--color-raised);
  }
  .logs-hero-band-inner {
    padding-block: 28px 24px;
    padding-inline: var(--page-inset);
  }
</style>

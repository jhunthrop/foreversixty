<!-- web/src/components/LogsHeroBand.svelte -->
<!-- Logs landing spec (2026-10-04) §4.A: the page's own header BAND -- full-bleed, no
     border, no card -- carrying `LogsHero`. Built from the same breakout
     `PlannerHeaderBand.svelte` already ships (`width: 100vw; margin-left: calc(50% - 50vw)`)
     rather than a second way of doing the same thing (tenet 10); the inner wrapper is the
     identical max-width-1344px, 48px-desktop/18px-phone-gutter column every other rebuilt
     page's band uses, so the hero's own content lines up with the page's own gutters at
     1024/1280/1440/1920 exactly as the planner's band does.

     One column, never the planner band's two: a report has no character card to put beside
     it (§4.B explains why the spine, not this band, keeps that job). §4.A: flat
     `--color-raised` floor, no art -- a report has no one class to paint, and principle 5
     keeps atmosphere out of a page with nothing to paint it from. -->
<script lang="ts">
  import LogsHero from './LogsHero.svelte';
</script>

<div class="logs-hero-band" data-testid="logs-hero-band">
  <div class="logs-hero-band-inner">
    <LogsHero />
  </div>
</div>

<style>
  .logs-hero-band {
    width: 100vw;
    margin-left: calc(50% - 50vw);
    background: var(--color-raised);
  }
  .logs-hero-band-inner {
    max-width: 1344px;
    margin: 0 auto;
    padding: 28px 18px 24px 18px;
  }
  @media (min-width: 1024px) {
    .logs-hero-band-inner {
      padding: 28px 48px 24px 48px;
    }
  }
</style>

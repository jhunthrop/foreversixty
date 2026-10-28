<!-- web/src/components/sim/ExampleResultCard.svelte -->
<!-- 2026-09-26 layout pass, review round 1 Finding 5: a static example of a finished sim, so
     a signed-out visitor sees the shape of the result before committing to a load method.
     Every string is static copy from landing-copy.ts -- there is no fetch here, and the
     numbers never change; the "Example" caption says so. -->
<script lang="ts">
  import { serverRunCaption } from '../../lib/billing/plans-page-copy';
  import { landingCopy } from '../../lib/sim/landing-copy';
</script>

<!-- No gutter of its own: every caller (SourceSwitcher.svelte's hero branch) already
     renders inside a section carrying the page's own mx-[18px]/md:mx-0 inset. -->
<div
  class="border-line bg-card-top rounded-panel flex flex-col gap-2 border p-4"
  data-testid="sim-example-result"
>
  <span class="label text-muted">{landingCopy.exampleCaption}</span>
  <p class="text-strong font-semibold">{serverRunCaption()}</p>
  <p class="text-muted text-[13px]">{landingCopy.exampleHint}</p>
</div>

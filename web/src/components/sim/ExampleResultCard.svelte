<!-- web/src/components/sim/ExampleResultCard.svelte -->
<!-- A real finished server run, so a signed-out visitor sees the shape of the result before
     committing to a load method. The label and hint are landing-copy.ts (B5); the picture and its
     caption come from the shared example-image reader /premium uses (spec B3.4). The 4:3 box is
     reserved so a failed image still leaves label, caption and hint. No fetch, no link. -->
<script lang="ts">
  import { EXAMPLE_IMAGE_FAILED_STYLE, exampleImage } from '../../lib/billing/example-image';
  import { serverRunCaption, splitCaptionNumbers } from '../../lib/billing/premium-example-copy';
  import { serverRunExample, type ServerRunExample } from '../../lib/billing/premium-examples';
  import { landingCopy } from '../../lib/sim/landing-copy';

  interface Props {
    example?: ServerRunExample;
  }
  const { example = serverRunExample }: Props = $props();

  const image = $derived(exampleImage(example));
  const captionParts = $derived(splitCaptionNumbers(serverRunCaption(example)));
  /** Spec B4: a failed load hides the img; the reserved box stays plain. */
  let failed = $state(false);
</script>

<!-- No gutter of its own: every caller (SourceSwitcher.svelte's hero branch) already
     renders inside a section carrying the page's own mx-[18px]/md:mx-0 inset. -->
<div
  class="border-line bg-card-top rounded-panel flex flex-col gap-2 border p-4"
  data-testid="sim-example-result"
>
  <span class="label text-muted">{landingCopy.exampleCaption}</span>
  <!-- From 768 the picture stops at 480 px, left in the card (spec B3.4): a 440 px capture is never blown up. -->
  <div class="bg-card-top rounded-[4px] md:max-w-[480px]" style="aspect-ratio: 4 / 3">
    <picture>
      <source srcset={image.avif} type="image/avif" />
      <source srcset={image.webp} type="image/webp" />
      <img
        class="border-line-soft box-border block h-auto w-full rounded-[4px] border"
        style={failed ? `aspect-ratio: 4 / 3; ${EXAMPLE_IMAGE_FAILED_STYLE}` : 'aspect-ratio: 4 / 3'}
        src={image.webp}
        width={image.width}
        height={image.height}
        alt={image.alt}
        loading={image.loading}
        decoding={image.decoding}
        onerror={() => (failed = true)}
      />
    </picture>
  </div>
  <p class="text-muted text-[13px]" data-testid="sim-example-caption">
    {#each captionParts as part, index (index)}{#if part.mono}<span class="font-mono">{part.text}</span
        >{:else}{part.text}{/if}{/each}
  </p>
  <p class="text-muted text-[13px]">{landingCopy.exampleHint}</p>
</div>

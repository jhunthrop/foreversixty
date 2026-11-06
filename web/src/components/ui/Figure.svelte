<!-- web/src/components/ui/Figure.svelte -->
<!-- A compared number (web layout spec 2026-11-05 §3.3, §3A.2): the body face 700 with tabular
     figures (no mono, no condensed face), an optional unit after it and a label above or a caption
     below it. Sizes: sm 16/20 (a paperdoll gain), default 28/32, big 36/40 (the page's answer), xl 56/60 (44/48
     below 768: the sim's answer, spec 2026-11-06-sim 4.3, its unit 18 px). -->
<script lang="ts">
  import { TONE_TEXT, type FigureTone } from '../../lib/ui/primitives';

  let {
    value,
    unit,
    label,
    caption,
    tone = 'strong',
    size = 'default',
    align = 'start',
    testid,
  }: {
    value: string;
    unit?: string;
    /** 12 px caps above the figure. */
    label?: string;
    /** 12 px muted line under the figure ("over next", "DPS"). */
    caption?: string;
    tone?: FigureTone;
    size?: 'sm' | 'default' | 'big' | 'xl';
    align?: 'start' | 'end';
    testid?: string;
  } = $props();

  const SIZE = {
    sm: 'text-[16px] leading-5',
    default: 'text-[28px] leading-8',
    big: 'text-[36px] leading-10',
    xl: 'text-[44px] leading-[48px] tracking-[-0.01em] md:text-[56px] md:leading-[60px]',
  } as const;
</script>

<span
  class={[
    'ui-figure inline-flex min-w-0 flex-col gap-[2px]',
    align === 'end' ? 'items-end text-right' : 'items-start',
  ]}
  data-testid={testid}
>
  {#if label}<span class="text-label tracking-label text-muted font-bold uppercase">{label}</span>{/if}
  <span class={['font-bold tabular-nums', SIZE[size], TONE_TEXT[tone]]}
    >{value}{#if unit}<span
        class={['text-muted font-semibold tracking-normal', size === 'xl' ? 'ml-2 text-[18px]' : 'ml-1 text-[13px]']}
        >{unit}</span
      >{/if}</span
  >
  {#if caption}<span class="text-muted text-[12px] leading-4">{caption}</span>{/if}
</span>

<!-- web/src/components/ui/Row.svelte -->
<!-- The one row (web layout spec 2026-11-05 §3A.2): lead | body | trail on a --color-line-soft
     hairline, 58 px (44 px `slim`), no box. The lead is a snippet (an icon, a crest) or a rank; the
     body a title (body face 700, a class colour from a prop) and a muted sub line; the trail a
     Figure, a Tag or one Button. A row is a link (`href`) or holds an action in its trail, never
     both. Hover is the S2 ground; `selected` the gold ground and a 3 px gold bar. `open` drops the
     side padding for rows in an open Section on S0. `foot` is a full-width line of controls under
     the body, for a row in a narrow rail (a link row never takes one). `fillTrail` gives the
     title a 180 px label column and the trail the rest of the row, a control at full width (the
     sim landing's "A planner build | field, Load"); below 768 the trail takes its own line. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { rowElement, titleColor, type TitleTone } from '../../lib/ui/primitives';

  let {
    title,
    titleTone,
    sub,
    lead,
    rank,
    trail,
    href,
    selected = false,
    density = 'default',
    open = false,
    wrapTrail = false,
    fillTrail = false,
    testid,
    id,
    titleContent,
    subContent,
    foot,
    item = false,
    labelledby,
  }: {
    title: string;
    titleTone?: TitleTone;
    sub?: string;
    lead?: Snippet;
    rank?: string;
    trail?: Snippet;
    href?: string;
    selected?: boolean;
    density?: 'default' | 'slim';
    open?: boolean;
    /** Below 768 the trail takes its own line under the body (an in-row Segmented); `xl`, below
     *  1280 (a trail of figures and actions too wide for a split page's main column). */
    wrapTrail?: boolean | 'xl';
    /** The title in a 180 px label column, the trail filling the rest (a field and its button). */
    fillTrail?: boolean;
    testid?: string;
    /** A link target on the row itself (`/sim/specs#warrior-fury`). */
    id?: string;
    /** Replaces the title's text (a name with its tooltip trigger). */
    titleContent?: Snippet;
    /** Replaces the sub line (a source line with a glyph). */
    subContent?: Snippet;
    /** A line of controls under the body, full width (a narrow rail: tags, votes, a button). */
    foot?: Snippet;
    /** A list item (`li`) of the caller's `ul`, so the rows keep their hairlines between them. Never with `href`. */
    item?: boolean;
    /** The id of the element that names the row (aria-labelledby). */
    labelledby?: string;
    /** Never rendered: an Astro caller's whitespace between named slots arrives as children. */
    children?: Snippet;
  } = $props();

  const element = $derived(rowElement(href, item));
  const hasLead = $derived(lead !== undefined || rank !== undefined);
</script>

<svelte:element
  this={element}
  href={element === 'a' ? href : undefined}
  class={[
    'ui-row border-line-soft relative grid items-center gap-x-3 border-b last:border-b-0',
    fillTrail && 'ui-row-fill',
    hasLead ? 'grid-cols-[auto_minmax(0,1fr)_auto]' : 'grid-cols-[minmax(0,1fr)_auto]',
    // A narrow phone row with nothing at its right gives the body that column and its gap.
    !trail &&
      (hasLead ? 'max-[600px]:grid-cols-[auto_minmax(0,1fr)]' : 'max-[600px]:grid-cols-[minmax(0,1fr)]'),
    density === 'slim' ? 'min-h-11 py-[6px]' : 'min-h-[var(--row-h)] py-[10px]',
    open && !selected ? 'px-0' : 'px-4',
    // Open and selected: the gold band reaches into the panel's padding, the content stays on the column.
    open && selected && '-mx-4',
    element === 'a' &&
      'hover:bg-hover active:bg-pressed focus-visible:outline-gold focus-visible:outline-2 focus-visible:-outline-offset-2',
    selected && 'bg-selected shadow-[inset_3px_0_0_var(--color-accent)]',
  ]}
  aria-current={selected ? 'true' : undefined}
  {id}
  aria-labelledby={labelledby}
  data-testid={testid}
>
  {#if hasLead}
    <span class="flex items-center gap-3">
      {#if rank !== undefined}<span
          class="text-muted w-6 text-right text-[15px] leading-5 font-bold tabular-nums">{rank}</span
        >{/if}
      {#if lead}{@render lead()}{/if}
    </span>
  {/if}
  <span class="min-w-0">
    <span
      class="block text-[15px] leading-5 font-bold {titleTone ? '' : 'text-strong'}"
      style:color={titleTone ? titleColor(titleTone) : undefined}
      >{#if titleContent}{@render titleContent()}{:else}{title}{/if}</span
    >
    {#if subContent}<span class="text-muted block text-[13px] leading-[18px]">{@render subContent()}</span>
    {:else if sub}<span class="text-muted block text-[13px] leading-[18px]">{sub}</span>{/if}
  </span>
  <span
    class={[
      'flex items-center gap-3',
      fillTrail && 'ui-row-trail',
      wrapTrail === 'xl'
        ? 'max-xl:col-span-full max-xl:mt-2 xl:justify-self-end'
        : wrapTrail
          ? 'max-md:col-span-full max-md:mt-2 md:justify-self-end'
          : 'justify-self-end',
      !trail && 'max-[600px]:hidden',
    ]}
    >{#if trail}{@render trail()}{/if}</span
  >
  {#if foot}<span class="col-span-full flex flex-wrap items-center gap-x-3 gap-y-2 pt-2"
      >{@render foot()}</span
    >{/if}
</svelte:element>

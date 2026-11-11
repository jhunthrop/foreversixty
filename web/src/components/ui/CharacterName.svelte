<!-- web/src/components/ui/CharacterName.svelte -->
<!-- A character named the way the game names one (component audit 2026-11-10 cluster 18; spec
     2026-11-05 §3.3): an optional ringed crest (ui/Crest), then the name in its class colour, a link
     when `href` is set. Every class-coloured character name on the web is this component; the
     caller sets only layout and size through `class`. crestSize 0 (the default) draws the name
     alone, for a row whose crest sits in its own column. `weight` inherit keeps a row's own weight. -->
<script lang="ts">
  import { classColorVar } from '../../lib/report/format';
  import Crest from './Crest.svelte';

  let {
    name,
    cls,
    crestSize = 0,
    href,
    onclick,
    title,
    weight = 'semibold',
    class: extra = '',
    testid,
    crestTestid,
  }: {
    name: string;
    /** A class name or slug; unknown or absent, the body colour and the LX crest. */
    cls: string | null | undefined;
    crestSize?: number;
    href?: string;
    onclick?: (event: MouseEvent) => void;
    title?: string;
    weight?: 'semibold' | 'inherit';
    /** The name's own size or truncation. */
    class?: string;
    testid?: string;
    crestTestid?: string;
  } = $props();

  const colour = $derived(classColorVar(cls ?? undefined));
  const nameClass = $derived(['ui-cname', weight === 'semibold' && 'font-semibold', extra]);
</script>

{#if crestSize > 0}<Crest {cls} size={crestSize} testid={crestTestid} />{/if}{#if href !== undefined}<a
    {href}
    {onclick}
    {title}
    class={[nameClass, 'hover:underline']}
    style:color={colour}
    data-testid={testid}>{name}</a
  >{:else}<span class={nameClass} {title} style:color={colour} data-testid={testid}>{name}</span>{/if}

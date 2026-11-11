<!-- web/src/components/home/HomeAbout.svelte -->
<!-- The signed-out home's foot (seo-fix 2026-11-11, lane G2): what the site does for a Forever
     player, in four short paragraphs read from the build's data (lib/home/about.ts), and every
     spec's leveling BiS by role, each spec in its class colour. Server-rendered, no island: it is the
     text a search engine and a player without JavaScript read, under the islands above it. -->
<script lang="ts">
  import { aboutCopy, type AboutView, type Part } from '../../lib/home/about';
  import { LINK, classColor } from '../../lib/ui/primitives';
  import Section from '../ui/Section.svelte';

  let { view }: { view: AboutView } = $props();
</script>

{#snippet part(item: Part)}{#if item.href === undefined}{item.text}{:else}<a
      class={LINK}
      href={item.href}
      style:color={item.classSlug === undefined ? undefined : classColor(item.classSlug)}>{item.text}</a
    >{/if}{/snippet}

<div class="home-about" data-testid="home-about">
  <Section title={aboutCopy.title} boxed testid="home-about-site">
    <div class="home-about-prose">
      {#each view.paragraphs as paragraph, index (index)}
        <p>
          {#each paragraph as item, at (at)}{@render part(item)}{/each}
        </p>
      {/each}
    </div>
  </Section>
  <Section title={aboutCopy.bisTitle} meta={view.bisMeta} boxed testid="home-about-bis">
    <dl class="home-about-roles">
      {#each view.roles as group (group.role)}
        <div>
          <dt class="text-label text-muted font-bold">{group.label}</dt>
          <dd>
            <ul class="home-about-specs">
              {#each group.specs as spec (spec.href)}<li>{@render part(spec)}</li>{/each}
            </ul>
          </dd>
        </div>
      {/each}
    </dl>
  </Section>
</div>

<style>
  .home-about {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
    gap: var(--gap);
    align-items: start;
  }
  @media (min-width: 1100px) {
    .home-about {
      grid-template-columns: minmax(0, 1fr) var(--home-rail, 440px);
    }
  }
  .home-about-prose {
    display: flex;
    flex-direction: column;
    gap: 12px;
    max-width: var(--measure-reading);
    font-size: 15px;
    line-height: 24px;
    color: var(--color-text);
  }
  .home-about-prose p {
    margin: 0;
  }
  .home-about-roles {
    display: flex;
    flex-direction: column;
    gap: 14px;
    margin: 0;
    font-size: 15px;
    line-height: 24px;
  }
  .home-about-roles dd {
    margin: 4px 0 0;
  }
  /* One spec per item, never broken across lines: a name wraps whole. */
  .home-about-specs {
    display: flex;
    flex-wrap: wrap;
    gap: 2px 18px;
    margin: 0;
    padding: 0;
    list-style: none;
  }
  .home-about-specs li {
    white-space: nowrap;
  }
</style>

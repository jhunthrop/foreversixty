<!-- web/src/components/sim/SimLanding.svelte -->
<!-- /sim with no character (spec 2026-11-06-sim §4.2, mock web-sim): the paste box on the page's
     one Plate, three other ways in as open rows, and the "Sim a best-in-slot set" rail (ruling S6:
     the published BiS sets, no figures). sim.astro server-renders it as the page's first paint and
     the island mounts the same component in its place, so nothing moves on hydration. Without
     handlers (the static render) every form still submits to /sim the plain way. -->
<script lang="ts">
  import { classTone } from '../../lib/ui/primitives';
  import { simPageCopy as copy } from '../../lib/sim/page-copy';
  import type { BisSetClass } from '../../lib/sim/bis-set-rail';
  import ClassCrestRing from '../character/ClassCrestRing.svelte';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import FormRow from '../ui/FormRow.svelte';
  import Page from '../ui/Page.svelte';
  import Plate from '../ui/Plate.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    bisSets,
    busy = false,
    pasteError = '',
    message = null,
    onpaste,
    onbuild,
    onfight,
    onsignin,
  }: {
    bisSets: readonly BisSetClass[];
    /** A character is loading: Run reads busy. */
    busy?: boolean;
    /** The paste did not decode. */
    pasteError?: string;
    /** A source's refusal (a build or fight that would not load). */
    message?: string | null;
    onpaste?: (code: string) => void;
    onbuild?: (ref: string) => void;
    onfight?: (ref: string) => void;
    onsignin?: () => void;
  } = $props();

  let paste = $state('');
  let buildRef = $state('');
  let fightRef = $state('');

  function handle(event: SubmitEvent, value: string, handler: ((value: string) => void) | undefined): void {
    if (handler === undefined) return;
    event.preventDefault();
    if (value.trim() !== '') handler(value.trim());
  }
</script>

<Page layout="split" testid="sim-landing">
  <Section title={copy.sections.simYourCharacter} testid="sim-cold-paste">
    <Plate label={copy.sections.simYourCharacter}>
      <div class="flex flex-col gap-3 p-4">
        <FormRow
          label={copy.paste.label}
          action="/sim"
          method="get"
          onsubmit={(event) => handle(event, paste, onpaste)}
        >
          <Field
            id="sim-cold-paste-input"
            name="code"
            label={copy.paste.label}
            placeholder={copy.paste.placeholder}
            grow
            bind:value={paste}
            error={pasteError === '' ? undefined : pasteError}
            testid="sim-cold-paste-input"
          />
          {#snippet submit()}
            <Button tone="primary" type="submit" disabled={busy} testid="sim-cold-paste-run"
              >{copy.paste.run}</Button
            >
          {/snippet}
        </FormRow>
        <p class="text-small text-muted m-0">
          {copy.paste.hint}
          <a class="text-gold font-semibold" href={copy.paste.setUpHref}>{copy.paste.setUp}</a>
        </p>
      </div>
    </Plate>
  </Section>
  <Section title={copy.sections.otherWaysIn} testid="sim-sources">
    <div>
      <Row title={copy.ways.characters} sub={copy.ways.charactersSub} open wrapTrail>
        {#snippet trail()}
          <Button href={onsignin ? undefined : '/account'} onclick={onsignin} testid="sim-sign-in"
            >{copy.ways.signIn}</Button
          >
        {/snippet}
      </Row>
      <Row title={copy.ways.build} sub={copy.ways.buildSub} open wrapTrail>
        {#snippet trail()}
          <FormRow
            label={copy.ways.build}
            action="/sim"
            method="get"
            onsubmit={(event) => handle(event, buildRef, onbuild)}
          >
            <Field
              id="sim-build-input"
              name="ref"
              label={copy.ways.build}
              placeholder={copy.ways.buildPlaceholder}
              labelHidden
              bind:value={buildRef}
              testid="sim-build-input"
            />
            {#snippet submit()}
              <Button type="submit" disabled={busy} testid="sim-build-load">{copy.ways.load}</Button>
            {/snippet}
          </FormRow>
        {/snippet}
      </Row>
      <Row title={copy.ways.fight} sub={copy.ways.fightSub} open wrapTrail>
        {#snippet trail()}
          <FormRow
            label={copy.ways.fight}
            action="/sim"
            method="get"
            onsubmit={(event) => handle(event, fightRef, onfight)}
          >
            <Field
              id="sim-fight-input"
              name="ref"
              label={copy.ways.fight}
              placeholder={copy.ways.fightPlaceholder}
              labelHidden
              bind:value={fightRef}
              testid="sim-fight-input"
            />
            {#snippet submit()}
              <Button type="submit" disabled={busy} testid="sim-fight-load">{copy.ways.load}</Button>
            {/snippet}
          </FormRow>
        {/snippet}
      </Row>
    </div>
    {#if message}
      <p class="text-small text-wipe mt-2" role="alert" data-testid="sim-landing-message">{message}</p>
    {/if}
  </Section>
  {#snippet rail()}
    <Section title={copy.sections.bisSet} meta={copy.sections.bisSetMeta} testid="sim-bis-sets">
      <div>
        {#each bisSets as klass (klass.slug)}
          <Row title={klass.name} titleTone={classTone(klass.slug)} open>
            {#snippet lead()}<ClassCrestRing characterClass={klass.slug} size={32} />{/snippet}
            {#snippet subContent()}
              <span class="flex flex-wrap gap-x-4 gap-y-1 pt-[2px]">
                {#each klass.specs as spec (spec.href)}
                  <a class="text-gold text-[14px] font-semibold" href={spec.href}>{spec.name} ›</a>
                {/each}
              </span>
            {/snippet}
          </Row>
        {/each}
      </div>
    </Section>
  {/snippet}
</Page>

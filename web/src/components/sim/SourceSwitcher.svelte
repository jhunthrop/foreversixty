<!-- web/src/components/sim/SourceSwitcher.svelte -->
<!-- A tool page with no character (mock design/mocks/ws_simtools.py `*-empty`), the landing's own
     grammar (SimLanding.svelte): the addon export on the page's framed panel, the other ways in on a
     panel of their own (a build, a logged fight, your characters). The addon export is first because it needs no
     account and gives the most accurate character; "Your characters" is last because it is the one
     that needs a sign-in. Signed in, that row is the way back to the character just left (the
     line's "Change source" opened this). -->
<script lang="ts">
  import { currentCharacterCopy } from '../../lib/current-character-copy';
  import { landingCopy } from '../../lib/sim/landing-copy';
  import { parseBuildInput } from '../../lib/sim/build-input';
  import { simCopy } from '../../lib/sim/copy';
  import { fightRefOf } from '../../lib/sim/sources';
  import { LINK } from '../../lib/ui/primitives';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import FormRow from '../ui/FormRow.svelte';
  import Page from '../ui/Page.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';

  let {
    busy,
    message,
    signedIn,
    onaddon,
    onbuild,
    onfight,
    onsignin,
    onback = () => {},
    hasCharacters = true,
  }: {
    busy: boolean;
    message: string | null;
    signedIn: boolean;
    onaddon: (code: string) => void;
    onbuild: (id: string) => void;
    onfight: (ref: string) => void;
    onsignin: () => void;
    /** Signed in: back to the character the line's "Change source" left. */
    onback?: () => void;
    /** False for a signed-in account with no characters: nothing to go back to. */
    hasCharacters?: boolean;
  } = $props();

  let addonCode = $state('');
  let buildRef = $state('');
  let fightRef = $state('');

  /** A saved link or id loads the saved build; an unsaved planner link loads its code. */
  function loadBuildInput(): void {
    const input = parseBuildInput(buildRef);
    if (input === null) return;
    if (input.kind === 'code') onaddon(input.code);
    else onbuild(input.id);
  }

  function submitWith(event: SubmitEvent, run: () => void): void {
    event.preventDefault();
    run();
  }
</script>

<Page layout="ledger" panels testid="sim-sources">
  <Section title={simCopy.sourceAddonTitle} panel frame="framed" testid="sim-source-addon">
    <div>
      <div class="flex flex-col gap-3">
        <FormRow
          label={simCopy.sourceAddonTitle}
          onsubmit={(event) => submitWith(event, () => onaddon(addonCode))}
        >
          <Field
            id="sim-addon"
            label={simCopy.sourceAddonTitle}
            labelHidden
            placeholder="FS1:…"
            grow
            disabled={busy}
            bind:value={addonCode}
            testid="sim-addon-input"
          />
          {#snippet submit()}
            <Button tone="primary" type="submit" disabled={busy} testid="sim-addon-load"
              >{simCopy.sourceLoad}</Button
            >
          {/snippet}
        </FormRow>
        <p class="text-small text-muted m-0">
          {simCopy.sourceAddonBody}
          <a
            href="/setup"
            class="{LINK} inline-flex min-h-11 items-center md:min-h-0"
            data-testid="sim-get-addon">{currentCharacterCopy.getTheAddon} ›</a
          >
        </p>
      </div>
    </div>
  </Section>
  <Section title={simCopy.sourceOtherWays} panel>
    <div>
      <Row title={simCopy.sourceBuildTitle} sub={simCopy.sourceBuildBody} open wrapTrail>
        {#snippet trail()}
          <FormRow label={simCopy.sourceBuildTitle} onsubmit={(event) => submitWith(event, loadBuildInput)}>
            <Field
              id="sim-build"
              label={simCopy.sourceBuildTitle}
              placeholder="foreversixty.gg/b/…"
              labelHidden
              disabled={busy}
              bind:value={buildRef}
              testid="sim-build-input"
            />
            {#snippet submit()}
              <Button type="submit" disabled={busy} testid="sim-build-load">{simCopy.sourceLoad}</Button>
            {/snippet}
          </FormRow>
        {/snippet}
      </Row>
      <Row title={simCopy.sourceFightTitle} sub={simCopy.sourceFightBody} open wrapTrail>
        {#snippet trail()}
          <FormRow
            label={simCopy.sourceFightTitle}
            onsubmit={(event) => submitWith(event, () => onfight(fightRefOf(fightRef)))}
          >
            <Field
              id="sim-fight"
              label={simCopy.sourceFightTitle}
              placeholder="foreversixty.gg/reports/…?fight=2"
              labelHidden
              disabled={busy}
              bind:value={fightRef}
              testid="sim-fight-input"
            />
            {#snippet submit()}
              <Button type="submit" disabled={busy} testid="sim-fight-load">{simCopy.sourceLoad}</Button>
            {/snippet}
          </FormRow>
        {/snippet}
      </Row>
      {#if signedIn && hasCharacters}
        <Row title={simCopy.sourceAccountTitle} sub={simCopy.backToCharactersSub} open wrapTrail>
          {#snippet trail()}
            <Button onclick={onback} testid="sim-back-to-characters">{simCopy.backToCharacters}</Button>
          {/snippet}
        </Row>
      {:else if signedIn}
        <Row title={simCopy.sourceAccountTitle} sub={simCopy.charactersOnTheirWay} open wrapTrail>
          {#snippet trail()}
            <Button href="/account" testid="sim-your-account">{simCopy.yourAccount}</Button>
          {/snippet}
        </Row>
      {:else}
        <Row title={simCopy.sourceAccountTitle} sub={simCopy.signInToFindCharacters} open wrapTrail>
          {#snippet trail()}
            <Button onclick={onsignin} testid="sim-signin">{landingCopy.signInWithBattlenet}</Button>
          {/snippet}
        </Row>
      {/if}
    </div>
    {#if message}
      <p class="text-small text-wipe m-0 mt-2" role="alert" data-testid="sim-source-message">{message}</p>
    {/if}
  </Section>
</Page>

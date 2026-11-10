# Companion releases and signing

The desktop companion (`companion/`) is built and released by `.github/workflows/companion.yml`
and published to the site's feed at `foreversixty.gg/dl/companion/` by `tools/publish-feed.sh`.
`/setup` links each platform's build from that feed (`web/src/data/companion.json`), and every
installed companion updates itself from the feed's `latest.json`.

## How a release happens

- A push to main that changes `companion/`, `logs/`, the workflow or
  `tools/package-mac-companion.sh` releases the next patch version after the newest
  `companion-v*` tag. A tag pushed by hand releases that version (a minor or major bump). A
  dispatch with `release: true` releases the commit as it is.
- Land branches and pull requests only test; no build, signing or release runs on them.
- `build` makes Windows (signed) and Linux. `build-mac` makes the Mac app on a macOS runner.
  `release` waits for both, signs every file with minisign, creates the GitHub release and
  publishes the same files to the feed. Only a minisign-signed release reaches the feed.
- The feed's `latest.json` names only the newest release's files, so every release must carry
  every platform. That is why `build-mac` runs on every release, not only on hand-pushed tags,
  and why `release` waits for it. A macOS runner bills at ten times the Linux rate, but it runs
  only when a release is cut (a handful a week), never for tests.
- A dispatch with `dry_run: true` builds, signs and notarizes everything as the next version and
  stops before `release`: no tag, no GitHub release, no feed entry. The builds stay as the run's
  artifacts, and the Windows job installs the setup on the runner, logs what it left behind,
  uninstalls it and keeps two screens as `dryrun-windows-captures`. Any branch may dispatch one:
  `gh workflow run companion.yml --ref <branch> -f dry_run=true`.

## Windows: Azure Artifact Signing

The Windows exe is signed in Commish LLC's name through Azure Artifact Signing. The `build` job
logs in to Azure with the GitHub OIDC federated credential on the `signing` environment and signs
with the certificate profile the repository variable `AZURE_SIGNING_PROFILE` names
(`foreversixty-companion`), then fails the release if the signature does not verify or does not
name Commish LLC. With the variable unset, Windows builds ship unsigned with a notice.
Secrets: `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, `AZURE_SUBSCRIPTION_ID` (environment `signing`);
variables `AZURE_SIGNING_ENDPOINT`, `AZURE_SIGNING_ACCOUNT`, `AZURE_SIGNING_PROFILE`.

The same profile signs the setup program. `build` compiles
`companion/packaging/windows/companion.iss` with Inno Setup 6 (`ISCC.exe`, on the hosted image;
choco installs it if an image drops it) around the signed exe, producing
`foreversixty-companion-setup.exe`. Inno signs as it compiles, through
`companion/packaging/windows/sign.ps1` (the ArtifactSigning module the signing action installed,
the Azure CLI login): the uninstaller it embeds, then the setup. The check step installs the setup
silently on the runner and verifies all three signatures (exe, setup, `unins000.exe`). The setup is
what `/setup` offers; the bare exe stays in every release because installed companions update from
it. The setup's `AppId` keys the Apps & Features entry, must never change, and is the updater's
`SetupAppID` too: each launch of an installed companion writes its version there
(`internal/updater/installed.go`), so Apps & Features follows self-updates.

If signing fails with AADSTS700213, GitHub changed the OIDC subject format again: add a federated
credential on the `foreversixty-companion-signing` app registration that matches the subject in
the error.

## macOS: Developer ID and notarization

`build-mac` builds an arm64 and an Intel binary on the Apple silicon runner, joins them with
`lipo`, and runs `tools/package-mac-companion.sh`, which:

1. wraps the binary in `Forever Sixty.app` (bundle id `gg.foreversixty.companion`, macOS 12 or
   later, the seal as its icon);
2. signs it with the hardened runtime and a secure timestamp;
3. notarizes it with `xcrun notarytool` using the App Store Connect API key, staples the ticket
   and validates it;
4. zips the stapled app as `foreversixty-companion_darwin_universal.app.zip` (the updater's asset);
5. makes `foreversixty-companion_darwin_universal.dmg` (the app beside a link to Applications),
   signs, notarizes and staples it. `/setup` offers the zip, which opens in fewer steps; the app
   moves itself to Applications on its first launch from either (`companion/internal/install`);
6. checks both with `spctl`, the way Gatekeeper will on a player's Mac.

Secrets and variables it reads:

| Name | Kind | What it is |
|---|---|---|
| `APPLE_CERTIFICATE_P12` | secret | base64 of the Developer ID Application certificate and its private key, exported as a .p12 |
| `APPLE_CERTIFICATE_PASSWORD` | secret | the password the .p12 was exported with |
| `ASC_KEY_ID`, `ASC_ISSUER_ID`, `ASC_KEY_P8` | secrets | the App Store Connect API key notarytool signs in with (App Manager role); `ASC_KEY_P8` is the .p8 file's text or base64 of it |
| `APPLE_SIGNING_IDENTITY` | variable, optional | pins the identity, `Developer ID Application: Commish Limited Liability Company (R93F485KX5)`; unset, the one Developer ID Application identity in the .p12 is used |

No Apple ID or app-specific password is needed: the API key covers notarization.

Until `APPLE_CERTIFICATE_P12` exists the job still builds, signs the app ad hoc, skips
notarization with a notice, and keeps the result as the run's `unsigned-companion-darwin`
artifact (seven days) for testing. The release publishes no Mac build then. An ad hoc app from
that artifact opens on a Mac only with a right-click, Open (or after
`xattr -dr com.apple.quarantine "Forever Sixty.app"`).

### The owner's one-time certificate steps

Team: Commish Limited Liability Company, Team ID `R93F485KX5`. Only the Account Holder can create
a Developer ID certificate.

1. On a Mac, open Keychain Access, then Keychain Access > Certificate Assistant > Request a
   Certificate From a Certificate Authority. Enter help@foreversixty.gg and "Commish LLC", choose
   **Saved to disk**, and save the `.certSigningRequest`. The private key is created in the login
   keychain now.
2. At developer.apple.com > Certificates, IDs & Profiles > Certificates, press **+**, choose
   **Developer ID Application**, the **G2 Sub-CA** profile, upload the request, and download
   `developerID_application.cer`.
3. Double-click the .cer so Keychain Access pairs it with the key. Under **My Certificates**,
   "Developer ID Application: Commish Limited Liability Company (R93F485KX5)" now has a
   disclosure triangle with the private key under it.
4. Right-click that certificate, **Export**, format **Personal Information Exchange (.p12)**, and
   set a strong password.
5. Add the two secrets (from the repository root):

   ```sh
   base64 -i Certificates.p12 | gh secret set APPLE_CERTIFICATE_P12
   gh secret set APPLE_CERTIFICATE_PASSWORD   # paste the export password
   gh variable set APPLE_SIGNING_IDENTITY \
     --body "Developer ID Application: Commish Limited Liability Company (R93F485KX5)"
   ```

6. Delete the exported .p12 from disk; keep the certificate and key in the login keychain (or a
   password manager) as the backup. A Developer ID certificate lasts five years; Apple allows a
   handful per team, so do not create spares.

Then cut the first signed release (the coordinator cuts tags) and turn the Mac download on.

### Turning the Mac download on in /setup

`/setup` shows "Download for Mac" only when `web/src/data/companion.json` lists a Mac build, and
the feed's Worker serves only the names that file lists. After the first release whose
`build-mac` job signed and notarized the app (its log says "signing as: Developer ID
Application: ..." and both `spctl` lines say "accepted, source=Notarized Developer ID"), check
the feed lists the disk image:

```sh
curl -s https://foreversixty.gg/dl/companion/latest.json | jq -r '.assets[].name' | grep darwin
```

then add the entry to `companion.json`'s `downloads` and land it:

```json
{
  "platform": "Mac",
  "asset": "foreversixty-companion_darwin_universal.app.zip",
  "firstRun": "Signed by Commish LLC and notarized by Apple. Open the download and Forever Sixty moves itself to Applications, then sets itself up. macOS 12 or later."
}
```

Until then a Mac visitor reads the export route, as before.

### When the Mac build fails

- **"Apple was still processing ... when the wait ran out"**: notarization took longer than the
  job's 30-minute wait (a team's first submissions can; 0.2.37's first run did). Rerun the failed
  job.
- **Notarization "Invalid"**: the job prints Apple's log. The usual causes are a binary signed
  without `--options runtime` or `--timestamp`, or a file in the bundle left unsigned.
- **"The certificate does not carry APPLE_SIGNING_IDENTITY"**: the variable and the .p12 name
  different certificates; fix the variable or export the right certificate.
- **`hdiutil: create failed - Resource busy`**: the script retries three times; rerun the job if
  all three failed.
- **"/setup offers a Mac build but APPLE_CERTIFICATE_P12 is not set"**: the certificate secret
  is gone while `companion.json` lists the Mac download; restore the secret (or take the Mac
  entry out of `companion.json`) and release again.
- **The App Store Connect key was revoked**: create a new key with the App Manager role at App
  Store Connect > Users and Access > Integrations, and replace the three `ASC_*` secrets.
- A failed `build-mac` holds the whole release, Windows included, so the feed never names a
  release without its Mac app. To ship Windows alone in an emergency, remove the Mac entry from
  `companion.json` first, then release with `build-mac` removed from `release`'s `needs`.

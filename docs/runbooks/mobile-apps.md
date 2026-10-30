# Mobile apps runbook

How the Forever Sixty iOS and Android apps are released, rolled out, halted, forced to
update and kept healthy. The plan is `docs/superpowers/plans/2026-10-30-mobile-apps-lanes.md`;
the API contract is `docs/contracts/app-v1.md`; the privacy answers are
`apps/shared/privacy.md`.

## What lives where

| Thing | Where |
|---|---|
| Version and build number of both apps | `apps/VERSION` (`APP_VERSION` x.y.z, `APP_BUILD` rises on every upload) |
| Release notes | `apps/CHANGELOG.md`, the `## Unreleased` section |
| Release command | `tools/release-apps.sh`, or `make -C apps release` |
| Release workflow | `.github/workflows/apps-release.yml` (manual dispatch) |
| fastlane (uploads, rollouts) | `apps/Gemfile` (pinned); run as `BUNDLE_GEMFILE=apps/Gemfile bundle exec fastlane ...` |
| Crash and hang reports | Sentry, projects `forever-ios` and `forever-android` |
| App traffic per version | Cloud Logging, the `api` service's `request` lines with `app_*` fields |
| Forced update | API env `APP_CLIENT_HEADERS`, `APP_MIN_VERSION_IOS`, `APP_MIN_VERSION_ANDROID`, `APP_STORE_URL_*` |

Secrets live in the GitHub `play-internal` environment (release) and in Google Secret Manager
(the API). Never paste one into a file, a commit or a chat.

## One-time setup (owner)

These are still open as of 2026-10-30:

1. **Play.**
   - The app exists in Play Console. Its first bundle must be uploaded by hand: the API cannot
     create a release for an app that has never had one.
   - Add these to the `play-internal` environment:
     - `PLAY_SERVICE_ACCOUNT_JSON`. Its service account needs "Release to testing tracks", and
       "Release to production" for the rollout steps below.
     - `ANDROID_UPLOAD_KEYSTORE_BASE64`, `ANDROID_UPLOAD_KEYSTORE_PASSWORD`,
       `ANDROID_UPLOAD_KEY_ALIAS`, `ANDROID_UPLOAD_KEY_PASSWORD`.
   - Play App Signing holds the app key. We hold only the upload key.
2. **Apple.** When enrollment clears:
   - Set the repository variable `APPLE_TEAM_ID`.
   - Create an App Store Connect API key with the App Manager role. Add `ASC_KEY_ID`,
     `ASC_ISSUER_ID` and `ASC_KEY_P8` (the .p8 file's text) to `play-internal`.
   - Create the app record (`gg.foreversixty.app`) in App Store Connect.
   - `project.yml` still signs ad hoc for the simulator. The release command overrides that
     with `DEVELOPMENT_TEAM` and automatic signing. Run a first archive by hand to confirm the
     profile resolves.
3. **Sentry.**
   - Create the projects `forever-ios` and `forever-android`.
   - Apply the privacy settings in `apps/shared/privacy.md`: Prevent Storing of IP Addresses,
     the scrubbers and the extra fields.
   - Set the repository variables `SENTRY_DSN_IOS` and `SENTRY_DSN_ANDROID`. Without them the
     builds run and report nothing.
4. **API counts.**
   - Set `APP_CLIENT_HEADERS=true` on the `api` service. This lands after Nov 6, per the plan's
     freeze.
   - Create the two log-based metrics in "Counting app traffic".

## Cutting a release

1. Write the player-facing notes under `## Unreleased` in `apps/CHANGELOG.md`, one line per
   change. Play caps them at 500 characters, and the command refuses longer notes. Commit and
   push.
2. Dry run, from a clean main:
   ```bash
   make -C apps release ARGS="--bump minor"
   ```
   It prints the version it would cut, the notes, what is missing (secrets, Apple) and every
   command. Nothing changes.
3. Release, either from GitHub or locally:
   - **GitHub (preferred):** Actions > apps-release > Run workflow on `main`. Choose the
     bump and platform, set execute, and leave play_status at draft until the app has passed
     its first review.
   - **Locally,** with the same environment variables exported:
     ```bash
     make -C apps release ARGS="--bump minor --execute"
     git push --follow-tags
     ```
4. The command:
   - bumps `apps/VERSION`;
   - moves the notes under `## <version> (<build>), <date>`;
   - commits `chore(apps): release <version> (<build>)` and tags `apps-v<version>-<build>`;
   - builds the signed Android bundle and uploads it to the Play internal track with its R8
     mapping;
   - archives iOS and uploads it to TestFlight.

   iOS is skipped, with the reason printed, until the Apple secrets exist. The workflow
   pushes the commit and tag even when an upload failed after the bump, so a spent build
   number is never reused.
5. Bumps:
   - `build` is a new upload of the same version, for a fix to a build still in testing.
   - `patch` is a bug-fix release, `minor` a feature release, `major` a rethink.
6. Check the build in Play Console (Testing > Internal) and TestFlight, install it on a real
   device, and sign in.

Internal testers then get it. Moving to production is the staged rollout below.

## Staged rollout

Never release to 100% at once. Each step holds about 24 hours, and moves on only if all of
these hold:

- **Sentry:** crash-free sessions ≥ 99.5% for the new release, and no new issue among the top
  five.
- **Play Console > Android vitals:** user-perceived crash rate under 1.09% and ANR rate under
  0.47% (Play's bad-behaviour thresholds).
- **Cloud Logging:** the new version's 5xx share is no higher than the last version's (see
  "Counting app traffic").
- **Support:** no new pattern in Discord or help@foreversixty.gg.

### Play: 1% > 5% > 20% > 50% > 100%

Promote the internal build to production at 1%. Use the version code, which is `APP_BUILD`:

```bash
export BUNDLE_GEMFILE=apps/Gemfile SUPPLY_JSON_KEY_DATA="$(cat play-key.json)"
bundle exec fastlane run upload_to_play_store package_name:gg.foreversixty.app \
  track:internal track_promote_to:production version_code:43 rollout:0.01 \
  skip_upload_aab:true skip_upload_metadata:true skip_upload_images:true \
  skip_upload_screenshots:true skip_upload_changelogs:true
```

Widen it (`rollout:0.05`, `0.2`, `0.5`), and finish with `release_status:completed`:

```bash
bundle exec fastlane run upload_to_play_store package_name:gg.foreversixty.app \
  track:production version_code:43 rollout:0.2 skip_upload_aab:true \
  skip_upload_metadata:true skip_upload_images:true skip_upload_screenshots:true
```

Play Console does the same: Release > Production > Manage rollout.

### App Store: phased release over 7 days

In App Store Connect, open the version and choose "Release update over 7-day period using
phased release". Apple's schedule is 1%, 2%, 5%, 10%, 20%, 50%, then 100%. It applies only to
automatic updates; anyone can update by hand from the store at any time.

- **Pause** stops the schedule, for up to 30 days in total.
- **Release to All Users** ends it early.

Hold the same gates as Play before letting it run.

## Halting a bad release

1. **Stop the spread.**
   - **Play:** halt the rollout with Release > Production > Halt rollout, or with fastlane:
     `track:production version_code:<n> release_status:halted`. Players who have not updated
     stay on the last good version. Resume later with `release_status:inProgress`.
   - **App Store:** Pause Phased Release. Apple cannot pull one version back. The fix is a new
     build.
2. **Fix forward.** Cut a `patch` release with the fix and roll it out from 1% again. For a
   crash at launch on iOS, ask for an expedited review (App Store Connect > Contact Us > App
   Review).
3. **If the bad version damages data or cannot reach the API correctly,** force everyone on it
   to update once the fix is live at 100% (next section).
4. **If the problem is in the API, not the app:** roll the API back, not the app. Use Cloud
   Run > api > Revisions > send traffic to the last good revision.

## Forcing an update (the minimum-version gate)

The API answers `426 update_required` to any app whose `X-App` version is below the minimum
for its platform, before any route runs. Both apps then show only "Update Forever Sixty",
with a button to the store page (`docs/contracts/app-v1.md`, "Forced update"). The site is
never affected.

1. Confirm the target version is live at 100% in that store. Gating ahead of that strands
   players with nothing to update to.
2. See who would be cut off, using the per-version counts below. A version still carrying a
   large share of traffic is a reason to wait.
3. Set the minimum. `APP_CLIENT_HEADERS` must already be `true`:
   ```bash
   gcloud run services update api --region us-east1 \
     --update-env-vars APP_MIN_VERSION_ANDROID=1.3.0
   # iOS also needs its store page, which has no default:
   gcloud run services update api --region us-east1 \
     --update-env-vars APP_MIN_VERSION_IOS=1.3.0,APP_STORE_URL_IOS=https://apps.apple.com/app/id<APP_STORE_ID>
   ```
   Startup fails, and the old revision keeps serving, if a minimum is malformed, set without
   `APP_CLIENT_HEADERS=true`, or set for iOS without `APP_STORE_URL_IOS`. The `app-client`
   startup line shows the gate in force.
4. Verify:
   ```bash
   curl -s -H 'X-App: android/1.2.9 (1)' https://api.foreversixty.gg/version
   ```
   This should answer `426` with `update_required`. Gated requests log `app_gated: true`.
5. Undo by removing the variable:
   ```bash
   gcloud run services update api --region us-east1 --remove-env-vars APP_MIN_VERSION_ANDROID
   ```

The gate compares only `x.y.z`, never the build number. It also gates sign-in, so an
outdated app cannot sign in either.

## Counting app traffic

With `APP_CLIENT_HEADERS=true`, each app request's `request` log line carries `app_platform`,
`app_version`, `app_build` and, when it is a UUID, `app_install`. Requests from the site
carry none of these.

Create two log-based metrics once:

```bash
cat > /tmp/app-requests.yaml <<'EOF'
filter: >-
  resource.type="cloud_run_revision" AND resource.labels.service_name="api"
  AND jsonPayload.msg="request" AND jsonPayload.app_platform:*
metricDescriptor:
  metricKind: DELTA
  valueType: INT64
  labels:
    - key: platform
    - key: version
    - key: status
labelExtractors:
  platform: EXTRACT(jsonPayload.app_platform)
  version: EXTRACT(jsonPayload.app_version)
  status: EXTRACT(jsonPayload.status)
EOF
gcloud logging metrics create app_requests --config-from-file=/tmp/app-requests.yaml
sed 's/jsonPayload.app_platform:\*/jsonPayload.app_platform:* AND jsonPayload.status>=500/' \
  /tmp/app-requests.yaml > /tmp/app-errors.yaml
gcloud logging metrics create app_errors --config-from-file=/tmp/app-errors.yaml
```

Chart them in Cloud Monitoring (Metrics explorer, `logging/user/app_requests`), grouped by
`version`. To count distinct installs per version, use Log Analytics:

```sql
SELECT JSON_VALUE(json_payload.app_version) AS version, COUNT(DISTINCT JSON_VALUE(json_payload.app_install)) AS installs
FROM `foreversixty.global._Default._AllLogs`
WHERE JSON_VALUE(json_payload.app_platform) = 'android' AND timestamp > TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 7 DAY)
GROUP BY version ORDER BY version DESC
```

The install id is random, tied to no account and reset on reinstall (`apps/shared/privacy.md`).
Never join it to an account.

## Where crash reports live

- **Sentry** (sentry.io), projects `forever-ios` and `forever-android`. Releases are named
  `gg.foreversixty.app@<version>+<build>`, which is how the rollout gates compare versions.
  - Hangs over 2 seconds (iOS) and ANRs (Android) are issues of their own.
  - Reports exist only from release builds made with a DSN. Debug, fixture and DSN-less builds
    never start the SDK.
- **Symbols are not uploaded to Sentry yet. This is a follow-up.**
  - Android stack traces from R8-minified builds need the mapping. Each release keeps
    `mapping.txt` as the `android-release` artifact of its apps-release run for 90 days, and
    Play gets it with the upload, so Android vitals are readable. To read a Sentry trace,
    retrace it with that run's `mapping.txt`.
  - iOS uploads its symbols to App Store Connect (`uploadSymbols`). The Sentry dSYM upload
    (sentry-cli) is still to add.
- **Play Console > Android vitals** (crashes and ANRs, from Google's own collection) and
  **Xcode Organizer > Crashes** (from Apple's, for players who share analytics) are the
  second sources.
- **TestFlight feedback:** testers' screenshots and crash reports, in App Store Connect >
  TestFlight > Feedback.

## Rotating the push keys

Lane A2's senders read these. The API service, the raid-posts job and the parse-report job all
need the same values, so update all three:

- `APNS_KEY_P8`, `APNS_KEY_ID`, `APNS_TEAM_ID`, `APNS_TOPIC`
- `FCM_SERVICE_ACCOUNT_JSON`

**APNs** (yearly, or at once if a key leaks). An Apple team can hold two APNs keys at once,
so rotation has no gap.

1. In Apple Developer > Certificates, IDs & Profiles > Keys, create a key with Apple Push
   Notifications service enabled, and download the .p8. Apple lets you download it only once.
2. Add the new versions in Secret Manager:
   ```bash
   gcloud secrets versions add APNS_KEY_P8 --data-file=AuthKey_NEWID.p8
   ```
   Update `APNS_KEY_ID` the same way, or as an env var if it is not a secret.
3. Redeploy the `api` service and update both jobs so they read the new versions.
4. Send a test push to your own device (raid posted on a test raid). Check the `push` log
   lines for `apns` successes.
5. Revoke the old key in Apple Developer.
6. Delete the downloaded .p8 from your machine.

**FCM:**

1. In Google Cloud > IAM > Service accounts, open the FCM sender account, then
   Keys > Add key > JSON.
2. Add it as a new version:
   ```bash
   gcloud secrets versions add FCM_SERVICE_ACCOUNT_JSON --data-file=key.json
   ```
   Then redeploy the service and update both jobs.
3. Send a test push to an Android device. Then delete the old key in IAM, and delete
   `key.json` locally.

`google-services.json` in the Android build identifies the Firebase project. It is not a
secret and does not rotate.

## When push fails

Work from the server out to the device:

1. **Are the senders on?**
   - `APP_NOTIFICATIONS_ENABLED=1` must be set, and the keys present.
   - The startup log names each sender, and names only variables, never values.
   - With no keys, no sender exists, and nothing is attempted.
2. **Did the event fire?** Each event is claimed once per account, dedupe key and channel in
   `notification_deliveries`.
   - A row means the send was attempted.
   - No row means one of these:
     - the producer never ran (for time-based events, check the raid-posts job's last run);
     - the player turned the event off;
     - quiet hours held it.
3. **Does the device have a token?** Look for the account's rows in `push_tokens`.
   - None means the app never registered. Possible causes:
     - permission was denied;
     - sign-in did not complete;
     - lane F7's registration failed. Check Sentry for the app's error.
   - Tokens that APNs or FCM call invalid are deleted on the spot, so a vanished token means
     the store rejected it.
4. **What did the store say?**
   - APNs:
     - `BadDeviceToken` usually means a sandbox token sent to production, or the reverse.
       `APNS_SANDBOX=1` is for debug builds only.
     - `TopicDisallowed` means `APNS_TOPIC` does not match the bundle id.
     - `InvalidProviderToken` means a revoked or mistyped key.
   - FCM:
     - `UNREGISTERED` means the app was uninstalled.
     - `SENDER_ID_MISMATCH` means the token belongs to another Firebase project.
     - A `401` or `403` means a bad service account key.
5. **On the device:**
   - iOS: notifications allowed, no Focus mode silencing the app, and Background App
     Refresh on.
   - Android 13 or later: the `POST_NOTIFICATIONS` permission is granted, the channel is not
     muted, and the app is not restricted by battery optimisation.

## When sign-in fails

Sign-in is lane A1's flow (`app-v1.md`, "Sign in"): Battle.net in the system browser, then
the `https://foreversixty.gg/app/auth` landing, then `POST /v1/auth/app/exchange`.

1. **Is the flow mounted?** It needs `APP_SIGN_IN=true` plus Battle.net configured. When the
   flow is off, the routes answer `404`.
2. **Does the hand-off reach the app?**
   - Check both well-known files and their contents:
     ```bash
     curl -s https://foreversixty.gg/.well-known/apple-app-site-association
     curl -s https://foreversixty.gg/.well-known/assetlinks.json
     ```
     They must carry the real Team ID and the Play App Signing SHA-256, never `TODO`.
   - If they are wrong, the landing opens in the browser instead of the app. The source is
     `web/app-links.json`.
   - iOS caches the association file through Apple's CDN. A fix can take a day to reach
     devices; reinstall the app to refresh it.
3. **What did the exchange answer?**
   - `400 invalid_grant`: the code was unknown, already used, more than 2 minutes old, or had
     the wrong verifier. Any attempt spends the code, so the player must start again. Repeated
     reports usually mean a slow hand-off or a device clock far off.
   - `400 invalid` at `/app/auth`: a forged or expired state. The player starts again in the
     app.
   - `401 session_expired` later: the token lapsed after 30 idle days or 180 days in total
     (`APP_TOKEN_IDLE_DAYS`, `APP_TOKEN_MAX_AGE_DAYS`). This is expected; the app signs in
     again.
   - `403 browser_session_required`: a billing write from the app. This is expected; the app
     sends the player to the site.
4. **Is Battle.net itself down?** Check the API's `bnet` log lines and Blizzard's status page.
   The web sign-in will be failing too.
5. **Was the device revoked?** A player who removes the device on the site's account page
   gets `401 unauthorized` on the next request, which is correct.
6. **Is the app outdated?** A `426 update_required` blocks sign-in as well. Check the
   minimum-version settings.

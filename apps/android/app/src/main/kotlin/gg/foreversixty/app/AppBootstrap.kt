package gg.foreversixty.app

import android.content.Context
import gg.foreversixty.app.auth.AndroidTokenStore
import gg.foreversixty.app.auth.TokenStore
import gg.foreversixty.app.auth.session.SessionTokens
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.network.ClientIdentity
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.network.FileAppConfigCache
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.UpdateGate
import gg.foreversixty.app.update.InstallId
import java.io.File

/**
 * What exists before the first app config is known: the API client, the feature switches and
 * the update gate. [AppGraph] is built on top of it, and builds its features from the switches as
 * they stand then, so on a fresh install MainActivity holds the graph until the first
 * `GET /v1/app/config` answers ([AppConfigRepository.awaitFirstAnswer]) or the hold passes.
 */
class AppBootstrap(context: Context, val config: ApiConfig = ApiConfig.fromBuildConfig()) {
    /** The device token, reporting every 401 discard to the session (lane F1). */
    val sessionTokens = SessionTokens(AndroidTokenStore.create(context))
    val tokens: TokenStore = sessionTokens

    /** Shuts when the API says this build is below its minimum; MainActivity then blocks. */
    val updateGate = UpdateGate()

    val identity = ClientIdentity(
        version = BuildConfig.VERSION_NAME,
        build = BuildConfig.VERSION_CODE,
        installId = InstallId.current(context),
    )

    val apiClient: ApiClient = ApiClient(
        ForeverHttpClient.create(
            config = config,
            tokens = tokens,
            engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets), { DebugFixtures.latencyMillis }) { DebugFixtures.variants },
            identity = identity,
        ),
        updateGate,
    )

    /** Read at launch (MainActivity) and cached; feature lanes ask it `isEnabled(feature)`. */
    val appConfig = AppConfigRepository(apiClient, FileAppConfigCache(File(context.filesDir, "app-config.json")))

    /** Whether a feature's screens show: on in debug and fixture builds, else as the config says. */
    fun isAvailable(feature: Feature): Boolean = appConfig.isAvailable(feature, config.mode, BuildConfig.DEBUG)
}

package gg.foreversixty.app

import android.content.Context
import android.os.Build
import gg.foreversixty.app.auth.AndroidTokenStore
import gg.foreversixty.app.auth.AppSignIn
import gg.foreversixty.app.auth.TokenStore
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.network.ForeverHttpClient

/**
 * The app's object graph, built once per process and passed down explicitly. Feature lanes add
 * their `XxxApi(apiClient)` here. A DI framework can replace this when the graph outgrows it.
 */
class AppGraph(context: Context, val config: ApiConfig = ApiConfig.fromBuildConfig()) {
    val tokens: TokenStore = AndroidTokenStore.create(context)

    val apiClient: ApiClient = ApiClient(
        ForeverHttpClient.create(
            config = config,
            tokens = tokens,
            engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)),
        ),
    )

    val accountApi = AccountApi(apiClient)

    val signIn = AppSignIn(
        baseUrl = config.baseUrl,
        client = apiClient,
        tokens = tokens,
        deviceName = Build.MODEL,
    )
}

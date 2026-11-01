package gg.foreversixty.app

import android.content.Context
import android.os.Build
import java.io.File
import gg.foreversixty.app.account.AccountDeps
import gg.foreversixty.app.auth.AndroidTokenStore
import gg.foreversixty.app.auth.AppSignIn
import gg.foreversixty.app.auth.CustomTabsLauncher
import gg.foreversixty.app.auth.TokenStore
import gg.foreversixty.app.auth.session.FixtureBrowser
import gg.foreversixty.app.auth.session.SessionManager
import gg.foreversixty.app.auth.session.SessionTokens
import gg.foreversixty.app.guides.GuidesFeature
import gg.foreversixty.app.home.HomeDependencies
import gg.foreversixty.app.premium.DebugStorefront
import gg.foreversixty.app.premium.PlansLink
import gg.foreversixty.app.premium.PlayStorefront
import gg.foreversixty.app.network.ApiMode
import kotlinx.coroutines.MainScope
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.bis.createBisModel
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.network.ClientIdentity
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.network.FileAppConfigCache
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.talents.TalentsFeature
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.UpdateGate
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.update.InstallId
import gg.foreversixty.app.widgets.WidgetsFeature
import gg.foreversixty.app.auth.session.me

/**
 * The app's object graph, built once per process and passed down explicitly. Feature lanes add
 * their `XxxApi(apiClient)` here. A DI framework can replace this when the graph outgrows it.
 */
class AppGraph(context: Context, val config: ApiConfig = ApiConfig.fromBuildConfig()) {
    /** The device token, reporting every 401 discard to [session] (lane F1). */
    private val sessionTokens = SessionTokens(AndroidTokenStore.create(context))
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
            engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)),
            identity = identity,
        ),
        updateGate,
    )

    val accountApi = AccountApi(apiClient)

    /** Read at launch (MainActivity) and cached; feature lanes ask it `isEnabled(feature)`. */
    val appConfig = AppConfigRepository(apiClient, FileAppConfigCache(File(context.filesDir, "app-config.json")))

    /** Whether a feature's screens show: on in debug and fixture builds, else as the config says. */
    fun isAvailable(feature: Feature): Boolean = appConfig.isAvailable(feature, config.mode, BuildConfig.DEBUG)

    /** Lane F6: the Talents screen pushed from Home; null while its switch is off. */
    private val talents: TalentsFeature? = if (isAvailable(Feature.Talents)) {
        TalentsFeature.create(context, ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)), appConfig)
    } else {
        null
    }

    /** Lane F8: whether See plans may link to web checkout here (decision P, fail closed). */
    val plansLink = PlansLink(appConfig, DebugStorefront.reader(PlayStorefront(context))) { isAvailable(Feature.PremiumLink) }

    /** Lane F2: the home answer, its static item data and its cache. */
    val home = HomeDependencies.create(
        context,
        config,
        accountApi,
        apiClient,
        talents,
        isSimsOn = { isAvailable(Feature.Sims) },
        isReportsOn = { isAvailable(Feature.Reports) },
        mayLinkToPlans = {
            plansLink.resolve()
            plansLink.mayLink.value
        },
    ) {
        isAvailable(Feature.CharacterHome)
    }

    /** Lane F4: guild raids and sign-up; their own static client reads the active build for icons. */
    val raids = RaidsFeature(
        context = context,
        apiClient = apiClient,
        statics = StaticDataClient(ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets))),
    ) { isAvailable(Feature.Raids) }

    val signIn = AppSignIn(
        baseUrl = config.baseUrl,
        client = apiClient,
        tokens = tokens,
        deviceName = Build.MODEL,
    )

    /** The account (lane F1): signed in or out, the current character, sign-in and sign-out. */
    val session = SessionManager(
        client = apiClient,
        tokens = sessionTokens,
        signInFlow = signIn,
        picked = PrefsValue(context, PrefsValue.PICKED_CHARACTER),
        scope = MainScope(),
        clearCaches = {
            home.cache.removeAll()
            raids.cache.removeAll()
            home.sims?.cache?.removeAll()
            home.reports?.cache?.removeAll()
            widgets?.signOut()
        },
    )

    /**
     * Lane F14: the widgets, the tile and the ongoing notifications, all from one snapshot; null
     * while the `widgets` switch is off (WidgetComponents then keeps the widgets out of the picker).
     */
    val widgets: WidgetsFeature? = if (isAvailable(Feature.Widgets)) {
        WidgetsFeature(
            context = context.applicationContext,
            apiClient = apiClient,
            raids = raids.service,
            signedIn = { tokens.bearerToken() != null },
            currentCharacter = { session.currentCharacter()?.key ?: session.pickedKey.value },
            serverSims = { session.phase.value.me?.entitlements?.serverSims == true },
        )
    } else {
        null
    }.also { feature -> home.sims?.liveProgress = feature?.let { it::onSimProgress } }

    /** Fixture builds sign in against the contract fixtures; live builds open a Custom Tab. */
    val account = AccountDeps(
        session = session,
        browser = { activity ->
            when (config.mode) {
                ApiMode.Fixture -> FixtureBrowser { session.handleRedirect(it) }
                ApiMode.Live -> CustomTabsLauncher(activity)
            }
        },
        versionName = identity.version,
        versionCode = identity.build,
        dataBuild = { appConfig.current?.activeBuild },
        plansLink = plansLink,
        liveSurfaces = widgets?.settings,
    )

    /**
     * Lane F9: the Guides tab and its offline guides; null while the `guides` switch is off. Tier
     * lists show only with `guide_tiers` on.
     */
    val guides: GuidesFeature? = if (isAvailable(Feature.Guides)) {
        GuidesFeature.create(
            context = context,
            engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)),
            tiersOn = { isAvailable(Feature.GuideTiers) },
            characters = { session.phase.value.me?.characters.orEmpty() },
        )
    } else {
        null
    }

    /**
     * BiS (lane F3): the static files through their offline cache; null while the switch is off.
     * Its character is the session's current one (lane F1), so the list follows the selector.
     */
    val bisModel: BisModel? = if (isAvailable(Feature.Bis)) {
        createBisModel(
            context = context,
            engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)),
            appConfig = appConfig,
            character = { session.currentCharacter() },
        )
    } else {
        null
    }
}

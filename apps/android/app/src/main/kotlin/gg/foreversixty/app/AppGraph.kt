package gg.foreversixty.app

import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.rankings.logic.LiveRankingsService
import gg.foreversixty.app.items.ItemIconLoader
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.system.SystemFeature
import gg.foreversixty.app.bis.staticDataStore
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
import gg.foreversixty.app.bis.createBisDataSource
import gg.foreversixty.app.bis.createBisModel
import gg.foreversixty.app.search.SearchFeature
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
import gg.foreversixty.app.planner.PlannerFeature
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.UpdateGate
import gg.foreversixty.app.raids.RaidsFeature
import gg.foreversixty.app.raids.logic.LiveRaidService
import gg.foreversixty.app.auth.EncryptedSecretStore
import gg.foreversixty.app.guild.GuildExtrasFeature
import gg.foreversixty.app.update.InstallId
import gg.foreversixty.app.widgets.WidgetsFeature
import gg.foreversixty.app.auth.session.me
import gg.foreversixty.app.auth.session.isSignedIn

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

    /**
     * Lane F15: the active build fresh from the site (never a saved copy), for the new-data strip.
     * A newer one than the config's refreshes the config, so an `Update` reloads on it.
     */
    suspend fun latestBuild(): String? {
        val answer = statics.load(StaticEndpoints.getActiveBuild()) as? ApiResult.Success ?: return null
        val build = answer.data.takeUnless { it.isFromCache }?.value?.build ?: return null
        if (build != appConfig.current?.activeBuild) appConfig.refresh()
        return build
    }

    /** Whether a feature's screens show: on in debug and fixture builds, else as the config says. */
    fun isAvailable(feature: Feature): Boolean = appConfig.isAvailable(feature, config.mode, BuildConfig.DEBUG)

    /**
     * The one static data store (lanes F12 and F15): guides, BiS lists, talents, the tooltip's files,
     * the search index and the item pages, saved in one folder for offline use until a newer build
     * swaps them out; sized and removed by You › Offline data; cleared at sign-out.
     */
    val statics: StaticDataStore = staticDataStore(context, ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)))

    /**
     * Lane F13: the talent planner (editor, shared builds, Your builds); null while `planner` is off.
     * Its sign-in and character are read from [session] when asked, never at build time.
     */
    val planner: PlannerFeature? = if (isAvailable(Feature.Planner)) {
        PlannerFeature.create(
            context = context,
            engine = ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets)),
            appConfig = appConfig,
            api = apiClient,
            signedIn = { session.phase.value.isSignedIn },
            character = { session.currentCharacter() },
        )
    } else {
        null
    }

    /** Lane F6: the Talents screen pushed from Home; null while its switch is off. */
    private val talents: TalentsFeature? = if (isAvailable(Feature.Talents)) {
        TalentsFeature.create(context, statics, appConfig, planner)
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
        statics,
        isSimsOn = { isAvailable(Feature.Sims) },
        isReportsOn = { isAvailable(Feature.Reports) },
        isRankingsOn = { isAvailable(Feature.Rankings) },
        mayLinkToPlans = {
            plansLink.resolve()
            plansLink.mayLink.value
        },
    ) {
        isAvailable(Feature.CharacterHome)
    }

    /** Lane F16: roster, raid history, calendar sync and link; null while `guild_extras` is off. */
    val guildExtras: GuildExtrasFeature? = if (isAvailable(Feature.GuildExtras)) {
        GuildExtrasFeature(
            context = context,
            apiClient = apiClient,
            raids = LiveRaidService(apiClient),
            secrets = EncryptedSecretStore.create(context, CALENDAR_LINK_FILE, CALENDAR_LINK_PURPOSE),
        )
    } else {
        null
    }

    /** Lane F4: guild raids and sign-up; their own static client reads the active build for icons. */
    val raids = RaidsFeature(
        context = context,
        apiClient = apiClient,
        statics = StaticDataClient(ForeverHttpClient.engineFor(config, AssetFixtureSource(context.assets))),
        extras = guildExtras,
    ) { isAvailable(Feature.Raids) }

    /**
     * Lane F12: universal search and the item page, over [statics]; the build is the one the BiS
     * tab reads (the app config's, else the bundled copy).
     */
    val search = SearchFeature.create(
        context = context,
        apiClient = apiClient,
        statics = statics,
        catalog = home.catalog,
        activeBuild = createBisDataSource(context, statics, appConfig)::activeBuild,
    ) { isAvailable(Feature.Search) }.also { search ->
        // Lane F15: F11.7's `Search again` opens Search with the name in the field.
        home.rankings?.let { rankings ->
            rankings.searchAgain = { name ->
                rankings.navigator.character = null
                if (search.enabled) search.navigator.openSearch(name)
            }
        }
    }

    val signIn = AppSignIn(
        baseUrl = config.baseUrl,
        client = apiClient,
        tokens = tokens,
        deviceName = Build.MODEL,
    )

    /** The account (lane F1): signed in or out, the current character, sign-in and sign-out. */
    val session: SessionManager = SessionManager(
        client = apiClient,
        tokens = sessionTokens,
        signInFlow = signIn,
        picked = PrefsValue(context, PrefsValue.PICKED_CHARACTER),
        scope = MainScope(),
        clearCaches = {
            home.cache.removeAll()
            raids.cache.removeAll()
            home.sims?.cache?.removeAll()
            statics.deleteAll()
            search.clear()
            home.reports?.cache?.removeAll()
            home.rankings?.cache?.removeAll()
            widgets?.signOut()
            guildExtras?.clear()
            system.signOut()
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
        guildExtras = guildExtras,
        offlineData = { system.offlineData().takeIf { system.enabled } },
    )

    /**
     * Lane F9: the Guides tab and its offline guides; null while the `guides` switch is off. Tier
     * lists show only with `guide_tiers` on.
     */
    val guides: GuidesFeature? = if (isAvailable(Feature.Guides)) {
        GuidesFeature.create(
            context = context,
            statics = statics,
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
            statics = statics,
            appConfig = appConfig,
            character = { session.currentCharacter() },
        )
    } else {
        null
    }

    /**
     * Lane F15: the launcher shortcuts, the Assistant routes, You › Offline data and `Sign up In`.
     * Always built (the static shortcuts are always declared); the `system` switch gates what it does.
     */
    val system: SystemFeature = SystemFeature(
        context = context,
        statics = statics,
        guides = guides,
        raids = raids.service,
        watching = { (LiveRankingsService(apiClient).watching() as? ApiResult.Success)?.data?.rows.orEmpty() },
        signedIn = { tokens.bearerToken() != null },
        currentCharacter = { session.currentCharacter() },
        connectivity = Connectivity.of(context),
        isOn = { isAvailable(Feature.System) },
    ).also { system ->
        ItemIconLoader.mayDownload = system::mayDownload
        guides?.mayPrefetch = system::mayPrefetch
    }
}

/** The calendar link's sealed store (W3-C17): beside the device token, under the same keyset. */
private const val CALENDAR_LINK_FILE = "datastore/calendar_link.bin"
private const val CALENDAR_LINK_PURPOSE = "gg.foreversixty.app/calendar-link/v1"

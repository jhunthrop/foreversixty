package gg.foreversixty.app.raids

import android.Manifest
import android.content.Context
import androidx.activity.compose.BackHandler
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateListOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.platform.LocalView
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.guild.GuildExtrasFeature
import gg.foreversixty.app.guild.GuildExtrasRows
import gg.foreversixty.app.guild.HistoryRoute
import gg.foreversixty.app.guild.RosterRoute
import gg.foreversixty.app.guild.logic.CalendarInput
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.dataOrNull
import gg.foreversixty.app.raids.logic.FileRaidCache
import gg.foreversixty.app.raids.logic.LiveRaidService
import gg.foreversixty.app.raids.logic.PlannedRaids
import gg.foreversixty.app.raids.logic.PrefsReminderAskStore
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidNightModel
import gg.foreversixty.app.raids.logic.RaidOpenAction
import gg.foreversixty.app.raids.logic.RaidPageModel
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.logic.RaidsTabModel
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.logic.ReminderAskStore
import gg.foreversixty.app.raids.logic.SignupController
import gg.foreversixty.app.raids.logic.SignupState
import gg.foreversixty.app.raids.logic.SystemNotificationPermission
import gg.foreversixty.app.raids.logic.ViewRaidHaptics
import gg.foreversixty.app.raids.ui.LocalRaidIconBuild
import gg.foreversixty.app.raids.ui.RaidNightScreen
import gg.foreversixty.app.raids.ui.RaidPageScreen
import gg.foreversixty.app.raids.ui.RaidsTabContentView
import java.io.File
import java.time.Instant
import kotlin.time.Duration.Companion.seconds
import kotlinx.coroutines.launch

/** A screen pushed on the Raids stack. */
sealed interface RaidsDestination {
    data class Raid(val guild: Long, val raid: Long, val roster: Boolean = false, val open: RaidOpenAction? = null) : RaidsDestination
    data class Night(val raid: Long) : RaidsDestination

    /** Lane F16: the guild's roster and its logged raid nights. */
    data class Roster(val guild: Long) : RaidsDestination
    data class History(val guild: Long) : RaidsDestination
}

/** The Raids tab's own back stack; deep links and notifications push here (spec 5.0, 6.4). */
class RaidsNavigator {
    val stack = mutableStateListOf<RaidsDestination>()

    /** Set when a link asks for the Raids tab; the scaffold selects it and clears this. */
    var wantsTab by mutableStateOf(false)

    fun push(destination: RaidsDestination) {
        stack.add(destination)
    }

    fun pop() {
        stack.removeLastOrNull()
    }

    /** Opens a raid link on a fresh stack, so back goes to the tab's root. False for other links. */
    fun open(link: DeepLink): Boolean {
        val destinations = when (link) {
            is DeepLink.Raid -> {
                val press = if (link.spec == "ask") SignupState.In else SignupState.of(link.state)
                listOf(RaidsDestination.Raid(link.guild, link.raid, open = press?.let(::RaidOpenAction)))
            }
            is DeepLink.RaidNight -> listOf(RaidsDestination.Raid(link.guild, link.raid), RaidsDestination.Night(link.raid))
            is DeepLink.Guild -> emptyList()
            else -> return false
        }
        stack.clear()
        stack.addAll(destinations)
        wantsTab = true
        return true
    }
}

/** The raids feature's dependencies, built once in the app graph. */
class RaidsFeature(
    context: Context,
    apiClient: ApiClient,
    private val statics: StaticDataClient,
    /** Lane F16: roster, raid history and calendar sync; null while `guild_extras` is off. */
    val extras: GuildExtrasFeature? = null,
    private val isAvailable: () -> Boolean,
) {
    val service: RaidService = LiveRaidService(apiClient)
    val cache: RaidCache = FileRaidCache(File(context.filesDir, "raids"))
    val askStore: ReminderAskStore = PrefsReminderAskStore(context.getSharedPreferences("forever_raids", Context.MODE_PRIVATE))
    val navigator = RaidsNavigator()
    val now: () -> Instant = Instant::now

    /** The Raids tab ships dark behind `raids` (AppGraph.isAvailable: on in debug and fixture builds). */
    val enabled: Boolean get() = isAvailable()

    suspend fun iconBuild(): String? = statics.activeBuild().dataOrNull()?.build
}

/**
 * F4: the Raids tab root and its stack. [characterKey] is the session's current character (lane
 * F1), null to follow `/v1/me`'s main; [selector] is the character selector row over the root.
 */
@Composable
fun RaidsRoute(feature: RaidsFeature, characterKey: String? = null, selector: @Composable () -> Unit = {}) {
    val context = LocalContext.current
    val view = LocalView.current
    val uri = LocalUriHandler.current
    val scope = rememberCoroutineScope()
    val res = androidx.compose.ui.platform.LocalResources.current
    val clock = remember(res) { RaidClock(res) }
    val permissionLauncher = rememberLauncherForActivityResult(ActivityResultContracts.RequestPermission()) { }
    val requestPermission = {
        // Lane F7 registers the FCM token once permission is granted.
        if (android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.TIRAMISU) {
            permissionLauncher.launch(Manifest.permission.POST_NOTIFICATIONS)
        }
    }
    val deps = remember(feature) {
        SignupController.Dependencies(
            service = feature.service,
            res = res,
            haptics = ViewRaidHaptics(view),
            permission = SystemNotificationPermission(context),
            askStore = feature.askStore,
            now = feature.now,
            scope = scope,
            hold = HOLD,
        )
    }
    val tab = remember(feature) { RaidsTabModel(feature.service, feature.cache, res, feature.now) }
    var next by remember { mutableStateOf<SignupController?>(null) }
    var iconBuild by remember { mutableStateOf<String?>(null) }

    val extras = feature.extras

    /** Lane F16: the calendar follows every refresh and every sign-up answer. */
    fun syncCalendar() {
        val loaded = (tab.state as? RaidsTabState.Loaded)?.content ?: return
        if (loaded.offlineSince != null || extras == null) return
        scope.launch { extras.syncCalendar(CalendarInput(loaded.head.id, loaded.head.name, loaded.planned)) }
    }

    val replaced: (gg.foreversixty.app.model.RaidView) -> Unit = { view ->
        tab.replace(view)
        syncCalendar()
    }

    fun syncNext() {
        val content = (tab.state as? RaidsTabState.Loaded)?.content
        val raid = content?.let { PlannedRaids(it.planned, feature.now()).next }
        val current = next
        next = when {
            content == null || raid == null -> null
            current != null && current.raid.id == raid.id -> current.also {
                it.replace(raid)
                it.offline = content.offlineSince != null
            }
            else -> SignupController(content.head.id, raid, deps).also {
                it.offline = content.offlineSince != null
                it.onRaidChanged = replaced
            }
        }
    }

    suspend fun reload() {
        tab.load()
        syncNext()
        syncCalendar()
        val loaded = (tab.state as? RaidsTabState.Loaded)?.content
        if (extras != null && loaded != null) extras.roster(loaded.head.id).load()
    }

    LaunchedEffect(feature) { iconBuild = feature.iconBuild() }
    LaunchedEffect(feature, characterKey) {
        tab.currentKey = characterKey
        reload()
    }
    val content = (tab.state as? RaidsTabState.Loaded)?.content
    BackHandler(enabled = feature.navigator.stack.isNotEmpty()) { feature.navigator.pop() }
    CompositionLocalProvider(LocalRaidIconBuild provides iconBuild) {
        when (val top = feature.navigator.stack.lastOrNull()) {
            null -> Column(Modifier.fillMaxSize().background(FsColor.bg)) {
                selector()
                RaidsTabContentView(
                    state = tab.state,
                    clock = clock,
                    now = feature.now,
                    nextController = next,
                    openRaid = { raid, roster -> feature.navigator.push(RaidsDestination.Raid(content?.head?.id ?: 0, raid.id, roster)) },
                    openNight = { feature.navigator.push(RaidsDestination.Night(it)) },
                    openWeb = { uri.openUri(it) },
                    retry = { scope.launch { reload() } },
                    requestPermission = requestPermission,
                    guildRows = { loaded ->
                        if (extras != null) {
                            GuildExtrasRows(
                                onRoster = extras.roster(loaded.head.id).onRoster,
                                openRoster = { feature.navigator.push(RaidsDestination.Roster(loaded.head.id)) },
                                openHistory = { feature.navigator.push(RaidsDestination.History(loaded.head.id)) },
                            )
                        }
                    },
                )
            }
            is RaidsDestination.Raid -> {
                val model = remember(top) { RaidPageModel(top.guild, top.raid, deps, feature.cache, top.open, replaced) }
                RaidPageScreen(
                    model = model,
                    clock = clock,
                    now = feature.now,
                    currentKey = content?.character?.key,
                    guildName = content?.head?.name.orEmpty(),
                    guildUrl = content?.head?.webUrl,
                    openNight = { feature.navigator.push(RaidsDestination.Night(it)) },
                    back = { feature.navigator.pop() },
                    requestPermission = requestPermission,
                )
            }
            is RaidsDestination.Night -> {
                val model = remember(top) { RaidNightModel(top.raid, feature.service, feature.cache, feature.now) }
                RaidNightScreen(model, clock, feature.now) { feature.navigator.pop() }
            }
            is RaidsDestination.Roster -> extras?.let {
                RosterRoute(it.roster(top.guild), clock, feature.now) { feature.navigator.pop() }
            }
            is RaidsDestination.History -> extras?.let {
                val pager = remember(top) { it.history(top.guild) }
                HistoryRoute(pager, tab.accountKeys, clock, feature.now, { raid -> feature.navigator.push(RaidsDestination.Night(raid)) }) {
                    feature.navigator.pop()
                }
            }
        }
    }
}

/** `Saved.`, `Copied` and the lock toast stay 4 s (spec F4.3, F4.18, F4.21). */
private val HOLD = 4.seconds

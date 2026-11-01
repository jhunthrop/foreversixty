package gg.foreversixty.app.guild

import android.content.ClipData
import android.content.ClipDescription
import android.content.Context
import android.content.Intent
import android.os.Build
import android.os.PersistableBundle
import android.provider.Settings
import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Column
import androidx.compose.material3.HorizontalDivider
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.ClipEntry
import androidx.compose.ui.platform.LocalClipboard
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.platform.LocalResources
import androidx.compose.ui.platform.LocalUriHandler
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.pluralStringResource
import androidx.compose.ui.res.stringResource
import androidx.core.net.toUri
import androidx.lifecycle.compose.LifecycleResumeEffect
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.SecretStore
import gg.foreversixty.app.guild.logic.CalendarInput
import gg.foreversixty.app.guild.logic.CalendarLinkModel
import gg.foreversixty.app.guild.logic.CalendarPlanner
import gg.foreversixty.app.guild.logic.CalendarStatus
import gg.foreversixty.app.guild.logic.CalendarSync
import gg.foreversixty.app.guild.logic.GuildService
import gg.foreversixty.app.guild.logic.HistoryPager
import gg.foreversixty.app.guild.logic.HistoryTarget
import gg.foreversixty.app.guild.logic.LiveGuildService
import gg.foreversixty.app.guild.logic.RosterStore
import gg.foreversixty.app.guild.platform.AndroidCalendarWriter
import gg.foreversixty.app.guild.platform.CALENDAR_PERMISSIONS
import gg.foreversixty.app.guild.platform.PrefsCalendarSettingsStore
import gg.foreversixty.app.guild.ui.CalendarActions
import gg.foreversixty.app.guild.ui.CalendarContent
import gg.foreversixty.app.guild.ui.CalendarScreenState
import gg.foreversixty.app.guild.ui.GuildLinkRow
import gg.foreversixty.app.guild.ui.HistoryContent
import gg.foreversixty.app.guild.ui.RosterContent
import gg.foreversixty.app.guild.ui.RosterControls
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.FileRaidCache
import gg.foreversixty.app.raids.logic.RaidCache
import gg.foreversixty.app.raids.logic.RaidClock
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.raids.ui.raidQuietCard
import java.io.File
import java.time.Instant
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

/**
 * F16's dependencies, built once in the app graph while `guild_extras` is on: the roster and
 * history loaders, the calendar sync and the calendar link's secure store.
 */
class GuildExtrasFeature(
    context: Context,
    apiClient: ApiClient,
    private val raids: RaidService,
    private val secrets: SecretStore,
) {
    val service: GuildService = LiveGuildService(apiClient)
    val cache: RaidCache = FileRaidCache(File(context.filesDir, "guild"))
    val now: () -> Instant = Instant::now
    private val app = context.applicationContext

    val calendar = CalendarSync(
        writer = AndroidCalendarWriter(app),
        store = PrefsCalendarSettingsStore(app.getSharedPreferences("forever_calendar", Context.MODE_PRIVATE)),
        planner = CalendarPlanner(app.resources, RaidClock(app.resources)),
        calendarName = app.getString(R.string.calendar_name),
        now = now,
    )

    /** What the switch shows; refreshed after each write and on resume. */
    var calendarStatus by mutableStateOf(CalendarStatus.Off)
        private set

    /** The planned raids the Raids tab loaded last, so a Which or Alert change rewrites at once. */
    private var lastInput: CalendarInput? = null

    private val rosters = mutableMapOf<Long, RosterStore>()

    /** One roster per guild, shared by the Raids tab's row and the Roster screen. */
    fun roster(guild: Long): RosterStore = rosters.getOrPut(guild) { RosterStore(guild, service, cache, now) }

    fun history(guild: Long): HistoryPager = HistoryPager(guild, service, cache, now)

    fun link(accountKey: String): CalendarLinkModel = CalendarLinkModel(service, secrets, accountKey)

    /** Every Raids refresh and sign-up answer (spec F16.4). */
    suspend fun syncCalendar(input: CalendarInput) {
        lastInput = input
        calendarStatus = withContext(Dispatchers.IO) { calendar.sync(input) }
    }

    suspend fun refreshCalendarStatus() {
        calendarStatus = withContext(Dispatchers.IO) { calendar.status() }
    }

    /** The switch: on after access is granted, off removes the calendar. */
    suspend fun setCalendar(on: Boolean, characterKey: String?) {
        calendarStatus = withContext(Dispatchers.IO) {
            if (on) {
                calendar.enable(lastInput ?: loadInput(characterKey))
            } else {
                calendar.disable()
                calendar.status()
            }
        }
    }

    suspend fun calendarDenied() {
        calendarStatus = withContext(Dispatchers.IO) {
            calendar.denied()
            calendar.status()
        }
    }

    suspend fun calendarChange(change: CalendarSync.(CalendarInput?) -> CalendarStatus) {
        calendarStatus = withContext(Dispatchers.IO) { calendar.change(lastInput) }
    }

    /** You › Calendar opened before the Raids tab: the current character's guild raids. */
    private suspend fun loadInput(characterKey: String?): CalendarInput? {
        val me = (raids.me() as? ApiResult.Success)?.data ?: return null
        val character = me.characters.firstOrNull { it.key == characterKey } ?: me.mainCharacter ?: return null
        val guild = character.guild?.takeIf { it.verified } ?: return null
        val planned = (raids.raids(guild.id) as? ApiResult.Success)?.data?.planned ?: return null
        val name = me.guilds.firstOrNull { it.id == guild.id }?.name ?: guild.name
        return CalendarInput(guild.id, name, planned).also { lastInput = it }
    }

    /** Sign-out: the guild caches and the calendar link go with the other caches. */
    suspend fun clear() {
        cache.removeAll()
        rosters.clear()
        lastInput = null
        secrets.clear()
    }
}

/** F16.7: the two rows under the Raids tab's guild head. */
@Composable
fun GuildExtrasRows(onRoster: Int?, openRoster: () -> Unit, openHistory: () -> Unit) {
    Column(Modifier.raidQuietCard()) {
        GuildLinkRow(
            painterResource(R.drawable.ic_guild_roster),
            stringResource(R.string.raids_rosterRow),
            onRoster?.let { pluralStringResource(R.plurals.raids_rosterRow_sub, it, it) },
            openRoster,
        )
        HorizontalDivider(color = FsColor.lineSoft)
        GuildLinkRow(
            painterResource(R.drawable.ic_guild_history),
            stringResource(R.string.raids_historyRow),
            stringResource(R.string.raids_historyRow_sub),
            openHistory,
        )
    }
}

/** F16.1: the roster pushed on the Raids stack; a row opens the character's web page. */
@Composable
fun RosterRoute(store: RosterStore, clock: RaidClock, now: () -> Instant, back: () -> Unit) {
    val uri = LocalUriHandler.current
    val scope = rememberCoroutineScope()
    val controls = remember(store) { RosterControls() }
    LaunchedEffect(store) { store.load() }
    RosterContent(
        state = store.state,
        controls = controls,
        stripTime = { clock.stripTime(it, now()) },
        open = { entry -> entry.webUrl?.let(uri::openUri) },
        retry = { scope.launch { store.load() } },
        back = back,
    )
}

/** F16.3: a night opens Your raid night when you were there, else its report on the web. */
@Composable
fun HistoryRoute(
    pager: HistoryPager,
    mine: Set<String>,
    clock: RaidClock,
    now: () -> Instant,
    openNight: (Long) -> Unit,
    back: () -> Unit,
) {
    val uri = LocalUriHandler.current
    val scope = rememberCoroutineScope()
    LaunchedEffect(pager) { pager.load() }
    HistoryContent(
        state = pager.state,
        more = pager.more,
        mine = mine,
        clock = clock,
        stripTime = { clock.stripTime(it, now()) },
        open = { facts ->
            when (val target = facts.target) {
                is HistoryTarget.Night -> openNight(target.raidId)
                is HistoryTarget.Report -> uri.openUri(target.url)
            }
        },
        loadMore = { scope.launch { pager.loadMore() } },
        retry = { scope.launch { pager.load() } },
        back = back,
    )
}

/**
 * F16.4: the switch asks for calendar access only after the flip; the link loads on open and
 * copies as a sensitive clip. [accountKey] keys the stored link; [characterKey] picks the guild.
 */
@Composable
fun CalendarRoute(feature: GuildExtrasFeature, accountKey: String, characterKey: String?, back: () -> Unit) {
    val context = LocalContext.current
    val clipboard = LocalClipboard.current
    val res = LocalResources.current
    val scope = rememberCoroutineScope()
    val link = remember(accountKey) { feature.link(accountKey) }
    val permissions = rememberLauncherForActivityResult(ActivityResultContracts.RequestMultiplePermissions()) { granted ->
        scope.launch {
            if (granted.values.all { it }) feature.setCalendar(true, characterKey) else feature.calendarDenied()
        }
    }
    LaunchedEffect(link) { link.load() }
    LifecycleResumeEffect(feature) {
        scope.launch { feature.refreshCalendarStatus() }
        onPauseOrDispose { }
    }
    CalendarContent(
        state = CalendarScreenState(
            status = feature.calendarStatus,
            settings = feature.calendar.settings,
            link = link.state,
            linkNotice = link.notice,
            offline = link.offline,
            resetting = link.resetting,
        ),
        actions = CalendarActions(
            onSwitch = { on ->
                if (on) permissions.launch(CALENDAR_PERMISSIONS) else scope.launch { feature.setCalendar(false, characterKey) }
            },
            onWhich = { which -> scope.launch { feature.calendarChange { setWhich(which, it) } } },
            onAlert = { alert -> scope.launch { feature.calendarChange { setAlert(alert, it) } } },
            openSettings = { context.startActivity(appSettings(context)) },
            copy = {
                link.copy()?.let { url -> scope.launch { clipboard.setClipEntry(ClipEntry(sensitiveClip(res.getString(R.string.calendar_link), url))) } }
            },
            reset = { scope.launch { link.reset() } },
        ),
        back = back,
    )
}

/** The url never shows in the clipboard preview (Android 13+ reads the extra). */
private fun sensitiveClip(label: String, url: String): ClipData = ClipData.newPlainText(label, url).apply {
    description.extras = PersistableBundle().apply {
        val key = if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) ClipDescription.EXTRA_IS_SENSITIVE else SENSITIVE_EXTRA
        putBoolean(key, true)
    }
}

private fun appSettings(context: Context) =
    Intent(Settings.ACTION_APPLICATION_DETAILS_SETTINGS, "package:${context.packageName}".toUri())

/** ClipDescription.EXTRA_IS_SENSITIVE's value, for Android 12 and below. */
private const val SENSITIVE_EXTRA = "android.content.extra.IS_SENSITIVE"

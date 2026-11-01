package gg.foreversixty.app.system

import gg.foreversixty.app.model.WidgetNextRaid
import android.content.Context
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.common.Connectivity
import gg.foreversixty.app.guides.GuidesFeature
import gg.foreversixty.app.items.ItemIconLoader
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.WatchedCharacter
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.system.logic.AppShortcut
import gg.foreversixty.app.system.logic.DynamicShortcuts
import gg.foreversixty.app.system.logic.ShortcutSignup
import gg.foreversixty.app.system.offline.FolderIconCache
import gg.foreversixty.app.system.offline.OfflineDataModel
import gg.foreversixty.app.system.offline.PrefsOfflineSettings
import gg.foreversixty.app.system.offline.UpdateOutcome
import gg.foreversixty.app.system.platform.DynamicShortcutPublisher

/** Where a launcher shortcut or the Assistant asked the app to go; the app shell reads and clears it. */
class SystemNavigator {
    var wantsTab: ForeverTab? by mutableStateOf(null)

    /** The tab a shortcut opens; Search covers the current tab instead (its own navigator opens it). */
    fun open(shortcut: AppShortcut) {
        wantsTab = tabOf(shortcut)
    }

    companion object {
        fun tabOf(shortcut: AppShortcut): ForeverTab? = when (shortcut) {
            AppShortcut.NextRaid, AppShortcut.SignupIn -> ForeverTab.Raids
            AppShortcut.Upgrades -> ForeverTab.Home
            AppShortcut.Bis -> ForeverTab.Bis
            AppShortcut.Guides -> ForeverTab.Guides
            AppShortcut.Search -> null
        }
    }
}

/**
 * Lane F15's dependencies, built once in the app graph. The static shortcuts (shortcuts.xml) are
 * always declared; [isOn] (the `system` switch) gates the dynamic shortcuts, the Assistant routes,
 * You › Offline data and `Sign up In`'s answer (off, it opens the app).
 */
class SystemFeature(
    context: Context,
    private val statics: StaticDataStore,
    private val guides: GuidesFeature?,
    private val raids: RaidService,
    private val watching: suspend () -> List<WatchedCharacter>,
    /** `GET /v1/me/widget`'s next raid for a character (W3-C12). */
    private val nextRaid: suspend (String?) -> WidgetNextRaid?,
    private val signedIn: suspend () -> Boolean,
    private val currentCharacter: () -> AccountCharacter?,
    val connectivity: Connectivity,
    private val isOn: () -> Boolean,
) {
    private val app = context.applicationContext
    val navigator = SystemNavigator()
    val offlineSettings = PrefsOfflineSettings(app.getSharedPreferences(PrefsOfflineSettings.FILE, Context.MODE_PRIVATE))
    private val unmetered = Connectivity.unmetered(app)
    private val publisher = DynamicShortcutPublisher(app)

    val enabled: Boolean get() = isOn()

    val signup = ShortcutSignup(raids, signedIn, { currentCharacter()?.key }, isOn)

    /** Icons and new guides wait for an unmetered network while `Download on Wi-Fi only` is on. */
    fun mayDownload(): Boolean = !isOn() || !offlineSettings.wifiOnly || unmetered()

    /** The guides of the player's classes are kept ready while the switch is on and the network allows. */
    fun mayPrefetch(): Boolean = !isOn() || (offlineSettings.keepClassesReady && mayDownload())

    fun offlineData(): OfflineDataModel = OfflineDataModel(
        statics = statics,
        icons = FolderIconCache(ItemIconLoader.folder(app)) { ItemIconLoader.clear(app) },
        settings = offlineSettings,
        connectivity = connectivity,
        update = ::updateNow,
    )

    /** `Update now`: the active build, its swap, the open guides' update and the player's classes saved again. */
    private suspend fun updateNow(): UpdateOutcome {
        val before = statics.cached(StaticEndpoints.getActiveBuild())?.value?.build
        val answer = statics.load(StaticEndpoints.getActiveBuild())
        val fresh = (answer as? ApiResult.Success)?.data?.takeUnless { it.isFromCache } ?: return UpdateOutcome.Failed
        val build = fresh.value.build
        guides?.let { feature ->
            feature.sync(force = true)
            if (feature.store.hasNewData) feature.store.applyUpdate()
        }
        offlineSettings.checkedAt = java.time.Instant.now()
        if (before == null || before == build) return UpdateOutcome.UpToDate
        statics.swap(build)
        return UpdateOutcome.Updated
    }

    /** Every foreground and sign-in: the dynamic shortcuts follow the data; off or signed out, none. */
    suspend fun refreshShortcuts() {
        if (!isOn() || !signedIn()) {
            publisher.clear()
            return
        }
        val current = currentCharacter()
        val watched = watching()
        // The same next raid the widget, voice answers and Spotlight read (one story, one clock).
        val next = nextRaid(current?.key)?.takeUnless { it.status == CANCELLED }
        publisher.publish(DynamicShortcuts.list(current, watched, next, app.resources))
    }

    private companion object {
        const val CANCELLED = "cancelled"
    }

    /** Sign-out: the dynamic shortcuts go with the account. */
    fun signOut() = publisher.clear()
}

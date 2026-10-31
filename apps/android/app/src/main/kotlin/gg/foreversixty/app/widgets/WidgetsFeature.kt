package gg.foreversixty.app.widgets

import android.content.ComponentName
import android.content.Context
import android.service.quicksettings.TileService
import androidx.glance.appwidget.updateAll
import gg.foreversixty.app.model.SimProgress
import gg.foreversixty.app.model.SimRow
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.raids.logic.RaidService
import gg.foreversixty.app.widgets.data.LiveKind
import gg.foreversixty.app.widgets.data.LiveWidgetService
import gg.foreversixty.app.widgets.data.WidgetRefresher
import gg.foreversixty.app.widgets.data.WidgetSignup
import gg.foreversixty.app.widgets.data.WidgetStore
import gg.foreversixty.app.widgets.glance.NextRaidWidget
import gg.foreversixty.app.widgets.glance.NextUpgradeWidget
import gg.foreversixty.app.widgets.live.LiveNotifier
import gg.foreversixty.app.widgets.live.LiveRegistrations
import gg.foreversixty.app.widgets.logic.LivePush
import gg.foreversixty.app.widgets.logic.RaidLiveContent
import gg.foreversixty.app.widgets.logic.SimLive
import gg.foreversixty.app.widgets.logic.SimLiveContent
import gg.foreversixty.app.widgets.logic.SimThrottle
import gg.foreversixty.app.widgets.logic.WidgetFigures
import gg.foreversixty.app.widgets.tile.NextRaidTileService
import java.time.Duration
import java.time.Instant

/**
 * Lane F14's dependencies, built once in the app graph while the `widgets` switch is on: the shared
 * snapshot store, its refresher, the widget sign-up, and the ongoing notifications. Every surface
 * redraws from the one snapshot through [WidgetFigures] ([redraw]).
 */
class WidgetsFeature(
    private val context: Context,
    apiClient: ApiClient,
    raids: RaidService,
    signedIn: suspend () -> Boolean,
    currentCharacter: () -> String?,
    /** `entitlements.server_sims`: the sim surface exists only for a server run (spec §1.1). */
    private val serverSims: () -> Boolean,
    private val now: () -> Instant = Instant::now,
) {
    val store = WidgetStore(context)
    private val service = LiveWidgetService(apiClient)
    private val notifier = LiveNotifier(context)
    private val registrations = LiveRegistrations(service, store)
    private val throttle = SimThrottle()

    val refresher = WidgetRefresher(service, store, signedIn, currentCharacter, now) { redraw() }
    val signup = WidgetSignup(raids, store, refresher, now, redraw = { redraw() })
    val settings = LiveSurfaceSettings(store, serverSims) { syncRaidLive() }

    fun figures(): WidgetFigures = WidgetFigures(context.resources)

    /** Redraws the widgets, the tile and the raid countdown from the stored snapshot. */
    suspend fun redraw() {
        NextRaidWidget().updateAll(context)
        NextUpgradeWidget().updateAll(context)
        TileService.requestListeningState(context, ComponentName(context, NextRaidTileService::class.java))
        syncRaidLive()
    }

    /** Shows the raid countdown inside its window (In or Late, switch on), else removes it. */
    suspend fun syncRaidLive() {
        val at = now()
        val figures = figures().raid(store.snapshot(), at, store.overlay())
        val content = RaidLiveContent.of(figures, context.resources)?.takeIf { store.liveSetting(LiveKind.Raid) }
        if (content == null) {
            notifier.cancelRaid()
            registrations.end(LiveKind.Raid)
        } else {
            notifier.showRaid(content)
            if (content.cancelled) registrations.end(LiveKind.Raid) else registrations.start(LiveKind.Raid, content.raidId.toString())
        }
        figures.liveTarget?.minus(WidgetFigures.LIVE_LEAD)?.takeIf { it > at }?.let {
            WidgetRefreshWork.atWindow(context, Duration.between(at, it))
        }
    }

    /** A server run's poll answer while the app runs (Sims lane's poll, at most every 15 s). */
    suspend fun onSimProgress(row: SimRow, progress: SimProgress, title: String) {
        val id = row.simId ?: return
        val state = when (progress.state) {
            "done" -> SimLive.State.Finished
            "error" -> SimLive.State.Failed
            else -> SimLive.State.Running
        }
        val total = progress.combosTotal.takeIf { it > 1 }
        showSim(SimLive(id, title, state, progress.combosDone, total, null, progress.dps))
    }

    private suspend fun showSim(sim: SimLive) {
        if (!serverSims() || !store.liveSetting(LiveKind.Sim) || !throttle.allows(sim, now())) return
        notifier.showSim(sim, SimLiveContent.of(sim, context.resources))
        if (sim.state == SimLive.State.Running) registrations.start(LiveKind.Sim, sim.simId) else registrations.end(LiveKind.Sim)
    }

    /** A live-surface FCM data message (fixtures surface.*.fcm). */
    suspend fun handlePush(data: Map<String, String>) {
        when (val push = LivePush.parse(data) ?: return) {
            LivePush.WidgetRefresh -> refresher.refresh()
            is LivePush.Raid -> refresher.refresh()
            is LivePush.Sim -> showSim(push.sim)
        }
    }

    /** Sign-out: every surface returns to its signed-out words and nothing stays registered. */
    suspend fun signOut() {
        registrations.endAll()
        notifier.cancelRaid()
        refresher.clear()
        WidgetRefreshWork.cancel(context)
    }
}

/** You › Live surfaces (spec §5.6 Rules): two switches, both on by default. */
class LiveSurfaceSettings(
    private val store: WidgetStore,
    private val serverSims: () -> Boolean,
    private val onChanged: suspend () -> Unit,
) {
    /** The sim switch shows only to an account that can run on the servers. */
    val showsSim: Boolean get() = serverSims()

    suspend fun isOn(kind: LiveKind): Boolean = store.liveSetting(kind)

    suspend fun set(kind: LiveKind, on: Boolean) {
        store.setLiveSetting(kind, on)
        onChanged()
    }
}

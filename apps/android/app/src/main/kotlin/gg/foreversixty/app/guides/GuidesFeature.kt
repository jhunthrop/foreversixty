package gg.foreversixty.app.guides

import android.content.Context
import gg.foreversixty.app.PrefsValue
import gg.foreversixty.app.auth.session.StoredValue
import gg.foreversixty.app.guides.logic.GuideFiles
import gg.foreversixty.app.guides.logic.GuidesNavigator
import gg.foreversixty.app.guides.logic.GuidesStore
import gg.foreversixty.app.guides.logic.LiveGuideFiles
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.network.StaticDataStore
import java.time.Duration
import java.time.Instant
import java.time.ZoneId

/**
 * F9's dependencies, built once in the app graph ([create]); tests pass fixture files. The tab ships
 * dark behind `guides`; the Tier lists segment behind `guide_tiers` ([tiersOn]).
 */
class GuidesFeature(
    val files: GuideFiles,
    /** The segment the tab opens on, per device (spec F9.1). */
    val segment: StoredValue,
    private val tiersOn: () -> Boolean,
    /** The account's characters: their classes' guides are kept ready offline. */
    private val characters: () -> List<AccountCharacter>,
    val now: () -> Instant = Instant::now,
    val zone: ZoneId = ZoneId.systemDefault(),
) {
    val navigator = GuidesNavigator()
    val store = GuidesStore(files)
    private var lastSync: Instant? = null

    val showsTiers: Boolean get() = tiersOn()

    /**
     * On launch and on foreground after an hour (spec F9, Sync): read the active build (a new one
     * waits behind the strip) and save the guides of every class the player has a character in.
     */
    suspend fun sync() {
        val at = now()
        val last = lastSync
        if (last != null && Duration.between(last, at) < SYNC_INTERVAL) return
        lastSync = at
        if (store.build == null) store.start() else store.checkForNewBuild()
        store.prefetch(characters().mapNotNull { it.className }.toSet())
    }

    companion object {
        private val SYNC_INTERVAL: Duration = Duration.ofHours(1)

        fun create(context: Context, statics: StaticDataStore, tiersOn: () -> Boolean, characters: () -> List<AccountCharacter>): GuidesFeature {
            return GuidesFeature(
                files = LiveGuideFiles(statics),
                segment = PrefsValue(context, PrefsValue.GUIDES_SEGMENT),
                tiersOn = tiersOn,
                characters = characters,
            )
        }
    }
}

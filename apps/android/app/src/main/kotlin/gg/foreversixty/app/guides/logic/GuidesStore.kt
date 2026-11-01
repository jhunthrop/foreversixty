package gg.foreversixty.app.guides.logic

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import gg.foreversixty.app.model.ActiveBuild
import gg.foreversixty.app.model.Changelog
import gg.foreversixty.app.model.GuideFile
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.KeyDate
import gg.foreversixty.app.model.TiersFile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.StaticEndpoints
import java.io.IOException

/**
 * The Guides tab's data (spec F9, Sync with the nightly data): the build on screen, its index and
 * what is saved, the tier lists and What changed. A newer build found while the tab is open waits
 * in [newBuild] behind the new-data strip until [applyUpdate].
 */
class GuidesStore(private val files: GuideFiles) {
    /** The build every screen of the tab reads; null until the active build is known. */
    var build: String? by mutableStateOf(null)
        private set

    /** A newer active build, found while [build]'s data is on screen. */
    var newBuild: String? by mutableStateOf(null)
        private set

    var updateFailed: Boolean by mutableStateOf(false)
        private set

    var index: Loadable<GuideIndex> by mutableStateOf(Loadable.Loading)
        private set

    /** The saved guides of [build] by spec id, with their sizes. */
    var saved: Map<String, Long> by mutableStateOf(emptyMap())
        private set

    var tiers: Loadable<TiersFile> by mutableStateOf(Loadable.Loading)
        private set

    var changelog: Loadable<Changelog> by mutableStateOf(Loadable.Loading)
        private set

    var dates: Loadable<List<KeyDate>> by mutableStateOf(Loadable.Loading)
        private set

    val hasNewData: Boolean get() = Guides.isNewData(build, newBuild)

    /** Opening the tab: the saved index at once, then the active build and its index. */
    suspend fun start() {
        val savedBuild = files.cached(StaticEndpoints.getActiveBuild())?.value?.build
        if (build == null && savedBuild != null) {
            build = savedBuild
            drawSavedIndex(savedBuild)
        }
        val active = files.load(StaticEndpoints.getActiveBuild())
        val latest = active.freshBuild()
        when {
            latest == null -> Unit
            index is Loadable.Loaded && latest != build -> newBuild = latest
            else -> build = latest
        }
        if (build == null) {
            index = Loadable.of(active.failureOnly(), index)
            return
        }
        loadIndex()
    }

    /** Pull to refresh and Try again: the index of the build on screen. */
    suspend fun loadIndex() {
        val current = build ?: return start()
        index = Loadable.of(files.load(StaticEndpoints.getGuideIndex(current)), index)
        refreshSaved()
    }

    /** On foreground: a newer build waits for Update; nothing on screen changes. */
    suspend fun checkForNewBuild() {
        val latest = files.load(StaticEndpoints.getActiveBuild()).freshBuild() ?: return
        if (build == null) build = latest else if (latest != build) newBuild = latest
    }

    /** Update on the strip: the new build's index; on failure the open build stays and the strip says so. */
    suspend fun applyUpdate(): Boolean {
        val target = newBuild ?: return true
        val answer = files.load(StaticEndpoints.getGuideIndex(target))
        if (answer !is ApiResult.Success || answer.data.isFromCache) {
            updateFailed = true
            return false
        }
        build = target
        newBuild = null
        updateFailed = false
        index = Loadable.of(answer, index)
        refreshSaved()
        return true
    }

    /** Saves every guide of [classSlugs] (the player's classes) in the background. */
    suspend fun prefetch(classSlugs: Set<String>) {
        val current = build ?: return
        val loaded = index.valueOrNull ?: return
        loaded.classes.filter { it.classSlug in classSlugs }.flatMap { it.specs }
            .forEach { files.load(StaticEndpoints.getGuide(current, it.spec)) }
        refreshSaved()
    }

    fun refreshSaved() {
        val current = build ?: return
        val loaded = index.valueOrNull ?: return
        saved = Guides.saved(loaded, current, files)
    }

    suspend fun loadTiers() {
        val current = build ?: return
        if (tiers == Loadable.Loading) Loadable.saved(files.cached(StaticEndpoints.getTiers(current)))?.let { tiers = it }
        tiers = Loadable.of(files.load(StaticEndpoints.getTiers(current)), tiers)
    }

    /** What changed: the key dates and the changelog, each its own region. */
    suspend fun loadChanges() {
        if (dates == Loadable.Loading) Loadable.saved(files.cached(StaticEndpoints.getKeyDates()))?.let { dates = it }
        if (changelog == Loadable.Loading) Loadable.saved(files.cached(StaticEndpoints.getChangelog()))?.let { changelog = it }
        dates = Loadable.of(files.load(StaticEndpoints.getKeyDates()), dates)
        changelog = Loadable.of(files.load(StaticEndpoints.getChangelog()), changelog)
    }

    private fun drawSavedIndex(build: String) {
        Loadable.saved(files.cached(StaticEndpoints.getGuideIndex(build)))?.let { index = it }
        refreshSaved()
    }

    /** The active build from the network; null offline or on failure (a saved copy is not news). */
    private fun ApiResult<CachedFile<ActiveBuild>>.freshBuild(): String? =
        (this as? ApiResult.Success)?.data?.takeUnless { it.isFromCache }?.value?.build

    /** The active build's failure as the index region's: offline is not saved, else the error state. */
    private fun ApiResult<CachedFile<ActiveBuild>>.failureOnly(): ApiResult<CachedFile<GuideIndex>> = when (this) {
        is ApiResult.Failure -> this
        is ApiResult.Success -> ApiResult.Failure(ApiError.Network(IOException(NO_BUILD)))
    }

    private companion object {
        const val NO_BUILD = "no active build"
    }
}

/** One guide (spec F9.2): the saved copy at once (no spinner over cached content), then the network. */
class GuideStore(val spec: String, private val files: GuideFiles) {
    var phase: Loadable<GuideFile> by mutableStateOf(Loadable.Loading)
        private set

    fun drawSaved(build: String) {
        if (phase == Loadable.Loading) Loadable.saved(files.cached(StaticEndpoints.getGuide(build, spec)))?.let { phase = it }
    }

    suspend fun load(build: String) {
        drawSaved(build)
        phase = Loadable.of(files.load(StaticEndpoints.getGuide(build, spec)), phase)
    }

    /** Update on the new-data strip: the guide of the new build replaces the open one. */
    suspend fun reload(build: String) {
        val answer = files.load(StaticEndpoints.getGuide(build, spec))
        if (answer is ApiResult.Success) phase = Loadable.of(answer, phase)
    }
}

package gg.foreversixty.app.talents

import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.home.HomeCharacter
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

/** The character the Talents screen is about: Home's current character (spec F6 §3). */
data class TalentCharacter(
    val classSlug: String,
    /** The primary tree's name as the API gives it (`Fury`); null when unknown. */
    val specName: String?,
    val level: Int?,
    val faction: String?,
    val race: String?,
    /** The export's per-tree rank strings; null without an export. */
    val trees: List<String>?,
    val dataBuild: String?,
) {
    companion object {
        fun from(character: HomeCharacter) = TalentCharacter(
            classSlug = character.className.orEmpty(),
            specName = character.specName?.takeIf { it.isNotEmpty() },
            level = character.level,
            faction = character.faction,
            race = character.race,
            // An export whose points are all 0 holds no ranks (the site's `exportCodeFromAccount`).
            trees = character.talents?.takeIf { talents -> talents.points.any { it > 0 } }?.trees,
            dataBuild = character.dataBuild,
        )
    }
}

/** The screen's states (spec F6 §6). */
sealed interface TalentsPhase {
    data object Loading : TalentsPhase
    /** What to draw; [offlineSinceMillis] is when the oldest saved copy shown was fetched. */
    data class Loaded(val presenter: TalentsPresenter, val offlineSinceMillis: Long?) : TalentsPhase
    /** The talent file is neither saved nor fetchable. */
    data object Failed : TalentsPhase
}

/** Loads the files behind the Talents screen and turns them into a presenter. iOS: TalentsModel.swift. */
class TalentsModel(
    private val source: TalentsSource,
    private val character: TalentCharacter,
    private val catalog: SpecCatalog,
) {
    private val mutable = MutableStateFlow<TalentsPhase>(TalentsPhase.Loading)
    val phase: StateFlow<TalentsPhase> = mutable.asStateFlow()
    var build: String = ""
        private set

    /** The first load, and Try again: a refresh never puts skeletons over content. */
    suspend fun load() {
        if (mutable.value == TalentsPhase.Failed) mutable.value = TalentsPhase.Loading
        build = source.activeBuild()
        val file = when (val result = source.talents(build, character.classSlug)) {
            is ApiResult.Success -> result.data
            is ApiResult.Failure -> {
                mutable.value = TalentsPhase.Failed
                return
            }
        }
        val guide = guideBuild(build)
        val export = export(file.value)
        val saved = listOfNotNull(file.savedAtIfCached, guide.savedAt, export.savedAt)
        val presenter = TalentsPresenter(file.value, guide.value, export.value, character.level, character.specName.orEmpty())
        mutable.value = TalentsPhase.Loaded(presenter, saved.minOrNull())
    }

    /** The guide build of the character's band (spec F6 §4.2). */
    private suspend fun guideBuild(build: String): Saved<TalentGuideBuild> {
        val spec = character.specName?.let { catalog.spec(character.classSlug, it) } ?: return Saved(TalentGuideBuild.NoSpec)
        val band = TalentGuide.bandForLevel(character.level) ?: return Saved(TalentGuideBuild.None)
        return when (val result = source.bis(build, spec.spec)) {
            is ApiResult.Success -> {
                val entry = TalentGuide.entry(result.data.value, band, character.faction)
                Saved(entry?.let { TalentGuideBuild.Build(it.talents, band) } ?: TalentGuideBuild.None, result.data.savedAtIfCached)
            }
            is ApiResult.Failure -> Saved(if (result.error.isMissingFile) TalentGuideBuild.None else TalentGuideBuild.Failed)
        }
    }

    /** The export, read on the active trees only when its own build lays them out the same. */
    private suspend fun export(active: TalentsFile): Saved<TalentExport> {
        val trees = character.trees ?: return Saved(TalentExport.None)
        val ranks = TalentExport.Ranks(trees.map(TalentRules::ranksOfTree))
        val dataBuild = character.dataBuild
        if (dataBuild == null || dataBuild == active.build) return Saved(ranks)
        val older = (source.talents(dataBuild, character.classSlug) as? ApiResult.Success)?.data ?: return Saved(TalentExport.Stale)
        return Saved(if (TalentGuide.sameLayout(older.value, active)) ranks else TalentExport.Stale, older.savedAtIfCached)
    }

    /** A value and, when it came from the offline copy, when that copy was fetched. */
    private data class Saved<T>(val value: T, val savedAt: Long? = null)
}

/** When the saved copy was fetched, or null for a fresh answer. */
val <T> CachedFile<T>.savedAtIfCached: Long? get() = if (isFromCache) savedAtMillis else null

/** A file the CDN does not have (404): an empty state, not an error. */
val ApiError.isMissingFile: Boolean
    get() = (this is ApiError.Status && status == HTTP_NOT_FOUND) || (this is ApiError.Server && status == HTTP_NOT_FOUND)

private const val HTTP_NOT_FOUND = 404

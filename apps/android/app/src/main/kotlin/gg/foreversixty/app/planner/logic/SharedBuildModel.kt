package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.RaceInfo
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.talents.TalentRules
import gg.foreversixty.app.talents.isMissingFile
import java.time.Instant
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

/** A shared build as the viewer draws it: the record, its own trees, and the names. */
data class SharedBuildView(
    val build: Build,
    /** The build's own tree version's talent file, never the active build's (spec F13.5). */
    val file: TalentsFile,
    val classSlug: String,
    val className: String,
    val raceName: String,
    /** When the offline copy shown was fetched, else null. */
    val offlineSinceMillis: Long? = null,
    /** The build against the guide build of the band it reaches (`Guide build`), when the spec's file has one. */
    val band: BandCompare? = null,
    val raceSlug: String = "",
) {
    /** A shared build has no character: its level is the one its points reach. */
    val level: Int get() = (build.pointOrder.size + TalentRules.BASE_LEVEL).coerceIn(PlannerEditor.MIN_LEVEL, PlannerEditor.MAX_LEVEL)

    val index: TalentIndex get() = TalentIndex(file)
    val totals: List<Int> get() = PlannerRules.pointsPerTree(index, build.pointOrder)
    val ranks: Map<Int, Int> get() = PlannerRules.ranks(build.pointOrder)
    val savedAt: Instant? get() = runCatching { Instant.parse(build.createdAt) }.getOrNull()
}

sealed interface SharedBuildPhase {
    data object Loading : SharedBuildPhase
    data class Loaded(val view: SharedBuildView) : SharedBuildPhase
    /** An unknown id: `This build does not exist.` */
    data object NotFound : SharedBuildPhase
    /** Offline with nothing saved. */
    data object Offline : SharedBuildPhase
    data object Failed : SharedBuildPhase
}

/** The `/b/{id}` viewer's loading (spec F13.5). iOS: SharedBuildModel.swift. */
class SharedBuildModel(private val source: PlannerSource, val id: String, private val catalog: SpecCatalog = SpecCatalog(emptyList())) {
    private val mutable = MutableStateFlow<SharedBuildPhase>(SharedBuildPhase.Loading)
    val phase: StateFlow<SharedBuildPhase> = mutable.asStateFlow()

    suspend fun load() {
        mutable.value = SharedBuildPhase.Loading
        mutable.value = when (val answer = source.build(id)) {
            is ApiResult.Failure -> failure(answer.error)
            is ApiResult.Success -> {
                val record = answer.data
                view(record.value, record.savedAtMillis.takeIf { record.isFromCache })
            }
        }
    }

    private suspend fun view(build: Build, offlineSince: Long?): SharedBuildPhase {
        val classes = (source.classes(build.treeVersion) as? ApiResult.Success)?.data?.value.orEmpty()
        val cls: ClassInfo = classes.firstOrNull { it.id == build.classId } ?: return SharedBuildPhase.Failed
        val file = when (val talents = source.talents(build.treeVersion, cls.slug)) {
            is ApiResult.Success -> talents.data.value
            // The record exists, so a missing tree file is a failure, never `does not exist`.
            is ApiResult.Failure -> return if (talents.error is ApiError.Network) SharedBuildPhase.Offline else SharedBuildPhase.Failed
        }
        val races: List<RaceInfo> = (source.races(build.treeVersion) as? ApiResult.Success)?.data?.value.orEmpty()
        val race = races.firstOrNull { it.id == build.raceId }
        val view = SharedBuildView(build, file, cls.slug, cls.name, race?.name.orEmpty(), offlineSince, raceSlug = race?.slug.orEmpty())
        return SharedBuildPhase.Loaded(view.copy(band = band(view, race?.faction)))
    }

    /** The guide build of the band the shared build reaches, read from the active build's BiS file. */
    private suspend fun band(view: SharedBuildView, faction: String?): BandCompare? {
        val index = view.index
        val spec = PlannerBand.specFor(catalog, index, view.build.pointOrder) ?: return null
        val file = (source.bis(source.activeBuild(), spec.spec) as? ApiResult.Success)?.data?.value ?: return null
        val build = PlannerBand.build(spec, file, PlannerBand.bandFor(view.build.pointOrder.size, null), faction) ?: return null
        return PlannerBand.compare(index, view.build.pointOrder, build, characterPoints = null)
    }

    private fun failure(error: ApiError): SharedBuildPhase = when {
        error is ApiError.Network -> SharedBuildPhase.Offline
        error.isMissingFile -> SharedBuildPhase.NotFound
        else -> SharedBuildPhase.Failed
    }
}

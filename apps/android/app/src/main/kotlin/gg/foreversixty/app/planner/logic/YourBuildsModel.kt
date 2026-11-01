package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.RaceInfo
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.talents.isMissingFile
import java.time.Instant
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow

/** One row of Your builds: the record and the names and totals it is shown with. */
data class BuildRow(
    val build: Build,
    val classSlug: String?,
    val className: String,
    val raceName: String,
    /** Points per tree, from the build's own trees; null when its tree file is not to hand. */
    val totals: List<Int>?,
) {
    val savedAt: Instant? get() = runCatching { Instant.parse(build.createdAt) }.getOrNull()
}

/** A row taken off the list, held for the snackbar's Undo (spec F13.6). */
data class PendingRemoval(val row: BuildRow, val position: Int)

/** Which write failed, for the line under the list. */
enum class BuildsProblem { RemoveFailed, UndoFailed }

sealed interface YourBuildsPhase {
    data object Loading : YourBuildsPhase
    data object Failed : YourBuildsPhase
    data class Loaded(
        val rows: List<BuildRow>,
        val pending: PendingRemoval? = null,
        val problem: BuildsProblem? = null,
    ) : YourBuildsPhase
}

/**
 * Your builds (spec F13.6): the list, the swipe that takes a row off it (`DELETE`; the link keeps
 * working) and the snackbar's Undo, which saves the same build again to put it back. iOS: YourBuildsModel.swift.
 */
class YourBuildsModel(private val source: PlannerSource) {
    private val mutable = MutableStateFlow<YourBuildsPhase>(YourBuildsPhase.Loading)
    val phase: StateFlow<YourBuildsPhase> = mutable.asStateFlow()

    private val loaded: YourBuildsPhase.Loaded? get() = mutable.value as? YourBuildsPhase.Loaded

    suspend fun load() {
        if (loaded == null) mutable.value = YourBuildsPhase.Loading
        val list = (source.myBuilds() as? ApiResult.Success)?.data
        if (list == null) {
            if (loaded == null) mutable.value = YourBuildsPhase.Failed
            return
        }
        val build = source.activeBuild()
        val classes = (source.classes(build) as? ApiResult.Success)?.data?.value.orEmpty()
        val races = (source.races(build) as? ApiResult.Success)?.data?.value.orEmpty()
        mutable.value = YourBuildsPhase.Loaded(list.rows.map { row(it, classes, races) })
    }

    private suspend fun row(build: Build, classes: List<ClassInfo>, races: List<RaceInfo>): BuildRow {
        val cls = classes.firstOrNull { it.id == build.classId }
        val file = cls?.let { (source.talents(build.treeVersion, it.slug) as? ApiResult.Success)?.data?.value }
        return BuildRow(
            build = build,
            classSlug = cls?.slug,
            className = cls?.name.orEmpty(),
            raceName = races.firstOrNull { it.id == build.raceId }?.name.orEmpty(),
            totals = file?.let { PlannerRules.pointsPerTree(TalentIndex(it), build.pointOrder) },
        )
    }

    /** The swipe: the row goes at once, then the DELETE; a failure puts it back and says so. */
    suspend fun remove(id: String) {
        val current = loaded ?: return
        val position = current.rows.indexOfFirst { it.build.id == id }
        if (position < 0) return
        val row = current.rows[position]
        mutable.value = current.copy(rows = current.rows - row, pending = PendingRemoval(row, position), problem = null)
        val answer = source.remove(id)
        // A 404 means it is already off the list: the same end state.
        if (answer is ApiResult.Failure && !answer.error.isMissingFile) restore(row, position, BuildsProblem.RemoveFailed)
    }

    /** The snackbar's Undo: the same build saved again, which puts it back on the list. */
    suspend fun undo() {
        val pending = loaded?.pending ?: return
        expire()
        when (val answer = source.save(PlannerShare.again(pending.row.build))) {
            is ApiResult.Success -> restore(pending.row.copy(build = pending.row.build.copy(id = answer.data.id)), pending.position, null)
            is ApiResult.Failure -> loaded?.let { mutable.value = it.copy(problem = BuildsProblem.UndoFailed) }
        }
    }

    /** The snackbar's 5 s ran out: the removal stands. */
    fun expire() {
        loaded?.let { mutable.value = it.copy(pending = null) }
    }

    private fun restore(row: BuildRow, position: Int, problem: BuildsProblem?) {
        val current = loaded ?: return
        val rows = current.rows.toMutableList().apply { add(position.coerceIn(0, size), row) }
        mutable.value = current.copy(rows = rows.toList(), pending = null, problem = problem)
    }
}

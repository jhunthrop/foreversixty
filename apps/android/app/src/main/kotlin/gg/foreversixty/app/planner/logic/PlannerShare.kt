package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.BuildInput
import gg.foreversixty.app.model.BuildSaved
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult

/** What `Share anyway` came to (spec F13.3). */
sealed interface ShareOutcome {
    data class Shared(val saved: BuildSaved) : ShareOutcome
    /** A 400: the first field's sentence, shown under the editor. */
    data class Field(val sentence: String) : ShareOutcome
    data object RateLimited : ShareOutcome
    data object Offline : ShareOutcome
    data object Failed : ShareOutcome
}

/** The share request and its answer (spec F13.3, the web's `POST /v1/builds`). iOS: PlannerShare.swift. */
object PlannerShare {
    /** The title field's limit, with its counter. */
    const val TITLE_MAX = 60
    private const val HTTP_BAD_REQUEST = 400
    private const val HTTP_TOO_MANY = 429

    /** The body: the class, race, active tree version and order, an empty gear map (the planner owns talents only). */
    fun input(classId: Int, raceId: Int, treeVersion: String, order: List<Int>, title: String): BuildInput = BuildInput(
        classId = classId,
        raceId = raceId,
        treeVersion = treeVersion,
        pointOrder = order,
        gear = emptyMap(),
        title = cleanTitle(title),
    )

    /** The same build sent again (Your builds' Undo): the server answers with the same record. */
    fun again(build: Build): BuildInput = BuildInput(
        classId = build.classId,
        raceId = build.raceId,
        treeVersion = build.treeVersion,
        pointOrder = build.pointOrder,
        gear = build.gear,
        title = build.title,
    )

    /** Trimmed, at most [TITLE_MAX] characters, and absent when empty. */
    fun cleanTitle(title: String): String? = title.trim().take(TITLE_MAX).ifEmpty { null }

    fun outcome(result: ApiResult<BuildSaved>): ShareOutcome = when (result) {
        is ApiResult.Success -> ShareOutcome.Shared(result.data)
        is ApiResult.Failure -> failure(result.error)
    }

    private fun failure(error: ApiError): ShareOutcome = when {
        error is ApiError.Network -> ShareOutcome.Offline
        error.status == HTTP_TOO_MANY -> ShareOutcome.RateLimited
        error is ApiError.Server && error.status == HTTP_BAD_REQUEST ->
            // The earliest field: `point_order[2]` before `point_order[10]`.
            error.fields.entries.minWithOrNull(compareBy({ it.key.length }, { it.key }))?.value?.let { ShareOutcome.Field(it) } ?: ShareOutcome.Field(error.message)
        else -> ShareOutcome.Failed
    }

    private val ApiError.status: Int?
        get() = when (this) {
            is ApiError.Server -> status
            is ApiError.Status -> status
            else -> null
        }
}

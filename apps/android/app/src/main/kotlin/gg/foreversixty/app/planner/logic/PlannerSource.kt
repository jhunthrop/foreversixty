package gg.foreversixty.app.planner.logic

import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.BuildInput
import gg.foreversixty.app.model.BuildList
import gg.foreversixty.app.model.BuildSaved
import gg.foreversixty.app.model.ClassInfo
import gg.foreversixty.app.model.RaceClassCombo
import gg.foreversixty.app.model.RaceInfo
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.talents.TalentsSource
import kotlinx.serialization.SerializationException

/** Everything the planner reads and writes (spec F13 Data). iOS: PlannerSource.swift. */
interface PlannerSource {
    suspend fun activeBuild(): String
    suspend fun talents(build: String, classSlug: String): ApiResult<CachedFile<TalentsFile>>
    suspend fun classes(build: String): ApiResult<CachedFile<List<ClassInfo>>>
    suspend fun races(build: String): ApiResult<CachedFile<List<RaceInfo>>>
    suspend fun combos(build: String): ApiResult<CachedFile<List<RaceClassCombo>>>

    /** `POST /v1/builds`: never queued offline. */
    suspend fun save(input: BuildInput): ApiResult<BuildSaved>

    /** `GET /v1/builds/{id}`, with the copy kept from the last good answer for offline. */
    suspend fun build(id: String): ApiResult<CachedFile<Build>>

    /** `GET /v1/builds?mine=1`. */
    suspend fun myBuilds(): ApiResult<BuildList>

    /** `DELETE /v1/builds/{id}`: off the list; the record and its link stay. */
    suspend fun remove(id: String): ApiResult<Unit>
}

/**
 * The live source: the CDN through the app's one static data store (lane F15), the API through
 * [api]; a shared build's last good answer is kept in the same store for offline.
 */
class LivePlannerSource(
    private val files: StaticDataStore,
    private val talentsSource: TalentsSource,
    private val api: ApiClient,
) : PlannerSource {
    override suspend fun activeBuild(): String = talentsSource.activeBuild()
    override suspend fun talents(build: String, classSlug: String) = talentsSource.talents(build, classSlug)
    override suspend fun classes(build: String) = files.load(StaticEndpoints.getClasses(build))
    override suspend fun races(build: String) = files.load(StaticEndpoints.getRaces(build))
    override suspend fun combos(build: String) = files.load(StaticEndpoints.getCombos(build))
    override suspend fun save(input: BuildInput) = api.send(ApiEndpoints.saveBuild(input))
    override suspend fun myBuilds() = api.send(ApiEndpoints.listMyBuilds())
    override suspend fun remove(id: String) = api.send(ApiEndpoints.removeBuild(id))

    override suspend fun build(id: String): ApiResult<CachedFile<Build>> {
        val path = cachePath(id)
        return when (val answer = api.send(ApiEndpoints.getBuild(id))) {
            is ApiResult.Success -> {
                files.saveBody(path, ApiJson.encodeToString(Build.serializer(), answer.data))
                ApiResult.Success(CachedFile(answer.data, files.savedBody(path)?.savedAtMillis ?: 0L, isFromCache = false), answer.requestId)
            }
            is ApiResult.Failure -> saved(path)?.takeIf { answer.error is ApiError.Network } ?: answer
        }
    }

    private fun saved(path: String): ApiResult<CachedFile<Build>>? {
        val entry = files.savedBody(path) ?: return null
        return try {
            ApiResult.Success(CachedFile(ApiJson.decodeFromString(Build.serializer(), entry.body), entry.savedAtMillis, isFromCache = true), "")
        } catch (e: SerializationException) {
            null
        }
    }

    private companion object {
        fun cachePath(id: String) = "/v1/builds/${id.filter { it.isLetterOrDigit() }}.json"
    }
}

package gg.foreversixty.app.planner

import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.model.Build
import gg.foreversixty.app.model.BuildInput
import gg.foreversixty.app.model.BuildList
import gg.foreversixty.app.model.BuildSaved
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.planner.logic.PlannerSource
import java.io.IOException

/** The planner's source over the repo's real files, with scripted API answers and a log of writes. */
class FakePlannerSource(
    private val files: Map<String, TalentsFile> = mapOf("shaman" to PlannerTestData.shaman(), "warrior" to PlannerTestData.warrior()),
    var saveAnswer: ApiResult<BuildSaved> = ApiResult.Success(BuildSaved("k7q2m9xa", "https://foreversixty.gg/b/k7q2m9xa"), "r"),
    var buildAnswer: ApiResult<CachedFile<Build>> = ApiResult.Success(CachedFile(PlannerTestData.sharedBuild(), 0, false), "r"),
    var listAnswer: ApiResult<BuildList> = ApiResult.Success(BuildList(listOf(PlannerTestData.sharedBuild()), 1, 1, 100), "r"),
    var removeAnswer: ApiResult<Unit> = ApiResult.Success(Unit, "r"),
) : PlannerSource {
    val saved = mutableListOf<BuildInput>()
    val removed = mutableListOf<String>()

    override suspend fun activeBuild() = BisTestData.BUILD
    override suspend fun talents(build: String, classSlug: String) =
        files[classSlug]?.let { ApiResult.Success(CachedFile(it, 0, false), "r") } ?: ApiResult.Failure(ApiError.Status(404))
    override suspend fun classes(build: String) = ApiResult.Success(CachedFile(PlannerTestData.classes(), 0, false), "r")
    override suspend fun races(build: String) = ApiResult.Success(CachedFile(PlannerTestData.races(), 0, false), "r")
    override suspend fun combos(build: String) = ApiResult.Success(CachedFile(PlannerTestData.combos(), 0, false), "r")
    override suspend fun save(input: BuildInput): ApiResult<BuildSaved> = saveAnswer.also { saved += input }
    override suspend fun build(id: String) = buildAnswer
    override suspend fun myBuilds() = listAnswer
    override suspend fun remove(id: String): ApiResult<Unit> = removeAnswer.also { removed += id }

    companion object {
        val offline = ApiResult.Failure(ApiError.Network(IOException("offline")))
    }
}

/** The app's recording facade (W3 §5.7): the planner's haptics go through the one Haptics. */
typealias RecordingHaptics = gg.foreversixty.app.raids.RecordingHaptics

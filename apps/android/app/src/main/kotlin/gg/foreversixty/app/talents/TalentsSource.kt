package gg.foreversixty.app.talents

import android.content.Context
import gg.foreversixty.app.bis.BisDataSource
import gg.foreversixty.app.bis.SpecCatalog
import gg.foreversixty.app.bis.bundledSpecCatalog
import gg.foreversixty.app.bis.createBisDataSource
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.TalentsFile
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.network.StaticEndpoints
import gg.foreversixty.app.planner.PlannerFeature

/**
 * What the Talents screen reads: the active build, the class's talent file and the spec's BiS
 * file (its bands carry the guide builds), each with its offline copy. iOS: TalentsSource.swift.
 */
interface TalentsSource {
    suspend fun activeBuild(): String
    suspend fun talents(build: String, classSlug: String): ApiResult<CachedFile<TalentsFile>>
    suspend fun bis(build: String, spec: String): ApiResult<CachedFile<BisFile>>
}

/** The live source: the CDN through the offline cache, the build as the BiS tab reads it. */
class LiveTalentsSource(
    private val files: StaticDataStore,
    private val bisSource: BisDataSource,
) : TalentsSource {
    override suspend fun activeBuild(): String = bisSource.activeBuild()
    override suspend fun talents(build: String, classSlug: String) = files.load(StaticEndpoints.getTalents(build, classSlug))
    override suspend fun bis(build: String, spec: String) = bisSource.bis(build, spec)
}

/**
 * What the Talents screen needs from the app graph (lane F6): its files and the spec catalogue; and
 * lane F13's planner, whose `Plan a build` and `Your builds` it links to (null while `planner` is off).
 */
class TalentsFeature(val source: TalentsSource, val catalog: SpecCatalog, val planner: PlannerFeature? = null) {
    companion object {
        fun create(context: Context, statics: StaticDataStore, appConfig: AppConfigRepository, planner: PlannerFeature? = null) = TalentsFeature(
            source = LiveTalentsSource(statics, createBisDataSource(context, statics, appConfig)),
            catalog = bundledSpecCatalog(context),
            planner = planner,
        )
    }
}

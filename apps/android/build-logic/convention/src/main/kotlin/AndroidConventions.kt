import org.gradle.api.Project
import org.gradle.api.artifacts.VersionCatalog
import org.gradle.api.artifacts.VersionCatalogsExtension
import org.gradle.kotlin.dsl.getByType

/** The one version catalog, read by every convention plugin so SDK levels live in one file. */
internal val Project.libs: VersionCatalog
    get() = extensions.getByType<VersionCatalogsExtension>().named("libs")

internal fun VersionCatalog.intVersion(alias: String): Int =
    findVersion(alias).orElseThrow { IllegalStateException("libs.versions.toml has no version '$alias'") }
        .requiredVersion
        .toInt()

/** The Android SDK levels and JVM toolchain every module builds against. */
internal data class AndroidLevels(
    val compileSdk: Int,
    val targetSdk: Int,
    val minSdk: Int,
    val jvmToolchain: Int,
)

internal val Project.androidLevels: AndroidLevels
    get() = AndroidLevels(
        compileSdk = libs.intVersion("compileSdk"),
        targetSdk = libs.intVersion("targetSdk"),
        minSdk = libs.intVersion("minSdk"),
        jvmToolchain = libs.intVersion("jvmToolchain"),
    )

/** The package every module namespaces under: gg.foreversixty.app.<suffix>. */
internal const val APP_NAMESPACE = "gg.foreversixty.app"

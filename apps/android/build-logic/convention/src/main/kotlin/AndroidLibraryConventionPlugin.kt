import com.android.build.api.dsl.CommonExtension
import com.android.build.api.dsl.LibraryExtension
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.kotlin.dsl.configure
import org.jetbrains.kotlin.gradle.dsl.KotlinAndroidProjectExtension

/** A :core library with a namespace derived from its name (:core:network -> gg.foreversixty.app.network). */
class AndroidLibraryConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        pluginManager.apply("com.android.library")
        extensions.configure<LibraryExtension> {
            namespace = "$APP_NAMESPACE.${name.replace('-', '.')}"
        }
        configureAndroid(extensions.getByType(LibraryExtension::class.java))
    }
}

/**
 * What every Android module shares, app or library: the SDK levels, the JDK toolchain, strict
 * lint and Kotlin warnings, and JVM tests that see Android resources so Robolectric can run them.
 */
internal fun Project.configureAndroid(android: CommonExtension) {
    val levels = androidLevels
    android.compileSdk = levels.compileSdk
    android.defaultConfig.minSdk = levels.minSdk
    android.testOptions.unitTests.isIncludeAndroidResources = true
    android.lint.apply {
        warningsAsErrors = true
        abortOnError = true
        // Dependency freshness is a scheduled bump's job, not a reason to fail a build; and a
        // library whose bundled lint jar predates this lint (lifecycle today) is its bug, not ours.
        disable += setOf(
            "GradleDependency",
            "NewerVersionAvailable",
            "AndroidGradlePluginVersion",
            "ObsoleteLintCustomCheck",
        )
    }
    extensions.configure<KotlinAndroidProjectExtension> {
        jvmToolchain(levels.jvmToolchain)
        compilerOptions { allWarningsAsErrors.set(true) }
    }
}

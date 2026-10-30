import com.android.build.api.dsl.ApplicationExtension
import org.gradle.api.GradleException
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.kotlin.dsl.configure

/** The two values apps/VERSION holds: the semver players see and the build number. */
internal data class AppVersion(val name: String, val build: Int) {
    companion object {
        private val SEMVER = Regex("""^\d+\.\d+\.\d+$""")

        /** Parses the file's `KEY = value` lines (xcconfig syntax, `//` comments); fails loudly. */
        fun parse(text: String): AppVersion {
            val values = text.lines()
                .map { it.trim() }
                .filter { it.isNotEmpty() && !it.startsWith("//") }
                .associate { line ->
                    val parts = line.split("=", limit = 2)
                    if (parts.size != 2) throw GradleException("apps/VERSION: not KEY = value: '$line'")
                    parts[0].trim() to parts[1].trim()
                }
            val name = values["APP_VERSION"] ?: throw GradleException("apps/VERSION has no APP_VERSION")
            if (!SEMVER.matches(name)) throw GradleException("apps/VERSION: APP_VERSION '$name' is not x.y.z")
            val build = values["APP_BUILD"]?.toIntOrNull()?.takeIf { it > 0 }
                ?: throw GradleException("apps/VERSION: APP_BUILD must be a positive integer")
            return AppVersion(name, build)
        }
    }
}

/** Sets versionName and versionCode from apps/VERSION, the one version file both apps read. */
class AppVersionConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        val file = rootProject.layout.projectDirectory.file("../VERSION")
        if (!file.asFile.isFile) throw GradleException("apps/VERSION is missing at ${file.asFile.path}")
        // Read through a provider so the configuration cache sees the file as an input.
        val version = AppVersion.parse(providers.fileContents(file).asText.get())
        extensions.configure<ApplicationExtension> {
            defaultConfig {
                versionName = version.name
                versionCode = version.build
            }
        }
    }
}

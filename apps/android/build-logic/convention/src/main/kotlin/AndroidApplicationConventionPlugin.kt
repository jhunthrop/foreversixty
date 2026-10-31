import com.android.build.api.dsl.ApplicationExtension
import org.gradle.api.Plugin
import org.gradle.api.Project
import org.gradle.kotlin.dsl.configure

/** The :app module: the shared Android setup plus the application id and target SDK. */
class AndroidApplicationConventionPlugin : Plugin<Project> {
    override fun apply(target: Project) = with(target) {
        pluginManager.apply("com.android.application")
        val android = extensions.getByType(ApplicationExtension::class.java)
        configureAndroid(android)
        extensions.configure<ApplicationExtension> {
            namespace = APP_NAMESPACE
            defaultConfig {
                applicationId = APP_NAMESPACE
                targetSdk = androidLevels.targetSdk
            }
        }
    }
}

#if DEBUG
import Foundation

/// Debug-only: `-FSRefreshHold 2` holds the drawer's Battle.net refresh that many seconds, so the
/// review walk captures its `Refreshing…` frame (the fixture answers at once). Never in a release.
enum MenuDebugLaunch {
    static func holdRefresh() async {
        let arguments = ProcessInfo.processInfo.arguments
        guard let index = arguments.firstIndex(of: "-FSRefreshHold"), arguments.indices.contains(index + 1),
              let seconds = Double(arguments[index + 1])
        else { return }
        try? await Task.sleep(for: .seconds(seconds))
    }
}
#endif

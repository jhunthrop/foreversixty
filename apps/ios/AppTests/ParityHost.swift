import SwiftUI
import UIKit

/// Renders a view for the parity helpers: a window of the device's size on the test host's scene,
/// laid out and given a moment for preferences and accessibility to settle.
@MainActor
enum ParityHost {
    /// Runs `body` with `view` on screen at `size`, then takes the window down.
    static func render<Content: View, Result>(
        _ view: Content,
        size: CGSize,
        _ body: (UIHostingController<Content>) -> Result
    ) -> Result {
        guard let scene = UIApplication.shared.connectedScenes.compactMap({ $0 as? UIWindowScene }).first else {
            fatalError("ParityHost: the test host has no window scene")
        }
        let host = UIHostingController(rootView: view)
        let window = UIWindow(windowScene: scene)
        window.frame = CGRect(origin: .zero, size: size)
        window.rootViewController = host
        window.makeKeyAndVisible()
        host.view.frame = window.bounds
        host.view.layoutIfNeeded()
        // Preferences arrive after the first layout pass; two run-loop turns settle them.
        for _ in 0 ..< 2 {
            RunLoop.main.run(until: Date().addingTimeInterval(0.05))
        }
        defer { window.isHidden = true }
        return body(host)
    }
}

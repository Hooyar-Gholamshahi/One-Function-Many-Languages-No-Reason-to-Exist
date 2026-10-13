#include <windows.h>
#include <iostream>

using HelloWorldFn = void (*)();

int main() {
    HMODULE library = LoadLibraryW(
        L"target\\debug\\rust_native_bridge.dll"
    );

    if (library == nullptr) {
        std::cerr << "LoadLibraryW failed: "
                  << GetLastError() << '\n';
        return 1;
    }

    auto hello_world = reinterpret_cast<HelloWorldFn>(
        GetProcAddress(library, "hello_world")
    );

    if (hello_world == nullptr) {
        std::cerr << "GetProcAddress failed: "
                  << GetLastError() << '\n';
        FreeLibrary(library);
        return 2;
    }

    hello_world();

    FreeLibrary(library);
    return 0;
}

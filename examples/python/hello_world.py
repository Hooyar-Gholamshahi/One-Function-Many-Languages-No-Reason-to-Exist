import ctypes
from pathlib import Path

project_root = Path(__file__).resolve().parents[2]
dll_path = project_root / "target" / "debug" / "rust_native_bridge.dll"

library = ctypes.CDLL(str(dll_path))

hello_world = library.hello_world
hello_world.argtypes = []
hello_world.restype = None

hello_world()

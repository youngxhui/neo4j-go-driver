"""
Executed in Go driver container.
Responsible for building driver and test backend.
"""

import sys
from pathlib import Path
import os
import subprocess


def run(args, env=None):
    subprocess.run(
        args, universal_newlines=True, check=True, env=env,
        stdout=sys.stdout, stderr=sys.stderr,
    )


if __name__ == "__main__":
    defaultEnv = os.environ.copy()
    defaultEnv["GOFLAGS"] = "-buildvcs=false"

    print("Building for current target", flush=True)
    run(["go", "build", "-tags", "internal_testkit", "-v", "./..."], env=defaultEnv)

    # Compile for 32 bits ARM to make sure it builds
    print("Building for 32 bits", flush=True)
    arm32Env = defaultEnv.copy()
    arm32Env["GOOS"] = "linux"
    arm32Env["GOARCH"] = "arm"
    arm32Env["GOARM"] = "7"
    run(["go", "build", "./neo4j/..."], env=arm32Env)

    print("Vet sources", flush=True)
    run(["go", "vet", "-tags", "internal_testkit", "./..."], env=defaultEnv)

    print("Install staticcheck", flush=True)
    run(["go", "install", "honnef.co/go/tools/cmd/staticcheck@v0.3.3"], env=defaultEnv)

    print("Run staticcheck", flush=True)
    gopath = Path(
        subprocess.check_output(["go", "env", "GOPATH"]).decode("utf-8").strip()
    )
    run([str(gopath / "bin" / "staticcheck"), "-tags", "internal_testkit", "./..."], env=defaultEnv)

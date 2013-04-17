
package main

import (
    "labix.org/v1/pipe"
    "log"
)

func main() {
    output, err := pipe.CombinedOutput(
        pipe.Line(
            pipe.Exec("ls", "-la", "/usr/bin"),
            pipe.System("grep pp"),
            pipe.WriteFile("grep.txt", 0644),
        ),
    )

    if err != nil {
        log.Fatal(err)
    }

    println(string(output))
}

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

go_binary(
    name = "hello",
    embed = [":bazel-bug-go-tool_lib"],
    visibility = ["//visibility:public"],
)

# gazelle:prefix github.com/jschaf/bazel-bug-go-tool
gazelle(name = "gazelle")

go_library(
    name = "bazel-bug-go-tool_lib",
    srcs = ["hello.go"],
    importpath = "github.com/jschaf/bazel-bug-go-tool",
    visibility = ["//visibility:private"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test", "go_tool_library")

go_tool_library(
    name = "lint",
    srcs = ["import_unsafe.go"],
    importpath = "github.com/jschaf/bazel-bug-go-tool/lint",
    visibility = ["//visibility:public"],
    deps = ["@org_golang_x_tools//go/analysis:go_tool_library"],
)

go_test(
    name = "lint_test",
    srcs = ["import_unsafe_test.go"],
    data = ["@go_sdk//:files"],
    embed = [":lint"],
    deps = [
        "@org_golang_x_tools//go/analysis:go_default_library",
        "@org_golang_x_tools//go/analysis/analysistest:go_default_library",
    ],
)

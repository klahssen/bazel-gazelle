getgazelle:
	go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle

gazelle:
	bazel run //:gazelle

buildgreeter: gazelle
	bazel build --experimental_convenience_symlinks=clean //cmd/greeter:greeter

rungreeter: gazelle
	time bazel run --experimental_convenience_symlinks=clean //cmd/greeter:greeter

clean:
	time bazel clean
#!/bin/bash
set -euo pipefail

echo "Running Go benchmarks..."
bazel test --test_output=streamed \
  --test_arg=-test.bench=. \
  --test_arg=-test.benchmem \
  --test_tag_filters=benchmark \
  --cache_test_results=no \
  //...

echo
echo "Running C++ benchmarks..."

# Query all benchmark-tagged cc_binary targets
cpp_benches=$(bazel query 'attr(tags, "benchmark", kind(cc_binary, //...))')

for target in $cpp_benches; do
  echo
  echo ">>> Running $target"
  bazel run -c opt "$target"
done

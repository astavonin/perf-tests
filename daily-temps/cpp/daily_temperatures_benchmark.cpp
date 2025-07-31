#include "daily_temperatures.h"

#include <benchmark/benchmark.h>
#include <vector>

std::vector<int> generateInput( int size )
{
    std::vector<int> input( size );
    for( int i = 0; i < size; ++i ) {
        input[i] = 30 + ( i % 40 );
    }
    return input;
}

static void BM_DailyTemperatures( benchmark::State &state )
{
    const auto input = generateInput( state.range( 0 ) );
    for( auto _ : state ) {
        benchmark::DoNotOptimize( dailyTemperatures( input ) );
    }
}

static void BM_DailyTemperaturesOpt( benchmark::State &state )
{
    const auto input = generateInput( state.range( 0 ) );
    for( auto _ : state ) {
        benchmark::DoNotOptimize( dailyTemperaturesOpt( input ) );
    }
}

// Register benchmarks with multiple input sizes
BENCHMARK( BM_DailyTemperatures )->Arg( 1000 )->Arg( 10000 )->Arg( 100000 );
BENCHMARK( BM_DailyTemperaturesOpt )->Arg( 1000 )->Arg( 10000 )->Arg( 100000 );

BENCHMARK_MAIN();

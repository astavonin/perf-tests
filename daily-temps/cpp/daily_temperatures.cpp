#include "daily_temperatures.h"

#include <stack>
#include <vector>

std::vector<int> dailyTemperatures( const std::vector<int> &temperatures )
{
    std::vector<int> result( temperatures.size(), 0 );
    std::stack<int>  s;
    for( int i = 0; i < temperatures.size(); ++i ) {
        while( !s.empty() && temperatures[i] > temperatures[s.top()] ) {
            int prev = s.top();
            s.pop();
            result[prev] = i - prev;
        }
        s.push( i );
    }
    return result;
}

std::vector<int> dailyTemperaturesOpt( const std::vector<int> &temperatures )
{
    std::vector<int> res( temperatures.size(), 0 );
    std::vector<int> track;
    track.reserve( temperatures.size() );

    for( int i = 0; i < temperatures.size(); ++i ) {
        int currTemp = temperatures[i]; // <- preventing CPU-cache misses here
        while( !track.empty() && currTemp > temperatures[track.back()] ) {
            int prev = track.back();
            track.pop_back();
            res[prev] = i - prev;
        }
        track.push_back( i );
    }
    return res;
}

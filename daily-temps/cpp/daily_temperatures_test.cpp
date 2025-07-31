#include "daily_temperatures.h"

#include <gtest/gtest.h>

TEST( DailyTemperaturesTest, Basic )
{
    std::vector<int> input    = { 73, 74, 75, 71, 69, 72, 76, 73 };
    std::vector<int> expected = { 1, 1, 4, 2, 1, 1, 0, 0 };
    EXPECT_EQ( dailyTemperatures( input ), expected );
}

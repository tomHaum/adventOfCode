using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day01
    {
        private static int SET_SIZE = 3;
        internal static void Run()
        {
            var depths = input.day01.Split("\r\n").Select(x => int.Parse(x)).ToArray();
            int increadingCount = 0;
            for(int i = 1; i < depths.Length; i++)
            {
                if (depths[i] > depths[i - 1])
                    increadingCount++;
            }
            Console.WriteLine($"Day 01 - Part 1: {increadingCount} increasing scans");

            int setSize = 3;
            int windowA = 0;
            int windowB = 0;
            increadingCount = 0;

            windowA = calculateWindowSum(depths, 0);

            for (int i = 1; i < depths.Length - SET_SIZE + 1; i++)
            {
                windowB = getNextWindowSum(depths, i, windowA);
                if (windowB > windowA)
                    increadingCount++;
                windowA = windowB;
            }
            Console.WriteLine($"Day 01 - Part 2: {increadingCount} increasing windows");
        }

        private static int calculateWindowSum(int[] depths, int index)
        {
            var windowSum = 0;
            for(int i = 0; i < SET_SIZE; i++)
            {
                windowSum += depths[index + i];
            }
            return windowSum;
        }

        private static int getNextWindowSum(int[] depths, int index, int previousWindowSum)
        {
            return previousWindowSum - depths[index - 1] + depths[index + SET_SIZE-1];
        }
    }
}

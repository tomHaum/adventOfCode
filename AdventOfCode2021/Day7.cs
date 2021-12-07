using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day7
    {
        internal static void Run()
        {
            var min = int.MaxValue;
            var max = int.MinValue;
            var crabs = input.day7.Split(",").Select(x =>
            {
                var y = int.Parse(x);
                min = Math.Min(min, y);
                max = Math.Max(max, y);
                return y;
            }).ToArray();

            var sum = crabs.Sum(); ;
            var avg = crabs.Average();


            var minCost = int.MaxValue;
            for (int i =min; i <= max; i++)
            {
                minCost = Math.Min(minCost, crabs.Aggregate(0, (accum, x) => accum + Math.Abs(i - x)));
            }
            Console.WriteLine("Day 7 - Part 1: " + minCost);

            minCost = int.MaxValue;
            for (int i = min; i <= max; i++)
            {
                minCost = Math.Min(minCost, crabs.Aggregate(0, (accum, x) => accum + Triangle(Math.Abs(i - x))));
            }

            Console.WriteLine("Day 7 - Part 2: " + minCost);
        }
        
        private static Dictionary<int, int> dict = new Dictionary<int, int>();
        private static int Triangle(int x)
        {
            if (!dict.ContainsKey(x))
            {
                var accum = 0;
                for(int i = 1; i <= x; i++)
                {
                    accum += i;
                }
                dict[x] = accum;
            }

            return dict[x];
        }
    }
}

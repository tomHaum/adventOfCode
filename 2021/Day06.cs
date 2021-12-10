using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Diagnostics;
using BenchmarkDotNet.Attributes;
namespace AdventOfCode2021
{
    [RPlotExporter] 
    public class Day06
    {
        
        internal static void Run()
        {


            var fishes = new long[9];
            var fishesTemp = new long[9];
            var txt = input.day06.Split(',').Select(x => long.Parse(x));
            foreach (var x in txt)
            {
                fishes[x]++;
            }

            for (int day = 0; day < 80; day++)
            {
                long newFishes = fishes[0];
                long[] tmp = fishes;
                for (int i = 0; i < fishes.Length; i++)
                {
                    fishesTemp[i] = fishes[(i + 1) % fishes.Length];
                }
                fishes = fishesTemp;
                fishesTemp = tmp;
                fishes[6] += newFishes;
            }

            Console.WriteLine("Day 06 - Part 1: " + fishes.Aggregate(0L, (accum, x) => accum + x));
            Stopwatch sw = new();
            var part2 = new Day06().Part2();
            Console.WriteLine("Day 06 - Part 2: " + part2);

        }
        [Benchmark]
        public long Part2()
        {
            var fishes = new long[9];
            var txt = input.day06.Split(',').Select(x => long.Parse(x));
            foreach (var x in txt)
            {
                fishes[x]++;
            }
            for (int day = 0; day < 256; day++)
            {
                long newFishes = fishes[0];
                for (int i = 0; i < fishes.Length - 1; i++)
                {
                    fishes[i] = fishes[(i + 1)];
                }
                fishes[6] += newFishes;
                fishes[8] = newFishes;
                //Console.WriteLine(String.Join(' ', fishes));
            }
            return fishes.Aggregate(0L, (accum, x) => accum + x);
        }
    }
}

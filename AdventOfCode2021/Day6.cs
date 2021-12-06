using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day6
    {

        internal static void Run()
        {
            var fishes = new long[9];
            var fishesTemp = new long[9];
            var txt = input.day6.Split(',').Select(x => long.Parse(x));
            foreach (var x in txt)
            {
                fishes[x]++;
            }

            Console.WriteLine(string.Join(' ', Enumerable.Range(0,9)));

            Console.WriteLine(String.Join(' ', fishes));
            for(int day = 0; day < 80; day++)
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
                //Console.WriteLine(String.Join(' ', fishes));
            }

            Console.WriteLine(fishes.Aggregate(0L, (accum, x) => accum + x));
            fishes = new long[9];
            fishesTemp = new long[9];
            txt = input.day6.Split(',').Select(x => long.Parse(x));
            foreach (var x in txt)
            {
                fishes[x]++;
            }
            for (int day = 0; day < 256; day++)
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
                //Console.WriteLine(String.Join(' ', fishes));
            }
            Console.WriteLine(fishes.Aggregate(0L, (accum, x) => accum + x));
        }
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day3
    {
        private static int LINE_WIDTH = 12;
        //private static int LINE_WIDTH = 5;
        internal static void Run()
        {
            var lines = input.day3.Split("\r\n").ToArray();


            int[] accum = new int[LINE_WIDTH];

            foreach(var l in lines)
            {
                for(int i = 0; i < LINE_WIDTH; i++)
                {
                    if (l[i] == '1')
                        accum[i] += 1;
                }
            }

            Console.WriteLine(string.Join(' ', accum));
            var gamma = accum.Select(x => x > lines.Length/2).Aggregate(0, (acc, x) => (acc << 1 )+ (x ?  1 : 0));
            var mask = Enumerable.Range(0, LINE_WIDTH).Aggregate(0, (acc, x) => (acc << 1) + 1);
            var epsilon = gamma ^ mask;
            Console.WriteLine($"Day 3 - Part 1: gamma [{gamma}] | epsilon [{epsilon}] | product [{gamma * epsilon}]");
        }
    }
}

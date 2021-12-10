using System;
using System.Collections;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day03
    {
        private static int LINE_WIDTH = 12;
        //private static int LINE_WIDTH = 5;
        internal static void Run()
        {
            var lines = input.day03.Split("\r\n").ToArray();


            int[] accum = new int[LINE_WIDTH];

            foreach(var l in lines)
            {
                for(int i = 0; i < LINE_WIDTH; i++)
                {
                    if (l[i] == '1')
                        accum[i] += 1;
                }
            }

            Debug.WriteLine(string.Join(' ', accum));
            var gamma = accum.Select(x => x > lines.Length/2).Aggregate(0, (acc, x) => (acc << 1 )+ (x ?  1 : 0));
            var mask = Enumerable.Range(0, LINE_WIDTH).Aggregate(0, (acc, x) => (acc << 1) + 1);
            var epsilon = gamma ^ mask;

            Console.WriteLine($"Day 03 - Part 1: gamma [{gamma}] | epsilon [{epsilon}] | product [{gamma * epsilon}]");

            var oxygenGenRating = GetRating(lines, true);
            var co2ScrubberRating = GetRating(lines, false);

            Console.WriteLine($"Day 03 - Part 2: {nameof(oxygenGenRating)} [{oxygenGenRating}] | {nameof(co2ScrubberRating)} [{co2ScrubberRating}] | product [{oxygenGenRating * co2ScrubberRating}]");
        }
        private static int GetRating(string[] lines, bool lookForOneOnMoreOnes)
        {
            var acceptedLeft = lines.Length;
            var accepted = new BitArray(lines.Length, true);
            var ratingStr = string.Empty;
            var lookForOnMoreOnes = lookForOneOnMoreOnes ? '1' : '0';
            var lookForOnMoreZeros = lookForOneOnMoreOnes ? '0' : '1';

            for (int bitIndex = 0; bitIndex < LINE_WIDTH; bitIndex++)
            {
                var onesCount = 0;

                for (int i = 0; i < lines.Length; i++)
                {
                    if (!accepted[i])
                        continue;
                    onesCount += lines[i][bitIndex] == '1' ? 1 : 0;
                }

                var zerosCount = (acceptedLeft - onesCount);
                var lookFor = onesCount >= zerosCount ? lookForOnMoreOnes : lookForOnMoreZeros;
                Debug.WriteLine(onesCount + " - " +  lookFor);

                for (int i = 0; i < lines.Length; i++)
                {
                    if (!accepted[i])
                        continue;
                    if (lines[i][bitIndex] != lookFor)
                    {
                        accepted[i] = false;
                        acceptedLeft--;
                        continue;
                    }
                    Debug.WriteLine(lines[i]);
                    ratingStr = lines[i];
                }
                Debug.WriteLine("");
                if (acceptedLeft == 1)
                    break;
            }
            var rating = Convert.ToInt32(ratingStr, 2);
            Debug.WriteLine($"rating [{rating}]");
            return rating;
        }
    }
}

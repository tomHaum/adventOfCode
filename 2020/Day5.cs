using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCodeCSharp
{
    class Day5
    {
        public static void Run()
        {
            Console.WriteLine("Day 5");

            var boardingPasses = data.day5.Split(Environment.NewLine);
            //boardingPasses = new string[] { "FBFBBFFRLR" };

            var max = 0;
            var list = new List<int>();
            foreach(var s in boardingPasses)
            {
                var b = s.Replace('F', '0')
                    .Replace('B', '1')
                    .Replace('R', '1')
                    .Replace('L', '0');

                var rowStr = b.Substring(0, 7);
                var colStr = b.Substring(7);

                var row = Convert.ToInt32(rowStr, 2);
                var col = Convert.ToInt32(colStr, 2);

                var seatId = row * 8 + col;
                list.Add(seatId);
                //Console.WriteLine($"{s} | {b} | {row} | {col} | {seatId}");
                if (seatId > max) max = seatId;
            }
            Console.WriteLine("Part 1: " + max);

            var last = list[0];
            foreach(var seat in list.OrderBy(x=>x))
            {
                //Console.WriteLine($"{last} | {seat}");
                if(seat - 2 == last)
                {
                    break;
                }
                last = seat;
            }
            Console.WriteLine("Part 2: "+ (last+1));
        }
    }
}

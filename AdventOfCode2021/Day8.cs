using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace AdventOfCode2021
{
    internal class Day8
    {

        /*
          
  0:      1:      2:      3:      4:   
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
          
         **/

        internal static void Run()
        {
            var lines = input.day8.Split("\r\n");

            int accum = 0;
            foreach (var line in lines)
            {
                var split = line.Split("|");
                var unique = split[0].Trim().Split(' ').OrderBy(x => x.Length).ToList();
                var output = split[1].Trim().Split(' ');

                foreach (var digit in output)
                {
                    // 1,7,4,8
                    if (digit.Length == 2 || digit.Length == 3 || digit.Length == 4 || digit.Length == 7)
                    {
                        accum++;
                    }
                }
            }

            Console.WriteLine($"Day 8 - Part 1: {accum}");

            accum = 0;
            foreach (var line in lines)
            {
                var split = line.Split("|");
                var unique = split[0].Trim().Split(' ').OrderBy(x => x.Length).ToList();
                var output = split[1].Trim().Split(' ');

                var one = unique.Single(x => x.Length == 2);
                var seven = unique.Single(x => x.Length == 3);
                var four = unique.Single(x => x.Length == 4);
                var eight = unique.Single(x => x.Length == 7);

                var top = seven.Except(one).ToList()[0];
                var six = unique.Where(x => x.Length == 6).Single(x => x.Intersect(one).Count() == 1);
                var bottomRight = one.Intersect(six).Single();
                var topRight = one.Single(x => x != bottomRight);

                var two = unique.Where(x => x.Length == 5).Single(x => x.Contains(topRight) && !x.Contains(bottomRight));
                var five = unique.Where(x => x.Length == 5).Single(x => !x.Contains(topRight) && x.Contains(bottomRight));
                var three = unique.Where(x => x.Length == 5).Single(x => x.Contains(topRight) && x.Contains(bottomRight));

                var bottomLeft = two.Except(five).Single(x => x != topRight);
                var zero = unique.Where(x => x.Length == 6 && x != six).Single(x => x.Contains(bottomLeft));
                var nine = unique.Single(x => x.Length == 6 && x != six && x != zero);


                //Console.WriteLine("zero  | " + zero);
                //Console.WriteLine("one   | " + one);
                //Console.WriteLine("two   | " + two);
                //Console.WriteLine("three | " + three);
                //Console.WriteLine("four  | " + four);
                //Console.WriteLine("five  | " + five);
                //Console.WriteLine("six   | " + six);
                //Console.WriteLine("seven | " + seven);
                //Console.WriteLine("eight | " + eight);
                //Console.WriteLine("nine  | " + nine);


                
                var lineNumber = 0;
                foreach (string digit in output.Where(x => x != null))
                {
                    int curr = 0;

                    if (zero.Length == digit.Length && zero.Intersect(digit).Count() == digit.Length)
                        curr = 0;
                    if (one.Length == digit.Length && one.Intersect(digit).Count() == digit.Length)
                        curr = 1;
                    if (two.Length == digit.Length && two.Intersect(digit).Count() == digit.Length)
                        curr = 2;
                    if (three.Length == digit.Length && three.Intersect(digit).Count() == digit.Length)
                        curr = 3;
                    if (four.Length == digit.Length && four.Intersect(digit).Count() == digit.Length)
                        curr = 4;
                    if (five.Length == digit.Length && five.Intersect(digit).Count() == digit.Length)
                        curr = 5;
                    if (six.Length == digit.Length && six.Intersect(digit).Count() == digit.Length)
                        curr = 6;
                    if (seven.Length == digit.Length && seven.Intersect(digit).Count() == digit.Length)
                        curr = 7;
                    if (eight.Length == digit.Length && eight.Intersect(digit).Count() == digit.Length)
                        curr = 8;
                    if (nine.Length == digit.Length && nine.Intersect(digit).Count() == digit.Length)
                        curr = 9;

                    lineNumber = lineNumber * 10 + curr;
                }
                Console.WriteLine("lineNumber " + lineNumber);
                accum += lineNumber;
            }
            Console.WriteLine($"Day 8 - Part 2: {accum}");


        }

    }
}
